# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/ezrfed/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/ezrfed/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][what is the rails way to use library that need instance?](https://www.reddit.com/r/rails/comments/f1633k/what_is_the_rails_way_to_use_library_that_need/)
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
## [3][Google Calendar API authorization help](https://www.reddit.com/r/rails/comments/f14fxe/google_calendar_api_authorization_help/)
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
## [4][How can i completely uninstall Ruby and Ruby on Rails completely from Windows to do a new installation](https://www.reddit.com/r/rails/comments/f18i86/how_can_i_completely_uninstall_ruby_and_ruby_on/)
- url: https://www.reddit.com/r/rails/comments/f18i86/how_can_i_completely_uninstall_ruby_and_ruby_on/
---
How can i completely uninstall Ruby and Ruby on Rails completely from Windows to do a new installation, there may be some issues of my previous installation.
## [5][Build a Reddit Clone in Rails 6 [PART 3 - BUILD COMPLETE]](https://www.reddit.com/r/rails/comments/f0uq9t/build_a_reddit_clone_in_rails_6_part_3_build/)
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
## [6][How can I document my API with swagger? Please help!](https://www.reddit.com/r/rails/comments/f126cn/how_can_i_document_my_api_with_swagger_please_help/)
- url: https://www.reddit.com/r/rails/comments/f126cn/how_can_i_document_my_api_with_swagger_please_help/
---
Hello everyone. I'm currently doing my engineer internships in a small company. However, as a project I was assigned to document the whole API using swagger in order to test its methods. I found [this](https://medium.com/@sushildamdhere/how-to-document-rest-apis-with-swagger-and-ruby-on-rails-ae4e13177f5d) tutorial, and also [this](https://medium.com/wolox-driving-innovation/lets-forget-painful-api-documentation-f5d0f5d0d06d). But I didn't have luck making it to work. I don't have too much experience with Ruby, so Rails is somehow foreign to me, and I feel lost and frustrated. There's any other alternative to swagger? If anyone is willing to help me, I would truly appreciate it!
## [7][Having trouble with bundle install](https://www.reddit.com/r/rails/comments/f16d3h/having_trouble_with_bundle_install/)
- url: https://www.reddit.com/r/rails/comments/f16d3h/having_trouble_with_bundle_install/
---
It's telling me I can't install kgio... But doesn't work?? This is code I pulled from github. They tell me to do a rake install, which errors and tells me to do a "bundle install" which is where this error is calming from.
## [8][Question about Google Client API OAuth](https://www.reddit.com/r/rails/comments/f11qol/question_about_google_client_api_oauth/)
- url: https://www.reddit.com/r/rails/comments/f11qol/question_about_google_client_api_oauth/
---
So, I used this page to incorporate the calendar api into my rails app:

https://developers.google.com/calendar/quickstart/ruby

But they put all the code into one file, so I wanted to break it up

I basically created a GoogleApi service module, and placed the `authorize` method into it.

And I created a `google_controller.rb` controller, and an action in it called `authenticate`, which will run that `authorize` method for the service module, and then try the rest of the code in the file provided by the link.

I don't get an error, but my entire server stops and gives me a link to click on that will take me to the accounts log in page. Which is great!

However, because my server hangs, there's no way to send that link to the front end (this is an api only rails app) and also, no one else can send http requests to the api because the server is hung.

It's a pretty funny situation, but a confusing one. Any ideas?

I've tried running the authenticate code in a separate thread so that it wouldn't run in the main thread and stop my server, however, doing so made me exceed the daily limit for google api requests, so that didn't work.

I really appreciate any help. If you need me to include any files, I'd be super happy to. I haven't included any code because I don't think this is a code issue so much as a "where should I put the code" kind of issue.
## [9][What datepicker do you use in your projects?](https://www.reddit.com/r/rails/comments/f0nxkp/what_datepicker_do_you_use_in_your_projects/)
- url: https://www.reddit.com/r/rails/comments/f0nxkp/what_datepicker_do_you_use_in_your_projects/
---

## [10][rails_admin and cancancan - Handle Unauthorized Access](https://www.reddit.com/r/rails/comments/f0l9a2/rails_admin_and_cancancan_handle_unauthorized/)
- url: https://www.reddit.com/r/rails/comments/f0l9a2/rails_admin_and_cancancan_handle_unauthorized/
---
Hello guys! got a new problem. I am using Rails [6.0.2.1](https://6.0.2.1).

In my project I use rails\_admin with cancancan. Naturally I need to handle access errors from cancan.

[https://github.com/sferik/rails\_admin/wiki/Cancancan](https://github.com/sferik/rails_admin/wiki/Cancancan) Here there is a solution:

" Also make sure RailsAdmin is inheriting from ApplicationController:

    # in config/initializers/rails_admin.rb
    
    config.parent_controller = 'ApplicationController'

"

I done it. In case of denied access to /admin I get this error:

## No route matches {:action=&gt;"index", :controller=&gt;"welcome"}

    class ApplicationController &lt; ActionController::Base
      protect_from_forgery with: :exception
    
      before_action :authenticate_user!
    
      rescue_from CanCan::AccessDenied do |exception|
        respond_to do |format|
          format.json { head :forbidden }
          format.html { redirect_to welcome_index_path, alert: exception.message }
        end
      end
    end

Here `routes.rb`

    Rails.application.routes.draw do
      mount RailsAdmin::Engine =&gt; '/admin', as: 'rails_admin'
      devise_for :users
      root to: 'home#index'
      get 'welcome/index'
      resources :companies
      resources :issues
      resources :buildings
      resources :confirmations
    end
## [11][Noob question about work with references](https://www.reddit.com/r/rails/comments/f0jfmw/noob_question_about_work_with_references/)
- url: https://www.reddit.com/r/rails/comments/f0jfmw/noob_question_about_work_with_references/
---
I have model `building` with migration:

    class CreateBuildings &lt; ActiveRecord::Migration[6.0]
      def change
        create_table :buildings do |t|
          t.string :name
          
          t.timestamps
        end
      end
    end

And model `issue` with migration:

    class CreateIssues &lt; ActiveRecord::Migration[6.0]
      def change
        create_table :issues do |t|
          t.string :title
          t.text :desc
          t.references :user, null: false, foreign_key: true
          t.references :building, null: false, foreign_key: true
          t.references :company, null: false, foreign_key: true
          t.date :date, default: Time.zone.today
          t.date :deadline
          t.text :report, default: '--put here your answer report--'
          t.string :status, default: '--waiting for response--'
    
          t.timestamps
        end
      end
    end

I need to get building's name by using foreign key in form `show.html.erb` coresponding to `issues controller`.

here is a  `IssuesController` method `show`:

      def show
        @issue = Issue.find(params[:id])
      end

* I tried to do this way:

&amp;#8203;

    &lt;p&gt;
      &lt;strong&gt;Building:&lt;/strong&gt;
      &lt;%= @issue.building_id.name %&gt;
    &lt;/p&gt;

But rails told that `name` is undefined method.

* In other way I trried to write a rails query:

&amp;#8203;

    &lt;p&gt;
      &lt;strong&gt;Building:&lt;/strong&gt;
      &lt;%= Building.select(:name).where("id == @issue.building_id") %&gt;
    &lt;/p&gt;

The result was:  #&lt;Building::ActiveRecord\_Relation:0x000000000e47e588&gt; 

change 

    &lt;%= Building.select(:name).where("id == u/issue.building_id") %&gt;

by

    &lt;%= Building.select(:name).where("id == @issue.building_id").reload %&gt;

trigger this exception:  

    SQLite3::SQLException: near ".": syntax error
