---
id: learning
title: Learning Resources
sidebar_label: Learning Resources
slug: /learning
---

## **What is SQL?**

SQL is often pronounced "sequel," but it stands for "structured query language." You can think of it as the computer equivalent of asking verbal questions.

Let's say you're interested in the Calls for Service (911 calls) dataset here on nolabase. You might want to know, "How many robberies have there been recently?” or “What's the most common emergency people call in for?"

If you were on the phone with someone, you would ask those questions verbally. On a computer, you would ask them using a query language like SQL. Luckily, there is a lot of overlap between the human and computer versions of these questions.

Typically, SQL is reserved for datasets that are too big to fit in spreadsheet programs like Google Sheets or Excel. It's popular among businesses, schools, software engineers, government agencies and others who need to present a large amount of information in a more manageable way.

If this is your first time using SQL, you can use this guide to write your first queries and get a sense of what you can do.


## **Getting ready**

To use SQL on datasets within nolabase, you need to connect with Excel, Tableau, or another method. If you're unsure how to do this, there are instructions on the nolabase site under the "Connecting" section.

With any dataset, the first thing you should do is get your bearings. On nolabase, the Overview page for each dataset shows the data that is available as well as important things to know. For instance, the Calls for Service (911 calls) data omits calls regarding juveniles. That could have a big effect on the outcome of your analysis, depending on what you're looking at.

There are many things that can go wrong with data - misspellings, empty rows, a strange column name you don't understand. To get an idea, you can use the [Quartz Guide to Bad Data](https://github.com/Quartz/bad-data-guide) to see dozens of the most common problems that crop up.

If you have any questions or are unsure, the most foolproof method is to contact the source of the data itself. With the Calls for Service dataset, that is the New Orleans Police Department. Many, if not most, sources will have a phone number or email you can contact with general questions.


## **SQL Queries**

To use SQL, you'll need to understand a few basic concepts and commands. The language uses a certain syntax, or word order, to converse with a database. The key word here is in the middle name of the language: query. In this case, a query is a question written in SQL syntax.

For example, this query asks, "Who has permits for short term rentals (AirBnBs) in New Orleans?"

```
SELECT contact_name
FROM str_permits
WHERE current_status = 'issued'
```

`SELECT` tells the computer what data you want to see. In this case, we're asking to see the licensees' names. If you want to see all the columns at once, you can type an asterisk `*` to mean "all."

"FROM" tells the computer where this data is coming from. In this case, our dataset is called str_permits, short for "short term rental permits."

"WHERE" is similar to its use in English, where you only want to see the data that meets certain criteria. In this example, we only want to see people with permits that have been issued, as opposed to denied, withdrawn or expired.

It's important to note that you must use the exact terms the computer uses. If you ask for data about "airbnb permits," for example, the computer won't know what you're talking about. You need to use "str_permits", the exact name of the dataset the computer uses.

You can find the table and column names in the Overview section of the datasource.

## **Analyzing Data**

The words at the beginning of each line - SELECT, FROM, and WHERE - are called "SQL commands." With a few basic commands, you can interrogate a database pretty deeply. 

"GROUP BY" groups rows into one row if they match somewhere. For instance, it groups all 911 calls for "arson" into the same row.

"ORDER BY" puts the output in a certain order, like alphabetical, chronological, or from smallest to largest.

For example, this query asks for the names of phone numbers of everyone with active permits in New Orleans (or rather, registered in the New Orleans database). But this time, it creates a single row for each person, and gives us the list in alphabetical order.

```
SELECT contact_name, contact_phone
FROM str_permits
WHERE current_status = issued
GROUP BY contact_name
ORDER BY contact_name ASC
```

Already, we've taken a large amount of information and arranged it in a more useful way. Out of tens of thousands of permit applicants, we found the ones we want to see and put them in order.

A lot of times, data analysis consists of doing math. One benefit to computers is that they can do math much more quickly than human beings. You can put mathematical functions like divide, add, count and average straight into your queries.

"COUNT" tells you the number of rows that meet a certain criteria. For example, you could count every STR permit in a certain neighborhood.

This one asks, "How many short term rental permits are expired?"

```
SELECT COUNT(license_number)
FROM str_permits
WHERE current_status = expired
```

In this query, "COUNT" is the mathematical function. It is listed in the first row, after SELECT, because what we want to see is the count of rows, not all the rows individually. In our first query, the SELECT command would show us every single row that met the criteria - hundreds or thousands of them.

There are many guides online to other [SQL commands](https://mode.com/sql-tutorial/sql-select-statement/). To practice more queries yourself, you can try the [W3Schools interactive tutorial](https://www.w3schools.com/sql/sql_syntax.asp). For more help, consider joining us on the [Nola Devs Slack channel](https://nola-slackin.herokuapp.com).
