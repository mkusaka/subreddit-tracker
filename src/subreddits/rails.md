# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/jsthk8/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/jsthk8/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Should models use direct associations or through?](https://www.reddit.com/r/rails/comments/juhtou/should_models_use_direct_associations_or_through/)
- url: https://www.reddit.com/r/rails/comments/juhtou/should_models_use_direct_associations_or_through/
---
I am wondering when it comes to scaling, is using through associations a better way to model a database? I am assuming the fewer columns, the better?

&amp;#x200B;

Assuming we have the following models

&amp;#x200B;

**Order**

has\_many line\_items

**Product**

*#product on the website*

has\_many line\_items, optional: true

belongs\_to vendor

**LineItem**

*#this is the products in the cart*

belongs\_to product

**Vendor**

*#vendor of the product*

has\_many products

has\_many line\_items through products

&amp;#x200B;

When an order comes in, the order will have many line items.  Is there any reason to have a vendor\_id association in the LineItem model? 

Or is \`through:\` the best case scenario?

I assume through is the best way to go, especially for database bloat when scaling on unneeded space being ocupied but I am wondering is there is any benefit to not using a through association
## [3][Best chatbot gems/libraries to use for rails app?](https://www.reddit.com/r/rails/comments/jucos4/best_chatbot_gemslibraries_to_use_for_rails_app/)
- url: https://www.reddit.com/r/rails/comments/jucos4/best_chatbot_gemslibraries_to_use_for_rails_app/
---
Just need a basic text chatbot to save information to a model &amp; wondering what you guys would recommend for gems/tutorials to do this with a rails app.
## [4][a bootstrap theme implementation with rails !!!](https://www.reddit.com/r/rails/comments/ju9s0n/a_bootstrap_theme_implementation_with_rails/)
- url: https://www.reddit.com/r/rails/comments/ju9s0n/a_bootstrap_theme_implementation_with_rails/
---
Hello Folks, spent 3 days trying to implement an html template to new project with rails the problem is  the theme based on Webpack and laravel-mix which compiles Sass, ES6 JavaScript, handles production builds, watchers, multiple CSS themes and more. This entire workflow is contained into an installable package named theme-mix. any idea what that means please help ?

wtf how i can use laravel-mix with rails webpacker ? any idea i would appreciate that .

&amp;#x200B;

https://preview.redd.it/5azi6qegx9z51.png?width=2056&amp;format=png&amp;auto=webp&amp;s=4ee0251ef67326c1c4617b9246716ad3d4121446

https://preview.redd.it/n586mregx9z51.png?width=2560&amp;format=png&amp;auto=webp&amp;s=30d1783293e80e5e20fb6cd26bb8d260f660407d

https://preview.redd.it/w1ngsf4xx9z51.png?width=2560&amp;format=png&amp;auto=webp&amp;s=e3e2417e7a21226c59a7385e7e4d5683e0c3dca3

https://preview.redd.it/lawkcregx9z51.png?width=2338&amp;format=png&amp;auto=webp&amp;s=112fbf66ca9ab1a5f629da46272c9278438fb75a

https://preview.redd.it/wb52dregx9z51.png?width=2416&amp;format=png&amp;auto=webp&amp;s=65e9debd41049ca5d06949c8c93ed193fffeb9e5

https://preview.redd.it/z034m8uwx9z51.png?width=2518&amp;format=png&amp;auto=webp&amp;s=f08e76ccdb00d3011fdb117e85c499c0b9f67b56
## [5][How do databases work with RoR?](https://www.reddit.com/r/rails/comments/jtwjx3/how_do_databases_work_with_ror/)
- url: https://www.reddit.com/r/rails/comments/jtwjx3/how_do_databases_work_with_ror/
---
I'm learning Rails and while I feel comfortable with a lot of it, this is the main area I'm struggling with. My past experience is with Flask/Express, where the database is more "separated" from the app.

With these Rails tutorials, they all have you sort of creating the db through rails-related commands like `rake db:migrate`. It's so opaque-- I don't really understand when/how the database is being initialized and configured. I mean, the prod database has to be set up with some sort of provider, doesn't it? This is one of those cases of "I don't know what I don't know" so I can't even talk about it intelligently.

Like if I use AWS db hosting, running rake commands on the web server isn't really a thing, right? I just don't know what the production workflow is for setting up a rails app with a SQL database.
## [6][Has anyone used BlitzJS ?](https://www.reddit.com/r/rails/comments/ju0gcx/has_anyone_used_blitzjs/)
- url: https://www.reddit.com/r/rails/comments/ju0gcx/has_anyone_used_blitzjs/
---
On their home page, it is clearly stated that the framework is inspired by Rails. Given the nice adoption curve, has anyone with a RoR background used it recently, and could give a fair opinion ? I find it weird that no Rails-like framework appeared in the JS world. SailsJS didn't get much traction.
## [7][I have a problem with a table that exists not loading?](https://www.reddit.com/r/rails/comments/jtva67/i_have_a_problem_with_a_table_that_exists_not/)
- url: https://www.reddit.com/r/rails/comments/jtva67/i_have_a_problem_with_a_table_that_exists_not/
---
When I go into the rails console 

I use to be able to use the command Business.all now when I type in Business it prints back like it's a string. How can I fix this problem?

It is causing me not to be able to run my code: Business.find\_by(user\_id: user\_id)

I believe I messed something up by not realizing.
## [8][6 Things to Do When Inheriting Legacy Rails Apps](https://www.reddit.com/r/rails/comments/jtgjla/6_things_to_do_when_inheriting_legacy_rails_apps/)
- url: https://www.reddit.com/r/rails/comments/jtgjla/6_things_to_do_when_inheriting_legacy_rails_apps/
---
One of our engineers wrote a guide to help people get off on the right foot when inheriting a Rails app and [I wanted to share it here.](https://nextlinklabs.com/insights/six-tips-for-inheriting-a-legacy-rails-app) Let me know what you think.
## [9][What’s the best way to integrate WordPress into a web application?](https://www.reddit.com/r/rails/comments/jtqntp/whats_the_best_way_to_integrate_wordpress_into_a/)
- url: https://www.reddit.com/r/rails/comments/jtqntp/whats_the_best_way_to_integrate_wordpress_into_a/
---
Hello all!

I just wanted to know what the best way to integrate WordPress into a web application.

I’ve been looking into a couple ways...

1. Using WordPress and pointing the server to a URL like blog.example.com
2. Using WordPress as a headless CMS
3. Using a completely different CMS like ButterCMS

I don’t know what would be the best option, just wondering if anyone else had experience with WordPress and how they integrated WordPress into their own web application.

Thanks!
## [10][How to access multiple images attached to a single active record instance?](https://www.reddit.com/r/rails/comments/jto91u/how_to_access_multiple_images_attached_to_a/)
- url: https://www.reddit.com/r/rails/comments/jto91u/how_to_access_multiple_images_attached_to_a/
---
Hi Folks, 

I've set up a model where I have a has_many_attached :images, but was wondering how to access each image that is attached. Do you need to iterate through or is there a way to pull them directly from each given record?
## [11][limit rails/ruby memory usage](https://www.reddit.com/r/rails/comments/jtjz9g/limit_railsruby_memory_usage/)
- url: https://www.reddit.com/r/rails/comments/jtjz9g/limit_railsruby_memory_usage/
---
Hello All - 

I'm running a rails app in AWS ECS (docker). I initially started with containers that were allocated 4gb of memory. This mostly worked, but memory usage seemed to grow over time, and if it approached or hit 100% memory usage, AWS ECS kills my container. Obviously not good.

I doubled my container memory limit to 8gb and I'm still seeing the same behavior. 

I understand that it's desirable to use "most" of the memory, but I'd like to limit to (for example) 90% usage. Is it possible to set a hard limit on ruby/rails memory usage (similar to how you can with java)? I did a bunch of googling and saw a bunch of GC setting, but it didn't see anything that looked like a setting for upper limit of memory allocation.

Thanks in advance!
