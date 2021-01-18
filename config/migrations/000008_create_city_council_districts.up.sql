CREATE TABLE geometries.council_districts (
    id SERIAL PRIMARY KEY,
    name varchar(254),
    district_id varchar(2),
    rep_name varchar(255),
    objectid numeric,
    authority varchar(255),
    last_update timestamp without time zone,
    geom geometry(POLYGON,4326),
    inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
);
COMMENT ON TABLE geometries.council_districts IS 'The city council district boundaries https://portal-nolagis.opendata.arcgis.com/datasets/4593a994e7644bcc91d9e1c096df1734_0';

CREATE INDEX geometries_council_districts_geom_index ON geometries.council_districts USING GIST (geom);
CREATE UNIQUE INDEX geometries_council_districts_name_index ON geometries.council_districts USING btree (name);

