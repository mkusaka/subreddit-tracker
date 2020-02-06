# aws
## [1][AWS VPC for Software Engineers](https://www.reddit.com/r/aws/comments/ezru46/aws_vpc_for_software_engineers/)
- url: https://blog.deleu.dev/aws-vpc-for-software-engineers/
---

## [2][I built a Google assistant integration to tell you about the latest AWS product announcements](https://www.reddit.com/r/aws/comments/ezl1qq/i_built_a_google_assistant_integration_to_tell/)
- url: https://www.reddit.com/r/aws/comments/ezl1qq/i_built_a_google_assistant_integration_to_tell/
---
Simply say "OK Google, talk to cloud computing news"

Should work on both iOS and Android.

It will give you an itemized list with hyperlinks about the latest product enhancements from AWS.

App details:
https://assistant.google.com/services/a/uid/0000006c6dc51de5

Source code:
https://github.com/circa10a/google-home-aws-news
## [3][Understanding Modern Cloud Architecture on AWS: A Concepts Series](https://www.reddit.com/r/aws/comments/ezcu39/understanding_modern_cloud_architecture_on_aws_a/)
- url: https://start.jcolemorrison.com/understanding-modern-cloud-architecture-on-aws-a-concepts-series/
---

## [4][Event registration webpage](https://www.reddit.com/r/aws/comments/ezseqy/event_registration_webpage/)
- url: https://www.reddit.com/r/aws/comments/ezseqy/event_registration_webpage/
---
Hey all.  I’m just starting my path into the world of cloud computing.  I volunteer for a nonprofit and every year we have a big sports weekend and we utilize jotforms that dumps all the info into a spreadsheet then once they submit it redirects to PayPal to pay for said event.  I want to transform this and utilize a public facing site that dumps into an sql dB and then can redirect for payments.  Does anyone have any resource or anything where I can get started?
## [5][For a Lambda function, is it possible to send a message to a connected client if they're connected via boto3?](https://www.reddit.com/r/aws/comments/ezrf3c/for_a_lambda_function_is_it_possible_to_send_a/)
- url: https://www.reddit.com/r/aws/comments/ezrf3c/for_a_lambda_function_is_it_possible_to_send_a/
---
Usually how I deal with sending data back to a connected client is I just use this function:

    def _send_to_connection(connection_id, data, event, domain_name, stage):
        gatewayapi = boto3.client(
            "apigatewaymanagementapi", endpoint_url="https://" + domain_name + "/" + stage
        )
        return gatewayapi.post_to_connection(
            ConnectionId=connection_id, Data=json.dumps(data).encode("utf-8")
        )

This works great for websocket connections, but how am I supposed to do this with a client that connects via something like this?:

    session = boto3.Session(profile_name="default-manager")
    lambda_client = session.client("lambda")
    
    context = json.dumps({"action": "handleMessage", "message": "Contents of the message"})
    
    response = lambda_client.invoke(
        FunctionName="my-function-manager-dev-handleCreation",
        InvocationType='RequestResponse',
        LogType='Tail',
        ClientContext=base64.b64encode(context.encode()).decode("utf8"),
        Payload=json.dumps({"statement": "my nose will now grow"}),
    )
## [6][Is anyone else having problems with the new redshift console UI?](https://www.reddit.com/r/aws/comments/ezrbg2/is_anyone_else_having_problems_with_the_new/)
- url: https://www.reddit.com/r/aws/comments/ezrbg2/is_anyone_else_having_problems_with_the_new/
---
The new Redshift console UI keeps freezing and locking up for me - in different browsers and on different machines. Is anyone else finding this?
## [7][Building an Alexa skill that reads the output of a bash script/runs a bash script](https://www.reddit.com/r/aws/comments/ezrb5e/building_an_alexa_skill_that_reads_the_output_of/)
- url: https://www.reddit.com/r/aws/comments/ezrb5e/building_an_alexa_skill_that_reads_the_output_of/
---
I'm new to AWS and I'm not sure how to integrate a simple bash script to read the temperature of my server into a voice feedback from Alexa.  


A friend of mine had done this before, I recall him using Ngrok to set a web server that Alexa could use to run the script and pipe the output? I'd like to do that too, or simply just to run a script based on on Alexa command.  


Are there any working examples or guides to setting up AWS Alexa Skills to start this? I didn't have much luck researching on my own.
## [8][Querying AWS Service Logs Using Amazon Athena](https://www.reddit.com/r/aws/comments/ezr759/querying_aws_service_logs_using_amazon_athena/)
- url: https://www.ibexlabs.com/querying-aws-service-logs-using-amazon-athena/
---

## [9][Should I be using signed urls for image hosting? [aws/S3]](https://www.reddit.com/r/aws/comments/ezqrs4/should_i_be_using_signed_urls_for_image_hosting/)
- url: /r/node/comments/ezqpiw/should_i_be_using_signed_urls_for_image_hosting/
---

## [10][Update PHP 7.3 to latest](https://www.reddit.com/r/aws/comments/ezqd0y/update_php_73_to_latest/)
- url: https://www.reddit.com/r/aws/comments/ezqd0y/update_php_73_to_latest/
---
i have an a1 instance and installed php7.3 via yum.

i can see that the latest version is 7.3.14 but when i tried to `yum update php` it doesnt update to the latest one so its stuck to 7.3.11

I know there's a remi repo but that repo doesnt have ARM packages.

so is there a way to update php to latest?
