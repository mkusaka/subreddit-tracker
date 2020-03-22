# aws
## [1][Unzipping s3 file in chunks using lambda](https://www.reddit.com/r/aws/comments/fmrwfb/unzipping_s3_file_in_chunks_using_lambda/)
- url: https://www.reddit.com/r/aws/comments/fmrwfb/unzipping_s3_file_in_chunks_using_lambda/
---
I used lambda to unzip S3 file back to s3.  Lambda can't unzip more than 3gb file.  I created a script that unzips by streaming chunks of file. 

[https://medium.com/@multiaki/unzipping-s3-files-back-to-s3-without-uncompressing-entire-file-streaming-82f662b5065a](https://medium.com/@multiaki/unzipping-s3-files-back-to-s3-without-uncompressing-entire-file-streaming-82f662b5065a)
## [2][Existing VPS image to EC2](https://www.reddit.com/r/aws/comments/fmxbzm/existing_vps_image_to_ec2/)
- url: https://www.reddit.com/r/aws/comments/fmxbzm/existing_vps_image_to_ec2/
---
For lift &amp; shifts (Rehosting), I think CloudEndure is used for Windows workloads, but what do people use for Linux based workloads?

I gather the best option is to boot the OS AMI on EC2 and rsync across the system as best you can and create your own AMI? Or am I missing a tool? VM import/export is only for VMware AFACIT.
## [3][Cloudfront really can't handle sub-directory indexes without a lambda trigger?](https://www.reddit.com/r/aws/comments/fmzcjg/cloudfront_really_cant_handle_subdirectory/)
- url: https://www.reddit.com/r/aws/comments/fmzcjg/cloudfront_really_cant_handle_subdirectory/
---
I was in disbelief to discover that using Cloudfront infront of a static S3 website you're unable to request sub-directory index files unless [you have a lambda trigger function that rewrites the request](https://aws.amazon.com/blogs/compute/implementing-default-directory-indexes-in-amazon-s3-backed-amazon-cloudfront-origins-using-lambdaedge/) for you.

Is this still the case or is there a better way to configure this in Cloudfront today?
## [4][How to configure cross-account Cognito authorizer for Appsync?](https://www.reddit.com/r/aws/comments/fmvuxf/how_to_configure_crossaccount_cognito_authorizer/)
- url: https://www.reddit.com/r/aws/comments/fmvuxf/how_to_configure_crossaccount_cognito_authorizer/
---
Is it possible to set up Appsync in a way that uses the Cognito user pool of another AWS account? API Gateway can use user pools from other accounts as an authorizer but I can't find docs for Appsync. I'd like to allow existing users to call appsync endpoints in a new account.
## [5][Display data in d3js visualizations(KPIs) using API Gateway, Lambda and Aurora](https://www.reddit.com/r/aws/comments/fmucc4/display_data_in_d3js_visualizationskpis_using_api/)
- url: https://www.reddit.com/r/aws/comments/fmucc4/display_data_in_d3js_visualizationskpis_using_api/
---
Hello,

We currently have data in S3, that gets transformed using EMR and copied to Redshift - which serves our internal users for building their visualizations using Tableau. We are looking for options to leverage the transformed data in S3, then build a few KPIs using D3js or something similar. Idea is to display the visualizations(usage stats) on the home page for all users when the user logs into the website. 

I am planning to use D3js, API Gateway, Lambda and Aurora Postgres to achieve this. The reasons I am leaning towards Aurora Postgres is 

* Transformed data that we build visualizations using D3js is less than 3GB.
* Redshift is expensive.
* our website can have 300-500 concurrent users at a time.

Here are the questions I've:

1. When I've so many concurrent users logged in, is triggering 300-500 Lambdas scalable?
2. Is there a better DB Backend solution than Aurora Postgres for less data and more concurrency?
3. Finally, is my approach correct?

Thanks in advance.
## [6][Monitor vpn traffic](https://www.reddit.com/r/aws/comments/fmuq3b/monitor_vpn_traffic/)
- url: https://www.reddit.com/r/aws/comments/fmuq3b/monitor_vpn_traffic/
---
Hi guys, i we have a bastion host in AWS infra, that we use it to tshoot vpn connectivity issues by simply pinging vpn peers.  
My purpose is to write the results of the ping in an html file and i was thinking to install apache web server and place there the html file so my colleagues that don't have ssh access to the bastion host, can see if pings fail through this html page.  
My concern is: isn't it dangerous if i expose our bastion to the public web by installing apache on it? can someone experienced on the matter offer me his advice of how to design a solution without security "holes"? It doesn't have to be necessarily via bastion host &amp; html, something with Cloudwatch perhaps ?  
Thank you!
## [7][AWS Active Directory Connector and Azure Active Directory Domain Services](https://www.reddit.com/r/aws/comments/fmogdc/aws_active_directory_connector_and_azure_active/)
- url: https://www.reddit.com/r/aws/comments/fmogdc/aws_active_directory_connector_and_azure_active/
---
I am trying to replicate the steps in this [post](https://aws.amazon.com/blogs/desktop-and-application-streaming/add-your-workspaces-to-azure-ad-using-azure-active-directory-domain-services/). However, When I get to the step where I create the ad connector, it fails with the following error:

\&gt; DNS unavailable (TCP port 53) for IP: [10.0.0.4](https://10.0.0.4), DNS unavailable (TCP port 53) for IP: [10.0.0.5](https://10.0.0.5)

I am very proficient with AWS. However, I'm struggling with Azure and feel I may have misconfigured something. I have carried out the following steps thus far:

&amp;#x200B;

In Azure, I used an existing resource group and created "Azure AD Domain Services" instance using default configuration

**Basics**

 \- Name: sy\*\*\*\*\*\*[k.com](https://k.com) 

 \- Subscription: Pay-As-You-Go 

 \- Resource Group:

 \- Default Region: UK South

 \- SKU: Standard  Forest type: User

**Network**

 \- Virtual network: (new) aadds-vnet

 \- Subnet: (new) aadds-subnet

 \- Subnet Address: [10.0.0.0/24](https://10.0.0.0/24)

 \- Network Security Group: (new) aadds-nsg

&amp;#x200B;

I created a site to site vpn connection with azure virtual network. However, I am not sure about this step in the post: "The tunnels must be configured to allow traffic from your AADDS endpoints and the Subnets" How exactly do I do this?

&amp;#x200B;

In AWS VPC cidr is [10.1.0.0/16](https://10.1.0.0/16) and both tunnels between AWS VPC and Azure Virtual Network are up and connected. I tried to contact the post author: "Justin Stokes" directly but can't find any emails for him. I cannot find a single online guide on how to set this up step by step along with the site to site ipsec setup. It would be very very helpful if someone can provide a video tutorial for this step by step from A-Z instead of leaving a chunk of the steps out of the guide.

&amp;#x200B;

The troubleshooting guide [here](https://docs.aws.amazon.com/directoryservice/latest/admin-guide/ad_connector_troubleshooting.html) suggest that the firewall i.e. network security group is not allowing port 53TCP/UDP inbound for AD Connector. But I updated the networks security group as a test with a rule to allow any source, any destination and any port and still I'm getting the same error.
## [8][One Lambda to Rule Them All - A Python Adventure in AWS (Feat. API Gateway)](https://www.reddit.com/r/aws/comments/fmffvo/one_lambda_to_rule_them_all_a_python_adventure_in/)
- url: https://medium.com/@cottenio/one-lambda-to-rule-them-all-44401893123f
---

## [9][Question: How to prevent direct linking from website.](https://www.reddit.com/r/aws/comments/fmevna/question_how_to_prevent_direct_linking_from/)
- url: https://www.reddit.com/r/aws/comments/fmevna/question_how_to_prevent_direct_linking_from/
---
Hello,
I have an S3 bucket setup to serve a static site along with CloudFront. I have the below bucket policy and CORS configuration, and what I am wondering now is how do I prevent direct access to assets except for when they are loaded through the site? That is, if I go to mySite.tld I want to see mySite.tld/images/someImage.jpg displayed on it, but if I go directly to mySite.tld/images/someImage.jpg in my browser I'd like it to give back forbidden or such. Is that feasible?

Motivation: I bought a commercial font and the license asks that I put such technical barriers in place.

Thank you in advance for any advice / help you can give.

Bucket Policy:
```
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "2",
      "Effect": "Allow",
      "Principal": {
        "AWS": "arn:aws:iam::cloudfront:user/CloudFront Origin Access Identity ------------"
      },
      "Action": "s3:GetObject",
      "Resource": "arn:aws:s3:::mySite.tld/*"
    }
  ]
}
```

CORS Configuration:
```
&lt;?xml version="1.0" encoding="UTF-8"?&gt;
&lt;CORSConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/"&gt;
  &lt;CORSRule&gt;
    &lt;AllowedOrigin&gt;https://mySite.tld&lt;/AllowedOrigin&gt;
    &lt;AllowedMethod&gt;GET&lt;/AllowedMethod&gt;
  &lt;/CORSRule&gt;
&lt;/CORSConfiguration&gt;
```
## [10][S3 Bucket Download Limit](https://www.reddit.com/r/aws/comments/fmiwaf/s3_bucket_download_limit/)
- url: https://www.reddit.com/r/aws/comments/fmiwaf/s3_bucket_download_limit/
---
Hi, I'm hoping you can help. I have an S3 bucket that I have created a TotalDownloaded metric, and now I'm trying to create an alarm that will go off if someone uses more than a monthly allotted amount of bandwidth but I don't see how to control the flow of data.

I don't know how to tame that metric and make it go into an alarm state when the limit is hit within 30 days.
