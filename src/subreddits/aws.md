# aws
## [1][Design help question - Trying to set up a raspberry pi + moisture sensor to trigger a email/alarm, want to use AWS. Should I do IoT or API Gateway + Lambda? Or something else?](https://www.reddit.com/r/aws/comments/eudpcb/design_help_question_trying_to_set_up_a_raspberry/)
- url: https://www.reddit.com/r/aws/comments/eudpcb/design_help_question_trying_to_set_up_a_raspberry/
---
Hello, I have a section in my basement that gets flooded often, I decided to buy a moisture sensor that was pretty cheap, and use a raspberry pi to connect to the sensor. The idea is that when the water starts to back up, it immediately sends me a text or email or something and I can shut the water in my house and fix the pipe...

Anyways, I've been using AWS a lot at work, and would like to continue to learn more about it. My first idea was to set up the pi so that when it detects water, it POSTs to an API gateway I set up, which runs a lambda, which then sends a text or email or something. I have worked with a lot of API gateway/lambda stuff in the past.

However I read about AWS IoT recently. I read this could be a good thing to use for sensors/pi's and stuff. I have never done anything IoT related, so it could be a cool learning oportunity.

For the above, is IoT a good solution? Is it better to stick with API Gateway? Is it stupid and complicated to use aws in the first place? Any other suggestions or recommendations before I dive deep into making this?

Thanks!
## [2][How to deploy artifacts of a particular version together?](https://www.reddit.com/r/aws/comments/eunnua/how_to_deploy_artifacts_of_a_particular_version/)
- url: https://www.reddit.com/r/aws/comments/eunnua/how_to_deploy_artifacts_of_a_particular_version/
---
Hello!

Imagine there is an application which consists of two services which run on two different EC2 instances. Both services must be compatible with each other. Different versions of service source code can be retrieved using version tags in git.

I want to create some automatic mechanism which

1. retrieves a particular version of service 1,
2. retrieves a compatible version of service 2, and
3. deploys both services on different EC2 instances.

In order for this to work, I must store somewhere the information which services belong together, e. g. in a table with columns

a. Configuration version
b. Version of service 1
c. Version of service 2

Most likely, I am not the first person to have this problem. Therefore there is probably an existing solution for this.

What piece of software can I use instead of that table (i. e. to keep track which versions of service 1 and 2 are compatible with each other)?

Options I have looked at:

1. Create a repository in the VCS which contains a flat file with versions of service 1 and 2. When a deployment is started, the CI/CD system (e. g. Jenkins) gets a version hash as input, then reads the versions from the git repository.
2. Someone recommended me to look at [Service Registration and Discover](https://spring.io/guides/gs/service-registration-and-discovery/), but I have no idea how to apply it to my problem.
3. Spinnaker seems to be able to do it. But I am not sure that it is a good idea to use it only because of that functionality.
4. Helm charts also seem to do just that, but they are made for Kubernetes. At the moment we don't use it and I am not sure that we need it. 

Thanks in advance

Dmitrii
## [3][Lambda to talk with Local Resources (Label Printer)](https://www.reddit.com/r/aws/comments/eunn6h/lambda_to_talk_with_local_resources_label_printer/)
- url: https://www.reddit.com/r/aws/comments/eunn6h/lambda_to_talk_with_local_resources_label_printer/
---
Hi all,

In my current company we use a label printer to get stories from our sprints into a physical board so everyone else can have some visibility (it's also by the board we do our stand ups).

Currently we have windows service app running locally in the devs laptops, obviously it's not a good setup as we often forget to start the service but I don't want to deploy the win service in a server (even tough we have several available)

This looks to me like the perfect thing to be a lambda, we already use SQS to submit print requests (user types the Jira Id in slack, a slack app adds the request into SQS), so basically I want to know if Lambda is able to access resources in our local network (mainly The Printer).

I know Lambda can access resources in our VPC but obviously the printer is not there.
Does this requires a connection between our local network and AWS? I'm fairly sure we don't have a DirectConnect Connection, although we might have a VPN setup in place (we have local servers that talk to aws instances).

Any thoughts?
## [4][Is it possible to do web automations like Selenium on AWS Lambda?](https://www.reddit.com/r/aws/comments/eunlh5/is_it_possible_to_do_web_automations_like/)
- url: https://www.reddit.com/r/aws/comments/eunlh5/is_it_possible_to_do_web_automations_like/
---
I need to do web automations for a site with no API or web hooks. I can't use selenium because it happens at arbitrary times so I need something on the cloud.
## [5][How can I launch ubuntu GUI on ec2 instance?](https://www.reddit.com/r/aws/comments/eulj1o/how_can_i_launch_ubuntu_gui_on_ec2_instance/)
- url: https://www.reddit.com/r/aws/comments/eulj1o/how_can_i_launch_ubuntu_gui_on_ec2_instance/
---
I know how to create an ec2 instance and connect to an ubuntu server via SSH connection on my mac terminal. 

But how can I use the ubuntu GUI? 

Do I just install the desktop like how i would normally do? 

Doesn't seem right tho
## [6][Cognito: Is there a way to verify JWT tokens at front-end using aws-sdk?](https://www.reddit.com/r/aws/comments/eu4ari/cognito_is_there_a_way_to_verify_jwt_tokens_at/)
- url: https://www.reddit.com/r/aws/comments/eu4ari/cognito_is_there_a_way_to_verify_jwt_tokens_at/
---
I'm using aws-sdk at front-end of my web application. Till now, I've set-up the flow to register new users, authenticate users that will get the access token, id token, and refresh token. 

I'm confused about what's next !!!

* The access and id tokens are valid for 1 hour and refresh token for 30days, and all are in JWT format.
* So, in order to check the log-in status of the user,  the access token needs to be parsed to check for the expiration time. Is it possible to do this at front end? 
* Consequently, if expired then using the refresh token will provide fresh access and id tokens.

Basically, I want to check the validity of the tokens and expiration time to maintain user log-in status.
## [7][Can I redirect my Elastic IP to https instead if http](https://www.reddit.com/r/aws/comments/eujkfa/can_i_redirect_my_elastic_ip_to_https_instead_if/)
- url: https://www.reddit.com/r/aws/comments/eujkfa/can_i_redirect_my_elastic_ip_to_https_instead_if/
---
If yes, how to do so. I tried Self Signed certificate, I works but webpage first pops up with a page saying Connection is not secure on this page. I don’t wan this.
## [8][Where is the AWS SSO API?](https://www.reddit.com/r/aws/comments/eu5xsb/where_is_the_aws_sso_api/)
- url: https://www.reddit.com/r/aws/comments/eu5xsb/where_is_the_aws_sso_api/
---
Hi All,

I've now implemented the SAML integration via the old method [https://docs.microsoft.com/en-us/azure/active-directory/saas-apps/amazon-web-service-tutorial](https://docs.microsoft.com/en-us/azure/active-directory/saas-apps/amazon-web-service-tutorial) and the new [https://aws.amazon.com/blogs/aws/the-next-evolution-in-aws-single-sign-on/](https://aws.amazon.com/blogs/aws/the-next-evolution-in-aws-single-sign-on/)

For what its worth the old method is easier to automate as there is terraform support for all things IAM. The new method is a much sleeker experience for the end user, but I can find zero information on how to automate it? How do you automate mapping users/groups to AWS accounts? How do you manage permissions sets via code?

The FAQ implies that there isn't one [https://aws.amazon.com/single-sign-on/faqs/](https://aws.amazon.com/single-sign-on/faqs/) 

&gt;**Is there an API available for AWS SSO?**  
&gt;  
&gt;No. You can use the AWS SSO console to perform all necessary operations.

Interesting tidbit [https://github.com/sportradar/aws-azure-login](https://github.com/sportradar/aws-azure-login) is required with the old method, while the AWS SSO will work with the new version of the AWS CLI v2 [https://aws.amazon.com/blogs/developer/aws-cli-v2-now-supports-aws-single-sign-on/](https://aws.amazon.com/blogs/developer/aws-cli-v2-now-supports-aws-single-sign-on/)
## [9][New to S3 Glacier -- How to send REST commands?](https://www.reddit.com/r/aws/comments/eu8o1a/new_to_s3_glacier_how_to_send_rest_commands/)
- url: https://www.reddit.com/r/aws/comments/eu8o1a/new_to_s3_glacier_how_to_send_rest_commands/
---
Hi there,

Over the past few days I've been toying with the idea of backing up my photos (about 1TB+ in RAW) and videos to a cloud storage solution... Currently I have them backed up to my NAS but it is not ideal in case of a disaster or fire or anything.

I looked at Google, Azure and Amazon and I think Glacier and Deep Archive seems to have the best balance of storage and retrieval cost for me. B2 is out of the question for me because the upload speeds are really slow over in Singapore (3MB/s vs 100MB/s on the 3 main providers. I have a 1gbps fibre connection.)

Before I proceed further, I wanted to make sure that data from a Glacier vault is easy to retrieve without any reliance on third party software, if possible. My parents used to keep lots of home videos of us when we were younger but unfortunately they are not readable anymore because they used the long obsolete Video8 format. So I am kinda wary about using CloudBerry or Duplicati etc in case it doesn't exist ten or twenty years down the road. I am thinking of simply zipping up files in logical folders sorted by years or something and then upload them to the cloud.

So it seems there are two main ways to get things into Glacier/ Deep Archive -- one is by life cycling it from an S3 bucket, and the other method is to upload and download directly through AWSCLI or REST API. The first method seems easy enough, but I am also determined to master the second method 'just in case'. And the second method is where I'm stuck right now. 

I saw that Amazon's documentation has quite a number of useful examples using the REST API (e.g. at [https://docs.aws.amazon.com/amazonglacier/latest/dev/api-initiate-job-post.html](https://docs.aws.amazon.com/amazonglacier/latest/dev/api-initiate-job-post.html)). One of their examples show this:

&amp;#x200B;

    POST /-/vaults/examplevault/jobs HTTP/1.1
    Host: glacier.us-west-2.amazonaws.com
    x-amz-Date: 20170210T120000Z
    x-amz-glacier-version: 2012-06-01
    Authorization: AWS4-HMAC-SHA256 Credential=AKIAIOSFODNN7EXAMPLE/20141123/us-west-2/glacier/aws4_request,SignedHeaders=host;x-amz-date;x-amz-glacier-version,Signature=9257c16da6b25a715ce900a5b45b03da0447acf430195dcb540091b12966f2a2
    
    {
      "Type": "archive-retrieval",
      "ArchiveId": "NkbByEejwEggmBz2fTHgJrg0XBoDfjP4q6iu87-TjhqG6eGoOY9Z8i1_AUyUsuhPAdTqLHy8pTl5nfCFJmDl2yEZONi5L26Omw12vcs01MNGntHEQL8MBfGlqrEXAMPLEArchiveId"
      "Description": "My archive description",
      "SNSTopic": "arn:aws:sns:us-west-2:111111111111:Glacier-ArchiveRetrieval-topic-Example",
      "Tier" : "Bulk"
    }

Now, this looks easy enough for me to modify to my needs. But the (stupid?) question that I have is: Where on earth do I key this in? I tried the API gateway but there doesn't seem to be a way for me to paste a whole chunk of text directly (I have to manually select GET, POST, etc...). I've Googled quite a bit and it came up with things like Docker, AWS Lambda, etc. but I have no idea what the correct tool is to use?

Can I even run these commands in AWSCLI? If so, how do I even input multiple lines?

Any help to point me in the right direction would be greatly appreciated!
## [10][Anyone have a good example (github repo ideally) of a Python backend using Cognito user pools and flows?](https://www.reddit.com/r/aws/comments/eu8hsi/anyone_have_a_good_example_github_repo_ideally_of/)
- url: https://www.reddit.com/r/aws/comments/eu8hsi/anyone_have_a_good_example_github_repo_ideally_of/
---
I've been fiddling with Cognito off and on for a bit, it's a beast although I feel like the primary issue is the lack of good documentation for it.

I'd like to set up a simple auth system with Cognito as the identity provider and some basic sign up, email verification and password update flows.

Any help is appreciated.
