- [Accessing the Nolabase](#accessing-the-nolabase)
  - [Getting an Account](#getting-an-account)
  - [Install a Client](#install-a-client)
  - [Connecting](#connecting)
  - [Writing a Query](#writing-a-query)
- [Resources for Learning](#resources-for-learning)
  - [Learn about SQL](#learn-about-sql)
  - [Learn about PostGIS](#learn-about-postgis)
- [DataSources](#datasources)
  - [Geometries](#geometries)
  - [Tax Assessor Data (properties sales and values)](#tax-assessor-data-properties-sales-and-values)
  - [Calls for Service](#calls-for-service)
  - [Vacation Rentals](#vacation-rentals)
  - [Short Term Rentals](#short-term-rentals)
  - [Restaurants](#restaurants)

## Accessing the Nolabase

### Getting an Account

The easiest way to get started querying the database is to connect to the community instance. You'll need a
username and password to connect. Currently, the only method to get an account is to ask for one in our Slack channel. Our channel is [#nolabase](https://nola.slack.com/archives/C01K1TBMRFA) in the "NOLA Devs" slack workspace. If you aren't already a memeber, you'll need to join using the following steps:

1. Enter your email in the [auto-invite tool](https://nola-slackin.herokuapp.com/) and follow the emailed instructions to get access to the workspace.
2. Once you are in, head to the [#nolabase](https://nola.slack.com/archives/C01K1TBMRFA) channel and say hi.

> *Note*:
> You can run the Nolabase locally on your computer, but this isn't supported yet. If you are interested in doing this, ask the community for help.

### Install a Client

The Nolabase is a [Postgres](https://www.postgresql.org/) database and thus you'll need a client to connect to it.
For an easy and user-friendly option, we recommend [sqlectron](https://sqlectron.github.io/). It's free, open source, and it works on all operating systems. Here is a more [complete list of clients](https://wiki.postgresql.org/wiki/PostgreSQL_Clients) that are supported. You can also connect to it from any programming language or other tool that can speak SQL.

### Connecting

Here are the details you'll need connect:

* *host* or *address*: `nolabase.codeforneworleans.org`
* *port*: `5432`
* *name* or *database*: `nolabase`
* *username*: `your-username`
* *password*: `your-password`
* *SSL*: `true` or `enabled`

> *Note*: Depending on your client, may need to specify that this is a `PostgreSQL` database. 

### Writing a Query

Try a test query to see that everything is working. This query
gives us the name of every neighborhood geometry in the database
sorted in the ascending direction:

```sql
SELECT name FROM geometries.neighborhoods ORDER BY name ASC;
```

Results:

```
            name
----------------------------
 ALGIERS POINT
 AUDUBON
 BAYOU ST. JOHN
 BEHRMAN
 BLACK PEARL
 .....
 (72 Rows)
```

## Resources for Learning

Here are some resources for learning the tools you'll likely need to use the database. Don't hesitate to reach out to the Slack community if you get stuck or need recommendations.

### Learn about SQL

[SQL](https://en.wikipedia.org/wiki/SQL) (**S**tructured **Q**uery **L**anguage) is a [domain specific language](https://en.wikipedia.org/wiki/Domain-specific_language) for querying and analyzing data. It's less powerful than a general purpose programming language and is thus much simpler and easier to learn.

[postgrestutorial.com](https://www.postgresqltutorial.com/) is a comprehensive guide if you want to understand SQL in depth. There are plenty of good courses to take and videos on youtube if you prefer to learn that way. When looking for places to learn, keep in mind that there are many *dialects* of SQL. They are mostly the same, but you'll probably have an easier time if the resource you are using is teaching with PostgreSQL.

### Learn about PostGIS

[PostGIS](https://postgis.net/) is a special extension for Postgres that enables you to query and analyze [geospatial information](https://en.wikipedia.org/wiki/Geographic_information_system) the same way you would normal tabluar information. You'll need it if your analysis needs to understand where things are in space.

The nolabase stores all geographic information using PostGIS types. e.g., we use `geometry.Point` for geographic locations and we use `gemoetry.MultiPolygon` for boundaries (like the shape of a neighborhood or a voter precinct). It also includes all the functions you need to query and do calculations on these types.

To learn PostGIS, see [this workshop](https://postgis.net/workshops/postgis-intro/) as an introduction.
See the [PostGIS documentation](https://postgis.net/docs/manual-3.1/) for reference.
Also, again, youtube and courses are good resources as well if you prefer to learn that way.

## DataSources 

DataSources are externally managed sources of data that are regularly pulled into the Nolabase. We try to keep the convention of creating a [postgres schema](https://www.postgresql.org/docs/9.1/ddl-schemas.html) for each DataSource in order to keep them isolated. If you aren't familiar with it, think of it as a namespace. This is where the tables, functions, triggers, etc live for that DataSource. The one place we break this convention is
`geometries` which is a special namespace for abstract geographic boundaries (think neighborhoods, police districts, etc) although I think we may change this before launching.

### Geometries

* Schema: `geometries`
  * Table: `neighborhoods`
    * Description: The statistical neighborhood areas.
    * DataSource: [NOLAGIS](https://portal-nolagis.opendata.arcgis.com/datasets/neighborhood-statistical-areas)
  * Table: `council_districts`
    * Description: City council districts and who represents them.
    * DataSource: [NOLAGIS](https://portal-nolagis.opendata.arcgis.com/datasets/4593a994e7644bcc91d9e1c096df1734_0)
  * Table: `voting_precincts`
    * Description: Voting precincts and registered voter demographic information.
    * DataSource: [NOLAGIS](https://portal-nolagis.opendata.arcgis.com/datasets/total-of-registered-voters)
  * Table: `school_districts`
    * Description: School Board Districts.
    * DataSource: [NOLAGIS](https://portal-nolagis.opendata.arcgis.com/datasets/school-board-districts)
  * Table: `police_districts`
    * Description: The NOPD Zones and Districts.
    * DataSource: [NOLAGIS](https://portal-nolagis.opendata.arcgis.com/datasets/nopd-police-zones)
  * Table: `police_subzones`
    * Description: The NOPD Subzones. The Subzone is the smallest level of jurisdiction used for reporting.
    * DataSource: [NOLAGIS](https://portal-nolagis.opendata.arcgis.com/datasets/nopd-police-subzones-reporting-districts)

### Tax Assessor Data (properties sales and values)

* Schema: `assessor`
  * Table: `properties`
    * Description: The details found in the `Owner and Parcel Information` table on the property page. Also has location data.
    * DataSource: [Tax Assessor Website](https://qpublic.net/la/orleans/)
  * Table: `property_sales`
    * Description: The details found in the `Value Information` table on the property page.
    * DataSource: [Tax Assessor Website](https://qpublic.net/la/orleans/)
  * Table: `property_values`
    * Description: The details found in the `Sale/Transfer Information` table on the property page.
    * DataSource: [Tax Assessor Website](https://qpublic.net/la/orleans/)

### Calls for Service

* Schema: `cfs`
  * Table: `calls_for_service`
    * Description: Calls for service from 2011 to today.
    * DataSource: [data.nola.gov](https://data.nola.gov/Public-Safety-and-Preparedness/Call-for-Service-2020/hp7u-i9hf)

### Vacation Rentals

* Schema: `vacation_rentals`
  * Table: `properties`
    * Description: A merged dataset of the Hotels, Motels, B&Bs, and Boarding Houses and the Short-Term Rentals datasets.
    * DataSource: [data.nola.gov](https://data.nola.gov/Housing-Land-Use-and-Blight/Vacation-Rentals-Hotels-B-B-short-term-rentals-etc/rbhq-zbz9)

### Short Term Rentals

* Schema: `str`
  * Table: `permits`
    * Description: Short term rental permits.
    * DataSource: [data.nola.gov](https://data.nola.gov/Housing-Land-Use-and-Blight/Short-Term-Rental-Permit-Applications/en36-xvxg)

### Restaurants

* Schema: `restaurants`
  * Table: `records`
    * Description: All restaurants in the parish.
    * DataSource: [NOLAGIS](https://portal-nolagis.opendata.arcgis.com/datasets/restaurants)


