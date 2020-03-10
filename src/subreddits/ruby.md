# ruby
## [1][I wrote some code demonstrating subjects covered in an awesome talk by Squirrel Eiserloh: "Math for Game Programmers: Fast and Funky 1D Nonlinear Transformations" (link to the talk included in the Ruby comments)](https://www.reddit.com/r/ruby/comments/fg08h1/i_wrote_some_code_demonstrating_subjects_covered/)
- url: http://fiddle.dragonruby.org/?share=https://gist.github.com/amirrajan/f22dd4f31a0616ca7822208e18f8650e
---

## [2][Python Bot for automating Tinder swipes: Would it be possible to replicate a similar project in Ruby?](https://www.reddit.com/r/ruby/comments/fgasij/python_bot_for_automating_tinder_swipes_would_it/)
- url: https://youtu.be/lvFAuUcowT4
---

## [3][(Un)successful ActiveRecord refactor](https://www.reddit.com/r/ruby/comments/fg0sgl/unsuccessful_activerecord_refactor/)
- url: https://kukicola.io/posts/unsuccessful-activerecord-refactor
---

## [4][Ways to interact with Sharepoint or MS Graph using Ruby?](https://www.reddit.com/r/ruby/comments/fg1bmp/ways_to_interact_with_sharepoint_or_ms_graph/)
- url: https://www.reddit.com/r/ruby/comments/fg1bmp/ways_to_interact_with_sharepoint_or_ms_graph/
---
Hey guys,

I'm trying to work with Sharepoint and the MS Graph API. I tried it with javascript but wasn't really feeling it, so I'm asking you all about using Ruby to develop on the new Sharepoint Framework or calling the Graph API?
## [5][Having trouble architecting a side project](https://www.reddit.com/r/ruby/comments/ffw0r2/having_trouble_architecting_a_side_project/)
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
## [6][Sinatra authentication from users in a database](https://www.reddit.com/r/ruby/comments/ffuvfs/sinatra_authentication_from_users_in_a_database/)
- url: https://www.reddit.com/r/ruby/comments/ffuvfs/sinatra_authentication_from_users_in_a_database/
---
Imagine we have a database which has a table for users. Including their id's, usernames, passwords and other information we got from user in the process of signup. Database might be anything and so, it really doesn't matter which database type we're using here.

I usually used this method:

[https://github.com/prp-e/sinatra-basic-auth/blob/master/main.rb](https://github.com/prp-e/sinatra-basic-auth/blob/master/main.rb)

To perform a simple authentication, but for now, I'm looking for a better way to do the authentication and as I store my users in a db (such as sqlite), I really want to read user's credentials from the db and check if they're correct, then let them login.

And the next question, is not really related to what I asked earlier, but it's a matter of authentication and stuff, so I ask it here. When I want to check someone is logged in or not, I check the value of `session[:username]`, I was thinking about some algorithmic way to use this to add user's credentials to what they create (such as a post), because I would need user id or username who created the post or any other entity.
## [7][DatoRSS - RSS Search Engine with RESTful API](https://www.reddit.com/r/ruby/comments/ffw6ff/datorss_rss_search_engine_with_restful_api/)
- url: https://datorss.com
---

## [8][Ruby 2.7 - Trying out without blowing up my dev.](https://www.reddit.com/r/ruby/comments/ffqy9h/ruby_27_trying_out_without_blowing_up_my_dev/)
- url: https://www.reddit.com/r/ruby/comments/ffqy9h/ruby_27_trying_out_without_blowing_up_my_dev/
---
Hello,

I'd like to start trying out the newest Ruby version, but I'd want to do so without breaking my current dev environment for work. At present I use *rbenv*. What would be the safest way to use 2.6 and 2.7 concurrently without issues?
## [9][opmlparser gem - read / parse outlines (incl. feed subscription lists) in the OPML (Outline Processor Markup Language) format in xml](https://www.reddit.com/r/ruby/comments/ffs6lk/opmlparser_gem_read_parse_outlines_incl_feed/)
- url: https://github.com/feedreader/pluto/tree/master/opmlparser
---

## [10][Installing older version of ruby with RVM?](https://www.reddit.com/r/ruby/comments/ffpdcn/installing_older_version_of_ruby_with_rvm/)
- url: https://www.reddit.com/r/ruby/comments/ffpdcn/installing_older_version_of_ruby_with_rvm/
---
I factory reset my mac to the current version of Mac OS. I installed RVM because that's what I normally use. However, for work, we still have projects that run on ruby 2.2.4, so I tried installing it but it didn't work. New versions install correctly, however, such as 2.6.3.

When I try installing 2.2.4, I get this error:

    Error running 'env GEM_HOME=/Users/garcia/.rvm/gems/ruby-2.2.4@global GEM_PATH= /Users/garcia/.rvm/rubies/ruby-2.2.4/bin/ruby -d /Users/garcia/.rvm/src/rubygems-3.0.8/setup.rb --no-document',
    please read /Users/garcia/.rvm/log/1583728534_ruby-2.2.4/rubygems.install.log

However, I'm still able to see 2.2.4 when I run `rvm list`. When I switch to it via `rvm use --default 2.2.4`, everything seems to work, but when I run `gem update --system`. I get this error:

    ERROR:  Loading command: update (LoadError)
    	cannot load such file -- openssl
    ERROR:  While executing gem ... (NoMethodError)
        undefined method `invoke_with_build_args' for nil:NilClass

&amp;#x200B;

Any ideas? I submitted an issue ticket to the RVM Github repo like a week ago, but no one has responded.

&amp;#x200B;

**UPDATE:**

So general consensus was to use rbenv. I uninstalled RVM and switched to rbenv and it still didn't work! :( Here's the error:

    WARNING: ruby-2.2.4 is past its end of life and is now unsupported.
    It no longer receives bug fixes or critical security updates.
    
    ruby-build: using readline from homebrew
      
    BUILD FAILED (OS X 10.15.3 using ruby-build 20200224)
    
    Inspect or clean up the working tree at /var/folders/y6/47yxhmmd5p90jgh2l95vsb980000gn/T/ruby-build.20200309143540.49296.Fp4D8d
    Results logged to /var/folders/y6/47yxhmmd5p90jgh2l95vsb980000gn/T/ruby-build.20200309143540.49296.log
    
    Last 10 log lines:
    installing capi-docs:         /Users/garcia/.rbenv/versions/2.2.4/share/doc/ruby
    The Ruby openssl extension was not compiled.
    ERROR: Ruby install aborted due to missing extensions
    Configure options used:
      --prefix=/Users/garcia/.rbenv/versions/2.2.4
      --with-openssl-dir=/usr/local/opt/openssl@1.1
      --with-readline-dir=/usr/local/opt/readline
      CC=clang
      LDFLAGS=-L/Users/garcia/.rbenv/versions/2.2.4/lib 
      CPPFLAGS=-I/Users/garcia/.rbenv/versions/2.2.4/include 

&amp;#x200B;

I am completely brand new to rbenv so I'm not sure the best way to troubleshoot this issue. Although, from my research (and also as someone has said in the comments) I've learned that ruby versions less than 2.3.x don't support openssl 1.1, however, I never had this problem before I reset my computer which is why it's so confusing.

Also, I followed the rbenv installation guide from it's github repo page.

Any help is super appreciated.

&amp;#x200B;

**Another Update**

So, the solution for me was to, yet again, restore my computer to factory settings and not even bother with RVM. I installed rbenv and was able to successfully install 2.2.4 without any warnings or errors.

However, when I try to run \``bundle install`\` on an old rails app (one that i'm working for my job), it gives me an error saying `The ruby version specified by your gemfile is 2.2.4, but your current version of ruby is 2.6.3`, but when I run `ruby -v` the output is `ruby 2.2.4p230 (2015-12-16 revision 53155) [x86_64-darwin19]` so I'm actually super confused because the errors are not making sense.

So, I guess this solves my ruby problem? My rails problem isn't solved yet, but I figured this is probably a question for r/rails  and not necessarily r/ruby. I'm actually not sure where this new problem belongs. It's a ruby error just within a rails context, so I guess it could belong in both.

If you any ideas, I would super appreciate it!

Also, I must say that this community has helped me so much more than stack overflow or even github. I'm definitely going to become a member of this subreddit and I'm in shock I wasn't already!
