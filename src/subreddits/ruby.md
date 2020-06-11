# ruby
## [1][DragonRuby in itch.io Bundle for Racial Justice and Equality ($5+)](https://www.reddit.com/r/ruby/comments/h0jp7v/dragonruby_in_itchio_bundle_for_racial_justice/)
- url: https://itch.io/b/520/bundle-for-racial-justice-and-equality
---

## [2][Ruby Desktop Development - Complete OSS Mac DMG App](https://www.reddit.com/r/ruby/comments/h0lnai/ruby_desktop_development_complete_oss_mac_dmg_app/)
- url: https://www.reddit.com/r/ruby/comments/h0lnai/ruby_desktop_development_complete_oss_mac_dmg_app/
---
A Ruby desktop complete application developed with [Glimmer](https://github.com/AndyObtiva/glimmer) has been released at version 2.1 (open-source and free):

[https://github.com/AndyObtiva/MathBowling](https://github.com/AndyObtiva/MathBowling)

Please use as a learning reference for how to leverage the [Glimmer](https://github.com/AndyObtiva/glimmer) desktop development GUI library.

Enjoy!!!
## [3][Help me to understand the triple equal with module](https://www.reddit.com/r/ruby/comments/h0s4gx/help_me_to_understand_the_triple_equal_with_module/)
- url: https://www.reddit.com/r/ruby/comments/h0s4gx/help_me_to_understand_the_triple_equal_with_module/
---
&amp;#x200B;

https://preview.redd.it/qyjc4bz6m7451.png?width=1058&amp;format=png&amp;auto=webp&amp;s=afbcc7dc64239ae644fee353e0a7171b195f9c95

Why `@controller.class.parent === Recall` doesn't returns `true`?

The `object_id` are the same but it is not equal. Why?

`Recall` is a module.
## [4][Expected output of the debug flag (`-d` or `--debug`)](https://www.reddit.com/r/ruby/comments/h0vor5/expected_output_of_the_debug_flag_d_or_debug/)
- url: https://www.reddit.com/r/ruby/comments/h0vor5/expected_output_of_the_debug_flag_d_or_debug/
---
On a fresh install of macOS 10.14.6, opening the Interactive Ruby Shell with the debug flag enabled results in the following output:

```console
$ ruby --debug
Exception `LoadError' at /System/Library/Frameworks/Ruby.framework/Versions/2.3/usr/lib/ruby/2.3.0/rubygems.rb:1242 - cannot load such file -- rubygems/defaults/operating_system
Exception `LoadError' at /System/Library/Frameworks/Ruby.framework/Versions/2.3/usr/lib/ruby/2.3.0/rubygems.rb:1251 - cannot load such file -- rubygems/defaults/ruby
Exception `LoadError' at /System/Library/Frameworks/Ruby.framework/Versions/2.3/usr/lib/ruby/2.3.0/rubygems/core_ext/kernel_require.rb:55 - cannot load such file -- did_you_mean
```

For what it's worth, similar output is shown when running a Ruby installation installed via `rbenv`.

I did not expect to see any exceptions. The closest thing I could find relating to the output I saw on GitHub was [this issue](https://github.com/rubygems/rubygems/issues/358) in the rubygems repository on GitHub. A contributor closed the issue with [this](https://github.com/rubygems/rubygems/issues/358#issuecomment-7114343) to say:

&gt; This is expected output when you run with `ruby -d`. There's nothing wrong with it.

Why is it the expected output?

I also came across [this Stack Overflow post](https://stackoverflow.com/questions/35240324/something-wrong-with-my-ruby?noredirect=1#comment58199906_35240324). To quote [the accepted answer](https://stackoverflow.com/a/35240886/11805749):

&gt; The system ruby installation seems to be messed up. Try removing the offending `/Library/Ruby/Site/2.0.0/rubygems.rb` (and may be more of enclosing folders)
&gt; 
&gt; Ruby have its bundled rubygems at `/System/Library/Frameworks/Ruby.framework/Versions/2.0/usr/lib/ruby/2.0.0` so it should work with these.
&gt; 
&gt; Once ruby is alive - you can try upgrading rubygems again by `sudo gem update --system` if needed, but I advise setting up a ruby version manager like `rvm` (rvm.io) and leave the system ruby be there only for emergences and backing the `brew`.
## [5][David Heinemeier Hansson Interview](https://www.reddit.com/r/ruby/comments/h0l5q5/david_heinemeier_hansson_interview/)
- url: https://evrone.com/dhh-interview
---

## [6][Introducing TestBench: a principled test framework for Ruby :-)](https://www.reddit.com/r/ruby/comments/h0cdc7/introducing_testbench_a_principled_test_framework/)
- url: https://blog.eventide-project.org/articles/introducing-test-bench/
---

## [7][Rails add support for bulk insert/upsert on relation](https://www.reddit.com/r/ruby/comments/h0dk0v/rails_add_support_for_bulk_insertupsert_on/)
- url: https://blog.saeloun.com/2020/06/10/rails-support-bulk-insert-on-relation
---

## [8][Pallets, the Ruby workflow engine, just got a big performance update](https://www.reddit.com/r/ruby/comments/h05vh5/pallets_the_ruby_workflow_engine_just_got_a_big/)
- url: https://github.com/linkyndy/pallets
---

## [9][RDatasets - Ruby gem for loading many of the data sets available in R](https://www.reddit.com/r/ruby/comments/h07ro8/rdatasets_ruby_gem_for_loading_many_of_the_data/)
- url: https://github.com/kojix2/rdatasets
---

## [10][Exploring Metaprogramming in Ruby](https://www.reddit.com/r/ruby/comments/gzsf5y/exploring_metaprogramming_in_ruby/)
- url: https://www.halcyon.hr/posts/exploring-metaprogramming-in-ruby/
---

