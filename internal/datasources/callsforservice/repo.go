package callsforservice

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

func (r *Repo) StoreCalls(calls []*ServiceCall) error {
	ctx := context.Background()
	tx, err := r.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	for _, c := range calls {
		sql := `
		INSERT INTO
			cfs.calls_for_service (
				nopd_item,
				type_text,
				priority,
				initial_type,
				initial_type_text,
				initial_priority,
				disposition,
				disposition_text,
				beat,
				block_address,
				zip,
				police_district,
				self_initiated,
				time_create,
				time_dispatch,
				time_closed,
				time_arrive,
				lng_lat_point_nad83
			)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
				$11, $12, $13, $14, $15, $16, $17, %s)
		ON CONFLICT (nopd_item) DO NOTHING
		RETURNING id;
        `
		var id int
		// the format could be NAD83 or WGS84, we want to transform the WGS84 ones to NAD83
		if c.LngLatPoint.SRID == 4326 {
			sql = fmt.Sprintf(sql, "ST_Transform(GeomFromEWKB($18), 3452)")
		} else {
			sql = fmt.Sprintf(sql, "GeomFromEWKB($18)")
		}
		err = tx.QueryRow(ctx, sql,
			c.NopdItem,
			c.TypeText,
			c.Priority,
			c.InitialType,
			c.InitialTypeText,
			c.InitialPriority,
			c.Disposition,
			c.DispositionText,
			c.Beat,
			c.BlockAddress,
			c.Zip,
			c.PoliceDistrict,
			c.SelfInitiated,
			c.TimeCreate.Time,
			c.TimeDispatch.Time,
			c.TimeClosed.Time,
			c.TimeArrive.Time,
			c.LngLatPoint,
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
