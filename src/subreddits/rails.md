# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/jsthk8/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/jsthk8/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][How do databases work with RoR?](https://www.reddit.com/r/rails/comments/jtwjx3/how_do_databases_work_with_ror/)
- url: https://www.reddit.com/r/rails/comments/jtwjx3/how_do_databases_work_with_ror/
---
I'm learning Rails and while I feel comfortable with a lot of it, this is the main area I'm struggling with. My past experience is with Flask/Express, where the database is more "separated" from the app.

With these Rails tutorials, they all have you sort of creating the db through rails-related commands like `rake db:migrate`. It's so opaque-- I don't really understand when/how the database is being initialized and configured. I mean, the prod database has to be set up with some sort of provider, doesn't it? This is one of those cases of "I don't know what I don't know" so I can't even talk about it intelligently.

Like if I use AWS db hosting, running rake commands on the web server isn't really a thing, right? I just don't know what the production workflow is for setting up a rails app with a SQL database.
## [3][Has anyone used BlitzJS ?](https://www.reddit.com/r/rails/comments/ju0gcx/has_anyone_used_blitzjs/)
- url: https://www.reddit.com/r/rails/comments/ju0gcx/has_anyone_used_blitzjs/
---
On their home page, it is clearly stated that the framework is inspired by Rails. Given the nice adoption curve, has anyone with a RoR background used it recently, and could give a fair opinion ? I find it weird that no Rails-like framework appeared in the JS world. SailsJS didn't get much traction.
## [4][I have a problem with a table that exists not loading?](https://www.reddit.com/r/rails/comments/jtva67/i_have_a_problem_with_a_table_that_exists_not/)
- url: https://www.reddit.com/r/rails/comments/jtva67/i_have_a_problem_with_a_table_that_exists_not/
---
When I go into the rails console 

I use to be able to use the command Business.all now when I type in Business it prints back like it's a string. How can I fix this problem?

It is causing me not to be able to run my code: Business.find\_by(user\_id: user\_id)

I believe I messed something up by not realizing.
## [5][6 Things to Do When Inheriting Legacy Rails Apps](https://www.reddit.com/r/rails/comments/jtgjla/6_things_to_do_when_inheriting_legacy_rails_apps/)
- url: https://www.reddit.com/r/rails/comments/jtgjla/6_things_to_do_when_inheriting_legacy_rails_apps/
---
One of our engineers wrote a guide to help people get off on the right foot when inheriting a Rails app and [I wanted to share it here.](https://nextlinklabs.com/insights/six-tips-for-inheriting-a-legacy-rails-app) Let me know what you think.
## [6][What’s the best way to integrate WordPress into a web application?](https://www.reddit.com/r/rails/comments/jtqntp/whats_the_best_way_to_integrate_wordpress_into_a/)
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
## [7][How to access multiple images attached to a single active record instance?](https://www.reddit.com/r/rails/comments/jto91u/how_to_access_multiple_images_attached_to_a/)
- url: https://www.reddit.com/r/rails/comments/jto91u/how_to_access_multiple_images_attached_to_a/
---
Hi Folks, 

I've set up a model where I have a has_many_attached :images, but was wondering how to access each image that is attached. Do you need to iterate through or is there a way to pull them directly from each given record?
## [8][limit rails/ruby memory usage](https://www.reddit.com/r/rails/comments/jtjz9g/limit_railsruby_memory_usage/)
- url: https://www.reddit.com/r/rails/comments/jtjz9g/limit_railsruby_memory_usage/
---
Hello All - 

I'm running a rails app in AWS ECS (docker). I initially started with containers that were allocated 4gb of memory. This mostly worked, but memory usage seemed to grow over time, and if it approached or hit 100% memory usage, AWS ECS kills my container. Obviously not good.

I doubled my container memory limit to 8gb and I'm still seeing the same behavior. 

I understand that it's desirable to use "most" of the memory, but I'd like to limit to (for example) 90% usage. Is it possible to set a hard limit on ruby/rails memory usage (similar to how you can with java)? I did a bunch of googling and saw a bunch of GC setting, but it didn't see anything that looked like a setting for upper limit of memory allocation.

Thanks in advance!
## [9][Does this spec pattern make sense? Having an issue of undefined variables.](https://www.reddit.com/r/rails/comments/jtjpg1/does_this_spec_pattern_make_sense_having_an_issue/)
- url: https://www.reddit.com/r/rails/comments/jtjpg1/does_this_spec_pattern_make_sense_having_an_issue/
---
I have a spec pattern like this:


    describe ...Controller do

      before :each do

        ...

      end

      describe '...' do

        before :all do

          // define variable here for it blocks below

        end

        it '...' do

          // variable defined in before :all not accessible

        end

      end

    end


Initially I was using context but I was trying to use that nested/internal before to setup some stuff in the inner describe.
## [10][Is it a code smell to use redirects in order to forward results from one controller to another?](https://www.reddit.com/r/rails/comments/jtms5n/is_it_a_code_smell_to_use_redirects_in_order_to/)
- url: https://www.reddit.com/r/rails/comments/jtms5n/is_it_a_code_smell_to_use_redirects_in_order_to/
---
There is an existing route/controller workflow that needs to be used as part of another route/controller. There a few problems eg. lack of controller context and in this specific case regarding timers.

But while it seems possible to make the response of a controller get redirected to another controller... that seems like a bad approach.

Are there any other obvious methods... other than doing two requests from the client side?(to route 1/route 2)
## [11][Company Workflows](https://www.reddit.com/r/rails/comments/jsy5xv/company_workflows/)
- url: https://www.reddit.com/r/rails/comments/jsy5xv/company_workflows/
---
I have a few questions regarding workflow and i'd love to hear some feedback from you guys about it and to compare it to my current workflow.

I'm mainly curious as to what your workflow is like with other team members, mainly designers.

I understand every project is different and dependent on the client, the workflow would be different.

I'm currently in a job where it's very 'get it done and get it out the door as quick as possible', this tends to leave a lot of issues when going live for example, the code breaking because it thinks a field is there and when it does html\_safe onto it. When this happens, I just tend to fix the error and get it deployed ASAP which works.

When working with a design, I tend to get everything done, have full control of all the margins, padding, font-sizes etc so I can change it but it's always just the little things that I don't see as i've been looking at it for too long and begins to just looks normal to me (I assume most people have this issue) which is fine and simple to change.

Some of these questions might not apply to freelancers as I imagine the client already has a design and they just want some product back ASAP (or within a time frame)

1) Before the developer get's the end design should everything should be tied up, all content in and approved by the client before the developer starts work on it? 

2) Do you communicate with the designer throughout and ask them to give feedback on something you've done and make changes if needed?

3) Most projects aren't smooth sailing as you'd expect due to client changes, design changes or config changes, how do you get around these? I know the codebase and what I should change and where it is. 

4) How do you tackle feedback? For me, i'm really bad at taking criticism but that's just due to the person I am and I hate receiving negative feedback as I just want to make people happy with what I do.

I'm really just looking at ways I could potentially make my workflow a little better as i'm seeing myself doing something and then doing it again and again and again and I feel like this could be easily avoided if it all was done and correct in the first place?
