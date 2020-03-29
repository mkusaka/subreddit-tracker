# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/foqc07/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/foqc07/personal_projects_show_off_your_own_project_andor/
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
## [2][Add react/Vue existing Rails app ?](https://www.reddit.com/r/rails/comments/fr32wy/add_reactvue_existing_rails_app/)
- url: https://www.reddit.com/r/rails/comments/fr32wy/add_reactvue_existing_rails_app/
---

#newbieHere 
Have you add react or vue on existing Rails  ? I have heard Webpacker would be done this to use React or Vue , would it affects on Rails version also if I add  FE framework?

If you use react or vue on existing Rails app without do "rails new my-app --webpack=react" ? Here is my another questions
 - should I use gem to use react or vue ? Such as "gem install react_on_rails"
- or I just use "yard add react react-dom" ?
## [3][noob questions](https://www.reddit.com/r/rails/comments/fqpsyp/noob_questions/)
- url: https://www.reddit.com/r/rails/comments/fqpsyp/noob_questions/
---
So, I've been a Microsoft stack guy for most of my career.  Especially on the web.  I've been happy enough with it to have never looked at Ruby or RoR at all.  In the lockdown, I'm spending some time learning "new" things and have just now watched some Ruby and RoR videos.

It looks to me like "ASP.old (before .NET) but with MVC and Active Record baked in."  Is that about right?

Also, if I were to write a Rails site, where would I host it?  I'm used to using Netlify or just S3 these days for "JAMstack" front ends, with back-end APIs on Azure.

If you had to give an elevator speech about why I might want to consider switching from Vue.js/.NET APIs to RoR, what wold it be?

Thanks!
## [4][Unable to access instance variables in .js.erb files](https://www.reddit.com/r/rails/comments/fqrdlr/unable_to_access_instance_variables_in_jserb_files/)
- url: https://www.reddit.com/r/rails/comments/fqrdlr/unable_to_access_instance_variables_in_jserb_files/
---
Hello,

I posted the following a while ago related to getting action specific javascript to wok in rails 6 (JQuery wasn't working)

https://www.reddit.com/r/rails/comments/dg2c0z/cant_get_action_specific_javascript_working_in/

I feel this is a different issue (of course I could be wrong). I am attempting to get a create action working via javascript `create.js.erb`, but it seems like I can't access the instance variable created by the controller. Just wondering what I am not understanding.

Here is what I have so far:

app/controllers/admin/permissions_controller.rb

      def create
        @permission = Permission.new permission_params

        respond_to do |format|
          if @permission.save
            format.html { redirect_to admin_role_path(@role), success: success_message(@permission) }
            format.js
          else
            format.html { redirect_to admin_role_path(@role), error: error_message }
            format.js
          end
        end
      end

app/views/admin/permissions/create.js.erb

    elm = document.getElementById(&lt;%= dom_id(@permission.user) %&gt;);
    elm.style.display = "none";

Console prints out

    VM45161:1 Uncaught SyntaxError: Unexpected string
        at processResponse (rails-ujs.js:283)
        at rails-ujs.js:196
        at XMLHttpRequest.xhr.onreadystatechange (rails-ujs.js:264)
## [5][Formatting Rails Exectuables](https://www.reddit.com/r/rails/comments/fqpkgu/formatting_rails_exectuables/)
- url: https://www.reddit.com/r/rails/comments/fqpkgu/formatting_rails_exectuables/
---
Hi all, so basically, I've been getting into "prettifying" and customizing my terminal experience. This includes git as well as the shell itself. Specifically, one that I am using now that I really enjoy is using the `--pretty` flag for `git log`. 

I was wondering if there's an equivalent for rails executables, specifically, in my case, `(bundle exec) rails routes`. Namely, I want a way to skip the printing of the default Active Storage routes when running this command. 

Thanks!
## [6][ActionText iframe embeds](https://www.reddit.com/r/rails/comments/fqj15c/actiontext_iframe_embeds/)
- url: https://www.reddit.com/r/rails/comments/fqj15c/actiontext_iframe_embeds/
---
Hi,

I am currently building some kind of inventory tracking system and I am using ActionText to capture rich text descriptions for products.

It would be great if I could just paste iframe embed and view it on `show` page. It doesn't need to be transformed to embed immediately in editor. I know this can be security risk, but this is in dashboard and only admins can edit this content so is there a way to disable ActionText html sanitization for `&lt;iframe&gt;` html tag?

Currently `@product.description.body.to_trix_html` outputs santized `&lt;iframe&gt;` tags.

    &amp;lt;iframe width=\"560\"...

Thanks r/rails for helping me, I have learned a lot about rails from answers from this community.

Bartol

&amp;#x200B;

Edit:

Solved!

    &lt;div class="trix-content"&gt;
      &lt;%= raw CGI.unescape_element @product.description.body.to_s, ['iframe'] %&gt;
    &lt;/div&gt;
## [7][When do you decide to rewrite a project?](https://www.reddit.com/r/rails/comments/fq9hhh/when_do_you_decide_to_rewrite_a_project/)
- url: https://www.reddit.com/r/rails/comments/fq9hhh/when_do_you_decide_to_rewrite_a_project/
---
I am revisiting an old rails app that I was developing, but the stack I used is completely different and the code quality and design is pretty bad. I'm considering just doing a complete rewrite. When do you guys say screw it and rewrite it?
## [8][using instance_of with mocha for ActionController::TestCase](https://www.reddit.com/r/rails/comments/fqez0r/using_instance_of_with_mocha_for/)
- url: https://www.reddit.com/r/rails/comments/fqez0r/using_instance_of_with_mocha_for/
---
In my controller test, I need to stub a method where a parameter should be an instance of a class. I decided I wanted to use instance\_of method here =&gt; [https://mocha.jamesmead.org/Mocha/API.html](https://mocha.jamesmead.org/Mocha/API.html)

I use `ActionController::TestCase` and I import both in the test class file. 

    require 'mocha/setup'
    require 'mocha/api' 

And from the documentation, I can see that parameter matching (instance\_of) is defined in the api.rb file but when I try 

    User.any_instance.stubs(:something).with(instance_of(String)).returns true

I get an error message like instance\_of is not defined

But I'm not able to include the required files to get this to work. Any help would be great :)
## [9][Successive Q and A](https://www.reddit.com/r/rails/comments/fq8r1f/successive_q_and_a/)
- url: https://www.reddit.com/r/rails/comments/fq8r1f/successive_q_and_a/
---
Trying to come up with a way for when a user is filling out the answers for multiple questions on a form to just see one at a time and when they answer, it stores the value and goes to the next question. Would this be something better done with using React or Angular, perhaps? Curious if anyone has tackled a similar issue and what route they went.
## [10][Has any of you worked on a HIPAA codebase? Advice?](https://www.reddit.com/r/rails/comments/fq49wk/has_any_of_you_worked_on_a_hipaa_codebase_advice/)
- url: https://www.reddit.com/r/rails/comments/fq49wk/has_any_of_you_worked_on_a_hipaa_codebase_advice/
---
I am about to start work on a multi-tenant HIPAA-compatible SAAS.

My understanding is that PII needs to be encrypted at rest. How do you go about this? 

Keep an encryption key for each tenant and encrypt these fields? But, if the encryption key is a part of the database which it would need to be, then how is that protection?

This in addition to heroku shield dynos and database of course. https://www.heroku.com/compliance

Or, would using the shield dyno and database be sufficient?
## [11][Using multiple parameters in scope using has_scope](https://www.reddit.com/r/rails/comments/fq5mu0/using_multiple_parameters_in_scope_using_has_scope/)
- url: https://www.reddit.com/r/rails/comments/fq5mu0/using_multiple_parameters_in_scope_using_has_scope/
---
Hi everyone,  


I'm using the [has\_scope](https://github.com/heartcombo/has_scope) gem for my scopes, and I'm trying to create a scope which takes two parameters, location and distance. I'm using the [geocoder](https://github.com/alexreisner/geocoder) gem to find venues within a certain location using the .near method.  


Here's my existing scope (which is just for location):

    scope :location, -&gt; location { near(location) }

But now want to give distance as an extra parameter.  
Here's essentially what I want to do (but does not work)

    scope :by_location_and_distance, -&gt; location, distance { near(location, distance) }

I think the reason this doesn't work is because has\_scope is looking for a param named "by\_location\_and\_distance". Which I do not have, as I have "location" and "distance" as two separate params.  


I've also tried this (which I hoped would allow the scope to access the params hash (first time I've ever used the lambda keyword so this was a complete guess - unsurprisingly it also does not work).

    scope :by_location_and_distance, lambda { |params| near(params["location"], params["distance"]) }

  
I've also tried to supply the params to the scope in the controller when calling has\_scope, like so:

    has_scope :by_location_and_distance do |controller, scope|
        scope.by_location_and_distance(params[:location], params[:distance])
      end

This also doesn't work. None of these throw an error, but it's as though the scope hasn't run.  
I'm running out of ideas a bit here. Has anyone used this gem and can explain to me how to give the scope two different params? All my other scopes only require one param so this is different to the logic I've been using up till now.  


Or even if you have advice on how I can debug this, that would be useful. I feel like the has\_scopes gem uses a lot of rails magic which is kind of working against me here, as I don't actually know how to debug it since all my scopes are being called in one line in controller#index (`@venues = apply_scopes(Venue).all`)  


Thanks.
