# aws
## [1][Switching career from on-prem to cloud with aws. Who's done it and do you regret it?](https://www.reddit.com/r/aws/comments/inbk91/switching_career_from_onprem_to_cloud_with_aws/)
- url: https://www.reddit.com/r/aws/comments/inbk91/switching_career_from_onprem_to_cloud_with_aws/
---
I'm currently a site reliability engineer where I manage around 50 servers and have tried to bring an infrastructure as code approach.  The thing is I like building out systems and desiging solutions that help our business more than I like dealing with hardware which I have to do a lot of on a small team.

How many of you have made a career switch from on-prem to cloud engineer or cloud architect?  What are your thoughts and do you regret it?
## [2][Do you host any of your wordpress websites on AWS?](https://www.reddit.com/r/aws/comments/inkv4e/do_you_host_any_of_your_wordpress_websites_on_aws/)
- url: https://www.reddit.com/r/aws/comments/inkv4e/do_you_host_any_of_your_wordpress_websites_on_aws/
---
Hello, I was wondering if any of you are hosting their WordPress + WooCommerce website on AWS, I was just wondering what services are you using for it? Just simple advice and guidance 

Thanks in advance
## [3][Have you guys found Guard Duty, Macie &amp; Security Hub to be cost effective and worth enabling?](https://www.reddit.com/r/aws/comments/inlguo/have_you_guys_found_guard_duty_macie_security_hub/)
- url: https://www.reddit.com/r/aws/comments/inlguo/have_you_guys_found_guard_duty_macie_security_hub/
---
Am considering these services where you just enable them, I am curious as to people's opinion on if they are worth it for smaller accounts and just generally peoples opinions on the services
## [4][[Trivia] Why is the AWS Systems Manager abbreviated as "SSM"](https://www.reddit.com/r/aws/comments/inctog/trivia_why_is_the_aws_systems_manager_abbreviated/)
- url: https://www.reddit.com/r/aws/comments/inctog/trivia_why_is_the_aws_systems_manager_abbreviated/
---
Want to take a guess?

&amp;#x200B;

&gt;&gt;! AWS Systems Manager (Systems Manager) was formerly known as "Amazon Simple Systems Manager (SSM)" and "Amazon EC2 Systems Manager (SSM)". The original abbreviated name of the service, "SSM", is still reflected in various AWS resources, including a few other service consoles.  !&lt;

Link: [https://docs.aws.amazon.com/systems-manager/latest/userguide/what-is-systems-manager.html#service-naming-history](https://docs.aws.amazon.com/systems-manager/latest/userguide/what-is-systems-manager.html#service-naming-history)
## [5][CloudFormation template not updating after CDK deployment](https://www.reddit.com/r/aws/comments/inixr6/cloudformation_template_not_updating_after_cdk/)
- url: https://www.reddit.com/r/aws/comments/inixr6/cloudformation_template_not_updating_after_cdk/
---
I have a generated Swagger json file. I have a CDK app that reads in this file and then I do

    const cfnRestApi = this.api.node.defaultChild as apigateway.CfnRestApi;
    cfnRestApi.body = 
    JSON.parse(JSON.stringify(swaggerDefinition));


After doing cdk deploy when I look at the Cloud Formation template it contains my Swagger info. However when I make changes to the Swagger json file, the template in CloudFormation is not updating even after cdk deploy runs successfully. How can I resolve this?
## [6][Multi value routing policy](https://www.reddit.com/r/aws/comments/ink5ho/multi_value_routing_policy/)
- url: https://www.reddit.com/r/aws/comments/ink5ho/multi_value_routing_policy/
---
Can multi value routing policy in route 53 serve as a substitute for ELB? What's the difference between the two? Both of them seems to do similar work.
## [7][Building and testing Lambda + Step Function projects locally?](https://www.reddit.com/r/aws/comments/inb0vp/building_and_testing_lambda_step_function/)
- url: https://www.reddit.com/r/aws/comments/inb0vp/building_and_testing_lambda_step_function/
---
EDIT: Appreciate the responses. Was looking at AWS SAM already, so I'm going to give that a whirl at least. Not sure why the Linux binaries are connected to brew, but that's a solvable problem.

&amp;nbsp;

Haven't been able to figure this one out.

My company recently pivoted from tomcat applications hosted on containers to "serverless" applications with ALBs if needed.

I'm no expert in this arena, as it's outside where I've built solutions in the past, but it seems straightforward enough. I can't figure out a way to actually prove out a solution without running a CFT deployment to the test account though, which means

 * 1: I can only try things out on my work laptop, within the rails set up by the company. This is slow and sometimes painful, with multiple levels of review before anything can actually run.

 * 2: Screwups (think infinite loop, erroneously starting process, unbounded output sizes) have a monetary cost to company, and a very real potential life (think "Ability to put food in mouth") risk to myself.

With normal containerized applications (or non-containerized, because lets face it, Process abstraction versus Machine abstraction doesn't change how this works), whether it was running inside of a Tomcat servlet, a Node.js process, or just a basic site running in httpd, I can always start a server on one of my scratch boxes, try stuff out, and if it doesn't work the worst-case scenario is that I delete the VM / wipe the disk on the machine / reboot my dev laptop, and I'm fine.

How do I protect myself in this kind of a model?

The closest I've found is to mock out any boto3 / AWS API calls, keep all I/O to stdio/stdout, and code in ALB-style support using one of the httpd-in-a-box libraries that programming languages have picked up recently.

Is that the best I can do?
## [8][How to apply for cdn/cache server as ISP ?](https://www.reddit.com/r/aws/comments/injdxx/how_to_apply_for_cdncache_server_as_isp/)
- url: https://www.reddit.com/r/aws/comments/injdxx/how_to_apply_for_cdncache_server_as_isp/
---
Hi we are ISP in Pakistan and we would like to have amazon cache servers, how to apply ? we do not have any aws server yet.
## [9][Is there any way to run 32-bit Windows inside EC2 [2020]?](https://www.reddit.com/r/aws/comments/in7ytt/is_there_any_way_to_run_32bit_windows_inside_ec2/)
- url: https://www.reddit.com/r/aws/comments/in7ytt/is_there_any_way_to_run_32bit_windows_inside_ec2/
---
I am trying to help my mom move the software she uses to run her business which is now at the office on a Windows XP machine (really old software). Because of COVID she can't go to the office and the machine running it keeps crashing. 

I thought of moving it to an EC2 instance, but the software needs 32 bits, and I can find any 32 bits instance in EC2.

I tried running the software in compatibility mode and many other alternative ways with no luck.

After that I made an attempt to install VirtualBox inside the EC2 instance in order to run a 32bit version of windows, but quickly found out that virtualization inside virtualization doesn't seem to work.

Any suggestion will be appreciated!  
Thanks!
## [10][AWS infrastructure visibility and drift detection.](https://www.reddit.com/r/aws/comments/in1s6s/aws_infrastructure_visibility_and_drift_detection/)
- url: https://www.reddit.com/r/aws/comments/in1s6s/aws_infrastructure_visibility_and_drift_detection/
---
Hello,

I am working in a backup company which backups into the AWS cloud. Until recently our cloud infrastructure was quite simple - had clusters of EC2s, ASGs and RDS. Recently we started moving to serverless and microservices . Lot of services are now either ECS and lambda based, along with existing EC2 based ones. Due to this, our infra is becoming complex, and we kinda have feeling of losing control. Now we are bit concerned about our infrastructure. 

We need to have a **clear visibility** into our cloud infrastructure and setting up baselines, and **drift detection**. If possible we would also like to be able to run some complex queries against our AWS infrastructure, e.g. What was the longest running lambdas, in region 'us-east-1' on the cloud push date. 

We want to be able to compare specific set of resources based on date/time, region etc as well - this is more for issue RCA.

Currently we have started with AWS Config, and also experimenting with a mix of open source tools as well as vendors (Lucidchart, cloud custodian etc).

Did anyone of you had similar experience/expections? How did you solve the issue cost effectively. If possible we would not want to host any such tool on our own but rather go for a SaaS offering.

Any suggestion in this would be of great help.

Thanks in advance.
