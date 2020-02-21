# aws
## [1][Savings Plan Update: Save Up to 17% On Your Lambda Workloads](https://www.reddit.com/r/aws/comments/f70qg5/savings_plan_update_save_up_to_17_on_your_lambda/)
- url: https://aws.amazon.com/blogs/aws/savings-plan-update-save-up-to-17-on-your-lambda-workloads/
---

## [2][Anyone using typescript in lambdas? How do you debug crash logs?](https://www.reddit.com/r/aws/comments/f78ks1/anyone_using_typescript_in_lambdas_how_do_you/)
- url: https://www.reddit.com/r/aws/comments/f78ks1/anyone_using_typescript_in_lambdas_how_do_you/
---
I am looking to use typescript in AWS lambda. Our current stack in a production ready app is Javascript lambdas exposed through gateway, and all deployment is done using terraform. Due to high number of errors in types, the team is looking to use typescript in lambdas.

I have couple of questions regarding to TS in lambda environment.

&amp;#x200B;

1. How do you transpile js files and how do you deploy? I know there are number of tutorials that use serverless framework, but we are using terraform and not looking to use another deployment framework.
2. How do you debug any crashes? If deployed files are JS, how do you map cloudwatch logs to your TS code? Currently in our JS code this is not much a question.
3. Is there any other tricky point we have to take care of in TS?
## [3][VPC with peering vs one bigger VPC?](https://www.reddit.com/r/aws/comments/f79vsa/vpc_with_peering_vs_one_bigger_vpc/)
- url: https://www.reddit.com/r/aws/comments/f79vsa/vpc_with_peering_vs_one_bigger_vpc/
---
Hi,

I considered having 2 separate AWS dev and test environments. However, the cost of having two RDS, DocumentDB clusters and MSK clusters became too large. Therefore I'd like to use same databases in TEST and DEV environments. Test is also used internally so there's no risk in breaking it. However I'm not sure which is a better approach to connect test and dev vpcs with peering or simple move applications into VPC?
## [4][RDS Certificate bundle 2019](https://www.reddit.com/r/aws/comments/f79d1g/rds_certificate_bundle_2019/)
- url: https://www.reddit.com/r/aws/comments/f79d1g/rds_certificate_bundle_2019/
---
AWS have provided new RDS certificates for 2019, however the link for the bundle is still the 2015 certificate.

The docs provide the links 2015 and 2019 root certificates, however the link to download the bundle (that contains both the intermediate and root certificates) only provides the 2015 certificate

[https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL.html](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL.html)

```
$ curl -L https://s3.amazonaws.com/rds-downloads/rds-combined-ca-bundle.pem -o rds-cert.pem
$ openssl x509 -in rds-cert.pem -noout -dates
notBefore=Feb  5 09:11:31 2015 GMT
notAfter=Mar  5 09:11:31 2020 GMT
```

Is there a way to get the certificate bundle including the 2019 certificate, or do we expect the bundle link to be updated soon?
## [5][New IAM condition key: aws:CalledViaFirst](https://www.reddit.com/r/aws/comments/f7b2nc/new_iam_condition_key_awscalledviafirst/)
- url: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_condition-keys.html#condition-keys-calledviafirst
---

## [6][AWS Savings Plans now include Lambda](https://www.reddit.com/r/aws/comments/f77dpk/aws_savings_plans_now_include_lambda/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/02/aws-lambda-participates-in-compute-savings-plans/
---

## [7][aws ec2 hosted website, not touched in months, images suddenly not loading but they exist with direct image path](https://www.reddit.com/r/aws/comments/f7akn0/aws_ec2_hosted_website_not_touched_in_months/)
- url: https://www.reddit.com/r/aws/comments/f7akn0/aws_ec2_hosted_website_not_touched_in_months/
---
What is going on? Is amazon updating stuff and caused bugs? seems unlikely
## [8][Error with push in code commit](https://www.reddit.com/r/aws/comments/f7a5yx/error_with_push_in_code_commit/)
- url: https://www.reddit.com/r/aws/comments/f7a5yx/error_with_push_in_code_commit/
---
I have cloned my codecommit repo and done changes but when I try to push it I get 
error: RPC failed; HTTP 403 curl 22 The requested URL returned error: 403
fatal: the remote end hung up unexpectedly
Someone on stack overflow said its a post buffer issue but I've increased and still not able to push.
Can someone help?
Thanks in advance.
## [9][.NET Core AWS Lambda Multiple Function in Single Project](https://www.reddit.com/r/aws/comments/f7491b/net_core_aws_lambda_multiple_function_in_single/)
- url: https://irensaltali.com/en/net-core-aws-lambda-multiple-function-in-single-project/
---

## [10][Connecting to a Postgres RDS instance](https://www.reddit.com/r/aws/comments/f71pp4/connecting_to_a_postgres_rds_instance/)
- url: https://www.reddit.com/r/aws/comments/f71pp4/connecting_to_a_postgres_rds_instance/
---
So I have a VPC with an RDS that is not publicly accessible as is typically recommended for obvious reasons. However, I need to connect a third party application (Sigma if you’re curious or this affects the answer). From everything I see the RDS instance needs to be publicly accessible. 

My idea is to create a read replica that is on a different subnet and is publicly accessible with a security group that only allows incoming connections on port 5432 and only to two specific IPs sigma provides. 

Questions:

1. Is this possible?

2. Are there security risks?

3. Is there a better way to do it?
