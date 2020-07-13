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
## [2][EKS now supports 1.17](https://www.reddit.com/r/aws/comments/hq0f4d/eks_now_supports_117/)
- url: https://github.com/awsdocs/amazon-eks-user-guide/blob/master/doc_source/kubernetes-versions.md#kubernetes-117
---

## [3][Best practises, when adding new Admin users to IAM](https://www.reddit.com/r/aws/comments/hqcc5e/best_practises_when_adding_new_admin_users_to_iam/)
- url: https://www.reddit.com/r/aws/comments/hqcc5e/best_practises_when_adding_new_admin_users_to_iam/
---
I'm a bit new to AWS and IAM. I'm wondering what is the best practises, when onboarding a new (admin) user, which has an existing aws account. 

Is it best practises to create a new account (form the IAM console or elsewhere), and make the user sign in with that account. Then use the process everytime a new user needs to access our AWS solution?

Or should i invite their existing user to our IAM group, and just le the users keep their original account (if this is the case, would could i do this?)?

Thank you in advance.
## [4][EC2 - what makes an instance "different"](https://www.reddit.com/r/aws/comments/hqbaaq/ec2_what_makes_an_instance_different/)
- url: https://www.reddit.com/r/aws/comments/hqbaaq/ec2_what_makes_an_instance_different/
---
Terrible title, but I couldn't think of a snappy way to word this.  I have a client that has installed a desktop version of some editing software on a Windows EC2 instance to make use of more power/memory.  If the instance is stopped for a couple of hours and then restarted the software in question thinks it has been installed on a different machine and needs to be reactivated.

I was wondering, how does the instance change so that it looks different to the software running on it? The IP address / machine name stays constant.  I'm thinking it's probably the MAC address that is different? (Which I'll be checking in an hour or two when we restart).

It's been a while since I did much with EC2 so I might be missing something very obvious here.
## [5][Where to find Master Username for AWS RDS?](https://www.reddit.com/r/aws/comments/hqc5ot/where_to_find_master_username_for_aws_rds/)
- url: https://www.reddit.com/r/aws/comments/hqc5ot/where_to_find_master_username_for_aws_rds/
---
See title. Want to connect to my database but every article regarding the Master Username seems to be outdated.

EDIT: Solved!
## [6][AWS SageMaker "default app" billing me?](https://www.reddit.com/r/aws/comments/hqc6mq/aws_sagemaker_default_app_billing_me/)
- url: https://www.reddit.com/r/aws/comments/hqc6mq/aws_sagemaker_default_app_billing_me/
---
Does the SageMaker default app being "ready" mean I am being billed for an instance? This is not clear. I can delete it everytime I finish using the studio (as it gets created when I open a studio(, but that would be one extra step I might not need to do.

https://preview.redd.it/bmezz2jrela51.png?width=1275&amp;format=png&amp;auto=webp&amp;s=4bf848f533c67dfb6f135f3bb16759e66e583023
## [7][Analysing up to 100k messages per second, what approach to take?](https://www.reddit.com/r/aws/comments/hpqf22/analysing_up_to_100k_messages_per_second_what/)
- url: https://www.reddit.com/r/aws/comments/hpqf22/analysing_up_to_100k_messages_per_second_what/
---
So I got the nice challenge to process messages from up to 100k users per second on a web page and the results need to be calculated in (near)realtime. The data coming in is structured and consists of a string and and amount and what I need to achieve is to get the combined sum of amounts for the unique strings (ie; some kind of voting system). So input looks something like:

TEAM\_A;10  
TEAM\_A:2  
TEAM\_B:1  
TEAM\_B:12  
TEAM\_A:4  
TEAM\_B:5  


output must be something like:  
TEAM\_A:16  
TEAM\_B:18  


and this must basically process every second. I've created 2 proof of concepts and both are using IoT Core over MQTT to send results to AWS; this works great. To process the incoming data I have 2 proof of concepts right now that work (not tested with 100k though!):  


A) Send IoT data to SQS and attach a Lambda. The Lambda will read up to 10 messages and store the data in ElastiCache Redis. While this works great for \~10k users, I don't think it will scale to 100k because it will require way too many concurrent Lambdas.

B) Send IoT data to Kinesis Firehose and then attach a Kinesis Analytics stream, the results will be pushed every second to a Lambda. This works really well however I'm reading Firehose has a 5k message/second limit which seems really low for a service called "Firehose". Also; if it could be adjusted to 100k/s, maybe the amount of data would be too big for Kinesis Analytics.

I'm leaning towards option B as this seems to be quite cost efficient. For scaling solution B I'm thinking I could create multiple IoT rules/firehose/analytics streams and let clients send their data to one of those rules round robin. Then combine results of these streams in Redis or DynamoDB and push them to the moderator.

Curious on your thoughts if I'm missing something here or if you have any experiences of your own ingesting/analysing this amount of data. I prefer to do everything using Cloudformation and serverless/managed solutions.
## [8][Open &amp; click event tracking](https://www.reddit.com/r/aws/comments/hq92l8/open_click_event_tracking/)
- url: https://www.reddit.com/r/aws/comments/hq92l8/open_click_event_tracking/
---
Actually  i am going to track open &amp; click event of sent mail via amazon ses  .I am using our subdomain actually for this event but still after  configuring it with sns We get notification for DELIVERY EVENT, but not  receiving notification for OPEN and CLICK Event

I  have read reagarding this on amazon page that we have to disable the  dkim settings but if suppose we want to getting such emails without  disabling the dkim settings is it possible to do it? If yes then how?

[https://forums.aws.amazon.com/thread.jspa?messageID=894418](https://forums.aws.amazon.com/thread.jspa?messageID=894418)
## [9][Sync Local Database Across Multiple Devices](https://www.reddit.com/r/aws/comments/hq5hzd/sync_local_database_across_multiple_devices/)
- url: https://www.reddit.com/r/aws/comments/hq5hzd/sync_local_database_across_multiple_devices/
---
Hi All

I am a complete beginner learning programming and currently working on an app built with Flutter.

The user will store data locally on the mobile device, but I am currently looking for a way whereby the user can still have access to that data across multiple devices or in the case they were to change phone etc they will still have retained that data.

I am looking for a way that this can happen automatically, so as whenever there is an update to the local DB i.e. users makes a change, input etc it is automatically updated elsewhere also.

I see that there is the CloudKit option for iOS, which is not an option as I am building in Flutter. Forgive me if I am way off here, however having done some reading is AWS Datastore and Appsync a similar thing? i.e. - user stores data locally using local SQLite and with Datastore and Appsync constant querying via GraphQL can see any updates made to local database and sync them across devices?
## [10][Static website on S3 + CloudFront. Is there a way to set up wildcard and splat URL redirect rules? (Aka all urls with ` /blog/* ` redirects to ` /articles/* `](https://www.reddit.com/r/aws/comments/hq0jed/static_website_on_s3_cloudfront_is_there_a_way_to/)
- url: https://www.reddit.com/r/aws/comments/hq0jed/static_website_on_s3_cloudfront_is_there_a_way_to/
---
I have around 200 old redirects on a website I need to preserve.  I see that one option is writing XML redirect rules in the S3 config, but that is limited to only 50 redirects. I saw another method with no rule limit is to set `x-amz-website-redirect-location` metadata on objects. However, these methods still require writing the exact URL and I can't use `/blog/*` or `/blog/:title`

Is there another way to be able to redirect with these type of rules? I read somewhere that Lamdba@Edge might be an option? I have not seen any examples and have never used it, so any resources I could be pointed to would be a big help!
## [11][(ACM) Certificate validation via DNS configuration](https://www.reddit.com/r/aws/comments/hq1x6h/acm_certificate_validation_via_dns_configuration/)
- url: https://www.reddit.com/r/aws/comments/hq1x6h/acm_certificate_validation_via_dns_configuration/
---
Hi everyone,

I've created an SSL certificate via ACM and I'm attempting to validate it via DNS configuration. 

I've got the CNAME record for the DNS configuration, but I've run into something that I'm not sure about. My registar, namecheap, won't allow me to set a CNAME on a domain while still keeping the custom DNS nameservers on the domain.

Eg. 

ns-1281.awsdns

[ns-1702.awsdns-26.co.uk](https://ns-1747.awsdns-26.co.uk)

[ns-211.awsdns-42.com](https://ns-343.awsdns-42.com)

&amp;#x200B;

In order for AWS to validate the domain, would the DNS nameservers not need to be set in addition to the CNAME?

&amp;#x200B;

Thanks!
