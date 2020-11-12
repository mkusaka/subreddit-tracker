# aws
## [1][re:Invent registration is now open](https://www.reddit.com/r/aws/comments/jkenu3/reinvent_registration_is_now_open/)
- url: https://register.virtual.awsevents.com/
---

## [2][Announcing AWS Glue DataBrew – A Visual Data Preparation Tool That Helps You Clean and Normalize Data Faster](https://www.reddit.com/r/aws/comments/jsgtxo/announcing_aws_glue_databrew_a_visual_data/)
- url: https://aws.amazon.com/blogs/aws/announcing-aws-glue-databrew-a-visual-data-preparation-tool-that-helps-you-clean-and-normalize-data-faster/
---

## [3][Accessing AWS Hong Kong region from Mainland China](https://www.reddit.com/r/aws/comments/jssxy3/accessing_aws_hong_kong_region_from_mainland_china/)
- url: https://www.reddit.com/r/aws/comments/jssxy3/accessing_aws_hong_kong_region_from_mainland_china/
---
I need a company website to be accessible from Mainland China with good performance. At the moment the version hosted in the eu-central region has latency and bandwidth problems with Mainland China. We have an ICP license and an AWS China account, but we would like to avoid using a different domain and dealing with separated infrastructure, so I was thinking about using the region in Honk Kong for traffic coming from China, but someone told me that performance might be poor anyway because of China's GFW (Hong Kong is outside of it). Does anyone has info on this? Thanks
## [4][AWS Glacier Restores not working - new web UI](https://www.reddit.com/r/aws/comments/jsszb0/aws_glacier_restores_not_working_new_web_ui/)
- url: https://www.reddit.com/r/aws/comments/jsszb0/aws_glacier_restores_not_working_new_web_ui/
---
Ever since the UI change recently I am having random restore problems. Restoring singular files from Glacier via the web-UI goes through the steps to 'initiate restore' but it doesn't seem to actually work.

Some files are fine to restore though. Same bucket? Anyone else?
## [5][Athena Engine v2](https://www.reddit.com/r/aws/comments/jslqhs/athena_engine_v2/)
- url: https://docs.aws.amazon.com/athena/latest/ug/engine-versions-reference.html#engine-versions-reference-0002
---

## [6][Which of the new interfaces is the worst?](https://www.reddit.com/r/aws/comments/jsgvwm/which_of_the_new_interfaces_is_the_worst/)
- url: https://www.reddit.com/r/aws/comments/jsgvwm/which_of_the_new_interfaces_is_the_worst/
---
I have been using AWS for the past 10+ years, and mostly championing it the entire time. Most of my work is CLI/API based, but I do use the console for various activities, like spot checking data in DynamoDB. The new Route53 UI is terrible, I think it is probably the worst of them. I ended up on the new S3 page today,  I am still digesting it to figure out if it is just a change, or really bad. I took a peek at the new DynamoDB page, it is really bad, basically unusable. What are your thoughts?
## [7][Want to Implement approval process for IAM users](https://www.reddit.com/r/aws/comments/jstaxb/want_to_implement_approval_process_for_iam_users/)
- url: https://www.reddit.com/r/aws/comments/jstaxb/want_to_implement_approval_process_for_iam_users/
---
 

Good Day!!!

i want to implement that all developer(IAM users) in AWS are undergone through approval process and get there account up by using active directory .

Can anyone suggest , how it would be implement. Which service i need to use here.

Regards Pratap
## [8][SSL certificate](https://www.reddit.com/r/aws/comments/jsseiv/ssl_certificate/)
- url: https://www.reddit.com/r/aws/comments/jsseiv/ssl_certificate/
---
Do we set up  separate SSL certificates and sticky sessions for the EC2 instance when  is to be accessed from different platform types (Windows, macOS, IOS, and Android)?
## [9][What on earth has happened to the spot instance console?](https://www.reddit.com/r/aws/comments/jsrzoa/what_on_earth_has_happened_to_the_spot_instance/)
- url: https://www.reddit.com/r/aws/comments/jsrzoa/what_on_earth_has_happened_to_the_spot_instance/
---
It used to be check the price, add an instance, choose the type and done. Now it's asking me what type of workload I want... I just want a single instance so I went for "Flexible" and removed all the instances except the m3.medium which is the one I want

I don't see anywhere to add my pricings anymore, went ahead anyway and it just terminates immediately. I just run a few things that I don't particularly care if they go down as they're personal projects...

Tried it in terraform and it's fine, so what *on earth* have they done to the really simple console? and *why*? More of a gripe since I'll just use terraform from now on but why have they made this way more complicated than it needs to be?
## [10][check_nt: Could not parse arguments || Windows Server 2016 Host Error](https://www.reddit.com/r/aws/comments/jsr8lp/check_nt_could_not_parse_arguments_windows_server/)
- url: https://www.reddit.com/r/aws/comments/jsr8lp/check_nt_could_not_parse_arguments_windows_server/
---
Hi,

I have Nagios Linux Server, I have added this Windows host to Nagios Server but I am getting below error

https://preview.redd.it/5dykvlb3rry51.png?width=1147&amp;format=png&amp;auto=webp&amp;s=f860218bd7410158d06bbe4e1cafd13252983fba

Nagios Server Configuration as below:

====================================================================================

**Windows.cfg  File**

&amp;#x200B;

define host {

use                     windows-server          

host\_name               [ec2-35-154-15-163.ap-south-1.compute.amazonaws.com](https://ec2-35-154-15-163.ap-south-1.compute.amazonaws.com)

alias                   My Windows Server       

address                 [35.154.15.163](https://35.154.15.163)

}

# Few Services Format as below:

&amp;#x200B;

define service {

use                     generic-service

host\_name               [ec2-35-154-15-163.ap-south-1.compute.amazonaws.com](https://ec2-35-154-15-163.ap-south-1.compute.amazonaws.com)

service\_description     NSClient++ Version

check\_command           check\_nt!nagios@2020sg!CLIENTVERSION

}

&amp;#x200B;

define service {

use                     generic-service

host\_name               [ec2-35-154-15-163.ap-south-1.compute.amazonaws.com](https://ec2-35-154-15-163.ap-south-1.compute.amazonaws.com)

service\_description     CPU Load

check\_command           check\_nt!nagios@2020sg!CPULOAD!20,40,60

}

&amp;#x200B;

====================================================================================

**commands.cfg File**

define command {

   command\_name    check\_nt

command\_line    $USER1$/check\_nt -H $HOSTADDRESS$ -p 12489 -v  $ARG1$ $ARG2$

}

=================================================================================

I am using this on AWS EC2 instances. Please help me to understand the cause of problem.
## [11][Real-world AWS Timestream ingest performance numbers?](https://www.reddit.com/r/aws/comments/jsgn9x/realworld_aws_timestream_ingest_performance/)
- url: https://www.reddit.com/r/aws/comments/jsgn9x/realworld_aws_timestream_ingest_performance/
---
I've been working for the last week to test AWS Timestream using some of their [supplied sample code for continuous data ingestion](https://github.com/awslabs/amazon-timestream-tools/tree/master/tools/continuous-ingestor) (Python/boto3) and the [Time-Series Benchmark Suite](https://github.com/timescale/tsbs) provided by Timescale (which uses the Timestream Go SDK).

No matter what combination of settings and VMs I use (lots of smaller clients w/fewer threads, limited XL clients w/lots of threads, short memory tier window, large memory tier window), I'm rarely seeing more than \~500 metrics/second from any instance (and sometimes worse). These same tests have performed significantly better on any mix of the databases that TSBS can be used with (TimescaleDB, Influx, Mongo, etc.). Obviously I can't compare the AWS ingestor on anything other than Timestream.

It seems I'm not the only one [seeing similar performance](https://crate.io/a/amazon-timestream-first-impressions/), but this just feels so impossibly poor that I want to ensure things are setup correctly.

Does anybody have any better experience or tips on what might be the issue here? What are some real-world performance metrics others could share?

(full disclosure, I work at Timescale and we're trying to honestly evaluate Timestream based on client questions)
