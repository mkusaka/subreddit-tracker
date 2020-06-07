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
## [2][usiing Jira to automate AWS workspace provision](https://www.reddit.com/r/aws/comments/gy1yx3/usiing_jira_to_automate_aws_workspace_provision/)
- url: https://www.reddit.com/r/aws/comments/gy1yx3/usiing_jira_to_automate_aws_workspace_provision/
---
I am working in a space that uses Aws workspace for every user to complete office task. We use jira service desk to kick off users onboarding. I am trying to see if anyone has automated the workspace creation using jira service desk. Ideally I would want something to create a workspace for a user soon as a ticket request comes into jira.
## [3][Trying to limit read access to S3 bucket from only one website](https://www.reddit.com/r/aws/comments/gy6i2m/trying_to_limit_read_access_to_s3_bucket_from/)
- url: https://www.reddit.com/r/aws/comments/gy6i2m/trying_to_limit_read_access_to_s3_bucket_from/
---
Hi!

I have a bucket, and I want to limit read access from only one website. Let's say, there are some images inside the bucket.

My bucket has AWS-KMS encryption.

Bucket policy is:

    {
        "Version": "2012-10-17",
        "Id": "http referer policy example",
        "Statement": [
            {
                "Sid": "Allow get requests originating from mywebsite.com.",
                "Effect": "Allow",
                "Principal": {
                    "AWS": "arn:aws:iam::MY_ACCOUNT_ID:user/s3-user-mywebsite"
                },
                "Action": "s3:GetObject",
                "Resource": "arn:aws:s3:::mybucket/*",
                "Condition": {
                    "StringLike": {
                        "aws:Referer": "https://mywebsite.com/*"
                    }
                }
            }
        ]
    }

`s3-user-mywebsite` has  `AmazonS3FullAccess` permission. Block public access settings are all off.

Anyway, I'm getting 400 when trying to render images from the bucket. I'm sure that there is correct `Referrer` header in request.

What am I missing?
## [4][What are the best services to learn for front end engineering?](https://www.reddit.com/r/aws/comments/gy13a8/what_are_the_best_services_to_learn_for_front_end/)
- url: https://www.reddit.com/r/aws/comments/gy13a8/what_are_the_best_services_to_learn_for_front_end/
---
I'd like to dive deeper into more AWS services, but typically when you Google learning AWS services, you get EC2 or VPN and the like. What services are particularly useful when engineering as a client-side JavaScript developer?

For example, monitoring metrics via CloudWatch is still useful. Code Suite for CI/CD is great. CloudFormation for infrastructure as code is great. What other services are interesting and useful?
## [5][Do tools like Terraform fill CloudFormation imperfections and gaps?](https://www.reddit.com/r/aws/comments/gxofte/do_tools_like_terraform_fill_cloudformation/)
- url: https://www.reddit.com/r/aws/comments/gxofte/do_tools_like_terraform_fill_cloudformation/
---
I'm two weeks into CloudFormation, and I've run into these three things in this short amount of time:

* Glue Job tags cannot be updated with CloudFormation: https://github.com/aws-cloudformation/aws-cloudformation-coverage-roadmap/issues/306

* RDS DBClusters cannot specify CopyTagsToSnapshots via CloudFormation: https://github.com/aws-cloudformation/aws-cloudformation-coverage-roadmap/issues/238

* If your RDS DBInstance specifies CaCertificateAuthority, all subsequent updates to anything in the entire stack will cause your instance to reboot, whether or not you make any changes to the CaCertificateAuthority attribute (this appears to not technically be a "bug", but CloudFormation gives no reboot warming and the CaCertificatrAuthority attribute does not show up in the ChangeSet if you didn't actually change it).

Only two weeks in, and I'm already needing to keep supplementary internal documentation about how to handle these cases, which makes me wonder how much worse this is going to get.

Do tools like Terraform fill these gaps?  Or do they themselves have their own similar failings?
## [6][Resources inside and outside of VPC](https://www.reddit.com/r/aws/comments/gy40q1/resources_inside_and_outside_of_vpc/)
- url: https://www.reddit.com/r/aws/comments/gy40q1/resources_inside_and_outside_of_vpc/
---
I am learning more about networking in AWS and I was wondering about what services go inside a VPC and what do not ?

For instance - RDS is deployed inside a VPC but S3 is not ? 

How is this determined ?
## [7][Static website authentication when Cognito is overkill?](https://www.reddit.com/r/aws/comments/gxsz0n/static_website_authentication_when_cognito_is/)
- url: https://www.reddit.com/r/aws/comments/gxsz0n/static_website_authentication_when_cognito_is/
---
I need to create a static website that has an authentication layer before it allows access to a web form. The web form will POST to API Gateway and trigger a Step Function.

Cognito seems like overkill because this is for low volume, internal staff use. A single password shared by staff should suffice.

What approach could I take for ensuring the service web form is only viewable/usable after the user inputs the correct password? I was thinking a basic password input web form followed by some mechanism to serve up the service web form (assuming the correct password is supplied).

Edit: To clarify, I'm using the serverless s3 static website approach.
## [8][How to check from when you use a given service within FREE TIER and when 12 months period will end?](https://www.reddit.com/r/aws/comments/gxzvfe/how_to_check_from_when_you_use_a_given_service/)
- url: https://www.reddit.com/r/aws/comments/gxzvfe/how_to_check_from_when_you_use_a_given_service/
---
Does each service have its own period of time, which starts from the first start of the service??
## [9][Ugly plan for cheaper snapshots](https://www.reddit.com/r/aws/comments/gxs48u/ugly_plan_for_cheaper_snapshots/)
- url: https://www.reddit.com/r/aws/comments/gxs48u/ugly_plan_for_cheaper_snapshots/
---
In my local region EBS snapshots cost $0.05 per GB/month. SC1 volumes cost $0.028 per GB/month.

I know it's not a fair direct comparison as snapshots do block level de-duplication. 

However, I could take my snapshots, split them into 1MB blocks, hardlink identical blocks, and compress all those blocks.

It should cost about half the price to store snapshots in this way compared to AWS's snapshots.
## [10][AppStream: How to save user fees?](https://www.reddit.com/r/aws/comments/gxpdk7/appstream_how_to_save_user_fees/)
- url: https://www.reddit.com/r/aws/comments/gxpdk7/appstream_how_to_save_user_fees/
---
Hi,

I am considering to use AppStream as part of my solution. The problem is that it comes with a rather steep fee of 4,19$ per user per month for licenses.

As my use-case would have thousands or more of users using the application just for a very short period of time (maybe 1 hour per months) this would be to pricey.

I don't have the user pool created yet so I am flexible where the users would come from. In general it would be an application where anybody could sign up online.

The application I am streaming is a proprietary one not under my control. So I can't modify it. Still I would prefer to be able to save the user settings so that a user can configure that application individually and maintain the settings when he logs in a some time later.

Any ideas how I could optimize cost if I

- Save the settings as described above

- do not save the settings and let the user re-configure everything from scratch
## [11][With the release of South African servers recently, does anyone know which region aws is going to work on building new servers?](https://www.reddit.com/r/aws/comments/gxxa84/with_the_release_of_south_african_servers/)
- url: https://www.reddit.com/r/aws/comments/gxxa84/with_the_release_of_south_african_servers/
---
They said in 2019 it was going to be Middle East and in 2020 it was going to be South Africa. Does anyone know what next?
