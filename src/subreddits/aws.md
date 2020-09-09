# aws
## [1][Week of Sept 8th - What are you building this week on AWS?](https://www.reddit.com/r/aws/comments/iot7kv/week_of_sept_8th_what_are_you_building_this_week/)
- url: https://www.reddit.com/r/aws/comments/iot7kv/week_of_sept_8th_what_are_you_building_this_week/
---
Share your victories with the community
## [2][Is it possible to record the speaker output of an EC2 instance?](https://www.reddit.com/r/aws/comments/ip15mq/is_it_possible_to_record_the_speaker_output_of_an/)
- url: https://www.reddit.com/r/aws/comments/ip15mq/is_it_possible_to_record_the_speaker_output_of_an/
---
SOLVED! I commented which solution worked in the comments below. Thank you everyone!

I'm a highschooler trying out AWS for the first time, so forgive me if this is a dumb question.

How can I record the speaker output of an AWS instance? I'm using the Amazon Linux 2 AMI. I see that when I install pulseaudio and list the sinks, there's a "Dummy Speaker" (or something along the lines of that).

I’ll describe my project really quick:
Essentially, I’m making a “Zoom bot” (similar to a discord one) that can do transcriptions, and also act as a Chatbot within in meeting chat. For transcriptions, I’m trying to record the audio and send it to Amazon Transcribe. Other Zoom apps that offer transcription capability need you to have a paid zoom account for either cloud recording or custom live streaming.
## [3][Security September: Escaping CodeBuild - The compromise that wasn't – One Cloud Please](https://www.reddit.com/r/aws/comments/ip477h/security_september_escaping_codebuild_the/)
- url: https://onecloudplease.com/blog/security-september-escaping-codebuild
---

## [4][With the fires in the West, I'd like to be able to plan accordingly. Does anyone know where us-west-2 is, exactly (or even vaguely)?](https://www.reddit.com/r/aws/comments/ip4iqe/with_the_fires_in_the_west_id_like_to_be_able_to/)
- url: https://www.reddit.com/r/aws/comments/ip4iqe/with_the_fires_in_the_west_id_like_to_be_able_to/
---
Subject pretty much says it all, I just want to know if there is a fast moving fire heading toward my organization's stuff.  Does anyone know where the various AZs of us-west-2 are?
## [5][Self hosted Docker Registry with S3 storage - high ListBucket usage](https://www.reddit.com/r/aws/comments/ipf0rr/self_hosted_docker_registry_with_s3_storage_high/)
- url: https://www.reddit.com/r/aws/comments/ipf0rr/self_hosted_docker_registry_with_s3_storage_high/
---
Hey everybody,

I just inherited an account and performed few initial Cost Explorer checks looking for quick easy wins. I noticed there's a recurring $70 fee for S3 even though the account itself does not utilize S3 that much. 99% of the cost comes from EU-Requests-Tier1 and after downloading a detailed report of the S3 costs I was able to extract useful piece of info - there are around 230-250k ListBucket requests per day(a constant stream of 15-18k per hour), to a bucket serving as a Storage to a self hosted [Docker Registry](https://docs.docker.com/registry/) on ECS. 

Of course, there are various applications running on the account(Consul, Hazelcast, RabbitMQ, Grafana, Vault, ELK stack and few self-made Java/Kotlin apps on ECS) which may be overloading the Registry with their requests. 

I just wanted to perform a reddit sanity check that this is not some known issue with Registry itself. This may of course be some app overloading the registry by a constant stream of requests. I'd rather not dive into code of the [S3 driver](https://docs.docker.com/registry/storage-drivers/s3/) to look up when exactly the `ListBucket` is triggered and then somehow set up a hook to see which app constantly triggers the call, unless absolutely necessary.
## [6][Considering moving our entire infrastructure to AWS. Could use some tips/information.](https://www.reddit.com/r/aws/comments/ipevga/considering_moving_our_entire_infrastructure_to/)
- url: https://www.reddit.com/r/aws/comments/ipevga/considering_moving_our_entire_infrastructure_to/
---
Hi,

I am the sole sysadmin for a SaaS shop which at some client locations also provides some hardware. Currently our entire infrastructure is self managed in a private VMWare cluster in a single DC. The hardware we have running for this is entirely overkill as we maybe use 20% of it's capacity. Other than a few VM's I have never used anything public cloud related for production use. Yet I am considering moving all our stuff to AWS to easy overall sysadmin tasks and give devs a bit more control/flexibility. Maybe even reduce cost.

Before I spent tons of money, I have a lot of questions. Hopefully some of you can help out.

**Some key info:**

* I am eying Frankfurt or Paris as the Region.
* We need Site to Site IPSec VPN's to ~30 of customers to be able to manage their hardware components. Some big clients need double NAT IPSec.
* Our office is connected to the DC with a IPSec VPN.
* We want to host everything in at least 2 AZ's with auto failover or active-active. I guess we could use Route53 and/or AWS Load Balancers?
* Stuff we use/need:
 * MariaDB (Can be moved to Multi-AZ RDS which solves headaches of managing our own cluster)
 * Elasticsearch (Containerized but AWS provides this as a service)
 * We currently run some of our services on a 3 node Docker Swarm cluster with HaProxy to provide some HA between the Docker nodes.
 * Webservers (Containerized)
 * Python API's (Containerized)
 * RabbitMQ (Containerized, but can be moved to SQS)
 * VerneMQ MQTT Broker (Containerized but AWS IoT Core should be able to handle this)
 * InfluxDB (Containerized)
 * Prometheus (with a bunch of exporters) for monitoring (Containerized)
 * A few internal services like Confluence for documentation. All of these run in Docker.
 * Ansible to manage most of the VM's.
 * We are currently also using JumpCloud as a SAML/SSO solution and SSH Key manager so I can easily decide who gets to access what and can let some devs use SSH on some VM's.

I did notice some of their services are missing from the pricing calculator, like Fargate. Is this an error or is not everything on there?

**Questions:**

* Multi AZ: Does AWS take care of replicating EKS, ECS, SQS etc. to another AZ?
* Containers: I have never used Kubernetes, it looks daunting. Should I take the deep dive and run EKS or just use ECS on EC2 or Fargate. What does Amazon use for ECS, is this "regular" Docker Swarm?
* Should I stand up a hosted Cisco/Fortigate/Juniper/etc firewall inside AWS for all the VPN's? Using a VPC VPN for each customer is pretty damn expensive and IIRC AWS does not support double NAT. If so, any recommendations?
* Running ECS/EKS and/or normal EC2 instances: How would you setup failover to another AZ? Does AWS do this for you?
* How would I go about organizing some security/network segmentation inside AWS? Especially if combined with a virtual firewall for example.
* I guess we could use their Application Load Balancers with Elastic IP's to organize where certain traffic goes. How does this work over multiple AZ's and with containers?

Any recommendations or remarks after reading the above? :-)

Thanks!
## [7][Stream Kinesis to a React App](https://www.reddit.com/r/aws/comments/ipesvg/stream_kinesis_to_a_react_app/)
- url: https://www.reddit.com/r/aws/comments/ipesvg/stream_kinesis_to_a_react_app/
---
Hi All,

I have a data stream in Kinesis that I would like to expose and display in a simple React app in close to real time.

What is the best way to do this? AppSync?

I could save to a DB then have react poll every X seconds but that seems like a waste of API calls if not much data is changing.

&amp;#x200B;

Thanks for any help
## [8][A GitHub repo template for maintaining a multiple environment infrastructure with Terraform in AWS](https://www.reddit.com/r/aws/comments/ios8d6/a_github_repo_template_for_maintaining_a_multiple/)
- url: https://www.reddit.com/r/aws/comments/ios8d6/a_github_repo_template_for_maintaining_a_multiple/
---
[https://github.com/unfor19/terraform-multienv](https://github.com/unfor19/terraform-multienv) \- This template includes a CI/CD process, that applies the infrastructure in an AWS account.

[CI\/CD status](https://preview.redd.it/5hb4mf96ywl51.png?width=1258&amp;format=png&amp;auto=webp&amp;s=7128a30365d8a1e6346cf610a268d5cdd31d6038)
## [9][Large file transfer to/from S3 with custom domain name](https://www.reddit.com/r/aws/comments/ipbja4/large_file_transfer_tofrom_s3_with_custom_domain/)
- url: https://www.reddit.com/r/aws/comments/ipbja4/large_file_transfer_tofrom_s3_with_custom_domain/
---
Is there a way to make large file transfers to/from S3 using my own domain name? My first thought was to use API Gateway as a proxy but there is a 10MB payload limit.

I basically want to abstract the S3 bucket endpoint from the transfer but still keep it RESTful, if possible. 

Any thoughts? Thanks!
## [10][An intro to EventBridge](https://www.reddit.com/r/aws/comments/ioz4ff/an_intro_to_eventbridge/)
- url: https://medium.com/@tschoffelen/migrating-to-serverless-building-the-bridge-840ddc7cc9d9
---

## [11][cPanel AWS - Multiple Dedicated IPs](https://www.reddit.com/r/aws/comments/ipcu5w/cpanel_aws_multiple_dedicated_ips/)
- url: https://www.reddit.com/r/aws/comments/ipcu5w/cpanel_aws_multiple_dedicated_ips/
---
Hi all,

I am setting up cPanel and would like to host multiple websites on it, however I only have one elastic IP. When I create a new user on cPanel and select Dedicated IP it says there are no free IPs.

&amp;#x200B;

How do I add more than one Elastic IP and add them to the cPanel please?

&amp;#x200B;

Many Thanks
