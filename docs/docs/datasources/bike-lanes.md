---
id: bike-lanes
title: Bike Lanes
sidebar_label: Bike Lanes
slug: /datasources/bike-lanes
---

## Overview

Completed bike lanes across the public street network of New Orleans. Updated regularly by DPW to reflect recently completed lanes. Attributed to show the type, for example a lane shared with cars or one protected from traffic.

[Source](https://data.nola.gov/Transportation-and-Infrastructure/Existing-Bike-Lanes/8npz-j6vy).

## Tables

### geometries.bike_lanes

```
     Column      |              Type              | Collation | Nullable | Default
-----------------+--------------------------------+-----------+----------+---------
 object_id       | integer                        |           | not null |
 install_year    | character varying(10)          |           |          |
 install_quarter | character varying(2)           |           |          |
 two_way         | boolean                        |           |          |
 neutral_ground  | boolean                        |           |          |
 divided_roadway | boolean                        |           |          |
 plan_source     | character varying(20)          |           |          |
 facility_type   | character varying(50)          |           |          |
 status          | character varying(20)          |           |          |
 geom            | geometry(MultiLineString,4326) |           |          |
Indexes:
    "bike_lanes_pkey" PRIMARY KEY, btree (object_id)
    "geometries_bike_lanes_geom_index" gist (geom)
```



