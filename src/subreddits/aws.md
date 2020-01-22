# aws
## [1][EKS price reduction](https://www.reddit.com/r/aws/comments/erz5l4/eks_price_reduction/)
- url: https://aws.amazon.com/blogs/aws/eks-price-reduction/
---

## [2][Difference between CDK, SAM and Serverless](https://www.reddit.com/r/aws/comments/es9yld/difference_between_cdk_sam_and_serverless/)
- url: https://www.reddit.com/r/aws/comments/es9yld/difference_between_cdk_sam_and_serverless/
---
I am new to serverless and I'm trying to use a framework for my first production application.  I have used the "Serverless" framework for a few projects but then I came across SAM and CDK and now I can't really make out what does what.
## [3][My Organization &amp; Created accounts problem](https://www.reddit.com/r/aws/comments/esbbae/my_organization_created_accounts_problem/)
- url: https://www.reddit.com/r/aws/comments/esbbae/my_organization_created_accounts_problem/
---
Hey guys, I have the following issue. We created an Organization in AWS, after which we created accounts for it, based on real email addresses, as outlined. 

Now, the problem is when I log into my account (the one that was created via My Organization) I can create IAM users, groups and policies - but none of the other accounts that got created via My Organization can see them.

What am I missing here - is the Master account supposed to create the IAM roles and users, or how do I share them with the rest of the accounts that got created through My Organization?
## [4][Creating a static website with S3, with a custom domain name purchased from elsewhere.](https://www.reddit.com/r/aws/comments/esbad8/creating_a_static_website_with_s3_with_a_custom/)
- url: https://www.reddit.com/r/aws/comments/esbad8/creating_a_static_website_with_s3_with_a_custom/
---
I've figured out how to create a static web page with AWS S3, and now I'm looking to attach a custom domain name to it. All of the domains on route 53 are very expensive so I want to buy from elsewhere. I also don't want my S3 bucket to have to be the same name as my domain name, as this will limit my bucket name choice. 

Is this possible? Does AWS allow you to attach go Daddy domains to it's resources?

U don't really know much about cloud front but I assume I could use it as part of my solution  My other thought was making a small EC2 instance and using it as a reverse proxy, so my bucket name doesn't have to match the domain name.
## [5][Does anyone have a template/tutorial for how to setup HTTPS for EC2 instances with CloudFormation?](https://www.reddit.com/r/aws/comments/esa125/does_anyone_have_a_templatetutorial_for_how_to/)
- url: https://www.reddit.com/r/aws/comments/esa125/does_anyone_have_a_templatetutorial_for_how_to/
---
I have these links but I don't know how to interpret them. 

https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-redirectconfig.html 
https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-defaultactions.html

Unfortunately with my account type it will only work by using templates on CloudFormation.
## [6][Only get existing resources with ResourceGroupTaggingAPI](https://www.reddit.com/r/aws/comments/es9kzw/only_get_existing_resources_with/)
- url: https://www.reddit.com/r/aws/comments/es9kzw/only_get_existing_resources_with/
---
Is there a way to only get the existing resources from the ResourceGroupTaggingAPI, instead of also including the previous resources, without actually checking? We're using the ResourceGroupTaggingAPI to find all RDS database instances which match a certain tag, but it's also returning all the deleted resources which means having to sanitize the list by checking whether the resources actually exist or not.

Of course a workaround would be to establish a blacklist of database identifiers... But there has to be a more elegant solution?
## [7][Manage source AMI with CodeDeploy](https://www.reddit.com/r/aws/comments/es9hni/manage_source_ami_with_codedeploy/)
- url: https://www.reddit.com/r/aws/comments/es9hni/manage_source_ami_with_codedeploy/
---
Hello, 

We use an AMI as part of our launch configuration for our autoscaling group. I know we can update the auto-scaling servers using CodeDeploy, but can we update the source AMI too? Or is there a better tool for that? The code is in github, and managed by TFS.
## [8][Advanced AWS CLI JMESPath Query Tricks](https://www.reddit.com/r/aws/comments/eruq96/advanced_aws_cli_jmespath_query_tricks/)
- url: https://opensourceconnections.com/blog/2015/07/27/advanced-aws-cli-jmespath-query/
---

## [9][How Do I Get a Task Definition Console?](https://www.reddit.com/r/aws/comments/es50c8/how_do_i_get_a_task_definition_console/)
- url: https://www.reddit.com/r/aws/comments/es50c8/how_do_i_get_a_task_definition_console/
---
I'm a beginner to AWS, so this might be a basic question, but I must not be Googling the right things to find my answer.

I have a task definition running in a console, how would I ssh into that task definition or otherwise get connect to a console on the image? I have the aws cli setup, but can't find any commands in there related to connecting to a task.
## [10][Need help with Cloudfront, S3 https redirection error 504](https://www.reddit.com/r/aws/comments/es7bw6/need_help_with_cloudfront_s3_https_redirection/)
- url: https://www.reddit.com/r/aws/comments/es7bw6/need_help_with_cloudfront_s3_https_redirection/
---
Hello All,

I am hoping that I could get some suggestions on how to go about fixing my issue with Cloudfront, S3 https redirection.  Background is I have followed the instructions provided here ([https://simonecarletti.com/blog/2016/08/redirect-domain-https-amazon-cloudfront/](https://simonecarletti.com/blog/2016/08/redirect-domain-https-amazon-cloudfront/)) to setup by https redirection.  When it came to verifying my changes, I got error 504 below.

[504 Error](https://preview.redd.it/js3nyh4cj9c41.png?width=1218&amp;format=png&amp;auto=webp&amp;s=cc07a7d21d948b63f5bb4c4ae99fcae24e090b55)

* The certificate I used is AWS issued.
* I tried opening the s3 website URL itself and I can get the redirection working fine.

I am hoping someone else has seen this error before and managed to fixed it.

Thank you in advance.
