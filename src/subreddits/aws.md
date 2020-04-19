# aws
## [1][Let's do AWS whiteboard sessions](https://www.reddit.com/r/aws/comments/g3d6yg/lets_do_aws_whiteboard_sessions/)
- url: https://www.reddit.com/r/aws/comments/g3d6yg/lets_do_aws_whiteboard_sessions/
---
Hey guys if you're interested in Tech subscribe to my YouTube channel. I'll be talking about tech, in general, with a huge concentration on AWS. I already have a couple of videos there. 

I will be releasing a video every week so if you have a topic you'd like me to dive-deep into let me know. I'd also be interested in having tech discussions with field professionals -- I can whiteboard and architect solutions based on your requirements. Subscribe and reach out! 

[https://www.youtube.com/channel/UCBl-ENwdTlUsLY05yGgXyxw](https://www.youtube.com/channel/UCBl-ENwdTlUsLY05yGgXyxw)

&amp;#x200B;

Here's one I did recently on the topic of Serverless Photo Recognition:

[https://www.youtube.com/watch?v=GIdJz7VnP58&amp;t=259s](https://www.youtube.com/watch?v=GIdJz7VnP58&amp;t=259s)

&amp;#x200B;

This isn't sponsored by AWS and all of the opinions are my own.
## [2][Lambda Logs to ANYTHING ELSE but CloudWatch](https://www.reddit.com/r/aws/comments/g3k1zr/lambda_logs_to_anything_else_but_cloudwatch/)
- url: https://www.reddit.com/r/aws/comments/g3k1zr/lambda_logs_to_anything_else_but_cloudwatch/
---
Hi guys,

I don't know about you but CloudWatch PutLog events are f**king killing me.

I have highly-intensive (and highly-invoked) Lambda functions running for a Serverless SaaS platform that does load testing (https://rungutan.com).

While the platform is still in beta, and I'm just testing it myself, I realized that my AWS Bill is going wild for these type of events, sometimes amounting to about 300 USD/day when doing some simple tests.

I've realized that there are 2 ways that this is happening:
1) StepFunctions EXPRESS puts logs into CloudWatch
2) Lambda functions itself do that

I've stopped the StepFunctions from putting logs into CloudWatch, but I STILL COULD use the logs from Lambda themselves in order to be able to push them into Thundra (https://www.thundra.io/) so that I can analyse them properly on a per-invocation type of thing.

So my question is:

-&gt;

Have you guys found any other way of pushing logs from Lambda to ANYTHING ELSE besides CloudWatch ?

I've even considered writing a Lambda function to process logs (by simply being invoked) and push them to somewhere else (aka ElasticSearch) but that makes me lose the nice logs-per-invocation logic from Thundra.

Follow-up question:

-&gt;

Is there maybe a way to limit those default log lines that Lambda pushes?
I'm talking about "START RequestId: adae3789-2dd8-4c6a-8c05-3c4e17d26d2e Version: $LATEST" and "END RequestId: adae3789-2dd8-4c6a-8c05-3c4e17d26d2e Version: $LATEST"
## [3][SAM/Cloudformation template syntax - multiple Authorizers](https://www.reddit.com/r/aws/comments/g3luo6/samcloudformation_template_syntax_multiple/)
- url: https://www.reddit.com/r/aws/comments/g3luo6/samcloudformation_template_syntax_multiple/
---
Hi AWS!  
I am currently using an AWS SAM template to provision a serverless API.  


Using paramters/Conditions I would like to modify the Authorizers on the `AWS::Serverless::Api`  


      MyApi:
        Type: AWS::Serverless::Api
        Properties:
          StageName: Prod
          Auth:
            Authorizers:
    # if MyCondition is true
              MyCognitoAuthorizer1:
                UserPoolArn: !GetAtt MyUserPool1.Arn
    # else
              MyCognitoAuthorizer2:
                UserPoolArn: !GetAtt MyUserPool2.Arn
              MyCognitoAuthorizer3:
                UserPoolArn: !GetAtt MyUserPool3.Arn
    # fi

I have tried using the `Condition` keyword on the Authorizer object as you would on a `Resource` but it takes no effect.

I cant find a way to get the template to achieve this outcome - does anyone have any advice?
## [4][OS storage type for t3a.nano](https://www.reddit.com/r/aws/comments/g3k8rr/os_storage_type_for_t3anano/)
- url: https://www.reddit.com/r/aws/comments/g3k8rr/os_storage_type_for_t3anano/
---
I'm interested in the EC2 t3a.nano instance type but can't find anywhere what kind of storage (and how much) this instance type comes with. So I'm talking about the default (included in the price) storage that the OS is installed on.

Does anyone know if it's a HDD or SSD and if it's local or EBS?
## [5][How to best architect a live video streaming platform using MediaLive, MediaPackage, and CloudFront?](https://www.reddit.com/r/aws/comments/g36gup/how_to_best_architect_a_live_video_streaming/)
- url: https://www.reddit.com/r/aws/comments/g36gup/how_to_best_architect_a_live_video_streaming/
---
Hi there,

I am building a live video streaming platform (users can push video streams via RTMP). I have gone through the AWS Live Video Streaming solution guide ([https://aws.amazon.com/solutions/live-streaming-on-aws/](https://aws.amazon.com/solutions/live-streaming-on-aws/)) and the CloudFormation template. I understand how inputs and channels work on MediaLive and I'm planning to use MediaPackage and CloudFront to distribute the video feed (HLS).

On my platform, when people create a new stream, I want to provide them with an RTMP URL and stream key (will be on MediaLive) to push their feed. Should I programmatically (via i.e. JS SDK) create a new input and channel for every user stream? And then should I create a new channel and endpoint for MediaPackage and CloudFront as well? Is this the best/recommended approach?

I understand that there is a minimal cost incurred by idle MediaLive inputs/channels (is there for MediaPackage as well?), and I know that I could start the MediaLive channels only when/as needed. However, I know that starting a channel takes time, and ideally, I would like my users to be able to start pushing video content right away after they set up their stream. How long does it take for a channel to start? And is there any event I can subscribe to (or create on SNS) to know when the channel is ready?

I can also see that there are specific quotas and limits ([https://docs.aws.amazon.com/medialive/latest/ug/limits.html](https://docs.aws.amazon.com/medialive/latest/ug/limits.html)). In my case, I might need to have tens if not hundreds of channels (if I do need a channel for every user indeed). Can I just apply to lift the quotas, same way it works for SES for example?

Last but not least, I see that the params object for the MediaLive createChannel function (JS: [https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/MediaLive.html#createChannel-property](https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/MediaLive.html#createChannel-property)) is quite extensive. Is there a shorter template I can use that contains the most important parameters?

Does my approach overall make sense or is there a better way (from an operational and cost perspective) to achieve what I have in mind?

&amp;#x200B;

Thanks in advance!
## [6][Python HTTP server on EC2 gives empty response](https://www.reddit.com/r/aws/comments/g3n8sg/python_http_server_on_ec2_gives_empty_response/)
- url: https://www.reddit.com/r/aws/comments/g3n8sg/python_http_server_on_ec2_gives_empty_response/
---
I have created a Python aiohttp server that runs on 0.0.0.0:80. It only has one single page /hook which will show basic Hello text.

When I run the server on my laptop and check with my browser it works.

Then I uploaded my code to my EC2 instance and configured the Security Group such that it allows all sources (0.0.0.0/0) inbound HTTP.

Now when I type &lt;instance public ip&gt;/hook, my browser says it has given an empty response.
Running curl http://localhost/hook when SSH into the instance also gives this empty response.

I don't think my server is running in HTTPS as I didn't configure the Python code to do so.
## [7][Stopped lightsail instances still incur charges?](https://www.reddit.com/r/aws/comments/g37d40/stopped_lightsail_instances_still_incur_charges/)
- url: https://www.reddit.com/r/aws/comments/g37d40/stopped_lightsail_instances_still_incur_charges/
---
If I understand this correctly. stopped lightsail instances are still charged as if they are active. Why would you stop a lightsail instance if they still get charged?

Edit: I’m dumb
## [8][Do you need any help with Devops work ?](https://www.reddit.com/r/aws/comments/g3m6tj/do_you_need_any_help_with_devops_work/)
- url: https://www.reddit.com/r/aws/comments/g3m6tj/do_you_need_any_help_with_devops_work/
---
Hey Guys

Hope you are doing good !!

Due to this Pandemic situation, we have a lot of free time on our hands.

If anybody needs any kind of help in Devops, Please DM me. We are a team of 2.

We can help you out  with below mentioned technologies :

Cloud : AWS &amp; Google Cloud 

Iac  : Terraform

Configuration Management   : Ansible

Container : Docker

Container Orchestration : Kubernetes

Monitoring : Prometheus &amp; Grafana

Centralized logging : ELK stack

Databases : MongoDB, Mysql
## [9][regional data transfer - in/out/between EC2 AZs or using elastic IPs or ELB](https://www.reddit.com/r/aws/comments/g3hjq7/regional_data_transfer_inoutbetween_ec2_azs_or/)
- url: https://www.reddit.com/r/aws/comments/g3hjq7/regional_data_transfer_inoutbetween_ec2_azs_or/
---
Hello. I was looking at our AWS bill and noticed 500GB of

" regional data transfer - in/out/between EC2 AZs or using elastic IPs or ELB "

I pretty much figured out that we have two servers talking to each other using their public DNS names that resolve to their elastic IPs.

I had two questions

1) Would the simple fix be to have their public dns addresses resolve to internal IPs? So on the servers themselves, if they ping [example.domain.com](https://example.domain.com) it would resolve to internal IP instead of going out to a DNS server and finding the real public IP?

I could do this either by adding the A reconds into the DNS servers the EC2 servers point to, or by manually adding hosts records on the servers.

2) Would doing this give us a performance gain? I'm not sure if the servers are currently going up to a router and back down each time they talk? When using the elastic IP?

So would having them talk to each other using internal IPs actually speed stuff up?

&amp;#x200B;

Thanks!
## [10][Ecommerce website architecture](https://www.reddit.com/r/aws/comments/g3f24l/ecommerce_website_architecture/)
- url: https://www.reddit.com/r/aws/comments/g3f24l/ecommerce_website_architecture/
---
Hi all, just want to get your opinion about how a simple ecommerce stack should look like.

The website will be quite simple showing a few products and has a cart that allows the customers to pay for what they've added to it. The site will be created with React and in the future will add authentication and profiles.

The stack in mind was host the content in S3, connect to dynamodb to fetch product details, lambda would handle the requests for the payments and triggered through api gateway. Also, route53 to direct to my domain name.

Thoughts?
