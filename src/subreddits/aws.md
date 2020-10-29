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
## [2][Did they JUST make a permanent change to the console UI? I was uploading items to my S3 bucket and I can't find a way to get back the view that it was before. Is there a way to revert back or did they finally get rid of that option? Not seeing any options to do so...](https://www.reddit.com/r/aws/comments/jk9pyi/did_they_just_make_a_permanent_change_to_the/)
- url: https://www.reddit.com/r/aws/comments/jk9pyi/did_they_just_make_a_permanent_change_to_the/
---
&amp;#x200B;

e.g. back when the buckets looked like this [https://miro.medium.com/max/1050/1\*xcOmWtHEhukc-VbgiLw8cg.png](https://miro.medium.com/max/1050/1*xcOmWtHEhukc-VbgiLw8cg.png)

I was looking for a quick way to verify MFA delete was enabled in the console. I know I can do this with the CLI.

Thanks!
## [3][100G networking in AWS, a network performance deep dive](https://www.reddit.com/r/aws/comments/jjsxjn/100g_networking_in_aws_a_network_performance_deep/)
- url: https://toonk.io/aws-network-performance-deep-dive/
---

## [4][AWS Nitro Enclaves – Isolated EC2 Environments to Process Confidential Data](https://www.reddit.com/r/aws/comments/jjxvjs/aws_nitro_enclaves_isolated_ec2_environments_to/)
- url: https://aws.amazon.com/blogs/aws/aws-nitro-enclaves-isolated-ec2-environments-to-process-confidential-data/
---

## [5][Inappropriate "update saved password" window for 2fa](https://www.reddit.com/r/aws/comments/jk64x5/inappropriate_update_saved_password_window_for_2fa/)
- url: https://www.reddit.com/r/aws/comments/jk64x5/inappropriate_update_saved_password_window_for_2fa/
---
Hello. My organisation has 2-factor auth turned on, so I have to type in the numbers from my smartphone every time I need to log into AWS consolse. The problem is that whenever I type those numbers and press the submit button, Chrome asks about changing my current AWS password to those numbers, which is mildly annoying. How do I prevent this from happening?
## [6][How to execute python code from a web app](https://www.reddit.com/r/aws/comments/jk9e7v/how_to_execute_python_code_from_a_web_app/)
- url: https://www.reddit.com/r/aws/comments/jk9e7v/how_to_execute_python_code_from_a_web_app/
---
Hello,

So I have a react web app that makes students submit their code. What I am trying to do is run this python code and check it with test units and return to the students the output of the code. So it's like an online IDE to smth. They don't need to install python locally to run. The web app is deployed on Google Firebase, is there a way I can achieve this functionality with AWS?
## [7][Does ElastiCache support Redis version 6+?](https://www.reddit.com/r/aws/comments/jk58of/does_elasticache_support_redis_version_6/)
- url: https://www.reddit.com/r/aws/comments/jk58of/does_elasticache_support_redis_version_6/
---
There are some version specific features to Redis version 6 that would make my dev experience much easier and faster but I also want to be able to use ElastiCache to enjoy its benefits...

Does ElastiCache support Redis version 6?
## [8][Sessions, tokens, cognito and jam stacks!](https://www.reddit.com/r/aws/comments/jjyvw3/sessions_tokens_cognito_and_jam_stacks/)
- url: https://www.reddit.com/r/aws/comments/jjyvw3/sessions_tokens_cognito_and_jam_stacks/
---
I’m big on the serverless movement and lately I’ve been trying to migrate my architecture to a jam stack model. The sticking point for me atm is security.

I’ve been playing around with oauth / pkce to create a client side only oauth flow using cognito hosted UI.

There is obviously a risk trade off that you have to decide upon.

In a trusted environment this setup works well.

I thought I was closing in on an acceptable balance, persisting access token only to session storage and if the user indicates ‘remember me’ then persist access token AND refresh token to local storage - which can create a long lived login.

Say access token = 1 hour and refresh = 30 days or whatever.

We know the trade off, access token can never really be invalidated, (and tbf I’ve been struggling to verify that cognito logout (at least to invalidate the refresh) actually works at all!)

But of course it then occurred to me that if the user opted not to be remembered and we discard (or even never issue) the refresh token, then there is no option but for the session to terminate after 1 hour (unlike cookie / server model where we can keep touching the expiry time stamp).

In this mind it seems the refresh token is the more sensitive asset here?
 
I’m really keen to understand what security model the aws console implements but I have no idea where I can find this information out.

What I do know is that it always forces a logout after about 12 hours or so, so perhaps it does work like an access token with 12 hour expiry and no refresh token?

There are a lot of considerations, I’m banging my head trying to reinvent what many others must have already settled on as best practices.

I have read the many articles on why jwt tokens should not be used for session tokens, but I’m all for jamstack, oauth and jwt,  it the idea that all api requests have to be routed through some session gateway doesn’t ring true with me.

Do I need to implement a session api endpoint that can ‘touch’ the cookie expiry, and let this api endpoint provide me with the api jwt tokens and store the refresh toke here - server side?

I would be really interested to learn other peoples thoughts and experiences in this domain.

Cheers!
## [9][Kubernetes Resources and Autoscaling — From Basics to Greatness](https://www.reddit.com/r/aws/comments/jjl7wt/kubernetes_resources_and_autoscaling_from_basics/)
- url: https://www.reddit.com/r/aws/comments/jjl7wt/kubernetes_resources_and_autoscaling_from_basics/
---
[In this blog](https://itnext.io/kubernetes-resources-and-autoscaling-from-basics-to-greatness-7cae17fbf27b?source=friends_link&amp;sk=ea2b30df277edaf345dc4868a63ef29b), we will try to describe the building blocks that are used by Kubernetes to implement autoscaling. We will see how Kubernetes collects and use metrics to autoscale pods and nodes. We will use AWS EKS to demonstrate the patterns and best practices used to implement autoscaling both at the pod and node levels.
## [10][Is there any way to get an url from createPresignedPost function like createPresignedUrl does?](https://www.reddit.com/r/aws/comments/jk4f95/is_there_any_way_to_get_an_url_from/)
- url: https://www.reddit.com/r/aws/comments/jk4f95/is_there_any_way_to_get_an_url_from/
---
If not, is there any tutorial for that? Or is there any way I can restrict the file size with createPresignedUrl?
It's pretty hard to find useful resources for those .
## [11][buff/cache is high, can i flush?](https://www.reddit.com/r/aws/comments/jk3buc/buffcache_is_high_can_i_flush/)
- url: https://www.reddit.com/r/aws/comments/jk3buc/buffcache_is_high_can_i_flush/
---
Hi,  


my server buff/cache is high, but checking from [linuxatemyram site](https://www.linuxatemyram.com) says its normal.  
as this speed up the system, but does not necessarily take away ram.  


my question is, would it be safe to run this to an ec2 server (AMI2), to clear cache memory?  
`echo 3 | sudo tee /proc/sys/vm/drop_caches`  
this is what the linuxatemyram recommend as a workaround.

&amp;#x200B;

thanks
