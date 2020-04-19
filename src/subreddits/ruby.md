# ruby
## [1][Ruby Hash Default Values](https://www.reddit.com/r/ruby/comments/g33oyz/ruby_hash_default_values/)
- url: https://www.reddit.com/r/ruby/comments/g33oyz/ruby_hash_default_values/
---
I've written a new doc section about Hash default values.  It has been merged into Ruby's [hash.c](https://github.com/ruby/ruby/commit/39c965f2306d29524feae04ff6e710b32702851a#diff-eff9999082c8ce7d8ba1fc1d79f439cf), but not yet released.

I wrote it over at my doc project [AboutRuby](https://github.com/BurdetteLamar/AboutRuby/blob/master/core/Hash/api/markdown.md#hash-api), then PRed it over to Ruby. (Thanks, @nobu, for the review and merge.)
## [2][Processing high volume of unique messages exactly once while preserving order in a queue using SQS and Ruby-on-Rails](https://www.reddit.com/r/ruby/comments/g37ygs/processing_high_volume_of_unique_messages_exactly/)
- url: https://medium.com/@rcdexta/processing-high-volume-of-unique-messages-exactly-once-while-preserving-order-in-a-queue-d8d6184ded01
---

## [3][About Hash Keys](https://www.reddit.com/r/ruby/comments/g36d0g/about_hash_keys/)
- url: https://www.reddit.com/r/ruby/comments/g36d0g/about_hash_keys/
---
I've written about [Hash Keys](https://github.com/BurdetteLamar/AboutRuby/blob/master/core/Hash/api/markdown.md#hash-keys).  (This is not yet merged into the Ruby code base.)
## [4][Three Examples of DTOs in Ruby](https://www.reddit.com/r/ruby/comments/g32yva/three_examples_of_dtos_in_ruby/)
- url: https://medium.com/@minghz42/three-examples-of-dtos-in-ruby-df22853ea128
---

## [5][Starting with Rails](https://www.reddit.com/r/ruby/comments/g2ya9n/starting_with_rails/)
- url: https://www.reddit.com/r/ruby/comments/g2ya9n/starting_with_rails/
---
Hey! I'm interested in getting into Ruby on Rails, and I'm wondering whether any of you might be able to point me in the direction of some good resources.

I have active subscriptions to [Lynda.com](https://lynda.com/) and [GoRails.com](https://gorails.com/), and I have the GitHub Student Developer Pack. My end-goal is to build a basic social-networking site for my school, not to become a paid web developer!

I have loads of experience in Python, HTML, JS and CSS, and I launched myself into a Basics of Ruby course on Lynda, so I have enough experience there too.

I was watching this free Udemy course, which looked perfect; [https://www.udemy.com/course/8-beautiful-ruby-on-rails-apps-in-30-days/learn/lecture/4336792?start=240](https://www.udemy.com/course/8-beautiful-ruby-on-rails-apps-in-30-days/learn/lecture/4336792?start=240), but in the Announcements section it was apparently severely outdated. Does anyone know of something similar? I honestly prefer video content to reading (with the exception of books).

And I'd prefer to not spend heaps. I've looked at The Odin Project and the Essential RoR Training courses on Lynda but the RoR course seems far too theoretical. I want to get creating ASAP.

Many thanks!
## [6][About Hash](https://www.reddit.com/r/ruby/comments/g39a34/about_hash/)
- url: https://www.reddit.com/r/ruby/comments/g39a34/about_hash/
---
And I've written [About Hash](https://github.com/BurdetteLamar/AboutRuby/blob/master/core/Hash/about/markdown.md#about-hash), also not yet merged into the Ruby core. This is everything that's not actually API-spec.
## [7][Irb and Visual Code wierd act](https://www.reddit.com/r/ruby/comments/g2wd30/irb_and_visual_code_wierd_act/)
- url: https://www.reddit.com/r/ruby/comments/g2wd30/irb_and_visual_code_wierd_act/
---
when I am calling irb from regular terminal I get:

    kun@lilpc ~/P/test&gt; irb
    Traceback (most recent call last):
    	25: from /home/kun/.rvm/gems/ruby-2.7.0/bin/ruby_executable_hooks:24:in `&lt;main&gt;'
    	24: from /home/kun/.rvm/gems/ruby-2.7.0/bin/ruby_executable_hooks:24:in `eval'
    	23: from /home/kun/.rvm/gems/ruby-2.7.0/bin/irb:23:in `&lt;main&gt;'
    	22: from /home/kun/.rvm/gems/ruby-2.7.0/bin/irb:23:in `load'
    	21: from /home/kun/.rvm/gems/ruby-2.7.0/gems/irb-1.2.3/exe/irb:11:in `&lt;top (required)&gt;'
    	20: from /home/kun/.rvm/gems/ruby-2.7.0/gems/irb-1.2.3/lib/irb.rb:399:in `start'
    	19: from /home/kun/.rvm/gems/ruby-2.7.0/gems/irb-1.2.3/lib/irb.rb:470:in `run'
    	18: from /home/kun/.rvm/gems/ruby-2.7.0/gems/irb-1.2.3/lib/irb.rb:470:in `catch'
    	17: from /home/kun/.rvm/gems/ruby-2.7.0/gems/irb-1.2.3/lib/irb.rb:471:in `block in run'
    	16: from /home/kun/.rvm/gems/ruby-2.7.0/gems/irb-1.2.3/lib/irb.rb:536:in `eval_input'
    	15: from /home/kun/.rvm/gems/ruby-2.7.0/gems/irb-1.2.3/lib/irb/ruby-lex.rb:134:in `each_top_level_statement'
    	14: from /home/kun/.rvm/gems/ruby-2.7.0/gems/irb-1.2.3/lib/irb/ruby-lex.rb:134:in `catch'
    	13: from /home/kun/.rvm/gems/ruby-2.7.0/gems/irb-1.2.3/lib/irb/ruby-lex.rb:135:in `block in each_top_level_statement'
    	12: from /home/kun/.rvm/gems/ruby-2.7.0/gems/irb-1.2.3/lib/irb/ruby-lex.rb:135:in `loop'
    	11: from /home/kun/.rvm/gems/ruby-2.7.0/gems/irb-1.2.3/lib/irb/ruby-lex.rb:138:in `block (2 levels) in each_top_level_statement'
    	10: from /home/kun/.rvm/gems/ruby-2.7.0/gems/irb-1.2.3/lib/irb/ruby-lex.rb:166:in `lex'
    	 9: from /home/kun/.rvm/gems/ruby-2.7.0/gems/irb-1.2.3/lib/irb.rb:517:in `block in eval_input'
    	 8: from /home/kun/.rvm/gems/ruby-2.7.0/gems/irb-1.2.3/lib/irb.rb:694:in `signal_status'
    	 7: from /home/kun/.rvm/gems/ruby-2.7.0/gems/irb-1.2.3/lib/irb.rb:518:in `block (2 levels) in eval_input'
    	 6: from /home/kun/.rvm/gems/ruby-2.7.0/gems/irb-1.2.3/lib/irb/input-method.rb:262:in `gets'
    	 5: from /home/kun/.rvm/rubies/ruby-2.7.0/lib/ruby/2.7.0/forwardable.rb:235:in `readmultiline'
    	 4: from /home/kun/.rvm/rubies/ruby-2.7.0/lib/ruby/2.7.0/forwardable.rb:235:in `readmultiline'
    	 3: from /home/kun/.rvm/gems/ruby-2.7.0/gems/reline-0.1.3/lib/reline.rb:174:in `readmultiline'
    	 2: from /home/kun/.rvm/gems/ruby-2.7.0/gems/reline-0.1.3/lib/reline.rb:207:in `inner_readline'
    	 1: from /home/kun/.rvm/gems/ruby-2.7.0/gems/reline-0.1.3/lib/reline.rb:339:in `may_req_ambiguous_char_width'
    /home/kun/.rvm/gems/ruby-2.7.0/gems/reline-0.1.3/lib/reline.rb:339:in `write': U+25BD from UTF-8 to US-ASCII (Encoding::UndefinedConversionError)

But when it's called from vs code nothing wrong goes on:

    kun@lilpc ~/P/population (master)&gt; irb
    2.7.0 :001 &gt;

I have an idea that solargraph gem which is my language server for ruby in vs code has some kind of permissions to block another irb starts rather than in vs integrated terminal.

Any ideas how to fix it?
## [8][Full Text Search in Milliseconds with Rails and PostgreSQL](https://www.reddit.com/r/ruby/comments/g2gnlb/full_text_search_in_milliseconds_with_rails_and/)
- url: https://pganalyze.com/blog/full-text-search-ruby-rails-postgres
---

## [9][The Ruby Blend: ViewComponent at GitHub with Joel Hawksley](https://www.reddit.com/r/ruby/comments/g2givv/the_ruby_blend_viewcomponent_at_github_with_joel/)
- url: https://www.reddit.com/r/ruby/comments/g2givv/the_ruby_blend_viewcomponent_at_github_with_joel/
---
Joel Hawksley is the guest on today’s episode. He is a Software Engineer at GitHub and works on the design system team, leading development on ViewComponent, formerly known as ActionView::Component. Joel is a very encouraging, uplifting, and inviting guest who tells it like it is. Andrew even throws in a little “homework” for those of you listening! If you want to learn more about how GitHub is leveraging ViewComponent, [give this episode a listen!](https://fireside.fm/s/ouBAUjGy+8JZwBXVx)
## [10][How to Elicit Exception for Flatten Reentered?](https://www.reddit.com/r/ruby/comments/g2mxvj/how_to_elicit_exception_for_flatten_reentered/)
- url: https://www.reddit.com/r/ruby/comments/g2mxvj/how_to_elicit_exception_for_flatten_reentered/
---
The C code for Array#flatten raises RuntimeError ("flatten reentered") for a certain condition.

The C-coded test tests this, but the Ruby-coded spec does not.

What Ruby code would cause this error?
