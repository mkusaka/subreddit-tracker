# ruby
## [1][Say hello to RuboCop::Packaging! 👋](https://www.reddit.com/r/ruby/comments/i87l2l/say_hello_to_rubocoppackaging/)
- url: https://www.reddit.com/r/ruby/comments/i87l2l/say_hello_to_rubocoppackaging/
---
Hello all,

We (the Debian core devs (DDs)) have been facing a lot of problems whilst maintaining Ruby libraries and applications in Ruby. We maintain around 1600 packages in the Debian archive.  
Some of the problems, for example, are using a relative path for `require_relative` calls from specs (tests directory) to the lib directory (`require_relative '../lib/foo'`) and others include using `git` in the gemspec files. And so on..

So we came up with the idea of writing a linter (an extension to RuboCop) for flagging those calls which can be fixed upstream (so a win-win situation for both, Downstream and Upstream!).  
And this is how RuboCop::Packaging (https://github.com/utkarsh2102/rubocop-packaging) is born! 💖

Please give it a look, let me know what you think about it?  
In case you find it good, please give it a ⭐, so the Downstream maintainers know that it was worth it!    
Suggestions and opinions are heartily welcomed! 🤗
## [2][Learn Ruby Regexp with hundreds of examples and exercises](https://www.reddit.com/r/ruby/comments/i85hm1/learn_ruby_regexp_with_hundreds_of_examples_and/)
- url: https://www.reddit.com/r/ruby/comments/i85hm1/learn_ruby_regexp_with_hundreds_of_examples_and/
---
Hello!

I updated my Ruby Regexp book recently, changes include updating to 2.7.1, adding epub/html version, new sections on escape sequences, conditional grouping, `\R`, etc. There's also significant increase in number of exercises and this time I added solutions for reference as well. The entire book can be read online at

https://learnbyexample.github.io/Ruby_Regexp/

Or, you can get free pdf/epub versions from:

* https://gumroad.com/l/rubyregexp
* https://leanpub.com/rubyregexp

I made **all my books free** at the end of March when the pandemic hit my country. And then I spent more than four months to update all my books, this is the last one. You can get all my books as a bundle (Python/Ruby/JS regex and grep/sed/awk cli tools) for free until this weekend:

* https://gumroad.com/l/regex
* https://leanpub.com/b/regex

I’d highly appreciate your feedback. That’s been a major motivating factor to keep writing as well as for improving the content.

Happy learning :)
## [3][TestBench now enables projects to support MRuby and Ruby at the same time seamlessly](https://www.reddit.com/r/ruby/comments/i7yktf/testbench_now_enables_projects_to_support_mruby/)
- url: https://twitter.com/realntl/status/1293271193634713601
---

## [4][JRuby 9.2.13.0 Released](https://www.reddit.com/r/ruby/comments/i82f55/jruby_92130_released/)
- url: https://www.jruby.org/2020/08/03/jruby-9-2-13-0.html
---

## [5][How Sidekiq really works](https://www.reddit.com/r/ruby/comments/i7qj7b/how_sidekiq_really_works/)
- url: https://pdabrowski.com/articles/how-sidekiq-really-works
---

## [6][Speed up Rails boot by 2.5x with a single line change if you use Docker on Mac](https://www.reddit.com/r/ruby/comments/i7skbj/speed_up_rails_boot_by_25x_with_a_single_line/)
- url: https://www.reddit.com/r/ruby/comments/i7skbj/speed_up_rails_boot_by_25x_with_a_single_line/
---
    # before:
    services:
      web:
        volumes:
          - .:/app
    
    # after, notice "cached":
    services:
      web:
        volumes:
          - .:/app:cached

And then re-create container with: 

    docker-compose up web

I've covered some background on this [here](https://levelup.gitconnected.com/speed-up-rails-on-docker-on-mac-2c9871c33f63?source=friends_link&amp;sk=ef3e03f93b0a149ddace1022c0565f98), but it's essentially that simple.
## [7][Creating unlogged PostgreSQL tables in Rails](https://www.reddit.com/r/ruby/comments/i7p96x/creating_unlogged_postgresql_tables_in_rails/)
- url: https://prathamesh.tech/2020/08/10/creating-unlogged-tables-in-rails/
---

## [8][Looping over an array with Ruby: Each, For Loops, While Loops, Do While Loops, Until Loops + Breaks](https://www.reddit.com/r/ruby/comments/i7sydn/looping_over_an_array_with_ruby_each_for_loops/)
- url: https://youtu.be/-7iy-o0MRS0
---

## [9][How to use Bootstrap and jQuery in Rails 6 with Webpacker](https://www.reddit.com/r/ruby/comments/i7sy06/how_to_use_bootstrap_and_jquery_in_rails_6_with/)
- url: https://rubyyagi.com/how-to-use-bootstrap-and-jquery-in-rails-6-with-webpacker/
---

## [10][When objects become super objects](https://www.reddit.com/r/ruby/comments/i7lfsd/when_objects_become_super_objects/)
- url: https://www.reddit.com/r/ruby/comments/i7lfsd/when_objects_become_super_objects/
---
I recently gave some thought into how I've _broken_ the single responsibiltiy principle throughout my career as a software engineer and came up with an interesting analogy using a pencil. [https://ramallo.pw/ruby/2020/08/11/when-objects-become-super-objects](https://ramallo.pw/ruby/2020/08/11/when-objects-become-super-objects)
