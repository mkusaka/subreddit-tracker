# aws
## [1][On-demand CI/CD infrastructure with GitLab and AWS Fargate - How to reduce costs and scale GitLab Runner down to zero](https://www.reddit.com/r/aws/comments/hwe4hx/ondemand_cicd_infrastructure_with_gitlab_and_aws/)
- url: https://www.reddit.com/r/aws/comments/hwe4hx/ondemand_cicd_infrastructure_with_gitlab_and_aws/
---
In his new article, Daniel Miranda shows how we can use **AWS Lambda** functions to stop the Runner manager hosted on **AWS Fargate** when there are no CI/CD jobs to process and start it when a new pipeline is triggered. This configuration can significantly reduce the costs when we have considerable idle times between builds.

[https://medium.com/ci-t/on-demand-ci-cd-infrastructure-with-gitlab-and-aws-fargate-376edc7afcda](https://medium.com/ci-t/on-demand-ci-cd-infrastructure-with-gitlab-and-aws-fargate-376edc7afcda)
## [2][IAM role for automatic AWS CDK deployment](https://www.reddit.com/r/aws/comments/hwc4jf/iam_role_for_automatic_aws_cdk_deployment/)
- url: https://www.reddit.com/r/aws/comments/hwc4jf/iam_role_for_automatic_aws_cdk_deployment/
---
Hi there! Is there a pre-defined policy or best-practices for creating an IAM role that the CI server uses to deploy via AWS CDK? I don't want to only list the services we are currently using, as this would mean manually logging into and updating this role whenever we need a new service, but I also don't think that the CI will need the ability to e. g. read from our database.

This is talking about production only, staging and local environments live in different AWS accounts.
## [3][Cross-account AWS Backup - possible or not?](https://www.reddit.com/r/aws/comments/hwcchr/crossaccount_aws_backup_possible_or_not/)
- url: https://www.reddit.com/r/aws/comments/hwcchr/crossaccount_aws_backup_possible_or_not/
---
Is it possible to set up AWS Backup to backup resources located on different account, or to store backups in a vault located on different account?
## [4][Cloud Batch](https://www.reddit.com/r/aws/comments/hwe7wa/cloud_batch/)
- url: https://www.reddit.com/r/aws/comments/hwe7wa/cloud_batch/
---
I was trying to come up with a technique to run the Offline processes on AWS. I tempted to use Lambda processes, but it appears to be lasting for 15 minutes only. Some of my jobs read huge files with 100,000 records for which processing takes hours.. Lambda will not support these long running processes. The other solution is to run an dedicated ec2 instance and spin it for the Job, and use SQS /Cloud watch to push the events.. this should work.. however i will always have to keep this containers up and running despite my jobs run for no more than 1-2 hours on different time of the day. Wanted to know your views?
## [5][CloudFormation - IAM policy already exists](https://www.reddit.com/r/aws/comments/hwfc31/cloudformation_iam_policy_already_exists/)
- url: https://www.reddit.com/r/aws/comments/hwfc31/cloudformation_iam_policy_already_exists/
---
Hi all,

I am running in to a situation where when updating a CFn stack it fails on an existing IAM policy.

&gt;dev-emr-policy already exists

I have several such policies in that stack but its complaining only for this one. To give you some background, this is a nested template setup where we have a main.yml which calls this stack and other stacks(TemplateURL) with the required parameters, this stack has several IAM resources we are creating/controlling through CFn. Just like this policy there are other pre existing IAM policies but its not complaining about those. This is happening when I modified some other stack which is part of the nested stack setup. There were no errors on the modified stack but this IAM stack is errorring out.  When I comment out this policy update works without any errors.

I am stuck here and could use some help please.

Thanks

&amp;#x200B;

    emrPolicy:
      Type: AWS::IAM::ManagedPolicy
      Properties:
        ManagedPolicyName: !Sub ${EnvironmentName}-emr-policy
        PolicyDocument:
          Version: "2012-10-17"
          Statement:
            -
              Effect: "Allow"
              Action:
                - "XXX"
                - "YYY"
                - "ZZZ"
              Resource:
                - !Sub arn:aws:AAA:${AWS::Region}:${AWS::AccountId}:something:/${EnvironmentName}/blah/ABC-??????

&amp;#x200B;
## [6][Lambda function is not invoking the attached destination, no matter how i invoke the function](https://www.reddit.com/r/aws/comments/hwesif/lambda_function_is_not_invoking_the_attached/)
- url: https://www.reddit.com/r/aws/comments/hwesif/lambda_function_is_not_invoking_the_attached/
---
Hi,

So i have the setup as following,

* A lambda function whose role has full access to EC2, SNS, SQS, Lambda, etc.
* A SNS to invoke the function
* A SNS topic with EMAIL subscription to send emails to a particular email
* Lambda function can execute without any issue or throw an error if something goes wrong.
* A destination attached which send Async success or failure updates to the SNS topic with EMAIL subscription

Now, I tried invoking this lambda using using the attached SNS topic, or using cli with invoke --invocation-type EVENT  
 or using invoke-async  
 but nothing is triggering the destination it seems as i'm not getting any email.

Can anyone please suggest or indicate what could be wrong or if i'm doing anything wrong?

Please let me know if you need any other information regarding this.
## [7][Is CodeCommit Really a Thing?](https://www.reddit.com/r/aws/comments/hvyd6h/is_codecommit_really_a_thing/)
- url: https://www.reddit.com/r/aws/comments/hvyd6h/is_codecommit_really_a_thing/
---
Just wondering if anyone out there actually uses it (and if so, why).

Seems kind of... forced.
## [8][AWS MediaLive and MediaStore](https://www.reddit.com/r/aws/comments/hw8z27/aws_medialive_and_mediastore/)
- url: https://www.reddit.com/r/aws/comments/hw8z27/aws_medialive_and_mediastore/
---
Hey everyone,
Has anyone implemented an encoder for IPTV using MediaStore and MediaLive??
We currently have a physical encoder receiving multicast traffic on one of its interfaces, after processing the traffic its configured to send the traffic to another server.
 The AWS MediaLive doesn't seem to be able to receive multicast traffic at all, at least I couldn't see a way.
Has anyone implemented an encoder in AWS able to receive that multicast traffic??? Is MediaLive and MediaStore the way to do that???? Im.not sure MediaLive is able.to received bare multicast traffic.
Any insights??
## [9][API Gateway: Include stage in custom URL](https://www.reddit.com/r/aws/comments/hwbete/api_gateway_include_stage_in_custom_url/)
- url: https://www.reddit.com/r/aws/comments/hwbete/api_gateway_include_stage_in_custom_url/
---
Let me start with that I am a pretty novice user of AWS and don't use CLI or anything for it. I am using AWS for a little Slack tool I built for me and others. I am familiar with CLI tools just don't use it for AWS yet.

I am using API Gateway to communicate to a set of Lambda functions for the Slack tool. It uses a custom domain that used to have the following URL pattern [api.example.com/v1/{stage}/{resource}](https://api.example.com/v1/{stage}/{resource}). I needed to reset the config because of an issue with certificates. I did that through the custom domain UI in AWS console and I can only get it to serve the API with a pattern of [api.example.com/v1/{resource}](https://api.example.com/v1/{resource}) 

This is a huge concern since some of the URLs are used by clients I don't control. Is there any way to replicate the old behaviour? Thanks for any help!
## [10][How to get ReadWriteMany persistent volume without losing its resilience?](https://www.reddit.com/r/aws/comments/hwfaz8/how_to_get_readwritemany_persistent_volume/)
- url: https://www.reddit.com/r/aws/comments/hwfaz8/how_to_get_readwritemany_persistent_volume/
---
To get a resilient ReadWriteMany persistent volume on AWS, you could use EFS provisioner. This tutorial shows you [how to use Kubernetes PVCs with EFS provisioner](https://www.padok.fr/en/blog/efs-provisioner-kubernetes) to get a more resilient solution.
