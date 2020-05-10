# aws
## [1][Is using AWS Organisations the best way structure side projects?](https://www.reddit.com/r/aws/comments/ggza9m/is_using_aws_organisations_the_best_way_structure/)
- url: https://www.reddit.com/r/aws/comments/ggza9m/is_using_aws_organisations_the_best_way_structure/
---
I've got a bunch of projects and ideas I'm working on. Is using a single master account that I log into and then creating Orgs from inside there the best way to structure things rather than set up the same 'org' level billing alerts and accounts each time?
## [2][Best Practice Guide for Setting a Company Up on AWS For The First Time](https://www.reddit.com/r/aws/comments/ggebgn/best_practice_guide_for_setting_a_company_up_on/)
- url: https://www.reddit.com/r/aws/comments/ggebgn/best_practice_guide_for_setting_a_company_up_on/
---
I am a senior AWS DevOps engineer and have just gotten a contract with a company to move their private cloud infrastructure to AWS. This company has a DevOps and TechSecOps team but no AWS experience. Aside from the migration itself, can anyone send me a link to a webpage, book, or other, that details a step-by-step best practices of setting a company up on AWS please? This company is ISO-27001 security compliant. I am thinking of stuff like this:

\- Create AWS Organizations, ie have a parent account and a couple of initial child accounts

\- Create a couple of IAM users

\- Turn on Cloud Trail

\- Configure AWS Config

\- Configure Billing Alerts

\- Maybe configure VPC endpoints?

\- Create a few Cost Allocation Tags

\- Create a standard VPC configuration with public and private subnets 

\- Anything else I have missed

&amp;#x200B;

And finally getting sample Terraform code to do all the above. Does anyone know of a comprehensive resource with at least all the above information as well as other basic stuff? If so, I'd be grateful for any information/suggestions/watch outs etc.

Thanks.
## [3][Is WEBDAV request allowed via AWS CloudFront?](https://www.reddit.com/r/aws/comments/gguwym/is_webdav_request_allowed_via_aws_cloudfront/)
- url: https://www.reddit.com/r/aws/comments/gguwym/is_webdav_request_allowed_via_aws_cloudfront/
---
Hello All,

&amp;#x200B;

WebDAV request (Port 443) is not working for me to pass through AWS CloudFront.

WEBSITE access via the INTERNET on port 443 with request flowing through CloudFront is working fine.

Looking for suggestions if WEBDAV is allowed.

&amp;#x200B;

Thanks in advance for advice.
## [4][Where are AWS Published Lambda Layers, I am looking for Numpy, Scipy and Pandas Lambda Layer published by AWS? Are there one?](https://www.reddit.com/r/aws/comments/gh1cge/where_are_aws_published_lambda_layers_i_am/)
- url: https://www.reddit.com/r/aws/comments/gh1cge/where_are_aws_published_lambda_layers_i_am/
---
Hello AWS Fellows, I have been searching for AWS Published layers with Numpy, SciPy, and Pandas, are there layers published by AWS? or do we have to create one ourselves? In the documentation for Lambda layers, it is mentioned that you can either create your own, use one provided by other customers or published by AWS but sadly I couldn't find the AWS ones.   


Would be great gesture if you can guide me through.   


Thanks
## [5][(amplify) How to use share data with ownerField](https://www.reddit.com/r/aws/comments/ggwqrz/amplify_how_to_use_share_data_with_ownerfield/)
- url: https://www.reddit.com/r/aws/comments/ggwqrz/amplify_how_to_use_share_data_with_ownerfield/
---
How do I structure a data type to be accessible by multiple users based on their email addresses? 

Say for example a todo-app; I want to be able to invite additional email addresses to have access to the data. I use cognition user pools with email signup. 

Something like this where the `owners` field would be an array with email addresses but I can't wrap my head around it yet

```
type SharableData @model @auth(rules: [{ allow: owner, ownerField: "owners" }]) {
  owners: [String!]!
}
```
## [6][Remote desktop software on ec2 instance for access?](https://www.reddit.com/r/aws/comments/ggrj10/remote_desktop_software_on_ec2_instance_for_access/)
- url: https://www.reddit.com/r/aws/comments/ggrj10/remote_desktop_software_on_ec2_instance_for_access/
---
Hi Still learning AWS

Would team viewer or log me in be safe to put on a ec2 instance?  I know RDP is not encrypted would ideally not like to use it.

I am looking at two ways to access my ec2

1) Setup a site to site vpn between my  router and VPC and then access it via rdp 

(would a remote vpn setup work here? rather then a site to site)

&amp;#x200B;

2) Setup log me in( or team viewer) to the ec2 instance.

&amp;#x200B;

Would either of those be acceptable?
## [7][Question about signing in through remote desktop](https://www.reddit.com/r/aws/comments/ggut0c/question_about_signing_in_through_remote_desktop/)
- url: https://www.reddit.com/r/aws/comments/ggut0c/question_about_signing_in_through_remote_desktop/
---
Welp, I messed up pretty bad I guess. Im taking a windows 2016 server class and we use AWS for out machines. I changed my administrator password through AD, and go to sign out and log back in but now whenever I log in through remote desktop its saying "User has to change password on next login" Well this is an issue I cant get to the windows login screen to change my password. I know this is a very stupid mistake and I really don't want to message my teacher telling her I fucked up this badly, is there anything I can do to sign into my machine?
## [8][Data Pipeline Scheduling](https://www.reddit.com/r/aws/comments/ggt2fa/data_pipeline_scheduling/)
- url: https://www.reddit.com/r/aws/comments/ggt2fa/data_pipeline_scheduling/
---
Hello, 

&amp;#x200B;

I want to automate the starting of some of their instances from monday to friday, so I set up a simple data pipeline on a daily schedule that runs an AWS CLI command similar to this:

`test $(date +"%u") -lt 6 &amp;&amp; aws ec2 stop-instances --instance-ids &lt;instance&gt; --region us-east-1`

When I checked it today, I saw there was an ERROR message, which I assume that it is that the day check failed (since it's Saturday) so it couldn't complete successfully. Will that stop the pipeline from running tomorrow or Monday?

Thanks for your time
## [9][AWS forecast ignores feature-attribute when making a prediction](https://www.reddit.com/r/aws/comments/ggpd0v/aws_forecast_ignores_featureattribute_when_making/)
- url: https://www.reddit.com/r/aws/comments/ggpd0v/aws_forecast_ignores_featureattribute_when_making/
---
Hey  fellows, 

I defined a TARGET dataset with 4 attributes (item\_id, store, date, demand) and need to forecast each item per day per shop. The model trains well, but when casting a prediction, the aws forecast service aggregates all items per day with disregards to the store attribute.

Has anyone ever had the same issue with AWS Forecast...?
## [10][ECS task cannot resolve hostnames](https://www.reddit.com/r/aws/comments/ggpaxs/ecs_task_cannot_resolve_hostnames/)
- url: https://www.reddit.com/r/aws/comments/ggpaxs/ecs_task_cannot_resolve_hostnames/
---
A colleague reached out me recently for help troubleshooting an issue with ECS, where when the host EC2 instance is in a certain VPC/subnet, the tasks cannot resolve hostnames when attempting to make requests to our company website, which is behind Akamai. This is a problem because the ECS cluster is running automated tests, and needs to be able to resolve our company site.

To be clear, the host EC2 instance has no issues resolving the same hostnames, just the containers.

What's strange is that an identical cluster running in another VPC/subnet does not have the same problem. Unfortunately, for reasons known only to our networking team, that subnet has a tiny CIDR block allocated, so the cluster cannot spin up enough tasks to get through the automation in a reasonable time.

I'm not much of a docker expert, so I haven't got a clear idea of how to debug this. The cluster is running in network\_mode "host" if that's useful information.
