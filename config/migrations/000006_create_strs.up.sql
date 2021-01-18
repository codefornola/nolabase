CREATE SCHEMA str;
COMMENT ON SCHEMA str IS 'All Orleans STR permit applications from data.nola.gov: https://data.nola.gov/Housing-Land-Use-and-Blight/Short-Term-Rental-Permit-Applications/en36-xvxg';

CREATE TABLE str.permits (
    id SERIAL PRIMARY KEY,
	reference_code varchar(6),
    license_number varchar(50),
    license_type varchar(50),
    residential_subtype varchar(50),
    address varchar(255),
    current_status varchar(20),
	expired boolean,
	link text,
	license_holder_name varchar(255),
    bedroom_limit int,
    guest_occupancy_limit int,
    operator_permit_number varchar(255),
    contact_name varchar(255),
    contact_phone varchar(50),
    contact_email varchar(255),
    application_date timestamp without time zone NOT NULL,
    expiration_date timestamp without time zone NOT NULL,
    issue_date timestamp without time zone NOT NULL,
    /* write to this column and a trigger will autoupdate the 4326 column */
    lng_lat_point_nad83 geometry(Point,3452),
    lng_lat_point geometry(Point,4326),
    inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
);
COMMENT ON TABLE str.permits IS 'The permits';

CREATE INDEX str_lng_lat_point_index ON str.permits  USING GIST (lng_lat_point);
CREATE INDEX str_lng_lat_point_nad83_index ON str.permits USING GIST (lng_lat_point_nad83);

CREATE TRIGGER str_wgs84_transformer BEFORE INSERT OR UPDATE ON str.permits
    FOR EACH ROW
    EXECUTE PROCEDURE wgs84_transform();
