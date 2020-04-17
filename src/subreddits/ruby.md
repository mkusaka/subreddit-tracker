# ruby
## [1][Starting with Rails](https://www.reddit.com/r/ruby/comments/g2ya9n/starting_with_rails/)
- url: https://www.reddit.com/r/ruby/comments/g2ya9n/starting_with_rails/
---
Hey! I'm interested in getting into Ruby on Rails, and I'm wondering whether any of you might be able to point me in the direction of some good resources.

I have active subscriptions to [Lynda.com](https://lynda.com/) and [GoRails.com](https://gorails.com/), and I have the GitHub Student Developer Pack. My end-goal is to build a basic social-networking site for my school, not to become a paid web developer!

I have loads of experience in Python, HTML, JS and CSS, and I launched myself into a Basics of Ruby course on Lynda, so I have enough experience there too.

I was watching this free Udemy course, which looked perfect; [https://www.udemy.com/course/8-beautiful-ruby-on-rails-apps-in-30-days/learn/lecture/4336792?start=240](https://www.udemy.com/course/8-beautiful-ruby-on-rails-apps-in-30-days/learn/lecture/4336792?start=240), but in the Announcements section it was apparently severely outdated. Does anyone know of something similar? I honestly prefer video content to reading (with the exception of books).

And I'd prefer to not spend heaps. I've looked at The Odin Project and the Essential RoR Training courses on Lynda but the RoR course seems far too theoretical. I want to get creating ASAP.

Many thanks!
## [2][Full Text Search in Milliseconds with Rails and PostgreSQL](https://www.reddit.com/r/ruby/comments/g2gnlb/full_text_search_in_milliseconds_with_rails_and/)
- url: https://pganalyze.com/blog/full-text-search-ruby-rails-postgres
---

## [3][The Ruby Blend: ViewComponent at GitHub with Joel Hawksley](https://www.reddit.com/r/ruby/comments/g2givv/the_ruby_blend_viewcomponent_at_github_with_joel/)
- url: https://www.reddit.com/r/ruby/comments/g2givv/the_ruby_blend_viewcomponent_at_github_with_joel/
---
Joel Hawksley is the guest on today’s episode. He is a Software Engineer at GitHub and works on the design system team, leading development on ViewComponent, formerly known as ActionView::Component. Joel is a very encouraging, uplifting, and inviting guest who tells it like it is. Andrew even throws in a little “homework” for those of you listening! If you want to learn more about how GitHub is leveraging ViewComponent, [give this episode a listen!](https://fireside.fm/s/ouBAUjGy+8JZwBXVx)
## [4][Irb and Visual Code wierd act](https://www.reddit.com/r/ruby/comments/g2wd30/irb_and_visual_code_wierd_act/)
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
## [5][How to Elicit Exception for Flatten Reentered?](https://www.reddit.com/r/ruby/comments/g2mxvj/how_to_elicit_exception_for_flatten_reentered/)
- url: https://www.reddit.com/r/ruby/comments/g2mxvj/how_to_elicit_exception_for_flatten_reentered/
---
The C code for Array#flatten raises RuntimeError ("flatten reentered") for a certain condition.

The C-coded test tests this, but the Ruby-coded spec does not.

What Ruby code would cause this error?
## [6][Aliases and unary operators](https://www.reddit.com/r/ruby/comments/g2m7vn/aliases_and_unary_operators/)
- url: https://www.reddit.com/r/ruby/comments/g2m7vn/aliases_and_unary_operators/
---
```ruby
class String
  alias !@ reverse
end

!"abc" # =&gt; "cba"
## [7][Ruby/Rails in NY](https://www.reddit.com/r/ruby/comments/g2fj11/rubyrails_in_ny/)
- url: https://www.reddit.com/r/ruby/comments/g2fj11/rubyrails_in_ny/
---
Hi Guys, perhaps some of you are located in NY or you cooperate with companies there. I wonder how is the situation for Ruby devs now? Is it the same case as in EU that there are too many devs currently on the market due to the reductions? 
Thanks in advance for your help!
## [8][What do you use for keeping track of your API?](https://www.reddit.com/r/ruby/comments/g2lijc/what_do_you_use_for_keeping_track_of_your_api/)
- url: https://www.reddit.com/r/ruby/comments/g2lijc/what_do_you_use_for_keeping_track_of_your_api/
---
Hey!

JSON Schema, Swagger, RAML to name a few. I wonder how do you *document* API HTTP endpoints in your applications?
## [9][Can this be built in ruby easily?](https://www.reddit.com/r/ruby/comments/g2wzuc/can_this_be_built_in_ruby_easily/)
- url: https://www.reddit.com/r/ruby/comments/g2wzuc/can_this_be_built_in_ruby_easily/
---
Hi All

i am new to web development and trying to build a project of a client portal to submit jobs to us, as per these images &gt;  [https://imgur.com/a/1fMaNOj](https://imgur.com/a/1fMaNOj) 

What would the easiest framework to use to achieve this? or is there something pre-existing i can modify?
## [10][Don't add database index if it already exists in Rails](https://www.reddit.com/r/ruby/comments/g2drmb/dont_add_database_index_if_it_already_exists_in/)
- url: https://prathamesh.tech/2020/04/16/dont-add-database-index-if-it-already-exists/
---

