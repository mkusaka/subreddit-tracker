# aws
## [1][Diagrams as code (Python) with AWS icon support](https://www.reddit.com/r/aws/comments/gjxil2/diagrams_as_code_python_with_aws_icon_support/)
- url: https://diagrams.mingrammer.com/docs/getting-started/examples
---

## [2][I wrote a guide on building your own resume site using AWS S3, CloudFront, and Route53. Its almost free to host and looks great on a resume.](https://www.reddit.com/r/aws/comments/gjl5wi/i_wrote_a_guide_on_building_your_own_resume_site/)
- url: https://seanjziegler.com/how-to-build-a-free-static-resume-site-with-aws-s3-cloudfront-and-route-53/
---

## [3][Has anyone tried using the new, revamped Macie?](https://www.reddit.com/r/aws/comments/gk95tq/has_anyone_tried_using_the_new_revamped_macie/)
- url: https://www.reddit.com/r/aws/comments/gk95tq/has_anyone_tried_using_the_new_revamped_macie/
---
I set up the new Macie yesterday and tested it by throwing some documents with a bunch of (fake) PII into a bucket: a CSV, a DOCX, and an XLSX. It appears when I create a job that the job is only scanning the CSV file because that's the only finding that's coming back. No mention of any findings from the other file types. Am I doing something wrong? I don't see how I could be as the new interface seems very, very simple. Thanks for any help!
## [4][I made a dedicated site for the Cloud Resume Challenge - an AWS project you can do to level up your skills and get noticed](https://www.reddit.com/r/aws/comments/gk94vd/i_made_a_dedicated_site_for_the_cloud_resume/)
- url: https://cloudresumechallenge.dev/
---

## [5][S3 Bucket Configured As Browsable Folders?](https://www.reddit.com/r/aws/comments/gk8x5m/s3_bucket_configured_as_browsable_folders/)
- url: https://www.reddit.com/r/aws/comments/gk8x5m/s3_bucket_configured_as_browsable_folders/
---
Hello Everyone:

First time posting here and I hope someone can help me out.

My company currently creates hardware solutions and provides firmware updates for said hardware via a public ftp server.  We're deprecating this server for a variety of reasons and we're thinking of moving it to an S3 bucket.  I've setup buckets before, but have always put up a web server or something in front where folks are not hitting the bucket directly.

The need here, because we have thousands of customers and don't want to manage usernames/passwords, is to setup an S3 bucket where someone can browse the folders, find the appropriate firmware for their particular hardware device and download it.

Is this doable with just and S3 bucket or should we be thinking of this differently?
## [6][Passing api token and properykey value with python 3 to the API gateway using GET.](https://www.reddit.com/r/aws/comments/gk8n5c/passing_api_token_and_properykey_value_with/)
- url: https://www.reddit.com/r/aws/comments/gk8n5c/passing_api_token_and_properykey_value_with/
---
I have the basic API call figured out, I can create an API and make calls to it. 

But for the life of me I can not figure out what the proper way to pass a propertykey value and API token to get a valid response from AWS API.

Can someone shine some light on this?
## [7][SSL for many customer whitelabel domains pointing to Elastic Beanstalk app?](https://www.reddit.com/r/aws/comments/gk8cz1/ssl_for_many_customer_whitelabel_domains_pointing/)
- url: https://www.reddit.com/r/aws/comments/gk8cz1/ssl_for_many_customer_whitelabel_domains_pointing/
---
We're considering deploying an app (maybe two) via Elastic Beanstalk. The complication is that the app is served via multiple (whitelabel) domains. We own most of the domain registrations - some customers own theirs and configure A records based on our instructions (we need to move them to CNAMES...).

  
**My question: What's the best way to generate and manage SSL certificates in such a setup?**  


My first thought was, for each customer:

* Move DNS hosting to Route 53 to use ACM DNS validation
* Delete existing A record and change to alias to point at Elastic Beanstalk load balancer
* Generate SSL certificate for domain, or add domain to existing SSL certificate(\*)

(\*) We have relatively large numbers of domains, let's say \~200 pointing at a serviceְ. Given that the load balancer has a limit of 25 SSL certificates, we might create an SSL certificate for each starting letter of the alphabet, and add domains to each as necessary. Is this possible? How would Elastic Beanstalk / Application Load Balancers cope with such a scenario.

  
We add domains all the time as customers join, and occasionally remove them as they leave. Currently everything is manual, but we do have an internal onboarding system for creating new customer platforms, and it would be relatively easy to use the AWS SDK to create domains and SSL certificates as needed.

&amp;#x200B;

Someone mentioned doing SSL termination using nginx / letsencrypt in a Docker instance with something like [https://github.com/linuxserver/docker-letsencrypt](https://github.com/linuxserver/docker-letsencrypt) and a NLB rather than ALB that Elastic Beanstalk provides, but  I'm not altogether excited about that because:

* We don't have any Docker infrastructure at the moment, so I have to learn ECS or similar and get all that infrastructure running.
* Would I need to run one container per domain? That could be \~1000 nginx processes, I'm not sure how that would scale. NB My Docker-fu is rubbish, mostly things I've read and forgotten, because I've never had to use it in production.
* Not sure how our internal onboarding system would spin containers up and down based on adding and removing domain names. Perhaps the ECS API? I'm really green here.

Happy to clarify if anything isn't clear, look forward to your feedback.
## [8][Change CloudWatch log interval for AWS Batch](https://www.reddit.com/r/aws/comments/gk6lgf/change_cloudwatch_log_interval_for_aws_batch/)
- url: https://www.reddit.com/r/aws/comments/gk6lgf/change_cloudwatch_log_interval_for_aws_batch/
---
I have an AWS Batch job that's running and having it's CloudWatch log under the `/aws/batch/job` log group. But what I've noticed is that my log events gets compiled and dumped all the same time every 25-30 minutes. Is there anyway to change the interval for it?
## [9][Amazon Lex with Discord bot, any idea?](https://www.reddit.com/r/aws/comments/gk6it9/amazon_lex_with_discord_bot_any_idea/)
- url: https://www.reddit.com/r/aws/comments/gk6it9/amazon_lex_with_discord_bot_any_idea/
---
So I have a task using AWS services and since I did spent a lot of time on discord, I'm thinking of making a chat-bot and deploy it as discord bot. Sadly there's not a lot of tutorial available out there. Does anyone has a clue if this is possible and what would the infrastructure from discord to amazon lex looks like? Thank you
## [10][IAM policy with tags to restrict access to the certain resources in the organisation](https://www.reddit.com/r/aws/comments/gk5p7t/iam_policy_with_tags_to_restrict_access_to_the/)
- url: https://www.reddit.com/r/aws/comments/gk5p7t/iam_policy_with_tags_to_restrict_access_to_the/
---
Hey guys,

&amp;#x200B;

so here's the situation. I have the IAM policy which supposed to deny access to the resources marked with the tag. The policy looks like this:

        {
          "Effect": "Deny",
          "Action": "*",
          "Resource": "*",
          "Condition": {
            "StringEquals": {
              "aws:ResourceTag/Project": "EMS"
            }
          }
        }

It looks like the policy should deny all actions on the resource with the \`Project=EMS\` tag. 

However, when I tag the bucket, EC2 instance, etc I can still see and execute actions (with a test user) on those particular resources just like on the others without the tag.

&amp;#x200B;

What am I doing wrong here?

&amp;#x200B;

Thanks!
