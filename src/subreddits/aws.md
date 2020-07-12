# aws
## [1][PSA: Don't take remote exams offered by Pearson Vue (OnVue) for AWS Certifications!](https://www.reddit.com/r/aws/comments/fscq7v/psa_dont_take_remote_exams_offered_by_pearson_vue/)
- url: https://www.reddit.com/r/aws/comments/fscq7v/psa_dont_take_remote_exams_offered_by_pearson_vue/
---
I can't describe how horrible this experience was.  I am not looking forward to how much work I am going to have to do to get my money back.  This is not my first AWS certification (I have SA Pro and DevOps Pro), but is my first online exam.  The short version is: Don't take AWS exams via the Pearson Vue at home option, even if it is offered.  AWS should not be offering this option as I can attest it is a waste of time.  Ironically, AWS would have us use their services because of their high availability and scaling but apparently they don't ask their test partners to do the same!

It started off easy enough: I passed the initial 'checks' as it confirmed my internet speed, camera access, and microphone access.  I started the process 15+ minutes before my scheduled exam time.  I was able to open the app, it again verified the technical requirements passed, and I went to the next screen.  It asked for my cell phone number and texted me a link which opened a web page which requested to take my photo.  Easy enough.  I did that and then the web page went to 'Uploading and verifying photo'.  A spinning circle started spinning.  This is where my test experience ended, but not where the poor experience ends.  I tried again, and then a third time.  Same experience.  As I write this, I left it on that page and the spinning is continuing.  This screen has been spinning for no less than 45 minutes.  At 8 minutes before my scheduled exam, I tried finding the help link.  A chat window opened, and I waited, and waited, and waited.  Still waiting as I write this.  My chat window has been open for 52 minutes and still no one to help.  Every two minutes I get ' All agents are currently assisting others. Thank you for your patience.' written in the window.  OK - what next?  They make it harder to find, but I got a phone number I can call.  I tried calling that.  Busy signal.  For the next 20 minutes I called back and back, busy signal.  Finally, I got it to actually pick up, but of course no human yet.  No estimate of time to when I can be helped.  They don't even have nice elevator music to listen to.  Who knows when I will be able to talk to someone.  This has been an exceedingly poor experience.

If you value your time, please do yourself a favor and don't even attempt a online exam with Pearson.  I worked hard to prepare for this exam and rescheduled things to fit around it.  Now, I will have to do that all again.

u/jeffbarr Is this the experience AWS is hoping to get with their testing partners?  This was a waste of my time and money.  Amazon should seriously reevaluate the quality of their test partners.  I understand everyone is trying to deal with all the issues.  However, if you can't offer quality testing, then please don't offer the option at all.  It isn't respectful to people's time.  Pearson is well aware of their capacity and if it isn't up to requirements, they shouldn't be scheduling test slots.

&amp;#x200B;

*EDIT*: A few background items I didn't initially share that may be relevant for others.  For the computer, I used a fully up to date Windows 10 laptop.  The laptop itself is only about a month old and is in near pristine condition.  Other than a few applications like Office, there is barely anything installed on there yet.  I used a hard wired connection, like recommended by Pearson through the use of a usb-to-ethernet adapter.  I have Verizon FIOS (980Mbps/840Mbps) and did do a speed test way after it was apparent this would not work.  I forget the exact numbers, but I was still pulling in hundreds of Mbps in both directions, despite everyone being at home and using the USB ethernet adapater which does put a cap on my speed, but I can't see hundreds of Mbps not being sufficent by orders of magnatude.  My phone is a fully up to date pixel 3.  I tried using my wifi in my house first (connected through FIOS), and then using the phone 4G LTE connection.  I can't imagine this was caused by my end.  It seemed like Pearson's servers were jammed at that point in time.

&amp;#x200B;

*Update*: After a LONG time, I did eventually get someone to answer from Pearson.  They were nice enough and were fairly easy to understand, although there was an delay echo introduced where whatever I said was echoed a quarter to half second later which was annoying, but bearable.  I was just happy she was able to hear me.  She said she could open a trouble ticket for me, but as it was well over an hour trying to get through to any human and doubtful it was on my side, I just told her to schedule me for the next available in person appointment.  She had to cancel my appointment and then rebook it as their sub-standard system wouldn't let her reschedule an at home appointment to at a location.  Surprisingly, she said they would refund my money and rebook me.  It was painless enough, but when I asked for a reference number on the refund, all she could do is say I 'should' get an email.  Perhaps unsurprisingly, this morning I see a fully posted charge for the rescheduled exam, but no sign of a refund.  Sigh.  I will give it a few days and then start this process over.

For what its worth, people should IGNORE the advice that the web chat is the fastest way of getting help.  Find the phone number and dial and re-dial it as fast as you can when you get a busy signal.  Despite the fact that it took 20+ minutes to get the number to pickup (and was 'waiting' 20 minutes less from the phones point of view) I got a faster response from someone on the phone.  Web based chat never picked up, even though I left it running during my entire phone conversation.

*Update #2*: It took two more days than the charge, but the refund did show up in the correct amount on my credit card.  I am actually quite surprised.
## [2][Analysing up to 100k messages per second, what approach to take?](https://www.reddit.com/r/aws/comments/hpqf22/analysing_up_to_100k_messages_per_second_what/)
- url: https://www.reddit.com/r/aws/comments/hpqf22/analysing_up_to_100k_messages_per_second_what/
---
So I got the nice challenge to process messages from up to 100k users per second on a web page and the results need to be calculated in (near)realtime. The data coming in is structured and consists of a string and and amount and what I need to achieve is to get the combined sum of amounts for the unique strings (ie; some kind of voting system). So input looks something like:

TEAM\_A;10  
TEAM\_A:2  
TEAM\_B:1  
TEAM\_B:12  
TEAM\_A:4  
TEAM\_B:5  


output must be something like:  
TEAM\_A:16  
TEAM\_B:18  


and this must basically process every second. I've created 2 proof of concepts and both are using IoT Core over MQTT to send results to AWS; this works great. To process the incoming data I have 2 proof of concepts right now that work (not tested with 100k though!):  


A) Send IoT data to SQS and attach a Lambda. The Lambda will read up to 10 messages and store the data in ElastiCache Redis. While this works great for \~10k users, I don't think it will scale to 100k because it will require way too many concurrent Lambdas.

B) Send IoT data to Kinesis Firehose and then attach a Kinesis Analytics stream, the results will be pushed every second to a Lambda. This works really well however I'm reading Firehose has a 5k message/second limit which seems really low for a service called "Firehose". Also; if it could be adjusted to 100k/s, maybe the amount of data would be too big for Kinesis Analytics.

I'm leaning towards option B as this seems to be quite cost efficient. For scaling solution B I'm thinking I could create multiple IoT rules/firehose/analytics streams and let clients send their data to one of those rules round robin. Then combine results of these streams in Redis or DynamoDB and push them to the moderator.

Curious on your thoughts if I'm missing something here or if you have any experiences of your own ingesting/analysing this amount of data. I prefer to do everything using Cloudformation and serverless/managed solutions.
## [3][Migration from On-prem to AWS](https://www.reddit.com/r/aws/comments/hpm1ak/migration_from_onprem_to_aws/)
- url: https://www.reddit.com/r/aws/comments/hpm1ak/migration_from_onprem_to_aws/
---
Does anyone have a diary of how they migrated their servers from on-prem to AWS? I'm interested in what AWS services you used to connect your on-prem network to AWS. I'm also interested to hear on how you were able to successfully migrated applications to AWS. And if you're going to repeat it again, what will you revise, improve or not do?

Thanks in advance!
## [4][Cognito: create a custom “username” in the background (for userIds)](https://www.reddit.com/r/aws/comments/hprudj/cognito_create_a_custom_username_in_the/)
- url: https://www.reddit.com/r/aws/comments/hprudj/cognito_create_a_custom_username_in_the/
---
I'm using **Amazon Cognito User Pools** for user management.

I want to use the ***username*** attribute from the user record inside the Cognito User Pool to store a custom **userId**, created internally by my application, without the user even knowing.

The reason is that this ***username*** attribute is a great solution for storing a userId: it's **non-mutable**, **forcibly unique**, and **queryable** by [Cognito's APIs](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_Operations.html) (ListUsers, AdminListGroupsForUser, AdminAddUserToGroup, etc.). Cognito's custom attributes for example are not a good alternative because they can't be used to query those APIs.

So by using the ***username*** attribute I'll be able to fully manage my users within Cognito, without the need to maintain user records in another database and keep them in sync.

Additionally, I want the users to sign up and sign in using an **email address** and **password** (while in the background I use the ***username*** attribute for my own purpose).

In [this page](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html) of Amazon Cognito's guide, it says this:

&gt;If your application does not require a username, you do not need to ask users to provide one. **Your app can create a unique username for users in the background**. This is useful if, for example, you want users to register and sign in with an email address and password.

This paragraph above suggests that what I need can actually be done with Cognito.

However, Amazon Cognito's documentation doesn't explain anything further on how to make that particular setup happen.

When setting up a Cognito User Pool, there's this configuration: “**How do you want your end users to sign-up and sign-in?**”

The options given are

* **Username**
* **Email address or phone number**

**Which of these options should I pick?** What other configuration must be set for this to wok?
## [5][Error lambda with gateway api](https://www.reddit.com/r/aws/comments/hprixr/error_lambda_with_gateway_api/)
- url: https://www.reddit.com/r/aws/comments/hprixr/error_lambda_with_gateway_api/
---
Hi, everybody, I wrote a lambda function that takes data from a postgres database in RDS. To do this I used the psycopg2 library, when I run the tests the function works correctly. so I added a gateway api get as lambda trigger, but when I call the url of the get it gives me this error:

`{"errorMessage": "Unable to import module 'lambda_function': /var/task/psycopg2/_psycopg.so: undefined symbol: PyUnicodeUCS4_DecodeUTF8", "errorType": "Runtime.ImportModuleError"}`

Can you help me? Why tests work and give me back the query data while if I call them through a get it doesn't work? Thank you.
## [6][RDS Snapshot](https://www.reddit.com/r/aws/comments/hpnx15/rds_snapshot/)
- url: https://www.reddit.com/r/aws/comments/hpnx15/rds_snapshot/
---
Hi Everyone,  


First, I want to thank you all for a very helpful community. I really appreciate how you help everyone here.  


I do have another question regarding RDS Snapshot. I've been using snapshot for 1 and a half yr now and it's my first time to come across this issue. After I restore my snapshot, I wouldn't be able to log in even I change the Master Password. Can anyone help me why I can't log in to it?  


But the original Instance is still running and I can still log into it.

https://preview.redd.it/vuvmfbc1mca51.png?width=732&amp;format=png&amp;auto=webp&amp;s=d5285098b386d57da0c97cd1d29330e95122b989
## [7][What is the best approach to providing managed AWS services to clients?](https://www.reddit.com/r/aws/comments/hpq65h/what_is_the_best_approach_to_providing_managed/)
- url: https://www.reddit.com/r/aws/comments/hpq65h/what_is_the_best_approach_to_providing_managed/
---
I am thinking of creating a managed cloud service for my clients, generally AWS. There are many services that already do this. But my agency has our own long term clients and I want our company to provide a complete ecosystem for them so that a single login is all they need.

I am not sure how to manage AWS account of hundreds of clients. What is the best approach to give them AWS services without them having to create an account? We thought about two ways to do this.

1. Programatically create AWS account for each client using our organization account. There are limits to the number of accounts, which we think we will be able to increase on requesting. And the advantage is monitoring of billing of each accounts.
2. Create different VPC for each client on our own AWS account, which comes at a cost of no individual client monitoring and ban on any services caused by one account will affect every other client. WE DON'T WANT TO DO THIS.

These are the two ways we though of implementing this. Is there any other approach? If someone else is already doing this, what is your approach to this.
## [8][Attempting to find private key for ssh into EC2](https://www.reddit.com/r/aws/comments/hpoq8v/attempting_to_find_private_key_for_ssh_into_ec2/)
- url: https://www.reddit.com/r/aws/comments/hpoq8v/attempting_to_find_private_key_for_ssh_into_ec2/
---
When creating my EC2 instance on aws (For ELASTIC BEANSTALK), I indicated that I wanted to setup ssh. I was ask for a user name and a password. In return, I was given an SHA256 key and the key's art. No idea what either of these are or how these can be used right now, but I have them saved.

&amp;#x200B;

I now want to connect to my instance via ssh, but I need a .pem or .ppk file for the private key. I have neither and have no idea where to find these or how to create them.

&amp;#x200B;

When I created my instance via awsebcli, would the private key have been saved somewhere on my computer?

&amp;#x200B;

Thanks!
## [9][Cloudformation GetAtt method from inside my resolver request mapping template](https://www.reddit.com/r/aws/comments/hpfs2p/cloudformation_getatt_method_from_inside_my/)
- url: https://www.reddit.com/r/aws/comments/hpfs2p/cloudformation_getatt_method_from_inside_my/
---
I have a template for updating cognito user attributes:

```
Resources:
  UpdateUserAttributeMutation:
    Type: AWS::AppSync::Resolver
    Properties:
      ApiId:
        Ref: AppSyncApi
      TypeName: Mutation
      FieldName: updateUserAttribute
      DataSourceName:
        Ref: AppsyncDynamoDBTableDataSource
      RequestMappingTemplate: |
        set($body = {})

        #set($attribute={})
        $util.qr($attribute.put("Name", "email_verified"))
        $util.qr($attribute.put("Value", "true"))

        #set($UserAttributes = [])
        $util.qr($UserAttributes.add($attribute))
        $util.qr($body.put("UserAttributes", $UserAttributes))

        #set($body.Username = $ctx.identity.sub)
        #set($body.UserPoolId = ***MY_COGNITO_ID***)

        {
          "version": "2018-05-29",
          "method": "POST",
          "resourcePath": "/",
          "params": {
            "headers": {
              "content-type": "application/x-amz-json-1.1",
              "x-amz-target":"AWSCognitoIdentityProviderService.AdminUpdateUserAttributes"
            },
            "body": $util.toJson($body)
          }
        }
      ResponseMappingTemplate: |
        #if($ctx.error)
          $util.error($ctx.error.message, $ctx.error.type)
        #end
        #if($ctx.result.statusCode == 200)
            true
        #else
         false
        #end
```

I have multiple environments, so I would love to be able to pass the `MY_COGNITO_ID` value into my template dynamically, is there a way to achieve this using GetAtt cloudformation helper method or do I need to use something different?
## [10][Can instance-store images not be used with current gen instance types?](https://www.reddit.com/r/aws/comments/hpi6qa/can_instancestore_images_not_be_used_with_current/)
- url: https://www.reddit.com/r/aws/comments/hpi6qa/can_instancestore_images_not_be_used_with_current/
---
For my application, I do not need EBS volumes (nor can I afford 8GB of EBS volume per server per month).

I'm using Ubuntu provided AMIs: https://cloud-images.ubuntu.com/locator/ec2/

I have chose ami-0921ae3f2ab9e0efc which is Ubuntu 18.04 and everything is great on it.  I have several running now.

The only problem is, I can only choose old instance types like "m1.medium", which is comparable in price to the newer "t3a.medium" except the T3A machines are faster and better.

What gives? Can instance-store only AMIs not be launched on t3a servers?
## [11][question about client vpn](https://www.reddit.com/r/aws/comments/hpmit0/question_about_client_vpn/)
- url: https://www.reddit.com/r/aws/comments/hpmit0/question_about_client_vpn/
---
when you create a client vpn to connect to a vpc, your instances don't see you with your client CIDR and instead they see you as connecting from the subnet that you associated in the vpn endpoint

when you take a look at the endpoint you see a route table being added and it shows up as nat. i used netstat and was able to find out the specific ip address inside the subnet that actually was connected to the instance.

is this nat information stored somewhere? can i control which exact ip with that subnet is being used?
