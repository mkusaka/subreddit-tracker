# aws
## [1][Week of Sept 28th - What are your favorite AWS Tips?](https://www.reddit.com/r/aws/comments/j1dndd/week_of_sept_28th_what_are_your_favorite_aws_tips/)
- url: https://www.reddit.com/r/aws/comments/j1dndd/week_of_sept_28th_what_are_your_favorite_aws_tips/
---
Share your AWS Tips
## [2][FYI: CDK Day 2020 Coming on September 30th](https://www.reddit.com/r/aws/comments/j1lpto/fyi_cdk_day_2020_coming_on_september_30th/)
- url: https://www.cdkday.com/
---

## [3][Is there an Amazon WorkSpaces client for OpenSuse?](https://www.reddit.com/r/aws/comments/j1yfox/is_there_an_amazon_workspaces_client_for_opensuse/)
- url: https://www.reddit.com/r/aws/comments/j1yfox/is_there_an_amazon_workspaces_client_for_opensuse/
---
Hi, I have been trying to find an Amazon WorkSpaces client for OpenSuse but no package seems to be available anywhere. I know amazon only officially supports deb but I was hoping that there would be some repackaged client at a community repo, or snapcraft, or somewhere. Why wouldn't amazon support RDH and SUSE as well? It wouldn't be that much of a task for amazon to repackage it, even as non-supported client.
## [4][AWS IoT and device states](https://www.reddit.com/r/aws/comments/j1w6aw/aws_iot_and_device_states/)
- url: https://www.reddit.com/r/aws/comments/j1w6aw/aws_iot_and_device_states/
---
If the state of a thing changes (eg. temperature sensor value changes) I can use an IoT Rule to update that state change directly to DynamoDB, right?  
So I can keep track of thing state in DynamoDB and query their current sensor values with lambdas? 

How about Fleet Indexing, it seems that it is designed to do just that.   
Should I rather be using Fleet Indexing and use it to keep track of each thing and query their state and sensor values?
## [5][Best way to automate ECS setup + deployments?](https://www.reddit.com/r/aws/comments/j1z3ch/best_way_to_automate_ecs_setup_deployments/)
- url: https://www.reddit.com/r/aws/comments/j1z3ch/best_way_to_automate_ecs_setup_deployments/
---
What do people suggest as a good way forward when it comes to having ESC infrastructure as code and deployments (blue/green) automated?

I have set up a couple of clusters / tasks / services (via the console) as sandbox to learn ECS but am now at the stage where I would like to deploy apps, built by jenkins in an automated fashion. It would also be interesting to use aws Code Deploy.

At the moment my focus is on terraform for the cluster configuration, is terraform also a good option for setting up and operating code deploy? I see some tutorials that use cloudformation but i'm really not a fan and would rather avoid it if at all possible. Another  alternative i could consider is ansible

Edit: images are already stored in ECR
## [6][Does ALB remove the need to put a NGINX server in front of my app servers?](https://www.reddit.com/r/aws/comments/j1f6uj/does_alb_remove_the_need_to_put_a_nginx_server_in/)
- url: https://www.reddit.com/r/aws/comments/j1f6uj/does_alb_remove_the_need_to_put_a_nginx_server_in/
---
I have a server for chat that handles websocket connections and a server for the core of my application with is just a rest API, users make posts and stuff. I was planning to put and NGINX server in front of both, to balance the load and act as a reverse proxy. However, now I am thinking the ALB might handle that for me. Is this assumption correct?
## [7][SUSE SAP BYOL image](https://www.reddit.com/r/aws/comments/j1xlhs/suse_sap_byol_image/)
- url: https://www.reddit.com/r/aws/comments/j1xlhs/suse_sap_byol_image/
---
Hy guys,

We have a plan to migrate our existing SAP to AWS, because of our existing subscription. We want to BYOS option when migrate to AWS.

I tried to search for the procedures, but didn't get a clue for that.

Like

\- Which AMI should we use for BYOS

\- Which is better use BYOS or On-Demand SUSE subscription on AWS

\- How to migrate with minimum downtime. We get suggestion for using cloudendure, but we considering too to use fresh installation and using system copy.

&amp;#x200B;

Maybe someone can help me get some enlightenment? BYOS is so confusing for us.

&amp;#x200B;

Thank you
## [8][What is the best approach to convert XML to PDF via XSLT](https://www.reddit.com/r/aws/comments/j1xhk1/what_is_the_best_approach_to_convert_xml_to_pdf/)
- url: https://www.reddit.com/r/aws/comments/j1xhk1/what_is_the_best_approach_to_convert_xml_to_pdf/
---
We have a team in our project who are responsible to convert 100000+ XML files to PDF via an XSLT. Currently they are running the task on 4 desktops but are looking for a quicker turnaround.  
This made me think to use some service in AWS to do this but I have no idea what AWS service we could deploy. I was thinking Lambda could be a candidate but not sure about the allowed running time per call (some XML's are quite large, resulting in over 20MB PDF files) and the cost of calling Lambda for so many files might not be cost effective.  
  
If anyone here has some idea or suggestions then that would be great!
## [9][EventBridge processing pipeline](https://www.reddit.com/r/aws/comments/j1p9ul/eventbridge_processing_pipeline/)
- url: https://www.reddit.com/r/aws/comments/j1p9ul/eventbridge_processing_pipeline/
---
I’m designing a custom event-based processing pipeline and want to use EventBridge as the bus.  This gives us a single endpoint to use across the application (multiple event types) and provide to customers for customer-specific events, route to targets based on the payload (instead of attributes), etc.  All the reasons why EventBridge exists.  Except for the cross-region limitation.  

I want to deploy the processing pipeline into 2 regions for either active/active processing (load balance the two) or active/passive processing (1 region acts as a failover for the other).  

Regardless, I know I can have the bus in the primary region target another bus in the secondary region, but I’m wondering if there is any other pattern that’s elegant and makes sense?  

How can I either failover to the secondary region or load balance between the two?

Note: I’m not necessarily trying to solve for a regional outage (the EventBridge bus itself would be down as well).  

But then that raises another question...how are people solving for a regional outage where the bus itself is down?  Just provide a list of endpoints in different regions?

Hope to hear how other people are architecting around this.
## [10][SNS doesn't deliver SMS](https://www.reddit.com/r/aws/comments/j1wyzx/sns_doesnt_deliver_sms/)
- url: https://www.reddit.com/r/aws/comments/j1wyzx/sns_doesnt_deliver_sms/
---
Hi,  
I'm using Amazon SNS to send SMS text messages to verify and notify my users. But these several days the SMS never came. There's no error at all, on the dashboard it says 100% delivery rate all green but in reality the SMS never came. I tried on several phone provider (SEA region Singapore) with SDK and directly from dashboard all got same result, it says successful but SMS never came. I also get billed.  


Anyone can help? Like where to check or report or anything? Or SNS has always been unreliable?
## [11][Automated help on quota increases?](https://www.reddit.com/r/aws/comments/j1u9rg/automated_help_on_quota_increases/)
- url: https://www.reddit.com/r/aws/comments/j1u9rg/automated_help_on_quota_increases/
---
I've found the quota system pretty frustrating. I dislike waiting for my quota request and find it infuriating when my request is declined. Is my money not green enough?

I'd like a website that automatically requests maximum quota for all services. When quota requests are denied or limited, I want insights into what I can do to get that request approved.

Is this something that would be useful for you as well? Curious how you guys handle quotas and if this website idea is worth building.
