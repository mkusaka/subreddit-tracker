# aws
## [1][Infrastructure As Code Security](https://www.reddit.com/r/aws/comments/faao6f/infrastructure_as_code_security/)
- url: https://www.ibexlabs.com/infrastructure-as-code-security/
---

## [2][330 dollar Lesson learned [Glue]](https://www.reddit.com/r/aws/comments/fa3xbk/330_dollar_lesson_learned_glue/)
- url: https://www.reddit.com/r/aws/comments/fa3xbk/330_dollar_lesson_learned_glue/
---
Dev Endpoints are expensive. I was only running it for 5 days and apparently it could have gotten upwards of $1,000/month

Oops.
## [3][Restore RDS from S3 keeps launching in wrong VPC](https://www.reddit.com/r/aws/comments/faa2c8/restore_rds_from_s3_keeps_launching_in_wrong_vpc/)
- url: https://www.reddit.com/r/aws/comments/faa2c8/restore_rds_from_s3_keeps_launching_in_wrong_vpc/
---
I keep trying to create a new RDS db from a s3 backup.  Even though I specifically set a VPC to use, it always launches into the wrong VPC.  Last week when I did this, it worked as expected.  I have several instances running, in the correct VPC, that had been restored from s3.

The only weird thing I've noticed is that the gui has changed.  The bucket selection says "Write audit logs to S3", and gives me the option to select a bucket.

[Here's the screen from last week](https://imgur.com/a/ixUmwPD)
and
[Here's the new restore screen](https://imgur.com/a/DcDvaq2)

Any thoughts?  I've done this before, and I don't know what's going on now.  Thanks in advance.
## [4][How exactly are we supposedly to BOTO3’s CloudFront.generate_presigned_url()](https://www.reddit.com/r/aws/comments/fa09t4/how_exactly_are_we_supposedly_to_boto3s/)
- url: https://www.reddit.com/r/aws/comments/fa09t4/how_exactly_are_we_supposedly_to_boto3s/
---
I am trying to generate presignedUrls for my CloudFront distribution. So far, I’ve seen only one way to do it in python - the example given in BOTO3’s CloudFront doc which uses botocore https://boto3.amazonaws.com/v1/documentation/api/latest/reference/services/cloudfront.html

The doc on it isn’t very expansive on how to use the boto3 presignedUrl api.

My question is: how do we use the boto3’s CloudFront generate_presigned_url()? What parameters should go in its ‘params’?
## [5][Cognito Cloudformation examples](https://www.reddit.com/r/aws/comments/fa9mnf/cognito_cloudformation_examples/)
- url: https://www.reddit.com/r/aws/comments/fa9mnf/cognito_cloudformation_examples/
---
Hi are there any good stack examples for deploying a user pool with SAML metadata etc? Ideally something that sets that up so I can demo using that configured auth with an ALB would be great. I'll continue to search github but can see there are quite a few new cf resources for cognito in the last few months, so want to avoid custom resources where possible
## [6][Serverless Project Example? Open Source](https://www.reddit.com/r/aws/comments/fa2zsw/serverless_project_example_open_source/)
- url: https://www.reddit.com/r/aws/comments/fa2zsw/serverless_project_example_open_source/
---
Anyone know of a large open source serverless application I can take a look at?

I’m interested in how other teams are packaging/deploying large amounts of lambdas and maintaining an API with it as well.

Thanks.
## [7][Easiest way bring up a single ECS container before updating an ECS service](https://www.reddit.com/r/aws/comments/fa95kn/easiest_way_bring_up_a_single_ecs_container/)
- url: https://www.reddit.com/r/aws/comments/fa95kn/easiest_way_bring_up_a_single_ecs_container/
---
I have a bunch of services running on ECS for which I would like to run DB migrations, these migrations are designed to be backwards and forwards compatible so we can run them automatically during deployments of new versions. I do not want all containers in a service to run migrations during updates, for example if an ECS service consists of 4 containers and I push a new image then I do not want 4 new containers all running DB migrations at the same time.

With this in mind, I would like to have a single container come up before all other containers and run DB migrations, then I would like to replace all the other containers in one go to keep things fast, has anyone come up with a nice pattern for this? I don't think that [Enhanced Container Dependency Management](https://aws.amazon.com/about-aws/whats-new/2019/03/amazon-ecs-introduces-enhanced-container-dependency-management/) will help here as far as I can tell as it would bring up an init container for each container that is being replaced.

My initial idea is to have 2 ECS services, a "migration" one that will only every have a single container and a "real" one that have any amount of containers and will actually serve traffic, the process of pushing a new image would be:

* Update migration ECS service with new task definition
* Migration container comes up and runs migrations
* Update real ECS service with new task definition
* Once the real service is up delete the migration service

My main issue with this method is that it requires 2 ECS services so updates might be a bit slow.
## [8][root account management in large organizations with many accounts and control tower.](https://www.reddit.com/r/aws/comments/fa94de/root_account_management_in_large_organizations/)
- url: https://www.reddit.com/r/aws/comments/fa94de/root_account_management_in_large_organizations/
---
How are people currently managing their new account root passwords as a result of using the control tower account factory? 

We have a fairly high level of automation surrounding deploying workloads into new accounts that we are provisioning with control tower, however the root password on the account is a sticking point. 

At present 

1. I forgot my password
2. reset that password
3. enable two factor

is (best I can tell) an entirely manual process. Does anybody have a more effective solution than doing this? Rolling out 10 accounts in a day is not uncommon and will get more and more common, but that manual procedure is both time costly and rather prone to error. Would love to hear opinions about how people are handling this.
## [9][Does this diagram make sense?](https://www.reddit.com/r/aws/comments/f9w4tr/does_this_diagram_make_sense/)
- url: https://www.reddit.com/r/aws/comments/f9w4tr/does_this_diagram_make_sense/
---
So basically I want to send image data and sensor from my raspberry pi to my cloud. My EC2 instance would act as a web server that takes the data from the database server (TYPO in image. I meant database server not storage server) and camera storage server and displays it. I'm using Lamda to compress the images/videos before storing it.

Should the Database server and camera storage be connected with each other? Since the image file path will be stored in the database server ?

Thank you (I'm currently learning AWS, so my solution/design might be atrocious.)

https://preview.redd.it/ibkxh7hxhaj41.png?width=724&amp;format=png&amp;auto=webp&amp;s=c42afa3584261b625d4ac7246503a76ebaf46f9c
## [10][Any workaround Lamda 250MB mighty limit?](https://www.reddit.com/r/aws/comments/fa7rh9/any_workaround_lamda_250mb_mighty_limit/)
- url: https://www.reddit.com/r/aws/comments/fa7rh9/any_workaround_lamda_250mb_mighty_limit/
---
Im using puppet, trying round some stuff, so NEED MOAR SPACE!
