# aws
## [1][`puppet apply` in EC2 User Data?](https://www.reddit.com/r/aws/comments/exl9fy/puppet_apply_in_ec2_user_data/)
- url: https://www.reddit.com/r/aws/comments/exl9fy/puppet_apply_in_ec2_user_data/
---
Hi,

I have taken over the maintenance of a relatively old CloudFormation project. In the User Data script of EC2 instances (i.e. launch configurations), there's a call to `puppet apply` to create/initialize file system, users/groups, directories, download files from S3, etc.

A few questions:

1. Are there any good reasons to do this instead of using a shell script?
2. I haven't used Puppet in the past but I have 2 years of experience using Ansible.
    1. Back when I was using Ansible, it was mainly inside my Jenkins pipelines, not in my CloudFormation User Data scripts.
3. Does it make any sense to use Ansible/Puppet in User Data at all? I always thought these tools are for remote provisioning of servers, not tools to be used inside User Data scripts.
    1. Now that [ppuppet apply](https://puppet.com/docs/puppet/6.10/architecture.html) has become obsoleted, would you suggest converting my User Data to plain shell scripts or migrating the puppet script to Ansible (in User Data)?

N.B: My only options are open source, free versions of Puppet and Ansible.
## [2][Best way to build a "wish list" notification system?](https://www.reddit.com/r/aws/comments/exdsob/best_way_to_build_a_wish_list_notification_system/)
- url: https://www.reddit.com/r/aws/comments/exdsob/best_way_to_build_a_wish_list_notification_system/
---
I'm playing around with some new stuff and building a little application as part of a learning experience. The app basically has two services at the moment, products and users. Users can sell or buy items, or they can add an item to their wish list. When an item on the wish list drops below a certain price or is back in stock, I want to notify those users that have the item on their wish list.

My data is stored in DynamoDB and I'm trying to make use of only serverless technologies. How can I go about building a solution like this? I want to build it in a "best practice" kind of way, so one of the things that's important is scalability. Therefore, it doesn't really make sense to create an SNS topic for each item being sold, nor does it make sense to go through every single user's wish list each time a product's price/availability changes.
## [3][Visibility of lambda deployment size](https://www.reddit.com/r/aws/comments/exnn8d/visibility_of_lambda_deployment_size/)
- url: https://www.reddit.com/r/aws/comments/exnn8d/visibility_of_lambda_deployment_size/
---
I'm deploying a lambda that uses a number of libraries via layers, and I'm running close to/into the 250Mb deployment limit on Lambda.

As I understand, this is the unzipped size of the layers I'm using, plus the code I deploy directly to the lambda.

I'd like to know if there's any visibility of the deployment size in AWS. At the moment, I get a message when I exceed it, but to work out what's taking the size I'm effectively going to have to unzip all the layers manually on my local PC and add up the various sizes.
## [4][Hosting Spring and a SPA on AWS](https://www.reddit.com/r/aws/comments/exn2k0/hosting_spring_and_a_spa_on_aws/)
- url: https://www.reddit.com/r/aws/comments/exn2k0/hosting_spring_and_a_spa_on_aws/
---
So for the past few days I've been killing myself with hosting my first actual app.. Today, I finally got it fully working and I am so happy. But very soon, I noticed that there's a big problem..

In my React app I am using a BrowserRouter to route around the app. I have my homepage, my about page, and an adminpage. When I tried to access my about page with a direct URL, I got a nice little 404. After googling for a few minutes, it was pretty clear what was wrong.. Since ReactJS is a SPA, routing around the page isn't actually routing around the page, it just unmounts the old and mounts the new Component.. I tried fixing this with the following code in Spring:

    @Controller
    public class Controller {
    	@GetMapping(value = {"/","/about"})
    	public String allIndex(){
    		return "index.html";
    	}
    	@GetMapping("/adminpage")
    	public String getIndex(){
    		return "index.html";
    	}
    }

It... *works*.. But for some reason /adminpage still doesn't want to work, but at least the pages that I want the user to see work, so I'm fine with that for now..

Now I am wondering, would it be better to host the React app on S3 by itself, and host the Spring app on beanstalk? Currently they're both ran from a single .jar and they're run on my beanstalk instance. Would splitting them up like that fix this? How do you deal with this issue? I'm sure the same thing happens with Angular, and all the other SPAs, right? I know React has the HashRouter, but I really don't want to deal with those ugly URLs with the hashtag.

But nonetheless, I probably won't bother fixing this app right now, because I am very happy with it, even if my admin panel doesn't work. So I am wondering how to prevent this from happening in the future. Thank you for any help.
## [5][Scoring A+ for SSL Labs on My Cloudfront-Hosted Static Website](https://www.reddit.com/r/aws/comments/exmsg5/scoring_a_for_ssl_labs_on_my_cloudfronthosted/)
- url: https://adamj.eu/tech/2020/02/02/scoring-a+-for-ssl-labs-on-my-cloudfront-hosted-static-site/
---

## [6][Cloud Predictions 2020: IOD Experts Make Their Calls - [IOD tech Content]](https://www.reddit.com/r/aws/comments/exl7r8/cloud_predictions_2020_iod_experts_make_their/)
- url: https://iamondemand.com/blog/cloud-predictions-2020-iod-experts-make-their-calls/
---

## [7][Beanstalk can't connect to RDS](https://www.reddit.com/r/aws/comments/ex6fht/beanstalk_cant_connect_to_rds/)
- url: https://www.reddit.com/r/aws/comments/ex6fht/beanstalk_cant_connect_to_rds/
---
Hello I have a EC2 instance with a RDS instance and a beanstalk that is hosting my Spring Boot application. When I deployed it, I got a 502 bad gateway.. I went to check the logs and I found that my Spring Boot application can't connect to the DB

    org.postgresql.util.PSQLException: The connection attempt failed.

I haven't configured any security groups or vpc or anything so I'm guessing that might be the issue, but I'm not sure what needs to be changed.. Any help? Thanks

Also when I open my EC2 console, and click on "Running Instances" I only see 1 instance, the beanstalk env, is that normal? I thought that I would find the RDS instance here too..

EDIT: I found the issue.. I connected to the database directly using pgAdmin. As soon as I connected to it, I knew what was wrong.. For some reason, the created database name was just postgresql, instead of the name that I gave it on RDS. After I changed the name of the database everything seems to run perfectly fine.
## [8][How can i automate the process of creating an account with cross-account roles for Terraform?](https://www.reddit.com/r/aws/comments/exez9t/how_can_i_automate_the_process_of_creating_an/)
- url: https://www.reddit.com/r/aws/comments/exez9t/how_can_i_automate_the_process_of_creating_an/
---
I'm trying out AWS's Control Tower service, and it looks like what we need. 

However, i need to have an easy way of adding new 'environment' accounts, that i can access with a cross account role. 

Essentially, in a few clicks, i need to be able to provision a new environment account for a set of infrastructure to live (VPC, ec2, etc) &amp; the roles/policies that allow another account to assume access through a role. It'd be even better if i could do this from the cli, but i need to at least be able to easily do this from the web console. 

I was hoping i could modify Account Factory to do this in all one step, but i have no idea how or where to start.
## [9][EC2 outage in us-east-1?](https://www.reddit.com/r/aws/comments/exh854/ec2_outage_in_useast1/)
- url: https://www.reddit.com/r/aws/comments/exh854/ec2_outage_in_useast1/
---
Curious if anyone else had some major outages at around 1900 EST in us-east-1? I can't find any pattern behind it, it isn't like an entire AZ went out, but simultaneously in multiple VPCs I had instances report failures in status checks, and 2 hours later the ASGs are still struggling to get back to capacity.
## [10][Terraform Vs Ansible - Friends Or Foes? Which Should I Choose?](https://www.reddit.com/r/aws/comments/exmbql/terraform_vs_ansible_friends_or_foes_which_should/)
- url: http://orocke.xyz
---

