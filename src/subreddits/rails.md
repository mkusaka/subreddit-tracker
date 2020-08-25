# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/i8dsvv/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/i8dsvv/personal_projects_show_off_your_own_project_andor/
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
## [2][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/id911w/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/id911w/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [3][after_commit callback not triggered from Concern](https://www.reddit.com/r/rails/comments/ig9fp6/after_commit_callback_not_triggered_from_concern/)
- url: https://www.reddit.com/r/rails/comments/ig9fp6/after_commit_callback_not_triggered_from_concern/
---
I have a model class like this

    class Task &lt; ActiveRecord::Base
      include Concerns::Tasks 
      self.table_name = "tasks" 
    end

and the concern is like in the app/models/concerns directory

    module Concerns::Tasks 
      extend ActiveSupport::Concern
      included do 
        after_commit :do_something
      end
      def do_something
        byebug
      end
    end

But my after\_commit callback isn't being hit at all. I'm on rails 4.2.1. Any ideas why?

EDIT: Other concern files are working fine. But only the newly added file isn't working fine. Do I have to register is somewhere?  


FINAL\_EDIT: The issue was in the filename. Should have known how rails picks up these files . Got it from here [https://stackoverflow.com/a/12306720/3575018](https://stackoverflow.com/a/12306720/3575018)
## [4][How to maintain CRUD when accepting nested attributes?](https://www.reddit.com/r/rails/comments/ifttz6/how_to_maintain_crud_when_accepting_nested/)
- url: https://www.reddit.com/r/rails/comments/ifttz6/how_to_maintain_crud_when_accepting_nested/
---
Let’s say these are my models:

Shift

\-- has\_many :employees

\-- has\_many :holiday\_shedules

\-- accepts\_nested\_attributes\_for :employees

\-- accepts\_nested\_attributes\_for :holiday\_schedules

Employees

\-- belongs\_to :shift

HolidaySchedule

\-- belongs\_to :shift

\-- has\_many :holidays

\-- accepts\_nested\_attributes\_for :holidays

Holiday

\-- belongs\_to HolidaySchedule

User steps:

1. Creates the shift

2. Adds staffing to the shift (multiple at a time)

3. Adds two, linked holiday\_schedules to the shift at a time, each with three holidays

Everything I’ve read says that custom controller actions are bad, but how do I maintain CRUD when accepts\_nested\_attributes\_for wants me to use a single controller update action for so many things?  That doesn’t even count updating the shift itself.

These are the options I’m seeing:

1.  Create or update multiple Nurses or HolidaySchedules through their respective controllers by namespacing and picking apart the parameters

2.  Send them all to Shift’s Update action and figure out what to do with them using conditionals

3.  Send them to custom actions in the shifts\_controller, i.e. def update\_nurses, def update\_holiday\_shifts

This is my first Rails project.  Are there options I’m not seeing?  What would be the most conventional, or Railsy, path?
## [5][Javascript not firing with Turbolinks](https://www.reddit.com/r/rails/comments/ifsu1z/javascript_not_firing_with_turbolinks/)
- url: https://www.reddit.com/r/rails/comments/ifsu1z/javascript_not_firing_with_turbolinks/
---
Hi Folks,
Using rails 5.2 with turbolinks and it's driving me crazy. Attempting to hide some items on a page when a button is clicked, but the js doesn't seem to fire. Even just trying to debug by using console.log and alert() methods doesn't seem help as neither fires when the button is clicked. Code works fine if I run it in the console, but clearly some issue with turbolinks here. Any ideas? Using an event listener on turbolinks:load (see below)

    document.addEventListener("turbolinks:load", function() {
        var btnWhiskey = document.getElementById('btn-Whiskey');
        btnWhiskey.addEventListener('click', function(){
             alert("testttt");
             console.log("TEST!");
        });
    });
## [6][Calling Database fails in secondary page.](https://www.reddit.com/r/rails/comments/ifzx3c/calling_database_fails_in_secondary_page/)
- url: https://www.reddit.com/r/rails/comments/ifzx3c/calling_database_fails_in_secondary_page/
---
I have a people database, that functions fine within it's own controller group (index, show, etc)

I want to use calls to this database in another controller, but it does not seem to be getting any data. I am sure I am missing something small, but I am hoping someone can help me.

outside\_controller:

    def pagename
        @people = Person.all
      end

pagename.html.erb

    &lt;% @people.each do |person| %&gt;
        &lt;li class="clearfix"&gt;
    	&lt;% if person.groffice == 'office' %&gt;
    		&lt;div class="image"&gt;
    		    &lt;%= image_tag("profiles/#{person.uname}S.jpg") %&gt;
    		&lt;/div&gt;
    		&lt;div class="body"&gt;
    			&lt;h5&gt;&lt;%= person.fname %&gt; &lt;%= person.lname %&gt;&lt;/h5&gt;
    		&lt;/div&gt;
    	&lt;% end %&gt;
        &lt;/li&gt;
    &lt;% end %&gt;

Person.rb

    class Person &lt; ApplicationRecord
    	has_many :pubs
    	default_scope { order('lname') }
    	belongs_to :boss, class_name: 'Person'
    	has_many :subordinates, class_name: 'Person', foreign_key: 'boss_id'
    
    	
    
        validates_presence_of :uname, :position, :fname, :lname # Needed for friendly URLs
    end
    

When I add the `person.groffice` call to my People Index, it pulls the data from that field and displays it without issue. I have a feeling I am missing something stupid. If I run `Person.where(:groffice =&gt; 'office')` in the console it will return the proper record.
## [7][Testing connection to db through irb, NameError Uninitialized constant "classname"](https://www.reddit.com/r/rails/comments/ifirdy/testing_connection_to_db_through_irb_nameerror/)
- url: https://www.reddit.com/r/rails/comments/ifirdy/testing_connection_to_db_through_irb_nameerror/
---
Hi all. I’m running into an issue where I want to test out the  connection to the db. I’m trying to access my Articles table. I have a  file in models “article.rb” and inside it has:class Article &lt; ApplicationRecordendWhen I go to my CMD and type the command “rails console” then follow up in IRB with “Article.all” I am receiving this error

    2.7.1 :001 &gt; Article.all Traceback (most recent call last):         1: from (irb):1 NameError (uninitialized constant Article) 

Here is a few pictures to show I have article.rb properly in modesl file and the code inside article.rb

edit: added database.yaml, migrate file, and schema pics in case it helps.

&amp;#x200B;

https://preview.redd.it/ko1lwyv37wi51.png?width=1920&amp;format=png&amp;auto=webp&amp;s=6a14daf5a0439fb03c0461b14ee18e8cc791df2b

https://preview.redd.it/eh51x8w37wi51.png?width=1913&amp;format=png&amp;auto=webp&amp;s=56c796b49cd034cb264be8868a7b83003b2e0c43

https://preview.redd.it/0ir6u8w37wi51.png?width=1916&amp;format=png&amp;auto=webp&amp;s=447a25b15fcc87014bb6cb5e4f18e67e5fe9977b

&amp;#x200B;

https://preview.redd.it/j29yuxbauvi51.png?width=1920&amp;format=png&amp;auto=webp&amp;s=5a6256dfdd9bbb2580491795159e0a1f0853b016

https://preview.redd.it/r2yyyubauvi51.png?width=1112&amp;format=png&amp;auto=webp&amp;s=bf43995799cbd02eee12ec47a2965a4e7a2ecb33
## [8][New Tutorial: How to Add Fields to a Devise User Signup in Rails 6](https://www.reddit.com/r/rails/comments/if3gcu/new_tutorial_how_to_add_fields_to_a_devise_user/)
- url: https://www.reddit.com/r/rails/comments/if3gcu/new_tutorial_how_to_add_fields_to_a_devise_user/
---
Do you want to add fields to your devise user signup in a rails 6 app?  In this tutorial we are going to add fields to our devise user sign up specifically: first\_name and last\_name.  We will also add a few validations to ensure some required fields are present.

Finally, we will add the confirmable module to ensure users sign up with a valid functioning email.  Confirmable is a Devise module that send the user an email to the address supplied to confirm accuracy — the user must click a link sent to the email to activate and use their account.  I have not seen content on how to add devise modules after a devise installation has already been complete, so I thought this could be useful for beginners.   

Please check it out and let me know what you think.

[https://youtu.be/0widKtkJONA](https://youtu.be/0widKtkJONA)
## [9][better use all data response in same API or use different API .](https://www.reddit.com/r/rails/comments/ifa4ef/better_use_all_data_response_in_same_api_or_use/)
- url: https://www.reddit.com/r/rails/comments/ifa4ef/better_use_all_data_response_in_same_api_or_use/
---
I have 7 models and i filter using 7 model from user to get data . data is big for example every key 50 keys (increase) i have 8 keys . it's good use one API to decrease go to db .

i ask better use one API or every key use API but every API go to db (make load in db )

more details

 models

    chat.rb 
    
    has_one :chat_statistic
    has_one :chat_message
    has_one :chat_rating
    has_one :chat_waiting_time
    has_many :messages_types
    has_one :chat_language
    has_one :chat_gender
    has_one :chat_subjects
    has_one :chat_type

i can filter using column for every other models .   


API have   


    1- overview {total_chats, number_chat_rating, largest_value_of_chat_type, total_waiting_time  }
    
    2- total_chats for every days ---&gt; have many keys for { day: number_total_chat_for this day }
    
    3- chat_type for every days  have many keys for { day: number_types_for this day }
    
    4- chat_analysis total value for every chat type --&gt; {x: 10, y:  20, z: 30}
    
    5- chat_subjects every days  have many keys for { day: number_subjects_for this day }
    
    6- chat_subjects_analysis  total value for every chat subjects --&gt; {x: 10, y:  20, z: 30}
    
    and another three keys 

**all values must include filter** 

&amp;#x200B;

&amp;#x200B;
## [10][node module not getting loaded in production (rails 6)](https://www.reddit.com/r/rails/comments/if5okd/node_module_not_getting_loaded_in_production/)
- url: https://www.reddit.com/r/rails/comments/if5okd/node_module_not_getting_loaded_in_production/
---
Hey,

I'm using the npm package [devicon](https://www.npmjs.com/package/devicon) which I added via `yarn add` to my project. In my `application.js` I have the following code:

```
// This file is automatically compiled by Webpack, along with any other files
// present in this directory. You're encouraged to place your actual application logic in
// a relevant structure within app/javascript and only use these pack files to reference
// that code so it'll be compiled.

require("@rails/ujs").start()
require("turbolinks").start()
require("@rails/activestorage").start()
require("channels")


// Uncomment to copy all static images under ../images to the output folder and reference
// them with the image_pack_tag helper in views (e.g &lt;%= image_pack_tag 'rails.png' %&gt;)
// or the `imagePath` JavaScript helper below.
//
// const images = require.context('../images', true)
// const imagePath = (name) =&gt; images(name, true)

require('jquery')
require('devicon')

// Sidebar
import { sidebarOn, sidebarOff } from './toggle_sidebar';
window.sidebarOn = sidebarOn;
window.sidebarOff = sidebarOff;

// Stimulus
import { Application } from "stimulus"
import { definitionsFromContext } from "stimulus/webpack-helpers"

const application = Application.start();
const context = require.context('../controllers', true, /\.js$/);
application.load(definitionsFromContext(context));
```  
In development everything works (devicon.css gets loaded and devicon font is available), but in production (on heroku) it's not getting loaded. What am I missing?  
  
  
  
**edit**  
`webpacker.yml`  
 
```
# Note: You must restart bin/webpack-dev-server for changes to take effect

default: &amp;default
  source_path: app/javascript
  source_entry_path: packs
  public_root_path: public
  public_output_path: packs
  cache_path: tmp/cache/webpacker
  check_yarn_integrity: false
  webpack_compile_output: true

  # Additional paths webpack should lookup modules
  # ['app/assets', 'engine/foo/app/assets']
  resolved_paths: []

  # Reload manifest.json on all requests so we reload latest compiled packs
  cache_manifest: false

  # Extract and emit a css file
  extract_css: false

  static_assets_extensions:
    - .jpg
    - .jpeg
    - .png
    - .gif
    - .tiff
    - .ico
    - .svg
    - .eot
    - .otf
    - .ttf
    - .woff
    - .woff2

  extensions:
    - .mjs
    - .js
    - .sass
    - .scss
    - .css
    - .module.sass
    - .module.scss
    - .module.css
    - .png
    - .svg
    - .gif
    - .jpeg
    - .jpg

development:
  &lt;&lt;: *default
  compile: true

  # Verifies that correct packages and versions are installed by inspecting package.json, yarn.lock, and node_modules
  check_yarn_integrity: true

  # Reference: https://webpack.js.org/configuration/dev-server/
  dev_server:
    https: false
    host: localhost
    port: 3035
    public: localhost:3035
    hmr: false
    # Inline should be set to true if using HMR
    inline: true
    overlay: true
    compress: true
    disable_host_check: true
    use_local_ip: false
    quiet: false
    pretty: false
    headers:
      'Access-Control-Allow-Origin': '*'
    watch_options:
      ignored: '**/node_modules/**'


test:
  &lt;&lt;: *default
  compile: true

  # Compile test packs to a separate directory
  public_output_path: packs-test

production:
  &lt;&lt;: *default

  # Production depends on precompilation of packs prior to booting for performance.
  compile: false

  # Extract and emit a css file
  extract_css: true

  # Cache manifest.json for performance
  cache_manifest: true
```
## [11][Input Mask for Rails 6 and Webpacker](https://www.reddit.com/r/rails/comments/ieikj8/input_mask_for_rails_6_and_webpacker/)
- url: https://www.reddit.com/r/rails/comments/ieikj8/input_mask_for_rails_6_and_webpacker/
---
I can’t imagine that there isn’t a decent JS or StimulusJS solution for input mask (for example as simple as comma separation for numbers) 

There are few libraries that I tried but I just can’t make it work with webpaker, turbolink. I do no want to put the jQuery in the environment.rb instead use CDN (I would rather not use CDN, my challenge is adding jQuery in the application level adds to the front-end of the site which I don’t want because the front-end uses a different version of jQuery and I would like to keep them separate).  Unless there is a way to add jQuery to packs config (I have two separate packs front-end and back-end)

Does anyone have a good example?  Something that you have actually used and can share the codes with?
## [12][Alternatives for Bootstrap framework?](https://www.reddit.com/r/rails/comments/iebb66/alternatives_for_bootstrap_framework/)
- url: https://www.reddit.com/r/rails/comments/iebb66/alternatives_for_bootstrap_framework/
---
I have a rails + stimulus app and using Bootstrap for styling. The thing is: Bootstrap depends on jquery, and  this last one is a heavy file after compilation. Because I don't use too much of Bootstrap components and neither Jquery, I would like something more lightweight after compilation. 
 
One more think, I'm not a expertise on css e.g I would need to search how to implement a system grid or how animate dropdown/cascading, for references. 

So my question is: There are others alternatives frameworks that are stable and lightweight? Also, do you guys usually wrote you on components (not talking about JS components e.g react), wrote your own system grid, etc...?

About the Jquery size, I don't remember right now about how much, but I use, sometime ago, a lib that shows by size the files compiled by webpack, and as result the Jquery was the biggest.
