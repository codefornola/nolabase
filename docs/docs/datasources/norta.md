---
id: norta
title: NORTA
sidebar_label: NORTA
slug: /datasources/norta
---

## Overview

New Orleans RTA data (public transportation). Follows the GTFS schema: https://developers.google.com/transit/gtfs.

[Source](https://www.norta.com/MyRTA/DataSubscription)

## Tables

### norta.routes

```
   Column    |            Type             | Collation | Nullable | Default
-------------+-----------------------------+-----------+----------+---------
 route_id    | character varying(10)       |           | not null |
 short_name  | character varying(255)      |           |          |
 long_name   | character varying(255)      |           |          |
 description | character varying(255)      |           |          |
 type        | integer                     |           |          |
 url         | character varying(255)      |           |          |
 color       | character varying(20)       |           |          |
 text_color  | character varying(20)       |           |          |
 inserted_at | timestamp without time zone |           | not null | now()
Indexes:
    "routes_pkey" PRIMARY KEY, btree (route_id)
```

### norta.trips

```
        Column         |            Type             | Collation | Nullable | Default
-----------------------+-----------------------------+-----------+----------+---------
 trip_id               | character varying(10)       |           | not null |
 route_id              | character varying(10)       |           | not null |
 service_id            | character varying(10)       |           | not null |
 trip_headsign         | character varying(255)      |           |          |
 trip_short_name       | character varying(255)      |           |          |
 direction_id          | boolean                     |           |          |
 block_id              | character varying(255)      |           |          |
 shape_id              | character varying(8)        |           |          |
 wheelchair_accessible | boolean                     |           |          |
 inserted_at           | timestamp without time zone |           | not null | now()
Indexes:
    "trips_pkey" PRIMARY KEY, btree (trip_id)
    "norta_trips_route_id_index" btree (route_id)
    "norta_trips_shape_id_index" btree (shape_id)
```

### norta.shapes

```
   Column    |            Type             | Collation | Nullable | Default
-------------+-----------------------------+-----------+----------+---------
 shape_id    | character varying(8)        |           | not null |
 geom        | geometry(LineString,4326)   |           |          |
 inserted_at | timestamp without time zone |           | not null | now()
Indexes:
    "shapes_pkey" PRIMARY KEY, btree (shape_id)
    "norta_shapes_geom_index" gist (geom)
```

### norta.stops

```
       Column        |            Type             | Collation | Nullable | Default
---------------------+-----------------------------+-----------+----------+---------
 stop_id             | character varying(10)       |           | not null |
 stop_code           | character varying(255)      |           |          |
 stop_name           | character varying(255)      |           | not null |
 stop_desc           | character varying(255)      |           |          |
 zone_id             | character varying(255)      |           |          |
 stop_url            | character varying(255)      |           |          |
 location_type       | boolean                     |           |          |
 parent_station      | character varying(255)      |           |          |
 wheelchair_boarding | boolean                     |           |          |
 lng_lat_point       | geometry(Point,4326)        |           |          |
 inserted_at         | timestamp without time zone |           | not null | now()
Indexes:
    "stops_pkey" PRIMARY KEY, btree (stop_id)
    "norta_stops_lng_lat_point_index" gist (lng_lat_point)
```

### norta.stop_times

```
     Column     |            Type             | Collation | Nullable | Default
----------------+-----------------------------+-----------+----------+---------
 stop_id        | character varying(10)       |           | not null |
 trip_id        | character varying(10)       |           | not null |
 arrival_time   | interval                    |           | not null |
 departure_time | interval                    |           | not null |
 stop_sequence  | integer                     |           | not null |
 stop_headsign  | character varying(255)      |           |          |
 pickup_type    | integer                     |           |          |
 drop_off_type  | integer                     |           |          |
 inserted_at    | timestamp without time zone |           | not null | now()
Indexes:
    "norta_stop_times_stop_id_index" btree (stop_id)
    "norta_stop_times_trip_id_index" btree (trip_id)
Check constraints:
    "stop_times_drop_off_type_check" CHECK (drop_off_type >= 0 AND drop_off_type <= 3)
    "stop_times_pickup_type_check" CHECK (pickup_type >= 0 AND pickup_type <= 3)
```

### norta.agency

```
     Column      |            Type             | Collation | Nullable | Default
-----------------+-----------------------------+-----------+----------+---------
 agency_id       | character varying(20)       |           | not null |
 agency_name     | character varying(255)      |           | not null |
 agency_url      | character varying(255)      |           | not null |
 agency_timezone | character varying(255)      |           | not null |
 agency_lang     | character varying(255)      |           |          |
 agency_phone    | character varying(255)      |           |          |
 inserted_at     | timestamp without time zone |           | not null | now()
Indexes:
    "agency_pkey" PRIMARY KEY, btree (agency_id)
```

### norta.calendar

```
   Column    |            Type             | Collation | Nullable | Default
-------------+-----------------------------+-----------+----------+---------
 service_id  | character varying(10)       |           | not null |
 monday      | boolean                     |           | not null |
 tuesday     | boolean                     |           | not null |
 wednesday   | boolean                     |           | not null |
 thursday    | boolean                     |           | not null |
 friday      | boolean                     |           | not null |
 saturday    | boolean                     |           | not null |
 sunday      | boolean                     |           | not null |
 start_date  | timestamp without time zone |           | not null |
 end_date    | timestamp without time zone |           | not null |
 inserted_at | timestamp without time zone |           | not null | now()
Indexes:
    "calendar_pkey" PRIMARY KEY, btree (service_id)
```


