# aws
## [1][I would like to build and deploy a full-stack service (frontend, backend, DB) on AWS. Which are the current best practice? Which AWS service should I use? How to configure them? (details in post)](https://www.reddit.com/r/aws/comments/gcnw4e/i_would_like_to_build_and_deploy_a_fullstack/)
- url: https://www.reddit.com/r/aws/comments/gcnw4e/i_would_like_to_build_and_deploy_a_fullstack/
---
Hi everyone, I'm totally new to the AWS world and I would like to understand how I could start deploying and managing apps. In this scenario, I would like to understand how I could easily start but at the same time be ready to scale. I'm willing to read a lot of resources, that's not an issue.

At the moment I find AWS really overwhelming because you have a lot of options but not a real best practice if you need to build a full-stack architecture with X, Y, Z technology (and maybe this is all because I'm new to this).

My current stack would be

* \[Frontend\] VueJS for the frontend (with Nuxt)
* \[Backend\] Express on the backend (willing to understand how to properly use lambdas with express)
* \[DB\] DynamoDB as a backend (if I don't see any limitations I would love to learn how to use it properly)

So I need to start from scratch everything on AWS. I would like to understand which services should I use for this architecture, how I should configure them (is there a framework/tool to manage them), and how I can configure a CI/CD for both the frontend and backend part.

I saw a lot of tutorials and a lot of technologies out there (like Amplify) but I didn't see the real best-practice I should follow, something that is a de-facto standard that will be maintained by AWS/community over the time.

I try to go into some details about my doubts

# Frontend

* should I host it on S3? I think that this is the common way to host Vue/React/AngularJS apps
* what should I do in order to set up a proper CI/CD env and not manually upload files on the S3 bucket?
* what should I do in order to be ready (something that could eventually never happen but I would like to understand best practice) in order to scale?
* how should the frontend connect with the backend (API Gateway?)? Are there some configurations that I should do?

# Backend

I used Lambda (single function, nothing complex, only for an Alexa skill) in the past and I really like it. I don't know if it's a good fit for every project (compared to EC2) but I would like to understand which are its limitation and where and when to use it.

So my backend would handle everything from user sign in, the user signs up and all the API endpoints to manage my DB entities. I saw that you can use express with lambda thanks to [AWS Serverless Express](https://github.com/awslabs/aws-serverless-express).

* Is express a good solution with lambda or should I use something else in order to use all the lambda's power?
* Which should be the best way in order to manage API Gateway / Lambda configuration in order to be exposed to my frontend?
* Which should be the best way to handle the CI / CD of my backend?
* How should I architecture my system in order to be ready to scale to different regions?
* I saw that many posts refer to ClaudiaJS. Is this still used or has AWS developed something to be used internally?

Maybe I'm just overthinking but I know that I need to understand how things work and how to manage them in order to not be error-prone or go crazy to manually do stuff (because that's the way to make things not scalable).
## [2][I need your help!!! Just charged for EC2 but I haven't touched AWS in years.](https://www.reddit.com/r/aws/comments/gcizgm/i_need_your_help_just_charged_for_ec2_but_i/)
- url: https://www.reddit.com/r/aws/comments/gcizgm/i_need_your_help_just_charged_for_ec2_but_i/
---
I honestly haven't touched AWS in years. But I just received a $400 bill for using EC2 last month, AWS is now saying that I may receive around $500 bill this month if I don't get it resolved, and honestly I'm dead broke, I can't afford that! I checked EC2 and it says there aren't any instances running. What should I do????

Edit: I sent a request to support, but they aren't getting back to me. Cost Management won't show me anything since it's the first time i'm accessing it (says it needs 24 hours).
## [3][Anybody using Config Engine for Compliance as code?](https://www.reddit.com/r/aws/comments/gcpxb4/anybody_using_config_engine_for_compliance_as_code/)
- url: https://www.reddit.com/r/aws/comments/gcpxb4/anybody_using_config_engine_for_compliance_as_code/
---
Is anyone using the below CF template/s and able to deploy Config rules successfully to other AWS accounts from the master account?

[AWS Config Engine as Code](https://github.com/awslabs/aws-config-engine-for-compliance-as-code)

Everything seems to be setup okay but when I follow the instructions on how to update the rule set in my master account, they don’t get pushed to my application accounts.
## [4][AWS Lambda + SSM - how to properly handle SSM value change?](https://www.reddit.com/r/aws/comments/gcneer/aws_lambda_ssm_how_to_properly_handle_ssm_value/)
- url: https://www.reddit.com/r/aws/comments/gcneer/aws_lambda_ssm_how_to_properly_handle_ssm_value/
---
Basically, I wan't to pass some easily-configurable parameters to the Lambda. This Lambda will consume quite a lot of traffic so I don't want to query SSM every time a request is handled. I thought about loading the param when the Lambda is initialized, outside the handler. 

I have found this article which does exactly it:

https://aws.amazon.com/blogs/compute/sharing-secrets-with-aws-lambda-using-aws-systems-manager-parameter-store/

However, when the SSM configuration changes, warm Lambdas will never get an update. This article mentions that:

&gt;  For more advanced use cases where configuration changes need to be received immediately, you could implement an expiry policy for your configuration entries or push notifications to your function.

I could just persist something like "last update date" in the "local configuration"(outside the handler) and whenever Lambda runs, it first checks if the date is older than lets say 5 minutes - if that is the case, it re-queries the SSM and saves the new param. This would basically be the "expiry policy" they're talking about and wouldn't cost that much.

Where can I find any examples of implementing "push notification to the function" though? Would that reconfigure the Lambda near real time?

Perhaps there's an easier way to achieve all that? Thanks.
## [5][AWS Combination Best Option for Django + React](https://www.reddit.com/r/aws/comments/gcpqcd/aws_combination_best_option_for_django_react/)
- url: https://www.reddit.com/r/aws/comments/gcpqcd/aws_combination_best_option_for_django_react/
---
 I have my backend written in django. I researched and understand that  AWS EC2 or AWS lightsail are cost effective solutions for me.

What I am confused is about frontend. Should I use the same instance or create a container and use Amazon Container services ?

The concerns I have is flexible load of containers during multiple  users coming to website, CORS/Same origin issues when deployed in same  instance, security issues when deployed in same instance, cost  effectiveness

Please help how do you decide in this situation

  
Other points -

I estimate very small load on both backend and frontend \~ 2000 visits per month with each visit having not more than 5 api calls to backend with each calls not more than 15kb of data transfer.

For my static media files, I will be using S3 storage.
## [6][EMR Cluster Config typo?](https://www.reddit.com/r/aws/comments/gcp3ic/emr_cluster_config_typo/)
- url: https://www.reddit.com/r/aws/comments/gcp3ic/emr_cluster_config_typo/
---
I'm just wondering if this is a typo:  

[From AWS Docs on EMR Cluster Configuration](https://preview.redd.it/9eq3garncjw41.png?width=1579&amp;format=png&amp;auto=webp&amp;s=15fb9983095eb9163eaf670f7e128c29ddf226d4)

  
Taken from: [https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-plan-instances-guidelines.html](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-plan-instances-guidelines.html)
## [7][Issue with deployment of node app to Elastic Beanstalk](https://www.reddit.com/r/aws/comments/gcm9q7/issue_with_deployment_of_node_app_to_elastic/)
- url: https://www.reddit.com/r/aws/comments/gcm9q7/issue_with_deployment_of_node_app_to_elastic/
---
OS of my PC: Windows 10

I am a beginner in AWS and have a node Application uses RDS and DynamoDB

When the app was deployed to ec2 , there was an error when running the app which said something on the lines of No credentials provider available, credentials are missing.

So I created a file called aws\_config.json at the root folder level, and added the accessKeyId, secretAccessKey and region and in the RDS Store js file retrieve the credentials from the path

const AWS = require('aws-sdk')

AWS.config.loadFromPath('/home/ec2-user/pizza-luvrs/aws\_config.json')

This solved the credentials error. 

Then when I deployed the application to Elastic Beanstalk, I am getting 502 error. 

When I check the logs for the reason behind the error, I see this log: 

&amp;#x200B;

May  2 18:33:53 ip-10-0-1-209 web: Error: EACCES: permission denied, open '/home/ec2-user/pizza-luvrs/aws\_config.json'

May  2 18:33:53 ip-10-0-1-209 web: at Object.openSync (fs.js:458:3)

May  2 18:33:53 ip-10-0-1-209 web: at Object.readFileSync (fs.js:360:35)

May  2 18:33:53 ip-10-0-1-209 web: at Object.readFileSync (/var/app/current/node\_modules/aws-sdk/lib/util.js:95:26)

May  2 18:33:53 ip-10-0-1-209 web: at Config.loadFromPath (/var/app/current/node\_modules/aws-sdk/lib/config.js:436:39)

May  2 18:33:53 ip-10-0-1-209 web: at Object.&lt;anonymous&gt; (/var/app/current/data/pizzaStore.js:16:12)

May  2 18:33:53 ip-10-0-1-209 web: at Module.\_compile (internal/modules/cjs/loader.js:1156:30)

May  2 18:33:53 ip-10-0-1-209 web: at Object.Module.\_extensions..js (internal/modules/cjs/loader.js:1176:10)

May  2 18:33:53 ip-10-0-1-209 web: at Module.load (internal/modules/cjs/loader.js:1000:32)

May  2 18:33:53 ip-10-0-1-209 web: at Function.Module.\_load (internal/modules/cjs/loader.js:899:14)

May  2 18:33:53 ip-10-0-1-209 web: at Module.require (internal/modules/cjs/loader.js:1042:19)

May  2 18:33:53 ip-10-0-1-209 web: at require (internal/modules/cjs/helpers.js:77:18)

May  2 18:33:53 ip-10-0-1-209 web: at Object.&lt;anonymous&gt; (/var/app/current/data/pizzas.js:5:20)

May  2 18:33:53 ip-10-0-1-209 web: at Module.\_compile (internal/modules/cjs/loader.js:1156:30)

May  2 18:33:53 ip-10-0-1-209 web: at Object.Module.\_extensions..js (internal/modules/cjs/loader.js:1176:10)

May  2 18:33:53 ip-10-0-1-209 web: at Module.load (internal/modules/cjs/loader.js:1000:32)

May  2 18:33:53 ip-10-0-1-209 web: at Function.Module.\_load (internal/modules/cjs/loader.js:899:14)

May  2 18:33:53 ip-10-0-1-209 web: at Module.require (internal/modules/cjs/loader.js:1042:19)

May  2 18:33:53 ip-10-0-1-209 web: at require (internal/modules/cjs/helpers.js:77:18)

May  2 18:33:53 ip-10-0-1-209 web: at Object.&lt;anonymous&gt; (/var/app/current/data/mock.js:1:16)

May  2 18:33:53 ip-10-0-1-209 web: at Module.\_compile (internal/modules/cjs/loader.js:1156:30)

May  2 18:33:53 ip-10-0-1-209 web: at Object.Module.\_extensions..js (internal/modules/cjs/loader.js:1176:10)

May  2 18:33:53 ip-10-0-1-209 web: at Module.load (internal/modules/cjs/loader.js:1000:32) {

May  2 18:33:53 ip-10-0-1-209 web: errno: -13,

May  2 18:33:53 ip-10-0-1-209 web: syscall: 'open',

&amp;#x200B;

I tried to resolve this issue by creating .ebextensions folder and adding a config file there which has the contents:

&amp;#x200B;

files:

  "/opt/elasticbeanstalk/hooks/appdeploy/post/99\_fix\_node\_permissions.sh":

mode: "000755"

owner: root

group: root

content: |

\#!/bin/bash

chown root:root /home/ec2-user/pizza-luvrs/aws\_config.json

&amp;#x200B;

Still getting 502 Bad Gateway error. I have also tried rebooting the servers in beanstalk, and have used the eb deploy command line method of deployment too.

&amp;#x200B;

What could be the solution to this?
## [8][IAM authentication with Postgres on RDS](https://www.reddit.com/r/aws/comments/gcfvjj/iam_authentication_with_postgres_on_rds/)
- url: https://www.reddit.com/r/aws/comments/gcfvjj/iam_authentication_with_postgres_on_rds/
---
Just curious, has anyone here used it and does it work well? There are some warnings about connection limitations in the doc. Also, does anyone know if this works with Aurora Postgres too?
## [9][API Gateway Custom Domain Names not working](https://www.reddit.com/r/aws/comments/gcb2iu/api_gateway_custom_domain_names_not_working/)
- url: https://www.reddit.com/r/aws/comments/gcb2iu/api_gateway_custom_domain_names_not_working/
---
I followed the documentation and the Custom Domain Names feature for API Gateway is not working. When I hit the URL (the auto-generated custom domain from AWS) OR the pretty domain I pointed it at, the page refuses to connect.

I followed these instructions for a REST API and Regional.

 [https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html)
## [10][After WorkSpaces upgrade, Alt+Tab no longer works](https://www.reddit.com/r/aws/comments/gc8p1h/after_workspaces_upgrade_alttab_no_longer_works/)
- url: https://www.reddit.com/r/aws/comments/gc8p1h/after_workspaces_upgrade_alttab_no_longer_works/
---
After months of putting it off, I clicked to allow my Amazon WorkSpaces client to upgrade to version 3.0.6.1437. Now when I login, Alt+Tab no longer cycles through applications on my cloud PC; it cycles through them on my LOCAL PC. I find this immensely irritating, just as much as not being able to find any settings that allow me to change it.

Does anyone know how I can get my previous functionality back so when I Alt+Tab, it cycles through apps on my AWS cloud machine? Thank you!
