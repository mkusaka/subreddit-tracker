# aws
## [1][re:Invent registration is now open](https://www.reddit.com/r/aws/comments/jkenu3/reinvent_registration_is_now_open/)
- url: https://register.virtual.awsevents.com/
---

## [2][Amazon CloudWatch launches tag based Metrics Explorer](https://www.reddit.com/r/aws/comments/joc21y/amazon_cloudwatch_launches_tag_based_metrics/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/11/amazon-cloudwatch-launches-metrics-explorer/
---

## [3][I built an open source version of the extension that changes the color of the navbar depending on the region](https://www.reddit.com/r/aws/comments/jnyolc/i_built_an_open_source_version_of_the_extension/)
- url: https://github.com/corollari/aws-color-region-navbar-extension
---

## [4][Use case question: I am accepting datasets of zipped images from users to be uploaded to s3 via the web browser, am I doing this right?](https://www.reddit.com/r/aws/comments/joi4e9/use_case_question_i_am_accepting_datasets_of/)
- url: https://www.reddit.com/r/aws/comments/joi4e9/use_case_question_i_am_accepting_datasets_of/
---
Hi /r/aws,

My use case is I am trying to accept large zips of data, sometimes 10+ GB, from the web browser to be uploaded to S3 under the users UUID (so /fjhsdjfds-432fase-fsdy7sd8-fayfasdiuf/projects/datasets/catpics.zip)

&amp;#x200B;

We later run some analysis on these zips downstream, however...there may be a time where they want to be able to browse the images in their datasets from the browser...in my mind the only way to enable this would be to unzip the files and store them in S3. This will obviously increase our storage costs by a ton, but we could charge accordingly for users that want this...

&amp;#x200B;

In any case, am I right to think the only way to handle this is to unzip the file immediately and store everything unzipped? Or is there something I'm missing?

&amp;#x200B;

Thanks for the help.
## [5][CloudFront Whitelisted Headers are returning "{"message":"Forbidden"}"](https://www.reddit.com/r/aws/comments/johf3g/cloudfront_whitelisted_headers_are_returning/)
- url: https://www.reddit.com/r/aws/comments/johf3g/cloudfront_whitelisted_headers_are_returning/
---
I Whitelisted the following header for CloudFront and everything works well.

https://preview.redd.it/i6v4aglxqex51.png?width=1812&amp;format=png&amp;auto=webp&amp;s=b22050e1c0fc145aec9bcbc92b30a539511a8b47

I then added Host and Origin. I noticed that now CloudFront responds with:

"{"message":"Forbidden"}"

https://preview.redd.it/vbd09fthrex51.png?width=2156&amp;format=png&amp;auto=webp&amp;s=7ae64aab317af6ea534f133bbe2d0662ed848a19

I then added removed host and it started working again:

&amp;#x200B;

https://preview.redd.it/snyc2r3vsex51.png?width=2024&amp;format=png&amp;auto=webp&amp;s=28dc2f8fc71ba2a423ea4033115422673b6d52a4

&amp;#x200B;

&amp;#x200B;

Why would this happen? My intuition tells me that if Host is not in the header, then cloudfront will fail.

Why do I think this? Because I as soon as I removed Host, the CloudFront worked again.

&amp;#x200B;

Looking to understand what is going on here. Thank you in advance.
## [6][Using Service Meshes in AWS](https://www.reddit.com/r/aws/comments/joh5fs/using_service_meshes_in_aws/)
- url: https://d1.awsstatic.com/whitepapers/using-service-meshes-in-aws.pdf?did=wp_card&amp;trk=wp_card
---

## [7][Amazon MQ Update – New RabbitMQ Message Broker Service](https://www.reddit.com/r/aws/comments/jo2esj/amazon_mq_update_new_rabbitmq_message_broker/)
- url: https://aws.amazon.com/blogs/aws/amazon-mq-update-new-rabbitmq-message-broker-service/
---

## [8][Amazon Kinesis Data Streams enables data stream retention up to one year](https://www.reddit.com/r/aws/comments/jobeyq/amazon_kinesis_data_streams_enables_data_stream/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/11/amazon-kinesis-data-streams-enables-data-stream-retention/
---

## [9][Is it possible to pass arguments to a glue workflow?](https://www.reddit.com/r/aws/comments/joijky/is_it_possible_to_pass_arguments_to_a_glue/)
- url: https://www.reddit.com/r/aws/comments/joijky/is_it_possible_to_pass_arguments_to_a_glue/
---
I'm trying to replace some of my glue jobs and crawlers with a workflow. These jobs currently take a date as an argument. If I invoke the workflow in node.js through the following  


`let params = {`  
  `Name: name,`  
  `Arguments: {`  
`selected_date: date`  
  `}`  
`}`

`Glue.startWorkflowRun(params);`

I get an error that Arguments is a unexpected parameter. I've also tried Properties. It's weird because I know that its possible to use default properties in a workflow and to read those via the jobs, but why then would it not be possible to pass properties in?
## [10][How can I Route users to the correct paths based on the domain the URL they type in?](https://www.reddit.com/r/aws/comments/joi3vr/how_can_i_route_users_to_the_correct_paths_based/)
- url: https://www.reddit.com/r/aws/comments/joi3vr/how_can_i_route_users_to_the_correct_paths_based/
---
Hi - I purchased a domain through route53. I connected it to CloudFront. I am struggling to map the URL 

&amp;#x200B;

Basically, if somebody types in "example dot com" without any slashes, I Want them to be routed to another website.

"example dot com/" -&gt; send person to "example2 dot com"

If they type in:

"example dot com/order-status" -&gt; send the person through to CloudFront, which will route them to API Gateway.
## [11][A New Integration for CloudWatch Alarms and OpsCenter](https://www.reddit.com/r/aws/comments/jo8ypl/a_new_integration_for_cloudwatch_alarms_and/)
- url: https://aws.amazon.com/blogs/aws/a-new-integration-for-cloudwatch-alarms-and-opscenter/
---

