# aws
## [1][Ingress and egress](https://www.reddit.com/r/aws/comments/f6qmhr/ingress_and_egress/)
- url: https://www.reddit.com/r/aws/comments/f6qmhr/ingress_and_egress/
---
Hi,

Any recommendations for ingress and egress controls, i realize aws don't have any native products and i don't want to create anything. some sort of appliance recommendation that works well and has url whitelisting. can be a simple single instance but needs to scale for prod use  later. integration with other aws a bonus e.g. cloudwatch

thank you
## [2][Using IAM Access Analyzer](https://www.reddit.com/r/aws/comments/f6smf3/using_iam_access_analyzer/)
- url: https://www.ibexlabs.com/using-iam-access-analyzer/
---

## [3][Is VPC the answer for this?](https://www.reddit.com/r/aws/comments/f6sg30/is_vpc_the_answer_for_this/)
- url: https://www.reddit.com/r/aws/comments/f6sg30/is_vpc_the_answer_for_this/
---
So I have an API application running on elastic beanstalk with 2 environments: one for production and one for development.  When the development environment is ready for production, I switch the URLs to change them over.  I use Google to manage my domain and have the cname pointing to the eb url that's used for production. The problem lies with the fact that I use MongoDB atlas, in which I have to whitelist I.P.s... obviously eb is elastic and switching URLs between environments changes it as well. I just started looking at VPC and was wondering if this was the answer, but it seems this only works per environment, correct? So if I switched I would have to whitelist 2 IPs in MongoDB atlas?
## [4][pros and cons of using private subnet](https://www.reddit.com/r/aws/comments/f6s34e/pros_and_cons_of_using_private_subnet/)
- url: https://www.reddit.com/r/aws/comments/f6s34e/pros_and_cons_of_using_private_subnet/
---
I tried posting this to /r/network. I hope it's OK to cross-post here.

I'm trying to improve an existing Elastic Beanstalk setup. There's a dozen instances behind a load balancer. Each instance has a public IP, although I've now closed it off such that only the load balancer can connect to them. Question: what are the benefits or disadvantages of leaving the instances inside a public subnet? I can move the instances to a private subnet, but they will no longer have Internet access. We'll have to pay for a NAT Gateway. I can't understand why it's more expensive to use private-facing instances. Maybe I'm missing something?
## [5][Lambda + DynamoDB acting weirdly](https://www.reddit.com/r/aws/comments/f6ryc9/lambda_dynamodb_acting_weirdly/)
- url: https://www.reddit.com/r/aws/comments/f6ryc9/lambda_dynamodb_acting_weirdly/
---
So here's the preface. I have an S3 bucket in which .csv files would be added frequently. Each time a .csv file is uploaded, I've set up a Lambda function which adds metadata like file names, timestamps, and a sequential file identifier to a DynamoDB table. Since DynamoDB doesn't have auto-increment functionality, I have the current number of files stored in a separate DynamoDB table. In the Lambda function, for each uploaded file, this counter value is retrieved, incremented, saved. This new incremented file count is used as the file identifier for the new file. 

Now here's the problem. When I upload multiple small .csv files, some files seem to go missing. Let's say I upload 14 files. The counter table correctly shows 14 but the table which contains the metadata only has 11 rows. I have no idea what is happening to the rest of these files.

On CloudWatch, sometimes logs don't appear at all. And in other times, the log for two separate files are shown in one log. I thought Lambda functions are individual for each file and would hence show up on different logs? For instance, in this picture showing the logs, it is only supposed to print the name of the .csv file once. But it does it twice. 

https://preview.redd.it/410uaxi0h2i41.png?width=1446&amp;format=png&amp;auto=webp&amp;s=00fa28051afc90695ee613333fc78176df18c1c9

Here's the Lambda function. I'm very new to all of this, so I apologize for any glaring mistakes.

    import boto3
    from datetime import datetime
    from uuid import uuid4
    
    s3_client = boto3.client('s3')
    dynamodb = boto3.resource('dynamodb')
    table = dynamodb.Table('s3_metadata')
    counter_table = dynamodb.Table('counter_table')
    
    def lambda_handler(event, context):
        
        csv_file_name = event['Records'][0]['s3']['object']['key'].split('/')[1]
        
        counter_table_response = counter_table.scan()
        new_csv_file_id = counter_table_response['Items'][0]['counts'] + 1
        
        timestamp = datetime.now().strftime('%Y%m%d-%H%M%S%f-') + str(uuid4())
        
        print('Current file: {}'.format(csv_file_name))
    
        #incrementing counter table
        counter_table.update_item(
            Key = {
                'file counts' : 'total'
            },
            UpdateExpression="SET counts = if_not_exists(counts, :start) + :inc",
        
            ExpressionAttributeValues={
                ':inc': 1,
                ':start': 0,
            }
        )
        
        csv_dict = {'file_id': new_csv_file_id, 'file_name': csv_file_name, 'file_timestamp' : timestamp}
        
        table.put_item(Item = csv_dict)
## [6][SSH over AWS SSM. No bastions or public-facing instances. SSH user management through IAM. No requirement to store SSH keys locally or on server.](https://www.reddit.com/r/aws/comments/f6bkju/ssh_over_aws_ssm_no_bastions_or_publicfacing/)
- url: https://github.com/elpy1/ssh-over-ssm
---

## [7][For the Love of Serverless Report - We analyzed trillions of invocations, and released a data packed report of the current state of serverless around the world. [No Signup Needed]](https://www.reddit.com/r/aws/comments/f6dsrk/for_the_love_of_serverless_report_we_analyzed/)
- url: https://newrelic.com/resources/ebooks/serverless-benchmark-report-aws-lambda-2020
---

## [8][Automate Cloudformation deployments?](https://www.reddit.com/r/aws/comments/f6pogw/automate_cloudformation_deployments/)
- url: https://www.reddit.com/r/aws/comments/f6pogw/automate_cloudformation_deployments/
---
Hi, we have few Cloudformation stacks (each with many nested stacks), currently our deployment process is to manually run the CLI commands (`cloudformation package/deploy`), this process works well but it is a real pain to manually write these commands and then set stack parameters each time, I was wondering if there's any tool to automate this process i.e. a single click-deploy type tool?
## [9][Looking to get off multiple on site domains and putting them on aws. Need some help](https://www.reddit.com/r/aws/comments/f6nmtz/looking_to_get_off_multiple_on_site_domains_and/)
- url: https://www.reddit.com/r/aws/comments/f6nmtz/looking_to_get_off_multiple_on_site_domains_and/
---
So I am one of the System Admins for my company. We have stores in 3 states and 8 domains. I have been tasked to see if we can combine them all under one main domain and host it in aws. 

We will be using active directory and possibly a nas at each site for roaming profiles. 

I have set up a free ec2 account with a Windows server and active directory, however I’m not sure how to get one of the pcs on site to connect to the test domain as a user. 

Anyone have ideas or places I can go to find more info? Most of the support I find just mentions RDP and we don’t really want to use that for the users to get on the domain.
## [10][RDS Data API Inconsistency and weird behavior](https://www.reddit.com/r/aws/comments/f6ag4c/rds_data_api_inconsistency_and_weird_behavior/)
- url: https://www.reddit.com/r/aws/comments/f6ag4c/rds_data_api_inconsistency_and_weird_behavior/
---
We are using RDS Data API for communication with the Postgres Serverless cluster from AWS Lambdas. Everything works well with small loads, like 1-10 simultaneous connections, but when load grows to 20+ simultaneous connections, RDS Data API becomes unstable and fails with weird things.

1. RDS Data API requests are mixed with their respective responses. When addressing about 50+ connections with SQL queries by RDS Data API, about 5% of requests can return results from another parallel request. For example, 50 lambdas invoke SQL query over table A and different 50 lambdas invoke SQL query over table B. As a result, some queries to table A returns result from table B and vice versa.
2. RDS Data API proxy connections can hang permanently. When 50+ lambdas have issued even one SQL query and even it has been responded with success, the following queries fails with different errors, but they can be summarized like Too many connections. Only connecting as administrator and force killing RDS connections can help.

Second issue can be avoided by wrapping requests into loop with jitter, and terminating old stuck connections from active connections. But first issue is terrible - when results from SQL requests are mixed, there is no way to workaround it from RDS-client side. What you can recommend in this situation?

~~Today we found new RDS Data API issue. If somebody tries to access RDS Data from different AWS account, he runs into error cluster &lt;ARN&gt; does not belong to the calling account id &lt;ACC-ID&gt;, which is good behavior. But after reaching RDS Data from different account, it stop works from current legitimate AWS account too with same error.~~  UPD: Last issue was erroneously recognized as RDS Data API error, but after research it has been qualified as our application per se.  But first and second issues above are still appeared.

Unfortunately, reaching RDS Data API team via [Rds-data-api-feedback@amazon.com](mailto:Rds-data-api-feedback@amazon.com)has had no result via last month. Maybe there are any AWS Data API developers here and can contact us about this really weird issue. Thanks
