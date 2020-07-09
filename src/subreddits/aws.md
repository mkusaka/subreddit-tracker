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
## [2][Full dark mode is now available for AWS documentation (text and code). To switch modes, choose Preferences](https://www.reddit.com/r/aws/comments/hnrdgd/full_dark_mode_is_now_available_for_aws/)
- url: https://docs.aws.amazon.com
---

## [3][AWS User Data is Being Stored &amp; Used by Amazon Outside Users' Chosen Regions](https://www.reddit.com/r/aws/comments/ho2dts/aws_user_data_is_being_stored_used_by_amazon/)
- url: https://www.cbronline.com/news/aws-user-data
---

## [4][Cognito "Forgot password" email is customized but I don't know how](https://www.reddit.com/r/aws/comments/ho0ium/cognito_forgot_password_email_is_customized_but_i/)
- url: https://www.reddit.com/r/aws/comments/ho0ium/cognito_forgot_password_email_is_customized_but_i/
---
Hello! I already tried posting on the AWS forums but seeing how very few posts recieve answers there, I thought I might try my luck here as well, so here it goes:

&amp;#x200B;

I'm going crazy and I don't know where else to look, please help me.  
In our codebase, whenever one calls the **ForgotPassword** api, a mail containing the 6 digits code to reset the password is sent to that user. So far so good, everything works as expected.  
The problem is that this e-mail is customized, and badly so: it's not good loking in general, and some mail clients even show a white code over a white background, so we really need to change that.  
It has been this way since before I started working here, and whoever set this up is long disappeared.  


The problem is that I don't know HOW they did it. In the screen **Message customizations** I can only change the *email verification* and *user invitations* messages. In **Triggers** there are no lambdas AT ALL. But the mail has to be customized somehow, right? How is this possible?
## [5][Access AWS server through corporate VPN](https://www.reddit.com/r/aws/comments/ho2d8r/access_aws_server_through_corporate_vpn/)
- url: https://www.reddit.com/r/aws/comments/ho2d8r/access_aws_server_through_corporate_vpn/
---
Hi,
I am very new to AWS, but something has come up at work that requires an urgent solution. We have a windows server in AWS that I can access from my desktop at work. But I am not able to access it when I am on my corporate VPN through my work laptop. We use Cisco AnyConnect SSL VPN.

What needs to be done to make the AWS server available through corporate Cisco VPN. Is it something to do with AWS or something needs to be changed on the Cisco VPN?

Thanks for all your help.
## [6][What are the best apps to help me monitor/manage my AWS implementation?](https://www.reddit.com/r/aws/comments/ho2a91/what_are_the_best_apps_to_help_me_monitormanage/)
- url: https://www.reddit.com/r/aws/comments/ho2a91/what_are_the_best_apps_to_help_me_monitormanage/
---
I like the AWS console app.

I was quite underwhelmed with the QuickSight app.

Are there others (published by Amazon...or others) which can help me manage/monitor my AWS plant/ops when I'm not at a computer? 

(For context, I use EC2, RDS, Route 53, SNS, Cloudwatch)
## [7][Can I create EBS snapshot for running instance volume without corruption risks?](https://www.reddit.com/r/aws/comments/ho29uj/can_i_create_ebs_snapshot_for_running_instance/)
- url: https://www.reddit.com/r/aws/comments/ho29uj/can_i_create_ebs_snapshot_for_running_instance/
---
Can I create EBS snapshot for running instance volume without corruption risks or do I have to stop instance to ensure that everything will be successful?
## [8][CI/CD For a static website on S3](https://www.reddit.com/r/aws/comments/hne21x/cicd_for_a_static_website_on_s3/)
- url: https://www.reddit.com/r/aws/comments/hne21x/cicd_for_a_static_website_on_s3/
---
Hi all , newbie here.

I just setup a static site with S3 ,CF and route53 , however I figure I need to update the site here and there.

Was just wondering what’s the best and less tedious way to setup a CI/CD for this? Do I use GitHub actions and Sync with the S3 , or is there an AWS specific tool I can use for this?
## [9][The AWS Well-Architected Framework Explained](https://www.reddit.com/r/aws/comments/hneuvo/the_aws_wellarchitected_framework_explained/)
- url: https://www.reddit.com/r/aws/comments/hneuvo/the_aws_wellarchitected_framework_explained/
---
# The AWS Well-Architected Framework Explained

Amazon Web Services (AWS) is the leading provider on today´s cloud market. In order to provide a better guidance for modern architects and developers, AWS has published its own architecture framework with its latest release in July 2019. In contrast to TOGAF, it is less holistic, provides less concrete guidance, but it is also much better designed to support in the era of cloud, connectivity, and integration platforms. It does this by addressing general design principles and by linking architectural decisions to concrete business impacts. It also provides five pillars to achieve a well-architected framework, which are operational excellence, security, reliability, performance efficiency, and cost optimization. 

# The General Well-Architected Framework Design Principles are as follows: 

1. **Stop guessing your capacity needs &gt;&gt;** Only decide on capacities when you are certain  
2. **Test systems at production scale &gt;&gt;** Cloud test systems on demand are much cheaper than testing on premises  
3. **Automate to make architectural experimentation easier &gt;&gt;** Automation where possible enhances cost and time savings, tracking of changes and impacts, and reverting to previous versions  
4. **Allow for evolutionary architectures &gt;&gt;** Evolvement over time enables an architecture to become business demand-driven and the business can react to innovations and changed requirements or market situations  
5. **Drive architectures using data &gt;&gt;** Cloud environments have data everywhere and it is easy to collect, so it should be used e.g. for enhanced decision-making or communication  
6. **Improve through game days &gt;&gt;** Regularly schedule simulated events to test your production environment and better understand how to improve it

In the following, I will provide an overview about the five pillars of the Well-Architected-Framework and the design principles corresponding to every pillar: 

&amp;#x200B;

https://preview.redd.it/h2nwtfv74m951.png?width=1824&amp;format=png&amp;auto=webp&amp;s=8e3632b9d160b2799fdbca65115cc149b1badc39

# AWS Operational Excellence

* **Perform operations as code &gt;&gt;** Using codes limits human error and enables a consistent response to events
* **Annotate documentation &gt;&gt;** The creation can be automated in a cloud environment and its benefits should be exploited
* **Make frequent, small, reversible changes &gt;&gt;** Changes should be small enough to be reversible again and they should have no big impact in case of failure
* **Refine operations procedures frequently &gt;&gt;** This can be done on “Game Days” as described in the general design principles section
* **Anticipate failure &gt;&gt;** Anticipate sources for failure, what scenarios there could be, and how to react
* **Learn from all operational failures &gt;&gt;** Share lessons learnt from failures across the organization so that everybody benefits from it  
 

# AWS Security

* **Implement a strong identity foundation &gt;&gt;** Give the least privilege as possible and centralize its management
* **Enable traceability &gt;&gt;** Monitor, alert, and audit actions and changes in real time and, if possible, react automatically
* **Apply security at all layers &gt;&gt;** Secure every instance and on every layer instead of only focusing on an outer layer of security
* **Automate security best practices &gt;&gt;** Include software-based security mechanisms and controls as code
* **Protect data in transit and at rest &gt;&gt;** Classify data, encrypt, and add time stamps to protect data
* **Keep people away from data &gt;&gt;** Minimize direct human access to data to minimize human errors
* **Prepare for security events &gt;&gt;** Align to the organizational incident management process and run regular simulations to automate and accelerate the speed of detection, investigation, and recovery  
 

# AWS Reliability

* **Test recovery procedures &gt;&gt;** Cloud enables the exact recreation of scenarios that led to a failure, so that different strategies can be tested
* **Automatically recover from failure &gt;&gt;** Use KPIs and add automated recovery processes where possible
* **Scale horizontally to increase aggregate system availability &gt;&gt;** Replace large resources with smaller and independent ones to mitigate the risk of one resource failing
* **Stop guessing capacity &gt;&gt;** Only decide on capacities when you are certain; see general design principles
* **Manage change in automation &gt;&gt;** Infrastructure changes should be automated and a “real change” should only occur if an automation itself needs to be changed  
 

# AWS Performance Efficiency

* **Democratize advanced technologies &gt;&gt;** Try to consume modern technologies that you lack the skills for via services from third party providers
* **Go global in minutes &gt;&gt;** Deploy your systems across different regions to provide lower latency to customers
* **Use serverless architectures &gt;&gt;** Minimize the usage of own servers to avoid the need to manage them
* **Experiment more often &gt;&gt;** Try different options for instances, storages, configurations etc. to optimize them
* **Mechanical sympathy &gt;&gt;** Align your technology approach to your goals  
 

# AWS Cost Optimization

* **Adopt a consumption model &gt;&gt;** Adopt your resources to the business demand and requirements
* **Measure overall efficiency &gt;&gt;** Understand what the architecture brings the business
* **Stop spending money on data center operations &gt;&gt;** Use cloud service providers for data center operations work to save resources and focus on projects and business
* **Analyze and attribute expenditure &gt;&gt;** Identify business owners of IT components to enable an improved ROI management
* **Use managed services to reduce cost of ownership &gt;&gt;** Reduce servers for operational tasks like sending emails or managing databases, as service providers can do this at higher scale and lower marginal costs

&amp;#x200B;

https://preview.redd.it/8vzp3fda4m951.png?width=903&amp;format=png&amp;auto=webp&amp;s=288f1678d912a47f1c88e5704a77c11f6b78bfce

# What is The Well-Architected Framework Review Process?

The review process describes in high-level terms, how the assessment of the principles should be done. For AWS, this should be a lightweight process, which is taking rather hours, instead of days and it should be repeated multiple times across the architecture lifecycle. AWS states that it is important to have a conversation (not an audit) and a “blame-free approach” during the review and it is also important to involve the right people. The results of the conversations should be a list of issues that can then be prioritized based on the business context and that can be formulated into a set of actions that help to improve the overall customer experience of the architecture.

# Is the AWS Well-Architected Framework Helpful?

Having analyzed the AWS white paper on its Well-Architected Framework, I realized that there are no guidelines on how to do an assessment of the framework. The chapter “Review Process” remains on a very high level and does not include any lists of activities that should be performed or deliverables that should be achieved. Neither are there overviews that help to understand the bigger context of the framework, how to get started, or how to combine it with other frameworks. Overall, the guidance stays very rough and provides only a broad direction. 

Nevertheless, a detailed guidance on the approach and the context of the principles is not the goal of AWS. On the contrary, they succeed in providing a good overview of good principles – as a starting point – for how to build cloud environments. It can also be well used to amend and enrich other architecture frameworks whose focus is not purely on cloud. 

A full version of the AWS Well-Architected Framework is available [here.](https://d1.awsstatic.com/whitepapers/architecture/AWS_Well-Architected_Framework.pdf)
## [10][Do I have to handle OPTIONS methods with my lambda with the new Api Gateway HttpApi](https://www.reddit.com/r/aws/comments/hnwjdd/do_i_have_to_handle_options_methods_with_my/)
- url: https://www.reddit.com/r/aws/comments/hnwjdd/do_i_have_to_handle_options_methods_with_my/
---
When using the old REST api gateway the options call from browsers would be handled in API gateway, I can't seem to figure out how to do this with the new HttpAPI? Surely I dont have to invoke my lambda function just for the options method?
## [11][So with s3 selecting a bucket and going to actions and clicking “get total size” are you charge for a list api call for every file ?](https://www.reddit.com/r/aws/comments/hnoed5/so_with_s3_selecting_a_bucket_and_going_to/)
- url: https://www.reddit.com/r/aws/comments/hnoed5/so_with_s3_selecting_a_bucket_and_going_to/
---

