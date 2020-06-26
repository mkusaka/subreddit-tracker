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
## [2][I wrote a thing that improves ui of several aws-cli cloudformation commands.](https://www.reddit.com/r/aws/comments/hg31lu/i_wrote_a_thing_that_improves_ui_of_several/)
- url: https://www.reddit.com/r/aws/comments/hg31lu/i_wrote_a_thing_that_improves_ui_of_several/
---
[https://github.com/molecule-man/stack-assembly/blob/master/docs/aws-drop-in.md](https://github.com/molecule-man/stack-assembly/blob/master/docs/aws-drop-in.md)

aws-cli is a great tool to automate cloudformation deployments. But a bit low-level. When I need to use some old script that does `aws cloudformation update-stack` then there are always two feelings inside me: from one side I fear to break my production as update-stack doesn't show if it's going to replace resource (data-loss) and from the other side I understand how hard it's going to be to rewrite all those small deployment scripts scattered across organisation's codebase with the ones that at least use change-sets.

I got tired of fearing and finally wrote the thing that allows to continue using those old deployment scripts and at the same time improve safety of the deployment. [The tool I wrote](https://github.com/molecule-man/stack-assembly/blob/master/docs/aws-drop-in.md) wraps aws-cli and replaces  [create-stack](https://docs.aws.amazon.com/cli/latest/reference/cloudformation/create-stack.html), [update-stack](https://docs.aws.amazon.com/cli/latest/reference/cloudformation/update-stack.html) and [deploy](https://docs.aws.amazon.com/cli/latest/reference/cloudformation/deploy.html) commands with improved versions of those commands that:

* use change sets
* present changes and give ability to approve/disapprove deployment
* give possibility to view diff of the template to be deployed
* display cloudformation events while stack is being deployed

https://i.redd.it/psdub5ur97751.gif
## [3][Automatically configure Linux Workspaces without a custom bundle](https://www.reddit.com/r/aws/comments/hfucgb/automatically_configure_linux_workspaces_without/)
- url: https://www.reddit.com/r/aws/comments/hfucgb/automatically_configure_linux_workspaces_without/
---
As Linux doesn’t support group policy, I’m looking for ways to configure my Linux workspaces. And I don’t want to create a custom bundle. Ideas? 

Chef and ansible seem like you need to create custom images.
## [4][Another Tax question](https://www.reddit.com/r/aws/comments/hg6vy3/another_tax_question/)
- url: https://www.reddit.com/r/aws/comments/hg6vy3/another_tax_question/
---
I have a new business that I have registered for VAT in the UK with AWS that has been accepted but I still have to pay VAT at 20%. Am I missing something?
## [5][How do I make the CloudWatch Alarm for Lambda Errors Default Email Body better?](https://www.reddit.com/r/aws/comments/hfulwq/how_do_i_make_the_cloudwatch_alarm_for_lambda/)
- url: https://www.reddit.com/r/aws/comments/hfulwq/how_do_i_make_the_cloudwatch_alarm_for_lambda/
---
The alarm only says when  a threshold is broken, I wish it would give me a print out of the lambda(s) that caused the alarm in the email body.
## [6][Hashtag in AWS cognito URL callback](https://www.reddit.com/r/aws/comments/hg5oqm/hashtag_in_aws_cognito_url_callback/)
- url: https://www.reddit.com/r/aws/comments/hg5oqm/hashtag_in_aws_cognito_url_callback/
---
I am creating a Single Page App that uses Oauth2 authentication using AWS Cognito. As a part of the configuration in Cognito I need to supply an url to reach the app after the authentication is performed, this is usually something like [http://domain/login](http://domain/login). The problem is that since the app I am creating is a SPA in React the url should be [http://domain/#/login](http://domain/#login) (note the hashtag) but the hashtag is a forbidden character for Cognito's callback URLs and it doesn't let me configure it that way. Any solutions?
## [7][I swear S3 just corrupted an event linked to a Lambda function. Now it doesn't work, I can't delete it, and I cant create any other event using the same prefix. Arg!](https://www.reddit.com/r/aws/comments/hg03at/i_swear_s3_just_corrupted_an_event_linked_to_a/)
- url: https://www.reddit.com/r/aws/comments/hg03at/i_swear_s3_just_corrupted_an_event_linked_to_a/
---

## [8][Which service combination should I use for hosting many Laravel app containers?](https://www.reddit.com/r/aws/comments/hg4tbt/which_service_combination_should_i_use_for/)
- url: https://www.reddit.com/r/aws/comments/hg4tbt/which_service_combination_should_i_use_for/
---
Hi,

Currently we have a SaaS like setup where you could spinup a new Laravel app on a subdomain.

Present solution is:

* Send a POST request to API Gateway
* Stepfunctions spin up RDS, S3 and EC2
* Register EC2 with Route53
* Send SSM RunCommand to deploy the app
* Send callback request to our "master"/management app.

While this is working fine there are many drawbacks and major underutilisation.

I am thinking of moving the setup to a more cloud friendly infra setup.

Right now reading and going through examples of AppMesh/ECS Ec2 type of setup (in the future Fargate but concepts are the same)

We will have thousands of these small Laravel app services each running on its own subdomain.

Ideally the app would be containerised with just the code and PHP-FPM process. Sidecar container would be Nginx with proxy to PHP-FPM.

How should I approach routing based on virtual hosts here? Could I use AWS ALB for SSL (wildcard) offload only and route using Envoy Proxy.

What are my limits here?

Any thoughts are welcome.
## [9][Serverless GitLab CI/CD on AWS Fargate](https://www.reddit.com/r/aws/comments/hfl31z/serverless_gitlab_cicd_on_aws_fargate/)
- url: https://www.reddit.com/r/aws/comments/hfl31z/serverless_gitlab_cicd_on_aws_fargate/
---
On this really nice article, [**Daniel Coutinho de Miranda**](https://www.linkedin.com/in/ACoAAAMhSJQBtT3u2qPU27N8D3dNyXP5wLFgybU) shows how to serverless run the GitLab Runner Manager and Fargate driver on AWS Fargate.

  
Really nice content!

https://medium.com/ci-t/serverless-gitlab-ci-cd-on-aws-fargate-da2a106ad39c
## [10][AWS Tax](https://www.reddit.com/r/aws/comments/hfyp09/aws_tax/)
- url: https://www.reddit.com/r/aws/comments/hfyp09/aws_tax/
---
How tax works if we consume resources in other countries' data centers? Do we only pay for the local one which we live in?
## [11][Backup from Win Server 2016 to AWS S3](https://www.reddit.com/r/aws/comments/hg0b76/backup_from_win_server_2016_to_aws_s3/)
- url: https://www.reddit.com/r/aws/comments/hg0b76/backup_from_win_server_2016_to_aws_s3/
---
I was wanting to do an automated backup from a window server 2016 to a AWS S3 of a file server.  Looking for recommendations possibly no or low cost.

Thanks in advance
