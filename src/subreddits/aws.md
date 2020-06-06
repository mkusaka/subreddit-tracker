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
## [2][Do tools like Terraform fill CloudFormation imperfections and gaps?](https://www.reddit.com/r/aws/comments/gxofte/do_tools_like_terraform_fill_cloudformation/)
- url: https://www.reddit.com/r/aws/comments/gxofte/do_tools_like_terraform_fill_cloudformation/
---
I'm two weeks into CloudFormation, and I've run into these three things in this short amount of time:

* Glue Job tags cannot be updated with CloudFormation: https://github.com/aws-cloudformation/aws-cloudformation-coverage-roadmap/issues/306

* RDS DBClusters cannot specify CopyTagsToSnapshots via CloudFormation: https://github.com/aws-cloudformation/aws-cloudformation-coverage-roadmap/issues/238

* If your RDS DBInstance specifies CaCertificateAuthority, all subsequent updates to anything in the entire stack will cause your instance to reboot, whether or not you make any changes to the CaCertificateAuthority attribute (this appears to not technically be a "bug", but CloudFormation gives no reboot warming and the CaCertificatrAuthority attribute does not show up in the ChangeSet if you didn't actually change it).

Only two weeks in, and I'm already needing to keep supplementary internal documentation about how to handle these cases, which makes me wonder how much worse this is going to get.

Do tools like Terraform fill these gaps?  Or do they themselves have their own similar failings?
## [3][A collection of awesome things related to the AWS Cloud Development Kit (CDK)](https://www.reddit.com/r/aws/comments/gx61ab/a_collection_of_awesome_things_related_to_the_aws/)
- url: https://github.com/eladb/awesome-cdk
---

## [4][AWS Services list in JSON](https://www.reddit.com/r/aws/comments/gxl3ly/aws_services_list_in_json/)
- url: https://www.reddit.com/r/aws/comments/gxl3ly/aws_services_list_in_json/
---
Long time no see, I've bee buys, but in the mean time I did a small repo that might be useful to some. I basically made a JSON file with all the AWS services. What for? because AWS dose not provide nor maintain such file, and some times you might want to build something around what AWS has to offer. Here is the repo: https://github.com/0x4447/0x4447_product_aws_services_list - I hope this helps some body.
## [5][Should I use Amplify or the CDK?](https://www.reddit.com/r/aws/comments/gxmkjd/should_i_use_amplify_or_the_cdk/)
- url: https://www.reddit.com/r/aws/comments/gxmkjd/should_i_use_amplify_or_the_cdk/
---
I want to build a small static website which does not have any dynamic content. The only dynamic content will be a contact form that should send an email. That's if. 

So what's your opinion? Should I just use [AWS Amplify](https://aws.amazon.com/amplify/) for that because "it just works"? Or should I create the stack (S3 / CloudFront / Lambda) by myself with the [AWS CDK](https://aws.amazon.com/cdk/) to have the full control an be flexible for the future? For example I could easily include [AWS SES](https://aws.amazon.com/ses/) that will send the emails from the contact form that gets send from a small Lambda. I didn't found an easy way to do that with `Amplify`. On the other site `Amplify` is pretty easy for hosting a static website including CDN cache invalidation, atomic and delta deploys (how can I do that with the CDK?) and a Lambda for the contact form, but without `SES`.

Or should I completely forget about and just use [Netlify](https://www.netlify.com) where everything just works out of the box?
## [6][Can I use a domain that is not supported by route 53?](https://www.reddit.com/r/aws/comments/gxo67p/can_i_use_a_domain_that_is_not_supported_by_route/)
- url: https://www.reddit.com/r/aws/comments/gxo67p/can_i_use_a_domain_that_is_not_supported_by_route/
---
I don't have any experience with domains and I was wondering if I would be able to attach a domain that is not supported by route 53 to my backend services? Thanks.
## [7][re:Invent 2020](https://www.reddit.com/r/aws/comments/gx6cej/reinvent_2020/)
- url: https://www.reddit.com/r/aws/comments/gx6cej/reinvent_2020/
---
What are your thoughts on re:Invent 2020 from happening? 

I’m kind of surprised they haven’t cancelled yet, but maybe they got a sweetheart deal to keep it active. Considering there are limited options for travel internationally right now, I’d be really surprised if the show went on. 

Also, the conference would be nearly impossible for me to get approval to attend. All travel is banned without sign-offs from several management layers above me plus budgets have been decimated to *nearly* zero.
## [8][Torn between lambda or EKS setup. I'm starting a new application supporting 500 users, probably 20 or so per hour. And in phase 2 looking to integrate with third party corporate customers looking to hit the API potentially 20+ concurrent calls every second. We will have SLAs to adhear to.](https://www.reddit.com/r/aws/comments/gxesst/torn_between_lambda_or_eks_setup_im_starting_a/)
- url: https://www.reddit.com/r/aws/comments/gxesst/torn_between_lambda_or_eks_setup_im_starting_a/
---
I work in a consultancy where I've been setting up alot of serverless setups. It was a no brainer setting up dashboard applications for our clients quickly. for my next project, I'm considering EKS due to requirements for higher concurrent API hits. I'm thinking cost and potentially a small risk we may be asked to move to Azure.
What do you guys think?

[View Poll](https://www.reddit.com/poll/gxesst)
## [9][Slack adopts Chime Video Call technology](https://www.reddit.com/r/aws/comments/gwt7vu/slack_adopts_chime_video_call_technology/)
- url: https://www.cnbc.com/2020/06/04/amazon-licenses-slack-for-workers-as-slack-adopts-aws-video-call-tech.html
---

## [10][Unexpectedly high costs with Cloudwatch](https://www.reddit.com/r/aws/comments/gx951u/unexpectedly_high_costs_with_cloudwatch/)
- url: https://www.reddit.com/r/aws/comments/gx951u/unexpectedly_high_costs_with_cloudwatch/
---
Our Cloudwatch costs suddenly ballooned from $200 per month to $1600 per month, and we're having trouble tracking down the source.

According to our billing dashboard, the last month saw 130 million API requests (not including GetMetricData). 

We have a few small auto scaling clusters, and we're also using Zabbix.

Has anyone had a similar unexpected increase and how did  you track the source down?
## [11][Suspicious activities on S3. Need help.](https://www.reddit.com/r/aws/comments/gx8fjm/suspicious_activities_on_s3_need_help/)
- url: https://www.reddit.com/r/aws/comments/gx8fjm/suspicious_activities_on_s3_need_help/
---
I'm seeing some suspicious activities on S3 buckets in my AWS account. These activities are coming from AWS Config Service from an unknown account. I wrote about this over here  [https://www.logsec.cloud/2020/06/04/investigating-on-aws/](https://www.logsec.cloud/2020/06/04/investigating-on-aws/) . Can someone please assist what should I do next?
