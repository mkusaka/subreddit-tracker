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
## [2][Building a podcast search engine using all AWS services.](https://www.reddit.com/r/aws/comments/hb9yue/building_a_podcast_search_engine_using_all_aws/)
- url: https://www.reddit.com/r/aws/comments/hb9yue/building_a_podcast_search_engine_using_all_aws/
---
I built a service that tracks a podcast, transcribes its episodes, and indexes it for search. I have a sample prototype rigged for the Joe Rogan Experience at [www.jamiepullup.com](https://www.jamiepullup.com). 

The service is a swift console app, all the processing, transcribing and indexing happens on various AWS services (EC2 to host the site and run the service, S3 for storage, Amazon Transcribe for transcription and Amazon Cloud Search for search)

Getting this up and running wasn't cheap, thank god for the activate program. 

If you have any questions regarding any of these technologies, I'd be more than happy to answer.
## [3][AWS said it mitigated a 2.3 Tbps DDoS attack, the largest ever](https://www.reddit.com/r/aws/comments/hau6np/aws_said_it_mitigated_a_23_tbps_ddos_attack_the/)
- url: https://www.zdnet.com/article/aws-said-it-mitigated-a-2-3-tbps-ddos-attack-the-largest-ever/
---

## [4][Can I reduce my cost by changing the way I'm running my infrastructure?](https://www.reddit.com/r/aws/comments/hbc5dd/can_i_reduce_my_cost_by_changing_the_way_im/)
- url: https://www.reddit.com/r/aws/comments/hbc5dd/can_i_reduce_my_cost_by_changing_the_way_im/
---
Currently I'm running 2 EC2 instances.

The main one is a Windows Server instance on t2.small with a 100GB gp2 storage volume which runs IIS web server, MSSQL database server, FTP server and hMailServer for email. 

The other is a Linux server running on t2.micro with the default 8GB gp2 storage which basically just runs a LAMP stack (no email) but is hardly used.

My average monthly bill over the last 2 years has been about $40/month but I'm wondering if there are other tools on AWS that would allow me to still host and maintain my web apps but possibly at a cheaper rate than using the EC2 instances mentioned.

Is there some sort of app hosting feature that I can use to host .NET web apps, databases, etc. to maintain my hosting environment or is EC2 the best I can get? In case EC2 is the best, I may have to look for cheaper alternatives to AWS which I really don't want to do because moving is going to be a pain.

Thanks in advance!
## [5][Have reserved CIDR blocks changed in the last year to eighteen months or so?](https://www.reddit.com/r/aws/comments/hbdbji/have_reserved_cidr_blocks_changed_in_the_last/)
- url: https://www.reddit.com/r/aws/comments/hbdbji/have_reserved_cidr_blocks_changed_in_the_last/
---
I just tried to add a new CIDR block to my VPC and stumbled upon this error.

As you can see, I already have 172.31.1.0/28 in my VPC which I added probably around 12-18 months ago, but according to the [documentation](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html#VPC_Sizing), this is a reserved address and isn't allowed.

Am I going crazy or has this changed recently?

https://preview.redd.it/mhlucmrqrn551.png?width=1718&amp;format=png&amp;auto=webp&amp;s=c281301322b5603bb9b835930623be046ce64c71

EDIT: Further to this, it looks like it's not accepting any subnets, even ones that are not restricted. I must be doing something wrong here.
## [6][SAM / Cloudformation API -&gt; Kinesis woes](https://www.reddit.com/r/aws/comments/hbcvm8/sam_cloudformation_api_kinesis_woes/)
- url: https://www.reddit.com/r/aws/comments/hbcvm8/sam_cloudformation_api_kinesis_woes/
---
Hi all,

Can anyone confirm whether or not I can establish an API gateway via cloudformation to receive HTTP posts and have it forward the payload directly to a kinesis data stream?

It appears to me I can only do that if I'm using http/2 / WebSocket gateway,  OR I intend to forward to a firehose stream which neither is not what I'm wanting.

When I attempt to deploy the stack, it's failing because I've got integration type set to AWS\_PROXY, and it claims it only supports lambdafunctions or firehose...  

Seems odd that if it really can't pass the request payload directly through to a data stream...?

Thanks
## [7][AWS Snowcone – A Small, Lightweight, Rugged, Secure Edge Computing, Edge Storage, and Data Transfer Device](https://www.reddit.com/r/aws/comments/hav7qe/aws_snowcone_a_small_lightweight_rugged_secure/)
- url: https://aws.amazon.com/blogs/aws/introducing-aws-snowcone-small-lightweight-edge-storage-and-processing/
---

## [8][Request for advice on which AWS service would best suit my purpose.](https://www.reddit.com/r/aws/comments/hbaez1/request_for_advice_on_which_aws_service_would/)
- url: https://www.reddit.com/r/aws/comments/hbaez1/request_for_advice_on_which_aws_service_would/
---
Hello, I would like to request some advice on which of the AWS NLP  related services would best suit my purpose as several of them seem like  they could be suitable candidates and there seems to be some overlap in  what they could be used for.  
 

As a disclaimer I will say that I am fairly new to all this stuff so please forgive any ignorance or naivety on my part.   
 

I have been tasked by the company I work for to find a solution for  detecting questions that have the same meaning but are worded  differently.  
 

For context, we develop apps for Google Home, Alexa, and things like that.  
 

Unfortunately I do not yet have all the details of the internal  mechanics of the application that this question pertains to but my  educated guess is that we will have a set of known questions and answers  and we will need a way to map a question asked by the user to one of  the existing known questions even if it is worded differently.  
 

I have done some reading and watched some videos the past few days and have been able to identify some possible candidates:  
 

&amp;#x200B;

* The first is using some pre-trained BERT models with Sagemaker and  Tensorflow/Pytorch, I still only have a vague notion of how this would  apply to what I'm trying to do but I will gladly learn all the specifics  if this turns out to be the best solution for my needs. I understand  that this system can be used to develop chatbots, and while what I am  trying to do is similar to a chatbot, I think it is less sophisticated  in the sense that we only need our application to be able to understand a  very specific set of questions whereas a typical chatbot has a much  broader scope of interaction (I may be misinterpreting this). I have a  feeling this solution may be overkill for my current purposes.  


&amp;#x200B;

* The second possible solution is Comprehend. Again, seems like it could be overkill.  


&amp;#x200B;

* The third possible solution is Lex. This seems like it could work,  but at the same time it also seems like it could be too specifically  oriented towards typical chatbot development and not flexible enough for  my purpose.  


Anyway, sorry for the long winded post.  
 

Any advice would be greatly appreciated.
## [9][Restricting TLS version when accessing S3 over the internet - is Cloudfront the only answer?](https://www.reddit.com/r/aws/comments/hbadu0/restricting_tls_version_when_accessing_s3_over/)
- url: https://www.reddit.com/r/aws/comments/hbadu0/restricting_tls_version_when_accessing_s3_over/
---
Hi,

I have a private S3 bucket that I need to grant access to a third party, who will download data out of it.

The current setup is they have an AWS account and programmatic  user that has explicit permissions granted to a role that can only access that bucket.  They then use one of the AWS SDK's to download the files to their servers over the internet.

The bucket policy is setup to deny non-SSL traffic and to restrict access to their IP ranges.

I have now been asked to ensure that this access is restricted to TLS V1.2 or greater.

While it is possible to do this via CloudFront, I am a bit unsure how this would fit with the above scenario without a lot of rework.

Does anyone know of any other way to restrict TLS version?

Thanks
## [10][Does Aurora Serverless support Multi-AZ by default ?](https://www.reddit.com/r/aws/comments/hbe8is/does_aurora_serverless_support_multiaz_by_default/)
- url: https://www.reddit.com/r/aws/comments/hbe8is/does_aurora_serverless_support_multiaz_by_default/
---
Hi,

When choosing to create a new RDS Aurora **Serverless** (Mysql or Postgres) there's no option to select **Multi-AZ**.

Does the service support Multi-AZ by default ?

In case of an AZ failure does Aurora Serverless stay running ?

Thank you.
## [11][EC2 instance can no longer call AWS ec2 stop-instances](https://www.reddit.com/r/aws/comments/hbax8r/ec2_instance_can_no_longer_call_aws_ec2/)
- url: https://www.reddit.com/r/aws/comments/hbax8r/ec2_instance_can_no_longer_call_aws_ec2/
---
I have a ec2 instance with the following rights on a instance policy

ec2:StopInstances
ec2:StartInstances
ec2:DescribeTags

Etc...

The instance (using amzn2-ami-hvm-2.0.20200520.1-x86_64-gp2) runs a go application within docker for performing certain actions. One of these actions include stopping/starting instances in a managed way within the same VPC. It use to work and the application was able to stop instances, but with the last few weeks I am now getting “UnauthorizedOperation” error. I even tried connecting to the server using ssh and performing a aws ec2 stop-instances command. The other command that I can no longer execute is ec2 describe-tags (again, not authorised). 

I am stuck. I do not know why I can no longer perform these actions. My next step is to revert to an older AMI (though that surely should have no effect). 

Any help, ideas?
