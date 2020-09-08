# aws
## [1][Week of Sept 8th - What are you building this week on AWS?](https://www.reddit.com/r/aws/comments/iot7kv/week_of_sept_8th_what_are_you_building_this_week/)
- url: https://www.reddit.com/r/aws/comments/iot7kv/week_of_sept_8th_what_are_you_building_this_week/
---
Share your victories with the community
## [2][Suggestions - AWS Developer Forums shouldn't require login to read](https://www.reddit.com/r/aws/comments/iojm35/suggestions_aws_developer_forums_shouldnt_require/)
- url: https://www.reddit.com/r/aws/comments/iojm35/suggestions_aws_developer_forums_shouldnt_require/
---
Like the title says. It would improve the usability of forums if you could quickly review posts without needing to login. Stackoverflow is much more low friction. It would be great if I didn't have to pull out my 2FA device and login just to read a forum post.
## [3][A GitHub repo template for maintaining a multiple environment infrastructure with Terraform in AWS](https://www.reddit.com/r/aws/comments/ios8d6/a_github_repo_template_for_maintaining_a_multiple/)
- url: https://www.reddit.com/r/aws/comments/ios8d6/a_github_repo_template_for_maintaining_a_multiple/
---
[https://github.com/unfor19/terraform-multienv](https://github.com/unfor19/terraform-multienv) \- This template includes a CI/CD process, that applies the infrastructure in an AWS account.

[CI\/CD status](https://preview.redd.it/5hb4mf96ywl51.png?width=1258&amp;format=png&amp;auto=webp&amp;s=7128a30365d8a1e6346cf610a268d5cdd31d6038)
## [4][AWS Lambda &amp; AWS RDS](https://www.reddit.com/r/aws/comments/iorxhu/aws_lambda_aws_rds/)
- url: https://www.reddit.com/r/aws/comments/iorxhu/aws_lambda_aws_rds/
---
Hey guys.

I want to access my RDS database from Lambda. I'm currently using the default VPC for Lambda and wondering how to configure the network access between the RDS Database and Lambda.

  
I've read that I need to either :

* make my database openly accessible 
* create a VPC with a security group, add a NAT Gateway (which is expensive)

Is there a (free) midway situation? I'm building a PoC for my startup so I don't have much budget lol.
## [5][Mongodb hosted on AWS - cheaper than Mongodb Atlas cloud hosting?](https://www.reddit.com/r/aws/comments/iopnal/mongodb_hosted_on_aws_cheaper_than_mongodb_atlas/)
- url: https://www.reddit.com/r/aws/comments/iopnal/mongodb_hosted_on_aws_cheaper_than_mongodb_atlas/
---
Hi all, so....I am currently making use of Mongodb atlas to host some production/test/dev servers. The problem is I find this quite expensive, and there is no way to control the number of nodes (the default is 3, whatever level of machine you select). And though these are hosted on aws VIA Mongodb atlas, I feel like I could potentially get better pricing (and a more granularly suitable machine) if I just spun everything up on AWS myself.

Am I wrong? Has anyone done this themselves?
## [6][Deploying a Docker App on AWS ECS](https://www.reddit.com/r/aws/comments/io82x6/deploying_a_docker_app_on_aws_ecs/)
- url: https://www.reddit.com/r/aws/comments/io82x6/deploying_a_docker_app_on_aws_ecs/
---
Hi all,

I recently set out to learn how to deploy a Docker based app to AWS using ECR and ECS. I put together a brief tutorial on how to do it and some of the learnings I gathered.

The video is available here: https://youtu.be/zs3tyVgiBQQ

Hope some find this helpful.
## [7][How to add swap memory in auto scaling group of ec2 instances](https://www.reddit.com/r/aws/comments/ioshkd/how_to_add_swap_memory_in_auto_scaling_group_of/)
- url: https://www.reddit.com/r/aws/comments/ioshkd/how_to_add_swap_memory_in_auto_scaling_group_of/
---
We are configuring an auto-scaling-group. We need that for each spawned instance a swap memory should get allocated and it should be deleted once the instance auto-terminates due to scale in property of ASG. Please help to achieve that
## [8][What are the SCP best practices?](https://www.reddit.com/r/aws/comments/iohmeb/what_are_the_scp_best_practices/)
- url: https://www.reddit.com/r/aws/comments/iohmeb/what_are_the_scp_best_practices/
---
I lost the count of companies that I talked and have no idea what Service control polices can be used for. Once I explain I have the follow-up question that I don’t have answer yet. What should I set on my SCP?

This is a open question that can go from blocking unused regions to blocking IAM user creation, restrict to just a group to be allowed to delete resources/snapshot, etc. 

Usually I share this site for them to start. https://asecure.cloud 

What do you think it is a “must have” for any medium/small company that is worried about their security regarding SCP?
## [9][cPanel - Name Server on GURU Domains](https://www.reddit.com/r/aws/comments/ior3qp/cpanel_name_server_on_guru_domains/)
- url: https://www.reddit.com/r/aws/comments/ior3qp/cpanel_name_server_on_guru_domains/
---
Hi,

I have just launched an EC2 for cPanel and am configuring it. I am a bit stuck as I have not done it before. It has given me these as the default name servers:

ns1.eu-west-2.compute.internal

ns2.eu-west-2.compute.internal

Do I just have to put these in my hosting provider (GURU)? I tried these values and its says the NS does not exist...

Do I have to do any additional steps in AWS to point my Elastic IP to the cPanel?

Please can somebody help, I aim to move 9 websites onto this instance once I have learnt it.
## [10][I want to have a powerful virtual Windows machine ready to spin up and experiment with RTX raytracing in Unreal Engine. Can AWS do this in the cloud?](https://www.reddit.com/r/aws/comments/ioplsj/i_want_to_have_a_powerful_virtual_windows_machine/)
- url: https://www.reddit.com/r/aws/comments/ioplsj/i_want_to_have_a_powerful_virtual_windows_machine/
---
Where should I start? I want to be able to connect to cloud storage also and, possibly, spin up many machines at once.

I'm not a sysAdmin, I'm a 3D artist - so probably something .... Simple-ish.
## [11][Best VPN setup to access EC2](https://www.reddit.com/r/aws/comments/iokll2/best_vpn_setup_to_access_ec2/)
- url: https://www.reddit.com/r/aws/comments/iokll2/best_vpn_setup_to_access_ec2/
---
We currently run a web dashboard on EC2, and for security have SSH access only available through a University VPN. We are bringing on some outside help, and it's impossible to get them on the University VPN, so looking at other solutions. AWS VPN looks like more of a site to site, hardware based solution, so AWS openVPN looks like the best option, but am I correct in thinking it's going to be a nightmare to setup? I have to run it on it's own EC2 instance? Am I missing a better option? If not, are there any good walk-throughs  of setting up openVPN on AWS?
