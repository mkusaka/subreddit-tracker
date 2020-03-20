# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/fgx7fz/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/fgx7fz/personal_projects_show_off_your_own_project_andor/
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
## [2][Postgres client / server version mismatch](https://www.reddit.com/r/rails/comments/flkqlo/postgres_client_server_version_mismatch/)
- url: https://www.reddit.com/r/rails/comments/flkqlo/postgres_client_server_version_mismatch/
---
I’m toying with Rails in containers using [this page](https://docs.docker.com/compose/rails/) as my guide.

I picked the latest Ruby, the latest Debian, and the latest Postgres (12) and bumped into a problem.

I create a table and then from the Rails container, go into “rails dbconsole” and do “\d accounts” (my table) and I get [this error](https://stackoverflow.com/questions/58461178/how-to-fix-error-column-c-relhasoids-does-not-exist-in-postgres)

Sure enough, “psql —version” from inside the Rails container says 11 while the DB container says 12.

My question is if I can somehow force the error to happen (and perhaps be more obvious) when Rails (using a version 11 library) connects to a version 12 DB server?
## [3][Alternative to Active Admin?](https://www.reddit.com/r/rails/comments/flj0e0/alternative_to_active_admin/)
- url: https://www.reddit.com/r/rails/comments/flj0e0/alternative_to_active_admin/
---
The app I am on is using ActiveAdmin for backend administration. It's been very painful to use lately and I was wondering which alternative to go after and why. There are so many options, but not all of them are maintained or even lightweight.
## [4][Get DB data ID in model class](https://www.reddit.com/r/rails/comments/flow2j/get_db_data_id_in_model_class/)
- url: https://www.reddit.com/r/rails/comments/flow2j/get_db_data_id_in_model_class/
---
Hello everyone! I have two almost equal tables in DB. I need to add a callback on update in 1 table will update same column or columns in other table.
So here is a question. How can I get the row ID in model class file?
## [5][devise changing routes to confirmation.user](https://www.reddit.com/r/rails/comments/flh8ym/devise_changing_routes_to_confirmationuser/)
- url: https://www.reddit.com/r/rails/comments/flh8ym/devise_changing_routes_to_confirmationuser/
---
Hi! I have this strange bug in my application that devise is changing the routes .... :S 

e.g when submitting the form here [http://localhost:3000/users/confirmation/new](http://localhost:3000/users/confirmation/new) it is sending the user to  [http://localhost:3000/confirmation.user](http://localhost:3000/confirmation.user)  :S 

and when signing up i get this mail

&lt;p&gt;&lt;a href="[http://localhost:3000/confirmation.16?confirmation\_token=yaQYvdu-ysiC1rkqtsxC](http://localhost:3000/confirmation.16?confirmation_token=yaQYvdu-ysiC1rkqtsxC)"&gt;Confirm my account&lt;/a&gt;&lt;/p&gt;

notice the .16 (this is the user id) in the url?

User.rb

`class User &lt; ApplicationRecord`

  `# Include default devise modules. Others available are:`

  `# :confirmable, :lockable, :timeoutable, :trackable and :omniauthable`

  `devise :database_authenticatable, :registerable,`

`:recoverable, :rememberable, :validatable , :confirmable`

&amp;#x200B;

  `has_many :products, dependent: :destroy`

  `has_many :trackers, dependent: :destroy`

  `validates_presence_of :firstname, :lastname`

  `has_one :plan, dependent: :destroy`

`end` 

&amp;#x200B;

routes.rb

`Rails.application.routes.draw do`

  `mount RailsAdmin::Engine =&gt; '/admin', as: 'rails_admin'`

&amp;#x200B;

`authenticate :user, lambda { |u| u.admin? } do`

  `require 'sidekiq/web'`

  `require 'sidekiq/cron/web'`

  `mount Sidekiq::Web =&gt; '/sidekiq', as: "sidekiq"`

  `end`

&amp;#x200B;

  `resources :trackers`

  `get 'home/index'`

  `resources :products do`

`resources :trackers`

  `end`

  `devise_for :users`

  `root to:'pages#index'`

  `get 'support', to: "pages#support"`

  `get 'confirmation', to: "pages#confirmation"`

  `get 'upgrade', to: 'pages#upgrade'`

  `get 'cancel', to: 'pages#cancel'`

  `get 'payments', to: 'pages#payments'`

  `post 'upgrade', to: 'pages#upgrade_product_limit'`

  `post 'cancel', to: 'pages#cancel_plan'`

&amp;#x200B;

  `#post 'payments', to: 'pages#list_subscriptions'`

`end`

What can be the issue?
## [6][Encrypted Chat App](https://www.reddit.com/r/rails/comments/flgbg7/encrypted_chat_app/)
- url: https://www.reddit.com/r/rails/comments/flgbg7/encrypted_chat_app/
---
I haven't worked with encryption/security much so I thought of a small demo app to create. I was thinking something like slack but the centralized server/webmasters have zero access to the messages. 

I wanted to ask what would be the best way to approach this. Do the messages need to be digested client side to prevent the server from receiving the clear text message? How would I go about secret keys for each user? 

I'd like some method that the server has literally zero ways to decode messages. Any tips, tricks or tutorials are totally welcomed! Thanks 😁
## [7][Social Media Platform](https://www.reddit.com/r/rails/comments/fl9vzd/social_media_platform/)
- url: https://www.reddit.com/r/rails/comments/fl9vzd/social_media_platform/
---
Hey guys,   
I'm working on a project using the Rails WAF. I'm relatively new to this and am getting lost here and there so please be patient.  


I'm looking to create a social media platform for a college project where users can sign up and organise events with one another.  


So far I have followed a few tutorials online and have been able to successfully implement a functioning devise based platform, which allows users to sign-up/sign-in and create 'events' which can be seen by other users and have basic CRUD functionality for the user who created the event.  


I am now looking to turn my attention to user profiles. From what I have gathered from my research it is okay to used the same database table to populate a user profile that I have used to create user login.  


Can you please advise the best way to create user profiles which allows the logged in user to flesh out their account with personal details which can be viewed by other users?
## [8][Rails Logging Help](https://www.reddit.com/r/rails/comments/fl31hn/rails_logging_help/)
- url: https://www.reddit.com/r/rails/comments/fl31hn/rails_logging_help/
---
I've set this in my development.rb and production.rb but the timestamp only shows up in dev. How do I set my logger to display correctly in PROD?

`logger = ActiveSupport::Logger.new(STDOUT)`  
  `logger.formatter = proc { |severity, datetime, progname, msg| "[#{severity}, #{datetime}] -  #{msg}\n" }`  
  `config.logger = logger`

Dev:

`[INFO, 2020-03-18 22:11:03 -0500] -  Redirected to` [`http://localhost:3000/`](http://localhost:3000/)

&amp;#x200B;

Prod:

`[8c173dc6-7acf-4b57-95e0-a8d6c610d3ee] Completed 302 Found in 128ms (ActiveRecord: 10.0ms | Allocations: 12865)`
## [9][Improving performance with Presenter and Eager Load of a big Rails app](https://www.reddit.com/r/rails/comments/fkvtdz/improving_performance_with_presenter_and_eager/)
- url: https://www.reddit.com/r/rails/comments/fkvtdz/improving_performance_with_presenter_and_eager/
---
Hey guys, I have written a post about how I helped to improve the performance of a Ruby on Rails application where I work.

The app is used by people all over the world and has around of 4 million request every month.

https://medium.com/@wlss.lucas/improving-performance-of-a-rails-app-12beb830319a
## [10][Best pattern for ActiveRecord Create objects](https://www.reddit.com/r/rails/comments/fl2103/best_pattern_for_activerecord_create_objects/)
- url: https://www.reddit.com/r/rails/comments/fl2103/best_pattern_for_activerecord_create_objects/
---
I have a business model, for example *CheckForWidgetOrder.*

Code is similar to:

    class CheckForWidgetOrder
      def check(user, widget_name, widget_area)
        if (user.on_west_coast)
          WidgetOrder.create(name: widget_name, area: widget_area)
        end
      end 
    end

Is this the right place for the WidgetOrder.create?  It seems like I am littering the business logic with ActiveRecord syntax...

I could inject the ActiveRecord model and still do the create, or is the preferred way to do the create action is to use a WidgetCreator class with the ActiveRecord call in there?  Or is that overkill?
## [11][[HELP] Restrict browsers mode on websites](https://www.reddit.com/r/rails/comments/fkw3kv/help_restrict_browsers_mode_on_websites/)
- url: https://www.reddit.com/r/rails/comments/fkw3kv/help_restrict_browsers_mode_on_websites/
---
Hi guys, this question is not related to Rails, but with websites in general. I work with a website that the browsers change to a restrict mode after you log in on the restrict area . What I mean with "restrict mode" is a browser with static address bar (users can't change), no button options (bookmark, reload, back/forward page, etc.). You basically have a static address bar and the rendered html page. 

My question is: why they use this restrict mode? I think is because of the back/forward page and the static address bar and how the website caches, or how they load Javascript on page, or because they don't want users reloading pages as the website load a lot of data on certain pages, or because they don't want users open more than one tab per user logged in, I don't know, but I was kind of curious.

Sorry if this question don't fit this sub, but I didn't know where to ask.
