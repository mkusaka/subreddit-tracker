# aws
## [1][AWS Wish List 2020](https://www.reddit.com/r/aws/comments/jbw85e/aws_wish_list_2020/)
- url: https://www.reddit.com/r/aws/comments/jbw85e/aws_wish_list_2020/
---
&amp;#x200B;

AWS always releases a bunch of features, sometimes everyday or atleast once a week. Here is my wish list of the features I want to see as a part of AWS infrastructure

1: AWS Managed Proxy Server(Rather than spinning own squid server)

2: EBS replication across different availability zones(Possible? Legal constraints?)

3: Multi-region VPC(Possible? Legal constraints?)

4: UI to debug boot issues(Better then EC2 Get Instance Screenshot and Instance logs)

5: Support tagging for every individual service(It's improving)

6: VPC endpoints support for every service (EKS?) 

7: EC2 instance live migration

8: Display AWS Cli  while resource creation(Similar to GCP)

9: Cost calculation while resource creation(AWS start supporting(for example, RDS) this feature but not for every service

10: More features in App Mesh(Circuit breaker, Rate Limiting)

P.S: Not sure if some features are already available, but if something is missing, please feel free to add
## [2][New – Amazon RDS on Graviton2 Processors](https://www.reddit.com/r/aws/comments/jbx1sd/new_amazon_rds_on_graviton2_processors/)
- url: https://aws.amazon.com/blogs/aws/new-amazon-rds-on-graviton2-processors/
---

## [3][Can Gitlab be substituted by AWS? (CodeCommit, CodePipeline, CodeBuild)?](https://www.reddit.com/r/aws/comments/jc7ucx/can_gitlab_be_substituted_by_aws_codecommit/)
- url: https://www.reddit.com/r/aws/comments/jc7ucx/can_gitlab_be_substituted_by_aws_codecommit/
---
I'd like to use all AWS tools. 

What are going to be the trade offs? will it be more expensive? More work needed?

What are the decision factors?
## [4][Graviton2 R6g and M6g instances are now generally available for Amazon RDS](https://www.reddit.com/r/aws/comments/jbx1jz/graviton2_r6g_and_m6g_instances_are_now_generally/)
- url: https://www.reddit.com/r/aws/comments/jbx1jz/graviton2_r6g_and_m6g_instances_are_now_generally/
---
[https://aws.amazon.com/about-aws/whats-new/2020/10/achieve-up-to-52-percent-better-price-performance-with-amazon-rds-using-new-graviton2-instances/](https://aws.amazon.com/about-aws/whats-new/2020/10/achieve-up-to-52-percent-better-price-performance-with-amazon-rds-using-new-graviton2-instances/)

[https://aws.amazon.com/blogs/aws/new-amazon-rds-on-graviton2-processors/](https://aws.amazon.com/blogs/aws/new-amazon-rds-on-graviton2-processors/)
## [5][Match multiple alarmName in event pattern of cloudwatch EventBridge?](https://www.reddit.com/r/aws/comments/jc7wtl/match_multiple_alarmname_in_event_pattern_of/)
- url: https://www.reddit.com/r/aws/comments/jc7wtl/match_multiple_alarmname_in_event_pattern_of/
---
I’m setting up event rules that send to lambda when an alarm triggered. 


But I want to match multiple alarms in the event pattern of cloudwatch EventBridge (detail.alarmName), how to?

I tried to use regrex in alarmName like: &lt;fixed-prefix&gt;-* to match all alarms have &lt;fixed-prefix&gt;-in their name but it didn’t work, event didn’t send to lambda.
## [6][AWS IAM:PassRole explained with pictures](https://www.reddit.com/r/aws/comments/jbwkh8/aws_iampassrole_explained_with_pictures/)
- url: https://blog.rowanudell.com/iam-passrole-explained/
---

## [7][Understanding DMS with SQL Server - CDC vs Replication](https://www.reddit.com/r/aws/comments/jc6ut1/understanding_dms_with_sql_server_cdc_vs/)
- url: https://www.reddit.com/r/aws/comments/jc6ut1/understanding_dms_with_sql_server_cdc_vs/
---
Hi,

I've been playing with DMS to replicate my on premises SQL Server to (yet to be determined) anything. But I am restricted to using CDC only and no distributors/publishers/MS-Replication. (I can use MS-Replication for the initial data load, thats fine but want to use CDC for ongoing replication).

According to the docs (see sources below) it says MS-CDC is used for non primary key tables and MS-Replication is used for primary key tables but SQL Server on RDS is only able to use MS-CDC for all tables, MS-Replication isn't an option.

Ideally I would like to use the RDS style for my on-premises hosted SQL Server. Is there an option I am missing somewhere for this, I couldn't see it? Or has anyone else tried to do the same thing?


Sources:

https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Source.SQLServer.html
See the dot points under: "Capturing data changes for SQL Server" and the following sentence.
## [8][Amplify CodePush and AB testing](https://www.reddit.com/r/aws/comments/jc5v34/amplify_codepush_and_ab_testing/)
- url: https://www.reddit.com/r/aws/comments/jc5v34/amplify_codepush_and_ab_testing/
---
Hello all,

I was wondering if AWS amplify has a service similar to microsoft's CodePush. ([CodePush](https://microsoft.github.io/code-push) is a cloud service that enables Cordova and React Native developers to deploy mobile app updates directly to their users' devices)

I was also wondering if there is an AB testing framework within AWS?

I have searched around, haven't found it yet. I assume it does... If any of you are aware of it, I would love it if you could push me in the right direction to learning how to set it up with my React-Native app!

&amp;#x200B;

Thanks!
## [9][Amazon CloudWatch Synthetics launches Recorder to generate user flow scripts for canaries](https://www.reddit.com/r/aws/comments/jbvehv/amazon_cloudwatch_synthetics_launches_recorder_to/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/10/amazon-cloudwatch-synthetics-launches-recorder-to-generate-user-flow-scripts-for-canaries/
---

## [10][aws product to use as varnish cache alternative?](https://www.reddit.com/r/aws/comments/jbrfcv/aws_product_to_use_as_varnish_cache_alternative/)
- url: https://www.reddit.com/r/aws/comments/jbrfcv/aws_product_to_use_as_varnish_cache_alternative/
---
Is there some product I can use in aws that can be used as a varnish cache replacement?    
   
Our un-cached application response speeds are very bad. like very bad.  We use cloudfront currently, but for app limitations cache expires every 2 hours and then requests goes against actual server which is rather painfully slow.   
  
Ive tried to convince my team to use varnish, but I think it's not getting traction because people perceive it as "the old way", "the way we did things before cloud" etc.

Is there is a similar piece of software I can use  - just under aws umbrella? that would be easier sell.
## [11][[serverless][JEST] How to test lambda which handles a publish message on SNS.](https://www.reddit.com/r/aws/comments/jbq35e/serverlessjest_how_to_test_lambda_which_handles_a/)
- url: https://www.reddit.com/r/aws/comments/jbq35e/serverlessjest_how_to_test_lambda_which_handles_a/
---
Hey everyone,

I faced to some issue when I wanted to test my handler which publishes a message on SNS.

Here is my code:

    // my handler 
    export const handler = (event) =&gt; {
        try { 
            await emitDeletionComplete(classified.metadata.classifiedRequestId);
         }                         
    catch(e) { 
            console.error(e);
             throw new Error(e); 
        }
    }
    
    // my SNS service 
    import SNS from 'aws-sdk/clients/sns';
    const snsClient = new SNS({region: process.env.AWS_REGION});
    
    export const emitDeletionComplete = async (id) =&gt; {
      try {
        await snsClient.publish({
          Message: JSON.stringify({
            type: 'DELETE_COMPLETE',
            data: {
              id
            }
          }),
          TopicArn: process.env.SNS_ARN
        }).promise();
      } catch(err) {
        console.error(err, err.stack);
        throw new Error('We do not succeed to publish the message DELETE_COMPLETE to ARN: ' + process.env.SNS_ARN);
      }
    };

When i want to test, i try to do :

    import { handler } from '../../../src/handler/dispatch-deletion-to-legacy';
    import SNS from 'aws-sdk/clients/sns';
    
    jest.mock('aws-sdk/clients/sns', () =&gt; {
      return {
        __esModule: true,
        default: jest.fn(() =&gt; {
          return {
            publish: jest.fn().mockReturnThis(),
            promise: jest.fn(),
          }
        }),
      };
    });
    
    [...]
    it('should delete', () =&gt; {
        let sns = new SNS();
        const event = {
          Records: [
            {
              body: JSON.stringify({...some event...})
            }
          ]
        }
    
        handler(event);
    
        expect(sns.publish().promise).toBeCalledTimes(1);
    });

Apparently, it is never called. I don't get why.Maybe my mock is completely wrong.

i'm stuck with it for few hours now...Any idea how can I mock correctly ?

&amp;#x200B;

thanks
