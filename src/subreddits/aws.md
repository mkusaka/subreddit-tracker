# aws
## [1][AWS security groups feel extremely limited &amp; restrictive](https://www.reddit.com/r/aws/comments/g4q6wh/aws_security_groups_feel_extremely_limited/)
- url: https://www.reddit.com/r/aws/comments/g4q6wh/aws_security_groups_feel_extremely_limited/
---
Hi,

We have been looking at automating the creation &amp; management of our security groups. We quickly hit the limits, of 5x SGs per machine &amp; 60x rules per SG.

These can be amended &amp; increased, but AWS maximums still fall short of our requirements.

We have AWS hosted Domain controllers (EC2 instances), and 20+ on-premise domain controllers.

For a successful trust, we need to open upwards of 20 ports inbound per domain controller.

&amp;#x200B;

A regular firewall would allow us to create an alias of ports, and an alias containing all the on-premise DC IPs. One rule would conquer them all in this setting. However AWS security groups doesn't seem to allow this. Instead each port &amp; IP must be added individually.

We have explored referencing a security group inside another security group, but this only appears to work with EC2 instances assigned to the group, not remote IP addresses.

We thought about splitting UDP &amp; TCP into their own ruleset, but with 25x DCs and 13x TCP ports, we're already over the max of 300:

[https://aws.amazon.com/premiumsupport/knowledge-center/increase-security-group-rule-limit/](https://aws.amazon.com/premiumsupport/knowledge-center/increase-security-group-rule-limit/)

&amp;#x200B;

We don't want to open any from on-premise, this needs to be extremely granular and controlled on a per-port basis.

&amp;#x200B;

How would you guys usually approach this scenario? I can't see how any of this would work for a large enterprise.

&amp;#x200B;

Thanks

&amp;#x200B;

&amp;#x200B;

Edit: So far people have only been able to advise use "any" rules, or "reduce the amount of DCs you communicate with".

How do either of these fit into the AWS security model? As quoted from the "Introduction to AWS security":

&gt;\&gt; All AWS customers benefit from a data center and network architecture built to satisfy the requirements of our most security-sensitive customers.  
&gt;  
&gt;\&gt; AWS and its partners offer a wide range of tools and features to help you to meet your security objectives. These tools mirror the familiar controls you deploy within your on-premises environments.

Following the model described above, without the use of aliases, the most security-sensitive customers will need tens of thousands of "rules".

The tools do not mirror familiar controls, all firewalls I know support aliases or named groups of ports &amp; IPs.  *Edit: I stand corrected by* /r/azz_kikkr\*, AWS\* ***and its partners*** *is a catchall clause.*

My scenario isn't a one off case either, what if we have VPN clients connecting in from hundreds or thousands of locations around the world?

We are either missing something major within AWS security, or AWS doesn't natively support scaling up to a very large enterprise.

&amp;#x200B;

Edit2:

Thanks everyone for your input.  So the general consensus seems to be "you'd buy a firewall to do this on-premise, so you should buy one to do it in the cloud".

We expected AWS to support this functionality natively, without the need to employ 3rd party tools.
## [2][Should I stop or terminate ec2 instances to save money overnight?](https://www.reddit.com/r/aws/comments/g4fj9w/should_i_stop_or_terminate_ec2_instances_to_save/)
- url: https://www.reddit.com/r/aws/comments/g4fj9w/should_i_stop_or_terminate_ec2_instances_to_save/
---
Hi,
We have dev ec2 instances where containers are deployed with ecs. The instances are using autoscalling groups. I wrote the lambda function that sets asg capacity to 0 in the evening and back to 1 in the morning. However, I'm thinking is it a good approach or would it be better to simply suspend autoscalligh healthchecks and stop instances?

Edit: sorry for poor english I was typing on my phone
## [3][How to schedule and run sending emails with AWS ?](https://www.reddit.com/r/aws/comments/g4pbyx/how_to_schedule_and_run_sending_emails_with_aws/)
- url: https://www.reddit.com/r/aws/comments/g4pbyx/how_to_schedule_and_run_sending_emails_with_aws/
---
So i have this project in which i can assign multiple users to an onboarding task and each of them will have a diffrent **start\_date** what i want to do is when a set period of time has passed from the users **start\_date** each user should recive a mail with some instructions.   


how can i do this kind of scheduling and email sending with AWS in rails. i know that i can spin up an worker mahine and do this king of things there. is there any way that i could do this without any kind of dedicated  instance ? like lambdas + SNS + SES for email delivery ?
## [4][Help network noob: ssh rule for DMZ home setup?](https://www.reddit.com/r/aws/comments/g4sjpj/help_network_noob_ssh_rule_for_dmz_home_setup/)
- url: https://www.reddit.com/r/aws/comments/g4sjpj/help_network_noob_ssh_rule_for_dmz_home_setup/
---
Hello!  
  
I've changed my home network setup recently, and the ole 'my IP' rule for Inbound SSH isn't doing the trick anymore for getting into my EC2.  
  
LTE Hotspot is 192.168.1.1   
connected to a  
Home Router is 192.168.2.1   
  
LTE Hotspot has a DMZ setup for 192.168.1.22, and Port Forwarding enabled. Internet works great. 
  
The 'My IP' setting sets it to the actual IP of the LTE Hotspot. Are there any other rules I need to set up? Thank you, this is preventing me from teleworking.
## [5][automate silent msi install](https://www.reddit.com/r/aws/comments/g4qjn4/automate_silent_msi_install/)
- url: https://www.reddit.com/r/aws/comments/g4qjn4/automate_silent_msi_install/
---
Hey, 

I'm currently running into the following issue:

We set up a Windows Server EC2 to be used with a third party service. They sent us an msi which needs to be installed on the machine to connect to their service - manually this works just fine. 

&amp;#x200B;

We now would like to automate this - but I'm having trouble to set up the silent install. 

Our current approach is to use a powershell script that will be called from user data that will download the msi from s3 (works) and then use msiexec to install it. 

We used Orca to have a look at the msi and were able to make things like accepting terms and conditions work but in the end and will not install . 

&amp;#x200B;

My question is - is there an easier way to do this in general or is this the right path and it just needs more work? 

If it is the right path, I  can share the current error .
## [6][Possible to automate uploading Lex slot values?](https://www.reddit.com/r/aws/comments/g4qaz7/possible_to_automate_uploading_lex_slot_values/)
- url: https://www.reddit.com/r/aws/comments/g4qaz7/possible_to_automate_uploading_lex_slot_values/
---
I noticed if you want to upload a massive set of slot values you have to upload a .zip that has a correct json object and format for Lex to use.

Is there a way to automate this? Say I want to use a 3rd party script to take a table that’s csv, translate that to a json object that fits the Lex requirements. It needs to be a .zip as far as I know.
## [7][How to trigger AWS Lambda by SMS?](https://www.reddit.com/r/aws/comments/g44lju/how_to_trigger_aws_lambda_by_sms/)
- url: https://www.reddit.com/r/aws/comments/g44lju/how_to_trigger_aws_lambda_by_sms/
---
&amp;#x200B;

https://preview.redd.it/usueom39oqt41.png?width=1599&amp;format=png&amp;auto=webp&amp;s=e02ffd69108b4cb8268bd7526d0a8b321e1f272a

In my last article I have placed a teaser that whether you can trigger Lambda by SMS. Today, We are going to do that! I am really excited to share this with you.

&amp;#x200B;

*NOTE: This article resources are not fully covered in the* **free tier***. Also, you cannot do it directly as there are some resources you will have to raise a ticket to get and adjust services limitations.*

&amp;#x200B;

So, I have been waiting for over a month to finish this article and I postponed it because of the workload that we faced with working remotely and waiting to get the ticket resolved by AWS support team. Which by the way, they were really helpful even in the free plan support.

# What are the involved parts?

There are three main involved part in this event, Customer Engagement, Application Integration and compute services to make this happen.

**Customer Engagement:**

Since we want to trigger a function via SMS, you will need a service or tool to get information from the user. Otherwise, how the function will get triggered?

In this part, we will use AWS Pinpoint. This service enables you to engage with customers through different channels like emails and transactional SMSs. Also, you can validate phone numbers if they are real too!

**Application Integration:**

We are working with SMS, which leads us to work with SNS. SNS is a service that enables you to organize and manage SMS process. Also, it can trigger Lambda too. You got the idea right?

**Compute**:

Since we want to do some computing processes for the SMS content, we will need a computing unit. The best and cheapest option is Lambda. Which is the reason for this article.

# Diagram:

&amp;#x200B;

https://preview.redd.it/f6e66i4ioqt41.png?width=602&amp;format=png&amp;auto=webp&amp;s=cd5b2de8550c352bf44f8c643fb6a103a4c11508

Straight forward scenario, We will contact Pinpoint through SMS, the message will be passed to SNS, which will be responsible for triggering Lambda. No rocket science here.

I am here to show you how-to not to describe the theory behind it. So, shall be begin?

# 1- Request a long/short code from Pinpoint:

Since we want to send an SMS to Pinpoint, it is required to have a code. To obtain one, Please follow the steps from the documentation [here](https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-sms-awssupport-long-code.html).

One point I want to bring attention to is some countries has both short and long code. But, as happened for me, which living in Kingdom of Bahrain, we have only long code, so far.

&gt;*NOTE: It took a while for me to get the code as I was a basic support plan user and there is no default, fast way to obtain this code in my region. Apply for it in advance.*

# 2- Create SNS topic to handle Pinpoint messages:

As we mentioned in the beginning, there is no way to invoke Lambda directly from Pinpoint, creating SNS topic is a must for this purpose.

* From Services, look for SNS and click on it.
* Open SNS console and from the left panel, select topics.

&amp;#x200B;

[s](https://preview.redd.it/wabveswloqt41.png?width=2876&amp;format=png&amp;auto=webp&amp;s=d797efa6c7879f4e27597b717cd6e31568a88984)

* Click on create topic.
* Fill the name for the topic and keep the default values the same.

&amp;#x200B;

https://preview.redd.it/j0x764vooqt41.png?width=2120&amp;format=png&amp;auto=webp&amp;s=f7e758baa2b16fef8cbc0ce18a8c6bc707e23759

* We are done with SNS!

# 3- Prepare IAM role for Lambda:

We will work with Lambda, and for that reason, let us make a proper role to consume SNS messages.

* Open IAM console and click on Roles.
* Click on Create Role.
* Select Lambda from the use cases list and click next.

&amp;#x200B;

https://preview.redd.it/cr7hvgdtoqt41.png?width=2879&amp;format=png&amp;auto=webp&amp;s=3ec774647ac36b729641a81a6e485b6b225b75e6

* In attach permission policies, search for SNS.
* Choose Read Only Access from the list.

&amp;#x200B;

https://preview.redd.it/9x0bg41woqt41.png?width=2062&amp;format=png&amp;auto=webp&amp;s=853fe24c83a31ecefaa7f9d8f25c4e6ce646aa17

* Finish the steps by given the role a descriptive name.

# 4- Lambda Time!!:

Now, we can prepare the function that will consume the received SMS. Let’s start:

* From Services, click on Lambda
* In Dashboard, click on Create Function button.
* Fill the needed information and do not forget to select an existing role, which is the one we created.

&amp;#x200B;

https://preview.redd.it/mw765awyoqt41.png?width=2594&amp;format=png&amp;auto=webp&amp;s=815a3a33682da600d0ca9f26b2ca9099d974a569

* In designer part, click on Add Trigger.

&amp;#x200B;

https://preview.redd.it/bznhouo1pqt41.png?width=2612&amp;format=png&amp;auto=webp&amp;s=70e13a419e95c2b358d9d82e906f96059fa32edf

* In trigger configuration, search for SNS and then look for the SNS topic we created earlier.

&amp;#x200B;

https://preview.redd.it/07ij37m6pqt41.png?width=1646&amp;format=png&amp;auto=webp&amp;s=5bc32d7e8aa20ca13709bdf5937b6e8e0c4ba445

* Done! This function will be triggered whenever SNS topic receives messages.

# 5- Pinpoint Configuration:

So, last step is here. Since we prepared all the resources, we need to configure the part the will trigger all the previous configurations.

* From Services, click on Pinpoint.
* In Pinpoint main page, click on Manage Projects.

&amp;#x200B;

https://preview.redd.it/g3g5io5cpqt41.png?width=2780&amp;format=png&amp;auto=webp&amp;s=2d19bce7cc351ee2e820a9fb81e727c547f821e0

* Create a new project.
* Skip Project Features for now.
* In Project Dashboard, click on Settings -&gt; SMS and Voice.

&amp;#x200B;

https://preview.redd.it/b9v4ne0fpqt41.png?width=2876&amp;format=png&amp;auto=webp&amp;s=bee125da3d3cb214e478e37f58dfc20242dc1b2b

* In SMS and voice page, under Number settings, click on the long/short code you have.

&gt;*NOTE: This is the code you will get after you asked for in the first step.*

&amp;#x200B;

https://preview.redd.it/ymx28kshpqt41.png?width=2282&amp;format=png&amp;auto=webp&amp;s=9332dc47a96d12ca4ee4504b5d2ef40e1d875ae4

* Go all the way down until you see Two-way SMS. Click on it.
* First thing, enable it.
* In incoming messages destination, choose an existing SNS topic, and select the one we created earlier.

&amp;#x200B;

https://preview.redd.it/fnzrcxmkpqt41.png?width=2296&amp;format=png&amp;auto=webp&amp;s=372462306c26763ad227e62704df958bfa805ad4

* Done! easy right?

# How to test?

So, you have everything in place and we need to trigger the function. Just send an SMS to the long/short code you have.

&amp;#x200B;

https://preview.redd.it/977vrkzmpqt41.jpg?width=1125&amp;format=pjpg&amp;auto=webp&amp;s=db34d147fdb35566c84684d2f365346e7c97fbc1

To validate that the function got triggered, check Lambda logs.

&amp;#x200B;

https://preview.redd.it/2lm21hnppqt41.png?width=2878&amp;format=png&amp;auto=webp&amp;s=d4b7098f657edebdbb0d1be77412f14a20527274

# Any useful use-case for this event?

I was thinking, what benefit I will get from triggering Lambda this way, and I guess I have a pretty good reason.

Imagine that you have a marketing dashboard in an EC2 that it gets turned off after working hours. One day, a colleague from marketing department called you and asked you to turn it on for few hours as he has something urgent. Imagine with me opening your laptop if you have it with you, connect it to Wify, open the console, login, don’t forget you enabled MFA, ext… Why can’t you have a Lambda function that when you send a short SMS, will do the job for you! Of course, you can validate the number is yours when you try to trigger it.

# Summary:

AWS Lambda is really great innovation. It amaze me every time I try a new thing to do with. Triggering a computing unit via SMS opens a whole new world of options and possibilities for sys admins, businesses and who like fun projects like me.

Until the next time, don’t forget to wash your hand and stay safe..
## [8][How to deal with ALB+EC2 and an SSL](https://www.reddit.com/r/aws/comments/g4ilur/how_to_deal_with_albec2_and_an_ssl/)
- url: https://www.reddit.com/r/aws/comments/g4ilur/how_to_deal_with_albec2_and_an_ssl/
---
We have an ALB set up with an SSL certificate coming from ACM.
How do we force using only HTTPS traffic from our ALB to our EC2s ? (i.e. "propagate" the SSL encryption)
## [9][AWS ProServe Travel/lifestyle?](https://www.reddit.com/r/aws/comments/g4bwuk/aws_proserve_travellifestyle/)
- url: https://www.reddit.com/r/aws/comments/g4bwuk/aws_proserve_travellifestyle/
---
I'm interviewing for a Cloud Infrastructure Architect role with AWS Professional Services (also called "ProServe"). They told me it would be 50% travel in my terrritory. However, no travel for the foreseeable future due to COVID obviously.

Has anyone here been in this role before? If so, what does the travel / work life balance look like? I'm curious how accurate that 50% really is, or if it's more like 75%.  


Thank you in advance for any insight you can provide! 
## [10][Cloning EC2, Security Groups, Gateway &amp; Subnet?](https://www.reddit.com/r/aws/comments/g4mm7a/cloning_ec2_security_groups_gateway_subnet/)
- url: https://www.reddit.com/r/aws/comments/g4mm7a/cloning_ec2_security_groups_gateway_subnet/
---
I'm wondering if it's possible to make a copy of a topology I have made in AWS?

I have 2 x EC2 Instances which are on a private network, they are both assigned to the same subset with separate security groups. I guess you could say this is an experimental network that will be used for local security testing. 

What I would like to essentially do, is copy the whole topology probably 5 times so that each topology will be on their own private network. 

I know I can create a snapshot of the EC2, but is there an easy way to clone the security groups &amp; network setup? Preferably I'd like to keep the same internal IP's for all services too.
