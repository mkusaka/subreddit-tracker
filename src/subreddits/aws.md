# aws
## [1][No internet access with NAT Gateway on a Private Subnet](https://www.reddit.com/r/aws/comments/f5qitt/no_internet_access_with_nat_gateway_on_a_private/)
- url: https://www.reddit.com/r/aws/comments/f5qitt/no_internet_access_with_nat_gateway_on_a_private/
---
Hi Everyone. 

I'm setting up a Private Subnet for my Lambdas but they don't seem to have internet access.

`Private Subnet B` and `Private Subnet D` are both set up the same. Here are screenshots of `Private Subnet B`. They have the route tables assigned with [0.0.0.0/0](https://0.0.0.0/0) forwarded to a NAT Gateway. The network ACLs looks correct. The lambda has a Security group that allows all outbound traffic.

YET, when I make a call inside the lambda to the internet, it fails. Any ideas? Thanks

&amp;#x200B;

[Private Subnet configuration](https://preview.redd.it/e39ji4mg4oh41.png?width=472&amp;format=png&amp;auto=webp&amp;s=13a9613526a0c57e2e8d57c41dd44d670bd8d673)

[Network ACL](https://preview.redd.it/90dl8wpd4oh41.png?width=938&amp;format=png&amp;auto=webp&amp;s=a401554e71cf09949313cb5fa63537252ecdcad0)

&amp;#x200B;

[Lambda VPC setup](https://preview.redd.it/vb2qvx0r4oh41.png?width=579&amp;format=png&amp;auto=webp&amp;s=50e97dff9e4c38db0f3196dd8eaa176dfd0b7ad1)

&amp;#x200B;

[Failed to connect to public ip \(Amazon SES\)](https://preview.redd.it/qw30vky85oh41.png?width=1110&amp;format=png&amp;auto=webp&amp;s=82c29592dcf547d00fb62060f17cb3595c4268a8)
## [2][AWS Released Managed Apache Cassandra. Here's What You Need to Know.](https://www.reddit.com/r/aws/comments/f5e8jl/aws_released_managed_apache_cassandra_heres_what/)
- url: https://blocks.directory/blog/aws-just-released-managed-cassandra-service
---

## [3][MS CAL License on AWS Workspaces](https://www.reddit.com/r/aws/comments/f5rklu/ms_cal_license_on_aws_workspaces/)
- url: https://www.reddit.com/r/aws/comments/f5rklu/ms_cal_license_on_aws_workspaces/
---
I'm analyzing Workspaces for an internal project and I'm trying to understand how are the CAL licenses for workstations connecting to AD for example?

Its already in the cost of the VDI provided by AWS ?
## [4][Will I pay less with async programming in Lambda](https://www.reddit.com/r/aws/comments/f5qe3v/will_i_pay_less_with_async_programming_in_lambda/)
- url: https://www.reddit.com/r/aws/comments/f5qe3v/will_i_pay_less_with_async_programming_in_lambda/
---
If I do a non blocking long http request I'm not using any cpu. Am I then billed less or is the total execution time?
## [5][Failover: Aurora Classic vs Aurora Serverless](https://www.reddit.com/r/aws/comments/f5qcl3/failover_aurora_classic_vs_aurora_serverless/)
- url: https://www.reddit.com/r/aws/comments/f5qcl3/failover_aurora_classic_vs_aurora_serverless/
---
Can somebody confirm if my understanding is correct

1. Aurora Serverless operates on a single AZ only, but can failover on another AZ by default. This also means your data is technically multi-AZ but you CAN'T use the other AZ's as replicas.
2. Aurora Classic Multi-AZ mode can failover to another AZ, and you can also use those replicas on the other AZ's anytime, even without disaster.
3. Aurora Classic Single-AZ mode will NOT failover to another AZ, but your data is alive in another AZ. You must manually instantiate an Aurora instance on a different AZ to recover your data.
## [6][A React Native starter app with fully-functioning AWS Cognito functionality (without Expo)](https://www.reddit.com/r/aws/comments/f5bn79/a_react_native_starter_app_with_fullyfunctioning/)
- url: https://github.com/vbudilov/react-native-starter
---

## [7][Expected DX Latency from Denver to us-west-2](https://www.reddit.com/r/aws/comments/f5nij8/expected_dx_latency_from_denver_to_uswest2/)
- url: https://www.reddit.com/r/aws/comments/f5nij8/expected_dx_latency_from_denver_to_uswest2/
---
We co-locate with one of the bigger datacenter providers in the Denver area, and from them I ordered a DirectConnect into the us-west-2 region.  This is a sub-1G, hosted connection.

I know the AWS DX Denver Location is at not my current datacenter, but at CoreSite DE1, and I assume the AWS router BGP peer IP is located here, since I get an average 1.2ms rtt to that IP address from my router. 
 
But from the same router to an EC2 instance in a us-west-2 VPC, the rtt is about 45ms, only a very marginal improvement over the 50ms average latency we see from our IPSEC VPN into the same region.  One of the reasons for purchasing the DX connection was to reduce the latency we saw over the VPN connection.
  
Before I open tickets with AWS and the datacenter provider, I was wondering what latency I should expect to see from the Denver AWS location into a VPC in us-west-2.  Thanks!
## [8][EC2 Domain join failing](https://www.reddit.com/r/aws/comments/f5piws/ec2_domain_join_failing/)
- url: https://www.reddit.com/r/aws/comments/f5piws/ec2_domain_join_failing/
---
Hi,

I have an account with an AWS AD domain shared to it from another account. When trying to configure an instance to join the domain automatically, it is failing, although there are no errors in the system.log (so not IAM related I don't think), and joining the domain manually works fine. I am configuring an IAM instance, and choosing to join the domain while launching the instance.

One thing I have noticed is that the DNS settings are not even being set on the instance.

Any suggestions on where I can look to figure this out? I was thinking that it is either a problem setting DNS (no idea why though)?

Here is my EC2Launch.log file.

&amp;#x200B;

&amp;#x200B;

2020/02/17 11:33:18Z: Initializing instance is started

2020/02/17 11:33:20Z: Opening port (COM1) handle to write to the console

2020/02/17 11:33:22Z: Windows sysprep configuration complete.

2020/02/17 11:33:22Z: Checking primary network interface

2020/02/17 11:33:22Z: Primary network interface found. Adding routes now...

2020/02/17 11:33:24Z: Successfully added the Route: [169.254.169.254/32](https://169.254.169.254/32), gateway: [172.31.32.1](https://172.31.32.1), NIC index: 7, Metric: 15

2020/02/17 11:33:24Z: Successfully added the Route: [169.254.169.250/32](https://169.254.169.250/32), gateway: [172.31.32.1](https://172.31.32.1), NIC index: 7, Metric: 15

2020/02/17 11:33:24Z: Successfully added the Route: [169.254.169.251/32](https://169.254.169.251/32), gateway: [172.31.32.1](https://172.31.32.1), NIC index: 7, Metric: 15

2020/02/17 11:33:24Z: Successfully added the Route: [169.254.169.249/32](https://169.254.169.249/32), gateway: [172.31.32.1](https://172.31.32.1), NIC index: 7, Metric: 15

2020/02/17 11:33:24Z: Successfully added the Route: [169.254.169.123/32](https://169.254.169.123/32), gateway: [172.31.32.1](https://172.31.32.1), NIC index: 7, Metric: 15

2020/02/17 11:33:24Z: Successfully added the Route: [169.254.169.253/32](https://169.254.169.253/32), gateway: [172.31.32.1](https://172.31.32.1), NIC index: 7, Metric: 15

2020/02/17 11:33:24Z: Message: Waiting for meta-data accessibility...

2020/02/17 11:33:24Z: Message: Meta-data is now available.

2020/02/17 11:33:24Z: Reading launch config from file

2020/02/17 11:33:25Z: Finished reading launch configs:

SetComputerName: False

SetMonitorAlwaysOn: True

SetWallpaper: True

AddDnsSuffixList: True

ExtendBootVolumeSize: True

HandleUserData: True

AdminPasswordType: Random

&amp;#x200B;

2020/02/17 11:33:25Z: Creating wallpaper setup cmd in startup directory

2020/02/17 11:33:26Z: EG is not available on this instance

2020/02/17 11:33:26Z: Error installing eGPU

2020/02/17 11:33:26Z: Setting computer name is disabled

2020/02/17 11:33:26Z: AMI Origin Version: 2020.02.12

2020/02/17 11:33:26Z: AMI Origin Name: Windows\_Server-2016-English-Full-Base

2020/02/17 11:33:26Z: OS: Microsoft Windows NT 10.0

2020/02/17 11:33:26Z: OsProductName: Windows Server 2016 Datacenter

2020/02/17 11:33:26Z: OsInstallOption: Full

2020/02/17 11:33:26Z: OsVersion: 10.0

2020/02/17 11:33:26Z: OsBuildLabEx: 14393.3503.amd64fre.rs1\_release.200131-0410

2020/02/17 11:33:26Z: Language: en-US

2020/02/17 11:33:26Z: TimeZone: Coordinated Universal Time

2020/02/17 11:33:26Z: Offset: UTC 00:00:00

2020/02/17 11:33:26Z: AMI-ID: ami-0648c16a1a9bd20dc

2020/02/17 11:33:26Z: Instance-ID: i-0cafe2ca75f5xxxxx

2020/02/17 11:33:26Z: Instance Type: m4.xlarge

2020/02/17 11:33:26Z: No SQL Billing Codes Associated With Instance

2020/02/17 11:33:43Z: Driver: AWS PV Driver Package v8.3.2

2020/02/17 11:33:44Z: Driver: Intel(R) 82599 Virtual Function v2.1.186.0

2020/02/17 11:33:44Z: Launch: EC2 Launch v1.3.2002240

2020/02/17 11:33:44Z: SSM: Amazon SSM Agent v2.3.722.0

2020/02/17 11:33:44Z: RDPCERTIFICATE-SUBJECTNAME: EC2AMAZ-FOHxxxx

2020/02/17 11:33:44Z: RDPCERTIFICATE-THUMBPRINT: DA3B23523FEB56AF4846905555DD59EC8Dxxxxxx

2020/02/17 11:33:47Z: Adding DNS suffixes in search list begins

2020/02/17 11:33:47Z: Adding DNS suffixes in search list done

2020/02/17 11:33:47Z: Executing boot volume extension

2020/02/17 11:33:49Z: Finish executing boot volume extension

2020/02/17 11:33:50Z: Failed to get metadata: The result from [http://169.254.169.254/latest/meta-data/hibernation/configured](http://169.254.169.254/latest/meta-data/hibernation/configured) was empty

2020/02/17 11:33:51Z: Metadata Check For Enabling Hibernation Failed: Exception calling "Parse" with "1" argument(s): "String was not recognized as a valid Boolean."

2020/02/17 11:33:51Z: HibernationEnabled: false

2020/02/17 11:33:51Z: Generating a random password...

2020/02/17 11:33:51Z: Failed to get metadata: The result from [http://169.254.169.254/latest/meta-data/public-keys/0/openssh-key](http://169.254.169.254/latest/meta-data/public-keys/0/openssh-key) was empty

2020/02/17 11:33:51Z: Unable to send the password to console: You cannot call a method on a null-valued expression.

2020/02/17 11:33:51Z: Message: Windows is Ready to use

2020/02/17 11:34:01Z: Instance initialization is scheduled successfully

2020/02/17 13:25:00Z: Initializing instance is started

2020/02/17 13:25:02Z: Opening port (COM1) handle to write to the console

2020/02/17 13:25:04Z: Windows sysprep configuration complete.

2020/02/17 13:25:04Z: Checking primary network interface

2020/02/17 13:25:04Z: Primary network interface found. Adding routes now...

2020/02/17 13:25:06Z: Successfully added the Route: [169.254.169.254/32](https://169.254.169.254/32), gateway: [10.4.6.1](https://10.4.6.1), NIC index: 8, Metric: 15

2020/02/17 13:25:06Z: Successfully added the Route: [169.254.169.250/32](https://169.254.169.250/32), gateway: [10.4.6.1](https://10.4.6.1), NIC index: 8, Metric: 15

2020/02/17 13:25:06Z: Successfully added the Route: [169.254.169.251/32](https://169.254.169.251/32), gateway: [10.4.6.1](https://10.4.6.1), NIC index: 8, Metric: 15

2020/02/17 13:25:06Z: Successfully added the Route: [169.254.169.249/32](https://169.254.169.249/32), gateway: [10.4.6.1](https://10.4.6.1), NIC index: 8, Metric: 15

2020/02/17 13:25:06Z: Successfully added the Route: [169.254.169.123/32](https://169.254.169.123/32), gateway: [10.4.6.1](https://10.4.6.1), NIC index: 8, Metric: 15

2020/02/17 13:25:06Z: Successfully added the Route: [169.254.169.253/32](https://169.254.169.253/32), gateway: [10.4.6.1](https://10.4.6.1), NIC index: 8, Metric: 15

2020/02/17 13:25:06Z: Message: Waiting for meta-data accessibility...

2020/02/17 13:25:06Z: Message: Meta-data is now available.

2020/02/17 13:25:06Z: Reading launch config from file

2020/02/17 13:25:07Z: Finished reading launch configs:

SetComputerName: False

SetMonitorAlwaysOn: True

SetWallpaper: True

AddDnsSuffixList: True

ExtendBootVolumeSize: True

HandleUserData: True

AdminPasswordType: Random

&amp;#x200B;

2020/02/17 13:25:07Z: Creating wallpaper setup cmd in startup directory

2020/02/17 13:25:07Z: EG is not available on this instance

2020/02/17 13:25:08Z: Error installing eGPU

2020/02/17 13:25:08Z: Setting computer name is disabled

2020/02/17 13:25:08Z: AMI Origin Version: 2020.02.12

2020/02/17 13:25:08Z: AMI Origin Name: Windows\_Server-2016-English-Full-Base

2020/02/17 13:25:08Z: OS: Microsoft Windows NT 10.0

2020/02/17 13:25:08Z: OsProductName: Windows Server 2016 Datacenter

2020/02/17 13:25:08Z: OsInstallOption: Full

2020/02/17 13:25:08Z: OsVersion: 10.0

2020/02/17 13:25:08Z: OsBuildLabEx: 14393.3503.amd64fre.rs1\_release.200131-0410

2020/02/17 13:25:08Z: Language: en-US

2020/02/17 13:25:08Z: TimeZone: Coordinated Universal Time

2020/02/17 13:25:08Z: Offset: UTC 00:00:00

2020/02/17 13:25:08Z: AMI-ID: ami-0648c16a1a9bd20dc

2020/02/17 13:25:08Z: Instance-ID: i-0dc99389f6d3xxxxx

2020/02/17 13:25:08Z: Instance Type: t3.large

2020/02/17 13:25:08Z: No SQL Billing Codes Associated With Instance

2020/02/17 13:25:25Z: Driver: AWS NVMe Driver v1.3.2.53

2020/02/17 13:25:25Z: Driver: AWS PV Driver Package v8.3.2

2020/02/17 13:25:26Z: Driver: Amazon Elastic Network Adapter v2.1.4.0

2020/02/17 13:25:26Z: Launch: EC2 Launch v1.3.2002240

2020/02/17 13:25:26Z: SSM: Amazon SSM Agent v2.3.722.0

2020/02/17 13:25:26Z: RDPCERTIFICATE-SUBJECTNAME: EC2AMAZ-FOHxxxx

2020/02/17 13:25:26Z: RDPCERTIFICATE-THUMBPRINT: DA3B23523FEB56AF4846905555DD59EC8Dxxxxxx

2020/02/17 13:25:29Z: Adding DNS suffixes in search list begins

2020/02/17 13:25:29Z: Adding DNS suffixes in search list done

2020/02/17 13:25:29Z: Executing boot volume extension

2020/02/17 13:25:31Z: Finish executing boot volume extension

2020/02/17 13:25:32Z: Failed to get metadata: The result from [http://169.254.169.254/latest/meta-data/hibernation/configured](http://169.254.169.254/latest/meta-data/hibernation/configured) was empty

2020/02/17 13:25:32Z: Metadata Check For Enabling Hibernation Failed: Exception calling "Parse" with "1" argument(s): "String was not recognized as a valid Boolean."

2020/02/17 13:25:32Z: HibernationEnabled: false

2020/02/17 13:25:32Z: Generating a random password...

2020/02/17 13:25:32Z: Username: Administrator

2020/02/17 13:25:32Z: Password: &lt;Password&gt;

xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxpqKJHzuftqrJ7wMA17wnBdYA2E2JBVkYQiHwLA2gru40KShlih9LDkVU54numaWJWAFHBzdl8G8DhuwPPU2hETWgDEYCFBPuMmFjoS2RJ2PpucZzPyEwkBEW9xxxxxHdQ6FMKHx+1me48y/LEE9Yu0IqwWOfJIgPAAxo37RrPxH4tDCODaWVPw3AOQ==

&lt;/Password&gt;  


2020/02/17 13:25:32Z: Message: Windows is Ready to use

2020/02/17 13:25:37Z: Job 'AmazonSSMAgent' Still In Running State, Sleeping For '5' Seconds

2020/02/17 13:25:42Z: Job 'AmazonSSMAgent' Still In Running State, Sleeping For '5' Seconds

2020/02/17 13:25:47Z: Job 'AmazonSSMAgent' Still In Running State, Sleeping For '5' Seconds

2020/02/17 13:25:52Z: Job 'AmazonSSMAgent' Still In Running State, Sleeping For '5' Seconds

2020/02/17 13:25:57Z: Job 'AmazonSSMAgent' Still In Running State, Sleeping For '5' Seconds

2020/02/17 13:26:02Z: Job 'AmazonSSMAgent' Still In Running State, Sleeping For '5' Seconds

2020/02/17 13:26:07Z: Job 'AmazonSSMAgent' Still In Running State, Sleeping For '5' Seconds

2020/02/17 13:26:12Z: Job 'AmazonSSMAgent' Still In Running State, Sleeping For '5' Seconds

2020/02/17 13:26:17Z: Job 'AmazonSSMAgent' Still In Running State, Sleeping For '5' Seconds

2020/02/17 13:26:22Z: Job 'AmazonSSMAgent' Still In Running State, Sleeping For '5' Seconds

2020/02/17 13:26:27Z: Job 'AmazonSSMAgent' Still In Running State, Sleeping For '5' Seconds

2020/02/17 13:26:32Z: Job 'AmazonSSMAgent' Still In Running State, Sleeping For '5' Seconds

2020/02/17 13:26:37Z: Job 'AmazonSSMAgent' Still In Running State, Sleeping For '5' Seconds

2020/02/17 13:26:42Z: Job 'AmazonSSMAgent' Still In Running State, Sleeping For '5' Seconds

2020/02/17 13:26:47Z: Job 'AmazonSSMAgent' Still In Running State, Sleeping For '5' Seconds

2020/02/17 13:26:52Z: Job 'AmazonSSMAgent' Still In Running State, Sleeping For '5' Seconds

2020/02/17 13:26:57Z: Job 'AmazonSSMAgent' Still In Running State, Sleeping For '5' Seconds

2020/02/17 13:27:02Z: Job 'AmazonSSMAgent' Still In Running State, Sleeping For '5' Seconds

2020/02/17 13:27:07Z: Job 'AmazonSSMAgent' Still In Running State, Sleeping For '5' Seconds

2020/02/17 13:27:12Z: Job 'AmazonSSMAgent' Still In Running State, Sleeping For '5' Seconds

2020/02/17 13:27:17Z: Job 'AmazonSSMAgent' Finished With Status 'Completed': ''

2020/02/17 13:27:17Z: Initializing instance is done
## [9][Multiple subnet communication between several EC2 instances broken](https://www.reddit.com/r/aws/comments/f5othv/multiple_subnet_communication_between_several_ec2/)
- url: https://www.reddit.com/r/aws/comments/f5othv/multiple_subnet_communication_between_several_ec2/
---
Recently an app update was made to my application which seems to have broken my subnet to subnet communication with my containers (same VPC) over HTTP. I have verified my routing table, security group, and network ACL, and they all look fine. But I'm unable to communicate via HTTP from one subnet to another. Any ideas on a direction?
## [10][I need some tips in reducing IOPS cost for MySQL Aurora Cluster](https://www.reddit.com/r/aws/comments/f5au0n/i_need_some_tips_in_reducing_iops_cost_for_mysql/)
- url: https://www.reddit.com/r/aws/comments/f5au0n/i_need_some_tips_in_reducing_iops_cost_for_mysql/
---
Recently, after launching a new app our Aurora IOPS costs have significantly gone up. This is mostly due to database write operations. 

I was wondering if there are ways to reduce this cost for example by storing all the write queries and combining them into a single large write.

I am not too familiar with how IOPS are calculated by AWS and hence this question.
