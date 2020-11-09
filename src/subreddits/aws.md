# aws
## [1][re:Invent registration is now open](https://www.reddit.com/r/aws/comments/jkenu3/reinvent_registration_is_now_open/)
- url: https://register.virtual.awsevents.com/
---

## [2][AWS Plans $2.8B Data Centers In India](https://www.reddit.com/r/aws/comments/jqgs2d/aws_plans_28b_data_centers_in_india/)
- url: https://www.thetechee.com/2020/11/aws-plans-28b-data-centers-in-india.html
---

## [3][SES Policy Generator and domain migration](https://www.reddit.com/r/aws/comments/jqwfzj/ses_policy_generator_and_domain_migration/)
- url: https://www.reddit.com/r/aws/comments/jqwfzj/ses_policy_generator_and_domain_migration/
---
I am migrating to a new domain. I have a Lambda that will send an email when an error occurs.

This is the current (and working) SES send authorization policy:

	{
		"Version": "2008-10-17",
		"Statement": [
			{
				"Sid": "stmt0000000000000",
				"Effect": "Allow",
				"Principal": {
					"AWS": "AIDARxxxxxxxxxxxxxxxx"
				},
				"Action": [
					"ses:SendEmail",
					"ses:SendRawEmail"
				],
				"Resource": "arn:aws:ses:us-east-1:yyyyyyyyyyyy:identity/zzzzzzzzzz"
			}
		]
	}
	
I don't remember how got the `Principal:AWS` value.  It is `21` characters long.  When I tried using the value from the old policy, I receive this error message:

    Error while applying policy: Policy contains an invalid principal

Since this is not a 12-digit AWS account ID, then I am assuming the is the ARN of an IAM user.  In the `IAM console`, I don't see where this value could live at.  How can I find it or how can I generate a new, valid value other than `*`?

Thanks.
## [4][Anyone ever use SES list/subscription management for a blog?](https://www.reddit.com/r/aws/comments/jqmkm4/anyone_ever_use_ses_listsubscription_management/)
- url: https://www.reddit.com/r/aws/comments/jqmkm4/anyone_ever_use_ses_listsubscription_management/
---
I'm creating a blog for one of my side projects and saw this press release: [Amazon SES now offers list and subscription management capabilities](https://aws.amazon.com/about-aws/whats-new/2020/10/amazon-ses-now-offers-list-and-subscription-management-capabilities/). I'm pretty familiar with AWS already, though I haven't really used SES outside of `sendEmail`. I'd like to **SES's list/subscription management feature for the blog's mailing list** however, there's not much written about this feature outside of the docs, so I want to see if anyone had any positive/negative experiences with SES list management.

From a technical standpoint - the blog is very simple and I'm not using a CMS. I'm just taking markdown files and compiling them into html.

I know there's services that cater to this need and I'm open to them if SES doesn't fit the use case. However, this project pretty much uses AWS for everything (except GitHub), so I'd like to stay in the ecosystem, if possible.
## [5][How to update in bulk G Suite users custom attributes with Google Admin SDK for AWS SAML federation](https://www.reddit.com/r/aws/comments/jqvpvc/how_to_update_in_bulk_g_suite_users_custom/)
- url: https://www.reddit.com/r/aws/comments/jqvpvc/how_to_update_in_bulk_g_suite_users_custom/
---
Hi everyone, after writing an article about how to manually federate G   Suite with AWS an interesting discussion pop-out on how to make these changes programmatically.

After a bit of research, I wrote another article and a GitHub repo to show my findings.

Wanted to share with you and know what do you think about the whole process and if you know of alternate methods to achieve the same thing.

Here are the links:

* [https://blog.pethron.me/how-to-update-in-bulk-g-suite-users-custom-attributes-with-google-admin-sdk](https://blog.pethron.me/how-to-update-in-bulk-g-suite-users-custom-attributes-with-google-admin-sdk)
* [https://github.com/pethron/gsuite-custom-schema-update](https://github.com/pethron/gsuite-custom-schema-update)
## [6][MFA Device Stolen](https://www.reddit.com/r/aws/comments/jqvph7/mfa_device_stolen/)
- url: https://www.reddit.com/r/aws/comments/jqvph7/mfa_device_stolen/
---
Need some advice. My cellphone that I used for MFA on my AWS Root and IAM user login has been stolen.. any idea how to go about regaining access?
## [7][Weird API Gateway Error, but not on console?](https://www.reddit.com/r/aws/comments/jquipp/weird_api_gateway_error_but_not_on_console/)
- url: https://stackoverflow.com/questions/64741872/api-gateway-get-action-works-in-the-console-but-not-in-postman-or-my-app
---

## [8][Weird Lambda behavior - Not executing email module when auto run but executing when when run manually](https://www.reddit.com/r/aws/comments/jqp51o/weird_lambda_behavior_not_executing_email_module/)
- url: https://www.reddit.com/r/aws/comments/jqp51o/weird_lambda_behavior_not_executing_email_module/
---
I have a simple Lambda.py which grabs a "requests" lib response , converts it into csv, sorts it , uploads to S3 and sends an email with attached csv.  

When I login to console and run it manually, the email is being sent, however when its run automatically by event bridge at 8 am, the email is not being sent. Rest of the modules are executed ( I can see the csv uploaded to S3 at 8 am every day).
anyone seen this ? How do I fix this?
## [9][Am I the only one who hates the new AWS console design updates?](https://www.reddit.com/r/aws/comments/jq3cha/am_i_the_only_one_who_hates_the_new_aws_console/)
- url: https://www.reddit.com/r/aws/comments/jq3cha/am_i_the_only_one_who_hates_the_new_aws_console/
---
I rarely use the old console except when I absolutely have to. It was slow and somewhat unappealing to look at. 

AWS just made some major updates to the console and I feel they did so with no user input. At least to me, everything I hate about the old one wasn't addressed or even made worse.

Is this just me or does anyone else feel same?
## [10][Suggestions on how to auto deploy lambdas from Github](https://www.reddit.com/r/aws/comments/jqpuke/suggestions_on_how_to_auto_deploy_lambdas_from/)
- url: https://www.reddit.com/r/aws/comments/jqpuke/suggestions_on_how_to_auto_deploy_lambdas_from/
---
I am building a system for clients to upload their own serverless/lambda JS/TS functions to a lambda from their GH repo. I have this working using CDK/CodePipeline but I don't think stacks are practical for this use case (many lambdas - don't want to deal with stack limits). I am looking at CodeDeploy as it seems this may do the trick. Basically the process should go like this:

\- User enters repo info in my dashboard (I am ok with them needing to create an oauth token)  
\- Press submit and a lambda fetches a builds code

Anyone have some experience with this and some potential pain points or alternative approaches? I am finding little on the subject of this process being automated from a lambda.
## [11][Best email provider for reliable/fast production integration with AWS SES](https://www.reddit.com/r/aws/comments/jqgrrq/best_email_provider_for_reliablefast_production/)
- url: https://www.reddit.com/r/aws/comments/jqgrrq/best_email_provider_for_reliablefast_production/
---
Hey,

I'm setting up a service that receives an email and integrates with SES where SES sends the attachment to S3 and then a lambda function parses the attachment.

I'm wondering if anyone has experience with which email provider to use for this. I was using a namecheap host but it's quite slow and buggy. What's the most reliable/fastest email provider that can integrate with SES?
