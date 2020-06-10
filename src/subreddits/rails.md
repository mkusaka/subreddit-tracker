# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gvu083/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gvu083/personal_projects_show_off_your_own_project_andor/
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
## [2][Issue with Joining to Another Model in Admin Namespace](https://www.reddit.com/r/rails/comments/h0a27t/issue_with_joining_to_another_model_in_admin/)
- url: https://www.reddit.com/r/rails/comments/h0a27t/issue_with_joining_to_another_model_in_admin/
---
I have products which belong to a category. Categories have many products. Both of these are in an admin namespace. I am trying to show the category name on the product show page when you look at the product but haven't been able to get it to work. What am I doing wrong here?

Category Model:

    class Admin::Category &lt; ApplicationRecord
      has_many :products
    end

Product Model:

    class Admin::Product &lt; ApplicationRecord
      belongs_to :category
      has_many_attached :images
    end

Products Controller

    class Admin::ProductsController &lt; Admin::ApplicationController
      before_action :set_admin_product, only: [:show, :edit, :update, :destroy]
      breadcrumb 'Products', :admin_products_path
      # GET /admin/products
      # GET /admin/products.json
      def index
        @admin_products = Admin::Product.all
      end
    
      # GET /admin/products/1
      # GET /admin/products/1.json
      def show
        @admin_product = Admin::Product.joins(:admin_category)
      end
    
      # GET /admin/products/new
      def new
        @admin_product = Admin::Product.new
      end
    
      # GET /admin/products/1/edit
      def edit
        @pagetitle = "Edit Product"
        breadcrumb 'Edit Product', edit_admin_product_path
      end
    
      # POST /admin/products
      # POST /admin/products.json
      def create
        @admin_product = Admin::Product.new(admin_product_params)
    
        respond_to do |format|
          if @admin_product.save
            format.html { redirect_to @admin_product, notice: 'Product was successfully created.' }
            format.json { render :show, status: :created, location: @admin_product }
          else
            format.html { render :new }
            format.json { render json: @admin_product.errors, status: :unprocessable_entity }
          end
        end
      end
    
      # PATCH/PUT /admin/products/1
      # PATCH/PUT /admin/products/1.json
      def update
        respond_to do |format|
          if @admin_product.update(admin_product_params)
            format.html { redirect_to @admin_product, notice: 'Product was successfully updated.' }
            format.json { render :show, status: :ok, location: @admin_product }
          else
            format.html { render :edit }
            format.json { render json: @admin_product.errors, status: :unprocessable_entity }
          end
        end
      end
    
      # DELETE /admin/products/1
      # DELETE /admin/products/1.json
      def destroy
        @admin_product.destroy
        respond_to do |format|
          format.html { redirect_to admin_products_url, notice: 'Product was successfully destroyed.' }
          format.json { head :no_content }
        end
      end
    
      private
        # Use callbacks to share common setup or constraints between actions.
        def set_admin_product
          @admin_product = Admin::Product.find(params[:id])
        end
    
    
        # Only allow a list of trusted parameters through.
        def admin_product_params
          params.require(:admin_product).permit(:name, :description, :available_on, :deleted_at, :permalink, :meta_description, :meta_keywords, :count_on_hand, :admin_category_id, images: [])
        end
    end
    

In the admin/products/show view:

    &lt;p class="mb-6"&gt;&lt;%= @admin_product.admin_category.name %&gt;&lt;/p&gt;

I'm getting an undefined method \`admin\_category' error on page.
## [3][Large JSON Column Value in PostgreSQL](https://www.reddit.com/r/rails/comments/h02iuu/large_json_column_value_in_postgresql/)
- url: https://www.reddit.com/r/rails/comments/h02iuu/large_json_column_value_in_postgresql/
---
I have a PostgreSQL database that has a table where I store output of reports in. It's a fairly simple table because it has only 5 columns one of which is a `data` column that is of type `json`. I don't need to search through it but the idea was to store complex json data to represent a report and have it rendered in my views.

However I'm getting a lot of timeouts in Heroku if I view some large reports. My hunch is that the JSON data is too big to transmit over the wire.

I guess my question is, for complex data structures (I don't have to search it), is this approach good practice?
## [4][Where to place image files for production?](https://www.reddit.com/r/rails/comments/h08362/where_to_place_image_files_for_production/)
- url: https://www.reddit.com/r/rails/comments/h08362/where_to_place_image_files_for_production/
---
I’m reading the [Rails guides](https://guides.rubyonrails.org/asset_pipeline.html) and they’re saying that “Assets can still be placed in the public hierarchy. Any assets under public will be served as static files by the application or web server when config.public_file_server.enabled is set to true. You should use app/assets for files that must undergo some pre-processing before they are served.”

If I have images for my static pages, should I place them in assets/images or in public? Because I encountered an issue while deploying (on a VPS with Phusion Passenger). It was saying:

ActionView::Template::Error (The asset "picture.jpeg" is not present in the asset pipeline.):

I changed the production.rb to:

config.assets.compile = true

because of that.

But maybe I encountered this issue because I should put my images in the public folder? If so, what kind of images should go in assets?

Or am I encountering this issue because I’m not precompiling the assets in development?
## [5][Is there a library for getting names like Heroku's servers?](https://www.reddit.com/r/rails/comments/gzci9n/is_there_a_library_for_getting_names_like_herokus/)
- url: https://www.reddit.com/r/rails/comments/gzci9n/is_there_a_library_for_getting_names_like_herokus/
---
You know how Heroku has server or app names like "dancing-catfish-12345" or something? Is there a library out there I can use to do the same thing or so I need to come up with that on my own?
## [6][A novice vs. Rails IDs](https://www.reddit.com/r/rails/comments/gzfu0v/a_novice_vs_rails_ids/)
- url: https://www.reddit.com/r/rails/comments/gzfu0v/a_novice_vs_rails_ids/
---
TL;DR it seems like the convention for :id is bad for some common use cases, but it also seems like it's a bad idea to go against the convention. What gives?

I'm a complete novice when it comes to Rails, but I've been writing software professionally for 6 years, so not a novice overall.

I used Rails for a college class, and figured why not break it out for a pet project?

I needed to have functionality similar to Google Drive link sharing: secure-ish, non-authenticated resource access.

I figured cool, let's use this SecureRandom module to generate a string I can use as the resource's primary key. I could do link sharing with mywebsite.com/resource/4kUgL2pdQMSCQtjE for example.

After some Googling, I came to the conclusion that this is pretty much impossible / not-recommended / more trouble than it's worth and I should leave the primary key as an auto-increment int.

Okay... so I guess I could use the this string as a property on the model, and then just... also declare it as an index in the migration so lookup is fast?

Gave that a go. And it wasn't great.
* Had to change the routes to use resources :foobar, param: :uuid. What did this actually do? Seems like it just changed the name.
* Set the before_create on the model to generate the SecureRandom. So far so good.
* Overrode the to_param to return the uuid. Okay...
* Now the Foobar.find calls in the controller don't work. Makes sense. Update those to find_by.
* Okay... now the router is doing /1-uuid or /2-uuid instead of just the uuid.

Seems like a huge battle to do something really simple. Is it just expected that the default behavior is leaky routing and you can't easily change it? Am I missing something?
## [7][Help with Hiding and Showing Elements with Stimulus](https://www.reddit.com/r/rails/comments/gz7i3t/help_with_hiding_and_showing_elements_with/)
- url: https://www.reddit.com/r/rails/comments/gz7i3t/help_with_hiding_and_showing_elements_with/
---
I'm working on an app that lets users create connections with other users, like a Facebook or LinkedIn. When a user clicks a 'connect' button, I want to create the Active Record (Connection), and I want to change the button (to allow the person to 'disconnect' to destroy the Connection) without a page refresh.

I was looking to change the look of the button by changing the visibility attribute to hidden (for the 'connect' button), while setting the visibility of another button ('disconnect') to visible.

I was trying to do this using Stimulus, but it's not working. The record is creating, but page refreshes.  I didn't get to the part where I unhide a button yet.

It's probably a super simple thing for someone who's worked with Stimulus, but it's new to me and I'm no good at JS to boot.

Thanks!

    # the view 
    &lt;div data-controller="connection_button"&gt; 
      &lt;div style="visibility:visible;" data-target="connection_button.hide"&gt; 
        &lt;%= link_to connections_path(referrer_id: user.id), method: :post, remote: true do %&gt;                                     &lt;button class="btn btn-primary" data-action="connection_button#hideConnectButton"&gt; 
            Connect 
          &lt;/button&gt; 
        &lt;% end %&gt; 
      &lt;/div&gt; 
    &lt;/div&gt; 
    
    
    
    # javascript/controllers/connection_button_controller.js 
    import { Controller } from "stimulus" 
    
    export default class extends Controller 
      { static targets = [ "hide"]    
      
      hideConnectButton() 
       { this.showTarget.style.visibility = "hidden"; } 
      } 
    
    
    
    # the controller 
    def create   
      @connection = Connections.create!(user_id: current_user.id) 
      
      if @connection.save     
        redirect_back(fallback_location: root_path) 
      end 
    end

I tried to get rid of the linked\_to to see if just the Stimulus works, and it does not. So I definitely did something wrong (no JS errors in the browser).

`&lt;%= link_to connections_path(referrer_id: user.id), method: :post, remote: true do %&gt;`

&amp;#x200B;

I also tried removing the visibility attribute on the div, so it was:

`&lt;div data-target="connection_button.hide"&gt;`

&amp;#x200B;

Based on [this article](https://boringrails.com/articles/better-stimulus-controllers/), I also tried changing the data target like this:

    this.contentTarget.classList.connection_button("hidden");
## [8][Does anyone know if Rails 4.x will run on modern Ruby?](https://www.reddit.com/r/rails/comments/gz22dw/does_anyone_know_if_rails_4x_will_run_on_modern/)
- url: https://www.reddit.com/r/rails/comments/gz22dw/does_anyone_know_if_rails_4x_will_run_on_modern/
---
Weird question, I know.

I've inherited an application that runs on Ruby 2.0.0 using Rails 4.x.  It's all old enough that I had to install Ubuntu 14.x just to get a working dev environment.  It's literally running Ubuntu 14.x in production.

Obviously this is a problem and it needs to be updated.  What I'd like to do is update the Ruby first if possible.  That would allow me to install a modern OS that hasn't been End of Life'd and then work on updating the framework from there.

Has anyone else had to deal with this before?  I haven't started digging into the problem just yet, but I'd love to hear people's opinions on the best way to go about this.

In theory I would first update to modern Ruby and then update the framework, but my fear is that I'll end up needing to do both at the same time which will probably effectively turn into a rewrite.
## [9][Weird performance issue with rails serializer](https://www.reddit.com/r/rails/comments/gza84p/weird_performance_issue_with_rails_serializer/)
- url: https://www.reddit.com/r/rails/comments/gza84p/weird_performance_issue_with_rails_serializer/
---
I'm one of my current projects I have an table with around 30 thousand rows, I've realized however that the index of said table takes some time to render, upwards of five seconds. While debugging and stripping the controller of every logic possible I was left with an request time of around 3\~3.5 seconds, which was still weird.  


At this time, all I had on my controller was:  


    class RecordController &lt; ApplicationController
      def index
        records = Records.all.page(params[:page] || 1).per(params[:per] ||  10)
        render json: records
      end
    end

Using kaminari pagination, I was rendering only 10 instances of Record. My serializer looked like this:

    class RecordSerializer &lt; ActiveModel::Serializer
      type 'engine-records'
      attributes :id 
    end

Then I realized that no matter how many Records I got through pagination, from 1 to 100, the response time was still the same. I only managed to get it lower by removing the serializer file.

By removing my RecordSerializer I've managed to get an response time of 150ms with 100 records per page. To me this seem absurd, almost 3 seconds of performance overhead caused by such an basic serializer. I can't find any posts online on the subject, and this seems too absurd to be "expected behaviour", anyone got any idea what might be causing this? Is there some sort of configuration that I might have missed or anything like that? Dropping the serializer is not an actual potion since I have to render Record and it's relationships on that endpoint, so I need the serializer, but I'd like to avoid those extra three seconds if possible
## [10][Custom Breakpoints for Mobile - Bootstrap for Rails](https://www.reddit.com/r/rails/comments/gzn5qh/custom_breakpoints_for_mobile_bootstrap_for_rails/)
- url: https://www.reddit.com/r/rails/comments/gzn5qh/custom_breakpoints_for_mobile_bootstrap_for_rails/
---
Hi Folks,

I have a rails app and I'd like to change the columns from col-6 for sm/md/lg viewing screens to col-12 for mobile. Would this require making custom mixins within a custom.scss file? Thanks in advace!
## [11][Javascript works in .html file within application, but doesn't work in .html.erb file](https://www.reddit.com/r/rails/comments/gz5wxi/javascript_works_in_html_file_within_application/)
- url: https://www.reddit.com/r/rails/comments/gz5wxi/javascript_works_in_html_file_within_application/
---
I'm using Lugo Labs Circle  [https://www.lugolabs.com/circles](https://www.lugolabs.com/circles) . I've created a test.html file in the root of my application that references circles.min.js and the javascript code works fine. I put the same code in my show.html.erb file referencing the same circles.min.js file and it doesn't work. Any ideas why?
