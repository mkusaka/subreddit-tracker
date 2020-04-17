# aws
## [1][AWS IQ waives fees until June 30, 2020, to help you stand up and scale remote work initiatives](https://www.reddit.com/r/aws/comments/g2prcr/aws_iq_waives_fees_until_june_30_2020_to_help_you/)
- url: https://aws.amazon.com/blogs/machine-learning/aws-iq-waives-fees-until-june-30-2020-to-help-you-stand-up-and-scale-remote-work-initiatives/
---

## [2][Elastic Beanstalk: Can I deploy using Dockerfiles instead of a Dockerrun.aws.json (v2) file when using a multicontainer Docker setup?](https://www.reddit.com/r/aws/comments/g2xzww/elastic_beanstalk_can_i_deploy_using_dockerfiles/)
- url: https://www.reddit.com/r/aws/comments/g2xzww/elastic_beanstalk_can_i_deploy_using_dockerfiles/
---
In Elastic Beanstalk in Single Container Docker configuration, it is possible to deploy without a `Dockerrun.aws.json` (v1) file as long as a Dockerfile is supplied.

Is this also possible in Multicontainer Docker setups? I prefer not to build the Docker images on my own. I just want to supply the required Dockerfiles and let Amazon build the images for me, just like what is possible in the case of Single Container Docker configuration.

Does anyone know?
## [3][Direct connect (DX) vs T1 internet provider](https://www.reddit.com/r/aws/comments/g2ywa8/direct_connect_dx_vs_t1_internet_provider/)
- url: https://www.reddit.com/r/aws/comments/g2ywa8/direct_connect_dx_vs_t1_internet_provider/
---
Hi folks,

We're currently investigating wether or not it's useful to invest in to DX. Currently the company I work for has a nice multi-gb internet connection through a T1 ISP. Bandwidth is not really the issue.

Traffic is secured using VPN's where possible.

Besides the guarantees you get on bandwidth and latency, and the lower cost of egress pricing, I'm not finding that many incentives to go this route.... or am I missing something?

Is there any difference in how AWS handles traffic coming from DX vs Internet?
## [4][How to "snapshot" a patch status with Systems Manager?](https://www.reddit.com/r/aws/comments/g31vo2/how_to_snapshot_a_patch_status_with_systems/)
- url: https://www.reddit.com/r/aws/comments/g31vo2/how_to_snapshot_a_patch_status_with_systems/
---
So we're looking to implement patching with Systems Manager. Technically it seems fine, but I'm struggling with figuring out the best way to implement what I think is a seemingly simple pattern.

1. Manually or automatically (e.g. scheduled for a specific day of the month) approve the latest group of patches for a set of non-production systems.
2. Kick of automated testing of non-production systems.
3. Assuming testing doesn't reveal any issues, approve the **same** updates that were installed on non-production systems to be installed on production systems.

The only thing I can come up with so far is running AWS-RunPatchBaseline on non-production systems on a specific day of the month (with a patch baseline configured to install all patches released up to that day), and then run AWS-RunPatchBaseline on production systems some number of days later and specify to install patches released up to the same date that AWS-RunPatchBaseline was run in non-production. This seems somewhat error prone and less than ideal.

I was able to implement a similar pattern in the past with Linux repo mirrors (and aptly) by just creating a snapshot, release it to non-prod, waiting for testing, and then releasing the same snapshot of patches to prod (and just waiting for nightly patch installation to pick up the newly released packages). Is there any way to implement something similar with Systems Manager that I'm missing?

Thanks!
## [5][Lambda function not able to handle load tests.](https://www.reddit.com/r/aws/comments/g30ty0/lambda_function_not_able_to_handle_load_tests/)
- url: https://www.reddit.com/r/aws/comments/g30ty0/lambda_function_not_able_to_handle_load_tests/
---
I am trying to load test my lambda function (connected to gateway) using JMeter. On the function, I have set the concurrency to "Use unreserved account concurrency" (default 1000 for my account).

I have set the following config in JMeter thread worker:

* Number of Threads (users): 1000

- Ramp Up Time: 0 secs (Because I want all the requests to hit at once. Please correct me if I'm wrong here)

- Input: I have provided a CSV with 4000 rows, so it will make 4000 requests in total

Most of my requests are failing. [Here is a screenshot showing how many requests fail along with the response received](https://imgur.com/GT72QQJ).

In production, I will have to handle such amount of concurrent requests. What can I do to make this work?

The lambda function is not compute heavy - all it does is reads a dynamo table and pushes the event payload to a SQS queue. So I don't think there is any space left to optimize there
## [6][Data Transfer Out (DTO) 40% Price Reduction in South America (São Paulo) Region](https://www.reddit.com/r/aws/comments/g2esea/data_transfer_out_dto_40_price_reduction_in_south/)
- url: https://aws.amazon.com/blogs/aws/aws-data-transfer-out-dto-40-price-reduction-in-south-america-sao-paulo-region/
---

## [7][AWS Forecast error saying filter id is missing or invalid](https://www.reddit.com/r/aws/comments/g2zep1/aws_forecast_error_saying_filter_id_is_missing_or/)
- url: https://www.reddit.com/r/aws/comments/g2zep1/aws_forecast_error_saying_filter_id_is_missing_or/
---
Whenever I want to do a Forecast lookup I get the following error message:

&gt;Bad request ResourceNotFoundException : The query did not return any forecast results as the required filter id is either missing or is invalid. 

I'm clueless about why I get this error. I checked my Forecast key (which is item\_id) and value (which is F11) and they fit with the .csv I used for my dataset. I also checked for the order of my attributes that I set for my dataset and they also fit with my .csv.

An example line in my csv looks like this:

    2016-12-07, 2, LAE, F11, 190, 190, 215, 58, 58, 79, 0, 23, **** 

And my data schema looks like this:

    {
    	"Attributes": [
    		{
    			"AttributeName": "timestamp",
    			"AttributeType": "timestamp"
    		},
    		{
    			"AttributeName": "werk",
    			"AttributeType": "string"
    		},
    		{
    			"AttributeName": "messpunkt",
    			"AttributeType": "string"
    		},
    		{
    			"AttributeName": "item_id",
    			"AttributeType": "string"
    		},
    		{
    			"AttributeName": "PL_atag",
    			"AttributeType": "string"
    		},
    		{
    			"AttributeName": "SO_atag",
    			"AttributeType": "string"
    		},
    		{
    			"AttributeName": "demand",
    			"AttributeType": "float"
    		},
    		{
    			"AttributeName": "plumlauf",
    			"AttributeType": "string"
    		},
    		{
    			"AttributeName": "sumlauf",
    			"AttributeType": "string"
    		},
    		{
    			"AttributeName": "umlauf",
    			"AttributeType": "string"
    		},
    		{
    			"AttributeName": "diff_mpp",
    			"AttributeType": "string"
    		},
    		{
    			"AttributeName": "diff_rp",
    			"AttributeType": "string"
    		},
    		{
    			"AttributeName": "fst_fber",
    			"AttributeType": "string"
    		}
    	]
    }

So where is the mistake here?
## [8][Amazon Chime: max participants with two way video?](https://www.reddit.com/r/aws/comments/g2jsv1/amazon_chime_max_participants_with_two_way_video/)
- url: https://www.reddit.com/r/aws/comments/g2jsv1/amazon_chime_max_participants_with_two_way_video/
---
Hello, was going through the docs to consider this solution for an upcoming opportunity and was not clear on the above point.

It says a user can host meeting with upto 250 attendees, but doesn't explicitly say whether all 250 can have interactive video or is it more like viewers with only upto x panelists/two-way speakers in the meeting. There are references of upto 16 video streams on first come basis, but it's not clear whether they only allow 16 users with a camera or its 16 can be displayed on the screen at a time and it wd dynamically switch to active ones/speaking ones i.e. all 250 can have two way video&amp;audio, but on the screen one can view upto recent 16 active users at any time. Similarly, on the SDK front, it talks about 100 attendees per meeting with 100 audio streams n 16 video streams.

So, if I am a paid user who is hosting a  video meeting, how many participants (with video) can I have in my meeting? Bluntly put, if I subscribe to $15/month as a host, what are rhe limits on the video conf participants for my meetings?

Can someone please clarify as to how many simultaneous participants (with two way audio n video) can join a meeting on chime? What are the other limitations / charges for a web based video conferencing / webinar solution (without any external pstn based voice etc)?

Thanks
## [9][Leveraging ULIDs to create order in unordered datastores (like S3)](https://www.reddit.com/r/aws/comments/g2mh01/leveraging_ulids_to_create_order_in_unordered/)
- url: https://www.trek10.com/blog/leveraging-ulids-to-create-order-in-unordered-datastores
---

## [10][Anyway to tag nodes in ASG set differently in CloudFormation?](https://www.reddit.com/r/aws/comments/g2jtko/anyway_to_tag_nodes_in_asg_set_differently_in/)
- url: https://www.reddit.com/r/aws/comments/g2jtko/anyway_to_tag_nodes_in_asg_set_differently_in/
---
I am using ASG w/ LaunchTemplate to create 3 systems in CloudFormation Script.  Is there a way I can tag these differently as alpha1, alpha2, alpha3, or add other tags, such a role tag?

I have a cluster, i have to configure nodes differently in the user data script, where one is a instance is a leader, while other instances are peers.  Additionally, on one of the systems, I would like to put a small management web service on it to reduce costs.
