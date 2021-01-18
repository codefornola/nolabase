CREATE SCHEMA geometries;
COMMENT ON SCHEMA geometries IS 'Postgis geometries. Shapes related to the city.';

CREATE TABLE geometries.neighborhoods (
    id SERIAL PRIMARY KEY,
    name varchar(254),
    objectid numeric,
    geom geometry(POLYGON,4326) 
);
COMMENT ON TABLE geometries.neighborhoods IS 'The statistical neighborhood geometries https://portal-nolagis.opendata.arcgis.com/datasets/neighborhood-statistical-areas';

CREATE INDEX geometries_neighborhoods_geom_index ON geometries.neighborhoods USING GIST (geom);
CREATE UNIQUE INDEX geometries_neighborhoods_name_index ON geometries.neighborhoods USING btree (name);
