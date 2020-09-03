# aws
## [1][168 AWS Services in 2 minutes](https://www.reddit.com/r/aws/comments/ilhhwk/168_aws_services_in_2_minutes/)
- url: https://www.youtube.com/watch?v=BtJAsvJOlhM
---

## [2][168 AWS Services in 2 Minutes](https://www.reddit.com/r/aws/comments/ilhhbp/168_aws_services_in_2_minutes/)
- url: https://twitter.com/forrestbrazeal/status/1301182968766107656?s=19
---

## [3][Can AWS Lambda be used to achieve my performance requirements, if so how?](https://www.reddit.com/r/aws/comments/ilsi0v/can_aws_lambda_be_used_to_achieve_my_performance/)
- url: https://www.reddit.com/r/aws/comments/ilsi0v/can_aws_lambda_be_used_to_achieve_my_performance/
---
I have a requirement where a single API request generates 1000 individual independent tasks.

Today we run as a C application in a single [c5.9xlarge](https://aws.amazon.com/ec2/instance-types/) EC2 instance and share the tasks across all the threads on the EC2 instance.

The tasks:

1. Are 100% cpu bound, there is no I/O or external interaction, they do a lot of calculations.
2. Are 100% independent, there is no relationship/dependancy between the tasks when they are executing.
3. Only use a small amount of memory.
4. On completion, generate some output. The output from all the tasks is merged/consolidated into a single response payload which is less than 3kb in size in json.

The response payload must include 100% of the tasks, there is no tolerance for missing task outputs.

With the current architecture ignoring network hops the processing time is less than 1.5 seconds.

I am considering porting this to AWS Lambda because:

1. The requests are sparse/intermittent with significant periods of inactivity
2. Lambda may be cheaper - When I take into account dev, test, qa and production, etc the cost for EC2 instances becomes a problem
3. Lambda may be more scaleable - Under the current architecture a single EC2 instance is required to process a single request. If I want to process 7 requests concurrently (during a period of bust activity) I would need to have 7 EC2 instances provisioned. Given my 1.5 seconds response time, I don't see dynamic ec2 scaling as helpful in this instance.
4. Perhaps with Lambda I can either improve my performance (reduce the response time down from 1.5 seconds) OR increase my calculation accuracy by running even more than 1000 tasks but still within 1.5 seconds.

I realise that the AWS Lambda supported languages [(Java, Go, PowerShell, Node.js, C#, Python, and Ruby code)](https://aws.amazon.com/lambda/faqs/#:%7E:text=AWS%20Lambda%20natively%20supports%20Java,%2C%20C%23%2C%20Go%20and%20PowerShell.), are unlikely to be as fast as C, however I think with some focussed optimisation the difference in processing time for a task will be minor and more than made up for by the greater potential concurrency of Lambda. I am happy to use any language. To be honest the focus of this question is not around the selection of the language, rather the architecture. I do realise that different language implementation will have different characteristics [regarding startup time and execution time](https://aws.amazon.com/blogs/compute/new-for-aws-lambda-predictable-start-up-times-with-provisioned-concurrency/).

My initial idea is to fan-out each task to run as an individual Lambda, then merge the results.

I noticed [this article](https://medium.com/circuitpeople/patterns-for-aws-lambda-fanout-a7662bb0b2aa) with the author describing a roughly 50ms blocking delay on just the invocation of **await lambda.InvokeAsync()**. That in itself is potentially a show stopper for my performance requirements.

I also wondered if I should use a queue based architecture by writing 1000 task request messages to a queue and allowing 1000 Lambda's to be triggered to process them concurrently. However then I need to merge the results, presumably by writing all the outputs to another SQS, where a results processor can merge them. I will actually have multiple requests running concurrently so the architecture needs to scale to support that.

I understand that by default AWS Lambda is [restricted to 1000 concurrent invocations per account](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-limits.html), however I understand that this can be increased by asking AWS support.

As you can see there are considerable unknowns, is AWS Lambda likely to work in the end for my performance requirement (1000 tasks in 1.5 seconds), if so, what is the recommended architecture, particularly in regards to the fanout-gather of responses.

&amp;#x200B;

NOTE: This is also on [Stack Overflow](https://stackoverflow.com/questions/63715548/can-aws-lambda-be-used-to-achieve-my-performance-requirements-if-so-how), a poster there suggested I put it here too.
## [4][WorkMail is Proving to be a bit expensive](https://www.reddit.com/r/aws/comments/ilpetb/workmail_is_proving_to_be_a_bit_expensive/)
- url: https://www.reddit.com/r/aws/comments/ilpetb/workmail_is_proving_to_be_a_bit_expensive/
---
Hello, 

I've recently moved one of my projects to aws, everything is great, except for the cost of Workmail. Before this i was using shared web hosting, in which email hosting was free so client had many email ids, But now in workmail its $4 per user per month, which if u compare to the pervious (free) plan is a lot. Now i understand it can't be free, but still the current rate is a bit high imo.  I am using many services like EC2, Route53, ECB, SES etc, but none are as expensive as workmail. Just for the last month its expense is around $40.

Can someone please guide me if i am doing something wrong and  how can i bring it lower or any alternatives to the workmail? Details below

* Number of users/ids right now: 7
* Used for sending and receiving emails some are hooked in code so alerts are also sent from and to them.

**Bill for August**

* AWS Service Charges: $137.49
* Data Transfer: $0.00
* Elastic Compute Cloud: $16.31
* Elastic Load Balancing: $18.60
* Route 53: $1.02
* Savings Plans for AWS Compute usage: $49.79
* Simple Email Service: $0.00
* **WorkMail: $39.27**
## [5][Confusion with savings plans "recommended commitments"](https://www.reddit.com/r/aws/comments/ilrjby/confusion_with_savings_plans_recommended/)
- url: https://www.reddit.com/r/aws/comments/ilrjby/confusion_with_savings_plans_recommended/
---
We've been running 3 x t3.small and 8 x t3.medium on-demand EC2 instances in the eu-west-2 region for nearly 6 months and don't predict that our usage patterns are going to change, so we want to buy a 3-year all-upfront EC2 savings plan to cover these instances.

The on-demand pricing for the t3 servers in the eu-west-2 region are:

    t3.small  = 0.0236/hour
    t3.medium = 0.0472/hour

Multiplying these figures by the number of instances we have:

    t3.small  0.0236 * 3 = 0.0708/hour
    t3.medium 0.0472 * 8 = 0.3776/hour

Therefore our total hourly spend on t3 instances is:

    0.0708 + 0.3776 = 0.4484/hour

However, in the AWS billing console the "recommended commitment" for the t3 family for a 3-year term is 0.170/hour. For a 1-year term it is 0.277/hour. Can someone explain why the hourly recommended commitment differs for a 1-year term vs a 3-year term, and why both those figures are so different from 0.4484/hour?
## [6][Estimated Cost](https://www.reddit.com/r/aws/comments/ilreig/estimated_cost/)
- url: https://www.reddit.com/r/aws/comments/ilreig/estimated_cost/
---
hello,

I have been looking how to get estimated cost of on prem env. There are three options on migration hub but none of them fit for me. Import csv and discovery agent is very manual. Discovery connector just available for esxi 6.5 and below. My env has esxi 6.7 and esxi 7.

Is there any other way to can get estimated cost with agentless model for esxi 6.7 and esxi 7 ?

Thanks,
## [7][What is the difference between Eq and Equals in GuardDuty filter??](https://www.reddit.com/r/aws/comments/ilrei1/what_is_the_difference_between_eq_and_equals_in/)
- url: https://www.reddit.com/r/aws/comments/ilrei1/what_is_the_difference_between_eq_and_equals_in/
---
I just queried a filter created in the console and it has both Eq and Equals statements.

[https://docs.aws.amazon.com/cli/latest/reference/guardduty/create-filter.html](https://docs.aws.amazon.com/cli/latest/reference/guardduty/create-filter.html)  


Eq -&gt; (list)

&gt;Represents the *equal* condition to be applied to a single field when querying for findings.

Equals -&gt; (list)

&gt;Represents an *equal* condition to be applied to a single field when querying for findings.

  
TO my mind the difference between "the" and "an" means that Eq is the only. But "an" is one of many. So if I specify both will Eq take priority?  


eg :

            "Criterion": {
                "resource.instanceDetails.instanceId": {
                    "Eq": [
                        "i-123"
                    ],
                    "Equals": [
                        "i-234" , "i-456"
                    ]
                }
             }

Because "the equal condition" is only i-123, the "an equal condition" is ignored?

Or will i- 123,234&amp;456 all trigger the condition?

Or do all statements have to match so that wouldn't actually match anything ever?
## [8][AWS Site-to-Site VPN now supports Internet Key Exchange (IKE) initiation (It's about freaking time!)](https://www.reddit.com/r/aws/comments/il5q9p/aws_sitetosite_vpn_now_supports_internet_key/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/08/aws-site-to-site-vpn-now-supports-internet-key-exchange-initiation/
---

## [9][What is the easiest way to do this?](https://www.reddit.com/r/aws/comments/iljuny/what_is_the_easiest_way_to_do_this/)
- url: https://www.reddit.com/r/aws/comments/iljuny/what_is_the_easiest_way_to_do_this/
---
I have a super super long eBook in a plain .txt file, it's 2M characters.

I'd like to convert this using Polly to a single audiobook as an .mp3 file final output, I am a newbie at coding languages. So, like, given the web console is extremely limited, what would be the easiest way I can do this simple task?
## [10][Practical problem/solution: Download file (~100mb size) from public site when site publishes updated version and store/overwrite in S3 bucket?](https://www.reddit.com/r/aws/comments/iloj7x/practical_problemsolution_download_file_100mb/)
- url: https://www.reddit.com/r/aws/comments/iloj7x/practical_problemsolution_download_file_100mb/
---
Thoughts are: write code (Python or PowerShell) in Lambda function to check if site has newer version uploaded, if so then download file to S3. Set CloudWatch event to run every 10 minutes with a target being this Lambda.  I think Lambda has a \~6mb file size limit though? Look into multi-part download or something along those lines maybe?

Could I somehow use Cloud9 for this? I am trying to think serverless so I dont need a server to run this job and for a practical serverless example. Surely someone somewhere has had to use Lambda to download large files from  some location to S3
