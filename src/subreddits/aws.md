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
## [2][ECS - our server response time has dropped from 0.3s to 2.5s - part 2](https://www.reddit.com/r/aws/comments/htxrzr/ecs_our_server_response_time_has_dropped_from_03s/)
- url: https://www.reddit.com/r/aws/comments/htxrzr/ecs_our_server_response_time_has_dropped_from_03s/
---
Hi everyone, wanted to thank you all for your contributions, your response was fantastic and so helpful. I resolved my CPU cloudwatch issue, which was due to a very low default cpu setting (thanks [rehevkor5](https://www.reddit.com/user/rehevkor5/) &amp; [jIsraelTurner](https://www.reddit.com/user/jIsraelTurner/)).

I have also ruled out a number of things in my first post which are not causing the 2.2s discrepancy. [Previous post here.](https://www.reddit.com/r/aws/comments/htgbnj/ecs_our_server_response_time_has_dropped_from_03s/)

1. It isn't related to the php version, apache version or the code as far as I can tell.
2. It isn't related to the RDS.
3. EFS isn't causing this issue.

I ruled these all out by setting up an identical site without a certificate. This site has a TTFB of 0.1s.

I'm now assuming my problem is related to my load balancer or is something to do with the certificate or Route53.

My ALB has two listeners:

`HTTP:80 - redirecting to HTTPS://#{host}:443/#{path}?#{query}HTTPS:443 - forwarding to http-target-group w/ ssl certificate`

I direct the domain to the ALB using an Alias record in Route53. I use google lighthouse to get the TTFB value. The http-target-group directs to a randomly assigned port on the EC2 target, which is created by ECS.

I use this meta tag `&lt;meta http-equiv="Content-Security-Policy" content="upgrade-insecure-requests"&gt;` as the server assumes it is running on HTTP because traffic enters on port 80. This ensures the browser loads everything over HTTPS.

On the "fast" version, I just have `HTTP: 80 forwarded to http-target-group` and it works fine.

&amp;#x200B;

Does anyone have any ideas? I'd also welcome advice on configuring the load balancer.
## [3][Looking for advice on how best to store data in a serverless project: DynamoDB, Athena/S3 or Aurora Postgres Serverless](https://www.reddit.com/r/aws/comments/htpjzi/looking_for_advice_on_how_best_to_store_data_in_a/)
- url: https://www.reddit.com/r/aws/comments/htpjzi/looking_for_advice_on_how_best_to_store_data_in_a/
---
I have a question about a project I'm doing to learn more about serverless technologies on AWS. Here are the goals of the project:

* Collect and consolidate Quarterly 13F filing data from the SEC
* Generate a single CSV file containing all holdings from every 13F filing since 2013
* Use only serverless technologies on AWS
* Define and deploy all of the project infrastructure using AWS CDK
* Learn more about the SEC, 13F and do analytics on the collected data
* Share my experience and get feedback on serverless architecture patterns and best practices

Here's my approach so far:

[AWS Serverless architecture for consolidating 13F filing data from the SEC](https://preview.redd.it/z4uoyr7hyob51.png?width=761&amp;format=png&amp;auto=webp&amp;s=bb2fcf2fea3a9742cfe09ed8de08441792162a11)

1. The SEC makes available a file each quarter that contains a list of all SEC filings, including 13F filings.

* The file is called the `Master Index of EDGAR Dissemination Feed`
* The file contains the following fields: `CIK|Company Name|Form Type|Date Filed|Filename`
* There are about 200,000+ filings per quarter, \~6,000 of which are 13F filings
* This file will be uploaded to an S3 bucket (`feed_bucket`) that triggers a Lambda function (`feed_lambda`)

1. `feed_lambda` parses the Quarterly file and sends messages to an SQS Queue

* The messages sent to SQS contain the information from each row that is a 13F filing: `CIK` (an company index), `Company Name`, `Form Type` (there are some variants to 13F filings), `Date Filed` and `Filename`
* `Filename` is the location of the 13F filing on `https://www.sec.gov/Archives/`, for example `https://www.sec.gov/Archives/edgar/data/1000097/0001000097-19-000001.txt`.
* SQS messages are processed by by a lambda function called `save_filing`

1. `save_filing` fetches 13F filings from `sec.gov` and saves these raw files to S3

* The `save_filing` Lambda function process up to 10 SQS messages at a time.
* `save_filing` saves raw `.txt` file filings to a `filings` S3 bucket that triggers a `process_filing` Lambda function.

1. `process_filing` uses regular expressions and Python's native XML parsing library to read the filing data

* Each filing contains a list of holdings that are read into a `namedtuple`. Additional processing, cleaning and validation are done for each holding listed in the given 13F filing.
* To save the cleaned holdings, I'm unsure of which option would the best fit. Some options I have considered are listed below:

## Options for storing 13F data

1. Store each holding as an Item in DynamoDB

* I think that getting data in to DynamoDB would be easy, but there's no easy way to export the full set of items from **DynamoDB** (possibly over 100 million items)
* I haven't used DynamoDB before and I'm interested in trying it out.

1. Save holding data to CSV files stored in an S3 bucket and export the data using a `select *` query on **AWS Athena**

* This could be a good option, but I would have lots of small files which might not be optimal for use with Athena
* The speed of the query to export data is not super important, but the size of the dataset is nowhere close to the petabyte scale that Athena can handle

1. Store data in **Aurora Postgres Serverless**

* I'm familiar with Postgres, but I'm not sure if I need it
* This would also involve a VPC which would make things slightly more complicated, I think

As I mentioned above, my goal is to export large CSV with all 13F holding data that I can use in a Business Intelligence tool such as Google Data Studio, Tableau or Qlik.

Does anyone have suggestions on how best to store data for this project?
## [4][AWS Glue Job Runs forever with a joined PySpark dataframe but not the other dataframes](https://www.reddit.com/r/aws/comments/htodbb/aws_glue_job_runs_forever_with_a_joined_pyspark/)
- url: https://www.reddit.com/r/aws/comments/htodbb/aws_glue_job_runs_forever_with_a_joined_pyspark/
---
As you can see below, at line 48, I am converting the dataframe df3 into a dynamic frame. There is code out of view that then writes the dynamic frame to a specific location.

The issue is using df3, which is a dataframe created via the joining of 2 other dataframes, causes the Glue Job to run forever. I know if I change line 48 to use the other dataframes, df2, df1, or df, the glue job runs successfully and stops.

I also know that df3 is created successfully, as I tested it in a jupyter notebook. I'm not sure why the glue job would run forever. Any ideas?

https://preview.redd.it/4z97v45xlob51.png?width=1093&amp;format=png&amp;auto=webp&amp;s=497e0c96b27c576f61ee02b449cbe9f5f55cdf42
## [5][If you were tasked to build a consumer Internet app (say Instagram), and have confidence that it grow big, would you build it entirely with serverless (api-gateway + lambda) or backend or server oriented (ECS)?](https://www.reddit.com/r/aws/comments/htrfco/if_you_were_tasked_to_build_a_consumer_internet/)
- url: https://www.reddit.com/r/aws/comments/htrfco/if_you_were_tasked_to_build_a_consumer_internet/
---
A couple of friends had an interesting thought experiment last week, roughly as stated in the title.

“and have confidence that it grow big” — this point seemed to lead to to go ECS route 
  - because of long term infrastructure cost seems to favor server model
  - also because of well understood scalability model

I felt it seemed reasonable, but I am not fully convinced because those discussions were fairly high level , not data driven.

Anyone had similar analysis?
## [6][EKS Fargate Spot capacity provider ever coming?](https://www.reddit.com/r/aws/comments/hty34m/eks_fargate_spot_capacity_provider_ever_coming/)
- url: https://www.reddit.com/r/aws/comments/hty34m/eks_fargate_spot_capacity_provider_ever_coming/
---
From the looks of it, the feature request is more dead than alive  
[https://github.com/aws/containers-roadmap/issues/622](https://github.com/aws/containers-roadmap/issues/622) 

Does anyone have any inside scoop?

IMHO, this is will be the ultimate kubernetes cloud offering once realized. What do you think?
## [7][ECS - our server response time has dropped from 0.3s to 2.5s](https://www.reddit.com/r/aws/comments/htgbnj/ecs_our_server_response_time_has_dropped_from_03s/)
- url: https://www.reddit.com/r/aws/comments/htgbnj/ecs_our_server_response_time_has_dropped_from_03s/
---
I've been updating a legacy PHP app (no version control for 10 years) and I've gotten it working pretty nicely on AWS now. I have some problems I can't really fix.

1. CPU usage for the ECS service is always above 130%. I don't understand why as the CPU for the EC2 box is only 8%, docker process says the same. This isn't an intensive site, it's just some really old PHP code.
2. We have a response time of 2.5s instead of 0.3s. In Google lighthouse this is indicated by \`Reduce server response times (TTFB)**\`.** The apache server setup is the same, and the code running the site is the same. Only difference is my code runs on ECS instances, and the old code runs directly on an IP exposed EC2 box.

Our setup is roughly this:

Application Load Balancer

2 target groups, HTTPS and HTTP.

HTTP does a 301 redirect to out HTTPS group. (I set this up as the site kept defaulting to HTTP - is this normal?)

At the moment we have 1 cluster, 1 service and 1 task running on ECS using EC2.

Our EC2 box is dedicated, t2 medium.

Our files are on EFS. Here we store all of our cache files, image files and session files so they are shared.

We have a certificate issued by Route53 and the site validates fine.

Docker is running Apache *20051115*, the site is on PHP5.4 and the database is MySQL 5.5.

Does anyone have any idea what could be happening? Thanks!
## [8][Essential knowledge javascript and typescript for AWS CDK](https://www.reddit.com/r/aws/comments/htk5ad/essential_knowledge_javascript_and_typescript_for/)
- url: https://www.reddit.com/r/aws/comments/htk5ad/essential_knowledge_javascript_and_typescript_for/
---
In our company, we've started the migration process to AWS CDK. Does someone know which level of typescript and javascript I should have to work properly with this IaaC? I'm no novice in programming and have three years of experience with python. Could you recommend me a list of books/course/videos for studying js and ts? P.S. I know that AWS CDK is also available in python, but we decided to use typescript because it's a native language for it.
## [9][Running my First API on AWS](https://www.reddit.com/r/aws/comments/htsihp/running_my_first_api_on_aws/)
- url: https://www.reddit.com/r/aws/comments/htsihp/running_my_first_api_on_aws/
---
I am planning on running my first API to connect Salesforce Data to my website.

&amp;#x200B;

I am using the Flask framework and am struggling with understanding:

&amp;#x200B;

1. Which Amazon Product Should I be Using - Lambda Functions?
2. What the price would be for me to host this for an application that would receive under 1 million requests.

&amp;#x200B;

I have never hosted anything on AWS and this feels a bit daunting. Reaching out to the community for some guidance!
## [10][Couchbase as a service](https://www.reddit.com/r/aws/comments/htp2a6/couchbase_as_a_service/)
- url: https://www.reddit.com/r/aws/comments/htp2a6/couchbase_as_a_service/
---
Short question as in topic, is there any couchbase as a service thing available on AWS? Thanks for feedback
## [11][Different NICE-DCV instances in unique EC2 instance](https://www.reddit.com/r/aws/comments/htpgdt/different_nicedcv_instances_in_unique_ec2_instance/)
- url: https://www.reddit.com/r/aws/comments/htpgdt/different_nicedcv_instances_in_unique_ec2_instance/
---
My friend and I would like to access the same Windows EC2 instance with different sessions.

We've read the documentation, yet we don't understand the notion of session with this tool. Does 1 NICE-DCV Server = 1 Session? Or is it something else?

I'd appreciate a clarification on this topic.

Thanks in advance!
