# aws
## [1][This is what happens when you move from Intel to Graviton2](https://www.reddit.com/r/aws/comments/j8w9l9/this_is_what_happens_when_you_move_from_intel_to/)
- url: https://i.redd.it/k2oeu3r99ds51.jpg
---

## [2][Can SNS deliver to Kinesis Firehose?](https://www.reddit.com/r/aws/comments/j956g7/can_sns_deliver_to_kinesis_firehose/)
- url: https://www.reddit.com/r/aws/comments/j956g7/can_sns_deliver_to_kinesis_firehose/
---
Without the use of Lambda?

Or is there a better service that achieve a fan out approach to data received?

Something like this

```
API Gateway =&gt; (x) =&gt; Kinesis Firehose =&gt; S3
                                  =&gt; Lambda
                                  =&gt; Kinesis =&gt; Lambda
```

Where (x) is a service that fans out to other services?
## [3][I wrote a tool to blacklist an AZ from ASG to avoid launching new instances in the faulty AZ. Followup to current issues with AWS Availability zones.](https://www.reddit.com/r/aws/comments/j8laxb/i_wrote_a_tool_to_blacklist_an_az_from_asg_to/)
- url: https://github.com/awsiv/asg-az-update
---

## [4][Permissions error trying to run Prometheus on EKS (Fargate only) with EFS](https://www.reddit.com/r/aws/comments/j93djz/permissions_error_trying_to_run_prometheus_on_eks/)
- url: https://stackoverflow.com/questions/64287430/permissions-error-trying-to-run-prometheus-on-aws-eks-fargate-only-with-efs
---

## [5][AWS hackathon abruptly cancelled?](https://www.reddit.com/r/aws/comments/j8tfna/aws_hackathon_abruptly_cancelled/)
- url: https://www.reddit.com/r/aws/comments/j8tfna/aws_hackathon_abruptly_cancelled/
---
A few months ago [this](https://dev.to/willydavidjr/aws-mission-to-mars-hackathon-10lp) was announced. AWS Mission to MARS hackathon - build whatever you want using AI, ML or Robotics services to win 65k in prizes. But yesterday I'm suddenly getting an email it was cancelled. Here's the [website](https://www.awsmissiontomars.com/#/event).

Anyone know anything? Kinda sucks to make noise about such a cool hackathon competition and then just cancel without giving a good reason.
## [6][What kind of database allows tools like AWS X-RAY to function?](https://www.reddit.com/r/aws/comments/j8um1t/what_kind_of_database_allows_tools_like_aws_xray/)
- url: https://www.reddit.com/r/aws/comments/j8um1t/what_kind_of_database_allows_tools_like_aws_xray/
---
I was testing AWS X-RAY out and it learned how to setup annotations for filtering and group by. I haven't collected a lot of data yet as I'm just toying around with a dozen traces or so. However I wouldn't doubt that x-ray can scale to millions of traces.
The main functionality I'm curious is the ability to have what seems to be a schemaless structure (50 arbitrary annotations) capable of offering filter (like a where clause) and grouping (like group by clause) and aggregation (avg, sum, count, etc).
I imagine building this on top of DynamoDB could be extremely challenging and a relational database would have performance issues on scale.
What's the ideal type of database to power applications like X-RAY and what makes it a great database for this type of process?

Thanks for any input!!
## [7][Is the cloudfront distribution automatically setted on the free trial when created? Can't find something related to the on the distro creation page](https://www.reddit.com/r/aws/comments/j922g7/is_the_cloudfront_distribution_automatically/)
- url: https://www.reddit.com/r/aws/comments/j922g7/is_the_cloudfront_distribution_automatically/
---

## [8][RCA for SA-EAST-1](https://www.reddit.com/r/aws/comments/j8w2h1/rca_for_saeast1/)
- url: https://www.reddit.com/r/aws/comments/j8w2h1/rca_for_saeast1/
---
Hello,

Does AWS publish public RCA’s for issues like the one that happened today with EC2 in South America?
## [9][Stitching Terraform and Ansible together via AWS SSM, plus Graviton2](https://www.reddit.com/r/aws/comments/j8xfwq/stitching_terraform_and_ansible_together_via_aws/)
- url: https://github.com/chadgeary/pihole
---

## [10][This AWS Quck Start tutorial needs to know my "Allowed Bastion External Access CIDR" but I don't know what to put here.](https://www.reddit.com/r/aws/comments/j8wnbo/this_aws_quck_start_tutorial_needs_to_know_my/)
- url: https://www.reddit.com/r/aws/comments/j8wnbo/this_aws_quck_start_tutorial_needs_to_know_my/
---
As the title goes i'm following this aws quick start tutorial [__here__](https://docs.aws.amazon.com/quickstart/latest/linux-bastion/step2.html). I'm at a point in the tutorial where I have to fill out [__this option here__](https://imgur.com/a/AYgsHUL). AWS defines this as...  

&gt; CIDR block that’s allowed SSH external access to the bastion hosts. We recommend that you set this value to a trusted CIDR block. For example, you might want to restrict access to your corporate network.  

My understanding is I need to provide a range of ip's but in CIDR format that I'll be ssh'n into from. Is this my public facing ip of my home network (e.g. the one showed by sites like [__ippig.com__](http://ippig.com/)? Or is this the internal gateway of my modem?  

Lastly, whether its the public facing ip or the internal gateway of my modem this is just one ip. So I'm still confused how I'm to put this in CIDR format as a range of ips. Can someone clear my confusion up with an example of what I'm to put here please.
