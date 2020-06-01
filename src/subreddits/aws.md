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
## [2][10 Recommendations for writing pragmatic AWS Lambdas in Python](https://www.reddit.com/r/aws/comments/gu90wn/10_recommendations_for_writing_pragmatic_aws/)
- url: https://medium.com/@jan.groth.de/10-recommendations-for-writing-pragmatic-aws-lambdas-in-python-5f4b038caafe
---

## [3][Can "Microsoft Always On VPN" work on AWS?](https://www.reddit.com/r/aws/comments/gugrhx/can_microsoft_always_on_vpn_work_on_aws/)
- url: https://www.reddit.com/r/aws/comments/gugrhx/can_microsoft_always_on_vpn_work_on_aws/
---
Hello,  I am a 3rd year student of computer science. I want to make the lab in this [post](http://blog.tofte-it.dk/tutorial-deploy-always-on-vpn/). But, I do not have a computer with enough hardware for virtualization.  The post in the link contains the following statement.

&gt;Do not attempt to deploy Remote Access on a virtual machine (VM) in Microsoft Azure. Using **Remote Access in Microsoft Azure is not supported**, including both Remote Access VPN and DirectAccess.

My goal is to set up the lab environment in a cloud environment, not on the local computer.  Because I don't have a computer with enough hardware.  Can I run the following network structure on AWS? Is Windows Remote Access running on AWS?  Can you help the student who wants to work on VPN technology?

[Always On VPN Overview](https://preview.redd.it/lt3yt2yxd9251.png?width=794&amp;format=png&amp;auto=webp&amp;s=ca67cf2288fd44ca67ce52994c910c4e5121c9b2)

*Note :  Sorry for my English knowledge. I just started learning.*
## [4][I'm studying how to migrate a relational database to DynamoDB and I'm having a lot of trouble understanding this article. Can someone breakdown how the queries get what you want?](https://www.reddit.com/r/aws/comments/gugwsl/im_studying_how_to_migrate_a_relational_database/)
- url: https://docs.amazonaws.cn/en_us/amazondynamodb/latest/developerguide/bp-modeling-nosql-B.html
---

## [5][Games performance on G4 instances.](https://www.reddit.com/r/aws/comments/gukjl5/games_performance_on_g4_instances/)
- url: https://www.reddit.com/r/aws/comments/gukjl5/games_performance_on_g4_instances/
---
I am testing G4 instance for Cloud Gaming, where my games are installed and running on Cloud machines and player is connected using RDP. Currently we are using G4 instances with T4 GPUs.  I am getting around 25 FPS which is very poor. When we are running same games on Azure NV6 with M60 GPUs  we are getting around 60 FPS. I cannot figure out why NV6 outperforming G4, which is less powerful hardware wise compare to G4.  Thanks.  
Sharing DxDiag of both the machines:  
[https://we.tl/t-OhHZ1kXFaX](https://we.tl/t-OhHZ1kXFaX)
## [6][How to know where are my credits being used?](https://www.reddit.com/r/aws/comments/gufhja/how_to_know_where_are_my_credits_being_used/)
- url: https://www.reddit.com/r/aws/comments/gufhja/how_to_know_where_are_my_credits_being_used/
---
I created an AWS Educate account using my university email and I was given some AWS credits for free. Now when I use the account my credits keep decreasing (which means I am paying for the services I am using), but since I am using several services, I want to know how much I am paying for the different services. How can I know this? When I go to the billing page on the link -  [https://console.aws.amazon.com/billing/home](https://console.aws.amazon.com/billing/home) it says you don't have the permission to access this page. This is perhaps because Amazon doesn't allow Credit / Debit cards to be used in AWS Educate accounts. But is there any way I can at least know how much I paid for which services on what date?
## [7][How to avoid unexpected charges when using AWS free tier](https://www.reddit.com/r/aws/comments/gu1u9t/how_to_avoid_unexpected_charges_when_using_aws/)
- url: https://www.reddit.com/r/aws/comments/gu1u9t/how_to_avoid_unexpected_charges_when_using_aws/
---
Hi, I’m really new to AWS. I’m a bit scared reading through posts about people getting unexpected charges. Besides setting up a billing alert, I want to get to know how the billing system works before starting my free tier instances.

I want to spin up a T3.Micro. To my understanding, there are limits on storage, network, and computation on EC2 instances. Where can I find these limits? If I exceed these limits, how much am I charged? 

I also want to try the DynamoDB. The free tier includes 25G storage, 25 WCU and 25 RCU. The description says it’s enough to handle 200M requests per month. Does it mean I can write and read 200M times without being charged?

Also, are services like S3 and Lambda less likely to generate a big amount of charges even when I exceed the limit since their pricing seems pretty cheap? Are there things to watch out for S3 and Lambda too?

Is there any general advice to avoid unexpected charges?

Thank you :)
## [8][IAM policy to deny EKS cluster creation with public access](https://www.reddit.com/r/aws/comments/guk79z/iam_policy_to_deny_eks_cluster_creation_with/)
- url: https://www.reddit.com/r/aws/comments/guk79z/iam_policy_to_deny_eks_cluster_creation_with/
---
    {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Sid": "EKSNoPublicAccess",
                "Effect": "Deny",
                "Action": [
                    "eks:CreateCluster",
                    "eks:UpdateClusterConfig"
                ],
                "Resource": "*",
                "Condition": {
                    "Bool": {
                        "eks:ResourcesVpcConfig:EndpointPublicAccess": true
                    }
                }
            }
        ]
    }
    

Hi,

My goal is to create an IAM policy rule that prevents clusters being created with public access. The above isn't it...any pointers appreciated.

Thanks!
## [9][Setup a Common EC2 filter](https://www.reddit.com/r/aws/comments/gujug4/setup_a_common_ec2_filter/)
- url: https://www.reddit.com/r/aws/comments/gujug4/setup_a_common_ec2_filter/
---
Quick tip - I hadn't thought of this until recently, this post started as a question and in the process of thinking through it I ended up solving my own problem.

Right now I'm working on a WSUS project, and I'm nearly always in EC2 filtering for "WSUS". This pulls up our list of servers just fine. Then when I make an EC2 change (networking etc)  it brings back to the default view everytime. So then I have to type in WSUS and filter again.

This is a minor annoyance, there is a way to have a button for common filters "WSUS" Button for me, Splunk button for somebody else etc. 

I found I can set a Chrome bookmark for this, and put it on my bookmarks bar in Chrome. 

For a tag we created called "service" the URL ends up being ... ec2/v2/home?region=us-east-1#Instances:tag:service=wsus;sort=tag:Name

Then I had the issue (that I prefer actually) that bookmarks bar doesn't show up unless you're on the New tab. In Chrome settings, Appearance you can have bookmarks bar open all the time.. so now I just hit the bookmark and my filters go back in place.
## [10][AWS CLI for ECS opposite to register-task-definition](https://www.reddit.com/r/aws/comments/gug5n8/aws_cli_for_ecs_opposite_to_registertaskdefinition/)
- url: https://www.reddit.com/r/aws/comments/gug5n8/aws_cli_for_ecs_opposite_to_registertaskdefinition/
---
We keep Task Definitions json files in .gitignore because of secrets in the environment variables.

    aws ecs register-task-definition --cli-input-json file://ecs/task-definition-sample.json --profile $AWS_PROFILE

Though the issue with this approach is that the initial developer needs to pass the file around some other method. I was hoping to use the following to download a copy using the aws cli

    aws ecs describe-task-definition --task-definition $REPONAME-sample--profile $AWS_PROFILE &gt; ecs/task-definition-sample.json

Though when I do *describe-task-definition* followed by *register-task-definition* it fails with:

    Missing required parameter in input: "family"
    Missing required parameter in input: "containerDefinitions"
    Unknown parameter in input: "taskDefinition", must be one of: family, taskRoleArn, executionRoleArn, networkMode, containerDefinitions, volumes, placementConstraints, requiresCompatibilities, cpu, memory, tags, pidMode, ipcMode, proxyConfiguration, inferenceAccelerators

Is *register-task-definition* the opposite to *describe-task-definition* or is there another I missed in the api docs.
## [11][API Gateway - Can you synchronously call integration backends?](https://www.reddit.com/r/aws/comments/gu91ue/api_gateway_can_you_synchronously_call/)
- url: https://www.reddit.com/r/aws/comments/gu91ue/api_gateway_can_you_synchronously_call/
---
So say I have a request that comes in to api gateway. For this request I would have to do 2 things:

1.) Create a user -&gt; lambda 1

2.) Process the user somehow -&gt; lambda 2

Say I now have two lambdas for these and hook them up to api gateway as integrations for a particular HTTP method/route.

For lambda 2 to work, I need a certain piece of information that only lambda 1 can provide after creating the user.

Is there anyway for lambda 2 to know that piece of information from lambda 1? Say through API Gateway orchestrating lambda 1’s reponse and sending it to lambda 2? If not, how would I be able to do that? My only way of thinking is to make ANOTHER SEPARATE API Gateway request would it not? I’m looking to save costs if possible, so trying to avoid that sort of way. It also kind of defeats the purpose I think of 1 api gateway request doing all the things necessary to process ALL the things necessary for that SPECIFIC request. 

So I just kinda wanna know how this would work in context of API Gateway (HTTP APIs specifically, not REST)

Solutions I can come up with only are:

1.) Monolithic lambdas that do too much/are too big with multiple routes/apis

2.) This somehow exists in API Gateway and I just don’t know about it (hopefully you can help?)
