# aws
## [1][Week of Aug 17th - What are your favorite AWS Tips?](https://www.reddit.com/r/aws/comments/ibdue5/week_of_aug_17th_what_are_your_favorite_aws_tips/)
- url: https://www.reddit.com/r/aws/comments/ibdue5/week_of_aug_17th_what_are_your_favorite_aws_tips/
---
Share your best AWS tips with the community!
## [2][Application and Classic Load Balancers are adding defense in depth with the introduction of Desync Mitigation Mode (and open sourcing HTTP Desync Guardian)](https://www.reddit.com/r/aws/comments/ibk57y/application_and_classic_load_balancers_are_adding/)
- url: https://www.reddit.com/r/aws/comments/ibk57y/application_and_classic_load_balancers_are_adding/
---
Elastic Load Balancing has released a new layer-7 security feature called “Desync Mitigation Mode” on both the Application and Classic Load Balancer. This feature offers protection from HTTP vulnerabilities arising from Desync, while maintaining the availability and performance of your application. More details here – https://aws.amazon.com/about-aws/whats-new/2020/08/application-and-classic-load-balancers-adding-defense-in-depth-with-introduction-of-desync-mitigation-mode/
 
PS: As always, we love hearing from you and take your feedback seriously! Thank you!

The underlying library has also been open-sourced, available on Github here https://github.com/aws/http-desync-guardian
## [3][We are the AWS EC2 Team - Ask the Experts - Aug 21st @ 9AM PT / 12PM ET / 4PM GMT!](https://www.reddit.com/r/aws/comments/iboiel/we_are_the_aws_ec2_team_ask_the_experts_aug_21st/)
- url: https://www.reddit.com/r/aws/comments/iboiel/we_are_the_aws_ec2_team_ask_the_experts_aug_21st/
---
Hey [r/aws](https://www.reddit.com/r/aws)! [u/AmazonWebServices](https://www.reddit.com/u/AmazonWebServices/) here.

The AWS EC2 team will be hosting an Ask the Experts session here **in this thread** to answer any questions you may have about running your workloads on the latest generation Amazon EC2 [M6g](https://aws.amazon.com/ec2/instance-types/m6/), [C6g](https://aws.amazon.com/ec2/instance-types/c6/), and [R6g](https://aws.amazon.com/ec2/instance-types/r6/) instances powered by the new [AWS Graviton2 processors](https://aws.amazon.com/ec2/graviton/). These instances enable up to 40% better price performance over comparable x86-based instances for a wide variety of workloads, including application servers, micro-services, high-performance computing, CPU-based machine learning inference, electronic design automation, gaming, open-source databases, and in-memory caches.

Already have questions? Post them below and we'll answer them starting at 9AM PT on Aug 21, 2020!
## [4][New – AWS Fargate for Amazon EKS now supports Amazon EFS](https://www.reddit.com/r/aws/comments/ibhlsz/new_aws_fargate_for_amazon_eks_now_supports/)
- url: https://aws.amazon.com/blogs/aws/new-aws-fargate-for-amazon-eks-now-supports-amazon-efs/
---

## [5][S3/Athena or S3/Glue/Athena for CSV data processing?](https://www.reddit.com/r/aws/comments/ibzkvd/s3athena_or_s3glueathena_for_csv_data_processing/)
- url: https://www.reddit.com/r/aws/comments/ibzkvd/s3athena_or_s3glueathena_for_csv_data_processing/
---
So my requirements are to perform audit logging in our application. Considering the large volume of data to be expected, we have decided to store the audit logs in a CSV file and store them to a bucket in S3 every hour or so. Then, later when a user requests an audit report, we need to filter these CSV data in the S3 bucket based on a given time frame and return the results for download.

Now, I have observed a couple of ways to do this but two have stuck out:

1. Use Athena to run `select *` queries on the CSV bucket directly.
2. Use Glue to crawl the data in the bucket from time to time and then use Athena to query the Glue data store.

While option 1 is simple, I wonder whether there is any advantage to using option 2. Considering the costs involved, I feel running a crawler frequently might be a bit more costly in addition to Athena's per query billing compared to running Athena solely.

The audit report generation activity is not expected to run very frequently, perhaps once every six months or so. 

So, what would be the best option in this case? The volume is expected to be very large. So, running a query in Athena might take a long time with the costs that come with it. So, will the cost of running a Glue crawler be justified if there are performance improvements running an Athena query?
## [6][Simplify AWS Lambda functions with middleware](https://www.reddit.com/r/aws/comments/ic0weu/simplify_aws_lambda_functions_with_middleware/)
- url: https://startup-cto.net/simplify-aws-lambda-functions-with-middleware/
---

## [7][I'm getting charged although i'm on the AWS free tier](https://www.reddit.com/r/aws/comments/ic0nqt/im_getting_charged_although_im_on_the_aws_free/)
- url: https://www.reddit.com/r/aws/comments/ic0nqt/im_getting_charged_although_im_on_the_aws_free/
---
Hello everyone,

I'm new to AWS and decided to use it for a side project. I'm on the t2.micro EC2 instance with a 10GB EBS volume. 

I just noticed i was charged $4 for last month, and i've already run up a $18 charge for the end of this month. I checked around and landed in billing and found out the charge comes from this `$0.065 per IOPS-month provisioned`. I'm finding this very confusing, please is there a way to resolve this issue and not be charged else is there a way to avoid this in the future.

NOTE: i don't know if this is important but i pulled a docker image in the instance.

Thanks
## [8][For our Cloud IoT users a useful comparison](https://www.reddit.com/r/aws/comments/ibxo7x/for_our_cloud_iot_users_a_useful_comparison/)
- url: https://www.reddit.com/r/aws/comments/ibxo7x/for_our_cloud_iot_users_a_useful_comparison/
---
This article compares AWS IoT Core to Google Cloud IoT Core

 [https://www.netburner.com/learn/iot-cloud-aws-iot-core-vs-google-cloud-iot-core/](https://www.netburner.com/learn/iot-cloud-aws-iot-core-vs-google-cloud-iot-core/)
## [9][How is working at AWS in COVID times?](https://www.reddit.com/r/aws/comments/ibm2ym/how_is_working_at_aws_in_covid_times/)
- url: https://www.reddit.com/r/aws/comments/ibm2ym/how_is_working_at_aws_in_covid_times/
---
I've read Glassdoor reviews about the pace being fast-moving, challenging, and rewarding. I saw reviews saying benefits are not as snazzy as other IT companies. I've read [this post](https://www.reddit.com/r/aws/comments/ash4en/working_for_aws/) commenting on a few of these aspects. 

Can anyone share how their experience has been working at AWS through COVID? Has leadership managed the change well? Any updates to salary or benefits?

Thanks in advance!
## [10][How to understand DynamoDB](https://www.reddit.com/r/aws/comments/ic038i/how_to_understand_dynamodb/)
- url: https://www.reddit.com/r/aws/comments/ic038i/how_to_understand_dynamodb/
---
DynamoDB is gaining popularity based on the amount of questions that people ask here and clients that I talk to. I always try to help by providing feedback and answers questions. While doing so, I decided to write down an article that answer the most common questions regarding DynamoDB. 

The goal is to help understand what should be the mindset  need to work with DynamoDB so you can have a performant application, that is easy to work with and extend.

Since lots of people paint themself In to a corner, which makes it impossible to extend the solution.

https://consulting.0x4447.com/articles/how_to/how-to-understand-dynamodb.html
## [11][AWS EKS and Kubernetes Notebook](https://www.reddit.com/r/aws/comments/ibz1o0/aws_eks_and_kubernetes_notebook/)
- url: https://www.linkedin.com/posts/adnanrashid1_kubernetes-cloud-publiccloud-activity-6701444471592153088-SnQS
---

