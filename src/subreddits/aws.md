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
## [2][What is AWS competitor to firebase database?](https://www.reddit.com/r/aws/comments/hd3zaw/what_is_aws_competitor_to_firebase_database/)
- url: https://www.reddit.com/r/aws/comments/hd3zaw/what_is_aws_competitor_to_firebase_database/
---
I found very easy to setup and use firebase database as a web dev.  Does AWS have a similar technology that is as simple to setup and use?
## [3][{"message":"Missing Authentication Token"} Chalice API Authentication Erorr?](https://www.reddit.com/r/aws/comments/hd2syh/messagemissing_authentication_token_chalice_api/)
- url: https://www.reddit.com/r/aws/comments/hd2syh/messagemissing_authentication_token_chalice_api/
---
Hello! 

I've been poking around AWS Chalice and Lambda, and I've been running into an issues with getting my API set up. I am consistiently running into a `{"message":"Missing Authentication Token"}` error.

My AWS rootkey credentials in the AWS configure settings are correct, however the error still remains.

The API never seems to update even when 200 POST requests are made from test clients such as Insomnia.

Anyone run into a similar issue?

`import requests, json`

`from chalice import Chalice, Response`

&amp;#x200B;

`app = Chalice(app_name='helloworld')`

&amp;#x200B;

`@app.route('/', methods=['POST'])`

`def index():`

`message = app.current_request.json_body`

`return {'webhook_message': message}`
## [4][AWS Phone Support: Just call me when a tech is available!](https://www.reddit.com/r/aws/comments/hcn13f/aws_phone_support_just_call_me_when_a_tech_is/)
- url: https://www.reddit.com/r/aws/comments/hcn13f/aws_phone_support_just_call_me_when_a_tech_is/
---
Every time I go to open a phone ticket with Amazon they call my landline immediately, then make me wait an hour or two for a tech. (At least I think they would... usually the call drops before the tech shows up.)

For the love of god, why not just call me when the tech is available.

Also how about having more than 16 bars of looping hold music. Maybe you could tell me how long the wait is going to be?

**Are phone tickets just a suckers game?**  


**Edit:** The support I get from Amazon is top-notch. I'm just nit-picky about the phone system.
## [5][Website (using Angular) accessing API with authorization? help!](https://www.reddit.com/r/aws/comments/hd4een/website_using_angular_accessing_api_with/)
- url: https://www.reddit.com/r/aws/comments/hd4een/website_using_angular_accessing_api_with/
---
I'm new to the web front-end programming, as I have always been more of a SQL and Python developer. I am building a website, and I have a MySQL DB (on AWS) with many tables of Sports data for example. I just want my web domain to display different analytics or reports. I found that I needed an API which hits the DB, which I've built in API Gateway.  I am learning Angular to make this site, and I have successfully connected to my API in Angular (using HttpRequest) and displayed data successfully locally on my browser (localhost:4200)! So that was great, but I then thought my API should have security/authorization though, so only my web code is allowed to hit my API. I added authorization in my API's settings (See pic: [https://imgur.com/a/xxBgksv](https://imgur.com/a/xxBgksv)) and I successfully tested connecting using Postman by plugging in my IAM AccessKey and SecretKey. But, every example or tutorial on Angular accessing an API either brushes over this signature part or uses Cognito as part of it's solution. I don't see why I'd use Cognito though. I am not making a pool of users for my site. I just need ONE access key for the website code only. I don't understand how there's no sample Angular code for a regular dynamic website accessing an API. This seems like a very common scenario for most people building a website. It's making me question if I am going down the right path as it seems more difficult than it should be. I'm assuming everyone has IAM authorization turned on for their API endpoint, right? I mean I guess it would be hard for someone to even find out my endpoint path, but I feel like it should have it regardless. I'm at a loss, please help. Thanks
## [6][Is AppSync the best way to set up a GraphQL API on aws for a web app?](https://www.reddit.com/r/aws/comments/hd5uhu/is_appsync_the_best_way_to_set_up_a_graphql_api/)
- url: https://www.reddit.com/r/aws/comments/hd5uhu/is_appsync_the_best_way_to_set_up_a_graphql_api/
---
Or is it mostly just for mobile apps?

If it’s not the one Th en are there any managed graphql solutions like API gateway or do you need to set up your own custom thing?
## [7][DynamoDB or MongoDB in ec2 instance?](https://www.reddit.com/r/aws/comments/hd2517/dynamodb_or_mongodb_in_ec2_instance/)
- url: https://www.reddit.com/r/aws/comments/hd2517/dynamodb_or_mongodb_in_ec2_instance/
---
We are planning to migrate our ecosystem completely to the AWS. Currently mongodb is used  as a backend with replica set . Our system is both write heavy and read heavy. And data is extremly unstructred. And it stores few millions of records. And very high throughput. 

So which is better in AWS for this use case. MongoDB in EC2 or DynamoDB interms of performance, cost, consistency and HA?
## [8][Rent AWS services to others](https://www.reddit.com/r/aws/comments/hd4g3p/rent_aws_services_to_others/)
- url: https://www.reddit.com/r/aws/comments/hd4g3p/rent_aws_services_to_others/
---
Hello people, Is there any place where I can give my AWS servers for rent for example? Let's say I'm gonna take them only half of the money that it costs me in AWS services and provide to them cheaper. I'm thinking that becauseI won 6,500$ credits in AWS by applying for an  idea
## [9][I have access key and secret, and a pem file. Can I get in?](https://www.reddit.com/r/aws/comments/hd40qf/i_have_access_key_and_secret_and_a_pem_file_can_i/)
- url: https://www.reddit.com/r/aws/comments/hd40qf/i_have_access_key_and_secret_and_a_pem_file_can_i/
---
It has been a while since I maintained my AWS account, and in the end I am only using it for a Lightsail instance.  I changed my phone a few months back and of course lost my virtual MFA key in the process... and the root registered phone number is a landline I no longer have access to. 

I am preparing to go and get the notarised documentation etc. as advised by  Customer Support when I contacted them for help yesterday accessing as root user.  I completely understand the reason, so this is not a whinge.  

I do however have the AWSAccessKeyID and the AWSSecretKey in a file on my laptop.  But I cannot find any documentation if it is possible to use the CLI and these keys, or a PEM file which I also have, to connect as root.  

Is anyone able to advise if this is even worth investigating, and maybe some documentation that would help me try it out?  

Thanks, long shot I guess.
## [10][Is one S3 bucket enough?](https://www.reddit.com/r/aws/comments/hd1fql/is_one_s3_bucket_enough/)
- url: https://www.reddit.com/r/aws/comments/hd1fql/is_one_s3_bucket_enough/
---
Hi, I'm developing an e-commerce platform. I thought using one bucket to store products images and user profile and store profile using the folders

like 

    mybucket/products
    mybucket/user
    mybucket/store

or should I use separate buckets?
## [11][12 Months free tier questions](https://www.reddit.com/r/aws/comments/hd1eso/12_months_free_tier_questions/)
- url: https://www.reddit.com/r/aws/comments/hd1eso/12_months_free_tier_questions/
---
1. Am I charged when downloading files from googledrive to my EC2?
2. Am I charge when uploading files from EC2 to google drive?
3. If 1 and 2 are yes, what is the FREE bandwidth limit?
4. I have 1 shapshot and I did not know if this is created automatically or not. Is it free?
