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
## [2][I've created a tool - spotcost.net The one-pager about all spot instances information. It helps find the cheapest region/az, compare specs, regions, price in time and etc. The difference between regions is huge (10-300% sic!). What do you think?](https://www.reddit.com/r/aws/comments/gwgdkm/ive_created_a_tool_spotcostnet_the_onepager_about/)
- url: https://spotcost.net
---

## [3][TLS 1.2 to become the minimum for all AWS FIPS endpoints | Amazon Web Services](https://www.reddit.com/r/aws/comments/gw1a47/tls_12_to_become_the_minimum_for_all_aws_fips/)
- url: https://aws.amazon.com/blogs/security/tls-1-2-to-become-the-minimum-for-all-aws-fips-endpoints/
---

## [4][Any reason NOT to use ACM Certificates ?](https://www.reddit.com/r/aws/comments/gw99u5/any_reason_not_to_use_acm_certificates/)
- url: https://www.reddit.com/r/aws/comments/gw99u5/any_reason_not_to_use_acm_certificates/
---
Is there any reason for someone to use traditional CA certificates (Comodo/Digicert etc) if they are fully on AWS, and do not plan to install the certificates in EC2 ? How would you convince corporate IT teams that traditional CA signed certificates offer no additional value;  and ACM in fact removes a lot of head aches offered by them like cert rotation?

Even if we decide to install certs on EC2, is there any advantage for paid certificates over those given by Lets Encrypt ?
## [5][RDS Script - Iterating over MS SQL databases and backing them up to s3](https://www.reddit.com/r/aws/comments/gwealg/rds_script_iterating_over_ms_sql_databases_and/)
- url: https://www.reddit.com/r/aws/comments/gwealg/rds_script_iterating_over_ms_sql_databases_and/
---
Hello Everybody,

I created this script a while ago and thought it might be useful for people here who use RDS for MS SQL and want to either migrate their databases or back them up.

&amp;#x200B;

This script iterates over all the databases except: master, model, rdsadmin, tmpdb, msdb.

&amp;#x200B;

prints the database name, and then backs it up as &lt;database\_name&gt;.bak to s3 (make sure to change the arn of the s3 bucket.)

&amp;#x200B;

 

`DECLARE @value VARCHAR(50)`

`DECLARE db_cursor CURSOR FOR`  

`SELECT name FROM master.dbo.sysdatabases`

`where name not in ('master','model','rdsadmin','tempdb','msdb')`

`OPEN db_cursor`   

`FETCH NEXT FROM db_cursor INTO @value`   

`WHILE @@FETCH_STATUS = 0`   

`BEGIN`   

`PRINT @value`

  `declare  @s3 nvarchar(MAX) = N'arn:aws:s3:::bucket_name/'+@value+'.BAK'`

  `exec msdb.dbo.rds_backup_database` 

`@source_db_name=@value,`

`@s3_arn_to_backup_to=@s3,` 

`@overwrite_S3_backup_file=1;`

`FETCH NEXT FROM db_cursor INTO @value`   

`END`   

`CLOSE db_cursor`   

`DEALLOCATE db_cursor`

&amp;#x200B;

&amp;#x200B;

Hopefully this helps someone!
## [6][Authorising HLS streaming files from Amazon S3 directory](https://www.reddit.com/r/aws/comments/gwglv0/authorising_hls_streaming_files_from_amazon_s3/)
- url: https://www.reddit.com/r/aws/comments/gwglv0/authorising_hls_streaming_files_from_amazon_s3/
---
We have setup that converts raw videos into HLS format (.m3u8 and .ts files) and organises them into a directory inside a s3 bucket. Each directory inside the bucket represents one video. Since s3 doesn't really have the concept of directory in its implementation, it does not allow us to get a signed url to read the content of the directory to feed into the video player.

I tried signing the URL for the .m3u8 file alone with getObject, but since tries to fetch the parts of the video to play, it will be thrown with an 403 by s3. Using cloudfront is not an option for us at this stage.

Is there a better and secure way to handle the streaming from s3 without making the entire bucket public?
## [7][API entrypoint with ACM certificates without ALB](https://www.reddit.com/r/aws/comments/gwh45t/api_entrypoint_with_acm_certificates_without_alb/)
- url: https://www.reddit.com/r/aws/comments/gwh45t/api_entrypoint_with_acm_certificates_without_alb/
---
 Hey!

I love to use ALB in front of EC2 instances because then I don’t need to worry about SSL on EC2 instances (on nginx side), but rather use ACM on ALB for 2 of my domains (site has to work on both [example1.com](http://example1.com/) and [example2.com](http://example2.com/), 2 certificates - [example2.com](http://example2.com/) is a CNAME to [example1.com](http://example1.com/), which right now is a CNAME to load balancer address). But is there a way of doing this when there’s only 1 EC2 instance and I don’t need ALB, but I still don’t want to manage certificates on the EC2 side?

The reason I’m asking is that I sometimes just need 1 instance, but sometimes I need to scale and need to add more EC2 instances and load balance them. The times I only need 1, I would rather not use the ALB to save money. Can I provision any sort of AWS resource that would serve as an entrypoint and would pass that request to the EC2 instance? I noticed CloudFront can be used to serve dynamic content as well, but it only supports 1 ACM certificate. Can API gateway be configured in such a way to pass the request down to EC2 instance but provide multiple certificates as well?  
 

I would like to be as flexible as possible, where I would CNAME my domains to some endpoint that stands in front of EC2, then if I need more instances, I would provision ALB, add those instances and swap out the CNAME to ALB’s address. I would love to hear out what you guys think.

All of this now works with ALB. I can just remove and add targeted instances, but in case I only need 1 instance, ALB is wasting money for my use case.

&amp;#x200B;

I'm not very experienced in AWS, so I would love to hear your opinions.
## [8][Domain still going to Godaddy website after setting up NS records with A record and CNAME records on AWS Lightsail](https://www.reddit.com/r/aws/comments/gwh3l1/domain_still_going_to_godaddy_website_after/)
- url: https://www.reddit.com/r/aws/comments/gwh3l1/domain_still_going_to_godaddy_website_after/
---
I purchased my domain name at Godaddy, then changed the nameservers in Godaddy's DNS settings over to my AWS nameservers. AWS now controls my DNS records so I added an A record and two CNAME records. 

The A record points to my static IP (@.example.com --&gt; staticIP) and my CNAME records point to my domain name ([www.example.com](https://www.example.com) \--&gt; [example.com](https://example.com))

Both my subdomains from the CNAME records take me to the correct website hosted on Lightsail, but going to the main domain ([example.com](https://example.com)) still takes me to Godaddy's website builder that I set up when I bought the domain. There are no other options in Godaddy's DNS settings except to transfer the domain which isn't available for another month. I'm not sure what else to do on the AWS side because I already set up the A record, but it's not working correctly. Anyone know what else I can do?
## [9][Different Deployment Types Of AWS Elastic Beanstalk](https://www.reddit.com/r/aws/comments/gwgzy3/different_deployment_types_of_aws_elastic/)
- url: https://www.ibexlabs.com/different-deployment-types-of-aws-elastic-beanstalk/
---

## [10][Automating RDS Start and Stop](https://www.reddit.com/r/aws/comments/gw7doh/automating_rds_start_and_stop/)
- url: https://www.reddit.com/r/aws/comments/gw7doh/automating_rds_start_and_stop/
---
I am using it (PostgreSql) only for developing purposes. So I want it to be open only on work hours. Is there a way of scheduling it?
## [11][Unable to map API Gateway to custom domain name](https://www.reddit.com/r/aws/comments/gwfprg/unable_to_map_api_gateway_to_custom_domain_name/)
- url: https://www.reddit.com/r/aws/comments/gwfprg/unable_to_map_api_gateway_to_custom_domain_name/
---
I am trying to create a custom domain name mapping to API Gateway. I followed the documentation and I have following things in place:

* SSL Certificate from ACM
* Entry for desired domain in Route 53 with routing set to API End point

Now only step remaining is adding API mapping. When I try to add an API mapping it throws errors saying  `Unable to complete operation due to concurrent modification. Please try again later.` 

Not able to figure out how to resolve this. Looking for assistance for the same.
