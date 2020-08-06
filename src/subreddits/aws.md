# aws
## [1][Infrastructure as Code with AWS CDK: Provision a fully-working Fargate cluster with all of the peripheral services...a web service included](https://www.reddit.com/r/aws/comments/i4c91z/infrastructure_as_code_with_aws_cdk_provision_a/)
- url: https://www.reddit.com/r/aws/comments/i4c91z/infrastructure_as_code_with_aws_cdk_provision_a/
---
Hey everyone,

Decided to open-source one of the projects I've been working on (cut it down and cleaned it up). It provisions the following services:

    Fargate
    ECR
    ALB
    SSM Parameter Store (adds values to it)
    ACM Certificate

To Do

* I've cut out a lot of other services that I used, like ElasticSearch and DynamoDB, but might add them back later on.
* Will be working on adding a CDK sub-project for the CodePipeline code as well.
* Working on a YouTube video diving deep into the code.

GitHub repo:

[https://github.com/vbudilov/fargate-cdk-webservice](https://github.com/vbudilov/fargate-cdk-webservice)

&amp;#x200B;

Enjoy,

Vladimir
## [2][State of the subreddit](https://www.reddit.com/r/aws/comments/i46np0/state_of_the_subreddit/)
- url: https://www.reddit.com/r/aws/comments/i46np0/state_of_the_subreddit/
---
* As the subreddit continues to grow, I wanted to take some time to summarize recent changes, ask for assistance, and allow an opportunity to provide feedback.  
* I (u/ckilborn) am back as a mod
   * As a follow up to my statement “I don't want to get involved with approving posts/comments unless they are egregious.” in [https://www.reddit.com/r/aws/comments/g179da/former\_mod\_asking\_for\_rawss\_blessing\_on\_becoming/](https://www.reddit.com/r/aws/comments/g179da/former_mod_asking_for_rawss_blessing_on_becoming/) \- I can confirm that I have not approved/deleted any posts/comments since my return. 
* We (u/Pi31415926/) and I talked about making the mod log public but in the end decided against it - Ex: Too many privacy concerns if someone accidentally posts PII 
* We’ve removed the inactive AWSSupport account as a mod
* Re "AWS Employee" Flair:
   * We’ve cleaned things up and validated all “AWS employee” flair 
   * Goal is to help the community with official voices. 
   * We will be limiting this to 30 users at a time as to not overrun the community. Trying to strike a balance between responsiveness and allowing the community freedom 
* Added a link to r/aws_cdk to the sidebar 
* Weekly discussion posts are created on Monday and pinned until Wed 
   * Ex [https://www.reddit.com/r/aws/comments/hysmye/week\_of\_july\_27th\_what\_are\_your\_favorite\_aws\_tips/](https://www.reddit.com/r/aws/comments/hysmye/week_of_july_27th_what_are_your_favorite_aws_tips/) and [https://www.reddit.com/r/aws/comments/hutwj1/week\_of\_july\_20th\_what\_are\_you\_building\_this\_week/](https://www.reddit.com/r/aws/comments/hutwj1/week_of_july_20th_what_are_you_building_this_week/)
* We are working on setting up some future AMAs with AWS teams 
   * REQUEST: Please let us know what topics you are interested in
* REQUEST: Looking for additional community moderators to help with comment and post approvals
## [3][Ever thought calling APIs from the Glue ETL Job. Free workshop to learn about it.](https://www.reddit.com/r/aws/comments/i4nxzj/ever_thought_calling_apis_from_the_glue_etl_job/)
- url: http://aws-dojo.com/workshoplists/workshoplist2
---

## [4][Traffic to unhealthy instance?](https://www.reddit.com/r/aws/comments/i4pv10/traffic_to_unhealthy_instance/)
- url: https://www.reddit.com/r/aws/comments/i4pv10/traffic_to_unhealthy_instance/
---
I have a target group with one instance and the health check path is a wrong path that causes instance marked as unhealthy instance in target group. 

I redirect traffic on port 80 and 443 to above target group on aws alb. 

So does the unhealthy instance in above target group will still get the traffic (80, 443) from alb, process and respond back to client ad normal with healthy instance?
## [5][How to Create Unlimited Rotating Proxies in AWS](https://www.reddit.com/r/aws/comments/i4ns0n/how_to_create_unlimited_rotating_proxies_in_aws/)
- url: https://medium.com/@devinjaystokes/using-proxycannon-ng-to-create-unlimited-rotating-proxies-fccffa70a728
---

## [6][SSL and proxy](https://www.reddit.com/r/aws/comments/i4nj2q/ssl_and_proxy/)
- url: https://www.reddit.com/r/aws/comments/i4nj2q/ssl_and_proxy/
---
I have a non-SSL private web portal in aws ec2.
http : / / &lt; private-ip&gt;:8080/mywebapp
I am trying to configure Nginx public proxy on this with SSL.
https : / / &lt; public-ip&gt;/myproxywebapp
When I open this proxy public url , I get only text ....no image and styles
In Chrome console , I see these error
&gt;Mixed Content: The page at '&lt;URL&gt;' was loaded over HTTPS, but requested an insecure stylesheet '&lt;URL&gt;'. This request has been blocked; the content must be served over HTTPS
Mixed Content: The page at '&lt;URL&gt;' was loaded over HTTPS, but requested an insecure script '&lt;URL&gt;'.
Mixed Content: The page at 'https : / / &lt; public-ip&gt;/myproxywebapp' was loaded over HTTPS, but requested an insecure image 'http://&lt; public-ip&gt;/xx/xx_resources/logos/xxx-logo-xxxxx.png'. This content should also be served over HTTPS.


This error is understandable , as it says proxy is in SSL but private portal is non SSL and content is getting mixed.

However , I can not change private portal.

What configuration changes I require in nginx to fix this issue ?
## [7][How to export the new AMI value from EC2 Image Builder](https://www.reddit.com/r/aws/comments/i4mtxn/how_to_export_the_new_ami_value_from_ec2_image/)
- url: https://www.reddit.com/r/aws/comments/i4mtxn/how_to_export_the_new_ami_value_from_ec2_image/
---
Hi,

So I have started using AWS reacently and wanted to automate a few processes. The one thing I am having extreme troubles with is automating Auto Scalig Group to automatically update it's AMI-Id after a 

What I want to achieve:  
*EC2 Image Builder builds a new image --//--&gt; Updating ASG's Launch Configuration with a new AMI-Id*

What I have now:  
*EC2 Image Builder builds a new image --&gt; (something missing) --&gt; Lambda Function --&gt; Updated ASG*

I have no problem generating a new Image in Image Builder and I have already set up a Lambda Function that succesfully generates an updated Launch Configuration and assignes it to the desired ASG, but I have no idea how to get the new AMI-Id variable/parameter out of Image Builder to be sent to the lambda function as a variable.

Does any one have something that might help us?
## [8][Start Workspace from Command Line Interface](https://www.reddit.com/r/aws/comments/i4mtma/start_workspace_from_command_line_interface/)
- url: https://www.reddit.com/r/aws/comments/i4mtma/start_workspace_from_command_line_interface/
---
Hello,

I'm looking for some help to script starting an Amazon Workspace from CLI. I'm familiar with the web interface, but it would be handy if I could batch operations. The documentation is not spelling out clear enough for me what I need to do, if someone could provide a sample script to start a workspace I would really appreciate it. I can authenticate from CLI.

Thanks,
## [9][Scaling SOEs in AWS](https://www.reddit.com/r/aws/comments/i4mbee/scaling_soes_in_aws/)
- url: https://www.reddit.com/r/aws/comments/i4mbee/scaling_soes_in_aws/
---
Over the past few years I've spent a bit of time building and remediating Windows and Linux AMI's across medium to large AWS environments.

I have put some of these learnings into a blog post that I hope will help people automate more and driver more reliable deployments in their environments.

* [Scaling SOEs in AWS](https://notes.keiran.io/posts/Scaling_SOEs_in_AWS/)
## [10][[CDK] Good way to create node_modules Lambda Layer](https://www.reddit.com/r/aws/comments/i4m86a/cdk_good_way_to_create_node_modules_lambda_layer/)
- url: https://www.reddit.com/r/aws/comments/i4m86a/cdk_good_way_to_create_node_modules_lambda_layer/
---
```typescript
const depsLayer = new lambda.LayerVersion(this.scope, 'nodeDependencies', {
  code: lambda.Code.fromAsset('./lambda-handler', {
    exclude: ['*', '!package.json', '!package-lock.json'],
    bundling: {
      image: lambda.Runtime.NODEJS_12_X.bundlingDockerImage,
      command: ['bash', '-c', 'mkdir /asset-output/nodejs &amp;&amp; cd $_ &amp;&amp;'
        + 'cp /asset-input/{package.json,package-lock.json} . &amp;&amp;'
        + 'npm ci'],
      environment: { HOME: '/tmp/home' },
    },
  }),
});

```

The node_modules itself can be excluded at calculating asset-hash, by using `exclude` and `bundling` option. The Lambda Layer is updated **only when the package{-lock}.json are changed**.
