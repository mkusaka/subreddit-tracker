# aws
## [1][PSA: the AWSLambdaExecute policy grants lambdas permission to put and get any object in any S3 bucket](https://www.reddit.com/r/aws/comments/fp6moa/psa_the_awslambdaexecute_policy_grants_lambdas/)
- url: https://www.reddit.com/r/aws/comments/fp6moa/psa_the_awslambdaexecute_policy_grants_lambdas/
---
I applied this policy to all my lambdas, because I thought it's an easy way to grant Cloudwatch permissions.

I was shocked when I later realised what I'd actually done.

I spend so long fine tuning most permissions. It's annoying to see a giant asterisk bulldozing through my IAM security measures.
## [2][Amazon Managed Cassandra Service now helps you automate the creation and management of resources by using AWS CloudFormation](https://www.reddit.com/r/aws/comments/foyp8f/amazon_managed_cassandra_service_now_helps_you/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/03/amazon-managed-cassandra-service-now-helps-you-automate-the-creation-and-management-of-resources-by-using-aws-cloudformation/
---

## [3][Migrating from one account to another](https://www.reddit.com/r/aws/comments/fpaz9w/migrating_from_one_account_to_another/)
- url: https://www.reddit.com/r/aws/comments/fpaz9w/migrating_from_one_account_to_another/
---
Does amazon provide a way to migrate an AWS account? I started building on a personal account and now need to migrate my entire setup to my work sponsored account. I’ve got ec2 machines, s3, Efs, workspaces.
## [4][S3 open source solution to transparently aggregate objects](https://www.reddit.com/r/aws/comments/fparmg/s3_open_source_solution_to_transparently/)
- url: https://www.reddit.com/r/aws/comments/fparmg/s3_open_source_solution_to_transparently/
---
Imagine we got a solution/apps that stores petazillions of small objects on s3. 

It obviously will cost us a lot in put and get and the performance will no bet good.

It can be acceptable for us to delay a bit the write and to store objects grouped. (same for read).

My question is there any existing solution that can do this job for me ? 

akka I know I can handle this in my code but it's a lot of work, so if something can do it for me I take.
## [5][Postgres index type for uuid[] array, supported on amazon rds?](https://www.reddit.com/r/aws/comments/fpa798/postgres_index_type_for_uuid_array_supported_on/)
- url: https://www.reddit.com/r/aws/comments/fpa798/postgres_index_type_for_uuid_array_supported_on/
---
PostgreSQL doesn't include an index type for uuid array columns. But it is possible to create an index type that understands how to compare in an array of uuids using operators.
This requires superuser privileges (to create an operator). I was wandering if this was possible if I host my pg database in amazon rds.
Kind regards
## [6][Savings Plan recommendations for Linked Accounts](https://www.reddit.com/r/aws/comments/fp7sfz/savings_plan_recommendations_for_linked_accounts/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/03/aws-cost-explorer-now-offers-savings-plans-recommendations-for-member-linked-accounts/
---

## [7][CLI install help](https://www.reddit.com/r/aws/comments/fpaww1/cli_install_help/)
- url: https://www.reddit.com/r/aws/comments/fpaww1/cli_install_help/
---
I'm upgrading my CLI from v1 to v2 on my OSX (10.13.6) machine. I thought I had v1 completely removed and v2 installed properly. 

When I go to a terminal window and type `which aws` I get back:

`/usr/local/bin/aws` 

However, when I type `aws --version`, I get back: 

`-bash: /Library/Frameworks/Python.framework/Versions/3.7/bin/aws: No such file or directory` which appears to me as the old directory where I had v1 CLI installed.

What can I do here to fix this and point `aws` to go look in the `/usr/local/bin` directory? 

Thanks!
## [8][Serverless .Net API Design](https://www.reddit.com/r/aws/comments/fp5is4/serverless_net_api_design/)
- url: https://www.reddit.com/r/aws/comments/fp5is4/serverless_net_api_design/
---
I'm trying to work out the best way to architect the back end of a React website to be serverless where possible and I'm having a hard time finding real world examples of how people are doing it.

[Amazon's own documentation](https://aws.amazon.com/blogs/developer/net-core-3-0-on-lambda-with-aws-lambdas-custom-runtime/) seems to upload an entire webAPI project into one Lambda function and uses API Gateway proxying to direct traffic inside it. While this makes it very easy to use, deploy and develop, it feels like a pretty bad way to do it given lack of control over the functions themselves.

It feels like a microservice approach would be better to allow each controller to be it's own Lambda function with API Gateway directing as appropriate. The problem here is managing access while developing, it seems that Docker is usually the way to go to run the microservices, but since the functions would be written as Lambda handlers I can't just change the API URL at run time in my front end.

Does anyone have experience with this, or know of any real world examples that go into some depth around actually developing and managing this process? Most blogs and walkthroughs just talk about the very basics of setting it up and don't integrate the entire process.
## [9][using awscli ssm send-command to run Powershell script](https://www.reddit.com/r/aws/comments/fowfkt/using_awscli_ssm_sendcommand_to_run_powershell/)
- url: https://www.reddit.com/r/aws/comments/fowfkt/using_awscli_ssm_sendcommand_to_run_powershell/
---
I'm trying to use aws ssm to run a powershell script, the script takes a parameter which is a file name.  The awscli command:

    aws ssm send-command --timeout-seconds 60 --instance-ids i-blahblah --document-name AWS-RunPowerShellScript --parameters commands=["c:\myscript.ps1 c:\filename"]

However script.ps1 fails when there's a space in filename.  If I'm running it on the remote server I can just quote like "c:\file name", and my script.ps1 would work.  But using the quote in the command resulted in

    aws ssm send-command --timeout-seconds 60 --instance-ids i-blahblah --document-name AWS-RunPowerShellScript --parameters commands=["c:\myscript.ps1 'c:\file name'"]

    Error parsing parameter '--parameters': Expected: ',', received: ''' for input:
    commands=[c:\myscript.ps1 c:\file name]
                                                     ^

How should I format --parameters?

EDIT: Found a workaround, backtick (\`) escapes space in powershell, so I can do:

    --parameters commands=["c:\myscript.ps1 c:\file` name"]
## [10][Does anyone have experience with the average implementation time of bug fixes in AWS?](https://www.reddit.com/r/aws/comments/fp8wa8/does_anyone_have_experience_with_the_average/)
- url: https://www.reddit.com/r/aws/comments/fp8wa8/does_anyone_have_experience_with_the_average/
---
Hi, I'm trying to configure AWS SES to send emails to S3 but failed because of a bug in SES (SES cannot deliver emails to a bucket with Object Lock).

I contacted AWS Support and they confirmed it and will implement a fix but cannot give me an ETA for that. 

Does anyone know if I should expect weeks, months or even longer here based on their previous experience?
