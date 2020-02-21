# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/ezrfed/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/ezrfed/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/f2r9mk/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/f2r9mk/personal_projects_show_off_your_own_project_andor/
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
## [3][I created a step-by-step tutorial demonstrating how to integrate React with Ruby on Rails](https://www.reddit.com/r/rails/comments/f7a0v8/i_created_a_stepbystep_tutorial_demonstrating_how/)
- url: https://www.reddit.com/r/rails/comments/f7a0v8/i_created_a_stepbystep_tutorial_demonstrating_how/
---
- [Live Demo](https://rails-react-example.herokuapp.com/)
- [Tutorial](https://stevepolito.design/blog/rails-react-tutorial/)

I really wanted to learn React and API development, so I went head first into building a simple application, and documented my experience. I think what sets this apart from other Rails and React tutorials is that I cover...

- API authorization
- API versioning
- Setting HTTP status codes
- Form validation on the front-end
- Handling errors
- Debouncing requests
- CSRF Countermeasures
## [4][Write test for data migrations](https://www.reddit.com/r/rails/comments/f717t9/write_test_for_data_migrations/)
- url: https://www.reddit.com/r/rails/comments/f717t9/write_test_for_data_migrations/
---
Today I heard a funny story at a local Ruby meetup. A guy corrupted data in production by a data migration. Don’t repeat his mistake, use this [gem](https://github.com/ka8725/migration_data). It’s not super popular, but has been tested over time. Surprisingly, people still use it and seem satisfied. Any contribution is welcomed!
## [5][What book/resources do you recommend for intermediate Rails developer?](https://www.reddit.com/r/rails/comments/f744ae/what_bookresources_do_you_recommend_for/)
- url: https://www.reddit.com/r/rails/comments/f744ae/what_bookresources_do_you_recommend_for/
---
I’ve been developing in Rails for about a year now and feel I’m ready to learn intermediate stuff. I have read Michael Hartl’s book, and did a couple of tutorials on udemy as well. What resources/books would you recommend? Thanks!
## [6][Making a desktop and mobile app from an existing rails app.](https://www.reddit.com/r/rails/comments/f76oug/making_a_desktop_and_mobile_app_from_an_existing/)
- url: https://www.reddit.com/r/rails/comments/f76oug/making_a_desktop_and_mobile_app_from_an_existing/
---
I have an existing Rails app I want to make a desktop app(using electron) and a mobile app. So I'm thinking of making a REST API. I want to ask if there is a better way then converting the Rails app to an API. Also, what is the best way to make Rails API out of an already built app?
## [7][What are some smaller Open source rails projects That are less intimidating to contribute to?](https://www.reddit.com/r/rails/comments/f713e8/what_are_some_smaller_open_source_rails_projects/)
- url: https://www.reddit.com/r/rails/comments/f713e8/what_are_some_smaller_open_source_rails_projects/
---
The current list (Will Update as People Comment)

[https://github.com/focallocal](https://github.com/focallocal)

[https://github.com/TheOdinProject/theodinproject](https://github.com/TheOdinProject/theodinproject)

[http://www.opensourcerails.com](http://www.opensourcerails.com/)

I am looking to share this list with my local user group to help other beginners like me.
## [8][Rails server on cloud9 not working?](https://www.reddit.com/r/rails/comments/f71fd6/rails_server_on_cloud9_not_working/)
- url: https://www.reddit.com/r/rails/comments/f71fd6/rails_server_on_cloud9_not_working/
---
As the title suggets, when i start up the rails server, it acts as if it wasn't even on. Any way to fix this?
## [9][Trying to use Simple Form and Enumerize gems to set a default value of an input. Help needed!](https://www.reddit.com/r/rails/comments/f737p3/trying_to_use_simple_form_and_enumerize_gems_to/)
- url: https://www.reddit.com/r/rails/comments/f737p3/trying_to_use_simple_form_and_enumerize_gems_to/
---
Hi all,  


I'm working with someone else's code and trying to work out how to set the value of an input as a default using simple form. Probably easier to explain with some code...

In the view I have this haml code that currently allows users to chose between two options from the payments attribute, which is enumerated in the model.  


view:

    = simple_form_for(@checkout, html: { class: 'edit-form edit_item' }, wrapper: :edit_form) do |f|
      %fieldset.large-2-col
        .fields-block
          = f.input :payments, as: :check_boxes
...

The payments attribute is enumerated in the model :

    enumerize :payments, in: [:cash, :card], default: :card, multiple: true

What I would like is for a user to only be offered the options of `:cash` and `:card` if they have those options set on their account. Currently if a user only has an account registered with card they can still select the cash radio button. In very crude terms it should do something like this:  


    - if current_account.cash_accounts.any?  == false
      default simple form input for card and don't show options
    - elsif current_account.card_accounts.any? == false
      default simple form input for cash and don't show options
    - else
  = f.input :payments, as: :check_boxes

I don't know how to just get one of :card or :cash from the enumerated attribute (depending on the user's account) and I don't know how to use simple form to have a default value selected whilst not showing anything in the view. Any help would be greatly appreciated :)
## [10][Best gem to limit zip codes at checkout?](https://www.reddit.com/r/rails/comments/f6wbfa/best_gem_to_limit_zip_codes_at_checkout/)
- url: https://www.reddit.com/r/rails/comments/f6wbfa/best_gem_to_limit_zip_codes_at_checkout/
---
Hi,

I'm building an ecommerce app, but I want to limit the geographic locations where orders can be shipped to. Wondering if there is a gem that's best recommended for this that I can incorporate into a function that won't allow customers to check out unless they have entered a zip code in the allowed range(s). Thanks in advance!
## [11][Rails 6/Webpacker/React/Bootstrap users....PLEASE HELP!!](https://www.reddit.com/r/rails/comments/f6yizg/rails_6webpackerreactbootstrap_usersplease_help/)
- url: https://www.reddit.com/r/rails/comments/f6yizg/rails_6webpackerreactbootstrap_usersplease_help/
---
I am completely stumped as to why a react-bootstrap Modal component throws an error in \`production\` but not \`development\`. In my app, if I run \`rails s\` and \`bin/webpack-dev-server\`, a 'create\` action that opens a modal with a form using react-bootstrap works fine. When I try it in \`production\` with \`RAILS\_ENV=production rake assets:precompile\` and \`rails s -e production\`, the create action that opens the bootstrap modal fails. The error:

&amp;#x200B;

    backend.js:6 TypeError: Cannot convert undefined or null to object
        at hasOwnProperty (&lt;anonymous&gt;)
        at Modal.js:21
        at Array.forEach (&lt;anonymous&gt;)
        at Modal.js:20
        at t.n.render (Modal.js:302)
        at Qi (react-dom.production.min.js:4243)
        at Ji (react-dom.production.min.js:4234)
        at wc (react-dom.production.min.js:6676)
        at yu (react-dom.production.min.js:5650)
        at Mu (react-dom.production.min.js:5639)
    r @ backend.js:6
    pc @ react-dom.production.min.js:4638
    n.callback @ react-dom.production.min.js:5083
    La @ react-dom.production.min.js:3061
    va @ react-dom.production.min.js:3049
    vu @ react-dom.production.min.js:6347
    t.unstable_runWithPriority @ scheduler.production.min.js:272
    Vo @ react-dom.production.min.js:2796
    gu @ react-dom.production.min.js:6107
    cu @ react-dom.production.min.js:5410
    (anonymous) @ react-dom.production.min.js:2831
    t.unstable_runWithPriority @ scheduler.production.min.js:272
    Vo @ react-dom.production.min.js:2796
    Qo @ react-dom.production.min.js:2826
    Jo @ react-dom.production.min.js:2816
    ue @ react-dom.production.min.js:7246
    kn @ react-dom.production.min.js:1712
    Modal.js:21 Uncaught TypeError: Cannot convert undefined or null to object
        at hasOwnProperty (&lt;anonymous&gt;)
        at Modal.js:21
        at Array.forEach (&lt;anonymous&gt;)
        at Modal.js:20
        at t.n.render (Modal.js:302)
        at Qi (react-dom.production.min.js:4243)
        at Ji (react-dom.production.min.js:4234)
        at wc (react-dom.production.min.js:6676)
        at yu (react-dom.production.min.js:5650)
        at Mu (react-dom.production.min.js:5639)
    (anonymous) @ Modal.js:21
    (anonymous) @ Modal.js:20
    n.render @ Modal.js:302
    Qi @ react-dom.production.min.js:4243
    Ji @ react-dom.production.min.js:4234
    wc @ react-dom.production.min.js:6676
    yu @ react-dom.production.min.js:5650
    Mu @ react-dom.production.min.js:5639
    cu @ react-dom.production.min.js:5395
    (anonymous) @ react-dom.production.min.js:2831
    t.unstable_runWithPriority @ scheduler.production.min.js:272
    Vo @ react-dom.production.min.js:2796
    Qo @ react-dom.production.min.js:2826
    Jo @ react-dom.production.min.js:2816
    ue @ react-dom.production.min.js:7246
    kn @ react-dom.production.min.js:1712

I'm not sure what in the production compile causes this to fail. This is all running from a local server and was discovered when trying to deploy to Heroku. 

&amp;#x200B;

Here is how my directory tree is laid out: 

&amp;#x200B;

    ├── assets
    │   ├── config
    │   ├── images
    │   └── stylesheets
    ├── channels
    │   └── application_cable
    ├── controllers
    │   └── concerns
    ├── helpers
    ├── javascript
    │   ├── channels
    │   ├── css
    │   └── packs
    │       ├── actions
    │       ├── components
    │       ├── containers
    │       ├── middleware
    │       ├── public
    │       └── reducers
    ├── jobs
    ├── mailers
    ├── models
    │   └── concerns
    ├── serializers
    └── views
        ├── current_outages
        ├── dashboard
        ├── future_outages
        ├── layouts
        ├── recurring_outages
        └── services

and the webpacker.yml file: 

&amp;#x200B;

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
      resolved_paths: ['app/javascript/css']
    
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
        - .jsx
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
      check_yarn_integrity: false
    
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
    

PLEASE HELP! Been on this for 3 days now.
## [12][Stripe: Tiered pricing for one time purchase?](https://www.reddit.com/r/rails/comments/f6uixy/stripe_tiered_pricing_for_one_time_purchase/)
- url: https://www.reddit.com/r/rails/comments/f6uixy/stripe_tiered_pricing_for_one_time_purchase/
---
I am working on a project where I need to capture one time payments from an account. For a bit of background: The account has many users, where users are part of teams as team\_members. I have another model we can call projects where teams are affixed to the project. It is on these individual projects that I would like to have a checkout button. The price for this one time payment should be calculated based on the number of individuals in a team, that are part of this project.

My issue is that I am not 100% sure how to achieve this for one time payments in stripe. Should I set this up alike that if I were to sell an individual item? Product (with name, description and price) and an Order to affix to the user? 

I've built subscriptions plans in Stripe before, but have never really used one time payments. Any guidance on setup here would be really great.
