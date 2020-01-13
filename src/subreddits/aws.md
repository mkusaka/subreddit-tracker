# aws
## [1][Path routing to Fargate tasks](https://www.reddit.com/r/aws/comments/eo46j7/path_routing_to_fargate_tasks/)
- url: https://www.reddit.com/r/aws/comments/eo46j7/path_routing_to_fargate_tasks/
---
Hi guys, I have a few doubts regarding networks, routing and Fargate that you fellas might have encountered before, or know how to circumvent :S.

&amp;#x200B;

Long story short, we are building a stack capable of handling, among other things, peaks of 200 plus Fargate tasks that need to be individually accessible using unique paths (1 path -&gt; 1 container/task).

&amp;#x200B;

My original idea was to go with an ALB and use the "path routing" feature it supports, sadly that's limited to 100 route groups (cannot be increased)... So not ideal.

&amp;#x200B;

With that limit in mind, I thought about setting 1 external ALB and 2 internal ones and chain them by using defaults, for a max of 300 route groups/paths, but testing that theory found out an ALB cannot point (as default at least) to another ALB... There goes that.

&amp;#x200B;

So I went to CloudFront, which we are using either way, and tried setting an Origin Group pointing to the first ALB, which uppon failure, would point to the next one and so on, but sadly origin groups have a limit of 2 origins, leaving me with 200 tasks at the most... Not sure if that's a fixed number or can be increased, it doesn't look like it judging by the UI.

&amp;#x200B;

How would you approach this? Is there a component for this?

&amp;#x200B;

&amp;#x200B;

Ps: As we are in the architecting/initial stages major changes can be done and work/complexity shouldn't be an issue.

&amp;#x200B;

Ps2: As a context, each task with both serve HLS to CF and have a small API to control both the container status and to perform certain operations.
## [2][Could I run League of Legends constantly on AWS?](https://www.reddit.com/r/aws/comments/eo4l4z/could_i_run_league_of_legends_constantly_on_aws/)
- url: https://www.reddit.com/r/aws/comments/eo4l4z/could_i_run_league_of_legends_constantly_on_aws/
---
Hey!

I've never used AWS and am starting a project to learn AWS Cloud Vmware.

I was thinking something that could be fun would be to run my favorite game League of Legends on AWS.

I wanted to ask if this is possible and if so, how would one go about it?

Thanks in advance :)
## [3][Building and scaling a micro-segmented enterprise network in AWS](https://www.reddit.com/r/aws/comments/enx5h6/building_and_scaling_a_microsegmented_enterprise/)
- url: https://medium.com/@matt_69071/building-and-scaling-a-micro-segmented-network-in-aws-for-regulated-industries-af4f931cd48c
---

## [4][AWS Border Protection - Is there a list of all AWS services/resources that can be configured to be "publicly" accessed?](https://www.reddit.com/r/aws/comments/ensmyk/aws_border_protection_is_there_a_list_of_all_aws/)
- url: https://www.reddit.com/r/aws/comments/ensmyk/aws_border_protection_is_there_a_list_of_all_aws/
---
Hi all -

There are obvious services that can be configured to be "publicly" accessible such as EC2 instances or S3 buckets; however, there are also some less known cases such as making an ECR repository public or publishing a public AMI.

*Does anyone know if there is a list maintained by AWS or by the community that lists all possible publicly facing endpoints for services/resources?* I'm looking at this from a cyber security perspective (i.e. border protection) with the use case of ensuring no services/resources from an AWS account can be accessed publicly. 

Thank you for any help in advance.

Best,

Andrew
## [5][Fargate Spot connection drainning?](https://www.reddit.com/r/aws/comments/eo4i21/fargate_spot_connection_drainning/)
- url: https://www.reddit.com/r/aws/comments/eo4i21/fargate_spot_connection_drainning/
---
When I was using EC2 spot, there was spot termination notice in metadata endpoint. All I had to do was to include check in my health check script. If it detected termination notice, it would report unhealthy and load balancer would drain the task.

I doesn't work the same way with Fargate spot. Notification is send to Event Bridge and ecs task receives stop signal which stop the task without draining.

Do you have any tip how to ensure draining on load balancer? One solution would be to highjack stop signal on all containers in task but that seems complicated and I am hopping there is a simpler solution.
## [6][Mapping multiple elastic IP addresses for nameservers on EC2](https://www.reddit.com/r/aws/comments/eo2tao/mapping_multiple_elastic_ip_addresses_for/)
- url: https://www.reddit.com/r/aws/comments/eo2tao/mapping_multiple_elastic_ip_addresses_for/
---
Is it possible to map multiple IP addresses to one EC2 server instance, for use as nameservers
## [7][Fargate load balancing &amp; routing](https://www.reddit.com/r/aws/comments/eo2mwn/fargate_load_balancing_routing/)
- url: https://www.reddit.com/r/aws/comments/eo2mwn/fargate_load_balancing_routing/
---
Hello,

I'm setting up my infra with Fargate (using AWS CDK scripting). As we have multiple applications running in the back-end, I'm currently investigating potential options for routing between applications. Below is an overview of the 3 options I currently have in mind. Do you have a strong opinion on which one I should go for, based on your experience ?

&amp;#x200B;

https://preview.redd.it/w2ya3o2l1ja41.png?width=911&amp;format=png&amp;auto=webp&amp;s=0bdcfdeff6ed45ca9751d8c1b8f2d1b1932be2f8

Thanks a ton in advance for the help!

Cheers
## [8][Semantic segmentation](https://www.reddit.com/r/aws/comments/eo24t6/semantic_segmentation/)
- url: https://www.reddit.com/r/aws/comments/eo24t6/semantic_segmentation/
---
Hello, is it possible to use Sagemaker's semantic segmentation model locally?  Because the only deployment in example notebook is Amazon SageMaker hosted endpoint. Reference: [https://github.com/awslabs/amazon-sagemaker-examples/blob/master/introduction\_to\_amazon\_algorithms/semantic\_segmentation\_pascalvoc/semantic\_segmentation\_pascalvoc.ipynb](https://github.com/awslabs/amazon-sagemaker-examples/blob/master/introduction_to_amazon_algorithms/semantic_segmentation_pascalvoc/semantic_segmentation_pascalvoc.ipynb). Thank you in advance!
## [9][amplify custom domain name](https://www.reddit.com/r/aws/comments/enzlq4/amplify_custom_domain_name/)
- url: https://www.reddit.com/r/aws/comments/enzlq4/amplify_custom_domain_name/
---
Hi,

I've started using amplify for an SPA with angular.

I'm trying to add my custom domain name with no luck. The domain name was bought/registered through AWS with Route53.

The error:

One or more of the CNAMEs you provided are already associated with a different resource.

What I've tried:

* I've tried recreating the app
* Checked cloudfront - it's empty
* Checked cloudformation for any old deployments
* Checked S3 for static websites
* I've tried adding another domain but now it's stuck "activating"

Any ideas?

Are there a better alternative to hosting a single page app in AWS?

Thank you!
## [10][What exactly are Region and different types of Zones in AWS? Describing Region, Availability Zone, Local Zone and Wavelength Zone](https://www.reddit.com/r/aws/comments/eo0kzi/what_exactly_are_region_and_different_types_of/)
- url: https://itnext.io/what-exactly-are-region-and-different-types-of-zones-in-aws-333d7e0ce373?source=friends_link&amp;sk=f2c81db6cd87ac0b63ccc3c99d615e2b
---

