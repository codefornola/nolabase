---
id: short-term-rentals
title: Short Term Rentals
sidebar_label: Short Term Rentals
slug: /datasources/ short-term-rentals
---

## Overview

All permit applications for properties to be used as short-term rentals.

[Source](https://data.nola.gov/Housing-Land-Use-and-Blight/Short-Term-Rental-Permit-Applications/en36-xvxg)

## Tables

### str_permits

```
         Column         |            Type             | Collation | Nullable |                 Default
------------------------+-----------------------------+-----------+----------+-----------------------------------------
 id                     | integer                     |           | not null | nextval('str.permits_id_seq'::regclass)
 reference_code         | character varying(6)        |           |          |
 license_number         | character varying(50)       |           |          |
 license_type           | character varying(50)       |           |          |
 residential_subtype    | character varying(50)       |           |          |
 address                | character varying(255)      |           |          |
 current_status         | character varying(20)       |           |          |
 expired                | boolean                     |           |          |
 link                   | text                        |           |          |
 license_holder_name    | character varying(255)      |           |          |
 bedroom_limit          | integer                     |           |          |
 guest_occupancy_limit  | integer                     |           |          |
 operator_permit_number | character varying(255)      |           |          |
 contact_name           | character varying(255)      |           |          |
 contact_phone          | character varying(50)       |           |          |
 contact_email          | character varying(255)      |           |          |
 application_date       | timestamp without time zone |           | not null |
 expiration_date        | timestamp without time zone |           | not null |
 issue_date             | timestamp without time zone |           | not null |
 lng_lat_point_nad83    | geometry(Point,3452)        |           |          |
 lng_lat_point          | geometry(Point,4326)        |           |          |
 inserted_at            | timestamp without time zone |           | not null | now()
Indexes:
    "permits_pkey" PRIMARY KEY, btree (id)
    "str_lng_lat_point_index" gist (lng_lat_point)
    "str_lng_lat_point_nad83_index" gist (lng_lat_point_nad83)
Triggers:
    str_wgs84_transformer BEFORE INSERT OR UPDATE ON str.permits FOR EACH ROW EXECUTE FUNCTION wgs84_transform()
```



