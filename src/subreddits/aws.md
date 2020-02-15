# aws
## [1][Multi-Attach for Provisioned IOPS (io1) EBS Volumes](https://www.reddit.com/r/aws/comments/f41kb2/multiattach_for_provisioned_iops_io1_ebs_volumes/)
- url: https://aws.amazon.com/blogs/aws/new-multi-attach-for-provisioned-iops-io1-amazon-ebs-volumes/
---

## [2][Infra.app - Desktop app for managing Kubernetes](https://www.reddit.com/r/aws/comments/f3u06h/infraapp_desktop_app_for_managing_kubernetes/)
- url: https://infra.app
---

## [3][rds rss to cloudformation](https://www.reddit.com/r/aws/comments/f40o0f/rds_rss_to_cloudformation/)
- url: https://www.reddit.com/r/aws/comments/f40o0f/rds_rss_to_cloudformation/
---
Hi all,

I'm new to aws and cloudformation.  I'm looking to convert one of my auroradb rds rss that I created in the website/gui into a cloudformation template.  I have gone to the cloudformation page, click create stack ==&gt; with existing resources (import resources) and it is asking for S3 bucket or upload a template, which I don't have.  Any useful article I can use to accomplish this?

&amp;#x200B;

Thanks,
## [4][Protect S3 Images with own IAM Concept](https://www.reddit.com/r/aws/comments/f3r7vh/protect_s3_images_with_own_iam_concept/)
- url: https://www.reddit.com/r/aws/comments/f3r7vh/protect_s3_images_with_own_iam_concept/
---
I have an App that should store and retrieve images from S3. Currently, upload happens via my own application server endpoint to S3 (image on app-&gt;application server-&gt;store in S3).

However, whats the best way to protect image access then, without giving every single user some IAM role on S3 nor share a direct S3 link?

My only solution would be that my application servers retrieve the image, and then forward them to the client. This causes two HTTP roundtrips though (S3-&gt;Application Server-&gt;App). This is massive for huge pictures.

Any other ideas?
## [5][AWS central authentication advises.](https://www.reddit.com/r/aws/comments/f44gne/aws_central_authentication_advises/)
- url: https://www.reddit.com/r/aws/comments/f44gne/aws_central_authentication_advises/
---
What is best way to build central authentication for people to access to AWS resources, like EC2? I am looking into Microsoft Active Directory, FreeIPA, AWS Directory Service. MS AD require license and support, FreeIPA is open source and can purchase 3rd party support, AWS Directory Service is more integrated with the rest of the services.

Also, I am researching on what is the best option for client base VPN. This is for people to connect to the VPN and then ssh or rdp EC2 resources. I know AWS has a client VPN service offering under the VPC. Is this difficult to setup? Or should I look into virtual firewall, like PfSense or Palo Alto?

My Goal is to setup an VPN for users to login, and then they can login to a bastion server or some AWS resources
## [6][DynamboDb and Databricks](https://www.reddit.com/r/aws/comments/f41es1/dynambodb_and_databricks/)
- url: https://www.reddit.com/r/aws/comments/f41es1/dynambodb_and_databricks/
---
Any clear explanation on how I can connect DynamoDB with databricks? I want to be able to send one data of row to dynamoDB from databricks.

If my question isn’t clear please let me know I will try my best to reword it.

Thanks
## [7][CloudFormation custom resource repository?](https://www.reddit.com/r/aws/comments/f3ws17/cloudformation_custom_resource_repository/)
- url: https://www.reddit.com/r/aws/comments/f3ws17/cloudformation_custom_resource_repository/
---
Hello,

Is there a CloudFormation custom resource repository, that supplements missing CFN support?

Specifically, I am looking for a way to create and assign Route53 reusable delegation sets. I can do this via the aws cli with one command, so hoping I could "quickly" spin up a Lambda function that would run the command to do the creation and another function to do the assignment.

Thanks!
## [8][Review: Amazon Connect – A Programmable Telephone System](https://www.reddit.com/r/aws/comments/f3xbeg/review_amazon_connect_a_programmable_telephone/)
- url: https://cloudonaut.io/review-amazon-connect-programmable-telephone-system/
---

## [9][Matlab on AWS Lambda](https://www.reddit.com/r/aws/comments/f433tj/matlab_on_aws_lambda/)
- url: https://www.reddit.com/r/aws/comments/f433tj/matlab_on_aws_lambda/
---
A new startup marrying Matlab and AWS Lambda :  
[http://www.matlambda.com/](http://www.matlambda.com/)
## [10][Moving from On Premise to AWS Aurora Serverless](https://www.reddit.com/r/aws/comments/f3xrtx/moving_from_on_premise_to_aws_aurora_serverless/)
- url: https://www.reddit.com/r/aws/comments/f3xrtx/moving_from_on_premise_to_aws_aurora_serverless/
---
Moving from on-premise to AWS Aurora Serverless

Hi devops,

We are currently maintaining a MySQL database on bare metal, which handles all of our production load. It has served us well, but in the past months, it has started to become too small for our load during peak times. We have a very inconsistent load on our production, where during the days it is very low, but peaks during evenings and weekends. While idle, we have somewhere around 100k hits/hr, while during peak it could go as high as 5M/hr or sometimes even 15M/hr. These spikes happens about every second to every third day, so not so frequently.

I have looked into different options on how to tackle this issue, and to find a long term solution that would suit our needs and take the issue off of our hands. Naturally, since we are in the AWS echo system with a lot of our services, we looked into what they could offer and found out about Aurora Serverless. It has got some limitations when it comes to functionality, but it all seems fine to us and should not bother us in any way as of how we are using MySQL right now.

Searching for experience on maintaining production workload on Aurora Serverless is quite small, at least according to my Google-fu. Most users seem to use it for test environments and staging since it is able to scale down to 0 ECU's, and therefore keeping the cost down.

Our main take is that it would suit our needs based on the spiky loads, and therefore not having to set up an instance able to handle those large loads while staying idle most of the time, at least if it works as good as AWS states that it does, and thereby keep the costs down.

Does anyone have actual experience with Aurora Serverless in production environments? Are there better solutions for scaling things like this, when it comes to AWS and MySQL?
