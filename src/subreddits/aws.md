# aws
## [1][Cognito: Is there a way to verify JWT tokens at front-end using aws-sdk?](https://www.reddit.com/r/aws/comments/eu4ari/cognito_is_there_a_way_to_verify_jwt_tokens_at/)
- url: https://www.reddit.com/r/aws/comments/eu4ari/cognito_is_there_a_way_to_verify_jwt_tokens_at/
---
I'm using aws-sdk at front-end of my web application. Till now, I've set-up the flow to register new users, authenticate users that will get the access token, id token, and refresh token. 

I'm confused about what's next !!!

* The access and id tokens are valid for 1 hour and refresh token for 30days, and all are in JWT format.
* So, in order to check the log-in status of the user,  the access token needs to be parsed to check for the expiration time. Is it possible to do this at front end? 
* Consequently, if expired then using the refresh token will provide fresh access and id tokens.

Basically, I want to check the validity of the tokens and expiration time to maintain user log-in status.
## [2][Golang vs Python when interfacing with S3](https://www.reddit.com/r/aws/comments/etvcy1/golang_vs_python_when_interfacing_with_s3/)
- url: https://www.reddit.com/r/aws/comments/etvcy1/golang_vs_python_when_interfacing_with_s3/
---
Hi all,

I'm going to start writing a long lived process that constantly interfaces with S3 (lists files and uploads files). The program will have \~5 workers that are continuously uploading files to S3 as they become available in the local filesystem. Which programing language would be a better fit for this task, Golang or Python? This is an I/O bounded task, so it seems like Python may not be that bad of a choice. Is Go's AWS SDK more performant than boto3 when interfacing with S3?

Thanks!
## [3][RDS SSL CA upgrade, if you have node in production, check it today](https://www.reddit.com/r/aws/comments/eu0yd3/rds_ssl_ca_upgrade_if_you_have_node_in_production/)
- url: https://medium.com/teamzerolabs/node-rds-mysql-ssl-upgrade-am-i-impacted-b4fdfa7b4677
---

## [4][Is anyone using AWS for email relay?](https://www.reddit.com/r/aws/comments/etwhj5/is_anyone_using_aws_for_email_relay/)
- url: https://www.reddit.com/r/aws/comments/etwhj5/is_anyone_using_aws_for_email_relay/
---
I am looking to use AWS to run a Docker container for email. Email won't actually be stored on the instance. It will more be an email relay using [https://www.simplelogin.io/](https://www.simplelogin.io/).

I can't find any articles of folks using AWS for this? Wondering if anyone is/has does this and what their experience was like.
## [5][My Lambda@Edge is 503'ing in other countries](https://www.reddit.com/r/aws/comments/eu1aj0/my_lambdaedge_is_503ing_in_other_countries/)
- url: https://www.reddit.com/r/aws/comments/eu1aj0/my_lambdaedge_is_503ing_in_other_countries/
---
Not sure why, but I'm in the US and hitting my URL that has Lambda@Edge running on it works fine, but when someone from Germany or Australia hit the URL, they get a 503 error that looks like this:

&amp;#x200B;

https://preview.redd.it/v5woxs8yi1d41.png?width=1366&amp;format=png&amp;auto=webp&amp;s=695ff6f1a1e9ad9783558274d84d7888b39fa5dd

It kind of botched the initial release of my app, so I removed it for now since it was just a download page and a download link, but does anyone know why this happened? I use Lambda@Edge at work also and I'd like to figure out what went wrong.
## [6][Endpoint service and DNS name](https://www.reddit.com/r/aws/comments/etz4ye/endpoint_service_and_dns_name/)
- url: https://www.reddit.com/r/aws/comments/etz4ye/endpoint_service_and_dns_name/
---
How can I customize dns name of endpoint service. Since our org has own dns server we map our custom domain name to dns name created for Private Link( end point service) . If we delete and endpoint and recreate it’s always creating a new dns name which then forces us to raise another request to remap this new dns to existing sub domain. I want to avoid it
How can I achieve ?
## [7][How to use cloud watch with google calendar API.](https://www.reddit.com/r/aws/comments/eu3lq2/how_to_use_cloud_watch_with_google_calendar_api/)
- url: https://www.reddit.com/r/aws/comments/eu3lq2/how_to_use_cloud_watch_with_google_calendar_api/
---
Hey guys complete noob here. How would I use cloud watch so when it is 1 day before an appointment created on google calendar I can use a lambda function to create an SES to the client.
## [8][Turnkey AWS with Paco: Create and Manage a WordPress server](https://www.reddit.com/r/aws/comments/etrfbo/turnkey_aws_with_paco_create_and_manage_a/)
- url: https://medium.com/waterbearcloud/turnkey-aws-with-paco-create-and-manage-a-wordpress-server-part-1-6a54f1b9b337
---

## [9][AWS Loadbalancer wrong rule caused forward](https://www.reddit.com/r/aws/comments/etylx4/aws_loadbalancer_wrong_rule_caused_forward/)
- url: https://www.reddit.com/r/aws/comments/etylx4/aws_loadbalancer_wrong_rule_caused_forward/
---
Hi,

I created a rule to ALB to point a site A to a new EC2 instance using (host based rule)

Then I quickly removed the rule when I realized the new server didn't have any Virtual Block Host block defined for the new site in Nginx configuration and it displayed a wrong site. 

Even I removed the rule from Load balancer, my browser (firefox) still redirects the site like there was the rule still on? So going to site B with a Firefox shows a site X. But with any other browser the site B is showing normali from the original server.  Why is this happening, I never touched the Domain DNS anyway, it has been all the time pointing to ALB DNS.
## [10][Using security groups and iptables](https://www.reddit.com/r/aws/comments/ety648/using_security_groups_and_iptables/)
- url: https://www.reddit.com/r/aws/comments/ety648/using_security_groups_and_iptables/
---
Hello! I have a question regarding the use of security groups and iptables. Specifically, I want to only worry about security groups rules, but I’m curious if iptable rules can bypass security group settings. For example, if my security group only allows port 22,80,443 from the world, but I have an iptables rule to allow port 8081 from the world, will that rule actually work and bypass what is defined in security groups?
