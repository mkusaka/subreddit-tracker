# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/f9t9kq/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/f9t9kq/personal_projects_show_off_your_own_project_andor/
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
## [2][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/fduax9/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/fduax9/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [3][Passwordless auth on Rails](https://www.reddit.com/r/rails/comments/fdr0z9/passwordless_auth_on_rails/)
- url: https://www.reddit.com/r/rails/comments/fdr0z9/passwordless_auth_on_rails/
---
 Is there a way to setup auth without password where user receives one-time password on every login.

Only auth option for Rails that I know of is [Devise](https://github.com/heartcombo/devise) but it doesn't seem to support "passwordless" auth.

Are there any other options and am I missing something about devise not supporting it?
## [4][Why use db migration when there's source control?](https://www.reddit.com/r/rails/comments/fdtpf6/why_use_db_migration_when_theres_source_control/)
- url: https://www.reddit.com/r/rails/comments/fdtpf6/why_use_db_migration_when_theres_source_control/
---
Why is checking db schema into source control not enough? What are the benefits of being able to vary both the db version and the app version?
## [5][How to iterate through json](https://www.reddit.com/r/rails/comments/fdtyfv/how_to_iterate_through_json/)
- url: https://www.reddit.com/r/rails/comments/fdtyfv/how_to_iterate_through_json/
---
Hello,

Am using the Woocommerce api to read product, customer and order data to a rails app. An example of a result for fetching products looks like this [https://pastebin.com/3Y1yfThQ](https://pastebin.com/3Y1yfThQ) How can I read this json (iterating a deeply nested array of hashes if am not wrong) and display the results in a html table...something like the layout below:  

&amp;#x200B;

https://preview.redd.it/9l3gyyl8buk41.png?width=1326&amp;format=png&amp;auto=webp&amp;s=118dc223ed2213180862a608e43e1094f2ce2cb9

Thanks.
## [6][Have you guys ever gone to a meet up?](https://www.reddit.com/r/rails/comments/fdl3ef/have_you_guys_ever_gone_to_a_meet_up/)
- url: https://www.reddit.com/r/rails/comments/fdl3ef/have_you_guys_ever_gone_to_a_meet_up/
---
Hey, guys! 

I'm new to this subreddit and still pretty new to rails, but I've been really enjoying it.  I've been trying to look for some local Rails groups to join(north Atlanta), but all I've found are dead groups on meetup.  

If you've gone to local meetups, where did you find yours?  What was your experience?
## [7][Template engine with layout yield and subtemplate?](https://www.reddit.com/r/rails/comments/fduifi/template_engine_with_layout_yield_and_subtemplate/)
- url: https://www.reddit.com/r/rails/comments/fduifi/template_engine_with_layout_yield_and_subtemplate/
---
I am looking for a template engine that comes with rendering layout with the yield keyword, and the ability to embed subtemplate like rails partials. Basically just like rails's view engine but without using rails.

(I checked out tilt, but it only supports layout yield, but not subtemplate.)
## [8][How to learn rails](https://www.reddit.com/r/rails/comments/fdugnp/how_to_learn_rails/)
- url: https://www.reddit.com/r/rails/comments/fdugnp/how_to_learn_rails/
---
Hi, i'm dot net developer. I want to start learn rails Framework, how tell me. What bests way to start, which ide use? What a trend use clear ror or ror api with spa? May be who can get link for site/book/video where build a rails program for beginners? Do use slim or hamal? I appreciate for any information
## [9][How to write a Factory for model that belongs_to?](https://www.reddit.com/r/rails/comments/fddts7/how_to_write_a_factory_for_model_that_belongs_to/)
- url: https://www.reddit.com/r/rails/comments/fddts7/how_to_write_a_factory_for_model_that_belongs_to/
---
I'm trying to test a model validation that belongs to another model. I write it the Factory like this: 

&amp;#x200B;

    FactoryBot.define do
      factory :outage do
        start_time { Time.now }
        end_time { start_time + 4e3 }
        outage_type { "FO" } 
        frequency { "None" }
        reason { Faker::Lorem.words(number: 5) } 
        service_id { 23 }
      end
    end

and my actual test is this: 

&amp;#x200B;

    require 'rails_helper'
    
    RSpec.describe Outage, type: :model do
      it 'has a valid Factory' do 
        build(:outage).should be_valid
      end
     
    end

The test doesn't pass because \`service must exist\`. I know for a fact that the service DOES exist and yet it doesn't get built in FactoryBot.
## [10][Net::SMTPAuthenticationError in Users::ConfirmationsController#create](https://www.reddit.com/r/rails/comments/fdbu2e/netsmtpauthenticationerror_in/)
- url: https://www.reddit.com/r/rails/comments/fdbu2e/netsmtpauthenticationerror_in/
---
I am about to deploy this project : [http://github.com/prp-e/dakhlokharj](http://github.com/prp-e/dakhlokharj) using docker (and my dockerfile is attached to the repository). I create a docker image and run it, it seems fine until I want to work on its sign up/login. 

When I want to sign up, it loads everything and seems works fine with captcha and even the database, but when it comes to the confirmation email, it gives me that error. What's the problem here?
## [11][Here's my undergraduate thesis application...How could I approach and what resources should I use (Rails inexperienced)](https://www.reddit.com/r/rails/comments/fd3wrf/heres_my_undergraduate_thesis_applicationhow/)
- url: https://www.reddit.com/r/rails/comments/fd3wrf/heres_my_undergraduate_thesis_applicationhow/
---
Hello, I'm currently at the last semester and I would like some thoughts of you guys about the web application that was requested by one of my professors.

Why I'm posting this ? 

It has almost 2 months that I have studying Ruby and I quite having a hard time to fully understand Rails. Before starting with Ruby, I had done so far a simple CRUD MVC in pure PHP. So, It's quite hard for me to grasp what I can done by myself  without following tutorials of others applications.

So, I would like to share the application Functional requirement and maybe someone could give me a brief guidence and gem recomendation for lightweight purposes for the first version of this application.

\- The application its a sharing videos system (channels)

Descrption: A plataform where professors can create many channels and post their videos. The professor has the power to validate (choose) which student can join their channel and see the videos. Student can make commentaries on a video.

So, there only two types of users. Professors and Students.

Functional Requirements:

*  Professors and Students registration and login
* Professors can create many channels and post many videos they want. (Professor X minister two curses, so he need to create two channel)
* Professors can add/invite students that its already on the plataform (I think It could have a search bar and the professors could search by name and invite students)
* Students can make commentaries on a video.

Its pretty just that. I could easily think important and cool functionality to add like a blog space in each channel (since each channel represent each course a professor is teaching) where they can add supporting content for the videos.

How could I start ? What should I search for beforehand ?

I already had draw a database schema and thought about the tables and columns.
## [12][ActiveModel::Errors question](https://www.reddit.com/r/rails/comments/fct5m5/activemodelerrors_question/)
- url: https://www.reddit.com/r/rails/comments/fct5m5/activemodelerrors_question/
---
From [https://medium.com/@rfleury2/a-quick-guide-to-model-errors-in-rails-965e2be3ac93](https://medium.com/@rfleury2/a-quick-guide-to-model-errors-in-rails-965e2be3ac93) this article I know that once a model is derived from `ActiveRecord::Base` it will have access to `ActiveModel::Errors` object which gets populated in case of validation or if we manually insert into it using `class.errors.add()` syntax.I want to know if there are any other ways that `ActiveModel::Errors` will be populated? Like when there is a rollback when that model is being saved and so on. Or when `update_attributes` is called on the model and so on
