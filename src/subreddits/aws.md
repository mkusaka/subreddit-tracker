# aws
## [1][AWS: Your complete guide to Amazon Web Services &amp; features](https://www.reddit.com/r/aws/comments/ey19oe/aws_your_complete_guide_to_amazon_web_services/)
- url: https://www.techradar.com/news/aws
---

## [2][Simple Cognito example to demonstrate sign-up, login and access S3](https://www.reddit.com/r/aws/comments/ey30il/simple_cognito_example_to_demonstrate_signup/)
- url: https://github.com/franzwong/cognito-s3-cloudformation-example
---

## [3][Understanding billing question](https://www.reddit.com/r/aws/comments/exya4z/understanding_billing_question/)
- url: https://www.reddit.com/r/aws/comments/exya4z/understanding_billing_question/
---
I've had a ~350 GB in IA storage for a few years. In early December, I moved it all over to Glacier Deep Archive. The data is all just offsite backup of a file server, and I hope to never have to access it.

The December bill was higher, which I expected since I paid for all the class transfers.

I just got my January bill, and it lists charges for about 350 GB in *each* of IA and 
Glacier Deep Archive. I can't figure out why. I logged into CloudWatch Management and it shows that I suddenly had 350 GB in Deep Archive starting in early December, and the amount of data in IA never went down.

Can anyone clue me in to what's going on?
## [4][Native Lambda](https://www.reddit.com/r/aws/comments/ey4c0m/native_lambda/)
- url: https://www.reddit.com/r/aws/comments/ey4c0m/native_lambda/
---
Gotta share this :)

Just deployed a native Linux binary function written in Java, using Micronaut and compiled with GraalVM to Lambda. This has a call latency of ~60ms :)

Is there anoyone else using Micronaut and Graal?
## [5][Is there a point in using crawlers on AWS Glue if I only use Python Shell and not PySpark?](https://www.reddit.com/r/aws/comments/ey264b/is_there_a_point_in_using_crawlers_on_aws_glue_if/)
- url: https://www.reddit.com/r/aws/comments/ey264b/is_there_a_point_in_using_crawlers_on_aws_glue_if/
---
Long story short, my company decided to use Python Shell instead of PySpark on Glue due to cost/benefit reasons. My understanding is that I'd be using boto3 to retrieve data directly from s3 client, instead of going through the trouble of setting up glue context and DynamicFrame. 

Is there any reason to use a Glue crawler/classifier when I can just define everything in my Python Shell job? If there is, are there any examples of how to utilize Glue catalog using Python Shell? I can't find examples, and [the documentation](https://boto3.amazonaws.com/v1/documentation/api/latest/reference/services/glue.html) doesn't give me a picture on why should and how could I use it.
## [6][The Lambda Layer Inspector (L2I)](https://www.reddit.com/r/aws/comments/ey5k1q/the_lambda_layer_inspector_l2i/)
- url: https://github.com/mhausenblas/l2i
---

## [7][New to AWS need some help](https://www.reddit.com/r/aws/comments/ey4izj/new_to_aws_need_some_help/)
- url: https://www.reddit.com/r/aws/comments/ey4izj/new_to_aws_need_some_help/
---
A friend of mine who uses AWS asked me a technical question regarding a particular feature which I have mentioned below. But I'm only familiar with CLI (I use AWS as a means to run some simulations) and this is related to DynamoDb. hope some of you whom are more experienced in  deploying DB can help me out in this regard.

"How to use AWS lambda to connect an app to a database server(say AWS DynamoDb) through API gateway without making it publicly accessible"
## [8][What's the difference between SSD and NVMe?](https://www.reddit.com/r/aws/comments/expdmc/whats_the_difference_between_ssd_and_nvme/)
- url: https://www.reddit.com/r/aws/comments/expdmc/whats_the_difference_between_ssd_and_nvme/
---
I am probably not making the comparison correctly, but my thinking is that SSD EBS is the same as an SSD drive we'd have on our own computer.  If I'm correct, then what exactly is NVMe?  Just faster SSD?  Or maybe it's directly mounted to the instance and so it's exponentially faster?
## [9][Postgres on Rds - horrible checkpoint behaviour](https://www.reddit.com/r/aws/comments/exsy3i/postgres_on_rds_horrible_checkpoint_behaviour/)
- url: https://www.reddit.com/r/aws/comments/exsy3i/postgres_on_rds_horrible_checkpoint_behaviour/
---
Hi,

Were bulk loading into a postgres db on rds and getting horrible behaviour where checkpoints take AGES and occur repeatedly.

Our data isn't even that big. Just a few dozen tables some with up to 20million rows. 

We played with the wal size but it just moves the problem rather than fixing it.

Never seen this in pg before so is there something out of whack in the default rds config?

Thanks!
Dan
## [10][Scoring A+ for SSL Labs on My Cloudfront-Hosted Static Website](https://www.reddit.com/r/aws/comments/exmsg5/scoring_a_for_ssl_labs_on_my_cloudfronthosted/)
- url: https://adamj.eu/tech/2020/02/02/scoring-a+-for-ssl-labs-on-my-cloudfront-hosted-static-site/
---

