# aws
## [1][Building an AI-powered Battlesnake with reinforcement learning on Amazon SageMaker | Amazon Web Services](https://www.reddit.com/r/aws/comments/fi4pbt/building_an_aipowered_battlesnake_with/)
- url: https://aws.amazon.com/blogs/machine-learning/building-an-ai-powered-battlesnake-with-reinforcement-learning-on-amazon-sagemaker/
---

## [2][Transfert files directly from server to S3 Bucket](https://www.reddit.com/r/aws/comments/fih9a2/transfert_files_directly_from_server_to_s3_bucket/)
- url: https://www.reddit.com/r/aws/comments/fih9a2/transfert_files_directly_from_server_to_s3_bucket/
---
Hello guys,

I’m using a service who provide many .gz files and I want to transfer them to my S3 bucket. The client update these files every day. I have to stay in sync with his server and my S3 Bucket.

The idea is to create an Lambda  (called every day) who would be in charge to transfer these .gz files directly from his server to my Bucket.

I’ve no idea how to do that.

They provide an API with all endpoints where these files are stored.

`{`  
`fileName: "xxxxx",`  
`url: "https://domaine.com/file.gz"`  
`}`

**Do you have articles or documentation about this use case?**
## [3][How is AWS helping your organization with COVID-19 response?](https://www.reddit.com/r/aws/comments/fifwbp/how_is_aws_helping_your_organization_with_covid19/)
- url: https://www.reddit.com/r/aws/comments/fifwbp/how_is_aws_helping_your_organization_with_covid19/
---
Just an opportunity to post new things that your organization is implementing to help your business cope with the changes taking place in response with COVID-19.  Are you:

* Deploying more VPN capacity?
* Setting up more virtual desktops?
* Changing your telephony setup?
## [4][AWS reMars conference cancelled](https://www.reddit.com/r/aws/comments/fi9v6l/aws_remars_conference_cancelled/)
- url: https://www.reddit.com/r/aws/comments/fi9v6l/aws_remars_conference_cancelled/
---
Not surprised ... but is it a bit early?
## [5][I am getting a 503 error...from a maxed S3 limit?](https://www.reddit.com/r/aws/comments/fi7puf/i_am_getting_a_503_errorfrom_a_maxed_s3_limit/)
- url: https://www.reddit.com/r/aws/comments/fi7puf/i_am_getting_a_503_errorfrom_a_maxed_s3_limit/
---
My site is getting a 503 "Service Unavailable

The server is temporarily unable to service your request due to maintenance downtime or capacity problems. Please try again later."

I am on the free tier and I knew S3 put requests were approaching the 2,000 max.

In fact I went and checked and it is 100%.

But with no bad credit...would they just stop the service rather than let me pay the extra?

How can I allow payments? Or is there another reason? Am i in a sandbox too?
## [6][Weird keyboard issue with AWS client](https://www.reddit.com/r/aws/comments/fia1q7/weird_keyboard_issue_with_aws_client/)
- url: https://www.reddit.com/r/aws/comments/fia1q7/weird_keyboard_issue_with_aws_client/
---
Hi, I hope this is the right forum. I use AWS with Linux, take a look at the keyboard layout. It is hosted in Ireland. Notice how the \~ and the @ are at near Enter. Any ideas here, is this a UK layout? This is persistent after I login. I am in the US.

https://preview.redd.it/2wx3acn8ajm41.png?width=1279&amp;format=png&amp;auto=webp&amp;s=de5ce2773658c91229235c19c8250011e7b7f1a5
## [7][How to find out which instances are in systems manager and which ones are not](https://www.reddit.com/r/aws/comments/fiaqqf/how_to_find_out_which_instances_are_in_systems/)
- url: https://www.reddit.com/r/aws/comments/fiaqqf/how_to_find_out_which_instances_are_in_systems/
---
I talked to support earlier today and they told me I could find out any many had the SSM agent installed using AWS Powershell Tools. I was able to pull down the instances that are in SSM vs all my instances. Any way to find out which ones DO NOT have the SSM agent installed? Support said there is no such methodology of doing this an easy way.
## [8][Howto remove Public Ip of instance created with “Auto-assign Public IP”?](https://www.reddit.com/r/aws/comments/fhy5x7/howto_remove_public_ip_of_instance_created_with/)
- url: https://www.reddit.com/r/aws/comments/fhy5x7/howto_remove_public_ip_of_instance_created_with/
---
I’m bringing env to private but I don’t find a way to remove public IP address of an instance once I had enabled vpn for private access to VPC. 

I’m using default vpc. 

So howto?
## [9][Sending multiple emails via SES in a lambda](https://www.reddit.com/r/aws/comments/fi5pp9/sending_multiple_emails_via_ses_in_a_lambda/)
- url: https://www.reddit.com/r/aws/comments/fi5pp9/sending_multiple_emails_via_ses_in_a_lambda/
---
I'll try and keep this quick, can add in code if people need more help understanding.

**TLDR**

Having issues usings SNS/SES inside a loop in a lambda

\- need to send email to multiple "sets" of emails all from one API call

\- My solution either doesn't send all SNS topics OR emails received twice

\- Using NodeJS

Currently developing APIs for a music streaming/learning app. Some features of the app are that you can create a "group" to share your songs to. Sometimes a song may be shared to multiple groups at once. Example, I want to share Song A to Group 1 and Group 2.

The request body would be something like this :

    {"songCuid":"xxx","groupCuids":["xxx","xxx"]}

What I have to do is

1. Updated the DB (RDS Postgresql) to reflect the new song in the groups
2. Retrieve members of each group
3. Send an email to each group (all members of each group)

&amp;#x200B;

The email will be something like

"A new song "Song A" has been shared to your group "**Group 1**", click this link to listen."

Therefore each "group" will need it's own email for it's members.

To get this working, I tried doing :

"FOREACH GROUP"

"get their members emails"

"send those with the groupName in an SNS topic"

Go to next group and do the same.

&amp;#x200B;

With various attempts, the lambda would finish before all SNS topics were sent, or Emails would be sent twice.

I can't seem to find a simple-ish solution to sending emails or using SES inside a loop and I just don't know how else I could get this to work.

&amp;#x200B;

Thanks in advance for reading.

&amp;#x200B;

**EDIT :** [**Repl.it**](https://Repl.it) link - [https://repl.it/repls/BeautifulBreakableValue](https://repl.it/repls/BeautifulBreakableValue)

index.js contains the main lambda - shareWithGroup, which uses publishSNS()

song-share-email.js contains the lambda that is triggered by the SNS message and sends the email
## [10][Can a lightsail load balancer handle websocket connections?](https://www.reddit.com/r/aws/comments/fhxg15/can_a_lightsail_load_balancer_handle_websocket/)
- url: https://www.reddit.com/r/aws/comments/fhxg15/can_a_lightsail_load_balancer_handle_websocket/
---
I have a high bandwidth websocket application that isnt feasible for me to run on ordinary EC2 due to the bandwidth costs being too high - this app does about 15TB per month. 

I'm looking at Lightsail that seems a good fit - an 8gb node comes with 5TB of traffic, so three of those nodes would sort me out. I understand Lightsail throttles sustained CPU usage but it seems like an 8gb node even if throttled down to 30% will still be fine for me as websockets aren't CPU intensive.

What I need to understand is:

1. Is the lightsail loadbalancer compatible with secure websockets (wss)?
2. Are there hidden charges I need to worry about? In addition to running Lightsail nodes for the websocket app, the main app itself will be hosted on EC2 nodes as the main app is highly demanding on CPU. Therefore I need the lightsail nodes to be able to access Redis and RDBS - I believe this can be done through the right VPC settings **but does it have charges?**
3. How much load can the lightsail loadbalance handle? Can lightsail in general handle big production traffic?

Thanks for any help!
