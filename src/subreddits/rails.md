# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/foqc07/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/foqc07/personal_projects_show_off_your_own_project_andor/
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
## [2][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/ftl6xc/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/ftl6xc/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [3][Properly Exporting JS Files per Page Only in Rails 6](https://www.reddit.com/r/rails/comments/ftimgw/properly_exporting_js_files_per_page_only_in/)
- url: https://www.reddit.com/r/rails/comments/ftimgw/properly_exporting_js_files_per_page_only_in/
---
I'm currently migrating an application from sprockets to webpacker under Rails 6. I have a situation where I name my js files after the page I attach it to for some crud modules. So I have:

* model_as/Show.js
* model_bs/Show.js

The style of each file is something like `var Show = (function() { })();` using the revealing pattern in vanilla js. The last line of each file is `window.Show = Show`.

In `ModelA`'s show page, I include it using `javascript_pack_tag 'model_as/Show.js'` and in`ModelB`'s show page, I include it using `javascript_pack_tag 'model_bs/Show.js'`. 

When I visit the show page for `A`, it works fine. However when I visit the show page for `B`, it seems to still have cache-ed `A`'s `Show.js` file. So I think it's an issue with `window.Show = Show`. 

My question is, given a setup like that, what's the best practice to export js files that we load only on a per page basis?
## [4][Launching Your First Freelance Application (A Tutorial Beginner Series)](https://www.reddit.com/r/rails/comments/ftgvwg/launching_your_first_freelance_application_a/)
- url: https://www.reddit.com/r/rails/comments/ftgvwg/launching_your_first_freelance_application_a/
---
Hey everyone. I've been holding off on sharing this resource I've been working on until I made it past the 2 week mark and had episodes under my belt.

I've started working on a series of Ruby on Rails tutorials called my "20 in 20" challenge, where I plan to build 20 different projects in 20 weeks, from varying scopes and difficulties.

Week 2 involved what was initially a very simple mock freelance gig. You went into your local diner, and Frank, the owner, asked if you could make a website for him if he would pay you. It's your first project like this, so you don't have a lot of experience but you say yes.

I covered Frank's requirements in episode 1, and then was met with several comments including one about one of my viewers whose friend just had to close his restaurant because of the Covid-19 outbreak. I decided to increase the scope, so that he could use this project to actually help his friend by creating an online platform for him.

The current scope has covered:

* Creating the initial pages, adding bootstrap, and how to source &amp; edit the images used in the background
* Attempting to implement a search field by the menu item names, as well as by whether the meals are: Gluten free, Peanut free, vegan, vegetarian, kosher, etc..
* Converting the menu into a list of products that you can add to an order
* The ability to manage your order on the shopping cart page
* Pushing to production and the challenges this poses with Active Storage

Part 6 just went up which covered deploying to Digital Ocean, using Digital Ocean Spaces for Active Storage, configuring the Capistrano files for handling Webpacker so that our bootstrap would work, and more. Hopefully this can serve as the basis for showing how to add Stripe, move to HTTPS with Cloudflare and a custom domain, and more.

I thought I would share it on this subreddit because I see others sharing their resources, and I feel like this might be a series for people looking to see a project go from an initial concept to a real world application (Helping out someone whose friend actually could use this project in practice, rather than theory) and I plan to continue this project well into what is now week 3, and soon to be week 4, by implementing the remaining features such as an admin panel, stripe checkout, and the customer reviews.

I'll include a link to the current episode as of this post. You can find all of the playlists by week, or just the master 20in20 playlist, on my channel. I really hope this helps someone. It was a ton of work to create, so I can only imagine what it would be like for someone without even my limited experience.

[https://www.youtube.com/watch?v=NOa3WEuimJc](https://www.youtube.com/watch?v=NOa3WEuimJc)
## [5][Production Rails - live-coding video series](https://www.reddit.com/r/rails/comments/ft90y5/production_rails_livecoding_video_series/)
- url: https://www.reddit.com/r/rails/comments/ft90y5/production_rails_livecoding_video_series/
---
I've been working on this series for a while, been holding back on sharing it until I had a decent number of videos uploaded.

Basically the idea is that I code a Rails application and show absolutely everything including when I get hopelessly stuck and don't know what to do.

I also include the stuff most tutorials don't include like setting up the production environment and writing tests.

[Here's the YouTube link.](https://www.youtube.com/watch?v=hBnzhi2kQ80&amp;list=PLYzYlrv4LHcgIuymTahL3UIlzS8fMTVOz&amp;index=2&amp;t=2s)

Hope you like it.
## [6][Deployment problem](https://www.reddit.com/r/rails/comments/ftjxv2/deployment_problem/)
- url: https://www.reddit.com/r/rails/comments/ftjxv2/deployment_problem/
---
I made updated my personal blog and website which used to be hosted on heroku, but this time I deployed to digitalocean using capistrino. In production on localhost it works more or less fine but when i visit the site at the hosted address it has a routing problem. The main pages controller works fine but the controllers for Blog and Profile don't work. I have no idea why, production.log is also empty so I am at a loss as how to fix it. I'm not sure the rules on posting code so contact me if you have any advice pleaase

&amp;#x200B;

Thanks
## [7][best way write inside pdf &amp; pptx](https://www.reddit.com/r/rails/comments/ftmag5/best_way_write_inside_pdf_pptx/)
- url: https://www.reddit.com/r/rails/comments/ftmag5/best_way_write_inside_pdf_pptx/
---
I need write in  pdf &amp; pptx  . use some data in side pdf &amp; pptx .  
Data like 

* Table with some style
* Add image and add above some text 
* add some Graph like pie graph
## [8][if trying to pick up Rails, AppAcademyOpen is pretty good](https://www.reddit.com/r/rails/comments/ft0zm2/if_trying_to_pick_up_rails_appacademyopen_is/)
- url: https://www.reddit.com/r/rails/comments/ft0zm2/if_trying_to_pick_up_rails_appacademyopen_is/
---
I've used Odin and benefited there, but if you feel you want further practice and engrain ideas more, AppAcademyOpen and its demos have been really nice, you have to expand the menu, but there are lots of lessons and modules such as:

https://imgur.com/a/9FVa6FK

Just a recommendation for those looking to get better. I've really enjoyed it.
## [9][Adding a rescue clause to middleware for a specific](https://www.reddit.com/r/rails/comments/ft9uyv/adding_a_rescue_clause_to_middleware_for_a/)
- url: https://www.reddit.com/r/rails/comments/ft9uyv/adding_a_rescue_clause_to_middleware_for_a/
---
Hi all,  
   I'm a newbie when it comes to Rails so looking for some help. I am looking to add some middleware that detects a specific request (in this case assets) and sends data to a logger should that request fail (for example, a 404).   


    class AssetLogger
      def initialize app
        @app = app
      end
    
      def call env
        path = env['PATH_INFO']
        @status, @headers, @response = @app.call(env)
        rescue Exception
          if path.include? "/asset_path/"
            logstuff
          end
    
    
        [@status, @headers, @response]
      end
    end
    
    config.middleware.use AssetLogger

When I test this out and try to load an asset that doesn't exist I get an ActionController::RoutingError and it does not go in to my rescue block. So clearly this approach will not work. I am wondering if anyone has any insight on a better way to do this?
## [10][Sorry dumb question: which database is schema_migrations supposed to be in?](https://www.reddit.com/r/rails/comments/ft3i9a/sorry_dumb_question_which_database_is_schema/)
- url: https://www.reddit.com/r/rails/comments/ft3i9a/sorry_dumb_question_which_database_is_schema/
---
dev or prod?

We've made changes to some configs and for some reason I can't remember if it is in the default location and the docs aren't helping me answer this question.   Thanks.
## [11][Can we see the generated metaprogramming method](https://www.reddit.com/r/rails/comments/fswsxp/can_we_see_the_generated_metaprogramming_method/)
- url: https://www.reddit.com/r/rails/comments/fswsxp/can_we_see_the_generated_metaprogramming_method/
---
I want to know if there is possible to check if a method that is generated by metaprogramming in rails can be viewed as a normal method definition. 

Suppose in the meta-programming definition I have defined a few methods which have a few conditional statements and for a particular such method, I want to see the finally constructed method.
## [12][Need help with AWS Cognito, S3, IAM, STS confusion](https://www.reddit.com/r/rails/comments/fsy7g7/need_help_with_aws_cognito_s3_iam_sts_confusion/)
- url: https://www.reddit.com/r/rails/comments/fsy7g7/need_help_with_aws_cognito_s3_iam_sts_confusion/
---
I'm using AWS Cognito User Pool for authentication to my Ruby on Rails app. I've also created an Identity Pool for the User Pool, with the intention that logged in users are able to read files from a specific S3 bucket. However, I'm having trouble access objects in S3 as the logged in user. I have the tokens for the logged in user (`id_token`, `access_token`), but I don't know where to supply them.

Following the [sample code](https://docs.aws.amazon.com/sdk-for-ruby/v3/developer-guide/s3-example-get-bucket-item.html) from the Ruby SDK doc only allows me to access S3 as a shared user, not the logged in user:
```ruby
s3 = Aws::S3::Resource.new(region: aws_region)
obj = s3.bucket('my-bucket').object('my-item')
obj.get(response_target: './my-code/my-item.txt')
```

Then I found [this doc](https://docs.aws.amazon.com/AmazonS3/latest/dev/AuthUsingTempSessionTokenRuby.html) about using temporary credentials, but I also don't know how use the user's tokens anywhere. I've also tried STS' [AssumeRoleWithWebIdentity](https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRoleWithWebIdentity.html) that allows me to pass in the `id_token`, and used the Identity Pool's `role_arn`, but I'm getting `AccessDenied` error. I'm confused of the relationship between STS, IAM, and Cognito Identity Pool. Can anyone help?

Note: I've also enabled S3 access logging and it should log the Cognito user who access the files.
