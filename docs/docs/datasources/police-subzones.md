---
id: police-subzones
title: Police Subzones
sidebar_label: Police Subzones
slug: /datasources/police-subzones
---

## Overview

The police subzone boundaries. The smallest level of jurisdiction by NOPD.

[Source](https://portal-nolagis.opendata.arcgis.com/datasets/nopd-police-subzones-reporting-districts)

## Tables

### police_subzones

```
   Column    |            Type             | Collation | Nullable |                        Default
-------------+-----------------------------+-----------+----------+--------------------------------------------------------
 id          | integer                     |           | not null | nextval('geometries.police_subzones_id_seq'::regclass)
 name        | character varying(10)       |           |          |
 zone        | character varying(10)       |           |          |
 district    | character varying(20)       |           |          |
 subzone     | character varying(20)       |           |          |
 geom        | geometry(MultiPolygon,4326) |           |          |
 inserted_at | timestamp without time zone |           | not null | now()
Indexes:
    "police_subzones_pkey" PRIMARY KEY, btree (id)
    "geometries_police_subzones_geom_index" gist (geom)
    "geometries_police_subzones_name_index" UNIQUE, btree (name)
```


