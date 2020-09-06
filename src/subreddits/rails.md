# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/igyvm1/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/igyvm1/personal_projects_show_off_your_own_project_andor/
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
## [2][Help with asset pipeline](https://www.reddit.com/r/rails/comments/ingy5q/help_with_asset_pipeline/)
- url: https://www.reddit.com/r/rails/comments/ingy5q/help_with_asset_pipeline/
---
Hey guys! I’m trying to figure out how to properly configure my asset pipeline. I’m having some trouble with this. In my production.rb file, I have config.assets.compile set to false and additionally, in my assets/config/manifest.js file, I have the following

//= link_tree ../images
//= link_directory ../stylesheets .css

However once I deploy the site, I get an error saying: “the asset &lt;image.jpg&gt; is not present in the asset pipeline” 

I’m accessing the images using: &lt;%= image_tag “image.jpg” %&gt;

It all works when I have config.assets.compile set to true but I know this is inefficient.

I would appreciate any help!
## [3][Ruby on Rails digest: 26 most popular repositories in July and August 2020](https://www.reddit.com/r/rails/comments/inkoxw/ruby_on_rails_digest_26_most_popular_repositories/)
- url: https://www.reddit.com/r/rails/comments/inkoxw/ruby_on_rails_digest_26_most_popular_repositories/
---
This is the second edition of Ruby/Rails digest. [This post](https://medium.com/@Iren.Korkishko/ruby-on-rails-digest-26-most-popular-repositories-in-july-and-august-2020-70ae593a5fe1?source=friends_link&amp;sk=477b6f0fb300b07f7fe068d12831c756) welcomes the most popular, most interesting, and useful Ruby on Rails repositories on GitHub in July and August!
## [4][multiple scopes and one def in controller](https://www.reddit.com/r/rails/comments/in3fbu/multiple_scopes_and_one_def_in_controller/)
- url: https://www.reddit.com/r/rails/comments/in3fbu/multiple_scopes_and_one_def_in_controller/
---
Hi guys, one question

if in home (view) I have this:

    
    &lt;div&gt;Today Offers&lt;/div&gt;
    &lt;%= render 'books/books', books: @query_of_the_day, skip_pages: true %&gt;

&amp;#x200B;

and in home\_controller.rb I have this

      def query_of_the_day
        day_of_the_week = Time.now.wday
        @query_of_the_day = if day_of_the_week == 0
                  Book.downloaded_today.includes(:screen, :online_links).limit(6).decorate
                ifelse day_of_the_week == 1
                  Book.rated_yesterday.includes(:screen, :online_links).limit(6).decorate
                ifelse day_of_the_week == 2
                  Book.over_50_readers.includes(:screen, :online_links).limit(6).decorate
                ifelse day_of_the_week == 3
                  Book.uploaded_yesterday.includes(:screen, :online_links).limit(6).decorate
                ifelse day_of_the_week == 4
                  Book.over_100_votes.includes(:screen, :online_links).limit(6).decorate
                else
                  Book.random.includes(:screen, :online_links).limit(6).decorate
                end

Where `downloaded_today`, `rated_yesterday`, `over_50_readers`, etc., are linked with a respective scope in scopes area of model/book.rb

&amp;#x200B;

My question is: Will it load ALL the `day_of_the_week` CASES (or all the  UNUSED scopes) when an user will visit the homepage?

...Because, in that case, my homepage will be very slow.
## [5][how do you deal with user timezones?](https://www.reddit.com/r/rails/comments/imym4y/how_do_you_deal_with_user_timezones/)
- url: https://www.reddit.com/r/rails/comments/imym4y/how_do_you_deal_with_user_timezones/
---
in a previous project i was saving the user timezone when signing up with an hidden field filled up with javascript to get the user's timezone and then used that timezone in views and mailers

but in my current project users can sign up with facebook so i can't add an hidden field

do i need to store the user timezone?

how would you localize datetimes in mailers without the user timezone stored?

and more generally how do you deal with user timezones?
## [6][Assets result in a 404](https://www.reddit.com/r/rails/comments/in0bqy/assets_result_in_a_404/)
- url: https://www.reddit.com/r/rails/comments/in0bqy/assets_result_in_a_404/
---
Hi. Tested whether my app would run on production (still on my local computer). However, all requests to them result in 404s.

What I did:

```sh
$ export RAILS_ENV=production
$ rails assets:precompile
$ rails s
```

At this point, neither my stylesheets nor my JavaScripts are loaded. I'm using the default configuration for `config/environments/production.rb`, meaning that the only assets settings is `compile`, which is set to `false`, which threw me off at first.

Tried turning it on, now the stylesheets loaded, but still no JavaScript. What is going on, I really can't figure it out?
## [7][What is your solution for page-specific javascript in rails?](https://www.reddit.com/r/rails/comments/imrqlk/what_is_your_solution_for_pagespecific_javascript/)
- url: https://www.reddit.com/r/rails/comments/imrqlk/what_is_your_solution_for_pagespecific_javascript/
---
I've been researching this throughout the day.  The most straight forward solution I've found is to put it in a content\_for in the view, and pull it into the layout right before the body tag, but that doesn't seem very 'unobtrusive.'

Is there is Rails standard, or community preferred method for this?
## [8][Podcasting platforms or tools built with Ruby on Rails?](https://www.reddit.com/r/rails/comments/imib6f/podcasting_platforms_or_tools_built_with_ruby_on/)
- url: https://www.reddit.com/r/rails/comments/imib6f/podcasting_platforms_or_tools_built_with_ruby_on/
---
Hey guys,

There are tons of Podcast platforms out there, but can any of you tell me if some of them are built with Rails?  

I'd be surprised if there were none.  Also, doesn't have to be the full platform, but maybe support tools like RSS feed builders or episode/video deployment tools.
## [9][Is there a stylistic reason NOT to with_indifferent_access every hash?](https://www.reddit.com/r/rails/comments/imhxf7/is_there_a_stylistic_reason_not_to_with/)
- url: https://www.reddit.com/r/rails/comments/imhxf7/is_there_a_stylistic_reason_not_to_with/
---
Hi,

Curious what the community thinks on this point. In many Rails codebases I've worked with `with_indifferent_access` is used _defensively_ to avoid the pain of `nil` returns.

```
# my_hash = { string_key: 'whoever did this must be a troll..' }

my_hash['string_key'] 
# =&gt; nil 
# argh!!!
```

Wondering if there are any stylistic reasons or situations that one would NOT convert all hashes `with_indifferent_access` and use a string key over a symbol (or mix of both) to semantically imply something about a hash value -- or perhaps this is just a quirk of Ruby?
## [10][API question](https://www.reddit.com/r/rails/comments/imgph5/api_question/)
- url: https://www.reddit.com/r/rails/comments/imgph5/api_question/
---
I have a decently successful iOS application that I am adding team/subscriptions to. Basically issuing licenses by the number of seats they want. 

To do so, I built a simple rails application that does three main functions: it allows a company to be created, to add employees (or seats) to the company roster, and to allow them the pay a monthly fee. 

Now I need to create an API that can speak with my iOS application to: 
1) verify that an account has paid (which my rails app w/ stripe integrated handles)
2) see what users (email addresses) are part of the account 

I am wondering what services you might use for this? Doorkeeper or something like this. Or JWT? I’m not that experienced with API’s and I’d love to simply allow api access for my iOS app server to then ingest the data from my rails app (in a secure way)
## [11][iCallendar issue](https://www.reddit.com/r/rails/comments/imdmnc/icallendar_issue/)
- url: https://www.reddit.com/r/rails/comments/imdmnc/icallendar_issue/
---
I am having some issues with icalendar gem,  it is adding events   in UTC+00 and I wanted it to be readjusted to match users  timezone, how can I do it?
