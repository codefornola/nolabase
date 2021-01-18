CREATE TABLE geometries.police_subzones (
    id SERIAL PRIMARY KEY,
    name varchar(10),
    zone varchar(10),
    district varchar(20),
    subzone varchar(20),
    geom geometry(MULTIPOLYGON,4326),
    inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
);
COMMENT ON TABLE geometries.police_subzones IS 'The police subzone boundaries. The smallest level of jurisdiction by NOPD https://portal-nolagis.opendata.arcgis.com/datasets/nopd-police-subzones-reporting-districts';

CREATE INDEX geometries_police_subzones_geom_index ON geometries.police_subzones USING GIST (geom);
CREATE UNIQUE INDEX geometries_police_subzones_name_index ON geometries.police_subzones USING btree (name);
