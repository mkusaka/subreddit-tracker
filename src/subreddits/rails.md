# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/ep2dw9/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/ep2dw9/personal_projects_show_off_your_own_project_andor/
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
## [2][ActiveInteractor v1.0.0 Release](https://www.reddit.com/r/rails/comments/eu4z1b/activeinteractor_v100_release/)
- url: https://www.reddit.com/r/rails/comments/eu4z1b/activeinteractor_v100_release/
---
Hey ruby friends! 

Over the weekend I released v1.0.0 of ActiveInteractor, an implementation of the Command Pattern for Ruby with ActiveModel::Validations based on the interactors gem. It comes with rich support for attributes, callbacks, and validations, and thread safe performance methods. 

This update has some major improvements to organizers as well as rails QOL improvements and a lot more.  Please check it out, let me know what you think!   

https://github.com/aaronmallen/activeinteractor 

[https://rubygems.org/gems/activeinteractor](https://rubygems.org/gems/activeinteractor)

[https://www.rubydoc.info/gems/activeinteractor](https://www.rubydoc.info/gems/activeinteractor/)
## [3][Best tools/database for a text-based search](https://www.reddit.com/r/rails/comments/etsj6z/best_toolsdatabase_for_a_textbased_search/)
- url: https://www.reddit.com/r/rails/comments/etsj6z/best_toolsdatabase_for_a_textbased_search/
---
I need to build a full-text search page in a rails application. I brainstormed with my team to choose which tools we could use, especially the database. We stand between Elasticsearch and Postgres. But we are not sure which would be the best choice. 

Today we have few records, but the expectation is that this solution will support a large scale.

Any tips for a search based on full-text? Database, tools, articles?
## [4][Checking user logged in or not (Devise)](https://www.reddit.com/r/rails/comments/etxy78/checking_user_logged_in_or_not_devise/)
- url: https://www.reddit.com/r/rails/comments/etxy78/checking_user_logged_in_or_not_devise/
---
I planned for adding a front-page (not static) to my app. And now, I actually want it to be like this : 

If user not logged in, shows a message like "Please login to your account" with a link to login page. And If user is logged in, shows something belongs to the user. I actually could do the same with help of you guys, but it was in scaffolds and not a view/controller different from that scaffold. 

Thanks!
## [5][How to handle large data imports with validations?](https://www.reddit.com/r/rails/comments/etptuq/how_to_handle_large_data_imports_with_validations/)
- url: https://www.reddit.com/r/rails/comments/etptuq/how_to_handle_large_data_imports_with_validations/
---
I have a database that has around a million records and growing. We have started importing records from an external system for reporting purposes and we are adding anywhere from 5k to 100k records per table. 

The tables with 100k records being imported takes hours to finish uploading. We are trying to speed up the process because during the upload window the system cannot be used, since the reports might be wrong with incomplete data.

We tried pgcopy and it is really fast and finishes in seconds, so one option is to go down that route, but it is just a data dump and we do it manually which is painful. The other problem is we have validations on the system and checks across tables that we do in rails. We also update records in case a duplicate is found. Which is again done in rails, using a first_or_initialize.

So I was wondering if anyone knew a faster way to import large data while keeping the rails validations and duplicate updates in rails. Or is the only option to go with a raw SQL to import the data?
## [6][Can someone help me in chat to deploy one GitHub Project on my hosting?](https://www.reddit.com/r/rails/comments/etw4in/can_someone_help_me_in_chat_to_deploy_one_github/)
- url: https://www.reddit.com/r/rails/comments/etw4in/can_someone_help_me_in_chat_to_deploy_one_github/
---
I have found one project and I would like to host it on my hosting.
But I really dont know how.
I know that hosting support rails but I am getting some errors like page not found and etc. Maybe I'm doing something wrong XD
Thanks

Hosting: HawkHost.com
## [7][Recording Building a Rails Based Accounting System](https://www.reddit.com/r/rails/comments/etk8a9/recording_building_a_rails_based_accounting_system/)
- url: https://www.reddit.com/r/rails/comments/etk8a9/recording_building_a_rails_based_accounting_system/
---
Hi guys

Recently added a video to the "reality tv" type of series I'm vlogging in youtube. It's not really a tutorial but rather a commentary / recording of how one would go about learning and building an application from scratch; in this case, an accounting system. It includes all mistakes / researching work / opinions etc. 

Warning: It may be a bit boring since I don't cut out anything to capture literally everything. 

Looking for suggestions or corrections or comments as to how I might go about a certain topic or even the video format itself.

Thanks.

Playlist here:

[https://www.youtube.com/playlist?list=PL2-7U6BzddIZ35bJdCFx6RZ-QR8n\_JD82](https://www.youtube.com/playlist?list=PL2-7U6BzddIZ35bJdCFx6RZ-QR8n_JD82)
## [8][I'm struggling to get a rails app running locally - any help would be much appreciated! It seems to freeze/hang during assets:precompile.](https://www.reddit.com/r/rails/comments/etpe2l/im_struggling_to_get_a_rails_app_running_locally/)
- url: https://www.reddit.com/r/rails/comments/etpe2l/im_struggling_to_get_a_rails_app_running_locally/
---
I've cloned the repo of a Rails 5.2 app that I will be working on and I am trying to get it working in development but having some issues. When I run  `rails s` it seems to start fine but when I go to localhost:3000 I just hangs/freezes and won't load the page. The server logs show this:

    =&gt; Booting Thin
    =&gt; Rails 5.2.2.1 application starting in development on http://0.0.0.0:3000
    =&gt; Run `rails server -h` for more startup options
    Both Evidence and its :state machine have defined a different default for "state". Use only one or the other for defining defaults to avoid unexpected behaviors.
    Thin web server (v1.7.2 codename Bachmanity)
    Maximum connections set to 1024
    Listening on 0.0.0.0:3000, CTRL+C to stop
    Started GET "/login" for 127.0.0.1 at 2020-01-25 11:23:04 +0100
    Processing by Users::SessionsController#new as HTML
       (1.8ms)  SET NAMES utf8,  @@SESSION.sql_mode = 'TRADITIONAL',  @@SESSION.sql_auto_is_null = 0, @@SESSION.wait_timeout = 2147483
      Rendering users/sessions/new.html.haml within layouts/sessions
      Rendered users/sessions/new.html.haml within layouts/sessions (50.5ms)

I also ran  `rails assets:precompile --trace` to try and see what's happening and it hangs/freezes at the point where it gets to  `** Execute assets:precompile` . Here's the full output:

    ** Invoke assets:precompile (first_time) 
    ** Invoke assets:environment (first_time) 
    ** Execute assets:environment 
    ** Invoke environment (first_time) 
    ** Execute environment 
    ** Invoke bower:before_precompile (first_time) 
    ** Execute bower:before_precompile 
    ** Invoke yarn:install (first_time) 
    ** Execute yarn:install 
    ** Execute assets:precompile

What could be causing this and how might  it be fixed?  I'm using Ubuntu 18.04, Rails 5.2 and Ruby 2.5.3.  Are there any other logs/files/settings that I could post here to give a better idea? Thanks in advance for any help or tips!

&amp;#x200B;

UPDATE: Installed and ran rbspy and the seems to hang/freeze on this: `fetch -` `/home/me/.rvm/gems/ruby-2.5.3/gems/sprockets-4.0.0.beta8/lib/sprockets/cache.rb`
## [9][Please critcize my personal site CV - ROR / Ruby developer looking to relocate to EUR or USA](https://www.reddit.com/r/rails/comments/etc6yk/please_critcize_my_personal_site_cv_ror_ruby/)
- url: https://www.reddit.com/r/rails/comments/etc6yk/please_critcize_my_personal_site_cv_ror_ruby/
---
[https://pablo-vizcay.tech/](https://pablo-vizcay.tech/)

Any feedback appreciated.. good, bad, suggestions. Especially if you are related or have exprience with the hiring process of software engineers.  


======= EDIT 01/25 =======  
Thanks everyone for your suggeestions, I've updated the CV online with a lot of them.
## [10][Need help with sending information to user](https://www.reddit.com/r/rails/comments/etihgu/need_help_with_sending_information_to_user/)
- url: https://www.reddit.com/r/rails/comments/etihgu/need_help_with_sending_information_to_user/
---
Hi. I want to create an app that is automatically sending an email/sms to a user after a payment is made. The email sent contains a username and password that are stored in a database. The payers are not required to generate login information in order to make a payment. The website looks more simple than craigslist. My question is how do I do this? :) What are the marks/checkpoints I have to code my way to? Are there specific gems that might help speed up the building of the app? I am a complete n00b :) 10x
## [11][Problem with Resque-Scheduler! Loading Schedule, Schedule Empty](https://www.reddit.com/r/rails/comments/etbxpn/problem_with_resquescheduler_loading_schedule/)
- url: https://www.reddit.com/r/rails/comments/etbxpn/problem_with_resquescheduler_loading_schedule/
---
Hello! Well, I had a problem while I were doing the installation of pageflow [https://github.com/codevise/pageflow](https://github.com/codevise/pageflow)

But, in the moment that I execute the foreman start, I receive a  warning and I cannot connect to the localhost. This is the warning:

    resque-scheduler: [INFO] 2020-01-22T14:39:17+00+00: Loading Schedule 
    resque-scheduler: [INFO] 2020-01-22T14:39:17+00+00: Schedule empty! Set Resque.schedule 
    resque-scheduler: [INFO] 2020-01-22T14:39:17+00+00: Schedules Loaded  
    
    Passing 'info' command to redis as is; administrative commands cannot be effectively namespaced and should be called on the redis connection directly; passthrought has been deprecated and will be removed in redis-namespace 2.0 (at /var/lib/gems/2.5.0/gems/resque-1.27.4/lib/resque/data_store.rb:59:in `method_missing') 
    
    Passing 'script' command to redis as is; administrative commands cannot be effectively namespaced and should be called on the redis connection directly; passthrought has been deprecated and will be removed in redis-namespace 2.0 (at /var/lib/gems/2.5.0/gems/resque-1.27.4/lib/resque/data_store.rb:59:in `method_missing') 

And I don't know why this is happening, I tried to look for a  question or solve this problem, but nothing. Now I'm stucked in this  point and I don't know how to go forward or what to do. Any help would be appreciated!

Thanks!
