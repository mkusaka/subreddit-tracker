# ruby
## [1][Ruby adds experimental support for end-less method definition](https://www.reddit.com/r/ruby/comments/g9hlar/ruby_adds_experimental_support_for_endless_method/)
- url: https://www.reddit.com/r/ruby/comments/g9hlar/ruby_adds_experimental_support_for_endless_method/
---
[https://blog.saeloun.com/2020/04/27/ruby-adds-endless-method-definition-experimental](https://blog.saeloun.com/2020/04/27/ruby-adds-endless-method-definition-experimental)

&amp;#x200B;

PS: Do not shoot the messenger. I come in peace.
## [2][Redis connection pool in Rails](https://www.reddit.com/r/ruby/comments/g9i2oy/redis_connection_pool_in_rails/)
- url: https://tejasbubane.github.io/posts/2020-04-22-redis-connection-pool-in-rails/
---

## [3][Code walkthrough - concurrent writes, race conditions and other fun stuff when you need to rename stream with zero downtime.](https://www.reddit.com/r/ruby/comments/g9k6nx/code_walkthrough_concurrent_writes_race/)
- url: https://blog.arkency.com/rename-stream-in-rails-event-store-with-zero-downtime/
---

## [4][Rails - Added sub seconds to Time inspect](https://www.reddit.com/r/ruby/comments/g9iey4/rails_added_sub_seconds_to_time_inspect/)
- url: https://blog.saeloun.com/2020/04/27/rails-time-subseconds
---

## [5][New to programming. This code works. I want to know if I am thinking about it right.](https://www.reddit.com/r/ruby/comments/g9ibqd/new_to_programming_this_code_works_i_want_to_know/)
- url: https://www.reddit.com/r/ruby/comments/g9ibqd/new_to_programming_this_code_works_i_want_to_know/
---
https://github.com/Ubuntu19019/learnruby/blob/master/abbreviate_sentence

This is a practice problem from the app academy open website. I spent way too long solving this problem. I want to make sure my comments reflect what is actually going on. I don't think I understand line 27 and line 7.

Line 7 wasn't working, because I was doing
&gt; no_vowels = new_word

To understand line 7 I think I need to understand line 27. This is what I think, but I think it's wrong. In the array in the first function let's say we have the word "elephant" that the computer is checking. It has more than 4 letters so it goes to the second function. The second function checks each letter of elephant and adds the non vowels to new_word. Return new_word is the answer function 1 gets back from function 2. Function 1 repeats the process of sending it to function 2 until it's looped through every word. If you don't do return new_word it doesn't work, because you have all the non vowels in a string, but the program hasn't been told to send it back to function 1.
## [6][Just released sequel-activerecord-adapter gem](https://www.reddit.com/r/ruby/comments/g8z5iv/just_released_sequelactiverecordadapter_gem/)
- url: https://www.reddit.com/r/ruby/comments/g8z5iv/just_released_sequelactiverecordadapter_gem/
---
As some of you might know, I'm a big advocate for the [Sequel](https://github.com/jeremyevans/sequel) database library. However, it can be difficult to use it an app that's already using ActiveRecord, because Sequel requires its own database connection.

So I've created the [sequel-activerecord-adapter](https://github.com/janko/sequel-activerecord-adapter) gem that allows Sequel to reuse the existing ActiveRecord connection for its database interaction.

    require "sequel-activerecord-adapter"

    ActiveRecord::Base.establish_connection("postgresql:///mydb")

    DB = Sequel.activerecord # reuses the ActiveRecord connection
   
    DB.tables #=&gt; [:users, :articles, ...]

    DB[:articles].insert(title: "Sequel ActiveRecord Adapter")
    DB[:articles].all #=&gt; [{ title: "Sequel ActiveRecord Adapter" }]
    DB[:articles].update(author_id: 2)

This means you can now easily use Sequel in parts of your app that might be more performance-sensitive or require more advanced database queries that ActiveRecord doesn't support.

It also means you can use libraries like [Rodauth](https://github.com/jeremyevans/rodauth) that use Sequel under-the-hood without the performance impact and mental overhead that would come with dealing with separate database connections.
## [7][IDE for JRuby?](https://www.reddit.com/r/ruby/comments/g95jl0/ide_for_jruby/)
- url: https://www.reddit.com/r/ruby/comments/g95jl0/ide_for_jruby/
---
I'd like to try JRuby for development for the JVM. It looks pretty nice, but I'm used to use VSCode for all of my development stuff and VSCode doesn't seem to have a good Ruby plugin. Is there a good and free IDE for JRuby development I could use? Obvioisly I know about RubyMine, but I'd prefer some free and not java based IDE. Thanks!
## [8][Dig](https://www.reddit.com/r/ruby/comments/g92d0v/dig/)
- url: https://www.reddit.com/r/ruby/comments/g92d0v/dig/
---
I'm working on project to enhance the rdoc for Ruby.  It's gradually being merged into the Ruby project.

Recently I was surprised by my research into dig.  Some of you may be, too.

\- My proposed rdoc:  [https://github.com/BurdetteLamar/AboutRuby/blob/master/core/Hash/api/markdown.md#dig](https://github.com/BurdetteLamar/AboutRuby/blob/master/core/Hash/api/markdown.md#dig)

\- The existing rdoc:  [https://ruby-doc.org/core-2.7.0/Hash.html#method-i-dig](https://ruby-doc.org/core-2.7.0/Hash.html#method-i-dig)
## [9][Idiomatic way to check if object responds to method, otherwise return blank string](https://www.reddit.com/r/ruby/comments/g8zydx/idiomatic_way_to_check_if_object_responds_to/)
- url: https://www.reddit.com/r/ruby/comments/g8zydx/idiomatic_way_to_check_if_object_responds_to/
---
I have a web scraper that's parses rows, and sometimes it cannot find elements and throws an exception.  I've added a check to see if the query returns nill object, and if so return a blank string instead.

I was curious if there's a better way to do this in general or more idiomatic way in Ruby to write this.

&amp;#x200B;

    def parse_single_post(row)
    title     =  row.at_css('.topictitle')
    replies   =  row.at_css('.replies')
    views     =  row.at_css('.posts')
    last_post =  row.at_css('.lastpost')
    author    =  row.at_css('.username')
    link      =  row.at_css('a.topictitle')

    posts &lt;&lt; OpenStruct.new(
      title:     title ? title.text : "",
      link:      link ? link['href'] : "",
      author:    author ? author.text : "",
      last_post: last_post ? last_post.text : "",
      replies:   replies ? replies.text : "",
      views:     views ? views.text : ""
    )
 end
## [10][LOOKING FOR A MENTOR](https://www.reddit.com/r/ruby/comments/g91gjq/looking_for_a_mentor/)
- url: https://www.reddit.com/r/ruby/comments/g91gjq/looking_for_a_mentor/
---
Hey ! just like the title states , I’m looking for a ruby/rails mentor for the next month or so until I get hired. I’ve taught myself ruby and rails but I know I need a little help gaining a better understanding of advanced concepts. Unfortunately code-mentor and similar sites charge up the arse and I really can’t afford it just yet. I’m extremely hard working and persistent, loyal and dedicated. If anyone could spare some time I would be forever grateful.
