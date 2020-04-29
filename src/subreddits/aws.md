# aws
## [1][Getting started on AWS - A Beginner Roadmap](https://www.reddit.com/r/aws/comments/g9xvi1/getting_started_on_aws_a_beginner_roadmap/)
- url: https://www.reddit.com/r/aws/comments/g9xvi1/getting_started_on_aws_a_beginner_roadmap/
---
I've been noticing quite a few threads about folks getting started on AWS but not knowing where to start.

I put together a youtube walkthrough of some of the most critical AWS services. In addition, I've also provided some background in the services you need to know for backend, frontend, big data, etc.

Video is available here:  https://youtu.be/lTyqzyk86f8

Thanks!
## [2][AWS Architecture: Architecting Twitter using AWS services](https://www.reddit.com/r/aws/comments/g9xh3g/aws_architecture_architecting_twitter_using_aws/)
- url: https://www.reddit.com/r/aws/comments/g9xh3g/aws_architecture_architecting_twitter_using_aws/
---
Hey guys, 

Some of you asked for a systems design session, so I picked Twitter and went with it. Here's a very simple way to approach architecting some of the most visible features of Twitter using AWS services:

[https://youtu.be/QHJzO5IQCw0](https://youtu.be/QHJzO5IQCw0)

I didn't prep specifically to put myself in the same spot as an interviewee, so I was thinking out-loud, a lot. But this is the way I would approach answering this question. 

Let me know if you find this useful (subscribe if you do) and also let me know what you'd want me to improve upon/elaborate/etc. 

Vladimir
## [3][AWS SDK Credentials Provider slow at start up](https://www.reddit.com/r/aws/comments/ga5j89/aws_sdk_credentials_provider_slow_at_start_up/)
- url: https://www.reddit.com/r/aws/comments/ga5j89/aws_sdk_credentials_provider_slow_at_start_up/
---
I have a local issue where the AWS SDK (Java) is invoking http://169.254.170. 2 at startup and the timeout is slow. 

Every time EC2ContainerCredentialsProviderWrapper is invoked there is a 2 second delay before the call times out. The rest of my team experience a much shorter delay so it is clearly local to my PC. I cannot figure out how to either stop the SDK making the call or reduce the timeout. Searching for solutions just turns up a huge list of search results of people have generic DHCP issues where windows cannot get an IP address.

The endpoint the SDK invokes is not configurable, it is a static final String (seen here  [https://github.com/aws/aws-sdk-java/blob/1.11.415/aws-java-sdk-core/src/main/java/com/amazonaws/auth/ContainerCredentialsProvider.java](https://github.com/aws/aws-sdk-java/blob/1.11.415/aws-java-sdk-core/src/main/java/com/amazonaws/auth/ContainerCredentialsProvider.java)). So it cannot be changed.

In the vain hope someone else has had this problem and come up with a solution I thought I'd ask here ...
## [4][Cloudfront + S3 - expiration of files](https://www.reddit.com/r/aws/comments/ga86ac/cloudfront_s3_expiration_of_files/)
- url: https://www.reddit.com/r/aws/comments/ga86ac/cloudfront_s3_expiration_of_files/
---
Hi,

I'm trying to understand how Cloudfront decides how long do the files stay in the cache. My situation is following: 

\- I have an S3 origin. My origin does not send any Cache-Control headers, the files do not have any metadata set. 

I understand that I can either leave the expiration times to the origin or I can customize the cache behavior by specifying min,max and default TTL. But what happens if I leave the expiration to the origin (Use Origin Cache Headers) and the origin does not send Cache-control headers? Will Cloudfront use the min,max and default values for TTLs even if they are greyed out? Or it does not cache the files at all? I have the situation shown on the picture, my origin does not send anything but the file does get cached, I don't know for how long. Is it the value of default TTL (86400?). I'm just trying to understand where it gets the value :) Thanks!

&amp;#x200B;

https://preview.redd.it/yqksz90jyqv41.png?width=854&amp;format=png&amp;auto=webp&amp;s=535bf055fc4a401d35f8309a2e82d240bddabdf6
## [5][Is AWS Key Management Service (KMS) of any personal use?](https://www.reddit.com/r/aws/comments/ga828l/is_aws_key_management_service_kms_of_any_personal/)
- url: https://www.reddit.com/r/aws/comments/ga828l/is_aws_key_management_service_kms_of_any_personal/
---
KMS creates a key for you and securely protects it. Some questions:

&amp;#x200B;

*  How secure is KMS? Surely US government and amazon have your key, don't they?
* Is this service for companies, or do individuals also find some use cases for it? I am thinking of using the key stored in amazon hardware for encryption of computer files. It costs more to put it on Yubikey and key back ups are headache.
## [6][How to enable vim keyboard mode in the Lambda console?](https://www.reddit.com/r/aws/comments/ga7lbb/how_to_enable_vim_keyboard_mode_in_the_lambda/)
- url: https://www.reddit.com/r/aws/comments/ga7lbb/how_to_enable_vim_keyboard_mode_in_the_lambda/
---
How can I enable vim keyboard mode in the Lambda console? I clearly remember that it worked fine a few years ago. Currently, vim mode works on Cloud 9 but not on the Lambda Console itself.
## [7][AWS S3 error: Please reduce your request rate](https://www.reddit.com/r/aws/comments/ga7j55/aws_s3_error_please_reduce_your_request_rate/)
- url: https://www.reddit.com/r/aws/comments/ga7j55/aws_s3_error_please_reduce_your_request_rate/
---
So, I have a few lambda functions chained together that perform a huge number of read/writes to a few S3 buckets. Today I started observing the following errors in the CloudWatch logs for one of the lambdas:

    com.amazonaws.services.s3.model.AmazonS3Exception: Please reduce your request rate. (Service: Amazon S3; Status Code: 503; Error Code: SlowDown; Request ID: 10F99AA6A9598781; S3 Extended Request ID: REDACTED), S3 Extended Request ID: REDACTED

The thing is that I have run this pipeline multiple times before and never faced this error. This only started today. I have looked up this error online but there is very scant information available.

Have any of you encountered this error before? What workaround can be used? I cannot afford to get such an error in production.
## [8][How to make a Cronjob to get data via a request and use this data to generate predictions via a machine learning model](https://www.reddit.com/r/aws/comments/ga7b5l/how_to_make_a_cronjob_to_get_data_via_a_request/)
- url: https://www.reddit.com/r/aws/comments/ga7b5l/how_to_make_a_cronjob_to_get_data_via_a_request/
---
As the title says, I'd like to make a cronjob that would execute daily, first this cronjob would get data from different APIs and use them to generate a prediction for a given day, and those predictions should be available by an API request, so that the model is run only once per day, and so are the requests to the different APIs, which AWS services would be handy when I am trying to achieve something like this?

Thanks!
## [9][Setting S3 Object ACL (canonical ID) on an S3 object uploaded by another account](https://www.reddit.com/r/aws/comments/ga6enh/setting_s3_object_acl_canonical_id_on_an_s3/)
- url: https://www.reddit.com/r/aws/comments/ga6enh/setting_s3_object_acl_canonical_id_on_an_s3/
---
Hi,

I'm trying to write objects to an S3 bucket from another account (secure logging) and then change the ACL on the uploaded objects to add the canonical ID of the account holding the S3 bucket.

I can do the cross account S3 upload fine as expected the file has no canonical id set (wether or not I add the -cannnedACLName switch with "private"|"bucket-owner-full-control" etc.).

I then do a Get ACL's on the bucket, save that as a variable and apply that to the S3 bucket object I've uploaded.  

If I do the above as a user on the account that owns the S3 bucket it all works fine, if I assume a role and do the same, the S3 uploads by the S3 set ACL fails with "Set-S3ACL: Access Denied"

I'm getting the same error both in my Lambda and locally (assuming the Lambda exec role after adding in a new trust for my user).

In PS the code is:

    $Acl = Get-S3Acl -BucketName $Bucketname
    Write-S3Object -BucketName $Bucketname -Key $File -File "c:\temp\$File" -CannedACLName Private -Credential $creds
    Set-S3ACL -BucketName $Bucketname -Key $File -Grant $acl.grants -OwnerID $acl.Owner.Id -Credential $creds    


The bucket policy on the S3 allows the following actions on the bucket and bucket/* for the Lambda exec role (which I impersonate testing locally).

    "s3:PutObject",
    "s3:PutObjectAcl",
    "s3:GetBucketAcl",
    "s3:GetObjectAcl"
## [10][How to proper setup HTTPS on EC2 [Help]](https://www.reddit.com/r/aws/comments/ga02mv/how_to_proper_setup_https_on_ec2_help/)
- url: https://www.reddit.com/r/aws/comments/ga02mv/how_to_proper_setup_https_on_ec2_help/
---
I'm running into problems while trying to run my API that's hosted on AWS EC2 through a HTTPS protocol.

The API runs normally **without** the ELB setup, however, after trying to configure (I follow the recommended steps), I get the **502 Bad Gateway** message.

Here's my configuration:

* AWS EC2 (t3a.small) running a docker container of my ExpressJS app listening on port 3000;
* Security group with http:80 and https:443 open;
* ACM that covers the following domains (mydomain.com, \*.mydomain.com);
* ELB listening to ports: http:80, https:443, https:3000;
* Route 53 with my hosted zone containing the A-type record with the ELB DNS value;

I'm running into problems while trying to run my API that's hosted on AWS EC2 through an HTTPS protocol.

http://ec2-ip-address.zone.compute.amazonaws.com:3000/api/

**Now**

https://api.mydomain.com:443/api/{resourceName}

I'm running into problems while trying to run my API that's hosted on AWS EC2 through an HTTPS protocol.
