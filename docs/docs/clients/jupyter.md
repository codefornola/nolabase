---
id: jupyter
title: Jupyter Notebooks
sidebar_label: Jupyter Notebooks
slug: /clients/jupyter
---

[Jupyter Notebooks](https://jupyter.org/) gives you the power of the [Python programming language](https://www.python.org/)
and all of it's associated libraries for data analysis and visualization.
However, it can be a steep learning curve for new users (unless you already know Python).

![Jupyter Notebooks](/img/jupyter-notebooks.png)

## Requirements

We recommend going through the instructions to install [JupyterLab](https://jupyter.org/install.html).
The one thing you'll need that doesn't come out of the box is a postgres client. We recommend [psycopg2](https://pypi.org/project/psycopg2/).
You can install this with `pip` or whatever package manager you are using.

Some optional libraries you may want to consider installing:

* [Pandas](https://pandas.pydata.org/)
  * For easy plotting and combining data.
  * You can pipe the results of a SQL query right into a `DataFrame` object.
* [Seaborn](https://seaborn.pydata.org/)
  * Useful visualization tools for statistical information
* [GeoPandas](https://geopandas.org/)
  * Like pandas, but includes support for spatial data (which the nolabase is full of)
* [contextily](https://github.com/geopandas/contextily)
  * Allows you to visualize data on top of a map

## Setup

Here's a pattern I've been using at the start of every notebook:

```python
import psycopg2
from getpass import getpass

password = getpass('Password: ')
con = psycopg2.connect(database="nolabase", user="benjamin.eckel", password=password,
    host="nolabase.codeforneworleans.org")
```

Creating a connection to the nolabase is the minimum you need to do to start querying data.
This code connects to the database and it prompts you for the password so you don't have to store the password
in your code where anyone can see it. You can now pass queries to the `con` object.

I also often find myself doing this pattern of loading up the result of a query into a GeoPandas
dataframe then plotting the result:

```python
import geopandas as gpd
import contextily as ctx

sql = """
SELECT 
  properties.neighborhood,
  SUM(permits.max_occupancy) AS max_occupancy,
  SUM(properties.livable_sqft) AS livable_sqft,
  CAST(SUM(permits.max_occupancy) AS decimal) / SUM(properties.livable_sqft) AS max_occupancy_per_sqft
FROM (
   SELECT
     n.geom AS neighborhood,
     greatest(SUM(guest_occupancy_limit), 0) AS max_occupancy
   FROM str_permits AS permits
   INNER JOIN voting_precincts AS n ON st_within(permits.lng_lat_point, n.geom)
   GROUP BY n.id
) AS permits
inner join (
   SELECT
     n.geom AS neighborhood,
     SUM(building_area_sq_ft) AS livable_sqft
   FROM assessor_properties AS properties
   INNER JOIN voting_precincts AS n ON st_within(properties.lng_lat_point, n.geom)
   WHERE properties.property_class = 'RESIDENTIAL'
   GROUP BY n.id
) AS properties
ON properties.neighborhood = permits.neighborhood
GROUP BY properties.neighborhood
ORDER BY max_occupancy_per_sqft DESC;
"""
# Execute the SQL query and load into a GeoDataFrame (GeoPandas)
df = gpd.GeoDataFrame.from_postgis(sql, con, geom_col='neighborhood')
# re-project to 3857 (web mercator) so it will fit on our basemap
df = df.to_crs(epsg=3857)
# Plot the data
ax = df.plot(column='max_occupancy_per_sqft', cmap='viridis', alpha=0.5, figsize=(15, 10))
# Add the basemap from contextily
ctx.add_basemap(ax, source=ctx.providers.Stamen.TonerLite)
```

Result:

![Jupyter Map Result](/img/jupyter-map.png)
