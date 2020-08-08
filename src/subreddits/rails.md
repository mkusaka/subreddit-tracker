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
## [2][Design Patterns and Anti-Patterns in Rails?](https://www.reddit.com/r/rails/comments/i5vg9b/design_patterns_and_antipatterns_in_rails/)
- url: https://www.reddit.com/r/rails/comments/i5vg9b/design_patterns_and_antipatterns_in_rails/
---
OK, it's more like a *software engineering* topic than a rails related one. But I asked one of my friends about deleting a table manually and re-do the migration (the project is written in Django and not rails) and he told me "This is an Anti-Pattern in Django". 

I knew possible dangers of the idea and I suggested it with the knowledge, but I jokingly answered him "You call everything you don't understand an anti-pattern". 

Now, I'm curious, is there a set of patterns and anti-patterns SPECIFICALLY for rails?
## [3][Deploying Rails 6 app into a subdirectory](https://www.reddit.com/r/rails/comments/i5yinq/deploying_rails_6_app_into_a_subdirectory/)
- url: https://www.reddit.com/r/rails/comments/i5yinq/deploying_rails_6_app_into_a_subdirectory/
---
I'm deploying a Rails 6 application into a subfolder (/dictionnaire-app). I set config.relative\_url\_root = "/dictionnaire-app"  
 in application.rb and also added the ENV variable according to [the documentation](https://edgeguides.rubyonrails.org/configuring.html#deploy-to-a-subdirectory-relative-url-root).

However, if assets are properly rendered, links generated with the link\_to  
 method are not updated accordingly.

I read many questions related to this problem and tried to create a scope in my routes.rb file, following [these instructions](https://gist.github.com/shad1w/75e7ebbbdc5382719cd06f56bbe6acd0).

            root to: "home#index"
    
    	scope 'dictionnaire-app' do
    		resources :words, path: "definition", param: :slug
    
    		post '/search', to: 'words#search'
    		get '/recherche', to: 'words#search_page'
    		get '/:letter', to: 'words#alphabet_page', param: :letter, as: "alphabetic_page"
    	end

Unfortunately the app throws an error saying the page doesn't exist for this URL : [https://www.lalanguefrancaise.com/dictionnaire-app/definition/rameal](https://www.lalanguefrancaise.com/dictionnaire-app/definition/rameal)

However, this URL works: [https://www.lalanguefrancaise.com/dictionnaire-app/dictionnaire-app/definition/rameal](https://www.lalanguefrancaise.com/dictionnaire-app/dictionnaire-app/definition/rameal)

If I remove the scope, the pages are displayed for the right URLs, but the links generated from link\_to are not including "dictionnaire-app"...

This issue seems to be from Rails ([see here](https://github.com/rails/rails/issues/24393)). Any way to find a work around?  
Thanks !
## [4][Is anyone using AnyCable on Heroku?](https://www.reddit.com/r/rails/comments/i5lf38/is_anyone_using_anycable_on_heroku/)
- url: https://www.reddit.com/r/rails/comments/i5lf38/is_anyone_using_anycable_on_heroku/
---
Any gotchas with the deployment or did the AnyCable guide work out for you?
## [5][Is it possible to access Rails database and model methods using an external script?](https://www.reddit.com/r/rails/comments/i5k2iv/is_it_possible_to_access_rails_database_and_model/)
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
## [6][Is fast_jsonapi dead?](https://www.reddit.com/r/rails/comments/i53vm0/is_fast_jsonapi_dead/)
- url: https://www.reddit.com/r/rails/comments/i53vm0/is_fast_jsonapi_dead/
---
It looks like this gem hasn't had any commits for over a year now. Did Netflix stop maintaining this project? Would you guys recommend using another API builder?
## [7][Is it possible to put Sidekiq retries into a different queue?](https://www.reddit.com/r/rails/comments/i54wdy/is_it_possible_to_put_sidekiq_retries_into_a/)
- url: https://www.reddit.com/r/rails/comments/i54wdy/is_it_possible_to_put_sidekiq_retries_into_a/
---
I would like to put Sidekiq retries into a lower priority queue. Is this possible?
## [8][I’m taking over an app which has zero test. I installed minitest and guard but running tests is extremely slow.](https://www.reddit.com/r/rails/comments/i5574g/im_taking_over_an_app_which_has_zero_test_i/)
- url: https://www.reddit.com/r/rails/comments/i5574g/im_taking_over_an_app_which_has_zero_test_i/
---
Every time I run the test suite or that guard detects a change, the tests take something like 45 seconds to more then one minute to start. That’s very unpractical especially since this app is quite big and has zero test written... Where should I look to try to understand and fix that first issue?
## [9][Checklist before putting moving codebase to a public repo?](https://www.reddit.com/r/rails/comments/i55cqa/checklist_before_putting_moving_codebase_to_a/)
- url: https://www.reddit.com/r/rails/comments/i55cqa/checklist_before_putting_moving_codebase_to_a/
---
I have a webapp that I made in rails a while ago and I want to share a copy of the code on a public repo (for potential use in a portfolio if I finally get the courage to try and move from hobbyist to professional).  

  

I had it on a private repo on bitbucket the whole time I developed and deployed it, and want to make sure when I make it public I'm not exposing anything. My site is verry lightly trafficed but I still wouldn't want something to happen to it and I do have a few hundred user emails in there to respect. 

&amp;#x200B;

Thing's I've checked:

\-Everything in config/production.rb is using environment variables

\-secrets.yml is using an environment variable for the production secret key

\-I changed the credentials for the admin account in the db/seeds.rb (in retrospect I should have had this set with an environment variable as well)

&amp;#x200B;

Anything else I may have overlooked or should consider?
## [10][Am I allowed to invite everybody to Ruby on Rails Discord server?](https://www.reddit.com/r/rails/comments/i5eq79/am_i_allowed_to_invite_everybody_to_ruby_on_rails/)
- url: https://www.reddit.com/r/rails/comments/i5eq79/am_i_allowed_to_invite_everybody_to_ruby_on_rails/
---
I'm hoping to gather all Ruby on Rails lovers to learn and communicate.
https://discord.gg/yUpUNAZ
## [11][Working with data from an SQLEXPRESS server](https://www.reddit.com/r/rails/comments/i4xrp4/working_with_data_from_an_sqlexpress_server/)
- url: https://www.reddit.com/r/rails/comments/i4xrp4/working_with_data_from_an_sqlexpress_server/
---
We have an existing SQLEXPRESS database that has some valuable information in it. How do I build a tunnel between rails in that server. Our primary db server is postgresql. Is it possible to establish an API based connection with that server? We simply need to read the data, and build reports with it.
