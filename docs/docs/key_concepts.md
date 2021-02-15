---
id: key_concepts
title: Key Concepts
sidebar_label: Key Concepts
slug: /key-concepts
---

## Datasources

Datasources are external sources of data our tools regularly pull from the web and put
into the nolabase.

## Table Layout

A datasource usually corresponds to one database table, but it may also have multiple tables.

## Spatial Data

Data is not just textual or numeric, we can also store "spatial" data. These are special types
that allow you to store and compute on things that exist in space, for example, bus stops,
voting precincts, etc. We utilize [PostGIS](https://postgis.net/) for these types.

## Walled Gardens

Our tools can pull data from anywhere. We often pull structured data from APIs and places
like [data.nola.gov](data.nola.gov), but we also have the capability to write [scrapers](https://en.wikipedia.org/wiki/Web_scraping)
which allow us to liberate data from sources who do not share their data in an accessible format.
The [Tax Assessor](/docs/datasources/tax-assessor) datasource is a good example of this.

