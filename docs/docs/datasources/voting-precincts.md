---
id: voting-precincts 
title: Voting Precincts
sidebar_label: Voting Precincts
slug: /datasources/voting-precincts
---

## Overview

The city voting precincts and registered voter info.

[Source](https://portal-nolagis.opendata.arcgis.com/datasets/total-of-registered-voters)

## Tables

### geometries.voting_precincts

```
             Column             |            Type             | Collation | Nullable |                         Default
--------------------------------+-----------------------------+-----------+----------+---------------------------------------------------------
id                             | integer                     |           | not null | nextval('geometries.voting_precincts_id_seq'::regclass)
name                           | character varying(254)      |           |          |
registered_voters              | integer                     |           |          |
registered_voters_white        | integer                     |           |          |
registered_voters_black        | integer                     |           |          |
registered_voters_other        | integer                     |           |          |
registered_democrats           | integer                     |           |          |
registered_democrats_white     | integer                     |           |          |
registered_democrats_black     | integer                     |           |          |
registered_republicans         | integer                     |           |          |
registered_republicans_white   | integer                     |           |          |
registered_republicans_black   | integer                     |           |          |
registered_republicans_other   | integer                     |           |          |
registered_other_parties       | integer                     |           |          |
registered_other_parties_white | integer                     |           |          |
registered_other_parties_black | integer                     |           |          |
registered_other_parties_other | integer                     |           |          |
geom                           | geometry(Polygon,4326)      |           |          |
inserted_at                    | timestamp without time zone |           | not null | now()
Indexes:
"voting_precincts_pkey" PRIMARY KEY, btree (id)
"geometries_voting_precincts_geom_index" gist (geom)
"geometries_voting_precincts_name_index" UNIQUE, btree (name)
```



