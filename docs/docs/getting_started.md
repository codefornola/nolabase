---
id: getting_started
title: Getting Started
sidebar_label: Getting Started
slug: /getting-started
---

## Meet the Community

In order to access the nolabase, you need someone from the community
to setup an account. Currently, the only way to reach us is to join our Slack channel:
the [#nolabase](https://nola.slack.com/archives/C01K1TBMRFA) channel of the "NOLA Devs" Slack workspace.
This is a public Slack workspace with over 1000 New Orleanians from the tech community.

1. Enter your email in the [auto-invite tool](https://nola-slackin.herokuapp.com/) and follow the emailed instructions to get access to the workspace.
2. Once you logged in, navigate to the [#nolabase](https://nola.slack.com/archives/C01K1TBMRFA) channel and say helo.

![nolabase-slack-channel.png](/img/nolabase-slack-channel.png)

:::note
This is also a great place to share what you are working on and ask questions. Remember,
we are here to empower each other.
:::

## Connecting

When a community member sets up your account, they will send you two pieces of information,
a `username` and a generated `password`. This is your [Postgres](https://www.postgresql.org) user.

To connect to the database, you need the following settings:

* *host* or *address*: `nolabase.codeforneworleans.org`
* *port*: `5432`
* *name* or *database*: `nolabase`
* *username*: `your-username`
* *password*: `your-password`
* *SSL*: `true` or `enabled`

:::tip
If you aren't sure what to do with this information, or you are new to 
SQL databases in general, we recommend you jump to the [learning](learning) section.
:::

## Writing a Query


Try a test query to see that everything is working. This query
gives us the name of every neighborhood geometry in the database
sorted in the ascending direction:

```sql
SELECT name FROM neighborhoods ORDER BY name ASC;
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






