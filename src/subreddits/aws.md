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
## [2][Introducing the Cloud Development Kit for Terraform](https://www.reddit.com/r/aws/comments/hsio6z/introducing_the_cloud_development_kit_for/)
- url: https://aws.amazon.com/pt/blogs/developer/introducing-the-cloud-development-kit-for-terraform-preview/
---

## [3][CodeDeploy and containers](https://www.reddit.com/r/aws/comments/hsvq9u/codedeploy_and_containers/)
- url: https://www.reddit.com/r/aws/comments/hsvq9u/codedeploy_and_containers/
---
I've been looking at using Code\* as a replacement for self-managed Jenkins, and I was a bit surprised that it doesn't really seem to natively support non-ECS Docker hosts *or* EKS.  

To clarify, I have code in CodeCommit which feeds into CodeBuild to produce a container image and push it into ECR.  I want to take that image and deploy it to either Docker running on an EC2 instance or an EKS cluster.

Googling has turned up a few work-arounds which leverage Lambdas to do the deployment, but they honestly feel pretty kludgey.

Is it just not possible, or did I miss something?  Does anyone have docs or recommended guides on how to do this?
## [4][Lambda function response to API Gateway not working on success](https://www.reddit.com/r/aws/comments/hsvjfy/lambda_function_response_to_api_gateway_not/)
- url: https://www.reddit.com/r/aws/comments/hsvjfy/lambda_function_response_to_api_gateway_not/
---
Hello,

I have a Lambda function that sends a contact-us email that is handling all statusCodes that need to be returned to API Gateway.

When I get to the final function to actually send an email via SES, the email sends and I receive it, but API Gateway returns an Internal Server Error after timing out after 10 seconds. My function executes in 1-2 seconds (depending on if it's warmed), so I know the timeout doesn't need to be increased.

Does the final callback have to be different than the previous callbacks? It actually logs "Email successfully sent.", the email arrives, but then it sits there for 10 seconds before it times out. It's not doing the final callback to API Gateway with a 200, so API Gateway responds with an Internal Server Error.

I've tried moving the callback out of the else, tried context.done(event). All of the other callbacks in the function, I have a return in front so that it stops execution of the Lambda. The last callback, I have tried with and without return.

Does anyone spot an error that I'm missing? Thank you so much!

    var DEBUG = false;
    var ses = new AWS.SES({ region: 'us-east-1' });
    ses.sendEmail(emailParams, function (error, sesResponse) {
      if (error) {
        console.log('Error:', error);
        console.log('sesResponse:', sesResponse);
        return callback(null, {
          statusCode: '500',
          headers: apiHeaders,
          body: JSON.stringify({message:'Error sending email.'})
        });
      } else {
        if (DEBUG) console.log('sesResponse:', sesResponse);
        console.log("Email successfully sent.");
      }
      callback(null, {
        statusCode: '200',
        headers: apiHeaders,
        body: JSON.stringify({message:'Email sent.'})
      });
    });
## [5][The DynamoDB Place - A forum for DDB questions](https://www.reddit.com/r/aws/comments/hsjk42/the_dynamodb_place_a_forum_for_ddb_questions/)
- url: https://dynamodbplace.com/
---

## [6][Why to avoid kubernetes:](https://www.reddit.com/r/aws/comments/hscvl4/why_to_avoid_kubernetes/)
- url: https://blog.coinbase.com/container-technologies-at-coinbase-d4ae118dcb6c#62f4
---

## [7][Upload image to AWS S3 made simple — Flutter](https://www.reddit.com/r/aws/comments/hssx1w/upload_image_to_aws_s3_made_simple_flutter/)
- url: https://medium.com//upload-image-to-aws-s3-made-simple-flutter-fd3356b5b5bb?source=friends_link&amp;sk=a032cf101eb66e20a305079f7348c97f
---

## [8][Amazon EKS Is Eating My IPs!](https://www.reddit.com/r/aws/comments/hssr1e/amazon_eks_is_eating_my_ips/)
- url: https://medium.com/better-programming/amazon-eks-is-eating-my-ips-e18ea057e045
---

## [9][Any plans to bundle a JSON schema extension and the new graph database extension with RDS PostgreSQL?](https://www.reddit.com/r/aws/comments/hssjel/any_plans_to_bundle_a_json_schema_extension_and/)
- url: https://www.reddit.com/r/aws/comments/hssjel/any_plans_to_bundle_a_json_schema_extension_and/
---
To be precise:

* https://www.postgresql.org/about/news/2050/
* https://github.com/furstenheim/is_jsonb_valid (or a similar extension developed by AWS)
## [10][Amazon Glacier - Am I doing this right?](https://www.reddit.com/r/aws/comments/hskmza/amazon_glacier_am_i_doing_this_right/)
- url: https://www.reddit.com/r/aws/comments/hskmza/amazon_glacier_am_i_doing_this_right/
---
I've been searching for an option to archive large amounts of data for a good price and I've found the Amazon Glacier service.

I've started uploading 225 GBs of data there.

But I'm confused about the pricing regarding downloading the data back.

Could someone more experienced give me an estimate on how much that would be? I'm very new to AWS stuff.
## [11][Migrate Hive Metastore to AWS Glue](https://www.reddit.com/r/aws/comments/hsqe9s/migrate_hive_metastore_to_aws_glue/)
- url: https://www.reddit.com/r/aws/comments/hsqe9s/migrate_hive_metastore_to_aws_glue/
---
Hi,

I am running a single node hadoop-hive cluster on a VM for test purposes. I want to link the hive metastore to AWS Glue Catalog for persistent storage. I searched but couldn’t find any good resources to do so. Can anyone help?
