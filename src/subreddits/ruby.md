# ruby
## [1][Backports 3.16 released: Support of Ruby 2.7's features in earlier Ruby versions](https://www.reddit.com/r/ruby/comments/f0erf3/backports_316_released_support_of_ruby_27s/)
- url: https://github.com/marcandre/backports/blob/master/CHANGELOG.rdoc#version-3160---feb-6th-2020
---

## [2][a bit off-top](https://www.reddit.com/r/ruby/comments/f0qd71/a_bit_offtop/)
- url: https://www.reddit.com/r/ruby/comments/f0qd71/a_bit_offtop/
---
What theme is this?

https://preview.redd.it/rhjadzngpof41.png?width=664&amp;format=png&amp;auto=webp&amp;s=2c6d88a72ff71b858a35f9cf872246540a890071
## [3][Setting up your next Rails project with GitlabCI - a 101 guide](https://www.reddit.com/r/ruby/comments/f0ahyv/setting_up_your_next_rails_project_with_gitlabci/)
- url: https://hixonrails.com/ruby-on-rails-tutorials/ruby-on-rails-set-up-on-gitlab-with-gitlabci/
---

## [4][Ruby Quiz - Challenge #17 - Build an HTML Template Engine Like It's 1999](https://www.reddit.com/r/ruby/comments/f0f3kj/ruby_quiz_challenge_17_build_an_html_template/)
- url: https://github.com/planetruby/quiz/tree/master/017
---

## [5][Question on ruby and websockets](https://www.reddit.com/r/ruby/comments/f0evse/question_on_ruby_and_websockets/)
- url: https://www.reddit.com/r/ruby/comments/f0evse/question_on_ruby_and_websockets/
---
Hey Guys.
Lately i have been experimenting with ruby and websockets.
So i created a new Rails 5 project with ActionCable, all seems to work fine with it.

But also i created a ruby plain script with the [Faye's ruby websocket client](https://github.com/faye/faye-websocket-ruby). Unlike most tutorials on internet i want to try a server side (as a client) script, not a frontend JS script inside an HTML file.

So i tried the basic usage of it and i successfully make the handshake to perform correctly but i can't continue testing because i cant figure where to subscribe after connected to a desired channel exposed in the Rails server.

Here is my ruby script code:


    require 'faye/websocket'
    require 'eventmachine'
    
    EM.run {
      ws = Faye::WebSocket::Client.new('ws://localhost:3001/cable',nil,{
        headers: {'Origin' =&gt; 'ws://localhost:3001/cable'}
      })
    
      ws.on :open do |event|
        p [:open]
        ws.send({data: 'hello'})
      end
    
      ws.on :message do |event|
        p [:message, event.data]
      end
    
      ws.on :close do |event|
        p [:close, event.code, event.reason]
        ws = nil
      end
    
      ws.on :error do |event|
        p [:close, event.code, event.reason]
        ws = nil
      end
    
      ws.send({data: 'yoyoyooy'}) # This gets sent to nowhere..
      # I was hoping into subscribing a channel and callbacks for that channel, something like:
      # ws.subscribe('my-channel',receive_message_callback,error_callback)
    }


I have reviewed many other websocket ruby gems but i don't see any of them mentioning anything about subscribing to channels or something like, only for JS libraries, so maybe there is something that i dont understand.

Any one can give me a light on this?
Thanks in advance!
## [6][So what have your thoughts been on Artichoke Ruby?](https://www.reddit.com/r/ruby/comments/ezx6t8/so_what_have_your_thoughts_been_on_artichoke_ruby/)
- url: https://www.reddit.com/r/ruby/comments/ezx6t8/so_what_have_your_thoughts_been_on_artichoke_ruby/
---
I saw [Artichoke Ruby](https://github.com/artichoke/artichoke) mentioned [several](https://www.reddit.com/r/ruby/comments/clu9pd/artichoke_a_ruby_made_with_rust/) months ago and it seemed like a lot of folks were excited about this at the time. Fast forward to now, looking at the Github [commits](https://github.com/artichoke/artichoke/commits/master), I can see it's ***really*** moving forward, but there's still only one contributor pushing this like a bulldozer. 

Does anyone have any updated thoughts or feelings on this? Do you feel it's simply not at a viable point yet to receive contributions? Or are there other things preventing entry by other contributors? Is this worth investment?

Ultimately this seems like an interesting and possibly solid option to MRI and even the 3x3 update everyone's excited about - and perhaps can take Ruby even further. Yes/No?
## [7][Multiple databases connection](https://www.reddit.com/r/ruby/comments/f02evm/multiple_databases_connection/)
- url: https://www.reddit.com/r/ruby/comments/f02evm/multiple_databases_connection/
---
Hi. I have an ongoing project which asks for many DB connections (9 total) in:

-SQL Server
-MySQL
-Oracle SQL
-PostgreSQL

I never did something like this before, so i might ask for some advices or ideas.

I'm using Rails 6.

In resume, it's a data warehouse. I must delevelop an ETL module which process all the data from these 9 databases and puts into only one DB (which is PostgreSQL). I'd like to read your experiences, advices or something like this.
## [8][Arel Cheatsheet on Steroids | Including Subqueries joins, Window Functions, named functions &amp; more](https://www.reddit.com/r/ruby/comments/ezto1r/arel_cheatsheet_on_steroids_including_subqueries/)
- url: https://gist.github.com/ProGM/c6df08da14708dcc28b5ca325df37ceb
---

## [9][Ruby Project News and Blog Recommendations Welcome - What's the State of the Ruby Feed-i-verse?](https://www.reddit.com/r/ruby/comments/ezw4td/ruby_project_news_and_blog_recommendations/)
- url: https://www.reddit.com/r/ruby/comments/ezw4td/ruby_project_news_and_blog_recommendations/
---
Hello, I'm looking for good to great web feeds for Ruby Project News, Blogs, and more. Here's my current subscription list:

    News.subscribe(
      'http://www.ruby-lang.org/en/feeds/news.rss',     # Ruby Lang News
      'http://www.jruby.org/atom.xml',                  # JRuby Lang News
      'http://blog.rubygems.org/atom.xml',              # RubyGems News
      'http://bundler.io/blog/feed.xml',                # Bundler News
      'http://weblog.rubyonrails.org/feed/atom.xml',    # Ruby on Rails News
      'http://sinatrarb.com/feed.xml',                  # Sinatra News
      'https://hanamirb.org/atom.xml',                  # Hanami News
      'http://jekyllrb.com/feed.xml',                   # Jekyll News
      'http://feeds.feedburner.com/jetbrains_rubymine?format=xml',  # RubyMine IDE News
      'https://blog.phusion.nl/rss/',                   # Phusion News
      'https://rubyinstaller.org/feed.xml',             # Ruby Installer for Windows News
      'http://planetruby.github.io/calendar/feed.xml',  # Ruby Conferences &amp; Camps News
      'https://rubytogether.org/news.xml',              # Ruby Together News
      'https://foundation.travis-ci.org/feed.xml',      # Travis Foundation News
      'https://railsgirlssummerofcode.org/blog.xml',    # Rails Girls Summer of Code News
    
      'http://blog.zenspider.com/atom.xml',          # Ryan Davis @ Seattle › Washington › United States
      'http://tenderlovemaking.com/atom.xml',        # Aaron Patterson @ Seattle › Washington › United States
      'http://blog.headius.com/feed.xml',            # Charles Nutter @ Richfield › Minnesota › United States
      'http://www.schneems.com/feed.xml',            # Richard Schneeman @ Austin › Texas › United States
      'https://eregon.me/blog/feed.xml',             # Benoit Daloze @ Zürich › Switzerland
      'http://samsaffron.com/posts.rss',             # Sam Saffron @ Sydney › Australia
      )

Do you know any other good news feeds / channels that you can recommend? Please, tell. Thanks.
## [10][I made a pseudo compiler for Ruby, actually.](https://www.reddit.com/r/ruby/comments/ezw9xv/i_made_a_pseudo_compiler_for_ruby_actually/)
- url: https://www.reddit.com/r/ruby/comments/ezw9xv/i_made_a_pseudo_compiler_for_ruby_actually/
---
So yesterday I got kinda bored so I searched up how to compile Ruby, I've got these 2 answers:

Use Rubinius, or compile it with the Ruby virtual machine and then run it, always with the vm.

Today I made it, and I am satisfied about the results; it's ~2 seconds faster at doing 4000000 iterations on my Raspberry Pi 4

https://github.com/phykos/rubybc

The question is: can I go even deeper and get even better results?
