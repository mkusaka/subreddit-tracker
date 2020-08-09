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
## [2][Rails 6 and the Secret Keys](https://www.reddit.com/r/rails/comments/i68uij/rails_6_and_the_secret_keys/)
- url: https://www.reddit.com/r/rails/comments/i68uij/rails_6_and_the_secret_keys/
---
I am in the process of attempting my first deployment of a new app in Rails 6 and when it goes to do the first db migration I get `ArgumentError: Missing \`secret_key_base\` for 'production' environment, set this string with \`rails credentials:edit\``

I am all for setting this up, but I at this point I am confused as to if this is all something I set up production space or development space. Is there a good walk-through of how and where to set this up? I tend to learn best for seeing something set up in front of me. Any help will be great!
## [3][Deploying a Rails app on Apache2](https://www.reddit.com/r/rails/comments/i689yt/deploying_a_rails_app_on_apache2/)
- url: https://www.reddit.com/r/rails/comments/i689yt/deploying_a_rails_app_on_apache2/
---
Hi 👋

I want to deploy a Rails app on my Apache2 web server. Coming from a PHP background, I'm used to the CGI-like approach of mod_php.

After reading through resources on this topic, I wanted to ask you whether it's a good or at least okay idea to do the following:

1.	Use Puma on the Rails app to start a local server
2.	Reverse proxy that through Apache2

I also read a lot about this Passenger thing, but it seems kinda weird to deploy. Any thoughts?
## [4][Design Patterns and Anti-Patterns in Rails?](https://www.reddit.com/r/rails/comments/i5vg9b/design_patterns_and_antipatterns_in_rails/)
- url: https://www.reddit.com/r/rails/comments/i5vg9b/design_patterns_and_antipatterns_in_rails/
---
OK, it's more like a *software engineering* topic than a rails related one. But I asked one of my friends about deleting a table manually and re-do the migration (the project is written in Django and not rails) and he told me "This is an Anti-Pattern in Django". 

I knew possible dangers of the idea and I suggested it with the knowledge, but I jokingly answered him "You call everything you don't understand an anti-pattern". 

Now, I'm curious, is there a set of patterns and anti-patterns SPECIFICALLY for rails?
## [5][New Video: Devise Profile by User Role](https://www.reddit.com/r/rails/comments/i60iis/new_video_devise_profile_by_user_role/)
- url: https://www.reddit.com/r/rails/comments/i60iis/new_video_devise_profile_by_user_role/
---
In this tutorial video I walk through how to create dynamic profile pages by user role.  We will use a technique I learned from my favorite Ruby on Rails tutorial creators, GoRails, to dynamically display the correct partial by current user role.

I received permission before posting from Chris at GoRails to demonstrate this technique which I have seen a few times, but never in this way.  I thought it may be an interesting application of the technique and also answered a subscriber request.  

I have only created 25 or so videos, so still very new to the process, all feedback is welcome as I just want to improve.  

[https://youtu.be/wbRDqZCchs0](https://youtu.be/wbRDqZCchs0)
## [6][How do I fix this error NameError: uninitialized constant AWS](https://www.reddit.com/r/rails/comments/i60c6i/how_do_i_fix_this_error_nameerror_uninitialized/)
- url: https://www.reddit.com/r/rails/comments/i60c6i/how_do_i_fix_this_error_nameerror_uninitialized/
---
Please help me deploy! I'm having this issue:  [https://stackoverflow.com/questions/63307710/a-do-i-fix-this-error-nameerror-uninitialized-constant-aws](https://stackoverflow.com/questions/63307710/a-do-i-fix-this-error-nameerror-uninitialized-constant-aws)  with AWS s3 deploying to Heroku. I'm also using Active Storage.
## [7][Is anyone using AnyCable on Heroku?](https://www.reddit.com/r/rails/comments/i5lf38/is_anyone_using_anycable_on_heroku/)
- url: https://www.reddit.com/r/rails/comments/i5lf38/is_anyone_using_anycable_on_heroku/
---
Any gotchas with the deployment or did the AnyCable guide work out for you?
## [8][Is it possible to access Rails database and model methods using an external script?](https://www.reddit.com/r/rails/comments/i5k2iv/is_it_possible_to_access_rails_database_and_model/)
- url: https://www.reddit.com/r/rails/comments/i5k2iv/is_it_possible_to_access_rails_database_and_model/
---
I have a Kafka consumer which keeps polling for new messages. Once a message comes in, I want to be able to update the database in my rails application using ActiveRecord. 


Model:


    class XYZ &lt; ApplicationRecord
         # ...
         def self.update_entity_record(en_id, status)
               # ...
         end
    end


Kafka consumer:


    # Setup Kafka connection
    # ...

    kafka.each_message(topic: my_topic) do |msg|
           msg = msg.value
           XYZ.update_entity_record(msg.fetch('en_id'), msg.fetch('status')
    end


Is it possible?
## [9][Is fast_jsonapi dead?](https://www.reddit.com/r/rails/comments/i53vm0/is_fast_jsonapi_dead/)
- url: https://www.reddit.com/r/rails/comments/i53vm0/is_fast_jsonapi_dead/
---
It looks like this gem hasn't had any commits for over a year now. Did Netflix stop maintaining this project? Would you guys recommend using another API builder?
## [10][Is it possible to put Sidekiq retries into a different queue?](https://www.reddit.com/r/rails/comments/i54wdy/is_it_possible_to_put_sidekiq_retries_into_a/)
- url: https://www.reddit.com/r/rails/comments/i54wdy/is_it_possible_to_put_sidekiq_retries_into_a/
---
I would like to put Sidekiq retries into a lower priority queue. Is this possible?
## [11][I’m taking over an app which has zero test. I installed minitest and guard but running tests is extremely slow.](https://www.reddit.com/r/rails/comments/i5574g/im_taking_over_an_app_which_has_zero_test_i/)
- url: https://www.reddit.com/r/rails/comments/i5574g/im_taking_over_an_app_which_has_zero_test_i/
---
Every time I run the test suite or that guard detects a change, the tests take something like 45 seconds to more then one minute to start. That’s very unpractical especially since this app is quite big and has zero test written... Where should I look to try to understand and fix that first issue?
