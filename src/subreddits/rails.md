# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gvu083/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gvu083/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/har6r7/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/har6r7/personal_projects_show_off_your_own_project_andor/
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
## [3][What are People Using for Localizing Date/Time?](https://www.reddit.com/r/rails/comments/hdkurj/what_are_people_using_for_localizing_datetime/)
- url: https://www.reddit.com/r/rails/comments/hdkurj/what_are_people_using_for_localizing_datetime/
---
The [Rails Locale Data Repository](https://github.com/svenfuchs/rails-i18n) seems to be the only comprehensive collection of formats but unfortunately I have not seen a single date/time format it uses that is an actual date time format **used in the given locale**. 

What are people using for their data/time localization formats?
## [4][Writing to a file is taking a lot of time in the prod environment](https://www.reddit.com/r/rails/comments/hdrubu/writing_to_a_file_is_taking_a_lot_of_time_in_the/)
- url: https://www.reddit.com/r/rails/comments/hdrubu/writing_to_a_file_is_taking_a_lot_of_time_in_the/
---
I know it's kind of weird. I have changed the in-memory generation of csv to a file-based one. Locally, it  generates the file relatively faster 20 k records in 20 mins or so. But in the production environment it's taking too much time. It's writing 1 k records over 20 mins. Is it because of I/O overhead? or am I looking at this wrong

This is the code I have

`tempfile = File.open(Rails.root.join('tmp', "#{SecureRandom.hex(8)}.csv"), 'wb')CSV.open(tempfile.path, 'w') do |csv|ActiveRecord::Base.uncached doModel.find_each do |model|csv &lt;&lt; ["foobar"]endendendtempfile`
## [5][How do you guys get acquainted with the huge codebase of an app with tons of functionality?](https://www.reddit.com/r/rails/comments/hd5xdu/how_do_you_guys_get_acquainted_with_the_huge/)
- url: https://www.reddit.com/r/rails/comments/hd5xdu/how_do_you_guys_get_acquainted_with_the_huge/
---
I got a job at a company where the app and the codebase are so huge and very poorly written tests. So, how should I get acquainted with this monstrosity?
## [6][Create Record As A User Signs Up](https://www.reddit.com/r/rails/comments/hd9yr4/create_record_as_a_user_signs_up/)
- url: https://www.reddit.com/r/rails/comments/hd9yr4/create_record_as_a_user_signs_up/
---
I am using devise facebook omniauth for User sign up. When a user signs up I want to create a record with user's name in particular table if it doesn't exist already, how do I go about doing this?
## [7][Does anyone know the answer to this? Has anyone faced this before? Sry if this is a basic error, but I'm new to programming.](https://www.reddit.com/r/rails/comments/hdbf7r/does_anyone_know_the_answer_to_this_has_anyone/)
- url: https://www.reddit.com/r/rails/comments/hdbf7r/does_anyone_know_the_answer_to_this_has_anyone/
---
Could not load 'active\_record/connection\_adapters/sqlite3\_adapter'. Make sure that the adapter in config/database.yml is valid. If you use an adapter other than 'mysql2', 'postgresql' or 'sqlite3' add the necessary adapter gem to the Gemfile.  
Couldn't create database for {"adapter"=&gt;"sqlite3", "pool"=&gt;5, "timeout"=&gt;5000, "database"=&gt;"db/development.sqlite3"}  
rails aborted! 

&amp;#x200B;

The adapter in my database.yml file is sqlite3 and even in my gemfile it is sqlite3. So after i host a server using 'rails server' command I get the above error. But after refreshing a few times, I get to the desired page which says 'Yay! You're on rails!' The problem is that i cant continue the scaffold command as it won't execute. When I write a command for scaffold, I get the error mentioned above and quite a lot of lines mentioning some error. Again, my error may be quite amateurish, but I legit have no idea what I should be changing. Thanks in advance.
## [8][How do I query for the most recently created records across multiple models?](https://www.reddit.com/r/rails/comments/hcu0i9/how_do_i_query_for_the_most_recently_created/)
- url: https://www.reddit.com/r/rails/comments/hcu0i9/how_do_i_query_for_the_most_recently_created/
---
## Models

```
class CoreCompetency &lt; ApplicationRecord
    has_many :scores
end
```

```
class TeamMember &lt; ApplicationRecord
    has_many :scores
end
```

```
class Score &lt; ApplicationRecord
    belongs_to :team_member
    belongs_to :core_competency
end
```

## Question

A `team_member` can create a `score` for several `core_competencies` at different times. How do I get a team_member's most recent `score` per `core_competency`?
## [9][Modify Instance variable in a haml?](https://www.reddit.com/r/rails/comments/hcohd4/modify_instance_variable_in_a_haml/)
- url: https://www.reddit.com/r/rails/comments/hcohd4/modify_instance_variable_in_a_haml/
---
I have a HAML with a modal. The modal contains a button, and the button’s href has instance variables so that when the Ajax call is made later (after clicking the button) the controller can find user data via java back end services (mysql).

I can’t seem to be able to set those instance variables correctly.

1. an Ajax call (of type get) is used to call the controller action which renders the modal containing the button.
2. The modal is supposed to pop up with the button containing the proper href (href containing the id and name, from the instance variables)
3. When the button is clicked it another Ajax call (of type put) should be made to act on that user.

I’m stuck between 1 and 2. I can post some code later but I hope this makes some sense about what I’m trying to do.

Thanks!

\---EDIT---

The modal to be rendered:

`.modal-body`

`= t('remittances.similar_refund_message', :transAm =&gt; @transAm, :EFTCheckDate =&gt; @EFTCheckDate)`

`=hidden_field_tag "action", @action`

`.footer.pull-right.padding-t-20`

`%a.btn.btn-primary{id:'save-refund-confirmation-button', href: update_refund_check_remittances_path(id: @remittance.id, payerName: @remittance.payer.name, checkDuplicate: false, filters: u/filters), 'data-remote'=&gt;'true', type: 'button'}= t('remittances.yes_continue')`

`%a.btn{"data-dismiss" =&gt; "modal", type: "button"}= t('remittances.no_go_back')`

The ruby controller action rendering the modal:

`def render_similar_update_refund_modal`

`@action = params['actionType']`

`@transAm = params['transAm']`

`@EFTCheckDate = params['EFTCheckDate']`

`render partial: 'similar_update_refund_modal_body'`

`end`

The ajax call to the controller action to render the modal:

`$.ajax({`

`url: errorModalUrl,`

`data: { actionType: actionType, transAm: transAmt, EFTCheckDate: EFTCheckDate }, success: function(response) {`

`$('.modal-body').replaceWith(response);`

`$('#remittance-error-button').click();`

`Terra.Overlay.hide();`

`},`

I think my confusion is how to use instance variables... this href works in the modal-body:`href: update_refund_check_remittances_path(checkDuplicate: false, isRefund: true)`

Basically what I want is to have the button in the modal body contain the href that has what I assume are instance variables. This is done on the save button on the bottom of my form page and it works just fine. That save button acts like a form submission button but instead of submitting, it serializes the form data and sends it via Ajax to the controller which directs the data to the java back end. When I take the href from the save button and put it in the modal button, it fails and the ajax call (from clicking the modal button) immediately hits the error branch.

&amp;#x200B;

&amp;#x200B;
## [10][How do you structure your models for a multi-lingual app?](https://www.reddit.com/r/rails/comments/hcdh15/how_do_you_structure_your_models_for_a/)
- url: https://www.reddit.com/r/rails/comments/hcdh15/how_do_you_structure_your_models_for_a/
---
I've been tasked with creating an e-commerce website, and I'm planning to build it in Rails. However, the client wants the content to be available in both English and Thai.

I'm a Rails newbie (and have also used Laravel and Django), so I'd like some advice on how to approach this. Would the best way be to have a `Product` model that contains the price, quantity, etc, and has a one-to-many relationship with the `ProductInfo` class, which contains the title and description. So `new_product` could have two `ProductInfo` objects, one with an English description and one with a Thai description. I'm guessing the same structure could be used for a `Category` model too.

What would be the practical way of structuring the data entry form for the `Product` class? Separate form for each language, or one form with the fields for both languages?
## [11][Different data presentation without querying database again](https://www.reddit.com/r/rails/comments/hcb7ah/different_data_presentation_without_querying/)
- url: https://www.reddit.com/r/rails/comments/hcb7ah/different_data_presentation_without_querying/
---
Hi I'm from nodejs/react background, and am learning ROR for fun for my side project. The reason why I chose ROR is SEO matters to me and I don't want to waste my time configuring SSR

Basically I want to display some data in two formats, cards and lists.

e.g. user hits url \`/cars\`, based on user saved preference, I'll render \`&lt;%= render card %&gt;\` or \`&lt;%= render list %&gt;\` in my template.

But I also want to allow user to toggle each view by a button in UI. However in ROR's server side mindset, if I understand correctly, when user changes card view to list view, I probably need to send an ajax and server returns \`\*.js.erb\` and replace the entire cards section with list section.

The problem is this inevitably query the database second time to render \`js.erb\`. On a higher level I  need to send ajax to reuse the erb template even though I have all the necessary data already in UI DOM.

One workaround I can think of is to render both \`cards\` and \`list\` in same page, and set \`display: none\` to either of it. But that means I need to render same data twice on the page. Is this considered a good approach?

from React POV, this is pretty straightforward where I can get the data, and use UI state to determine if I need to render cards or list without sending extra ajax. 

What's considered the best practice of doing this without much hassle? Thanks in advance!
## [12][Would you build your app in Rails in 2020?](https://www.reddit.com/r/rails/comments/hbpmc1/would_you_build_your_app_in_rails_in_2020/)
- url: https://www.reddit.com/r/rails/comments/hbpmc1/would_you_build_your_app_in_rails_in_2020/
---
I have a couple of (what I think are) good ideas floating around in my head, and trying to decide what to build them in.

I've dabbled in Rails, as well as full stack JS, but never built anything **serious** with either.

A few years ago I was really looking into becoming a professional Web Developer. I was working tech support for a web design company but had hit a dead end with my career there. So I was doing the typical self studying thing, Coursera, Codecademy, Hartl's tutorial, etc, getting ready to go down that route, before I kind of fell into an Application Support engineer position with a startup my friend was working in, and I have been doing that ever since. Now I support the same application for a self driving car startup making well over 6 figures.

Now that I am in a pretty good place financially and career wise, and I am a little older and have a little better work ethic, I want to take a serious crack at actually making something out of these ideas I have. Back when I was first thinking about this, a few years ago, I was pretty sure I was going to build them in Rails. Rails was pretty popular then, I was learning about it and found it really easy to use and I really liked how fast you could get something up. But it seems like in the few years I wasn't really paying attention, Rails has really fallen out of vogue, as it were. It doesn't scale. People are listing Ruby as their most hated language.  It's really hard to make anything other than an old fashioned Monolith style app with it. Etc. etc.

If you were starting a brand new project in 2020 one that you were hoping to actually monetize and go commercial with, would still doing it in Rails be a decision that you'd regret later, or is Rails fine, just everyone likes to crap on it because it's not the new hotness anymore?
