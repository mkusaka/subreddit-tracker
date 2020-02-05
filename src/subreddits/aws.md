# aws
## [1][Converting existing API Gateway REST API into AWS Amplify + AppSync](https://www.reddit.com/r/aws/comments/ez7x6k/converting_existing_api_gateway_rest_api_into_aws/)
- url: https://www.reddit.com/r/aws/comments/ez7x6k/converting_existing_api_gateway_rest_api_into_aws/
---
I haven't been able to find any information about anyone doing this online: I'm trying to make a react-native app using AWS Amplify. My team has built a full REST and Websocket API through API Gateway, which works perfectly with our react web app, however there doesn't seem to be any tools to convert this into an API that could be used with a mobile application. Some issues I've encountered are that:

1. AWS Amplify is really pushing for the usage of GraphQL through AppSync - and there are seemingly no tools for me to convert the many REST/Websocket endpoints I have into GraphQL
2. Amplify doesn't have any documentation about how to integrate an existing API (whether it be GraphQL or REST), only the creation of a new one. Based on the configuration files that are generated when I create a new API it seems like I could possibly modify these to access my previous API Gateway


I'm curious as to whether anyone has gone through this process before? Based on my research there are no alternative to using AWS Amplify for the creation of a mobile application on AWS (as it has absorbed AWS Mobile Hub)
## [2][AWS Shield Feedback!](https://www.reddit.com/r/aws/comments/ez213c/aws_shield_feedback/)
- url: https://www.reddit.com/r/aws/comments/ez213c/aws_shield_feedback/
---
Hi all! Looking for some feedback/thoughts/experiences with AWS and DDoS attacks/mitigation! 

Thx!
## [3][Show your Tool: Parliament](https://www.reddit.com/r/aws/comments/ez6rfe/show_your_tool_parliament/)
- url: https://cloudonaut.io/show-your-tool-parliament/
---

## [4][Opinion: ECS (Perhaps Fargate) vs Docker on EC2](https://www.reddit.com/r/aws/comments/ez19h4/opinion_ecs_perhaps_fargate_vs_docker_on_ec2/)
- url: https://www.reddit.com/r/aws/comments/ez19h4/opinion_ecs_perhaps_fargate_vs_docker_on_ec2/
---
We have grown rather organically as most startups, and it turns out we are running 90+% of our apps in containers on EC2, sometimes multiple containers per instance.  I have heard from an friend/co-worker that ECS was kind of hard to get into/PITA 3+ years ago, but wondering about thoughts now. The best way to form my own opinion is going to be build a cluster (or two), add all containers and deploy, with our restrictions on how we can load balance and from where (external only)  .  I'm not very interested in jumping into EKS at the moment, so please stick to ECS.  This is more a question about the ecosystem of ECS, possible blue/green, ease of use, security etc.  As a side question, does Fargate mitigate some peoples concerns about serverless docker?
## [5][Is the "Billing &amp; Cost Management Dashboard" the ultimate place to see how much AWS costs me?](https://www.reddit.com/r/aws/comments/eyugiv/is_the_billing_cost_management_dashboard_the/)
- url: https://www.reddit.com/r/aws/comments/eyugiv/is_the_billing_cost_management_dashboard_the/
---
Hi,

&amp;#x200B;

I am a paranoid person by nature, so I just want to be very sure that though the "Billing &amp; Cost Management Dashboard"  I can see how much AWS will charge me?

&amp;#x200B;

Let's imagine someone hacks me and starts a EC2 Instance. Will the second the EC2 instance start to cost money (after the Free tier thing), will the "Billing &amp; Cost Management Dashboard" show me that cost?

So, if I keep an eye on the "Spend Summary" in the "Billing &amp; Cost Management Dashboard" will I know how much they will charge me?

&amp;#x200B;

&amp;#x200B;

&amp;#x200B;

I am talking about this part of AWS in the picture below.

[https://i.imgur.com/MKYdF7q.png](https://i.imgur.com/MKYdF7q.png)

&amp;#x200B;

I also set up a budget and MFA with my phone

Thank you
## [6][How to setup RDS Aurora with aws-cli. Documentation and error messages are not aligned. At least I can't find the connection.](https://www.reddit.com/r/aws/comments/ez81ro/how_to_setup_rds_aurora_with_awscli_documentation/)
- url: https://www.reddit.com/r/aws/comments/ez81ro/how_to_setup_rds_aurora_with_awscli_documentation/
---
aws-cli version `aws-cli/2.0.0dev2 Python/3.7.3 Linux/4.15.0-47-generic botocore/2.0.0dev1`

`aws rds create-db-instance \`

`--db-name database \`

`--db-instance-identifier databaseid \`

`--db-instance-class db.t3.small \`

`--db-subnet-group-name subnet-0f7471222887cd5c4 \`

`--engine aurora-mysql \`

`--engine-version 5.7.23 \`

`--master-username username \`

`--master-user-password 123password \`

`--profile profile`

The error I get is `The parameter AllocatedStorage must be provided and must not be null.`

But in the doc the parameter is not applicable to Aurora: `Not applicable. Aurora cluster volumes automatically grow as the amount of data in your database increases, though you are only charged for the space that you use in an Aurora cluster volume.`

How can I create Aurora Mysql DB with CLI?
## [7][Sagemaker not using GPU even though it is deployed in GPU instance.](https://www.reddit.com/r/aws/comments/ez7ysj/sagemaker_not_using_gpu_even_though_it_is/)
- url: https://www.reddit.com/r/aws/comments/ez7ysj/sagemaker_not_using_gpu_even_though_it_is/
---
I was able to successfully deploy my object detection model into Sagemaker and get predictions from it but it turns out the GPU is not being utilized when the prediction is to be done. (In short terms, the inference is not using GPU)

I deployed my custom model using the code:

```
predictor = model.deploy(initial_instance_count=1, instance_type='ml.p2.xlarge') 
```

Then, I get predictions by:

```
results = predictor.predict(imageData)
```

This process takes around ~15+ seconds and when I check the utilization graph in the endpoint dashboard I get peaks in GPU memory but not on the process usage.

The question is also present here. 

https://stackoverflow.com/questions/59421569/sagemaker-not-using-gpu-even-though-it-is-deployed-in-gpu-instance
## [8][Security Windows Server](https://www.reddit.com/r/aws/comments/ez7tfx/security_windows_server/)
- url: https://www.reddit.com/r/aws/comments/ez7tfx/security_windows_server/
---
Hello,

  
I'm newbie on AWS and I would like create a new server windows with SQL &amp; IIS.

How secure access to this server ? Create a VM with firewall ? Or a firewall is included on AWS ?

Thanks a lot for your help .
## [9][CloudEndure failed machine conversion](https://www.reddit.com/r/aws/comments/ez6krg/cloudendure_failed_machine_conversion/)
- url: https://www.reddit.com/r/aws/comments/ez6krg/cloudendure_failed_machine_conversion/
---
Hi !

At work today, we tried to migrate a VM from VmWare to AWS. It is a Windows 2016.

Agent is installed, Data replication seems ok, everything seems fine. 

Logs do not bring us much information :

 

`04/02/2020 à 17:12:25 Job started`

`04/02/2020 à 17:12:25 Started waiting for latest snapshot`

`04/02/2020 à 17:17:27 The latest state of the machine 4209275b-9c0f-xxxx-xxxx-804389039078 could not have been fetched. Creating replica that is up-to-date to 2020-02-04 16:10:26.178717UTC instead.`

`04/02/2020 à 17:17:27 Finished waiting for latest snapshot`

`04/02/2020 à 17:17:35 Started machine conversions`

`04/02/2020 à 17:32:37 Failed machine conversions`

`04/02/2020 à 17:32:38 Failed creating a replica for volume vol-4209275b-9c0f-xxxx-xxxx-804389039078:e:0: No snapshot for volume cloning, probably a failed conversion.`

`04/02/2020 à 17:32:38 Failed creating a replica for volume vol-4209275b-9c0f-xxxx-xxxx-804389039078:c:0: No snapshot for volume cloning, probably a failed conversion.`

`04/02/2020 à 17:32:38 Failed creating a replica for volume vol-4209275b-9c0f-xxxx-xxxx-804389039078:f:0: No snapshot for volume cloning, probably a failed conversion.`

`04/02/2020 à 17:32:38 Job finished`

For information :

We have to shut down source for network reason when replication process have begin. 

We do not understand what is going on because everything should, thechnicaly, work well. 

Do you guys have any suggestion ? Already tried to reinstall from 0 the agent and create a new porject. Data replication seems fine, only the snapshots seems causing trouble.
## [10][Trying to find the status of my table in dynamoDB using Python.](https://www.reddit.com/r/aws/comments/ez53ow/trying_to_find_the_status_of_my_table_in_dynamodb/)
- url: https://www.reddit.com/r/aws/comments/ez53ow/trying_to_find_the_status_of_my_table_in_dynamodb/
---
Lets say I add a new item to my table: I want my python script to recognize that change. Is there any way I can do that? Thank you!
