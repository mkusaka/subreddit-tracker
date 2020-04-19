# aws
## [1][How to trigger AWS Lambda by SMS?](https://www.reddit.com/r/aws/comments/g44lju/how_to_trigger_aws_lambda_by_sms/)
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
## [2][Vault on AWS - A Terraform Project for Secrets Management Anywhere](https://www.reddit.com/r/aws/comments/g3t4ay/vault_on_aws_a_terraform_project_for_secrets/)
- url: https://github.com/jcolemorrison/vault-on-aws
---

## [3][Avoiding ECR costs (NAT/PrivateLink)](https://www.reddit.com/r/aws/comments/g47lj5/avoiding_ecr_costs_natprivatelink/)
- url: https://www.reddit.com/r/aws/comments/g47lj5/avoiding_ecr_costs_natprivatelink/
---
Hi,

I'm trying to deploy a basic app. One container running in a fargate ECS task. I have pushed my container to ECR, everything is configured but the container was failing to start as it couldn't pull the container from ECR.

I thought fargate running in a private subnet without NAT (I don't need NAT) could pull from ECR, but apparently not.

It appears that if you don't have NAT then you need to set up PrivateLink VPN connection from fargate to ECR.

Looking at NAT vs PrivateLink, these two options both cost roughly the same, starting at around $36 per month just to have it running 24/7, plus more for data transfer. 

My app is small and I won't be deploying/scaling very often.

It seems crazy that I need to pay for NAT or a VPN to just pull from another internal aws service.

Do I just need to suck it up and pay, or are there any other options?

Thanks in advance.
## [4][AWS CLI :: increase quota](https://www.reddit.com/r/aws/comments/g43vmt/aws_cli_increase_quota/)
- url: https://www.reddit.com/r/aws/comments/g43vmt/aws_cli_increase_quota/
---
I tried to dig the AWS CLI documentation on how do I increase a specific account quota limit  
but the seems like I can't define in the call \`account-id\` or something similar.

[https://docs.aws.amazon.com/cli/latest/reference/service-quotas/request-service-quota-increase.html](https://docs.aws.amazon.com/cli/latest/reference/service-quotas/request-service-quota-increase.html)

suggestions?

Thanks
## [5][Amazon's JEDI lawsuit is on hold until mid-August](https://www.reddit.com/r/aws/comments/g426dy/amazons_jedi_lawsuit_is_on_hold_until_midaugust/)
- url: https://www.engadget.com/amazons-jedi-lawsuit-is-on-hold-until-mid-august-181403851.html
---

## [6][Best way to use Redis srever and DynamoDb from a single Lambda function?](https://www.reddit.com/r/aws/comments/g3y0fd/best_way_to_use_redis_srever_and_dynamodb_from_a/)
- url: https://www.reddit.com/r/aws/comments/g3y0fd/best_way_to_use_redis_srever_and_dynamodb_from_a/
---
I want to connect to a Redis server in my Lambda function.  I also want to connect to DyanmoDb.  My Lambda function is used to serve HTTP requests for a REST API.  A single Lambda function handles all requests.  My application does not need to scale right now, but I do want to be able to plan ahead.

**Option 1**:  Use ElastiCache as my Redis server.  The problem with that is it requires my Lambda function to use VPC.  I have heard that VPC and Lambda is not a good combination.  I also know that my function will need to access the internet in order to use DynamoDb normally.  So to make this work I could:

* **Option 1.1**. Set up a managed NAT Gateway for my lambda.
   * Pros: Not hard to do
   * Cons:  Will cost me $40 a month.  I don't really want to pay that for this project until I really need it.
* **Option 1.2**. Set up a NAT instance and use that
   * Pros: Fits into free tier of EC2
   * Cons:  More complicated, and I have read it's not a good idea for production?
* **Option 1.3**. Configure DynamoDb work with VPC
   * Pros: Would not need to set up a NAT gateway.  No cost.
   * Cons: Still using VPC for my Lambda function which might cause cold start problems?  Lambda can no longer access internet, but I don't need it to right now.

**Option 2:** Run Redis on an EC2 instance.

* Pros: I don't think this needs a VPC for my Lambda function to connect, which is a plus.  Also fits into the free tier of EC2.
* Cons: Harder to set up than ElastiCache, less robust.  Maybe not good for production?  Security issues?

**Option 3:** Use Redis Enterprise Cloud

* Pros: Easy setup, also has a free tier
* Cons: Hosted on a different cloud provider which I imagine adds some latency?  Might defeat the purpose of Redis if we have to tack on like 30ms for each operation.

&amp;#x200B;

So I'm wondering if anyone else has dealt with this before and how they solved it.  I'm leaning towards **Option 1.3:** Make DynamoDb work with VPC, but then I'm wondering if VPC/Lambda cold start problems are still a problem.   If there is virtually no penalty for VPC in Lambda, I'll just do that.  Otherwise I might do **Option 2:** Run Redis on an EC2 instance, but it sounds like that requires a lot of time and effort to do it right.

**tl;dr** Does VPC on lambda still suck?  Should I host my own Redis instance on EC2?

Edit: I'm 2/2 on screwing up post titles in this sub
## [7][Agnita: Authentication for Create React App using AWS Cognito](https://www.reddit.com/r/aws/comments/g3qhlg/agnita_authentication_for_create_react_app_using/)
- url: https://github.com/bartw/agnita
---

## [8][Optional approvals CodePipeline](https://www.reddit.com/r/aws/comments/g3qweb/optional_approvals_codepipeline/)
- url: https://www.reddit.com/r/aws/comments/g3qweb/optional_approvals_codepipeline/
---
Happy Saturday! I'm considering having optional approvals workflows in CodePipeline. They can be less priority tests, non functional tests that need not be run for all changes, but should be when needed automatically.

There's some ways of getting this done, I could add a static file to the build artifact, based on the contents of the file, those approval flows can skip or run. It's kinda clumsy. I also prefer to not update the CodePipeline once setup. I'm curious to know how the community has tackled this, it's not an uncommon problem.

Edit: Case in point, performance tests. Running them in an ideal world on test environments regularly is trivial. But it usually isn't. Test environments are flaky, running tests that push them need more coordination, and not all changes need to be perf tested. It takes time, effort, sometimes isn't just possible. There could be some applications that can be load tested whenever, legacy systems with many dependencies are typically not in the list.
## [9][Using SQLite with Elastic Beanstalk?](https://www.reddit.com/r/aws/comments/g3tebl/using_sqlite_with_elastic_beanstalk/)
- url: https://www.reddit.com/r/aws/comments/g3tebl/using_sqlite_with_elastic_beanstalk/
---
I wrote a Flask application following [this](https://flask.palletsprojects.com/en/1.1.x/tutorial/) tutorial, which uses SQLite and recommended Elastic Beanstalk as an option for deployment. However, I can't find any documentation that uses SQLite with EB. How do I initialize my SQLite database on my EB app, and is there good documentation for using these tools together?
## [10][Confusion regarding EBS gp2 throughput calculation](https://www.reddit.com/r/aws/comments/g3rxvm/confusion_regarding_ebs_gp2_throughput_calculation/)
- url: https://www.reddit.com/r/aws/comments/g3rxvm/confusion_regarding_ebs_gp2_throughput_calculation/
---
Reference: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-volume-types.html

For gp2, it says:

T = VIR

Where V = volume size, I = I/O size, R = I/O rate, and T = throughput.

Then it gives an example calculating the smallest volume to reach 250 MiB/s throughput:

        250 MiB/s
    =  ---------------------
        (256 KiB)(3 IOPS/GiB)

I assume 256 KiB here is the I/O size, how does it arrive at that value?  Earlier in the document  however (the Volume Characteristics table), for gp2, max IOPS per volume, it says:

&gt; 16,000 (16 KiB I/O)

so 16 KiB seems to be the I/O size, where does 256 come from?
