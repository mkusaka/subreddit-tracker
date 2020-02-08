# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/ezrfed/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/ezrfed/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][What datepicker do you use in your projects?](https://www.reddit.com/r/rails/comments/f0nxkp/what_datepicker_do_you_use_in_your_projects/)
- url: https://www.reddit.com/r/rails/comments/f0nxkp/what_datepicker_do_you_use_in_your_projects/
---

## [3][rails_admin and cancancan - Handle Unauthorized Access](https://www.reddit.com/r/rails/comments/f0l9a2/rails_admin_and_cancancan_handle_unauthorized/)
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
## [4][Rails n00b, questions about authorization + URL schemes](https://www.reddit.com/r/rails/comments/f0fvsg/rails_n00b_questions_about_authorization_url/)
- url: https://www.reddit.com/r/rails/comments/f0fvsg/rails_n00b_questions_about_authorization_url/
---
Hello everyone,

I'm working on a Rails backend for a mobile app for event management, where mutiple different clients will be creating shows, events, etc, and accessing event data through Rails to the database. I want to make sure each client sees only their own data. Once the client authenticates, I can then query for events, then filter based on their client ID.

I'm thinking about a couple different url schemes:

`/event/18489202`

or

`/client/1/event/1`

In the first example, an event id is global across the application. In the second, an event is specific to the client. Or, I could have even this

`/event/1`

Out of the box, it would seem the first option is easiest to set up using a `resources` entry in `routes.rb`  . However I'm wondering if it would be more convenient, or aesthetically pleasing for clients, to use smaller client-specific numbers in the URL and then parse that on the backend for the actual database ID value. Plus I am a Rails noob so plenty of technical considerations maybe you all can help me with.

I'd be interested to hear your opinions. Thanks !
## [5][Noob question about work with references](https://www.reddit.com/r/rails/comments/f0jfmw/noob_question_about_work_with_references/)
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
## [6][Communication between scaffolds](https://www.reddit.com/r/rails/comments/f0iq7h/communication_between_scaffolds/)
- url: https://www.reddit.com/r/rails/comments/f0iq7h/communication_between_scaffolds/
---
I was thinking of some sort of "favorites", "wish list" or "cart" in a commerce website (btw, "favorites" can happen in a link sharing website as well :P). I mean we have two entities which actually define an online shopping website : Products and Users.

My question is, if we implement those two things, we need something else, a cart. I'm not going to dive deep into what actually should happen in a shop. Just wondering what happens when something enters to the cart. 

In cart, we have the product and its features and its quantity. For example, I'm going to amazon right now, and I choose a hair dryer. My mother may want one too. So, I increase the quantity. 

It means I need two scaffolds, right? I need one for product and one for the cart. But how can I make a communication between them?
## [7][Ruby on Rails advanced courses?](https://www.reddit.com/r/rails/comments/f00t14/ruby_on_rails_advanced_courses/)
- url: https://www.reddit.com/r/rails/comments/f00t14/ruby_on_rails_advanced_courses/
---
Hello, I consider I have experience in Rails, but I want to go to the next level, for this I want to know if there's an advanced course on Rails which has some or most of the following topics:

* Examples of good use of rails conventions (service objects, query object, null objects, creating concerns to DRY your controllers/models, etc.).
* Make performant apps (performant queries, performant ruby code, fix n+1 queries, etc.).
* Examples of good use of design patterns (observable, chain of responsibility, factory, strategy pattern, etc.).
* Serializing data (blueprinter, fast json api, etc.)
* Make performant tests (build or build_stubbed instead of create, parallel tests, etc.)

Or even building your own gems, metaprogramming, etc. (I know this is ruby, but to be an advanced Rails developer you have to be an advanced Ruby developer too IMO)
## [8][Issues with Stripe Connect... API saying customer + plan doesn't exist](https://www.reddit.com/r/rails/comments/f0b01n/issues_with_stripe_connect_api_saying_customer/)
- url: https://www.reddit.com/r/rails/comments/f0b01n/issues_with_stripe_connect_api_saying_customer/
---
I'm attempting to create a rails application with stripe connect to essentially act as middleman for services. With my current controller and JS, I have no problem charging a customer for a subscription for my platform. The issue becomes when I attempt to charge on behalf of a connected account as I'll typically get an error saying the customer doesn't exist (and when I update the connected account with customer object) it says the plan doesn't exist. I may have a misunderstanding of how connect works, but does the customer\_id and plan\_id also need to exist within the connected stripe accounts along with the application/platform's account? Attached a snippet of my controller below for reference. Problematic line is where subscription object is created.

 

`def create`  
 `#Make sure we change this to production when the time comes`  
 `Stripe.api_key = Rails.application.credentials.development[:stripe_api_key]`  
 `#Make sure that the credentials file has the appropriate plan_ids`  
`plan_id = params[:plan_id]`  
`plan = Stripe::Plan.retrieve(plan_id)`  
 `# flash[:warning] = plan`  
`token = params[:stripeToken]`  
 `# flash[:warning] = Stripe.api_key`  
 `#Let's add subscription value to the Library.`  
`subscription_plans = PlanType.all`  
 `#calling private function find_plan`  
`plan_type = find_plan(plan, subscription_plans)`  
`customer = if current_user.stripe_id.present?`  
 `Stripe::Customer.retrieve(current_user.stripe_id)`  
 `# flash[:danger] = "User already has a stripe ID!"`  
 `else`  
 `Stripe::Customer.create({`  
`email: current_user.email,`   
`source:token,`  
`}, {stripe_account: plan_type.stripe_id})`  
 `# Stripe::Customer.create(description: 'Test Customer')`  
 `#Save the stripe id to the database`  
 `end`  
 `#Update the subscription creation with stripe connected account param &amp; application_fee_percent params. Sent via connect`  
 `#transfer_data{amount_percent: 95, destination: plan_type.stripe_id }`  
`subscription = customer.subscriptions.create({plan: plan.id, application_fee_percent: 10,}, stripe_account: plan_type.stripe_id)`  
 `#Update the hash`  
`current_user.stripe_subscription_id[plan.nickname.downcase] = subscription.id`  
`options = {`  
`stripe_id: customer.id,`  
`subscribed: true`  
`}`  
`current_user.plan_subscription_library_additions &lt;&lt; plan_type`
## [9][There are two Ruby version manager RVM vs RBENV. what's your choice? and why?](https://www.reddit.com/r/rails/comments/f009mb/there_are_two_ruby_version_manager_rvm_vs_rbenv/)
- url: https://www.reddit.com/r/rails/comments/f009mb/there_are_two_ruby_version_manager_rvm_vs_rbenv/
---
I was at work, I installed ruby and rails using RBENV. 

My senior colleague made me install RVM and uninstalled RBENV. Insists that RVM handles better somehow. 

What's your thought on this? 

What is the industrial standard?
## [10][Yet another Arel Cheatsheet, but on Steroids](https://www.reddit.com/r/rails/comments/ezuuup/yet_another_arel_cheatsheet_but_on_steroids/)
- url: https://www.reddit.com/r/rails/comments/ezuuup/yet_another_arel_cheatsheet_but_on_steroids/
---
Hi guys,

I've created a cheatsheet of Arel that collects my personal experience with it. I've decided to share it since Arel is pretty much undocumented, but I think it's a really powerful tool.

[https://gist.github.com/ProGM/c6df08da14708dcc28b5ca325df37ceb](https://gist.github.com/ProGM/c6df08da14708dcc28b5ca325df37ceb)
## [11][Multiple database connections?](https://www.reddit.com/r/rails/comments/f02fxm/multiple_database_connections/)
- url: https://www.reddit.com/r/rails/comments/f02fxm/multiple_database_connections/
---
Hi. I have an ongoing project which asks for many DB connections (9 total) in:

-SQL Server
-MySQL
-Oracle SQL
-PostgreSQL

I never did something like this before, so i might ask for some advices or ideas.

I'm using Rails 6.

In resume, it's a data warehouse. I must delevelop an ETL module which process all the data from these 9 databases and puts into only one DB (which is PostgreSQL). I'd like to read your experiences, advices or something like this.
