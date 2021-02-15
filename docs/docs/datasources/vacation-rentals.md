---
id: vacation-rentals
title: Vacation Rentals
sidebar_label: Vacation Rentals
slug: /datasources/vacation-rentals
---

## Overview

A merged dataset of the Hotels, Motels, B&Bs, and Boarding Houses and the Short-Term Rentals datasets.

[Source](https://data.nola.gov/Housing-Land-Use-and-Blight/Vacation-Rentals-Hotels-B-B-short-term-rentals-etc/rbhq-zbz9)

## Tables

### vacation_rentals

```
Column        |            Type             | Collation | Nullable |                         Default
---------------------+-----------------------------+-----------+----------+---------------------------------------------------------
id                  | integer                     |           | not null | nextval('vacation_rentals.properties_id_seq'::regclass)
name                | character varying(100)      |           |          |
address_name        | character varying(255)      |           |          |
type                | character varying(100)      |           |          |
bedroom_limit       | integer                     |           |          |
guest_limit         | integer                     |           |          |
expiration_date     | timestamp without time zone |           |          |
lng_lat_point_nad83 | geometry(Point,3452)        |           |          |
lng_lat_point       | geometry(Point,4326)        |           |          |
inserted_at         | timestamp without time zone |           | not null | now()
Indexes:
"properties_pkey" PRIMARY KEY, btree (id)
"vacation_rentals_lng_lat_point_index" gist (lng_lat_point)
"vacation_rentals_lng_lat_point_nad83_index" gist (lng_lat_point_nad83)
Triggers:
vacation_rentals_wgs84_transformer BEFORE INSERT OR UPDATE ON vacation_rentals.properties FOR EACH ROW EXECUTE FUNCTION wgs84_transform()
```



