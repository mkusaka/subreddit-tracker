# aws
## [1][Howto remove Public Ip of instance created with “Auto-assign Public IP”?](https://www.reddit.com/r/aws/comments/fhy5x7/howto_remove_public_ip_of_instance_created_with/)
- url: https://www.reddit.com/r/aws/comments/fhy5x7/howto_remove_public_ip_of_instance_created_with/
---
I’m bringing env to private but I don’t find a way to remove public IP address of an instance once I had enabled vpn for private access to VPC. 

I’m using default vpc. 

So howto?
## [2][Building faster, lower cost, better APIs – HTTP APIs now generally available | Amazon Web Services](https://www.reddit.com/r/aws/comments/fhnh0g/building_faster_lower_cost_better_apis_http_apis/)
- url: https://aws.amazon.com/blogs/compute/building-better-apis-http-apis-now-generally-available/
---

## [3][Can a lightsail load balancer handle websocket connections?](https://www.reddit.com/r/aws/comments/fhxg15/can_a_lightsail_load_balancer_handle_websocket/)
- url: https://www.reddit.com/r/aws/comments/fhxg15/can_a_lightsail_load_balancer_handle_websocket/
---
I have a high bandwidth websocket application that isnt feasible for me to run on ordinary EC2 due to the bandwidth costs being too high - this app does about 15TB per month. 

I'm looking at Lightsail that seems a good fit - an 8gb node comes with 5TB of traffic, so three of those nodes would sort me out. I understand Lightsail throttles sustained CPU usage but it seems like an 8gb node even if throttled down to 30% will still be fine for me as websockets aren't CPU intensive.

What I need to understand is:

1. Is the lightsail loadbalancer compatible with secure websockets (wss)?
2. Are there hidden charges I need to worry about? In addition to running Lightsail nodes for the websocket app, the main app itself will be hosted on EC2 nodes as the main app is highly demanding on CPU. Therefore I need the lightsail nodes to be able to access Redis and RDBS - I believe this can be done through the right VPC settings **but does it have charges?**
3. How much load can the lightsail loadbalance handle? Can lightsail in general handle big production traffic?

Thanks for any help!
## [4][SSO logs me out too quickly](https://www.reddit.com/r/aws/comments/fhvu1w/sso_logs_me_out_too_quickly/)
- url: https://www.reddit.com/r/aws/comments/fhvu1w/sso_logs_me_out_too_quickly/
---
https://signin.aws.amazon.com keeps me logged in for a day or 12 hours. So SSO feels a little two steps forward, one step back.

I'm also a little puzzled how in Firefox I can keep multiple accounts open in one Profile session.

Any tips guys?
## [5][Donating AWS services to charities?](https://www.reddit.com/r/aws/comments/fhjs7v/donating_aws_services_to_charities/)
- url: https://www.reddit.com/r/aws/comments/fhjs7v/donating_aws_services_to_charities/
---
I am looking to start giving back to charitable causes. 

Are there any charities out there that are in need of an IT guy skilled in AWS and cloud services?
## [6][Internet Lag in AWS WorkSpaces and AppStream](https://www.reddit.com/r/aws/comments/fhtlih/internet_lag_in_aws_workspaces_and_appstream/)
- url: https://www.reddit.com/r/aws/comments/fhtlih/internet_lag_in_aws_workspaces_and_appstream/
---
I did a survey on our users using WorkSpaces and AppStream. The major concern we received is Internet lag. Some of our users use laptops to do their work and lags when using these services. I notice the same when using laptops, but not on desktops.

Is there to mitigate this issue?
## [7][Bug report: `aws apigateway get-export` sometimes returns reordered items (no harm but noise in Git)](https://www.reddit.com/r/aws/comments/fhu7k6/bug_report_aws_apigateway_getexport_sometimes/)
- url: https://www.reddit.com/r/aws/comments/fhu7k6/bug_report_aws_apigateway_getexport_sometimes/
---
Good day, I have a file `pull.sh`:

    AWS_PROFILE=gms-ai \
        aws apigateway get-export \
            --rest-api-id xxxxxxxxxx \
            --stage-name prod \
            --export-type swagger \
            --parameters '{"extensions": "integrations,authorizers"}' \
            --accepts application/yaml \
            swagger.yaml

And sometimes after executing I see this sad picture from `git diff`:

[Pointless reorder :\(](https://preview.redd.it/4rzjc374ddm41.png?width=1170&amp;format=png&amp;auto=webp&amp;s=e520b21119fb07679c0e2302c7199634ae19760e)

Not mentioning that it is better to see body mapping template in readable/editable form, without  \\n\\n\\n.
## [8][Can you use AWS Workspaces to reach on premises resources over VPN, not just AD.](https://www.reddit.com/r/aws/comments/fhq0wk/can_you_use_aws_workspaces_to_reach_on_premises/)
- url: https://www.reddit.com/r/aws/comments/fhq0wk/can_you_use_aws_workspaces_to_reach_on_premises/
---
Can AWS workspaces access on-premises resources like file shares etc over Direct connect or VPN? I know it can auth against AD, but what about everything else?
## [9][Is there a way to customize the lambda by installing required libraries and compilers ?](https://www.reddit.com/r/aws/comments/fhu0my/is_there_a_way_to_customize_the_lambda_by/)
- url: https://www.reddit.com/r/aws/comments/fhu0my/is_there_a_way_to_customize_the_lambda_by/
---
I would like to use Lambda services but my use case need a compiler installed in lambda but i am not sure if its possible.
## [10][As an Office 365 shop, what are compelling reasons to choose AWS over Azure?](https://www.reddit.com/r/aws/comments/fhlkfr/as_an_office_365_shop_what_are_compelling_reasons/)
- url: https://www.reddit.com/r/aws/comments/fhlkfr/as_an_office_365_shop_what_are_compelling_reasons/
---
We have been using O365 for about 1.5 years.  This primarily includes Exchange for email and SharePoint for our intranet and other internal documents.  We are a health care organization and also a Windows shop.

We use AD Connect to sync our accounts, passwords, and groups to the O365 cloud but still have our 2016 DCs on-prem.  All of our other servers / services / etc. are hosted on-prem using VMWare and Horizon View.  

On my own time, I have recently earned 2 AWS certs and have been playing around with AWS for my own side projects and really like it.  I only have just a small bit of experience with Azure.  My experience was not as enjoyable, but I honestly have not spent nearly as much time with Azure as I have with AWS.

The cost of Office 365 is not a concern as we've been very happy with the results.  Are there any compelling reasons to use AWS vs Azure given our current situation?  If so, I'd love to hear about them, but it seems like future integration with Azure would be easier.  AWS seems like it would be a hard sell to others in our org.

Either way, we would most certainly keep using O365.

Thanks.

**EDIT:** We are not a software company, but rather more akin to a small hospital -- but no overnight stays or surgeries.
