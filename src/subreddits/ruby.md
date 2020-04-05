# ruby
## [1][I've written a tiny gem which allows you to rerun your last IRB command using the "did you mean" suggestion if you make a typo. Hopefully it's a handy time-saver!](https://www.reddit.com/r/ruby/comments/fuyo4u/ive_written_a_tiny_gem_which_allows_you_to_rerun/)
- url: https://github.com/AaronC81/yes_i_did
---

## [2][Stimulus Reflex](https://www.reddit.com/r/ruby/comments/fvd1pb/stimulus_reflex/)
- url: /r/rails/comments/fvd1cs/stimulus_reflex/
---

## [3][Small web crawler ruby-based for practice purpose. How should I approach my project ?](https://www.reddit.com/r/ruby/comments/fv4dcw/small_web_crawler_rubybased_for_practice_purpose/)
- url: https://www.reddit.com/r/ruby/comments/fv4dcw/small_web_crawler_rubybased_for_practice_purpose/
---
So, paralleling with my rails course, I do like to code as practice few ruby scripts to get used to more and more ruby gems.  

I did a few projects using Nokogiri, rest client, open-uri, watir, mechanize and few others to scrape things I like. But crawling intere website it’s far more difficult. 

About the project...

I do like buy mangas. I have a list of 20 online stores to buy mangas. Most of it is small websites, so a few people know them. Every time I’m looking for a rare manga I do a search on every store from this list. 

I would like to do a web crawler to automate this for me.  But I would reduce to 3 or 4 stores which the program will crawl to find if the manga  I’m looking for is available.  

I need some guidance on how should I do that...


The scrape programa I made was based on inputs from the user and built the url using these values.  
I don’t think this is a good approach to use in this case. I would have to deal with a lot problems messing with the URL of every website. 

So, I thought using the home page of every store and use the Mechanize gem to find the search field from each store.  (Its that a good way of thinking ?)

So I though two program structure for this project and I don’t know if those are the best pra ctice for doing that.

Every website has its own html structure. Acessing the page results through css classes will be different for each one. 

So how should I approach ? Make exclusively functions to each store? Like each store will be a class and it has its own functions to return what I looking for.

Or the best way is to “global” functions which will have to cover a lot of exceptions due the different structure each website has. 

You can check out my scrape codes [here](https://github.com/nicweeaboo) . It’s really nothing special.
## [4][Tell us a feature written in Ruby that gives you a real sense of achievement](https://www.reddit.com/r/ruby/comments/fva8ll/tell_us_a_feature_written_in_ruby_that_gives_you/)
- url: https://www.reddit.com/r/ruby/comments/fva8ll/tell_us_a_feature_written_in_ruby_that_gives_you/
---
For me, I implemented a slightly modified version of the Quota Rule algorithm that dispatches application forms to different callcenters. 😊

**PS: Don't be shy! This is not a competition. We can all learn from each other!**
## [5][Advanced Sneakers adapter for ActiveJob](https://www.reddit.com/r/ruby/comments/fus9k8/advanced_sneakers_adapter_for_activejob/)
- url: https://www.reddit.com/r/ruby/comments/fus9k8/advanced_sneakers_adapter_for_activejob/
---
Usually [Sidekiq](https://github.com/mperham/sidekiq/) is the best choice for ruby background processing. But if a project has advanced requirements for background processing (routing, heterogeneous consumers, message priorities, etc), it might be a good idea to switch to [RabbitMQ](https://www.rabbitmq.com/)\-based solutions: build ad-hoc consumers using [Bunny](https://github.com/ruby-amqp/bunny) or use a framework (like [Sneakers](https://github.com/jondot/sneakers) and [Hutch](https://github.com/jondot/sneakers)). Any of these approaches require some boilerplate code to be written to play nice with rails.

Since Rails 4.2 there is ActiveJob abstraction over background processing solutions which is great improvement. Developers have ability to replace one backend with another without much changes in codebase. But not all backends cover complete ActiveJob interface. Sneakers already has [ActiveJob adapter](https://github.com/rails/rails/blob/master/activejob/lib/active_job/queue_adapters/sneakers_adapter.rb), but it is too basic and requires a lot of work to be done to at least cover same functionality Sidekiq has (like retries with [exponential back-off](https://en.wikipedia.org/wiki/Exponential_backoff) in case of failures or even delayed execution). Hutch [does not have ActiveJob adapter](https://github.com/gocardless/hutch/issues/117).

Here is the another custom ActiveJob adapter for Sneakers - [:advanced\_sneakers](https://github.com/veeqo/advanced-sneakers-activejob). It supports delays and retries. It also allows to customize message params (like routing keys or headers) and exposes AMQP metadata to jobs.
## [6][The Real Difference: `self.method_name` vs `class &lt;&lt; self`](https://www.reddit.com/r/ruby/comments/fummrh/the_real_difference_selfmethod_name_vs_class_self/)
- url: https://emmanuelhayford.com/the-difference-between-self-method-name-and-class-self-in-ruby/
---

## [7][CSV Shaper: DSL for creating CSV output in Ruby &amp; Rails](https://www.reddit.com/r/ruby/comments/fuerse/csv_shaper_dsl_for_creating_csv_output_in_ruby/)
- url: https://github.com/paulspringett/csv_shaper
---

## [8][How to improve my skill of finding where ruby variables are defined?](https://www.reddit.com/r/ruby/comments/fufxwc/how_to_improve_my_skill_of_finding_where_ruby/)
- url: https://www.reddit.com/r/ruby/comments/fufxwc/how_to_improve_my_skill_of_finding_where_ruby/
---
Hi folks,

I've not yet developed the skill of quickly determining where a variable in ruby is defined.

This leads me to feel frustrated when programming ruby. Imagine if you were cooking and every time you
needed a knife, you had to spend 3 minutes hunting around for it. 3 minutes is not a lot, but compared to
the 3-8 seconds you're used to, it makes the cooking really frustrating.

I would like to becomes as fast at finding variable definitions in ruby as I am in python or javascript

Could you please suggest some

 * courses I can watch/read

 * exercises I can do

 * tools I can use

?

------

Tools I'm going to explore:

For VSCode: Solargraph, suggested by /u/efreed09

For Vim: [Vim-Jump](https://github.com/pechorin/any-jump.vim), suggested by /u/shvedchenko

And credit to /u/fl00pz for the early ask of a pretty important focusing question.
## [9][emacs mode-line color according guard-rspec result?](https://www.reddit.com/r/ruby/comments/fubyhk/emacs_modeline_color_according_guardrspec_result/)
- url: https://www.reddit.com/r/ruby/comments/fubyhk/emacs_modeline_color_according_guardrspec_result/
---
Is there any support / integration between guard-rspec and Emacs / zenburn-theme?

(I think zenburn-theme has integration with flycheck, though it seems not working for me)

&amp;#x200B;

I've seen such ,,, color changing by rspec result... now i'm thinking that maybe i'm crazy or dreamt sth...
## [10][jinb.in | elegant pastebins](https://www.reddit.com/r/ruby/comments/fu4mni/jinbin_elegant_pastebins/)
- url: /r/rubyonrails/comments/fu2g2g/jinbin_elegant_pastebins/
---

