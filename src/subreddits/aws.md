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
## [2][Migrating from GCP to AWS. Need advice to get up to speed](https://www.reddit.com/r/aws/comments/gqst9c/migrating_from_gcp_to_aws_need_advice_to_get_up/)
- url: https://www.reddit.com/r/aws/comments/gqst9c/migrating_from_gcp_to_aws_need_advice_to_get_up/
---
Our management has decided to move from GCP to AWS.  
 We use the following services in GCP:

Compute: 

GKE, Compute Engine VMs,  

Storage: 

SQL ( for Postgres DB), Google Storage bucket,

More helper services: 

Pubsub, cloud functions, cloud scheduler, Stackdriver ( logging and monitoring), cloud tasks, cloud build, bigQuery   


In the first phase I'll be migrating just the kubernetes, databases. The helper services will require changes in application code, which i will be doing in the next phase. 

Now, I just logged into my AWS account. GCP has a very neat step by step documentation for those migrating from AWS to GCP. But for moving from GCP to AWS I dont find any documentation. Can someone pls give me a few pointers that I need to keep in mind before I start creating the infra on AWS? Are there any gotchas that I need to be aware of?
## [3][Send cloudinit output to cloudwatch logs](https://www.reddit.com/r/aws/comments/gqvq8j/send_cloudinit_output_to_cloudwatch_logs/)
- url: https://www.reddit.com/r/aws/comments/gqvq8j/send_cloudinit_output_to_cloudwatch_logs/
---
When having a lot of EC2 instances that use UserData, it gets difficult to track the output, as you have to log in to each instance to see the output of the cloud-init logs. I wrote a small golang program that can send the output of your commands straight to cloudwatch logs. That way, you can easily see and track errors in userdata. PR's &amp; feedback welcome! GitHub Repo: [https://github.com/in4it/tee2cloudwatch](https://github.com/in4it/tee2cloudwatch)
## [4][AWS SSO and Jenkins/Terraform/some other CI](https://www.reddit.com/r/aws/comments/gqsr0a/aws_sso_and_jenkinsterraformsome_other_ci/)
- url: https://www.reddit.com/r/aws/comments/gqsr0a/aws_sso_and_jenkinsterraformsome_other_ci/
---
Hey all!

We have been working on implementing SSO with GSuite for Central auth, and working great(ish).

One issue I am working on is that our infra built in Terraform. Locally, its okay, it still can find the users SSO credentials, even if you have them post into a credentials file manually.

My question is regarding Jenkins or other CI services. How do you provision your credentials, so it can authenticate with SSO? 
Do you use some helper, or just generate an IAM role in the account you use to deploy in?
## [5][Show content from S3 on the frontend](https://www.reddit.com/r/aws/comments/gqweuw/show_content_from_s3_on_the_frontend/)
- url: https://www.reddit.com/r/aws/comments/gqweuw/show_content_from_s3_on_the_frontend/
---
Hi all, 

I am working on an app, and where I need to find a way to show content from S3 on my frontend. 

Before, I wrote a lambda that runs once a day, scrapes certain content and saves a file to S3 as a txt file. But I don't know how to access it/show it, since I have 0 experience with React. What would you suggest as a solution?  


Is it better to write the logic in the react file or there is a better way? I could write a python lambda,  that will loop through the bucket and return the content, but still don't know how can I connect it to the FE.

I am still new to programming, so I hope this makes sense.
## [6][Which tool for querying semi-structured data?](https://www.reddit.com/r/aws/comments/gqvf2h/which_tool_for_querying_semistructured_data/)
- url: https://www.reddit.com/r/aws/comments/gqvf2h/which_tool_for_querying_semistructured_data/
---
Advice needed. We are building a modular pipeline in AWS for text processing. I am building a lot of timing and other metrics, which will be collected as JSON objects with various tags and values.

I'd like to gather all this and query it, such as finding min, max and avg values over time spans, finding patterns in the data etc.

What's the most cost effective and/or best tool for storing these metrics, and doing queries that can do basic aggregation and other data mining on generic data objects that may evolve over time (not all objects might have the same exact schema as we evolve the system)  I don't have that much familiarity with NoSQL but I wonder if that's the direction I should go? We integrate with Cloudwatch already, is that useful for more general data processing?

I guess this may be two questions: which data store and which query tool. Hoping to find something that doesn't require a huge time investment to learn. 

Thanks!
## [7][AWS SSO: no reports for user and group based access?](https://www.reddit.com/r/aws/comments/gqsgis/aws_sso_no_reports_for_user_and_group_based_access/)
- url: https://www.reddit.com/r/aws/comments/gqsgis/aws_sso_no_reports_for_user_and_group_based_access/
---
Where do I get a high level view of every AWS account a particular user or group is assigned to in AWS SSO? I can't see this information anywhere within the Console (side rant: and there's no API yet).

I can go into each individual AWS Account in SSO and see which groups and users are assigned to that _specific_ account. But doing that for 20+ accounts, one by one, to provide a report for compliance and audit seems obscene.
## [8][Is there a better way to write this buildspec.yml file? Newbie trying to use codebuild / codepipeline](https://www.reddit.com/r/aws/comments/gqqpgf/is_there_a_better_way_to_write_this_buildspecyml/)
- url: https://www.reddit.com/r/aws/comments/gqqpgf/is_there_a_better_way_to_write_this_buildspecyml/
---
Pardon for my ignorance, I have tried a few hours to figure this out

My buildspec file I wrote:

    version: 0.2
    
    phases:
      install:
        commands:
          - echo Installing Serverless...
          - npm install -g serverless
      pre_build:
        commands:
          - echo Installing frontend dependencies...
          - cd frontend-serverless &amp;&amp; npm install
          - echo Installing backend dependencies...
          - cd .. &amp;&amp; cd backend-serverless &amp;&amp; npm install
      build:
        commands:
          - echo Deploying serverless backend...
          - cd .. &amp;&amp; cd backend-serverless &amp;&amp; serverless deploy
          - echo Deploying serverless frontend...
          - cd .. &amp;&amp; cd frontend-serverless &amp;&amp; serverless
      post_build:
        commands:
          - echo Deployment completed on `date`
    

My Project structure :

    [frontend]
    [backend]
    buildspec.yml
frontend and backend are two different project folders both of which rely on the serverless framework.

My codebuild works and deploys correctly but notice how I am using ```cd ..``` in a few of the commands to get back to the previous directory.  It is very ugly however I cannot find a better way to make this work at the moment.

Is there a better way to reference the root working directory in codebuild?  I can't find any documentation.

Thank you for your help.
## [9][So You Suddenly Need to Reduce Your AWS Bill: 4 Things We Did](https://www.reddit.com/r/aws/comments/gqcbqo/so_you_suddenly_need_to_reduce_your_aws_bill_4/)
- url: https://www.patientcolife.com/so-you-suddenly-need-to-reduce-your-aws-bill-4-things-we-did/
---

## [10][Can anyone help me figure out a workaround or a better solution for my current problem? I can't figure out how to make my CFT access S3 bucket from different regions.](https://www.reddit.com/r/aws/comments/gqx0uj/can_anyone_help_me_figure_out_a_workaround_or_a/)
- url: https://www.reddit.com/r/aws/comments/gqx0uj/can_anyone_help_me_figure_out_a_workaround_or_a/
---
**The use case** \- I have a lambda that will auto-tag every newly created ec2 and rds instances with the creator's IAM user id. I uploaded the lambda code in an S3 bucket in the form of a zip.

  This is how my bucket looks -   
**bucket name:** resources-for-autotag  
|  lambda-function.zip  

I've written a CFT which deploys a CloudWatch rule and some IAM policies. This CFT takes the code from the bucket and deploys the lambda along with the CW rules and IAM policies.  
My CFT requires the name of the bucket which should be provided by us.

**The problem** is that CFT can't access the bucket if launched in different regions than where my resource bucket is located. I tried making the bucket public(which isn't the best practice, I know) but that didn't work either.   
I'm forced to create a bucket in every region with the same content. It's becoming pretty redundant and time taking and honestly the repetitive work is getting to me more than I'd like to admit.   
I'm working on an Organization account, and my permissions and access to other AWS services are pretty limited. 

In addition to all this, I have more than 50 accounts I have to deploy this to.

   
**Trials  -** 

* I tried bucket replication. But it doesn't make much sense as I have to use the content only once, and the process is more time taking than just creating buckets in different regions and uploading the zip file. 
* CFT StackSet won't work either for the same reason that S3 buckets can't be accessed from other regions.
* I cannot add the lambda function directly in the CFT, as it's not usually a best practice and moreover the code is more than the size that is allowed inline CFT.
* I tried writing a lambda, using boto3 which creates S3 buckets in different regions, uploads the file, and wrote another lambda to empty and delete all the redundant buckets, but I keep getting the Illegal constraint error, even when I gave the Location Constraint and specified the default AWS-REGION
* I tried creating S3 buckets through CLI, but the same location constraint error is haunting me.
* I tried Cloud Custodian auto-tag policy, but it's not working for some reason on newly created instances, and their documentation wasn't of much help.

**Questions I'm hoping I can get answers to -** 

1. Is there a way I can modify my CFT, so I can create S3 and use it in the same deployment stack?
2. Are there any workarounds for working with Amazon s3 in different regions?

I'm sorry if it is a long and boring problem, I've been trying some workaround for a week and the manual work is getting to me.
## [11][Athena create partition.](https://www.reddit.com/r/aws/comments/gqs1fk/athena_create_partition/)
- url: https://www.reddit.com/r/aws/comments/gqs1fk/athena_create_partition/
---
So the simple question is... I'm familiar with CTAS queries and how I'm able to create an entirely new table as the output of a merge of two other tables.

But in this instance, what I would like to do is run a query against two date partitioned tables every day to add a partition to a table for a particular partition key.

Essentially something along the lines of:

     ALTER TABLE merged ADD PARTITION (date = '20200526') WITH (format='parquet', external_location='s3://my-bucket/my-key') AS select * from table1

But I can't see any documentation around creating a partition from query results... just creating tables.  This feels lacking...
