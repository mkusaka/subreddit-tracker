# aws
## [1][re:Invent registration is now open](https://www.reddit.com/r/aws/comments/jkenu3/reinvent_registration_is_now_open/)
- url: https://register.virtual.awsevents.com/
---

## [2][Week of Nov 2nd - What are you building this week in AWS?](https://www.reddit.com/r/aws/comments/jmqb02/week_of_nov_2nd_what_are_you_building_this_week/)
- url: https://www.reddit.com/r/aws/comments/jmqb02/week_of_nov_2nd_what_are_you_building_this_week/
---
Share what you are working on
## [3][AWS preps its own library of public Docker container images](https://www.reddit.com/r/aws/comments/jnimsn/aws_preps_its_own_library_of_public_docker/)
- url: https://www.zdnet.com/article/aws-preps-its-own-library-of-public-docker-container-images/
---

## [4][Why does Artifact require an NDA? Who can I share the files with?](https://www.reddit.com/r/aws/comments/jnvg3l/why_does_artifact_require_an_nda_who_can_i_share/)
- url: https://www.reddit.com/r/aws/comments/jnvg3l/why_does_artifact_require_an_nda_who_can_i_share/
---
I'm not clear as to why Artifact requires an NDA if the intent is to share the files with third parties (e.g. auditors). Who, exactly, may I share these files with? Can I give them to my insurance company, my HOA, my bookie? Can I email them, or must I grant IAM access to my auditors?
## [5][KMS Policy help](https://www.reddit.com/r/aws/comments/jntdng/kms_policy_help/)
- url: https://www.reddit.com/r/aws/comments/jntdng/kms_policy_help/
---
I have an S3 bucket which is used to store the organization trail, and the bucket is encrypted with KMS. 

The KMS policy is as follows:

    {
        "Version": "2012-10-17",
        "Id": "key-consolepolicy-3",
        "Statement": [
            {
                "Sid": "Enable IAM User Permissions",
                "Effect": "Allow",
                "Principal": {
                    "AWS": "arn:aws:iam::MyAccountNumber:root"
                },
                "Action": "kms:*",
                "Resource": "*"
            },
            {
                "Sid": "Allow CloudTrail to encrypt logs",
                "Effect": "Allow",
                "Principal": {
                    "Service": "cloudtrail.amazonaws.com"
                },
                "Action": "kms:GenerateDataKey*",
                "Resource": "*",
                "Condition": {
                    "StringLike": {
                        "kms:EncryptionContext:aws:cloudtrail:arn": "arn:aws:cloudtrail:*:*:trail/*"
                    }
                }
            },
            {
                "Sid": "Enable cross account log decryption",
                "Effect": "Allow",
                "Principal": {
                    "AWS": "*"
                },
                "Action": [
                    "kms:Decrypt",
                    "kms:ReEncryptFrom"
                ],
                "Resource": "*",
                "Condition": {
                    "StringEquals": {
                        "aws:PrincipalOrgID": "MyOrgID"
                    },
                    "StringLike": {
                        "kms:EncryptionContext:aws:cloudtrail:arn": "arn:aws:cloudtrail:*:*:trail/*"
                    }
                }
            },
            {
                "Sid": "Allow CloudTrail to describe key",
                "Effect": "Allow",
                "Principal": {
                    "Service": "cloudtrail.amazonaws.com"
                },
                "Action": "kms:DescribeKey",
                "Resource": "*"
            },
            {
                "Sid": "Allow thirdparty read the data",
                "Effect": "Allow",
                "Principal": {
                    "AWS": "arn:aws:iam::THIRDPARTY-Account-ID:role/THIRDPARTY-Role"
                },
                "Action": "kms:Decrypt",
                "Resource": "*"
            }
        ]
    }

The trail is working fine, now I need to allow a third party account which is not part of the org to read this bucket. I added the statement with SID "Allow thirdparty read the data" in the key policy, as well as the bucket policy as below:

    {
                "Sid": "Thirdparty read access",
                "Effect": "Allow",
                "Principal": {
                    "AWS": "arn:aws:iam::THIRDPARTY-Account-ID:role/THIRDPARTY-Role"
                },
                "Action": [
                    "s3:GetObject",
                    "s3:ListBucket"
                ],
                "Resource": [
                    "arn:aws:s3:::mybucket",
                    "arn:aws:s3:::mybucket/*"
                ]
            }

The IAM role used by third party also has s3:\* and kms:\* permissions. However, when they try to download any file, they are getting access denied. 

I created a support ticket, and AWS is saying that :

"This KMS policy will allow the "kms:Decrypt" and  "kms:ReEncryptFrom" action on any resource only if the request is coming from an IAM entity (users or roles) belonging to an AWS account which is part of an AWS Organization whose ID  is:   "MyOrgID". 

 In this scenario, the ec2 instance " i-xxxxxx“ which is assuming the role "arn:aws:iam::THIRDPARTY-Account-ID:role/THIRDPARTY-Role" is in an AWS account \[Account ID: xxx\] which is not a part of the AWS Organization with ID:  "MyOrgID". IAM entities (users or roles) present in accounts outside of their organization do not have access to the KMS key to be able to decrypt the objects and thus you are getting access denied while downloading files from S3.  "

I agree that, however, there is no explicit deny , and I explicitly allow the third party account to decrypt, so I do not think that this statement is correct from AWS. I also do not see any access denied error in KMS. ( I do not have bucket access logs enabled ). 

Is the support right ? If not, what should I do to allow the third party to read the file ?
## [6][Any plans for Aurora MySQL 5.6 supporting TLS1.1/1.2?](https://www.reddit.com/r/aws/comments/jntegz/any_plans_for_aurora_mysql_56_supporting_tls1112/)
- url: https://www.reddit.com/r/aws/comments/jntegz/any_plans_for_aurora_mysql_56_supporting_tls1112/
---
What the title says, basically. 

RDS MySQL 5.6 engines have this support for minor server versions 5.6.46 and up. But even the latest Aurora5.6 (1.23.0) engine is still stuck at 5.6.10. Any news or pointers on that?
## [7][How do I successfully authorize to AWS EKS cluster which was created in AWS console?](https://www.reddit.com/r/aws/comments/jnudh5/how_do_i_successfully_authorize_to_aws_eks/)
- url: https://www.reddit.com/r/aws/comments/jnudh5/how_do_i_successfully_authorize_to_aws_eks/
---
I use my organization's AWS account which has OrganizationAccountAccessRole applied. I created EKS cluster using the AWS console - in that moment there were no Users under Amazon IAM.

Cluster was created successfully.

I created new user under IAM and attached AdministratorAccess policy to him, so I am sure he has all the permissions needed for managing clusters.

After creating the user I downloaded csv with credentials for CLI access.

Then I ran `aws configure` command and set the user's credentials.

After that I ran aws eks `--region us-east-2 update-kubeconfig --name mycluster`  
 and checked that my cluster was successfully added to kubeconfig by executing `kubectl config view --minify` \- all changes are there in config.

My current kubectl context is set properly for my EKS cluster, but running `kubectl`  
 commands such as `kubectl get pods` ends up with a following error: *You must be logged in to the server (Unauthorized).*

In AWS documentation I found some aws commands for the cluster owner that can be executed to let non-owners access the cluster, but I can not access my cluster in my Linux console as I did create the cluster in AWS console.

Please, how can I troubleshoot this error and connect to my cluster from Linux console?
## [8][Can I dump keys from an AWS request? (To debug IAM conditions)](https://www.reddit.com/r/aws/comments/jnu4x3/can_i_dump_keys_from_an_aws_request_to_debug_iam/)
- url: https://www.reddit.com/r/aws/comments/jnu4x3/can_i_dump_keys_from_an_aws_request_to_debug_iam/
---
Is there any way to see which keys (and their values) are being passed to IAM with a request?  


Other than reading and trying to make sense of the various AWS lists of keys, is there a way for someone outside of the IAM rules engine to see what was submitted to the rules engine?

&amp;#x200B;

(Background :  
I'm trying to create a trust relationship that needs to include a wildcard for the assuming Principal's role name (because that role is created by CloudFormation, controlled by eksctl).

I have set the principal to the account ID and a condition statement with the wildcard. But now anything on that cluster can use the role. Which is not ideal.

I think the answer is another condition on role-session-name. But I was wondering if there were any other keys I could use. The session name isn't really a "secret" and anyone can choose to set it when they know what it is. So I'd prefer something more immutable. Although flexible so that if someone rebuilds the cluster it doesn't break..!

Last time I spent days trying to figure the rules out and discovered it was a bug. Only some SSM actions support the global condition key "aws:sourceVpce". Erm, so it isn't really "global" then is it? I'd like to avoid wasting time again.)
## [9][Networking compliance in AWS accounts](https://www.reddit.com/r/aws/comments/jnsyzl/networking_compliance_in_aws_accounts/)
- url: https://www.reddit.com/r/aws/comments/jnsyzl/networking_compliance_in_aws_accounts/
---
Hi,

currently my organization is moving to AWS and has strict guidelines enforcing that all internet traffic must go through a third party proxy. 

The direct result of these guidelines is that we need to restrict all lambdas in the accounts so that they need to be in an isolated VPC, these lambdas would need to use the proxy credentials from application code to access internet.

This has the issue that IAC tools, such as SAM, CDK, Terraform, etc... They won't work properly as these tools rely heavily on lambdas to perform their job, these lambdas are created of course with the default vpc settings, eventually these settings depending on the tool used can be modified so that these let's call them 'worker lambdas' are deployed.

The only issue is that these lambdas won't work as they will try to perform certain actions in your account, in some cases accessing certain services in AWS to do something in the account. We could overcome the issue by adding vpc endpoints for the services in the vpc, but this seems overkilled just to comply with the fact that all traffic needs to be monitored.

Questions:

1. Is the requirement of the traffic being monitored centrally overkilled in a cloud environment?
2. Is there a solution to make all internet traffic flow to this 3rd party proxy from the VPC, in a transparent way for the resources in the VPC? provided the VPC is configured correctly.
3. Related to 1, since AWS logs every account API call, combining this with VPC flow logs?, we could be compliant by saving this traffic data for future audits, am I wrong?
4. What'd be the easiest way to meet these compliance rules? (I think 3)

&amp;#x200B;

Thanks in advance.
## [10][Receiving Cookie-Header in Lambda using AWS::Serverless::Function and AWS::Serverless::HttpApi?](https://www.reddit.com/r/aws/comments/jnsyui/receiving_cookieheader_in_lambda_using/)
- url: https://www.reddit.com/r/aws/comments/jnsyui/receiving_cookieheader_in_lambda_using/
---
Hi,

I have a lambda function behind an APIGateway (type HTTP API) and everything deployed by SAM/CloudFormation.

    Resources:
      API:
        Type: AWS::Serverless::HttpApi
        Properties:
          StageName: !Ref ENV
    
      SomeFunction:
        Type: AWS::Serverless::Function
        Properties:
          FunctionName: some-function
          Role: !Ref LambdaRoleArn
          CodeUri: dist/lambda/somefunc/
          Events:
            HttpApiEvent:
              Type: HttpApi
              Properties:
                ApiId: !Ref API
                Method: GET
                Path: /somefunc

Within that lambda (written in go) I can receive all headers sent by the client except `Cookie`, so the APIGateway seems to filter these.

I have found out that on `AWS::ApiGateway::Method` resources there is a property `RequestParameters` where I could set `method.request.header.Cookie: true`, but I see no way to make SAM set this. Maybe I am confusing things, but I guess the reason is, that [SAM does not support the `x-amazon-apigateway-integration.requestParameters Object` extension](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-specification-api-gateway-extensions.html).

So is it actually possible to receive the `Cookie` header in my lambda or do I have to implement a separate lambda and use it as a custom Authorizer? ([https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-httpapi-httpapiauth.html](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-httpapi-httpapiauth.html))
## [11][Noob questions relating AWS, sorry for any inconvenience caused](https://www.reddit.com/r/aws/comments/jngkda/noob_questions_relating_aws_sorry_for_any/)
- url: https://www.reddit.com/r/aws/comments/jngkda/noob_questions_relating_aws_sorry_for_any/
---
Hello folks, I am working on a project and used the streamlit framework for the same to display some analytics on time series data. I want to host it on AWS as I have got some free credits from my uni. But I can't seem to figure out which services to use for my requirements, the options are too vast and it's confusing too. I have some timeseries\[16000 rows\] that increases at an average of 100 rows per day(columns are constant at 8), while building the streamlit app, I mostly used ML models\[to predict the next 10-20 timestamps\] and it runs on that data chosen from the options. Now, coming to the AWS part,

1. Where do I save my data? I mean which database? From my reading I learnt that Amazon Timestream is promising but I couldn't find a proper code samples to transit my data from csv files to Timestream data base(appreciate any example for the same), or should I rather go with typical database(relational dbms)?
2. I want to add some exploratory data analytics options to my app which I built based out of streamlit, could I do it on some AWS service and integrate it into my app? or do I need to build that on streamlit itself?
3. I have heard a lot of Sagemaker in the market, is it of any use to me? Can it offset my need to use streamlit and deploy the app directly?
4. Please highlight any services I might be using based on the details, would save me loads of time  
PS: everywhere I mentioned an app, I basically meant a web based app, appreciate your time for reading, thank you very much.
## [12][EC2 Image Builder now supports AMI distribution across AWS accounts](https://www.reddit.com/r/aws/comments/jn4ucq/ec2_image_builder_now_supports_ami_distribution/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/10/ec2-image-builder-now-supports-ami-distribution-across-aws-accounts/
---

