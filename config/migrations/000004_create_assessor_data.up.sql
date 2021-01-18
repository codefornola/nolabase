CREATE SCHEMA assessor;
COMMENT ON SCHEMA assessor IS 'Tax assessor data scraped from the Orleans Parish tax assessor website: https://www.qpublic.net/la/orleans/';

CREATE TABLE assessor.properties (
    id SERIAL PRIMARY KEY,
    assessor_id character varying(50),
    owner_name text,
    land_area_sq_ft integer,
    location_address text,
    mailing_address text,
    property_class character varying(255),
    municipal_district character varying(255),
    assessment_area character varying(255),
    tax_bill_number character varying(255),
    /* write to this column and a trigger will autoupdate the 4326 column */
    lng_lat_point_nad83 geometry(Point,3452),
    lng_lat_point geometry(Point,4326),
    parcel_no character varying(255),
    building_area_sq_ft integer,
    inserted_at timestamp without time zone NOT NULL DEFAULT NOW(),
    scraped_at timestamp without time zone NULL
);
COMMENT ON TABLE assessor.properties IS 'All Orleans properties / parcels';

CREATE UNIQUE INDEX properties_assessor_id_index ON assessor.properties USING btree (assessor_id);
CREATE INDEX properties_tax_bill_number_index ON assessor.properties USING btree (tax_bill_number);
CREATE INDEX assessor_lng_lat_point_index ON assessor.properties USING GIST (lng_lat_point);
CREATE INDEX assessor_lng_lat_point_nad83_index ON assessor.properties USING GIST (lng_lat_point_nad83);

CREATE TRIGGER assessor_properties_wgs84_transformer BEFORE INSERT OR UPDATE ON assessor.properties
    FOR EACH ROW
    EXECUTE PROCEDURE wgs84_transform();

CREATE TABLE assessor.property_sales (
    id SERIAL PRIMARY KEY,
    property_id integer,
    date timestamp without time zone,
    price integer,
    grantor character varying(255),
    grantee character varying(255),
    notarial_archive_number character varying(255),
    instrument_number character varying(255),
    inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
);
COMMENT ON TABLE assessor.property_sales IS 'All Orleans sales of properties / parcels';

CREATE TABLE assessor.property_values (
    id SERIAL PRIMARY KEY,
    property_id integer,
    year integer,
    land_value integer,
    building_value integer,
    total_value integer,
    assessed_land_value integer,
    assessed_building_value integer,
    total_assessed_value integer,
    homestead_exemption_value integer,
    taxable_assessment integer,
    neighborhood_id integer,
    age_freeze integer,
    disability_freeze integer,
    assment_change integer,
    tax_contract integer,
    inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
);
COMMENT ON TABLE assessor.property_values IS 'All Orleans tax valuations of properties / parcels';
