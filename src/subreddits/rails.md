# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/jfcv1r/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/jfcv1r/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/jnwqje/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/jnwqje/personal_projects_show_off_your_own_project_andor/
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
## [3][Quick question about how to create a non-RESTful route](https://www.reddit.com/r/rails/comments/joe6dg/quick_question_about_how_to_create_a_nonrestful/)
- url: https://www.reddit.com/r/rails/comments/joe6dg/quick_question_about_how_to_create_a_nonrestful/
---
I want to be able to use a URL like `/goals/:date`.  

My route currently looks like this:  

      resources :goals do
        get ':date', to: 'goals#filter_by_date'
      end


But that makes the URL `/goals/:goal_id/:date`. How do I remove `:goal_id`?
## [4][undefined method `auto_strip_attributes` in rails console?](https://www.reddit.com/r/rails/comments/jo67j8/undefined_method_auto_strip_attributes_in_rails/)
- url: https://www.reddit.com/r/rails/comments/jo67j8/undefined_method_auto_strip_attributes_in_rails/
---
I'm using this gem [https://rubygems.org/gems/auto\_strip\_attributes/versions/2.0.6](https://rubygems.org/gems/auto_strip_attributes/versions/2.0.6) for stripping white space before saving the field to the db. I'm using it in my model like:

    class Brand &lt; ApplicationRecord
      auto_strip_attributes :name
    end

No idea what I did, but now whenever I go into rails console, I'm getting:

    Traceback (most recent call last):
            3: from (irb):1
            2: from app/models/person.rb:1:in `&lt;top (required)&gt;'
            1: from app/models/person.rb:12:in `&lt;class:Brand&gt;'
    NoMethodError (undefined method `auto_strip_attributes' for Person (call 'Person.connection' to establish a connection):Class)

Version of gem I'm using:

    $ bundle list | grep auto_strip
      * auto_strip_attributes (2.6.0)

I tried:

\- To force bundle install: `$ bundle install --redownload`

\- stop any servers and only use rails console
## [5][Big open source API project built with Rails to learn good patterns from?](https://www.reddit.com/r/rails/comments/jnso6v/big_open_source_api_project_built_with_rails_to/)
- url: https://www.reddit.com/r/rails/comments/jnso6v/big_open_source_api_project_built_with_rails_to/
---
Hi guys, I recently picked rails for my next side project since it offers everything I need out of the box :)

I need to develop a rest api backend, so I created an app with the "only api" flag. I'm following some tutorials on GoRails but I'm looking for a good open source project to learn from.

Since I'm building an API, I'm looking for a project that includes a API.

Thanks :)
## [6][Best way to keep Rails site CSS in sync with a separate Wordpress site?](https://www.reddit.com/r/rails/comments/jnvcxz/best_way_to_keep_rails_site_css_in_sync_with_a/)
- url: https://www.reddit.com/r/rails/comments/jnvcxz/best_way_to_keep_rails_site_css_in_sync_with_a/
---
Is there an accepted best practice way of doing this, or do I just have to manually copy and paste any changes I make?
## [7][Where to get application templates?](https://www.reddit.com/r/rails/comments/jntpj9/where_to_get_application_templates/)
- url: https://www.reddit.com/r/rails/comments/jntpj9/where_to_get_application_templates/
---
I found a few cool templates, but most of them are made by people for educational purposes. I really suck at frontend design and everything related. I just need to know where's the treasure box!
## [8][How to get full stacktrace for minitest](https://www.reddit.com/r/rails/comments/jnvw74/how_to_get_full_stacktrace_for_minitest/)
- url: https://www.reddit.com/r/rails/comments/jnvw74/how_to_get_full_stacktrace_for_minitest/
---
This is how the message I get for a failure in minitest   


    Expected response to be a &lt;200&gt;, but was &lt;500&gt;.
            Expected: 200
              Actual: 500

How can I get the full backtrace to understand where it gave a 500? 

  
Something like mentioned in this [https://gist.github.com/jeremyvdw/1001999](https://gist.github.com/jeremyvdw/1001999)
## [9][adding CSS to a single view](https://www.reddit.com/r/rails/comments/jn9nfy/adding_css_to_a_single_view/)
- url: https://www.reddit.com/r/rails/comments/jn9nfy/adding_css_to_a_single_view/
---
I'm working on a rails 6 application ( i generated a controller with his view (index/show/edit) ) 

but on the show page i have to add a show some images so i'm going to use swiperJS  (like this exemple not mine btw) https://codepad.co/snippet/swiper-thumbs-gallery-with-two-way-control

so i need to add the JS and CSS to the show page only 

a did a:

    yarn add swiper 

and to the javascript folder i added a new file called GallerieSwiper.js which contant my JS 

and in my show.html.erb page i added:

    &lt;%= javascript_pack_tag 'GallerieSwiper' %&gt;


the issue is how to add the swiper.min.css to only the show view ? 

i can't added in the top of the page as it is a subview of the application.html.erb that contant the whole website structure
## [10][The best Rails developer that no one ever was](https://www.reddit.com/r/rails/comments/jmxfci/the_best_rails_developer_that_no_one_ever_was/)
- url: https://www.reddit.com/r/rails/comments/jmxfci/the_best_rails_developer_that_no_one_ever_was/
---
What resources helped you become a better rails developer? Or projects?
## [11][How to fix Rendered ActiveModel::Serializer::Null with Hash](https://www.reddit.com/r/rails/comments/jn0wim/how_to_fix_rendered_activemodelserializernull/)
- url: https://www.reddit.com/r/rails/comments/jn0wim/how_to_fix_rendered_activemodelserializernull/
---
Hi guys, I'm using Active Model Serializer and I'm having trouble updating a user. I'm not sure what to do to solve this problem? 

I get this error message.   

`app/controllers/registrations_controller.rb:19
[active_model_serializers] Rendered ActiveModel::Serializer::Null with Hash (0.09ms)` 

My registrations_controller is:

	`def update
		@user = User.find(params[:id])
		if @user.assign_attributes(registration_params)
			render json: @user
		else
			render json: { status: :bad_request }
		end	
	end`	


Here is my User Serializer:

`
class UserSerializer &lt; ActiveModel::Serializer
  include Rails.application.routes.url_helpers
  attributes  :id, :first_name, :last_name, :email, :photo_url, :login_status
  
  attributes :photo_url
  def photo_url
    variant = object.photo.variant(resize: "80x80")
    return rails_representation_url(variant, only_path: true)
  end  

  def login_status
    {
      status: :created,
      logged_in: true
    }
  end   
end`
## [12][Seo Optimizer gem](https://www.reddit.com/r/rails/comments/jmmmnu/seo_optimizer_gem/)
- url: https://www.reddit.com/r/rails/comments/jmmmnu/seo_optimizer_gem/
---
Hey,

I wrote a gem to manage sitemap, robots.txt, error pages, etc.

Things to optimize SEO, you know \^\^ ...

It uses the "sitemap\_generator" gem. 

I would like to have your opinions, comments and contributions on this. 

The first goal is to easily and efficiently manage SEO on a Rails application. 

You can find the source code here: [https://github.com/RonanLOUARN/seo\_optimizer](https://github.com/RonanLOUARN/seo_optimizer)

Have a great day!
