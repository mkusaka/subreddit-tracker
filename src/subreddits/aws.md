# aws
## [1][Store and Access Time Series Data at Any Scale with Amazon Timestream – Now Generally Available](https://www.reddit.com/r/aws/comments/j2wz4x/store_and_access_time_series_data_at_any_scale/)
- url: https://aws.amazon.com/blogs/aws/store-and-access-time-series-data-at-any-scale-with-amazon-timestream-now-generally-available/
---

## [2][Amazon S3 on Outposts Now Available](https://www.reddit.com/r/aws/comments/j2vkja/amazon_s3_on_outposts_now_available/)
- url: https://aws.amazon.com/blogs/aws/amazon-s3-on-outposts-now-available/
---

## [3][Amazon Timestream is now Generally Available](https://www.reddit.com/r/aws/comments/j2yz2v/amazon_timestream_is_now_generally_available/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/09/amazon-timestream-now-generally-available/
---

## [4][How to deploy Static app to AWS S3](https://www.reddit.com/r/aws/comments/j38ecn/how_to_deploy_static_app_to_aws_s3/)
- url: https://codingflamingo.com/blog/how-to-deploy-static-site-to-aws-s3-bucket-from-github
---

## [5][AWS Web Scraper architecture advice](https://www.reddit.com/r/aws/comments/j3860g/aws_web_scraper_architecture_advice/)
- url: https://www.reddit.com/r/aws/comments/j3860g/aws_web_scraper_architecture_advice/
---
  I have an EC2 instance with a Web Server which reads from a local MongoDB. I'd like the data in the database to be updated on a weekly basis by performing web scraping (100k+ rows). For that I've set up 2 SQS queues, Lambda functions for doing scraping in parallel, and 3 scripts: one is reading from the DB and inserting data in SQS, another one is checking for the job to be completed and the third one is reading from SQS and updating the database. Since the database is located on the same instance and not externally available I had to put all 3 scripts on the instance as well, but in order to separate them somehow I ran them inside docker containers. This is the solution I ended up with, but I think it's far from perfect, and would like to hear any advice on how this can be improved. Thanks. 

&amp;#x200B;

https://preview.redd.it/wkejeswoehq51.png?width=1800&amp;format=png&amp;auto=webp&amp;s=6e9c4ebb1518c727fa9aad7f271c9799765afd62
## [6][How can I mass audit user access for an environment using federated on premise Active Directory?](https://www.reddit.com/r/aws/comments/j37ujd/how_can_i_mass_audit_user_access_for_an/)
- url: https://www.reddit.com/r/aws/comments/j37ujd/how_can_i_mass_audit_user_access_for_an/
---
I would like to be able to show auditors a report of all of my AD users that have access into AWS, what they have done and when. 

Cloudtrail is on but how can I get everyone's access at once?
## [7][Managing long running scripts](https://www.reddit.com/r/aws/comments/j37hsj/managing_long_running_scripts/)
- url: https://www.reddit.com/r/aws/comments/j37hsj/managing_long_running_scripts/
---
I have a series of node scripts which listen to different sources (websockets or long polling mainly) for the purpose of data collection. The actual script is modified slightly depending on the customer it is being run for. The script is basically a while(true) which runs for 6 to 72 hours. The scripts have grown and gotten pretty convoluted over time, rewriting them isn't on the table at the moment. 

Currently, we are just ssh'ing into an ec2 instance and running them manually. This is quickly becoming something we can no long handle manually. We've also had issues with our ec2's underlying hardware being degraded. 

We'd like to have a web interface where we can "launch" these scripts. I have no concerns for the web interface, but I do have some about how to manage running these scripts. 

*It seems like ec2 or ecs are the only real options for running this sort of workload. Am I missing any others?* I'd like to avoid managing ec2 instances as much as possible because of the previous degradation issues, but *I wonder if the same sort of degredation can happen with fargate?* If it can (or even if ecs isn't susptable), *what is the recommended way or pattern for running a process in the cloud which needs to listen for incoming data when you don't want to miss any data?*

Because these scripts are tweaked on a regular basis, we'd like to store them on S3 so we can replace them easily.  Based on that, *how can we automate taking a script from S3 and running it somewhere*? If we go the ecs approach, I guess we'd have to have our web interface create a docker container, deploy it to ecr, deploy that to ecs? I am not familiar with ecs, is that approach recommended, is there a more standard approach? And if there is hardware degradation, what are the recommendations for restarting a process like this?
## [8][How do you manage your Cloud Credentials?](https://www.reddit.com/r/aws/comments/j3630o/how_do_you_manage_your_cloud_credentials/)
- url: https://www.reddit.com/r/aws/comments/j3630o/how_do_you_manage_your_cloud_credentials/
---
It's something that I and my teammates struggled a lot with.  


Today I launched my open-source tool to help us managing credentials in a secure way.  


How do you manage them? It's the AssumeRole and AssumeRoleWithSAML the better method?

[https://github.com/Noovolari/leapp](https://github.com/Noovolari/leapp)  


To me save credentials from an AWS user to the \~/.aws file is insecure and I preferred to store them in the Keychain. It is the right way?
## [9][How do you get visibility over cloud workloads?](https://www.reddit.com/r/aws/comments/j2vto0/how_do_you_get_visibility_over_cloud_workloads/)
- url: https://www.reddit.com/r/aws/comments/j2vto0/how_do_you_get_visibility_over_cloud_workloads/
---
I oversee our cloud infrastructure and we have hundreds of EC2 instances, S3 buckets, other cloud resources. How do you know the *purpose* or the *business reason* for the cloud resource? Tagging is the only way I know how, but tags can be inconsistent and unreliable.  It's pretty much all a black box for me, but I want better visibility into the box.  For example, which are SFTP servers, which are pipelines, which are web servers, etc.  Also, even with cost explorer I'm not able to pinpoint where any cost increases are coming from - it only gives me the AWS service which is not helpful if I'm running hundreds of EC2 and the workload is ephemeral. Does anyone else have this problem? What do you use? Do I just need to shore up my tags?
## [10][My EBS volume is frozen and when I attach it to an instance the instance is frozen. How can I unfreeze the EBS volume?](https://www.reddit.com/r/aws/comments/j34hpt/my_ebs_volume_is_frozen_and_when_i_attach_it_to/)
- url: https://www.reddit.com/r/aws/comments/j34hpt/my_ebs_volume_is_frozen_and_when_i_attach_it_to/
---

