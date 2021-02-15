---
slug: hello-world
title: Hello World
author: bhelx
author_url: https://github.com/bhelx
author_image_url: https://github.com/bhelx.png?size=200
tags: []
---

I'd like to start this blog by writing about the motivation behind this project. 
Hopefully, it can be a helpful introduction as well as a lodestar for the contributors moving forward.

## How We Got Here

As long as I've been programming professionally, I've been using my free time and knowledge to help
with projects that I think are beneficial to the city of New Orleans. Known as "civic hacking", this
type of advocacy has been around for a long time and is the founding principle behind
[Code For New Orleans](http://codeforneworleans.org/).

Along the way I've worked with locals doing a diverse amount of work in fields like
journalism, crime, politics, taxes, housing affordability, public transportation, city governance, government
accountability, and others.

The work I do generally involves building tools or helping with data analysis. The one constant problem I encounter
is using and understanding public data. Here in New Orleans, we are lucky to have some quality sources
of public data in machine-readable form. However, it was not always that way, and it's still far from
perfect. Even with good data, I started to notice something was missing. I found that nearly every project I
worked on had a number of barriers to entry that often required a programmer to solve.

## The Problem

In each project, the primary barriers to entry were around extracting the data (esp. if it's not in a machine-readble
form) and transforming and loading all the unique datasources into a tool for analysis. I also noticed
that every local group doing this type of work is basically doing this prep work over and over again
in their own silos. And because each process is unique, when they finish their analysis, they have a unique
and complicated set of steps to reproduce the results. My idea was to solve all of these problems one
time as a community, and then let the community worry about the parts they actually care about
(doing analysis and building helpful apps and tools).

## Nolabase

The nolabase is a community-shared [Postgres](https://www.postgresql.org/) database that contains all
known public information about the city. The data is already clean and always kept up to date by the community.
And combining sources of data is as simple as performing a [SQL join](https://en.wikipedia.org/wiki/Join_(SQL)) across 
as many tables as you want. Querying the data involves writing a [SQL](https://en.wikipedia.org/wiki/SQL) query
which makes sharing easy. This means reproducing someone's results, or copying and modifying their analysis
for your own purposes, is as simple as copying and pasting the queries. And since
many tools speak SQL, you aren't necessarily locked into using the nolabase and writing SQL directly.
It also operates as a rich and stable platform for building tools by offloading the complicated data
integrity and computation problems to our team.

## Goals

Although the nolabase is a tool, our focus is on the community around it. The nolabase provides
the community with an **interface** to agree on which enables us to work together and help each other.
We hope that we never have to solve these problems alone again.


