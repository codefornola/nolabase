CREATE TABLE geometries.police_districts (
    id SERIAL PRIMARY KEY,
    zone varchar(10),
    district varchar(20),
    geom geometry(POLYGON,4326),
    inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
);
COMMENT ON TABLE geometries.police_districts IS 'The police district boundaries https://portal-nolagis.opendata.arcgis.com/datasets/nopd-police-zones';

CREATE INDEX geometries_police_districts_geom_index ON geometries.police_districts USING GIST (geom);
CREATE UNIQUE INDEX geometries_police_districts_name_index ON geometries.police_districts USING btree (zone);
