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
## [2][Static Hosting in S3 bucket question](https://www.reddit.com/r/aws/comments/hlhomf/static_hosting_in_s3_bucket_question/)
- url: https://www.reddit.com/r/aws/comments/hlhomf/static_hosting_in_s3_bucket_question/
---
I am going through the guide of hosting a static website in an s3 bucket. The guidances says i should create:

1. S3 bucket of [example.com](https://example.com)
2. Another S3 bucket of [www.example.com](https://www.example.com) and redirect to [example.com](https://example.com)

Questions:

1. Can i not just point both [example.com](https://example.com) and [www.example.com](https://www.example.com) to the same S3 bucket instead of doing a redirect on the buckets?
2. Why do the names of the buckets have to match my DNS? Can i not have [example.com](https://example.com) point to s3 bucket randomly named?
## [3][Permissions denied when using cross account roles for Jenkins](https://www.reddit.com/r/aws/comments/hldel3/permissions_denied_when_using_cross_account_roles/)
- url: https://www.reddit.com/r/aws/comments/hldel3/permissions_denied_when_using_cross_account_roles/
---
I am trying to get a Jenkins server in the root account of an organization to be able to push a serverless application (nodejs using serverless framework) to a new development account. I have setup a JenkinsAccessRole that has a trust relationship with the main account.

IAM Policy for JenkinsAccessRole in the new development account

    {
      "Version": "2012-10-17",
      "Statement": [
        {
          "Effect": "Allow",
          "Principal": {
            "AWS": "arn:aws:iam::&lt;account ID for root account&gt;:root"
          },
          "Action": "sts:AssumeRole",
          "Condition": {}
        }
      ]
    }

The role has permissions for cloudformation, sqs, sns and s3

Error message seen in Jenkins

    com.amazonaws.services.securitytoken.model.AWSSecurityTokenServiceException: User: arn:aws:sts::&lt;root account ID&gt;:assumed-role/Jenkins/i-015333655393dd020 is not authorized to perform: sts:AssumeRole on resource: arn:aws:iam::&lt;new dev account ID&gt;:role/JenkinsAccessRole (Service: AWSSecurityTokenService; Status Code: 403; Error Code: AccessDenied;

Jenkins code

    withAWS(role: 'JenkinsAccessRole', roleAccount: '&lt;main&gt;', duration: 3600, roleSessionName: 'Serverless-Deploy') {
         sh "npm run deployDev"
    }

Can anyone spot the issue or give suggestions on what might be wrong?

&amp;#x200B;

EDIT\*\* 

Figured out my issue, I had a policy for the other dev account that allowed my Jenkins server to assume the role which connected the accounts. 
## [4][Easy Web Server in AWS Ec2](https://www.reddit.com/r/aws/comments/hlh0br/easy_web_server_in_aws_ec2/)
- url: https://www.reddit.com/r/aws/comments/hlh0br/easy_web_server_in_aws_ec2/
---
I wrote some super detailed documentation awhile ago on how to setup a web server on Ec2 in AWS and it always helps me time and time again.   
So I thought it could be helpful to anyone here trying to start a web server on AWS:  
[https://github.com/stevegardiner26/aws-webservers-creation-guide](https://github.com/stevegardiner26/aws-webservers-creation-guide)
## [5][Amazon Polly - where do I choose regular TTS vs neural network?](https://www.reddit.com/r/aws/comments/hljd5t/amazon_polly_where_do_i_choose_regular_tts_vs/)
- url: https://www.reddit.com/r/aws/comments/hljd5t/amazon_polly_where_do_i_choose_regular_tts_vs/
---
I'm new to AWS, I'm in the console on the page that says:

Text-to-Speech

Listen, customize, and download speech. Integrate when you're ready.

Type or paste your text in the window, choose your language and region, choose a voice, choose Listen to speech, and then integrate it into your applications and services.

With up to 3000 characters you can listen, download, or save immediately. For up to 100,000 characters, your task must be saved to an S3 bucket.
## [6][Can IAM user with AWSPriceListServiceFullAccess policy access any other resources if keys are exposed?](https://www.reddit.com/r/aws/comments/hlg7b7/can_iam_user_with_awspricelistservicefullaccess/)
- url: https://www.reddit.com/r/aws/comments/hlg7b7/can_iam_user_with_awspricelistservicefullaccess/
---
I have created a user with AWSPriceListServiceFullAccess policy attached. I wanted to use this to call AWS pricelist API for an application that I was building. The code of the application was on Github in a private repository. As soon as I changed it to public repo, the AWS warned me about a potential security breach and they told (forced) me to delete the user.

I don't think that on exposing those keys, there is any threat to my AWS account. Can anyone explain if there is any threat to that?
## [7][Amazon SES Security?](https://www.reddit.com/r/aws/comments/hl7wpk/amazon_ses_security/)
- url: https://www.reddit.com/r/aws/comments/hl7wpk/amazon_ses_security/
---
I logged into my SES today to see that over 1400 emails were sent, but I haven't run any campaigns which is a problem.   


I'm not sure how to isolate where these emails were sent from, perhaps it was because of a server error or perhaps not. But how do I isolate it so that I know for sure that I wasn't hacked or had my SMTP details compromised?
## [8][Build quickly a system that filter CloudWatch logs and post to Slack, via CDK.](https://www.reddit.com/r/aws/comments/hkx584/build_quickly_a_system_that_filter_cloudwatch/)
- url: https://github.com/TheDesignium/cdk-log-notifier
---

## [9][Custom CloudWatch Metrics vs DynamoDB](https://www.reddit.com/r/aws/comments/hl55p7/custom_cloudwatch_metrics_vs_dynamodb/)
- url: https://www.reddit.com/r/aws/comments/hl55p7/custom_cloudwatch_metrics_vs_dynamodb/
---
I want to track the number of requests by user to my API, and the API runs via Lambda in every AWS region  (20 regions) because of latency.

Does it make sense to use CloudWatch metrics when operating in a multi-region format?

CloudWatch endpoints are per region, meaning a custom metric will have to be created for each user in each region where they run.

If I had one metric per user that means I'd be spending $0.30 \* 20 = $6 a month per user to track their requests over time.  That seems a bit too much to spend.

If I had one DynamoDB table in a single region that tracked traffic with TTLs on the records it seems that is the more efficient way, but would suffer the problem of implementing the retry logic if that region isn't reachable for some region.

How have you all solved this problem before?
## [10][Cost effectively connecting to Aurora Serverless for personal project](https://www.reddit.com/r/aws/comments/hla44b/cost_effectively_connecting_to_aurora_serverless/)
- url: https://www.reddit.com/r/aws/comments/hla44b/cost_effectively_connecting_to_aurora_serverless/
---
I'm attempting to use Aurora Serverless, but I'm getting tripped up by how to cost effectively manage my VPC. My requirements:

1. Aurora must be accessible from a Lambda (the Lambda must have public internet access)
2. I'd like to use a desktop client for connecting to my DB (for configuring tables, etc)

I believe I'd want to use a private subnet with a NAT Gateway in a production environment, but I'd like to avoid that here since NAT Gateways are expensive. Should I just use a public subnet and accept the security implications (not that there are many for this project)?

If I go with the public subnet, should I be able to connect to the DB from a desktop client without needing a VPN? I was expecting to, but I can't get it working. I'm not sure if I just need to update something like my security group/ACL, or if I have a misunderstanding in how this all works.
## [11][Front or back end for Cognito](https://www.reddit.com/r/aws/comments/hl4tdh/front_or_back_end_for_cognito/)
- url: https://www.reddit.com/r/aws/comments/hl4tdh/front_or_back_end_for_cognito/
---
I'm building an application that will use AWS Cognito for handling user authentication and such. I have seen a few tutorials that use everything in the front end including the pool and client ID, this seems like a security issue to me. However when I try to make it in a backend service it results in an error about missing fetch and some googling shows it's because I'm using Node.js isnstead of something like Vue.js on the front end.

&amp;#x200B;

Should the request (and client/pool ID) come from the front end or the back end? How big of security issue is having the ID's in the front end?
