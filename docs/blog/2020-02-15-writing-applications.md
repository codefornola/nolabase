---
slug: writing-applications
title: Writing Applications
author: bhelx
author_url: https://github.com/bhelx
author_image_url: https://github.com/bhelx.png?size=200
tags: []
---

Although the nolabase is primarily used to empower data analysts, it can also support
application developers. I've mentioned before how I see it as a platform
for both groups of people. Here I want to demonstrate how you can use it to build an
application by re-writing an existing one to use the nolabase.

## Neighborhood Annotation

In 2016, I was reading an article from local crime analyst [Jeff Asher](https://twitter.com/Crimealytics)
in which he mentioned using a datasource from [data.nola.gov](https://data.nola.gov)
which had locations (lat/longs), but not the neighborhood the record was in. So his inability to
group the data by neighborhood was impeding more granular analysis.

The nolabase would be a perfect solution to this, but it didn't exist yet. So I quickly wrote a
[custom python script to help out](https://github.com/bhelx/nola-neighborhood-annotation).
Over time the scope grew. We added datasets and features. This was all working fine, but the
problem was that we couldn't get all of these tools to do geospatial computation compiling on
his Windows machine. Even if I could have gotten it to compile, I'd have to regularly ship him any updated
data or fixes to the data. I was looking to spend less time on this, not more. I ended up
writing a web application to run the script for him in a sandbox which increased the amount of time
and now money I was spending on it. We also didn't have a community formed (like Code for New Orleans)
to help out.

## Using the Nolabase

If we re-write this to use the nolabase, the user of the script now only needs python and a postgres client (psycopg2 in 
this case) installed which should be no problem on a Windows machine. So we could easily compile an exe and even make a simple
GUI for windows users. And now that the data is stored remotely, it can
always stay up to date. Furthermore, we can expand it to give the user an option to choose any geometry table to join the
data with (not just neighborhoods). I've left those extra features as an exercise to the reader though:

```python
#!/usr/local/bin/python
# -*- coding: utf-8 -*-

# This demonstrates using the nolabase as a dependency to a community run application or tool.
# Run it by passing a connection string and an input and output csv file:
# ```
# # get 2021 calls for service data as a test
# curl https://data.nola.gov/resource/3pha-hum9.csv > calls.csv
# python3 annotate.py postgresql://nolabaseuser:nolabasepassword@nolabase.codeforneworleans.org/nolabase calls.csv output.csv
# ```

import sys
import csv
import re
import psycopg2
from urllib.parse import urlparse

lat_lng_rg = re.compile('.*?([+-]?\\d*\\.\\d+)(?![-+0-9\\.]).*?([+-]?\\d*\\.\\d+)(?![-+0-9\\.])')

def parse_lat_lng(lat_lng_string):
    """
    Turns the Location column into (lat, lng) floats
    May look like this "(29.98645605, -90.06910049)"
    May have degree symbol "(29.98645605°,-90.06910049°)"
    """
    m = lat_lng_rg.search(lat_lng_string)

    if m:
        return (float(m.group(1)), float(m.group(2)))
    else:
        return (None, None)

# If I were to use this in production, I'd change this to be a batch
# call to the database instead of doing one row at a time
def find_neighborhood(conn, lat, lng):
    cur = conn.cursor()
    sql = """
            SELECT name FROM geometries.neighborhoods AS n WHERE
            st_within(ST_GeomFromText('POINT(%f %f)', 4326), n.geom) LIMIT 1
            """ % (lat, lng)
    cur.execute(sql)
    rows = cur.fetchone()
    if rows is None or len(rows) <= 0:
        return "N/A"
    return rows[0]

def annotate_csv(conn, in_file, out_file):
    """
    Goes row by row through the in_file and
    writes out the row to the out_file with
    the new Neighbhorhood column
    """

    reader = csv.reader(in_file)
    writer = csv.writer(out_file)

    # Write headers first, add new neighborhood column
    headers = next(reader)
    headers.append('Neighborhood')

    writer.writerow(headers)

    for row in reader:
        # WGS84 point, "Location" column, is last element
        lat, lng = parse_lat_lng(row[-1])

        if lat and lng:
            neighborhood = find_neighborhood(conn, lat, lng)
        else:
            neighborhood = 'N/A'

        row.append(neighborhood)
        writer.writerow(row)

        print("#%s lat: %s lng: %s -> %s" % (reader.line_num, lat, lng,
            neighborhood))


def print_help():
    help = """
    Usage:
    python annotate.py postgresql://username:password@nolabase.codefornola.org/nolabase input.csv output.csv
    """
    print(help)


if __name__ == '__main__':

    if len(sys.argv) < 3:
        print_help()
        sys.exit()
 
    parts = urlparse(sys.argv[1])

    conn = psycopg2.connect(
        user=parts.username,
        password=parts.password,
        dbname=parts.path[1:],
        host=parts.hostname,
        port=parts.port
    )

    in_file_path = sys.argv[2]
    out_file_path = sys.argv[3]

    with open(in_file_path, 'r') as in_file:
        with open(out_file_path, 'w') as out_file:
            annotate_csv(conn, in_file, out_file)
```
