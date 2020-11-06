# aws
## [1][re:Invent registration is now open](https://www.reddit.com/r/aws/comments/jkenu3/reinvent_registration_is_now_open/)
- url: https://register.virtual.awsevents.com/
---

## [2][In the Works – AWS Region in Hyderabad, India](https://www.reddit.com/r/aws/comments/joz9nc/in_the_works_aws_region_in_hyderabad_india/)
- url: https://aws.amazon.com/blogs/aws/in-the-works-aws-region-in-hyderabad-india/
---

## [3][A password encryption application built with Nitro Enclaves](https://www.reddit.com/r/aws/comments/joyvpx/a_password_encryption_application_built_with/)
- url: https://www.sentiatechblog.com/ultra-secure-password-storage-with-nitropepper?utm_source=reddit&amp;utm_medium=social&amp;utm_campaign=nitropepper
---

## [4][New Route53 Interface?](https://www.reddit.com/r/aws/comments/jokgqn/new_route53_interface/)
- url: https://www.reddit.com/r/aws/comments/jokgqn/new_route53_interface/
---
What do you guys think? I think it's incredibly terrible. I've left them feedback and I'm actually seriously concerned they might make it permanent.
## [5][Does anybody feels Lambda frustrating?](https://www.reddit.com/r/aws/comments/jovis9/does_anybody_feels_lambda_frustrating/)
- url: https://www.reddit.com/r/aws/comments/jovis9/does_anybody_feels_lambda_frustrating/
---
Feels like an old drama queen friend that you like once you get to know the quirks. 

Testing is very heavy and takes a long time to deploy, even with SAM. The examples are unnecessary complex. The python packages should not live in the same folder as the code. A lot of nit pick details to make simple code to work. Random errors that gets fixed redeploying the same exact code.
## [6][Trying to get the OnDemand price for the reservations available in my AWS account using pricing Api](https://www.reddit.com/r/aws/comments/jp4ms1/trying_to_get_the_ondemand_price_for_the/)
- url: https://www.reddit.com/r/aws/comments/jp4ms1/trying_to_get_the_ondemand_price_for_the/
---
As reservations are  part of ec2 instances, I was getting cost for ec2 instances but not specifically for the reservation. Is there any possibility to get only cost for reservation or any attribute that available to get details of reservation

    client = boto3.client('pricing', region_name='ap-south-1',                                                              aws_access_key_id='abc',                                                              aws_secret_access_key='')
    
            response = client.get_products(
                ServiceCode='AmazonEC2',
                Filters=[
                    {'Type': 'TERM_MATCH', 'Field': 'operatingSystem', 'Value':'Linux'},
                    {'Type': 'TERM_MATCH', 'Field': 'location','Value': 'US East (N. Virginia)'},
                    {'Type': 'TERM_MATCH', 'Field': 'instanceType', 'Value': 't3.micro'},
                    {'Type': 'TERM_MATCH', 'Field': 'tenancy', 'Value': 'Default'},
                ]
    
            )

This is the code snippet
## [7][EventBridge event storage and replay](https://www.reddit.com/r/aws/comments/jp417a/eventbridge_event_storage_and_replay/)
- url: https://www.reddit.com/r/aws/comments/jp417a/eventbridge_event_storage_and_replay/
---
Looks like EventBridge is getting event archive and replay. Seems like big news but I think it just leaked in the API update, haven’t seen a post yet!

https://awsapichanges.info/archive/changes/034487-events.html
## [8][Anyone who can help me with policies?](https://www.reddit.com/r/aws/comments/jp12cm/anyone_who_can_help_me_with_policies/)
- url: https://www.reddit.com/r/aws/comments/jp12cm/anyone_who_can_help_me_with_policies/
---
Sorry, but I'm a big noob on AWS Policies, I've only used the defaults until now.

What I'm trying to do is the following: -Create a custom policy that allows only for certain regions (currently only 2 regions only) -In the same policy, allow full access to s3, ec2, elastic beanstalk, rds, route 53, cloudfront -In the same policy, limit the instance types that these users can launch (e.g. t3 only or micro and small types only) this is applicable for both EC2 and RDS

Currently I'm having problem with the "Deny" part for EC2 instances, it seems to deny everything currently. Also if I only allow ec2, I can't even create an instance (it throws an error with no clue what's missing) because it requires other permissions such as security group and etc.

Here's what I currently have:

    {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Sid": "AllowGlobals",
                "Effect": "Allow",
                "Action": [
                    "route53:*"
                ],
                "Resource": "*"
            },
            {
                "Sid": "AllowedRegions",
                "Effect": "Allow",
                "Action": [
                    "*",
                    "cloudwatch:*",
                    "ec2:*"
                ],
                "Resource": "*",
                "Condition": {
                    "StringEquals": {
                        "aws:RequestedRegion": [
                            "ap-southeast-1",
                            "ap-southeast-2"
                        ]
                    }
                }
            },
            {
                "Sid": "LimitEc2Instances",
                "Effect": "Deny",
                "Action": [
                    "ec2:RunInstances",
                    "ec2:StartInstances",
                    "ec2:RebootInstances",
                    "ec2:StopInstances"
                ],
                "Resource": [
                    "arn:aws:ec2:*:*:*"
                ],
                "Condition": {
                    "StringNotEquals": {
                        "ec2.InstanceType": [
                            "*.nano",
                            "*.micro",
                            "*.small"
                        ]
                    }
                }
            }
        ]
    }
## [9][API Gateway - Gateway Responses](https://www.reddit.com/r/aws/comments/jp3k6v/api_gateway_gateway_responses/)
- url: https://www.reddit.com/r/aws/comments/jp3k6v/api_gateway_gateway_responses/
---
Hi - I set up Gateway Responses for Response status codes of:

&amp;#x200B;

403 and 404 in the Gateway Responses tab of API Gateway.

&amp;#x200B;

I setup the Response templates to text/html and put in an html body to make the errors more human friendly and helpful.

&amp;#x200B;

I then deployed the new updates and for some reason nothing changes on my end. I still receive:

&amp;#x200B;

{"message":"Missing Authentication Token"}

&amp;#x200B;

Then I look in the google developer console and see the following:

&amp;#x200B;

Failed to load resource: the server responded with a status of 403 ()

&amp;#x200B;

&amp;#x200B;

What does Gateway Responses do? It seems like I made the updates and these templates are not showing.
## [10][How can I inject environment variables into a ECR container (Python app)](https://www.reddit.com/r/aws/comments/jp3702/how_can_i_inject_environment_variables_into_a_ecr/)
- url: https://www.reddit.com/r/aws/comments/jp3702/how_can_i_inject_environment_variables_into_a_ecr/
---
Hey guys,  
Trying to deploy a Python app to AWS using Docker.  
I did push my Docker image to ECR and hooked it to Cluster, Task Definition, Services, etc.

When creating the container in ECR I remember defining my environment variables (it's basically my Stripe Secret Keys + some MySQL credentials) but for some reasons, the app is not picking those up.

I read about a couple of approaches, such as uploading a .env file into a S3 bucket and using that.

Has anyone been successful setting that up?

In the Python code, the variables are defined as follow :

`STRIPE_SECRET_KEY = os.getenv('STRIPE_SECRET_KEY')`
## [11][Email alert when EC2 are left running over 24hrs](https://www.reddit.com/r/aws/comments/jozxui/email_alert_when_ec2_are_left_running_over_24hrs/)
- url: https://www.reddit.com/r/aws/comments/jozxui/email_alert_when_ec2_are_left_running_over_24hrs/
---
Every EC2 instance in my account are for lab purpose that should never be left running &gt; 24hrs.

Objectives:

* Check the launch time on every single EC2 instance (don't care about whether its tagged or what the tag key/value are).
* Every region should be checked.
* Send email to a distribution group listing the Name tag or instance id if there's no tag if EC2 has been running for &gt; 24hrs. Repeat when &gt; 36hrs then again &gt; 48hrs etc.

[This post] (https://www.reddit.com/r/aws/comments/41l2s4/alerts_when_lab_instances_are_left_running_for/) came close to what I was looking for.

Is Lamdba + CloudWatch + SNS Topic the most efficient method to implement this?
Would I need to implement scripts for every region I need checked?
