# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/jfcv1r/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/jfcv1r/personal_projects_show_off_your_own_project_andor/
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
## [2][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/jsthk8/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/jsthk8/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [3][6 Things to Do When Inheriting Legacy Rails Apps](https://www.reddit.com/r/rails/comments/jtgjla/6_things_to_do_when_inheriting_legacy_rails_apps/)
- url: https://www.reddit.com/r/rails/comments/jtgjla/6_things_to_do_when_inheriting_legacy_rails_apps/
---
One of our engineers wrote a guide to help people get off on the right foot when inheriting a Rails app and [I wanted to share it here.](https://nextlinklabs.com/insights/six-tips-for-inheriting-a-legacy-rails-app) Let me know what you think.
## [4][Company Workflows](https://www.reddit.com/r/rails/comments/jsy5xv/company_workflows/)
- url: https://www.reddit.com/r/rails/comments/jsy5xv/company_workflows/
---
I have a few questions regarding workflow and i'd love to hear some feedback from you guys about it and to compare it to my current workflow.

I'm mainly curious as to what your workflow is like with other team members, mainly designers.

I understand every project is different and dependent on the client, the workflow would be different.

I'm currently in a job where it's very 'get it done and get it out the door as quick as possible', this tends to leave a lot of issues when going live for example, the code breaking because it thinks a field is there and when it does html\_safe onto it. When this happens, I just tend to fix the error and get it deployed ASAP which works.

When working with a design, I tend to get everything done, have full control of all the margins, padding, font-sizes etc so I can change it but it's always just the little things that I don't see as i've been looking at it for too long and begins to just looks normal to me (I assume most people have this issue) which is fine and simple to change.

Some of these questions might not apply to freelancers as I imagine the client already has a design and they just want some product back ASAP (or within a time frame)

1) Before the developer get's the end design should everything should be tied up, all content in and approved by the client before the developer starts work on it? 

2) Do you communicate with the designer throughout and ask them to give feedback on something you've done and make changes if needed?

3) Most projects aren't smooth sailing as you'd expect due to client changes, design changes or config changes, how do you get around these? I know the codebase and what I should change and where it is. 

4) How do you tackle feedback? For me, i'm really bad at taking criticism but that's just due to the person I am and I hate receiving negative feedback as I just want to make people happy with what I do.

I'm really just looking at ways I could potentially make my workflow a little better as i'm seeing myself doing something and then doing it again and again and again and I feel like this could be easily avoided if it all was done and correct in the first place?
## [5][Push a Dockerized Rails API on Heroku](https://www.reddit.com/r/rails/comments/jtco98/push_a_dockerized_rails_api_on_heroku/)
- url: https://www.reddit.com/r/rails/comments/jtco98/push_a_dockerized_rails_api_on_heroku/
---
hope you are fine!    
Actually, I have a dockerized rails API I'm working on.&amp; I have to push it on Heroku.  


1. I add pg gem to my production environment &amp; move MySql gem to the development environment. 

&amp;#8203;

    group :development, :test do
      # Call 'byebug' anywhere in the code to stop execution and get a   console
      gem 'byebug', platforms: [:mri, :mingw, :x64_mingw]
      # Adds support for Capybara system testing and selenium driver
      gem 'capybara', '~&gt; 2.13'
      gem 'selenium-webdriver'
      gem 'mocha'
      gem 'webmock'
      gem 'mysql2'
      
    end
    
    group :production do
      gem 'pg', '1.2.3'
    end

And after that, I run `docker-compose up --build`    &amp; after pushing my code to git I run `git push heroku &lt;branch name&gt;` everything works fine. 

but when I run `heroku run rails db:migrate`  i got an error  

    rails aborted!
    NoMethodError: Cannot load database configuration:
    undefined method `[]' for nil:NilClass

This is my first time to push a dockerized project on Heroku,  i don't know how to do this.  
can anyone please guide me?
## [6][Any good resources for beginners on how to write good tests?](https://www.reddit.com/r/rails/comments/jsmecm/any_good_resources_for_beginners_on_how_to_write/)
- url: https://www.reddit.com/r/rails/comments/jsmecm/any_good_resources_for_beginners_on_how_to_write/
---
I know there are tons of tutorials/docs out there about testing, but are there good ones that teaches good practical/pragmatic approaches on testing rails apps? Something that teaches good practice on what to test, instead of just how to test things; what are bad tests, how to avoid bugs in tests etc
## [7][Getting the response format of requests](https://www.reddit.com/r/rails/comments/jswgto/getting_the_response_format_of_requests/)
- url: https://www.reddit.com/r/rails/comments/jswgto/getting_the_response_format_of_requests/
---
The performance metrics we use - by default - keys requests like `&lt;controller&gt;#&lt;action&gt;` but we want it to also separate requests by format. Currently I overwrite the action name in an after_action callback and append the format.

    trx.set_action("#{controller.class}##{controller.action_name}.#{request.format.to_sym || :unknown}")

The problem is that request.format is undetermined for requests with `Accept: */*`. Also `request.format` only reflects the preferred response type, a json-only action might have a HTML request.format but responds with a json content type.

So all in all is there a good way to get the actual format that was rendered in a *action filter? Only thing I really found is `response.content_type` but there I would need to fish out the format from the MIME type string. I was wondering if there is a way to get Mime::Type instance of the response.

Edit: Best I could come up with so far `response.media_type&amp;.split("/")&amp;.last`
## [8][Issues with Google Maps &amp; Places API for Rails app on Mobile Browsers](https://www.reddit.com/r/rails/comments/jsvie7/issues_with_google_maps_places_api_for_rails_app/)
- url: https://www.reddit.com/r/rails/comments/jsvie7/issues_with_google_maps_places_api_for_rails_app/
---
Hi All - working on a rails app that leverages google maps/places APIs for address autocompletion. It works quite well on desktop browsers, but when I attempt to enter an address on mobile it won't work/ will clear the entire form. I've attached some snippets of the code below:

#autcomplete.js

    $(document).on('turbolinks:load', function() {
    function initializeAutocomplete(id) {     
        var element = document.getElementById(id);
        if (element) {
            var autocomplete = new google.maps.places.Autocomplete(element, { types: ['geocode'], componentRestrictions: {country: 'us'} });
            google.maps.event.addListener(autocomplete, 'place_changed', onPlaceChanged);
        }
    }

    function onPlaceChanged() {
        var place = this.getPlace();       
        for (var i in place.address_components) {
            var component = place.address_components[i];
            for (var j in component.types) { 
                var type_element = document.getElementById(component.types[j]);
                if (type_element) {
                    type_element.value = component.long_name;
                }
            }
        }
    }   
    google.maps.event.addDomListener(window, 'load', function() {
        initializeAutocomplete('autocomplete_address');   
        }); 
    }); 

and within my form, the address field is like so:

    &lt;%= f.text_field :address, class: 'form-control form-control-mobile description-mobile', placeholder: "Enter your home address", id: :autocomplete_address %&gt;&lt;br&gt;
## [9][How large is your test suite?](https://www.reddit.com/r/rails/comments/jseicl/how_large_is_your_test_suite/)
- url: https://www.reddit.com/r/rails/comments/jseicl/how_large_is_your_test_suite/
---
Just ran my test suite for the app I am working on.   


Here's are the stats:

`Finished in 303.213471s, 1.6787 runs/s, 4.6601 assertions/s.`  
`509 runs, 1413 assertions, 0 failures, 0 errors, 0 skips`

How large are you test suites and how long does it take to go through them.
## [10][What would be a good way to make an app with sound files that are streamable and downloadable?](https://www.reddit.com/r/rails/comments/jser47/what_would_be_a_good_way_to_make_an_app_with/)
- url: https://www.reddit.com/r/rails/comments/jser47/what_would_be_a_good_way_to_make_an_app_with/
---
Hi all!

I’m building a Rails app where users should be able to stream and/or download sound files. I’d like it to have similar functionality to this site: [https://www.soundsnap.com/animals/birds](https://www.soundsnap.com/animals/birds) I’ll record all the sounds myself.

What would be the best way to go about a project like this? I've built a couple of simple Rails apps before but nothing that deals with a lot of static assets like. Any advice on how to approach something like this would be much appreciated!

Some questions that come to mind are:

* Should I store the files on some kind of cloud storage like AWS S3 or Digitalocean Spaces? Are there any others that are reliable and cheap?
* Should I use a CDN for caching?
* How could I stream the audio? Would using HTML5 audio tags work well enough? Or maybe Soundcloud's embedded player?

Any suggestions or guidance in general would be fantastic!
## [11][Does ActiveStorage auto-purge when updating a record?](https://www.reddit.com/r/rails/comments/js96mq/does_activestorage_autopurge_when_updating_a/)
- url: https://www.reddit.com/r/rails/comments/js96mq/does_activestorage_autopurge_when_updating_a/
---
I am away from a Rails development machine to test this, and I think I want an answer sooner, as my Googling gives me vague answers.

Let's say I have this model.
```
class Product &lt; ApplicationRecord
  has_one_attached :image
end
```

When creating a record with an attachment it goes something like this:
```
product = Product.new
product.image.attach(attachment_image) ## This attaches and uploads an image.
product.save
```

When updating a record with a new attachment, a similar code looks like this.
```
product = Product.find(1)
product.image.attach(new_attachment_image) ## Does this purge the first uploaded file and then uploads the new file? Or it keeps them both?
product.save
```

Do I need to `purge` it first before attaching a new file? This is just a shower thought but I don't have a way to confirm.
## [12][Help fixing bug on tempusdominus-bootstrap-4 package](https://www.reddit.com/r/rails/comments/jsfnk9/help_fixing_bug_on_tempusdominusbootstrap4_package/)
- url: https://www.reddit.com/r/rails/comments/jsfnk9/help_fixing_bug_on_tempusdominusbootstrap4_package/
---
Hi, guys. I need some help here. I'm using the tempusdominus-bootstrap-4 package with Rails 6 &amp; Webpacker and I'm getting a error message on browser console. While on development is working fine, on production, where all js file are copy &amp; transpiled into one file, I'm getting the following error message:

    tempusdominus-bootstrap-4.js:31 Uncaught TypeError:  t is not a function
      at tempusdominus-bootstrap-4.js:31
      at Object.&lt;anonymous&gt; (tempusdominus-bootstrap-4.js:30)
      at Object.&lt;anonymous&gt; (application-a354d07e2c01c614e691.js:1)
      at n (bootstrap:19)   
      at Module.&lt;anonymous&gt; (application-a354d07e2c01c614e691.js:1)   
      at n (bootstrap:19)   at bootstrap:83   at application-a354d07e2c01c614e691.js:1 

When I open the the file on line 31, I got the following:

    +function () {
      var _typeof = typeof Symbol === "function" &amp;&amp; _typeof Symbol.iterator === "symbol" ? function (obj) {     
        return _typeof obj;   
      } : function (obj) {
        return obj &amp;&amp; typeof Symbol === "function" &amp;&amp; obj.constructor === Symbol &amp;&amp; obj !== Symbol.prototype ? "symbol" : _typeof obj;   
      };   
      // more stuffs here 
    }; 

While on the original code and same line we have the following:

    +function () {   
      var _typeof = typeof Symbol === "function" &amp;&amp; typeof Symbol.iterator === "symbol" ? function (obj) {
        return typeof obj;    
      } : function (obj) {      
        return obj &amp;&amp; typeof Symbol === "function" &amp;&amp; obj.constructor === Symbol &amp;&amp; obj !== Symbol.prototype ? "symbol" : typeof obj;    
      };   
      // more stuffs here 
    }; 

I'm not a expertise on javascript nor webpack, but I could see that, after webpack copy &amp; transpile all imported packages into one file, some `typeof` defined keyword were changed to `_typeof`. I thought that was the cause of error, so I tried to fixed it changing all references (found 2) of `_typeof`to `_testTypeof` on the original code. Then it worked! I got the following code from webpack precompile (copy &amp; transpile):

    +function () {   
      var _testTypeof = typeof Symbol === "function" &amp;&amp; _typeof(Symbol.iterator) === "symbol" ? function (obj) {
        return _typeof(obj);   
      } : function (obj) {     
        return obj &amp;&amp; typeof Symbol === "function" &amp;&amp; obj.constructor === Symbol &amp;&amp; obj !== Symbol.prototype ? "symbol" : _typeof(obj);   
      };   
      // more stuffs here 
    }; 

However, I wasn't expecting `typeof` to become `_typeof`, again. But, still, the change was enough to start to working on production and doesn't raise an error message on browser console. Does anyone already a problem like this? The problem may be the transpiler (babel)? I don't know what to do.

Link to the source on git: [https://github.com/tempusdominus/bootstrap-4/blob/master/build/js/tempusdominus-bootstrap-4.js](https://github.com/tempusdominus/bootstrap-4/blob/master/build/js/tempusdominus-bootstrap-4.js)
