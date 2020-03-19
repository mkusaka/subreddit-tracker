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
## [2][Rails Logging Help](https://www.reddit.com/r/rails/comments/fl31hn/rails_logging_help/)
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
## [3][Social Media Platform](https://www.reddit.com/r/rails/comments/fl9vzd/social_media_platform/)
- url: https://www.reddit.com/r/rails/comments/fl9vzd/social_media_platform/
---
Hey guys,   
I'm working on a project using the Rails WAF. I'm relatively new to this and am getting lost here and there so please be patient.  


I'm looking to create a social media platform for a college project where users can sign up and organise events with one another.  


So far I have followed a few tutorials online and have been able to successfully implement a functioning devise based platform, which allows users to sign-up/sign-in and create 'events' which can be seen by other users and have basic CRUD functionality for the user who created the event.  


I am now looking to turn my attention to user profiles. From what I have gathered from my research it is okay to used the same database table to populate a user profile that I have used to create user login.  


Can you please advise the best way to create user profiles which allows the logged in user to flesh out their account with personal details which can be viewed by other users?
## [4][Improving performance with Presenter and Eager Load of a big Rails app](https://www.reddit.com/r/rails/comments/fkvtdz/improving_performance_with_presenter_and_eager/)
- url: https://www.reddit.com/r/rails/comments/fkvtdz/improving_performance_with_presenter_and_eager/
---
Hey guys, I have written a post about how I helped to improve the performance of a Ruby on Rails application where I work.

The app is used by people all over the world and has around of 4 million request every month.

https://medium.com/@wlss.lucas/improving-performance-of-a-rails-app-12beb830319a
## [5][Best pattern for ActiveRecord Create objects](https://www.reddit.com/r/rails/comments/fl2103/best_pattern_for_activerecord_create_objects/)
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
## [6][[HELP] Restrict browsers mode on websites](https://www.reddit.com/r/rails/comments/fkw3kv/help_restrict_browsers_mode_on_websites/)
- url: https://www.reddit.com/r/rails/comments/fkw3kv/help_restrict_browsers_mode_on_websites/
---
Hi guys, this question is not related to Rails, but with websites in general. I work with a website that the browsers change to a restrict mode after you log in on the restrict area . What I mean with "restrict mode" is a browser with static address bar (users can't change), no button options (bookmark, reload, back/forward page, etc.). You basically have a static address bar and the rendered html page. 

My question is: why they use this restrict mode? I think is because of the back/forward page and the static address bar and how the website caches, or how they load Javascript on page, or because they don't want users reloading pages as the website load a lot of data on certain pages, or because they don't want users open more than one tab per user logged in, I don't know, but I was kind of curious.

Sorry if this question don't fit this sub, but I didn't know where to ask.
## [7][Taking Rails 6’s Action Mailbox for a Spin](https://www.reddit.com/r/rails/comments/fkmhsu/taking_rails_6s_action_mailbox_for_a_spin/)
- url: https://www.reddit.com/r/rails/comments/fkmhsu/taking_rails_6s_action_mailbox_for_a_spin/
---
In this article the author will take Action Mailbox for a spin and write a simple working Rails 6 app to demonstrate Action Mailbox in action.

Link to article: https://medium.com/merajulislam/taking-rails-6s-action-mailbox-for-a-spin-104d0e34d379
## [8][Google Oauth2 not applying image_size](https://www.reddit.com/r/rails/comments/fkombb/google_oauth2_not_applying_image_size/)
- url: https://www.reddit.com/r/rails/comments/fkombb/google_oauth2_not_applying_image_size/
---
I'm trying to apply image\_size to omniauth-google-oauth2, but everything I tried didn't work.

My devise.rb

    config.omniauth :google_oauth2, google_client_id, google_client_secret, {image_aspect_ratio: "square", image_size: 100}

I'm using my application\_helper to display the image with:

    module ApplicationHelper def avatar_url(user)     user.image   end end

If I check the image\_url it still shows [https://lh3.googleusercontent.com/a-/AOh14Gg9KNrUy\_97f3wDhevAJGE52nT8Y6ho48Uo9vJwHw](https://lh3.googleusercontent.com/a-/AOh14Gg9KNrUy_97f3wDhevAJGE52nT8Y6ho48Uo9vJwHw) without any change in size

I've also tried auth.info.image.sub('?sz=32','?sz=100'), but it didn't work.
## [9][ActionText uploading a gif is being shown as an image](https://www.reddit.com/r/rails/comments/fk8anx/actiontext_uploading_a_gif_is_being_shown_as_an/)
- url: https://www.reddit.com/r/rails/comments/fk8anx/actiontext_uploading_a_gif_is_being_shown_as_an/
---
Trying to upload a gif through action text and it gets uploaded correctly but on update or creation something happens and an image variant is created and is used to display the post  

would appreiciate any help, here is the SO link with a bit more detail

[https://stackoverflow.com/questions/60721539/actiontext-gif-attachment-converts-to-photo-when-submitting-rails-form](https://stackoverflow.com/questions/60721539/actiontext-gif-attachment-converts-to-photo-when-submitting-rails-form)
## [10][Example of Search Form with Google Places Autocomplete in Rails 6?](https://www.reddit.com/r/rails/comments/fkch08/example_of_search_form_with_google_places/)
- url: https://www.reddit.com/r/rails/comments/fkch08/example_of_search_form_with_google_places/
---
Does anyone have a tutorial of an example of Google Places Autocomplete being used in a search form in Rails 6? I am trying to add it to my app, but it looks like most of the resources available are for Rails 3/4.
## [11][Aliasing polymorphic relationships](https://www.reddit.com/r/rails/comments/fkdccv/aliasing_polymorphic_relationships/)
- url: https://www.reddit.com/r/rails/comments/fkdccv/aliasing_polymorphic_relationships/
---
Hi, I have an address model. It has an belongs\_to :addressable relationship which is polymorphic. One of the models that can be an addressable is "property".

I'm aware there's a syntax that allows me to somewhat alias the addressable relationship with an specific relation.

Somewhat like

    belongs_to :addressable, polymorphic: true
    belongs_to :property, as: :addressable, -&gt; {where(addressable_type: 'Property')}

&amp;#x200B;

What's the correct syntax for it? I have done this before but I can't find anything online or figure out where I've done it.
