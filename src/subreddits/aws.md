# aws
## [1][How about cost alert latency?](https://www.reddit.com/r/aws/comments/idv9gc/how_about_cost_alert_latency/)
- url: https://www.reddit.com/r/aws/comments/idv9gc/how_about_cost_alert_latency/
---
Hello!

I'm wondering how reliable are cost alerts with regard to latency?

As in what point in time does the underlying cost calculation represent? -1 hour, -10 hours, -1 day?

Does somebody have some experience with this?

Let's say I program a kill switch to go off at $1000 of costs accumulated since the beginning of the month. Then it would we be inconvenient if those $1000 represent the cost from several hours ago and due to an unexpected extreme spike the actual cost are already like $2000 f.x.
## [2][How to choose a Database on AWS](https://www.reddit.com/r/aws/comments/idpggn/how_to_choose_a_database_on_aws/)
- url: https://www.reddit.com/r/aws/comments/idpggn/how_to_choose_a_database_on_aws/
---
Hi all,

I recently put together a video resource highlighting a methodology to choose the most appropriate database for your use-case on AWS. I wanted to share for folks here who may be interested.

Video is available here: https://youtu.be/eK_umMYxZfM

Cheers
## [3][Just took over as sysadmin, it director, Helpdesk etc. Noticed they were using AWS. Everyone is logging in under the root account, no MFA. How do I proceed?](https://www.reddit.com/r/aws/comments/idegep/just_took_over_as_sysadmin_it_director_helpdesk/)
- url: https://www.reddit.com/r/aws/comments/idegep/just_took_over_as_sysadmin_it_director_helpdesk/
---
Only their marketing department uses AWS. I was thinking of creating a marketing group, assigning users to that group. Giving that group access to S3 and whatever else they are using. It seems they built everything under root and store their data there. Will this be a problem? I am studying for me SAA so I know some of this but not everything.

edit: so I told them about it and they said thank you and want me to fix it. I am studying for my SAA and have enough knowledge to get things working and all. The only problem is I was hired to be a sys admin at a sys admin pay. They have me doing things way outside the scope of my responsibilities that was discussed during the interview. Bait and switch... I have to do IT consulting, help desk, syadmin, it director stuff, now AWS? My first day on the job they wanted me to document their entire network, and create a proposal. Should I renegotiate my hourly rate?
## [4][Speed up data sync from S3 to ec2](https://www.reddit.com/r/aws/comments/idwio0/speed_up_data_sync_from_s3_to_ec2/)
- url: https://www.reddit.com/r/aws/comments/idwio0/speed_up_data_sync_from_s3_to_ec2/
---
Im looking for advice, I have a compute job that runs on an EC2 once a month.  I've optimized the job so that it runs within an hour, however the biggest bottleneck to date is syncing thousands of csv files to the machine before the job starts.  

If it helps the files are collected every minute from hundreds of weather stations, what are the options?
## [5][S3-hosted website - cross-region failover](https://www.reddit.com/r/aws/comments/idw0nk/s3hosted_website_crossregion_failover/)
- url: https://www.reddit.com/r/aws/comments/idw0nk/s3hosted_website_crossregion_failover/
---
Hi, I'm designing a cross-region disaster recovery plan, migrating my system to another AWS region. I'm having trouble with migrating my S3-hosted frontend website. 

I'll use [subdomain.domain.com](https://subdomain.domain.com) as my example URL. I've currently got a bucket called [subdomain.domain.com](https://subdomain.domain.com), and a CNAME in domain.com's DNS config, with \`name: subdomain.domain.com\`, \`value: [s3-website.eu-west-2.amazonaws.com](https://s3-website.eu-west-2.amazonaws.com)\`. 

For my cross-region DR, I want to point incoming [subdomain.domain.com](https://subdomain.domain.com) traffic at a different S3 bucket, hosted in a different region. This other bucket is called [disaster-recovery.domain.com](https://disaster-recovery.domain.com) and has a CNAME with \`name: disaster-recovery.domain.com\`, and \`value: s3-website.eu-west-1.amazonaws.com\`. 

For DR testing purposes, I am trying to direct incoming [subdomain.domain.com](https://subdomain.domain.com) traffic to the [disaster-recovery.domain.com](https://disaster-recovery.domain.com) S3 bucket. I would expect this to work if I change my [subdomain.domain.com](https://subdomain.domain.com) CNAME's value from [s3-website.eu-west-2.amazonaws.com](https://s3-website.eu-west-2.amazonaws.com) to [disaster-recovery.domain.com](https://disaster-recovery.domain.com). However, accessing [subdomain.domain.com](https://subdomain.domain.com) on a cache-cleared browser serves me content from the original [subdomain.domain.com](https://subdomain.domain.com) bucket, not from the [disaster-recovery.domain.com](https://disaster-recovery.domain.com) bucket.

Could someone please give me some insights on how I can get this to work?

&amp;#x200B;

Notes:

* I know that CloudFront offers automated regional failover, but I would prefer to not have to add CloudFront to my stack
* I use CloudFlare to manage my domain, and it's CDN service caches my frontend content. I ran custom purges before every refresh
## [6][API Gateway HTTP APIs adds integration with five AWS services](https://www.reddit.com/r/aws/comments/idl6kf/api_gateway_http_apis_adds_integration_with_five/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/08/api-gateway-http-apis-adds-integration-with-five-aws-services/
---

## [7][How To Make Your Custom Software HIPAA Compliant](https://www.reddit.com/r/aws/comments/idtg44/how_to_make_your_custom_software_hipaa_compliant/)
- url: http://hipaa-compliant.icu
---

## [8][Is possible to use Lightsail without buying a Database?](https://www.reddit.com/r/aws/comments/idwqg4/is_possible_to_use_lightsail_without_buying_a/)
- url: https://www.reddit.com/r/aws/comments/idwqg4/is_possible_to_use_lightsail_without_buying_a/
---
Lightsail comes by default with WordPress installed which means that the server already has MySQL DB.

Can I continue using WordPress without buying a database plan?
## [9][Announcing the newest AWS Heroes – August 2020](https://www.reddit.com/r/aws/comments/idg34l/announcing_the_newest_aws_heroes_august_2020/)
- url: https://aws.amazon.com/blogs/aws/announcing-the-newest-aws-heroes-august-2020/
---

## [10][[Question] How to deploy lambdas and RDS with bitbucket pipelines for CI/CD?](https://www.reddit.com/r/aws/comments/idut7m/question_how_to_deploy_lambdas_and_rds_with/)
- url: https://www.reddit.com/r/aws/comments/idut7m/question_how_to_deploy_lambdas_and_rds_with/
---
Hey everyone,

Hope everyone is well.

I am looking for some advice, and to learn from the way you deploy at your current workplace.

We started doing serverless at my company, and automated deploys to 3 different envs (dev, staging, prod) with bitbucket pipelines. These pipelines run a linter, unit testing, and type checking (mypy as we're using python). At the end of it all, it uses AWS SAM to deploy.

Now we are getting to the stage were we want to deploy a "test" version before we replace a currently working version - a canary. 

I did that by hacking the template with \`{{$TEST\_SUFFIX}}\` and doing a \`sed\` command to create one test template. It creates all the new infrastructure (api gateway + lambda) then runs requests with pytest against it, and if it all passes, it tears it down and deploys to the actual environment. This seems unwieldy, and not the right way to do it... How do you do it at your workplace?

Furthermore, I now would like to spin up an RDS instance so that my tests can run against the lambda, and I can check the database to see if the item created is what I expect. This now raises questions around doing migrations, seeding data, etc. It keeps growing and getting more confusing.

So how do you guys keep all of this well managed, and clear? I hope what I wrote down makes sense, if you guys need any further explanation, let me know.

Appreciate your help - hope I can write a blog post when this ends to help clear the confusion others might also have :)

Best, ze
