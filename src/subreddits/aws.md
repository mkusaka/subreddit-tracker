# aws
## [1][Week of Aug 17th - What are your favorite AWS Tips?](https://www.reddit.com/r/aws/comments/ibdue5/week_of_aug_17th_what_are_your_favorite_aws_tips/)
- url: https://www.reddit.com/r/aws/comments/ibdue5/week_of_aug_17th_what_are_your_favorite_aws_tips/
---
Share your best AWS tips with the community!
## [2][End to End Machine Learning Tutorial — From Data Collection to Deployment on AWS 🚀](https://www.reddit.com/r/aws/comments/ickq5y/end_to_end_machine_learning_tutorial_from_data/)
- url: https://www.reddit.com/r/aws/comments/ickq5y/end_to_end_machine_learning_tutorial_from_data/
---
\[Updated Post: End To End machine learning and beyond notebooks 🚀\]

Hello everyone!

With the great collaboration of a friend of mine, we built and deployed a machine learning application to AWS using Python.

To tell you more about this fun journey, we wrote a post where we'll walk you through the steps to:

* Collect and scrape data with Scrapy / Selenium
* Train a deep character CNN for (English) sentiment analysis using PyTorch
* Build an interactive web app with Dash to serve the model in real-time
* Put everything in Docker Compose
* Deploy to AWS on a custom domain name

Here the link: [https://medium.com/datadriveninvestor/end-to-end-machine-learning-from-data-collection-to-deployment-ce74f51ca203](https://medium.com/datadriveninvestor/end-to-end-machine-learning-from-data-collection-to-deployment-ce74f51ca203)

All the code is available in Github, the blocks are independent and reusable in other projects [https://github.com/MarwanDebbiche/post-tuto-deployment](https://github.com/MarwanDebbiche/post-tuto-deployment)

Feel free to reach out to me if you have any question,
## [3][AWS SSO with GSuite as IDp - Login fails from GSuite Dashboard, OK from AWS SSO Start](https://www.reddit.com/r/aws/comments/icga2m/aws_sso_with_gsuite_as_idp_login_fails_from/)
- url: https://www.reddit.com/r/aws/comments/icga2m/aws_sso_with_gsuite_as_idp_login_fails_from/
---
EDIT:  Found the fix. 

Found [this post](https://old.reddit.com/r/aws/comments/eadzp8/a_little_help_needed_with_aws_sso_with_gsuite/fckckpx/) after looking on the [AWS Support forum thread](https://forums.aws.amazon.com/thread.jspa?threadID=323869&amp;tstart=0)

A) Setting the Start URL in GSuite to blank, and 
B) Setting Name ID Format = EMAIL 

then *waiting a few minutes* updates the button in Gsuite Dashboard. Now it works! 



---- 

Original Post: 

There's a [post a few months ago here](https://www.reddit.com/r/aws/comments/gmn67i/how_to_setup_aws_organizations_with_aws_sso_using/) which goes into detail on how to set up AWS SSO using GSuite as the IDp. 

This is almost identical to the [post on Amazon's Security Blog from a month ago on the same process](https://aws.amazon.com/blogs/security/how-to-use-g-suite-as-external-identity-provider-aws-sso/) - the only differences are /u/bartheletf's post has leaving Signed Response in the GSuite Dashboard unchecked, whereas Amazon have it checked, and Amazon goes on to mention using their [ssosync](https://github.com/awslabs/ssosync) project to provision users on a schedule. 

Both guides work great in terms of getting logins via the AWS SSO Start page working. 

However, and as mentioned by others in the other reddit post - trying to log into AWS SSO by launching it from the GSuite Dashboard doesn't work, and just gives you a SAML Error over on Amazon's side. [Screenshot here](https://i.imgur.com/Sd1e3jI.png).  Clicking the Sign In button (only appears if you're not already signed in) makes it work same as if you'd gone to the AWS SSO Start page and signed in yourself. 

I've tried with the Signed Response checkbox both enabled or not, but it doesn't change anything. 

I'm no SAML expert, but it seems like GSuite isn't doing it's own SAML call correct

HTTP Calls when using the AWS SSO Start page:   

- `POST https://accounts.google.com/o/saml2/idp?idpid=&lt;idpid&gt;` returns HTTP 200  -- contains a SAML AuthnRequest payload on the request   
- `POST https://us-east-1.signin.aws.amazon.com/platform/saml/acs/&lt;acsid&gt;`  returns HTTP 200 -- contains a SAML Response payload on the request

HTTP Calls when using GSuite Dashboard: 

- `GET https://accounts.google.com/o/saml2/initsso?idpid=&lt;idpid&gt;&amp;spid=&lt;spid&gt;&amp;forceauthn=false` returns HTTP 200
- `POST https://us-east-1.signin.aws.amazon.com/platform/saml/acs/&lt;acsid&gt;` returns HTTP 400 -- contains a SAML Response payload on the request

When comparing the content of the SAML being posted to the Amazon ACS URLs, the first difference I noticed was that the NameID format beign sent when launched from GSuite is "urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified", but when from the AWS SSO it was "urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress" 
Changing the SSO Application in GSuite from UNSPECIFIED to EMAIL fixed that so that both were using the "urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress" format. 

Another difference is the Response when launched from AWS SSO has InResponseTo attribute in two elements, referencing the ID send in the AuthRequest payload. 

Other than that, the rest of it is differences in Timestamps (both look to be correct time), as well as various session ID and signatures/digests. 

Has anyone else had any luck getting this to work?
## [4][Cognito User Pool Backup &amp; Restore?](https://www.reddit.com/r/aws/comments/iclxfo/cognito_user_pool_backup_restore/)
- url: https://www.reddit.com/r/aws/comments/iclxfo/cognito_user_pool_backup_restore/
---
Is this still not a thing? How are people doing DR across regions with Cognito? I don't see a way to export user pool users with their password hashes. In case of a DR event where I need to fail over to another region, I don't see any way around having to send every user a password reset.
## [5][AWS announces WorldForge in AWS RoboMaker](https://www.reddit.com/r/aws/comments/ic27tx/aws_announces_worldforge_in_aws_robomaker/)
- url: https://aws.amazon.com/blogs/aws/aws-announces-worldforge-in-aws-robomaker/
---

## [6][How I Write Meaningful Tests for AWS Lambda Functions](https://www.reddit.com/r/aws/comments/icmjue/how_i_write_meaningful_tests_for_aws_lambda/)
- url: https://medium.com/@psingman/how-i-write-meaningful-tests-for-aws-lambda-functions-f009f0a9c587
---

## [7][Get ENI IP address after creating VPC Endpoint](https://www.reddit.com/r/aws/comments/icmgxd/get_eni_ip_address_after_creating_vpc_endpoint/)
- url: https://www.reddit.com/r/aws/comments/icmgxd/get_eni_ip_address_after_creating_vpc_endpoint/
---
Hello,  


I need some help figuring this out. Using Cloudformation I am creating a VPC endpoint and would like to get the IP addresses of the created ENIs.  


I have read through the docs and there is no direct way to do this.   


Any suggestions?
## [8][Using Current_date in Redshuft materialised view](https://www.reddit.com/r/aws/comments/icig9n/using_current_date_in_redshuft_materialised_view/)
- url: https://www.reddit.com/r/aws/comments/icig9n/using_current_date_in_redshuft_materialised_view/
---
Hi All,

I found in documentation that current\_date function can't be used in materialised views, is there an alternative to this? I need current date in a filter condition.

Thnx,

mc
## [9][Does SageMaker GroundTruth support cases in which the input S3 bucket for the dataset has new data added to it once a labeling jobs starts?](https://www.reddit.com/r/aws/comments/ickis1/does_sagemaker_groundtruth_support_cases_in_which/)
- url: https://www.reddit.com/r/aws/comments/ickis1/does_sagemaker_groundtruth_support_cases_in_which/
---
I have a service where my team can submit their own images to the s3 bucket to be included in the labeling process, and would like to incorporate that new data in the current and future labeling jobs. 

From what I can tell, GroundTruth relies on the manifest file (which can be autogenerated) to know which images/resources need labeling. However when creating a job, it only takes into account the current state of the input S3 bucket. 

I don't think having a lambda trigger on S3 upload to manually update the manifest file is a good idea or event something that works. 

Anyone encounter a similar situation? How did you solve it? 
Cannot seem to find a reference online in the documentation/youtube videos about handling that.
## [10][Advice needed: Setting up CI/CD for serverless?](https://www.reddit.com/r/aws/comments/ici0eb/advice_needed_setting_up_cicd_for_serverless/)
- url: https://www.reddit.com/r/aws/comments/ici0eb/advice_needed_setting_up_cicd_for_serverless/
---
I've recently taken over a large codebase using serverless framework. The application is broken into several 'services' and uses a single `serverless.yml` file in the root of the repository. Each 'service' contains several functions with API Gateway as the event source. The `serverless.yml` file is highly dynamic, taking advantages of variables in several places. Services are deployed by running `serverless deploy --stage dev --service profile` where `profile` is the name of the Profile service. The command bundles and deploys a single service thus each service is it's own CloudFormation stack.

I'm the 3rd dev to take over this project. It seems like the original dev who setup the project planned on following a microservices pattern however the 2nd dev, who preceded me, has made a bit of a mess of it by importing code across the service boundaries and having code in one service importing code from another service. I'm in the process of fixing this mess by decoupling the services by using messaging (via custom EventBridge events). Despite the 2nd dev making a mess I think the intent of the original dev is pretty good.

I'd like to setup CI/CD for this project using the AWS Code\* suite of tools but given how this repository is structured I can't think of the best approach. For example if a change is commited to one service is doesn't make sense to build and deploy **all** of the other services since they haven't changed.

What's the best approach for this kind of setup? Would it be better to move each 'service' into it's own repository and give each service it's own CI/CD pipeline? Any tips would be appreciated.
## [11][Connect QuickSight to Aurora Serverless MySQL compatibility](https://www.reddit.com/r/aws/comments/ichdd4/connect_quicksight_to_aurora_serverless_mysql/)
- url: https://www.reddit.com/r/aws/comments/ichdd4/connect_quicksight_to_aurora_serverless_mysql/
---
I trying to connect AWS QuickSight to my aurora serverless database but always get timeout error

&gt;Connecting to your data source took too long. Retrying this request may help. Contact technical support for further assistance.

I have created a VPC connection and open all ports in the security group but still cannot connect.

I am missing anything?
