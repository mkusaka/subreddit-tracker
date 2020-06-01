# rails
## [1][Absolute beginner building a web app! :D](https://www.reddit.com/r/rails/comments/gugj7a/absolute_beginner_building_a_web_app_d/)
- url: https://www.reddit.com/r/rails/comments/gugj7a/absolute_beginner_building_a_web_app_d/
---
Hey there Rails community! I am building a beginner web app . 

The app contains users that register/log in (using devise) and bus tickets that are shown on the welcome page. Everyone can see the bus tickets, but only users that are registered+logged in can buy (its fake buying :D). Additionally, users that have bought the ticket have the option to cancel it (no later than 1 hour before bus starts).

The tricky part here for me is how to display the tickets that the user already bought, and then to add the option of cancelling. I thought maybe to add a third table that would connect users and tickets, and then display that on the "My tickets" page? Not sure how to accomplish that. 

If anyone could chip in with a bit of advice I would be really grateful! Thanks guys.

tickets(bus:string, time:datetime, quantity:integer, price:integer)
users (standard devise table with email, password and so on)
## [2][How to rescue custom exceptionin transaction block while still rolling back everything?](https://www.reddit.com/r/rails/comments/guk18l/how_to_rescue_custom_exceptionin_transaction/)
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
## [3][How can I sort an array of hashes into separate arrays based on key value uniqueness?](https://www.reddit.com/r/rails/comments/gu8nmc/how_can_i_sort_an_array_of_hashes_into_separate/)
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
## [4][JavaEE ala Rails](https://www.reddit.com/r/rails/comments/gu4h9q/javaee_ala_rails/)
- url: https://www.reddit.com/r/rails/comments/gu4h9q/javaee_ala_rails/
---
I took something I learned with rails and revisited JavaEE recreating the rails crud with JSF, it is not fashionable but seems better than nowaday development.  Rails developers can identify what I did here?  [JSF-PERFECT-CRUD](https://github.com/lazaronixon/jsf-perfect-crud)
## [5][My first VSCode extension for Rails tests](https://www.reddit.com/r/rails/comments/gttg2k/my_first_vscode_extension_for_rails_tests/)
- url: https://www.reddit.com/r/rails/comments/gttg2k/my_first_vscode_extension_for_rails_tests/
---
Having switched recently from rspec to minitest I missed an easy way to switch between test files and code. There is an extension for Rspec already - [Rails-go-to-spec](https://marketplace.visualstudio.com/items?itemName=sporto.rails-go-to-spec), based on that I created a [version for minitest](https://marketplace.visualstudio.com/items?itemName=SimonBo.rails-go-to-test). It creates a new test file with boilerplate or opens the existing test file. Enjoy!
## [6][How do I add 'belongs_to' to existing Model](https://www.reddit.com/r/rails/comments/gtxdnr/how_do_i_add_belongs_to_to_existing_model/)
- url: https://www.reddit.com/r/rails/comments/gtxdnr/how_do_i_add_belongs_to_to_existing_model/
---
So I was following this guide  [The has\_one :through Association](https://guides.rubyonrails.org/association_basics.html#the-has-one-through-association) All seem good, just I was wondering how do I add 't.belongs\_to :supplier' part to existing model/scheme?

    class CreateAccountHistories &lt; ActiveRecord::Migration[5.0]
      def change
        create_table :suppliers do |t|
          t.string :name
          t.timestamps
        end
     
        create_table :accounts do |t|
          t.belongs_to :supplier   &lt;&lt;&lt;&lt;&lt;&lt; THIS!!!
          t.string :account_number
          t.timestamps
        end
     
        create_table :account_histories do |t|
          t.belongs_to :account
          t.integer :credit_rating
          t.timestamps
        end
      end
    end
## [7][Rails but not only for API](https://www.reddit.com/r/rails/comments/gtyuqw/rails_but_not_only_for_api/)
- url: https://www.reddit.com/r/rails/comments/gtyuqw/rails_but_not_only_for_api/
---
Hi guys, I was wondering if a normal rails app (created with just a rails new) can be used as an API. In my situation i want to create a website and a react native app.

So do i need work on my rails like usual, or do i need to create 2 different path in my routes ? Or maybe something else?
## [8][Question about retrieving/streaming media from AWS S3 in Rails.](https://www.reddit.com/r/rails/comments/gtttn8/question_about_retrievingstreaming_media_from_aws/)
- url: https://www.reddit.com/r/rails/comments/gtttn8/question_about_retrievingstreaming_media_from_aws/
---
Hi guys!

So I’m new to PosgreSQL usage with Rails so please bear with me.

I’m building a Udemy clone as a small project and I was wondering how I would stream videos from my Amazon S3 buckets.

So I’m using PosgreSQL to store the name of the video, author, and date published. But how do I retrieve the video files in a way where it would only allow subscribed users to view the videos?

Also, how would I be able to refer to the media in S3 from PostgreSQL?

Would this be handled with the controller rather than with AWS S3?

Any help or insight would be much appreciated! Thank you!
## [9][Stackoverflow 2020 developer survey results](https://www.reddit.com/r/rails/comments/gtm7kg/stackoverflow_2020_developer_survey_results/)
- url: https://www.reddit.com/r/rails/comments/gtm7kg/stackoverflow_2020_developer_survey_results/
---
IMHO the results of the survey show Ruby and Rails in bad company:

* Most loved / dreaded languages (Ruby):

 [https://insights.stackoverflow.com/survey/2020#technology-most-loved-dreaded-and-wanted-languages](https://insights.stackoverflow.com/survey/2020#technology-most-loved-dreaded-and-wanted-languages) 

**=&gt;** Ruby is 7th most dreaded after VBA, Objective, Perl, Assembly, C, PHP

* Most loved / dreaded frameworks (Rails):

[https://insights.stackoverflow.com/survey/2020#technology-most-loved-dreaded-and-wanted-web-frameworks](https://insights.stackoverflow.com/survey/2020#technology-most-loved-dreaded-and-wanted-web-frameworks)

**=&gt;** Rails is 6th most dreaded after Angular, Drupal, jQuery, ASP, Symfony 

&amp;#x200B;

What do you think about it?
## [10][Oauth/OmniAuth/etc.](https://www.reddit.com/r/rails/comments/gtq9et/oauthomniauthetc/)
- url: https://www.reddit.com/r/rails/comments/gtq9et/oauthomniauthetc/
---
Quick question... Would it be unwise to limit website logins to OAuth services? Meaning, I don't use devise and have them create users, instead just have them sign in via whatevertf they wanna use? I've personally enjoyed having that flexibility on my end and use it fairly often. That being said:

1) Would you use a website that only takes in OAuth-type registrations?

2) Would devise be the best way to handle this still? or is there a better option available with ruby 2.7 &amp; Rails 6? (Saw the omniauth section of Devise says it works best with 2.5 so I'm not sure on this).
