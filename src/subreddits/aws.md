# aws
## [1][Week of Sept 21st - What are you building this week in AWS?](https://www.reddit.com/r/aws/comments/ix0dmq/week_of_sept_21st_what_are_you_building_this_week/)
- url: https://www.reddit.com/r/aws/comments/ix0dmq/week_of_sept_21st_what_are_you_building_this_week/
---
Share what you are working on
## [2][AWS billing is like the Babadook](https://www.reddit.com/r/aws/comments/ixgafg/aws_billing_is_like_the_babadook/)
- url: https://www.reddit.com/r/aws/comments/ixgafg/aws_billing_is_like_the_babadook/
---
tl;dr: I cancelled an AWS account months ago. Had some hassle with inflexible billing policies already. Now Route 53 is auto-renewing a domain name on the account unless I log in to prevent it. I cannot log in because the account is closed.

Sorry for the long rant. Just blowing off steam here.

I opened an AWS account for my own research and teaching needs while at a former employer. The bills were handled through our department, but the account was my responsibility. All was good at the time, but I no longer work there, and terminating an AWS account turns out to be nearly impossible!

Two months after leaving that job, I got a message from the accountant in my former office. It turned out I was still being billed for a reserved EC2 instance, and there was no possible way to just pay the remaining charges all at once. Instead, I will keep getting bills until the instance's yearly agreement has run out. It would easily be possible to pay off the remainder via institutional funds now, but I can't keep sending tiny bills to my former employer months after I've left that job. So, now I'm paying those out of pocket until January.

Dealing with that was a hassle, because I no longer have access to the institutional email address with which I created the AWS account, and contacting customer support from another address doesn't fly, for obvious reasons. When I finally did get things sorted out, the customer service person assured me that those were the only remaining charges.

But now it turns out that a domain name I started via Route 53 will auto renew this month unless I log in and change it (which I can't do, because the account is closed). Now I'm going to have to go through all the same rigamarole as last time to get an agent to disable the Route 53 domain registration.

Why in the world can I not just say "I no longer want this service. Please stop taking money from me." Why do I have to comb through every one of AWS's hundreds of surprisingly non-integrated services to make sure there are no hidden automatic renewals? And why is it possible to "close" an account to the point that I cannot make changes, even if they will keep billing me for stuff unless I make changes?

Every time I think I've ridden myself of this cursed account, it shows back up like a goddam Babadook!

Edit: I think I should have been more clear about my complaint here. I understand what a reserved instance is. It would be nice if they would bill out the remaining 3 months at once, but that's not the biggest problem. The two things that are truly frustrating are:

  1. **Nothing should auto renew on a closed account!!** As it stands, Route 53 will auto renew a domain unless I log in to stop it. I cannot log in, because the account is closed.

  2. When canceling an account, there should be clear information about what remaining charges exist on the account. The account should not truly lock until there are no outstanding charges to be billed.
## [3][Which serverless framework is used more in AWS related jobs?](https://www.reddit.com/r/aws/comments/ixliza/which_serverless_framework_is_used_more_in_aws/)
- url: https://www.reddit.com/r/aws/comments/ixliza/which_serverless_framework_is_used_more_in_aws/
---
There is SAM, Serverless, Claudia, etc.   Which one seems to be more prevalent in the workforce?
## [4][Security September: Cataclysms in the Cloud Formations – One Cloud Please](https://www.reddit.com/r/aws/comments/ixjh7u/security_september_cataclysms_in_the_cloud/)
- url: https://onecloudplease.com/blog/security-september-cataclysms-in-the-cloud-formations
---

## [5][nab3: An async boto3 ORM](https://www.reddit.com/r/aws/comments/ixmv93/nab3_an_async_boto3_orm/)
- url: https://www.reddit.com/r/aws/comments/ixmv93/nab3_an_async_boto3_orm/
---
[https://github.com/WillNye/nab3](https://github.com/WillNye/nab3)

nab3 provides an easy to use interface for searching and inspecting a service and its related resources within AWS. It also creates useful relationships that aren't returned within boto3. An example of this is returning the Security Groups for an ASG or an ECS Cluster (outlined in the README).

More AWS services still need to be added so any contributors are welcome. There is a dedicated doc on contributing because as you may imagine trying to normalize boto3 can get hairy.
## [6][Amazon Athena Custom/Dynamic Front End?](https://www.reddit.com/r/aws/comments/ixmifw/amazon_athena_customdynamic_front_end/)
- url: https://www.reddit.com/r/aws/comments/ixmifw/amazon_athena_customdynamic_front_end/
---
I know building a front end for Athena is fairly common, but I’m looking for some resources on a previously executed frontend for Athena with user generated query inputs.

I want to use some kind of Intellisense to help the user out too, so if you have any library recommendations pls let me know.

Thanks!
## [7][Dear AWS, Please stop spamming me](https://www.reddit.com/r/aws/comments/iwzrqz/dear_aws_please_stop_spamming_me/)
- url: https://www.reddit.com/r/aws/comments/iwzrqz/dear_aws_please_stop_spamming_me/
---
Dear AWS, 

Like many, I'm trying to follow 'best practices' - using AWS Control Tower and Multi-Account configuration.

However you're making this rather painful - every account created results in me getting yet another subscription to your marketing emails. 

Today I got a dozen more "Welcome to your Getting Started series" emails.   
As if that wasn't painful enough, un-subscribing is made unnecessarily difficult - I need to copy the email address over and tell you, once again, that I really don't want more marketing material. 

Given the amount of tracking in all those links, the least you could do would be to provide a one-click unsubscribe.  

Better still - how about for every account created from AWS Orgs / Control Tower you opt them out of marketing emails. 

Sincerely Frustrated, 

L. Extension.
## [8][cloudwatch dashboard - account costs single widget](https://www.reddit.com/r/aws/comments/ixmeso/cloudwatch_dashboard_account_costs_single_widget/)
- url: https://www.reddit.com/r/aws/comments/ixmeso/cloudwatch_dashboard_account_costs_single_widget/
---
hi,

i've created a single custom cloudwatch dashboard for my product with relevant metrics, what i wouldn't mind is a single number shown in the widget with cost estimate, ideally for the tagged stuff i'm using but for the whole account is fine. 

i think the answer is no, but doesn't anyone know if this is possible and if so how?

thanks.
## [9][anyone have ever done - docker vault image hosted in AWS ECS](https://www.reddit.com/r/aws/comments/ixm619/anyone_have_ever_done_docker_vault_image_hosted/)
- url: https://www.reddit.com/r/aws/comments/ixm619/anyone_have_ever_done_docker_vault_image_hosted/
---
if so, a little help of guidance may be a great of help  
requirement  
vault docker image in AWS ECS (EC2)
## [10][Status check of stopped ec2 instance in Ok state?](https://www.reddit.com/r/aws/comments/ixlskc/status_check_of_stopped_ec2_instance_in_ok_state/)
- url: https://www.reddit.com/r/aws/comments/ixlskc/status_check_of_stopped_ec2_instance_in_ok_state/
---
I have aws ec2 instance and I created status check alarm in “Status Checks” tab of the instance. 

I expect when instance have fulldisk or overload cpu, the alarm will be triggered as [aws doc](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/monitoring-system-instance-status-check.html). 

But I also expect when the instance in stopped state, alarm should also be triggered as status check should fail?

Is there a way to test or manually test this alarm of ec2 instance? As I want to trigger alarm and send mesg to sns topic for lambda to process. 

Is there a json object for this sns come with this kind of alarm?
## [11][What language for good paying AWS based jobs?](https://www.reddit.com/r/aws/comments/ixexpb/what_language_for_good_paying_aws_based_jobs/)
- url: https://www.reddit.com/r/aws/comments/ixexpb/what_language_for_good_paying_aws_based_jobs/
---
I've been a .NET developer my whole career, primarily in large corporate environments.  Several years ago I started embarking on a journey with AWS out of personal interest.   I'm now thinking I want to venture out of the corporate world and into more modern cloud based full stack development, with a focus on AWS services. 

However, one thing has become very clear to me in looking at job postings for the past several months.  There seems to be a fairly big pay disparity between .NET developers and Python/Node developers in the context of cloud based development, with Node/Python generally showing higher salaries. 

I've had some experience with Node, but zero with Python.  I know this puts me at a disadvantage but I'm hoping my umpteen years as a senior developer would give me some leverage. 

My question for this post is, what is the more popular language for AWS based development?  Node or Python?
