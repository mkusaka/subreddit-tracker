# aws
## [1][Week of Aug 10th - What AWS questions do you have?](https://www.reddit.com/r/aws/comments/i75xfw/week_of_aug_10th_what_aws_questions_do_you_have/)
- url: https://www.reddit.com/r/aws/comments/i75xfw/week_of_aug_10th_what_aws_questions_do_you_have/
---
What questions do you have about AWS?
## [2][The new Route53 console isn't great](https://www.reddit.com/r/aws/comments/i84o7w/the_new_route53_console_isnt_great/)
- url: https://www.reddit.com/r/aws/comments/i84o7w/the_new_route53_console_isnt_great/
---
Today I logged in to route53 to see it has a new console much like the rest of AWS, I didnt think much of it until I started to use it. 

Creating new records is now a 3 page wizard whereas it used to be a simple sidepanel where you could create a record fast and move on. 

Honestly im finding some of the new consoles to be similar where there is less detail and more wizard driven menu's...  its not great.
## [3][STS role chaining with Terraform?](https://www.reddit.com/r/aws/comments/i8bz6w/sts_role_chaining_with_terraform/)
- url: https://www.reddit.com/r/aws/comments/i8bz6w/sts_role_chaining_with_terraform/
---
Hello,

I have the following setup: 

* **userA** can assume the role **roleA** in **accountA**
* **accountB** has a role **roleB** which can be assumed by any user of accountA having the role roleA

What I want to achieve with Terraform, which works with the CLI, is to have my userA assume the roleA in accountA, then using the newly generated STS credentials, assume roleB in accountB in another Terraform module.

I tried doing it this way:

    # main.tf
    provider "aws" {
      # userA credentials are exported as environment variables
      version = "~&gt; 2.0"
      assume_role {
        role_arn     = "arn:aws:iam::accountA:role/roleA"
      }
    }
    
    module "module" {
      source = "./mymodule"
    }
    

&amp;#x200B;

    # mymodule/main.tf
    provider "aws" {
      version = "~&gt; 2.0"
      assume_role {
        role_arn     = "arn:aws:iam::accountB:role/roleB"
      }
    }

But this does not seem to work: the main module properly assumes roleA, but it doesn't "keep the roleA context" when calling the submodule, and what happens in there is that terraform tries to assume roleB as userA, which is not what's intended.

Is there any way to perform this role chaining with terraform?

Thank you
## [4][How could ASG assign EIP ?](https://www.reddit.com/r/aws/comments/i8ds97/how_could_asg_assign_eip/)
- url: https://www.reddit.com/r/aws/comments/i8ds97/how_could_asg_assign_eip/
---
I'm baffled there doesn't seem to be a trivial way of doing this. I was just charged a few dollars because my EIP wasn't allocated after an EC2 instance was restarted...

How can we ensure an EIP gets assigned by the ASG which starts a new EC2 ?

User Data seems to be the best solution right now, but I don't really like it (and requires inputting AWS\_SECRET):  [https://forums.aws.amazon.com/message.jspa?messageID=498761](https://forums.aws.amazon.com/message.jspa?messageID=498761)
## [5][What is the best tool for creating visual diagrams of existing AWS infra?](https://www.reddit.com/r/aws/comments/i7tu8p/what_is_the_best_tool_for_creating_visual/)
- url: https://www.reddit.com/r/aws/comments/i7tu8p/what_is_the_best_tool_for_creating_visual/
---
I have an existing AWS deployment of mostly EC2 instance with some S3 and a decent amount of advanced networking pieces (transit gateways, direct connect, etc.) I am looking for the best tool to diagram this out. I have used CloudCraft in the past but wondered if there was anything else I should consider.
## [6][How good is AWS ECS in 2020?](https://www.reddit.com/r/aws/comments/i8ac24/how_good_is_aws_ecs_in_2020/)
- url: https://www.reddit.com/r/aws/comments/i8ac24/how_good_is_aws_ecs_in_2020/
---
It seems to have some powerful features, while still being fairly simple. However, our ECS set-up is small and the ECS set-up yet and I'm actually doing most stuff manually.

For a very long time Docker orchestration was very painful and while it solved 30 issues and made many things simpler, it introduced also another 10 problems and added some additional complexities. I'm re-navigating the Docker orchestration landscape and ECS is starting to look fairly good. It's even compatible with docker-compose, which has very simple  clean yaml configuration files and an extremely low learning curve. 

What is your experience with ECS in 2020? 

If you know a good articles instead of personal experiences/reviews, please do post them in the comments! :)
## [7][ECS Autoscale EC2 Instances](https://www.reddit.com/r/aws/comments/i8cwrn/ecs_autoscale_ec2_instances/)
- url: https://www.reddit.com/r/aws/comments/i8cwrn/ecs_autoscale_ec2_instances/
---
Hello,

For simplification, I currently have a cluster running 1 t2.medium instance that has one container that consumes all of the containers resources. I'm trying to do a rolling update, where I can update the container, run the task and once ready shuts down the old container. 

It should also be noted that I have added a placement constraint on the task "ec2Instance = t2.medium". 

The problem, is when I try to force a new deployment, it says "insufficient memory". Now, if I go to cloud formation and create a new instance, it works with no problem. 

However, I'm trying to automate this using ASG to spin up a new instance when needed and then drain the old one gracefully. I know this has to be possible, but the ECS autoscaling group seems like it only pertains to tasks and not ec2 instances. 

Can anyone point me in the right direction?

&amp;#x200B;

Thanks
## [8][should I use roles when developing locally?](https://www.reddit.com/r/aws/comments/i8ck7s/should_i_use_roles_when_developing_locally/)
- url: https://www.reddit.com/r/aws/comments/i8ck7s/should_i_use_roles_when_developing_locally/
---
I'm developing locally an app that will eventually run on EC2. The app needs IAM credentials to access an AWS database and other resources. AWS docs recommend [using roles for apps running on EC2](https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html#use-roles-with-ec2). When developing locally, is it worth [configuring local profile to use a role](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-role.html), or is it ok practice to attach a bunch of policies directly to IAM user with an access key?
## [9][Updating GuardDuty via API woes](https://www.reddit.com/r/aws/comments/i8b5sv/updating_guardduty_via_api_woes/)
- url: https://www.reddit.com/r/aws/comments/i8b5sv/updating_guardduty_via_api_woes/
---
GuardDuty seems pretty decent, I like it. However I am constantly running into issues configuring it via API and am wondering if I'm doing something wrong or whether anyone knows of work arounds.

The setup:

- Org is managed via an unholy mix of Control Tower and our own in-house org management tooling that we created because even ALZ wasn't a thing at the time. I'm consolidating everything into CT + the CT customisations (https://aws.amazon.com/solutions/implementations/customizations-for-aws-control-tower/).

- GuardDuty is not in Control Tower. Our previous in house solution used Stack Sets to deploy things to every account in every region, this was before the "organisation wide" functionality of GuardDuty existed.

- I am now using the Org GuardDuty features, though not deploying via CT customisation because it only needs to touch two accounts: 1) The org master account to delegate the CT audit account as the GuardDuty master via a StackSet deployed to org master in every region that has a Custom Resource to execute the API call, 2) The audit account to again run Custom Resources to activate GuardDuty as org wide in every region via Stack Set and also to create IP Set and Threat Intel lists.

Issues I have run into:

- When designating a GuardDuty master from the Org master and activating the GuardDuty master as being organisation wide it implicitly creates the Detectors in all accounts/regions. This is nice. However during the initial deployment some of these detectors got IDs of the form `REDACTED-CD1e-REDACTED--D5c54-5A89B` (as opposed to `fC51REDACTEDDEbCB2REDACTEDDadCE`, note the hyphens). Whilst these detectors could get listed via the API, other calls, such as even getting or deleting it, would say either that ID does not exist or the ID is malformed. After a weekend this had gone away and I was then able to delete and re-create the detectors.

- I am now experimenting with the S3 detections. Because the detectors previously existed they either needed to be re-created to get this feature or updated to opt-in. I chose to re-create because I wanted the experience of doing it from scratch (i.e. if a customer asked us for it, we had a DR situation, etc). Disabling the org wide functionality and then removing the designation of a GuardDuty master worked fine, however the detectors still remained in the accounts (without the S3 feature opt-in). I chose to write a little script to purge them, however after 24 hours they are still not gone. I can still make the API call to delete and it succeeds but they remain and appear functional. I can successfully delete them via the console, but a) I'm not much of a point-and-click guy, b) I need to know what the problem is.

- When creating Threat Intel IP Sets the API it will often give a permission error (nothing else useful in CloudTrail) but when re-running the command manually it will then work. I am choosing to use the API in a Custom Resource to do this because I want the list of IP sets to be "dynamic" via parameters and so can't just specify them as distinct resources in CloudFormation.

I would lodge a ticket, however we don't have support on all accounts... and getting this working is in fact the last step before finishing the migration of accounts to our Control Tower setup.
## [10][Is there anything besides the AWS Free Tier to play around with on AWS?](https://www.reddit.com/r/aws/comments/i8a5jw/is_there_anything_besides_the_aws_free_tier_to/)
- url: https://www.reddit.com/r/aws/comments/i8a5jw/is_there_anything_besides_the_aws_free_tier_to/
---
I've been using AWS for 5 years, but I want to extensively play around with it.

But AWS is expensive and I don't want to personally pay a lot.

Does AWS have any plans for developers where they can extensively use the AWS Cloud for educational purposes without spending too much money?
## [11][Hands-on workshop to create mini Data Lake with Amazon S3, Lake Formation and Glue](https://www.reddit.com/r/aws/comments/i89prr/handson_workshop_to_create_mini_data_lake_with/)
- url: http://aws-dojo.com/workshoplists/workshoplist3
---

