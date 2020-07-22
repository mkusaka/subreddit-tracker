# aws
## [1][Introducing The CIS Amazon EKS Benchmark | Amazon Web Services](https://www.reddit.com/r/aws/comments/hvlj36/introducing_the_cis_amazon_eks_benchmark_amazon/)
- url: https://aws.amazon.com/blogs/containers/introducing-cis-amazon-eks-benchmark/
---

## [2][Amazon EC2 Spot Instances integrations public roadmap](https://www.reddit.com/r/aws/comments/hvg3ja/amazon_ec2_spot_instances_integrations_public/)
- url: https://www.reddit.com/r/aws/comments/hvg3ja/amazon_ec2_spot_instances_integrations_public/
---
We just launched the [Amazon EC2 Spot Instances integrations public roadmap](https://github.com/aws/ec2-spot-instances-integrations-roadmap)! Tell us what open source software and frameworks you'd like us to integrate EC2 Spot Instances with!

[https://github.com/aws/ec2-spot-instances-integrations-roadmap/projects/1](https://github.com/aws/ec2-spot-instances-integrations-roadmap/projects/1)
## [3][Network Chaos Testing?](https://www.reddit.com/r/aws/comments/hvtnqy/network_chaos_testing/)
- url: https://www.reddit.com/r/aws/comments/hvtnqy/network_chaos_testing/
---
Hello everyone,

Does something like this exist? It seems like most of the chaos tools I've been seeing mostly target architectures. Like: ECS, EC2, Lambda, EKS, Cloudwatch, Elasticache, ELB, RDS, IAM, etc. [Gremlin/ChaosIQ/Toolkit]

But nothing on ALB, VPC, Route53, etc.

Or is this one of those things where we have to test manually? Or can Lambda handle account wide tasks for network chaos testing?
## [4][AWS Batch - why Job Queue can have multiple Compute Environments?](https://www.reddit.com/r/aws/comments/hvprgt/aws_batch_why_job_queue_can_have_multiple_compute/)
- url: https://www.reddit.com/r/aws/comments/hvprgt/aws_batch_why_job_queue_can_have_multiple_compute/
---
I am looking for concrete examples why Job Queue could have more than one Compute Environment specified.

At first I thought "maybe I have multiple compute environments with different instance types and I let the queue know which order is preferred". However, one Compute Environment itself can have multiple types of instances specified and with proper `allocationStrategy` on a `MANAGED` cluster will automatically provision a suitable instance(in theory at least). That fact ruled out my initial idea.

So far I thought about following use cases:

* run tasks on SPOT instances, if none are available then fallback to compute environment with on demand ones
* use my own ECS Cluster but if no compute resources are available fallback to Batch managed cluster
* use environment with compute optimised instances but if none are available fallback to some generic, cheaper ones defined in other compute environment

Why else could that feature be useful?
## [5][create a workspaces usage report with cloudwatch events](https://www.reddit.com/r/aws/comments/hvnww0/create_a_workspaces_usage_report_with_cloudwatch/)
- url: https://www.reddit.com/r/aws/comments/hvnww0/create_a_workspaces_usage_report_with_cloudwatch/
---
so i want to create a usage report to track the use of workspaces and there's documentation mentioning to use a cloudwatch event with workspaces as a source event. however the article does not explain how to configure the target

i selected cloudwatch log group as a target but i don't see any log group there. I even tried creating one. i am using an admin account so i don't think it's a permission issue. 

what am i doing wrong? i don't want to overcomplicate things by using lambda
## [6][How would you do continuous archiving on a 30TB windows file share](https://www.reddit.com/r/aws/comments/hvgzy3/how_would_you_do_continuous_archiving_on_a_30tb/)
- url: https://www.reddit.com/r/aws/comments/hvgzy3/how_would_you_do_continuous_archiving_on_a_30tb/
---
We have an analytics company with a quickly growing 30TB windows file share. Most of the data is automatically created in the share (by an on-prem app, that must stay on-prem) and NEVER accessed. There is possibility that any of it may NEED to be accessed at any time. So the data needs to be kept indefinitely. The on-prem storage is quickly running out of space and we're working on a solution to archive data to AWS.

The goal would be to move any data older than 30 days to S3 glacier. 

We're unsure of the best way to accomplish this is an automated and reliable fashion.

3rd party solutions that I've experience with COPY/Backup the data, requiring us to manually delete it from on-premise.

Anyone have any ideas for a use case like this? Otherwise we'll have to look at manually archiving data on a periodic basis.
## [7][Database solution for small serverless website?](https://www.reddit.com/r/aws/comments/hvh6ht/database_solution_for_small_serverless_website/)
- url: https://www.reddit.com/r/aws/comments/hvh6ht/database_solution_for_small_serverless_website/
---
Aurora serverless has a ~30 second startup time from paused. What is a cheap solution to having a serverless website with light database use that won't break the moment that there is more than 1 user?
## [8][Where are the AWS Reference Architecture Diagrams?!](https://www.reddit.com/r/aws/comments/hv38ac/where_are_the_aws_reference_architecture_diagrams/)
- url: https://www.reddit.com/r/aws/comments/hv38ac/where_are_the_aws_reference_architecture_diagrams/
---
"No documents found that match the selected criteria." is what I see when looking for *AWS Reference Architecture Diagrams* upon https://aws.amazon.com/architecture

I'm looking for https://aws.amazon.com/architecture/icons/ in the shape of basic Well Architected architecture diagram. You know something showing:

1) Account for IAM
2) Account for Audit/Logs with CloudTrail enabled
3) Maybe a staging, but certainly a production workload account
4) Icing on the cake is Guard duty etc setup in accordance to https://d1.awsstatic.com/whitepapers/Security/AWS_Security_Checklist.pdf

Are there such diagrams out there I can refer to?

Thank you in advance!
## [9][Amazon EBS Fast Snapshot Restore for Shared EBS Snapshots](https://www.reddit.com/r/aws/comments/hve8b0/amazon_ebs_fast_snapshot_restore_for_shared_ebs/)
- url: https://aws.amazon.com/blogs/aws/amazon-ebs-fast-snapshot-restore-for-shared-ebs-snapshots/
---

## [10][How do you access your private subnets in your VPC?](https://www.reddit.com/r/aws/comments/hvby7x/how_do_you_access_your_private_subnets_in_your_vpc/)
- url: https://www.reddit.com/r/aws/comments/hvby7x/how_do_you_access_your_private_subnets_in_your_vpc/
---
I'm building my small startup infrastructure in AWS and am trying to solve the access problem. Everything is internal to the VPC with NAT Gateways for internet access. (We have a single public facing load balancer). I need to access things like the Kubernetes control plane, ssh to ec2, load internal dashboards, etc.

&amp;#x200B;

Previously I've always solved this with OpenVPN on a small EC2 with a public IP, but I'm curious if there's something better? I've seen more and more about people using a Bastion and dynamically opening/closing the security group as needed. But this still makes things like dashboards hard. I was looking into AWS SSM Session Manager, but no support for macOS, so that can't work. Interested to hear what others are doing!

&amp;#x200B;

EDIT: I was incorrect here, You can use Session Manager on macOS. You just can't install the \_agent\_ on it. So It looks like I'm going to try our Session Manager, looks really promising! Thanks
