# aws
## [1][re:Invent registration is now open](https://www.reddit.com/r/aws/comments/jkenu3/reinvent_registration_is_now_open/)
- url: https://register.virtual.awsevents.com/
---

## [2][Like I don't already spend enough hours a day with a headset on, Amazon are offering 3 whole weeks.](https://www.reddit.com/r/aws/comments/jkuid9/like_i_dont_already_spend_enough_hours_a_day_with/)
- url: https://www.reddit.com/r/aws/comments/jkuid9/like_i_dont_already_spend_enough_hours_a_day_with/
---
&gt;"AWS re:Invent is the world's largest, most comprehensive cloud computing event. This year, for the first time ever, [re:Invent is available as a free 3-week virtual event](https://email.awscloud.com/dc/c1xUhc3_T1MSZBn2C4tcV93yFLmM1X2-9yPbcR96HQKjALTzTGSheu047jEXJpODpP9v4g4vAecZhHKUX1Y1eBD271GwFaxuAj7910nLOCcd9D1upskWxTtCObEqfllgruhQrSgOfdI7Px0UcWMqdWvie9DE_Hvu-wjXG3U2_k6veOUsqNSEF7sRhqTxE4pfBimYPPm7WrzRib_dcvlq-6Xy4EuYZn_xTy0FF9vvgJMhP_IcZbQC3llL4iwNZ2eiMFejI97vp91ffNAl8C5-Zv6E8hl1tNDUP56O1gpwY1R9q1eWEnOD8L-Qu_G6f9gF7FwZ1g1d22mh8E_QqjDKaQ==/Emt0T912ZM80k0BM0a0BO01)."

&amp;#x200B;

I get fed up of headphones after a few hours on a call. **3 weeks!** I might have to change my audio settings to external speakers. (I am fortunate to be working from home.)
## [3][Elastic Load Balancing launches gRPC support for Application Load Balancer](https://www.reddit.com/r/aws/comments/jkkwi5/elastic_load_balancing_launches_grpc_support_for/)
- url: https://www.reddit.com/r/aws/comments/jkkwi5/elastic_load_balancing_launches_grpc_support_for/
---
Just like the title says, full support of gRPC as first class protocol. This includes unary, service-side streaming, client-side streaming, and bidirectional RPC. Your target group is gRPC type, and have gRPC health checks.

Checkout the blog post:

https://aws.amazon.com/blogs/aws/new-application-load-balancer-support-for-end-to-end-http-2-and-grpc/
## [4][New – Application Load Balancer Support for End-to-End HTTP/2 and gRPC](https://www.reddit.com/r/aws/comments/jkjv2c/new_application_load_balancer_support_for/)
- url: https://aws.amazon.com/blogs/aws/new-application-load-balancer-support-for-end-to-end-http-2-and-grpc/
---

## [5][No Secure Nitro Enclave for T3 instances](https://www.reddit.com/r/aws/comments/jkt7w9/no_secure_nitro_enclave_for_t3_instances/)
- url: https://www.reddit.com/r/aws/comments/jkt7w9/no_secure_nitro_enclave_for_t3_instances/
---
Why would this be so?
## [6][API GW api key per user - solution?](https://www.reddit.com/r/aws/comments/jkut5x/api_gw_api_key_per_user_solution/)
- url: https://www.reddit.com/r/aws/comments/jkut5x/api_gw_api_key_per_user_solution/
---
Hi,

I've used the API gateway a bunch for different use cases but i'm having trouble see how to use the API GW with different levels of users. Think of a website that provides an API, you have different tiers, sign up and you get your key. What does that look like in AWS?

I know you can have custom keys, is it a matter of something like lamdba, dynamo (track keys) to create, update keys in API GW? This seems like something that would be common, any links to solutions, examples would be good. I've tried searching, not having much luck.
## [7][Nitro Enclaves lets you use ACM certs directly on EC2 instances!](https://www.reddit.com/r/aws/comments/jkaee1/nitro_enclaves_lets_you_use_acm_certs_directly_on/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/10/announcing-aws-certificate-manager-for-nitro-enclaves/
---

## [8][Any current or former TAMs around?](https://www.reddit.com/r/aws/comments/jku3na/any_current_or_former_tams_around/)
- url: https://www.reddit.com/r/aws/comments/jku3na/any_current_or_former_tams_around/
---
Just have a few questions on the role!

Thanks
## [9][Did they JUST make a permanent change to the console UI? I was uploading items to my S3 bucket and I can't find a way to get back the view that it was before. Is there a way to revert back or did they finally get rid of that option? Not seeing any options to do so...](https://www.reddit.com/r/aws/comments/jk9pyi/did_they_just_make_a_permanent_change_to_the/)
- url: https://www.reddit.com/r/aws/comments/jk9pyi/did_they_just_make_a_permanent_change_to_the/
---
&amp;#x200B;

e.g. back when the buckets looked like this [https://miro.medium.com/max/1050/1\*xcOmWtHEhukc-VbgiLw8cg.png](https://miro.medium.com/max/1050/1*xcOmWtHEhukc-VbgiLw8cg.png)

I was looking for a quick way to verify MFA delete was enabled in the console. I know I can do this with the CLI.

Thanks!
## [10][Can you get cognito user access token in lambda without user password present?](https://www.reddit.com/r/aws/comments/jkq7an/can_you_get_cognito_user_access_token_in_lambda/)
- url: https://www.reddit.com/r/aws/comments/jkq7an/can_you_get_cognito_user_access_token_in_lambda/
---
I am working on my user update-email function and when they update their email in my front end app, I am sending them a custom email link to trigger an api gateway call that triggers a lambda function. I would like my user to be able to simply click this link and have their new email verified whether they are currently logged into a browser or not. 

To do so, I was going to use the verifyUserAttribute function from the CognitoIdentityServiceProvider. [https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/CognitoIdentityServiceProvider.html#verifyUserAttribute-property](https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/CognitoIdentityServiceProvider.html#verifyUserAttribute-property)

However, as you can see in the link, it requires a cognito user access token from a user session. Does anyone know a way to initiate a user session just based off username or email using admin privileges? I thought adminGetUser would work for second but upon examining what it returns, it only holds user attributes and not session information.
## [11][S3 Bucket for Book Covers?](https://www.reddit.com/r/aws/comments/jkpo9z/s3_bucket_for_book_covers/)
- url: https://www.reddit.com/r/aws/comments/jkpo9z/s3_bucket_for_book_covers/
---
Hello. I am a noob in the cloud world. I have a Node js app on Heroku and am trying to figure out where to host my images which are mostly book covers.

*Back in my day we had one website host and we liked it! \*shakes fist\**

I am a bit lost in the clouds (sorry). Should I be using an S3 bucket to host all of the images for my website? And if so do I need to set Access control list (ACL) to Everyone (Read/Read) even though it yells at me? (How else would people view the images?)

If everyone can see how lost I am, maybe there is a nooblier way to host images and static content. I doubt I'll be getting 100 billion downloads a month. Maybe just a few hundred to thousand, and probably mostly just North America.
