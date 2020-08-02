# aws
## [1][Should I use serverless or Amplify + CloudFront to host a SPA + REST backend?](https://www.reddit.com/r/aws/comments/i28r3o/should_i_use_serverless_or_amplify_cloudfront_to/)
- url: https://www.reddit.com/r/aws/comments/i28r3o/should_i_use_serverless_or_amplify_cloudfront_to/
---
I'm new to AWS so it's a bit overwhelming with all the options and info.

The website I want to build is relatively basic and standard. The website being a SPA and including a contact form might just be the most advanced features.

If I want to build a REST backend with Django Rest Framework (or perhaps I'll learn to use FastAPI instead), and a frontend with Vue; is it better to use serverless or Amplify for the frontend, and CloudFront for the REST backend?

Can you explain the pros and cons? What about price?
## [2][Using AWS Code* to build and deploy is neat but so much manual work](https://www.reddit.com/r/aws/comments/i1xl4l/using_aws_code_to_build_and_deploy_is_neat_but_so/)
- url: https://www.reddit.com/r/aws/comments/i1xl4l/using_aws_code_to_build_and_deploy_is_neat_but_so/
---
I went to several services like CodeBuild, Pipeline, Beanstalk in AWS console to generate a project name, a build, to pipeline, etc. There were so many manual steps I had to do to deploy a basic web app. How can I automate this so that it can be part of IaC? I used CodeCommit and I'd like the pipeline to be converted into IaC and commit that as part of my web project in CodeCommit.
## [3][Lambda/Api Gateway](https://www.reddit.com/r/aws/comments/i2bbyk/lambdaapi_gateway/)
- url: https://www.reddit.com/r/aws/comments/i2bbyk/lambdaapi_gateway/
---
Right now I have a test lambda handler which I'm integrating with an Api Gateway proxy endpoint:

    
    
    export async function func1(event: any) {
      console.log('next line is the full request event object, 2 depth');
      console.log(JSON.stringify(event, undefined, 2));
    
      return {
        statusCode: 200,
        headers: { "content-type" : "text/plain" },
        body: `From the return body. This route is: ${event.path}`
      };
    }

I'm looking at the docs but don't see how to access query or route parameters: [https://docs.aws.amazon.com/cdk/api/latest/docs/aws-lambda-readme.html](https://docs.aws.amazon.com/cdk/api/latest/docs/aws-lambda-readme.html)

I assume those are somewhere on the `event` object. Does anyone know where?
## [4][Does Elastic Beanstalk support tensorflow with python 3.8 ?](https://www.reddit.com/r/aws/comments/i27dwp/does_elastic_beanstalk_support_tensorflow_with/)
- url: https://www.reddit.com/r/aws/comments/i27dwp/does_elastic_beanstalk_support_tensorflow_with/
---
Hi everyone!

I have never used aws before and I have a django project I want to deploy. I am using Tensorflow (version 2.3) and python 3.8.2. I am also using tensorflow hub and not using any databases. I was wondering if AWS elastic beanstalk is compatible with tensorflow 2.3 and python 3.8?

Any insight would be greatly appreciated
## [5][Deploying serverless Django applications to AWS with CDK on a tiny budget using Lambda, API Gateway, awsgi and the Lambda proxy pattern](https://www.reddit.com/r/aws/comments/i1u5hv/deploying_serverless_django_applications_to_aws/)
- url: https://briancaffey.github.io/2020/08/01/django-and-lambda-with-cdk-and-api-gateway.html
---

## [6][AWS-SQS-how delayed retry works?](https://www.reddit.com/r/aws/comments/i26qg2/awssqshow_delayed_retry_works/)
- url: https://www.reddit.com/r/aws/comments/i26qg2/awssqshow_delayed_retry_works/
---
Hi guys,

I want the retry of message to be delayed like one minute so that everytime the error occurs, it will not trigger the same message immediately.

&amp;#x200B;

In order to do this, can I just set the visibilityTimeout to be 1 min? And what is the mechanism under the hood? 

&amp;#x200B;

Because for retrying the message is still in the queue, so how we polling the message depends on two things (1) message exist. 2) visibityTimeout is reached) ? And everytime the same message gets polled, the visibilityTimeout wii be reset?

&amp;#x200B;

Is that understanding correct?

&amp;#x200B;

Thanks!
## [7][Is a shared EKS control plane possible in a multi-account AWS environment?](https://www.reddit.com/r/aws/comments/i1yq1f/is_a_shared_eks_control_plane_possible_in_a/)
- url: https://www.reddit.com/r/aws/comments/i1yq1f/is_a_shared_eks_control_plane_possible_in_a/
---
In an AWS environment where there are shared VPC's for Dev,Test,Prod,Shared and each project team has accounts broken out by resources such as subnets, do you see issues running an EKS control plane either in each VPC or in the Shared VPC?

The goal being that each team has its own compute nodes for isolation and billing, while trying to simplify the EKS management.
## [8][CodePipeline: when to use stages vs actions?](https://www.reddit.com/r/aws/comments/i2586d/codepipeline_when_to_use_stages_vs_actions/)
- url: https://www.reddit.com/r/aws/comments/i2586d/codepipeline_when_to_use_stages_vs_actions/
---
A pipeline contains stages, stages contain actions, actions act upon artifacts.

The [AWS CodePipeline concepts docs](https://docs.aws.amazon.com/codepipeline/latest/userguide/concepts.html) say:

&gt;A **stage** might be a build stage, where the source code is built and tests are run. It can also be a deployment stage, where code is deployed to runtime environments

An example of this is in the [Terraform CodePipeline example](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/codepipeline#example-usage), which has a "Source", "Build", and "Deploy" stage. The [AWS four stage pipeline tutorial](https://docs.aws.amazon.com/codepipeline/latest/userguide/tutorials-four-stage-pipeline.html#tutorials-four-stage-pipeline-pipeline-create) does something similar.

The AWS CodePipeline concepts docs also say:

&gt;\[An **action**\] can include things like a source action from a code change, an action for deploying the application to instances, and so on... Valid CodePipeline action types are source, build, test, deploy, approval, and invoke

The overlap between the two levels confuses me. Why structure stages in such a manner if actions do those things? In doing so, a "Source" stage could hypothetically have a "build", "test", and "deploy" actions, which doesn't make sense to me. Wouldn't it stand to reason that each stage should correspond to an environment (dev, QA, staging, prod, etc), where each action takes care of a step toward releasing to the environment?
## [9][Help downloading large files from an SFTP server to an EC2 instance](https://www.reddit.com/r/aws/comments/i1zzf3/help_downloading_large_files_from_an_sftp_server/)
- url: https://www.reddit.com/r/aws/comments/i1zzf3/help_downloading_large_files_from_an_sftp_server/
---
Hi all-

Tl:dr at the bottom! 

I’m a newbie to all of this (AWS and computing in general) and would really appreciate some advice. 

My task is to download a set of files from an SFTP server and then upload them to an S3 bucket. Conceptually, this isn’t difficult; for smaller files, I wrote a Lambda function (w/Python 3.8 runtime) that works just fine.

However, I’ve been having trouble doing this with the larger files on the SFTP server (ranging from 1-150 GB). For these files, I set up an EC2 instance with sufficient memory to store the largest file and then ran a Python script (similar to the handler function that I use to download/upload smaller files via the Lambda function) that uses the Paramiko module to download the large files and then the boto3 module to subsequently upload them to S3. This fails each time; paramiko throws an EOFError that seems to be a known issue with the module, but which nobody quite knows how to fix. I’ve tried various pseudo-fixes and monkey patches, but nothing has worked. In short, paramiko doesn’t seem to be all that great for handling large files. 

My question, then, is this: can any of you recommend a Python module, AWS service, or scripting language that will allow me to reliably download large files from an SFTP server? It’s important to note that I do not control this server, and I can only access it via username/password (I don’t have SSH keys). Thank you!

TL;DR I’m having trouble downloading large files from an SFTP server, and would greatly appreciate any recommendations on Python modules, AWS services, or other scripting languages that I can use to reliably download large files. Thank you!
## [10][EC2 Auto Scale using a variable predefined schedule](https://www.reddit.com/r/aws/comments/i1rg0y/ec2_auto_scale_using_a_variable_predefined/)
- url: https://www.reddit.com/r/aws/comments/i1rg0y/ec2_auto_scale_using_a_variable_predefined/
---
I want to conduct online examinations on EC2 instances.   
I know how many students there will be at each time instant as the tests are pre-scheduled.  
So I want to scale using the schedule I have, it is variable not fixed or weekly, Can anyone help me where to look for this.

Thank you
