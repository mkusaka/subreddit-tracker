# aws
## [1][Week of Aug 31st - What AWS tips do you have?](https://www.reddit.com/r/aws/comments/ik3van/week_of_aug_31st_what_aws_tips_do_you_have/)
- url: https://www.reddit.com/r/aws/comments/ik3van/week_of_aug_31st_what_aws_tips_do_you_have/
---
Share your tips with the community.
## [2][Amazon Relational Database Service - The Basics Of AWS RDS](https://www.reddit.com/r/aws/comments/il3mpc/amazon_relational_database_service_the_basics_of/)
- url: https://catalins.tech/amazon-relational-database-service-the-basics-of-aws-rds
---

## [3][Best way to put an externally manage ALB in front of an EKS cluster?](https://www.reddit.com/r/aws/comments/il032s/best_way_to_put_an_externally_manage_alb_in_front/)
- url: https://www.reddit.com/r/aws/comments/il032s/best_way_to_put_an_externally_manage_alb_in_front/
---
I use Terraform to manage all of my cloud resources. I've been using https://www.getambassador.io/docs/latest/topics/running/ambassador-with-aws/ for a while with a default classic load balancer.

In general (and I wouldn't dare post this in r/kubernetes), I am not a big fan of granting access to my K8s nodes to manage resources outside of the cluster e.g. load balancers. I would rather set up an ALB in front of my cluster in Terraform, point it at an ingress controller in the cluster and call it a day.

This can be done with Amabassador but it requires the use of NodePort and statically choose ports for each cluster node. I _think_ that will work but I'd love to hear about any downsides. I plan to use an ALB with ACM, terminating SSL at the ALB.

All that being said, does anyone else do this? I can't be the only one who wants to manage as much of my infra as possible from Terraform (or CloudFormation). The "normal" way people do this feels like a clash of concerns e.g. letting K8s manage resources outside of a cluster.
## [4][Best way to move 2 large quantities of data to AWS?](https://www.reddit.com/r/aws/comments/iksw7h/best_way_to_move_2_large_quantities_of_data_to_aws/)
- url: https://www.reddit.com/r/aws/comments/iksw7h/best_way_to_move_2_large_quantities_of_data_to_aws/
---
What is the best way to move 2x 150gb PHI (protected health information) files into S3? Can this be done over standard internet connection or is SnowBall necessary?

Also, same question for upcoming larger scenarios: moving 5 2TB PHI files into S3.

&amp;#x200B;

EDIT: Is it possible to make transfer via flash drive HIPAA compliant? Our office has a much faster connection than the client site.
## [5][AWS Site-to-Site VPN now supports Internet Key Exchange (IKE) initiation (It's about freaking time!)](https://www.reddit.com/r/aws/comments/il5q9p/aws_sitetosite_vpn_now_supports_internet_key/)
- url: https://aws.amazon.com/about-aws/whats-new/2020/08/aws-site-to-site-vpn-now-supports-internet-key-exchange-initiation/
---

## [6][Books/resources with focus on aws CLI?](https://www.reddit.com/r/aws/comments/il59iu/booksresources_with_focus_on_aws_cli/)
- url: https://www.reddit.com/r/aws/comments/il59iu/booksresources_with_focus_on_aws_cli/
---
I wasn't able to find any. Anyone here can recommend me something?
## [7][Best way to visualize resources](https://www.reddit.com/r/aws/comments/ikvwg4/best_way_to_visualize_resources/)
- url: https://www.reddit.com/r/aws/comments/ikvwg4/best_way_to_visualize_resources/
---
What’s the best way to visualize the resources in your account?

Follow up-If you join an account how do you go about finding what’s in use/ how it’s architected?
## [8][Switching to Self-Hosted MongoDB.](https://www.reddit.com/r/aws/comments/il4hxb/switching_to_selfhosted_mongodb/)
- url: https://www.loginradius.com/engineering/blog/self-hosted-mongo/
---

## [9][How to get an realistic vision of my monthly S3 Costs?](https://www.reddit.com/r/aws/comments/il4cdb/how_to_get_an_realistic_vision_of_my_monthly_s3/)
- url: https://www.reddit.com/r/aws/comments/il4cdb/how_to_get_an_realistic_vision_of_my_monthly_s3/
---
Hi,

Currently I have a self-hosted Minio server with images on it for my website. I stumbled upon a couple Issues I cannot resolve right now so I decided to look for the actual S3. Only I have no clue what all the settings are about.  


I'm working with about 50GB of high resolution images and 22GB of low resolution images so roughly 75GB together. These are all in my S3 storage (now Minio). I want to move all this data over to the S3.  


The low res images or used in a live search and the first 9 results will allows be loaded and the other 18 if scroll down. These changes based on your search query. The images that are loaded are maximum 1.2MB (that should be minified even more).   


How can I configure something realistic to have a proper cost estimate, preferred higher then it will be. Can someone help me create this estimation?

https://preview.redd.it/fghenegsopk51.png?width=1328&amp;format=png&amp;auto=webp&amp;s=ec06e9b7f83828f8e5cec84d612d74c314a378a9
## [10][CodeBuild OnStateChange firing out of order](https://www.reddit.com/r/aws/comments/il3eyw/codebuild_onstatechange_firing_out_of_order/)
- url: https://www.reddit.com/r/aws/comments/il3eyw/codebuild_onstatechange_firing_out_of_order/
---
I am seeing some weird behaviour with CodeBuild OnStateChange event processing which I don't understand.

For some background, CodeBuild does not natively support integrating with private git repositories.  [Bitbucket.org](https://Bitbucket.org) and [Github.com](https://Github.com) are supported, but private / self-hosted Bitbucket is not.  To accomodate this, a solution is to sync the private repo into CodeCommit, and trigger CodeBuild off that.  To complete the puzzle we must send the build status updates back to the self-hosted Bitbucket server.

What we have in place is a lambda function that triggers of CodeBuild `onStateChange` defined via CDK like so:

         codeBuildProject.onStateChange('StateChange', {
             target: new targets.LambdaFunction(statusUpdateFn)
         })

The function then reads the state of the event object `event.detail['build-status']` and relays that to the Bitbucket server.  Trivial, right?.

&amp;#x200B;

Now to the problem, on short build jobs (less than 20 second), the `onStateChange` event triggers the lambda function 3 times.. in this order:

Lambda start time: 12:06:33.    
`event.time:"12:06:29"`  
`event.detail["build-status"] : "IN_PROGRESS"`

Lambda start time: 12:07:22  
`event.time: "12:07:18"`  
`event.detail["build-status"]: "SUCCEEDED"`

Lambda start time: 12:07:26  
`event.time: "12:06:29"`  
`event.detail["build-status"]: "IN_PROGRESS"`

&amp;#x200B;

The problem is that last run, which has the same \`event.time\` as the first run, but it overwrites/reverses the `"SUCCEEDED"` state back to `"IN_PROGRESS"`.

Can anyone reason as to why that third event (or the second `"IN_PROGRESS"` event) triggers at all?

Is it a bug, or a mis-configuration on our end?

I also tried an alternate version of the CDK code with separate triggers for `onBuildStart` and `onBuildSucceeded`, but we still see 3 events, not 2.

 

&amp;#x200B;

For reference, on a longer build, the `onStateChange` event also triggers the lambda 3 times, but in a more agreeable order:

Lambda start time: 11:02:23  
`event.time:"11:02:16"`  
`event.detail["build-status"] : "IN_PROGRESS"\``

Lambda start time: 11:03:20  
`event.time: "11:02:16"`  
`event.detail["build-status"]: "IN_PROGRESS"`

Lambda start time: 11:03:33  
`event.time: "11:03:27"`  
`event.detail["build-status"]: "SUCCEEDED"`
## [11][Best option for Temporary persisted storage](https://www.reddit.com/r/aws/comments/il2xst/best_option_for_temporary_persisted_storage/)
- url: https://www.reddit.com/r/aws/comments/il2xst/best_option_for_temporary_persisted_storage/
---
Hi.

I have a system and i need to store some data for processing, any ideas on the options i have for places to store it?

1. I have a system that captures some info.
2. I then use that to send to a 3rd party (service1) which returns me a uuid.
3. A user then performs an action
4. Service 1 then calls a webhook and sends me the uuid and the name of the event the user performed 
5. When the event happens (up to 5 days later) i need to send some info captured on step 1 to another 3rd party (service2)

After i send the info to service2 i don't need it anymore, service2 is the source of truth and it would be confusing to store it in my app.

I am thinking dynomodb with they key as the uuid i get from service1 and the value the info i send later. Once sent i delete the record.

 
Are there any other aws services that might suit?
