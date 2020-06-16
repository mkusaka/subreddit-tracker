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
## [2][Dealing with growing load on RDS](https://www.reddit.com/r/aws/comments/ha1bn2/dealing_with_growing_load_on_rds/)
- url: https://www.reddit.com/r/aws/comments/ha1bn2/dealing_with_growing_load_on_rds/
---
Hi everyone!

I have a hypothetical scenario, maybe you are able to help me out.

Say you have a increasingly stressed production DB in RDS, 3TB of data with 500GB increase each month. You want to offload data for analytics purposes and reduce the stress on your production DB, making scaling in the future easier.

So I know that by using ETL tools like Glue, EMR, LakeFormation you can load your data into e.g. a data lake (or even Redshift) for further analysis.

But then the scaling problem in RDS is still an issue. How would you tackle an increasingly stressed main DB? By rethinking the entire persistence layer architecture, splitting up a databases into several databases? Move to something more scalable like Aurora or even DynamoDB?

Thanks! Really appreciate the help!
## [3][Cluster Cloner: Copying cluster infrastructure between GKE, EKS, and AKS](https://www.reddit.com/r/aws/comments/ha2tou/cluster_cloner_copying_cluster_infrastructure/)
- url: https://blog.doit-intl.com/you-can-handle-the-pods-but-what-about-the-clusters-486fbdb5345d
---

## [4][Enforcing Compliance using the AWS Ansible Collection](https://www.reddit.com/r/aws/comments/ha1gg7/enforcing_compliance_using_the_aws_ansible/)
- url: https://steampunk.si/blog/enforcing-compliance-using-aws-ansible-collection/
---

## [5][Don't forget that you can ask Jeff Barr anything (right here on r/aws) for the upcoming AWS Community Online event on 6 July](https://www.reddit.com/r/aws/comments/h9i8op/dont_forget_that_you_can_ask_jeff_barr_anything/)
- url: https://www.reddit.com/r/aws/comments/h15ie8/ask_jeff_anything_help_inform_jeff_barrs_keynote/
---

## [6][Do you feel that AWS needs to step up its game in following the Java release cadence?](https://www.reddit.com/r/aws/comments/ha3p2z/do_you_feel_that_aws_needs_to_step_up_its_game_in/)
- url: https://www.reddit.com/r/aws/comments/ha3p2z/do_you_feel_that_aws_needs_to_step_up_its_game_in/
---
The latest release (not LTS though) is   **Java** 14  and aws java is at 11. I think its not good enough. I wish AWS would step up its game and work a little bit faster at keeping up.

&amp;#x200B;

What do you think?
## [7][AWS XRay: Out of the box spring](https://www.reddit.com/r/aws/comments/ha3jyc/aws_xray_out_of_the_box_spring/)
- url: https://www.reddit.com/r/aws/comments/ha3jyc/aws_xray_out_of_the_box_spring/
---
https://docs.aws.amazon.com/xray/latest/devguide/xray-sdk-java-aop-spring.html

This guide shows how we can use AWS XRay out of the box for spring projects without making any major code changes. However, if we follow these to the dot, there is no auto instrumentation of code. Is there any extra configuration I have to add?
## [8][Is AWS Batch an analogous product to Cisco Tidal or Stonebranch?](https://www.reddit.com/r/aws/comments/ha00ir/is_aws_batch_an_analogous_product_to_cisco_tidal/)
- url: https://www.reddit.com/r/aws/comments/ha00ir/is_aws_batch_an_analogous_product_to_cisco_tidal/
---
I'm trying to wrap my head around the use cases for AWS Batch. From what I've gathered it's best suited for batch processing tasks, some of which can be run in parallel and some which have dependencies on other jobs. 

Currently at my work we schedule jobs on Cisco Tidal. We have several servers available to us locally that we can run jobs on. We define jobs in Tidal to either run at a specific time or after some job dependencies have been met. Jobs are allocated to run on our local servers depending on which servers have suitable resources available.

Is this essentially what AWS Batch does, just on AWS infrastructure?
## [9][Mounting a volume for AWS Batch](https://www.reddit.com/r/aws/comments/ha1yiw/mounting_a_volume_for_aws_batch/)
- url: https://www.reddit.com/r/aws/comments/ha1yiw/mounting_a_volume_for_aws_batch/
---
Hi,

We have some AWS batch processes that run nicely, using images from ECS.  We do not assign any volumes or storage, and it seems we get 8gb by default. I'm not actually sure why/where that is defined.

Anyway we now have a situation where we need more space. It's only temporary processing space - we need to extract an archive, convert it, re-compress it and then upload it to S3. We already have this process, it's just that we've now ran out of space in our 8gb allowance.

So; Just to be absolutely sure, how should we go about adding this space?  I see a few things about connecting EFS to the instance, is that a good use case?  Are there considerations regarding to multiple jobs running at the same time etc?  (There are scenarios where this is allowed - since it's a generic unzipper process, that gets called many times). 

So the requirement is a throwaway storage volume, that doesn't need to persist, it can disappear once the AWS batch job finishes.  The data files that have currently blown it up are 9gb, i'm not sure how much our image itself uses.  Alpine linux so presumably not a huge amount.

Or of course, if we can simply tune that initial 8gb up by a couple of gb then we're laughing...

Thanks!

Dan
## [10][AWS Cognito: after a successful login form AWS kibana using Google Account and loginout. My second login is unsuccessful and loop me back to the login page.](https://www.reddit.com/r/aws/comments/h9wmfh/aws_cognito_after_a_successful_login_form_aws/)
- url: https://www.reddit.com/r/aws/comments/h9wmfh/aws_cognito_after_a_successful_login_form_aws/
---
AWS Cognito: after a successful login form AWS kibana using Google Account and loginout. My second login is unsuccessful and loop me back to the login page. How do i fix this Error?
## [11][SSL and cloudfront with a domain registered from somewhere other than route 53](https://www.reddit.com/r/aws/comments/h9snsr/ssl_and_cloudfront_with_a_domain_registered_from/)
- url: https://www.reddit.com/r/aws/comments/h9snsr/ssl_and_cloudfront_with_a_domain_registered_from/
---
Hi,

I have 2 domain names bought from GoDaddy and I've been doing a lot of research and it seems like I can't get SSL working with cloudfront with domains from GoDaddy because GoDaddy doesn't offer SSL  domain forwarding? I was wondering if anyone else had this issue or if they transferred their domains to another provider?

resource I found: [https://stackoverflow.com/questions/56914945/https-connections-to-cloudfront-s3-using-godaddy-domain](https://stackoverflow.com/questions/56914945/https-connections-to-cloudfront-s3-using-godaddy-domain)
