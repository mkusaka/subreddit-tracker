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
## [2][A utility to edit secrets from AWS Secrets Manager without storing them](https://www.reddit.com/r/aws/comments/hcj9lf/a_utility_to_edit_secrets_from_aws_secrets/)
- url: https://github.com/zeapo/barberousse
---

## [3][Do I really need to provide resource-level tenant isolation with IAM roles and policies?](https://www.reddit.com/r/aws/comments/hcdv70/do_i_really_need_to_provide_resourcelevel_tenant/)
- url: https://www.reddit.com/r/aws/comments/hcdv70/do_i_really_need_to_provide_resourcelevel_tenant/
---
I'm developing a web app for a SaaS business.

This business is aspiring/planning to have hundreds or thousands of customers.

I've watched some (awesome) talks by Tod Golding on AWS multi-tenant architectures for SaaS businesses.

Tod recommends us to add tenant isolation not only at the token level, but also at the resource level (DynamoDB tables or partition keys, S3 buckets or tags, etc.), by adding policies on those resources that will isolate tenants and also handle permission levels.

After creating these policies, we would utilize the services of Cognito Identity Pools (or AWS STS) to exchange the user token for AWS credentials, through which the users would assume roles.

I have to say: I don't really see the need for that. Because if I have a token with a secret key, I can assume the user is who they are, and then I can just take the user ID, the tenant ID and query my DynamoDB table or S3 bucket according to what I need.

**What precisely other security would such IAM policies add?** Can someone illuminate me?

Because as I see it, I think this will just add unnecessary complexity and make the app slower (demanding extra requests to STS each time I need to access AWS resources).
## [4][I wrote a free app for sketching cloud architecture diagrams](https://www.reddit.com/r/aws/comments/hbztrc/i_wrote_a_free_app_for_sketching_cloud/)
- url: https://www.reddit.com/r/aws/comments/hbztrc/i_wrote_a_free_app_for_sketching_cloud/
---
I wrote a free app for sketching cloud architecture diagrams. All AWS, Azure, GCP, Kubernetes, Alibaba Cloud, Oracle Cloud icons and more are preloaded in the app. Hope the community finds it useful: [cloudskew.com](https://www.cloudskew.com/)

Notes:

1. The app's just a simple diagram editor, it doesn't need access to any AWS, Azure, GCP accounts.
2. You can see some sample diagrams [here](https://www.cloudskew.com/docs/samples.html).

[CloudSkew - Free AWS, Azure, GCP, Kubernetes diagram tool](https://preview.redd.it/9jm111zn1v551.png?width=1438&amp;format=png&amp;auto=webp&amp;s=c33c6eb8c76a0c52408e0c672d36b6eac62a3fed)
## [5][Best ec2 instance type(s) for remote full-stack development?](https://www.reddit.com/r/aws/comments/hcl879/best_ec2_instance_types_for_remote_fullstack/)
- url: https://www.reddit.com/r/aws/comments/hcl879/best_ec2_instance_types_for_remote_fullstack/
---
Was using a t2.small for remote dev through vscode but the performance was abysmal.  Any sort of load on the device seemed to caused very high response times or kick me out ssh(even through the aws webconsole).  However the status checks on the server never showed any issues.  The rest of my internet is fine and I've never had this issue on DigitalOcean.  

Using a docker-compose MERN stack.

I'm wondering if anyone can recommend a good ec2 instance type for full-stack remote development with vscode(via ssh)?  Or at least what to avoid?  8GB of memory is the most I would ever need I think.
## [6][How do I make boto3 output more readable?](https://www.reddit.com/r/aws/comments/hcmalq/how_do_i_make_boto3_output_more_readable/)
- url: https://www.reddit.com/r/aws/comments/hcmalq/how_do_i_make_boto3_output_more_readable/
---
Hello, 

If I get an output from something like ec2.describe\_instances(), I just get an enormous wall of text. How can I format it to make it more readable in the terminal? I have tried playing around with json (output=json.loads(response) for example, but can't seem to find something that works.
## [7][Amazon Glacier Stack for Synology](https://www.reddit.com/r/aws/comments/hcm6wa/amazon_glacier_stack_for_synology/)
- url: https://www.reddit.com/r/aws/comments/hcm6wa/amazon_glacier_stack_for_synology/
---
Another weekend project, this time I made a CloudFormation file for anyone with a Synology device who would like to have an offsite backup using Amazon Glacier. As always the CF file is setup as one click deployment. 

Worth knowing: the downside of the Synology implementation of Glacier is that they don't ask for a Vault ARN, they make one for you, so there is no way to configure the Vault using a CF file. Not to mention that they need way more privileges than needed. Thankfully they have a standard naming scheme, so I locked Synology down with a policy that has a wildmaks resource ARN.

This is the URL of the project: https://github.com/0x4447/0x4447_product_synology_backup
## [8][Elastic Beanstalk - Command '/opt/python/run/venv/bin/pip install -r /opt/python/ondeck/app/requirements.txt' returned non-zero exit status 1](https://www.reddit.com/r/aws/comments/hck7u0/elastic_beanstalk_command_optpythonrunvenvbinpip/)
- url: https://www.reddit.com/r/aws/comments/hck7u0/elastic_beanstalk_command_optpythonrunvenvbinpip/
---
Trying to use a git link in order to install a package which is on github. For this, I included the following line in my requirements.txt file:

        [-e] git+https://github.com/deribit/deribit-api-clients#egg=python

But I get the following error when I run \`eb deploy -v\`

        CalledProcessError: Command '/opt/python/run/venv/bin/pip install -r /opt/python/ondeck/app/requirements.txt' returned non-zero exit status 1.

What's the right way to list a requirement like that? Pip version of the instance is 9.0.3
## [9][how do i change Node.js version , Proxy server,Listeners from config or yml? instead of going over Configuration overview in aws web UI](https://www.reddit.com/r/aws/comments/hcjyiu/how_do_i_change_nodejs_version_proxy/)
- url: https://www.reddit.com/r/aws/comments/hcjyiu/how_do_i_change_nodejs_version_proxy/
---
currently i usually end up doing config over web UI

Node.js version, Proxy server,  Listeners,   Instance type
## [10][Question about lightsail and databases](https://www.reddit.com/r/aws/comments/hcgaxz/question_about_lightsail_and_databases/)
- url: https://www.reddit.com/r/aws/comments/hcgaxz/question_about_lightsail_and_databases/
---
Hey guys, I’m new to AWS and kind of confused about their pricing. I’m starting off with the free lightsail. My understanding is you pay for each hour you use (I assume that’s both me playing around in the backend, and anybody visiting the site) up to the 3.50 for the cheapest plan. Seems like a great deal. I kept reading and it says you have to pay 15/month for managed databases. 

This is where I’m confused. With the 3.50/month can I not create sql databases? Or are managed databases something different? Once the free trial ends, I don’t want to get with a bunch of costs I don’t quite understand. 

I’m looking to run php and python scripts, and to do so I obviously need to utilize MySQL or postgresql etc. can I create those with just the 3.50/month, or would I need the 15/month managed databases to be able to do that?

Thanks so much!
## [11][Moving a file share to FSx for Windows in a hybrid environment?](https://www.reddit.com/r/aws/comments/hcab3a/moving_a_file_share_to_fsx_for_windows_in_a/)
- url: https://www.reddit.com/r/aws/comments/hcab3a/moving_a_file_share_to_fsx_for_windows_in_a/
---
How well does this work compared to a traditional Windows File Server or on-prem NAS? How did you handle access for remote workers?
