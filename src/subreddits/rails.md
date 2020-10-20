# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/jbmadq/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/jbmadq/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Where do people host their Rails 6 apps in 2020?](https://www.reddit.com/r/rails/comments/jeo0xy/where_do_people_host_their_rails_6_apps_in_2020/)
- url: https://www.reddit.com/r/rails/comments/jeo0xy/where_do_people_host_their_rails_6_apps_in_2020/
---
I’m a returning Rails developer from the 4.2 days. Been out a couple years. 

I used to use AWS Elastic Beanstalk, because we had other AWS resources that our app used.

But here I am in October 2020, trying to get even the Rails 6 “yay you’re on Rails” default app running on Beanstalk, and it’s an endless slog of deploy, fail, check the logs, google the error, try a fix, repeat.

Is Beanstalk the wrong solution now? How do people host Rails apps on AWS? I’ve never used containers - is that the approach I should take?
## [3][Setting up a dev environment database to work on app locally?](https://www.reddit.com/r/rails/comments/jekdzx/setting_up_a_dev_environment_database_to_work_on/)
- url: https://www.reddit.com/r/rails/comments/jekdzx/setting_up_a_dev_environment_database_to_work_on/
---
Say you join a project and you need to set up your local dev environment. You install Rails, bundle install etc., but how exactly does the database portion work?

What I mean is, there's a production database that's pretty large with a couple dozen tables. When I work on the app locally, it won't run without connecting to a database for logins etc. Should I clone the whole production database and reproduce it locally, or what is the best practice here?
## [4][data model for charting stock-prices/changes?](https://www.reddit.com/r/rails/comments/jedezf/data_model_for_charting_stockpriceschanges/)
- url: https://www.reddit.com/r/rails/comments/jedezf/data_model_for_charting_stockpriceschanges/
---
Stocks are often charted out with prices and their given date of change. 

IF you have a single stock and wanted to show the price changing over time, how would you model that in your database? 

Sounds like TONS of data...
## [5][Count No of Words From an Html String Blog](https://www.reddit.com/r/rails/comments/jendjr/count_no_of_words_from_an_html_string_blog/)
- url: https://www.reddit.com/r/rails/comments/jendjr/count_no_of_words_from_an_html_string_blog/
---
Hi, guys hope you are fine,

  
I'm implementing an API in which I receive a Blog from the Front end in an HTML String with base 64 converted image and save this as a text in my API.

But now I want to count the No of Words of a Blog but not the base 64 converted image words as it is a huge string.   


Any kind of help will be appreciated?
## [6][Question: What are you guys using for centralized logging that is not papertrail app](https://www.reddit.com/r/rails/comments/jehhqd/question_what_are_you_guys_using_for_centralized/)
- url: https://www.reddit.com/r/rails/comments/jehhqd/question_what_are_you_guys_using_for_centralized/
---
Looking for an open sources solution to collect logs (centralize logging). Please suggest if you are using something in production. What has worked out well for you. 

\- graylog  
\- logstash  
\- fluend  
\- something else (hopefully)
## [7][[Help] Seeding database adds duplicate files to aws s3 bucket](https://www.reddit.com/r/rails/comments/jehzr0/help_seeding_database_adds_duplicate_files_to_aws/)
- url: https://www.reddit.com/r/rails/comments/jehzr0/help_seeding_database_adds_duplicate_files_to_aws/
---
So I was attempting to troubleshoot an issue with aws when I noticed that every time I run rails db:setup, it sends in duplicate files, but with random string names. The issue I initially had was fixed but now I have extra files in the bucket. This is a few examples of the names: 

1gk9BGENPhY9N8b1DXBbjZ7j 

UcvvBR9N3WDxPCx4yXT5uXhk

g2mx2c5XpgV3dMWUFj1Gof2R


windows 10 wsl2 running Ubuntu 18-04
Some of my settings are:

* ruby 2.5.1p57
* Rails 5.2.4.4
* npm 6.4.1
* Files were loaded directly to the bucket from aws upload, rather than pushing them up through the rails app
* the duplicate files are not default-ly public

The seeds file has the setup something like this:

    movie1 = Movie.create({title: "Splash page", description: "some description", rating: "pg-13", director: "Me", video_url: ""})

    clip1 = open('https://pulled_straight_from_the_Object_url.amazonaws.com/test-video.webm')
    cover_pic1 = open('https://pulled_straight_from_the_Object_url.amazonaws.com/test-cover_picture.jpg')

    movie1.movie_clip.attach(io: clip1, filename: 'test-video.webm ')
    movie1.cover_art.attach(io: cover_pic1, filename: 'test-cover_picture.jpg')

I tried googling for something similar but couldn't find anything. It isn't a problem, since I can delete the files. I just want to know why it's happening.
## [8][Tool for tracking clicks on buttons in Ruby on Rails Apps](https://www.reddit.com/r/rails/comments/jeh5qv/tool_for_tracking_clicks_on_buttons_in_ruby_on/)
- url: https://www.reddit.com/r/rails/comments/jeh5qv/tool_for_tracking_clicks_on_buttons_in_ruby_on/
---
We have an app that does various things, but it's not so clear to us which users are clicking which buttons.

We'd like to use this to build a profile of user activity.

Most solutions from google analytics don't seem to let us know which user(s) clicked which things. Each logged in user corresponds to a `user_id` in our DB, and what we'd need to know precisely is which user_id clicked or didn't click certain buttons, or browsed certain pages, etc.
## [9][Adding users &amp; authorisation to a Rails app | Live-streaming Ruby on Rails](https://www.reddit.com/r/rails/comments/je15gu/adding_users_authorisation_to_a_rails_app/)
- url: https://www.reddit.com/r/rails/comments/je15gu/adding_users_authorisation_to_a_rails_app/
---
I've done a couple more live-streams on the Sundae Club channel, where each week we live-stream building an app and look at features from a learner's perspective. In the last couple of weeks we've looked at:

[https://www.youtube.com/watch?v=b-J1TuimRtk](https://www.youtube.com/watch?v=b-J1TuimRtk) - Adding users to a Rails app

[https://www.youtube.com/watch?v=77Z_u2QDTZA](https://www.youtube.com/watch?v=77Z_u2QDTZA) - Adding user permissions to a Rails app

Both videos follow roughly the same format where we look at some of the options that are available and then implement one of them on the app and try to see it in action. Chapters are listed on each videos to help you navigate and save you having to watch the whole stream.

New streams are each Sunday at 10.30am (UK time) and available any time on YouTube.
## [10][Combine Dynamic Params Into Field](https://www.reddit.com/r/rails/comments/je6xqc/combine_dynamic_params_into_field/)
- url: https://www.reddit.com/r/rails/comments/je6xqc/combine_dynamic_params_into_field/
---
I have some drop down select fields that are dynamic. They are displayed on the page based on how many options are selected. How do I combine the results of them in value field on create. 

So if the model is variant\_values, and within the form I have this:

        &lt;% u/product_variants.each do |product_variant| %&gt;
        &lt;div class="col"&gt;
          &lt;div class="form-group"&gt;
          &lt;%= form.label product_variant.name %&gt;
          &lt;%= form.select :value, product_variant.value.split(","), { name: 'variant[product_variant_ids][]', class: 'form-control' } %&gt;
          &lt;/div&gt;
        &lt;/div&gt;
        &lt;% end %&gt;

How would I combine them into one value field when I create the varaint\_value? right now it only submits the last product\_varaint field to be displayed.  I am think of combining them with a comma or hyphen between the values so I can separate them later.
## [11][Trying to generate an app with a specific version, but Rails is updating itself.](https://www.reddit.com/r/rails/comments/jeg6be/trying_to_generate_an_app_with_a_specific_version/)
- url: https://www.reddit.com/r/rails/comments/jeg6be/trying_to_generate_an_app_with_a_specific_version/
---
I'm trying to generate an app like so:

`rails _6.0.1_ new myapp`

But it updates itself when generating the app:

`Fetching rails 6.0.3.4`    

`Installing rails 6.0.3.4`

And then when I do 

`bin/rails about`

It of course shows Rails 6.0.3.4

Why is this happening and how can I prevent it?
