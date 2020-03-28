# aws
## [1][Is there a need for more CPUs if its utilisation doesn’t hit 100%?](https://www.reddit.com/r/aws/comments/fqf7w1/is_there_a_need_for_more_cpus_if_its_utilisation/)
- url: https://www.reddit.com/r/aws/comments/fqf7w1/is_there_a_need_for_more_cpus_if_its_utilisation/
---
I’ve a small service running on Fargate. At the moment, I’ve set it on 0.25 vCPU and 0.5gb of memory. 

So far, I occasionally see the CPU utilisation hitting 100% but falls back to around 40-60% most of the time. Since the average CPU utilisation usage seldom hits 100% for long period of time, does this then mean that 0.25 vCPU is sufficient in my case and increasing the vCPU is not going to do any difference to the service?

The reason I’m asking is I wanted to know whether I need to consider situations where I may still need to increase the number of vCPU (or CPUs) to improve the performance of an app even if the CPU utilisation in Cloudwatch appears to be utilising less than 100% most of the time.
## [2][How To Connect To ElastiCache Server From Node.js Express Application?](https://www.reddit.com/r/aws/comments/fqehjq/how_to_connect_to_elasticache_server_from_nodejs/)
- url: https://www.reddit.com/r/aws/comments/fqehjq/how_to_connect_to_elasticache_server_from_nodejs/
---
I set up a basic memcached elasticache cluster and I am trying to now connect to it through express-session in node.js (using the connect-memcached module). I have the following code:

&amp;#x200B;

https://preview.redd.it/ujhf0xfjrcp41.png?width=946&amp;format=png&amp;auto=webp&amp;s=fe2668d01efa66d3b79e1adcd90842870174e90d

On line 15, I am specifically providing the host server which is my cluster's configuration endpoint, but it keeps saying either the server is not available or timed out. Note I am running this application on my local machine, but I configured the security group to include all incoming traffic from anywhere. I'm not sure how i'm suppose to connect to the elasticache server. I tried looking everywhere but I can only find a couple of really complicated examples.
## [3][Serverless Components (as in the Serverless Framework)](https://www.reddit.com/r/aws/comments/fq5nhb/serverless_components_as_in_the_serverless/)
- url: https://www.reddit.com/r/aws/comments/fq5nhb/serverless_components_as_in_the_serverless/
---
Anyone using Serverless Components? As in [here](https://github.com/serverless/components)?

I really like the declarative nature of it, but it is still in beta and it shows a fair bit. More documentation is definitely needed as I'm looking over the source code way more than I should be. If you can find a component or a example template that fits your needs, Bob's your uncle. If not, the learning curve is pretty steep because there isn't much transfer from the Serverless Framework.

Just wondering if it is catching on with anyone. I really like it, or what I imagine it will be after another couple of releases, but sometimes I'm the only one laughing in a packed theatre.
## [4][How can I discover what kind of output cloudwatch event rule produces?](https://www.reddit.com/r/aws/comments/fqjvdt/how_can_i_discover_what_kind_of_output_cloudwatch/)
- url: https://www.reddit.com/r/aws/comments/fqjvdt/how_can_i_discover_what_kind_of_output_cloudwatch/
---
Lets say I have lambda functions configured to listen to cloudwatch event rule. Is there some place where I can find what kind structure the event will produce?

I know I can print it in lambda itself but it still requires the event rule to run. In some cases it's simple to make it happen but what if I have something that isn't easily reproducible like ASG failure.
## [5][How to handle schema migrations with Aurora Serverless?](https://www.reddit.com/r/aws/comments/fqjtfq/how_to_handle_schema_migrations_with_aurora/)
- url: https://www.reddit.com/r/aws/comments/fqjtfq/how_to_handle_schema_migrations_with_aurora/
---
Hi,

I’m thinking about using Aurora Serverless through the Data API from Lambda functions. One thing I’m concerned about is if it’s possible to handle schema migrations automatically. Are there any tools or best practices for this?

Thanks!
## [6][How to upload a file to all EC2 instances without using SCP/SSH?](https://www.reddit.com/r/aws/comments/fqiqyo/how_to_upload_a_file_to_all_ec2_instances_without/)
- url: https://www.reddit.com/r/aws/comments/fqiqyo/how_to_upload_a_file_to_all_ec2_instances_without/
---
Hi! I have an application running on two EC2 instances and I would like to automate config change deployments for it. The process is very simple, I just need to build a config file in GitLab CI, upload it to a particular path within my EC2 instances, and then run a command to restart my application so that the new config is picked up.

What would be the proper way of copying my config file to all EC2 instances? I want to avoid manually SCPing to them, as that would require hardcoding the instance URLs. I imagine the right tool is SSM, but I haven't found a way to do this with it.
## [7][Issue with Cloud Map](https://www.reddit.com/r/aws/comments/fqe2a4/issue_with_cloud_map/)
- url: https://www.reddit.com/r/aws/comments/fqe2a4/issue_with_cloud_map/
---
 have a Private Load Balancer ( ALB) DNS url which I can open in private subnet. I want to map this in Cloud Map.

I did this following configuration in Cloud Map

namespace = mynamespace

service name = myservice

CNAME = &lt;Private ALB DNS url &gt;

I can open http://&lt;Private ALB DNS url &gt;:8080/swagger-ui.html in private subnet

But I am unable to open it using Cloud Map url in private subnet.

This Cloud Map url

http://myservice.mynamespace:8080/swagger-ui.html does not open in private subnet.

What is wrong ? Is there any fix to this ?
## [8][AZ down in us-east-1?](https://www.reddit.com/r/aws/comments/fq439y/az_down_in_useast1/)
- url: https://www.reddit.com/r/aws/comments/fq439y/az_down_in_useast1/
---
Anyone experiencing issues with AZ's in us-east-1?   
Our Lambda's that are monitoring connectivity from every AZ in every account all went on at once.   
I don't see anything on AWS health page.
## [9][New AZ in ca-central-1 (Canada) region - skipped C?](https://www.reddit.com/r/aws/comments/fq1my5/new_az_in_cacentral1_canada_region_skipped_c/)
- url: https://www.reddit.com/r/aws/comments/fq1my5/new_az_in_cacentral1_canada_region_skipped_c/
---
It looks like there is a new AZ in the Canada region - ca-central-1-d.  What's odd is that this is the 3rd AZ in the region but AWS looks to have skipped "C".  I checked all the other regions and it doesn't look like this is the norm

Causing me a bad day because I have some terrafrom that dynamically creates subnets in VPCs based on number of AZs and assumed they are lettered sequentially.  Anyone know if it's a new normal that AZs are no longer sequentially lettered?
## [10][Question regarding T3.large (burstable) and C5.Large](https://www.reddit.com/r/aws/comments/fq4xhv/question_regarding_t3large_burstable_and_c5large/)
- url: https://www.reddit.com/r/aws/comments/fq4xhv/question_regarding_t3large_burstable_and_c5large/
---
T3.Large and C5.Large both have two vCPU.   Besides a differnce in RAM, is the other difference that T3.large is burstable CPU while C5.Large has a consistent CPU available (no burstable)?

Also, if I set my T3 instance to T3 unlimited is that basically the same as a C5.large (except for cost)?
