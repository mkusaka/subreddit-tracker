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
## [2][What is the proper way to build many Lambda functions and updated them later?](https://www.reddit.com/r/aws/comments/htfa9z/what_is_the_proper_way_to_build_many_lambda/)
- url: https://www.reddit.com/r/aws/comments/htfa9z/what_is_the_proper_way_to_build_many_lambda/
---
I want to make a bot that makes other bots on Telegram platform. I want to use AWS infrastructure, look like their Lamdba functions are perfect fit, pay for them only when they are active. In my concept, each bot equal to one lambda function, and they all share the same codebase.

At the starting point, I thought to make each new Lambda function programmatically, but this will bring me problems later I think, like need to attach many services programmatically via AWS SDK: Gateway API, DynamoDB. But the main problem, how I will update the codebase for these 1000+ functions later? I think that bash script is a bad idea here.

So, I moved forward and found SAM (AWS Serverless Application Model) and CloudFormatting, which should help me I guess. But I can't understand the concept. I can make a stack with all the required resources, but how will I make new bots from this one stack? Or should I build a template and make new stacks for each new bot programmatically via AWS SDK from this template?

Next, how to update them later? For example, I want to update all bots that have version 1.1 to version 1.2. How I will replace them? Should I make a new stack or can I update older ones? I don't see any options in UI of CloudFormatting or any related methods in AWS SDK for that.

Thanks
## [3][Do I need to rewrite my app to use Aurora?](https://www.reddit.com/r/aws/comments/htdj3m/do_i_need_to_rewrite_my_app_to_use_aurora/)
- url: https://www.reddit.com/r/aws/comments/htdj3m/do_i_need_to_rewrite_my_app_to_use_aurora/
---
I see that Aurora has a write endpoint and a read endpoint. Does. That mean to take advantage of read endpoints I need to specify a different endpoint for read queries?

i.e go through my existing application, figure out the reads, and then change the endpoints for those queries?

I tried searching online, but couldn’t get a definitive answer. Most apps I know have only one IP/DNS specified for the DB, making no difference between reads and writes.
## [4][Approzium: Observable, password-less authentication to databases - built on AWS IAM](https://www.reddit.com/r/aws/comments/ht2bnj/approzium_observable_passwordless_authentication/)
- url: https://github.com/cyralinc/approzium
---

## [5][ECS - our server response time has dropped from 0.3s to 2.5s](https://www.reddit.com/r/aws/comments/htgbnj/ecs_our_server_response_time_has_dropped_from_03s/)
- url: https://www.reddit.com/r/aws/comments/htgbnj/ecs_our_server_response_time_has_dropped_from_03s/
---
I've been updating a legacy PHP app (no version control for 10 years) and I've gotten it working pretty nicely on AWS now. I have some problems I can't really fix.

1. CPU usage for the ECS service is always above 130%. I don't understand why as the CPU for the EC2 box is only 8%, docker process says the same. This isn't an intensive site, it's just some really old PHP code.
2. We have a response time of 2.5s instead of 0.3s. In Google lighthouse this is indicated by \`Reduce server response times (TTFB)**\`.** The apache server setup is the same, and the code running the site is the same. Only difference is my code runs on ECS instances, and the old code runs directly on an IP exposed EC2 box.

Our setup is roughly this:

Application Load Balancer

2 target groups, HTTPS and HTTP.

HTTP does a 301 redirect to out HTTPS group. (I set this up as the site kept defaulting to HTTP - is this normal?)

At the moment we have 1 cluster, 1 service and 1 task running on ECS using EC2.

Our EC2 box is dedicated, t2 medium.

Our files are on EFS. Here we store all of our cache files, image files and session files so they are shared.

We have a certificate issued by Route53 and the site validates fine.

Docker is running Apache *20051115*, the site is on PHP5.4 and the database is MySQL 5.5.

Does anyone have any idea what could be happening? Thanks!
## [6][Completely new to AWS and lost on what to do next](https://www.reddit.com/r/aws/comments/ht7xfb/completely_new_to_aws_and_lost_on_what_to_do_next/)
- url: https://www.reddit.com/r/aws/comments/ht7xfb/completely_new_to_aws_and_lost_on_what_to_do_next/
---
I'm trying to learn aws by building an infrastructure for a website. I built an vpc, made ec2 instances, made a elastic load balencer with an auto scaling group , an rds database and connected it to one of my instances, an efs to share a file system between my ec2 instances, and even transferred my ec2 logs to cloudwatchlogs  all with the help of youtube. The thing that I am completely stumped on is how to transfer my my cloudwatchlogs to my s3 bucket. Thank you for taking time to read my post and I am truly sorry if this sounds like gibberish, I am under the weather and frustrated. Thanks again
## [7][Cloud formation or Terraform](https://www.reddit.com/r/aws/comments/ht3oec/cloud_formation_or_terraform/)
- url: https://www.reddit.com/r/aws/comments/ht3oec/cloud_formation_or_terraform/
---
Hi. 
Just starting a Greenfield SaaS product. Could go with Terraform or Cloud formation. SaaS will be mainly lamdas, rds. 
Which would you go for?
## [8][What is a notebook instance type?](https://www.reddit.com/r/aws/comments/htcew9/what_is_a_notebook_instance_type/)
- url: https://www.reddit.com/r/aws/comments/htcew9/what_is_a_notebook_instance_type/
---
When you create a new notebook instance you have to specify an instance type, choosing which between different size(?). But what do this sizes mean. Are they an allocation of computer capacity?
## [9][Configure ALB health check based on JSON response](https://www.reddit.com/r/aws/comments/htc1ih/configure_alb_health_check_based_on_json_response/)
- url: https://www.reddit.com/r/aws/comments/htc1ih/configure_alb_health_check_based_on_json_response/
---
Backend servers behind ALB returns JSON response `curl http://backend_srv1/healthapi`

    "health" : {
      "status" : "success"
    },

Is there a way to configure ALB health check based on success/fail JSON response?
## [10][Question on aws s3 cp](https://www.reddit.com/r/aws/comments/hszkea/question_on_aws_s3_cp/)
- url: https://www.reddit.com/r/aws/comments/hszkea/question_on_aws_s3_cp/
---
If the source and destination are s3 buckets, does the file ever leave AWS, or is the copy direct from s3 bucket to s3 bucket?

Example:

    aws s3 cp s3://foo/bar.tgz s3://baz/
## [11][CloudFormation Designer - How do you know what you need to define?](https://www.reddit.com/r/aws/comments/ht7kvx/cloudformation_designer_how_do_you_know_what_you/)
- url: https://www.reddit.com/r/aws/comments/ht7kvx/cloudformation_designer_how_do_you_know_what_you/
---
I am trying to build a stack that runs an ECS cluster for an application that acts as a server to listen for HTTP requests.

With that simple premise, how am I supposed to throw together a CF design? Is there any way to validate it as you go? When I open the designer it's pretty intimidating to see literally a blank slate.

For example, I added an ECS cluster to the screen. Now what? This is not very intuitive - any tips are appreciated.
