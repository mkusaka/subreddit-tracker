# ruby
## [1][How to run slow RSpec files on Github Actions with parallel jobs by doing an auto split of the spec file by test examples](https://www.reddit.com/r/ruby/comments/h9beb4/how_to_run_slow_rspec_files_on_github_actions/)
- url: https://docs.knapsackpro.com/2020/how-to-run-slow-rspec-files-on-github-actions-with-parallel-jobs-by-doing-an-auto-split-of-the-spec-file-by-test-examples
---

## [2][datoji - Remote JSON storage server. Create, Read, Update, Delete and Search JSON data.](https://www.reddit.com/r/ruby/comments/h9fvqs/datoji_remote_json_storage_server_create_read/)
- url: https://datoji.dev
---

## [3][Noob ruby/Linux cron job question](https://www.reddit.com/r/ruby/comments/h93lk9/noob_rubylinux_cron_job_question/)
- url: https://www.reddit.com/r/ruby/comments/h93lk9/noob_rubylinux_cron_job_question/
---
Thanks for reading this. I have very little experience with Ruby.  I have a RedHat Linux server that has a homemade Ruby application running on it. The person who created it is long gone. It's started manually now, but I'd like it to start on reboot. I was going to add a cron job for the script, but it's started with the command "rvmsudo bundle exec puma". Is adding this to a "@reboot" cron job the best path here, and if so do I just put that path and command in a script?

&amp;#x200B;

Thanks again, any help is appreciated.
## [4][I made a template for running a ruby app (sinatra) in Docker, feedback wanted](https://www.reddit.com/r/ruby/comments/h8qmr3/i_made_a_template_for_running_a_ruby_app_sinatra/)
- url: https://github.com/SebastianThorn/ruby-docker-skeleton
---

## [5][Puma NewRelic plugin to monitor pools](https://www.reddit.com/r/ruby/comments/h8rsgh/puma_newrelic_plugin_to_monitor_pools/)
- url: https://github.com/benoist/puma-newrelic
---

## [6][Good (can be paid) courses for learning Ruby + Rails for PHP developer](https://www.reddit.com/r/ruby/comments/h7p6tm/good_can_be_paid_courses_for_learning_ruby_rails/)
- url: https://www.reddit.com/r/ruby/comments/h7p6tm/good_can_be_paid_courses_for_learning_ruby_rails/
---
Hello! I am a PHP developer with few years experience and I want to switch to ruby and get a job as a junior ROR dev. I was hoping you could point me to some good and tested tutorials or courses where the focus is put on creating a real-life application.
## [7][Reverting to an older string time format](https://www.reddit.com/r/ruby/comments/h7pq7v/reverting_to_an_older_string_time_format/)
- url: https://www.reddit.com/r/ruby/comments/h7pq7v/reverting_to_an_older_string_time_format/
---
Currently with ruby 2.7.1, when you do something like `Date.today.to_s(:long)` you'll get something like

`June 09, 2020` 

But with older versions of ruby (for instance 2.2.4), that format would look like this: 

`June 9, 2020`

Without that 0 padding on the day. Is there anyway to revert back to this old formatting, i.e **without** the 0 padding on the day? Without changing the `to_s(:long)` part and without going back to an older  version of ruby? 

Or is there anyway to specify that `:long` should follow this structure: `%B %e, %Y`

All help is appreciated!
## [8][Is my algorithm's time complexity O(n log n)?](https://www.reddit.com/r/ruby/comments/h7s4yg/is_my_algorithms_time_complexity_on_log_n/)
- url: https://www.reddit.com/r/ruby/comments/h7s4yg/is_my_algorithms_time_complexity_on_log_n/
---
    def two_sum?(array, value)
    	array.sort! # O(nlogn)
    	array.each do |element|
    		return true if bsearch(array - [element], value - element) == true
    	end
    	return false
    end
    
    def bsearch(array, value)
    	return false if array.empty?
    	mid_idx = array.length / 2
    	mid_value = array[mid_idx]
    	return true if mid_value == value
    	mid_value &lt; value ? bsearch(array[0...mid_idx], value) : bsearch(array[mid_idx+1..-1], value)
    end

I'm trying to create a function that finds a two unique numbers in an array whose sum equals the value in the second argument. I believe my implementation has a time complexity of O(n log n). However, when I run it with another function whose time complexity is also O(n log n), the total time is way different (calculated using the Benchmark gem) using the same input. For my function, it takes about 0.9 seconds. For the other function, it is taking 0.003 seconds. Is there any error in my algorithm analysis? Is my implementation not O(n log n)?
## [9][Spark Joy by Running Fewer Tests](https://www.reddit.com/r/ruby/comments/h7cclo/spark_joy_by_running_fewer_tests/)
- url: https://engineering.shopify.com/blogs/engineering/spark-joy-by-running-fewer-tests
---

## [10][Fun facts about Ruby #4](https://www.reddit.com/r/ruby/comments/h12yg6/fun_facts_about_ruby_4/)
- url: https://i.redd.it/n22dyxmq8b451.png
---

