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
## [2][A Shared File System for Your Lambda Functions](https://www.reddit.com/r/aws/comments/habrhi/a_shared_file_system_for_your_lambda_functions/)
- url: https://aws.amazon.com/blogs/aws/new-a-shared-file-system-for-your-lambda-functions/
---

## [3][Amazon EC2 Auto Scaling now supports Instance Refresh within Auto Scaling Groups](https://www.reddit.com/r/aws/comments/hafi7q/amazon_ec2_auto_scaling_now_supports_instance/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/06/amazon-ec2-auto-scaling-now-supports-instance-refresh-within-auto-scaling-groups/
---

## [4][We are thinking of migrating from Classic ECS to Fargate - costs?](https://www.reddit.com/r/aws/comments/haop4g/we_are_thinking_of_migrating_from_classic_ecs_to/)
- url: https://www.reddit.com/r/aws/comments/haop4g/we_are_thinking_of_migrating_from_classic_ecs_to/
---
I am thinking of moving a 'classical' ECS docker cluster running 67 services to Fargate. The cluster autoscales in/out but has about 12 nodes. It costs us about $960/month.

I ran some calculations on the running the same tasks in Fargate and calculated that 67 Fargate tasks with 0.5 CPU and 1GB of memory would cost $424/month. I'm a little surprised by this since most web forums claim Fargate is MORE expensive.

I also wonder whether the 0.5CPU in Fargate is equivalent to 200 CPU Units in ECS if the instance type is C5.large.

Is there anything else I might be missing?
## [5][AWS Workspaces](https://www.reddit.com/r/aws/comments/happs5/aws_workspaces/)
- url: https://www.reddit.com/r/aws/comments/happs5/aws_workspaces/
---
Why the fuck hasn't a solution been created for the problem of  "User must change password at next login" not been created yet? How is everyone doing on-boarding?
## [6][Reminder: AWS Summit Online is starting today!](https://www.reddit.com/r/aws/comments/han0w1/reminder_aws_summit_online_is_starting_today/)
- url: https://www.reddit.com/r/aws/comments/han0w1/reminder_aws_summit_online_is_starting_today/
---
See more info and how to register [here](https://aws.amazon.com/events/summits/online/emea/agenda/)

Can't wait!

EDIT: looks like they should have used scheduled auto scaling instead of target tracking:) site seems to be up and working now...

Update: it seems like their whole schedule has shifted mid-day by an hour. I'm utterly confused at what's going on everywhere. It's a prerecorded summit, how hard is it really? you were supposed to be the chosen one AWS....  
Do what I say, not what I do much?  

## [7][What would be the best way to build Gravatar like service with using AWS services and least amount of code?](https://www.reddit.com/r/aws/comments/har7l0/what_would_be_the_best_way_to_build_gravatar_like/)
- url: https://www.reddit.com/r/aws/comments/har7l0/what_would_be_the_best_way_to_build_gravatar_like/
---
How could you build a system with AWS services that:
1. A user can upload an image, and the only detail they give out is their email.
2. The uploaded image can be accessed by some hash (generated from email?) in public.
3. The uploaded image GET must also have support for width and height parameters that allows to fetch custom size uploaded image. If no width and height is specified, a default 250x250 image is returned.

How would build it with AWS services with using least amount of code?
## [8][SNS sends two emails for RDS snapshot creation](https://www.reddit.com/r/aws/comments/haquel/sns_sends_two_emails_for_rds_snapshot_creation/)
- url: https://www.reddit.com/r/aws/comments/haquel/sns_sends_two_emails_for_rds_snapshot_creation/
---
I'm setting up SNS to send out emails confirming the snapshots went off and completed successfully. Right now cloudwatch has an event set up for RDS DB Snapshot Events but I get two emails.

One is "creating autmoated snapshot" and the second 2 minutes later is "Automated snapshot created" telling me that it completed.  
  
What JSON should I have in cloudwatch to only receive the second one? Here is what was input by amazon by default  
  
    {
      "source": [
        "aws.rds"
      ],
      "detail-type": [
        "RDS DB Snapshot Event"
      ]
    }
## [9][Appstream: Files question](https://www.reddit.com/r/aws/comments/haqsm4/appstream_files_question/)
- url: https://www.reddit.com/r/aws/comments/haqsm4/appstream_files_question/
---
In our use case we have a program where we need to at times update the files folder, doing this through image builder is time consuming - is there an easier way (perhaps mapping a S3 drive at run time?) to update these files without updating the image?.
## [10][How to store data/file-uploads securely where a visitor cant guess URLs for file uploads using S3?](https://www.reddit.com/r/aws/comments/hamazy/how_to_store_datafileuploads_securely_where_a/)
- url: https://www.reddit.com/r/aws/comments/hamazy/how_to_store_datafileuploads_securely_where_a/
---
* Using Django and AWS S3 for file uploads
* I dont want visitors to be able to guess eg. [www.whatever.com/user/file/2332fdsfSDFsdfsfd234.jpg](http://www.whatever.com/user/file/2332fdsfSDFsdfsfd234.jpg) for another users. All data should be private. A user will upload a private document or a photo of themselves and it should only be viewable by the logged in user or the admin of the web app.
* What is the term/method I should look up to handle this situation?
## [11][Trouble with AWS Glue and adding column names through a custom classifier](https://www.reddit.com/r/aws/comments/haptk2/trouble_with_aws_glue_and_adding_column_names/)
- url: https://www.reddit.com/r/aws/comments/haptk2/trouble_with_aws_glue_and_adding_column_names/
---
Hey guys,

important: we are deploying everything with terraform but a manual workaround for this would also be sufficient!

For a client we are required to read a .tsv file that is located in a S3. This file does NOT have any headers and currently if we run a crawler to populate a database with a table in Glue, we get the col1, col2 etc.. Not desired behaviour so we got the requirement to add column names.

We tried adding a custom classifier that adds column names to the table but this does not seem to work. I am wondering if anyone has experience here with Glue and how to solve this. If you guys require more info (scrubbed) let me know and I can see what to provide!  


Cheers, ConsultantCat
