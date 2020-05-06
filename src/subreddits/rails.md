# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/g616hm/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/g616hm/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gek2rk/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gek2rk/personal_projects_show_off_your_own_project_andor/
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
## [3][RSpec and Minitest](https://www.reddit.com/r/rails/comments/gedoal/rspec_and_minitest/)
- url: https://www.reddit.com/r/rails/comments/gedoal/rspec_and_minitest/
---
Do you create RSpec and Minitest in your app? or you choose only one?
## [4][Database design for rails application](https://www.reddit.com/r/rails/comments/gefv84/database_design_for_rails_application/)
- url: https://www.reddit.com/r/rails/comments/gefv84/database_design_for_rails_application/
---
This is my [database design](https://dbdiagram.io/d/5ea1b31e39d18f5553fe19e6). I am not sure if my associations are correct. Should I use composite keys or foreign keys? Can you guys tell me what is wrong with my database design. Thankss
## [5][Creating model to controller to view or create another model and make associations](https://www.reddit.com/r/rails/comments/geatl7/creating_model_to_controller_to_view_or_create/)
- url: https://www.reddit.com/r/rails/comments/geatl7/creating_model_to_controller_to_view_or_create/
---
Hello there, I am done creating my unit testing for my first created model. What to do next? Do I need to create the next model and associations of each or I will continue build my model until controller and view?
## [6][Saving JSON data from shopify_app Webhook?](https://www.reddit.com/r/rails/comments/ge4wzn/saving_json_data_from_shopify_app_webhook/)
- url: https://www.reddit.com/r/rails/comments/ge4wzn/saving_json_data_from_shopify_app_webhook/
---
Hi -

I am somewhat of a novice as it relates to API's and JSON.

**MY Goal:**

I am playing around with the shopify\_app gem, and am having trouble saving JSON data to my local db.

**For some background on my setup:**

1. Created an embeddable shopify app using the [shopify\_app gem](https://github.com/shopify/shopify_app)
2. Have a webhook for orders/create that I generated in the config/initializers/shopify\_app.rb file
3. Created a job that works with the webhook. Whereas whenever a shopify store that has my app creates an order, shopify sends over the order data via JSON to my webhook

app/controllers/webhooks\_controller.rb

    class WebhooksController &lt; ApplicationController
        include ShopifyApp::WebhookVerification
        skip_before_action :verify_authenticity_token, raise: false
    
        def orders_create
          params.permit!
          OrdersCreateJob.perform_now(shop_domain: shop_domain, webhook: webhook_params.to_h)
          head :no_content
        end
    
        def orders_create_save
          @order = Order.new(shopify_order_id: webhook_params[:webhook][:id])
          @order.save
        end
    
        private
          def webhook_params
            params.except(:controller, :action, :type)
          end
    end

app/jobs/orders\_create\_job.rb

    class OrdersCreateJob &lt; ActiveJob::Base
      def perform(shop_domain:, webhook:)
        shop = Shop.find_by(shopify_domain: shop_domain)
    
        if shop.nil?
          logger.error("#{self.class} failed: cannot find shop with domain '#{shop_domain}'")
          return
        end
    
        shop.with_shopify_session do
        end
      end
    end

config/initializers/shopify\_app.rb

    ShopifyApp.configure do |config|
      config.application_name = "My Shopify App"
      config.api_key = ENV['SHOPIFY_API_KEY']
      config.secret = ENV['SHOPIFY_API_SECRET']
      config.old_secret = ""
      config.scope = "read_products, read_orders, read_customers" 
      config.embedded_app = true
      config.after_authenticate_job = false
      config.api_version = "2020-04"
      config.shop_session_repository = 'Shop'
      config.webhooks = [
        {topic: 'orders/create', address: 'https://12345.ngrok.io/webhooks/orders_create', format: 'json'},
        # {topic: 'orders/paid', address: 'https://12345.ngrok.io/webhooks/orders_paid', format: 'json'},
        # {topic: 'orders/create', address: 'https://12345.ngrok.io/webhooks/orders_create'}
      ]
    end
    
    ShopifyApp::Utils.fetch_known_api_versions                        # Uncomment to fetch known api versions from shopify servers on boot
    ShopifyAPI::ApiVersion.version_lookup_mode = :raise_on_unknown    # Uncomment to raise an error if attempting to use an api version that was not previously known

config/routes.rb

    Rails.application.routes.draw do
      mount ShopifyApp::Engine, at: '/'
    
      root :to =&gt; 'home#index'
      post '/webhooks/orders_create', :to =&gt; 'webhooks#orders_create'
      get '/webhooks/orders_create', to: 'webhooks#orders_create_save'
    end

Any thoughts on this would be helpful. Thank you very much!

**Edit:**

Order (belongs\_to :shop) schema:

      create_table "orders", force: :cascade do |t|
        t.string "shopify_order_id", null: false
        t.string "shopify_order_name", default: ""
        t.datetime "shopify_order_created_at"
        t.integer "shop_id"
        t.datetime "created_at", precision: 6, null: false
        t.datetime "updated_at", precision: 6, null: false
        t.index ["shop_id"], name: "index_orders_on_shop_id"
        t.index ["shopify_order_id"], name: "index_orders_on_shopify_order_id", unique: true
      end

&amp;#x200B;
## [7][Nested form not showing up.](https://www.reddit.com/r/rails/comments/ge3imk/nested_form_not_showing_up/)
- url: https://www.reddit.com/r/rails/comments/ge3imk/nested_form_not_showing_up/
---
Hi, I'm currently working on a project and possible employee is using to see if they'll hire me. I've pretty much finished the test with time to spare and am just trying to add some extra fluff.

The whole codebase is [here](https://github.com/Kutomore/music-event-application) but I'll write the relevant info now.

Basically I'm trying to setup the registration form so the user will both create his account and his profile at the same time, which in theory should be pretty simple.

But no matter what I do, I can't get the "profile" for to show up on the registration form.

&lt;h2&gt;Sign up&lt;/h2&gt;    


    &lt;%= form_for(resource, as: resource_name, url: registration_path(resource_name)) do |f| %&gt;
      &lt;%= render "devise/shared/error_messages", resource: resource %&gt;
      &lt;div class="field"&gt;
        &lt;%= f.label :email %&gt;&lt;br /&gt;
        &lt;%= f.email_field :email, autofocus: true, autocomplete: "email" %&gt;
      &lt;/div&gt;
      &lt;div class="field"&gt;
        &lt;%= f.label :password %&gt;
        &lt;% if u/minimum_password_length %&gt;
        &lt;em&gt;(&lt;%= u/minimum_password_length %&gt; characters minimum)&lt;/em&gt;
        &lt;% end %&gt;&lt;br /&gt;
        &lt;%= f.password_field :password, autocomplete: "new-password" %&gt;
      &lt;/div&gt;
      &lt;div class="field"&gt;
        &lt;%= f.label :password_confirmation %&gt;&lt;br /&gt;
        &lt;%= f.password_field :password_confirmation, autocomplete: "new-password" %&gt;
      &lt;/div&gt;
      &lt;%= f.fields_for :profile_attributes,@profile = Profile.new do |ff|%&gt;
        &lt;div class="field"&gt;
          &lt;%= ff.text_field :name, placeholder: 'Name' %&gt;
        &lt;/div&gt;
        &lt;div class="field second-step"&gt;
          &lt;%= ff.text_field :email, placeholder: 'Email' %&gt;
        &lt;/div&gt;
        &lt;div class="field second-step"&gt;
          &lt;%= ff.text_field :phone, placeholder: 'Phone' %&gt;
        &lt;/div&gt;
        &lt;div class="field second-step"&gt;
          &lt;%= ff.number_field :gender, placeholder: 'Gender' %&gt;
        &lt;/div&gt;
        &lt;div class="field second-step mb-5"&gt;
          &lt;%= ff.text_field :birthdate, placeholder: 'Birthdate' %&gt;
        &lt;/div&gt;
      &lt;% end %&gt;
      &lt;div class="actions"&gt;
        &lt;%= f.submit "Sign up" %&gt;
      &lt;/div&gt;
    &lt;% end %&gt;
    
    &lt;%= render "devise/shared/links" %&gt;

This is the whole code for the view. User belongs\_to profile and accepts\_nested\_attributes for it. I've tried building the relationship with `.build_profile` by overriding the devise controller in the routes. I've tried setting it up as it currently is `@profile =` [`Profile.new`](https://Profile.new). But the problem persists, no matter what I do the form won't show up, only the user fields are visible.

Anyone got any ideas what might be causing this? (Ruby 2.6.3 with Rails 6.0.2)

&amp;#x200B;

edit 1 - I've again tried with the `resource.build_profile` setup, to no avail, but with a pry in the controller I was able to confirm that the profile is indeed being built on the controller, but still, it doesn't show up.

&amp;#x200B;

edit 2 - Found the issue, basically the controller was setup like 

      def new
        super
        resource.build_profile
      end

Which was rendering the view before I was able to build the profile, I was using super since I didn't want to overwrite the default behaviour, but it seems that I had to, anyway, thanks ppl.
## [8][How do you make the routing as private?](https://www.reddit.com/r/rails/comments/geawq5/how_do_you_make_the_routing_as_private/)
- url: https://www.reddit.com/r/rails/comments/geawq5/how_do_you_make_the_routing_as_private/
---
Hello, I have a case here 
User has many Rooms , and Roms belongs to User
In the Rooms table I have a column PIN 
So when people into that room, people must insert the PIN into that Room
In my routes.rb I put resources :rooms 

So it gonna have route "localhost:3000/rooms/:id"

If user know the Id of the rooms or just put dummy I'd in that routes, they might be able to access that Room 

So how to make that Room only be able if people insert PIN into that Rooms ?

Should we add more PIN params =  "localhost:3000/rooms/:id/:pin"  ? 

Would it gonna make SQL injection ? To believe of that params on that routes,?
## [9][Anyone here using tailwindui? Thinking of pulling the trigger](https://www.reddit.com/r/rails/comments/gdn2w7/anyone_here_using_tailwindui_thinking_of_pulling/)
- url: https://www.reddit.com/r/rails/comments/gdn2w7/anyone_here_using_tailwindui_thinking_of_pulling/
---
I really like tailwindcss, and im thinking of getting tailwindui. It would save me alot of time in the design process. Anyone working with tailwindui? How has it changed your productivity?
## [10][Testing testing testing](https://www.reddit.com/r/rails/comments/gdo3kx/testing_testing_testing/)
- url: https://www.reddit.com/r/rails/comments/gdo3kx/testing_testing_testing/
---
Hi all, I'm fairly new to Rails but experienced in other languages.

Something I'm finding difficult is testing in Rails. Out of the box there are just so many types of tests and so I have so many questions.

1) MiniTest or RSpec? Every project I've seen in my career has used RSpec, do I really need to learn both?

2) Why so many types of spec? Request specs, controller specs, model specs, job specs, service specs, I've probably missed some! Do I need to learn them all? 

How? 

Not to sound petty but the RSpec documentation website is one of the worst websites I've ever encountered with regards to the volume of information on there (not much). I say this as someone who has had to learn and digest other testing frameworks in other languages such as PHP and JS.

3) Shoulda? Or shouldn't? I keep seeing Shoulda popup in projects but then friends tell me "DON'T USE THAT CRAP". Is it good to use or not? If not, why not?

4) Factory bot? Is this bad? Similar to Shoudla, people seem to dislike it.

5) Are there any decent resources out there with plenty of examples of how to write good tests, in the most recent syntax with modern tools? I keep seeing examples online with "the old RSpec syntax" showing something like `obj.stub()`

6) Do you all really, actually T,D,D? As in, you honestly, 100%, seriously write tests first, then code, then refractor, etc? Everyone I've met in my Rails career journey says they like to TDD but then in practice when pairing or whatever, they go guns blazing writing or debugging code without even looking at the tests until later.

7) Any other wisdom you can offer with regards to being a testing ninja? It's already confusing enough with all of the various _types_ of tests (unit, integration, feature, end to end, etc) but then Rails adds another level of complexity for me with all of the various styles and syntaxes and stuff.

Thanks in advance!
## [11][Verify identity of user](https://www.reddit.com/r/rails/comments/gdr2sq/verify_identity_of_user/)
- url: https://www.reddit.com/r/rails/comments/gdr2sq/verify_identity_of_user/
---
Hey guys... I’ve been trying to integrate an authentication system to verify if a given user is in fact a student from my university. I’m having a really hard time finding a good documentation on it. 

I looked into omniauth and it seems like the way to go but I’m super confused on how to integrate it with my university login. Does anyone here have any experience with integrating such a custom verification system?
## [12][How to Build a Twitter Clone with rails, ActionCable and React](https://www.reddit.com/r/rails/comments/gdjgd6/how_to_build_a_twitter_clone_with_rails/)
- url: https://www.reddit.com/r/rails/comments/gdjgd6/how_to_build_a_twitter_clone_with_rails/
---
A week or so ago, a video about making a Twitter Clone with Rails CableReady and StimulusReflex hit the web. While I think it is really cool what those teams are doing, I am a little more old school (in JS years &lt;/sarcasm&gt;) and just use React.

&amp;#x200B;

I figured you all would like to see a tutorial on a clone of the clone with React and bare ActionCable

[https://robrace.dev/build-a-twitter-clone-with-rails-actioncable-and-react/](https://robrace.dev/build-a-twitter-clone-with-rails-actioncable-and-react/)
