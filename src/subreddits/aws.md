# aws
## [1][Python 2 End of Life (EOL): What It Means for You, Plus Services Like AWS Lambda and Heroku](https://www.reddit.com/r/aws/comments/fcjl8o/python_2_end_of_life_eol_what_it_means_for_you/)
- url: https://www.icanteven.io/blog/python2-eol/
---

## [2][S3 download all objects](https://www.reddit.com/r/aws/comments/fcr0q9/s3_download_all_objects/)
- url: https://www.reddit.com/r/aws/comments/fcr0q9/s3_download_all_objects/
---
Hey all,

I’m looking for a way to download all files (and all versions) from an s3 bucket recursively, to a local filesystem. I’m hoping there’s an easily distinguishable way to denote thee versioned files locally. 

Is there an easy way to do this with the aws-cli or is that beyond what it can do? 

I don’t mind using a 3rd party tool either. Thanks!
## [3][is it possible to leave no trail behind in this case?](https://www.reddit.com/r/aws/comments/fcthjq/is_it_possible_to_leave_no_trail_behind_in_this/)
- url: https://www.reddit.com/r/aws/comments/fcthjq/is_it_possible_to_leave_no_trail_behind_in_this/
---
Hello!

My instances are locked behind a security group that only allows traffic through ports 80 and 443. When I need access, I use a custom batch script to allow traffic through ports 22 and 5432 exclusively to my IP address. Then I proceed to access it with putty using my key pair. Once I'm done, I use another custom script to close ports 22 and 5432.

AWS has CloudTrail, which records all activity for your account. I've noticed that I can monitor security group changes (such as those that I explained above) and I want to know if having these records is enough to tell if someone got into my instance.

So, my questions are:

1) Can anyone access the instances behind that security group without having to open port 22 AND physically having access to my key pair file?

2) Can I trust CloudTrail records, so that all breaches are guaranteed to be logged just like normal access?

Thanks in advance!
## [4][Caching content lambda or gateway](https://www.reddit.com/r/aws/comments/fcrxc1/caching_content_lambda_or_gateway/)
- url: https://www.reddit.com/r/aws/comments/fcrxc1/caching_content_lambda_or_gateway/
---
I have a proxy apigateway -&gt; lambda -&gt; dynamodb stack which is exposed as GET and POSTs.  I'm trying to find out the best way to cache content for a certain amount of time on the GET, and to check if the data had changed or not (for example using etags).

I'm trying to prevent unnecessary calls to my lambda (which read from dynambodb) - although the data could change at any given point so its not clear how to handle that.

As far as I can tell - you can cache at api gateway, or we can use the user browser caching.

&amp;#x200B;

1. Are there any examples of caching at the api gateway layer, *specifically* using the Aws CDK?
2. If caching at the user browser level, should I do this in my lambda and set the headers manually in the response?  I cannot seem to find any examples of this, perhaps its not a common approach?  A python example would be ideal.  Can we use etags, and if so do we have to manually deal with the etag as my understanding is that the server normally deals with this.

So something like:

Cache-Control: private, max-age=0, no-cache 

Etag: &lt;some\_tag&gt;

Any advantages to either approach?

Hopefully someone can point me in the right direction on how to cache responses and only pull the data again in the GET if it has changed.  Thanks!
## [5][AWS CDK best practices (moving from Terraform+Terragrunt)](https://www.reddit.com/r/aws/comments/fcubc1/aws_cdk_best_practices_moving_from/)
- url: https://www.reddit.com/r/aws/comments/fcubc1/aws_cdk_best_practices_moving_from/
---
Well, I've been playing with CDK for a while and became huge fan of it. I am about to introduce it to my team, but first I need to clarify some of the best practices. We've been using Terraform heavily with its "extension" Terragrunt. I would love to map my current experiences to the CDK world.

First question - **how do you version your infrastructure?**  Imagine you want to add a new module with a fresh AWS RDS instance.

Here's our Terraform flow - you play on your branch with the infrastructure by adding all the required resources. After it is ready, you merge it into master branch. Lastly a `git tag` is created and the change is then considered released(not deployed). When somebody wants to update the environment to the new version(deploy your changes), they bump the version in `source:@git/blahblahblah` in `terragrunt.hcl` and perform apply. Basically, we use git tags and Terragrunt to do it.

How does it look like in CDK? Do you version your infrastructure by creating node packages and the actual update is performed by updating the version in package.json, installing the new one and performing the `cdk deploy`? This is what came into my mind first, but perhaps I am missing something, and there is another way to do it.

Second question concerns **multiple environments - you always have qa, preprod, prod and "local" environments used by developers.** In Terraform/grunt, we have a repository with mutliple directories  `qa/terragrunt.hcl`, `preprod/terragrunt.hcl` and so on. If you want to tweak the environment(by changing its params or bumping the infra version in the source git tag), you change the adequate `terragrunt.hcl`, perform apply and commit your changes on succes. How do you solve this usecase in AWS CDK?
## [6][ELI5: When should you use RDS over DynamoDB?](https://www.reddit.com/r/aws/comments/fchhac/eli5_when_should_you_use_rds_over_dynamodb/)
- url: https://www.reddit.com/r/aws/comments/fchhac/eli5_when_should_you_use_rds_over_dynamodb/
---
I heard that even at AWS, all newly services/products are built on DDB unless they have a compelling reason not to because

Of the maintainability, performance and scalability benefits that DDB offers 

With that in mind, when would you want to use RDS over DynamoDB?

If you need joins that are offered in SQL, you could define multiple tables in DDB and load from those tables as needed.

One other thing I noticed was that RDS has a set cost whereas in DDB, you pay for WCU and RCU usage
## [7][Cloudwatch Log Insights - @logStream has no link](https://www.reddit.com/r/aws/comments/fcuhkb/cloudwatch_log_insights_logstream_has_no_link/)
- url: https://www.reddit.com/r/aws/comments/fcuhkb/cloudwatch_log_insights_logstream_has_no_link/
---
For some reason when retrieving the @logStream field when using cloudwatch insights, the field is no longer a link to the messages position in the log stream. 
When using the regular search, I can just click on the log stream link and get taken to the context of a message, this was very useful.

Anyone know if something changed and whether it's possible to get this behaviour again?
## [8][Rewrite Forwarded Paths from an Application Load Balancer](https://www.reddit.com/r/aws/comments/fctruy/rewrite_forwarded_paths_from_an_application_load/)
- url: https://www.reddit.com/r/aws/comments/fctruy/rewrite_forwarded_paths_from_an_application_load/
---
Is there a way for me to rewrite the paths of forwarded traffic?

For example, I have a listener rule that forwards `/test/*` to target group `tg-1`. If I hit `https://url.com/test/endpoint`, I want the target container to see the request going to `/endpoint` and not `/test/endpoint`. Is this something that is possible, or am I just going to have to rewire my endpoints in the service?
## [9][Refunds for unused resources?](https://www.reddit.com/r/aws/comments/fctpaq/refunds_for_unused_resources/)
- url: https://www.reddit.com/r/aws/comments/fctpaq/refunds_for_unused_resources/
---
I got charged $156 for setting up Kubernetes and an Elastic service. I'm not too seasoned with AWS and actually ended up going with GCS, leaving the AWS services untouch.

Simple question: has anyone any experience getting refunds for services that went completely unused? I opened a ticket last night but then found this thread of people pleading to get their refunds honored so I'm not 100% confident. https://forums.aws.amazon.com/thread.jspa?threadID=61842
## [10][Is there way to get a response for successful AWS CLI execution?](https://www.reddit.com/r/aws/comments/fcteaq/is_there_way_to_get_a_response_for_successful_aws/)
- url: https://www.reddit.com/r/aws/comments/fcteaq/is_there_way_to_get_a_response_for_successful_aws/
---
Hi,

I tried `echo $?` but then I think it just returns a response if the command is started executing successfully. It doesn't tell if the command did what it was supposed to do. Is there a way to track that in code?

For example if I run

`aws ec2 stop-instance &lt;instance_ID&gt;`

it should return a code after stopping the instances successfully or returns some other code if it failed due to access issues or may be due to instance not found.

Thanks
