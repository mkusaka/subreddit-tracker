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
## [2][Getting the Most Out of AWS’ Online Summit](https://www.reddit.com/r/aws/comments/h9djvs/getting_the_most_out_of_aws_online_summit/)
- url: https://www.contino.io/insights/aws-emea-online-summit
---

## [3][I made a tool that makes me faster listing AWS resources and their attributes](https://www.reddit.com/r/aws/comments/h8yhbo/i_made_a_tool_that_makes_me_faster_listing_aws/)
- url: https://github.com/jckuester/awsls
---

## [4][List and delete snapshots](https://www.reddit.com/r/aws/comments/h9blzg/list_and_delete_snapshots/)
- url: https://www.reddit.com/r/aws/comments/h9blzg/list_and_delete_snapshots/
---
Could someone assist me with the right syntax to delete the  snapshot i created: 

I am unable to to delete the snapshot  with snap.delete()  by refe the boto3 doc . ----&gt;

 snap.delete()

AttributeError: 'dict' object has no attribute 'delete'

 

=====================================================

import boto3  
session=boto3.session.Session(profile\_name='XXXX',region\_name='us-east-1')  
ec2\_cli=session.client('ec2')  
all\_reg=ec2\_cli.describe\_regions()  
\#print(all\_reg\['Regions'\])  
list\_of\_regions=\[\]  
for each\_reg in all\_reg\['Regions'\]:  
 \#print(each\_reg\['RegionName'\])  
    list\_of\_regions.append(each\_reg\['RegionName'\])  
\#print(list\_of\_regions)  
for region in list\_of\_regions:  
 print('region:' , region)  
    ec2=boto3.client('ec2',region\_name =region)  
    response = ec2.describe\_snapshots(OwnerIds=\['self'\])  
 \#print(response\['Snapshots'\])  
 for snap in response\['Snapshots'\]:  
 \#print(snap\['SnapshotId'\])  
        snap.delete()
## [5][Cloudwatch logs for ECS in different region. Good or terrible idea?](https://www.reddit.com/r/aws/comments/h9d4tq/cloudwatch_logs_for_ecs_in_different_region_good/)
- url: https://www.reddit.com/r/aws/comments/h9d4tq/cloudwatch_logs_for_ecs_in_different_region_good/
---
Hi, 
We have our application deployed in the eu-west-1, now we're expanding in the US and I launched our stack in the us-east-1. However, I just found out that it's possible to send ecs logs in us-east-1 to the log group in eu. Do I need to pay data transfer fees for that? If not that sounds like a good idea because we also have a log analytics stack that is quite expensive to run.
## [6][Breaking in to the Top 10 of AWS Deepracer Competition - May 2020](https://www.reddit.com/r/aws/comments/h9g56l/breaking_in_to_the_top_10_of_aws_deepracer/)
- url: https://mickqg.github.io/DeepracerBlog/
---

## [7][Site-to-Site VPN, are you billed if the tunnel state is down?](https://www.reddit.com/r/aws/comments/h9crqw/sitetosite_vpn_are_you_billed_if_the_tunnel_state/)
- url: https://www.reddit.com/r/aws/comments/h9crqw/sitetosite_vpn_are_you_billed_if_the_tunnel_state/
---
As the title, will AWS bill me if the managed VPN tunnels are down?

As per here: [https://aws.amazon.com/vpn/pricing/](https://aws.amazon.com/vpn/pricing/) "You will be charged for your AWS Site-to-Site VPN connection on an hourly basis, for each hour the connection is active."

"Active" in this case is a little vague, but knowing the AWS pricing model, I feel strongly that they will be charging regardless of the state of the tunnels (i.e. the moment you create a AWS VPN in the console, regardless of if you set-it up on the other side), but wanted to run this past you beautiful people of r/aws
## [8][RedShift HealthCheck tool](https://www.reddit.com/r/aws/comments/h9buvm/redshift_healthcheck_tool/)
- url: https://www.reddit.com/r/aws/comments/h9buvm/redshift_healthcheck_tool/
---
We have open-sourced a small utility to perform health checks on the RedShift cluster. It'll check your

1. table-level issues
2. Design problems
3. Vacuum, stats
4. WLM
5. and Performance-related issues

We want the AWS Redshift community to provide feedback on this also we welcome your contribution.

Link to the repo:  [https://github.com/searceinc/RStoolKit](https://github.com/searceinc/RStoolKit)
## [9][Cognito Architectural Advice](https://www.reddit.com/r/aws/comments/h9bdoh/cognito_architectural_advice/)
- url: https://www.reddit.com/r/aws/comments/h9bdoh/cognito_architectural_advice/
---
 

I need some help with some architectural advice. I have a shared application load balancer (By shared I mean multiple customers use this and all requests are routed a set of EC2 instances running IIS)

I'm trying to implement SSO using cognito user pools with OpenID being the identity providers.

I want to be able to have my customers access use it by going to their URL

www.customerwebsite.com/sso and www.customerwebsite.com/admin/sso

I also want to create a backdoor so i'm also able to access it with my own OpenID, I'd call it something like www.CustomerWebsite.com/Backdoor the ( FYI the customers are okay this and i will be naming it something better than Backdoor :')  )

should I create a User Pool per customer and have my backdoor separate as a separate User Pool (therefore accounts won't appear in multiple user pools Or should i  create 1 user pool per customer with an additional App client for myself. I also worry about hitting the limit for Load Balncer Limit rule and input would be greatly appreciated
## [10][AWS/Cloud based distributed DAG task runner](https://www.reddit.com/r/aws/comments/h995od/awscloud_based_distributed_dag_task_runner/)
- url: https://www.reddit.com/r/aws/comments/h995od/awscloud_based_distributed_dag_task_runner/
---
Long time reader, first time poster. At work we have an internally developed distributed task runner system (Java) with a single main/delegator server and multiple task runner servers.

Jobs are a hierarchical tree of tasks that each have their own parameters and dependencies, parameters can be passed down into sub-tasks, tasks can spawn more sub-tasks that the main server redistributes to executors. Executors report status and logs back to the main server for quick examination of how a process ran.

It works really well but is showing its age a little. It was developed before Hadoop was a thing, before containers were mainstream and definitely before the AWS revolution. The problem is it assumes on-prem zero cost servers, so a dozen executor machines sit there idle until the one time a week/month they have to run their tasks. Worse there are multiple clusters and one cluster of five EC2 instances maxes out while several other ten-instance clusters sit idle. It's a crushing waste of resources.

My question is- is there an on-prem solution (or cloud solution I guess) for this situation where multiple clusters can join in on a process or a single larger cluster can be configured to work on whatever is needed if idle but snap to a certain task if it starts? Or alternatively an AWS or cloud solution where only processing time is charged?

I have looked into various task runner applications.

Apache Airflow (from AirBnB) is nice but needs serious bolted-on Celery job queue configuration to handle multiple executor servers.

Luigi (from Spotify) has a task dependency concept but I don't think it can farm out tasks to multiple servers. Also jobs are in Python which isn't a dealbreaker but would prefer declarative JSON, YAML, XML, even INI.

Nomad (from Hashicorp- vagrant, packer, etc.) is a really nice cluster manager, very professional UI and documentation but unfortunately doesn't support DAGs or a tree of tasks to perform with dependencies.

I looked in to a few more but it became clear that there wasn't a single offering that supported multiple task runner servers, tree-shaped (DAG) jobs

Am I missing something here? Running dependency-based high I/O distributed tasks seems like a thing every medium-sized company would do but there aren't any options that tick all the boxes.
## [11][Best way for secure access to S3?](https://www.reddit.com/r/aws/comments/h9736a/best_way_for_secure_access_to_s3/)
- url: https://www.reddit.com/r/aws/comments/h9736a/best_way_for_secure_access_to_s3/
---
What is the best way to allow **secure** access to data stored in S3 to other folks within your organization? Is it better to attach an identity policy to the users or to a resource policy to the S3 bucket and why?
