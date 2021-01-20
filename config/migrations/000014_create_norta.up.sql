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


-- CREATE TABLE norta.stop_times
-- (
--   stop_id           varchar(255) PRIMARY KEY,
--   trip_id           varchar(255) NOT NULL,
--   arrival_time      interval NOT NULL,
--   departure_time    interval NOT NULL,
--   stop_sequence     integer NOT NULL,
--   stop_headsign     varchar(255) NULL,
--   pickup_type       integer NULL CHECK(pickup_type >= 0 and pickup_type <=3),
--   drop_off_type     integer NULL CHECK(drop_off_type >= 0 and drop_off_type <=3),
--   inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
-- );

-- CREATE TABLE norta.agency
-- (
--   agency_id         varchar(255) UNIQUE NULL,
--   agency_name       varchar(255) NOT NULL,
--   agency_url        varchar(255) NOT NULL,
--   agency_timezone   varchar(255) NOT NULL,
--   agency_lang       varchar(255) NULL,
--   agency_phone      varchar(255) NULL,
--   inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
-- );
-- COMMENT ON TABLE norta.agency IS '';

-- CREATE TABLE norta.feed_info (
--   feed_publisher_name varchar(255) NOT NULL,
--   feed_publisher_url  varchar(255) NOT NULL,
--   feed_lang varchar(255) NOT NULL,
--   feed_start_date numeric(8) NULL,
--   feed_end_date numeric(8) NULL,
--   feed_version varchar(255) NULL,
--   inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
-- );
-- CREATE TABLE norta.calendar
-- (
--   service_id        varchar(255) PRIMARY KEY,
--   monday            boolean NOT NULL,
--   tuesday           boolean NOT NULL,
--   wednesday         boolean NOT NULL,
--   thursday          boolean NOT NULL,
--   friday            boolean NOT NULL,
--   saturday          boolean NOT NULL,
--   sunday            boolean NOT NULL,
--   start_date        numeric(8) NOT NULL,
--   end_date          numeric(8) NOT NULL,
--   inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
-- );

-- CREATE TABLE norta.calendar_dates
-- (
--   service_id varchar(255) NOT NULL,
--   date numeric(8) NOT NULL,
--   exception_type integer NOT NULL,
--   inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
-- );

-- CREATE TABLE norta.frequencies
-- (
--   trip_id           varchar(255) NOT NULL,
--   start_time        interval NOT NULL,
--   end_time          interval NOT NULL,
--   headway_secs      integer NOT NULL,
--   inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
-- );

-- CREATE TABLE norta.transfers
-- (
--     from_stop_id  varchar(255) NOT NULL,
--     to_stop_id    varchar(255) NOT NULL,
--     transfer_type   integer NOT NULL,
--     min_transfer_time integer,
--     inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
-- );
