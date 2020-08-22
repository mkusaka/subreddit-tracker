# aws
## [1][How to organize infrastructure responsibilities if we build Micro-services with AWS CDK](https://www.reddit.com/r/aws/comments/ieece7/how_to_organize_infrastructure_responsibilities/)
- url: https://www.reddit.com/r/aws/comments/ieece7/how_to_organize_infrastructure_responsibilities/
---
I really like AWS CDK very much because it allows us to organize/align our team purely into a developer first way, i.e., each repo (say \`billing-service\`, or \`checking-out-service\`) directly corresponds to a vertical function team -- the repo contains the node.js service code and also contains the infrastructure setup code (i.e, how to setup Beanstalk cluster).

&amp;#x200B;

However, it seems that we still need a horizontal repo (and team) to take care the shared things across the vertical function repos (and teams) -- for example, if \`billing-service\` and  \`checking-out-service\` shares the same VPC, or even share the same ECS cluster, then the repo that is in charge of the shared VPC and ECS cluster seemingly have to be an independent repo, say 'vpc-and-ecs'

&amp;#x200B;

My question here are the following two:

1. In the above example, if we have to have the third repo of \`vpc-and-ecs\`, how can the  \`billing-service\` and  \`checking-out-service\` repo knows the output of \`vpc-and-ecs\`, such as CIDR block or ECS cluster ID and etc.? (I guess hard coding is ok at the very beginning, but I feel it's very hard to maintain across the team)
2. If we need to update the shared infrastructure code (vpc-and-ecs), say we want to change the VPC CIDR or change the subnets, it probably will have inevitable effect towards the  \`billing-service\` and  \`checking-out-service\` repo, how can we manage the repos change dependency and cross team communication?

Anyone thought of how to work with CDK in a large cross team?
## [2][Is there any eBook / website with sample aws architecture diagrams? Something similar to the aws this is my architectur series?](https://www.reddit.com/r/aws/comments/iehi9b/is_there_any_ebook_website_with_sample_aws/)
- url: https://www.reddit.com/r/aws/comments/iehi9b/is_there_any_ebook_website_with_sample_aws/
---

## [3][Good blog about SNS-SQS Fanout using message filtering + CloudFormation](https://www.reddit.com/r/aws/comments/ie65d5/good_blog_about_snssqs_fanout_using_message/)
- url: https://www.reddit.com/r/aws/comments/ie65d5/good_blog_about_snssqs_fanout_using_message/
---
Just saw this recently published blog about SNS-SQS filtering and using message attributes to filter. Especially the CloudFormation part was very useful for me.  


[https://medium.com/better-programming/how-to-fan-out-to-different-sqs-queues-using-sns-message-filtering-84cd23ed9d07](https://medium.com/better-programming/how-to-fan-out-to-different-sqs-queues-using-sns-message-filtering-84cd23ed9d07)
## [4][Auto provision IAM roles and permissions with AWS Control Tower](https://www.reddit.com/r/aws/comments/ie7ahp/auto_provision_iam_roles_and_permissions_with_aws/)
- url: https://www.reddit.com/r/aws/comments/ie7ahp/auto_provision_iam_roles_and_permissions_with_aws/
---
We've been using AWS Control Tower for about 6 months now to help us quickly provision accounts for various business units and its worked well for the most part. However, I have yet to find any straightforward approaches for adding customizations to the accounts during the creation process - specifically when trying to automate the addition of custom IAM roles upon account creation.

This is frustrating because we manage our AWS environments with Terraform, which is a great tool, but in order to start using Terraform with an AWS account we have to manually go into each account and create the IAM role Terraform will assume when it's interacting with the account.

Looking for a recommended process for automating account customizations with AWS Control Tower - specifically the ability to have a custom IAM role added to the account upon creation.

Any advice or links to documentation is appreciated.

Thanks,
## [5][Any ETA on Aurora PostgreSQL 12?](https://www.reddit.com/r/aws/comments/ie67ho/any_eta_on_aurora_postgresql_12/)
- url: https://www.reddit.com/r/aws/comments/ie67ho/any_eta_on_aurora_postgresql_12/
---
Just like for 11, the wait is excruciating...
## [6][Autoscalling IIS webser (AWS)](https://www.reddit.com/r/aws/comments/ieg1pf/autoscalling_iis_webser_aws/)
- url: https://www.reddit.com/r/aws/comments/ieg1pf/autoscalling_iis_webser_aws/
---
hello,

I create a basic webpage in IIS webserver and link it with private IP, after that created an AMI and link it with Autoscaling.

&amp;#x200B;

https://preview.redd.it/rfdwr2ml4ji51.png?width=527&amp;format=png&amp;auto=webp&amp;s=2a723350a91a46531cd31eacf396fa2cd2805438

Autoscaling creating specified instance but with different private IPs, so how would I access my IIS webpage with different IP, to access my webpages i have to access those machines and change their IP address. 

if I need 100 system, I can't log in and changed their assigned IP address for 100, is there any easy way to do it?
## [7][Secure API Gateway for React in Amplify](https://www.reddit.com/r/aws/comments/iedgkx/secure_api_gateway_for_react_in_amplify/)
- url: https://www.reddit.com/r/aws/comments/iedgkx/secure_api_gateway_for_react_in_amplify/
---
I'm creating a react application in Amplify. The react will query API Gateway. Eventually, this react application will be hosted in S3 + CloudFront for the public to use. 

I was wondering if it's secure if I create an API key for API Gateway and I use this API key to authenticate my application when calling the API. My concern is that react is essentially a front-end and my API key will be exposed. 

Is this an appropriate way to authenticate or is it insecure if so, what better way can I use to authenticate the react to the API?  

PS: Cognito is unable to do the job as my website can be used by anyone without any form of authentication.
## [8][Speed up data sync from S3 to ec2](https://www.reddit.com/r/aws/comments/idwio0/speed_up_data_sync_from_s3_to_ec2/)
- url: https://www.reddit.com/r/aws/comments/idwio0/speed_up_data_sync_from_s3_to_ec2/
---
Im looking for advice, I have a compute job that runs on an EC2 once a month.  I've optimized the job so that it runs within an hour, however the biggest bottleneck to date is syncing thousands of csv files to the machine before the job starts.  

If it helps the files are collected every minute from hundreds of weather stations, what are the options?
## [9][Upcoming AWS Livestream on Twitch: AWS Resource Tagging. Monday, August 24 – 11:00 am PST – https://www.twitch.tv/aws](https://www.reddit.com/r/aws/comments/ie3gdi/upcoming_aws_livestream_on_twitch_aws_resource/)
- url: https://i.redd.it/rsh4vj50qei51.jpg
---

## [10][AWS Amplify + Redux Saga: Proper way to create an Auth Listener? Event Channel + Hub?](https://www.reddit.com/r/aws/comments/ieatuj/aws_amplify_redux_saga_proper_way_to_create_an/)
- url: https://www.reddit.com/r/aws/comments/ieatuj/aws_amplify_redux_saga_proper_way_to_create_an/
---
**What I'm Implementing:**

I'm trying to implement an auth listener for my Redux Saga + AWS Amplify application. I'm using a [Redux Saga Event Channel](https://redux-saga.js.org/docs/advanced/Channels.html) and [AWS Amplify Hub (Auth listener)](https://docs.amplify.aws/lib/utilities/hub/q/platform/js)

&amp;#x200B;

**My Issues:**

1.  I'm having an issue where an event is being emitted on \`configured\` and \`signOut\`, but it won't emit an event for \`signIn\`.
2.  My second issue is with when I refresh the page. I'm using Redux Persist to persist the state (Cognito User object) when refreshed, but how can I verify that the user is currently logged in. Is there an event for that?

&amp;#x200B;

**Root Saga:**

    // Root Saga
    export default function* rootSaga() {
      // AWS Auth Channel
      yield fork(awsAuthChannelSaga);
    
      yield all([
        // Sagas: Auth
        fork(watchLoginSaga),
        fork(watchLogoutSaga),
        fork(watchResetPasswordSaga),
        fork(watchSendVerificationCodeSaga),
      ]);
    };

&amp;#x200B;

**Saga:**

    // Redux Saga: AWS Auth Channel
    export function* awsAuthChannelSaga() {
      try {
        // Channel
        const channel = eventChannel((emit: any) =&gt; {
          // Emitter
          const emitter = (event: any) =&gt; emit(event);
    
          // AWS: Auth Listener
          Hub.listen('auth', emitter);
    
          // AWS: Remove Auth Listener
          return () =&gt; Hub.remove('auth', emitter);
        });
    
    
        yield takeEvery(channel, function*(event: any) {
    
          // Event: Configured
          if (event.payload.event === 'configured') {
            console.log('AWS: Auth configured');
          }
          // Event: Sign In (NOT WORKING)
          else if (event.payload.event === 'signIn') {
            // Redux: Login Success
            yield put(loginSuccess());
          }
          // Event: Sign Out
          else if (event.payload.event === 'signOut') {
            // Redux: Logout Success
            yield put(logoutSuccess());
          }
        });
      }
      catch (error) {
        console.log(error);
      }
    };
