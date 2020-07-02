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
## [2][Does anyone else feel that Step Functions have great potential, but the implementation was half-arsed, so they're not very practical?](https://www.reddit.com/r/aws/comments/hju8or/does_anyone_else_feel_that_step_functions_have/)
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
## [3][DDB-Table v1.0: Safe TypeScript schemas for DynamoDB](https://www.reddit.com/r/aws/comments/hjtaj6/ddbtable_v10_safe_typescript_schemas_for_dynamodb/)
- url: https://github.com/neuledge/ddb-table
---

## [4][Announcing the Porting Assistant for .NET](https://www.reddit.com/r/aws/comments/hjjfnv/announcing_the_porting_assistant_for_net/)
- url: https://aws.amazon.com/blogs/aws/announcing-the-porting-assistant-for-net/
---

## [5][What are some cool things you can do with dead-letter queues (DLQ)?](https://www.reddit.com/r/aws/comments/hjcwhz/what_are_some_cool_things_you_can_do_with/)
- url: https://www.reddit.com/r/aws/comments/hjcwhz/what_are_some_cool_things_you_can_do_with/
---
I've learned that SNS, SQS, and Lambda support DLQs for when an event fails all processing attempts or expires without being processed. The obvious use case is to use DLQs to implement some kind of retry logic.

I'm sure there are other interesting use cases for DLQs as well. What kinds of interesting things can you do with DLQs?
## [6][EKS Fargate pricing question](https://www.reddit.com/r/aws/comments/hjvdqz/eks_fargate_pricing_question/)
- url: https://www.reddit.com/r/aws/comments/hjvdqz/eks_fargate_pricing_question/
---
It appears that pricing is based on vCPU/memory *requests* but I can't find any mention of *limits* to the pods. Are limits automatically set the same as requests? If not....How can that be fair to AWS?
## [7][How to invalidate cache of an S3 bucket object?](https://www.reddit.com/r/aws/comments/hjv44n/how_to_invalidate_cache_of_an_s3_bucket_object/)
- url: https://www.reddit.com/r/aws/comments/hjv44n/how_to_invalidate_cache_of_an_s3_bucket_object/
---
Hi, since I can manually set metadata cache-control max-age on an s3 bucket object using a browser, I was wondering how can I manually invalidate that same object?
## [8][Serving private content with cloudfront](https://www.reddit.com/r/aws/comments/hjpi90/serving_private_content_with_cloudfront/)
- url: https://www.reddit.com/r/aws/comments/hjpi90/serving_private_content_with_cloudfront/
---
Today I am setting up a cloudfront distribution serving documents from a s3 bucket. I am intending to use cloudfront signed urls, but I am really shocked that you need to use the root account to generate/upload the cloudfront keys to do url signing. 


What I expected to be able to do was either call a service to generate a signed url, or store the key in secrets manager that the application can retrieve in order to sign urls. Also being able to generate / rotate the keys using iac.


I'm curious if anyone else is using this feature, and how you manage key generation and rotation. Requiring the root account for this seems counter intuitive to aws best practices.
## [9][How to run/re-run subset of jobs AWS Glue workflow?](https://www.reddit.com/r/aws/comments/hjtcqv/how_to_runrerun_subset_of_jobs_aws_glue_workflow/)
- url: https://www.reddit.com/r/aws/comments/hjtcqv/how_to_runrerun_subset_of_jobs_aws_glue_workflow/
---
Hi guys, I was wondering if it's possible to run only a subset of steps on my AWS Glue Workflow in case one of the jobs fail.

Thank you so much in advance.
## [10][Cloudflare Workers vs AWS Lambda@Edge?](https://www.reddit.com/r/aws/comments/hjss5d/cloudflare_workers_vs_aws_lambdaedge/)
- url: https://www.reddit.com/r/aws/comments/hjss5d/cloudflare_workers_vs_aws_lambdaedge/
---
 What is the difference between these two services? They appear to do the exact same thing.

Are there certain things which one is able to do that the other is not?
## [11][Pretty new to AWS Have some questions about the usage of public IPs and general instance management](https://www.reddit.com/r/aws/comments/hjsf5x/pretty_new_to_aws_have_some_questions_about_the/)
- url: https://www.reddit.com/r/aws/comments/hjsf5x/pretty_new_to_aws_have_some_questions_about_the/
---
Hi! I'm currently starting my journey in the AWS management certification world.

I've so far taken the most basic (free) courses and have been managing a pretty simple account for some months now.

I have a couple of questions that would probably be no-brainers for experienced AWS admins, so please, bear with me.

1. I keep having to request more public IP addressess everytime I create new VMs for different clients (I've already reached my max number of public IP address). So far, it's just pretty simple ~~VMs~~  instances that I plan to optimise and improve once I get my certs and my knowledge on place but until then.. how do you manage with so many public IPs? - I understand I should assign a new different public IP addesss for each client at least but, how to manage, for example, multiple services that need public visibility for the same client? Should I set up a load balancer so it redirects traffic depending on the dns the request comes through and set up a unique public IP address?
2. Probably a stupid question but... it wouldn't be the first time I mess up an ~~VM~~ instance's net configuration (bear with me, I come from VMware on-premise world) -  
I'm used to having the direct console, so I can "attach" myself to the VM and just operate as if I was in front of the VM's screen but can't for the life of me, find the way to do so on AWS instances... I understand this is due to the different approach of the whole concept of instances VS VMs... but I still wouldn't know how to deal with issues like these: I rely on the premise that the networking works as intended and don't know how to deal with this kind of issues.

I know I'm still pretty green on some basic concepts regarding the whole AWS infrastructure world and usage but I'm trying my best to gain the knowledge ASAP. :)
