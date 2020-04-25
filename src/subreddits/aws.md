# aws
## [1][Is it allowed in ToS to discuss enterprise support cases online?](https://www.reddit.com/r/aws/comments/g7qbd4/is_it_allowed_in_tos_to_discuss_enterprise/)
- url: https://www.reddit.com/r/aws/comments/g7qbd4/is_it_allowed_in_tos_to_discuss_enterprise/
---
I find myself asking a lot of questions to support for stuff that I can't find already-existing info on, and so I probably open a support case at least once every two weeks or so.

The responses are always amazing, and I want to share the info I get online.  In a recent example, I found an flaw (not a security flaw) in one of the AWS provided default IAM policies, and when I wrote to support, they were like, "indeed, we tested it too, the policy is mistaken, we've escalated this to the IAM team, and to work around it here are some ways you can fill in the gaps." It seems like a shame that info is just sitting in my inbox and not on a forum somewhere.

Can I copy paste those responses somewhere?  Do I have to redact the agents name?  Or do I have to reword the issue and solution in my own words?
## [2][The Extended AWS Security Ramp-Up Guide](https://www.reddit.com/r/aws/comments/g7orbm/the_extended_aws_security_rampup_guide/)
- url: https://research.nccgroup.com/2020/04/24/the-extended-aws-security-ramp-up-guide/
---

## [3][Making real time dashboard (website) using API gateway, s3 and lambda?](https://www.reddit.com/r/aws/comments/g7spzt/making_real_time_dashboard_website_using_api/)
- url: https://www.reddit.com/r/aws/comments/g7spzt/making_real_time_dashboard_website_using_api/
---
I am currently making a website that monitors my raspberry pi data (which is stored in dynamoDB) and show it in real time. My aws design goal is to create a serverless realtime web app. So from what I have heard, I can add my static site (I am using vueJS) in S3 and then use the help of API gateway and lambda to get the data. The problem is I am not familiar with websockets. So please forgive me if I say something stupid; So my biggest question is the client side code. Once I have created a websocket API gateway, what should I add in my client side to make it work? I know I can use rest api instead, but since I want the site to be in realtime, I would need to call the API every few second (which is not cost effective).

If you have any advice regarding my project or implementation please feel free to give suggestions.

Thank you!
## [4][We are the AWS AI / ML Team - Ask the Experts - May 18th @ 9AM PT / 12PM ET / 4PM GMT!](https://www.reddit.com/r/aws/comments/g7gjgg/we_are_the_aws_ai_ml_team_ask_the_experts_may/)
- url: https://www.reddit.com/r/aws/comments/g7gjgg/we_are_the_aws_ai_ml_team_ask_the_experts_may/
---
Hey [r/aws](https://www.reddit.com/r/aws)! u/AmazonWebServices here.

The AWS AI/ML team will be hosting an Ask the Experts session here **in this thread** to answer any questions you may have about building and training machine learning models with [Amazon SageMaker](https://aws.amazon.com/sagemaker/), as well as any questions you might have about machine learning in general.

Already have questions? Post them below and we'll answer them starting at 9AM PT on May 18, 2020!
## [5][Architecture deep-dive: How a news aggregator collects and serves up news articles](https://www.reddit.com/r/aws/comments/g7smug/architecture_deepdive_how_a_news_aggregator/)
- url: https://www.reddit.com/r/aws/comments/g7smug/architecture_deepdive_how_a_news_aggregator/
---
Hey guys, 

For today's AWS Deep Dive I've interviewed the creators of a news aggregator and conducted an architecture review of their design.  

[https://www.youtube.com/watch?v=XaO8BBgBRH0](https://www.youtube.com/watch?v=XaO8BBgBRH0)

Thanks,

Vladimir

P.S. Subscribe if you find it useful and let me know what else you'd like me to make
## [6][AWS Amplify - How do I allow federated identity users to only access data that belongs to them via GraphQL queries?](https://www.reddit.com/r/aws/comments/g7nu6y/aws_amplify_how_do_i_allow_federated_identity/)
- url: https://www.reddit.com/r/aws/comments/g7nu6y/aws_amplify_how_do_i_allow_federated_identity/
---
Hi folks, I've got a mobile app where users sign up/sign in with 'Sign In With Apple' in which a federated identity is created upon sign in. I need to store and fetch data that is only accessible by the user.

How can I make it so users can only access their own data when the app makes GraphQL queries? Which value can I extract from their federated identities that is reliable? Is there a better way to do this?

Thanks in advance, I've searched quite hard for this but I can't find a clear cut answer.
## [7][Possible to point a domain to both an S3 static website and a ELB fronted service?](https://www.reddit.com/r/aws/comments/g7th9c/possible_to_point_a_domain_to_both_an_s3_static/)
- url: https://www.reddit.com/r/aws/comments/g7th9c/possible_to_point_a_domain_to_both_an_s3_static/
---
So how would this scenario be pulled off?  I've got a static website hosted at www.domain.com.  I've got a service running behind an ELB that I want accessed from www.domain.com/api.  Is this possible? Or is the only way to pull this off is by use of another subdomain such as api.domain.com?
## [8][How to secure API gateway with PingFederate?](https://www.reddit.com/r/aws/comments/g7suut/how_to_secure_api_gateway_with_pingfederate/)
- url: https://www.reddit.com/r/aws/comments/g7suut/how_to_secure_api_gateway_with_pingfederate/
---
Has anyone integrated REST APIs deployed to AWS API Gateway with PingFederate to implement OAuth 2.0 protocol?
## [9][Advice with my S3 strategy](https://www.reddit.com/r/aws/comments/g7o36k/advice_with_my_s3_strategy/)
- url: https://www.reddit.com/r/aws/comments/g7o36k/advice_with_my_s3_strategy/
---
This is my first experience with S3 so bear with me.

I'm looking to use S3 to store all the user assets for an app I'm building.  The assets would mostly just include videos and photos, and maybe user avatars.

Since the server I'm hosting the back end on has a slow upstream connection, I was hoping I could offload the work to the client.

So my idea is

1. Client hits my API with a request to upload some content
2. My server would validate the user, and validate they are a member of the group they are trying to upload too. Then return a temporary signed POST url
3. Client would then upload to that url
4. Once they complete the upload, they hit my API again so I can record the asset in my database. 

Then to fetch the data they would:

1. Hit my API for a download link
2. Server verifies they are authenticated, and are a member of the group
3. Generate a signed temporary download link, then redirect the client to the URL

Does that make sense?

I'm assuming I can't skip a step and just pipe the request straight to AWS through my server right? As that would still require my severs bandwidth?

In case it matters, the back end is written in Kotlin using the ktor framework, and I'm using the Java SDK 2
## [10][Cannot ssh to an ecs ec2 instance](https://www.reddit.com/r/aws/comments/g7qq3c/cannot_ssh_to_an_ecs_ec2_instance/)
- url: https://www.reddit.com/r/aws/comments/g7qq3c/cannot_ssh_to_an_ecs_ec2_instance/
---
I cannot ssh into my ecs ec2 instance even there's a ssh on the ssg inbound rules.

SSH TCP 22 [0.0.0.0/0](https://0.0.0.0/0)

SSH TCP 22 ::/0

Does anyone able to fix this? I can ssh to an ec2 instance not created through ecs though.
