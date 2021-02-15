---
id: school-districts
title: School Districts
sidebar_label: School Districts
slug: /datasources/school-districts
---

## Overview

The school board district boundaries.

[Source](https://portal-nolagis.opendata.arcgis.com/datasets/school-board-districts)

## Tables

### geometries.school_districts

```
   Column    |            Type             | Collation | Nullable |                         Default
-------------+-----------------------------+-----------+----------+---------------------------------------------------------
 id          | integer                     |           | not null | nextval('geometries.school_districts_id_seq'::regclass)
 name        | character varying(254)      |           |          |
 geom        | geometry(Polygon,4326)      |           |          |
 inserted_at | timestamp without time zone |           | not null | now()
Indexes:
    "school_districts_pkey" PRIMARY KEY, btree (id)
    "geometries_school_districts_geom_index" gist (geom)
    "geometries_school_districts_name_index" UNIQUE, btree (name)
```


