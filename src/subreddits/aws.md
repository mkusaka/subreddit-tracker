# aws
## [1][Managing AWS accounts for an organization: AWS Control Tower v.s. Landing Zone v.s. Terraform?](https://www.reddit.com/r/aws/comments/hwwy1u/managing_aws_accounts_for_an_organization_aws/)
- url: https://www.reddit.com/r/aws/comments/hwwy1u/managing_aws_accounts_for_an_organization_aws/
---
Hello,

I'm looking into the best way to manage a relatively small number (&lt;50) of AWS accounts in an organization. AWS is trying to push for Control Tower or Landing Zone, but I'm not sure of the added-value compared to Terraform.

More precisely, I've PoC'ed managing the following with Terraform:

* Master account: Organization, children accounts, OUs, service control policies, tagging policies
* Children accounts: SAML IdP, IAM roles, CloudTrail...

I have a setup where I use 1 Terraform module for children accounts, allowing the AWS Terraform to assume a role in the child account to provision the right resources in here - so far it works like a charm and seems to be automation-friendly.

Why would I want to go for Control Tower / Landing Zone? Any experience with these?

Thank you
## [2][[Announcement] Lightsail now offers a CDN to optimize your website delivery](https://www.reddit.com/r/aws/comments/hwplbo/announcement_lightsail_now_offers_a_cdn_to/)
- url: https://www.reddit.com/r/aws/comments/hwplbo/announcement_lightsail_now_offers_a_cdn_to/
---
Hello redditors,

Today we announced Lightsail CDN, a new feature on Amazon Lightsail that allows you to create CDN distributions to accelerate the delivery of your website’s content to users all over the world. You can create Lightsail CDN distributions with Lightsail instances or load balancers as origins. We even made it easy to setup CDN for common use cases like Wordpress, Static sites with predefined configurations. So, you can enable CDN for your applications with just a few clicks.

We offer an introductory plan that's free for 12 months to make it easy to get started. Give it a try!

What’s New post: [https://aws.amazon.com/about-aws/whats-new/2020/07/amazon-lightsail-now-offers-cdn-distributions-to-accelerate-content-delivery/](https://aws.amazon.com/about-aws/whats-new/2020/07/amazon-lightsail-now-offers-cdn-distributions-to-accelerate-content-delivery/)

Blog post: [https://aws.amazon.com/blogs/compute/improving-website-performance-with-lightsail-content-delivery-network/](https://aws.amazon.com/blogs/compute/improving-website-performance-with-lightsail-content-delivery-network/)

Pricing -  [https://aws.amazon.com/lightsail/pricing](https://aws.amazon.com/lightsail/pricing)

*Here's a general discussion thread on Lightsail-*  [*https://www.reddit.com/r/aws/comments/ftpvbb/heard\_of\_lightsail\_use\_lightsail\_let\_us\_know\_your/*](https://www.reddit.com/r/aws/comments/ftpvbb/heard_of_lightsail_use_lightsail_let_us_know_your/)

We’re excited for you to try it out and would love to hear your feedback.

Thanks!
## [3][Prometheus: yet-another-cloudwatch-exporter — collecting AWS CloudWatch metrics](https://www.reddit.com/r/aws/comments/hwyke6/prometheus_yetanothercloudwatchexporter/)
- url: https://medium.com//prometheus-yet-another-cloudwatch-exporter-collecting-aws-cloudwatch-metrics-806bd34818a8?source=friends_link&amp;sk=46fc121c07b7de4939eb69f78a1cca24
---

## [4][[Humour] The "UK power cable" that came with our most recent Snowball. Note the amazing disappearing earth...](https://www.reddit.com/r/aws/comments/hwep2d/humour_the_uk_power_cable_that_came_with_our_most/)
- url: https://i.imgur.com/QJW53Hc.jpg
---

## [5][Managing users in AWS Directory Services Microsoft AD](https://www.reddit.com/r/aws/comments/hwxt4w/managing_users_in_aws_directory_services/)
- url: https://www.reddit.com/r/aws/comments/hwxt4w/managing_users_in_aws_directory_services/
---
According to the docs, to manage users in an AWS Directory Services Microsoft AD domain we need a windows machine that can join the new domain, and then use the AD Tools within windows to manage users and groups.

&amp;#x200B;

Is there an alternative to this?

We really really don't want to run any EC2 windows servers for *reasons*.

So, looking for some way of managing users and groups outwith windows - any other AD tools that would do the job - any CLI options?
## [6][How do you provide read access on AWS Console and write access through Terraform provisions to enforce security and transparency among developers?](https://www.reddit.com/r/aws/comments/hx1jlq/how_do_you_provide_read_access_on_aws_console_and/)
- url: https://www.reddit.com/r/aws/comments/hx1jlq/how_do_you_provide_read_access_on_aws_console_and/
---
It is generally best practice to maintain your infrastructure through Terraform state if you've provisioned through it in the first place because otherwise any manual provisions, directly on AWS console per se, will cause a mismatch from the state and it will cause troubles going forward.

So how do we give developers read access on AWS Console but write access through Terraform provisions? This enforces consistency such that all infrastructure changes require an explicit transparent PR/merge.

Please let me know if there's a better practice that I'm not aware of.
## [7][V2.0 of ElectricEye is now ready!](https://www.reddit.com/r/aws/comments/hwq4km/v20_of_electriceye_is_now_ready/)
- url: https://www.reddit.com/r/aws/comments/hwq4km/v20_of_electriceye_is_now_ready/
---
Hey folks, a few months back I posted about a new tool I wrote when I left AWS to add additional checks and capabilities into AWS Security Hub called ElectricEye:  [https://github.com/jonrau1/ElectricEye](https://github.com/jonrau1/ElectricEye) 

It has come a long way since I first wrote it, featuring more 3rd party integrations, over 200 checks, added extra compliance requirements to the checks, aligns to the latest versions of the AWS Security Finding Format and the latest rendition makes it much more efficient. This is in a very large part due to the awesome crew from DisruptOps (no , they didn't pay me to say this). I got to work with them closely when I was at AWS and they really made ElectricEye into an awesome tool.

New features include the addition of tests, a developer guide, a new integration with DisruptOps, addition of more checks and a ton of efficiencies such as less AWS API calls, the addition of a CLI and the ability to export findings into CSV (in talks about other formats too, but, you can convert CSV easy enough to JSON or something).

As always, here is the language from my repo:

Continuously monitor your AWS services for configurations that can lead to degradation of confidentiality, integrity or availability. All results will be sent to Security Hub for further aggregation and analysis.

* 100% native Security Hub integration &amp; 100% serverless with full CloudFormation &amp; Terraform support in AWS Commercial and GovCloud Regions
* 220+ security &amp; best practice detections including services not covered by Security Hub/Config (AppStream, Cognito, EKS, ECR, DocDB, etc.)
* Detections aligned to NIST CSF, NIST 800-53, AICPA TSC and ISO 27001:2013 using the Compliance.RelatedRequirements field.
* 60+ multi-account SOAR playbooks
* AWS &amp; 3rd Party Integrations: DisruptOps, Config Recorder, Pagerduty, Slack, ServiceNow Incident Management, Jira, Azure DevOps, Shodan and Microsoft Teams
## [8][Why does a subnet with a /28 cider have 11 available ips not 16?](https://www.reddit.com/r/aws/comments/hwqkng/why_does_a_subnet_with_a_28_cider_have_11/)
- url: https://www.reddit.com/r/aws/comments/hwqkng/why_does_a_subnet_with_a_28_cider_have_11/
---
What are the other 5 used for.


I have a lambda useing 3 subnets each with a /28 cider with the Lambda connected to an api gateway.

I’m I’m not sure how many simsltanius invocations of the lambda function this will allow?

Each invocation should last 3 seconds.

I was restricted to using /28 ciders due to limited available ip range.

I discussed the usage with a backend developer on my team and It’s invoked from the front end with 700 users I was told to build it like this and under pressure to do so quickly so have already built it and applied it to dev as of today so backend is able to test there code on the lambda function but I’m still concerned about if the number of invocations allowed at once could potentially cause performance issues?
## [9][[S3] How to manage file accessibility of s3 objects while working via a web app?](https://www.reddit.com/r/aws/comments/hx0trk/s3_how_to_manage_file_accessibility_of_s3_objects/)
- url: https://www.reddit.com/r/aws/comments/hx0trk/s3_how_to_manage_file_accessibility_of_s3_objects/
---
I am working with a flask app where I use s3 to store images(using the boto3 python library), and also dynamically render the images to the front-end after getting the objects(images) from s3. My question is what is the correct way to render back the images to the front-end for files stored in s3.

For eg. if I was storing the files in the server's file system, every image requested via the user would go through the server's validation to check the accessibility of the user and only then give the result. But if I directly give back the object's s3 URL(eg [https://bucketName.s3.amazonaws.com/fileName](https://bucketName.s3.amazonaws.com/fileName)), I won't get to check the user's access(like for those images only belonging to one user).

I want to get the image from the s3 bucket and dynamically create the &lt;img&gt; tag in my HTML  and render the image there.

So how do I approach this problem?

I know the question sounds dumb, but I'm not very familiar with aws and s3. I'll try to clarify the problem if you have any problem understanding the question.
## [10][Amazon Elastic File System doubles per-client throughput](https://www.reddit.com/r/aws/comments/hwq0sy/amazon_elastic_file_system_doubles_perclient/)
- url: https://www.reddit.com/r/aws/comments/hwq0sy/amazon_elastic_file_system_doubles_perclient/
---
[https://aws.amazon.com/about-aws/whats-new/2020/07/amazon-elastic-file-system-increases-per-client-throughput/](https://aws.amazon.com/about-aws/whats-new/2020/07/amazon-elastic-file-system-increases-per-client-throughput/)
