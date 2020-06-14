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
## [2][How to host this django project on AWS](https://www.reddit.com/r/aws/comments/h8sya9/how_to_host_this_django_project_on_aws/)
- url: https://www.reddit.com/r/aws/comments/h8sya9/how_to_host_this_django_project_on_aws/
---
I made a project which includes a django-celery-beat task ([https://github.com/celery/django-celery-beat](https://github.com/celery/django-celery-beat)) which is used in order to query an API every second and store the result in a database (SQLite). Is it straightforward to host this on AWS, given that it's already working as it should locally? I've never done something like this using AWS, so am open to all possible suggestions regarding what's the best way to do it. Should I for instance host it on EC2 or something else?

The project is quite basic - it contains a view which has a live chart (highcharts) that is fed through websockets with the data I'm putting into the DB (that comes from the previously mentioned API). I also have another view which contains a form where the user inputs a date range and then gets a file to download. But my aprehension is coming mainly with how farfetched it will be in respect to the celery beat, given that on my local machine I'm using redis and so on.
## [3][I am getting billed if my RDS instance is stopped?](https://www.reddit.com/r/aws/comments/h8p7yt/i_am_getting_billed_if_my_rds_instance_is_stopped/)
- url: https://www.reddit.com/r/aws/comments/h8p7yt/i_am_getting_billed_if_my_rds_instance_is_stopped/
---
Hi!

I have a situation where I need to access my database for probably like 10 hours out of the whole month. Obviously I don't want to pay for the whole month and I also don't want to spin it up (if it's avoidable) just to access it. Will I be billed if instance is stopped/paused? Or what other solutions can there be for such a case?

Thanks in advance!
## [4][How to flush RAM in Ubuntu EC2 instance?](https://www.reddit.com/r/aws/comments/h8ixph/how_to_flush_ram_in_ubuntu_ec2_instance/)
- url: https://www.reddit.com/r/aws/comments/h8ixph/how_to_flush_ram_in_ubuntu_ec2_instance/
---
I tried commands such as: sync; echo 1 &gt; /proc/sys/vm/drop_caches

But I get permission denied even using sudo.  How can I clear memory on a running EC2.
I’m troubleshooting memory leaks but while I do I need to restart the instance very often and it’s taking much time.
## [5][Safe to use programmatic access over public internet?](https://www.reddit.com/r/aws/comments/h8s8dg/safe_to_use_programmatic_access_over_public/)
- url: https://www.reddit.com/r/aws/comments/h8s8dg/safe_to_use_programmatic_access_over_public/
---
I have a relatively straightforward question; I have sometimes used programmatic access (Access Key/Secure Key) to do stuff over public internet (for hobbyist purposes), and always somewhat assumed that it is safe to do.

However, I'm looking for a real explanation of how this works to prove that it is secure (or quantify how secure it is). Mostly for my own learning due to being somewhat obsessive about security.

* Assuming that you use key rotation and your keys are not exposed in any way.
* When there is no other alternative, do you consider programmatic access a fully secure solution?
* Do you know exactly how it works?
## [6][Making diagram of existing cloud infrastructure](https://www.reddit.com/r/aws/comments/h8p6rj/making_diagram_of_existing_cloud_infrastructure/)
- url: https://github.com/Cloud-Architects/cloudiscovery
---

## [7][Aggregate result of all lambda invocation](https://www.reddit.com/r/aws/comments/h8ormf/aggregate_result_of_all_lambda_invocation/)
- url: https://www.reddit.com/r/aws/comments/h8ormf/aggregate_result_of_all_lambda_invocation/
---
My use case is that I need ti check how many objects from the list are missing  in s3, and return number of missing documents. Note that the list size can be of milion order of magnitude. 

From what I gather I have to use HeadObject but doing that sequentially would be too slow. So I could send all messages to the queue and have each lambda head few objects. However I'm stuck on how to aggregate that data further, i.e. how to sum results from  all parallel executions. 

I've seen that maybe Kinesis can be used to this effect, but I never worked with it so I don't know if it will fit my use case.
## [8][Is there a way to implement a system where an IAM admin can allow or ignore every resource an IAM user tries to create?](https://www.reddit.com/r/aws/comments/h8tmwl/is_there_a_way_to_implement_a_system_where_an_iam/)
- url: https://www.reddit.com/r/aws/comments/h8tmwl/is_there_a_way_to_implement_a_system_where_an_iam/
---
Hi, is there a way in AWS where an IAM user will request the AWS IAM Admin User to grant him permission to create a particular resource?

For example, John (IAM User) wants to create a t2.large EC2. But to do that, he will need to send a request to Mark (AWS Admin User). Mark will review the request and can either grant him permission to create t2.large EC2 or can ignore it. Is this possible? Thanks.
## [9][In S3 bucket, uploading files with different extensions like .glb,.gltf or USDZ are getting converted into .bin file.](https://www.reddit.com/r/aws/comments/h8g5zt/in_s3_bucket_uploading_files_with_different/)
- url: https://www.reddit.com/r/aws/comments/h8g5zt/in_s3_bucket_uploading_files_with_different/
---
When I am trying to upload it directly from AWS console, its working but not with PHP Laravel.  
I m having trouble with 3D files via Laravel file uploader,   


files that are working  
.html, mp4, jpeg, jpg, png and more.
## [10][Restrict API gateway by IP for POST only](https://www.reddit.com/r/aws/comments/h8q9wy/restrict_api_gateway_by_ip_for_post_only/)
- url: https://www.reddit.com/r/aws/comments/h8q9wy/restrict_api_gateway_by_ip_for_post_only/
---
Hey all. Please let me know if this isn’t the right place to ask. Is it possible to restrict an API gateway that’s running an express app using proxy+ to allow GETs from anywhere but to restrict POSTs to a certain set of IPs? Would they need to be separate resources to support this? Thanks in advance
## [11][Help with separating data in AWS Glue](https://www.reddit.com/r/aws/comments/h8bft9/help_with_separating_data_in_aws_glue/)
- url: https://www.reddit.com/r/aws/comments/h8bft9/help_with_separating_data_in_aws_glue/
---
I have a support ticket out but I might as well exhaust all resources. 

To try and give a brief overview of what I'm doing, I have a table of records that have 2 different IDs. One ID is a unique identifier, the other ID is used to group 1-3 records.  
EX:  
* recordID: 1, groupID: abc
* recordID: 2, groupID: abc
* recordID: 3, groupID: def  

I don't really care about recordID, I want to get all unique groupIDs. In MySQL where I come from, that's a simple GROUP BY, but I'm not sure how to do it in AWS. My Glue job already separated out the data nicely how I want it, but it's currently going to give me duplicate records without a group by or something similar. How do I separate them out?
