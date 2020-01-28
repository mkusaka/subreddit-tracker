# aws
## [1][New: AWS Security Documentation by category for different services](https://www.reddit.com/r/aws/comments/eupx5l/new_aws_security_documentation_by_category_for/)
- url: https://docs.aws.amazon.com/security/
---

## [2][An Open Source Alternative to AWS SageMaker](https://www.reddit.com/r/aws/comments/euuy4u/an_open_source_alternative_to_aws_sagemaker/)
- url: https://www.datanami.com/2020/01/27/an-open-source-alternative-to-aws-sagemaker/
---

## [3][Is anyone here using AWS Batch for their projects?](https://www.reddit.com/r/aws/comments/ev51fa/is_anyone_here_using_aws_batch_for_their_projects/)
- url: https://www.reddit.com/r/aws/comments/ev51fa/is_anyone_here_using_aws_batch_for_their_projects/
---
Please share your experience with it so far if you're using it. Like Pros, Cons, etc. Thanks
## [4][Lambda Layer Database - A crowd-sourced database of Useful Hosted lambda layers &amp; layer zips for your serverless projects](https://www.reddit.com/r/aws/comments/euzmf6/lambda_layer_database_a_crowdsourced_database_of/)
- url: https://refinery.io/lambda-layers-database/#/
---

## [5][Handling Lambda Function Throttling](https://www.reddit.com/r/aws/comments/euvqt9/handling_lambda_function_throttling/)
- url: https://www.reddit.com/r/aws/comments/euvqt9/handling_lambda_function_throttling/
---
Hey folks, wanted to share a youtube channel that I've been working on dedicated to providing simple and easy to digest tutorials on various AWS services.

My newest video discusses the implications of Lambda Throttling and how to mitigate its potential impacts to your serverless application.

The video is available here: https://youtu.be/2tmsyweB3k4

Thank you!
## [6][Does AWS AI Services Accuracy Improve over time?](https://www.reddit.com/r/aws/comments/ev3e89/does_aws_ai_services_accuracy_improve_over_time/)
- url: https://www.reddit.com/r/aws/comments/ev3e89/does_aws_ai_services_accuracy_improve_over_time/
---
Hi Guys,

I'm currently performing some audio to text market analysis to compare high-level AI services from AWS, GCP and Azure. An ideally wonderful part of the research is to include how the transcription services improved over the past years.

Based on experience, how do you find AWS high level AI services?

Regards,

Allan
## [7][Any suggestions for my AWS Data Architecture? [Part 2]](https://www.reddit.com/r/aws/comments/ev2itw/any_suggestions_for_my_aws_data_architecture_part/)
- url: https://www.reddit.com/r/aws/comments/ev2itw/any_suggestions_for_my_aws_data_architecture_part/
---
Hi everyone, first of all, thanks for helping me with my [first post](https://www.reddit.com/r/aws/comments/elwt84/any_suggestion_to_my_aws_data_architecture/). 

As you advised, I ditched Airflow and Postgres. 

&amp;#x200B;

DATA FLOW OUTLINE (no precise tech stuff here)

Most of the news publishers have RSS/Atom feeds (something similar to an API if you do not know what it is) to give programmatic access to the latest published news.

By checking each RSS once in a while you can see new records (here's [the one of the NY Times](https://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml)). Therefore, by collecting all the endpoints for all the news providers' feeds that I need, I will be able to collect my data.

Each article itself (to simplify) is a title, published date, author, description, etc.

Let's say I have the place where I store all the articles. After that, I would have to query that with my API.

I plan to store only up to 7-30 days of recent data in the Elasticsearch.

&amp;#x200B;

SOLUTION ARCHITECTURE \[REVISED\]

&amp;#x200B;

https://preview.redd.it/rd4xyzaf0hd41.png?width=3840&amp;format=png&amp;auto=webp&amp;s=5c77c9d3c2715092c9c21db9cd7e22e4d811afb8

1. CloudWatch event triggers the "orchestration lambda"
2. "orchestration lambda" gets the RSS feeds URLs from DynamoDB
3. "orchestration lambda" sends all those RSS URLs to SQS
4. each RSS URL from SQS is picked-up by Lambda that reads and normalizes the feed. 
5. Title+description is md5'ed so we use it as an ID for DynamoDB record. Using `attribute_not_exists` of DDB, only items with unseen ID are promoted, so it solves the deduping problem
6. DDB streams collect up to 100 newly inserted records and send it to Lambda
7. Lambda insert records into Elasticsearch cluster

&amp;#x200B;

My main questions are:

1) The laziest way to monitor it?

2) What if 7. fails? 

3) Any other suggestions?

&amp;#x200B;

Thanks!
## [8][Is it possible to do web automations like Selenium on AWS Lambda?](https://www.reddit.com/r/aws/comments/eunlh5/is_it_possible_to_do_web_automations_like/)
- url: https://www.reddit.com/r/aws/comments/eunlh5/is_it_possible_to_do_web_automations_like/
---
I need to do web automations for a site with no API or web hooks. I can't use selenium because it happens at arbitrary times so I need something on the cloud.
## [9][(NLB) How does health checks works for target group?](https://www.reddit.com/r/aws/comments/eutpbm/nlb_how_does_health_checks_works_for_target_group/)
- url: https://www.reddit.com/r/aws/comments/eutpbm/nlb_how_does_health_checks_works_for_target_group/
---
I am struggling to understand why my targets are unhealthy in the target group.

I have an NLB (Network Load Balancer) listens to port 6443 and forwards the request to targets in the target group. There are 3 EC2 instances attached to the target group. And the target group uses TCP to send a request to each registered target to check its health status.

I am allowing any incoming traffic to port 6443 in the security group. I don't see anything unusual in the Network ACL. NLB and the EC2 instances are all under the same subnet. I don't know what else am I missing here...

I have the Terraform code [here](https://github.com/shoukoo/kubernetes/blob/master/kubernetes-the-hard-way-aws/terraform/main.tf)
## [10][How to permanently delete EC2 instance?](https://www.reddit.com/r/aws/comments/ev1ta5/how_to_permanently_delete_ec2_instance/)
- url: https://www.reddit.com/r/aws/comments/ev1ta5/how_to_permanently_delete_ec2_instance/
---
Hi. I created an AWS account a month and a half ago to test some of its features.

To this point I had only used S3 buckets, CloudWatch and CloudFront, I thought I was going light with their limits, however today I got an email from them telling me that I was at 85% of my free tier limit, so I checked and it looks like I had an EC2 instance running. I didn't recall creating it but maybe I did by accident so I deleted it, however after a while I checked again and there was a new one, and each time I delete it another one appears.

I googled my issue but could only find a forum post from a few years ago where they mention Beanstalk, which I already verified I had not started.

I do not want to get charged and would not like to end up deleting my account, does anyone have a clue on what might be happening?
