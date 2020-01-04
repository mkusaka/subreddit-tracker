# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/e5z9l0/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/e5z9l0/personal_projects_show_off_your_own_project_andor/
---
In this thread you can showcase your personal pet project to other redditors.

Need help with a specific problem or just wanna have some extra eyeballs on your code? Ask away!

A suggested format to get you started:

1. **Name of your project**
2. **A short description**
3. **Application stack**
4. **Link to Live app**
5. **Link to GitHub**
6. **You experience level**
7. **Other information or areas that you would like advice on**

 

^(Many thanks to Kritnc for getting the ball rolling.)
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/eiidds/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/eiidds/personal_projects_show_off_your_own_project_andor/
---
In this thread you can showcase your personal pet project to other redditors.

Need help with a specific problem or just wanna have some extra eyeballs on your code? Ask away!

A suggested format to get you started:

1. **Name of your project**
2. **A short description**
3. **Application stack**
4. **Link to Live app**
5. **Link to GitHub**
6. **You experience level**
7. **Other information or areas that you would like advice on**

 

^(Many thanks to Kritnc for getting the ball rolling.)
## [3][Migrate Your Rails App from Heroku to AWS Lambda](https://www.reddit.com/r/rails/comments/ejpqku/migrate_your_rails_app_from_heroku_to_aws_lambda/)
- url: https://www.reddit.com/r/rails/comments/ejpqku/migrate_your_rails_app_from_heroku_to_aws_lambda/
---
**Using The New ActiveRecord Aurora Serverless Adapter.** Are you someone with a pet Rails project running on a Free, Hobby, or Professional Heroku plan? Perhaps your company or freelance gig has a valuable, but infrequently used, Rails application? Such applications make great candidates for both AWS Lambda &amp; Aurora Serverless using our [Lamby](https://lamby.custominktech.com) gem.

Hoping other Rails developers find [this guide](https://technology.customink.com/blog/2020/01/03/migrate-your-rails-app-from-heroku-to-aws-lambda/) useful as a way to learn AWS Lambda and other tools within the AWS ecosystem.
## [4][What is the best way? (Rails app, smartTVs and views)](https://www.reddit.com/r/rails/comments/ejip1v/what_is_the_best_way_rails_app_smarttvs_and_views/)
- url: https://www.reddit.com/r/rails/comments/ejip1v/what_is_the_best_way_rails_app_smarttvs_and_views/
---
 I  am working in my project about smart TVs that will be placed in some  points of my city. Those points display content uploaded by my customers  to my app, which is made in Ruby on rails.

The appoach I have used for the TVs, is a raspberry PI which load a  web view full screen on start. Problem is... if sometimes does not load  correctly, then the only option is a manual restart... that is not a  choice.

I have been considering to do a little client which just ask for the  JSONs and make like a little game UI with some visual libraries (PyGame,  Ruby2d..etc). The advantages would be full control and the posibility to  implement a so-called 'manual caché' that will help me to spend less  money on AWS.

The question is: How can I do that as easy as possible? Any suggestion? Any bad or good aspect I have not considered?

PD: I have tech background but not consider myself really a professional, but an ethernal hack-ish amateur.

Thanks in advance :)
## [5][Concurrent Ruby vs Parallel? Has anyone tried both?](https://www.reddit.com/r/rails/comments/ejlt9z/concurrent_ruby_vs_parallel_has_anyone_tried_both/)
- url: https://www.reddit.com/r/rails/comments/ejlt9z/concurrent_ruby_vs_parallel_has_anyone_tried_both/
---
In our codebase, we're often fetching data sequentially. I come from a JS/Hack background, so I'm used to running things in parallel using async/await. 

I found Parallel ([https://github.com/grosser/parallel](https://github.com/grosser/parallel)) and Concurrent Ruby ([https://github.com/ruby-concurrency/concurrent-ruby](https://github.com/ruby-concurrency/concurrent-ruby)). They can both do what I want, but Parallel is simpler.

I'm not experienced with Ruby, so I can't evaluate these packages. Has anyone tried them? What are some pros and cons of each?
## [6][How to organize project for web app that has different "portals"/main pages for different products?](https://www.reddit.com/r/rails/comments/ejlrh6/how_to_organize_project_for_web_app_that_has/)
- url: https://www.reddit.com/r/rails/comments/ejlrh6/how_to_organize_project_for_web_app_that_has/
---
I've already built a web app for a company that sells product A. Now, I've been hired again to expand on the same web app for product B. I envision each product to have its own main or home pages. The only tables I imagine being shared between both products is the User table. There are also some more general-purpose files and functions I imagine will be used by both "sub-apps".

I'm a little lost into how to organize this project. I imagine my URL will look like "[websitename.com/productA](https://websitename.com/productA)/" and "[websitename.com/productB/](https://websitename.com/productB/)". My initial instinct to use "namespaces" which I unfortunately know nothing about and haven't used yet. I was wondering if any of you have done this before and could point me to some resources on how to go about making this decision.

Edit: For example I'm not sure whether to have my apps folder split into appA and appB, or whether I should split in specific folders like models, eg: app/models/productA and app/models/productB
## [7][How to upload Live Video in Rails?](https://www.reddit.com/r/rails/comments/ejhr30/how_to_upload_live_video_in_rails/)
- url: https://www.reddit.com/r/rails/comments/ejhr30/how_to_upload_live_video_in_rails/
---
I want to implement Live Streaming in Rails. Users can record video on their device and send it to other clients while recording. I've googled and found out about WebRTC which uses P2P protocol to achieve this. However I also want to save the video on the server using ActiveStorage. For that I want the recorder client to send a stream of the video to Rails back end as well. First question is how can I send a stream of video to back end while the client is recording it. Now I know how to store files in ActiveStorage if I have the entire data. So my second question is how to achieve storing a stream of data in ActiveStorage?
## [8][How to seed a Postgres Production Database on Heroku?](https://www.reddit.com/r/rails/comments/ejljxj/how_to_seed_a_postgres_production_database_on/)
- url: https://www.reddit.com/r/rails/comments/ejljxj/how_to_seed_a_postgres_production_database_on/
---
I've deployed a rails app (rails v 5.2 and ruby 2.6) to heroku and have no problem migrating my database with commands like `heroku run rake db:migrate` but haven't found an efficient mechanism for seeding that database. Running `heroku run rake db:seed` doesn't seem to yield any results and all posts on stack overflow seem outdated. Is it possible to seed a production pg database on heroku?
## [9][Any tools and tutorials on rails and activitypub?](https://www.reddit.com/r/rails/comments/ejfj4p/any_tools_and_tutorials_on_rails_and_activitypub/)
- url: https://www.reddit.com/r/rails/comments/ejfj4p/any_tools_and_tutorials_on_rails_and_activitypub/
---
I recently discovered [pump.io](https://pump.io), activitypub and the concept of "federated" networks. One of my friends invited me to the network mastodon (in case you don't know, there's also [ruby.social](https://ruby.social) instance for people like us, ruby lovers!) and I made an account. I realized mastodon itself is written in ruby. I'm curious, is there anyway to make our rails apps federated? I did a bit of search but I found some of those projects are discontinued. And if yes, is there any tutorials? 

Thanks a lot.
## [10][Schedule tasks at Heroku?](https://www.reddit.com/r/rails/comments/ejeg6t/schedule_tasks_at_heroku/)
- url: https://www.reddit.com/r/rails/comments/ejeg6t/schedule_tasks_at_heroku/
---
I'm using an [addon](https://devcenter.heroku.com/articles/scheduler) which lets you run a rake or rails command every 10 minutes, but I need to run it more often than that, like every 2 or 3 minutes.

How can I do this?
## [11][Add to Cart functionality with multiple models involved](https://www.reddit.com/r/rails/comments/ejl9xn/add_to_cart_functionality_with_multiple_models/)
- url: https://www.reddit.com/r/rails/comments/ejl9xn/add_to_cart_functionality_with_multiple_models/
---
I'm creating a functionality where users can add a product to their cart. However, this product can have 0 or many 'add ons' that the user selects before clicking 'add to cart'. The Model is user &gt; cart &gt; cart items &gt; product\_type\_x &gt; product\_type\_x\_add\_ons. I've done it like that, because later on there will be other product types which have different attributes and no add ons (more traditional products). In other words, if the user selects a product\_x and then 2 add ons, and clicks add to cart, I would like to first: check if the user has a cart, if not, create it. Add a record to cart\_items, add a record to cart\_product\_type\_x, add two records to cart\_product\_type\_x\_add\_ons. All of the cart\_ product tables link back to the product and add on tables. So I was wondering what is a good way to implement this, given people recommend putting logic out of controller. Validations will need to be performed of course to ensure users are not adding add ons that do not apply to that product ( through a direct request or view source 'hacking' of ids). So all saves must fail if one validation/save to db fails.
## [12][Speed up your RoR test suite](https://www.reddit.com/r/rails/comments/ejbnrg/speed_up_your_ror_test_suite/)
- url: https://www.reddit.com/r/rails/comments/ejbnrg/speed_up_your_ror_test_suite/
---
https://medium.com/better-programming/cut-your-rspec-minitest-runtime-with-testprof-d19e55783050
I published a piece on Better Programming today. It's about making your RoR tests run much much faster.

Would love to hear your thoughts
