# aws
## [1][Longer Console Sessions](https://www.reddit.com/r/aws/comments/hxbq9i/longer_console_sessions/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/07/now-gain-longer-access-to-your-aws-resources-when-switching-roles-in-the-aws-management-console/
---

## [2][HTTP API microservice separation](https://www.reddit.com/r/aws/comments/hxlpu8/http_api_microservice_separation/)
- url: https://www.reddit.com/r/aws/comments/hxlpu8/http_api_microservice_separation/
---
How are folks managing things like resource-ownership for microservices with the new HTTP APIs? As far as I can tell it is fairly simple to allow each microservice to maintain control over its own routes/integrations/etc. The main problem is that I haven't been able to work out is how to enable route-level throttling since, as far as I can tell, it needs to be defined on the stage rather than the route itself.

The only solution I can come up with is having each service export a list of routes and their desired throttling limits, then importing all of them into a service that auto-updates the stage as needed.
## [3][How to video stream for a month for less than $10](https://www.reddit.com/r/aws/comments/hxh6yl/how_to_video_stream_for_a_month_for_less_than_10/)
- url: https://www.reddit.com/r/aws/comments/hxh6yl/how_to_video_stream_for_a_month_for_less_than_10/
---
I am trying to understand how can people video stream HD for a whole month for less than $10. I understand Netflix uses AWS but wouldnt their EC2 costs be huge for video streaming.

How is it made efficient ?
## [4][AWS new private IP address inside my bash script](https://www.reddit.com/r/aws/comments/hxk1yw/aws_new_private_ip_address_inside_my_bash_script/)
- url: https://www.reddit.com/r/aws/comments/hxk1yw/aws_new_private_ip_address_inside_my_bash_script/
---
Hello guys,

I am trying to create a GHE (GitHub Enterprise) inside AWS and once it's up and running I need to perform a few curls to install the license, settings.json file and my public key.

How can I grab the IP of this new instance during "building" time, and use this IP inside my curls commands?

packer.json file

[https://gist.github.com/rmarcandier/64786b129b0e936042961f2007879a78](https://gist.github.com/rmarcandier/64786b129b0e936042961f2007879a78)

conf.sh script

[https://gist.github.com/rmarcandier/9e412169656b7b17f154251c4c8bf338](https://gist.github.com/rmarcandier/9e412169656b7b17f154251c4c8bf338)

Thank you in advance

Regards

RG
## [5][Does a dynamodb scan return items in any particular order?](https://www.reddit.com/r/aws/comments/hxa1pr/does_a_dynamodb_scan_return_items_in_any/)
- url: https://www.reddit.com/r/aws/comments/hxa1pr/does_a_dynamodb_scan_return_items_in_any/
---
In particular say I begin scanning my table, is it guaranteed that items written to the table mid-scan will also be returned? For instance, if the primary key were integers, and it returned in sorted batches, then the first batch might return the entries 1,2,5,7,11,13 where 13 is the LastEvaluatedKey. While this is happening say that 12 is written to the table, under such an ordering some newly written entries may be missed, while others would be included. This question isn't application specific, but more out of curiosity since I don't see much written about it in the aws docs.
## [6][My Lambda Function is not inforcing CORS.](https://www.reddit.com/r/aws/comments/hxjot6/my_lambda_function_is_not_inforcing_cors/)
- url: https://www.reddit.com/r/aws/comments/hxjot6/my_lambda_function_is_not_inforcing_cors/
---
I have already put the CORS headers and body in my Lambda function code however it doesn't seem to work, I still have to use a CORS remover/blocker to view my Javascript code in the console.

Also in my Lambda function results it shows the headers,body and status code which I don't need for my script, does anyone know how to fix all of this and get my CORS to work? I just want the number to show the "357".
Do i have to add a wild card?

    {
    "statusCode": 200,
    "headers": {
    "Content-Type": "application/json",
    "Access-Control-Allow-Origin": "*",
    "Allow": "GET, OPTIONS, POST",
    "Access-Control-Allow-Methods": "GET, OPTIONS, POST",
    "Access-Control-Allow-Headers": "*"
    },
    "body": "\"357\""
    }


----------

    import boto3
    import json
    def lambda_handler(event, context):
    dynamodb = boto3.client('dynamodb')
    response = dynamodb.update_item(
    TableName='resumecounter', 
    Key={
        'url':{'S': "etc.com"}
    },
    UpdateExpression='SET visits = visits + :inc',
    ExpressionAttributeValues={
        ':inc': {'N': '1'}
    },
    ReturnValues="UPDATED_NEW"
    )

    print("UPDATING ITEM")
    response = {
    'statusCode': 200,
    'headers': {
        "Content-Type" : "application/json",
        "Access-Control-Allow-Origin" : "*",
        "Allow" : "GET, OPTIONS, POST",
        "Access-Control-Allow-Methods" : "GET, OPTIONS, POST",
        "Access-Control-Allow-Headers" : "*"
    },
    'body': json.dumps(response["Attributes"]["visits"]["N"])
    }
     return response
## [7][Is my S3 Bucket secure from public access?](https://www.reddit.com/r/aws/comments/hxj1fq/is_my_s3_bucket_secure_from_public_access/)
- url: https://www.reddit.com/r/aws/comments/hxj1fq/is_my_s3_bucket_secure_from_public_access/
---
I have my mongodb backed up to my S3 Bucket. Is my data reasonably secure? 

\- I have no CORS configurations defined. All Public Access is blocked.

Here is my S3 bucket policy: 

{

"Version": "2012-10-17",

"Id": "Policy",

"Statement": \[

{

"Sid": "statement",

"Effect": "Allow",

"Principal": {

"AWS": "arn:aws:iam::xxxxxxxxxxxx:user/s3-bucket"

},

"Action": "s3:\*",

"Resource": \[

"arn:aws:s3:::my-backups/\*",

"arn:aws:s3:::my-backups"

\]

}

\]

}

I believe I need my policies as such to allow my Iam user to run the AWS sync command. 

&amp;#x200B;

  
Let me know if all is good or its full of obvious flaws.
## [8][Api gateway returning cors error instead of token expired](https://www.reddit.com/r/aws/comments/hxitbk/api_gateway_returning_cors_error_instead_of_token/)
- url: https://www.reddit.com/r/aws/comments/hxitbk/api_gateway_returning_cors_error_instead_of_token/
---
I am working on a project that is configured with cognito for user authentication.usually after login , the APIs return responses without failure (refreshing token has not been configured yet). After the token expires , though each api call from browser is returning cors instead of token expired error. Invoking in postman gives the correct error that the incoming token has expired.please advice wat I am missing here
## [9][HTTP web proxy for browser (IP:port) - cheapest (on-demand?) AWS solution?](https://www.reddit.com/r/aws/comments/hxl168/http_web_proxy_for_browser_ipport_cheapest/)
- url: https://www.reddit.com/r/aws/comments/hxl168/http_web_proxy_for_browser_ipport_cheapest/
---
For example, to bypass geo-blocking.
## [10][Experience with choosing AWS Glue as an ETL platform](https://www.reddit.com/r/aws/comments/hx8i6s/experience_with_choosing_aws_glue_as_an_etl/)
- url: https://www.reddit.com/r/dataengineering/comments/hx8h1j/experience_with_choosing_aws_glue_as_an_etl/
---

