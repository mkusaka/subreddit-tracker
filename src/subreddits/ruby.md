# ruby
## [1][RailsConf 2020 Cancelled](https://www.reddit.com/r/ruby/comments/fgjhcx/railsconf_2020_cancelled/)
- url: https://railsconf.com/
---

## [2][Roast my terminal app :)](https://www.reddit.com/r/ruby/comments/fgslkn/roast_my_terminal_app/)
- url: https://www.reddit.com/r/ruby/comments/fgslkn/roast_my_terminal_app/
---
Hey guys :)

Im currently working on a project for school and I would be stoked if you guys could roast it and tell me what i should change :)

any feedback is welcome as I am just learning

Essentially I've built a terminal app to track staffs moral

&amp;#x200B;

[https://github.com/alex1092/MoralChecker](https://github.com/alex1092/MoralChecker)

Also if anyone has any tips with making uml diagrams that would be awesome :) 
## [3][Ressources to improve after a few years?](https://www.reddit.com/r/ruby/comments/fgfx05/ressources_to_improve_after_a_few_years/)
- url: https://www.reddit.com/r/ruby/comments/fgfx05/ressources_to_improve_after_a_few_years/
---
I've been a Ruby developer for four years, and feel like I'm not as good as I should be.

For example a recruiter asked me if I used XXX (don't remember the word). It was just a fancy word for blocks, but I didn't know the word and it was one of the main reason I wasn't hired (or so he told me).

Most things I've read online is for new developpers, not really to improve as an already pretty experienced dev, do you have any places to recommend?  
(I know about the Ruby Bibliography)
## [4][How to create rake task for importing javascript libraries](https://www.reddit.com/r/ruby/comments/fgjuph/how_to_create_rake_task_for_importing_javascript/)
- url: https://www.reddit.com/r/ruby/comments/fgjuph/how_to_create_rake_task_for_importing_javascript/
---
I started using ruby on rails with its 5.2 version and since the newest rails version(not correlated) was released I haven't been coding that much in rails/ruby. I got used to using asset pipeline for importing javascript or css libraries. And then the switch to webpacker happened, which in my opinion hasn't made things much simpler. Long story short, I went through few tutorials for implementing bootstrap into rails project and I thought it could be nice to have a rake task for importing javascript libraries. 

I fell in love with rails because it frees you from great amount of "manual" labor and gives you time to spend energy on being creative. 

So basically I'd like to perform automatic installation of bootstrap into rails project with the command like this: 

`rake bootstrap:install`

I've came up with an idea that this task should do this

1. `system("yarn install bootstrap")` \- to install javascript dependency
2. create file app/javascript/stylesheets/application.scss
   1. fill it with this text "`@import '~bootstrap';`"
3. append to app/javascript/packs/applications.js text below:

* `require("bootstrap/dist/js/bootstrap")`
* `require("stylesheets/application")`

I'd like to get some opinions about this, is there something wrong with my approach, what's the best way to do it etc.
## [5][Python Bot for automating Tinder swipes: Would it be possible to replicate a similar project in Ruby?](https://www.reddit.com/r/ruby/comments/fgasij/python_bot_for_automating_tinder_swipes_would_it/)
- url: https://youtu.be/lvFAuUcowT4
---

## [6][I wrote some code demonstrating subjects covered in an awesome talk by Squirrel Eiserloh: "Math for Game Programmers: Fast and Funky 1D Nonlinear Transformations" (link to the talk included in the Ruby comments)](https://www.reddit.com/r/ruby/comments/fg08h1/i_wrote_some_code_demonstrating_subjects_covered/)
- url: http://fiddle.dragonruby.org/?share=https://gist.github.com/amirrajan/f22dd4f31a0616ca7822208e18f8650e
---

## [7][(Un)successful ActiveRecord refactor](https://www.reddit.com/r/ruby/comments/fg0sgl/unsuccessful_activerecord_refactor/)
- url: https://kukicola.io/posts/unsuccessful-activerecord-refactor
---

## [8][Ways to interact with Sharepoint or MS Graph using Ruby?](https://www.reddit.com/r/ruby/comments/fg1bmp/ways_to_interact_with_sharepoint_or_ms_graph/)
- url: https://www.reddit.com/r/ruby/comments/fg1bmp/ways_to_interact_with_sharepoint_or_ms_graph/
---
Hey guys,

I'm trying to work with Sharepoint and the MS Graph API. I tried it with javascript but wasn't really feeling it, so I'm asking you all about using Ruby to develop on the new Sharepoint Framework or calling the Graph API?
## [9][Having trouble architecting a side project](https://www.reddit.com/r/ruby/comments/ffw0r2/having_trouble_architecting_a_side_project/)
- url: https://www.reddit.com/r/ruby/comments/ffw0r2/having_trouble_architecting_a_side_project/
---
Hey guys. I'm a pretty new developer, and I want to start working on a side project. Here's sort of a high level idea of what I want to do (mainly focusing on the backend side of things, don't really care about the front end.. for now)

The app (or just a microservice? not sure) will basically have an endpoint a user can hit, and it will retrieve real-time transportation data from another endpoint (which I have access to). Since the data retrieved from this endpoint doesn't have everything I need, I need to hit another endpoint to retrieve more data, then essentially stitch everything together into a JSON and return it to the user in a friendly way. I'm having a lot of trouble figuring out how to architect this...

1. I want to use Rails, because its good for my career progression, but I feel like this app will be lightweight enough to just use Sinatra, but people discourage it for beginners. Any thoughts? I won't be really dealing with creating users or anything - solely just having an endpoint that returns a bunch of data to the user.
2. Lets say I were to scale this app, and I need to handle a bunch of requests at the same time, but don't want to do synchronous processing. Should I use something like Sidekiq, to have background jobs that just listen to a queue where data gets pushed to every time we get a request? Was thinking something like...

\- Have a worker to fetch the realtime data

\- Have another worker to get the data from another endpoint

\- When done, these workers will push the data to Redis, and I can query it to stitch together a JSON to return to the user

3) The feed I get realtime data from gets updated every 30 seconds. What is the most efficient way to pull down this feed? Should I use something like Cron to set up some schedules for some workers?

Really sorry if any of these questions sound stupid or not, I'm just trying to figure out the best way to approach things, rather than wing it and build a crappy, architected app.
## [10][Sinatra authentication from users in a database](https://www.reddit.com/r/ruby/comments/ffuvfs/sinatra_authentication_from_users_in_a_database/)
- url: https://www.reddit.com/r/ruby/comments/ffuvfs/sinatra_authentication_from_users_in_a_database/
---
Imagine we have a database which has a table for users. Including their id's, usernames, passwords and other information we got from user in the process of signup. Database might be anything and so, it really doesn't matter which database type we're using here.

I usually used this method:

[https://github.com/prp-e/sinatra-basic-auth/blob/master/main.rb](https://github.com/prp-e/sinatra-basic-auth/blob/master/main.rb)

To perform a simple authentication, but for now, I'm looking for a better way to do the authentication and as I store my users in a db (such as sqlite), I really want to read user's credentials from the db and check if they're correct, then let them login.

And the next question, is not really related to what I asked earlier, but it's a matter of authentication and stuff, so I ask it here. When I want to check someone is logged in or not, I check the value of `session[:username]`, I was thinking about some algorithmic way to use this to add user's credentials to what they create (such as a post), because I would need user id or username who created the post or any other entity.
