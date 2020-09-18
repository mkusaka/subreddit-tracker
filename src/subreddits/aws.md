# aws
## [1][Congrats to r/aws on 140k subscribers!](https://www.reddit.com/r/aws/comments/iv4gen/congrats_to_raws_on_140k_subscribers/)
- url: https://www.reddit.com/r/aws/comments/iv4gen/congrats_to_raws_on_140k_subscribers/
---
Great to see a growing and vibrant community
## [2][Api Gateway now supports mTLS](https://www.reddit.com/r/aws/comments/iv4ex9/api_gateway_now_supports_mtls/)
- url: https://aws.amazon.com/blogs/compute/introducing-mutual-tls-authentication-for-amazon-api-gateway/
---

## [3][Yikes: AWS Aurora PostgreSQL versions vanish from the mega-cloud for days, leaving customers in the dark](https://www.reddit.com/r/aws/comments/iuojz9/yikes_aws_aurora_postgresql_versions_vanish_from/)
- url: https://www.reddit.com/r/aws/comments/iuojz9/yikes_aws_aurora_postgresql_versions_vanish_from/
---
Surprisingly, I saw this via /u/quinnypig's Twitter account. Has anybody here been affected by this? 

[AWS Aurora PostgreSQL versions vanish from the mega-cloud for days, leaving customers in the dark](https://www.theregister.com/2020/09/16/aws_aurora_postgresql_versions_disappeared/)
## [4][Features of a proper pipeline service - my improvement roadmap for AWS CodePipeline](https://www.reddit.com/r/aws/comments/iv5k1q/features_of_a_proper_pipeline_service_my/)
- url: https://www.sentiatechblog.com/features-of-a-proper-pipeline-service?utm_source=reddit&amp;utm_medium=social&amp;utm_campaign=proper_pipeline
---

## [5][Curated AWS SSM Scripts](https://www.reddit.com/r/aws/comments/iv5h7p/curated_aws_ssm_scripts/)
- url: https://rewind.io/blog/curated-aws-ssm-scripts
---

## [6][AmazonLinux2 package repos | S3 VPC endpoint sufficient or Internet access required? seems needed](https://www.reddit.com/r/aws/comments/iv2a12/amazonlinux2_package_repos_s3_vpc_endpoint/)
- url: https://www.reddit.com/r/aws/comments/iv2a12/amazonlinux2_package_repos_s3_vpc_endpoint/
---
Hi all,

I was under the impression that the Amazon Linux package repo's were hosted in Amazons [own s3 buckets](https://aws.amazon.com/amazon-linux-ami/faqs/) ...but that article is old and referencing the Amazon Linux 1 AMI, not 2. Im not sure why AWS wouldn't host a mirror in their own S3 buckets, that would seem to 'finish off' Systems Manager capabilities in not requiring instances to leave private link, but allowing package updates via repos in S3 via S3 endpoint. Seems like this would be a easy win for them to satisfy customers egress concerns.

[Amazon Linux AMI basics](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/amazon-linux-ami-basics.html#package-repository) AWS page shows that EC2 instances 'must have internet access'. To test I setup NACL and SG to only allow traffic to/from S3 VPC endpoint. I do see in the yum configs for Amazon Linux 2 that there are a few repo files in /etc/yum.repos.d/, but for brevity there is:

amzn2-core.repo which has \[amzn2-core\] enabled with mirror list:

http//amazonlinux.$awsregion.$awsdomain/$releasever/$product/$target/$basearch/mirror.list

If I run a yum update I can see traffic timing out to:

http//amazonlinux.us-east-1.amazonaws.com/2/core/2.0/x86\_64/7bd66e56187a1363571cafcbc92d6f3a57fa4ac69ed6e5859f3f5rd3b166c8/repodata/repomd.xml?instance\_id=i-*instanceId*&amp;region=us-west-2:

Which If i hit from an internet connected machine I receive the xml which shows its a public hosted mirror at Duke University: [http://linux.duke.edu/metadata/repo](http://linux.duke.edu/metadata/repo)

Now I could setup a proxy like Squid, host a mirror locally or in S3, or just allow my subnet out to the internet, but it seems like AWS hosting their own mirrors in their own S3 buckets wouldn't be that much farther to go, I mean they have their own distro flavor and a repo named 'amzn-\*'. Thoughts? Its 1:30 am and I may not be thinking clearly, I was working on setting up patching for AmznLinux2 instances today and ran into this since we dont allow outbound internet from this particular environment.
## [7][Availability of t4g instances in zone eu-est-1 (Ireland)](https://www.reddit.com/r/aws/comments/iv20gc/availability_of_t4g_instances_in_zone_euest1/)
- url: https://www.reddit.com/r/aws/comments/iv20gc/availability_of_t4g_instances_in_zone_euest1/
---
Following the announcement of new t4g instances I updated my autoscaler launch configuration to use it as *"*[*T4g instances*](https://aws.amazon.com/ec2/instance-types/t4/) *are available today \[14 Sept.\] in US East (N. Virginia, Ohio), US West (Oregon), Asia Pacific (Tokyo, Mumbai), Europe (Frankfurt, Ireland)."*

But when launching a new instance I got the message :

&gt;Launching a new EC2 instance. Status Reason: The requested configuration is currently not supported. Please check the documentation for supported configurations. Launching EC2 instance failed.

Is there something I missed ?
## [8][Network outage sa-east-1??](https://www.reddit.com/r/aws/comments/iusmks/network_outage_saeast1/)
- url: https://www.reddit.com/r/aws/comments/iusmks/network_outage_saeast1/
---
different services (EC2 &amp; RDS) are suffering to communicate here. São Paulo (sa-east-1).

Someone else with the same problem?
## [9][How do I check how much data I have analyzed?](https://www.reddit.com/r/aws/comments/iuyyvp/how_do_i_check_how_much_data_i_have_analyzed/)
- url: https://www.reddit.com/r/aws/comments/iuyyvp/how_do_i_check_how_much_data_i_have_analyzed/
---
Well, the title's pretty vague, so I'll try to explain my situation here. I'm attempting to use Rekognition to do some simple facial emotion recognition. The problem is that I do not want to spend any money. Good news is that Amazon allows for up to 5,000 images to be processed a month for 12 months for free on Rekognition. Bad news is that I can't figure out how many photos I have analyzed so far. I can't seem to find a dashboard, or webpage, on Amazon Web Services that states how many images or how long I have (in terms of images left to analyze) until I'm out of free images for the month. Would appreciate any help, thanks!

&amp;#x200B;

(also, this is my first time using AWS, its frickin awesome)
## [10][App architecture](https://www.reddit.com/r/aws/comments/iuyqkh/app_architecture/)
- url: https://www.reddit.com/r/aws/comments/iuyqkh/app_architecture/
---
Hey guys, I’m new to the dev side of AWS. Long time on AWS infrastructure, have a idea for an app but unsure what services to go with. 

The general flow is user logs in, fills out a form, based on their choices it will generate a pdf with dynamic content. The user could then view their content and share it. 

I’ve heard good things about Amplify, but am not sure if it would work here. As I’ve read it is for static websites. 

I’m thinking - Some combination of S3, api gateway and lambda and Cognito. 

Any advice??
