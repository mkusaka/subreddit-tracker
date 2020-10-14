# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/j6qvuh/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/j6qvuh/personal_projects_show_off_your_own_project_andor/
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
## [2][Watch me build a new Ruby/Rails community in public 📹](https://www.reddit.com/r/rails/comments/jaoh8a/watch_me_build_a_new_rubyrails_community_in_public/)
- url: https://www.reddit.com/r/rails/comments/jaoh8a/watch_me_build_a_new_rubyrails_community_in_public/
---
Hey all,

I've been wanting to see a place/resource build up more hype about Rails and Ruby again like the good ole' days so I'm building one in public.

Some background about what exactly I plan on building is here: [https://web-crunch.com/posts/lets-build-for-ruby-and-rails-developers](https://web-crunch.com/posts/lets-build-for-ruby-and-rails-developers)

I plan to commit to a few hours of recording my progress a week for an initial MVP of the job board portion of the app. Once that's complete I'll do a tentative soft launch and continue adding more of the community features. 

#### Primary goals:

- Build some buzz around Ruby / Rails.
- Bring more Ruby and Rails developers together.
- Have a centralized place to find Ruby and Rails jobs.
- Showcase profiles developers can use to get hired.
- Have a community forum for all things Ruby and Rails + more web/programming topics.
- Maybe make some side $$

The ongoing YouTube playlist is here: [https://www.youtube.com/playlist?list=PL01nNIgQ4uxOfChhPf3jxq8H6fUWnCeLC](https://www.youtube.com/playlist?list=PL01nNIgQ4uxOfChhPf3jxq8H6fUWnCeLC)

The ongoing collection of articles/videos is here:
[https://web-crunch.com/collections/lets-build-for-ruby-and-rails-developers](https://web-crunch.com/collections/lets-build-for-ruby-and-rails-developers)

The buzz around the Laravel community is inspiring and is a big driver for me to take on this project. 

It's worth noting that I'm not a super-advanced Rails developer like I'm sure many of you are but I believe I know enough to build something such as this (no matter how scrappy to start). Feedback and pointers are welcomed as I'm also using this as a learning experience and potential to earn some side income. 

**TL;DR;** I'm building a community and job board for Ruby/Rails developers in public. Maybe follow along if you're interested?
## [3][Webpacker Issues](https://www.reddit.com/r/rails/comments/jaxxx0/webpacker_issues/)
- url: https://www.reddit.com/r/rails/comments/jaxxx0/webpacker_issues/
---
Hi, just a stressed out dev trying to get some JS to work through Webpacker. The error i'm getting is that a method is undefined. I'm trying to install/compile a third party js libary (masonry) into Webpacker and i'm having an absolute nightmare with it. I've installed the package I want through Yarn and I can see that it's the correct one as when I do 'yarn list', I can see it. Great, so we have it installed and can see in the package.json that it's installed but how do I know what I need to require/import? Surely it should just be a matter of installing it through a package managers (bower, npm or yarn) and then just  requiring the file in the app/javascript/packs/application.js using the name of the file in the package.json?

Is there something i'm missing? I've looked everywhere on how to do this but every article I go to seems to say something else...

**Update**: SOLVED. Issue was I wasn't referencing the file correctly within the application.js file. It's clicked with me now and I think I now understand how it all links up together. To my understanding, whenever you 'require' a file in the application.js, it looks within the .node\_modules folder which I did not know about. I got the solution by looking at my yarn.lock file to see the name of the folder and then looked in the .node\_modules/#{directory} so when I put 'require("#{directory}"), it thought it was the directory than the file. It just confused me as some of the other requires (actiontext, actionstorage etc) didn't reference their files, it references the directory to include (Tell me if i'm wrong).
## [4][Minitest vs RSpec for Rails Development](https://www.reddit.com/r/rails/comments/jay6xk/minitest_vs_rspec_for_rails_development/)
- url: https://www.reddit.com/r/rails/comments/jay6xk/minitest_vs_rspec_for_rails_development/
---
Hi guys. Recently completed a bootcamp, covering a lot of Ruby on Rails and using RSpec as the testing framework from this. I also see from looking online a lot of people set up RSpec as their framework for Rails `rails new -T` without minitest (the default)

  
What are peoples opinions on using Minitest, as I have been approached to interview for a role which uses it and from looking online it seems inferior in many ways. I am trying to get it set up on my rails app (just to play with it) and cant seem to find a good coverage solution (like simplecov) as I found it to be a problem setting this up with Minitest (with the little documentation on it also)

So yeah just a question, what do you guys use to test Rails apps and if its minitest what resources and coverage do you use!
## [5][ActiveStorage query optionals](https://www.reddit.com/r/rails/comments/jazt1a/activestorage_query_optionals/)
- url: https://www.reddit.com/r/rails/comments/jazt1a/activestorage_query_optionals/
---
Hi guys, I was wondering if AS has building queries for doing joins. That way, I would avoiding the n+1 query problem. E.g: In the example bellow, I fetched all the attachments for 'files\_attached' (query 1) and later I display each file's filename (the 3 blob queries for 3 files). 

    ##query 1
    ActiveStorage::Attachment Load (0.3ms)  SELECT "active_storage_attachments".* FROM "active_storage_attachments" WHERE "active_storage_attachments"."record_id" = $1 AND "active_storage_attachments"."record_type" = $2 AND "active_storage_attachments"."name" = $3
    
    ## displaying filenames for each file fetched
    ##query 2
    ActiveStorage::Blob Load (0.2ms)  SELECT "active_storage_blobs".* FROM "active_storage_blobs" WHERE "active_storage_blobs"."id" = $1 LIMIT $2
    
    ##query 3
    ActiveStorage::Blob Load (0.2ms)  SELECT "active_storage_blobs".* FROM "active_storage_blobs" WHERE "active_storage_blobs"."id" = $1 LIMIT $2
    
    ##query 4
    ActiveStorage::Blob Load (0.2ms)  SELECT "active_storage_blobs".* FROM "active_storage_blobs" WHERE "active_storage_blobs"."id" = $1 LIMIT $2

I'm aware that I could make one single query using raw SQL, but I would like to avoid it if possible.
## [6][What can a Noob put in his portfolio ?](https://www.reddit.com/r/rails/comments/jagmh4/what_can_a_noob_put_in_his_portfolio/)
- url: https://www.reddit.com/r/rails/comments/jagmh4/what_can_a_noob_put_in_his_portfolio/
---
As the title said, i'm completely new to ruby/rails what i should put in my portfolio ? any ideas ?
## [7][devise doesn't work after adding a new field](https://www.reddit.com/r/rails/comments/jaeea4/devise_doesnt_work_after_adding_a_new_field/)
- url: https://www.reddit.com/r/rails/comments/jaeea4/devise_doesnt_work_after_adding_a_new_field/
---
Recently, I have created a project and I want people to be known by their usernames instead of emails. So, I added a migration like this:

    rails g migration add_username_to_users username:string:uniq

This is not a problematic thing, is it? By the way, now when I run my app and try to login, it says "invalid email, username or password" and doesn't log me in. What have I missed?
## [8][Need help: Rails credentials in Github actions pipeline](https://www.reddit.com/r/rails/comments/jaa7vc/need_help_rails_credentials_in_github_actions/)
- url: https://www.reddit.com/r/rails/comments/jaa7vc/need_help_rails_credentials_in_github_actions/
---
Hello dear Rails community, I am relatively new to Rails and to learn I am building a new application which I want to deploy to my Ubuntu v-server via Github actions. I am really struggling with deploying the app since using this particular stack (Rails / Docker / Github actions) is new to me.

I dockerized the app and I want to push it to Dockerhub, so I can retrieve the image from Dockerhub on my server. The image pushing and pulling works fine, but the app wouldn't build on the server with the image because I guess I didn't (and still don't) fully understand the process. This was when I posted this question on Stackoverflow:

[https://stackoverflow.com/questions/64239706/cannot-start-dockerized-rails-application-when-pulling-it-from-docker-hub](https://stackoverflow.com/questions/64239706/cannot-start-dockerized-rails-application-when-pulling-it-from-docker-hub)

After studying on how to set up a Rails app in Github actions I have come really far imho and currently I get an error in the pipeline that googling the error suggests something is wrong with my credentials:

    ActiveSupport::MessageEncryptor::InvalidMessage: Cannot load database configuration:
    ActiveSupport::MessageEncryptor::InvalidMessage
    ... (I removed the rest of the stack to save space)
    Caused by:
    OpenSSL::Cipher::CipherError: 

I have /config/master.key and also /config/credentials/&lt;environment&gt;.key for each environment. I also added the master.key to the Github actions secrets and pass the master.key secret like this:

    env:   
      RAILS_ENV: test
      RAILS_MASTER_KEY: ${{ secrets.RAILS_MASTER_KEY }}

When I added my test environment key to the main.yml

    env: 
     RAILS_ENV: test
     RAILS_TEST_KEY: ${{ secrets.RAILS_TEST_KEY }}

as mentioned [here](https://blog.saeloun.com/2019/10/10/rails-6-adds-support-for-multi-environment-credentials.html#storing-encryption-key-in-environment-variables) i got the original error again:

    rails aborted!
    NoMethodError: Cannot load database configuration:
    undefined method `[]' for nil:NilClass

I hope I provided all the necessary info and I would appreciate any help I can get.
## [9][Best practices for ruby gems continuous integrate](https://www.reddit.com/r/rails/comments/jantiu/best_practices_for_ruby_gems_continuous_integrate/)
- url: https://www.reddit.com/r/rails/comments/jantiu/best_practices_for_ruby_gems_continuous_integrate/
---
I have a project which is a Rails engine. It's built using Gitlab CI and published to a private on-premise gem repo. The gem repo is set to disallow re-deploy so that if the gem's version has not changed, the publish step will fail.  I'm wondering what some best practices are around this. Does it make sense to try to incorporate some sort of automation that auto-increments the minor or patch-level version numbers? Or should I do the opposite and have an earlier test that checks to see whether the current version is greater than the previous version?
## [10][What do multiple conditions in the if for callback do?](https://www.reddit.com/r/rails/comments/jadz5n/what_do_multiple_conditions_in_the_if_for/)
- url: https://www.reddit.com/r/rails/comments/jadz5n/what_do_multiple_conditions_in_the_if_for/
---
Suppose I have a user record like this   


    class User &lt; ActiveRecord::Base
      has_many :user_emails, dependent: :destroy
      accepts_nested_attributes_for :user_emails, allow_destroy: true
      before_update :call_main_update, if: [:first_condition, :second_condition]
    
      def first_condition
        puts "FIRST CONDITION"
        return true
      end
      
      def second_condition
        puts "SECOND CONDITION"
        return true 
      end
    
      def call_main_update
        puts "MAIN UPDATE"
      end
    end

is the above code equal to this? 

    before_update :call_main_update, if: -&gt; { first_condition &amp;&amp; second_condition }

From what I tested it out it looks like it. But I couldn't find any documentation here [https://guides.rubyonrails.org/active\_record\_callbacks.html](https://guides.rubyonrails.org/active_record_callbacks.html).   


Could you help me out with what it actually does?
## [11][Way to test your application under big traffic - jMeter](https://www.reddit.com/r/rails/comments/j9nrkh/way_to_test_your_application_under_big_traffic/)
- url: https://www.reddit.com/r/rails/comments/j9nrkh/way_to_test_your_application_under_big_traffic/
---
If you would like to check where are endpoints "to-be-improved" in your application, before users will encounter it during big traffic - there you have quick blogpost about our experience apache-jmeter [https://www.2n.pl/blog/apache-jmeter](https://www.2n.pl/blog/apache-jmeter) . Feedback / other approaches would be appreciated!
