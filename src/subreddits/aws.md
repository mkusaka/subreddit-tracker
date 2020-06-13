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
## [2][Get custom metric using get-metric-statistics](https://www.reddit.com/r/aws/comments/h869ad/get_custom_metric_using_getmetricstatistics/)
- url: https://www.reddit.com/r/aws/comments/h869ad/get_custom_metric_using_getmetricstatistics/
---
Hi, 

This works:

    aws cloudwatch get-metric-statistics --metric-name 'CPUUtilization' --start-time 2020-06-13T10:00:00Z --end-time 2020-06-13T10:05:00Z --period 3600 --namespace AWS/EC2 --statistics Maximum --dimensions Name=InstanceId,Value=i-0b7765711d215xxxx

I am trying to do the same for a custom metric that I can see in Cloudwatch called "'Memory % Committed Bytes In Use'":

    aws cloudwatch get-metric-statistics --metric-name 'Memory % Committed Bytes In Use' --start-time 2020-06-13T10:00:00Z --end-time 2020-06-13T10:05:00Z --period 3600 --namespace AWS/EC2 --statistics Maximum --dimensions Name=InstanceId,Value=i-0b7765711d215xxxx

That isn't finding it, it just returns:

    {
        "Label": "Memory % Committed Bytes In Use",
        "Datapoints": []
    }

Any idea on how I can get that memory metric? If I try and graph it, I can see it under under the "All Metrics" tab, and to the left of the search bar, I have 

    London | All&gt; CWAgent &gt; AutoScalingGroupName, ImageId, InstanceId, InstanceType, objectname'

If that makes sense?
## [3][We are the AWS ML Heroes - Ask the Experts - June 15th @ 9AM PT / 12PM ET / 4PM GMT!](https://www.reddit.com/r/aws/comments/h7rayy/we_are_the_aws_ml_heroes_ask_the_experts_june/)
- url: https://www.reddit.com/r/aws/comments/h7rayy/we_are_the_aws_ml_heroes_ask_the_experts_june/
---
Hey r/aws!

u/AmazonWebServices here.

Several AWS Machine Learning Heroes will be hosting an Ask the Experts session here in this thread to answer any questions you may have about training and tuning ML models, as well as any questions you might have about Amazon SageMaker or machine learning in general. You don’t want to miss this one!

Already have questions? Post them below and we'll answer them starting at 9AM PT on June 15, 2020!
## [4][Serverless - 0 logs in CloudWatch, but full logs when debugging](https://www.reddit.com/r/aws/comments/h8210j/serverless_0_logs_in_cloudwatch_but_full_logs/)
- url: https://www.reddit.com/r/aws/comments/h8210j/serverless_0_logs_in_cloudwatch_but_full_logs/
---
Hi,

I've managed to do something awesome, and I think other people would want to do it as well - I've 100% disabled CloudWatch logging for all my Lambda functions (using a "logs:*" DENY policy) while still being able to have all logs that are needed when debugging with 0 costs.

Pre-requisites:
- Sentry (either hosted or on-premise)
- a basic understanding of the "logging" library in Python

Reasons to do this:
- as mentioned in my other post (https://www.reddit.com/r/aws/comments/g3k1zr/lambda_logs_to_anything_else_but_cloudwatch/), I had shitload of costs triggered by CloudWatch PutEvents that were killing my AWS budget
- I hate scrolling through the new CloudWatch Web UI
- I'm not interested in paying for another service (eg custom ElasticSearch/Datadog) to query/filter my logs


How to do it:
- create your projects in Sentry and get a DSN (preferrably one for each of your Lambda repositories)
- include the Sentry SDK in your custom Layer
- and then:

```
from sentry import capture_exception
from sentry import capture_message
import logging

def lambda_handler(event, context):
  log = logging.getLogger("my-logger")
  log.debug("Hello, world")
  capture_message("something happened")

```

And that's it.

Now, when an exception is triggered through Sentry, you also get the benefit of having all the logs up until that point, WITHOUT them being stored or even sent to CloudWatch.

Proof -&gt; https://drive.google.com/file/d/12KWGTjizUAfAzGBFHmy0OHUHtVwoBNm3/view?usp=drivesdk


Full disclosure: I'm the CTO of https://rungutan.com , an api-driven SaaS load testing platform written 100% on Serverless technologies
## [5][AWS Anti Patterns - Mixing Accounts](https://www.reddit.com/r/aws/comments/h7jy11/aws_anti_patterns_mixing_accounts/)
- url: https://www.reddit.com/r/aws/comments/h7jy11/aws_anti_patterns_mixing_accounts/
---
Guilty of this myself and seeing many others fall into the same trap, I wanted to make a quick technical resource to highlight the problems with mixing accounts and why its important to avoid this trap from the beginning.

Video is available here: https://youtu.be/A_XnXc-5i8Y
## [6][Can I use AWS WAF with my own web server that is hosted on an EC2 instance?](https://www.reddit.com/r/aws/comments/h7ofot/can_i_use_aws_waf_with_my_own_web_server_that_is/)
- url: https://www.reddit.com/r/aws/comments/h7ofot/can_i_use_aws_waf_with_my_own_web_server_that_is/
---
Or AWS WAF works only with Cloud Front and Application Load Balancer?
## [7][How to Achieve the parallelism for these distributed tasks handling with AWS](https://www.reddit.com/r/aws/comments/h84nqr/how_to_achieve_the_parallelism_for_these/)
- url: https://www.reddit.com/r/aws/comments/h84nqr/how_to_achieve_the_parallelism_for_these/
---
While developing an infra on aws having certain questions and issues I am facing with it. I want to parallelism distributed tasks handling for API.

Scenario- I have an app that is related to banking where every minute  millions of requests generate for scheduling reports. Report are of  several types. Like say Account report, Transaction report, and so on.  which can be further schedule daily, weekly, monthly or custom. It  should be send using email and report attach to email.

Now what I require is I want build an API to handle such request and  schedule the reports. Once it scheduled I want to make sure it execute  the N  
 number of whenever it is scheduled.(parallel) \[The reason I want Parallel execution of all the reports which is  scheduled at same time is because lets say if user(s) schedule report(s)  at 11:00 UTC it should receive the report at the same time. I don't  want user(s) end up receiving reports after 1hr or 2hr or N  
  number of hours. No matter how much reports are scheduled, Ever time it  should execute the report at given time stamp. That's why I am looking  for some solution of handling such case of parallel distributed tasks  handling.\]

I tried -  I am not sure what AWS services would be more reliable for  such kind of situation. I built API using API Gateway with SQS &gt;  Lambda which create event on cloudwatch and execute it with lambda  trigger to send SES(email). (But this approach leave me at certain issues if delaying the report  delivery where N  
 number of users scheduled report at the same time.)

Can anyone help me here with forming some very efficient solution that can handle such cases seamlessly without any delay.
## [8][ACM now allows Route53 Hosted Zone validation via Cloudformation](https://www.reddit.com/r/aws/comments/h7v0hp/acm_now_allows_route53_hosted_zone_validation_via/)
- url: https://www.reddit.com/r/aws/comments/h7v0hp/acm_now_allows_route53_hosted_zone_validation_via/
---
ACM at some point updated their Domain Validation option to include a Hosted Zone ID for a domain name. You can specify a hosted zone for ACM to validate your certificates, so it is fully automated (unless I am wrong). How long has this been a thing? 

[Domain Validation option spec] (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-certificate-domainvalidationoption.html)
## [9][Custom metrics, ASGs, and Cloudwatch alarms](https://www.reddit.com/r/aws/comments/h7xk89/custom_metrics_asgs_and_cloudwatch_alarms/)
- url: https://www.reddit.com/r/aws/comments/h7xk89/custom_metrics_asgs_and_cloudwatch_alarms/
---
Hello,

We have an application running on instances in an ASG that has an issue with memory. To get around this until we have a code fix, I wrote a lambda function to do some checks on the health of the TG, then take the instance out, reboot it, and add it back in. The aim was to trigger this when RAM utilization gets above a certain amount.

The problem is that frustratingly I can't create an alarm for this as it looks like the RAM metric from CWAgent needs to be selected on an instance by instance basis, and these are in ASG.

So do I need another Lambda function to update the Cloudwatch Alarm on a regular basis? Or is there another way?

The only other thing I was thinking of was writing something that runs inside the instance itself, but that could be unreliable...
## [10][How exactly do applications handle auto scaling?](https://www.reddit.com/r/aws/comments/h7rnex/how_exactly_do_applications_handle_auto_scaling/)
- url: https://www.reddit.com/r/aws/comments/h7rnex/how_exactly_do_applications_handle_auto_scaling/
---
Been working with AWS for a few years and am taking my SAA cert in a few days, however with my experience and all the prep I did for the exam, I still don't fully understand how applications know how to work with auto scaling groups.

The common architecture is ELB &lt;-&gt; EC2 Auto Scaling Group &lt;-&gt; Database.  I fully understand how ASGs work to dynamically route traffic between the ELB and database tiers, build/destroy instances, and monitor health of targets, but what I don't understand is how do the actual compute application itself work with this dynamic environment?

It seems to be that there is some magic going on behind the scenes where an application just "knows" how to deal with it (obviously not the case), or somehow the application just doesn't care about the infrastructure.  Do developers have to develop their applications in a way to "work" with auto scaling?  Are there any examples of applications that would not work with ASGs by default or by the nature of the program?
## [11][How to SSM or Similar to Automatically Attach Elastic Graphics on Instance Launch??](https://www.reddit.com/r/aws/comments/h7upw6/how_to_ssm_or_similar_to_automatically_attach/)
- url: https://www.reddit.com/r/aws/comments/h7upw6/how_to_ssm_or_similar_to_automatically_attach/
---
I am looking for a means to have a particular AMI always launch with elastic graphics. The software my client uses (not my software), calls on AWS to create a spot instance fleet with a selected AMI and instance type, but does not provide the option to create and attach elastic graphics to the instances being launched.

&amp;#x200B;

Ideally, I would just append the following to the code that controls launching instances, but I don't have access to the code to add it:

 \--elastic-gpu-specification Type=eg1.2xlarge 

&amp;#x200B;

Unfortunately, the instances are being launched from this cloud management app doesn't support this, and I cannot add in the appropriate flags to create and attach elastic graphics on launch from there. My task is to find a workaround to this particular software's limitation and just have aws automatically do the deed anytime the instance is launched.  

&amp;#x200B;

I have been trying to find a means via a cloudwatch event -&gt; SSM function, but it seems that there are only means to set up monitoring.  


Any ideas here?  Thanks.
