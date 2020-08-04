# aws
## [1][Week of Aug 3rd - What are you building this week on AWS?](https://www.reddit.com/r/aws/comments/i2ydcv/week_of_aug_3rd_what_are_you_building_this_week/)
- url: https://www.reddit.com/r/aws/comments/i2ydcv/week_of_aug_3rd_what_are_you_building_this_week/
---
Discuss your challenges and successes
## [2][Why You Should Never, Ever print() in a Lambda Function](https://www.reddit.com/r/aws/comments/i3h6rn/why_you_should_never_ever_print_in_a_lambda/)
- url: https://medium.com/@psingman/why-you-should-never-ever-print-in-a-lambda-function-f997d684a705
---

## [3][Free Workshop to build enterprise search service using Amazon Kendra](https://www.reddit.com/r/aws/comments/i3fo4e/free_workshop_to_build_enterprise_search_service/)
- url: http://aws-dojo.com/workshoplists/workshoplist5
---

## [4][Aws AppFlow: Unlock your SaaS data](https://www.reddit.com/r/aws/comments/i3itox/aws_appflow_unlock_your_saas_data/)
- url: https://medium.com/@yvescallaert/aws-appflow-unlock-your-saas-data-d3308cc99c7f?source=friends_link&amp;sk=00adfe5d606c6c4107885c1441dbafd9
---

## [5][List of resource types on AWS that can be made public](https://www.reddit.com/r/aws/comments/i39opb/list_of_resource_types_on_aws_that_can_be_made/)
- url: https://github.com/SummitRoute/aws_exposable_resources
---

## [6][Awesome AWS workshops](https://www.reddit.com/r/aws/comments/i2wcl2/awesome_aws_workshops/)
- url: https://awesome-aws-workshops.com/
---

## [7][EBS and long running request](https://www.reddit.com/r/aws/comments/i3glkf/ebs_and_long_running_request/)
- url: https://www.reddit.com/r/aws/comments/i3glkf/ebs_and_long_running_request/
---
Hi, I have a dot net app which can sometimes take 2-3 hours to process a request, not all requests, just an occasional large file import. This process sends status updates to the client using SignalR and redis so everyone knows whats happening. Can someone tell me what the best way is to to keep the EC2 instance running while this process is running. The EBS is setup to auto scale with a classic load balancer.
## [8][I made this with an AWS SageMaker Notebook... :)](https://www.reddit.com/r/aws/comments/i3aow6/i_made_this_with_an_aws_sagemaker_notebook/)
- url: https://www.reddit.com/r/aws/comments/i3aow6/i_made_this_with_an_aws_sagemaker_notebook/
---
https://reddit.com/link/i3aow6/video/phezog3q8we51/player

I'm having fun with Machine Learning!

The colors of the nodes show the bias and the width of the lines show the weights of the model while it trains.

It's running on an AWS SageMaker Notebook with a lot of hideously hacky code.  But I like the end result :)
## [9][Security Documentation for 165+ AWS Services](https://www.reddit.com/r/aws/comments/i33t61/security_documentation_for_165_aws_services/)
- url: https://docs.aws.amazon.com/security/
---

## [10][Make shared S3 bucket read-only for most of our team?](https://www.reddit.com/r/aws/comments/i3goel/make_shared_s3_bucket_readonly_for_most_of_our/)
- url: https://www.reddit.com/r/aws/comments/i3goel/make_shared_s3_bucket_readonly_for_most_of_our/
---
Hello all.   


Apologies for the newbie question. My team shares an AWS account. We use federated (SSO) login that is managed by our IT department to access the console. We have then each set up for ourselves one or more IAM user(s) with access keys for CLI or SDK use. 

Currently, we're working with data that we would like to remain read-only to our team (and blocked to the world) to comply with best practices in our industry. The only people with write access should be me and our line manager (CLI-only is fine). I have put this data into a bucket and I've modified the bucket policy to be as follows:

    {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Effect": "Deny",
                "Sid": "MakeReadOnly",
                "Action": [
                    "s3:GetBucketPolicy",
                    "s3:PutBucketPolicy",
                    "s3:putObject"
                ],
                "NotPrincipal": {
                    "AWS": [
                        "arn:aws:iam::123456789012:user/line_manager_iam_user",
                        "arn:aws:iam::123456789012:user/my_iam_user"
                    ]
                },
                "Resource": [
                    "arn:aws:s3:::read-only-bucket",
                    "arn:aws:s3:::read-only-bucket/*"
                ]
            }
        ]
    }

This works fine, but since everyone in the team has access to the IAM console, it's easy enough for any of them to just generate a new set of access keys for my IAM user, and then use those keys to modify or delete the data.

What is the best way to do this? Do I have to request that the IT team limit our ability to generate new access keys for IAM users? That would be less than ideal because we make and delete access keys often.
## [11][What the difference between Lambda and EC2 in VPC?](https://www.reddit.com/r/aws/comments/i3jicc/what_the_difference_between_lambda_and_ec2_in_vpc/)
- url: https://www.reddit.com/r/aws/comments/i3jicc/what_the_difference_between_lambda_and_ec2_in_vpc/
---
I have VPC with one subnet that connected to Internet Gateway. That subnet contains few Lambda functions and ElastiCache server. I see that these Lambda functions does not have access to Internet because there is no NAT. But if I put EC2 in the same subnet then this server does have access to Internet. Why? What the difference between Lambda and EC2 in this VPC?

I ask because I don't want to pay for NAT (it's expensive) but I would like to give internet access to my Lambda functions and have access to ElastiCache on the same time. Looks like it's impossible to implement.

If I use ECS and put all services and functions into containers it will be cheaper?
