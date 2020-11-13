# aws
## [1][re:Invent registration is now open](https://www.reddit.com/r/aws/comments/jkenu3/reinvent_registration_is_now_open/)
- url: https://register.virtual.awsevents.com/
---

## [2][Meet the newest AWS Heroes including the first DevTools Heroes!](https://www.reddit.com/r/aws/comments/jt1o78/meet_the_newest_aws_heroes_including_the_first/)
- url: https://aws.amazon.com/blogs/aws/meet-the-newest-aws-heroes-including-the-first-devtools-heroes/
---

## [3][Improving amazon athena performance](https://www.reddit.com/r/aws/comments/jtg5gk/improving_amazon_athena_performance/)
- url: https://www.reddit.com/r/aws/comments/jtg5gk/improving_amazon_athena_performance/
---
How would you improve your athena queries? I have a 10GB of raw data that will be fed to aws athena and I've noticed that the queries have been taking too long. To the point where I've experienced it timing out after 30 mins. Is there any way around this?
## [4][How to paginate Aws s3 list object?](https://www.reddit.com/r/aws/comments/jtd7jn/how_to_paginate_aws_s3_list_object/)
- url: https://www.reddit.com/r/aws/comments/jtd7jn/how_to_paginate_aws_s3_list_object/
---
I have over 20000 images in s3 and I want to paginate the first 100 after clicking on pagination 2nd link it should load the second 100 images and so on.

  

`const params = { Bucket: "test-bucket", Delimiter: '/', MaxKeys: 100, Prefix: "thumbnail_images/Q" };`

`async function* listAllKeys(params) {`  
  `try {`  
    `do {`  
      `const data = await s3.listObjectsV2(params).promise();`  
      `params.ContinuationToken = data.NextContinuationToken;`  
      `yield data;`  
    `} while (params.ContinuationToken);`  
  `} catch (err) {`  
    `return err;`  
  `}`  
`};`

I am using **aws-sdk** node package.
## [5][Greengrass Stream Manager &amp; Kinesis Data Stream cooperation](https://www.reddit.com/r/aws/comments/jte569/greengrass_stream_manager_kinesis_data_stream/)
- url: https://www.reddit.com/r/aws/comments/jte569/greengrass_stream_manager_kinesis_data_stream/
---
Hi,

I've been working with Stream Manager synchronising my local streams to Kinesis for a while.  
There's one thing that bothers me a lot..

Well, let me use an example.  
Let's say my local stream batch size is 100 and each message has 100kB. When local message buffer reaches 100, data gets synchronised with remote Kinesis stream by Stream Manager. 

1) Do all 100 messages get added to the Kinesis Data Stream at same time? So it gives 100\*100kB = 10MB? 

2) How is data being read by shards? If I put those 10MB at once, does it mean it will take 10 seconds to read for 1 shard or 1 second to read for 10 shards?

3) I also faced RateExceeded exception (found in SM logs on the device). However, I can't understand how did it happen. My Data Stream is tied to Delivery Stream which puts data to S3. And.. I found some duplicated data there once and metrics for the Data Stream went crazy - they almost lined up with limits, even though I was sending much less than 1MB/sec (with 1 shard). Well, much less when we calculate this over time, but more than 1MB at one write... So here we kinda switch back to the 2nd question.   

 I'd like to estimate shards amount using calculator. I set "Average record size (in KiB)" to 100kB (I believe it's about a single message out of those 100, isn't it?). But I'm not sure what to set for "Maximum records written per second" - would it be always same as batch size, so 100? (assuming this is the only data source for Kinesis Data Stream).

Thanks in advance!
## [6][Can we connect to AWS RDS Standby DB when Primary is Active and running ?](https://www.reddit.com/r/aws/comments/jtdwqx/can_we_connect_to_aws_rds_standby_db_when_primary/)
- url: https://www.reddit.com/r/aws/comments/jtdwqx/can_we_connect_to_aws_rds_standby_db_when_primary/
---
Can we connect to RDS Standby DB when Primary is Active and running ?
## [7][How do I create the user authentication ( sign in )using Cognito which will give access to existing users in RDS database?](https://www.reddit.com/r/aws/comments/jtb11t/how_do_i_create_the_user_authentication_sign_in/)
- url: https://www.reddit.com/r/aws/comments/jtb11t/how_do_i_create_the_user_authentication_sign_in/
---
I have imported data in MySQL DB, it has thousands of entries for users. I need to give authentication to the users present in the MySQL db table using AWS Cognito. Can anyone please explain to me is it possible, if so how to implement the solution?
## [8][Using AWS Lambda extensions to send logs to custom destinations | Amazon Web Services](https://www.reddit.com/r/aws/comments/jt00uu/using_aws_lambda_extensions_to_send_logs_to/)
- url: https://aws.amazon.com/blogs/compute/using-aws-lambda-extensions-to-send-logs-to-custom-destinations/
---

## [9][Majority of Alexa Now Running on Faster, More Cost-Effective Amazon EC2 Inf1 Instances](https://www.reddit.com/r/aws/comments/jszq36/majority_of_alexa_now_running_on_faster_more/)
- url: https://aws.amazon.com/blogs/aws/majority-of-alexa-now-running-on-faster-more-cost-effective-amazon-ec2-inf1-instances/
---

## [10][How to implement multiple Lambda REST APIs with single code base](https://www.reddit.com/r/aws/comments/jtbi7p/how_to_implement_multiple_lambda_rest_apis_with/)
- url: https://www.reddit.com/r/aws/comments/jtbi7p/how_to_implement_multiple_lambda_rest_apis_with/
---
Hello! I’m trying to find the best way implement multiple instances of the same REST API code (C#) via Lambda, in which each API would connect to a different database.  Being completely new to AWS, I could see the scalability could easily get out of control and Im really open to any guidance or recommendations regarding this.


The idea is that each customer (company) would have their own respective API that’s locked down that only their users can access and is accessed via mobile and desktop clients. The API would then connect to its own individual RDS database. The API code can either connect to its respective database by a stored connection file, or I’ve experimented with using Lambda variables that could be set per Lambda function to manage the database connection.


I’ve seen from uploading via Visual Studio (C# / AWS Plugins) that it automatically creates a Cloud Formation stack that seems to contain the following:
-ServerlessRestApi
-AspNetCoreFunction


Ive also noticed it creates an S3 bucket where it stores the code in a zip file within an AwsApi/ subfolder along with a template file


Optimally I’m wondering if it’s a matter of making multiple lambda functions reference the same S3 bucket (and not sure how Id do that), so we just update code deployments in that one location? Or is there a way to have a single REST API (In Lambda or otherwise) point to a different database per connection request using their keys, stages, or other variables somehow?


Thank you!
## [11][Running Postgres raw query with Glue Connection](https://www.reddit.com/r/aws/comments/jta0xf/running_postgres_raw_query_with_glue_connection/)
- url: https://www.reddit.com/r/aws/comments/jta0xf/running_postgres_raw_query_with_glue_connection/
---
I want to run query over a Postgres db from my glue job and fetch the results into a dataframe. I'm able to connect to the database, even able to extract a db table into dataframe. But, when I try to run a raw query by passing &lt;query&gt; parameter in the connection\_options, I am receiving this error:

{..........An error occurred while calling o71.getDynamicFrame. : java.util.NoSuchElementException: key not found: location at scala.collection.MapLike$class.default(MapLike.scala:228).........}

I am able to achieve similar thing with redshift, I just pass &lt;query&gt; in my connection\_options and it fetched the results into the DF, but same thing is not working with Postgres, even though the documentation seems to be show no difference. The piece of code that I'm using to achieve this:

query = "SELECT \* FROM schema.table\_name where &lt;condition&gt; limit X"

connection\_options = {  
"url": "jdbc:postgresql://xxxxxx.yyyyy.rds.amazonaws.com:&lt;port&gt;/dbname", "query": query, "user": &lt;username&gt;, "password": &lt;password&gt; }

df1 = glueContext.create\_dynamic\_frame.from\_options(connection\_type="postgresql",connection\_options=connection\_options)

df1.show()

However, the following works fine:

connection\_options = {  
"url": "jdbc:postgresql://xxxxxx.yyyyy.rds.amazonaws.com:&lt;port&gt;/dbname", "dbtable": schema.table, "user": &lt;username&gt;, "password": &lt;password&gt;

}

df1 = glueContext.create\_dynamic\_frame.from\_options(connection\_type="postgresql",connection\_options=connection\_options)

df1.show()
