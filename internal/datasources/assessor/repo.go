package assessor

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"

	pgx "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repo struct {
	conn *pgxpool.Pool
}

func NewRepo(conn *pgxpool.Pool) *Repo {
	return &Repo{
		conn: conn,
	}
}

func (c *Repo) Close() {
	c.conn.Close()
}

func (c *Repo) StorePropertyPage(page *PropertyPage) error {
	ctx := context.Background()
	tx, err := c.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	sql := `
	UPDATE 
		assessor.properties
	SET 
		assessor_id = $1,
		owner_name = $2,
		land_area_sq_ft = $3,
		location_address = $4,
		mailing_address = $5,
		property_class = $6,
		assessment_area = $7,
		tax_bill_number = $8,
		parcel_no = $9,
		building_area_sq_ft = $10,
		lng_lat_point_nad83 = GeomFromEWKB($11)
	WHERE
		assessor_id = $1
	RETURNING id;
	`
	var propertyId int
	err = tx.QueryRow(ctx, sql,
		page.property.AssessorId,
		page.property.OwnerName,
		page.property.LandAreaSqFt,
		page.property.LocationAddress,
		page.property.MailingAddress,
		page.property.PropertyClass,
		page.property.AssessmentArea,
		page.property.TaxBillNumber,
		page.property.ParcelNo,
		page.property.BuildingAreaSqFt,
		page.property.LngLatPoint,
	).Scan(&propertyId)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	if len(page.values) > 0 {
		sql = `
	INSERT INTO
		assessor.property_values (
             property_id,
			 year,
			 land_value,
			 building_value,
			 total_value,
			 assessed_land_value,
			 assessed_building_value,
			 total_assessed_value,
			 homestead_exemption_value,
			 taxable_assessment,
			 age_freeze,
			 disability_freeze,
			 assment_change,
			 tax_contract
		)
	VALUES %s;
	`
		numVals := 14
		valueArgs := make([]interface{}, 0, len(page.values)*numVals)
		var valueStrings []string
		i := 1
		for _, v := range page.values {
			// templates
			values := strings.Join(generateValues(i, i+numVals-1), ",")
			i += numVals
			valueStrings = append(valueStrings, "("+values+")")
			// arguments
			valueArgs = append(valueArgs, propertyId)
			valueArgs = append(valueArgs, v.Year)
			valueArgs = append(valueArgs, v.LandValue)
			valueArgs = append(valueArgs, v.BuildingValue)
			valueArgs = append(valueArgs, v.TotalValue)
			valueArgs = append(valueArgs, v.AssessedLandValue)
			valueArgs = append(valueArgs, v.AssessedBuildingValue)
			valueArgs = append(valueArgs, v.TotalAssessedValue)
			valueArgs = append(valueArgs, v.HomesteadExemptionValue)
			valueArgs = append(valueArgs, v.TaxableAssessment)
			valueArgs = append(valueArgs, v.AgeFreeze)
			valueArgs = append(valueArgs, v.DisabilityFreeze)
			valueArgs = append(valueArgs, v.AssessmentChange)
			valueArgs = append(valueArgs, v.TaxContract)
		}
		sql = fmt.Sprintf(sql, strings.Join(valueStrings, ","))
		_, err = tx.Exec(ctx, sql, valueArgs...)
		if err != nil {
			return err
		}

	}

	if len(page.sales) > 0 {
		sql = `
	INSERT INTO
		assessor.property_sales (
			property_id,
			price,
			grantor,
			grantee,
			notarial_archive_number,
			instrument_number
		)
	VALUES %s;
	`
		numVals := 6
		valueArgs := make([]interface{}, 0, len(page.sales)*numVals)
		var valueStrings []string
		i := 1
		for _, s := range page.sales {
			// templates
			values := strings.Join(generateValues(i, i+numVals-1), ",")
			i += numVals
			valueStrings = append(valueStrings, "("+values+")")
			// arguments
			valueArgs = append(valueArgs, propertyId)
			valueArgs = append(valueArgs, s.Price)
			valueArgs = append(valueArgs, s.Grantor)
			valueArgs = append(valueArgs, s.Grantee)
			valueArgs = append(valueArgs, s.NotarialArchiveNumber)
			valueArgs = append(valueArgs, s.InstrumentNumber)
		}
		sql = fmt.Sprintf(sql, strings.Join(valueStrings, ","))
		_, err = tx.Exec(ctx, sql, valueArgs...)
		if err != nil {
			return err
		}
	}

	// finally, update the scraped_at time for this property
	sql = "UPDATE assessor.properties SET scraped_at = now() where id = $1;"
	_, err = tx.Exec(ctx, sql, propertyId)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	err = tx.Commit(ctx)
	return err
}

func (c *Repo) AllAssesorIds(action func(string) error) {
	sql := `
		SELECT
			assessor_id
		FROM assessor.properties;
	`
	rows, err := c.conn.Query(context.Background(), sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		err = action(id)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (c *Repo) StoreNewProperties(properties []*Property) error {
	valueStrings := make([]string, 0, len(properties))
	valueArgs := make([]interface{}, 0, len(properties))
	for i, prop := range properties {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d)", i+1))
		valueArgs = append(valueArgs, prop.AssessorId)
	}
	stmt := fmt.Sprintf("INSERT INTO assessor.properties(assessor_id) VALUES %s ON CONFLICT (assessor_id) DO NOTHING;",
		strings.Join(valueStrings, ","))

	_, err := c.conn.Exec(context.Background(), stmt, valueArgs...)
	return err
}

func (c *Repo) FindUnseenAssessorIds(ids []string) ([]string, error) {
	buf := bytes.NewBufferString("SELECT assessor_id FROM assessor.properties WHERE assessor_id IN(")
	for i, v := range ids {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString("'")
		buf.WriteString(v)
		buf.WriteString("'")
	}
	buf.WriteString(")")
	foundIds := make(map[string]bool)
	rows, err := c.conn.Query(context.Background(), buf.String())
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		foundIds[id] = true
	}
	var newIds []string
	for _, id := range ids {
		if _, ok := foundIds[id]; !ok {
			newIds = append(newIds, id)
		}
	}
	return newIds, nil
}

func (c *Repo) FindProperty(id string) (*Property, error) {
	property := &Property{}
	sql := `
		SELECT
			id,	
			assessor_id,
			owner_name,
			land_area_sq_ft,
			location_address,
			mailing_address,
			property_class,
			assessment_area,
			tax_bill_number,
			parcel_no,
			building_area_sq_ft,
			inserted_at,
			updated_at
		FROM assessor.properties
		WHERE assessor_id = $1
	`
	err := c.conn.QueryRow(
		context.Background(),
		sql, id,
	).Scan(
		&property.Id,
		&property.AssessorId,
		&property.OwnerName,
		&property.LandAreaSqFt,
		&property.LocationAddress,
		&property.MailingAddress,
		&property.PropertyClass,
		&property.AssessmentArea,
		&property.TaxBillNumber,
		&property.ParcelNo,
		&property.BuildingAreaSqFt,
		&property.InsertedAt,
		&property.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return property, nil
}

func generateValues(start, end int, extra ...string) (result []string) {
	for i := start; i <= end; i++ {
		result = append(result, fmt.Sprintf("$%d", i))
	}
	for _, v := range extra {
		result = append(result, v)
	}
	return result
}
