# aws
## [1][Amazon Redshift update – ra3.4xlarge instances](https://www.reddit.com/r/aws/comments/fu14lp/amazon_redshift_update_ra34xlarge_instances/)
- url: https://aws.amazon.com/blogs/aws/amazon-redshift-update-ra3-4xlarge-instances/
---

## [2][Face Recognition Attendence with Python AWS Rekognition Raspberry Pi3](https://www.reddit.com/r/aws/comments/fu5hhc/face_recognition_attendence_with_python_aws/)
- url: https://www.reddit.com/r/aws/comments/fu5hhc/face_recognition_attendence_with_python_aws/
---
Face Recognition Attendence with AWS Rekognition Raspberry Pi3

https://github.com/Arbazkhan4712/Face-Recognition-Attendence-with-AWS-Rekognition-Raspberry-Pi3
## [3][AWS management console takes forever to load](https://www.reddit.com/r/aws/comments/fu8e3u/aws_management_console_takes_forever_to_load/)
- url: https://www.reddit.com/r/aws/comments/fu8e3u/aws_management_console_takes_forever_to_load/
---
us-east-2

Tried with Firefox, Opera, Chrome. It's the same.

Anybody facing similar issue?
## [4][Heard of Lightsail? Use Lightsail? Let us know your feedback!](https://www.reddit.com/r/aws/comments/ftpvbb/heard_of_lightsail_use_lightsail_let_us_know_your/)
- url: https://www.reddit.com/r/aws/comments/ftpvbb/heard_of_lightsail_use_lightsail_let_us_know_your/
---
Hello redditors,You may or may not have heard of [Lightsail](https://aws.amazon.com/lightsail/). For those of us who haven’t heard, to quote our marketing folks -“Lightsail is an easy-to-use cloud platform that offers you everything needed to build an application or website, plus a cost-effective, monthly plan.”

TL;DR we provide you basic building blocks like instances, load balancers and databases with an easy-to-use interface (no really!) and simple monthly bundle pricing (just like paying for your netflix). We use the same infra as other AWS services like EC2 - hence same the reliability and availability. Give us a try if you haven’t (first month’s on us).

OK. Those of you who are already using us, Thank you! We from the service team are looking forward for your feedback - What do you like? What do you dislike? Something blocking you from using Lightsail? Something that could improve? You have a feature idea? Is there a cloud technology which could benefit from Lightsailification?Even a “How do I do &lt;this&gt; on Lightsail?” question.ANY and ALL feedback welcome!

As we plan our roadmap for the year ahead, we take this feedback seriously. So, thank you in advance for your feedback!I am joined by /u/mgcataws , /u/emwaws from my team to respond to your feedback.

Thank you and stay safe!

Some resources that may come in handy-Bunch of getting started guides - [https://aws.amazon.com/lightsail/resources/](https://aws.amazon.com/lightsail/resources/)Lightsail pricing- [https://aws.amazon.com/lightsail/pricing/ ](https://aws.amazon.com/lightsail/pricing/)

&amp;#x200B;

EDIT: Thank you all for taking time to give your feedback! Please keep them coming. We will read through every comment and take that feedback. Even if you don't see us comment on your comment, we still have taken your feedback!
## [5][Help to create Aurora Read Replica](https://www.reddit.com/r/aws/comments/fu7jdj/help_to_create_aurora_read_replica/)
- url: https://www.reddit.com/r/aws/comments/fu7jdj/help_to_create_aurora_read_replica/
---
I have a Primary/Master RDS with Mysql 5.7 running and whenever i create a Aurora Read replica, it always create like this:    
https://imgur.com/a/vEEwX2D


why its creating a write role when i clearly click on "Create Aurora Read Replica"?


even if i connect using the 2nd endpoint that says "ro" with "reader" role, i can still see its not read only(value is = 0) by doing this query:  SHOW GLOBAL VARIABLES LIKE 'innodb_read_only';    
(Why it created 2 endpoints anyway?)


I try to find that "role" so i can modify it but i cant find it.


I did try to create a normal mysql replica(Create Read Replica) before this and it was as read only as expected.


I want to use the Aurora as i read that it can use autoscale for up to 15 Read Replica compared to a normal one of 5 replicas without autoscale, Plus i think i dont need HAProxy for Aurora read replica compared to the normal read replica
## [6][AWS Cognito SignUp and LogIn with MobileID / Unique Phone Id](https://www.reddit.com/r/aws/comments/fu6ooj/aws_cognito_signup_and_login_with_mobileid_unique/)
- url: https://www.reddit.com/r/aws/comments/fu6ooj/aws_cognito_signup_and_login_with_mobileid_unique/
---
Hello people

I want to perform the following SingUp/ Login flow with  amazon cognito:

Option 1.) LogIn with E-Mail and Password

Option 2.a) LogIn with unique device id.               

2.b) If account doesn't exist: Sign Up User with unique device id

Basically we let the user acces the App directly and save him as a  user with his unique device id, so that we can save his input, even if  he deletes  and reinstalls the app / the app crashes etc.

We didn't find a working solution as of right now, is there a way to realize this?

Thanks a lot!!
## [7][Is it possible to only use a phone number as a sender for AWS SNS SMS?](https://www.reddit.com/r/aws/comments/fu3uav/is_it_possible_to_only_use_a_phone_number_as_a/)
- url: https://www.reddit.com/r/aws/comments/fu3uav/is_it_possible_to_only_use_a_phone_number_as_a/
---
Currently I'm able to send SMS using AWS SNS but it defaults to using the Sender ID in countries where it is supported. Is it possible to turn the Sender ID off completely in all countries?
## [8][How to get details of terminated EC2 instance?](https://www.reddit.com/r/aws/comments/fttbzn/how_to_get_details_of_terminated_ec2_instance/)
- url: https://www.reddit.com/r/aws/comments/fttbzn/how_to_get_details_of_terminated_ec2_instance/
---
I only have the internal IP of an EC2 instance that has been terminated. I don’t know any other details like the InstanceId. How do I get info (like who created/terminated) on that terminated instance in cloudtrail by private IP?
## [9][Every time a Lambda function fails (synchronous and async) I want to receive an email with the logs and traceback. How can I do this?](https://www.reddit.com/r/aws/comments/ftzf3r/every_time_a_lambda_function_fails_synchronous/)
- url: https://www.reddit.com/r/aws/comments/ftzf3r/every_time_a_lambda_function_fails_synchronous/
---
Currently I've used my tooling to set up a DLQ for each lambda. One error handler lambda processes events from those SNS topics.

So I currently get emails like below, but only for asynchronously invoked lambdas.

&gt; Lambda x was invoked at time y with payload z. The error message was "blah". Here's a URL to the full Cloudwatch logs, but also the logs are printed here, below.

(Note: at my scale one email per lambda won't overwhelm my mail client. And I can get my error handler lambda to deduplicate messages if the same fault happens repeatedly)

The code to achieve this is pretty messy. The SNS payload only includes the error message sometimes. It also doesn't have the complete logs.

But worst of all, DLQ doesn't work for synchronously invokes Lambdas. (E.g. ones serving a Rest API.)

How else could I go about doing this?
How can I trigger a lambda every time a synchronously invoked Lambda fails?
## [10][How to ensure that a role can be assumed but a single resource?](https://www.reddit.com/r/aws/comments/ftllyb/how_to_ensure_that_a_role_can_be_assumed_but_a/)
- url: https://www.reddit.com/r/aws/comments/ftllyb/how_to_ensure_that_a_role_can_be_assumed_but_a/
---
I have a Lambda function so I had to create a role with some permissions and also had to add a trusted relationship to the Lambda service. But how to ensure that just my lambda function can assume that role?

I've tried everything but nothing is working.
