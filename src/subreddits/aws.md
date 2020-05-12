# aws
## [1][Tell us more about your AWS Step Functions business use cases.](https://www.reddit.com/r/aws/comments/gi9xi3/tell_us_more_about_your_aws_step_functions/)
- url: https://www.reddit.com/r/aws/comments/gi9xi3/tell_us_more_about_your_aws_step_functions/
---
How many of you are using AWS Step Functions **to handle business logic**? Are you satisfied? Where do you develop the state machines? What tooling/frameworks to you use? How do you handle refactors? Do you test your code? How much do you pay? Is it worth the hassle? Is local development hell or is it actually doable/mockable?

I personally liked my experience with AWS Step Functions to automate some "technical burden" like replacing old EC2 AMIs or upgrading a database. Step Functions also play huge part in automating machine learning in AWS (as can be read in AWS WAF Machine Learning Lens). I have yet to meet anyone using it for business logic though and this is how AWS advertises it(mostly). 

I'd love to hear your stories.
## [2][EC2 M6g Instances, powered by AWS Graviton2](https://www.reddit.com/r/aws/comments/ghxy21/ec2_m6g_instances_powered_by_aws_graviton2/)
- url: https://aws.amazon.com/blogs/aws/new-m6g-ec2-instances-powered-by-arm-based-aws-graviton2/
---

## [3][Serverless Authentication with AWS Amplify: A Practical Guide](https://www.reddit.com/r/aws/comments/gi6v30/serverless_authentication_with_aws_amplify_a/)
- url: https://iamondemand.com/blog/serverless-authentication-with-aws-amplify-a-practical-guide/
---

## [4][How to put files on S3 that only be downloaded by my code on any machine.](https://www.reddit.com/r/aws/comments/gi9g4u/how_to_put_files_on_s3_that_only_be_downloaded_by/)
- url: https://www.reddit.com/r/aws/comments/gi9g4u/how_to_put_files_on_s3_that_only_be_downloaded_by/
---
I want to put my deep learning models on S3, and download them wherever I put my code. I obviously can't make them public as it's our IP. What is the way?
## [5][When Should You Use AWS Spot Instances?](https://www.reddit.com/r/aws/comments/gia7y7/when_should_you_use_aws_spot_instances/)
- url: https://www.cloudforecast.io/blog/are-aws-spot-instances-worth-it-in-production/?utm_source=reddit&amp;utm_medium=post&amp;utm_campaign=spot
---

## [6][C# .NET Core Serverless Application multiple origin CORS header](https://www.reddit.com/r/aws/comments/gi8f4k/c_net_core_serverless_application_multiple_origin/)
- url: https://www.reddit.com/r/aws/comments/gi8f4k/c_net_core_serverless_application_multiple_origin/
---
Hi, I'm wondering if anyone can help me with a CORS issue I've been unable to resolve

I have a serverless application utilising AWS lambda and API Gateway to provide a GraphQL API that needs to be able to:

Inspect the origin header sent by the browser, check it against a whitelist of origins, if it matches, return the incoming origin as the Access-Control-Allow-Origin header, else return a default origin. 

This is because some axios calls from the client require the withCredentials flag to be true, which means we can not specify a wildcard for this header in the serverless.template, but due to having a dev, stage and production environment (and the want to be able to call the endpoints from an unspecified origin such as testing requests using Insomnia) we need the ability to respond with the appropriate origin header when it matches a specified entry in the whitelist.

The serverless.template json file has a globals section for CORS settings which only allows us to specify one origin

    "Api": { "Cors": { "AllowMethods": "'GET,OPTIONS,POST'", "AllowHeaders": "'Content-Type'", "AllowOrigin": "'https://somedomain.com'", "AllowCredentials": "true", "MaxAge": "'600'" } }

A .net core serverless application has a LambdaEntryPoint and then the middleware in startup.cs but I'm not sure where I am able to create an APIGatewayProxyResponse to provide the logic to check the incoming origin against the whitelist and then return the appropriate header to satisfy the client requirements and not upset API Gateway in the process 

The request pipeline is handled by dotnet graphql middleware, in the startup.cs using app.useGraphQL&lt;Schema&gt;("/graphql") - there isn't a controller for example that we can manipulate,

So the question is in a serverless application in .NET Core, where is the correct place to manipulate the headers.

Thanks
## [7][Serverless Thought Leaders Converge Online on May 21 &amp; May 28](https://www.reddit.com/r/aws/comments/ghroot/serverless_thought_leaders_converge_online_on_may/)
- url: https://www.reddit.com/r/aws/comments/ghroot/serverless_thought_leaders_converge_online_on_may/
---
Hi, r/aws!

The first-ever AWS Serverless-First Function is a free, virtual event intended to help you learn the latest about serverless approaches on AWS. The event occurs over two Thursdays, May 21 and May 28, and each day focuses on specific topics.

* May 21: Learn about adopting serverless across your organization
* May 28: Learn about building serverless into your applications

Live chat is included on both days – so bring all of your questions to be answered by serverless experts!

You can check out the full session agenda and speaker and event details on the info page: [https://pages.awscloud.com/GLOBAL-event-OE-serverless-first-function-2020-reg-event.html](https://pages.awscloud.com/GLOBAL-event-OE-serverless-first-function-2020-reg-event.html)
## [8][m6g instances showing up in Oregon region](https://www.reddit.com/r/aws/comments/ghv4jd/m6g_instances_showing_up_in_oregon_region/)
- url: https://www.reddit.com/r/aws/comments/ghv4jd/m6g_instances_showing_up_in_oregon_region/
---
It looks like they're Generally Available now. They're in the price list, and I'm running one now... they're also in Virginia, Ohio, Ireland, Frankfurt, and Tokyo... they're available as reserved instances...

Update: here's the [blog post](https://aws.amazon.com/blogs/aws/new-m6g-ec2-instances-powered-by-arm-based-aws-graviton2/). And the [What's New](https://aws.amazon.com/about-aws/whats-new/2020/05/amazon-ec2-m6g-instances-powered-by-aws-graviton2-processors-generally-available/) post.
## [9][Awesome AWS Workshops](https://www.reddit.com/r/aws/comments/ghfgx7/awesome_aws_workshops/)
- url: https://www.reddit.com/r/aws/comments/ghfgx7/awesome_aws_workshops/
---
Following the idea of Awesome Lists, I went about our public workshops that we could find and created this weekend project to help the community. Feel free to add your own workshops, guides, and walkthroughs.

[https://awesome-aws-workshops.com/](https://awesome-aws-workshops.com/)
## [10][Policing sandbox EC2 instances](https://www.reddit.com/r/aws/comments/gi15lz/policing_sandbox_ec2_instances/)
- url: https://www.reddit.com/r/aws/comments/gi15lz/policing_sandbox_ec2_instances/
---
So my boss has asked me to reign in the costs of an account that is shared amongst many people to try and learn AWS things.

One issue we'd like to address is EC2 instances that are left running. I know I could setup an AWS Config rule to perhaps kill running EC2s, though the approach appears to be:
If EC2 instance is costing a lot and running more than a day, get a notification of what, who, when and then follow up by email in a "policing" manner.

I've enabled CloudTrail, and I am wondering how on earth do I parse this to get the "costing a lot and running more than a day" event in order to act?!
