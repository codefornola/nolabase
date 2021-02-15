package shorttermrentals

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
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

func (r *Repo) StorePermits(permits []*Permit) error {
	ctx := context.Background()
	tx, err := r.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	for _, p := range permits {
		sql := `
		INSERT INTO
			str_permits (
				license_number,
				license_type,
				residential_subtype,
				address,
				current_status,
				expired,
				link,
				license_holder_name,
				bedroom_limit,
				guest_occupancy_limit,
				reference_code,
				operator_permit_number,
				contact_name,
				contact_phone,
				contact_email,
				application_date,
				expiration_date,
				issue_date,
				lng_lat_point_nad83
			)
		VALUES (
			$1, $2, $3, $4, $5, $6,
			$7, $8, $9, $10, $11,
			$12, $13, $14, $15, $16, $17, $18, %s)
		RETURNING id;
        `
		var id int
		// the format could be NAD83 or WGS84, we want to transform the WGS84 ones to NAD83
		if p.LngLatPoint.SRID == 4326 {
			sql = fmt.Sprintf(sql, "ST_Transform(GeomFromEWKB($19), 3452)")
		} else {
			sql = fmt.Sprintf(sql, "GeomFromEWKB($19)")
		}
		err = tx.QueryRow(ctx, sql,
			p.LicenseNumber,
			p.LicenseType,
			p.ResidentialSubtype,
			p.Address,
			p.CurrentStatus,
			p.Expired,
			p.Link,
			p.LicenseHolderName,
			p.BedroomLimit,
			p.GuestOccupancyLimit,
			p.ReferenceCode,
			p.OperatorPermitNumber,
			p.ContactName,
			p.ContactPhone,
			p.ContactEmail,
			p.ApplicationDate.Time,
			p.ExpirationDate.Time,
			p.IssueDate.Time,
			p.LngLatPoint,
		).Scan(&id)
		if err != nil && err != pgx.ErrNoRows {
			tx.Rollback(ctx)
			return err
		}
	}

	err = tx.Commit(ctx)
	return err
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
