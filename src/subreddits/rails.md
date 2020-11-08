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
## [3][How to add hard_wrap: true in a custom Markdown Redcarpet](https://www.reddit.com/r/rails/comments/jq8upc/how_to_add_hard_wrap_true_in_a_custom_markdown/)
- url: https://www.reddit.com/r/rails/comments/jq8upc/how_to_add_hard_wrap_true_in_a_custom_markdown/
---
In my application\_helper I have this

    def parse_markdown(text)
     markdown = Redcarpet::Markdown.new(MarkdownRenderer, hard_wrap: true, autolink: true, space_after_headers: true)
     markdown.render(text)
    end 

But probably to add **hard\_wrap: true** is not the right thing to do.

We are using a "*custom markdown*". So in **facedes/markdown\_render.rb** I have this

    class MarkdownRenderer &lt; Redcarpet::Render::HTML
     include Rails.application.routes.url_helpers
     include ActionView::Helpers::UrlHelper
        def paragraph(text)
         "#{text}&lt;br&gt;"
        end
       (and a lot of other things) 

But in this way if I write a comment with

    line one
    line two 

I see

    line one line two 

and only if I use the "doubble", in this way

    line one
    
    line two 

I see

    line one
    line two 

**How to solve?**

I want to write

    line one
    line two
    
    line three 

and to see

    line one
    line two
    line three

I also try to add in 

     def initialize(options={})
       super options.merge(:hard_wrap =&gt; true)
     end 

But it doesn't work.
## [4][Getting a syntax error, but I can't figure out why](https://www.reddit.com/r/rails/comments/jq6s9h/getting_a_syntax_error_but_i_cant_figure_out_why/)
- url: https://www.reddit.com/r/rails/comments/jq6s9h/getting_a_syntax_error_but_i_cant_figure_out_why/
---
https://imgur.com/a/0far2GQ

I am following a twitter clone tutorial and I keep getting a syntax error when trying to render a partial form. I can't figure out what I am doing wrong. Could you please take a look?
## [5][Add library css in Ruby on Rails 6 with webpacker](https://www.reddit.com/r/rails/comments/jqb6u3/add_library_css_in_ruby_on_rails_6_with_webpacker/)
- url: https://www.reddit.com/r/rails/comments/jqb6u3/add_library_css_in_ruby_on_rails_6_with_webpacker/
---
I'm trying to install a library, specifically bootstrap-table via yarn. I had tried the gem found at the following link [Click](https://github.com/bjevanchiu/bootstrap-table-rails) but it is an old version and I was unable to get it to work. I have read around that however with Rails 6 it is better to use the webpacker to install libraries. My problem now is I can't figure out how to insert bootstrap-table css.

Here are my files:

application.js

    import 'bootstrap'
    require("@rails/ujs").start()
    require("turbolinks").start()
    require("@rails/activestorage").start()
    require("channels")
    require("jquery")
    require("jquery-ui")
    import "@fortawesome/fontawesome-free/js/all";
    require("leaflet");
    require("easy-autocomplete")
    
    require bootstrap-table


application.scss

     *= require jquery-ui
     *= require_tree .
     *= require_self
    
    
    
     */
    
    @import "leaflet/dist/leaflet";
    @import "bootstrap/scss/bootstrap";
    @import '@fortawesome/fontawesome-free/css/fontawesome.css';
    @import '@fortawesome/fontawesome-free/css/all.css';
    @import 'easy-autocomplete/dist/easy-autocomplete';


In app/views/layouts/application.htlm.erb I have:

    &lt;!DOCTYPE html&gt;
    &lt;html&gt;
      &lt;head&gt;
        &lt;title&gt;Resto&lt;/title&gt;
        &lt;%= csrf_meta_tags %&gt;
        &lt;%= csp_meta_tag %&gt;
    
        &lt;%= stylesheet_link_tag 'application', media: 'all', 'data-turbolinks-track': 'reload' %&gt;
        &lt;!-- This refers to app/javascript/stylesheets/application.scss--&gt;
        &lt;%= javascript_pack_tag 'application', 'data-turbolinks-track': 'reload' %&gt;
    
    
      &lt;/head&gt;
    
      &lt;body&gt;
        &lt;p class = "notice" &gt;&lt;%= notice %&gt; &lt;/p&gt;
        &lt;p class = "alert" &gt;&lt;%= alert %&gt;&lt;/p&gt;
        &lt;%= yield %&gt;
    
      &lt;/body&gt;
    &lt;/html&gt;


I read on the internet that to insert the css of the libraries installed by yarn I have to paste the following code in application.html.erb file:

    &lt;% = stylesheet_pack_tag 'application', 'data-turbolinks-track': 'reload'%&gt;


So what I did was put this code in the file and I edited `config/webpacker.yml` to be sure it has `extract_css: true` in the default section at the top.


I also started the `bin / webpack-dev-server` command in the console but I got this error:



    You want to set webpacker.yml value of compile to true for your environment
    
    unless you are using the webpack -w or the webpack-dev-server.
    
    webpack has not yet re-run to reflect updates.
    
    You have misconfigured Webpacker's config/webpacker.yml file.
    
    Your webpack configuration is not creating a manifest.
    
    Your manifest contains:
    
    {
      "application.js": "/packs/js/application-105ced549e0eccca7ef2.js",
      "application.js.map": "/packs/js/application-105ced549e0eccca7ef2.js.map",
      "entrypoints": {
        "application": {
          "js": [
            "/packs/js/application-105ced549e0eccca7ef2.js"
          ],
          "js.map": [
            "/packs/js/application-105ced549e0eccca7ef2.js.map"
          ]
        }
      }
    }



I would like not to have to use all this configuration but to just add the bootstrap-table css file to my page. I don't feel like good in running this `bin / webpack-dev-server` command. This is because I had a lot of errors after running this command and I would like to avoid having to keep it running all the time. It's possible? If that's not possible, what am I doing wrong?
## [6][General models or multiple specific ones?](https://www.reddit.com/r/rails/comments/jq4vcx/general_models_or_multiple_specific_ones/)
- url: https://www.reddit.com/r/rails/comments/jq4vcx/general_models_or_multiple_specific_ones/
---
I am building an inventory tracker with some extra features as a project and wondering in the scenario  I am having whether to use 2 models (ProductVariantTitle and ProductVariantTag) and allow the user to add their own Tagging Title and Description/Tag like Size, Color, Category, etc. Or do more specified models that are more likely to be required like "Category" or "Vendor"?

&amp;#x200B;

Assuming there is a Product Model and ProductVariant Model in both cases... (Product model is more of a title and generalized information, the ProuctVariant is the specific SKU --- no matter if there are no variants, there will always be a ProductVariant created within the app)

For instance, I can maybe do one less specified model structure like this:

&amp;#x200B;

**ProductVariantTitle**

⦁	Name

Examples:

NAME is Size

NAME is Color

NAME is Category

NAME is Vendor

NAME is Collection

&amp;#x200B;

**ProductVariantTag**

⦁	Name

⦁	ProductVariantTitleId

⦁	ProductVariantId

Examples:

NAME is Small

NAME is Black

NAME is T-Shirt

NAME is Bob

NAME is Summer

&amp;#x200B;

Or should I be doing something more specified like:

**Vendor**

* Name

**ProductVariantVendor**  (join)

* VendorId
* ProductVariantId

&amp;#x200B;

**Collection**

* Name

**ProductVariantCollection**  (join)

* CollectionId
* ProductVariantId

&amp;#x200B;

**Category**

* Name

**ProductVariantCategory** (join)

* CategoryId
* ProductVariantId

&amp;#x200B;

as models?

&amp;#x200B;

But maybe at the same time, keep the **ProductVariantTitle and ProductVariantTag** for more dynamic titles like **"Size", "Style", "Color" etc.?** Because not every product will always have a color specified, such as a battery or food.  But Category and Vendor are characteristics that will always be required for any product.

&amp;#x200B;

Both seem to make sense in their own worlds but I am wondering if there are any opinions on either?
## [7][How to incorporate an application within a Rails application?](https://www.reddit.com/r/rails/comments/jpyk29/how_to_incorporate_an_application_within_a_rails/)
- url: https://www.reddit.com/r/rails/comments/jpyk29/how_to_incorporate_an_application_within_a_rails/
---
Hello Rails community!

Just a quick question.

Creating a Rails app and serving data is pretty straight-forward for me. But I wanted to know how I could use data from my Rails app with another Python-based application.

I was thinking of having a shared database and using it as a shared place where both applications can utilize each others data.

I think I have an idea, but was wondering if I was on the right track.

Thanks for any help!
## [8][Bootstrap-table-rails](https://www.reddit.com/r/rails/comments/jpnu5h/bootstraptablerails/)
- url: https://www.reddit.com/r/rails/comments/jpnu5h/bootstraptablerails/
---
hi, I'm trying to install this gem: https://github.com/bjevanchiu/bootstrap-table-rails in rails 6. The problem is that when I put it in the gemfile and run the bundle install command, it doesn't appear in node_modules. With linux find I looked for the bootstrap-table-rails directory but I can't find it anywhere. How can I integrate Bootstrap-table into rails? Thanks so much for the future help
## [9][Securing Rails API access - what method?](https://www.reddit.com/r/rails/comments/jpgg3d/securing_rails_api_access_what_method/)
- url: https://www.reddit.com/r/rails/comments/jpgg3d/securing_rails_api_access_what_method/
---
Imagine your typical Rails API scenario: you're building an API but you want to restrict the access to it (also to have opportunity to see how many requests a certain consumer of the API made for statistical purposes).

One way which seems simple enough and is often used in tutorials on the net is securing it using API keys which are tied to the other Rails apps consuming this new fancy API. 

My question is: where do you put the list of generated API keys? Inside the API itself in some table (e.g. approved_clients)? Or do you use some other solution?

I'm trying to brainstorm here which approach I should take when building an API shared between several Rails-based apps.

Using oAuth makes no sense here to me because the said API would not be user-centric but Rails app (client) centric: e.g. you'd ask the shared geo API to return you a list of cities close to given geo coordinates. So we don't need to login individual users here to do something on them, just to have a shared services that's generic enough no to involve oAuth or other more complicated flows. 

Or could JWTs somehow fit here?

Thanks for your ideas.
## [10][I created an Xbox backward compatible game catalogue with Rails](https://www.reddit.com/r/rails/comments/jpc1tb/i_created_an_xbox_backward_compatible_game/)
- url: https://www.reddit.com/r/rails/comments/jpc1tb/i_created_an_xbox_backward_compatible_game/
---
Hi there,

I've recently purchased an Xbox Series S and excited about the games from the previous generation as I didn't have much time to dive in. I created [this small web app](https://backwardible.com/) with Rails and thought it might help other people too.

Oh, I tend to write blog posts about the side-projects I do, so if anyone's interested in terms of what I used and why I used, the concerning post is [here](https://arkan.me/backwardible/) as well!

I would love to hear your feedback.
## [11][Confused about pulling changes without overwriting local env](https://www.reddit.com/r/rails/comments/jpgc5y/confused_about_pulling_changes_without/)
- url: https://www.reddit.com/r/rails/comments/jpgc5y/confused_about_pulling_changes_without/
---
Can somebody explain with a TL;DR on this? I'm trying to figure out the following scenario:

\- I clone a repo and get it up and running with bundle install, setup local db, create env file, compile webpacker, and so on. This goes fine.

\- A dev makes changes to that remote repo which I want to get on my local machine. What do I do now? How do I get those changes onto my machine without overwriting my env file or any other setup I've done locally?
## [12][Best React resource for Rails developers](https://www.reddit.com/r/rails/comments/jpb17u/best_react_resource_for_rails_developers/)
- url: https://www.reddit.com/r/rails/comments/jpb17u/best_react_resource_for_rails_developers/
---
hey friends, I'm a purely backend developer and I barely played with React so far. What are the best resources that you recommend for me to go over to get up to speed when it comes to React on top of Rails.
