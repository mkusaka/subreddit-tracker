# aws
## [1][Week of Aug 24th - What are you building this week on AWS?](https://www.reddit.com/r/aws/comments/ifoebu/week_of_aug_24th_what_are_you_building_this_week/)
- url: https://www.reddit.com/r/aws/comments/ifoebu/week_of_aug_24th_what_are_you_building_this_week/
---
Share what you are creating on AWS
## [2][How are you guys saving money on AWS?](https://www.reddit.com/r/aws/comments/ifb1rg/how_are_you_guys_saving_money_on_aws/)
- url: https://www.reddit.com/r/aws/comments/ifb1rg/how_are_you_guys_saving_money_on_aws/
---
Hey guys,  


Times are tough and I am looking for ways to save money on AWS and maybe help somebody else seeing this post. What are some recent ways that you have been able to save a little extra money? Please provide the obvious suggestions too, as they may not be so obvious to me or somebody else.
## [3][AWS SSM Agent for Kubernetes](https://www.reddit.com/r/aws/comments/ifmlkv/aws_ssm_agent_for_kubernetes/)
- url: https://www.reddit.com/r/aws/comments/ifmlkv/aws_ssm_agent_for_kubernetes/
---
Access Kubernetes worker node with AWS SSM agent running as a DaemonSet. Install the AWS SSM Agent only when you need to access the K8s node, it's recommended to use it with [IAM Roles for Service Accounts](https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html). 

I've created the auto-updated (stay up today with the latest SSM Agent) SSM Agent Docker image [https://github.com/alexei-led/kube-ssm-agent](https://t.co/VNVYcawwno?amp=1) and a short setup guide.
## [4][Announcing a new open source tool: eks-auth-sync](https://www.reddit.com/r/aws/comments/ifl7hf/announcing_a_new_open_source_tool_eksauthsync/)
- url: https://polarsquad.com/blog/announcing-a-new-open-source-tool-eks-auth-sync
---

## [5][Lint your IAM policies the browser](https://www.reddit.com/r/aws/comments/iffk4e/lint_your_iam_policies_the_browser/)
- url: https://parliament.summitroute.com/
---

## [6][Extend an Existing on on-prem AD to EC2](https://www.reddit.com/r/aws/comments/ifnob6/extend_an_existing_on_onprem_ad_to_ec2/)
- url: https://www.reddit.com/r/aws/comments/ifnob6/extend_an_existing_on_onprem_ad_to_ec2/
---
Hello. 

I have to extend an existing on-prem AD to EC2 for DR/Backup purpose. I do not want to use AWS Managed AD but want the AD to be deployed on IaaS and us any native process for the replication. 

&amp;#x200B;

Can someone guide to what the process will be for this? Any article will do.
## [7][Is there ever a reason *not* to use a VPC endpoint if one is available?](https://www.reddit.com/r/aws/comments/ifk4js/is_there_ever_a_reason_not_to_use_a_vpc_endpoint/)
- url: https://www.reddit.com/r/aws/comments/ifk4js/is_there_ever_a_reason_not_to_use_a_vpc_endpoint/
---
I'm curious about VPC endpoints, and why there is the need to set them up manually for each VPC and service, and configure routing to use them when their use would generally be considered best practice. 

Why aren't they automatically set up by default for AWS services? Is it a matter of being purely another layer of least-privilege onion skinning or is there another reason?
## [8][How do we deal with awseducate for s3 rotating Access Keys &amp; Ids?](https://www.reddit.com/r/aws/comments/ifjlwv/how_do_we_deal_with_awseducate_for_s3_rotating/)
- url: https://www.reddit.com/r/aws/comments/ifjlwv/how_do_we_deal_with_awseducate_for_s3_rotating/
---
So I have an IAM role which is supervised by Vocareum and we cannot create access key or ids for accessing **S3FullAmazonAccess**, but how do I do this? Should I opt out of this or there is another way that I'm unaware of? I'm using Django and boto as SDK and I followed step by step from a Corey Schafer's Django tutorial. I also contacted AWS, but they don't say anything either.  

&amp;#x200B;

Error when accessing a picture:

`&lt;Error&gt;`

`&lt;Code&gt;InvalidAccessKeyId&lt;/Code&gt;`  
`&lt;Message&gt;The AWS Access Key Id you provided does not exist in our records.&lt;/Message&gt;`  
`&lt;AWSAccessKeyId&gt;ASMYACCESSKEYID&lt;/AWSAccessKeyId&gt;`  
`&lt;RequestId&gt;06E90B1C351AD95D&lt;/RequestId&gt;`  
`&lt;HostId&gt;SuL8rq6akfmQVsi3etxLR4OE/GI0WepXgOPWCLLFcESWVOddqV8dk4lhBJpa5fuk5rRy2ckaf1Y=&lt;/HostId&gt;`

`&lt;/Error&gt;`
## [9][How does Fargate Spot work exactly?](https://www.reddit.com/r/aws/comments/if6436/how_does_fargate_spot_work_exactly/)
- url: https://www.reddit.com/r/aws/comments/if6436/how_does_fargate_spot_work_exactly/
---
Hello all,

This should be a fairly easy question to answer.

I've been toying with a new project that uses ECS, running 4 containers on Fargate using Fargate spot as a capacity provider. They have been running for about a week now and I noticed that they have not gone down (or stopped) once. I expected them to run like EC2 spot instances, where after sometime they will go down and a new one will spin up. 

Is this not how Fargate spot works?
## [10][Is it still possible to take aws exams?](https://www.reddit.com/r/aws/comments/iflx12/is_it_still_possible_to_take_aws_exams/)
- url: https://www.reddit.com/r/aws/comments/iflx12/is_it_still_possible_to_take_aws_exams/
---
I know in-person test centers have been shut down, but aws's website won't allow me too register for a test online either. Is anyone else having that problem?
## [11][my AWS ssh is not connecting to instance .debug log in below text.i am not getting the terminal of Ubuntu instance](https://www.reddit.com/r/aws/comments/ifltzk/my_aws_ssh_is_not_connecting_to_instance_debug/)
- url: https://www.reddit.com/r/aws/comments/ifltzk/my_aws_ssh_is_not_connecting_to_instance_debug/
---
ssh -v -i C:\\Users\\raja\\Downloads\\testDevelopmentAsgb.pem [ubuntu@ec2-15-207-204-117.ap-south-1.compute.amazonaws.com](mailto:ubuntu@ec2-15-207-204-117.ap-south-1.compute.amazonaws.com)

OpenSSH\_for\_Windows\_7.7p1, LibreSSL 2.6.5

debug1: Reading configuration data C:\\\\Users\\\\raja/.ssh/config

debug1: Connecting to [ec2-15-207-204-117.ap-south-1.compute.amazonaws.com](https://ec2-15-207-204-117.ap-south-1.compute.amazonaws.com) \[[15.207.204.117](https://15.207.204.117)\] port 22.

debug1: Connection established.

debug1: key\_load\_public: No such file or directory

debug1: identity file C:\\\\Users\\\\raja\\\\Downloads\\\\testDevelopmentAsgb.pem type -1

debug1: key\_load\_public: No such file or directory

debug1: identity file C:\\\\Users\\\\raja\\\\Downloads\\\\testDevelopmentAsgb.pem-cert type -1

debug1: Local version string SSH-2.0-OpenSSH\_for\_Windows\_7.7
