# aws
## [1][AWS WAF and AWS Shield: The Ultimate Web App Protection Combo - IOD](https://www.reddit.com/r/aws/comments/ish6vb/aws_waf_and_aws_shield_the_ultimate_web_app/)
- url: https://iamondemand.com/blog/aws-waf-and-aws-shield-the-ultimate-web-app-protection-combo/
---

## [2][Learning Lambda - Suggest A Project?](https://www.reddit.com/r/aws/comments/isk7if/learning_lambda_suggest_a_project/)
- url: https://www.reddit.com/r/aws/comments/isk7if/learning_lambda_suggest_a_project/
---
I was tasked with learning about Lambda, but not given any actual direction on what to do with it.

What task could I tackle in a test environment that would teach me good fundamentals about it?
## [3][Why do new accounts only get access to 2 AZs in us-west-1?](https://www.reddit.com/r/aws/comments/isdvi0/why_do_new_accounts_only_get_access_to_2_azs_in/)
- url: https://www.reddit.com/r/aws/comments/isdvi0/why_do_new_accounts_only_get_access_to_2_azs_in/
---
I normally never have anything in the US regions (beyond whatever AWS has centralised) as I’m on the other side of the planet, however we are looking to expand into the US and I’m trying to decide on a region. 

For me, west coast will probably be best latency-wise so I’m trying to decide between us-west-1 and us-west-2 but I’m puzzled by Northern California having 3 zones but only 2 available to new accounts. Anyone know what that’s about?

Sorry if this has been asked before I tried searching and couldn’t find anything. 

See here, there’s an asterisks: https://aws.amazon.com/about-aws/global-infrastructure/regions_az/
## [4][Which AWS service to use for backing up office files?](https://www.reddit.com/r/aws/comments/ishmja/which_aws_service_to_use_for_backing_up_office/)
- url: https://www.reddit.com/r/aws/comments/ishmja/which_aws_service_to_use_for_backing_up_office/
---
Use case:
My fathers company wants to backup all their office data on AWS from their in-house server as they are moving to a remote first model.
Their requirements are, CRD files and have access levels for which employees can access which files and folders. 

I have zero experience with AWS but I looked into S3 and it’s very easy to use from the console for managing the files. But I couldn’t see how to set access levels for storage system.

I want avoid creating a custom console and write all the APIs for accessing the storage system.

Help would be very much appreciated.

EDIT:

Thanks a lot for the help. I’ll look into the provided solutions 🙏🏻
## [5][Slapping up a temp site on AWS while the devs work on the real one](https://www.reddit.com/r/aws/comments/is657y/slapping_up_a_temp_site_on_aws_while_the_devs/)
- url: https://www.reddit.com/r/aws/comments/is657y/slapping_up_a_temp_site_on_aws_while_the_devs/
---
Hi all, I'm working on a Code for America project called [BallotNav](https://github.com/hackforla/ballotnav), we're creating a tool to help people find in-person locations they can drop off their mail-in ballot if they are concerned about the reliability of USPS and COVID safety at the polls.

I'm trying to poop a temp site up at our URL to send press to, recruit non-tech volunteers, and let ppl sign up for launch notification emails, just quick n dirty until we can launch the real site.

I am a non-tech volunteer who can handle a wordpress site etc. and would like to be able to build and update this myself without drawing time/effort away from the real site build.  Is there any kind of simplistic CMS like that for AWS? My devs don't know, cuz they've never needed one, thought I'd toss it out to the crowd.

ETA: Never mind, turns out I can build one in wix or whatever and they can point to it, I got confused because I had to kill the redirect in Namecheap where I registered it when we did the custom nameservers.

With that said, anyone who’d like to volunteer in any capacity, come aboard!
## [6][How to setup Lake Formation to handle daily full datasets](https://www.reddit.com/r/aws/comments/isjzvc/how_to_setup_lake_formation_to_handle_daily_full/)
- url: https://www.reddit.com/r/aws/comments/isjzvc/how_to_setup_lake_formation_to_handle_daily_full/
---
Hi everyone I am trying to use lake formation to register my tables. I have a process that will put my entire dataset into an s3 bucket daily. The structure looks like this currently.

Bucket -&gt; databasename -&gt; table name1 -&gt; dated parquetfolder1
Bucket -&gt; databasename -&gt; table name1 -&gt; dated parquetfolder2
Bucket -&gt; databasename -&gt; table name1 -&gt; json file telling which parquet is the newest
Bucket -&gt; databasename -&gt; table name2-&gt; dated parquetfolder1
Bucket -&gt; databasename -&gt; table name2 -&gt; dated parquetfolder2
Bucket -&gt; databasename -&gt; table name2 -&gt; json file tellling which parquet is the newest

But currently when I setup a glue crawler it literally flattens out database name and table name and just gives me 6 tables. On for tech dated parquet and two for the json config. Since they both have the same name one gets a hash.

In reality I want two tables and I want the path to swap when the json file changes. Or if I need to save the data in a different structure to achieve snapping of my data that is also fine
## [7][Best way to move CloudTrail logs to Glacier?](https://www.reddit.com/r/aws/comments/is8k2e/best_way_to_move_cloudtrail_logs_to_glacier/)
- url: https://www.reddit.com/r/aws/comments/is8k2e/best_way_to_move_cloudtrail_logs_to_glacier/
---
I can't simply use lifecycle rule for that, because for lots of smaller files Glacier just increases costs.

The only way to move log files to Glacier without incurring too much cost is concating and pushing to Glacier. As far as I know, there are no existing solution that allows this, so I would need to create lambda function that package log files in S3 and move it to Glacier.

I want to do the same for log files created by other services, such as CloudFront logs, CloudWatch logs exprorted to S3, etc.
## [8][AWS Load Balancing and Web App redirecting to SSL](https://www.reddit.com/r/aws/comments/iskm73/aws_load_balancing_and_web_app_redirecting_to_ssl/)
- url: https://www.reddit.com/r/aws/comments/iskm73/aws_load_balancing_and_web_app_redirecting_to_ssl/
---
We have a [legacy] web application that redirects non-SSL requests (port 80) to SSL (port 443). The web server is IIS (if that matters).

Now I wish to put a load balancer in front of this app. That will be step one towards scaling in a future project phase.

However, the redirecting to SSL is causing problems that I can't quite get my head around.

I have the AWS Network Load Balancer (NLB) listening on both 80 and 443. Apparently it can only send to the target group (my web server) on **one** port.

* If I have it sending to port 80 then we go into a redirect loop because the web server assumes every request is non-SSL.
* If I have it sending to port 443 then we never redirect because, again, the web server assumes all requests are already secure

I chose NLB because it will allow me to use my existing web server's elastic IP which makes the cutover a bit more seamless (and, I'm told, some of our API clients have that IP address allow-listed already so a change is hard). I think that logic rules out using an Application Load Balancer (can't use elastic IP)? NLB is layer 4 though, so I can't play any games with headers.

I can't help but think this "redirect to SSL" thing has already been solved by smarter people than me... so how'd you do it?
## [9][Charges in AWS Educate Account](https://www.reddit.com/r/aws/comments/isjc8w/charges_in_aws_educate_account/)
- url: https://www.reddit.com/r/aws/comments/isjc8w/charges_in_aws_educate_account/
---
Hi all!

&amp;#x200B;

I have a AWS account, that came with Github Education, and I'm wondering why I'm being charged for a EC2 instance, using the free tier (**t2.micro**)

As I have a Educate account, I can't acess the billing console.

&amp;#x200B;

Thank you!
## [10][AWS work question](https://www.reddit.com/r/aws/comments/isj6fp/aws_work_question/)
- url: https://www.reddit.com/r/aws/comments/isj6fp/aws_work_question/
---
I'm using for work AWS, and I was wondering, can they see what device I connect with, like am I on my laptop or iPad? (Because it's strictly not allowed on tablets/phones etc just PC/desktop)  


Sorry if I'm not in the right place, but I didn't know where to ask this!

Thank you!
