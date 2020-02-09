# ruby
## [1][Rails adds support for if_exists/if_not_exists on remove_column/add_column in migrations](https://www.reddit.com/r/ruby/comments/f176xt/rails_adds_support_for_if_existsif_not_exists_on/)
- url: https://blog.saeloun.com/2020/02/04/rails-support-for-if_exists-if_not_exists-on-remove_column-add_column-in-migrations
---

## [2][How to source external files in Jekyll kramdown?](https://www.reddit.com/r/ruby/comments/f1310m/how_to_source_external_files_in_jekyll_kramdown/)
- url: https://www.reddit.com/r/ruby/comments/f1310m/how_to_source_external_files_in_jekyll_kramdown/
---
What is the best way to source external files into the kramdown of my Jekyll site post? I am creating code tutorials and would like the code snippets to reside in an external git repo in script form.
## [3][Is there anyway to build a parser in Ruby](https://www.reddit.com/r/ruby/comments/f0vthc/is_there_anyway_to_build_a_parser_in_ruby/)
- url: https://www.reddit.com/r/ruby/comments/f0vthc/is_there_anyway_to_build_a_parser_in_ruby/
---
Hi guys, Ruby Noob on here. 

I was thinking of building a basic browser, and I was told that at the very least, I would need a HTML parser and a way to get HTTP request. 

So my question is, can I build a parser for a markup like HTML in ruby? And if so, what are the constraints and factors I should consider?
## [4][Backports 3.16 released: Support of Ruby 2.7's features in earlier Ruby versions](https://www.reddit.com/r/ruby/comments/f0erf3/backports_316_released_support_of_ruby_27s/)
- url: https://github.com/marcandre/backports/blob/master/CHANGELOG.rdoc#version-3160---feb-6th-2020
---

## [5][Setting up your next Rails project with GitlabCI - a 101 guide](https://www.reddit.com/r/ruby/comments/f0ahyv/setting_up_your_next_rails_project_with_gitlabci/)
- url: https://hixonrails.com/ruby-on-rails-tutorials/ruby-on-rails-set-up-on-gitlab-with-gitlabci/
---

## [6][a bit off-top](https://www.reddit.com/r/ruby/comments/f0qd71/a_bit_offtop/)
- url: https://www.reddit.com/r/ruby/comments/f0qd71/a_bit_offtop/
---
What theme is this?

https://preview.redd.it/rhjadzngpof41.png?width=664&amp;format=png&amp;auto=webp&amp;s=2c6d88a72ff71b858a35f9cf872246540a890071
## [7][Question on ruby and websockets](https://www.reddit.com/r/ruby/comments/f0evse/question_on_ruby_and_websockets/)
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
## [8][Ruby Quiz - Challenge #17 - Build an HTML Template Engine Like It's 1999](https://www.reddit.com/r/ruby/comments/f0f3kj/ruby_quiz_challenge_17_build_an_html_template/)
- url: https://github.com/planetruby/quiz/tree/master/017
---

## [9][So what have your thoughts been on Artichoke Ruby?](https://www.reddit.com/r/ruby/comments/ezx6t8/so_what_have_your_thoughts_been_on_artichoke_ruby/)
- url: https://www.reddit.com/r/ruby/comments/ezx6t8/so_what_have_your_thoughts_been_on_artichoke_ruby/
---
I saw [Artichoke Ruby](https://github.com/artichoke/artichoke) mentioned [several](https://www.reddit.com/r/ruby/comments/clu9pd/artichoke_a_ruby_made_with_rust/) months ago and it seemed like a lot of folks were excited about this at the time. Fast forward to now, looking at the Github [commits](https://github.com/artichoke/artichoke/commits/master), I can see it's ***really*** moving forward, but there's still only one contributor pushing this like a bulldozer. 

Does anyone have any updated thoughts or feelings on this? Do you feel it's simply not at a viable point yet to receive contributions? Or are there other things preventing entry by other contributors? Is this worth investment?

Ultimately this seems like an interesting and possibly solid option to MRI and even the 3x3 update everyone's excited about - and perhaps can take Ruby even further. Yes/No?
## [10][Multiple databases connection](https://www.reddit.com/r/ruby/comments/f02evm/multiple_databases_connection/)
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
