# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/hja8kx/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/hja8kx/personal_projects_show_off_your_own_project_andor/
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
## [2][Best and cheapest hosting options?](https://www.reddit.com/r/rails/comments/hlgze5/best_and_cheapest_hosting_options/)
- url: https://www.reddit.com/r/rails/comments/hlgze5/best_and_cheapest_hosting_options/
---
What would you recommend for a good, Rails-tailored hosting plan (like Heroku) for around $7-8 per month? That is what my client told me she would like to pay. Heroku Hobby is $7, but is that good enough for a production e-commerce site?

If not, what else would you recommend? I'd like something like Heroku that is specifically tailored to Rails where you can push your code from GitHub and it works. I tried and failed to deploy my app to Digital Ocean's Rails droplet (and I'm far from being a noob when it comes to the Linux CLI).
## [3][Devs who dislike or has been bitten by Active Record's default_scope, how about we include it in style guides?](https://www.reddit.com/r/rails/comments/hlakkq/devs_who_dislike_or_has_been_bitten_by_active/)
- url: https://www.reddit.com/r/rails/comments/hlakkq/devs_who_dislike_or_has_been_bitten_by_active/
---
[https://github.com/rubocop-hq/rails-style-guide/issues/266](https://github.com/rubocop-hq/rails-style-guide/issues/266)
## [4][Turbolinks and preserving focus state](https://www.reddit.com/r/rails/comments/hlmmhy/turbolinks_and_preserving_focus_state/)
- url: https://www.reddit.com/r/rails/comments/hlmmhy/turbolinks_and_preserving_focus_state/
---
Howdy.

Is there a way of preserving focus state? My use case is that I want the "new" form to autosave with a debounce, and once it saves it should become the "edit" form of the saved entry. I can manually replace the history state but I think that it would be cool to get it for free from Turbolinks, almost like Inertia.
## [5][Trying to understand If-Modified-Since HTTP Header](https://www.reddit.com/r/rails/comments/hlgzev/trying_to_understand_ifmodifiedsince_http_header/)
- url: https://www.reddit.com/r/rails/comments/hlgzev/trying_to_understand_ifmodifiedsince_http_header/
---
I'm building a Rails app that consumes an API that accepts this header.

&gt; "The server then determines based on these headers if the client's cached copy is the most recent version of this file. If so, the server returns an HTTP status 304 code, letting the client know it can reuse its cached copy."

But I don't know if I understand the idea, does this mean I have to always save the result of each API request to a file or database? and if the request's result has an HTTP status 304 should I rescue it from the database? What if I don't have the result in the database? Do I send you a request without that header to get the result again?
## [6][What is your Editor and setup for Rails development.](https://www.reddit.com/r/rails/comments/hl609b/what_is_your_editor_and_setup_for_rails/)
- url: https://www.reddit.com/r/rails/comments/hl609b/what_is_your_editor_and_setup_for_rails/
---
I really love vs code and use it even though it takes a lot of memory. I recently started getting my feet wet with vim and have learned some movements in it. So I use the vscode vim plugin to move within the file using vim motions. Honestly, it's a good intro to vim motions even though sometimes the extension crashes.

I huge drawback is finding where the method being used is defined, and to fo to the definition and find all references to a method. I use ruby and solargraph extensions but it doesn't work most of the time in non-framework code and take a huge amount of time even if it does so. Anybody has any suggestions to work within vs-code.   


Other than that I use go-to-spec which easily gets you to the spec file and vice versa. 

and end-wise which adds end to your method but doesn't work that well. 

Any other extension which you believe is of great use?

Also is it possible to key the rails application running and infer the method definition and other information since most of the things are runtime?
## [7][New bootstrap toasts, pretty neat](https://www.reddit.com/r/rails/comments/hli2tq/new_bootstrap_toasts_pretty_neat/)
- url: https://www.reddit.com/r/rails/comments/hli2tq/new_bootstrap_toasts_pretty_neat/
---
&amp;#x200B;

[toasts](https://i.redd.it/q4zhf2oxez851.gif)

&amp;#x200B;

My application.html.erb  


    &lt;!DOCTYPE html&gt;
    &lt;html&gt;
      &lt;head&gt;
        ...
        &lt;%= javascript_pack_tag 'flash', 'data-turbolinks-track': 'reload' %&gt;
      &lt;/head&gt;
    
      &lt;body&gt;
        ...
        &lt;%= render '/flashes' %&gt;
        &lt;%= render '/toasts' %&gt;
        &lt;%= yield %&gt;
      &lt;/body&gt;
    &lt;/html&gt;


 app/views/\_flashes.html.erb   


    &lt;div class="acme-flashes notice d-none"&gt;
      &lt;%= flash[:notice] %&gt;
    &lt;/div&gt;
    &lt;div class="acme-flashes alert d-none"&gt;
      &lt;%= flash[:alert] %&gt;
    &lt;/div&gt;

 

app/views/\_toasts.html.erb 

    &lt;div class="acme-toasts position-absolute w-100 p-4 d-flex flex-column align-items-end" style="z-index: 1"&gt;
      &lt;div class="w-25 inner"&gt;
        &lt;!-- empty --&gt;
      &lt;/div&gt;
    &lt;/div&gt;

&amp;#x200B;

app/javascript/packs/flash.js

    require("acme/flash").start()

app/javascript/acme/flash.js

    export function start() {
      document.addEventListener("turbolinks:load", () =&gt; {
        create()
        show()
      })
    }
    
    function show() {
      $('.toast')
        .toast({ delay: 2000 })
        .toast('show')
    }
    
    function create() {
      $.each( getFlashesInDom(), function(index, value) {
        $(".acme-toasts .inner").append(
          template(value)
        )
      })
    }
    
    function template(body) {
      return `
        &lt;div class="toast" role="status" aria-live="polite" aria-atomic="true" zdata-delay=3000 style="z-index: 1"&gt;
          &lt;div class="toast-header"&gt;
            &lt;!-- an icon can be here --&gt;
            &lt;strong class="mr-auto"&gt;
              &lt;!-- title here --&gt;
            &lt;/strong&gt;
            &lt;small class="text-muted"&gt;just now&lt;/small&gt;
            &lt;button type="button" class="ml-2 mb-1 close" data-dismiss="toast" aria-label="Close"&gt;
              &lt;span aria-hidden="true"&gt;&amp;times;&lt;/span&gt;
            &lt;/button&gt;
          &lt;/div&gt;
          &lt;div class="toast-body"&gt;
            ${body}
          &lt;/div&gt;
        &lt;/div&gt;`
    }
    
    function getFlashesInDom() {
      return $.map( $(".acme-flashes"), function(val,i) {
        let html = $.trim( $(val).html() )
        if ( html.length == 0 ) { return }
        return html
      })
    }

config/webpack/environment.js  


    const { environment } = require('@rails/webpacker')
    
    // https://github.com/webpack-contrib/expose-loader/blob/master/README.md
    environment.loaders.append('expose', {
            test: require.resolve('jquery'),
            loader: 'expose-loader',
            options: {
              // For `underscore` library, it can be `_.map map` or `_.map|map`
              exposes: ['jquery', '$', 'jQuery']
            }
    });
    
    module.exports = environment

My yarn.lock has  bootstrap@\^4.5.0  
btw.
## [8][How to "merge" two regex format (with and without)](https://www.reddit.com/r/rails/comments/hl9sst/how_to_merge_two_regex_format_with_and_without/)
- url: https://www.reddit.com/r/rails/comments/hl9sst/how_to_merge_two_regex_format_with_and_without/
---
Can I write this in just one regex expression?

    format: { with: /\A[^\.\%\&amp;]+\z/ },
    format: { without: /\s/ }
## [9][How to translate message error in model files](https://www.reddit.com/r/rails/comments/hl9v85/how_to_translate_message_error_in_model_files/)
- url: https://www.reddit.com/r/rails/comments/hl9v85/how_to_translate_message_error_in_model_files/
---
I have a message error in user.rb mode, in validates area.

It is in english. I have a multilanguage website. Usually in views folder i use something like `&lt;%= t('.error_message') %&gt;`

How to add this into a model file like this?

    validates :password, format: {with: /^\w{3,15}$/, message: 'must be 3-15 letters'}
## [10][Automatic Book Tracker built with ruby on rails](https://www.reddit.com/r/rails/comments/hlflwq/automatic_book_tracker_built_with_ruby_on_rails/)
- url: https://www.reddit.com/r/rails/comments/hlflwq/automatic_book_tracker_built_with_ruby_on_rails/
---
Hey everyone!

I have recently finished my tutorial on building a ruby on rails based tracker. My videos assume no knowledge of the rails framework and teach how to get started whilst building a simple projects.

If you are interested starting out with ruby on rails with a cool project this is the place for you!

 [https://www.youtube.com/watch?v=uEwu7D5G-hU&amp;list=PLB4RncStK2LUbl9VWLQAHznLJrYz2YMB4](https://www.youtube.com/watch?v=uEwu7D5G-hU&amp;list=PLB4RncStK2LUbl9VWLQAHznLJrYz2YMB4) 

Hope you enjoy!
## [11][Job Queue Design for Multiple Small Transactions](https://www.reddit.com/r/rails/comments/hkz9pf/job_queue_design_for_multiple_small_transactions/)
- url: https://www.reddit.com/r/rails/comments/hkz9pf/job_queue_design_for_multiple_small_transactions/
---
Hi everyone

I have a task that I run via Sidekiq + ActiveJob + Redis. The task is a collection of ActiveRecord instances with a one to many relation to another model which performs updates on the "to many" table for each "belongs\_to". Is it considered best practice to have a single task for it in a loop or should I break it down to multiple smaller tasks to be put on queue?

Thanks
