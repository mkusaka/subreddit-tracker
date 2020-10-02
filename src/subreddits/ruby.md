# ruby
## [1][Ruby 2.7.2 Released](https://www.reddit.com/r/ruby/comments/j3t9hg/ruby_272_released/)
- url: https://www.ruby-lang.org/en/news/2020/10/02/ruby-2-7-2-released/
---

## [2][Code Golf: Count vowels in each word of a sentence](https://www.reddit.com/r/ruby/comments/j3qkbw/code_golf_count_vowels_in_each_word_of_a_sentence/)
- url: https://www.reddit.com/r/ruby/comments/j3qkbw/code_golf_count_vowels_in_each_word_of_a_sentence/
---
I just placed 2nd in a 15 minute code golf challenge. The challenge was to take an input sentence and print a space-separated list of integers for the number of vowels in each word (e.g. "This example sentence I made." -&gt; "1 3 3 1 2")

My code (50 characters) was:

    $&gt;&lt;&lt;gets.split.map{|w|w.scan(/[aeiou]/i).size}*" "

The winner also used Ruby and his code was 44 characters, but he didn't share. I have no idea how it's possible to shave 6 characters off this solution. Can anyone see a way?
## [3][some explanations on fibers?](https://www.reddit.com/r/ruby/comments/j3t1z3/some_explanations_on_fibers/)
- url: https://www.reddit.com/r/ruby/comments/j3t1z3/some_explanations_on_fibers/
---
let's say I have the following (multithreaded) code

`def dputs(string = '', delay = 1, newline = true)`

`string.chars.each do |i|`

`print(i)`

`sleep(delay)`

`end`

`print("\n") if newline`

`end`

&amp;#x200B;

`[Thread.new { dputs('a' * 5) },`

`Thread.new { dputs('b' * 5) }].map(&amp;:join)`

How do I have to rewrite this using fibers rather than threads?

I have read [this article](https://www.rubyguides.com/2019/11/what-are-fibers-in-ruby/) but still can't understand
## [4][faster_s3_url: Optimized S3 url generation in ruby](https://www.reddit.com/r/ruby/comments/j3dygu/faster_s3_url_optimized_s3_url_generation_in_ruby/)
- url: https://github.com/jrochkind/faster_s3_url
---

## [5][Installing Ruby with ruby-build and ruby-install](https://www.reddit.com/r/ruby/comments/j3idi7/installing_ruby_with_rubybuild_and_rubyinstall/)
- url: https://nts.strzibny.name/ruby-build-and-ruby-install/
---

## [6][What surprised us in PostgreSQL-schema based multitenancy](https://www.reddit.com/r/ruby/comments/j365oy/what_surprised_us_in_postgresqlschema_based/)
- url: https://blog.arkency.com/what-surprised-us-in-postgres-schema-multitenancy/
---

## [7][One-liner introduction - Ruby one-liners cookbook](https://www.reddit.com/r/ruby/comments/j34eyl/oneliner_introduction_ruby_oneliners_cookbook/)
- url: https://learnbyexample.github.io/learn_ruby_oneliners/one-liner-introduction.html
---

## [8][Procedure - new rails design pattern](https://www.reddit.com/r/ruby/comments/j36wjs/procedure_new_rails_design_pattern/)
- url: https://longliveruby.com/articles/rails-procedure-design-pattern
---

## [9][Ruby 3.0 Freezes Range Objects](https://www.reddit.com/r/ruby/comments/j332ai/ruby_30_freezes_range_objects/)
- url: https://scriptday.com/blog/2020/10/01/ruby-3-0-freezes-range-objects
---

## [10][Gathering ideas: Ruby is the best programming language for beginners](https://www.reddit.com/r/ruby/comments/j2tu7k/gathering_ideas_ruby_is_the_best_programming/)
- url: https://www.reddit.com/r/ruby/comments/j2tu7k/gathering_ideas_ruby_is_the_best_programming/
---
I had an epiphany: Ruby is much better than Python for someone learning to program. :-) I realize those are fighting words in many corners of the Internet. Probably also in academia.

I'm putting together my thoughts for blog posts. Does anyone here have any points or anecdotes about learning Ruby as a first language, or teaching someone?

I had my realization [while coping with Python's cruft and arbitrary design decisions](https://www.reddit.com/r/learnpython/comments/ispgri/any_way_to_make_textsplitjoin_capitalize_possible/). It had been a while since I'd done much Python, and I was surprised at how bad the UX is. Python is held up as an amazing teaching language, yet it presents an uneven smorgasbord of global functions (which are misleadingly called "builtins"), classes, mutating "functions" and intentionally-difficult-to-use fp features.
