# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/ezrfed/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/ezrfed/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Migrating from Rails 5 to Rails 6](https://www.reddit.com/r/rails/comments/f26c6v/migrating_from_rails_5_to_rails_6/)
- url: https://www.reddit.com/r/rails/comments/f26c6v/migrating_from_rails_5_to_rails_6/
---
I have a long running Rails 5 application with "a lot" of javascript files. It currently has webpack as well but not the default configuration used by Rails 6. Anyone who could recommend an article / tutorial / best way to migrate from sprockets asset pipeline Rails 5 to webpack based Rails 6?
## [3][Manage Redis by Ruby On Rails](https://www.reddit.com/r/rails/comments/f1x280/manage_redis_by_ruby_on_rails/)
- url: https://www.reddit.com/r/rails/comments/f1x280/manage_redis_by_ruby_on_rails/
---
I come to present a tool that I have developed with 2 others developers [https://github.com/OpenGems/redis\_web\_manager](https://github.com/OpenGems/redis_web_manager)

This tool allows you to manage Redis, thanks to a web interface you  can easily manage your Redis instance (see your keys, the memory used,  clients connected, etc …).

Enjoy !
## [4][How to speed up Ruby and JavaScript tests with CI parallelisation](https://www.reddit.com/r/rails/comments/f263lr/how_to_speed_up_ruby_and_javascript_tests_with_ci/)
- url: https://www.reddit.com/r/rails/comments/f263lr/how_to_speed_up_ruby_and_javascript_tests_with_ci/
---
Want to know how to do split tests on parallel continuous integration servers — static and dynamic? Follow the link below to the article: https://medium.com/arturtrzop/how-to-speed-up-ruby-and-javascript-tests-with-ci-parallelisation-a2324a62022a?source=friends_link&amp;sk=05c8bc7d416797aafe2d92fda90d6495
## [5][How to Efficiently Understand/Debug Rails Code for someone new to ORMs](https://www.reddit.com/r/rails/comments/f1yaam/how_to_efficiently_understanddebug_rails_code_for/)
- url: https://www.reddit.com/r/rails/comments/f1yaam/how_to_efficiently_understanddebug_rails_code_for/
---
Hello, all.

Semi-recently, I transitioned into a job that uses Ruby on Rails, and I'm really digging it so far!

I'd never used anything professionally that involved an ORM, so I find myself a little lost sometimes. Not that the ORM is hard to learn (it's awesome!), but the abstraction makes it harder for me to understand the code - my old tricks for debugging code no longer work. 

In my last job, I was working with Java Spring and JDBC query templates. We had dedicated files for housing our query templates, and it was very easy to find where queries were used. So whenever I needed to find where and when a specific thing would happen (or if it could happen) in the code, I would build a vague understanding of the data model and search for a query that selected/updated/deleted the relevant data. Now, since DB interactions are all abstracted away into the ORM, I feel a bit sluggish when it comes to debugging or figuring out our code. Currently, I try to search for what I need through model usages and whatnot, but that can be tough since the ORM is pretty flexible in how you use the model. For example, you could do 

    model.update(somekey: "someValue")

 OR 

    model.someKey = "someValue"

followed by a 

    model.save

, not to mention that the name 

    model

could vary, as you could be aliasing it in lots of different ways, especially if it's being accessed through *another* model as a nested attribute, or through a model association, or whatever else rails is capable of. Also, the key I'm looking for could be fairly ambiguous (i.e., the key "approved" could be used across multiple models). And of course, you can always update multiple attributes in any order, like 

    model.update(some_other_key: "blah", the_key_i_want: "blahhh")

, so this has eliminated a lot of go-to options for me. I always have to try to search multiple things and mentally union all of the results I've found. Because of all of this, I've been finding it hard to confidently (and quickly) track down *every* instance of the model being used in the way I'm looking for, which produced plenty of cortisol when the lead is AFK and someone posts in our team chat with a question or issue. I'm trying not to resort to regex every time I do a global search 😕  


I think I might be stuck in my old way of thinking, and I just need a way to break out of it... Any tricks or tips? What can I be doing differently?
## [6][puts_debuggerer (debugger-less debugging FTW)](https://www.reddit.com/r/rails/comments/f1tb3w/puts_debuggerer_debuggerless_debugging_ftw/)
- url: https://www.reddit.com/r/rails/comments/f1tb3w/puts_debuggerer_debuggerless_debugging_ftw/
---
[https://github.com/AndyObtiva/puts\_debuggerer](https://github.com/AndyObtiva/puts_debuggerer)
## [7][Rails 6, Webpacker and adding a simple JS library](https://www.reddit.com/r/rails/comments/f1zvs3/rails_6_webpacker_and_adding_a_simple_js_library/)
- url: https://www.reddit.com/r/rails/comments/f1zvs3/rails_6_webpacker_and_adding_a_simple_js_library/
---
Seems to be a lot of confusion in the Rails community nowadays about best Webpacker practices. Tons of questions but with very few answers. I hope this changes soon.

Anyway, would anybody here mind having a look?

[https://gist.github.com/ratahtatah/7c2e5a2ea8eb9a77b8f9eb4a259f8828](https://gist.github.com/ratahtatah/7c2e5a2ea8eb9a77b8f9eb4a259f8828)

Thanks!
## [8][Strategic - Painless Strategy Pattern in Ruby and Rails](https://www.reddit.com/r/rails/comments/f1t8sf/strategic_painless_strategy_pattern_in_ruby_and/)
- url: https://www.reddit.com/r/rails/comments/f1t8sf/strategic_painless_strategy_pattern_in_ruby_and/
---
[http://andymaleh.blogspot.com/2020/01/strategic-v080-painless-strategy.html?m=1](http://andymaleh.blogspot.com/2020/01/strategic-v080-painless-strategy.html?m=1)
## [9][React on Rails question for the community- Incorporating a Map API.](https://www.reddit.com/r/rails/comments/f1txc8/react_on_rails_question_for_the_community/)
- url: https://www.reddit.com/r/rails/comments/f1txc8/react_on_rails_question_for_the_community/
---
Hey I'm a junior dev working on a group project. We have been trying to use leaflets map API within our website and can't seem to figure it out how to get it working! If anyone has some good resources or advice for us it would be greatly appreciated.
## [10][Rails 6 multiple databases &amp; test parallelisation](https://www.reddit.com/r/rails/comments/f1qz8e/rails_6_multiple_databases_test_parallelisation/)
- url: https://www.reddit.com/r/rails/comments/f1qz8e/rails_6_multiple_databases_test_parallelisation/
---
Hi all!

Having some issues with implementing multiple databases, specifically a write and a replica. I've been following the [rails guide](https://guides.rubyonrails.org/active_record_multiple_databases.html) and am not setting up automatic connection switching because I believe that we still make some writes for GET connections and I'm wanting to go with the smallest possible changes to start with. I'm mainly interested in switching two heavier operations over to the read server to relieve pressure from the write DB and so am using manual connection switching e.g.

&amp;#x200B;

    ActiveRecord::Base.connected_to(role: :reading) do
        # larger read to the replica
    end

&amp;#x200B;

I'm running into two problems with my tests:

1. Database connections default to the writing role, but after the above manual connection switch, the database connection continues to use the replica which causes issues when I want to write after the block. I've been debugging this by checking the configuration through  
`ApplicationRecord.connection.instance_variable_get(:@config)`  
I was under the impression that the DB connection should only swap for the duration of the block?
2. Another issue I get is that the database used is different between the read and write roles. I can see this because the database suffix is different (e.g. `database-3` for the write vs `database` for the read).

Unless I'm missing some article on the rails site, I can't see much info for multiple databases and testing so hoping for some tips.

&amp;#x200B;

Thanks!
## [11][Good resources for learning about ports and proxies with Rails?](https://www.reddit.com/r/rails/comments/f1tbde/good_resources_for_learning_about_ports_and/)
- url: https://www.reddit.com/r/rails/comments/f1tbde/good_resources_for_learning_about_ports_and/
---
Do you know any good resources for learning about ports and proxies with regards to Ruby on Rails? Thanks
