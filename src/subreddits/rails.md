# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gnbebg/personal_projects_show_off_your_own_project_andor/)
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
## [2][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/gs526t/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/gs526t/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [3][My first VSCode extension for Rails tests](https://www.reddit.com/r/rails/comments/gttg2k/my_first_vscode_extension_for_rails_tests/)
- url: https://www.reddit.com/r/rails/comments/gttg2k/my_first_vscode_extension_for_rails_tests/
---
Having switched recently from rspec to minitest I missed an easy way to switch between test files and code. There is an extension for Rspec already - [Rails-go-to-spec](https://marketplace.visualstudio.com/items?itemName=sporto.rails-go-to-spec), based on that I created a [version for minitest](https://marketplace.visualstudio.com/items?itemName=SimonBo.rails-go-to-test). It creates a new test file with boilerplate or opens the existing test file. Enjoy!
## [4][How do I add 'belongs_to' to existing Model](https://www.reddit.com/r/rails/comments/gtxdnr/how_do_i_add_belongs_to_to_existing_model/)
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
## [5][Question about retrieving/streaming media from AWS S3 in Rails.](https://www.reddit.com/r/rails/comments/gtttn8/question_about_retrievingstreaming_media_from_aws/)
- url: https://www.reddit.com/r/rails/comments/gtttn8/question_about_retrievingstreaming_media_from_aws/
---
Hi guys!

So I’m new to PosgreSQL usage with Rails so please bear with me.

I’m building a Udemy clone as a small project and I was wondering how I would stream videos from my Amazon S3 buckets.

So I’m using PosgreSQL to store the name of the video, author, and date published. But how do I retrieve the video files in a way where it would only allow subscribed users to view the videos?

Also, how would I be able to refer to the media in S3 from PostgreSQL?

Would this be handled with the controller rather than with AWS S3?

Any help or insight would be much appreciated! Thank you!
## [6][Rails but not only for API](https://www.reddit.com/r/rails/comments/gtyuqw/rails_but_not_only_for_api/)
- url: https://www.reddit.com/r/rails/comments/gtyuqw/rails_but_not_only_for_api/
---
Hi guys, I was wondering if a normal rails app (created with just a rails new) can be used as an API. In my situation i want to create a website and a react native app.

So do i need work on my rails like usual, or do i need to create 2 different path in my routes ? Or maybe something else?
## [7][Stackoverflow 2020 developer survey results](https://www.reddit.com/r/rails/comments/gtm7kg/stackoverflow_2020_developer_survey_results/)
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
## [8][Oauth/OmniAuth/etc.](https://www.reddit.com/r/rails/comments/gtq9et/oauthomniauthetc/)
- url: https://www.reddit.com/r/rails/comments/gtq9et/oauthomniauthetc/
---
Quick question... Would it be unwise to limit website logins to OAuth services? Meaning, I don't use devise and have them create users, instead just have them sign in via whatevertf they wanna use? I've personally enjoyed having that flexibility on my end and use it fairly often. That being said:

1) Would you use a website that only takes in OAuth-type registrations?

2) Would devise be the best way to handle this still? or is there a better option available with ruby 2.7 &amp; Rails 6? (Saw the omniauth section of Devise says it works best with 2.5 so I'm not sure on this).
## [9]["block in execute_hook". My app wont show up on heroku after implementing aws s3](https://www.reddit.com/r/rails/comments/gtp4dw/block_in_execute_hook_my_app_wont_show_up_on/)
- url: https://www.reddit.com/r/rails/comments/gtp4dw/block_in_execute_hook_my_app_wont_show_up_on/
---
Could it be my active storage configuration.  I specified has\_one\_attached in the model as well as adding strong params. I believe I also did the configuration correctly as well on heroku. Does anyone have any ideas? 

&amp;#x200B;

https://preview.redd.it/df4cbtimuz151.png?width=1398&amp;format=png&amp;auto=webp&amp;s=0fd1ae736c3647ebb8b7b40afdf771966c7dd760
## [10][Error loading the 'sqlite3' Active Record adapter. Missing a gem it depends on? sqlite3 is not part of the bundle. Add it to your Gemfile for heroku](https://www.reddit.com/r/rails/comments/gtnvj8/error_loading_the_sqlite3_active_record_adapter/)
- url: https://www.reddit.com/r/rails/comments/gtnvj8/error_loading_the_sqlite3_active_record_adapter/
---
 Hello guys I do not know why but I am getting this strange error on deployment. I have been spending three hours on this but can not get it to work. It has successfully worked before many times and even work once in the three hours deploying a new app. I do not know why I am getting this error on heroku and only heroku. 

&amp;#x200B;

[database.yml](https://preview.redd.it/o5cqganngz151.png?width=1901&amp;format=png&amp;auto=webp&amp;s=400833825be8ae5cc485a28891df7a90c922e3ae)

[Heroku Error](https://preview.redd.it/ccs3acnngz151.png?width=1570&amp;format=png&amp;auto=webp&amp;s=aa93dc062088e39392aa32237052b92fd4e1b98a)

[gemfile](https://preview.redd.it/q26b5dnngz151.png?width=1272&amp;format=png&amp;auto=webp&amp;s=278170a0de474c157e56213ea26618c21425edda)

&amp;#x200B;

https://preview.redd.it/ymca5nuwgz151.png?width=987&amp;format=png&amp;auto=webp&amp;s=76426e30f4357e9dd9883af34cd539ca4af8abd5

Any help would be appreciated !!!!
## [11][Ahoy vs Google Analytics](https://www.reddit.com/r/rails/comments/gtfbf9/ahoy_vs_google_analytics/)
- url: https://www.reddit.com/r/rails/comments/gtfbf9/ahoy_vs_google_analytics/
---
Right now I'm making a SaaS and want to have a dashboard for each user to track their own sales. I've heard about Ahoy gem being the go-to for this job but I wonder about the feature comprison between these two choices apart from privacy issue. Is there a significant reason to choose one over another? Feature-wise, which is your choice of tool?
## [12][Link_to 'blah', post.show(id), etc.](https://www.reddit.com/r/rails/comments/gtj5c4/link_to_blah_postshowid_etc/)
- url: https://www.reddit.com/r/rails/comments/gtj5c4/link_to_blah_postshowid_etc/
---
I know that won't work, but that's basically what I'd like to do. I already have a working link for the post, but there's another section on the page that I'd like to link specifically to that post ID in a static manner. The post will always be post id #1. What's the proper syntax here?
