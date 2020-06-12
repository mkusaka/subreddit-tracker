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
## [2][Show /r/aws: periodic sanitized snapshot of production RDS database for developers/QA](https://www.reddit.com/r/aws/comments/h7ew99/show_raws_periodic_sanitized_snapshot_of/)
- url: https://github.com/CloudSnorkel/RDS-sanitized-snapshots
---

## [3][How to make signed URLs less unwieldy and more usable?](https://www.reddit.com/r/aws/comments/h7dt02/how_to_make_signed_urls_less_unwieldy_and_more/)
- url: https://www.reddit.com/r/aws/comments/h7dt02/how_to_make_signed_urls_less_unwieldy_and_more/
---
Since signed URLs are so long (~400 extra chars) are there any good solutions for sharing shorter URLs? 
Do people build their own URL shorteners to get deal with this? Or are there simpler workarounds?
## [4][AWS Anti Patterns - Mixing Accounts](https://www.reddit.com/r/aws/comments/h7jy11/aws_anti_patterns_mixing_accounts/)
- url: https://www.reddit.com/r/aws/comments/h7jy11/aws_anti_patterns_mixing_accounts/
---
Guilty of this myself and seeing many others fall into the same trap, I wanted to make a quick technical resource to highlight the problems with mixing accounts and why its important to avoid this trap from the beginning.

Video is available here: https://youtu.be/A_XnXc-5i8Y
## [5][Fargate task on private subnet getting "CannotPullContainerError" even though its got a working NAT](https://www.reddit.com/r/aws/comments/h7iqme/fargate_task_on_private_subnet_getting/)
- url: https://www.reddit.com/r/aws/comments/h7iqme/fargate_task_on_private_subnet_getting/
---
I have a Fargate task failing to deploy with a CannotPullContainerError that suggests it cannot reach the internet. Its on a private subnet that has a working NAT and other EC2 instances that have no problem reaching the outside world.  I have IAM permissions setup, and everything works fine inside a public subnet. According to the docs the error message Im getting suggests there is no path to the repository. Im using ECR.

I cant figure out what is going on, does anyone have any suggestions? Thanks!
## [6][How to restart EC2 instance in auto scaling group?](https://www.reddit.com/r/aws/comments/h7gbu2/how_to_restart_ec2_instance_in_auto_scaling_group/)
- url: https://www.reddit.com/r/aws/comments/h7gbu2/how_to_restart_ec2_instance_in_auto_scaling_group/
---
I have an EC2 auto scaling group configured to run from 1 to 5 instance depending on load. I need the instances to restart for some configuration changes to happen. 4 out of 5 instances have already been restarted by auto scaling group. But 1 is always on, I want to restart that instance. I want to know the following:

1. If I restart that instance manually when another is running (so 2 out of 5 are running ) will AWS understand that and run another instance?
2. If 1 is not the correct way of restarting instances in EC2 auto scaling group, the what is?
## [7][Have codedeploy launch a new instance before update](https://www.reddit.com/r/aws/comments/h7g9cw/have_codedeploy_launch_a_new_instance_before/)
- url: https://www.reddit.com/r/aws/comments/h7g9cw/have_codedeploy_launch_a_new_instance_before/
---
Hello
Is there a way I can have no downtime deployments with CodeDeploy without the Blue/Green option?

I have a single instance behind a load balancer. I would like to have CodeDeploy launch a new instance before deployment so that I have no downtime with a OneAtATime deploy config. It should then turn off the old instance once complete.

I launched my CF stack including a ASG and for some reason CD will create a new ASG when doing a Blue/Green which then deletes the old ASG. Then all be CF updates fail because the ASG resource no longer exists.

Does anybody have a workaround?
## [8][Show /r/aws: A Better AWS Console](https://www.reddit.com/r/aws/comments/h109wr/show_raws_a_better_aws_console/)
- url: https://www.reddit.com/r/aws/comments/h109wr/show_raws_a_better_aws_console/
---
Hi folks,

I saw the post about how rough the AWS Console is and I couldn't agree more. A few of my friends are working on https://vantage.sh/ - a console experience for AWS focused on better UI/UX. I'm using it in early access right now and it is extremely well done and getting better every week. 

I know they're actively looking for early access user feedback so if you're interested you should sign up. I also messaged them about posting here so hopefully they have/will make a reddit account and come in here to chat.
## [9][AWS Backup - cron command](https://www.reddit.com/r/aws/comments/h79yin/aws_backup_cron_command/)
- url: https://www.reddit.com/r/aws/comments/h79yin/aws_backup_cron_command/
---
Hi all

Currently configuring AWS Backup and I would like to use a cron command to take care of quarterly backups

It seems that some sites online have slightly different syntax or less values that AWS doesn't accept

I was wondering if anyone had an example that they use in their configuration?
## [10][Advice on AWS services to use for photo-sharing mobile app](https://www.reddit.com/r/aws/comments/h7enuc/advice_on_aws_services_to_use_for_photosharing/)
- url: https://www.reddit.com/r/aws/comments/h7enuc/advice_on_aws_services_to_use_for_photosharing/
---
Hi AWS experts - 

I'm a software engineer but very new to cloud computing.

I'm building a photo-sharing mobile app that I hope will eventually attract high traffic and be spun into a start-up. I have about 20% of the backend and proof-of-concept ready and running on my local workstation. The overall system will consist of a Java Spring web server, a Python machine learning engine, a website and iOS and Android mobile apps. AWS is so vast and confusing I'm at a loss what infrastructure to use. Could any of you kindly give me options to look at, pros, cons, etc. of the various AWS services I could use to achieve a system like this? Could you also go into things like monitoring? Not just infrastructure monitoring, but how I can write metrics from my code, etc. If you've launched a similar app from AWS could you also share your experience?

Thanks very much!
## [11][ECS isn't using the Codedeploy images file on deploy](https://www.reddit.com/r/aws/comments/h13zfe/ecs_isnt_using_the_codedeploy_images_file_on/)
- url: https://www.reddit.com/r/aws/comments/h13zfe/ecs_isnt_using_the_codedeploy_images_file_on/
---
I am using Terraform for a Codepipeline deployed ECS cluster. The codepipeline works up until the deploy stage. I have been digging into this issue for about 3 days now and what I have found, is that when I review the codepipeline execution, I do find the proper imagedefinitions.json file in S3 artifact button. The formatting looks right to me (see below), but when I look at the tasks that failed to start in ECS, I find that it is looking for the wrong image in ECR. See image attached. It is looking for the one from the original task definition (:latest), not the one in the imagedefinitions.json (:320ac83).

[ECS Screenshot](https://atlchris-screenshots.s3.amazonaws.com/Screen+Shot+2020-06-10+at+7.38.01+PM.png)

Does anyone have any theories on what might be going on or things I can look at further to help narrow down the issue?

    # imagedefinitions.json
     
    [
       {
          "name":"php",
          "imageUri":"xxxxxxxx07.dkr.ecr.us-east-1.amazonaws.com/filmfed-php:320ac83"
       },
       {
          "name":"nginx",
          "imageUri":"xxxxxxxx07.dkr.ecr.us-east-1.amazonaws.com/filmfed-nginx:320ac83"
       }
    ]
