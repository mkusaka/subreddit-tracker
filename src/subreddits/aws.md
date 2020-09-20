# aws
## [1][Acloudguru is scamming people. Secretly removed Linuxacademy courses and replaced it with their inferior content](https://www.reddit.com/r/aws/comments/ivwj2w/acloudguru_is_scamming_people_secretly_removed/)
- url: https://www.reddit.com/r/aws/comments/ivwj2w/acloudguru_is_scamming_people_secretly_removed/
---
**Acloudguru is scamming people and going back on their promise.**

When Acloudguru took over LinuxAcademy they assured us that we will have access to both catalog of courses. This was a lie.

I paid for Linuxacademy yearly subscription to access their AWS Architect Pro and Devops Pro courses. 

**When I logged in a few days ago I found out that ACG removed 50 hour Aws Architect Pro Linuxacademy course by Adrian Cantrill and replaced it with their ACG inferior 14 hour course by Scott Pelter**

**ACG removed 32 hour Devops Pro course and replaced it with their garbage 6 hour course. In actuality it’s only 4 hours!! Because they sneakily marked each section quiz as 4 hours long and added it to course total.**

This is clearly not what I and other Linuxacademy members paid for. We would like the content that we paid for. Ryan Kroonenburg should be ashamed of himself for scamming people.

I opened a ticket and was told by ACG rep that if I didn’t watch any video from Linuxacademy AWS Pro courses before then I won’t have access to them. Which is completely the opposite of what we were told when ACG took over.

They are slowly replacing all LinuxAcademy courses with shorter, vomit inducing ACG products. 

**Also they sneakily inflate course length by making their quizzes as 4 hour long each. For example there are 6 quiz for AWS Devops Pro exam. So 6 x 4 is 24 hours. The total length of AWS Devops pro course advertised by ACG is 27 hours. So there is only 3 hours of content. No really, go check!**

Linux academy had such great courses and content. Acloudguru is completely destroying all of its credibility and scamming people on top of it. I advise not to get any subscription with them. 

Rather support people like Stephen Maarek, Adrian Cantrill, Eissa Sharif, Neal Davis etc.
## [2][Holy Mother of God. . . AWS should just drop CloudFormation support and let Terraform be the official IaC provider for the platform](https://www.reddit.com/r/aws/comments/iw0xeb/holy_mother_of_god_aws_should_just_drop/)
- url: https://www.reddit.com/r/aws/comments/iw0xeb/holy_mother_of_god_aws_should_just_drop/
---
Just spend some time today fiddling with CloudFormation for a side project (helping a friend). . . and I ran into these two issues:

https://github.com/aws-cloudformation/aws-cloudformation-coverage-roadmap/issues/239

https://github.com/aws/containers-roadmap/issues/631

I actually love AWS. . . most of the time but it's bullshit like which makes me go nuts.  

u/JeffBarr, I have a proposal for you:

Either just give up on CloudFormation or make official CloudFormation support for all resource properties criteria for any service to go GA. 

Another idea: have your service teams manage one instance of their service with CloudFormation somewhere, ongoing.

Another idea: buy Hashicorp.

Another idea: stop sucking so much, you can do better than this. You SHOULD be doing better than this.
## [3][Does VPC can handle VLAN tagged traffic?](https://www.reddit.com/r/aws/comments/iwc30u/does_vpc_can_handle_vlan_tagged_traffic/)
- url: https://www.reddit.com/r/aws/comments/iwc30u/does_vpc_can_handle_vlan_tagged_traffic/
---
Hi All,  


We are in the stage of developing a SDN component within an EC2 instance(preferably metal instances ) and have some concern regarding the traffic from this SDN part to the VPC.

1- If we send VLAN tagged packets from an EC2 instance through ENI, is there any possible way VPC can handle this traffic and route accordingly?

2- Is it possible to configure the ENI from VPC more similar to a trunk port from switch?

3- While setting subnets in VPC is it possible , each of these subnet bounding to specific vlan? 

4- Does the VPC subnets can be configured with complete isolation(no inter-subnet routing) 

Please shed some light to VPC backbone.  


Thanks and Regards,
## [4][A Cloud Guru slowly removing all Linux Academy courses!](https://www.reddit.com/r/aws/comments/ivyg4f/a_cloud_guru_slowly_removing_all_linux_academy/)
- url: /r/AWSCertifications/comments/ivi9nt/a_cloud_guru_slowly_removing_all_linux_academy/
---

## [5][What are the Biggest Issues You've Had With AWS Resource Management?](https://www.reddit.com/r/aws/comments/iw3z7h/what_are_the_biggest_issues_youve_had_with_aws/)
- url: https://www.reddit.com/r/aws/comments/iw3z7h/what_are_the_biggest_issues_youve_had_with_aws/
---
I've been working with AWS for about a year now in some environments that had less than ideal controls in place, and I'm curious what other people have ran into when it comes to managing things in AWS on a metadata perspective.

Biggest thing for me has probably been tag enforcement, especially with app teams setting up tags that had typos or capitalization issues. 

Also a lot of garbage snapshots/AMIs left over, people making changes to resources that were deployed by CFTs instead of updating the CFT, etc.

Any other big problems that people have run into that I should be on the lookout for so I can setup controls to prevent/remediate?
## [6][No way to restore a backup that was exported from RDS to S3?](https://www.reddit.com/r/aws/comments/ivt9p6/no_way_to_restore_a_backup_that_was_exported_from/)
- url: https://www.reddit.com/r/aws/comments/ivt9p6/no_way_to_restore_a_backup_that_was_exported_from/
---
Just got informed by AWS support that there is no way to restore a backup that you have taken from RDS (Aurora MySQL) to S3. It's a Parquet backup, which they say they can't restore. If that's the case, what is the point of backing up your RDS DB to S3? Also what can you do with a Parquet backup? And how do you backup your DB outside of RDS to be restored in the future? Help me guys, my world has been shaken.
## [7][Can ACM be used for client certificates on physical devices?](https://www.reddit.com/r/aws/comments/iw6mql/can_acm_be_used_for_client_certificates_on/)
- url: https://www.reddit.com/r/aws/comments/iw6mql/can_acm_be_used_for_client_certificates_on/
---
Is ACM only for devices hosted in AWS?

I’d like to see if ACM could be used a complete substitute for an on premises Microsoft PKI or is it only useable for AWS-related services. 

For instance is there any way to effectively distribute (such as push certificates to devices with an MDM like Intune, SCCM,or ADCS auto-enrolment SCEP etc.) client authentication certificates to AD users and devices for uses such as EAP-TLS authentication for WPA2 Enterprise 802.11x WiFi, email digital signatures, certificates for identifying a device as a trusted device to access an Amazon Workspace etc.?

What about SSL web server certificates for internal servers and intranet sites not hosted in AWS?
## [8][Having trouble debugging an issue with table_import_from_s3](https://www.reddit.com/r/aws/comments/iw69k6/having_trouble_debugging_an_issue_with_table/)
- url: https://www.reddit.com/r/aws/comments/iw69k6/having_trouble_debugging_an_issue_with_table/
---
I have an RDS instance with an attached IAM role which allows it to sts and talk to the "athena-results" bucket listed below. It's in a public subnet with a security group that allows outbound 443 and 80. I'm running the following query:

    select aws_s3.table_import_from_s3(
                   'etl.stating_table',
                   '',
                   '(format csv)',
                           'athena-results',
                           '0cae18bb-6fb5-4b23-a735-aede565bc345.csv',
                           'us-east-2'
               );

And it's returning this

    [XX000] CURL error code: 28 when attempting to validate pre-signed URL, 1 attempt(s) remaining
    [XX000] HINT: make sure your instance is able to connect with S3.
    [XX000] CURL error code: 28 when attempting to validate pre-signed URL, 0 attempt(s) remaining
    [XX000] HINT: make sure your instance is able to connect with S3.

It's pretty clear this is some sort of network routing issue and not a credentials issue, as when you make a bad request to the API it tends to give you a 401 immediately, not time out. However, im just not sure how to continue debugging because i've checked and rechecked the RDS instance and its subnet group is three \*public\* subnets with an internet gateway attached and a security group with sufficient outbound TCP / UDP rules such that it shouldnt interfere with cURL.

Can anyone help? I feel really up the creek without a paddle here.

&amp;#x200B;

( I did [post on the forums](https://forums.aws.amazon.com/thread.jspa?messageID=957175#957175) about this but i see so many unanswered questions on there I'm not gonna take my chances.)

&amp;#x200B;

***\*EDIT\**** I figure it out but im not happy and i'd really like some feedback if anyone feels like it.

When i change the RDS instance "Public accessibility" option to "yes" I get a different error immediately:

    [XX000] CURL error code: 51 when attempting to validate pre-signed URL, 1 attempt(s) remaining
    [XX000] CURL error code: 51 when attempting to validate pre-signed URL, 0 attempt(s) remaining

From there I took a wild guess, after working with AWS for 8 years, i removed the "." characters from my bucket name and it worked. Sometimes I really wish i could write a stern letter to some of the people who maintain this crap.
## [9][Receiving GET 400 when trying to implement socket.io communication. Node.js, ELB, Route 53, CloudFront, and s3](https://www.reddit.com/r/aws/comments/iw67wi/receiving_get_400_when_trying_to_implement/)
- url: https://www.reddit.com/r/aws/comments/iw67wi/receiving_get_400_when_trying_to_implement/
---
Wasting all my days away :-( Please send help. 

https://stackoverflow.com/questions/63964233/receiving-get-400-when-trying-to-implement-socket-io-communication-node-js-elb
## [10][How to host a Node js program on AWS](https://www.reddit.com/r/aws/comments/iw4x20/how_to_host_a_node_js_program_on_aws/)
- url: https://www.reddit.com/r/aws/comments/iw4x20/how_to_host_a_node_js_program_on_aws/
---
I am new to AWS as of today so please bare with me.

I successfully set up a mssql RDS database today. Then I created a node js program that creates a local host to host an API for the database. When hosted locally, I was able to connect to my AWS RDS database and query it.  What is the best option in the free tier to now host this node js program so that I can access it from any device? I tried using elastic beanstalk, but have not had any success with that yet.
