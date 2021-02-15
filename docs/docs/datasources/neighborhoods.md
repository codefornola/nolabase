---
id: neighborhoods
title: Neighborhoods
sidebar_label: Neighborhoods
slug: /datasources/neighborhoods
---

## Overview

The statistical neighborhood geometries.

[Source](https://portal-nolagis.opendata.arcgis.com/datasets/neighborhood-statistical-areas)

## Tables

### neighborhoods

```
  Column  |          Type          | Collation | Nullable |                       Default
----------+------------------------+-----------+----------+------------------------------------------------------
 id       | integer                |           | not null | nextval('geometries.neighborhoods_id_seq'::regclass)
 name     | character varying(254) |           |          |
 objectid | numeric                |           |          |
 geom     | geometry(Polygon,4326) |           |          |
Indexes:
    "neighborhoods_pkey" PRIMARY KEY, btree (id)
    "geometries_neighborhoods_geom_index" gist (geom)
    "geometries_neighborhoods_name_index" UNIQUE, btree (name)
```


