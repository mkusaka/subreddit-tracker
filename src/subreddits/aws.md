# aws
## [1][Problems with Pearson Proctoring](https://www.reddit.com/r/aws/comments/gf0g5p/problems_with_pearson_proctoring/)
- url: https://www.reddit.com/r/aws/comments/gf0g5p/problems_with_pearson_proctoring/
---
Has anyone experienced problems with the Pearson online proctoring system? I can't seem to pass the internet test. My speedtest results look fine

https://www.speedtest.net/result/9398879276

But it still fails the internet test. Has anyone had this issue and tried to resolve it?

Update: still no luck on the internet test and my exam is in 5 mins. I even documented the experience in a video to show Pearson I was using the same machine. I hope they refund me. 

https://youtu.be/XWK_ul36B5M
## [2][When will AWS Backup support Aurora RDS PostgreSQL, if ever? In the meantime, which solution would you recommend for daily/weekly/monthly backup archival on S3?](https://www.reddit.com/r/aws/comments/geqhor/when_will_aws_backup_support_aurora_rds/)
- url: https://www.reddit.com/r/aws/comments/geqhor/when_will_aws_backup_support_aurora_rds/
---
It's disappointing for such a major AWS service to be missing proper backup support.

What solution would you recommend to perform daily/weekly/monthly backup rotation and archival on S3? Preferably without significant additional cost on top of S3.

Edit, to clarify:

We're already making use of native tools (pg_dump) to execute daily backups, which are then sent to S3, but it's not a very resilient solution currently.


I was looking for a managed way to deal with backups instead of improving our in-house scripts. The feature set in AWS Backup would be a perfect fit, alas it doesn't support Aurora.


We're already making use of snapshots. This backup to S3 is our second backup set (you can never be too careful!), which is archived to Glacier. This is in case we ever need to figure out what happened in the distant past, or restore data that a client deleted and now wants.
## [3][Route 53 forward inbound to local homelab](https://www.reddit.com/r/aws/comments/getoef/route_53_forward_inbound_to_local_homelab/)
- url: https://www.reddit.com/r/aws/comments/getoef/route_53_forward_inbound_to_local_homelab/
---
I'm just really getting into AWS after years of occasionally spinning up an EC2 instance or so. I'll be taking the exam for the Cloud Practitioner cert in a couple weeks!

So I have a local homelab that runs many different webapps. I have registered a domain with Route 53. I want some addresses like blah.blahblah.com to go to EC2 instances while others like blah2.blahblah.com to be forwarded to my homelab. What's the best way to go about this, or any resources I can read to be more familiar?
## [4][Cognito UserPool Sub or IdentityPool ID as the User ID?](https://www.reddit.com/r/aws/comments/gf660d/cognito_userpool_sub_or_identitypool_id_as_the/)
- url: /r/serverless/comments/gei5kp/cognito_userpool_sub_or_identitypool_id_as_the/
---

## [5][Connecton issues using aws cli on eu-central-1](https://www.reddit.com/r/aws/comments/gf5blf/connecton_issues_using_aws_cli_on_eucentral1/)
- url: https://www.reddit.com/r/aws/comments/gf5blf/connecton_issues_using_aws_cli_on_eucentral1/
---
Holla,

since today, I regurarily get 'connection reset by peer' errors using aws cli commands...

Anyone else?

Greetings
M
## [6][Allowing web and SSH access from hundreds of hosts?](https://www.reddit.com/r/aws/comments/geuwzd/allowing_web_and_ssh_access_from_hundreds_of_hosts/)
- url: https://www.reddit.com/r/aws/comments/geuwzd/allowing_web_and_ssh_access_from_hundreds_of_hosts/
---
My company is switching from a straightforward corporate VPN for which we have a single CIDR-friendly IP range to add to our security groups, to a janky one with hundreds of randomly assigned possible egress points.

I'm told that there's no way we can reduce the number of egress points for this new VPN. Since there are way too many possible source addresses we would need to allow in, it looks like traditional security groups are out the window, and NACLs can handle even fewer rules.

I could theoretically ask for a limit increase for rules per SG or manage multiple SGs, but we don't want to face possible performance degradation from having so many rules. We also maintain many VPCs with overlapping address space, so establishing a site-to-site tunnel to each of them would be burdensome.

If this was only for web traffic to ALBs/CF, AWS WAF would be the obvious choice, but we need something for more general traffic, mostly web with CLB/NLB and SSH. I also looked into Firewall Manager, but it has to be managed at the organization level, which I don't have.

I'm at a bit of a loss. Assuming I can't push back on our VPN operators, what other choices do I have?
## [7][Will pay someone to help me turn my python script into a lambda function](https://www.reddit.com/r/aws/comments/gepesz/will_pay_someone_to_help_me_turn_my_python_script/)
- url: https://www.reddit.com/r/aws/comments/gepesz/will_pay_someone_to_help_me_turn_my_python_script/
---
I have a python script that monitors twitter for specific keywords, and I would like it to run every minute and scan twitter. Right now I have it on a virtual machine in AWS and it's scheduled every minute but because the machine is always running it's costing me a lot monthly. I am very new to this sort of stuff, but from my research it looks like converting it to a lambda function would be better. If interested DM me
## [8][Filter pattern contain both OR and exclude?](https://www.reddit.com/r/aws/comments/gf3k40/filter_pattern_contain_both_or_and_exclude/)
- url: https://www.reddit.com/r/aws/comments/gf3k40/filter_pattern_contain_both_or_and_exclude/
---
I have filter pattern as below:

    ?Pattern1 ?pattern2 - "exclude me"

I want to match "Pattern1" OR "pattern2" that not contain "exclude me"

I tried but it didn't work, so I believe aws cloudwatch doesn't support it yet, correct?

How can I achieve the goal?
## [9][Building AWS environment documentation](https://www.reddit.com/r/aws/comments/geiuhv/building_aws_environment_documentation/)
- url: https://www.reddit.com/r/aws/comments/geiuhv/building_aws_environment_documentation/
---
Does anyone have a nice way of building some sort of documentation of a new environment? Like you walked into a new place and there was no documentation type of deal? Basically looking for a starting point vs reinventing the wheel here.
## [10][When will AWS ec2 support AMD Epyc Rome instances?](https://www.reddit.com/r/aws/comments/gf2f2h/when_will_aws_ec2_support_amd_epyc_rome_instances/)
- url: https://www.reddit.com/r/aws/comments/gf2f2h/when_will_aws_ec2_support_amd_epyc_rome_instances/
---
When will AWS ec2 support AMD Epyc Rome instances?

t3a sizes ?

Please forward this to Jeff B.
