# aws
## [1][Example serverless data pipeline for crawling PDFs from the Web and transforming their contents into structured data using AWS Textract. Built with AWS CDK + TypeScript.](https://www.reddit.com/r/aws/comments/fbwtr2/example_serverless_data_pipeline_for_crawling/)
- url: https://github.com/aeksco/aws-pdf-textract-pipeline
---

## [2][On prem over AWS?](https://www.reddit.com/r/aws/comments/fc720o/on_prem_over_aws/)
- url: https://www.reddit.com/r/aws/comments/fc720o/on_prem_over_aws/
---
People who choose on-prem hardware over cloud providers like AWS, what's your reasoning behind you decision? With cloud pricing decreasing overtime, I'm curious as to why people invest on building on-prem hardware instead.
## [3][Limit public IPs to be available just for ELB/ALBs?](https://www.reddit.com/r/aws/comments/fcbfsj/limit_public_ips_to_be_available_just_for_elbalbs/)
- url: https://www.reddit.com/r/aws/comments/fcbfsj/limit_public_ips_to_be_available_just_for_elbalbs/
---
Hi,

for security reasons, I'd like to deny users in a multi-account environment to create ressources with public IP addresses unless it is an ELB/ALB. Would this be possible with an SCP?

AFAIK IAM actions only exist for "AllocateAddress" and "AssociateAddress" which is about Elastic IPs but not normal public IPs.
## [4][Migrating t2 instance to t3 with RAID](https://www.reddit.com/r/aws/comments/fcbf9t/migrating_t2_instance_to_t3_with_raid/)
- url: https://www.reddit.com/r/aws/comments/fcbf9t/migrating_t2_instance_to_t3_with_raid/
---
Hello!
We had some issues on friday when we wanted to migrate our t2.med instance that used raid10.

Upon migration to t3 the RAID field that we had didn't work, because the new t3 drives were NVMe and labled as nvmeXn1.

The old drives were xvda.

Is there an easy-ish way of migrating this?
## [5][Help me clear something regarding MFA-Delete](https://www.reddit.com/r/aws/comments/fcb6fo/help_me_clear_something_regarding_mfadelete/)
- url: https://www.reddit.com/r/aws/comments/fcb6fo/help_me_clear_something_regarding_mfadelete/
---
Hi,To my understanding, Bucket versioning can be enabled by any authorized user (root/bucket owner/IAM user) through IAM policies. MFA-Delete though can only be enabled by the root user (through CLI/API with issued keys). Bearing this in mind, what happens in this scenarios:

- What if you want to enable MFA-Delete on buckets not owned by root? Do you have to contact the root user in order to do this as the bucket owner?

- This [AWS blog](https://aws.amazon.com/blogs/security/securing-access-to-aws-using-mfa-part-3/) post states that "you cannot make version DELETE actions with MFA using IAM user credentials. You must use your root AWS account.". Does this mean that other authorized users can only enable/disable versioning on the bucket with MFA and cannot delete objects/versions/delete-markers (And only the root user can)?

- In case users are allowed to delete objects on an MFA-Delete enabled bucket: If a bucket policy states that S3 API actions are MFA protected through policy conditions "aws:MultiFactorAuthPresent" and "aws:MultiFactorAuthAge", Do said users still require the mfa header (or the --mfa flag when using the CLI) when performing versioning state changes or object deletion on the bucket?

Thanks to anyone who can shine a light on this.
## [6][Why can't my lambda invoke another lambda.](https://www.reddit.com/r/aws/comments/fc1jc9/why_cant_my_lambda_invoke_another_lambda/)
- url: https://www.reddit.com/r/aws/comments/fc1jc9/why_cant_my_lambda_invoke_another_lambda/
---
I understand this is an anti-pattern. But I'm just doing a small project so I have no need for scaling etc.

I have two (python) lambdas, `watcher` and `registrar`.

Under certain conditions I want `watcher` to invoke `registrar`.

    invoke_lambda = boto3.client("lambda")
    response = invoke_lambda.invoke(
        FunctionName="registrar",
        InvocationType="Event",
        LogType="Tail",
        Payload=&lt;some payload&gt;),
    )

But I get an error:

    [ERROR] ClientError: An error occurred (AccessDeniedException) when calling the Invoke operation: User: arn:aws:sts::1234567890:assumed-role/LambdaExecutionRole/watcher is not authorized to perform: lambda:InvokeFunction on resource: arn:aws:lambda:eu-west-2:1234567890:function:registrar

But as far as I can tell, I have the correct permissions (role document):

    {
    "permissionsBoundary": {},
    "roleName": "LambdaExecutionRole",
    "policies": [
        {
        "document": {
            "Version": "2012-10-17",
            "Statement": [
            {
                "Sid": "",
                "Effect": "Allow",
                "Action": "lambda:InvokeFunction",
                "Resource": "arn:aws:lambda:eu-west-2:1234567890:function:registrar"
            }
            ]
        },
        "name": "lambda-invoke-policy",
        "id": "&lt;obfuscated&gt;",
        "type": "managed",
        "arn": "arn:aws:iam::1234567890:policy/lambda-invoke-policy"
        },
        {
        "document": {
            "Version": "2012-10-17",
            "Statement": [
            {
                "Sid": "",
                "Effect": "Allow",
                "Action": [
                "logs:PutLogEvents",
                "logs:CreateLogStream",
                "logs:CreateLogGroup"
                ],
                "Resource": "arn:aws:logs:*:*:*"
            }
            ]
        },
        "name": "lambda-cloudwatch-policy",
        "id": "&lt;obfuscated&gt;",
        "type": "managed",
        "arn": "arn:aws:iam::1234567890:policy/lambda-cloudwatch-policy"
        },
        {
        "document": {
            "Version": "2012-10-17",
            "Statement": [
            {
                "Sid": "",
                "Effect": "Allow",
                "Action": "dynamodb:Query",
                "Resource": "arn:aws:dynamodb:eu-west-2:1234567890:table/Matches/index/RegDate-index"
            }
            ]
        },
        "name": "lambda-dynamodb-policy",
        "id": "&lt;obfuscated&gt;",
        "type": "managed",
        "arn": "arn:aws:iam::1234567890:policy/lambda-dynamodb-policy"
        }
    ],
    "trustedEntities": [
        "lambda.amazonaws.com"
    ]
    }

What am I missing/doing wrong? :(


EDIT: 11 hours later...

I changed the resource for `lambda:InvokeFunction` to "*" and it started working.

I then changed it back to the explicit ARN and it was still working...

FYI: I'm using Terraform.
## [7][AWS Lambda Downloading CSV?](https://www.reddit.com/r/aws/comments/fc7npa/aws_lambda_downloading_csv/)
- url: https://www.reddit.com/r/aws/comments/fc7npa/aws_lambda_downloading_csv/
---
I have a simple AWS Lambda doing two calls - One to SWAPI (Star Wars API) and one to the NASDAQ website. Both work locally, but in AWS Lambda, only the SWAPI call resolves and the NASDAQ call hangs until the Lambda function times out. Is it possible that the NASDAQ website is blocking the call? How can I confirm this? (Note: no errors show up when running, just hanging).

Here is the code:

    import axios from "axios";

    // If this is to be invoked directly, we can create our own event
    export interface Event {
      name: string;
    }

    const main = async (): Promise&lt;string&gt; =&gt; {
      try {
        console.log(
          (
            await axios({
              url: "https://swapi.co/api/people/2",
              method: "GET",
              responseType: "blob"
            })
          ).data
        );

        console.log(
          (
            await axios({
              url:
                "https://old.nasdaq.com/screening/companies-by-name.aspx?letter=0&amp;exchange=nasdaq&amp;render=download",
              method: "GET",
              responseType: "blob"
            })
          ).data
        );

        return "Success";
      } catch (e) {
        return `Failed + ${e}`;
      }
    };

    export { main };
## [8][Backing up DynamoDb](https://www.reddit.com/r/aws/comments/fcavih/backing_up_dynamodb/)
- url: https://www.reddit.com/r/aws/comments/fcavih/backing_up_dynamodb/
---
Is there any way in which I can download the whole dynamo DB data locally onto my system and then restore it latter.

Yes, I know Dynamo has backup option and you can restore that in any region. But I need a local copy of it.
## [9][Help with updating dynamoDB](https://www.reddit.com/r/aws/comments/fc82lv/help_with_updating_dynamodb/)
- url: https://www.reddit.com/r/aws/comments/fc82lv/help_with_updating_dynamodb/
---
Here is my code:  


    "payload": {
        "TableName": "Test1",
        "Key": {
          "ID": "aosdjndidj"
        },
    
        "ExpressionAttributeNames": {
          "#R": "Region"
        },
        "UpdateExpression": "Set #R = :r",
        "ExpressionAttributeValues": {
          ":r": 2
        },
        "ReturnValues": "UPDATED_NEW"
      }
    
        dynamo.update(payload, callback);

It gives the error: 

    "errorType": "ValidationException",
      "errorMessage": "The provided key element does not match the schema"

Here is the database:  


    ID, Region, Name

Let me know what is wrong
## [10][What's the best way to deploy Docker container that runs some shell scripts to ECS ?](https://www.reddit.com/r/aws/comments/fc9xwd/whats_the_best_way_to_deploy_docker_container/)
- url: https://www.reddit.com/r/aws/comments/fc9xwd/whats_the_best_way_to_deploy_docker_container/
---
I am new to ecs. What I am looking is a simple way to run by Docker container on ecs. My Container doesn't run any application, it has some shell scripts to run in crontab and some adhoc tasks that I need to run once in a while using \`docker exec &lt;container\_id&gt;\`  . Could you help how CONTAINER DEFINITION and TASK DEFINITION should look like/
