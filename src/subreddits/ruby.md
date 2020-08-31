# ruby
## [1][Ruby Creator Yukihiro Matsumoto on the challenges of updating a programming language](https://www.reddit.com/r/ruby/comments/ijltu6/ruby_creator_yukihiro_matsumoto_on_the_challenges/)
- url: https://thenewstack.io/ruby-creator-yukihiro-matsumoto-on-the-challenges-of-updating-a-programming-language/
---

## [2][Version number bumped to 3.0.0 from 2.8.0 (tentative). · ruby/ruby@21c62fb · GitHub](https://www.reddit.com/r/ruby/comments/ijwbei/version_number_bumped_to_300_from_280_tentative/)
- url: https://github.com/ruby/ruby/commit/21c62fb670b1646c5051a46d29081523cd782f11
---

## [3][dry-types &amp; dry-struct - why? (or why I stopped working on virtus)](https://www.reddit.com/r/ruby/comments/ijtni4/drytypes_drystruct_why_or_why_i_stopped_working/)
- url: https://www.youtube.com/watch?v=eTrOeTiLZGk
---

## [4][Buff exception backtrace with local variables, arguments and instance variables!](https://www.reddit.com/r/ruby/comments/ijwoco/buff_exception_backtrace_with_local_variables/)
- url: https://www.reddit.com/r/ruby/comments/ijwoco/buff_exception_backtrace_with_local_variables/
---
A while ago, I wrote a debugging tool called [tapping\_device](https://github.com/st0012/tapping_device) and received some pretty good responses here. So today I want to introduce another new tool I'm working on recently: [power\_trace](https://github.com/st0012/power_trace). 

It adds all the information we desperately need to exception backtraces:

1. Local variables
2. Method call arguments
3. Instance variables defined in the frame

And it now supports integration with `rails`, `minitest` and `rspec`

Here's an example output of the `rspec` integration  


https://preview.redd.it/laipf1mmvbk51.png?width=1792&amp;format=png&amp;auto=webp&amp;s=5778de0990eda73d3fbcedcd059df641b2f9fb11

Please give it a try let me know what you think about it 😄
## [5][Comparable module in Ruby](https://www.reddit.com/r/ruby/comments/ijr9tn/comparable_module_in_ruby/)
- url: https://www.sandipmane.dev/comparable-module-in-ruby
---

## [6][Ruqqus API for Ruby](https://www.reddit.com/r/ruby/comments/ijr3xx/ruqqus_api_for_ruby/)
- url: https://www.reddit.com/r/ruby/comments/ijr3xx/ruqqus_api_for_ruby/
---
While I'm not the author, I participate in both Ruqqus and Reddit and I [saw this gem there](https://ruqqus.com/post/2q92/made-a-ruqqus-api-for-ruby), which IMO is pretty well made.

I will leave the links here in case you want to contribute.

* [GitHub](https://github.com/ForeverZer0/ruqqus)
* [RubyGems](https://rubygems.org/gems/ruqqus)
* [Documentation](https://www.rubydoc.info/gems/ruqqus/)
## [7][A local JSON database using Rudash for Hash traversal](https://www.reddit.com/r/ruby/comments/ijqkri/a_local_json_database_using_rudash_for_hash/)
- url: https://github.com/rowdb/rowdb
---

## [8][Writing a Discord Bot in Ruby using DiscordRB](https://www.reddit.com/r/ruby/comments/ijgg1l/writing_a_discord_bot_in_ruby_using_discordrb/)
- url: https://www.maxbits.net/blog/post/Writing-a-Discord-Bot-in-Ruby-using-DiscordRB/
---

## [9][A new version of Ruby Firebird Extension Library using Rust instead of C](https://www.reddit.com/r/ruby/comments/ijcfgc/a_new_version_of_ruby_firebird_extension_library/)
- url: https://github.com/fernandobatels/rbfbclient
---

## [10][Which line is faster?](https://www.reddit.com/r/ruby/comments/ijf3xb/which_line_is_faster/)
- url: https://www.reddit.com/r/ruby/comments/ijf3xb/which_line_is_faster/
---
Tiny experiment ahead, I just want to concatenate strings from a collection of hashes. 

I'll share my analysis (and what I've got as the correct answer) tomorrow.

To create a collection of hashes with strings we'll use: 

```ruby
Array.new(count, { text: 'text' }) #=&gt; [{ text: 'text' }, { text: 'text' }, ... "count" times]
```

First option is a classic `map` and `join`.

The second option is a lovely `inject`.

Cast your votes!

[View Poll](https://www.reddit.com/poll/ijf3xb)
