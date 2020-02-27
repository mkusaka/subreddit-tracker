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
## [2][Freelancing in Rails](https://www.reddit.com/r/rails/comments/fa7ojc/freelancing_in_rails/)
- url: https://www.reddit.com/r/rails/comments/fa7ojc/freelancing_in_rails/
---
hey.. how are you all?? I hope you all doing fine.. I want advice from experience people.. Anyone here doing freelancing work in ruby on rails?? can you tell me the struggles of finding a rails job in freelancing?? if someone learning rails and want a career of rails in freelancing, what advice would you like to give??
## [3][Would anyone be interested in taking over development of a passion project that is a tool for guitarists and other musicians?](https://www.reddit.com/r/rails/comments/f9y0ge/would_anyone_be_interested_in_taking_over/)
- url: https://www.reddit.com/r/rails/comments/f9y0ge/would_anyone_be_interested_in_taking_over/
---
The site is www.whatkeyamiin.com. It gets 2,000+ visits per month and has room to grow, but I haven't had time to work on it or develop new features in years. I'd love to see it continue on with someone else who wants to help give back to the guitarist community!

It is built on Rails 4 and jQuery, no front end framework (although it would be a good fit for one now).
## [4][Looking For Projects to Build a Portfolio](https://www.reddit.com/r/rails/comments/fa6f2q/looking_for_projects_to_build_a_portfolio/)
- url: https://www.reddit.com/r/rails/comments/fa6f2q/looking_for_projects_to_build_a_portfolio/
---
Hi All,

I'm a Business Intelligence Consultant looking to pivot careers.   I'm looking for some projects I could work on that would look good as part of a portfolio.  Any recommendations are greatly appreciated, as well as any advice on getting my foot in the door with a job.
## [5][Problems installing ActiveAdmin - allows assigning of a non existent parameter?](https://www.reddit.com/r/rails/comments/fa5ues/problems_installing_activeadmin_allows_assigning/)
- url: https://www.reddit.com/r/rails/comments/fa5ues/problems_installing_activeadmin_allows_assigning/
---
Hi all,  


 I'm trying to experiment with Active Admin but having problems with getting set up. I'm at the final step but can't seem to understand what's going on.  


So during the install it gives me the following line in my seeds file:  


    AdminUser.create!(email: 'admin@example.com', password: 'password', password_confirmation: 'password') if Rails.env.development?

This seems to create the AdminUser instance fine, but then when I try to log in, I get an error:

    Invalid Email or password.

So, I then checked the schema, but it appears that a password field doesn't even exist for that model:  


    create_table "admin_users", force: :cascade do |t|
        t.string "email", default: "", null: false
        t.string "encrypted_password", default: "", null: false
        t.string "reset_password_token"
        t.datetime "reset_password_sent_at"
        t.datetime "remember_created_at"
        t.datetime "created_at", null: false
        t.datetime "updated_at", null: false
        t.index ["email"], name: "index_admin_users_on_email", unique: true
        t.index ["reset_password_token"], name: "index_admin_users_on_reset_password_token", unique: true
      end

But for some reason, it still seems to accept the password being set when we create the instance in seeds. No error.  


When I check the model, it looks like this:  


    =&gt; AdminUser(id: integer, email: string, encrypted_password: string, reset_password_token: string, reset_password_sent_at: datetime, remember_created_at: datetime, created_at: datetime, updated_at: datetime)

So I guess my question is twofold:  
1) Why is it allowing me to assign a password when that column seemingly doesn't exist in the database.  
2) How do I get round this issue?  


Thanks.
## [6][How are you doing structured logging?](https://www.reddit.com/r/rails/comments/f9xi4m/how_are_you_doing_structured_logging/)
- url: https://www.reddit.com/r/rails/comments/f9xi4m/how_are_you_doing_structured_logging/
---
I want to bring some sanity to my logs. I run a solo SaaS business, and want quick and easy insight into debugging/support, business metrics, and general application health. 

I'm on Heroku (STDOUT), using lograge and Papertail at the moment. Lograge is great, but it's not really a logging framework as it doesn't provide a DSL for custom logging, and it leaves a lot of unstructured logs as is: ActiveJob, Sidekiq, custom `Rails.logger.info ...` etc.

There are others like Timber[0], and semantic_logger, but neither seem actively maintained or have widespread adoption.

I think JSON is the way to go, paired with some hosted ELK (Logz.io, Scalyr, Coralogix) or other cloud platform (Datadog, LogDNA, Timber, etc). I've been trialing a bunch of them, but none seem to really click for me. Previously I've been on Papertrail. I want an all-in-one logging + metrics/dashboard service. I like the concept of Datadog, but the Heroku integration is still beta and doesn't work optimally.

Choosing a cloud provider has been a little overwhelming, so I think I should start with optimizing how I log (switching to JSON, ensuring all logs are structured (AJ, Sidekiq, etc), have a good DSL for writing custom business logs w/ context, etc, and then go from there.



[0] [https://docs.timber.io/setup/languages/ruby#](https://docs.timber.io/setup/languages/ruby#)
[1] [http://rocketjob.github.io/semantic_logger/rails.html](http://rocketjob.github.io/semantic_logger/rails.html)
## [7][Tracking Online Users within Your Rails 6 App [Tutorial]](https://www.reddit.com/r/rails/comments/f9xizg/tracking_online_users_within_your_rails_6_app/)
- url: https://www.reddit.com/r/rails/comments/f9xizg/tracking_online_users_within_your_rails_6_app/
---
Hey guys, I'm working on some Rails video tutorials that I want to share with the community. Recently I've seen some questions posted about how to display "online users" within a Rails app and figured it would be a useful video to create for my YouTube channel.

If this sounds interesting to you, then check out the video:

[https://www.youtube.com/watch?v=koXdcLCmGEI](https://www.youtube.com/watch?v=koXdcLCmGEI)

I'm a long term fan of Rails and enjoy creating and sharing content in this space. If you enjoy the video or have new content suggestions then feel free to leave a comment or send me a message.
## [8][I made my side project with Rails!](https://www.reddit.com/r/rails/comments/f9xu5y/i_made_my_side_project_with_rails/)
- url: https://www.reddit.com/r/rails/comments/f9xu5y/i_made_my_side_project_with_rails/
---
Hi y'all!

I've been hard at work building my side project Debugg for the past month and now I think it's good enough to be shared. 

[https://debugg.me](https://debugg.me) 

It's a platform for learning systems development (part by part) by solving exercises, kind of like how Leetcode works. I'd love it if you checked it out and gave me you feedback.

Lastly, I just wanna say how grateful I am that Rails exists. I'd started learning Rails in grade 12 from The Odin Project, with the aim of one day building something of my own that I would be proud of, and Rails is helping me achieve that. I don't think I would've kept interest in building this if I hadn't been able to get it up so quick (scaffolding is awesome!).

Anyways, thanks again. I'll always be grateful to the Rails community.
## [9][Help with permit value](https://www.reddit.com/r/rails/comments/f9v371/help_with_permit_value/)
- url: https://www.reddit.com/r/rails/comments/f9v371/help_with_permit_value/
---
Hi, can someone please help me with permit value from submit button in form. I have four voting buttons and I want to save votes to db. I'm getting undefined method permit for String. 

Form:

    &lt;%= form_with scope: :votes1, url: votes_path, local: true do |form| %&gt;
     
    &lt;div class="row"&gt;
          &lt;div class="col-sm"&gt;
            &lt;%= form.submit :value =&gt; "Nuda", :class =&gt; 'btn btn-danger' %&gt;
          &lt;/div&gt;
          &lt;div class="col-sm"&gt;
            &lt;%= form.submit :value =&gt; "To se mi líbí!", :class =&gt; 'btn btn-success' %&gt;
          &lt;/div&gt;
          &lt;div class="col-sm"&gt;
            &lt;%= form.submit :value =&gt; "Málo detailů", :class =&gt; 'btn btn-secondary'  %&gt;
          &lt;/div&gt;
          &lt;div class="col-sm"&gt;
            &lt;%= form.submit :value =&gt; "Příliš mnoho detailů.", :class =&gt; 'btn btn-warning' %&gt;
          &lt;/div&gt;
        &lt;/div&gt;
    
        &lt;% end %&gt;

Controller:

        def create
            @vote = Vote.new(vote_params)
            @vote.save
            redirect_to root_path
          end
           
        private
            def vote_params
              params.require(:commit).permit(:result)
            end
    end

dbschema:

    ActiveRecord::Schema.define(version: 2020_02_26_081717) do
    
      # These are extensions that must be enabled in order to support this database
      enable_extension "plpgsql"
    
      create_table "votes", force: :cascade do |t|
        t.string "result"
        t.datetime "created_at", precision: 6, null: false
        t.datetime "updated_at", precision: 6, null: false
      end
    
    end
## [10][How to test how much users can use the website in same time - using Camaleon CMS](https://www.reddit.com/r/rails/comments/f9v2h1/how_to_test_how_much_users_can_use_the_website_in/)
- url: https://www.reddit.com/r/rails/comments/f9v2h1/how_to_test_how_much_users_can_use_the_website_in/
---
Hello,   
I need to integrate some website with offers from another. I want to use Camaleon CMS for it. The most basic thing in this project are filters. For example there will be 20 fields that I can filter offers.

My questions are mostly pointed to test and scaling. 

I know it is dependand on how I will write sqls, that will determine the speed of query but I would like to get to know how to make those calculations about how many users can handle Camaleon CMS.  


200 people hitting the page in the same time ? Will it handle ? What kind of server I need at least

1000 people hitting the page in the same time ? Will it handle ? What kind of server I need at least

I want to put this website on specific hosting. I think they share all their machine power with users and you are restricted by "number of processes and RAM". They dont have offer showing what is the power of cpu. So its hard to say for me what are the specifics of servers.   


I would apprecieate any help.
## [11][Questions about TDD and fixtures.](https://www.reddit.com/r/rails/comments/f9koam/questions_about_tdd_and_fixtures/)
- url: https://www.reddit.com/r/rails/comments/f9koam/questions_about_tdd_and_fixtures/
---
Hi all,  


I have just started to learn about TDD, and I am finding the information/tutorials online to be a little bit scattered. I have a few questions piling up and I'd like to just share them with you and hopefully iron out some of my initial hurdles:  


1) Why do I need rspec?

A lot of the tutorials I'm watching use rspec, but it seems to function almost identically to the rails tests that are included with ActiveSupport. What are the advantages of rspec, and should I learn how to use the ActiveSupport testing framework before moving on to rspec?  


2) How do I generate fixtures?  
A lot of the tutorials I'm following just go ahead and open up these fixtures yml files. But although I have the file structure for them, I don't have any .yml files in my fixtures folder. I didn't get these generated when my models were generated.. should they have been? Or am I expected to create these files manually?  


3) What's the best practice for isolating parameters on models when testing?  
Most of the tutorials I've seen show very basic test cases like "should be invalid without a title". This is fine, but normally the example only has a title, and nothing else, so we can be sure that is what is failing. But my model has about 12 fields. So, short of creating 12 different fixtures, each one having a different unique field which is set up to fail the test, how do I run specific tests for all these fields?  


Hope my questions are clear but I often find that when I'm new to something my terminology is all over the place, so let me know if something is unclear and I'll try to explain better.  


Thanks.
