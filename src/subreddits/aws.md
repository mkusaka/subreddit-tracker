# aws
## [1][Amazon Athena adds support for running SQL queries across relational, non-relational, object, and custom data sources.](https://www.reddit.com/r/aws/comments/jtxnwl/amazon_athena_adds_support_for_running_sql/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/11/aws-what-s-new-for-athena-federated-query/
---

## [2][Lightsail Containers: An Easy Way to Run your Containers in the Cloud](https://www.reddit.com/r/aws/comments/jtk41e/lightsail_containers_an_easy_way_to_run_your/)
- url: https://aws.amazon.com/blogs/aws/lightsail-containers-an-easy-way-to-run-your-containers-in-the-cloud/
---

## [3][Amazon Textract now supports handwriting.](https://www.reddit.com/r/aws/comments/jtn8xf/amazon_textract_now_supports_handwriting/)
- url: https://aws.amazon.com/blogs/machine-learning/amazon-textract-recognizes-handwriting-and-adds-five-new-languages/
---

## [4][How to implement a system that process files in order of arrival with AWS services?](https://www.reddit.com/r/aws/comments/ju148z/how_to_implement_a_system_that_process_files_in/)
- url: https://www.reddit.com/r/aws/comments/ju148z/how_to_implement_a_system_that_process_files_in/
---
I have files incoming in several S3 buckets.

Files in different buckets can be process in parallel, they are independent.

However files in the same bucket must be processed in strict order of arrival and CANNOT be processed in parallel.

For processing, lambda functions could be used.

I'm not sure however on how to orchestrate the whole thing.

The ordering of processing is key.

I was thinking on something like SQS, that preserve orders but I'm not sure how to avoid parallel processing in the same bucket.

Maybe also step functions?

I also would like a solution as serverless as possible, as I expect the system to be able to scale automatically.

Files are coming in every 2-3 minutes in each bucket, and processing it's usually faster than that, and mostly under 5 minutes. Files are not big, max 5-10MB.
## [5][Session Revoke after Completing the Exam](https://www.reddit.com/r/aws/comments/jtwz1b/session_revoke_after_completing_the_exam/)
- url: https://www.reddit.com/r/aws/comments/jtwz1b/session_revoke_after_completing_the_exam/
---
 So I just finished taking the AWS SAA Exam with Pearson VUE and PASSED. But I needed to pee during the exam. I told the examiners bf he said I wasn’t allowed to leave. I held it for about an hour then I finished the exam. Couldn’t hold it in any longer and didn’t bother to review my flagged questions. Clicked next next next and what do I know? Freaking survey questions. I was so annoyed and couldn’t hold it in anymore and half way through answering the survey questions I told the examiner I’m going to pee and left. When I got back he told me he was revoking my session. I clicked end exam and saw that I passed. What’s gonna happen to me now?
## [6][AWS Glue to RDS help](https://www.reddit.com/r/aws/comments/jtzn9x/aws_glue_to_rds_help/)
- url: https://www.reddit.com/r/aws/comments/jtzn9x/aws_glue_to_rds_help/
---
I have created an AWS Glue connection to my RDS database from my S3 bucket, but every time I run it, it gives the same error: Check that your connection definition references your JDBC database with correct URL syntax, username, and password. Unknown database '**DB identifier**'

I think my problem is that I don't know what my Database name is. All I have done in RDS so far is created my database, and I have my database and my database instance. Is the database name something that I create in Glue or is it something else?
## [7][Best practice when making a lambda function (with Gateway) to read and update a dynamodb table?](https://www.reddit.com/r/aws/comments/jtwyr2/best_practice_when_making_a_lambda_function_with/)
- url: https://www.reddit.com/r/aws/comments/jtwyr2/best_practice_when_making_a_lambda_function_with/
---
So, basically, I am doing a very simple project to learn lambda/gateway/dynamodb.

Basically, I will have one single item in the dynamodb table. The item in the table is a number. Whenever the lambda function is triggered by the gateway, I want to update the count by 1 and then get that data and send it back via the gateway.

I have been reading up up lambda and built a simple function that can get_item from the table successfully, built it using python.

Next, I plan on writing code to simply update that item by the count of one.

My thought process to do this is to simply (in the lambda_handler function), first put the code that updates the count in the table. Then, right after that code, put the code that gets the count.

I want to make sure though that the get is getting the NEW count after the first part of the function is ran. 

Would doing the above be the best practice for this? Or would some weird async stuff be going on where, even though the update part of the code is ran first, the function would still get the old data? I guess this would also be potentially an issue with dynamodb because it is "eventually" consistent I heard?

Just curious on what the best practice is with the above? Thanks for any help.
## [8][SES - sends from me to me on a gmail business account flagged as possible fishing](https://www.reddit.com/r/aws/comments/jtv9ae/ses_sends_from_me_to_me_on_a_gmail_business/)
- url: https://www.reddit.com/r/aws/comments/jtv9ae/ses_sends_from_me_to_me_on_a_gmail_business/
---
I have been replacing all of our metric-alert endpoints to be a lambda as terraform does not support email b/c the ARN is generated at the time of sending.  I don't thinks to relevant.  Not so important, but each of these emails gets that big yellow bar, and i click "Looks Safe".  My email is verified, every security measure is passing except for DMARC at the moment, but i think this is more about the text, which is formatted json wrapped in a &lt;pre&gt; tag.  Is there a setting, or 8, that i can do to get gmail to believe it.  


TBH- Sorry this is really not completely an SES question, but i feel like people may have ran into this here.  If its a gmail switch, i can have my boss change it monday.   


Thanks
## [9][Improving amazon athena performance](https://www.reddit.com/r/aws/comments/jtg5gk/improving_amazon_athena_performance/)
- url: https://www.reddit.com/r/aws/comments/jtg5gk/improving_amazon_athena_performance/
---
How would you improve your athena queries? I have a 10GB of raw data that will be fed to aws athena and I've noticed that the queries have been taking too long. To the point where I've experienced it timing out after 30 mins. Is there any way around this?
## [10][Securely accessing RDS database from desktop GUI: any pitfalls?](https://www.reddit.com/r/aws/comments/jtjlo9/securely_accessing_rds_database_from_desktop_gui/)
- url: https://www.reddit.com/r/aws/comments/jtjlo9/securely_accessing_rds_database_from_desktop_gui/
---
Hi folks,

hopefully I dont ask something something that has been asked (I tried finding a clear answer, but maybe my search skills are rusty). 

&amp;#x200B;

I am migrating my database (one previously hosted at Digital Ocean) to a Postgres database at RDS, one which occasionally I need to manually access (in case some migrations fail). I personally use TablePlus to do some easy operations.

&amp;#x200B;

Now I am aware that to achieve this, I would have to enable that my database can be accessed "Publically". I am a bit wary of doing so, as I worry that I do not fully comprehend whether this may expose me to any potential dangers. 

&amp;#x200B;

I assume that I would want to create a "whitelist" in my VPC, which allow only specific IPs to access this "public Database"? Or are there better, more secure ways of doing so? Any particular pitfalls one needs to way off when doing so?
