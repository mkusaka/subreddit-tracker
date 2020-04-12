# aws
## [1][Under the hood: AWS Fargate data plane | Amazon Web Services](https://www.reddit.com/r/aws/comments/fzupjm/under_the_hood_aws_fargate_data_plane_amazon_web/)
- url: https://aws.amazon.com/blogs/containers/under-the-hood-fargate-data-plane/
---

## [2][Should I always verify JSON Web Tokens that get returned from Cognito?](https://www.reddit.com/r/aws/comments/fzo9b6/should_i_always_verify_json_web_tokens_that_get/)
- url: https://www.reddit.com/r/aws/comments/fzo9b6/should_i_always_verify_json_web_tokens_that_get/
---
I've been working with Cognito recently, and I'm a little confused on the importance and use cases of verifying JSON web tokens. AWS has some docs on verifying JWTs here: [https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-verifying-a-jwt.html](https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-verifying-a-jwt.html)

I'm wanting to use Cognito in a very basic way (for the time being):

\- user accesses login page and attempts to sign in

\- credentials are sent to Cognito

\- if authentication is successful, frontend can redirect to a dashboard page (or otherwise update the app state to represent successful sign-in)

Seems like I can use Cognito without ever exposing JWTs to the user or to the front-end (maybe I'm wrong), and the front code can just listen for success or failure and act accordingly.

Should JSON Web Tokens always be verified? Or is it alright to not do that if the case is this simple? Is verifying JWTs mainly if the user is attempting to access other services like data in a DynamoDB table?
## [3][How to deploy Flask API with large files?](https://www.reddit.com/r/aws/comments/fzroac/how_to_deploy_flask_api_with_large_files/)
- url: https://www.reddit.com/r/aws/comments/fzroac/how_to_deploy_flask_api_with_large_files/
---
Been trying to use lambda and elastic beanstalk, but they have hard limits to the source code size. My API uses pre-trained ML models, which is around 600MB total. How do I deal with this? Is it possible to store the models in a S3 bucket and access them from Python? And how slow would it be?
## [4][S3 + Cloudfront + GoDaddy domain: root domain doesn't route to www subdomain](https://www.reddit.com/r/aws/comments/fzkvdo/s3_cloudfront_godaddy_domain_root_domain_doesnt/)
- url: https://www.reddit.com/r/aws/comments/fzkvdo/s3_cloudfront_godaddy_domain_root_domain_doesnt/
---
I want [https://example.](https://example.com)com  to route to [https://www.example.](https://www.example.com)com

But I get the error:

(Firefox) : [example.com](https://example.com) has a security policy called HTTP Strict Transport Security (HSTS), which means that Firefox can only connect to it securely.

Chrome: NET::ERR\_CERT\_COMMON\_NAME\_INVALID

&amp;#x200B;

I have:

An S3 bucket configured to host a static site. A Cloudfront distribution that references the bucket and lists the root and www sub-domain as CNAMES. ACM SSL cert with both root and subdomains listed. A GoDaddy domain with CNAME reference to the distribution and forwarding to [http://www.example.com](http://www.example.com).

What am I missing here?

&amp;#x200B;

&amp;#x200B;

[GoDaddy DNS records](https://preview.redd.it/n4sz934dp9s41.png?width=1924&amp;format=png&amp;auto=webp&amp;s=d8e651796a48f2acc32c0653d3c3a0d273542fa8)

&amp;#x200B;

[GoDaddy DNS forwarding](https://preview.redd.it/kb8kwtvhp9s41.png?width=906&amp;format=png&amp;auto=webp&amp;s=020fa72c93142251ea47eace428b45f23ee65435)

&amp;#x200B;

[Cloudfront Distribution](https://preview.redd.it/cwgi88ojp9s41.png?width=1774&amp;format=png&amp;auto=webp&amp;s=bbedcb413c880a292afc730ae8b1f8fe1205793b)
## [5][CI/CD just using AWS CLI](https://www.reddit.com/r/aws/comments/fzri42/cicd_just_using_aws_cli/)
- url: https://www.reddit.com/r/aws/comments/fzri42/cicd_just_using_aws_cli/
---
In the last few months I've been researching CI/CD in AWS while also looking at Azure DevOps.

A handful of my projects are now using the standard CodeCommit, CodeBuild and CodeDeploy toolchain in AWS. But I am left wondering about dumping and retrieving artefacts from S3 and the length of the CloudFormation templates!

I am sure that I am not getting it so please give me some inputs on why this non-standard way of re-creating and deploying a Spring Boot web service backend is bad?

1. On commit to the stage or master branches in CodeCommit trigger a CodeBuild run.
2. CodeBuild builds a Docker image and pushes to ECR using the AWS CLI
3. The CodeBuild YAML further uses the AWS CLI to check if there is an ECS cluster running the image
   1. If no, create the ECS cluster and run the latest image as a container.
   2. If yes, update it ECS cluster to run the latest image as a container.

I guess bullet 3.1 should be broken up into a CloudFormation template and 3.2 should be broken up into a CodeDeploy step but is the solution presented here really so bad? What principles am I violating? 

Thanks!
## [6][Route53 Cloudfront S3 routing issues with apex domain](https://www.reddit.com/r/aws/comments/fzo2gg/route53_cloudfront_s3_routing_issues_with_apex/)
- url: https://www.reddit.com/r/aws/comments/fzo2gg/route53_cloudfront_s3_routing_issues_with_apex/
---
I have a static site in S3 that I am routing to Cloudfront. 

I am able to reach my website https://www.spinanddestroy.com with the www subdomain 
But I have a problem with my apex domain (https://spinanddestroy.com) 

I have an S3 bucket for www.spinanddestroy.com and I have another bucket for spinanddestroy.com that I redirect to the www subdomain bucket. 

In route53, I have spinanddestroy.com have an A record to www.spinanddestroy.com 
and then I have an A record from www.spinanddestroy.com to the cloudfront distribution (which is sourced from the www.spinanddestroy.com site) 

I need help with why when I navigate to https://spinanddestroy.com I get a 403 error. I want to redirect it back to the www subdomain.  

I also have a hard time trying to troubleshoot this problem. Trace and other network tools are not helping so much. I'd like some advice there as well.
## [7][Optimized way to get data from Kinesis data stream](https://www.reddit.com/r/aws/comments/fzs4ae/optimized_way_to_get_data_from_kinesis_data_stream/)
- url: https://www.reddit.com/r/aws/comments/fzs4ae/optimized_way_to_get_data_from_kinesis_data_stream/
---
What is the best way to ingest the stream from Kinesis? It can be ingested through Spark streaming in AWS EMR and also through AWS Firehose but which way is better or efficient?

The idea is to make it suitable for real time and batch processing both.
## [8][The new AWS calculator is terrible.](https://www.reddit.com/r/aws/comments/fyvifi/the_new_aws_calculator_is_terrible/)
- url: https://www.reddit.com/r/aws/comments/fyvifi/the_new_aws_calculator_is_terrible/
---
Am I the only one that hates the new AWS calculator, [calculator.aws](https://calculator.aws)? I don't need a wizard to hold my hand through getting pricing. I need something simple, like the already existing simply monthly calculator. Now I see that the simple monthly calculator is going away on June 30, 2020, and we have to start using calculator.aws? Say it ain't so.
## [9][Stopping all the ecs tasks excluding some](https://www.reddit.com/r/aws/comments/fzm9u3/stopping_all_the_ecs_tasks_excluding_some/)
- url: https://www.reddit.com/r/aws/comments/fzm9u3/stopping_all_the_ecs_tasks_excluding_some/
---
Hi
 I want to shut down all the tasks running in the ec2 backed ecs cluster excluding some. I also want keep a list of the ones I shut down so i can bring back those when needed.  

There are already some stopped tasks in the cluster I won’t touch those.

I am planning to use boto3, is that a Good way to do it ? 

Any other suggestions?
## [10][EC2 t2.2xlarge instance. CloudWatch CPU Utilization never goes above 12%, but Linux htop reports over 50% usage](https://www.reddit.com/r/aws/comments/fzdjtl/ec2_t22xlarge_instance_cloudwatch_cpu_utilization/)
- url: https://www.reddit.com/r/aws/comments/fzdjtl/ec2_t22xlarge_instance_cloudwatch_cpu_utilization/
---
My webserver is a  t2.2xlarge instance (8 core). 

T2Unlimited is switched on.

The webserver receives lots of web requests every second, so I do expected high cpu usage like linux htop reports. But I don't understand why the cloudwatch dashboard only reports it as 5-12% permanently, never above 12%, and when it hits 12% it seems to be peaking.

Am I doing something completely stupid / are these dashboards just incorrect?

&amp;#x200B;

Thanks!
