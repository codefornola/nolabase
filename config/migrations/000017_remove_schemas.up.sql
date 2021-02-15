ALTER SEQUENCE vacation_rentals.properties_id_seq RENAME TO vacation_rental_properties_id_seq;
ALTER INDEX vacation_rentals.properties_pkey RENAME TO vacation_rental_pkey;
ALTER TABLE vacation_rentals.properties SET SCHEMA public;
ALTER TABLE properties RENAME TO vacation_rentals;
DROP SCHEMA vacation_rentals;

ALTER TABLE str.permits SET SCHEMA public;
ALTER TABLE permits RENAME TO str_permits;
DROP SCHEMA str;

ALTER TABLE cfs.calls_for_service SET SCHEMA public;
DROP SCHEMA cfs;

ALTER TABLE restaurants.records SET SCHEMA public;
ALTER TABLE records RENAME TO restaurants;
DROP SCHEMA restaurants;

ALTER TABLE geometries.neighborhoods SET SCHEMA public;
ALTER TABLE geometries.council_districts SET SCHEMA public;
ALTER TABLE geometries.voting_precincts SET SCHEMA public;
ALTER TABLE geometries.school_districts SET SCHEMA public;
ALTER TABLE geometries.police_districts SET SCHEMA public;
ALTER TABLE geometries.police_subzones SET SCHEMA public;
ALTER TABLE geometries.bike_lanes SET SCHEMA public;
DROP SCHEMA geometries;

ALTER TABLE assessor.properties SET SCHEMA public;
ALTER TABLE properties RENAME TO assessor_properties;
ALTER TABLE assessor.property_sales SET SCHEMA public;
ALTER TABLE property_sales RENAME TO assessor_property_sales;
ALTER TABLE assessor.property_values SET SCHEMA public;
ALTER TABLE property_values RENAME TO assessor_property_values;
DROP SCHEMA assessor;

ALTER TABLE norta.routes RENAME TO norta_routes;
ALTER TABLE norta.norta_routes SET SCHEMA public;
ALTER TABLE norta.trips RENAME TO norta_trips;
ALTER TABLE norta.norta_trips SET SCHEMA public;
ALTER TABLE norta.shapes RENAME TO norta_shapes;
ALTER TABLE norta.norta_shapes SET SCHEMA public;
ALTER TABLE norta.stops RENAME TO norta_stops;
ALTER TABLE norta.norta_stops SET SCHEMA public;
ALTER TABLE norta.stop_times RENAME TO norta_stop_times;
ALTER TABLE norta.norta_stop_times SET SCHEMA public;
ALTER TABLE norta.agency RENAME TO norta_agency;
ALTER TABLE norta.norta_agency SET SCHEMA public;
ALTER TABLE norta.calendar RENAME TO norta_calendar;
ALTER TABLE norta.norta_calendar SET SCHEMA public;
ALTER TABLE norta.frequencies RENAME TO norta_frequencies;
ALTER TABLE norta.norta_frequencies SET SCHEMA public;
DROP SCHEMA norta;