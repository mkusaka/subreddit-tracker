# ruby
## [1][Does Ruby work well for blockchain development? My experience and code](https://www.reddit.com/r/ruby/comments/hv3qdq/does_ruby_work_well_for_blockchain_development_my/)
- url: https://syndicode.com/2020/05/06/blockchain-monitoring-in-the-example-of-bitcoin-and-ruby/
---

## [2][PropCheck: Library to do property-based testing in Ruby](https://www.reddit.com/r/ruby/comments/hv73pm/propcheck_library_to_do_propertybased_testing_in/)
- url: https://github.com/Qqwy/ruby-prop_check/
---

## [3][Solution for those who like Action Cable but hate writing JavaScript](https://www.reddit.com/r/ruby/comments/hv6uk2/solution_for_those_who_like_action_cable_but_hate/)
- url: https://pdabrowski.com/articles/cable-ready-with-action-cable
---

## [4][Help needed with bundler on ubuntu 18.04 LTS](https://www.reddit.com/r/ruby/comments/hv2q7d/help_needed_with_bundler_on_ubuntu_1804_lts/)
- url: https://www.reddit.com/r/ruby/comments/hv2q7d/help_needed_with_bundler_on_ubuntu_1804_lts/
---
Hi guys! New to ruby, trying to set up a website with jekyll (that's currently the only thing I'm using ruby actively for) and ran into this issue with bundler vs ubuntu package manager: [https://github.com/rubygems/rubygems/issues/3068](https://github.com/rubygems/rubygems/issues/3068)

Symptoms: whenever I run bundle it get something like this:

    ...
    /usr/lib/ruby/vendor_ruby/rubygems/defaults/operating_system.rb:10: warning: constant Gem::ConfigMap is deprecated
    Fetching gem metadata from https://rubygems.org/...........
    /usr/lib/ruby/vendor_ruby/rubygems/defaults/operating_system.rb:10: warning: constant Gem::ConfigMap is deprecated
    Fetching gem metadata from https://rubygems.org/. Resolving dependencies...
    ...

Maybe also relevant: when running bundle install with a new package in a gemfile I get asked if I want to install with superuser privileges. I'm not sure what to choose and why, don't know enough about the implications.

The same errors pop up with gems:

    $ gem env version /usr/lib/ruby/vendor_ruby/rubygems/defaults/operating_system.rb:10: warning: constant Gem::ConfigMap is deprecated
    /usr/lib/ruby/vendor_ruby/rubygems/defaults/operating_system.rb:10: warning: constant Gem::ConfigMap is deprecated
    /usr/lib/ruby/vendor_ruby/rubygems/defaults/operating_system.rb:29: warning: constant Gem::ConfigMap is deprecated
    /usr/lib/ruby/vendor_ruby/rubygems/defaults/operating_system.rb:30: warning: constant Gem::ConfigMap is deprecated
    /usr/lib/ruby/vendor_ruby/rubygems/defaults/operating_system.rb:10: warning: constant Gem::ConfigMap is deprecated
    /usr/lib/ruby/vendor_ruby/rubygems/defaults/operating_system.rb:10: warning: constant Gem::ConfigMap is deprecated
    /usr/lib/ruby/vendor_ruby/rubygems/defaults/operating_system.rb:10: warning: constant Gem::ConfigMap is deprecated
    /usr/lib/ruby/vendor_ruby/rubygems/defaults/operating_system.rb:29: warning: constant Gem::ConfigMap is deprecated
    /usr/lib/ruby/vendor_ruby/rubygems/defaults/operating_system.rb:30: warning: constant Gem::ConfigMap is deprecated
    3.1.4

According to the [bug report](https://github.com/rubygems/rubygems/issues/3068) linked above, this is because at some point I mixed apt packages and bundler packages and / or then upgraded bundler. The thing is: the bug report is too technical for me to understand how to resolve this and I don't want to abuse it as a support forum. Concerning that: if this is not the appropriate place for my issue, I appreciate a pointer to the right place.

I'm running ubuntu 18.04 LTS, have su rights and apart from my jekyll project, don't depend on ruby. What should I do? Would appreciate some help, talking me through this!
## [5][Declarative and auditable object-oriented library for Redis](https://www.reddit.com/r/ruby/comments/hulcfy/declarative_and_auditable_objectoriented_library/)
- url: https://github.com/ntd251/sideroo?source=reddit
---

## [6][Rails 6.1 allows enums attributes to configure the default value](https://www.reddit.com/r/ruby/comments/hv39r6/rails_61_allows_enums_attributes_to_configure_the/)
- url: https://blog.bigbinary.com/2020/07/21/rails-6-1-allows-enums-attributes-to-have-default-value.html
---

## [7][Newest tty-prompt gem packs many new features: conversion of input into Array or Hash object, ability to silence output to allow for dynamic menus, a way to control help display (on_start/always/never) and few more.](https://www.reddit.com/r/ruby/comments/humiec/newest_ttyprompt_gem_packs_many_new_features/)
- url: https://github.com/piotrmurach/tty-prompt
---

## [8][Interesting or cool ruby features?](https://www.reddit.com/r/ruby/comments/huip5j/interesting_or_cool_ruby_features/)
- url: https://www.reddit.com/r/ruby/comments/huip5j/interesting_or_cool_ruby_features/
---
If you were to demo or explain what Ruby is to a bunch of non-ruby devs (e.g. Java or C# developers) what would you demonstrate or explain about Ruby which you think other devs might find interesting?  Any specific methods or facts that might impress a developer with little or no ruby knowledge?

Perhaps there is already a good website with some examples?

Thanks
## [9][Making RSpec Tests More Robust](https://www.reddit.com/r/ruby/comments/hukup6/making_rspec_tests_more_robust/)
- url: http://jakeyesbeck.com/2020/07/19/making-rspec-tests-more-robust/?src=reddit
---

## [10][Roadmap Ruby and Ruby On Rails](https://www.reddit.com/r/ruby/comments/huf714/roadmap_ruby_and_ruby_on_rails/)
- url: https://www.reddit.com/r/ruby/comments/huf714/roadmap_ruby_and_ruby_on_rails/
---
You know of a roadmap (similar [https://roadmap.sh/roadmaps](https://roadmap.sh/roadmaps) [https://www.quora.com/What-is-Ruby-On-Rails-framework-learning-road-map](https://www.quora.com/What-is-Ruby-On-Rails-framework-learning-road-map)) but for Ruby or Ruby On Rails, I'd like to create one for Ruby and another for Ruby On Rails, but first I'd like to know if you know of one that already exists and thus extend it or develop it as a learning path that can help novice programmers in their learning. Thanks for your input!
