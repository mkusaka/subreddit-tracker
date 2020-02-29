# aws
## [1][React+nodejs: Amplify authenticates the user using cognito, but what is the best way to access protected routes and persist the session for the user?](https://www.reddit.com/r/aws/comments/fb72pq/reactnodejs_amplify_authenticates_the_user_using/)
- url: https://www.reddit.com/r/aws/comments/fb72pq/reactnodejs_amplify_authenticates_the_user_using/
---
There are a few options which I am not too sure on to use as the primary one.

a) Do I pass the JWT from amplify to nodejs backend to access protected routes? - if this is the case, then i am good, I know how to do this now after a previous post.

b) Do I pass the cookies that amplify creates after authenticating? - how best to do this on a node app?

c) Do I stick with the current way amplify persists the userbase using cookies in the browser? Or do I implement a different way? if different what would be the right architecture?

&amp;#x200B;

Sorry if this comes off as a little bit noobish, but I am trying to figure out the most secure way of keeping the user session alive and also accessing the backend for protected routes.
## [2][AWS Noob - Connected a PLC (IoT Device) to AWS as a "thing", how do I go about building a simple app to collect the data?](https://www.reddit.com/r/aws/comments/fb4252/aws_noob_connected_a_plc_iot_device_to_aws_as_a/)
- url: https://www.reddit.com/r/aws/comments/fb4252/aws_noob_connected_a_plc_iot_device_to_aws_as_a/
---
I am using MQTT to pull data off of a PLC system system and I have successfully setup a data stream into AWS - [https://imgur.com/a/mY9MCN3](https://imgur.com/a/mY9MCN3)

What are some resources I should look into to build a very simple iOS and Android app off of this data so it can been seen on a phone? I programmed in JAVA back in 2006 but haven't been programming since so I am just trying to make sometime super basic for the time being. I am building a monitoring system for a PLC system.
## [3][Is it possible to view the number of Invalidation requests I've sent to cloudfront?](https://www.reddit.com/r/aws/comments/fb87k9/is_it_possible_to_view_the_number_of_invalidation/)
- url: https://www.reddit.com/r/aws/comments/fb87k9/is_it_possible_to_view_the_number_of_invalidation/
---
I cant seem to find anything in the Cost explorer or the usage tab withing cloudfront.

I'm trying to stick within the free tier limits so this information would be very useful.

Cheers
## [4][Why do my aws calls not retry with a timeout set?](https://www.reddit.com/r/aws/comments/fb61iq/why_do_my_aws_calls_not_retry_with_a_timeout_set/)
- url: https://www.reddit.com/r/aws/comments/fb61iq/why_do_my_aws_calls_not_retry_with_a_timeout_set/
---
I have some aws client calls in a aws lambda function that have a timeout set to less than the lambda execution time. This is so the lambda doesn’t time out waiting on a dynamo query for example. I set the timeout property of the dynamo client and I see some TaskCanceledException but i don’t believe that the call is being retried since the whole lambda immediately exits around the time the exception happens. Im not sure what setting the retry property of the dynamo client is doing if it doesn’t seem to retry the operation on timeout. Here is the documentation I was using for information about aws sdk and timeouts https://docs.aws.amazon.com/sdk-for-net/v2/developer-guide/retries-timeouts.html#retries

 Am I misunderstanding when retries are supposed to happen? If so am I supposed to implement my own retry logic in the timeout case
## [5][One large VPC with lots of smaller subnets VS lots of smaller VPC](https://www.reddit.com/r/aws/comments/fbcji4/one_large_vpc_with_lots_of_smaller_subnets_vs/)
- url: https://www.reddit.com/r/aws/comments/fbcji4/one_large_vpc_with_lots_of_smaller_subnets_vs/
---
I'm wondering if it makes sense to have one large VPC, say /16, and chopping up into lots of smaller subnets, say /27. Then using 1-4 of said subnets for each of our clients. Or would it be better to have 1 smaller VPC per client, that is chopped in the same fashion. With the route tables on subnets and security groups on instances I'm not seeing much of a benefit, if any, in smaller VPCs. Smaller VPCs feels like just an administrative nightmare,  but I could be wrong.

A few requirements must be met.
1, one of the subnets will allow access from the internet (web tier)
O2, some of our clients require VPN access into one of their subnets. The build needs to support this with us not knowing the client network ahead of time and with bringing new clients on whenever we get them.
3 , all servers are on a domain, so they will all need to be able to access the domain controllers we manage on ec2 instances for all the domainy stuff.
4, we have centralized monitoring servers running that need to be able to access the subnets, as well as jump boxes.

Things to note.
Our application does not really scale so it has single points of failure and auto scaling groups aren't really an option just yet. 
If something goes wrong on a server  we need to fix the server, not delete and rebuild. That is coming but still a ways out.

Any thoughts or documents you can provide would be amazing.
## [6][AWS MediaConvert down for anyone else?](https://www.reddit.com/r/aws/comments/fbbtuk/aws_mediaconvert_down_for_anyone_else/)
- url: https://www.reddit.com/r/aws/comments/fbbtuk/aws_mediaconvert_down_for_anyone_else/
---
I've got jobs that are getting hung up with "PROGRESSING" status. Haven't changed any code myself. I'm processing small videos that normally take a few seconds to transcode.
## [7][Anyway to get list of SES Email recepients?](https://www.reddit.com/r/aws/comments/fbbsxg/anyway_to_get_list_of_ses_email_recepients/)
- url: https://www.reddit.com/r/aws/comments/fbbsxg/anyway_to_get_list_of_ses_email_recepients/
---
Hi All,

We use SES to send email from a our applications to our customers - I was wondering if I could generate a list/report of all sent emails with sender address, recipient address &amp; subject line. 

I honestly cannot find out how to do that - the documentation on creating configration sets &amp; integrating with CloudWatch is frustratingly incomplete (no mention of dimensions, values, etc). and I cannot understand why this has to be so complex, when literally every other SMTP provider can provide this basic report out of the box.

Thanks in advance.
## [8][CodePipeline + CloudFormation + Lambda](https://www.reddit.com/r/aws/comments/faumli/codepipeline_cloudformation_lambda/)
- url: https://www.reddit.com/r/aws/comments/faumli/codepipeline_cloudformation_lambda/
---
Hey all,

I created this post on medium to help the ones who like me had difficulties to understand how to implement a DevOps pipeline in AWS using AWS developer tools (CodeCommit, CodeBuild, CodeDeploy and CodePipeline). And from understanding the DevOps pipeline create an real pipeline using cloudformation and deploying the code to Lambda (Python and Java).

https://link.medium.com/blysiuvJr4

I shared in the post files that I created that you can use to start an sample project in your account.
## [9][Reclone mount inside ECS, doable?](https://www.reddit.com/r/aws/comments/fb9ft3/reclone_mount_inside_ecs_doable/)
- url: https://www.reddit.com/r/aws/comments/fb9ft3/reclone_mount_inside_ecs_doable/
---
I’ll admit I’m a noob here, my only experience with docker is some physical things running on Rancher inside FreeNAS on metal. 

I’m looking to be able to have a common storage to a number of dockers that would be backed by Rclone mount. I believe Rclone mount is available as a docker itself, but not sure if one container can provide storage to multiple other containers?
## [10][AWS Linux restore from AMI](https://www.reddit.com/r/aws/comments/fb2jqs/aws_linux_restore_from_ami/)
- url: /r/linuxadmin/comments/fb2j2g/aws_linux_restore_from_ami/
---

