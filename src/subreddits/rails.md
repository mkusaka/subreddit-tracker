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
## [3][Any good resources for beginners on how to write good tests?](https://www.reddit.com/r/rails/comments/jsmecm/any_good_resources_for_beginners_on_how_to_write/)
- url: https://www.reddit.com/r/rails/comments/jsmecm/any_good_resources_for_beginners_on_how_to_write/
---
I know there are tons of tutorials/docs out there about testing, but are there good ones that teaches good practical/pragmatic approaches on testing rails apps? Something that teaches good practice on what to test, instead of just how to test things; what are bad tests, how to avoid bugs in tests etc
## [4][How large is your test suite?](https://www.reddit.com/r/rails/comments/jseicl/how_large_is_your_test_suite/)
- url: https://www.reddit.com/r/rails/comments/jseicl/how_large_is_your_test_suite/
---
Just ran my test suite for the app I am working on.   


Here's are the stats:

`Finished in 303.213471s, 1.6787 runs/s, 4.6601 assertions/s.`  
`509 runs, 1413 assertions, 0 failures, 0 errors, 0 skips`

How large are you test suites and how long does it take to go through them.
## [5][What would be a good way to make an app with sound files that are streamable and downloadable?](https://www.reddit.com/r/rails/comments/jser47/what_would_be_a_good_way_to_make_an_app_with/)
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
## [6][Does ActiveStorage auto-purge when updating a record?](https://www.reddit.com/r/rails/comments/js96mq/does_activestorage_autopurge_when_updating_a/)
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
## [7][Help fixing bug on tempusdominus-bootstrap-4 package](https://www.reddit.com/r/rails/comments/jsfnk9/help_fixing_bug_on_tempusdominusbootstrap4_package/)
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
## [8][Rails 6 form_with (remote: true) not working on production, only on development](https://www.reddit.com/r/rails/comments/jrt60p/rails_6_form_with_remote_true_not_working_on/)
- url: https://www.reddit.com/r/rails/comments/jrt60p/rails_6_form_with_remote_true_not_working_on/
---
Hi, guys! I need some help here. I updated my gems and packages recently and while everything is work on development, on production I noticed that the form\_with helper are not working properly. It seems to me that form\_with is not making a XHR/AJAX, or if does it change the whole page for the requested partial (e.g flash message) - this doesn't happen on development.

The following is how I'm defining the form\_with and the controller.

* One note is that I'm using StimulusJS to load the requested content to the view. I can confirm that Stimulus is working because only with form\_with I'm having problems. I believe that the XHR/AJAX actions are not trigger because rails is not making a XHR/AJAX. Another note is that I'm requesting a html page (e.g: flash message or error message).

&amp;#x200B;

    # edit view
    &lt;%= 
      form_with model: @entity, 
      url: entity_path(@entity),
      method: :put,
      data: { 
        type: "html",
        action: 'ajax:success-&gt;form#scrollUp
        ajax:success-&gt;form#successContentReplace 
        ajax:error-&gt;form#failureContentReplace', 
        content: "content_form_message" 
      } do |f| %&gt; 
        
        /* html inputs */
     &lt;%= f.submit "Save", class: "btn btn-success btn-block" %&gt;
    %&gt;
    
    # entity controller
    def update
      @entity = ::Entity.find(params[:id])
      respond_to do |format|
        if @entidade.update(entidade_params)
          format.html do
            flash.now[:success] = "Atualização realizada com sucesso!"
            render partial: 'layouts/shared/flash_messages'
          end
        else
          format.html do
            render partial: 'layouts/shared/error_messages',
            locals: { object: @entity }
          end
        end
      end
    end

My guess is that the problem is with Rails\_ujs (or, unlikely, Turbolink). They may be not recognizing that the form\_with is making a XHR/AJAX because of some updated syntax. On the other hand, this only happens on production. Rails\_ujs makes use of JQuery and I have it load by webpacker.

A bonus: The Devise's sign\_out method is not work on production, too. I'm getting 'route not found', but the route exist because it is working on development. I think this is happen because of the same problem. To call the sign\_out method on view I use `remote: true`, in other words, rails\_ujs.

Does anyone pass through this problem too when updating to rails 6?

&amp;#x200B;

# #Edit

The problem was a conflict between packages. After removing the package *tempusdominus-bootstrap-4* everything start to work fine. I believe the conflict was on the rails\_ujs and tempusdominus-bootstrap-4 packages, because both uses JQuery and was related to my problem.

The error message the was

    Uncaught TypeError: t is not a function
      at tempusdominus-bootstrap-4.js:31
      at Object.&lt;anonymous&gt; (tempusdominus-bootstrap-4.js:30)
      at Object.&lt;anonymous&gt; (application-2e0b3414c2caaddca3bc.js:1)
      at n (bootstrap:19)
      at Module.&lt;anonymous&gt; (application-2e0b3414c2caaddca3bc.js:1)
      at n (bootstrap:19)
      at bootstrap:83
      at application-2e0b3414c2caaddca3bc.js:1

I will raise a issue on the tempusdominus-bootstrap-4 github.
## [9][What's the best way to find where a template appears in a Rails app?](https://www.reddit.com/r/rails/comments/jrr6s6/whats_the_best_way_to_find_where_a_template/)
- url: https://www.reddit.com/r/rails/comments/jrr6s6/whats_the_best_way_to_find_where_a_template/
---
I started a new job recently as a Rails developer. I'm only a couple of weeks in so the codebase is still very new to me. I've started tackling support tickets, where often a support person will describe a change to a page that they want but not include the URL or a great description of how to actually get there. But I can usually find the template in the code. Obviously in these instances I can just ask for more information, but I want to get better at figuring this kind of thing out myself. My question is: given a view in a Rails app, what's the best way to find the URL that displays that template? Sometimes it's not as easy as running rake routes
## [10][How to query for only one post from each user?](https://www.reddit.com/r/rails/comments/jrm32k/how_to_query_for_only_one_post_from_each_user/)
- url: https://www.reddit.com/r/rails/comments/jrm32k/how_to_query_for_only_one_post_from_each_user/
---
Environment: Rails 6, Postgres.

Models: User has\_many posts.

Request: I wish to list the latest single published post from users with a limit of 10. I only want to see one post per user in the list.

I tried these but neither is working:

&gt;posts = Post.published.select("DISTINCT ON (posts.user\_id) posts.\*).order(id: :desc).limit(10)

This above gives an error: PG::InvalidColumnReference: ERROR:  SELECT DISTINCT ON expressions must match initial ORDER BY expressions

&gt;posts = Post.published.distinct(user\_id).order(id: :desc).limit(10)

This one gives no error, but disregards the distinct part. I get multiple posts by user.

Please help!
## [11][What OS do you use?](https://www.reddit.com/r/rails/comments/jrk7kz/what_os_do_you_use/)
- url: https://www.reddit.com/r/rails/comments/jrk7kz/what_os_do_you_use/
---
Hey Ruby People, 

I am thinking of switching from MacOS to Linux for my development machine. Have you made the switch, or do you dev on a linux machine? 

I'd love to know what peoples opinions are. :)
## [12][Get user information from Access-Token ( omniauth-google-oauth2)](https://www.reddit.com/r/rails/comments/jrlz0e/get_user_information_from_accesstoken/)
- url: https://www.reddit.com/r/rails/comments/jrlz0e/get_user_information_from_accesstoken/
---
hi u/rails hope you are fine!   
Actually, I'm making an API in which I'm getting the Access token From the frontend &amp;  I have to log in user by getting information from that token. 

I'm using omniauth but I don't know if I get a token from the front-end, how can I get the user info from that token by requesting back to google.  
it can be quite easy if im following monolithical structure but for Api's I don't know the flow.

can anyone please help me ?
