---
id: tax-assessor
title: Tax Assessor
sidebar_label: Tax Assessor
slug: /datasources/tax-assessor
---

## Overview

This is a DataSource containing tax data on ~160,000 properties in New Orleans. I heard the assessor's office was unwilling
to share the data with the city's open data initiative so I'm releasing it here.  I've scraped this data from the [tax assessor's web site](http://nolaassessor.com/).

[Here](http://qpublic9.qpublic.net/la_orleans_display.php?KEY=935-GRAVIERST) is an example page of a single property.

![Screenshot of Assessors Site](/img/tax-assessor-screen.png)

## Tables

### assessor_properties

```
       Column        |            Type             | Collation | Nullable |                     Default
---------------------+-----------------------------+-----------+----------+-------------------------------------------------
 id                  | integer                     |           | not null | nextval('assessor.properties_id_seq'::regclass)
 assessor_id         | character varying(50)       |           |          |
 owner_name          | text                        |           |          |
 land_area_sq_ft     | integer                     |           |          |
 location_address    | text                        |           |          |
 mailing_address     | text                        |           |          |
 property_class      | character varying(255)      |           |          |
 municipal_district  | character varying(255)      |           |          |
 assessment_area     | character varying(255)      |           |          |
 tax_bill_number     | character varying(255)      |           |          |
 lng_lat_point_nad83 | geometry(Point,3452)        |           |          |
 lng_lat_point       | geometry(Point,4326)        |           |          |
 parcel_no           | character varying(255)      |           |          |
 building_area_sq_ft | integer                     |           |          |
 inserted_at         | timestamp without time zone |           | not null | now()
 scraped_at          | timestamp without time zone |           |          |
Indexes:
    "properties_pkey" PRIMARY KEY, btree (id)
    "assessor_lng_lat_point_index" gist (lng_lat_point)
    "assessor_lng_lat_point_nad83_index" gist (lng_lat_point_nad83)
    "properties_assessor_id_index" UNIQUE, btree (assessor_id)
    "properties_tax_bill_number_index" btree (tax_bill_number)
Triggers:
    assessor_properties_wgs84_transformer BEFORE INSERT OR UPDATE ON assessor.properties FOR EACH ROW EXECUTE FUNCTION wgs84_transform()
```

### assessor_property_sales

```
         Column          |            Type             | Collation | Nullable |                       Default
-------------------------+-----------------------------+-----------+----------+-----------------------------------------------------
 id                      | integer                     |           | not null | nextval('assessor.property_sales_id_seq'::regclass)
 property_id             | integer                     |           |          |
 date                    | timestamp without time zone |           |          |
 price                   | integer                     |           |          |
 grantor                 | character varying(255)      |           |          |
 grantee                 | character varying(255)      |           |          |
 notarial_archive_number | character varying(255)      |           |          |
 instrument_number       | character varying(255)      |           |          |
 inserted_at             | timestamp without time zone |           | not null | now()
Indexes:
    "property_sales_pkey" PRIMARY KEY, btree (id)
```

### assessor_property_values

```
          Column           |            Type             | Collation | Nullable |                       Default
---------------------------+-----------------------------+-----------+----------+------------------------------------------------------
 id                        | integer                     |           | not null | nextval('assessor.property_values_id_seq'::regclass)
 property_id               | integer                     |           |          |
 year                      | integer                     |           |          |
 land_value                | integer                     |           |          |
 building_value            | integer                     |           |          |
 total_value               | integer                     |           |          |
 assessed_land_value       | integer                     |           |          |
 assessed_building_value   | integer                     |           |          |
 total_assessed_value      | integer                     |           |          |
 homestead_exemption_value | integer                     |           |          |
 taxable_assessment        | integer                     |           |          |
 neighborhood_id           | integer                     |           |          |
 age_freeze                | integer                     |           |          |
 disability_freeze         | integer                     |           |          |
 assment_change            | integer                     |           |          |
 tax_contract              | integer                     |           |          |
 inserted_at               | timestamp without time zone |           | not null | now()
Indexes:
    "property_values_pkey" PRIMARY KEY, btree (id)
```
