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
## [2][AWS Landing zone 2.4.0](https://www.reddit.com/r/aws/comments/gtrp6p/aws_landing_zone_240/)
- url: https://github.com/clouddrove/cloudformation-aws-landing-zone
---

## [3][T3 unlimited pricing question](https://www.reddit.com/r/aws/comments/gttoto/t3_unlimited_pricing_question/)
- url: https://www.reddit.com/r/aws/comments/gttoto/t3_unlimited_pricing_question/
---
I'm analyzing the cost of moving a cluster of servers from Rackspace to AWS.  Doing a simple test I've concluded that this cluster will burn 8 CPU credits per hour per server pretty much every hour 24x7.  This is a very CPU consistent usecase by design.  If I use a t3.nano I'm burning 2 CPU credits per hour more than I'm allowed.  If I go with a t3.micro I'm accumulating 4 CPU credits per hour more than I'm using.  Can someone tell me what the cost difference is by going with t3.nano vs t3.micro using t3 unlimited?  I assume the nano would actually be more expensive?
## [4][SELECT * FROM S3 - Using S3 Select to Query S3 files with SQL](https://www.reddit.com/r/aws/comments/gtf1hw/select_from_s3_using_s3_select_to_query_s3_files/)
- url: https://www.reddit.com/r/aws/comments/gtf1hw/select_from_s3_using_s3_select_to_query_s3_files/
---
Hey folks,

Recently put together a tutorial video for using AWS' newish feature, S3 Select, to run SQL commands on your JSON, CSV, or Parquet files in S3.

Ontop of it being super easy to use, using S3 Select over traditional S3 Get + Filtering has a 400% performance improvement + cost reduction. Very useful for those of you pulling down large config files in your apps.

The tutorial video is located here: https://youtu.be/yqJwN8EBCw8

Thanks!
## [5][Setup Question: PHP Deployment](https://www.reddit.com/r/aws/comments/gtrr9g/setup_question_php_deployment/)
- url: https://www.reddit.com/r/aws/comments/gtrr9g/setup_question_php_deployment/
---
I have a question about the best way to do PHP production deployments. Right now, I push to a production github enterprise branch, increase the number of nodes in my autoscaling group, let the new node be built, and then delete the old node so we're running on the new version. If I can have downtime, I'll just blow the server away and let it rebuild. 

I'd love to have the application deploy as soon as a push it to github enterprise with no manual action. I prefer to have the server rebuilt because it avoids oddities with dependencies and things like that to just let everything pull and rebuild. 

Any recommendations about how to do this better?
## [6][Point elastic beanstalk to subdomain [ Beginner ]](https://www.reddit.com/r/aws/comments/gtw00a/point_elastic_beanstalk_to_subdomain_beginner/)
- url: https://www.reddit.com/r/aws/comments/gtw00a/point_elastic_beanstalk_to_subdomain_beginner/
---
Hello everyone beginner here, I have setup elastic beanstalk now i want to point it to my subdomain (namecheap)
Given im using the free tier can anyone here help give some ideas how to do it.
## [7][Can I serve .php file from S3 hosting? If not, what to use?](https://www.reddit.com/r/aws/comments/gtyr8j/can_i_serve_php_file_from_s3_hosting_if_not_what/)
- url: https://www.reddit.com/r/aws/comments/gtyr8j/can_i_serve_php_file_from_s3_hosting_if_not_what/
---
Hi, is it possible to serve a index.php file from S3 hosting? I think it's not but I'd rather know for sure.

If it's not possible, what should I use to do that? I'm very new to hosting and servers so I'd be very grateful to get some advice and recommendations.

What I'm trying to do is create a gallery but instead of hard coding tens of &lt;img&gt; tags, I want to get file names from S3 where they are stored and run a script that creates &lt;img&gt; tags for me. I tested locally and I get the local file names so everything should work, I just need help solving the index.php file serving instead of index.html.
## [8][Question: SQL File for Local Dev Environment](https://www.reddit.com/r/aws/comments/gtros0/question_sql_file_for_local_dev_environment/)
- url: https://www.reddit.com/r/aws/comments/gtros0/question_sql_file_for_local_dev_environment/
---
Looking for some advice. I have a local setup for PHP development that I'll pull and load with a copy of my production data every so often. I also have a database that we take to load locally into R for offline analysis. 

Right now, I have a shell script that runs every six hours that takes a database and EFS backup and put it in S3 for backup purposes and a copy in an S3 bucket for replicating to dev/analysis environments.

I want to move to AWS backup and go serverless for my backups. However, i'm wondering the best way to get the copy of the RDS and EFS data into S3 in a serverless wway where it can be downloaded. Any recommendations?
## [9][SAM and CORS support](https://www.reddit.com/r/aws/comments/gtn45w/sam_and_cors_support/)
- url: https://www.reddit.com/r/aws/comments/gtn45w/sam_and_cors_support/
---
Can someone please help me with CORS, since there's no simple and wholesome explanation how to do it using SAM. My static website is served from S3 bucket and my api is built using SAM (api gateway, lambda, dynamodb, rds(+vpc lambda)). I understand that the source (my api) needs to serve allowed origins in http header. Now, the question is where and how? In template I gave my api gateway cors: allowed-origin..: "httpmys3link.com" but it failed because  option method is required for every function. What do I do now? Thanks.
## [10][Unable to connect (including SSH) to EC2 instance randomly?](https://www.reddit.com/r/aws/comments/gtjn1f/unable_to_connect_including_ssh_to_ec2_instance/)
- url: https://www.reddit.com/r/aws/comments/gtjn1f/unable_to_connect_including_ssh_to_ec2_instance/
---
Occasionally, there would be a window of time where I would be unable to connect (including SSH and HTTP) to my EC2 instance. The instance type is a T2 micro, and I've had it for ~3 years.

Status checks seem OK (2/2), and there's some CPU activity on CloudWatch. I attempted the provided diagnostics but they always fail when this happens. I should have EBS storage on it, one of my suspicions is lack of RAM.

Any ideas for debugging? It's very frustrating to debug. Thanks!

Thanks in advance :)

EDIT: Forgot to add that I reboot the instance and the inability to connect remains.
## [11][What do your "day to day" roles look like?](https://www.reddit.com/r/aws/comments/gtk8ou/what_do_your_day_to_day_roles_look_like/)
- url: https://www.reddit.com/r/aws/comments/gtk8ou/what_do_your_day_to_day_roles_look_like/
---
For your developers, what permission do you give them on the roles they assume? I'm trying to figure out a good middle ground without giving them everything, while not impeding them while doing their job.
