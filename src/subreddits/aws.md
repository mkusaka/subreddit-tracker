# aws
## [1][Maximising uptime with single EC2 Spot instance](https://www.reddit.com/r/aws/comments/f4o40u/maximising_uptime_with_single_ec2_spot_instance/)
- url: https://www.reddit.com/r/aws/comments/f4o40u/maximising_uptime_with_single_ec2_spot_instance/
---
I would like to run a single server as a spot instance with disk persistance.It seems like if it's interrupted then I just have to wait until there is sufficient free capacity in that AZ for it to be restarted by EC2.

I understand that I can use a spot fleet or scaling group to get another instance started in another AZ or size, but I'd rather not use these as I'm constantly using this VM and modifying its disk contents, so I'd rather not have to recreate an AMI each time I make a modification on disk.

I also have lots of servers like this and would rather not have to create an auto scaling group for each server.

Is there any way of getting the size/type flexibility for a single instance that spot fleets provide, so that it can be instantly restarted if it needs to be stopped by EC2 Spot services?

Edit 2: These are Windows servers. Ignore previous comment about only being in use 50hrs per week - I mixed up my requirements.
## [2][Cheapest Serverless Architecture to help out a charity that was abused by a "Lead network Engineer" that had hosted everything of their's on an unpatched win8 pc under a desk with no upc.](https://www.reddit.com/r/aws/comments/f4e249/cheapest_serverless_architecture_to_help_out_a/)
- url: https://www.reddit.com/r/aws/comments/f4e249/cheapest_serverless_architecture_to_help_out_a/
---
I'm using my long weekend to try to assist a local kids charity by moving them off their old desktop PC running under a desk to being hosted in AWS. On Tuesday I'm going to talk to the right people and get this approved, and then do the work for free (I'm unemployed and really bored.)

I'm wanting to go with a serverless approach to keep costs down as much as possible, its a low usage system only accessed by around 1000 people per month and so I would rather they be able to use as much of their budget to their mission as much as possible.  Besides, the old code was bad enough, I wanted to throw up looking at it.

The charity has a person who is trustworthy who knows Python but the prior system was not in python (it was in php3) because of the other "more sr." but still younger tech guy was worried about being replaced.  So any new code I write will be in a modern version of python so that it's maintainable by that organization and the brain trust they already have and can actually trust already.

I'm doing that in an attempt to inoculate them from being abused again by tech-savvy people with no ethics; And the more I can save them on maintenance and operation costs, the worse the guy that was charging them thousands a month to "administrate and maintain" - (read: remote desktop into that win8 pc randomly to run windows update because he had turned the windows update service off in the name of "security") - that production system that was just running under a persons desk will look.

I wont say the name of the charity; They got abused by this other guy and I don't want that mistake by them to reflect badly on their mission or brand; They didn't do anything wrong, they just don't understand tech and got scammed as an innocent victim.  So let's help them, give me your ideas.

Either way, I know I'm possibly biased here.. but I just dont think a crappy php3 app running on a win8 laptop under a desk without a upc is good enough for them.  So I want to save these people as much money as possible going forward.  It helps them, and the more they save going forward the worse it looks for the other guy that abused them.  I really want to make this system as cheap as possible so I can also see the look on his face when we compare the numbers.

I might be biased; But at least my hearts in the right place.

Anyway, Here is what I'm thinking:

- Move domain to route53

- Enable Static HTML/site hosting out of an S3 bucket and host all static content there. That covers a lot of the data move.

- Move the subdomain that handles forms/scripts/login to a subdomain that is covered by an API gateway and calls AWS Lambda for running REST API endpoints needed.

- Find a good but cheap solution for common DB operations that is still consistent/ACID-compliant and can scale as needed as they grow without too much difficulty or technical debt for them to worry about when that happens, using Lambda to run code that queries/updates the database as needed. 

- Cache/memoize as needed.

We all know Lambda etc could handle this easily, but I'm pondering the database system as I want it to be as cheap as possible and also be able to scale up; Last I checked dynamodb was utterly the most expensive option last I looked at this problem so I want to hear the ideas of the community on how this can be best designed, on the off chance that something has changed that I don't know about or somebody has a good idea that I have not considered.

At this point, I'm tempted to give them a small RDS instance and then call it a day.. but I'm hoping there is a better option as I remember from when I did this months/years ago last that Lambda needed a bunch of extra junk to be able to access RDS that made it a lot more expensive and complicated than it really should be IMHO, and I would rather not make this more complicated than I need to because I'm thinking about this from a "how are they going to maintain this" perspective.

I have been explicitly told that Cognito is not an option so I need to stay away from it, but I still need login via email/password with some extra constraints so I figure that's a lambda for login, one for registration, etc.  If you have a better idea please speak up.

My original thinking was "find a vendor that does what they need and let them pay a much smaller bill every month" but due to what they do, and due to other things outside of my or their control, it has to be a custom system.  So I'm here.

**Ok, big EDIT/update based on peoples responses:**

The biggest thing I can call out here is that I also have some pretty big non-technical requirements.  This means that like most projects technical choices are being made based on both technical needs and limitations imposed by the org to mitigate the risk of them not being able to support anything after it starts costing them so much less.  Due to this many of the systems mentioned by people - Elastic Beanstalk, Lightsail, ec2, etc - can't be used due to massively increased cost compared to this more simple solution, or massively increased maintenance requirements that the org in question is NOT READY FOR as they have no full-time devs and couldn't afford to hire one without it deeply impacting their daily mission.

Thank you all for everybody's responses, And my inbox is overflowing and I will get to you all as I can.
## [3][Single-Table Design in DynamoDB](https://www.reddit.com/r/aws/comments/f4q5ha/singletable_design_in_dynamodb/)
- url: http://alexdebrie.com/posts/dynamodb-single-table/
---

## [4][Feature request: Auto-expiring security group rules](https://www.reddit.com/r/aws/comments/f4nbal/feature_request_autoexpiring_security_group_rules/)
- url: https://www.reddit.com/r/aws/comments/f4nbal/feature_request_autoexpiring_security_group_rules/
---
I would make use of security group rules that automatically expired and were deleted.  Sometimes I have a developer that needs temporary SFTP access to a location and it would be great to set a rule that gives his IP access and the rule expires in the configured time period.

(ducking for cover)
## [5][Blue green deployment with namespace segregation. Advice needed](https://www.reddit.com/r/aws/comments/f4ogbs/blue_green_deployment_with_namespace_segregation/)
- url: https://www.reddit.com/r/aws/comments/f4ogbs/blue_green_deployment_with_namespace_segregation/
---
We have an EKS cluster serving a mobile application. We're considering implementing blue green deployment strategy to minimize downtime and to ease rollbacks.

Is it possible to do a blue green setup with namespace segregation? Any resources that can help with getting started?
## [6][When should I be using AWS IoT?](https://www.reddit.com/r/aws/comments/f4lugm/when_should_i_be_using_aws_iot/)
- url: https://www.reddit.com/r/aws/comments/f4lugm/when_should_i_be_using_aws_iot/
---
Hey guys,
I found out about AWS IoT recently, and wanted to learn more about it. So, I got started on a project where I read the humidity and temperature of my room, log it, show it on a website, and send an email if the values pass a threshold. I'm using a Pi Zero W for taking the readings. 
I initially thought of using IoT Core for this, creating a thing, and making a Rule to send the sensor data to the database, and another rule for sending the email. Then I realized that I can just use my Pi to do both those things directly (boto3, and creating an IAM user).

So, I'm curious as to what kind of projects AWS IoT would be more suitable for. Any input on this would be appreciated!

Thank you!
## [7][Glue and the write_dynamic_frame preactions and postactions options help](https://www.reddit.com/r/aws/comments/f4h67l/glue_and_the_write_dynamic_frame_preactions_and/)
- url: https://www.reddit.com/r/aws/comments/f4h67l/glue_and_the_write_dynamic_frame_preactions_and/
---
I am working with a large number of files that hit S3 throughout the the day from several sources.  The are all the same format but can have overlapping records, the good news is that when the records do overlap the are duplicates.

The destination for my ETL is redshift and I am very comfortable with the stage / dedupe / merge techniques.  

&amp;#x200B;

to help control costs i want to fire my glue jobs on a schedule rather than triggering on files arriving.  So when the job runs I may have 10-100 files to process all with potential for some duplicate records.  I am typically using bookmarks and this all works nice when i do not have the potential for duplicates.

&amp;#x200B;

My goal is to use the options above as per  [https://aws.amazon.com/premiumsupport/knowledge-center/sql-commands-redshift-glue-job/](https://aws.amazon.com/premiumsupport/knowledge-center/sql-commands-redshift-glue-job/)   but this only works if the pre and post jobs are run PER FILE.  so Is my glue job issuing a COPY command per file or is it reading all of the available files into the dataframe and performing a single COPY with the pre and post command run one time ?
## [8][EKS with EFS correct way to serve static assets?](https://www.reddit.com/r/aws/comments/f4hqoh/eks_with_efs_correct_way_to_serve_static_assets/)
- url: https://www.reddit.com/r/aws/comments/f4hqoh/eks_with_efs_correct_way_to_serve_static_assets/
---
I'm writing an application to run on EKS. I wanted to use EFS to store the HTML template files and assets (css, images, etc) and attach it as a volume to my EKS clusters for beta and prod. In EFS I planned to have a beta folder and a prod folder, and rsync between them when deploying the templates/assets. The docker image with the code would be deployed separately.

I don't want to build all of the assets into my docker image, because then my image would be huge. This application uses the same code to serve many websites, so we have about 12 GB of asset files and templates, although that's mostly asset files.

&amp;#x200B;

Is this the correct way to handle this? Is EFS fast enough to serve files like this?

&amp;#x200B;

I know that S3 is good for static files, like images, css, etc, but it's not clear to me how that doesn't break my development workflow:

* The HTML templates need to be accessed by my application to be filled in, so can't be stored in S3.
* Seems like it would break relative links when developing, unless putting assets in S3 was part of the deployment process, but I'm not sure if there's a quick efficient way to sync large asset directories with an S3 bucket.
* In production, I'm not sure how I could use the same domain to serve all static assets but then let my application handle the rest of the requests.

I feel like it's possible that the template files could be copied into the image, and asset files could be saved to S3, but I'm not sure if this is easy to do, and I also still feel like this would break my 2nd and 3rd points above.

&amp;#x200B;

Does anyone have any pointers on how to handle this? Sorry if I've missed something obvious, I've tried to do lots of reading on AWS, but this is also my first AWS app.
## [9][How to keep my EC2 WordPress site spinning 24 hours?](https://www.reddit.com/r/aws/comments/f4k0hu/how_to_keep_my_ec2_wordpress_site_spinning_24/)
- url: https://www.reddit.com/r/aws/comments/f4k0hu/how_to_keep_my_ec2_wordpress_site_spinning_24/
---
I followed this link to set up a WordPress site on EC2:
https://www.youtube.com/watch?v=U3VSJhaC4kc&amp;t=18s

At 2:47:49 of the course the EC2 instance was terminated, which means the website is no longer accessible. If we need it spinning 24 hours, do we just leave the instance as it is and log out our account afterwards?

I also configured my domain name on Route53 to point to the public IP of the EC2 instance following this link: https://www.youtube.com/watch?v=DYqfdw4Kvbg.

And then I entered my domain url in laptop browser and cell phone browser. However, only my cell phone browser opened the url successfully and my laptop browser failed.

Please give me some help on how to route the domain name to public IP of EC2 instance and keep the instance spinning so we can keep and save our work on the WordPress page.
## [10][Multi-Attach for Provisioned IOPS (io1) EBS Volumes](https://www.reddit.com/r/aws/comments/f41kb2/multiattach_for_provisioned_iops_io1_ebs_volumes/)
- url: https://aws.amazon.com/blogs/aws/new-multi-attach-for-provisioned-iops-io1-amazon-ebs-volumes/
---

