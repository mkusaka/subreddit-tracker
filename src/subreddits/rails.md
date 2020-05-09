# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gek2rk/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gek2rk/personal_projects_show_off_your_own_project_andor/
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
## [2][Profiling Rails app that uses websockets](https://www.reddit.com/r/rails/comments/gg7dpw/profiling_rails_app_that_uses_websockets/)
- url: https://www.reddit.com/r/rails/comments/gg7dpw/profiling_rails_app_that_uses_websockets/
---
Hi. I'm new to Ruby and I am trying to profile a Rails app that uses websockets for streaming audio packets and other messages. I am interested in CPU time and wall clock time taken by all the methods in the app. Most of the tracing gems I see show the information for a single HTTP request, but it doesn't help in the case of websockets, since they remain alive even after a request. If it helps, 

How would you do this? How do I profile arbitrary points in code, like just do a "start" in the websocket open and "stop" in the websocket close, and generate a report? I tried ruby-prof, but it crashes with cryptic errors...
## [3][Good places to start with IOS development?](https://www.reddit.com/r/rails/comments/ggeq6a/good_places_to_start_with_ios_development/)
- url: https://www.reddit.com/r/rails/comments/ggeq6a/good_places_to_start_with_ios_development/
---
Hey everyone! 

I have recently completed a full stack development course (Rails 6, Ruby, HTML, CSS, SQL, simple JavaScript) and feel relatively comfortable creating simple websites and hosting them on Heroku. 

I now wanted to look into IOS development, as I really want my next project (time management software) to also have a mobile application for users. 

Do you think React Native is the right way to go? I was unsure if it was wise to venture to far from the skills I recently learned, as I hoped this project would allow me to build upon and strengthen them.

Are there any good books / guides or tutorials that I could follow to get a coherent understanding?

Looking forward to reading your responses, thanks a lot in advance and have a great weekend! :)
## [4][link_to without the &lt;a&gt;. To add in the &lt;head&gt;](https://www.reddit.com/r/rails/comments/ggd97p/link_to_without_the_a_to_add_in_the_head/)
- url: https://www.reddit.com/r/rails/comments/ggd97p/link_to_without_the_a_to_add_in_the_head/
---
Hi guys,

in the "structured data" area of my website, I added a potentialAction, SearchAction,

in this way

    	"potentialAction": {
    		"@type": "SearchAction",
    		"target": "https://www.mywebsite.com/search?utf8=%E2%9C%93&amp;amp;search%5Bq%5D={search_term_string}",
    		"query-input": "required name=search_term_string"
    	}

the link/target is not the best thing to see. I want to replace it with a link\_to, but if I should add

    &lt;%= link_to search_path(search: { q: #{search_term_string} }) %&gt;

obviously I will have the &lt;a&gt; tag. **How to put it without the &lt;a&gt;?**

Just to create the url
## [5][Best course resource for ruby and rails?](https://www.reddit.com/r/rails/comments/gfx93d/best_course_resource_for_ruby_and_rails/)
- url: https://www.reddit.com/r/rails/comments/gfx93d/best_course_resource_for_ruby_and_rails/
---
TL;DR - what course website should I get my company to pay for that has best rails content?


I’ve been using Rails for about a year. The purpose is business applications which is based around CRUD plus sometimes tricky client reqs. I get along fine in Rails but I feel there are some fundamentals and good habits which may be missing.  Also looking at improving things such as DRY, refactoring, StimulusReflex, and generally beautiful but maintainable code. 

I read this article and found it also interesting http://jeromedalbert.com/how-dhh-organizes-his-rails-controllers/. 

It wouldn’t hurt to improve my JS, HTML5, Bootstrap, etc. 

So looking for a nice resource which will cover these and be somewhat relevant to rails. 

I currently have access to a friends pluralsight. They also have a sale going on. Should I just get the company to purchase that?

Any other recommendations? Thanks.
## [6][Best way to convert data from a view or model into a spreadsheet](https://www.reddit.com/r/rails/comments/gfxr5k/best_way_to_convert_data_from_a_view_or_model/)
- url: https://www.reddit.com/r/rails/comments/gfxr5k/best_way_to_convert_data_from_a_view_or_model/
---
haven't done this before, was wondering if anyone had a nice article or video for using caxlsx to do so
## [7][Webpacker and CSS image calls](https://www.reddit.com/r/rails/comments/gg05c9/webpacker_and_css_image_calls/)
- url: https://www.reddit.com/r/rails/comments/gg05c9/webpacker_and_css_image_calls/
---
I am still new to a non asset pipeline world. I am working on transitioning my app from 4 to 6. I have a number of image calls in my CSS that webpack without error, but do not load. 

    .logo {
    background: image-url("../images/seal.png") no-repeat scroll 0 0 transparent;
    background-size: 45px auto;
        ...
    }
      .page-header a {
        background: image-url("../images/neuro_banner_w.jpg") no-repeat scroll 0px 0 transparent;
    background-size: 100% auto;
       ... 
        }

My images now all reside in the app/javascript/images folder (with subfolders included, but neither of these images are in a subfolder.

When I load my page, no images appear. When I open the inspector the I see:

    background: image-url(/packs/media/images/seal-275a0ac1ae6079ecad8e70e483ffbfee.png) no-repeat scroll 0 0 transparent;

The line is crossed out with the error: Invalid Property value. 

Also the following two lines are also crossed out (background size and height)

What am I doing wrong?

I am not sure what files you need to see code from, so call out if there is something that I need to show.
## [8][[Question/help] Shopify + GraphQL](https://www.reddit.com/r/rails/comments/gfxjz3/questionhelp_shopify_graphql/)
- url: https://www.reddit.com/r/rails/comments/gfxjz3/questionhelp_shopify_graphql/
---
I am new to building shopify applications, and am a bit confused regarding GraphQL.

Currently:

I have built a super simple Shopify application using the shopify_app gem. The application has some webhooks associated with orders. When a shop creates an order, I receive the JSON from the order, run a job in my code and submit the data I want to my local database.

What I am trying to do:

I need to access all orders from a shopify store that installs my application, and store all orders for that store. I would want to fetch all orders every 4-months or so. Shopify gave my application permission to fetch all orders.

My issue/question:

It appears the best course of action here is to use GraphQL. What I am confused with is do I need to use the shopify_api gem in addition to the shopify_app gem to get GraphQL to work?
## [9][Tailwind form question.](https://www.reddit.com/r/rails/comments/gg022m/tailwind_form_question/)
- url: https://www.reddit.com/r/rails/comments/gg022m/tailwind_form_question/
---
Using tailwind on Rails 6 with Webpack and all that good stuff. It's my first time with tailwind so I'm just stumbling through just how "Open Source" this framework is in nature. Could almost be called option paralysis if I'm being honest. Now, unto the question:

&amp;#x200B;

I have a sign-up &amp; new post form that I'm very happy with so far. However, I am thoroughly disappointed that the text box looks absolutely retarded (really small initially) and has to be manually resized. Is there a way that I can make it start at a designated size and have it self-adjust as needed? Or is this something tailwind's out-of-the-box CSS doesn't care for?
## [10][react-rails vs rails-api and react app](https://www.reddit.com/r/rails/comments/gfo0vg/reactrails_vs_railsapi_and_react_app/)
- url: https://www.reddit.com/r/rails/comments/gfo0vg/reactrails_vs_railsapi_and_react_app/
---


[View Poll](https://www.reddit.com/poll/gfo0vg)
## [11][How to set variable on page load?](https://www.reddit.com/r/rails/comments/gft14w/how_to_set_variable_on_page_load/)
- url: https://www.reddit.com/r/rails/comments/gft14w/how_to_set_variable_on_page_load/
---
So I have a hidden form field that gets set with a value onclick from a couple of links. Depending on what link is clicked a different value gets set. The form field is called activeselector.

What I want to do is set an `@activeselector` variable to the form param if it is there, but set it to '12345' or something if it is not so the page loads with a default value.

Whats the best way to do this? I have the onclick set up and setting the value, but I'm not sure the best  way to return it back to the page on submit and default it on first page load.
