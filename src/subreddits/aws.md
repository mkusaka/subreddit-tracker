# aws
## [1][Is it fair to say Oracle is ~3x more expensive than Aurora?](https://www.reddit.com/r/aws/comments/g5anrh/is_it_fair_to_say_oracle_is_3x_more_expensive/)
- url: https://www.reddit.com/r/aws/comments/g5anrh/is_it_fair_to_say_oracle_is_3x_more_expensive/
---
Comparing Multi-AZ RDS between MySQL Aurora and Oracle SE2 Multi-AZ https://s.natalian.org/2020-04-21/oracle.png

Or are there better rules of thumbs here to better make sense of the AWS database landscape?
## [2][AWS Elastic Kubernetes Service: running ALB Ingress controller](https://www.reddit.com/r/aws/comments/g5du9t/aws_elastic_kubernetes_service_running_alb/)
- url: https://www.reddit.com/r/aws/comments/g5du9t/aws_elastic_kubernetes_service_running_alb/
---
A Kubernetes controller which actually controls AWS Application Load Balancers (ALB) in an AWS account when an Ingress resource with the kubernetes.io/ingress.class: alb annotation is created in a Kubernetes cluster.

Article: https://medium.com/setevoy4/aws-elastic-kubernetes-service-running-alb-ingress-controller-8d0d457615fa?source=friends_link&amp;sk=f68ab36f4a1e0f4398a406109a303068
## [3][Migrate Unix to AWS?](https://www.reddit.com/r/aws/comments/g5dbqo/migrate_unix_to_aws/)
- url: https://www.reddit.com/r/aws/comments/g5dbqo/migrate_unix_to_aws/
---
How are people approaching a UNIX workload conversion to cloud infrastructure?

Context: a customer called my firm asking about cloud infra migration and it includes x86/x64 as well as Unix (not sure if aix/ FreeBSD).

Is it always Unix —&gt; Linux (app port /rebuild)—&gt; cloud ?
## [4][Urgently need to blacklist an IP from my API Gateway/Lambda](https://www.reddit.com/r/aws/comments/g4wf4p/urgently_need_to_blacklist_an_ip_from_my_api/)
- url: https://www.reddit.com/r/aws/comments/g4wf4p/urgently_need_to_blacklist_an_ip_from_my_api/
---
Someone is spamming my services and it's beginning to incur costs. I shut down my services for now, to the dismay of my users, until I can blacklist this jerk's IP address.

However, I have one of those weird Lambda/API Gateway proxy setups, and it's not matching any of the documentation on how to do this. Can anyone help me out?

&amp;#x200B;

EDIT:

Unfortunately, my lambda is an HTTP API, which does not support Resource Policies NOR WAF, according to Amazon themselves.

I need to take a step back and re-think security. It's secure as is, just vulnerable to spamming. I may use a REST API with API Keys and/or WAF. Thanks all for the suggestions, I'm still welcome to more.
## [5][Hooks/Test cases for Athena queries before pushing to Github](https://www.reddit.com/r/aws/comments/g5edt8/hookstest_cases_for_athena_queries_before_pushing/)
- url: https://www.reddit.com/r/aws/comments/g5edt8/hookstest_cases_for_athena_queries_before_pushing/
---
Hi all!


Is there a way to add tests cases to see if Athena Queries are parsed properly before pushing to Github?

Wanted to make sure the query has no errors before pushing to Github.
## [6][EBS which volume type for ECS cluster](https://www.reddit.com/r/aws/comments/g5ecdn/ebs_which_volume_type_for_ecs_cluster/)
- url: https://www.reddit.com/r/aws/comments/g5ecdn/ebs_which_volume_type_for_ecs_cluster/
---
Hi, we have ecs cluster that run stateless web application containers. I created ecs cluster with terraform and it uses standard ebs volume time by default however I looked that there are also other volume types available. How can I pick the best type or calculate what i need? We actually only need 5 Gib of storage but ECS optimized ami requires at least 30. We plan to have at least 6 ecs clusters so it's important for to understand how to optimize ECS cluster storage
## [7][Kinesis streaming to multiple lambdas, will all of them receive the same message?](https://www.reddit.com/r/aws/comments/g5ayjd/kinesis_streaming_to_multiple_lambdas_will_all_of/)
- url: https://www.reddit.com/r/aws/comments/g5ayjd/kinesis_streaming_to_multiple_lambdas_will_all_of/
---
I have 3 Lambdas subscribed to a same Kinesis data stream. When that single Kinesis stream sends out `MessageA`, will all the 3 Lambdas that are subscribed to it receive a copy of `MessageA` too? 

Or will only one of the Lambdas receive `MessageA`?
## [8][Handling binary uploads using AWS Lambda](https://www.reddit.com/r/aws/comments/g56q6l/handling_binary_uploads_using_aws_lambda/)
- url: https://www.reddit.com/r/aws/comments/g56q6l/handling_binary_uploads_using_aws_lambda/
---
Background: I'm implementing a script that sends some analytics from a website to our server. Because of the size, I used zlib to compress the logs to around 30KB or so. The uncompressed JSON is about 1-2MB.

Problem: What's the best approach to let AWS Lambda read the compressed file? Implementing it via traditional EC2 is easy, but if I use Lambda, most of the online articles suggest that I upload the files to S3. Is there a better way to do this? Can I upload files to Lambda directly?
## [9][Recent review or experiences with Amazon Connect](https://www.reddit.com/r/aws/comments/g562r9/recent_review_or_experiences_with_amazon_connect/)
- url: https://www.reddit.com/r/aws/comments/g562r9/recent_review_or_experiences_with_amazon_connect/
---
I saw some older reviews of Connect from a year or two back, but wondering if anyone is using it for a production call center and can comment on their experience. I'm curious in general experiences, but also if you feel the product has been continuously improving as I know there's some things AWS just "tries out" and then doesn't really invest in. I want to make sure it's a product I can grow with.

For context:

\- I have an outsourced call center I run with 12 agents today. We are growing fast though so I'd like something that can easily scale to 100+ in a year or so.-

\- I have a few other services running in AWS (lightsail for a POS server and AWS workspaces for agent terminals in the cloud) but not sure how much that makes a difference having it all in one place.
## [10][Getting "cannot open shared object file: No such file or directory - /var/task/vendor/bundle/ruby/2.7.0/gems/pg-1.2.3/lib/pg_ext.so" when trying to run ruby code on AWS lambda.](https://www.reddit.com/r/aws/comments/g4zn8b/getting_cannot_open_shared_object_file_no_such/)
- url: https://www.reddit.com/r/aws/comments/g4zn8b/getting_cannot_open_shared_object_file_no_such/
---
i zipped all my code and uploaded it to AWS lambda, when i try running test i am getting this error. its  almost a week now since i have been stuck on this issue.

    {
      "errorMessage": "libldap_r-2.4.so.2: cannot open shared object file: No such file or directory - /var/task/vendor/bundle/ruby/2.7.0/gems/pg-1.2.3/lib/pg_ext.so",
      "errorType": "Init&lt;LoadError&gt;",
      "stackTrace": [
        "/var/lang/lib/ruby/2.7.0/rubygems/core_ext/kernel_require.rb:92:in `require'",
        "/var/lang/lib/ruby/2.7.0/rubygems/core_ext/kernel_require.rb:92:in `require'",
        "/var/task/vendor/bundle/ruby/2.7.0/gems/pg-1.2.3/lib/pg.rb:5:in `&lt;top (required)&gt;'",
        "/var/lang/lib/ruby/2.7.0/rubygems/core_ext/kernel_require.rb:92:in `require'",
        "/var/lang/lib/ruby/2.7.0/rubygems/core_ext/kernel_require.rb:92:in `require'",
        "/var/task/lambda_function.rb:2:in `&lt;top (required)&gt;'",
        "/var/lang/lib/ruby/2.7.0/rubygems/core_ext/kernel_require.rb:92:in `require'",
        "/var/lang/lib/ruby/2.7.0/rubygems/core_ext/kernel_require.rb:92:in `require'"
      ]
    }

i tried deploying the \`vendor\` and \`lib\` as a layer on lambda and i still get the same error.

i tried this answer too but no help [https://www.reddit.com/r/ruby/comments/a3e7a1/postgresql\_on\_aws\_lambda\_ruby/](https://www.reddit.com/r/ruby/comments/a3e7a1/postgresql_on_aws_lambda_ruby/)this is my repo is you want to give it a try [https://bitbucket.org/nijeeshjoshy/aws-lambda-ruby-postgres](https://bitbucket.org/nijeeshjoshy/aws-lambda-ruby-postgres)  


i added this repo [https://github.com/mphsi/ruby\_lambda\_pg\_layer](https://github.com/mphsi/ruby_lambda_pg_layer)  as a layer and then it fixed all the errors related to \`postgres\` but now it says active\_record not found.
