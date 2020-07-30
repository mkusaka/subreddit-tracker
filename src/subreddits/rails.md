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
## [3][Shopify on Rails is doing well](https://www.reddit.com/r/rails/comments/i0f8dx/shopify_on_rails_is_doing_well/)
- url: https://www.reddit.com/r/rails/comments/i0f8dx/shopify_on_rails_is_doing_well/
---
I think this is the story that we want Rails to be mentioned in, but unfortunately no one cares much about the stack. Nonetheless, it’s a Rails win to me. They’re doing well as a business, and they’re using Rails.

https://www.wsj.com/articles/shopifys-revenue-nearly-doubles-as-covid-19-pushes-shopping-online-11596057094
## [4][How do you handle real-time in your applications?](https://www.reddit.com/r/rails/comments/i076f5/how_do_you_handle_realtime_in_your_applications/)
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
## [5][I have a coding test coming up, but I'm not sure what this instruction means?](https://www.reddit.com/r/rails/comments/i0am0q/i_have_a_coding_test_coming_up_but_im_not_sure/)
- url: https://www.reddit.com/r/rails/comments/i0am0q/i_have_a_coding_test_coming_up_but_im_not_sure/
---
Hi all,  


I have received.a coding test for a job where I have to create a little rails app. The instructions for sharing the code back with them are like this:

    - A link to the Git repository containing the code
    - A separate, initial commit of any automatically-generated code (For example, from rails new, so we can easily see your changes)

  
The first point is a no brainer, but I'm a bit unclear on the second point. Is someone able to clarify for me what this means? When I create a new rails app, all I do is "rails new my-project-name". Does this mean I can ignore this second point? Or am I misinterpreting it?  


I just wanted to check here before going back to them, as it's my first coding test like this so I don't want to blow it by fumbling over a minor misunderstanding.  


Thanks.
## [6][after 4.0 =&gt; 4.2 upgrade, user.authenticate fails. Any insight?](https://www.reddit.com/r/rails/comments/i080rr/after_40_42_upgrade_userauthenticate_fails_any/)
- url: https://www.reddit.com/r/rails/comments/i080rr/after_40_42_upgrade_userauthenticate_fails_any/
---
Hey guys,

just as the title said.

The user model is setting up the has_secure_password attribute, which uses bcrypt for authentication against a password_digest DB column.

The problem is that after upgrade from rails 4.0 to 4.2, the authentication is now returning false.  I've verified that the password_digest is the exact same between the working 4.0 version and the non-working 4.2 version.

Does anyone have an insight into why this would be?

If I have to I'll start diving into the rails code directly, but I'd prefer not doing that if it's just something easy I'm missing.

I'm not super rails savvy so I apologize if this is something well known, but no upgrade guides mentioned it.
## [7][Trying to use ActionCable gem locally, but getting LoadError with require statements](https://www.reddit.com/r/rails/comments/i0a19i/trying_to_use_actioncable_gem_locally_but_getting/)
- url: https://www.reddit.com/r/rails/comments/i0a19i/trying_to_use_actioncable_gem_locally_but_getting/
---
I needed to modify action cable and wanted to monkey patch a local version I placed in the lib/ folder.

I added this line to my Gemfile:

    gem 'actioncable', '5.2.3', :path =&gt; "lib/gems/actioncable-5.2.3"

But then I get an error after startup when I load a page that says  **LoadError (cannot load such file -- nio):** on actioncable-5.2.3/lib/action\_cable/connection/stream\_event\_loop.rb

If I go into rails console I can call  `require "nio"` and it works fine.

I've been following the instructions I've found for installing a gem locally [https://rubyglasses.blogspot.com/2009/12/how-to-monkey-patch-gem.html](https://rubyglasses.blogspot.com/2009/12/how-to-monkey-patch-gem.html)  (couldn't find any detailed guides, and I know it's from 2009)

Anyone have an idea what is causing this?
## [8][Add Authenticity Token to JavaScript AJAX request](https://www.reddit.com/r/rails/comments/hzvvai/add_authenticity_token_to_javascript_ajax_request/)
- url: https://www.reddit.com/r/rails/comments/hzvvai/add_authenticity_token_to_javascript_ajax_request/
---
I need to send it manually and cannot submit it through a form. Is there a specific request param that Rails looks for in order to verify the Authenticity token? 

Perhaps, I’m going at it wrong. The other idea I had was to make it an API-only route/controller and, maybe then I wouldn’t have to use the authenticity token and just use an API auth token? Sorry, I am new to Rails :D
## [9][Why redirect_to proc?](https://www.reddit.com/r/rails/comments/hzk7uy/why_redirect_to_proc/)
- url: https://www.reddit.com/r/rails/comments/hzk7uy/why_redirect_to_proc/
---
What is the use case of redirecting to a proc?

As in 

    redirect_to proc { edit_post_url(@post) }

vs 

    redirect_to edit_post_url(@post)

?

What are some uses of this?

Edit: from https://api.rubyonrails.org/classes/ActionController/Redirecting.html
## [10][Turning a list item into an edit form](https://www.reddit.com/r/rails/comments/hzpgfb/turning_a_list_item_into_an_edit_form/)
- url: https://www.reddit.com/r/rails/comments/hzpgfb/turning_a_list_item_into_an_edit_form/
---
I have a list of employees and an edit button by each one. Instead of the edit button redirecting to a form, I want it to turn the employee's info into text boxes that are auto-filled with the current value of that attribute. The edit button would turn into a submit button. The user would make the changes needed, and after submitting, the employee's new attributes would be listed. 

Here's my understanding of how this would work. 

1. The user clicks on the edit button, which triggers a function that removes the current contents of the employee's div and replaces it with the form built as a document fragment by another function. 
2. The new submit button is a link to the update action of the employee controller with a data-remote="true" attribute. 
3. The update action sends back the new employee's instance, which somehow triggers a function that will delete the form and replace it with a new document fragment containing the updated employee information. 

My biggest concern is pulling off the Ajax call properly. I'm trying to use ujs and vanilla JavaScript, but the resources I've found for doing Ajax calls in rails use coffee script and/or jQuery. I've found a few newer articles on medium, but they don't go in enough depth for me to be able to adapt they're examples for my propose.

Do you guys know of any resources that may help me fill in the gaps?

Thanks!
## [11][Has many through association not working (ActiveRecord::HasManyThroughAssociationNotFoundError)](https://www.reddit.com/r/rails/comments/hzp1k0/has_many_through_association_not_working/)
- url: https://www.reddit.com/r/rails/comments/hzp1k0/has_many_through_association_not_working/
---
Hi I'm getting the following error when I test out either of the following associations:

    Venue.last.users
    ActiveRecord::HasManyThroughAssociationNotFoundError: Could not find the association :venue_users in model Venue

&amp;#x200B;

    User.last.venues
    ActiveRecord::HasManyThroughAssociationNotFoundError: Could not find the association :venue_users in model User

  
Here's what my models look like:

    class Venue &lt; ApplicationRecord
      has_many :users, through: :venue_users
    end

    class User &lt; ApplicationRecord
      has_many :venues, through: :venue_users
    end

  
This is what the join table looks like in my schema:

    create_table "venue_users", force: :cascade do |t|
        t.bigint "user_id", null: false
        t.bigint "venue_id", null: false
        t.integer "user_type", default: 0
        t.index ["user_id"], name: "index_venue_users_on_user_id"
        t.index ["venue_id"], name: "index_venue_users_on_venue_id"
      end

  
Can someone help me where I've gone wrong?  


Thanks.
## [12][Rails way of trimming down controller action code size?](https://www.reddit.com/r/rails/comments/hzmath/rails_way_of_trimming_down_controller_action_code/)
- url: https://www.reddit.com/r/rails/comments/hzmath/rails_way_of_trimming_down_controller_action_code/
---
I'd like to move some business logic out of the controller actions.
