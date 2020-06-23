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
## [3][how to improve multi insert](https://www.reddit.com/r/rails/comments/hebb6x/how_to_improve_multi_insert/)
- url: https://www.reddit.com/r/rails/comments/hebb6x/how_to_improve_multi_insert/
---
i have list of values and need insert in db for general i use `each` for loop.

`data = {`

`row1: {id:1, name: 'test',  type: 'x'},`

`........... etc`

`}`

&amp;#x200B;

how can improve it  . i hear  about \`activerecord-import\` for reduce queries
## [4][Wordpress vs Rails](https://www.reddit.com/r/rails/comments/hee2b3/wordpress_vs_rails/)
- url: https://www.reddit.com/r/rails/comments/hee2b3/wordpress_vs_rails/
---
lets say you know how to make Wordpress websites with all their different plugins and themes (no php) and you know Rails (medior).

In what case would you pick making a website in RoR over WP?
## [5][Stripe Users: How do you handle product data with Rails?](https://www.reddit.com/r/rails/comments/hdwode/stripe_users_how_do_you_handle_product_data_with/)
- url: https://www.reddit.com/r/rails/comments/hdwode/stripe_users_how_do_you_handle_product_data_with/
---
I noticed that Stripe wants you to create products in their API, which sets off alarm bells about data going out of sync locally if I also have a \`products\` table.

How do people in the community integrate Stripe product data with Rails and ActiveRecord. Are there any best practices?
## [6][How to implement a ThemeForest front end the "Rails 6 way"?](https://www.reddit.com/r/rails/comments/hdyjjz/how_to_implement_a_themeforest_front_end_the/)
- url: https://www.reddit.com/r/rails/comments/hdyjjz/how_to_implement_a_themeforest_front_end_the/
---
Hey everyone, I've been bouncing around for a week now, watching videos, reading tutorials, and experimenting, and I feel like I'm doing this all wrong. Instead of asking how to fix this particular problem or that problem, I think I need to step back and ask **how I should be doing this on a brand new project... What's the Rails 6 way?**

So I'm building a new Rails 6 app, and since I'm not a front-end person, I purchased a couple ThemeForest HTML bootstrap admin templates. The first one was pretty simple. After trying to implement it with WebPacker, I gave up, and I sort of got it working through the asset pipeline - but I had to do a lot of tweaking.  
By that I mean I put all the fonts, css, images and js in subfolders under /apps/assets. Renamed some conflicting css and js file names. Renamed the icons.css (where the font references are) to icons.scss and then line-by-line changed the urls to something like this:

    src: url(font-path('fa-brands-400.eot')

Well, now I've decided I don't like the theme, and I've moved to a much better one called Metronic (very popular).  
 [https://keenthemes.com/metronic/?page=index](https://keenthemes.com/metronic/?page=index) 

The theme's docs quickstart section shows two way of getting started: Using WebPack or using Gulp. You decide which of 11 demos you like best, then if you want, tweak files under the src directory, then run " npm run build --demo1" (for webpack) or " gulp *--demo1" (for gulp).*  
That compiles all the assets into a dist directory with all the css, js, images and fonts.

Both options produce a small number of xxxx.bundle.css and xxx.bundle.js files that are called in the HTML examples. You also get a packages.json with all the yarn dependencies. 

Well, since WebPacker is the new hotness in Rails 6, and only sort of understanding what WebPack/WebPacker is and does, I started with the WebPack version. I know right now you're supposed to use WebPack for javascript, and asset pipeline for css, fonts and media - but most of the videos I watched just went ahead and used WebPack for everything.

So I tried it, and the CSS, images and fonts seem to work using WebPacker. No CSS tweaking to find the fonts.  
The javascript - not so much - because the javascript isn't getting loaded into the global scope (apparently that's how WebPacker is supposed to work). With tons of lines of javascript, it's pretty intimidating to figure out how to get this pre-written code into the right scopes. I'm not even sure how I would go about that.

That being the case, I put the javascript in the asset pipeline, and css/etc in WebPacker. Totally flipping the script. But it works. 

Ok, not totally. The pre-compiling in development is so slow that half the time my pages wont load (1500ms timeout). And when I tried to deploy to Heroku, all hell broke loose.

I've also tried the Gulp version exclusively in the asset pipeline. Sort of works in dev, but pukes on precompiling at Heroku.

So, instead of banging my head against trying to get either approach to work, it's time to step back and ask the experts: Which route should I be going? What's the best approach to take? What's the Rails 6 way?  
Bonus question: Is there a good reference on how to bring a modern HTML template into a Rails 6 app? Video, book, blog post, etc?  
Bonus bonus question: Since it's so popular, has anyone worked specifically with Metronic?

For reference, this is the docs for the HTML template:  
 [https://keenthemes.com/metronic/?page=docs&amp;section=html-webpack-quick-start](https://keenthemes.com/metronic/?page=docs&amp;section=html-webpack-quick-start) 

Thanks for any and all guidance!
## [7][Upgrading ruby version on production server](https://www.reddit.com/r/rails/comments/he0zcr/upgrading_ruby_version_on_production_server/)
- url: https://www.reddit.com/r/rails/comments/he0zcr/upgrading_ruby_version_on_production_server/
---
Does anyone have experience with this? Is it any different from doing it in the development environment?

If anyone could point me to a tutorial or just give me some tips, I'd very much appreciate it
## [8][Please review my ServiceHelper](https://www.reddit.com/r/rails/comments/hdulyf/please_review_my_servicehelper/)
- url: https://www.reddit.com/r/rails/comments/hdulyf/please_review_my_servicehelper/
---
I'm working on app which have many external APIs calls. I don't want to have fat wrapper, so I made smart(at least I'm tried) base_service class, to make few dozens light services, but I have doubts if is it overcomplicated, imho I like now my tiny services which now have about 20 lines of code, but I want to know if my approach is proper way. Here is code: https://github.com/rwegrzyniak/FaradayBaseService
Can you make fast code review?
## [9][Recommended books](https://www.reddit.com/r/rails/comments/he2zqd/recommended_books/)
- url: https://www.reddit.com/r/rails/comments/he2zqd/recommended_books/
---
What are the best books you would recommend about Ruby/Rails?
## [10][How to have distinct forms for multiple objects in the same view](https://www.reddit.com/r/rails/comments/hdv8vt/how_to_have_distinct_forms_for_multiple_objects/)
- url: https://www.reddit.com/r/rails/comments/hdv8vt/how_to_have_distinct_forms_for_multiple_objects/
---
I'm building a shopping cart and one of the features I'd like to be able to implement is to allow the shopper to select the quantity of item they want. I figure I could do this with a form_for for each object that's currently in the cart, but worried that there would be conflicts. How can I assign a unique form input (which will then be available as params for the next page)? Would it be something like:

     &lt;%= form_for(:product_name, url: cart_path), html: {id: "quantity-form"}) do |f| %&gt;
        &lt;%= f.label :quantity, class: 'field-label emphasis contact-form-item-mobile' do %&gt;
            &lt;span class = 'field-label emphasis'&gt; Quantity &lt;/span&gt;
        &lt;% end %&gt;
        &lt;%= f.number_field :quantity, class: 'form-control form-control-mobile', placeholder: "Quantity" %&gt;&lt;br&gt;
    &lt;% end %&gt;
## [11][Help optimizing SQL query generation in rails active record](https://www.reddit.com/r/rails/comments/hdw20h/help_optimizing_sql_query_generation_in_rails/)
- url: https://www.reddit.com/r/rails/comments/hdw20h/help_optimizing_sql_query_generation_in_rails/
---
I have a \`user\` model that has\_many \`posts\`. For this example suppose I have to do two things.

1. Check if the user any posts present
2. If test, check if the post is present in one of the following types. 

This example is a bit contrived but it like the posts exist only if the user is of a certain kind. something like that. so please bear with me on this. 

Approach one

user = User.find(1)

return user.posts.present? &amp;&amp; (user.posts.map(&amp;:type) &amp; types).present?

Approach 2

user = User.find(1)

return true unless user.posts.exists?

return user.posts.where(type: types).exists?

While I thought the approach 2 would be more efficient, I noticed that in the 1st approach since posts are already loaded in the present? check the subsequent map doesn't fire a new query.  So say just 1 DB query. 

In the second approach probably since there is a where query it ends up firing 2 queries to the DB. 

Could you let me know if this is this is the case and any information on how rails caches these queries? I tried it only in a rails console.
## [12][Writing to a file is taking a lot of time in the prod environment](https://www.reddit.com/r/rails/comments/hdrubu/writing_to_a_file_is_taking_a_lot_of_time_in_the/)
- url: https://www.reddit.com/r/rails/comments/hdrubu/writing_to_a_file_is_taking_a_lot_of_time_in_the/
---
I know it's kind of weird. I have changed the in-memory generation of csv to a file-based one. Locally, it  generates the file relatively faster 20 k records in 20 mins or so. But in the production environment it's taking too much time. It's writing 1 k records over 20 mins. Is it because of I/O overhead? or am I looking at this wrong

This is the code I have

`tempfile = File.open(Rails.root.join('tmp', "#{SecureRandom.hex(8)}.csv"), 'wb')CSV.open(tempfile.path, 'w') do |csv|ActiveRecord::Base.uncached doModel.find_each do |model|csv &lt;&lt; ["foobar"]endendendtempfile`

&amp;#x200B;

Looks like writing to file is indeed an overhead, what are my options now? I cannot create file in memory due to RAM reasons

This is how much Benchmark results looks like

0.019205 0.001835 0.021040 ( 0.021893) - local

0.455934 0.035429 0.491363 ( 0.725257) - prod

&amp;#x200B;

&amp;#x200B;
