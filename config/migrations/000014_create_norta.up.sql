CREATE SCHEMA norta;
COMMENT ON SCHEMA restaurants IS 'New Orleans RTA data (public transportation). Follows the GTFS schema: https://developers.google.com/transit/gtfs';

CREATE TABLE norta.routes (
    route_id varchar(10) PRIMARY KEY,
	short_name varchar(255),
	long_name  varchar(255),
	description  varchar(255),
	type       int,
	url        varchar(255),
	color      varchar(20),
	text_color varchar(20),
    inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
);

CREATE TABLE norta.trips
(
  trip_id           varchar(10) NOT NULL PRIMARY KEY,
  route_id          varchar(10) NOT NULL,
  service_id        varchar(10) NOT NULL,
  trip_headsign     varchar(255) NULL,
  trip_short_name   varchar(255) NULL,
  direction_id      boolean NULL,
  block_id          varchar(255) NULL,
  shape_id          varchar(8) NULL,
  wheelchair_accessible boolean NULL,
  --bikes_allowed boolean NULL,
  inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
);
CREATE INDEX norta_trips_route_id_index ON norta.trips USING btree (route_id);
CREATE INDEX norta_trips_shape_id_index ON norta.trips USING btree (shape_id);

CREATE TABLE norta.shapes
(
  shape_id    varchar(8) PRIMARY KEY,
  geom        geometry(LINESTRING,4326),
  inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
);

CREATE INDEX norta_shapes_geom_index ON norta.shapes USING GIST (geom);

CREATE TABLE norta.stops
(
  stop_id           varchar(10) PRIMARY KEY,
  stop_code         varchar(255) NULL,
  stop_name         varchar(255) NOT NULL,
  stop_desc         varchar(255) NULL,
  zone_id           varchar(255) NULL,
  stop_url          varchar(255) NULL,
  location_type     boolean NULL,
  parent_station    varchar(255) NULL,
  wheelchair_boarding boolean NULL,
  lng_lat_point     geometry(Point,4326),
  inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
);
CREATE INDEX norta_stops_lng_lat_point_index ON norta.stops USING GIST (lng_lat_point);

