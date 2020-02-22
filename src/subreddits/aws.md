# aws
## [1][Cost Analysis](https://www.reddit.com/r/aws/comments/f7kgom/cost_analysis/)
- url: https://www.reddit.com/r/aws/comments/f7kgom/cost_analysis/
---
How do you guys analyze your costs? I am the finance guy so anything too technical will be lost on me (I can always buy the Cloud guys some beer though to get some help). 

Previously we were exporting the Detailed Billing Report with Resources and Tags into an S3 bucket on a monthly basis. The lack of more frequent updates and the size of the file (70+ GB/month) made it less than ideal but I could get to what I needed utilizing Power BI to summarize the report. 

With that report being deprecated and the Cost and Usage report (CUR) being the new standard, I would like to use the CUR, but I’m being told by corporate cloud ops that there is no longer a single monthly file, and that I would have to download and aggregate thousands of files each month to be able to get a view of the entire bill. 

We do use a third party product called Cloudhealth to help analyze our costs but I have never been able to get a sufficient level of detail out of it without a lot of manual processing.  Some of our products are legacy and single-tenant. As such there are thousands of individual environments that I want to be able to run analysis against and compare.  

Do you all look at your costs? How do you handle it?
## [2][API Gateway endpoint URL for web service on Fargate via NLB and VPC Link?](https://www.reddit.com/r/aws/comments/f7runb/api_gateway_endpoint_url_for_web_service_on/)
- url: https://www.reddit.com/r/aws/comments/f7runb/api_gateway_endpoint_url_for_web_service_on/
---
Hello!

This is the intended setup:

    [API Gateway] - [VPC Link] - { [NLB] - [Fargate] }
    
    # { ... private subnet ... }

I have an HTTP web service running in a container on a Fargate cluster in a private subnet. This service is accessed via an NLB which is also private.

Now I want to access that web service through the NLB via a VPC Link and API Gateway. For setting up a method I need to choose an endpoint URL under Method Execution.

What is this endpoint URL?

On a side note: The NLB is currently publicly accessible for debugging purposes but the goal is that the only component reachable from the internet will be API Gateway.

https://preview.redd.it/iezqq9cuwgi41.png?width=578&amp;format=png&amp;auto=webp&amp;s=2fa4558603dc05c0b42bb897e826e15ae6d95d18

Thanks

Raffael
## [3][Best Practices for Secondary Indexes with DynamoDB](https://www.reddit.com/r/aws/comments/f7g2os/best_practices_for_secondary_indexes_with_dynamodb/)
- url: https://www.trek10.com/blog/best-practices-for-secondary-indexes-with-dynamodb/
---

## [4][New IAM condition key: aws:CalledViaFirst](https://www.reddit.com/r/aws/comments/f7b2nc/new_iam_condition_key_awscalledviafirst/)
- url: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_condition-keys.html#condition-keys-calledviafirst
---

## [5][Labs that define a complete solution?](https://www.reddit.com/r/aws/comments/f7ihcr/labs_that_define_a_complete_solution/)
- url: https://www.reddit.com/r/aws/comments/f7ihcr/labs_that_define_a_complete_solution/
---
Is anyone aware of full-scale scenario labs in which an entire solution is required? Something that says something along the lines of

*Company A is requesting a cloud solution that meets the following requirements.... &lt;list of requirements&gt;. With this information, design a solution that meets their needs and deploy it to AWS.*

As it is, I've been working on labs that hit individual services, but outside of reading the Well Architected Framework and related whitepapers, I'm not getting much in the way of how the "theory" is applied in practice. My thought is that if I can start to work on full-scale scenarios, I might better put these things together.
## [6][Moving from old Serverless System to new Serverless System?](https://www.reddit.com/r/aws/comments/f7fu74/moving_from_old_serverless_system_to_new/)
- url: https://www.reddit.com/r/aws/comments/f7fu74/moving_from_old_serverless_system_to_new/
---
Currently, when the user adds a new item for the first time on the website - PHP app will make an API call to add an item with the sessionId in the DynamoDB.  The table has a TTL setup which will delete the items when it has expired. They can update the Items using the different endpoint.

Here is the current workflow as an example:

&amp;#x200B;

https://preview.redd.it/l99x7obyobi41.jpg?width=667&amp;format=pjpg&amp;auto=webp&amp;s=ea73a6e54f3c429a495e5e6b56cbeb666e153bb1

I need to rewrite new lambdas functions completely and better table structure in DynamoDB (new table: DataV2 with SortKey). What the best approach for the current users that has sessionId on the DataV1 on the website continues with the old serverless system (diagram above) and when the new user adds an item for the first time - it will point to a new Lambda function which will then insert records to a new table Datav2?
## [7][Beginner Question: How can I allow my instance on my private subnet to `wget` and `yum install` dependencies?](https://www.reddit.com/r/aws/comments/f7iqu1/beginner_question_how_can_i_allow_my_instance_on/)
- url: https://www.reddit.com/r/aws/comments/f7iqu1/beginner_question_how_can_i_allow_my_instance_on/
---
I have an enormous install script that needs to go out to the internet for dependencies. Is there a way to give my instance internet access?
## [8][Has anyone used Account Tags for automating account customization?](https://www.reddit.com/r/aws/comments/f7hza8/has_anyone_used_account_tags_for_automating/)
- url: https://www.reddit.com/r/aws/comments/f7hza8/has_anyone_used_account_tags_for_automating/
---
I was looking at the docs here: [https://docs.aws.amazon.com/organizations/latest/userguide/orgs\_tagging.html](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_tagging.html) and was considering automating provisioning of baseline IAM roles based on the account's attributes defined by tags. 

I had seen this solution in the AWS Labs github repo: [https://github.com/awslabs/aws-deployment-framework](https://github.com/awslabs/aws-deployment-framework) which has support for account tags. 

My idea was based on environment = Prod | Non-Prod, a Developer Role for example would have a read-only policy in Prod but more privileges including permission boundaries in Non-Prod. 

Looking for examples and feedback/suggestions from the community on how you guys have used Account tags.
## [9][Savings Plan Update: Save Up to 17% On Your Lambda Workloads](https://www.reddit.com/r/aws/comments/f70qg5/savings_plan_update_save_up_to_17_on_your_lambda/)
- url: https://aws.amazon.com/blogs/aws/savings-plan-update-save-up-to-17-on-your-lambda-workloads/
---

## [10][ALB: Can't seem to get path based routing to work](https://www.reddit.com/r/aws/comments/f7d0ex/alb_cant_seem_to_get_path_based_routing_to_work/)
- url: https://www.reddit.com/r/aws/comments/f7d0ex/alb_cant_seem_to_get_path_based_routing_to_work/
---
Deployed web app successfully on an ECS cluster with ALB.  Everything works great.  I then add a Listener rule to Forward /app1/\* to the existing Target Group, which then results in a 404.

Does the web app itself have to be aware of this new root path "/app1" in order to respond to it?  I would hope not.    I'm using [asp.net](https://asp.net) with Kestrel if that makes any difference here.

I'm trying to achieve the ability to deploy multiple web apps to the same Cluster and use the same Load Balancer URL with specific path to get to the appropriate target group/web app, such as:

[alb.domain.com/app1](https://alb.domain.com/app1) \-&gt; webapp1  
[alb.domain.com/app2](https://alb.domain.com/app2) \-&gt; webapp2

and so on...

What am I missing?
