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
## [2][Does anyone build build their rails API using the JSON API standard? What server side library do you use?](https://www.reddit.com/r/rails/comments/fbzmuj/does_anyone_build_build_their_rails_api_using_the/)
- url: https://www.reddit.com/r/rails/comments/fbzmuj/does_anyone_build_build_their_rails_api_using_the/
---
The [JSON API](https://jsonapi.org/) standard is a general standard for formatting API responses for your JSON API. It's one of many standards, but I kind of like the simplicity of it and it seems fairly popular. 

Their website mentions [several ruby/rails gems](https://jsonapi.org/implementations/#server-libraries-ruby) that implement the standard. All of them promise that you can just add them and it generates links, routes, responses, etc... for you. 

I looked through a few of them and found the amount of choices overwhelming, and a few of them weren't really maintained that well. 

Do any of you use the JSON API standard for your own APIs, and if so what would you recommend? 

Thanks!
## [3][Suggestions for incorporating a low level recommendation engine for my v1 app](https://www.reddit.com/r/rails/comments/fbyv7x/suggestions_for_incorporating_a_low_level/)
- url: https://www.reddit.com/r/rails/comments/fbyv7x/suggestions_for_incorporating_a_low_level/
---
Hey everyone. I basically have a simple app that involved rating things. In this case you upvote/downvote dishes from restaurants. Dishes are tagged categorically. For example, the dish named, *Casey's Hometown Burger* will presumably be tagged as a burger. There's also an option for users to upload photos of dishes and users collectively can view recently uploaded on an, *explore-like* page similar to Instagram's.

I'd like to tie in some sort of recommendation logic where if a user has a tendency to upvote multiple dishes with the same tag then they'd get shown more burger related photos on the explore page. I would do the same when they would search for restaurants. If a restaurant offers burgers and the burgers are highly rated there, this listing would show up further up in results for this user vs. another user that has a tendency to rate pasta related dishes.

I'm aware that this falls under the realm of machine learning, something I have absolutely no experience in unfortunately. I was wondering if there was something like a reliable gem or article that could at the very least help me get to proof of concept. Thanks!
## [4][rails new project --webpack=react NOT working](https://www.reddit.com/r/rails/comments/fbxitf/rails_new_project_webpackreact_not_working/)
- url: https://www.reddit.com/r/rails/comments/fbxitf/rails_new_project_webpackreact_not_working/
---
&amp;#x200B;

https://preview.redd.it/p16bibs9l3k41.png?width=1909&amp;format=png&amp;auto=webp&amp;s=9a11bc9f20ce70da5cee1e8444d28e7e8a1f50c0

&amp;#x200B;

I get this error.

I installed rails using  [http://railsinstaller.org/en](http://railsinstaller.org/en)

Windows ruby 2.3

I don't see javascript folder inside the app folder

what should I do?Thanks
## [5][Creating a third item by dragging and dropping one item on top of another](https://www.reddit.com/r/rails/comments/fbhgei/creating_a_third_item_by_dragging_and_dropping/)
- url: https://www.reddit.com/r/rails/comments/fbhgei/creating_a_third_item_by_dragging_and_dropping/
---
I'm writing an app where I have files that I want to be able to group together by dragging and dropping one on top of another. This should create a folder that can be named. It should also be possible to nest folders like this, by dragging and dropping folders on top of each other.

I got the associations working in Rails and I got the Javascript part of the drag and drop working as well, but I'm having trouble figuring out how to integrate them. 

Every "drag and drop" Rails article I found online is about reordering items in a list and relies on gems for nested sets or hierarchical data, which are irrelevant to my requirements, so it doesn't even help me by showing how to get JS and Rails talking in the way I'm looking for.

How should I go about doing something like this? Any ideas? Thanks!
## [6][How should I go about adding business hours and appointment slots to my model?](https://www.reddit.com/r/rails/comments/fb6j1e/how_should_i_go_about_adding_business_hours_and/)
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
## [7][JSON-RPC infraestructure](https://www.reddit.com/r/rails/comments/fbedv2/jsonrpc_infraestructure/)
- url: https://www.reddit.com/r/rails/comments/fbedv2/jsonrpc_infraestructure/
---
Hello im a begginner coder and I have been learning with ruby on rails. I have learn the basics like CRUD and REST. I can create features on my app. THis time  Im trying to create an app but he requires me to use JSON-RPC. I dont understand how to use it or set it up. Anyone can help me please? if is usefull is in order to create a Ruby on Rails app in order to connect to NEO Blockchain. Thanks
## [8][Advice on massive data sets](https://www.reddit.com/r/rails/comments/fb4fj8/advice_on_massive_data_sets/)
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
## [9][ActiveStorage::FileNotFoundError for existing files](https://www.reddit.com/r/rails/comments/fb02tp/activestoragefilenotfounderror_for_existing_files/)
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
## [10][html.slim and react?](https://www.reddit.com/r/rails/comments/fb4192/htmlslim_and_react/)
- url: https://www.reddit.com/r/rails/comments/fb4192/htmlslim_and_react/
---
Has anyone incorporated a react front end using html.slim? 

we have tried doing 

= react\_component 'MainApp'

&amp;#x200B;

but nothing and thats the only research I can find that connect react and slim
## [11][Why does my variable behave differently when I assign using different ActiveRecord search method?](https://www.reddit.com/r/rails/comments/faoadl/why_does_my_variable_behave_differently_when_i/)
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
