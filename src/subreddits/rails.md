# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/hja8kx/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/hja8kx/personal_projects_show_off_your_own_project_andor/
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
## [2][How to render a form partial from a model in other model template if it fails to save into database ?](https://www.reddit.com/r/rails/comments/ho09wu/how_to_render_a_form_partial_from_a_model_in/)
- url: https://www.reddit.com/r/rails/comments/ho09wu/how_to_render_a_form_partial_from_a_model_in/
---
The title is kind confusing but I don't know how to express my issue...

I'm working on a simple blog app. I have an Article model and Comment model. The comments are only shown on the /views/articles/show from given article.

The page has a form to post a comment. It's working fine to add a comment but the issue is to how to reload this page again with my shared/\_errors partial if the user try to post a blanked commentary.

For instance, in my articles controller my create function just render 'new' again if the article can't be saved.

How it would be in my create function on comments controller since comments only have two partials to render  (\_form and the \_comment to display all the comments). I don't know how to simply keep in this show page for article and render again the form with the errors message. You can check the code [here](https://github.com/Gregory280/alpha-blog-5.1.4).

&amp;#x200B;

    def createarticle = Article.find(params[:article_id])
    
    u/comment = u/article.comments.create(params[:comment].permit(:name, :comment))
    
    u/comment.user = current_user
    
    if u/comment.saveflash[:notice] = "Comment posted successfully!"redirect_to article_path(@article)
    
    elseflash.now[:notice] = 'Cant be blank'
    
    render 'comments/_form'
    end
    end 

(Don't why but reddit is changing a little bit my code)

Currently the controller is sending me from the route 'http://localhost:3000/articles/26 ' to '[http://localhost:3000/articles/26/comments](http://localhost:3000/articles/26/comments)' witch its a view with just the comment form.

&amp;#x200B;
## [3][Job uniqueness for ActiveJob](https://www.reddit.com/r/rails/comments/hnmbq8/job_uniqueness_for_activejob/)
- url: https://www.reddit.com/r/rails/comments/hnmbq8/job_uniqueness_for_activejob/
---
The [activejob-uniqueness](https://github.com/veeqo/activejob-uniqueness) is an attempt to implement something similar to sidekiq-unique-jobs, but working on more high-level abstraction, like ActiveJob callbacks, what makes it compatible with any ActiveJob adapter (including Sidekiq). It uses [redlock-rb](https://github.com/leandromoreira/redlock-rb) (implementation of [Redlock algorithm](https://redis.io/topics/distlock)) and therefore depends on Redis.
## [4][Migrate WordPress site to Rails CMS](https://www.reddit.com/r/rails/comments/hnmder/migrate_wordpress_site_to_rails_cms/)
- url: https://www.reddit.com/r/rails/comments/hnmder/migrate_wordpress_site_to_rails_cms/
---
Hey everyone! 

I just finished a Full Stack Development bootcamp (based on the Rails framework) and I loved everything about Rails !

Before doing the bootcamp, I was working on my own website about dogs, built with WordPress. Now, I look at it and I think about all the stuff I could add with my Rails knowledge (instead of only being a pure content website, i could really add web app features, like a way to find the nearest vets, breeders, etc.)

I could try to learn some PHP and do it on WordPress but I honestly really don't want to do that, and Rails and PHP work differently.

I was then thinking about taking all the already existing content and migrate it in a Rails CMS. That could be doable but i dont really know which CMS suits my needs.

Do you guys have any experience with this? Thanks a lot in advance !
## [5][Ruby on rails and React stack](https://www.reddit.com/r/rails/comments/hnjauk/ruby_on_rails_and_react_stack/)
- url: https://www.reddit.com/r/rails/comments/hnjauk/ruby_on_rails_and_react_stack/
---
Just curious. While going for a Rails and React webapp which code architecture is recommended.

1. React stack in the same Rails codebase OR

2.  a separate Rails api code base and a separate React stack.  

When it comes to community support and efficiency, which one is recommended?
## [6][How to make friendly_id backfilling migration faster? You can skip all the callbacks.](https://www.reddit.com/r/rails/comments/hndvj5/how_to_make_friendly_id_backfilling_migration/)
- url: https://www.reddit.com/r/rails/comments/hndvj5/how_to_make_friendly_id_backfilling_migration/
---
I am currently working on integrating friendly_id gem into some of the models in Talenox. Basically, it makes our in app URLs look nicer with human and company names in front, instead of just incremental primary key IDs. Oh boy… `Employee.all.each(&amp;:save)` is fucking slow in production.

There are several things that can cause update and insert to slow down a lot for an ActiveRecord model:

* Validations - especially when it involves multiple models
* Callbacks - especially when they cause a chain of callbacks in other models
* `belongs_to :parent, touch: true` - technically a callback to bust russian doll caches, but adding a slug does not necessitate busting caches

Guess what, we can skip all those. How? By backfilling with an empty model class.

Assuming we have an Employee model with a relation employees, what you can do is: Create an ActiveRecord model class in that migration class with none of the callbacks EXCEPT friendly_id and slug_candidate method.

    class BackfillEmployeesWithFriendlyId &lt; ActiveRecord::Migration[5.0]
    
      # Using a blank class allows us to easily skip all callbacks that can make
      # mass migration slow.
      class FriendlyIdEmployee &lt; ActiveRecord::Base
        self.table_name = 'employees'
        extend FriendlyId
        friendly_id :slug_candidate, use: [:slugged, :finders]
    
        def slug_candidate
          if first_name || last_name
            "#{first_name} #{last_name}"[0, 20]
          else
            "employee"
          end + " #{SecureRandom.hex[0, 8]}"
        end
      end
    
      def up
        print "Updating friendly_id slug for employees"
        FriendlyIdEmployee.where(slug: nil).each do |row|
          row.save; print('.')
        end
        puts ''
      end
    end
    
However, I couldn’t get the friendly_id history plug in to work properly yet. friendly_id history is implemented using ActiveRecord polymorphic. When the backfilling migration above is run, it will end up creating FriendlyId::Slug records with sluggable type of `BackfillEmployeesWithFriendlyId::FriendlyIdEmployee` instead of just `Employee`. That also means you can’t do subclassing of ActiveRecord models with friendly_id and expect history to work. Luckily we don’t need it.

[Source](https://anonoz.github.io/tech/2020/07/08/faster-friendly_id-mass-migration.html)
## [7][Rails Frontend Development Resources](https://www.reddit.com/r/rails/comments/hnhez6/rails_frontend_development_resources/)
- url: https://www.reddit.com/r/rails/comments/hnhez6/rails_frontend_development_resources/
---
Hi everyone, 

I've been doing Ruby and RoR for a couple of years now, but have always worked on API only projects with little or no frontend at all. Im looking for resources to learn the concepts of structuring a frontend rails app, guides on how to properly use layouts, view helpers etc. If anyone has any tutorials, books or videos that would be helpful I would really appreciate if you could share them. Thanks
## [8][Best architecture for preferred sort for e-commerce products?](https://www.reddit.com/r/rails/comments/hn78ue/best_architecture_for_preferred_sort_for/)
- url: https://www.reddit.com/r/rails/comments/hn78ue/best_architecture_for_preferred_sort_for/
---
I have an ecommerce app on rails and was wondering what's the best architecture for ordering products on a page. I'm aware of how to use .order(:table_attribute), but was wondering if there is a cleaner/more dynamic way to bubble offerings to the top that you want to highlight. Any recommended practices for this?
## [9][Database administration tool like ActiveAdmin/Rails_admin](https://www.reddit.com/r/rails/comments/hnchui/database_administration_tool_like/)
- url: https://www.reddit.com/r/rails/comments/hnchui/database_administration_tool_like/
---
I am looking for ways to work on an external (i.e. not the rails-backend) database from within a rails app, preferably in a very user-friendly and powerful way like ActiveAdmin/Rails\_admin allow. 

&amp;#x200B;

Does anyone know of a ready-made solution? To my knowledge you can't just point the two former tools to a different DB.
## [10][If you were creating a new Rails API, what serializer would you use and why?](https://www.reddit.com/r/rails/comments/hmymab/if_you_were_creating_a_new_rails_api_what/)
- url: https://www.reddit.com/r/rails/comments/hmymab/if_you_were_creating_a_new_rails_api_what/
---
We have a relatively new Rails API project (full Rails, not API only) that is using ActiveModelSerializer. We're considering switching due to some frustrations with the project, it's documentation, the current state of it, and a desire to get away from the confusing DSL. If you were starting a new project is there a serializer you recommend? fast-json-api comes to mind, but I haven't fully explored all the alternatives.
## [11][Why does Stripe only have documentation for Sinatra?](https://www.reddit.com/r/rails/comments/hn72ht/why_does_stripe_only_have_documentation_for/)
- url: https://www.reddit.com/r/rails/comments/hn72ht/why_does_stripe_only_have_documentation_for/
---
The ruby documentation for Stripe is specific to Sinatra and not rails. Heroku for the ruby documentation also defaults to Sinatra, and has Rails in a separate section. With rails being the most popular ruby framework, is there some a reason these products don't include rails as the default? Is there some historical reason for this? Just curious.
