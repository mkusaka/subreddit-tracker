# aws
## [1][Amazon Fraud Detector is now Generally Available](https://www.reddit.com/r/aws/comments/hzhzc9/amazon_fraud_detector_is_now_generally_available/)
- url: https://aws.amazon.com/blogs/aws/amazon-fraud-detector-is-now-generally-available/
---

## [2][Video stream service recommendation](https://www.reddit.com/r/aws/comments/hzx43v/video_stream_service_recommendation/)
- url: https://www.reddit.com/r/aws/comments/hzx43v/video_stream_service_recommendation/
---
I'm prototyping a new idea that involves on-demand live stream sessions, with 1-2 hosts and about 20-30 participants for a max of 30mins to 1hour sessions. While, we can manage some sessions with Zoom, its not practical when we want to build a platform for streamers and participants. 

While exploring I found Twilio Programmable Video to be a pretty good option for our needs. But, looking at costs - I'm trying to figure out if AWS has any solutions, and glad they do. 

The problem: Too many AWS Media live stream solutions? 

Media Services:

Kinesis Video Streams

MediaLive

Amazon Interactive Video Service

Does anyone have any recommendations on what I should start looking at? Or any other alternative platforms one would recommend for such use case? Happy to provide more info, if necessary.
## [3][MFA info](https://www.reddit.com/r/aws/comments/hzybgj/mfa_info/)
- url: https://www.reddit.com/r/aws/comments/hzybgj/mfa_info/
---
Hi guys

Recently I created the policy to force MFA to users, but I'm wondering is there any way to display additional info to every user in group for example:
User without MFA getting message or pop-up that he can't do anything until he will configure the mfa?
Users with MFA will not getting any info.
## [4][AWS IAM Footguns &amp; Security Basics](https://www.reddit.com/r/aws/comments/i005rd/aws_iam_footguns_security_basics/)
- url: https://cloud-fundamentals.com/blog/iam-footguns-security-basics/
---

## [5][new r53 interface - love or hate it ?](https://www.reddit.com/r/aws/comments/hzzi82/new_r53_interface_love_or_hate_it/)
- url: https://www.reddit.com/r/aws/comments/hzzi82/new_r53_interface_love_or_hate_it/
---
as the title says, whats your impression of the new interface?
## [6][S3 Eventual Consistency Issues with updating recently archived objects (NoSuchKey errors being thrown)](https://www.reddit.com/r/aws/comments/hzy7ss/s3_eventual_consistency_issues_with_updating/)
- url: https://www.reddit.com/r/aws/comments/hzy7ss/s3_eventual_consistency_issues_with_updating/
---
I have an ETL application, where a prefix is sending object create events to an SQS queue that triggers a series of lambda functions to process the data. We are running into a problem where the S3 object has not fully propagated across all of S3 and a NoSuchKey error is thrown. There is a separate process that will move the original object to a separate prefix should a notification be sent to a dead letter queue if the original process fails (this assist with retrying the failed data).

When data is retried, however, the key original name of the key needs to stay the same. Attempting to retry the data lease to additional NoSuchKey errors being thrown when the SQS queue receives the create object event. In the lambda function that original creates the s3 object to be processed, we are using boto3's upload_fileobj method. Is there a better way to upload this data to ensure the eventual consistency, and is there a separate event trigger we should be using on the bucket (i.e., instead of all object create events, use POST/PUT/multi-part upload)? I have tried using boto3's waiter.wait in our downstream lambda function, and that fixed the issue for about a day but it started creeping up again, to no avail.

In theory, we could send the sqs message directly from the application code after the object is created, but that would make replaying data slightly difficult as now instead of simply being able to move the original object back into the original prefix to be processed, we would have to also manually send the sqs message instead of it being triggered by the bucket creation event.
## [7][EKS 1.13 EOL](https://www.reddit.com/r/aws/comments/hzxx8o/eks_113_eol/)
- url: https://www.reddit.com/r/aws/comments/hzxx8o/eks_113_eol/
---
I have a EKS cluster in 1.13 version. Console  invite me to update my cluster to Kubernetes version 1.15 or higher in order to avoid service interruption.

Can somebody give me the least date to do this update and avoid service interruption ?

Thanks
## [8][us-east-1 DNS issues](https://www.reddit.com/r/aws/comments/hzbcry/useast1_dns_issues/)
- url: https://www.reddit.com/r/aws/comments/hzbcry/useast1_dns_issues/
---
Is anyone else experiencing DNS resolver issues right now in US-east-1?  Started noticing it around 4:45 AM EST.
## [9][Vpc peering two same cidr vpcs with two different subnets?](https://www.reddit.com/r/aws/comments/hzw15o/vpc_peering_two_same_cidr_vpcs_with_two_different/)
- url: https://www.reddit.com/r/aws/comments/hzw15o/vpc_peering_two_same_cidr_vpcs_with_two_different/
---
I have two vpcs that have same cidr: 10.0.0.0/16

In vpc A, I have subnet: 10.0.1.0/24

In vpc B, I have subnet: 10.0.2.0/24

I want two subnet can talk to each other privately, I can only think about vpc peering but they are in vpcs with same cidr. 

how to have two subnets connect to each other? 

Is it possible to peering specific subnet between vpcs that have same cidr?
## [10][My EC2 Microservice gets progressively slower!](https://www.reddit.com/r/aws/comments/hzn3mh/my_ec2_microservice_gets_progressively_slower/)
- url: https://www.reddit.com/r/aws/comments/hzn3mh/my_ec2_microservice_gets_progressively_slower/
---
Hi. So I am using EC2 primarily to generate thumbnails and previews for files. And it gets progressively slower for some reason. I restarted the process. Not much of a difference. Restarted the app and suddenly the time reduced to almost 1/3rd. I am failing to find the cause. I am keeping track of memory and cpu usage and logging them. Also, monitoring if there are any zombie processes. Nothing. Only thing I noticed is that after restart, the disk usage reduced from 8GB to like 5GB. Also, memory usage reduced to around 50% from 60%.

The library I am using is a python library. And it uses ImageMagick, libreoffice, ffmpeg to generate previews.  Anyone can give me why it could be happening? Is there any aws solution I can use or a some tool I can install on my instance that help me solve this.

EDIT: I am using t3a.small for now. Usually the load is pretty much negligible hence I felt that a more powerful instance is a waste. uploads happen in bulk and thats when the problem happens. Everything comes to crawl and I have to manually restart the instance. I am thinking about creating another instance and only starting it when there is an incoming job.

EDIT: Thanks. I have turned on T3 unlimited. Hopefully that solves it.
