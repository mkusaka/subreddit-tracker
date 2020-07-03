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
## [2][CloudFront IP Ranges](https://www.reddit.com/r/aws/comments/hkg4ms/cloudfront_ip_ranges/)
- url: https://www.reddit.com/r/aws/comments/hkg4ms/cloudfront_ip_ranges/
---
Hi All,

On one of our applications that is deployed though cloudFront,  when i ping the website it gives me an IP address that matches one of the IP's listed by AWS that are used by CloudFront which is expected.

Looking further into the CloudFront IP ranges you can search for them through the AWS PowerShell Tools, and filter the IP addresses by region for example:

    Get-AWSPublicIpAddressRange -Region global -ServiceKey CLOUDFRONT | select IpPrefix

Searching this gives me a nice list of all the global IP ranges used by CloudFront. However you can search for specific regions by changing the search to "-Region eu-west-2" instead of "-Region global" and it only returns one IP range - [52.56.127.0/25](https://52.56.127.0/25)

&amp;#x200B;

Does anyone know what the difference between the two IP ranges are? i dont understand why they would have region specific IP ranges aswell as global ones since i would assume the global ones could be used by CloudFront in any region.
## [3][Network load balancer with many ports](https://www.reddit.com/r/aws/comments/hkj9kr/network_load_balancer_with_many_ports/)
- url: https://www.reddit.com/r/aws/comments/hkj9kr/network_load_balancer_with_many_ports/
---
Ive been asked to look at adding a load balancer to one of our AWS environments.

The application servers which will sit behind the LB currently listen on ports anywhere in the tcp/6800-6999 range and can change frequently.

I did a few initial checks but when I go to configure an NLB, I am unable to pick a range of ports when configuring listeners. I could manually enter each port but this would take a long time and it got me thinking if the AWS NLB is the right tool for this job?

Would there be a better way of doing this?
## [4][AWS Security Maturity Roadmap for 2020](https://www.reddit.com/r/aws/comments/hk3nm8/aws_security_maturity_roadmap_for_2020/)
- url: https://summitroute.com/downloads/aws_security_maturity_roadmap-Summit_Route.pdf
---

## [5][AWS Fargate for Amazon Elastic Kubernetes Service](https://www.reddit.com/r/aws/comments/hkj7va/aws_fargate_for_amazon_elastic_kubernetes_service/)
- url: https://caylent.com/aws-fargate-for-amazon-elastic-kubernetes-service
---

## [6][A Python -&gt; Step Functions compiler](https://www.reddit.com/r/aws/comments/hk7fqn/a_python_step_functions_compiler/)
- url: https://www.reddit.com/r/aws/comments/hk7fqn/a_python_step_functions_compiler/
---
Hey all, 

I've been working on a Python -&gt; Step Functions compiler, and after seeing /u/mlda065's [post](https://www.reddit.com/r/aws/comments/hju8or/does_anyone_else_feel_that_step_functions_have/), I figured it might be interesting to post a preview/demo here.

It's called [Cohesion.dev](https://cohesion.dev) and I've been working on it for a few months.  There's a web-based demo at [https://preview.cohesion.dev](https://preview.cohesion.dev).

The way it works: it adds a library called Cohesion that lets you call Lambdas, Activities etc. as if they are just Python functions.  Then it parses your code and splits it up into little pieces that it outputs as Lambdas, and outputs a Step Function that stitches it all back together.

For now it's a bit rough around the edges, but if you'd like a preview and to follow along as it improves, please check it out at that link!
## [7][Does Quicksight have an option to add customized objects, like Tableau allows users to develop extensions?](https://www.reddit.com/r/aws/comments/hkj5g5/does_quicksight_have_an_option_to_add_customized/)
- url: https://www.reddit.com/r/aws/comments/hkj5g5/does_quicksight_have_an_option_to_add_customized/
---
The title basically says it. E.g. Tableau cannot render dynamically images, so there are extensions for that. Is there such an option in Quicksight (for developing, not for rendering images).
## [8][Change lambda code with SAM](https://www.reddit.com/r/aws/comments/hkhfiw/change_lambda_code_with_sam/)
- url: https://www.reddit.com/r/aws/comments/hkhfiw/change_lambda_code_with_sam/
---
Hello everyone. Can someone please tell me how can I deploy only code to my SAM's lambdas? When I want to deploy new changes, I have to run sam build and then run sam deploy and it is overkill if I want to make small bug fixes to my lambdas. Is there a way I can change lambda code and not having to go directly to lambda code over aws ui. Thanks.
## [9][Open Source Multi-Cloud Security Assessment Tool ScoutSuite 5.9.0 Released](https://www.reddit.com/r/aws/comments/hk5ccj/open_source_multicloud_security_assessment_tool/)
- url: https://research.nccgroup.com/2020/07/02/tool-release-scoutsuite-5-9-0/
---

## [10][Does anyone else feel that Step Functions have great potential, but the implementation was half-arsed, so they're not very practical?](https://www.reddit.com/r/aws/comments/hju8or/does_anyone_else_feel_that_step_functions_have/)
- url: https://www.reddit.com/r/aws/comments/hju8or/does_anyone_else_feel_that_step_functions_have/
---
Does anyone else feel like Step Functions have a lot of potential, but the implementation was kind of half arsed? There's so much basic functionality lacking. There are so few comparison operators for Choice statements, and error handling is ridiculous.


* AWS documentation says that `States.ALL` catches all errors. But if you reference `$.myvar` and that doesn't exist, it will cause an error which is *not* caught by `States.ALL`. **There is no way to catch all errors*. I had to make a second step function which does nothing but invoke the first one, and then handle an error.
* You cannot catch errors from a Choice task (e.g. from looking up a path that doesn't exist), because `Catch` is not supported for `Choice`
* You cannot compare two variables in a Choice task. You can only compare variable to hard coded constant  
* You cannot check the length of an array in a Choice task  
* You cannot check whether a variable is defined in a Choice task. ie. It's not possible to say `if 'x' in d and d['x'] == 1: foo(d['x']) else: bar()`  
* They say you can use JSONPath, but when I try JSONPath errors I get errors. I've had different support agents give me confliction advice about whether this is possible.  
* AWS recommend adding `TimeoutSeconds` to tasks, and say they can be used on any task. But they're not allowed in Parallel tasks. **The documentation is incorrect.**  
* AWS recommend using Pass tasks as placeholders for Lambdas while you're constructing your step function. But Pass tasks can't take `Retry` or `Catch` fields. This means it's not possible to draft a step function with Pass tasks as placeholders if you have tasks that are only used to handle caught errors.
* You can't write step functions in Yaml. That means you can't have comments, and it's very easy to write invalid code because of a comma trailing at the end of a list.  
* If one branch in a Map or Parallel task fails, all other tasks are immediately cancelled. It's not feasible to configure it to continue executing the other branches. I want to say "for each piece of equipment, send a signal". If piece 500 fails to be actuated, I don't want pieces 1 to 499 to be cancelled mid-way through actuation
* you can't do a dictionary lookup
* you can't decrement a variable to do `for(a = 10; a &gt; 0; a-- )`
## [11][Discount for partial storage month](https://www.reddit.com/r/aws/comments/hkgchl/discount_for_partial_storage_month/)
- url: https://www.reddit.com/r/aws/comments/hkgchl/discount_for_partial_storage_month/
---
Can you please explain **Discount for partial storage month** in the context of this estimate https://calculator.aws/#/estimate?id=ff3c92481f6df0d8d4dc5913eb79fcba1aba8e08 ?

https://s.natalian.org/2020-07-03/1593759825_2560x1440.png
