CREATE TABLE norta.stop_times
(
  stop_id           varchar(10) NOT NULL,
  trip_id           varchar(10) NOT NULL,
  arrival_time      interval NOT NULL,
  departure_time    interval NOT NULL,
  stop_sequence     integer NOT NULL,
  stop_headsign     varchar(255) NULL,
  pickup_type       integer NULL CHECK(pickup_type >= 0 and pickup_type <=3),
  drop_off_type     integer NULL CHECK(drop_off_type >= 0 and drop_off_type <=3),
  inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
);

CREATE INDEX norta_stop_times_stop_id_index ON norta.stop_times USING btree (stop_id);
CREATE INDEX norta_stop_times_trip_id_index ON norta.stop_times USING btree (trip_id);

CREATE TABLE norta.agency
(
  agency_id         varchar(20) PRIMARY KEY,
  agency_name       varchar(255) NOT NULL,
  agency_url        varchar(255) NOT NULL,
  agency_timezone   varchar(255) NOT NULL,
  agency_lang       varchar(255) NULL,
  agency_phone      varchar(255) NULL,
  inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
);

CREATE TABLE norta.calendar
(
  service_id        varchar(10) PRIMARY KEY,
  monday            boolean NOT NULL,
  tuesday           boolean NOT NULL,
  wednesday         boolean NOT NULL,
  thursday          boolean NOT NULL,
  friday            boolean NOT NULL,
  saturday          boolean NOT NULL,
  sunday            boolean NOT NULL,
  start_date        timestamp NOT NULL,
  end_date          timestamp NOT NULL,
  inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
);

-- CREATE TABLE norta.calendar_dates
-- (
--   service_id varchar(10) NOT NULL,
--   date numeric(8) NOT NULL,
--   exception_type integer NOT NULL,
--   inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
-- );

-- CREATE INDEX norta_calendar_dates_service_id_index ON norta.calendar_dates USING btree (service_id);

CREATE TABLE norta.frequencies
(
  trip_id           varchar(10) NOT NULL,
  start_time        interval NOT NULL,
  end_time          interval NOT NULL,
  headway_secs      integer NOT NULL,
  inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
);

CREATE INDEX norta_frequencies_trip_id_index ON norta.frequencies USING btree (trip_id);

-- CREATE TABLE norta.transfers
-- (
--     from_stop_id  varchar(255) NOT NULL,
--     to_stop_id    varchar(255) NOT NULL,
--     transfer_type   integer NOT NULL,
--     min_transfer_time integer,
--     inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
-- );