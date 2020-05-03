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
## [3][Ruby on Rails authorization using CanCanCan](https://www.reddit.com/r/rails/comments/gcoxya/ruby_on_rails_authorization_using_cancancan/)
- url: https://www.reddit.com/r/rails/comments/gcoxya/ruby_on_rails_authorization_using_cancancan/
---
Hi ruby family,

As an initiative to give back to the community, I have started writing a series of blogs on ruby and ruby on rails. Planning to create more content in the future to help share the knowledge. I just published a post about Authorization on Ruby on Rails using CanCanCan. Do check it out and let me know your thoughts.

[https://addytalks.tech/2020/05/03/ruby-on-rails-authorization-with-cancancan/](https://addytalks.tech/2020/05/03/ruby-on-rails-authorization-with-cancancan/)
## [4][Crucial Resources](https://www.reddit.com/r/rails/comments/gcjze8/crucial_resources/)
- url: https://www.reddit.com/r/rails/comments/gcjze8/crucial_resources/
---
Hey everyone! I'm somewhat new to programming (8 months in). Learning rails and so far its awesome! I had a bunch of local files and was told to put them in a repo for others to use. They have helped me tremendously. They are also great when I'm traveling or without internet as they are local. Feel free to use and would love if you had any helpful so called "cheat sheets" you would want to throw in there.

Cheers!

Heres the repo

[https://github.com/tylertomlinson/crucial\_resources](https://github.com/tylertomlinson/crucial_resources)
## [5][markdown user tag @](https://www.reddit.com/r/rails/comments/gcp5d2/markdown_user_tag/)
- url: https://www.reddit.com/r/rails/comments/gcp5d2/markdown_user_tag/
---
I edited my markdown redcarpet in this way

    class MarkdownRenderer &lt; Redcarpet::Render::HTML
      include Rails.application.routes.url_helpers
    
      def paragraph(text)
        text.gsub! (/@(\w+)/) do |match|
          user = User.find_by_username({$1})
          if user.exists?
            link_to match, user_path({$1})
          else
           link_to match, search_path(search: { author: {$1} })
          end
        end
        text
      end

But I have problems just with the first line and it says that no method for exists. Any tips?
## [6][How do I make my rails app run smoothly while mongodb isn't available?](https://www.reddit.com/r/rails/comments/gcoijo/how_do_i_make_my_rails_app_run_smoothly_while/)
- url: https://www.reddit.com/r/rails/comments/gcoijo/how_do_i_make_my_rails_app_run_smoothly_while/
---
Is there a better way to handle situations where the mongodb server goes down? I have a monolithic rails application that uses mongoid for certain tasks. 

However, from keeping my web app from failing/lagging, is there a better way I could handle the scenarios where mongo isn't available? 

Right now, all I could come up with was reducing the server selection timeout and handling the exception in every single place individually.
## [7][Blob column of model in Rails 6 with ActiveStorage](https://www.reddit.com/r/rails/comments/gciqy3/blob_column_of_model_in_rails_6_with_activestorage/)
- url: https://www.reddit.com/r/rails/comments/gciqy3/blob_column_of_model_in_rails_6_with_activestorage/
---
Hello there, how to generate a model with a column of blob. I have a db design of table "document" with the column of id:int, file:blob, status:enum and exp\_date:datetime. I am using postgresql
## [8][How to setup testing suite for front-end Javascript e2e?](https://www.reddit.com/r/rails/comments/gch5wl/how_to_setup_testing_suite_for_frontend/)
- url: https://www.reddit.com/r/rails/comments/gch5wl/how_to_setup_testing_suite_for_frontend/
---
I'm ready to dip my toes into front-end testing, but I'm a little confused on where to begin. 

From what I've read Jest and Puppeteer seem like solid choices, at least good enough for what I'm needing.

My question, though, has to deal with how my various *.js.erb files will fit into the mix. Is there anything I need to do in order to make the `ActiveRecord` interpolations "usable" in the sense of testing things after the `respond_to` *.js.erb renderings?

So many of the well-written articles I've read (like [this one](https://blog.arkency.com/testing-react-dot-js-components-with-jest-in-rails-plus-webpacker-plus-webpack-environment/)) use React. Are there any roadblocks I can expect if I'm only using "old-school" Javascript and jQuery? I've gone through a few tutorials that involve Webpacker/Webpack so I'm at least somewhat familiar with what's going on there.

Thanks in advance for any links, help, or insights you can offer!
## [9][Question about React and Rails.](https://www.reddit.com/r/rails/comments/gcjx1i/question_about_react_and_rails/)
- url: https://www.reddit.com/r/rails/comments/gcjx1i/question_about_react_and_rails/
---
I'm fairly new to Rails so bear with me!

I saw that you can create routes using React inside of a Rails application using the react-router-dom.

I know that a Rails app without React you usually create routes in routes.rb, but I'm seeing you can do all that inside of a App.js file and using the line `match '*path', to: 'pages#index', via: :all` within the routes.rb file.

What's the difference? Does creating routes using React allow for your app to be more dynamic? Is this way preferred when using React with Rails?

Thanks!
## [10][Best way to introduce/suggest Open ID Connect to enterprise vendors?](https://www.reddit.com/r/rails/comments/gc1a0i/best_way_to_introducesuggest_open_id_connect_to/)
- url: https://www.reddit.com/r/rails/comments/gc1a0i/best_way_to_introducesuggest_open_id_connect_to/
---
My rails app is the client (I plan to use omniauth-openid-connect gem). I need web single sign on authentication so that my rails app doesn't require another login when end users are coming from my customer's web portal. I want to make sure my customer's developers are onboard with implementing the authentication server side of open id connect (identity provider). 

What are the best ways to go about this process?

Edit: sorry title should say 'enterprise' not 'enterprise vendors'
## [11][Will using Ruby 2.6.5 with Rails 5 mess anything up?](https://www.reddit.com/r/rails/comments/gbohj9/will_using_ruby_265_with_rails_5_mess_anything_up/)
- url: https://www.reddit.com/r/rails/comments/gbohj9/will_using_ruby_265_with_rails_5_mess_anything_up/
---
Specifically, will using Ruby 2.6.5 with rails 5.0.0.1 mess anything up for me as I follow a book using Ruby 2.3.1? 

I started going through Agile Web Development with Rails 5 but wasn’t able to get the same versions installed. I installed Ruby 2.3.1 but was unable to install Rails with it. I had to instead use 2.6.5 and install Rails 5. Now if I switch back to 2.3.1, it says rails is not a recognized command. 

Will that mess up anything on my end or make it harder to follow the book? Any help would be really appreciated!
## [12][Devise + Wicked](https://www.reddit.com/r/rails/comments/gbp525/devise_wicked/)
- url: https://www.reddit.com/r/rails/comments/gbp525/devise_wicked/
---
Hey, I hope you're all safe.

I'm currently trying to create a smooth sign up form with devise and wicked but I'm having some trouble, since I use devise "Email only signup".

After a user signs up, they get redirected to the wicked controller but then I get the following error. "first argument in form cannot contain nil or be empty".

&amp;#x200B;

This is the critical part:

    class RegistrationsController &lt; Devise::RegistrationsController
    
      def after_inactive_sign_up_path_for(resource)
        user_registration_step_path(user_id: @user.id, id: :personal)
      end
    end

&amp;#x200B;

    class RegistrationStepsController &lt; ApplicationController
      include Wicked::Wizard
      before_action :set_user
      steps :personal
    
      def show
        render_wizard
      end
    
      def update
        render_wizard @user
      end
    
    private
      def set_user
        @user = current_user
      end
    end

personal.html.erb

    &lt;%= form_for @user, url: wizard_path, method: :put do |f| %&gt;
    &lt;% end %&gt;
