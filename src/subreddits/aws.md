# aws
## [1][The Complete CloudFormation Guide](https://www.reddit.com/r/aws/comments/fbi106/the_complete_cloudformation_guide/)
- url: https://start.jcolemorrison.com/the-complete-cloudformation-guide/
---

## [2][SQS and Lambda with Async calls](https://www.reddit.com/r/aws/comments/fbtlyd/sqs_and_lambda_with_async_calls/)
- url: https://www.reddit.com/r/aws/comments/fbtlyd/sqs_and_lambda_with_async_calls/
---
Hello all, 

I have a question regarding the fact that you can call Lambda as an async call and then it has by default retry option for async calls. 

SQS has this too, as I noticed from a basic test that SQS retries for 4 times if I'm not mistaken, and Lambda has a default value of 2 retires. So, I'm getting 6 calls in case of failure in Lambda because of the retires default options in both of them?? 

Did I missed any info here and that makes me think that this is the case?

&amp;#x200B;

Thanks..
## [3][Everywhere I read about oatuh 2 tokens, they say you should not store them on client side/localstorage, but AWS amplify does store them on the client's local storage, can somebody more in the know explain why it is ok for Amplify to do so?](https://www.reddit.com/r/aws/comments/fbnalo/everywhere_i_read_about_oatuh_2_tokens_they_say/)
- url: https://www.reddit.com/r/aws/comments/fbnalo/everywhere_i_read_about_oatuh_2_tokens_they_say/
---
For reference:

  
Auth0 specifying you should store them on localstorage/client: [https://auth0.com/docs/tokens/guides/store-tokens](https://auth0.com/docs/tokens/guides/store-tokens)  


Here is docs specifying amplify stores them on the client: [https://aws-amplify.github.io/docs/js/authentication#managing-security-tokens](https://aws-amplify.github.io/docs/js/authentication#managing-security-tokens)  


How exactly are you supposed to store these tokens using SPA apps which use S3 as a backend if its a security issue to store them on the client. Does that mean that every single SPA app which remembers user's across sessions with federated login is actually going against security best practices, and hence is  vulnerable?
## [4][SES SMTP limitations with printers, etc.](https://www.reddit.com/r/aws/comments/fbrshf/ses_smtp_limitations_with_printers_etc/)
- url: https://www.reddit.com/r/aws/comments/fbrshf/ses_smtp_limitations_with_printers_etc/
---
I try to use SES as our outbound SMTP infrastructure for incidental devices (printers, etc...) because I like being able to restrict to specific approved sender email addresses via IAM policies, and I have a mature Terraform setup in place for managing accounts/policies.

HOWEVER

I keep running into devices that for some reason just... don't work with SES, but work with any other email provider I try.  I've encountered this on a Brother printer in the past, and an HP Color LaserJet Pro MFP M479fdw right now on my desk.

Behavior: HP printer simply says "System failure" when attempting to send a test email when configured using SES.

Symptoms / troubleshooting so far:

* Tried changing 1 character in username and password to rule out bad credentials or a bad webUI silently truncating long passwords.  Error message changed to "Invalid credentials", so I know the credentials are being stored correctly; it's something else.
* Tried the same SES SMTP credentials in other SMTP applications, and they do work, so the credentials are good.
* Tried switching to Mailgun on the printer, email started working, so scan-to-email \*does\* work in some capacity.  Also Mailgun also has really long passwords, so it's not a password length thing.  Also tested with Fastmail and scan-to-email works with them as well.
* Tried to find any sort of diagnostic/debugging logs anywhere on the printer -- no luck.  I hooked it up to external syslog at the highest debugging level, but the printer doesn't log anything about SMTP unfortunately.
* Upgraded printer firmware to the latest published version, no change.
* Tried all combinations of TLS / non-TLS, all SES ports, direct IP address vs. DNS entry for SMTP server hostname.
* Extensive Googling, no luck finding anyone in a similar predicament.

So -- I'm asking two things:

1. Has anyone fought this specific problem and has tips for what I should try?  I wish I could get diagnostic logs out of the printer, but no luck.
2. At a higher level, does anyone know what specific things SES does differently that would cause it to fail when used in conjunction with brittle SMTP clients in embedded devices like printers, etc?

This is low-stakes -- Mailgun is working fine, I just want to know what's wrong and simplify the stack to use AWS if at all possible.  Thanks!
## [5][Aws single VM monitoring](https://www.reddit.com/r/aws/comments/fbr0f6/aws_single_vm_monitoring/)
- url: https://www.reddit.com/r/aws/comments/fbr0f6/aws_single_vm_monitoring/
---
In short: is there a quick and dirty way to generate a weekly (CentOs)  VM performance report(CPU/MEM)? I've got CloudWatch with detailed monitoring enabled. Or is the standard to simply build a dashboard and screenshot it?

Background: I come from a sysadmin background and run a sizeable GCP environment. I'm now dipping into AWS as it seems to be the way to go.
## [6][Terraform vs cloud formation?](https://www.reddit.com/r/aws/comments/fbdmjz/terraform_vs_cloud_formation/)
- url: https://www.reddit.com/r/aws/comments/fbdmjz/terraform_vs_cloud_formation/
---
Is it worth going over cloud formation if terraform already goes the job ? Is there something that cloud formation does better than terraform or chef,etc that one should be aware of?
## [7][Trouble curl-ing into API Gateway to Lambda](https://www.reddit.com/r/aws/comments/fbpm3m/trouble_curling_into_api_gateway_to_lambda/)
- url: https://www.reddit.com/r/aws/comments/fbpm3m/trouble_curling_into_api_gateway_to_lambda/
---
I'll preface this by saying I'm not an expert in this area and it might be a problem with how I'm using curl but: I've created an api gateway to a lambda function and I'm trying to test it using curl. I have successfully tested the lambda function.

The lambda event uses an event:

{"phone":"phone number","message":"message to send"}

I created the api using HTTP API and auto deploy. I noticed that it did not deploy initially as format 2.0 seems to have problems. I had since then changed it to format 1.0, which did deploy.

I attempted to curl it using:

curl -v -X POST [https://myapiid.execute-api.ap-southeast-2.amazonaws.com/prod](https://dj9u44aoze.execute-api.ap-southeast-2.amazonaws.com/prod) \-d '{"phone":"+myphone","message":"curl lambda test"}'

The results I get are:

Note: Unnecessary use of -X or --request, POST is already inferred.

\*   Trying [00.00.0](https://54.79.66.233)0.00...

\* TCP\_NODELAY set

\* Connected to [myapiid.execute-api.ap-southeast-2.amazonaws.com](https://dj9u44aoze.execute-api.ap-southeast-2.amazonaws.com) ([00.00.0](https://54.79.66.233)0.00) port 443 (#0)

\* ALPN, offering h2

\* ALPN, offering http/1.1

\* Cipher selection: ALL:!EXPORT:!EXPORT40:!EXPORT56:!aNULL:!LOW:!RC4:@STRENGTH

\* successfully set certificate verify locations:

\*   CAfile: /etc/ssl/cert.pem

CApath: none

\* TLSv1.2 (OUT), TLS handshake, Client hello (1):

\* TLSv1.2 (IN), TLS handshake, Server hello (2):

\* TLSv1.2 (IN), TLS handshake, Certificate (11):

\* TLSv1.2 (IN), TLS handshake, Server key exchange (12):

\* TLSv1.2 (IN), TLS handshake, Server finished (14):

\* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):

\* TLSv1.2 (OUT), TLS change cipher, Client hello (1):

\* TLSv1.2 (OUT), TLS handshake, Finished (20):

\* TLSv1.2 (IN), TLS change cipher, Client hello (1):

\* TLSv1.2 (IN), TLS handshake, Finished (20):

\* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256

\* ALPN, server accepted to use h2

\* Server certificate:

\*  subject: CN=\*.execute-api.ap-southeast-2.amazonaws.com

\*  start date: Nov 20 00:00:00 2019 GMT

\*  expire date: Dec 20 12:00:00 2020 GMT

\*  subjectAltName: host "[myapiid.execute-api.ap-southeast-2.amazonaws.com](https://dj9u44aoze.execute-api.ap-southeast-2.amazonaws.com)" matched cert's "\*.execute-api.ap-southeast-2.amazonaws.com"

\*  issuer: C=US; O=Amazon; OU=Server CA 1B; CN=Amazon

\*  SSL certificate verify ok.

\* Using HTTP2, server supports multi-use

\* Connection state changed (HTTP/2 confirmed)

\* Copying HTTP/2 data in stream buffer to connection buffer after upgrade: len=0

\* Using Stream ID: 1 (easy handle 0x7fae9980ec00)

\&gt; POST /prod HTTP/2

\&gt; Host: [myapiid.execute-api.ap-southeast-2.amazonaws.com](https://dj9u44aoze.execute-api.ap-southeast-2.amazonaws.com)

\&gt; User-Agent: curl/7.54.0

\&gt; Accept: \*/\*

\&gt; Content-Length: 53

\&gt; Content-Type: application/x-www-form-urlencoded

\&gt;

\* Connection state changed (MAX\_CONCURRENT\_STREAMS updated)!

\* We are completely uploaded and fine

&lt; HTTP/2 404

&lt; date: Sun, 01 Mar 2020 05:22:43 GMT

&lt; content-length: 23

&lt; x-amzn-requestid: reqidgiven=

&lt;

\* Connection #0 to host [myapiid.execute-api.ap-southeast-2.amazonaws.com](https://dj9u44aoze.execute-api.ap-southeast-2.amazonaws.com) left intact

{"message":"Not Found"}

I get the same {"message":"Not Found"} outcome when I click on the invocation link, I'm not sure which part I'm getting wrong
## [8][boto3 dynamo](https://www.reddit.com/r/aws/comments/fbmfol/boto3_dynamo/)
- url: https://www.reddit.com/r/aws/comments/fbmfol/boto3_dynamo/
---
Hello,

I'm using boto3 to add items to a dynamoDB table. It works fine but I can't see the added items in the AWS console. I can't even see the created tables. Am I doing something wrong?
## [9][Identify Windows disks to EBS volumes](https://www.reddit.com/r/aws/comments/fbmaaz/identify_windows_disks_to_ebs_volumes/)
- url: https://www.reddit.com/r/aws/comments/fbmaaz/identify_windows_disks_to_ebs_volumes/
---
Hi, How do we know which EBS volume is which drive in Windows OS?

I need to expand D:/ and E:/ drives on my EC2 instance and could not find a way to relate the drives in AWS Console.

Thanks in advance
## [10][Attempting to create a MediaLive/MediaPackage config for streaming](https://www.reddit.com/r/aws/comments/fblfmu/attempting_to_create_a_medialivemediapackage/)
- url: https://www.reddit.com/r/aws/comments/fblfmu/attempting_to_create_a_medialivemediapackage/
---
I'm attempting to create a demo service where one can livestream video/audio to an audience.

via the SDKs, I am spinning up new/unique MediaLive + MediaPackage channel &amp; inputs per created stream, as well as a Cloudfront distribution pointing to my MediaPackage for cdn, however CF distros can take up to 20 minutes to propagate which is obviously not optimal for on demand stream creation. How I should approach resolving this?
