# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/ep2dw9/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/ep2dw9/personal_projects_show_off_your_own_project_andor/
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
## [2][AMP + IMG + RAILS](https://www.reddit.com/r/rails/comments/esqb14/amp_img_rails/)
- url: https://www.reddit.com/r/rails/comments/esqb14/amp_img_rails/
---
Error about width and height (random?).

I created the AMP pages on my website. To add the &lt;amp-img&gt; we usted this system:

in helpers/application\_helper.rb there is this:

      def amp_image_tag(url, options = {})
        options[:width], options[:height] = *safely_extract_image_size(url) unless options[:width]
        content_tag('amp-img', nil, options.merge(src: asset_url(url)))
      end
    
      def safely_extract_image_size(url)
        FastImage.size(asset_url(url))
      rescue StandardError =&gt; e
        Rails.env.development? ? raise(e) : Rails.logger.warn(e)
        [640, 480]
      end

and in the view page, we just made this:

      &lt;% @movie.screens.each do |screen| %&gt;
      &lt;div&gt;
        &lt;%= amp_image_tag(screen.file.url(:medium), alt: @movie.title, layout: 'responsive') %&gt;
      &lt;/div&gt;
    &lt;% end %&gt;

In amp-img tag, the width and height values are obligatory.

I don't undestad why but sometime those values are showed and sometime they miss.

What is wrong?

p.s. the images have different sizes. They are uploaded by users.
## [3][Is there a boilerplate repo that has authentication configured?](https://www.reddit.com/r/rails/comments/esgmnf/is_there_a_boilerplate_repo_that_has/)
- url: https://www.reddit.com/r/rails/comments/esgmnf/is_there_a_boilerplate_repo_that_has/
---
I'm building an app from scratch using Rails and React. Is there any boilerplate project out there that has a simple log in page and Devise (or other authentication gem) configured that I can work off of? The closest thing I can find is thoughtbot's https://github.com/thoughtbot/clearance
## [4][asdf version manager and bundle exec question](https://www.reddit.com/r/rails/comments/esi2s7/asdf_version_manager_and_bundle_exec_question/)
- url: https://www.reddit.com/r/rails/comments/esi2s7/asdf_version_manager_and_bundle_exec_question/
---
I'm hoping some asdf users here can shed some light on this for me:

I recently set up a new machine and used thoughtbot's laptop script to speed up the process. one of the things it setups for it asdf as your version manager. i come from using rvm, but i was interested in giving asdf a try anyway.

a new coworker of mine also set her machine up similarly and is also using asdf as her version manager.

i'm having to prepend bundle exec before executing gem commands like rspec, sidekiq or rails, but NOT before running rake tasks

my coworker does NOT have to prepend bundle exec for anything. we checked her zshrc file and she doesn't have any bundle exec aliases either.

Why am i having to prepend bundle exec and she doesn't? Why is it specific to gems so far and not rake? Is it typical with asdf to have to prepend it without an alias?
## [5][Reading ActiveRecord Callbacks guide, is there any difference between using a method vs a proc in these code examples?](https://www.reddit.com/r/rails/comments/escllx/reading_activerecord_callbacks_guide_is_there_any/)
- url: https://www.reddit.com/r/rails/comments/escllx/reading_activerecord_callbacks_guide_is_there_any/
---
I am reading the below in the callbacks guide, and am wondering, is there a difference/benefit between just using a private class method `Order#paid_with_card?` versus the second option of having a Proc used `Proc.new { paid_with_card? }`? 

-------

https://guides.rubyonrails.org/active_record_callbacks.html#using-if-and-unless-with-a-symbol

&gt;You can associate the :if and :unless options with a symbol corresponding to the name of a predicate method that will get called right before the callback. When using the :if option, the callback won't be executed if the predicate method returns false; when using the :unless option, the callback won't be executed if the predicate method returns true. This is the most common option. Using this form of registration it is also possible to register several different predicates that should be called to check if the callback should be executed.

    class Order &lt; ApplicationRecord
      before_save :normalize_card_number, if: :paid_with_card?
    end

https://guides.rubyonrails.org/active_record_callbacks.html#using-if-and-unless-with-a-proc

&gt;It is possible to associate :if and :unless with a Proc object. This option is best suited when writing short validation methods, usually one-liners:

    class Order &lt; ApplicationRecord
      before_save :normalize_card_number, if: Proc.new { paid_with_card? }
    end

-----

I think I need to review my Proc understanding, basics of it is just storing a block as a variable that can be called later with `.call` of course, according to the doc there is more nuances involved. Referred to this:

https://ruby-doc.org/core-2.6/Proc.html
## [6][On estimating, my response to a client who indicated his budget of $2k (#freelancing #rails )](https://www.reddit.com/r/rails/comments/esdtx4/on_estimating_my_response_to_a_client_who/)
- url: https://www.reddit.com/r/rails/comments/esdtx4/on_estimating_my_response_to_a_client_who/
---
Recently I had the lawn re-done for my front yard. And the first thing I did was ask for an estimate.Am I a hypocrite for shying away from estimates for my own clients?

I'm not sure. There's two kinds of work.While collar work - mostly mental and very little repetition - i do that.Blue collar work - the steps to re sod the grass on 99% of homes is quite similar.

Take for example, a rails application. My client at example.com uses coffeescript for the frontend.A hot technology some time ago, nowadays its quite rare to come across it, and  most devs (including me) have forgotten the syntax for coffeescript. You application might use ES5 script.The point is that, on a rails application, there's many ways of doing things, unlike a lawn.Thousands on "variables" in software, in a re-sod or a plumbing situation, not many.

So someone comes and asks for an estimate. When it's the first thing I'm doing on a codebase, it's a foggy situation.

Let's say you're standing in a thick fog, and you client wants to you walk to the shed. Then he says, oh before you start walking can you estimate it, because I'm trying to keep it under 2k steps.As the "walker in the fog", you can see seven feet ahead. If you walk 7 feet, you can see 7 feet further. This is how it is when working with a rails application, especially when you're new to someones application.So the walkers estimate is: I don't know how many feet away the shed is, but I'm pretty sure it's not within 7 feet -  I'd see the shed if that's the case.

My answer is the same for the estimate. It's possible I might spent only 5 hours ($500)In the case that it costs over $2k, I won't know until I've spent $1.3k, because I can only see $7k/feet ahead.
## [7][AppSignal users, have you implemented Deploy Markers?](https://www.reddit.com/r/rails/comments/ese2zr/appsignal_users_have_you_implemented_deploy/)
- url: https://www.reddit.com/r/rails/comments/ese2zr/appsignal_users_have_you_implemented_deploy/
---
We currently use AppSignal for our logging, and now we want to implement deploy markers. My impression is that I just need to create a file './config/initializers/appsignal.rb' and add a single line:
```
ENV['APP_VERSION'] = Rails.application.secrets.version
```
My concern is that I'm unsure of which happens first, initializers getting run versus when the AppSignal gem loads './config/appaignal.yml.' In addition, does AppSignal read the environment variable APP\_VERSION for each log, or only once upon startup?
## [8][API + Frontend - Should I use webpack or isolate them?](https://www.reddit.com/r/rails/comments/esb7fx/api_frontend_should_i_use_webpack_or_isolate_them/)
- url: https://www.reddit.com/r/rails/comments/esb7fx/api_frontend_should_i_use_webpack_or_isolate_them/
---
I usually build monolithic rails apps, often with a React or Vue SPA frontend.

I am about to start developping a simple app with a RoR API. I would like to use React or Vue for the frontend, and I plan to develop a mobile app that will access the same API, but later, when I have more time to code.

Also, since this is a personal project I want to build to showcase my dev skills, I intend to develop the API and the frontend with several frameworks/languages (React, Crystal, Elixir, React, Vue, Svelte, Webcomponents) and compare them.

Would you advise me to create two separate projects/git repos (API + frontend) or should I go for webpack and bundle the frontend and the API together? Should I use git submodules?
## [9][Build a Reddit Clone in Rails 6 [New Video Series]](https://www.reddit.com/r/rails/comments/es6m26/build_a_reddit_clone_in_rails_6_new_video_series/)
- url: https://www.reddit.com/r/rails/comments/es6m26/build_a_reddit_clone_in_rails_6_new_video_series/
---
Hi guys, I'm back with a new Ruby on Rails build. This time we are building a lightweight Reddit clone. I've shared some previous Rails builds here and thought this might be on interest to anyone who enjoyed those past videos. Hope you guys enjoy the video.

**Part 1:**

In this video I cover the setup of a new rails app, installing the devise gem and creating user accounts. Add functionality for users to create a new community/subreddit. Add functionality to create new posts for a community (using nested routes).

[https://www.youtube.com/watch?v=aD6JvHKNPPM](https://www.youtube.com/watch?v=aD6JvHKNPPM)

**Part 2 (Coming soon):**

Signed in users can subscribe to communities and create new posts within them. Add functionality to allow subscribers to upvote / downvote posts and order them in community based on popularity. Add comments to posts.

&amp;#x200B;

The build is part of a collection of MVP's that I have been building after working in the tech startup space for some time and assisting in building minimum viable products.

I post new web app builds on YouTube using Ruby on Rails 6 so feel free to check out the channel and add some recommendations for upcoming builds.
## [10][Can a rails response contain a measurement of its own slowness?](https://www.reddit.com/r/rails/comments/esder2/can_a_rails_response_contain_a_measurement_of_its/)
- url: https://www.reddit.com/r/rails/comments/esder2/can_a_rails_response_contain_a_measurement_of_its/
---
Can a rails response contain a measurement of how much time was spent on activerecord, and view rendering.So this statistic:

![img](sph7dmx7kcc41 "
")

show here like this:   


    &lt;html&gt; 
    &lt;body&gt; 
    &lt;!-- stats  Views: 0.3ms .  ActiveRecord: 4.4ms --&gt;  
    &lt;/body&gt;  
    &lt;/html&gt;
## [11][Authlogic::Session::Existence::SessionInvalidError](https://www.reddit.com/r/rails/comments/es30kk/authlogicsessionexistencesessioninvaliderror/)
- url: https://www.reddit.com/r/rails/comments/es30kk/authlogicsessionexistencesessioninvaliderror/
---
I'm running minitests and I'm getting this random error 

 Authlogic::Session::Existence::SessionInvalidError:  Authlogic::Session::Existence::SessionInvalidError: Your session is invalid and has the following errors: Activate &amp; Log in

This doesn't appear for the say first 10 minitests and subsequently, all the later ones have this error. 

I'm not sure where to look for since this error started to appear suddenly. Any help would be great
