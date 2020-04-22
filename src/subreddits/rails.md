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
## [2][How good is good enough for finding a rails job?](https://www.reddit.com/r/rails/comments/g5s5yd/how_good_is_good_enough_for_finding_a_rails_job/)
- url: https://www.reddit.com/r/rails/comments/g5s5yd/how_good_is_good_enough_for_finding_a_rails_job/
---
I started learning Rails a few years ago because I wanted to build an application that could automate a good chunk of someone's job in the company I work for. 

I started with very little programming knowledge. Fast forward 2 years later (not all build time) and I have an application with tests using RSpec, integrations with several API's, background jobs, role-based authorization, data secured at the application level and hosted on a HIPAA-compliant platform. 

I wanted to use this project as a launchpad into a Rails job. I haven't had the success I expected from a job search. I thought a robust (for a beginner) project and my background in marketing tech and project management would make me an attractive candidate for an entry level rails job.

I haven't been able to break past the technical interviews, and the interviewers haven't asked any followup questions about it when I mention my relevant experience.

Am I even close to qualified for getting a Rails job, or should I change course and go into something else like Product Management?
## [3][Rails pimped to look modern](https://www.reddit.com/r/rails/comments/g5zk0f/rails_pimped_to_look_modern/)
- url: https://www.reddit.com/r/rails/comments/g5zk0f/rails_pimped_to_look_modern/
---
I am relatively new to Ruby and Rails, but studying hard! Yet I don't want to go down a dead end so thought I better ask some pertinent questions regarding how companies do certain types of apps.

I am using apps like todoist and TickTick. They have nice slick interfaces, with drag and drop etc. Now I am worried that perhaps Rails cannot achieve such loveliness. Or is it a case of me needing some kind of front-end type layer that has many of these functionalities built in?

Could someone enlighten me on this foggy area for me?

Thanks.
## [4][Rails 6 + webpacker + yarn + jquery.fancytree + less](https://www.reddit.com/r/rails/comments/g5xmxw/rails_6_webpacker_yarn_jqueryfancytree_less/)
- url: https://www.reddit.com/r/rails/comments/g5xmxw/rails_6_webpacker_yarn_jqueryfancytree_less/
---
We had to migrate a gem from using fancytree-rails as a ruby gem to a  new rails 6 gem using webpacker and jquery.fancytree coming from npm.  On top of that jquery.fancytree is using LESS (CSS) and you have to do a  few configurations.

App is available at [https://github.com/thebravoman/rails6\_webpacker\_fancytree\_less](https://github.com/thebravoman/rails6_webpacker_fancytree_less)

Here is how to do it it a few simple commands. Hope you find it useful.

[https://kmitov.com/posts/rails-6-webpacker-yarn-fancytree-less/](https://kmitov.com/posts/rails-6-webpacker-yarn-fancytree-less/)
## [5][Try out Cloud 66 for rails developement!!](https://www.reddit.com/r/rails/comments/g5xbhx/try_out_cloud_66_for_rails_developement/)
- url: https://www.reddit.com/r/rails/comments/g5xbhx/try_out_cloud_66_for_rails_developement/
---
Hey everyone, I am part of a new organization which offer cheap and great services for developers to deploy their rails application! Check our product and give me some review. If you are interested just comment below and I will send you a code for 1 month free. Thank you! 

 [https://www.cloud66.com/rails/](https://www.cloud66.com/rails/)
## [6][Cache-control header is getting ignored](https://www.reddit.com/r/rails/comments/g5mrfw/cachecontrol_header_is_getting_ignored/)
- url: https://www.reddit.com/r/rails/comments/g5mrfw/cachecontrol_header_is_getting_ignored/
---
Hey!

I've some problems dealing with the static assets on my site.  I've managed to get the assets served via Cloudfront but I would also like to have them cached but that doesn't seem to work. I've placed config.static\_cache\_control = "public, max-age=31536000" in production.rb but it seems like it's getting ignored 

The response header in chrome looks like this

1. **age:** 41777
2. **content-encoding:** gzip
3. **content-type:** image/svg+xml
4. **date:** Tue, 21 Apr 2020 08:30:39 GMT
5. **last-modified:** Tue, 23 Oct 2018 06:20:26 GMT
6. **server:** Cowboy
7. **status:** 200
8. **strict-transport-security:** max-age=31536000; includeSubDomains
9. **vary:** Accept-Encoding
10. **via:** 1.1 vegur, 1.1 8b82a0c44466382daf259dbb61c8f23c.cloudfront.net (CloudFront)
11. **x-amz-cf-id:** FiTQR0j9SCWMtt5SJzazhyLkW9pg2TB0XltWdygkjQM-ZIkwKQC95w==
12. **x-amz-cf-pop:** ARN1-C1
13. **x-cache:** Hit from cloudfront

I expect a line like "**Cache-control:** Public, max-age=31536000" but it's nowhere to be found.

&amp;#x200B;

My production.rb looks like this as a whole

* config.cache\_classes = true
* config.eager\_load = true
* config.consider\_all\_requests\_local = false
* config.action\_controller.perform\_caching = true
* config.public\_file\_server.enabled = ENV\["RAILS\_SERVE\_STATIC\_FILES"\].present?
* config.assets.js\_compressor = *Uglifier*.new(harmony: true)
* config.assets.compile = false
* config.action\_controller.asset\_host = "1a2b3c4d5e.cloudfront.net"
* config.static\_cache\_control = "public, max-age=31536000"
* config.active\_storage.service = :local
* config.force\_ssl = true
* config.log\_level = :debug
* config.log\_tags = \[:request\_id\]
* config.action\_mailer.perform\_caching = false
* config.i18n.fallbacks = true
* config.active\_support.deprecation = :notify
* config.log\_formatter = ::*Logger*::*Formatter*.new
* if ENV\["RAILS\_LOG\_TO\_STDOUT"\].present?  
logger           = *ActiveSupport*::*Logger*.new(STDOUT)  
logger.formatter = config.log\_formatter  
config.logger = *ActiveSupport*::*TaggedLogging*.new(logger)  
 end
* config.active\_record.dump\_schema\_after\_migration = false

&amp;#x200B;

Can you guys see if there is something I'm missing or if there is something overriding my cache-control?
## [7][Looking for a rails front-end framework](https://www.reddit.com/r/rails/comments/g5ch3n/looking_for_a_rails_frontend_framework/)
- url: https://www.reddit.com/r/rails/comments/g5ch3n/looking_for_a_rails_frontend_framework/
---
I am developing a new app from scratch in rails 6 and I'm looking for a simple frontend framework to:

1.  Allow users to access the app offline and
2.  Access some native features of the phone: like Geolocation, Camera and Bluetooth.*

Indeed, I'm looking for a frontend framework to keep the architecture as simple as possible and avoid having two code-bases.

The app itself is not very complex (think of a "listing app", similar to AirBnb) but given the targeted users, I would like to have strong control over the design. 

I should underline that I am a hobbyist developer (in my last project I used rails 3 --yep, that's a long time ago! ; ) so I would prefer to avoid learning "yet another framework". For that reason, I would like to use html 5 / css as much as possible, in order to keep the development experience as close as possible to a simple web-app.


Question: what frontend framework and / or approach would answer to these requirement?
## [8][Rails database creation issue](https://www.reddit.com/r/rails/comments/g5fm8e/rails_database_creation_issue/)
- url: https://www.reddit.com/r/rails/comments/g5fm8e/rails_database_creation_issue/
---
Hi I am having an issue with the database, whenever I create a database, I get this error below:

&amp;#x200B;

rails db:create

rails aborted!

TypeError: no implicit conversion of Array into Hash

/Users/muhammadkhan/Desktop/DevcampPortfolio/config/environment.rb:5:in \`&lt;main&gt;'

/Users/muhammadkhan/Desktop/DevcampPortfolio/bin/rails:9:in \`&lt;top (required)&gt;'

/Users/muhammadkhan/Desktop/DevcampPortfolio/bin/spring:15:in \`&lt;top (required)&gt;'

bin/rails:3:in \`load'

bin/rails:3:in \`&lt;main&gt;'

Tasks: TOP =&gt; db:create =&gt; db:load\_config =&gt; environment

(See full trace by running task with --trace)

&amp;#x200B;

If someone can be able to help that would be great as I really don't know which file I need to go and fix which function.
## [9][would appreciate hearing your thoughts](https://www.reddit.com/r/rails/comments/g5f59c/would_appreciate_hearing_your_thoughts/)
- url: https://www.reddit.com/r/rails/comments/g5f59c/would_appreciate_hearing_your_thoughts/
---
I have database where i think it not well designed and many problems will show up in future, it like that : a polymorphic relation between "Claim" and "wgclaim" and "rtaclaim" they are a types of this claim so the claim is the father and wgclaim, rtaclaims are sons and soon future we can add more sons.
example : 
rtaclaim : has relation with : police, policy, car, passenger ...
wgclaim: has relation with : plan 
and both share claimable : reference , status, user, team ...
do you think the polymorphic is the right path ? or there another way ??
## [10][Controller is not creating my ActiveRecord models](https://www.reddit.com/r/rails/comments/g5b5i4/controller_is_not_creating_my_activerecord_models/)
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
## [11][update_without_password send re-confirmation email](https://www.reddit.com/r/rails/comments/g4w8sw/update_without_password_send_reconfirmation_email/)
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
