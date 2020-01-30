# aws
## [1][AWS SES US-West-2 Blacklisted](https://www.reddit.com/r/aws/comments/evr5v0/aws_ses_uswest2_blacklisted/)
- url: https://www.reddit.com/r/aws/comments/evr5v0/aws_ses_uswest2_blacklisted/
---
FYI I've been troubleshooting emails getting bounced from our SES account and I noticed that all US-West-2 IPs in SES now appear to be blacklisted [according to mxtoolbox](https://imgur.com/a/Cgzk33A).  I've opened an incident with AWS support on this to investigate.
## [2][What cheap relational-like DB can I use serverless for a simple web app?](https://www.reddit.com/r/aws/comments/evp5mn/what_cheap_relationallike_db_can_i_use_serverless/)
- url: https://www.reddit.com/r/aws/comments/evp5mn/what_cheap_relationallike_db_can_i_use_serverless/
---
Hello!

I have a small web app where family members track their daily expenses. They are divided into categories, and there is also "budgets" feature where users can set for themselves a recommended monthly limits for a combination of categories. 

Currently I am self-hosting it, but I would like to migrate it to AWS.

Frontend is simple JS that will live in S3/CloudFront. Login will probably be Cognito. Backend will be Lambda/API Gateway. I am not sure what should I use for storing data.

Currently I am using MySQL with a few tables - expense categories, expense amounts entries and budget limits. 

Volume-wise, currently it is about 10k records (&lt;10MB) total. It doesn't grow a lot.

Usage-wise, we are speaking about single write operations per day and a little more reading operations. Dashboard displays latest N expenses, and there is filtering by category and date. There is also on-demand "generate report" feature that can read really a lot of data (all of it), but this is used maybe once per a few months/year.

Money-wise, I am targeting a solution that will cost me &lt;$5 per month. All the mentioned above services should fall in sub-$1 expense and what is troubling me is the DB I can't decide on:

From what I know about AWS I have the following options:

1. DynamoDB - not relational. Will have to break my brain to adjust filtering to it, but let's say I can do it. I will need maybe 3 tables (categories, budgets, actual expense items) and I am not sure how to do capacity. Auto scale, or provisioned? 1 WCU should be enough, but I am not sure about RCU. Lowest case scenario (1WCU/1RCU) is about $.60 per table or $2 for 3 tables. Could grow more if I need GSI for filtering per category/date. 

2. Aurora Serverless - it looks like drop-in replacement for my current MySQL implementation. Haven't used it before and I expect there will be surprises, though. What scares me is the cold start everyone complains about. Waiting 20+ seconds to awaken the DB is way too much for a comfortable user experience. 

3. Aurora/RDS always-on - the smallest possible instance db.t3.micro costs $0.017 per hour =~ $11 per month. Way above limits.

4. Data files in S3 and custom read-write operations on them - not really a fan of this approach... It will be cheapest, but brings in a number of additional problems.

5. Non-AWS solution - what? 

What are your recommendations? (Price checks are for EU regions)
## [3][Deployed website, but can't tell where it is managed.](https://www.reddit.com/r/aws/comments/evqtag/deployed_website_but_cant_tell_where_it_is_managed/)
- url: https://www.reddit.com/r/aws/comments/evqtag/deployed_website_but_cant_tell_where_it_is_managed/
---
So, a while back, I a deployed a website, static no database, and it is up and running. I can see the record in route53, points to cloudfront, but there doesn't seem to be a distribution. 

I remember just adding a repository and it autodeployed on changes. I was thinking it was Lightsail, but there is nothing there either.

Anyone have an idea where it might be? There are a million services, and I must be skipping over the one I need or not know how to correctly check billing to work from there.
## [4][I'm new, which services do I want to use for this?](https://www.reddit.com/r/aws/comments/ew4pgw/im_new_which_services_do_i_want_to_use_for_this/)
- url: https://www.reddit.com/r/aws/comments/ew4pgw/im_new_which_services_do_i_want_to_use_for_this/
---
I'm making a simple blog site for a group of us (5 people). The site has a login, someone writes text, hits post and that's all there is to it. Maybe I'll allow images.

What services do I want to use? 1) I want page loads to be &lt;200 milliseconds. 2) I have already used up all 12 month of free tier discounts on previous projects. There's a few ways I can think of doing this one

1. Have the site entirely run statically on s3 except for when I do login and a http POST, which I use lambda to handle
2. Use DynamoDB for static HTML pages since they change daily and may shorten my lambda function runtime
3. Use EC2 instead of lambda
4. Use EBS instead of DynamoDB for the static HTML files
5. Other combinations

For logging in and creating a post I don't mind the page being slow. But for all visitors I'd like the html, css, js and images to all load in &lt;200.

\-Edit- I suspect my app needs &lt;128mb to execute but I'm not 100% sure. It's written in C#. Might need 256 but I doubt anything more. I also not sure how EC2 is billed. If I use on average 10% of my CPU and I want an always up CPU to run my site, do I pay 24hrs or would it be 2.4ish hours?
## [5][CentOS 7.7 custom AMI with ENA driver installed but can't ping](https://www.reddit.com/r/aws/comments/evxbhs/centos_77_custom_ami_with_ena_driver_installed/)
- url: https://www.reddit.com/r/aws/comments/evxbhs/centos_77_custom_ami_with_ena_driver_installed/
---
I created a custom AMI using a OVA image I uploaded. The image works fine with T2 instance types but when I switch to T3 it could not get a network. After looking into it, it appears I needed to install the ENA driver. I followed the guide provided here:  [https://aws.amazon.com/premiumsupport/knowledge-center/install-ena-driver-rhel-ec2/](https://aws.amazon.com/premiumsupport/knowledge-center/install-ena-driver-rhel-ec2/)  and installed driver 2.2.2. However, shutting the instance down and changing it to T3 still powers up without network connectivity. What am I missing here?
## [6][Installing the CloudWatch Agent incur in additional charges?](https://www.reddit.com/r/aws/comments/evv5s1/installing_the_cloudwatch_agent_incur_in/)
- url: https://www.reddit.com/r/aws/comments/evv5s1/installing_the_cloudwatch_agent_incur_in/
---
I only  want the mem\_\* metrics, are they contabilized as basic metrics?
## [7][Improving AWS Performance for the Future](https://www.reddit.com/r/aws/comments/ew36qk/improving_aws_performance_for_the_future/)
- url: https://www.ibexlabs.com/improving-aws-performance-for-the-future/
---

## [8][AWS Budget per user](https://www.reddit.com/r/aws/comments/evt4yj/aws_budget_per_user/)
- url: https://www.reddit.com/r/aws/comments/evt4yj/aws_budget_per_user/
---
Can you use AWS Budget per user of an account? I have been trying to google this, but all documentation on AWS Budget seems to be about per account budgeting, not per user.

Use case: Corporations may want to create a separate account for development. To prevent their developers from running expensive services, you can set a limit *per developer*, say $100, that once they exceed that limit, they cannot use any service anymore.
## [9][Use CloudFront SSL with React website on S3](https://www.reddit.com/r/aws/comments/ew0av4/use_cloudfront_ssl_with_react_website_on_s3/)
- url: https://www.reddit.com/r/aws/comments/ew0av4/use_cloudfront_ssl_with_react_website_on_s3/
---
I have a react site on S3, an SSL certificate on CloudFront and a domain name on GoDaddy.  The certificate is active and deployed.

But when I got to [https://example.com](https://example.com) nothing is reached.  How do I connect the SSL to the website for https?
## [10][Serverless framework unable to setup lambda trigger](https://www.reddit.com/r/aws/comments/evwgmt/serverless_framework_unable_to_setup_lambda/)
- url: https://www.reddit.com/r/aws/comments/evwgmt/serverless_framework_unable_to_setup_lambda/
---
I am using the serverless framework with Node.js to create and deploy my Lambdas. It works great except now after recreating my Lambda (new account) I am unable to set up a trigger using the serverless.yml file

&amp;#x200B;

Here is part of my serverless.yml file which shoud be enough to setup the trigger.

&amp;#x200B;

    functions:
      email:
        handler: handler.email
        memorySize: 128 # in MB
        events:
          - sqs: 
            arn: arn:aws:sqs:us-east-1:&lt;account number&gt;:email_queue_${opt:stage}
            batchSize: 1

I've checked and currently there is 1 message in the queue and the queue arn matches

&amp;#x200B;

Here is the documentation on it

[https://serverless.com/framework/docs/providers/aws/events/sqs/](https://serverless.com/framework/docs/providers/aws/events/sqs/)
