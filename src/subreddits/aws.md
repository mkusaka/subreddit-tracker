# aws
## [1][Hosting a small JIRA instance on AWS: A case study](https://www.reddit.com/r/aws/comments/erj232/hosting_a_small_jira_instance_on_aws_a_case_study/)
- url: https://www.reddit.com/r/aws/comments/erj232/hosting_a_small_jira_instance_on_aws_a_case_study/
---
We decided to get off of our cloud version of Atlassian JIRA and host it ourselves, for a variety of reasons. We have credits to burn, and I wanted to build some recommendations on small-instance hosting since hosting recommendations are so sparse. A google search turned up a lot of "best practices", but nothing in terms of "Do X, Do Y, get up and running". 

Here's the basics:

* JIRA for a team of 6 
* Evaluation License
* 24/7 access required, but the team is all in EDT

Here's what I started with:

* Spot instance arrangement, with a fleet floor of T3.Small, with a maximum spot price set to the on-demand price of a T3.Small
* EBS at 40Gb
* RDS MySQL at M5.xlarge, with storage set at 20Gb
* SES set up for email outbounds

Key Learnings: 

* So when I spun up RDS, I had completely forgotten to change the default spinup configs, and it spun up a beefy M5.xlarge. I will have to fix this on the next go
* The instance spun up and JIRA installed fine. On configuration using the web browser, it asked for the admin credentials, then crashed. I restarted the JIRA instance and everything seem to pick up the where it left off. Logs show nothing amiss, which was weird.
* The installation supported the basics, but when I installed BigGantt, the instance died. Logs show it ran out of memory. I will have to adjust on the next go
* MySQL and JIRA: UGH. Had to install extra JDBC driver, change configs in command line, just burned an hour just getting the additional driver to work properly.

Here's what I settled on:

* Spot Instance Arrangement, with a fleet floor of T3.medium, with a maximum spot price set to on-demand price of T3.medium
* EBS at 40Gb
* RDS Postgres at T3.small, with storage set to 20Gb
* SES still active

Final takeways:

* Postgres is a great "fire and forget" solution for JIRA. As comfortable as I am with MySQL, it wasn't worth my time to fiddle with the JDBC drivers on the second go
* EC2 CPU utilization never went above 2% (??!?) according to cloudwatch, even when we had 4 concurrent users on the system
* RDS CPU Utilization never went above 5% (??!?) according to cloudwatch
* EC2 Memory usage is TIGHT, but manageable for the evaluation instance. Available memory even at max usage never dipped below 110mb, though memory utilization always seems to be close to 95-100%
* Costs in 20 days so far are:
- $9.73 for EC2 Spot Fleet
- $12.54 for RDS instnace
-  **Total after 20 days** $22.27

Is it more expensive than the cloud implementation? Sure is. But while setting this up I had a chance to learn some AWS quirks and built a baseline for the future. Would I do this again? Sure. I like pain.

*EDITED due to garbage formatting on my part**
## [2][Amazon EC2 Spot instances can now be stopped and started similar to On-Demand instances](https://www.reddit.com/r/aws/comments/erk3or/amazon_ec2_spot_instances_can_now_be_stopped_and/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/01/amazon-ec2-spot-instances-stopped-started-similar-to-on-demand-instances/
---

## [3][Deep Dive into Querying Elasticsearch. Filter vs Query. Full text search](https://www.reddit.com/r/aws/comments/erqfy0/deep_dive_into_querying_elasticsearch_filter_vs/)
- url: https://towardsdatascience.com/deep-dive-into-querying-elasticsearch-filter-vs-query-full-text-search-b861b06bd4c0
---

## [4][Advanced AWS CLI JMESPath Query Tricks](https://www.reddit.com/r/aws/comments/eruq96/advanced_aws_cli_jmespath_query_tricks/)
- url: https://opensourceconnections.com/blog/2015/07/27/advanced-aws-cli-jmespath-query/
---

## [5][ASG with Spot priority then On-Demand](https://www.reddit.com/r/aws/comments/erq8g6/asg_with_spot_priority_then_ondemand/)
- url: https://www.reddit.com/r/aws/comments/erq8g6/asg_with_spot_priority_then_ondemand/
---
I've been digging around the docs and web for a simple solution: Modify my existing AutoScaling Group to add spot instances (if available) otherwise revert to default an add the standard On-Demand instance instead.

I don't want the ASG to have a minimum base on-demand because my workload scales from 0, meaning I can go 12 hours without any load ("jobs").

I want 0% On-Demand Percentage if spot is available. But if I put 0% on-demand (100% spot), what if there is no spot availability? The ASG won't scale?

&amp;#x200B;

Perhaps my question is more related to a Launch Template. If \`Purchasing Option=Request Spot Instances\` is checked, then does that mean ONLY spot instances are purchased?

My concerns rise from the need to utilize p2.xlarge as there is no other similar instance like it and I'm worried that this limits spot availability versus everyone else who's specifying 4+ different instance types in their spot pool.
## [6][There is a Time and a Place for Anomaly Detection](https://www.reddit.com/r/aws/comments/eruavt/there_is_a_time_and_a_place_for_anomaly_detection/)
- url: https://www.sentiatechblog.com/there-is-a-time-and-a-place-for-anomaly-detection?utm_source=reddit&amp;utm_medium=social&amp;utm_campaign=cw_ad_post
---

## [7][Osaka Local Region to be upgraded to a full region in early 2021](https://www.reddit.com/r/aws/comments/ercig3/osaka_local_region_to_be_upgraded_to_a_full/)
- url: https://aws.amazon.com/blogs/aws/in-the-works-aws-osaka-local-region-expansion-to-full-region/
---

## [8][Is anyone actually using AWS Control Tower Guardrails?](https://www.reddit.com/r/aws/comments/erjpmy/is_anyone_actually_using_aws_control_tower/)
- url: https://www.reddit.com/r/aws/comments/erjpmy/is_anyone_actually_using_aws_control_tower/
---
I have just started using Control Tower since I have a greenfield account structure and I really like the idea of centrally enforcing best practice policies etc. But each Guardrail can only be enabled on a single OU at a time. 

https://docs.aws.amazon.com/controltower/latest/userguide/guardrails.html#enable-guardrails

This process takes several minutes to complete and blocks any other changes in Control Tower. Since there are 18 recommended and elective Guardrails and they each take 3+ minutes to enable you are looking at at least an hour of manual clicking, sitting and waiting PER OU to enable these. Is this just a pain you must bear since Control Tower is so new?
## [9][Cloud Economist recommendations](https://www.reddit.com/r/aws/comments/err2ym/cloud_economist_recommendations/)
- url: https://www.reddit.com/r/aws/comments/err2ym/cloud_economist_recommendations/
---
I hope this is the right place to ask this. I'm in the process of setting up a Cloud Economist tribe in my company, to share and organise knowledge about the joys of Cloud Economy. Are there any resources you can recommend that helped you understanding cost optimization, cost strategies, ...? Particular websites that regularly publish articles, specific books that I should check out. Thank you very much in advance!
## [10][Cloudwatch customize alarm using lambda keep sending empty "Logs" field](https://www.reddit.com/r/aws/comments/eroian/cloudwatch_customize_alarm_using_lambda_keep/)
- url: https://www.reddit.com/r/aws/comments/eroian/cloudwatch_customize_alarm_using_lambda_keep/
---
I'm using this repo as an example to send customized alarm using lambda [https://github.com/awslabs/cloudwatch-logs-customize-alarms](https://github.com/awslabs/cloudwatch-logs-customize-alarms)

&amp;#x200B;

This portion:

        var text = 'Alarm Name: ' + '&lt;b&gt;' + message.AlarmName + '&lt;/b&gt;&lt;br/&gt;' + 
                   'Runbook Details: &lt;a href="http://wiki.mycompany.com/prodrunbook"&gt;Production Runbook&lt;/a&gt;&lt;br/&gt;' +
                   'Account ID: ' + message.AWSAccountId + '&lt;br/&gt;'+
                   'Region: ' + message.Region + '&lt;br/&gt;'+
                   'Alarm Time: ' + date.toString() + '&lt;br/&gt;'+
                   logData;

I received emails via aws ses but it keeps sending email with empty "Logs" (logdata is empty).

Why is it happen and is there a possible solution to fix it?

Is there possible that alarm needs time to recovery and that caused empty "logData"?

&amp;#x200B;

Many thanks.
