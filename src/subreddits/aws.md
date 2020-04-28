# aws
## [1][Now Open – AWS Europe (Milan) Region](https://www.reddit.com/r/aws/comments/g9i7wc/now_open_aws_europe_milan_region/)
- url: https://aws.amazon.com/blogs/aws/now-open-aws-europe-milan-region/
---

## [2][Free AWS Workspaces until 6/30/2020](https://www.reddit.com/r/aws/comments/g9bkev/free_aws_workspaces_until_6302020/)
- url: https://www.reddit.com/r/aws/comments/g9bkev/free_aws_workspaces_until_6302020/
---
Apparently AWS announced that you can get up to 50 workspaces for free until 6/30/2020.

[https://aws.amazon.com/blogs/desktop-and-application-streaming/work-from-home-offer-for-amazon-workspaces](https://aws.amazon.com/blogs/desktop-and-application-streaming/work-from-home-offer-for-amazon-workspaces)
## [3][How to stop the running instance using AWS CLI?](https://www.reddit.com/r/aws/comments/g9liub/how_to_stop_the_running_instance_using_aws_cli/)
- url: https://www.reddit.com/r/aws/comments/g9liub/how_to_stop_the_running_instance_using_aws_cli/
---
Hello there,

I have created a shell script to stop the running instances using AWS CLI - 

    InstanceId=( $(aws ec2 describe-instances --filters "Name=tag:Name,Values=main-instance" "Name=tag:Environment,Values=dev" "Name=instance-state-name,Values=running" --output text --query 'Reservations[*].Instances[*].InstanceId') )
    
    aws ec2 stop-instances --instance-ids ${InstanceId}

It works fine except when console shows some older terminated instances with name tag main-instance. stop-instance command fails while trying to stop them with followin error message - 

 `An error occurred (IncorrectInstanceState) when calling the StopInstances operation: This instance 'i-9djd30j93930d3d33' is not in a state from which it can be stopped.` 

Instance `i-9djd30j93930d3d33` is already terminated, while my script doesnt do anything to the running instance.

Anything wrong I did up there?

Thanks
## [4][Ssh connection gets timed out after installing apache2 , though Http port is accessible.](https://www.reddit.com/r/aws/comments/g9mt1l/ssh_connection_gets_timed_out_after_installing/)
- url: https://www.reddit.com/r/aws/comments/g9mt1l/ssh_connection_gets_timed_out_after_installing/
---
I tried to add custom domain to aws ec2 instance by editing Route 53 setting and installed apache2 server. Now I am unable to ssh into the instance, though I can access the apache server via http: The ip4 address is - [3.7.45.198](https://3.7.45.198).  I only edited inbound rule to open HTTP port and did not touch already opened ssh port 22.

&amp;#x200B;

The error is: The connection has been closed because the server is taking too long to respond..
## [5][SageMaker: do not lose your packages (and conda envs) after machine restarts](https://www.reddit.com/r/aws/comments/g94hn2/sagemaker_do_not_lose_your_packages_and_conda/)
- url: https://www.reddit.com/r/aws/comments/g94hn2/sagemaker_do_not_lose_your_packages_and_conda/
---
Hey all,

I really like using AWS SageMaker as dev machine, but losing installed packages and conda environments after machine restarts is very annoying. 

Here is the guide to prevent that: [https://biasandvariance.com/sagemaker-save-your-conda-environments/](https://biasandvariance.com/sagemaker-save-your-conda-environments/)

Hope it's useful!
## [6][[Question] Using Global Accelerator with endpoint port mappings](https://www.reddit.com/r/aws/comments/g9jtpy/question_using_global_accelerator_with_endpoint/)
- url: https://www.reddit.com/r/aws/comments/g9jtpy/question_using_global_accelerator_with_endpoint/
---
When deploying a Global Accelerator instance,  
is there a way to specify which port will the traffic be forwarded to on the endpoint (EC2 instance)

I know i can add a loadbalancer as a middleman but I need tcp connections with client ip preservation.  


My end goal is to have tcp traffic from port 443 on the global accelerator reach and EC2 instance on port 4443.
## [7][AWS Lambda Not Consistent Across Browsers?](https://www.reddit.com/r/aws/comments/g9jq16/aws_lambda_not_consistent_across_browsers/)
- url: https://www.reddit.com/r/aws/comments/g9jq16/aws_lambda_not_consistent_across_browsers/
---
Not sure if anyone else is having this issue, but I created an AWS Lambda project built off of the space facts template, and added two custom functions. When I open the project via the invocation name while on chrome, it works fine consistently. When I try to invoke either of my two other commands, I *sometimes* get a response, and most of the time get "There was an issue processing your input."

Tried doing the exact same thing in the testing console in firefox and it worked fine! The only thing I could think of that would cause the inconsistency is some kind of firewall or antivirus exception for one of the browsers, but I don't have any exceptions in my antivirus software and my firewall is disabled.

If anyone has any insight on this I'd be interested in hearing what you guys have to say!


Edit: Now it's not working on firefox either! I have no idea what's going on.
## [8][How to share security groups between AWS regions?](https://www.reddit.com/r/aws/comments/g9grre/how_to_share_security_groups_between_aws_regions/)
- url: https://www.reddit.com/r/aws/comments/g9grre/how_to_share_security_groups_between_aws_regions/
---
I have servers in multiple AWS regions, and they all use the identical security group. Any time I need to make an adjustment, I need to go to each of the regions and add or remove an IP address. 

How can I share a single security group across regions?
## [9][AWS Storage](https://www.reddit.com/r/aws/comments/g9errv/aws_storage/)
- url: https://www.reddit.com/r/aws/comments/g9errv/aws_storage/
---
Hello everyone, i deal with many data, documents, photos, videos and so on.

So i saw that AWS has cheap rates (despite not knowing wich plan to get) to bulk storage. I want to ask your help, what plan should i get, is there any storage with Ubuntu or windows so that is easier to manage?

I plan to move close to 5Tb right now, and i will download some files once or twice a month, but manly upload constantly, like a backup.
## [10][Django HelloWorld Changes not showing up.](https://www.reddit.com/r/aws/comments/g9g4kw/django_helloworld_changes_not_showing_up/)
- url: https://www.reddit.com/r/aws/comments/g9g4kw/django_helloworld_changes_not_showing_up/
---
I created a HelloWorld starter app.

I pulled the code into my machine and started making changes, I can see the changes locally but not one AWS.

I can see my changes in CodeCommit's master branch. Project is builidng successfully.

What should I try?

&amp;#x200B;

Thank you.
