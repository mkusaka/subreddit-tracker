# ruby
## [1][Benchmarking Ruby 2.7.0's Numbered Parameters](https://www.reddit.com/r/ruby/comments/fat90p/benchmarking_ruby_270s_numbered_parameters/)
- url: https://blog.schembri.me/post/benchmarking-ruby-2.7.0-numbered-parameters/
---

## [2][Faster Excel Parsing in Ruby](https://www.reddit.com/r/ruby/comments/faszl8/faster_excel_parsing_in_ruby/)
- url: https://blog.schembri.me/post/faster-excel-parsing-in-ruby/
---

## [3][Has anyone used Kiba?](https://www.reddit.com/r/ruby/comments/fal2l2/has_anyone_used_kiba/)
- url: https://www.reddit.com/r/ruby/comments/fal2l2/has_anyone_used_kiba/
---
Hey! I've been looking for data engineering tools, found tons of them, to manage workflows, to generate ETL jobs, most of them use Python and I'd rather use Ruby. (i.e. amazon glue, airflow, astronomer, stitchdata, ...)

A few weeks ago, the Ruby Weekly featured Kiba v3 https://github.com/thbar/kiba 

I've been reading thorough their wiki and I feel very inclined to start using it, although I would also need to set up a server to run these ETL tasks, maybe a serverless instance or something similar. 

Just to give you some broad context, my main goal is to extract data from a tsv file, transform the data into a set of attributes I want, and lastly store the transformed data into a csv file so that later I can easily load it into redshift.

My question is, do you have any experiences to share regarding data engineering in Ruby? Or have you used Kiba previously? Any insights or comments? 

Anything would be appreciated!
## [4][I get nil for sort_by on an array of objects](https://www.reddit.com/r/ruby/comments/far4kd/i_get_nil_for_sort_by_on_an_array_of_objects/)
- url: https://www.reddit.com/r/ruby/comments/far4kd/i_get_nil_for_sort_by_on_an_array_of_objects/
---
    def sorted_objects
          @objects.sort_by do |l, r|
            if !l || !r
              @sort_results = 0
            else
              @sort_results = 1
            end
          end.to_a
          #p sort_results
         # p @objects
        end

so I have printed r and it shows nil every time

here's an example of what objects is storing  


    @objects &lt;&lt; Room.new(10 - x, y, tile.to_i, 220, 300, @max_x.to_i, @max_y.to_i,"floor",0)

no idea why r returns nil.
## [5][The state of ruby blogs and news (2020/2 Edition) - 36 channels, 1464 items](https://www.reddit.com/r/ruby/comments/fahwbr/the_state_of_ruby_blogs_and_news_20202_edition_36/)
- url: https://www.reddit.com/r/ruby/comments/fahwbr/the_state_of_ruby_blogs_and_news_20202_edition_36/
---
Hello,

   I've put together a little survey about the state of the ruby feed-iverse that includes personal blogs, ruby project news and more.

  See the [planet.ini](https://github.com/planetruby/planet/blob/master/planet.ini) for all feeds included in the survey.

Q: What feed formats are in use?

    Formats  (n=36)
    ---------------------------------
      atom        23 (63%) | *************************************
      rss 2.0     13 (36%) | *********************


Q: What servers are in use?

    Servers  (n=36)
    ---------------------------------
      GitHub.com  12 (33%) | *******************
      nginx        8 (22%) | *************
      Cowboy       5 (13%) | *******
      Apache       4 (11%) | ******
      cloudflare   3 ( 8%) | ****
      GSE          2 ( 5%) | ***
      Netlify      2 ( 5%) | ***


Q: What (web site) publishing tools are in use?

    Generators  (n=36)
    ---------------------------------
      ?           22 (61%) | ************************************
      Jekyll       9 (25%) | ***************
      wordpress    3 ( 8%) | ****
      webgen       1 ( 2%) | *
      Ghost        1 ( 2%) | *


Q: What's the update frequency of posts?

    Update Frequency  (n=35)
    ---------------------------------
      &lt;=   7 days   3 ( 8%) | ****
      &lt;=  14 days   4 (11%) | ******
      &lt;=  30 days   7 (20%) | ************
      &lt;=  90 days  12 (34%) | ********************
      &lt;= 180 days   5 (14%) | ********
      &lt;= 365 days   4 (11%) | ******

That's all for now. Comments and ideas on how to improve
the stats are more than welcome. Cheers. Prost.
## [6][Ruby Hash#transform_keys now accepts a hash that maps existing keys to new keys](https://www.reddit.com/r/ruby/comments/fae2jf/ruby_hashtransform_keys_now_accepts_a_hash_that/)
- url: https://blog.saeloun.com/2020/02/26/ruby-hash-transform_keys-now-accepts-a-hash-that-maps-existing-keys-to-new-keys
---

## [7][Building a reddit bot using Ruby](https://www.reddit.com/r/ruby/comments/fam6w7/building_a_reddit_bot_using_ruby/)
- url: https://www.reddit.com/r/ruby/comments/fam6w7/building_a_reddit_bot_using_ruby/
---
I apologize if this has been asked recently, I had trouble finding anything semi-current on the subject. If there's recent threads on this that I didn't find, please let me know as I'd love to see them.

I'm a newish Ruby developer and I have been toying with the idea of making a simple reddit bot with Ruby, mainly to help me learn various skills, but also to be functional and serve a sub I mod. I want it to fetch info from one outside site and post the data it finds there when requested. Simple enough I hope (not simple for me lol). 

I'm having trouble finding much info on how to implement a reddit bot using Ruby. I know Python is what's used most often, and I do have some experience with it but I'm focusing on Ruby atm and I don't want to go off on a Python tangent. I found the \[Reddit\_Bot\]([https://rubygems.org/gems/reddit\_bot](https://rubygems.org/gems/reddit_bot)) gem which seems to be popular and recently updated. I'm going through the docs which are helpful and have a lot of examples, but there's basic info missing that is probably common knowledge to more experienced devs but a vast abyssal mystery to newbies like me. 

So to be specific:

1. Is Ruby okay to use for this, or should I set aside cash now for when I inevitably smash my computer in frustration?
2. If (1) is yes, are there any videos, tutorials, anything you know of that are Ruby-specific that might be helpful for this project? (I did already watch several youtube vids on Python bots which were helpful, but since I'm still a baby dev I have a hard time translating what I learn there to Ruby context) 
3. Any advice, tips, or whatever you have for this would be appreciated. I have mainly worked in a Rails framework for general web development and have only a little experience working with API's, so pretty much every part of this process is going to be a learning experience for me.

Thanks!
## [8][Telegram Bot wrapper for the RubyGems.org API](https://www.reddit.com/r/ruby/comments/fadr8d/telegram_bot_wrapper_for_the_rubygemsorg_api/)
- url: https://github.com/davidesantangelo/rubyg-bot
---

## [9][Anyone used Sled with Ruby?](https://www.reddit.com/r/ruby/comments/faeebb/anyone_used_sled_with_ruby/)
- url: https://www.reddit.com/r/ruby/comments/faeebb/anyone_used_sled_with_ruby/
---
Has anyone used [Sled](https://github.com/spacejam/sled) (embedded DB) with Ruby so far? Was curious to hear about any experiences or observations so far... 

I was particularly wondering about how it might with RubyMotion.
## [10][Introducing Hanami::API](https://www.reddit.com/r/ruby/comments/f9tmgv/introducing_hanamiapi/)
- url: https://www.reddit.com/r/ruby/comments/f9tmgv/introducing_hanamiapi/
---
It's a minimal, extremely fast, lightweight Ruby framework for HTTP APIs.

(The article contains benchmark results)

[http://hanamirb.org/blog/2020/02/26/introducing-hanami-api.html](http://hanamirb.org/blog/2020/02/26/introducing-hanami-api.html)
