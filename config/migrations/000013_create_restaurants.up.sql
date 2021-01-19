CREATE SCHEMA restaurants;
COMMENT ON SCHEMA restaurants IS 'All New Orleans Restaurants';

CREATE TABLE restaurants.records (
    id SERIAL PRIMARY KEY,
    address varchar(255),
    business_name varchar(255),
    business_type varchar(100),
    city varchar(100),
    owner_name varchar(255),
    phone_number varchar(20),
    state varchar(2),
    suite varchar(20),
    -- lng_lat_point_nad83 geometry(Point,3452),
    lng_lat_point geometry(Point,4326),
    inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
);
COMMENT ON TABLE restaurants.records IS 'The records: https://portal-nolagis.opendata.arcgis.com/datasets/restaurants';

CREATE INDEX restaurants_lng_lat_point_index ON restaurants.records USING GIST (lng_lat_point);
--CREATE INDEX restaurants_lng_lat_point_nad83_index ON restaurants.records USING GIST (lng_lat_point_nad83);
-- CREATE TRIGGER restaurants_wgs84_transformer BEFORE INSERT OR UPDATE ON restaurants.records
--     FOR EACH ROW
--     EXECUTE PROCEDURE wgs84_transform();
