# ruby
## [1][GitHub's Upgrade Journey to Ruby 2.7](https://www.reddit.com/r/ruby/comments/igimha/githubs_upgrade_journey_to_ruby_27/)
- url: https://github.blog/2020-08-25-upgrading-github-to-ruby-2-7/
---

## [2][Turn your Ruby on Rails REST API to GraphQL using Hasura Actions](https://www.reddit.com/r/ruby/comments/igwf09/turn_your_ruby_on_rails_rest_api_to_graphql_using/)
- url: https://hasura.io/blog/turn-your-ruby-on-rails-rest-api-to-graphql-using-hasura-actions/
---

## [3][Google Cloud Functions On Ruby](https://www.reddit.com/r/ruby/comments/igltnc/google_cloud_functions_on_ruby/)
- url: https://docs.google.com/forms/d/e/1FAIpQLSfEgsbch9pCO52W1fLFdmIibCLhB_eU7MTzZWX4E2PfOvwa2w/viewform
---

## [4][Can't figure out why my app doesn't scale as I add instances](https://www.reddit.com/r/ruby/comments/igby63/cant_figure_out_why_my_app_doesnt_scale_as_i_add/)
- url: https://www.reddit.com/r/ruby/comments/igby63/cant_figure_out_why_my_app_doesnt_scale_as_i_add/
---
Hi all. I am having a weird problem and hope that someone can give me some advice. 

I have a Kubernetes cluster for my Rails app, a CMS that lets users create their own sites. So there is the admin app itself, and then user sites are made of templates with Liquid which are stored in the database (Postgres); pages for user sites are cached in Memcached by Rails using the full page caching gem; then there is a small middleware at the top of the chain that checks if the requested page is in cache, if it is, it returns that response quickly and halts the chain so the rest of the Rails stack is not used and the request is a lot faster. 

When I add more instances, I can see from benchmarks that throughput for user sites scales more or less as I would expect, kinda of linearly up to a point. However if I benchmark the admin pages adding more instances doesn't make much difference to the req/sec figures, which is super weird. These pages aren't fully cached like user pages because content is dynamic and requires authentication and personalised content for each user, so in this case I am just doing fragment caching where possible. The other significant difference is that the admin pages use the database of course for each page with a varying number of queries per page. 

I have been testing things with the database setup and the storage in use to see if there are bottlenecks there, but I can't figure out the problem. The database cluster is one master and two read replicas (I am using Rails 6' automatic connection switching to split reads and writes between replicas and master), and has enough memory for a few 100s of connections; the block storage gives 7500 IOPS both read and write, and the dataset is tiny since I haven't launched yet, so it should be all in memory meaning that the storage wouldn't affect the performance much at the moment. 

Taking these things into account, I wouldn't think that the database is the bottleneck but I have done an experiment which may suggest otherwise: I have added a simple query to the aforementioned middleware and disabled the memcached part. So what the middleware does now is perform a simple database query with ActiveRecord and return a dummy response so the rest of the Rails app is not in use. To my surprise, even with that stupid simple query on a tiny table, I get 2 to 4 times more requests per second when I remove that query from the middleware, compared to the middleware with the database query. WTH? How can a so small query on a small table make such a difference? 

What can I try next to see if I can figure out the problem? How do I investigate if the database setup is the bottleneck? The other weird thing is that I have tried benchmarking Postgres against the master only, and I get between 500 and 800+ transactions per second. Not a huge number, but still shows that my app should handle more than 70-80 req/sec with 9 instances, also considering that the app can use the replicas too for the reads. The app is using Puma with two workers and 5 threads per instance. Any advice on how to investigate and find the problem would be MUCH appreciated. Thanks in advance!
## [5][How to add authorization to a Ruby app using oso, an open source policy engine](https://www.reddit.com/r/ruby/comments/igf2x8/how_to_add_authorization_to_a_ruby_app_using_oso/)
- url: https://www.reddit.com/r/ruby/comments/igf2x8/how_to_add_authorization_to_a_ruby_app_using_oso/
---
 

https://preview.redd.it/udhb44xm96j51.png?width=403&amp;format=png&amp;auto=webp&amp;s=87e8470d2a455bb01458763d8b40045b172b55e4

Hi all!

We built an open source policy engine for authorization that's embedded in your application, called oso. The idea is that you write policies using the oso policy language to govern who can do what inside your application, then you integrate them with a few lines of code using our Ruby library.

In the Rails world, there are some great precedents for authorization (e.g., Pundit and CanCanCan). We tried to build on these and pick up where they left off.  The big difference between them and oso is that the oso language itself (Polar) is purpose-built for authorization – e.g., representing common domain concepts and patterns like roles, hierarchies and relationships.

You can use oso in any Ruby application. We're planning to build a Rails integration soon, and would welcome any specific feature requests if you have them.

We're continuing to build on oso, so we're keen to hear any feedback and happy to answer any questions.

Some useful links:

Quickstart: [https://docs.osohq.com/getting-started/quickstart.html](https://docs.osohq.com/getting-started/quickstart.html)  
Ruby library docs: [https://docs.osohq.com/using/libraries/ruby/index.html](https://docs.osohq.com/using/libraries/ruby/index.html)  
Source code: [https://github.com/osohq/oso/tree/main/languages/ruby](https://github.com/osohq/oso/tree/main/languages/ruby)  
Feel free to join us on Slack for any questions: [join-slack.osohq.com](http://join-slack.osohq.com/)

PS We also support Python and Java, and are actively working on other languages.
## [6][Curious how the Ruby garbage collector works? I've written an article all about it](https://www.reddit.com/r/ruby/comments/ify6nu/curious_how_the_ruby_garbage_collector_works_ive/)
- url: https://blog.peterzhu.ca/notes-on-ruby-gc/
---

## [7][The great Rubykon Benchmark 2020: CRuby vs JRuby vs TruffleRuby](https://www.reddit.com/r/ruby/comments/ifsjwd/the_great_rubykon_benchmark_2020_cruby_vs_jruby/)
- url: https://pragtob.wordpress.com/2020/08/24/the-great-rubykon-benchmark-2020-cruby-vs-jruby-vs-truffleruby/
---

## [8][[Gem] active_record-events: Timestamp management made easy](https://www.reddit.com/r/ruby/comments/ifr6r4/gem_active_recordevents_timestamp_management_made/)
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
## [9][Ruby TRICKS of 2018](https://www.reddit.com/r/ruby/comments/ifs5p2/ruby_tricks_of_2018/)
- url: https://idiosyncratic-ruby.com/75-ruby-tricks-of-2018.html
---

## [10][Understanding Ruby blocks](https://www.reddit.com/r/ruby/comments/ifqr28/understanding_ruby_blocks/)
- url: https://www.reddit.com/r/ruby/comments/ifqr28/understanding_ruby_blocks/
---
I know plenty has already been written about Ruby blocks but I decided to take a swing at writing an even better explanation of blocks than the other stuff that's out there right now.

This post is adapted from an exercise from a class I taught last year. It takes you through the process writing your own Ruby block, a block that will output any text you give it inside an ASCII "box".

Hope you like it. Here's the post: [Understanding Ruby blocks](https://www.codewithjason.com/understanding-ruby-blocks/)
