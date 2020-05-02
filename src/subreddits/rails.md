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
## [3][Best way to introduce/suggest Open ID Connect to enterprise vendors?](https://www.reddit.com/r/rails/comments/gc1a0i/best_way_to_introducesuggest_open_id_connect_to/)
- url: https://www.reddit.com/r/rails/comments/gc1a0i/best_way_to_introducesuggest_open_id_connect_to/
---
My rails app is the client (I plan to use omniauth-openid-connect gem). I need web single sign on authentication so that my rails app doesn't require another login when end users are coming from my customer's web portal. I want to make sure my customer's developers are onboard with implementing the authentication server side of open id connect (identity provider). 

What are the best ways to go about this process?

Edit: sorry title should say 'enterprise' not 'enterprise vendors'
## [4][Will using Ruby 2.6.5 with Rails 5 mess anything up?](https://www.reddit.com/r/rails/comments/gbohj9/will_using_ruby_265_with_rails_5_mess_anything_up/)
- url: https://www.reddit.com/r/rails/comments/gbohj9/will_using_ruby_265_with_rails_5_mess_anything_up/
---
Specifically, will using Ruby 2.6.5 with rails 5.0.0.1 mess anything up for me as I follow a book using Ruby 2.3.1? 

I started going through Agile Web Development with Rails 5 but wasn’t able to get the same versions installed. I installed Ruby 2.3.1 but was unable to install Rails with it. I had to instead use 2.6.5 and install Rails 5. Now if I switch back to 2.3.1, it says rails is not a recognized command. 

Will that mess up anything on my end or make it harder to follow the book? Any help would be really appreciated!
## [5][Devise + Wicked](https://www.reddit.com/r/rails/comments/gbp525/devise_wicked/)
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
## [6][[Help] Can I update multiple records with different attributes in one query?](https://www.reddit.com/r/rails/comments/gbm8fy/help_can_i_update_multiple_records_with_different/)
- url: https://www.reddit.com/r/rails/comments/gbm8fy/help_can_i_update_multiple_records_with_different/
---
I have multiple records I need to update with different attributes for each record. For example, I need to keep the order in tact and am keeping track through an attribute called 'index'. When one updates, I need to update the others to reflect their new position in the order. Is there any way to do so with one update or update_all, or do I have to go through them individually and update them one by one?

```
{id: 1, index: 1},
{id: 2, index: 2},
{id: 3, index: 3}
```

should become

```
{id: 3, index: 1},
{id: 1, index: 2},
{id: 2, index: 3}
```
## [7][[Podcast] The Ruby Blend: Parentheses and typosquatting](https://www.reddit.com/r/rails/comments/gbf5ag/podcast_the_ruby_blend_parentheses_and/)
- url: https://www.reddit.com/r/rails/comments/gbf5ag/podcast_the_ruby_blend_parentheses_and/
---
In this episode, Nate has to check his emails, Andrew uses parentheses, and Ron drowns in the DigitalOcean. We also find out why Nate was right (again). Also, what happened with typosquatting? Nate dives into the Manager vs. Maker's schedule. And why does Andrew have "needless paranoia?" \[Listen to this week's episode to find out about all these and more\]([https://fireside.fm/s/ouBAUjGy+4ZxpTNPN](https://fireside.fm/s/ouBAUjGy+4ZxpTNPN)).
## [8][Any good tutorials for building a Todo app with Rails as a backend?](https://www.reddit.com/r/rails/comments/gbih3f/any_good_tutorials_for_building_a_todo_app_with/)
- url: https://www.reddit.com/r/rails/comments/gbih3f/any_good_tutorials_for_building_a_todo_app_with/
---
Wanna try my hand at building a simple todo app with rails as a backend and html, css, and javascript as a frontend, with javascript handling the crud functionality. Any good tutorials for that?
## [9][Slightly Confused about Bundler](https://www.reddit.com/r/rails/comments/gb41qc/slightly_confused_about_bundler/)
- url: https://www.reddit.com/r/rails/comments/gb41qc/slightly_confused_about_bundler/
---
I mainly work with Elixir/Phoenix and there seems to be some similarities with how they handle dependencies but the more I worked in Rails it became evident that's not exactly the case.

In Phoenix you add dependencies to a mix file then run a command to install that dependency to your project.

In Rails you add dependencies to your gem file then run a command to install that dependency to your project but when working with a legacy rails project I had to install an older version of bundler directly in my project. When I run gem ls on my global directory(not project) I see they are all installed globally. What am I missing here? In phoenix everything is self contained like a virtual environment. In rails I have bundler installed globally and also in projects. Any clarification would be great. Thanks,
## [10][Where is rails generate?](https://www.reddit.com/r/rails/comments/gbakw3/where_is_rails_generate/)
- url: https://www.reddit.com/r/rails/comments/gbakw3/where_is_rails_generate/
---
Noob here, I've created a hello world app and ran it with webrick dev server BUT I want to host a webpage with a ruby app on nginx webserver ubuntu...

&amp;#x200B;

why dont I have rails generate? 

I can execute 'rails' but generate is not a switch for my rails command

&amp;#x200B;

how do I get my app.rb to show up on nginx webserver so I can execute if from the world wide webs?

&amp;#x200B;

Hetzner virtual cloud server dual cpu cores super smokin dawg
## [11][Twitter-bootstrap-rails](https://www.reddit.com/r/rails/comments/gb5myl/twitterbootstraprails/)
- url: https://www.reddit.com/r/rails/comments/gb5myl/twitterbootstraprails/
---
I know it's a bit of an outdated gem, but I tried it anyways on rails 6 to see if I could cheat my way through some stuff (also miss the sidebars). Looks absolutely retarded. You guys found any alternatives?
## [12][Bye, Rails Assets](https://www.reddit.com/r/rails/comments/gaefn2/bye_rails_assets/)
- url: https://www.reddit.com/r/rails/comments/gaefn2/bye_rails_assets/
---
Remove rails-assets.org dependencies from your Gemfile and move on.

[https://www.ramblingcode.dev/posts/bye\_rails\_assets/](https://www.ramblingcode.dev/posts/bye_rails_assets/)
