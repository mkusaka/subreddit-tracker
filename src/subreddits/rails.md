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
## [2][Do I still need to bundle install if project gems are stored in application path?](https://www.reddit.com/r/rails/comments/gyykem/do_i_still_need_to_bundle_install_if_project_gems/)
- url: https://www.reddit.com/r/rails/comments/gyykem/do_i_still_need_to_bundle_install_if_project_gems/
---
If I git pull a project that has the \`bundle\_gems\` path inside of, say, \`vendor\`, is it still necessary to run \`bundle install\`.  

I'm actually working on a docker image – similar question: Do I need to run \` bundle install\` if I COPY in the project files inc the gems?
## [3][Is Michael Hartls Ruby on Rails Tutorial good for a complete beginner to Rails?](https://www.reddit.com/r/rails/comments/gyjr4q/is_michael_hartls_ruby_on_rails_tutorial_good_for/)
- url: https://www.reddit.com/r/rails/comments/gyjr4q/is_michael_hartls_ruby_on_rails_tutorial_good_for/
---
Hi, I have some knowledge and understanding of how Ruby works and I'm looking to build on that.

I was reccomended this book, looked it up and noticed it teaches rails 5.

Is this still a sufficient learning tool?

Also does anyone else have any thoughts on the book or any other learning tools?

Thank you in advance!
## [4][Rails 6.1's ActiveModel Errors Revamp](https://www.reddit.com/r/rails/comments/gy80or/rails_61s_activemodel_errors_revamp/)
- url: https://www.reddit.com/r/rails/comments/gy80or/rails_61s_activemodel_errors_revamp/
---
[https://code.lulalala.com/2020/0531-1013.html](https://code.lulalala.com/2020/0531-1013.html)

As Rails developers, we are all used to the \`book.errors(:title)\` interface. This has remained relatively stable up until now, but is soon going to change.

I'd like to share the new model errors changes, before Rails 6.1rc1 gets released. The article contains a list of deprecation and recommended replacements offered in the new implementation. I hope to get some feedback, and see if we need to improve the upgrade guide a bit, to make the migration process less painful.

And if you have any suggestion on the actual code changes it self, please also let me know. Thanks you!
## [5][redirect_to vs. render performance](https://www.reddit.com/r/rails/comments/gyhfg0/redirect_to_vs_render_performance/)
- url: https://www.reddit.com/r/rails/comments/gyhfg0/redirect_to_vs_render_performance/
---
I have a simple MVC app for form submission. The form is quite long, so we give the users to save drafts of their form.

There are two types of flows. 

**Flow 1:**
1. GET `/form/new`
2. POST `/form` (for save draft update)
3. in controller, `redirect_to edit_form_path(@form)`
4. GET `/form/1/edit`

This is semantic, allowed me to render a flash notice, etc. But when the redirect is executed, it actually hits the DB again to get the current user and fetch the form object from the DB. This doesn't seem performant because I already have those values in memory from my current request.

**Flow 2:**
1. GET `/form/new`
2. POST `/form` (for save draft update)
3. in controller, `render :edit`
4. return `:edit` template, URL stays as `/form`

This isn't super semantic, I have to stupidly write out `flash[:now]='blah'` on its own line, but I save the DB hit.

This seems like a common enough case that there should be a way to have my cake and eat it to, but I simply can't find it out there. Any suggestions? Is one "the right way" to go about this? Obviously my app will always be indie and I'm not super worried about performance concerns now, and with indexes it should be pretty speedy to go with Flow 1, but I figured I would ask now.

Thanks!
## [6][Speed up queries on text columns using :gin indexes with [Rails &amp; PostgreSQL]](https://www.reddit.com/r/rails/comments/gyeczs/speed_up_queries_on_text_columns_using_gin/)
- url: https://www.reddit.com/r/rails/comments/gyeczs/speed_up_queries_on_text_columns_using_gin/
---
[https://www.ramblingcode.dev/posts/adding\_index\_to\_text\_columns\_in\_rails/](https://www.ramblingcode.dev/posts/adding_index_to_text_columns_in_rails/)
## [7][assign_attributes inside Transaction block](https://www.reddit.com/r/rails/comments/gygcvr/assign_attributes_inside_transaction_block/)
- url: https://www.reddit.com/r/rails/comments/gygcvr/assign_attributes_inside_transaction_block/
---
I update a model using assign\_attributes inside a transaction block like so. 

    ActiveRecord::Base.transaction do
      model.assign_attributes(params)
    end

But sometimes what happens in the params gets persisted to the database. I feel it is happening in the \` commit\_transaction\` method inside \`transaction.rb \`file when I debug it.   


Could someone tell me if this is the expected behavior?
## [8][Final video on automatic book ranker.](https://www.reddit.com/r/rails/comments/gyebim/final_video_on_automatic_book_ranker/)
- url: https://www.reddit.com/r/rails/comments/gyebim/final_video_on_automatic_book_ranker/
---
The concept of the videos ( playlist here  [https://www.youtube.com/watch?v=uEwu7D5G-hU&amp;list=PLB4RncStK2LUbl9VWLQAHznLJrYz2YMB4](https://www.youtube.com/watch?v=uEwu7D5G-hU&amp;list=PLB4RncStK2LUbl9VWLQAHznLJrYz2YMB4) ) was to develop a book tracker which tracks how many books people are reading and ranks users on how many books you read.

Link: [https://youtu.be/Cqlp8DZ7UI4](https://youtu.be/Cqlp8DZ7UI4)
## [9][How to identify between new and edit actions in javascript with webpack](https://www.reddit.com/r/rails/comments/gy6o5o/how_to_identify_between_new_and_edit_actions_in/)
- url: https://www.reddit.com/r/rails/comments/gy6o5o/how_to_identify_between_new_and_edit_actions_in/
---
I am hiding some elements in the new action, but I don't want to do on the edit action. Is there a way to do this, without checking whether the URL has edit in it.  


I am using webpack, so there is no way to pass instance variables inside js
## [10][How to configure a private app in Shopify](https://www.reddit.com/r/rails/comments/gxxiu1/how_to_configure_a_private_app_in_shopify/)
- url: https://www.reddit.com/r/rails/comments/gxxiu1/how_to_configure_a_private_app_in_shopify/
---
Hi! I'm new to this community. I have a quick question about configuring a private app in Shopify. My app is in Ruby on Rails and I am trying to install it on to Shopify. I created a private app on Shopify through the partners dashboard. When I enter the domain (example.shopify.com) I created in Shopify in the login page, I get this error:

Oauth error invalid\_request: The Shopify API application does not support oauth

The url i'm redirected to is this: [https://example.myshopify.com/admin/oauth/authorize?client\_id=c54817f9011be80129703e1987a4a37e&amp;redirect\_uri=http%3A%2F%2Flocalhost%3A3000%2Fauth%2Fshopify%2Fcallback&amp;response\_type=code&amp;scope=read\_products&amp;state=2c0c64c657dd9a212e088f7018c6e8fe4ba87c39e4542ec1](https://example.myshopify.com/admin/oauth/authorize?client_id=c54817f9011be80129703e1987a4a37e&amp;redirect_uri=http%3A%2F%2Flocalhost%3A3000%2Fauth%2Fshopify%2Fcallback&amp;response_type=code&amp;scope=read_products&amp;state=2c0c64c657dd9a212e088f7018c6e8fe4ba87c39e4542ec1)

My code for configuration is this:

\`\`\`

ShopifyApp.configure do |config|  
  config.application\_name = "APP NAME"  
  config.api\_key = ENV\['SHOPIFY\_API\_KEY'\]  
  config.secret = ENV\['SHOPIFY\_API\_SECRET'\]  
  config.old\_secret = ""  
  config.scope = ENV\['SHOPIFY\_PERMISSIONS'\] # Consult this page for more scope options:  
 \# https://help.shopify.com/en/api/getting-started/authentication/oauth/scopes  
  config.embedded\_app = true  
  config.after\_authenticate\_job = false  
  config.api\_version = "2020-01"  
  config.shop\_session\_repository = 'Shop'  
  config.webhooks = \[  
{topic: 'orders/create', address: "#{ENV\['APP\_HOST'\]}/webhooks/orders/order\_create", format: 'json'},  
{topic: 'fulfillments/create', address: "#{ENV\['APP\_HOST'\]}/webhooks/orders/order\_fulfilled", format: 'json'}  
  \]  
end

\`\`\`

Am i using the wrong url format? I'm not sure what to set 'APP\_HOST' to be in env file either.

Would be grateful for any answers. Thanks!
## [11][specifying ruby version in gemfile](https://www.reddit.com/r/rails/comments/gy07ne/specifying_ruby_version_in_gemfile/)
- url: https://www.reddit.com/r/rails/comments/gy07ne/specifying_ruby_version_in_gemfile/
---
howdy, working through the Hartl book and I noticed in his gem file he says to copy, doesn't include Ruby whereas it is in my default new rails app gem file. Is the best practice the leave out Ruby in your gem file? how does it work with including Ruby in your gem file vs. not including? Will the rails app look for Ruby/Ruby version in a different file?
