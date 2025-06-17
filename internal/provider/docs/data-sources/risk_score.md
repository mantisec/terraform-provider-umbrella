---
page_title: "umbrella_risk_score Data Source - risk_score"
description: |-
  The Investigate Risk Score is based on an analysis of
the lexical characteristics of the domain name
and patterns in queries and requests to the domain.
The risk score is scaled from 0 to 100 where 100 is the highest risk
and 0 represents no risk at all. Periodically, Investigate updates this score based
on additional inputs.
A domain blocked by Umbrella receives a score of 100.
---

# umbrella_risk_score (Data Source)

The Investigate Risk Score is based on an analysis of
the lexical characteristics of the domain name
and patterns in queries and requests to the domain.
The risk score is scaled from 0 to 100 where 100 is the highest risk
and 0 represents no risk at all. Periodically, Investigate updates this score based
on additional inputs.
A domain blocked by Umbrella receives a score of 100.

## Example Usage


### Basic Usage

Basic usage of the risk_score data source

```hcl
data "umbrella_risk_score" "example" {
  # Add filter attributes here
  id = "12345"
}
```



## Schema

### Optional



### Optional



### Read-Only

- `id` (String) The unique identifier for this resource



