# aws
## [1][AWS Global Accelerator: Terrible Name, Awesome Service](https://www.reddit.com/r/aws/comments/fsba7m/aws_global_accelerator_terrible_name_awesome/)
- url: https://www.sentiatechblog.com/aws-global-accelerator-terrible-name-awesome-service?utm_source=reddit&amp;utm_medium=social&amp;utm_campaign=global_accelerator
---

## [2][Docker desktop creators built a Kubernetes management tool](https://www.reddit.com/r/aws/comments/fs1lvy/docker_desktop_creators_built_a_kubernetes/)
- url: https://infra.app/
---

## [3][Now Open – Third Availability Zone in the AWS Canada (Central) Region](https://www.reddit.com/r/aws/comments/frrqky/now_open_third_availability_zone_in_the_aws/)
- url: https://aws.amazon.com/blogs/aws/now-open-third-availability-zone-in-the-aws-canada-central-region/
---

## [4][ECS Task IAM Role for cross-account resource sharing.](https://www.reddit.com/r/aws/comments/fsa78f/ecs_task_iam_role_for_crossaccount_resource/)
- url: https://www.reddit.com/r/aws/comments/fsa78f/ecs_task_iam_role_for_crossaccount_resource/
---
I have a ECS task that needs to access resources in another aws account. 

My ECS tasks are running in Account A and the resources that I need to access are in account B.

For this I have created a role in account B with necessary  permission to the required resources and I have added a trust relationship with account A as well.

Now in account A where ECS is running I have added  sts:AssumeRole to the Remote IAM role in account B.

This however is not working. My ECS task is still not being able to access resourcecs in account B.

&amp;#x200B;

How do i set up cross-account IAM roles/policies that work with ECS  Task.
## [5][How to automate workspace?](https://www.reddit.com/r/aws/comments/fsaz6p/how_to_automate_workspace/)
- url: https://www.reddit.com/r/aws/comments/fsaz6p/how_to_automate_workspace/
---
Hello guys,

how do you automate your virtual desktops? Because after the creation of the image, the sysprep/general process killed understandably my user settings. Now I have to set things like language, default browser, etc. manually on every machine.  


Do you use something like PS scripts, autostart scripts in the default image or something similiar? I cant find anything about answer.xlm files or something else. 

PS. my native language isnt EN, sorry for some grammatical mistakes and I am a beginner with AWS.   


Thanks in advance. Greetings Manuel
## [6][Manage infrastructure Programmatically](https://www.reddit.com/r/aws/comments/fs5831/manage_infrastructure_programmatically/)
- url: https://www.reddit.com/r/aws/comments/fs5831/manage_infrastructure_programmatically/
---
I am currently working on a project which allows engineers to manage their infrastructure by using Terrafrom in conjunction with Protocall Buffers.  This allows the developer to use a language of their choice without relying on data definition languages. If you have any comments, thoughts or ideas about the project, raise an issue.  
[https://github.com/jackskj/terraform-pb](https://github.com/jackskj/terraform-pb)
## [7][Nearest Lightsail instance region for South Africa](https://www.reddit.com/r/aws/comments/fsaqpg/nearest_lightsail_instance_region_for_south_africa/)
- url: https://www.reddit.com/r/aws/comments/fsaqpg/nearest_lightsail_instance_region_for_south_africa/
---
I'm currently setting up lightsail for wordpress and I'm trying to find out which is the instance location with the lowest latency for South Africa. Does anyone here knows or has a similar experience?

Geographically, Mumbai is the nearest. Would it be a safe bet to go with Mumbai?

https://preview.redd.it/6a2wzpqcmzp41.jpg?width=3657&amp;format=pjpg&amp;auto=webp&amp;s=9d4524fc2c25ee3c48e875c7a10ac7131c71c63c
## [8][HELP: CloudFront subdirectory hosting multiple angular apps](https://www.reddit.com/r/aws/comments/fs9nnn/help_cloudfront_subdirectory_hosting_multiple/)
- url: https://www.reddit.com/r/aws/comments/fs9nnn/help_cloudfront_subdirectory_hosting_multiple/
---
Hi there,

I have several angular apps, which I want to host on S3 in a single bucket but under different sub directories(Only for the sake of maintenance of a single bucket )

Note: I'm aware and have been hosting single app in its own bucket.

I have been trying to do three things 

1) Cloudfront CDN for s3 sub directory hosting

2) preventing Subdirectory folder name in the URL

3) Making Angular routes work with the cloudfront. 

So far I have had success with the 1st point.

I have been trying to do this for like 3 hours now. but unable to achieve a whole solution.

Please guide me if its even possible to do that or not as no solution is available on the net properly.
## [9][AWS Cognito](https://www.reddit.com/r/aws/comments/fs9hrc/aws_cognito/)
- url: https://www.reddit.com/r/aws/comments/fs9hrc/aws_cognito/
---
I don't quite understand the documentation and I hope this does not sound like a stupid question. 

I have a subdomain (xxx.yyy.com) which i need to secure with MFA. This sub-domain is running a on-prem setup of Dynamics CRM.  It is possible for me to use AWS Cognito, create a user pool with a fixed list of users and at the domain to it? I am also using EC2 to run the AD server instead of managed AD.
## [10][Monitor ALL EC2 instances in a region for status check failures.](https://www.reddit.com/r/aws/comments/fs3t0h/monitor_all_ec2_instances_in_a_region_for_status/)
- url: https://www.reddit.com/r/aws/comments/fs3t0h/monitor_all_ec2_instances_in_a_region_for_status/
---
I am looking for ideas on what combination of AWS services I can use.

The idea is to have a check in place to send an alert (email, SNS, Pagerduty) of some kind if health checks on ANY of the EC2 instances in a REGION fails. Ignore any EC2 instances that are stopped.

The alarm should work if new instances are added or old instances are removed from the region. Would a lambda function do the job for something like this?
