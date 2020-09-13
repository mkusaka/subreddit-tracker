# aws
## [1][Moving 25TB data from one S3 bucket to another took 7 engineers, 4 parallel sessions each and 2 full days](https://www.reddit.com/r/aws/comments/irkshm/moving_25tb_data_from_one_s3_bucket_to_another/)
- url: https://www.reddit.com/r/aws/comments/irkshm/moving_25tb_data_from_one_s3_bucket_to_another/
---
We recently moved 25tb data from s3 bucket to another. Our estimate was 2 hours for one engineer. After starting the process, we quickly realized it's going pretty slow. Specifically because there were millions of small files with few mbs. All 7 engineers got behind the effort and we finished it in 2 days with help of 7 engineers, keeping the session alive 24/7

We used aws cli and cp/mv command.

We used 

"Run parallel uploads using the AWS Command Line Interface (AWS CLI)" 

"Use Amazon S3 batch operations" 

from following link 
 https://aws.amazon.com/premiumsupport/knowledge-center/s3-large-transfer-between-buckets/

I believe making network request for every small file is what caused the slowness. Had it been bigger files, it wouldn't have taken as long. 

There has to be a better way. Please help me find the options for the next time we do this.
## [2][What does your cloud stack looks like?](https://www.reddit.com/r/aws/comments/irwi76/what_does_your_cloud_stack_looks_like/)
- url: https://www.reddit.com/r/aws/comments/irwi76/what_does_your_cloud_stack_looks_like/
---
So I’ve been expanding my knowledge and skill set about and around cloud and was wondering what does your company’s cloud stack comprises of? This would give a better insight to what tech to learn about.
## [3][Does AWS Amplify remove the DynamoTable after change the schema?](https://www.reddit.com/r/aws/comments/irvb7h/does_aws_amplify_remove_the_dynamotable_after/)
- url: https://www.reddit.com/r/aws/comments/irvb7h/does_aws_amplify_remove_the_dynamotable_after/
---
Right now has 4 dynamodb tables with data. 

I want to add one property to one of that tables, (imageURL). 

If I update the model, AWS Amplify will remove all tables? 

I guess no, but I'm not pretty sure about this.
## [4][Amazon Career Day](https://www.reddit.com/r/aws/comments/ir5svb/amazon_career_day/)
- url: https://www.reddit.com/r/aws/comments/ir5svb/amazon_career_day/
---
https://www.amazoncareerday.com/

Disclaimer: I work at AWS as a Technical Consultant. This is *not* a third party recruiting link. I do not receive any kickbacks, referral fees, nor am I recruiter. I just thought that this would be of interest to this subreddit.
## [5][HOW TO: AWS Lambda CloudWatch Logs Structured as JSON with Python Runtime](https://www.reddit.com/r/aws/comments/irt7tu/how_to_aws_lambda_cloudwatch_logs_structured_as/)
- url: https://www.reddit.com/r/aws/comments/irt7tu/how_to_aws_lambda_cloudwatch_logs_structured_as/
---
hello, my AWS lambda function written in Python runtime runs fine but CloudWatch is rendering the function’s JSON messages in one line. 

    {'version': '0', 'id': '4d26b655-60d4-422a-27d3-9142221a19b1', 'detail-type': 'New Event', 'source': 'new payment', 'account': '397521619882', 'time': '2020-09-13T06:19:59Z', 'region': 'us-east-1', 'resources': [], 'detail': {'state': 'deposit', 'id': '123', 'amount': '100'}}

Same function written in Node.Js when invoked shows up structured in JSON in CloudWatch logs. Which is easy to ready.

How can I get Python to log function's message in JSON? Is there a sample I can look at?
## [6][How to choose between AWS and other Cloud providers](https://www.reddit.com/r/aws/comments/irk7b2/how_to_choose_between_aws_and_other_cloud/)
- url: https://www.reddit.com/r/aws/comments/irk7b2/how_to_choose_between_aws_and_other_cloud/
---
I am struggling to get in to the cloud. Most of my programming life has been on embedded products, and websites back in my day were just simple things. Now we have AWS et al and it's extremely complex in both good and confusing ways.

My app is a dead simple Node.js app that converts HTML pages into PDFs and then concatenate the PDFs together. For this I require two packages

1. Puppeteer, which requires an installation of Chromium (which it does on its own)
2. node-pdftk, which requires an installation of PDFTK

I have this exact app running on my local Windows environment, and I even deployed it to Heroku using some buildpacks. I was so happy when it finally ran, only to find out that you cannot actually write files like PDFs on your Heroku environment and even if you could, all files are erased on reboot. I don't think any of these providers explain any of there products to newbies properly... How could I have no idea that you cannot write files!? Anyways..

Now I am looking at AWS because it has S3 which I guess is where I will store my PDFs. But trying to choose between EC2, EBS, and Lambda is driving me crazy. I even looked at Firebase but the more I read the less confident I get in any of these systems.

All I want is to be able to run my simple Node.js app and be able to install a couple things like PDFTK and Chromium. I only expect a peak of 100s to 1000s of users at a time trying to generate PDFs and not in the realm of 10s of 1000s or more. 

So I was starting to spin up a Lambda server but I think it's the same thing as my Heroku server. It seems that there are some ways to use puppeteer and pdftk with lambda but there seems to be lots of issues with EC2/EBS (dependencies, having to change the OS, etc). 

In the ideal world I would have a linux or windows server that I could just install whatever I want on it and it was easy to start up my node  app and handle HTTP requests.
## [7][Howdy. New to IaC.](https://www.reddit.com/r/aws/comments/irr3oc/howdy_new_to_iac/)
- url: https://www.reddit.com/r/aws/comments/irr3oc/howdy_new_to_iac/
---
Where do I start? I am transitioning from Linux system administration to AWS IaC. I need to pick up typescript for Cloud Development Kit. My coding background is bash and ansible playbooks. Any suggestions for learning typescript?
## [8][One or more ALB / Listeners / TG's per application or microservice](https://www.reddit.com/r/aws/comments/irqrz8/one_or_more_alb_listeners_tgs_per_application_or/)
- url: https://www.reddit.com/r/aws/comments/irqrz8/one_or_more_alb_listeners_tgs_per_application_or/
---
I have a Fargate service behind a ALB. HTTPS listener works fine and Fargate tasks register fine with the TargetGroup.

I'm adding a few more (Fargate) services for different application. I can't find a lot of good information on best practices when using ALB and running more then one services.   
Do you add more listeners and TG's on the same ALB? Or do you add more ALB's? Any pointers?
## [9][Any one used AWS Step Functions with Map state for anything. Example code would be really useful. Thank You.](https://www.reddit.com/r/aws/comments/irmu8x/any_one_used_aws_step_functions_with_map_state/)
- url: https://www.reddit.com/r/aws/comments/irmu8x/any_one_used_aws_step_functions_with_map_state/
---

## [10][Thoughts on Interoperability Between AWS CDK, AWS SAM and AWS Chalice](https://www.reddit.com/r/aws/comments/irluzi/thoughts_on_interoperability_between_aws_cdk_aws/)
- url: https://softwhat.com/thoughts-on-interoperability-between-aws-cdk-aws-sam-and-aws-chalice/
---

