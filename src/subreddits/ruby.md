# ruby
## [1][Rubygem release heat maps on the Ruby Toolbox](https://www.reddit.com/r/ruby/comments/emh8tr/rubygem_release_heat_maps_on_the_ruby_toolbox/)
- url: https://www.ruby-toolbox.com/blog/2020-01-09/rubygem-release-heatmaps
---

## [2][new app with no javascript and webpack](https://www.reddit.com/r/ruby/comments/emfip7/new_app_with_no_javascript_and_webpack/)
- url: https://www.reddit.com/r/ruby/comments/emfip7/new_app_with_no_javascript_and_webpack/
---
How does using the --skip-javascript and --skip-webpack-install to create a new app affect anything?
## [3][is there a tty: true equivalent for docker run?](https://www.reddit.com/r/ruby/comments/emo2o2/is_there_a_tty_true_equivalent_for_docker_run/)
- url: https://www.reddit.com/r/ruby/comments/emo2o2/is_there_a_tty_true_equivalent_for_docker_run/
---
I want to start a container in detached mode, then I can exec -it into it in my own pace. I'm currently able to do that via docker-compose with **tty: true** in the yml.

Is there a version without docker-compose, that is, I can run with the **docker run** command? And without a running service in the container. (I'm able to do that without a running service with the docker-compose method)
## [4][Deep dive into Did You Mean](https://www.reddit.com/r/ruby/comments/em9fdc/deep_dive_into_did_you_mean/)
- url: https://shime.sh/deep-dive-into-did-you-mean
---

## [5][Ruby low memory and cpu on file upload](https://www.reddit.com/r/ruby/comments/eme8e4/ruby_low_memory_and_cpu_on_file_upload/)
- url: https://www.reddit.com/r/ruby/comments/eme8e4/ruby_low_memory_and_cpu_on_file_upload/
---
I'm uploading pdf files to the cloud storage using fog-backblaze gem. I dig little in to the source code. And now I'm thinking of, Instead of ruby handle IO objects is there any better way to totally delegate this on to a better library like curl, by just passing file path. So MRI threads get little weight. Do you know any libraries to do that?
## [6][Acts_as_taggable_on and JSONAPI::Resources for react-rails app, how should I collect information to read for my front-end?](https://www.reddit.com/r/ruby/comments/emcls0/acts_as_taggable_on_and_jsonapiresources_for/)
- url: https://www.reddit.com/r/ruby/comments/emcls0/acts_as_taggable_on_and_jsonapiresources_for/
---
There is Rails code for example (from the github page) which is used to find information

    ActsAsTaggableOn::Tag.most_used
    User.tagged_with("awesome").by_join_date
    @tom.find_related_skills # =&gt; [&lt;User name="Bobby"&gt;, &lt;User name="Frankie"&gt;]

But I cannot for the life of me try and get them into JSONAPI data using the JSONAPI::Resources gem.

Should I be coding in reactJS to find similar information? I could simply just display the raw data of the tags. It is doable (through tags and taggings), but I don't want to repeat myself if any methods are available.

Is there a way to run rails code from a react app?
## [7][Is this a StringIO bug in ruby 2.7.0?](https://www.reddit.com/r/ruby/comments/emd6q4/is_this_a_stringio_bug_in_ruby_270/)
- url: https://www.reddit.com/r/ruby/comments/emd6q4/is_this_a_stringio_bug_in_ruby_270/
---
When running tests on one of my projects for 2.7.0, I ran into some weird failures, which I reduced/isolated to this. 

StringIO no longer preserves encoding of strings it was initialized with in 2.7.0. 

Is this a bug? I thought I'd ask for feedback here before trying to submit it as a bug on ruby tracker (and figure out how to do so). 

      binary_string = "\x11\x99".force_encoding("ASCII-8BIT") # note this is invalid UTF-8, for realism
      puts "binary_string.encoding: #{binary_string.encoding}" # "ASCII-8BIT"

      string_io = StringIO.new(binary_string)
      retrieved = string_io.read
      puts "retrieved.encoding: #{retrieved.encoding}"

      # On Ruby 2.7.0, UTF-8
      # On Ruby 2.6.5, ASCII-8BIT

**Oh except OOPS**, I can't reproduce this in a straight `irb` terminal. But it is reproducing in my actual project!

So it must not be StringIO by itself, it must be one of my dependencies rudely changing the behavior of StringIO somehow... this is painful. 

Does Rails monkey-patch StringIO somehow?  Confirmed I can reproduce this in a project that does "require 'rails'". 

I might go ahead and report this as a bug to rails then.... unless it's somehow intended or a bugfix? Very odd. Still curious if anyone has any advice. I think I may have trouble getting this report consideration on Rails tracker. I *really* don't understand what's going on.
## [8][What's new in Ruby 2.7](https://www.reddit.com/r/ruby/comments/em4keo/whats_new_in_ruby_27/)
- url: https://nithinbekal.com/posts/ruby-2-7/
---

## [9][Improving Database performance and overcoming common N+1 issues in Active Record using includes, preload, eager_load, pluck, select, exists?](https://www.reddit.com/r/ruby/comments/em5e4a/improving_database_performance_and_overcoming/)
- url: https://blog.saeloun.com/2020/01/08/activerecord-database-performance-n-1-includes-preload-eager-load-pluck
---

## [10][Open Source Progress Report](https://www.reddit.com/r/ruby/comments/em3f3v/open_source_progress_report/)
- url: https://www.codeotaku.com/journal/2019-12/open-source-progress-report/index
---

