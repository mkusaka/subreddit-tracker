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
## [2][At my wits end with trying to access an RDS Postgres DB from home](https://www.reddit.com/r/aws/comments/gv6d9n/at_my_wits_end_with_trying_to_access_an_rds/)
- url: https://www.reddit.com/r/aws/comments/gv6d9n/at_my_wits_end_with_trying_to_access_an_rds/
---
I created an RDS PG DB on a company account created for me with full root access. I set the RDS to Public, set inbound/outbound rules on the security group to allow All Traffic on 0.0.0./0 (I know this is crazy but I just need to get it working first). I attached an Internet Gateway to the VPC (default). I've done everything recommended by the docs to make the RDS accessible to a Rails app on my laptop for development. 

&amp;#x200B;

Whenever I try to either test it via \`psql -h &lt;end point&gt; -U &lt;account username&gt; -p 5432 --password\` and enter the password, it timesout and asks if I'm sure it's running on 5432 and that it's taking IP/TCP traffic. 

&amp;#x200B;

Even if I try to ping it with Telnet via \`telnet &lt;endpoint&gt; &lt;port&gt;\`, same thing. 

&amp;#x200B;

My endpoint is in this format (where it says 'endpoint' in the security description of RDS DB instance): 

&amp;#x200B;

my-postgres-db[.cy5ehlf0towl.us-east-1.rds.amazonaws.com](https://tml-dashboard-development.cy5ehlf0towl.us-east-1.rds.amazonaws.com)

&amp;#x200B;

Can anyone give any other clues? I've spent 3 days on this and it's driving me crazy.
## [3][Detecting inventory stock + SMS Text Notifications with AWS](https://www.reddit.com/r/aws/comments/gust4m/detecting_inventory_stock_sms_text_notifications/)
- url: https://www.reddit.com/r/aws/comments/gust4m/detecting_inventory_stock_sms_text_notifications/
---
Hey all,

I've recently been dealing with some frustration trying to purchase some seemingly common products off online retailers. Getting sick of seeing 'out of stock' messages, I decided to build a script that would automatically detect when a retailer receives new inventory of a product, and send me a SMS Text Notification.

It ended up working really well for the first item - I got a notification in less than 30 minutes and was able to nab one. Not so good luck on a second product - been running it for over 10 hours and no dice.

I put together a video showing how I did it: https://youtu.be/6ixBJZ2vnYk

I hope this shows some practical applications of using AWS and software in general. 

Cheers
## [4][[bug] Impossible to change value of SSM parameter](https://www.reddit.com/r/aws/comments/gv5du9/bug_impossible_to_change_value_of_ssm_parameter/)
- url: https://www.reddit.com/r/aws/comments/gv5du9/bug_impossible_to_change_value_of_ssm_parameter/
---
[Screenshot: parameter name is not set and is impossible to set](https://imgur.com/I6Eknv6)

As you can see above, trying to edit a parameter from SSM results in a page that's impossible to use.

Couldn't find any way to report this bug otherwise... which can also be improved.
## [5][Windows DHCP server](https://www.reddit.com/r/aws/comments/gv37ri/windows_dhcp_server/)
- url: https://www.reddit.com/r/aws/comments/gv37ri/windows_dhcp_server/
---
Hi 

I'm sure there are many other ways of fulfill this requirement but our management team would like to uplift our current on prem Windows DHCP Server and move this into AWS as a EC2 instance

Has anyone or does anyone have their Windows DHCP server running out of AWS ?

The Windows Server configured for DHCP will service our office users computers, this will not affect any AWS servers

The plan is to update the ip helper address which will point to the new DHCP server in AWS, on our core switch so that clients know where to go when looking for a IP address
## [6][Why redirect on AWS workmail aliases dont work ?](https://www.reddit.com/r/aws/comments/gv7815/why_redirect_on_aws_workmail_aliases_dont_work/)
- url: https://www.reddit.com/r/aws/comments/gv7815/why_redirect_on_aws_workmail_aliases_dont_work/
---
Creating rules:

*When the message is received to myalias@mydomain* 

then

*Then copy the messeg to folder (etc)*

&amp;#x200B;

It is not working
## [7][Amazon FSx for Windows File Server – Storage Size and Throughput Capacity Scaling](https://www.reddit.com/r/aws/comments/gv76b6/amazon_fsx_for_windows_file_server_storage_size/)
- url: https://aws.amazon.com/blogs/aws/amazon-fsx-for-windows-file-server-storage-size-and-throughput-capacity-scaling/
---

## [8][Fast File Transfer from EC2](https://www.reddit.com/r/aws/comments/guyld3/fast_file_transfer_from_ec2/)
- url: https://www.reddit.com/r/aws/comments/guyld3/fast_file_transfer_from_ec2/
---
I'm a graduate student with access to an EC2 instance. I'm about to run an experiment that is likely to generate 50-100 GB of data. I've noticed in the past that even small files take a very long time to transfer from my AWS instance to my laptop via \`scp\` or \`rsync\`. Is there a better way to get large files off of an EC2 instance?
## [9][SSM Patch manager and RHEL 8 - is it possible?](https://www.reddit.com/r/aws/comments/gv4rjf/ssm_patch_manager_and_rhel_8_is_it_possible/)
- url: https://www.reddit.com/r/aws/comments/gv4rjf/ssm_patch_manager_and_rhel_8_is_it_possible/
---
I've gotten a task that is to set up automatic patching on some servers, we use "Red Hat Enterprise Linux 8" machines and to my surprise the System manager patch manager seems to not support this. 

I tested it on the RHEL 7, and it works perfectly. From trouble shooting i found out it has something todo with how the patchmanager uses yum from a python script (  i think ), and it chokes when trying to do "import yum". 

The patch manager is valuable for our organisation and it would be a bummer if we cant figure out how to get it to work on these machines.

Does anyone have experience with this or something similar?
## [10][Deploying a AWS Lambda written on F# from cloudformation](https://www.reddit.com/r/aws/comments/guzf30/deploying_a_aws_lambda_written_on_f_from/)
- url: https://www.reddit.com/r/aws/comments/guzf30/deploying_a_aws_lambda_written_on_f_from/
---
Hello all,

I'd like to write and deploy a lambda function written in F#/dotnet; I've found tutorials that allow me to do it manually (something like this):

dotnet tool install -g Amazon.Lambda.Tools

dotnet new lambda.EmptyFunction -lang F# -o FSharpBasicFunction --region us-west-2 --profile default

dotnet lambda deploy-function MyFSharpFunction

yet I need to do it from a cloudformation template and not manually; for this I've found nothing on the internet.

Any idea on how can I archive this?
Thanks!
## [11][AWS workspaces v/s AWS EC2](https://www.reddit.com/r/aws/comments/guxgon/aws_workspaces_vs_aws_ec2/)
- url: https://www.reddit.com/r/aws/comments/guxgon/aws_workspaces_vs_aws_ec2/
---
So we use EC2 instances window's MIA(windows 10 image with all pre build configuration) for our developer's. I was thinking about pushing an idea to my manager about using AWS workspaces desktops environment instance instead of EC2. Anyone here would be able to help with me with pros and cons of using workspaces over EC2? Like scalability,cost,etc.
We do use workspaces for third parties to enable the access for our internal sysytem but that's about it.
