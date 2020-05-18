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
## [2][Shareable Rails app templates](https://www.reddit.com/r/rails/comments/glljbs/shareable_rails_app_templates/)
- url: https://www.reddit.com/r/rails/comments/glljbs/shareable_rails_app_templates/
---
There's been some discussion lately about Rails app templates and how they could be a lot more useful. After having spent a ton of time building Rails app templates for GoRails and Jumpstart, I can definitely feel the pain point. Nate Berkopec [tweeted](https://twitter.com/nateberkopec/status/1259634098353537024) about templates as well recently.

One of the things that is definitely frustrating is installing something like Bootstrap. You've got a bunch of different files to configure and it's easy to miss a step. We've packaged all that into a [Rails app template](https://railsbytes.com/public/templates/x9Qsqx) so all you have to do is run a single command to install it.

These app templates are definitely trickier to build because they'll need to verify Rails version and other dependencies (like is webpacker installed, etc), but I think if done well they can save Rails devs a ton of time. Plus, app templates can be run against existing Rails apps, not just brand new ones. 

So here you have it: [railsbytes.com](https://railsbytes.com)

It took u/king601 and I about 8 hours to build this first version so we know there's a lot to improve, but we hope it can serve as a foundation to help make adding common gems, features, and other things in Rails much easier.

Some interesting examples:

Installing Devise is interactive. It can ask you what model you want and what extra fields so you don't have to remember any of the commands: [https://railsbytes.com/public/templates/X8Bsjx](https://railsbytes.com/public/templates/X8Bsjx) 

Same thing applies to gems like HoneyBadger that need to ask for an API key: [https://railsbytes.com/public/templates/zNPsmV](https://railsbytes.com/public/templates/zNPsmV)

You can have a Rails app template trigger chain other Rails app templates together like this one that installs Rspec, Factorybot, and Standardrb: [https://railsbytes.com/public/templates/V4YsyX](https://railsbytes.com/public/templates/V4YsyX) 

Wanted to share this with you guys and we're curious to hear what you guys think. It seems like it's got a ton of potential to do all sorts of interesting things.
## [3][ISeries Rails Environment and development](https://www.reddit.com/r/rails/comments/glv47u/iseries_rails_environment_and_development/)
- url: https://www.reddit.com/r/rails/comments/glv47u/iseries_rails_environment_and_development/
---
Happy Quarantine everyone. We are looking to start running some web apps on top of some of our clients iSeries systems (System i, AS/400). We got the default rails environment up and running on one of the servers. Does anyone have any recommendations on resources getting started guides or something like that for rails on the iSeries? 

Thanks in advance
## [4][I cannot get rails to allow a react app to post data to it. And I'm losing my mind.](https://www.reddit.com/r/rails/comments/glqmgb/i_cannot_get_rails_to_allow_a_react_app_to_post/)
- url: https://www.reddit.com/r/rails/comments/glqmgb/i_cannot_get_rails_to_allow_a_react_app_to_post/
---
I have a rails 6 api project running on localhost:3000. And I have a React project running on localhost:3001. 

I'm using rack-cors, but it's having no luck. In my application.rb I have the following.

    config.middleware.insert_before 0, Rack::Cors do
      allow do
        origins '*'
        resource '*', headers: :any, methods: :any
      end
    end

and on the react side I have this, trying to post a simple login json object.

    axios.post(urls.post_login, 
           {auth: {email: email, password: password}}, 
           {headers: {'Content-Type': 'application/json'}}
    ).then(blah blah blah)    

And no matter what I try, this is what I get back...

       Access to XMLHttpRequest at 'localhost:3000/user_token' from origin 'http://localhost:3001' has been blocked by CORS policy: Cross origin requests are only supported for protocol schemes: http, data, chrome, chrome-extension, https. 

Worth nothing that posting to the endpoint works fine using postman.

someone please help, Im full on losing it.
## [5][Newbie Question about collection_select](https://www.reddit.com/r/rails/comments/gle01p/newbie_question_about_collection_select/)
- url: https://www.reddit.com/r/rails/comments/gle01p/newbie_question_about_collection_select/
---
Hey,

I am pretty new to Rails, so don't judge me ;) Basically I want to select the article's category using collection\_select. As I understand collection\_select gives me string and I need it to be an object of category Model how do I go about doing it? Does collection\_select have some feature for it or do I need to create my own methods to get the category object by id?
## [6][Feature Request - rails console - quickly switch to dbconsole and back](https://www.reddit.com/r/rails/comments/glh1pz/feature_request_rails_console_quickly_switch_to/)
- url: https://www.reddit.com/r/rails/comments/glh1pz/feature_request_rails_console_quickly_switch_to/
---
  

Ability to quickly switch from the console (bundle exec rails console) to the dbconsole (bundle exec rails dbconsole)
and vice versa
## [7][Best resources for an in-depth look at RSpec?](https://www.reddit.com/r/rails/comments/gl11y3/best_resources_for_an_indepth_look_at_rspec/)
- url: https://www.reddit.com/r/rails/comments/gl11y3/best_resources_for_an_indepth_look_at_rspec/
---

## [8][Anyone know a Repl where I can practice rails?](https://www.reddit.com/r/rails/comments/glben7/anyone_know_a_repl_where_i_can_practice_rails/)
- url: https://www.reddit.com/r/rails/comments/glben7/anyone_know_a_repl_where_i_can_practice_rails/
---
I'm trying to learn rails but I don't have a personal computer. Is there a Repl that supports rails ?
## [9][Rails 6/Webpack question.](https://www.reddit.com/r/rails/comments/gl0z7p/rails_6webpack_question/)
- url: https://www.reddit.com/r/rails/comments/gl0z7p/rails_6webpack_question/
---
So... on the premise of keeping things as clean and consistent as possible, is there any value to adding font awesome into your application via gem or w/e other options are available as opposed to the link in application.html.erb? I feel kind of guilty just pasting a 3rd party link there and using it app-wide.
## [10][How many users are watching the page?](https://www.reddit.com/r/rails/comments/gkxk10/how_many_users_are_watching_the_page/)
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
## [11][Best way to update two records based on each others values.](https://www.reddit.com/r/rails/comments/gkto3w/best_way_to_update_two_records_based_on_each/)
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
