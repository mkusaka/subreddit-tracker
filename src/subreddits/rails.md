# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gnbebg/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gnbebg/personal_projects_show_off_your_own_project_andor/
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
## [2][railsnew.io: the simplest way to generate a Rails app with (or without!) all the bells and whistles](https://www.reddit.com/r/rails/comments/gr7g5d/railsnewio_the_simplest_way_to_generate_a_rails/)
- url: https://www.reddit.com/r/rails/comments/gr7g5d/railsnewio_the_simplest_way_to_generate_a_rails/
---
There’s been a lot of discussion lately about generating new Rails apps. There’s an endless number of tweets lamenting over the default choices.
It’s one of the [hottest topics in ‘May of WTFs’](https://discuss.rubyonrails.org/t/interactive-rails-new/74355).

Even though Rails is more than 15 years old, we are still using the same mechanism to create a new Rails app: `rails new`. And that’s not a problem in and of itself: `rails new` is undoubtedly very powerful and customizable using the template API. But that’s the thing: developers are lazy and do NOT want to customize. This is especially true for __Rails__ developers: convention over configuration is the name of the game!

However… we grew increasingly opinionated about those conventions. DHH’s omakase swiss-army knife grew significantly over the years, and  some (most?) people think it’s more of a kitchen sink now. 

There’s no consensus on what a slimmed-down starter Rails stack should look like, either. Some would go as far as dropping everything and just start with the minimum. Others are __almost__ fine with the omakase stack, except a few things: typically Postgres, RSpec, or perhaps, the Javascript/frontend choices. And there’s everything in-between, centering around the idea of a ‘circa-2009’ stack.

[DHH himself acknowledged the issue](https://discuss.rubyonrails.org/t/interactive-rails-new/74355/50) and gave his blessing to add a —minimal  and an —interactive flag to the official `rails new` generator (as seen on Create React App, Vue CLI, Nuxt.js etc.)

[railsnew.io](https://railsnew.io) is aiming to solve the same problem, using a different approach (for starters, it’s a web application, rather than part of the `rails new` CLI.) [railsnew.io](https://railsnew.io) started out as a weekend fun project. However, with the integration of [railsbytes.com](https://railsbytes.com) and other features added after some initial feedback, we believe it has the potential to become something truly useful.

The app is rough around the edges right now - we are planning to fix things/add more features if it proves to be useful to the community. However, even in its current beta state, it is simple, fast and intuitive to create a new Rails app with everything you (don’t) need. 

Let’s say, you’d like to use Postgres, Stimulus Reflex, and Tailwind, ignoring some things (e.g. spring, various Rails sub-frameworks, sprockets, Turbolinks etc.). With [railsnew.io](https://railsnew.io), this means a few clicks - and it just works!

Once you choose your app’s ingredients and generate the app, you’ll get step-by-step instructions on how to verify it - tailored to that exact stack (provided that you are using any railsbytes, like Stimulus (Reflex) or Tailwind - there’s no use to verify the standard stuff).

I guess that’s enough rambling for now - please [give it a spin](https://railsnew.io) and let us know what do you think!
## [3][Is it OK to catch StandardError in all controller actions, do something, and re-raise the exception?](https://www.reddit.com/r/rails/comments/grgrns/is_it_ok_to_catch_standarderror_in_all_controller/)
- url: https://www.reddit.com/r/rails/comments/grgrns/is_it_ok_to_catch_standarderror_in_all_controller/
---
Is this OK or bad coding practice?

    class ApplicationController &lt; ActionController::Base
    
    around_action :catch_and_rescue
    
    def catch_and_rescue
        yield 
    rescue StandardError =&gt; e
        do something here such as logging and/or send email
        raise e
    end

I know there are gems that do this but I don't want to complicate my application further. Do you think this would cause any problems in the future? Or should I just swallow the hard pill and install a gem like [exception\_notification](https://github.com/smartinez87/exception_notification)? This app won't be used by more than 100 people concurrently so I don't expect to be spammed by this.
## [4][How to avoid N+1 query using SQL views (materialized) in Rails application](https://www.reddit.com/r/rails/comments/gqy90z/how_to_avoid_n1_query_using_sql_views/)
- url: https://www.reddit.com/r/rails/comments/gqy90z/how_to_avoid_n1_query_using_sql_views/
---
In this article, we consider a solution using the SQL view to avoid query problem N+1 when calculating the average values in Ruby on Rails application.  


 Tutorial and link to GitHub is available at: 

[https://jtway.co/how-to-avoid-n-1-query-using-sql-views-materialized-in-rails-application-7cf415cd112f](https://jtway.co/how-to-avoid-n-1-query-using-sql-views-materialized-in-rails-application-7cf415cd112f)
## [5][noob question about duplicating data for less db queries](https://www.reddit.com/r/rails/comments/grde7i/noob_question_about_duplicating_data_for_less_db/)
- url: https://www.reddit.com/r/rails/comments/grde7i/noob_question_about_duplicating_data_for_less_db/
---
hello ! I'm a new ish Rails dev using it on a side project which will hopefully result in a commercial product down the line.

My question, to reduce the number of db queries, does it make sense to duplicate a subset of data in a model / db record which appears in another model ? Then our views only need one db query. If you want to edit things, then you a) get the data from the relation and b) make sure you change the duplicated fields to match the new relation value.

A concrete example, say I have a Ticket model, which can be assigned to a User. Ticket has an int id field which relates to the User ID. If my Ticket model has a field for "user name", then I for all my read operations, I can access just that ticket for everything I need, and I save myself a query for [user.id](https://user.id) .

Am I nuts ?

Thanks \~\~
## [6][SAML SSO Invalid signature in SAML response](https://www.reddit.com/r/rails/comments/gr03b1/saml_sso_invalid_signature_in_saml_response/)
- url: https://www.reddit.com/r/rails/comments/gr03b1/saml_sso_invalid_signature_in_saml_response/
---
I'm using saml-ruby to validate a saml response. And the error message I'm getting is because of this particular line [https://github.com/onelogin/ruby-saml/blob/master/lib/xml\_security.rb#L357](https://github.com/onelogin/ruby-saml/blob/master/lib/xml_security.rb#L357)

I do not understand the signature, public key, and what is being signed. But what I can understand is the certificate in the response x.509's public key is used to verify the signature using the signing algorithm mentioned as part of the response. 

So in my case, the signature does not match and I get an "Invalid SAML signature in the response" error. Could you please help me out which why this is failing, what exactly is being searched for in the signature of the response, and how the IDP generates it.
## [7][How would model cuts of meat?](https://www.reddit.com/r/rails/comments/gr0mok/how_would_model_cuts_of_meat/)
- url: https://www.reddit.com/r/rails/comments/gr0mok/how_would_model_cuts_of_meat/
---
I'm not quite sure how to organize my meat model. For example, a cow is broken into primal cuts, which then have sub-primals, which can be processed into steaks, etc.

What is a good way to structure this? 4 models for animal, primal, sub-primal, and cut?  e.g. cow.primal.subprimal.cut = "delicious"? It feels awkward to access the cuts like this, doesn't it?

&amp;#x200B;

EDIT:

Sorry I didn't respond yesterday. My brain was fried by the time I finished work. 

I'm thinking about two different uses for the model:

1. An inventory for a user. Say I buy a 1/2 cow (which I have), I would like to track what I have left. e.g. `User.inventory` lists all meat in my freezer.
2. A cut explorer so that I can see where a `Piece` is from and identify if there are any exclusion rules. As u/beejamin noted, you can't get a full rib roast &amp; ribeyes.
## [8][C++ HTTP library for POSTing/GETing to Rails 6?](https://www.reddit.com/r/rails/comments/gr7nsd/c_http_library_for_postinggeting_to_rails_6/)
- url: https://www.reddit.com/r/rails/comments/gr7nsd/c_http_library_for_postinggeting_to_rails_6/
---
I need to develop a client in C++14 that talks to a Rails 6 application. Can anybody recommend any libraries? Ideally, I'd like something that works with Devise with me having to manage as little as possible.

The client application will need to run on Windows.

Also, I will need to send small binary attachments to Rails, so it would be nice if the library made that easy.

I see a variety of C++ HTTP libraries out there (cpr, libcurl, cpp-httplib, poco) but I'm not sure if any one is more or less easier to use with Rails.
## [9][Best Way to Implement Multi Column Search Feature](https://www.reddit.com/r/rails/comments/gqtri1/best_way_to_implement_multi_column_search_feature/)
- url: https://www.reddit.com/r/rails/comments/gqtri1/best_way_to_implement_multi_column_search_feature/
---
So I wanted to add search feature which would look for LIKE results in multiple columns of the table, ie author (via relation), description, comment et cetc. Can anyone recommend me best approach, gem for this? :) I saw multiple gems but most of them seem too bulky for my simple needs.
## [10][Custom Validator](https://www.reddit.com/r/rails/comments/gr4vz9/custom_validator/)
- url: https://www.reddit.com/r/rails/comments/gr4vz9/custom_validator/
---
Does anyone have a Rails 6 resource for making a custom validator that would accept params in order#new controller. I’ve built an order lifecycle that is customized for a small business’ web store.  At the moment of payment, instantiating the order from a cart, I need to validate not only the Order model but also many parts of the Contact model (which are never required aside from this moment of declaring shipping/billing info).  To account for guest checkout, I don’t wish put these validations in an existing model; I want the validation to be freestanding and perform only in this checkout route.


I completed a bootcamp several years ago, but after Dev Boot Camp closed, it took down its github; and nearly all of the code I wrote was in clones of project prompts from the DBC github that were removed.  I was financially sunk after camp and had to resort to immediate employment instead of rebuilding my portfolio from nothing.  So!  I’m grasping for current resources, and would consider hiring a tutor if anyone if interested!
## [11][Are redis connections and action cable connections the same thing?](https://www.reddit.com/r/rails/comments/gr3s8b/are_redis_connections_and_action_cable/)
- url: https://www.reddit.com/r/rails/comments/gr3s8b/are_redis_connections_and_action_cable/
---
Sorry if this is a dumb question, but I've never set up action cable before.  I have it all working in dev, and on my staging server, but I want to make sure that once this goes into production, I'm ready.
