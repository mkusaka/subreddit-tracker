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
## [3][Bye, Rails Assets](https://www.reddit.com/r/rails/comments/gaefn2/bye_rails_assets/)
- url: https://www.reddit.com/r/rails/comments/gaefn2/bye_rails_assets/
---
Remove rails-assets.org dependencies from your Gemfile and move on.

[https://www.ramblingcode.dev/posts/bye\_rails\_assets/](https://www.ramblingcode.dev/posts/bye_rails_assets/)
## [4][Deploy a new Rails app to AWS Lambda using Lamby (my notes)](https://www.reddit.com/r/rails/comments/gafh0n/deploy_a_new_rails_app_to_aws_lambda_using_lamby/)
- url: https://www.reddit.com/r/rails/comments/gafh0n/deploy_a_new_rails_app_to_aws_lambda_using_lamby/
---
Link: [https://gist.github.com/joshuap/4ab915f6b94366f570fcb63f38660262](https://gist.github.com/joshuap/4ab915f6b94366f570fcb63f38660262)

Should you do this? I don't know. \*Can\* you do this? Yep. :) ([Here is my production endpoint](https://2ayxcm8ff3.execute-api.us-west-1.amazonaws.com/production/)--not sure how long this will be up.)

A few unanswered questions I have:

1. What about max function size? Will this work with a large Rails app?
2. What about cold starts? Even a boilerplate Rails app takes a long time to boot.
3. Why is Docker so slow on macOS 😭

Am I missing anything?
## [5][Rails 5.2 to 6 migration could be more challenging than you think. Let's talk about it.](https://www.reddit.com/r/rails/comments/gab35q/rails_52_to_6_migration_could_be_more_challenging/)
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
## [6][Help getting this recently create app with rails 4.2.3 getting deployed on heroku](https://www.reddit.com/r/rails/comments/gap36j/help_getting_this_recently_create_app_with_rails/)
- url: https://www.reddit.com/r/rails/comments/gap36j/help_getting_this_recently_create_app_with_rails/
---
Before anything, I'm a newbie in rails or anything related to web development. I'm taking this course about rails and its kind outdated this current section. Its an app that uses milia and devise. Unfortunately, milia only works with rails 5.1 and the course I'm taking is uses rails 4.2.3 and their outdated versions of milia and devise gems. I tried before to do things with Rails 5.1 but I had some many problems that I restarted things so many times and now I trying to code with the downgrade of things. Its not the best choice I know, but I'm feeling that wasting time struggling everyday to get things worked out and nothing change.

So, I wanna give it a shot doing this with rails 4 and get the better from the course. And if nothing works out, I'll try again with rails 5...

The app I'm trying to deploy on heroku I hadn't code any functionality yet. It just a new app and I just set a home page. Nothing else.

I had to downgrade the rails to 4.3.2 and ruby to 2.2.10 to get this at least deployed to heroku. The things works fine in the local environment (using aws cloud9) but I have and error on heroku. Does not show anything, just say to check the logs which I couldnt trace the error. 

You can check out [here](https://imgur.com/a/Itt5nuz) the heroku logs 

My [github repo](https://github.com/Gregory280/saas-project) 

Important informations:

I set puma as my web server. I followed the heroku instructions and seems fine this.

In my gem file I have a group production with gem pg 0.18.4 version and rails\_12factor 

For development I'm using sqlite3.

I kind think I had traced where the error might be, I think it is something with the pg gem and activerecord, but I don't have a clue how to solve it.
## [7][Lightweight Auth solution for Rails/React/Redux app?](https://www.reddit.com/r/rails/comments/gac0qd/lightweight_auth_solution_for_railsreactredux_app/)
- url: https://www.reddit.com/r/rails/comments/gac0qd/lightweight_auth_solution_for_railsreactredux_app/
---
I'm building a dashboard using Rails 6 and using the included webpacker to build the frontend in React and Redux for state. This app will have a logon but no sign up. The admin will maintain user access since it's an internal application. However, there will be a front public facing aspect of the app but no sign up or CRUD action. My goto would be Devise but wondering if this is too 'heavy handed' Or should I just use it and take what I need?
## [8][How do I handle bot requests?](https://www.reddit.com/r/rails/comments/ga9teg/how_do_i_handle_bot_requests/)
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
## [9][App that responds and to API requests and send requests to an external API](https://www.reddit.com/r/rails/comments/gacrrt/app_that_responds_and_to_api_requests_and_send/)
- url: https://www.reddit.com/r/rails/comments/gacrrt/app_that_responds_and_to_api_requests_and_send/
---
I’ve just started a new project that will be mostly API-based but with a weird twist because the app will act as a client and server: it respond to requests and it sends the same requests to an external end-point, using the same protocol.

The app will be an intermediary between a shop and a wholesaler and it will forward orders.   
Simplifying things, it will be something like this:

    Request: shop -&gt; myapp -&gt; wholesaler
    Response: wholesaler -&gt; myapp -&gt; shop

The shop will be authenticated on myapp using JWT and myapp will be authenticated on the wholesaler also with JWT.

The API used is the same and I would almost be able to do a “forward” (if that's possible) because I would just need to manipulate the token and add a string for the order number.

Example

    # From shop to myapp
    POST /api/orders
    Authorization: Bearer &lt;shop_jwt&gt;
    { order_number: "1234"
      order_items: [...] }
    
    # From myapp to wholesaler
    POST /api/orders
    Authorization: Bearer &lt;mayapp_jwt&gt;
    { order_number: "MYAPP-1234"
      order_items: [...] }

This is something new for me and I don't know how to approach this kind of problem. Could you point me in some direction?

I should also add that I'll need to store the orders that will pass back and forth through myapp, although that's the easy part
## [10][Rails 6 MiniTest &amp; FactoryBot: extra records existing in my tests](https://www.reddit.com/r/rails/comments/gadsy0/rails_6_minitest_factorybot_extra_records/)
- url: https://www.reddit.com/r/rails/comments/gadsy0/rails_6_minitest_factorybot_extra_records/
---
Hey all, I'm pretty new to MiniTest and figured I'd give it a shot with a new codebase. I keep running into having additional records in my test database at random.

I've double checked my factories, I don't use any setup blocks, and I haven't disabled transactional tests.

I'm honestly kind of at a loss.

Anyone have any ideas as to what my cause the test DB to have mystery guests like this?
## [11][Working around a library that doesn't have direct Ruby support](https://www.reddit.com/r/rails/comments/gaat92/working_around_a_library_that_doesnt_have_direct/)
- url: https://www.reddit.com/r/rails/comments/gaat92/working_around_a_library_that_doesnt_have_direct/
---


I'm currently building an API in Rails to control devices registered with an Azure IoT Hub.

Now, in the Azure documentation, I only see Python/Node/.Net guides to control the end devices.

I require Rails for certain libraries, and hence used it to build my API. But now I've hit this road block.

 I don't see any official Rails support. I only require c2d messaging 

The only other way I can think of with my limited knowledge is to deploy a small Python/Flask app on the same server and have my Rails code make HTTP requests to it, locally. 

My question : Is this a bad approach? If yes, is there a better one? 

Because I imagine this being a common problem, where some small functionality doesn't have a gem, and is only available in other languages officially.
## [12][How to read in .gifv files into Active Storage](https://www.reddit.com/r/rails/comments/gaeg79/how_to_read_in_gifv_files_into_active_storage/)
- url: https://www.reddit.com/r/rails/comments/gaeg79/how_to_read_in_gifv_files_into_active_storage/
---
I'm trying to load a bunch of .gifv files into Active storage, but I can't seem to figure out how to go about doing this.  I have a solution for .gif:

    image.picture_img.attach(io: open(img_picture), filename: "#{img_name}.gif", content_type: "image/gif")

But the same thing doesn't work when I change .gif to .gifv

Any ideas?
