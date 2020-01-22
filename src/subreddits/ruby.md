# ruby
## [1][Hey guys... I made something I’m pretty proud of!! It is a game called Startup Simulator that is completely in the command line. The only gem needed is colorize. Please feel free to clone and play the game. It’s similar to lemonade stand but you build a startup! Please Star my repo! It’s super fun!](https://www.reddit.com/r/ruby/comments/es4jlv/hey_guys_i_made_something_im_pretty_proud_of_it/)
- url: https://github.com/arichards4814/startup-simulator
---

## [2][Been learning Ruby for 6 months, just built my first side project and wanted to share :)](https://www.reddit.com/r/ruby/comments/es9umh/been_learning_ruby_for_6_months_just_built_my/)
- url: https://www.reddit.com/r/ruby/comments/es9umh/been_learning_ruby_for_6_months_just_built_my/
---
Hey guys,

Here is the [side project](http://www.careerfair.io/).

[Career Fair](https://www.careerfair.io/) is a website that lets you learn about jobs by people who have done them.

Too often, we only think about salary whilst taking a job -- I think there are many more factors that matter and this is my effort in adding transparency to that.

For some context: wrote my first line of code 6 months ago and immediately took a liking to the simplicity of ruby. 

Built using Ruby (Sinatra Framework), HTML, CSS, Javascript, and PostgreSQL on the backend. Site is responsive but the menu is currently not working on smartphones - bear with me whilst I fix that :)

Was lots of fun to code - theme is built from scratch. Sinatra is a great framework to start with - made it really easy to understand networking basics (i.e GET / POST requests etc) and how to persist a database. 

Would love your feedback and let me what you think.
## [3][Rails 6.1 adds query method missing to find orphan records](https://www.reddit.com/r/ruby/comments/esasfy/rails_61_adds_query_method_missing_to_find_orphan/)
- url: https://blog.saeloun.com/2020/01/21/rails-6-1-adds-query-method-missing-to-find-orphan-records
---

## [4][Cannot require highline/import even though successfully installed](https://www.reddit.com/r/ruby/comments/es9qfd/cannot_require_highlineimport_even_though/)
- url: https://www.reddit.com/r/ruby/comments/es9qfd/cannot_require_highlineimport_even_though/
---
Hey all. I have run into a strange problem where I get the below error message, when trying to run a script which requires highline/import, even after installing highline/import using **sudo gem install highline:**

    Traceback (most recent call last):
    	2: from ./myscript.rb:16:in `&lt;main&gt;'
    	1: from /home/user/.rbenv/versions/2.7.0/lib/ruby/2.7.0/rubygems/core_ext/kernel_require.rb:92:in `require'
    /home/user/.rbenv/versions/2.7.0/lib/ruby/2.7.0/rubygems/core_ext/kernel_require.rb:92:in `require': cannot load such file -- highline/import (LoadError) 

And this is the output when installing highline:

    # sudo gem install highline
    Successfully installed highline-2.0.3
    Parsing documentation for highline-2.0.3
    Done installing documentation for highline after 0 seconds
    1 gem installed

Why is it still reporting that highline could not be loaded even though it is successfully installed? I am running Ruby 2.7.0 using rbenv. I also tried running the below commands before I installed highline but it did not work:

    # rbenv shell 7.2.0
    # rbenv global 7.2.0
## [5][Rails is Fast: Optimize Your View Performance](https://www.reddit.com/r/ruby/comments/esbysh/rails_is_fast_optimize_your_view_performance/)
- url: https://blog.appsignal.com/2020/01/22/rails-is-fast-optimize-your-view-performance.html
---

## [6][Can beginners give talks at meetups like Philly.rb or NYC.rb?](https://www.reddit.com/r/ruby/comments/es116g/can_beginners_give_talks_at_meetups_like_phillyrb/)
- url: https://www.reddit.com/r/ruby/comments/es116g/can_beginners_give_talks_at_meetups_like_phillyrb/
---
If I were to give a talk it would definitely be a short one. I'm not sure (yet) what I would give a talk about but If I make a gem or something I would love to talk about it!

I'm also afraid of being too beginner to do a talk.
## [7][The Complete Guide to Migrate to Strong Parameters](https://www.reddit.com/r/ruby/comments/es100w/the_complete_guide_to_migrate_to_strong_parameters/)
- url: https://www.fastruby.io/blog/rails/upgrades/strong-parameters-migration-guide.html
---

## [8][RClone: a small "git clone" shortcut](https://www.reddit.com/r/ruby/comments/eryssu/rclone_a_small_git_clone_shortcut/)
- url: https://www.reddit.com/r/ruby/comments/eryssu/rclone_a_small_git_clone_shortcut/
---
Typing the full link to a GitHub repository is annoying, right?  
Well fear not, RClone is here for you! [**https://github.com/Sevodric/rclone**](https://github.com/Sevodric/rclone)

I'm a beginner in both git and Ruby and I made this small script to save the time I... lost... writing it... I guess...

Edit: thanks for interesting feedbacks, I applied some of them and updated in consequences.
## [9][Error installing ruby-debug-ide](https://www.reddit.com/r/ruby/comments/erwfqx/error_installing_rubydebugide/)
- url: https://www.reddit.com/r/ruby/comments/erwfqx/error_installing_rubydebugide/
---
Hello. I am using ruby 2.5.5p157. I am experiencing an issue when trying to install ruby-debug-ide on Kali Linux 2019.3 via: `gem install ruby-debug-ide`. I get the below output:

    Building native extensions. This could take a while...
    ERROR:  Error installing ruby-debug-ide:
    	ERROR: Failed to build gem native extension.
    
        current directory: /var/lib/gems/2.5.0/gems/ruby-debug-ide-0.7.0/ext
    /usr/bin/ruby2.5 mkrf_conf.rb
    Installing base gem
    Building native extensions. This could take a while...
    
    current directory: /var/lib/gems/2.5.0/gems/ruby-debug-ide-0.7.0/ext
    /usr/bin/ruby2.5 -rrubygems /usr/share/rubygems-integration/all/gems/rake-12.3.1/exe/rake RUBYARCHDIR=/var/lib/gems/2.5.0/extensions/x86_64-linux/2.5.0/ruby-debug-ide-0.7.0 RUBYLIBDIR=/var/lib/gems/2.5.0/extensions/x86_64-linux/2.5.0/ruby-debug-ide-0.7.0
    /usr/bin/ruby2.5: No such file or directory -- /usr/share/rubygems-integration/all/gems/rake-12.3.1/exe/rake (LoadError)
    
    rake failed, exit code 1
    

It seems that the issue is with this path not existing: `/usr/share/rubygems-integration/all/gems/rake-12.3.1/exe/rake`. However, if I run `rake --version`, I get `rake, version 12.3.1`, so rake is installed. I have already installed the prerequisite debase gem.

I also installed ruby-dev via `apt install ruby-dev`, but this didn't help. Are there any additional packages that I need to install to get it to work?
## [10][Only 15% of the Basecamp operations budget is spent on Ruby](https://www.reddit.com/r/ruby/comments/erj73x/only_15_of_the_basecamp_operations_budget_is/)
- url: https://m.signalvnoise.com/only-15-of-the-basecamp-operations-budget-is-spent-on-ruby/
---

