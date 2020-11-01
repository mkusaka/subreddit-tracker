# aws
## [1][re:Invent registration is now open](https://www.reddit.com/r/aws/comments/jkenu3/reinvent_registration_is_now_open/)
- url: https://register.virtual.awsevents.com/
---

## [2][5 Tips to Make Your Lambda Functions Run Faster (and Cheaper)](https://www.reddit.com/r/aws/comments/jlt5fn/5_tips_to_make_your_lambda_functions_run_faster/)
- url: https://www.webiny.com/blog/5-tips-to-make-your-lambda-functions-run-faster-and-cheaper
---

## [3][Ways to increase availability of my API Gateway -&gt; Lambda -&gt; DynamoDB Setup](https://www.reddit.com/r/aws/comments/jlyv26/ways_to_increase_availability_of_my_api_gateway/)
- url: https://www.reddit.com/r/aws/comments/jlyv26/ways_to_increase_availability_of_my_api_gateway/
---
I'm currently setting up a very small lambda function. Basically it takes some information related to website analytics (where the user click, who the user is, what time it was performed).

This saves it to DynamoDB and accessed via API Gateway (currently has no authentication method).

I'm trying to test it using Artillery and I'm getting 502 after some time (when the load starts to increase).

I need some tips on what settings I could change to reach 10,000 API requests per minute. I'm very new to this set-up so please pardon this noob.
## [4][Multiple Cost Allocation Tags in Cost Explorer or Cost Categories?](https://www.reddit.com/r/aws/comments/jlyw7i/multiple_cost_allocation_tags_in_cost_explorer_or/)
- url: https://www.reddit.com/r/aws/comments/jlyw7i/multiple_cost_allocation_tags_in_cost_explorer_or/
---
I've setup and activated seperate Cost Allocation Tags for Jenkins CI EC2 runs "win2012_qa",  "win2012_qa_manual". I enter these tags in the corresponding AMI/EC2 entries in Jenkins and launch some jenkins jobs which use the tags.

It works as expected, I can view each tag individually fine in Cost Explorer.  But I cannot get multiple Cost Allocation Tags to present in Cost Explorer when I select more than 1 tag.

I then tried setting up a "QA" Cost category with "win2012_qa" and "win2012_qa_manual" as seperate rules with Dimension=Tag and Tag key+value = "win2012_qa".   Same for "win2012_qa_manual".   Wait 8 hours for it to process.  But it seems to only display 1 tag.

Is there any way I can have multiple tags presented in Cost Explorer?
## [5][AWS MFA with the YubiKey Manager CLI](https://www.reddit.com/r/aws/comments/jlmi8k/aws_mfa_with_the_yubikey_manager_cli/)
- url: https://www.reddit.com/r/aws/comments/jlmi8k/aws_mfa_with_the_yubikey_manager_cli/
---
I've been searching for a powerful CLI based application to manage hard tokens for my root and iam accounts. Recently I settled on the [__YubiKey Manager CLI app__](https://support.yubico.com/hc/en-us/articles/360016614940-YubiKey-Manager-CLI-ykman-User-Manual) in combination with the [__5c__](https://www.yubico.com/product/yubikey-5c/). Since the CLI app uses [__OATH-TOTP__](https://en.wikipedia.org/wiki/Time-based_One-time_Password_algorithm) it has made the experience quite pleasing and wanted to share my [__three part article__](https://blog.shaneyost.io/). A couple people on here had great advice on this application so I wanted to say thanks.
## [6][Lightsail vs. EC2 for my project](https://www.reddit.com/r/aws/comments/jlyryd/lightsail_vs_ec2_for_my_project/)
- url: https://www.reddit.com/r/aws/comments/jlyryd/lightsail_vs_ec2_for_my_project/
---
I’m making a project where my Arduino sends an email/sms every time a sensor condition changes. To accomplish this I need a server to connect to SES and Twilio’s SMS API.

I have limited experience with AWS so I’m really not sure if I should use Lightsail or EC2 for this. If the Arduino would only need to connect to connect to the server a couple times a day (about 600/year), which would be cheaper, the $3.50 Lightsale tier or EC2 (would this usage be within EC2’s free 750 hours?)
## [7][Error during validation of WorkSpace image](https://www.reddit.com/r/aws/comments/jm1g58/error_during_validation_of_workspace_image/)
- url: https://www.reddit.com/r/aws/comments/jm1g58/error_during_validation_of_workspace_image/
---
hey,

 I'm slowly despairing... I try to create an image with AWS but everytime I get the error " Error during validation of WorkSpace image". I installed a lot of software in it. Last time (6 months ago) everything worked fine.

I have checked all of these steps: [https://aws.amazon.com/de/premiumsupport/knowledge-center/workspaces-image-creation-issue/](https://aws.amazon.com/de/premiumsupport/knowledge-center/workspaces-image-creation-issue/)  


The imagechecker tool says that everything is passed and fine.   
 https://imgur.com/47qkidd   
Does anybody have an advice for me? 

Many greetings
## [8][Help using S3 as Google One replacement for sharing](https://www.reddit.com/r/aws/comments/jlz411/help_using_s3_as_google_one_replacement_for/)
- url: https://www.reddit.com/r/aws/comments/jlz411/help_using_s3_as_google_one_replacement_for/
---
Hi,

I was wondering if anybody is using S3 to replace Google Drive. I am considering this because Google options for having storage shared with family do not work for me (we live in different countries).

My plan would be to make a IAM user for my parents with privileges to log into the console and only up/download from a bucket. I guess they should be OK to learn how to do this.

I would be grateful for any insights on possible showstopers and/or hidden costs that my crop up from other necessary infrastructure.

\- Looking at costs per GB / transfer sound like I'd be paying only slightly more than Google Drive. I am assuming transfer cost "to the internet" are the highest transfer costs, correct?

\- Are IAM users and roles free?

\- Could I add other permissions/limitations (different users for different family members, quotas, upload limits) using only S3, or would this need extra components? If only in S3, would this incurr extra costs?

\- What could my parents do to screw this up? XD

&amp;#x200B;

My quick google of a simple solution to add, e.g., a Web UI for such an application did not bring up much, but if there is somethig ready-made, it'd also be nice to take a look.

&amp;#x200B;

Thanks in advance for any help!

Juan
## [9][writing from EMR to AWS Elasticsearch](https://www.reddit.com/r/aws/comments/jlz9v1/writing_from_emr_to_aws_elasticsearch/)
- url: https://www.reddit.com/r/aws/comments/jlz9v1/writing_from_emr_to_aws_elasticsearch/
---
Hi,

I'm convertin my code to from spark on (EC2) to EMR, but I can't find any relevant document regarding   
writing from spark job to aws elasticsearch.

I used ( on my EC2)  spark-hadoop jar, but now I not sure where i should add \\ configure this jar( es-haddop):  
[http://www.elastic.co/downloads/hadoop](http://www.elastic.co/downloads/hadoop)

Thanks 

Yanka
## [10][RDS MySQL keep 5% of disk space free permanently ?](https://www.reddit.com/r/aws/comments/jlz12b/rds_mysql_keep_5_of_disk_space_free_permanently/)
- url: https://www.reddit.com/r/aws/comments/jlz12b/rds_mysql_keep_5_of_disk_space_free_permanently/
---
A team is doing some bad SQL queries which are causing the innodb temporary tables on-disk to use up the full remaining disk space of the DB instance.

Because the DB Disk space gets full, I cannot restart/reboot the DB instance (the option is not available in the GUI).  I either have to get AWS support involved to reboot the DB, or manually kill the DB Connections and processes and increase the DB Disk size which will then allow me reboot the DB to clear the temporary tables and return the DB size to normal.

While the team investigate and fix the bad queries, is there any way I can make the DB instance keep a % of disk space free so I can reboot the DB instance ?

The below has great information on how to set the temporary table sizing, but unfortunately the queries are forcing the tables to write to disk.

https://aws.amazon.com/blogs/database/best-practices-for-configuring-parameters-for-amazon-rds-for-mysql-part-1-parameters-related-to-performance/
## [11][What have you automated with python ?](https://www.reddit.com/r/aws/comments/jlstkv/what_have_you_automated_with_python/)
- url: https://www.reddit.com/r/aws/comments/jlstkv/what_have_you_automated_with_python/
---
Hey everyone, I really need real case example on what you have automated in AWS. I feel like I hit a brick wall. I have been learning so much python for my sys role but no idea where to start automating in my company.Perhaps I would get some ideas if I get to hear to your automation achievement.
