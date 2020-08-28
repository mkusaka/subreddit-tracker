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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/igyvm1/personal_projects_show_off_your_own_project_andor/)
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
## [3][Why is my update controller action changing ID of model record belongs_to?](https://www.reddit.com/r/rails/comments/ii0upw/why_is_my_update_controller_action_changing_id_of/)
- url: https://www.reddit.com/r/rails/comments/ii0upw/why_is_my_update_controller_action_changing_id_of/
---
Odd one here. I have some nested resources, and nested forms. On new, everything saves just fine. However, on update, even with no changes, one of my IDs updates and I can't figure out why. I'll try to include everything relevant.

Model ClientSite

    class ClientSite &lt; ApplicationRecord
      include Hashid::Rails
      has_many :sale_vehicle_specials, dependent: :destroy
      accepts_nested_attributes_for :sale_vehicle_specials, allow_destroy: true
    end

Model SaleVehicleSpecial

    class SaleVehicleSpecial &lt; ApplicationRecord
      include Hashid::Rails
      belongs_to :client_site
      has_one :sale_special_detail, dependent: :destroy, inverse_of: :sale_vehicle_special
      
      validates :client_site, presence: true
      
      accepts_nested_attributes_for :sale_special_detail
    end

Controller SaleVehicleSpecial

    class SaleVehicleSpecialsController &lt; ApplicationController
      before_action :set_special_site
      before_action :set_sale_vehicle_special
    
      def edit
      end
    
      def update
        respond_to do |format|
          if @sale_vehicle_special.update_attributes(sale_vehicle_special_params)
            format.html { redirect_to client_site_sale_vehicle_specials_path(@site), notice: 'Special was successfully updated.' }
            format.json { render :show, status: :ok, location: client_site_sale_vehicle_specials_path(@site) }
          else
            format.html { render :edit }
            format.json { render json: @sale_vehicle_special.errors, status: :unprocessable_entity }
          end
        end
      end
    
      private
        def set_special_site
          @site = ClientSite.find_by_hashid(params[:client_site_id])
        end
        
        def set_sale_vehicle_special
          @sale_vehicle_special = @site.sale_vehicle_specials.find_by_hashid(params[:id])
        end
    
        def sale_vehicle_special_params
          params.require(:sale_vehicle_special).permit(:client_site_id, 
                        sale_special_detail_attributes: [ :id, :sale_vehicle_special_id,  :_destroy] )
        end
    end
    

Form \_form\_edit.html.erb

    &lt;%= form_with(model: [@site, @sale_vehicle_special], local: true) do |form| %&gt;
    
      &lt;div class="form-group"&gt;
        &lt;%= form.label :client_site_id, 'Dealership Name:' %&gt;
        
        &lt;% assigned_names = Assignment.where("user_id = " + current_user.id.to_s).map { |site| [ClientSite.find_by(id: site.client_site_id).client_site_name, ClientSite.find_by(id: site.client_site_id).hashid] } %&gt;
        &lt;%= form.select(:client_site_id, options_for_select(assigned_names, selected: @site.hashid), {prompt: "Select Client"}, {class: "client-site form-control"}) %&gt;
      &lt;/div&gt;
    
      &lt;h2&gt;Vehicle Details&lt;/h2&gt;
      &lt;%= form.fields_for :sale_special_detail do |ff| %&gt;
        ...
      &lt;% end %&gt;
    
      &lt;div class="actions"&gt;
        &lt;%= form.submit "Save Special", class: 'btn btn-large btn-primary' %&gt;
        &lt;%= link_to 'Cancel', client_site_sale_vehicle_specials_path(@site), {class: 'btn btn-small btn-outline-secondary ml-1'} %&gt;
      &lt;/div&gt;
    &lt;% end %&gt;
    

Yes, I know some of this code shouldn't be in the controller, but I'm wondering if somehow that select is causing my problem...it just seems unlikely.

I have noticed that on form submission for new, if there is an error with validating any of the fields I get an error `Client site is invalid` which goes away  when I fix the other errors.

However, here's the problem. When I update, whether or not I've changed anything, my client\_site\_id updates.

New submission

    Started POST "/client_sites/7YpTPB/sale_vehicle_specials" for ::1 at 2020-08-28 00:35:14 -0400
       (0.7ms)  SELECT "schema_migrations"."version" FROM "schema_migrations" ORDER BY "schema_migrations"."version" ASC
    Processing by SaleVehicleSpecialsController#create as HTML
      Parameters: {"authenticity_token"=&gt;"q08yPsYCjHaa9ty/uSxbXSnftiQBU4wtSAkTuOdCeDvoXyKG4FjcoRj0Iyync4GXxPv7d+zknYPXc1KJgPCGpw==", "sale_vehicle_special"=&gt;{"client_site_id"=&gt;"7YpTPB", "campaign_end_date(1i)"=&gt;"2020", "campaign_end_date(2i)"=&gt;"8", "campaign_end_date(3i)"=&gt;"1", "sale_special_detail_attributes"=&gt;{"stock_no"=&gt;"F10067", "vin"=&gt;"H78JJ3HDY789EOKH4"}}}, "commit"=&gt;"Save Special", "client_site_id"=&gt;"7YpTPB"}
      ClientSite Load (0.2ms)  SELECT "client_sites".* FROM "client_sites" WHERE "client_sites"."id" = $1 LIMIT $2  [["id", 28], ["LIMIT", 1]]
      ↳ app/controllers/sale_vehicle_specials_controller.rb:119:in `set_special_site'
      SaleVehicleSpecial Load (0.5ms)  SELECT "sale_vehicle_specials".* FROM "sale_vehicle_specials" WHERE "sale_vehicle_specials"."client_site_id" = $1 AND "sale_vehicle_specials"."id" IS NULL LIMIT $2  [["client_site_id", 28], ["LIMIT", 1]]
      ↳ app/controllers/sale_vehicle_specials_controller.rb:123:in `set_sale_vehicle_special'
      ClientSite Load (0.4ms)  SELECT "client_sites".* FROM "client_sites" WHERE "client_sites"."id" = $1 LIMIT $2  [["id", 28], ["LIMIT", 1]]
      ↳ app/controllers/sale_vehicle_specials_controller.rb:127:in `set_brand_id'
      SaleVehicleSpecial Create (2.4ms)  INSERT INTO "sale_vehicle_specials" ("created_at", "updated_at", "client_site_id", "campaign_end_date") VALUES ($1, $2, $3, $4) RETURNING "id"  [["created_at", "2020-08-28 04:35:15.171281"], ["updated_at", "2020-08-28 04:35:15.171281"], ["client_site_id", 28], ["campaign_end_date", "2020-08-01"]]
      ↳ app/controllers/sale_vehicle_specials_controller.rb:82:in `block in create'
      SaleVehicleSpecial Update (0.5ms)  UPDATE "sale_vehicle_specials" SET "updated_at" = $1 WHERE "sale_vehicle_specials"."id" = $2  [["updated_at", "2020-08-28 04:35:15.202998"], ["id", 5]]
      ↳ app/controllers/sale_vehicle_specials_controller.rb:82:in `block in create'
       (0.8ms)  COMMIT
      ↳ app/controllers/sale_vehicle_specials_controller.rb:82:in `block in create'
    Redirected to http://localhost:3000/client_sites/7YpTPB/sale_vehicle_specials
    Completed 302 Found in 273ms (ActiveRecord: 57.2ms | Allocations: 110534)

And update

    Started PATCH "/client_sites/7YpTPB/sale_vehicle_specials/jmzSWB" for ::1 at 2020-08-28 00:41:11 -0400
    Processing by SaleVehicleSpecialsController#update as HTML
      Parameters: {"authenticity_token"=&gt;"m/NdQvRiACbI81HEI6eGdKgtCYuA6VYrnv4RbPbIjR84Yxx7cIiHE9N00s8iLEY4YjXAeH7beQDp7h/B+Sf7qg==", "sale_vehicle_special"=&gt;{"client_site_id"=&gt;"7YpTPB", "campaign_end_date(1i)"=&gt;"2020", "campaign_end_date(2i)"=&gt;"8", "campaign_end_date(3i)"=&gt;"1", "sale_special_detail_attributes"=&gt;{"stock_no"=&gt;"F10067", "vin"=&gt;"H78JJ3HDY789EOKH4", "id"=&gt;"5"}}}, "commit"=&gt;"Save Special", "client_site_id"=&gt;"7YpTPB", "id"=&gt;"jmzSWB"}
      ClientSite Load (0.5ms)  SELECT "client_sites".* FROM "client_sites" WHERE "client_sites"."id" = $1 LIMIT $2  [["id", 28], ["LIMIT", 1]]
      ↳ app/controllers/sale_vehicle_specials_controller.rb:119:in `set_special_site'
      SaleVehicleSpecial Load (0.4ms)  SELECT "sale_vehicle_specials".* FROM "sale_vehicle_specials" WHERE "sale_vehicle_specials"."client_site_id" = $1 AND "sale_vehicle_specials"."id" = $2 LIMIT $3  [["client_site_id", 28], ["id", 5], ["LIMIT", 1]]
      ↳ app/controllers/sale_vehicle_specials_controller.rb:123:in `set_sale_vehicle_special'
    DEPRECATION WARNING: update_attributes is deprecated and will be removed from Rails 6.1 (please, use update instead) (called from block in update at /app/controllers/sale_vehicle_specials_controller.rb:96)
       (0.2ms)  BEGIN
      ↳ app/controllers/sale_vehicle_specials_controller.rb:96:in `block in update'
      SaleSpecialDetail Load (0.3ms)  SELECT "sale_special_details".* FROM "sale_special_details" WHERE "sale_special_details"."sale_vehicle_special_id" = $1 LIMIT $2  [["sale_vehicle_special_id", 5], ["LIMIT", 1]]
      ↳ app/controllers/sale_vehicle_specials_controller.rb:96:in `block in update'

Here's the problem line

      SaleVehicleSpecial Update (4.6ms)  UPDATE "sale_vehicle_specials" SET "client_site_id" = $1, "updated_at" = $2 WHERE "sale_vehicle_specials"."id" = $3  [["client_site_id", 7], ["updated_at", "2020-08-28 04:41:11.356725"], ["id", 5]]
      ↳ app/controllers/sale_vehicle_specials_controller.rb:96:in `block in update'

Rest of the update logs

       (1.4ms)  COMMIT
      ↳ app/controllers/sale_vehicle_specials_controller.rb:96:in `block in update'
    Redirected to http://localhost:3000/client_sites/7YpTPB/sale_vehicle_specials
    Completed 302 Found in 156ms (ActiveRecord: 15.1ms | Allocations: 18699)

(Obviously) I stripped out everything that I don't think is relevent. But, no matter what, everytime I update one of these SaleVehicleSpecials the client\_site\_id updates to 7. Doesn't matter what it is to start, always updates to 7. I can't find any hardcoded values anywhere.

I've been scratching my head over this for hours and I'm at a loss.
## [4][Do you still use SJR responses in your app?](https://www.reddit.com/r/rails/comments/ihubfa/do_you_still_use_sjr_responses_in_your_app/)
- url: https://www.reddit.com/r/rails/comments/ihubfa/do_you_still_use_sjr_responses_in_your_app/
---
I’m wondering if this pattern is still popular amongst new Rails applications. I’ve avoided using it up to now since I always thought it was quite messy. I just used a stimulus’s controller and handled the ajax:success, error etc myself.
## [5][Rails server as a websocket client](https://www.reddit.com/r/rails/comments/ihrx23/rails_server_as_a_websocket_client/)
- url: https://www.reddit.com/r/rails/comments/ihrx23/rails_server_as_a_websocket_client/
---
I want to build a slack bot server and want to connect to the real time API using a websocket connection. this means rails would be connecting to slack websocket server. actioncable deosnt apply here. 

I know there are repos for ruby slack bots and ruby slack bot server but those use faye and grape, and i want to use rails. i know using rails as a websocket client isn’t convention but i would rather manage the connection myself and have all the benefits of rails than use grape/faye/eventmachine.

I don’t want to use thin either because (1) couldn’t find docs about hooking into the EM reactor and (2) i don’t want the websocket logic to impact incoming http requests. 

could i create some sort of singleton class that makes a connection on boot? is that a footgun?
## [6][Job Opportunity | Fully Remote | Sr Ruby Dev with Shopify Experience](https://www.reddit.com/r/rails/comments/ihldu0/job_opportunity_fully_remote_sr_ruby_dev_with/)
- url: https://www.reddit.com/r/rails/comments/ihldu0/job_opportunity_fully_remote_sr_ruby_dev_with/
---
I am a Ruby developer and I am looking for a Sr Ruby on Rails developer who has experience working with Shopify and integrating with other third party API's. 

&amp;#x200B;

This is a fully remote and temporary position   ( 3+ months ). 

&amp;#x200B;

I need someone who is comfortable working in a rather large scale private app, is able to make local changes without breaking things, and of course has GitHub experience.

&amp;#x200B;

Please DM me if you are interested and I'd be happy to share more details. Cheers!
## [7]["Dual-purpose APIs" -- Does this make sense?](https://www.reddit.com/r/rails/comments/ihcfz4/dualpurpose_apis_does_this_make_sense/)
- url: https://www.reddit.com/r/rails/comments/ihcfz4/dualpurpose_apis_does_this_make_sense/
---
I have looked around and I haven't really found any blog posts or "received knowledge" about dual-purpose APIs.

By "dual-purpose", I mean this:  I have a management application that has a normal web interface, but it also has a smartphone companion app for doing a certain subset of procedures that benefit greatly from being able to scan barcodes using the phone's camera.  So, there's an API for the smartphone app to connect, log on, perform the operations once the barcodes have been scanned, etc.

But there is also a need for an API for modern front-end web interfaces.  I'm thinking specifically of things like type-ahead controls on forms, "query the server for a list of cities to populate a dropdown after the user has chosen a country", that sort of thing.

In the past, I have used Grape to design my APIs, and done some "hacks" to be able to use either the current rack session (which would exist if there is a user interacting with the web app, and needs to query that data) or use non-session-based token authentication for when the smartphone app is making requests.  I like that Grape makes it easy to add swagger documentation for the API.

I'm not really **unhappy**, per se, with my solution, but as I'm starting a new app, I thought I'd check and see if there's some other different thing that would make more sense, before getting too far into it.  I'd be interested in hearing how others have approached a similar problem.

Thanks!
## [8][Small script to reduce the image size of Ruby and Ruby on Rails Docker images](https://www.reddit.com/r/rails/comments/igxju5/small_script_to_reduce_the_image_size_of_ruby_and/)
- url: https://www.reddit.com/r/rails/comments/igxju5/small_script_to_reduce_the_image_size_of_ruby_and/
---
In the recent months I was migrating one of my Rails applications I was maintaining in the past years from capistrano to Docker. I did know that the gems are leaving files behind and therefore the Docker images became quite large but I was shocked when I realized how big the difference is.

So I sat down and I wrote a small gem called [cleanup\_vendor](https://github.com/raszi/cleanup_vendor) which cleans up the leftover files and reduces the Docker image size significantly.

Comments and suggestions are welcome.
## [9][Interview Questions &amp; Algorithm Mastery (or lack thereof)](https://www.reddit.com/r/rails/comments/ih0ikx/interview_questions_algorithm_mastery_or_lack/)
- url: https://www.reddit.com/r/rails/comments/ih0ikx/interview_questions_algorithm_mastery_or_lack/
---
Hey guys, so I just had a very disappointing experience yesterday with my first major Ruby on Rails interview.

It was actually my 2nd interview and I felt I did a great job in the first interview, and I had a github and some live sites with examples of my work in Rails 6 (along with earlier versions) and React and other JS frameworks/libraries.

In the first round, I was able to answer most of the basic development and experience questions confidently and honestly and it went great.

However, round 2 went downhill pretty fast.

There were some other developers brought in and I was asked to do live coding for some mathematical algorithm solutions.  The person leading the interview told me there would be a series of these and that this first one was really easy but tests a lot of different computer science concepts.

It had something to do with a missing data set and calculating a rolling average over a given time period.  I was able to "think outloud" as I considered that I would need some sort of loop and a couple of variables to keep track of which time period I was currently on - outside the loop - and then increment that within the loop, and remove the previous dataset from the array since it was rolling.

But... I'm a very nervous person by nature and don't do well with coding in front of people.  So I totally blew it and couldn't come up with the solution within the minute or so I was given.

The coding aspect of the interview stopped there and I wasn't given anymore questions about actual code and we talked about processes and development strategies - which I think I did well on.

I was very upfront in the first interview about how I wasn't an algorithm master and would need time to think about things and look up possible solutions online.

In the past I've done things like Project Euler and CodeWars where I learn algorithms and I've coded solutions for finding duplicates in arrays, finding the largest number - both looping and recursive, using binary search, bubble sort, etc.

But not being able to come up with a solution on the spot cost me my first Rails job.  They spoke to the recruiter and passed on me.

I'm not a "natural born" programmer and mathematical algorithms do not come natural to me.  Am I just in the wrong business, or can you be a Rails developer without necessarily needing to be able to do stuff like this on the spot?
## [10][Rails heroku: Is it possible to select worker dyno type instead of the default?](https://www.reddit.com/r/rails/comments/ih9445/rails_heroku_is_it_possible_to_select_worker_dyno/)
- url: https://www.reddit.com/r/rails/comments/ih9445/rails_heroku_is_it_possible_to_select_worker_dyno/
---
My web application processes documents uploaded by users using sidekiq on heroku. Less than 5% of documents are too large and my heroku workers die without warning due to out of memory issues.

Is it possible to boot different worker dynos in sidekiq based on the upload size?

For example, one can specify one-off dyno types by using --size

    heroku run rails c --size=standard-2x rake heavy:job 
    heroku run rails c --size=performance-l rake heavy:job

Another possible solution is maybe loading fewer gems in my workers? (doesn't seem possible)
## [11][What is this rails server log error? (attached)](https://www.reddit.com/r/rails/comments/ihb00l/what_is_this_rails_server_log_error_attached/)
- url: https://www.reddit.com/r/rails/comments/ihb00l/what_is_this_rails_server_log_error_attached/
---
As of 5 days ago, my rails server started spewing errors in my log to the point where I am getting Heroku L10 issues non-stop. I have no idea how to fix it since I don't understand the error message (the app still seems to run fine.)

Anybody with suggestions?

&amp;#x200B;

`WARNING: V8 isolate was forked, it can not be disposed and memory will not be reclaimed till the Ruby process exits.`

`/app/vendor/bundle/ruby/2.7.0/gems/parallel-1.19.2/lib/parallel.rb:445: [BUG] Segmentation fault at 0x00007f8521f109d0`

`ruby 2.7.1p83 (2020-03-31 revision a0c7c23c9c) [x86_64-linux]`

&amp;#x200B;

`-- Machine register context ------------------------------------------------`

 `RIP: 0x00007f8581f73bd8 RBP: 0x00007f8569eb02d0 RSP: 0x00007f8569eb0250`

 `RAX: 0x0000000000000000 RBX: 0x00007f8521f10700 RCX: 0x0000000000000001`

 `RDX: 0x0000000000000000 RDI: 0x00007f8521f10700 RSI: 0x0000000000000000`

  `R8: 0x00000000000000ca  R9: 0x00007f8522f119d0 R10: 0x000000000000000d`

 `R11: 0x00007f85238693c0 R12: 0x00005640dd835740 R13: 0x00005640dd095a28`

 `R14: 0x00005640dd835750 R15: 0x00005640dd0959f0 EFL: 0x0000000000010206`

&amp;#x200B;

`-- C level backtrace information -------------------------------------------`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_vm_bugreport+0x7ce) [0x7f8582a4402e] vm_dump.c:755`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_bug_for_fatal_signal+0xe7) [0x7f8582863207] error.c:658`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(sigsegv+0x4b) [0x7f85829a807b] signal.c:946`

`/lib/x86_64-linux-gnu/libc.so.6(0x7f85823e5fd0) [0x7f85823e5fd0]`

`/lib/x86_64-linux-gnu/libpthread.so.0(__GI___pthread_timedjoin_ex+0x28) [0x7f8581f73bd8]`

`/app/vendor/bundle/ruby/2.7.0/gems/sq_mini_racer-0.2.5.0.1.beta3/lib/sq_mini_racer_extension.so(_ZN2v88platform12WorkerThreadD0Ev+0x1c) [0x7f85232b945c]`

`/app/vendor/bundle/ruby/2.7.0/gems/sq_mini_racer-0.2.5.0.1.beta3/lib/sq_mini_racer_extension.so(0x7f85232b777e) [0x7f85232b777e]`

`/app/vendor/bundle/ruby/2.7.0/gems/sq_mini_racer-0.2.5.0.1.beta3/lib/sq_mini_racer_extension.so(0x7f85232b666a) [0x7f85232b666a]`

`/app/vendor/bundle/ruby/2.7.0/gems/sq_mini_racer-0.2.5.0.1.beta3/lib/sq_mini_racer_extension.so(0x7f85232b683e) [0x7f85232b683e]`

`/lib/x86_64-linux-gnu/libc.so.6(0x7f85823ea0f1) [0x7f85823ea0f1]`

`/lib/x86_64-linux-gnu/libc.so.6(0x431ea) [0x7f85823ea1ea]`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(ruby_stop+0x10) [0x7f8582870cc0] eval.c:290`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(ruby_run_node) (null):0`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_f_fork+0x74) [0x7f85829576e4] process.c:4130`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_call_cfunc_with_frame+0xc9) [0x7f8582a206ba] vm_insnhelper.c:2514`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_call_cfunc) vm_insnhelper.c:2539`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_sendish+0x58) [0x7f8582a2bded] vm_insnhelper.c:4023`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_exec_core) insns.def:782`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_vm_exec+0xab) [0x7f8582a3211b] vm.c:1920`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(invoke_iseq_block_from_c+0x15b) [0x7f8582a3527e] vm.c:1116`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(invoke_block_from_c_bh) vm.c:1134`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_yield) vm.c:1179`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_yield_0) vm_eval.c:1227`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_yield_values2) vm_eval.c:1273`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(each_with_index_i+0x7d) [0x7f8582852b2d] enum.c:2365`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_yield_with_cfunc+0x115) [0x7f8582a28025] vm_insnhelper.c:3220`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(invoke_block_from_c_bh+0x2c) [0x7f8582a32f34] vm.c:1139`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_yield) vm.c:1179`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_yield_0) vm_eval.c:1227`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_yield_1) vm_eval.c:1233`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_yield) vm_eval.c:1243`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_array_len+0x0) [0x7f85827ce8ac] array.c:2135`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_ary_each) array.c:2134`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_call0_cfunc_with_frame+0x10c) [0x7f8582a3656c] vm_eval.c:91`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_call0_cfunc) vm_eval.c:105`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_call0_body) vm_eval.c:140`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_vm_call0+0xbf) [0x7f8582a36c1f] vm_eval.c:52`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_vm_call_kw+0x6d) [0x7f8582a36ebd] vm_eval.c:268`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(iterate_method+0x39) [0x7f8582a37949] vm_eval.c:718`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_iterate0+0xd1) [0x7f8582a25211] vm_eval.c:1415`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_block_call+0x54) [0x7f8582a253c4] vm_eval.c:1480`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(enum_each_with_index+0x44) [0x7f858284bd04] enum.c:2395`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_call_cfunc_with_frame+0xc9) [0x7f8582a206ba] vm_insnhelper.c:2514`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_call_cfunc) vm_insnhelper.c:2539`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_sendish+0x58) [0x7f8582a2bded] vm_insnhelper.c:4023`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_exec_core) insns.def:782`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_vm_exec+0xab) [0x7f8582a3211b] vm.c:1920`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(invoke_iseq_block_from_c+0x98) [0x7f8582a32e0e] vm.c:1116`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(invoke_block_from_c_bh) vm.c:1134`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_yield) vm.c:1179`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_yield_0) vm_eval.c:1227`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_yield_1) vm_eval.c:1233`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_yield) vm_eval.c:1243`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(each_slice_i+0x6f) [0x7f8582852a0f] enum.c:2517`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_yield_with_cfunc+0x115) [0x7f8582a28025] vm_insnhelper.c:3220`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(invoke_block_from_c_bh+0x2c) [0x7f8582a32f34] vm.c:1139`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_yield) vm.c:1179`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_yield_0) vm_eval.c:1227`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_yield_1) vm_eval.c:1233`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_yield) vm_eval.c:1243`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_array_len+0x0) [0x7f85827ce8ac] array.c:2135`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_ary_each) array.c:2134`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_call0_cfunc_with_frame+0x10c) [0x7f8582a3656c] vm_eval.c:91`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_call0_cfunc) vm_eval.c:105`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_call0_body) vm_eval.c:140`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_vm_call0+0xbf) [0x7f8582a36c1f] vm_eval.c:52`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_vm_call_kw+0x6d) [0x7f8582a36ebd] vm_eval.c:268`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(iterate_method+0x39) [0x7f8582a37949] vm_eval.c:718`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_iterate0+0xd1) [0x7f8582a25211] vm_eval.c:1415`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_block_call+0x54) [0x7f8582a253c4] vm_eval.c:1480`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(enum_each_slice+0x9d) [0x7f858284b91d] enum.c:2579`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_call_cfunc_with_frame+0xc9) [0x7f8582a206ba] vm_insnhelper.c:2514`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_call_cfunc) vm_insnhelper.c:2539`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_sendish+0x58) [0x7f8582a2bded] vm_insnhelper.c:4023`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_exec_core) insns.def:782`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_vm_exec+0xab) [0x7f8582a3211b] vm.c:1920`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(invoke_iseq_block_from_c+0x98) [0x7f8582a32e0e] vm.c:1116`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(invoke_block_from_c_bh) vm.c:1134`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_yield) vm.c:1179`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_yield_0) vm_eval.c:1227`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_yield_1) vm_eval.c:1233`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_yield) vm_eval.c:1243`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_array_len+0x0) [0x7f85827ce8ac] array.c:2135`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_ary_each) array.c:2134`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_call_cfunc_with_frame+0xc9) [0x7f8582a206ba] vm_insnhelper.c:2514`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_call_cfunc) vm_insnhelper.c:2539`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_sendish+0x58) [0x7f8582a2bded] vm_insnhelper.c:4023`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_exec_core) insns.def:782`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_vm_exec+0xab) [0x7f8582a3211b] vm.c:1920`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(invoke_iseq_block_from_c+0xa7) [0x7f8582a34c28] vm.c:1116`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(invoke_block_from_c_bh) vm.c:1134`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_yield) vm.c:1179`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_yield_0) vm_eval.c:1227`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(catch_i) vm_eval.c:2228`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_catch_protect+0xae) [0x7f8582a2567e] vm_eval.c:2310`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_catch_obj+0x2c) [0x7f8582a2577c] vm_eval.c:2336`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_call_cfunc_with_frame+0xc9) [0x7f8582a206ba] vm_insnhelper.c:2514`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_call_cfunc) vm_insnhelper.c:2539`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_sendish+0x58) [0x7f8582a2bded] vm_insnhelper.c:4023`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_exec_core) insns.def:782`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_vm_exec+0xab) [0x7f8582a3211b] vm.c:1920`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_ec_vm_ptr+0x0) [0x7f8582a35b3d] vm.c:1074`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_vm_global_hooks) vm_core.h:1932`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(invoke_bmethod) vm.c:1076`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(invoke_iseq_block_from_c) vm.c:1119`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(invoke_block_from_c_proc) vm.c:1216`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_vm_invoke_bmethod) vm.c:1245`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_call_bmethod+0xac) [0x7f8582a3c42c] vm_insnhelper.c:2570`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_sendish+0xa2) [0x7f8582a2bd36] vm_insnhelper.c:4023`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_exec_core) insns.def:801`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_vm_exec+0xab) [0x7f8582a3211b] vm.c:1920`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_ec_vm_ptr+0x0) [0x7f8582a35b3d] vm.c:1074`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_vm_global_hooks) vm_core.h:1932`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(invoke_bmethod) vm.c:1076`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(invoke_iseq_block_from_c) vm.c:1119`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(invoke_block_from_c_proc) vm.c:1216`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_vm_invoke_bmethod) vm.c:1245`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_call_bmethod+0xac) [0x7f8582a3c42c] vm_insnhelper.c:2570`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_sendish+0xa2) [0x7f8582a2bd36] vm_insnhelper.c:4023`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_exec_core) insns.def:801`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_vm_exec+0xab) [0x7f8582a3211b] vm.c:1920`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_ec_vm_ptr+0x0) [0x7f8582a35b3d] vm.c:1074`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_vm_global_hooks) vm_core.h:1932`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(invoke_bmethod) vm.c:1076`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(invoke_iseq_block_from_c) vm.c:1119`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(invoke_block_from_c_proc) vm.c:1216`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_vm_invoke_bmethod) vm.c:1245`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_call_bmethod+0xac) [0x7f8582a3c42c] vm_insnhelper.c:2570`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_sendish+0xa2) [0x7f8582a2bd36] vm_insnhelper.c:4023`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_exec_core) insns.def:801`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_vm_exec+0x74f) [0x7f8582a327bf] vm.c:1929`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_ec_vm_ptr+0x0) [0x7f8582a35b3d] vm.c:1074`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_vm_global_hooks) vm_core.h:1932`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(invoke_bmethod) vm.c:1076`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(invoke_iseq_block_from_c) vm.c:1119`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(invoke_block_from_c_proc) vm.c:1216`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_vm_invoke_bmethod) vm.c:1245`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_call_bmethod+0xac) [0x7f8582a3c42c] vm_insnhelper.c:2570`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_sendish+0xa2) [0x7f8582a2bd36] vm_insnhelper.c:4023`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_exec_core) insns.def:801`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_vm_exec+0x74f) [0x7f8582a327bf] vm.c:1929`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(invoke_iseq_block_from_c+0x183) [0x7f8582a360e3] vm.c:1116`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(invoke_block_from_c_proc) vm.c:1216`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(vm_invoke_proc) vm.c:1238`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_vm_invoke_proc) vm.c:1259`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(thread_do_start+0x289) [0x7f85829ebd99] thread.c:697`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(thread_start_func_2+0x257) [0x7f85829ee117] thread.c:745`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(rb_native_cond_initialize+0x0) [0x7f85829ee74b] thread_pthread.c:969`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(register_cached_thread_and_wait) thread_pthread.c:1021`

`/app/vendor/ruby-2.7.1/bin/../lib/libruby.so.2.7(thread_start_func_1) thread_pthread.c:976`

`/lib/x86_64-linux-gnu/libpthread.so.0(0x76db) [0x7f8581f726db]`

`/lib/x86_64-linux-gnu/libc.so.6(clone+0x3f) [0x7f85824c8a3f]`

&amp;#x200B;

`-- Other runtime information -----------------------------------------------`

&amp;#x200B;

`* Process memory map:`

&amp;#x200B;

`150d77480000-150d77500000 rw-p 00000000 00:00 0` 

`17f601700000-17f601780000 rw-p 00000000 00:00 0` 

`1926f3080000-1926f3100000 rw-p 00000000 00:00 0` 

`259cb9880000-259cb9881000 rw-p 00000000 00:00 0` 

`278c9d400000-278c9d480000 rw-p 00000000 00:00 0` 

`3164e3400000-3164e3480000 rw-p 00000000 00:00 0` 

`3752d3c00000-3752d3c80000 rw-p 00000000 00:00 0` 

`38cde2c80000-38cde2c9d000 rw-p 00000000 00:00 0` 

`38fb691da000-38fb69200000 ---p 00000000 00:00 0` 

`38fb69200000-38fb69201000 rw-p 00000000 00:00 0` 

`38fb69201000-38fb69202000 ---p 00000000 00:00 0` 

&amp;#x200B;
## [12][Setting up a blog on AWS Cloud9](https://www.reddit.com/r/rails/comments/igw3bw/setting_up_a_blog_on_aws_cloud9/)
- url: https://www.reddit.com/r/rails/comments/igw3bw/setting_up_a_blog_on_aws_cloud9/
---
Hi everyone,

I'm new to programming, and i've been following the upskill course for about 2 weeks now. Everything was going great untill yesterday. I started the Deep Dive: Build a Blog section that uses Cloud9 and Ruby to build a blog. I have tried many times to follow the exact same path as the instructor but i keep on having the same problem.

The commands that I input in the terminal are the following (I do this on a new environment with the default settings).

    $ rails new blog

Then the instructor ask us to change the 'sqlite3' line in the file called "gemfile" into gem 'sqlite3', '1.3.13'. To use the same version as him.

I then inpunt :

    cd
    
    cd environment
    
    cd blog
    
    bundle install
    
    bundle update
    
    rails generate scaffold Post title:string body:text
    
    rake db:migrate
    
    rails server

It is at this point that the problem occurs. The instructor terminal's response 3 more lines that i don't have [https://imgur.com/a/rAx4gCZ](https://imgur.com/a/rAx4gCZ). In my case the output stops at Ctrl c to shut down.

Plus after that the instructor click preview running application and is sent to the "ruby welcome page". He then get rid of a part of the URL to go to the blog. In my case the URL is completely different from his, it looks like that [https://c187d78accd944209c8f91023e991d71.vfs.cloud9.us-east-2.amazonaws.com/](https://c187d78accd944209c8f91023e991d71.vfs.cloud9.us-east-2.amazonaws.com/) (actual URL). and his like that [https://imgur.com/a/ghdq9wd](https://imgur.com/a/ghdq9wd)

He then get rid of the gray part of the URl and add "/posts" to be redirected to the blog.

Do you guys know what I've been doing wrong ?

Thank you for reading all of these it's my first time posting so i hope it's kind of understandable.

Have a great day
