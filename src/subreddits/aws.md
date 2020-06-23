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
## [2][Introducing AWS Solutions Constructs](https://www.reddit.com/r/aws/comments/he38t0/introducing_aws_solutions_constructs/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/06/introducing-aws-solutions-constructs/
---

## [3][websocket for publically available chatbot](https://www.reddit.com/r/aws/comments/hebnhe/websocket_for_publically_available_chatbot/)
- url: https://www.reddit.com/r/aws/comments/hebnhe/websocket_for_publically_available_chatbot/
---
I'm developing a chatbot for a publically available website (kind of like a customer support chatbot) and I was wondering how to properly protect it while still using it without authentication.

I was thinking about using an invisible recaptcha or something along those lines but from previous projects I remember we had an issue with stealing sessions, but that was on Azure. Does anyone if API Gateway has something to handle that?
## [4][Multiple security groups or just one?](https://www.reddit.com/r/aws/comments/hee44q/multiple_security_groups_or_just_one/)
- url: https://www.reddit.com/r/aws/comments/hee44q/multiple_security_groups_or_just_one/
---
I'm wondering what everyone thinks on this topic. I personally like multiple security groups but I'm not sure if that's proper.

In regards to EC2, I feel, if you have rules that apply to more than 1 instance; say allowing RDP/SSH from the office IP, then you make one SG for that and apply to all instances. However, one issue is that for windows, you dont need port 22 open by default. This means your attack surface is larger than needed, even if just from a single IP. 

I like applying one SG to multiple instances because it allows for easy updates across all instances if required. For example, Office IP changes for whatever reason, you just update one SG rather than 100s or 1000s.

Also, I like having multiple, more targeted SGs on instances that need custom rules. For example, one SG handling HTTP/HTTPS and another handling say port 22 for FTP etc. Keeping the 5 SG limit in mind.

I've seen where you give each instance it's own SG and put absolutely everything in that and it just seems super messy and a bit of management nightmare.

Thoughts?
## [5][AWS Solutions Constructs – A Library of Architecture Patterns for the AWS CDK](https://www.reddit.com/r/aws/comments/he1g74/aws_solutions_constructs_a_library_of/)
- url: https://aws.amazon.com/blogs/aws/aws-solutions-constructs-a-library-of-architecture-patterns-for-the-aws-cdk/
---

## [6][Athena "SELECT * not allowed in queries without FROM clause"/"Column 'column_name' cannot be resolved"](https://www.reddit.com/r/aws/comments/hedmgi/athena_select_not_allowed_in_queries_without_from/)
- url: https://www.reddit.com/r/aws/comments/hedmgi/athena_select_not_allowed_in_queries_without_from/
---
Hi all!

I have data, which comes from webhook as JSON string. I save it as JSON document to S3 using Lambda (Python). Then I use Glue Crawler to create schema and when I want to take a peek at my data, Athena gives me `SELECT * not allowed in queries without FROM clause` error when I try to query all columns and  `Column 'column_name' cannot be resolved` when I want to query specific column.

I searched for similar error and everyone has it due to their CSV encoding, but no answers with JSON files. Does anyone know what can possibly be wrong?  


**UPD**: Okay, nvm, all I had to do is to grant myself a permissions at Lake Formation on specified tables
## [7][s3 bucket permissions](https://www.reddit.com/r/aws/comments/hecpwp/s3_bucket_permissions/)
- url: https://www.reddit.com/r/aws/comments/hecpwp/s3_bucket_permissions/
---
Hi all

We currently use cyberduck to perform some Windows file uploads to s3

Does anyone know if its possible to create a IAM User + Policy to provide the ability to  move/rename/delete files and folders within a specified bucket? Or is this simply not available/allowed ?
## [8][What is/are your DR plan?](https://www.reddit.com/r/aws/comments/he3v6v/what_isare_your_dr_plan/)
- url: https://www.reddit.com/r/aws/comments/he3v6v/what_isare_your_dr_plan/
---
Can you guys share yours? also where should i start? like reading resources.

initially i thought theres no need to have a DR but then came to my mind that i can deploy a standby on another region and just clone the servers

our aws are mostly rds mysql, elasticache and ec2
## [9][How to use EFS with AWS Lambda?](https://www.reddit.com/r/aws/comments/hdy144/how_to_use_efs_with_aws_lambda/)
- url: https://www.reddit.com/r/aws/comments/hdy144/how_to_use_efs_with_aws_lambda/
---
&amp;#x200B;

https://preview.redd.it/sv9how6a9i651.png?width=641&amp;format=png&amp;auto=webp&amp;s=d94df6c5bc1aafa9b5c1466c3472760a4d17156e

AWS recently launched a new feature that lets the customer make an EFS (Elastic File System) that works with Lambda. This is really cool! But, Why EFS?

EFS is a storage unit that lives independently inside a machine. This file can be attached to other services like EC2 and can be accessed from multiple instances at once. Files that are inside this storage unit is accessible from any connected instance or Lambda.

**Why do we need something like this to work with Lambda? Extra complicated step?**

Actually, this feature is amazing if you looked at it from different angles, let’s start with some of them:

*1- Consistency:*

If you need multiple Lambdas to use (read and write) BIG files, you’ll need them to be in a place that doesn’t delay the function to get the resources, which leads to less computing power and time -&gt; less money.

*2- More space:*

When you’re working with files from S3, you’re limited to the max storage size of 512 MiB. Which is not enough is some cases. Plus, you might need to work with this file in different processes stages, like cleaning, segmenting, processing, and exporting readable reports/formats of this file. Imagine the amount of code involved in this scenario.

*3- More space 2:*

Using layers will share the resources between the functions, but Layers sometimes can’t handle the size of the resources and binaries you called to run this function. Using EFS will give you more room to store these resources and call when is needed.

I can list more points. But, you got the idea. Let’s dive into how to use it with lambda.

# Creating EFS:

1- Open AWS console and search for “EFS”.

&amp;#x200B;

https://preview.redd.it/rh3jgvmg9i651.png?width=2876&amp;format=png&amp;auto=webp&amp;s=39732c98655bc1f1509fe9756ec9c438c2f96cf5

2- Click on “Create file system”.

3- At step 1, select your VPC. If your lambda is configured within a VPC, choose it, if not, remember when VPC you’ve selected.

&amp;#x200B;

https://preview.redd.it/hoycogli9i651.png?width=2880&amp;format=png&amp;auto=webp&amp;s=8a3f9bdad9fd0b8d618406e4bad749ee6941a7ae

4- Step 2, Add descriptive name for your file system. Then click next step.

&amp;#x200B;

https://preview.redd.it/osowihal9i651.png?width=2880&amp;format=png&amp;auto=webp&amp;s=da1ec399e14736501c2ac5ca2f1932253fcae9ac

5- In step 3, go down and click on “Add access point”. Then fill it with what’s in the image.

&amp;#x200B;

https://preview.redd.it/hgggcc6o9i651.png?width=2238&amp;format=png&amp;auto=webp&amp;s=8a6a96daac64826bc0ab721ee9983aba4b12ce32

6- Review the configuration, then click “Create file system”.

&amp;#x200B;

https://preview.redd.it/e0g20ilq9i651.png?width=2878&amp;format=png&amp;auto=webp&amp;s=c0e8b7c1b97ae4374c82da702ac876e72318b4ba

7- Done ! wait for a few seconds and your EFS will be active.

&amp;#x200B;

https://preview.redd.it/d5ggxlks9i651.png?width=2874&amp;format=png&amp;auto=webp&amp;s=0b90bf69010a1195eea178b81f16c300a0eda84f

# Connect it with Lambda:

1- Click on Service and search for “Lambda”.

&amp;#x200B;

https://preview.redd.it/4487tf4u9i651.png?width=2880&amp;format=png&amp;auto=webp&amp;s=cfeebe0297dadd38183d01e37fc355d2aa949a88

2- Create a new function and choose your preferred runtime language. In this article, I’ll use Python.

&amp;#x200B;

https://preview.redd.it/bpf99wtv9i651.png?width=2616&amp;format=png&amp;auto=webp&amp;s=448c0c3e7d13038a8c38298e90ed6c7dac265021

3- Go down a little bit and you’ll see a section called “VPC”, click on “Edit”.

&amp;#x200B;

https://preview.redd.it/8tizq57x9i651.png?width=2600&amp;format=png&amp;auto=webp&amp;s=0947937eff2593757ea236d0169a0f187382b56e

4- Select your VPC, and choose the Subnets and the security group. Then save.

&amp;#x200B;

https://preview.redd.it/251at16z9i651.png?width=2880&amp;format=png&amp;auto=webp&amp;s=8ef40c9544bfb57286b37ffbdbde54f52be6cff9

5- Under the VPC section, click on “Add file system” from “file System” section.

&amp;#x200B;

https://preview.redd.it/hy0uu0eeai651.png?width=2588&amp;format=png&amp;auto=webp&amp;s=9c52f8b67bb6734815fd076537575fdcd2249fd4

6- Select the EFS File System we made, remember that we gave it a descriptive name. Then choose the Access point that is associated with the Access Point and finally, give it a path.

&amp;#x200B;

&gt;*NOTE: This path needs to start with ‘/mnt/‘. You can keep it as it is or if you want to have a custom folders, defiantly you can.*

&amp;#x200B;

https://preview.redd.it/qi9rm8mgai651.png?width=1656&amp;format=png&amp;auto=webp&amp;s=b5ecbf64d578261479a131b73e8119865a5057fc

7- A small piece of code to test if the file system is attached to the function.

&amp;#x200B;

https://preview.redd.it/k402xs7jai651.png?width=2666&amp;format=png&amp;auto=webp&amp;s=53b6b731d5a26176bf96deccbc68b30e8827ac5d

8- Bingo!! We made it!

&amp;#x200B;

https://preview.redd.it/watgdwpkai651.png?width=2578&amp;format=png&amp;auto=webp&amp;s=753ee8eed2eab80384cb55c658803550490efaef

# Conclusion:

Adding EFS to Lambda is a huge new milestone in The Serverless architecture. You can have new use-cases that will be doable and before were a nightmare to accomplish. Easy steps with the cheap price make Lambda an option that can battle EC2 in some new modern use-cases.
## [10][AWS EKS multi region service discovery](https://www.reddit.com/r/aws/comments/he4x4w/aws_eks_multi_region_service_discovery/)
- url: https://www.reddit.com/r/aws/comments/he4x4w/aws_eks_multi_region_service_discovery/
---
Hi Everyone, 

I have a scenario where I have EKS cluster in multiple AWS regions. ALL the VPC's are paired using TGW to route the traffic between them. 

For DNS resolution between the region for k8s to discover pod in other region looks like we can use AWS R53 resolver. But turns out we can also CoreDNS instead of R53 resolver. 

Anyone have used CoreDNS for service discovery between the EKS cluster in multiple AWS regions? 

It will be super helpful if you have any lesson learned that you can share?

&amp;#x200B;

Cheers.
## [11][AWS WorkSpaces resilience - recovery from an AZ failure?](https://www.reddit.com/r/aws/comments/hdzl3z/aws_workspaces_resilience_recovery_from_an_az/)
- url: https://www.reddit.com/r/aws/comments/hdzl3z/aws_workspaces_resilience_recovery_from_an_az/
---
We're looking into putting our desktop in the cloud by using WorkSpaces as a VDI solution.  I believe there are a lot of problems, but one which stands out is high availability.

How does WorkSpaces cope with a AZ failure?  Their 'Best Practices for Deploying Amazon Workspaces' whitepaper ([https://d1.awsstatic.com/whitepapers/workspaces/Best\_Practices\_for\_Deploying\_Amazon\_WorkSpaces.pdf](https://d1.awsstatic.com/whitepapers/workspaces/Best_Practices_for_Deploying_Amazon_WorkSpaces.pdf)) shows deploying over two AZs, but what if one fails?

I think WorkSpaces are balanced across the two, with the backups on S3, so losing an AZ would mean restoring half the WorkSpaces to the remaining AZ.  Sounds horrible!  Or do the WorkSpaces exist in both AZ's?

cheers.
