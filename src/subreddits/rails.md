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
## [3][How do you guys get acquainted with the huge codebase of an app with tons of functionality?](https://www.reddit.com/r/rails/comments/hd5xdu/how_do_you_guys_get_acquainted_with_the_huge/)
- url: https://www.reddit.com/r/rails/comments/hd5xdu/how_do_you_guys_get_acquainted_with_the_huge/
---
I got a job at a company where the app and the codebase are so huge and very poorly written tests. So, how should I get acquainted with this monstrosity?
## [4][The Power of the Front-End Dev (in rails). A question about new DB Areas](https://www.reddit.com/r/rails/comments/hd6np2/the_power_of_the_frontend_dev_in_rails_a_question/)
- url: https://www.reddit.com/r/rails/comments/hd6np2/the_power_of_the_frontend_dev_in_rails_a_question/
---
I'm just a front end dev. And I use ONLY notepad++ to make my edits.

But recently I noted that to create a new DB area, a boolean titled adult\_content, the back end dev made this edit

in book.rb (in model) he added `#  adult_content    :boolean          default(FALSE)`

    # == Schema Information
    #
    # Table name: books
    #
    #  id               :integer          not null, primary key
    #  title            :string           not null
    #  slug             :string           not null
    #  description      :text             not null
    #  adult_content    :boolean          default(FALSE)

in db/schema.rb, he changed the version of ActiveRecord::Schema.define (in this way: `ActiveRecord::Schema.define(version: 20200610100708) do)`  and he added     `t.boolean  "adult_content",    default: false`

    
      create_table "books", force: :cascade do |t|
        t.string   "title",                            null: false
        t.string   "slug",                             null: false
        t.text     "description",                      null: false
        t.boolean  "adult_content",    default: false
      end

He edited also two tiles in spec folder and created a new file in db/migration titled `20200610100708_add_adult_content_to_books.rb`

**Now my question is: can I create a new boolean areas in the DB making edits like this using only notepad++?**
## [5][How do I query for the most recently created records across multiple models?](https://www.reddit.com/r/rails/comments/hcu0i9/how_do_i_query_for_the_most_recently_created/)
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
## [6][Modify Instance variable in a haml?](https://www.reddit.com/r/rails/comments/hcohd4/modify_instance_variable_in_a_haml/)
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
## [7][How do you structure your models for a multi-lingual app?](https://www.reddit.com/r/rails/comments/hcdh15/how_do_you_structure_your_models_for_a/)
- url: https://www.reddit.com/r/rails/comments/hcdh15/how_do_you_structure_your_models_for_a/
---
I've been tasked with creating an e-commerce website, and I'm planning to build it in Rails. However, the client wants the content to be available in both English and Thai.

I'm a Rails newbie (and have also used Laravel and Django), so I'd like some advice on how to approach this. Would the best way be to have a `Product` model that contains the price, quantity, etc, and has a one-to-many relationship with the `ProductInfo` class, which contains the title and description. So `new_product` could have two `ProductInfo` objects, one with an English description and one with a Thai description. I'm guessing the same structure could be used for a `Category` model too.

What would be the practical way of structuring the data entry form for the `Product` class? Separate form for each language, or one form with the fields for both languages?
## [8][Different data presentation without querying database again](https://www.reddit.com/r/rails/comments/hcb7ah/different_data_presentation_without_querying/)
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
## [9][Would you build your app in Rails in 2020?](https://www.reddit.com/r/rails/comments/hbpmc1/would_you_build_your_app_in_rails_in_2020/)
- url: https://www.reddit.com/r/rails/comments/hbpmc1/would_you_build_your_app_in_rails_in_2020/
---
I have a couple of (what I think are) good ideas floating around in my head, and trying to decide what to build them in.

I've dabbled in Rails, as well as full stack JS, but never built anything **serious** with either.

A few years ago I was really looking into becoming a professional Web Developer. I was working tech support for a web design company but had hit a dead end with my career there. So I was doing the typical self studying thing, Coursera, Codecademy, Hartl's tutorial, etc, getting ready to go down that route, before I kind of fell into an Application Support engineer position with a startup my friend was working in, and I have been doing that ever since. Now I support the same application for a self driving car startup making well over 6 figures.

Now that I am in a pretty good place financially and career wise, and I am a little older and have a little better work ethic, I want to take a serious crack at actually making something out of these ideas I have. Back when I was first thinking about this, a few years ago, I was pretty sure I was going to build them in Rails. Rails was pretty popular then, I was learning about it and found it really easy to use and I really liked how fast you could get something up. But it seems like in the few years I wasn't really paying attention, Rails has really fallen out of vogue, as it were. It doesn't scale. People are listing Ruby as their most hated language.  It's really hard to make anything other than an old fashioned Monolith style app with it. Etc. etc.

If you were starting a brand new project in 2020 one that you were hoping to actually monetize and go commercial with, would still doing it in Rails be a decision that you'd regret later, or is Rails fine, just everyone likes to crap on it because it's not the new hotness anymore?
## [10][Trying to make regex combo for password validation actually work](https://www.reddit.com/r/rails/comments/hc2t0e/trying_to_make_regex_combo_for_password/)
- url: https://www.reddit.com/r/rails/comments/hc2t0e/trying_to_make_regex_combo_for_password/
---
Hello,

I am building a simple authentication app and for password and username validation I need correct regex. I have tried quite a few combinations but none of them seem to actually work(either false flag or error screen)

    validates :username, presence: true, uniqueness: true, format: { with: /\A[a-zA-Z]|[0-9]|[a-zA-Z0-9]\z/ }
    validates :password, presence: true, format: { with: /\A(?=.*\d)(?=.*[A-Z])(?=.*[^a-zA-Z0-9]).{8,15}\z/ }

For username, only alphanumerical strings permitted 

For password, atleast one upper case and atleast one digit
## [11][Need help understanding module inclusion](https://www.reddit.com/r/rails/comments/hbylab/need_help_understanding_module_inclusion/)
- url: https://www.reddit.com/r/rails/comments/hbylab/need_help_understanding_module_inclusion/
---
I'm trying to understand why I'm not getting the intended results:

I have two models, CustomerAddress and SupplierAddress that each include the module AddressValidation.

In the AddressValidation module, I am trying to determine if I am dealing with a CustomerAddress or a SupplierAddress.

&amp;nbsp;&amp;nbsp;

So in AddressValidation I wanted to have:

**def self.included(base)**

&amp;nbsp;  if base.is_a?(CustomerAddress)

 &amp;nbsp; &amp;nbsp;    [run code 1]

&amp;nbsp;   else

&amp;nbsp; &amp;nbsp;     [run code 2]

&amp;nbsp;   end

**end**

&amp;nbsp;&amp;nbsp;

The problem is, that is_a?(CustomerAddress) always returns false regardless which model is including the module. It also always returns false if I use instance_of?(CustomerAddress)

&amp;nbsp;&amp;nbsp;

So what I did instead is:

**def self.included(base)**

&amp;nbsp;  if base.name == "CustomerAddress"

&amp;nbsp;&amp;nbsp;     [run_code_1]

&amp;nbsp;  else

&amp;nbsp;&amp;nbsp;     [run_code_2]

&amp;nbsp;  end

**end**

&amp;nbsp;&amp;nbsp;

This method works as intended.

So my question is why does calling .name work as intended but not .is_a? or .instance_of?


**EDIT: Formatting**

**EDIT 2** 

Here is the detailed answer to the question, for future reference.

My class CustomerAddress, as far as Ruby is concerned, is actually an instance of the class **Class** and has an attributed called "name" that has the value "CustomerAddress".

Thus base.is_a?(Class) will return true, but nothing else will.

To ensure that it is the right class, I have to either compare its name attribute, or us the equality operator as mentioned below by /u/xire28:

if base == CustomerAddress
## [12][GET http://localhost:3030/users/user_id/file_name net::ERR_ABORTED 404 (Not Found)](https://www.reddit.com/r/rails/comments/hc2rxx/get_httplocalhost3030usersuser_idfile_name_neterr/)
- url: https://www.reddit.com/r/rails/comments/hc2rxx/get_httplocalhost3030usersuser_idfile_name_neterr/
---
I'm getting this error in my console and in the terminal ActionController::RoutingError (No route matches \[GET\] "/users/user\_id/file\_name"). Everything loads correctly, except when I go from the invoices page to the dashboard page, the dashboard page loads twice. Any suggestions is greatly appreciated
