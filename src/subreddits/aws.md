# aws
## [1][Can I replace NLB with ECS Service Discovery?](https://www.reddit.com/r/aws/comments/inwung/can_i_replace_nlb_with_ecs_service_discovery/)
- url: https://www.reddit.com/r/aws/comments/inwung/can_i_replace_nlb_with_ecs_service_discovery/
---
I currently have ECS Service running with an endpoint serve through a Network Load Balancer. The Network Load Balancer is transparent and directs all traffic to ECS. 

&amp;#x200B;

From my understanding, if I remove NLB and direct traffic to ECS Service, I could potentially have downtime if the task is removed and over-loaded. However, I have recently learned that NLB can be removed and ECS Service Discovery will take care of serving the healthy tasks. 

&amp;#x200B;

I am learning about Service Discovery right now and having difficulty understanding how it can replace NLB?
## [2][When and why use AWS Elastic Beanstalk?](https://www.reddit.com/r/aws/comments/inxcd4/when_and_why_use_aws_elastic_beanstalk/)
- url: https://www.reddit.com/r/aws/comments/inxcd4/when_and_why_use_aws_elastic_beanstalk/
---
Hi guys.  


The company I work for has several different applications running on EC2 instances - they are all Java-based applications - and I was wondering if I could migrate them from EC2 to Beanstalk? and if I should?

What are the benefits of Beanstalk? We have Rest APIs that use Jetty. Would Beanstalk handle that?

Is there any reason why I shouldn't move to Beanstalk?
## [3][How to allow more instance types in Cloudformation?](https://www.reddit.com/r/aws/comments/io6img/how_to_allow_more_instance_types_in_cloudformation/)
- url: https://www.reddit.com/r/aws/comments/io6img/how_to_allow_more_instance_types_in_cloudformation/
---
Hi,

I'm using Elastic Beanstalk to launch EC2 instances. I'm trying to upgrade from t2.micro to t3a.micro, but I keep getting the following error when I try to change instancetype in Beanstalk: 

    Service:AmazonCloudFormation, Message:Parameter InstanceType failed to satisfy constraint: must be a valid EC2 instance type.

I saw that Cloudformation template didn't allow t3a.micro instances as it was rather old template so I updated "InstanceType"-field and its AllowedValues array to include t3a.micro, but I still get the error above.

Is there something that I'm missing or why Elastic Beanstalk is not letting me upgrade to t3?
## [4][Why no interface endpoint for AWS Inspector?](https://www.reddit.com/r/aws/comments/io6bxu/why_no_interface_endpoint_for_aws_inspector/)
- url: https://www.reddit.com/r/aws/comments/io6bxu/why_no_interface_endpoint_for_aws_inspector/
---
We have been using interface endpoints to keep AWS command traffic within our environment as much as possible.

Looking at our firewall now, the majority of the AWS bound traffic that leaves our environment is for arsenal..amazonaws.com which is telemetry data from the AWS Inspector agent.

Is there a reason why this traffic isn't supported via a VPC interface endpoint?
## [5][Lambda with nodejs and middleware](https://www.reddit.com/r/aws/comments/inub4l/lambda_with_nodejs_and_middleware/)
- url: https://www.reddit.com/r/aws/comments/inub4l/lambda_with_nodejs_and_middleware/
---
The app I'm building has a descent number of entities, about 50. Assuming each entity will have at the minimum one set of crud functions, if I build them as skinny, one lambda per endpoint, I'm up to 200 lambdas without event starting to build the app. 
I'm afraid of negative impact in maintenance and cost around this quantity. Deploying is already starting to take some time and all these lambdas needing to be up for every tenant. 
Option would be to have a number of lambdas serving multiple entities grouped by domain, or service. 
If this is a better option, what kind of middleware do you recommend (trying to avoid express here and pick something aws lambda specific). Thanks.
## [6][Port Forward in EC2 instance](https://www.reddit.com/r/aws/comments/io5dhi/port_forward_in_ec2_instance/)
- url: https://www.reddit.com/r/aws/comments/io5dhi/port_forward_in_ec2_instance/
---
Hello, I'm new in AWS and recently created a web server and wanted to access it remotely. Then I saw that you can port forward by setting up an E2C instance. I did it successfully but my question now is if I have multiple web servers, should you create one instance for every web server or maybe there is a way to access your web servers through one EC2 instance?  

I created two web servers and I accessed them through two different ports 8080 and 8081

I used this command to forward these ports: 

**ssh -R :8081:192.168.0.106:8081 -i C:\\Users\\PC\\Desktop\\PForw.pem .......**

Sorry if I'm not clear... just a beginner in AWS.
## [7][How to setup aws VPN with on-prem Windows 10](https://www.reddit.com/r/aws/comments/inwxkt/how_to_setup_aws_vpn_with_onprem_windows_10/)
- url: https://www.reddit.com/r/aws/comments/inwxkt/how_to_setup_aws_vpn_with_onprem_windows_10/
---
I am studying for the CSAP exam and I’m running into a ton of on-prem topics that I’m not that familiar with. 

Does anyone know of any free VPN server software that is 1) compatible with Win 10, and 2) compatible with aws site-to-site VPN?

I got pretty far with OpenVPN, but from what I can tell it uses SSL only and does not seem to support IPSec which appears to be required by aws.
## [8][Can someone explain what AWS Lambda is in a way you would explain it to a person with basic cloud knowledge](https://www.reddit.com/r/aws/comments/io536s/can_someone_explain_what_aws_lambda_is_in_a_way/)
- url: https://www.reddit.com/r/aws/comments/io536s/can_someone_explain_what_aws_lambda_is_in_a_way/
---

## [9][EC2 Access to EFS volumes](https://www.reddit.com/r/aws/comments/inwzp2/ec2_access_to_efs_volumes/)
- url: https://www.reddit.com/r/aws/comments/inwzp2/ec2_access_to_efs_volumes/
---
I am studying for the Solutions Architect cert and got to the EFS section had this situation and question:

I created an EC2 instance and then mounted the EFS volume that I previously created. The instance was not associated with an IAM role.

What permission/policy/priv, etc., allows this EC2 instance to access the EFS volume?

My understanding was that EVERYTHING needs some kind of explicit permission.

I couldn't find anything specific about the instance that would suggest that it can mount/access and EFS volume ... except that the instance was created by a User with the "AdministratorAccess" policy.
## [10][How to find cost of untagged resources in AWS using Cost explorer API](https://www.reddit.com/r/aws/comments/io1cw3/how_to_find_cost_of_untagged_resources_in_aws/)
- url: https://www.reddit.com/r/aws/comments/io1cw3/how_to_find_cost_of_untagged_resources_in_aws/
---
We have 2 mandatory tags for every resource in AWS, we want to find the cost of the untagged resources, how is it possible to get the cost for the untagged resources ? 

I checked through the Cost explorer API docs but didnt get much for untagged cost findings, any suggestions ?

&amp;#x200B;

Thank you !
