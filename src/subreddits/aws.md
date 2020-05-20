# aws
## [1][DynamoDB now supports empty values for non-key String and Binary attributes](https://www.reddit.com/r/aws/comments/gn5hhi/dynamodb_now_supports_empty_values_for_nonkey/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/05/amazon-dynamodb-now-supports-empty-values-for-non-key-string-and-binary-attributes-in-dynamodb-tables/
---

## [2][CORS No 'Access-Control-Allow-Origin' header Using AWS Lambda/S3/VS2019](https://www.reddit.com/r/aws/comments/gnav0b/cors_no_accesscontrolalloworigin_header_using_aws/)
- url: https://www.reddit.com/r/aws/comments/gnav0b/cors_no_accesscontrolalloworigin_header_using_aws/
---
&amp;#x200B;

 I have a Lambda server function uploaded to AWS through visual studio 2019 and enabled CORS through the api gateway and added CORS to the S3 bucket that hosted a static website. When I use curl command:

curl -H "Origin: [https://denisejames.dev/](https://denisejames.dev/)" -H "Access-Control-Request-Method: POST" -H "Access-Control-Request-Headers: X-Requested-With" -X OPTIONS --verbose [https://6j7r3vkufh.execute-api.us-east-1.amazonaws.com/Prod/api/dynamodb/getitems](https://6j7r3vkufh.execute-api.us-east-1.amazonaws.com/Prod/api/dynamodb/getitems)

&amp;#x200B;

    * schannel: stored credential handle in session cache
    &gt; OPTIONS /Prod/api/dynamodb/getitems HTTP/1.1
    &gt; Host: 6j7r3vkufh.execute-api.us-east-1.amazonaws.com
    &gt; User-Agent: curl/7.55.1
    &gt; Accept: */*
    &gt; Origin: https://denisejames.dev/
    &gt; Access-Control-Request-Method: POST
    &gt; Access-Control-Request-Headers: X-Requested-With
    &gt;
    * schannel: client wants to read 102400 bytes
    * schannel: encdata_buffer resized 103424
    * schannel: encrypted data buffer: offset 0 length 103424
    * schannel: encrypted data got 570
    * schannel: encrypted data buffer: offset 570 length 103424
    * schannel: decrypted data length: 541
    * schannel: decrypted data added: 541
    * schannel: decrypted data cached: offset 541 length 102400
    * schannel: encrypted data buffer: offset 0 length 103424
    * schannel: decrypted data buffer: offset 541 length 102400
    * schannel: schannel_recv cleanup
    * schannel: decrypted data returned 541
    * schannel: decrypted data buffer: offset 0 length 102400
    &lt; HTTP/1.1 200 OK
    &lt; Content-Type: application/json
    &lt; Content-Length: 3
    &lt; Connection: keep-alive
    &lt; Date: Wed, 20 May 2020 12:28:40 GMT
    &lt; x-amzn-RequestId: 9bcaeff3-1d2f-42db-826c-9a28ff3beeb8
    &lt; Access-Control-Allow-Origin: *
    &lt; Access-Control-Allow-Headers: *
    &lt; x-amz-apigw-id: M1JSyHxEIAMFing=
    &lt; Access-Control-Allow-Methods: OPTIONS,POST,GET,DELETE
    &lt; X-Cache: Miss from cloudfront
    &lt; Via: 1.1 aefca35da479fa09516461bbcf9ed01c.cloudfront.net (CloudFront)
    &lt; X-Amz-Cf-Pop: ORD51-C2
    &lt; X-Amz-Cf-Id: P-PYU57bCpHajQXPp40BLpLagy1snQ_4Z-bmY0O7XJkmEFjlohKWKQ==
    &lt;
    {}
    * Connection #0 to host 6j7r3vkufh.execute-api.us-east-1.amazonaws.com left intact

 

This is the error I get when I use use firefox at www.denisejames.dev

Cross-Origin Request Blocked: The Same Origin Policy disallows reading the remote resource at [https://6j7r3vkufh.execute-api.us-east-1.amazonaws.com/Prod/api/dynamodb/getitems](https://6j7r3vkufh.execute-api.us-east-1.amazonaws.com/Prod/api/dynamodb/getitems). (Reason: CORS header ‘Access-Control-Allow-Origin’ missing).
## [3][How to setup AWS Organizations with AWS SSO using G Suite as an identity provider. Made account management, centralized billing and resource sharing much easier in my own company. Hope this helps :) !](https://www.reddit.com/r/aws/comments/gmn67i/how_to_setup_aws_organizations_with_aws_sso_using/)
- url: https://medium.com/serverless-transformation/setup-aws-organizations-with-google-suite-saml-sso-7e676f5ed3e1
---

## [4][Roughly, how long does it take for a New Datacenter (Cape Town, is the one I’m interested in) to gain support for EKS?](https://www.reddit.com/r/aws/comments/gnaukl/roughly_how_long_does_it_take_for_a_new/)
- url: https://www.reddit.com/r/aws/comments/gnaukl/roughly_how_long_does_it_take_for_a_new/
---
This question is rather ambiguous, but just out of curiosity I was wondering how long it typically takes for a DataCenter to gain Containers EKS Support? Specifically the one in Cape Town has been under the Heading “We’re Working on it” for 28 Days now, but I can’t seem to find anything online that suggests how long it took previous data centres; so I don’t have any frame of reference for the time period of giving a Datacenter EKS support.
## [5][WAF vs NACL](https://www.reddit.com/r/aws/comments/gnaptg/waf_vs_nacl/)
- url: https://www.reddit.com/r/aws/comments/gnaptg/waf_vs_nacl/
---
Which is better option to restrict IP address. Can someone suggest which is better for application protection . I read that for public web application waf is a better choice and for private web app whitelisting add range in NAcl is good choice .
## [6][A beginner looking for Deployment Suggestion](https://www.reddit.com/r/aws/comments/gnaho4/a_beginner_looking_for_deployment_suggestion/)
- url: https://www.reddit.com/r/aws/comments/gnaho4/a_beginner_looking_for_deployment_suggestion/
---
Hello Everyone, Beginner here, me and my friend is working on a backend server for an app  
for which we need to setup a node js server &amp; would use MySQL for data, that would point to a subdomain also with some security and scalability , and stability.

Our current setup is we have deployed the node js server in ec2 and we are using it with a public ip.can you guys here suggest an ideal setup for our purpose. so that the up time is good and better scalability.

i have gone through some Reddit post in which some suggestions were using ALB, ACM etc. It would be very helpful if you guys can help me out here.
## [7][AWS Secrets Manager Alternative Without Limits: Vault From HashiCorp](https://www.reddit.com/r/aws/comments/gnaecf/aws_secrets_manager_alternative_without_limits/)
- url: http://selleoblogone.xyz
---

## [8][Client VPN to VPC with address range that overlaps with client's LAN](https://www.reddit.com/r/aws/comments/gn9pye/client_vpn_to_vpc_with_address_range_that/)
- url: https://www.reddit.com/r/aws/comments/gn9pye/client_vpn_to_vpc_with_address_range_that/
---
I'm trying to setup AWS Client VPN (with split tunnelling) to connect from my home network to a VPC, but I'm facing problems, which I think are related to overlapping IP address ranges.

The VPC's CIDR is [192.168.0.0/16](https://192.168.0.0/16). My home wifi network uses 192.168.0.x.

The Client VPN has a default route with the destination of [192.168.0.0/16](https://192.168.0.0/16) which I can't remove. The Client VPN's DNS server is 192.168.0.2.

I can connect to the VPN, but can't access any resources in the VPC. I don't think my traffic is getting blocked by security groups, ingress authorization or Network ACLs.

Am I wasting my time trying to get this working? Should it be possible or should I just make a new VPC with a different CIDR and use that instead?
## [9][What is the safest way to create Authenticator ?](https://www.reddit.com/r/aws/comments/gn8c5b/what_is_the_safest_way_to_create_authenticator/)
- url: https://www.reddit.com/r/aws/comments/gn8c5b/what_is_the_safest_way_to_create_authenticator/
---
I think it is best to have the same application on several devices so that if you lose your phone, you can still access other devices and log in.

Am I thinking right? - if yes What apps are suitable for this?

&amp;#x200B;

Or maybe there are other options to protect yourself from devices loss, theft, destruction, etc.
## [10][Heroku Django update static files EC2](https://www.reddit.com/r/aws/comments/gn258f/heroku_django_update_static_files_ec2/)
- url: https://www.reddit.com/r/aws/comments/gn258f/heroku_django_update_static_files_ec2/
---
So I am making a project where I serve the static css and images as well as upload pdfs to S3. Is there a way to update the static files if they are updated locally so I dont have to upload the css files through S3 console every time I make a change? I know that there is a S3 commandline
