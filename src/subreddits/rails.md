# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/ep2dw9/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/ep2dw9/personal_projects_show_off_your_own_project_andor/
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
## [2][Speed of development on Rails](https://www.reddit.com/r/rails/comments/eul9bm/speed_of_development_on_rails/)
- url: https://www.reddit.com/r/rails/comments/eul9bm/speed_of_development_on_rails/
---
Hi, guys, I'm Alex and I work as a Product designer right now, but always wanted to be able to build products/side-projects by myself. So I made some research on tech and found that the **RoR** is the fastest framework in terms of development speed (fix me if I'm wrong).

But for me, it's quite hard to measure the development speed, so that's why I'm here.

I want to ask you, guys, how long it will take you to build a [Dribble](https://dribbble.com/) (I read that it was built using RoR) clone with basic features like:  
\- Account system (Where user can just make basic changes in acc info, e.x.change password, email, name);  
\- Image uploading;  
\- Commenting;  
\- Likes &amp; post views;  
\- Filtering by categories;

Thanks in advance to all RoR community.
## [3][The common way of documenting ruby and typescript code.](https://www.reddit.com/r/rails/comments/eulopd/the_common_way_of_documenting_ruby_and_typescript/)
- url: https://www.reddit.com/r/rails/comments/eulopd/the_common_way_of_documenting_ruby_and_typescript/
---
Hello.   
I have quite a big application which is using rails for backend and angular for the frontend. I want to write in-code documentation for it (and latter generate the documentation with some tool). The problem is that I don't want to have separate and different ways of doing it, instead, I would like to have one tool for both. So basically:  
1. Writing in-code documentation for ruby parts  
2. Writing in-code documentation for angular parts   
3. Generate the documentation with the code snippets and put in online  
4. To have the ability to see the documentation in preview in my IDE.   


Once again the question:

Is there are a tool that will be suitable for generating the documentation for ruby and angular parts altogether?
## [4][Help needed: what's a fast, preferably non-video way to brush up on recent, corporate-level, advanced rails knowledge and techniques?](https://www.reddit.com/r/rails/comments/eun48t/help_needed_whats_a_fast_preferably_nonvideo_way/)
- url: https://www.reddit.com/r/rails/comments/eun48t/help_needed_whats_a_fast_preferably_nonvideo_way/
---
Got an upcoming interview for a rails job that will involve quizzing on hands on knowledge that I'd really like to ace.

I've got a lot of experience in Rails, but less corporate knowledge over more recent times.

So I've got the basics very down pat and a lot of advance usage, but I'd like something that can show me the idiomatic, best of the best Rails techniques for large, enterprise applications.

Any tips?

Thanks
## [5][URI.escape is obsolete. Percent-encoding your query string](https://www.reddit.com/r/rails/comments/eul9tb/uriescape_is_obsolete_percentencoding_your_query/)
- url: https://www.reddit.com/r/rails/comments/eul9tb/uriescape_is_obsolete_percentencoding_your_query/
---
Have you encountered one of those warnings in your Ruby 2.7.0 project?
&gt; warning: URI.escape is obsolete warning: URI.encode is obsolete

Find out how to fix it --&gt; https://medium.com/arturtrzop/uri-escape-is-obsolete-percent-encoding-your-query-string-cd998c4e847c?source=friends_link&amp;sk=4850768e1dcf986890f70f8a24b47413
## [6][Help me choose](https://www.reddit.com/r/rails/comments/eumvd5/help_me_choose/)
- url: https://www.reddit.com/r/rails/comments/eumvd5/help_me_choose/
---
So I have to implement a project that is entirely Real-Time. Do you suggest Ruby or Node for this kind of  project?  
Thank you :D
## [7][What is the catch with Hatchbox compared to Heroku?](https://www.reddit.com/r/rails/comments/eu799q/what_is_the_catch_with_hatchbox_compared_to_heroku/)
- url: https://www.reddit.com/r/rails/comments/eu799q/what_is_the_catch_with_hatchbox_compared_to_heroku/
---
[Hatchbox](https://www.hatchbox.io/) seems really great in what they offer. You just supply the servers, give them a fixed price per month and you don't have to worry about all that DevOps stuff. For small but real production apps that use anthing beyond the classic Heroku hobby dyno this seems like a great deal. However, looking through this sub and the internet it looks like no one is really talking about them. Is there anything that I'm missing that makes them not as attractive as I'm making them out to be right now? I know that Heroku has their plugins which allow you to easily utilize any supported service and Hatchbox doesn't offer anyting similiar, but this surely can't be the sole reason. Does anyone have experience with Hatchbox and can tell a little bit about it or can maybe someone enlighten me on what I'm missing exactly? Because to me this seems like a really great service and I think it's strange that next to no one is talking about them.
## [8][Dynamically reference credentials... casting issue?](https://www.reddit.com/r/rails/comments/euf3dq/dynamically_reference_credentials_casting_issue/)
- url: https://www.reddit.com/r/rails/comments/euf3dq/dynamically_reference_credentials_casting_issue/
---
Hi,

I'm attempting to dynamically reference a plan\_id for a given product (plan names are associated with products in my model) without hard coding it. Current version of rails is 5.2.3 &amp; ruby 2.6.3.  As an example I have a plan called "carnivore" which is mapped to its corresponding id in the `credentials.yml.enc` file. However in my code I don't want to directly call `Rails.application.credentials.carnivore,` instead I'm looking to do something like:

`plan_name = 'Rails application credentials '+ @product.plan_type_name.downcase`

`connection_string = plan_name.parameterize(separator:".", preserve_case: true)`

calling `connection_string` will output the string `Rails.application.credentials.carnivore`, but won't actually reference the plan\_id in the encoded file. How can I change this so that rails realizes that I don't want the string itself, but a reference to the credentials?
## [9][ActiveInteractor v1.0.0 Release](https://www.reddit.com/r/rails/comments/eu4z1b/activeinteractor_v100_release/)
- url: https://www.reddit.com/r/rails/comments/eu4z1b/activeinteractor_v100_release/
---
Hey ruby friends!

Over the weekend I released v1.0.0 of ActiveInteractor, an implementation of the Command Pattern for Ruby with ActiveModel::Validations heavily inspired by the interactors gem. It comes with rich support for attributes, callbacks, and validations, and thread safe performance methods.

This update has some major improvements to organizers as well as rails QOL improvements and a lot more.  Please check it out, let me know what you think!

[https://github.com/aaronmallen/activeinteractor](https://github.com/aaronmallen/activeinteractor)

[https://medium.com/@aaronmallen/activeinteractor-8557c0dc78db](https://medium.com/@aaronmallen/activeinteractor-8557c0dc78db)

[https://github.com/aaronmallen/activeinteractor/wiki](https://github.com/aaronmallen/activeinteractor/wiki)

[https://rubygems.org/gems/activeinteractor](https://rubygems.org/gems/activeinteractor)

[https://www.rubydoc.info/gems/activeinteractor](https://www.rubydoc.info/gems/activeinteractor/)

&amp;#x200B;

Update: It should be noted though this is NOT the interactor gem by collective idea, this is inspired by the interactor gem by collective idea. The main difference between the two gems is ActiveInteractor supports ActiveSupport validation and callbacks for your interactor run.
## [10][Best tools/database for a text-based search](https://www.reddit.com/r/rails/comments/etsj6z/best_toolsdatabase_for_a_textbased_search/)
- url: https://www.reddit.com/r/rails/comments/etsj6z/best_toolsdatabase_for_a_textbased_search/
---
I need to build a full-text search page in a rails application. I brainstormed with my team to choose which tools we could use, especially the database. We stand between Elasticsearch and Postgres. But we are not sure which would be the best choice. 

Today we have few records, but the expectation is that this solution will support a large scale.

Any tips for a search based on full-text? Database, tools, articles?
## [11][Checking user logged in or not (Devise)](https://www.reddit.com/r/rails/comments/etxy78/checking_user_logged_in_or_not_devise/)
- url: https://www.reddit.com/r/rails/comments/etxy78/checking_user_logged_in_or_not_devise/
---
I planned for adding a front-page (not static) to my app. And now, I actually want it to be like this : 

If user not logged in, shows a message like "Please login to your account" with a link to login page. And If user is logged in, shows something belongs to the user. I actually could do the same with help of you guys, but it was in scaffolds and not a view/controller different from that scaffold. 

Thanks!
