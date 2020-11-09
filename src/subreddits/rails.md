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
## [3][Is there a convention for the controller/view of the root route?](https://www.reddit.com/r/rails/comments/jqn37p/is_there_a_convention_for_the_controllerview_of/)
- url: https://www.reddit.com/r/rails/comments/jqn37p/is_there_a_convention_for_the_controllerview_of/
---
I keep thinking I should call it main or default or root or maybe application.   
I plan to have it check the whether a session exists and redirect to a main page otherwise show some unauthenticated information. So in that case should it be called welcome?  
... Thoughts?
## [4][I want to build a CMS for Photos. But I want the photos saved on a CDN like Amazoin CLoudfront, how do I go about this?](https://www.reddit.com/r/rails/comments/jqc2tz/i_want_to_build_a_cms_for_photos_but_i_want_the/)
- url: https://www.reddit.com/r/rails/comments/jqc2tz/i_want_to_build_a_cms_for_photos_but_i_want_the/
---
What CDN would you suggest?

And if anyone has done this before, could they share a link to their source code if it was open source.   
What pitfalls should I avoid?
## [5][Permitted attribute but still Cannot mass assign warning](https://www.reddit.com/r/rails/comments/jqrb8v/permitted_attribute_but_still_cannot_mass_assign/)
- url: https://www.reddit.com/r/rails/comments/jqrb8v/permitted_attribute_but_still_cannot_mass_assign/
---
I wrap the params in my controller so I get something like this.

    params = { "name": "user_name", "user": { "name": "user_name" } }

and in the application controller I permit these like so

     params["user"].permit(:name)

But durinng .new like so 

    User.new(params[:user]))

I get a warning like

 

Can't mass-assign protected attributes for User: name

&amp;#x200B;

I do not have name in attr\_accessible since I belive it needs to be handled in the controller level. But I'm not sure why I'm getting this cannot mass-assign exception
## [6][Rails custom domains for users on heroku](https://www.reddit.com/r/rails/comments/jqhvaa/rails_custom_domains_for_users_on_heroku/)
- url: https://www.reddit.com/r/rails/comments/jqhvaa/rails_custom_domains_for_users_on_heroku/
---
Really appreciate the community help here. I'm working on an app that exists on Heroku. The app has users through devise and those users should be able to add their own domain name to the app. [www.app.com/user1](https://www.app.com/user1) should be accessible from [user1.com](https://user1.com) if they have a domain set up. I've been researching like crazy and I don't have a solid solution yet. If anyone has achieved something similar I would love to know how. Thank you family! If the solution requires a paid third party service like Cloudflare that is acceptable.
## [7][How to add hard_wrap: true in a custom Markdown Redcarpet](https://www.reddit.com/r/rails/comments/jq8upc/how_to_add_hard_wrap_true_in_a_custom_markdown/)
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
## [8][Getting a syntax error, but I can't figure out why](https://www.reddit.com/r/rails/comments/jq6s9h/getting_a_syntax_error_but_i_cant_figure_out_why/)
- url: https://www.reddit.com/r/rails/comments/jq6s9h/getting_a_syntax_error_but_i_cant_figure_out_why/
---
https://imgur.com/a/0far2GQ

I am following a twitter clone tutorial and I keep getting a syntax error when trying to render a partial form. I can't figure out what I am doing wrong. Could you please take a look?
## [9][Add library css in Ruby on Rails 6 with webpacker](https://www.reddit.com/r/rails/comments/jqb6u3/add_library_css_in_ruby_on_rails_6_with_webpacker/)
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
## [10][General models or multiple specific ones?](https://www.reddit.com/r/rails/comments/jq4vcx/general_models_or_multiple_specific_ones/)
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
## [11][How to incorporate an application within a Rails application?](https://www.reddit.com/r/rails/comments/jpyk29/how_to_incorporate_an_application_within_a_rails/)
- url: https://www.reddit.com/r/rails/comments/jpyk29/how_to_incorporate_an_application_within_a_rails/
---
Hello Rails community!

Just a quick question.

Creating a Rails app and serving data is pretty straight-forward for me. But I wanted to know how I could use data from my Rails app with another Python-based application.

I was thinking of having a shared database and using it as a shared place where both applications can utilize each others data.

I think I have an idea, but was wondering if I was on the right track.

Thanks for any help!
## [12][Bootstrap-table-rails](https://www.reddit.com/r/rails/comments/jpnu5h/bootstraptablerails/)
- url: https://www.reddit.com/r/rails/comments/jpnu5h/bootstraptablerails/
---
hi, I'm trying to install this gem: https://github.com/bjevanchiu/bootstrap-table-rails in rails 6. The problem is that when I put it in the gemfile and run the bundle install command, it doesn't appear in node_modules. With linux find I looked for the bootstrap-table-rails directory but I can't find it anywhere. How can I integrate Bootstrap-table into rails? Thanks so much for the future help
