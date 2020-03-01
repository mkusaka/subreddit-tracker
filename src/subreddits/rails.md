# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/f9t9kq/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/f9t9kq/personal_projects_show_off_your_own_project_andor/
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
## [2][Creating a third item by dragging and dropping one item on top of another](https://www.reddit.com/r/rails/comments/fbhgei/creating_a_third_item_by_dragging_and_dropping/)
- url: https://www.reddit.com/r/rails/comments/fbhgei/creating_a_third_item_by_dragging_and_dropping/
---
I'm writing an app where I have files that I want to be able to group together by dragging and dropping one on top of another. This should create a folder that can be named. It should also be possible to nest folders like this, by dragging and dropping folders on top of each other.

I got the associations working in Rails and I got the Javascript part of the drag and drop working as well, but I'm having trouble figuring out how to integrate them. 

Every "drag and drop" Rails article I found online is about reordering items in a list and relies on gems for nested sets or hierarchical data, which are irrelevant to my requirements, so it doesn't even help me by showing how to get JS and Rails talking in the way I'm looking for.

How should I go about doing something like this? Any ideas? Thanks!
## [3][How should I go about adding business hours and appointment slots to my model?](https://www.reddit.com/r/rails/comments/fb6j1e/how_should_i_go_about_adding_business_hours_and/)
- url: https://www.reddit.com/r/rails/comments/fb6j1e/how_should_i_go_about_adding_business_hours_and/
---
Hi everyone,  


This is going to be quite a generalised question but I'm about to implement the most complex feature I've tried to create, and I just want to get an idea of the best way to go about this.  


Basically I want to have two time-based features on my Business model - OpeningHours and AppointmentSlots.

It's a little complicated because a Business will have multiple opening\_hours, and multiple appointment\_slots for a single day. So I think this will have to be split up into two separate models - perhaps something like:

    class Business
      has_many :business_days
      has_many :opening_hours, through week_schedule
      has_many :appointment_slots, through week_schedule
    end
    
    class BusinessDay
      belongs_to: business
      has_many: opening_hours
      has_many: appointment_slots
    end
    
    class OpeningHours
      belongs_to: business_day
    end
    
    class AppointmentSlots
      belongs_to: business_day
    end

And the schema would be something like this:

    create_table "business_days", force: :cascade do |t|
        t.bigint "business_id"
        t.integer "day" #this will be an enum representing day of the week
        t.boolean "is_open"
      end
    
    create_table "opening_hours", force: :cascade do |t|
        t.bigint "business_day_id"
        t.time "open"
        t.time "close"
      end

This seems logical so far, but I kind of feel like this is probably a common thing to implement. Seems like there should be a gem that will add this kind of functionality to my model.  
I've looked at Biz ([https://github.com/zendesk/biz](https://github.com/zendesk/biz)) and Business Time ([https://github.com/bokmann/business\_time](https://github.com/bokmann/business_time)) and both of them seem somewhat useful, but the documentation for both of them doesn't actually mention how these incorporate with an existing model. So I'm not sure if they are right for what I'm looking for.  


To add to this, I want AppointmentSlots to have very similar functionality to Opening\_Hours - the actual data will look very similar (numerous appointment\_slots for each day).  


So it seems like I should be doing this with integration in mind somehow so that my code remains DRY. But I'm very new to all this and not really sure exactly how to approach the problem.  


I have a habit of going down the wrong route and then creating problems for myself down the road so if anyone can advise me on whether I'm approaching this correctly and any advise you might have?  


Thanks a lot.
## [4][JSON-RPC infraestructure](https://www.reddit.com/r/rails/comments/fbedv2/jsonrpc_infraestructure/)
- url: https://www.reddit.com/r/rails/comments/fbedv2/jsonrpc_infraestructure/
---
Hello im a begginner coder and I have been learning with ruby on rails. I have learn the basics like CRUD and REST. I can create features on my app. THis time  Im trying to create an app but he requires me to use JSON-RPC. I dont understand how to use it or set it up. Anyone can help me please? if is usefull is in order to create a Ruby on Rails app in order to connect to NEO Blockchain. Thanks
## [5][Advice on massive data sets](https://www.reddit.com/r/rails/comments/fb4fj8/advice_on_massive_data_sets/)
- url: https://www.reddit.com/r/rails/comments/fb4fj8/advice_on_massive_data_sets/
---
Howdy folks.  


I'm working on my app, [scientificmealplanner.com](https://scientificmealplanner.com), and thinking of re-engineering my meal plan generation algorithm significantly and I'd love thoughts on this direction.  


Currently, it generates meal plans essentially mashing recipes together in a semi-random fashion until it finds ones that work together. This approach works ok for now, but I am working on adding more and more parameters (calorie ranges, prep time ranges, protein ranges, etc etc), which makes a successful generation take longer with this approach.   


So my new possible direction is to pre-calculate stats for EVERY combination of 3 recipes (a day of the meal plan). In which case I can simply select from the pre-calculated combinations that match the desired parameters. I think, theoretically a lot easier.  


But this introduces a massive dataset with it's own challenges. I currently have 150 recipes, but eventually there may be like 1000 in the system. Using the combinations formula, we're talking \~551K records at 150 recipes to \~166M records at 1000 recipes.   


In addition, each month I update all ingredients via food data central API, and then it will need to run a method on every one of these recipe combination records to recalculate stats. So essentially re-calculating \~551K - 166M record stats once per month.  


I have no experience with that level of large dataset. Normally dealing with like a few thousand.

\- Is that amount of records feasible? Any idea what sort of database instance I'd need on AWS for this amount of records? Currently just on a t2.micro RDS, presumably need a much larger one.  
\- Any advice on handling millions of small background jobs? I'm using sidekiq/redis on AWS. Currently on a cache.t3.small.   


Thanks for any thoughts
## [6][ActiveStorage::FileNotFoundError for existing files](https://www.reddit.com/r/rails/comments/fb02tp/activestoragefilenotfounderror_for_existing_files/)
- url: https://www.reddit.com/r/rails/comments/fb02tp/activestoragefilenotfounderror_for_existing_files/
---
On a standard, out of the box Rails 6 with ActiveStorage set to disk storage,  I get his error from time to time:

    ActionView::Template::Error (ActiveStorage::FileNotFoundError): 
    2:     TODO cache this like forever     
    3: --&gt; 
    4: &lt;% if (current_website.icon.attached? rescue nil) %&gt; 
    5: &lt;link rel="apple-touch-icon" href="&lt;%= current_website.icon.variant(resize: "57x57").service_url %&gt;" /&gt; 
    6: &lt;link rel="shortcut icon" href="&lt;%= current_website.icon.variant(resize: "16x16").service_url %&gt;" /&gt; 
    7: &lt;link rel="icon" href="&lt;%= current_website.icon.variant(resize: "16x16").service_url %&gt;" /&gt; 
    app/views/common/_page_head_icons.html.erb:5 
    app/views/common/_page_head.html.erb:61 
    app/views/layouts/admin.html.erb:3

Restarting the server solves it. It happens on nginx, apache+mod\_passenger, puma. Haven't tested others, at this point chances of it being web server related are way too small.

Can anyone point a middle finger to what could be wrong? TIA

I got 500 coins saying I can find someone on reddit who knows wtf is going on
## [7][html.slim and react?](https://www.reddit.com/r/rails/comments/fb4192/htmlslim_and_react/)
- url: https://www.reddit.com/r/rails/comments/fb4192/htmlslim_and_react/
---
Has anyone incorporated a react front end using html.slim? 

we have tried doing 

= react\_component 'MainApp'

&amp;#x200B;

but nothing and thats the only research I can find that connect react and slim
## [8][Why does my variable behave differently when I assign using different ActiveRecord search method?](https://www.reddit.com/r/rails/comments/faoadl/why_does_my_variable_behave_differently_when_i/)
- url: https://www.reddit.com/r/rails/comments/faoadl/why_does_my_variable_behave_differently_when_i/
---
So I've noticed some unexpected behaviour when dealing with ActiveRecord.  


First off I assign my instance using 4 different ActiveRecord methods:  


    a = MyModel.first
    
    =&gt; #&lt;MyModel:0x00007f9ed138a4b0
     id: 1,
     name: "foobar",
     rating: 1
    
    
    
    b = MyModel.where(name: "foobar")
    
    =&gt; [#&lt;MyModel:0x00007f9ed706db00
      id: 1,
      name: "foobar",
      rating: 1
    
    
    
    c = MyModel.find(1)
    
    =&gt; #&lt;MyModel:0x00007f9ed7142da0
     id: 1,
     name: "foobar",
     rating: 1
    
    
    
    d = MyModel.where(rating: 1)
    
    =&gt; [#&lt;MyModel:0x00007f9ed1bc5328
      id: 1,
      name: "foobar",
      rating: 1

Now I try returning data from them:  


    a) WILL RETURN ALL DATA
    
    b) WILL ONLY RETURN NAME
        i.e. b.name =&gt; foobar
        b.id =&gt; undefined method `id
    
    c) WILL RETURN ALL DATA
    
    d) WILL ONLY RETURN NAME
        i.e. b.name =&gt; foobar
        b.rating =&gt; undefined method 'rating'

The actual instance seems to have been returned each time and I can see all the data when I just return the variable (e.g. just hitting a, b, c, or d), but I can only seem to access "name" for all of them. Only when assigning using .find(id) and .first will allow me to fully access all data.  


If someone could please explain why this is, that would be very much appreciated??
## [9][Ruby Hash#transform_keys now accepts a hash that maps existing keys to new keys](https://www.reddit.com/r/rails/comments/fae2xz/ruby_hashtransform_keys_now_accepts_a_hash_that/)
- url: https://blog.saeloun.com/2020/02/26/ruby-hash-transform_keys-now-accepts-a-hash-that-maps-existing-keys-to-new-keys
---

## [10][How can I disable migrations in the second database in Rails 6?](https://www.reddit.com/r/rails/comments/fad72g/how_can_i_disable_migrations_in_the_second/)
- url: https://www.reddit.com/r/rails/comments/fad72g/how_can_i_disable_migrations_in_the_second/
---
I'm doing a new project on a database in Rails 6.0.2.1, we're going to use two databases, one of them is the database of the Rails project and the other will be connecting to the database of an ERP we have, I want to completely disable migrations on this second database database (so a `rails db:drop` or a `rails db:reset` doesn't affect our ERP's database) and only use a `schema.rb` so that Rails can know what tables are in this db.

What can I do?
## [11][When developing a big app with different parts, is it smart to split every part into their own app?](https://www.reddit.com/r/rails/comments/faed8e/when_developing_a_big_app_with_different_parts_is/)
- url: https://www.reddit.com/r/rails/comments/faed8e/when_developing_a_big_app_with_different_parts_is/
---
I'm currently developing quite a big application (atleast for my standards). The application basically consists of a user-facing part with a "User" user model and a dashboard part with a "Vendor" user model (there will most likely also be a third part with it's own dashboard-like application in the future). The User part has the most business logic and is way bigger then the Vendor part, however, both parts are operating on and sharing the same data. But the data they share is the only thing they have in common. Everything else in the application, like for example the frontend or the domains, is pretty clear cut into these two parts. 

That's why I'm now wondering if it would be a better idea to really split these two parts into their own apps. So instead of having one big rails app, I would have one for the User side and one for the Vendor side which both share the same database. 

There are three big positives that I see with this approach:

1. **Seperation of concern.** Since they are essentially two different applications that just work on the same underlying data, having them on two different Rails apps should make everything more secure/stable and more easily extentable in the future.

2. **Simpler auth logic.** I was previously thinking about going with devise since both apps have entirely different user models and devise is the only auth gem that can handle multiple user models. However, I'm not a huge fan of devise. Splitting everything into two apps would allow me to choose a simpler auth gem like Clearance which has a way simpler api.

3. **Easily switchable technology.** Going with two apps instead of one big rails app would also allow me two choose and experiment with different technologies more easily. For example, I'm pretty interested in choosing Hanami but wouldn't choose it for the entire stack if everything would be a single app. However, splitting the one really big app into one big and one small app would allow me to choose Hanami for the smaller one and keep Rails for the big one. And if I find Hanami to be the wrong fit somewhere down the road I would only need to convert a small part of the original really big app back.

Does this all make sense and is this a common approach to handling bigger apps with multiple different parts or am I setting myself up for a lot of issues in the future if I go this path?
