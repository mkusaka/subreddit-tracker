# aws
## [1][AWS Dev Support is actually quite good](https://www.reddit.com/r/aws/comments/g1d06w/aws_dev_support_is_actually_quite_good/)
- url: https://www.reddit.com/r/aws/comments/g1d06w/aws_dev_support_is_actually_quite_good/
---
Well worth the $29/month while you're learning AWS, IMO. Their response time can take a good 10+ hours, but the responses are never template/off mark as I read about when researching. 

I'm posting an example response below so you see what I mean. I was trying to figure out why a Lambda function was erring.

[https://nimbusweb.me/box/attachment/4091230/pt90k9p2fg6v0xas2zal/HcLxl2v15YhMyawu/Screenshot\_2020-04-15\_03.15.10.jpg](https://nimbusweb.me/box/attachment/4091230/pt90k9p2fg6v0xas2zal/HcLxl2v15YhMyawu/Screenshot_2020-04-15_03.15.10.jpg)
## [2][Former Mod asking for r/aws's blessing on becoming a limited mod again](https://www.reddit.com/r/aws/comments/g179da/former_mod_asking_for_rawss_blessing_on_becoming/)
- url: https://www.reddit.com/r/aws/comments/g179da/former_mod_asking_for_rawss_blessing_on_becoming/
---
Hello /r/aws 

I was a mod here for a few years and (humble brag) I helped implement many things: post flair, chat room, AMAs with AWS, re:Invent "event" posts and collections, weekly "What are you working on/What do you need help with/Cert Discussons/Favorite AWS Tip discussions etc., subreddit logo and more. 

I took a job at AWS a few months ago and decided to step down in the spirit of transparency while I adjusted. ([https://www.reddit.com/r/aws/comments/ec1t04/fyi\_in\_the\_spirit\_of\_transparency\_stepping\_down/](https://www.reddit.com/r/aws/comments/ec1t04/fyi_in_the_spirit_of_transparency_stepping_down/)) - 

I'd like to ask the community if it would be alright if I became a limited mod again. I intend to use mod privs to do things like pin posts and organize AMAs where appropriate. My job at AWS is NOT in marketing, I am in a technical role. I don't want to get involved with approving posts/comments unless they are egregious.

What are you thoughts?
## [3][EC2 Windows Server 2019 Nvida AMI problems(?)](https://www.reddit.com/r/aws/comments/g1onp2/ec2_windows_server_2019_nvida_ami_problems/)
- url: https://www.reddit.com/r/aws/comments/g1onp2/ec2_windows_server_2019_nvida_ami_problems/
---
Hi, I'm not familiar to remote PC and server configurations as most people are here, and please correct me if I'm doing anything wrong.

I'm trying to launch some games I installed on my EC2 machine from Nvdia's gaming PC AMI

g4dn.2xlarge (Windows Server 2019)

 found here:  [https://aws.amazon.com/marketplace/pp/NVIDIA-NVIDIA-Gaming-PC-Windows-Server-2019/B07STLTHM8#pdp-pricing](https://aws.amazon.com/marketplace/pp/NVIDIA-NVIDIA-Gaming-PC-Windows-Server-2019/B07STLTHM8#pdp-pricing) 

It says it has pre-loaded basic drivers for graphics cards and things for gaming. The first time I attempted to load a newer driver:  **Version:442.50  WHQL** (Tesla T4), and when I rebooted the machine the graphics card stopped showing up in Task Manager. 

So I got another instance, same settings. I installed steam and installed Rocket League, which runs on DX11. When launching the game it would show in the task bar :\[DX11,Cooked\], but shortly after a second the window does not appear and it terminates itself without an error. Similar things also happened to CS:GO, where it started at first, I changed the setting in the game so it would be full screen. It froze and when I tried to turn it back on it would give me a DirectXError. I'm not sure if I'm missing some drivers for them, what should I do?
## [4][Lambda HTTP trigger?](https://www.reddit.com/r/aws/comments/g1psf1/lambda_http_trigger/)
- url: https://www.reddit.com/r/aws/comments/g1psf1/lambda_http_trigger/
---
With Azure Functions, you can trigger it via a straight HTTP call, without having to set up (and pay for) an API Gateway.   Why can't we do that with Lambda?
## [5][Is there a way to check service usage through the AWS API?](https://www.reddit.com/r/aws/comments/g1oq5q/is_there_a_way_to_check_service_usage_through_the/)
- url: https://www.reddit.com/r/aws/comments/g1oq5q/is_there_a_way_to_check_service_usage_through_the/
---
I am currently using Boto3 with python to call Amazon Polly TTS. I do not want to exceed the free tier limit and was wondering if I could get my current character count usage through the API?
## [6][GSuite as IdP for AWS SSO](https://www.reddit.com/r/aws/comments/g1mf3p/gsuite_as_idp_for_aws_sso/)
- url: https://www.reddit.com/r/aws/comments/g1mf3p/gsuite_as_idp_for_aws_sso/
---
Anyone had any luck setting up [AWS SSO](https://aws.amazon.com/single-sign-on/) (Not IAM, not Cognito) to use GSuite as an IdP? 

Aside from [one forum post last year](https://forums.aws.amazon.com/thread.jspa?messageID=913518) requesting it as a feature, all the results seem to be about using it through either Cognito or IAM.  Which isn't what I'm after. 

I've got the flow working where either you can do the initiation from either the IDP or SP side, but AWS just comes back with 

&gt; An unexpected error has occurred  
&gt; Please try signing in again. If the error persists, please contact your administrator


Checking in my browser, the SAML response from GSuite seems to be completely successful, so I suspect there's probably just an attribute mapping issue or something - but I can't figure out what the problem is. 

There's no Cloudwatch logs or anything that I can find, so I'm not sure where to go from here.
## [7][What is the best way to learn AWS and PostgreSQL](https://www.reddit.com/r/aws/comments/g1j5ne/what_is_the_best_way_to_learn_aws_and_postgresql/)
- url: https://www.reddit.com/r/aws/comments/g1j5ne/what_is_the_best_way_to_learn_aws_and_postgresql/
---
We have been a Microsoft /Azure shop for quite a long long time and I was told to learn AWS and PostgreSQL for a new project we are doing. I have decent experience with Azure and SQL server. I am not looking for any specific resources, but wanted to check what would be the best way to learn both together ?
## [8][Voice echo on Skype/Microsoft Teams meetings while using AWS workspaces](https://www.reddit.com/r/aws/comments/g1m88z/voice_echo_on_skypemicrosoft_teams_meetings_while/)
- url: https://www.reddit.com/r/aws/comments/g1m88z/voice_echo_on_skypemicrosoft_teams_meetings_while/
---
I have been using AWS workspaces for past two weeks due to a business exigency. While I am attending meetings on VOIP clients, other participants are constantly hearing an echo in the background. Kindly let me know if this is a known issue and if there is any solution to it.
## [9][Is there a relational database with DynamoDB type pricing?](https://www.reddit.com/r/aws/comments/g12o01/is_there_a_relational_database_with_dynamodb_type/)
- url: https://www.reddit.com/r/aws/comments/g12o01/is_there_a_relational_database_with_dynamodb_type/
---
As per question. DynamoDB is fully managed and charged by read/writes, so for a hobby application that gets used say once a day or once per week the monthly cost is virtually NIL. The closest relational database product I could find on AWS is Aurora Serverless, but that is charged per hour so relies on the database going to sleep when not in use to keep the cost down. And that in turn causes a 30 second delay for the database to restart each time the application is accessed. So is that it or is there another relational database product on AWS, which is also charged on read/writes/storage that I can look at for a hobby application that would have a sub $5/month cost?
## [10][Get-ec2InstanceMetadata IMDSv1 or IMDSv2](https://www.reddit.com/r/aws/comments/g1d502/getec2instancemetadata_imdsv1_or_imdsv2/)
- url: https://www.reddit.com/r/aws/comments/g1d502/getec2instancemetadata_imdsv1_or_imdsv2/
---
Does anyone know off hand if the powershell Get-ec2InstanceMetadata uses IMDSv1 or has transitioned to IMDSv2? I'm not able to spin up a ec2 to check the token status cloudwatch metric when running the command, so asking here.
