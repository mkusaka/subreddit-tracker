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
## [2][Week of Sept 26th - AWS Open Discussion](https://www.reddit.com/r/aws/comments/jifk9h/week_of_sept_26th_aws_open_discussion/)
- url: https://www.reddit.com/r/aws/comments/jifk9h/week_of_sept_26th_aws_open_discussion/
---
Doing this for a few reasons. We know some folks are hesitant to create a new post for a small question they may have. Well, this is the place to ask, and discuss. At the same time, with a growing community we find ourselves having to limit the posts that may be off-topic to the primary purpose of the sub.
## [3][aurora multi master in production anyone?](https://www.reddit.com/r/aws/comments/jizg99/aurora_multi_master_in_production_anyone/)
- url: https://www.reddit.com/r/aws/comments/jizg99/aurora_multi_master_in_production_anyone/
---
Hi, 

I am evaluating a migration on aurora multimaster with a database-based sharding.

I'd like to hear some real-life experience from someone which is using this setup.

Any tips or gotchas to know before crossing that bridge?
## [4][How to create a role for clients so that they can access my S3 bucket and simply upload the required folder?](https://www.reddit.com/r/aws/comments/jizl6r/how_to_create_a_role_for_clients_so_that_they_can/)
- url: https://www.reddit.com/r/aws/comments/jizl6r/how_to_create_a_role_for_clients_so_that_they_can/
---
Hi, I am a complete AWS noob, so please excuse my lack of knowledge. I have only started working on AWS3 days ago.

I have created an S3 Bucket so that my clients can upload their folders to it. What I want is to use a simple third-party client for Amazon S3 so that they can very easily go to my bucket and just drop the folder they wanted to upload and leave, since I believe, accessing the AWS console will confuse them. As I understand, I will have to create a role for them. Can anyone please guide me as to how to create a role for my clients with write-only access? Also what client should I install on the clients machine?
## [5][The annual “Last Week in AWS” charity t-shirt drive has begun, featuring Route 53](https://www.reddit.com/r/aws/comments/jijfqo/the_annual_last_week_in_aws_charity_tshirt_drive/)
- url: https://www.lastweekinaws.com/2020-charity-t-shirt/
---

## [6][Best AWS service for implamenting CI/CD service](https://www.reddit.com/r/aws/comments/jiuf3n/best_aws_service_for_implamenting_cicd_service/)
- url: https://www.reddit.com/r/aws/comments/jiuf3n/best_aws_service_for_implamenting_cicd_service/
---
I'me creating a \`flutter build\` feature for my service.

Here is the question, Since flutter build requires to install flutter, and then build it, I'm not sure Lambda will be the best choice, It might be overpriced. What other option do i have for performing semi-havy operation per request?

&amp;#x200B;

It usually takes 5 minutes to install, and build web on local machine.

&amp;#x200B;

&amp;#x200B;

input : source file

process: install flutter, build dist with source file

output: dist as a zip file
## [7][How do large enterprises manage security in AWS?](https://www.reddit.com/r/aws/comments/jilcqf/how_do_large_enterprises_manage_security_in_aws/)
- url: https://www.reddit.com/r/aws/comments/jilcqf/how_do_large_enterprises_manage_security_in_aws/
---
Imagine you are working in a security role in a very large enterprise with many AWS roles/accounts and they want a central point to manage AWS security and look after API access limits and similar things across the enterprise. Is there any guidance or best practice available for this sort of thing?

Thanks!
## [8][What AWS Technologies would I want to use?](https://www.reddit.com/r/aws/comments/jizn9x/what_aws_technologies_would_i_want_to_use/)
- url: https://www.reddit.com/r/aws/comments/jizn9x/what_aws_technologies_would_i_want_to_use/
---
Hi Everyone,

Formerly a firebase user looking to start up with AWS for a new project.

The project is simply going to be a web app (possibly server side rendered for speed), sitting on a node.js backend with a MonogoDB server. I also have a java 'app' which on my PC I've been using to run backend data collection tasks. 

In firebase this would be using the firebase cloud store database, and firebase functions.

Is there anything like this on AWS? I don't want to think a ton about scaling for (hopefully) a lot of pageviews. But I'm also willing to do something more custom if its cheaper
## [9][How secure are pods running on EKS Fargate?](https://www.reddit.com/r/aws/comments/jizduy/how_secure_are_pods_running_on_eks_fargate/)
- url: https://www.reddit.com/r/aws/comments/jizduy/how_secure_are_pods_running_on_eks_fargate/
---
I'm currently evaluating running workloads on EKS Fargate and found this in the [documentation](https://docs.aws.amazon.com/eks/latest/userguide/fargate.html).

&gt;Fargate runs each pod in a VM-isolated environment without sharing resources                                              with other pods. However, because Kubernetes is a single-tenant orchestrator,                                              Fargate cannot guarantee pod-level security isolation. You should run sensitive                                              workloads or untrusted workloads that need complete security isolation using                                              separate Amazon EKS clusters.

What does this mean exactly? My interpretation is that Kubernetes doesn't have a concept of multi-tenancy and since you're sharing infrastructure with other users there's a chance they can access your data? If so, that would pretty heavily limit its use cases from my perspective.

It's a bit vague, it kinda sounds like they're saying "there's isolation between pods so it should be secure" followed by a bit of hand waving and the disclaimer "except for when it's not..."

Is my understanding correct? How great is the risk?
## [10][Should I create VPCs with IPv6 at this point?](https://www.reddit.com/r/aws/comments/jiqus8/should_i_create_vpcs_with_ipv6_at_this_point/)
- url: https://www.reddit.com/r/aws/comments/jiqus8/should_i_create_vpcs_with_ipv6_at_this_point/
---
Understand that there is still room with IPv4, but would it be a long term play to go with IPv6? Is there any tradeoffs with this?
## [11][Restrict access to your internal websites on AWS with BeyondCorp (defined with Terraform)](https://www.reddit.com/r/aws/comments/jiqmyp/restrict_access_to_your_internal_websites_on_aws/)
- url: https://transcend.io/blog/restrict-access-to-internal-websites-with-beyondcorp
---

## [12][ALB design w/ NLB? Use ALB to handle requests, and NLB to forward data for logging? (Splunk)](https://www.reddit.com/r/aws/comments/jipp9m/alb_design_w_nlb_use_alb_to_handle_requests_and/)
- url: https://www.reddit.com/r/aws/comments/jipp9m/alb_design_w_nlb_use_alb_to_handle_requests_and/
---
Hey all,   


I am setting up Splunk Enterprise in AWS. I want to set up ALB for the front end, and an NLB for the backend to send logs to the indexer when necessary for any systems that are SaaS.  


Is it possible to route all traffic through the ALB, but should it hit a certain port or rule it directs it to the NLB and begins logging TCP/TLS? I have been putzing around all day trying to get this set up, the only portion I seem to have up and running is ALB w/ Route 53. Just want to be able to point Domain w/ proper ports to Splunk and begin logging (and proper SG of course)  


Thanks!
