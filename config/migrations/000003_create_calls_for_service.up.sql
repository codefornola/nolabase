CREATE SCHEMA cfs;
COMMENT ON SCHEMA cfs IS 'The Orleans calls for service data from data.nola.gov. 2011 to present https://data.nola.gov/Public-Safety-and-Preparedness/Call-for-Service-2020/hp7u-i9hf';

CREATE TABLE cfs.calls_for_service (
    id SERIAL PRIMARY KEY,
    nopd_item varchar(50) UNIQUE NOT NULL,
    type_text varchar(100),
    priority varchar(10),
    initial_type varchar(10),
    initial_type_text varchar(100),
    initial_priority varchar(10),
    disposition varchar(50),
    disposition_text varchar(100),
    beat varchar(10),
    block_address varchar(255),
    zip varchar(10),
    police_district varchar(10),
    self_initiated boolean,
    time_create timestamp without time zone,
    time_dispatch timestamp without time zone,
    time_closed timestamp without time zone,
    time_arrive timestamp without time zone,
    /* write to this column and a trigger will autoupdate the 4326 column */
    lng_lat_point_nad83 geometry(Point,3452),
    lng_lat_point geometry(Point,4326),
    inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
);
COMMENT ON TABLE cfs.calls_for_service IS 'The service calls';

CREATE UNIQUE INDEX calls_for_service_nopd_item ON cfs.calls_for_service USING btree (nopd_item);
CREATE INDEX calls_for_service_lng_lat_point_index ON cfs.calls_for_service USING GIST (lng_lat_point);
CREATE INDEX calls_for_service_lng_lat_point_nad83_index ON cfs.calls_for_service USING GIST (lng_lat_point_nad83);

CREATE TRIGGER calls_for_service_wgs84_transformer BEFORE INSERT OR UPDATE ON cfs.calls_for_service
    FOR EACH ROW
    EXECUTE PROCEDURE wgs84_transform();