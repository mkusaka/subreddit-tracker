# aws
## [1][EC2 inside a private subnet and S3](https://www.reddit.com/r/aws/comments/gkrtz0/ec2_inside_a_private_subnet_and_s3/)
- url: https://www.reddit.com/r/aws/comments/gkrtz0/ec2_inside_a_private_subnet_and_s3/
---
How is it possible for an ec2 inside a private subnet to use boto3 in order to save or load objects on S3? 

Thanks
## [2][a1.metal experience](https://www.reddit.com/r/aws/comments/gks1tv/a1metal_experience/)
- url: https://www.reddit.com/r/aws/comments/gks1tv/a1metal_experience/
---
Hi,
I am currently on AWS a1.metal and trying to see some SPE events,
given the core is a Neoverse one.

$ perf record -e arm_spe/ts_enable=1,pa_enable=1/ dd
event syntax error: 'arm_spe/ts_enable=1,pa_enable=1/'
                     \___ Cannot find PMU `arm_spe'. Missing kernel support?
Run 'perf list' for a list of valid events

 Usage: perf record [&lt;options&gt;] [&lt;command&gt;]
    or: perf record [&lt;options&gt;] -- &lt;command&gt; [&lt;options&gt;]

    -e, --event &lt;event&gt;   event selector. use 'perf list' to list available events

I'm using Ubuntu 20.4 LTS server, so I think the support is included, no?
## [3][Premium Support, quarantine feelings and the importance of saying thank you](https://www.reddit.com/r/aws/comments/gkazht/premium_support_quarantine_feelings_and_the/)
- url: https://www.reddit.com/r/aws/comments/gkazht/premium_support_quarantine_feelings_and_the/
---
Hello guys,

Probably a few of you could find this post stupid, but this is something that I had on my mind the whole week and I would like to share it.

I'm part of Premium Support on AWS and the other day I saw a post here about how awesome we are, well, not al of us but most, you know.

Maybe is the pandemic that is making us more emotional in general but this week I received 3 messages from customers saying how great my answer and my help was and I can't explain how happy that made me feel.

I know that is my job and I shouldn't expect this because basically there is someone paying me for doing an activity. Fair, capitalism is that and is completely reasonable. But as you may be already aware, the level of quality and customer obsession that Amazon expect from us is too high, and we can end up dealing with a lot of stress trying to provide the best solutions for other people. 

At the end sometimes, we are going out of scope, reviewing things about we didn't have any idea till 10 minutes ago, googling errors on behalf others and reading documentation to try to explain the same things as simple as possible. 

In my case, I was so afraid at the beginning and I suffered a lot the pressure of not having the same level of knowledge when I just started than a \~3 years user of a service.

Summarizing, because I believe I wrote more than I expected and the point of this was only say that if you consider that the answer really made your work easier and really helped you, if you have a minute please let us know. For me at least is very rewarding when I feel that I'm not just doing my job because someone else is paying me for doing it. I like to know that the users that I'm helping are not just numbers and they really appreciate the work.

If you read this all the way, thank you.

PS: Sorry if there are mistakes on my English, I'm not a native speaker.

—————————————

Edited: I would like to say thank you to everyone that took a few minutes to answer this post.  

I wasn’t expecting at all such a nice impact and good feedback, I only wanted to write this just to share a few feelings.

Makes me very happy read all your nice comments. 

Thank YOU for the support guys, is very rewarding and means a lot ⭐️⭐️⭐️⭐️⭐️
## [4][Central ingress point for all traffic in a mostly serverless environment with multiple accounts](https://www.reddit.com/r/aws/comments/gkrdk3/central_ingress_point_for_all_traffic_in_a_mostly/)
- url: https://www.reddit.com/r/aws/comments/gkrdk3/central_ingress_point_for_all_traffic_in_a_mostly/
---
In traditional environments, it is a common practice to put a perimeter security services around the whole environment. Is it possible to do something similar in a multi account serverless environment ?

If we use mostly VPC based services, I think I can put one VPC as the ingress point from the internet, and put my thirdparty firewalls and other perimeter security services, and then route them to multiple VPCs in multiple accounts using TGWs/Peering. But how would someone does this in a centralized manner for serverless applications ? Mostly lambdas, API GW that does not sit in VPCs ?

I think that even if it is possible to do it, it is not a scalable approach, and puts a lot of limitations to the developers. So, do you think that in this scenario, a centralized approach or decentralized approach will be better ? By decentralization, I mean that the public facing endpoints can sit in any account, but they must adhere certain policies that are enforced automatically by the central IT team - eg, all API GWs must have a WAF that sits in front. But then what if I need to inspect the packets that are coming from outside ? 

Sorry, so many questions, I am just brainstorming on what would be a better approach for standardization of a messy environment.
## [5][difference between cloudwatch and cloudtrail?](https://www.reddit.com/r/aws/comments/gksla8/difference_between_cloudwatch_and_cloudtrail/)
- url: https://www.reddit.com/r/aws/comments/gksla8/difference_between_cloudwatch_and_cloudtrail/
---
i noticed you can send api gateway logs to both cloudtrail and cloudwatch. is there any difference in the logs sent to either service?
## [6][Db migration in master slave db scenario](https://www.reddit.com/r/aws/comments/gku2t7/db_migration_in_master_slave_db_scenario/)
- url: https://www.reddit.com/r/aws/comments/gku2t7/db_migration_in_master_slave_db_scenario/
---
 I am planning to have a AWS aurora deployement in master slave setup . The AWS infra for aurora db  is setup from azure devops pipelines. Now what strategy should I follow for schema creation or update on these db instances. Will db migration that happen as part of the backend application startup work in these master slave scenario . My backend app is a .net api app running as fargate services .
## [7][m6g](https://www.reddit.com/r/aws/comments/gkraal/m6g/)
- url: https://www.reddit.com/r/aws/comments/gkraal/m6g/
---
Is it possible for a Free Tier user launch a bare metal m6g instance? I am still not sure if it is okay or not.
## [8][CloudFront: Is it not possible to route different CNAMEs to different paths within the same distribution?](https://www.reddit.com/r/aws/comments/gkpdjj/cloudfront_is_it_not_possible_to_route_different/)
- url: https://www.reddit.com/r/aws/comments/gkpdjj/cloudfront_is_it_not_possible_to_route_different/
---
Let's say I have a single distribution with S3 bucket as origin with folders such as /env1/files... and /env2/files....   Then I have [env1.somedomain.com](https://env1.somedomain.com) and [env2.somedomain.com](https://env2.somedomain.com) as the CNAMEs in the distribution.   Why can't I specify that env1. routes to S3/env1 and env2. routes to S3/env2?  Seems that you can only specify behavior based on path, not CNAME.
## [9][Is there any reason not to have awscli on an instance?](https://www.reddit.com/r/aws/comments/gkh9lw/is_there_any_reason_not_to_have_awscli_on_an/)
- url: https://www.reddit.com/r/aws/comments/gkh9lw/is_there_any_reason_not_to_have_awscli_on_an/
---
I'm setting up an ECS cluster using the AWS ECS optimized AMI, and I noticed awscli isn't on there.

Is there a particular reason?  Our standard bootstrap involves copying some stuff out of S3.
## [10][Diagrams as code (Python) with AWS icon support](https://www.reddit.com/r/aws/comments/gjxil2/diagrams_as_code_python_with_aws_icon_support/)
- url: https://diagrams.mingrammer.com/docs/getting-started/examples
---

