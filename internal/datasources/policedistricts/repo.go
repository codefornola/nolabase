package policedistricts

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

func (r *Repo) StorePoliceDistricts(districts *PoliceDistricts) error {
	ctx := context.Background()
	tx, err := r.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	for _, feature := range districts.Features.Features {
		sql := `
		INSERT INTO
			police_districts (
				zone,
				district,
				geom
			)
		VALUES ($1, $2, $3);
		`
		ewkb, err := ewkbhex.Encode(feature.Geometry, ewkbhex.NDR)
		if err != nil {
			return err
		}
		_, err = tx.Exec(ctx,
			sql,
			feature.Properties["Zone"],
			feature.Properties["District"],
			ewkb,
		)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)
	return err
}
