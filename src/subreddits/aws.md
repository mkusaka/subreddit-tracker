# aws
## [1][Amazon EKS now supports Kubernetes version 1.18](https://www.reddit.com/r/aws/comments/jat1y3/amazon_eks_now_supports_kubernetes_version_118/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/10/amazon-eks-supports-kubernetes-version-1-18/
---

## [2][What RDS instance do you use and how many simultaneous connection are you handling?](https://www.reddit.com/r/aws/comments/jb0j36/what_rds_instance_do_you_use_and_how_many/)
- url: https://www.reddit.com/r/aws/comments/jb0j36/what_rds_instance_do_you_use_and_how_many/
---
* just gathering information on user experience with RDS
## [3][Is there a way to generate AWS CLI command from exisiting resources?](https://www.reddit.com/r/aws/comments/jaxf5w/is_there_a_way_to_generate_aws_cli_command_from/)
- url: https://www.reddit.com/r/aws/comments/jaxf5w/is_there_a_way_to_generate_aws_cli_command_from/
---
Hello there,

I was wondering if there is a AWS resource say RDS, is there a way we can generate AWS CLI command to create this exact RDS with all the settings same as the running one?

Thanks
## [4][Serverless Swagger UI for API Gateway](https://www.reddit.com/r/aws/comments/jadw8n/serverless_swagger_ui_for_api_gateway/)
- url: https://www.reddit.com/r/aws/comments/jadw8n/serverless_swagger_ui_for_api_gateway/
---
Quite long but detailed (I hope!) article on how to create an always up-to-date Swagger UI website for API Gateway, protected with Cognito authentication - missing part for sharing API documentation. For TL;DR see the repo with full example - link at the bottom.

[https://betterdev.blog/serverless-swagger-ui-for-api-gateway/](https://betterdev.blog/serverless-swagger-ui-for-api-gateway/)
## [5][Python package to get eks-token (Alternate to aws eks get-token CLI)](https://www.reddit.com/r/aws/comments/jaw7o1/python_package_to_get_ekstoken_alternate_to_aws/)
- url: https://github.com/peak-ai/eks-token
---

## [6][Account Vending Machine](https://www.reddit.com/r/aws/comments/jaw02p/account_vending_machine/)
- url: https://www.reddit.com/r/aws/comments/jaw02p/account_vending_machine/
---
Hi everyone!

So I have a need to create new accounts on the fly, and to bootstrap these new accounts with correct IAM roles so that our application pipelines can deploy their things.

Anyone have a good way to do this that is not customizable control tower? :)

Cheers!
## [7][AWS::SSM::MaintenanceWindowTask output to cloudwatch](https://www.reddit.com/r/aws/comments/jayg6h/awsssmmaintenancewindowtask_output_to_cloudwatch/)
- url: https://www.reddit.com/r/aws/comments/jayg6h/awsssmmaintenancewindowtask_output_to_cloudwatch/
---
Hello, 

If I create a task via the console, I can tick a box to send output to CloudWatch. I can't find anything in the CF documentation about doing this. 

Does anyone know how this can be done?
## [8][Suitable AWS services to run Python script from website](https://www.reddit.com/r/aws/comments/jaxs0w/suitable_aws_services_to_run_python_script_from/)
- url: https://www.reddit.com/r/aws/comments/jaxs0w/suitable_aws_services_to_run_python_script_from/
---
Hi everyone,

I need help to decide which AWS service or maybe also pointed to tutorial how to set up one (since I am new to AWS).

I have a uni project where we need to deploy our script in the cloud and we decided to place it on AWS.

The idea or the process how it is gonna look like would be: 

* The website will have a form that when submitted will also trigger a python script in AWS.
* Python script that will do computation in the cloud and outputting some file output.

That is the step that I need to do, at least for this stretch.

I did some research of what might be a possibility service to use.

I created EC2 instance where, I can call the address from anywhere and outputting "Hello, World" in html format. I think I can call that address to trigger the python script.

However are there better way to do this? I heard a lot about Lambda and S3, but not sure if those  are the one that suits the requirement.

Thank you so much in advance for the help!
## [9][A History of AWS Services](https://www.reddit.com/r/aws/comments/jaghu7/a_history_of_aws_services/)
- url: https://www.awsgeek.com/AWS-History/
---

## [10][AWS CDK for AWS Budget with notification and CF deploy stack](https://www.reddit.com/r/aws/comments/jan2x2/aws_cdk_for_aws_budget_with_notification_and_cf/)
- url: https://www.reddit.com/r/aws/comments/jan2x2/aws_cdk_for_aws_budget_with_notification_and_cf/
---
&amp;#x200B;

https://preview.redd.it/y9t33214lxs51.png?width=2272&amp;format=png&amp;auto=webp&amp;s=07ed922c1eeb87f3ce9b275fba70e49f74eeea42

Hi All,

I wrote an article recently about using AWS Budgets and setting up notifications with thresholds for your AWS budget to keep an eye on costs. Which something we should all do but most of the examples out there use CloudWatch + SNS. 

I'm a big fan of CDK but there is often a lack of real world examples. So all of my examples will be in CDK but I also provide the CF deploy stack launch button so you can hit the ground running if you just want to try it out.

It's very basic but hopefully helps someone, open source project:

[https://github.com/talkncloud/aws/tree/main/essential-billing](https://github.com/talkncloud/aws/tree/main/essential-billing)

There is a link to the article in the github repo if you want to read up on it.

Let me know if questions or feedback, keen to improve it if people it's useful.

Thanks,

Mick
