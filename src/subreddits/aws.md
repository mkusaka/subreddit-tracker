# aws
## [1][Four months of 100% website uptime! Thanks AWS!](https://www.reddit.com/r/aws/comments/fua8ye/four_months_of_100_website_uptime_thanks_aws/)
- url: https://www.reddit.com/r/aws/comments/fua8ye/four_months_of_100_website_uptime_thanks_aws/
---
We used to host our website using a dedicated server with each web service in separate Docker containers. Although it worked, it wasn't elegant and we often had downtime. We're now using Elastic Beanstalk and Aurora through AWS.

When we initially moved over, our I/O charges with Aurora were horrendous. It forced us to tweak our code to decrease the database queries. As a result, lower charges and our web app runs quicker.

Even after multiple tweaks, AWS is still roughly twice the price of hosting the website on a dedicated server but I've now gained something invaluable which is better sleep.

Thought I'd share my experience in-case anybody else is thinking of switching from dedicated to cloud.
## [2][Internet outage scenario between regions?](https://www.reddit.com/r/aws/comments/fuqmvk/internet_outage_scenario_between_regions/)
- url: https://www.reddit.com/r/aws/comments/fuqmvk/internet_outage_scenario_between_regions/
---
I have two aws regions, one near me and one far from other continent. 

My ISP have problems with their internet that causes extremely lags or timeouts frequently to the far region. 

I must keep them both. How to solve connectivity issue with far region from current region?

Mostly I can’t access our website that written using MERN stack - with over 500mil records in db. 

I’m thinking of putting a cloudfront in front of website’s load balancer?

And maybe use vpc peering? Can vpc peering hep in this case? If yes then how?
## [3][Amazon RDS for SQL Server now supports In-Region Read Replicas](https://www.reddit.com/r/aws/comments/fuepgt/amazon_rds_for_sql_server_now_supports_inregion/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/04/amazon-rds-for-sql-server-now-supports-in-region-read-replicas/
---

## [4][TIFU. Left my instances paused](https://www.reddit.com/r/aws/comments/fuktaa/tifu_left_my_instances_paused/)
- url: https://www.reddit.com/r/aws/comments/fuktaa/tifu_left_my_instances_paused/
---
The bill was awful.

:(
## [5][Stepfunction sending SNS Issue](https://www.reddit.com/r/aws/comments/fuklrt/stepfunction_sending_sns_issue/)
- url: https://www.reddit.com/r/aws/comments/fuklrt/stepfunction_sending_sns_issue/
---
I have a stepfunction which utilizes a lambda function that will call send a email SNS via the sdk.  Unfortunately it doesnt seem to make publish the SNS as I dont receive any emails.  My other lambda functions not inside a stepfunction work as intended.  Does a step function need specific privileges?  Im all this via serverless and my top level iamrolesstatement lets all lambdas publish SNS messages.  Thoughts or suggestions please. Thank you!
## [6][How to set scopes for a JWT Authorizer (from HTTP APIs) with Cognito as issuer?](https://www.reddit.com/r/aws/comments/fujfwa/how_to_set_scopes_for_a_jwt_authorizer_from_http/)
- url: https://www.reddit.com/r/aws/comments/fujfwa/how_to_set_scopes_for_a_jwt_authorizer_from_http/
---
The API Gateway Developer Guide instructs us in [this page](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-jwt-authorizer.html) on how to create their built-in **JWT Authorizers** for **HTTP API** routes.

According to this guide, we can create a JWT Authorizer for a given route. The authorizer evaluates the token by checking the token's **"aud" claim** (which should match acceptable clients configured for the authorizer), as well as the token's **scopes** (which should match acceptable scopes configured for the authorizer) and other token claims (kid, iss, etc).

However, I'm using AWS Cognito as IdP and token issuer. And the tokens we get back from Cognito after login seem to be incompatible with JWT Authorizers: Cognito's **Access Tokens** include scopes but don't include an "aud" claim (they use "client\_id" instead), and their **ID Tokens** include an "aud" claim, but don't include scopes.

I really liked the simplicity of JWT Authorizers. **But how am I supposed to scope my authorizers if Cognito can't give me a workable token**?

Can anybody suggest me a workaround this issue that can target a large user base (like \~50,000 users)?
## [7][Any news on Aurora multi-region multi-master?](https://www.reddit.com/r/aws/comments/fuf8t1/any_news_on_aurora_multiregion_multimaster/)
- url: https://www.reddit.com/r/aws/comments/fuf8t1/any_news_on_aurora_multiregion_multimaster/
---
Back in 2017, AWS announced [Aurora multi-master](https://www.youtube.com/watch?v=4XL1VZymTA8), and that went GA in August last year. In that announcement AWS said that multi-region support for multi-master would come in 2018, but as far as I'm aware nothing else has been announced. Delays are understandable, I'm really just curious. Has anyone heard anything? Anyone know if multi-region Aurora masters are any closer to reality?
## [8][Face Recognition Attendence with Python AWS Rekognition Raspberry Pi3](https://www.reddit.com/r/aws/comments/fu5hhc/face_recognition_attendence_with_python_aws/)
- url: https://www.reddit.com/r/aws/comments/fu5hhc/face_recognition_attendence_with_python_aws/
---
Face Recognition Attendence with AWS Rekognition Raspberry Pi3

https://github.com/Arbazkhan4712/Face-Recognition-Attendence-with-AWS-Rekognition-Raspberry-Pi3
## [9][How to create an encrypted AMI?](https://www.reddit.com/r/aws/comments/fugrju/how_to_create_an_encrypted_ami/)
- url: https://www.reddit.com/r/aws/comments/fugrju/how_to_create_an_encrypted_ami/
---
I started with a standard AMI (unbuntu).  I installed various software, edited config files, etc.  My EC2 instance is working how I want.  Now I want to create an AMI so I can stop the instance and restart it later with all the same setup, possibly on a more capable instance type as circumstances require.  But as some of the config files contain passwords, I need the AMI to be encrypted.  How do I do that?

I have a Customer Master Key, and enabled default encryption, but that does not seem to be enough, as when I created a new AMI from my running instance, the underlying snapshot was marked NOT encrypted.
## [10][Cors issues](https://www.reddit.com/r/aws/comments/fug7hy/cors_issues/)
- url: https://www.reddit.com/r/aws/comments/fug7hy/cors_issues/
---
My lambda api is throwing cors issue from my browser. I have a react front end.
Lambda is nodejs.
I have enabled cors headers in the response.
Access-control-allow-origin but the issue doesn’t go away.
Any tips on how to resolve.
Thanks
