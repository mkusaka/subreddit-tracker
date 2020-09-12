# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/ipfxw8/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/ipfxw8/personal_projects_show_off_your_own_project_andor/
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
## [2][cache warming strategy recs](https://www.reddit.com/r/rails/comments/ira56v/cache_warming_strategy_recs/)
- url: https://www.reddit.com/r/rails/comments/ira56v/cache_warming_strategy_recs/
---
In my rails app, I generate a lot of dashboards which aggregates data on the order of GBs. using the pull based caching strategy won’t work for me because the cache miss would take about 10 seconds to repopulate the cache.

the data is also high dimensionality. the dashboards can group by multiple columns over user specified dates.

I’m doing all sorts of gymnastics in order to warm up the cache in a lot of different places in the app, but it is getting unruly to the point where i’m not even sure if a cache populating line of code is getting used in the app. 

are there any gems that help with this problem?
## [3][How to get the most out of GoRails as a newbie?](https://www.reddit.com/r/rails/comments/iqv8si/how_to_get_the_most_out_of_gorails_as_a_newbie/)
- url: https://www.reddit.com/r/rails/comments/iqv8si/how_to_get_the_most_out_of_gorails_as_a_newbie/
---
I managed to get a free year subscription of GoRails thanks to the Github Student Development Pack. I'm still learning Rails. I managed to get the basics of Rails down thanks to Michael Hartl Rails book.


I'm looking to learn more about Rails and follow along step by step projects, but for the most part, GoRails seems to be just a collection of mini videos on how stuff works. I'm not really sure what to start watching or if the site has any "how to build X". I'm still exploring the site though. 

Anyways, what's the best way to use this site to get the most out of my free 1 year subscription and make myself a better Rails Dev?

Thanks.
## [4][Real life application with Rails and Stimulus Reflex compared to the SPA approach](https://www.reddit.com/r/rails/comments/iqqnvg/real_life_application_with_rails_and_stimulus/)
- url: https://www.reddit.com/r/rails/comments/iqqnvg/real_life_application_with_rails_and_stimulus/
---
Hi Rails and JS devs, 

&amp;#x200B;

I never really understood the hype around SPAs and "modern JavaScript" development workflow.

It always seemed so much complicated to me for basically no real advantages.

I will talk about it more precisely in a future blog post. 

So I created a repo with the two stacks to compare how much they are different in a real life scenario. 

&amp;#x200B;

Turbolinks was a game changer and Stimulus Reflex adds the last thing to create reactive apps.

&amp;#x200B;

Here's the repo: [https://github.com/guillaumebriday/modern-datatables](https://github.com/guillaumebriday/modern-datatables)

&amp;#x200B;

Let me know how do you build reactive apps these days!
## [5][The best and the simplest authentication method for Rails API](https://www.reddit.com/r/rails/comments/iqrjxj/the_best_and_the_simplest_authentication_method/)
- url: https://www.reddit.com/r/rails/comments/iqrjxj/the_best_and_the_simplest_authentication_method/
---
Hey everyone,   
I've been trying to complete a rails API backend and VueJS frontend movie browser app where you can signup, login, change the password, and follow movies, genres and stars then get recommendations according to your follows for a technical test.  
It's really simple when you think about it but I'm having a crisis with authentication methods. I'm not experienced with Rails API. So please give me some insights and tips about where to look for a simple authentication method.  
First I've tried Knock and JWT, which went well it was working like a charm on localhost until I switch to production and push the project to Heroku and Netlify.  It was working with login/signup but when I do follow actions it was giving 401 unauthorized even tough there was a Bearer token in Headers.   
Then on the same project, I switched back to devise and simple\_auth\_token which causes huge errors.  
Then I decided to code the backend from scratch and I've tried the followings: 

  
https://github.com/madeindjs/api\_on\_rails  
https://armaizadenwala.com/blog/social-media-app/  
[https://gorails.com/series/how-to-build-apis-with-rails](https://gorails.com/series/how-to-build-apis-with-rails)  
None of the above really helped me and I'm still struggling to finish the authentication so I can land a job.   
What is the ultimate resource for a rails API?

Thank you.
## [6][Managing Rails cache low-level concurrency](https://www.reddit.com/r/rails/comments/iqy3i3/managing_rails_cache_lowlevel_concurrency/)
- url: https://www.reddit.com/r/rails/comments/iqy3i3/managing_rails_cache_lowlevel_concurrency/
---
Let's say I have the following Worker:

```
class ModelWorker
  include Sidekiq::Worker

  def perform(id)
    cached_request = do_request
    Model.find(id).process_using_request(cached_request)
  end

  def do_request
    Rails.cache.fetch('key') do
      Client.new.fetch_data
    end
  end
end
```

Ideally, what I would like to happen, is that the code wrapped by the `fetch` call is only called once. Is this possible if I fire, simultaneously, tons of `ModelWorker` at the same time? Or will all jobs call `Client.new.fetch_data`?
## [7][Anyone know of a product that acts like an internal stackoverflow?](https://www.reddit.com/r/rails/comments/iqn0mm/anyone_know_of_a_product_that_acts_like_an/)
- url: https://www.reddit.com/r/rails/comments/iqn0mm/anyone_know_of_a_product_that_acts_like_an/
---
I’d like to be able to log any blockers that devs come across in our team, from small annoying things to larger issues.

I’ve had countless times where one of us will encounter an issue that someone else in our team has already found a solution to but it hasn’t been logged anywhere.

Are there any products out there that do this sort of thing?
## [8][Authorizing social media even when logged in](https://www.reddit.com/r/rails/comments/iqink2/authorizing_social_media_even_when_logged_in/)
- url: https://www.reddit.com/r/rails/comments/iqink2/authorizing_social_media_even_when_logged_in/
---
A social media analytics app will get data from many sources like facebook, twitter, instagram,etc. A user who logs in my analytics app will see the option to connect these 3 social media to periodically fetch from those sources.  


How is that handled via devise? I dont want code I want to understand the flow. My confusion is that I am already logged in then how can I connect those 3 social media? Is that still oauth?
## [9][Web-Based bug tracking system.](https://www.reddit.com/r/rails/comments/iqffxx/webbased_bug_tracking_system/)
- url: https://www.reddit.com/r/rails/comments/iqffxx/webbased_bug_tracking_system/
---
Hello people. I need some help finding out gems that i can use to build bug tracking system.  
I'm new to rails and programming. I'm looking to expand my knowledge and try to build some good projects.
## [10][I need a Rails developer](https://www.reddit.com/r/rails/comments/iqvbd3/i_need_a_rails_developer/)
- url: https://www.reddit.com/r/rails/comments/iqvbd3/i_need_a_rails_developer/
---
Hello! I am working on a restaurant reservations web application. Already have a react developer who is taking care of doing SSR, app routing, and different kind of changes in front end, that require to commit changes in Rails too, so I want somebody to take care of the Rails part through coordinating the Rails code with the React one. Also, there are 2 bugs to be fixed and finally deployment in Heroku. If you want to have a look at the codebase, please send me a message so I can send it to you in order to have a clearer idea. My budget is $100, since I am a college student who will have other expenses related to this start up, but would love to work with the rails developer and hire for other tasks even after the web app will be launched. PM me if you're interested. Thank you.
## [11][Dockerizing and deploying rails on AWS](https://www.reddit.com/r/rails/comments/iq112b/dockerizing_and_deploying_rails_on_aws/)
- url: https://www.reddit.com/r/rails/comments/iq112b/dockerizing_and_deploying_rails_on_aws/
---
Hi, please excuse me if this is a simple question, but I am new to Docker and AWS, but not to rails.

I want to get my Docker skills better, and I want to get my AWS skills better!  I want to deploy my application using Docker on AWS. I dont know where to start.

I have a rails API only application (created using the API flag), which now has a Dockerfile and docker-compose.yml file. There are so many Docker tutorials and I understand how Docker works in its basic form - but here is what I don't understand:

1) if the database is created on a separate container, how does the application connect with it, and how do I deploy it on AWS? does it need a separate cluster? Does it need a separate Dockerfile? is docker-compose only used for development, but not for production? 

I don't expect anyone to give me a long explanation, but if you could point me in the right direction of some good learning materials, or give me a brief overview, or tell me I'm barking up the right/wrong tree, that would be great. 

&amp;#x200B;

Thanks so much
