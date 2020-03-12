# aws
## [1][I publicly tore an AWS engineer's AWS bill apart to find savings, then shared the story with the world.](https://www.reddit.com/r/aws/comments/fh41zb/i_publicly_tore_an_aws_engineers_aws_bill_apart/)
- url: https://www.lastweekinaws.com/blog/an-aws-bill-analysis-changelogs-md/
---

## [2][Amazon Redshift launches pause and resume](https://www.reddit.com/r/aws/comments/fh5dpb/amazon_redshift_launches_pause_and_resume/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/03/amazon-redshift-launches-pause-resume/
---

## [3][S3 Replication Time Control pricing??](https://www.reddit.com/r/aws/comments/fhdypa/s3_replication_time_control_pricing/)
- url: https://www.reddit.com/r/aws/comments/fhdypa/s3_replication_time_control_pricing/
---
I'm trying to estimate time control fees in advance of enabling it and reviewing my bill later.  Crazy I know...

Here's my example use case: 5,000 objects, 50 GB in total, replicated between two regions, say over 1 month.

From the **S3 replication pricing** section on this link:  [https://aws.amazon.com/s3/pricing/](https://aws.amazon.com/s3/pricing/) 

*"When you use S3 Replication Time Control, you also pay a* ***Replication Time Control Data Transfer fee*** *and* ***S3 Replication Metrics*** *charges that are billed at the same rate as Amazon CloudWatch custom metrics."*

The RTC Data Transfer fee is $0.015 per GB = simple

But how do the S3 Replication Metrics translate to CloudWatch custom metrics?  5,000 custom metrics?
## [4][Best way to upload REALLY large files to S3 from my web app](https://www.reddit.com/r/aws/comments/fh6zg6/best_way_to_upload_really_large_files_to_s3_from/)
- url: https://www.reddit.com/r/aws/comments/fh6zg6/best_way_to_upload_really_large_files_to_s3_from/
---
So I have a Django app running on an EC2 instance, with a React front-end, and I'm trying to figure out the best way to upload really large DNA/RNA FASTQ files. I'm fairly new to AWS and it's a bit overwhelming. The file sizes range from a couple of gigs up to 13 GB and users will upload potentially 10-15 of these files at a time. 

First I tried uploading them to my web server sequentially and transferring them to s3 using boto3. That took forever. Now I'm getting a presigned POST request from my server using , returning it to the client, and uploading it directly from the client. That still seems to be pretty slow. It's frustrating that presigned POST requests don't support transfer accelerated endpoints. However, PUT requests do.

&amp;#x200B;

That brings me to two main questions:

1. Is the way I'm currently doing this (uploading directly from the client using a multipart POST) the best/fastest way of doing this?
2. Is a presigned PUT with transfer acceleration faster than a multipart POST? Is it even possible to upload a file with a PUT request to s3? I'm getting mixed messages after scouring the internet a bit. 

Thanks for any insight you could provide!
## [5][Amazon Redshift launches pause and resume](https://www.reddit.com/r/aws/comments/fhd1iq/amazon_redshift_launches_pause_and_resume/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/03/amazon-redshift-launches-pause-resume/
---

## [6][Using AWS Athena To Convert A CSV File To Parquet](https://www.reddit.com/r/aws/comments/fh3tmp/using_aws_athena_to_convert_a_csv_file_to_parquet/)
- url: https://www.cloudforecast.io/blog/Athena-to-transform-CSV-to-Parquet/
---

## [7][Need help understanding the "renewal" of RDS Reserved Instances.](https://www.reddit.com/r/aws/comments/fhceul/need_help_understanding_the_renewal_of_rds/)
- url: https://www.reddit.com/r/aws/comments/fhceul/need_help_understanding_the_renewal_of_rds/
---
I have 2 RDS on-demand instances (1 micro, 1 large) and 1 reserved instance for the large one.

The RI for the large one will expire on April 5th. While I understand that there are no "auto-renewals" of reserved instances, I'm trying to find out if I can purchase the reserved instance now and have it start only after the current RI expires on April 5th.

However, what I read online and in this sub-reddit got me confused. It seems to be that if I purchase a new RI now, it will start immediately and apply the RI hours to my micro instance instead of waiting for the old RI to expire.

If that's the case, then how should I configure the RIs such that it will only start after current RI expire?

TIA for any assistance.
## [8][What's the difference between these two policies using NotPrincipal?](https://www.reddit.com/r/aws/comments/fhccmr/whats_the_difference_between_these_two_policies/)
- url: https://www.reddit.com/r/aws/comments/fhccmr/whats_the_difference_between_these_two_policies/
---
 One has Deny and the other has Allow, but how do they both work with NotPrincipal?

    {
      "Version": "2012-10-17",
      "Statement": [{
          "Effect": "Deny",
          "NotPrincipal": {"AWS": [
              "arn:aws:iam::123456789012:user/John",
              "arn:aws:iam::246810121416:user/Mary"
          ]},
          "Action": "s3:*",
          "Resource": [
              "arn:aws:s3:::BUCKETNAME",
              "arn:aws:s3:::BUCKETNAME/*"
          ]
      }]
    }

and

    {
      "Version": "2012-10-17",
      "Statement": [{
          "Effect": "Allow",
          "NotPrincipal": {"AWS": [
              "arn:aws:iam::123456789012:user/John",
              "arn:aws:iam::246810121416:user/Mary"
          ]},
          "Action": "s3:*",
          "Resource": [
              "arn:aws:s3:::BUCKETNAME",
              "arn:aws:s3:::BUCKETNAME/*"
          ]
      }]
    }
## [9][is AWS required to tell anyone who asks for; the serial numbers of all servers, switching, routers, etc inside their data centers?](https://www.reddit.com/r/aws/comments/fhdper/is_aws_required_to_tell_anyone_who_asks_for_the/)
- url: https://www.reddit.com/r/aws/comments/fhdper/is_aws_required_to_tell_anyone_who_asks_for_the/
---
is it public information or something that needs to be guarded, protected and encrypted and kept away from everyone.
## [10][Verifying c5n/m5n/r5n encryption in-transit?](https://www.reddit.com/r/aws/comments/fh7xyc/verifying_c5nm5nr5n_encryption_intransit/)
- url: https://www.reddit.com/r/aws/comments/fh7xyc/verifying_c5nm5nr5n_encryption_intransit/
---
Hey Everyone,

The c5n/m5n/r5n (and a few other) families of instance types support in-transit encryption at the hardware layer according to AWS' documentation: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/data-protection.html#encryption-transit

Is there any way to actually verify this is happening? I tried a before/after VPC mirror session when moving from m5 to m5n on 2 instances, but that's not really showing any differences. I assume the encryption is happening at too low of a level that we can't see at the AWS control pane layer, but I'd be curious if anyone has any insight into this.

Thanks!
