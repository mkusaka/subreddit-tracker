# aws
## [1][Cortex - Open source alternative to SageMaker for model serving](https://www.reddit.com/r/aws/comments/fkmpnr/cortex_open_source_alternative_to_sagemaker_for/)
- url: https://github.com/cortexlabs/cortex
---

## [2][Is Workmail still being developed?](https://www.reddit.com/r/aws/comments/fko8vm/is_workmail_still_being_developed/)
- url: https://www.reddit.com/r/aws/comments/fko8vm/is_workmail_still_being_developed/
---
We use Workmail as our companies email service. It is generally ok, but it doesn't seem to have had any new features since it's release 3 years ago. The interface is clunky but we expected this to improve. 

We are concerned that there seems to be a lack of development on the platform. It is very basic but we pay almost the same amount as we would pay for Gmail. 

Also, we receive calendar invites for AWS events, from AWS, and the only options we have are Outlook, Google or iCal.  No mention of Workmail or any web based integration to their events. This just comes across as a product that AWS doesn't even recognise.

Does anybody use Workmail and is there any reason not to move onto Gmail?
## [3][Automated Terraform Deployments to AWS with Github Actions](https://www.reddit.com/r/aws/comments/fkcqi3/automated_terraform_deployments_to_aws_with/)
- url: https://medium.com/@dnorth98/automated-terraform-deployments-to-aws-with-github-actions-c590c065c179
---

## [4][Understand pod networking in your eks!](https://www.reddit.com/r/aws/comments/fkmm7i/understand_pod_networking_in_your_eks/)
- url: https://yashmehrotra.com/post/2020-03-16-case-of-missing-packet/
---

## [5][Using SSM for patching (this is the most confusing patch interface I've ever seen)](https://www.reddit.com/r/aws/comments/fki1yl/using_ssm_for_patching_this_is_the_most_confusing/)
- url: https://www.reddit.com/r/aws/comments/fki1yl/using_ssm_for_patching_this_is_the_most_confusing/
---
Hello all,

 I've spend hours trying to understand AWS's patching interface with SSM. I get the concept of:

* patch groups
* maintenance windows
* patch baselines

but for the love of god how do you actually put this stuff together?  The interface is absolutely horrible. 

Does anyone have a good resource (yes I've read the docs) on how to actually go about doing this? Even our AWS engineers at work are struggling to put this all together. 

&amp;#x200B;

Kind regards
## [6][CodePipeline to deploy an AWS CloudFormation stack](https://www.reddit.com/r/aws/comments/fklulr/codepipeline_to_deploy_an_aws_cloudformation_stack/)
- url: https://www.reddit.com/r/aws/comments/fklulr/codepipeline_to_deploy_an_aws_cloudformation_stack/
---
hi all, while I am also trying to figure this out, I thought I will ask the helpful peeps on reddit.

I am attempting to deploy a cloudformation stack using code pipeline but my stack takes some parameters which I am not sure how to integrate with code pipeline.

Those parameters are simple user inputs which can filled when deploying just using cloudformation.

Is there a way I can tell code pipeline values to those parameters?
## [7][How do you handle promoting lambda functions to different environments with respect to their attached environment variables?](https://www.reddit.com/r/aws/comments/fko1r2/how_do_you_handle_promoting_lambda_functions_to/)
- url: https://www.reddit.com/r/aws/comments/fko1r2/how_do_you_handle_promoting_lambda_functions_to/
---
Lambda does not support per-alias environment variables, which makes promoting a single through different environments very difficult. We have opted for simply running two functions, with the same code, but different env vars.

I've heard of other solutions like parsing the alias from the context and then including a certain set of secrets depending on the alias.

How have you guys solved this limitation?
## [8][Can S3 Select be used on hidden files?](https://www.reddit.com/r/aws/comments/fkjxdr/can_s3_select_be_used_on_hidden_files/)
- url: https://www.reddit.com/r/aws/comments/fkjxdr/can_s3_select_be_used_on_hidden_files/
---
The title says it all really.  I've only discovered S3 Select tonight after struggling to find a solution to a particular ETL problem I've run into.  It works great on the latest version of the file,  but I can't seem to figure out how to use it to read older versions.  I'm currently using boto3 and the select\_object\_content method doesn't appear to take VersionId as a parameter.  Is there a way around this?  I'm fairly new to S3 so apologies if I'm missing something obvious.
## [9][Can someone explain the real differences between Amzn EFS and Azure Files?](https://www.reddit.com/r/aws/comments/fkh7c9/can_someone_explain_the_real_differences_between/)
- url: https://www.reddit.com/r/aws/comments/fkh7c9/can_someone_explain_the_real_differences_between/
---
And additionally, why should I NOT choose one over the other? What are the shortcomings of each? I've heard EFS isn't considered enterprise. Why? Thank you!
## [10][Has anyone done auto-tagging AWS resources by their IAM username across all AWS accounts in the Org? Currently, I'm able to auto-tag resources in one account, I need for all the accounts.](https://www.reddit.com/r/aws/comments/fkmpuy/has_anyone_done_autotagging_aws_resources_by/)
- url: https://www.reddit.com/r/aws/comments/fkmpuy/has_anyone_done_autotagging_aws_resources_by/
---
We currently have around 90+ AWS accounts in the Org, is there any way I can achieve Auto-tagging across all regions in all the accounts?
