# aws
## [1][Awesome AWS workshops](https://www.reddit.com/r/aws/comments/i2wcl2/awesome_aws_workshops/)
- url: https://awesome-aws-workshops.com/
---

## [2][If I disable mutli-az on a running MySQL rds, is there downtime?](https://www.reddit.com/r/aws/comments/i2t647/if_i_disable_mutliaz_on_a_running_mysql_rds_is/)
- url: https://www.reddit.com/r/aws/comments/i2t647/if_i_disable_mutliaz_on_a_running_mysql_rds_is/
---
Willl currently running transactions be affected?
## [3][How to recover a suspended account?](https://www.reddit.com/r/aws/comments/i2ujwk/how_to_recover_a_suspended_account/)
- url: https://www.reddit.com/r/aws/comments/i2ujwk/how_to_recover_a_suspended_account/
---
I am in a very strange situation with my account apparently having been suspended.   
I created an aws personal account over a year ago and have always been within the free tier limits.

I use it seldomly as I mostly utilise it for proof of concept apps. After some time away from it, I logged in today only to be greeted with the message: \`Authentication failed because your account has been suspended.\`

I have nothing running and I did not exceed any limits. I use my real contact information and have no outstanding balance to be paid. I was able to reach the account details area and saw that the debit card which I used had expired. 

So I updated it to my new one. I received a charge of 1$ but my account remains suspended.

Looking into the support section, I see that the sole suggestion for cases of a suspended account is to pay any outstanding balance which is not my case.  
In addition to that, in order to receive technical support , I need to Sign in &amp; Submit AWS Support Request, which is not possible for me as I get the \`account suspended\` message.  


I don't even know why my account was suspended in the first place, nor how I can find out this information, and even if I knew, I have no idea who to contact or what the next steps are to resolve this issue.  
Has anyone else experienced anything similar?
## [4][ACM additional verification in a huge organization](https://www.reddit.com/r/aws/comments/i2szqi/acm_additional_verification_in_a_huge_organization/)
- url: https://www.reddit.com/r/aws/comments/i2szqi/acm_additional_verification_in_a_huge_organization/
---
Hey,

we've got a pretty huge organization (200+ accounts and still in early migration, will increase much more). Two of our main domains are listed in the Alexa top 500, so we need additional verification for each account to be able to generate ACM certificates. So far this is a super annoying manual process, where you first have to register for the forum with that account, wait a few hours, then make a public post, and then wait again for some unpredictable amount of time when someone unlocks that account. Rinse and repeat hundreds of times.

Has anyone managed to tie this authorization to the organizational root? I've asked our rep about this, but they're currently on vacation with an auto-responder.
## [5][How to run scheduled job (e.g. midnight) that scales depending on needs?](https://www.reddit.com/r/aws/comments/i2j905/how_to_run_scheduled_job_eg_midnight_that_scales/)
- url: https://www.reddit.com/r/aws/comments/i2j905/how_to_run_scheduled_job_eg_midnight_that_scales/
---
I want to run scheduled job (e.g. once a day, or once a month) that will perform some operation (e.g. deactivate those users who are not paying, or generate reminder email to those who are due payment more than few days).

The amount of work each time can vary (it can be few users to process or few hundred thousands). Depending on the amount of data to process, I want to benefit from lambda auto scalability.

Because sometimes there can be huge amount of data, I can't process it in the single scheduled lambda. The only architecture that comes to my mind is to have a single "main" lambda (aka the scheduler) and SQS, and multiple worker lambdas.

The scheduler reads the DB, and finds all users that needs to be processed (e.g. 100k users). Then the scheduler puts 100k messages to SQS (separate message for each user) and worker lambdas are being triggered to process it.

I see following drawbacks here:

- the scheduler is obvious bottleneck and single point of failure
- the infrastructure contains of 3 elements (scheduler, sqs, workers)

Is this approach correct? Is there any other simpler way that I'm not aware of?
## [6][Aurora Question](https://www.reddit.com/r/aws/comments/i2u4au/aurora_question/)
- url: https://www.reddit.com/r/aws/comments/i2u4au/aurora_question/
---
I am relatively experienced with many AWS services - but I do have a large gap around Aurora/RDS

&amp;#x200B;

I'm trying to create a multi-region multi-master (write replicas) setup

The purpose is to give low latency to users (if each read and write replica is in the user's region) and to give resilience (if there is a region outage, the users can have their requests routed to another region (the latency will be higher, but reduced service is better than no service))

&amp;#x200B;

I'm trying to learn about AWS Aurora and I've created a toy cluster to learn.  It seems I can create a cluster that is served out of multiple regions (and Aurora replicates data between regions automatically).  I've also read that it is possible to have a multi-master setup (in my toy cluster, it only had one write partition, I couldn't work out how to create another write partition in another region, which made me question if it's possible?)

&amp;#x200B;

Here is a diagram of what I'm thinking:

 https://imgur.com/DzoSpHL 

&amp;#x200B;

Thank you in advance!

&amp;#x200B;

tl;dr:

multi-master over multi-region Aurora - possible?
## [7][Looking for CloudFormation and RDS Aurora serverless best practices](https://www.reddit.com/r/aws/comments/i2w6su/looking_for_cloudformation_and_rds_aurora/)
- url: https://www.reddit.com/r/aws/comments/i2w6su/looking_for_cloudformation_and_rds_aurora/
---
Hi everyone,

Updating an RDS resource in a CloudFormation template *(e.g EngineVersion property)* may result in the replacement of the current DB instance by a fresh one, thus loosing data. To avoid this, using **snapshots** seems to be the only solution *(by passing a* ***SnapshotId*** *to the CF template)*.

However, I have a lot of other resources accessing the RDS instance through its **ARN** which is exported in the CloudFormation template **Outputs**. So having a new RDS instance would break my other resources.

\- Is the initialization from **snapshots** the way to go ?

\- Would storing the **ARN** in the **Parameter Store** be a correct solution ?

\- Is there something I'm missing regarding the automation of this kind of db maintenance ?

Thanks in advance !
## [8][[ECS] Running multiple services from same ECR](https://www.reddit.com/r/aws/comments/i2vssy/ecs_running_multiple_services_from_same_ecr/)
- url: https://www.reddit.com/r/aws/comments/i2vssy/ecs_running_multiple_services_from_same_ecr/
---
Hello,

my usecase is that I have a single code base which is built by Jenkins into single ECR image. From that, I want to have multiple Services, that will run from the same ECR but have different environment variables. Each of the Services needs to be independently scalable.

Example: I have ECR MY\_CONTAINER and based on it I want to have the following

MY\_CONTAINER -e PRESET=A, (min, max) = (0, 10)  
MY\_CONTAINER -e PRESET=B, (min, max) = (5, 100)

I would have Lambda that would, based on the external event, spawn new tasks with either preset = A or preset=B. 

&amp;#x200B;

I'm wondering for the cleanest way to do this. My preset is far more complex than a single variable, and I'd prefer to have these somewhat "documentable" (in an IaaC fashion maybe).  

&amp;#x200B;

One idea that comes to mind is to keep `my-container-task-def-preset-A.json` and `my-container-task-def-preset-B.json` in the repository, and then Jenkins could update everything at build phase, but this solution looks somewhat clunky to me.

Due to my lack of knowledge, I'm not sure if the following is possible.

I have two ECS services, i.e. `my-container-service-A` and `my-container-service-B` which are spawned from the same task def and have empty env. These services have separate scaling rules. Then, on the Lambda side, I add ENV on the Lambda side and send the task to the appropriate service. This way I could keep all preset mappings in one JSON, that can be used as python mapping and sort of documentation. 

&amp;#x200B;

On the same note, if I submit a task to ECS, how long it can stay in PENDING state until its execution is cancelled? Is there a way to increase this quota to inf?
## [9][Can RDS read replicas function as a backup?](https://www.reddit.com/r/aws/comments/i2uwh7/can_rds_read_replicas_function_as_a_backup/)
- url: https://www.reddit.com/r/aws/comments/i2uwh7/can_rds_read_replicas_function_as_a_backup/
---
I'm creating an RDS database and I am wondering if a read replica can function as a backup DB. I have read the AWS page regarding read replicas and they mention that the replica can be promoted to be a standalone database. Does this mean that the replica can be a backup?
## [10][Lambda zip file format](https://www.reddit.com/r/aws/comments/i2ugym/lambda_zip_file_format/)
- url: https://www.reddit.com/r/aws/comments/i2ugym/lambda_zip_file_format/
---
I found I have to use the zip upload method if using npm libraries with my lambda handler. Right now my directory structure is like this:

    node_modules // the library folder
    lambda
      handler.js  

I'm guessing this is not the right structure. Can someone advise on how to set it up for upload?

I'm curious about both the proper file/folder structure, and also if `node_modules` should be included at all.

Usually the library folder is not included, but I sense it may need to be included here (if AWS doesn't handle the install step itself).
