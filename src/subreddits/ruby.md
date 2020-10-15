# ruby
## [1][How to create a cross-plathorm SDK library?](https://www.reddit.com/r/ruby/comments/jbl93u/how_to_create_a_crossplathorm_sdk_library/)
- url: https://www.reddit.com/r/ruby/comments/jbl93u/how_to_create_a_crossplathorm_sdk_library/
---
Let's say there's a library written in C++ for which I want to create an SDK or gem in Ruby.

&amp;#x200B;

A C++ library provides 3 plathorm-specific pre-compiled libraries:

\* windows, \*.dll

\* linux, \*.so

\* macOS, \*.dylib

&amp;#x200B;

I'll use FFI gem and refer to a particular type of a precompiled library, from my Ruby SDK/gem, depending on a plathorm of the end user.

&amp;#x200B;

(1) What's a conventional way to do it? Where, how is it done?

Namely:

     main_precompiled_lib = 
      if plathorm == Windows
        "./win-lib.dll"
      elseif plathorm == Linux
        "./linux-lib.so"
      elseif ....

(2) How can I provide a way to the user to choose his plathorm?
## [2][Best intermediate - expert training tools and resources?](https://www.reddit.com/r/ruby/comments/jbapcq/best_intermediate_expert_training_tools_and/)
- url: https://www.reddit.com/r/ruby/comments/jbapcq/best_intermediate_expert_training_tools_and/
---
Been using rails and ruby for a while now I'm fairly competent but like to step my game up to the next level with the quality of apps I create. Has anyone found any or have any favourite resources for doing so?
## [3][OpenStruct in Ruby](https://www.reddit.com/r/ruby/comments/jbig5f/openstruct_in_ruby/)
- url: https://medium.com/rubycademy/openstruct-in-ruby-ab6ba3aff9a4
---

## [4][Dotenv-linter v2.2.0: find and fix problems in .env files. What's new?](https://www.reddit.com/r/ruby/comments/jaz0th/dotenvlinter_v220_find_and_fix_problems_in_env/)
- url: https://evrone.com/dotenv-linter-v220
---

## [5][Hijack-test gem to test device for clipboard hijackers](https://www.reddit.com/r/ruby/comments/jb1fqk/hijacktest_gem_to_test_device_for_clipboard/)
- url: https://cybersecrs.github.io/project/hijack-test
---

## [6][Rails multitenancy story in 11 snippets of code](https://www.reddit.com/r/ruby/comments/jaz3x9/rails_multitenancy_story_in_11_snippets_of_code/)
- url: https://blog.arkency.com/rails-multitenancy-story-in-11-snippets-of-code/
---

## [7][Calling a Rust function from Ruby](https://www.reddit.com/r/ruby/comments/javslp/calling_a_rust_function_from_ruby/)
- url: https://www.reddit.com/r/ruby/comments/javslp/calling_a_rust_function_from_ruby/
---
How to call a Rust function from Ruby?

Are there gems that are ready for production the most?
## [8][Ruby Type Signatures and Typing at Stripe](https://www.reddit.com/r/ruby/comments/jaldo8/ruby_type_signatures_and_typing_at_stripe/)
- url: https://brandur.org/nanoglyphs/015-ruby-typing#ruby-typing
---

## [9][Font Awesome on Rails 6 Application using Yarn](https://www.reddit.com/r/ruby/comments/jasnu8/font_awesome_on_rails_6_application_using_yarn/)
- url: https://www.youtube.com/watch?v=7ybtzjpNw44
---

## [10][Ruby 2.7.2 has revised Rdoc for ENV](https://www.reddit.com/r/ruby/comments/jairwg/ruby_272_has_revised_rdoc_for_env/)
- url: https://www.reddit.com/r/ruby/comments/jairwg/ruby_272_has_revised_rdoc_for_env/
---
Some of the revisions I've made were in Ruby 2.7.0.  The rest are now in Ruby 2.7.2.

\- 2.6.3 (before revisions): [https://ruby-doc.org/core-2.6.3/ENV.html](https://ruby-doc.org/core-2.6.3/ENV.html)

\- 2.7.2 (with all revisions): [https://ruby-doc.org/core-2.7.2/ENV.html](https://ruby-doc.org/core-2.7.2/ENV.html)
