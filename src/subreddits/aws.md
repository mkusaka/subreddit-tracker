# aws
## [1][Policy Sentry: IAM Least Privilege Policy Generator, auditor, and analysis database](https://www.reddit.com/r/aws/comments/evbx18/policy_sentry_iam_least_privilege_policy/)
- url: https://github.com/salesforce/policy_sentry/
---

## [2][What exactly are Region and different types of Zones in AWS?](https://www.reddit.com/r/aws/comments/evm4kv/what_exactly_are_region_and_different_types_of/)
- url: https://medium.com/@zarinfam/what-exactly-are-region-and-different-types-of-zones-in-aws-333d7e0ce373?source=friends_link&amp;sk=f2c81db6cd87ac0b63ccc3c99d615e2b
---

## [3][AWS Lambda maximum bandwidth 0.5 gbps](https://www.reddit.com/r/aws/comments/ev9u8f/aws_lambda_maximum_bandwidth_05_gbps/)
- url: https://www.reddit.com/r/aws/comments/ev9u8f/aws_lambda_maximum_bandwidth_05_gbps/
---
With 3008mb memory allocated for the Lambda function the maximum bandwidth is 0.5 gbps (or read or write at \~75megabytes/second) to S3.  Have you all encountered this same limit?

It does not appear to be documented at:

[https://docs.aws.amazon.com/lambda/latest/dg/limits.html](https://docs.aws.amazon.com/lambda/latest/dg/limits.html)
## [4][[HELP] New EC2 bootup. how to enable PasswordAuthentication in sshd_config after booting up?](https://www.reddit.com/r/aws/comments/evl052/help_new_ec2_bootup_how_to_enable/)
- url: https://www.reddit.com/r/aws/comments/evl052/help_new_ec2_bootup_how_to_enable/
---
im setting up a auto scaling group from an AMI template. i tried putting the commands during AS setup.. didnt work, i also put a @reboot cron but it didnt work.

the template already has PasswordAuthentication yes but for some reason the AMI image went back to 'no'

i just need this to be run:    
sed -i 's/PasswordAuthentication no/PasswordAuthentication yes/g' /etc/ssh/sshd_config &amp;&amp; service sshd restart

EDIT: was solved by u/jamsan920 . please consider this closed. thanks
## [5][Multiple Interfaces causing asymmetric routing on ec2 instance](https://www.reddit.com/r/aws/comments/evjrzn/multiple_interfaces_causing_asymmetric_routing_on/)
- url: https://www.reddit.com/r/aws/comments/evjrzn/multiple_interfaces_causing_asymmetric_routing_on/
---
Hi,

I am having two interfaces attached to ec2 ubuntu instance where one interface having public ip and other interface only having private ip and i could see its causing routing loops.Can some one guide me on this one ?
## [6][Elastic beanstalk instances can't be placed in private subnets (without NAT gateway) ?](https://www.reddit.com/r/aws/comments/evhbwx/elastic_beanstalk_instances_cant_be_placed_in/)
- url: https://www.reddit.com/r/aws/comments/evhbwx/elastic_beanstalk_instances_cant_be_placed_in/
---
I spent the entire an day trying to deploy my flask based API engine on elastic beanstalk and figured that if I deploy it in a custom vpc where I keep the load balancer in public subnet (with internet gateway) and the compute instances in the private subnet (without Nat gateway) the deployment fails with the message - "The EC2 instances failed to communicate with AWS Elastic Beanstalk, either because of configuration problems with the VPC or a failed EC2 instance. Check your VPC configuration and try launching the environment again."
But then I deployed everything in a public subnet and it worked. 

So, I just wanted to confirm the behavior.

Is it that the requests coming to the API engines are routed through the load balancer but the actual instances behind the load balancer respond to individual requests (which they received) themselves ? 
(this should be the only reason they need internet connectivity as far as I can tell).
## [7][Is anyone here using AWS Batch for their projects?](https://www.reddit.com/r/aws/comments/ev51fa/is_anyone_here_using_aws_batch_for_their_projects/)
- url: https://www.reddit.com/r/aws/comments/ev51fa/is_anyone_here_using_aws_batch_for_their_projects/
---
Please share your experience with it so far if you're using it. Like Pros, Cons, etc. Thanks
## [8][Feedback request - getting started experience with AWS products](https://www.reddit.com/r/aws/comments/ev8yea/feedback_request_getting_started_experience_with/)
- url: https://www.reddit.com/r/aws/comments/ev8yea/feedback_request_getting_started_experience_with/
---
Hi r/aws!

u/amazonwebservices here. We are exploring  ways to improve the experience of getting started with AWS products.

Several approaches to lay out this content are in consideration, including learning pathways based on personas (e.g., full stack, DevOps, mobile) and / or product categories (e.g., serverless, containers).

* When you want to learn how to do something on AWS, is it typically in the lens of a specific use case you already have in mind or to generally go deeper on a particular product category? Or is there a different motivation altogether?
* How would you want to consume this information (e.g., hands-on tutorials, code snippets, videos)?
* What would make it easier and more engaging to get started with AWS?

Share your feedback below and help inform what the new getting started experience could look like!
## [9][Has anyone gotten CORS working with the new API Gateway V2 (HTTP API)?](https://www.reddit.com/r/aws/comments/evj9i5/has_anyone_gotten_cors_working_with_the_new_api/)
- url: https://www.reddit.com/r/aws/comments/evj9i5/has_anyone_gotten_cors_working_with_the_new_api/
---
According to here [https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html) "If you configure CORS for an API, API Gateway automatically sends a response to preflight OPTIONS requests, even if there isn't an OPTIONS route configured for your API."

However when I configure CORS the preflight responses just gives a 204 response with none of the necessary response headers (Access-Control-Allow-Origin) to get web browsers to accept it. I know something is happening when I configure CORS as without it I get 404 responses, but it's not giving the headers that I thought it would

Has anyone else gotten CORS working with the new HTTP API?
## [10][Multi Region and subnet schemes](https://www.reddit.com/r/aws/comments/evir6r/multi_region_and_subnet_schemes/)
- url: https://www.reddit.com/r/aws/comments/evir6r/multi_region_and_subnet_schemes/
---
Let's say I wanted to build out a second region for HA purposes. would it make sense to use a different cidr block? I don't foresee peering the VPCs but you never know. Does it hurt to use the same ip range if we don't peer.
