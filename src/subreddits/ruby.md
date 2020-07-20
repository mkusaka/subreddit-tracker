# ruby
## [1][Roadmap Ruby and Ruby On Rails](https://www.reddit.com/r/ruby/comments/huf714/roadmap_ruby_and_ruby_on_rails/)
- url: https://www.reddit.com/r/ruby/comments/huf714/roadmap_ruby_and_ruby_on_rails/
---
You know of a roadmap (similar [https://roadmap.sh/roadmaps](https://roadmap.sh/roadmaps) [https://www.quora.com/What-is-Ruby-On-Rails-framework-learning-road-map](https://www.quora.com/What-is-Ruby-On-Rails-framework-learning-road-map)) but for Ruby or Ruby On Rails, I'd like to create one for Ruby and another for Ruby On Rails, but first I'd like to know if you know of one that already exists and thus extend it or develop it as a learning path that can help novice programmers in their learning. Thanks for your input!
## [2][Interesting or cool ruby features?](https://www.reddit.com/r/ruby/comments/huip5j/interesting_or_cool_ruby_features/)
- url: https://www.reddit.com/r/ruby/comments/huip5j/interesting_or_cool_ruby_features/
---
If you were to demo or explain what Ruby is to a bunch of non-ruby devs (e.g. Java or C# developers) what would you demonstrate or explain about Ruby which you think other devs might find interesting?  Any specific methods or facts that might impress a developer with little or no ruby knowledge?

Perhaps there is already a good website with some examples?

Thanks
## [3][Making RSpec Tests More Robust](https://www.reddit.com/r/ruby/comments/hukup6/making_rspec_tests_more_robust/)
- url: http://jakeyesbeck.com/2020/07/19/making-rspec-tests-more-robust/?src=reddit
---

## [4][Ruby (Squib) error when running](https://www.reddit.com/r/ruby/comments/huc7bh/ruby_squib_error_when_running/)
- url: https://www.reddit.com/r/ruby/comments/huc7bh/ruby_squib_error_when_running/
---
I was able to get this running yesterday, but it started throwing errors even when I reverted back to an earlier draft of my Ruby file. Now it still throws the same error when I'm trying to just use a verbatim copy-and-paste (that worked in the past) example from the documentation.

[https://squib.rocks/](https://squib.rocks/)

This is the error it throws on the sample from the above link (on the bottom of the page, which I've been able to run before).

I'm running this through WSL (Ubuntu).

    $ ruby deck.rb
    WARNING: Nokogiri was built against LibXML version 2.9.9, but has dynamically loaded 2.9.10  
    Traceback (most recent call last):  
    9: from deck.rb:3:in '&lt;main&gt;'  
    8: from deck.rb:3:in 'new'  
    7: from /var/lib/gems/2.7.0/gems/squib-0.15.3/lib/squib/deck.rb:74:in 'initialize'  
    6: from /var/lib/gems/2.7.0/gems/squib-0.15.3/lib/squib/deck.rb:74:in 'instance_eval'  
    5: from deck.rb:5:in 'block in &lt;main&gt;'  
    4: from /var/lib/gems/2.7.0/gems/squib-0.15.3/lib/squib/api/data.rb:124:in 'xlsx'  
    3: from /var/lib/gems/2.7.0/gems/squib-0.15.3/lib/squib/api/data.rb:18:in 'xlsx'  
    2: from /var/lib/gems/2.7.0/gems/squib-0.15.3/lib/squib/api/data.rb:18:in 'upto'  
    1: from /var/lib/gems/2.7.0/gems/squib-0.15.3/lib/squib/api/data.rb:20:in 'block in xlsx' /var/lib/gems/2.7.0/gems/squib-0.15.3/lib/squib/api/data.rb:20:in 'strip!': can't modify frozen String: "" (FrozenError)  

&amp;#x200B;
## [5][Request for code review](https://www.reddit.com/r/ruby/comments/hu5xxd/request_for_code_review/)
- url: https://www.reddit.com/r/ruby/comments/hu5xxd/request_for_code_review/
---
Hello everyone!

I have started learning ruby and I've just finished my first OOP project (I am learning through the odin project). I want to know what kind of mistakes  am I making early on. So I would appreciate any reviews/critiques/mistake identification.

[https://github.com/e-km0d/Ruby-Projects/tree/master/Tic\_Tac\_Toe](https://github.com/e-km0d/Ruby-Projects/tree/master/Tic_Tac_Toe)

What can I improve here?

Thanks!
## [6][How we migrated application servers from Unicorn to Puma (GitLab)](https://www.reddit.com/r/ruby/comments/htw2iy/how_we_migrated_application_servers_from_unicorn/)
- url: https://about.gitlab.com/blog/2020/07/08/migrating-to-puma-on-gitlab/
---

## [7][RFC: Naming complex getter/setter methods](https://www.reddit.com/r/ruby/comments/hu22bk/rfc_naming_complex_gettersetter_methods/)
- url: https://www.reddit.com/r/ruby/comments/hu22bk/rfc_naming_complex_gettersetter_methods/
---
I'm currently re-evaluating how to name methods that are, to the user, getter and setter methods but are more complex in their implementation than just setting a variable.

If just a variable needs to be set, we have the `attr_reader`, `attr_writer` and `attr_accessor` methods. That's easy and okay for simple cases.

In more complex cases where, for example, multiple arguments are required the `attr_*` methods won't work because only one value can be given as argument.

The way I'm currently doing this in most cases is to have a single method like `#line_dash_pattern(value = nil, phase = 0, &amp;block)` which acts as getter if `value` is `nil` or as a setter if `value` isn't.  The definition for this method can be found [here](https://github.com/gettalong/hexapdf/blob/master/lib/hexapdf/content/canvas.rb#L615) and its implementation [here](https://github.com/gettalong/hexapdf/blob/master/lib/hexapdf/content/canvas.rb#L2027).

I find it a bit ugly due to big `if/else` statement but haven't found anything better. For example, using methods named `get_line_dash_pattern` and `set_line_dash_pattern` feels a bit un-Ruby like, though it allows for consistent naming, even if only getter or setter is needed.

Since, I guess, this "problem" is something that many have encountered, I would be interested in hearing how you solved it! Thanks!
## [8][TCP/UDP explanation with ruby (using log delivery as an example)](https://www.reddit.com/r/ruby/comments/htxyhs/tcpudp_explanation_with_ruby_using_log_delivery/)
- url: https://www.reddit.com/r/ruby/comments/htxyhs/tcpudp_explanation_with_ruby_using_log_delivery/
---
Wrote an article about log delivery and in this article, I'll explain TCP/UDP using ruby. Hope this will help you understand TCP/UDP better

[https://dev.sweatco.in/centralized-logging-delivery/](https://dev.sweatco.in/centralized-logging-delivery/)
## [9][Using async/await syntax in ruby with async-await gem](https://www.reddit.com/r/ruby/comments/htpwuf/using_asyncawait_syntax_in_ruby_with_asyncawait/)
- url: https://www.reddit.com/r/ruby/comments/htpwuf/using_asyncawait_syntax_in_ruby_with_asyncawait/
---
Hello! My question is about such a popular concept in other programming languages - async/await. Have you tried out a gem [async-await](https://github.com/socketry/async-await) which is built on top of [async](https://github.com/socketry/async) that gives us async decorator for functions and await method (accepting a block)?

[code example](https://imgur.com/I3kACSF)

Please share your thoughts/experience about it.
## [10][How can I create class instance and run a class method when .txt file gets added to the directory in Ruby?](https://www.reddit.com/r/ruby/comments/htz846/how_can_i_create_class_instance_and_run_a_class/)
- url: https://www.reddit.com/r/ruby/comments/htz846/how_can_i_create_class_instance_and_run_a_class/
---

