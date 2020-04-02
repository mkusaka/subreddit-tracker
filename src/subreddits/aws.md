# aws
## [1][Amazon Elastic File System announces 400% increase in read operations for General Purpose mode file systems](https://www.reddit.com/r/aws/comments/ft2wqq/amazon_elastic_file_system_announces_400_increase/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/04/amazon-elastic-file-system-announces-increase-in-read-operations-for-general-purpose-file-systems/
---

## [2][How to ensure that a role can be assumed but a single resource?](https://www.reddit.com/r/aws/comments/ftllyb/how_to_ensure_that_a_role_can_be_assumed_but_a/)
- url: https://www.reddit.com/r/aws/comments/ftllyb/how_to_ensure_that_a_role_can_be_assumed_but_a/
---
I have a Lambda function so I had to create a role with some permissions and also had to add a trusted relationship to the Lambda service. But how to ensure that just my lambda function can assume that role?

I've tried everything but nothing is working.
## [3][How to use a Terraform state stored in one AWS account and deploy it in another?](https://www.reddit.com/r/aws/comments/ftlq05/how_to_use_a_terraform_state_stored_in_one_aws/)
- url: https://www.reddit.com/r/aws/comments/ftlq05/how_to_use_a_terraform_state_stored_in_one_aws/
---
When using Terraform with other people it’s often useful to store your state in a bucket. For example, an S3 bucket if you deploy on AWS. And then you may want to use the same bucket for different AWS accounts for consistency purposes. Or you may also want your S3 bucket to be stored in a different AWS account for right management reasons. This article explains [how to use a Terraform state stored in one AWS account](https://www.padok.fr/en/blog/terraform-s3-bucket-aws) and deploy it in another one.
## [4][Serverless Cognito app - where to store secure data?](https://www.reddit.com/r/aws/comments/ftlaw6/serverless_cognito_app_where_to_store_secure_data/)
- url: https://www.reddit.com/r/aws/comments/ftlaw6/serverless_cognito_app_where_to_store_secure_data/
---
Hello,

I am developing a serverless application and managing user data inside Cognito. However, I need to store bank account details (sort and account number). 

What would be the best method to save this highly secure information? 

I am not sure if Cognito has a method of saving this information (I don't think so). 

Possibly in parameter store as a secret? Or possibly store as an encrypted value with the encryption key being stored in secrets manager? Could even use KMS for this?

There are a lot of possible solutions but I am not sure which would be the most appropriate.
## [5][Turning my already deployed HTTP site to use HTTPS](https://www.reddit.com/r/aws/comments/ftf8np/turning_my_already_deployed_http_site_to_use_https/)
- url: https://www.reddit.com/r/aws/comments/ftf8np/turning_my_already_deployed_http_site_to_use_https/
---
The site works fine but it's served in http. I want https.  
 
- I went to CloudFront → selected the static site bucket’s distribution → distribution settings → `Behaviors` → `Edit` → Selected `HTTPS only`  
- I also went to the bucket, → Properties → Static website hosting → Redirect requests → entered in the target domain and https.  
  
I did that because before, I had the default behavior selected to reroute HTTP to HTTPS. That definitely didn’t work on any browser I used. Either way, the site still loads as an HTTP site. What’s worse, if I manually type in `https://www.mysite.xyz` the page doesn’t load at all.  
  
Some context:  
- Its CloudFront distribution is enabled and deployed.  
- Its SSL Certificate is the custom certificate that I requested. Setting it to this was no different from when the site used the default CloudFront Certificate.  
- Bucket policy:  
    {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Sid": "PublicReadGetObject",
                "Effect": "Allow",
                "Principal": "*",
                "Action": "s3:GetObject",
                "Resource": "arn:aws:s3:::my-site.xyz/*"
            }
        ]
    }

Here's one of the [better tutorials](https://dcurt.is/how-to-host-a-static-website-on-aws-with-https) I found on the matter, since most only detail how to get to the https setting in CloudFront.  
Pardon the shitty markdown. [Reddit's instructions](https://www.reddit.com/wiki/markdown) don't work.
## [6][Assume role vs changing S3 bucket acl](https://www.reddit.com/r/aws/comments/ftguh6/assume_role_vs_changing_s3_bucket_acl/)
- url: https://www.reddit.com/r/aws/comments/ftguh6/assume_role_vs_changing_s3_bucket_acl/
---
If an application on an EC2 instance from Account A needs to upload files to a S3 bucket in Account B owned by someone else, is it better to assume a role in Account B or change the acl like the example below in order to grant bucket owner permission to access the uploaded file in the S3 bucket?

    aws s3 cp --acl bucket-owner-full-control localFile s3://bucketname/path/filename
## [7][Should I switch my ALB to use the Least Outstanding Requests algorithm? YES!](https://www.reddit.com/r/aws/comments/fsy2sq/should_i_switch_my_alb_to_use_the_least/)
- url: https://medium.com/dazn-tech/aws-application-load-balancer-algorithms-765be2eca158?source=friends_link&amp;sk=2da467951a30524ef398b9b333707d43
---

## [8][Scalable Serverless Microservice Demo AWS Lambda Kinesis Terraform](https://www.reddit.com/r/aws/comments/ftik1i/scalable_serverless_microservice_demo_aws_lambda/)
- url: https://medium.com//scalable-serverless-microservice-demo-aws-lambda-kinesis-terraform-cbe6036bf5ac?source=friends_link&amp;sk=074614683a6641cab9b6067929bdc660
---

## [9][Invoke Async Lambda Function from Go](https://www.reddit.com/r/aws/comments/ftbzaa/invoke_async_lambda_function_from_go/)
- url: https://www.reddit.com/r/aws/comments/ftbzaa/invoke_async_lambda_function_from_go/
---
I'm trying to invoke Lambda function from another Lambda function asynchronously with Go.

There seems to be a method to invoke them as listed in the documentation [here](https://docs.aws.amazon.com/sdk-for-go/api/service/lambda/#Lambda.Invoke) but I cannot find any documentation on what the handler should look like in Go.

If anyone has done this before and has an example of what the handler should look like that would be fantastic.
## [10][Is this security group normal for Workspaces directory?](https://www.reddit.com/r/aws/comments/ftdl6f/is_this_security_group_normal_for_workspaces/)
- url: https://www.reddit.com/r/aws/comments/ftdl6f/is_this_security_group_normal_for_workspaces/
---
with all this working from home going on i spun up a workspaces machine in AWS and in part of doing this workspaces requires a directory. so i opted for the AWS Simple AD directory service.

once the simple ad directory service was created i noticed a new security group was also created and upon looking at it, it's open to everyone and their brother...

&amp;#x200B;

[security group](https://preview.redd.it/96zj4aw59bq41.png?width=550&amp;format=png&amp;auto=webp&amp;s=7778a5fad006660052990eef776619fd680ba33a)

is this normal? this certainly can't be best practice is it? how best can i make this more secure?
