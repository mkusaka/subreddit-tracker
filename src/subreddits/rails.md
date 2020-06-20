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
## [3][How do you structure your models for a multi-lingual app?](https://www.reddit.com/r/rails/comments/hcdh15/how_do_you_structure_your_models_for_a/)
- url: https://www.reddit.com/r/rails/comments/hcdh15/how_do_you_structure_your_models_for_a/
---
I've been tasked with creating an e-commerce website, and I'm planning to build it in Rails. However, the client wants the content to be available in both English and Thai.

I'm a Rails newbie (and have also used Laravel and Django), so I'd like some advice on how to approach this. Would the best way be to have a `Product` model that contains the price, quantity, etc, and has a one-to-many relationship with the `ProductInfo` class, which contains the title and description. So `new_product` could have two `ProductInfo` objects, one with an English description and one with a Thai description. I'm guessing the same structure could be used for a `Category` model too.

What would be the practical way of structuring the data entry form for the `Product` class? Separate form for each language, or one form with the fields for both languages?
## [4][Different data presentation without querying database again](https://www.reddit.com/r/rails/comments/hcb7ah/different_data_presentation_without_querying/)
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
## [5][Trying to make regex combo for password validation actually work](https://www.reddit.com/r/rails/comments/hc2t0e/trying_to_make_regex_combo_for_password/)
- url: https://www.reddit.com/r/rails/comments/hc2t0e/trying_to_make_regex_combo_for_password/
---
Hello,

I am building a simple authentication app and for password and username validation I need correct regex. I have tried quite a few combinations but none of them seem to actually work(either false flag or error screen)

    validates :username, presence: true, uniqueness: true, format: { with: /\A[a-zA-Z]|[0-9]|[a-zA-Z0-9]\z/ }
    validates :password, presence: true, format: { with: /\A(?=.*\d)(?=.*[A-Z])(?=.*[^a-zA-Z0-9]).{8,15}\z/ }

For username, only alphanumerical strings permitted 

For password, atleast one upper case and atleast one digit
## [6][Would you build your app in Rails in 2020?](https://www.reddit.com/r/rails/comments/hbpmc1/would_you_build_your_app_in_rails_in_2020/)
- url: https://www.reddit.com/r/rails/comments/hbpmc1/would_you_build_your_app_in_rails_in_2020/
---
I have a couple of (what I think are) good ideas floating around in my head, and trying to decide what to build them in.

I've dabbled in Rails, as well as full stack JS, but never built anything **serious** with either.

A few years ago I was really looking into becoming a professional Web Developer. I was working tech support for a web design company but had hit a dead end with my career there. So I was doing the typical self studying thing, Coursera, Codecademy, Hartl's tutorial, etc, getting ready to go down that route, before I kind of fell into an Application Support engineer position with a startup my friend was working in, and I have been doing that ever since. Now I support the same application for a self driving car startup making well over 6 figures.

Now that I am in a pretty good place financially and career wise, and I am a little older and have a little better work ethic, I want to take a serious crack at actually making something out of these ideas I have. Back when I was first thinking about this, a few years ago, I was pretty sure I was going to build them in Rails. Rails was pretty popular then, I was learning about it and found it really easy to use and I really liked how fast you could get something up. But it seems like in the few years I wasn't really paying attention, Rails has really fallen out of vogue, as it were. It doesn't scale. People are listing Ruby as their most hated language.  It's really hard to make anything other than an old fashioned Monolith style app with it. Etc. etc.

If you were starting a brand new project in 2020 one that you were hoping to actually monetize and go commercial with, would still doing it in Rails be a decision that you'd regret later, or is Rails fine, just everyone likes to crap on it because it's not the new hotness anymore?
## [7][Need help understanding module inclusion](https://www.reddit.com/r/rails/comments/hbylab/need_help_understanding_module_inclusion/)
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
## [8][GET http://localhost:3030/users/user_id/file_name net::ERR_ABORTED 404 (Not Found)](https://www.reddit.com/r/rails/comments/hc2rxx/get_httplocalhost3030usersuser_idfile_name_neterr/)
- url: https://www.reddit.com/r/rails/comments/hc2rxx/get_httplocalhost3030usersuser_idfile_name_neterr/
---
I'm getting this error in my console and in the terminal ActionController::RoutingError (No route matches \[GET\] "/users/user\_id/file\_name"). Everything loads correctly, except when I go from the invoices page to the dashboard page, the dashboard page loads twice. Any suggestions is greatly appreciated
## [9][Help me to figure out ActiveRecord logic](https://www.reddit.com/r/rails/comments/hbz770/help_me_to_figure_out_activerecord_logic/)
- url: https://www.reddit.com/r/rails/comments/hbz770/help_me_to_figure_out_activerecord_logic/
---
So here's what I want:

* Seller that has many Testimonials;
* Author that has many Testimonials;
* ConfirmedProfile of The User that gets access to all testimonials as Seller and Author, the confirmation would be requested by user and confirmed by Admin.

I am not sure which association types to choose, I believe the polymorphic? and probably I need 'through' association for confirmation part too? Any thoughts?
## [10][Gems to build order-tracking software?](https://www.reddit.com/r/rails/comments/hbqtiz/gems_to_build_ordertracking_software/)
- url: https://www.reddit.com/r/rails/comments/hbqtiz/gems_to_build_ordertracking_software/
---
Curious what gems you've used to build order-tracking software whether it be to integrate with shipping like UPS/FedEx etc or to track local deliveries.
## [11][Rails 6 multiple databases support in Rails Event Store](https://www.reddit.com/r/rails/comments/hbaxq8/rails_6_multiple_databases_support_in_rails_event/)
- url: https://www.reddit.com/r/rails/comments/hbaxq8/rails_6_multiple_databases_support_in_rails_event/
---
Rails 6 released in August 2019 has brought us several new features. One of the notable changes is support for multiple databases. All details have been described in Rails guides and I’ve read already several blog posts describing how to do it. But how to use this feature to allow Rails Event Store data to be stored in a separate database? Check my [blog post where I document my experiments with it](https://blog.arkency.com/rails-multiple-databases-support-in-rails-event-store/).
## [12][Docker and rails issue](https://www.reddit.com/r/rails/comments/hbney4/docker_and_rails_issue/)
- url: https://www.reddit.com/r/rails/comments/hbney4/docker_and_rails_issue/
---
Hello Folks, i was trying to follow docker guide to setup a new rails app with, but it in rails 5 and i wanna rails 6, is there any guide you can share with me guys.

the docker example with rails 5 : [https://docs.docker.com/compose/rails](https://docs.docker.com/compose/rails/)

**note i don't wanna install rails in my machine or any other gem just like docker docs did!!** 
