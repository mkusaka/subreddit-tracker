# aws
## [1][What's the best way to design a DynamoDB table to store cloudwatch alarms](https://www.reddit.com/r/aws/comments/f56ufi/whats_the_best_way_to_design_a_dynamodb_table_to/)
- url: https://www.reddit.com/r/aws/comments/f56ufi/whats_the_best_way_to_design_a_dynamodb_table_to/
---
 Actually i'm using Cloudwatch to generate alarms for my different services (Elasticsearch, RDS, ELB, ...) and my goal is to store these alarms in a DynamoDB table.

These are the fieds that i will store :

    - Service (Elasticsearch, RDS, ELB, ...)  
    - Priority of the alert(P1,P2,P3,P4)  
    - Datetime  
    - Resource Name  
    - Status  
    - Reason 

Is it ok to design the table like that :

    id = Service#Priority 
    sort_key = Datetime 

Or simply :

    id = Service 
    sort_key = Datetime 

I'm not sure what's the best approach to use to better design the table.

NOTE : About 50 alerts will be stored each day.

Thanks.
## [2][CloudFormation Linter](https://www.reddit.com/r/aws/comments/f4w7qy/cloudformation_linter/)
- url: https://github.com/aws-cloudformation/cfn-python-lint/
---

## [3][Any good article to read about continuous optimisation?](https://www.reddit.com/r/aws/comments/f56jfp/any_good_article_to_read_about_continuous/)
- url: https://www.reddit.com/r/aws/comments/f56jfp/any_good_article_to_read_about_continuous/
---
Continuous optimisation in terms of performance and cost.
## [4][AWS SES how to log sender using IAM smtp username/password?](https://www.reddit.com/r/aws/comments/f57hw5/aws_ses_how_to_log_sender_using_iam_smtp/)
- url: https://www.reddit.com/r/aws/comments/f57hw5/aws_ses_how_to_log_sender_using_iam_smtp/
---
My SES IAM smtp username is leaked.

Is it possibile to log sender source?
## [5][How to invoke a single Lambda function for multiple S3 file uploads?](https://www.reddit.com/r/aws/comments/f56wyk/how_to_invoke_a_single_lambda_function_for/)
- url: https://www.reddit.com/r/aws/comments/f56wyk/how_to_invoke_a_single_lambda_function_for/
---
I have a bucket in S3 to which two .csv files or more come in each day. Right now, I have set up an event trigger which invokes a Lambda function for **each** .csv file and uploads the names and the eTags of these files to a DynamoDB table. There are two things I want to know:

* How to make it so that only one Lambda function is triggered for multiple uploaded files?
* I'm using eTag as a unique file identifier as DynamoDB doesn't have auto-increment functionality. Is there any better method to uniquely identify each file?
## [6][Best way to add security to api-gateway which will be hit by an API on azure.](https://www.reddit.com/r/aws/comments/f58m3j/best_way_to_add_security_to_apigateway_which_will/)
- url: https://www.reddit.com/r/aws/comments/f58m3j/best_way_to_add_security_to_apigateway_which_will/
---
Hey guys,

Basically we have an API on Azure which we want to hit an api-gateway on AWS.

We were thinking of simply having an API key stored on the Azure API which we can use to authenticate through a lambda authoriser, since the API key exists only on the Azure API it should be relatively safe.

Is there a better way to make sure that calls that hit our gateway are only coming from our Azure API?
## [7][Single-Table Design in DynamoDB](https://www.reddit.com/r/aws/comments/f4q5ha/singletable_design_in_dynamodb/)
- url: http://alexdebrie.com/posts/dynamodb-single-table/
---

## [8][Can we set default schemas for AWS Redshift roles to virtually achieve different databases under ONE AWS Redshift instance?](https://www.reddit.com/r/aws/comments/f58hog/can_we_set_default_schemas_for_aws_redshift_roles/)
- url: https://www.reddit.com/r/aws/comments/f58hog/can_we_set_default_schemas_for_aws_redshift_roles/
---
We have a huge `AWS RDS POSTGRES` database we are looking to migrate to `AWS Redshift`.
Right now, under our RDS instance, we have created two databases;
1. For QA.
2. For Production.

We are looking to switch to `AWS Redshift` now.
There's a problem that under ONE instance of AWS Redshift, we can only create ONE database.
Now we can not afford to create two `AWS REDSHIFT` instances with 4 cores as they are expensive.

I came to know from [here](https://stackoverflow.com/a/3282682/10190191) that we can create multiple schemas in the database and then we can set default schema for a certain role.
This way, we can (without changing our queries by putting schema names in front of tables in `FROM` clause and hence without changing our queries) create two (virtually) separate databases for `QA` and `PROD` by creating two roles having different default schemas.

My question is, is it possible to achieve this in `AWS Redshift`? 
As it is different at places from normal relational databases implementation wise at many places.



&gt; Instance is costly so we can not GET an instance and then try this. We have to be sure before buying it.
## [9][Amazon RDS SSL/TLS update does not work for global database clusters](https://www.reddit.com/r/aws/comments/f53fxm/amazon_rds_ssltls_update_does_not_work_for_global/)
- url: https://www.reddit.com/r/aws/comments/f53fxm/amazon_rds_ssltls_update_does_not_work_for_global/
---
FYI, from the UI it does not work

"DBInstance not found: xxx-global-db-instance-1"

state is still "Requires Update"

not really an issue for us, but maybe for others
## [10][Cognito + AWS Amplify](https://www.reddit.com/r/aws/comments/f54jeu/cognito_aws_amplify/)
- url: https://www.reddit.com/r/aws/comments/f54jeu/cognito_aws_amplify/
---
Pretty committed to Cognito and AWS Amplify, but I keep running into issues not explained in the documentation. I've been running into an odd (accidental) case where aws-amplify doesn't know what to do when two different users try logging in at the same time. 

Having another problem where, when a user logs out, my server can't just verify that the JWT is valid - but has to run `getUser({accessToken}`... which can be stupid slow when it comes to response times. As you only get 5 rps. I totally get that with stateless I need to check a blacklist of JWT's - but why would they limit this to only 5 rps? 

It just seems like a half-baked service. I wish I wasn't baited in with the 50k free users. 

Has anyone else dealt with this? Any advice on migrating (3k+ users)? Or venting any similar experiences would be appreciated!
