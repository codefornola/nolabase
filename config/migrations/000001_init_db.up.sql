CREATE EXTENSION IF NOT EXISTS postgis;
COMMENT ON EXTENSION postgis IS 'PostGIS geometry, geography, and raster spatial types and functions';

CREATE SCHEMA infra;

CREATE TABLE infra.jobs (
                      id SERIAL PRIMARY KEY,
                      url text NOT NULL,
                      metadata text,
                      error text,
                      scraper_name varchar(255),
                      state int NOT NULL DEFAULT 0,
                      inserted_at timestamp without time zone NOT NULL DEFAULT NOW(),
                      UNIQUE (url, metadata)
);

CREATE INDEX jobs_state_index ON infra.jobs USING btree (state);

CREATE FUNCTION public.wgs84_transform() RETURNS trigger AS $wgs84_transform$
BEGIN
    IF NEW.lng_lat_point_nad83 IS NOT NULL
        THEN
            IF TG_OP = 'INSERT'
                THEN
                    NEW.lng_lat_point := st_transform(NEW.lng_lat_point_nad83, 4326);
                    RETURN NEW;
            ELSIF TG_OP = 'UPDATE' AND OLD.lng_lat_point_nad83 IS DISTINCT FROM NEW.lng_lat_point_nad83
                THEN
                    NEW.lng_lat_point := st_transform(NEW.lng_lat_point_nad83, 4326);
                    RETURN NEW;
            END IF;
    END IF;
    RETURN NEW;
END;
$wgs84_transform$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION notify_job() RETURNS TRIGGER AS $$

    DECLARE 
        data json;
        notification json;
    
    BEGIN
        IF (TG_OP != 'INSERT') THEN
            RETURN NULL;
        END IF;
        
        data = row_to_json(NEW);

        -- Contruct the notification as a JSON string.
        notification = json_build_object(
                          'table',TG_TABLE_NAME,
                          'action', TG_OP,
                          'data', data);
                        
        -- Execute pg_notify(channel, notification)
        PERFORM pg_notify('new_jobs',notification::text);
        
        -- Result is ignored since this is an AFTER trigger
        RETURN NULL; 
    END;
    
$$ LANGUAGE plpgsql;

CREATE TRIGGER new_job_notification
AFTER INSERT OR UPDATE OR DELETE ON infra.jobs
    FOR EACH ROW EXECUTE PROCEDURE notify_job();