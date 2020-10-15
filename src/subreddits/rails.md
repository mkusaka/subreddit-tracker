# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/j6qvuh/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/j6qvuh/personal_projects_show_off_your_own_project_andor/
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
## [2][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/jbmadq/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/jbmadq/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [3][validating if an email is REAL](https://www.reddit.com/r/rails/comments/jbbtw7/validating_if_an_email_is_real/)
- url: https://www.reddit.com/r/rails/comments/jbbtw7/validating_if_an_email_is_real/
---
Validating an email by REGEX is usually not enough.

If you use `validates :email, format: { with: URI::MailTo::EMAIL_REGEXP }, allow_blank: true` this will not allow to submit something like "**arebefrvevervv**" in the `email` field, but a user will still be able to submit an email address like "**vedsvear@vfdsvdf.c**" or "**test@test.test**" - we don't want this happening.

So, we need a solution to check if the "**@domain**" or "**email@domain**" actually exists. Here's a [nice article](https://rubygarage.org/blog/how-to-validate-emails) that I found on this topic.

There seem to be a few gems that help to add this additional validation layer:

* [https://github.com/rubygarage/truemail](https://github.com/rubygarage/truemail)
* [https://github.com/afair/email\_address](https://github.com/afair/email_address)
* [https://github.com/micke/valid\_email2](https://github.com/micke/valid_email2)

Do you have experience using any of these gems?  
Which one should one go for?
## [4][convert an array of nested arrays into an (array of arrays )](https://www.reddit.com/r/rails/comments/jblhwm/convert_an_array_of_nested_arrays_into_an_array/)
- url: https://www.reddit.com/r/rails/comments/jblhwm/convert_an_array_of_nested_arrays_into_an_array/
---
im learning rails and i have output like this  e.g 

`rows = [[[a,b,c]],[[d,e,f]],[[g,h,I]]]`  

Now I want to convert this to this one  `rows = [[a,b,c],[d,e,f],[g,h,I]]`  

can anyone please help me to remove that extrabrackets ?
## [5][how is error_messages_for populated for an object](https://www.reddit.com/r/rails/comments/jbjpc9/how_is_error_messages_for_populated_for_an_object/)
- url: https://www.reddit.com/r/rails/comments/jbjpc9/how_is_error_messages_for_populated_for_an_object/
---
For a form, I have the error messages to be displayed like this 

    error_messages_for :model

I'm not able to figure out how a particular error message is populated. 

Like "Name is invalid"

But the thing is there is no column called name in the model and I do not see things being explicitly added to the model. 

like model.errors.add(:base, 'Name is invalid')

so I'm not sure if for some validation this record, these messages get automatically added?
## [6][Creating users as an Admin with devise](https://www.reddit.com/r/rails/comments/jb5vdp/creating_users_as_an_admin_with_devise/)
- url: https://www.reddit.com/r/rails/comments/jb5vdp/creating_users_as_an_admin_with_devise/
---
I want to be able to register Users as an **admin** but I'm not able because when I try to register a new user I get the 'You are already signed in' alert. Is there any way I can go around that and just be able to register users. I tried a few things that I saw on stackoverflow, but none worked. 

Please help
## [7][Rails EasyAutocomplete gem not shows suggestions](https://www.reddit.com/r/rails/comments/jb1b1s/rails_easyautocomplete_gem_not_shows_suggestions/)
- url: https://www.reddit.com/r/rails/comments/jb1b1s/rails_easyautocomplete_gem_not_shows_suggestions/
---
I was following this tutorial  [click][1] and I wrote the following code:

HTML:

    &lt;input type="text" id="prova" placeholder="Enter name1..." data-behavior="autocomplete" /&gt;

JS:

    $("#prova").keyup(function() {
    
      $.ajax({
        type: "GET",
        url: "https://photon.komoot.de/api/?q=" + $("#prova").val(),
    
        success: function(results) {
    
          var aList = results.features;
          var aOptions = [];
          let htmlVal = '';
          for (i = 0; i &lt; aList.length; i++) {
            optLabel = aList[i].properties.name;
            aOptions.push(optLabel);
          }
    
    
          var options = {
            data: aOptions
          };
    
          $('*[data-behavior="autocomplete"]').easyAutocomplete(options);
    
        }
      });
    
    });

as you can see through the ajax call I connect to an external service that returns addresses. The problem is that they don't appear. Where am I wrong?
(PS: The api rest works, I'm able to print in console.log aOptions array)
  


  [1]: https://joelc.io/dynamic-autocomplete-rails-6
## [8][Watch me build a new Ruby/Rails community in public 📹](https://www.reddit.com/r/rails/comments/jaoh8a/watch_me_build_a_new_rubyrails_community_in_public/)
- url: https://www.reddit.com/r/rails/comments/jaoh8a/watch_me_build_a_new_rubyrails_community_in_public/
---
Hey all,

I've been wanting to see a place/resource build up more hype about Rails and Ruby again like the good ole' days so I'm building one in public.

Some background about what exactly I plan on building is here: [https://web-crunch.com/posts/lets-build-for-ruby-and-rails-developers](https://web-crunch.com/posts/lets-build-for-ruby-and-rails-developers)

I plan to commit to a few hours of recording my progress a week for an initial MVP of the job board portion of the app. Once that's complete I'll do a tentative soft launch and continue adding more of the community features. 

#### Primary goals:

- Build some buzz around Ruby / Rails.
- Bring more Ruby and Rails developers together.
- Have a centralized place to find Ruby and Rails jobs.
- Showcase profiles developers can use to get hired.
- Have a community forum for all things Ruby and Rails + more web/programming topics.
- Maybe make some side $$

The ongoing YouTube playlist is here: [https://www.youtube.com/playlist?list=PL01nNIgQ4uxOfChhPf3jxq8H6fUWnCeLC](https://www.youtube.com/playlist?list=PL01nNIgQ4uxOfChhPf3jxq8H6fUWnCeLC)

The ongoing collection of articles/videos is here:
[https://web-crunch.com/collections/lets-build-for-ruby-and-rails-developers](https://web-crunch.com/collections/lets-build-for-ruby-and-rails-developers)

The buzz around the Laravel community is inspiring and is a big driver for me to take on this project. 

It's worth noting that I'm not a super-advanced Rails developer like I'm sure many of you are but I believe I know enough to build something such as this (no matter how scrappy to start). Feedback and pointers are welcomed as I'm also using this as a learning experience and potential to earn some side income. 

**TL;DR;** I'm building a community and job board for Ruby/Rails developers in public. Maybe follow along if you're interested?
## [9][Minitest vs RSpec for Rails Development](https://www.reddit.com/r/rails/comments/jay6xk/minitest_vs_rspec_for_rails_development/)
- url: https://www.reddit.com/r/rails/comments/jay6xk/minitest_vs_rspec_for_rails_development/
---
Hi guys. Recently completed a bootcamp, covering a lot of Ruby on Rails and using RSpec as the testing framework from this. I also see from looking online a lot of people set up RSpec as their framework for Rails `rails new -T` without minitest (the default)

What are peoples opinions on using Minitest, as I have been approached to interview for a role which uses it and from looking online it seems inferior in many ways. I am trying to get it set up on my rails app (just to play with it) and cant seem to find a good coverage solution (like simplecov) as I found it to be a problem setting this up with Minitest (with the little documentation on it also)

So yeah just a question, what do you guys use to test Rails apps and if its minitest what resources and coverage do you use!  


EDIT: beautiful simplecov (and simplecov-console), minitest works fine. Thanks for all the help!

https://preview.redd.it/r0tl3klhb8t51.png?width=603&amp;format=png&amp;auto=webp&amp;s=c1745b697775e0dfbace8d3acb3fa6b329aaa117

  

## [10][ActiveStorage query optionals](https://www.reddit.com/r/rails/comments/jazt1a/activestorage_query_optionals/)
- url: https://www.reddit.com/r/rails/comments/jazt1a/activestorage_query_optionals/
---
Hi guys, I was wondering if AS has building queries for doing joins. That way, I would avoiding the n+1 query problem. E.g: In the example bellow, I fetched all the attachments for 'files\_attached' (query 1) and later I display each file's filename (the 3 blob queries for 3 files). 

    ##query 1
    ActiveStorage::Attachment Load (0.3ms)  SELECT "active_storage_attachments".* FROM "active_storage_attachments" WHERE "active_storage_attachments"."record_id" = $1 AND "active_storage_attachments"."record_type" = $2 AND "active_storage_attachments"."name" = $3
    
    ## displaying filenames for each file fetched
    ##query 2
    ActiveStorage::Blob Load (0.2ms)  SELECT "active_storage_blobs".* FROM "active_storage_blobs" WHERE "active_storage_blobs"."id" = $1 LIMIT $2
    
    ##query 3
    ActiveStorage::Blob Load (0.2ms)  SELECT "active_storage_blobs".* FROM "active_storage_blobs" WHERE "active_storage_blobs"."id" = $1 LIMIT $2
    
    ##query 4
    ActiveStorage::Blob Load (0.2ms)  SELECT "active_storage_blobs".* FROM "active_storage_blobs" WHERE "active_storage_blobs"."id" = $1 LIMIT $2

I'm aware that I could make one single query using raw SQL, but I would like to avoid it if possible.
## [11][Webpacker Issues](https://www.reddit.com/r/rails/comments/jaxxx0/webpacker_issues/)
- url: https://www.reddit.com/r/rails/comments/jaxxx0/webpacker_issues/
---
Hi, just a stressed out dev trying to get some JS to work through Webpacker. The error i'm getting is that a method is undefined. I'm trying to install/compile a third party js libary (masonry) into Webpacker and i'm having an absolute nightmare with it. I've installed the package I want through Yarn and I can see that it's the correct one as when I do 'yarn list', I can see it. Great, so we have it installed and can see in the package.json that it's installed but how do I know what I need to require/import? Surely it should just be a matter of installing it through a package managers (bower, npm or yarn) and then just  requiring the file in the app/javascript/packs/application.js using the name of the file in the package.json?

Is there something i'm missing? I've looked everywhere on how to do this but every article I go to seems to say something else...

**Update**: SOLVED. Issue was I wasn't referencing the file correctly within the application.js file. It's clicked with me now and I think I now understand how it all links up together. To my understanding, whenever you 'require' a file in the application.js, it looks within the .node\_modules folder which I did not know about. I got the solution by looking at my yarn.lock file to see the name of the folder and then looked in the .node\_modules/#{directory} so when I put 'require("#{directory}"), it thought it was the directory than the file. It just confused me as some of the other requires (actiontext, actionstorage etc) didn't reference their files, it references the directory to include (Tell me if i'm wrong).
## [12][What can a Noob put in his portfolio ?](https://www.reddit.com/r/rails/comments/jagmh4/what_can_a_noob_put_in_his_portfolio/)
- url: https://www.reddit.com/r/rails/comments/jagmh4/what_can_a_noob_put_in_his_portfolio/
---
As the title said, i'm completely new to ruby/rails what i should put in my portfolio ? any ideas ?
