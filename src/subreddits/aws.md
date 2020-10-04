# aws
## [1][AWS Lightsail Deep Dive: What is it and when to use it](https://www.reddit.com/r/aws/comments/j4s7y4/aws_lightsail_deep_dive_what_is_it_and_when_to/)
- url: https://www.learnaws.org/2020/09/24/aws-lightsail-deep-dive/
---

## [2][AWS CloudFront route traffic through private DNS name](https://www.reddit.com/r/aws/comments/j4whk6/aws_cloudfront_route_traffic_through_private_dns/)
- url: https://www.reddit.com/r/aws/comments/j4whk6/aws_cloudfront_route_traffic_through_private_dns/
---
About 10 years ago, I setup AWS Cloudfront origin pointed at my EC2 private DNS so that I can close off all outside traffic from the internet to my EC2 instance directly using security groups.  In other words, all traffic had to go through CloudFront.

I tried doing the same and it won't let me.

Does anybody know how I might able to get this to work?

Thanks in advance.
## [3][Aim for recovery from regional failure or zonal failure? TRICKY EXAM QUESTION](https://www.reddit.com/r/aws/comments/j4ohcp/aim_for_recovery_from_regional_failure_or_zonal/)
- url: https://www.reddit.com/r/aws/comments/j4ohcp/aim_for_recovery_from_regional_failure_or_zonal/
---
So since AWS wants us to build while keeping fault and failure tolerance in mind, should we do that on a regional scale or a zonal scale?   


for example i can duplicate my resources in different zones in the same region right? but then what if a disaster hit the whole region or a regional outage happened? in this case i should have built my infrastructure to rely on regional availability no? so incase of one region outage accord then at least i have a different region serving my users?  


The question in this picture is from jon bonsos practice test for CCP

  


https://preview.redd.it/nmpwf8n2myq51.png?width=770&amp;format=png&amp;auto=webp&amp;s=6744f06c85f980082377602313d496bc5e08502d
## [4][Can Transit Gateway ENIs be configured as Traffic Mirror sources?](https://www.reddit.com/r/aws/comments/j4vzw5/can_transit_gateway_enis_be_configured_as_traffic/)
- url: https://www.reddit.com/r/aws/comments/j4vzw5/can_transit_gateway_enis_be_configured_as_traffic/
---
As per title - there's an archived thread with the same question but no confirmation.

This seems like a much cleaner approach to traffic monitoring in our VPC than setting up each individual EC2 host/eni.
## [5][Question: AWS Amplify Graphql S3 Integration](https://www.reddit.com/r/aws/comments/j4r4lh/question_aws_amplify_graphql_s3_integration/)
- url: https://www.reddit.com/r/aws/comments/j4r4lh/question_aws_amplify_graphql_s3_integration/
---
Reading through the documentation I see this line:

"The GraphQL Transform handles creating the relevant input types and will store pointers to S3 objects in Amazon DynamoDB. The AppSync SDKs and Amplify library handle uploading the files to S3 transparently."

From what I understand, this means by running a graphql request with the described S3Object definition I should be able to do a mutation and by running that mutation upload the data to S3. However, that doesn't happen. I'm curious to know if this functionality exists and how to do it or if I am misinterpreting what's being said, thanks!

[https://docs.amplify.aws/cli/graphql-transformer/storage#basics](https://docs.amplify.aws/cli/graphql-transformer/storage#basics)
## [6][How Do I Get Graphics Output From My EC2 Instance To My Local Machine?](https://www.reddit.com/r/aws/comments/j4kmnq/how_do_i_get_graphics_output_from_my_ec2_instance/)
- url: https://www.reddit.com/r/aws/comments/j4kmnq/how_do_i_get_graphics_output_from_my_ec2_instance/
---
The EC2 instance is running the Ubuntu 18.04 Deep Learning AMI and my local machine is running Ubuntu 20.04. I have some Python scripts that output Matplotlib graphs and I would like to run the scripts on my EC2 instance and then see the output graphs on my local machine. Is there a good way to do this? Some googling turned up a bunch of things about X11 and Windows machines but I couldn't find a good resource concerning Linux local machines. Any help is appreciated. Thanks in advance.
## [7][AWS Amplify and scaling](https://www.reddit.com/r/aws/comments/j4tn0f/aws_amplify_and_scaling/)
- url: https://www.reddit.com/r/aws/comments/j4tn0f/aws_amplify_and_scaling/
---
I have deployed various pages in the past with AWS Amplify and enjoy developing with it but I can't for the life of me figure out if AWS Amplify supports some kind of auto scaling (for instance if the site were to skyrocket in usage) 

All my clients to date have had 100 or less users on their page at a given time so we never noticed anything or stress tested it.  What does AWS Amplify do in those cases though?
## [8][Utterly confused by Cloudwatch custom metrics GUI](https://www.reddit.com/r/aws/comments/j4nua5/utterly_confused_by_cloudwatch_custom_metrics_gui/)
- url: https://www.reddit.com/r/aws/comments/j4nua5/utterly_confused_by_cloudwatch_custom_metrics_gui/
---
Let's say I have an app that simulates a train, I want to collect metrics on the number of passengers, and also their names, I use commands such as this to add metrics:

`aws cloudwatch put-metric-data --namespace MyTrainApp --metric-name passengers --value 2
--unit Count --dimensions "passengers=abe;bob"`

However, over in the console, when I try to look at a graph of the passengers, the data is sorted up by the passengers list (`abe;bob;carl`, etc), so I need to manually select each unique passengers list in order to have Cloudwatch graph the whole thing.  This is very counter intuitive to me as I'd expect it'd be sorted by metric-name.

Is there something I can do to make it do what I expect?
## [9][S3 Update – Three New Security &amp; Access Control Features](https://www.reddit.com/r/aws/comments/j422w2/s3_update_three_new_security_access_control/)
- url: https://aws.amazon.com/blogs/aws/amazon-s3-update-three-new-security-access-control-features/
---

## [10][Cracking my head over proper routing for multiple API gateways &amp; cloudfronts](https://www.reddit.com/r/aws/comments/j4gezr/cracking_my_head_over_proper_routing_for_multiple/)
- url: https://www.reddit.com/r/aws/comments/j4gezr/cracking_my_head_over_proper_routing_for_multiple/
---
I've been struggling all day long how to best do routing for a new (non-technical) client. The project is small enough to use it as a test bench for our standardization.

They have a main website, hosted by a third part, on [website1.com](https://website.com) \+ a few sub websites also on different domains. They asked us to develop a chatbot for [website1.com](https://website.com) and since the third party could become a bottleneck in terms of DNS and support we decided to register [client.com](https://client.com) in route 53 and let the third party load it through an iframe.

So now I have [client.com](https://client.com), loaded through an iframe, serving the chatbot and [api.client.com](https://api.client.com) to act as the API gateway with websockets.

The client has said that they want a dashboard for the chatbot1 + expand into the other websites that they have with new bots. I think the best would be to expand the current API gateway to accommodate the dashboard API calls as well for chatbot1. Before we continue though I want to make sure that we have a proper plan in place for expanding, including DTAP (dev, test, acceptance, prod).

Taking the chatbot1 as an example, what would be the best course of action?

&amp;#x200B;

|frontend|API gateway|
|:-|:-|
|client.com/web/chatbot1 |client.com/api/chatbot1 |
|chatbot1.client.com |chatbot1.client.com/api |
|client.com/chatbot1 |api.client.com/chatbot1 |

&amp;#x200B;
