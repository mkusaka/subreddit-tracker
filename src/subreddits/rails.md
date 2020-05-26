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
## [2][Best Way to Implement Multi Column Search Feature](https://www.reddit.com/r/rails/comments/gqtri1/best_way_to_implement_multi_column_search_feature/)
- url: https://www.reddit.com/r/rails/comments/gqtri1/best_way_to_implement_multi_column_search_feature/
---
So I wanted to add search feature which would look for LIKE results in multiple columns of the table, ie author (via relation), description, comment et cetc. Can anyone recommend me best approach, gem for this? :) I saw multiple gems but most of them seem too bulky for my simple needs.
## [3][Rails 6 and parallel tests: get ID/number of current executor](https://www.reddit.com/r/rails/comments/gqtb87/rails_6_and_parallel_tests_get_idnumber_of/)
- url: https://www.reddit.com/r/rails/comments/gqtb87/rails_6_and_parallel_tests_get_idnumber_of/
---
I need to try something with Capybara to attempt to solve an issueand I need to set a different server port for each executor when I am using parallel tests. With old parallelisation gems you could use the env variable TEST_ENV_NUMBER to get a different number for each executor process, but that doesn't seem to work with Rails' own parallel tests and I can't find anything similar. Any idea? Thanks
## [4][When to use Engines for breaking up monolith?](https://www.reddit.com/r/rails/comments/gqjmjv/when_to_use_engines_for_breaking_up_monolith/)
- url: https://www.reddit.com/r/rails/comments/gqjmjv/when_to_use_engines_for_breaking_up_monolith/
---
We have two Rails repos that use a Rails models gem we have.  This gem contains all the Rails models.  One is the main repo that serves as a backend for the frontend application.  The other small repo is an API for direct client access.

It's a pain to use this models gem as we need to tag releases and it's a lot of overhead.  Now we're combining the two repos together while separating them via Rails Engines.  I personally don't like the approach and wish we just had a regular monolith.  I don't understand why we don't just namespace the API.

It's been a long project that isn't complete yet that's converting it into this monolith with Engines.  I'm concerned because if we had a normal monolith we would have been done already with I think easier-to-manage code.  We could have spent that time upgrading Rails.  Now that will probably be delayed I'm guessing.

Am I wrong thinking this is over-engineered?  If it is over-engineered, how do you convince your coworkers and manager?
## [5][Working with Rails 6 / Webpack / Syntax](https://www.reddit.com/r/rails/comments/gqq7jm/working_with_rails_6_webpack_syntax/)
- url: https://www.reddit.com/r/rails/comments/gqq7jm/working_with_rails_6_webpack_syntax/
---
I'm having some errors when trying to compile react js stuff. I'm currently migrating a project from rails 5 to 6 so the additional dependencies had to be configured manually.

It happens when I use something like this in a js file:

```
someArray.push(&lt;SomeComponent/&gt;)
```

The error message:

```
Module parse failed: Unexpected token (34:6)
You may need an appropriate loader to handle this file type, currently no loaders are configured to process this file.
```

What I've already done:

* included @babel/core, @babel/preset-env, @babel-preset-react, @babel-preset-typescript
* Did a `rails webpacker:install:react`
## [6][RSpec book - Aaron Sumner](https://www.reddit.com/r/rails/comments/gqh0q9/rspec_book_aaron_sumner/)
- url: https://www.reddit.com/r/rails/comments/gqh0q9/rspec_book_aaron_sumner/
---
Hello everyone, I´m trying to improve my tests and learn more about RSpec. I found this e-book from [Aaron Sumner - Every day Rails Testing with RSpec](https://leanpub.com/everydayrailsrspec) and I would like to know your opinions since I found a lot of "tutorials" and "videos" but nothing concrete except for this publication. There are even some references from other authors to this one. I would like to purchase it but also would like to know your thoughts about this one in specific plus other recommendations if possible.
## [7][Solo semi-noob working on complex site to start side business](https://www.reddit.com/r/rails/comments/gqm0bq/solo_seminoob_working_on_complex_site_to_start/)
- url: https://www.reddit.com/r/rails/comments/gqm0bq/solo_seminoob_working_on_complex_site_to_start/
---
Has there been anyone in this situation that can help me see the light at the end of the tunnel?

I'm a decent programmer, although instead of mastering any language I've been kind of forced to jump around in my career in IT. I fell in love with Ruby for automating my sysadmin tasks, self wrote an internal rails app to automate even more, and now have a lot more free time.

I have an idea for a side business, but it's a big task with lots of code to write which requires a lot of reading the documentation (on both advanced rails stuff as well as the gems I'm using) since I'm not a webdev or software engineer. I'd love to partner with someone who knows rails better than me, but obviously choosing a business partner is a major decision and I don't know any rails devs in my area, much less someone I'd trust as a partner. I also can't afford to hire anyone, so having help on this is out.

Has anyone here been in this situation that can give me advice? Mistakes to avoid? Or a success story to brag about and give me more hope? Today has been a bit hard as I remember a friend I served with that once told me if I ever started a business he would join my team in a heartbeat. As a side note (and I hope this is allowed here), I'd like to plug the Special Operations Warrior Foundation, which I'll be donating a percentage of profits to if I ever have any.
## [8][Restart nginx throws error and crashes rails app](https://www.reddit.com/r/rails/comments/gqg1v5/restart_nginx_throws_error_and_crashes_rails_app/)
- url: https://www.reddit.com/r/rails/comments/gqg1v5/restart_nginx_throws_error_and_crashes_rails_app/
---
Ruby on Rails beginner here.

I have an application that makes requests to an application server that is run on Rails, Nginx, and Passenger.

I want to edit one file on the application server. This is a small file called mailer.rb that controls sending mail to clients. When by reading online, to see my changes I needed to restart my Nginx service. I did so with the code `sudo service nginx reload`.

When I reloaded nginx, the application throws an error page, as attached.

https://preview.redd.it/bpxhi2tncy051.png?width=1263&amp;format=png&amp;auto=webp&amp;s=fd761c7d443f7a6f79391ee6d59ad2a26a3cfdcb

When looking into the error logs, the response it is giving me is as follows:

    App 27766 output: Error: The application encountered the following error: cannot load such file -- /var/www/my-backend/code/config/environment (LoadError)
    App 27766 output:
    App 27766 output:     config.ru:3:in \require_relative'`
    App 27766 output:     config.ru:3:in \block in &lt;main&gt;'`
    App 27766 output: /var/www/my-backend/code/vendor/bundle/ruby/2.4.0/gems/rack-2.0.7/lib/rack/builder.rb:55:in \instance_eval'`
    App 27766 output: /var/www/my-backend/code/vendor/bundle/ruby/2.4.0/gems/rack-2.0.7/lib/rack/builder.rb:55:in \initialize'`
    App 27766 output:     config.ru:1:in \new'`
    App 27766 output:     config.ru:1:in \&lt;main&gt;'`
    App 27766 output: /usr/share/passenger/helper-scripts/rack-preloader.rb:101:in \eval'`
    App 27766 output: /usr/share/passenger/helper-scripts/rack-preloader.rb:101:in \preload_app'`
    App 27766 output: /usr/share/passenger/helper-scripts/rack-preloader.rb:189:in \block in &lt;module:App&gt;'`
    App 27766 output: /usr/lib/ruby/vendor_ruby/phusion_passenger/loader_shared_helpers.rb:380:in \run_block_and_record_step_progress'`
    App 27766 output: /usr/share/passenger/helper-scripts/rack-preloader.rb:188:in \&lt;module:App&gt;'`
    App 27766 output: /usr/share/passenger/helper-scripts/rack-preloader.rb:30:in \&lt;module:PhusionPassenger&gt;'`
    App 27766 output: /usr/share/passenger/helper-scripts/rack-preloader.rb:29:in \&lt;main&gt;'`
    [ E 2020-05-25 11:52:13.5984 1351/T2x age/Cor/App/Implementation.cpp:221 ]: Could not spawn process for application /var/www/my-backend/code: The application encountered the following error: cannot load such file -- /var/www/my-backend/code/config/environment (LoadError)
    Error ID: 7e0ef459
    Error details saved to: /tmp/passenger-error-FqS6Dm.html

&amp;#x200B;

I reverted my changes to the file and tried restarting again but the problem persists. It seems that the file change itself wasn't the problem, but anytime I restart Nginx this happens. Any help would be appreciated.
## [9][Web App Development](https://www.reddit.com/r/rails/comments/gqahgr/web_app_development/)
- url: https://www.reddit.com/r/rails/comments/gqahgr/web_app_development/
---
Hello Everyone. I've been surfing through the internet on how to create a web app but so far I haven't gotten enough Information. 
First Question: Can I create a cross platform web app with Ruby on Rails?
2nd: Can I create the app alone with no other team member?
3rd: Will the back end,frontend and database be created on one the same framework?
## [10][dynamically creating ActionCable subscriptions in asset pipeline](https://www.reddit.com/r/rails/comments/gpx1hi/dynamically_creating_actioncable_subscriptions_in/)
- url: https://www.reddit.com/r/rails/comments/gpx1hi/dynamically_creating_actioncable_subscriptions_in/
---
In expanding that DHH YouTube tutorial for an ActionCable chatroom, I came across a [much more fleshed out example](https://www.thegreatcodeadventure.com/rails-5-action-cable-with-multiple-chatroom-subscriptions/) that loops through all existing conversations and creates a subscription for each. This approach "works," but it essentially hardcodes the app's current number of conversations into the asset pipeline, and then leaves it as is until the assets are regenerated. So really, it's not too usable.

What approaches have you taken to dynamically create ActionCable subscriptions?

**assets/javascripts/channels/messages.coffee**

    &lt;% Conversation.all.each do |conversation| %&gt;
        App['conversation_' + &lt;%= conversation.id %&gt;] = App.cable.subscriptions.create { 
          channel: "MessagesChannel", conversation: &lt;%= conversation.id %&gt; 
        }, 
        {
          connected: -&gt;
            console.log 'Connected'
    
          disconnected: -&gt;
            console.log 'Disconnected'
    
          received: (data) -&gt;
            $("[data-conversation-id='" + data.conversation_id + "']").append(data.message)
            $("#conversation-main").scrollTop($("#conversation-main")[0].scrollHeight);
    
          speak: (body, conversation_id) -&gt;
            @perform 'speak', body: body, conversation_id: conversation_id
    
          set_conversation_id: (conversation_id) -&gt;
            console.log conversation_id
            this.conversation_id = conversation_id
        }
    
      &lt;% end %&gt;
    
      $(document).on 'turbolinks:load', -&gt;
        submit_message()
        scroll_bottom()
    
      submit_message = () -&gt;
        $('#response').on 'keydown', (event) -&gt;
          if event.keyCode is 13
            conversation_id = $("[data-conversation-id]").data("conversation-id")
            # values = $(this).serializeArray()
            App['conversation_' + conversation_id].set_conversation_id(conversation_id)
            App['conversation_' + conversation_id].speak(event.target.value, conversation_id)
            event.target.value = ""
            event.preventDefault()
    
      scroll_bottom = () -&gt;
        if $('#messages').length &gt; 0
          $('#messages').scrollTop($('#messages')[0].scrollHeight)
## [11][Action Text Question](https://www.reddit.com/r/rails/comments/gq5s7r/action_text_question/)
- url: https://www.reddit.com/r/rails/comments/gq5s7r/action_text_question/
---
Anybody happen to know whether Action Text fields take in html code and display it on the view?
