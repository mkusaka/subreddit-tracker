# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/hwehh6/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/hwehh6/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/i00rha/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/i00rha/personal_projects_show_off_your_own_project_andor/
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
## [3][ActiveStorage and local directory with dockerized app](https://www.reddit.com/r/rails/comments/i14po4/activestorage_and_local_directory_with_dockerized/)
- url: https://www.reddit.com/r/rails/comments/i14po4/activestorage_and_local_directory_with_dockerized/
---
I have a `docker-compose.yml` with a Rails/Postgres app which uses **ActiveStorage** and local storage setup. 

    local:
 service: Disk
 root: &lt;%= Rails.root.join("storage") %&gt;

I want all file uploads to be kept safely in the host's disk. What is the safest way of keeping those files intact while pulling/rebuilding the docker images ?
## [4]["All-to-all" association](https://www.reddit.com/r/rails/comments/i133s9/alltoall_association/)
- url: https://www.reddit.com/r/rails/comments/i133s9/alltoall_association/
---
Obviously this isn't the name for what I want to do, but I can't think of any other way to describe it.  I want to make an "all to all" association between two models.

Imagine I have two models:  A "Customer" model which is what it says on the tin, and probably has hundreds or thousands of records, and some get added every day.  And a "Warehouse" model which probably has a handful of records, and don't change much but still once in a while.

I'd like to store, in every Customer record, the cost of traveling to that customer from each warehouse.

So, as a first crack, you have something like this:

    class WarehouseCost &lt; ApplicationRecord
      belongs_to :warehouse
      belongs_to :customer
     
      # assume this model has an attribute :travel_cost
    end

    class Customer &lt; ApplicationRecord
      has_many :warehouse_costs
      accepts_nested_attributes_for :warehouse_costs
    end

Now, this can store the data, and with the right validations you can prevent it from having more than 1 record for each combination of warehouse and customer, but where I'm actually running into trouble is how to make the forms work.  This isn't a job for cocoon, which lets you arbitrarily add records.  I'd like a single input displayed for each Warehouse, regardless of whether any data is already saved in the record or not.

What I've done so far is to put a very kludgy hack in the edit action of the controller:

    Warehouse.all.each do |w|
      wc = @customer.warehouse_costs.find_by(:warehouse =&gt; w)
      @customer.warehouse_costs.create(:warehouse =&gt; w) unless wc
    end

Because the Customer now has a WarehouseCost record explicitly created for each Warehouse, the form "fields_for" works correctly and can draw an input for each one:

    &lt;%= f.fields_for :warehouse_costs do |wc| %&gt;
      &lt;b&gt;&lt;%= "Cost from " + wc.object.warehouse.name+" warehouse:" %&gt;&lt;/b&gt;
      &lt;%= wc.text_field :travel_cost, :class =&gt; 'form-control' %&gt;
    &lt;% end %&gt;

So, even though this works, it feels like I'm going about it the wrong way.  For starters, it doesn't work on new records.  Also, it feels kludgy to add the associated records before drawing the edit form, as surely I could somehow generate the correct form for the associated records without them already being there?

Any suggestions?
## [5][Rails caching in production](https://www.reddit.com/r/rails/comments/i0tv17/rails_caching_in_production/)
- url: https://www.reddit.com/r/rails/comments/i0tv17/rails_caching_in_production/
---
Could someone explain me how should I configure a rails app in production if I'm fragment caching some views, how I can configure the memcache server In \`Heroku\`, is there a step by step documentation that can help me setting it up?
## [6][Working with APIs and data in rails.](https://www.reddit.com/r/rails/comments/i0vv42/working_with_apis_and_data_in_rails/)
- url: https://www.reddit.com/r/rails/comments/i0vv42/working_with_apis_and_data_in_rails/
---
So, I'm trying to set up an app that has a feature for making payments with credit ( or debit) cards. I have been looking for tutorials that are type agnostic. Many tutorials use Stripe or paypal as examples but none really teaches the basics. Most also concentrate on building an API app in rails. I want to consume an API

I found a book by Noel Rappin "Take My Money" but it is extremely and needlessly complex (note to rails authors : if you write a book or tutorial, dont use slim or haml in the views and write so even inexperienced newcomers can follow along ...)

Can anyone direct me to a good tutorial on working with apis, data and payments in rails ?
## [7][Shopify on Rails is doing well](https://www.reddit.com/r/rails/comments/i0f8dx/shopify_on_rails_is_doing_well/)
- url: https://www.reddit.com/r/rails/comments/i0f8dx/shopify_on_rails_is_doing_well/
---
I think this is the story that we want Rails to be mentioned in, but unfortunately no one cares much about the stack. Nonetheless, it’s a Rails win to me. They’re doing well as a business, and they’re using Rails.

https://www.wsj.com/articles/shopifys-revenue-nearly-doubles-as-covid-19-pushes-shopping-online-11596057094
## [8][How do you handle real-time in your applications?](https://www.reddit.com/r/rails/comments/i076f5/how_do_you_handle_realtime_in_your_applications/)
- url: https://www.reddit.com/r/rails/comments/i076f5/how_do_you_handle_realtime_in_your_applications/
---
I can't find anyone covering this issue in a practical manner.

Let's say we have a fairly large SaaS multi tenant Rails app, with a bunch of ActiveRecord models, and a react/vue client to power the SPA front-end.

How do you approach making this app update data for all users in realtime?

I understand most articles out there show that you can use websockets to emit events to the client and listen to them on the frontend, but it's often an over-simplified view that doesn't cover:

- How to abstract out ActiveRecord data sync on both the backend and frontend? (similar to firestore data bind)
- What about race conditions when emitting the update events from activerecord? should there be versioning to avoid the possible issue of an old update event being received after the newest.

I'm asking because i built out a hackish, standardized way to emit changes from Rails via pusher on all models:

    class ApplicationRecord &lt; ActiveRecord::Base
      self.abstract_class = true
    
      after_commit :sync_payload
    
      def sync_payload
        if respond_to?("pusher_channel_name")
          channel = pusher_channel_name
          event = "#{self.class.name.underscore.dasherize}-updated"
    
          return if ch.blank? || ch.is_a?(Array) &amp;&amp; ch.empty?
          
          if channel
            if destroyed?
              push_payload(channel, event, { id: id, _destroyed: true })
            else
              push_payload(channel, event, as_json)
            end
          end
        end
      end
    end

On the client side (Vue), i built out mixin methods to listen for these events and change data.

As you can see this is subject to race conditions and it doesn't make sure events are sent in order, and there's various concerns on how reliable this is and if users will always have up-to-date data.

I'm curious to see how others approach this problem.
## [9][I have a coding test coming up, but I'm not sure what this instruction means?](https://www.reddit.com/r/rails/comments/i0am0q/i_have_a_coding_test_coming_up_but_im_not_sure/)
- url: https://www.reddit.com/r/rails/comments/i0am0q/i_have_a_coding_test_coming_up_but_im_not_sure/
---
Hi all,  


I have received.a coding test for a job where I have to create a little rails app. The instructions for sharing the code back with them are like this:

    - A link to the Git repository containing the code
    - A separate, initial commit of any automatically-generated code (For example, from rails new, so we can easily see your changes)

  
The first point is a no brainer, but I'm a bit unclear on the second point. Is someone able to clarify for me what this means? When I create a new rails app, all I do is "rails new my-project-name". Does this mean I can ignore this second point? Or am I misinterpreting it?  


I just wanted to check here before going back to them, as it's my first coding test like this so I don't want to blow it by fumbling over a minor misunderstanding.  


Thanks.
## [10][after 4.0 =&gt; 4.2 upgrade, user.authenticate fails. Any insight?](https://www.reddit.com/r/rails/comments/i080rr/after_40_42_upgrade_userauthenticate_fails_any/)
- url: https://www.reddit.com/r/rails/comments/i080rr/after_40_42_upgrade_userauthenticate_fails_any/
---
Hey guys,

just as the title said.

The user model is setting up the has_secure_password attribute, which uses bcrypt for authentication against a password_digest DB column.

The problem is that after upgrade from rails 4.0 to 4.2, the authentication is now returning false.  I've verified that the password_digest is the exact same between the working 4.0 version and the non-working 4.2 version.

Does anyone have an insight into why this would be?

If I have to I'll start diving into the rails code directly, but I'd prefer not doing that if it's just something easy I'm missing.

I'm not super rails savvy so I apologize if this is something well known, but no upgrade guides mentioned it.
## [11][Trying to use ActionCable gem locally, but getting LoadError with require statements](https://www.reddit.com/r/rails/comments/i0a19i/trying_to_use_actioncable_gem_locally_but_getting/)
- url: https://www.reddit.com/r/rails/comments/i0a19i/trying_to_use_actioncable_gem_locally_but_getting/
---
I needed to modify action cable and wanted to monkey patch a local version I placed in the lib/ folder.

I added this line to my Gemfile:

    gem 'actioncable', '5.2.3', :path =&gt; "lib/gems/actioncable-5.2.3"

But then I get an error after startup when I load a page that says  **LoadError (cannot load such file -- nio):** on actioncable-5.2.3/lib/action\_cable/connection/stream\_event\_loop.rb

If I go into rails console I can call  `require "nio"` and it works fine.

I've been following the instructions I've found for installing a gem locally [https://rubyglasses.blogspot.com/2009/12/how-to-monkey-patch-gem.html](https://rubyglasses.blogspot.com/2009/12/how-to-monkey-patch-gem.html)  (couldn't find any detailed guides, and I know it's from 2009)

Anyone have an idea what is causing this?
## [12][Add Authenticity Token to JavaScript AJAX request](https://www.reddit.com/r/rails/comments/hzvvai/add_authenticity_token_to_javascript_ajax_request/)
- url: https://www.reddit.com/r/rails/comments/hzvvai/add_authenticity_token_to_javascript_ajax_request/
---
I need to send it manually and cannot submit it through a form. Is there a specific request param that Rails looks for in order to verify the Authenticity token? 

Perhaps, I’m going at it wrong. The other idea I had was to make it an API-only route/controller and, maybe then I wouldn’t have to use the authenticity token and just use an API auth token? Sorry, I am new to Rails :D
