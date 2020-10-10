# aws
## [1][How to Develop Serverless Vuejs Application with AWS Amplify?](https://www.reddit.com/r/aws/comments/j8i5st/how_to_develop_serverless_vuejs_application_with/)
- url: https://tekkiwebsolutions.com/blog/serverless-vuejs-application-with-aws-amplify/
---

## [2][Why Doesn't AWS Have a Cloud Run Equivalent?](https://www.reddit.com/r/aws/comments/j84h22/why_doesnt_aws_have_a_cloud_run_equivalent/)
- url: https://www.reddit.com/r/aws/comments/j84h22/why_doesnt_aws_have_a_cloud_run_equivalent/
---
Does anyone know why AWS doesn't have something similar to [Cloud Run](https://cloud.google.com/run) where you run your container and are billed only when your container receives incoming requests? It is similar to Lambda but instead of FaaS, it is CaaS but with the billing model of FaaS, unlike ECS and EKS where your container runs all the time. I would think that this would be an attractive option for companies that are still building traditional apps that can be containerized but don't want the complexities of ECS or EKS and want to move to the cloud and benefit from the auto-scaling, per second billing, etc. In Lambda, AWS is already running a full container but to serve a single request at a time. Using Cloud Run, you can serve dozens or more concurrent requests using the same processing footprint
## [3][I am attempting to get Create an SSL Certificate using AWS Certificate Manager. My DNS is google domains. Followed steps and it won't get out of pending status for status, help?](https://www.reddit.com/r/aws/comments/j8eht7/i_am_attempting_to_get_create_an_ssl_certificate/)
- url: https://www.reddit.com/r/aws/comments/j8eht7/i_am_attempting_to_get_create_an_ssl_certificate/
---
So, basically, I have a domain name already bought and the DNS is google domains. With AWS Certificate Manager, I followed the steps here:

https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-request-public.html#request-public-console

Also, I both did *.example.com and example.com for domains when following the steps (not example.com, but the domain I have. The * is for wildcard).

For step five in that step by step, I selected DNS validation. At the end, I was given instructions to add a CNAME record to the DNS configurations.

So I went to google domains, entered the NAME section given as it was written in AWS, Selected CNAME for type, and for the Value I entered the exact value given in AWS. They literally give you all the values and I copied and pasted them in and verified they are correct.

Yet, after all of that, I am still given the validation status in AWS as "Pending Validation". It has been over 12 hours now and the status is still pending validation.

What is going on? To be clear, I have used this domain before with github pages and it worked fine and linked up fine. I since removed those records from google domains and so they should no longer interfere with anything. Also verified that nothing shows anymore either when going to my URL.

Does anyone have any suggestions as to how I can get the validation status to go from "pending validation" to "success"?

Just very confused and frusterated, as I am basically following what the documentation says. Hopefully someone can help, thanks.
## [4][Blacklist/whitelist AZs for Autoscaling-group when an AZ is misbehaving](https://www.reddit.com/r/aws/comments/j8jz7a/blacklistwhitelist_azs_for_autoscalinggroup_when/)
- url: https://www.reddit.com/r/aws/comments/j8jz7a/blacklistwhitelist_azs_for_autoscalinggroup_when/
---
Follow-up to current issue with AWS availability zone - [I wrote a script to remove/add AZs from/to an ASG to avoid launching new instances in the faulty AZ.](https://github.com/awsiv/asg-az-update)

Comments welcome :)
## [5][Amazon EMR integration with AWS Lake Formation is now generally available](https://www.reddit.com/r/aws/comments/j84ld1/amazon_emr_integration_with_aws_lake_formation_is/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/10/amazon-emr-integration-with-aws-lake-formation-is-now-generally-available/
---

## [6][US-EAST-1 outage use1-az2](https://www.reddit.com/r/aws/comments/j7xvib/useast1_outage_use1az2/)
- url: https://www.reddit.com/r/aws/comments/j7xvib/useast1_outage_use1az2/
---
We are facing issue for new instance scaled from 1:20 AM PDT having netowrk connectivity issue.. Tried other 2 AZ B,A &amp; D..

Is anyone else facing problem? What other service.. I believe all major service calling EC2 or network interface will be affected..

&gt;Network connectivity issues  \[01:48 AM PDT\] We are investigating networking connectivity issues for a small subset of newly launched EC2 instances within a single Availability Zone (use1-az2) in the US-EAST-1 Region. We have identified root cause and are working towards resolution. Network connectivity for existing instances is not affected by this issue. For newly launched instances that are affected, relaunching a new instance may resolve the issue.  
&gt;  
&gt;\[02:47 AM PDT\] We continue to work toward recovery for the networking connectivity issues affecting a small subset of newly launched EC2 instances within a single Availability Zone (use1-az2) in the US-EAST-1 Region. Network connectivity for existing instances remains unaffected by this issue. For newly launched instances that are affected, relaunching a new instance may resolve the issue.  
&gt;  
&gt;\[04:53 AM PDT\] We are still working toward recovery for the networking connectivity issues affecting a small subset of newly launched EC2 instances within a single Availability Zone (use1-az2) in the US-EAST-1 Region. Network connectivity for existing instances remains unaffected by this issue. For newly launched instances that are affected, launching a replacement instance may resolve the issue.

We beleive the issue started well before this.. probably 3-4 hrs before,  when we started noticing the higher error rates but were not sure of incident then. Can AWS come clean on when it actually started &amp; how soon can this be resolved??  sitting ducks for now!.
## [7][Managed Service SLA's](https://www.reddit.com/r/aws/comments/j8bjjb/managed_service_slas/)
- url: https://www.reddit.com/r/aws/comments/j8bjjb/managed_service_slas/
---
Does anyone self-measure managed service SLA's for AWS services like RDS or MSK. They claim 99.95% but how do you (1) know you are receiving the advertised uptime, and (2) detect/measure outages to submit claims for reimbursement?

Is this automated by AWS? Is there an un-biased 3rd party service/auditor that does this?
## [8][AWS Site-to-Site VPN Routing Configuration Trouble](https://www.reddit.com/r/aws/comments/j8gxqu/aws_sitetosite_vpn_routing_configuration_trouble/)
- url: https://www.reddit.com/r/aws/comments/j8gxqu/aws_sitetosite_vpn_routing_configuration_trouble/
---
Hello everyone,

New to AWS and trying to configure a VPN connection to our data centre and wanted to confirm something. I have the following subnets and VPC configured in AWS and in the Data Centre

* VPC - 192.168.0.0/16
* Local subnet- 192.168.100.0/24
* Data centre subnet - 192.168.200.0/24

Neither subnets overlap and I have successfully brought up both the IPsec tunnels. I have added the propagated route to the route table and for the moment have all traffic allowed on the NACL and Security group. Route Table as follows:

* 192.168.0.0/16 - Local
* 192.168.200.0/24 - VGW (propagated)
* 0.0.0.0/0 - Nat Gateway

Now I can see traffic hitting my ENI in VPC flow logs, but nothing is being received back at the data centre. Upon doing some research I found the following documentation on routing that states the following:

* If propagated routes from a Site-to-Site VPN connection or AWS Direct Connect connection overlap with the local route for your VPC, the local route is most preferred even if the propagated routes are more specific.

[https://docs.aws.amazon.com/vpn/latest/s2svpn/VPNRoutingTypes.html#vpn-route-priority](https://docs.aws.amazon.com/vpn/latest/s2svpn/VPNRoutingTypes.html#vpn-route-priority)

So if my understanding is correct, I cannot use a remote subnet that is in the range of my VPC. If this is the case why is that not something that in enforced when creating the subnet? Given that most VPNs to my knowledge only care about conflicting Subnets. Which in this case aren't conflicting.

Can someone confirm that my understanding is correct? Or if there is a way that I can configure the routing to allow the traffic to my data centre without recreating the entire subnet?

Many thanks. Really appreciate any help possible.
## [9][AWS Cloud Map simplifies service discovery with optional parameters](https://www.reddit.com/r/aws/comments/j86d9u/aws_cloud_map_simplifies_service_discovery_with/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/10/aws-cloud-map-simplifies-service-discovery-optional-parameters/
---

## [10][DynamoDB streams to ingest to Elasticsearch](https://www.reddit.com/r/aws/comments/j8ah2k/dynamodb_streams_to_ingest_to_elasticsearch/)
- url: https://www.reddit.com/r/aws/comments/j8ah2k/dynamodb_streams_to_ingest_to_elasticsearch/
---
We have an elastic search domain which we put some of our dynamodb records into to facilitate with queries.  Right now we have a dynamodb stream that insert  or updates records to elasticsearch when the dynamodb table changes, but we are seeing it take 1 minute to move over 40 records.  We also have parallelism set to 10, but the performance is still bad.  Is there a special configuration that needs to be made to make dynamodb streams processed at a faster rate?
