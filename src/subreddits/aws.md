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
## [2][Cognito User Pool: storing user ids &amp; tenant ids](https://www.reddit.com/r/aws/comments/hp73lc/cognito_user_pool_storing_user_ids_tenant_ids/)
- url: https://www.reddit.com/r/aws/comments/hp73lc/cognito_user_pool_storing_user_ids_tenant_ids/
---
I'm building a **multi-tenant** web application in AWS, with Amazon Cognito as user directory and DynamoDB as data store.

Ps: “Tenant” here means a paying customer account. Each tenant can contain one or multiple users.

**My architecture choices**:

1. I would like to use Cognito (User Pool) as an autonomous source and “single source of truth” for my app's users.
2. I need to create my own custom **userIds** and **tenantIds** for each user (they will be referred to in DynamoDB items also), and store both of these in Cognito through **User Pool attributes**.  
(These attributes are automatically passed to the **id tokens** issued by Cognito upon user logins.)
3. I need to be able to retrieve a list of users (from my User Pool) that belong to any specific tenant within the system.

I could store the **userId** and **tenantId** as User Pool “**custom attributes**”. However, these custom attributes are not searchable by Cognito's [ListUsers API](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_ListUsers.html). This API can however search by the Cognito **standard user attributes** (username, email, name, given\_name, family\_name, sub, etc.).

To address this limitation of the ListUsers API, the idea I had was to store the **userId** and **tenantId** in one of these standard attributes, like ***given\_name*** and ***family\_name***, which I won't need to use, since I already have a "***name***" attribute that I can use to store the user's name.

**My question is****:**

**Is there any problem with this strategy? Anything I'm not seeing that could create problems?**

Or maybe there is a simpler way of dealing with this problem?
## [3][HTTPS on EC2 instance running python project](https://www.reddit.com/r/aws/comments/hp49to/https_on_ec2_instance_running_python_project/)
- url: https://www.reddit.com/r/aws/comments/hp49to/https_on_ec2_instance_running_python_project/
---
I'm having considerable difficulty getting HTTPS to resolve on my EC2 instance, which runs a python project. The request just times out (ERR\_CONNECTION\_TIMED\_OUT). HTTP runs ok, however. The steps I've taken are as follows.

1. I've created a certificate in ACM for the following domains: \*.mywebsite.com and mywebsite.com

[https://i.stack.imgur.com/QCTbF.png](https://i.stack.imgur.com/QCTbF.png)

&amp;#x200B;

2. I've setup Route 53 as follows:

[https://i.stack.imgur.com/qsdAm.png](https://i.stack.imgur.com/qsdAm.png)

Routing policy on the A records is Simple.

&amp;#x200B;

3. I've gone into the Listener for my Load Balancer for my EC2 instance and **CHANGED** the port from 80 (HTTP) TO 443 (HTTPS) and added my certificate.

Note: the "Forward To" is a Target Group running on port 80 (HTTP). I've read that this is correct.

[https://i.stack.imgur.com/8yYxQ.png](https://i.stack.imgur.com/8yYxQ.png)

&amp;#x200B;

4. I've then gone into the Inbound Rules for my Security group, and added HTTPS

[https://i.stack.imgur.com/TO8Wz.png](https://i.stack.imgur.com/TO8Wz.png)

&amp;#x200B;

At this point, I've got the following questions:

a) Given that this is a python/Django project, is enabling HTTPS for EC2 possible to do this through the AWS website or do I need to add config files and deploy to my instance?

b) Do I need to create a target group running on HTTPS?

c) Do I need listeners on my load balance for port 80 and port 443 or just port 443?

d) On my security group, do I need port 80 to go to 0.0.0.0/0 and ::0/?

e) Should the A record by the DNS name of the load balancer or should it be the CNAME of my environment?

Thanks for your help!
## [4][Is it possible to authenticate users with cognito ~without~ the web UI?](https://www.reddit.com/r/aws/comments/hp7rkm/is_it_possible_to_authenticate_users_with_cognito/)
- url: https://www.reddit.com/r/aws/comments/hp7rkm/is_it_possible_to_authenticate_users_with_cognito/
---
I'm familiar with using a user pool client id+secret to obtain a scoped access token but I'm having a hard time figuring out if it's possible to log in as a user in a similar manner—eg, using cURL or the AWS CLI.

I have gotten as far as using the `InitiateAuth` method with the `USER_PASSWORD_AUTH` flow to get a token but its only scope is `aws.cognito.signin.user.admin`.
## [5][pagination for workspaces, boto3, python, sdk](https://www.reddit.com/r/aws/comments/hp9qqj/pagination_for_workspaces_boto3_python_sdk/)
- url: https://www.reddit.com/r/aws/comments/hp9qqj/pagination_for_workspaces_boto3_python_sdk/
---
not a programmer by trade but trying to put together script to go through Workspaces and change the idle time shutdown settings on the autostop unit. Got the script to list the units that are set too high... didn't realize there was pagination and now am kinda confused. Was going to loop through NextToken but first set doesn't have a token. And yes, love stackoverflow, but can't find anything there to help (or is clear to me). Thx!
## [6][I cannot login into AWS using Vivaldi web browser, anyone having the same issue?](https://www.reddit.com/r/aws/comments/hp76ht/i_cannot_login_into_aws_using_vivaldi_web_browser/)
- url: https://www.reddit.com/r/aws/comments/hp76ht/i_cannot_login_into_aws_using_vivaldi_web_browser/
---
I've tried both the root account and the IAM user I've created for the console. Login request just returns me 400 Bad Request and says to clean cookies, however, it doesn't help at all.

Tried Chrome Stable and Firefox - works, so it must be something wrong with Vivaldi or how I set it up.

Anyone having the same issue?
## [7][AWS and Docker collaborate to simplify the developer experience](https://www.reddit.com/r/aws/comments/hoifac/aws_and_docker_collaborate_to_simplify_the/)
- url: https://aws.amazon.com/blogs/containers/aws-docker-collaborate-simplify-developer-experience/
---

## [8][Very long start times for RDS export to S3 tasks](https://www.reddit.com/r/aws/comments/hozoxu/very_long_start_times_for_rds_export_to_s3_tasks/)
- url: https://www.reddit.com/r/aws/comments/hozoxu/very_long_start_times_for_rds_export_to_s3_tasks/
---
I started playing around with RDS export to S3 feature. So far it seems like every task takes about 30 minutes just to start-up. Even before it accesses the data it just sits there on `STARTING` for 30 minutes. The export itself then takes about 10 minutes which makes sense. But having to wait 40 minutes total just to get an error that I got the table name wrong is disappointing.

Has anyone else had similar experience? Or are the servers doing the export just overloaded today?
## [9][Videos from fwd:cloudsec 2020 are now online](https://www.reddit.com/r/aws/comments/hoqw4u/videos_from_fwdcloudsec_2020_are_now_online/)
- url: https://www.reddit.com/r/aws/comments/hoqw4u/videos_from_fwdcloudsec_2020_are_now_online/
---
fwd:cloudsec is a non-profit community organized cloud security conference that was planned for the day before re:Inforce, but switched to being a virtual conference when re:Inforce cancelled.

Playlist: [https://www.youtube.com/playlist?list=PLCPCP1pNWD7OBQvDY7vLCFhxWxok9DITl](https://www.youtube.com/playlist?list=PLCPCP1pNWD7OBQvDY7vLCFhxWxok9DITl)

Sign up on the mailing list for announcements for next year's conference: [https://fwdcloudsec.org/](https://fwdcloudsec.org/)
## [10][Can I set up my lambda to not trigger more than once every X minutes?](https://www.reddit.com/r/aws/comments/hp2z71/can_i_set_up_my_lambda_to_not_trigger_more_than/)
- url: https://www.reddit.com/r/aws/comments/hp2z71/can_i_set_up_my_lambda_to_not_trigger_more_than/
---
I have a Lambda that I'm invoking, hopefully, once every X minutes. But I'd like to be safe and make sure it's not being invoked any more often than that. Is this possible?
## [11][Creating a comment section tutorial using AWS Lambda, DynamoDB, SES](https://www.reddit.com/r/aws/comments/hp353d/creating_a_comment_section_tutorial_using_aws/)
- url: https://www.reddit.com/r/aws/comments/hp353d/creating_a_comment_section_tutorial_using_aws/
---
Hey guys!

I created a tutorial on how to build a simple comment section as part of getting familiar with some of  AWS' services.

[https://github.com/mannyray/AWScommentSection](https://github.com/mannyray/AWScommentSection)

&amp;#x200B;

What do you guys recommend I explore next in AWS?
