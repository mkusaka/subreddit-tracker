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
## [2][I built an IP geolocation module optimized for AWS lambda](https://www.reddit.com/r/aws/comments/hqmhia/i_built_an_ip_geolocation_module_optimized_for/)
- url: https://github.com/corollari/fast-geoip
---

## [3][1 VPC with 80 IPSEC VPNs](https://www.reddit.com/r/aws/comments/hqy6jv/1_vpc_with_80_ipsec_vpns/)
- url: https://www.reddit.com/r/aws/comments/hqy6jv/1_vpc_with_80_ipsec_vpns/
---
Hi all,

I need to connect a VPC with 80 on-premise networks, using industrial cellular routers and IPSEC. Running 24/7, dataflow is expected to be negligible. No high availability is requested for the "individual" VPN connections. And even if not requested yet, we must be prepared to support a redundant VPN server.

VGW and Site-to-site VPNs look like a great option but it's nearly 3k USD/month and there are some soft limits (50 site-to-site VPNs/regions + 10 site-to-site VPNs/VGW). 

To overcome VGW limits, we could use a Transit Gateway but it's 3k USD/month only because of 81 attachements. I also have to add the price of site-to-site VPNs, correct ? 3k + 3k = 6k USD/month ?

&amp;#x200B;

Do you think we could push VGW soft limits to 80 VPNs/VGW ?

Is it a good architecture to use Site-to-site VPNs &amp; VGW/Transit Gateway for that kind of use ?

What would you recommend instead for a more cost effective solution ? Have a pair of IPSec server instances (strongswan?) in my public subnets ?

&amp;#x200B;

Thanks
## [4][Attempting to get SSM param gives error: ParameterNotFound: null](https://www.reddit.com/r/aws/comments/hqx65m/attempting_to_get_ssm_param_gives_error/)
- url: https://www.reddit.com/r/aws/comments/hqx65m/attempting_to_get_ssm_param_gives_error/
---
I have simple nodejs code that grabs a param, but I get `ParameterNotFound: null` for reasons I don't understand. What could I be doing wrong here?

`const AWS = require('aws-sdk')`  
`const ssm = new AWS.SSM({region: '&lt;region&gt;'});`  
`const getParameterValue = async (params) =&gt; {`  
 `var request = await ssm.getParameter(params).promise();`  
 `console.log(request.Parameter.Value);`   
`};`  


&amp;#x200B;

Params:  
 `var params = {`  
 `Name: '&lt;param_name&gt;',`   
 `WithDecryption: true`  
 `};`
## [5][SQS IAM EC2 policy access](https://www.reddit.com/r/aws/comments/hqymke/sqs_iam_ec2_policy_access/)
- url: https://www.reddit.com/r/aws/comments/hqymke/sqs_iam_ec2_policy_access/
---
my inline sqs policy has receive message on an ec2 iam role; however i can't dequeue a message from the cli on the box.  what might be wrong?

aws sqs receive-message --queue-url &lt;my url&gt;

An error occurred (AccessDenied) when calling the ReceiveMessage operation: Access to the resource https://queue.amazonaws.com/ is denied.

policy:   
"Effect": "Allow",  
"Action": \[  
"sqs:DeleteMessage",  
"sqs:GetQueueUrl",  
"sqs:ChangeMessageVisibility",  
"sqs:SendMessageBatch",  
"sqs:ReceiveMessage",  
"sqs:SendMessage",  
"sqs:GetQueueAttributes",  
"sqs:ListQueueTags",  
"sqs:ListDeadLetterSourceQueues",  
"sqs:DeleteMessageBatch",  
"sqs:PurgeQueue",  
"sqs:DeleteQueue",  
"sqs:CreateQueue",  
"sqs:ChangeMessageVisibilityBatch",  
"sqs:SetQueueAttributes"  
\],
## [6][How to analyze a form via textract and A2I using lambda function?](https://www.reddit.com/r/aws/comments/hqxyub/how_to_analyze_a_form_via_textract_and_a2i_using/)
- url: https://www.reddit.com/r/aws/comments/hqxyub/how_to_analyze_a_form_via_textract_and_a2i_using/
---
Hey guys, I am trying to send the form extraction from textract to A2I using a lambda function but I have no clue how to go about implementing it. 

Currently, after following some guides, I can send the output of a textract back to S3 in a .txt format. 

How can I send the output to A2I for human review? Thank you!
## [7][Learning Opportunity.](https://www.reddit.com/r/aws/comments/hqg29q/learning_opportunity/)
- url: https://www.reddit.com/r/aws/comments/hqg29q/learning_opportunity/
---
I have not  yet seen this posted here, but I wanted to share something a co-worker has been developing and wanted to share with a wider audience.  A great way to learn to code and use AWS at the same time.

 [https://learn-to-code.workshop.aws](https://learn-to-code.workshop.aws/) 

All feedback is welcome.
## [8][Static Website Hosting &amp; Redirect Error](https://www.reddit.com/r/aws/comments/hqxd3n/static_website_hosting_redirect_error/)
- url: https://www.reddit.com/r/aws/comments/hqxd3n/static_website_hosting_redirect_error/
---
[http://example.org/](http://algostem.org/), [http://www.example.org/](http://www.algostem.org/), and [https://www.example.org/](https://www.algostem.org/) all resolve to [https://www.example.org/](https://www.woorank.com/en/www/algostem.org#) (correct). However, I get a timeout error for [https://example.org/](https://algostem.org/). I'm hosting a static website using S3, CloudFront, and Route 53. The domain is registered through Google Domains. [This](https://dev.to/razcodes/how-to-host-a-static-website-with-aws-s3-and-ssl-using-cloudfront-3e37) is the tutorial I followed (I know no one likes to open links on Reddit but it's just too long to summarize here). Any help would be greatly appreciated!
## [9][Migrate password-less, no oidc compatible user database to cognito](https://www.reddit.com/r/aws/comments/hqwm6u/migrate_passwordless_no_oidc_compatible_user/)
- url: https://www.reddit.com/r/aws/comments/hqwm6u/migrate_passwordless_no_oidc_compatible_user/
---
So I have a user database, which is a combination of facebook and email authentication, catch is that instead of password for email auth, we use some kind of "passcode" which has different meaning in the business. the fields for my user database are general, email, first name, last name, created at, updated at and uuid, used to identify users in our business. also those users are attached to another table, which contains these "passcodes". how could i migrate to cognito, and also introduce normal passwords as well?
## [10][Why some EC2 instances have "N/A" for ECU? What does N/A mean here?](https://www.reddit.com/r/aws/comments/hqyyrj/why_some_ec2_instances_have_na_for_ecu_what_does/)
- url: https://www.reddit.com/r/aws/comments/hqyyrj/why_some_ec2_instances_have_na_for_ecu_what_does/
---
For example r5d.large is a bit more expensive than r5ad.large, while having the same specs. The former has 10 ECUs, the latter has "N/A"
## [11][Egis - a handy Ruby interface for AWS Athena](https://www.reddit.com/r/aws/comments/hqye6m/egis_a_handy_ruby_interface_for_aws_athena/)
- url: https://www.reddit.com/r/aws/comments/hqye6m/egis_a_handy_ruby_interface_for_aws_athena/
---
Together with my team at [https://www.u2i.com/](https://www.u2i.com/) we have recently built a wrapper for AWS Athena SDK providing a convenient, higher-level interface for defining schemas, creating tables and executing queries. It reduces boilerplate and repetitive code allowing you to process large data sets easily.

[https://github.com/u2i/egis](https://github.com/u2i/egis)

Do you find it useful? Please give it a star. ;)
