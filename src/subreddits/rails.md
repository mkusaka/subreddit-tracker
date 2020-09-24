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
## [3][automatically getting more inputs and how to handle them in the backend](https://www.reddit.com/r/rails/comments/iynsdf/automatically_getting_more_inputs_and_how_to/)
- url: https://www.reddit.com/r/rails/comments/iynsdf/automatically_getting_more_inputs_and_how_to/
---
I know this might be more of a frontend issue, but I have the habit of having anything in one place (I think this might be a harmful thought, but I do most of my projects alone, so I keep them the way I can handle!). So, I'm a musician and when I was uploading an album on bandcamp, I noticed something. 

When you use "album", the damn thing automatically gives you a field of new song. I thing that be a really good feature to have, but I don't know what that's called and how I can have on that a rails project. If you guys have any experiences, I'd love to know. 

Thanks!
## [4][Having trouble with wizard gem](https://www.reddit.com/r/rails/comments/iyegbd/having_trouble_with_wizard_gem/)
- url: https://www.reddit.com/r/rails/comments/iyegbd/having_trouble_with_wizard_gem/
---
I can't seem to access the wizard without receiving this error "**The requested step did not match any steps defined for this controller**". I don't know why its not starting the steps process when I send the user over to the controller for Wicked with the business form id. I'm also using friendly\_id btw.

I'm passing the id of the object to the path.

*routes.rb*

`resources  :business_form_wizard`

*link in view*

`&lt;%= link_to 'form wizard', form_wizard_path(current_provider.business_form), class:...%&gt;`

business\_*form\_wizard\_controller.rb*

`class BusinessFormWizardController &lt; ApplicationController`

`include Wicked::Wizard`

`steps :info, :scope_of_serv,...`

`def show`

 `@business_form = BusinessForm.friendly.find(params[:id])`

 `render_wizard`

`end`

&amp;#x200B;

`def update`

 `@ece_form.update_attributes(params[:business_form])`

`render_wizard @business_form`

`end`
## [5][Rails: testing with MySQL](https://www.reddit.com/r/rails/comments/iyc18w/rails_testing_with_mysql/)
- url: https://www.reddit.com/r/rails/comments/iyc18w/rails_testing_with_mysql/
---
Hi. My Rails app uses MySQL for prod and dev, but uses SQLite for testing. The problem with that is that everytime I run the tests, it overwrites the schema.rb file to fit SQLite instead of MySQL.

The problem with using MySQL for testing is that Rails always tries to access a database called something like `name_here-1`. That `-1` is illegal in MySQL, so I can't even create that database.

Is it possible to do MySQL-based tests or do I have to stick to SQLite for that?
If I have to stick to SQLite, how can I stop Rails from overwriting `schema.rb`?
## [6][flag_shih_tzu gem , return same value for any key](https://www.reddit.com/r/rails/comments/iy9u1k/flag_shih_tzu_gem_return_same_value_for_any_key/)
- url: https://www.reddit.com/r/rails/comments/iy9u1k/flag_shih_tzu_gem_return_same_value_for_any_key/
---
I'm using **flag\_shih\_tzu** gem and at my User model I include that lines  


    include FlagShihTzu
    
       has_flags  2 =&gt; :admin,
                  3 =&gt; :moderator

I have a column **flags** as well in User table  **t.integer "flags", default: 1, null: false**  


but when I set Flags value to 2 or 3 it doesn't matter It will always return me the same value admin? for example  


    user = User.first
    user.flags = 2
    user. save 
    
    User.first.admin? 
    true 
    
    user = User.second  
    user.flags = 3
    user.save
    
    User.second.admin?
    true
    User.second.moderator?
    false
## [7][Where to start](https://www.reddit.com/r/rails/comments/iy3dmv/where_to_start/)
- url: https://www.reddit.com/r/rails/comments/iy3dmv/where_to_start/
---
Hey everyone,

I have an interview for a Java/Ruby position next Monday. I am proficient in Java but I have no exposure to Ruby and I am looking for resources to get started. I read the first few chapters of 'Eloquent Ruby' this afternoon and it seems to share similarities with Python/Elixir - both of which I have experience with. 

Any advice would be appreciated :)

Thanks in advance!
## [8][Why is dependent: :destroy not working for me?](https://www.reddit.com/r/rails/comments/ixwox5/why_is_dependent_destroy_not_working_for_me/)
- url: https://www.reddit.com/r/rails/comments/ixwox5/why_is_dependent_destroy_not_working_for_me/
---
I have two models:

    class Detection &lt; ApplicationRecord
      has_many :links, foreign_key: 'start_id', dependent: :destroy
      has_many :links, foreign_key: 'end_id', dependent: :destroy
    end
    
    class Link &lt; ApplicationRecord
      belongs_to :start, class_name: 'Detection', foreign_key: :start_id
      belongs_to :end, class_name: 'Detection', foreign_key: :end_id
    end

When I delete a detection, I want the associated links to be destroyed. However, I'm getting a failed foreign-key constraint from pg when I try to delete a Detection, suggesting that my associations are not being described correctly here:

    PG::ForeignKeyViolation: ERROR: update or delete on table "detections" violates foreign key constraint "fk_rails_8633ff3bb0" on table "links" DETAIL: Key (id)=(1) is still referenced from table "links".


How do I tell Rails to delete any links where start\_id or end\_id match the deleted Detection's id?
## [9][Is there a better way to do what I just did??](https://www.reddit.com/r/rails/comments/ixw4kt/is_there_a_better_way_to_do_what_i_just_did/)
- url: https://www.reddit.com/r/rails/comments/ixw4kt/is_there_a_better_way_to_do_what_i_just_did/
---
Hi all, I’m a Rails beginner and just implemented a feature that has resulted in using multiple conditionals in my controllers and views, and am just wondering if there might be a more elegant way to do this?

I’m working on a practice app that lets you create online Invoices, and now finished a letting you create Estimates as well....which are simply Invoice records, but with a “estimate” Boolean attribute set to TRUE. 

So I have a “New Invoice” button that calls the ‘new’ action in the Invoices controller, AND I have a “New Estimate” button that also calls that same action, but passes in a “estimate: true” hash. 

Then, in the controller, I use an if/else conditional to check for that hash, and if it’s there, I create a new record with the ‘estimate’ Boolean set to true. Otherwise, I create a new record without setting the estimate boolean. 

Finally, for my Invoice views (show and  \_form, along with Active Mailer views) I also do a conditional check so the views properly display the words “Invoice” or “Estimate” depending on what the record’s estimate Boolean is, and also displays a different ‘Back’ or ‘Cancel’ button which brings you back to either the Invoices or Estimates index page (speaking of which, buttons to the Invoices page and Estimates page work the same way, with a conditional passed to the Index controller action, and that action using a conditional to serve up either invoices or estimates).

So that’s how it works, and it DOES work, but I just wonder if there might be a better, cleaner way to do something like this? I haven’t figured one out despite some research, but Rails is filled with so many clever tools, I thought I’d ask here!

Thanks very much!
## [10][Becoming a regular Rails contributor?](https://www.reddit.com/r/rails/comments/ixn3qh/becoming_a_regular_rails_contributor/)
- url: https://www.reddit.com/r/rails/comments/ixn3qh/becoming_a_regular_rails_contributor/
---
Are there any regular Rails contributors here?

It's been an old goal of mine to become one. Now I"m not expecting to reach Tenderlove's level, but I want to get to the level where I have at least a vague idea how to solve regular bugs in ActiveRecord. That I know enough about the internals I can come up with solutions given enough time. When I now look at Rails bugs, especially ActiveRecord, I usually wouldn't know even what files are causing the problem - is it AssociationReflection? AssociationProxy maybe?. When I look at someone's bug fix it sorta makes sense,  but the interaction between all the different classes and modules is too much for my brain to handle.

I guess I need to practice more, but it seems like ordinary bugs are often times too hard for me to practice on. It seems like regular contributors have a tons of contextual knowledge on ActiveRecord that I'm missing.

P.S : I'm not sure why but I'm more interested in ActiveRecord than other parts of Rails, and it makes sense to me to focus on one part in any case.

Has anyone here made the transition to becoming a regular contributor and care to share how this came to be?
## [11][3rd party vendor SaaS is sending us us 'action' variable in POST](https://www.reddit.com/r/rails/comments/ixj48w/3rd_party_vendor_saas_is_sending_us_us_action/)
- url: https://www.reddit.com/r/rails/comments/ixj48w/3rd_party_vendor_saas_is_sending_us_us_action/
---
Hi,

Got a problem. 3rd party vendor SaaS is sending us 'action' variable in POST, which is getting overwritten by the Rails routing with the method name.

Do you have an idea for a workaround? 

The only idea we have now is to create a front app that will receive such payload and will change the action parameter to action2 (or something) and then it will send it to our rails app.
## [12][Decorating API response before](https://www.reddit.com/r/rails/comments/ixkeic/decorating_api_response_before/)
- url: https://www.reddit.com/r/rails/comments/ixkeic/decorating_api_response_before/
---
Hi everyone, so I'm back after [this](https://www.reddit.com/r/rails/comments/it350w/connecting_to_wordpress_db_from_rails/). I've essentially created an adapter to connect the crud operations from WP API so that it can be used in our app. We've stored some of the posts' data inside our DB and sync regularly for updated content.

Currently, the front end is fetching the posts via the WP API but now, we want to add some fields to the posts records using information we have in our DB.

The original plan would be to simply send the related information (with post ids) to the front end and cross-check them with the fetched post ids, then combine them. However, it doesn't feel really clean IMO.

My question is would it be better to fetch the posts on the back end, "decorate" the response with the new fields, then sending the decorated response to the front end? Seems like a roundabout way of doing things but is there a pros/cons to it? Or if anyone else have any suggestions would be much appreciated!
