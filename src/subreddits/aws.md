# aws
## [1][Amazon Career Day](https://www.reddit.com/r/aws/comments/ir5svb/amazon_career_day/)
- url: https://www.reddit.com/r/aws/comments/ir5svb/amazon_career_day/
---
https://www.amazoncareerday.com/

Disclaimer: I work at AWS as a Technical Consultant. This is *not* a third party recruiting link. I do not receive any kickbacks, referral fees, nor am I recruiter. I just thought that this would be of interest to this subreddit.
## [2][Reminder - New S3 Buckets after 9/30/2020 won't support Path-Style Requests](https://www.reddit.com/r/aws/comments/iqwv17/reminder_new_s3_buckets_after_9302020_wont/)
- url: https://www.reddit.com/r/aws/comments/iqwv17/reminder_new_s3_buckets_after_9302020_wont/
---
If you want a bucket you can access using path style requests (e.g. [https://s3.Region.amazonaws.com/bucket-name/key\_name](https://s3.Region.amazonaws.com/bucket-name/key_name)) you have just over two weeks to create your bucket.

[https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html](https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html)

&amp;#x200B;

&gt;Important  
&gt;  
&gt;Buckets created after September 30, 2020, will support only virtual hosted-style requests.                                                                                  Path-style requests will continue to be supported for buckets created on or before                                           this date.
## [3][I'm building a search engine for terminal commands and just added support for aws-cli. What do you think?](https://www.reddit.com/r/aws/comments/ir0ln7/im_building_a_search_engine_for_terminal_commands/)
- url: https://www.reddit.com/r/aws/comments/ir0ln7/im_building_a_search_engine_for_terminal_commands/
---
Hey all!

I'm working on an app that works as a search engine for anything terminal related. It's called [Sidekick](https://getsidekick.app). The idea is instead of searching stuff on Google and going through websites you just hit a shortcut, type what you want with natural language and the app gives you the command you need and copies it into the clipboard.

Currently, the app supports full-text search through your history, history of your team, and popular CLIs like aws-cli, \`gcloud, heroku, etc. You can even quickly add flags, options, and parameters + look into some basic documentation for the CLI. 

Searching for commands with Sidekick is much faster than with Google. You don't have even use your mouse. More importantly, you don't wait for websites to load and don't need to manually copy and paste commands from these websites.  


I added a short video showcasing the basic functionality.

[Sidekick - AWS search](https://reddit.com/link/ir0ln7/video/36isu8yuflm51/player)

The app is in private beta right now but I'm sending out invites every week or so (happy to prioritize you if you want to try it out). **You can sign up** [**here**](https://getsidekick.app).  


I'd love to know what you think! Thanks.
## [4][AWS Taxation](https://www.reddit.com/r/aws/comments/irbd8s/aws_taxation/)
- url: https://www.reddit.com/r/aws/comments/irbd8s/aws_taxation/
---
I have a UK business with a valid tax number yet I'm still charged for tax? Do all business pony up 20% and wait for a tax return or am I missing something?
## [5][Global Accelerator with Media Services](https://www.reddit.com/r/aws/comments/ira01m/global_accelerator_with_media_services/)
- url: https://www.reddit.com/r/aws/comments/ira01m/global_accelerator_with_media_services/
---
Hi there,

Wondering if anyone has a workflow to allow MediaLive and MediaStore to be used as an endpoint for Global Accelerator?
## [6][Does AWS translate ever produce other artifacts, more so than google translate?](https://www.reddit.com/r/aws/comments/ir9x0q/does_aws_translate_ever_produce_other_artifacts/)
- url: https://www.reddit.com/r/aws/comments/ir9x0q/does_aws_translate_ever_produce_other_artifacts/
---
I'm doing a translation and aws seems to have produced a new artifact.  I translated the material  in aws, then translated it back with google translate to check it.  Then translated it back into aws and now there is a new word that completely changes the meaning of the material.  google translate is translating it perfectly now from the aws original translation, which should I go with, neural machine or google?
## [7][Best way to organize Discord Bot data in DynamoDB?](https://www.reddit.com/r/aws/comments/ir77vq/best_way_to_organize_discord_bot_data_in_dynamodb/)
- url: https://www.reddit.com/r/aws/comments/ir77vq/best_way_to_organize_discord_bot_data_in_dynamodb/
---
With my table, I want to answer these primary questions:

* How many wins does the user have in rock, paper, scissors?
* How many wins does the user have in boxing?
* Who has the top 10 wins in the server?

Keep in mind, I'm trying to go with the 'single' table approach. On a 'guild' object, there will be many configurations that exist: moderation, user levels, etc.

I also am going to be using the 'user' entity to keep track of economy as well: item purchases, inventory, etc.

[My current table design is here.](https://imgur.com/YxGbQoD) This works the way I need it to, for the most part; however, there's virtually no way to link a 'wins' column with the guild id. I only want to showcase the user wins that are *in* the guild on the leaderboard.
## [8][Elastic Beanstalk introduces support for shared load balancers!](https://www.reddit.com/r/aws/comments/iqjxcw/elastic_beanstalk_introduces_support_for_shared/)
- url: https://aws.amazon.com/blogs/containers/amazon-elastic-beanstalk-introduces-support-shared-load-balancers/
---

## [9][Suggestions on reducing application startup time?](https://www.reddit.com/r/aws/comments/iquord/suggestions_on_reducing_application_startup_time/)
- url: https://www.reddit.com/r/aws/comments/iquord/suggestions_on_reducing_application_startup_time/
---
I am working on an interactive application that currently requires one Windows g4dn.xlarge instance per active user (used to be a desktop app). At the moment it uses Auto Scaling Groups (ASG) and increases the desired capacity by 1 when a user logs in. 

The problem is that it is slow to start - \~5 minutes to launch the instance and another 5-10 minutes for the initial application start (which takes 1-2 minutes locally on a freshly booted machine). I understand it is slow because when using snapshots, each block is initially loaded from S3.

&amp;#x200B;

I would like to reduce startup times while keeping costs down. Is that possible? I can think of several alternatives:

1. Keep a few instances on (desired capacity &gt; 0) and increase/decrease when users log in/out. Application start time will be fast, but if I set capacity to 5 and 6 users log in at once the 6th one will face a slow starting time. At $0.71/hr per instance it is also pretty expensive.

2. Use EBS fast snapshot restore with desired capacity at 0. This should have been pretty fast, but I am getting inconsistent results, even with subsequent launches in the same AZ. I am seeing about 1-3 minutes to launch the instance and another 1-5 minutes for initial startup.

3. Avoid ASG and keep a few stopped machines in each AZ (to account for capacity issues). This should be fast and consistent, the main issue is that I need to implement ASG features. Standby cost would be $0.1/GB/hr, so not so expensive with 100 GB machines.

&amp;#x200B;

It looks like #3 would be the best option if considering both time and cost, am I missing something? It does require more dev work. I am also working on supporting multiple users per instance, switching to Linux and reducing load times, but that will take a while.
## [10][Tiny library for generating safe DynamoDB sort-keys in JavaScript](https://www.reddit.com/r/aws/comments/iqtn5x/tiny_library_for_generating_safe_dynamodb/)
- url: https://github.com/neuledge/sort-key
---

