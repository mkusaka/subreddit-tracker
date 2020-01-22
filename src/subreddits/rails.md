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
## [2][Build a Reddit Clone in Rails 6 [New Video Series]](https://www.reddit.com/r/rails/comments/es6m26/build_a_reddit_clone_in_rails_6_new_video_series/)
- url: https://www.reddit.com/r/rails/comments/es6m26/build_a_reddit_clone_in_rails_6_new_video_series/
---
Hi guys, I'm back with a new Ruby on Rails build. This time we are building a lightweight Reddit clone. I've shared some previous Rails builds here and thought this might be on interest to anyone who enjoyed those past videos. Hope you guys enjoy the video.

**Part 1:**

In this video I cover the setup of a new rails app, installing the devise gem and creating user accounts. Add functionality for users to create a new community/subreddit. Add functionality to create new posts for a community (using nested routes).

[https://www.youtube.com/watch?v=aD6JvHKNPPM](https://www.youtube.com/watch?v=aD6JvHKNPPM)

**Part 2 (Coming soon):**

Signed in users can subscribe to communities and create new posts within them. Add functionality to allow subscribers to upvote / downvote posts and order them in community based on popularity. Add comments to posts.

&amp;#x200B;

The build is part of a collection of MVP's that I have been building after working in the tech startup space for some time and assisting in building minimum viable products.

I post new web app builds on YouTube using Ruby on Rails 6 so feel free to check out the channel and add some recommendations for upcoming builds.
## [3][API + Frontend - Should I use webpack or isolate them?](https://www.reddit.com/r/rails/comments/esb7fx/api_frontend_should_i_use_webpack_or_isolate_them/)
- url: https://www.reddit.com/r/rails/comments/esb7fx/api_frontend_should_i_use_webpack_or_isolate_them/
---
I usually build monolithic rails apps, often with a React or Vue SPA frontend.

I am about to start developping a simple app with a RoR API. I would like to use React or Vue for the frontend, and I plan to develop a mobile app that will access the same API, but later, when I have more time to code.

Also, since this is a personal project I want to build to showcase my dev skills, I intend to develop the API and the frontend with several frameworks/languages (React, Crystal, Elixir, React, Vue, Svelte, Webcomponents) and compare them.

Would you advise me to create two separate projects/git repos (API + frontend) or should I go for webpack and bundle the frontend and the API together? Should I use git submodules?
## [4][Authlogic::Session::Existence::SessionInvalidError](https://www.reddit.com/r/rails/comments/es30kk/authlogicsessionexistencesessioninvaliderror/)
- url: https://www.reddit.com/r/rails/comments/es30kk/authlogicsessionexistencesessioninvaliderror/
---
I'm running minitests and I'm getting this random error 

 Authlogic::Session::Existence::SessionInvalidError:  Authlogic::Session::Existence::SessionInvalidError: Your session is invalid and has the following errors: Activate &amp; Log in

This doesn't appear for the say first 10 minitests and subsequently, all the later ones have this error. 

I'm not sure where to look for since this error started to appear suddenly. Any help would be great
## [5][Rails and Apache (Phusion Passenger): how does the Apacche vhost DocumentRoot fit into routing?](https://www.reddit.com/r/rails/comments/es3r9w/rails_and_apache_phusion_passenger_how_does_the/)
- url: https://www.reddit.com/r/rails/comments/es3r9w/rails_and_apache_phusion_passenger_how_does_the/
---
I understand that I need to set up a virtualhost setting DocumentRoot to the Public/ directory at the same level as the app/ and other directories.  I'm trying to understand how the handover from Apache / httpd and rails is handled.

Any links would be appreciated.  Does Apache handover connections to Passenger or what is going on?
## [6][Creating a join table relation on create of has_many model?](https://www.reddit.com/r/rails/comments/es0ywu/creating_a_join_table_relation_on_create_of_has/)
- url: https://www.reddit.com/r/rails/comments/es0ywu/creating_a_join_table_relation_on_create_of_has/
---
I'm trying to recall how to create a join table relationship on create or update. For example, I have 

\`Outage\` has\_many \`services\` through \`service\_outages\` 

&amp;#x200B;

On create of an \`outage\`, I need to also create a join table relationship to \`services\` through \`service\_outages\`.

&amp;#x200B;

Here's the code: 

&amp;#x200B;

    class OutagesController &lt; ApplicationController
        before_action :find_outage, only: [:show, :update, :destroy]
    
        def index
            @outages = Outage.all
            render :json =&gt; @outages, status: @ok 
        end
    
        def show
            render :json =&gt; @outage, status: @ok
        end
    
        def create 
            @outage = Outage.create!(outage_params)  
            if @outage.save
                render :json =&gt; @outage, status: @ok 
            else         
                render :json =&gt; { errors: @outage.errors.full_messages }, status: @unprocessible_entity
            end
        end
    
        def update 
            @outage.update!(outage_params)
            if @outage.save
                render :json =&gt; @outage, status: @ok 
            else         
                render :json =&gt; { errors: @outage.errors.full_messages }, status: @unprocessible_entity
            end
        end
    
        def destroy 
            @outage.delete
            if @outage.delete
                render :json =&gt; "Outage instance destroyed"
            else
                render :json =&gt; { errors: "Unable to delete" }
            end
        end
    
    
        private 
    
        def find_outage
            @outage = Outage.find(params[:id])
        end
    
        def outage_params
            byebug 
            params.permit(:start_time, :end_time, :outage_type, :frequency, :reason)
        end
    end
    

I pass \`service.id\` in params from the front end on create. Do I create this in the \`create\` method here?
## [7][Need some help with Model/Schema/Relationship for a Rails API](https://www.reddit.com/r/rails/comments/es6xrb/need_some_help_with_modelschemarelationship_for_a/)
- url: https://www.reddit.com/r/rails/comments/es6xrb/need_some_help_with_modelschemarelationship_for_a/
---
Hi all,

So im working on a personal project to increase my skills in React particularly. I know Rails decent enough (Enough to finish the "Facebook Clone" on TheOdinProject. So at least an intermediate level I guess?

Anyways, for my personal project im going to be using a Rails backend to serve JSON back to my React frontend. My main concern right now is with one specific resource.

Basically my app is going to display Sensor and Sensor Data (Sensors are being handled via some ESP8266 devices sending POST requests with Temperature/Humidity/etc... to the API).

So I have 2 directions of data:  


* Sensor devices sending Historical measurements (Temp/Humidity/etc.) to my API via POST
* API sending Sensor data in JSON format to my frontend.

I think setting up the API should be simple enough, but im trying to figure out the best way to model/schema it. So far I think it needs to look something like this:

**SENSOR**: (Probably Contains a Name/ID/Device Image/Geographic Location (Eventually plan on having a map with geo coordinate). I may need a "Type" to differentiate? 

**SENSOR** has\_many **MEASUREMENTS** (A measurement would be one instance of data recording, IE: Temp/Humidity/Light/etc...). If you can think of a better name let me know

**SENSOR/MEASUREMENT** has\_one **TYPE** (Not sure if this is necessary It could help differentiate the measurements. So when I see "33" I know it's a humidity measurement and not a Temp measurement. Not sure if I necessarily need the "TYPE" for the Sensor though (Outside of maybe some sort of front-end categorization which may be helpful).

That's all I can think of right now. Also I wonder if it would be better to do a join table. IE: **SENSOR\_MEASUREMENTS** or something like that. Seems like it may be easier? IE: **SENSOR** has\_many **MEASUREMENTS** through **SENSOR\_MEASUREMENTS**

Let me know if anything seems obviously missing, or any suggestions. This is sort of a WIP project idea, and im def. no rails expert (Especially when it comes to architecting something myself.

Thanks!
## [8][How to access params hash in lib directory rails 6](https://www.reddit.com/r/rails/comments/es6bvy/how_to_access_params_hash_in_lib_directory_rails_6/)
- url: https://www.reddit.com/r/rails/comments/es6bvy/how_to_access_params_hash_in_lib_directory_rails_6/
---
In my rails app, I am using Kramdown to parse Markdown. I want to extend the functionality of the convert\_a method in the HTML converter. Part of this involves accessing the  database, but it is dependent on a parameter in the URL. Because I am  not directly calling the method that I am overriding I cannot simply  pass the method the params hash. Is there a way to access this hash, or even just get the current URL in a module in the libdirectory?

to give a bit more context, the method call is in a helper method here:

    # in app/helpers/myhelper.rb 
    def to_html(text)     
        Kramdown::Document.new(text, parse_block_html: true).to_custom_html 
    end

and here is the file in which I override the convert\_a:

    # in lib/custom_html.rb 
    class CustomHtml &lt; Kramdown::Converter::Html
        def convert_a(el, indent) 
            # use params[:foo] to make query         
            format_as_span_html(el.type, el.attr, inner(el, indent)) 
        end
    end
## [9][Understanding Apartment Gem Documentation](https://www.reddit.com/r/rails/comments/es0kvb/understanding_apartment_gem_documentation/)
- url: https://www.reddit.com/r/rails/comments/es0kvb/understanding_apartment_gem_documentation/
---
Hey guys, I'm trying to understand how apartment gem works with postgres.

I was Googling to understand certain methods at first, but then I decided to go through the docs which I thought would make my life easier.

I was going through apartment gem documentation on [rubydoc](https://www.rubydoc.info/gems/apartment/2.2.1)

To list all the methods that Apartment's classes provide, I clicked on Methods on the top left corner.

&amp;#x200B;

Here's what I don't understand:

To check the current Tenant (got this from stackoverflow), you run:

    Apartment::Tenant.current

I don't see 'current' method under [Tenant class](https://www.rubydoc.info/gems/apartment/2.2.1/Apartment/Tenant) at all.

&amp;#x200B;

Assuming docs is the only way of understanding the methods available, here are my questions:

1.	How am I supposed to know that Apartment::Tenant even has a method named 'current'?
2.	I see that `switch` method is an instance method defined in [AbstractAdapter](https://www.rubydoc.info/gems/apartment/2.2.1/Apartment/Adapters/AbstractAdapter#switch-instance_method). How is that possible if I can call `Apartment::Tenant.switch()`? Shouldn't it be a class method defined under Tenant class?
3.	I imagine Tenant class is extended from AbstractAdapter class, but where can I find the docs for it?

&amp;#x200B;
Any help is appreciated! Thank you
## [10][readyException.js:6 Uncaught TypeError: Cannot read property 'mData' of undefined](https://www.reddit.com/r/rails/comments/erzo8h/readyexceptionjs6_uncaught_typeerror_cannot_read/)
- url: https://www.reddit.com/r/rails/comments/erzo8h/readyexceptionjs6_uncaught_typeerror_cannot_read/
---
Hello guys I am having that error while loading my page when i added data tables, in a rails 6.0.0 app im working on. I used data tables CDN to download it and followed the official documentation. 

this is my html table 

`&lt;table id="example" class="title table index-table"&gt;`

  `&lt;tr&gt;`

`&lt;th&gt;Title&lt;/th&gt;`

`&lt;% if u/user['role'] === 'admin'  or u/user['role'] === 'instructor' %&gt;`

`&lt;th&gt;Tags&lt;/th&gt;`

`&lt;% end %&gt;`

  `&lt;/tr&gt;`

  `&lt;% u/assignments.each do |hw| %&gt;`

`&lt;tr&gt;`

`&lt;% if u/user['role'] === 'admin'  or u/user['role'] === 'instructor' %&gt;`

`&lt;td&gt;`

`&lt;span class="hw-title"&gt;&lt;%= link_to hw.title, assignment_path(hw) %&gt;&lt;/span&gt;`

`&lt;/td&gt;`

`&lt;% else %&gt;`

`&lt;td&gt;`

`&lt;span class="hw-title"&gt;&lt;%= link_to hw.title, new_assignment_solution_path(hw) %&gt;&lt;/span&gt;`

`&lt;/td&gt;`

`&lt;% end %&gt;`

`&lt;% if u/user['role'] === 'admin'  or u/user['role'] === 'instructor' %&gt;`

`&lt;td&gt;`

`&lt;% if hw.tags.empty? %&gt;`

`&lt;em&gt;&lt;/em&gt;`

`&lt;% else %&gt;`

`&lt;% hw.tags.each do |tag| %&gt;`

`&lt;% unless tag.name.include? ' ' %&gt;`

`&lt;span class="hw-tag"&gt;#&lt;%=` [`tag.name`](https://tag.name) `%&gt;&lt;/span&gt;`

`&lt;% else %&gt;`

`&lt;% words = tag.name.split(' ') %&gt;`

`&lt;% words.each do |word| %&gt;`

`&lt;span&gt;`

`#&lt;%= word.strip %&gt;`

`&lt;/span&gt;`

`&lt;% end %&gt;`

`&lt;% end %&gt;`

`&lt;% end %&gt;`

`&lt;% end %&gt;`

`&lt;/td&gt;`

`&lt;% end %&gt;`

`&lt;/tr&gt;`

  `&lt;% end %&gt;`

`&lt;/table&gt;`

&amp;#x200B;

  `&lt;script type="text/javascript"&gt;`

`$(document).ready(function() {`

`$('#example').DataTable();`

`});`

  `&lt;/script&gt;`

&amp;#x200B;

And i added the tags on the application html

`&lt;link rel="stylesheet" type="text/css" href="`[`https://cdn.datatables.net/1.10.20/css/jquery.dataTables.css`](https://cdn.datatables.net/1.10.20/css/jquery.dataTables.css)`"&gt;`

&amp;#x200B;

`&lt;script type="text/javascript" charset="utf8" src="`[`https://cdn.datatables.net/1.10.20/js/jquery.dataTables.js`](https://cdn.datatables.net/1.10.20/js/jquery.dataTables.js)`"&gt;&lt;/script&gt;`
## [11][Issue with Code Mirror and Heroku](https://www.reddit.com/r/rails/comments/erpm2k/issue_with_code_mirror_and_heroku/)
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
