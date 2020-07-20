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
## [2][Amplify: What is the best approach to customize sign up?](https://www.reddit.com/r/aws/comments/huivpb/amplify_what_is_the_best_approach_to_customize/)
- url: https://www.reddit.com/r/aws/comments/huivpb/amplify_what_is_the_best_approach_to_customize/
---
I'm building a web application, using Amplify's Auth module and its components to help me out interfacing with Cognito and providing authentication.  
**I need to customize Amplify's sign up process so that users are signed up (to Cognito) with an auto-generated username**, instead of them creating their own username. The users will just fill in their email and password during sign up, and then also for sign in (the username will be used only internally, hidden from users).

So I need a slight modification to the sign up process.

**What is the recommended approach to do this?**  
(taking in consideration also the recent updates to the Amplify framework)

So far I've found these possibilities, but I don't know which is right:  


1. **Extend the** `Component` **class overriding its** `signUp` **method** (in which I will call [Auth.signUp](https://aws-amplify.github.io/amplify-js/api/classes/authclass.html#signup) myself and pass it my auto-generated username together with the email and password the user filled in the form).  
Example: [https://egghead.io/lessons/react-native-manually-sign-up-new-users-in-react-with-aws-amplify-auth-class](https://egghead.io/lessons/react-native-manually-sign-up-new-users-in-react-with-aws-amplify-auth-class)   

2. **Extend the** `SignUp` **component class overriding its** `handleSubmit` **event handler** (calling Auth.signUp as described above).  
Example: [https://github.com/aws-amplify/amplify-js/issues/1436](https://github.com/aws-amplify/amplify-js/issues/1436)   

3. **Pass a custom function as a prop to** `handleSubmit` **to the sign up component** (calling Auth.signUp as described above).  
Example: [https://github.com/aws-amplify/amplify-js/issues/5480](https://github.com/aws-amplify/amplify-js/issues/5480)  

4. **Use the sign up component** `formField` **property to pass a default value to the username form field, then hide this field with CSS** (I don't even know if this is possible, and it sounds too gimmicky to me and maybe risky)  
Example: [https://docs.amplify.aws/ui/auth/authenticator/q/framework/react#custom-form-fields](https://docs.amplify.aws/ui/auth/authenticator/q/framework/react#custom-form-fields)

**I think one of these ways (or some other) may be the proper way of accomplishing it, but I don't have the code-wisdom to know it. So, if anyone can illuminate me on this it would be very helpful.**

Ps: I'm using Svelte as framework, so I'm using the Web Components version of Amplify's components (and not React, Vue, etc.).

Thanks for reading this!
## [3][Question about Launch Templates](https://www.reddit.com/r/aws/comments/huj9g8/question_about_launch_templates/)
- url: https://www.reddit.com/r/aws/comments/huj9g8/question_about_launch_templates/
---
 is it possible to have a Launch Template with a list of AMIs to spin up differently configured instances? 

As far as I’m aware, the ratio between AMIs and LT is 1:1. However in my case I want to have a list of AMIs in my launch template not just one.
## [4][Cognito CSV Import Broken?](https://www.reddit.com/r/aws/comments/huey2g/cognito_csv_import_broken/)
- url: https://www.reddit.com/r/aws/comments/huey2g/cognito_csv_import_broken/
---
I'm trying to test CSV importing for users into Cognito however am getting some strange errors in the console:

https://preview.redd.it/44m54ggwoxb51.png?width=1920&amp;format=png&amp;auto=webp&amp;s=eb568983e58c9057f970dea0d08f6f5a594a2d4e

The error message to look at is:

    Refused to connect to ...

This looks like an error in the Cognito javascript.

I'm pretty certain it's not an issue with the CSV file; in the past, incorrect CSV files can still be uploaded but when you go to process it (by clicking Start), an error occurs in the web interface (the failed jobs show this):

https://preview.redd.it/1dmx2gdzwxb51.png?width=1202&amp;format=png&amp;auto=webp&amp;s=7ec8361c5b22ea76a0c64b79dd84a64dbb95b2e0

Is anyone else able to create import jobs with a CSV file, or having a similar error?
## [5][I need help to move multiple subdomains on single EC2 instance?](https://www.reddit.com/r/aws/comments/hujd0m/i_need_help_to_move_multiple_subdomains_on_single/)
- url: https://www.reddit.com/r/aws/comments/hujd0m/i_need_help_to_move_multiple_subdomains_on_single/
---
Dear AWS community,

I need a bit of help to save my EC2 server costs.

I have 4 EC2 (+4 RDS) instances, each running its own subdomain - www.&lt;domain&gt;.com, abc.&lt;domain&gt;.com, def.&lt;domain&gt;.com, xyz.&lt;domain&gt;.com.

Currently, none of the is generating any revenue, so to save the costs, I wish to merge them all to single EC2 instance.

Each of the webcode is setup to deploy through bitbucket code based CI/CD (continuous development/integration)

Any pointers are really helpful.

Thanks!
## [6][Question about free tier for AWS Lambda](https://www.reddit.com/r/aws/comments/hughfl/question_about_free_tier_for_aws_lambda/)
- url: https://www.reddit.com/r/aws/comments/hughfl/question_about_free_tier_for_aws_lambda/
---
I want to start learning AWS Lambda and deploy a few functions. The services I believe I will use are 

* Lambda
* Gateway (to convert to my lambda functions to REST APIs)
* DynamoDB

I know that AWS has a 12 month free trial in which all should be covered. What I am not sure is what will happen after this trial period is over.

**Lambda**: Is it true that I get 1 million requests per month even after the free tier?

**Gateway:** After the 12 months, do you have to pay no matter what? I think this is true after reading the pricing page: [https://aws.amazon.com/api-gateway/pricing/](https://aws.amazon.com/api-gateway/pricing/)

**DynamoDB**: I was a little confused about their free tier here: [https://aws.amazon.com/dynamodb/pricing/](https://aws.amazon.com/dynamodb/pricing/)

Base on [this page](https://aws.amazon.com/free/?sc_icontent=awssm-evergreen-free_tier&amp;sc_iplace=2up&amp;trk=ha_awssm-evergreen-free_tier&amp;sc_ichannel=ha&amp;sc_icampaign=evergreen-free_tier&amp;all-free-tier.sort-by=item.additionalFields.SortRank&amp;all-free-tier.sort-order=asc&amp;awsf.Free%20Tier%20Types=tier%23always-free&amp;awsm.page-all-free-tier=1), I see that I'm getting 1 million requests per month on **Lambda** and 25 GB of **DynamoDB** storage. Just to confirm, will I continue getting these services after the 12 month free tier? Does this mean that I will only have to pay for AWS **Gateway**?

Sorry for the basic questions. I just wanted to confirm before creating an account, and didn't want to risk getting a high bill
## [7][What should be the IAM policy to provide S3 access to a federated user of a different AWS account?](https://www.reddit.com/r/aws/comments/huh86g/what_should_be_the_iam_policy_to_provide_s3/)
- url: https://www.reddit.com/r/aws/comments/huh86g/what_should_be_the_iam_policy_to_provide_s3/
---
I do know to provide IAM policies to normal IAM users of other accounts but how do I provide access to federated login users? Please help me here. Thank you!
## [8][Looking for a website I remember being posted here](https://www.reddit.com/r/aws/comments/huh7nm/looking_for_a_website_i_remember_being_posted_here/)
- url: https://www.reddit.com/r/aws/comments/huh7nm/looking_for_a_website_i_remember_being_posted_here/
---
Hey, sorry if this is a dumb post.   
  
I remember seeing a comment somewhere here regarding a website of an AWS expert who will send your CV to his mail list if you complete an AWS hosted website with a load of steps including good source control management and CI/CD.  
  
Does anyone know the link?  
  
Looking for a mate who is trying to build up a GitHub portfolio for his CV.
## [9][AWS Static Website Cost Risks?](https://www.reddit.com/r/aws/comments/hu857x/aws_static_website_cost_risks/)
- url: https://www.reddit.com/r/aws/comments/hu857x/aws_static_website_cost_risks/
---
Are there any cost associated risks to hosting a static personal blog on AWS with resources stored in S3?

For example if there was a DDOS attack would I be 100% responsible for footing the bill if it somehow wasn’t or only partially blocked by AWS shield? (I’m sure I could go a million years before getting some kind of attack on a blog, but I still want to know if there are risks no matter how insignificant)

I can can create a website with HTML, bootstrap, css, and JavaScript, but I’m definitely not a developer or architect. I’m just looking for cheap web hosting options that give me complete control of the website.
## [10][odd question is it possible to separate a single beanstalk api that does login,registration,verification to their own route 53 record set without the other routes being accessible and without splitting beanstalk project?](https://www.reddit.com/r/aws/comments/hukri6/odd_question_is_it_possible_to_separate_a_single/)
- url: https://www.reddit.com/r/aws/comments/hukri6/odd_question_is_it_possible_to_separate_a_single/
---
or is it better to split them up to their own beanstalk
like Dashboard , Verification ,Registration.
## [11][Attempting to upload files to AWS - Error - "No etag was provided by S3."](https://www.reddit.com/r/aws/comments/huef1s/attempting_to_upload_files_to_aws_error_no_etag/)
- url: https://www.reddit.com/r/aws/comments/huef1s/attempting_to_upload_files_to_aws_error_no_etag/
---
First time I've received this error, have been using AWS for years...  I've repeatedly got this error trying to upload images to my AWS today, and can't seem to find anything online regarding this.

Any chance someone around here knows what could be causing this?
