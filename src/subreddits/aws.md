# aws
## [1][90% Price cut for AWS IoT Core Jobs](https://www.reddit.com/r/aws/comments/gfq14z/90_price_cut_for_aws_iot_core_jobs/)
- url: https://www.reddit.com/r/aws/comments/gfq14z/90_price_cut_for_aws_iot_core_jobs/
---
AWS just announced a 90% price reduction for their IoT Core Jobs service. While price reductions from AWS are common, the magnitude of this one is not. I wrote a bit about that this week plus other DevOps relevant news in the weekly roundup for [This Week in DevOps](https://thisweekindevops.com/2020/05/08/weekly-roundup-may-8th-2020/).

Has anyone here used IoT Jobs or have any idea why they reduced prices so drastically?
## [2][AWS: CloudFormation — using lists in Parameters](https://www.reddit.com/r/aws/comments/gfondy/aws_cloudformation_using_lists_in_parameters/)
- url: https://medium.com//aws-cloudformation-using-lists-in-parameters-52982e78384f?source=friends_link&amp;sk=e2bcf0b78581010200e619e8da7c016f
---

## [3][What's the best way to clean up abandoned resources on multiple region/services?](https://www.reddit.com/r/aws/comments/gfre80/whats_the_best_way_to_clean_up_abandoned/)
- url: https://www.reddit.com/r/aws/comments/gfre80/whats_the_best_way_to_clean_up_abandoned/
---
Well, I don't know nearly 90% of the resources and its usage (ec2, ebs, rds, ecs, etc...) and trying to find a way to determine if they are still actively using.
Can anyone recommend me any software/practice to handle this?
## [4][How to manage multiple aws cli config ?](https://www.reddit.com/r/aws/comments/gfp3xp/how_to_manage_multiple_aws_cli_config/)
- url: https://www.reddit.com/r/aws/comments/gfp3xp/how_to_manage_multiple_aws_cli_config/
---
I have access to 3 AWS accounts and using profile in aws config we can manage multiple keys . But I am trying for something like creating a source file for each profile instead of hardcoding them  and source the required file whenever required in bash , so no need to hardcode these values in .aws/credentials . Once I exit the current shell the access keys shouldn't be available until I source them again .
## [5][Running a scheduled fargate task](https://www.reddit.com/r/aws/comments/gfp29y/running_a_scheduled_fargate_task/)
- url: https://www.reddit.com/r/aws/comments/gfp29y/running_a_scheduled_fargate_task/
---
Hi,

I’m having a problem running a scheduled fargate task. What I want to happen is for the task to run to completion once, then wait for the next execution time and run again.

What happens is that the task runs once and gets restarted by fargate immediately because the container exited.

What is the secret to letting it complete the first time, get unloaded, and then run again later?
## [6][How is Reddit hosted on S3 when it's clearly a dynamic and not static website?](https://www.reddit.com/r/aws/comments/gf9obs/how_is_reddit_hosted_on_s3_when_its_clearly_a/)
- url: https://www.reddit.com/r/aws/comments/gf9obs/how_is_reddit_hosted_on_s3_when_its_clearly_a/
---
From my understanding, S3 can only host static websites, so how is Reddit supposedly hosted on S3? It uses JavaScript and other dynamic website features.
## [7][Method to maintain absolute least privileged automatically?](https://www.reddit.com/r/aws/comments/gft779/method_to_maintain_absolute_least_privileged/)
- url: https://www.reddit.com/r/aws/comments/gft779/method_to_maintain_absolute_least_privileged/
---
I was wondering what the best way would be to remove a role or policy from a user on say a 90 day or just a X time basis automatically. The goal is say User was granted read S3, but that user only needed it one time and doesn't use read S3 for say 90 days, then I would want to revoke that policy automatically. Anyone know an efficient way to go about this?

&amp;#x200B;

My idea was to have a script that runs once a day to scan logs for users to see the last day they have done that action and if it is greater then the cut off point then revoke the role. This seems like it would take alot of time and processing to do this though.

&amp;#x200B;

Just looking for some ideas to get me going in the right direction thanks!
## [8][Can you use both DTMF and Lex in a Connect Flow?](https://www.reddit.com/r/aws/comments/gft51o/can_you_use_both_dtmf_and_lex_in_a_connect_flow/)
- url: https://www.reddit.com/r/aws/comments/gft51o/can_you_use_both_dtmf_and_lex_in_a_connect_flow/
---
We have a lex bot that handles 1 lookup based on speech input from the customer. It works well.

Our connect flow has a few text-to-speech responses. 

We have one get customer input block that’s handling all of this... 

Press 1 starts the Lex
Press 2, 3, 4 plays a prompt

Looks like 2, 3, and 4 go through lex even though there’s not for lex to do.

Is there a way to separate these so DTMF doesn’t have to go through lex.

Thx!
## [9][AWS Architecture idea](https://www.reddit.com/r/aws/comments/gfcpnv/aws_architecture_idea/)
- url: https://www.reddit.com/r/aws/comments/gfcpnv/aws_architecture_idea/
---
Hi all,

I recently joined a company as an intern/junior Python Developer and I am working on an internal project. They want to use AWS and app need to be serverless. The app itself is pretty simple, they want an 'address-book' type of program, so they can track their clients. In short, I need to make a CRUD app and my idea is to use: S3 for the fronted, Lambda for the computing/code, DynamoDB as a database and Cognito for authentication and also a Gateway API for the communication. Do you think it is a good idea? If not, why?

I am still very new to AWS, around 2 weeks of experience and there is too much choice, so I wanted to share this and hear thoughts of more experienced people here.

Thanks in advance!

Edit: I wrote verification instead of authentication.

Edit2: First of all, thank you all for the constructive answers! You're awesome! Some people mentioned that it would be like reinventing the wheel and that there are already cheap/free solutions for the problem, which I am aware of and I agree, but I probably needed to mention that one of the main goals of the project, besides its use, will be me gaining an actual experience by going through the whole process of working on a real-world project. Hope it helps!
## [10][Lightsail instance firewall now supports sourceIP rules...and also PING!](https://www.reddit.com/r/aws/comments/gfe0ti/lightsail_instance_firewall_now_supports_sourceip/)
- url: https://www.reddit.com/r/aws/comments/gfe0ti/lightsail_instance_firewall_now_supports_sourceip/
---
Hello redditors,

Hope y'all are keeping safe. We released a few enhancements to Lightsail instance firewall which many of you requested. Lightsail firewall rules now support source IP conditions and also, ICMP-PING. So, you can now open PING to your Lightsail instances.

Here's a blog to get started - [https://aws.amazon.com/blogs/compute/enhancing-site-security-with-new-lightsail-firewall-features/](https://aws.amazon.com/blogs/compute/enhancing-site-security-with-new-lightsail-firewall-features/)

As always, please keep the feedback coming. Thank you! 

&amp;#x200B;

PS: Some resources that may come in handy if you haven't used Lightsail yet-

Bunch of getting started guides - [https://aws.amazon.com/lightsail/resources/](https://aws.amazon.com/lightsail/resources/)

Lightsail pricing- [https://aws.amazon.com/lightsail/pricing/ ](https://aws.amazon.com/lightsail/pricing/)
