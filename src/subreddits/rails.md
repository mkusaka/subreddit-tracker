# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/i00rha/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/i00rha/personal_projects_show_off_your_own_project_andor/
---
In this thread you can showcase your personal pet project to other redditors.

Need help with a specific problem or just wanna have some extra eyeballs on your code? Ask away!

A suggested format to get you started:

1. **Name of your project**
2. **A short description**
3. **Application stack**
4. **Link to Live app**
5. **Link to GitHub**
6. **You experience level**
7. **Other information or areas that you would like advice on**

 

^(Many thanks to Kritnc for getting the ball rolling.)
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/i8dsvv/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/i8dsvv/personal_projects_show_off_your_own_project_andor/
---
In this thread you can showcase your personal pet project to other redditors.

Need help with a specific problem or just wanna have some extra eyeballs on your code? Ask away!

A suggested format to get you started:

1. **Name of your project**
2. **A short description**
3. **Application stack**
4. **Link to Live app**
5. **Link to GitHub**
6. **You experience level**
7. **Other information or areas that you would like advice on**

 

^(Many thanks to Kritnc for getting the ball rolling.)
## [3][Creating multiple objects through one form](https://www.reddit.com/r/rails/comments/i8dl92/creating_multiple_objects_through_one_form/)
- url: https://www.reddit.com/r/rails/comments/i8dl92/creating_multiple_objects_through_one_form/
---
I am a beginner rails developer and I am stuck with what seems to be a very specific problem.

I have three tables: User, UserApp, and App.

When a user logs in I want them to have input for 20 apps that they use (netflix, amazon, chase etc.) and each record to be saved in UserApp as you can imagine. I know the simpler way is to just use MongoDB but the client I am working for insisted on using a relational database. 

How do I create such a form where the user can input multiple apps (if the app doesn't already exist on the database then a new app should be created)? Are there any articles or prior posts that I can refer to? Thank you for your help
## [4][I curated all the remote job openings from Hacker News who is hiring thread - August](https://www.reddit.com/r/rails/comments/i7pl8b/i_curated_all_the_remote_job_openings_from_hacker/)
- url: https://www.reddit.com/r/rails/comments/i7pl8b/i_curated_all_the_remote_job_openings_from_hacker/
---
This list contains 398 remote jobs and you can filter them by location or skills.

Here I would like to share the entire remote jobs list from the big list of opportunities. All these are 100% remote jobs not just allowed to work from home during this crisis. These are 100% remote jobs and will continue to follow that after the crisis.

https://remoteleaf.com/whoishiring.   

Note: Select "ruby" in the skills filter to see ruby/rails dev jobs.

- 100% remote full-time jobs
- Spent more than 14 hours to curate this information
## [5][Deploying Rails 6, React, Ruby-2.7.1 via Elastic Beanstalk. Unable to install a later version of nodejs via ebextensions to enable assets:precompile](https://www.reddit.com/r/rails/comments/i7t0js/deploying_rails_6_react_ruby271_via_elastic/)
- url: https://www.reddit.com/r/rails/comments/i7t0js/deploying_rails_6_react_ruby271_via_elastic/
---
Webpacker seems to want a nodejs greater then version 8 so I've tried installing it via ebextensions. 

In the application's configuration:

    RAILS_SKIP_ASSET_COMPILATION true

I've stripped down the ebextension's file to jus tonstall nodejs...

    commands:
      01_node_get:
        cwd: /tmp
        command: 'sudo curl --silent --location https://rpm.nodesource.com/setup_13.x | sudo bash -'

      02_node_install:
        cwd: /tmp
        command: 'sudo yum -y install nodejs'

I have even tried the commands on the AWS instance but it still insists on installing nodejs 6.17

    % sudo yum remove nodejs
    % sudo yum clean all
    % sudo rm /var/cache/yum
    % cd /tmp
    % sudo curl --silent --location https://rpm.nodesource.com/setup_13.x | sudo bash -
    % sudo yum install nodejs

What am I doing wrong. How can I get round this? Any pointers greatly appreciated.
## [6][Limiting Users from AD with Devise](https://www.reddit.com/r/rails/comments/i7sptq/limiting_users_from_ad_with_devise/)
- url: https://www.reddit.com/r/rails/comments/i7sptq/limiting_users_from_ad_with_devise/
---
I have been looking to install Devise into my app to a limited access namespace, but as I look at Devise, I am not seeing if there is a way to limit the users that able to log in. I work for a larger company with a more than substantial AD. I really only want 4-6 people to be able to log in out of the entire AD. Am I approaching this the best way? Is there a better option? Is there an easy way to accomplish this within Devise?
## [7][it seems basic question but i don't find the answer](https://www.reddit.com/r/rails/comments/i7soq9/it_seems_basic_question_but_i_dont_find_the_answer/)
- url: https://www.reddit.com/r/rails/comments/i7soq9/it_seems_basic_question_but_i_dont_find_the_answer/
---
Hello Devs.. i'm beginner rails developer

i'm adding some features for running application and this app have model User and Organization

user is belong to organization and organization has many users

&amp;#x200B;

now what i'm trying to do that when i render organizations i want to include all users that belong to organization with each one

&amp;#x200B;

&amp;#x200B;

Code snap:

&gt;def index  
&gt;  
&gt;\\@organizations = Organization.all     
&gt;  
&gt;render json: \\@organizations  
&gt;  
&gt;end

&amp;#x200B;

and here's user and organization models

&amp;#x200B;

&gt;class Organization &lt; ApplicationRecord      
&gt;  
&gt;has\_many :users, dependent: :delete\_all      
&gt;  
&gt;has\_many :statuses, dependent: :delete\_all      
&gt;  
&gt;has\_many :leads, dependent: :delete\_all      
&gt;  
&gt;has\_attached\_file :pdf      
&gt;  
&gt;validates\_attachment\_content\_type :pdf, content\_type: \['image/jpeg', 'image/png', 'image/gif', 'application/pdf'\]  
&gt;  
&gt;end

&amp;#x200B;

&amp;#x200B;

&gt;class User &lt; ApplicationRecord  
&gt;  
&gt; has\_secure\_password    
&gt;  
&gt;has\_and\_belongs\_to\_many :roles    
&gt;  
&gt;has\_many :leads    
&gt;  
&gt;has\_many :schedules    
&gt;  
&gt;has\_many :comments, dependent: :destroy    
&gt;  
&gt;belongs\_to :supervisor, :class\_name =&gt; "User",foreign\_key: 'supervisor\_id', optional: true    
&gt;  
&gt;has\_many :salesAgents, :class\_name =&gt; "User",foreign\_key: 'supervisor\_id'    
&gt;  
&gt;has\_many :vacation\_requests    
&gt;  
&gt;has\_many :targets    
&gt;  
&gt;belongs\_to :organization , optional:true  
&gt;  
&gt;end

&amp;#x200B;

&amp;#x200B;

&amp;#x200B;

So what i have to do as i tried many things but i can't do it
## [8][What kind of knowledge should a 1year experience Rails developer have?](https://www.reddit.com/r/rails/comments/i7flbp/what_kind_of_knowledge_should_a_1year_experience/)
- url: https://www.reddit.com/r/rails/comments/i7flbp/what_kind_of_knowledge_should_a_1year_experience/
---
I have been working for a year at a company with not much mentoring. I was wondering exactly how much knowledge should I have as someone with one year of experience. I can write API's and generally when I am give a task I can figure my way out. On the other hand, I was tongue tied when I was asked something that a senior developer considered to be basic knowledge.
## [9][Could someone ELI5 how to use docker and automated testing for my rails app?](https://www.reddit.com/r/rails/comments/i7eq5b/could_someone_eli5_how_to_use_docker_and/)
- url: https://www.reddit.com/r/rails/comments/i7eq5b/could_someone_eli5_how_to_use_docker_and/
---
Hello all, I am working on a rails app that is in prod (deployed to DO) and in use by mostly me but still a production app nonetheless.  

I typically do development/deploy by building features, running rspec locally, merging to master, and pushing master to my prod server (which is running on dokku).  

I've read a little about CI/CD and docker and understand them conceptually but not sure how to convert my rails app to use them.  Ideally i would like a test VM that mirrors my production server, that I can push my develop branch to periodically and automatically run tests with each push.

I would also like to run it in a docker container for ease of distribution (it could be a handy thing to self-host for others that are interested).  I know the terminology but not the how-to of building this stuff, and not sure what other products/software is necessary.

I've read about Jenkins, CircleCI, TravisCI, Github Actions, etc - are all these basically the same thing?

Sorry for the rambling and unorganized questions, it's just an area I would like to learn more about with my rails app. Thanks!
## [10][Sooo, any gems out there that help end users create custom queries?](https://www.reddit.com/r/rails/comments/i7fwnv/sooo_any_gems_out_there_that_help_end_users/)
- url: https://www.reddit.com/r/rails/comments/i7fwnv/sooo_any_gems_out_there_that_help_end_users/
---
Thanks in advance for suggestions!


Ok, so the use case is that our users need to be able to pick columns from multiple different tables, select a predicate per column , then select a value. They want a csv that has the columns and values for just the columns they picked... basically just like the results you’d get running a query from the dB. Users don’t know sql and need a GUI. 


Here’s an example. 
The user has a Program (which is associated with things like teams and more). The program can have many years. Per year, there are different kinds of projects. Per project, there are different dollar amounts from varying sources of funding. 


The users want to be able to choose their own attrs and values to end up with queries like “find all programs between year X and year Y where a fund was approved and then unapproved” and get a result with these columns: program’s name, year’s name, team’s name, project’s name, fund amount, fundsource’s name. This is just one example, and they want to be able to make their own queries any time. 


I have some front end ideas, so I’m mostly looking for back-end focused gems that can help generate safe sql queries based on Params from form(s).


We are already using Ransack for some searches, but I can’t seem to figure out how to get what I need from Ransack. My understanding of Ransack is that it’s just returning a collection of instances of one Model —like Program.ransack(q).result will just give me Programs that match the form params, not the details for the related resources. 


I tried out Metabase to see if it would be better to just give the users an actual GUI for the dB, but the users don’t understand joins tables or foreign keys, or the fact that some names are stored in their own tables and the id is saved to the table instead of a string (e.g. team.name_id)


I didn’t want to reinvent the wheel if it’s been done by more talented folks than I!!! Thanks again!
## [11][[FRONT-END] loading="lazy" in the images](https://www.reddit.com/r/rails/comments/i7btsp/frontend_loadinglazy_in_the_images/)
- url: https://www.reddit.com/r/rails/comments/i7btsp/frontend_loadinglazy_in_the_images/
---
Hi guys,

recently an user suggested me to add  loading="lazy" in all my img tag. 

Is it a good idea? Is there a "right way" to do it on rails?

Should I just add loading="lazy" and ... nothing else?

Thanks for your help!
## [12][how to add parse_markdown and truncate at the same time?](https://www.reddit.com/r/rails/comments/i7bwjr/how_to_add_parse_markdown_and_truncate_at_the/)
- url: https://www.reddit.com/r/rails/comments/i7bwjr/how_to_add_parse_markdown_and_truncate_at_the/
---
Usually to "render" the markdown I use this system

    &lt;%= parse_markdown(@item.description).html_safe %&gt;

and to add the truncate I use this system

    &lt;%= truncate(@item.last_update.text, length: 280, separator: ' ') %&gt;

**How to add the markdown in the last example?**

  
I tried with  

    &lt;%= truncate(parse_markdown(@item.last_update.text).html_safe, length: 280, separator: ' ') %&gt;

 but it doesn't work correctly. If the message is short (less than 280), it works. But if it is longer, the markdown doesn't work
