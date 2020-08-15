# aws
## [1][The scaling latency of Lambda with the cost of EC2?](https://www.reddit.com/r/aws/comments/ia1op0/the_scaling_latency_of_lambda_with_the_cost_of_ec2/)
- url: https://www.reddit.com/r/aws/comments/ia1op0/the_scaling_latency_of_lambda_with_the_cost_of_ec2/
---
The problem: create a system that can handle thousands of http requests per second. Traffic cycles up and down daily but is always at a high level, and there are unpredictable, massive spikes. Each request is fairly quick and easy to fulfill (basically just forward to Kinesis).

My initial solution used Lambda + API Gateway. It worked great, but the cost of running it ended up being just too much.

My current solution uses an App ELB + AutoScaling Group of EC2 (reserved/spot) instances. It works pretty well for the daily sinusoidal cycle, but when the massive traffic spikes occur it doesn't respond nearly fast enough, and requests instantly start failing as the instances become overburdened. The Cloudwatch alarm goes off eventually, then the new instance has to boot up, get registered in the ELB, etc. By the time the new instance is running, the spike has usually mostly calmed down, and the new instance is taken down shortly thereafter.

Is there a way to have my cake and eat it too here? The AG solution is more cost-effective and works 90% of the time, but I feel bad about failing those requests and having to rely on clients retrying them.

Additionally, the AG setup I have is very manual at the moment. Deployment consists of manually launching a new instance with the AMI, deploying to it, saving that as a new AMI, updating the AG to use that going forward, and deploying to live instances. The process is scriptable and that's TODO but still annoying. Are there some better tools to do this with? (I've tried my hand at CloudFormation and I want to stay far away from that oh god)
## [2][I want to learn cloudformation, docker, kubernetes etc.](https://www.reddit.com/r/aws/comments/ia5x1k/i_want_to_learn_cloudformation_docker_kubernetes/)
- url: https://www.reddit.com/r/aws/comments/ia5x1k/i_want_to_learn_cloudformation_docker_kubernetes/
---
I am really confused where and how can I start. I am in last year of my college, companies are coming from placements, really need to gain some real world skills. Guide me out a little here please. I just know how to launch EC2 instances.
## [3][Amazon S3 Path Deprecation Plan (Deprecate date: Sep 30, 2020)](https://www.reddit.com/r/aws/comments/ia5wy7/amazon_s3_path_deprecation_plan_deprecate_date/)
- url: https://aws.amazon.com/blogs/aws/amazon-s3-path-deprecation-plan-the-rest-of-the-story/
---

## [4][aw cognito how to see the country of login?](https://www.reddit.com/r/aws/comments/ia66er/aw_cognito_how_to_see_the_country_of_login/)
- url: https://www.reddit.com/r/aws/comments/ia66er/aw_cognito_how_to_see_the_country_of_login/
---
I am fairly new to cognito and I am just wondering how can I check where did the user logged in?
## [5][Hide my API Endpoints (Serverless)](https://www.reddit.com/r/aws/comments/i9szse/hide_my_api_endpoints_serverless/)
- url: https://www.reddit.com/r/aws/comments/i9szse/hide_my_api_endpoints_serverless/
---
Hey guys,

I'm trying to build a platform using a Serverless Architecture using S3, Lambda and API Gateway.

Now, while this worked great for a hobby project, for an actual production grade project, I cannot have my API Endpoints exposed for obvious reasons. It wouldn't take any half decent programmer more than  5 minutes to build an API Abuse script using Python where they just keep calling it.

API Keys won't do it. Reason: Let's say a new user visits my website and to signup, they need to enter their mobile and recieve an OTP. Now this can be easily exploited causing me to rack up bills for millions of SMS. And I cannot assign an API Key for every unknown user (Again, theoretically, I could but it's pointless)

Any and all suggestions are welcome, thanks !
## [6][Redshift + Kinesis connectivity issues](https://www.reddit.com/r/aws/comments/ia6jxg/redshift_kinesis_connectivity_issues/)
- url: https://www.reddit.com/r/aws/comments/ia6jxg/redshift_kinesis_connectivity_issues/
---
I set up a kinesis data firehose delivery stream that has a kinesis data stream source and the destination is redshift. The redshift cluster is not publicly accessible and is in a private subnet group. Because of this, the firehose connection to the redshift cluster is failing. I whitelisted the kinesis firehose public IP address in my security group which is also included in the redshift cluster's security group but the connection is still failing. What else should I look at?
## [7][App sync unique field](https://www.reddit.com/r/aws/comments/ia3uxh/app_sync_unique_field/)
- url: https://www.reddit.com/r/aws/comments/ia3uxh/app_sync_unique_field/
---
Does anyone know how to make a dynamo db field have a unique “constraint” on it?  Haven’t seen anything in the docks. 

Say I have an email field, and I want that field to have only unique values in the table, how would I achieve that?
## [8][AWS Online Tech Talks for August 2020](https://www.reddit.com/r/aws/comments/i9tufw/aws_online_tech_talks_for_august_2020/)
- url: https://aws.amazon.com/blogs/aws/aws-online-tech-talks-for-august-2020/
---

## [9][Best Practice - Updating dns record for multiple subdomains when migrating to new ELB](https://www.reddit.com/r/aws/comments/ia38ej/best_practice_updating_dns_record_for_multiple/)
- url: https://www.reddit.com/r/aws/comments/ia38ej/best_practice_updating_dns_record_for_multiple/
---
Suppose if I have 100s of subdomains like [client1.example.com](https://client1.example.com) [client2.example.com](https://client2.example.com)...etc

All these subdomains are mapped to an ELB's dns name using an alias record, Now in case I want to migrate all my subdomains to new ELB I have to update the alias record for all subdomains using a bash script. So instead of doing this now I am planning to create a new subdomain [my-prod-elb.example.com](https://my-prod-elb.example.com) which is mapped to new elb using alias record and rest of my subdomains will be using this domain as cname record. With this setup i don't have to update all other subdomains when i have to migrate to new subdomain i just only have to change alias record for [my-prod-elb.example.com](https://my-prod-elb.example.com).

&amp;#x200B;

* Current setup: [client1.example.com](https://client1.example.com) \&gt; alias record -&gt; elb dns name
* Planned setup  :    [client1.example.com](https://client1.example.com) \&gt; cname record -&gt; [my-prod-elb.example.com](https://my-prod-elb.example.com) \- &gt; alias record -&gt; elb dns name

Is this a bad practice ? Will it cause dns issues ?

According to aws route53 docs quires to alias records are free but quires to cname records are billable, is this correct ?
## [10][Unable to import Pandas in AWS Lambda layer](https://www.reddit.com/r/aws/comments/i9wn3i/unable_to_import_pandas_in_aws_lambda_layer/)
- url: https://www.reddit.com/r/aws/comments/i9wn3i/unable_to_import_pandas_in_aws_lambda_layer/
---
I uploaded python pandas to Lambda and when I run the Lambda, I see the following error:

        "errorMessage": "Unable to import module 'lambda_function': C extension: No module named 'pandas._libs.interval' not built. If you want to import pandas from the source directory, you may need to run 'python setup.py build_ext --inplace --force' to build the C extensions first."

I am really not sure what is supposed to be going on here. But for more context, I created a directory called "python" and then I run the following line on my terminal

        python3.8 -m pip install pandas -t .

I then zip the "python" directory and then create a new layer and upload the zip file.

&amp;#x200B;

Struggling to see where I am following short.
