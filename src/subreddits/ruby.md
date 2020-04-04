# ruby
## [1][The Real Difference: `self.method_name` vs `class &lt;&lt; self`](https://www.reddit.com/r/ruby/comments/fummrh/the_real_difference_selfmethod_name_vs_class_self/)
- url: https://emmanuelhayford.com/the-difference-between-self-method-name-and-class-self-in-ruby/
---

## [2][Advanced Sneakers adapter for ActiveJob](https://www.reddit.com/r/ruby/comments/fus9k8/advanced_sneakers_adapter_for_activejob/)
- url: https://www.reddit.com/r/ruby/comments/fus9k8/advanced_sneakers_adapter_for_activejob/
---
Usually [Sidekiq](https://github.com/mperham/sidekiq/) is the best choice for ruby background processing. But if a project has advanced requirements for background processing (routing, heterogeneous consumers, message priorities, etc), it might be a good idea to switch to [RabbitMQ](https://www.rabbitmq.com/)\-based solutions: build ad-hoc consumers using [Bunny](https://github.com/ruby-amqp/bunny) or use a framework (like [Sneakers](https://github.com/jondot/sneakers) and [Hutch](https://github.com/jondot/sneakers)). Any of these approaches require some boilerplate code to be written to play nice with rails.

Since Rails 4.2 there is ActiveJob abstraction over background processing solutions which is great improvement. Developers have ability to replace one backend with another without much changes in codebase. But not all backends cover complete ActiveJob interface. Sneakers already has [ActiveJob adapter](https://github.com/rails/rails/blob/master/activejob/lib/active_job/queue_adapters/sneakers_adapter.rb), but it is too basic and requires a lot of work to be done to at least cover same functionality Sidekiq has (like retries with [exponential back-off](https://en.wikipedia.org/wiki/Exponential_backoff) in case of failures or even delayed execution). Hutch [does not have ActiveJob adapter](https://github.com/gocardless/hutch/issues/117).

Here is the another custom ActiveJob adapter for Sneakers - [:advanced\_sneakers](https://github.com/veeqo/advanced-sneakers-activejob). It supports delays and retries. It also allows to customize message params (like routing keys or headers) and exposes AMQP metadata to jobs.
## [3][CSV Shaper: DSL for creating CSV output in Ruby &amp; Rails](https://www.reddit.com/r/ruby/comments/fuerse/csv_shaper_dsl_for_creating_csv_output_in_ruby/)
- url: https://github.com/paulspringett/csv_shaper
---

## [4][How to improve my skill of finding where ruby variables are defined?](https://www.reddit.com/r/ruby/comments/fufxwc/how_to_improve_my_skill_of_finding_where_ruby/)
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
## [5][emacs mode-line color according guard-rspec result?](https://www.reddit.com/r/ruby/comments/fubyhk/emacs_modeline_color_according_guardrspec_result/)
- url: https://www.reddit.com/r/ruby/comments/fubyhk/emacs_modeline_color_according_guardrspec_result/
---
Is there any support / integration between guard-rspec and Emacs / zenburn-theme?

(I think zenburn-theme has integration with flycheck, though it seems not working for me)

&amp;#x200B;

I've seen such ,,, color changing by rspec result... now i'm thinking that maybe i'm crazy or dreamt sth...
## [6][jinb.in | elegant pastebins](https://www.reddit.com/r/ruby/comments/fu4mni/jinbin_elegant_pastebins/)
- url: /r/rubyonrails/comments/fu2g2g/jinbin_elegant_pastebins/
---

## [7][Most common Ruby on Rails vulnerabilities and how to deal with them in your projects](https://www.reddit.com/r/ruby/comments/ftp02g/most_common_ruby_on_rails_vulnerabilities_and_how/)
- url: https://hixonrails.com/ruby-on-rails-tutorials/ruby-on-rails-security-best-practices/
---

## [8][Join RubyWine Online conference Saturday 10:00 UTC](https://www.reddit.com/r/ruby/comments/ftwffd/join_rubywine_online_conference_saturday_1000_utc/)
- url: https://rubywine.org
---

## [9][Veteran developers, how do you get acclimated to a large new codebase?](https://www.reddit.com/r/ruby/comments/ft5fbc/veteran_developers_how_do_you_get_acclimated_to_a/)
- url: https://www.reddit.com/r/ruby/comments/ft5fbc/veteran_developers_how_do_you_get_acclimated_to_a/
---
I just landed my second Ruby job a few weeks ago. The codebase is HUGE and (for me) mind-bendingly complex, with layers upon layers of ... interesting ... architectural decisions in various stages of deprecation.

For those of you who've been around the block a few times, how do you get acclimated to a new codebase when a new place onboards you? What questions are you asking yourself? Are you looking for or noticing anything in particular that might help you down the line? Are you taking notes?

It definitely takes time and practice. But are there any other ideas that might help in the meantime?

EDIT: The amount of great advice in this thread blows me away. Thanks!
## [10][Programming languages not required!](https://www.reddit.com/r/ruby/comments/ftk7kf/programming_languages_not_required/)
- url: https://medium.com/the-developers-journey/programming-languages-not-required-6fd0422e9dec?source=friends_link&amp;sk=227ce47f96f0aab7dcd803538d500953
---

