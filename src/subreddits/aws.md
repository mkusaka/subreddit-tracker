# aws
## [1][Can I run amazon-ssm-agent on Flatcar Container Linux?](https://www.reddit.com/r/aws/comments/ijauww/can_i_run_amazonssmagent_on_flatcar_container/)
- url: https://www.reddit.com/r/aws/comments/ijauww/can_i_run_amazonssmagent_on_flatcar_container/
---
Flatcar does not have package managers like yum or apt, so I can't simply install official rpm or deb.

I extracted official rpm, and found that it doesn't contain lots of files. Looks like only files I need for running ssm-agent are:

* binaries in /usr/bin: ssm-document-worker, ssm-session-logger, ssm-session-worker, amazon-ssm-agent, ssm-cli
* /etc/amazon-ssm-agent.service

So If I extract those files from archive, will it work even on Container Linux?

I've also tried running daemonset provided by [https://github.com/mumoshu/kube-ssm-agent](https://github.com/mumoshu/kube-ssm-agent) on one of my kubernetes cluster. While I could access to journald logs or process list with this, since it runs ssm-agent inside container, I could not access to host file system. I couldn't scroll console with up/down cursor key when reading journald logs. PgUp/PgDn didn't work either. Not sure if this is the issue with the docker image.

The docker image does not contain those files which existed in rpm, I'm not sure these files are really needed.

* /etc/amazon/ssm/seelog.xml.template
* /etc/amazon/ssm/amazon-ssm-agent.json.template
## [2][Moving to AWS - How can I get in touch with an Startup Account Manager ?](https://www.reddit.com/r/aws/comments/ij2gto/moving_to_aws_how_can_i_get_in_touch_with_an/)
- url: https://www.reddit.com/r/aws/comments/ij2gto/moving_to_aws_how_can_i_get_in_touch_with_an/
---
We are a steadily growing dev company we want to move to AWS and need some hand holding. We manage web apps for clients and have couple of our own digital products.
## [3][Deploying a flask microservice EBS or ECS](https://www.reddit.com/r/aws/comments/ijch7i/deploying_a_flask_microservice_ebs_or_ecs/)
- url: https://www.reddit.com/r/aws/comments/ijch7i/deploying_a_flask_microservice_ebs_or_ecs/
---
I have created a flask application, which has a connectivity to DynamoDB (so no internal databases). 

I have tried to deploy it to EBS, but i had a lot of issues with configuration, and i found it pretty hard to check the serverlogs, because most of the logs very server (NGINX and WSGI) specific, and it always gave a 502 bad gateway. (I'm new to AWS and trying to learn it, but i think it was hard to setup). 

In the context i wanted to try out ECS, the idea would be, that Docker Images, would be a "safer" solution, because you can almost guarantee, that if it works on your machine, it will work on all machines. 

Furthermore, i think the ECS cluster solution looks very promising, since you can manage multiple docker containers simultaneously, and it makes scalability easier.

We plan to deploy multiple microservices, which needs scalability, would you recommend ECS, or would it be more cost-efficient to deploy to EBS or an EC2 instance.
## [4][AWS Lambda function is being run by an internal system(?)](https://www.reddit.com/r/aws/comments/iiu7d4/aws_lambda_function_is_being_run_by_an_internal/)
- url: https://www.reddit.com/r/aws/comments/iiu7d4/aws_lambda_function_is_being_run_by_an_internal/
---
Hi guys! I've a Cloudwatch event which runs an AWS Lambda every minute. Usually it runs on the 50th second of every minute.

From the Cloudwatch group log I noticed that sometimes the lambda is invoked on unusual time, but my code is not executed. It seems it loads the modules contained in the layer and execute some code which is not mine.
I can see a lot of these string logged:
`[DEBUG] 2020-08-28T19:25:34.220Z Changing event name from creating-client-class.iot-data to creating-client-class.iot-data-plane`

The following minute Cloudwatch runs my code as usual, but it crashes due to time not synced in a library which is probably loaded 2 times because of this behaviour.
## [5][EB environment refuses to terminate](https://www.reddit.com/r/aws/comments/ij7rko/eb_environment_refuses_to_terminate/)
- url: https://www.reddit.com/r/aws/comments/ij7rko/eb_environment_refuses_to_terminate/
---
I have an environment that I will press terminate, it’ll run for 30m, but then stop and do nothing. No EC2 instances tied to it.

More info:

https://forums.aws.amazon.com/thread.jspa?threadID=327056

Would really appreciate the help.
## [6][Can you add extra information such as a user’s name with their email in SES?](https://www.reddit.com/r/aws/comments/ij2va3/can_you_add_extra_information_such_as_a_users/)
- url: https://www.reddit.com/r/aws/comments/ij2va3/can_you_add_extra_information_such_as_a_users/
---
Hello everyone,

I’m looking at switching from Mailjet to AWS SES. Only thing is that I need a way to place extra information to my email list such as a user’s name. This is something that can be easily done in mailjet. Also, if this feature is possible, can you change/add a particular user’s name via some kind of api?

Thanks!
## [7][Multi-cdn with Route53 ?](https://www.reddit.com/r/aws/comments/ij0g28/multicdn_with_route53/)
- url: https://www.reddit.com/r/aws/comments/ij0g28/multicdn_with_route53/
---
Can I use Route53 to setup a multi-cdn with say Cloudfront, Google CDN, Azure CDN together. I guess using a weighted policy while creating a record works but how do I point a weighted record created in aws to a resource outside say Google CDN ? Any help would be appreciated.

Thanks in advance.
## [8][The new route 53 UI is terrible](https://www.reddit.com/r/aws/comments/ii8ts4/the_new_route_53_ui_is_terrible/)
- url: https://www.reddit.com/r/aws/comments/ii8ts4/the_new_route_53_ui_is_terrible/
---
Didn't I already post this? Oh wait no, I'm sorry. That was the new calculator UI.

AWS...please stop with all the wizard nonsense. Again. I don't need a wizard to hold my hand through creating a TXT record. I need something simple, or as you now call it, the "old console". I get the desire to create an experience, but please do it where it is warranted. Who in the community is asking for you to complicate the process of creating DNS records? I would rather you take us back to the days of editing BIND files with VIM than have to work in your new console. And I am not alone! A colleague of mine today just shared his feelings to me about your new console. He said, " real DNS ballers edit BIND files with vim". If you need a wizard to create DNS records, you should not be creating DNS records.
## [9][Over 54,000 scanned NSW (Australia) driver's licences found in open cloud storage - Security - iTnews](https://www.reddit.com/r/aws/comments/iikq97/over_54000_scanned_nsw_australia_drivers_licences/)
- url: https://www.itnews.com.au/news/over-54000-scanned-nsw-drivers-licenses-found-in-open-cloud-storage-552544
---

## [10][AWS Web App Static File Storage](https://www.reddit.com/r/aws/comments/iixw9v/aws_web_app_static_file_storage/)
- url: https://www.reddit.com/r/aws/comments/iixw9v/aws_web_app_static_file_storage/
---
I have a generic question.  We are moving a a dynamic web application from onsite to AWS EC2.  We plan on using multi-AZ RDS and a load balancer.    The web application has A LOT of static files.   My questions is what are the pros and cons of using S3 vs EBS vs EFS?  I know that S3 is WAY cheaper than EBS or EFS.  EBS you have to pay for the storage capacity upfront.   But if you expand EBS, do you have to take the application down to do the expansion?   EFS can dynamically grow.   Does S3 have the capacity (latency, throughput, etc)  to handle a large web application that will be used regionally.   Most of our users will be in one geographic location.   I don't see any reason to use Cloudfront for caching of static file.  Any other recommendation about this setup would be great!  


Clarification:  My wording for "EBS you have to pay for the storage capacity upfront" may be mis-interpreted.  I meant that if you have 5TB of data, but I provisioned the system with 10TB (for growth), you will be paying for the provisioned capacity. 
