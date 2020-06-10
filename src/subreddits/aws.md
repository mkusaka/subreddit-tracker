# aws
## [1][PSA: Don't take remote exams offered by Pearson Vue (OnVue) for AWS Certifications!](https://www.reddit.com/r/aws/comments/fscq7v/psa_dont_take_remote_exams_offered_by_pearson_vue/)
- url: https://www.reddit.com/r/aws/comments/fscq7v/psa_dont_take_remote_exams_offered_by_pearson_vue/
---
I can't describe how horrible this experience was.  I am not looking forward to how much work I am going to have to do to get my money back.  This is not my first AWS certification (I have SA Pro and DevOps Pro), but is my first online exam.  The short version is: Don't take AWS exams via the Pearson Vue at home option, even if it is offered.  AWS should not be offering this option as I can attest it is a waste of time.  Ironically, AWS would have us use their services because of their high availability and scaling but apparently they don't ask their test partners to do the same!

It started off easy enough: I passed the initial 'checks' as it confirmed my internet speed, camera access, and microphone access.  I started the process 15+ minutes before my scheduled exam time.  I was able to open the app, it again verified the technical requirements passed, and I went to the next screen.  It asked for my cell phone number and texted me a link which opened a web page which requested to take my photo.  Easy enough.  I did that and then the web page went to 'Uploading and verifying photo'.  A spinning circle started spinning.  This is where my test experience ended, but not where the poor experience ends.  I tried again, and then a third time.  Same experience.  As I write this, I left it on that page and the spinning is continuing.  This screen has been spinning for no less than 45 minutes.  At 8 minutes before my scheduled exam, I tried finding the help link.  A chat window opened, and I waited, and waited, and waited.  Still waiting as I write this.  My chat window has been open for 52 minutes and still no one to help.  Every two minutes I get ' All agents are currently assisting others. Thank you for your patience.' written in the window.  OK - what next?  They make it harder to find, but I got a phone number I can call.  I tried calling that.  Busy signal.  For the next 20 minutes I called back and back, busy signal.  Finally, I got it to actually pick up, but of course no human yet.  No estimate of time to when I can be helped.  They don't even have nice elevator music to listen to.  Who knows when I will be able to talk to someone.  This has been an exceedingly poor experience.

If you value your time, please do yourself a favor and don't even attempt a online exam with Pearson.  I worked hard to prepare for this exam and rescheduled things to fit around it.  Now, I will have to do that all again.

u/jeffbarr Is this the experience AWS is hoping to get with their testing partners?  This was a waste of my time and money.  Amazon should seriously reevaluate the quality of their test partners.  I understand everyone is trying to deal with all the issues.  However, if you can't offer quality testing, then please don't offer the option at all.  It isn't respectful to people's time.  Pearson is well aware of their capacity and if it isn't up to requirements, they shouldn't be scheduling test slots.

&amp;#x200B;

*EDIT*: A few background items I didn't initially share that may be relevant for others.  For the computer, I used a fully up to date Windows 10 laptop.  The laptop itself is only about a month old and is in near pristine condition.  Other than a few applications like Office, there is barely anything installed on there yet.  I used a hard wired connection, like recommended by Pearson through the use of a usb-to-ethernet adapter.  I have Verizon FIOS (980Mbps/840Mbps) and did do a speed test way after it was apparent this would not work.  I forget the exact numbers, but I was still pulling in hundreds of Mbps in both directions, despite everyone being at home and using the USB ethernet adapater which does put a cap on my speed, but I can't see hundreds of Mbps not being sufficent by orders of magnatude.  My phone is a fully up to date pixel 3.  I tried using my wifi in my house first (connected through FIOS), and then using the phone 4G LTE connection.  I can't imagine this was caused by my end.  It seemed like Pearson's servers were jammed at that point in time.

&amp;#x200B;

*Update*: After a LONG time, I did eventually get someone to answer from Pearson.  They were nice enough and were fairly easy to understand, although there was an delay echo introduced where whatever I said was echoed a quarter to half second later which was annoying, but bearable.  I was just happy she was able to hear me.  She said she could open a trouble ticket for me, but as it was well over an hour trying to get through to any human and doubtful it was on my side, I just told her to schedule me for the next available in person appointment.  She had to cancel my appointment and then rebook it as their sub-standard system wouldn't let her reschedule an at home appointment to at a location.  Surprisingly, she said they would refund my money and rebook me.  It was painless enough, but when I asked for a reference number on the refund, all she could do is say I 'should' get an email.  Perhaps unsurprisingly, this morning I see a fully posted charge for the rescheduled exam, but no sign of a refund.  Sigh.  I will give it a few days and then start this process over.

For what its worth, people should IGNORE the advice that the web chat is the fastest way of getting help.  Find the phone number and dial and re-dial it as fast as you can when you get a busy signal.  Despite the fact that it took 20+ minutes to get the number to pickup (and was 'waiting' 20 minutes less from the phones point of view) I got a faster response from someone on the phone.  Web based chat never picked up, even though I left it running during my entire phone conversation.

*Update #2*: It took two more days than the charge, but the refund did show up in the correct amount on my credit card.  I am actually quite surprised.
## [2][Advanced filtering in AWS CloudWatch](https://www.reddit.com/r/aws/comments/h08p73/advanced_filtering_in_aws_cloudwatch/)
- url: https://www.reddit.com/r/aws/comments/h08p73/advanced_filtering_in_aws_cloudwatch/
---
I am using JSON logs in CloudWatch.

I am trying to filter logs in CloudWatch but one of my fields is an array of objects. And I want to check a value of an object with a key "message".

For a sample log entry:

    { "appLog": {
        "payload": {
          "responses": [
            {
              "key": "message",
              "value": "Hello World!"
            }
          ]
        }
    }

I want filter all logs which has "message" key object in `responses`.

I am using this query works in Logs Insights and it works as expected. But If my `responses` array has multiple entries and the "message" key entry is not at the first index, my query fails.

    fields @message | filter appLog.payload.responses.0.key == 'message'

Is there any way of selecting an object with some property irrespective of the index? Something like:

    filter appLog.payload.responses.*.key == 'message'
## [3][AWS Fargate platform being upgraded from 1.3 to 1.4, has network impacting changes](https://www.reddit.com/r/aws/comments/gzy7l3/aws_fargate_platform_being_upgraded_from_13_to_14/)
- url: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html#platform-version-migration
---

## [4][Can I register multiple target groups with an ALB?](https://www.reddit.com/r/aws/comments/h064sz/can_i_register_multiple_target_groups_with_an_alb/)
- url: https://www.reddit.com/r/aws/comments/h064sz/can_i_register_multiple_target_groups_with_an_alb/
---
Hi, I have services running on different EC2 nodes, these services will be under different subdomains. [site1.example.com](https://site1.example.com), [site2.example.com](https://site2.example.com), [site2.example.com](https://site2.example.com) 

Can I point all of the DNS records at my load balancer, and then from there route to the different target groups?

Or is the best way to do this to put the records into route 53? (they currently are externally hosted)
## [5][Why are my instances in eu-west-3 geting a "US" (or default ?) answer from Route53 record with georouting ?](https://www.reddit.com/r/aws/comments/h0aizs/why_are_my_instances_in_euwest3_geting_a_us_or/)
- url: https://www.reddit.com/r/aws/comments/h0aizs/why_are_my_instances_in_euwest3_geting_a_us_or/
---
Hi there

I've a Route53 record which has a georouting policy attached. Default to us-east, if continent = Europe then eu-west.

Next to that, I just created a new production VPC / instances in eu-west-3. From the instances there, the georouting answer US, while it should answer EU. From others VPCs (staging, development in other EU zones), my own personnal server in a german datacenter, my laptop in france, the georouting works properly, getting "EU" as an answer

```bash
ᐅ ssh production-eu-west dig +short app.production.company.com
app-us-east.production.company.com.
123.123.123.123 # WRONG
ᐅ ssh production-us-east dig +short app.production.company.com
app-us-east.production.company.com.
123.123.123.123 # OK
ᐅ ssh staging dig +short app.production.company.com # eu-west-2
app-eu-west.production.company.com.
111.111.111.111 # OK
ᐅ ssh development dig +short app.production.company.com # eu-west-1
app-eu-west.production.company.com.
111.111.111.111 # OK
```

I guess there's a messup in the VPC, but it's terraformed using [this module](https://registry.terraform.io/modules/terraform-aws-modules/vpc/). I'm not "fluent" enought in AWS to dig properly in the routing tables, subnets, NAT gateway etc &amp; find an hint about that.

Can anyone help me finding the issue ?
## [6][Anyone switched from ECS with EC2 to Fargate for 24/7 workloads?](https://www.reddit.com/r/aws/comments/h09m21/anyone_switched_from_ecs_with_ec2_to_fargate_for/)
- url: https://www.reddit.com/r/aws/comments/h09m21/anyone_switched_from_ecs_with_ec2_to_fargate_for/
---
It seems with the recent price reduction, and the optional compute savings plans, Fargate has possibly become much more attractive for 24/7 workloads now.

Just curious about real world experience feedback such as cost, performance and general feedback.

Thinking about making the switch.
## [7][All files except SQL files deleted](https://www.reddit.com/r/aws/comments/h075f1/all_files_except_sql_files_deleted/)
- url: https://www.reddit.com/r/aws/comments/h075f1/all_files_except_sql_files_deleted/
---
I work for a small business in Argentina who hosts all of its files in an AWS server, for some reason the automatic backup did not work and the files were then deleted, the SQL data was the only thing that remained which was used to restore the lost files, do you have an idea why this happened and how to prevent from happening again in the future? Could it be a virus?
## [8][Is it possible to retrieve Tag metadata for EC2 instances that are scheduled for retirement from Cloudwatch Event Rules and send them to SNS for a Lambda function?](https://www.reddit.com/r/aws/comments/h05c06/is_it_possible_to_retrieve_tag_metadata_for_ec2/)
- url: https://www.reddit.com/r/aws/comments/h05c06/is_it_possible_to_retrieve_tag_metadata_for_ec2/
---
 I'm trying to trigger a lambda with an instance retirement Cloudwatch Event \["AWS\_EC2\_INSTANCE\_RETIREMENT\_SCHEDULED"\]

Does anyone know how I could simulate an instance retirement in my account and filter the instance id resource using TAGS? I believe you can only specify all ec2 resources or by specific (instance id only) .  I manage thousands of EC2s and everything is handled by ASG so specific instance id is not an option. I only want to receive AWS\_EC2\_INSTANCE\_RETIREMENT\_SCHEDULED notifications on a specific group of instances possibly using TAGS. Please help.
## [9][cookies with dev/stage/prod in same region](https://www.reddit.com/r/aws/comments/h0awax/cookies_with_devstageprod_in_same_region/)
- url: https://www.reddit.com/r/aws/comments/h0awax/cookies_with_devstageprod_in_same_region/
---
Hi,

we have the following three environments in the same region, all with sticky sessions.

dev, stage, and prod.

It seems when I go to dev, go to stage, and go to certain urls, and return to dev, I need to clear my cookies for it to work, just wondering if anyone else has experienced this?

thanks.
## [10][The Path Towards Enterprise Level AWS Infrastructure series by Michał Kapiczyński](https://www.reddit.com/r/aws/comments/h09xox/the_path_towards_enterprise_level_aws/)
- url: https://www.reddit.com/r/aws/comments/h09xox/the_path_towards_enterprise_level_aws/
---
A mini-series which will walk you through the process of creating an enterprise-level AWS infrastructure. By the end of this series, we will have created an infrastructure comprising a VPC with four subnets in two different availability zones with a client application, backend server and a database deployed inside. 

Series:

* [Architecture Scaffolding](https://itnext.io/the-path-towards-enterprise-level-aws-infrastructure-architecture-scaffolding-d244d0c80364?source=friends_link&amp;sk=fe6e16d39aba934189a148dcdd51ccab)
* [EC2, AMI, Bastion Host, RDS](https://itnext.io/the-path-towards-enterprise-level-aws-infrastructure-part-2-ec2-ami-bastion-host-rds-3109c73dc913?source=friends_link&amp;sk=8d4dd8a3f6a30621c5d660e8632549f8)
* [Load Balancing and Application Deployment](https://itnext.io/the-path-towards-enterprise-level-aws-infrastructure-load-balancing-and-application-deployment-47c48e4c343d?source=friends_link&amp;sk=e59789ba257f74696acf8f17e2e87814)
## [11][Golf Factor](https://www.reddit.com/r/aws/comments/gzn0fv/golf_factor/)
- url: https://www.reddit.com/r/aws/comments/gzn0fv/golf_factor/
---
Unfortunately this community has no shitpost flair, as I would mark it (as it kinda is a shitpost).

I'm a fan of cloud solutions and abstracting the data center, but in many organizations the cost outweights the benefits of migrating. I've seen migrations that quadrupled the cost of ownership for infrastructure with no clear benefit. And the only viable explanation for such a move would be something I've called the "Golf Factor".

So what is Golf Factor? Pretty simple, it's the ability for CEO/CTO, to say that they too use "The Cloud" during meetings with other C*Os. 

Have you guys noticed such behaviour? Has it influenced the decision to move to cloud, if yes how?
