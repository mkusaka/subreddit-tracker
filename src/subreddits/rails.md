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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/evmx0w/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/evmx0w/personal_projects_show_off_your_own_project_andor/
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
## [3][ActionCable (Redis) -&gt; reconnect on redis connection error](https://www.reddit.com/r/rails/comments/ewn1e4/actioncable_redis_reconnect_on_redis_connection/)
- url: https://www.reddit.com/r/rails/comments/ewn1e4/actioncable_redis_reconnect_on_redis_connection/
---
Hello! 

I work on an application which has ActionCable implemented with Redis.

I would like to know whether it is possible to reset the entire ActionCable setup in case the Redis throws Connection error.

I would like to connect to a different Redis url when the connection error occurs. The infrastructure of this Redis service is that there are two Redis databases mirrored, and there are frequent planned maintanance on the service which can cause the shut down of one of the databases. Unfortunately the switch is not solved by the service itself, however it provides us 2 connection urls. These are the ones I would like to switch between automatically.

Is it possible to do with ActionCable.

I know it is possible with ActiveRecord database connection
## [4][What is everyone using as wrapper for hybrid app?](https://www.reddit.com/r/rails/comments/ewlmre/what_is_everyone_using_as_wrapper_for_hybrid_app/)
- url: https://www.reddit.com/r/rails/comments/ewlmre/what_is_everyone_using_as_wrapper_for_hybrid_app/
---
I've got a working PWA built with Rails and Turbolinks.  It has notifications too (only working on Android).

But now a requirement has arisen that we need to be able to use Bluetooth (both on Android and iOS) so I'm afraid a PWA will no longer suffice.

My first thought was to use the Turbolinks wrappers:
https://github.com/turbolinks/turbolinks-android and https://github.com/turbolinks/turbolinks-ios but the one for Android has been deprecated for a long time and there don't seem to come any updates.

What would be good alternatives?  Has anyone tried Cordova in combination with Rails (and Turbolinks)?  Any other suggestions.  Mind you: it's not a single page app, it's a normal Rails app, that currently also works as a PWA.

Curious to hear what others are using.
## [5][What is everyone doing for caching images served by ActiveStorage? I see the pull-request 34477 has not been merged](https://www.reddit.com/r/rails/comments/ewe9fw/what_is_everyone_doing_for_caching_images_served/)
- url: https://www.reddit.com/r/rails/comments/ewe9fw/what_is_everyone_doing_for_caching_images_served/
---
[https://github.com/rails/rails/pull/34477/files](https://github.com/rails/rails/pull/34477/files)

&amp;#x200B;

looks like a great solution. Has anyone turned this into a GEM that needs a signal boost or what?!?!
## [6][Tests and CI/Cd - How to start from almost nothing?](https://www.reddit.com/r/rails/comments/ewa9ah/tests_and_cicd_how_to_start_from_almost_nothing/)
- url: https://www.reddit.com/r/rails/comments/ewa9ah/tests_and_cicd_how_to_start_from_almost_nothing/
---
Our codebase is a mess. It's a huge app and for the longest time we managed to survive without tests. Now we have a bit of time to do some tests. Aside from making sure that new code must have tests, how can we slowly improve our test coverage? Do you have tips? We have more than 15 people pushing code to the app's repo every day.
## [7][installed bootstrap 4.4 on rails 6](https://www.reddit.com/r/rails/comments/ewfyqt/installed_bootstrap_44_on_rails_6/)
- url: https://www.reddit.com/r/rails/comments/ewfyqt/installed_bootstrap_44_on_rails_6/
---
I installed bootstrap 4.4 on rails 6 by following this gist: [https://gist.github.com/bazzel/2c64e2e5804077f9a61938a93ed54823](https://gist.github.com/bazzel/2c64e2e5804077f9a61938a93ed54823)

What I'm not sure of is:

is this line

    &lt;%= stylesheet_pack_tag 'styles', media: 'all', 'data-turbolinks-track': 'reload' %&gt;

supposed to replace this pre-existing line:

    &lt;%= stylesheet_link_tag 'application', media: 'all', 'data-turbolinks-track': 'reload' %&gt;

or added next to it.
## [8][Create a rest SSL request](https://www.reddit.com/r/rails/comments/ewfca1/create_a_rest_ssl_request/)
- url: https://www.reddit.com/r/rails/comments/ewfca1/create_a_rest_ssl_request/
---
Hello guys I have a problem, I'm trying to create a connection with a external server with ssl credentials but I get errors with the authentication the members from the service told me that I need to add a pem file in my keystore but it's from java and Im not sure about how do it in rails I try with this code but I can't connect with the server, if anyone can help me I will be very grateful

&amp;#x200B;

RestClient::Request.execute(

method: "PUT",

url: "#{ENV\['STP\_HOST'\]}/ordenPago/registra",

payload: json\_stp\_registra,

headers: {

"Content-Type" =&gt; "application/json",

accept: "application/json"

}).body
## [9][How do you write request specs for APIs?](https://www.reddit.com/r/rails/comments/ew8tyr/how_do_you_write_request_specs_for_apis/)
- url: https://www.reddit.com/r/rails/comments/ew8tyr/how_do_you_write_request_specs_for_apis/
---
Hi all,

I've been trying to push myself deeper into thinking about effective and elegant testing strategies and I'm struggling to come up with a flow for writing request specs for my JSON API. Right now, I feel like I'm just randomly coming up with things to test and the code comes out really repetitive. Do you guys have any tools/blogs/resources to help me figure out where to start here?
## [10][Runtime error](https://www.reddit.com/r/rails/comments/ew9oy3/runtime_error/)
- url: https://www.reddit.com/r/rails/comments/ew9oy3/runtime_error/
---
 

I tried doing Rails s and was prompted with this error Webpacker configuration file not found /home/user1/Mycollect/config/webpacker.yml. Please run rails webpacker:install Error: No such file or directory @ rb\_sysopen - /home/user1/Mycollect/config/webpacker.yml (RuntimeError) 

W​hen I then try to run the command rails webpacker:install it gives me this error,   


rails aborted!  
ArgumentError: Malformed version number string 0.32+git  
/home/user1/Mycollect/vendor/cache/gems/webpacker-4.2.2/lib/tasks/webpacker/check\_yarn.rake:12:in \`block (2 levels) in &lt;main&gt;'  
/home/user1/Mycollect/vendor/cache/gems/railties-6.0.2.1/lib/rails/commands/rake/rake\_command.rb:23:in \`block in perform'  
/home/user1/Mycollect/vendor/cache/gems/railties-6.0.2.1/lib/rails/commands/rake/rake\_command.rb:20:in \`perform'  
/home/user1/Mycollect/vendor/cache/gems/railties-6.0.2.1/lib/rails/command.rb:48:in \`invoke'  
/home/user1/Mycollect/vendor/cache/gems/railties-6.0.2.1/lib/rails/commands.rb:18:in \`&lt;main&gt;'  
/home/user1/Mycollect/vendor/cache/gems/bootsnap-1.4.5/lib/bootsnap/load\_path\_cache/core\_ext/kernel\_require.rb:22:in \`require'  
/home/user1/Mycollect/vendor/cache/gems/bootsnap-1.4.5/lib/bootsnap/load\_path\_cache/core\_ext/kernel\_require.rb:22:in \`block in require\_with\_bootsnap\_lfi'  
/home/user1/Mycollect/vendor/cache/gems/bootsnap-1.4.5/lib/bootsnap/load\_path\_cache/loaded\_features\_index.rb:92:in \`register'  
/home/user1/Mycollect/vendor/cache/gems/bootsnap-1.4.5/lib/bootsnap/load\_path\_cache/core\_ext/kernel\_require.rb:21:in \`require\_with\_bootsnap\_lfi'  
/home/user1/Mycollect/vendor/cache/gems/bootsnap-1.4.5/lib/bootsnap/load\_path\_cache/core\_ext/kernel\_require.rb:30:in \`require'  
/home/user1/Mycollect/vendor/cache/gems/activesupport-6.0.2.1/lib/active\_support/dependencies.rb:325:in \`block in require'  
/home/user1/Mycollect/vendor/cache/gems/activesupport-6.0.2.1/lib/active\_support/dependencies.rb:291:in \`load\_dependency'  
/home/user1/Mycollect/vendor/cache/gems/activesupport-6.0.2.1/lib/active\_support/dependencies.rb:325:in \`require'  
/home/user1/Mycollect/bin/rails:9:in \`&lt;top (required)&gt;'  
/home/user1/Mycollect/vendor/cache/gems/spring-2.1.0/lib/spring/client/rails.rb:28:in \`load'  
/home/user1/Mycollect/vendor/cache/gems/spring-2.1.0/lib/spring/client/rails.rb:28:in \`call'  
/home/user1/Mycollect/vendor/cache/gems/spring-2.1.0/lib/spring/client/command.rb:7:in \`call'  
/home/user1/Mycollect/vendor/cache/gems/spring-2.1.0/lib/spring/client.rb:30:in \`run'  
/home/user1/Mycollect/vendor/cache/gems/spring-2.1.0/bin/spring:49:in \`&lt;top (required)&gt;'  
/home/user1/Mycollect/vendor/cache/gems/spring-2.1.0/lib/spring/binstub.rb:11:in \`load'  
/home/user1/Mycollect/vendor/cache/gems/spring-2.1.0/lib/spring/binstub.rb:11:in \`&lt;top (required)&gt;'  
/home/user1/Mycollect/bin/spring:15:in \`&lt;top (required)&gt;'  
bin/rails:3:in \`load'  
bin/rails:3:in \`&lt;main&gt;'  
Tasks: TOP =&gt; webpacker:install =&gt; webpacker:check\_yarn  
(See full trace by running task with --trace)  
Wondering if anyone knows a quick fix?
## [11][Handy cheat sheet comparing all the different ways to set attributes with ActiveRecord on Rails 6](https://www.reddit.com/r/rails/comments/ewb0fv/handy_cheat_sheet_comparing_all_the_different/)
- url: https://www.reddit.com/r/rails/comments/ewb0fv/handy_cheat_sheet_comparing_all_the_different/
---
I put together a handy cheat sheet for quickly comparing all the different ways you can set attributes on an Active Record object in Rails 6. At a glance the cheat sheet shows you which methods save to the database, run validations, run callbacks, update timestamps, etc.

[https://scottbartell.com/2020/01/30/set-attributes-in-active-record-rails-6/](https://scottbartell.com/2020/01/30/set-attributes-in-active-record-rails-6/)

&amp;#x200B;

A bit of context... a few years ago I came across a similar cheat sheet created by [David Verhasselt for Rails 4](https://davidverhasselt.com/set-attributes-in-activerecord/) and I found it to be such a useful reference that I eventually made my own for [Rails 5](https://scottbartell.com/2019/07/15/set-attributes-in-active-record-rails-5/). The Rails 5 guide has continued to be useful in my day-to-day so I made this new one for [Rails 6](https://scottbartell.com/2020/01/30/set-attributes-in-active-record-rails-6/)!

&amp;#x200B;

I'm very open to feedback / suggestions to make this an even more valuable resource -- please share your thoughts!
## [12][Best ways to do RESTful authentication](https://www.reddit.com/r/rails/comments/ew3pwo/best_ways_to_do_restful_authentication/)
- url: https://www.reddit.com/r/rails/comments/ew3pwo/best_ways_to_do_restful_authentication/
---
I have finished my project and it's ready for being deployed. I used different technologies and tools such as MySQL and Devise. And of course, I asked a lot of questions (I can say at least 5 of them were in this sub :D) to learn more. 

Now, I have a bunch of questions and ambiguities. I have plans of having an android app for our app. I even talked to an android developer and she said she needs an API (which is not surprising) and these came to my mind : 

* For login, I can't simply send a string as password. I need a token authentication. But I couldn't find any clear explanation on how to do that. I found libraries and gems, but most of them had no clear documentation. I used devise for managing my users, so I of course need a way to do token authentication with devise. 
* As the project will be launched in a day or two, and the app takes time to be published, I have another ambiguity. How can force previous users to get a token? 
* And what about sign-ups? Isn't it unsafe to send information as plain text to the server? or sign-ups can have a token too? 

These were what came into my mind. I'll be thankful of your answers and discussions.
