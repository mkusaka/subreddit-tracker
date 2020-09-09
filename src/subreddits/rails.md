# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/igyvm1/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/igyvm1/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/ipfxw8/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/ipfxw8/personal_projects_show_off_your_own_project_andor/
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
## [3][5 Star-Rating in Rails API](https://www.reddit.com/r/rails/comments/ipbq5c/5_starrating_in_rails_api/)
- url: https://www.reddit.com/r/rails/comments/ipbq5c/5_starrating_in_rails_api/
---
Hi guys!

I'm learning rails and my work is from the backend side and I'm going to make a restaurant rating system in which a user can rate a restaurant he checkin before. I'm getting user\_id &amp; rating e.g (4,5,3) from front-end now I don't know how to get the overall restaurant rating. I have an idea below

s = sum of rating on a particular restaurant  t = total no of rating on that restaurant   Restaurant overall rating = s/t 

&amp;#x200B;

Here I'm getting these values from my Rating model where I store restaurant\_ID, User\_id, user\_rating, Overall\_rating.

I don't know here, I'm going on the right path or not but I need guidness from seniors.
## [4][Multi-step form structure](https://www.reddit.com/r/rails/comments/ipbf32/multistep_form_structure/)
- url: https://www.reddit.com/r/rails/comments/ipbf32/multistep_form_structure/
---
Hey Redditors,

I'm working on a community website where each member can have their own company page. The company page has the ability to have a regular "contact us form" but it's not very flexible. Many of our current users have asked for the ability to have a multi-step form and custom forms for their industry.

Let's say you are an electrician and you would like to get more information out of a potential customer before you call them. So you would like to add fields like "*What's the purpose for the new wiring?*" or "*What's the approximate square footage of the home?*".

**The functionality I'm looking to implement**

* Unique forms for a specific industry (electrician, roofers, plumbers) with text fields, number fields, radio buttons, checkboxes, and photo uploads
* The ability to make changes to the form on a specific company page
* Admin users should be able to create forms for each industry and modifying them
* "Clone" a premade form for an industry then have the ability to modifying it

&amp;#x200B;

&gt;***Porch, Thumbtack and many others have this kind of functionality***

[https:\/\/porch.com\/checkout\/start?appId=homePage&amp;companyTypeId=1071&amp;context=home-page&amp;conversionId=ae29f83c-5f5e-4c7f-b3e5-5707b2cf49c5&amp;pageName=HomePage&amp;postalCode=10001&amp;r=https&amp;#37;3A&amp;#37;2F&amp;#37;2Fporch.com&amp;#37;2F&amp;serviceTypeId=](https://preview.redd.it/3j80oavon2m51.png?width=2572&amp;format=png&amp;auto=webp&amp;s=02c4f116d802fbc05f02b528bfee405ae4a35eb5)

&amp;#x200B;

**What I'm asking the rails community is the following:**

* How would you structure this kind of functionality?
   * Database tables
   * Schema
   * Relationships
* Would you make API calls for each step of the multi-step form or just use localStorage until submission?
* What security aspects are easy to forget about?
* Are there any gems out there solving this?

&amp;#x200B;

Thanks, Daniel!
## [5][validate x number of model boolean attribute](https://www.reddit.com/r/rails/comments/ip6xqo/validate_x_number_of_model_boolean_attribute/)
- url: https://www.reddit.com/r/rails/comments/ip6xqo/validate_x_number_of_model_boolean_attribute/
---
Young on my Rails journey, but loving it so far. I've got a nested Resource as something like `Parent` =&gt; `Child`. The `Child` model has an `is_active` boolean attribute on it, which I can toggle on a form (yay!)

What I'm trying to get working next, is to allow any number of Child records for a Parent, but ONLY allow 4 of them where `is_active` is true. I'm trying to write this as a validate callback on the Child model...am I on the right track?
## [6][Image tag if logic?](https://www.reddit.com/r/rails/comments/ios0qj/image_tag_if_logic/)
- url: https://www.reddit.com/r/rails/comments/ios0qj/image_tag_if_logic/
---
Would it be possible to check if the image is square and set a particular resize option and if not square set it to other option. Something like this:

    image_tag(logo_path, if width==height ? size: 167 : width: 167)
## [7][Is there a way to use trailblazer reform with vuejs forms](https://www.reddit.com/r/rails/comments/ionf5y/is_there_a_way_to_use_trailblazer_reform_with/)
- url: https://www.reddit.com/r/rails/comments/ionf5y/is_there_a_way_to_use_trailblazer_reform_with/
---
Hey, we are transforming some of our views with a lot of unmaintainable coffeescript to vuejs components. For form validation we are using dry/reform including the trailblazer stack. Is there a handy way to migrate these forms without to doublicate our validations to the front end with proper user error messaging?
## [8][A few questions about Rails.ajax](https://www.reddit.com/r/rails/comments/ioajn2/a_few_questions_about_railsajax/)
- url: https://www.reddit.com/r/rails/comments/ioajn2/a_few_questions_about_railsajax/
---
A few days back I asked how to weave a form through a table:  [Post](https://www.reddit.com/r/rails/comments/ilglc6/how_can_you_distribute_input_elements_outside_a/)

I'v mostly gotten overmotion’s suggestion of submitting formless inputs using AJAX to function, but I’m a little fuzzy on how Rails.ajax actually works.  I couldn’t find anything about it in the rails docs and all the tutorials either use coffeescript, assume you're familiar with the jquery method it’s derived from, or both.

Here’s my solution:

    Rails.ajax({
        type: "PATCH",
        url: "nurses/&lt;%= @nurse.id %&gt;",
        data: `first=${first}&amp;last=${last}&amp;start_date=${startDate}&amp;can_charge=${can_charge}`,
        success: function(data) {
        }
      });

&amp;#x200B;

and my questions:

\-- Is my entire update.js.erb file being wrapped in a callback function and executed upon a successful ajax request?  Is it being passed through the ‘data’ parameter in the above success method?  And if so, why doesn’t it need to be called in the body?

\-- Is there a failed-validations version of update.js.erb I can use to send feedback to the user?  Where would that be stored and what would its name be?

If you guys have any good resources on this, I’d be happy to read them.  Thanks!

This is the full edit.js.erb file for anyone curious:

    function ajaxSubmit() {
      const first = document.querySelector("#first").value;
      const last = document.querySelector("#last").value;
      const startDate = document.querySelector("#start_date").value;
      const canCharge = document.querySelector("#can_charge").value;
    
      Rails.ajax({
        type: "PATCH",
        url: "nurses/&lt;%= @nurse.id %&gt;",
        data: `first=${first}&amp;last=${last}&amp;start_date=${startDate}&amp;can_charge=${can_charge}`,
        success: function(data) {
        }
      });
    };
    
    function switchRow() {
      let clickedTR = document.querySelector(".nurse&lt;%= @nurse.id %&gt;");
      clickedTR.innerHTML = "&lt;%= escape_javascript(render partial: 'nurses/inline_form', locals: { nurse: @nurse } ) %&gt;";
      let update_button = document.querySelector("#update-nurse");
    
      update_button.addEventListener("click", ajaxSubmit, true);
    };
    
    switchRow(); 
## [9][When to make external API connection](https://www.reddit.com/r/rails/comments/iobg0s/when_to_make_external_api_connection/)
- url: https://www.reddit.com/r/rails/comments/iobg0s/when_to_make_external_api_connection/
---
I want to make a site that uses the k8s-client gem to interact with kubernetes clusters. This is my first time writing something that interacts with an external API. I have written a job to make the connection, but I'm wondering where I should call it. Should I spawn a new connection each time I run a job that makes an API call? What's the best practice for doing this?
## [10][Using Google places in Rails Api](https://www.reddit.com/r/rails/comments/io452h/using_google_places_in_rails_api/)
- url: https://www.reddit.com/r/rails/comments/io452h/using_google_places_in_rails_api/
---
hi guys!

I am making Rails API which will take user location in params and return restaurants around the radius of 500 meters in form of an array, actually, I don't know the steps that how can I return those restaurants by using google places. Any help will be highly appreciated.
## [11][Ruby on Rails digest: 26 most popular repositories in July and August 2020](https://www.reddit.com/r/rails/comments/inkoxw/ruby_on_rails_digest_26_most_popular_repositories/)
- url: https://www.reddit.com/r/rails/comments/inkoxw/ruby_on_rails_digest_26_most_popular_repositories/
---
This is the second edition of Ruby/Rails digest. [This post](https://medium.com/@Iren.Korkishko/ruby-on-rails-digest-26-most-popular-repositories-in-july-and-august-2020-70ae593a5fe1?source=friends_link&amp;sk=477b6f0fb300b07f7fe068d12831c756) welcomes the most popular, most interesting, and useful Ruby on Rails repositories on GitHub in July and August!
## [12][Help with asset pipeline](https://www.reddit.com/r/rails/comments/ingy5q/help_with_asset_pipeline/)
- url: https://www.reddit.com/r/rails/comments/ingy5q/help_with_asset_pipeline/
---
Hey guys! I’m trying to figure out how to properly configure my asset pipeline. I’m having some trouble with this. In my production.rb file, I have config.assets.compile set to false and additionally, in my assets/config/manifest.js file, I have the following

//= link_tree ../images
//= link_directory ../stylesheets .css

However once I deploy the site, I get an error saying: “the asset &lt;image.jpg&gt; is not present in the asset pipeline” 

I’m accessing the images using: &lt;%= image_tag “image.jpg” %&gt;

It all works when I have config.assets.compile set to true but I know this is inefficient.

I would appreciate any help!
