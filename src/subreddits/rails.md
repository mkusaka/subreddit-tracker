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
## [2][API architecture design for fast reads with 150 million records](https://www.reddit.com/r/rails/comments/ik0sbe/api_architecture_design_for_fast_reads_with_150/)
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
## [3][Starting a new app... public and admin best approach for today?](https://www.reddit.com/r/rails/comments/ik0ikq/starting_a_new_app_public_and_admin_best_approach/)
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
## [4][How to create a form that saves 2 records for the same model?](https://www.reddit.com/r/rails/comments/ijva9y/how_to_create_a_form_that_saves_2_records_for_the/)
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
## [5][How does Devise authentication works?](https://www.reddit.com/r/rails/comments/ijke0i/how_does_devise_authentication_works/)
- url: https://www.reddit.com/r/rails/comments/ijke0i/how_does_devise_authentication_works/
---
Hi, guys. I have some question about how  Devise authentication works. The question are in the sequence:

- I have notice that the session cookie changes for every new request. From what I have read, Devise doesn't hold session on server, only on client. If so, why the session cookie is changing constantly? If no value is saved on server, Devise has no way to invalidate a session key - only changing the password would invalidate the session key. Or there a session saved on server for every login user? 

- Based on the question above, and if there are no session saved on server, how is it considerate different than the main approach (using only a access token passed on header bearer token) of OAuth protocol? Or it isn't?
## [6][Heroku - Deploy Broken](https://www.reddit.com/r/rails/comments/ijjqgk/heroku_deploy_broken/)
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
## [7][How do you share the master.key with GitHub collaborators](https://www.reddit.com/r/rails/comments/ijd66h/how_do_you_share_the_masterkey_with_github/)
- url: https://www.reddit.com/r/rails/comments/ijd66h/how_do_you_share_the_masterkey_with_github/
---
Hello. So I created a new Rails project and I am working with another developer on it. I pushed the project to GitHub but, of course, the `master.key` isn't sent to GitHub. How do both of us safely handle this so there are no conflicts? Do I just send the other dev the `master.key`? We need it for credentials.
## [8][Submitting a get request via ajax does not trigger csv download](https://www.reddit.com/r/rails/comments/ijfotf/submitting_a_get_request_via_ajax_does_not/)
- url: https://www.reddit.com/r/rails/comments/ijfotf/submitting_a_get_request_via_ajax_does_not/
---
Basically I have a get(verb) form, where has remote true and on submission, I am rendering a partial via .js.erb.  


Now I would like to add export as csv button, with the filters selected in the above form.   


My options are   
`&lt;%= link_to "CSV", filters_reports_path(format: "csv") %&gt;`  


But this does not give the filters selected previously.  


So I have added a button, and on click I am getting the form data and sending the request like this  

```
&lt;button id="download_csv"&gt; Download &lt;/button&gt; 

  $('#download_csv').on("click", function (e) {
    e.preventDefault()
    $.ajax({type: "GET",
      url: "/reports/filters.csv",
      data: $('#reports_filter').serialize(),
      success:function(result) {
        console.log('ok');
      },
      error:function(result) {
        alert('error');
      }
    });
  })
```

```
--&gt; log for link_to
Started GET "/reports/filters.csv" for ::1 at 2020-08-30 21:54:28 +0530
  User Load (0.9ms)  SELECT "users".* FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" ASC LIMIT $2  [["id", "4c0f729f-ee51-4e38-9c65-86042e709f78"], ["LIMIT", 1]]
Processing by ReportsController#filters as CSV
Completed 200 OK in 1ms (ActiveRecord: 0.0ms | Allocations: 220)


  TimeEntry Load (1.2ms)  SELECT "time_entries".* FROM "time_entries" ORDER BY "time_entries"."id" ASC LIMIT $1  [["LIMIT", 1000]]
  ↳ app/controllers/reports_controller.rb:43:in `block (3 levels) in filters'



-----------
-&gt; log for ajax request
Started GET "/reports/filters.csv?project_id=&amp;daterange=23%2FAug%2F2020%20-%2030%2FAug%2F2020" for ::1 at 2020-08-30 21:54:45 +0530
  User Load (1.4ms)  SELECT "users".* FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" ASC LIMIT $2  [["id", "4c0f729f-ee51-4e38-9c65-86042e709f78"], ["LIMIT", 1]]
Processing by ReportsController#filters as CSV
  Parameters: {"project_id"=&gt;"", "daterange"=&gt;"23/Aug/2020 - 30/Aug/2020"}
Completed 200 OK in 1ms (ActiveRecord: 0.0ms | Allocations: 222)
```

My question is if there a way to get the filters in the link_to itself other than storing in session? or what needs to be done to get the ajax request trigger download
## [9][Rails 5.2.4, Active Storage, and S3: is it possible to implement an effective IAM/ACL policy when the app's default opinion dumps everything into the root of a single bucket?](https://www.reddit.com/r/rails/comments/ij5e6v/rails_524_active_storage_and_s3_is_it_possible_to/)
- url: https://www.reddit.com/r/rails/comments/ij5e6v/rails_524_active_storage_and_s3_is_it_possible_to/
---
I have switched from Paperclip to Active Storage. While Active Storage is great, it presents a new set of problems. Paperclip let me set a bucket directory (or even a new bucket), whereas Active Storage dumps everything into the root of a single bucket. I fear that this could create security issues. Improperly set IAM policies are an easy to make no-no, and my app had been using two buckets to avoid the problem: one for public content, and one for private admin accessible content.

In making that two bucket setup work with Active Storage, I tried to use the `service` option for `has_one_attached`. However, I get an ArgumentError: `unknown keyword: service`. The [edge guide](https://edgeguides.rubyonrails.org/active_storage_overview.html#setup) implies that this option works in Rails 5.2, but the [guide for 5.2.4](https://guides.rubyonrails.org/v5.2.4/active_storage_overview.html#has-one-attached) doesn't mention the `service` option at *all*.

How do you lock down access to certain S3 files while being less restrictive with everything else? Is this even possible? Do I have to move to Shrine?
## [10][Beginner ActiveStorage Help](https://www.reddit.com/r/rails/comments/ij4hh0/beginner_activestorage_help/)
- url: https://www.reddit.com/r/rails/comments/ij4hh0/beginner_activestorage_help/
---
# SOLUTION:
See this SO post: https://stackoverflow.com/questions/58373159/unknown-attribute-service-name-for-activestorageblob


# Routes

    
    # routes.rb
      resources :properties,          only: [:new, :create, :edit, :update, :index, :show] do
        resources :expense_items
      end
    

# Model

    # model
    class ExpenseItem &lt; ApplicationRecord
      has_many_attached :attachments
    end

    # Works fine in rails console:
    @expense_item.attachments.map { |f| f.filename } 
    # =&gt;[#&lt;ActiveStorage::Filename:0x000056033874b940 @filename="test1.png"&gt;, 
      #&lt;ActiveStorage::Filename:0x0000560338748d08 @filename="test2.mov"&gt;, 
      #&lt;ActiveStorage::Filename:0x000056033893e540 @filename="test3.pdf"&gt;] 


# Error
But when I try to use it in my edit view, I get an error saying:
    
    undefined method `service_name' for #&lt;ActiveStorage::Blob:0x00007fdd4844eaa8&gt;
    Did you mean?  service

    
# View

    # edit view
      &lt;div class="col-sm-6"&gt;
        &lt;%= render 'shared/error_messages', object: @expense_item %&gt;
          &lt;ul&gt;
            &lt;% @expense_item.attachments.each do |file| %&gt;
              &lt;li&gt;
                &lt;% if file.previewable? %&gt;
                  &lt;p&gt;file.filename&lt;/p&gt;
                  &lt;%#= image_tag file.preview(resize_to_limit: [100,100]) %&gt;
                &lt;% end %&gt;
              &lt;/li&gt;
            &lt;% end %&gt;
          &lt;/ul&gt;      &lt;%= f.submit 'Update Expense', class: 'btn btn-primary mt-sm-3' %&gt;
        &lt;% end %&gt;

The exception gets thrown at the if file.previewable? line.

# Controller

    # expense items controller
      def show
        @expense_item = property.expense_items.find(params[:id])
      end  

      def edit
        @expense_item = property.expense_items.find(params[:id])
      end
  
      def update
        @expense_item = property.expense_items.find(params[:id])
        if @expense_item&amp;.update_attributes(expense_item_params)
          flash[:success] = "Updated Expense Item: #{@expense_item.name}"
          redirect_to @expense_item.property
        else
          render 'edit'
        end
      end
    
      def expense_item_params
        params.require(:expense_item).permit(:name, :cost, :expense_date, :property_id, attachments: [])
      end
## [11][Javascript + Rails](https://www.reddit.com/r/rails/comments/ij4gvx/javascript_rails/)
- url: https://www.reddit.com/r/rails/comments/ij4gvx/javascript_rails/
---
I need help guys, i have to do a form where i select a commune (a big list of this) and, when i have to selectthe condominium  it only shows me those of the selected commune. I try with onChange(Javascript), but i dont know how get this value and use in the condominium place.

&amp;#x200B;

I really hope you can understand to me, is difficult to me explain this in english
