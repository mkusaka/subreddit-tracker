# aws
## [1][You’ve heard it before, outrageous bill on accident](https://www.reddit.com/r/aws/comments/eq08s9/youve_heard_it_before_outrageous_bill_on_accident/)
- url: https://www.reddit.com/r/aws/comments/eq08s9/youve_heard_it_before_outrageous_bill_on_accident/
---
I created a s3 event to compliment my lambda function with a object created event. What I didn’t think of was, my lambda creates s3 objects. So after testing my cool new event and seeing it worked I unknowingly tucked into bed and fell into a deep slumber. Upon waking in the morning I got a AWS free tier limit email. This got me thinking, logged into AWS and see the event is working too good. It’s been running my lambda non stop since last night. I finally got the event to stop after disabling didn’t work I deleted. Got it. As I head over to the billing console I notice a $900 bill. Oh shhhhh. 

Can this be undone?
## [2][Suggestion: AWS console browser notification](https://www.reddit.com/r/aws/comments/epu5i9/suggestion_aws_console_browser_notification/)
- url: https://www.reddit.com/r/aws/comments/epu5i9/suggestion_aws_console_browser_notification/
---
Not sure if this is the best place to post this

I think it would be nice if AWS console sends a browser notification when something completes, eg, when an ec2 instance is up and running, or when an ami image is created

Like gmail notifications, it is up to the user to decide if they want it or not

What do you think?
## [3][How do set-up HTTP caching for client applications in aws Free tier?](https://www.reddit.com/r/aws/comments/eq0b8g/how_do_setup_http_caching_for_client_applications/)
- url: https://www.reddit.com/r/aws/comments/eq0b8g/how_do_setup_http_caching_for_client_applications/
---
My web application is requesting server for all the pages. This in inefficient for lot of reasons.   
I've set up my account in Free tier and currently no way I can move to paid version. 

I want to have HTTP caching enabled for all the requests. [Link to what I'm ranting about.](https://developers.google.com/web/fundamentals/performance/optimizing-content-efficiency/http-caching)

My current set-up :  Client --&gt; API Gateway --&gt; Lambdas --&gt; DynamoDB. 

API Gateway has API caching setting which is for paid users only. 

So, what can I do to enable the most important feature of applications in this case?
## [4][How to get the IAM ARN of an assumed role with Terraform?](https://www.reddit.com/r/aws/comments/epz117/how_to_get_the_iam_arn_of_an_assumed_role_with/)
- url: https://www.reddit.com/r/aws/comments/epz117/how_to_get_the_iam_arn_of_an_assumed_role_with/
---
Hi,

I'm trying to create IAM policies that limit access to just the current role creating them. This is a role assumed from another account. Using the data source `aws_caller_identity.current.arn` I get the ARN from STS, not the IAM one. Something like `arn:aws:sts::1111111111:assumed-role/my-role/111111111836523000` . Does anyone know how to properly create a policy with Terraform where either the principal is the role assumed by the current user or there is a condition that limits access to only this role?

As an example, this is with I currently have that doesn't work.

    data "aws_iam_policy_document" "this" {
      statement {
        effect = "Allow"
        principals {
          type = "AWS"
          identifiers = [
            data.aws_caller_identity.current.arn
          ]
        }
        resources = "*"
        actions = ["secretsmanager:*"]
      }
    }

Every time I run `terraform apply` with that I get the following change:

    Principal = { 
       ~ AWS = "arn:aws:sts::1111111111:assumed-role/my-role/111111111836523000" -&gt; "arn:aws:sts::11111111111:assumed-role/my-role/111111111119111810000"
    }

&amp;#x200B;

Thanks.
## [5][EBS Snapshots: How incremental are they?](https://www.reddit.com/r/aws/comments/eps5hg/ebs_snapshots_how_incremental_are_they/)
- url: https://www.reddit.com/r/aws/comments/eps5hg/ebs_snapshots_how_incremental_are_they/
---
I'm trying to determine if a process of restoring volumes from snapshots and snapshotting the volume is going to save me money or cost me more.

Currently, we have an EBS volume that is inactive most of the time. Once a day, it backs up other data, and then gets snapshotted. In this case, it's fairly obvious to me that those snapshots are incremental as they all are created from the same volume.

Because the volume is inactive most of the day, I was considering deleting the volume after the snapshot was done, and restoring a new volume from that snapshot. I'm wondering if snapshots are incremental in that case, as essentially it'd be a chain of snapshots and volumes: Volume A gets snapshotted to A' and then vol A deleted, A' gets restored to volume B, volume B gets snapshotted to B' and vol B gets deleted, B' gets restored to volume C, etc.

Is snapshot B' incremental off snapshot A', or is it a full volume snapshot because it's the first snapshot off volume B?
## [6][What are you using Amazon Polly for?](https://www.reddit.com/r/aws/comments/epy2gp/what_are_you_using_amazon_polly_for/)
- url: https://www.reddit.com/r/aws/comments/epy2gp/what_are_you_using_amazon_polly_for/
---
I have a product that uses Polly to mainly create audio versions of articles but we are not seeing much engagement from users and therefore we are looking at other use cases that we can pivot to.

Are you or your industry using Amazon Polly? If so, I'd love to learn about the use case.
## [7][Aws network security like ports, cidr scanner](https://www.reddit.com/r/aws/comments/epvebl/aws_network_security_like_ports_cidr_scanner/)
- url: https://www.reddit.com/r/aws/comments/epvebl/aws_network_security_like_ports_cidr_scanner/
---
Hi folks,
Wondering if there is any scanning tool available for aws for network/vpc centric view. For example want to know of an ec2 in a subnet/vpc can talk to another one in different subnet/same vpc, which ports are open via which sgs etc.

Thanks
## [8][AWS Step Functions to trigger a windows CLI](https://www.reddit.com/r/aws/comments/eppbig/aws_step_functions_to_trigger_a_windows_cli/)
- url: https://www.reddit.com/r/aws/comments/eppbig/aws_step_functions_to_trigger_a_windows_cli/
---
I am trying to trigger a windows CLI is it possible in AWS Lambda / Step Functions

&amp;#x200B;

the CLI is from Oculus and I could not locate their API

[https://developer.oculus.com/distribute/latest/concepts/publish-reference-platform-command-line-utility/](https://developer.oculus.com/distribute/latest/concepts/publish-reference-platform-command-line-utility/)
## [9][Django + worker: Beanstalk -&gt; ECS?](https://www.reddit.com/r/aws/comments/epul9r/django_worker_beanstalk_ecs/)
- url: https://www.reddit.com/r/aws/comments/epul9r/django_worker_beanstalk_ecs/
---
Currently using Beanstalk to run a Django app with some async tasks which rely on the django codebase running on the worker tier, managed by the [eb\_sqs library](https://github.com/cuda-networks/django-eb-sqs). I have dockerized my Django app and have it running in ECS backed by Fargate. However, I'm a bit at a loss as to how to go forward with my task running. The options I have considered are:

&amp;#x200B;

1. Use ECS to run celery workers. However, I have read that it takes minutes to spin up new tasks with Fargate, so I'm worried about scaling. Maybe this is not a problem with good autoscaling rules.
2. Use [zappa](https://github.com/Miserlou/Zappa) to run tasks in lambda. No tasks take over the time limit and it would be nice to not worry about scaling workers. However, I'd have to manage 2 deploys and is zappa prod ready?
3. Keep worker tier beanstalk running with django still using that library. Would have 2 deploys, though.

Is there an obvious choice here? Or a choice I'm missing? Thanks in advance :)
## [10][What is the best approach to implement HTML to PDF lambda function?](https://www.reddit.com/r/aws/comments/epu7qx/what_is_the_best_approach_to_implement_html_to/)
- url: https://www.reddit.com/r/aws/comments/epu7qx/what_is_the_best_approach_to_implement_html_to/
---
Currently using wkhtmltopdf and nodejs8.10. Apparently, AWS will no longer support nodejs8.10 and I am forced to use nodejs10.x but wkhtmltopdf doesn't support newer versions. 

&amp;#x200B;

I am now considering changing the language just for this specific function.
