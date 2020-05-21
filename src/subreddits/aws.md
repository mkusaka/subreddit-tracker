# aws
## [1][Is there a way to block a Lambda function to be run no more than once per second for a any given IP address?](https://www.reddit.com/r/aws/comments/gnpwsw/is_there_a_way_to_block_a_lambda_function_to_be/)
- url: https://www.reddit.com/r/aws/comments/gnpwsw/is_there_a_way_to_block_a_lambda_function_to_be/
---
Is there a way to block a Lambda function to be run no more than once per second for a any given IP address?
## [2][Has anyone here made the transition from AWS enthusiast to AWS employee?](https://www.reddit.com/r/aws/comments/gnezee/has_anyone_here_made_the_transition_from_aws/)
- url: https://www.reddit.com/r/aws/comments/gnezee/has_anyone_here_made_the_transition_from_aws/
---
Being on the 'customer' side of a 'customer-obsessed' company is one thing, but I would be curious if anyone has made the transition to working at AWS and how that went.
## [3][AWS logging &amp; Monitoring Strategy Guide (w/ CloudFormation and Terraform templates)](https://www.reddit.com/r/aws/comments/gndkrg/aws_logging_monitoring_strategy_guide_w/)
- url: https://asecure.cloud/g/strategy_logging/
---

## [4][Azure blob storage to AWS S3 data migration](https://www.reddit.com/r/aws/comments/gnpxsb/azure_blob_storage_to_aws_s3_data_migration/)
- url: https://www.reddit.com/r/aws/comments/gnpxsb/azure_blob_storage_to_aws_s3_data_migration/
---
Hi everyone, 

I am planning on migrating 2 years worth of data that has approximately 45000 blobs/day. What would be a good and efficient way to do this?
I am currently running a python script in an ecs fargate container - the script lists and downloads the blob, then uploads it to s3. It is an extremely time consuming procedure and it is not fail proof either. 

Please let me know if there are any aws services that I can use to make this better
## [5][CloudFormation now supports Blue/Green deployments for ECS!](https://www.reddit.com/r/aws/comments/gni34t/cloudformation_now_supports_bluegreen_deployments/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/05/aws-cloudformation-now-supports-blue-green-deployments-for-amazon-ecs/
---

## [6][Preventing cross account roles external to Organization?](https://www.reddit.com/r/aws/comments/gniv3j/preventing_cross_account_roles_external_to/)
- url: https://www.reddit.com/r/aws/comments/gniv3j/preventing_cross_account_roles_external_to/
---
Hi All,

Currently our developers have limited access in IAM, they can create roles provided they follow a naming convention and also attach a boundary policy to the role. But how can I prevent them from creating a role that trusts a role from an account outside our organization?

It's totally feasible that they would do a cross account role for accounts within our organization for collaboration work, but not for outside accounts.

I tried setting a deny on iam:CreateRole with the following condition, but it does not work:

            "Condition": {
                "StringNotEquals": {
                    "aws:PrincipalOrgID": [
                        "o-blahblah"
                    ]
                }
            },

Is this the incorrect approach? What have you guys done to secure your organization from this issue?
## [7][Unity -&gt; .net sdk for using pinpoint](https://www.reddit.com/r/aws/comments/gnsjld/unity_net_sdk_for_using_pinpoint/)
- url: https://www.reddit.com/r/aws/comments/gnsjld/unity_net_sdk_for_using_pinpoint/
---
Hi there!  


Relative noob question.   
On the game-day chat yesterday, I was advised that unity can post analytics to pinpoint using the .net SDK.   


&gt;For Unity, you can use the AWS SDK for .NET. Download the Pinpoint package from Nuget.  
&gt;  
&gt;  
&gt;  
&gt;[https://www.nuget.org/packages/AWSSDK.Pinpoint/3.3.109.16](https://www.nuget.org/packages/AWSSDK.Pinpoint/3.3.109.16)  
&gt;  
&gt;  
&gt;  
&gt;Add the Pinpoint assembly and its dependencies to the Assets\\Plugins in your Unity Project. Then you should be able to reference that in your C# Code. Don’t forget to use  
&gt;  
&gt;  
&gt;  
&gt;using Amazon.Pinpoint  
&gt;  
&gt;  
&gt;  
&gt;at the top of the file to enable the code to use the Pinpoint features. There may be other namespaces in the Pinpoint SDK assembly.  
&gt;  
&gt;  
&gt;  
&gt;The Unity project ‘API Compatibility’ setting in the project settings must be set to ‘.NET 4.x’

  
Could anyone advise me as to ***how*** to add the right assemblies / dependencies to my project?   


Alternatively, if anyone has the depreciated mobile-analytics sdk working, could they confirm that events are piped into pinpoint?   


Many thanks!
## [8][DynamoDB now supports empty values for non-key String and Binary attributes](https://www.reddit.com/r/aws/comments/gn5hhi/dynamodb_now_supports_empty_values_for_nonkey/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/05/amazon-dynamodb-now-supports-empty-values-for-non-key-string-and-binary-attributes-in-dynamodb-tables/
---

## [9][Is there anyway to achieve unlimited domains/subdomains pointing to ALB?](https://www.reddit.com/r/aws/comments/gnps7t/is_there_anyway_to_achieve_unlimited/)
- url: https://www.reddit.com/r/aws/comments/gnps7t/is_there_anyway_to_achieve_unlimited/
---
So I just ran into a brick wall with the stupid 25 SSL cert limit on ALB.  I need to point UNLIMITED domains/subdomains to our ALB and our proxy app on the instance will take care of doing stuff behind the scenes.  Is there anyway to achieve this without creating a bunch of ALBs?
## [10][IAM time bound session termination](https://www.reddit.com/r/aws/comments/gnpkv1/iam_time_bound_session_termination/)
- url: https://www.reddit.com/r/aws/comments/gnpkv1/iam_time_bound_session_termination/
---
Whenever the support engineers are required to perform any level of elevated actions within an account they request a time bound access and then they are added to an AD security group that ties back to the support role. 

Now let's say the time bound access is set to 3 hrs and after 3 hrs sailpoint IAG automatically removes the users from the AD groups what we have noticed here is the  users can still access the AWS accounts after 3hrs via the CLI even after they are removed from the AD SG until they are asked to revalidate themselves again i.e. log out and log back into the CLI. 

There must be a way of session synchronization for a logged in user. This cannot be done at the role as multiple users from that role might have session access at the same time. 

Please let know if there's a better way to revoke time bound access via the CLI as soon as the user is removed from the AD security group.
