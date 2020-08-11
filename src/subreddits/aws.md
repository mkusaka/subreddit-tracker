# aws
## [1][Week of Aug 10th - What AWS questions do you have?](https://www.reddit.com/r/aws/comments/i75xfw/week_of_aug_10th_what_aws_questions_do_you_have/)
- url: https://www.reddit.com/r/aws/comments/i75xfw/week_of_aug_10th_what_aws_questions_do_you_have/
---
What questions do you have about AWS?
## [2][AWS Glue version 2.0 featuring 10x faster job start times and 1-minute minimum billing duration](https://www.reddit.com/r/aws/comments/i7eat0/aws_glue_version_20_featuring_10x_faster_job/)
- url: https://aws.amazon.com/blogs/aws/aws-glue-version-2-0-featuring-10x-faster-job-start-times-and-1-minute-minimum-billing-duration/
---

## [3][ELI5: Why shouldn't we all move to Graviton2 instances if they're 40%* cheaper? What's the catch?](https://www.reddit.com/r/aws/comments/i7jk3z/eli5_why_shouldnt_we_all_move_to_graviton2/)
- url: https://www.reddit.com/r/aws/comments/i7jk3z/eli5_why_shouldnt_we_all_move_to_graviton2/
---
\* that's what their [marketing](https://aws.amazon.com/ec2/instance-types/m6/) claims 

Beginner here - not familiar with CPU architectures and OS stuff, I've only used Ubuntu Server on EC2. Curious about what it'd take for everyone to move to ARM-based instances if they're cheaper? For instance, if I'm running a Docker Flask app on M5/T3 already, what exactly stops me from moving it to M6g?
## [4][Would you like more polls?](https://www.reddit.com/r/aws/comments/i7rnnu/would_you_like_more_polls/)
- url: https://www.reddit.com/r/aws/comments/i7rnnu/would_you_like_more_polls/
---
I was thinking of creating some polls for the subreddit. Ex: Desired features/Biggest Challenges/Etc Please indicate your interest below. 

[View Poll](https://www.reddit.com/poll/i7rnnu)
## [5][A little Help with Stitching Two Clips Together with E-Transcoder CLI ?](https://www.reddit.com/r/aws/comments/i7qep9/a_little_help_with_stitching_two_clips_together/)
- url: https://www.reddit.com/r/aws/comments/i7qep9/a_little_help_with_stitching_two_clips_together/
---
Hello,

I am trying to Stitch two clips together with Elastic Transcoder CLI,

Command : 

    aws elastictranscoder create-job --pipeline-id 1581555518949-p2e73l --inputs --input  Key=intro.mp4 --input Key=video.mp4 --outputs Key=New-video.mp4,PresetId=1594777774704-fupili

This one is not working, I would be really grateful if you can point me into the right direction :)  


Best Regards.
## [6][⚡️ Dynatron - Bridge between AWS DynamoDB Document Client and Real World usage](https://www.reddit.com/r/aws/comments/i7cz0e/dynatron_bridge_between_aws_dynamodb_document/)
- url: https://www.reddit.com/r/aws/comments/i7cz0e/dynatron_bridge_between_aws_dynamodb_document/
---
This library is a result of years of working with AWS DynamoDB and overcoming underwater rocks, missing optimizations and hidden issues that are very hard to catch (like hanging SSL connections in 0.2% of cases).

Homepage - [https://93v.github.io/dynatron/](https://93v.github.io/dynatron/)

Github - [https://github.com/93v/dynatron](https://github.com/93v/dynatron)

NPM Package - [https://www.npmjs.com/package/dynatron](https://www.npmjs.com/package/dynatron)
## [7][Where are photos saved by default in AWS Batch?](https://www.reddit.com/r/aws/comments/i7pf52/where_are_photos_saved_by_default_in_aws_batch/)
- url: https://www.reddit.com/r/aws/comments/i7pf52/where_are_photos_saved_by_default_in_aws_batch/
---
First of all, I'm completely new to AWS and am confused a lot so please explain like I'm a 14-15-year-old kid. 

So, I created my first Batch job with the following dockerfile:

    FROM python
    RUN pip3 install matplotlib
    RUN pip3 install seaborn
    RUN mkdir /src
    COPY . /src
    CMD ["python3", "/src/run.py"]

This python script generates only one graph in the form of a JPEG image and is saved using the command

    fig.savefig('graph.jpeg', format='jpeg', dpi=100, bbox_inches='tight')

Then the job starts and runs successfully. But now I can't find where the image got saved because the job didn't fail. Where are images saved? And do I have to start a new service to save photos?
## [8][Need help with using EC2 server to host Discord bot to control my smart lights.](https://www.reddit.com/r/aws/comments/i7lyr6/need_help_with_using_ec2_server_to_host_discord/)
- url: https://www.reddit.com/r/aws/comments/i7lyr6/need_help_with_using_ec2_server_to_host_discord/
---
Let me start off by saying I have literally never used AWS before. So basically I made a discord bot in Javascript so I can control my Philips Hue lights in my room, it works perfectly, so I set it up in an EC2 instance so it would always be running. Now the problem is, none of the commands work because the AWS server cannot reach my Hue bridge's IP, which I'm assuming only works on my IP. So my question is, how do I make this work? Sorry if this was confusing.
## [9][AWS TechU program](https://www.reddit.com/r/aws/comments/i78w38/aws_techu_program/)
- url: https://www.reddit.com/r/aws/comments/i78w38/aws_techu_program/
---
Hello,

Has anyone been interviewed for the TechU program for Solution architect in AWS? I wanted to know about the rounds and some général details about the interview process. Thanks in advance.
## [10][Anyone having connection problems with Aurora serverless in the last few days in Ireland?](https://www.reddit.com/r/aws/comments/i7o7ik/anyone_having_connection_problems_with_aurora/)
- url: https://www.reddit.com/r/aws/comments/i7o7ik/anyone_having_connection_problems_with_aurora/
---
Hi, 

we are running an aurora serverless in Ireland, in the last 7 seven days we had two incident: suddenly for a few minutes we were unable to login then everything went back to normal, no loss of data only angry clients ;-)

I'm waiting for the findings of AWS support, but I'm wondering if others have experienced the same problem.
## [11][User Groups and Reddit - Let's Come Together!](https://www.reddit.com/r/aws/comments/i779oc/user_groups_and_reddit_lets_come_together/)
- url: https://www.reddit.com/r/aws/comments/i779oc/user_groups_and_reddit_lets_come_together/
---
No idea if this is the best way to go about this, but I figured it was worth asking the community...

I'm one of the organizers of the Greater Philadelphia AWS User Group ([meetup.com/gpawsug](https://meetup.com/gpawsug)).  We try our best to bring informative and interesting content to our users in the Philly area every month.  We do this not for profit, but rather to spread the word about the platform and hopefully win some hearts and minds.  We shy away from 3rd party salesy content, always looking for ways that we can solve problems with the AWS platform alone.  We recently passed 1000 users, so it seems like we must be doing something right.

Most of the user groups are free, and are not directly affiliated with Amazon or AWS.  We try to stay firmly planted in that "USER" segment.

Lately, my co-admins and I have been talking about how we could use Reddit to help us make our community better.  We definitely don't want to throw spammy advertising posts for our meetups (Now virtual! Thanks, COVID!).  But I know there is some way we can bring goodness to redditors and our community alike.

So finally, my points...

* Anyone interested in presenting a topic to our group?  If so, what?  (Note:  we try to cater to all levels of experience, across all topics in the platform)
* Anyone have a good idea for how we can bring User Groups and subreddits closer together? We have a global community of user groups who would certainly be interested in this answer!
* Go register for your local user group and get involved!  There's a directory on the AWS Developer site.  [https://aws.amazon.com/developer/community/usergroups/](https://aws.amazon.com/developer/community/usergroups/)
