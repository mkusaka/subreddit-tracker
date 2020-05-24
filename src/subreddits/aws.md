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
## [2][I wrote a complete guide to the AWS Systems Manager Parameter Store. It's a stellar way to make your applications more scalable and less failure prone](https://www.reddit.com/r/aws/comments/gpnqin/i_wrote_a_complete_guide_to_the_aws_systems/)
- url: https://seanjziegler.com/a-complete-guide-to-using-the-aws-systems-manager-parameter-store/
---

## [3][Breaking down NAT costs with AWS flow logs](https://www.reddit.com/r/aws/comments/gp6d7u/breaking_down_nat_costs_with_aws_flow_logs/)
- url: https://www.reddit.com/r/aws/comments/gp6d7u/breaking_down_nat_costs_with_aws_flow_logs/
---
We used AWS flow logs to build some dashboards with a detailed breakdown of what our internal services were costing in terms of data transmission, allowing us to decrease NAT costs by some 40%.

I couldn't find a way to do this with the docs, so I wrote about the strategy we came up with on my blog: [https://badgateway.qc.to/breaking-down-aws-flow-logs/](https://badgateway.qc.to/breaking-down-aws-flow-logs/)
## [4][Amazon Virtual Private Cloud (VPC) now supports Bring Your Own IPv6 Addresses (BYOIPv6)](https://www.reddit.com/r/aws/comments/gpipzz/amazon_virtual_private_cloud_vpc_now_supports/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/05/amazon-virtual-private-cloud-supports-bring-own-ipv6-addresses/
---

## [5][RDS DB/serverless Lambda/S3 static React app not accessible in same account?](https://www.reddit.com/r/aws/comments/gppj34/rds_dbserverless_lambdas3_static_react_app_not/)
- url: https://www.reddit.com/r/aws/comments/gppj34/rds_dbserverless_lambdas3_static_react_app_not/
---
I have a Postgresql RDS DB with a Ruby on Jets Lambda and React S3 static app with public access. I'm getting a 403 forbidden when trying the given link. I'm also unable to view the json data when trying to access \`index\` or \`show\`. I've tried to create inbound/outbound rules on RDS allowing all traffic but not working. The only item that is on a VPC is the RDS DB. Is there something else I should know about having all 3 parts of the stack on the same AWS account?
## [6][Hi guys. I have a dumb question. What’s the difference between Amazon Web Services and Amazon Workspace?](https://www.reddit.com/r/aws/comments/gpp9m1/hi_guys_i_have_a_dumb_question_whats_the/)
- url: https://www.reddit.com/r/aws/comments/gpp9m1/hi_guys_i_have_a_dumb_question_whats_the/
---
Are they the same thing? Are they totally different things. I’m very confused. My job uses Amazon workspace so we can remotely access everything on the cloud. From my understanding, that’s the same thing as Amazon Web Services. So are they the same thing? Help pls
## [7][Does AWS Have IRCnet Servers or Connections Easily Available](https://www.reddit.com/r/aws/comments/gpnkmn/does_aws_have_ircnet_servers_or_connections/)
- url: https://www.reddit.com/r/aws/comments/gpnkmn/does_aws_have_ircnet_servers_or_connections/
---
I've noticed that it's very hard to connect from EC2, AWS Central Europe/Frankfurt/Germany to the IRC network called IRCnet.

By hard I mean that the connection gets blocked all the time when I try. And by try I mean that I went through some of the servers on [http://irc.tu-ilmenau.de/all\_servers/](http://irc.tu-ilmenau.de/all_servers/) and those ones that I tried were negative.

Does Amazon happen to host their own server or servers on IRCnet? I mean ones that are easily accessible from EC2. Or does someone know some IRCnet servers that do not block EC2 connections?

Are there some extra steps that can be taken in order to bypass this problem? I know that Freenode does not allow simple connection attempts from EC2 but they do allow SASL attempts with registered Freenode user credentials. I also know that IRCnet is anarchistic network without internal structure compared to Freenode, so I'm not getting my hopes up, but you know, mayyyybe there is some way.

Thank you in advance, if you happen have anything cool to chime in on this!
## [8][EC2 Instance User Accounts](https://www.reddit.com/r/aws/comments/gpn3ef/ec2_instance_user_accounts/)
- url: https://www.reddit.com/r/aws/comments/gpn3ef/ec2_instance_user_accounts/
---
Hi everyone,

I'm working my way through the course content for the AWS CSA certification, and am busy learning about setting up EC2 instances. I'm quite a noob i.t.o. AWS, so apologies if this is a stupid question.

I have some questions on user accounts. When launching an instance (let's say Amazon Linux AMI), as per the documentation "each Linux instance launches with a default Linux system user account". You can then create additional user accounts when you're SSH'd into the instance in the root account.

1. Does the user account used to access the EC2 Instance differ from the Users created in IAM? I assume this is the case? Are there any relationship between these accounts or an "easier way" to grant permissions through IAM to access the EC2 Instance?
2. According to the documentation on creating new user accounts ( [https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/managing-users.html](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/managing-users.html)), when setting up the account you add an SSH public key to the account. I assume this won't be the same public key from the root account, so you first then have to create a new key pair and generate the public key?
3. In practice, how many users would typically need to have access to the VM, and would thus need user accounts? Know this is a wide question, but just wondering.

Hope this make sense - thanks in advance :)
## [9][Getting http response code monitoring on NLB](https://www.reddit.com/r/aws/comments/gpn2m8/getting_http_response_code_monitoring_on_nlb/)
- url: https://www.reddit.com/r/aws/comments/gpn2m8/getting_http_response_code_monitoring_on_nlb/
---
I know this is ridiculous since NLB is L4 and ALB is L7 which is why ALB gets response code monitoring. But we really liked the fact that we could see how many requests returned 500 errors (including gateway timeout). 

But at the same time, NLB’s tls termination is one of the those “magic” features on aws and I was wondering if I could get similar support for it. So the question is, is there any way to get this response code monitoring at the load balancer level or at least something similar to it?
## [10][I am trying to download something off of AWS but I have no idea how](https://www.reddit.com/r/aws/comments/gpluxc/i_am_trying_to_download_something_off_of_aws_but/)
- url: https://www.reddit.com/r/aws/comments/gpluxc/i_am_trying_to_download_something_off_of_aws_but/
---
I have never used this service, and I've tried creating an account to download what I need but I have no idea what I'm doing.

Basically, I need to download some [BAM files](https://www.ncbi.nlm.nih.gov/sra/?term=SRP132033) on NCBI. The problem is that each download address for the BAM files is [like this](https://i.imgur.com/PyfWZ5i.png). I think it's indicating this is some kind of address on AWS (S3)? I have no idea what that is or how to access it. I tried following some instructions on the NCBI which made me create an account of some sorts on AWS, but it was asking for my card number and 50 other things which it shouldn't be if I'm just wanting to download a number of files.

Please help.
## [11][Infrastructure as Code Best Practices?](https://www.reddit.com/r/aws/comments/gp5tca/infrastructure_as_code_best_practices/)
- url: https://www.reddit.com/r/aws/comments/gp5tca/infrastructure_as_code_best_practices/
---
Hello All,

I've recently been approached about potentially publishing an Infrastructure as Code book. While I certainly have plenty of ideas, I wanted to get a general sense of what do actual practitioners need to learn in order to be successful with IAC?

1. I was thinking to have 3 parts to the book, 1 that focuses on executive leadership and tells the "why" story of IAC, 1 that focuses on the technical implementation of IAC, and the last that focuses on architectural patterns and use cases. What are your thoughts about this structure?
2. How did your organization adopt IAC? Did it start with the cloud? What does your maturity journey look like?
3. What're the most common pitfalls you've run across trying to do IAC? Are you able to do full automation or are there still manual steps?
4. How did you learn IAC? What was the most helpful?
