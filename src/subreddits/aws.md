# aws
## [1][Week of Aug 24th - What are you building this week on AWS?](https://www.reddit.com/r/aws/comments/ifoebu/week_of_aug_24th_what_are_you_building_this_week/)
- url: https://www.reddit.com/r/aws/comments/ifoebu/week_of_aug_24th_what_are_you_building_this_week/
---
Share what you are creating on AWS
## [2][WDYT about Streaming/Recording of workshops walkthrough](https://www.reddit.com/r/aws/comments/ihbpu6/wdyt_about_streamingrecording_of_workshops/)
- url: https://www.reddit.com/r/aws/comments/ihbpu6/wdyt_about_streamingrecording_of_workshops/
---
Hello fellow members of /r/aws, I've created the https://awesome-aws-workshops.com portal and I'm polling the community to see if there's any interest of watching videos/live streaming of workshop walkthrough commented by me or any other guest.

The idea is to have weekly workshops done depending on the demand and eventually integrate that back into the portal.

Workshops that are larger, might be divided in video series (think about the EKS and ECS ones with multiple modules)

If you don't think this is a good use of your time and also my time, please do comment bellow on what would make the portal more interesting for you.

[View Poll](https://www.reddit.com/poll/ihbpu6)
## [3][Managed AD vs Stand Alone/Dedicated DCs?](https://www.reddit.com/r/aws/comments/ihk7us/managed_ad_vs_stand_alonededicated_dcs/)
- url: https://www.reddit.com/r/aws/comments/ihk7us/managed_ad_vs_stand_alonededicated_dcs/
---
Hi all,

I can't seem to find any concrete rational for using AWS' Managed AD solution over Dedicated DCs or Vice Versa. I'm wondering if anyone out there has any opinions, insight, or things we should be considering when making this decision.

I know there are cost implications to going the dedicated route and Managed AD doesn't have the same configuration capabilities (like GPO's) as what stand-alone DCs would have. 

Thanks much for your time and help!
## [4][How do you handle lots of lambda functions with dependencies in a project?](https://www.reddit.com/r/aws/comments/ihbr68/how_do_you_handle_lots_of_lambda_functions_with/)
- url: https://www.reddit.com/r/aws/comments/ihbr68/how_do_you_handle_lots_of_lambda_functions_with/
---
I have a project with about \~25 Node lambdas in a repo, is it bad that each one has their own package.json and node\_modules? How do you organize large serverless projects and their dependencies? Looking for inspiration. Thanks.
## [5][AWS Server Side Encryption](https://www.reddit.com/r/aws/comments/ih7aot/aws_server_side_encryption/)
- url: https://www.reddit.com/r/aws/comments/ih7aot/aws_server_side_encryption/
---
Other than compliance, what real risk are we mitigating by using server side encryption? AWS employees removing a hard drive with your data on it? It seems so extremely unlikely to even be worth considering. Thoughts?
## [6][ballpark direct connect costs?](https://www.reddit.com/r/aws/comments/ihe20x/ballpark_direct_connect_costs/)
- url: https://www.reddit.com/r/aws/comments/ihe20x/ballpark_direct_connect_costs/
---
Anyone care to share approximate cost to provision a 1Gbps direct connect connection from a location about 35 miles outside of NYC to the CoreSite NY1 location?  

I'm trying to determine feasibility--I understand I won't know exact carrier costs until I contact, well, a carrier.  I'm just looking for anecdotal pricing if possible.  Thanks in advance!
## [7][How to connect on premise datacenter to AWS S3](https://www.reddit.com/r/aws/comments/ihec4i/how_to_connect_on_premise_datacenter_to_aws_s3/)
- url: https://www.reddit.com/r/aws/comments/ihec4i/how_to_connect_on_premise_datacenter_to_aws_s3/
---
Hi All

I was looking at different ways to connect our on premise datacenters to access S3.

One option was to route all the AWS public S3 subnets using public VIF back to the on prem DC. 

Option 2 was to connect a storage gateway to the direct connect.

&amp;#x200B;

With the storage gateway, there is another additional failure domain where you have to manage the virtual appliances and also becomes a critical device to manage in case the HA VM's go down.

Could someone please explain what are the pros and cons with the 2 options and which one is normally preferred?

&amp;#x200B;

Thanks in advance. I'm still reading through docs but nothing beats real world scenarios from hearing experiences from the community.
## [8][Website for 3000 visits at the same time with AWS](https://www.reddit.com/r/aws/comments/ihe929/website_for_3000_visits_at_the_same_time_with_aws/)
- url: https://www.reddit.com/r/aws/comments/ihe929/website_for_3000_visits_at_the_same_time_with_aws/
---
Hello everyone, I would like to ask your help with something.

The company I work for decided to hold a virtual congress in October, and part of that congress includes an expo, for that we will have to create 2 web pages (wordpress) that at its peak will reach at least 3000 visits at the same time. We are thinking of using AWS for this, but we don't really know which service would be the best for this scenario.

We had thought about Lightsail, but we are not sure.

Any suggestion?
## [9][Can you hide the bucket region in the URL on S3?](https://www.reddit.com/r/aws/comments/ihgilo/can_you_hide_the_bucket_region_in_the_url_on_s3/)
- url: https://www.reddit.com/r/aws/comments/ihgilo/can_you_hide_the_bucket_region_in_the_url_on_s3/
---
https://bucket.s3.us-east-1.amazonaws.com/joespdf.pdf
## [10][Hosting a Flask API on EC2 - best tips/tricks - basic questions](https://www.reddit.com/r/aws/comments/ih6dqw/hosting_a_flask_api_on_ec2_best_tipstricks_basic/)
- url: https://www.reddit.com/r/aws/comments/ih6dqw/hosting_a_flask_api_on_ec2_best_tipstricks_basic/
---
Hey guys, cross-posted this to r/learnpython but this seems like a more relevant subreddit actually. Apologies if this isn't the correct place for it.

I'm hosting a simple flask API on an EC2 instance.

When you call it, it launches a headless browser in selenium that then loads a website, scrapes some info, and returns it to the user. I'm expecting traffic of occasionally up to 10 people calling it in a given second.

I have a few questions about this:

1 - What is the best practice for hosting this? Do I just run the python script in a tmux shell and then leave it running when I disconnect from my ssh to the EC2? Or should I be using some fancy tool to keep it running when I'm not logged in such as systemd

2 - How does Flask handle multiple queries at once? Does it automatically know to distribute queries separately between multiple cores? If it doesn't, is this something I could set up? I have no great understanding of how an API hosted on EC2 would handle even just two requests simultaneously.

3 - A friend mentioned I should have a fancier setup involving the API hosted behind an nginx which serves requests to dif versions of it or something like this, what's the merit in this?

Thank you kindly, would love to know the best practise here and there's surprisingly little documentation on the industry standards.

Best regards and thanks in advance for any responses

(Side note: When I run it, it says WARNING: Do not use the development server in a production environment. This makes me think I'm probably doing something wrong here? Is flask not meant to be used in production like this?)
## [11][why is cloudformation support not totally automated by now?](https://www.reddit.com/r/aws/comments/ih76o9/why_is_cloudformation_support_not_totally/)
- url: https://www.reddit.com/r/aws/comments/ih76o9/why_is_cloudformation_support_not_totally/
---
Lately, I'm seeing the aws terraform provider being updated much faster than cloudformation can support AWS' own features (for example, elasticsearch service warm node support).  As someone stuck in a world of cloudformation, that's hilarious and frustrating.  But isn't the question why does support have to wait?  Couldn't AWS have created some generic module that the different services could code their configuration knobs towards?  Instant CFN support would make this platform much more attractive.
