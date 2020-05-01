# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/g616hm/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/g616hm/personal_projects_show_off_your_own_project_andor/
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
## [2][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/gauf3h/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/gauf3h/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [3][[Podcast] The Ruby Blend: Parentheses and typosquatting](https://www.reddit.com/r/rails/comments/gbf5ag/podcast_the_ruby_blend_parentheses_and/)
- url: https://www.reddit.com/r/rails/comments/gbf5ag/podcast_the_ruby_blend_parentheses_and/
---
In this episode, Nate has to check his emails, Andrew uses parentheses, and Ron drowns in the DigitalOcean. We also find out why Nate was right (again). Also, what happened with typosquatting? Nate dives into the Manager vs. Maker's schedule. And why does Andrew have "needless paranoia?" \[Listen to this week's episode to find out about all these and more\]([https://fireside.fm/s/ouBAUjGy+4ZxpTNPN](https://fireside.fm/s/ouBAUjGy+4ZxpTNPN)).
## [4][Slightly Confused about Bundler](https://www.reddit.com/r/rails/comments/gb41qc/slightly_confused_about_bundler/)
- url: https://www.reddit.com/r/rails/comments/gb41qc/slightly_confused_about_bundler/
---
I mainly work with Elixir/Phoenix and there seems to be some similarities with how they handle dependencies but the more I worked in Rails it became evident that's not exactly the case.

In Phoenix you add dependencies to a mix file then run a command to install that dependency to your project.

In Rails you add dependencies to your gem file then run a command to install that dependency to your project but when working with a legacy rails project I had to install an older version of bundler directly in my project. When I run gem ls on my global directory(not project) I see they are all installed globally. What am I missing here? In phoenix everything is self contained like a virtual environment. In rails I have bundler installed globally and also in projects. Any clarification would be great. Thanks,
## [5][Where is rails generate?](https://www.reddit.com/r/rails/comments/gbakw3/where_is_rails_generate/)
- url: https://www.reddit.com/r/rails/comments/gbakw3/where_is_rails_generate/
---
Noob here, I've created a hello world app and ran it with webrick dev server BUT I want to host a webpage with a ruby app on nginx webserver ubuntu...

&amp;#x200B;

why dont I have rails generate? 

I can execute 'rails' but generate is not a switch for my rails command

&amp;#x200B;

how do I get my app.rb to show up on nginx webserver so I can execute if from the world wide webs?

&amp;#x200B;

Hetzner virtual cloud server dual cpu cores super smokin dawg
## [6][Twitter-bootstrap-rails](https://www.reddit.com/r/rails/comments/gb5myl/twitterbootstraprails/)
- url: https://www.reddit.com/r/rails/comments/gb5myl/twitterbootstraprails/
---
I know it's a bit of an outdated gem, but I tried it anyways on rails 6 to see if I could cheat my way through some stuff (also miss the sidebars). Looks absolutely retarded. You guys found any alternatives?
## [7][Bye, Rails Assets](https://www.reddit.com/r/rails/comments/gaefn2/bye_rails_assets/)
- url: https://www.reddit.com/r/rails/comments/gaefn2/bye_rails_assets/
---
Remove rails-assets.org dependencies from your Gemfile and move on.

[https://www.ramblingcode.dev/posts/bye\_rails\_assets/](https://www.ramblingcode.dev/posts/bye_rails_assets/)
## [8][Deploy a new Rails app to AWS Lambda using Lamby (my notes)](https://www.reddit.com/r/rails/comments/gafh0n/deploy_a_new_rails_app_to_aws_lambda_using_lamby/)
- url: https://www.reddit.com/r/rails/comments/gafh0n/deploy_a_new_rails_app_to_aws_lambda_using_lamby/
---
Link: [https://gist.github.com/joshuap/4ab915f6b94366f570fcb63f38660262](https://gist.github.com/joshuap/4ab915f6b94366f570fcb63f38660262)

Should you do this? I don't know. \*Can\* you do this? Yep. :) ([Here is my production endpoint](https://2ayxcm8ff3.execute-api.us-west-1.amazonaws.com/production/)--not sure how long this will be up.)

A few unanswered questions I have:

1. What about max function size? Will this work with a large Rails app?
2. What about cold starts? Even a boilerplate Rails app takes a long time to boot.
3. Why is Docker so slow on macOS 😭

Am I missing anything?
## [9][Rails 5.2 to 6 migration could be more challenging than you think. Let's talk about it.](https://www.reddit.com/r/rails/comments/gab35q/rails_52_to_6_migration_could_be_more_challenging/)
- url: https://www.reddit.com/r/rails/comments/gab35q/rails_52_to_6_migration_could_be_more_challenging/
---
You know how migrating to a new version of rails is taking away resource and time from answering customer needs. It is not just the framework but all the gems and dependencies  now with rails  6 and webpacker as defualt packing it is even more time consuming. I've been on a rails6+webpack+rails ujs migration for the last 8-9 days for a not very small app. We were postponing migrations for a few months. We are far from finish.

I just spend 6 hours as a dependency was using a removed jQuery method and jQuery now comes through webpacker and yarn. The whole experience of migrating a real app to a kind of a new stack (npm modules) including turbolinks and stimulus is tedious. It is still relatively easy compared to other stacks but is still taking a lot of time. A lot of configurations, sprockets on the top, feature specs failing. I spend 2 days figuring out how to connect sprockets with webpack since not everything could be migrated to webpack instantly.

The end result is nice and powerful, but reaching it is taking us away from answering customer needs and the whole point of the software we write is to answer customer needs not to fight with a miss behavoir of a configuration.

So I thought I could write articles to also help our colleagues and have a reference for other apps. I wrote about 4 and then I just started writing notes for articles. I am at about 22 notes for new articles that should be developed. 

But articles are not the solution I thought. I've been through hundreds of GitHub issues and I've looked through tens of pull request up until now to understand what is happening. Every app is different and every app faces similar but different challenges when migrating 

So I thought - nobody should go through this. The process is not bringing value - you just need to have it magically migrated and you should focus on talking with customers and developing features to address their pains. You learn a couple of things a long the way while.migrating, but you already know them. There is not much new except for the many new configurations for different things. You need to just receive a pull request with everything resolved in a couple of days and then a meeting with discussion on what was change and why. 

We've already lost about 120 hours on this and we are far from finish. Many people share that this takes months. Months of migration especially when webpack, sprockets and not trivial JavaScript is involved.

This is what I thought I could help with. You might need to talk with somebody, or you might need to work with somebody or to delegate to someone to resolve it and give you a new branch for you to just merge and then go over all the problems. 

Message me if you are feeling the pain. Share a comment here to discuss.
## [10][Interesting article](https://www.reddit.com/r/rails/comments/gawia2/interesting_article/)
- url: https://www.reddit.com/r/rails/comments/gawia2/interesting_article/
---
 [https://blog.cloud66.com/home-buying-selling-platform-orchard-deploys-kubernetes-to-aws-with-maestro/](https://blog.cloud66.com/home-buying-selling-platform-orchard-deploys-kubernetes-to-aws-with-maestro/)
## [11][Lightweight Auth solution for Rails/React/Redux app?](https://www.reddit.com/r/rails/comments/gac0qd/lightweight_auth_solution_for_railsreactredux_app/)
- url: https://www.reddit.com/r/rails/comments/gac0qd/lightweight_auth_solution_for_railsreactredux_app/
---
I'm building a dashboard using Rails 6 and using the included webpacker to build the frontend in React and Redux for state. This app will have a logon but no sign up. The admin will maintain user access since it's an internal application. However, there will be a front public facing aspect of the app but no sign up or CRUD action. My goto would be Devise but wondering if this is too 'heavy handed' Or should I just use it and take what I need?
## [12][How do I handle bot requests?](https://www.reddit.com/r/rails/comments/ga9teg/how_do_i_handle_bot_requests/)
- url: https://www.reddit.com/r/rails/comments/ga9teg/how_do_i_handle_bot_requests/
---
So, I have this VPS with my Rails application that is open to internet, and I keep receiving a lot of requests from bots that are trying to find loopholes, for example:

`[81edb7aa-08a1-4204-8429-5ffbf1d1a374] ActionController::RoutingError (No route matches [GET] "/home.asp"):` 

`[d916cd40-c474-4bf4-85aa-8cfa8538201e] ActionController::RoutingError (No route matches [GET] "/login.cgi"):` 

`[cf8bc9a3-b4e5-4df1-9c04-fc1fb595266c] ActionController::RoutingError (No route matches [GET] "/dana-na/auth/url_default/welcome.cgi"):`

`[d033854a-15a1-4e7f-b66f-f0008727b847] ActionController::RoutingError (No route matches [GET] "/remote/login"):`

 `[e79c5fe3-1d08-49f0-8031-a2129724ea5b] ActionController::RoutingError (No route matches [GET] "/index.asp"):`

 `[9851fade-d0f5-4a52-992e-81f824705874] ActionController::RoutingError (No route matches [GET] "/SQlite/main.php"):`

And so on.

They keep coming from random IP's, so I though I could create a rule on my nginx.conf that block any request that results in an error or something like this.

So, how do you guys handle this bruteforce?

I'm using Nginx + Passenger, but I'm no expert on configuring it.
