# aws
## [1][Which is better for S3 uploads from client side - HTTP POST form or using a Cognito identity](https://www.reddit.com/r/aws/comments/ewjyqa/which_is_better_for_s3_uploads_from_client_side/)
- url: https://www.reddit.com/r/aws/comments/ewjyqa/which_is_better_for_s3_uploads_from_client_side/
---
Originally posted on stackoverflow: [https://stackoverflow.com/questions/59967346/which-is-better-for-s3-uploads-from-client-side-http-post-form-or-using-a-cogn](https://stackoverflow.com/questions/59967346/which-is-better-for-s3-uploads-from-client-side-http-post-form-or-using-a-cogn)

&gt;I want to upload a file to S3 directly from client browser. I don't want that file to go through my server.    
&gt;  
&gt;After doing some research I am seeing two ways to do this:   
&gt;  
&gt;  
&gt;  
&gt; 1. \[Cognito Identity\]\[1\]  
&gt;  
&gt; 2. \[Upload using HTTP POST form\]\[2\]    
&gt;  
&gt;  
&gt;  
&gt;  
&gt;  
&gt;Issue with Cognito Identity is that we are not using a user pool. We are authenticating using our database without AWS Cognito User Pool. And without that we only have 'IdentityPoolId', which according to me, anyone can copy from our website's JS code and upload objects to our S3 bucket. Correct me if I am wrong with this assumption.    
&gt;  
&gt;  
&gt;  
&gt;Using HTTP POST form seems OK to me but while researching I am seeing that Cognito is the new way to do this and I'm afraid that I'm missing something.  
&gt;  
&gt;  
&gt;  
&gt;  
&gt;  
&gt;  \[1\]: [https://aws.amazon.com/blogs/mobile/building-fine-grained-authorization-using-amazon-cognito-user-pools-groups/](https://aws.amazon.com/blogs/mobile/building-fine-grained-authorization-using-amazon-cognito-user-pools-groups/)  
&gt;  
&gt;  \[2\]: [https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-authentication-HTTPPOST.html](https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-authentication-HTTPPOST.html)
## [2][Setup Kibana with Fargate](https://www.reddit.com/r/aws/comments/ewo5sz/setup_kibana_with_fargate/)
- url: https://www.reddit.com/r/aws/comments/ewo5sz/setup_kibana_with_fargate/
---
I'm trying to get Kibana set up correctly with Fargate. I have followed this simple guide: [https://aws.amazon.com/blogs/compute/building-deploying-and-operating-containerized-applications-with-aws-fargate/](https://aws.amazon.com/blogs/compute/building-deploying-and-operating-containerized-applications-with-aws-fargate/)

The Fargate Task Definition is set up to run a Kibana docker image, based on this Dockerfile [https://github.com/docker-library/kibana/blob/4eef267ab6d1e6a655f4e44dabe9aef7915e4b99/7/Dockerfile](https://github.com/docker-library/kibana/blob/4eef267ab6d1e6a655f4e44dabe9aef7915e4b99/7/Dockerfile)  


Since Kibana demands incoming traffic on port 5601, I have opend this port in my awsvpc.

The reponse I get is **502 Bad Gateway**, when I go to the DNS name for the loadbalancer (port 5601).  


What can the problem be? I think maybe it has something to do with the fargate task and the docker image, that has problem routing the traffic to the correct port or have trouble making the port 5601 container port available for incoming traffic.

&amp;#x200B;

When I run the dockerimage on my local computer, it is possible to add parameters to define portmapping. Is this possible for fargate-task aswell?  


See attached image, for how the container is set up for fargate task:  


https://preview.redd.it/ay1bzjiz34e41.png?width=1033&amp;format=png&amp;auto=webp&amp;s=c8234dc8eb0cfd0f4bc35396e1799aa30acf6ea9
## [3][RDS Useage Question](https://www.reddit.com/r/aws/comments/ewmfui/rds_useage_question/)
- url: https://www.reddit.com/r/aws/comments/ewmfui/rds_useage_question/
---
For two months, I've been using the free tier of RDS to use MySQL with a CRUD webapp for my CV. 

It gives you 750 hours a month and if you exceed that you have to pay, without an option to disable the service once the limit has been reached.

So far it has come slightly below the 750 hours, so I have not had to pay. The thing is, I don't know what would factors contribute towards the hours used. So far it has only been me using the webapp. I read on documentation that it has to do with instances used. Does that mean that it is affected by different IPs checking out the webapp, like when I show it to prospective employers?

I am concerned because if it is coming to just under the 750 hours with just me using it, if employers start having a look, I will have to pay a lot of money.
## [4][Understanding VPCs?](https://www.reddit.com/r/aws/comments/ewbhvb/understanding_vpcs/)
- url: https://www.reddit.com/r/aws/comments/ewbhvb/understanding_vpcs/
---
Trying to learn VPC, see below garbage pic for reference:

https://preview.redd.it/8eu28rb23zd41.png?width=626&amp;format=png&amp;auto=webp&amp;s=d6eccea0d5a0137c73fe71f92204a4169bac1507

VPC questions:

* I (generally) want my databases in private subnets?
   * If for some reason they need internet I'd do a NAT GW in (another) public subnet?
* I (generally) want my EC2 instances in private subnets?
   * If for some reason they need internet I'd do a NAT GW in (another) public subnet?
* I generally want my ELB in a public subnet?
* Does the ASG for the EC2s go in a public or private subnet? 
   * Or is this not really a thing that has a "location"?
* I only need 1 internet gateway per VPC?
* In what scenario would 1 single VPC need NAT GWs in 2 separate (public) subnets.
* What makes a subnet private vs public?
   * Is this just weather or not it's routed to an IG?
* Is the only difference between a NAT gateway and an internet gateway that a NAT gateway is outbound only?
## [5][Cost of Nested Queries on AWS Athena](https://www.reddit.com/r/aws/comments/ewgdw2/cost_of_nested_queries_on_aws_athena/)
- url: https://www.reddit.com/r/aws/comments/ewgdw2/cost_of_nested_queries_on_aws_athena/
---
Hi folks, 

Something at work is requiring me to use Athena to query something like 2 TB of data from a client S3 bucket.

I was wondering about the pricing structure of this. I designed my query based off of a very small testing file with the same shape as the client data, so I designed the process using something like 5 queries that produced interim tables, 3 of which add columns, 1 which splits the data and re-unions, and one which groups the data by a set of columns. 

I have the option of consolidating these queries into one nested query.

So according to what AWS says on their website about Athena pricing, this single nested query should cost around $10 ($5/TB * 2 TB), while running the process using interim tables described above would incur about $50 of costs. Does anyone know if this is actually the case? Are you charged based on the number of SELECT statements are contained within your query, or can you actually save money by consolidating SELECTs into a nested query?

Thanks for the help everyone!
## [6][ELI5: NLB vs ALB](https://www.reddit.com/r/aws/comments/ewby6y/eli5_nlb_vs_alb/)
- url: https://www.reddit.com/r/aws/comments/ewby6y/eli5_nlb_vs_alb/
---
I've read through the AWS documentation on ELB vs ALB and still don't have a good idea on when to use one vs the other - https://aws.amazon.com/elasticloadbalancing/features/

From reading the documentation

ALB allows for layer 7 routing based on HTTP headers 

Whereas

NLB allows for layer 4 routing which is strictly UDP/TCP - network layer.

They share commonalities - high availability, security features, health checks etc

It's not as simple as if you don't require routing based off HTTP headers, use NLB?

I've read that NLB can scale better but I don't understand why it can scale better. Is it because it doesn't have to do any processing of any layers above layer 4 - network?
## [7][Why cloudwatch keeps charging me and refund credits every month?](https://www.reddit.com/r/aws/comments/ewljot/why_cloudwatch_keeps_charging_me_and_refund/)
- url: https://i.redd.it/na32f0i3v2e41.jpg
---

## [8][AWS ElasticSearch Node](https://www.reddit.com/r/aws/comments/ewlbs9/aws_elasticsearch_node/)
- url: https://i.imgur.com/WsjmsO0.png
---

## [9][What is "AWS Security Scanner" in my server logs?](https://www.reddit.com/r/aws/comments/ew5lzt/what_is_aws_security_scanner_in_my_server_logs/)
- url: https://www.reddit.com/r/aws/comments/ew5lzt/what_is_aws_security_scanner_in_my_server_logs/
---
I am getting these in my ec2 nginxn logs:

"GET /latest/dynamic/instance-identity/document HTTP/1.1" 400 264 "-" "AWS Security Scanner

Is this just a bot trying to do smt shady or a AWS internal "scanner" ?
## [10][Step Functions Express Workflows](https://www.reddit.com/r/aws/comments/ewkv5q/step_functions_express_workflows/)
- url: https://www.reddit.com/r/aws/comments/ewkv5q/step_functions_express_workflows/
---

Step Functions, which were launched in November 2016, are a valuable resource for the orchestration of both simple and complex processes in AWS. 
Restricted to 2,000 execution starts per second, you may ask, how can we manage workflows for greater volumes of data? This blog post explores AWS' answer to this problem: express workflows:

https://medium.com/@trrhodes/step-functions-express-workflows-cb8436259af6?source=friends_link&amp;sk=a26219e52771af6da7068ec1fc05d629
