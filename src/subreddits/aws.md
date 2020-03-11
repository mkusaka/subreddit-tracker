# aws
## [1][Amazon EKS now supports Kubernetes 1.15](https://www.reddit.com/r/aws/comments/fgo244/amazon_eks_now_supports_kubernetes_115/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/03/amazon-eks-now-supports-kubernetes-version-1-15/
---

## [2][New AWS OS for running containers](https://www.reddit.com/r/aws/comments/fgnhjd/new_aws_os_for_running_containers/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/03/announcing-bottlerocket-a-new-open-source-linux-based-operating-system-optimized-to-run-containers/
---

## [3][How to download files from S3 with Lambda?](https://www.reddit.com/r/aws/comments/fguobe/how_to_download_files_from_s3_with_lambda/)
- url: https://www.reddit.com/r/aws/comments/fguobe/how_to_download_files_from_s3_with_lambda/
---
I want to grant users access to some PDF files stored in a private S3 bucket. Upon request of the file I call a Lambda function that checks whether they're allowed to access the file (e.g. they can't have it if they're over 18 or if their name starts with an A - those are just dumb examples of the pre-access checks I do). I wanted to use this Lambda to also get the file from the bucket and return it in the response but I don't think I can given the body payload limit of 6MB.I read about pre-signed URLs but what bothers me is that anyone with the link can access the file, granted they do it before it expires.Is there a way to download my S3 files and return them within the Lambda function? If not, is there another way to do so that would guarantee the download link cannot be shared?
## [4][Control your AWS Lambda with Provisioned Concurrency](https://www.reddit.com/r/aws/comments/fgvrgv/control_your_aws_lambda_with_provisioned/)
- url: https://epsagon.com/blog/development/control-your-aws-lambda-with-provisioned-concurrency/
---

## [5][AWS Summit San Francisco 2020 Cancelled](https://www.reddit.com/r/aws/comments/fgmzqw/aws_summit_san_francisco_2020_cancelled/)
- url: https://www.reddit.com/r/aws/comments/fgmzqw/aws_summit_san_francisco_2020_cancelled/
---
Just got the email today. 
 

Important Event Update

Our top priority is the well-being of our customers, partners, and employees. After careful review of the current situation with COVID-19 in San Francisco and listening to the guidance provided by the local authorities, Amazon Web Services has made the decision to cancel the AWS Summit San Francisco 2020, which was scheduled for April 14 at the Moscone Center. If you booked a hotel reservation through the AWS Summit room block at the Marriott Marquis San Francisco you will need to take action to cancel it to avoid charges. You can do this via the confirmation email you received from Marriott Marquis San Francisco or by calling them directly at 1-877-622-3056.

We are reimagining the AWS Summit San Francisco to be a fully digital experience in May. In this new format, you can still hear about exciting new product announcements from AWS leaders, dive into educational and technical content, and engage with AWS experts. More information will be released in the coming weeks.
## [6][AWS Lambda and Node.js 12: Support and Benchmark](https://www.reddit.com/r/aws/comments/fgx209/aws_lambda_and_nodejs_12_support_and_benchmark/)
- url: https://epsagon.com/blog/aws-lambda-and-node-js-12-support-and-benchmark/
---

## [7][AWS S3 bucket storage type and sync](https://www.reddit.com/r/aws/comments/fguafp/aws_s3_bucket_storage_type_and_sync/)
- url: https://www.reddit.com/r/aws/comments/fguafp/aws_s3_bucket_storage_type_and_sync/
---
I'm using CLI to sync assets over from one bucket to another (back up in different region) . The main bucket has Intelligent Tiering and the back up has Standard-IA.

When I run a sync between the 2 buckets, the assets appearing in the backup bucket are appearing as standard storage class rather than standard-IA.

Why's this?
## [8][AWS CloudWatch - Part 1/3: Logs and Insights](https://www.reddit.com/r/aws/comments/fgwk7o/aws_cloudwatch_part_13_logs_and_insights/)
- url: https://epsagon.com/blog/aws-cloudwatch-logs-and-insights/
---

## [9][AppSpec.yml not found issue when deploying to Windows EC2 from CodeDeploy](https://www.reddit.com/r/aws/comments/fgvb2g/appspecyml_not_found_issue_when_deploying_to/)
- url: https://www.reddit.com/r/aws/comments/fgvb2g/appspecyml_not_found_issue_when_deploying_to/
---
Hi,

I've had this running months ago, so I know it works, but have created a new EC2 instance to deploy my code and stuck at the first hurdle.

My Deployment Details runs as follows:

* Application Stop - succeeded
* Download Bundle - succeeded
* BeforeInstall - Failed

Upon looking at the failed event, I get:

&gt; The CodeDeploy agent did not find an AppSpec file within the unpacked revision directory at revision-relative path "appspec.yml". The revision was unpacked to directory "C:\\ProgramData/Amazon/CodeDeploy/57f7ec1b-0452-444e-840c-4deb4566e82d/d-WH9HTZAW0/deployment-archive", and the AppSpec file was expected but not found at path "C:\\ProgramData/Amazon/CodeDeploy/57f7ec1b-0452-444e-840c-4deb4566e82d/d-WH9HTZAW0/deployment-archive/appspec.yml". Consult the AWS CodeDeploy Appspec documentation for more information at http://docs.aws.amazon.com/codedeploy/latest/userguide/reference-appspec-file.html 

Thing is, if I jump onto my EC2 and copy and paste the full path, sure enough I see the YML file, along with the files that were in a ZIP file within my S3 bucket, so they've been successfully sent to the EC2 and unzipped.

So I'm sure it's not a permissions things, the connection is being clearly made, and the S3 Bucket, CodeDeploy and my EC2 are all happy.

I read somewhere about changing the AppSpec.yml file to "appspec.yml", "AppSpec.yaml", "appspec.yaml", and still nothing works.

Anything obvious to try out?
## [10][Permissions Needed For Waiter In Lambda?](https://www.reddit.com/r/aws/comments/fgv9v5/permissions_needed_for_waiter_in_lambda/)
- url: https://www.reddit.com/r/aws/comments/fgv9v5/permissions_needed_for_waiter_in_lambda/
---
I'm in the middle of migrating a Python3 lambda from a fairly unrestrictive account, to a more restricted account.  I'm running into an issue where my code:

```python
s3Waiter = s3Conn.get_waiter( 'object_exists' )
```

Triggers an exception in the restricted account:

```python
[ERROR] WaiterError: Waiter ObjectExists failed: Forbidden
Traceback (most recent call last):
  File "/var/task/lambda_function.py", line 165, in lambda_handler
    s3Waiter.wait( Bucket = s3Bucket, Key = s3Key )
  File "/var/runtime/botocore/waiter.py", line 53, in wait
    Waiter.wait(self, **kwargs)
  File "/var/runtime/botocore/waiter.py", line 310, in wait
    raise WaiterError(
```

Are there any resources that detail what permissions are required for such actions?
