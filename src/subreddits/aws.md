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
## [2][CPU showing vulnerability on boot that was supposed to be fixed on EC2 instances last year](https://www.reddit.com/r/aws/comments/gvhpun/cpu_showing_vulnerability_on_boot_that_was/)
- url: https://www.reddit.com/r/aws/comments/gvhpun/cpu_showing_vulnerability_on_boot_that_was/
---
Hello all.  I have the following issue, it started with the `oom_reaper` killing the only memory intensive process on the server, `mariadb`, as seen when looking at oom_scores.  So I'm not positive, maridb is actually causing it to run out of memory.    


In debugging the issue i was looking through the `kern.log` and see the following.
```
MDS CPU bug present and SMT on, data leak possible. See https://www.kernel.org/doc/html/latest/admin-guide/hw-vuln/mds.html for more details.
TAA CPU bug present and SMT on, data leak possible. See https://www.kernel.org/doc/html/latest/admin-guide/hw-vuln/tsx_async_abort.html for more details.
```
And i have gone over those links which lead me to CVEs addressed by Amazon in this release. https://aws.amazon.com/security/security-bulletins/AWS-2019-004/.  

The values on all of my servers in two VPCs in the file `cat /sys/devices/system/cpu/vulnerabilities/mds` is `Vulnerable: Clear CPU buffers attempted, no microcode; SMT Host state unknown`.  

From my understanding that statement is there b/c `md_clean` is not passed to kernel. So i installed `intel-microcode` package, i can see the available and chosen ones with with `/usr/sbin/iucode_tool -tb -lS /lib/firmware/intel-ucode/*`

When i reboot the server the message don't change and `dmesg | grep microcode` returns 
```
[    0.820544] TAA: Vulnerable: Clear CPU buffers attempted, no microcode
[    0.820544] MDS: Vulnerable: Clear CPU buffers attempted, no microcode
```

Now I may be going down a rabbit hole with my OOM, but this was a clear well documented issue and well hell at this point i don't know exactly what i want, but i'll take any input. 

I'm confused about why it doesn't actually show the chosen microcodes or change the message. All updates are applied to all servers.

```
$&gt; uname -a
Linux prodweb01 5.3.0-1019-aws #21~18.04.1-Ubuntu SMP Mon May 11 12:33:03 UTC 2020 x86_64 x86_64 x86_64 GNU/Linux
```
The impacted server is a `c5.xlarge`  I see this on older T2 instances and newer rebuilt T3 instances.

Thanks for any ideas, speculations etc.
## [3][Cutting costs - Do I need a NAT gateway?](https://www.reddit.com/r/aws/comments/gvqvjg/cutting_costs_do_i_need_a_nat_gateway/)
- url: https://www.reddit.com/r/aws/comments/gvqvjg/cutting_costs_do_i_need_a_nat_gateway/
---
I'm running a small application in AWS. I have a web server that provides an API, hosted in AWS Fargate, with an Elastic Load Balancer in front of it. It connects to a RDS instance to store user data. I have a SPA client hosted in S3 with CloudFront.

1. I set the stack up with AWS CDK and it apparently spun up a NAT Gateway for me which is one of the greatest costs. Do I even need that?
2. At the moment I'm at roughly 50 users per day, so I guess I should be fine with the lowest tier CPU/memory specs on all my instances, right? The application is basically a CRUD application at this point.
3. I don't want to spend too much cash on this, but my costs are already estimated to $100+/month. Is this reasonable?
## [4][AWS WorkSpaces - How do you protect / monitor it?](https://www.reddit.com/r/aws/comments/gvteo6/aws_workspaces_how_do_you_protect_monitor_it/)
- url: https://www.reddit.com/r/aws/comments/gvteo6/aws_workspaces_how_do_you_protect_monitor_it/
---
I'm trying to create a proof of concept for whether a company could use AWS WorkSpaces as a DaaS. I keep finding that a lot of services I try to use only work with EC2 instances or need IAM roles.

* CloudWatch - have to treat WorkSpaces as an on-premise device to get custom metrics and use an agent which was a pain to setup. Took 2 minutes to get working with an EC2 instance however took me a week to get it working properly with a WorkSpace
* Amazon Inspector - only EC2, I can't run it at all against my WorkSpaces to see how vulnerable they are.
* Guard Duty - I haven't use this yet however it looks like it needs a service role, since you can't put service roles on WorkSpaces I'm assuming this won't work but may be wrong..

Anyone got any advice maybe I'm missing something obvious? Thank you very much.
## [5][RDS PostgreSQL Logical Replication COPY from AWS RDS Snapshot](https://www.reddit.com/r/aws/comments/gvsqyd/rds_postgresql_logical_replication_copy_from_aws/)
- url: https://medium.com/searce/rds-postgresql-logical-replication-copy-from-aws-rds-snapshot-6983446472a9
---

## [6][Updated an RDS instance tag via CloudFormation, and the instance rebooted](https://www.reddit.com/r/aws/comments/gvsnzf/updated_an_rds_instance_tag_via_cloudformation/)
- url: https://www.reddit.com/r/aws/comments/gvsnzf/updated_an_rds_instance_tag_via_cloudformation/
---
I have a CloudFormation stack of an Aurora PostgreSQL cluster with currently a single instance.  There are some other supplementary resources such as ParameterGroup, SecurityGroup, etc.

Today I initiated a stack update which did one single thing: it changed a single tag value that is applied to all the resources.  That's it.

My stack changeset showed nothing unusual, no unexpected Replacements, etc.

I execute the update, and things when it gets to updating the instance, I notice in the window behind my current one, the DB instance rebooted.

I have two questions:

1) why?

2) in the future is there any way I can anticipate disruptions like this?
## [7][An abnormally high number of S3 GET requests on a development account?](https://www.reddit.com/r/aws/comments/gvs9v4/an_abnormally_high_number_of_s3_get_requests_on_a/)
- url: https://www.reddit.com/r/aws/comments/gvs9v4/an_abnormally_high_number_of_s3_get_requests_on_a/
---
Hi all,

I was just taking a look at my billing when I noticed a high number of S3 GET requests (around 4000). I know this isn't huge, but it is not a public or project-tied account. I have only been doing some development work using CloudFormation and SAM. 

I checked again this morning and now there are \~5700 requests in total. I thought that originally my GET requests might be coming from CodePipeline as part of the CI solution I'm using. 

Has anyone seen this before?

UPDATE: I forgot to submit this post, and since re-checking there are now over 6500 requests.
## [8][VPN Solution](https://www.reddit.com/r/aws/comments/gvm8sm/vpn_solution/)
- url: https://www.reddit.com/r/aws/comments/gvm8sm/vpn_solution/
---
Looking for a robust remote access vpn solution. We’ve been using Sophos UTM’s SSL vpn client but it’s pretty unreliable where it sometimes refuses to resolve dns names until client machines are rebooted. Are there any AWS recommended firewall appliances that have good vpn support or perhaps stand alone vpn? Looked into native AWS client vpn but looks too new to consider and a bit cumbersome.
## [9][Whitelisting paths in AWS WAF](https://www.reddit.com/r/aws/comments/gvqw24/whitelisting_paths_in_aws_waf/)
- url: https://www.reddit.com/r/aws/comments/gvqw24/whitelisting_paths_in_aws_waf/
---
Hello, I'm a student which is creating  an IaC infrastructure  (with Terraform) on AWS for a launching company. 

We have setup the waf with the Owasp top 10 rules (using a terraform module) in front of our load balancer. The problem I am facing here is I need to whitelist certain URI (like /api/comment) to allow some blocked words (like the SQL injection words, they can be use by our visitors in the comment section) and I don't understand how I am supposed to do that.  
Can some of you guys help me here ? thanks in advance
## [10][Using the built in ALB authentication](https://www.reddit.com/r/aws/comments/gvllv4/using_the_built_in_alb_authentication/)
- url: https://www.reddit.com/r/aws/comments/gvllv4/using_the_built_in_alb_authentication/
---
When using this:  [https://aws.amazon.com/blogs/aws/built-in-authentication-in-alb/](https://aws.amazon.com/blogs/aws/built-in-authentication-in-alb/) 

since the application is no longer handling the authentication, how do you know who the authenticated user is in the app?  Does it pass through the JWT token?
## [11][Internal Documentation Tool](https://www.reddit.com/r/aws/comments/gv7gox/internal_documentation_tool/)
- url: https://www.reddit.com/r/aws/comments/gv7gox/internal_documentation_tool/
---
Curious what everyone else uses for enterprise internal documentation. Wiki? Sharepoint? Build a custom tool in AWS?  
For example, what would you use if you want to share architecture diagrams, run books, ops books, etc. to anyone in your IT org?
