# ruby
## [1][Rewrite Kernel#tap with Ruby by k0kubun](https://www.reddit.com/r/ruby/comments/hker0s/rewrite_kerneltap_with_ruby_by_k0kubun/)
- url: https://github.com/ruby/ruby/pull/3281
---

## [2][Fun Facts about Ruby #8](https://www.reddit.com/r/ruby/comments/hkf8zk/fun_facts_about_ruby_8/)
- url: https://twitter.com/RubyCademy/status/1278962830637662208
---

## [3][I thought I understood "self", but I don't](https://www.reddit.com/r/ruby/comments/hk6upv/i_thought_i_understood_self_but_i_dont/)
- url: https://www.reddit.com/r/ruby/comments/hk6upv/i_thought_i_understood_self_but_i_dont/
---
https://github.com/Ubuntu19019/learnruby/blob/master/Understanding%20self

So I'm doing a battleship project. It has an rspec file and I'm doing what it tells me to do. The attack method (which I believe is an instance method)  is tripping me up. 

I thought self is used for an instance of the class Board. new_board = Board.new(n).  new_board would be an instance. There's no function to return @grid. But attr_reader for @size. So I could do self.size.(some other instance method) right? Since I can't do that with @grid I thought I would have to do 
&gt;if @grid[position] == :S

That doesn't work. My theory is that since there's no method for returning @grid I can't call another method on it. If there was a method I could call either self.grid[position] == :S or @grid[position] == :S

If I had a method for returning @grid is it better to directly call @grid or do self.grid(is this an instance)?

I also need clarification on what self refers to. 

&gt;if self[position] == :S

Is self referring to an instance of the class? If that's the case what exactly does the initialize method do? Does initialize build the grid inside of the Board class, then self[postion] would be like new_board[position] since new_board is an instance of the Board class?

I'm really confused. After I digest this information I want to know more about class methods vs instance methods and why I would choose to make one over the other.
## [4][How often do you write scrapping scripts?](https://www.reddit.com/r/ruby/comments/hk8glk/how_often_do_you_write_scrapping_scripts/)
- url: https://www.reddit.com/r/ruby/comments/hk8glk/how_often_do_you_write_scrapping_scripts/
---
I was doing some website and the client needed me to copy some data from a blog. So naturally I said, this would be a good time to write a webscrapper. It took me a while to get the hang of Nokogiri but I did. Which made me wonder, how often do you write scrapping scripts in your experience?
## [5][Trailblazer version 2.1 - official launch with new website and documentation.](https://www.reddit.com/r/ruby/comments/hjyybu/trailblazer_version_21_official_launch_with_new/)
- url: https://www.reddit.com/r/ruby/comments/hjyybu/trailblazer_version_21_official_launch_with_new/
---
For those who know trailblazer - you know that TRB evolved few times, introduced a lot of great ideas but also had his problems:  missing documentation, lack of communication with the community, and not consistent approach to some problems. We strongly believe that the current version with the support of core team and the new landing page with complete documentation is something that is worth working with.  


For those who somehow didn't hear about Trailblazer - check it out, probably for some complex projects where Rails aren't sufficient that tool could be really helpful and joyfull.  


Either you know TRB already or you never heard about it, check out the blog post about the history and changes of trailblazer 2.1 : [https://trailblazer.to/2.1/blog.html](https://trailblazer.to/2.1/blog.html) and give us feedback!
## [6][Can you Scrap images with Nokogiri? Or do I have to go the Selenium route?](https://www.reddit.com/r/ruby/comments/hk8r1b/can_you_scrap_images_with_nokogiri_or_do_i_have/)
- url: https://www.reddit.com/r/ruby/comments/hk8r1b/can_you_scrap_images_with_nokogiri_or_do_i_have/
---
As the title says, I'm looking to get some images from a blog. Any help would be appreciated
## [7][Gem with complex parameters](https://www.reddit.com/r/ruby/comments/hjs5x7/gem_with_complex_parameters/)
- url: https://www.reddit.com/r/ruby/comments/hjs5x7/gem_with_complex_parameters/
---
Hi there.

Imagine that I'm creating a gem that takes care of some heavy and complex calculations.

    def compute_stuff(...)

One of the inputs that gem needs is also non trivial. The gem needs to receive a collection of objects, each one of those objects would have a few fields itself. Something like this:

    [
        {"name": "John", "location": "US", "order": {...} },
        {"name": "Jane", "location": "UK", "order": {...} }
        ...
    ]

Which strategy would you user to pass that kind of inputs to the gem's method?

1. A simple hash: just send an hash to the gem
2. Create a class mapping each object and "make" the caller use that class?
3. Other option?

&amp;#x200B;

I have a past writing c# code and I really like interfaces to be as explicit as possible, how can we try to achieve that for a gem like this one?
## [8][Release JRuby 9.2.12.0 Released · jruby/jruby · GitHub](https://www.reddit.com/r/ruby/comments/hjh3vu/release_jruby_92120_released_jrubyjruby_github/)
- url: https://github.com/jruby/jruby/releases/tag/9.2.12.0
---

## [9][AnyCable 1.0: Four years of real-time web with Ruby and Go](https://www.reddit.com/r/ruby/comments/hjb2td/anycable_10_four_years_of_realtime_web_with_ruby/)
- url: https://evilmartians.com/chronicles/anycable-1-0-four-years-of-real-time-web-with-ruby-and-go
---

## [10][digest-crc 0.6.0 released. ~40x performance improvement!](https://www.reddit.com/r/ruby/comments/hjndvt/digestcrc_060_released_40x_performance_improvement/)
- url: https://www.reddit.com/r/ruby/comments/hjndvt/digestcrc_060_released_40x_performance_improvement/
---
[digest-crc](https://github.com/postmodern/digest-crc#readme) [0.6.0](https://github.com/postmodern/digest-crc/blob/master/ChangeLog.md#060--2020-07-01) has been released. This release introduces _optional_ C extensions, that when built on CRuby will override the pure-Ruby CRC methods with C equivalents, providing a ~40x performance improvement. If the C extensions cannot be built (no compiler installed, no `ruby.h` headers, not using CRuby) digest-crc will fallback to using the pure-Ruby CRC methods.

**Note:** If you use [ruby-kafka](https://github.com/zendesk/ruby-kafka#readme), run `bundle update digest-crc` to take advantage of the improved performance (see: https://github.com/zendesk/ruby-kafka/issues/620).
