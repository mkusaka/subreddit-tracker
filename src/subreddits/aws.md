# aws
## [1][[HELP] Using CloudFront as whole site routing - implications of the limited amount of cache behaviours and other ideas?](https://www.reddit.com/r/aws/comments/i12a6y/help_using_cloudfront_as_whole_site_routing/)
- url: https://www.reddit.com/r/aws/comments/i12a6y/help_using_cloudfront_as_whole_site_routing/
---
Hey guys, i need your assistance :)

My company is in the process of migrating to AWS and we are discussing how to implement frontend routing for our main website. We have quite a load of micro-frontends which are served on a single host, so we have like example.com/search, example.com/details/&lt;id&gt;, example.com/profile etc. which all are different apps under the hood, served via Fargate or Lambda. Routing still is legacy code which uses some old reverse Proxy, so we currently have no on-edge routing (except for assets which are mostly served from CloudFront)

My idea was to create a CloudFront Distribution which uses Cache-Behaviours as routing for the different apps. For more dynamic routing the default behavior could use an ELB as fallback or maybe like a Lambda or Fargate service that serves as a proxy. I also stumbled across [this article](https://aws.amazon.com/de/blogs/networking-and-content-delivery/dynamic-whole-site-delivery-with-amazon-cloudfront/).

What i am not quite sure about is whether the limited amount of cache behaviours (25) has any implications. I know that it can be increased by contacting support, but does the limit have a reason? Like, does it get more expensive if we increase the contingent?

Someone else had the idea to use a lambda@edge with a DynamoDB or something like that for route matching and proxying.

Another idea was that the default behavior for the CloudFront routes to another CloudFront which itself can have another set of cache behaviors to circumvent the 25 limit.

Would love to get some feedback on this.
## [2][Join the AWS Community Builders Program](https://www.reddit.com/r/aws/comments/i0s77b/join_the_aws_community_builders_program/)
- url: https://www.reddit.com/r/aws/comments/i0s77b/join_the_aws_community_builders_program/
---
I was part of the beta program of AWS Community Builders Program, which became public today.

That means you can apply to join the community too! Check the link for more details! 

[https://aws.amazon.com/developer/community/community-builders/](https://aws.amazon.com/developer/community/community-builders/)
## [3][AWS ec2 instance autoscaling guidance?](https://www.reddit.com/r/aws/comments/i16ii3/aws_ec2_instance_autoscaling_guidance/)
- url: https://www.reddit.com/r/aws/comments/i16ii3/aws_ec2_instance_autoscaling_guidance/
---
ok so i have successfully created an autoscaling group but i need help with the scale out policy.

So currently my min instances = 1, desired = 1, max = 10..policy = CPU over 70%

so i did load tests and it didnt create new instances, and i was getting lots of 500s

i manually changed my min instances = 4, desired = 4, max = 10..ran the test again and the 500s droppd down to 0...

i looked at the policy and theres another option for network in/out, and load balancer but not sure what numbers to put in?

https://preview.redd.it/c1o8w8g6z6e51.png?width=1012&amp;format=png&amp;auto=webp&amp;s=d262e68599327b3d59946b37cedde272762c5cb2
## [4][Ebextensions increase max_input_vars allowed](https://www.reddit.com/r/aws/comments/i173vd/ebextensions_increase_max_input_vars_allowed/)
- url: https://www.reddit.com/r/aws/comments/i173vd/ebextensions_increase_max_input_vars_allowed/
---
Hi All, 
I've got the instance to update the max_input_vars to 3000, but I'm still seeing a message for`PHP Max Input Vars (allowed) 1000`.

Current ebextension file

    files:
        "/etc/php.d/99uploadsize.ini"
            mode: "00064"
            owner: root
            group: root
            content: |
               upload_max_filesize = 100M
               post_max_size = 100M
               max_input_vars = 3000
    commands:
      remove_old_ini:
          command: "rm -f /etc/php.d/99uploadsize.ini.bak"

further info: this is for a client's wordpress site. they are insisting that we use this particular backend content builder, and it is throwing these warnings.
## [5][please help](https://www.reddit.com/r/aws/comments/i158zy/please_help/)
- url: https://www.reddit.com/r/aws/comments/i158zy/please_help/
---
i have started aws free tier a month ago but why am i being charged!??

&amp;#x200B;

https://preview.redd.it/wuks53f476e51.png?width=1124&amp;format=png&amp;auto=webp&amp;s=f6914de62e8772ac6f5486f4a3e4cd14258238d4
## [6][DLSB Workshop - Access Denied for Java resources](https://www.reddit.com/r/aws/comments/i15490/dlsb_workshop_access_denied_for_java_resources/)
- url: https://www.reddit.com/r/aws/comments/i15490/dlsb_workshop_access_denied_for_java_resources/
---
Hey

I recently attended the QLDB workshop and the Java classes used in the workshop ([App.java](https://App.java) and [SlowUpdate.java](https://SlowUpdate.java)) were provided in the slide deck in the form of URLs. 

I tried to access the java files but it is throwing Access Denied error. Any reason as to why this is happening? And from where can I get these resources if not the links provided in the slide deck.

Thank You.
## [7][2nd Network Interface Ubuntu EC2](https://www.reddit.com/r/aws/comments/i10tpl/2nd_network_interface_ubuntu_ec2/)
- url: https://www.reddit.com/r/aws/comments/i10tpl/2nd_network_interface_ubuntu_ec2/
---
Hey everyone,
I have a EC2 instance running Ubuntu. I created a network interface, and I attached it to the instance.
I SSH'd into the Ubuntu machine and the new second interface is showing down and is not showing any IP address at all.
What do I need to do to add a 2nd interface to a running EC2 ubuntu instance????
Thank you all!!!
## [8][Refactoring a distributed monolith to microservices using Amazon Event Bridge and implementing BASE consistency to decouple our microservices.](https://www.reddit.com/r/aws/comments/i0sok6/refactoring_a_distributed_monolith_to/)
- url: https://www.rehanvdm.com/serverless/refactoring-a-distributed-monolith-to-microservices/index.html
---

## [9][Managing more than 5 accounts](https://www.reddit.com/r/aws/comments/i0z8uu/managing_more_than_5_accounts/)
- url: https://www.reddit.com/r/aws/comments/i0z8uu/managing_more_than_5_accounts/
---
How are people managing more than 5 accounts? The console has that fun limit where it will remember the last 5 accounts, I have 6 which I move between regularly enough that I feel the friction.

I have seen chrome extensions like 
https://github.com/tilfin/aws-extend-switch-roles

But worry about chrome extensions being compromised and session data exfiled. 

I could fork, audit, install locally?

Thoughts?
## [10][re:Invent 2020 will be free and virtual!](https://www.reddit.com/r/aws/comments/i06b6w/reinvent_2020_will_be_free_and_virtual/)
- url: https://reinvent.awsevents.com
---

