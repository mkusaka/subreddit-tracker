# aws
## [1][Webscraper on steroids, using 2,000 Lambda invokes to scan 1,000,000 websites in under 7 minutes.](https://www.reddit.com/r/aws/comments/gd6xss/webscraper_on_steroids_using_2000_lambda_invokes/)
- url: /r/Python/comments/gcq18f/a_serverless_web_scraper_built_on_the_lambda/
---

## [2][I unknowingly left EC2 instances running on an old account last year and accumulated $3,700 in charges. Does Amazon pursue/sell these debts? Do they file it against your credit report?](https://www.reddit.com/r/aws/comments/gcxbar/i_unknowingly_left_ec2_instances_running_on_an/)
- url: https://www.reddit.com/r/aws/comments/gcxbar/i_unknowingly_left_ec2_instances_running_on_an/
---
So yes I know I messed up here. I was using AWS sometime last year to mess with linux VMs on higher end hardware than I have available, and messing with Plex on there.

I stopped messing with it maybe around ~10 months ago due to other unrelated reasons. Before I switched it off I was having trouble encoding on t3 small and similar instances on plex, and was periodically switching them to e.g. the larger m5 machines.

Anyway it looks (well this is what I guess) like I must have left an instance on a  much more powerful machine before the last time I stopped using it for ~10 months. At this time I also changed my email address to a custom domain, so any email notifications didn't get to me. They didn't bother sending any actual real life mail.

I wanted to use AWS again today and signed in, only to find my account has been suspended with $3,700 worth of bills. These were accumulating at around $700/month. I don't know why they didn't suspend the account sooner, and let the debt reach $3,700 over several months,, but they did.

I have spoke to support and submitted a request to have the bill amended/dropped, but am obviously worried it will not.

My question is, if they don't drop them, do they actually try to chase these debts, and at this value? Do they take people to court, or sell their debt to 3rd party companies?

Also do they file the unpaid bills on your credit report?
## [3][How to deploy web apps to AWS using S3 &amp; CloudFront](https://www.reddit.com/r/aws/comments/gdbihm/how_to_deploy_web_apps_to_aws_using_s3_cloudfront/)
- url: https://www.reddit.com/r/aws/comments/gdbihm/how_to_deploy_web_apps_to_aws_using_s3_cloudfront/
---
In this article you can learn how to deploy your website and frontend applications on AWS using AWS S3, CloudFront and Route 53. We’re going to be using S3 to store our source code and assets, CloudFront as a CDN to distribute our website and provide HTTPS support, and Route53 to create a custom DNS record that points to CloudFront distribution. Lets jump right through it:

[https://medium.com/@khaledosman/deploying-a-web-app-on-aws-using-s3-cloudfront-45f77cd652c6](https://medium.com/@khaledosman/deploying-a-web-app-on-aws-using-s3-cloudfront-45f77cd652c6)
## [4][Is it recommended to use Docker with AWS Elastic beanstalk?](https://www.reddit.com/r/aws/comments/gdbhmd/is_it_recommended_to_use_docker_with_aws_elastic/)
- url: https://www.reddit.com/r/aws/comments/gdbhmd/is_it_recommended_to_use_docker_with_aws_elastic/
---
I'm deciding how to deploy my Nodejs API to AWS Elastic beanstalk. I have 2 options: deploy it as a normal nodejs service or deploy the api dockerized.

In which situations is it more recomendable to use docker over a normal nodejs service? (Giving the fact that I will use Elastic beanstalk)
## [5][Setting up a CI/CD pipeline for a ReactJS app using GitLab and AWS CDK](https://www.reddit.com/r/aws/comments/gd0lj8/setting_up_a_cicd_pipeline_for_a_reactjs_app/)
- url: https://www.reddit.com/r/aws/comments/gd0lj8/setting_up_a_cicd_pipeline_for_a_reactjs_app/
---
Created a sample CI/CD ReactJS + AWS CDK project to deploy a website to S3, fronted by CloudFront. All of the resources are created with CDK (including Cognito and lambda functions). I'm using GitLab for the repo and CI/CD. The link to the project is in the description of the video.

https://youtu.be/6OC6AqF-IMc
## [6][Confused on how to go e my IAM user permissions without access to using my credit card](https://www.reddit.com/r/aws/comments/gdaoug/confused_on_how_to_go_e_my_iam_user_permissions/)
- url: https://www.reddit.com/r/aws/comments/gdaoug/confused_on_how_to_go_e_my_iam_user_permissions/
---
Hello, I am towards the end of completing an application and the company has access to my aws account, I was charged a dollar when opening it and today was charged a dollar for another server in Mumbai where they're based off of. What permission can I remove to not give them ny access for purchases as I do not want to worry about receiving an expensive bill. Thank you!
## [7][Trigger Appsync mutation from lambda function](https://www.reddit.com/r/aws/comments/gdangk/trigger_appsync_mutation_from_lambda_function/)
- url: https://www.reddit.com/r/aws/comments/gdangk/trigger_appsync_mutation_from_lambda_function/
---
Trigger Appsync mutation from lambda function

Hi,

so  I keep getting this error message when I try to trigger a mutation  through lambda when I add/modify/delete items from DynamoDB : "message" :  "Invalid request, \`query\` can't be null.",

So  what  i am trying to do is whenever a modification is done directly on  dynamodb, the subscribed users will be notified with the changes.

I have created the new mutation with None type data source. I tested it out directly on the query console and it works fine.

I  have also created the lambda function based on Python that is able to  retrieve the dynamodb streams and have tested it out with cloudwatch.

Now the issue is when I try to do a http post request from the lambda, i get:

error: MalformedHttpRequestException

message: invalid request, \`query\` can't be null.

\-----------------------------------------------------------------------------------------------------------------------------------------------------

name of mutation: addTodo

the data i am sending over post:

{'operationName':  'addTodo', 'variables': {'id': '400', 'name': 'some name',   'description': 'some description', 'query': 'mutation addTodo($id:  ID,$name: String,  $description: String) {addTodo(id: $id, name: $name,  description:  $description) {id name description}}'}}
## [8][AD Connector with third party vpn?](https://www.reddit.com/r/aws/comments/gda6vs/ad_connector_with_third_party_vpn/)
- url: https://www.reddit.com/r/aws/comments/gda6vs/ad_connector_with_third_party_vpn/
---
I have a VPN tunnel using ec2 vms to establish communication from on prem (where ad is located) to AWS.

Can i use AD connector without having to configure the AWS vpn?
## [9][Amazon S3 permissions question](https://www.reddit.com/r/aws/comments/gd548h/amazon_s3_permissions_question/)
- url: https://www.reddit.com/r/aws/comments/gd548h/amazon_s3_permissions_question/
---
I'm a beginner developer and I have a basic question that I can't seem to find a concrete answer to online. Let's say I allow users to upload images and each image has a series of users who are allowed to view it. What is the best practice in node for securing these images to allow only the permitted users to view them? Should I just add my AWS credentials to the server and request the correct image and pass the entire image from the server to the client? Or should I generate a time-limited, publicly accessible URL and pass that down to the client so that the entire image object doesn't pass through my server? Just a bit confused on how node interacts with S3 when users are allowed to upload private photos that are only to be accessed by certain users.
## [10][Lambda provisioned concurrency with Aurora Serverless - should I use the Data API?](https://www.reddit.com/r/aws/comments/gcwtqt/lambda_provisioned_concurrency_with_aurora/)
- url: https://www.reddit.com/r/aws/comments/gcwtqt/lambda_provisioned_concurrency_with_aurora/
---
I'm researching using an AWS serverless stack instead of Heroku.  I'm building an API that's called synchronously, so I'd like to make architecture decisions that result in consistently fast responses.

My budget supports provisioned concurrency and never pausing Aurora Serverless.

My understanding is:

* With Aurora Serverless, I don't need to worry about my lambdas exceeding a connection limit.
* With Provisioned Concurrency, I don't need to worry about the time it takes to establish a database connection.
* If I can pre-establish a database connection, running queries directly will always be faster than using the Data API.

Putting together all these points, I believe I should use traditional database connections to Aurora Serverless instead of the Data API.  This should be fastest, and provisioned concurrency will allow me to avoid any cold starts.

Am I thinking about this right?

Also - it appears I can do all of this with High Availability, with the exception that Aurora Serverless has an unknown failover delay: [https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.how-it-works.html#aurora-serverless.failover](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.how-it-works.html#aurora-serverless.failover)

Is that the primary caveat?

thank you!
