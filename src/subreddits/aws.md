# aws
## [1][Live Stream Today: Building Serverless Fullstack Applications (AWS Lambda, Express, React &amp; More)](https://www.reddit.com/r/aws/comments/ge2knj/live_stream_today_building_serverless_fullstack/)
- url: https://www.reddit.com/r/aws/comments/ge2knj/live_stream_today_building_serverless_fullstack/
---
Hi all,

I'm doing a live stream w/ AWS today on building serverless fullstack applications today at 2pm PT if you'd like to join!

* **Stream:** [https://www.twitch.tv/aws](https://www.twitch.tv/aws)
* **Project:** [https://github.com/serverless-components/fullstack-app](https://github.com/serverless-components/fullstack-app)
## [2][Building AWS environment documentation](https://www.reddit.com/r/aws/comments/geiuhv/building_aws_environment_documentation/)
- url: https://www.reddit.com/r/aws/comments/geiuhv/building_aws_environment_documentation/
---
Does anyone have a nice way of building some sort of documentation of a new environment? Like you walked into a new place and there was no documentation type of deal? Basically looking for a starting point vs reinventing the wheel here.
## [3][AWS Lambda - Chalice/Zappa equivalents in other languages?](https://www.reddit.com/r/aws/comments/gek11y/aws_lambda_chalicezappa_equivalents_in_other/)
- url: https://www.reddit.com/r/aws/comments/gek11y/aws_lambda_chalicezappa_equivalents_in_other/
---
Hi.

Are there any equivalents of Zappa or Chalice frameworks in other languages, namely TypeScript or Java?




(Zappa/Chalice are Python web app building frameworks. Instead of multiple Lambda functions, you only have one monolithic Lambda, which also contains the routing part inside. You target your "one-to-rule-them-all" function with ALB or API Gateway and voila. If you swap `zappa` for `flask` in your app and transform it into container, you can quickly make it run on ECS. This is useful if you'd like for example to start your project on Lambdas then easily migrate to containers when costs are too high/traffic patterns of your customers are known)
## [4][How and when to use IAM for new services?](https://www.reddit.com/r/aws/comments/ge7qka/how_and_when_to_use_iam_for_new_services/)
- url: https://www.reddit.com/r/aws/comments/ge7qka/how_and_when_to_use_iam_for_new_services/
---
This is a dumb question, but I'm asking because I don't feel confident in how I should be using IAM. I feel like I understand the technical nature of IAM, just not the practical implementation of it in my day to day use of AWS.

Today I wanted to launch a third party cloud formation script. It created a lambda function, sns topic, and other small things.

I feel like I should be using IAM to create a user with a minimum policy, and launch this script under that user. I'm not sure if this is really a practical use of IAM.

I am curious how people typically use IAM/users when launching new services. Are there any general guidelines I should follow for "this is a situation I need a new IAM user/policy for"?
## [5][New AL2 Elastic Beanstalk python environments use pipenv?](https://www.reddit.com/r/aws/comments/ge6y8w/new_al2_elastic_beanstalk_python_environments_use/)
- url: https://www.reddit.com/r/aws/comments/ge6y8w/new_al2_elastic_beanstalk_python_environments_use/
---
Just noticed this in the release notes: [https://i.imgur.com/1laNVXK.png](https://i.imgur.com/1laNVXK.png)

I haven't really used pipenv, but from the occasional headlines/chatter I see, it sounds like a shitshow (and possibly dead?). So it's pretty surprising to find that the next gen EB platform is using it over pip (though I assume/hope pip is still available).

I couldn't find any discussion on this. What are everyone's thoughts on this?
## [6][How can I know what is the size of /tmp for MySQL in an RDS instance?](https://www.reddit.com/r/aws/comments/ge1jdc/how_can_i_know_what_is_the_size_of_tmp_for_mysql/)
- url: https://www.reddit.com/r/aws/comments/ge1jdc/how_can_i_know_what_is_the_size_of_tmp_for_mysql/
---
Hello,

We are trying to do a big query in one of our RDS instances, but we keep getting errors regarding a table being locked. I have searched high and low but seem to be hitting a wall (probably due to my knowledge or lack thereof). First, we had this error showing up in the logs:

    09:09:08 UTC - mysqld got signal 11 ;
    This could be because you hit a bug. It is also possible that this binary
    or one of the libraries it was linked against is corrupt, improperly built,
    or misconfigured. This error can also be caused by malfunctioning hardware.
    We will try our best to scrape up some info that will hopefully help
    diagnose the problem, but since we have already crashed,
    something is definitely wrong and this may fail.

Which only info I found was: [https://dba.stackexchange.com/questions/86915/mysqld-got-signal-11-its-my-fault-but-how-do-i-fix-it](https://dba.stackexchange.com/questions/86915/mysqld-got-signal-11-its-my-fault-but-how-do-i-fix-it) but, we are can't apply its fix.

After that, we thought the query (which is doing a massive temp table with all the selects/joins) might be causing the storage to fall (we had less than 2% free). We've now fixed that, along some optimization in the query, and see the following error:

    [ERROR] /rdsdbbin/mysql/bin/mysqld: Incorrect key file for table '/rdsdbdata/tmp/#sql_616d_2'; try to repair it

There's still plenty of free storage, so according to some other people facing the same problems, like for example [http://www.mysqlperformancetuning.com/a-fix-for-incorrect-key-file-for-table-mysql](http://www.mysqlperformancetuning.com/a-fix-for-incorrect-key-file-for-table-mysql) it's down to the free space of the /tmp dir.

Which brings me to my original question. How do I know what is the size of /tmp? Is it possible to increase? Do we need to raise a ticket with AWS?

Thanks!
## [7][Concern of running SQL statements over network vs importing a dump?](https://www.reddit.com/r/aws/comments/ge15k4/concern_of_running_sql_statements_over_network_vs/)
- url: https://www.reddit.com/r/aws/comments/ge15k4/concern_of_running_sql_statements_over_network_vs/
---
I have an RDS instance (t4 large, provisioned IOPs up to 2500) where I'll have an application connect to. This application will run Flyway migrations against the database. The initial setup, these migrations could be up to almost 8 or 9 million lines or sql over a spam of a couple days. Is there a concern that with so much sql going at a high rate that some things could get "lost" over the network? Would it be safer to create the database as a localhost on some server then dump/import into RDS?
## [8][How would you dynamically generate a gatsby site &amp; deploy it programmatically to s3?](https://www.reddit.com/r/aws/comments/ge85q1/how_would_you_dynamically_generate_a_gatsby_site/)
- url: https://www.reddit.com/r/aws/comments/ge85q1/how_would_you_dynamically_generate_a_gatsby_site/
---
Lets say I have a web app with a form, that form contains data that I want to use to generate a simple gatsby page. How would I generate the gatsby sites on AWS?  


Right now i'm thinking I could pass the data to a lambda, and generate the site there, then save all of the output to s3. I just can't wrap my head around how I would generate the site in a lambda function?
## [9][Stopping Cloudwatch Billing](https://www.reddit.com/r/aws/comments/ge3w57/stopping_cloudwatch_billing/)
- url: https://www.reddit.com/r/aws/comments/ge3w57/stopping_cloudwatch_billing/
---
Hello, I'm new to AWS, so I wasn't sure where to look. I created an endpoint from a sagemaker model that I then connected to the API gateway to do POST requests. I am not currently using the endpoint right now but will do so later, and in the meantime, Cloudwatch continues to log api requests, which results in me being billed. How can I stop Cloudwatch from logging these api requests?
## [10][cal on lightsail](https://www.reddit.com/r/aws/comments/gebf70/cal_on_lightsail/)
- url: https://www.reddit.com/r/aws/comments/gebf70/cal_on_lightsail/
---
greetings guys.   
it seems lightsail only lets 2 users to rdp in.we require 3 NON ADMIN users to rdp simultaneously. so i need 1 cal or 3 cal ?  
and where can i buy cal online?is it a permanent license or monthly or yearly ?   


thank you for your time !
