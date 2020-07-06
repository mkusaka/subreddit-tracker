# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/hja8kx/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/hja8kx/personal_projects_show_off_your_own_project_andor/
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
## [2][Rails is about giving the power of a 10 person development team to one person through simplicity](https://www.reddit.com/r/rails/comments/hm7mex/rails_is_about_giving_the_power_of_a_10_person/)
- url: https://www.reddit.com/r/rails/comments/hm7mex/rails_is_about_giving_the_power_of_a_10_person/
---
Having built SPA using front end frameworks before, and now building the same web app with Rails I was thinking about the following:

Rails is a great tool for solo / small team developers because it gives you the leverage to deal with the same set of problems as a larger team by providing a simplistic mental model through which to reason about the software application.

For example, let's discuss the often heated SPA / JSON Api versus Server Side Application debate. I think this debate is not actually a technology problem - but a personnel problem. For example, suppose you  have an application that is highly interactive. Giving a dedicated front end team (3-5 of 10 people) a front end framework makes complete sense because it allows those individuals to organize and handle the complexity of five parallel change sets at any one time. That is to say, the front end framework provides the opportunity to create a "\*communicable\* model" for this squad to maintain in their collective understanding \*and to reason about / collaborate on\*.  Without a \*communicable\* model (e.g. spaghetti javascript / css), the difficulty of proposing new change sets and evolving software in a multiplayer environment grows in a non-linear way. This is a really misrepresented risk in technology choices I think.

Frameworks linearize the learning curve for the application's mechanics, allowing \*new collaborators\* to ramp on and quickly evolve their end of the software.

Flipping the script now, if you are just one or two developers, it might not make sense to use a full blown front end framework. The reason is that the primary issue you're dealing with is not managing a \*communicable\* model between many contributors to handle parallel change sets, rather it is a question of "how much you can get done". Rails provides the \*leverage\* here by abstracting certain application decisions / choices and doing the "heavy lifting" allowing you, a solo developer, to hold the entire application model in \*your head\*. That is to say, the application model might not be a \*communicable\* model (e.g. Sever Generated JavaScript responses is probably not as easily grasped as parsing a JSON request) but it is simple enough to hold in one person's mind and still evolve the application with a sufficient speed.

There's some nuances here and I'm probably wrong in some regards, but an interesting thought (I think) none the less

&amp;#x200B;

edit; spelling
## [3][How to build guest cart/checkout?](https://www.reddit.com/r/rails/comments/hm6xjd/how_to_build_guest_cartcheckout/)
- url: https://www.reddit.com/r/rails/comments/hm6xjd/how_to_build_guest_cartcheckout/
---
Currently I have an ecommerce app that allows users to add items to a cart and checkout if they have an existing user account (the cart object is associated with user object and product object). I want to create the ability for guest checkout, but not sure how this would work from an architecture perspective if a user doesn't have an account. Any thoughts? Thanks in advance.
## [4][Spree Commerce Help](https://www.reddit.com/r/rails/comments/hm6hfg/spree_commerce_help/)
- url: https://www.reddit.com/r/rails/comments/hm6hfg/spree_commerce_help/
---
Hi There,

Trying out Spree Commerce and having trouble rendering editted templates in my app/views/spree directory. I am able to override templates in Solidus by recreating them in app/views/spree but it doesn't seem to work in Spree.

Anyone experienced this issue?

Rails 6.0.3 &amp; Spree 4.1

&amp;#x200B;

Thank you in advance!
## [5][Unable to disable asset fingerprinting for emails only?](https://www.reddit.com/r/rails/comments/hlymis/unable_to_disable_asset_fingerprinting_for_emails/)
- url: https://www.reddit.com/r/rails/comments/hlymis/unable_to_disable_asset_fingerprinting_for_emails/
---
tl:dr; Images in emails work when fingerprinted, but not for static images in the public folder

Problem:

I'm embedding images in my ActionMailer / Devise sent emails, like logos etc.

This works when I use images in my assets folder, but are fingerprinted.

I don't want them to be fingerprinted or to use the fingerprinted assets that I use do use elsewhere in the application, so that these images don't break in future.

Putting images in the public folder works locally, but not on higher environments.

Asset folder working with fingerprinting:

    &lt;%= image_url('emails/header-Logo.png') %&gt; 
    &lt;%= asset_url('emails/header-Logo.png', :skip_pipeline =&gt; false) %&gt; 
    # Generates fingerprint local + remote src=http://cdn.local.me:3000/assets/emails/header-Logo-57b82ca2de25bbafe57e4b21ed884d8c972344fca646ee08873fdbbfd9db8259.png 

Public folder working (locally only), with no fingerprinting:

    &lt;%= asset_url('images/emails/Email-Welcome-Features.png', :skip_pipeline =&gt; true) %&gt; 
    # Generates locally src="http://cdn.local.me:3000/images/emails/Email-Welcome-Features.png" 
    # Generates remote src="http:///images/emails/Email-Welcome-Features.png"

Prod config:

    if ENV['CDN_HOSTNAME']   
    config.action_controller.asset_host = ENV['CDN_HOSTNAME']   config.action_mailer.asset_host = ENV['CDN_HOSTNAME'] 
    end

Other assets, like on the view templates, js content etc, do come correctly (fingerprinted) from the cdn in prod..

Rails version: rails ([6.0.2.2](https://6.0.2.2))

Any help appreciated :)
## [6][Insight into cause of "Completed 401 Unauthorized" error after being redirected to app from an external service](https://www.reddit.com/r/rails/comments/hlz7bn/insight_into_cause_of_completed_401_unauthorized/)
- url: https://www.reddit.com/r/rails/comments/hlz7bn/insight_into_cause_of_completed_401_unauthorized/
---
Users are directed to Stripe Connect as part of the onboarding process for my app. When they finish with Stripe, they are sent back to my site with two params in the url: `code` and an encoded `state`. In this case, the latter param was previously saved to the db in decoded form, then encoded and passed to Stripe to prevent CSRF forgery.

My problem is that users are signed out when redirected back to my app upon the first pass through onboarding, whereas the second pass goes through just fine. In other words, if they complete the sign in and do everything over again, they are allowed to proceed. A `Completed 401 Unauthorized` in the console is all I have to go off of. Presumably this has something to do with Devise and it's `authenticate_user!` method?

    2020-07-06T00:29:24.121304+00:00 app[web.1]: Started GET "/onboard/id_verification?code=ac_Tau1R4hyn47LGrwiisUpFDEEMmpWVBrw&amp;state=YTVmNGZmOTUtMjc3ZC00YjIyLThjZDctNTliODBlMzY1OThk" for ###.##.##.## at 2020-07-06 00:29:24 +0000
    2020-07-06T00:29:24.123838+00:00 app[web.1]: Processing by Users::OnboardsController#id_verification as HTML
    2020-07-06T00:29:24.123891+00:00 app[web.1]: Parameters: {"code"=&gt;"ac_Tau1R4hyn47LGrwiisUpFDEEMmpWVBrw", "state"=&gt;"YTVmNGZmOTUtMjc3ZC00YjIyLThjZDctNTliODBlMzY1OThk"}
    2020-07-06T00:29:24.127212+00:00 app[web.1]: Completed 401 Unauthorized in 3ms (ActiveRecord: 0.0ms)
    2020-07-06T00:29:24.261889+00:00 app[web.1]: Started GET "/users/sign_in" for ###.##.##.## at 2020-07-06 00:29:24 +0000
    2020-07-06T00:29:24.265479+00:00 app[web.1]: Processing by Users::SessionsController#new as HTML

Assuming a happy path (decoded `state` param equals what is saved to the db, `code` is `present?`, and `stripe_user_id` is either `present?` or `nil`) the user is taken to a form for uploading some documentation. The stuff they upload is encrypted and saved to a private AWS bucket with IAM policies appropriately set, but the way that works is outside the scope of this question. Here's the relevant controller action:

    class Users::OnboardsController &lt; ApplicationController
    
      def id_verification
        if params[:state].blank? || params[:code].blank?
          redirect_to onboard_stripe_path
        elsif Base64.decode64(params[:state]) != current_user.stripe_state_token
          flash[:error] = "CSRF forgery detected"
          redirect_to onboard_stripe_path
        elsif Base64.decode64(params[:state]) == current_user.stripe_state_token &amp;&amp; params[:code].present? &amp;&amp; current_user.stripe_user_id.present?
          @user = current_user
          respond_to do |format| 
            format.html { render '/users/onboard/id_verification.html.erb' } 
          end
        elsif Base64.decode64(params[:state]) == current_user.stripe_state_token &amp;&amp; params[:code].present? &amp;&amp; current_user.stripe_user_id == nil
          get_stripe_id = StripeUser.connect(params[:code])
          current_user.stripe_user_id = get_stripe_id.stripe_user_id
          current_user.save
          @user = current_user
    
          respond_to do |format|
            format.html { render '/users/onboard/id_verification.html.erb' } 
          end
        end
      end
    end

Does anyone have any insight as to the source of that `Completed 401 Unauthorized` error, and why it is only an issue on the first pass through onboarding? Thank you!
## [7][Turbolinks and preserving focus state](https://www.reddit.com/r/rails/comments/hlmmhy/turbolinks_and_preserving_focus_state/)
- url: https://www.reddit.com/r/rails/comments/hlmmhy/turbolinks_and_preserving_focus_state/
---
Howdy.

Is there a way of preserving focus state? My use case is that I want the "new" form to autosave with a debounce, and once it saves it should become the "edit" form of the saved entry. I can manually replace the history state but I think that it would be cool to get it for free from Turbolinks, almost like Inertia.

Edit: \`data-turbolinks-permanent\` didn't work for me. Not sure if that's only an issue I'm experimenting because I'm writing buggy code. I use it somewhere else in the app and it works well, but it does not preserve the focus state when I'm inside an element, as if they still removing the element from the DOM.
## [8][Best and cheapest hosting options?](https://www.reddit.com/r/rails/comments/hlgze5/best_and_cheapest_hosting_options/)
- url: https://www.reddit.com/r/rails/comments/hlgze5/best_and_cheapest_hosting_options/
---
What would you recommend for a good, Rails-tailored hosting plan (like Heroku) for around $7-8 per month? That is what my client told me she would like to pay. Heroku Hobby is $7, but is that good enough for a production e-commerce site?

If not, what else would you recommend? I'd like something like Heroku that is specifically tailored to Rails where you can push your code from GitHub and it works. I tried and failed to deploy my app to Digital Ocean's Rails droplet (and I'm far from being a noob when it comes to the Linux CLI).
## [9][Devs who dislike or has been bitten by Active Record's default_scope, how about we include it in style guides?](https://www.reddit.com/r/rails/comments/hlakkq/devs_who_dislike_or_has_been_bitten_by_active/)
- url: https://www.reddit.com/r/rails/comments/hlakkq/devs_who_dislike_or_has_been_bitten_by_active/
---
[https://github.com/rubocop-hq/rails-style-guide/issues/266](https://github.com/rubocop-hq/rails-style-guide/issues/266)
## [10][Trying to understand If-Modified-Since HTTP Header](https://www.reddit.com/r/rails/comments/hlgzev/trying_to_understand_ifmodifiedsince_http_header/)
- url: https://www.reddit.com/r/rails/comments/hlgzev/trying_to_understand_ifmodifiedsince_http_header/
---
I'm building a Rails app that consumes an API that accepts this header.

&gt; "The server then determines based on these headers if the client's cached copy is the most recent version of this file. If so, the server returns an HTTP status 304 code, letting the client know it can reuse its cached copy."

But I don't know if I understand the idea, does this mean I have to always save the result of each API request to a file or database? and if the request's result has an HTTP status 304 should I rescue it from the database? What if I don't have the result in the database? Do I send you a request without that header to get the result again?
## [11][New bootstrap toasts, pretty neat](https://www.reddit.com/r/rails/comments/hli2tq/new_bootstrap_toasts_pretty_neat/)
- url: https://www.reddit.com/r/rails/comments/hli2tq/new_bootstrap_toasts_pretty_neat/
---
&amp;#x200B;

[toasts](https://i.redd.it/q4zhf2oxez851.gif)

&amp;#x200B;

My application.html.erb  


    &lt;!DOCTYPE html&gt;
    &lt;html&gt;
      &lt;head&gt;
        ...
        &lt;%= javascript_pack_tag 'flash', 'data-turbolinks-track': 'reload' %&gt;
      &lt;/head&gt;
    
      &lt;body&gt;
        ...
        &lt;%= render '/flashes' %&gt;
        &lt;%= render '/toasts' %&gt;
        &lt;%= yield %&gt;
      &lt;/body&gt;
    &lt;/html&gt;


 app/views/\_flashes.html.erb   


    &lt;div class="acme-flashes notice d-none"&gt;
      &lt;%= flash[:notice] %&gt;
    &lt;/div&gt;
    &lt;div class="acme-flashes alert d-none"&gt;
      &lt;%= flash[:alert] %&gt;
    &lt;/div&gt;

 

app/views/\_toasts.html.erb 

    &lt;div class="acme-toasts position-absolute w-100 p-4 d-flex flex-column align-items-end" style="z-index: 1"&gt;
      &lt;div class="w-25 inner"&gt;
        &lt;!-- empty --&gt;
      &lt;/div&gt;
    &lt;/div&gt;

&amp;#x200B;

app/javascript/packs/flash.js

    require("acme/flash").start()

app/javascript/acme/flash.js

    export function start() {
      document.addEventListener("turbolinks:load", () =&gt; {
        create()
        show()
      })
    }
    
    function show() {
      $('.toast')
        .toast({ delay: 2000 })
        .toast('show')
    }
    
    function create() {
      $.each( getFlashesInDom(), function(index, value) {
        $(".acme-toasts .inner").append(
          template(value)
        )
      })
    }
    
    function template(body) {
      return `
        &lt;div class="toast" role="status" aria-live="polite" aria-atomic="true" zdata-delay=3000 style="z-index: 1"&gt;
          &lt;div class="toast-header"&gt;
            &lt;!-- an icon can be here --&gt;
            &lt;strong class="mr-auto"&gt;
              &lt;!-- title here --&gt;
            &lt;/strong&gt;
            &lt;small class="text-muted"&gt;just now&lt;/small&gt;
            &lt;button type="button" class="ml-2 mb-1 close" data-dismiss="toast" aria-label="Close"&gt;
              &lt;span aria-hidden="true"&gt;&amp;times;&lt;/span&gt;
            &lt;/button&gt;
          &lt;/div&gt;
          &lt;div class="toast-body"&gt;
            ${body}
          &lt;/div&gt;
        &lt;/div&gt;`
    }
    
    function getFlashesInDom() {
      return $.map( $(".acme-flashes"), function(val,i) {
        let html = $.trim( $(val).html() )
        if ( html.length == 0 ) { return }
        return html
      })
    }

config/webpack/environment.js  


    const { environment } = require('@rails/webpacker')
    
    // https://github.com/webpack-contrib/expose-loader/blob/master/README.md
    environment.loaders.append('expose', {
            test: require.resolve('jquery'),
            loader: 'expose-loader',
            options: {
              // For `underscore` library, it can be `_.map map` or `_.map|map`
              exposes: ['jquery', '$', 'jQuery']
            }
    });
    
    module.exports = environment

My yarn.lock has  bootstrap@\^4.5.0  
btw.
