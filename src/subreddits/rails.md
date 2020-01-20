# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/ep2dw9/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/ep2dw9/personal_projects_show_off_your_own_project_andor/
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
## [2][Example project with Vue + Rails + SPA](https://www.reddit.com/r/rails/comments/erad9q/example_project_with_vue_rails_spa/)
- url: https://www.reddit.com/r/rails/comments/erad9q/example_project_with_vue_rails_spa/
---
Looking for a guide or a repo that has vuejs on a rails project, and it's an SPA. (will gold lol)
## [3][Is Facebook lite a good portfolio project?](https://www.reddit.com/r/rails/comments/erajnl/is_facebook_lite_a_good_portfolio_project/)
- url: https://www.reddit.com/r/rails/comments/erajnl/is_facebook_lite_a_good_portfolio_project/
---
Here is a good example

[https://dev-communities.herokuapp.com/](https://dev-communities.herokuapp.com/)

This is the last of many rails projects in the Odin project (an opensource curriculum that takes over a year to complete)

[https://www.theodinproject.com/courses/ruby-on-rails/lessons/final-project](https://www.theodinproject.com/courses/ruby-on-rails/lessons/final-project)

Is this actually going to impress any employers?

A lot of you have told me in the past to contribute to Open source and that has been NOTED but how does this go for a portfolio project?
## [4][Webpacker, how to include a JS library's image path?](https://www.reddit.com/r/rails/comments/erag0g/webpacker_how_to_include_a_js_librarys_image_path/)
- url: https://www.reddit.com/r/rails/comments/erag0g/webpacker_how_to_include_a_js_librarys_image_path/
---
I've been on Sprockets until this current project where I don't have a choice. I'm trying to add Leaflet, which I install like this:

    yarn add leaflet

Which gives me this structure in `node_modules/`

    leaflet
      dist
        images
          layers.png
          marker-icon.png
          ...
        leaflet.js
        leaflet.css

I've required in `application.js`

    require("leaflet");

And imported in `application.scss`

    @import "leaflet/dist/leaflet";

Which is loading everything fine, until I try to add a marker and I get this error in the web inspector, and obviously no marker image:

    GET http://localhost:3000/assets/images/marker-icon.png 404 (Not Found)

Because `leaflet.css` has its own references to images:

    .leaflet-default-icon-path {
      background-image: url(images/marker-icon.png);
    }

How to I let the leaflet library CSS and JS know about the images directory that’s now in `node_modules/`? Is there something I can do in the webpacker config to move the images as part of the packing?
## [5][What IRC or other chat system has ongoing rails discussions?](https://www.reddit.com/r/rails/comments/er0vab/what_irc_or_other_chat_system_has_ongoing_rails/)
- url: https://www.reddit.com/r/rails/comments/er0vab/what_irc_or_other_chat_system_has_ongoing_rails/
---
I joined some IRC channels and though there were a lot of people, there were no active discussions. Where do I go for active rails channels?
## [6][Is it possible to use the GraphiQL Playground in a API-only Rails app?](https://www.reddit.com/r/rails/comments/er7o5h/is_it_possible_to_use_the_graphiql_playground_in/)
- url: https://www.reddit.com/r/rails/comments/er7o5h/is_it_possible_to_use_the_graphiql_playground_in/
---
I have an API-only Rails app that I am adding GraphQL to. I am going through this GraphQL tutorial: [https://www.howtographql.com/graphql-ruby/2-queries/](https://www.howtographql.com/graphql-ruby/2-queries/) , and they recommend using GraphQL's built-in Playground interface, which looks awesome. However, because my Rails app is API-only, it can't be built, I think because it isn't setup to display templates. Does anyone know if there is a way to get around this?
## [7][Rails + React issue with onChange callback](https://www.reddit.com/r/rails/comments/er8432/rails_react_issue_with_onchange_callback/)
- url: https://www.reddit.com/r/rails/comments/er8432/rails_react_issue_with_onchange_callback/
---
I'm making a todo application to help me better understand React. Everything works with one exception. If a user types too quickly within an input field, the data gets out of sync. For example, if a user were to quickly type "123456" into the `input`, the final value would be "6".

I'm almost certain it's because by the time the `put` request is made and the `state` is updated, the user has already added more text to the `input`.

- [Video demonstrating the problem](https://share.getcloudapp.com/YEudz01k)
- [Application](https://rails-react-example.herokuapp.com/)
- [Source Code](https://github.com/stevepolitodesign/rails-react-example)
    - [Component in question](https://github.com/stevepolitodesign/rails-react-example/blob/master/app/javascript/packs/components/TodoItem.jsx#L86)
## [8][Best Rails CMS for multi-site/multi-tenant app?](https://www.reddit.com/r/rails/comments/er511o/best_rails_cms_for_multisitemultitenant_app/)
- url: https://www.reddit.com/r/rails/comments/er511o/best_rails_cms_for_multisitemultitenant_app/
---
I'm working with a client who wants to essentially white-label his current rental software site and offer it to similar businesses. I'm new to the CMS world (although I have 13 years of Rails experience) so I'm not sure what the best choice here is. What I don't want to do is have an install of this per business using it. I'd like to run them all on one app that can swap out assets/views/databases/etc based on URL. Seems like a multi-site with a multi-tenant database could be the right tool for the job. So:

* Are my assumptions correct so far?
* What's the best CMS for this?

I checked out Refinery but the multi-site libraries for it were lacking to say the least. Sofa looks like it might be better but it doesn't seem to have as big of a community. Any help would be much appreciated!!
## [9][I want to create an website like EngVid.com (website to watch videos from YouTube)](https://www.reddit.com/r/rails/comments/er8byw/i_want_to_create_an_website_like_engvidcom/)
- url: https://www.reddit.com/r/rails/comments/er8byw/i_want_to_create_an_website_like_engvidcom/
---
[Dont need to read these first paragraph ]

engvid.com 


Hi. I’m currently studying Ruby but I have read about Rails recently. This is my first attempt to learn how to built web application. I have done so far just a simple CRUDE application with MVC architecture nothing special in PHP without any framework. 

I’m very unfamiliar with web application since I have just code crud with simple “strings” information, so is quite hard to grasp how to get done this new personal project. 

First, I would like if know with is worth doing that with rails. You might say I even don’t need any of that to do a website like EngVid.com. but the thing is that I want to turn this project into a undergraduate thesis. I want to do basically the exactly same website to my country. Without doing that with any kind of plataform such as blogger or Wordpress. Build by myself all the pages the users can see. 



—Probelms

Is could bd difficult to implement these 3 features: commentaries to each page video, quiz system to each page video and tags system.



The main issue I would like to ask is how could I implement the Quiz system which they have to every video. I can’t grasp how I could get something like this done in my case. 
In the page of submitting a new video I could show to the Teacher User a series of fields to fill related to the Quiz. It would be something like 10 question, each one with 4 incorrect anwser and 1 correct.

Capture these input values it not a mystery to my. But where is going and how to deal with them it is a mystery to me. 
Display the capture value of inputs in a quiz-like format  is not the issue. It will be display with a title and radios input types. 
My concern is how to do the validation of correct and incorrect question.
## [10][Anyone successfully using Less (css) with a Rails 6 project?](https://www.reddit.com/r/rails/comments/er0t2a/anyone_successfully_using_less_css_with_a_rails_6/)
- url: https://www.reddit.com/r/rails/comments/er0t2a/anyone_successfully_using_less_css_with_a_rails_6/
---
I have a Rails 5 app using the less-rails gem I want to migrate to Rails 6. When I try to use less in Rails 6 I get an error, "ActionView::Template::Error (undefined method \`pathname' " I think it is the same error as reported here: 

[https://github.com/metaskills/less-rails/issues/152](https://github.com/metaskills/less-rails/issues/152)

&amp;#x200B;

Has anyone successfully used Less in a Rails 6 app? I was thinking it might be better to drop the less-rails gem and sprockets, and try to do it all with webpacker. I gave it a try but couldn't get it to work. I don't think I fully understand webpacker, and how the loader and enviornment files work. If anyone has a sample working config they could share, I'd appreciate it!
## [11][Submitting](https://www.reddit.com/r/rails/comments/er0p4l/submitting/)
- url: https://www.reddit.com/r/rails/comments/er0p4l/submitting/
---
This might be too html/webdev for this sub, but overall I've been so happy to not be hate-shamed here that I'm going to try  r/rails first!

Anyway, I have a form on a web app I'm working on to get comfortable with rails. The code for the edit page is below...

    &lt;%= form_for @gallery do |f| %&gt;
    
    ...
    
    &lt;table class="table table-condensed table-bordered table-striped"&gt;
                    &lt;tr&gt;
                        &lt;th&gt;&lt;strong&gt;Thumbnail&lt;/strong&gt;&lt;/th&gt;
                        &lt;th&gt;&lt;strong&gt;Filename&lt;/strong&gt;&lt;/th&gt;
                        &lt;th&gt;&lt;strong&gt;Added&lt;/strong&gt;&lt;/th&gt;
                        &lt;th&gt;&lt;strong&gt;Move Image&lt;/strong&gt;&lt;/th&gt;
                        &lt;th class="text-center"&gt;&lt;strong&gt;Delete&lt;/strong&gt;&lt;/th&gt;
                    &lt;/tr&gt;
                    &lt;% @gallery.images.each do |image| %&gt;
                        &lt;tr&gt;
                            &lt;td class="text-center"&gt;&lt;%= link_to( image_tag(image, alt: image.filename, class: "img-table"), image, target: "_blank") %&gt;&lt;/td&gt;
                            &lt;td&gt;&lt;%= image.filename %&gt;&lt;/td&gt;
                            &lt;td&gt;&lt;%= image.created_at %&gt;&lt;/td&gt;
                            &lt;td&gt;
                                &lt;div class="form-group"&gt;
                                    &lt;%= f.collection_select("id", Gallery.all, :id, :name, {prompt: true}, {class: "form-control"} ) %&gt;
                                &lt;/div&gt;
                            &lt;/td&gt;
                            &lt;td&gt;&lt;%= link_to "Delete", delete_image_attachment_gallery_path(image.signed_id, @gallery.id), method: :delete, data: { confirm: 'Are you sure?' }, class: "btn btn-block btn-sml btn-danger" %&gt;&lt;/td&gt;
                        &lt;/tr&gt;
                    &lt;% end %&gt;
                &lt;/table&gt;
    
    ...
    
    &lt;% end %&gt;

Anyway, the rendered page/form looks roughly like this...

[Interface Screenshot](http://i.imgur.com/nNE4sX6.png)

The Screenshot above is an example of an Gallery resource, which *has\_many\_attached: images* using the Active Storage gem and Amazon S3 (learning a TON about that stuff at the moment too!)

My problem is—as can be seen in the interface screenshotted above—I included a column to let the user move an image from one gallery to another. In the pic, I have changed first image's gallery to "Fashion", and my hope is that it would be moved in the #update action in gallery\_controller.

My general though process was that all of these select options would be submitted in an array I could loop through, and as I proceeded I could check the values against the current gallery name, and move each image as needed. I'm not sure yet how exactly to move them, but that's not important now. Anyway, I'm almost embarrassed to admit I have no idea where this hypothetical array of selectbox options lives. I checked in params\[\] and found only the following:

    Parameters: {"utf8"=&gt;"✓", "authenticity_token"=&gt;"ZVNF4+pzodvyu4KY9jAXHdkMF9XOCweowaqBgc1+lHSpioNer9DcxbFWo3AHMjANAmaFqtE10c4NKcK2d/HEHA==", "gallery"=&gt;{"name"=&gt;"Weddings", "description"=&gt;"Weddings", "id"=&gt;"11", "visible"=&gt;"0"}, "button"=&gt;"", "id"=&gt;"11"}

Anyway, as usual I'm sure I could figure it out, but I'm struggling to think of how to start my search! Any help at all would be profoundly appreciated!

Also let me know if you'd like to see anything else!
