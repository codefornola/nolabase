CREATE TABLE geometries.school_districts (
    id SERIAL PRIMARY KEY,
    name varchar(254),
    geom geometry(POLYGON,4326),
    inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
);
COMMENT ON TABLE geometries.school_districts IS 'The school board district boundaries https://portal-nolagis.opendata.arcgis.com/datasets/school-board-districts';

CREATE INDEX geometries_school_districts_geom_index ON geometries.school_districts USING GIST (geom);
CREATE UNIQUE INDEX geometries_school_districts_name_index ON geometries.school_districts USING btree (name);

