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
## [2][AWS Cognito SignInWithApple throws apple_sub Attribute does not exist in the schema error](https://www.reddit.com/r/aws/comments/gq9nvw/aws_cognito_signinwithapple_throws_apple_sub/)
- url: https://www.reddit.com/r/aws/comments/gq9nvw/aws_cognito_signinwithapple_throws_apple_sub/
---
copy paste from:  [https://stackoverflow.com/questions/62001500/aws-cognito-signinwithapple-throws-apple-sub-attribute-does-not-exist-in-the-sch](https://stackoverflow.com/questions/62001500/aws-cognito-signinwithapple-throws-apple-sub-attribute-does-not-exist-in-the-sch)

I am currently developing SignInWithApple federation using amazon cognito. During the user cognito apple signup I receive a callback with apple\_sub Attribute does not exist in the schema error. Full callback url below:

[`https://www.localhost:3000/cognito/oauth/callback?error_description=Invalid+user+attributes%3A+apple_sub%3A+Attribute+does+not+exist+in+the+schema.%0A+&amp;state=provider%3DSignInWithApple&amp;error=invalid_request`](https://www.localhost:3000/cognito/oauth/callback?error_description=Invalid+user+attributes%3A+apple_sub%3A+Attribute+does+not+exist+in+the+schema.%0A+&amp;state=provider%3DSignInWithApple&amp;error=invalid_request)

I have tried connecting my cognito to a PreSignUp\_ExternalProvider lambda trigger and deleting the apple sub attribute via the code below:

`// appleSignIn has an obsolete apple_sub userAttributes`

`// that is not present inside AWS infrastructure and causes crashes`

`if (eventData?.request?.userAttributes?.apple_sub) {`

`delete eventData.request.userAttributes.apple_sub`

`}`

This worked marvelously until today. Today, For some reason, even when I delete the apple\_sub user attribute before passing the event further, I am still receiving the error above.

I have also tried adding a custom apple\_sub attribute, but that didn't help me also. There is also no apple\_sub standard attribute to be found.

Is there a way how to add standardized apple\_sub attribute to my cognito schema? Why is it not there by default when I allow SignInWithApple federation?
## [3][Best practices for custom authentication for calling api from lambda](https://www.reddit.com/r/aws/comments/gq5c7k/best_practices_for_custom_authentication_for/)
- url: https://www.reddit.com/r/aws/comments/gq5c7k/best_practices_for_custom_authentication_for/
---
Any tips or tricks anyone has for calling an API built in api gateway from a lambda? Additional information, building an app that uses Userpool jwt and a custom auth lambda. Currently checking the validity of the token and the request itself, how could I work a lambda into this flow?
## [4][Is Cloudfront effective when visitors aren't near any data centers?](https://www.reddit.com/r/aws/comments/gq39k5/is_cloudfront_effective_when_visitors_arent_near/)
- url: https://www.reddit.com/r/aws/comments/gq39k5/is_cloudfront_effective_when_visitors_arent_near/
---
If I serve a website that has images, how effective is CloudFront if my users are all in Hawaii?

My servers are in west-2, as are my S3 buckets. I guess I'm wondering if there's an advantage to serving images through a CDN, given my users are always in the same location.

For this scenario, disregard users on VPNs, users travelling, etc. Assume all users are always in Hawaii.

Thanks
## [5][How to copy/move an EC2 instance/volume/snapshot to another region/AZ?](https://www.reddit.com/r/aws/comments/gqabtl/how_to_copymove_an_ec2_instancevolumesnapshot_to/)
- url: https://www.reddit.com/r/aws/comments/gqabtl/how_to_copymove_an_ec2_instancevolumesnapshot_to/
---
I'm studying for my SAA-C02 and out of all the things, I can't figure out what the best practice is for moving EC2 instances between AZs and regions. Would you make a snapshot of the volume, somehow transfer the snapshot to another place (how?) and then make a new boot volume from that snapshot (how?) ? How do you back up snapshots to S3?

Also, I tried creating a new instance from an existing snapshot of mine and I couldn't figure out how. When launching a new instance, I can only create a new root volume, then when the instance is up and running I can go from "Volumes"-&gt;"Attach Volume to instance" which is not what I wanted.

Thank you for your help!
## [6][College Student Brand New to AWS - Personal Project for the Summer?](https://www.reddit.com/r/aws/comments/gq0khx/college_student_brand_new_to_aws_personal_project/)
- url: https://www.reddit.com/r/aws/comments/gq0khx/college_student_brand_new_to_aws_personal_project/
---
Hello!

Just looking for some advice. I'm studying Information Systems at my University but recently became interested in the business/strategic side of Cloud Computing (think Cloud/Infrastructure Advisory roles such as [this one](https://www.accenture.com/us-en/careers/jobdetails?id=00752687_en)).

I know to have a strong high-level understanding of AWS, it's important to get some hands-on experience so I can hold my own in future interviews/jobs. This summer, after getting my AWS CCP, I want to do a simple personal project in AWS. There's a student/trial version of Azure I can use, as I'm not looking to break the bank. 

Any recommendations on a simple project or two to do in AWS this summer? 

I really need some insight, as I have very limited knowledge of what's possible after the AWS CCP. Please help. Thanks!

UPDATE: Wow this really blew up! Thanks so much, I have so many ideas to choose from now, and I can’t wait to get started!
## [7][Are there any sites where people share common cloud formation, or other templates?](https://www.reddit.com/r/aws/comments/gpwuy7/are_there_any_sites_where_people_share_common/)
- url: https://www.reddit.com/r/aws/comments/gpwuy7/are_there_any_sites_where_people_share_common/
---
Are there any sites that are collecting re-usable templates? I don't really care whether it's terraform, pulumi, amazon cdk, or cloud formation.

Most large orgs have common sets of templates people can modify to checkout resources. Sometimes it interfaces with service catalog sometimes it does not, but in general some kind of streamlined self checkout for common requests. Is there a database of open sourced resources like this?
## [8][Triggering an event from an update to a table in our MySQL RDS.](https://www.reddit.com/r/aws/comments/gq7k0d/triggering_an_event_from_an_update_to_a_table_in/)
- url: https://www.reddit.com/r/aws/comments/gq7k0d/triggering_an_event_from_an_update_to_a_table_in/
---
We have a table in our MySQL RDS that gets updated whenever a customer scans a certain code. 

However, this update does not reflect immediately on our CRM's record of that customer (The CRM service we use is completely independent of any of AWS' services). Right now, we are therefore having to update the CRM manually on a daily basis to reflect the updates from the RDS table.

Essentially, what I'm proposing is being able to automate the following process:


1. RDS table updated 
2. Lambda Function triggered, the customer ID is retrieved from the row that was updated
3. CRM APIs are used by the lambda function to also update the record in the CRM, based on the customer ID that both the CRM and RDS share in common. 

My problem however seems to be that we can't trigger an event from the RDS. I'm curious as to what the best workaround for this is. 

I'm also wondering if this is considered best-practice for this use-case, or if there's a more elegant method of achieving the same thing?
## [9][I wrote a complete guide to the AWS Systems Manager Parameter Store. It's a stellar way to make your applications more scalable and less failure prone](https://www.reddit.com/r/aws/comments/gpnqin/i_wrote_a_complete_guide_to_the_aws_systems/)
- url: https://seanjziegler.com/a-complete-guide-to-using-the-aws-systems-manager-parameter-store/
---

## [10][What is the best way to make a SSML voice more realistic ?](https://www.reddit.com/r/aws/comments/gq79w1/what_is_the_best_way_to_make_a_ssml_voice_more/)
- url: https://www.reddit.com/r/aws/comments/gq79w1/what_is_the_best_way_to_make_a_ssml_voice_more/
---
I'm working on a project to make shows and songs based on amazon polly voices and i would like to know what kind of editing i should make to achieve realism on the voice.

Here is one example of the realism i want to achieve : [https://youtu.be/PyjmfrFUZ\_4](https://youtu.be/PyjmfrFUZ_4)
## [11][Struggling with Route53 and changing an instance's associated domain.](https://www.reddit.com/r/aws/comments/gq79bl/struggling_with_route53_and_changing_an_instances/)
- url: https://www.reddit.com/r/aws/comments/gq79bl/struggling_with_route53_and_changing_an_instances/
---
Hi all,

I’m hoping someone can help me with what I’m sure isn’t too difficult, but I’m having a hard time figuring out.


The situation:
We were using AWS for a Woocommerce site, but recently we’ve moved our .com domain away from AWS. The instance (EC2, RDS, etc.) has been untouched, I simply pointed our .com domain elsewhere via our registrar.

The desired outcome:
We already have a .net domain associated with AWS, and would look this to now point to the old .com instance as a legacy reference. 

The issue:
This is my first time working with Route53, so I’m unsure which DNS entries I need to copy over from the previous .com zone to the .net zone. Is Route53 all that is needed here or are there changes that need to be made to the instance? I have already changed the domain in Wordpress wp-config on the instance.


Many thanks in advance for any light you can shine on this for me! I’m relatively well-versed in CPanel but was kind of dropped into AWS unexpectedly.
