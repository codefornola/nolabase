CREATE SCHEMA vacation_rentals;
ALTER TABLE vacation_rentals SET SCHEMA vacation_rentals;
ALTER TABLE vacation_rentals.vacation_rentals RENAME TO properties;
ALTER SEQUENCE vacation_rentals.vacation_rental_properties_id_seq RENAME TO properties_id_seq;
ALTER INDEX vacation_rentals.vacation_rental_pkey RENAME TO properties_pkey;

CREATE SCHEMA str;
ALTER TABLE str_permits SET SCHEMA str;
ALTER TABLE str.str_permits RENAME TO permits;

CREATE SCHEMA cfs;
ALTER TABLE calls_for_service SET SCHEMA cfs;

CREATE SCHEMA restaurants;
ALTER TABLE restaurants SET SCHEMA restaurants;
ALTER TABLE restaurants.restaurants RENAME TO records;

CREATE SCHEMA geometries;
ALTER TABLE neighborhoods SET SCHEMA geometries;
ALTER TABLE council_districts SET SCHEMA geometries;
ALTER TABLE voting_precincts SET SCHEMA geometries;
ALTER TABLE school_districts SET SCHEMA geometries;
ALTER TABLE police_districts SET SCHEMA geometries;
ALTER TABLE police_subzones SET SCHEMA geometries;
ALTER TABLE bike_lanes SET SCHEMA geometries;

CREATE SCHEMA assessor;
ALTER TABLE assessor_properties SET SCHEMA assessor;
ALTER TABLE assessor.assessor_properties RENAME TO properties;
ALTER TABLE assessor_property_sales SET SCHEMA assessor;
ALTER TABLE assessor.assessor_property_sales RENAME TO property_sales;
ALTER TABLE assessor_property_values SET SCHEMA assessor;
ALTER TABLE assessor.assessor_property_values RENAME TO property_values;

CREATE SCHEMA norta;
ALTER TABLE norta_routes SET SCHEMA norta;
ALTER TABLE norta.norta_routes RENAME TO routes;
ALTER TABLE norta_trips SET SCHEMA norta;
ALTER TABLE norta.norta_trips RENAME TO trips;
ALTER TABLE norta_shapes SET SCHEMA norta;
ALTER TABLE norta.norta_shapes RENAME TO shapes;
ALTER TABLE norta_stops SET SCHEMA norta;
ALTER TABLE norta.norta_stops RENAME TO stops;
ALTER TABLE norta_stop_times SET SCHEMA norta;
ALTER TABLE norta.norta_stop_times RENAME TO stop_times;
ALTER TABLE norta_agency SET SCHEMA norta;
ALTER TABLE norta.norta_agency RENAME TO agency;
ALTER TABLE norta_calendar SET SCHEMA norta;
ALTER TABLE norta.norta_calendar RENAME TO calendar;
ALTER TABLE norta_frequencies SET SCHEMA norta;
ALTER TABLE norta.norta_frequencies RENAME TO frequencies;