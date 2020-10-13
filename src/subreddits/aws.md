# aws
## [1][Week of Oct 12th - What are you building this week in AWS?](https://www.reddit.com/r/aws/comments/j9s5xz/week_of_oct_12th_what_are_you_building_this_week/)
- url: https://www.reddit.com/r/aws/comments/j9s5xz/week_of_oct_12th_what_are_you_building_this_week/
---
Share what you are working on
## [2][I built a GitHub Action that deploys static sites to Cloudfront](https://www.reddit.com/r/aws/comments/ja1k9g/i_built_a_github_action_that_deploys_static_sites/)
- url: https://github.com/onramper/action-deploy-aws-static-site
---

## [3][Timestream vs. InfluxDB](https://www.reddit.com/r/aws/comments/ja8hg3/timestream_vs_influxdb/)
- url: https://www.reddit.com/r/aws/comments/ja8hg3/timestream_vs_influxdb/
---
We have been using influxdb in the past as a time series database, however, I would be curious as to how it compares to Timestream? Apart from the obvious advantages of Timestream being a fully managed database, is there anything Timestream does that influxdb can’t do?
Thanks!
## [4][Set alarm on sending rate SES](https://www.reddit.com/r/aws/comments/jadisj/set_alarm_on_sending_rate_ses/)
- url: https://www.reddit.com/r/aws/comments/jadisj/set_alarm_on_sending_rate_ses/
---
Unfortunately one of our systems went into a loop and started sending replies to a out of office reply (internally luckily), but resulting in over 100k emails sent in the last 24 hours.

https://imgur.com/a/6sF0gTL

As a result that we've hit our limit for 24 hours, I requested a temporary limit increase to circumvent this, but I don't want to run into this issue again.

I tried setting up a cloudwatch alarm for my SES metrics, but those metrics don't correlate to what is in my SES Sending Statistics screen:
https://imgur.com/a/3vmJ4iH

What parameters do I need to put in here to get the same graphs?
## [5][EC2 VPN Data Transfer price question](https://www.reddit.com/r/aws/comments/jadi6c/ec2_vpn_data_transfer_price_question/)
- url: https://www.reddit.com/r/aws/comments/jadi6c/ec2_vpn_data_transfer_price_question/
---
So I've setup a VPN on my EC2 server. How will Amazon calculate my data transfer usage? Say I downloaded 1GB of data via VPN, will they count it as 1GB IN (From internet to EC2) + 1GB OUT (From EC2 to my device)?
## [6][Connect EC2 instances without IP (using DNS?)](https://www.reddit.com/r/aws/comments/ja9zsb/connect_ec2_instances_without_ip_using_dns/)
- url: https://www.reddit.com/r/aws/comments/ja9zsb/connect_ec2_instances_without_ip_using_dns/
---
I have now searched for a couple of hours for this without a result... this shows me that my approach is maybe flawed or I am using the wrong search terms...

I want multiple EC2 instances to work together. They are in a private subnet. They don't need to be reachable on the internet. 

Of course they can reach one another using their IPs. However, if I rebuild an instance it will get a new IP. Therefore I want them to have internal DNS names within their VPC so they can reach each other using "server1.local" or "server2.local".

Is that possible using AWS tools?
## [7][Ever Wonder About Snowflake and Databricks' Relationship with AWS as Allies &amp; Competitors?](https://www.reddit.com/r/aws/comments/jacvi8/ever_wonder_about_snowflake_and_databricks/)
- url: https://www.reddit.com/r/aws/comments/jacvi8/ever_wonder_about_snowflake_and_databricks/
---
Last month the Future Data Conference had Ben Horowitz (partner at Andreessan-Horowitz) give his thoughts,[I thought it was interesting](https://medium.com/whispering-data/the-3-most-interesting-ideas-from-the-future-data-conference-caf849fa4958?source=friends_link&amp;sk=b6fbdd7e4bc35f9d1298c981a7f12820)!
## [8][How can I make a real time dashboard (with a time lag of less than a minute)?](https://www.reddit.com/r/aws/comments/ja9u5n/how_can_i_make_a_real_time_dashboard_with_a_time/)
- url: https://www.reddit.com/r/aws/comments/ja9u5n/how_can_i_make_a_real_time_dashboard_with_a_time/
---
  

Question- How can I make a real time dashboard (with a time lag of less than a minute)? 

My database is available in AWS Postgres RDS instance. I also made dashboards in Quicksight but it can refresh data hourly only and now we need real time dashboarding.
## [9][Windows CIFS replacement options](https://www.reddit.com/r/aws/comments/jab71r/windows_cifs_replacement_options/)
- url: https://www.reddit.com/r/aws/comments/jab71r/windows_cifs_replacement_options/
---
Our team utilizes several applications and systems that are being migrated to AWS. We also use an extensive amount of excel spreadsheets for Third-Party compliance report tracking. We have been told by our internal security group that once the transition of the systems servers and application servers are completed (December 2020) they will not open ports to our Windows file share

I am researching, frantically I might add, for a solution in AWS that fits the requirements.

* Ability to write (Alteryx) excel file outputs programmatically 
* Teams to be able to access files in a file directory type of interface (GUI) team is best described as technically autistic (They are highly skilled in working our specific systems, but new technology introduction is a painstaking, loooooong training process.)
* Read, write files as needed. Techs pull generated docs (Alteryx) from file system, research critical issues, add notes/ticket creations/etc and save back to file store.

&amp;#x200B;

It seems a simple list. I thought a EC2 with EFS mounted was our answer. Our data needs is small, under 200GB with the vast majority being archival (120GB historical archive) access for a small team of techs (30), on Windows machines.  


But the rub there was Windows, since EFS has no native windows support. 

I just found  Amazon FSx but haven't researched it enough.  


Any suggestions?
## [10][API Gateway &lt;-&gt; EC2 connection](https://www.reddit.com/r/aws/comments/jaaqr8/api_gateway_ec2_connection/)
- url: https://www.reddit.com/r/aws/comments/jaaqr8/api_gateway_ec2_connection/
---
Hi guys,

I'm trying to set up an EC2 with an application that I want to be consumed only by API Gateway and hides the endpoints from everything else.

I tried with a private VPC but then API Gateway doesn't find the endpoints. 

&amp;#x200B;

I'm really lost now, if API Gateway had a static Ip I could make a security group that blocked everything else and use a pubic VPC but doens't seem to be the case.

&amp;#x200B;

Any ideas I could try? 

&amp;#x200B;

Thanks!
## [11][We now have AWS TLS termination at the network load balancer. What’s the difference with ALB and NLB? Both can perform the same actions now?](https://www.reddit.com/r/aws/comments/ja822t/we_now_have_aws_tls_termination_at_the_network/)
- url: https://www.reddit.com/r/aws/comments/ja822t/we_now_have_aws_tls_termination_at_the_network/
---

