# aws
## [1][LucidChart; 'AWS Import' tool being removed from all Pro accounts](https://www.reddit.com/r/aws/comments/f9b57g/lucidchart_aws_import_tool_being_removed_from_all/)
- url: https://www.reddit.com/r/aws/comments/f9b57g/lucidchart_aws_import_tool_being_removed_from_all/
---
  

Hi All, 

I was just on the phone with the LucidChart sales team to enquire about their ‘cloud insights’ product. They informed me that the “AWS Import’ Tool was going to sunset, and ‘cloud insights’ was going to be the replacement. The new tool will not be part of the ‘pro’ licensing model. It will be an add-on and will also have a per resource import cost attached to it. Needless to say, this tool is fantastic and probably the best solution out there for all my diagramming needs. I have no quarrels with paying what it's worth, but I have negative ideas about companies removing features only to add the same feature at a higher cost with usage costs on top.
## [2][AWS Free Tier usage limit alerting via AWS Budgets](https://www.reddit.com/r/aws/comments/f9q2lk/aws_free_tier_usage_limit_alerting_via_aws_budgets/)
- url: https://www.reddit.com/r/aws/comments/f9q2lk/aws_free_tier_usage_limit_alerting_via_aws_budgets/
---
Hoping someone can help me with this. I received an email this morning stating   "AWS Free Tier usage limit alerting" my account has exceeded 85% of the usage limit for one or more AWS Free Tier-eligible services for the month of February.

I've 100% not got any instances of anything running. I've not even used the account in months and I've checked everywhere I can in my portal to be sure and can't see anything.

I've used the portal probably once so not 100% familiar with it and hoping someone could help to make sure I've not missed looking somewhere that would show me what's happening.
## [3][Has anyone built/maintained a medium-size web application fully on AWS Lambda? My experience has been horrific.](https://www.reddit.com/r/aws/comments/f9fowb/has_anyone_builtmaintained_a_mediumsize_web/)
- url: https://www.reddit.com/r/aws/comments/f9fowb/has_anyone_builtmaintained_a_mediumsize_web/
---
I joined a team a few months ago and took over a web application that prides itself fully serverless and uses microservice architecture. The entire application is written in NodeJs and built on AWS Lambda using Serverless framework.

While I can understand that using Lambdas can help us at completing new features pretty quickly, which the product manager and senior management love, I have wondered if it's a good idea in the long run to have everything on Lambda.

The project has only been slightly more than 2 years old and the entire "microservice" architecture and codebase have become so difficult to maintain. While it's really convenient to start a new feature entirely separate from the existing codebase, it's extremely difficult and messy whenever I need to change anything in the existing codebase. 

Using NodeJs doesn't make things any better because now everyone is littering functions everywhere in files and folders. There are so many functions spread across different files that I need to figure out which one to import. The non-typed codes are so messy that I have to keep printing to the console to check what's in a variable.

On top of all the usual issues in a microservice architecture, the hell lot of Lambda functions makes the already complicated architecture extremely difficult to trace and debug!

It's so frustrating to be working on this project!

I'm not sure what is the "best" practice when it comes to building applications on serverless. Is building a medium-size web application entirely on AWS Lambda a common thing in the industry and is it a good idea to do it this way? Or has someone just started this entire project the wrong way, left the company and everyone else who took over it just followed along the mess?

If anyone has any experience building an a medium-size web application entirely on Lambda, I would  love to hear about your experience with it. I would also love to know what is the correct and better way of building an entire application serverless platforms.
## [4][Lambda to do Cloudwatch logs to Elastic domain is timing out on POST](https://www.reddit.com/r/aws/comments/f9qfhd/lambda_to_do_cloudwatch_logs_to_elastic_domain_is/)
- url: https://www.reddit.com/r/aws/comments/f9qfhd/lambda_to_do_cloudwatch_logs_to_elastic_domain_is/
---
Trying to follow [https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL\_ES\_Stream.html](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_ES_Stream.html) , it seems pretty straight forward and the lambda is called correctly. By adding some console.log calls I can see the lambda runs up to posting to the Elastic endpoint and then times out.

I have followed the article twice, once creating the ES domain through terraform, and once through the web console. Unfortunately times out still happens in both cases.

The lambda has the AWSLambdaVPCAccessExecutionRole attached, as well as additional rights to get triggered by cloudwatch (and this works).

Both the ES domain and the lambda are sitting in the same VPC and the same (private) subnet. The VPC has DNS resolution set to true.

The security group attached to the ES domain has https/443 allowed for the VPC CIDR on ingress and egress is open wide ([0.0.0.0/0](https://0.0.0.0/0)).

I can tunnel into the subnet through a bastion and I am able to see the Elastic instance and open kibana but no data is coming in.

What am I missing here ?

What can I check to figure out what's missing ?

update: protocol
## [5][Private VPC needs access to mail relay](https://www.reddit.com/r/aws/comments/f9nn1i/private_vpc_needs_access_to_mail_relay/)
- url: https://www.reddit.com/r/aws/comments/f9nn1i/private_vpc_needs_access_to_mail_relay/
---
We have a private VPC with a few EC2 instances. I’m trying to telnet from the instances to mail relay on port 25 but it does not connect. The mail relay is a public IP address so I’m not sure it is going to work or how it would work. I have read some stuff about port 25 not being allowed unless you specifically submit a form to aws asking for it to be available. Should I be able to Telnet? I can ping the mail relay no problem and all outbound access is wide open. Is there anything inbound that I need to add? From the instances, I can launch a web browser and browse public internet (google, espn, Facebook, etc.) so there is a path to the internet that way. I am new to AWS so there could be something I’m missing too.
## [6][A Better Way to SSH in AWS](https://www.reddit.com/r/aws/comments/f9abzy/a_better_way_to_ssh_in_aws/)
- url: https://nullsweep.com/a-better-way-to-ssh-in-aws/
---

## [7][Cross account S3 buckets and AWS SSO?](https://www.reddit.com/r/aws/comments/f9ibln/cross_account_s3_buckets_and_aws_sso/)
- url: https://www.reddit.com/r/aws/comments/f9ibln/cross_account_s3_buckets_and_aws_sso/
---
I'm using AWS SSO to give access to my developers from their GSuite accounts, which is working great. The only problem is that I don't know what principal to give access to in my cross-account IAM roles. If I want a user to log in to one account and access an S3 bucket in the other, none of the normal solutions seem to work.

Any ideas?
## [8][Amazon FSx for Lustre Update: Persistent Storage for Long-Term, High-Performance Workloads](https://www.reddit.com/r/aws/comments/f9d6cb/amazon_fsx_for_lustre_update_persistent_storage/)
- url: https://aws.amazon.com/blogs/aws/amazon-fsx-for-lustre-persistent-storage/
---

## [9][Could not turn on the AWS S3 Hosted website](https://www.reddit.com/r/aws/comments/f9bx7y/could_not_turn_on_the_aws_s3_hosted_website/)
- url: https://www.reddit.com/r/aws/comments/f9bx7y/could_not_turn_on_the_aws_s3_hosted_website/
---
&amp;#x200B;

[A NAME, NS and SOA](https://preview.redd.it/qbvb0g9z83j41.png?width=2512&amp;format=png&amp;auto=webp&amp;s=fcbb8a4897d3355acab00c5113186c660dc99663)

&amp;#x200B;

[Settings](https://preview.redd.it/si0ml6c293j41.png?width=3094&amp;format=png&amp;auto=webp&amp;s=24748488dfdd30203cc6ae02d47b51cc055b863a)

&amp;#x200B;

[Bucket Policy](https://preview.redd.it/k68rjl1b93j41.png?width=2146&amp;format=png&amp;auto=webp&amp;s=90922eed5e6313b802dcefaf8687ea517e1f98ff)

The source is here

[http://www.devops.engineering.s3-website-us-east-1.amazonaws.com](http://www.devops.engineering.s3-website-us-east-1.amazonaws.com)

I got the domain from Amazon

&amp;#x200B;

[Domain record](https://preview.redd.it/harc1pszz8j41.png?width=3360&amp;format=png&amp;auto=webp&amp;s=2fcd82b5ecc4fb82a2a253578fff5b19a8498252)

&amp;#x200B;

What is missing in this picture? ( other than my lack of knowledge :) )

&amp;#x200B;

[http://www.devops.engineering](http://www.devops.engineering) return error ERR\_NAME\_NOT\_RESOLVED
## [10][Suspicious CentOS 7 AMIs on AWS Marketplace](https://www.reddit.com/r/aws/comments/f9aruy/suspicious_centos_7_amis_on_aws_marketplace/)
- url: https://www.reddit.com/r/aws/comments/f9aruy/suspicious_centos_7_amis_on_aws_marketplace/
---
I use Packer to build AMIs based on CentOS 7. My filter configuration looks like this:

        "source_ami_filter": {
          "filters": {
            "name": "CentOS Linux 7 x86_64 HVM EBS *"
          },
          "owners": ["679593333241"],
          "most_recent": true
        },

This has worked fine for a number of years, allowing me to use the latest available AMIs when CentOS release updates. This week, builds have started failing:

    ==&gt; amazon-ebs: Error launching source instance: OptInRequired: In order to use this AWS Marketplace product you need to accept terms and subscribe. To do so please visit https://aws.amazon.com/marketplace/pp?sku=67xglex2rdpaymxh17620nfoy

It looks like "[SupportedImages.com](https://SupportedImages.com)" have created a new AMI with a very similar name to the existing CentOS-provided AMIs. I almost went ahead and accepted the Marketplace terms and conditions before realising these AMIs are not free - there is a subscription charge associated with them of $399 up to $23,999 depending on instance size!

The official CentOS AMI with AMI name "CentOS Linux 7 x86\_64 HVM EBS ENA 1901\_01-b7ee8a69-ee97-4a49-9e68-afaee216db2e-ami-05713873c6794f575.4" is on the Marketplace here:

[https://aws.amazon.com/marketplace/pp/Centosorg-CentOS-7-x8664-with-Updates-HVM/B00O7WM7QW](https://aws.amazon.com/marketplace/pp/Centosorg-CentOS-7-x8664-with-Updates-HVM/B00O7WM7QW)

The SupportedImages AMI with AMI name "CentOS Linux 7 x86\_64 HVM EBS ENA 1901\_01-691155db-037a-4082-81f2-6936b7091732-ami-07aeae1d312ef8dce.4" is here:

[https://aws.amazon.com/marketplace/pp/B07X8R39Q3/](https://aws.amazon.com/marketplace/pp/B07X8R39Q3/)

The SupportedImages website is very basic and the whole thing seems quite shady; to me this looks like a deliberate attempt to mislead with the AMI name and repackaging of CentOS but with a charge attached. However, I am also not really sure what can be done here - is this just a side effect of allowing public images with no restrictions on names on the Marketplace? Who would I even report this to?

Also, are these images even trustworthy? How would I be able to tell?

For anyone who has the same issue I have worked around this issue for now by adding the CentOS product code to my filter:

        "source_ami_filter": {
          "filters": {
            "name": "CentOS Linux 7 x86_64 HVM EBS *",
    	"product-code": "aw0evgkw8e5c1q413zgy5pjce"
          },
          "owners": ["679593333241"],
          "most_recent": true
        },
