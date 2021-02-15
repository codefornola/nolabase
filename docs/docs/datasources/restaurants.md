---
id: restaurants
title: Restaurants
sidebar_label: Restaurants
slug: /datasources/restaurants
---

## Overview

Locations of restaurants throughout New Orleans, as indicated by occupational licenses.

[Source](https://data.nola.gov/dataset/Restaurants/yc3w-jdut)

## Tables

### restaurants

```
    Column     |            Type             | Collation | Nullable |                     Default
---------------+-----------------------------+-----------+----------+-------------------------------------------------
 id            | integer                     |           | not null | nextval('restaurants.records_id_seq'::regclass)
 address       | character varying(255)      |           |          |
 business_name | character varying(255)      |           |          |
 business_type | character varying(100)      |           |          |
 city          | character varying(100)      |           |          |
 owner_name    | character varying(255)      |           |          |
 phone_number  | character varying(20)       |           |          |
 state         | character varying(2)        |           |          |
 suite         | character varying(20)       |           |          |
 lng_lat_point | geometry(Point,4326)        |           |          |
 inserted_at   | timestamp without time zone |           | not null | now()
Indexes:
    "records_pkey" PRIMARY KEY, btree (id)
    "restaurants_lng_lat_point_index" gist (lng_lat_point)
```


