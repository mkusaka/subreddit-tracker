# aws
## [1][Cognito is confusing...](https://www.reddit.com/r/aws/comments/eqsi1u/cognito_is_confusing/)
- url: https://www.reddit.com/r/aws/comments/eqsi1u/cognito_is_confusing/
---
I had set-up my web-app to use cognito for user sign-in, sign-up, MFA and password management some time ago using  `amazon-cognito-identity-js` library. Import statement: 

`import * as AmazonCognitoIdentity from "amazon-cognito-identity-js";`

Now, after some time, I've come back to add few other implementations like fetching user attributes and others following resources like: 

[https://aws.amazon.com/blogs/mobile/accessing-your-user-pools-using-the-amazon-cognito-identity-sdk-for-javascript/](https://aws.amazon.com/blogs/mobile/accessing-your-user-pools-using-the-amazon-cognito-identity-sdk-for-javascript/)

As I was writing code, most of the things are not recognized by above imported library. Example: 

    AWS.config.credentials = new AWS.CognitoIdentityCredentials({})  &lt;-- seems like these are not part of amazon-cognito-identity-js library.

Another resource, from aws js docs: [https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/CognitoIdentityServiceProvider.html#getUser-property](https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/CognitoIdentityServiceProvider.html#getUser-property)

    var cognitoidentityserviceprovider = new AWS.CognitoIdentityServiceProvider(); &lt;-- same error.

What am I missing here? Note that I've tried importing library as AWS.

Edit: I don't want to use amplify sdk since most of my code is already written using js sdk.
## [2][SQS : the specified queue does not exist for this wsdl version](https://www.reddit.com/r/aws/comments/equis1/sqs_the_specified_queue_does_not_exist_for_this/)
- url: https://www.reddit.com/r/aws/comments/equis1/sqs_the_specified_queue_does_not_exist_for_this/
---
We have a microservice running in EKS which is throwing this error. We checked the obvious reason of region being incorrect. It is set correctly.

Is there anything else that we're missing?
## [3][IAM timeouts in us-east 1 with "Http request timed out enforced after 999ms"](https://www.reddit.com/r/aws/comments/eqiac7/iam_timeouts_in_useast_1_with_http_request_timed/)
- url: https://www.reddit.com/r/aws/comments/eqiac7/iam_timeouts_in_useast_1_with_http_request_timed/
---
Seeing numerous timeouts and fails of IAM.  AWS CLI unable to locate new keys.  No info on status pages yet.

UPDATE
12:09PM UTC-5 new keys were recognized by AWS CLI.
## [4][JSON Python Serialization Question](https://www.reddit.com/r/aws/comments/eqr5dn/json_python_serialization_question/)
- url: https://www.reddit.com/r/aws/comments/eqr5dn/json_python_serialization_question/
---
Hi there, 

I'm building a Reddit bot using AWS Lambda, SQS, EC2 and S3. 

Essentially, the bot scans Reddit for a comment, verifies in the database that this comment hasn't been replied to, and replies. The main handler imports a Python file, which creates a class for the user and interacts with the SQL database from that class.

It's telling me that I need to serialize the object in my module, but I'm at a loss to understand why. Why does AWS need a serialized JSON object to run a Python script? Does it need to be in JSON to transfer the information to my SQL database?

Here is my error message:

 "errorMessage": "Object of type module is not JSON serializable",

  "errorType": "TypeError",

  "stackTrace": \[

"  File \\"/var/task/main.py\\", line 5, in handler\\n    encoder = json.JSONEncoder().encode({'output': Scrapey})\\n",

"  File \\"/var/lang/lib/python3.7/json/encoder.py\\", line 199, in encode\\n    chunks = self.iterencode(o, \_one\_shot=True)\\n",

"  File \\"/var/lang/lib/python3.7/json/encoder.py\\", line 257, in iterencode\\n    return \_iterencode(o, 0)\\n",

"  File \\"/var/lang/lib/python3.7/json/encoder.py\\", line 179, in default\\n    raise TypeError(f'Object of type {o.\_\_class\_\_.\_\_name\_\_} '\\n"
## [5][Get a list of instances running Windows](https://www.reddit.com/r/aws/comments/eqlwfq/get_a_list_of_instances_running_windows/)
- url: https://www.reddit.com/r/aws/comments/eqlwfq/get_a_list_of_instances_running_windows/
---
Is there a way to get a list of instances running Windows and it’s version via the CLI or boto3?
## [6][What data fields for images to CSV?](https://www.reddit.com/r/aws/comments/eqpf6x/what_data_fields_for_images_to_csv/)
- url: https://www.reddit.com/r/aws/comments/eqpf6x/what_data_fields_for_images_to_csv/
---
i am creating a program to fetch images (AKIs, AMIs, ARIs) information (not the actual content) and store it in various formats as requested.  i intend to make this program freely available.  what i would like to know is what data fields people would like to have saved in CSV and TSV formats.  obvious fields would include the region name, object ID, name, and creation date/time.  what else?

this is being written in Python.  so it should be fairly portable.

it will have options to select regions and other aspects such as image type.
## [7][Send RDS data to Elasticsearch](https://www.reddit.com/r/aws/comments/eqlxx0/send_rds_data_to_elasticsearch/)
- url: https://www.reddit.com/r/aws/comments/eqlxx0/send_rds_data_to_elasticsearch/
---
I have a postgres db hosted on RDS and am having a hard time figuring out how to send my data over to Elasticsearch. I have tried setting it up with Kinesis data streams however the entire system is starting to look very complex.

Surely I am not the only one that has tried to do this before, so I was wondering if anyone had experiences to share.
## [8][How long does a quota increase generally take?](https://www.reddit.com/r/aws/comments/eqosye/how_long_does_a_quota_increase_generally_take/)
- url: https://www.reddit.com/r/aws/comments/eqosye/how_long_does_a_quota_increase_generally_take/
---
I’m completely new to AWS, so sorry if this sounds impatient. Does it take more than a day to get a quota increase, or should I submit another support request?
## [9][Amazon ECS Preview Support for EFS file systems Now Available](https://www.reddit.com/r/aws/comments/eq8hi6/amazon_ecs_preview_support_for_efs_file_systems/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/01/amazon-ecs-preview-support-for-efs-file-systems-now-available/
---

## [10][What's the difference between the SDK and CDK?](https://www.reddit.com/r/aws/comments/eqmu73/whats_the_difference_between_the_sdk_and_cdk/)
- url: https://www.reddit.com/r/aws/comments/eqmu73/whats_the_difference_between_the_sdk_and_cdk/
---
I made a tool to automate the creation and configuration of several resources (like pipelines and all associated resources) to solve repetitive tasks. I'm currently working on a feature that allows makes using git-flow easier and the code base is becoming almost too big for 1 person to work on (it's basically a CodeStar clone without the user management features and without using cloud formation).

But, I just found out something about a CDK and I'm questioning my methods (as usual) on whether my methods of setting up resources and managing them are correct. As far as I understand it, CDK is a tool to make cloud formation (IAC) easier through languages like javascript and python. Does anyone have a solid idea of what the difference is between the SDK and CDK, and whether I should use the CDK instead of the SDK?
