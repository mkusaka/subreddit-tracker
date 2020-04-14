# aws
## [1][Is there a relational database with DynamoDB type pricing?](https://www.reddit.com/r/aws/comments/g12o01/is_there_a_relational_database_with_dynamodb_type/)
- url: https://www.reddit.com/r/aws/comments/g12o01/is_there_a_relational_database_with_dynamodb_type/
---
As per question. DynamoDB is fully managed and charged by read/writes, so for a hobby application that gets used say once a day or once per week the monthly cost is virtually NIL. The closest relational database product I could find on AWS is Aurora Serverless, but that is charged per hour so relies on the database going to sleep when not in use to keep the cost down. And that in turn causes a 30 second delay for the database to restart each time the application is accessed. So is that it or is there another relational database product on AWS, which is also charged on read/writes/storage that I can look at for a hobby application that would have a sub $5/month cost?
## [2][Set up an alert for malware detected on AWS WorkSpace](https://www.reddit.com/r/aws/comments/g14rd6/set_up_an_alert_for_malware_detected_on_aws/)
- url: https://www.reddit.com/r/aws/comments/g14rd6/set_up_an_alert_for_malware_detected_on_aws/
---
Hi there, for the life of me I can't figure this one out. I basically want a Cloudwatch alert to be generated if one of my AWS WorkSpaces detects malware. I'm using Windows Defender.

The Windows event ID for malware detected is  **Event ID: 1007**  so essentially all I'd like is an alert in AWS if that event is ever generated.

Can some please let me know the easiest way to do this?

Many thanks :).
## [3][Best practice for upgrading EKS cluster?](https://www.reddit.com/r/aws/comments/g0wp5j/best_practice_for_upgrading_eks_cluster/)
- url: https://www.reddit.com/r/aws/comments/g0wp5j/best_practice_for_upgrading_eks_cluster/
---
Hey everyone. Just wondering what are best practices to upgrade a cluster version. I currently am super behind and its time. Is it best to just create a new cluster completely and deploy to that while moving everything off or upgrading workers and master? Also does anyone have any reliable resources for this? There is virtually nothing I can see so far that has really helped me get a solution.

Any help is welcomed! EKSctl, ETC is an option!
## [4][Assume a role in another account from an assumed role behind autoscaling group](https://www.reddit.com/r/aws/comments/g13lmq/assume_a_role_in_another_account_from_an_assumed/)
- url: https://www.reddit.com/r/aws/comments/g13lmq/assume_a_role_in_another_account_from_an_assumed/
---
I have an autoscaling group of CI/CD servers in ACCOUNT_A which deploys to an s3 server in ACCOUNT_B. I want the servers in ACCOUNT_A to be able to assume a role in ACCOUNT_B to do the upload. I have attached an instance profile to the ACCOUNT_A servers. I can't find a way of doing it as AWS Principals in assumed role policy documents don't support wildcards. It would be fine if the instances never changed as I could add their assumed-role users to the principal in the assume role policy document, however as they change, I can't do that.

The only way I can think of doing it is to add ACCOUNT_A's ID as the principal in ACCOUNT_B but that seems really insecure.

Am I going about this the wrong way or have I missed something?

    $ aws sts assume-role --role-arn arn:aws:iam::ACCOUNT_B:role/deployment --role-session-name test
    An error occurred (AccessDenied) when calling the AssumeRole operation: User: arn:aws:sts::ACCOUNT_A:assumed-role/cicd/INSTANCE_ID is not authorized to perform: sts:AssumeRole on resource: arn:aws:iam::ACCOUNT_B:role/deployment
## [5][Cloud Partner Programs: A Numbers Game](https://www.reddit.com/r/aws/comments/g135wu/cloud_partner_programs_a_numbers_game/)
- url: https://fika.works/blog/cloud-partner-programs/
---

## [6][Cant pass aws sign up using Debit card EPS by BPI](https://www.reddit.com/r/aws/comments/g117kg/cant_pass_aws_sign_up_using_debit_card_eps_by_bpi/)
- url: https://www.reddit.com/r/aws/comments/g117kg/cant_pass_aws_sign_up_using_debit_card_eps_by_bpi/
---
Im having trouble making an account in AWS. I have money on my Debit Card. Is there any workaround for this? Anyone got this problem before? Is it probably because I have an EPS card instead of the MasterCard?
## [7][Q: how do you keep up with AWS new announcements?](https://www.reddit.com/r/aws/comments/g0itdu/q_how_do_you_keep_up_with_aws_new_announcements/)
- url: https://www.reddit.com/r/aws/comments/g0itdu/q_how_do_you_keep_up_with_aws_new_announcements/
---
I am almost sure that at some point you have all joked about how many new products are launched by AWS every year.

As people using AWS products at work, how do you guys keep up with news that are relevant to what you are using?

I found this to be really challenging. I have tried feedly and following up over Twitter but I can't say this was effective.
## [8][Lambda User Experience Survey With Open Results](https://www.reddit.com/r/aws/comments/g151lz/lambda_user_experience_survey_with_open_results/)
- url: https://www.reddit.com/r/aws/comments/g151lz/lambda_user_experience_survey_with_open_results/
---
I'm trying to understand peoples' opinions on AWS Lambda as part of my research on serverless technologies and filling out this survey would help enormously.

The results are open so we can all learn something and I'll post them back here once the responses stop coming in. Plus, you get a programming joke if you complete it.

Thanks everyone! The link is here: [https://forms.gle/FDeg2geN1JgvXnUZ9](https://forms.gle/FDeg2geN1JgvXnUZ9)
## [9][How to automate response to Workspace is ready/available or unhealthy?](https://www.reddit.com/r/aws/comments/g14ee5/how_to_automate_response_to_workspace_is/)
- url: https://www.reddit.com/r/aws/comments/g14ee5/how_to_automate_response_to_workspace_is/
---
We're close to being able to deploy workspaces to a bunch of remote workers but are hitting a snag on  automate the provisioning of it -

&amp;#x200B;

* We're looking to provision a bunch of workspaces using a custom image. 
* We have a simple approval workflow setup in servicenow
* We know how to kick off the createworkspaces part - yay!
* That works most of the time, but fails enough (workspace goes to unhealthy) to be annoying
* so, how do we know when a build is unhealthy or when it successful in an automated manner?
* We were going to use describeworkspace and notify the client when its "available"
* This is OK, but is there a better way than continually polling and then reacting if it's unhealthy?
* Can we use cloudwatch, sqs, lambda, etc.? So, Workspaces sends a healthy or unhealthy event to CloudWatch, on "unhealthy" it kicks off a lambda function that destroys the workspaces and rebuilds it... on healthy (the first time) sends an email to the user saying the unit is ready.

I've been googling around and the available docs are confusing.
## [10][Question about AWS Security Groups and IPv6](https://www.reddit.com/r/aws/comments/g0vc9n/question_about_aws_security_groups_and_ipv6/)
- url: https://www.reddit.com/r/aws/comments/g0vc9n/question_about_aws_security_groups_and_ipv6/
---
Here's a curious question.  I ran into a scenario where a friend spun up an RDS (we split an amazon account).  I typically lock down resources by nat'd IP address, just to play it safe.  I saw his was IPv6, came up with a rule to allow only is... and no dice.

Came across this document:

https://aws.amazon.com/premiumsupport/knowledge-center/rds-ipv6/

Looks like IPv6 connections are blocked.  Okay, fair enough.  Did a bit of testing.  Tried a few workarounds... but eventually I added back in 0.0.0.0/0 to create a few tables for him.

Friend: "Hey it works.  What did you do."

Me: "I added a rule for 0.0.0.0/0."

Friend: "Uh... what?  Why did that work."

I have no idea why that worked.  My only thoughts are:

* Is local connection is IPv4.  But his WAN connection is IPv6.  So AWS is respecting his local IPv4 address.  Which kind of breaks my understanding of security rules.

* Something along the way is translating IPv6 to IPv4.  Will need to do a bit of research for this one.

I'm a bit stumped on this issue.  It's not something I've come across before.  Anyone know why his connection to rds works?
