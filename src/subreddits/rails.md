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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/jnwqje/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/jnwqje/personal_projects_show_off_your_own_project_andor/
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
## [3][Rails 6 form_with (remote: true) not working on production, only on development](https://www.reddit.com/r/rails/comments/jrt60p/rails_6_form_with_remote_true_not_working_on/)
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
## [4][What's the best way to find where a template appears in a Rails app?](https://www.reddit.com/r/rails/comments/jrr6s6/whats_the_best_way_to_find_where_a_template/)
- url: https://www.reddit.com/r/rails/comments/jrr6s6/whats_the_best_way_to_find_where_a_template/
---
I started a new job recently as a Rails developer. I'm only a couple of weeks in so the codebase is still very new to me. I've started tackling support tickets, where often a support person will describe a change to a page that they want but not include the URL or a great description of how to actually get there. But I can usually find the template in the code. Obviously in these instances I can just ask for more information, but I want to get better at figuring this kind of thing out myself. My question is: given a view in a Rails app, what's the best way to find the URL that displays that template? Sometimes it's not as easy as running rake routes
## [5][What OS do you use?](https://www.reddit.com/r/rails/comments/jrk7kz/what_os_do_you_use/)
- url: https://www.reddit.com/r/rails/comments/jrk7kz/what_os_do_you_use/
---
Hey Ruby People, 

I am thinking of switching from MacOS to Linux for my development machine. Have you made the switch, or do you dev on a linux machine? 

I'd love to know what peoples opinions are. :)
## [6][How to query for only one post from each user?](https://www.reddit.com/r/rails/comments/jrm32k/how_to_query_for_only_one_post_from_each_user/)
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
## [7][Get user information from Access-Token ( omniauth-google-oauth2)](https://www.reddit.com/r/rails/comments/jrlz0e/get_user_information_from_accesstoken/)
- url: https://www.reddit.com/r/rails/comments/jrlz0e/get_user_information_from_accesstoken/
---
hi u/rails hope you are fine!   
Actually, I'm making an API in which I'm getting the Access token From the frontend &amp;  I have to log in user by getting information from that token. 

I'm using omniauth but I don't know if I get a token from the front-end, how can I get the user info from that token by requesting back to google.  
it can be quite easy if im following monolithical structure but for Api's I don't know the flow.

can anyone please help me ?
## [8][Asking for recommended way to sell digital goods](https://www.reddit.com/r/rails/comments/jrkpxm/asking_for_recommended_way_to_sell_digital_goods/)
- url: https://www.reddit.com/r/rails/comments/jrkpxm/asking_for_recommended_way_to_sell_digital_goods/
---
I have been looking around for this kind of platforms, and stumbled upon Gumroad, BuyMeACoffee, and etc as possible solutions.

&amp;#x200B;

Any of you have experiences with these platforms or something similar to that? Or is it better to just make our own solution?
## [9][How to invalidate nested pertain cache?](https://www.reddit.com/r/rails/comments/jrgpbg/how_to_invalidate_nested_pertain_cache/)
- url: https://www.reddit.com/r/rails/comments/jrgpbg/how_to_invalidate_nested_pertain_cache/
---

    cache “1” do
      cache “2” do 
         cache “3” do


If I change temple or code that produced cache 3, how can I invalidate 2, which then invalidates 1 aswell?

The idea was to have a cache based on filehash + templatehash. 

Children are called within the template , so I don’t know yet what children will be called 

Having an additional VERSION inside the classes is ok but then I need to change version +1 for 1 and 2, just if I change 3


General programming question I guess
## [10][How to link to github file for current rails view file](https://www.reddit.com/r/rails/comments/jrhand/how_to_link_to_github_file_for_current_rails_view/)
- url: https://www.reddit.com/r/rails/comments/jrhand/how_to_link_to_github_file_for_current_rails_view/
---
I am working on a project with other people who are not too familiar with where files for different parts of my app are to be found. I'd like to add links that "open source file for this component on github".

Does anyone know if there is already such a gem available? If not, would anyone like to work with me on creating one. I imagine it could be very useful.
## [11][For an Index page, should I calculate this info on the fly, or store it in the database?](https://www.reddit.com/r/rails/comments/jrbgnm/for_an_index_page_should_i_calculate_this_info_on/)
- url: https://www.reddit.com/r/rails/comments/jrbgnm/for_an_index_page_should_i_calculate_this_info_on/
---
Hey there, I am working on a simple invoicing app and for my Clients index view, which will probably show 12-15 records at a time, I want each Client's table line to include some data like:

* How many outstanding invoices that client has (ie, 23 invoices)
* How many Paid invoices they have
* The dollar amount of all their invoices (ie, $18,453.52 total)
* The dollar amount of outstanding invoices

This seems like it might be a burden to have Rails calculate all this for 12-15 clients per index page, yes? This is my first project, so I'm just not sure, but it seems like it involves a lot of database work every time a user simply looks at their Client index page. So I was going to create a couple of Count values and Totals values columns the Client table, and just change them anytime a new invoice is created/deleted/paid.

Does this seem reasonable?

Thanks much!
## [12][Amazing course for on RAILS preferably PAID](https://www.reddit.com/r/rails/comments/jr26sa/amazing_course_for_on_rails_preferably_paid/)
- url: https://www.reddit.com/r/rails/comments/jr26sa/amazing_course_for_on_rails_preferably_paid/
---
I can take courses at work, and I'm thinking of taking a rails course to up my skills. The course can be paid one. What I'm really looking for in the course.  

1. Internals of rails app, lesser-known features, identifying performance bottlenecks, and optimizing code with design patterns and good testing techniques. 
2. Using APM's like new relic to monitor the health of the app
3. The devops side of deploying, CI, CD etc of a rails app. 

I do not want a beginner's course on creating a simple CRUD app. I want it to be something a little advanced than that. 

Could someone please suggest a good, paid course on this, or even multiple courses that cover this? 

Thanks.
