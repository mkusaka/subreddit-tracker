# ruby
## [1][Curious how the Ruby garbage collector works? I've written an article all about it](https://www.reddit.com/r/ruby/comments/ify6nu/curious_how_the_ruby_garbage_collector_works_ive/)
- url: https://blog.peterzhu.ca/notes-on-ruby-gc/
---

## [2][Can't figure out why my app doesn't scale as I add instances](https://www.reddit.com/r/ruby/comments/igby63/cant_figure_out_why_my_app_doesnt_scale_as_i_add/)
- url: https://www.reddit.com/r/ruby/comments/igby63/cant_figure_out_why_my_app_doesnt_scale_as_i_add/
---
Hi all. I am having a weird problem and hope that someone can give me some advice. I have a Kubernetes cluster for my Rails app, a CMS that lets users create their own sites. So there is the admin app itself, and then user sites are made of templates with Liquid which are stored in the database (Postgres); pages for user sites are cached in Memcached by Rails using the full page caching gem; then there is a small middleware at the top of the chain that checks if the requested page is in cache, if it is, it returns that response quickly and halts the chain so the rest of the Rails stack is not used and the request is a lot faster. When I add more instances, I can see from benchmarks that throughput for user sites scales more or less as I would expect, kinda of linearly up to a point. However if I benchmark the admin pages adding more instances doesn't make much difference to the req/sec figures, which is super weird. These pages aren't fully cached like user pages because content is dynamic and requires authentication and personalised content for each user, so in this case I am just doing fragment caching where possible. The other significant difference is that the admin pages use the database of course for each page with a varying number of queries per page. I have been testing things with the database setup and the storage in use to see if there are bottlenecks there, but I can't figure out the problem. The database cluster is one master and two read replicas (I am using Rails 6' automatic connection switching to split reads and writes between replicas and master), and has enough memory for a few 100s of connections; the block storage gives 7500 IOPS both read and write, and the dataset is tiny since I haven't launched yet, so it should be all in memory meaning that the storage wouldn't affect the performance much at the moment. 
Taking these things into account, I wouldn't think that the database is the bottleneck but I have done an experiment which may suggest otherwise: I have added a simple query to the aforementioned middleware and disabled the memcached part. So what the middleware does now is perform a simple database query with ActiveRecord and return a dummy response so the rest of the Rails app is not in use. To my surprise, even with that stupid simple query on a tiny table, I get 2 to 4 times more requests per second when I remove that query from the middleware, compared to the middleware with the database query. WTH? How can a so small query on a small table make such a difference? 
What can I try next to see if I can figure out the problem? How do I investigate if the database setup is the bottleneck? The other weird thing is that I have tried benchmarking Postgres against the master only, and I get between 500 and 800+ transactions per second. Not a huge number, but still shows that my app should handle more than 70-80 req/sec with 9 instances, also considering that the app can use the replicas too for the reads. The app is using Puma with two workers and 5 threads per instance. Any advice on how to investigate and find the problem would be MUCH appreciated. Thanks in advance!
## [3][The great Rubykon Benchmark 2020: CRuby vs JRuby vs TruffleRuby](https://www.reddit.com/r/ruby/comments/ifsjwd/the_great_rubykon_benchmark_2020_cruby_vs_jruby/)
- url: https://pragtob.wordpress.com/2020/08/24/the-great-rubykon-benchmark-2020-cruby-vs-jruby-vs-truffleruby/
---

## [4][Ruby TRICKS of 2018](https://www.reddit.com/r/ruby/comments/ifs5p2/ruby_tricks_of_2018/)
- url: https://idiosyncratic-ruby.com/75-ruby-tricks-of-2018.html
---

## [5][[Gem] active_record-events: Timestamp management made easy](https://www.reddit.com/r/ruby/comments/ifr6r4/gem_active_recordevents_timestamp_management_made/)
- url: https://www.reddit.com/r/ruby/comments/ifr6r4/gem_active_recordevents_timestamp_management_made/
---
Hello everyone! 👋

If you tend to use a lot of timestamp fields in your Rails applications, you might be interested in a gem I recently released called [active\_record-events](https://github.com/pienkowb/active_record-events).

By adding convenience methods on top of a `datetime` field, the gem allows you to manage custom timestamps in a similar way to how ActiveRecord handles the `created_at` and `updated_at` fields.

Here's an example. Assuming we have a `Task` model with a `completed_at` field, let's add a `has_event` macro inside the model class:

    class Task &lt; ActiveRecord::Base
      has_event :complete
    end

As a result, we get a bunch of useful methods without the need to define them explicitly.

    task = Task.create!
    
    task.completed? # =&gt; false
    
    task.complete
    
    task.completed? # =&gt; true
    task.completed_at # =&gt; Sat, 22 Aug 2020 20:45:39 UTC +00:00

This and plenty of other features are described in detail on the gem's GitHub page.

You can check it out here: [https://github.com/pienkowb/active\_record-events](https://github.com/pienkowb/active_record-events)
## [6][Understanding Ruby blocks](https://www.reddit.com/r/ruby/comments/ifqr28/understanding_ruby_blocks/)
- url: https://www.reddit.com/r/ruby/comments/ifqr28/understanding_ruby_blocks/
---
I know plenty has already been written about Ruby blocks but I decided to take a swing at writing an even better explanation of blocks than the other stuff that's out there right now.

This post is adapted from an exercise from a class I taught last year. It takes you through the process writing your own Ruby block, a block that will output any text you give it inside an ASCII "box".

Hope you like it. Here's the post: [Understanding Ruby blocks](https://www.codewithjason.com/understanding-ruby-blocks/)
## [7][New Gem simple_assets Released!](https://www.reddit.com/r/ruby/comments/ifwvkz/new_gem_simple_assets_released/)
- url: https://www.reddit.com/r/ruby/comments/ifwvkz/new_gem_simple_assets_released/
---
I have just released a new gem called simple_assets. It is a Dead simple HTML-based assets helper for Ruby. The main idea here is to promote re-usability for projects. I would appreciate any feedback on this design. https://github.com/westonganger/simple_assets
## [8][[JOB] [REMOTE] Drizly is Hiring Remote Software Engineers](https://www.reddit.com/r/ruby/comments/ifyb37/job_remote_drizly_is_hiring_remote_software/)
- url: https://www.reddit.com/r/ruby/comments/ifyb37/job_remote_drizly_is_hiring_remote_software/
---
Hello everyone!  


Cant seem to find the community rules page, so definitely attack me if I am wrong, i probably deserve it.  


So referencing my title, we here at Drizly is looking for Remote Fullstack (can favor front end or backend!) Software Engineers to help build out our eCommerce platform solution for retailers, consumers and alcohol brands across the nation. We are utilizing Ruby on the backend but we are quite language agnostic and happy to teach! We are open to non traditional backgrounds in both experience and education.    


Unfortunately at this time we are unable to sponsor at this time and candidates need to be based in the US.    


Feel free to apply here, email me [kenneth.han@drizly.com](mailto:kenneth.han@drizly.com) or DM me! [https://jobs.lever.co/drizly/25edf13d-923d-48de-b99c-10106873ac27?lever-origin=applied&amp;lever-source%5B%5D=reddit](https://jobs.lever.co/drizly/25edf13d-923d-48de-b99c-10106873ac27?lever-origin=applied&amp;lever-source%5B%5D=reddit)
## [9][sports gem - sport data structures for matches, scores, leagues, seasons, rounds, groups, teams, clubs and more](https://www.reddit.com/r/ruby/comments/ifkkb2/sports_gem_sport_data_structures_for_matches/)
- url: https://github.com/sportdb/sport.db/tree/master/sports
---

## [10][Exploring Ruby's Enumerable Module](https://www.reddit.com/r/ruby/comments/ifiybu/exploring_rubys_enumerable_module/)
- url: https://www.sandipmane.dev/exploring-rubys-enumerable-module
---

