# aws
## [1][AWS service names - AWS vs Amazon. How do they make the distinction?](https://www.reddit.com/r/aws/comments/fvbkqa/aws_service_names_aws_vs_amazon_how_do_they_make/)
- url: https://www.reddit.com/r/aws/comments/fvbkqa/aws_service_names_aws_vs_amazon_how_do_they_make/
---
AWS makes distinction between services like Amazon DynamoDB vs AWS CloudTrail.

When I worked for AWS I did my speaker certification - an internal cert that allows one to speak on behalf of AWS. In that, they were very specific that one could not say "S3", but one had to say "Amazon S3". Now I've just realised that the homepage of the documentation ([https://docs.aws.amazon.com/](https://docs.aws.amazon.com/)) actually shows the various services prefixed by Amazon or AWS (wish I had figured this out while doing my speaker cert).

But I wonder how they make the distinction? 

Initially I thought it was "older services like S3 would have Amazon preceding them" - like Amazon S3 or Amazon DynamoDB, but I now see that DocumentDB is an "Amazon" service and not an "AWS" service and that's a new'ish service.

Anyone have any thoughts on this?
## [2][Project Ideas for Beginner DEs](https://www.reddit.com/r/aws/comments/fv5tdn/project_ideas_for_beginner_des/)
- url: https://www.reddit.com/r/aws/comments/fv5tdn/project_ideas_for_beginner_des/
---
I’m a data analyst who is striving to become a data engineer. What are some aws project ideas that I could start working on and gain some DE experience?
## [3][What security, networking and storage features are missing from AWS?](https://www.reddit.com/r/aws/comments/fv5f75/what_security_networking_and_storage_features_are/)
- url: https://www.reddit.com/r/aws/comments/fv5f75/what_security_networking_and_storage_features_are/
---
I have a few weeks to code out new OSS functionality for AWS. What security/network/storage functionality would make your life easier?
## [4][Tutorial: Standing up an EKS Cluster with Terraform](https://www.reddit.com/r/aws/comments/fuzgze/tutorial_standing_up_an_eks_cluster_with_terraform/)
- url: https://www.reddit.com/r/aws/comments/fuzgze/tutorial_standing_up_an_eks_cluster_with_terraform/
---
I've been playing with EKS lately and wanted to ensure I define my IaC so I wrote the following article on how to do this with terraform. Hopefully its helpful! [https://link.medium.com/0doYsmoXp5](https://link.medium.com/0doYsmoXp5)
## [5][Powershell New-EC2Tag](https://www.reddit.com/r/aws/comments/fuw6mi/powershell_newec2tag/)
- url: https://www.reddit.com/r/aws/comments/fuw6mi/powershell_newec2tag/
---
Can't take credit for the script:   [https://www.yobyot.com/aws/tag-aws-ec2-ebs-volumes-with-the-instance-name-tag/2017/02/05/](https://www.yobyot.com/aws/tag-aws-ec2-ebs-volumes-with-the-instance-name-tag/2017/02/05/) 

But I'm looking at running the following:

    Get-EC2Instance -Filter @(@{name = 'tag:MigrationPhase'; values = '432020' }) -ProfileName prd | ` # Get EC2 instances and pass to pipeline
    ForEach-Object -Process {
        # Get the name tag of the current instance ID; Amazon.EC2.Model.Tag is in the Instances object
        $instanceName = $_.Tags | Where-Object -Property Key -EQ "Name" | Select-Object -ExpandProperty Value
        $_.BlockDeviceMappings | ` # Pass all the current block device objects down the pipeline
        ForEach-Object -Process {
            $volumeid = $_.ebs.volumeid # Retrieve current volume id for this BDM in the current instance
            # Get the current volume's Name tag
            $volumeNameTag = Get-EC2Tag -Filter @(@{ name = 'tag:Name'; values = "*" }; @{ name = "resource-type"; values = "volume" }; @{ name = "resource-id"; values = $volumeid }) -ProfileName prd | Select-Object -ExpandProperty Value
            
            if (-not $volumeNameTag) # Replace the tag in the volume if it is blank
            {
                New-EC2Tag -Resources $volumeid -ProfileName prd -Tags @{ Key = "Name"; Value = $instanceName }  # Add volume name tag that matches InstanceID
            }
        }
    }

The problem I'm running into is, each time New-EC2Tag runs, I'm left with the following exception:

&gt;New-EC2Tag : The request must contain the parameter resourceIdSet  
&gt;  
&gt;At line:13 char:13  
&gt;  
&gt;\+             New-EC2Tag -Resources $volumeid -ProfileName prd -Tags @{ ...  
&gt;  
&gt;\+             \~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~\~  
&gt;  
&gt;\+ CategoryInfo          : InvalidOperation: (Amazon.PowerShe...NewEC2TagCmdlet:NewEC2TagCmdlet) \[New-EC2Tag\], InvalidOperationException  
&gt;  
&gt;\+ FullyQualifiedErrorId : Amazon.EC2.AmazonEC2Exception,Amazon.PowerShell.Cmdlets.EC2.NewEC2TagCmdlet

Curious if anyone has experience with this particular exception and what it is I might be missing.

  
Thanks
## [6][Postgres aurora major version upgrade](https://www.reddit.com/r/aws/comments/fuzapb/postgres_aurora_major_version_upgrade/)
- url: https://www.reddit.com/r/aws/comments/fuzapb/postgres_aurora_major_version_upgrade/
---
Anyone know if this is coming soon? Only upgrade path is manual pg_dump and at this point might as well move to RDS. 

Sadly, have lost faith in postgres aurora at this point.
## [7][Automated Ingesting help needed](https://www.reddit.com/r/aws/comments/fv1dde/automated_ingesting_help_needed/)
- url: https://www.reddit.com/r/aws/comments/fv1dde/automated_ingesting_help_needed/
---
Hi All,

**Problem**

I am looking to build something on AWS which will allow users to login (using something like AWS AD) and upload files. These files need to then be tagged or put into a folder which is associated to their account and stored somewhere where they can then be downloaded by a server (onsite), virus checked and moved through to the users home folder on our intranet. We do not want incoming direct connections to our DMZ which is why we are looking to cloud so only external connections are being made.

**Current Idea**

My current idea is to have a Nextcloud EC2 instance to provide a nice web front end, with single sign on connected to an AWS AD. Nextcloud should (if my thinking is correct), then connect to something like an S3 bucket with each user getting there own folder to store files.

A server should then be able look at that S3 bucket for changes, and download them, knowing from the file name the final location.

**Need Help On**

1. Am I going about this all wrong, is there a better way to do this?
2. Can you see any flaws in my current plan?
3. Any helpful articles or guides to do what I would like to do?

Many Thanks
## [8][Howto use both NACL and SG correctly?](https://www.reddit.com/r/aws/comments/fv3d0r/howto_use_both_nacl_and_sg_correctly/)
- url: https://www.reddit.com/r/aws/comments/fv3d0r/howto_use_both_nacl_and_sg_correctly/
---
I want to enable both nacl and sg for an ec2 instance that need to open port 80, 443 and 22 to internet access and it can do update software only. 

1. As my understanding sg still have entries of inbound that allow http, https and tcp/22 from 0.0.0.0 and outbound use default - allow all, is that correct?

2. With NACL, I don’t have much knowledge, so I guess I’ll need to allow http, https, tcp/22 AND port above 32???-32??? from 0.0.0.0 for inbound and for outbound I’ll need to allow port 80, 443 (??? means a range of ports that need to open for traffic from updates/downloads software when it outgoing on port 80/443) — what is for the range 32xxx? I’ve just remembered it but not quite sure. So is this logic correct?
## [9][Best way to share connection/configuration info among different resources](https://www.reddit.com/r/aws/comments/fuxqr9/best_way_to_share_connectionconfiguration_info/)
- url: https://www.reddit.com/r/aws/comments/fuxqr9/best_way_to_share_connectionconfiguration_info/
---
I have applications running in Lambda, Docker images in Batch, EC2, etc. I'm wondering what the best way is to share connection info, like RDS hosts/user/password. I've seen Parameter Store and Secrets Manager, but I'm not sure if this is the proper way to do it? Sometimes I think I've found a solution in AWS, and later I find out there's a totally other thing I've never heard of that's more appropriate :)
## [10][Route 53 subdomain pointing to EC2 fails](https://www.reddit.com/r/aws/comments/fv3akj/route_53_subdomain_pointing_to_ec2_fails/)
- url: https://www.reddit.com/r/aws/comments/fv3akj/route_53_subdomain_pointing_to_ec2_fails/
---
I set up a minecraft server on an EC2 instance last year. I created a subdomain record in my hosted zone and pointed the A record to the EC2 IP address. This was easy and worked well.

It's been a few months since I used the server, but I wanted to join today. I can a name resolution failure. Do an NS Lookup and a dig and nothing can resolve the hostname.

If I test the records in the Route 53 console it works. The server is up and I can connect via IP address.

Any ideas what to try next? I have deleted and recreated the record several times, and dropped the TTL too.
