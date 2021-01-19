- [Accessing the Nolabase](#accessing-the-nolabase)
  - [Getting an Account](#getting-an-account)
  - [Install a Client](#install-a-client)
  - [Connecting](#connecting)
  - [Writing a Query](#writing-a-query)
- [Resources for Learning](#resources-for-learning)
  - [Learn about SQL](#learn-about-sql)
  - [Learn about PostGIS](#learn-about-postgis)
- [DataSources](#datasources)

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

### Learn about SQL

* TODO

### Learn about PostGIS

* TODO

## DataSources 

DataSources are externally managed sources of data that are regularly pulled into the Nolabase. We try to keep the convention of creating a [postgres schema](https://www.postgresql.org/docs/9.1/ddl-schemas.html) for each DataSource in order to keep them isolated. If you aren't familiar with it, think of it as a namespace. This is where the tables, functions, triggers, etc live for that DataSource. The one place we break this convention is
`geometries` which is a special namespace for abstract geographic boundaries (think neighborhoods, police districts, etc) although I think we may change this before launching.

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
* Schema: `cfs`
  * Table: `calls_for_service`
    * Description: Calls for service from 2011 to today.
    * DataSource: [data.nola.gov](https://data.nola.gov/Public-Safety-and-Preparedness/Call-for-Service-2020/hp7u-i9hf)
* Schema: `vacation_rentals`
  * Table: `properties`
    * Description: A merged dataset of the Hotels, Motels, B&Bs, and Boarding Houses and the Short-Term Rentals datasets.
    * DataSource: [data.nola.gov](https://data.nola.gov/Housing-Land-Use-and-Blight/Vacation-Rentals-Hotels-B-B-short-term-rentals-etc/rbhq-zbz9)
* Schema: `str`
  * Table: `permits`
    * Description: Short term rental permits.
    * DataSource: [data.nola.gov](https://data.nola.gov/Housing-Land-Use-and-Blight/Short-Term-Rental-Permit-Applications/en36-xvxg)
* Schema: `restaurants`
  * Table: `records`
    * Description: All restaurants in the parish.
    * DataSource: [NOLAGIS](https://portal-nolagis.opendata.arcgis.com/datasets/restaurants)


