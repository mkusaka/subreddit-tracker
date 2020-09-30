# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/iui4p8/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/iui4p8/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/iya619/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/iya619/personal_projects_show_off_your_own_project_andor/
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
## [3][Are monoliths becoming cool again?](https://www.reddit.com/r/rails/comments/j2kuxq/are_monoliths_becoming_cool_again/)
- url: https://www.reddit.com/r/rails/comments/j2kuxq/are_monoliths_becoming_cool_again/
---
Just a gut feeling, but it seems like micro-service everything for small-medium companies is getting a bit out of flavour, and monoliths start to make more sense again. The hype pendulum keeps swinging.

This together with some JS fatigue and the new Deno project (which might deprecate Node) makes me think choosing boring old monolith tech like Rails may become cool again. Or am I reading too much into this?
## [4][ApplicationController.render could generate an invalid CSRF token– and you have one more way for a CSRF token to annoy you.](https://www.reddit.com/r/rails/comments/j2hl1o/applicationcontrollerrender_could_generate_an/)
- url: https://www.reddit.com/r/rails/comments/j2hl1o/applicationcontrollerrender_could_generate_an/
---
A recent annoying CSRF problem in our platform. Hope you could find the information useful.

Basically there was an ActionController::InvalidAuthenticityToken error on deleting a group.

This was occurring because we were using ApplcationController.render to render one of the buttons for a delete request. This was not working. ApplicationController has no information about the session. We should have used 'render'.

Be careful when using ApplicationController.render for rendering forms. 

[https://kmitov.com/posts/applicationcontroller-render-one-more-way-for-a-csrf-token-to-annoy-you/](https://kmitov.com/posts/applicationcontroller-render-one-more-way-for-a-csrf-token-to-annoy-you/)
## [5][Deploy Ruby on Rails on Google cloud](https://www.reddit.com/r/rails/comments/j26lkz/deploy_ruby_on_rails_on_google_cloud/)
- url: https://www.reddit.com/r/rails/comments/j26lkz/deploy_ruby_on_rails_on_google_cloud/
---
Hi, does anyone know what's the best way to deploy a ruby on rails application on google cloud also using the SQL instance on Google Cloud and a process to use sidekiq.

Do you guys have any pratical experience?
## [6][Trying to upload the file to rails server from React? I receive permitted: false in parameters of file](https://www.reddit.com/r/rails/comments/j25sdg/trying_to_upload_the_file_to_rails_server_from/)
- url: https://www.reddit.com/r/rails/comments/j25sdg/trying_to_upload_the_file_to_rails_server_from/
---
I'm using active\_storage to upload pictures. The file doesn't show on rails. Is the something with rails or Javascript? Where should I be looking

&amp;#x200B;

 

&lt;ActionController::Parameters {} permitted: false&gt;
## [7][Using Carrierwave, cannot access image_url in view](https://www.reddit.com/r/rails/comments/j2694q/using_carrierwave_cannot_access_image_url_in_view/)
- url: https://www.reddit.com/r/rails/comments/j2694q/using_carrierwave_cannot_access_image_url_in_view/
---
Images are saving to the database and are nested within a parent. Parent model is Nitrogen, Child model is Photos. For example, in the view when I try to access the image by  &lt;%=nitrogen.photos.ids%&gt;, the id is displayed.  Attached is a image of what I get in the rails console when I call Nitrogen.last.photos. How can I extract :picture within the ActiveRecord::Associations::CollectionProxy?
## [8][Trouble Understaning how the Rails framework works...](https://www.reddit.com/r/rails/comments/j1zum4/trouble_understaning_how_the_rails_framework_works/)
- url: https://www.reddit.com/r/rails/comments/j1zum4/trouble_understaning_how_the_rails_framework_works/
---
Hello there. I have been learing Ruby and Rails through the Odin Project. I am in the Rails section and I have a lot of trouble understaiding how many of the things work. To me they seem like magic and I don't really get it. I think the reason is because I learn better through video tutorials and explanations so I enrolled in the "The complete ruby on rails developer course" on udemy. I helped me a lot but there are a lot of things that the instructor does not go in depth or just will ignore and assume we already know. Are there any video tutorials, courses or lessons on Rails that you would recommend me?
## [9][Rails and Solidus](https://www.reddit.com/r/rails/comments/j1na6n/rails_and_solidus/)
- url: https://www.reddit.com/r/rails/comments/j1na6n/rails_and_solidus/
---
Hi guys,

coming more or less form python and starting recently with ruby. I am looking for a multi-vendor platform and from what I could see and read - Solidus might be the solution. Generally, I am looking for a multi-vendor platform. Does anyone from you have any experiences with it? Does it work good or bad? Any recommendations?  


Thanks
## [10][Which React/Rails setup should I use?](https://www.reddit.com/r/rails/comments/j1ig5o/which_reactrails_setup_should_i_use/)
- url: https://www.reddit.com/r/rails/comments/j1ig5o/which_reactrails_setup_should_i_use/
---
There is an enormous amount of advice on using React with Rails, so much that I'm struggling to decide on a setup to learn.

What I have managed to decide on is that I want to use Rails as an API, serving JSON to a separate React front end. What I haven't decided on is:

1. Whether to use Redux
2. Whether to store the React app in the same repo as the Rails app
3. Whether to use Rails' asset pipeline to serve the React app
4. Whether the app should be a SPA
5. Whether I should use React Hooks

From what I know about state and mutability it seems as though Redux is a smart option. That being said, I've read that it adds a lot of complexity and dev time to simple tasks, and considering I'm going to be working on something alone, which is just an idea, I wonder whether Redux might be overkill and have a negative effect on my motivation.

I know this is quite a wide question but really any pros/cons/successes/failures that anyone has would be really useful in helping me make a decision. And if you do have a recommended setup, and also know a good tutorial for that setup, then that would also be incredibly helpful.

Thanks
## [11][MySQL vs PostgreSQL - rare chance at an easy upgrade](https://www.reddit.com/r/rails/comments/j1gawq/mysql_vs_postgresql_rare_chance_at_an_easy_upgrade/)
- url: https://www.reddit.com/r/rails/comments/j1gawq/mysql_vs_postgresql_rare_chance_at_an_easy_upgrade/
---
I just got a new job for a very small company that had a Rails 4 app that was so poorly made that they had a brand new Rails 5.2/6 app built to replace it.

The old outsource company also made a migration task to transform a large part of the data in the db because the original was a total cluster fuck. And this migration task works perfectly. Other aspects of the new app do not, and hence why they decided on hiring someone in house to take over.

This is rare opportunity to migrate from the old MySQL db to a Postgres db. I like Postgres, but I have to wonder... is it worth it?

I know Postgres and MySQL mostly have differences at the massive scale level, and right now this app is no where near that level and probably will never get there. It’s got 7-10 daily users who could be using the app at the same time, and db size is currently under 1gb with 9 years of data. Chances are it will not grow to 10x or even come close to needing any of the enterprise/large scale features that differentiate MySQL and Postgres.

So, is it worth it? Feels to me like it is, but I’m biased.
## [12][what causes the file_field to be invisible?](https://www.reddit.com/r/rails/comments/j1l1m2/what_causes_the_file_field_to_be_invisible/)
- url: https://www.reddit.com/r/rails/comments/j1l1m2/what_causes_the_file_field_to_be_invisible/
---
I would like to make my app upload multiple files with Shrine, but [one doc](https://shrinerb.com/docs/multiple-files#3-create-the-view) suggests two `file_field`s whereas the [other](https://github.com/shrinerb/shrine/wiki/Adding-Direct-App-Uploads#5-form) suggests only one. After posting a question to their discourse forum, it was suggested that I hide the one named `files[]`. Whether I do this or not, the first `file_field` always fails to render. Why does this field not display?

    &lt;%= form_for @item, html: { enctype: "multipart/form-data" } do |f| %&gt;
     &lt;%= f.fields_for :photos do |i| %&gt;
      &lt;%= i.label :image %&gt;
      &lt;%= i.hidden_field :image, value: i.object.cached_photos_data, class: "upload-data" %&gt;
      &lt;%= i.file_field :image, class: "upload-file" %&gt; /// why is this not rendering?😢
     &lt;% end %&gt;
     &lt;%= file_field_tag "files[]", multiple: true %&gt; // what purpose does this one serve?
     
     &lt;%= f.text_field :title %&gt;
          
     &lt;%= f.submit "Submit" %&gt;    
    &lt;% end %&gt;

Models:

    class Item &lt; ApplicationRecord
     has_many :photos, as: :imageable, dependent: :destroy
    end
    
    class Photo &lt; ApplicationRecord
     include ImagesUploader::Attachment(:image)
     belongs_to :imageable, polymorphic: true
     validates_presence_of :image
    end

Controller:

    class ItemsController &lt; ApplicationController
    
     def new
      @item = current_user.items.new
     end
    
     def create
      @item = current_user.items.create(item_params)
      @item.save
     end
    
     private
     def item_params
      params.require(:item).permit(:title, photos_attributes: { image: [] })
     end
    end
