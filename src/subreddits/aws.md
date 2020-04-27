# aws
## [1][I thought the AWS docs were unclear, so I tried to write a plain English guide to deploying from GitHub to EC2 using CodePipeline](https://www.reddit.com/r/aws/comments/g8n0mg/i_thought_the_aws_docs_were_unclear_so_i_tried_to/)
- url: https://seanjziegler.com/deploying-code-from-github-to-aws-ec2-with-codepipeline/
---

## [2][Serverless AD Password reset tool](https://www.reddit.com/r/aws/comments/g8zuw9/serverless_ad_password_reset_tool/)
- url: https://www.reddit.com/r/aws/comments/g8zuw9/serverless_ad_password_reset_tool/
---
Is it technically possible to create an AD user password reset tool using Lambda ? My concern is since Lambda's cannot be joined to the domain, will it work ? I can provide the admin credentials in a parameter store if that is required. 

Basically I am looking for something very simple like [this](https://github.com/jirutka/ldap-passwd-webui), but serverless. Are you aware of anything that can do this ? If it is technically possible to do it, I think I will invest some time do do it.
## [3][Middy, Node.js middleware framework for AWS Lambda, goes 1.0.0](https://www.reddit.com/r/aws/comments/g8hwsb/middy_nodejs_middleware_framework_for_aws_lambda/)
- url: https://loige.co/middy-1-is-here
---

## [4][Building an AWS IoT App with React.js](https://www.reddit.com/r/aws/comments/g8wk0h/building_an_aws_iot_app_with_reactjs/)
- url: https://www.reddit.com/r/aws/comments/g8wk0h/building_an_aws_iot_app_with_reactjs/
---
After floundering through the sea of AWS documentation, I finally came across AWS Amplify. I found it really useful and would have saved me a lot of time had I just used it in the first place.

I wanted to document my progress in building a real-time app for sensor data streaming to AWS IoT Core. Hoping others find inspiration or utility in my post.

[https://theskenengineering.com/building-a-react-js-app-with-aws-iot/](https://theskenengineering.com/building-a-react-js-app-with-aws-iot/)

Thank you for your time.
## [5][On-prem MS SQL](https://www.reddit.com/r/aws/comments/g8yaba/onprem_ms_sql/)
- url: https://www.reddit.com/r/aws/comments/g8yaba/onprem_ms_sql/
---
If I were to deploy MS SQL via the EC2 instance under the launch configuration, what sort of licensing model does it comes with? I was able to extract the license key using powershell script provided by AWS support, but I didn't seem to work when I was using it for PowerBI setup.

Does the MS SQL come with Software Assurance?
## [6][Load balancer architect question](https://www.reddit.com/r/aws/comments/g90ko9/load_balancer_architect_question/)
- url: https://www.reddit.com/r/aws/comments/g90ko9/load_balancer_architect_question/
---
We have our servers all set up the same way. One load balancer with an SSL cert, that points to a target group with 2 servers in it. Pretty simple. We have this setup replicated 4 times in one region as we have 4 variations on this same app. Would it be smarter to consolidate the load balancers and have 1 load balancer with 4 rules (with 4 certs) that point to the 4 target groups, or is it better to keep the load balancers separate? Just trying to figure out best practice and most cost-effective.

&amp;#x200B;

Thanks in advance!
## [7][Problem updating Workspaces Client](https://www.reddit.com/r/aws/comments/g8zxxg/problem_updating_workspaces_client/)
- url: https://www.reddit.com/r/aws/comments/g8zxxg/problem_updating_workspaces_client/
---
Anyone else having problems to update the Workspaces Client?

On the website there's an outdated version and the installer stops when it notices that.

When I open the client I can connect and use my workspace but if try to update it just won't do nothing.

I'm on Windows 10 with the latest updates my client is in version 3.0.2.1288 and newest version is 3.0.6.
## [8][Question about hosting via EC2?](https://www.reddit.com/r/aws/comments/g8wvqy/question_about_hosting_via_ec2/)
- url: https://www.reddit.com/r/aws/comments/g8wvqy/question_about_hosting_via_ec2/
---
Hi all,

&amp;#x200B;

I want to host my website on an EC2 instance using AWS and want to be able to update the server via a GitHub repository by SSHing into the server and have  the server to redirect to my already existing domain. How would I go about doing this?
## [9][New Amazon Macie version coming soon?](https://www.reddit.com/r/aws/comments/g8z8fj/new_amazon_macie_version_coming_soon/)
- url: https://www.reddit.com/r/aws/comments/g8z8fj/new_amazon_macie_version_coming_soon/
---
I went to enable Macie in my account this morning and got a message suggesting I wait for the new, lower-cost, globally-available Macie to become available. Any idea when that might be happening?
## [10][Not able to populate AWS Glue Job monitoring metrics.](https://www.reddit.com/r/aws/comments/g8wbyx/not_able_to_populate_aws_glue_job_monitoring/)
- url: https://www.reddit.com/r/aws/comments/g8wbyx/not_able_to_populate_aws_glue_job_monitoring/
---
I am trying to populate maximum possible Glue job metrics for some testing, below is the setup I have created:

- A crawler reads data (dummy customer data of 500 rows) from a CSV file placed in an S3 bucket.
- Used another crawler to crawl tables created in Redshift cluster.
- An ETL job finally reads data from csv file in s3 and dumps it into a Redshift table.

The job is running without any issue and i am able to see final data getting dumped into Redshift table, however, in the end, only below 5 Cloudwatch metrics are being populated:

- glue.jvm.heap.usage
- glue.jvm.heap.used
- glue.s3.filesystem.read_bytes
- glue.s3.filesystem.write_bytes
- glue.system.cpuSystemLoad

There are approximately 20 more metrics which are not getting populated.

Any suggestions on how to populate those remaining metrics as well?
