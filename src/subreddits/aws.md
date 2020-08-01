# aws
## [1][New – Using Amazon GuardDuty to Protect Your S3 Buckets](https://www.reddit.com/r/aws/comments/i1czaa/new_using_amazon_guardduty_to_protect_your_s3/)
- url: https://aws.amazon.com/blogs/aws/new-using-amazon-guardduty-to-protect-your-s3-buckets/
---

## [2][Is serverless cheaper for your use case? Find out with this calculator.](https://www.reddit.com/r/aws/comments/i19di3/is_serverless_cheaper_for_your_use_case_find_out/)
- url: https://medium.com/serverless-transformation/is-serverless-cheaper-for-your-use-case-find-out-with-this-calculator-2f8a52fc6a68
---

## [3][CodeBuild — what is the meaning of these IAM policy statements?](https://www.reddit.com/r/aws/comments/i1mwsg/codebuild_what_is_the_meaning_of_these_iam_policy/)
- url: https://www.reddit.com/r/aws/comments/i1mwsg/codebuild_what_is_the_meaning_of_these_iam_policy/
---
For creating a CodeBuild project with Terraform, the TF docs provide [this example](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/codebuild_project#example-usage) which includes an IAM role policy containing several statements. Two of them, which involve network interfaces, are unclear to me. What do they do in the context of the CodeBuild project? And what resource should `"arn:aws:ec2:us-east-1:123456789012:network-interface/*"` be attached to? I'm a bit in the dark on this so any amount of explanation will probably be useful to me. Thanks!

    [
      {
        "Effect": "Allow",
        "Action": [
          "ec2:CreateNetworkInterface",
          "ec2:DescribeDhcpOptions",
          "ec2:DescribeNetworkInterfaces",
          "ec2:DeleteNetworkInterface",
          "ec2:DescribeSubnets",
          "ec2:DescribeSecurityGroups",
          "ec2:DescribeVpcs"
        ],
        "Resource": "*"
      },
      {
        "Effect": "Allow",
        "Action": [
          "ec2:CreateNetworkInterfacePermission"
        ],
        "Resource": [
          "arn:aws:ec2:us-east-1:123456789012:network-interface/*"
        ],
        "Condition": {
          "StringEquals": {
            "ec2:Subnet": [
              "${aws_subnet.example1.arn}",
              "${aws_subnet.example2.arn}"
            ],
            "ec2:AuthorizedService": "codebuild.amazonaws.com"
          }
        }
      }
    ]
## [4][CloudFormation validation for faster iteration](https://www.reddit.com/r/aws/comments/i1gxhe/cloudformation_validation_for_faster_iteration/)
- url: https://www.reddit.com/r/aws/comments/i1gxhe/cloudformation_validation_for_faster_iteration/
---
I know that CloudFormation YAML/JSON can be generated, but suppose you were writing it by hand, and you’d want a fast turnaround time for validation. What would you use? 

I use `yamllint`, `cfn-lint` and `aws cloudformation validate-template ...`. But even with all of these, something that I’d think of as simple to validate, such as a malformed IP, isn’t caught, and that error is only caught when actually applying the stack to AWS. 

That means slow feedback as well as dollars lost for setting the gears in motion. 

This seems silly and avoidable. Is there some better validation that I'm not seeing?

Another example of something that isn't caught is a reference to a non-existing `X`. For example, I had referenced a `configSet` whose definition I had indented wrongly, so that it "didn't exist". Perhaps it's more complicated than I realize, but this also seems like something that could be found via a dry run. I'm sure there are other examples.
## [5][EC2 Auto Scale using a variable predefined schedule](https://www.reddit.com/r/aws/comments/i1rg0y/ec2_auto_scale_using_a_variable_predefined/)
- url: https://www.reddit.com/r/aws/comments/i1rg0y/ec2_auto_scale_using_a_variable_predefined/
---
I want to conduct online examinations on EC2 instances.   
I know how many students there will be at each time instant as the tests are pre-scheduled.  
So I want to scale using the schedule I have, it is variable not fixed or weekly, Can anyone help me where to look for this.

Thank you
## [6][Maven deploy Failed to deploy artifacts: Could not find artifact](https://www.reddit.com/r/aws/comments/i1qoae/maven_deploy_failed_to_deploy_artifacts_could_not/)
- url: https://www.reddit.com/r/aws/comments/i1qoae/maven_deploy_failed_to_deploy_artifacts_could_not/
---
 

I'm a newcomer to all this and i'm trying to deploy a maven app on elastic beanstalk while following this guide [https://aws.amazon.com/pt/blogs/developer/deploying-java-applications-on-elastic-beanstalk-from-maven/](https://aws.amazon.com/pt/blogs/developer/deploying-java-applications-on-elastic-beanstalk-from-maven/).

However when i run the mvn -Ps3-deploy package deploycommand I get an error saying Failed to deploy artifacts: Could not find artifact Timzudo:MinguitoOnline:jar:1 in 20278788 ([https://gitlab.com/Msgouveia/ok.git](https://gitlab.com/Msgouveia/ok.git))

I don't really understand why this is happening since the repo is public and the project that is there has the same groupId, artifactId, version, and packaging.

Info and distributionManagement:

&lt;modelVersion&gt;4.0.0&lt;/modelVersion&gt;

&lt;groupId&gt;Timzudo&lt;/groupId&gt;

&lt;artifactId&gt;MinguitoOnline&lt;/artifactId&gt;

&lt;version&gt;1&lt;/version&gt;

&lt;packaging&gt;jar&lt;/packaging&gt;

&lt;name&gt;Java Sample Jetty App&lt;/name&gt;

&lt;distributionManagement&gt;

&lt;repository&gt;

&lt;uniqueVersion&gt;false&lt;/uniqueVersion&gt;

&lt;id&gt;20278788&lt;/id&gt;

&lt;name&gt;MinguitoOnline&lt;/name&gt;

&lt;url&gt;[https://gitlab.com/Msgouveia/ok.git](https://gitlab.com/Msgouveia/ok.git)&lt;/url&gt;

&lt;layout&gt;default&lt;/layout&gt;

&lt;/repository&gt;

&lt;/distributionManagement&gt;
## [7][Help understanding launch configuration](https://www.reddit.com/r/aws/comments/i1py8u/help_understanding_launch_configuration/)
- url: https://www.reddit.com/r/aws/comments/i1py8u/help_understanding_launch_configuration/
---
From the documentation: 
“When you create a launch configuration, the default value for the instance placement tenancy is null and the instance tenancy is controlled by the tenancy attribute of the VPC. If you set the Launch Configuration Tenancy to default and the VPC Tenancy is set to dedicated, then the instances have dedicated tenancy. If you set the Launch Configuration Tenancy to dedicated and the VPC Tenancy is set to default, then again the instances have dedicated tenancy.”

In the last sentence, why would the instance tenancy be dedicated if the VPC tenancy is default. The first sentence says the instance tenancy is controlled by the attribute of the VPC. In that case, shouldn’t it be default aswell? 

Can anyone help me understand this please
## [8][DynamoDB data modeling question. Many to many relationship?](https://www.reddit.com/r/aws/comments/i1kgz5/dynamodb_data_modeling_question_many_to_many/)
- url: https://www.reddit.com/r/aws/comments/i1kgz5/dynamodb_data_modeling_question_many_to_many/
---
Hello! I'm trying to use DynamoDB to keep track of the count of certain items and tags. The main idea is, I want to be able to have a relationship between items and tags with their count; a many-to-many relationship (?) since an item can be submitted with multiple tags attached to it, and a tag can be counted towards many items.

Currently I have a small table kind of like this:

PK | SK | count | name
--|--|-----|----
PROFILE#username1 | TAG#tag1 | 3 | tag1
PROFILE#username1 | ITEM#item1 | 5 | item1
PROFILE#username1 | ITEM#item2 | 7 | item2

`PK` and `SK` are the primary and sort keys. With this configuration I can get the count of a certain item, and a certain tag per user. I can also get the list of the items and tags per user sorted by count thanks to a Local Secondary Index using the attribute `count` as the sort key.

What I want now is to know how much of the `count` attribute of an item come from certain tags. For example, if item1 has a count of 5, I want to be able to tell that 3 of them are from the tag1 and the remaining 2 are from tag2; all this for a particular user since users will be able to create tags and count items (with said tags) without virtually any restriction.

I was thinking just adding an extra record with an attribute of the tag from which it comes from since Dynamo has no problem with several records? 

I really can't wrap my mind around how would I do this, also, maybe in the future there would be like a history (per user) of every time (timestamp) an item with or without tag is added. It just keeps getting more complex, for me at least. I have read the AWS doc on Adjacency List Design Pattern but it is still not very clear to me.

I chose Dynamo because I believed it would be better for maintaining (serverless) and scaling, but it is my first time trying to model something real and it is proving to be counter productive time-wise; maybe I should stick with regular SQL?

If my explanation is confusing I apologize because it is also confusing for me (heh); maybe I should start by getting my ideas more in order first.
## [9][Has RDS PostgreSQL auto minor version increment ever worked?](https://www.reddit.com/r/aws/comments/i1desc/has_rds_postgresql_auto_minor_version_increment/)
- url: https://www.reddit.com/r/aws/comments/i1desc/has_rds_postgresql_auto_minor_version_increment/
---
I have always had the minor version increment on for all of my PostgreSQL databases.  At this point I have around 10 that are all stuck on 12.2 and a few that are on 11.7.  Neither one of those are the latest versions available.  I for a long time had a ton of databases stuck on an early version of 10 and they never updated until I manually updated them to 12.2.  At this point I have been on RDS for a couple of years and never seen a single one of our databases automatically update minor versions.  Is there some black magic besides checking the box I am supposed to do to get this to actually work?

I remember seeing that box and wondering why it never worked years ago.  Then a while back they announced that they "enhanced" it with  [https://aws.amazon.com/about-aws/whats-new/2018/12/amazon-rds-enhances-auto-minor-version-upgrades/](https://aws.amazon.com/about-aws/whats-new/2018/12/amazon-rds-enhances-auto-minor-version-upgrades/) , which I was hoping meant they actually made it work.  But nope still nothing.
## [10][Route 53 + Lightsail link not working, getting DNS_PROBE_FINISHED_NXDOMAIN error](https://www.reddit.com/r/aws/comments/i1kclo/route_53_lightsail_link_not_working_getting_dns/)
- url: https://www.reddit.com/r/aws/comments/i1kclo/route_53_lightsail_link_not_working_getting_dns/
---
I've just moved all my domains to AWS Route 53 to be registered and have their DNS hosted zones there from another provider. 

I'm trying to point my Route 53 hosted domains to my Lightsail instances.

I've followed all the steps given in these two links, and to no avail.

* [https://lightsail.aws.amazon.com/ls/docs/en\_us/articles/amazon-lightsail-using-route-53-to-point-a-domain-to-an-instance](https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-using-route-53-to-point-a-domain-to-an-instance)
* [https://cloudconfusing.com/2019/04/18/how-to-use-lightsail-with-route-53-dns/](https://cloudconfusing.com/2019/04/18/how-to-use-lightsail-with-route-53-dns/)

I have more than 3 domains, so the Lightsail DNS Zones are not an option for me, is there any fix to avoid getting this error 'DNS\_PROBE\_FINISHED\_NXDOMAIN' in Google Chrome
