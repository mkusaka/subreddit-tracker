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
## [2][ActionText iframe embeds](https://www.reddit.com/r/rails/comments/fqj15c/actiontext_iframe_embeds/)
- url: https://www.reddit.com/r/rails/comments/fqj15c/actiontext_iframe_embeds/
---
Hi, 

I am currently building some kind of inventory tracking system and I am using ActionText to capture rich text descriptions for products.

It would be great if I could just paste iframe embed and view it on `show` page. It doesn't need to be transformed to embed immediately in editor. I know this can be security risk, but this is in dashboard and only admins can edit this content so is there a way to disable ActionText html sanitization for `&lt;iframe&gt;` html tag?

Currently `@product.description.body.to_trix_html` outputs santized `&lt;iframe&gt;` tags.

    &amp;lt;iframe width=\"560\"...

Thanks r/rails for helping me, I have learned a lot about rails from answers from this community.

Bartol
## [3][When do you decide to rewrite a project?](https://www.reddit.com/r/rails/comments/fq9hhh/when_do_you_decide_to_rewrite_a_project/)
- url: https://www.reddit.com/r/rails/comments/fq9hhh/when_do_you_decide_to_rewrite_a_project/
---
I am revisiting an old rails app that I was developing, but the stack I used is completely different and the code quality and design is pretty bad. I'm considering just doing a complete rewrite. When do you guys say screw it and rewrite it?
## [4][using instance_of with mocha for ActionController::TestCase](https://www.reddit.com/r/rails/comments/fqez0r/using_instance_of_with_mocha_for/)
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
## [5][Successive Q and A](https://www.reddit.com/r/rails/comments/fq8r1f/successive_q_and_a/)
- url: https://www.reddit.com/r/rails/comments/fq8r1f/successive_q_and_a/
---
Trying to come up with a way for when a user is filling out the answers for multiple questions on a form to just see one at a time and when they answer, it stores the value and goes to the next question. Would this be something better done with using React or Angular, perhaps? Curious if anyone has tackled a similar issue and what route they went.
## [6][Has any of you worked on a HIPAA codebase? Advice?](https://www.reddit.com/r/rails/comments/fq49wk/has_any_of_you_worked_on_a_hipaa_codebase_advice/)
- url: https://www.reddit.com/r/rails/comments/fq49wk/has_any_of_you_worked_on_a_hipaa_codebase_advice/
---
I am about to start work on a multi-tenant HIPAA-compatible SAAS.

My understanding is that PII needs to be encrypted at rest. How do you go about this? 

Keep an encryption key for each tenant and encrypt these fields? But, if the encryption key is a part of the database which it would need to be, then how is that protection?

This in addition to heroku shield dynos and database of course. https://www.heroku.com/compliance

Or, would using the shield dyno and database be sufficient?
## [7][Using multiple parameters in scope using has_scope](https://www.reddit.com/r/rails/comments/fq5mu0/using_multiple_parameters_in_scope_using_has_scope/)
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
## [8][Pragmatic Studo - Latest Courses?](https://www.reddit.com/r/rails/comments/fq42cb/pragmatic_studo_latest_courses/)
- url: https://www.reddit.com/r/rails/comments/fq42cb/pragmatic_studo_latest_courses/
---
Hey ya'll, apologies if I'm in the wrong forum, but I just completed the Pragmatic Studio course in Rails 6 and it was really great.

They also offer a Ruby language course, and a special course on "Mastering Ruby Blocks".  They also offer new courses on Elixir/OTP and Phoenix.

Has anyone tried any of these other courses out?  I would imagine that they are equally good but just wanted to make sure since they are quite a bit more expensive than other training sites.

https://pragmaticstudio.com/
## [9][Dynamic Feed That Pulls Data from multiple sources in Rails](https://www.reddit.com/r/rails/comments/fq2ytc/dynamic_feed_that_pulls_data_from_multiple/)
- url: https://www.reddit.com/r/rails/comments/fq2ytc/dynamic_feed_that_pulls_data_from_multiple/
---
I was taking a look at the Grailed website - [https://www.grailed.com](https://www.grailed.com) and was wondering what's the best way to replicate something like this in rails. Specifically, the home page; in the way it's a dynamic feed that pulls data from a number of different sources; I assume like a CMS and a SQL database of some sort.

I really have no idea how to start building something like this or where to look, so any little bit of guidance (even if it's what to search for on google) would be great.
## [10][No luck changing activeRecord naming convention with 'authorized'](https://www.reddit.com/r/rails/comments/fq0j4y/no_luck_changing_activerecord_naming_convention/)
- url: https://www.reddit.com/r/rails/comments/fq0j4y/no_luck_changing_activerecord_naming_convention/
---
I'm trying to create a table named \`authorized\`. Naturally, active record changes it to the plural \`authorizeds\`. I've tried over riding this with \`self.table\_name = 'authroized' but no luck when actually migrating the tables (\`rails db:migrate\`). Any other ways around this?
## [11][Database - is there any benefit to index a list of names?](https://www.reddit.com/r/rails/comments/fpv4h6/database_is_there_any_benefit_to_index_a_list_of/)
- url: https://www.reddit.com/r/rails/comments/fpv4h6/database_is_there_any_benefit_to_index_a_list_of/
---
Say I have a model Company, it has one single :name column. The purpose of this model is just to store company names. The only use for this model is to create a select list when needed (i.e Select All), and to **find by id** when a company name is required (i.e I won't search the Company model directly). The user can manage the company list (create/delete/etc).

Is there any reason to (or not to) index the :name column?

---
Edit: Thanks guys your answers are very helpful!

A little summary:

- Add index when you need to search the column, of course.
- Add index when you need to sort a column e.g sort by company name in this case.
- Add index can be helpful if the column needs to be unique.
- If the model is simple and will be used very often (like the company name in this case), it doesn't hurt to index the column early on. Because you may need to search or sort the column later.
