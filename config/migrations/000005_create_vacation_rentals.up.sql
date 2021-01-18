CREATE SCHEMA vacation_rentals;
COMMENT ON SCHEMA vacation_rentals IS 'Orleans parish vacation rentals (Hotels, B&B, short-term rentals, etc.) from data.nola.gov: https://data.nola.gov/Housing-Land-Use-and-Blight/Vacation-Rentals-Hotels-B-B-short-term-rentals-etc/rbhq-zbz9';

CREATE TABLE vacation_rentals.properties (
    id SERIAL PRIMARY KEY,
    name varchar(100),
    address_name varchar(255),
    type varchar(100),
    bedroom_limit int,
    guest_limit int,
    expiration_date timestamp without time zone,
    /* write to this column and a trigger will autoupdate the 4326 column */
    lng_lat_point_nad83 geometry(Point,3452),
    lng_lat_point geometry(Point,4326),
    inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
);
COMMENT ON TABLE vacation_rentals.properties IS 'All properties with their limitations and expiration dates.';

CREATE INDEX vacation_rentals_lng_lat_point_index ON vacation_rentals.properties USING GIST (lng_lat_point);
CREATE INDEX vacation_rentals_lng_lat_point_nad83_index ON vacation_rentals.properties USING GIST (lng_lat_point_nad83);

CREATE TRIGGER vacation_rentals_wgs84_transformer BEFORE INSERT OR UPDATE ON vacation_rentals.properties
    FOR EACH ROW
    EXECUTE PROCEDURE wgs84_transform();
