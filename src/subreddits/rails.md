# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gek2rk/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gek2rk/personal_projects_show_off_your_own_project_andor/
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
## [2][Setting up an Automatic Book reader with Devise + Rails - 1](https://www.reddit.com/r/rails/comments/ggzr52/setting_up_an_automatic_book_reader_with_devise/)
- url: https://www.reddit.com/r/rails/comments/ggzr52/setting_up_an_automatic_book_reader_with_devise/
---
Hey guys,

I hate how much coding channels focus on a basic 'ToDo list / blog' so  i wanted wanted to share my progress in building an automatic book reader (not that good in making vids but meh). Probably will do several more videos until the end goal of having a tracker for how many books I've read.

&amp;#x200B;

 [https://www.youtube.com/watch?v=AnDLaaTXeWg&amp;feature=youtu.be](https://www.youtube.com/watch?v=AnDLaaTXeWg&amp;feature=youtu.be) 

&amp;#x200B;

Thx!
## [3][What is the most pleasant frontend stack of your choice?](https://www.reddit.com/r/rails/comments/gglkff/what_is_the_most_pleasant_frontend_stack_of_your/)
- url: https://www.reddit.com/r/rails/comments/gglkff/what_is_the_most_pleasant_frontend_stack_of_your/
---

## [4][Rails with React](https://www.reddit.com/r/rails/comments/ggvbcn/rails_with_react/)
- url: https://www.reddit.com/r/rails/comments/ggvbcn/rails_with_react/
---
Hi just learnt to build things rails as back end. Is there any good tut where I can learn front end with rails? I want to learn react and node js with rails.
## [5][How to add "{" and "}" in an url](https://www.reddit.com/r/rails/comments/ggn5t6/how_to_add_and_in_an_url/)
- url: https://www.reddit.com/r/rails/comments/ggn5t6/how_to_add_and_in_an_url/
---
Hi guys.

In my structured\_data (in &lt;head&gt;), I have this

    "potentialAction": {
     "@type": "SearchAction",
     "target": "&lt;%= search_url(search: {q: "{search_term_string}" }) %&gt;",
     "query-input": "required name=search_term_string"
     }

 

Watch "target". It show me the link in this way

It show me `https://www.mywebsite.com/search?utf8=%E2%9C%93&amp;search%5Bq%5D=%7Bsearch_term_string%7D`

It doesn't show { or } but %7B and %7D

I need to have this link `https://www.mywebsite.com/search?utf8=%E2%9C%93&amp;search%5Bq%5D={search_term_string}`

(following the [Google Guide](https://developers.google.com/search/docs/data-types/sitelinks-searchbox))

**How how to solve?**

I already try to fix using

`"{search_term_string}".html_safe` or `"&amp;#123;search_term_string&amp;#125;".html_safe` or `%({search_term_string})` or `%(&amp;#123;search_term_string&amp;#125;).html_safe`

**but nothing.**   
I was  thinking that there is the bug because it is an URL. So I should that I have to use URI.encode, but I'm not still sure about this (and the tests didn't go very well...).
## [6][Good places to start with IOS development?](https://www.reddit.com/r/rails/comments/ggeq6a/good_places_to_start_with_ios_development/)
- url: https://www.reddit.com/r/rails/comments/ggeq6a/good_places_to_start_with_ios_development/
---
Hey everyone! 

I have recently completed a full stack development course (Rails 6, Ruby, HTML, CSS, SQL, simple JavaScript) and feel relatively comfortable creating simple websites and hosting them on Heroku. 

I now wanted to look into IOS development, as I really want my next project (time management software) to also have a mobile application for users. 

Do you think React Native is the right way to go? I was unsure if it was wise to venture to far from the skills I recently learned, as I hoped this project would allow me to build upon and strengthen them.

Are there any good books / guides or tutorials that I could follow to get a coherent understanding?

Looking forward to reading your responses, thanks a lot in advance and have a great weekend! :)
## [7][Rspec: running certain tests without transaction](https://www.reddit.com/r/rails/comments/ggfxjj/rspec_running_certain_tests_without_transaction/)
- url: https://www.reddit.com/r/rails/comments/ggfxjj/rspec_running_certain_tests_without_transaction/
---
I really don't want to go back to database_cleaner but I currently have the issue that I cannot test a critical part that changes the transaction isolation level. It works in dev/prod but since rspec wraps all the tests in transactions I get `ActiveRecord::TransactionIsolationError: cannot set transaction isolation in a nested transaction`.

I'm afraid there is no way to not wrap tests in transactions on a selective basis or is there?

Alternatively I could move the transaction to a little wrapper class which I then don't test and just test the process under test conditions.

Any ideas on what to do in this case?

Cheers
## [8][How do I set up this ActiveRecord association where a model has both a "owner" and a "user"?](https://www.reddit.com/r/rails/comments/ggjwsa/how_do_i_set_up_this_activerecord_association/)
- url: https://www.reddit.com/r/rails/comments/ggjwsa/how_do_i_set_up_this_activerecord_association/
---
Hi all,

I've just created the following migration to allow a Venue to have an owner (by association with the users table) according to [these instructions](https://stackoverflow.com/questions/27809342/rails-migration-add-reference-to-table-but-different-column-name-for-foreign-ke).

    class AddOwnerToVenue &lt; ActiveRecord::Migration[6.0]
      def change
        add_reference :venues, :owner, foreign_key: { to_table: :users }
      end
    end

&amp; I have the following association set up in venue.rb:

    class Venue &lt; ApplicationRecord
      has_one :owner belongs_to :user
    end

Now for User, I have the corresponding association:

    class User &lt; ApplicationRecord
      has_many :venues
    end

But in order for my new owner association to work, I need to do something similar to this, but for it to look in the venue.owner column, not the venue.user.

Something like:

    class User &lt; ApplicationRecord
      has_many :venues, as: owner
    end

I can't find a way to do this. Hopefully someone has come across this kind of association before and can help me out?

Thanks.
## [9][Profiling Rails app that uses websockets](https://www.reddit.com/r/rails/comments/gg7dpw/profiling_rails_app_that_uses_websockets/)
- url: https://www.reddit.com/r/rails/comments/gg7dpw/profiling_rails_app_that_uses_websockets/
---
Hi. I'm new to Ruby and I am trying to profile a Rails app that uses websockets for streaming audio packets and other messages. I am interested in CPU time and wall clock time taken by all the methods in the app. Most of the tracing gems I see show the information for a single HTTP request, but it doesn't help in the case of websockets, since they remain alive even after a request. If it helps, 

How would you do this? How do I profile arbitrary points in code, like just do a "start" in the websocket open and "stop" in the websocket close, and generate a report? I tried ruby-prof, but it crashes with cryptic errors...
## [10][Best course resource for ruby and rails?](https://www.reddit.com/r/rails/comments/gfx93d/best_course_resource_for_ruby_and_rails/)
- url: https://www.reddit.com/r/rails/comments/gfx93d/best_course_resource_for_ruby_and_rails/
---
TL;DR - what course website should I get my company to pay for that has best rails content?


I’ve been using Rails for about a year. The purpose is business applications which is based around CRUD plus sometimes tricky client reqs. I get along fine in Rails but I feel there are some fundamentals and good habits which may be missing.  Also looking at improving things such as DRY, refactoring, StimulusReflex, and generally beautiful but maintainable code. 

I read this article and found it also interesting http://jeromedalbert.com/how-dhh-organizes-his-rails-controllers/. 

It wouldn’t hurt to improve my JS, HTML5, Bootstrap, etc. 

So looking for a nice resource which will cover these and be somewhat relevant to rails. 

I currently have access to a friends pluralsight. They also have a sale going on. Should I just get the company to purchase that?

Any other recommendations? Thanks.
## [11][Best way to convert data from a view or model into a spreadsheet](https://www.reddit.com/r/rails/comments/gfxr5k/best_way_to_convert_data_from_a_view_or_model/)
- url: https://www.reddit.com/r/rails/comments/gfxr5k/best_way_to_convert_data_from_a_view_or_model/
---
haven't done this before, was wondering if anyone had a nice article or video for using caxlsx to do so
