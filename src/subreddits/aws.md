# aws
## [1][re:Invent 2020 will be free and virtual!](https://www.reddit.com/r/aws/comments/i06b6w/reinvent_2020_will_be_free_and_virtual/)
- url: https://reinvent.awsevents.com
---

## [2][How to automate AWS resource deployment the right way?](https://www.reddit.com/r/aws/comments/i0lim2/how_to_automate_aws_resource_deployment_the_right/)
- url: https://www.reddit.com/r/aws/comments/i0lim2/how_to_automate_aws_resource_deployment_the_right/
---
Over the last few years, I built a rather complex platform on AWS. I used Terraform for everything, and I am pretty happy with it.

Now I am bootstrapping a new project on AWS.

Here are my options (I ignored native CloudFormation on purpose) :

* The easy option is to stick with Terraform. Despite all its quirks. At least I know it well, and I'll be productive with it.
* Then there is the easy upgrade: using Terragrunt from day one. Still Terraform. But probably fewer headaches. (no experience with it, it just smells good)
* I could also go with the CDK way. After all, AWS looks committed to make it the reference way to manage infrastructure. No experience with it either. And apparently, new AWS features lag behind the Terraform AWS provider because AWS itself slowly integrates new APIs in CloudFormation. And I have no experience with CF.
* I was already struggling to pick some tools and stick to it, but there is the new kid on the block: CDK for Terraform. Now, TBH, I'm lost.

In my former platform, I've never achieved full automation: PR -&gt; validation -&gt; infrastructure updated.

What's the fastest but still clean way to achieve this with a blank slate?

PS: I know a missed a few options. Please only raise them if you truly believe they are _much_ better for my use case. :-)
## [3][Starting CDK](https://www.reddit.com/r/aws/comments/i0j3il/starting_cdk/)
- url: https://www.reddit.com/r/aws/comments/i0j3il/starting_cdk/
---
Hi guys, I've been tinkering with CDK and trying to get some infrastructure going. I'm familiar with CFN but I don't have a coding background so I'm struggling a bit with CDK. Trying to use Typescript. If anyone can point me to some resources, finished projects or guides that'd be very helpful.
## [4][Modifying schema manually on AWS glue](https://www.reddit.com/r/aws/comments/i0lzdz/modifying_schema_manually_on_aws_glue/)
- url: https://www.reddit.com/r/aws/comments/i0lzdz/modifying_schema_manually_on_aws_glue/
---
Hey guys,

I'm trying to avoid the cost and complexity of glue crawlers.

I was reading AWS glue documentation and tried a few things in the console, and I'm not sure I understand the need for a glue crawler in order to analyze the schema. In case that I want to store the data as is, and I know about every schema change that is taking place, will there be any issue updating the schema manually, via the console -&gt; edit schema? 

From what I read at the documentation: [https://docs.aws.amazon.com/glue/latest/dg/start-console-overview.html](https://docs.aws.amazon.com/glue/latest/dg/start-console-overview.html)  
I see 

&gt;You can also populate the Data Catalog with manually created tables. With this method, you provide the schema and other metadata to create table definitions in the Data Catalog. Because this method can be a bit tedious and error prone, it's often better to have a crawler create the table definitions.

Why is it error-prone?

Thank you!
## [5][IAM support for prefix lists is missing?](https://www.reddit.com/r/aws/comments/i0lvs2/iam_support_for_prefix_lists_is_missing/)
- url: https://www.reddit.com/r/aws/comments/i0lvs2/iam_support_for_prefix_lists_is_missing/
---
I'm trying to use the new (1 month ago) customer-managed prefix list functionality (announcement: [https://aws.amazon.com/about-aws/whats-new/2020/06/amazon-virtual-private-cloud-customers-use-prefix-lists-simplify-configuration-security-groups-route-tables/](https://aws.amazon.com/about-aws/whats-new/2020/06/amazon-virtual-private-cloud-customers-use-prefix-lists-simplify-configuration-security-groups-route-tables/)) to try and simplify some management. The announcement says this feature is available in every public region, and the documentation is available publicly. However IAM users are documented as not having permissions by default, with this link ([https://docs.aws.amazon.com/vpc/latest/userguide/managed-prefix-lists.html#managed-prefix-lists-iam](https://docs.aws.amazon.com/vpc/latest/userguide/managed-prefix-lists.html#managed-prefix-lists-iam)) provided in order to document the process of allowing IAM users to operate on managed prefix lists.

But in my account, if I add the Actions outlined from that documentation page (plus I've guessed that the Create and Delete operations should also exist as they are generic verbs available for almost every resource type in IAM) then I get the messages in this screenshot (https://imgur.com/ZhD8QzK). It seems like IAM doesn't know what these operations are even though this service has been available in all public regions for a month.

Is there some kind of "IAM upgrade" I need to do to make my account aware of new stuff like this?
## [6][Lambda Node.js deleteObjects() doesn't work](https://www.reddit.com/r/aws/comments/i0jstu/lambda_nodejs_deleteobjects_doesnt_work/)
- url: https://www.reddit.com/r/aws/comments/i0jstu/lambda_nodejs_deleteobjects_doesnt_work/
---
Hi, I'm using AWS Lambda and trying to delete objects from an S3 bucket but it doesn't work, I don't seem to even get a callback. I don't get any error messages. The function executes and nothing happens. Other operations like uploading and getting object data works. I copied the code from [AWS S3 SDK for JS](https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/S3.html#deleteObjects-property). My bucket doesn't have versioning enabled. I have given my Lambda function full s3 access role. Where could the problem be? My code is this:

&amp;#x200B;

        const params = {
          Bucket: bucketName, 
          Delete: { 
            Objects: [ 
              {
                Key: 'images/3---compr/compressed-small/img7.jpg' 
              },
              {
                Key: 'images/3---compr/compressed-small/img2.jpg' 
              }
            ]
          }
        };
        
        s3.deleteObjects(params, function(err, data) {
          if (err) console.log(err, err.stack);
          else     console.log(data);
          console.log("this doesn't get logged");
        });

edit: I successfully fixed my error. The problem was that i was using async/await incorrectly + I changed s3.deleteObjects(params); to s3.deleteObjects(params).promise();
## [7][How to find when the vpc or subnets were last used?](https://www.reddit.com/r/aws/comments/i0j0qw/how_to_find_when_the_vpc_or_subnets_were_last_used/)
- url: https://www.reddit.com/r/aws/comments/i0j0qw/how_to_find_when_the_vpc_or_subnets_were_last_used/
---
Is there a way to find out when the vpc or the subnets were last used?
Trying to delete a few dozen VPC's and subnets created by users years back but is not in use now.
## [8][Future Scheduled Event Triggers (many, programatic) - AWS Native Solution?](https://www.reddit.com/r/aws/comments/i0gh3d/future_scheduled_event_triggers_many_programatic/)
- url: https://www.reddit.com/r/aws/comments/i0gh3d/future_scheduled_event_triggers_many_programatic/
---
This seems like a bit of an odd thing that AWS appears to be missing. We have a large number of events that occur throughout the course of each day. Events occur, and at a period of X hours later, another event needs to happen. The X can be variable for each event, and there are literally thousands of events each day. 

&amp;#x200B;

We currently have a system where we are just scanning the database looking for the initial event X hours ago, which was okay at small scale, but performance in this model does not scale well. 

&amp;#x200B;

We have designed a system where the initial event occurs, calculates the X time for itself, and puts a row in the SQL database with the "trigger date/time" -- Then we poll every few minutes that table for any trigger date/time that has now passed, and execute it and remove it from the table. We feel this will work, and will certainly scale better, but it got me thinking that it seems odd that there isn't really a "scheduler" type function that can scale well. CloudWatch events has a pretty low threshold, and even eventBridge has a low cap of how many schedules could exist at a given time. 

&amp;#x200B;

I was thinking about how Alexa works, where you can set a timer, or a reminder for a given duration of time, and it runs it. I figure its not doing it on-device, as you can see the event set in the Alexa app, and you (can) get push notifications back via the app too. So I wonder what Amazon would be using to do that (setting an event to trigger X time in the future) -- Somehow I can't see them using CloudWatch events to fire off Alexa alarms, but I can't imagine they just have tasks scheduling up on an EC2 instance either. 

&amp;#x200B;

I did find this:  
[https://www.freecodecamp.org/news/how-to-schedule-ad-hoc-tasks-with-dynamodb-ttl-and-lambda-421fa5778993/#:\~:text=The%20default%20limit%20on%20CloudWatch,millions%20of%20ad%2Dhoc%20tasks](https://www.freecodecamp.org/news/how-to-schedule-ad-hoc-tasks-with-dynamodb-ttl-and-lambda-421fa5778993/#:~:text=The%20default%20limit%20on%20CloudWatch,millions%20of%20ad%2Dhoc%20tasks).

Which actually seemed REALLY clever and cool, until I got to the bit where the average delay time from the trigger was 11 minutes, and the specs allow for up to 48 hours delay!
## [9][Error while running eb init on terminal](https://www.reddit.com/r/aws/comments/i0itun/error_while_running_eb_init_on_terminal/)
- url: https://www.reddit.com/r/aws/comments/i0itun/error_while_running_eb_init_on_terminal/
---
 I'm getting a service error when I run eb init that says "ServiceError- 'express' not a valid key=value pair...". Someone please help me out! What should I do in order for it to run properly?
## [10][AWS IAM Footguns &amp; Security Basics](https://www.reddit.com/r/aws/comments/i005rd/aws_iam_footguns_security_basics/)
- url: https://cloud-fundamentals.com/blog/iam-footguns-security-basics/
---

