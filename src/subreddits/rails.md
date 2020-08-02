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
## [3][[Help]Rounding problems when trying to calculate taxes on order](https://www.reddit.com/r/rails/comments/i1uhyu/helprounding_problems_when_trying_to_calculate/)
- url: https://www.reddit.com/r/rails/comments/i1uhyu/helprounding_problems_when_trying_to_calculate/
---
Hi, I'm trying to calculate the total cost for a given set of fixed orders inclusive of tax(which should be rounded up to the nearest 0.05. Despite trying both .round and .ceil, I'm still slightly off the target values for each of them. I was wondering if anyone could identify what I am missing in my orders controller?

    module Api
      class V1::OrdersController &lt; ActionController::API	
        def calculate
    			calculated = {}
    			calculated["sales_tax"] = 0
    			params["item"].values.each do |item|
    				calculated["#{item["quantity"]} #{item["name"]}"] = item["quantity"].to_f * item["price"].to_f
    		
    				unless ["book", "chocolate bar", "imported box of chocolates", "packet of headache pills"].include? item["name"].strip
    					calculated["#{item["quantity"]} #{item["name"]}"] = ('%.2f' %(calculated["#{item["quantity"]} #{item["name"]}"]*1.10)).to_f
    					calculated["sales_tax"] += item["quantity"].to_f * item["price"].to_f * 0.10
    				end
    
    				if item["name"].split(" ").include? "imported"
    					calculated["#{item["quantity"]} #{item["name"]}"] = ('%.2f' % (calculated["#{item["quantity"]} #{item["name"]}"]*1.05)).to_f
    					calculated["sales_tax"] += item["quantity"].to_f * item["price"].to_f * 0.05
    				end
    			end
    			
    			calculated["total"] =('%.2f' % (calculated.values.sum - calculated["sales_tax"])).to_f
    			calculated["sales_tax"] = ('%.2f' % calculated["sales_tax"]).to_f
    			render json: calculated, status: 200
    		end
    	end
    end
## [4][How do I access "current_user" in my RegistrationsController?](https://www.reddit.com/r/rails/comments/i1x41i/how_do_i_access_current_user_in_my/)
- url: https://www.reddit.com/r/rails/comments/i1x41i/how_do_i_access_current_user_in_my/
---
Hey,  


I'm trying to access current\_user" in my RegistrationsController, but it's returning nil.  


The only answer I could find online was [this one](https://stackoverflow.com/questions/11369941/rails-devise-current-user-is-nil-when-overriding-registrationscontroller) suggesting to add the line:  


    include Devise::Controllers::Helpers

But this didn't work for me.  


Any ideas?  


Thanks.
## [5][Can't verify CSRF token authenticity Error](https://www.reddit.com/r/rails/comments/i1sd6c/cant_verify_csrf_token_authenticity_error/)
- url: https://www.reddit.com/r/rails/comments/i1sd6c/cant_verify_csrf_token_authenticity_error/
---
Hello,  


I'm building a rails API and I'm having the problem of getting the following error message:  


    Can't verify CSRF token authenticity.

  
This is happening on my UsersController when I try to UPDATE a user. Even though I have the following lines in the controller (I've also tried them in ApplicationController).  


    skip_before_action :verify_authenticity_token
    protect_from_forgery with: :null_session

Having these lines has resolved the issue in my other controllers, but not in the users controller. Can anyone help me with this?  


Thanks.
## [6][Rails API mode with native app: How to handle social provider JWT authentication.](https://www.reddit.com/r/rails/comments/i1ofh8/rails_api_mode_with_native_app_how_to_handle/)
- url: https://www.reddit.com/r/rails/comments/i1ofh8/rails_api_mode_with_native_app_how_to_handle/
---
Hi guys

I've been searching for the best way to handle authentication (without breaking the bank) in my latest project, being a native app with a Rails backend.  


I looked at Auth0, Firebase, Okta, FusionAuth but after reading an extensive amount of threads on here regarding the different solutions, I was wondering if just using ruby-jwt would actually be enough. This is a production app so I am looking for something robust but I'm a bit lost on why I would need an external provider or a specific gem if I am just using social providers and returning JWT's.   


But do let me know what you think and where my ideas may be wrong!  


Regards
## [7][Auto generated migrations like in Django?](https://www.reddit.com/r/rails/comments/i1m8yy/auto_generated_migrations_like_in_django/)
- url: https://www.reddit.com/r/rails/comments/i1m8yy/auto_generated_migrations_like_in_django/
---
With Django, I just modify models.py and then maybe rename a field and then the migration is generated. I’m wondering if there is a way to do this with Rails migrations. I’ve heard of Squasher but it seems unmaintained.
## [8][Most popular database for Rails?](https://www.reddit.com/r/rails/comments/i1dx5k/most_popular_database_for_rails/)
- url: https://www.reddit.com/r/rails/comments/i1dx5k/most_popular_database_for_rails/
---
I would assume PostgreSQL because I know many Rails developers like to use Heroku which seems to favor PostgreSQL in terms of pricing.
## [9][From Internal to External CSS](https://www.reddit.com/r/rails/comments/i1nt6b/from_internal_to_external_css/)
- url: https://www.reddit.com/r/rails/comments/i1nt6b/from_internal_to_external_css/
---
On my website I have all the css files in the folder \\app\\assets\\stylesheets (in .scss, obviously).

And in the head I have this:

    &lt;% if Rails.env.production? %&gt;
    	&lt;style type="text/css" media="all"&gt;&lt;%= inline_asset('application.css') %&gt;&lt;/style&gt;
    &lt;% else %&gt;
    	&lt;%= stylesheet_link_tag 'application', media: 'all' %&gt;
    &lt;% end %&gt;
    &lt;%= csrf_meta_tags %&gt;

In this way, watching the source of my online website, I see that the CSS is showed as "internal".

How to turn it as external? I know that this is the best practice.
## [10][Creating a new route to a separate index](https://www.reddit.com/r/rails/comments/i1949h/creating_a_new_route_to_a_separate_index/)
- url: https://www.reddit.com/r/rails/comments/i1949h/creating_a_new_route_to_a_separate_index/
---
i've read and re-read the guides and I don't see how to distinguish between the different get requests.

If I have the following:

resources :widgets

   get     :separate\_index

end

&amp;#x200B;

The get becomes equivalent to a show b/c it routes with an id.

I need it to route as an index style get.

&amp;#x200B;

I'm sorry, i realize this might not be clear, but i don't know how else to explain it.
## [11][ActiveStorage and local directory with dockerized app](https://www.reddit.com/r/rails/comments/i14po4/activestorage_and_local_directory_with_dockerized/)
- url: https://www.reddit.com/r/rails/comments/i14po4/activestorage_and_local_directory_with_dockerized/
---
I have a `docker-compose.yml` with a Rails/Postgres app which uses **ActiveStorage** and local storage setup. 

    local:
 service: Disk
 root: &lt;%= Rails.root.join("storage") %&gt;

I want all file uploads to be kept safely in the host's disk. What is the safest way of keeping those files intact while pulling/rebuilding the docker images ?
## [12]["All-to-all" association](https://www.reddit.com/r/rails/comments/i133s9/alltoall_association/)
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
