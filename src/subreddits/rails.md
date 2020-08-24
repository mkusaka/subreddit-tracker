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
## [3][Testing connection to db through irb, NameError Uninitialized constant "classname"](https://www.reddit.com/r/rails/comments/ifirdy/testing_connection_to_db_through_irb_nameerror/)
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
## [4][New Tutorial: How to Add Fields to a Devise User Signup in Rails 6](https://www.reddit.com/r/rails/comments/if3gcu/new_tutorial_how_to_add_fields_to_a_devise_user/)
- url: https://www.reddit.com/r/rails/comments/if3gcu/new_tutorial_how_to_add_fields_to_a_devise_user/
---
Do you want to add fields to your devise user signup in a rails 6 app?  In this tutorial we are going to add fields to our devise user sign up specifically: first\_name and last\_name.  We will also add a few validations to ensure some required fields are present.

Finally, we will add the confirmable module to ensure users sign up with a valid functioning email.  Confirmable is a Devise module that send the user an email to the address supplied to confirm accuracy — the user must click a link sent to the email to activate and use their account.  I have not seen content on how to add devise modules after a devise installation has already been complete, so I thought this could be useful for beginners.   

Please check it out and let me know what you think.

[https://youtu.be/0widKtkJONA](https://youtu.be/0widKtkJONA)
## [5][better use all data response in same API or use different API .](https://www.reddit.com/r/rails/comments/ifa4ef/better_use_all_data_response_in_same_api_or_use/)
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
## [6][node module not getting loaded in production (rails 6)](https://www.reddit.com/r/rails/comments/if5okd/node_module_not_getting_loaded_in_production/)
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
## [7][Input Mask for Rails 6 and Webpacker](https://www.reddit.com/r/rails/comments/ieikj8/input_mask_for_rails_6_and_webpacker/)
- url: https://www.reddit.com/r/rails/comments/ieikj8/input_mask_for_rails_6_and_webpacker/
---
I can’t imagine that there isn’t a decent JS or StimulusJS solution for input mask (for example as simple as comma separation for numbers) 

There are few libraries that I tried but I just can’t make it work with webpaker, turbolink. I do no want to put the jQuery in the environment.rb instead use CDN (I would rather not use CDN, my challenge is adding jQuery in the application level adds to the front-end of the site which I don’t want because the front-end uses a different version of jQuery and I would like to keep them separate).  Unless there is a way to add jQuery to packs config (I have two separate packs front-end and back-end)

Does anyone have a good example?  Something that you have actually used and can share the codes with?
## [8][Alternatives for Bootstrap framework?](https://www.reddit.com/r/rails/comments/iebb66/alternatives_for_bootstrap_framework/)
- url: https://www.reddit.com/r/rails/comments/iebb66/alternatives_for_bootstrap_framework/
---
I have a rails + stimulus app and using Bootstrap for styling. The thing is: Bootstrap depends on jquery, and  this last one is a heavy file after compilation. Because I don't use too much of Bootstrap components and neither Jquery, I would like something more lightweight after compilation. 
 
One more think, I'm not a expertise on css e.g I would need to search how to implement a system grid or how animate dropdown/cascading, for references. 

So my question is: There are others alternatives frameworks that are stable and lightweight? Also, do you guys usually wrote you on components (not talking about JS components e.g react), wrote your own system grid, etc...?

About the Jquery size, I don't remember right now about how much, but I use, sometime ago, a lib that shows by size the files compiled by webpack, and as result the Jquery was the biggest.
## [9][How to make the comment’s content unique based on its parent post?](https://www.reddit.com/r/rails/comments/ie8zri/how_to_make_the_comments_content_unique_based_on/)
- url: https://www.reddit.com/r/rails/comments/ie8zri/how_to_make_the_comments_content_unique_based_on/
---
I don’t want to make it globally unique in the model file. I just want the content of the comment to be unique if the post already has a comment with same content.
## [10][Is it possible to output custom text to the console/server log on startup?](https://www.reddit.com/r/rails/comments/idyetp/is_it_possible_to_output_custom_text_to_the/)
- url: https://www.reddit.com/r/rails/comments/idyetp/is_it_possible_to_output_custom_text_to_the/
---
Bit of a silly request but I want to add some ASCII art to my applications startup so I can see the app logo every time it runs.  

I’m sure it’s possible but where would I need to add my puts statement for this to work?
## [11][Wits-end aws elb ssl eks ingress rails not recognizing https scheme in auth0](https://www.reddit.com/r/rails/comments/ie0kum/witsend_aws_elb_ssl_eks_ingress_rails_not/)
- url: https://www.reddit.com/r/rails/comments/ie0kum/witsend_aws_elb_ssl_eks_ingress_rails_not/
---
First apologies for breaking any rules, we're in panic mode and I didn't have a chance to review rules.
We have in fact searched for a solution frequently and have attempted many fixes suggested on stack overflow / medium etc to know success. 

Long story short
Terminating ssl in aws l7 elb. 
Ingress controller routing host to service-&gt;deployment-&gt;pod-&gt;container of a rails app attempting to use auth0 as the IDP.

Initial site loads properly on https.
However it appears as if generated links (e.g. Redirect and call back urls) are missing awareness of the https scheme.

We've attempted modification of the ingress annotations but seems to be coming up short.

X forward for and proto and scheme and host headers have been tested but we lack visibility into whether or not they are actually getting passed to the rails app.

Working on adding verbosity to that effect now but curious if anyone has experience with this pattern.

Also forgive mobile formatting. 

Thank you!


Just a quick update:
We resolved the issue initially by moving ssl termination into the ingress controller and changing from L7 load balancer to L4 load balancer. This isn't ideal and will continue to investigate the L7 config.  I'll be sure to post an update when we resolve.
## [12][Newbie dev here, got some questions about Rails as a backend](https://www.reddit.com/r/rails/comments/idxys7/newbie_dev_here_got_some_questions_about_rails_as/)
- url: https://www.reddit.com/r/rails/comments/idxys7/newbie_dev_here_got_some_questions_about_rails_as/
---
\- Can you use Rails as a back end for a react native app? 

And if you can, could anyone link me to resources that explain how to do it
