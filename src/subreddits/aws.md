# aws
## [1][I made a thing to display SNS messages on the command line](https://www.reddit.com/r/aws/comments/ewv00r/i_made_a_thing_to_display_sns_messages_on_the/)
- url: https://github.com/ziggy42/ontopic
---

## [2][Cloudfront not pointing to updated S3 bucket](https://www.reddit.com/r/aws/comments/ewsurw/cloudfront_not_pointing_to_updated_s3_bucket/)
- url: https://www.reddit.com/r/aws/comments/ewsurw/cloudfront_not_pointing_to_updated_s3_bucket/
---
I'm hosting a static website on S3 and am using Cloudfront for the SSL certificate.  Domain is through GoDaddy.  I updated the code and pushed to S3 but the website still isn't showing the updated version.

Any idea why that is?

&amp;#x200B;

&amp;#x200B;

UPDATE

Just took a little bit of time :)   
Thanks everyone.
## [3][Using CodeDeploy for ECS deployment](https://www.reddit.com/r/aws/comments/ewpudk/using_codedeploy_for_ecs_deployment/)
- url: https://www.reddit.com/r/aws/comments/ewpudk/using_codedeploy_for_ecs_deployment/
---
Hello.Out of curiosity, I was wondering if any of you used CodeDeploy as the DeploymentController for your ECS Services over the native ECS one, and what lessons you learnt from it, what made you go for it as opposed to the native one, etc.

&amp;#x200B;

Also, in the case of single account / multiple environments, do you create 1 application globally and then different deployment groups for each environment or multiple applications representing the same app?

&amp;#x200B;

Thank you!
## [4][Question - Can't SSH into an EC2 Instance](https://www.reddit.com/r/aws/comments/ewp0aa/question_cant_ssh_into_an_ec2_instance/)
- url: https://www.reddit.com/r/aws/comments/ewp0aa/question_cant_ssh_into_an_ec2_instance/
---
Hi all, so for some reason I can't ssh into an Ec2 instance  


Below is basically everything - its a Kali aws marketplace box "I killed this box so its not up right now"  


* I set up multiple rules the SG to allow ssh in and out to both my IP and anywhere. 
*  I made sure the SG was assigned.  
* I allowed ping and couldnt ping
* i ensure vpc had rules as well
* I followed the rules to save the ppk key in putty
   * I downloaded the prem key
   * loaded the prem key in puttygen
   * saved the prem key without a password as a private key
   * loaded the prem key into putty per directions (SSH &gt; AUTH)  

* I KNOW EXTERNAL SSH WORKS - I can ssh into azure.
* this failed on both kalli and nessus
* tried t2.med and t2.micro
* Connection just times out???  

* tried on two different AWS accounts
   * 6 different instances
   * 2 rebuilt instances
   * 2 new SGs and VPCs
   * 2 New inbound and outbound rules
* Tried in linux using the prem key and ssh -i

&amp;#x200B;

&amp;#x200B;

Theres got to be a simple fix and i really dont know what it is  


  


https://preview.redd.it/b6gql3t6f4e41.png?width=1549&amp;format=png&amp;auto=webp&amp;s=3633a17ea597afde7db8e9329fb69cd38138fd8c
## [5][Provisioned IOPS vs GP2 vs throughput optimized](https://www.reddit.com/r/aws/comments/ewpna9/provisioned_iops_vs_gp2_vs_throughput_optimized/)
- url: https://www.reddit.com/r/aws/comments/ewpna9/provisioned_iops_vs_gp2_vs_throughput_optimized/
---
From what I'm understanding provisioned IOPS is for shit being accessed alot (costs a fuckton)

However I don't really get a the use cases for GP2 over throughput optimized or visa versa?

Can someone provide a use case where GP2 would be better than provisioned throughput and a use case where the opposite is true?
## [6][CFN volume attachments and userdata order of execution](https://www.reddit.com/r/aws/comments/ewyphv/cfn_volume_attachments_and_userdata_order_of/)
- url: https://www.reddit.com/r/aws/comments/ewyphv/cfn_volume_attachments_and_userdata_order_of/
---
I have tried searching here, and well, all over and have likely missed what to do. 


The dilemma is this: from what I'm experiencing, ebs volume attachments using the AWS::EC2::Volume construct appear to be attached after userdata executes. If I spin in a until loop waiting for them to exist it'll spin until timeout and never complete.  If I don't wait, userdata executes and complains about the volumes not existing. 


Now, part of the issue here is that all Nitrobased instances e.g. the r5 series now map all EBS volumes as NVME drives. That means any Device e.g. /dev/xvdb you assign in attaching the volume is reduced in value unless you use ebsnvme-id to link the current mapping of xvd(x) to nvmeNn1. So that's fine, I'm pulling that down for my RHEL systems and would expect to be able to get the nvme device by specifying known /dev/xvd(x) devices to map to the one I want. 


This is great in theory, but because the volumes don't seem to be attached until after userdata completes, I can't do it. I need to attach several ebs volumes for use with Oracle ASM on Linux, or simple volumes in Windows for SQLServer. I could use blockdevicemapping rather than separate volume attachments, but that makes adding additional devices less hot add and more having to update the stack requiring an instance reboot. Also, you can't tag within the BlockDeviceMapping construct under 'AWS::EC2::Instance'. I know you can tag items after the fact, so that might be what I have to do with the awscli to get around this, but am hoping someone can point out something simple that's eluding me. 


Anyone? Please feel free to call me an idiot and point me to the simple solution that for the life of me I can't seem to find. 

TIA
## [7][Part 7 Building an Imgur clone in Vue.js and Serverless - Uploading Images To Our App - Now Live!](https://www.reddit.com/r/aws/comments/ewoope/part_7_building_an_imgur_clone_in_vuejs_and/)
- url: https://tutorialedge.net/projects/building-imgur-clone-vuejs-nodejs/part-7-uploading-images/
---

## [8][VPN and Private/Public secure isolation strategy. Plus, transit gateway!](https://www.reddit.com/r/aws/comments/eww1sw/vpn_and_privatepublic_secure_isolation_strategy/)
- url: https://www.reddit.com/r/aws/comments/eww1sw/vpn_and_privatepublic_secure_isolation_strategy/
---
Hello All,

I'm POC-ing transit gateway right now, and I'm having a bit of trouble understanding the concept of public/private subnets and routes when applied to the OpenVPN server I've put into my Transit VPC.  It comes down to this:

&amp;#x200B;

VPN server must be public, otherwise we can't hit it.

Public subnets are attached to public route tables, which route 0.0/0 to the IGW, and VPC /16 locally but nothing else.  This is by design, and is considered best security practice.

Private subnets have all the fun, and can go pretty much everywhere, including back to the on-premise network over the S2S VPN.

&amp;#x200B;

How do I present a VPN server to our users over the public internet, and allow them to get to private resources without expanding the public route table in a way that may be considered insecure?

One ring out from the scenario I've described, our users will need to hit the Transit Gateway.  This will allow them access to the (mostly private) resources in the other Accounts/VPCs/Environments.  How do I further expand the users' access to these resources, without stomping all over my security boundaries?  Is this  even possible?
## [9][what is the access policy for Cognito Authentication with ElasticSearch Kibana ?](https://www.reddit.com/r/aws/comments/ews909/what_is_the_access_policy_for_cognito/)
- url: https://i.imgur.com/zUNC1nN.png
---

## [10][Is AWS S3 static hosting suitable for a small business?](https://www.reddit.com/r/aws/comments/ewq5k6/is_aws_s3_static_hosting_suitable_for_a_small/)
- url: https://www.reddit.com/r/aws/comments/ewq5k6/is_aws_s3_static_hosting_suitable_for_a_small/
---
I work for a small start up which builds websites for small businesses. Our sites use an npm package to render the pages differently based on the users device type. Currently we don’t have clients that have server needs.

Would an S3 bucket be able to host a React application which has routes and npm dependencies?
