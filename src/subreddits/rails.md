# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gek2rk/personal_projects_show_off_your_own_project_andor/)
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
## [2][Newbie Question about collection_select](https://www.reddit.com/r/rails/comments/gle01p/newbie_question_about_collection_select/)
- url: https://www.reddit.com/r/rails/comments/gle01p/newbie_question_about_collection_select/
---
Hey,

I am pretty new to Rails, so don't judge me ;) Basically I want to select the article's category using collection\_select. As I understand collection\_select gives me string and I need it to be an object of category Model how do I go about doing it? Does collection\_select have some feature for it or do I need to create my own methods to get the category object by id?
## [3][Anyone know a Repl where I can practice rails?](https://www.reddit.com/r/rails/comments/glben7/anyone_know_a_repl_where_i_can_practice_rails/)
- url: https://www.reddit.com/r/rails/comments/glben7/anyone_know_a_repl_where_i_can_practice_rails/
---
I'm trying to learn rails but I don't have a personal computer. Is there a Repl that supports rails ?
## [4][Best resources for an in-depth look at RSpec?](https://www.reddit.com/r/rails/comments/gl11y3/best_resources_for_an_indepth_look_at_rspec/)
- url: https://www.reddit.com/r/rails/comments/gl11y3/best_resources_for_an_indepth_look_at_rspec/
---

## [5][Rails 6/Webpack question.](https://www.reddit.com/r/rails/comments/gl0z7p/rails_6webpack_question/)
- url: https://www.reddit.com/r/rails/comments/gl0z7p/rails_6webpack_question/
---
So... on the premise of keeping things as clean and consistent as possible, is there any value to adding font awesome into your application via gem or w/e other options are available as opposed to the link in application.html.erb? I feel kind of guilty just pasting a 3rd party link there and using it app-wide.
## [6][How many users are watching the page?](https://www.reddit.com/r/rails/comments/gkxk10/how_many_users_are_watching_the_page/)
- url: https://www.reddit.com/r/rails/comments/gkxk10/how_many_users_are_watching_the_page/
---
I have a def in `user.rb` to check if the user (logged) is online.

I use it in views just checking `&lt;% if u/user.online? %&gt;`

I can also check which is the last movie-page watched by the user, because in `user.rb` I have also this

    has_many :visited_movies, -&gt; { distinct }, through: :visits, source: :movie

Now I want to create a new feature on my website.

In the Movie Pages **I want to show how many logged users are watching now that page**.

I should use the previous datas, right?

How to do? I don't know how to start
## [7][Best way to update two records based on each others values.](https://www.reddit.com/r/rails/comments/gkto3w/best_way_to_update_two_records_based_on_each/)
- url: https://www.reddit.com/r/rails/comments/gkto3w/best_way_to_update_two_records_based_on_each/
---
Hi, I have two 'Creature' records and I would like to perform a calculation based on both records' stats, and then update both records.

This is my current Creatures controller:

    def hunt
        predator = Creature.find(params[:id])
        prey = Creature.find(predator.prey)
        @predator = predator.hunt_predator
        @prey = prey.hunt_prey
    end

And here is a snippet from my Creature model:

    def hunt_prey
        prey = self.clone
        prey.population = prey.population*0.8
        prey.save!
    end

My challenge is that the model changes the records irrespective of the other creatures stats. Is there a way of calling the `predator.hunt_predator` model that takes the other record as input so I can modify both?

Alternatively, am I approaching this in a non-Rails way, and is there a much better practice way of doing this?

Looking online I found a lot of suggestions to use `Creature.update_all`, but couldn't find anything for running the calculations using both records as input. I've considered running the calculations in the controller and then using `Creature.update_all` but I know having a bloated controller is bad practice.

Thanks so much for the help, I'm very new to Rails and been stumped on this for weeks.
## [8][Common issues deploy to production and how to fix that ?](https://www.reddit.com/r/rails/comments/gkrf01/common_issues_deploy_to_production_and_how_to_fix/)
- url: https://www.reddit.com/r/rails/comments/gkrf01/common_issues_deploy_to_production_and_how_to_fix/
---
Hey guys, I wanna deploy my first rails app into production this week or next week in GCP, before going do this I wanna know the common issues when you deploy to production and how to fix that?


In my Rails app, I am using Active Storage with GCP bucket, Whenever Cron for run rake task, Rbenv, Delayed Job, PostgreSql, I saw some articles how to deploy to Ubuntu server , and there are some ways to deploy 
From using Mina, Passenger , Unicorn, and Capistrano. Which you use for this ? And if we use this gem to help us to deploy , would it make issue on the things I use in my Rails app like active job , active storage , or else ? And will make some more configuration for it ??
Thank you for attention
## [9][Setting up an Automatic Book reader with Devise + Rails - 2](https://www.reddit.com/r/rails/comments/gkssv6/setting_up_an_automatic_book_reader_with_devise/)
- url: https://www.reddit.com/r/rails/comments/gkssv6/setting_up_an_automatic_book_reader_with_devise/
---
Hey guys,  I hate how much coding channels focus on a basic 'ToDo list / blog' so i wanted wanted to share my progress in building an automatic book reader (not that good in making vids but meh). 

Here is my second video where we explore the CRUD section of the application.

 [https://www.youtube.com/watch?v=dxFTUA1leRc&amp;feature=youtu.be](https://www.youtube.com/watch?v=dxFTUA1leRc&amp;feature=youtu.be) 

Videos will get better as time goes on ;)
## [10][how should i update my config/webpack/environment.js for bootstrap 4 ?](https://www.reddit.com/r/rails/comments/gkgsl9/how_should_i_update_my_configwebpackenvironmentjs/)
- url: https://www.reddit.com/r/rails/comments/gkgsl9/how_should_i_update_my_configwebpackenvironmentjs/
---
i've looked at three guides, and ended up confused even more:  


  


https://preview.redd.it/v3gbl91ymzy41.png?width=2400&amp;format=png&amp;auto=webp&amp;s=e61348825d05d95585dd5230def6679ec0a295a8
## [11][How to implement SEO in Rails?](https://www.reddit.com/r/rails/comments/gkt6xt/how_to_implement_seo_in_rails/)
- url: https://www.reddit.com/r/rails/comments/gkt6xt/how_to_implement_seo_in_rails/
---
How do you make rails SEO friendly.? Have you implementation it and works very well in rails ?
