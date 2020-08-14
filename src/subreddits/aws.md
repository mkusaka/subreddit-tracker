# aws
## [1][c5ad instances inbound!](https://www.reddit.com/r/aws/comments/i9l3n1/c5ad_instances_inbound/)
- url: https://www.reddit.com/r/aws/comments/i9l3n1/c5ad_instances_inbound/
---
The c5ad instances are in the [price list](https://aws.amazon.com/ec2/pricing/on-demand/), and I was able to "change instance type" to one in the Oregon region. However, when I tried to start the instance, it got stuck in the Pending state. (I got tired of waiting and stopped it.)

I expect they will fix this, then (if they follow their usual m.o.) there will be an announcement... stay tuned!

**Edit:** I'm running one now!
## [2][Amazon Braket – Go Hands-On with Quantum Computing](https://www.reddit.com/r/aws/comments/i97yi4/amazon_braket_go_handson_with_quantum_computing/)
- url: https://aws.amazon.com/blogs/aws/amazon-braket-go-hands-on-with-quantum-computing/
---

## [3][AWS Step Functions adds updates to ‘choice’ state, global access to context object, dynamic timeouts, result selection, and intrinsic functions to Amazon States Languages](https://www.reddit.com/r/aws/comments/i96png/aws_step_functions_adds_updates_to_choice_state/)
- url: https://aws.amazon.com/blogs/aws/aws-step-functions-adds-updates-to-choice-state-global-access-to-context-object-dynamic-timeouts-result-selection-and-intrinsic-functions-to-amazon-states-languages/
---

## [4][Hack to support front-line developers - 3rd Annual AWS Hackathon for Good](https://www.reddit.com/r/aws/comments/i9loxh/hack_to_support_frontline_developers_3rd_annual/)
- url: https://awshackforgood.devpost.com
---

## [5][[S3] Cross-Region replication times increased](https://www.reddit.com/r/aws/comments/i9itjk/s3_crossregion_replication_times_increased/)
- url: https://www.reddit.com/r/aws/comments/i9itjk/s3_crossregion_replication_times_increased/
---
Hey! I noticed some time ago that the cross-region replication (Ireland - SP) increased greatly comparatively to last year.  I had an experiment running that uses S3 and I noticed that last year, replication time was around 1 second, then it the beginning of this summer it changed to a bimodal distribution (some around 1 second, some around 25 seconds) and currently it takes over 20 seconds.  
Has anyone else noticed this effect?  
Thanks in advance!
## [6][Hands-on exercise to use Lambda Function with Systems Manager](https://www.reddit.com/r/aws/comments/i9k04z/handson_exercise_to_use_lambda_function_with/)
- url: http://aws-dojo.com/excercises/excercise13
---

## [7][Lock down access to AWS web gui](https://www.reddit.com/r/aws/comments/i9i3xo/lock_down_access_to_aws_web_gui/)
- url: https://www.reddit.com/r/aws/comments/i9i3xo/lock_down_access_to_aws_web_gui/
---
I'm on a restricted network where only certain websites or ip addresses can be whitelisted. However, I need to use web interface for AWS. But it needs access to randomname.cloudfront.com (and others) domains to work and I think these change regularly. 

Whitelisting all of cloudfront is a big security issue.
Any ideas to lock this down? Is there maybe some website which lists all the IP addresses/domains from amazon which are required for working with the web interface ?
## [8][A piece on how to use AWS Landing Zone to its best abilities](https://www.reddit.com/r/aws/comments/i9lynv/a_piece_on_how_to_use_aws_landing_zone_to_its/)
- url: https://caylent.com/optimize-aws-landing-zone
---

## [9][[LONG-FORM VIDEO]: Complete AWS DeepRacer -&gt; Evo hardware upgrade.](https://www.reddit.com/r/aws/comments/i9is5d/longform_video_complete_aws_deepracer_evo/)
- url: https://www.reddit.com/r/aws/comments/i9is5d/longform_video_complete_aws_deepracer_evo/
---
&amp;#x200B;

https://reddit.com/link/i9is5d/video/4crpn2j0rxg51/player
## [10][AWS VPN Client from CLI](https://www.reddit.com/r/aws/comments/i9imtq/aws_vpn_client_from_cli/)
- url: https://www.reddit.com/r/aws/comments/i9imtq/aws_vpn_client_from_cli/
---
Hi guys, I was wondering if there’s a way to launch the AWS VPN client from CLI and have it automatically connect to the VPC without having to launch the exe and enter user name and password and then click Connect. 

I could save the login details in a key store and reference them in the command so it’s secure, just not sure if it would work though. 

Or the only thing I can think of is a macro...just trying to get everything done in a 1-liner rather than a few stepped macro.
