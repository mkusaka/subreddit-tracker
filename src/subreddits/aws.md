# aws
## [1][Anyone having acquired Outposts, and willing to share some feedback on usage, so far?](https://www.reddit.com/r/aws/comments/gbh5aj/anyone_having_acquired_outposts_and_willing_to/)
- url: https://www.reddit.com/r/aws/comments/gbh5aj/anyone_having_acquired_outposts_and_willing_to/
---
Among things of interest:

- reason for acquisition &amp; expectations being met?!?

- performance compared to the region, both vis-a-vis on-prem

- LGW - w/CoIP, or else?!?

- services on-prem delivered for workloads in Outposts (e.g. are you load balancing with on-prem devices like Citrix, F5, A10?)

- ... anything else that comes to mind ...
## [2][Can't pass and pull values from Mysql to AWS](https://www.reddit.com/r/aws/comments/gbg61q/cant_pass_and_pull_values_from_mysql_to_aws/)
- url: https://www.reddit.com/r/aws/comments/gbg61q/cant_pass_and_pull_values_from_mysql_to_aws/
---
I have the connection established from Mysql to AWS RDS, or I think I do at least. The connection seems to be working on both sides. However, I can't pass or pull values from my website using PHP, it just comes back as a connection error when I try to pull them or nothing gets passed when I try to push them.
## [3][FORMULA 1 DeepRacer ProAm Special Event](https://www.reddit.com/r/aws/comments/gb6rwq/formula_1_deepracer_proam_special_event/)
- url: https://aws.amazon.com/blogs/aws/join-the-formula-1-deepracer-proam-special-event/
---

## [4][AWS Workmail - Who uses it?](https://www.reddit.com/r/aws/comments/gbapw9/aws_workmail_who_uses_it/)
- url: https://www.reddit.com/r/aws/comments/gbapw9/aws_workmail_who_uses_it/
---
I know they have their sample companies on the website... but I haven’t heard of any of them before. Are there any major name companies that use workmail?
## [5][Accessing multiple APIs from the same base URL](https://www.reddit.com/r/aws/comments/gbfz4l/accessing_multiple_apis_from_the_same_base_url/)
- url: https://www.reddit.com/r/aws/comments/gbfz4l/accessing_multiple_apis_from_the_same_base_url/
---
I would like to split my template with all the lambda functions into several templates for each domain (products, users etc). The lambdas are supposed to be accessed from an API. I cannot create a single API for these, CloudFormation arouses issues on that (ApiEvents should be intruduced in the same template as the API itself). So my approach would be that each template has its own API from which the lambdas are accessible. But each API naturally has its own URL when CF deploys them. 

Obviously I would like to access all the lambdas from the same base URL. How could this be done, should I use AWS::ApiGateway::DomainName and AWS::ApiGateway::BasePathMapping? Have I understood correctly that those could be used in that way? So [test.myproject.com/products](https://test.myproject.com/products) would access functions in products template via its API and [test.myproject.com/users](https://test.myproject.com/users) would access functions in users teamplate via its API.  


[https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html)
## [6][How to handle over 200 lambdas with Cloud Formation?](https://www.reddit.com/r/aws/comments/gb4ht8/how_to_handle_over_200_lambdas_with_cloud/)
- url: https://www.reddit.com/r/aws/comments/gb4ht8/how_to_handle_over_200_lambdas_with_cloud/
---
I have a few stacks, one for the network, another for database and such. And then I have a stack for all the Serverless::Api and the Serverless::Functions. 

I have rached the limit of 200 resources in that stack. I tried to separate some of the functions to a different stack and referencing to the Api with "!ImportValue MyApi" where needed, ie. function events. But when trying to deploy, I get: "Api Event must reference an Api in the same template". So this cannot be done.

I cannot introduce all the api events in one stack with the api since I would hit the 200 limit again. How about nesting stacks? If I have api in one stack and two stacks for functions that depend on the api stack, would that help me or would I get the same error again (events in the same temolate as the api)?

What would be the best approach here?

Edit: The title is wrong, there aren't over 200 lambdas but over 200 resources. I have about 80 lambdas in the template but CF creates AWS::Lamda::Permission for each lambda when deployed. I know that is too much and that is why I'm seeking help to how to resolve this and split it into smaller stacks and not getting the "Api Event must reference an Api in the same template" error.

Edit2: When trying to nest stacks so that the Api is in one stack and some of the lambdas in another, nested stack, I get error: "The REST API doesn't contain any methods". I tried adding one lamda to the same template as the Api is in and nest the other functions in other templates. But then I still get that "Api Event must reference an Api in the same template. So either I have to introduce all the api events in the same template as the api is in (pretty cumbersome) OR have several templates with lambdas and each having its own api, but I would need a way to access all the endpoints via the same base URL.
## [7][AWS Lambda occasional 500 error](https://www.reddit.com/r/aws/comments/gbdn5t/aws_lambda_occasional_500_error/)
- url: https://www.reddit.com/r/aws/comments/gbdn5t/aws_lambda_occasional_500_error/
---
For past few months I have a strange problem. Sometimes (maybe once in 10 days) all of my Lambda functions show 500 or 502 errors. This lasts only for a few minutes, so I can't properly diagnose the problem. When this error occurs I check cloudwatch logs, but it seems that API calls don't even reach the code. During this errors, software works on a local machine and database seems to be fine. Runtime for all of the code is Node.js10.x.

Does anybody have this kind of problem?

What can be the possible solution?
## [8][Best way to programmatically get authorization to hit an AWS endpoint?](https://www.reddit.com/r/aws/comments/gb66gf/best_way_to_programmatically_get_authorization_to/)
- url: https://www.reddit.com/r/aws/comments/gb66gf/best_way_to_programmatically_get_authorization_to/
---
I'm familiar with writing nodejs scripts to test endpoints, usually with a bearer token, but I'm not familiar with hitting an AWS endpoint. What's the best way to generate a token in nodejs?
## [9][Stopping RDS via Lambda](https://www.reddit.com/r/aws/comments/gbeu5r/stopping_rds_via_lambda/)
- url: https://www.reddit.com/r/aws/comments/gbeu5r/stopping_rds_via_lambda/
---
Hi I am learning how to stop my RDS via Lambda and I came across this  [https://dzone.com/articles/create-an-aws-lambda-function-to-stop-and-start-an](https://dzone.com/articles/create-an-aws-lambda-function-to-stop-and-start-an) 

I followed the instruction to configure but when I ran the Lambda function, there is an error "Task timed out after 3.00 second" I am unable to determine why.  Is the Python script correct?
## [10][How to do RDS PostgreSQL point in time recovery ?](https://www.reddit.com/r/aws/comments/gbed54/how_to_do_rds_postgresql_point_in_time_recovery/)
- url: https://www.reddit.com/r/aws/comments/gbed54/how_to_do_rds_postgresql_point_in_time_recovery/
---
You can go and do a PITR in Actions&gt;Restore to point in time , but you can do this up to a certain period, like up until 10 minutes ago. 
Is there a way to make it available to be recovered until 1 minute ago ?
