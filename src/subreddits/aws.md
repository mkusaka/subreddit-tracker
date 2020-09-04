# aws
## [1][Amazon CloudFront announces support for TLSv1.3 for viewer connections](https://www.reddit.com/r/aws/comments/imazwl/amazon_cloudfront_announces_support_for_tlsv13/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/09/cloudfront-tlsv1-3-support/
---

## [2][How does EBS expand volumes?](https://www.reddit.com/r/aws/comments/im8ob4/how_does_ebs_expand_volumes/)
- url: https://kichik.com/2020/09/03/how-does-aws-ebs-expand-volumes/
---

## [3][AWS Step Functions increases the maximum payload to 256kb](https://www.reddit.com/r/aws/comments/im3tec/aws_step_functions_increases_the_maximum_payload/)
- url: https://www.reddit.com/r/aws/comments/im3tec/aws_step_functions_increases_the_maximum_payload/
---
[https://aws.amazon.com/about-aws/whats-new/2020/09/aws-step-functions-increases-payload-size-to-256kb/](https://aws.amazon.com/about-aws/whats-new/2020/09/aws-step-functions-increases-payload-size-to-256kb/)
## [4][IAM offline](https://www.reddit.com/r/aws/comments/im2hhu/iam_offline/)
- url: https://www.reddit.com/r/aws/comments/im2hhu/iam_offline/
---
Can't pull up anything in the IAM console or cli right now.

Hopefully it's minor? I wonder if this could affect policy on active resources?

Edit: was down about 15 minutes, 17:10-17:25 EDT or thereabouts.
## [5][Root MFA physical device Storage](https://www.reddit.com/r/aws/comments/imeaqt/root_mfa_physical_device_storage/)
- url: https://www.reddit.com/r/aws/comments/imeaqt/root_mfa_physical_device_storage/
---
AWS best practice is to avoid using the root user to access AWS services and to secure that user with a hardware MFA token.

The company I work at has many AWS accounts and has historically stored the MFA devices in a locked safe in the office. Due to COVID-19 the office is closed and the devices have been sitting with a named individual who can be contacted to get the code. However this is not sustainable when anyone might need root access due to an emergency at any time of day. So we are currently evaluating alternative solutions.

I was wondering if anyone else has run into similar problems and how you have dealt with them?
## [6][GraphQL subscriptions and API gateway](https://www.reddit.com/r/aws/comments/imfngf/graphql_subscriptions_and_api_gateway/)
- url: https://www.reddit.com/r/aws/comments/imfngf/graphql_subscriptions_and_api_gateway/
---
Hey,

Has anyone here been able to hook up graphql subscriptions with API gateway before?

My team is running a node graphql server in elastic bean stalk. We’re not using appsync or amplify with it. Everything is currently working fine with queries and mutations, but we’re having trouble with subscriptions. It’s an authorization issue.

Our current solution for the issue is emitting messages through redis, and picking up the messages on the client with an open socket.

I read some blog post where someone is using a lambda to handle the subscriptions. I feel like that can’t be the only way to do this.
## [7][Changes for AWS customers in Brazil](https://www.reddit.com/r/aws/comments/imfifi/changes_for_aws_customers_in_brazil/)
- url: https://i.redd.it/jmtles60q4l51.png
---

## [8][Is there a way to query the AWS RDS Oracle DB Identifier from SQL?](https://www.reddit.com/r/aws/comments/imfdph/is_there_a_way_to_query_the_aws_rds_oracle_db/)
- url: https://www.reddit.com/r/aws/comments/imfdph/is_there_a_way_to_query_the_aws_rds_oracle_db/
---
When you create a DB in RDS you give it a name (specifically "DB Identifier"). Is there a way to query in SQL from a given instance and get that DB Identifier?

It is not select sys\_context('USERENV', 'DB\_NAME') from dual;

DB\_NAME returns the SID, which is NOT what I want. Unfortunately I have tried all of the SYSCONTEXTs and AWS doesn't set them to anything useful to me. All of the IP related calls return the non-routable addresses on the 192 or 10 network. I want to get to the DNS CNAME or the AWS DB Identifier.
## [9][Is it really worth it to go through a partner for your AWS account to save costs?](https://www.reddit.com/r/aws/comments/im93zn/is_it_really_worth_it_to_go_through_a_partner_for/)
- url: https://www.reddit.com/r/aws/comments/im93zn/is_it_really_worth_it_to_go_through_a_partner_for/
---
What are the pros and cons?  
Is it safe?  
What if you just get a partner for the discounts?   
Since partners cam give discounts right?  
They can lower your AWS billing?
## [10][Why use Amazon Quicksight with other AWS tooling over BusinessObjects/Web Intelligence for reporting?](https://www.reddit.com/r/aws/comments/imf6gw/why_use_amazon_quicksight_with_other_aws_tooling/)
- url: https://www.reddit.com/r/aws/comments/imf6gw/why_use_amazon_quicksight_with_other_aws_tooling/
---
I have my own opinion, but I'd love to hear some answers from the community. We are currently piloting Quicksight, BusinessObjects/WEBI (our current environment), and a bunch of other tools. Quicksight is fun to work with in some regards, but why use it over other tools? 

I understand the pricing model can make the cost of these reports cheaper, but are there other glaring advantages?
