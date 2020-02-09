# aws
## [1][RDS and KMS Access - A follow up](https://www.reddit.com/r/aws/comments/f17a25/rds_and_kms_access_a_follow_up/)
- url: https://www.reddit.com/r/aws/comments/f17a25/rds_and_kms_access_a_follow_up/
---
I was recently struggling to understand why RDS had access to a CMK I created in KMS and came across [an old post on this sub](https://www.reddit.com/r/aws/comments/b8nrar/rds_and_kms_im_doing_something_wrong/) of someone asking the same question. 

The post is archived but I believe the comments are not correct and it sent me down the wrong path for a while.

The comment saying
&gt; Your $ACCOUNTID:root principal allows all principals in $ACCOUNTID to have those kms:* permissions...

is not true. From the [documentation](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam), it says:
&gt;  It does not by itself give any IAM users or roles access to the CMK, but it enables you to use IAM policies to do so

For my Aurora Serverless RDS, it was gaining access to the key from a grant. If you run `aws kms list-grants --key-id yourkey` you should see the grant which contains an encryption context with your DB instance ID. You should also see a CreateGrant entry in CloudTrail. The grant is created by the entity who creates the database.

I'm making this post in the hopes it saves someone from hours of head scratching in the future!
## [2][Should I encrypt, hash or encode?](https://www.reddit.com/r/aws/comments/f0tewr/should_i_encrypt_hash_or_encode/)
- url: https://blog.deleu.dev/should-i-encrypt-hash-or-encode/
---

## [3][Issues with AWS Educate and using Elastic Beanstalk](https://www.reddit.com/r/aws/comments/f13vun/issues_with_aws_educate_and_using_elastic/)
- url: https://www.reddit.com/r/aws/comments/f13vun/issues_with_aws_educate_and_using_elastic/
---
I've been following this [guide](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/nodejs-dynamodb-tutorial.html) for creating a node.js web application. I got as far as creating an EC2 instance running Amazon Linux, and I successfully connected to it using an SSH client called PuTTY. I did that just to run the commands in the tutorial. When I actually got the part about launching an Elastic Beanstalk environment, I followed this [link to Elastic Beanstalk](https://console.aws.amazon.com/elasticbeanstalk/home#/newApplication?applicationName=tutorials&amp;environmentType=LoadBalanced) provided by the guide. When I clicked "Get started" to launch a sample application, I flickered to a loading screen, and then back to the Elastic Beanstalk home. It refuses to allow me to launch the sample application. I did a search on this subreddit to see if anyone else had this issue, and apparently AWS Educate users aren't allowed to use this feature. I was wondering how I could get permissions to launch this, or if I should use another approach entirely. For context, I am trying to create a website which includes some features like messaging and a whiteboard to draw on. I am new to aws, servers, and environments, and the like, so I'm sorry if I left out any important details. Thank you for reading.
## [4][What retailers block access from AWS IP blocks?](https://www.reddit.com/r/aws/comments/f0xxxb/what_retailers_block_access_from_aws_ip_blocks/)
- url: https://www.reddit.com/r/aws/comments/f0xxxb/what_retailers_block_access_from_aws_ip_blocks/
---
I just recently moved a shopping application for a customer from Rackspace to AWS and we noticed that all connections to [lowes.com](https://lowes.com) from AWS are denied/blocked. I find this interesting/frustrating. Why would [lowes.com](https://lowes.com) as a whole block all access from AWS? It seems the blacklist connections from any/all AWS regions. Do other retailers do this as well?
## [5][Part 8 - Deploying a Vue.js and Serverless Imgur Clone with Travis-CI - Now Live!](https://www.reddit.com/r/aws/comments/f17pnb/part_8_deploying_a_vuejs_and_serverless_imgur/)
- url: https://tutorialedge.net/projects/building-imgur-clone-vuejs-nodejs/part-8-deploying-our-app/
---

## [6][Downside only enabling AWS "Business" support as-needed?](https://www.reddit.com/r/aws/comments/f0wfzf/downside_only_enabling_aws_business_support/)
- url: https://www.reddit.com/r/aws/comments/f0wfzf/downside_only_enabling_aws_business_support/
---
We currently pay monthly for AWS "Business" support. We use it \~6x/year. In December when we prepaid a bunch of costs, the % hit was significant.

I saw an AWS savings tip that recommended keeping support off, because you could always enable it as-needed.

Are there downsides to this approach I'm not considering? Does anyone here do this? This is for a business, so if we encountered an emergency we need to contact someone in &lt;1 hour. But would that timeframe work within this approach?
## [7][Do you experience many outrages on AWS?](https://www.reddit.com/r/aws/comments/f0udvx/do_you_experience_many_outrages_on_aws/)
- url: https://www.reddit.com/r/aws/comments/f0udvx/do_you_experience_many_outrages_on_aws/
---
I have some production servers (file, database, web, and backups) for my business hosted on Linode.com.  So far, I'm pretty happy with the support and they have great documentation.  My only complaint is that I'll receive an email once every other week or so telling me there's an issue on the physical box and that they need to do emergency maintenance and that my server will be down anywhere from 30 to 60 minutes.

Usually my business website is down anywhere from a couple of minutes to an hour, much to the dismay of my customers.  I'm wondering, how frequently do you experience outages with your AWS virtual machines?  I'm really tired of this and am thinking of switching.  Thanks!
## [8][How do engineers/developers use custom domain names?](https://www.reddit.com/r/aws/comments/f13rz1/how_do_engineersdevelopers_use_custom_domain_names/)
- url: https://www.reddit.com/r/aws/comments/f13rz1/how_do_engineersdevelopers_use_custom_domain_names/
---
So, If you're building something with amazon web services and instead of having apis coded with lengthy amazon-id and other stuff, how do you replace this with domain names? 

Do all devs keep a registered domain name? 

And if only the product being built is for gauging the users before going fully functional, are there options I'm unaware for using domain names ?
## [9][Program that will sync AWS S3 files from local computer?](https://www.reddit.com/r/aws/comments/f1722v/program_that_will_sync_aws_s3_files_from_local/)
- url: https://www.reddit.com/r/aws/comments/f1722v/program_that_will_sync_aws_s3_files_from_local/
---
Hi so I keep updating files in my S3 bucket. I want a program (for Mac) that will sync only changed files from my computer to the s3 bucket. 

I know Forklift is available for this. Any other program for the mac? Thanks!
## [10][Image processing lambda](https://www.reddit.com/r/aws/comments/f13mxs/image_processing_lambda/)
- url: https://www.reddit.com/r/aws/comments/f13mxs/image_processing_lambda/
---
Im looking to replace a memory hogging image processing call in my .Net API with a nice fast lambda that does the same things.  I need resize, yes, but also some effects like Mono, BW, contrast, sharpness, transparency, posterization and a couple of others.

I can see plenty of examples using 'sharp' but Im not sure if this will do everything I need.

Any suggestions? Would prefer node, but can cope with python if I HAVE TO :)

Thanks
