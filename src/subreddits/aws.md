# aws
## [1][Easily create production ready serverless app powered by multi-account CI/CD pipeline in just few minutes, with my 1st open-source project](https://www.reddit.com/r/aws/comments/fpje9b/easily_create_production_ready_serverless_app/)
- url: https://github.com/eficode/serverless-ops-boilerplate
---

## [2][Policy for a Cloud Economist job function?](https://www.reddit.com/r/aws/comments/fpvvak/policy_for_a_cloud_economist_job_function/)
- url: https://www.reddit.com/r/aws/comments/fpvvak/policy_for_a_cloud_economist_job_function/
---
arn:aws:iam::aws:policy/job-function/Billing + compute-optimizer:*

Or a more limited https://s.natalian.org/2020-03-27/billing.json + compute-optimizer:*

Any suggestions please?
## [3][[LAMBDA/DYNAMODB] PUT /resource/:id how to handle the metadata ?](https://www.reddit.com/r/aws/comments/fpuh9x/lambdadynamodb_put_resourceid_how_to_handle_the/)
- url: https://www.reddit.com/r/aws/comments/fpuh9x/lambdadynamodb_put_resourceid_how_to_handle_the/
---
Hi People, 

I will illustrate my problem with an example.

Imagine you have the User resource. Let's say we have :  
POST /users =&gt; add a user  
PUT /users/:id =&gt; edit a user. 

When we add a user, we persist it in a dynamoDB with : id, lastname, firstname, age, creationDate, lastModificationDate. 

How would you handle the update ? Because the client will just send us in body :   
name, firstname, age.   


To avoid erase creationDate - lastModificationDate, how would you handle this ?   
I know it exists just partial update by using update() instead of put() in aws-sdk.   


But in my real case, we have a lot of properties, it seems to complicated to handle each one in a update query.   


So we thought a basic solution :   
\- Get the user, retrieve the id/creationDate/modificationDate to add it in the new object, then put it. =&gt; seems also complex because when we will need to add new properties, we don't have to forget to edit this part of code.   
To facilitate, we could add a "metadata" property.

How do you handle this in your project ?

If I was not clear, don't hesitate. \^\^  


Thanks
## [4][AWS Route53 CNAME record](https://www.reddit.com/r/aws/comments/fpwpy7/aws_route53_cname_record/)
- url: https://www.reddit.com/r/aws/comments/fpwpy7/aws_route53_cname_record/
---
Hello!

I am about to edit an old CNAME record which looks something like this.

[abc.set.com](https://abc.set.com) CNAME [ns.te.com](https://ns.te.com).

I will report [ns.te.com](https://ns.te.com) with [st.set.com](https://st.set.com).

I can see that the old CNAME contains a . in the last. 

Does this make any difference.

[abc.set.com](https://abc.set.com) CNAME [ns.te.com](https://ns.te.com).

[abc.set.com](https://abc.set.com) CNAME [st.set.com](https://st.set.com)

Whats the purpose of "." in the last in the CNAME record?

&amp;#x200B;

Thanks
## [5][EC2 with Elastic IP and S3 Data Transfer cost](https://www.reddit.com/r/aws/comments/fpvnjj/ec2_with_elastic_ip_and_s3_data_transfer_cost/)
- url: https://www.reddit.com/r/aws/comments/fpvnjj/ec2_with_elastic_ip_and_s3_data_transfer_cost/
---
Please assure me that interconnection traffic between EC2 instance with Elastic IP attached toward S3 bucket **within same region** doesn't count against VPC outbound data transfer cost (0.09$ per GB).

*Conditions:*  
EC2 has single network adapter.  
No VPC Endpoint used.  
Same region.
## [6][What's the maximum time /tmp files are stored in AWS Lambda?](https://www.reddit.com/r/aws/comments/fpdbnv/whats_the_maximum_time_tmp_files_are_stored_in/)
- url: https://www.reddit.com/r/aws/comments/fpdbnv/whats_the_maximum_time_tmp_files_are_stored_in/
---
I've read that "AWS Lambda **maintains the execution context for some time in anticipation of another Lambda function invocation**". - but what does that actually mean in practice? How frequently would the function have to be called to still be in the same execution context? What's the maximum possible time that something would stay in the tmp folder?

My project requires that it stays 1 hour or less, but I can't find out if what the AWS docs say would mean that that's what will happen.
## [7][NLB --&gt; EC2 possible?](https://www.reddit.com/r/aws/comments/fpu3cq/nlb_ec2_possible/)
- url: https://www.reddit.com/r/aws/comments/fpu3cq/nlb_ec2_possible/
---
Hello, im trying to use an ELB with an EIP attached to balance traffic for my EC2 hosted website. However i can see AWS for situations like this recommend that you use a setup like:

NLB (with EIP) --&gt; ALB --&gt; EC2

And i wonder why? Why not just NLB --&gt; EC2?
## [8][Challenges with network connectivity between VPN Clients and EC2 instances](https://www.reddit.com/r/aws/comments/fplxre/challenges_with_network_connectivity_between_vpn/)
- url: https://www.reddit.com/r/aws/comments/fplxre/challenges_with_network_connectivity_between_vpn/
---
Hi r/aws!

&amp;#x200B;

I'm having some challenges with full connectivity between my VPN endpoint clients and EC2 instances.

VPC subnet: [172.31.0.0/20](https://172.31.0.0/20)

Client VPN: [192.168.200.0/22](https://192.168.200.0/22)

&amp;#x200B;

When my VPN clients try to access EC2 instances everything works fine.  For example, users can RDP to my remote desktop server without issue.

When I try to reach out to my VPN clients through (using Lansweeper, or even ping them), I can't make a connection.

To test, i've connected my PC to the VPN, then turned off my firewall, still the same result.

My Network ACLs and Security Groups are all very permissive currently (as a results of testing), but I still the same problems.

My thought was that my EC2 instances don't know how to route to my VPN clients, but I can't seem to add an entry into the VPC route table for this.  There is no target I can reference for the Client VPN Endpoint subnet.

&amp;#x200B;

Is there something simple I am missing here?
## [9][Confused by Billing Alerts](https://www.reddit.com/r/aws/comments/fphgqq/confused_by_billing_alerts/)
- url: https://www.reddit.com/r/aws/comments/fphgqq/confused_by_billing_alerts/
---
I only have experience with Cost and Usage reporting, but that is a bit heavy for my current need. Thus, I'm reading https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/monitor_estimated_charges_with_cloudwatch.html and getting confused by billing alerts (not alarms, which I think I understand). It says that once turned on, billing alerts cannot be turned off. From that page, it looks like billing alerts is in fact a Cloudwatch metric called Estimated Charge. Questions:

- Why would I want to turn billing alerts off?
- If they are metrics, do I get charged normal CW metric rates for them?
## [10][Undoing transaction error](https://www.reddit.com/r/aws/comments/fpklyl/undoing_transaction_error/)
- url: https://www.reddit.com/r/aws/comments/fpklyl/undoing_transaction_error/
---
Hi Gurus,

While running the copy command to load the data from S3 into Redshift I am getting below error in logs:
Undoing 1 transactions ON TABLE 456834 with current xid ...
Could you please let me know where I am going wrong?
I am running the copy command from from python using psycopg2 .

Thanks,
S
