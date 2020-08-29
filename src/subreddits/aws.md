# aws
## [1][The new route 53 UI is terrible](https://www.reddit.com/r/aws/comments/ii8ts4/the_new_route_53_ui_is_terrible/)
- url: https://www.reddit.com/r/aws/comments/ii8ts4/the_new_route_53_ui_is_terrible/
---
Didn't I already post this? Oh wait no, I'm sorry. That was the new calculator UI.

AWS...please stop with all the wizard nonsense. Again. I don't need a wizard to hold my hand through creating a TXT record. I need something simple, or as you now call it, the "old console". I get the desire to create an experience, but please do it where it is warranted. Who in the community is asking for you to complicate the process of creating DNS records? I would rather you take us back to the days of editing BIND files with VIM than have to work in your new console. And I am not alone! A colleague of mine today just shared his feelings to me about your new console. He said, " real DNS ballers edit BIND files with vim". If you need a wizard to create DNS records, you should not be creating DNS records.
## [2][Over 54,000 scanned NSW (Australia) driver's licences found in open cloud storage - Security - iTnews](https://www.reddit.com/r/aws/comments/iikq97/over_54000_scanned_nsw_australia_drivers_licences/)
- url: https://www.itnews.com.au/news/over-54000-scanned-nsw-drivers-licenses-found-in-open-cloud-storage-552544
---

## [3][Automatically deploying your Container Apps with AWS Copilot](https://www.reddit.com/r/aws/comments/iihe4b/automatically_deploying_your_container_apps_with/)
- url: https://aws.amazon.com/blogs/containers/automatically-deploying-your-container-application-with-aws-copilot/
---

## [4][Labs to learn about AWS Analytics Services](https://www.reddit.com/r/aws/comments/iink54/labs_to_learn_about_aws_analytics_services/)
- url: https://aws-dojo.com/digests/analytics.html
---

## [5][How to use cloudfront with s3 origin from different account](https://www.reddit.com/r/aws/comments/iil99g/how_to_use_cloudfront_with_s3_origin_from/)
- url: https://www.reddit.com/r/aws/comments/iil99g/how_to_use_cloudfront_with_s3_origin_from/
---
Hello everyone. I have an account A1 with S3 where my content resides and I want to use cloudfront from account A2 to serve it. If I would want to achieve something like this, what needs to be done. 

I could see cloudfront asking for origin url but has no mention of access key and secret. Does that mean if I have to access S3 content of A1 from cloudfront of A2, I need to make A1 account S3 public ? Is there any other way to solve this problem ?
## [6][CloudFormation, Terraform, or CDK? A guide to infrastructure as code on AWS](https://www.reddit.com/r/aws/comments/ii73jl/cloudformation_terraform_or_cdk_a_guide_to/)
- url: https://acloudguru.com/blog/engineering/cloudformation-terraform-or-cdk-guide-to-iac-on-aws
---

## [7][Free tier a scam or just buggy?](https://www.reddit.com/r/aws/comments/iipz3h/free_tier_a_scam_or_just_buggy/)
- url: https://www.reddit.com/r/aws/comments/iipz3h/free_tier_a_scam_or_just_buggy/
---
So I've signed up to the free tier. Uploaded the sample app with a minor tweak. Refreshed the website maybe 5 or 6 times. Then I get this email... 

Dear AWS Customer,
Your AWS account ******** has exceeded 85% of the usage limit for one or more AWS Free Tier-eligible services for the month of August.
AWS Free Tier Usage as of 08/28/2020AWS Free Tier Usage Limit2000.0 Requests; 2,000 Put, Copy, Post or List Requests of Amazon S3

I can't imagine how I could have reached 2,000 requests in under one day of slow, lazy testing... So, what's going on?
## [8][OCI Artifact Support In Amazon ECR](https://www.reddit.com/r/aws/comments/iicumq/oci_artifact_support_in_amazon_ecr/)
- url: https://aws.amazon.com/blogs/containers/oci-artifact-support-in-amazon-ecr/
---

## [9][I cant write to my Cognito Custom attributes](https://www.reddit.com/r/aws/comments/iik1wa/i_cant_write_to_my_cognito_custom_attributes/)
- url: https://www.reddit.com/r/aws/comments/iik1wa/i_cant_write_to_my_cognito_custom_attributes/
---
So, I have a custom attribute that I am able to write to, but then I realized I would like to have one more variable stored alongside the user login, so I added another custom attribute, however, when I write to it from my website, I get the error: "A client attempted to write unauthorized attribute". I am not sure why it is doing this. I have verified that the name is the same as what I am saving it as. Any advice?
## [10][How can I allow users to map their own domain to a subdomain on my site?](https://www.reddit.com/r/aws/comments/iijouo/how_can_i_allow_users_to_map_their_own_domain_to/)
- url: https://www.reddit.com/r/aws/comments/iijouo/how_can_i_allow_users_to_map_their_own_domain_to/
---
So I run a service where a user can create a site and they get their own subdomain. So if a user has site `a-site-name` their subdomain will be `a-site-name.domain.com`. Some of the users want to be able to have their own domains instead of the subdomain on my domain. I know I should be able to do this with a CNAME record in DNS (so `sub.the-users-domain.com` resolves to `a-site-name.domain.com`). However when I try to do this, I get a 403 from Cloudfront.

To describe my setup:

- Files are served by AWS Cloudfront using an S3 bucket as the origin point.
- Each site `a-site-name` is stored in a folder in the S3 bucket of the same name.
- I have a Lambda@Edge function running so that whenever a request is made to `a-site-name.domain.com/whatever-resource` (wildcard subdomains) it actually resolves to and grabs data from `distribution-id.cloudfront.net/a-site-name/whatever-resource`.

I'm guessing the fact that I'm getting a 403 has to do with the configuration of my Cloudfront deployment or with that Lambda@Edge function. Does anyone have any ideas as to how to change the configuration so that people can just use a CNAME record to use their own domains? I'm lost to be honest.
