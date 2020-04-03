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
## [3][Pluralsight suggestions for Rails?](https://www.reddit.com/r/rails/comments/ftztw7/pluralsight_suggestions_for_rails/)
- url: https://www.reddit.com/r/rails/comments/ftztw7/pluralsight_suggestions_for_rails/
---
Pluralsight is making all of their videos free for the month so I want to get as much out of it as I can. So right now I'm trying to find as many useful courses as I can for full stack development with Rails. 

I'm currently reading *Secrets of the JavaScript Ninja* to get a better understanding of JS and I have a copy of *Agile Web Development with Ruby On Rails* to read next for a good intro to Rails. Are there any Pluralsight courses I should look into to help?

I'm thinking about looking for a more in depth Ruby, HTML or CSS. I have experience with all of these but I could always benefit from learning more. Are there an specific courses I should check out or any other areas I should look into more? There is a "The MVC Request Life Cycle" course I'm thinking about too. Any other suggestion?
## [4][Animated page transitions with Ruby on Rails and Turbolinks](https://www.reddit.com/r/rails/comments/ftr4x3/animated_page_transitions_with_ruby_on_rails_and/)
- url: https://www.reddit.com/r/rails/comments/ftr4x3/animated_page_transitions_with_ruby_on_rails_and/
---
Here is something I wanted to have a look at for a long time. How to provide animated page transitions, while maintaining a non Javascript first approach and keeping the simplicity of server side rendering.

My team had a look at the internals of Turbolinks and how to achieve something as close as a single page app experience without the adding complexity on the tech stack [https://www.cookieshq.co.uk/posts/page-animations-with-turbolinks](https://www.cookieshq.co.uk/posts/page-animations-with-turbolinks)
## [5][Cast Raw SQL results into ActiveModel](https://www.reddit.com/r/rails/comments/fttdq6/cast_raw_sql_results_into_activemodel/)
- url: https://www.reddit.com/r/rails/comments/fttdq6/cast_raw_sql_results_into_activemodel/
---
Hey,

I'm sure there's a proper way of doing that but either I'm not looking for the right terms or it's harder than I expected it to be.

**Status**

I have two models: `Comments` and `Plays` that are `ApplicationRecord`s. Now in my app I have a need for listing information that is coming from various models. In my case it's an activity feed with comments a user created, and plays a user recorded.

These "events" don't exist in the DB, I just aggregate them on demand at the moment so I don't have an UserEvent ApplicationRecord that's persisted in the DB yet. To do this without creating a DB table I use an ActiveModel like this:

    class UserEvent
      include ActiveModel::Model
      attr_accessor :event_type,
                :event_title,
                :object_id,
                :object_type,
                :event_created_at
    end

I use a raw SQL query to aggregate the user events from these tables but I'm not sure how the next step is done best. What I want is "cast" my DB results into the UserEvent so the DB columns returned by the query map onto the fields of my ActiveModel.

**What I tried**

1) This can't work because UserEvent doesn't have `find_by_sql` (which makes sense)

```
      events_for_user = UserEvent.find_by_sql([events_for_user_sql, {user_id: current_user.id}])
```

2) This way feels a bit awkward, especially if using binds so I'm not sure if it's a good approach: https://stackoverflow.com/a/48391562/989919

      binds = [ ActiveRecord::Relation::QueryAttribute.new(
          "user_id", current_user.id, ActiveRecord::Type::Integer.new
      )]

      r = ApplicationRecord.connection.exec_query(
          events_for_user_sql, 'sql', binds
      )

Is this a bad way of doing it? What's the most common way?

Thank you!
## [6][ERR_ADDRESS_UNREACHABLE when trying to connect to localhost server on phone](https://www.reddit.com/r/rails/comments/ftw383/err_address_unreachable_when_trying_to_connect_to/)
- url: https://www.reddit.com/r/rails/comments/ftw383/err_address_unreachable_when_trying_to_connect_to/
---
Hi all,  


Trying to test my site on my mobile but I am getting this error when I try to access it.  


Here's the steps I've taken:  
Mobile and MacBook connected to same WiFi network  
`rails s -b MY_IP_ADDRESS`  
`on mobile: http://MY_IP_ADDRESS:3000`

&amp;#x200B;

I've also tried using:  
`rails s -b 0.0.0.0`  
`on mobile:` [`http://0.0.0.0:3000`](http://0.0.0.0:3000)  


and  
`rails s -b 0.0.0.0 -p 3001`  
`on mobile: http://0.0.0.0:3001`  
`on mobile: http://MY_IP_ADDRESS:3001`  


Same error for all of the above. Any ideas?
## [7][Capistrano 2 vs 3](https://www.reddit.com/r/rails/comments/ftrchd/capistrano_2_vs_3/)
- url: https://www.reddit.com/r/rails/comments/ftrchd/capistrano_2_vs_3/
---
We have a Rails app that is deployed to a few hundred EC2 instances and we currently use Capistrano 2 with a fairly heavily modified deploy script to deploy it. We've noticed a lot of reliability issues with deploys lately including failures to connect to specific machines and generic connection failures that are harder to pin down. The machines that have connection failures frequently seem just fine when manually connecting to them after the deploy has failed, which makes the issues harder to diagnose.

Is Capistrano 3 is generally "better" at this sort of thing than Capistrano 2?

I'd also love to hear about people's experience with Capistrano 3 for similar deployment sizes and/or experience upgrading from Capistrano 2 to 3.
## [8][Active Admin Fetching All Associations on A Model](https://www.reddit.com/r/rails/comments/ftnro3/active_admin_fetching_all_associations_on_a_model/)
- url: https://www.reddit.com/r/rails/comments/ftnro3/active_admin_fetching_all_associations_on_a_model/
---
I have a model, `listing.rb` with about 10 different associations including active storage attachments. Currently, to `GET /admin/listings` takes a whopping 1550ms and it's to do with the fact that ActiveAdmin is fetching all 10+ associations.

My admin resource looks something like this:

    ActiveAdmin.register Listing do
      includes :shop, :variants, shop: :user
    
      actions :all, only: %i[index show]
    
      config.per_page = [15, 30, 50, 75, 100]
    
      scope :all
      scope :accepted
      scope :pending
      scope :rejected
      scope :withdrawn
    
      index do
        selectable_column
        id_column
    
        column :shop, link: true do |l|
          l.shop.user.name || l.shop.user.username
        end
        column :title
        column :description do |l|
          l.description.truncate(50)
        end
        column :price do |l|
          money = Money.new l.variants.first.price, l.currency
          money.format
        end
        column :state
      end
    
      batch_action :accept do |ids|
        Listing.where(id: ids).update(state: :accepted)
    
        redirect_to request.referer, notice: "Accepted #{ids.size} listings"
      end
    end

I'm only fetching three different relations, which means only 4 queries should be ran in total, but instead I'm getting this: [https://pastebin.com/XNcB3xAy](https://pastebin.com/XNcB3xAy) (way too large to put it here)

So my question is, **can you set what associations will be fetched with ActiveAdmin or does it just fetch all of them despite not needing half of them**
## [9][Properly Exporting JS Files per Page Only in Rails 6](https://www.reddit.com/r/rails/comments/ftimgw/properly_exporting_js_files_per_page_only_in/)
- url: https://www.reddit.com/r/rails/comments/ftimgw/properly_exporting_js_files_per_page_only_in/
---
I'm currently migrating an application from sprockets to webpacker under Rails 6. I have a situation where I name my js files after the page I attach it to for some crud modules. So I have:

* model_as/Show.js
* model_bs/Show.js

The style of each file is something like `var Show = (function() { })();` using the revealing pattern in vanilla js. The last line of each file is `window.Show = Show`.

In `ModelA`'s show page, I include it using `javascript_pack_tag 'model_as/Show.js'` and in`ModelB`'s show page, I include it using `javascript_pack_tag 'model_bs/Show.js'`. 

When I visit the show page for `A`, it works fine. However when I visit the show page for `B`, it seems to still have cache-ed `A`'s `Show.js` file. So I think it's an issue with `window.Show = Show`. 

My question is, given a setup like that, what's the best practice to export js files that we load only on a per page basis?
## [10][Launching Your First Freelance Application (A Tutorial Beginner Series)](https://www.reddit.com/r/rails/comments/ftgvwg/launching_your_first_freelance_application_a/)
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
## [11][I want to have anchor links that jump down to different sections of a partial that is rendered in a show view, how to?](https://www.reddit.com/r/rails/comments/ftncem/i_want_to_have_anchor_links_that_jump_down_to/)
- url: https://www.reddit.com/r/rails/comments/ftncem/i_want_to_have_anchor_links_that_jump_down_to/
---
So I have a `show` page view with ids on the same page that I jump to with links like so at the top of the page, to help jump down to them further below in the page:

    &lt;a href="#info"&gt;Info&lt;/a&gt;&lt;br&gt;
    &lt;a href="#contact"&gt;Contact&lt;/a&gt;&lt;br&gt;
    etc

And those ids on the same `show` page. But I also have a partial further down that I render such as 

    &lt;%= render 'partials/categories' %&gt;

That partial `'partials/categories'` it's own ids that I want to jump to with anchor links, from the `show` page itself using the above partial.  So at the top of the `show` page, how do I refer to ids in the partial?

I tried just straight up referencing ids from the partial but it didn't work

    &lt;a href="#info"&gt;Info&lt;/a&gt;&lt;br&gt; # "info" is an id on the show page itself
    &lt;a href="#contact"&gt;Contact&lt;/a&gt;&lt;br&gt; # "contact" is an id on the show page itself
    &lt;a href="#basketball"&gt;Basketball&lt;/a&gt;&lt;br&gt; # from the partial 'partials/categories', breaks page
    &lt;a href="#hockey"&gt;Hockey&lt;/a&gt;&lt;br&gt; # from the partial 'partials/categories', breaks page

by the way, if my terminology needs fixing, please have at it.  I haven't done much front-end training as of yet, need to!

------

RESOLVED: my bad! i was just missing the # for the one for the partial. woops. thanks, pretty simple
## [12][best way write inside pdf &amp; pptx](https://www.reddit.com/r/rails/comments/ftmag5/best_way_write_inside_pdf_pptx/)
- url: https://www.reddit.com/r/rails/comments/ftmag5/best_way_write_inside_pdf_pptx/
---
I need write in  pdf &amp; pptx  . use some data in side pdf &amp; pptx .  
Data like 

* Table with some style
* Add image and add above some text 
* add some Graph like pie graph
