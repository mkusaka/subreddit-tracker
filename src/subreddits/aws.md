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
## [2][Introducing AWS CodeArtifact: A fully managed software artifact repository service](https://www.reddit.com/r/aws/comments/h0l5b2/introducing_aws_codeartifact_a_fully_managed/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/06/introducing-aws-codeartifact-a-fully-managed-software-artifact-repository-service/
---

## [3][Dear AWS, stop ruining the freaking console UI [rant]](https://www.reddit.com/r/aws/comments/h09wl5/dear_aws_stop_ruining_the_freaking_console_ui_rant/)
- url: https://www.reddit.com/r/aws/comments/h09wl5/dear_aws_stop_ruining_the_freaking_console_ui_rant/
---
*I need to get this off my chest, and since this is one of the few places online where people that might share my view on this might see it, I figured it's a good place to go off.*

*If someone from AWS is actually reading this, please pay special attention to the last bit on accessibility, because I'm pretty sure most of the frustration is due to that.*

Dear AWS, please STOP ruining the console UI! I'm not the kind of person that hates change just cause I'm stubborn. If you were improving it, power to ya, but you're not. You are busy making the experience worse. I guess I should thank you because I've been telling coworkers for years to use the CLI and that it's better, and now you are going out of your way to prove my point and drive people there. But sometimes it's just simpler to view a dashboard or play around with a new service using the console. Well, it used to be.

Your transition over to the new UI aren't even smooth on some services. Take EC2 for instance. You rolled out the new look for the Autoscaling section, but most of the time when I navigate there I get the old UI with an error message. When I reload the page, the new UI loads and I can see my resources. Next, CloudWatch Logs. WHY THE HECK WOULD YOU MAKE IT LESS USER-FRIENDLY!? Usually you go to view logs when stuff is broken, often production systems, which is stressful enough. Now you've gone and changed the UI and made it worse. Something as stupid as switching between viewing logs as "Text" vs "Row" is now in a sub menu in a drop down, why?

That leads me to my next point, sub menus and drop downs. Everything is in a collapsible element. That's freaking annoying. Sometimes you want to copy some text to share with a colleague, but as soon as you click to highlight, the blooming thing expands or retracts  and moves the element. Ultimately you can do what you want to do, yes, but it takes longer. In high paced, high pressure environments, crap like that is something no one needs.

It's one thing to make something look better, but most people that uses AWS don't care about looks. We want functionality and ease of use. It can look like a dog's breakfast for all we care, it just has to work!!

**Accessibility**

As I said at the start, I'm sure most of my frustrations is because you are making the UI less user-friendly for people with vision problems. You are making it harder for me to do my job, and I really don't need anyone to do that.

The old UI was basic, simple, and it was really clear where one section ended and another started. There was less collapsable elements and hidden menus. Yes, sometimes you had to scroll till your fingers went numb, but at least it didn't require clicking on 4 different little arrows and two sub-menus to get to the info you want.

I highlight text that I want my screen reader to read out loud. But it feels like 70% of the time I try that technique with the new UI it doesn't work. The text is either some kind of link or action button that opens a collapsable element, or the reader doesn't pick it up as text. Now I know the first response to that last one will be "maybe your screen reader is the issue." But why then is it only on your website? I don't know what kind of UI framework you use, but it's not very accessibility-friendly. It's pretty much impossible to read text in a table. It either doesn't read, or it reads the entire table, no matter which cell I'm highlighting. The worst part is that you're now using this same thing for your documentation pages. I'm basically losing my mind cause I can't read the freaking docs!

Then there is the moving of buttons and options and inconsistent UI's. I'm not talking about the UI being inconsistent across services, it's always been like that. That's something I learned to love about the old UI. I'm talking about something like the Lambda console. Select a function and navigate to the "Configuration" tab. All the config sections are full screen-width blocks, except the X-Ray one. In addition to the screen reader, I use a screen zoom function. So I don't see the whole screen. So I basically scrolled up and down and up and down in search of the X-Ray section, thinking I'm not seeing it. Only to find out, nope, that one config block is sitting on the right side of the page, outside the view of the zoom. Again, you could say that's not your problem, but it kinda is. If all the configs were side-by-side, I would be hovering left to right all the way down the page.

The moving of buttons is one of those things that make me want to scream. With the old UI, most of the action buttons is on the left hand side at the top. Now you moved it to the right, but not on all pages. Why? Why would you move something just for the sake of moving it? "It looks better there.", no it doesn't. It looks the same, it's just orange instead of blue and on the right instead of the left. Most people don't know this, but people with vision problems don't read all the menus/buttons. They memorize button names, link text, and the placement of it to speed up their workflow. Now I basically have to start over.

And finally let's get to colors, fonts, and shadows. The old UI, again, was basic. Black text on a white page, when highlighted it was substantially bolder, and when on a button it was Bol white text on a dark blue background. Here and there there was a menu with white text on black backgrounds. Now everything is a much more modern font, which is thinner and harder to read when highlighted since it doesn't get much bolder. Some pages have colors that are so light that's impossible to see white text, and pages are so busy to cram all the info into a single view, that everything just feels cramped and the font feels smaller.

I can go on, but I'd be pretty surprised if anyone made it this far. I also feel a bit better now, even though as soon as I navigate away from here I'm going back to the console and that kinda sucks.

As I said, I'm not a person that hates change. You updated the Support Center to have the new UI, and apart from the fact that I can't use my screen reader to read the table with all the open cases, it's nice. There's not much wrong with that page and you did a good job there. It's still user-friendly, even for me. Yeah the font/color issue is there too, but other than that.

I'm not the kind of person to just bitch and moan about something and not do something about it. This rant must sound like me bitching and moaning, and honestly, if I was allowed to use all the cuss words that came to mind, it probably would sound more like a rant. But I am willing to help wherever I can to help you improve the console experience. If I have to submit all my suggestions or take screen recordings to explain my situation, I'd gladly do that. I'm just not going to do it if it's going to get ignored. Rather ignore this then.

PS: It's not just AWS that's making this mistake. Even the folks here at Reddit made that mistake with their new look. It's impossible for me to use with my assistive technologies, so I'm still using the old UI. Yeah it looks like something that was created 20 years ago, but it works, and that's what matters.
## [4][What the hell is costing me money?](https://www.reddit.com/r/aws/comments/h0wu18/what_the_hell_is_costing_me_money/)
- url: https://www.reddit.com/r/aws/comments/h0wu18/what_the_hell_is_costing_me_money/
---
I don't have any S3 buckets or EC2 instances as far as I can tell yet in the billing section I am being charged for using 'EC2'?

I was only playing around with a test account for half a day.

How exactly do I see what is costing me because I am at a loss.
## [5][AWS RDS Backup question](https://www.reddit.com/r/aws/comments/h0mrao/aws_rds_backup_question/)
- url: https://www.reddit.com/r/aws/comments/h0mrao/aws_rds_backup_question/
---
So, I'm in a bit of a weird situation.  The company I work for is asking that on top of the 7 daily backups that are currently scheduled in RDS, they would also like to take a backup once a week and keep 13 of them.

What it boils down to is running multiple backup schedules with different retention periods on the RDS databases.

I've had several thoughts on this... even setting up a cron job to do a snap shot that exports to S3 and then setup a lifecycle policy on the storage... but I'm curious if anyone else has any ideas on how to tackle this?
## [6][Does Terraform essentially compile down to CloudFormation?](https://www.reddit.com/r/aws/comments/h0q76p/does_terraform_essentially_compile_down_to/)
- url: https://www.reddit.com/r/aws/comments/h0q76p/does_terraform_essentially_compile_down_to/
---
I noticed Serverless just generates CF stacks. Does Terraform work the same way or does it tie into AWS at a lower level?
## [7][CloudFormation Parameters file formats for AWS CLI vs CodePipeline](https://www.reddit.com/r/aws/comments/h0fjxe/cloudformation_parameters_file_formats_for_aws/)
- url: https://www.reddit.com/r/aws/comments/h0fjxe/cloudformation_parameters_file_formats_for_aws/
---
We're a big CloudFormation shop, and we have some pretty robust pipelines to handle deployment of many layers of infrastructure code via CodePipeline.

When launching CloudFormation via CodePipeline using the CloudFormation ActionTypeId, you can provide a JSON file from the artifacts like so:

      Configuration:
        ActionMode: REPLACE_ON_FAILURE
        RoleArn: !Ref Role
        StackName: "mystack"
        TemplateConfiguration: "source::mytemplate.yaml"
        TemplatePath: "source::parameters.json"


That parameters.json file looks like this:


    {
        "Parameters": {
            "MyParam": "foo",
            "OtherParam": "bar"
        }
    }


However, if you're running through the CLI, something we'd typically be doing while testing out ideas and local builds...



    aws cloudformation create-stack \
        --stack-name mystack \
        --template-body file://mytemplate.yaml \
        --parameters file://parameters.json


That parameters.json file needs to be like:


    [
        { "ParameterKey": "MyParam", "ParameterValue": "foo" },
        { "ParameterKey": "OtherParam", "ParameterValue": "bar" }
    ]


Has anyone managed to figure out a way to make this a bit easier on the development process?

Ideally I'd love to use the parameters file format that CodePipeline uses for CloudFormation across the board - it's much nicer and also supports tags in the same file.

I'm thinking of pre-commit hook on my git repos to translate the CodePipeline format into the AWS CLI format in some ignored local artifacts path, but that seems like overkill.

We run most things through our pipelines once they're going into environments used by others, but as we're writing things the first time.. being able to just make the CLI call saves a ton of time. The different formats for the parameters just gets in the way though.

I'd love to hear what others are doing.
## [8][Anyway to identify network connectivity outage from VPC flow logs etc](https://www.reddit.com/r/aws/comments/h0opr2/anyway_to_identify_network_connectivity_outage/)
- url: https://www.reddit.com/r/aws/comments/h0opr2/anyway_to_identify_network_connectivity_outage/
---
Hi, 
  I have an EC2 running some service on AWS, this depends on on-premise connectivity.. I was suspecting to have some connectivity issue and started to ping the on premise ip and dump the output to a file.. Issue happened again today and I checked the ping activity, there was indeed an almost 1min disconnect. Is there any other place within AWS I can check this ? I tried to make some sense of vpc flow logs for that VPC where I experienced the issue, but there is nothing unusual there.  

We have a direct connect to on prem, I don't have access to those assets..  AFAIK there is no cloud-watch metrics for those ..

Any suggestions that may help me to collect some more facts before I bother network team ?

Thanks
## [9][Advanced filtering in AWS CloudWatch](https://www.reddit.com/r/aws/comments/h08p73/advanced_filtering_in_aws_cloudwatch/)
- url: https://www.reddit.com/r/aws/comments/h08p73/advanced_filtering_in_aws_cloudwatch/
---
I am using JSON logs in CloudWatch.

I am trying to filter logs in CloudWatch but one of my fields is an array of objects. And I want to check a value of an object with a key "message".

For a sample log entry:

    { "appLog": {
        "payload": {
          "responses": [
            {
              "key": "message",
              "value": "Hello World!"
            }
          ]
        }
    }

I want filter all logs which has "message" key object in `responses`.

I am using this query works in Logs Insights and it works as expected. But If my `responses` array has multiple entries and the "message" key entry is not at the first index, my query fails.

    fields @message | filter appLog.payload.responses.0.key == 'message'

Is there any way of selecting an object with some property irrespective of the index? Something like:

    filter appLog.payload.responses.*.key == 'message'
## [10][Anyone switched from ECS with EC2 to Fargate for 24/7 workloads?](https://www.reddit.com/r/aws/comments/h09m21/anyone_switched_from_ecs_with_ec2_to_fargate_for/)
- url: https://www.reddit.com/r/aws/comments/h09m21/anyone_switched_from_ecs_with_ec2_to_fargate_for/
---
It seems with the recent price reduction, and the optional compute savings plans, Fargate has possibly become much more attractive for 24/7 workloads now.

Just curious about real world experience feedback such as cost, performance and general feedback.

Thinking about making the switch.
## [11][Why are my instances in eu-west-3 geting a "US" (or default ?) answer from Route53 record with georouting ?](https://www.reddit.com/r/aws/comments/h0aizs/why_are_my_instances_in_euwest3_geting_a_us_or/)
- url: https://www.reddit.com/r/aws/comments/h0aizs/why_are_my_instances_in_euwest3_geting_a_us_or/
---
Hi there

I've a Route53 record which has a georouting policy attached. Default to us-east, if continent = Europe then eu-west.

Next to that, I just created a new production VPC / instances in eu-west-3. From the instances there, the georouting answer US, while it should answer EU. From others VPCs (staging, development in other EU zones), my own personnal server in a german datacenter, my laptop in france, the georouting works properly, getting "EU" as an answer

```bash
ᐅ ssh production-eu-west dig +short app.production.company.com
app-us-east.production.company.com.
123.123.123.123 # WRONG
ᐅ ssh production-us-east dig +short app.production.company.com
app-us-east.production.company.com.
123.123.123.123 # OK
ᐅ ssh staging dig +short app.production.company.com # eu-west-2
app-eu-west.production.company.com.
111.111.111.111 # OK
ᐅ ssh development dig +short app.production.company.com # eu-west-1
app-eu-west.production.company.com.
111.111.111.111 # OK
```

I guess there's a messup in the VPC, but it's terraformed using [this module](https://registry.terraform.io/modules/terraform-aws-modules/vpc/). I'm not "fluent" enought in AWS to dig properly in the routing tables, subnets, NAT gateway etc &amp; find an hint about that.

Can anyone help me finding the issue ?
