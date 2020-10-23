# aws
## [1][AWS Wish List 2020](https://www.reddit.com/r/aws/comments/jbw85e/aws_wish_list_2020/)
- url: https://www.reddit.com/r/aws/comments/jbw85e/aws_wish_list_2020/
---
&amp;#x200B;

AWS always releases a bunch of features, sometimes everyday or atleast once a week. Here is my wish list of the features I want to see as a part of AWS infrastructure

1: AWS Managed Proxy Server(Rather than spinning own squid server)

2: EBS replication across different availability zones(Possible? Legal constraints?)

3: Multi-region VPC(Possible? Legal constraints?)

4: UI to debug boot issues(Better then EC2 Get Instance Screenshot and Instance logs)

5: Support tagging for every individual service(It's improving)

6: VPC endpoints support for every service (EKS?) 

7: EC2 instance live migration

8: Display AWS Cli  while resource creation(Similar to GCP)

9: Cost calculation while resource creation(AWS start supporting(for example, RDS) this feature but not for every service

10: More features in App Mesh(Circuit breaker, Rate Limiting)

P.S: Not sure if some features are already available, but if something is missing, please feel free to add
## [2][Introducing Amazon SNS FIFO – First-In-First-Out Pub/Sub Messaging](https://www.reddit.com/r/aws/comments/jg99wv/introducing_amazon_sns_fifo_firstinfirstout/)
- url: https://aws.amazon.com/blogs/aws/introducing-amazon-sns-fifo-first-in-first-out-pub-sub-messaging/
---

## [3][Finally! Amazon CloudFront announces support for public key management through IAM user permissions for signed URLs and signed cookies](https://www.reddit.com/r/aws/comments/jgdb37/finally_amazon_cloudfront_announces_support_for/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/10/cloudfront-iam-signed-url/
---

## [4][Amazon Prime Day 2020 – Powered by AWS](https://www.reddit.com/r/aws/comments/jg55ig/amazon_prime_day_2020_powered_by_aws/)
- url: https://aws.amazon.com/blogs/aws/amazon-prime-day-2020-powered-by-aws/
---

## [5][AWS CloudFormation now supports increased limits on five service quotas](https://www.reddit.com/r/aws/comments/jg8x9t/aws_cloudformation_now_supports_increased_limits/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/10/aws-cloudformation-now-supports-increased-limits-on-five-service-quotas/
---

## [6][rain - A development workflow tool for working with AWS CloudFormation](https://www.reddit.com/r/aws/comments/jg37gs/rain_a_development_workflow_tool_for_working_with/)
- url: https://github.com/aws-cloudformation/rain
---

## [7][Introducing Amazon SNS FIFO – First-In-First-Out Pub/Sub Messaging](https://www.reddit.com/r/aws/comments/jg8uoh/introducing_amazon_sns_fifo_firstinfirstout/)
- url: https://aws.amazon.com/blogs/aws/introducing-amazon-sns-fifo-first-in-first-out-pub-sub-messaging/
---

## [8][Amazon Prime Day 2020 – Powered by AWS](https://www.reddit.com/r/aws/comments/jg5auh/amazon_prime_day_2020_powered_by_aws/)
- url: https://aws.amazon.com/blogs/aws/amazon-prime-day-2020-powered-by-aws/
---

## [9][OCR for consistently structured ID cards](https://www.reddit.com/r/aws/comments/jgm6x6/ocr_for_consistently_structured_id_cards/)
- url: https://www.reddit.com/r/aws/comments/jgm6x6/ocr_for_consistently_structured_id_cards/
---
My application needs to extract information from customer ID cards. We're only operating in one geography and so I know exactly the format of the information (coordinates of where all the text fields are, etc).

I've had a look at Textract which isn't bad, but I think performance could be improved if I could somehow feed in the domain knowledge about the ID card structure. Is there a way to do that, or should I be looking at other solutions? I'd rather avoid having to do the image manipulation/normalisation ourselves if possible.
## [10][AWS RDS - Move from one master cluster to multiple master cluster.](https://www.reddit.com/r/aws/comments/jgm0i2/aws_rds_move_from_one_master_cluster_to_multiple/)
- url: https://www.reddit.com/r/aws/comments/jgm0i2/aws_rds_move_from_one_master_cluster_to_multiple/
---
Hi, currently at work we use a simple cluster of Mysql with AWS RDS, which has one master only. Now because of business requirement (we want more "write" instances, we have write-heavy database) . Is there any way to reconfigure our existing cluster ? Or do we have to create a new cluster with multiple masters and migrate the data to there ?
## [11][Origin and Referrer request header key-value pairs not working properly - Is this standard?](https://www.reddit.com/r/aws/comments/jglyfp/origin_and_referrer_request_header_keyvalue_pairs/)
- url: https://www.reddit.com/r/aws/comments/jglyfp/origin_and_referrer_request_header_keyvalue_pairs/
---
When I check the chrome developer tools to understand the request header key-value pairs are being sent with my requests, I notice that sometimes Referrer displays and Origin never displays.

For my API I need the Origin and Referrer Headers to always display for the GET request. Is there a specific reason why this is happening?

For Example - 

I access example.com/get/resource from salesforce.com/abc

What I am expecting in the header request is:

origin: salesforce.com

Referrer: salesforce.com/abc

For some reason, origin does not display in the chrome developer tools and referrer fails to show up when I load the website from a new tab.

Is this expected? I ask, because I want to use either of these fields in my authorizer function.
