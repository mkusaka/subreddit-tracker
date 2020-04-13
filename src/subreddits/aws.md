# aws
## [1][Prospective user wanting to get in to AWS](https://www.reddit.com/r/aws/comments/g0igj4/prospective_user_wanting_to_get_in_to_aws/)
- url: https://www.reddit.com/r/aws/comments/g0igj4/prospective_user_wanting_to_get_in_to_aws/
---
I was wanting to run a game server from an actual server, like AWS does...Or at least that's what I heard they do.

I have no idea how I would even begin to do this. I mean, I guess my ideal case would be if it were basically a virtual machine that would run the Game Server application.

Is this a feasible possibility? How can I get started for learning how to do this?
## [2][Customer support nightmare](https://www.reddit.com/r/aws/comments/g05rs9/customer_support_nightmare/)
- url: https://www.reddit.com/r/aws/comments/g05rs9/customer_support_nightmare/
---
I'm trying to resolve an issue with my AWS account at the moment. I have not logged into the account in quite some time, probably 5+ years and attempted to a few days ago. When trying to log in, I was told that there wasn't an account associated with the email I tried and then when attempting to sign up with the same email, I'm told that the email is currently in use. I have brought up this issue to AWS support, but strangely enough, their support staff keep instructing me to get further support through [amazon.com](https://amazon.com) channels. The [amazon.com](https://amazon.com) support naturally just want to help me with any issues that might be going on with my [amazon.com](https://amazon.com) account, which uses a completely different email address and really as far as I can tell, has nothing to do with my AWS account. No matter how many times I tell them that they're suggestion wasn't helpful and what the [amazon.com](https://amazon.com) support is trying to do, they keep trying to direct me to that support channel. Not sure what to do at this point as I am about 10 emails deep and I have repeated steps, such as sending screenshots and doing as they ask multiple times now. Any help in this matter would be greatly appreciated, thanks!
## [3][EC2 rebooted without warning](https://www.reddit.com/r/aws/comments/g0iyzc/ec2_rebooted_without_warning/)
- url: https://www.reddit.com/r/aws/comments/g0iyzc/ec2_rebooted_without_warning/
---
Hi! I have some applications deployed to a pair of EC2 instances behind an ALB. Today I woke up to find a lot Slack messages saying something was broken. Upon closer inspection, it turns out one of my two instances got rebooted, and now has a different IP. Bad on me for not having it set up so that the apps would come back online on their own, but:

* Does AWS regularly reboot your EC2 instances without warning?
* If so, is there anywhere I could look to find out why this happened?

The only clue I have is that the public IPs are much more similar now, they used to be in very different subnets and now they are on the same /16. Is that a good enough reason to do a forced reboot?

Thanks!
## [4][Q: how do you keep up with AWS new announcements?](https://www.reddit.com/r/aws/comments/g0itdu/q_how_do_you_keep_up_with_aws_new_announcements/)
- url: https://www.reddit.com/r/aws/comments/g0itdu/q_how_do_you_keep_up_with_aws_new_announcements/
---
I am almost sure that at some point you have all joked about how many new products are launched by AWS every year.

As people using AWS products at work, how do you guys keep up with news that are relevant to what you are using?

I found this to be really challenging. I have tried feedly and following up over Twitter but I can't say this was effective.
## [5][Cloud Gaming on Amazon Web Services](https://www.reddit.com/r/aws/comments/g03d95/cloud_gaming_on_amazon_web_services/)
- url: https://medium.com/tensoriot/cloud-gaming-on-amazon-web-services-4be806c0051b
---

## [6][Under the hood: AWS Fargate data plane | Amazon Web Services](https://www.reddit.com/r/aws/comments/fzupjm/under_the_hood_aws_fargate_data_plane_amazon_web/)
- url: https://aws.amazon.com/blogs/containers/under-the-hood-fargate-data-plane/
---

## [7][RDS performance with UUIDs](https://www.reddit.com/r/aws/comments/g04tol/rds_performance_with_uuids/)
- url: https://www.reddit.com/r/aws/comments/g04tol/rds_performance_with_uuids/
---
I've read articles (not specific to RDS) that discourages the use of UUIDs as primary key (vs the usual auto-increment id), if the database is going to be "large" someday. They say it will eventually have performance degration.

Just looking for feedback from anyone using UUIDs on RDS, Mysql or Aurora, that already has 9-digit rows. How's the performance so far?
## [8][What is the most optimized way to the upload large number of file to S3 from local?](https://www.reddit.com/r/aws/comments/fzz4gg/what_is_the_most_optimized_way_to_the_upload/)
- url: https://www.reddit.com/r/aws/comments/fzz4gg/what_is_the_most_optimized_way_to_the_upload/
---
So my use case - 
This is a Personal Project for my images/videos backup.
I have almost 200GBs worth of images/videos, which i need to upload to  S3 for backup. 
I need to keep the requests to a minimum when uploading.

I was thinking if I could write a PY script to ZIP the folders and upload them as single files to S3. Does that makes sense, to achieve my requirement. Or is their better ways of doing this?

How you guys handling similar use cases? Any suggestion is welcome. 

Thank you!
## [9][Newbie Question](https://www.reddit.com/r/aws/comments/g02o4k/newbie_question/)
- url: https://www.reddit.com/r/aws/comments/g02o4k/newbie_question/
---
Is it okay to post questions to this reddit as a newbie and beginner? I have about 2TB of website backups and files to store so am looking for help to get started being able to use AWS S3 to store files and folders and possibly edit them and then re-upload if possible.
## [10][Lambda: trouble importing module, "cannot find module iso-morphic fetch"](https://www.reddit.com/r/aws/comments/g07sil/lambda_trouble_importing_module_cannot_find/)
- url: https://www.reddit.com/r/aws/comments/g07sil/lambda_trouble_importing_module_cannot_find/
---
I'm using AWS Cloud 9 to import a Lambda function "locally". But I can't get it to import isomorphic-fetch: [https://nimbusweb.me/box/attachment/4082785/m493xez0xn2hl8uw1her/fu3KZ9RodmtcNoKK/Screenshot\_2020-04-13\_07.08.06.jpg](https://nimbusweb.me/box/attachment/4082785/m493xez0xn2hl8uw1her/fu3KZ9RodmtcNoKK/Screenshot_2020-04-13_07.08.06.jpg)

Theories why it's not working:

1. `node_modules` is a sibling to the directory with the lambda function. But since `node_modules` is usually searched for recursively I don't think this should cause issues
2. `require` statement doesn't assign fetch to a variable. But with isomorphic fetch I don't believe you're supposed to: [https://www.npmjs.com/package/isomorphic-fetch](https://www.npmjs.com/package/isomorphic-fetch)

Hopefully someone with more experience than me knows what's going on. Here's the full response body:

    Response
    {
        "errorType": "Runtime.ImportModuleError",
        "errorMessage": "Error: Cannot find module 'isomorphic-
    fetch'\nRequire stack:\n- /var/task/index.js\n-
     /var/runtime/UserFunction.js\n- /var/runtime/index.js"
    }
