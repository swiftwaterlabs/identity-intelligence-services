# Identity Intelligence Services
The Identity Intelligence Services solution is a collection of components that enable deep insights into identities in use at any enterprise organization, indepdendent of identity providers used.  By leveraging the capabilities of big data offerings to quickly and easily analyze large amounts of data, a comprehensive set of insights across all users and groups can be obtained across all identity providers used.  

# Supported Directory Providers
The following directory providers are supported:

|Name|Status|
|---|---|
|LDAP / Active Directory Domains|Available|
|Okta Organizations|In Development|
|Google Workspaces|Planned|

# How It Works
For every onboarded directory, data about each user and group is queried in an incremental manner and pushed to a secure Amazon S3 bucket for storage.  Once the data is collected, items in the S3 buckets are wrapped by big data tools such as Hive (Amazon Glue Database) and Presto (Amazon Athena) so they can be queried in real time.

Using queries defined in the [analytics/queries](analytics/queries) directory, this data can then exported to providers such as Tableau to gain insights from the raw data.  A sample set of analytical workbooks are available at [analytics/tableau/workbooks](analytics/tableau/workbooks/) to showcase what is possible.

# Available Insights
Using the sample Tableau workbooks, when the Identity Intelligence Services are used the following insights can be gained:

## Users
* What user types are present and how many of each are there?
* How many active, terminated, or disabled users are in my directories?
* How many users have a valid Manager assigned to them?
* What are the most common job titles?
* How many Non Human Accounts are there?
  * What is the breakdown between different types of Non Human Accounts?
  * How many active or disabled accounts are there?
  * How many accounts have a valid sponsor defined (and who are the ones that do not)?
  * What departments in the enterprise manage / sponsor the most accounts?
  * Who sponsors / owns the most accounts in ther enterprise?
  * For a given person, what accounts are they a sponsor of?

## Groups
* What types of groups are present and how many of each are there?
* Are there groups with no members?  If so, how many are there and what type are they?
* Which groups have the largest number of members?
* Which directory entities belong to the most groups?  Filterable by type to investigate possible anomalies
* For a given user or group, what groups are they a member of?

# How To Use
The Identity Intelligence Services solution is broken down into two pieces - data collection and analytics.

## Data Collection
To install, configure, and collect data so it is available for analysis, follow information in the [src/ReadMe](src/README.md)

## Analytics
When data is collected, use the [analytics/queries](analytics/queries) to extract data either locally or to a Tableau Server so it is available for the pre-built dashboards at [analytics/tableau/workbooks](analytics/tableau/workbooks/).

# License
This projects is made available under the [MIT License](LICENSE).

