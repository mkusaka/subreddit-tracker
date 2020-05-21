# ruby
## [1][A prototype of an interactive tutorial environment for building games with Ruby (link in the comments).](https://www.reddit.com/r/ruby/comments/gnpvp8/a_prototype_of_an_interactive_tutorial/)
- url: https://v.redd.it/w3w1mpldl1051
---

## [2][5 Ways to Splat in Ruby](https://www.reddit.com/r/ruby/comments/gnve20/5_ways_to_splat_in_ruby/)
- url: https://hint.io/blog/5-ways-to-splat-in-ruby
---

## [3][Why did Why the Lucky Stiff quit?](https://www.reddit.com/r/ruby/comments/gnq0aq/why_did_why_the_lucky_stiff_quit/)
- url: https://www.reddit.com/r/ruby/comments/gnq0aq/why_did_why_the_lucky_stiff_quit/
---

## [4][The new gem that helps you rule all your classes :)](https://www.reddit.com/r/ruby/comments/gnvy31/the_new_gem_that_helps_you_rule_all_your_classes/)
- url: https://www.reddit.com/r/ruby/comments/gnvy31/the_new_gem_that_helps_you_rule_all_your_classes/
---
Github: [https://github.com/andriy-baran/mature\_factory](https://github.com/andriy-baran/mature_factory) 

Gitbook docs:  [https://app.gitbook.com/@andriy-baran-v/s/mature-factory/](https://app.gitbook.com/@andriy-baran-v/s/mature-factory/)
## [5][Speeding up Rails with Memoization](https://www.reddit.com/r/ruby/comments/gnl9xi/speeding_up_rails_with_memoization/)
- url: https://www.reddit.com/r/ruby/comments/gnl9xi/speeding_up_rails_with_memoization/
---
Whoever first said that "the fastest code is no code" must have really liked memoization. After all, memoization speeds up your application by running less code. In this article, Jonathan Miles introduces us to memoization. We'll learn when to use it, how to implement it in Ruby, and how to avoid common pitfalls. Buckle up! [https://www.honeybadger.io/blog/ruby-rails-memoization/](https://www.honeybadger.io/blog/ruby-rails-memoization/)
## [6][New OData server library for Ruby](https://www.reddit.com/r/ruby/comments/gntthv/new_odata_server_library_for_ruby/)
- url: https://www.reddit.com/r/ruby/comments/gntthv/new_odata_server_library_for_ruby/
---
Hello rubyists,

Although OData is around for a long time ago now, and some already predicted it's dead, i still believe there is value in it. The OData Protocoll has been open sourced, but when i looked for open source server implemententions i found only python, php and javascript implementations, but nothing in Ruby. I tried to fill this gap with the [Safrano library](https://gitlab.com/dm0da/safrano#safrano) .

As an example of what can be done with it, you can have a look at the [Fiveapples demo app](https://gitlab.com/dm0da/fiveapples#fiveapples) which runs safrano and an openui5 demo app in a single standalone Rack application.

If you know  a bit of database modeling and are looking for REST/OData server, this might be for you.

&amp;#x200B;
## [7][How to Migrate from Capybara Webkit to Webdrivers](https://www.reddit.com/r/ruby/comments/gndimm/how_to_migrate_from_capybara_webkit_to_webdrivers/)
- url: https://www.fastruby.io/blog/rails/upgrades/how-to-migrate-from-capybara-webkit-to-webdrivers.html
---

## [8][anyone used `sdbm` from stdlib? Anyone found it buggy?](https://www.reddit.com/r/ruby/comments/gneh4r/anyone_used_sdbm_from_stdlib_anyone_found_it_buggy/)
- url: https://www.reddit.com/r/ruby/comments/gneh4r/anyone_used_sdbm_from_stdlib_anyone_found_it_buggy/
---
So the ruby stdlib actually has a few different file-based key/value store implementations. 

There's [PStore](https://ruby-doc.org/stdlib-2.5.3/libdoc/pstore/rdoc/PStore.html), which gets talked about maybe the most. There's [gdbm](https://ruby-doc.org/stdlib-2.5.3/libdoc/gdbm/rdoc/GDBM.html) and [dbm](https://ruby-doc.org/stdlib-2.6.1/libdoc/dbm/rdoc/DBM.html) -- which weren't available on my system. Apparently while kind of in stdlib, ruby has to be compiled on a system that has dbm or gdbm library headers available for the ruby packages to actually be on the system. 

And there's also [SDBM](https://ruby-doc.org/stdlib-2.6.3/libdoc/sdbm/rdoc/SDBM.html). Which was available on my system, and for some reason I tried to use (under MRI 2.6.6)

I *think* I ran into some severe bugs. I was storing around 300,000 key/values in there. The keys are all URLs (and fairly similar to URLs to each other), the values are all "true" -- I was basically trying to use it as a simple persistent quick-lookup Set. 

I found that some things I stored there `store[key] = value`, would NOT come back with the value when retrieved `store[key] # =&gt; nil`. 

Investigating more, I found that `store.keys.count` was N, but `store.values.count` was a different number M &lt; N.  That doesn't seem like it should ever happen!  `store.each_pair.count` was also that same M &lt; N.  It would seem the sdbm had become inconsistent with itself somehow -- there were keys listed in `store.keys` which if you asked `store.has_key?(key)` would return false!

My guess is some kind of hash collision error in it's internal implementation -- my 300,000 keys which were all URLs (and all began with `https://samedomain`) perhaps triggered lots of key collisions?

When I find a bug like this, I feel like it's my responsibility to report it upstream. BUT, I was unable to simplify this to a nice reproduction script, a subset of 100 of my keys didn't reproduce, I just know my complete 300,000 key set reproduces. 

And after having lost a few hours to this bug (my brain was not considering the bug could be in sdbm itself, part of stdlib! So I was stumped trying to figure out why my code wasn't producing the results I expected, took me a while to consider and confirm it seemed to be sdbm's fault after all) -- not looking to spend any more on reproducing. 

But i figured reporting it on /r/ruby was at least better than not mentioning it to anyone. 

Anyone else ever used sdbm?  Anyone found any problems with it?

[I switched to `pstore`, and it seems to be working fine now. I was using sdbm because i liked it's API better, and using pstore's transaction API is a bit more confusing, and I was worried if there'd be performance or RAM problems if I kept a transaction open while writing hundreds of thousands of objects... but I'm just doing that, and it seems to be okay].
## [9][Tokens -- Is there a list of reserved tokens for Ruby's syntax?](https://www.reddit.com/r/ruby/comments/gna19o/tokens_is_there_a_list_of_reserved_tokens_for/)
- url: https://www.reddit.com/r/ruby/comments/gna19o/tokens_is_there_a_list_of_reserved_tokens_for/
---
I need list for the reserved tokens Ruby has.

For example, in the expression `"hello " &lt;&lt; "world"`, `&lt;&lt;` is a method of the `String` class, not a token.
## [10][Scaling Rails: Docker &amp; AWS Beanstalk](https://www.reddit.com/r/ruby/comments/gnaiyu/scaling_rails_docker_aws_beanstalk/)
- url: http://jetrockets.pro/blog/scaling-rails-docker-aws-beanstalk
---

