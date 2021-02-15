---
id: council-districts
title: City Council Districts
sidebar_label: City Council Districts
slug: /datasources/council-districts
---

## Overview

The city council district boundaries.

[Source](https://portal-nolagis.opendata.arcgis.com/datasets/4593a994e7644bcc91d9e1c096df1734_0)

## Tables

### geometries.council_districts

```
   Column    |            Type             | Collation | Nullable |                         Default
-------------+-----------------------------+-----------+----------+----------------------------------------------------------
 id          | integer                     |           | not null | nextval('geometries.council_districts_id_seq'::regclass)
 name        | character varying(254)      |           |          |
 district_id | character varying(2)        |           |          |
 rep_name    | character varying(255)      |           |          |
 objectid    | numeric                     |           |          |
 authority   | character varying(255)      |           |          |
 last_update | timestamp without time zone |           |          |
 geom        | geometry(Polygon,4326)      |           |          |
 inserted_at | timestamp without time zone |           | not null | now()
Indexes:
    "council_districts_pkey" PRIMARY KEY, btree (id)
    "geometries_council_districts_geom_index" gist (geom)
    "geometries_council_districts_name_index" UNIQUE, btree (name)
```


