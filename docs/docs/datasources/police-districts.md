---
id: police-districts
title: Police Districts
sidebar_label: Police Districts
slug: /datasources/police-districts
---

## Overview

The police district boundaries.

[Source](https://portal-nolagis.opendata.arcgis.com/datasets/nopd-police-zones)

## Tables

### geometries.police_districts

```
   Column    |            Type             | Collation | Nullable |                         Default
-------------+-----------------------------+-----------+----------+---------------------------------------------------------
 id          | integer                     |           | not null | nextval('geometries.police_districts_id_seq'::regclass)
 zone        | character varying(10)       |           |          |
 district    | character varying(20)       |           |          |
 geom        | geometry(Polygon,4326)      |           |          |
 inserted_at | timestamp without time zone |           | not null | now()
Indexes:
    "police_districts_pkey" PRIMARY KEY, btree (id)
    "geometries_police_districts_geom_index" gist (geom)
    "geometries_police_districts_name_index" UNIQUE, btree (zone)
```


