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
## [3][PigCI Now Is Open Source 🎁🎉🥳](https://www.reddit.com/r/rails/comments/gcwvv4/pigci_now_is_open_source/)
- url: https://www.reddit.com/r/rails/comments/gcwvv4/pigci_now_is_open_source/
---
Main Website: [https://pigci.com/](https://pigci.com/)

Source Code : [https://github.com/PigCI/App](https://github.com/PigCI/App)

video : [https://youtu.be/uodK1rY8J2c](https://youtu.be/uodK1rY8J2c)
## [4][Anyone know of a Rails/Ruby Discord?](https://www.reddit.com/r/rails/comments/gd38tl/anyone_know_of_a_railsruby_discord/)
- url: https://www.reddit.com/r/rails/comments/gd38tl/anyone_know_of_a_railsruby_discord/
---
Anyone apart of any discord server that’s about Rails or Ruby? I think that’s be a fun place to connect!
## [5][How do you check type of file in ActiveStorage ?](https://www.reddit.com/r/rails/comments/gd9q4g/how_do_you_check_type_of_file_in_activestorage/)
- url: https://www.reddit.com/r/rails/comments/gd9q4g/how_do_you_check_type_of_file_in_activestorage/
---
Hey anyone, do you use ActiveStorage for multiple type files? How do you check that type ? In my case I jus want to check if that is images or videos
## [6][Query Problem Possible on ActiveRecord or Low Level Query?](https://www.reddit.com/r/rails/comments/gd4rap/query_problem_possible_on_activerecord_or_low/)
- url: https://www.reddit.com/r/rails/comments/gd4rap/query_problem_possible_on_activerecord_or_low/
---
I have table A, B and C.

A has many B and C is a category of B (B has an a\_id and B has a c\_id).

Let's say I want to query all A's and its corresponding B's based on selected C's. B has a date field so it is a possible that a record in A can have multiple B's of the same C. Example:

A with id 1 has 2 B records (1, 2) with dates January 2020 and February 2020 as well as the same C\_id 1. If I select C with id 1, it should give me record A with id 1 and only B with id 2 since it has the latest date value.

Wondering if this is possible with ActiveRecord or do I have to do a manual query?
## [7][Ruby on Rails authorization using CanCanCan](https://www.reddit.com/r/rails/comments/gcoxya/ruby_on_rails_authorization_using_cancancan/)
- url: https://www.reddit.com/r/rails/comments/gcoxya/ruby_on_rails_authorization_using_cancancan/
---
Hi ruby family,

As an initiative to give back to the community, I have started writing a series of blogs on ruby and ruby on rails. Planning to create more content in the future to help share the knowledge. I just published a post about Authorization on Ruby on Rails using CanCanCan. Do check it out and let me know your thoughts.

[https://addytalks.tech/2020/05/03/ruby-on-rails-authorization-with-cancancan/](https://addytalks.tech/2020/05/03/ruby-on-rails-authorization-with-cancancan/)
## [8][Markdown redcarpet and user tag (with link_to)](https://www.reddit.com/r/rails/comments/gcs4ms/markdown_redcarpet_and_user_tag_with_link_to/)
- url: https://www.reddit.com/r/rails/comments/gcs4ms/markdown_redcarpet_and_user_tag_with_link_to/
---
Hi guys, I improved my markdown in this way to create a link to the user page

&amp;#x200B;

    class MarkdownRenderer &lt; Redcarpet::Render::HTML
      include Rails.application.routes.url_helpers
    
      def paragraph(text)
        text.gsub! (/@(\w+)/) do |match|
          user = User.find_by_username($1)
          if user != nil
            link_to match, user_path($1)
          else
            link_to match, search_path(search: { author: $1 })
          end
        end
        text
      end

But it says me 

     No route marches {: action=&gt;"show", : controller=&gt;"users", :locale=&gt;"Freank"} missing required keys: [:id]

how is it possible? `$1` is my username.

I also tried to replace 

            link_to match, user_path($1)

with 

            link_to match, user_path(user.username)

but I have the same problem.

How to fix?
## [9][How do I make my rails app run smoothly while mongodb isn't available?](https://www.reddit.com/r/rails/comments/gcoijo/how_do_i_make_my_rails_app_run_smoothly_while/)
- url: https://www.reddit.com/r/rails/comments/gcoijo/how_do_i_make_my_rails_app_run_smoothly_while/
---
Is there a better way to handle situations where the mongodb server goes down? I have a monolithic rails application that uses mongoid for certain tasks. 

However, from keeping my web app from failing/lagging, is there a better way I could handle the scenarios where mongo isn't available? 

Right now, all I could come up with was reducing the server selection timeout and handling the exception in every single place individually.
## [10][Crucial Resources](https://www.reddit.com/r/rails/comments/gcjze8/crucial_resources/)
- url: https://www.reddit.com/r/rails/comments/gcjze8/crucial_resources/
---
Hey everyone! I'm somewhat new to programming (8 months in). Learning rails and so far its awesome! I had a bunch of local files and was told to put them in a repo for others to use. They have helped me tremendously. They are also great when I'm traveling or without internet as they are local. Feel free to use and would love if you had any helpful so called "cheat sheets" you would want to throw in there.

Cheers!

Heres the repo

[https://github.com/tylertomlinson/crucial\_resources](https://github.com/tylertomlinson/crucial_resources)
## [11][Blob column of model in Rails 6 with ActiveStorage](https://www.reddit.com/r/rails/comments/gciqy3/blob_column_of_model_in_rails_6_with_activestorage/)
- url: https://www.reddit.com/r/rails/comments/gciqy3/blob_column_of_model_in_rails_6_with_activestorage/
---
Hello there, how to generate a model with a column of blob. I have a db design of table "document" with the column of id:int, file:blob, status:enum and exp\_date:datetime. I am using postgresql
## [12][Question about React and Rails.](https://www.reddit.com/r/rails/comments/gcjx1i/question_about_react_and_rails/)
- url: https://www.reddit.com/r/rails/comments/gcjx1i/question_about_react_and_rails/
---
I'm fairly new to Rails so bear with me!

I saw that you can create routes using React inside of a Rails application using the react-router-dom.

I know that a Rails app without React you usually create routes in routes.rb, but I'm seeing you can do all that inside of a App.js file and using the line `match '*path', to: 'pages#index', via: :all` within the routes.rb file.

What's the difference? Does creating routes using React allow for your app to be more dynamic? Is this way preferred when using React with Rails?

Thanks!
