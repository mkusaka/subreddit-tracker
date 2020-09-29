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
## [3][Rails and Solidus](https://www.reddit.com/r/rails/comments/j1na6n/rails_and_solidus/)
- url: https://www.reddit.com/r/rails/comments/j1na6n/rails_and_solidus/
---
Hi guys,

coming more or less form python and starting recently with ruby. I am looking for a multi-vendor platform and from what I could see and read - Solidus might be the solution. Generally, I am looking for a multi-vendor platform. Does anyone from you have any experiences with it? Does it work good or bad? Any recommendations?  


Thanks
## [4][Which React/Rails setup should I use?](https://www.reddit.com/r/rails/comments/j1ig5o/which_reactrails_setup_should_i_use/)
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
## [5][MySQL vs PostgreSQL - rare chance at an easy upgrade](https://www.reddit.com/r/rails/comments/j1gawq/mysql_vs_postgresql_rare_chance_at_an_easy_upgrade/)
- url: https://www.reddit.com/r/rails/comments/j1gawq/mysql_vs_postgresql_rare_chance_at_an_easy_upgrade/
---
I just got a new job for a very small company that had a Rails 4 app that was so poorly made that they had a brand new Rails 5.2/6 app built to replace it.

The old outsource company also made a migration task to transform a large part of the data in the db because the original was a total cluster fuck. And this migration task works perfectly. Other aspects of the new app do not, and hence why they decided on hiring someone in house to take over.

This is rare opportunity to migrate from the old MySQL db to a Postgres db. I like Postgres, but I have to wonder... is it worth it?

I know Postgres and MySQL mostly have differences at the massive scale level, and right now this app is no where near that level and probably will never get there. It’s got 7-10 daily users who could be using the app at the same time, and db size is currently under 1gb with 9 years of data. Chances are it will not grow to 10x or even come close to needing any of the enterprise/large scale features that differentiate MySQL and Postgres.

So, is it worth it? Feels to me like it is, but I’m biased.
## [6][what causes the file_field to be invisible?](https://www.reddit.com/r/rails/comments/j1l1m2/what_causes_the_file_field_to_be_invisible/)
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
## [7][Best Way to Do Form Within a Form (not nested)?](https://www.reddit.com/r/rails/comments/j1nrsc/best_way_to_do_form_within_a_form_not_nested/)
- url: https://www.reddit.com/r/rails/comments/j1nrsc/best_way_to_do_form_within_a_form_not_nested/
---
Hi, I’m learning Rails by building an invoicing app, and I’m trying to figure out how to open a form for one model (Invoice, which belongs to Client) but within that form, open a new form to create an instance of another model (Client) that, once created, will become available to choose in the original form. 

Each invoice is tied to a Client, of course, and right now, I have to create a new Client in a separate New Client form, and then the New Invoice form uses a drop-down select to show all the clients that you can tie the invoice to before saving the invoice. 

But what I really want is to be able to create a new client from WITHIN the new Invoice form. So, within the New Invoice form, I have a “Add New Client” link and it opens a Tailwind modal, with a “client name” input field and “a href” styled as a Save button for the modal. That fake button links to a StimulusJS controller action, which simply saves the client name input field to a Javascript variable, and also I use “event.preventDefault()” and “event.stopImmediatePropagation” to prevent the modal’s Save ‘button’ from submitting the host Invoice form. That all works so far. 

**So my first question is:** Now that I have the new client’s name in my Javascript Stimulus controller, how can I pass that back to a Rails Client controller action, so I can create the new client? I can’t find any information about how to pass a value OUT of Stimulus back INTO a Rails controller. The most I can come up with is the vague idea that I could use Ajax to create a URL in my Stimulus controller, and insert the client name into the URL, and then use a custom Rails Create action to grab the name as a param. But I’m still really sketchy on how to, again, move that or any data from a Stimulus controller back to a Rails controller. 

**And my second question is:** if I get this working, once I close the New Client modal and return to its open Invoice form, will the Invoice form still not show the newly created Client in its drop-down select? If Rails loads those existing Client names into the drop-down when the new Invoice form is opened, then it seems like the drop-down would not display the new client from the modal. Any ideas for how to deal with this? I’ve thought of switching from a drop-down select to a autocomplete text field using Stimulus, which might be more dynamic, but haven’t tried that out yet…

Thanks for reading this, and any guidance/advice is much appreciated!
## [8][redirect_to the previous page but changing the language](https://www.reddit.com/r/rails/comments/j1fwh5/redirect_to_the_previous_page_but_changing_the/)
- url: https://www.reddit.com/r/rails/comments/j1fwh5/redirect_to_the_previous_page_but_changing_the/
---
 What am I trying to do? To redirect the user **in the previous page** with the website translated in new language selected

My previous script was (language\_controller):

    class LanguagesController &lt; ApplicationController
     def edit
         @language = Language.find_by(locale: I18n.locale)
         render layout: false
     end
    
     def update
         @language = Language.find(update_params[:id])
         if user_signed_in?
           setting = current_user.setting || current_user.build_setting
       setting.language = @language.id
           setting.save
         end
         redirect_to root_path(locale: @language.locale)
       end
    
        private
        def update_params
         params.require(:language).permit(:id)
       end
     end 

As you can see, there is **redirect\_to root\_path** that redirect the user in the home page. Even if he is changing the language in an article-page

So I edit it and I added in **edit** `session[:return_to] ||= request.referer`

and in my **update** action I replaced `redirect_to root_path(locale: @language.locale)`  
 with `redirect_to session.delete(:return_to)`

What is happening? If the user is logged, **it works very good** and it redirects the user in the previous page.

If the user is **not logged** he is redirected in previous page **BUT now** the language is the same. **Probably because I removed** `(locale: @language.locale)`**, right?**

I also try to edit redirect\_to in this way: `redirect_to session.delete(:return_to, locale: @language.locale)`  
 but it said that I can not add two elements there.

How to solve?
## [9][&lt;%- if and &lt;% end -%&gt;](https://www.reddit.com/r/rails/comments/j1g3xp/if_and_end/)
- url: https://www.reddit.com/r/rails/comments/j1g3xp/if_and_end/
---
What is it? `&lt;%- if and &lt;% end -%&gt;`

Why `-` ?!

Where I can see documentations about how does it work?
## [10][Migrating users to Amazon Cognito; keep everything the same and just change the current_user method?](https://www.reddit.com/r/rails/comments/j1h7xh/migrating_users_to_amazon_cognito_keep_everything/)
- url: https://www.reddit.com/r/rails/comments/j1h7xh/migrating_users_to_amazon_cognito_keep_everything/
---
I'm currently trying to decouple our User model in our original Rails app from the application authentication.  We're starting to branch out to a couple different applications and I'd like to use Amazon Cognito to manage user identities. 

We currently use JWTs for authentication, with the Rails server generating them and a React SPA sending them with each request so that current_user can be set in the application controller.  Am I missing anything or could this be as simple as importing the existing users into Cognito, and then just using the current_user method to find the `User` instance based on the JWT user_id as usual?  I imagine there'd be a little extra logic around keeping the Cognito users in sync, but I'm wondering if I'm missing anything major here.
## [11][I curated all the remote job openings from Hacker News who is hiring - September](https://www.reddit.com/r/rails/comments/j0s1hv/i_curated_all_the_remote_job_openings_from_hacker/)
- url: https://www.reddit.com/r/rails/comments/j0s1hv/i_curated_all_the_remote_job_openings_from_hacker/
---
Here I would like to share all the remote jobs that I've curated from Hacker News Who is hiring thread. All these are 100% remote jobs not just allowed to work from home during COVID-19. These are 100% remote jobs and will continue to follow that after the covid.

https://remoteleaf.com/whoishiring.   
Note: Please select "Ruby" in the category filter to view Ruby/Rails jobs

✅ 100% remote full-time jobs.    
✅ Each and every job is manually curated and verified. Spent more than 12 hours for this.
## [12][Add controller actions + discussing updating an attribute in different places? | Sundae Club | Ruby on Rails Livestream](https://www.reddit.com/r/rails/comments/j0uhcp/add_controller_actions_discussing_updating_an/)
- url: https://www.reddit.com/r/rails/comments/j0uhcp/add_controller_actions_discussing_updating_an/
---
Hi everyone,

More Ruby on Rails live-streaming this week, thanks to everyone that's watched over the last few weeks - hopefully this is useful or interesting to learners. I'm afraid I was running behind today, so didn't get a chance to post a link before the stream, so here's the link from earlier:

[https://www.youtube.com/watch?v=MWSngfjNKS0](https://www.youtube.com/watch?v=MWSngfjNKS0)

Here's the YouTube description of what's in the stream, as well as the chapter links (chapter's now appear in the YouTube video scrubber itself, which is neat).  


&gt;This week we went down a minor rabbit-hole with Tailwind CSS (totally my fault, I didn't research it enough and it's more different from the UI libraries I'm used to than I realised, I'll try and do some research during the week!).  We also built some initial edit and update controller actions for our Channel model and started looking at how we might allow admins to edit attributes on a Channel that a standard user can't edit. This comes from seeing this question in multiple places, so I was interested to raise it early.  
&gt;  
&gt;  
[02:35](https://www.youtube.com/watch?v=MWSngfjNKS0&amp;t=155s) \- Intro: What I hope to cover  
[07:34](https://www.youtube.com/watch?v=MWSngfjNKS0&amp;t=454s) \- Quick recap of where we're up to  
[08:49](https://www.youtube.com/watch?v=MWSngfjNKS0&amp;t=529s) \- A detour into Tailwind which totally didn't work  
[27:00](https://www.youtube.com/watch?v=MWSngfjNKS0&amp;t=1620s) \- Adding edit action to our Channels controller  
[53:12](https://www.youtube.com/watch?v=MWSngfjNKS0&amp;t=3192s) \- Adding a form for channels (more Tailwind stuff)  
[1:10:47](https://www.youtube.com/watch?v=MWSngfjNKS0&amp;t=4247s) \- Add an update action to our Channels controller  
[1:24:45](https://www.youtube.com/watch?v=MWSngfjNKS0&amp;t=5085s) \- How can we update the new attribute in a different place?
