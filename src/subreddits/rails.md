# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/ezrfed/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/ezrfed/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][An in-depth guide on how to create a social media app using Ruby on Rails and React Native](https://www.reddit.com/r/rails/comments/f1keb9/an_indepth_guide_on_how_to_create_a_social_media/)
- url: https://www.reddit.com/r/rails/comments/f1keb9/an_indepth_guide_on_how_to_create_a_social_media/
---
Hi! I wanted to write a blog post series going into detail with creating a production-ready Rails app. So far it is only 5 parts covering the Rails side.

It's going to be a long series, but it helps me fine tune my skills with Rails as well as provide a new resource to those wanting to know how to create a production-ready app.

[Part 1: Setting Up Rails](https://armaizadenwala.com/blog/social-network-app-using-rails-and-react-native-rails-setup/)

[Part 2: PostgreSQL Setup](https://armaizadenwala.com/blog/social-network-app-using-rails-and-react-native-postgresql-setup/)

[Part 3: Authentication Using Devise](https://armaizadenwala.com/blog/social-network-app-using-rails-and-react-native-auth-setup/)

[Part 4: Creating A User](https://armaizadenwala.com/blog/social-network-app-using-rails-and-react-native-creating-a-user/)

[Part 5: Creating Login And Register Endpoints](https://armaizadenwala.com/blog/social-network-app-using-rails-and-react-native-creating-auth-endpoints/)
## [3][Built my own inline editor in rails!!!!!!](https://www.reddit.com/r/rails/comments/f1ps1f/built_my_own_inline_editor_in_rails/)
- url: https://www.reddit.com/r/rails/comments/f1ps1f/built_my_own_inline_editor_in_rails/
---
[https://www.youtube.com/watch?v=4pYYiM1B5nw](https://www.youtube.com/watch?v=4pYYiM1B5nw)
## [4][Implications of increasing the database connection pool](https://www.reddit.com/r/rails/comments/f1af9u/implications_of_increasing_the_database/)
- url: https://www.reddit.com/r/rails/comments/f1af9u/implications_of_increasing_the_database/
---
Hello! I've just launched my app on product hunt yesterday, getting tons of traffic, and now a couple of dev ops issues have cropped up that I'd love your thoughts on. I'm seeing some errors coming in:**ActiveRecord::ConnectionTimeoutError: could not obtain a connection from the pool within 5.000 seconds (waited 5.001 seconds); all pooled connections were in use**

The app builds meal plans, and the 'build' endpoint is a big crazy process that can take anywhere from like 1s - 60s. My database pool size is set to 5, I'm using puma web server on AWS with thread count set to 5, and I'm running 2 - 4 t3.medium instances.

So I think this means essentially I can only have \~5 people hitting the generate mealplan button at any time before it's likely to hit this error as all the connections are in use.

So my questions:

1. What do I need to think about when changing the database connection pool to a larger quantity? Trade offs etc
2. Does the number of instances change my actual database connection pool size? For instance if the pool size is set to 5, and I have 2 instances, does that mean I can have 10 connections at once? I think no it's still 5, but I'm not sure on that.

Thanks!
## [5][Deploy scraper on heroku and in rails?](https://www.reddit.com/r/rails/comments/f1c750/deploy_scraper_on_heroku_and_in_rails/)
- url: https://www.reddit.com/r/rails/comments/f1c750/deploy_scraper_on_heroku_and_in_rails/
---
Hi! Learning Ruby since a few months and built this scraper for learning purposes. I want to add it to a rails app and deploy on heroku later. Would that be possible? 

If you can point me to a direction where i can learn more would be great!

See here the github [https://github.com/sljmn/bol\_scraper](https://github.com/sljmn/bol_scraper) and demo in readme. Any feedback on my code?

With this scraper you can find the inventory for a certain product, based on ean.
## [6][[Windows 10] I can't run "bundle install" due to libv8](https://www.reddit.com/r/rails/comments/f1aqxu/windows_10_i_cant_run_bundle_install_due_to_libv8/)
- url: https://www.reddit.com/r/rails/comments/f1aqxu/windows_10_i_cant_run_bundle_install_due_to_libv8/
---
I am running this on Windows 10.I am trying build this repository (incase the Gemfile and Gemfile.lock are helpful):[https://github.com/huginn/huginn](https://github.com/huginn/huginn)

Error:

&gt;C:/Ruby24-x64/lib/ruby/gems/2.4.0/gems/libv8-7.3.492.27.1/ext/libv8  
&gt;  
&gt;C:/Ruby24-x64/bin/ruby.exe -r ./siteconf20200209-30680-1j9g8t0.rb extconf.rb  
&gt;  
&gt;creating Makefile  
&gt;  
&gt;The system cannot find the path specified.  
&gt;  
&gt;The system cannot find the path specified.  
&gt;  
&gt;The system cannot find the path specified.  
&gt;  
&gt;C:/Ruby24-x64/lib/ruby/gems/2.4.0/gems/libv8-7.3.492.27.1/ext/libv8/builder.rb:57:in  
&gt;  
&gt;\`setup\_python!': libv8 requires python 2 to be installed in order to build, but  
&gt;  
&gt;it is currently not available (RuntimeError)  
&gt;  
&gt;from  
&gt;  
&gt;C:/Ruby24-x64/lib/ruby/gems/2.4.0/gems/libv8-7.3.492.27.1/ext/libv8/builder.rb:39:in  
&gt;  
&gt;\`build\_libv8!'  
&gt;  
&gt;from  
&gt;  
&gt;C:/Ruby24-x64/lib/ruby/gems/2.4.0/gems/libv8-7.3.492.27.1/ext/libv8/location.rb:24:in  
&gt;  
&gt;\`install!'  
&gt;  
&gt;from extconf.rb:7:in \`&lt;main&gt;'  
&gt;  
&gt;extconf failed, exit code 1  
&gt;  
&gt;Gem files will remain installed in  
&gt;  
&gt;C:/Ruby24-x64/lib/ruby/gems/2.4.0/gems/libv8-7.3.492.27.1 for inspection.  
&gt;  
&gt;Results logged to  
&gt;  
&gt;C:/Ruby24-x64/lib/ruby/gems/2.4.0/extensions/x64-mingw32/2.4.0/libv8-7.3.492.27.1/gem\_make.out  
&gt;  
&gt;An error occurred while installing libv8 (7.3.492.27.1), and Bundler cannot  
&gt;  
&gt;continue.  
&gt;  
&gt;Make sure that \`gem install libv8 -v '7.3.492.27.1' --source  
&gt;  
&gt;'[https://rubygems.org/'\`](https://rubygems.org/'`) succeeds before bundling.  
&gt;  
&gt;In Gemfile:  
&gt;  
&gt;mini\_racer was resolved to 0.2.9, which depends on  
&gt;  
&gt;libv8

What ever I google, the top answer is always the same:I tried both:

    gem install libv8 -v '3.16.14.1' -- --with-system-v8

or

    gem install libv8 -v '3.11.8.13' -- --with-system-v8

[https://stackoverflow.com/questions/16514758/gem-install-libv8-version-3-11-8-17-on-ruby-windows](https://stackoverflow.com/questions/16514758/gem-install-libv8-version-3-11-8-17-on-ruby-windows)[https://stackoverflow.com/questions/17346681/libv8-required-python-2-to-be-installed-in-order-to-build-windows/22579851](https://stackoverflow.com/questions/17346681/libv8-required-python-2-to-be-installed-in-order-to-build-windows/22579851)[https://stackoverflow.com/questions/23536893/therubyracer-gemextbuilderror-error-failed-to-build-gem-native-extension](https://stackoverflow.com/questions/23536893/therubyracer-gemextbuilderror-error-failed-to-build-gem-native-extension)

I see the Gemfile.lock has this

&gt;libv8 (7.3.492.27.1)  
.....  
mini\_racer (0.2.9)  
  libv8 (&gt;= 6.9.411)

So I assume the Gemfile.lock is causing this to install the version I don't want. But I don't know how to resolve this.

I am still learning Ruby and this is my first big project I decided to start trying out.

Also I do have Python 2 on my computer. So I'm not sure why it's saying it doesn't exist.
## [7][How can I know which N+1 queries should stay as a N+1 queries?](https://www.reddit.com/r/rails/comments/f1amyx/how_can_i_know_which_n1_queries_should_stay_as_a/)
- url: https://www.reddit.com/r/rails/comments/f1amyx/how_can_i_know_which_n1_queries_should_stay_as_a/
---
I've always heard this, "fixing" some N+1 queries could hurt performance, but I've never heard how can I know how can I determine this.
## [8][what is the rails way to use library that need instance?](https://www.reddit.com/r/rails/comments/f1633k/what_is_the_rails_way_to_use_library_that_need/)
- url: https://www.reddit.com/r/rails/comments/f1633k/what_is_the_rails_way_to_use_library_that_need/
---
hello all I'm trying to use this fcm library 

https://github.com/spacialdb/fcm

problem is to use this library I need to create a instance of it

    require 'fcm'

    fcm = FCM.new("my_server_key")
    # you can set option parameters in here
    #  - all options are pass to HTTParty method arguments
    #  - ref: https://github.com/jnunemaker/httparty/blob/master/lib/httparty.rb#L29-L60
    #  fcm = FCM.new("my_server_key", timeout: 3)

    registration_ids= ["12", "13"] # an array of one or more client registration tokens

    # See https://firebase.google.com/docs/reference/fcm/rest/v1/projects.messages for all available options.
    options = { "notification": {
                  "title": "Portugal vs. Denmark",
                  "body": "5 to 1"
              }
    }
    response = fcm.send(registration_ids, options)


so whenever I need to use the `fcm.send` I redo all the initiate code and I thought it as not a rails way,

because I should have 1 fcm object instead 1 on each of class


what is the rails way to use this gem?


I've read about module but not really sure how to execute it
## [9][How can i completely uninstall Ruby and Ruby on Rails completely from Windows to do a new installation](https://www.reddit.com/r/rails/comments/f18i86/how_can_i_completely_uninstall_ruby_and_ruby_on/)
- url: https://www.reddit.com/r/rails/comments/f18i86/how_can_i_completely_uninstall_ruby_and_ruby_on/
---
How can i completely uninstall Ruby and Ruby on Rails completely from Windows to do a new installation, there may be some issues of my previous installation.
## [10][Google Calendar API authorization help](https://www.reddit.com/r/rails/comments/f14fxe/google_calendar_api_authorization_help/)
- url: https://www.reddit.com/r/rails/comments/f14fxe/google_calendar_api_authorization_help/
---
So, I'm able to use omniauth to authenticate a user with rails. I've also gotten to this point where I want to start using the google calendar service. However, I think this code is a bit outdated. 

I have the \`@token\`  of the user, but I'm not sure how to use it to authorize the service.

    #What data comes back from OmniAuth?     
        @auth = request.env["omniauth.auth"]
        #Use the token from the data to request a list of calendars
        @token = @auth["credentials"]["token"]
        client = Google::APIClient.new
        client.authorization.access_token = @token
        service = client.discovered_api('calendar', 'v3')
        @result = client.execute(
          :api_method =&gt; service.calendar_list.list,
          :parameters =&gt; {},
          :headers =&gt; {'Content-Type' =&gt; 'application/json'})

Any ideas? Or do any of you know the more updated code for this? 

Thanks so much!
## [11][Build a Reddit Clone in Rails 6 [PART 3 - BUILD COMPLETE]](https://www.reddit.com/r/rails/comments/f0uq9t/build_a_reddit_clone_in_rails_6_part_3_build/)
- url: https://www.reddit.com/r/rails/comments/f0uq9t/build_a_reddit_clone_in_rails_6_part_3_build/
---
Hi guys, I've recently released part 3 of the new Reddit clone build in Ruby on Rails. I shared part 1 and 2 here over the past 2 weeks for the community to check out. Those first 2 videos got some interest, so this latest release might be useful to anyone who has enjoyed the series so far.

**Part 1:**

In this video I cover the setup of a new rails app, installing the devise gem and creating user accounts. Add functionality for users to create a new community/subreddit. Add functionality to create new posts for a community (using nested routes).

[https://www.youtube.com/watch?v=aD6JvHKNPPM](https://www.youtube.com/watch?v=aD6JvHKNPPM)

**Part 2:**

Add dropdown nav menu allowing users to view their profile, edit account settings and sign out. Extend devise to add first name, last name and username. Signed in users can subscribe to communities and create new posts within them. Add profile for users which shows some basic details and their recent posts. Add new comments table and allow subscribers to add comments to posts. Improve layout and design of our app frontend.

[https://www.youtube.com/watch?v=kSj3pCT3r6Y](https://www.youtube.com/watch?v=kSj3pCT3r6Y)

**Part 3:**

Add ability to upvote / downvote posts and order them in their community based on popularity. Add user karma based on number of upvotes received on posts.

[https://www.youtube.com/watch?v=Rsqj34unV1c](https://www.youtube.com/watch?v=Rsqj34unV1c)

The build is part of a collection of MVP's that I have been building after working in the tech startup space for some time and assisting in building minimum viable products.

I post new web app builds on YouTube using Ruby on Rails 6 so feel free to check out the channel and add some recommendations for upcoming builds.
