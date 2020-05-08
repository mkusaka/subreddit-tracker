# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/g616hm/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/g616hm/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gek2rk/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gek2rk/personal_projects_show_off_your_own_project_andor/
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
## [3][react-rails vs rails-api and react app](https://www.reddit.com/r/rails/comments/gfo0vg/reactrails_vs_railsapi_and_react_app/)
- url: https://www.reddit.com/r/rails/comments/gfo0vg/reactrails_vs_railsapi_and_react_app/
---


[View Poll](https://www.reddit.com/poll/gfo0vg)
## [4][Help deciding where to host an app.](https://www.reddit.com/r/rails/comments/gfmlwu/help_deciding_where_to_host_an_app/)
- url: https://www.reddit.com/r/rails/comments/gfmlwu/help_deciding_where_to_host_an_app/
---
Hi folks, I want to know your opinions on where should I host a react rails-application. This project is currently a rails api and a react frontend, they're in separate repos(for now but I could have one repo if you think it is better). This will be a small app used by 10-15 users, so I believe traffic is not an issue.

What would you suggest to host the app, users won't likely grow, but my biggest unknown is the pricing, this is my first freelance project. I don't want anything fancy, since I don't think this is going to scale.

I'm no devops expert and I don't think I want to be setting up everything manually, so I would like a hosting solution where I  wont have to be constantly turning on/off things, in order to not be billed a lot of money for charges I'm not even using at full. What would you suggest?

For example the company I'm working on now uses AWS but seeing what the infrastructure team does, it does not seems like an easy task. I also have used heroku free plans to host college projects and I actually like it, but I don't know how the non free projects work(or if is even worth it to pay for it).

Any other hosting solutions are welcome, keep in mind that this is a very small project and I also have a react frontend which lives in another repo, but I believe I can make it live inside the rails repo if that something you guys recommend.
## [5][How to set variable on page load?](https://www.reddit.com/r/rails/comments/gft14w/how_to_set_variable_on_page_load/)
- url: https://www.reddit.com/r/rails/comments/gft14w/how_to_set_variable_on_page_load/
---
So I have a hidden form field that gets set with a value onclick from a couple of links. Depending on what link is clicked a different value gets set. The form field is called activeselector.

What I want to do is set an `@activeselctor` variable to the form param if it is there, but set it to '12345' or something if it is not so the page loads with a default value. 

Whats the best way to do this? I have the onclick set up and setting the value, but I'm not sure the best  way to return it back to the page on submit and default it on first page load.
## [6][dynamic query builder (help!)](https://www.reddit.com/r/rails/comments/gfhdcn/dynamic_query_builder_help/)
- url: https://www.reddit.com/r/rails/comments/gfhdcn/dynamic_query_builder_help/
---
Hi everybody!

I have been trying to figure it out how to create a model for `report` that dynamically allows the following things:

1. Could call any model in DB
2. Create nested joins with N-levels (as many as needed, there are cases with 8 nested tables or more!!)
3. Avoid collitions with polymorphic relations
4. Filter the table with dynamic `where()`

&amp;#x200B;

For better understanding this is an example of a daily basis use:

`Isa.includes({funding_option: [{application: [{funding_opportunity_from_resource: [{fund: :company}]}]}]}).where(.....`

&amp;#x200B;

&amp;#x200B;

Any idea?
## [7][Advice on finding technical advisors?](https://www.reddit.com/r/rails/comments/gf7gzv/advice_on_finding_technical_advisors/)
- url: https://www.reddit.com/r/rails/comments/gf7gzv/advice_on_finding_technical_advisors/
---
Hi /rails !

Question:

To anyone who has written a technical book or contributed to one before: 

* Did you get technical advisors to edit / or read before publishing? 
* If so how did you find those individuals / did you find it added value to the book?

Context:

I'm currently writing a book which explains how to build a project that bridges several technologies major ones being Sinatra / Rails (others being go, vuejs).

I'd like to think I know my way around a Rails app but am certainly no Ultra-Rails-Architecture-BestPractice-Guru™️©, so I'd love to hear from the community!

Thanks y'all!
## [8][Notes from DHH's RailsConf keynote interview](https://www.reddit.com/r/rails/comments/geyr61/notes_from_dhhs_railsconf_keynote_interview/)
- url: https://www.reddit.com/r/rails/comments/geyr61/notes_from_dhhs_railsconf_keynote_interview/
---
For those of you who haven't had the chance to watch yet:

https://www.joshuawood.net/notes/2020-railsconf-dhh-keynote

One of the parts I liked best was this comment on repeating history as an industry:

&gt; The amnesia is partially caused by so many new people entering the industry. They haven’t experienced it any other way.

There are a ton of new people who haven’t been exposed to the joy of Rails because JavaScript is mainstream and/or someone told them not to bother.

The “Rails Way” is completely foreign to many of the new JavaScript developers I’ve met. Don’t assume people get it just because it’s been around forever. Spread the word:

&gt; Ruby is great because you can know a little JavaScript and then jump to Ruby, understand it, and own the full stack.

What do y'all think about the state of Ruby and Rails in 2020?
## [9][Where should decorator methods that don't belong to a model go?](https://www.reddit.com/r/rails/comments/gf5r1v/where_should_decorator_methods_that_dont_belong/)
- url: https://www.reddit.com/r/rails/comments/gf5r1v/where_should_decorator_methods_that_dont_belong/
---
Hi all,

I've got pretty familiar with using the draper gem now to store decorator methods for my models in order to clean up my views.  


But now I have the following method in my home controller:

    request.location.city ? @location = request.location.city : @location = "..."

And this is how I'm calling it in my home page:

    placeholder: "Search for results in #{@location}",

The controller seems like the wrong place to put logic purely related to the view.  


I can't see anything in the draper documentation about decorator methods for views. It's all about decorator methods for models.  


So I'm wondering, where should I be putting a method like the one just mentioned?  


Thanks
## [10][Is Everyday Rails Testing with RSpec a Good Read for a Mid-Level Developer?](https://www.reddit.com/r/rails/comments/gf4gqo/is_everyday_rails_testing_with_rspec_a_good_read/)
- url: https://www.reddit.com/r/rails/comments/gf4gqo/is_everyday_rails_testing_with_rspec_a_good_read/
---
As it says on the tin, I'd be keen to hear from developers with, say, 2+ years of experience on whether they would recommend the book?

I appreciate there will be many elements that are for juniors but if it also caters beyond that then I'll be purchasing a copy.

Thanks!
## [11][Please help. Running out of hairs to pull.](https://www.reddit.com/r/rails/comments/geza8v/please_help_running_out_of_hairs_to_pull/)
- url: https://www.reddit.com/r/rails/comments/geza8v/please_help_running_out_of_hairs_to_pull/
---
This is likely more of an OS issue than it is a dev/coding  issue, but I trust you guys' judgement. I am constantly having to reboot my terminal (running on zsh since apple forced it) to run the devise:install and all other gem variants. 

Here's what happens: 

I do something, whatever on the terminal. Then, I run one of the install operators. It works fine. Next time I try it - just finished running the simple\_forms install and now doing devise views - it just freezes up and does nothing. Opening a new tab and running it there doesn't work. Need to cmd + Q out of that b\*\*\*\* and reopen the terimnal, then cd back into the app, then it works fine. I've done this now 6 times in the last 45 minutes. I will be bald by the end of the night if I can't fix this nuance.
## [12][Should I use multi-tenancy in my project ?](https://www.reddit.com/r/rails/comments/geo7h6/should_i_use_multitenancy_in_my_project/)
- url: https://www.reddit.com/r/rails/comments/geo7h6/should_i_use_multitenancy_in_my_project/
---
So, I was taking this course in which  you code a SaaS Management Project. This app uses milia gem for multi tenenacy. There is two types of user. Company and Employee. A company can create a project and then invite employees to be part of this project. And in the projects you can upload files, texts., etc...

Unfortunately I couldn’t get through this course's section due the very outdated content such as milia gem and rails 4.

I was looking forward to code this app because I could imagine my undergraduate thesis project being built using the very bases of this SaaS management project app. 

I’m trying to code a simple video sharing plataform where a professor could have their own private “YouTube channel” where their could invite their students to participate and watch the lectures. 

So Insted of Employee and Company, my models would be Professors and Students.

So, I’m looking for this “privacy” characteristic. That you only would be able to see the content if you were invited or something like that. 

The multi tenancy for my project is the correct approach?
