# aws
## [1][The new AWS calculator is terrible.](https://www.reddit.com/r/aws/comments/fyvifi/the_new_aws_calculator_is_terrible/)
- url: https://www.reddit.com/r/aws/comments/fyvifi/the_new_aws_calculator_is_terrible/
---
Am I the only one that hates the new AWS calculator, [calculator.aws](https://calculator.aws)? I don't need a wizard to hold my hand through getting pricing. I need something simple, like the already existing simply monthly calculator. Now I see that the simple monthly calculator is going away on June 30, 2020, and we have to start using calculator.aws? Say it ain't so.
## [2][Service Meshes: An Introduction for Infrastructure/DevOps Teams (AWS App Mesh)](https://www.reddit.com/r/aws/comments/fyu1lb/service_meshes_an_introduction_for/)
- url: https://www.singlestoneconsulting.com/blog/what-are-service-meshes/
---

## [3][Test/Dev and Prod environments](https://www.reddit.com/r/aws/comments/fyxk54/testdev_and_prod_environments/)
- url: https://www.reddit.com/r/aws/comments/fyxk54/testdev_and_prod_environments/
---
I have been tasked with setting up test/dev and production environments in AWS. I'm curious how you would approach this. 

My initial instinct is to create two separate VPCs with different private subnets. I've read a couple of other posts where people separate the accounts. I can see where this comes into play for IAM among other things. I believe IAM is account specific so going with two VPC's i'd have to have to take that into account for groups, etc. 

Any input and suggestions are welcome!
## [4][New to making a server, need guidance/advice](https://www.reddit.com/r/aws/comments/fz52jt/new_to_making_a_server_need_guidanceadvice/)
- url: https://www.reddit.com/r/aws/comments/fz52jt/new_to_making_a_server_need_guidanceadvice/
---
Hi, I am new to making server/backend with aws. 

Currently I have a setup where clients can communicate to our server via grpc and we have a python3 backend using boto3. 

The concern we have is that it seems like we are adding another layer that only seems to create an overhead between our clients and the server. For instance, if client wishes to authenticate itself, it would send its username and password via grpc to our python backend, and our boto3 would relay that information to cognito. Cognito tells the server whether user was authenticated with accessToken, and the server sends back exact same info to the client. 

This seems cumbersome and defeats the purpose of having an accessToken. The problem is that aws sdk does not support the language we wish to deploy to our clients, and this seems to be the only way to go about it. Even if our client side language is in python, is it a good practice to have boto3 do the work on the server side like this? 

Thanks in advance.

By the way, I'm not actually sure how JWT tokens work. Once the client has a hold of it, how does the client use it?
## [5][27 DynamoDB Best Practices [Bite-sized Tips]](https://www.reddit.com/r/aws/comments/fyjdr2/27_dynamodb_best_practices_bitesized_tips/)
- url: https://dynobase.dev/dynamodb-best-practices/
---

## [6][S3 and Cloudfront web deployment - inconsistent rendering](https://www.reddit.com/r/aws/comments/fyqmvr/s3_and_cloudfront_web_deployment_inconsistent/)
- url: https://www.reddit.com/r/aws/comments/fyqmvr/s3_and_cloudfront_web_deployment_inconsistent/
---
I've deployed a React app using S3 and Cloudfront and the content being displayed is very inconsistent. For one, I changed the app favicon and page titles to reflect my brand and the current active page.

After building and deploying to S3, I open the app and it is still using the default React favicon and title (React App). Then, if I navigate to another page and reload, the correct favicon and page title are shown for every subsequent page I navigate to until I go back to Home and reload.

This is the case with every new deployed feature (new components, css).

If I understand correctly, bucket object changes take 24 to reflect in Cloudfront, but in that case, shouldn't everything still be as the default? It seems very strange to me that things are switching back and forth between different versions of the same file in S3.

As far as I can tell, I'm supposed to be invalidating bucket files, or changing the Cache-Control header.
## [7][Why does aws route table have no source but only destination?](https://www.reddit.com/r/aws/comments/fz1hfh/why_does_aws_route_table_have_no_source_but_only/)
- url: https://www.reddit.com/r/aws/comments/fz1hfh/why_does_aws_route_table_have_no_source_but_only/
---
I'm trying to conceptually understand route tables and am a bit confused. When I create a route in a route table, I can only specify destination and target. Why not a source and target for inbound traffic?
## [8][Get info about current VPC/subnet within lambda?](https://www.reddit.com/r/aws/comments/fyqgoj/get_info_about_current_vpcsubnet_within_lambda/)
- url: https://www.reddit.com/r/aws/comments/fyqgoj/get_info_about_current_vpcsubnet_within_lambda/
---
Is there an easy way to determine the current VPC/subnet information from within a lambda function?

I'd like to be able to have a simple function that does some information gathering for me. I can think of a few possible ways to do it, just wondering if anyone knows of anything natively direct.

1. Get current IP address
2. Find and describe ENI with that IP
3. Get info about subnet/vpc that ENI resides in

Thanks
## [9][AWS IPv6 traffic routing through vps](https://www.reddit.com/r/aws/comments/fz0mas/aws_ipv6_traffic_routing_through_vps/)
- url: https://www.reddit.com/r/aws/comments/fz0mas/aws_ipv6_traffic_routing_through_vps/
---
I am wanting to use a VPS firewall to manage IPv6 traffic from AWS LAN to the internet. You can do this in IPv4 using nat.

How could this be done using IPv6? If i set a default ipv6 route to point at a VPS instance, how would I get the VPS firewall instance to then onforward that to the internet and vice versa?
## [10][Lightsail : how to auto-renew SSL certificate ?](https://www.reddit.com/r/aws/comments/fys545/lightsail_how_to_autorenew_ssl_certificate/)
- url: https://www.reddit.com/r/aws/comments/fys545/lightsail_how_to_autorenew_ssl_certificate/
---
Referring to this piece of documentation:  [https://lightsail.aws.amazon.com/ls/docs/en\_us/articles/amazon-lightsail-using-lets-encrypt-certificates-with-wordpress](https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-using-lets-encrypt-certificates-with-wordpress)

I do not need a Load Balancer, and thus don't want to pay for one.

How do I get the SSL certificate to automatically renew itself ? I do not want to have to SSH every 2 months to repeat all these steps...

\`certbot renew\` apparently does the trick. But when I run it, it says that the certificate isn't yet set to expire. This leads me to think that it's a simple command that I'll have to run every 2 months, rather than a command which takes care of running the appropriate set of commands automatically every 2 months.

**How do I get this process automatized ?**
