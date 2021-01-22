package bikelanes

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/twpayne/go-geom/encoding/ewkbhex"
)

type Repo struct {
	conn *pgxpool.Pool
}

func NewRepo(conn *pgxpool.Pool) *Repo {
	return &Repo{
		conn: conn,
	}
}

func (r *Repo) StoreBikeLanes(bikelanes *BikeLanes) error {
	ctx := context.Background()
	tx, err := r.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	for _, feature := range bikelanes.Features.Features {
		sql := `
		INSERT INTO
			geometries.bike_lanes (
				object_id,
				install_year,
				install_quarter,
				two_way,
				neutral_ground,
				divided_roadway,
				plan_source,
				facility_type,
				status,
				geom
			)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, ST_Multi($10));
		`
		ewkb, err := ewkbhex.Encode(feature.Geometry, ewkbhex.NDR)
		if err != nil {
			return err
		}
		_, err = tx.Exec(ctx,
			sql,
			feature.Properties["objectid"],
			feature.Properties["install_year"],
			feature.Properties["install_quarter"],
			feature.Properties["two_way"] == "1",
			feature.Properties["neutral_ground"] == "1" || feature.Properties["neutral_ground"] == "True",
			feature.Properties["divided_roadway"] == "1",
			feature.Properties["plan_source"],
			feature.Properties["facility_type"],
			feature.Properties["status"],
			ewkb,
		)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)
	return err
}
