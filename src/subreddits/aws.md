# aws
## [1][What a typical 100% Serverless Architecture looks like in AWS!](https://www.reddit.com/r/aws/comments/gomeoy/what_a_typical_100_serverless_architecture_looks/)
- url: https://medium.com/serverless-transformation/what-a-typical-100-serverless-architecture-looks-like-in-aws-40f252cd0ecb
---

## [2][Rocketchat on AWS - problem with route 53](https://www.reddit.com/r/aws/comments/gp39fz/rocketchat_on_aws_problem_with_route_53/)
- url: https://www.reddit.com/r/aws/comments/gp39fz/rocketchat_on_aws_problem_with_route_53/
---
Hello my friends,

I just installed rocketchat on AWS, and now I'm struggling with route 53.  I have a domain and I want to point to aws. It was parked on godaddy, so I change the DNS for the DNS provided by aws and I'm waiting more tha 2 hours and nothing happens. 

Maybe I'm doing something wrong? Can someone please help me?
## [3][Cloudfront VS Cloudflare which one is better](https://www.reddit.com/r/aws/comments/gp42jb/cloudfront_vs_cloudflare_which_one_is_better/)
- url: https://www.reddit.com/r/aws/comments/gp42jb/cloudfront_vs_cloudflare_which_one_is_better/
---
I have a website, I am thinking which CDN service should i use

Can someone please advise me which one should I use, I wanna speed up my site. 

My site with images(not much, one page around 2-3 images)

I am using lightsail as a hosting. 

Or should I use both?
## [4][AWS Lambda and OracleDB integration](https://www.reddit.com/r/aws/comments/gp2te8/aws_lambda_and_oracledb_integration/)
- url: https://www.rehanvdm.com/serverless/an-unexpected-journey-with-lambda-oracledb/index.html
---

## [5][How can I connect my web application to query data from AWS mysql RDS instance?](https://www.reddit.com/r/aws/comments/gp280a/how_can_i_connect_my_web_application_to_query/)
- url: https://www.reddit.com/r/aws/comments/gp280a/how_can_i_connect_my_web_application_to_query/
---
I have created instance in mysql RDS. Currently I'm using mysql workbench to import and query the data. Now, I want to create a web application where I can select field values and query the data using the user interface. 

Please suggest how to do it. Thanks in advance..
## [6][Strange behavior with an NLB Target Group](https://www.reddit.com/r/aws/comments/gozxn1/strange_behavior_with_an_nlb_target_group/)
- url: https://www.reddit.com/r/aws/comments/gozxn1/strange_behavior_with_an_nlb_target_group/
---
I'm not sure what is causing this behavior.  For some reason an NLB based Target Group is registering a bunch of targets (ports) that are incorrect until it eventually hits upon the correct port and it becomes healthy.  This obviously causes a very long delay in the service becoming responsive to clients.  Observing the task container during the above event shows the container having the same SINGLE host port assigned to it, so I'm not sure what is causing this strange behavior.

I don't see this same thing with ALB based Target Group.  What should I be looking at to fix this?
## [7][Is there a way to copy (not move) objects from S3 to Glacier directly?](https://www.reddit.com/r/aws/comments/gotvsu/is_there_a_way_to_copy_not_move_objects_from_s3/)
- url: https://www.reddit.com/r/aws/comments/gotvsu/is_there_a_way_to_copy_not_move_objects_from_s3/
---
Maybe I'm missing something, but the Lifecycle rules only allow moving objects from S3 to Glacier so the files will be gone from S3. Is that correct?

In my case, the data needs to be accessible on S3 for reading and backed up on Glacier. Here's what happens:

1. server sends pre-signed post links to client
2. client app uploads stuff to S3 directly via pre-signed post links
3. server gets notified via SNS
4. server downloads file from S3 and uploads it to Glacier **&lt; this part sucks**
5. server queues up some processing to do on the uploaded files (without modifying them)

The Glacier serves as a backup and a long-term archive for client data.

How can I copy files from S3 to Glacier without having to download them to the server?

Is this the right way to use Glacier or should I be using something else?
## [8][Is there a way to open the services menu with a keyboard shortcut?](https://www.reddit.com/r/aws/comments/goi6og/is_there_a_way_to_open_the_services_menu_with_a/)
- url: https://www.reddit.com/r/aws/comments/goi6og/is_there_a_way_to_open_the_services_menu_with_a/
---
Seriously, if I want to navigate to a new service I have to take my hands off the keyboard and click the menu icon before I can search! 

Am I just an idiot? What’s the point of auto focusing the search bar if you still have to use the mouse to open the services menu in the first place?
## [9][Are Your AWS S3 Buckets Secure ?](https://www.reddit.com/r/aws/comments/goolnp/are_your_aws_s3_buckets_secure/)
- url: https://www.reddit.com/r/aws/comments/goolnp/are_your_aws_s3_buckets_secure/
---
S3 is a commonly used AWS service which is also prone to access policy misconfigurations by users, resulting in security issues. My 2 cents on probable steps that you could take to define a bucket policy which ensures that your data is safe !  


[https://towardsdatascience.com/are-your-aws-s3-buckets-secure-5cc07f63788?source=friends\_link&amp;sk=3938670e8684e35c35d3661c4c4feddf](https://towardsdatascience.com/are-your-aws-s3-buckets-secure-5cc07f63788?source=friends_link&amp;sk=3938670e8684e35c35d3661c4c4feddf)
## [10][Iterate a value in DynamoDB](https://www.reddit.com/r/aws/comments/gopzbk/iterate_a_value_in_dynamodb/)
- url: https://www.reddit.com/r/aws/comments/gopzbk/iterate_a_value_in_dynamodb/
---
Hello, I was wondering what the best way to do something like this would be:

for an object in dynamoDB such as Item:{int numLeft: 10}, I would like to have a "claim" lambda function that updates the value of numleft to be one less than it currently is, unless numLeft =0, then I want to return an error message to tell the user that no more can be claimed. 

&amp;#x200B;

I know how to update in with node.js, but I only know how by reading the table for the item, then looking at the value, then calling the update value function with the new iterated down value. If there is a way to do this in a single query to DynamoDB, I would love to know it! 

Thanks!
