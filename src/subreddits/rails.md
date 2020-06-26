# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/har6r7/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/har6r7/personal_projects_show_off_your_own_project_andor/
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
## [2][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/hfkxk4/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/hfkxk4/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [3][Issue with nested form](https://www.reddit.com/r/rails/comments/hg6fc8/issue_with_nested_form/)
- url: https://www.reddit.com/r/rails/comments/hg6fc8/issue_with_nested_form/
---
I am having a problem with a nested form setup. I have a products model that has product\_options from that model. What I am trying to do is have a button to add a new product option name/value within the product form. I also want to list within the product the current options and be able to delete them if needed. I do not see any form fields displaying though for the product\_options fields. What am I doing worng?

Product model

    class Product &lt; ApplicationRecord
      belongs_to :category
      has_many :product_options
      accepts_nested_attributes_for :product_options, allow_destroy: true
      has_many_attached :images
      has_one_attached :main_image
    end
    

Product\_options model

    class ProductOption &lt; ApplicationRecord
      belongs_to :product, optional: true
    end
    

product form (product\_options at bottom)

    &lt;%= form_with(model: [:admin, product], local: true) do |form| %&gt;
    &lt;% if product.errors.any? %&gt;
    &lt;div id="error_explanation"&gt;
       &lt;h2&gt;&lt;%= pluralize(product.errors.count, "error") %&gt; prohibited this product from being saved:&lt;/h2&gt;
       &lt;ul&gt;
          &lt;% product.errors.full_messages.each do |message| %&gt;
          &lt;li&gt;&lt;%= message %&gt;&lt;/li&gt;
          &lt;% end %&gt;
       &lt;/ul&gt;
    &lt;/div&gt;
    &lt;% end %&gt;
    &lt;!-- &lt;h3 class="mb-4 text-2xl"&gt;Add New admin_product&lt;/h3&gt; --&gt;
    &lt;%= form.hidden_field :category_id %&gt;
    &lt;div class="mb-6"&gt;
       &lt;%= form.label :name, class: "label" %&gt;
       &lt;%= form.text_field :name, class: "input" %&gt;
    &lt;/div&gt;
    &lt;div class="mb-6"&gt;
       &lt;%= form.label "Product Category", class: "label" %&gt;
       &lt;%= form.collection_select(:category_id, Category.order(:category), :id, :category, {include_blank: true}, { class: "fmt-1 block form-select w-full py-2 px-3 py-0 border border-gray-300 bg-gray-100 rounded-md shadow-sm focus:outline-none focus:shadow-outline-blue focus:border-blue-300 transition duration-150 ease-in-out sm:text-sm sm:leading-5" }) %&gt;
    &lt;/div&gt;
    &lt;div class="mb-6"&gt;
       &lt;%= form.label :description, class: "label" %&gt;
       &lt;%= form.text_area :description, class: "input" %&gt;
    &lt;/div&gt;
    &lt;div class="mb-6"&gt;
       &lt;%= form.label :available_on, class: "label" %&gt;
       &lt;div class="flex items-center justify-between max-w-md"&gt;
          &lt;%= form.text_field :available_on, data: {behavior: "flatpickr"} %&gt;
       &lt;/div&gt;
    &lt;/div&gt;
    &lt;div class="mb-6"&gt;
       &lt;%= form.label "Number in Stock", class: "label" %&gt;
       &lt;%= form.number_field :stock, class: "input"  %&gt;
    &lt;/div&gt;
    &lt;div class="mb-6"&gt;
       &lt;%= form.label "Upload Images", class: "label" %&gt;
       &lt;%= form.file_field :images, multiple: true, :class =&gt; "input" %&gt;
    &lt;/div&gt;
    &lt;hr class="border mb-4" /&gt;
    &lt;div class="border-2 border-dashed rounded-md mt-3 pt-4"&gt;
       Add a product option: &lt;br /&gt;&lt;br /&gt;
       &lt;%= form.fields_for :product_options do |builder| %&gt;
       &lt;%= builder.hidden_field :product_id, value: product.id %&gt;
       &lt;div class="mb-6"&gt;
          &lt;%= builder.label :name, class: "label" %&gt;
          &lt;%= builder.text_field :name, class: "input" %&gt;
       &lt;/div&gt;
       &lt;div class="mb-6"&gt;
          &lt;%= builder.label :value, class: "label" %&gt;
          &lt;%= builder.text_area :value, class: "input" %&gt;
       &lt;/div&gt;
       &lt;% end %&gt;
    &lt;/div&gt;
    &lt;%= form.submit class: "inline-block px-5 py-3 shadow-lg rounded-lg bg-indigo-600 text-white hover:bg-indigo-700 cursor-pointer" %&gt;
    &lt;%= link_to 'Cancel', admin_products_path, { :class=&gt;"inline-block px-5 py-3 shadow-lg rounded-lg bg-gray-600 text-white hover:bg-gray-700 cursor-pointer" } %&gt;
    &lt;% end %&gt;

Products controller

    class Admin::ProductsController &lt; Admin::ApplicationController
    before_action :set_product, only: [:show, :edit, :update, :destroy]
    
    def index
      @products = Product.all
      add_breadcrumb('Products')
      @pagetitle = "View Products"
    end
    
    # GET /products/1
    # GET /products/1.json
    def show
      add_breadcrumb('Products', admin_products_path)
      add_breadcrumb(@product.name)
      @pagetitle = "View Product"
    end
    
    # GET /products/new
    def new
      @product = Product.new
      @product.product_options.build
      add_breadcrumb('Products', admin_products_path)
      add_breadcrumb(@product.name)
      @pagetitle = "Create New Product"
    end
    
    # GET /products/1/edit
    def edit
      add_breadcrumb('Products', admin_products_path)
      add_breadcrumb(@product.name)
      @pagetitle = "Edit Product"
    end
    
    # POST /products
    # POST /products.json
    def create
      @product = Product.new(product_params)
      @product_options = @product.product_options.create(product_params)
      respond_to do |format|
        if @product.save
          format.html { redirect_to admin_products_url, notice: 'Product was successfully created.' }
          format.json { render :show, status: :created, location: @product }
        else
          format.html { render :new }
          format.json { render json: @product.errors, status: :unprocessable_entity }
        end
      end
    end
    
    # PATCH/PUT /products/1
    # PATCH/PUT /products/1.json
    def update
      respond_to do |format|
        if @product.update(product_params)
          format.html { redirect_to @product, notice: 'Product was successfully updated.' }
          format.json { render :show, status: :ok, location: @product }
        else
          format.html { render :edit }
          format.json { render json: @product.errors, status: :unprocessable_entity }
        end
      end
    end
    
    # DELETE /products/1
    # DELETE /products/1.json
    def destroy
      @product.destroy
      respond_to do |format|
        format.html { redirect_to admin_products_url, notice: 'Product was successfully destroyed.' }
        format.json { head :no_content }
      end
    end
    
    def delete_image_attachment
      @image = ActiveStorage::Attachment.find(params[:id])
      @image.purge
      redirect_back(fallback_location: request.referer)
    end
    
    private
      # Use callbacks to share common setup or constraints between actions.
      def set_product
        @product = Product.find(params[:id])
      end
    
      # Only allow a list of trusted parameters through.
      def product_params
        params.require(:product).permit(:name, :description, :stock, :available_on, :permalink, :category_id, :main_image, :properties, images: [], product_options_attributes: [:name, :value])
      end
    
    end
    

Any help is appreciated. Thanks!
## [4][Show a list of each database entry on html page](https://www.reddit.com/r/rails/comments/hfo5hs/show_a_list_of_each_database_entry_on_html_page/)
- url: https://www.reddit.com/r/rails/comments/hfo5hs/show_a_list_of_each_database_entry_on_html_page/
---
Currently, I use &lt;%= @ places.each do |place| %&gt; within a table to display all records of places that have been created. they order out one after the other, and I am searching for an easy way to list them out one beneath the other on the page. Any ideas? And thank you in advance.
## [5][Attaching a source to customer for payment method with rails &amp; stripe?](https://www.reddit.com/r/rails/comments/hfm4hf/attaching_a_source_to_customer_for_payment_method/)
- url: https://www.reddit.com/r/rails/comments/hfm4hf/attaching_a_source_to_customer_for_payment_method/
---
I have a rails app and in test mode I have no problem accepting payments, but in live mode I get an error like "Stripe::InvalidRequestError (A source must be attached to a customer to be used as a `payment_method`.)". I've attached a token to the customer when it is created (see code below), but not sure where else it needs to be.

    customer = if current_user.stripe_id[connected_acct].present?
        Stripe::Customer.retrieve(current_user.stripe_id[connected_acct], {stripe_account: item.stripe_id})
    else
        Stripe::Customer.create({
                    email: current_user.email, 
                    source: token,
                },
                {
                    stripe_account: item.stripe_id,
                })

Confirming the payment intent looks like this:

    payment_method_card = params[:user][:card_id]
     confirm_payment = Stripe::PaymentIntent.confirm(
                payment_intent.id,
                {payment_method: payment_method_card},
            )
## [6][Confirm a payment intent in Live Mode (Rails &amp; Stripe)](https://www.reddit.com/r/rails/comments/hfd8uv/confirm_a_payment_intent_in_live_mode_rails_stripe/)
- url: https://www.reddit.com/r/rails/comments/hfd8uv/confirm_a_payment_intent_in_live_mode_rails_stripe/
---
I have a rails application that uses stripe to handle payments &amp; fees. In test mode my code works well, but when I attempt to move it over to live mode I run into issues. I'm now using production/live API keys, the issue is around confirming a basic intent. In test mode the confirmation looks something like so (where card_brand = 'visa' or 'mastercard' etc)

    payment_method_card = 'pm_card_' + card_brand
    confirm_payment = Stripe::PaymentIntent.confirm(
                payment_intent.id,
                {payment_method: payment_method_card},
    )

This doesn't fly in production and I'm just wondering what parameter needs to be passed to the payment_method attribute in the confirmation. Any/all help is appreciated.
## [7][Advice for background multiple long-running tasks](https://www.reddit.com/r/rails/comments/hfag0c/advice_for_background_multiple_longrunning_tasks/)
- url: https://www.reddit.com/r/rails/comments/hfag0c/advice_for_background_multiple_longrunning_tasks/
---
Hi guys,

I currently have numerous sidekiq workers that interact with ActiveRecord to store/retrieve data that is being pushed to via HTTP POST requests to an API Controller.

When I run a health check on my customers' networks, an agent on their system interacts with my API, which  looks in ActiveRecord for records that my sidekiq workers have created, retrieves the commands to run and then sends it back to the agents via HTTP response. As the agent completes each task, it submits an HTTP POST request back to the API, which then stores the data in ActiveRecord. The sidekiq worker obtains this data and updates other attributes within the application.

Since it is not a best practice to store lots of code in sidekiq workers, or even have them run for a long time, what are some other alternatives?

Ultimately, I just need to push a series of commands and get their responses, regardless if it takes 5 seconds or 15 minutes.

Here's an example of what several workers are doing:

`Task.create(command: "ipconfig")`  
`# waits for task.status == "completed" by running continuous task.status calls every second`  
`if task.output.include? "this"`  
   `do this`  
`else`   
   `do this instead`  
`end`  
`Task.create(command: "next command")`  
`# waits for task.status == "completed" by running continuous task.status calls every second`
## [8][Should I load the pipeline JS and CSS using a CDN like Cloudinary?](https://www.reddit.com/r/rails/comments/hf1r1h/should_i_load_the_pipeline_js_and_css_using_a_cdn/)
- url: https://www.reddit.com/r/rails/comments/hf1r1h/should_i_load_the_pipeline_js_and_css_using_a_cdn/
---
My webapp's pipeline asset size is making the first page load quite delayed. Hence, I am thinking of uploading the precompiled and minified CSS and JS files to Cloudinary. Is it a good practice?
## [9][Is Ruby on Rails front and back end development](https://www.reddit.com/r/rails/comments/hf2w7s/is_ruby_on_rails_front_and_back_end_development/)
- url: https://www.reddit.com/r/rails/comments/hf2w7s/is_ruby_on_rails_front_and_back_end_development/
---

## [10][How should I write my frontend code with rails?](https://www.reddit.com/r/rails/comments/hf6ta8/how_should_i_write_my_frontend_code_with_rails/)
- url: https://www.reddit.com/r/rails/comments/hf6ta8/how_should_i_write_my_frontend_code_with_rails/
---
Using "vanilla" tecnologies such as css, js and so goes on or creating/customizing themes from bulma/bootstrap? I'm confused.
## [11][Solidus won't display images out of the box - Paperclip and ImageMagick error](https://www.reddit.com/r/rails/comments/hf0te0/solidus_wont_display_images_out_of_the_box/)
- url: https://www.reddit.com/r/rails/comments/hf0te0/solidus_wont_display_images_out_of_the_box/
---
I'm just starting with Solidus, though I do know my way around Rails. I've created a Solidus project on my Mac, and I installed all the required gems as well as ImageMagick through Homebrew.  Both `identify` and `magick` work as terminal commands as my `/usr/local/bin` directory is in my PATH. I've also specified the path to `identify` in the `application.rb` file (I'm guessing it's supposed to go in side the `Application` class). However, none of the images show up in Solidus, and when I try and manually add an image, I get the following error (it's related to paperclip). 

Any help here would be - Google has not been my friend this time. BTW, the same thing happens on both Mac and Windows running WSL (Debian). The error message is shown below:

    [paperclip] Trying to link /tmp/RackMultipart20200622-9041-8gndhx.jpg to /tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-1t35lnu.jpg
    [paperclip] Trying to link /tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-1t35lnu.jpg to /tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-62aumc.jpg
    Command :: PATH=/usr/local/bin/identify:$PATH; file -b --mime '/tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-62aumc.jpg'
    Command :: PATH=/usr/local/bin/identify:$PATH; identify -format '%wx%h,%[exif:orientation]' '/tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-1t35lnu.jpg[0]' 2&gt;/dev/null
    [paperclip] An error was received while processing: #&lt;Paperclip::Errors::NotIdentifiedByImageMagickError: Paperclip::Errors::NotIdentifiedByImageMagickError&gt;
    Command :: PATH=/usr/local/bin/identify:$PATH; identify -format '%wx%h,%[exif:orientation]' '/tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-1t35lnu.jpg[0]' 2&gt;/dev/null
    [paperclip] An error was received while processing: #&lt;Paperclip::Errors::NotIdentifiedByImageMagickError: Paperclip::Errors::NotIdentifiedByImageMagickError&gt;
    Command :: PATH=/usr/local/bin/identify:$PATH; identify -format '%wx%h,%[exif:orientation]' '/tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-1t35lnu.jpg[0]' 2&gt;/dev/null
    [paperclip] An error was received while processing: #&lt;Paperclip::Errors::NotIdentifiedByImageMagickError: Paperclip::Errors::NotIdentifiedByImageMagickError&gt;
    Command :: PATH=/usr/local/bin/identify:$PATH; identify -format '%wx%h,%[exif:orientation]' '/tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-1t35lnu.jpg[0]' 2&gt;/dev/null
    [paperclip] An error was received while processing: #&lt;Paperclip::Errors::NotIdentifiedByImageMagickError: Paperclip::Errors::NotIdentifiedByImageMagickError&gt;
    [paperclip] Trying to link /tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-1t35lnu.jpg to /tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-1dlkgdw.jpg
    Command :: PATH=/usr/local/bin/identify:$PATH; file -b --mime '/tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-1dlkgdw.jpg'
    [paperclip] Trying to link /tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-1t35lnu.jpg to /tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-16yjulo.jpg
    Command :: PATH=/usr/local/bin/identify:$PATH; file -b --mime '/tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-16yjulo.jpg'
      Rendering /Library/Ruby/Gems/2.6.0/gems/solidus_backend-2.10.1/app/views/spree/admin/images/create.js.erb
      Rendered /Library/Ruby/Gems/2.6.0/gems/solidus_backend-2.10.1/app/views/spree/admin/images/create.js.erb (Duration: 3.4ms | Allocations: 874)
    Completed 200 OK in 836ms (Views: 20.7ms | ActiveRecord: 7.6ms | Allocations: 29100)
## [12][undefined local variable or method `resource' for #&lt;PostController:0x00007fd90d028990&gt;](https://www.reddit.com/r/rails/comments/hf7kck/undefined_local_variable_or_method_resource_for/)
- url: https://www.reddit.com/r/rails/comments/hf7kck/undefined_local_variable_or_method_resource_for/
---
I can create posts from  [http://localhost3030/posts/new](http://localhost3030/posts/new) and they save to the database. but I get the above error on the \* line below. Also below is how I called user\_path(resource) within a after\_sign\_in method within my applications controller yesterday and it worked. I figured calling the same in my post controller create method would re\_direct to the same route, but it resolved the above error. Any ideas why? And thank you in advance. 

def create

  @ post = current\_user.posts.create(post\_params)

  if @ post.valid?

\*    redirect\_to user\_path(resource) 

  else

render :new, status: :unprocessable\_entity

  end

  end

&amp;#x200B;

 [https://www.reddit.com/r/rails/comments/heku7v/no\_route\_matches\_actionshow\_controllerusers/](https://www.reddit.com/r/rails/comments/heku7v/no_route_matches_actionshow_controllerusers/)
