package restaurants

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

func (r *Repo) StoreRestaurantss(restaurants *Restaurants) error {
	ctx := context.Background()
	tx, err := r.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	for _, feature := range restaurants.Features.Features {
		sql := `
		INSERT INTO
			restaurants (
				address,
				business_name,
				business_type,
				city,
				owner_name,
				phone_number,
				state,
				suite,
				lng_lat_point
			)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
		`
		ewkb, err := ewkbhex.Encode(feature.Geometry, ewkbhex.NDR)
		if err != nil {
			return err
		}
		_, err = tx.Exec(ctx,
			sql,
			feature.Properties["Address"],
			feature.Properties["BusinessName"],
			feature.Properties["BusinessType"],
			feature.Properties["City"],
			feature.Properties["OwnerName"],
			feature.Properties["PhoneNumber"],
			feature.Properties["State"],
			feature.Properties["Suite"],
			ewkb,
		)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)
	return err
}
