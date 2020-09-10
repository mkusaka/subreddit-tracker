# aws
## [1][API Gateway HTTP APIs authorization now supports Lambda and IAM](https://www.reddit.com/r/aws/comments/iptsk8/api_gateway_http_apis_authorization_now_supports/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/09/api-gateway-http-apis-now-supports-lambda-and-iam-authorization-options/
---

## [2][In praise of S3, the greatest cloud service of all time](https://www.reddit.com/r/aws/comments/iq20le/in_praise_of_s3_the_greatest_cloud_service_of_all/)
- url: https://acloudguru.com/blog/engineering/brazeal-in-praise-of-s3-the-greatest-cloud-service-of-all-time
---

## [3][Lamdba functions: does adding layers add cost?](https://www.reddit.com/r/aws/comments/iq33gf/lamdba_functions_does_adding_layers_add_cost/)
- url: https://www.reddit.com/r/aws/comments/iq33gf/lamdba_functions_does_adding_layers_add_cost/
---
Hi all,

I'm running some lambdas and they import a few layers. When the instance boots up, is this time that you get charged for, or is it only the 'computation' time later on that you get charged for?

What's a way to optimise this process, don't worry about being overly detailed, just point me at a link - I'd be very grateful.

Thanks in advance!
## [4][[OC] Boolean parameters in CloudFormation](https://www.reddit.com/r/aws/comments/iq30k6/oc_boolean_parameters_in_cloudformation/)
- url: https://awholenother.com/2020/06/20/boolean-parameters-in-cloudformation.html
---

## [5][Create a Cost Intelligence Dashboard (AWS Well-Architected Lab)](https://www.reddit.com/r/aws/comments/iphlhv/create_a_cost_intelligence_dashboard_aws/)
- url: https://wellarchitectedlabs.com/cost/200_labs/200_enterprise_dashboards/
---

## [6][AWS Lambda and PyTorch](https://www.reddit.com/r/aws/comments/ipzy43/aws_lambda_and_pytorch/)
- url: https://www.reddit.com/r/aws/comments/ipzy43/aws_lambda_and_pytorch/
---
I wish to deploy a python flask application on aws lambda, but it has PyTorch as a dependency. And thus I am unable to do so, due to the [AWS Lambda limits](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-limits.html). 

I have found a solution to this, explained in this [post](https://aws.amazon.com/blogs/aws/new-a-shared-file-system-for-your-lambda-functions/) which uses shared files system for the lambda function. It would be great if someone could explain to me the pros and cons of this approach.

And how is it better than this approach (taken from [here](https://github.com/mattmcclean/sam-pytorch-example))

&gt;AWS Lambda has a limit of 250 MB for the deployment package size including lamba layers. PyTorch plus its dependencies is more than this so we need to implement a trick to get around this limit. We will create a zipfile called \`.requirements.zip\` with all the PyTorch and associated packages. We will then add this zipfile to the Lambda Layer zipfile along with a python script called \`unzip\_requirements.py\`. The python script will extract the zipfile \`.requirements.zip\` to the /tmp when the Lambda execution context is created.

And are there any other ways to do the same?
## [7][How to store private patient records in an S3 bucket?](https://www.reddit.com/r/aws/comments/iq2hdi/how_to_store_private_patient_records_in_an_s3/)
- url: https://www.reddit.com/r/aws/comments/iq2hdi/how_to_store_private_patient_records_in_an_s3/
---
Say I have bucket of hospital records s3://hospital and I want to store patient records by prefix, s3://hospital/patientZero/covid_results.txt.

I know I could use ${aws:username} as described https://aws.amazon.com/premiumsupport/knowledge-center/iam-s3-user-specific-folder/ to limit access to that prefix to ${aws:username}, assuming that username is "patientZero".

But I am aware of HIPAA where you're not allowed to have user identifiable names (metadate) in the paths like "patientZero". So is there a way to **hash** it in the policy? Or is there some other approach I'm missing?

The idea is only the assumed role/user can access his/her prefix/path of object data. If there is good guidelines / best practices for private file hosting in a bucket, do please let me know! Perhaps I should ask this question on SO?
## [8][What is the the fastest instance/networking/way to transfer S3 bucket from one account to another?](https://www.reddit.com/r/aws/comments/iq2a1i/what_is_the_the_fastest_instancenetworkingway_to/)
- url: https://www.reddit.com/r/aws/comments/iq2a1i/what_is_the_the_fastest_instancenetworkingway_to/
---
I am Copying an S3 bucket from one AWS account to another. I did some research and am now using an instance using m5n.2xlarge with enhanced networking enabled for the transfer machine. I have seen speeds between 5 and 10 MiB/s.

What would you recommend to use or do to get a faster transfer of the S3 bucket?
## [9][Consume data from external Kafka](https://www.reddit.com/r/aws/comments/iq24ts/consume_data_from_external_kafka/)
- url: https://www.reddit.com/r/aws/comments/iq24ts/consume_data_from_external_kafka/
---
I’m curious if there is a best practice for consuming data from an Kafka (not AWS managed, but within the AWS IaaS environment) into an AWS account via VPC peering.

Some options I’ve explored are:

- Kafka to S3
- Kafka to CloudWatch Logs
- Kafka to Vector.io (Fargate) and then to CloudWatch or S3
## [10][ffprobe - aws sagemaker](https://www.reddit.com/r/aws/comments/iq1b7w/ffprobe_aws_sagemaker/)
- url: https://www.reddit.com/r/aws/comments/iq1b7w/ffprobe_aws_sagemaker/
---
Hi everyone ,

I'm a new user of AWS. I want to run my machine learning model in this platform.  

When I try to import that library  `from pydub import AudioSegment` , 

I am taking that error: `FileNotFoundError: [Errno 2] No such file or directory: 'ffprobe': 'ffprobe'` 

I tried to use that commands -&gt; `!sudo yum install AudioSegment` and `!pip install ffprobe` and `!sudo yum install -y pydub` and other combinations but there is still same error.  

Some one help to me? Thanx.
