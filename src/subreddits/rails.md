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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/har6r7/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/har6r7/personal_projects_show_off_your_own_project_andor/
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
## [3][Best PWA options for Rails?](https://www.reddit.com/r/rails/comments/hal2l7/best_pwa_options_for_rails/)
- url: https://www.reddit.com/r/rails/comments/hal2l7/best_pwa_options_for_rails/
---
Looking for the current and best methods for building Progressive Web Apps with Rails. Any suggestions?

Thinking about maybe dabbling in StimulusReflex, but not seeing any solid discussion on PWA implementation / planning.

Goal is to build an app that feels as native as possible with a cohesive UI that eliminates building and maintaining multiple user interface styles for various devices.
## [4][Heroku Add-ons I am working on](https://www.reddit.com/r/rails/comments/haq5lv/heroku_addons_i_am_working_on/)
- url: https://www.reddit.com/r/rails/comments/haq5lv/heroku_addons_i_am_working_on/
---
Hi Devs!

If you are on Heroku and use Heroku Scheduler, I am working on 2 projects that might interest you! Feel free to reach out!

[https://schedulerctl.com](https://schedulerctl.com/)

[https://schedulermon.com](https://schedulermon.com/)
## [5][Integrating Keycloak](https://www.reddit.com/r/rails/comments/han42k/integrating_keycloak/)
- url: https://www.reddit.com/r/rails/comments/han42k/integrating_keycloak/
---
Has anyone integrated [Keycloak](https://keycloak.org) with a Rails app?   

I'm not very experienced with authentication in Rails, aside from the basics of setting up Devise, and it is unclear to me what I need to do to make the first connection.

There is a Rack middleware gem, [keycloak-api-rails](https://github.com/looorent/keycloak-api-rails), which handles JWT token management for Keycloak.  However, I must do something in my app that will start the authN process.

I believe it should go something like this:

1. User arrives at my Rails site, but user is not yet authenticated
2. User is automatically redirected to my Keycloak server, and Keycloak provides the login page
3. User logs in with a local Keycloak username/password, or they login via one of the configured social integrations (Github, for example)
4. Upon successful authN, Keycloak redirects user back to my Rails app
5. The keycloak gem looks at the header and sees the Bearer token, then does the appropriate inspection to determine that it is valid; if so, it allows Rails to process the request fully.

**I don't know what must be done to get from step 1 to step 2.**  There are many gems for doing many things regarding user authentication, but I don't know which is right for this use case.  I would like to be as precise as possible with my gem choices such that I do not include a lot of unnecessary dependencies or configurations.

Suggestions?  TIA.
## [6][Help with configuring PostgreSQL and locale in Rails app](https://www.reddit.com/r/rails/comments/hamond/help_with_configuring_postgresql_and_locale_in/)
- url: https://www.reddit.com/r/rails/comments/hamond/help_with_configuring_postgresql_and_locale_in/
---
Hi, I'm using a Mac running Catalina to develop Rails applications. Since I'm going to use PostgreSQL in production, I've switched to using it in development (from sqlite). However, despite the fact that my Rails app specifies 'unicode' in the database.yml file, I cannot use it with a PostgreSQL database with UTF8 encoding. It keeps giving me the following error message:

    ERROR:  encoding "UTF8" does not match locale "en_NZ.ISO8859-15" DETAIL:  The chosen LC_CTYPE setting requires encoding "LATIN9". 

What can I do to allow PostgreSQL to let me use UTF8 encoding, despite the fact that my Mac is set to NZ English?

This is definitely a problem for me as I've been tasked with creating a website that will have both English and Thai options.
## [7][Is the Ruby Cookbook out of date now?](https://www.reddit.com/r/rails/comments/hah1x4/is_the_ruby_cookbook_out_of_date_now/)
- url: https://www.reddit.com/r/rails/comments/hah1x4/is_the_ruby_cookbook_out_of_date_now/
---
Is the Ruby Cookbook out of date now or is it still mostly still relevant? If it's out of date, is there anything you would suggest? Either a ruby or a rails cookbook (or similar).
## [8][Ruby on Rails cheatsheet from Michael Hartl tutorials](https://www.reddit.com/r/rails/comments/ham7y6/ruby_on_rails_cheatsheet_from_michael_hartl/)
- url: https://www.reddit.com/r/rails/comments/ham7y6/ruby_on_rails_cheatsheet_from_michael_hartl/
---
Hi folks,   

Hope everyone is staying safe!

I started learning Ruby-on-Rails framework two weeks before &amp; made a cheat sheet with minimal explanations of different concepts, referred from **Michael Hartl**’s "[Learn Web Development with Rails - Fourth Edition](https://www.pdfdrive.com/ruby-on-rails-tutorial-learn-web-development-with-rails-4th-edition-e184254589.html)" tutorials.  

The textbook is really awesome &amp; helped me to build a simple blogging website similar to the basic version of Twitter and the source code is available in github: [https://github.com/ddlogesh/rails-tutorial](https://github.com/ddlogesh/rails-tutorial) 

I would like to share these two cheat sheets, which may be a good start for preparing ruby-on-rails framework or recollect the concepts learned from the above-mentioned textbook.  

Rails: [https://www.notion.so/Ruby-on-Rails-Cheat-Sheet-61106a73031d46ec81c285daaebf1409](https://www.notion.so/Ruby-on-Rails-Cheat-Sheet-61106a73031d46ec81c285daaebf1409)

Ruby: [https://www.notion.so/Ruby-Cheat-Sheet-7c8aabc9268b4e65b5691245dd19068c](https://www.notion.so/Ruby-Cheat-Sheet-7c8aabc9268b4e65b5691245dd19068c)

If you are a beginner to Ruby, please watch this 4-hours tutorial from **freeCodeCamp**: [https://www.youtube.com/watch?v=t\_ispmWmdjY](https://www.youtube.com/watch?v=t_ispmWmdjY). Though it's quite a long, it's worth watching it!

Please do share these cheatsheets with your Ruby-on-Rails enthusiasts, who may be interested to learn Ruby-on-Rails.
## [9][Help with internal table join](https://www.reddit.com/r/rails/comments/hahf0f/help_with_internal_table_join/)
- url: https://www.reddit.com/r/rails/comments/hahf0f/help_with_internal_table_join/
---
hello:

I have two tables:

`Player --&gt; has_many FantasyStarts`

`FantasyStarts --&gt;  belongs_to Player`

&amp;#x200B;

In english, I am trying to find all SeasonStats from the last year and week that were the first entry for that player in the whole table (ie who started last week and it was their first 'start' ever recorded in the table). I have constructed a bare sql query that works great, but I am trying to improve my rails query skills:

`FantasyStart.includes(:player).find_by_sql("with count_table as (select player_id, count(*) as num_starts from fantasy_starts where position != 'BN' group by player_id) select * from fantasy_starts join count_table on fantasy_starts.player_id = count_table.player_id where week = 13 and year = 2019 and count_table.num_starts = 1 and position != 'BN' ")`

I have also figure out a 'rails way' of creating the map of counts I would join to:

`FantasyStart.where.not(position: 'BN').group(:player_id).count`

which yields:

`=&gt; {184=&gt;2, 2562721=&gt;5, 2540215=&gt;12, 100004=&gt;57, 100001=&gt;26, 100006=&gt;62, 2505785=&gt;5, 2561029=&gt;1, 2541316=&gt;1, 2558954=&gt;1, 2552408=&gt;12, 2532820=&gt;60, 2507999=&gt;25, 2506194=&gt;16, 2505600=&gt;18, 2532977=&gt;6, 2507164=&gt;75, 2495441=&gt;4, 100022=&gt;23, 2543704=&gt;2, 2532807=&gt;1, 81288=&gt;1, 2550658=&gt;4, 2506386=&gt;15, 2560809=&gt;11, 2533349=&gt;4, 2560735=&gt;3, 2557976=&gt;19, 2556521=&gt;15,....`

any good ideas for a next step?
## [10][naming css classes and the placement of css styles](https://www.reddit.com/r/rails/comments/hahqvu/naming_css_classes_and_the_placement_of_css_styles/)
- url: https://www.reddit.com/r/rails/comments/hahqvu/naming_css_classes_and_the_placement_of_css_styles/
---

where do I put this bit bit of css in a small/madium size rails 6 app?

    .clients-favorite-button-style { font-size: 300%; }

idea 1:
It's first use was in views/posts/new.html.erb, so i guess it can go in

    app/assets/stylesheets/posts.scss

But, this button style is not specific to posts, so I feel posts.scss might not be the best place.

idea 2:
The second idea was to put it in app/assets/stylesheets/instacart.scss
(suppose the site we're building is instacart.com)
But this doesn't scale well, because after some months this file has grown
and contains to pretty much contain all the styles the site uses.

idea 3:
The application has a app/assets/javascripts folder, I can put it somewhere in
there but this style is not specific to any js/webpacker packs, and often the style is used outside of the webpacker world.

do you have any other ideas?


also
how do you namespace the styles, I've been using the lame approach like this:

    .instacart-color-1 { ... }
    .instacart-font-1 { ... }
    .instacart-bolder-1 { ... }
    etc
## [11][default_scope is evil?](https://www.reddit.com/r/rails/comments/ha6gqj/default_scope_is_evil/)
- url: https://www.reddit.com/r/rails/comments/ha6gqj/default_scope_is_evil/
---
I've read articles [like this](https://rails-bestpractices.com/posts/2013/06/15/default_scope-is-evil/) or [this one](https://stackoverflow.com/questions/25087336/why-is-using-the-rails-default-scope-often-recommend-against), but I think it depends on the use case.

For example at my job we want to add a soft delete in the user model, I think adding a `default_scope where(active: true)` would solve this issue pretty fast and efficiently, all views would show active users and maybe some views would need a little touch here and there. Hard deletions would be done at the terminal anyway and in only very specific cases/records. Also, the fact that it affects users initialization isn't a drawback in this case, but actually a pro.

What do you think?
## [12][Trying to deploy Rails and VueJS app to Heroku but...](https://www.reddit.com/r/rails/comments/halfdg/trying_to_deploy_rails_and_vuejs_app_to_heroku_but/)
- url: https://www.reddit.com/r/rails/comments/halfdg/trying_to_deploy_rails_and_vuejs_app_to_heroku_but/
---
I seem to be running into an issue with the Rails UJS module not being found according to the Heroku log.

&amp;#x200B;

The exact error is " ModuleNotFoundError: Module not found: Error: Can't resolve 'rails/ujs' "

I have it in my appication.js file require("@rails/ujs").start();

In my package.json file it 

"@rails/ujs": "\^6.0.0", 

Still saying it can't resolve it... not sure what else I  am missing.  Anyone run into this?  Any other code I can share let me know
