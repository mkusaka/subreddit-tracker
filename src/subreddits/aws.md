# aws
## [1][Another outage?](https://www.reddit.com/r/aws/comments/iu686c/another_outage/)
- url: https://www.reddit.com/r/aws/comments/iu686c/another_outage/
---
Looks like this one might be IAM related.

&gt;\[03:17 PM PDT\] We are investigating increased authentication error rates and latencies affecting IAM. IAM related requests to other AWS services may also be impacted. 

My eb and aws cli are both affected, can't push anything.
## [2][New AWS Console with favorites menu inside the services button](https://www.reddit.com/r/aws/comments/iufnk5/new_aws_console_with_favorites_menu_inside_the/)
- url: https://www.reddit.com/r/aws/comments/iufnk5/new_aws_console_with_favorites_menu_inside_the/
---
Now i need two clicks to visit a service instead of right clicking my favorite small icon service and opening in a new tab.

I know this console is not at all used if you do IaC but sometimes this needs to be shown to others in a graphical way
## [3][Τ-family instances](https://www.reddit.com/r/aws/comments/iug6j1/τfamily_instances/)
- url: https://www.reddit.com/r/aws/comments/iug6j1/τfamily_instances/
---
After the announcement of the new T4g we are checking the wider T-family. Unpredictable pricing is a little awkward. On the other hand, the bursting concept seems convenient.

Interestingly enough their base price is lower than respective Ms. For example, in Ohio:

T3.xlarge (4 vCPUs, 16GB RAM) is at $0.1664/hr
M5.xlarge (4 vCPUs, 16GB RAM) is at $0.192/hr 

Have you used T-family instances? What do you think about them?
## [4][S3 Replication RTC with CloudFormation?](https://www.reddit.com/r/aws/comments/iuj0ww/s3_replication_rtc_with_cloudformation/)
- url: https://www.reddit.com/r/aws/comments/iuj0ww/s3_replication_rtc_with_cloudformation/
---
Hey Guys!

I've been writing a CF template that will create two S3 buckets and setup SRR (Same Region Replication) between them. Got everything working fine and the buckets replicate no bother. But when i try to add RTC (and get the 15 minutes replication time) to the template it all fails and i can't even deploy it.

I've followed along with the S3 CloudFormation docs and did exactly as it said. The problem seems to be from this document and the one it links to. I've found no examples online or anything and im beginning to think this feature barely exists lol.

Error im getting inside CloudFormation is :

*Encountered unsupported property ReplicationConfiguration*

My code is below that im using for the bucket creation that im adding RTC to (with the bucket names changed), any help would be so appreciated!

      OriginalBucket:
        Type: AWS::S3::Bucket
        Properties:
          BucketName: original-bucket
          VersioningConfiguration:
            Status: Enabled
            ReplicationConfiguration:
              Role: !GetAtt ReplicationRole.Arn
              Rules:
              - Destination:
                  Bucket: !GetAtt DestinationBucket.Arn
                  ReplicationTime:
                    Status: Enabled
                    Time: 
                      Minutes: 15
                  Metrics:
                    Status: Enabled
                    EventThreshold:
                      Minutes: 15
                Prefix: ""
                Status: Enabled
## [5][docker container - simplest way to host?](https://www.reddit.com/r/aws/comments/iuiuza/docker_container_simplest_way_to_host/)
- url: https://www.reddit.com/r/aws/comments/iuiuza/docker_container_simplest_way_to_host/
---
Hi,

All my experience uptil now is on kubernetes.

I am prototyping an idea and need a cheap and simple way to deploy it on aws. 

Requirements:

\- supports container

\- needs RDS

\- needs internet access (for external apis)

\- should support basic ci/cd pipeline. (I use gitlab)

&amp;#x200B;

ECS seems the route but as soon as I put it inside vpc, it lose internet and nat gateway is way expensive for small prototype!

&amp;#x200B;

Thanks.
## [6][Anyone else having console issues in us-east-1 right now?](https://www.reddit.com/r/aws/comments/itwd8q/anyone_else_having_console_issues_in_useast1/)
- url: https://www.reddit.com/r/aws/comments/itwd8q/anyone_else_having_console_issues_in_useast1/
---
Getting a lot of errors in EC2 right now.
## [7][Following the Juli 2020 ruling of the EU court that the EU/US Privacy shield is invalid: Will we be able to use AWS for customer projects in the near future?](https://www.reddit.com/r/aws/comments/iuf9uu/following_the_juli_2020_ruling_of_the_eu_court/)
- url: https://www.reddit.com/r/aws/comments/iuf9uu/following_the_juli_2020_ruling_of_the_eu_court/
---
I have heard several opinions on the EU/Privacy Shield Rulings and the german politics is pretty much silent on the specifics so far. Given that we (In Germany) have no real alternative to AWS here, the situation seems to be that AWS, Azure, Google Cloud, Digital Ocean etc. are not good for projects storing customer sensitive data (aka everything with a login).   
So fellow Europeans, Germanz and other people of the world: What do you think is the case?
## [8][Licensing of MS products in AWS - mostly for "desktop-like" use](https://www.reddit.com/r/aws/comments/iuhuke/licensing_of_ms_products_in_aws_mostly_for/)
- url: https://www.reddit.com/r/aws/comments/iuhuke/licensing_of_ms_products_in_aws_mostly_for/
---
For those in need to run MSOffice (incl. Outlook) apps in a "desktop like" setup, in AWS, I think there are a few options:

- AWS workspaces - which is actually Windows server, with a "Windows 10 like skin" (a.k.a. "desktop experience")

- EC2 w/Windows 10 (maybe as described [here](https://gist.github.com/peterforgacs/abebc777fcd6f4b67c07b2283cd31777)) - vague concerning licensing

- Windows 10 in dedicated instance or dedicated hosting - [per MS](https://aws.amazon.com/windows/faq/#byol-win-cl), although I have no idea why someone would choose a dedicated **host** for a desktop need

- an actual Windows server

Someone pointed us the other day to the [following document](https://blogs.gartner.com/stephen-white/2020/06/10/organizations-running-office-365-windows-applications-and-windows-10-in-amazon-google-and-alibaba-should-re-check-licensing/), which adds to the concern I wanted to address, thus the true reason for my post: has someone actually figured out the **cost optimal** (OS license, O365 license, Outlook license, etc., etc.) solution of running applications (someone was even suggesting Office Plus) - mostly Excel, Word and also Outlook, in a desktop like fashion (users logged in and using the environment as "desktop")?
## [9][We are the AWS EC2 Team - Ask the Experts - Sep 24th @ 9AM PT / 12PM ET / 4PM GMT!](https://www.reddit.com/r/aws/comments/iu0c8d/we_are_the_aws_ec2_team_ask_the_experts_sep_24th/)
- url: https://www.reddit.com/r/aws/comments/iu0c8d/we_are_the_aws_ec2_team_ask_the_experts_sep_24th/
---
Hey r/aws! u/AmazonWebServices here.

The AWS EC2 team will be hosting an Ask the Experts session here **in this thread** to answer any questions you may have about **deploying your machine learning models to Amazon EC2 Inf1 instances powered by the** [**AWS Inferentia chip**](https://aws.amazon.com/machine-learning/inferentia/), which is custom designed by AWS to provide high performance and cost-effective machine learning inference in the cloud. These instances provide up to 30% higher throughput, and 45% lower cost per inference over comparable GPU-based instances for a wide variety of machine learning use cases such as image and video analysis, conversational agents, fraud detection, financial forecasting, healthcare automation, recommendation engines, text analytics, and transcription. It's easy to get started and popular frameworks such as TensorFlow, PyTorch, and MXNet are supported.

Already have questions? Post them below and we'll answer them starting at 9AM PT on Sep 24, 2020!
## [10][Starting with AWS SDK](https://www.reddit.com/r/aws/comments/iugy1y/starting_with_aws_sdk/)
- url: https://www.reddit.com/r/aws/comments/iugy1y/starting_with_aws_sdk/
---
Going to build some application to integrate with aws, e.g. getting some cloudwatch metrics. Any way i can make request to aws and test the api free? Any sample data response?
