# ruby
## [1][Ruby for Good (Online)](https://www.reddit.com/r/ruby/comments/hw718c/ruby_for_good_online/)
- url: https://www.reddit.com/r/ruby/comments/hw718c/ruby_for_good_online/
---
Hey all,

I'm delighted to share that [registration ](https://ti.to/codeforgood/rubyforgood)for our second virtual Ruby for Good (RfG) event is now open! This is an annual event held every year since 2014, based out of the DC-metro area, where Ruby programmers (of all skill levels), from all over the globe, get together to build technology solutions in support of nonprofits with meaningful missions that benefit our communities. 

Last year we built applications for diaper banks, refuge restrooms, the amazon conservation team, and many other nonprofits serving critical missions, that really need support from people with diverse technology skill-sets, that would not otherwise be able to afford this support. This year, we will be focusing on projects that affect our local communities, and given travel and gathering restrictions, the 2020 event will be virtual.

You can find more information about the event [here](https://rubyforgood.org/attend), a list of the most frequently asked questions [here](https://rubyforgood.org/faq), and the event schedule [here](https://rubyforgood.org/2020).I'd be thrilled to answer any questions!

Happiness,

Sean and the RfG team
## [2][How to speed up assets precompile for Ruby on Rails apps](https://www.reddit.com/r/ruby/comments/hwcu7e/how_to_speed_up_assets_precompile_for_ruby_on/)
- url: https://www.reddit.com/r/ruby/comments/hwcu7e/how_to_speed_up_assets_precompile_for_ruby_on/
---
You spend too much time to deploy your rails project especially on the assets:precompile step, or maybe sometimes you see the following error during the assets precompilation:

*FATAL ERROR: Ineffective mark-compacts near heap limit Allocation failed - JavaScript heap out of memory*

So for sure, this short article will help you to have 50% improvement only with some minor changes:

[https://medium.com/@ali.sepehri.kh/how-to-speed-up-assets-precompile-for-ruby-on-rails-apps-e0338d8d7301?source=friends\_link&amp;sk=995eaef8fdbc73e533c6384f6d47d37e](https://medium.com/@ali.sepehri.kh/how-to-speed-up-assets-precompile-for-ruby-on-rails-apps-e0338d8d7301?source=friends_link&amp;sk=995eaef8fdbc73e533c6384f6d47d37e)
## [3][Am I implementing Dragonfly Gem correctly to my Ruby on Rails app ? [joke]](https://www.reddit.com/r/ruby/comments/hwf5t1/am_i_implementing_dragonfly_gem_correctly_to_my/)
- url: https://github.com/markevans/dragonfly/issues/509
---

## [4][How to handle STI with Fast JSON API serializer for Ruby objects](https://www.reddit.com/r/ruby/comments/hw1u9u/how_to_handle_sti_with_fast_json_api_serializer/)
- url: https://www.reddit.com/r/ruby/comments/hw1u9u/how_to_handle_sti_with_fast_json_api_serializer/
---
 A most popular and faster gem for serializing data in Ruby on Rails: [Fast JSON API](https://github.com/Netflix/fast_jsonapi)

Apparently, this gem is not supporting STI. Check issue: [Lookup for STI models](https://github.com/Netflix/fast_jsonapi/issues/225)

My only way around this is by [Conditional Attributes](https://github.com/Netflix/fast_jsonapi#conditional-attributes) and check if an object is some particular inherited class to parse additional attributes and relations.

How do you handle this?
## [5][Please help](https://www.reddit.com/r/ruby/comments/hwd5wz/please_help/)
- url: https://www.reddit.com/r/ruby/comments/hwd5wz/please_help/
---
I am fairly new to ruby and am trying to create a calculator.However the program enters into a loop and does not return to terminal or exits.Please help me.

code in imgur

[https://imgur.com/gallery/aCCxPtd](https://imgur.com/gallery/aCCxPtd)
## [6][How to please standardrb/rubocop Style/MissingRespondToMissing when using method_missing style delegation?](https://www.reddit.com/r/ruby/comments/hw8t6a/how_to_please_standardrbrubocop/)
- url: https://www.reddit.com/r/ruby/comments/hw8t6a/how_to_please_standardrbrubocop/
---
I am working on conforming a project to [standardrb](https://github.com/testdouble/standard) and am trying to avoid adding exceptions. I have the following class that I use to act as a decorator for page objects in a scraping project: https://gist.github.com/ttilberg/90eee60fccc9ca8b41601d831789e9f1

I'm getting two offenses. I don't believe the first one even makes sense in this context. The second I don't technically understand.

&gt; doc_wrapper.rb:34:3: Style/MethodMissingSuper: When using `method_missing`, fall back on `super`.

I can understand how this makes sense in a `method_missing` implementation that doesn't fully delegate all method calls to another object. You'd want to continue to pass the method call up the chain if you weren't going to handle it. However, in this case, I am fully expecting `@doc` to handle any calls I'm not explicitly defining, so I want the buck to stop at `@doc`. If he can't handle it, no one can... right?

&gt; /doc_wrapper.rb:34:3: Style/MissingRespondToMissing: When using `method_missing`, define `respond_to_missing?`.

I don't understand this one from a technical sense. I've read a bit about it's usage, but everything surrounds dynamic method selection (if the method =~ /something/). It seems odd to have to do that, since I'm already delegating _everything_ to `@obj`. But a simple test case confirmed it doesn't delegate that. `page.respond_to? :css # =&gt; false ;  page.doc.respond_to? :css  #=&gt; true`. Weird.

This raises exception about calling a private method:

    def respond_to_missing? *args
      @obj.respond_to_missing? *args
    end

Trying to use `@obj.send :respond_to_missing? *args` also doesn't work. What is the correct implementation?

I want to note that I'm choosing to use `method_missing` delegation rather than `SimpleDelegator` because I wish to conform the incoming object at #initialize. For a moment, I thought I liked that more. It kept the `doc` easily separated just in case I overload a method that I _do_ still need from the target object... But I guess that's what `__getobj__` is for? /shrug.

What are your thoughts on these two errors with this file? Would you make an exception in standardrb, or would you correct these somehow?
## [7][Ruby Method Overloading](https://www.reddit.com/r/ruby/comments/hvpwll/ruby_method_overloading/)
- url: https://lucaguidi.com/2020/07/22/ruby-method-overloading/
---

## [8][Rollout UI - Admin Dashboard for Feature Flags](https://www.reddit.com/r/ruby/comments/hvsip1/rollout_ui_admin_dashboard_for_feature_flags/)
- url: https://github.com/fetlife/rollout-ui
---

## [9][Ractor - Ruby's Actor-like concurrent abstraction](https://www.reddit.com/r/ruby/comments/hvpwe0/ractor_rubys_actorlike_concurrent_abstraction/)
- url: https://github.com/ko1/ruby/blob/ractor_parallel/doc/ractor.md
---

## [10][A Deeper Look at Ruby's Enumerable](https://www.reddit.com/r/ruby/comments/hvr9xe/a_deeper_look_at_rubys_enumerable/)
- url: https://tutswiki.com/ruby-enumerable/
---

