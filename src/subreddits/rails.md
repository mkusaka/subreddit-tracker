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
## [2][How y'all doing performance wise with Rails 6 (compared to 5.2)?](https://www.reddit.com/r/rails/comments/ep8odb/how_yall_doing_performance_wise_with_rails_6/)
- url: https://www.reddit.com/r/rails/comments/ep8odb/how_yall_doing_performance_wise_with_rails_6/
---
We recently upgraded our monolith to Rails 6 and unfortunately there is a clear trend: https://i.imgur.com/6Ca3pT4.png  
*edit/*  
x-axis = time (this is like 4 days)  
y-axis = response time (mean, 90th, 95th)  
*/edit*

We did a few things at once (updated several gems to new major versions, finally using cache versioning) but no real change in *our* code. We switched back to non-versioned caching for a while but that doesn't change anything in the timings.

I don't want to exclude the possibility that it's not Rails but any of the gems we are using but before I do some more testing I thought I'd ask if any of you guys have some insights. Did your apps perform better or worse after Rails 5&gt;6 upgrade? Did you encounter anything particular while migrating?

Thanks in advance :)
## [3][prefer "dependent: destroy" in ruby/model or in postgres on delete cascade?](https://www.reddit.com/r/rails/comments/epf2d6/prefer_dependent_destroy_in_rubymodel_or_in/)
- url: https://www.reddit.com/r/rails/comments/epf2d6/prefer_dependent_destroy_in_rubymodel_or_in/
---
with rails 6  
when generating a model like this "be rails g model step step:text page:references"  
it adds foreign keys in postgres, and destroy can fail like this:  


[destroy can fail](https://preview.redd.it/f55cbmzlv2b41.png?width=1388&amp;format=png&amp;auto=webp&amp;s=9cc4d726bdb7f4e5aea573a7e6e43ed75f71ba40)

  
what is better?  
let ruby handle deletions or let postgres handle it via "on delete cascade" ?
## [4][Had anyone here packaged a rails app for on prem licensing and installation?](https://www.reddit.com/r/rails/comments/ephv2t/had_anyone_here_packaged_a_rails_app_for_on_prem/)
- url: https://www.reddit.com/r/rails/comments/ephv2t/had_anyone_here_packaged_a_rails_app_for_on_prem/
---
Any tips?
## [5][How to update attribute on notifications after load?](https://www.reddit.com/r/rails/comments/ep9f5y/how_to_update_attribute_on_notifications_after/)
- url: https://www.reddit.com/r/rails/comments/ep9f5y/how_to_update_attribute_on_notifications_after/
---
In my application a user can have a group with many alerts. When an alert is created it has an `t.boolean "read", default: false` attribute, which I want to set to `true` once the users has seen it.  


I have the following:

**groups\_controller.rb**

    class GroupsController &lt; ApplicationController
      def index
        @groups = current_user.groups.includes(:alerts, :new_alerts)
      end
    end

**index.html.erb**

    &lt;% @groups.each do |group| %&gt;
      &lt;% group.alerts.limit(5).each do |alert| %&gt;
         &lt;% if alert.unread? %&gt;
            &lt;span class="badge pending-status"&gt;New&lt;/span&gt;
         &lt;% end %&gt;
         ...
      &lt;% end %&gt;
    &lt;% end %&gt;

And in the **show.html.erb** on a group I show all the alerts.  


I want to know if it's possible to mark those 5 in **index.html.erb** as read once they have been loaded into the view? So if I have 5 alerts and 3 are unread, I want the 3 to be shown as **New**, but once the page is refresh or **show.html.erb** is loaded, they will be shown an read.

Any suggestions?

&amp;#x200B;

I have been fiddling with an idea of having the read attribute as an integer instead, with a default of 0 and first time they are loaded `alert.update(read: 1)` and then alerts with `read &gt; 2` would be the unread ones. What do you think about that?
## [6][Best Gems for Recurring Jobs](https://www.reddit.com/r/rails/comments/ep6hey/best_gems_for_recurring_jobs/)
- url: https://www.reddit.com/r/rails/comments/ep6hey/best_gems_for_recurring_jobs/
---
I am currently working on an app that requires a user to send out questions to other users at an interval either from when created and/or when answered. For example, one question is set to recur on the first of every month. Another is set to recur 30 days after the user has answered it. I started using the ice\_cube gem but am thinking there might be a better option out there. Does anyone know a good gem to handle this?
## [7][Why is querying against an includes so slow in this query?](https://www.reddit.com/r/rails/comments/ep5eoj/why_is_querying_against_an_includes_so_slow_in/)
- url: https://www.reddit.com/r/rails/comments/ep5eoj/why_is_querying_against_an_includes_so_slow_in/
---
So I have a an app where users can be in relationships with each other (purely platonic, of course).  There's a couple of included associations on each relationship, such as the inverse relationship, and I want to filter out relationships where the other user isn't active:

My first solution, to query against the included table, is incredibly slow
```
user.relationships.includes(
                :related_user,
                :other_table,
                inverse_relationship: [:other_table],
              )
                .where(users: { active: true })
            )
```

in the server logs I can see a massive query like this: 
```
SELECT "relationships"."id" AS t0_r0, "relationships"."relating_user_id" AS t0_r1, "relationships"."related_user_id" AS t0_r2, "relationships"."inverse_relationship_id" AS t0_r3, "relationships"."created_at" AS t0_r4, "relationships"."updated_at" AS t0_r5, "users"."id" AS t1_r0, "users"."first_name" AS t1_r1, "users"."last_name" AS t1_r2, "users"."email" AS t1_r3, "users"."password_digest" AS t1_r4, "users"."avatar" AS t1_r5, "users"."admin" AS t1_r6, "users"."manager" AS t1_r7, "users"."title" AS t1_r8, "users"."created_at" AS t1_r9, "users"."updated_at" AS t1_r10,
```

The second query is this: 
```
user.relationships.includes(
                :related_user,
                :other_table,
                inverse_relationship: [:other_table],
              )
                .reject { |r| !r.related_user.active? }
            )
```

This query is blazing fast

I'm curious what rails is doing under the hood here, as commonly you would keep this kind of filtering in the query,
## [8][[Help] Custom method works in development (local) but returns 301 in production (Heroku).](https://www.reddit.com/r/rails/comments/ep5iar/help_custom_method_works_in_development_local_but/)
- url: https://www.reddit.com/r/rails/comments/ep5iar/help_custom_method_works_in_development_local_but/
---
Hi everyone! As the title says I'm having trouble with a custom method when I deploy to production.

What I'm trying to do is to dynamically populate as much info as possible when a user is registering their address. Basically, the city field is disabled and when a user inputs their postal code the app looks it up and returns the potal code's city_id so it can then call cities/id and get the city's name to display it. 

Here's some code:

**controllers/postal_codes_controller.rb**

    # Get postal_code data by postal_code.number
    def get_postal_code_by_number
      @postal_code = PostalCode.find_by_number(params[:number])
      render json: @postal_code.to_json, status: :ok, location: @postal_code
    end

**config/routes.rb**

      get 'get_postal_code_by_number/:number', :to =&gt; 'postal_codes#get_postal_code_by_number'


This works 100% fine locally but when I try it in heroku it responds with a 301 code and redirects to home ('/').


**Example**

When I enter the url *app-name.herokuapp.com/get_postal_code_by_number/91000* this is what I get:

[Chrome: Inspect element &gt; Network](https://imgur.com/cawFzbg)	


Please note that the postal codes that will be used in the app are just numbers and are not continuous, so I can't do id=number.

Any ideas?

Cheers!

------


**[UPDATE: SOLVED!]**

Turns out it didn't have to do with the method or the ajax but rather it was a routing problem. I'm still not sure why it wasn't working on production but I was able to get it to work by adding the method as a collection resource in routes.rb:

**config/routes.rb**

    resources :postal_codes do
      collection do
        get 'number/:number', :action =&gt; 'get_by_number'
      end
    end

This makes it possible to have a url like:

&gt;herokuapp.com/postal_codes/number/91000

As you can see, I also changed the method name and route so that they make a bit more sense now that the url is structured in such a way. If the original names were kept, the url would have looked like:

&gt;herokuapp.com/postal_codes/get_postal_code_by_number/91000

which is not as good looking.

The only issue I'm now facing with this solution is that it's not generating a path, if I were to do something like:

**config/routes.rb**

    ... 
    get 'number/:number', :action =&gt; 'get_by_number', :as =&gt; by_number
    ...

The path would be *by_number_postal_codes_path* and not *postal_codes_by_number_path* as I would've liked. But nevertheless, it works in both development and production so it's all good now.

Hopefully this will help someone if they hit the same problem.

Cheers! :D
## [9][Difference between to_h and to_hash for ActiveRecord::Errors class in rails](https://www.reddit.com/r/rails/comments/ep6dtr/difference_between_to_h_and_to_hash_for/)
- url: https://www.reddit.com/r/rails/comments/ep6dtr/difference_between_to_h_and_to_hash_for/
---
For a class like this, ActiveRecord::Errors, to\_h and to\_hash gives two different outputs although they are alias of each other.

&gt;class Person  
&gt;  
&gt;\# Required dependency for ActiveModel::Errors  
&gt;  
&gt;extend ActiveModel::Naming  
&gt;  
&gt;  
&gt;  
&gt;def initialize  
&gt;  
&gt;u/errors = ActiveModel::Errors.new(self) # wanted to type "@errors"  
&gt;  
&gt;end  
&gt;  
&gt;  
&gt;  
&gt;attr\_accessor :name, :phone, :age  
&gt;  
&gt;attr\_reader   :errors  
&gt;  
&gt;  
&gt;  
&gt;def validate!  
&gt;  
&gt;errors.add(:name, :blank, message: "Name cannot be nil") if name.nil?  
&gt;  
&gt;errors.add(:name, :blank, message: "Name is mandatory") if name.nil?  
&gt;  
&gt;errors.add(:phone, :blank, message: "Phone cannot be nil") if phone.nil?  
&gt;  
&gt;end  
&gt;  
&gt;\# The following methods are needed to be minimally implemented  
&gt;  
&gt;  
&gt;  
&gt;def read\_attribute\_for\_validation(attr)  
&gt;  
&gt;send(attr)  
&gt;  
&gt;end  
&gt;  
&gt;  
&gt;  
&gt;def self.human\_attribute\_name(attr, options = {})  
&gt;  
&gt;attr  
&gt;  
&gt;end  
&gt;  
&gt;  
&gt;  
&gt;def self.lookup\_ancestors  
&gt;  
&gt;\[self\]  
&gt;  
&gt;end  
&gt;  
&gt;end  
&gt;  
&gt;  
&gt;  
&gt;person = [Person.new](https://Person.new)()  
&gt;  
&gt;person.validate!  
&gt;  
&gt;person.errors.message =&gt; {:name=&gt;\["can't be blank", "can't be blank"\], :phone=&gt;\["can't be blank"\]}  
&gt;  
&gt;person.to\_hash =&gt; {:name=&gt;\["can't be blank", "can't be blank"\], :phone=&gt;\["can't be blank"\]}  
&gt;  
&gt;person.to\_h =&gt; {:name=&gt;"can't be blank", :phone=&gt;"can't be blank"}

&amp;#x200B;

Why is this difference between to\_h and to\_hash for this particular class?  


to\_hash has documentation in here [https://apidock.com/rails/v5.2.3/ActiveModel/Errors/to\_hash](https://apidock.com/rails/v5.2.3/ActiveModel/Errors/to_hash)
## [10][Read that the next Ruby release would update Ruby to be 3xs faster. Will that also improve/affect Rails performance?](https://www.reddit.com/r/rails/comments/ep5lzs/read_that_the_next_ruby_release_would_update_ruby/)
- url: https://www.reddit.com/r/rails/comments/ep5lzs/read_that_the_next_ruby_release_would_update_ruby/
---
was speaking to someone who said things can get quite bloated as a Rails app gets larger and larger, in their view, so just wondering.
## [11][The best approach for deleting existing association while updating](https://www.reddit.com/r/rails/comments/ep12dg/the_best_approach_for_deleting_existing/)
- url: https://www.reddit.com/r/rails/comments/ep12dg/the_best_approach_for_deleting_existing/
---
I'm writing an update API where when the user sends the updated parameter the already existing records for that association must be deleted and the ones in the API must be updated.   
If I have a relationship like Book has many Authors and if a particular book has a few authors and I'm trying to update the book by setting the new authors. The already existing authors for this book must be deleted and the new one updated. 

The association is defined as below

    has_many :authors, class_name: 'Authors', dependent: :destroy, autosave: true, after_add: :compute_something,
after_remove: :compute_some_other_thing

In my update method, I have something like this 

    book.authors.delete_all
    book.update_attributes(params[:book])

I'm creating new authors even if they exist after removing them by using the accepts\_nested\_attributes\_for property in rails.   


My question is do you see any pitfalls or bottlenecks with this approach in terms of the after\_add and after\_remove method execution or something else too.   
The authors aren't huge always it's mostly around 5 on average and around 25 at max.
