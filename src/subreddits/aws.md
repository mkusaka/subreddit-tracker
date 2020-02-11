# aws
## [1][Ansible over AWS Systems Manager Sessions – a perfect solution for high security environments](https://www.reddit.com/r/aws/comments/f25bxn/ansible_over_aws_systems_manager_sessions_a/)
- url: https://luktom.net/en/e1693-ansible-over-aws-systems-manager-sessions-a-perfect-solution-for-high-security-environments
---

## [2][SQL vs NoSQL](https://www.reddit.com/r/aws/comments/f1x0qk/sql_vs_nosql/)
- url: https://www.reddit.com/r/aws/comments/f1x0qk/sql_vs_nosql/
---
Hey folks, wanted to share a youtube channel that I've been working on dedicated to providing simple and easy to digest tutorials on various AWS services.

My newest video discusses the difference between SQL and NoSQL in the AWS context.

The video is available here: https://youtu.be/ruz-vK8IesE

Thank you!
## [3][Help: RCS &lt;-&gt; EC2 latency? Has anyone seen this issue before?](https://www.reddit.com/r/aws/comments/f277nl/help_rcs_ec2_latency_has_anyone_seen_this_issue/)
- url: https://www.reddit.com/r/aws/comments/f277nl/help_rcs_ec2_latency_has_anyone_seen_this_issue/
---
Hi! I'm a front-end / design guy currently trying to help an AWS customer resolve their database issues (so way out of my depth here!).

They have outsourced their development to an external third-party development company and that company doesn't seem to be able to solve their issue, so I'm calling on Reddit to help!

* They have a MySQL database running on RDS and an Express server running on EC2
* RDS is t2.medium right now
* one of the queries is taking 8sec\~ to respond with data from RDS to EC2
   * the query is very fast (sub 10ms I believe) but the payload is 18MB uncompressed.
   * the third party company is claiming that that 18MB is a huge payload and that the issue is coming from network speed?
      * I've not personally built anything in MySQL in many years so I'm unsure whether this is normally an issue?
      * Surely 18MB would normally transfer very quickly from EC2&lt;-&gt;RDS?

What possible solutions should they be looking at here? Right now we're trying to see if upgrading from **t2.large** to **t3.medium** will fix the problem (the developer company says that this will resolve rate limiting issues, but they've led us down this black hole for months now with nothing fruitful in sight).

My gut instinct is that there's something more sinister at play here?
## [4][502 Bad Gateway only when accessing event body in Lambda](https://www.reddit.com/r/aws/comments/f23yok/502_bad_gateway_only_when_accessing_event_body_in/)
- url: https://www.reddit.com/r/aws/comments/f23yok/502_bad_gateway_only_when_accessing_event_body_in/
---
I have been trying to troubleshoot this issue for hours. I have a nodejs function running in aws lambda hooked up to an api gateway rest api. I am trying to get an id out of the body to use as a parameter in my query to RDS, however if i access the variable, i get a 502 internal server error. I've been using JSONparse to get the body into the code as the proxy integration brings it in as string by default. What is really weird is that i can just return that same body through JSON.stringify(body) and it shows that the id is indeed being passed all the way through the entire process. However if I try using the variable in my code by accessing it from the [body.id](https://body.id) after doing JSONparse, it returns the error. I am lost!
## [5][Some of my node lambdas have errors in their metric - but nothing in their CloudWatch logs](https://www.reddit.com/r/aws/comments/f26an4/some_of_my_node_lambdas_have_errors_in_their/)
- url: https://www.reddit.com/r/aws/comments/f26an4/some_of_my_node_lambdas_have_errors_in_their/
---
We have quite a few lambdas, and roughly half are experiencing this problem. We see logs from all of them, so its unlikely that this is a problem with permissions writing to CloudWatch.

The symptoms are as follows:

We have alarms set up in DataDog to raise an alarm if the aws.lambda.errors metric goes above 1 in a 30 minute window. These alarms have started to be triggered - but we can find no corresponding error in the logs for the lambda -or indeed any lambda.

My questions:

* Why do we not see an error in the log? 
* How would you go about debugging this?

&amp;#x200B;

Thanks in advance for any tips!
## [6][Boto3 Instance Filter?](https://www.reddit.com/r/aws/comments/f27nxg/boto3_instance_filter/)
- url: https://www.reddit.com/r/aws/comments/f27nxg/boto3_instance_filter/
---
Hey Everyone,

I'm over here failing at python and would like some help :) I'm trying to print the instance name into cloudwatch, as we leverage that through our Active Directory and recognize systems based on name.  Here's my guess as to the filter, but it's not working - 

&amp;#x200B;

            instances = ec2.instances.filter(
                Filters=[
                        {'Name' : 'instance-state-name','Values' : ['stopped']},
                        {'Name' : 'Key', 'Values': ['Name']}
                ])
## [7][SES: Tons of abuse messages coming from Yahoo?](https://www.reddit.com/r/aws/comments/f21pc6/ses_tons_of_abuse_messages_coming_from_yahoo/)
- url: https://www.reddit.com/r/aws/comments/f21pc6/ses_tons_of_abuse_messages_coming_from_yahoo/
---
Whenever I send out a newsletter to confirmed opted in customers I get a lot of these messages back.  

I am being told yahoo uses an older authentication method, but I don't have a clue on how to resolve these.

Does anyone have an answer?

Thank you!

`Feedback-Type: abuse`  
`User-Agent: Yahoo!-Mail-Feedback/2.0`  
`Version: 0.1`  
`Original-Mail-From: xxxx`  
`Original-Rcpt-To:` `x``xxx`  
`Received-Date: Wed, 22 Jan 2020 03:38:34 +0000`  
`Reported-Domain: amazonses.com`  
`Authentication-Results: authentication result string is not available`
## [8][Is creating throwaway AWS accounts against terms of use?](https://www.reddit.com/r/aws/comments/f2605h/is_creating_throwaway_aws_accounts_against_terms/)
- url: https://www.reddit.com/r/aws/comments/f2605h/is_creating_throwaway_aws_accounts_against_terms/
---
At work, I'm getting an error in an account that doesn't have a default VPC. At home I want to create a throwaway account and try to reproduce the problem. Is this against the terms of use?
## [9][Feasibility of using fargate spot containers for sessionless services?](https://www.reddit.com/r/aws/comments/f1uplv/feasibility_of_using_fargate_spot_containers_for/)
- url: https://www.reddit.com/r/aws/comments/f1uplv/feasibility_of_using_fargate_spot_containers_for/
---
I'm considering using only fargate spot instances to host a service for my website.  Here are my concerns:

1. Are there periods where no spot containers are available?  If so, the whole premise fails.  I am hoping that there is a constant supply of spot containers, but each instance may only be available for a few minutes.
2. A SIGTERM is issued to the docker container 2 minutes prior to termination.  Does AWS also automatically change the state to draining?  I don't want my ALB to keep sending requests to a container that is going to be terminated
## [10][GuardDuty Functionality](https://www.reddit.com/r/aws/comments/f201l7/guardduty_functionality/)
- url: https://www.reddit.com/r/aws/comments/f201l7/guardduty_functionality/
---
I'm trying to learn GuardDuty, but I can't get it to report on some of the findings.

In my first test account, I setup GuardDuty and immediately tried to get it to produce findings. I've got some for VPC/DNS findings, but non for the user related findings. Since, I've seen that GuardDuty has a 14 day learning period for the non-DNS/VPC related findings. That GuardDuty would learn all that bad behaviour and my attack attempts as regular user behaviour. Woops. 

I've set up a new account, I've created a handful of users, I've waited 14 days (every so often using the users to login and create an S3 bucket, then delete it, etc). And I still can't make some of the findings pop. 

For example, I tried to make this finding happen:  [https://docs.aws.amazon.com/guardduty/latest/ug/guardduty\_unauthorized.html#unauthorized4](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_unauthorized.html#unauthorized4) . I tried to login to the same user account from a machine in France, Australia and Indonesia... All within 5 minutes. No luck. 

Any ideas on what I'm doing wrong? Does anyone have documentation on this learning mode?
