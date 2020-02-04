# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/ep2dw9/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/ep2dw9/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/evmx0w/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/evmx0w/personal_projects_show_off_your_own_project_andor/
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
## [3][Debugging story: Mysteriously truncated timestamps with ActiveRecord](https://www.reddit.com/r/rails/comments/eypj02/debugging_story_mysteriously_truncated_timestamps/)
- url: https://www.reddit.com/r/rails/comments/eypj02/debugging_story_mysteriously_truncated_timestamps/
---
Saving [\#Ruby](https://www.linkedin.com/feed/hashtag/?highlightedUpdateUrns=urn%3Ali%3Aactivity%3A6628670995878289408&amp;keywords=%23Ruby&amp;originTrackingId=FZCWQu%2BAcSd9sZ%2F%2BpT1cog%3D%3D) objects in [\#PostgreSQL](https://www.linkedin.com/feed/hashtag/?highlightedUpdateUrns=urn%3Ali%3Aactivity%3A6628670995878289408&amp;keywords=%23PostgreSQL&amp;originTrackingId=FZCWQu%2BAcSd9sZ%2F%2BpT1cog%3D%3D)  is not a rocket science, and there is little that could surprise me. Have you ever thought that? Well, we had too until we tried to debug a flaky test.     


Read the whole story: [https://www.toptal.com/ruby-on-rails/timestamp-truncation-rails-activerecord-tale](https://www.toptal.com/ruby-on-rails/timestamp-truncation-rails-activerecord-tale) to learn why saving timestamps may cause headache :)
## [4][ActiveAdmin + WYSIWIG editor + Image uploads...How would you handle this?](https://www.reddit.com/r/rails/comments/ey7sou/activeadmin_wysiwig_editor_image_uploadshow_would/)
- url: https://www.reddit.com/r/rails/comments/ey7sou/activeadmin_wysiwig_editor_image_uploadshow_would/
---
Hello guys,

Am using activeadmin on a blog app to be used internally in our company. I have a post model (with fields: featured image, title, body and extra\_images) configured with image uploading using active storage. The ideal situation is for every post to have a featured image and extra 3 other images in between the wall of text that results from the text-body. So far am able to upload and show the featured image but am experiencing some difficulties with how to use the image uploader on the text editor am using (Trumbowyg editor for activeadmin) - how do I configure it to upload using active storage?

Thanks.
## [5][Help pls](https://www.reddit.com/r/rails/comments/eyg56r/help_pls/)
- url: https://www.reddit.com/r/rails/comments/eyg56r/help_pls/
---
Hello!

I have 4 models: company, user(created by devise), building, issue

Issue looks this way:

class CreateIssues &lt; ActiveRecord::Migration\[6.0\]def change    create\_table :issues do |t|      t.string :title      t.text :desc      t.references :user, null: false, foreign\_key: true      t.references :building, null: false, foreign\_key: true      t.references :company, null: false, foreign\_key: true      t.date :date      t.date :deadline      t.text :report      t.string :status      t.timestampsendendend

So here there is 3 foreign keys. I also have a form where user input name of building and company.

In controller I have to get an ID of database record with inputed values. Here is my problem. How can we get an id of a record?

I tried to do this way but it doesn t work:

def create\\@issue = Issue.new(issue\_params)\\@issue.user\_id = current\_user\\@issue.building = Building.find\_by(name: params\[:building\_name\])\\@issue.company = Company.find\_by(name: params\[:company\_name\])if \\@issue.save      redirect\_to \\@issueelse      render 'new'endend
## [6][Object inheritance and structuring some objects.](https://www.reddit.com/r/rails/comments/eyfua2/object_inheritance_and_structuring_some_objects/)
- url: https://www.reddit.com/r/rails/comments/eyfua2/object_inheritance_and_structuring_some_objects/
---
Ok, so I was JUST ABOUT to try implementing something and I thought I'd do a search for it before I jump in and I'm thinking I'm glad I did so at this point.

#Where I'm At
I'm a little over a month into my rails journey and I'm having a very frustrating/amazing time so far! 

#What I Want To Do
I am building a site for me to display my own portfolio of... well almost anything:

* Graphic Design shit
* Retouching Samples
* Webdesign stuff
* Etc...

My vision is to to build a dynamic website in a sort of squarespace-y sort of way that lets me add pages in-browser rather than forcing myself to go back into development and add another page 'manually'. I know it's not hard to do, but I like the idea of being able to just add stuff as I feel called.

Anyway, I figured I could define a table of Site_Items with name:string, title:string, hidden:boolean, and order:integer properties, since all page types would need those basics. Then I could have other objects inherit that basic structure and go from there. I kind of sketched it out below...

Site_Item |  |  | | &lt; ApplicationRecord
---|---|----|----|----
\^------------ | Group| | | &lt; Site_Item
\^------------ | Page | | | &lt; Site_Item
 | \^------ | TEMPLATE_1 | | &lt; Page
 | \^------ | TEMPLATE_2 | | &lt; Page
 | \^------ | ...etc | | &lt; Page

I don't know what's obvious and what isn't, so essentially I want to be able to call @site_items = Site_Item.all, and since all Groups, Pages, and Templates are ultimately based on Site_Items, I'll just be able to iterate over this collection.

However, having searched the internet and stumbling across STIs, MTIs, Polymorphic Associations, and a [VOCAL DEBATE](https://www.reddit.com/r/rails/comments/8cr6u9/i_am_wondering_whether_i_should_use_sti_or_mti/) about how to use each, I wanted to ask here about my specific intended goal and how best to pursue it.

Moreover (and I'm about to betray my ignorance here...) I'm struggling to conceptualize which thing should be doing the inheriting. Should it be a controller? model? migration? I'm just not clear on that stuff I guess!

Any ideas or guidance would be incredibly helpful!
## [7][Receive payments for Side Project](https://www.reddit.com/r/rails/comments/ey5q90/receive_payments_for_side_project/)
- url: https://www.reddit.com/r/rails/comments/ey5q90/receive_payments_for_side_project/
---
What would you use in order to receive small payments ($3-$30)  for a side project ?

There are not different types of subscriptions and no complicated pricing schemas. Only pay by month ($3)  or pay annual price ($30).  

Stripe? Paypal? Braintree?  Any other suggestions?

Please keep in mind this is a low maintenance project that I may abandon after one month (If it fails to generate earnings) and I also want to find a way to accept payments quickly in case I build another and another and another project.
## [8][devise_ldap_authenticatable &amp; react_native](https://www.reddit.com/r/rails/comments/exqdrq/devise_ldap_authenticatable_react_native/)
- url: https://www.reddit.com/r/rails/comments/exqdrq/devise_ldap_authenticatable_react_native/
---
I'm adding a React Native App to my monolith rails app. I was wondering, what would be the best approach to dealing with user authorization?  I check the user based on their username/active directory group. How do I handle authenticating users through devise\_ldap\_authenticatable with React Native?
## [9][Nginx + Puma, shows nothing but the root page](https://www.reddit.com/r/rails/comments/expomw/nginx_puma_shows_nothing_but_the_root_page/)
- url: https://www.reddit.com/r/rails/comments/expomw/nginx_puma_shows_nothing_but_the_root_page/
---
First, I followed this :

[https://www.digitalocean.com/community/tutorials/how-to-deploy-a-rails-app-with-puma-and-nginx-on-ubuntu-14-04](https://www.digitalocean.com/community/tutorials/how-to-deploy-a-rails-app-with-puma-and-nginx-on-ubuntu-14-04)

and my nginx configuration and puma.rb matches this (I actually modified them a bit, but important parts are just like this)

and now, when I go to the domain, it:

1. doesn't load the assets (and I wrote a lot of css to fix RTL)
2. doesn't load any route except root.

in both cases, it gives me error 404.

as an update, it only happens on https. http is ok :|
## [10][How to handle CRUD on join table in backend/frontend?](https://www.reddit.com/r/rails/comments/exph1s/how_to_handle_crud_on_join_table_in/)
- url: https://www.reddit.com/r/rails/comments/exph1s/how_to_handle_crud_on_join_table_in/
---
I'm trying to understand how to handle a many\_to\_many relationship on CRUD using Rails/React. 

On CRUD of the has\_many model, do I handle CRUD of the belongs\_to join table in the backend controller? Or do I route the join table to the frontend and handle the CRUD ops there with \`fetch\`?

In the backend, would the controller look something like this on create? 

&amp;#x200B;

    def create
        @table = Table.create(table_params)
        if @table.save
            JoinTable.create(table_id: params.permit!(:table_id), other_table_id: 
        params.permit!(:other_table_id))
          render :json =&gt; @table, status: @ok
        else  
           render :json =&gt; { errors: @table.erros }
        end
    end

or would I handle this in the frontend and treat the join table like the other tables.
## [11][Deploying rails app using Nginx](https://www.reddit.com/r/rails/comments/exnchp/deploying_rails_app_using_nginx/)
- url: https://www.reddit.com/r/rails/comments/exnchp/deploying_rails_app_using_nginx/
---
I just put my project on a server, then I ran its production env.

Now, I configured nginx like this :

           root /home/prp-e/dakhlokharj/public;
    
            # Add index.php to the list if you are using PHP
            index index.html index.htm index.nginx-debian.html;
    
            server_name _;
    
            location / {
                    # First attempt to serve request as file, then
                    # as directory, then fall back to displaying a 404.
                    #try_files $uri $uri/ =404;
                    proxy_pass       https://localhost:3000;
                    proxy_set_header Host      $host;
                    proxy_set_header X-Real-IP $remote_addr;
            }

But when I go to the domain, it still shows me the old static index.

UPDATE: 

It's okay now, but doesn't load assets (I added another location for assets) and also doesn't load other routes, only root! It gave me 404.

P.S : Ruby installed using RVM.
## [12][How to find working gem?](https://www.reddit.com/r/rails/comments/exg1ek/how_to_find_working_gem/)
- url: https://www.reddit.com/r/rails/comments/exg1ek/how_to_find_working_gem/
---
Been in ruby for a month now!  
How do you guys find gem? i am looking for a gem for appointment scheduling, but everything i see is at least 2 years old. How to gauge if it still works. Do you perform unit test for each case?
