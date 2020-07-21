# aws
## [1][Week of July 20th - What are you building this week on AWS?](https://www.reddit.com/r/aws/comments/hutwj1/week_of_july_20th_what_are_you_building_this_week/)
- url: https://www.reddit.com/r/aws/comments/hutwj1/week_of_july_20th_what_are_you_building_this_week/
---
Share your success and challenges
## [2][Where are the AWS Reference Architecture Diagrams?!](https://www.reddit.com/r/aws/comments/hv38ac/where_are_the_aws_reference_architecture_diagrams/)
- url: https://www.reddit.com/r/aws/comments/hv38ac/where_are_the_aws_reference_architecture_diagrams/
---
"No documents found that match the selected criteria." is what I see when looking for *AWS Reference Architecture Diagrams* upon https://aws.amazon.com/architecture

I'm looking for https://aws.amazon.com/architecture/icons/ in the shape of basic Well Architected architecture diagram. You know something showing:

1) Account for IAM
2) Account for Audit/Logs with CloudTrail enabled
3) Maybe a staging, but certainly a production workload account
4) Icing on the cake is Guard duty etc setup in accordance to https://d1.awsstatic.com/whitepapers/Security/AWS_Security_Checklist.pdf

Are there such diagrams out there I can refer to?

Thank you in advance!
## [3][An Overview of Amazon Redshift](https://www.reddit.com/r/aws/comments/husyna/an_overview_of_amazon_redshift/)
- url: https://www.reddit.com/r/aws/comments/husyna/an_overview_of_amazon_redshift/
---
Hi,

I wanted to share a video I put together overviewing Amazon Redshift. I've recently begun using it on a work project and thought it would be helpful to share some of my learnings.

The video is here: https://youtu.be/IAagaCN4CIk

Feedback / Constructive Criticism welcome!
## [4][Advice on deploying Machine Learning model (alternative to Lamda)](https://www.reddit.com/r/aws/comments/hv4cji/advice_on_deploying_machine_learning_model/)
- url: https://www.reddit.com/r/aws/comments/hv4cji/advice_on_deploying_machine_learning_model/
---
Hello everybody,

&amp;#x200B;

I have a Keras machine learning model, that i would like to have deployed as an api on AWS. I have hired an external consultant to perform this task, but i have cloud and ML experience myself. 

The current solution is hosting the model *h5 keras file* on an AWS s3 Bucket with a lambda service loading in the model, and exposing an endpoint to get a prediction. 

Since Lambda's memory is limited to 512 mb by default, the consultant, have now requested for an extension on memory to the AWS support team. 

I liked the Lambda solution, because it is a serverless solution, which is definitely a high priority in comparison to EC2 or similar. 

Though i have a feeling, that extending the default memory on the Lambda, does not seem like a best practice. 

Anybody, that has been in a similar solution?

&amp;#x200B;

Thank you in advance
## [5][How to make sure all of my file uploads are done on s3?](https://www.reddit.com/r/aws/comments/hv3no1/how_to_make_sure_all_of_my_file_uploads_are_done/)
- url: https://www.reddit.com/r/aws/comments/hv3no1/how_to_make_sure_all_of_my_file_uploads_are_done/
---
I have a use case where I am generating many CSVs' and pushing it to  S3 bucket, with meta properties with batch key. example one batch is  executing which has 8 CSVs then all 8 CSVs will have same batch key on  S3 meta properties.

Now here I want to make sure that all the file is getting generated and uploaded on S3. After that I have to send email of all the files as attached in email of single batch created. With the same example above I will be sending email of 8 attachment from that batch with single email as attachments of 8 CSVs.

How Can I make sure that all the CSVs are generated and stored on S3?  Is there anything that can help me to achieve such functionality

Below is my handler function  


    /**
     * handler function
     */
    exports.handler = async (event, context) =&gt; {
    
        var derived = {}
    
        const store = JSON.parse(event);
    
        const { data } = store;
    
        derived.filebatch = filebatchKey.uuid();
    
       
        derived.report_details = await get_details(data.id);
    
    
        const user = derived.user.filters.filter(function(x){return x.origin === "DAILY"});
        
        const teachers = derived.teachers.filters.filter(function(x){return x.origin === "MONTHLY"});
        
        const students = derived.students.filters.filter(function(x){return x.origin === "WEEKLY"});
    
        if(user.length &gt; 0) {
    
          for (const row of user) {
    
                const fileName = `user-report${random.math(4)}`
            
                derived.user = await user(data, row);
                
                // here uploading to S3 with data, file name and batchid that will be stored in meta
    
                await csv_file_upload(derived.user, fileName, derived.filebatch);
          }
    
        }
    
        if(teachers.length &gt; 0) {
    
          for (const row of teachers) {
          
              const fileName = `teachers-report${random.math(4)}`
    
              derived.teachers = await teachers(data, row)	
    
    			// here uploading to S3 with data, file name and batchid that will be stored in meta
              await csv_file_upload(derived.teachers, fileName, derived.filebatch);
    
            }
        } 
    
        if(students.length &gt; 0){
    
          for (const row of students) {
    
              const fileName = `students-report${random.math(4)}`
    
              derived.students = await students(data, row);
    
    		  // here uploading to S3 with data, file name and batchid that will be stored in meta
              await csv_file_upload(derived.students, fileName, derived.filebatch);
    
            }
        }
        
        return { "meta": { "message": "success" }};
    
    };

It could happen that I might have `N` numbers of CSVs that  generate, but the point is I should be getting the last upload or  something like that so that I know that all files are executed and  uploaded to S3 and Now I can send email with that batch attachment.

How can I achieve this? Any help?
## [6][AWS Sumerian with AWS Lex chatbot integration. Getting output 'undefined' for every input text I give it.](https://www.reddit.com/r/aws/comments/hv58j5/aws_sumerian_with_aws_lex_chatbot_integration/)
- url: https://www.reddit.com/r/aws/comments/hv58j5/aws_sumerian_with_aws_lex_chatbot_integration/
---
Hey guys, I am following this documentation:  [https://docs.sumerian.amazonaws.com/tutorials/create/beginner/lex-html/](https://docs.sumerian.amazonaws.com/tutorials/create/beginner/lex-html/) 

And it is working so far. However, sumerian bot is giving me 'undefined' for every input I give it. I am currently using the BookCar template bot with AWS Lex. In the lex console, the bot works fine but when I link it to Sumerian, its  always undefined.

How am I able to view the console logs to pin point the point of failure? Thank you!
## [7][Do I get the free 62,000 emails if I send them from AWS Lambda as well?](https://www.reddit.com/r/aws/comments/hv50k8/do_i_get_the_free_62000_emails_if_i_send_them/)
- url: https://www.reddit.com/r/aws/comments/hv50k8/do_i_get_the_free_62000_emails_if_i_send_them/
---
Or do I \*have\* to use EC2 for that free tier?
## [8][Manage IAM roles and policies at scale](https://www.reddit.com/r/aws/comments/hv4dw8/manage_iam_roles_and_policies_at_scale/)
- url: https://www.reddit.com/r/aws/comments/hv4dw8/manage_iam_roles_and_policies_at_scale/
---
Hi,

I'm building a micro-service which just dumps the data in redshift. Traffic pattern peaks at 12pm and vanish after 5pm and cycle repeats starting 7am. There will be more than 100,000 clients within the whole company!

The way I see it:
1. Create an API - Need to scale it and take care of throttling etc.
2. Put an SQS up front and I can poll at my pace.


If I go with option#2, how do I give permissions to clients to put their data dump in SQS at scale? IAM policies have a total character limit and just won't scale!  Any known work arounds? I don't want to end up in a situation where I'm creating unmanageable IAM roles either!

Thanks!
## [9][Fargate, EKS, and running like a Lambda](https://www.reddit.com/r/aws/comments/hum1f4/fargate_eks_and_running_like_a_lambda/)
- url: https://www.reddit.com/r/aws/comments/hum1f4/fargate_eks_and_running_like_a_lambda/
---
If my questions seem naive, it's because they are.  I'm new to both AWS and Kubernetes, having been tossed into the blender recently on the job.

I need to write a new... service, X.  It will be called sporadically (i.e., on-demand, not scheduled) by another service, Y.  It will run a body of code that we don't own, and will not modify, so that code will be run as-is, essentially by fork &amp; exec.  The results of that run will be communicated by X to Y via Kafka topic.

We will be using Kubernetes (EKS), in which service Y will run as a traditional service.

I'm being strongly urged to use Fargate for the new "service" X, although we have no expertise with it.  The thinking relayed to me is that we'd like this to act like a Lambda, but as a job for X may take hours to complete, Lambda can't be used.  Also, I'm told that we believe it will be cheaper to use Fargate for this than a traditional "always running" service for X, though I disbelieve we can know that, since Fargate costs more per minute than ECS or EKS (right?) and we have a poor understanding of how frequently X will be used.

My vague understanding of Fargate is that it's not truly "serverless" like Lambda, in the sense that it's not a function you're building with it, it's a Docker container, with arbitrary functionality?

But I'm being told that an application built for Fargate *can* behave like a function, in that you trigger it, the app gets provisioned, runs, and then exits and is torn down?

I really need someone to help me understand how this could work.  If my new service X were a traditional, always-hot service, I'd write it to take HTTP REST requests, or subscribe to a Kafka topic for requests.  How, if I'm writing X to behave like a "take one request, then exit" "service", would / could Y communicate with it?

Would I write X to take HTTP / Kafka requests, and then just exit after completing one request?

Or is there a pattern here I should use, like taking command arguments to X when it starts?

If the latter, how am I to trigger X?  Is there a way in AWS for me to place a REST API in front of X, which could translate the request into a form that invoking X with arguments would want, and then invoke X?  (I'm not yet understanding the mechanism that triggers a Fargate thing to run.)

As I said, naive questions, thanks in advance for any help.  AWS docs are voluminous, but not much help for my level of newbie-ness.
## [10][Server Migration from VMware to AWS troubleshooting (Storage not showing)](https://www.reddit.com/r/aws/comments/huyyy3/server_migration_from_vmware_to_aws/)
- url: https://www.reddit.com/r/aws/comments/huyyy3/server_migration_from_vmware_to_aws/
---
First time i've had this issue and its happen with multiple servers today (Windows). I am migrating VMs from Vmware to AWS. When i get to the storage configuration it is not showing any drives. Normally i see the drives of the original servers. 

Has anyone seen this before? i have done the import a couple of times today.
## [11][Seeking information on getting started selling my AWS Creations](https://www.reddit.com/r/aws/comments/hv3bto/seeking_information_on_getting_started_selling_my/)
- url: https://www.reddit.com/r/aws/comments/hv3bto/seeking_information_on_getting_started_selling_my/
---
I'm looking to create a side business that runs AWS services for my client. For example, I might have a custom Lambda set up that performs an action with their API software (outside of AWS).

Are there any resources out there that show how I can set this up as a SaaS for billing purposes?

Thanks
