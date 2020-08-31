# aws
## [1][AWS workspace choppy mouse over text](https://www.reddit.com/r/aws/comments/ijwtbo/aws_workspace_choppy_mouse_over_text/)
- url: https://www.reddit.com/r/aws/comments/ijwtbo/aws_workspace_choppy_mouse_over_text/
---
So I've check this on a couple different workspaces on a couple different performance levels.  Full screen or windows.  Whenever I have something open that allows typing..   notepad, notepad++, the url bar in a browser, etc..   anything I can "type" in.   The mouse movement over that text area is very choppy.  The rest of the virtual machine is very smooth.

Seems to be anytime the mouse pointing turns to the "I" text cursor.

EDIT:  this is a windows desktop.  Using amazon workspace client.
## [2][Log Management solutions](https://www.reddit.com/r/aws/comments/ijlx0k/log_management_solutions/)
- url: https://www.reddit.com/r/aws/comments/ijlx0k/log_management_solutions/
---
I’m creating an application in AWS that uses Kubernetes and some bare EC2. I’m trying to find a good log management solution but all hosted offerings seem so expensive. I’m starting my own company and paying for hosting myself so cost is a big deal. I’m considering running my own log management server but not sure on which one to choose. I’ve also considered just uploading logs to CloudWatch even though their UI isn’t very good. What has others done to manage logs that doesn’t break the bank?
## [3][My first technical blog post: A Deep Dive into Serverless Tracing with AWS X Ray &amp; Lambda](https://www.reddit.com/r/aws/comments/ijolqf/my_first_technical_blog_post_a_deep_dive_into/)
- url: https://www.reddit.com/r/aws/comments/ijolqf/my_first_technical_blog_post_a_deep_dive_into/
---
Hey all,

This post is a long time coming. After spending quite a few hours emailing back and forth with support and digging deep into documentation, git issues, and stack overflow (as well as good old fashion experimentation) I decided to document my research as a resource. It's quite long even after taking out a lot of irrelevant details, but I hope it gets you as far as you need to go in adding X Ray to your lambdas. 

If you're interested, check it out below:

[https://medium.com/@5t33/a-deep-dive-into-serverless-tracing-with-aws-x-ray-lambda-5ff1821c3c70](https://medium.com/@5t33/a-deep-dive-into-serverless-tracing-with-aws-x-ray-lambda-5ff1821c3c70)  


I'd also appreciate any feedback.
## [4][Cloudfront Compress Objects Automatically not working](https://www.reddit.com/r/aws/comments/ijwhcb/cloudfront_compress_objects_automatically_not/)
- url: https://www.reddit.com/r/aws/comments/ijwhcb/cloudfront_compress_objects_automatically_not/
---
 I have setup S3 as a host for my website with cloudfront as cdn. In gtmetrix I got feedback to compress the static files to improve the performance. I have enabled Content length in CORS configuration for S3 and could see content length in response field of S3 and also have switched on Compress Objects Automatically in cloudfront but still the delivered filed is not in gzip format. Below are the s3 and cloudfront url links.

S3 link :- [https://s3.ap-south-1.amazonaws.com/www.peervadoo.com/assets/img/ill/ill1.jpg](https://s3.ap-south-1.amazonaws.com/www.peervadoo.com/assets/img/ill/ill1.jpg)

Cloudfront link :- [https://www.peervadoo.com/assets/img/ill/ill1.jpg](https://www.peervadoo.com/assets/img/ill/ill1.jpg)
## [5][How do I auto-start my Python Flask web app every time my AWS EC2 instance boots up?](https://www.reddit.com/r/aws/comments/ijuoog/how_do_i_autostart_my_python_flask_web_app_every/)
- url: https://www.reddit.com/r/aws/comments/ijuoog/how_do_i_autostart_my_python_flask_web_app_every/
---
Currently what i do is &gt; Boot up instance &gt; then manually execute the following .sh file:

    sudo systemctl daemon-reload
    sudo service nginx restart
    sudo service gunicorn3 restart

Accordingly, i put the following in the "View/Change User Data" field, as i believe that's where startup commands are to be inputted:

    #!/bin/bash
    
    sudo systemctl daemon-reload
    sudo service nginx restart 
    sudo service gunicorn3 restart

However, the Python Flask web app still doesn't automatically start. Any idea what i'm doing wrong? Thanks for reading.
## [6][Isn't AmazonSSMManagedInstanceCore still too permissive?](https://www.reddit.com/r/aws/comments/ijsrj7/isnt_amazonssmmanagedinstancecore_still_too/)
- url: https://www.reddit.com/r/aws/comments/ijsrj7/isnt_amazonssmmanagedinstancecore_still_too/
---
The policy allows `ssm:GetParameter(s)` for everything, why is this required? This looks insecure default, it allows reading configuration for other apps.
It probably can't read encrypted value without kms permission though.

I don't know if I can change statement Resource to something other than '*'. Do some permissions runs action for ec2? Most instances in my account runs on ASG, so I can't just pass instance id. Or, could I restrict actions to EC2 running on certain ASG?
## [7][AWS CloudFront - We're logging TLS 1.3 traffic?](https://www.reddit.com/r/aws/comments/ijsjpx/aws_cloudfront_were_logging_tls_13_traffic/)
- url: https://www.reddit.com/r/aws/comments/ijsjpx/aws_cloudfront_were_logging_tls_13_traffic/
---
TLS 1.3 for CloudFront has been on the roadmap for S2N for a while now -&gt; [https://github.com/awslabs/s2n/projects](https://github.com/awslabs/s2n/projects)

I've noticed in the last few months though, we're seeing TLS 1.3 traffic on our CloudFront distribution logs. Are we seeing TLS 1.3 sneakily snuck in for testing?

This is pretty awesome. Wondering if anyone can shed any light on this? I haven't seen any comms recently.
## [8][Can someone help me figure out aws api gateway?](https://www.reddit.com/r/aws/comments/ijvg8m/can_someone_help_me_figure_out_aws_api_gateway/)
- url: https://www.reddit.com/r/aws/comments/ijvg8m/can_someone_help_me_figure_out_aws_api_gateway/
---
We have 2 concurrent ruby applications running which handle our systems. We're planning on phasing out the old one, but some of our clients still use the old system to make some api calls. Now we cannot ask the clients to change their process, so someone suggested that we use api gateway and emulate the old system and push data to new system though it. Would this be feasible? Can someone share a case study around this? 

Any help would be appreciated. 
Thanks.
## [9][Prevent AWS from reading your step functions data](https://www.reddit.com/r/aws/comments/ijxv47/prevent_aws_from_reading_your_step_functions_data/)
- url: https://blog.theodo.com/2020/08/secure-aws-step-functions-sensitive-data/
---

## [10][AWS Step functions alternative](https://www.reddit.com/r/aws/comments/ijxqk5/aws_step_functions_alternative/)
- url: https://www.reddit.com/r/aws/comments/ijxqk5/aws_step_functions_alternative/
---
Hello everyone

I'm working on a startup that is creating a product/service that can be an alternative to AWS Step functions. 
In essence we're creating a software orchestration platform that allows you to configure a workflow of various steps that trigger, for example, a Lamba function or another piece of code. 

While working on a PoC/MVP I wanted to ask some feedback on how you are currently using AWS Step functions, what are the good things and the not so good things in your opinion. 

Thanks for any feedback you can give me!
