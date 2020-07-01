# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/har6r7/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/har6r7/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/hja8kx/personal_projects_show_off_your_own_project_andor/)
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
## [3][Any gems that will create a seeds script from my current DB?](https://www.reddit.com/r/rails/comments/hj783p/any_gems_that_will_create_a_seeds_script_from_my/)
- url: https://www.reddit.com/r/rails/comments/hj783p/any_gems_that_will_create_a_seeds_script_from_my/
---
Looking for a gem that will look at my current DB and create a script so I can drop it and recreate it if required.

My goal is to push my local db up to Heroku but you can’t do this on free plans.
## [4][Gem for passwordless Auth in Rails API](https://www.reddit.com/r/rails/comments/hj8zf6/gem_for_passwordless_auth_in_rails_api/)
- url: https://www.reddit.com/r/rails/comments/hj8zf6/gem_for_passwordless_auth_in_rails_api/
---
I am on the lookout for a gem that enables passwordless auth (ie you only have to enter an email to be sent a magic link to login) in a rails API. Having looked this morning I’ve found examples that work when using rails for views but not any that are setup to work with only rails APIs. Just wondered if anyone knew of any?
## [5][what best way to Encrypt &amp; Decrypt data using ruby(rails)](https://www.reddit.com/r/rails/comments/hj7jre/what_best_way_to_encrypt_decrypt_data_using/)
- url: https://www.reddit.com/r/rails/comments/hj7jre/what_best_way_to_encrypt_decrypt_data_using/
---
i take from user token and i think must encrypt this token before save in database . and need decrypt token for use get some data .

i generate private key and save it in db &amp;  public key save in csv in local server .  


i need suggest way or gem to help  me .
## [6][Just a thank you](https://www.reddit.com/r/rails/comments/hipz4t/just_a_thank_you/)
- url: https://www.reddit.com/r/rails/comments/hipz4t/just_a_thank_you/
---
One of the main reasons why I dived into Rails was because I heard about how good the community is - and boy did I experience that. Here and the Slack group.

I dived into this with pretty much zero experience in Rails about two months ago and a month ago I deployed a production ready API using ECS. 

I have this wonderful community here to guide me along the way and clear any pesky little doubts I had. 

So, thank y’all!
## [7][Apps with multiple webpacker packs, how do you organize your app/javascript folder?](https://www.reddit.com/r/rails/comments/hixnme/apps_with_multiple_webpacker_packs_how_do_you/)
- url: https://www.reddit.com/r/rails/comments/hixnme/apps_with_multiple_webpacker_packs_how_do_you/
---
We have a large app with multiple `packs` entry points. Then each pack gets its own folder under `app/javascript`. There's one `shared` folder where we obviously put shared stuff. It's not bad, but I'm wondering how other people are solving this.

Are you doing something different?
## [8][Confusion about Rails API Security](https://www.reddit.com/r/rails/comments/hj0isx/confusion_about_rails_api_security/)
- url: https://www.reddit.com/r/rails/comments/hj0isx/confusion_about_rails_api_security/
---
So far I've set up my rails grape API to use JWT for authentication using the [**grape-jwt-authentication**](https://github.com/hausgold/grape-jwt-authentication) gem.

I am confused about whether the data payload being exchanged between my server and client's server should also be encrypted using the RSA key pairs? Or should there be a separate encryption key between the server for this?

Or does HTTPS already take care of encrypting the data in transition?
## [9][Shopify App + Devise. Company website registration/login for stores?](https://www.reddit.com/r/rails/comments/hiwn21/shopify_app_devise_company_website/)
- url: https://www.reddit.com/r/rails/comments/hiwn21/shopify_app_devise_company_website/
---
I have built a few different Shopify Applications using the shopify\_app gem (Shopify Application Rails engine and generator). This is great for redirecting a user from an app developer's website to the shopify store to download the application.

However, how would one setup a registration/login at the app developer's website (using devise) to then be able to connect to the store owners shopify app/install it in the shopify store?

Note: using the shopify\_app gem a route is created that directs the user to a shopify login page, like so:

    Rails.application.routes.draw do
      root :to =&gt; 'home#index'
      mount ShopifyApp::Engine, at: '/'
      # For details on the DSL available within this file, see https://guides.rubyonrails.org/routing.html
    end

Also I am wondering how one would allow a shop that has installed your application to then be able to go to your company website and register/login? I was thinking about setting a new account for each app install (store name would be the subdomain - so it is reserved), and allowing the user to sign up by using the forget password route? Would this be ok?
## [10][How do you handle form spam? What about scoring/grading/filtering of content/leads/etc?](https://www.reddit.com/r/rails/comments/hinhsm/how_do_you_handle_form_spam_what_about/)
- url: https://www.reddit.com/r/rails/comments/hinhsm/how_do_you_handle_form_spam_what_about/
---
Hey group!  


I am almost done writing the first version of a gem that's meant to easily and intelligently handle form spam. It's a process that I had already incorporated into a current project, and I've gone ahead and extracted it and abstracted everything into a more flexible set of classes and methods.

But the reason for this post is... well I'm experiencing a little imposter syndrome. I was all fired up about writing the gem and getting it out into the world, but now I'm just wondering if it would even be useful to people.  


So, can you the community help me out? I'd really like to know what you do to handle spammy form submissions from public forms. We all know, those dam bots are AWFUL. But I felt like there were so many different ways that I saw people handling it.  


For what it's worth, here's the breakdown of my gem's initial core features:  


1. A rule based system that allows for detecting matching keywords, links, WP formatted comments. emails and domains, and scoring the object being analyzed. (90% Done)
   1. Allows for positive &amp; negative point assignments to a matching keyword, email, or domain, which use either the default point values from the initializer or the point value you assign in the rule.
   2. Has a 'Ban' list to automatically 'fail' any input that matches anything you absolutely want blocked from the site. However, it does not block anything, it just calculates a score that would be considered 'failing' automatically.
   3. Has a 'Pass' list to automatically bypass the rest of the scoring process and assign the lowest/highest passing score possible.
   4. I also have a handful of public methods that allow you to run your own detections for counting the number of links in a string (of any size), that allow you to detect keywords, emails, or domains and return a True/False value if the value is found.
   5. I have a couple of public methods that also have a customizable sensitivity level, so if you've having issues where some submissions are getting crafty with how they are submitting content you want scored, then you can pass in an additional param to set the sensitivity to 'High' and get a slightly more broad match. Useful for spammers who write URLs like 'http blahblahblah com'
   6. The scoring can also be used in a positive way, like scoring inbound leads to set them as a priority based on their email domain, or keywords in their submission details.

&amp;#x200B;

Features I plan to extract from my current app into the gem:

1. Adding a 'honeypot' feature to form\_with and form\_for methods.
   1. The honeypot would detect the fields that the programmer created, and then inject commonly abused/used fields that are not in the original field list as hidden input fields.
   2. During scoring, honeypot fields would be checked for any value being present (because they should be blank, they are hidden fields after all). Anything being detected means a bot filled out the form and the submission automatically fails scoring. The IP address gets logged for the abusive behavior.
2. Add the IP address logging for submissions that have failed the scoring for spam. Other scoring models would/could be excluded since you can also use the scoring for positive methods as well that are entirely unrelated to bad behavior.
3. Integrating with Rack to intercept bad actors and BAN THOSE MOFOS. Seriously, I'm bringing down the ban-hammer on these bots. I have no sympathy for abusers.
   1. Over time, the log of bad actors submitting can keep track of abusive/unwanted behavior and allow you to put restrictions on actions taken.
   2. Outright ban IP addresses who have bad behavior patterns. I'd like to setup some defaults, but allow this to be configurable and flexible for however someone decides to use this in their app. It's really meant to work at the Rack level so that it intercepts and stops processing an entire request through the Rails stack and saves on processing. It's very useful when getting spam bombed.

&amp;#x200B;

Anyway, that's my idea that I've been slowly working on in the dark. Would this even be useful to people out there? I know it's useful to me, so I'm going to do it regardless, but it would be nice to know if I'm actually solving a problem for others.
## [11][Anyone figure out a way to add structure to actiontext columns? ActionList?](https://www.reddit.com/r/rails/comments/hit4l8/anyone_figure_out_a_way_to_add_structure_to/)
- url: https://www.reddit.com/r/rails/comments/hit4l8/anyone_figure_out_a_way_to_add_structure_to/
---
I would like to add a "list" to a product table without adding a new JSON Column. 

would love `has_rich_list :features` with a form input that is strictly list elements. 

I have not yet seen any custom "actiontext" gems but i imagine this coming as people get more familiar with the data models. 

yolo
## [12][Best way to add a search/filter/sort to a list of records?](https://www.reddit.com/r/rails/comments/hinagw/best_way_to_add_a_searchfiltersort_to_a_list_of/)
- url: https://www.reddit.com/r/rails/comments/hinagw/best_way_to_add_a_searchfiltersort_to_a_list_of/
---
I'm looking for a package or gem that will let me add this functionality rather than having to code it from scratch.  What's the best solution for this?

Ideally, it would be nice to have it async so I don't have to refresh the page but that's not too important.
