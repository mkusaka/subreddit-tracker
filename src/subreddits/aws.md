# aws
## [1][AWS Wish List 2020](https://www.reddit.com/r/aws/comments/jbw85e/aws_wish_list_2020/)
- url: https://www.reddit.com/r/aws/comments/jbw85e/aws_wish_list_2020/
---
&amp;#x200B;

AWS always releases a bunch of features, sometimes everyday or atleast once a week. Here is my wish list of the features I want to see as a part of AWS infrastructure

1: AWS Managed Proxy Server(Rather than spinning own squid server)

2: EBS replication across different availability zones(Possible? Legal constraints?)

3: Multi-region VPC(Possible? Legal constraints?)

4: UI to debug boot issues(Better then EC2 Get Instance Screenshot and Instance logs)

5: Support tagging for every individual service(It's improving)

6: VPC endpoints support for every service (EKS?) 

7: EC2 instance live migration

8: Display AWS Cli  while resource creation(Similar to GCP)

9: Cost calculation while resource creation(AWS start supporting(for example, RDS) this feature but not for every service

10: More features in App Mesh(Circuit breaker, Rate Limiting)

P.S: Not sure if some features are already available, but if something is missing, please feel free to add
## [2][CodePipeline with multiple branches seems absurdly complex](https://www.reddit.com/r/aws/comments/ji5gql/codepipeline_with_multiple_branches_seems/)
- url: https://www.reddit.com/r/aws/comments/ji5gql/codepipeline_with_multiple_branches_seems/
---
[https://aws.amazon.com/blogs/devops/multi-branch-codepipeline-strategy-with-event-driven-architecture/](https://aws.amazon.com/blogs/devops/multi-branch-codepipeline-strategy-with-event-driven-architecture/)

We are considering moving to CodePipeline and we use GitFlow as our release methodology in which we deploy "release branches" which are branches cut off develop when we decide new features and fixes are ready. GitFlow seems like a time and battle tested strategy for organizations that want to have deliberate releases.

Its crazy to me that CodePipeline doesn't have native support for GitFlow which is a multi-branch deploy workflow. The solution provided by the link above suggests using lambda to create CloudFormation stacks. This seems crazy and an abuse of what CloudFormation is designed for. The CloudFormation interface isn't designed for thousands of stacks. Creating a stack for every branch will add an incredible amount of noise to the interface if you use Cfn for standard infrastructure-as-code uses.

Am I missing something?

Further, I'm not sure if there is a language barrier issue in the article, but the following paragraph seems very passive aggressive:

&gt;It’s important to note that trunk-based is, by far, the best strategy for taking full advantage of a DevOps approach; this is the branching strategy that AWS recommends to its customers. On the other hand, many customers like to work with multiple branches and believe it justifies the effort and complexity in dealing with branching merges. This solution is for these customers.

It's as if they are saying, if you're not using trunk based workflow, you're inferior and overly complicating things. When in reality, it feels like AWS is overly complicating the workflow necessary to support a very standard methodology.

Anyone with me on this?  
Are there other GitFlow users out there who use CodePipeline?  
Are there other GitFlow users out there who use something else to release to AWS?
## [3][Failing over with falling over - Stack Overflow Blog](https://www.reddit.com/r/aws/comments/jibnsu/failing_over_with_falling_over_stack_overflow_blog/)
- url: https://stackoverflow.blog/2020/10/23/adrian-cockcroft-aws-failover-chaos-engineering-fault-tolerance-distaster-recovery/
---

## [4][Need SMB shared folder on AWS](https://www.reddit.com/r/aws/comments/jienh1/need_smb_shared_folder_on_aws/)
- url: https://www.reddit.com/r/aws/comments/jienh1/need_smb_shared_folder_on_aws/
---
Hello,

Like the title says, I need to provide a shared folder (SMB) hosted on AWS, it will be accesible from a local LAN. I don't know which option is better: S3 Bucket with AWS Appliance? Amazon FSx? Or I missing something? Help please.
## [5][AWS IAM Introduction](https://www.reddit.com/r/aws/comments/jhu004/aws_iam_introduction/)
- url: https://www.reddit.com/r/aws/comments/jhu004/aws_iam_introduction/
---
Hey everyone - thought I'd post a brief overview of IAM that I initially wrote for work on here. 

Typically IAM is something devs learn "on the way" to building things, but I figure a brief rundown could be useful for a lot of people:

[https://towardsdatascience.com/aws-iam-introduction-20c1f017c43](https://towardsdatascience.com/aws-iam-introduction-20c1f017c43)
## [6][AWS Inspector HELP!!!! Plz....](https://www.reddit.com/r/aws/comments/jieha9/aws_inspector_help_plz/)
- url: https://www.reddit.com/r/aws/comments/jieha9/aws_inspector_help_plz/
---
Greetings community

Does anyone know how Amazon inspector actually works?

Looking at the results for a Linux instance it had Windows CVEs on it and vise versa.

My instances are at the latest patch level but still showing 500+ vulnerabilities?!?

Any help graciously accepted :)
## [7][Rest API on AWS: Will It Bankrupt Me?](https://www.reddit.com/r/aws/comments/jidckk/rest_api_on_aws_will_it_bankrupt_me/)
- url: https://www.reddit.com/r/aws/comments/jidckk/rest_api_on_aws_will_it_bankrupt_me/
---
I'm building a SAAS product. Its not overly complicated but for the sake of brevity on this post:

It will allow users to create projects and within each project create "Diaries". In simple terms its a daily record of events which include time, description + images and video. In return visualisations and reports can be created using the collected data.

I'm considering using AWS Lambdas + API Gateway as the backend. I don't expect each user to generate more than a million requests per month with an entry subscription price of approx. $300 per user (I know it very high but there's very good reason for it  as I've worked in the target market industry for several years).

The simplicity of user authentication and setup of such as a system on AWS is too attractive to ignore as opposed to using Node or Flask for my backend.

Has anyone had experience with such a setup? It seems that a lot of posts here are warning me about the AWS bill but is it really that expensive??
## [8][CDK best practice for CI/CD with stages+config+env](https://www.reddit.com/r/aws/comments/jialp1/cdk_best_practice_for_cicd_with_stagesconfigenv/)
- url: https://www.reddit.com/r/aws/comments/jialp1/cdk_best_practice_for_cicd_with_stagesconfigenv/
---
Hi everyone, I'm new to CDK but I'm used to serverless-framework. Our current approach is to have a git repository for each project and each repository has at least 2 main branches: 

* master for production env on AWS
* staging for staging env on AWS

Every env we have a different configuration (account and region) + specific config for what we are going to deploy, for example, a different domain/subdomain based on the env.

Is there a best practice that allows me to push to a branch, load automatically the config (account+region) + env variables and deploy on that specific AWS env?

I've tried to look for that best practice on the official documentation, forums, and slack channels but to be honest, I never found something official.
## [9][Force Workspaces to only show up on 2 of the 3 monitors](https://www.reddit.com/r/aws/comments/ji5hkh/force_workspaces_to_only_show_up_on_2_of_the_3/)
- url: https://www.reddit.com/r/aws/comments/ji5hkh/force_workspaces_to_only_show_up_on_2_of_the_3/
---
I've been trying to research this, but I cannot find a solution, nor a 3rd party application that can give me what I need. Basically, I need to prevent AWS WorkSpaces from showing up on 1 of my 3 monitors, but I'd like it to be on the other 2. I cannot adjust the display settings within the workspaces application to try and "disconnect" one of the displays, so that isn't going to work. Essentially, the reason is I need 2 of my screens dedicated to running certain work software and my 3rd monitor to be running other software, such as Outlook, Teams, etc.

&amp;#x200B;

Can anyone help me find a solution? I feel like I am all out of ideas and would like to save myself from more pain of figuring this out. Thanks!
## [10][Having trouble setting up load balancers with HIPAA compliance](https://www.reddit.com/r/aws/comments/ji8k18/having_trouble_setting_up_load_balancers_with/)
- url: https://www.reddit.com/r/aws/comments/ji8k18/having_trouble_setting_up_load_balancers_with/
---
Forst of all, I'm pretty new to AWS and I started with this quick start here: https://aws.amazon.com/quickstart/architecture/compliance-hipaa/

The issue I'm having is getting the public proxy server load balancer to redirect to the internal app server load balancer. It just times out. The quickstart used nginx on the proxy server so. I completely replaced the app server EC2 instance so it can work with our app. It's using Django on Ubuntu instead of wordpress on Amazon Linux that the quickstart used. I'm not sure if something is not configured properly with the subnets, security groups or something else. I'm pretty sure I set everything to be open as much as possible but still no luck. Is there a way I can pinpoint the exact cause of the issue? Nginx logs just shows that the connection was closed. Any help or tips would be appreciated!
## [11][Cloud formation manually?](https://www.reddit.com/r/aws/comments/ji6eyn/cloud_formation_manually/)
- url: https://www.reddit.com/r/aws/comments/ji6eyn/cloud_formation_manually/
---
I’m kind of new to AWS, but a long time C++ developer. Trying to find out best practices in terms of development process beyond just using the console for everything. 

How are you all creating your cloud formation templates?  Are you writing the YAML by hand, googling syntax and pasting in stuff. Are you using some sort of cloud formation designer, or are you using the AWS toolkit built into VSCode?
