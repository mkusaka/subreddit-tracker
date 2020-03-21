# aws
## [1][AWS launches initiative to accelerate COVID-19 diagnostics, research, and testing](https://www.reddit.com/r/aws/comments/fluzsm/aws_launches_initiative_to_accelerate_covid19/)
- url: https://blog.aboutamazon.com/innovation/aws-launches-initiative-to-accelerate-covid-19-diagnostics-research-and-testing
---

## [2][IOST Becomes First Blockchain Project Advertised by Amazon AWS China](https://www.reddit.com/r/aws/comments/fmd9br/iost_becomes_first_blockchain_project_advertised/)
- url: https://www.reddit.com/r/aws/comments/fmd9br/iost_becomes_first_blockchain_project_advertised/
---
 

**IOST’s Blockchain as a Service (BaaS) platform** [**announced**](https://medium.com/iost/iost-became-the-first-blockchain-project-interviewed-advertised-by-amazon-aws-china-aac0dab38b57) **on March 20, 2020, that it has strengthened its relationship with AWS China, the global leading cloud computing platform.** 

## Amazon Sponsors IOST

[IOST](https://btcmanager.com/?s=iost), a China-based enterprise-level blockchain system, had already [announced](https://medium.com/iost/iost-now-live-on-aws-launch-your-blockchain-solution-in-5-minutes-6a93632ccfe5) its launch on the Amazon Web Service (AWS) marketplace last year. The news had been received very positively by the community and represented an important step in the development of IOST. Thanks to this collaboration, developers can easily deploy a blockchain infrastructure with dev tools to develop dApps, projects and test out IOST’s high-speed infrastructure with just one click.

[https://btcmanager.com/iost-first-blockchain-project-amazon-aws-china/](https://btcmanager.com/iost-first-blockchain-project-amazon-aws-china/)
## [3][Using Lambda and CloudWatch to Start/Stop Lightsail Instance automatically](https://www.reddit.com/r/aws/comments/fmc9qr/using_lambda_and_cloudwatch_to_startstop/)
- url: https://www.reddit.com/r/aws/comments/fmc9qr/using_lambda_and_cloudwatch_to_startstop/
---
So I'm trying to write a python script to automatically turn on/off my Lightsail instance using Lambda and CloudWatch. But after figuring out one after another error, I'm stuck with `"errorMessage": "Could not connect to the endpoint URL: \"`[`https://lightsail.us-west-1.amazonaws.com/\`](https://lightsail.us-west-1.amazonaws.com/\)`"", "errorType": "EndpointConnectionError"`.

BTW, total Python newbie, but here is my really simple but definitely not working code:

`import boto3`

`session = boto3.Session(`

`aws_access_key_id=ACCESS_KEY,`

`aws_secret_access_key=SECRET_KEY,`

`region_name='ap-southeast-2'`

`)`

`client = boto3.client('lightsail')`

`response = client.start_instance(`   

`instanceName='ap-southeast-2b'`

`)`

I had to hardcoded my access key and secret key in since it throws error when I don't: `"errorMessage": "name 'ACCESS_KEY' is not defined", "errorType": "NameError"` (I have my keys in the environment variables and test event, but I'm not sure why when I run test, it didn't use those).

&amp;#x200B;

One more thing, why the `Encryption in transit` option in the environment variables doesn't stick? I toggled the `Enable helpers for encryption in transit`, selected my key (SYMMETRIC\_DEFAULT), pressed encrypt, and then save, but all of this for nothing, why?
## [4][Slashing CloudFront change propagation times in 2020 – recent changes and looking forward | Amazon Web Services](https://www.reddit.com/r/aws/comments/flzeix/slashing_cloudfront_change_propagation_times_in/)
- url: https://aws.amazon.com/blogs/networking-and-content-delivery/slashing-cloudfront-change-propagation-times-in-2020-recent-changes-and-looking-forward/
---

## [5][school me on aws privatelink please](https://www.reddit.com/r/aws/comments/fmb9lu/school_me_on_aws_privatelink_please/)
- url: https://www.reddit.com/r/aws/comments/fmb9lu/school_me_on_aws_privatelink_please/
---
i'm pretty unfamiliar with this service, and am most interested in how it works with a snowflake database, ie. [https://docs.snowflake.com/en/user-guide/admin-security-privatelink.html](https://docs.snowflake.com/en/user-guide/admin-security-privatelink.html)

i'd like to confirm my understanding of the basics of how it works and how a 3rd party trusted relationship could be set up.

company ACME vpc and company ACME account on snowflake, both on us-east-1, \_without\_ privatelink, i'm assuming the data traverses only the internal amazon network but it takes place in network pipes shared with internet traffic.  ACME is reaching out to their public facing https snowflake account URL [acc-12345.us-east-1.snowflakecomputing.com](https://acc-12345.us-east-1.snowflakecomputing.com).

\*with\* privatelink, a trust relationship is set up between ACME's aws vpc and snowflake, now with a private ACME DNS entry; [acme.snowflakecomputing.com](https://acme.snowflakecomputing.com).  this only resolves from ACME's vpc and for users on their vpn.  when snowflake is accessed via it, the network traffic is routed on internal aws pipes, although in some methodology where the traffic is not shared with other accounts or public data.

at this point [acc-12345.us-east-1.snowflakecomputing.com](https://acc-12345.us-east-1.snowflakecomputing.com) is still accessible to the world.  but it could be restricted to allow only ACME employees to access it by providing whitelisted blocks of ACME IPs?  [https://docs.snowflake.com/en/user-guide/admin-security-privatelink.html#blocking-public-access-optional](https://docs.snowflake.com/en/user-guide/admin-security-privatelink.html#blocking-public-access-optional)

after which both [acc-12345.us-east-1.snowflakecomputing.com](https://acc-12345.us-east-1.snowflakecomputing.com) and [acme.snowflakecomputing.com](https://acme.snowflakecomputing.com) will still work, but the former will still put traffic through shared public pipes and the latter through private ones, and only ACME employees will be able to reach the former address and be able to even resolve the latter address.

if i'm a 3rd party provider providing a cloud service, and want to access ACME's snowflake, i essentially have to do the same as above?  which is either get them to whitelist a subset or block of my IPs to access [acc-12345.us-east-1.snowflakecomputing.com](https://acc-12345.us-east-1.snowflakecomputing.com), or to alternatively set up a privatelink trust between MYCO vpc and the customer's snowflake account, hosting my own DNS to [acme.snowflakecomputing.com](https://acme.snowflakecomputing.com), at which point no IP addresses or ranges must be specified?

in the case of snowflake, it also appears that there is some additional complexity in regards to also needing to make S3 buckets visible as well as some additional client configuration items.

am i understanding this correctly?  or where am i falling apart?
## [6][EC2 Spot instance](https://www.reddit.com/r/aws/comments/fm55q0/ec2_spot_instance/)
- url: https://www.reddit.com/r/aws/comments/fm55q0/ec2_spot_instance/
---
How long can my ec2 spot request stay on pending-fullfilment? It’s been stuck on this for an hour already.
## [7][Stupid question -- where is the Elasticache Redis API for Java? Or am I misunderstanding what that is?](https://www.reddit.com/r/aws/comments/fm9m6e/stupid_question_where_is_the_elasticache_redis/)
- url: https://www.reddit.com/r/aws/comments/fm9m6e/stupid_question_where_is_the_elasticache_redis/
---
Started a new job, was given the task to create a lambda that writes to Redis. Is that such a thing? Or was it shorthand for write to a DynamoDb table fronted by Redis?

Sorry I feel so dumb
## [8][How to use (prediction) from a deployed model in Sagemaker for IOS](https://www.reddit.com/r/aws/comments/fm8ub3/how_to_use_prediction_from_a_deployed_model_in/)
- url: https://www.reddit.com/r/aws/comments/fm8ub3/how_to_use_prediction_from_a_deployed_model_in/
---
Hi, i can't seem to find any documentation on using a deployed model for IOS. Is Amplify framework able to connect to a deployed model in sagemaker and make predictions using that?
## [9][Common use cases for Aurora](https://www.reddit.com/r/aws/comments/fm3tiq/common_use_cases_for_aurora/)
- url: https://www.reddit.com/r/aws/comments/fm3tiq/common_use_cases_for_aurora/
---
What are some common use cases for Aurora? Trying to understand exactly how it stacks up against AWS' db offerings in comparison to the other options and in a more concrete day-to-day use case sense.
## [10][Difficulty Getting Lamda Edge function to work with CF](https://www.reddit.com/r/aws/comments/fm8f3n/difficulty_getting_lamda_edge_function_to_work/)
- url: https://www.reddit.com/r/aws/comments/fm8f3n/difficulty_getting_lamda_edge_function_to_work/
---
I've been having on-going CORS issues with some assets on my client's server. Everything I can find suggests it's the fact that S3/CF does not always x'mit the Vary header, so the workaround is to do it ourselves using Lamda to append to the HTTP response.

So I created this Lamda Function that I sourced from others in a similar situation

    'use strict';
    
    // If the response lacks a Vary: header, fix it in a CloudFront Origin Response trigger.
    
    exports.handler = async (event) =&gt; {
        const response = event.Records[0].cf.response;
        const headers = response.headers;
        
        console.log("vary headers before", headers['vary']);
    
        if (!headers['vary'])
        {
            headers['vary'] = [
                { key: 'Vary', value: 'Access-Control-Request-Headers' },
                { key: 'Vary', value: 'Access-Control-Request-Method' },
                { key: 'Vary', value: 'Origin' },
            ];
        }
        
        console.log("vary headers after", headers['vary']);
        // Return modified response
        return response
    }

Note, I added the console.log to test that it was actually working. When I run it through the Test engine in Lamda's UI, it works exactly how I expect, the console.log()s fire and output. Perfect. It even shows up in CloudWatch logs.

So I deploy it to my CF distro, I select "Origin Response" as the event and save it. I go over to CF and see that it's deploying. Great. After it is finished deploying, I run an Invalidation on \* (not sure if this step is necessary if CF is being re-deployed).

Once the Invalidation completes, I go test a couple URLs via curl. The Vary headers are not there. The file loads, as expected, but no Vary headers.

I checked CloudWatch and there are no logs about it at all. No errors, no successes, nothing -- other than the test event I did that did work. I checked all US based regions, all of the service groups, no relevant logs anywhere, which tells me that the lamda function isn't be invoked at all.

I double checked the CF config. Checked the behavior tab and confirmed that there is a Origin Response trigger that is indeed pointing to the lamda function and the correct version.

&amp;#x200B;

I'm all out of ideas at this point, especially being an admitted AWS newb. I'd appreciate any pointers!

&amp;#x200B;

**UPDATE:** Looks like the non-logging might have been a permissions issue. I noticed some warnings in the IAM role relating to the CloudWatch needing resources specified or w/e. So I added some wildcards, and boom, the logging happens. But still no header rewriting. Might there be a permission for rewriting headers?
