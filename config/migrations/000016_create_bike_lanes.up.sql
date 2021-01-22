CREATE TABLE geometries.bike_lanes (
    object_id integer PRIMARY KEY,
    install_year varchar(10) NULL,
    install_quarter varchar(2) NULL,
    two_way boolean NULL,
    neutral_ground boolean NULL,
    divided_roadway boolean NULL,
    plan_source varchar(20) NULL,
    facility_type varchar(50) NULL,
    status varchar(20) NULL,
    geom geometry(MULTILINESTRING,4326)
);

CREATE INDEX geometries_bike_lanes_geom_index ON geometries.bike_lanes USING GIST (geom);
