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
## [2][AWS Encyclopedia -- A place to share all the useful AWS documentation, repos, and solutions you have bookmarked!](https://www.reddit.com/r/aws/comments/gskv0k/aws_encyclopedia_a_place_to_share_all_the_useful/)
- url: https://github.com/cschultz82/aws_encyclopedia
---

## [3][Best Practices for CloudFormation/Serverless](https://www.reddit.com/r/aws/comments/gssiz4/best_practices_for_cloudformationserverless/)
- url: https://www.reddit.com/r/aws/comments/gssiz4/best_practices_for_cloudformationserverless/
---
I am curious what the best practices are with respect to CI/CD and rebuilding/deploying your applications.

Should certain resources be permanent (S3 Buckets, CloudFront Distributions) or should I wipe away everything and rebuild?

I'd rather not have to remember which steps I did manually when trying to rebuild our entire stack from scratch.
## [4][www.amazon.com is down!](https://www.reddit.com/r/aws/comments/gscvwp/wwwamazoncom_is_down/)
- url: https://www.reddit.com/r/aws/comments/gscvwp/wwwamazoncom_is_down/
---
% nslookup www.amazon.com

Server: 8.8.8.8

Address: 8.8.8.8#53

server can't find www.amazon.com: SERVFAIL

Colorado, USA

&amp;#x200B;

HN Thread: [https://news.ycombinator.com/item?id=23341170](https://news.ycombinator.com/item?id=23341170)
## [5][AWS ECS spot instances availability/creation](https://www.reddit.com/r/aws/comments/gsqqje/aws_ecs_spot_instances_availabilitycreation/)
- url: https://www.reddit.com/r/aws/comments/gsqqje/aws_ecs_spot_instances_availabilitycreation/
---
I have a m5.large  
 spot instance running on an ECS cluster. What will happen if this instance becomes unavailable? Will ECS spawn a new instance automatically? if not, then how do we handle this situation so that our application doesn't go down?

I am aware of services like Spotinist but I am trying to understand the flow here and also what does AWS provide by default?
## [6][Apparently I don't understand CloudWatch metrics, or how math works](https://www.reddit.com/r/aws/comments/gsn6zi/apparently_i_dont_understand_cloudwatch_metrics/)
- url: https://www.reddit.com/r/aws/comments/gsn6zi/apparently_i_dont_understand_cloudwatch_metrics/
---
Could someone explain this to me like I'm 5?

I'm trying to see (on a chart) how much S3 outgoing bandwidth I'm burning every month. My bill says it's been 457GB last month and 400GB the month before.

Let's say I want to set up a CloudWatch dashboard that shows me this, and I want to look at the past X months, and know how much bandwidth was consumed each month.

I set the statistic to "Sum", period to 30 days, and when I choose the past 3 months, suddenly it says I'm using 800GB over the past couple months. With the configuration I've chosen, doesn't this mean "show the past 3 months, summing up total bandwidth over 30 days"?

If I switch period to 15 minutes and mouse over the data points, it's showing about 300MB every 15 minutes. Is that 300MB of data consumption every 15 minutes? I don't see how that's possible.

(is CloudWatch Metrics maybe not the best place to do this, and I should use the Cost Explorer instead?)
## [7][Best practice API design for](https://www.reddit.com/r/aws/comments/gsqqwp/best_practice_api_design_for/)
- url: https://www.reddit.com/r/aws/comments/gsqqwp/best_practice_api_design_for/
---
The system was designed with only end users in mind. Primary authentication method is a custom auth lambda using userpools as the authorisation server. That provides a policy based on the user account. The project has expanded a little and requires a small rethink on design. I'm now in the position where the best approach may be to get the internal lambdas to call the api, however as far as I know I cannot create a userpool account for each of the lambdas to access the authorisation server to receive jwts.

One approach I can think of is to decouple the calls to the lambdas and the api and adding a layer between them so that the architecture arguably would look something more akin to that of a microservice ecosystem.
## [8][Cognito refresh token expiry?](https://www.reddit.com/r/aws/comments/gsne83/cognito_refresh_token_expiry/)
- url: https://www.reddit.com/r/aws/comments/gsne83/cognito_refresh_token_expiry/
---
How can I tell when a refresh token is due to expire?

I know how long it lasts, but I don't know when it was issued, so that's not helpful.

I can't decode it like an access token or id token.

Is there something in the SDK that can give me info about a refresh token? Struggling to find any useful docs on this.
## [9][Slow query even after upgrading to larger instance?](https://www.reddit.com/r/aws/comments/gsnrpe/slow_query_even_after_upgrading_to_larger_instance/)
- url: https://www.reddit.com/r/aws/comments/gsnrpe/slow_query_even_after_upgrading_to_larger_instance/
---
* i'm currently running the free tier
* currently using rds
* cpu is just running 1% cloudwatch
* read and write bearly hitting 5 iops cloudwatch
* i changed it to instance Class db.r3 large

&amp;#x200B;

is this considered slow?

* 3 index scans 2 seq scan,1  Bitmap Heap Scan only running 1 query
* Aggregate4 , CTE Scan 7 , Group 2 , Index Scan 4 , Nested Loop Inner Join 5 , Nested Loop Left Join4 , Sort 4 , Subquery Scan 1

do i need to modify something in postgres?

&amp;#x200B;

when i query it doesn't not seams to improve it still queries the same amount of time
## [10][NLB now supports ALPN on TLS listeners](https://www.reddit.com/r/aws/comments/gsdo3w/nlb_now_supports_alpn_on_tls_listeners/)
- url: https://www.reddit.com/r/aws/comments/gsdo3w/nlb_now_supports_alpn_on_tls_listeners/
---
Elastic Load Balancing now supports Application-Layer Protocol Negotiation (ALPN) policies on Network Load Balancers. ALPN is a TLS extension supported by all major browsers that enables negotiation of the protocol used after establishing a TLS connection, such as HTTP/2. Using ALPN policies, you can now offload your application’s TLS HTTP/2 traffic decryption/encryption to the Network Load Balancer, improving your service security posture and reducing operational complexity.

To get started, simply attach an ALPN policy to your Network Load Balancer TLS listener. The policy can be viewed and changed at any time based on your application’s protocol requirements. When ALPN is enabled, you can use Network Load Balancer TLS access logs to track successful and unsuccessful ALPN negotiations, view clients’ protocol preference lists, identify anomalies and debug connection issues.

Network Load Balancer ALPN policies are now available in all AWS Regions. To learn more, please refer to the Network Load Balancer documentation http://docs.aws.amazon.com/elasticloadbalancing/latest/network/introduction.html
## [11][Need Help - eCommerce Architecture](https://www.reddit.com/r/aws/comments/gsrvvm/need_help_ecommerce_architecture/)
- url: https://www.reddit.com/r/aws/comments/gsrvvm/need_help_ecommerce_architecture/
---
Folks,

I am looking to develop a simple eCommerce website architecture in AWS.

Could you please suggest any example architecture.

Thank you.
