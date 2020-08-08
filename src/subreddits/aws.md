# aws
## [1][DynamoDB atomic counter](https://www.reddit.com/r/aws/comments/i5wgz5/dynamodb_atomic_counter/)
- url: https://www.reddit.com/r/aws/comments/i5wgz5/dynamodb_atomic_counter/
---
Hi,

DynamoDB UpdateItem documentation has an example of how to make an atomic counter.

[https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API\_UpdateItem.html#API\_UpdateItem\_Examples](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_UpdateItem.html#API_UpdateItem_Examples)

&amp;#x200B;

https://preview.redd.it/r6n96jzd0rf51.png?width=591&amp;format=png&amp;auto=webp&amp;s=e1f3d23566b51a2597eb5d370bc94c206d4f1b80

My question is it a "concurrency"  safe operation? Because without transaction support and locking access to the document looks like there is a race condition.
## [2][Why does AWS does not have a region in Switzerland?](https://www.reddit.com/r/aws/comments/i5xwm3/why_does_aws_does_not_have_a_region_in_switzerland/)
- url: https://www.reddit.com/r/aws/comments/i5xwm3/why_does_aws_does_not_have_a_region_in_switzerland/
---
Both GCP, Azure have regions in Switzerland, but there is not a AWS region for Switzerland.

Yes Switzerland is expensive, so does the Australia. AWS have a region in Australia though.
## [3][Singular and Plural problems](https://www.reddit.com/r/aws/comments/i5sdd5/singular_and_plural_problems/)
- url: https://www.reddit.com/r/aws/comments/i5sdd5/singular_and_plural_problems/
---
[https://docs.aws.amazon.com/lambda/latest/dg/services-cloudwatchevents-expressions.html](https://docs.aws.amazon.com/lambda/latest/dg/services-cloudwatchevents-expressions.html)

&gt;For a singular value the unit must be singular (for example, rate(1 day)), otherwise plural (for example, rate(5 days)).

Actually I am working on a Terraform script like below. It doesn't work when the `scanning_minutes` is 1. I switch to use `cron` instead of `rate` now. 🤷🏻‍♂️

    locals {
      scanning_minutes = 1
    }
    
    resource "aws_cloudwatch_event_rule" "scanner_lambda" {
      schedule_expression = "rate(${local.scanning_minutes} minutes)"
      is_enabled = true
    }
## [4][Automating Account Creation &amp; CI/CD with CDK/CodeStar?](https://www.reddit.com/r/aws/comments/i5o4xv/automating_account_creation_cicd_with_cdkcodestar/)
- url: https://www.reddit.com/r/aws/comments/i5o4xv/automating_account_creation_cicd_with_cdkcodestar/
---
Hey all!

I work with a consulting company trying to build out a dashboard of sorts for our developers to handle CI/CD and account creation. Generally, I do not handle development related tasks, so my understanding of CI/CD and coding pipelines is limited, so please bear with my noob-like approach to this question. 

In short, we heavily utilize the CDK in projects we complete for clients, and are trying to build a serverless (if possible), in-house, all-in-one interface for devs to complete work assigned to them, and handle development and coding from inside the solution we are building. 

We recently came across CodeStar, which appears to do a lot of what we are trying to do with respect to managing different projects, automating account creation for different clients, etc. 

However, I have not been able to find any documentation or resources demonstrating how or even if the CDK would work with CodeStar.

With the release of CDK Pipelines, we were alternatively planning on utilizing CodePipeline for CD component of this interface; however, the CI portion of this interface starts to become unclear, as I am unsure if the CDK would play nice with CodeBuild?

In short, is CodeStar a good solution for being able to automate account creation/IAM roles for new clients/projects that come in, as well as full CI/CD, testing, etc, specifically utilizing the CDK? Or is there an alternative set of services that would be recommended to accomplish this? Would Control Tower be appropriate for a situation like this, or is that more intended for enterprise-level organizations that need multiple AWS accounts within their company?

TIA for any direction you all may be able to provide!
## [5][Cloudfront SSL](https://www.reddit.com/r/aws/comments/i5s73d/cloudfront_ssl/)
- url: https://www.reddit.com/r/aws/comments/i5s73d/cloudfront_ssl/
---
Wanted to host a basic static website from an s3 bucket and serve it with cloudfront (HTTPS). Does the custom ssl really cost $600 a month if I get it throuch ACM? If so how do I choose the SNI option that is supposedly free?

&amp;#x200B;

If none of these work, does disabling the cloudfront distribution stop the costs or do i need to delete?

&amp;#x200B;

Thanks
## [6][Lightsail VS Firebase for my project?](https://www.reddit.com/r/aws/comments/i5vkhg/lightsail_vs_firebase_for_my_project/)
- url: https://www.reddit.com/r/aws/comments/i5vkhg/lightsail_vs_firebase_for_my_project/
---
Hey, I basically hardly know anything about these two services but these two were recommended to me and I'm not knowledgeable enough to make a decision. 

My website will include a database of two different account types. A geographical search function. A review system. And host some media (only profile pictures to begin, hopefully adding media to reviews later depending on price and success.) It's a national website, but I cant really predict how many users it will have. Lets say a few hundred to a few thousand a month. Not crazy, but not like 50 a month.

Id say those are the main things to note of the website. This will either take off and be successful in its industry or being a nice project on my resume, so that's where I'm coming from. I only have experience with projects in vanilla JS/HTML/CSS currently and I'd like to have this up as fast as possible, so no, I'm not looking to build out my own entire backend for example. 

Any input on which to use, that could potentially be useful later as well but still quick, easy, and cheap to setup that will fit my websites goals? Also like I said, I'm not that experienced so excuse me if my question sounds stupid.
## [7][AWS EKS + Terraform = EC2 instances?](https://www.reddit.com/r/aws/comments/i5xkxj/aws_eks_terraform_ec2_instances/)
- url: https://www.reddit.com/r/aws/comments/i5xkxj/aws_eks_terraform_ec2_instances/
---
Hey guys.  
I'm designing the CI/CD architecture and wondering whether it's possible to use AWS EKS to run CI/CD Servers to run Terraform that will create and destroy EC2 Instances.

&amp;#x200B;

Thanks in advance!
## [8][What service am I looking for? Proxy internal applications behind SSO.](https://www.reddit.com/r/aws/comments/i5ojfs/what_service_am_i_looking_for_proxy_internal/)
- url: https://www.reddit.com/r/aws/comments/i5ojfs/what_service_am_i_looking_for_proxy_internal/
---
If you're familiar with Cloudflare Argo tunnels, it basically is a service where you can setup SSO in front of a server and Cloudflare will authenticate, and then after you're authenticated, it will proxy the connection to the server you were trying to connect to.

We have used it to have SSH open to the internet by forcing people to auth against Okta, and we have also used it to protect a remote desktop server. I am also interested in protecting some internal web applications that aren't secure when public to the internet, like LibreNMS.

I see WorkLink got close with having a secure browser and proxy, but it appears to only be available for mobile devices. GCP has Identity-Aware Proxy, but I could not find something similar on AWS. Am I missing something?
## [9][How large does data need to be before you use S3?](https://www.reddit.com/r/aws/comments/i5o4gu/how_large_does_data_need_to_be_before_you_use_s3/)
- url: https://www.reddit.com/r/aws/comments/i5o4gu/how_large_does_data_need_to_be_before_you_use_s3/
---
I know the simple rule that “large files should use s3 and small data should use a DB” but what general rule do you follow before you switch to S3? I am going to let users write descriptions and reviews. How long would a description need to be before you’d push them into s3?
## [10][AWS Organizations and prevent IAM user creation in member accounts](https://www.reddit.com/r/aws/comments/i5pe65/aws_organizations_and_prevent_iam_user_creation/)
- url: https://www.reddit.com/r/aws/comments/i5pe65/aws_organizations_and_prevent_iam_user_creation/
---
Hello,

I'm doing some googling/reading but I don't see any obvious way to prevent or lock down the member accounts of an AWS Organization from having IAM users created within them. 

To state it differently, there is a desire to control all IAM account creation such that it can only be done by the root user in the master account of the AWS Organization (I understand the best practices around AWS root account usage; this is still the desired outcome of this exercise.)

Is there some means of achieving this?

Thanks!
