# aws
## [1][How are you guys performing cost optimisation?](https://www.reddit.com/r/aws/comments/f37s6n/how_are_you_guys_performing_cost_optimisation/)
- url: https://www.reddit.com/r/aws/comments/f37s6n/how_are_you_guys_performing_cost_optimisation/
---
We are using scheduling policies for ASG and lambda functions to stop RDSs during weekends. I was wondering what other approach can be taken to reduce the costs. Specifically around changing instance type if it's more than enough to perform the task.

Any thoughts?
## [2][Use CloudFormation StackSets for Multiple Accounts in an AWS Organization](https://www.reddit.com/r/aws/comments/f2vp8f/use_cloudformation_stacksets_for_multiple/)
- url: https://aws.amazon.com/blogs/aws/new-use-aws-cloudformation-stacksets-for-multiple-accounts-in-an-aws-organization/
---

## [3][How do you manage contact forms with static sites?](https://www.reddit.com/r/aws/comments/f38oke/how_do_you_manage_contact_forms_with_static_sites/)
- url: https://www.reddit.com/r/aws/comments/f38oke/how_do_you_manage_contact_forms_with_static_sites/
---
I've got a static site that needs a contact page and I've tried tonnes of CloudFormation templates but they are either too old or not designed for the SES sandbox. It's driving me nuts and I'm tempted to just use a 3rd party form but I like the idea of messing around with things like SES, SNS, and DynamoDB etc.. Thanks
## [4][How long from keys in github to a compromised account?](https://www.reddit.com/r/aws/comments/f39jna/how_long_from_keys_in_github_to_a_compromised/)
- url: https://www.reddit.com/r/aws/comments/f39jna/how_long_from_keys_in_github_to_a_compromised/
---
I'm trying to teach people a security lesson and wondered if anyone had any stats to back up my ideas.

If an AWS key/secret were posted to Github (or other popular public service), how quickly are you likely to see attempts to use the keys?

I say Github, because there are people out there scanning it for just such leaks. I could post them to twitter and there is a much wider audience, but I don't really expect anyone to be watching twitter for AWS\_....=AKIA....

Obviously you don't need to confess to leaking the keys yourself! :-) "Someone I know put their key in github and within x minutes there were attempted logins" ?

(I will follow up with the "Code Spaces" story)
## [5][Best Practice For Configuration Management](https://www.reddit.com/r/aws/comments/f39jc9/best_practice_for_configuration_management/)
- url: https://www.reddit.com/r/aws/comments/f39jc9/best_practice_for_configuration_management/
---
At the moment, we're using Cloudformation with gitlab runners to do our build and deployments. We're ramping up with new customers quite quickly, and my template mappings are blowing out in size with all the new variables I've had to pull out. 

I thought about moving them all to parameters, but it feels very haphazard to have such a huge deploy command that's difficult to untangle. Does anyone here have some suggestions? I'm thinking about looking into either a custom resource to pull in my variables, or using something like Chef or Ansible to do my deployments instead
## [6][AWS Client VPN seems like a rip off, no?](https://www.reddit.com/r/aws/comments/f2t9ds/aws_client_vpn_seems_like_a_rip_off_no/)
- url: https://www.reddit.com/r/aws/comments/f2t9ds/aws_client_vpn_seems_like_a_rip_off_no/
---
Is it just me being a cheap bastard and used to buying OpenVPN AS licenses for my BYOL OpenVPN AS instances or is AWS Client VPN ridiculously over priced?

10 VPN users connected for 1 hour is $0.60 with AWS Client VPN. 

Also I don’t think it goes across accounts with VPC peering. Meaning normally I do 1 AWS account per environment with a central core services (Admin) VPC/Account and have the accounts peer into that one. Usually I only need a single OpenVPN AS box in that Admin VPC/Account and I’m good.

Thoughts? Can we expect this new service offering to come down in price any time soon? Even if it does I think it’s hard to compete with a $150 one time license for 10 users and whatever monthly run cost for the box.
## [7][Converting my CloudFront Lambda@Edge Function from JavaScript to Python](https://www.reddit.com/r/aws/comments/f37qx0/converting_my_cloudfront_lambdaedge_function_from/)
- url: https://adamj.eu/tech/2020/02/13/converting-my-cloudfront-lambda-at-edge-function-from-javascript-to-python/
---

## [8][AWS Storage Gateway for software development company](https://www.reddit.com/r/aws/comments/f35i3e/aws_storage_gateway_for_software_development/)
- url: https://www.reddit.com/r/aws/comments/f35i3e/aws_storage_gateway_for_software_development/
---
Hey. I run a 40+ developer software development agency and use AWS EC2 for dev/staging instances of our customers. During local project setups, the download of the database and files of the project usually takes around 30 minutes at \~7/8 MB/s. I was thinking if Storage Gateway would have a big impact with local project setups. Have anyone here done or seen something similar? Would you recommend Storage Gateway for this kind of problem? Or do you think it is better to just keep downloading from AWS EC2 instances at a quite-okay speed?
## [9][Amazon hyperledger fabric](https://www.reddit.com/r/aws/comments/f34zly/amazon_hyperledger_fabric/)
- url: https://www.reddit.com/r/aws/comments/f34zly/amazon_hyperledger_fabric/
---
Amazon Managed Blockchain

When using amazon managed blockchain's hyperledger fabric, I am wondering if each peer node can run a dockerized web application.  Or would I need to run an application on a separately hosted container that has logic built in for connecting to the VPC for submitting transactions to the hyperledger fabric? 


In my case I have specific GPU and OS requirements for my container to run properly before data can be submitted and the transaction recorded. If this is possible on the amazon managed blockchain nodes via the VPC, can I specify the hardware requirements and OS for each node to host this application? Or is it wasteful to request such significant resources for each blockchain node if the end goal is simply to record a transaction? Would it be better to run the computationally costly application on a separate kubernetes instance then write transaction to ledger? This is for labeling and then uploading imaging data to a longterm storage location. I am using hyperledger fabric to record each time data is submitted and retrieved from long term storage.
## [10][Network Load Balancer Target Group Health Check TCP Expected Response](https://www.reddit.com/r/aws/comments/f349wh/network_load_balancer_target_group_health_check/)
- url: https://www.reddit.com/r/aws/comments/f349wh/network_load_balancer_target_group_health_check/
---
Been scouring the documentation site for hours now, can anyone point me to the page on what the target group expects as a healthy response on a TCP health check?

I have a port open that just replies with OK like so

`TcpClient client = server.AcceptTcpClient();NetworkStream ns = client.GetStream();byte[] hello = Encoding.Default.GetBytes("OK");ns.Write(hello, 0, hello.Length);`

and have done a nc on a separate ec2 and it all works

`root@ip-10-200-10-93:/home/ubuntu# nc -vz` [`10.200.31.240`](https://10.200.31.240) `4416Connection to` [`10.200.31.240`](https://10.200.31.240) `4416 port [tcp/*] succeeded!`

But somehow, target group health checks still return unhealthy. I read somewhere it just needs a TCP connection, then I read another where it needs an HTTP reply? Really need someone to point me to the right direction. Thanks

This is for Network Load Balancer TCP
