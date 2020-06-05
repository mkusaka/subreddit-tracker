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
## [2][Slack adopts Chime Video Call technology](https://www.reddit.com/r/aws/comments/gwt7vu/slack_adopts_chime_video_call_technology/)
- url: https://www.cnbc.com/2020/06/04/amazon-licenses-slack-for-workers-as-slack-adopts-aws-video-call-tech.html
---

## [3][I've created a tool - spotcost.net The one-pager about all spot instances information. It helps find the cheapest region/az, compare specs, regions, price in time and etc. The difference between regions is huge (10-300% sic!). What do you think?](https://www.reddit.com/r/aws/comments/gwgdkm/ive_created_a_tool_spotcostnet_the_onepager_about/)
- url: https://spotcost.net
---

## [4][Why are NLB healthchecks coming from public IP address?](https://www.reddit.com/r/aws/comments/gx46m3/why_are_nlb_healthchecks_coming_from_public_ip/)
- url: https://www.reddit.com/r/aws/comments/gx46m3/why_are_nlb_healthchecks_coming_from_public_ip/
---
1. Create a VPC (10.0.0.0) and subnet (10.0.0.0/24) with internet gateway and normal routing table.
2. Create internal NLB (10.0.0.151), listener (8000), and target group (instance).
3. Launch EC2 instance (10.0.0.184) and register to target group.
4. Register instnace to target group.
5. On the instance, run python -m SimpleHTTPServer

Output

    209.17.97.106 - - [05/Jun/2020 12:38:42] "GET / HTTP/1.1" 200 -
    209.17.96.242 - - [05/Jun/2020 12:39:53] "GET / HTTP/1.1" 200 -
    209.17.96.122 - - [05/Jun/2020 12:40:58] "GET / HTTP/1.1" 200 -
    209.17.97.10 - - [05/Jun/2020 12:41:10] "GET / HTTP/1.1" 200 -

Why are the health checks coming from a public IP address? As I read the documentation, they should come from the NLB's private IP address?
## [5][New – Amazon EC2 C5a Instances Powered By 2nd Gen AMD EPYC™ Processors](https://www.reddit.com/r/aws/comments/gws4qo/new_amazon_ec2_c5a_instances_powered_by_2nd_gen/)
- url: https://www.reddit.com/r/aws/comments/gws4qo/new_amazon_ec2_c5a_instances_powered_by_2nd_gen/
---
&amp;#x200B;

https://preview.redd.it/jjnuvb51gz251.jpg?width=680&amp;format=pjpg&amp;auto=webp&amp;s=f02a2f00b5d054579c5e26db7e7297757cbddb91

[https://aws.amazon.com/blogs/aws/new-amazon-ec2-c5a-instances-powered-by-2nd-gen-amd-epyc-processors/](https://aws.amazon.com/blogs/aws/new-amazon-ec2-c5a-instances-powered-by-2nd-gen-amd-epyc-processors/)
## [6][Update auth token whenever user permissions are modified elsewhere](https://www.reddit.com/r/aws/comments/gx1cvh/update_auth_token_whenever_user_permissions_are/)
- url: https://www.reddit.com/r/aws/comments/gx1cvh/update_auth_token_whenever_user_permissions_are/
---
I'm creating a web application with user authentication.

I'll have multiple customers, each account with multiple users. The user permissions will be stored in his auth token upon login (issued by **Amazon Cognito**).

I'm facing the problem of **possible old auth tokens**: if **userA** modifies **userB**'s permissions, but **userB** currently has an active session, there's an ≤1h window of time where **userB** will have his old auth token (before it refreshes), containing claims that represent his old permission set. So **userB** may still be accessing and changing information that he shouldn't be allowed to anymore.

What to do in this situation?

I would like to be able to refresh a user's active sessions (i.e. update all currently valid auth tokens) remotely whenever someone modifies that user's permission set, so he gets his tokens updated.

This sounds to me like event listening. So, is **EventBridge** the tool I should use here? Can I make a web client listen to and react whenever my updatePermissions Lambda function is executed **with that user as argument**, and then refresh the user's token? If possible, is it cost-effective?

Or am I thinking it wrong?

Thanks
## [7][Whenever I update my Fargate container via CloudFormation, ECS deploys the change in minutes, but CloudFormation marks the service UPDATE_IN_PROGRESS for many minutes more. Why?](https://www.reddit.com/r/aws/comments/gwz5yw/whenever_i_update_my_fargate_container_via/)
- url: https://www.reddit.com/r/aws/comments/gwz5yw/whenever_i_update_my_fargate_container_via/
---
I've had this happen reproducably. It happens for tasks with a load balancer, and an explicit, short health check period. It also happens with tasks with no load balancer.

What is CloudFormation waiting for?
## [8][AWS S3 Monetization](https://www.reddit.com/r/aws/comments/gx2ofg/aws_s3_monetization/)
- url: https://www.reddit.com/r/aws/comments/gx2ofg/aws_s3_monetization/
---
Hello guys, I'm starting a movie watching website and the way I'll win money is by ads, I am wondering how could I monetize my S3 videos(movies) with ads before the video starts? Is there any way?

Thanks :)
## [9][ELI5: Asynchronous API push notifications](https://www.reddit.com/r/aws/comments/gwur3y/eli5_asynchronous_api_push_notifications/)
- url: https://www.reddit.com/r/aws/comments/gwur3y/eli5_asynchronous_api_push_notifications/
---
I have a couple of lambdas that need to execute in sequence calling external APIs and may possibly fail for various reasons depending on the health of the third party. For this reason the API calls that need to go externally outside the bounds of my AWS, once a response has been received, how do I alert the react client  (am using Amplify) that the process has completed, other than just making requests to my api asking for job updates?
## [10][Lambda vs EC2 vs locally running](https://www.reddit.com/r/aws/comments/gwrnzd/lambda_vs_ec2_vs_locally_running/)
- url: https://www.reddit.com/r/aws/comments/gwrnzd/lambda_vs_ec2_vs_locally_running/
---
Hi guys, I'm extremely new to AWS. I watched a few of the Amazon AWS training modules and was introduced to Lambda and EC2. From my understanding, they are both computing services. I googled what the difference was between the two and what I read was that Lambda is used for smaller and quicker executions whereas EC2 is used for more high-performance executions. But why would you get lambda in the first place if all you're doing is computing simple tasks? Wouldn't it be easier to compute it locally? I would appreciate if someone could give examples of what kind of situation you would use each kind of resource. Sorry if the question is a bit elementary, I was just introduced to the subject.
## [11][Media Services / CloudFront - amount of viewers](https://www.reddit.com/r/aws/comments/gwpg0q/media_services_cloudfront_amount_of_viewers/)
- url: https://www.reddit.com/r/aws/comments/gwpg0q/media_services_cloudfront_amount_of_viewers/
---
Hey all,

Question from a novice in the AWS environment. I successfully deployed a MediaLive, MediaPackage and CloudFront in AWS for live streaming. Works like a charm. 

Is there a way to see how many viewers/connections you have during a stream? I see number of Request Count in CloudFront, but no actually of real viewers/connections. My customers would love to know that. Is it possible to see that info?

Thanks!
