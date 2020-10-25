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
## [2][Many many MySQL Aurora database issues. Anyone else?](https://www.reddit.com/r/aws/comments/jhi0yi/many_many_mysql_aurora_database_issues_anyone_else/)
- url: https://www.reddit.com/r/aws/comments/jhi0yi/many_many_mysql_aurora_database_issues_anyone_else/
---
We've had a frustrating few months with AWS Aurora (various version up to 2.08.2) . Some issues we've had:

- Under heavy loads, database returns incorrect response (response was supposed to be delivered to a different connection for a different query). Mismatched responses to requests.
Workaround: we found we had an expensive recursive query for circular dependencies existed. Pulling this data from the database would cause the database or Aurora proxy to become confused, and return incorrect responses for small random percentage of various queries moving forward. Fix would be to restart the database or failover between the master and replica.

- Refuses connections even though the database is only using 200 out of 2,000 allowed (XLarge instance). Database metrics across the board (CPU %, Memory %, Disk Space) are all within 10% of actual capacity
Seems to happen where we are creating a high number of connections (and then releasing them after). Seems to be a bug where Aurora or the Aurora proxy believes it has a higher number of connections than it actually has. Aurora dashboard reports a very low number of connections. Could be related to the proxy/load balancer Aurora automatically configures between our application servers and the database server.

- Locks on table rows being acquired and never being released
Every few weeks the database would be unable to acquire a row lock for one of our tables we were using to increment a counter in a thread-safe way. Simple query "SELECT col FROM table WHERE &lt;bool&gt; FOR UPDATE;" which is supposed to lock a single row and implicitly unlock at the end of transaction did not appear to happen. in these times, "SHOW ENGINE INNODB STATUS" would show zero deadlocks.

We never had these problems with a plain-vanilla MySQL 5.7 Community Edition.  It's frustrating that we cannot find anyone with similar experiences to us on the internet. Seems large companies with presumably more traffic are using Aurora without any issues. A lot of these issue do sound like they could be related to the proxy between our application servers and the Aurora database, has anyone had anything at all similar?

We have not tried updating to 2.09.0 yet, since we are at the end of our rope, and looking to move back to either RDS or our own hosted version of MySQL but I see it fixes an [extensive list of deadlock issues](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Updates.2090.html) which makes me feel we cannot be the only ones
## [3][Suggestions AutoScalingGroup to serve as little 50x errors as possible](https://www.reddit.com/r/aws/comments/jht18t/suggestions_autoscalinggroup_to_serve_as_little/)
- url: https://www.reddit.com/r/aws/comments/jht18t/suggestions_autoscalinggroup_to_serve_as_little/
---
Hi, hope someone can give me some suggestions please.

I try to set up an AutoScalingGroup in order to serve as little 50x errors as possible when scaling aggressively to accommodate huge spikes of traffic, for example -thousands of users online within a couple of minutes.
## [4][Slack outage Oct 5, 2020 CoE?](https://www.reddit.com/r/aws/comments/jhswuc/slack_outage_oct_5_2020_coe/)
- url: https://www.reddit.com/r/aws/comments/jhswuc/slack_outage_oct_5_2020_coe/
---
Why do neither Slack nor AWS offer a root cause of the performance issues [1] on Oct 5, 2020? Was it something Slack misconfigured or did AWS drop the ball? Since Slack is “all in” on AWS, this is disconcerting. 

[1] https://www.cnn.com/2020/10/05/tech/slack-down-today/index.html
## [5][Organization’s dns server via transit gateway is not resolving AmazonAWS.com addresses](https://www.reddit.com/r/aws/comments/jhsw8h/organizations_dns_server_via_transit_gateway_is/)
- url: https://www.reddit.com/r/aws/comments/jhsw8h/organizations_dns_server_via_transit_gateway_is/
---
As part of AWS organizations, my IT setup my dhcp options sets to to have my company’s dns servers. However I cannot resolve amazonaws.com addresses with these Name servers. While the issue could be my IT, I want to make sure I’ve done all i can on my end. 

Can I modify resolv.conf to be able to use both my orgs dns servers and Amazon provided dns servers ? Or something with route 53?
## [6][AWS AppSync + WAF + Cognito + DynamoDB (CDK + CF)](https://www.reddit.com/r/aws/comments/jhrcgu/aws_appsync_waf_cognito_dynamodb_cdk_cf/)
- url: https://www.reddit.com/r/aws/comments/jhrcgu/aws_appsync_waf_cognito_dynamodb_cdk_cf/
---
&amp;#x200B;

https://preview.redd.it/7gzmkp3iz7v51.jpg?width=5616&amp;format=pjpg&amp;auto=webp&amp;s=f63dd95aadda9b372e3b8ac0467d7fe856b5d236

Hi Again,

AWS announced support for WAF with AppSync this month so I thought it would be a good time for a write up. There are very few AppSync CDK examples so I've also put a bit of effort into this one to provide a decent example.

If you don't know what AppSync is it is a managed GraphQL API, it's gaining popularity in the mobile app dev space with support for iOS.

The CDK stack is open source on github and for those who just want to deploy the example there is a deploy stack button to get started in a single click.

High-level overview:

* AppSync
   * Basic graphQL API
   * Cognito enabled
   * WAF enabled
* WAF
   * Associated to AppSync endpoint
   * Three rules, Geo, Rate, Core
* DynamoDB
   * As a datasource with our AppSync app
* Cognito
   * Local cognito users, no guest logins

Access the gitub repo here:

[https://github.com/talkncloud/aws/tree/main/appsync-waf](https://github.com/talkncloud/aws/tree/main/appsync-waf)

If you want to read about this stack with more detail on AppSync you can here, if you're deploying WAF to AppSync check it out because the core rule set denies AppSync requests:

[https://www.talkncloud.com/aws-appsync-with-waf-wooo-cdk-cf/](https://www.talkncloud.com/aws-appsync-with-waf-wooo-cdk-cf/)

Keen for feedback and to hear how others are tackling AppSync with CDK and/or CF.
## [7][EC2 Request Ticket always denied or can't be processed](https://www.reddit.com/r/aws/comments/jhr8d4/ec2_request_ticket_always_denied_or_cant_be/)
- url: https://www.reddit.com/r/aws/comments/jhr8d4/ec2_request_ticket_always_denied_or_cant_be/
---
Title, I'm pretty much requesting one GPU instance to do my deep learning and work on, but it either gets denied, or can't be processed. I've been sending tickets for over a month now, what do I do?
## [8][Reasons for different Cloudfront distributions for different S3 buckets](https://www.reddit.com/r/aws/comments/jhqqqw/reasons_for_different_cloudfront_distributions/)
- url: https://www.reddit.com/r/aws/comments/jhqqqw/reasons_for_different_cloudfront_distributions/
---
Hi, 

What are some reason why you would (or  wouldn't) have multiple cloudfront distributions for different buckets? Part of our business have multiple buckets served by the same CloudFront distribution working fine but would rather have them split out but can't really come up with a reason why.
## [9][AWS Amplify JavaScript/iOS Resources](https://www.reddit.com/r/aws/comments/jhqeig/aws_amplify_javascriptios_resources/)
- url: https://www.reddit.com/r/aws/comments/jhqeig/aws_amplify_javascriptios_resources/
---
Hey everyone, I’ve been working with AWS Amplify quite a bit and decided to help people that are interested in creating an app or web app with an AWS backend get started.

Initializing an AWS Amplify project can be a little confusing, but once you get it it all works! I’ve made a JavaScript and iOS tutorial linked below for free!

Let me know if you need help. If you found it useful please like, comment, and subscribe!

Thanks guys. 

JavaScript - https://youtu.be/JzU_U3QEP7o

iOS - https://youtu.be/XvhQkXNv1_I
## [10][What does AWS EventBridge mean to ISV?](https://www.reddit.com/r/aws/comments/jhoa1s/what_does_aws_eventbridge_mean_to_isv/)
- url: https://www.reddit.com/r/aws/comments/jhoa1s/what_does_aws_eventbridge_mean_to_isv/
---
I read people saying AWS EventBridge is as revolutionary as Lambda (btw, I do agree lambda is revolutionary in cloud software development) [https://www.trek10.com/blog/amazon-eventbridge/](https://www.trek10.com/blog/amazon-eventbridge/)

Then, it begs a question whether that means some good opportunities for ISVs? For example, does that mean it's easier to develop a cloud service like Zapier or ITFFF?
## [11][AWS IAM Policy based on EC2 tag](https://www.reddit.com/r/aws/comments/jhnzru/aws_iam_policy_based_on_ec2_tag/)
- url: https://www.reddit.com/r/aws/comments/jhnzru/aws_iam_policy_based_on_ec2_tag/
---
I created an IAM Role with the following Permissions policy

    {
        "Version": "2012-10-17",
        "Statement": [
            {
                "Sid": "VisualEditor0",
                "Effect": "Allow",
                "Action": "ec2:DescribeInstances",
                "Resource": "*",
                "Condition": {
                    "StringEquals": {
                        "aws:ResourceTag/Project": "Phoenix"
                    }
                }        
            }
        ]
    }

* A 3rd party App uses the Role ARN associated with the IAM Policy to describe EC2 instances. 

* The objective is for the App to only be allowed describing EC2 instances tagged Project=Phoenix. 

* The issue is the App gets an error "AWS credentials have insufficient permission to perform action" so no EC2 are described at all. There are 2 instances with the matching tag.

* If I remove the Condition from the policy, DescribeInstances works but the App gets all EC2 from any specified AWS Region. I need to limit the list of EC2 returned to only instances tagged Project=Phoenix.

* Also tried tagging the IAM Role with Project=Phoenix as a test but same error.

Am I using &amp; configuring the IAM Role and Policy correctly for the objective? What do I need to adjust?
