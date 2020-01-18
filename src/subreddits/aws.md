# aws
## [1][Amazon ECS Preview Support for EFS file systems Now Available](https://www.reddit.com/r/aws/comments/eq8hi6/amazon_ecs_preview_support_for_efs_file_systems/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/01/amazon-ecs-preview-support-for-efs-file-systems-now-available/
---

## [2][How do you audit user actions when assuming roles?](https://www.reddit.com/r/aws/comments/eqaxsy/how_do_you_audit_user_actions_when_assuming_roles/)
- url: https://www.reddit.com/r/aws/comments/eqaxsy/how_do_you_audit_user_actions_when_assuming_roles/
---
Title. When IAM users just have permission to assume roles into other accounts that then do things, how do you use Cloud Trail to audit what specific users did and what specific user made some change?
## [3][You’ve heard it before, outrageous bill on accident](https://www.reddit.com/r/aws/comments/eq08s9/youve_heard_it_before_outrageous_bill_on_accident/)
- url: https://www.reddit.com/r/aws/comments/eq08s9/youve_heard_it_before_outrageous_bill_on_accident/
---
I created a s3 event to compliment my lambda function with a object created event. What I didn’t think of was, my lambda creates s3 objects. So after testing my cool new event and seeing it worked I unknowingly tucked into bed and fell into a deep slumber. Upon waking in the morning I got a AWS free tier limit email. This got me thinking, logged into AWS and see the event is working too good. It’s been running my lambda non stop since last night. I finally got the event to stop after disabling didn’t work I deleted. Got it. As I head over to the billing console I notice a $900 bill. Oh shhhhh. 

Can this be undone?

Update 1: I was refunded $900 by AWS support to take care of the bill!

Update 2: After the last call to my lambda function was this morning, and all of my AWS resources were deleted around 12pm with support of AWS. 5 hours later my bill jumped from $900 to $1500. Not sure what happened
## [4][Today I was just voluntold that I'll be responsible for drafting guidelines and policy for how my team will be operating in AWS. I could use some wisdom.](https://www.reddit.com/r/aws/comments/eq9enk/today_i_was_just_voluntold_that_ill_be/)
- url: https://www.reddit.com/r/aws/comments/eq9enk/today_i_was_just_voluntold_that_ill_be/
---
Preface: I'm not technically deep into AWS.  Most of what we use it for is EC2 instances and RDS so there is a tremendous amount.  
Recently we started growing our presence in AWS.  Prior to that, it was largely a dev/test playground for my developers with only two chiefs (one being myself).  Now it's looking like we're going to start putting more services in AWS.  This will involve me bringing the rest of my team (6-8 more admins) into the environment.  Most of them won't need to operate in AWS on a daily basis.  Some will have a more frequent need to do so than others.  The bulk of it is controlled by our corporate IT department and most of what we'll be doing is in the console.  


From the beginning, I've had the feeling that if you don't get a handle on your AWS environment early and set some ground rules, things can get messy and out of control very quickly.   I feel as though the technical guidelines, things like encrypting volumes, discouraging the use of local user accounts,  principles of least privilege, complying with corporate security requirements, are pretty straight forward (though I wouldn't mind suggestions or lessons learned)  What I'm most curious about is how other folks set guidelines for how your team operates.  


While my instincts tell me I should delegate specific tasks (creation, deletion, etc.) to just a few people  I'm looking for ideas, wisdom, lessons learned, or suggestions so I don't fuck this up too bad.  My ideal scenario would be to empower my guys to do what they need to without having to seek permission too often but make sure the boundaries are clear so things stay tidy and everyone is playing by the same rules.  Basically, no one should be asking "wtf is that and why is it running?"  I'm open to ideas and I'm sure there are things I'm not thinking about but my main goal is to lay out a policy that keeps things from becoming a shit show.  Thanks!
## [5][Autoscale ECS - SQS workloads](https://www.reddit.com/r/aws/comments/eqbung/autoscale_ecs_sqs_workloads/)
- url: https://www.reddit.com/r/aws/comments/eqbung/autoscale_ecs_sqs_workloads/
---
Happy Friday!

I have an ECS service with 1 desired task. It runs a python script that polls an SQS queue. We don't have autoscale set-up, rather unsure of the best way to do so.

The workload is an ML job. The task processes one message at a time. CPU is fairly low, usually. We started doing a CW alarm that fires when #Mesaages &gt; 1, we scale up to as many tasks by using a lambda. The trick is knowing when to scale down. The #messagws visible becomes 0 immediately after they get picked up. so we wouldn't know when they get over. We could wait for N minutes and set desired to 0. We're over engineering already, adding a timed scaling worsens it. 

Hoping to hear from the community about how to do this better, or even go upstream and change input from SQS to something else. Frankly, we'd love to use the sweet AWS autoscaling features instead of hacking around.
## [6][Licensing Adobe CC on Amazon AppStream: Named-User Licensing or Shared Device Licensing](https://www.reddit.com/r/aws/comments/eq9ajv/licensing_adobe_cc_on_amazon_appstream_nameduser/)
- url: https://www.reddit.com/r/aws/comments/eq9ajv/licensing_adobe_cc_on_amazon_appstream_nameduser/
---
Hello,

To the people that installed Adobe CC on AppStream, do you use Named-User Licensing or Shared Device Licensing?

Thank you.
## [7][Simple Email Service High Availability](https://www.reddit.com/r/aws/comments/eq7laf/simple_email_service_high_availability/)
- url: https://www.reddit.com/r/aws/comments/eq7laf/simple_email_service_high_availability/
---
I'm working on high availability architecture for a web app hosted AWS. I know I can host EC2 and RDS instances in multiple AZ's, but I can't seem to find high availability information on SES (Simple Email Service). My assumption is that it's hosted in multiple AZ's by default, but you know what they say about assumptions. 

What do I need to know about SES high availability?
## [8][CloudFormation: Move resources to another stack](https://www.reddit.com/r/aws/comments/eq720t/cloudformation_move_resources_to_another_stack/)
- url: https://www.reddit.com/r/aws/comments/eq720t/cloudformation_move_resources_to_another_stack/
---
Hello,

I’ve set some Route53 resources with a DeletionPolicy of Retain. I’m now ready to delete them from the template so I can add them to another template (hit the arbitrary 200 resource record limit in a CF template managing DNS).

The blog post shows using the CF console to move the resources to another template. Is there a way to do this via CloudFormation, by using an ARN or some other identifier? I know hosted zones have an identifier but haven’t found anything for CNAMEs, etc.

Basically, trying to do it all through CF and not manually.

Thanks!
## [9][Route53: Create white label name servers via CloudFormation](https://www.reddit.com/r/aws/comments/eq4g46/route53_create_white_label_name_servers_via/)
- url: https://www.reddit.com/r/aws/comments/eq4g46/route53_create_white_label_name_servers_via/
---
Hello,

I'm looking at this guide: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/white-label-name-servers.html

In it, they list only three options to accomplish creating a reusable delegation set:

* Route 53 API

* AWS CLI

* AWS SDKs

I don't see any examples using CloudFormation, which is how I am creating the hosted zones and all of the relevant DNS entries.

Does anyone have a working example of how one might do this? I'm in the process of migrating zones and records to Route 53 for 50+ domains and the randomized 4 name servers (each different) for each domain is killing me, little by little, with every domain.

Thanks!
## [10][Learning AWS for a complete noob](https://www.reddit.com/r/aws/comments/eq5zzr/learning_aws_for_a_complete_noob/)
- url: https://www.reddit.com/r/aws/comments/eq5zzr/learning_aws_for_a_complete_noob/
---
I'm a product manager who is completely new to AWS - is there a good resource online to learn about the whole thing at the 10,000 feet level? Thank you :)
