# aws
## [1][Amazon FSx file share - multi-region](https://www.reddit.com/r/aws/comments/fvxusk/amazon_fsx_file_share_multiregion/)
- url: https://www.reddit.com/r/aws/comments/fvxusk/amazon_fsx_file_share_multiregion/
---
Hi all,

Does anyone have experience in setting up FSx in multi-region? I'd like to copy or synchronize the FSx file shares (windows) to another region. I know multi-az is possible. We're using DFS. The only thing I can think of is VPC peering, creating new file shares in the other region and setting up DFS replication by adding these new file shares to the AD in the first region. Anyone that has experience with this? Or another idea maybe?

&amp;#x200B;

Thanks!
## [2][When to define permissions in identity-based policy vs resource-based policy?](https://www.reddit.com/r/aws/comments/fvtlrq/when_to_define_permissions_in_identitybased/)
- url: https://www.reddit.com/r/aws/comments/fvtlrq/when_to_define_permissions_in_identitybased/
---
When defining permissions for AWS resources like S3, KMS, etc, how do I determine whether to define the permission in an identity-based policy attached to the IAM user/group/role or a resource-based policy attached to the resource?
## [3][AWS Ireland virtual hiring event](https://www.reddit.com/r/aws/comments/fvi5ia/aws_ireland_virtual_hiring_event/)
- url: https://www.reddit.com/r/aws/comments/fvi5ia/aws_ireland_virtual_hiring_event/
---
AWS RDS and ElastiCache are looking to recruit for our Dublin #ireland  teams 👍🏻 software development engineers.Interviews will be scheduled in the last 2 weeks of April.

Our software developers build the next generation technologies that change how millions of #AWS customers connect, and interact with the AWS services ecosystem. We use ideas from every facet of computer science including distributed computing, large-scale design, big and real-time data processing, data storage, service oriented architecture, and networking. We are looking for highly-motivated and passionate engineers to build our next generation high performance data storage platforms that solve real-time query, transaction and analytics processing needs for large scale data applications.  https://www.amazon.jobs/en/jobs/1076231/software-development-engineer-hiring-event-in-april?sc_channel=sm&amp;sc_campaign=Recruiting_Hiring&amp;sc_publisher=TWITTER&amp;sc_country=Global&amp;sc_geo=EMEA&amp;sc_outcome=awareness&amp;trkCampaign=virtual_event_dublin_sde_April&amp;trk=virtual_event_dublin_sde_April_TWITTER&amp;sc_content=virtual_event_dublin_sde_April&amp;linkId=85853464  #amazonwebservices #awscloud #cloud
## [4][Building website(eCommerce) on AWS vs GCP](https://www.reddit.com/r/aws/comments/fvxssw/building_websiteecommerce_on_aws_vs_gcp/)
- url: https://www.reddit.com/r/aws/comments/fvxssw/building_websiteecommerce_on_aws_vs_gcp/
---
Hello all,

I  have been using GCP for a couple of years now to host websites on cloud  mainly, I have never tried AWS. For now  I'm trying to built my first eCommerce website, how much can AWS benefit  me through that and how is it different from GCP?

Does AWS support Blockchain?

Do you recommend building all the website database on a blockchain server?
## [5][Logs out of files in pods in EKS into Cloudwatch logs streams](https://www.reddit.com/r/aws/comments/fvwvny/logs_out_of_files_in_pods_in_eks_into_cloudwatch/)
- url: https://www.reddit.com/r/aws/comments/fvwvny/logs_out_of_files_in_pods_in_eks_into_cloudwatch/
---
I've got a legacy app that has been dropped into EKS. The app creates logs in different files. eg catalina.out, application.log, access.log.

I don't really want all that going down stdout/stderr. I'd like it to use Cloudwatch and take each file and send it to a different log group. And ideally the cloudwatch config to come from an SSM parameter rather than being baked in at build time.

Someone has suggested sidecars. And then went off on a total tangent about all sort of other applications they can use and still end up pushing to stdout/err (using "fluentd").

Is it possible to put Cloudwatch into a sidecar in a way that it can access files in the application pods and send the logs to groups correctly?? It seems like this should be so simple but the people who claim to know EKS/k8s keep getting tied up in knots and not fixing it the right way.

Any examples in github/whitepapers/blogs etc. I can show them?

(I'm torn between this being an AWS question and a k8s question!)
## [6][AWS service names - AWS vs Amazon. How do they make the distinction?](https://www.reddit.com/r/aws/comments/fvbkqa/aws_service_names_aws_vs_amazon_how_do_they_make/)
- url: https://www.reddit.com/r/aws/comments/fvbkqa/aws_service_names_aws_vs_amazon_how_do_they_make/
---
AWS makes distinction between services like Amazon DynamoDB vs AWS CloudTrail.

When I worked for AWS I did my speaker certification - an internal cert that allows one to speak on behalf of AWS. In that, they were very specific that one could not say "S3", but one had to say "Amazon S3". Now I've just realised that the homepage of the documentation ([https://docs.aws.amazon.com/](https://docs.aws.amazon.com/)) actually shows the various services prefixed by Amazon or AWS (wish I had figured this out while doing my speaker cert).

But I wonder how they make the distinction? 

Initially I thought it was "older services like S3 would have Amazon preceding them" - like Amazon S3 or Amazon DynamoDB, but I now see that DocumentDB is an "Amazon" service and not an "AWS" service and that's a new'ish service.

Anyone have any thoughts on this?
## [7][lambda layer issue](https://www.reddit.com/r/aws/comments/fvnxik/lambda_layer_issue/)
- url: https://www.reddit.com/r/aws/comments/fvnxik/lambda_layer_issue/
---
when trying to add a layer to my function i keep getting "You are not authorized to perform: lambda:GetLayerVersion."

what am i doing wrong?

&amp;#x200B;

EDIT: i think i figured it out. i think the account where this layer is shared doesn't have proper permission for me. 
## [8][When should you use AWS Secrets Manager vs KMS?](https://www.reddit.com/r/aws/comments/fvj0nm/when_should_you_use_aws_secrets_manager_vs_kms/)
- url: https://www.reddit.com/r/aws/comments/fvj0nm/when_should_you_use_aws_secrets_manager_vs_kms/
---
Trying to understand the difference
## [9][scale specific tasks?](https://www.reddit.com/r/aws/comments/fvpo4s/scale_specific_tasks/)
- url: https://www.reddit.com/r/aws/comments/fvpo4s/scale_specific_tasks/
---
if I have 100 tasks, all with their own specific purpose, how can I choose which to scale in? when one task completes its purpose, I want to scale it in, but scaling in tasks at random screws that up badly. there’s only a 1/100 chance I scale the right task!
## [10][I have a hard time committing to ECS vs EKS](https://www.reddit.com/r/aws/comments/fvem57/i_have_a_hard_time_committing_to_ecs_vs_eks/)
- url: https://www.reddit.com/r/aws/comments/fvem57/i_have_a_hard_time_committing_to_ecs_vs_eks/
---
I understand that it's supposed to be easier to learn and that the fargate service is well integrated with ECS but nobody in the industry uses it on the same level that they use EKS (or kubernetes in general).

So the question is if you were starting out with containers on AWS would you use EKS or ECS:

[View Poll](https://www.reddit.com/poll/fvem57)
