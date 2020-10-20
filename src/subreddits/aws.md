# aws
## [1][AWS Wish List 2020](https://www.reddit.com/r/aws/comments/jbw85e/aws_wish_list_2020/)
- url: https://www.reddit.com/r/aws/comments/jbw85e/aws_wish_list_2020/
---
&amp;#x200B;

AWS always releases a bunch of features, sometimes everyday or atleast once a week. Here is my wish list of the features I want to see as a part of AWS infrastructure

1: AWS Managed Proxy Server(Rather than spinning own squid server)

2: EBS replication across different availability zones(Possible? Legal constraints?)

3: Multi-region VPC(Possible? Legal constraints?)

4: UI to debug boot issues(Better then EC2 Get Instance Screenshot and Instance logs)

5: Support tagging for every individual service(It's improving)

6: VPC endpoints support for every service (EKS?) 

7: EC2 instance live migration

8: Display AWS Cli  while resource creation(Similar to GCP)

9: Cost calculation while resource creation(AWS start supporting(for example, RDS) this feature but not for every service

10: More features in App Mesh(Circuit breaker, Rate Limiting)

P.S: Not sure if some features are already available, but if something is missing, please feel free to add
## [2][Week of Oct 19th - What have you learned recently about AWS?](https://www.reddit.com/r/aws/comments/je2wmx/week_of_oct_19th_what_have_you_learned_recently/)
- url: https://www.reddit.com/r/aws/comments/je2wmx/week_of_oct_19th_what_have_you_learned_recently/
---
Weekly Discussion post   
Sharing is caring
## [3][How to SAML federate your AWS account with G Suite](https://www.reddit.com/r/aws/comments/jencr5/how_to_saml_federate_your_aws_account_with_g_suite/)
- url: https://blog.pethron.me/how-to-saml-federate-your-aws-account-with-g-suite
---

## [4][Is there a way to output all the security groups and NACLs to something like a spreadsheet?](https://www.reddit.com/r/aws/comments/jednxz/is_there_a_way_to_output_all_the_security_groups/)
- url: https://www.reddit.com/r/aws/comments/jednxz/is_there_a_way_to_output_all_the_security_groups/
---
The only script I have dumps me out a json file that is like 15,000 lines. I just want something more easily read by a human.
## [5][Should I one put their Application Load Balancer behind CloudFront?](https://www.reddit.com/r/aws/comments/jelo7r/should_i_one_put_their_application_load_balancer/)
- url: https://www.reddit.com/r/aws/comments/jelo7r/should_i_one_put_their_application_load_balancer/
---
Is it a best practice to put an Application Load Balancer behind CloudFront as a dynamic origin? Are there factors one should consider?

Also, if one were to serve a webapp at `www.site.com` backed by S3 and wanted to serve the api at `api.site.com` backed by a server behind a load balancer, that would entail two separate CloudFront distributions right? Because I couldn't find how to keep it on just one with a behavior that chose the origin based on the hostname. Are there considerations to keep in mind from creating two?
## [6][Nginx logs into cloudwatch not appearing](https://www.reddit.com/r/aws/comments/jenfcr/nginx_logs_into_cloudwatch_not_appearing/)
- url: https://www.reddit.com/r/aws/comments/jenfcr/nginx_logs_into_cloudwatch_not_appearing/
---
I am trying to load the nginx error logs and access logs into cloudwatch. I have atatched an IAM role to the EC2 instance, installed cloudwatch agent &amp; ran the cloudwatch wizard which has stored the cloudwatch agent config in SSM Parameter store. Also restarted the cloudwatch agent + nginx service. Cloudwatch appears to be 'recognized' as I have new Metrics available (disk\_used\_percent). Problem is, it hasn't created any logs, log groups or streams. Any suggestions much appreciated. First part of the cloudwatch config.json file below:

`{ "agent": {` 

`"metrics_collection_interval": 60,` 

`"run_as_user": "root"` 

`},` 

`"logs":` 

`{ "logs_collected": {` 

`"files": {` 

`"collect_list": [` 

`{` 

`"file_path": "/var/log/nginx/error.log",` 

`"log_group_name": "nginx error.log",` 

`"log_stream_name": "{instance_id}"` 

`},` 

`{` 

`"file_path": "/var/log/nginx/access.log",` 

`"log_group_name": "nginx access.log",`

`"log_stream_name": "{instance_id}" }`
## [7][which rds should i use?](https://www.reddit.com/r/aws/comments/jem836/which_rds_should_i_use/)
- url: https://www.reddit.com/r/aws/comments/jem836/which_rds_should_i_use/
---
The project I am working on requires me to have a relational database for a small business. But the requests made to the database is very low somedays none. Which form of rds will be the most helpful to keep the prices as low as possible.
## [8][s3fs gets unmounted when files pushed to its directory](https://www.reddit.com/r/aws/comments/jeozu1/s3fs_gets_unmounted_when_files_pushed_to_its/)
- url: https://www.reddit.com/r/aws/comments/jeozu1/s3fs_gets_unmounted_when_files_pushed_to_its/
---
Mounted a s3 bucket using s3fs in a SUSE Linux machine, we have a sap system there that pushes files to that directory which then gets uploaded in the s3 bucket. Whenever we push these files, the s3 gets unmounted and it fails. Any other solution for this scenario since s3fs seems unreliable.
## [9][Web hook strategies](https://www.reddit.com/r/aws/comments/jeo07n/web_hook_strategies/)
- url: https://www.reddit.com/r/aws/comments/jeo07n/web_hook_strategies/
---
Hi all, my company is embarking on moving stuff from our DC into AWS. Dipping our toes in the water, one of the first things we want to do is move our mass mailing structures partially into AWS.

We are looking at upgrading our email sending (via SendGrid) from approx 10 emails per second to probably on the order of 1000+ per second. For those familiar with ESPs you’ll know they return data in real time by having you provide a callback webhook URL where they dump json events/status updates.

Currently we handle this by having a load balanced local website that just drops the request to disk and then another device asynchronously processes the result an distributes them to where’ve customer database it belongs to.

Our current thoughts are to use API gateway and then deliver the payloads into SQS, we would then continue to asynchronously process the results like we do today.

For the purpose of this upgrade, we just want to make sure we aren’t ‘losing’ we hook requests or overloading our servers with capacity, we aren’t yet moving the db’s to the cloud nor attempting to containerise the processing logic or something like that, just baby steps.

Are the 2 services (API gateway and SQS) the right tools for this job?
## [10][Need help determining why AWS SAM Hello World app does not deploy in ap-southeast-1](https://www.reddit.com/r/aws/comments/jeo017/need_help_determining_why_aws_sam_hello_world_app/)
- url: https://www.reddit.com/r/aws/comments/jeo017/need_help_determining_why_aws_sam_hello_world_app/
---
I am following the instructions outlined by AWS [here](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-getting-started-hello-world.html) to deploy the AWS SAM Hello World (Python 3.8) template in ap-southeast-1. However, when I try to access the API, I am greeted with the following:

    {"message":"Forbidden"}

I have tried deploying in us-east-1 as well as ap-northeast-1 and the API returns the hello world message as expected. 


AWS documentation shows that SAM is supported in ap-southeast-1. I have also checked the deployed configs on API Gateway and Lambda to be the same in all regions.


I am turning to this sub now as I am at wits end trying to understand why I get a Forbidden message on ap-southeast-1 when everything was deployed successfully. Thank you in advance for any advice!
## [11][API Gateway blocks CORS requests, but still invokes Lambda integration](https://www.reddit.com/r/aws/comments/jenzqx/api_gateway_blocks_cors_requests_but_still/)
- url: https://www.reddit.com/r/aws/comments/jenzqx/api_gateway_blocks_cors_requests_but_still/
---
Hey, I have a simple contact form setup on my personal website. I have API Gateway CORS configured to only take requests from my site. Everything seems to be working, when I try to send something from my local network, Chrome tells me:

    Access to fetch at 'xxx' from origin 'http://localhost:3000' has been blocked by CORS policy: No 'Access-Control-Allow-Origin' header is present on the requested resource. If an opaque response serves your needs, set the request's mode to 'no-cors' to fetch the resource with CORS disabled.

So everything seems good, but yet the Lambda is still invoked and I wind up getting the email anyway. Am I misunderstanding something about how this works?
## [12][ELB: Trailing slash (/) issue :: without slash(/) site not resolving](https://www.reddit.com/r/aws/comments/jenqts/elb_trailing_slash_issue_without_slash_site_not/)
- url: https://www.reddit.com/r/aws/comments/jenqts/elb_trailing_slash_issue_without_slash_site_not/
---
Hi,

***Scenario:***

* My website resolves to Aws ELB with 443 listener 
* Backend Target group is an EC2 instance with apache running on port 8080

***Problem***:

* If I miss the trailing slash, the website does not resolve and when i put the trailing slash (/) everything works fine.
* [www.example.com/foldername](https://www.example.com/folder)  &lt;-- ***Does not work***, returns url [www.example.com:8080/foldername/](http://sg-test.pushtrack.co:8080/TwoClickDemo/) and resolution fails
* [www.example.com/foldername](https://www.example.com/folder)/   &lt;-- Does Works fine
*  To check if everything fine on the backend I queried the EC2 instance IP on port 8080 with and without /
* http://ec2\_ip:8080/foldername  and  http://ec2\_ip:8080/foldername/   
*  Above both url with instance ip worked as expected

I'm guessing something fails between ELB and EC2 instance 

Has anyone faced similar issue before?
