# ruby
## [1][Top 10 Most Popular Programming Languages (PYPL) - 2004/ October 2020](https://www.reddit.com/r/ruby/comments/jh63jv/top_10_most_popular_programming_languages_pypl/)
- url: https://youtu.be/DL37toLMCJ8
---

## [2][Stay Up To Date with Ruby on Rails - Automatic PRs in your repo to ease your future Rails upgrade](https://www.reddit.com/r/ruby/comments/jgqsvp/stay_up_to_date_with_ruby_on_rails_automatic_prs/)
- url: https://www.producthunt.com/posts/stay-up-to-date-with-ruby-on-rails
---

## [3][flow-lite gem - a lite workflow engine; let's you define your workflow steps in Flowfiles; incl. the flow command line tool](https://www.reddit.com/r/ruby/comments/jgmdmd/flowlite_gem_a_lite_workflow_engine_lets_you/)
- url: https://github.com/rubycoco/git/tree/master/flow-lite
---

## [4][Understanding create_or_find_by vs find_or_create_by one could be 4X faster, with benchmarks](https://www.reddit.com/r/ruby/comments/jg9mbm/understanding_create_or_find_by_vs_find_or_create/)
- url: https://www.mayerdan.com/ruby/2020/10/22/ar-find_or_create
---

## [5][Rails 6.1 supports order DESC for find_each, find_in_batches, and in_batches | BigBinary Blog](https://www.reddit.com/r/ruby/comments/jg6rgs/rails_61_supports_order_desc_for_find_each_find/)
- url: https://blog.bigbinary.com/2020/10/22/rails-6-1-supports-order-desc-for-find_each-find_in_batches-and-in_batches.html
---

## [6][Object#deep_freeze(skip_shareable: false) and shareable pragma proposals](https://www.reddit.com/r/ruby/comments/jg1e8w/objectdeep_freezeskip_shareable_false_and/)
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
## [7][Create Toast Notifications on Rails 6 application using Yarn](https://www.reddit.com/r/ruby/comments/jgfzl0/create_toast_notifications_on_rails_6_application/)
- url: https://www.youtube.com/watch?v=XWtP9RSIbLo
---

## [8][How I set up a Rails application for testing](https://www.reddit.com/r/ruby/comments/jg50p8/how_i_set_up_a_rails_application_for_testing/)
- url: https://www.codewithjason.com/set-rails-application-testing/
---

## [9][Strange Rubocop behaviour](https://www.reddit.com/r/ruby/comments/jgbt9n/strange_rubocop_behaviour/)
- url: https://www.reddit.com/r/ruby/comments/jgbt9n/strange_rubocop_behaviour/
---
Can anyone explain this for me, please.

* I'm working in a gem project directory.
* Gemfile.lock specifies rubocop 0.50.0
* OS - x86_64 Linux
* shell - zsh
* ruby -v 2.6.6 (managed by rbenv)

If I run `bundle exec rubocop` it uses the `.rubocop.yml` from my `$HOME` directory.

If I run `bundle exec rubocop --config .rubocop.yml` it uses the `.rubocop.yml` from the project root directory.

Why am I having to pass the `--config` flag to force rubocop to use the project specific configuration file? I thought rubocop was supposed to use the first config file it found?

*edit:*

A little more info, after discovering the `--debug` flag:

Without the `--config` flag, the local config file is being used and `AllCops/Exclude configuration` being taken from `$HOME` config. This is as per documented behaviour. My issue was that the config in `$HOME` specified `TargetRubyVersion: "2.6"` which is not recognized by rubocop 0.50.0 resulting in 

    Error: Unknown Ruby version 2.6 found in `TargetRubyVersion` parameter (in /home/442401/.rubocop.yml).
    Known versions: 1.9, 2.0, 2.1, 2.2, 2.3, 2.4
## [10][Can my_hash1.delete(123) be unsafe multithreading-wise?](https://www.reddit.com/r/ruby/comments/jg3y56/can_my_hash1delete123_be_unsafe_multithreadingwise/)
- url: https://www.reddit.com/r/ruby/comments/jg3y56/can_my_hash1delete123_be_unsafe_multithreadingwise/
---
Even if multiple threads try to delete a key with the same number from hash

    my_hash1.delete(123)

it won't be matter because deletion from a hash an item with a certain key can only change a state of hash once. Calling \`delete(key)\` multiple times, or from multiple threads won't cause any issue.  

Therefore, **my\_hash1.delete(key)** always thread-safe, correct?

&amp;#x200B;
