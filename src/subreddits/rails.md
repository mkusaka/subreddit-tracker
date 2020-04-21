# rails
## [1][Looking for a rails front-end framework](https://www.reddit.com/r/rails/comments/g5ch3n/looking_for_a_rails_frontend_framework/)
- url: https://www.reddit.com/r/rails/comments/g5ch3n/looking_for_a_rails_frontend_framework/
---
I am developing a new app from scratch in rails 6 and I'm looking for a simple frontend framework to:

1.  Allow users to access the app offline and
2.  Access some native features of the phone: like Geolocation, Camera and Bluetooth.*

Indeed, I'm looking for a frontend framework to keep the architecture as simple as possible and avoid having two code-bases.

The app itself is not very complex (think of a "listing app", similar to AirBnb) but given the targeted users, I would like to have strong control over the design. 

I should underline that I am a hobbyist developer (in my last project I used rails 3 --yep, that's a long time ago! ; ) so I would prefer to avoid learning "yet another framework". For that reason, I would like to use html 5 / css as much as possible, in order to keep the development experience as close as possible to a simple web-app.


Question: what frontend framework and / or approach would answer to these requirement?
## [2][would appreciate hearing your thoughts](https://www.reddit.com/r/rails/comments/g5f59c/would_appreciate_hearing_your_thoughts/)
- url: https://www.reddit.com/r/rails/comments/g5f59c/would_appreciate_hearing_your_thoughts/
---
I have database where i think it not well designed and many problems will show up in future, it like that : a polymorphic relation between "Claim" and "wgclaim" and "rtaclaim" they are a types of this claim so the claim is the father and wgclaim, rtaclaims are sons and soon future we can add more sons.
example : 
rtaclaim : has relation with : police, policy, car, passenger ...
wgclaim: has relation with : plan 
and both share claimable : reference , status, user, team ...
do you think the polymorphic is the right path ? or there another way ??
## [3][Controller is not creating my ActiveRecord models](https://www.reddit.com/r/rails/comments/g5b5i4/controller_is_not_creating_my_activerecord_models/)
- url: https://www.reddit.com/r/rails/comments/g5b5i4/controller_is_not_creating_my_activerecord_models/
---
Hi all, I am having troubles creating a model in my payments controller for a stripe checkout. Some background info, I am just testing around with the stripe integration and I want to create a PaymentSession model which is a regular ActiveRecord model right before creating my Stripe Checkout session. This is to confirm payments are completed via stripe webhooks. Everything is working fine except for creating this PaymentSession record. Can someone explain how this works, probably has something do with the create method, right? 

This is my code in the controller and this line does not work: 

    @paymentsession = @product.payment_sessions.build(stripe_checkout_session_id: "1234")

&amp;#x200B;

    class PaymentsController &lt; ApplicationController
    
        def create
          @product = Product.find(params[:id])
          ---- Below is not working------
          @paymentsession = @product.payment_sessions.build(stripe_checkout_session_id: "1234")
          -------------------------------
          application_fee = ((@product.price / 100) * 8 + 0.30).round
          @session = Stripe::Checkout::Session.create({
            payment_method_types: ['card', 'ideal'],
            line_items: [{
              name: @product.title,
              amount: @product.price,
              images: ['https://picsum.photos/400'],
              currency: 'eur',
              quantity: 1,
              
            }],
            payment_intent_data: {
              application_fee_amount: application_fee,
            },
            success_url: payments_success_url,
            cancel_url: 'https://example.com/cancel',
          }, {stripe_account: @product.user.stripe_user_id})
            
            respond_to do |format| 
              format.js
            end
    
        end
    
        def success
        end
    
        def cancel
        end
    
    end
## [4][update_without_password send re-confirmation email](https://www.reddit.com/r/rails/comments/g4w8sw/update_without_password_send_reconfirmation_email/)
- url: https://www.reddit.com/r/rails/comments/g4w8sw/update_without_password_send_reconfirmation_email/
---
I would like to resend a confirmation e-mail, if the user change the email,  after call update\_without\_password.

Today, i make this in code:

 `def update_resource(resource, params)`  
 `if params[:password].present? &amp;&amp; params[:password_confirmation].present?`  
 `if params[:password] == params[:password_confirmation]`  
        `resource.update(password: params[:password])`  
        `resource.save`  
 `end`  
      `edit_user_registration_path(resource)`  
 `else`  
      `result = resource.update_without_password(params)`  
 `end`  
 `end`
## [5][Webpacker : how to hot-reload HTML ?](https://www.reddit.com/r/rails/comments/g4p458/webpacker_how_to_hotreload_html/)
- url: https://www.reddit.com/r/rails/comments/g4p458/webpacker_how_to_hotreload_html/
---
With Docker +  Webpacker, it is possible to hot-reload JS, it is as simple as set hmr: true in the webpacker.yml file.

It is also possible to hot-reload CSS, even if they are sprockets-generated : a few configuration in development.js (as some listed here [https://github.com/rails/webpacker/issues/1879](https://github.com/rails/webpacker/issues/1879)) and you're done.

But I'm still struggling to hot-reload HTML (in .html.erb files). In old days it was achieved thanks to guard, but I'm looking for a solution with webpacker to avoid a tool-stacking nightmare. How can it be achieved ?
## [6][Realtime, reactive web apps without Javascript using Stimulus Reflex](https://www.reddit.com/r/rails/comments/g48fm2/realtime_reactive_web_apps_without_javascript/)
- url: https://www.reddit.com/r/rails/comments/g48fm2/realtime_reactive_web_apps_without_javascript/
---
Thought you guys would be interested in this.

Stimulus Reflex lets you build realtime, reactive apps in Rails similar to Phoenix LiveView. You don't have to write any Javascript and it will use Rails to render all the HTML updates server-side. It uses ActionCable to trigger updates and then tells Rails to render the current page again and send that back over the websocket. It'll then use DOM diffing to update the page automatically for you.

Definitely one of the coolest projects going on right now in the Rails world I'd say. It looks like we might see some similar things in Turbolinks 6 when that comes out too which is exciting.

[https://gorails.com/episodes/stimulus-reflex-basics?autoplay=1](https://gorails.com/episodes/stimulus-reflex-basics?autoplay=1)
## [7][Having difficulty understanding `has_many :something, through: :something_else, source: :another_thing`, here is a good brief lesson](https://www.reddit.com/r/rails/comments/g4dgsj/having_difficulty_understanding_has_many/)
- url: https://www.reddit.com/r/rails/comments/g4dgsj/having_difficulty_understanding_has_many/
---
primarily recommend reading this lesson

https://open.appacademy.io/learn/full-stack-online/sql/associations--has_many--through

What's important is to note that _existing associations_ are what is being referred to in the `through` and `source` parts of the association, and they are not references to a particular model, but to associations/methods already declared.

take this code:

    class Physician &lt; ApplicationRecord
      has_many(
        :appointments,
        class_name: 'Appointment',
        foreign_key: :physician_id,
        primary_key: :id
      )

      has_many :patients, through: :appointments, source: :patient
    end

    class Appointment &lt; ApplicationRecord
      belongs_to(
        :physician,
        class_name: 'Physician',
        foreign_key: :physician_id,
        primary_key: :id
      )

      belongs_to(
        :patient,
        class_name: 'Patient',
        foreign_key: :patient_id,
        primary_key: :id
      )
    end

    class Patient &lt; ApplicationRecord
      has_many(
        :appointments
        class_name: 'Appointment',
        foreign_key: :patient_id,
        primary_key: :id
      )

      has_many :physicians, through: :appointments, source: :physician
    end


In `Physician`, the `has_many :patients, through: :appointments, source: :patient`:

1. the `:patients` will be a new method for your `Physician` model, `Physician#patients`.  

2. the `through: :appointments` is using your existing `Physician#appointments` method/association already declared in the `Physician` model of course.

3. the `source: :physician` is using an _existing_ association/method that is already declared too, namely, the join model `Appointments#physician` method/association in the `Appointments` model.

I struggle sometime to conceptualize this, above is a very brief synopsis, but I recommend going to the lesson, it's free.
## [8][GitHub - styd/pry-diff-routes: Inspect routes changes in Rails console.](https://www.reddit.com/r/rails/comments/g45csl/github_stydprydiffroutes_inspect_routes_changes/)
- url: https://www.reddit.com/r/rails/comments/g45csl/github_stydprydiffroutes_inspect_routes_changes/
---
Hi fellow Rails developer,

First of all, let's keep being safe and healthy and enjoy staying at home. Secondly, I want to say thank you to those of you who have contributed to my older gem, [ApexCharts.RB](https://github.com/styd/apexcharts.rb), by using, reporting issues, sending pull requests, or just giving the repo a star (I couldn't believe it at first but even [Chartkick](https://chartkick.com) author Andrew Kane and [ApexCharts](https://apexcharts.com) author Juned Chhipa also gave a star. Thank you!). I sincerely appreciate all of your efforts.

Today, I want to tell you about a new gem that I've been working on since last month called [PryDiffRoutes](https://github.com/styd/pry-diff-routes). It's a Pry plugin that works in Rails console to see the route changes that we make.

The way it works is that it would save the current routes as Rails console is starting and then from that point on the routes you change -- don't forget to save the route file -- is comparable with the routes when you started the Rails console. When you type the Pry command `diff-routes`, it will output your changes as removed, modified, or new routes.

What's "modified"? And why don't I just display the **before** and the **after**?
IMHO, when we talk about an endpoint, what we really meant usually is a specific combination of a _verb_ and a _path_. When they don't change, we consider it as the same endpoint. By changing either the _verb_ or the _path_, we can say that we are removing the old route and creating a new route. When we just change where it routes to, we are modifying the route. So, what I'm doing is just group routes changes based on our intention for those changes.

I hope it is useful for you and let me know what you think about it.
Happy weekend!
## [9][Can I use Active Storage to store files in a user's personal storage service?](https://www.reddit.com/r/rails/comments/g4c9se/can_i_use_active_storage_to_store_files_in_a/)
- url: https://www.reddit.com/r/rails/comments/g4c9se/can_i_use_active_storage_to_store_files_in_a/
---
I would like to entertain the idea of "not owning" the users' uploaded files. Some applications store the data or files related to your account in your own Google Drive etc. Is this possible with Active Storage?
## [10][Need help properly configuring my database.yml](https://www.reddit.com/r/rails/comments/g4gzth/need_help_properly_configuring_my_databaseyml/)
- url: https://www.reddit.com/r/rails/comments/g4gzth/need_help_properly_configuring_my_databaseyml/
---
Aloha new rails jr dev here. Ive been stressing for days trying to run this repo: [https://github.com/chinweenie/Full-Stack-Project-Etsy](https://github.com/chinweenie/Full-Stack-Project-Etsy)

i have learned A LOT so far , but i am struggling to with my database.yml configuration, here is my code : [https://pastebin.com/YkhkhYwb](https://pastebin.com/YkhkhYwb)

and here is the error im getting in the browser: Puma caught this error: Cannot load database configuration: undefined method \`each' for "gem 'pg'":String (NoMethodError)

ive busted my as to get this far, here is the guide that helped me the most : [https://vitux.com/how-to-install-latest-ruby-on-rails-on-ubuntu/](https://vitux.com/how-to-install-latest-ruby-on-rails-on-ubuntu/)

i know im close  can anyone take the time and patience to help me run this app ? 

mahalo
