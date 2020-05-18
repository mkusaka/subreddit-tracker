# aws
## [1][TIL AWS has tooling to stop/start instances - Scheduler CLI](https://www.reddit.com/r/aws/comments/gls34w/til_aws_has_tooling_to_stopstart_instances/)
- url: https://www.reddit.com/r/aws/comments/gls34w/til_aws_has_tooling_to_stopstart_instances/
---
https://docs.aws.amazon.com/solutions/latest/instance-scheduler/appendix-a.html

I can't help but think this is perhaps only useful for dev/staging environments.
## [2][CDK for production workloads](https://www.reddit.com/r/aws/comments/glm4k0/cdk_for_production_workloads/)
- url: https://www.reddit.com/r/aws/comments/glm4k0/cdk_for_production_workloads/
---
I was recently getting into learning the AWS CDK and I am super excited about it. The elegant approach to provision infrastructure from python code is amazing.

But since it is pretty new and maybe does not offer all features, I was wondering: is the CDK actually used by organizations for real production environments?
## [3][[HELP] Cannot connect to RDS from Fargate](https://www.reddit.com/r/aws/comments/gm0ps5/help_cannot_connect_to_rds_from_fargate/)
- url: https://www.reddit.com/r/aws/comments/gm0ps5/help_cannot_connect_to_rds_from_fargate/
---
Hi, I'm new to AWS and I'm stuck with this issue for 2 weeks now.

[https://stackoverflow.com/questions/61869022/cannot-connect-to-rds-from-fargate](https://stackoverflow.com/questions/61869022/cannot-connect-to-rds-from-fargate)

Can someone please help me with this?
## [4][AWS tutorials from an ex-AWS engineer](https://www.reddit.com/r/aws/comments/glzvt5/aws_tutorials_from_an_exaws_engineer/)
- url: https://www.reddit.com/r/aws/comments/glzvt5/aws_tutorials_from_an_exaws_engineer/
---
Hi everyone,

I'm an ex-AWS software engineer (moved on to try something different).

I've been thinking of writing up some tutorials on using AWS to do things which I think far more people should be aware (and not scared) of. For example, hosting a static site on S3 + Cloudfront.

Is there a specific tutorial you'd be interested in?
## [5][MediaConvert Server Side Encryption of HLS Stream](https://www.reddit.com/r/aws/comments/gly5ym/mediaconvert_server_side_encryption_of_hls_stream/)
- url: https://www.reddit.com/r/aws/comments/gly5ym/mediaconvert_server_side_encryption_of_hls_stream/
---
Here's the MediaConvert code that encrypts the segments:

&amp;#x200B;

'Encryption' =&gt; \[

	'EncryptionMethod' =&gt; 'AES128',

	'InitializationVectorInManifest' =&gt; 'EXCLUDE',

	'OfflineEncrypted' =&gt; 'DISABLED',

	'StaticKeyProvider' =&gt; \[

		'StaticKeyValue' =&gt; 'ecd0d06eaf884d8226c33928e87efa33',

		'Url' =&gt; '[https://Example.com/keyinfo](https://Example.com/keyinfo)',

	\],

	'Type' =&gt; 'STATIC\_KEY', //'SPEKE'|'STATIC\_KEY'

\],

&amp;#x200B;

[Example.com/keyinfo](https://Example.com/keyinfo) is a text file with ONLY the following 2 lines of text:

&amp;#x200B;

[https://Example.com/key](https://Example.com/key)

key

&amp;#x200B;

&amp;#x200B;

Then, at [Example.com/key](https://Example.com/key) is a text file containing just 1 single line of text:

&amp;#x200B;

ecd0d06eaf884d8226c33928e87efa33

&amp;#x200B;

(and that's the same as "StaticKeyValue" from the earlier MediaConvert job)

&amp;#x200B;

And when I play this video using VideoJS, it is able to get to the main playlist (.m3u8), it is able to locate the keyinfo file (I can see it in the browser logs), but then it fails when trying to open the first .ts segment which is part 1 of the encrypted video.

&amp;#x200B;

So what do you think is wrong here? I am happy to pay someone to troubleshoot this for me.

&amp;#x200B;

Thanks in advance for your help!
## [6][LightSail instance suddenly stopped working](https://www.reddit.com/r/aws/comments/gly4tj/lightsail_instance_suddenly_stopped_working/)
- url: https://www.reddit.com/r/aws/comments/gly4tj/lightsail_instance_suddenly_stopped_working/
---
Seems like there is a recently undocumented introduction of the [bursting "feature"](https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-frequently-asked-questions-faq) for LightSail that completely render my VPS useless?

Does someone else get into the same situation or It's only my VPS that is faulty?

https://preview.redd.it/d3g7wi0uphz41.png?width=912&amp;format=png&amp;auto=webp&amp;s=54f29f97fb3b67c7a426467f083dbad8a77d53a6

&amp;#x200B;

https://preview.redd.it/9nl3gfdeqhz41.png?width=909&amp;format=png&amp;auto=webp&amp;s=9f0df67f5a096c8defe69cd9601ca483ff391821

Edit : Seems like they always behaved that way. ( Like a burstable T2. It's just now that they decided to include their metric calculation into the metric dashboard ). Wierd thing that I never experienced instances being taken down before but now they do...
## [7][AWS: CloudFormation — using Conditions, Fn::Equals, and Fn::If — an example](https://www.reddit.com/r/aws/comments/gly2s1/aws_cloudformation_using_conditions_fnequals_and/)
- url: https://www.reddit.com/r/aws/comments/gly2s1/aws_cloudformation_using_conditions_fnequals_and/
---
The author has aCloudFormation stack with VPC Peerings.

The task: add an ability to chose if CloudFormation have to create the peering mentioned above — or skip this step.

The solution: use the AWS CloudFormation Conditions: will add a new parameter VPCPeeringCreate which will accept a true value false from a Jenkins job and then depending on this value CloudFormation will decide if need to create such a peering and related resources - the peering itself and two Routes.

The article: https://medium.com/setevoy4/aws-cloudformation-using-conditions-fn-equals-and-fn-if-an-example-bf147336a25f?source=friends_link&amp;sk=30143435f35dc77da8a400e49a7190fe
## [8][When does Lambda@Edge stop being counted as running for billing purposes?](https://www.reddit.com/r/aws/comments/gltr1g/when_does_lambdaedge_stop_being_counted_as/)
- url: https://www.reddit.com/r/aws/comments/gltr1g/when_does_lambdaedge_stop_being_counted_as/
---
Let's say I'm using Lambda@Edge in front of an S3 Bucket that is used for streaming large videos or large file downloads.  Would the Lambda@Edge function be counted as running for the full duration of the streaming video or large download?  Or is it only considered running while it's doing it's initial run to determine the origin and once the origin is set, it stops running at that point (for purposes of billing)?

I hope its not considered running for the duration of the data transfer.  That could get insanely expensive.
## [9][Glue Crawler creating thousands of tables](https://www.reddit.com/r/aws/comments/glr75g/glue_crawler_creating_thousands_of_tables/)
- url: https://www.reddit.com/r/aws/comments/glr75g/glue_crawler_creating_thousands_of_tables/
---
I'm struggling a bit with AWS Glue Crawler and wondering if anyone can help set me in the right direction. I have thousands of xml files on S3 that are daily snapshots of data that I'm trying to convert to 2 partitioned parquet tables (to query with Athena).

**Example S3 files:**

* s3://bucketname/2020/04/26/07/31/index.xml
* s3://bucketname/2020/04/26/08/08/records/491.xml
* s3://bucketname/2020/05/17/05/00/index.xml
* s3://bucketname/2020/05/17/05/53/records/116.xml
* s3://bucketname/2020/05/17/05/14/records/45.xml

**General pattern:**

* s3://bucketname/yyyy/MM/dd/HH/mm/index.xml
* s3://bucketname/yyyy/MM/dd/HH/mm/records/&lt;recordId&gt;.xml

&amp;#x200B;

I hoped the default Crawler would discover there are two schemas (index &amp; records), identify the partitions, and create only 2 tables. But when I run it with include\_path = s3://bucketname/2020, it creates thousands of tables, one for each file I think.

Should I be structuring my files differently? Or is there context I should be providing to the Crawler?
## [10][renaming Elastic Beanstalk Windows host &amp; domain join / remove](https://www.reddit.com/r/aws/comments/glw8gh/renaming_elastic_beanstalk_windows_host_domain/)
- url: https://www.reddit.com/r/aws/comments/glw8gh/renaming_elastic_beanstalk_windows_host_domain/
---
just looking for advice on the best way to achieve;

-rename of Windows host produced via Elastic Beanstalk, they are all the same, which doesn't work when you want multiple hosts joining domain
-join of the newly renamed host to AWS managed Microsoft AD, what is the best way to do this ?
-removal from domain, what is the best way to remove computer from domain on termination for cleanup ?
