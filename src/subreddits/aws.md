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
## [2][Looking to stay up to date on the latest #AWS Public Roadmaps? Here is a a Github repo which lists them so that you can find them in one place and stay updated.](https://www.reddit.com/r/aws/comments/hicf1p/looking_to_stay_up_to_date_on_the_latest_aws/)
- url: https://github.com/johnculkin/UnofficialListOfPublicAWSRoadmaps
---

## [3][Elemental Video on Demand diagram](https://www.reddit.com/r/aws/comments/himckb/elemental_video_on_demand_diagram/)
- url: https://www.reddit.com/r/aws/comments/himckb/elemental_video_on_demand_diagram/
---
As an exercise, I decided to make a diagram of AWS's [Video on Demand](https://aws.amazon.com/solutions/implementations/video-on-demand-on-aws/) solution. I did this in part because I felt the diagram they provided in the docs wasn't anywhere near detailed enough.

Diagram Link:  [**https://app.ilograph.com/demo.ilograph.AWS%2520Video-On-Demand/Ingest%2520Workflow**](https://app.ilograph.com/demo.ilograph.AWS%2520Video-On-Demand/Ingest%2520Workflow)

A few thoughts I had while working on this:

* As you might expect, ingesting video it is a complex process (60 + steps), even with AWS Elemental doing the heavy lifting.
* The example solution provided by AWS makes some interesting choices, such as re-using a lambda to kick off one of three different state machines depending on the input. This lambda is also invoked three different ways (S3 trigger, CloudWatch trigger, and invoked by another lambda). Hopefully this is conveyed clearly in the diagram above.
* Some of the names the solution gives the resources are... not great. Examples include "Endpoint" (a MediaConvert endpoint), "DynamoDB Table", and "Sns Topic". As far as I can tell, there is no easy way to change these in the solution before deploying (or after, for that matter).
## [4][Your own AI-powered Call Center on AWS in less than 1 hour](https://www.reddit.com/r/aws/comments/hinf9j/your_own_aipowered_call_center_on_aws_in_less/)
- url: https://www.reddit.com/r/aws/comments/hinf9j/your_own_aipowered_call_center_on_aws_in_less/
---
Amazon Connect has always been a curiosity of mine, and I briefly worked in a startup whose primary focus was to make Call Center’s life easier for its operators. During that time, I went to Manilla to shadow one of those places, and the amount of energy, effort, and patience to run such a business is fantastic!

[https://medium.com/@thiagodefaria/your-own-ai-powered-call-center-on-aws-in-less-than-1-hour-d0494049977c](https://medium.com/@thiagodefaria/your-own-ai-powered-call-center-on-aws-in-less-than-1-hour-d0494049977c)
## [5][Amazon Virtual Private Cloud (VPC) customers can now use their own Prefix Lists to simplify the configuration of security groups and route tables](https://www.reddit.com/r/aws/comments/hia8fb/amazon_virtual_private_cloud_vpc_customers_can/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/06/amazon-virtual-private-cloud-customers-use-prefix-lists-simplify-configuration-security-groups-route-tables/
---

## [6][Сloud Security Fundamentals: Infrastructure Hardening and Cloud Computin...](https://www.reddit.com/r/aws/comments/hi19mo/сloud_security_fundamentals_infrastructure/)
- url: https://www.reddit.com/r/aws/comments/hi19mo/сloud_security_fundamentals_infrastructure/
---
Learn cloud security fundamentals, explore popular benchmarks, learn how to protect different types of services, automate cloud compliance checks, and enhance operational security:  [**https://youtu.be/d52-fYISrLQ**](https://youtu.be/d52-fYISrLQ)  

**About the Webinar:**

During the Сloud Security Fundamentals webinar, DataArt’s experts delved into the aspects of shared responsibility in cloud security and compliance check automation. They reviewed the popular benchmarks such as TrendMicro and CloudSploit, shared the approaches to identity and access management hardening, and much more.

In this cloud security tutorial, get insight on:

— Key differences in security requirements for on-prem and cloud security architecture

— Your duties keeping the cloud resources safe as a part of responsibility shared with the cloud provider

— Cloud infrastructure hardening methods and ultimate cloud security goals

— Popular security auditing benchmarks and tools, and how to adapt them to achieve your desired security level

— Installing software updates and vulnerability management after migrating to the cloud
## [7][Terraform / AWS &amp; Secrets](https://www.reddit.com/r/aws/comments/hik0vj/terraform_aws_secrets/)
- url: https://www.reddit.com/r/aws/comments/hik0vj/terraform_aws_secrets/
---
Hi all

How are you guys currently storing secrets created via Terraform (on AWS)?

&amp;#x200B;

Previously I have used Hashicorp Vault for this but I want to explore other options as well with a view to avoiding storing the secret in the terraform state (which is what the Vault solution successfully achieved).

Do you just accept that it will be in the terraform state and lock down accordingly or are you using some other solution?

Vault is very good, and I would recommend it, but I want to check out the other options too whilst I have some down time.

Thanks
## [8][another billing post](https://www.reddit.com/r/aws/comments/hijnqr/another_billing_post/)
- url: https://www.reddit.com/r/aws/comments/hijnqr/another_billing_post/
---
hi all,

i know this gets asked often but i'd like to see if anyone has made recent progress.

what i'd like to do is basically create stuff in aws, add tags, and have a single dashboard that shows everything with this tag show x amount dollars per time query. i have started using grafana because its flexible and powerful but i can't see any examples of people using tags to filter dashboards.

has anyone used another products to query aws billing data with some success? do i need to get the data into another service or db to then query it.

thanks.
## [9][How to stay sane when managing tens-to-hundreds of lambda-backed repositories.](https://www.reddit.com/r/aws/comments/hi0wkv/how_to_stay_sane_when_managing_tenstohundreds_of/)
- url: https://www.reddit.com/r/aws/comments/hi0wkv/how_to_stay_sane_when_managing_tenstohundreds_of/
---
If you already use or are considering adopting lambdas/kinesis for data processing, [I wrote an article](https://medium.com/whispering-data/tackling-fragmentation-in-serverless-data-pipelines-b4027f506ee5) on dealing with deploying tens-to-hundreds of lambda-based repos in a sane manner.
## [10][Permission boundary policy denying VPC "edit dns resolution" and "edit dns hostnames" how?](https://www.reddit.com/r/aws/comments/hiibcf/permission_boundary_policy_denying_vpc_edit_dns/)
- url: https://www.reddit.com/r/aws/comments/hiibcf/permission_boundary_policy_denying_vpc_edit_dns/
---
Hello fellow Cloud-fans

Im currently hardening our organizations primary VPC and a subtask here is to strengthen the Permission Boundary on the role that all regular users are using, and here occurs a problem for me. I need to deny the users access to edit  "dns resolution" and "dns hostnames" for this specific VPC, however, I see no IAM actions that cover these 2 actions, so how do I deny it without using a wildcard?  
   
For reference: Other hardening actions I've already implemented to protect this specific VPC(if you have suggestions for additions please do share):

* Block modification/Deletion of DHCP options for our specific VPC
   * ec2:DeleteDhcpOptions 
   * ec2:AssociateDhcpOptions 
* Block modification of CIDR blocks for our specific VPC
   * ec2:DisassociateVpcCidrBlock 
* Block deletion of our specific VPC
   * ec2:DeleteVpc 
* **Block "edit dns resolution" for our specific VPC?**
* **Block "edit dns hostnames" for our specific VPC**
* Block modification of tags for our specific VPC
   * ec2:DeleteTags
* Block modification of ACLs for our specific VPC
   * ec2:ReplaceNetworkAclAssociation
   * ec2:DeleteNetworkAclEntry 
   * ec2:CreateNetworkAclEntry 
* Block modification of RouteTable for our specific VPC
   * ec2:DisassociateRouteTable 
   * ec2:ReplaceRoute 
* Block modification of subnets for our specific VPC
   * ec2:DisassociateSubnetCidrBlock 
   * ec2:AssociateSubnetCidrBlock 
   * ec2:DeleteSubnet 

 I don't want to block them from all VPC changes as they need to be able to perform certain ones.

I guess an alternative could be to use the deny everything wildcard and then use NotActions for the select few actions that the user should be allowed to perform. I would very much prefer to only deny those I need to deny though as I might unexpectedly block the users from performing actions I'm not aware of right now.
## [11][Best way to log very small amounts of data from Lambda/python?](https://www.reddit.com/r/aws/comments/hiht4o/best_way_to_log_very_small_amounts_of_data_from/)
- url: https://www.reddit.com/r/aws/comments/hiht4o/best_way_to_log_very_small_amounts_of_data_from/
---
I'm writing a small lambda function to be used by the public, and I'd like to keep a record of the data submitted. It will be very low volume, so even SES could be used for this; but ideally I'd like to just have a csv of the data. 

Can anyone recommend the simplest way to log a few pieces of data from python/lambda? Wondering if cloudtrail might be the easiest.
