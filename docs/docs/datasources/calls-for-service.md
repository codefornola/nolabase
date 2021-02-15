---
id: calls-for-service
title: Calls For Service
sidebar_label: Calls For Service
slug: /datasources/calls-for-service
---

## Overview

This dataset reflects incidents that have been reported to the New Orleans Police Department. Data is provided by Orleans Parish Communication District (OPCD), the administrative office of 9-1-1 for the City of New Orleans. Please request 911 audio via our public records request system here: https://nola.nextrequest.com.
In the OPCD system, NOPD may reclassify or change the signal type for up to 36 hours after the incident is marked up. For information about an incident after this time period, citizens may request police reports from the NOPD Public Records Division. In order to protect the privacy of victims, addresses are shown at the block level and the call types cruelty to juveniles, juvenile attachment and missing juvenile have been removed in accordance with the Louisiana Public Records Act, L.R.S. 44:1. Map coordinates (X,Y) have been removed for the following call types: Aggravated Rape, Aggravated Rape - MA, Crime Against Nature, Mental Patient, Oral Sexual Battery, Prostitution, Sexual Battery, Simple Rape, Simple Rape - Male V, and Soliciting for Prost.
Disclaimer: These incidents may be based upon preliminary information supplied to the Police Department by the reporting parties that have not been verified. The preliminary crime classifications may be changed at a later date based upon additional investigation and there is always the possibility of mechanical or human error. Therefore, the New Orleans Police Department does not guarantee (either expressed or implied) the accuracy, completeness, timeliness, or correct sequencing of the information and the information should not be used for comparison purposes over time. The New Orleans Police Department will not be responsible for any error or omission, or for the use of, or the results obtained from the use of this information. All data visualizations on maps should be considered approximate and attempts to derive specific addresses are strictly prohibited. The New Orleans Police Department is not responsible for the content of any off-site pages that are referenced by or that reference this web page other than an official City of New Orleans or New Orleans Police Department web page. The user specifically acknowledges that the New Orleans Police Department is not responsible for any defamatory, offensive, misleading, or illegal conduct of other users, links, or third parties and that the risk of injury from the foregoing rests entirely with the user. Any use of the information for commercial purposes is strictly prohibited. The unauthorized use of the words "New Orleans Police Department," "NOPD," or any colorable imitation of these words or the unauthorized use of the New Orleans Police Department logo is unlawful. This web page does not, in any way, authorize such use.

:::info
This table contains data from 2011 to present day
:::

## Tables

### calls_for_service

```
       Column        |            Type             | Collation | Nullable |                      Default
---------------------+-----------------------------+-----------+----------+---------------------------------------------------
 id                  | integer                     |           | not null | nextval('cfs.calls_for_service_id_seq'::regclass)
 nopd_item           | character varying(50)       |           | not null |
 type_text           | character varying(100)      |           |          |
 priority            | character varying(10)       |           |          |
 initial_type        | character varying(10)       |           |          |
 initial_type_text   | character varying(100)      |           |          |
 initial_priority    | character varying(10)       |           |          |
 disposition         | character varying(50)       |           |          |
 disposition_text    | character varying(100)      |           |          |
 beat                | character varying(10)       |           |          |
 block_address       | character varying(255)      |           |          |
 zip                 | character varying(10)       |           |          |
 police_district     | character varying(10)       |           |          |
 self_initiated      | boolean                     |           |          |
 time_create         | timestamp without time zone |           |          |
 time_dispatch       | timestamp without time zone |           |          |
 time_closed         | timestamp without time zone |           |          |
 time_arrive         | timestamp without time zone |           |          |
 lng_lat_point_nad83 | geometry(Point,3452)        |           |          |
 lng_lat_point       | geometry(Point,4326)        |           |          |
 inserted_at         | timestamp without time zone |           | not null | now()
Indexes:
    "calls_for_service_pkey" PRIMARY KEY, btree (id)
    "calls_for_service_lng_lat_point_index" gist (lng_lat_point)
    "calls_for_service_lng_lat_point_nad83_index" gist (lng_lat_point_nad83)
    "calls_for_service_nopd_item" UNIQUE, btree (nopd_item)
    "calls_for_service_nopd_item_key" UNIQUE CONSTRAINT, btree (nopd_item)
Triggers:
    calls_for_service_wgs84_transformer BEFORE INSERT OR UPDATE ON cfs.calls_for_service FOR EACH ROW EXECUTE FUNCTION wgs84_transform()
```

