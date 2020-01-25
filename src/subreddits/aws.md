# aws
## [1][Original Snowball retirement, Snowball Edge only from 4th Feb](https://www.reddit.com/r/aws/comments/et8d99/original_snowball_retirement_snowball_edge_only/)
- url: https://www.reddit.com/r/aws/comments/et8d99/original_snowball_retirement_snowball_edge_only/
---
Anyone know any more about the original Snowball retirement?  I haven't been able to find any announcement of it or any more details, but the following notification appears in the AWS console:

&gt; On February 4th, 2020, the first-generation 48 TB and 80 TB Snowball devices will be retired from the Snowball Service. Devices ordered by February 3rd, 2020 will be fulfilled. These 48 TB and 80 TB devices will no longer be orderable on February 4th, 2020. The Snowball Edge Storage Optimized device replaces the first-generation Snowball device for data migration requirements.

This is a bit frustrating for us, as we found the Edge to be more awkward to use in our specific workflows, and we're not geared up to start working with Edges again, certainly not with a week and a half notice.
## [2][Export Aurora or RDS snapshots to Amazon S3](https://www.reddit.com/r/aws/comments/et3mxm/export_aurora_or_rds_snapshots_to_amazon_s3/)
- url: https://www.reddit.com/r/aws/comments/et3mxm/export_aurora_or_rds_snapshots_to_amazon_s3/
---
You can now export Aurora or RDS snapshots to Amazon S3 for analytics or long term retention with just a few clicks on the Amazon RDS Management Console or using the AWS SDK or CLI. No need to write any code or provision any instances! 

https://aws.amazon.com/about-aws/whats-new/2020/01/announcing-amazon-relational-database-service-snapshot-export-to-s3/
## [3][AWS engineer uploads customer keys, passwords to GitHub](https://www.reddit.com/r/aws/comments/et08r0/aws_engineer_uploads_customer_keys_passwords_to/)
- url: https://www.reddit.com/r/aws/comments/et08r0/aws_engineer_uploads_customer_keys_passwords_to/
---
https://www.theregister.co.uk/2020/01/23/aws_engineer_customer_credentials_github/

IAM secret key/secrets were available for five hours before being taken down.

As a precaution we’ve rotated all our IAM keys , no word yet from AWS
## [4][AwsApiChanges.info: Updated API methods which reflect recent AWS updates](https://www.reddit.com/r/aws/comments/esuw0d/awsapichangesinfo_updated_api_methods_which/)
- url: https://awsapichanges.info
---

## [5][What's the difference between CloudFront and ELB?](https://www.reddit.com/r/aws/comments/eta7e3/whats_the_difference_between_cloudfront_and_elb/)
- url: https://www.reddit.com/r/aws/comments/eta7e3/whats_the_difference_between_cloudfront_and_elb/
---
Both of them distribute the traffic based on geolocation. So what's the difference?

**Edit:** And how does S3 fit into this? Isn't it distributed also?
## [6][Running a Lambda in response to updates to SSM parameters owned by Amazon?](https://www.reddit.com/r/aws/comments/et9tv9/running_a_lambda_in_response_to_updates_to_ssm/)
- url: https://www.reddit.com/r/aws/comments/et9tv9/running_a_lambda_in_response_to_updates_to_ssm/
---
Hello - I run an ECS cluster and it's been incredibly annoying to have to update the ECS-optimized AMIs in my launch templates/configurations every time a new AMI is published. I wrote a Lambda to deal with changing out the AMI in my ASGs but I'd like to trigger the Lambda every time Amazon updates the recommended AMI in Parameter Store.   


Using the documentation on this provided by AWS ([https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-paramstore-cwe.html](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-paramstore-cwe.html)), I set up a CloudWatch event rule that looks like this: 

`{`

  `"detail-type": [`

`"Parameter Store Change"`

  `],`

  `"source": [`

`"aws.ssm"`

  `],`

  `"detail": {`

`"name": [`

`"/aws/service/ecs/optimized-ami/amazon-linux/recommended"`

`],`

`"operation": [`

`"Update"`

`]`

  `}`

`}`

However, I'm not sure this is going to work as intended. I'm skeptical because I assume the events for that parameter may not be in my event bus but would anyone know for sure if this is going to work?
## [7][What language are you using to Develop in AWS CDK?](https://www.reddit.com/r/aws/comments/et8w84/what_language_are_you_using_to_develop_in_aws_cdk/)
- url: /r/aws_cdk/comments/et8sig/what_language_are_you_using_to_develop_in_cdk/
---

## [8][Results of the 2019 AWS Container Security Survey](https://www.reddit.com/r/aws/comments/esv8f2/results_of_the_2019_aws_container_security_survey/)
- url: https://aws.amazon.com/blogs/containers/results-of-the-2019-aws-container-security-survey/
---

## [9][[ECS] Detect and generate a cloudwatch Alarm when a Task is killed or restarted](https://www.reddit.com/r/aws/comments/et80n6/ecs_detect_and_generate_a_cloudwatch_alarm_when_a/)
- url: https://www.reddit.com/r/aws/comments/et80n6/ecs_detect_and_generate_a_cloudwatch_alarm_when_a/
---
Hi,

I'm trying to set an Alarm using Cloudwatch to detect when a Task is killed in an ECS Cluster.

I followed this  [https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cloudwatch-metrics.html#cw\_running\_task\_count](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cloudwatch-metrics.html#cw_running_task_count)  to set up a metric to track the number of tasks running by a service.

This is my alarm : 

**Namespace**        : AWS/ECS  
**MetricName**      : CPUUtilization  
**ServiceName**     : my\_service  
**ClusterName**     : my\_cluster  
**Statistic**               : Sample count  
**Period**                  : 1minute  
**Conditions** :   
   **Threshold type** : Static  
**Whenever CPUUtilization is**... : Lower Than 1  
 

But it's not working and doesn't generate an Alarm as expected. I think that this is because if the task is killed, it's recreated again automatically and quickly under 1 minute (the period set in the Alarm).

I tried to change the Period to less than a minute but AWS says ***Only a period greater than 60s is supported for metrics in the "AWS/" namespaces*** 

 

So is there a way to detect if a task is killed ?

Thanks
## [10][Assign public IP range to the auto scalable containers in ECS?](https://www.reddit.com/r/aws/comments/et6i88/assign_public_ip_range_to_the_auto_scalable/)
- url: https://www.reddit.com/r/aws/comments/et6i88/assign_public_ip_range_to_the_auto_scalable/
---
Hi, When i spin up an ECS cluster, it assigns public IP dynamically to all the containers it creates.

Can we configure something so that the IP assigned are within a particular range? Something like public CIDR?
