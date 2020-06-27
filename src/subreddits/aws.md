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
## [2][Reminder: Terminate your free WorkSpaces by 30-Jun-20](https://www.reddit.com/r/aws/comments/hgl5hz/reminder_terminate_your_free_workspaces_by_30jun20/)
- url: https://www.reddit.com/r/aws/comments/hgl5hz/reminder_terminate_your_free_workspaces_by_30jun20/
---
Just a quick reminder, AWS' free WorkSpaces offer ends 30-Jun-2020. If you were using it to test out the service, please remember to terminate it to avoid additional charges.

Thank you to Amazon for offering this.
## [3][AWS Copilot CLI: create, release and manage production ready containerized applications on Amazon ECS and AWS Fargate](https://www.reddit.com/r/aws/comments/hgbvsa/aws_copilot_cli_create_release_and_manage/)
- url: https://github.com/aws/copilot-cli
---

## [4][DynamoDB Filtering by Date Partition Keys](https://www.reddit.com/r/aws/comments/hgs0t0/dynamodb_filtering_by_date_partition_keys/)
- url: https://www.reddit.com/r/aws/comments/hgs0t0/dynamodb_filtering_by_date_partition_keys/
---
Hi, I have been exploring DynamoDb a lot recently but I am stuck on learning how to model the data appropriately.

I have an example table with columns as follows:

|shop\_id|last\_visited|opened|revenue|balance|average\_daily\_revenue|
|:-|:-|:-|:-|:-|:-|
|b4f4f00c07f|2020-06-27T08:20:14Z|2004-03-21T08:20:14Z|80000|1000|500|

&amp;#x200B;

Following these points from AWS documentation:

* **Hashkey/Parition Key should be a high cardinality attribute** \- therefore shop\_id is an obvious choice.
* **If you are retrieving items by a single field - that field should either be a hash\_key or a GSI**. So it seems fine to have shop\_id as the partition key as a large number of cases requires me to retrieve a single shop.

So at this point I am fairly certain that I can select shop\_id as the partition key for the table, making inserting and selecting single record by id very fast. So far so good.

The next challenge is that I would like to be able to filter based on all the other columns in the table: last\_visited, opened, revenue, balance etc. without knowing the shop\_id. All filters are optional except last\_visited which is compulsory. So I make the following decisions:

1. Since using **query** call requires you to use the partition key (shop\_id). I need to create a GSI. Ok great.
2. So now to select a GSI hashkey, knowing that last\_visited is the most common attribute in queries, I select that as the **partition key**.
3. BUT it turns out, I cannot perform a "BETWEEN" condition on the partition key, and therefore must use **scan** (which is discouraged).

The way I see it, inside dynamodb it is impossible to have a date key as partition key while using the query operation. This seems like a fairly common scenario. **(Please prove me wrong).**

The only other method I can think of is using last\_visited as both the partition key and the range key inside a GSI. But this feels a bit wrong. **Am I missing something?**

Side question: are all other filter fields are included as LSIs. **Are LSIs also operational on GSIs?**
## [5][Apache - (104)Connection reset by peer: [client x.x.x.x:4712] AH03308: ap_proxy_transfer_between_connections: error on sock - ap_get_brigade](https://www.reddit.com/r/aws/comments/hgsd7l/apache_104connection_reset_by_peer_client/)
- url: https://www.reddit.com/r/aws/comments/hgsd7l/apache_104connection_reset_by_peer_client/
---
I have a django app deployed to elastic beanstalk:

[http://django-env.eba-bcvm9cxz.us-west-2.elasticbeanstalk.com/chart/](http://django-env.eba-bcvm9cxz.us-west-2.elasticbeanstalk.com/chart/)

I'm following this tutorial: [https://medium.com/@elspanishgeek/how-to-deploy-django-channels-2-x-on-aws-elastic-beanstalk-8621771d4ff0](https://medium.com/@elspanishgeek/how-to-deploy-django-channels-2-x-on-aws-elastic-beanstalk-8621771d4ff0) and using architecture number 2.

My relevant configuration files are the folllowing:

        option_settings:  
          aws:elasticbeanstalk:application:environment:
            DJANGO_SETTINGS_MODULE: dashboard.settings
            PYTHONPATH: /opt/python/current/app/dashboard:$PYTHONPATH
          aws:elasticbeanstalk:container:python:
            WSGIPath: dashboard/wsgi.py
        
          aws:elbv2:listener:80:
            ListenerEnabled: 'true'
            Protocol: HTTP
          aws:elbv2:listener:5000:
            ListenerEnabled: 'true'
            Protocol: HTTP

\--

        files:
          "/etc/httpd/conf.d/proxy.conf":
            mode: "000644"
            owner: root
            group: root
            content: |
              ProxyPass /websockets/ ws://127.0.0.1:5000/websockets/ upgrade=NONE
              ProxyPassReverse /websockets/ ws://127.0.0.1:5000/websockets/

I needed to add \`upgrade=NONE\` in proxy.conf, which isn't in the tutorial, but I for me I always got a 503 on the client when it was establishing the websocket.

I have a daphne server listening on 5000, which is successfuly receiving the websocket request from the client, but the client never gets to receive the response from the server. On chrome console I get:

    WebSocket connection to 'ws://django-env.eba-bcvm9cxz.us-west-2.elasticbeanstalk.com/websockets/' failed: Error during WebSocket handshake: Unexpected response code: 408

Apache's logs show the following:

        [Sat Jun 27 11:42:11.962267 2020] [watchdog:debug] [pid 9855] mod_watchdog.c(590): AH02981: Watchdog: Created child worker thread (_proxy_hcheck_).
        [Sat Jun 27 11:42:11.962286 2020] [proxy:debug] [pid 9855] proxy_util.c(1940): AH00925: initializing worker ws://127.0.0.1:5000/websockets/ shared
        [Sat Jun 27 11:42:11.962298 2020] [proxy:debug] [pid 9855] proxy_util.c(2000): AH00927: initializing worker ws://127.0.0.1:5000/websockets/ local
        [Sat Jun 27 11:42:11.962319 2020] [proxy:debug] [pid 9855] proxy_util.c(2035): AH00930: initialized pool in child 9855 for (127.0.0.1) min=0 max=4 smax=4
        [Sat Jun 27 11:42:11.962330 2020] [proxy:debug] [pid 9855] proxy_util.c(1940): AH00925: initializing worker proxy:reverse shared
        [Sat Jun 27 11:42:11.962334 2020] [proxy:debug] [pid 9855] proxy_util.c(2000): AH00927: initializing worker proxy:reverse local
        [Sat Jun 27 11:42:11.962341 2020] [proxy:debug] [pid 9855] proxy_util.c(2035): AH00930: initialized pool in child 9855 for (*) min=0 max=4 smax=4
        [Sat Jun 27 11:42:11.962345 2020] [proxy:debug] [pid 9855] proxy_util.c(1935): AH00924: worker ws://127.0.0.1:5000/websockets/ shared already initialized
        [Sat Jun 27 11:42:11.962347 2020] [proxy:debug] [pid 9855] proxy_util.c(1992): AH00926: worker ws://127.0.0.1:5000/websockets/ local already initialized
        [Sat Jun 27 11:42:30.521413 2020] [proxy:debug] [pid 9625] proxy_util.c(4059): (104)Connection reset by peer: [client 172.31.42.96:4452] AH03308: ap_proxy_transfer_between_connections: error on sock - ap_get_brigade
        [Sat Jun 27 11:42:30.521453 2020] [proxy:debug] [pid 9625] proxy_util.c(2353): AH00943: WS: has released connection for (127.0.0.1)

What's wrong with my configuration?
## [6][Event bus, kinesis or sns + sqs?](https://www.reddit.com/r/aws/comments/hg9jod/event_bus_kinesis_or_sns_sqs/)
- url: https://www.reddit.com/r/aws/comments/hg9jod/event_bus_kinesis_or_sns_sqs/
---
Which one should I use for a generic event bus that would be used in multiple services? I have lots of handlers in my data pipeline and to do sns fanout I would need a queue per function. Is that too many queues? Can you create queues at runtime?
## [7][Anyone else seeing a lot of random timeouts with lambda? Started about an hour ago in the us-west-2 region](https://www.reddit.com/r/aws/comments/hgel50/anyone_else_seeing_a_lot_of_random_timeouts_with/)
- url: https://www.reddit.com/r/aws/comments/hgel50/anyone_else_seeing_a_lot_of_random_timeouts_with/
---
We haven't really changed anything or deployed anything new but all of a sudden starting around 11:20am Pacific Time, a lot of our lambda functions started timing out randomly.

Wondering if it's just us?

Edit: We make a lot of S3 calls in these aforementioned lambdas, as one of the commenters mentioned here, it might be related to that

&amp;#x200B;
## [8][Network packages analysis?](https://www.reddit.com/r/aws/comments/hgibs1/network_packages_analysis/)
- url: https://www.reddit.com/r/aws/comments/hgibs1/network_packages_analysis/
---
A customer is asking to implement a Firewall that can analyze network packages using Sandbox in their AWS infrastructure.

I haven't seen such product on AWS. Does WAF or Shield include that? This is a simple website that would be using Cloudflare.
## [9][Provisioned IOPS throughput calculation](https://www.reddit.com/r/aws/comments/hgn2nt/provisioned_iops_throughput_calculation/)
- url: https://www.reddit.com/r/aws/comments/hgn2nt/provisioned_iops_throughput_calculation/
---
I am trying to understand how throughput is calculated on provisioned IOPS drives. Can someone help me understand this graph?  Apparently an io1 drive set to 2K IOPS can go up to 500 MBps throughput. Does that just mean 32K to 64K can go up to 1,000 MBps? The graph in the documentation is difficult to read.

[https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-volume-types.html#EBSVolumeTypes\_piops](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-volume-types.html#EBSVolumeTypes_piops) 

&amp;#x200B;

Thank you!
## [10][Difficulty connecting to RDS instance](https://www.reddit.com/r/aws/comments/hggviz/difficulty_connecting_to_rds_instance/)
- url: https://www.reddit.com/r/aws/comments/hggviz/difficulty_connecting_to_rds_instance/
---
I have a MariaDB 10.3.13 instance running, but I think I'm missing something, cause can't connect from a Lightsail instance.

Public accessibility: Yes  
VPC security groups: Inbound [0.0.0.0/0](https://0.0.0.0/0)

What else am I overlooking here?
## [11][Dow Jones Hammer: A multi-account cloud security tool for AWS.](https://www.reddit.com/r/aws/comments/hgdqxc/dow_jones_hammer_a_multiaccount_cloud_security/)
- url: https://github.com/dowjones/hammer
---

