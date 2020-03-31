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
## [2][What do you use to measure API endpoint calls?](https://www.reddit.com/r/rails/comments/fs7v2f/what_do_you_use_to_measure_api_endpoint_calls/)
- url: https://www.reddit.com/r/rails/comments/fs7v2f/what_do_you_use_to_measure_api_endpoint_calls/
---
I‘m searching for an inexpensive product, should  also be GDPR compliant.


Do you know an overview of popular products that solve such and similar tasks?
## [3][Using ES6 in Rails 4.2.x](https://www.reddit.com/r/rails/comments/fs8pio/using_es6_in_rails_42x/)
- url: https://www.reddit.com/r/rails/comments/fs8pio/using_es6_in_rails_42x/
---
Everyone here using ES6 in Rails 4.2.x ? 
I heard that I should use Sprockets 3 , but my sprockets is 2.11.0 and there is only a documentation from upgrading from 3 to 4 

Did you think it gonna break my project if I upgrade this from 2 to 3 ?
And I would say from 2 to 4 is impossible for me because my rails still 4.2 

Any guide for this ? 

When I want to deploy to production, and if I use ES6 it is always error , And I think they way here using Babel to transpile it 

If anyone here in 4.2 and using Es6 in their rails
I d love to get know your guide for this 
Thank you :)
## [4][How to Rspec this?](https://www.reddit.com/r/rails/comments/fsbo6p/how_to_rspec_this/)
- url: https://www.reddit.com/r/rails/comments/fsbo6p/how_to_rspec_this/
---
I've been assigned to write rspec tests for a bit dated legacy project as a part of training.

I am to write factories and spec for existing models. Some code in those models look weird and even silly, but I can not replace/refactor any of it and should create tests for model as the are.

The model code has a class method:

*def self*.filter\_type(type\_id)  
 *case* type\_id  
 *when* 1  
5  
 *when* 2  
6  
 *when* 3  
7  
 *when* 4  
3  
 *when* 5  
9  
 *when* 6  
10  
 *else*  
 8  
 *end*  
*end*

So I can just create examples for each 'when' one-by-one to get rid of it but it's the 'else' case that bugs me. How do I make a proper test to include any value othere then (1..6). 

There is no restriction to use only integers as 'type\_id' in the code but let's assume there is.

Any hints to this?
## [5][Best way to Solve O(n + 1) in a table display](https://www.reddit.com/r/rails/comments/frtx0b/best_way_to_solve_on_1_in_a_table_display/)
- url: https://www.reddit.com/r/rails/comments/frtx0b/best_way_to_solve_on_1_in_a_table_display/
---
Let's say I have model `A` that `has_many` model `B` that in turn `has_many` model `C`. 

If I want to create a page that has a table where given an instance of `A`, I should display rows of `B`'s where each column in each row is defined by `C`s of each `B`. Assuming each `C` in `B` always equal number of instances and can be sorted by `C`'s column called `priority`.

What's the most efficient way to go about it?

```
# Some controller:
@a = A.find(a_id)
@b_records = B.where(a_id: @a.id)
```

Haml O(n+1) style:

```
%table
  - @b_records.each do |r|
    %tr
      - r.cs.order('priority ASC').each do |c|
        %td= c.some_value
```
## [6][ActiveRecord::RecordInvalid: Validation failed: Users must exist on db:seed](https://www.reddit.com/r/rails/comments/frvagj/activerecordrecordinvalid_validation_failed_users/)
- url: https://www.reddit.com/r/rails/comments/frvagj/activerecordrecordinvalid_validation_failed_users/
---
I am trying to seed an sqlite3 db from a json file as a project. I have two models user and logins.

    `require 'json'  file = File.read('db/people.json')`
     data_hash = JSON.parse(file, symbolize_names: true) 
     records = JSON.parse(File.read('db/people.json'))
     records.each do |record| User.create!(record.except('logins').merge('password' =&gt; 'encrypted password'))
     end 
    records.each do |record| Login.create!(record['logins']) 
    end

When I run my rails db:seed it successfully seeds the users and then fails when creating the logins with this error **ActiveRecord::RecordInvalid: Validation failed: Users must exist** It may be something with my schema or my seed script im not sure which .

       ActiveRecord::Schema.define(version: 2020_03_30_164743) do
    
       create_table "logins", force: :cascade do |t|
        t.datetime "date"
        t.datetime "created_at", null: false
        t.datetime "updated_at", null: false
        t.integer "user_id"
        t.index ["user_id"], name: "index_logins_on_user_id"
       end
    
       create_table "users", force: :cascade do |t|
        t.string "first_name"
        t.string "last_name"
        t.string "city"
        t.string "state"
        t.string "email", default: "", null: false
        t.string "encrypted_password", default: "", null: false
        t.string "reset_password_token"
        t.datetime "reset_password_sent_at"
        t.datetime "remember_created_at"
        t.datetime "created_at", null: false
         t.datetime "updated_at", null: false
         t.index ["email"], name: "index_users_on_email", unique: true
         t.index ["reset_password_token"], name: "index_users_on_reset_password_token", unique: true
        end
    
       end

Here is a link to my source code [https://github.com/jslack2537/apiDemoApp](https://github.com/jslack2537/apiDemoApp)

&amp;#x200B;

I was able to make it work like this 

    records.each do |record|    
     u = User.new(record.except('logins').merge('password' =&gt; 'encrypted password'))    
     u.logins = record['logins'].map{|l| Login.new(l)}     
     u.save! 
    end

&amp;#x200B;
## [7][best implementation for a model with an owner and users?](https://www.reddit.com/r/rails/comments/frkvhe/best_implementation_for_a_model_with_an_owner_and/)
- url: https://www.reddit.com/r/rails/comments/frkvhe/best_implementation_for_a_model_with_an_owner_and/
---
I have a Team model in my rails app. The team should have an owner. The owner should be able to invite other users to his/her team.

I'm using a Team as an "account" for my app. This account can have many users on it, but ultimately belongs to the owner of the team.

Therefore a team should also have many users. The users can interact with the account/app a certain way. But, the owner has more authorization with what he/she can do with the account/app because of his/her owner status.

Also, the owner should be able to assign different roles to any of the users, with each role having it's own unique set of authorizations.

Lastly, users may have their own teams, separate from the ones the members of. Meaning, that users can have multiple teams, being members of some, and being owners of others (although, it's likely they'll only ever be the owner of one, for convenience reasons).

&amp;#x200B;

How can I implement this? My initial thought was this:

**3 models**: Team, User, UserTeam

**Many-to-many assocation (has\_many through):** using user\_team as the middle man.

Lastly, I thought to keep track of roles on the user\_team records.

But this doesn't *feel* right. Ideally, I'd like to be able to write

    team.users # =&gt; returns all users including owner
    team.owner # =&gt; returns user that is the owner
    
    user.teams # =&gt; returns all teams user belongs to (including owned teams)
    user.owned_teams # =&gt; returns only the teams that the user owns.
    
    # and most importantly
    team.owner == @owner # =&gt; true
    team.users.includes?(@owner) # =&gt; true

because owners are technically users with an owner role.

What do you all think? Was my first thought on the right track? is there a better way? or should I just completely give up programming and become a musician?

Thanks!
## [8][Building a rails app that can upload pictures to an S3 bucket via CarrierWave and fog-aws, but it's giving me a Errno::EPIPE at /posts Broken pipe error](https://www.reddit.com/r/rails/comments/frmi1y/building_a_rails_app_that_can_upload_pictures_to/)
- url: https://www.reddit.com/r/rails/comments/frmi1y/building_a_rails_app_that_can_upload_pictures_to/
---
Edit: Note to anyone with the same issue in the future: 

SOLVED - Remember to specify the region in fog.credentials 

&amp;#x200B;

Hi there.

First of all this is my first post, I am really new at both Ruby and Rails, so I am glad to be here.

The last few days I've tried following a youtube tutorial creating an "instagram clone" cause I for my second rails project wanted to learn about implementing a postgres db, devise user authentication and also file hosting on AWS.

For the last 2 days I have been stuck on this issue where I cannot push anything to S3, and I feel like either the error is extremely vague or I am just not equipped to understand it. I tried using better\_errors and binding\_of\_caller to get a better feel for what's going on, but I am as stuck as can be.

things I've done:

I generated a CarrierWave uploader via the cli, and changed it to use :fog for storageI created the CarrierWave initializer and filled in the required fog credentials and directory as such:

&amp;#x200B;

https://preview.redd.it/65f2byqgjrp41.png?width=1310&amp;format=png&amp;auto=webp&amp;s=9837ef56e4d1a0baec84e83d9a96405a6d10d183

I have the variables for my credentials and bucket in a separate application.yml file living in the config directory

I then have a Post model in which I mount the uploader mentioned, and I have a controller for it. It is the controller specifically that throws the error

&amp;#x200B;

https://preview.redd.it/emqer83ckrp41.png?width=1490&amp;format=png&amp;auto=webp&amp;s=94ad7b7745fbf906b4d2e0136406410decadfa6d

Please let me know what other information is required cause right now I am feeling desperate and that's never a fun experience :)

(also, if someone would be willing to help me out with screensharing or similar, I would appreciate and welcome it)

Best regards

&amp;#x200B;

ps: forgot to mention that it looks like it times out when trying to upload/connect to aws, and that I verified the credentials and bucket with the awscli tool
## [9][Ruby 2.7.0 Warnings](https://www.reddit.com/r/rails/comments/frjwgd/ruby_270_warnings/)
- url: https://www.reddit.com/r/rails/comments/frjwgd/ruby_270_warnings/
---
I started using Ruby 2.7.0 for my projects. However there seems to be a lot of warning messages with Rails 6 libraries like activesupport. Will this be an issue moving forward?
## [10][From %&lt;a href=... to &lt;%=link_to in text.gsub!](https://www.reddit.com/r/rails/comments/frmzyn/from_a_href_to_link_to_in_textgsub/)
- url: https://www.reddit.com/r/rails/comments/frmzyn/from_a_href_to_link_to_in_textgsub/
---
Hi, I'm customizing my  **markdown redcarpet (**`class MarkdownRenderer &lt; Redcarpet::Render::HTML`**)**.

I found this part

      def paragraph(text)
        text.gsub!(/@(\w+)/) do |match|
          %(&lt;a href="/user/#{match[1..-1]}"&gt;#{match}&lt;/a&gt;)
        end

Can I replace `%(&lt;a href="/user/#{match[1..-1]}"&gt;#{match}&lt;/a&gt;)` using `&lt;%= link_to #{match}, user_path(#{match}) etc. etc.` ?

**How to do?** What is the right syntax?
## [11][Dtos in rails](https://www.reddit.com/r/rails/comments/frn6cd/dtos_in_rails/)
- url: https://www.reddit.com/r/rails/comments/frn6cd/dtos_in_rails/
---
Hi, im from c#. How to create dto in ruby on rails?
