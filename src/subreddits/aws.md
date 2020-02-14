# aws
## [1][Protect S3 Images with own IAM Concept](https://www.reddit.com/r/aws/comments/f3r7vh/protect_s3_images_with_own_iam_concept/)
- url: https://www.reddit.com/r/aws/comments/f3r7vh/protect_s3_images_with_own_iam_concept/
---
I have an App that should store and retrieve images from S3. Currently, upload happens via my own application server endpoint to S3 (image on app-&gt;application server-&gt;store in S3).

However, whats the best way to protect image access then, without giving every single user some IAM role on S3 nor share a direct S3 link?

My only solution would be that my application servers retrieve the image, and then forward them to the client. This causes two HTTP roundtrips though (S3-&gt;Application Server-&gt;App). This is massive for huge pictures.

Any other ideas?
## [2][In which I figure out to run Docker Compose successfully on ECS :-)](https://www.reddit.com/r/aws/comments/f3pp0q/in_which_i_figure_out_to_run_docker_compose/)
- url: https://rmoff.net/2020/02/13/adventures-in-the-cloud-part-94-ecs/
---

## [3][Hello AWS Session Manager; Farewell SSH](https://www.reddit.com/r/aws/comments/f3buy3/hello_aws_session_manager_farewell_ssh/)
- url: https://medium.com/@dnorth98/hello-aws-session-manager-farewell-ssh-7fdfa4134696
---

## [4][Amazon S3 - Price for huge files/data retrieval ?](https://www.reddit.com/r/aws/comments/f3quyv/amazon_s3_price_for_huge_filesdata_retrieval/)
- url: https://www.reddit.com/r/aws/comments/f3quyv/amazon_s3_price_for_huge_filesdata_retrieval/
---
Hi everyone,

I'm creating a file hosting website. Currently, the goal is for users to share files to a lot of people, and files will be deleted after a time of inactivity. I've been looking over Google Cloud, AWS and Azure, and something strikes me.   
In GC/AWS, the biggest cost is clearly data retrieval (90$/Tb approximately). But in Azure Storage, retrieval of data is absolutely free. Is there something i'm missing? Because in that case, Azure Storage would be far better than AWS S3 or GC Storage.
## [5][A 7KB AWS lambda Node.js library with zero runtime dependencies](https://www.reddit.com/r/aws/comments/f3rse9/a_7kb_aws_lambda_nodejs_library_with_zero_runtime/)
- url: https://www.npmjs.com/package/micro-aws-lambda
---

## [6][Can a lambda in child account(Trusting account) change/update the Org wide tag policies in the master account(Trusted account) via sts:AssumeRole?](https://www.reddit.com/r/aws/comments/f3rhsz/can_a_lambda_in_child_accounttrusting_account/)
- url: https://www.reddit.com/r/aws/comments/f3rhsz/can_a_lambda_in_child_accounttrusting_account/
---
I was reading up on Assume Role and how it works and it got me thinking will the lambda which is in my child account(Trusting account), be able to modify or update the Organization-wide Tag Policies in the master account(Trusted account)?

I haven't tried this till now, as I don't have access to the Master account. 

My main concern is: As the tag policies and the control policies will be only accessible from master account, will the assume role method actually work in this case?
## [7][Getting the current shadow state for a thing on IoTCore](https://www.reddit.com/r/aws/comments/f3rhln/getting_the_current_shadow_state_for_a_thing_on/)
- url: https://www.reddit.com/r/aws/comments/f3rhln/getting_the_current_shadow_state_for_a_thing_on/
---
Hey,

I'm trying to find a way to retrieve the current shadow state for a device I have on IoTCore using Python SDK but no luck. What am I missing? what easy programmable way do I have for getting the shadow state of a device?

&amp;#x200B;

    c.connect()
    MQTTClient = c.getMQTTConnection()
    MQTTClient.configureOfflinePublishQueueing(100)
    
    def callback(payload, responseStatus, token):
        print("inside the callback")
        print("--", payload)
        print("--", responseStatus)
        print("--", token)
    
    handler = c.createShadowHandlerWithName("0f231d3c-31d3-bb25-a571-9cbb25ec2724", True)
    print(handler.shadowGet(callback, 5))

I only get the Token that is returned from the callback, and it prints nothing to the screen.

Thanks!
## [8][Cannot get video calling working with the new WorkSpace Streaming Protocol](https://www.reddit.com/r/aws/comments/f3qw2y/cannot_get_video_calling_working_with_the_new/)
- url: https://www.reddit.com/r/aws/comments/f3qw2y/cannot_get_video_calling_working_with_the_new/
---
Hey All, 

Testing the new WorkSpace Streaming Protocol which is replacing PCoIP 

https://aws.amazon.com/workspaces/wsp/

One of the big new features for us is bi-directional video so we can now use it for video calling. 

Spun up new WorkSpaces in a new Directory and logged in using the new WorkSpace client but its still not detecting my webcam. 

Is there anything  additional I need to do get video calling working? 

Thanks.
## [9][AWS Exam - Passed/Failed on screen message](https://www.reddit.com/r/aws/comments/f3qgxm/aws_exam_passedfailed_on_screen_message/)
- url: https://www.reddit.com/r/aws/comments/f3qgxm/aws_exam_passedfailed_on_screen_message/
---
I know this has been brought up before, but does anyone know what the final "on screen message" is for the AWS exam is if you pass or fail??

*Like many, I quickly read the message expecting the final email would contain the results (which it doesn't)... now i'm second guessing if I actually passed... frustrating waiting for the final results.*

From memory, it went something like "....passed....", then a blur of somewhat relief. just hoping the failed message doesn't start with "you have NOT passed". Would love any input :)
## [10][EKS pricing question](https://www.reddit.com/r/aws/comments/f3ogx5/eks_pricing_question/)
- url: https://www.reddit.com/r/aws/comments/f3ogx5/eks_pricing_question/
---
How do I calculate the cost of running an 8-node cluster for 100 hours a month in EKS - the Simple Monthly Calculator doesn't contain that service at all to work out an estimate?! I read somewhere that you have to pay for the master plane on top too....
