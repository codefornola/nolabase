package vacationrentals

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

func (r *Repo) StoreRentals(rentals []*VacationRental) error {
	ctx := context.Background()
	tx, err := r.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	for _, v := range rentals {
		sql := `
		INSERT INTO
			vacation_rentals (
				name,
				address_name,
				type,
				bedroom_limit,
				guest_limit,
				expiration_date,
				lng_lat_point_nad83
			)
		VALUES ($1, $2, $3, $4, $5, $6, %s)
		RETURNING id;
        `
		var id int
		// the format could be NAD83 or WGS84, we want to transform the WGS84 ones to NAD83
		if v.LngLatPoint.SRID == 4326 {
			sql = fmt.Sprintf(sql, "ST_Transform(GeomFromEWKB($7), 3452)")
		} else {
			sql = fmt.Sprintf(sql, "GeomFromEWKB($7)")
		}
		err = tx.QueryRow(ctx, sql,
			v.Name,
			v.AddressName,
			v.Type,
			v.BedroomLimit,
			v.GuestLimit,
			v.ExpirationDate.Time,
			v.LngLatPoint,
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
