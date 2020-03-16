# aws
## [1][CloudFront Updates Are No Longer Soul Destroying - Adam Johnson](https://www.reddit.com/r/aws/comments/fjjdkg/cloudfront_updates_are_no_longer_soul_destroying/)
- url: https://adamj.eu/tech/2020/03/16/cloudfront-updates-no-longer-soul-destroying/
---

## [2][New AWS security tool for Security Hub - ElectricEye](https://www.reddit.com/r/aws/comments/fj8ij5/new_aws_security_tool_for_security_hub_electriceye/)
- url: https://www.reddit.com/r/aws/comments/fj8ij5/new_aws_security_tool_for_security_hub_electriceye/
---
Hey folks, wanted to share a new side project I made public today called ElectricEye. The TL;DR is it is yet another continuous compliance monitoring tool for AWS, but I built it with Security Hub in mind. All of the checks are Python3.8, run on a container on top of Fargate and also support some additional integrations with Config Recorder, Shodan, Slack, JIRA and ServiceNow.

Some high level stats:

\- 162 checks across 50 services / components (components being defined as a part of a large service, like RDS Snapshots as part of RDS, etc.). It supports things not supported by Security Hub or Config like DocDb, MSK, Cognito, EKS, etc.

\- 60 response &amp; remediation playbooks that support ElectricEye checks and Security Hub security standards, I was the author of the original CIS playbooks while I was at AWS and wanted to greatly improve and expand that

\- CloudFormation and Terraform support for everything, least privileged IAM policies set where they were able to be defined to include both the core module &amp; add-on's.

It is no where near complete or a mature project like Prowler, Scout2, PacBot, etc. but being a former member of the Security Hub team it is a service near and dear to me and was one of the motivating factors for publishing this.

All feedback is welcomed, go ahead and take a look here:  [https://github.com/jonrau1/ElectricEye](https://github.com/jonrau1/ElectricEye)
## [3][Any plans for expanding online delivery of exams during a pandemic](https://www.reddit.com/r/aws/comments/fjgi39/any_plans_for_expanding_online_delivery_of_exams/)
- url: https://www.reddit.com/r/aws/comments/fjgi39/any_plans_for_expanding_online_delivery_of_exams/
---
hey AWS support, are there any plans for including professional &amp; specialty level certifications to the online proctoring catalog? I know that CCP is being delivered online.

In a situation like this, where candidates like me want to continue AWS certification journey without stressing about COVID-19, do you have something in pipeline that we can look forward to?

I am reliant on public transport and the nearest test centre is roughly 50 miles away. I am sure there are more candidates like me.
## [4][Cognito: aws console users](https://www.reddit.com/r/aws/comments/fjk8cx/cognito_aws_console_users/)
- url: https://www.reddit.com/r/aws/comments/fjk8cx/cognito_aws_console_users/
---
Hi, 

I am developing a "meta" dashboard that enables our ops team to have an overview on a bunch of aws accounts that we manage.

I wanted to use cognito to authenticate to that dashboard, my idea was to use the existing users in our "management" aws console account to login into the dashboard using cognito.

Is this possible? How do I federate with the aws console? If is not possible what do you recomend?
## [5][6 Ways To Reduce Your AWS Bill - Powered by Volumetree](https://www.reddit.com/r/aws/comments/fji7cg/6_ways_to_reduce_your_aws_bill_powered_by/)
- url: https://www.volumetree.com/reduce-your-aws-bill/?utm_source=Reddit&amp;utm_medium=Pallvi
---

## [6][SageMaker: automatically stop your instances when idle](https://www.reddit.com/r/aws/comments/fj69au/sagemaker_automatically_stop_your_instances_when/)
- url: https://biasandvariance.com/sagemaker-stop-your-instances-when-idle/
---

## [7][Pre Token Generation](https://www.reddit.com/r/aws/comments/fjhrjm/pre_token_generation/)
- url: https://www.reddit.com/r/aws/comments/fjhrjm/pre_token_generation/
---
I've written a simple NodeJS lambda to add a new field to the JWT.When I decode the JWT I do not see the field. am I doing something blatantly wrong?

Thanks in advance

    const bearerFunction = (event, context, callback) =&gt; {   
      event.response = { 
        claimsOverrideDetails: { 
          claimsToAddOrOverride: { 
            "ThisIsANewClaim": "Testymctestest" 
          } 
        } 
      }; 
      callback(null, event); 
    }; 

edit - got code formatting
## [8][Increase AWS EC2's EBS Volume without Downtime](https://www.reddit.com/r/aws/comments/fj5pn7/increase_aws_ec2s_ebs_volume_without_downtime/)
- url: https://www.devopscycle.com/increase-aws-ec2-ebs-volume-without-downtime
---

## [9][Why use AWS Systems Manager Patch Manager?](https://www.reddit.com/r/aws/comments/fjdcy3/why_use_aws_systems_manager_patch_manager/)
- url: https://www.reddit.com/r/aws/comments/fjdcy3/why_use_aws_systems_manager_patch_manager/
---
I'm studying for a Pro exam and there are questions suggesting creating a patch baseline etc. 

Surely a better approach is to use a new OS AMI with the updates?

I know it's not an option, but still, I am wondering why AWS suggest this method at all. It feels error prone to me.
## [10][Rate the architecture. What would you like to improve?](https://www.reddit.com/r/aws/comments/fjfw84/rate_the_architecture_what_would_you_like_to/)
- url: https://i.redd.it/z77q7hpy1zm41.png
---

