# aws
## [1][Best Practice Guide for Setting a Company Up on AWS For The First Time](https://www.reddit.com/r/aws/comments/ggebgn/best_practice_guide_for_setting_a_company_up_on/)
- url: https://www.reddit.com/r/aws/comments/ggebgn/best_practice_guide_for_setting_a_company_up_on/
---
I am a senior AWS DevOps engineer and have just gotten a contract with a company to move their private cloud infrastructure to AWS. This company has a DevOps and TechSecOps team but no AWS experience. Aside from the migration itself, can anyone send me a link to a webpage, book, or other, that details a step-by-step best practices of setting a company up on AWS please? This company is ISO-27001 security compliant. I am thinking of stuff like this:

\- Create AWS Organizations, ie have a parent account and a couple of initial child accounts

\- Create a couple of IAM users

\- Turn on Cloud Trail

\- Configure AWS Config

\- Configure Billing Alerts

\- Maybe configure VPC endpoints?

\- Create a few Cost Allocation Tags

\- Create a standard VPC configuration with public and private subnets 

\- Anything else I have missed

&amp;#x200B;

And finally getting sample Terraform code to do all the above. Does anyone know of a comprehensive resource with at least all the above information as well as other basic stuff? If so, I'd be grateful for any information/suggestions/watch outs etc.

Thanks.
## [2][EC2 Price Reduction – For EC2 Instance Saving Plans and Standard Reserved Instances](https://www.reddit.com/r/aws/comments/gfznd6/ec2_price_reduction_for_ec2_instance_saving_plans/)
- url: https://aws.amazon.com/blogs/aws/ec2-price-reduction-for-ec2-instance-saving-plans-and-standard-reserved-instances/
---

## [3][Alternative to ECS Container Insights?](https://www.reddit.com/r/aws/comments/ggegmc/alternative_to_ecs_container_insights/)
- url: https://www.reddit.com/r/aws/comments/ggegmc/alternative_to_ecs_container_insights/
---
Anyone using something else to get some insights into performance / running count etc then Container Insights? The pricing for this is relatively high, and by the looks of it is 10% of our spending with AWS.

Any suggestions?
## [4][AWS for practice on a budget](https://www.reddit.com/r/aws/comments/ggc8fe/aws_for_practice_on_a_budget/)
- url: https://www.reddit.com/r/aws/comments/ggc8fe/aws_for_practice_on_a_budget/
---
I want to try out AWS, mostly for practice so it will be used by myself and will probably have no traffic. I only have a debit card so I was wondering if its possible to minimize expenditure or set some kind of limit on my account so that the budget is not exceeded
## [5][Best searchable DB for car data](https://www.reddit.com/r/aws/comments/ggb6al/best_searchable_db_for_car_data/)
- url: https://www.reddit.com/r/aws/comments/ggb6al/best_searchable_db_for_car_data/
---
Hello,

For a project I've been tasked to store data of cars for dealerships. The model is this data is something as:

{

Vehicle Identification Number ,Color,Brand,type,dealership,...

}

From this data, only the dealership field may change over time.

The user should be able to text search in any field, get all yellow cars, get all Ford cars,...

In the current prototype, I'm using dynamoDB. At this point, all the data (100 entries) can still be fetched by a combination of GSI s and sorted in the front end.

However, once this application is launched, the data sets might reach 10.000-100.000 . At this point, DynamoDB queries would require a GSI for each field, or per dealership have a LSI per field. The query limit of 1MB might be

After some googling, I've found that some combine dynamoDB and elastisearch. Why not use only elastisearch then? Is a SQL DB like Aurora better suited for this?

All advice is welcome
## [6][Suggestions on what to do with credits that are about to expire](https://www.reddit.com/r/aws/comments/gfu8f6/suggestions_on_what_to_do_with_credits_that_are/)
- url: https://www.reddit.com/r/aws/comments/gfu8f6/suggestions_on_what_to_do_with_credits_that_are/
---
I have about $500 in credits that will expire by the end of this month. I have been trying to think of something interesting to do with it before it goes away.

Any suggestions of interesting short-term projects that I could use this credits with ?

Thanks.
## [7][Is anyone else having troubles paying AWS with a check during the pandemic?](https://www.reddit.com/r/aws/comments/gg29vb/is_anyone_else_having_troubles_paying_aws_with_a/)
- url: /r/startup/comments/gg1xjw/is_anyone_else_having_troubles_paying_aws_with_a/
---

## [8][Elastic beanstalk with 404 API not found](https://www.reddit.com/r/aws/comments/gg2979/elastic_beanstalk_with_404_api_not_found/)
- url: https://www.reddit.com/r/aws/comments/gg2979/elastic_beanstalk_with_404_api_not_found/
---
Hi- I'm relatively new to AWS. My elastic beanstalk server just stopped working despite server being health listed as ok. Nothing was changed and the app has been the same and deployed since 2018. I have re-deployed several times without success. Now when I try to load the application URL I get

{"status":404,"message":"Api not found"}  

&amp;#x200B;

Here is the elastic beanstalk console 

&amp;#x200B;

Any tips appreciated! Thanks

https://preview.redd.it/hf06efw21mx41.png?width=1520&amp;format=png&amp;auto=webp&amp;s=4e8313400238c0d5cbb29c9c81fdc48d7c645345
## [9][Which AWS Service should I use for processing (like batch) something with consuming Rest API that will be triggering from another AWS Service?](https://www.reddit.com/r/aws/comments/gg1tp4/which_aws_service_should_i_use_for_processing/)
- url: https://www.reddit.com/r/aws/comments/gg1tp4/which_aws_service_should_i_use_for_processing/
---
Which AWS Service should I use for processing (like batch) something with consuming Rest API that will be triggering from another AWS Service? In fact, My first concern is to get a minimum cost to achieve that.

Short Explanation: think about there is an app and the app will spin up with some conditions to consume a Rest API. The app will be triggered in some cases from another AWS Service that might host a simple Java app, But that Java app may trigger this consuming app multiple times before this consuming app don't end up previous processing job/task. So at this point, There might be consuming app more than one simultaneously and I think the consuming job/task will take max 3 hours in one shot.
## [10][AWS Beanstalk and Laravel - Mac vs Windows Upload](https://www.reddit.com/r/aws/comments/gg15u2/aws_beanstalk_and_laravel_mac_vs_windows_upload/)
- url: https://www.reddit.com/r/aws/comments/gg15u2/aws_beanstalk_and_laravel_mac_vs_windows_upload/
---
I'm still new to AWS, and I think I might be having file permission issues on Beanstalk.

I've seen that file permissions are automatically set by Beanstalk. When you upload a new zip with new directories and files in it, are those permissions run again to apply to the new files?

I'm working with an existing Beanstalk app that was previously uploaded from a Mac, but when I upload my Windows zip with new files, they aren't found when I load the Laravel application in the browser.

Initially, the new pages files were loading fine. Certain new images weren't rendering on the new pages that I created on Windows. So it was finding the page files fine, but not the images.

Now, I pointed the images to an S3 bucket and it won't even load the page file when I upload a new zip. Laravel responds with an error saying it cannot find the view file, which it was able to find previously.

When I download the uploaded zip from AWS, the new files and images do exist there, so they are being processed. 

I don't think it matters, but there aren't any .gitignore files that would cause it to ignore the new files.

Does this sound like a permissions issue from Windows, and is there a way to refresh the permissions without connecting to the EC2 via SSH? Can I include a config file of some way to force the permissions to set correctly every time a new zip is deployed?

I'm working on getting a Linux machine running so I can develop on that instead of Windows, but I'm trying to understand what's happening.
