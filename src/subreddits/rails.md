# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/igyvm1/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/igyvm1/personal_projects_show_off_your_own_project_andor/
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
## [2][Is it a good idea learning Rails by the book of 2017 year?](https://www.reddit.com/r/rails/comments/ikph1s/is_it_a_good_idea_learning_rails_by_the_book_of/)
- url: https://www.reddit.com/r/rails/comments/ikph1s/is_it_a_good_idea_learning_rails_by_the_book_of/
---
I am reading Ruby on Rails Tutorial book by Michael Hartl (2017) now. I am feeling some troubles with Ruby version and Gems version. 

So, is it a good idea or I need to find some newer resource? What resource you can recommend?

P.S. I went from Python &amp; Django. I got new project and part of it wrote on Ruby. Stack: Django, Rails, React - looks like Frankenstein. :)
## [3][Email from contact form in Rails 6](https://www.reddit.com/r/rails/comments/iky3nn/email_from_contact_form_in_rails_6/)
- url: https://www.reddit.com/r/rails/comments/iky3nn/email_from_contact_form_in_rails_6/
---
Hello. I'm adding a contact form in my web development portfolio. I would like people to be able to send me a message to my email directly from this form. I'm trying to find a way to do this with Rails 6 and I'm not finding anything new. While Googlilng, I notice that the JS has changed slightly in Rails. It's moved from assets to it's own folder because of webpacker. I'm such a newbie so I don't know much about why that's important. I'm pretty sure I have to add the "mail\_form", "simple\_form", "jquery-rails", and I guess "bootstrap\_sass" gems. Now I'm just wondering what to do next.

Older ways I see are basically adding a controller for the contact form. However, I my form will just be at the bottom of my homepage. Most of my portfolio will just be on this page, so when people are done looking at it, I want them to be able to just scroll to the bottom and see the form there, then message me if that's how they want to contact me quickly.  If someone could point me in the right direction of doing this for Rails 6, I'd appreciate it!
## [4][building REST APIs with Rails and MongoDB](https://www.reddit.com/r/rails/comments/ikw6q5/building_rest_apis_with_rails_and_mongodb/)
- url: https://www.reddit.com/r/rails/comments/ikw6q5/building_rest_apis_with_rails_and_mongodb/
---
Hello everybody

Any good resources for building APIs with Rails and MongoDB? 

PS: I have tried a course on Coursera but it was outdated.
## [5][transaction_include_any_action or use :create or :update in model concern](https://www.reddit.com/r/rails/comments/ikmq5u/transaction_include_any_action_or_use_create_or/)
- url: https://www.reddit.com/r/rails/comments/ikmq5u/transaction_include_any_action_or_use_create_or/
---
For Concern does this make more sense in rails 4.2.1

    included do
      after_commit :update_actions, on: :update
      after_commit :create_action, on: :create
    end
    
    def update_actions
    end
    
    def create_action
    end

Or does this make more sense?

    included do
      after_commit :after_commit_actions, on: [:create, :update]
    end
    
    def after_commit_actions
     if(send(:transaction_include_any_action?, [:create]))
       //do create thigs
     elsif (send(:transaction_include_any_action?, [:update])
       do update things
     end

I'm more favored towards the first approach because

1. it doesn't look like we are checking the same thing two times.
2. \`transaction\_include\_any\_action\` protected and in rails 5.1.7 it is being converted to a private method. I don't feel good about calling protected and private methods like this.

What do you think about both approaches? Which is better and why?
## [6][will_paginate returns nil](https://www.reddit.com/r/rails/comments/iklzzl/will_paginate_returns_nil/)
- url: https://www.reddit.com/r/rails/comments/iklzzl/will_paginate_returns_nil/
---
Hi! Just came across this weird problem:

`will_paginate` will always return `nil`, no matter what.

My controller action looks like this:

```ruby
        page = params[:page].to_i
        if page &lt; 1
            page = 1
        end

        @posts = Post.paginate page: page, per_page: 10
```

And the corresponding view like this:

```erb
&lt;%= will_paginate @posts %&gt;
```

However, the pagination is not rendered. Trying to inspect it through the CLI gives me `"nil"`. Why is that? How can I fix it? Thank you!

Edit: I know, will paginate is not part of Rails, but it's commonly used with it, so I thought I might give it a try here. Sorry if it doesn't really fit
## [7][API architecture design for fast reads with 150 million records](https://www.reddit.com/r/rails/comments/ik0sbe/api_architecture_design_for_fast_reads_with_150/)
- url: https://www.reddit.com/r/rails/comments/ik0sbe/api_architecture_design_for_fast_reads_with_150/
---
I have a text file with 150 million unique records.

Each record has two columns: (1) a string and (2) an integer. The string has a unique label, and the integer is that label's value.

There's only a single query available: return the integer value for any given label.

This text file is regenerated every 72 hours. ~90% of the data remains the same across regeneration, but this regeneration is controlled by a 3rd party. I simply get a new text file every 72 hours.

I'm exploring multiple architectures for exposing this text file as an API. I want to use Ruby/Rails.

Ideally, a query shouldn't take more than 100 - 500ms (per read).

## Architecture 1

* Store the text file on disk. Query the text file. Cache queries in memory.
* Pros: Simple implementation. Easy to update data.
* Cons: Uncached read queries are slow.

## Architecture 2

* Parse the text file into a traditional/NoSQL database, with each line treated as a database record/document. Run queries against the database.
* Pros: Seems like the common architecture.
* Cons: Updating 150m database records is slow and seems wasteful, especially since ~90% of records remain the same.

## Architecture 3

* Use Redis or in-memory database to store the 5GB text file. Run queries against the in-memory database.
* Pros: Fast queries. Easy to update data.
* Cons: Expensive.

## Architecture 4

* Use ElasticSearch to query records.
* Pros: ElasticSearch is designed for search.
* Cons: ElasticSearch may be overkill for such simple queries.

Questions:

* Can you suggest any other approaches (if any)?

* Are there additional pros/cons I overlooked?

* What is the most "common" architecture for balancing cost/performance when trying to produce fast reads against a data store (of 150m) records that change?
## [8][Starting a new app... public and admin best approach for today?](https://www.reddit.com/r/rails/comments/ik0ikq/starting_a_new_app_public_and_admin_best_approach/)
- url: https://www.reddit.com/r/rails/comments/ik0ikq/starting_a_new_app_public_and_admin_best_approach/
---
Some quick context. Long time, small, Rails development team. We've done the same thing the same way for many years. Usually admin based monolith with public facing API. Most of our work is pure business management.

Now we have a project that would have several hundred "public" uses but really the main crux of the project is still the backend business side for the "staff" users. Sort of a CRM+CMS + Event management. (Nothing we haven't done before in parts.) Public would access their personal history and sign-up for events. But the admin side has lots of additional business functions that would not be exposed to other users.

My team is small and our views seem narrow on this, so I'm reaching out for ideas.

I see a couple of options:

**All-In**: There is no real admin side, just expose more options for logged in users who have permissions.

**Admin First**: Since this project is mostly about the business domain start here and add client access (public side) on top of it. Scoping /admin

**The Core**: Build the core functionality into a gem that is shared by a Public Engine, and an Admin Engine. (sort of the Spree concept)

We've done each of these with success, granted the Engine route has the most friction for our team. We are used to a "dashboard" concept for admin, so the All-In approach is uncomfortable but also exciting to explore. Our fallback has always been Admin First.

&amp;#x200B;

What other approaches am I missing?
## [9][How to create a form that saves 2 records for the same model?](https://www.reddit.com/r/rails/comments/ijva9y/how_to_create_a_form_that_saves_2_records_for_the/)
- url: https://www.reddit.com/r/rails/comments/ijva9y/how_to_create_a_form_that_saves_2_records_for_the/
---
Hello,

I am building a small personal finance app.  I have models Accounts and Transactions.

I want to build a feature to transfer money between 2 accounts, by way of creating the appropriate transactions.

I have created a "transfer" controller action in my AccountController, and added some attr\_accessors to my Account model:

```ruby
  attr_accessor :transfer_to_account
  attr_accessor :transfer_amount
```

I'm getting stuck on how to build the form for this, because I want the fields to be:

* transfer_to_account (which should auto populate a dropdown based on all of the user's accounts)
* transfer_amount

The "Source_account" will be pre-determined based on which account the user is in. Then the controller action should create a debit transaction in the source account (fk), and a credit transaction in the target account (fk).

The part I'm stuck on is how to 
a. specify the form_for code and fields
b. populate a dropdown with all current accounts (to use in the "transfer_to_account" field

Thanks!
## [10][How does Devise authentication works?](https://www.reddit.com/r/rails/comments/ijke0i/how_does_devise_authentication_works/)
- url: https://www.reddit.com/r/rails/comments/ijke0i/how_does_devise_authentication_works/
---
Hi, guys. I have some question about how  Devise authentication works. The question are in the sequence:

- I have notice that the session cookie changes for every new request. From what I have read, Devise doesn't hold session on server, only on client. If so, why the session cookie is changing constantly? If no value is saved on server, Devise has no way to invalidate a session key - only changing the password would invalidate the session key. Or there a session saved on server for every login user? 

- Based on the question above, and if there are no session saved on server, how is it considerate different than the main approach (using only a access token passed on header bearer token) of OAuth protocol? Or it isn't?
## [11][Heroku - Deploy Broken](https://www.reddit.com/r/rails/comments/ijjqgk/heroku_deploy_broken/)
- url: https://www.reddit.com/r/rails/comments/ijjqgk/heroku_deploy_broken/
---
I just merged a PR that changed a few things:

  - An application feature (setting up attachments on a model via ActiveStorage)
  - Upgrading to `Rails 6` from `Rails 5.2`
  - Update circle ci config to use newer version of bundler

I suspect it was upgrading to Rails 6. Heroku says the app has been deployed successfully but I get a 503.

When I try loading the app, I get the 503 and I see my logs have:


    2020-08-30T20:00:42.831439+00:00 heroku[web.1]: Starting process with command `bundle exec puma -C config/puma.rb`
    2020-08-30T20:00:46.014937+00:00 app[web.1]: [heroku-exec] Starting
    2020-08-30T20:00:46.991311+00:00 app[web.1]: [4] Puma starting in cluster mode...
    2020-08-30T20:00:46.992288+00:00 app[web.1]: [4] * Version 3.12.6 (ruby 2.5.1-p57), codename: Llamas in Pajamas
    2020-08-30T20:00:46.992345+00:00 app[web.1]: [4] * Min threads: 5, max threads: 5
    2020-08-30T20:00:46.992391+00:00 app[web.1]: [4] * Environment: production
    2020-08-30T20:00:46.992475+00:00 app[web.1]: [4] * Process workers: 2
    2020-08-30T20:00:46.992519+00:00 app[web.1]: [4] * Preloading application
    2020-08-30T20:00:54.223474+00:00 app[web.1]: [4] * Listening on tcp://0.0.0.0:16562
    2020-08-30T20:00:54.223966+00:00 app[web.1]: [4] ! WARNING: Detected 1 Thread(s) started in app boot:
    2020-08-30T20:00:54.224082+00:00 app[web.1]: [4] ! #&lt;Thread:0x000055b93cbae460@/app/vendor/bundle/ruby/2.5.0/bundler/gems/rails-7101489293ec/activerecord/lib/active_record/connection_adapters/abstract/connection_pool.rb:332 sleep&gt; - /app/vendor/bundle/ruby/2.5.0/bundler/gems/rails-7101489293ec/activerecord/lib/active_record/connection_adapters/abstract/connection_pool.rb:335:in `sleep'
    2020-08-30T20:00:54.224251+00:00 app[web.1]: [4] Use Ctrl-C to stop
    2020-08-30T20:00:54.288824+00:00 app[web.1]: /app/vendor/bundle/ruby/2.5.0/bundler/gems/rails-7101489293ec/activesupport/lib/active_support/fork_tracker.rb:8:in `fork': super: no superclass method `fork' for #&lt;Puma::Cluster:0x000055b93a8947b8&gt; (NoMethodError)
    2020-08-30T20:00:54.288894+00:00 app[web.1]: from /app/vendor/bundle/ruby/2.5.0/gems/puma-3.12.6/lib/puma/cluster.rb:139:in `block in spawn_workers'
    2020-08-30T20:00:54.288964+00:00 app[web.1]: from /app/vendor/bundle/ruby/2.5.0/gems/puma-3.12.6/lib/puma/cluster.rb:135:in `times'
    2020-08-30T20:00:54.288987+00:00 app[web.1]: from /app/vendor/bundle/ruby/2.5.0/gems/puma-3.12.6/lib/puma/cluster.rb:135:in `spawn_workers'
    2020-08-30T20:00:54.289009+00:00 app[web.1]: from /app/vendor/bundle/ruby/2.5.0/gems/puma-3.12.6/lib/puma/cluster.rb:468:in `run'
    2020-08-30T20:00:54.289028+00:00 app[web.1]: from /app/vendor/bundle/ruby/2.5.0/gems/puma-3.12.6/lib/puma/launcher.rb:186:in `run'
    2020-08-30T20:00:54.289047+00:00 app[web.1]: from /app/vendor/bundle/ruby/2.5.0/gems/puma-3.12.6/lib/puma/cli.rb:80:in `run'
    2020-08-30T20:00:54.289066+00:00 app[web.1]: from /app/vendor/bundle/ruby/2.5.0/gems/puma-3.12.6/bin/puma:10:in `&lt;top (required)&gt;'
    2020-08-30T20:00:54.289084+00:00 app[web.1]: from /app/vendor/bundle/ruby/2.5.0/bin/puma:23:in `load'
    2020-08-30T20:00:54.289103+00:00 app[web.1]: from /app/vendor/bundle/ruby/2.5.0/bin/puma:23:in `&lt;main&gt;'
    

I had also seen this similar error message earlier:
    
    2020-08-30T19:07:18.337125+00:00 app[web.1]: [4] * Listening on tcp://0.0.0.0:47116
    2020-08-30T19:07:18.337144+00:00 app[web.1]: [4] ! WARNING: Detected 1 Thread(s) started in app boot:
    2020-08-30T19:07:18.337184+00:00 app[web.1]: [4] ! #&lt;Thread:0x0000562860d548a8@/app/vendor/bundle/ruby/2.5.0/bundler/gems/rails-7101489293ec/activerecord/lib/active_record/connection_adapters/abstract/connection_pool.rb:332 sleep&gt; - /app/vendor/bundle/ruby/2.5.0/bundler/gems/rails-7101489293ec/activerecord/lib/active_record/connection_adapters/abstract/connection_pool.rb:335:in `sleep'
    2020-08-30T19:07:18.337270+00:00 app[web.1]: [4] Use Ctrl-C to stop
    2020-08-30T19:07:18.337560+00:00 app[web.1]: bundler: failed to load command: puma (/app/vendor/bundle/ruby/2.5.0/bin/puma)
    2020-08-30T19:07:18.388901+00:00 app[web.1]: NoMethodError: super: no superclass method `fork' for #&lt;Puma::Cluster:0x000056285cc5b928&gt;
    2020-08-30T19:07:18.388903+00:00 app[web.1]: /app/vendor/bundle/ruby/2.5.0/bundler/gems/rails-7101489293ec/activesupport/lib/active_support/fork_tracker.rb:8:in `fork'
    2020-08-30T19:07:18.388904+00:00 app[web.1]: /app/vendor/bundle/ruby/2.5.0/gems/puma-3.12.6/lib/puma/cluster.rb:139:in `block in spawn_workers'
    2020-08-30T19:07:18.388904+00:00 app[web.1]: /app/vendor/bundle/ruby/2.5.0/gems/puma-3.12.6/lib/puma/cluster.rb:135:in `times'
    2020-08-30T19:07:18.388904+00:00 app[web.1]: /app/vendor/bundle/ruby/2.5.0/gems/puma-3.12.6/lib/puma/cluster.rb:135:in `spawn_workers'
    2020-08-30T19:07:18.388905+00:00 app[web.1]: /app/vendor/bundle/ruby/2.5.0/gems/puma-3.12.6/lib/puma/cluster.rb:468:in `run'
    2020-08-30T19:07:18.388905+00:00 app[web.1]: /app/vendor/bundle/ruby/2.5.0/gems/puma-3.12.6/lib/puma/launcher.rb:186:in `run'
    2020-08-30T19:07:18.388906+00:00 app[web.1]: /app/vendor/bundle/ruby/2.5.0/gems/puma-3.12.6/lib/puma/cli.rb:80:in `run'
    2020-08-30T19:07:18.388906+00:00 app[web.1]: /app/vendor/bundle/ruby/2.5.0/gems/puma-3.12.6/bin/puma:10:in `&lt;top (required)&gt;'
    2020-08-30T19:07:18.388907+00:00 app[web.1]: /app/vendor/bundle/ruby/2.5.0/bin/puma:23:in `load'
    2020-08-30T19:07:18.388907+00:00 app[web.1]: /app/vendor/bundle/ruby/2.5.0/bin/puma:23:in `&lt;top (required)&gt;'
    2020-08-30T19:07:18.492258+00:00 heroku[web.1]: Process exited with status 1
    2020-08-30T19:07:18.547903+00:00 heroku[web.1]: State changed from starting to crashed
    

Which changed to the first block pasted above when I [set an env var like mentioned in this post](https://github.com/puma/puma/issues/1572#issuecomment-411288015).

When I load up a rails console, I see the ActiveStorage migration that was a part of this PR went out successfully because I see `service_name` available on `ActiveStorage::Blob` :

    [3] pry(main)&gt; ActiveStorage::Blob
    =&gt; ActiveStorage::Blob(id: integer, key: string, filename: string, content_type: string, metadata: text, byte_size: integer, checksum: string, created_at: datetime, service_name: string)
