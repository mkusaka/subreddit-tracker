# aws
## [1][Announcing Amazon Relational Database Service (RDS) Snapshot Export to S3](https://www.reddit.com/r/aws/comments/etepdt/announcing_amazon_relational_database_service_rds/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/01/announcing-amazon-relational-database-service-snapshot-export-to-s3/
---

## [2][R53 service health issue](https://www.reddit.com/r/aws/comments/etgoet/r53_service_health_issue/)
- url: https://www.reddit.com/r/aws/comments/etgoet/r53_service_health_issue/
---
The service health statement is vague. I am seeing problems across many of our apps that utilize EC2, ECS, RDS, R53, along with vendor sites I know run on AWS. Waiting to chat with support. [https://status.aws.amazon.com](https://status.aws.amazon.com/)

EDIT: Problems cleared up for my RDS Aurora Mysql cluster connection timeouts when AWS posted that they found the root cause of R53 issue at 6:08pm ET. Frustrating they never include all affected services.
## [3][Questions on setup of online shopping website with aws.](https://www.reddit.com/r/aws/comments/etl3k4/questions_on_setup_of_online_shopping_website/)
- url: https://www.reddit.com/r/aws/comments/etl3k4/questions_on_setup_of_online_shopping_website/
---
I'm creating a shopping website, using angular for the frontend and java for the backend, and am a little overwhelmed with the following.

Basically I'm stuck on where should the java application live? 

- Should I put it in a docker container and run it on EC2? 

- Should I just run it on EC2 without docker? 

- Is there any benefit in me thinking about fargate or something along those lines?

- Should I forget about the EC2 option and just go with lambda?
## [4][Best practice to back up existing dedicated server data to S3](https://www.reddit.com/r/aws/comments/etmlp1/best_practice_to_back_up_existing_dedicated/)
- url: https://www.reddit.com/r/aws/comments/etmlp1/best_practice_to_back_up_existing_dedicated/
---
Hoping for advice from someone here who has experience (or at least an idea for) copying data from a dedicated server environment to S3.

We have a website with alot of video content currently managed via cPanel on our dedicated server. We would like to back up everything in the public\_html directory to an AWS S3 bucket, retaining the directory structure.

This would just be a backup - we would manage everything backed up on S3 via normal AWS tools (no cPanel) once the data is moved across.

Would appreciate any advice on the best way transfer/back up this data securely and directly without intermediate storage.

Thanks
## [5][AWS Transcribe Sensitivity](https://www.reddit.com/r/aws/comments/etm3ff/aws_transcribe_sensitivity/)
- url: https://www.reddit.com/r/aws/comments/etm3ff/aws_transcribe_sensitivity/
---
I want to make sure the transcription only works if the person speaks clearly and pronounces their words correctly. Is there a way to tell AWS Transcribe to turn up/down the sensitivity of how well it picks up words?
## [6][T3 Instances on Dedicated Single-Tenant Hardware](https://www.reddit.com/r/aws/comments/etd5rz/t3_instances_on_dedicated_singletenant_hardware/)
- url: https://aws.amazon.com/blogs/aws/new-t3-instances-on-dedicated-single-tenant-hardware/
---

## [7][AWS DataSync Update – Support for Amazon FSx for Windows File Server](https://www.reddit.com/r/aws/comments/etgt1k/aws_datasync_update_support_for_amazon_fsx_for/)
- url: https://aws.amazon.com/blogs/aws/aws-datasync-update-support-for-amazon-fsx-for-windows-file-server/
---

## [8][Original Snowball retirement, Snowball Edge only from 4th Feb](https://www.reddit.com/r/aws/comments/et8d99/original_snowball_retirement_snowball_edge_only/)
- url: https://www.reddit.com/r/aws/comments/et8d99/original_snowball_retirement_snowball_edge_only/
---
Anyone know any more about the original Snowball retirement?  I haven't been able to find any announcement of it or any more details, but the following notification appears in the AWS console:

&gt; On February 4th, 2020, the first-generation 48 TB and 80 TB Snowball devices will be retired from the Snowball Service. Devices ordered by February 3rd, 2020 will be fulfilled. These 48 TB and 80 TB devices will no longer be orderable on February 4th, 2020. The Snowball Edge Storage Optimized device replaces the first-generation Snowball device for data migration requirements.

This is a bit frustrating for us, as we found the Edge to be more awkward to use in our specific workflows, and we're not geared up to start working with Edges again, certainly not with a week and a half notice.
## [9][VPC private IP address change](https://www.reddit.com/r/aws/comments/etka4i/vpc_private_ip_address_change/)
- url: https://www.reddit.com/r/aws/comments/etka4i/vpc_private_ip_address_change/
---
i would like to change the private IP (IPv4) address of a VPC but *not* change the IPv6 address.  is there a way to do this?
## [10][AWS Transfer for SFTP or set up a local Linux SFTP EC2 Instance](https://www.reddit.com/r/aws/comments/etiila/aws_transfer_for_sftp_or_set_up_a_local_linux/)
- url: https://www.reddit.com/r/aws/comments/etiila/aws_transfer_for_sftp_or_set_up_a_local_linux/
---
Hi Everyone,

I have been given the task of setting up a SFTP server.    To be honest, I would rather setup an SFTP  Server myself on an EC2 Instance.

My Question is to inquire if anyone has ever made use of the SFTP service that AWS offers? What would be the pros and cons of it to against setting up ground up?

Thanks
