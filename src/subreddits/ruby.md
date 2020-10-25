# ruby
## [1][Benchable 0.2.0 adds support for memory benchmarks](https://www.reddit.com/r/ruby/comments/jhjefw/benchable_020_adds_support_for_memory_benchmarks/)
- url: https://i.redd.it/fj6eb1mlq4v51.jpg
---

## [2][Nested each](https://www.reddit.com/r/ruby/comments/jhr2gx/nested_each/)
- url: https://www.reddit.com/r/ruby/comments/jhr2gx/nested_each/
---
Hi guys, I’m on day three of learning ruby, I’m trying to use nested each statements to look for matches in two separate arrays and return the matches in a hash with the quantity of each match as the value. I cannot for love nor money get the syntax right. Can someone please give me a hand with how it should look? Tia
## [3][Tailwind CSS + Rails via webpack](https://www.reddit.com/r/ruby/comments/jhcsjd/tailwind_css_rails_via_webpack/)
- url: https://www.boringgenerators.com/blog/2020-10-18-install-tailwind/
---

## [4][Top 10 Most Popular Programming Languages (PYPL) - 2004/ October 2020](https://www.reddit.com/r/ruby/comments/jh63jv/top_10_most_popular_programming_languages_pypl/)
- url: https://youtu.be/DL37toLMCJ8
---

## [5][Stay Up To Date with Ruby on Rails - Automatic PRs in your repo to ease your future Rails upgrade](https://www.reddit.com/r/ruby/comments/jgqsvp/stay_up_to_date_with_ruby_on_rails_automatic_prs/)
- url: https://www.producthunt.com/posts/stay-up-to-date-with-ruby-on-rails
---

## [6][flow-lite gem - a lite workflow engine; let's you define your workflow steps in Flowfiles; incl. the flow command line tool](https://www.reddit.com/r/ruby/comments/jgmdmd/flowlite_gem_a_lite_workflow_engine_lets_you/)
- url: https://github.com/rubycoco/git/tree/master/flow-lite
---

## [7][Understanding create_or_find_by vs find_or_create_by one could be 4X faster, with benchmarks](https://www.reddit.com/r/ruby/comments/jg9mbm/understanding_create_or_find_by_vs_find_or_create/)
- url: https://www.mayerdan.com/ruby/2020/10/22/ar-find_or_create
---

## [8][Rails 6.1 supports order DESC for find_each, find_in_batches, and in_batches | BigBinary Blog](https://www.reddit.com/r/ruby/comments/jg6rgs/rails_61_supports_order_desc_for_find_each_find/)
- url: https://blog.bigbinary.com/2020/10/22/rails-6-1-supports-order-desc-for-find_each-find_in_batches-and-in_batches.html
---

## [9][Object#deep_freeze(skip_shareable: false) and shareable pragma proposals](https://www.reddit.com/r/ruby/comments/jg1e8w/objectdeep_freezeskip_shareable_false_and/)
- url: https://www.reddit.com/r/ruby/comments/jg1e8w/objectdeep_freezeskip_shareable_false_and/
---
Exciting work by Koichi, Nobu and others to make Ractors more usable.

Object#deep\_freeze(skip\_shareable: false)

    a = ['s', Ractor.new{}].deep_freeze(skip_shareable: true)
    a[0].frozen? # =&gt; true    
    a[1].frozen? # =&gt; false
    
    Ractor.new do
      p a
    end.take

[https://github.com/ruby/ruby/pull/3669](https://github.com/ruby/ruby/pull/3669)

The shareable pragma proposal allows you to avoid having to call freeze on every object:

    # shareable_constant_value: true
    
    A = [1, [2, [3, 4]]]
    H = {a: "a"}
    
    Ractor.new do
      p A
      p H
    end.take

[https://bugs.ruby-lang.org/issues/17273](https://bugs.ruby-lang.org/issues/17273)
## [10][Create Toast Notifications on Rails 6 application using Yarn](https://www.reddit.com/r/ruby/comments/jgfzl0/create_toast_notifications_on_rails_6_application/)
- url: https://www.youtube.com/watch?v=XWtP9RSIbLo
---

