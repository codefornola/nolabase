---
id: qgis
title: QGIS
sidebar_label: QGIS
slug: /clients/qgis
---

[QGIS](https:/qgis.org) is free and open source geographic information system
software that you can connect to the nolabase, to make maps, integrate with other spatial data,
or run advanced spatial analysis operations.

## Setup

After [installing QGIS](https://www.qgis.org/en/site/forusers/download.html),
open a new project and in the [Browser](https://docs.qgis.org/3.16/en/docs/user_manual/introduction/browser.html)
panel you'll see a number of database and service connection options,
among them `PostGIS`.

![QGIS Browser Panel](/img/qgis-browser-panel.png)

Right-click `PostGIS` to create a New Connection, and enter your nolabase
connection information:

![QGIS Create PostGIS Connection](/img/qgis-new-postgis-connection.png)

Leave the credentials blank and you'll be prompted for them when you
make (or test) the connection.

## Getting Started

Once connected, you'll get a list of the spatial tables in the nolabase, each with
an icon indicating what type of geometry (point, line, or polygon) it includes.
Just drag one into the main QGIS workspace to begin working with it.

You'll see the table appear as a new layer in the `Layers` panel. Right-click on
the layer and choose `Open Attribute Table` to inspect the data in tabular form.

Use the [Layer Styling Panel](https://docs.qgis.org/3.16/en/docs/user_manual/introduction/general_tools.html#layer-styling-panel)
to begin changing the style of features in your layer based on data in the
attribute table.

Use the [QuickMapServices](https://nextgis.com/blog/quickmapservices/) plugin to
get access to a wide variety of basemaps. To install, go to `Plugins` -> `Manage
and Install Plugins...`. Adding a basemap from this plugin (like Open Street Map
or Stamen Toner) allows you to to see nolabase data against a familiar backdrop.

## Resources

* [QGIS Official Documentation](https://docs.qgis.org/3.16/en/docs/user_manual/)
to learn more about how to use the software.
* [Getting Started with QGIS](https://geoservices.leventhalmap.org/cartinal/guides/get-started-qgis/)
from the Leventhal Map & Education Center (Boston Public Library) is a nice
introduction to the software. They provide other good GIS-related guides as well.

QGIS has an active worldwide user (and developer) community. Some good places to
access it are

* [QGIS on Twitter](https://twitter.com/qgis) for news and events.
* [QGIS on Github](https://github.com/qgis/qgis) to track issues and development.
* [QGIS N.A. User Group](http://qgis.us/) one of many user groups around the world.
