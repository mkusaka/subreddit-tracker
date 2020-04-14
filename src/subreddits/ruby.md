# ruby
## [1][Ruby Concurrency Final Report - not the end, just the beginning!](https://www.reddit.com/r/ruby/comments/g12yai/ruby_concurrency_final_report_not_the_end_just/)
- url: https://www.codeotaku.com/journal/2020-04/ruby-concurrency-final-report/index
---

## [2][Using named captures to extract information from Strings - Aliou Diallo](https://www.reddit.com/r/ruby/comments/g13ud1/using_named_captures_to_extract_information_from/)
- url: https://aliou.me/posts/2020/04/regexp-named-captures/
---

## [3][How to work with external services in tests](https://www.reddit.com/r/ruby/comments/g13hvi/how_to_work_with_external_services_in_tests/)
- url: https://jtway.co/how-to-work-with-external-services-in-tests-f06050343aeb?source=friends_link&amp;sk=76776e03d94c5443dad5bd6da8f44b36
---

## [4][Ruby 2.7 deprecates conversion of keyword arguments](https://www.reddit.com/r/ruby/comments/g143it/ruby_27_deprecates_conversion_of_keyword_arguments/)
- url: https://blog.bigbinary.com/2020/04/14/ruby-2-7-deprecates-conversion-of-keyword-arguments.html
---

## [5][Feature Flags: The stupid simple way to de-stress production releases](https://www.reddit.com/r/ruby/comments/g0ivt2/feature_flags_the_stupid_simple_way_to_destress/)
- url: https://boringrails.com/articles/feature-flags-simplest-thing-that-could-work/
---

## [6][I created a gem for creating auto-similar fractals](https://www.reddit.com/r/ruby/comments/g0grqs/i_created_a_gem_for_creating_autosimilar_fractals/)
- url: https://github.com/StefanoMartin/FractalsRB
---

## [7][DudePolicy - gem for policy objects in Ruby on Rails application from perspective of current_user/current_account](https://www.reddit.com/r/ruby/comments/g0g3rj/dudepolicy_gem_for_policy_objects_in_ruby_on/)
- url: https://github.com/equivalent/dude_policy
---

## [8][Shoutouts to dgutov for creating robe.el](https://www.reddit.com/r/ruby/comments/g0lpxt/shoutouts_to_dgutov_for_creating_robeel/)
- url: https://www.reddit.com/r/ruby/comments/g0lpxt/shoutouts_to_dgutov_for_creating_robeel/
---
After finishing school I managed to get a job as a Ruby on Rails developer without having any prior experience with Ruby. Setting up a decent development environment caused some difficulties, though. I wanted to use Emacs but Solargraph didn’t seem to work too well. The examples on the website seemed to work, I got code completion for classes, class methods and functions from included modules. However, everything related to objects and local variables completely failed. To make sure the problem wasn’t caused by Emacs, I tried it out in VSCode. Same result.

Later I switched to Doom Emacs which uses Robe by default for Ruby development. It works pretty well and seems to give a quite accurate list of all the available methods. Because Robe wasn't really mentioned often on this subreddit, I decided to spread the word a bit.

Thanks for creating this awesome package!
## [9][RVM/Bundler : gem uninstall -aIX fails. -- Why are there two different directories for gems that are being used?](https://www.reddit.com/r/ruby/comments/g0or5c/rvmbundler_gem_uninstall_aix_fails_why_are_there/)
- url: https://www.reddit.com/r/ruby/comments/g0or5c/rvmbundler_gem_uninstall_aix_fails_why_are_there/
---
Hi everyone,

I'm gonna start with telling you what is the goal.

**Goal:**

1. Uninstalling everything related to rvm, bundler, and gems. Literally making things gone.
2. Then re-installing everything back again as if this is a new fresh server.
3. Installing gems after becoming a correct system user through bundler to avoid any permission related errors on any ruby on rails app.
4. This should work on any system regardless of how f\*\*\*ed up the rvm/ruby/gem configuration.

**The problem:**

Hi, gem uninstall -aIX fails (Goal #1) because turns out the gems are installed in different directories etc.

And the error log tells me to go to a different directory and run the commands over there.

$GEM\_PATH: /usr/local/rvm/gems/ruby-2.6.3:/usr/local/rvm/gems/ruby-2.6.3@global

$GEM\_HOME: /usr/local/rvm/gems/ruby-2.6.3

**Error log:**

**$ gem uninstall -aIX**

ERROR:  While executing gem ... (Gem::InstallError)

bundler-unload is not installed in GEM\_HOME, try:

gem uninstall -i /usr/local/rvm/rubies/ruby-2.6.3/lib/ruby/gems/2.6.0 bundler-unload

**The question:**

Why the second path exists for some of the gems? (and why the second directory ends with 2.6.0)

\- /usr/local/rvm/gems/ruby-2.6.3/gems/

\- /usr/local/rvm/rubies/ruby-2.6.3/lib/ruby/gems/2.6.0
## [10][Improve code readability with closures in Ruby](https://www.reddit.com/r/ruby/comments/g0flu7/improve_code_readability_with_closures_in_ruby/)
- url: https://railsguides.net/improve-code-readability-with-closures-in-ruby/
---

