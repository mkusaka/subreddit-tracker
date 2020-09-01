# aws
## [1][Week of Aug 31st - What AWS tips do you have?](https://www.reddit.com/r/aws/comments/ik3van/week_of_aug_31st_what_aws_tips_do_you_have/)
- url: https://www.reddit.com/r/aws/comments/ik3van/week_of_aug_31st_what_aws_tips_do_you_have/
---
Share your tips with the community.
## [2][Any ideas on why my Amazon workspace account gets progressively more laggy from logging on](https://www.reddit.com/r/aws/comments/ikiq0t/any_ideas_on_why_my_amazon_workspace_account_gets/)
- url: https://www.reddit.com/r/aws/comments/ikiq0t/any_ideas_on_why_my_amazon_workspace_account_gets/
---
(Apologies if this is not the place to ask this) I am currently working from home using Amazon Workspace to log into my work desktop. I did this for a couple months between March and May this year with no issue, was put on to furlough and have returned to working from home since early July.

Since then I have had continuous lag both on the internet and over the phone (via the ring central app to call using my browser). However this is not my connection (150 download 15 upload) nor my hardware. Also it seemingly gets progressively worse over the work day, both my phone and my general work has no issues then begins to get worse.

I have asked my IT department to have a look at it but they are pretty useless and have ‘found no issues’ with my account although around 2 weeks ago they reset my account to a previous state and one of the IT guys advised there was a major fault of some description with my account although he could not specify what nor did he fix it.

Apologies for the boring, long explanation but If anyone has any suggestions or ideas as to what the issue could be I would be very grateful as I have been told that I will have to return to the office if this can’t be fixed and I would rather not have to travel in 2 hours+ each day (plus the costs of trains) but I have no IT knowhow so have no idea how to fix this...
## [3][Bottlerocket OS 1.0 release](https://www.reddit.com/r/aws/comments/ik3smc/bottlerocket_os_10_release/)
- url: https://github.com/bottlerocket-os/bottlerocket/releases/tag/v1.0.0
---

## [4][JWT auth using React + Node (Cognito+S3+EC2)](https://www.reddit.com/r/aws/comments/ikjnw0/jwt_auth_using_react_node_cognitos3ec2/)
- url: https://www.reddit.com/r/aws/comments/ikjnw0/jwt_auth_using_react_node_cognitos3ec2/
---
Hello folks,

So I have configured my node application to fetch the jwks.json with the pair of keys and verify the token received in the headers (Autorization: Bearer \[token\]) in order to allow access to the API routes.

The problem is within React app. So I send the token in the request headers, however I put the token string by hand. How I make to extract it (by code) from the cookie and send it along the request?

Also, is there a more simple approach out there?
## [5][Anyone ever successfully exported an EC2 instance and run it in virtualbox?](https://www.reddit.com/r/aws/comments/ike3ex/anyone_ever_successfully_exported_an_ec2_instance/)
- url: https://www.reddit.com/r/aws/comments/ike3ex/anyone_ever_successfully_exported_an_ec2_instance/
---
I launched and successfully ran t-mobile's [tpotce](https://github.com/telekom-security/tpotce) in my AWS account to capture honeypot data. Once I was done, I wanted to shut off and export the instance to keep for review purposes. I exported the instance following the documentation [here](https://docs.aws.amazon.com/vm-import/latest/userguide/vmexport.html). The syntax for the CLI command I used was 

    aws ec2 create-instance-export-task --description "debian honeypot" --instance-id i-xxxxxxxxxxxxxxx --target-environment vmware --export-to-s3-task DiskImageFormat=vmdk,ContainerFormat=ova,S3Bucket=bucket_name

The AMI I used was debian-10-amd64-20200610-293 (ami-0c24eddbea3a65909). Anytime i try to boot the VM in virtualbox, the image stops booting [here](https://raw.githubusercontent.com/remotephone/remotephone.github.io/master/images/stuck_ec2_boot.png).

I can boot the VM successfully in VMWare player, but networking is completely jacked and I am much more comfortable troubleshooting that in virtualbox. It can't see my interfaces that VMWare player is providing and I've tried a bunch of different things, including using VMWare's network manager tool to add, modify, remove, replace, and whatever else I could do to the network interfaces to get them to come up and no dice. 

In virtualbox, I've tried a whole host of other things, but the main ones is reimporting the OVA an unchecking the import as VDI option (so running it both as a VMDK and a VDI). I've also taken the disk I successfully booted in VMWare and used it as the boot disk in virtualbox with the same problem. Finally, I can mount the disk if I add it as a secondary disk, I  just can't boot from it. 

Thanks!
## [6][ECS Task vs AWS Batch](https://www.reddit.com/r/aws/comments/ikhe17/ecs_task_vs_aws_batch/)
- url: https://www.reddit.com/r/aws/comments/ikhe17/ecs_task_vs_aws_batch/
---
I always use ECS tasks when I have long running processes that needs to be scheduled or needs to run when a specific S3 event occours. So I was wondering: what's the benefit or use case for AWS Batch compared to an ECS task?

&amp;#x200B;

Thanks!
## [7][Finding and deleting orphaned snapshots using AWS CLI in command prompt](https://www.reddit.com/r/aws/comments/ikk2x8/finding_and_deleting_orphaned_snapshots_using_aws/)
- url: https://www.reddit.com/r/aws/comments/ikk2x8/finding_and_deleting_orphaned_snapshots_using_aws/
---
Hi there guys, hope you're well.

I have a pretty straightforward problem. I have a bunch of orphaned EBS snapshots that need to be deleted. When upgrading our software we create these snapshots should anything go wrong. However, as some of you may be aware, even if EC2s are terminated, the snapshots remain and thus need to be deleted manually. There are quite a few articles on this issue online, however I noticed most of them have solutions which requires usage of bash. 

I've also found others such as:

    comm -23 &lt;(aws ec2 describe-snapshots --owner-ids AWS-ACCOUNT-ID --query 'Snapshots[*].SnapshotId' --output text | tr '\t' '\n' | sort) &lt;(aws ec2 describe-volumes --query 'Volumes[*].SnapshotId' --output text | tr '\t' '\n' | sort | uniq)

Any ideas on a workable solution that could be used in CMD??? Thanks.
## [8][[NEED SOME INPUT] What’re some tips for someone starting with Big Data on AWS?](https://www.reddit.com/r/aws/comments/ik78vd/need_some_input_whatre_some_tips_for_someone/)
- url: https://www.reddit.com/r/aws/comments/ik78vd/need_some_input_whatre_some_tips_for_someone/
---
Hello,

I am just starting my journey of big data on AWS. I am still in the researching phase and am looking to indulge in as much information as I can about the subject before making any decisions. 

Do you guys have any tips or tricks with using big data services on AWS? Provide all of the input you can. Hoping to help the next guy who is looking for help as well. (: 

Thanks!
## [9][Performance penalty of multi-az postgres with synchronous_commit=on?](https://www.reddit.com/r/aws/comments/ikhuyn/performance_penalty_of_multiaz_postgres_with/)
- url: https://www.reddit.com/r/aws/comments/ikhuyn/performance_penalty_of_multiaz_postgres_with/
---
I’m currently running a single AZ rds instance with one replica in production. We’re looking to potentially enable multi az for more fault tolerance.

I’m seeing that replication between master and replica in the same AZ is asynchronous. But I was wondering if anyone had experience about the performance penalty of synchronous replication between master/failover nodes in a Multi AZ setup.

Thanks.
## [10][AWS Toolkit keeps creating default security groups for ELB and VPC when I've supplied it with my own](https://www.reddit.com/r/aws/comments/ikdjet/aws_toolkit_keeps_creating_default_security/)
- url: https://www.reddit.com/r/aws/comments/ikdjet/aws_toolkit_keeps_creating_default_security/
---
Hi, 

I currently have this VPC setup:

2 public subnets - for ELB

2 private subnets - for application

2 private subnets - for RDS (db subnet group)

I'm trying to deploy a load balanced elastic beanstalk application and I supply it with my own security groups, but it seems to want to create 2 default ones, one for a VPC and one for the load balancer, is there anyway around this?
## [11][Dynamically play a menu to the user Amazon Connect](https://www.reddit.com/r/aws/comments/ikapvq/dynamically_play_a_menu_to_the_user_amazon_connect/)
- url: https://www.reddit.com/r/aws/comments/ikapvq/dynamically_play_a_menu_to_the_user_amazon_connect/
---
Hello, everyone. I'm working on an AWS Connect project that will allow the user to know her balance on a set of accounts. The user can have a maximum of 9 accounts, or as little as one. So for example if  the user calls an audio like "To check your savings balance, press 1, to know your cheque account balance press 2." And so on depending on the user's accounts. Is this possible or some workaround?
