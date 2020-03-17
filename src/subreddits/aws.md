# aws
## [1][Locked myself out of my Lightsail instance - UPDATE](https://www.reddit.com/r/aws/comments/fk2hc2/locked_myself_out_of_my_lightsail_instance_update/)
- url: https://www.reddit.com/r/aws/comments/fk2hc2/locked_myself_out_of_my_lightsail_instance_update/
---
Considering I wasn't able to find any elegant solution online, I posted a question on this sub about a week ago looking for a way to regain SSH access to my Lightsail instance. [Original post](https://www.reddit.com/r/aws/comments/ff2r7u/i_got_locked_out_of_my_aws_lightsail_instance/)

You guys proposed transfering over to EC2 from my snapshot in order to externally disable the UFW and that seemed like it would work, but I wanted to minimize my expenses since the budget wasn't really high. I upgraded my support plan to create a technical case and a support representative came up with this solution:

&amp;#x200B;

1. Create a snapshot of your Lightsail instance
2. Select "Create instance from snapshot"
3. While configuring the new instance, locate the "Launch script" section and paste the following:

&amp;#x200B;

        --//
        Content-Type: text/cloud-config; charset="us-ascii"
        MIME-Version: 1.0
        Content-Transfer-Encoding: 7bit
        Content-Disposition: attachment; filename="cloud-config.txt"
        
        #cloud-config
        cloud_final_modules:
        - [scripts-user, always]
        
        --//
        Content-Type: text/x-shellscript; charset="us-ascii"
        MIME-Version: 1.0
        Content-Transfer-Encoding: 7bit
        Content-Disposition: attachment; filename="userdata.txt"
        
        #!/bin/bash
        sudo ufw disable
        --//

&amp;#x200B;

4. Detach your static IP from your old instance and attach to the new one 

5. SSH into your new instance

&amp;#x200B;

Hope this helps if anyone finds themselves in the same situation I did. Cheers.
## [2][Amazon ElastiCache Global Datastore for Redis](https://www.reddit.com/r/aws/comments/fjvaac/amazon_elasticache_global_datastore_for_redis/)
- url: https://aws.amazon.com/blogs/aws/now-available-amazon-elasticache-global-datastore-for-redis/
---

## [3][Tracking environmental impact of AWS usage?](https://www.reddit.com/r/aws/comments/fk0qn1/tracking_environmental_impact_of_aws_usage/)
- url: https://www.reddit.com/r/aws/comments/fk0qn1/tracking_environmental_impact_of_aws_usage/
---
Is there a good way to estimate carbon footprint/KwH or other consumption metrics for AWS services? I'd like to offset my usage and am struggling to find applicable data or resources.

If not, I'd be interested in spinning up a project to develop the tooling to do so. I imagine there are plenty of users and businesses that could use this info.
## [4][Is it possible to enforce a smaller max payload on API Gateway?](https://www.reddit.com/r/aws/comments/fk2wbd/is_it_possible_to_enforce_a_smaller_max_payload/)
- url: https://www.reddit.com/r/aws/comments/fk2wbd/is_it_possible_to_enforce_a_smaller_max_payload/
---
I'm building an API with API Gateway. The service it is exposing is inherently slow and cannot be made faster. The service is experimental so speed is not being prioritised while the proof of concept is being tested. Is it possible to reduce the max payload allowed by API Gateway to reduce load on the system. I can see in the docs that payload cannot be increased but can it be decreased?  [https://docs.aws.amazon.com/apigateway/latest/developerguide/limits.html](https://docs.aws.amazon.com/apigateway/latest/developerguide/limits.html)
## [5][Cloning Production RDS DB for Staging](https://www.reddit.com/r/aws/comments/fk4rka/cloning_production_rds_db_for_staging/)
- url: https://www.reddit.com/r/aws/comments/fk4rka/cloning_production_rds_db_for_staging/
---
I’m just getting started with RDS for my Postgres database and my question is, how would I be able to clone my production database so that I could connect my app to it and run test in a staging environment?
## [6][CloudFront Updates Are No Longer Soul Destroying - Adam Johnson](https://www.reddit.com/r/aws/comments/fjjdkg/cloudfront_updates_are_no_longer_soul_destroying/)
- url: https://adamj.eu/tech/2020/03/16/cloudfront-updates-no-longer-soul-destroying/
---

## [7][Can data stored in EBS Volume on the master node be replicated on the slave nodes?](https://www.reddit.com/r/aws/comments/fk0rvb/can_data_stored_in_ebs_volume_on_the_master_node/)
- url: https://www.reddit.com/r/aws/comments/fk0rvb/can_data_stored_in_ebs_volume_on_the_master_node/
---
I'm using Kubernetes with AWS. I've created a master and 2 slave nodes. The EBS volume path is attached to the master. I can see the volume directory on the slaves(/data). When I add data to that directory on my master node i can view it in the path set on my pods but the data doesn't get updated on my slave nodes. What I wanted to know is if its possible for the data to be replicated on my slaves and if so, then how do I do it?
## [8][Amazon WorkSpaces](https://www.reddit.com/r/aws/comments/fjxiur/amazon_workspaces/)
- url: https://www.reddit.com/r/aws/comments/fjxiur/amazon_workspaces/
---
I've been looking into Amazon WorkSpaces for a group project regarding a business's IT infrastructure. I understand that it is a desktop as a service, however, would the business still require on-site servers? Or would data and information be stored over the cloud and managed by AWS? Would it be possible to have private on-site servers? 

I've noticed a lot of Universities starting to use WorkSpaces, and I'm uncertain if they would be okay with just keeping all data in the cloud storage.
## [9][Does API gateway support multiple levels of wildcard domains](https://www.reddit.com/r/aws/comments/fjwpw4/does_api_gateway_support_multiple_levels_of/)
- url: https://www.reddit.com/r/aws/comments/fjwpw4/does_api_gateway_support_multiple_levels_of/
---
Hey all, I’m trying to setup and API with Lambda and API Gateway with a custom domain (cert with ACM). I’m wondering if it’s possible to setup multiple levels of wildcard domains as it doesn’t seem possible from my tests. By multiple levels I mean *.*.mydomain.com and *.*.*.mydomain.com.

I’m wondering if this is more so a certificate issue than API gateway since multiple levels of domains do find the API service properly, but I get a certificate error and the API gateway endpoints are forbidden
## [10][Amazon Athena now publishes CloudWatch Events for Athena query state transitions](https://www.reddit.com/r/aws/comments/fjp9bl/amazon_athena_now_publishes_cloudwatch_events_for/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/03/amazon-athena-now-publishes-cloudwatch-events-for-athena-query-state-transitions/
---

