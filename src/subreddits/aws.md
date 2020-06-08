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
## [2][Hard time getting TLS/SSL to work with Network Load Balancer](https://www.reddit.com/r/aws/comments/gyyf04/hard_time_getting_tlsssl_to_work_with_network/)
- url: https://www.reddit.com/r/aws/comments/gyyf04/hard_time_getting_tlsssl_to_work_with_network/
---
 Setup: ASP.NET Core Web App on ECS

We are attempting to use an NLB to pass SSL all the way through to the host. So far, I've been able to get non SSL TCP to work over port 80, but when trying TLS over port 443 the health checks keep failing which is causing ECS to keep shutting down the Task.

Do I need to do anything special with the ASP.NET Core app to allow this? When running it locally in Docker, both 80 and 443 work just fine. I'm assuming in that case it's a self signed cert. The Dockerfile does contain "Expose 443", which again, works locally, so I would've expected this to transfer the same to ECS.

Without any logs contain the reason as to why the healthchecks are failing, I flying blind right now and have no idea where to look. My assumption is that it's SSL related, but I'm not sure.
## [3][Calculated field in Cloudwatch Log Insights](https://www.reddit.com/r/aws/comments/gyx7qf/calculated_field_in_cloudwatch_log_insights/)
- url: https://www.reddit.com/r/aws/comments/gyx7qf/calculated_field_in_cloudwatch_log_insights/
---
Hi, I have a query, but want to be able to sort by the duration, which I am calculating as "abs(start-end).

&amp;#x200B;

    fields @timestamp, dstPort, bytes, srcAddr, dstAddr, @message, abs(start-end)
    | filter  dstPort = 80 and dstAddr = "10.2.6.242" 
    | sort @timestamp desc
    | limit 2000

How can I do that? I had a play around with Parse but I don't think that is what I am looking for, but basically I think I want something like:\~  


    fields @timestamp, dstPort, bytes, srcAddr, dstAddr, @message, duration as abs(start-end)
    | filter  dstPort = 80 and dstAddr = "10.2.6.242" 
    | sort duration desc
    | limit 2000

But that doesn't work..

Also, does anyone know the standard size (in bytes) of an HTTP healthcheck from an NLB? I am seeing a lot of connections that are 164 bytes or 362 bytes so wondering if it is either of those? Or is it just not that easy?
## [4][.Net Worker Service on Fargate vs Lambda](https://www.reddit.com/r/aws/comments/gyur6y/net_worker_service_on_fargate_vs_lambda/)
- url: https://www.reddit.com/r/aws/comments/gyur6y/net_worker_service_on_fargate_vs_lambda/
---
Hi, 

I’m relatively new to dotnet and trying to understand the benefits of a background service on Fargate (specifically .NET Core Worker Service)

My use case is that I want to query an RDS Database at 6pm everyday and check if a user has met specific conditions or not. If not, then send them a text.

Would a .NET Worker Service running on Fargate be a good solution for this or having a lambda run at those times be a better approach?

Trying to understand if a worker service would be an overkill for this use case. 

Thank you.
## [5][AWS newbie -- setting up periodic data collection with lambda functions](https://www.reddit.com/r/aws/comments/gyqq3s/aws_newbie_setting_up_periodic_data_collection/)
- url: https://www.reddit.com/r/aws/comments/gyqq3s/aws_newbie_setting_up_periodic_data_collection/
---
This post is to see if the AWS solution I came up with is reasonable and see if anyone suggests alternatives:

The goal is to gather and save traffic data from here.com's api every few minutes.  It's to collect sample data for analysis tool -- and to get familiarity with some AWS tools.  It's not mission-critcal.

Here.com uses a bearer token setup -- you authenticate with oauth secret keys, and get an access token that goes in request headers, the token expires in 24 hours.  

I set up:
  * two lambda functions, refresh-token and save-traffic-stats
  * three ssm parameters: oauth-id, oauth-secret, and bearer-token, all SecureString

Every 10 hour refresh-token, using the oauth parms,  updates the bearer-token parm (thinking odds are it'll never be down, but if there's some transient error it'll probably be better in ten hours)

Every few minutes save-traffic-stats runs, reads the bearer token, gets the traffic data, and saves to s3.

The functions are scheduled with CloudWatch fixed rate event rules.

Is that a reasonable/typical solution for this goal?

Other thoughts... 

I read that if you were doing a lot of traffic, the ssm Parameter Store standard tier might throttle you -- the limits are so far beyond what I'm doing that's of no concern for me.

For this application it's not clear to me that it's any easier, more flexible, or more secure than running scripts by cron/at on an ec2 host.  I did hit one complication: I'm using python and the popular "request" http library is deprecated, so I had to fall back to comes-with-python urllib.

The setup so far has been  entirely thru the AWS Console and ad-hoc cli scripts.  My next goal is to set up terraform or similar to make the set up paramaterized/recreatable.
## [6][Why are all new posts pending moderator review now?](https://www.reddit.com/r/aws/comments/gyy4a2/why_are_all_new_posts_pending_moderator_review_now/)
- url: https://www.reddit.com/r/aws/comments/gyy4a2/why_are_all_new_posts_pending_moderator_review_now/
---
Very frustrating because I have come to rely on this sub for timely help and now o have to wait for a moderator to approve.  Why is that now?  None of the other subs I frequent require that.
## [7][Slack’s new integration deal with AWS could also be about tweaking Microsoft](https://www.reddit.com/r/aws/comments/gyiq2i/slacks_new_integration_deal_with_aws_could_also/)
- url: https://techcrunch.com/2020/06/05/slacks-new-integration-deal-with-aws-could-also-be-about-tweaking-microsoft/
---

## [8][question on API gateway client certificate](https://www.reddit.com/r/aws/comments/gyuvcm/question_on_api_gateway_client_certificate/)
- url: https://www.reddit.com/r/aws/comments/gyuvcm/question_on_api_gateway_client_certificate/
---
I have a REST API that's using Lambda as the "backend". My boss hired a third party VA/PT engineer to check the configuration of the application and then I got a report that I should be enabling API gateway's client certificate to let my back end know that requests are coming from API Gateway. 

Now please take note that I'm fairly new to AWS services (or security, for that matter) and from what I know AWS already requires API gateway to have "permissions" before it can access lambda. From that, I assume Lambda would already know that requests come from API gateway. 

My question is should I enable API gateway's client certificate when connecting to Lambda? and if so, how do I do that?
## [9][Migrating old instance type to the latest gen question](https://www.reddit.com/r/aws/comments/gyz5hy/migrating_old_instance_type_to_the_latest_gen/)
- url: https://www.reddit.com/r/aws/comments/gyz5hy/migrating_old_instance_type_to_the_latest_gen/
---
So we have lot of database and pub/sub server such as Kafka and Cassandra in dev and production in our start up. Amazon notified us our instance type are t2-3, m1 and m3 and they need to be migrated to the latest gen which is m5 I believe.The story with them is that AWS is considering the underlying type oh hardware as obsolete and more failure prone than their best availability standards. My question is how do I achieve this. Do I have to snapshot this first and deploy new instance ? Never migrated instances before so curious on how to do this safely ?

TLDR: Need to migrate old gen instance to the latest gen. How do I achieve in the best secure way without losing data and minimal downtime and outage. 
## [10][Would there be any interruption, when doing an AMI backup to a production server?](https://www.reddit.com/r/aws/comments/gyuhou/would_there_be_any_interruption_when_doing_an_ami/)
- url: https://www.reddit.com/r/aws/comments/gyuhou/would_there_be_any_interruption_when_doing_an_ami/
---
Hi,

we need to AMI backup a production server. we are worried that there will be interruption doing so.  
we don't want interruption, as the server is online to serve clients.
## [11][Configuring Free SSL certificate on IIS in EC-2 Instance](https://www.reddit.com/r/aws/comments/gyus2h/configuring_free_ssl_certificate_on_iis_in_ec2/)
- url: https://www.reddit.com/r/aws/comments/gyus2h/configuring_free_ssl_certificate_on_iis_in_ec2/
---
I have hosted my PHP MySQL web application on EC-2 on windows 2019 on free tier. I have assigned the elastic IP to the instance and pointed to the our domain which is hosted on third party. Now i would like to enable the Free SSL certificate by amazon certificate manager and i have created that. I am little bit confused in the next steps can anyone help me there.

1. Do i need to use Amazon Route 53 service for this? Is this service is available in free tier? Do i need to create an hosted zone in Route 53 with the domain name and create records and update these records into my third party domain provider portal? Is this step is really required in the free SSL setup to work on my domain?
2. Or Elastic Load Balancer is the only one i need to set up in the further process to work the SSL?

Can anyone advise me here and it will be really appreciated.

Thanks in advance!
