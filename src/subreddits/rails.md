# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gvu083/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gvu083/personal_projects_show_off_your_own_project_andor/
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
## [2][List of 211 remote jobs hand-picked from "Hacker News: Who is hiring?"](https://www.reddit.com/r/rails/comments/h9vihg/list_of_211_remote_jobs_handpicked_from_hacker/)
- url: https://www.reddit.com/r/rails/comments/h9vihg/list_of_211_remote_jobs_handpicked_from_hacker/
---
This list contains 211 remote jobs and you can filter them by location or skills.

Here I would like to share the entire remote jobs list from the big list of opportunities. All these are 100% remote jobs not just allowed to work from home during this crisis. These are 100% remote jobs and will continue to follow that after the crisis.

https://remoteleaf.com/whoishiring

Note: Select "ruby" in the category filter to see rails jobs. 

✅ 100% remote full-time jobs    
✅ Spent more than 30 hours to curate this information
## [3][Getting db:migrate error eitherway](https://www.reddit.com/r/rails/comments/ha1aop/getting_dbmigrate_error_eitherway/)
- url: https://www.reddit.com/r/rails/comments/ha1aop/getting_dbmigrate_error_eitherway/
---
Hello,   

I am new to rails. I have been working on a basic sign-up/login module. On localhost, I am getting the following error:  

&gt; Migrations are pending. To resolve this issue, run:
&gt; rails db:migrate RAILS_ENV=development

Then when I follow the that instruction, I am getting the following error:


&gt; An error has occurred, all later migrations are canceled: Mysql2::error: Table 'users' already exists


How can this be fixed ?
## [4][MySQL vs PostgreSQL for production?](https://www.reddit.com/r/rails/comments/h9tsrj/mysql_vs_postgresql_for_production/)
- url: https://www.reddit.com/r/rails/comments/h9tsrj/mysql_vs_postgresql_for_production/
---
Which of these two would you recommend for a production database for a Rails app? I've worked with both before, but am much more familiar with MySQL. Does one have any major advantages over the other when it comes to Rails? 

Also, is it worth exploring NoSQL options for Rails, like MongoDB or Redis?
## [5][Generate Dynamic URLs/Paths in rails based off table parameter](https://www.reddit.com/r/rails/comments/h9mozp/generate_dynamic_urlspaths_in_rails_based_off/)
- url: https://www.reddit.com/r/rails/comments/h9mozp/generate_dynamic_urlspaths_in_rails_based_off/
---
I'm looking to generate a dynamic url within rails based off a table parameter. As an example, let's say I have a bunch of partners and want to have a dedicated path /some-partner-name based off the saved name value in an Active Record value. What is the best approach to creating these paths/URLs?
## [6][Devise with Multiple User Types](https://www.reddit.com/r/rails/comments/h9kslp/devise_with_multiple_user_types/)
- url: https://www.reddit.com/r/rails/comments/h9kslp/devise_with_multiple_user_types/
---
I've got an application with 3 different kinds of users: administrators, representatives and cardholders.

I'm using Pundit to authorize different actions based on the user type but I'm having trouble configuring devise (and devise invitable) to play nice with the different user types. 

Here are my current Devise Routes:

```
  #
  # Devise Routes
  #
  devise_for :users, controllers: { invitations: 'invitations/administrators' }
  devise_scope :user do
    get 'sign_in', to: 'devise/sessions#new'
    get 'sign_up', to: 'invitations/administrators#new'
    get 'success', to: 'invitations/administrators#success', as: 'invitation_success'
    get 'representative_inquiry', to: 'invitations/administrators#representative_inquiry'
    post 'representative_inquiry_submission', to: 'administrators/users#representative_inquiry_submission'
    post 'sign_up', to: 'invitations/administrators#create'
    post 'invitations/representatives', to: 'invitations/representatives#create'
    post 'invitations/cardholders', to: 'invitations/cardholders#create'
  end
```

The problem is that when a representative accepts an invitation, they will call `accept_resource` in the AdministratorsController. That's not ideal because I need different actions to happen when different user types accept an invitations.

I've also tried a setup like this:

```
  devise_for :users, controllers: { invitations: 'invitations/administrators' }
  devise_for :representatives, class_name: 'User', controllers: { invitations: 'invitations/representatives' }
```

This works a little better but the `warden` session is still iffy (note: `devise_for :users` _must_ be the first line and I can't use `devise_for :administrators` because then the warden session is "warden.user.administrator.key" instead of "warden.user.user.key" and devise can't authenticate the user).

Any ideas on the best way to manage this? I've never used devise with multiple user types and it's a bit weird.

Thanks!
## [7][easyAutocomplete in create form](https://www.reddit.com/r/rails/comments/h9gu3z/easyautocomplete_in_create_form/)
- url: https://www.reddit.com/r/rails/comments/h9gu3z/easyautocomplete_in_create_form/
---
I want to implement find easy autocomplete in a form, where I would search and select seller name, but when I submit the form I get seller name as a string. I need it to be seller ID. Is there a way for me to pass ID instead the name using  text\_field or I need to use other type of field?

&amp;#x200B;

https://preview.redd.it/l1y5l0dxx2551.png?width=1296&amp;format=png&amp;auto=webp&amp;s=9d596510ce62311dcb8620ba3100b94c9c45d1f2
## [8][Bulma Nav Dropdown - Quick flash on page load after converting to Webpacker](https://www.reddit.com/r/rails/comments/h9iwmq/bulma_nav_dropdown_quick_flash_on_page_load_after/)
- url: https://www.reddit.com/r/rails/comments/h9iwmq/bulma_nav_dropdown_quick_flash_on_page_load_after/
---
**EDIT**: It seems I have corrected the issue by removing some no longer necessary lines in my assets.rb initializer.  Nevermind! Mods feel free to delete or leave this post if you think it may help others.

---

Hello all, I just converted a big project from using the Asset pipeline and self-hosted Bulma framework to using Webpacker and the Bulma npm package, along with Stimulus.js and Stimulus Reflex.  After making the conversion, I'm noticing that my navbar dropdown (class `navbar-dropdown`) is flashing briefly when loading a page or moving amongst my app routes.  

I don't believe this to be a Bulma issue, but rather something with my webpacker or javascript pack configuration somewhere.  I am using the dropdown with the native CSS `is-hoverable` class, so no JS is written for this at all.

If anyone could help me out on this I would appreciate it.  I'm new to Webpack and probably missing something about how it works.  I can provide a video of the issue if needed. Thanks!

Here is my nav code and partial:

**_navbar.html.erb**

```
&lt;nav class="navbar has-shadow is-spaced" role="navigation" aria-label="main-navigation" data-controller="nav"&gt;
  &lt;div class="navbar-brand"&gt;
    &lt;%= link_to root_path, class: "navbar-item" do %&gt;
      &lt;%= image_tag("logo.png") %&gt;
    &lt;% end %&gt;
    &lt;a
      role="button"
      class="navbar-burger burger"
      aria-label="menu"
      aria-expanded="false"
      data-target="nav.burger"
      data-action="click-&gt;nav#toggleNav"
    &gt;
      &lt;span aria-hidden="true"&gt;&lt;/span&gt;
      &lt;span aria-hidden="true"&gt;&lt;/span&gt;
      &lt;span aria-hidden="true"&gt;&lt;/span&gt;
    &lt;/a&gt;
  &lt;/div&gt;

  &lt;div class="navbar-menu" id="navMenu" data-target="nav.menu"&gt;

    &lt;%= render "shared/userMenu" if user_signed_in? %&gt;

    &lt;%= render "shared/environment" if !Rails.env.production? %&gt;
    
    &lt;div class="navbar-end"&gt;
      &lt;div class="navbar-item"&gt;
        &lt;div class="buttons"&gt;
          &lt;%= render "shared/guestlinks" if !user_signed_in? %&gt;
        &lt;/div&gt;
      &lt;/div&gt;
    &lt;/div&gt;
  &lt;/div&gt;
&lt;/nav&gt;
```

**_userMenu.html.erb**

```
&lt;div class="navbar-item has-dropdown is-hoverable"&gt;
  &lt;a class="navbar-link has-text-grey-dark"&gt;
    &lt;span&gt;&lt;%= current_user.full_name %&gt;&lt;/span&gt;
  &lt;/a&gt;

  &lt;div class="navbar-dropdown ob-main-menu"&gt;
    &lt;div class="navbar-item"&gt;
      &lt;%= render "shared/userDetails" %&gt;
    &lt;/div&gt;

    &lt;hr class="navbar-divider"/&gt;

    &lt;%= link_to edit_user_registration_path, class: "navbar-item" do %&gt;
      &lt;span class="icon has-text-link is-medium ob-icon-rpad"&gt;
        &lt;i class="fas fa-cog"&gt;&lt;/i&gt;
      &lt;/span&gt;
      Settings
    &lt;% end %&gt;

    &lt;%= link_to accounts_inactive_path, class: "navbar-item" do %&gt;
      &lt;span class="icon has-text-link is-medium ob-icon-rpad"&gt;
        &lt;i class="fas fa-archive"&gt;&lt;/i&gt;
      &lt;/span&gt;
      Inactive Accounts
    &lt;% end %&gt;

    &lt;%= link_to destroy_user_session_path, class: "navbar-item", method: :delete do%&gt;
      &lt;span class="icon has-text-link is-medium ob-icon-rpad"&gt;
        &lt;i class="fas fa-sign-out-alt"&gt;&lt;/i&gt;
      &lt;/span&gt;
      Logout
    &lt;% end %&gt;
  &lt;/div&gt;
&lt;/div&gt;
```

**packs/application.js**

```
import "controllers"
import "@fortawesome/fontawesome-free/js/all"
import "../application/stylesheets/application"
```

**controllers/index.js**
```
import { Application } from "stimulus"
import { definitionsFromContext } from "stimulus/webpack-helpers"
import StimulusReflex from 'stimulus_reflex'
import consumer from '../channels/consumer'
import controller from './application_controller'

const application = Application.start()
const context = require.context("controllers", true, /_controller\.js$/)
application.load(definitionsFromContext(context))
StimulusReflex.initialize(application, { consumer, controller, debug: false })
```

**package.json**
```
{
  "name": "myapp",
  "private": true,
  "dependencies": {
    "@creativebulma/bulma-tooltip": "^1.2.0",
    "@fortawesome/fontawesome-free": "^5.13.0",
    "@rails/webpacker": "4.2.2",
    "bulma": "^0.8.2",
    "bulma-checkradio": "^1.1.1",
    "bulma-toast": "^2.0.1",
    "stimulus": "^1.1.1",
    "stimulus_reflex": "^3.2.1",
    "url-loader": "^4.1.0"
  },
  "devDependencies": {
    "webpack-dev-server": "^3.10.0"
  }
}
```

**Gemfile**

```
# frozen_string_literal: true

source 'https://rubygems.org'

git_source(:github) do |repo_name|
  repo_name = "#{repo_name}/#{repo_name}" unless repo_name.include?('/')
  "https://github.com/#{repo_name}.git"
end

ruby '2.6.6'

gem 'aws-sdk-s3', '~&gt; 1.60.0'
gem 'devise', '~&gt; 4.7.1'
gem 'draper', '~&gt; 3.1.0'
gem 'faker', '~&gt; 2.9.0'
gem 'figaro', '~&gt; 1.1.1'
gem 'jquery-rails'
gem 'jquery-turbolinks'
gem 'mini_magick', '~&gt; 4.9.5'
gem 'pagy', '~&gt; 3.7.1'
gem 'pg', '~&gt; 1.2.3'
gem 'puma', '~&gt; 4.3.3'
gem 'rails', '~&gt; 6.0.2.2'
gem 'rails-ujs', '~&gt; 0.1.0'
gem 'sassc', '~&gt; 2.2.1'
gem 'sassc-rails', '~&gt; 2.1.2'
gem 'stimulus_reflex', '~&gt; 3.2'
gem 'turbolinks', '~&gt; 5.2.1'
gem 'uglifier', '~&gt; 4.2.0'
gem 'webpacker', '~&gt; 4.2.2'

group :development, :test do
  gem 'capybara'
  gem 'database_cleaner', '~&gt; 1.7.0'
  gem 'factory_bot_rails', '~&gt; 5.1.1'
  gem 'pry-byebug'
  gem 'pry-rails'
  gem 'rails-controller-testing', '~&gt; 1.0.4'
  gem 'rspec-rails', '~&gt; 3.9.0'
  gem 'shoulda-matchers', '~&gt; 4.1.2'
  gem 'simplecov', require: false
end

group :development do
  gem 'better_errors'
  gem 'binding_of_caller'
  gem 'listen'
  gem 'rubocop', require: false
  gem 'rubocop-rails', require: false
  gem 'spring'
  gem 'spring-watcher-listen'
end
```
## [9][I have http://localhost:3001/items/product1, how do I add to the filepath to make it http://localhost:3001/items/product1/prices?](https://www.reddit.com/r/rails/comments/h9j3rp/i_have_httplocalhost3001itemsproduct1_how_do_i/)
- url: https://www.reddit.com/r/rails/comments/h9j3rp/i_have_httplocalhost3001itemsproduct1_how_do_i/
---
Hey all- I'm looking to add `/prices` to `http://localhost:3001/items/product1/`  so I have [`http://localhost:3001/items/product1/prices`](http://localhost:3001/items/product1/prices) . I have logic in the items controller for product1 already that I would like to use in /prices as well.  


How do I add this route? What is this type of routing called in [https://guides.rubyonrails.org/routing.html](https://guides.rubyonrails.org/routing.html) ? I'm getting lost and confused with the amount of routing documentation out there.
## [10][Which is better? Puma or Passenger?](https://www.reddit.com/r/rails/comments/h9em7o/which_is_better_puma_or_passenger/)
- url: https://www.reddit.com/r/rails/comments/h9em7o/which_is_better_puma_or_passenger/
---
Which server is best for handling high volume of light weight requests? I did load testing with both servers, Puma won the contest with very slight difference. The main reason I'm asking to reduce running servers for my production app. As of now 15 servers are running with passenger. Is this possible to reduce servers count to atleast 13 or 14 by switching to Puma from Passenger?

Note: It's an API app.
## [11][StimulusReflex, I was waiting for you](https://www.reddit.com/r/rails/comments/h92jon/stimulusreflex_i_was_waiting_for_you/)
- url: https://www.reddit.com/r/rails/comments/h92jon/stimulusreflex_i_was_waiting_for_you/
---
Hello to this new tool, without being really aware, I was actually waiting for it. I felt pain about this topic. There were no place were to store the front state in Ruby-on-Rails. Not properly. How to know which accordion is opened or not, how to know which tab is opened or not, etc. The only way I found until then was to use a centralized state (root object with Vue, or Vuex, or redux alone or with React). And store/restore it from localStorage. Without this centralized state, things get messy very quickly for complex scenarii. I wasn't completly convinced by StimulusJS itself because the state is assumed by the DOM, therefore it is very difficult to reason about it. Now with Reflex, it seems you can rely on a front state that is 100% assumed by the backend, removing the need to reason about a frontend state, and all the heavy mandatory JS tooling. Guys, this is a huge promise. Thanks a lot !!
