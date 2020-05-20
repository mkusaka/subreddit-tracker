# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gek2rk/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gek2rk/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gnbebg/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gnbebg/personal_projects_show_off_your_own_project_andor/
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
## [3][Why won't Heroku run the assets:precompile task (Rails6 + Webpacker)](https://www.reddit.com/r/rails/comments/gn4cpf/why_wont_heroku_run_the_assetsprecompile_task/)
- url: https://www.reddit.com/r/rails/comments/gn4cpf/why_wont_heroku_run_the_assetsprecompile_task/
---
What am I doing wrong? Pulling my hair out here. I was on webpacker 4.2.2 then upgraded to 5.1.1 thinking that was my problem...still no dice. Not much I haven't tried. This isn't a new app but one that I'm migrating to using webpacker. **Thanks for any help!**

In Heroku, the nodejs buildpack is before the ruby buildpack.

Logging: [https://gist.github.com/leesmith/2cc55b46e023f2cc29dbf239bf9401bd](https://gist.github.com/leesmith/2cc55b46e023f2cc29dbf239bf9401bd)

Code: [https://github.com/leesmith/decent\_authentication/tree/tailwind](https://github.com/leesmith/decent_authentication/tree/tailwind)

**Edit:** Suspicion confirmed. Running `heroku run rake -T --app my-app` returns nothing. WTF?

**Edit:** This had nothing to do with Heroku or webpack...I tried to use rubocop in my `Rakefile` and reverted it back to the standard `Rakefile` that you get with a new rails project. :)
## [4][Database structure for a serialized inventory system](https://www.reddit.com/r/rails/comments/gn770m/database_structure_for_a_serialized_inventory/)
- url: https://www.reddit.com/r/rails/comments/gn770m/database_structure_for_a_serialized_inventory/
---
Hello guys, would appreciate some help structuring my Database correctly. I'm building an inventory system where i need to track the location of every single product. Here's what I have so far:-

    class Product &lt; ApplicationRecord
      has_many :serial_numbers
    end
    
    class SerialNumber &lt; ApplicationRecord
      belongs_to :product
    end
    
    class SerializedStockMovement &lt; ApplicationRecord
      belongs_to :bin
      belongs_to :document, polymorphic: true # Sales invoice/Stock Adjustment etc
      belongs_to :serial_number
    end
    
    class Bin &lt; ApplicationRecord
      has_many :serialized_stock_movements
    end

Every time i wish to add an individual item into a bin,

    SerializedStockMovement.create(
      bin: @bin,
      serial_number: @serial_number,
      quantity: 1
    )

and to transfer to another bin,

    SerializedStockMovement.create(
      bin: @bin_one,
      serial_number: @serial_number,
      quantity: -1
    )
    
    SerializedStockMovement.create(
      bin: @bin_two,
      serial_number: @serial_number,
      quantity: 1
    )

It's working decently, but I just have a feeling that the structure can be better. A few of the problems with the structure currently is

* Quantity is always +1 or -1. Would it be better to be an enum "in" "out"?
* To get all the serial numbers currently in a bin, i'm having to do some fancy grouping.

&amp;#x200B;

    class Bin &lt; ApplicationRecord 
      has_many :serial_numbers, -&gt; { 
         group("serial_numbers.id, serial_numbers.*")
         .having("sum(serialized_stock_movements.quantity) &gt; 0") 
      }, through: :serialized_stock_movements
    end

&amp;#x200B;

* How best can I enforce a validation where the serial number is always only in ONE bin? I know how to do it in rails, but is there a way to enforce it in the DB?  


**Edit**  
Added UML

https://preview.redd.it/dv2zq25jgwz41.png?width=2508&amp;format=png&amp;auto=webp&amp;s=e336e2e835705cd7d0981bba1ff674b50a008719
## [5][Rails form for, form tag, form with - does anyone else have problems with these?](https://www.reddit.com/r/rails/comments/gn50bx/rails_form_for_form_tag_form_with_does_anyone/)
- url: https://www.reddit.com/r/rails/comments/gn50bx/rails_form_for_form_tag_form_with_does_anyone/
---
Hi guys, beginner developer here.  I was stuck on building my form for HOURS today and I'm just wondering if anyone else has issues.  I ended up resorting to an html form which worked fine in the end, but what I was trying to do was build a form without a model, but that would pass back to a strong params method in my controller.  For the life of me I couldn't get it to work without running into a no method error when I rendered the form again (if the user did not type in the correct input -- it's a question/answer type form).  When the form rendered again, the no method error would trigger because it seems like the form was expecting an instance attribute.  

So frustrating :(
## [6][Tracking and displaying recent posts.](https://www.reddit.com/r/rails/comments/gmziwa/tracking_and_displaying_recent_posts/)
- url: https://www.reddit.com/r/rails/comments/gmziwa/tracking_and_displaying_recent_posts/
---
Working on a (portfolio) site that has a private blog (only admin can create posts). I have most of the app already built, but I'm having a hard time trying to get attention from the main (home) page to the blog since it's not technically a blog site. 

I'm thinking of just making some (post)cards on the home page and making them live (as opposed to static) and tracking by the three most recent posts do display there. It will only be text displaying the title and a few (2-3) lines of body text. My problem specifically is I'm not sure how to get the cards to display the three most recent blog posts automatically. Is it just a simple post.last callback? with :title and :body attached? Or is there a tougher runaround on this?
## [7][Devise not hitting database.](https://www.reddit.com/r/rails/comments/gn6e64/devise_not_hitting_database/)
- url: https://www.reddit.com/r/rails/comments/gn6e64/devise_not_hitting_database/
---
Getting some kind of freak incident here, hoping some of you guys can help me figure out where the cycle is breaking.

&amp;#x200B;

`&lt;%= simple_form_for(resource, as: resource_name, url: registration_path(resource_name)) do |f| %&gt;`

  `&lt;%= f.error_notification %&gt;`

&amp;#x200B;

`&lt;div&gt;`

`&lt;%= f.input :email,`

`required: true,`

`autofocus: true,`

`input_html: { autocomplete: "email",`

`class: "blah" } %&gt;`

`&lt;/div&gt;`

&amp;#x200B;

`&lt;div&gt;`

`&lt;div&gt;`

`&lt;%= f.input :password,`

`required: true,`

`hint: ("#{@minimum_password_length} characters minimum" if u/minimum_password_length),`

`input_html: { autocomplete: "new-password",`

`class: "blah" } %&gt;`

&amp;#x200B;

`&lt;div&gt;`

`&lt;label class="block text-gray-700 text-sm font-bold mb-2"&gt;`

`&lt;%= f.input :password_confirmation,`

`required: true,`

`input_html: { autocomplete: "new-password",`

`class: "blah" } %&gt;`

`&lt;/div&gt;`

`&lt;div&gt;`

`&lt;div&gt;`

`&lt;%= f.button :submit, "Sign up", class: "blah" %&gt;`

`&lt;% end %&gt;`

`more closing divs and mumbo jumbo...`

When I submit the form, the db/terminal reflect the transaction, but it loops back around with this:

`Started GET "/users/sign_up?authenticity_token=Yfj0DZUtHHTk7LqhctXttGIsGUD0MHEyANKMXkaMaPKcNuFt6rKgRrZbJc2KFkwL%2B7fW0Re9%2F529uat1MwNXMg%3D%3D&amp;user%5Bemail%5D=jsmitty%`[`40email.com`](https://40email.com)`&amp;user%5Bpassword%5D=[FILTERED]&amp;user%5Bpassword_confirmation%5D=[FILTERED]&amp;commit=Sign+up" for ::1 at 2020-05-20 02:37:05 -0400`

`Processing by Devise::RegistrationsController#new as HTML`

  `Parameters: {"authenticity_token"=&gt;"Yfj0DZUtHHTk7LqhctXttGIsGUD0MHEyANKMXkaMaPKcNuFt6rKgRrZbJc2KFkwL+7fW0Re9/529uat1MwNXMg==", "user"=&gt;{"email"=&gt;`[`"jsmitty@email.com`](mailto:"jsmitty@email.com)`", "password"=&gt;"[FILTERED]", "password_confirmation"=&gt;"[FILTERED]"}, "commit"=&gt;"Sign up"}`

  `Rendering devise/registrations/new.html.erb within layouts/application`

  `Rendered devise/registrations/new.html.erb within layouts/application (Duration: 10.0ms | Allocations: 6465)`

`[Webpacker] Everything's up-to-date. Nothing to do`

  `Rendered layouts/_navbar.html.erb (Duration: 0.1ms | Allocations: 53)`

  `Rendered layouts/_footer.html.erb (Duration: 0.0ms | Allocations: 5)`

`Completed 200 OK in 22ms (Views: 20.6ms | ActiveRecord: 0.0ms | Allocations: 11351)`

&amp;#x200B;

I can't for the life of me figure out where this is going wrong. App already has a root route, I already tried manually redirecting to root route. I already tried the helpers (explicitly stating the user relations in the resource tag at the beginnig) but nothing seems to get the form through. I also tried manually creating the user via rails console, that worked. But I'm having much the same issue with the login bit as well. It just loops back around as I see it. The only thing that seems a little off on this app as opposed to the other apps that I've used devise on is a db:migration I ran for active text that ended up running an active storage bit that I don't understand. I'm assuming it's for file upload, as would be expected from a rich text editor. I don't see how that would have anything to do with the users though.
## [8][Help understanding if my transaction is nested.](https://www.reddit.com/r/rails/comments/gmtoig/help_understanding_if_my_transaction_is_nested/)
- url: https://www.reddit.com/r/rails/comments/gmtoig/help_understanding_if_my_transaction_is_nested/
---
I was following [https://api.rubyonrails.org/classes/ActiveRecord/Transactions/ClassMethods.html](https://api.rubyonrails.org/classes/ActiveRecord/Transactions/ClassMethods.html) documentation and came across the below example

    Account.transaction do
      balance.save!
      account.save!
    end

And from that link and this [https://makandracards.com/makandra/42885-nested-activerecord-transaction-pitfalls](https://makandracards.com/makandra/42885-nested-activerecord-transaction-pitfalls) I realize that nested transactions is part of the parent transaction and calling a rollback on the nested transaction doesn't propagate to the parent.   
Also, that save!, update\_attributes! creates its own transaction and hence making it a nested transaction. 

My code looks something like this, 

          ActiveRecord::Base.transaction do
            # destroy an related object (must be rolled back if below updates fail)
            Book.update_attributes!(params)
            Author.update_attributes!(params[nscname])
          end

and I write tests like so

    it "should not destroy related object if Book update fails"
      Book.any_instance.stubs(update_attributes!).raises(some_exception)
      assert related object is not deleted
    end

Now my problems are, as per my understanding the Book.update\_attributes! runs in a separate transaction and hence it would not rollback the related object, but it turns out it does and the tests succeed.

I have a few questions, 

1. Should I include requires\_new: true in the transaction?
2. Will the entire transaction rolled back in case book update fails, author update fails. Will this code lead to an invalid state?
## [9][Error to set up GCP active storage to be public](https://www.reddit.com/r/rails/comments/gn3orn/error_to_set_up_gcp_active_storage_to_be_public/)
- url: https://www.reddit.com/r/rails/comments/gn3orn/error_to_set_up_gcp_active_storage_to_be_public/
---
I active the public Access for GCP on Storage.yml  to be true 
When I tried to upload my file into my Rails app
I got error "ArgumentError (unknown keyword :public)"

How do you set public Access for active storage in GCP?
## [10][belongs_to &amp; has_many - do you really need both?](https://www.reddit.com/r/rails/comments/gmnmt6/belongs_to_has_many_do_you_really_need_both/)
- url: https://www.reddit.com/r/rails/comments/gmnmt6/belongs_to_has_many_do_you_really_need_both/
---
I am following a training example with microposts and users, and it says you need belongs\_to and has\_many. Is that not duplication? What happens if you only have one and not the other?
## [11][Friendly-id and a bug with the sequence of a title with two words](https://www.reddit.com/r/rails/comments/gmlktq/friendlyid_and_a_bug_with_the_sequence_of_a_title/)
- url: https://www.reddit.com/r/rails/comments/gmlktq/friendlyid_and_a_bug_with_the_sequence_of_a_title/
---
Following the tips of the author's gem (soruce: [https://github.com/norman/friendly\_id/issues/480](https://github.com/norman/friendly_id/issues/480)), I added a title\_and\_sequence like this

     extend FriendlyId   
     friendly_id :slug_candidates, use: [:slugged, :finders]
    
     def slug_candidates
         [:title, :title_and_sequence] end def title_and_sequence
         slug = title.to_param
         sequence = Book.where("slug LIKE '#{slug}--%'").count + 2
     "#{slug}--#{sequence}"
     end

but I noted that there is a bug. If I use as title "one", it works correctly

    /books/one 
    /books/one-2 
    /books/one-3 
    /books/one-4 etc.

But if I use a title with two words, for example "one two" I have a "bug" after the second result.

    /books/one-two 
    /books/one-two-2 
    /books/one-two-5j12-123j-afu4-jasdk 
    /books/one-two-9as6-k273-ewu1-87srt 
    etc.

**Do you have any idea about how to solve it?**
## [12][Explicitly rollback a transaction in a controller method and render a response as well.](https://www.reddit.com/r/rails/comments/gmqm8z/explicitly_rollback_a_transaction_in_a_controller/)
- url: https://www.reddit.com/r/rails/comments/gmqm8z/explicitly_rollback_a_transaction_in_a_controller/
---
I have an update action like so, 

    def update
       ActiveRecord::Base.transaction do
        # destroy an item
        result = call to an external method that returns false if an exception is raised
        unless result
          raise ActiveRecord::RollBack
        end
        # method that raises exception on update, not false
      end
      rescue
          # render_errors
          # same in case of rollback or any other exception
      end
    end

My problem, if the result is false then the transaction is rolled back but the rescue block isn't called and hence I do not get the required HTTP code and message. In case it failed because of any other exception the rescue block is hit and all is well. 

Could you please help me figure out how I can raise a rollback and handle that exception to call render\_errors as well.
