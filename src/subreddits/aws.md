# aws
## [1][How-to Guide: Debugging a Kubernetes Application](https://www.reddit.com/r/aws/comments/ffssxi/howto_guide_debugging_a_kubernetes_application/)
- url: https://epsagon.com/blog/development/how-to-guide-debugging-a-kubernetes-application/
---

## [2][Is there any use-case for multiple IPs to single instance?](https://www.reddit.com/r/aws/comments/ffjtj3/is_there_any_usecase_for_multiple_ips_to_single/)
- url: https://www.reddit.com/r/aws/comments/ffjtj3/is_there_any_usecase_for_multiple_ips_to_single/
---
Hi, I know that we can attach a second ENI to an EC2 instance. But I was wondering what can be a possible use-case to do that? How can it be useful?

Thanks
## [3][confusing aws lambda kickoff performance](https://www.reddit.com/r/aws/comments/ffol6k/confusing_aws_lambda_kickoff_performance/)
- url: https://www.reddit.com/r/aws/comments/ffol6k/confusing_aws_lambda_kickoff_performance/
---
i have a lambda that kicks off an api request.  here is what happens; database dumps 5 MB files to S3.  S3 initiates lambda.  lambda copies 5 MB in, sends it to an api, gets the response, writes the response to S3.  4 lambdas are allowed to run concurrently.  all succeed.

i expect my api service to have a pretty steady state of 4 things hitting it.  what do i actually see instead?

i see lots of dead space.  example is 13:08 here, nothing is going on.  [https://i.imgur.com/f2DTdrz.png](https://i.imgur.com/f2DTdrz.png)

i have 4 log streams, and all look similar.  the lambdas themselves are like 10-25 seconds.  [https://i.imgur.com/wEHDOKL.png](https://i.imgur.com/wEHDOKL.png)

they all have nothing going on in 13:08.  no lambda initiated in 13:07 could run long enough to result in no lambda being initiated in 13:08, yet all 4 log streams look like this.  amazon is simply doing nothing at all in the 13:08 minute, for no reason i can tell.  what is going on here?  it is super annoying, as it is killing my throughput.

my best guess speculation is that it's some poor/strange limitation of the concurrency limitation.  what is extra confusing though is that it gets even worse over time.  my process was just hanging around while nothing was happening for those last requests at the tail to complete.  i'll note in the kibana log that the first hit in my actual process was 13:03:52 and the last was at that tail before 13:20.  it's about 25% the throughput end to end that i was expecting.

1 minute granularity is the lowest i can get the concurrent charts to produce anything to plot.

max: [https://i.imgur.com/85U5V40.png](https://i.imgur.com/85U5V40.png)

sum: [https://i.imgur.com/hYQpdwi.png](https://i.imgur.com/hYQpdwi.png)

what's with all my dead spots?
## [4][Benchmarking the AWS Graviton2 with KeyDB – M6g up to 65% faster](https://www.reddit.com/r/aws/comments/ffupap/benchmarking_the_aws_graviton2_with_keydb_m6g_up/)
- url: https://docs.keydb.dev/blog/2020/03/02/blog-post/
---

## [5][Distributed Tracing in Asynchronous Applications](https://www.reddit.com/r/aws/comments/ffud9h/distributed_tracing_in_asynchronous_applications/)
- url: https://epsagon.com/blog/distributed-tracing-in-asynchronous-applications/
---

## [6][Is it possible to choose different subnets per service like Workspaces and EC2?](https://www.reddit.com/r/aws/comments/ffub6w/is_it_possible_to_choose_different_subnets_per/)
- url: https://www.reddit.com/r/aws/comments/ffub6w/is_it_possible_to_choose_different_subnets_per/
---
I have set up 3 /24 networks and I am trying to set everything up so that our EC2 instances will use one while our workspaces are on another but they are all just getting DHCP from the same subnet.  I don't see anywhere on how to set them apart.  I'm still new to this so I hope this makes sense.
## [7][Noob Cloudformation Egress Security Groups.... multiple ports.](https://www.reddit.com/r/aws/comments/ffu7a4/noob_cloudformation_egress_security_groups/)
- url: https://www.reddit.com/r/aws/comments/ffu7a4/noob_cloudformation_egress_security_groups/
---
Please can anyone advise me on how to get a list of ports opened up in Cloudformation Security Group INgress rules.

I know it should be a LIST but my YAML keeps failing. When i try other syntax, the security group only seems to pick the LAST port in the list.

Help me Obi wan ... you're my only hope.

    #Setting Up Security Groups
      PublicSecurityGroup:
        Type: AWS::EC2::SecurityGroup
        Properties:
            GroupName: PublicSecurityGroup
            GroupDescription: Public Security Group
            VpcId:
                Ref: VPC
      PublicOutboundRule1:
        Type: AWS::EC2::SecurityGroupEgress
        Properties:
            GroupId: !Ref PublicSecurityGroup
            IpProtocol: tcp
            FromPort: 0
            ToPort: 65535
            CidrIp: 0.0.0.0/0
      PublicInboundRule1:
        Type: AWS::EC2::SecurityGroupIngress
        Properties:
            GroupId: !Ref PublicSecurityGroup
            - IpProtocol: tcp
            FromPort: 80
            ToPort: 80
            - IpProtocol: tcp
            FromPort: 443
            ToPort: 443
            CidrIp: 0.0.0.0/0
## [8][Creating S2S VPN](https://www.reddit.com/r/aws/comments/ffkjwx/creating_s2s_vpn/)
- url: https://www.reddit.com/r/aws/comments/ffkjwx/creating_s2s_vpn/
---
I'm creating IPSec Site to Site VPNs in Azure which I'm confident work but am having difficulty testing them. Despite my best efforts our office firewall isn't letting all the packets through that I need so I'm having to test from AWS.

&amp;#x200B;

I want to create a VPN from AWS to the Azure VPN but the AWS VPN only seems to be for connecting from an office into AWS. It doesn't even give any options for cipher suites of DH groups. Can this be done natively in AWS or do I need to run something like pfSense in an EC2?
## [9][deleting AWS managed KMS keys](https://www.reddit.com/r/aws/comments/fftuf0/deleting_aws_managed_kms_keys/)
- url: https://www.reddit.com/r/aws/comments/fftuf0/deleting_aws_managed_kms_keys/
---
Hi guys,

Noobie here, I was practising on AWS basic plan and created an IAM administrator account for the study guides. As part of the practise I created used Lambda, EFS and EBS which in turn created theyKMS keys for encyrption. 

When I was done I deleted "everything" including the IAM admin account that created those keys. I wish to delete those keys but I don't know how. The guides I'm getting are for deleting CMKs. Please help, I'm tired of paying for keys I'm not using.

(I don't want to delete the root account for fear of being unable to use it again in future)
## [10][AWS re:Invent 2019 Announcements Recap](https://www.reddit.com/r/aws/comments/fftt4s/aws_reinvent_2019_announcements_recap/)
- url: https://epsagon.com/blog/webinars/aws-reinvent-2019-recap/
---

