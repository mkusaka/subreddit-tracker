# aws
## [1][AWS Lambda now supports EFS. Practice hands-on with this exercise.](https://www.reddit.com/r/aws/comments/i5a20j/aws_lambda_now_supports_efs_practice_handson_with/)
- url: http://aws-dojo.com/excercises/excercise7
---

## [2][Hands-on beginner tutorial: Implementing a highly available, scalable and cost-efficient video resizing service using serverless architecture with AWS Lambda, SNS and SQS](https://www.reddit.com/r/aws/comments/i4wd66/handson_beginner_tutorial_implementing_a_highly/)
- url: http://blog.nghdang.com/2020/08/03/implementing-a-highly-available-scalable-and-cost-efficient-video-processing-service-using-serverless-architecture-with-aws-lambda-sns-and-sqs/
---

## [3][Help with Slackbot Chatbot](https://www.reddit.com/r/aws/comments/i5aolq/help_with_slackbot_chatbot/)
- url: https://www.reddit.com/r/aws/comments/i5aolq/help_with_slackbot_chatbot/
---
Hi Everyone,

I’m trying to develop a slack bot that’s is hosted on AWS. 

Originally I thought about using Lex, Lambda, cloud watch and maybe cognito. 

Or could use elastic beanstalk with ec2 and s3? 

I am wanting to implement a CI/CD pipeline as well from my GitHub Repo.

I also need to use API requests on a server not hosted on AWS to gather information for the slack bot.

Any thoughts on how I could achieve this ? Never done anything like this before so looking for some ideas. TIA
## [4][How to get Pipeline name from an SNS message in AWS Lambda?](https://www.reddit.com/r/aws/comments/i5d2sk/how_to_get_pipeline_name_from_an_sns_message_in/)
- url: https://www.reddit.com/r/aws/comments/i5d2sk/how_to_get_pipeline_name_from_an_sns_message_in/
---
Hi all,

I am having big trouble with getting the sourcePipelineName value out of the SNS message triggering one of my functions.

What I have as an example is this:

    def get_ami_id_from_ib_notification(ib_notification):
        for resource in ib_notification['outputResources']['amis']:
            if resource['region'] == os.environ['AWS_REGION']:
                return(resource['image'])
            else:
                return(None)

... and it works perfectly for getting the AMI-ID back. I tried to get sourcePipelineName the same way, but no luck (though I don't need to check for anything additional):

    def get_asg_name_from_ib_notification(ib_notification):
    	for resource in ib_notification:
    		return(resource['sourcePipelineName'])

But it results in:

    [ERROR] TypeError: string indices must be integers
    Traceback (most recent call last):
      File "/var/task/lambda_function.py", line 99, in lambda_handler
        asg_name = get_asg_name_from_ib_notification(ib_notification)
      File "/var/task/lambda_function.py", line 25, in get_asg_name_from_ib_notification
        return(resource['sourcePipelineName'])

According to [https://docs.aws.amazon.com/cli/latest/reference/imagebuilder/get-image.html](https://docs.aws.amazon.com/cli/latest/reference/imagebuilder/get-image.html) there should be a key called sourcePipelineName available.

What am I doing wrong?
## [5][Is there a way to obtain the list of users of a user pool from another account?](https://www.reddit.com/r/aws/comments/i5a7dr/is_there_a_way_to_obtain_the_list_of_users_of_a/)
- url: https://www.reddit.com/r/aws/comments/i5a7dr/is_there_a_way_to_obtain_the_list_of_users_of_a/
---
Basically I'm building a user management portal for my company, and we have users spread across different accounts. 

When I try to get the users from a difference account, I'm getting the error "ResourceNotFoundException: User pool #### does not exist."

I figured that this means that the user pool does not exist in this account. Whenever I switch my AWS account on my IDE, it'll get the users from that account.

    const AWS = require('aws-sdk');
    var docClient = new AWS.DynamoDB.DocumentClient();

    exports.main = (event, context, callback) =&gt; {

    let id = event.pathParameters.id;
    let account = event.pathParameters.account;
    let accessKeyId, secretAccessKey;
    function getAccount() {
      var table = "accounts";
      var params = {
        TableName: table,
        Key: {
          "account_name": account
        }
      };
    docClient.get(params, function (err, data) {
      if (err) {
        callback(err, null);
      } else {
        console.log(data.Item);
        try {
          accessKeyId = data.Item.accessKeyId;
          secretAccessKey = data.Item.secretAccessKey;
          getUsers(accessKeyId, secretAccessKey);
        }
        catch (e) {
          callback(e, null);
        }

      }
    });
    };

     getAccount();

     function getUsers({ accessKeyId, secretAccessKey }) {

    AWS.config.update(
      {
        accessKeyId: accessKeyId,
        secretAccessKey: secretAccessKey,
        region: 'eu-west-1'
      });

    const cognitoidentityserviceprovider = new AWS.CognitoIdentityServiceProvider();
    var params = {
      UserPoolId: id,
      AttributesToGet: [
        'email',
      ],
    };
    cognitoidentityserviceprovider.listUsers(params, (err, data) =&gt; {
      if (err) {
        console.log(err);
        callback(err, null);
      }
      else {
        console.log("data", data);
        callback(null, {
          statusCode: 200,
          body: JSON.stringify(data),
          headers: {
            "Access-Control-Allow-Origin": "*",
            "Access-Control-Allow-Methods": "GET, OPTIONS",
            'Access-Control-Allow-Credentials': true,
            "Access-Control-Allow-Headers": "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With"
          }
        });

      }
    });
  }
};
## [6][AWS custom metric resets to '--' without a regular flow of log events](https://www.reddit.com/r/aws/comments/i5c0bo/aws_custom_metric_resets_to_without_a_regular/)
- url: https://www.reddit.com/r/aws/comments/i5c0bo/aws_custom_metric_resets_to_without_a_regular/
---
I'm a pretty new dev, so I decided to do something fun to help get exposure to other services. I created an ec2 instance to run a minecraft server. I'm using cloudwatch agent to send the minecraft logs to cloudwatch, where I have custom filters that trigger whenever a player leaves or joins to increment a custom metric called playerCount.

The problem that occurs is if there is a gap of time between triggering events, the metric goes back to '--' which I believe is the equivalent of "no data." Is there a way I can set the custom metric so that it holds onto its value over time?
## [7][Redirect 302 with just API gateway](https://www.reddit.com/r/aws/comments/i59hdd/redirect_302_with_just_api_gateway/)
- url: https://www.reddit.com/r/aws/comments/i59hdd/redirect_302_with_just_api_gateway/
---
Hello guys.

Me and our teammates would like to create a redirect **302** using just API gateway or, in general, with less components possible.

AFAIK the API gateway is able to do just 301. [The solution with less moving parts seems to be having a lambda that make this redirect hooked to the API record.](https://kennbrodhagen.net/2016/04/02/how-to-return-302-using-api-gateway-lambda/)

&amp;#x200B;

Do you have other ideas? Are you able to explain us why they didn't provide this ability? Is there any negative/hidden outcomes in REDIRECT 302 instead 301?
## [8][is it possible to run an exe in beanstalk?or do i have to use a different service?](https://www.reddit.com/r/aws/comments/i59d7v/is_it_possible_to_run_an_exe_in_beanstalkor_do_i/)
- url: https://www.reddit.com/r/aws/comments/i59d7v/is_it_possible_to_run_an_exe_in_beanstalkor_do_i/
---
* i want to use it to host my game it consist of exe,mono dlls,and folders, 
* ill be using as a game server, its written in .net made with unity

&amp;#x200B;

i apologize if my question is odd i'm not sure where to put my game server in aws.

thank you.
## [9][What's the correct way to connect Lambdas to an API?](https://www.reddit.com/r/aws/comments/i58ps8/whats_the_correct_way_to_connect_lambdas_to_an_api/)
- url: https://www.reddit.com/r/aws/comments/i58ps8/whats_the_correct_way_to_connect_lambdas_to_an_api/
---
Hello, all. I'm a complete beginner when it comes to AWS. I'm doing a medium scale project that is using an instance of Connect to handle customer service interactions. We have a REST API that handles the requests of whatever information is needed by Connect. Let's say the API exposes 20 end-points, is it necessary to create 20 Lambda instances to handle all possible request or is there any better way?

Thanks in advance.
## [10][SaaS logging solutions](https://www.reddit.com/r/aws/comments/i54i8w/saas_logging_solutions/)
- url: https://www.reddit.com/r/aws/comments/i54i8w/saas_logging_solutions/
---
I'm building a SaaS web app for small businesses, target at 10s or 100s of thousands of customers.

It is a basic accounting app, similar to Quickbooks, but aiming to simplify money management for a more math-anxious, number-averse audience.

As customers use the application, I want to have a **real-time picture of how they're hitting our services** (`/user/create`, `/transactions/list`, `/image/upload`, etc.):

* how many requests (aggregations per hour, day, month and year)
* who the customer is (aggregations per user id and tenant id)
* where the customer is located (aggregations per city, province and country)
* latencies they're getting
* errors they're getting

I'm handling requests with **API Gateway**, and for log emission I will simply activate my API to send **Access Logs** to the **CloudWatch Logs** services. So the log emission and sending to CloudWatch Logs I've already figured out.

But now I need to process these logs and make the aggregations I listed above.

I've contemplated these options:

* Streaming the logs to **Kinesis Firehose**, then processing them and sending to **Amazon Elasticsearch Service** to make the aggregations and visualize with **Kibana**
* Streaming the logs to **Kinesis Firehose**, then processing them and making the aggregations with **Kinesis Data Analytics**, then saving the results to a **DynamoDB** table

I would like a solution that will be:

* **cost-efficient**, both today (with few customers and very sparse logging) and in the possible future as I get thousands of customers hitting the API all day from various locations
* **highly scalable and self-managed** at every level of scale

I know very little about these services. At this point, the **Elasticsearch** path sounds expensive to me, it sounds compute-heavy and storage-heavy. While the **Kinesis Data Analytics** seems more neat and cheap. Or, maybe both solutions are overkill for my case? I could also try **CloudWatch Insights**, but it doesn't seem to handle the particular aggregations I need.

I'd like to note that I don't need very advanced ad-hoc analytics, I already know upfront the aggregations I need (as listed above).

Thoughts?
