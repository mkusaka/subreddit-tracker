# aws
## [1][re:Invent registration is now open](https://www.reddit.com/r/aws/comments/jkenu3/reinvent_registration_is_now_open/)
- url: https://register.virtual.awsevents.com/
---

## [2][New – Archive and Replay Events with Amazon EventBridge](https://www.reddit.com/r/aws/comments/jp9405/new_archive_and_replay_events_with_amazon/)
- url: https://aws.amazon.com/blogs/aws/new-archive-and-replay-events-with-amazon-eventbridge/
---

## [3][Boto3 update_auto_scaling_group failing](https://www.reddit.com/r/aws/comments/jpptpz/boto3_update_auto_scaling_group_failing/)
- url: https://www.reddit.com/r/aws/comments/jpptpz/boto3_update_auto_scaling_group_failing/
---
Trying to update ASGs with new launchtemplate / lt version using update\_auto\_scaling\_group (Python3.7 in Lambda function) but get the following error:

"An error occurred (AccessDenied) when calling the UpdateAutoScalingGroup operation: You are not authorized to use launch template"

Cloudtrail shows:

"errorCode": "AccessDenied","errorMessage": "An unknown error occurred",

The lambda role has EC2FullAdmin assigned to it, which [according to the documentation is sufficient](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-launch-template-permissions.html) for accessing autoscaling and launchtemplates. I've also added AutoScalingFullAccess (though EC2FullAdmin includes autoscaling:\*).

Code (based on boto3 documentation)

`response = asg_client.update_auto_scaling_group(`

`AutoScalingGroupName=ASGName,`

`LaunchTemplate={`

`'LaunchTemplateId': 'lt-123abc',`

`'Version': '15'`

`}`

`)`

Strangely, if I instead use MixedInstancesPolicy, it updates the ASG with the new launch template version without an issue, implying access is not a problem.

`response = asg_client.update_auto_scaling_group(`

`AutoScalingGroupName=ASGName,`

`MixedInstancesPolicy = {`

`"LaunchTemplate":{`

`"LaunchTemplateSpecification":{`

`"LaunchTemplateId": 'lt-123abc',`

`"Version": '15'`

`},`

`"Overrides": CurrentMixedInstancesPolicy['LaunchTemplate']['Overrides']`

`}`

`}`

`)`

What am I missing here?

&amp;#x200B;

Edit: So I granted the lambda role AdministratorAccess as a test, the first code example then worked. So AmazonEC2FullAccess / AutoScalingFullAccess is not enough. CloudTrail is not giving me sufficient info to tell what API permission I am missing...
## [4][Cognito REST API](https://www.reddit.com/r/aws/comments/jpmffo/cognito_rest_api/)
- url: https://www.reddit.com/r/aws/comments/jpmffo/cognito_rest_api/
---
Hey guys. I’ve been playing with Amplify and I’m currently waiting for SES to be approved for my cognito MFA. 

Long story short, I’d like for users who sign up have the opportunity to CRUD through the REST API, but I’d like for guests to only be able to read. Whenever I set up the amplify API, I can only allow the cognito users to have permissions, but not guests. What’s the best way to do this?
## [5][How can I get the latest added key in s3 without having to iterate through it's contents?](https://www.reddit.com/r/aws/comments/jpne0d/how_can_i_get_the_latest_added_key_in_s3_without/)
- url: https://www.reddit.com/r/aws/comments/jpne0d/how_can_i_get_the_latest_added_key_in_s3_without/
---
I have a lambda function that adds folders and files to s3. Is there any way for a second lambda function to get the latest added key to s3 without iterating through the s3 contents or listening to s3 events.
## [6][ec2: How to install/configure pfsense to work properly?](https://www.reddit.com/r/aws/comments/jpmvuw/ec2_how_to_installconfigure_pfsense_to_work/)
- url: https://www.reddit.com/r/aws/comments/jpmvuw/ec2_how_to_installconfigure_pfsense_to_work/
---
I have 2 ec2 instances and they both have elastic addresses. I have 2 VPC on the [10.0.0.0](https://10.10.10.0) range and [172.31.0.0](https://172.31.0.0)

Any idea what I have to do when launching a pfsense instance? 

Thanks very much
## [7][In the Works – AWS Region in Hyderabad, India](https://www.reddit.com/r/aws/comments/joz9nc/in_the_works_aws_region_in_hyderabad_india/)
- url: https://aws.amazon.com/blogs/aws/in-the-works-aws-region-in-hyderabad-india/
---

## [8][Starting the conversation from Lex Chatbot?](https://www.reddit.com/r/aws/comments/jpcr3g/starting_the_conversation_from_lex_chatbot/)
- url: https://www.reddit.com/r/aws/comments/jpcr3g/starting_the_conversation_from_lex_chatbot/
---
I’m trying to integrate Lex onto my website, which uses AWS indirectly and ultimately deployed from Heroku.  I just had a few questions.

- Will deploying the bot on my website be the same for other users who use just AWS to deploy the website? 
- is there a way for Lex to initiate the conversation when the user lands on the page, such as a welcome message?
## [9][Amazon S3 Website won't load quickly and in-page links don't work, yet things load from Google](https://www.reddit.com/r/aws/comments/jpk0e0/amazon_s3_website_wont_load_quickly_and_inpage/)
- url: https://www.reddit.com/r/aws/comments/jpk0e0/amazon_s3_website_wont_load_quickly_and_inpage/
---
I'm very new to Amazon S3, and have been having trouble with my website that I host on S3 loading quickly. When I type the address of the website into the address bar, the page doesn't load until I manually select the address and hit enter again. The other issue I'm having is that the in-page links I have (essentially, when I click the About Me button, I want the page to scroll down to the About Me section) don't load when I select them. The page begins to load but never finishes. I also can't open my resume from the website. However, when I search the website in the Google search engine and click on it from there, it loads right away. Furthermore, I can select the About Me, Projects, Contact Me, and Resume parts of my page from Google, and when I do so, they load right away. I've had this problem in both Safari and Google Chrome, so I don't think it's a browser issue. Any help would be appreciated, and let me know if you need more information!

&amp;#x200B;

&amp;#x200B;

https://preview.redd.it/dcblb5bxpqx51.png?width=1236&amp;format=png&amp;auto=webp&amp;s=28a51375046c8b63bf5d96ba951f9e120bda3303
## [10][ec2 instances "freezing"](https://www.reddit.com/r/aws/comments/jp7uef/ec2_instances_freezing/)
- url: https://www.reddit.com/r/aws/comments/jp7uef/ec2_instances_freezing/
---
I've had a few instances that i lose the ability to SSH or RDP into them until i reboot them, but i can still ping them.

what causes this, or am i missing something?
## [11][Can I use python to ssh into an EC2 instance to place files from my local pc?](https://www.reddit.com/r/aws/comments/jp7vgx/can_i_use_python_to_ssh_into_an_ec2_instance_to/)
- url: https://www.reddit.com/r/aws/comments/jp7vgx/can_i_use_python_to_ssh_into_an_ec2_instance_to/
---
Disclaimer: I couldn't find any other sub to post this to other than here, so well:

I have an AWS EC2 instance, and I want to ssh into it and place some  of my local files on the server using a python script. I've heard that  you can use a module called paramiko  
 which is great, but the EC2 instance gives you a key pair (.pem  
), not a password, and I don't know how to login using that instead of a password.

Any help will be greatly appreciated.
