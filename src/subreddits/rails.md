# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/f9t9kq/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/f9t9kq/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/fgx7fz/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/fgx7fz/personal_projects_show_off_your_own_project_andor/
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
## [3][How do I use Devise in normal AND API mode](https://www.reddit.com/r/rails/comments/fgwc24/how_do_i_use_devise_in_normal_and_api_mode/)
- url: https://www.reddit.com/r/rails/comments/fgwc24/how_do_i_use_devise_in_normal_and_api_mode/
---
Hi guys,  
I want to use Devise for creating users with mobile app via REST API but I also want use it in classic fashion.  
So in API I want to return JSON and appropriate HTTP codes. Also I want to issue JWT token for authentication for mobile app.  
On web I want normal devise functionality.  
How I should structure controllers, router and other stuff? I cant find any good resources about this.  
Thank you guys.
## [4][Question: Rails 5.2 Testing Javascript and Manipulating the Session](https://www.reddit.com/r/rails/comments/fgmzfk/question_rails_52_testing_javascript_and/)
- url: https://www.reddit.com/r/rails/comments/fgmzfk/question_rails_52_testing_javascript_and/
---
Hey everyone,

I've been at this problem for too long, so I'm reaching out for a steer on it.  
I have some code that works in the browser, but I can't work out how to write a test for it.

Basically, I have a Bootstrap 4 modal in a view. There's some Coffeescript that will open the modal if the user has come from a particular HTTP\_REFERRER.

I cannot find a test framework that supports JS as well as supporting manipulating request headers. I've looked at MiniTest, RSpec and Capybara.

I've seen a StackOverflow post where BrowserMob was suggested - but it seems to have been abandoned, and I'm loathe to add another dependency for this one feature.

Any suggestions? I'm probably overlooking something obvious.  The aim is be able to manipulate the HTTP\_REFERRER header in a test to cause the JS to fire and show the modal.

Thanks in advance.
## [5][Rails with nginx/unicorn - 2gb per process??](https://www.reddit.com/r/rails/comments/fgkwuf/rails_with_nginxunicorn_2gb_per_process/)
- url: https://www.reddit.com/r/rails/comments/fgkwuf/rails_with_nginxunicorn_2gb_per_process/
---
Hi all,

My unicorn processes are taking up about 2GB each of memory. The internet tells me this is rather abnormal and high, but I'm not sure what size app people are benchmarking against. 

I have a production consumer web app which does about a million people per month in traffic. My AWS instance is 16gb memory. I have my unicorn config set to `worker_processes 5`. 

When I run a command like `ps -aux --sort -rss | head -n 10` the top six are the unicorn master and 5 workers, each taking up between 2 and 2.5GB of memory. 

I've already gone down several rabbit holes here, but before following any one path I would love to get a sense if this is at all normal for the kind of traffic my website does. 

Thanks!
## [6][Deployment Questions](https://www.reddit.com/r/rails/comments/fgnxqg/deployment_questions/)
- url: https://www.reddit.com/r/rails/comments/fgnxqg/deployment_questions/
---
Hello, I have a few questions about your preferred way to deploy. I typically just launch an EC2, ssh in, configure capistrano, nginx, route 53 &amp; deploy.

I want to explore some other areas of deployment to try and speed this process up without increasing monthly cost too much. What i am thinking is to start using docker, deploy with beanstalk, RDS for a DB but i have one last area to figure out. Redis.   


ElastiCache is the price of a small EC2. Using this service could double the bill which isn't the goal. Maybe i don't understand the pricing the model. If i turn one on, am i charged $0.05/hr like a EC2 or is it on-demand? My other thought is to put redis into the docker container. My fear is something like this: 50 jobs in queue, deploy, queue is gone. Is this possible? Maybe i don't understand how the deployment works nor redis? Would the queue be gone if i put redis in the container?   


Regardless, what's your preferred way deploying RoR while not being backed by a corporate budget? Open to exploring other options!
## [7][Migrate a Rails Project from Paperclip to ActiveStorage](https://www.reddit.com/r/rails/comments/fgiswf/migrate_a_rails_project_from_paperclip_to/)
- url: https://www.reddit.com/r/rails/comments/fgiswf/migrate_a_rails_project_from_paperclip_to/
---
I've recently started a blog and the first article I wrote was based around my experience migrating one of my personal projects from Paperclip to ActiveStorage.  Sharing here in hopes someone finds it useful!

[https://www.kevin-custer.com/blog/migrate-a-rails-project-from-paperclip-to-active-storage/](https://www.kevin-custer.com/blog/migrate-a-rails-project-from-paperclip-to-active-storage/)
## [8][Single Sign In across Rails app and Vue app](https://www.reddit.com/r/rails/comments/fglms1/single_sign_in_across_rails_app_and_vue_app/)
- url: https://www.reddit.com/r/rails/comments/fglms1/single_sign_in_across_rails_app_and_vue_app/
---
I have a RoR website that handles authentication with Devise and is hosted at [www.domain.com](https://www.domain.com).  I have another app (Vue, no rails) that is hosted at [subdomain.domain.com](https://subdomain.domain.com) and using axios to send a sign in request to my RoR website with username and password.  The Vue app receives an `auth_token` as a response and sends it in the header with each subsequent request to ensure it belongs to a user.

This all works fine, but if the user was to then sign into the RoR app, the Vue app does not get notified and the Vue app's token will no longer be valid, resulting in the need to sign in again.

I've looked through a lot of single-sign-on solutions and while they allow you to sign in with 1 app, it doesn't seem like they discuss much about how to inform the other apps when an `auth_token` is updated.

If my Vue app does not have direct database access, how can it remain signed in after the user signs into the other app?
## [9][people dealing with multiple clients/apps,](https://www.reddit.com/r/rails/comments/fgp4ap/people_dealing_with_multiple_clientsapps/)
- url: https://www.reddit.com/r/rails/comments/fgp4ap/people_dealing_with_multiple_clientsapps/
---
how do you keep a seperate bash history?  
how do you retain history across reboots/close-tabs   


im on latest macos with not so latest bash
## [10][Is there any part of Rails you wished you understood better?](https://www.reddit.com/r/rails/comments/fgcy4s/is_there_any_part_of_rails_you_wished_you/)
- url: https://www.reddit.com/r/rails/comments/fgcy4s/is_there_any_part_of_rails_you_wished_you/
---
I'm currently putting together a post about how the [Action Dispatch Request object](https://github.com/rails/rails/blob/1ef8c60dfcab2f0594b468088680fe1054d17452/actionpack/lib/action_dispatch/http/request.rb) works under the hood. Is there any part of Rails you wish you understood better? I'd be happy to write about it. [This is my latest about Active Storage](https://medium.com/@liroy).
## [11][Procedure to generate site map for rails app on Heroku?](https://www.reddit.com/r/rails/comments/fgi65i/procedure_to_generate_site_map_for_rails_app_on/)
- url: https://www.reddit.com/r/rails/comments/fgi65i/procedure_to_generate_site_map_for_rails_app_on/
---
Looking to generate a sitemap to submit to google search for quicker indexing, but haven't been able to find any comprehensive walkthroughs. Any links or tutorials would be appreciated. Thanks!
## [12][How to define custom exceptions errors?](https://www.reddit.com/r/rails/comments/fgmseu/how_to_define_custom_exceptions_errors/)
- url: https://www.reddit.com/r/rails/comments/fgmseu/how_to_define_custom_exceptions_errors/
---
Hello everyone! I do hope my post finds everyone well. I would like to know how can I define a custom exception error? My controllers returns standard errors when an invalid request is provided. However, I would like to define my own exception messages with other HTTP error codes in some controllers. 

E.g: 

    {
      status: "error",
      message: "Whoops! No content here",
      code: 204
    }

What can I do to achieve this? Thank you very much in advance.
