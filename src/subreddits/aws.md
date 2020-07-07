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
## [2][PSA: Anyone Just Getting Into AWS, Set Budget Alerts and Remove Unused Databases / Instances](https://www.reddit.com/r/aws/comments/hmtmtw/psa_anyone_just_getting_into_aws_set_budget/)
- url: https://i.redd.it/ui2es8bgif951.jpg
---

## [3][Best practices for allowing an outside developer access to my AWS instance?](https://www.reddit.com/r/aws/comments/hmtyc1/best_practices_for_allowing_an_outside_developer/)
- url: https://www.reddit.com/r/aws/comments/hmtyc1/best_practices_for_allowing_an_outside_developer/
---
We have an outside developer that we have hired to build a piece of enterprise software for us.  They are going to be using an S3 bucket (creating it and filling it) as part of the development process.  Post-development, the application will obviously need access to the bucket, but the developer will not.  

What are the best practices for giving the outside guy enough permissions to develop and test the app, but not so much they can do damage within my AWS instance.  I also want to be able to revoke their access once the development process is complete and add them back for future development and feature add-ons.  

Thoughts?  Anything I should definitely NOT do in this process?

Thanks in advance for your comments...
## [4][Is it possible for an S3 to trigger a Lambda but the Lambda be too fast to get the actual file the first run?](https://www.reddit.com/r/aws/comments/hmt15e/is_it_possible_for_an_s3_to_trigger_a_lambda_but/)
- url: https://www.reddit.com/r/aws/comments/hmt15e/is_it_possible_for_an_s3_to_trigger_a_lambda_but/
---
after a frustratingly long several hours of investigation, i’ve come to the conclusion that it is not wholly uncommon for an s3 triggered event to fail on the first attempt to do something with it, then try again later.

eg. i write 100 59 MB files to an S3 prefix.  a lambda executes for each to process it.  99 work as expected.  1 fails.  not because the code logic is bad, or there was a timeout, actually the exception i catch is that it *fails to retrieve the s3 object that triggered it* with a 404.   it then ran again (automatically), about a minute later, successfully.  is this normal expected behavior?
## [5][How does Amazon Linux relate to CentOS and RHEL?](https://www.reddit.com/r/aws/comments/hmnrb6/how_does_amazon_linux_relate_to_centos_and_rhel/)
- url: https://www.reddit.com/r/aws/comments/hmnrb6/how_does_amazon_linux_relate_to_centos_and_rhel/
---
My understanding of Amazon Linux is like a mixture of CentOS and RHEL. Some of the "stuff" from both like yum for example would work on Amazon Linux. However I am searching for a more clear answer on this, for example if I am installing a software that is supported on CentOS7 / RHEL7 and CentOS6 / RHEL6. Given that my machine is Amazon Linux, I am not sure which version to install. Thanks.
## [6][13 lessons learned from taking 8 AWS certification tests in 4 weeks](https://www.reddit.com/r/aws/comments/hma3bk/13_lessons_learned_from_taking_8_aws/)
- url: https://medium.com/@quinn.richard/14-lessons-learned-from-taking-8-aws-certification-tests-in-4-weeks-b10b2c296c14
---

## [7][Uploading multipart/data-form image through api gateway with mapping templates generates “cannot read property buffer of undefined” NodeJS](https://www.reddit.com/r/aws/comments/hmskvq/uploading_multipartdataform_image_through_api/)
- url: https://www.reddit.com/r/aws/comments/hmskvq/uploading_multipartdataform_image_through_api/
---


I'm trying to use api gateway to upload an image to s3, i'm using http integration (no proxy, no lambda), the problem is that i have to set a condition on the headers in order to accept a POST request, so i have to edit template mappings, the image upload works fine with no mapping templates applied (passthrough option), but when i add the conditions or anything else in the mapping templates, the request will fail and i will get this error "cannot read property buffer of undefined" I suspect that the body has been changed, ans so is the buffer lost in there, i want to know how to use template mappings to keep the body (image) intact when uploading, while at the same time keeping my header conditions functional.

I'm using :

content-type:multipart/form-data. Backend is in Nodejs. Multer in NodeJS.
## [8][Made this serverless URL shortener using Python and AWS as a learning project, feedbacks are welcome](https://www.reddit.com/r/aws/comments/hm98b9/made_this_serverless_url_shortener_using_python/)
- url: https://github.com/SkullTech/shorty.serverless
---

## [9][New – Create Amazon RDS DB Instances on AWS Outposts](https://www.reddit.com/r/aws/comments/hmg1at/new_create_amazon_rds_db_instances_on_aws_outposts/)
- url: https://aws.amazon.com/blogs/aws/new-create-amazon-rds-db-instances-on-aws-outposts/
---

## [10][Deploying Docker app to AWS ElasticBeanstalk](https://www.reddit.com/r/aws/comments/hmp3n2/deploying_docker_app_to_aws_elasticbeanstalk/)
- url: https://www.reddit.com/r/aws/comments/hmp3n2/deploying_docker_app_to_aws_elasticbeanstalk/
---
Recently I experienced a strange problem that I hadn't had before with AWS EB.

My AWS EB foobar-api application has 2 environments:
- foobar-api-staging
- foobar-api-production

It's a Docker environment with single container setup. Containers are built and deployed to AWS ECR. 

All of a sudden stuff from my development branch ended up in production environment that was obviously not wanted. 

In my Dockerrun.aws.json the docker repo URL is specified like that:
```
 "Image": {
    "Name": "foo.dkr.ecr.us-east-1.amazonaws.com/bar/foobar-api:latest",
    "Update": "true"
  }
```

So I thought about it and it seems to me there are 2 possible issues:
1. Update should be set to "false"?
2. I could use separate repository URL-s / tags for production and staging.

While 1. is simple fix the second is a bit trickier.

So here's the question - how does anyone handle setups like these? How have you configured it?

Thanks for thinking along.
## [11][Linux running on Amazon Workspaces has discrepancy with IP addresses](https://www.reddit.com/r/aws/comments/hmovvg/linux_running_on_amazon_workspaces_has/)
- url: https://www.reddit.com/r/aws/comments/hmovvg/linux_running_on_amazon_workspaces_has/
---
In my AWS console, it lists a particular IP address. I see this same IP address listed as one of the addresses listed when I do ifconfig -a. However there are a lot of other addresses so I'm not sure which one is the "right" one. And then finally when I do one of those online "what's my IP address sites", I get an IP address that I have never seen before. Just trying to understand these discrepancies. What am I not understanding? Is this a known Workspaces issue? Thanks.
