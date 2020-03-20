# aws
## [1][Run Folding@home on AWS spot instances and help crack COVID-19](https://www.reddit.com/r/aws/comments/flkdgm/run_foldinghome_on_aws_spot_instances_and_help/)
- url: https://www.reddit.com/r/aws/comments/flkdgm/run_foldinghome_on_aws_spot_instances_and_help/
---
Spare AWS credits? [This CloudFormation template](https://github.com/jkataja/cfn-foldingathome) creates an AWS spot instance fleet for running the [Folding@Home](https://foldingathome.org/) client. Folding@home is a computing platform to assist disease research, for instance to find a cure for the COVID-19 virus. The template uses g4dn.xlarge instances with GPUs.

The template installs Ubuntu 18.04 LTS, NVidia CUDA 10.2 and Folding@home client 7.5.1. Folding@home client is started automatically after instance initialization is complete. Client runs until the template is removed, auto scaling group is scaled in or spot instance is reclaimed. The template bids 100% of on-demand price. Spot instance pricing varies: I have seen approximately 60-70% discount. The Folding@home servers have been recently busy and you cannot always retrieve a work package, so there may be some idle.

GitHub link: [https://github.com/jkataja/cfn-foldingathome](https://github.com/jkataja/cfn-foldingathome)
## [2][Amazon WorkSpaces and Amazon WorkDocs for up to 50 users no charge through June 30, 2020.](https://www.reddit.com/r/aws/comments/flcwwo/amazon_workspaces_and_amazon_workdocs_for_up_to/)
- url: https://www.reddit.com/r/aws/comments/flcwwo/amazon_workspaces_and_amazon_workdocs_for_up_to/
---
 [https://aws.amazon.com/blogs/desktop-and-application-streaming/new-offers-to-enable-work-from-home-from-amazon-workspaces-and-ama](https://aws.amazon.com/blogs/desktop-and-application-streaming/new-offers-to-enable-work-from-home-from-amazon-workspaces-and-amazon-workdocs/) 

&amp;#x200B;

*We are announcing two new offers that enable you to use Amazon WorkSpaces and Amazon WorkDocs for up to 50 users at no charge. Both offers are for new customers that have not previously used these services and are available through June 30, 2020. The WorkSpaces offer is available beginning on April 1st and will include our Standard, Value, and Performance bundles. The WorkDocs offer is available immediately. Together, WorkSpaces and WorkDocs provide remote workers the ability to securely use the applications they need, while collaborating on documents with their colleagues around the world. If you want to get started with WorkSpaces immediately, you can use our ongoing* [*Free Tier*](https://aws.amazon.com/free/?all-free-tier.sort-by=item.additionalFields.SortRank&amp;all-free-tier.sort-order=asc&amp;all-free-tier.q=Desktop%2B%26%2BApp%2BStreaming%2BAmazon%2BWorkSpaces%2B40%2Bhours&amp;all-free-tier.q_operator=AND) *offer of the WorkSpaces Standard bundle* [zon-workdocs/](https://aws.amazon.com/blogs/desktop-and-application-streaming/new-offers-to-enable-work-from-home-from-amazon-workspaces-and-amazon-workdocs/)  

&amp;#x200B;

I'm an AWS Employee.
## [3][Exam reschedule due to COVID-19](https://www.reddit.com/r/aws/comments/flsugg/exam_reschedule_due_to_covid19/)
- url: https://www.reddit.com/r/aws/comments/flsugg/exam_reschedule_due_to_covid19/
---
Has anyone facing issues with connecting to PSI Support ?

We have a total lockdown on 22nd March due to Coronavirus, and my exam is scheduled on the same day. Due to the outbreak, I have already exhausted my 2x rescheduling options and now I am unable to reschedule it for a later date.

Do we have someone here from PSI ? or can AWS help in this regard ?
## [4][AWS signature for SNS?](https://www.reddit.com/r/aws/comments/fltq0c/aws_signature_for_sns/)
- url: https://www.reddit.com/r/aws/comments/fltq0c/aws_signature_for_sns/
---
I'm using the AWS key derivation for python listed in the [example](https://docs.aws.amazon.com/general/latest/gr/sigv4-signed-request-examples.html#sig-v4-examples-get-auth-header) for use without boto3. 

Im submitting a request to publish to an SNS Topic but keep getting the error message `The request signature we calculated does not match the signature you provided. Check your AWS Secret Access Key and signing method. Consult the service documentation for details.`. The Access Key and Secret Key match whats in IAM.

The only thing I changed in the script was the URI params:

    method = 'GET'
    service = 'sns'
    host = 'sns.us-east-1.amazonaws.com'
    region = 'us-east-1'
    endpoint = 'https://sns.us-east-1.amazonaws.com'
    request_parameters = 'Action=Publish&amp;Version=2010-03-31&amp;Message=HowBowDah&amp;TopicArn=arn%3aws%3sns%3us-east-1%3151341732511%3test'



Does this underlying request look right?

    GET /?Action=Publish&amp;Version=2010-03-31&amp;Message=HowBowDah&amp;TopicArn=arn%253Aws%253sns%253us-east-1%253151341732511%253test HTTP/1.1
    Host: sns.us-east-1.amazonaws.com
    User-Agent: python-requests/2.22.0
    Accept-Encoding: gzip, deflate
    Accept: */*
    Connection: close
    x-amz-date: 20200320T113335Z
    Authorization: AWS4-HMAC-SHA256 
    Credential=AKIASGPFLSKPVMSJFLML/20200320/us-east-1/sns/aws4_request, SignedHeaders=host;x-amz-date, Signature=8d5a555626c592146768499aee1e85d64f6d4898cfb568ea2544cf7ad0d6d9f0
## [5][Signed Cookies AWSSDK - Cloudfront](https://www.reddit.com/r/aws/comments/floatz/signed_cookies_awssdk_cloudfront/)
- url: https://www.reddit.com/r/aws/comments/floatz/signed_cookies_awssdk_cloudfront/
---
Anyone have a working solution with the new Cloudfront AWSSDK to get the signed cookies working with alternate domain.  Created the CNAME for the hosted domain, e.g., (cdn.example.com).  Signed URLs works fine but for the life of me cannot get the signed cookies to work.
## [6][aws-cli S3 bucket setting Content-Type text/html ?](https://www.reddit.com/r/aws/comments/flt4g0/awscli_s3_bucket_setting_contenttype_texthtml/)
- url: https://www.reddit.com/r/aws/comments/flt4g0/awscli_s3_bucket_setting_contenttype_texthtml/
---
Hello,

I'm trying to upload a bunch of files with no extensions that are in fact html

I'm doing it like this:

aws s3 sync . s3://$FRONTEND\_BUCKET\_NAME --exclude '\*.\*' --content-type text/html --content-language html --metadata-directive REPLACE

Yet content-type is set to  binary/octet-stream in the bucket? What am I missing?

&amp;#x200B;

UPDATE:

Ok I figured it out, trash this post if you want but solution is to:

dist: bionic  
install: pyenv global 3.8.1 &amp;&amp; pip3 install --user awscli==1.17.7

because reasons I guess.
## [7][Non-USD billing and COVID-19](https://www.reddit.com/r/aws/comments/flqem9/nonusd_billing_and_covid19/)
- url: https://www.reddit.com/r/aws/comments/flqem9/nonusd_billing_and_covid19/
---
Has anyone seen any communications about what will happen for those of us billed in a non-USD currency? Exchange rates have taken a plunge with the global pandemic so I'm wondering what the options are... 😨
## [8][Programatically modify a security group?](https://www.reddit.com/r/aws/comments/flst8y/programatically_modify_a_security_group/)
- url: https://www.reddit.com/r/aws/comments/flst8y/programatically_modify_a_security_group/
---
It doesn't look like this is specifically possible. To be specific, I want to modify a member IP permision within a group. I'm using the .NET API, but since it's feature parity seemingly with the command line, I've been looking at that too.

&amp;#x200B;

I thought I was doing well, I was able to find my way through describeSecurityGroups, and find the appropriate SecurityGroup. Then i can iterate the IpPermissions, into IPv4Ranges and loop that and find the one I want to change (by Description in this case)

&amp;#x200B;

Ideally, I'd be wanting to update the IP when I find it, but there is only UpdateDescription for it, not the IP. I can't seem to find any commands to remove a specific Ipv4Range entry. 

&amp;#x200B;

To give some back story, we have a web-based product that is used by many users, including our staff. Our staff also have RDP access to a Bastion host that they need to use on occasion. Our office IPs are whitelisted, and some of our staff have static Ips, so those aren't an issue either. However a lot of staff are on dynamic connections. We're getting a bit overrun of "can you add my IP to the bastion" every day with so many now working from home. We don't have a VPN into the production network, it's what this Bastion host is for. 

What I had hoped to do was to use the Description field, and put the username from the web product in there, and then during logon to the product, pickup their IP, and update the security group accordingly, so as soon as they web-login, they would also be granted access on the Bastion. 

Since there doesn't seem to be function to do this, I guess the only option is to make a security group per user, but is there a limit of how many security groups can exist on an EC2 instance, and this seems a lot messier. Short of a VPN into production, can anyone think of any other way to approach this?
## [9][Determining read/write capacity on DynamoDB](https://www.reddit.com/r/aws/comments/flscvx/determining_readwrite_capacity_on_dynamodb/)
- url: https://www.reddit.com/r/aws/comments/flscvx/determining_readwrite_capacity_on_dynamodb/
---
Can someone please help me determine the right read/write capacity units for my dynamodb tabels? I have 2 tables already in dynamodb, but one of the engineers set up read and write cus too high when created the tables, so the costs are higher that they should be. How can I determine the right amount of units to minimize the costs and maintain the performance? Thanks
## [10][a standard IT infrastructure for ROS robots and IoT devices.](https://www.reddit.com/r/aws/comments/flsa2h/a_standard_it_infrastructure_for_ros_robots_and/)
- url: https://github.com/rdbox-intec/rdbox
---

