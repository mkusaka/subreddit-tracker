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
## [2][Image tag if logic?](https://www.reddit.com/r/rails/comments/ios0qj/image_tag_if_logic/)
- url: https://www.reddit.com/r/rails/comments/ios0qj/image_tag_if_logic/
---
Would it be possible to check if the image is square and set a particular resize option and if not square set it to other option. Something like this:

    image_tag(logo_path, if width==height ? size: 167 : width: 167)
## [3][Is there a way to use trailblazer reform with vuejs forms](https://www.reddit.com/r/rails/comments/ionf5y/is_there_a_way_to_use_trailblazer_reform_with/)
- url: https://www.reddit.com/r/rails/comments/ionf5y/is_there_a_way_to_use_trailblazer_reform_with/
---
Hey, we are transforming some of our views with a lot of unmaintainable coffeescript to vuejs components. For form validation we are using dry/reform including the trailblazer stack. Is there a handy way to migrate these forms without to doublicate our validations to the front end with proper user error messaging?
## [4][A few questions about Rails.ajax](https://www.reddit.com/r/rails/comments/ioajn2/a_few_questions_about_railsajax/)
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
## [5][When to make external API connection](https://www.reddit.com/r/rails/comments/iobg0s/when_to_make_external_api_connection/)
- url: https://www.reddit.com/r/rails/comments/iobg0s/when_to_make_external_api_connection/
---
I want to make a site that uses the k8s-client gem to interact with kubernetes clusters. This is my first time writing something that interacts with an external API. I have written a job to make the connection, but I'm wondering where I should call it. Should I spawn a new connection each time I run a job that makes an API call? What's the best practice for doing this?
## [6][Using Google places in Rails Api](https://www.reddit.com/r/rails/comments/io452h/using_google_places_in_rails_api/)
- url: https://www.reddit.com/r/rails/comments/io452h/using_google_places_in_rails_api/
---
hi guys!

I am making Rails API which will take user location in params and return restaurants around the radius of 500 meters in form of an array, actually, I don't know the steps that how can I return those restaurants by using google places. Any help will be highly appreciated.
## [7][Ruby on Rails digest: 26 most popular repositories in July and August 2020](https://www.reddit.com/r/rails/comments/inkoxw/ruby_on_rails_digest_26_most_popular_repositories/)
- url: https://www.reddit.com/r/rails/comments/inkoxw/ruby_on_rails_digest_26_most_popular_repositories/
---
This is the second edition of Ruby/Rails digest. [This post](https://medium.com/@Iren.Korkishko/ruby-on-rails-digest-26-most-popular-repositories-in-july-and-august-2020-70ae593a5fe1?source=friends_link&amp;sk=477b6f0fb300b07f7fe068d12831c756) welcomes the most popular, most interesting, and useful Ruby on Rails repositories on GitHub in July and August!
## [8][Help with asset pipeline](https://www.reddit.com/r/rails/comments/ingy5q/help_with_asset_pipeline/)
- url: https://www.reddit.com/r/rails/comments/ingy5q/help_with_asset_pipeline/
---
Hey guys! I’m trying to figure out how to properly configure my asset pipeline. I’m having some trouble with this. In my production.rb file, I have config.assets.compile set to false and additionally, in my assets/config/manifest.js file, I have the following

//= link_tree ../images
//= link_directory ../stylesheets .css

However once I deploy the site, I get an error saying: “the asset &lt;image.jpg&gt; is not present in the asset pipeline” 

I’m accessing the images using: &lt;%= image_tag “image.jpg” %&gt;

It all works when I have config.assets.compile set to true but I know this is inefficient.

I would appreciate any help!
## [9][I need an experienced rails developer](https://www.reddit.com/r/rails/comments/inqvjj/i_need_an_experienced_rails_developer/)
- url: https://www.reddit.com/r/rails/comments/inqvjj/i_need_an_experienced_rails_developer/
---
Hello everyone! I am working on a start up for restaurant reservations web app and really need a rails developer who is able to work and coordinate with a react developer, who is actually working for this project. There are only some tasks that need to be done since most of the work is already done! PM if any of you is interested. :)
## [10][multiple scopes and one def in controller](https://www.reddit.com/r/rails/comments/in3fbu/multiple_scopes_and_one_def_in_controller/)
- url: https://www.reddit.com/r/rails/comments/in3fbu/multiple_scopes_and_one_def_in_controller/
---
Hi guys, one question

if in home (view) I have this:

    
    &lt;div&gt;Today Offers&lt;/div&gt;
    &lt;%= render 'books/books', books: @query_of_the_day, skip_pages: true %&gt;

&amp;#x200B;

and in home\_controller.rb I have this

      def query_of_the_day
        day_of_the_week = Time.now.wday
        @query_of_the_day = if day_of_the_week == 0
                  Book.downloaded_today.includes(:screen, :online_links).limit(6).decorate
                ifelse day_of_the_week == 1
                  Book.rated_yesterday.includes(:screen, :online_links).limit(6).decorate
                ifelse day_of_the_week == 2
                  Book.over_50_readers.includes(:screen, :online_links).limit(6).decorate
                ifelse day_of_the_week == 3
                  Book.uploaded_yesterday.includes(:screen, :online_links).limit(6).decorate
                ifelse day_of_the_week == 4
                  Book.over_100_votes.includes(:screen, :online_links).limit(6).decorate
                else
                  Book.random.includes(:screen, :online_links).limit(6).decorate
                end

Where `downloaded_today`, `rated_yesterday`, `over_50_readers`, etc., are linked with a respective scope in scopes area of model/book.rb

&amp;#x200B;

My question is: Will it load ALL the `day_of_the_week` CASES (or all the  UNUSED scopes) when an user will visit the homepage?

...Because, in that case, my homepage will be very slow.
## [11][how do you deal with user timezones?](https://www.reddit.com/r/rails/comments/imym4y/how_do_you_deal_with_user_timezones/)
- url: https://www.reddit.com/r/rails/comments/imym4y/how_do_you_deal_with_user_timezones/
---
in a previous project i was saving the user timezone when signing up with an hidden field filled up with javascript to get the user's timezone and then used that timezone in views and mailers

but in my current project users can sign up with facebook so i can't add an hidden field

do i need to store the user timezone?

how would you localize datetimes in mailers without the user timezone stored?

and more generally how do you deal with user timezones?
