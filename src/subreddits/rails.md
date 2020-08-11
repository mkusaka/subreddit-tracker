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
## [2][I curated all the remote job openings from Hacker News who is hiring thread - August](https://www.reddit.com/r/rails/comments/i7pl8b/i_curated_all_the_remote_job_openings_from_hacker/)
- url: https://www.reddit.com/r/rails/comments/i7pl8b/i_curated_all_the_remote_job_openings_from_hacker/
---
This list contains 398 remote jobs and you can filter them by location or skills.

Here I would like to share the entire remote jobs list from the big list of opportunities. All these are 100% remote jobs not just allowed to work from home during this crisis. These are 100% remote jobs and will continue to follow that after the crisis.

https://remoteleaf.com/whoishiring.   

Note: Select "ruby" in the skills filter to see ruby/rails dev jobs.

- 100% remote full-time jobs
- Spent more than 14 hours to curate this information
## [3][What kind of knowledge should a 1year experience Rails developer have?](https://www.reddit.com/r/rails/comments/i7flbp/what_kind_of_knowledge_should_a_1year_experience/)
- url: https://www.reddit.com/r/rails/comments/i7flbp/what_kind_of_knowledge_should_a_1year_experience/
---
I have been working for a year at a company with not much mentoring. I was wondering exactly how much knowledge should I have as someone with one year of experience. I can write API's and generally when I am give a task I can figure my way out. On the other hand, I was tongue tied when I was asked something that a senior developer considered to be basic knowledge.
## [4][Could someone ELI5 how to use docker and automated testing for my rails app?](https://www.reddit.com/r/rails/comments/i7eq5b/could_someone_eli5_how_to_use_docker_and/)
- url: https://www.reddit.com/r/rails/comments/i7eq5b/could_someone_eli5_how_to_use_docker_and/
---
Hello all, I am working on a rails app that is in prod (deployed to DO) and in use by mostly me but still a production app nonetheless.  

I typically do development/deploy by building features, running rspec locally, merging to master, and pushing master to my prod server (which is running on dokku).  

I've read a little about CI/CD and docker and understand them conceptually but not sure how to convert my rails app to use them.  Ideally i would like a test VM that mirrors my production server, that I can push my develop branch to periodically and automatically run tests with each push.

I would also like to run it in a docker container for ease of distribution (it could be a handy thing to self-host for others that are interested).  I know the terminology but not the how-to of building this stuff, and not sure what other products/software is necessary.

I've read about Jenkins, CircleCI, TravisCI, Github Actions, etc - are all these basically the same thing?

Sorry for the rambling and unorganized questions, it's just an area I would like to learn more about with my rails app. Thanks!
## [5][Sooo, any gems out there that help end users create custom queries?](https://www.reddit.com/r/rails/comments/i7fwnv/sooo_any_gems_out_there_that_help_end_users/)
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
## [6][[FRONT-END] loading="lazy" in the images](https://www.reddit.com/r/rails/comments/i7btsp/frontend_loadinglazy_in_the_images/)
- url: https://www.reddit.com/r/rails/comments/i7btsp/frontend_loadinglazy_in_the_images/
---
Hi guys,

recently an user suggested me to add  loading="lazy" in all my img tag. 

Is it a good idea? Is there a "right way" to do it on rails?

Should I just add loading="lazy" and ... nothing else?

Thanks for your help!
## [7][how to add parse_markdown and truncate at the same time?](https://www.reddit.com/r/rails/comments/i7bwjr/how_to_add_parse_markdown_and_truncate_at_the/)
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
## [8][Where do you validate user input/url params?](https://www.reddit.com/r/rails/comments/i7fdld/where_do_you_validate_user_inputurl_params/)
- url: https://www.reddit.com/r/rails/comments/i7fdld/where_do_you_validate_user_inputurl_params/
---
Where do you put the validations that check the user input(via url in mycase) that's used to fetch data from a model?

lets says I have a has of countries &amp; currencies { us: 'USD', cn: 'CNY', in: 'INR' ..}

Since its a url param are you doing it in controller or within a controller-concern or model or model-concern or seperate class in model which has all these validation rules? I'm kind of puzzled with this..

Thanks!
## [9][Rails 6 Devise and the Locked Namespace](https://www.reddit.com/r/rails/comments/i79b78/rails_6_devise_and_the_locked_namespace/)
- url: https://www.reddit.com/r/rails/comments/i79b78/rails_6_devise_and_the_locked_namespace/
---
In my first attempt to use my company's AD for my website authentication, I have hit a quandary.  I only want the AD access for a specific namespace of the site as the rest of the site is open to the public. I have googled for a good run down of how to set this properly but I have yet to find one. 

I have been following [this article](https://medium.com/@takatoyo/step-by-step-implement-active-directory-auth-with-devise-33590bd3e3f1) in my initial setup.

So I am hoping that someone might tell me if this is how I am supposed to go about doing this.

I have a feeling that it will involve:  
namespace\_controller.rb  
users/sessions\_controller.rb (as setup by devise)  
namespace/thing1\_controller.rb  
namespace/thing2\_controller.rb
## [10][Using a Custom Image as Toggle for Form Input?](https://www.reddit.com/r/rails/comments/i7gkib/using_a_custom_image_as_toggle_for_form_input/)
- url: https://www.reddit.com/r/rails/comments/i7gkib/using_a_custom_image_as_toggle_for_form_input/
---
Hi All,

I am new to Rails and had a question about form inputs. I have an image that will be green for switched on and gray for switched off. What kind of form input should I research? The closest thing I see is a radio button, but I need it to be the gray/green image instead of the default one. 

Thanks in advance!
## [11][The best Ruby/Rails profiler?](https://www.reddit.com/r/rails/comments/i74jms/the_best_rubyrails_profiler/)
- url: https://www.reddit.com/r/rails/comments/i74jms/the_best_rubyrails_profiler/
---
What is the best profiling tool for Ruby on Rails? I'm looking for a tool which could return executed functions stack with function names, number of calls and time elapsed on each function. The goal is to find slow methods which are causing performance issues.
