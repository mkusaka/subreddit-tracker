# aws
## [1][Week of Aug 24th - What are you building this week on AWS?](https://www.reddit.com/r/aws/comments/ifoebu/week_of_aug_24th_what_are_you_building_this_week/)
- url: https://www.reddit.com/r/aws/comments/ifoebu/week_of_aug_24th_what_are_you_building_this_week/
---
Share what you are creating on AWS
## [2][The good parts of AWS - A visual summary](https://www.reddit.com/r/aws/comments/igusc5/the_good_parts_of_aws_a_visual_summary/)
- url: https://hassenchaieb.com/aws-good-parts/
---

## [3][Code-first GraphQL server development](https://www.reddit.com/r/aws/comments/igyemq/codefirst_graphql_server_development/)
- url: https://blog.graphqleditor.com/graphql-nexus-codefirst-graphql-server/
---

## [4][Best practise for running monthly webscraping and processing jobs.](https://www.reddit.com/r/aws/comments/igxgmt/best_practise_for_running_monthly_webscraping_and/)
- url: https://www.reddit.com/r/aws/comments/igxgmt/best_practise_for_running_monthly_webscraping_and/
---
I have a webscraping module, (currently inside a jupyter notebook using AWS sagemaker), i want to run the code once a month, and store it inside a  *sqlite3* database (that is currently also hosted on the notebook). 

I have looked into *aws batch jobs,* and having *multiple EC2 spot instances (as a cluster)*, to run something like this. 

I'm still pretty new to AWS, so I'm not sure, what the best practice is regarding this. 

In terms of processing, it is a faily short job, and could be completed within a 5-10 minutes by a regular ec2 t2.small - t2.medium machine.

&amp;#x200B;

What are you guys opinion on this?
## [5][[AWS Sagemaker] Running a Jupyter Notebook for more than 24 hours](https://www.reddit.com/r/aws/comments/igwybs/aws_sagemaker_running_a_jupyter_notebook_for_more/)
- url: https://www.reddit.com/r/aws/comments/igwybs/aws_sagemaker_running_a_jupyter_notebook_for_more/
---
I am currently using Dask to run operations on a large dataset and simultaneously storing the contents into smaller files in S3. The operations take a significant amount of time (over 24 hours). The problem is that writing into S3 stops after some hours and the generation of new files is incomplete. How can I ensure that process can finish outputting files without stopping? Currently I am monitoring the process by going to S3 and calculating the size and the number of files. If it is unchanged, I know that the process suddenly stopped before completion.


Should I keep the jupyter notebook instance on-off? Should I shut down the notebook instance? Is it safe to shut down my PC and open Sagemaker the next day? Most importantly, how can I ensure that the process is completed and doesn't stop abruptly?
## [6][Managing Workspaces patching and software deployment with SCCM?](https://www.reddit.com/r/aws/comments/igowba/managing_workspaces_patching_and_software/)
- url: https://www.reddit.com/r/aws/comments/igowba/managing_workspaces_patching_and_software/
---

We would like to avoid having a different process for application and Windows updates deployment to AWS Workspaces than we have for our physical workstations. 

Has anyone set up Workspaces as SCCM clients?  How did licensing them as SCCM clients work out for you?
## [7]["Hacking" AWS to get more cpu burst credits?](https://www.reddit.com/r/aws/comments/igyr0o/hacking_aws_to_get_more_cpu_burst_credits/)
- url: https://www.reddit.com/r/aws/comments/igyr0o/hacking_aws_to_get_more_cpu_burst_credits/
---
Thanks to all the lovely people in this forum to answer my stupid question [post](https://old.reddit.com/r/aws/comments/igjam5/aws_ec2_performance_degrades_after_three_hours/) yesterday. Here's another one:

The t2 instance I'm using gets throttled after the 4th hour of running because I run out of cpu burst credits on the instance. I've given a certain number as the instance starts and then AWS throttles usage after. Makes sense, right? 

So my question is... is it possible to use the AWS API to:

* Create an instance

* Run it for four hours

* Terminate the instance

* Reactivate a new instance

And still pay the standard t2.small rate? Would I be charged for the extra credits that I use that go "above and beyond" what I ordinarily would use when the t2 instance would be throttled? 

I must assume that AWS has a way to preventing this sort of nonsense by charging for additional credit usage but I figured it's worth a shot and I can't find answers to this question anywhere.
## [8][What happened to Amazon SNS-&gt;SMS?](https://www.reddit.com/r/aws/comments/igy2y5/what_happened_to_amazon_snssms/)
- url: https://www.reddit.com/r/aws/comments/igy2y5/what_happened_to_amazon_snssms/
---
Anyone know what's going on with the SMS service? I am getting "internal error" for all my outgoing messages and my spending limit did not go to $25 as all previous months, it shows the standard $1.

Haven't received any reply to my quota increase nor regarding any possible issue with the system.

Not sure if this is a temporary issue or if I should look for a new gateway service.
## [9][AWS Cognito and problem with sending invitation email](https://www.reddit.com/r/aws/comments/igurdh/aws_cognito_and_problem_with_sending_invitation/)
- url: https://www.reddit.com/r/aws/comments/igurdh/aws_cognito_and_problem_with_sending_invitation/
---
 My desired scenario is this:

1. User is added to user pool with AWS Amplify.
2. User's email is auto verified with presignup trigger.
3. User is auto confirmed with presignup trigger.
4. User receives an invitation email from Cognito.

My problem is with the invitation email. It is not sent. Everything else prior to that is working. I tried without auto-verify and auto-confirm and then the confirmation code is sent by email just fine.

How can I fix this issue? Is there some issue that the email isn't sent if users are auto confirmed?
## [10][Recommendations for docker container management/static IPs for containers](https://www.reddit.com/r/aws/comments/igsyr9/recommendations_for_docker_container/)
- url: https://www.reddit.com/r/aws/comments/igsyr9/recommendations_for_docker_container/
---
A team and I are currently working on a project that involves running multiple applications on different machines. We are using AWS, Terraform, and ECR.

Ideally we would use ECS (via terraform) to manage different tasks; however, we need to be able to specify what tasks run on what EC2 instances. This is necessary because some applications must run on separate VMs since they need different public IPs. Fargate would solve this issue but it is too costly. We tried using the tag feature for ECS but tags cannot be dynamically generated via Terraform, and manually configuring them on the AWS UI is not a scalable solution.

On another note our team has considered using Kubernetes via Terraform to restrict containers to specific pods, but setting up Kubernetes via Terraform on AWS is a hassle as AWS tries to restrict you to EKS.

Our team is restricted to AWS since other infrastructure already exists there.

It is important to note that this project is still in its early stages so changing the overall stack (not the cloud provider) is still an option.

TLDR; what is a scalable way via IAC to provision machines and specifically manage docker containers across different machines? or what is a scalable way via IAC to manage docker containers with different public IPs?
## [11][Anyone else haivng large AWS incident / outtage causing issues. Appears mostly related to EU-WEST2a](https://www.reddit.com/r/aws/comments/ig90ou/anyone_else_haivng_large_aws_incident_outtage/)
- url: https://i.imgur.com/Ajikl0m.png
---

