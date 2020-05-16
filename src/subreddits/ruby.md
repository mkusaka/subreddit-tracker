# ruby
## [1][Can you have a "literal" string literal in Ruby? Can I make single-quoted string *not* escape \\?](https://www.reddit.com/r/ruby/comments/gkuphz/can_you_have_a_literal_string_literal_in_ruby_can/)
- url: https://www.reddit.com/r/ruby/comments/gkuphz/can_you_have_a_literal_string_literal_in_ruby_can/
---
I've been doing the 2015 Advent of Code problems to practice my Ruby and I ran across an interesting issue on Day 8.  Day 8 gives you a list of strings and a list of "escape sequences" --- `\"`, `\\`, and `\x[0-9a-f]{2}`. You need to count how many characters in each string are escape characters.

The problem I ran into is that I can't figure out how to get Ruby to read a literal string as having two backslashes (`\\`) in a row. I thought single-quoted strings in Ruby ignored escape sequences, but even in single-quoted Ruby strings, `\\` is an escaped backslash character.

    irb(main):003:0&gt; string1 = 'test\test'
    =&gt; "test\\test"
    irb(main):004:0&gt; string2 = 'test\\test'
    =&gt; "test\\test"
    irb(main):005:0&gt; string1 == string2
    =&gt; true
    irb(main):006:0&gt; string1.size
    =&gt; 9
    irb(main):007:0&gt; string2.size
    =&gt; 9
    irb(main):008:0&gt; string1.bytes
    =&gt; [116, 101, 115, 116, 92, 116, 101, 115, 116]
    irb(main):009:0&gt; string2.bytes
    =&gt; [116, 101, 115, 116, 92, 116, 101, 115, 116]

I solved the AoC problem by reading the strings character by character into an array and working with that, but that doesn't feel like a very Ruby-like solution to me.

My question is: Is there a way to create a string literal in Ruby that ignores \*all\* escaped characters, or a way to get at the "raw" unescaped data that was passed into the string when it was created?
## [2][Why Fullstaq Ruby?](https://www.reddit.com/r/ruby/comments/gk5i19/why_fullstaq_ruby/)
- url: https://www.joyfulbikeshedding.com/blog/2020-05-15-why-fullstaq-ruby.html
---

## [3][sportdb-importers Gem v1.0 - tools 'n' scripts for importing sports (incl. football) data in comma-separated values (csv) format](https://www.reddit.com/r/ruby/comments/gke7v6/sportdbimporters_gem_v10_tools_n_scripts_for/)
- url: https://github.com/sportdb/sport.db/tree/master/sportdb-importers
---

## [4][A heavily tested (2k lines) and commented classic Red Black Tree implementation in Python and Ruby. Great for learning the material.](https://www.reddit.com/r/ruby/comments/gk51nf/a_heavily_tested_2k_lines_and_commented_classic/)
- url: https://www.reddit.com/r/ruby/comments/gk51nf/a_heavily_tested_2k_lines_and_commented_classic/
---
Back when I was trying to implement the structure, I could not find any open source implementations that were well written and commented. I did not manage to find any implementation that had any significant amount of tests and as such was not sure if it even worked correctly.

I tried my best to describe the different operations needed thoroughly and have written a lot of tests (functional too) covering all operations, with drawn out trees in comments.

https://github.com/stanislavkozlovski/Red-Black-Tree

Any feedback is greatly appreciated :)
## [5][Foot Traffic: pure Ruby DSL for Chrome scripting based on Ferrum](https://www.reddit.com/r/ruby/comments/gjqaji/foot_traffic_pure_ruby_dsl_for_chrome_scripting/)
- url: https://github.com/lewagon/foot_traffic
---

## [6][Aaron Patterson (tenderlove) and Zamith are live streaming an hour long conversation on Ruby](https://www.reddit.com/r/ruby/comments/gjqexq/aaron_patterson_tenderlove_and_zamith_are_live/)
- url: https://twitch.tv/wearesubvisual
---

## [7][LeetCode Challenges Contains Duplicate | Jewels and Stones](https://www.reddit.com/r/ruby/comments/gk2iby/leetcode_challenges_contains_duplicate_jewels_and/)
- url: https://youtu.be/3g7BGGzGw5I
---

## [8][How to Fix Slow Code in Ruby](https://www.reddit.com/r/ruby/comments/gjqwh9/how_to_fix_slow_code_in_ruby/)
- url: https://engineering.shopify.com/blogs/engineering/how-fix-slow-code-ruby
---

## [9][Matz is calling for feedback on Ruby 2.7/3.0 keyword argument pain](https://www.reddit.com/r/ruby/comments/gjjayp/matz_is_calling_for_feedback_on_ruby_2730_keyword/)
- url: https://discuss.rubyonrails.org/t/new-2-7-3-0-keyword-argument-pain-point/74980
---

## [10][Thread scheduler for light weight concurrency. by ioquatix · Pull Request #3032 · was merged. :heart:](https://www.reddit.com/r/ruby/comments/gjjtla/thread_scheduler_for_light_weight_concurrency_by/)
- url: https://github.com/ruby/ruby/pull/3032
---

