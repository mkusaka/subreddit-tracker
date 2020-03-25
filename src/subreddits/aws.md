# aws
## [1][CRUD App backend: API Gateway also, or just Lambda?](https://www.reddit.com/r/aws/comments/fom3ah/crud_app_backend_api_gateway_also_or_just_lambda/)
- url: https://www.reddit.com/r/aws/comments/fom3ah/crud_app_backend_api_gateway_also_or_just_lambda/
---
My backend will handle user auth/session management, and CRUD requests to a database on Amazon RDS (MySQL). I'm not clear on the role of API Gateway. 

Lambda gives a direct URL for REST queries for every function, right? What is the additional benefit of Amazon API gateway then? Should a small web application generally be concerned with API gateway, or focus on Lambda plus the database service?

Thank you
## [2][What do you think is the best up-to-date AWS AppSync tutorial for a complex Webapp?](https://www.reddit.com/r/aws/comments/foq7tc/what_do_you_think_is_the_best_uptodate_aws/)
- url: https://www.reddit.com/r/aws/comments/foq7tc/what_do_you_think_is_the_best_uptodate_aws/
---
Currently I am planning to build a project with the following stack:

* ReactJS
* AppSync 
* Amazon DynamoDB 
* Amazon Cognito 
* AWS Lambda

and I am looking for a tutorial which shows me a general overview how it would play together in a production app, so no todo lists or similar. Maybe a ecommerce shop would be interesting even if that's not what I am looking to build but it has many similarities.

What would you recommend?

Thank you for your time :)
## [3][[QUESTION] Using Lambda to change files and create a Github pull request](https://www.reddit.com/r/aws/comments/fonc7g/question_using_lambda_to_change_files_and_create/)
- url: https://www.reddit.com/r/aws/comments/fonc7g/question_using_lambda_to_change_files_and_create/
---
Hi all, I have a pipeline that goes like this:

1.  I created RDS using Terraform, something like:

&amp;#8203;

    module database {
      # other details
      db_storage              = "400"
    }

  2.  I have an RDS storage checker running on Lambda. This checks the remaining amount of storage in RDS and automatically increase the AllocatedStorage if it is below a certain threshold.

3. However, this Lambda function only increases the storage in RDS, and it is not reflected in the repo. So, we have to manually change it in the repo and create a pull request. 

So, I want to automate changing the `db_storage` value in the repo (which means I have to checkout a branch first), and create a pull request.

&amp;#x200B;

I went to check the Github webhook, but there is only an outgoing webhook and no incoming webhook. 

&amp;#x200B;

What should be the approach to achieve this?
## [4][Does anyone work for AWS in Fed Gov Solutions? Deciding if I want to respond to a recruiter.](https://www.reddit.com/r/aws/comments/fo97zt/does_anyone_work_for_aws_in_fed_gov_solutions/)
- url: https://www.reddit.com/r/aws/comments/fo97zt/does_anyone_work_for_aws_in_fed_gov_solutions/
---
Was recently pinged by an AWS recruiter about a cloud infrastructure architect role with Federal Gov Solutions in the DC/VA area.

I've been curious about working at AWS in the past, but always read reviews about how many departments are meat grinders/sweat shops with no sense of work life balance. I'm very motivated and passionate about my work but I'm not at the point in life where I'm interested in 12 hour days as the regular.

I have been told AWS Education Solutions is a nice department to work for in this regard and I'm curious if the same is true for federal.

Does anyone work for AWS in Fed Solutions? What do you think?
## [5][[Question] API as the registered target in the target group of NLB](https://www.reddit.com/r/aws/comments/foo7xb/question_api_as_the_registered_target_in_the/)
- url: https://www.reddit.com/r/aws/comments/foo7xb/question_api_as_the_registered_target_in_the/
---
I am trying to build a API Gateway fronted by an NLB linked to a VPC endpoint so API can be exposed to other VPCs. However the targets in NLB target group needs to be either an IP address or an instance. How can I direct the traffic from NLB to API? Do I need to write a Lambda to get the IP address of APIs?
## [6][Request more than 5 eip?](https://www.reddit.com/r/aws/comments/fopgoy/request_more_than_5_eip/)
- url: https://www.reddit.com/r/aws/comments/fopgoy/request_more_than_5_eip/
---
Default aws only gives max 5 eip. I want to more, how to have more eips and what’s its cost?
## [7][Free Unlimited Load Testing for 30 days if you help me by completing a 2 minute Google Form](https://www.reddit.com/r/aws/comments/fop2ck/free_unlimited_load_testing_for_30_days_if_you/)
- url: https://www.reddit.com/r/aws/comments/fop2ck/free_unlimited_load_testing_for_30_days_if_you/
---
Everyone is offering something these days, in order to help other people or other businesses. Being a startup that is not launched yet, is very hard to contribute with something valuable to the community.

So, here at [Rungutan](https://rungutan.com), we decided to offer you useful data instead 🙂. 

In order to do this, we will need your help as well: share with us your Load Testing routine in times of crisis and __get our FREE report on this__ 📰 + __1 Month FREE access to the Rungutan Load Testing Platform (in Beta Test version)__ 📣

PS: We welcome all DevOps, Testers and Developers to contribute with their responses

[https://docs.google.com/forms/d/e/1FAIpQLSe7vt7SXagYg3wMgNPIEz4vW8j0rIh-rqv9ydiwLQkhFYQEaA/viewform](https://docs.google.com/forms/d/e/1FAIpQLSe7vt7SXagYg3wMgNPIEz4vW8j0rIh-rqv9ydiwLQkhFYQEaA/viewform)
## [8][Amazon CA Question](https://www.reddit.com/r/aws/comments/fod47f/amazon_ca_question/)
- url: https://www.reddit.com/r/aws/comments/fod47f/amazon_ca_question/
---
I'm sure you have seen that Amazon communicated their plans to migrate all SSL/TLS certs to their own CA, Amazon Trust Services. 

I have been doing the recommended test by fetching an S3 object they provide and looking for a 200 OK in the response. Funny thing is, we have some instances that are 8-10 years old, some of which haven't really been kept up-to-date and yet they are passing with a 200 OK, just fine. Taking that a step further, some of those very same instances do not show Amazon Trust Services in the .crt that gets used. (I verify the file being used by doing a `curl -v` .)

So, my question is, have any of you seen a failure yet? If so, what does it look like? And if you're willing to share, what type of instance is it and how old?

Cheers!
## [9][Looking for Advice on Systems Manager Patch Manager](https://www.reddit.com/r/aws/comments/fo935v/looking_for_advice_on_systems_manager_patch/)
- url: https://www.reddit.com/r/aws/comments/fo935v/looking_for_advice_on_systems_manager_patch/
---
I've been tasked with getting a handle on the current environment patch compliance and software inventory. The environments are all AWS. I have the SSM agent installed and upgraded to the newest version pretty much across the board. I have almost all instances showing inventory.

When it comes to patching compliance, I am struggling to come up with a plan.

What I have so far are 4 environments: dev, staging, qa, and perf. Most of the instances are tagged with the appropriate tag for the respective environment. So I have a different maintenance windows setup for each environment to only scan for needed updates so I can report compliance. Eventually, I will need to patch these instances and this is where I am really struggling. 

We have ASGs and those instances spun up via scaling policy will get the environment tag I mentioned above. What I don't want to do is kick off an install maintenance window and have these ASG machines get rebooted and potentially triggering more scaling events. Ideally, I wouldn't target my ASGs at all for patching, but instead automate the AMI to be updated and attached to the launch template.

So my first couple questions are:
How are others patching their static instances separate from their instances spun up via ASG?
How are you leveraging tagging and Patch Groups to ensure the right patches are getting applied to the right instances?
## [10][What is the best way to tags to AWS resource to oversee cost](https://www.reddit.com/r/aws/comments/fodth4/what_is_the_best_way_to_tags_to_aws_resource_to/)
- url: https://www.reddit.com/r/aws/comments/fodth4/what_is_the_best_way_to_tags_to_aws_resource_to/
---
Due to the whole Corona Virus outbreak, the company want to save some money and review existing AWS resource and cut cost.  What is the best way to go about this. I want to see cost of EC2/EBS/SNAPSHOT as well S3/Kinesis stream etc ?  Also we monitor usage via Wavefront hence we do lot of Data transfer, is there way to see cost of those too ?
