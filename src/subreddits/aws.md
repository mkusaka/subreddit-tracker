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
## [2][Is it cheaper to run a wordpress server using spot instances and external files(s3) and db behind an LB on Autoscale, than buying a reserved instance?](https://www.reddit.com/r/aws/comments/hh5rmj/is_it_cheaper_to_run_a_wordpress_server_using/)
- url: https://www.reddit.com/r/aws/comments/hh5rmj/is_it_cheaper_to_run_a_wordpress_server_using/
---
We're running a few wordpress sites, some get spikes but generally low traffic. I keep seeing, 'save up to 90% on spot intsances' but it seems that the disappearance of a spot instance would bring the service down. Does anyone have any experience with trying this setup?
## [3][Locked out of EC2 instance](https://www.reddit.com/r/aws/comments/hhdxbg/locked_out_of_ec2_instance/)
- url: https://www.reddit.com/r/aws/comments/hhdxbg/locked_out_of_ec2_instance/
---
I've installed firewalld on a linux instance to make some port configurations and I remember setting up SSH ports permanent. But for some reason I'm locked out and no port responds which makes me wonder if I somehow haven't made the changes permanent.

I can't get into it. Does anyone have an idea? Thanks ahed!
## [4][Calling IOT Shadow Rest APIs with API Gateway?](https://www.reddit.com/r/aws/comments/hh9v2l/calling_iot_shadow_rest_apis_with_api_gateway/)
- url: https://www.reddit.com/r/aws/comments/hh9v2l/calling_iot_shadow_rest_apis_with_api_gateway/
---
I'm working on a project where I'm trying to get my IOT thing's shadow's state and update it from a web application. I'm thinking to doing this by using API gateway. I've created an IAM Role to allow Get/Update requests for that Thing.

I've tried a couple of ways to do this:

1. I set the integration type to HTTP, and set the endpoint to the shadow's URI. The issue with this is I can't set an IAM role that authenticates this request
2. I change the integration type to AWS Service, and then to either IoT or IoT Data, and set the IAM Role, but neither of them work, because I can't set the endpoint.

Not sure how to proceed with this, the only other thing I can think of is using Lambda, but I'd like to know if it's possible to do this without lambda.

Thank you.
## [5][Does the Elastic Beanstalk setup of Nginx have caching or similar enabled by default?](https://www.reddit.com/r/aws/comments/hh0ye2/does_the_elastic_beanstalk_setup_of_nginx_have/)
- url: https://www.reddit.com/r/aws/comments/hh0ye2/does_the_elastic_beanstalk_setup_of_nginx_have/
---
I just deployed my first Python Flask website through Elastic Beanstalk. I've been slowly working through the issues and have mostly resolved them. But now I'm trying to submit data through my API so I can populate my database. I was able to send my initial data for one table in which consisted of about 12500 POST requests. But as soon as I started on my second table's data I was only able to get one POST submitted before I started getting 500 errors. 

In reading the logs it *feels* like something is cached by Nginx but I've tried restarting the application servers with no success getting past the issue.

**Question**: Does the Nginx instance for Elastic Beanstalk do some sort of caching by default that will not allow my data to change for new POST requests?

This is the related application log of note:

* Jun 27 20:01:09 ip-172-31-33-164 web: sqlalchemy.exc.IntegrityError: (pymysql.err.IntegrityError) (1062, "Duplicate entry '**1111**' for key 'PRIMARY'")
* Jun 27 20:01:09 ip-172-31-33-164 web: [SQL: INSERT INTO product (sku, title, media_id) VALUES (%(sku)s, %(title)s, %(media_id)s)]
* Jun 27 20:01:09 ip-172-31-33-164 web: [parameters: {'sku': **2222**, 'title': 'Product1', 'media_id': 3}]

The above is what happens when my API tries writing my POST request to the DB. But for some reason the first highlighted value of **1111** keeps being used for sku field instead of the one sent by the POST request, in this example **2222**. When running my code locally I don't have any issues and can make successful POST requests and input the data. Which is why I believe it's something with the deployment and not the code itself. The code is doing what it should and rejecting the database write but no matter how many requests with different data I send it always uses that same initial sku value of **1111** and won't change. I've looked for a hard-coded value in my code but there is none to be found.

Here is the relevant python code from my API which should be adjusting the sku value every time.

        product = Product(
            sku=request.json['sku'],
            title=request.json['title'],
            media_id=determine_media_id(request.json['media'])
        )

Does anyone have any ideas of what may be going on?
## [6][Trying to send data to API Gateway from Wordpress](https://www.reddit.com/r/aws/comments/hh8rto/trying_to_send_data_to_api_gateway_from_wordpress/)
- url: https://www.reddit.com/r/aws/comments/hh8rto/trying_to_send_data_to_api_gateway_from_wordpress/
---
I am new to using AWS, and I have a wordpress website with a form and I am trying to send the strings from the form to be processed by a Lambda function. I was told I should use an API Gateway which would work with a REST API to send to the Lambda function. I installed the AWS SDK for PHP on my wordpress website, but am lost at how I am supposed to send the data to the API Gateway. Any help is appreciated!
## [7][AWS Rekognition Custom Label model failing](https://www.reddit.com/r/aws/comments/hgzxc4/aws_rekognition_custom_label_model_failing/)
- url: https://www.reddit.com/r/aws/comments/hgzxc4/aws_rekognition_custom_label_model_failing/
---
When I try to train a model using Rekognition Custom Labels, it gives me this error:

&amp;#x200B;

Less than 50% of labels overlap between the training and test datasets.

&amp;#x200B;

&amp;#x200B;

What does this mean and how can I fix it?
## [8][CloudFormation - Get AutoscalingGroup ARN for use in IAM?](https://www.reddit.com/r/aws/comments/hh5wvr/cloudformation_get_autoscalinggroup_arn_for_use/)
- url: https://www.reddit.com/r/aws/comments/hh5wvr/cloudformation_get_autoscalinggroup_arn_for_use/
---
I'm sure this is a stupid question with an easy solution, but I can't seem to figure this out.

I create an autoscaling group in my CF template and want to create an IAM role granting permissions to users to adjust the desired count. The IAM policy expects the ARN of the scaling group, but the autoscaling group resources only returns the Name if you use the Ref function and there is no Get::Att values. Since those names have UUID's in them, you can't  construct a very specific ARN yourself, since you don't know what the UUID is.

How would I set the permissions then? Is there a way to get the ARN?
## [9][Stream Processing Q for twitter data streaming](https://www.reddit.com/r/aws/comments/hh2yyw/stream_processing_q_for_twitter_data_streaming/)
- url: https://www.reddit.com/r/aws/comments/hh2yyw/stream_processing_q_for_twitter_data_streaming/
---
Hello,

I'm in the process of hashing out my ideas for a streaming data personal project. My intention is to capture streaming data from the twitter API based on hashtags that interest me. I have the data capture logic working locally and intend on deploying the py script over to an EC2 instance.

I have a separate py script where I'd like to do the NLP-related tasks in a streaming fashion (e.g.  tokenization, sentiment analysis, bigram frequency, etc) and want to create a real-time dashboard where I can start doing some exploration with simple frequency-based bar charts in JS and put on my resume.

I originally intended on using an Amazon Elasticsearch cluster but ran into a wall when I saw the complexity in sharing my kibana dashboard to the public.

**My main concern is in the streaming ETL stage**, as there are endless options out there, but since i'm within the AWS ecosystem I'd like to use Kinesis. From what I read, kinesis data firehose would be my best bet where I can rapidly load json documents onto S3 and have an event-triggered AWS lambda that picks up new JSON objects as they come in to the S3 store, loads them and possibly calculates the NLP stuff via the lambda function or using a Amazon EMR pyspark job? (kind of lost if this is the best approach for the nlp-stage). Also, I'm a bit lost on why It would be best to use s3 as a data lake and not do the ETL as it's streamed onto kinesis, then store it to something like PostgreSQL on RDS?

With the AWS architecture, is this the best pipeline to use for my use case? Or would it be more efficient to run batch jobs and do the analysis at the end of the day? ultimately I'd like to have a dashboard where I can view people's sentiment on a political hashtag that interests me and thought it would be cool to have this automated and in near-real time.
## [10][Reminder: Terminate your free WorkSpaces by 30-Jun-20](https://www.reddit.com/r/aws/comments/hgl5hz/reminder_terminate_your_free_workspaces_by_30jun20/)
- url: https://www.reddit.com/r/aws/comments/hgl5hz/reminder_terminate_your_free_workspaces_by_30jun20/
---
Just a quick reminder, AWS' free WorkSpaces offer ends 30-Jun-2020. If you were using it to test out the service, please remember to terminate it to avoid additional charges.

Thank you to Amazon for offering this.
## [11][Storage from Cellphone to S3](https://www.reddit.com/r/aws/comments/hh6opq/storage_from_cellphone_to_s3/)
- url: https://www.reddit.com/r/aws/comments/hh6opq/storage_from_cellphone_to_s3/
---
Hello

I want to stop using google photos, but I still want to do the storage of my photos in some cloud service, I was wondering if I could use S3 Glacier Deep Archive, based on the simulation I did, I would pay something like $0.3 month (I just need to upload 35Gb one time, this is the size my backup of photos since 2010).

I don't need to do a lot of retrievals. The "tricky part" is I want to sync my photos from cell phone directly to S3.

I was thinking in using a python script on my phone and using AWS SDK to upload to S3 and after a success upload, delete the photo from my phone.

1. Do you think it's possible to do that?
2. Any other idea?
