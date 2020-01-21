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
## [2][Issue with Code Mirror and Heroku](https://www.reddit.com/r/rails/comments/erpm2k/issue_with_code_mirror_and_heroku/)
- url: https://www.reddit.com/r/rails/comments/erpm2k/issue_with_code_mirror_and_heroku/
---
Hello Everyone,

I recently added the code mirror editor to my rails (6.0.0) app. I was able to get it working just fine in my local environment (Ubuntu 18.04), however when I deploy it to heroku the editors actual input area is very badly offset and has random x's in it.

Here is my application.js file:

    require("@rails/ujs").start()
    require("turbolinks").start()
    require("@rails/activestorage").start()
    require("channels")
    require("jquery")
    import { autocompleteExport } from '../custom/autocomplete';
    import { initialize } from '../custom/editor';
    import { formatToc } from '../custom/page_display';
    
    window.jQuery = $;
    window.$ = $;
    
    $(document).on("turbolinks:load", autocompleteExport.categories)
    $(document).on("turbolinks:load", autocompleteExport.search)
    $(document).on("turbolinks:load", autocompleteExport.admins)
    $(document).on("turbolinks:load", initialize)
    $(document).on("turbolinks:load", formatToc)

&amp;#x200B;

The relevent editor.js:

    import CodeMirror from 'codemirror/lib/codemirror.js'
    import 'codemirror/lib/codemirror.css'
    import 'codemirror/mode/markdown/markdown.js'
    
    function initialize(){
        let textArea = document.getElementById('page_edit_content')
        let editor = document.getElementById('content-editor')
        
        if(textArea &amp;&amp; !editor){
            var contentEditor = CodeMirror.fromTextArea(textArea, {
                                                        lineWrapping: true,
                                                        mode: "markdown",
                                                    });
    
            contentEditor.display.wrapper.id = "content-editor"
        }
    
        textArea = null
        editor = null
    
        textArea = document.getElementById('page_edit_summary')
        editor = document.getElementById("summary-editor")
    
        if(textArea &amp;&amp; !editor){
            var contentEditor = CodeMirror.fromTextArea(textArea, {
                lineWrapping: true,
                mode: "markdown",
            });
    
            contentEditor.display.wrapper.id = "summary-editor"
        }
    
        textArea = null
        editor = null
    
        textArea = document.getElementById('page_content')
        editor = document.getElementById("new-content-editor")
    
        if(textArea &amp;&amp; !editor){
            var contentEditor = CodeMirror.fromTextArea(textArea, {
                lineWrapping: true,
                mode: "markdown",
            });
    
            contentEditor.display.wrapper.id = "new-content-editor"
        }
    
        textArea = null
        editor = null
    
        textArea = document.getElementById('page_summary')
        editor = document.getElementById("new-summary-editor")
    
        if(textArea &amp;&amp; !editor){
            var contentEditor = CodeMirror.fromTextArea(textArea, {
                lineWrapping: true,
                mode: "markdown",
            });
    
            contentEditor.display.wrapper.id = "new-summary-editor"
        }
    }
    
    export {initialize}

and lastly one of the views that isn't working

    &lt;% @page_title = "New Page" %&gt;
    
    &lt;div class="container form-container"&gt;
      &lt;%= form_for @page, url: world_pages_path(params[:world_name]), html: {style: "width: 100%"} do |f| %&gt;
        &lt;%= render 'shared/error_messages', errors: flash[:errors] %&gt; 
        &lt;div class="form-group"&gt;
          &lt;%= f.label :title %&gt;
          &lt;%= f.text_field :title, autofocus: true, class: "form-control" %&gt;
        &lt;/div&gt;
        &lt;div class="form-group"&gt;
          &lt;%= f.label :summary %&gt;
          &lt;%= f.text_area :summary, class: "form-control" %&gt;
        &lt;/div&gt;
        &lt;div class="form-group"&gt;
          &lt;%= f.label :content %&gt;
          &lt;%= f.text_area :content, class: "form-control" %&gt;
        &lt;/div&gt;
        &lt;div class="form-group"&gt;
          &lt;%= f.submit "Create!", class: "btn btn-primary" %&gt;
        &lt;/div&gt;
        &lt;% if params[:category] %&gt;     
        &lt;%= hidden_field_tag :category, params[:category] %&gt;
        &lt;% end %&gt;
      &lt;% end %&gt;
    &lt;/div&gt;

&amp;#x200B;

[Those x's at the top are actually part of the top editor, and the ones way at the bottom of the main editor you can see.](https://preview.redd.it/hp2q249g92c41.png?width=3354&amp;format=png&amp;auto=webp&amp;s=35ef6547cbf75ba627ad950b7f8538a2993e5596)
## [3][How do you track or list online users sin rails?](https://www.reddit.com/r/rails/comments/erkddk/how_do_you_track_or_list_online_users_sin_rails/)
- url: https://www.reddit.com/r/rails/comments/erkddk/how_do_you_track_or_list_online_users_sin_rails/
---

## [4][Trouble getting javascript to function!](https://www.reddit.com/r/rails/comments/erj6td/trouble_getting_javascript_to_function/)
- url: https://www.reddit.com/r/rails/comments/erj6td/trouble_getting_javascript_to_function/
---
Hey guys, I have an in-process webapp deployed at [here](https://safe-mesa-78631.herokuapp.com/)

Please forgive the awful stand-in copy I wrote. The issue is with the top nav bar, and in particular the "Galleries" drop-down. It seems like SOMETIMES the dropdown work, and sometimes (especially after clicking a dropped-down link) it doesn't. Whenever it doesn't work, I can hit the down-arrow on my keyboard to get it to drop down, but other than that I'm just not sure why it's not working.

I'm using the bootstrap framework (obviously), if that matters.

[Here](https://github.com/Chucksef/photo_app) is a link to the github which will show just about everything you might need, though of course I can answer any questions. I honestly just don't know why it would work sometimes and not others and it's driving me INSANE!!! So any help at all would be incredibly appreciated!
## [5][`bundle install` commands not working, error not clear](https://www.reddit.com/r/rails/comments/erf1re/bundle_install_commands_not_working_error_not/)
- url: https://www.reddit.com/r/rails/comments/erf1re/bundle_install_commands_not_working_error_not/
---
I did a `brew install imagemagick` for image manipulation, and went to update my gemfile with `image_processing` and `mini_magic` gems, and then did a `bundle install` and it errors out.

I tried `brew update` and `brew upgrade`, and they ran but still getting an error on `bundle install`, see the below:


    Mac-Users-Apple-Computer:sample_app_v6 z$ bundle
    Traceback (most recent call last):
      31: from /Users/z/.rbenv/versions/2.6.3/bin/bundle:23:in `&lt;main&gt;'
      30: from /Users/z/.rbenv/versions/2.6.3/bin/bundle:23:in `load'
      29: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/gems/2.6.0/gems/bundler-2.1.4/exe/bundle:34:in `&lt;top (required)&gt;'
      28: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/gems/2.6.0/gems/bundler-2.1.4/lib/bundler/friendly_errors.rb:123:in `with_friendly_errors'
      27: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/gems/2.6.0/gems/bundler-2.1.4/exe/bundle:46:in `block in &lt;top (required)&gt;'
      26: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/gems/2.6.0/gems/bundler-2.1.4/lib/bundler/cli.rb:24:in `start'
      25: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/gems/2.6.0/gems/bundler-2.1.4/lib/bundler/vendor/thor/lib/thor/base.rb:476:in `start'
      24: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/gems/2.6.0/gems/bundler-2.1.4/lib/bundler/cli.rb:30:in `dispatch'
      23: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/gems/2.6.0/gems/bundler-2.1.4/lib/bundler/vendor/thor/lib/thor.rb:399:in `dispatch'
      22: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/gems/2.6.0/gems/bundler-2.1.4/lib/bundler/vendor/thor/lib/thor/invocation.rb:127:in `invoke_command'
      21: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/gems/2.6.0/gems/bundler-2.1.4/lib/bundler/vendor/thor/lib/thor/command.rb:27:in `run'
      20: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/gems/2.6.0/gems/bundler-2.1.4/lib/bundler/cli.rb:255:in `install'
      19: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/gems/2.6.0/gems/bundler-2.1.4/lib/bundler/settings.rb:124:in `temporary'
      18: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/gems/2.6.0/gems/bundler-2.1.4/lib/bundler/cli.rb:256:in `block in install'
      17: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/gems/2.6.0/gems/bundler-2.1.4/lib/bundler/cli/install.rb:66:in `run'
      16: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/2.6.0/rubygems/core_ext/kernel_require.rb:54:in `require'
      15: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/2.6.0/rubygems/core_ext/kernel_require.rb:54:in `require'
      14: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/gems/2.6.0/gems/bundler-2.1.4/lib/bundler/installer.rb:4:in `&lt;top (required)&gt;'
      13: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/2.6.0/rubygems/core_ext/kernel_require.rb:54:in `require'
      12: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/2.6.0/rubygems/core_ext/kernel_require.rb:54:in `require'
      11: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/2.6.0/rubygems/dependency_installer.rb:4:in `&lt;top (required)&gt;'
      10: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/2.6.0/rubygems/core_ext/kernel_require.rb:54:in `require'
      9: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/2.6.0/rubygems/core_ext/kernel_require.rb:54:in `require'
      8: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/2.6.0/rubygems/package.rb:44:in `&lt;top (required)&gt;'
      7: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/2.6.0/rubygems/core_ext/kernel_require.rb:54:in `require'
      6: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/2.6.0/rubygems/core_ext/kernel_require.rb:54:in `require'
      5: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/2.6.0/rubygems/security.rb:12:in `&lt;top (required)&gt;'
      4: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/2.6.0/rubygems/core_ext/kernel_require.rb:54:in `require'
      3: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/2.6.0/rubygems/core_ext/kernel_require.rb:54:in `require'
      2: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/2.6.0/openssl.rb:13:in `&lt;top (required)&gt;'
      1: from /Users/z/.rbenv/versions/2.6.3/lib/ruby/2.6.0/rubygems/core_ext/kernel_require.rb:54:in `require'
    /Users/z/.rbenv/versions/2.6.3/lib/ruby/2.6.0/rubygems/core_ext/kernel_require.rb:54:in `require': dlopen(/Users/z/.rbenv/versions/2.6.3/lib/ruby/2.6.0/x86_64-darwin18/openssl.bundle, 9): Library not loaded: /usr/local/opt/openssl/lib/libssl.1.0.0.dylib (LoadError)
      Referenced from: /Users/z/.rbenv/versions/2.6.3/lib/ruby/2.6.0/x86_64-darwin18/openssl.bundle
      Reason: image not found - /Users/z/.rbenv/versions/2.6.3/lib/ruby/2.6.0/x86_64-darwin18/openssl.bundle

I've done some google searching but it isn't clear


-----

EDIT: I did see an SO answer saying to do the below (I'm on mac of course), `brew remove openssl; brew install openssl`, but get an error so am not sure if this is a safe route to try:

https://stackoverflow.com/questions/39051887/image-not-found-in-openssl-bundle-when-install-cocoapods-on-mac?noredirect=1&amp;lq=1

    Mac-Users-Apple-Computer:sample_app_v6 z$ brew remove openssl
    Error: Refusing to uninstall /usr/local/Cellar/openssl@1.1/1.1.1d
    because it is required by glib, httpie, imagemagick, krb5, libheif, mysql, postgresql, python, python@3.8, rbenv and shared-mime-info, which are currently installed.
    You can override this and force removal with:
      brew uninstall --ignore-dependencies openssl

-----

EDIT 2: RESOLVED

moomaka had me uninstall and reinstall my particular version of ruby and eventually all is well.  Details in comments below.
## [6][ActiveAdmin Scoped Collection Actions](https://www.reddit.com/r/rails/comments/ere9m2/activeadmin_scoped_collection_actions/)
- url: https://www.reddit.com/r/rails/comments/ere9m2/activeadmin_scoped_collection_actions/
---
Hey,

According to the gem authors: "...By default we don't allow perform actions on **all** the records. We want protect you from accidental deleting...". Well, I'd like to be able to bulk select a bunch of records (or even all records) using the selectable column in the index page of a particular resource and perform an action on those records via scoped AA collection actions...

How can I override this disabled setting?

Thanks.
## [7][Is Facebook lite a good portfolio project?](https://www.reddit.com/r/rails/comments/erajnl/is_facebook_lite_a_good_portfolio_project/)
- url: https://www.reddit.com/r/rails/comments/erajnl/is_facebook_lite_a_good_portfolio_project/
---
This is the last of many rails projects in the Odin project (an opensource curriculum that takes over a year to complete)

[https://www.theodinproject.com/courses/ruby-on-rails/lessons/final-project](https://www.theodinproject.com/courses/ruby-on-rails/lessons/final-project)

Is this actually going to impress any employers?

A lot of you have told me in the past to contribute to Open source and that has been NOTED but how does this go for a portfolio project?
## [8][Example project with Vue + Rails + SPA](https://www.reddit.com/r/rails/comments/erad9q/example_project_with_vue_rails_spa/)
- url: https://www.reddit.com/r/rails/comments/erad9q/example_project_with_vue_rails_spa/
---
Looking for a guide or a repo that has vuejs on a rails project, and it's an SPA. (will gold lol)
## [9][Is it possible to use the GraphiQL Playground in a API-only Rails app?](https://www.reddit.com/r/rails/comments/er7o5h/is_it_possible_to_use_the_graphiql_playground_in/)
- url: https://www.reddit.com/r/rails/comments/er7o5h/is_it_possible_to_use_the_graphiql_playground_in/
---
I have an API-only Rails app that I am adding GraphQL to. I am going through this GraphQL tutorial: [https://www.howtographql.com/graphql-ruby/2-queries/](https://www.howtographql.com/graphql-ruby/2-queries/) , and they recommend using GraphQL's built-in Playground interface, which looks awesome. However, because my Rails app is API-only, it can't be built, I think because it isn't setup to display templates. Does anyone know if there is a way to get around this?
## [10][Webpacker, how to include a JS library's image path?](https://www.reddit.com/r/rails/comments/erag0g/webpacker_how_to_include_a_js_librarys_image_path/)
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
## [11][Rails + React issue with onChange callback](https://www.reddit.com/r/rails/comments/er8432/rails_react_issue_with_onchange_callback/)
- url: https://www.reddit.com/r/rails/comments/er8432/rails_react_issue_with_onchange_callback/
---
I'm making a todo application to help me better understand React. Everything works with one exception. If a user types too quickly within an input field, the data gets out of sync. For example, if a user were to quickly type "123456" into the `input`, the final value would be "6".

I'm almost certain it's because by the time the `put` request is made and the `state` is updated, the user has already added more text to the `input`.

- [Video demonstrating the problem](https://share.getcloudapp.com/YEudz01k)
- [Application](https://rails-react-example.herokuapp.com/)
- [Source Code](https://github.com/stevepolitodesign/rails-react-example)
    - [Component in question](https://github.com/stevepolitodesign/rails-react-example/blob/master/app/javascript/packs/components/TodoItem.jsx#L86)
