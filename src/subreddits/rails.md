# rails
## [1][Has anyone ever had extremely high response times on the ActionCable endpoint?](https://www.reddit.com/r/rails/comments/g3hq8q/has_anyone_ever_had_extremely_high_response_times/)
- url: https://www.reddit.com/r/rails/comments/g3hq8q/has_anyone_ever_had_extremely_high_response_times/
---
On localhost, inspector's network tab either shows these /cable requests with *absurd* times:

https://preview.redd.it/q1vjzezt8it41.png?width=2704&amp;format=png&amp;auto=webp&amp;s=08bd5f10b2737ead8ef6cb0af3c9ada16b3d3f69

...or it gives me this nanoscale domino stack of pending requests:

https://preview.redd.it/skp07rw39it41.png?width=2718&amp;format=png&amp;auto=webp&amp;s=4c001441853c3c680129f8f1ffeca02148f57ab8

I have a similar problem with Heroku. I'll be the only user on the site, I'll be doing *nothing*, and the logentries addon will bombard my email with consecutive High Response Time Alerts:

    25 Dec 2019 14:13:31.061299 &lt;158&gt;1 2019-12-25T22:13:28.421009+00:00 heroku router - - at=info method=GET path="/cable" host=www.######.com request_id=#### fwd="####" dyno=web.1 connect=1ms service=531114ms status=101 bytes=174 protocol=https High Response Time

I'm using Ruby version 2.5.0, Rails version 5.2.0, and Puma version: 4.3.1. I've tried both Iodine and AnyCable, and neither fixes this problem. Leading me to believe that the problem is elsewhere. But my ActionCable setup is very standard - an almost exact copy+paste of the setups suggested in docs and tutorials everywhere.

I've found a very small handful of people experiencing a similar issue, but none of them have had any luck. So I'm trying here. Have you experienced this, and did you fix it? If so - how? If not - what could possibly be causing this? Thank you so much!
## [2][Webpack on aws Cloud9](https://www.reddit.com/r/rails/comments/g3fw6w/webpack_on_aws_cloud9/)
- url: https://www.reddit.com/r/rails/comments/g3fw6w/webpack_on_aws_cloud9/
---
I thought i installed webpack properly in my app but im missing folders and files in my file tree.

I am missing the webpack folder that is suppose to be under the config folder.

Its suppose to look like this I believe:

&amp;#x200B;

[This is not what my file tree looks like](https://preview.redd.it/hlydr97emht41.png?width=193&amp;format=png&amp;auto=webp&amp;s=a9c242ea6756c29f142c1a91ca6cacf77317b25d)

I'm using aws cloud9 environment.

I followed the instructions here:  [https://docs.aws.amazon.com/sdk-for-javascript/v2/developer-guide/webpack.html](https://docs.aws.amazon.com/sdk-for-javascript/v2/developer-guide/webpack.html) 

How do I get webpack to integrate/work properly with aws cloud9?

Ive installed nod.js and nvm.

Thanks
## [3][Questions about scaling workers with customers](https://www.reddit.com/r/rails/comments/g3dquq/questions_about_scaling_workers_with_customers/)
- url: https://www.reddit.com/r/rails/comments/g3dquq/questions_about_scaling_workers_with_customers/
---
Hi guys,

Need some advice about scaling the use of workers with my rails app. Right now, for each customer that kicks off a task, it creates about 20 workers for that one particular task, and each of those workers can take approximately 10-15 minutes to run. However, I feel like this may become a problem as the list of customers grow and may not be the best practice for growth.

I have two queues at the moment as well -- developer, and production -- and the purpose is just so that I can manage sidekiq workers in both production and developer independently without any confusion.

Is this something that should be OK as it grows to, say, 500 workers running simultaneously? In my credentials file I do have the pool set to 200 for now, but wasn't sure how much higher I could raise it without having issues unless it just really depends on system resources.

I should also make it known that I am not a professional by any means and have just used rails in the last few years to get by.

Any help or advice would be greatly appreciated
## [4][How to setup PostgreSQL database with models and controllers](https://www.reddit.com/r/rails/comments/g3gybz/how_to_setup_postgresql_database_with_models_and/)
- url: https://www.reddit.com/r/rails/comments/g3gybz/how_to_setup_postgresql_database_with_models_and/
---
I am rather new to Rails but I have gone through the tutorial on their website. I want to make something similar to PCPartPicker. I have gone ahead an planned out my structure for the tables in the database. They go something like this:

    build
      id title
      1  My dream desktop
    
    build_component
      build_id component_id count
      1        1            1   (1 ryzen9 cpu)
      1        2            2   (2 hyper-x 8G memory sticks)
      
    
    component
      id   type_id  name           price   (all components have name and a price)
       1   1            Ryzen 9 3900X  400
       2   2            Hyper-X DDR4   89
       
    component_attribute
      component_id attribute_id value
      1            2            3900
      1            4            AM4
      2            1            8000
      2            2            3000
           
    
    component_type 
      id      name
       1      CPU
       2      Memory Stick
       3      Motherboard
       4      Storage
    
    type_attribute (not strictly necessary, but would make creating forms and list easier)
      id_type  id_attribute
      1            2
      1            4  (cpu has speed, socket type)
      2            1
      2            2
      2            5 (memory has speed, size and memory_type)
       
    attribute
      id      name
      1       size_mbytes
      2       speed_mhz
      4       cpu_socket
      5       memory_type

So each of those are an individual table. Now for the question. In the example provide by the website. One would create a model for each table and then that model was assigned or worked with a controller. In this case, if I want to have something similar to the articles example where my parts are just listed out, would I make multiple models and all assign them to the same controller? 

My second question is: Would I create these tables in PostgreSQL and then just use Rails to read and write to them or would I create just an empty database and use Rails to add all the tables?
## [5][Eager Loads + Calculations (.pluck, .sum, etc) gotcha that we just ran into.](https://www.reddit.com/r/rails/comments/g338xm/eager_loads_calculations_pluck_sum_etc_gotcha/)
- url: https://www.reddit.com/r/rails/comments/g338xm/eager_loads_calculations_pluck_sum_etc_gotcha/
---
If you eager-load a has\_many associations your .sum or .calculate or .pluck will return duplicate values.

So:

`@products.sum(:price) != @products.includes(:reviews).sum(:price)`

Even this returns a wrong value:

`@products.includes(:reviews).pluck(:price).sum # this will include duplicate price values`
## [6][Need some help with getting started](https://www.reddit.com/r/rails/comments/g2y98l/need_some_help_with_getting_started/)
- url: https://www.reddit.com/r/rails/comments/g2y98l/need_some_help_with_getting_started/
---
Hey! I'm interested in getting into Ruby on Rails, and I'm wondering whether any of you might be able to point me in the direction of some good resources.

I have active subscriptions to [Lynda.com](https://lynda.com/) and [GoRails.com](https://gorails.com/), and I have the GitHub Student Developer Pack. My end-goal is to build a basic social-networking site for my school, not to become a paid web developer!

I have loads of experience in Python, HTML, JS and CSS, and I launched myself into a Basics of Ruby course on Lynda, so I have enough experience there too.

I was watching this free Udemy course, which looked perfect; [https://www.udemy.com/course/8-beautiful-ruby-on-rails-apps-in-30-days/learn/lecture/4336792?start=240](https://www.udemy.com/course/8-beautiful-ruby-on-rails-apps-in-30-days/learn/lecture/4336792?start=240), but in the Announcements section it was apparently severely outdated. Does anyone know of something similar? I honestly prefer video content to reading (with the exception of books).

And I'd prefer to not spend heaps. I've looked at The Odin Project and the Essential RoR Training courses on Lynda but the RoR course seems far too theoretical. I want to get creating ASAP.

Many thanks!
## [7][You're not stubbing, stupid!](https://www.reddit.com/r/rails/comments/g2zi9v/youre_not_stubbing_stupid/)
- url: https://www.reddit.com/r/rails/comments/g2zi9v/youre_not_stubbing_stupid/
---
Hi, I've been struggling with writing a passing cucumber that deletes a resource and shows a confirmation dialog to the user. The controller sends a request to an external API, so I stubbed this request. I finally found out what was going wrong and wrote this blog post about it. Hope it'll help some (some day).

[https://www.theguild.nl/youre-not-stubbing-stupid/](https://www.theguild.nl/youre-not-stubbing-stupid/)
## [8][Why is it not saving nested attributes?](https://www.reddit.com/r/rails/comments/g2r7z6/why_is_it_not_saving_nested_attributes/)
- url: https://www.reddit.com/r/rails/comments/g2r7z6/why_is_it_not_saving_nested_attributes/
---
I have created a model called lead and it `has_many text_field` and `text_field has_many text_value` , I have been testing it in the rails console with, `text.update(text_value: [value])`  and it saves!

but within the views it does not, please help.

[The is the new.html.erb \(Highlighted text is my nested attribute text\_value\)](https://preview.redd.it/g26ixzm9s9t41.png?width=1280&amp;format=png&amp;auto=webp&amp;s=ca6d6aec6a4ba94e82afa5507d6d08299a135010)

&amp;#x200B;

[lead controller \(highlighted with test\_value nested within text\_field\)](https://preview.redd.it/22zlxyxns9t41.png?width=1280&amp;format=png&amp;auto=webp&amp;s=e7013fc36e26e34ee34f3ba696a37f0d4fdac42a)

&amp;#x200B;

[text\_field model](https://preview.redd.it/xj6slas3t9t41.png?width=1280&amp;format=png&amp;auto=webp&amp;s=af00c1b59eb256d4d50dbe831204e2e3697d3f9e)

&amp;#x200B;

[text\_value model](https://preview.redd.it/7tznnyw6t9t41.png?width=1280&amp;format=png&amp;auto=webp&amp;s=aa2b886bbbc3f3da9aa702e4e7a81d721a6ed59d)

&amp;#x200B;

[lead model](https://preview.redd.it/o4xykotbt9t41.png?width=1280&amp;format=png&amp;auto=webp&amp;s=cdd985669de13803e45edff0488a31ba0b5bcf1e)

&amp;#x200B;

[schema.rb \(highlighted red is text\_value and text\_field\)](https://preview.redd.it/cznrrv0kt9t41.png?width=1280&amp;format=png&amp;auto=webp&amp;s=f4f01465109a77d3b4674f31ee6d5b7c37a7676a)
## [9][The Ruby Blend: ViewComponent at GitHub with Joel Hawksley](https://www.reddit.com/r/rails/comments/g2gy1t/the_ruby_blend_viewcomponent_at_github_with_joel/)
- url: https://www.reddit.com/r/rails/comments/g2gy1t/the_ruby_blend_viewcomponent_at_github_with_joel/
---
Joel Hawksley is the guest on today’s episode. He is a Software Engineer at GitHub and works on the design system team, leading development on ViewComponent, formerly known as ActionView::Component. Joel is a very encouraging, uplifting, and inviting guest who tells it like it is. Andrew even throws in a little “homework” for those of you listening! If you want to learn more about how GitHub is leveraging ViewComponent, [give this episode a listen!](https://fireside.fm/s/ouBAUjGy+8JZwBXVx)
## [10][Rails 6 + docker-compose : simplest possible Hello World project™](https://www.reddit.com/r/rails/comments/g2c48p/rails_6_dockercompose_simplest_possible_hello/)
- url: https://www.reddit.com/r/rails/comments/g2c48p/rails_6_dockercompose_simplest_possible_hello/
---
Hello mates, I spend some time to find the simplest possible configuration to start a Rails project using Docker.

[https://github.com/bdavidxyz/simplest-rails-docker](https://github.com/bdavidxyz/simplest-rails-docker)

I picked some ideas from other projects from GitHub, trying to remove as much files and as much lines per file as possible.

With Ruby-on-Rails, you can't start a project without a properly initialized, configured, and started database. You also can't start the local server without a webpack dev server running (from Rails 6 I guess). So that's make 3 services running at least before even to print a "hello world" page.

This kind of project is extremely useful to test various tools and versions without polluting your local machine. I find it also more professional to use Docker instead of brittle, quickly-too-old documentation. So it could be a starting point for a more serious configuration.

What do you think ? What could be simplified ? Do you have more simple examples ? How do you test tools in isolation ? Thanks for your returns, and stay safe :)
