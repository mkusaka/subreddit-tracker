# rails
## [1][Resources for Fancy Links? Specifically trying to populate user/email field if redirected to login after click](https://www.reddit.com/r/rails/comments/gv3ydx/resources_for_fancy_links_specifically_trying_to/)
- url: https://www.reddit.com/r/rails/comments/gv3ydx/resources_for_fancy_links_specifically_trying_to/
---
The business that I’m designing for has a lot of stubborn, crotchety customers, and in this new infrastructure, I’m trying to save the employees as much time as possible by empowering customers (who don’t know how to search their email, or google a tracking number, or make it two clicks past a login page).

What’s the best way for including params in a link that then update the value of an input?

I feel like I could just sit down and build a param trigger to render JS and put logic in a js.erb file, but something in my gut is saying that tailoring to unique params situations is usually a job for a pack full of jquery
## [2][Order of Callbacks in rails 4](https://www.reddit.com/r/rails/comments/gv3k0n/order_of_callbacks_in_rails_4/)
- url: https://www.reddit.com/r/rails/comments/gv3k0n/order_of_callbacks_in_rails_4/
---
I have a model say \`Item\` which has a method called on \`after\_commit\` but at the same time, the item model also has an Observer called \`ItemObserver\`. The observer also has an \`after\_commit\` callback which resets the cache.

Now my issue is on the update of the \`item\` model the \`after\_commit\` on the model is called first before the \`after\_commit\` on the observer hence the cache isn't reset and the \`after\_commit\` runs on stale data learning to invalid actions.

I wanted to know if it's possible to force the observer's \`after\_commit\` to run first so that is refreshed with the latest portal for all subsequent callbacks.
## [3][Is it poor form to have to supply an array as an argument to initialize, even if there's only one object in the array?](https://www.reddit.com/r/rails/comments/guv73y/is_it_poor_form_to_have_to_supply_an_array_as_an/)
- url: https://www.reddit.com/r/rails/comments/guv73y/is_it_poor_form_to_have_to_supply_an_array_as_an/
---
Hi all,

I have a class which is initialised with 3 arguments.

`def initialize(arg1:, arg2:, array_of_objects:)
  @arg1 = arg1
  @arg2 = arg2
  @array_of_objects = array_of_objects
end`

My code requires that I run a method on each object in my_class_instance.array_of_objects. There can be one or more objects in that array.

At the moment I'm initializing like this:

`MyClass.new("foo", "bar", [MyOtherClass.new])`

If I don't do it this way, then when I try to run .each on the array_of_objects, it throws an error.

Is this bad practice? And if so, how should I be supplying n number of objects to the initialize method?

Thanks and apologies for the formatting, I only have access to my phone right now.
## [4][Anybody have any experience with House of Dubs?](https://www.reddit.com/r/rails/comments/gutn5u/anybody_have_any_experience_with_house_of_dubs/)
- url: https://www.reddit.com/r/rails/comments/gutn5u/anybody_have_any_experience_with_house_of_dubs/
---
Just got a freelance offer from them stating a budget and what they need. It's a static informational site which I'm hoping to code top-bottom from scratch. More for practice, though I could just as easily go with a hosting service like GoDaddy, WordPress, Wix, etc. by the looks of it. Curious if anyone has experience with this and knows more or less how I should handle billing to make sure everything plays out well for both parties. Rails related because I'm enjoying using rails and want to bring a hammer to do a scissor's job.
## [5][donorbox - what is it like working for them?](https://www.reddit.com/r/rails/comments/guynpd/donorbox_what_is_it_like_working_for_them/)
- url: https://www.reddit.com/r/rails/comments/guynpd/donorbox_what_is_it_like_working_for_them/
---

## [6][Conventions...](https://www.reddit.com/r/rails/comments/guuqg5/conventions/)
- url: https://www.reddit.com/r/rails/comments/guuqg5/conventions/
---
Ruby on Rails Conventions is important to develop as a  team. 😊

&amp;#x200B;

Controller -&gt; plural.

Model -&gt; singular

Migration -&gt; plural.

&amp;#x200B;

but sometimes weird when generate them 😅
## [7][How to rescue custom exceptionin transaction block while still rolling back everything?](https://www.reddit.com/r/rails/comments/guk18l/how_to_rescue_custom_exceptionin_transaction/)
- url: https://www.reddit.com/r/rails/comments/guk18l/how_to_rescue_custom_exceptionin_transaction/
---
Suppose I would like to add a condition inside a transaction block which would raise an exception that would be raised, rescued, but I would still like all previous saves to be rolled back:

    ActiveRecord::Base.transaction do
        object1.save!
        object2.save!
        if condition_is_met
            raise CustomNameStandardError.new error_msg_string
        end
        object3.save!
    rescue CustomNameStandardError =&gt; e
        flash[:danger] = e.message
        redirect_to page
    end

with defining CustomNameStandardError within the same controller class like so

    class CustomNameStandardError &lt; StandardError; end

This will not rollback saves to object1 and object2. This is because by rescuing the exception I do not trigger rollback. How can I trigger rollback, but still be able to redirect myself back to the current page with an error message?

**EDIT: I found the answer and should have read the docs carefully. Apperantly raising ActiveRecord::Rollback does exactly what I wanted in my post:**

    ActiveRecord::Base.transaction do
        object1.save!
        object2.save!
        if condition_is_met
            flash[:danger] = error_msg_string
            raise ActiveRecord::Rollback
        end
        object3.save!
        redirect_to success_page
    end
    
    redirect_to where_i_came_from
## [8][Absolute beginner building a web app! :D](https://www.reddit.com/r/rails/comments/gugj7a/absolute_beginner_building_a_web_app_d/)
- url: https://www.reddit.com/r/rails/comments/gugj7a/absolute_beginner_building_a_web_app_d/
---
Hey there Rails community! I am building a beginner web app . 

The app contains users that register/log in (using devise) and bus tickets that are shown on the welcome page. Everyone can see the bus tickets, but only users that are registered+logged in can buy (its fake buying :D). Additionally, users that have bought the ticket have the option to cancel it (no later than 1 hour before bus starts).

The tricky part here for me is how to display the tickets that the user already bought, and then to add the option of cancelling. I thought maybe to add a third table that would connect users and tickets, and then display that on the "My tickets" page? Not sure how to accomplish that. 

If anyone could chip in with a bit of advice I would be really grateful! Thanks guys.

tickets(bus:string, time:datetime, quantity:integer, price:integer)
users (standard devise table with email, password and so on)
## [9][How can I sort an array of hashes into separate arrays based on key value uniqueness?](https://www.reddit.com/r/rails/comments/gu8nmc/how_can_i_sort_an_array_of_hashes_into_separate/)
- url: https://www.reddit.com/r/rails/comments/gu8nmc/how_can_i_sort_an_array_of_hashes_into_separate/
---
Hi all,  


Having trouble implementing a coding challenge.  


So I have an array of objects:

    [{
        :color=&gt;"red",
        :fruit=&gt;"apple"
      },
      {
        :color=&gt;"green",
        :fruit=&gt;"pear"
      },
      {
        :color=&gt;"red",
        :fruit=&gt;"tomato"
      }]

Now, I want to sort this into separate arrays so that I can count the number of fruit of each color. So in this case I'll have one array with `apple` and `tomato`, and another with `pear`.  


One additional catch is that I can have infinite fruit to sort, and I can have infinite different colors represented by the fruit, and I don't know what those colors will be. So this method needs to be dynamic enough to handle that.  


I just need it to recognise each unique color which is present in the objects within the array, and then put them in a corresponding array of other fruits of that unique color. So desired output will in this case be:  


    [
      {
        :color=&gt;"red",
        :fruit=&gt;"apple"
      },
      {
        :color=&gt;"red",
        :fruit=&gt;"tomato"
      }
    ],
    [
      {
        :color=&gt;"green",
        :fruit=&gt;"pear"
      }
    ]


Hope this makes sense.  


Thanks
## [10][JavaEE ala Rails](https://www.reddit.com/r/rails/comments/gu4h9q/javaee_ala_rails/)
- url: https://www.reddit.com/r/rails/comments/gu4h9q/javaee_ala_rails/
---
I took something I learned with rails and revisited JavaEE recreating the rails crud with JSF, it is not fashionable but seems better than nowaday development.  Rails developers can identify what I did here?  [JSF-PERFECT-CRUD](https://github.com/lazaronixon/jsf-perfect-crud)
