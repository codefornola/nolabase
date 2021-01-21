package norta

import (
	"context"

	"github.com/geops/gtfsparser"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/twpayne/go-geom"
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

func (r *Repo) StoreGtfs(feed *gtfsparser.Feed) error {
	ctx := context.Background()
	tx, err := r.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	for _, route := range feed.Routes {
		sql := `
		INSERT INTO
			norta.routes (
				route_id,
				short_name,
				long_name,
				description,
				type,
				url,
				color,
				text_color
			)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
		`
		_, err = tx.Exec(ctx,
			sql,
			route.Id,
			route.Short_name,
			route.Long_name,
			route.Desc,
			route.Type,
			route.Url,
			route.Color,
			route.Text_color,
		)
		if err != nil {
			return err
		}
	}

	for _, trip := range feed.Trips {
		sql := `
		INSERT INTO
			norta.trips (
				trip_id,
				route_id,
				service_id,
				trip_headsign,
				trip_short_name,
				direction_id,
				block_id,
				shape_id,
				wheelchair_accessible
			)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
		`
		_, err = tx.Exec(ctx,
			sql,
			trip.Id,
			trip.Route.Id,
			trip.Service.Id,
			trip.Headsign,
			trip.Short_name,
			trip.Direction_id == 1,
			trip.Block_id,
			trip.Shape.Id,
			trip.Wheelchair_accessible == 1,
		)
		if err != nil {
			return err
		}
		for _, t := range trip.StopTimes {
			sql := `
			INSERT INTO
				norta.stop_times (
					stop_id,
					trip_id,
					arrival_time,
					departure_time,
					stop_sequence,
					stop_headsign,
					pickup_type,
					drop_off_type
				)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
			`
			_, err = tx.Exec(ctx,
				sql,
				t.Stop.Id,
				trip.Id,
				t.Arrival_time,
				t.Departure_time,
				t.Sequence,
				t.Headsign,
				t.Pickup_type,
				t.Drop_off_type,
			)
			if err != nil {
				return err
			}
		}
	}

	for _, shape := range feed.Shapes {
		var coords []float64
		for _, p := range shape.Points {
			// longitude then latitude (XY) layout
			coords = append(coords, float64(p.Lon))
			coords = append(coords, float64(p.Lat))
		}
		line := geom.NewLineStringFlat(geom.XY, coords)
		sql := `
		INSERT INTO
			norta.shapes (
				shape_id,
				geom
			)
		VALUES ($1, $2);
		`
		ewkb, err := ewkbhex.Encode(line, ewkbhex.NDR)
		_, err = tx.Exec(ctx,
			sql,
			shape.Id,
			ewkb,
		)
		if err != nil {
			return err
		}
	}

	for _, stop := range feed.Stops {
		sql := `
		INSERT INTO
			norta.stops (
				stop_id,
				stop_code,
				stop_name,
				stop_desc,
				zone_id,
				stop_url,
				location_type,
				parent_station,
				wheelchair_boarding,
				lng_lat_point
			)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);
		`
		point := geom.NewPointFlat(geom.XY, []float64{float64(stop.Lon), float64(stop.Lat)})
		ewkb, err := ewkbhex.Encode(point, ewkbhex.NDR)
		_, err = tx.Exec(ctx,
			sql,
			stop.Id,
			stop.Code,
			stop.Name,
			stop.Desc,
			stop.Zone_id,
			stop.Url,
			stop.Location_type == 1,
			stop.Parent_station,
			stop.Wheelchair_boarding == 1,
			ewkb,
		)
		if err != nil {
			return err
		}
	}

	for _, agency := range feed.Agencies {
		sql := `
		INSERT INTO
			norta.agency (
				agency_id,
				agency_name,
				agency_url,
				agency_timezone,
				agency_lang,
				agency_phone
			)
		VALUES ($1, $2, $3, $4, $5, $6);
		`
		_, err = tx.Exec(ctx,
			sql,
			agency.Id,
			agency.Name,
			agency.Url,
			agency.Timezone,
			agency.Lang,
			agency.Phone,
		)
		if err != nil {
			return err
		}
	}

	for _, service := range feed.Services {
		sql := `
		INSERT INTO
			norta.calendar (
				service_id,
				monday,
				tuesday,
				wednesday,
				thursday,
				friday,
				saturday,
				sunday,
				start_date,
				end_date
			)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);
		`
		_, err = tx.Exec(ctx,
			sql,
			service.Id,
			service.Daymap[1],
			service.Daymap[2],
			service.Daymap[3],
			service.Daymap[4],
			service.Daymap[5],
			service.Daymap[6],
			service.Daymap[0],
			service.Start_date.GetTime(),
			service.End_date.GetTime(),
		)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)
	return err
}
