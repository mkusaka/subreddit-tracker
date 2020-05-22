# aws
## [1][Is there a way to open the services menu with a keyboard shortcut?](https://www.reddit.com/r/aws/comments/goi6og/is_there_a_way_to_open_the_services_menu_with_a/)
- url: https://www.reddit.com/r/aws/comments/goi6og/is_there_a_way_to_open_the_services_menu_with_a/
---
Seriously, if I want to navigate to a new service I have to take my hands off the keyboard and click the menu icon before I can search! 

Am I just an idiot? What’s the point of auto focusing the search bar if you still have to use the mouse to open the services menu in the first place?
## [2][How AWS Web application firewall (WAF) was implemented to prevent an in-production bug; and how you can implement it in less than 5 minutes](https://www.reddit.com/r/aws/comments/go3974/how_aws_web_application_firewall_waf_was/)
- url: https://medium.com/swlh/how-aws-waf-saved-us-6dc21343f1f5
---

## [3][What should i use to cache my API response ( Beginner )](https://www.reddit.com/r/aws/comments/gog8du/what_should_i_use_to_cache_my_api_response/)
- url: https://www.reddit.com/r/aws/comments/gog8du/what_should_i_use_to_cache_my_api_response/
---
Hi All,

I am making an App for my Univ, where Students can see all the events going on in the Univ.

The website is hosted on AWS Lightsail ($3.5/month) plan. As the website will only be visited by few organizers that  have the ability to post events , i choose the low end server. 

The problem is with the App , as it makes API calls to the server to get the details about the events and most of them are same response . Even bigger problem when a new event is posted, everyone will make the same request at the same time.

For around 2000 users what can i use that will be able to caches the API requests and only request the server if a new one is made. Also if the event is updated i want to be able to tell the caching server to drop the current cache for the event.

Or should i just upgrade to a better AWS Lightsail to handel all users.
## [4][Global Accelerator Experience?](https://www.reddit.com/r/aws/comments/go8dv4/global_accelerator_experience/)
- url: https://www.reddit.com/r/aws/comments/go8dv4/global_accelerator_experience/
---
It sure as hell makes a difference for me. 

&amp;#x200B;

&amp;#x200B;

https://preview.redd.it/zmpkmi8fi7051.png?width=917&amp;format=png&amp;auto=webp&amp;s=9cbe47e75c0fa5ae745db0c61a3e6ea40986928e

https://preview.redd.it/9f89sg8fi7051.png?width=895&amp;format=png&amp;auto=webp&amp;s=92005f27e8c7c9dc1ea10217f91379a57250d475

https://preview.redd.it/tm9x0g8fi7051.png?width=930&amp;format=png&amp;auto=webp&amp;s=9f13c9d024b0feb6da4ac62285ce6c0c1996c1df

https://preview.redd.it/6g0mmf8fi7051.png?width=976&amp;format=png&amp;auto=webp&amp;s=4470f9c9fc2475ad94c39528150650128d273f34
## [5][If I use S3 Get Size in the console, which cost do I incur?](https://www.reddit.com/r/aws/comments/gohp3z/if_i_use_s3_get_size_in_the_console_which_cost_do/)
- url: https://www.reddit.com/r/aws/comments/gohp3z/if_i_use_s3_get_size_in_the_console_which_cost_do/
---
Is the request made GET or LIST HTTP request? If it's LIST, then it's 10 times more expensive than I thought :D
## [6][How to enable Two way SSL in the AWS Load Balancer? Or is there an alternative service by AWS which supports that?](https://www.reddit.com/r/aws/comments/gohiwl/how_to_enable_two_way_ssl_in_the_aws_load/)
- url: https://www.reddit.com/r/aws/comments/gohiwl/how_to_enable_two_way_ssl_in_the_aws_load/
---
I have recently started a new job, and have been asked to research for a way to enable two way SSL on an AWS Load Balancer. I have never personally used AWS services as I am a developer, and not someone who knows the infrastructure/DevOps side much. I know that AWS Load Balancer does not support two way SSL, so I instead said that we can go for nginx which supports that and then route the request to the AWS Load Balancer which would be configured to accept the requests only from the nginx server. 

Wondering if there's any service which supports LB and two way SSL out of the box by AWS.

Any other suggestions are also welcome!
## [7][DynamoDB for fetching API key from user details table](https://www.reddit.com/r/aws/comments/gohdvd/dynamodb_for_fetching_api_key_from_user_details/)
- url: https://www.reddit.com/r/aws/comments/gohdvd/dynamodb_for_fetching_api_key_from_user_details/
---
Hi - I'm still fairly new to dynamodb coming from a largely MySQL background, so I was wondering if this is the right use case.

I'm looking at using dynamodb to store user information and an api key.

Structure would be as follows:

key for this table would be user cognito id, and then data would include first name, second name, email and an API key.

When my API gets a request, I need to be able to fetch the user cognito id from the API key.

As far as I'm aware, the best way to do this would be to have a second table called API Keys, where key would be API key, and the data stored would just be user id. - Issue with this is ensuring both of these tables are always kept in sync.

Does this sound right?

&amp;#x200B;

Thanks in advance.
## [8][Securing EC2 instances that host HTTP/S &amp; SFTP services](https://www.reddit.com/r/aws/comments/go75j1/securing_ec2_instances_that_host_https_sftp/)
- url: https://www.reddit.com/r/aws/comments/go75j1/securing_ec2_instances_that_host_https_sftp/
---
 Hi All,

Have some EC2 instances that are hosting HTTP/S and SFTP services. Since these were recently moved to AWS, we have been seeing a pretty good deal of port scanning going on - and in a couple of instances it may have caused in interruption in service. We are not necessarily AWS savvy - we do have an MSP that officially manages AWS for us - but I want to ask for some perspective here :)

Certainly, provisioning a WAF for HTTP/S services would be an improvement, but in our use case we need to support very large file uploads and downloads. This has been a limitation of WAF in the past, and last I asked our vendor, still is. Along with that, a WAF would not help as far as SFTP traffic goes - and I am not sure what to do there.

From what I can see, an NLB and security groups are all that are protecting these instances at the moment, so possibly additional firewall or IDS/IPS services would be the next steps?

Thanks in advance;
## [9][Cognito - User pool per tenant, or one user pool for all tenants](https://www.reddit.com/r/aws/comments/gog02b/cognito_user_pool_per_tenant_or_one_user_pool_for/)
- url: https://www.reddit.com/r/aws/comments/gog02b/cognito_user_pool_per_tenant_or_one_user_pool_for/
---
Hi! As the title suggests I am having a hard time diciding between using a single user for all of my SaaS system's tenants, or apon team creation provisioning a user pool for that team.   


All other AWS resources will be shared, such as DynamoDB, the API Gateway, and the webpage will be served from a single S3 bucket (with Cloud Front).  


Does anyone have any advice/experience that would point one way or the other?
## [10][Serverless video processing architecture?](https://www.reddit.com/r/aws/comments/godb2v/serverless_video_processing_architecture/)
- url: https://www.reddit.com/r/aws/comments/godb2v/serverless_video_processing_architecture/
---
I want to record video from web clients. Don't know when recording ends and webapp is closed, so I need to send the video chunks in real time.

Staff member needs to perform manual video analysis after session is over. In the future video analysis maybe done by a machine AI.
Want to keep my architecture serverless (pay-as-you-go, auto-scaling).
Which aws service architecture do you recommend for this?
Thanks!
