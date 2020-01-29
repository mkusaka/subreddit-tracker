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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/evmx0w/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/evmx0w/personal_projects_show_off_your_own_project_andor/
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
## [3][Build a Reddit Clone in Rails 6 [PART 2]](https://www.reddit.com/r/rails/comments/evg45k/build_a_reddit_clone_in_rails_6_part_2/)
- url: https://www.reddit.com/r/rails/comments/evg45k/build_a_reddit_clone_in_rails_6_part_2/
---
Hi guys, I've released part 2 of the new Reddit clone build in Ruby on Rails. I posted the first part of the build here about 1 week ago and it got some interest so this latest release might be helpful to anyone who enjoyed the previous video.

**Part 1:**

In this video I cover the setup of a new rails app, installing the devise gem and creating user accounts. Add functionality for users to create a new community/subreddit. Add functionality to create new posts for a community (using nested routes).

[https://www.youtube.com/watch?v=aD6JvHKNPPM](https://www.youtube.com/watch?v=aD6JvHKNPPM)

**Part 2:**

Add dropdown nav menu allowing users to view their profile, edit account settings and sign out. Extend devise to add first name, last name and username. Signed in users can subscribe to communities and create new posts within them. Add profile for users which shows some basic details and their recent posts. Add new comments table and allow subscribers to add comments to posts. Improve layout and design of our app frontend.

[https://www.youtube.com/watch?v=kSj3pCT3r6Y](https://www.youtube.com/watch?v=kSj3pCT3r6Y)

**Part 3 (to be confirmed):**

Add ability to upvote / downvote posts and order them in their community based on popularity. Add user karma based on number of upvotes received on posts.  


The build is part of a collection of MVP's that I have been building after working in the tech startup space for some time and assisting in building minimum viable products.

I post new web app builds on YouTube using Ruby on Rails 6 so feel free to check out the channel and add some recommendations for upcoming builds.
## [4][Rails Version Used In Jobs 🤔](https://www.reddit.com/r/rails/comments/evj708/rails_version_used_in_jobs/)
- url: https://www.reddit.com/r/rails/comments/evj708/rails_version_used_in_jobs/
---
Was thinking about hopping back to rails, was wondering what version is being used lately on the job market, if anyone has any insight would be appreciated ^^
Last version I picked up was 5.2
## [5][Routes For Portfolio Project](https://www.reddit.com/r/rails/comments/evmkco/routes_for_portfolio_project/)
- url: https://www.reddit.com/r/rails/comments/evmkco/routes_for_portfolio_project/
---
I was thinking here if I should use resources routes or make custom one for my portfolio projects 🤔🤔 
any insights ?
## [6][Weird Devise error when new users signup](https://www.reddit.com/r/rails/comments/evfoq0/weird_devise_error_when_new_users_signup/)
- url: https://www.reddit.com/r/rails/comments/evfoq0/weird_devise_error_when_new_users_signup/
---
OK, so I setup Devise on a Rails 6 app that I've created. Additionally, I followed the [email only signup](https://github.com/heartcombo/devise/wiki/How-To:-Email-only-sign-up) instructions.

In general, everything seems ok *except* that a bunch of the users appear to be getting logged out during the signup process which causes them to click the reset password link. An anonymized version of the production log is below.

Any ideas on what might be causing this? 

&amp;#x200B;

&gt;*From rotated app log* production.log  
13:52:05: new user  
13:52:07: signup  
13:52:08: /almost\_finished loads  
13:52:09: /newsletters loads? (can't say it's the same user for sure)  
13:52:13: email event received  
13:52:14: goes to /users/edit (involves users/registrations/edit.html.erb around this time)  
**13:52:15 signout happens (same user's IP from signup) -** "(No route matches \[GET\] "/users/sign\_out"):"  
13:52:15 starting /users/confirmation/new  
13:52:16 new post /users/confirmations with commit "**Send confirmation again**"  
13:52:16 mail event  
13:52:17 /users/sign\_in  
13:52:18 /newsletters redirect  
13:52:21 another "Send confirmation again"  
13:52:22 get /users/sign\_in  
13:52:22 /newsletters redirect  
13:52:22: /newsletters/9

In case it's helpful, the site is [Pew Pew Training](https://pewpewtraining.com).
## [7][Solopreneur owner of 20k/month Rails app wants to give inside view](https://www.reddit.com/r/rails/comments/ev2s0n/solopreneur_owner_of_20kmonth_rails_app_wants_to/)
- url: https://www.reddit.com/r/rails/comments/ev2s0n/solopreneur_owner_of_20kmonth_rails_app_wants_to/
---
Ten years ago I built a niche marketplace and have been living from the revenue ever since. It was designed to be a lifestyle business from the get-go so I spent a ton of time initially optimizing and automating everything.

I want to create a stream that gives a behind-the-scenes, nothing-hidden, look at what I've built and how I work, but I've no idea which topics are most interesting. If I learned one thing in creating my business, it's that it's all too easy to create something that no-one wants.

So what topics would you like to be covered?

Some ideas I had were

**Coding**

1. generating a maximally enjoyable development environment (e.g. seeding data, simulating cron, mirroring production as much as possible etc.)
2. removing brittleness from integration tests that run on circleci
3. dealing with the shitshow that is sales tax accounting across multiple currencies
4. detecting and recovering from production bugs asap
5. dealing with the real-world mess that is imperfect user input (e.g. when they type emails with a leading space or inconsistent capitalization; when they create a tag that is almost the same as a previous one — like E Guitar vs. Electric Guitar—and now your data is split across two areas)
6. discussing the 8+ year consequences of certain architectural/software design issues
7. streamlining massive amounts of config
8. multi-redundant systems of backup to prevent disaster
9. designing error messages and a logging strategy that speeds up recovery from errors
10. a tour of the most evil, insidious bugs I dealt with over the years (I keep a diary for them)
11. payment systems in-depth (refunds, errors etc.)
12. caching systems for performance

**Business/marketing**

1. how I use data to decide to add/remove a feature
2. AB testing a web app
3. technical SEO (microdata, site structure for internal links, google's tools, sitemaps, etc.) — I get 85% of my traffic (and therefore revenue) from SEO, so I know a thing or two
4. how I use JS and integration tests on all tracking code (critical to get right in my business)
5. auto-email systems to previous customers for extra sales
6. Adwords workflow to drive revenue
7. Analytics workflow to figure out what content working
8. Writing copy that gets sales (what worked for me vs. didn't)

\_If there's something you're interested in but don't see mentioned above, please do share.\_
## [8][How do you guys store constants that are model-specific and to be used for testing and validations?](https://www.reddit.com/r/rails/comments/evgsdl/how_do_you_guys_store_constants_that_are/)
- url: https://www.reddit.com/r/rails/comments/evgsdl/how_do_you_guys_store_constants_that_are/
---
Example: Customers can either be a "Person" or "Company", nothing else.

Right now I'm thinking of making a method in Customer.rb:

&gt;def self.types  
&gt;  
&gt;\["Person","Company"\]  
&gt;  
&gt;end

Then I refer to it as Customer.types in my validations and Rspec files. I'm doing this as an attempt to be DRY, and to make it easier to add to this array in the future if needed.

Alternatively I can store it as a constant in config/initializers, but this isn't truly a global constant since it is model-specific.

Edit: Also, how do you guys store global constants that are NOT model-specific?

Thoughts?
## [9][Has anyone recently setup a heroku with a custom subdomain with namecheap?](https://www.reddit.com/r/rails/comments/evjbde/has_anyone_recently_setup_a_heroku_with_a_custom/)
- url: https://www.reddit.com/r/rails/comments/evjbde/has_anyone_recently_setup_a_heroku_with_a_custom/
---
It has been a while since I setup a custom domain on heroku. The things I have done so far.

1. I added the custom domain to heroku using heroku domains:add test.example.com
2. Added a cname record on namecheap that points to the heroku app. I got the URL from the heroku domains:add command 

The DNS has propagated. But when I go to test.example.com in Chrome I get the following error

    This site can’t provide a secure connection
    test.example.com sent an invalid response.
    ERR_SSL_PROTOCOL_ERROR

Anyone know why this is happening? Do I have to setup SSL in the app? Or is it something else?
## [10][Editing form properties in Rails?](https://www.reddit.com/r/rails/comments/evfkwq/editing_form_properties_in_rails/)
- url: https://www.reddit.com/r/rails/comments/evfkwq/editing_form_properties_in_rails/
---
I have a form view, which is like this : 

    &lt;%= form_with(model: post, local: true) do |form| %&gt;
      &lt;% if post.errors.any? %&gt;
        &lt;div id="error_explanation"&gt;
          &lt;h2&gt;&lt;%= pluralize(post.errors.count, "error") %&gt; prohibited this post from being saved:&lt;/h2&gt;
    
          &lt;ul&gt;
            &lt;% post.errors.full_messages.each do |message| %&gt;
              &lt;li&gt;&lt;%= message %&gt;&lt;/li&gt;
            &lt;% end %&gt;
          &lt;/ul&gt;
        &lt;/div&gt;
      &lt;% end %&gt;
    
      &lt;div class="field"&gt;
        &lt;%= form.label :title %&gt;
        &lt;%= form.text_field :title %&gt;
      &lt;/div&gt;
    
      &lt;div class="field"&gt;
        &lt;%= form.label :text %&gt;
        &lt;%= form.text_area :text %&gt;
      &lt;/div&gt;
    
      &lt;div class="field"&gt;
        &lt;%= form.label :date %&gt;
        &lt;%= form.datetime_select :date %&gt;
      &lt;/div&gt;
    
      &lt;div class="actions"&gt;
        &lt;%= form.submit %&gt;
      &lt;/div&gt;
    &lt;% end %&gt;

And as in our country, we have a different calendar than the gregorian (and our calendar has a gem!) I got curious, how can I edit that "form.datetime\_select" to be in our calendar's format.
## [11][it's tdd tuesday where I'm at - how do i write a test for this?](https://www.reddit.com/r/rails/comments/evewi4/its_tdd_tuesday_where_im_at_how_do_i_write_a_test/)
- url: https://www.reddit.com/r/rails/comments/evewi4/its_tdd_tuesday_where_im_at_how_do_i_write_a_test/
---
  
suppose you'r working on a newish rails application, and you're tasked with adding bootstrap css to rails 6.  
in TDD style, how do you write a test first?   


can you write a test for this visual broken'ness (screenshot.png) ?  


https://preview.redd.it/364uuzcg6md41.png?width=800&amp;format=png&amp;auto=webp&amp;s=574c90db0119334d51b80d8e734a022ddf4bc3cc
## [12][RSpec allow to receive and return or raise an exception](https://www.reddit.com/r/rails/comments/evbw0h/rspec_allow_to_receive_and_return_or_raise_an/)
- url: https://www.reddit.com/r/rails/comments/evbw0h/rspec_allow_to_receive_and_return_or_raise_an/
---
I know you can mock an object's method and tell it to return different values depending if it was called by the first time, second time, etc., with `allow(object).to receive(:foo).and_return('a', 'b' ...)` and you can also tell it to raise an exception replacing `and_return` with `and_raise ...`

But is there a way to tell it to raise an exception at the first call and then to return a value at the second call?
