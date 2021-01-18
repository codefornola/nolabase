CREATE TABLE geometries.voting_precincts (
    id SERIAL PRIMARY KEY,
    name varchar(254),
    registered_voters int,
    registered_voters_white int,
    registered_voters_black int,
    registered_voters_other int,
    registered_democrats int,
    registered_democrats_white int,
    registered_democrats_black int,
    registered_republicans int,
    registered_republicans_white int,
    registered_republicans_black int,
    registered_republicans_other int,
    registered_other_parties int,
    registered_other_parties_white int,
    registered_other_parties_black int,
    registered_other_parties_other int,
    geom geometry(POLYGON,4326),
    inserted_at timestamp without time zone NOT NULL DEFAULT NOW()
);
COMMENT ON TABLE geometries.voting_precincts IS 'The city voting precincts and registered voter info https://portal-nolagis.opendata.arcgis.com/datasets/total-of-registered-voters';

CREATE INDEX geometries_voting_precincts_geom_index ON geometries.voting_precincts USING GIST (geom);
CREATE UNIQUE INDEX geometries_voting_precincts_name_index ON geometries.voting_precincts USING btree (name);

