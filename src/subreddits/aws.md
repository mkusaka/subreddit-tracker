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
## [2][Announcing CDK Pipelines Preview, continuous delivery for AWS CDK applications](https://www.reddit.com/r/aws/comments/hrtcpm/announcing_cdk_pipelines_preview_continuous/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/07/announcing-cdk-pipelines-preview/
---

## [3][EC2 instance connect - black screen of death, unable to SSH](https://www.reddit.com/r/aws/comments/hs7gkv/ec2_instance_connect_black_screen_of_death_unable/)
- url: https://www.reddit.com/r/aws/comments/hs7gkv/ec2_instance_connect_black_screen_of_death_unable/
---
I am not able to use Instance Connect to SSH into EC2 instances via the console.

My instances are in US East 1, so I allowed SSH traffic on the defined IP address range of  18.202.216.48/29 , which I pulled from here :  [https://ip-ranges.amazonaws.com/ip-ranges.json](https://ip-ranges.amazonaws.com/ip-ranges.json) 

&amp;#x200B;

However, when I click on the instance and try to Connect (selecting "EC2 Instance Connect (browser-based SSH connection)  " instead of being able to SSH into my instance, I get a pop up / black screen of death that I can't do anything with.  


The black screen eventually times out (maybe 3 min later) and is replaced by the following error message: 

"There was a problem setting up the instance connection

An error occurred and we were unable to connect or stay connected to your instance. If this instance has just started up, try again in a minute or two."
## [4][How To Build A GitOps Pipeline On A Stack Of AWS Services](https://www.reddit.com/r/aws/comments/hs6ukc/how_to_build_a_gitops_pipeline_on_a_stack_of_aws/)
- url: https://www.reddit.com/r/aws/comments/hs6ukc/how_to_build_a_gitops_pipeline_on_a_stack_of_aws/
---
The author is using ArgoCD and other AWS services like CodeCommit, CodePipeline, CodeBuild, Amazon EKS. You can find the complete tutorial --&gt; [here](https://itnext.io/how-to-build-a-gitops-pipeline-on-a-stack-of-aws-services-63f7670b5f95?source=friends_link&amp;sk=f8674e6bcd583995a7b7e78b2cf3ce84) &lt;--
## [5][Amazon Interactive Video Service – Add Live Video to Your Apps and Websites](https://www.reddit.com/r/aws/comments/hrvmlw/amazon_interactive_video_service_add_live_video/)
- url: https://aws.amazon.com/blogs/aws/amazon-interactive-video-service-add-live-video-to-your-apps-and-websites/
---

## [6][AWS S3-User Defined Metadata](https://www.reddit.com/r/aws/comments/hs8f68/aws_s3user_defined_metadata/)
- url: https://www.reddit.com/r/aws/comments/hs8f68/aws_s3user_defined_metadata/
---
Hi ,

I have a question regarding the user defined metadata.

I understand that we need to apply 'x-amz-meta-key:value' using aws sdk , to seperate from HTTP headers and aws s3 stores them in lowercase and returns them in header on GET request.

Is there any document where it explains how aws S3 reads this user defined metadata on PUT/POST?

&amp;#x200B;

Thanks
## [7][Kafka vs SQS](https://www.reddit.com/r/aws/comments/hs87ny/kafka_vs_sqs/)
- url: https://www.reddit.com/r/aws/comments/hs87ny/kafka_vs_sqs/
---
I have an application that uses AWS SQS with Lambda to process the messages pushed on the Queue. The Lambda keeps on polling the Queue, and when a new message appears it process the message.

For this scenario, is it possible to replace the SQS with Kafka on the AWS. In other words, can we use Kafka as a Queue for this use case?
## [8][Terraform's data object equivalent in AWS CloudFormation](https://www.reddit.com/r/aws/comments/hs7q2b/terraforms_data_object_equivalent_in_aws/)
- url: https://www.reddit.com/r/aws/comments/hs7q2b/terraforms_data_object_equivalent_in_aws/
---
Hey there,

could anyone tell me if there is an equivalent of `data` from Terraform in CFN's world? 

I guess it is possible, because CDK often has `fromXXX()` so that you can "import" an existing VPC to your Stack to get some data about it, and CDK is just a wrapper over CFN(kinda).
## [9][CloudFront to Ingress Controler giving 308 redirects](https://www.reddit.com/r/aws/comments/hs5ov2/cloudfront_to_ingress_controler_giving_308/)
- url: https://www.reddit.com/r/aws/comments/hs5ov2/cloudfront_to_ingress_controler_giving_308/
---
In the company we have an ingress controller (nginx), which we use for redirecting HTTP to HTTPS for some services. I wanted to put the URL presented by the ingress controller behind CloudFront by setting the origin of the CloudFront distribution to the URL. However, I am experiencing 308 redirects to the ingress controller URL.  


What could be the problem here?
## [10][Intermittent roblems with apt get on Xenial EC2 instance](https://www.reddit.com/r/aws/comments/hs6knk/intermittent_roblems_with_apt_get_on_xenial_ec2/)
- url: https://www.reddit.com/r/aws/comments/hs6knk/intermittent_roblems_with_apt_get_on_xenial_ec2/
---
Anyone had any issues doing an apt-get install on the ubuntu xenial image (ami-09623236df3ab2b4e) whilst running in EC2.   


I've got a packer job in codebuild that 70% of the time works OK, but, intermittently fails when installing packages. Retrying the build usually succeeds immediately.

    amazon-ebs: Get:22 http://archive.ubuntu.com/ubuntu xenial-backports/universe Translation-en [4476 B]
    amazon-ebs: Fetched 16.4 MB in 4s (3622 kB/s)
    amazon-ebs: Reading package lists...
    amazon-ebs: Reading package lists...
    amazon-ebs: Building dependency tree...
    amazon-ebs: Reading state information...
    amazon-ebs: Package python is not available, but is referred to by another package.
    amazon-ebs: This may mean that the package is missing, has been obsoleted, or
    amazon-ebs: is only available from another source
    amazon-ebs:
    ==&gt; amazon-ebs: E: Package 'autoconf' has no installation candidate
        amazon-ebs: Package pkg-config is not available, but is referred to by another package.
        amazon-ebs: This may mean that the package is missing, has been obsoleted, or
        amazon-ebs: is only available from another source
    ==&gt; amazon-ebs: E: Package 'bison' has no installation candidate
    ==&gt; amazon-ebs: E: Package 'build-essential' has no installation candidate
    ==&gt; amazon-ebs: E: Unable to locate package libyaml-dev
        amazon-ebs:
        amazon-ebs: Package libreadline-dev is not available, but is referred to by another package.
    ==&gt; amazon-ebs: E: Package 'libreadline-dev' has no installation candidate
        amazon-ebs: This may mean that the package is missing, has been obsoleted, or
        amazon-ebs: is only available from another source
    ==&gt; amazon-ebs: E: Unable to locate package libncurses5-dev
        amazon-ebs:
    ==&gt; amazon-ebs: E: Package 'libffi-dev' has no installation candidate
        amazon-ebs: Package libffi-dev is not available, but is referred to by another package.
        amazon-ebs: This may mean that the package is missing, has been obsoleted, or
        amazon-ebs: is only available from another source
        amazon-ebs:
        amazon-ebs: Package autoconf is not available, but is referred to by another package.
    ==&gt; amazon-ebs: E: Unable to locate package libgdbm-dev
    ==&gt; amazon-ebs: E: Unable to locate package libgeos-3.5.0
    ==&gt; amazon-ebs: E: Couldn't find any package by glob 'libgeos-3.5.0'
    ==&gt; amazon-ebs: E: Couldn't find any package by regex 'libgeos-3.5.0'
    ==&gt; amazon-ebs: E: Unable to locate package libgeos++-dev
        amazon-ebs: This may mean that the package is missing, has been obsoleted, or
    ==&gt; amazon-ebs: E: Couldn't find any package by regex 'libgeos++-dev'
        amazon-ebs: is only available from another source
        amazon-ebs:
    ==&gt; amazon-ebs: E: Unable to locate package libproj-dev
        amazon-ebs: Package build-essential is not available, but is referred to by another package.
        amazon-ebs: This may mean that the package is missing, has been obsoleted, or
        amazon-ebs: is only available from another source
        amazon-ebs:
        amazon-ebs: Package bison is not available, but is referred to by another package.
        amazon-ebs: This may mean that the package is missing, has been obsoleted, or
        amazon-ebs: is only available from another source
        amazon-ebs:
    ==&gt; amazon-ebs: E: Package 'pkg-config' has no installation candidate
    ==&gt; amazon-ebs: E: Package 'python' has no installation candidate
    ==&gt; amazon-ebs: E: Unable to locate package stunnel4

From booting the source image, the apt.sources has 

    deb http://eu-west-2.ec2.archive.ubuntu.com/ubuntu/ xenial main restricted
    deb http://eu-west-2.ec2.archive.ubuntu.com/ubuntu/ xenial-updates main restricted
    deb http://eu-west-2.ec2.archive.ubuntu.com/ubuntu/ xenial universe
    deb http://eu-west-2.ec2.archive.ubuntu.com/ubuntu/ xenial-updates universe
    deb http://eu-west-2.ec2.archive.ubuntu.com/ubuntu/ xenial multiverse
    deb http://eu-west-2.ec2.archive.ubuntu.com/ubuntu/ xenial-updates multiverse
    deb http://eu-west-2.ec2.archive.ubuntu.com/ubuntu/ xenial-backports main restricted universe multiverse
    deb http://security.ubuntu.com/ubuntu xenial-security main restricted
    deb http://security.ubuntu.com/ubuntu xenial-security universe
    deb http://security.ubuntu.com/ubuntu xenial-security multiverse
    

Which seems to be injected by cloud-init; from my packer output this doesn't appear to be happening (see [archive.ubuntu.com](https://archive.ubuntu.com)) but I can't see why that wouldn't be. I'm not suppressing that in my packer builder.

      "builders": [
        {
          "type": "amazon-ebs",
          "profile": "xxxx",
          "access_key": "{{user `aws_access_key`}}",
          "secret_key": "{{user `aws_secret_key`}}",
          "region": "eu-west-2",
          "source_ami": "ami-09623236df3ab2b4e",
          "instance_type": "t3.medium",
          "ssh_username": "ubuntu",
          "ami_name": "xxxx-webapp-{{timestamp}}"
        }
      ],
    

It feels like the upstream is having problems, I'm not sure if forcing the ec2 mirrors is causing a problem or would solve the problem and I'm not sure what to try next.  


Thanks for any advice.
## [11][rds postgres: db restore with geo location column returns string instead of object](https://www.reddit.com/r/aws/comments/hs4u9e/rds_postgres_db_restore_with_geo_location_column/)
- url: https://www.reddit.com/r/aws/comments/hs4u9e/rds_postgres_db_restore_with_geo_location_column/
---
I have an automated restore script which has been running for about 1y, but within the last month'ish something has changed, which is causing the geo request to come back as a string instead of an object. Older databases in the same cluster return correctly with exactly the same value, so i am struggling to work out whats wrong with the column/data after I have done a restore?? 

I have a test nodejs script, which is using sequelize as the ORM to test agains the dbs in isolation. I originally thought it was an issue with the ORM but considering it works on one db and not the other, with exactly the same data, and in the same instance, makes me think there is something wrong with the restore. 

Also of note, this same dump file restored into non RDS instances works fine. 

From the database it "looks" like they are exactly the same but obviously i am missing something? 

query executed by the framework : 

    SELECT "name", "address", "geoPointLocation" FROM "Locations" AS "Locations" WHERE "Locations"."name" = 'this-issue-sucks';

Working database &lt;&lt;

    postgres@working_db=&gt; select ST_IsValid("geoPointLocation") from "Locations" where name = 'this-issue-sucks';
    -[ RECORD 1 ]-
    st_isvalid | t
    
    postgres@working_db=&gt; select st_asText("geoPointLocation") from "Locations" where name = 'this-issue-sucks';
    -[ RECORD 1 ]----------------------------
    st_astext | POINT(-77.0365739 38.8976633)
    
    postgres@working_db=&gt; select "geoPointLocation" from "Locations" where name = 'this-issue-sucks';
    -[ RECORD 1 ]----+-------------------------------------------
    geoPointLocation | 0101000000C7180E3A574253C0E3288AA1E6724340

Failing database &lt;&lt;

    postgres@temp_geo=&gt; select ST_IsValid("geoPointLocation") from "Locations" where name = 'this-issue-sucks';
    -[ RECORD 1 ]-
    st_isvalid | t
    
    postgres@temp_geo=&gt; select st_asText("geoPointLocation") from "Locations" where name = 'this-issue-sucks';
    -[ RECORD 1 ]----------------------------
    st_astext | POINT(-77.0365739 38.8976633)
    
    postgres@temp_geo=&gt; select "geoPointLocation" from "Locations" where name = 'this-issue-sucks';
    -[ RECORD 1 ]----+---------------------------------------------------
    geoPointLocation | 0101000020E6100000C7180E3A574253C0E3288AA1E6724340
