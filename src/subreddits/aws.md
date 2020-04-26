# aws
## [1][Help me find that blog with nice AWS diagrams...](https://www.reddit.com/r/aws/comments/g82rue/help_me_find_that_blog_with_nice_aws_diagrams/)
- url: https://www.reddit.com/r/aws/comments/g82rue/help_me_find_that_blog_with_nice_aws_diagrams/
---
Hi guys, I'm not able to find a blog I saw some time ago. Its articles were about AWS technologies and services, explained with nice hand drawn diagrams. They were very thorough and detailed, and explained complex stuff like VPCs with just a glance! Help!
## [2][Coordinating resource creation and modification in a small team](https://www.reddit.com/r/aws/comments/g8bx6i/coordinating_resource_creation_and_modification/)
- url: https://www.reddit.com/r/aws/comments/g8bx6i/coordinating_resource_creation_and_modification/
---
I'm the AWS coordinator for our small company and I'm having a little trouble making sure that our usage is well coordinated.There are a few team members who are responsible for creating and configuring EC2 servers and other resources and often not enough thought goes into resource creation.I'm not directly in their chain of command, so I'm not always part of the discussion of what's required.

I'd like to have some oversight when servers are created or significantly modified, so that I can discuss requirements before the server is created and when it needs to be modified.

I'm imagining a system where a team member can create a request that represents a potential resource with a number of config options selected.They then pass it to me for approval and then I can make suggestions and have a discussion.It would also be good if this system could be a place for us to notify each other of changes that we made that were too minor to require a discussion.

A feed of events would also make sense.

Any insights here of how you're doing things or what processes or products you would suggest?  


Edit:   
We're not yet using CloudFormation or CI/CD.  
We're pretty new to the cloud, until now we were distributing software packages to customers to install on prem.  
For now almost all our interaction with AWS is via the GUI.
## [3][Architecture deep-dive: How a news aggregator collects and serves up news articles](https://www.reddit.com/r/aws/comments/g7smug/architecture_deepdive_how_a_news_aggregator/)
- url: https://www.reddit.com/r/aws/comments/g7smug/architecture_deepdive_how_a_news_aggregator/
---
Hey guys,

For today's AWS Deep Dive I've interviewed the creators of a news aggregator and conducted an architecture review of their design.

[https://www.youtube.com/watch?v=XaO8BBgBRH0](https://www.youtube.com/watch?v=XaO8BBgBRH0)

Thanks,

Vladimir

P.S. Let me know what else you'd like me to make
## [4][Scaling Our AWS Infrastructure](https://www.reddit.com/r/aws/comments/g80yg0/scaling_our_aws_infrastructure/)
- url: https://medium.com/swlh/scaling-our-aws-infrastructure-9e64e6817b8c
---

## [5][EC2 Import rejected kernel version](https://www.reddit.com/r/aws/comments/g86qju/ec2_import_rejected_kernel_version/)
- url: https://www.reddit.com/r/aws/comments/g86qju/ec2_import_rejected_kernel_version/
---
I created recently an EC2 instance of Gentoo running kernel 5.4.28 and it's working just fine. I did an export of this instance to my local server and started under VirtualBox, did some few updates and changes to export back as new instance but surprisingly I got an error saying that "Kernel version 5.4.28 is not supported".

Considering that I have already an instance running on 5.4.28, what am I doing wrong?
## [6][AWS Macie and cloudformation](https://www.reddit.com/r/aws/comments/g85y5l/aws_macie_and_cloudformation/)
- url: https://www.reddit.com/r/aws/comments/g85y5l/aws_macie_and_cloudformation/
---
I have been doing some research on using cloudformation to turn on macie and integrating member accounts in macie but cannot get any clear documentation of achieving this with cloudformation. I know I can turn it on via console but want to do this via iac(cloudformation or terraform). Anyone have experience or documentation on this?
## [7][Is it allowed in ToS to discuss enterprise support cases online?](https://www.reddit.com/r/aws/comments/g7qbd4/is_it_allowed_in_tos_to_discuss_enterprise/)
- url: https://www.reddit.com/r/aws/comments/g7qbd4/is_it_allowed_in_tos_to_discuss_enterprise/
---
I find myself asking a lot of questions to support for stuff that I can't find already-existing info on, and so I probably open a support case at least once every two weeks or so.

The responses are always amazing, and I want to share the info I get online.  In a recent example, I found an flaw (not a security flaw) in one of the AWS provided default IAM policies, and when I wrote to support, they were like, "indeed, we tested it too, the policy is mistaken, we've escalated this to the IAM team, and to work around it here are some ways you can fill in the gaps." It seems like a shame that info is just sitting in my inbox and not on a forum somewhere.

Can I copy paste those responses somewhere?  Do I have to redact the agents name?  Or do I have to reword the issue and solution in my own words?
## [8][Can I embed Python snippets that run on lambda on my website](https://www.reddit.com/r/aws/comments/g83a2x/can_i_embed_python_snippets_that_run_on_lambda_on/)
- url: https://www.reddit.com/r/aws/comments/g83a2x/can_i_embed_python_snippets_that_run_on_lambda_on/
---
Hi,

&amp;#x200B;

I want to embed Python code snippets that the reader can edit and run like programming websites.

Some people recommend using AWS lambda.

Can I achieve that in WordPress? Is there any tutorial or any material that can help?

&amp;#x200B;

Regards,
## [9][Convert AWS EC2 Reserved instance to EC2 Spot instance. Possible?](https://www.reddit.com/r/aws/comments/g82nq8/convert_aws_ec2_reserved_instance_to_ec2_spot/)
- url: https://www.reddit.com/r/aws/comments/g82nq8/convert_aws_ec2_reserved_instance_to_ec2_spot/
---
 

Hello everyone,

I'm running an application developed in-house for a research project, in an EC2 instance (c5d.4xlarge). I have used a Reserve Instance to manage the billing, as this project was supported by AWS credits. But I would like to optimize the instance due to shortage of credits we're facing right now. I want to switch from a reserved instance to a EC2 Spot instance. However, I cannot sell my instance on the AWS marketplace because the research lab is Non-profit and part of an educational institution.   Is there a way to convert my reserved instance to a spot instance, so that I can save up on my credits? The credits are set to expire in September, but with the on-going speed of credits spent, I will be out of credit balance by July.   #ec2 #spot-instance #reserved-instance #help #amazon-web-services
## [10][Making real time dashboard (website) using API gateway, s3 and lambda?](https://www.reddit.com/r/aws/comments/g7spzt/making_real_time_dashboard_website_using_api/)
- url: https://www.reddit.com/r/aws/comments/g7spzt/making_real_time_dashboard_website_using_api/
---
I am currently making a website that monitors my raspberry pi data (which is stored in dynamoDB) and show it in real time. My aws design goal is to create a serverless realtime web app. So from what I have heard, I can add my static site (I am using vueJS) in S3 and then use the help of API gateway and lambda to get the data. The problem is I am not familiar with websockets. So please forgive me if I say something stupid; So my biggest question is the client side code. Once I have created a websocket API gateway, what should I add in my client side to make it work? I know I can use rest api instead, but since I want the site to be in realtime, I would need to call the API every few second (which is not cost effective).

If you have any advice regarding my project or implementation please feel free to give suggestions.

Thank you!
