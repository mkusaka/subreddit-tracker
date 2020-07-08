# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/hja8kx/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/hja8kx/personal_projects_show_off_your_own_project_andor/
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
## [2][How to make friendly_id backfilling migration faster? You can skip all the callbacks.](https://www.reddit.com/r/rails/comments/hndvj5/how_to_make_friendly_id_backfilling_migration/)
- url: https://www.reddit.com/r/rails/comments/hndvj5/how_to_make_friendly_id_backfilling_migration/
---
I am currently working on integrating friendly_id gem into some of the models in Talenox. Basically, it makes our in app URLs look nicer with human and company names in front, instead of just incremental primary key IDs. Oh boy… `Employee.all.each(&amp;:save)` is fucking slow in production.

There are several things that can cause update and insert to slow down a lot for an ActiveRecord model:

* Validations - especially when it involves multiple models
* Callbacks - especially when they cause a chain of callbacks in other models
* `belongs_to :parent, touch: true` - technically a callback to bust russian doll caches, but adding a slug does not necessitate busting caches

Guess what, we can skip all those. How? By backfilling with an empty model class.

Assuming we have an Employee model with a relation employees, what you can do is: Create an ActiveRecord model class in that migration class with none of the callbacks EXCEPT friendly_id and slug_candidate method.

    class BackfillEmployeesWithFriendlyId &lt; ActiveRecord::Migration[5.0]
    
      # Using a blank class allows us to easily skip all callbacks that can make
      # mass migration slow.
      class FriendlyIdEmployee &lt; ActiveRecord::Base
        self.table_name = 'employees'
        extend FriendlyId
        friendly_id :slug_candidate, use: [:slugged, :finders]
    
        def slug_candidate
          if first_name || last_name
            "#{first_name} #{last_name}"[0, 20]
          else
            "employee"
          end + " #{SecureRandom.hex[0, 8]}"
        end
      end
    
      def up
        print "Updating friendly_id slug for employees"
        FriendlyIdEmployee.where(slug: nil).each do |row|
          row.save; print('.')
        end
        puts ''
      end
    end
    
However, I couldn’t get the friendly_id history plug in to work properly yet. friendly_id history is implemented using ActiveRecord polymorphic. When the backfilling migration above is run, it will end up creating FriendlyId::Slug records with sluggable type of `BackfillEmployeesWithFriendlyId::FriendlyIdEmployee` instead of just `Employee`. That also means you can’t do subclassing of ActiveRecord models with friendly_id and expect history to work. Luckily we don’t need it.

[Source](https://anonoz.github.io/tech/2020/07/08/faster-friendly_id-mass-migration.html)
## [3][Best architecture for preferred sort for e-commerce products?](https://www.reddit.com/r/rails/comments/hn78ue/best_architecture_for_preferred_sort_for/)
- url: https://www.reddit.com/r/rails/comments/hn78ue/best_architecture_for_preferred_sort_for/
---
I have an ecommerce app on rails and was wondering what's the best architecture for ordering products on a page. I'm aware of how to use .order(:table_attribute), but was wondering if there is a cleaner/more dynamic way to bubble offerings to the top that you want to highlight. Any recommended practices for this?
## [4][Database administration tool like ActiveAdmin/Rails_admin](https://www.reddit.com/r/rails/comments/hnchui/database_administration_tool_like/)
- url: https://www.reddit.com/r/rails/comments/hnchui/database_administration_tool_like/
---
I am looking for ways to work on an external (i.e. not the rails-backend) database from within a rails app, preferably in a very user-friendly and powerful way like ActiveAdmin/Rails\_admin allow. 

&amp;#x200B;

Does anyone know of a ready-made solution? To my knowledge you can't just point the two former tools to a different DB.
## [5][If you were creating a new Rails API, what serializer would you use and why?](https://www.reddit.com/r/rails/comments/hmymab/if_you_were_creating_a_new_rails_api_what/)
- url: https://www.reddit.com/r/rails/comments/hmymab/if_you_were_creating_a_new_rails_api_what/
---
We have a relatively new Rails API project (full Rails, not API only) that is using ActiveModelSerializer. We're considering switching due to some frustrations with the project, it's documentation, the current state of it, and a desire to get away from the confusing DSL. If you were starting a new project is there a serializer you recommend? fast-json-api comes to mind, but I haven't fully explored all the alternatives.
## [6][Why does Stripe only have documentation for Sinatra?](https://www.reddit.com/r/rails/comments/hn72ht/why_does_stripe_only_have_documentation_for/)
- url: https://www.reddit.com/r/rails/comments/hn72ht/why_does_stripe_only_have_documentation_for/
---
The ruby documentation for Stripe is specific to Sinatra and not rails. Heroku for the ruby documentation also defaults to Sinatra, and has Rails in a separate section. With rails being the most popular ruby framework, is there some a reason these products don't include rails as the default? Is there some historical reason for this? Just curious.
## [7][r/rails salary poll](https://www.reddit.com/r/rails/comments/hn6p3s/rrails_salary_poll/)
- url: https://www.reddit.com/r/rails/comments/hn6p3s/rrails_salary_poll/
---
Here's a salary poll for the group, if anybody's interested in playing.

If so, how about some guidelines?

- amounts are in USD (or equivalent)
- intended for those working as FTE doing Rails/Ruby work either mostly or exclusively

Can contractors play too? I think so, but too bad there's no radio button option.

Should amounts include total compensation?

[View Poll](https://www.reddit.com/poll/hn6p3s)
## [8][Creating a rest API with devise-jwt, advice and guide needed](https://www.reddit.com/r/rails/comments/hmwm3s/creating_a_rest_api_with_devisejwt_advice_and/)
- url: https://www.reddit.com/r/rails/comments/hmwm3s/creating_a_rest_api_with_devisejwt_advice_and/
---
Greetings guys. 

I was searching about JWT authentication in rails, and first I found this guy explained how you can make an API with JWT authentication system : 

[https://medium.com/@nandhae/2019-how-i-set-up-authentication-with-jwt-in-just-a-few-lines-of-code-with-rails-5-api-devise-9db7d3cee2c0](https://medium.com/@nandhae/2019-how-i-set-up-authentication-with-jwt-in-just-a-few-lines-of-code-with-rails-5-api-devise-9db7d3cee2c0) 

he uses devise, which is really good. but I couldn't find how I can do signup and sign-in and get tokens, etc. Should I write methods? (his git project didn't have methods written). 

And I also find this one: 

[https://medium.com/better-programming/build-a-rails-api-with-jwt-61fb8a52d833](https://medium.com/better-programming/build-a-rails-api-with-jwt-61fb8a52d833)

and this guy, created user model himself. I think it's less secure and became my second priority to implement my project like this. 

So, I ask you guys, do you have any advice/guide for me?
## [9][Unpermitted Parameter](https://www.reddit.com/r/rails/comments/hmuzt3/unpermitted_parameter/)
- url: https://www.reddit.com/r/rails/comments/hmuzt3/unpermitted_parameter/
---
I am trying to update a product model and also create any new option\_types/option\_values at that time. I am getting an error that option\_type is an unpermitted parameter. I'm all tripped up on this. Any ideas what I have wrong?

    09:51:10 web.1     |   ↳ app/controllers/admin/products_controller.rb:88:in `set_product'
    09:51:10 web.1     | Unpermitted parameter: :option_type
    09:51:10 web.1     | Redirected to http://localhost:5000/admin/products/30

products form section:

&amp;#x200B;

    &lt;%= form_with(model: [:admin, product], local: true) do |form| %&gt;        
    
    &lt;div class="pt-8" &gt;
              &lt;div class="grid md:grid-cols-1 row-gap-6 col-gap-4 lg:grid-cols-3"&gt;
    
               &lt;% u/option_types.each do |option_type| %&gt;
                 &lt;%= content_tag :tr, id: dom_id(option_type), class: dom_class(option_type) do %&gt;
                   &lt;td&gt;&lt;%= link_to option_type.name, edit_admin_product_option_type_path(@product, option_type.id) %&gt;&lt;/td&gt;
                 &lt;% end %&gt;
               &lt;% end %&gt;
    
               &lt;%= form.fields_for :option_type, OptionType.new do |options|%&gt;
    
                 &lt;div&gt;
                   &lt;%= options.label "Option Type Name", class: "text-gray-700" %&gt;
                   &lt;%= options.text_field :name, class: 'w-full mt-2 px-4 py-2 block rounded bg-gray-200 text-gray-800 border border-gray-300 focus:outline-none focus:bg-white' %&gt;
                 &lt;/div&gt;
    
                 &lt;div data-controller="nested-form"&gt;
                   &lt;%= options.fields_for :option_values, OptionValue.new, child_index: 'NEW_RECORD' do |ov| %&gt;
                     &lt;%= render "admin/option_types/option_values_fields", form: ov %&gt;
                   &lt;% end %&gt;
                 &lt;/div&gt;
    
                 &lt;div class="form-group"&gt;
                   &lt;%= options.hidden_field 'product_id', :value =&gt; params[:product_id] %&gt;
                   &lt;%= options.submit class: 'inline-block px-5 py-3 shadow-lg rounded-lg bg-indigo-600 text-white hover:bg-indigo-700 cursor-pointer' %&gt;
                 &lt;/div&gt;
    
               &lt;% end %&gt;
             &lt;/div&gt;
          &lt;/div&gt;
    &lt;% end %&gt;

&amp;#x200B;

product Controller:

&amp;#x200B;

    class Admin::ProductsController &lt; Admin::ApplicationController
      before_action :set_product, only: [:show, :edit, :update, :destroy]
    
      # GET /products
      # GET /products.json
      def index
    
        @products = Product.all
        add_breadcrumb('Products')
        @pagetitle = "View Products"
      end
    
      # GET /products/1
      # GET /products/1.json
      def show
      end
    
      # GET /products/new
      def new
        @option_types = @product.option_types
        @product = Product.new
        add_breadcrumb('Products', admin_products_path)
        add_breadcrumb(@product.name)
        @pagetitle = "Create Product"
      end
    
      # GET /products/1/edit
      def edit
        @option_types = @product.option_types
        add_breadcrumb('Products', admin_products_path)
        add_breadcrumb(@product.name)
        @pagetitle = "Edit Product"
      end
    
      # POST /products
      # POST /products.json
      def create
        @product = Product.new(product_params)
    
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
    
            format.html { redirect_to admin_product_url, notice: 'Product was successfully updated.' }
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
          params.require(:product).permit(:name, :description, :available_on, :slug, :meta_description, :meta_keywords, :price, :main_image, :option_type_id, :option_type, images: [])
        end
    end

product model:

    class Product &lt; ApplicationRecord
      has_many_attached :images, :dependent =&gt; :delete_all
      has_one_attached :main_image, :dependent =&gt; :delete_all
      has_many :product_option_types, dependent: :destroy, inverse_of: :product
      has_many :option_types, through: :product_option_types, dependent: :destroy
      accepts_nested_attributes_for :option_types
      has_many :variants, inverse_of: :product, dependent: :destroy
    end

option\_type model:

    class OptionType &lt; ApplicationRecord
      belongs_to :product
      has_many :option_values, dependent: :destroy
      accepts_nested_attributes_for :option_values, reject_if: :all_blank, allow_destroy: true
      has_many :product_option_types, dependent: :destroy
    end
## [10][How to create a form to embed into other websites?](https://www.reddit.com/r/rails/comments/hmrp3q/how_to_create_a_form_to_embed_into_other_websites/)
- url: https://www.reddit.com/r/rails/comments/hmrp3q/how_to_create_a_form_to_embed_into_other_websites/
---
I want to create a form that lives in my app but I can then embed it into any other website that I give a piece of code to place into their website.

&amp;#x200B;

What is the best way to go about this with Rails?

&amp;#x200B;

I assume I should be using an iframe but wondering any obstacles I may need to watch out for and any opinions on how I should go about doing this?
## [11][Up to date resource for starting to learn rails](https://www.reddit.com/r/rails/comments/hmod8i/up_to_date_resource_for_starting_to_learn_rails/)
- url: https://www.reddit.com/r/rails/comments/hmod8i/up_to_date_resource_for_starting_to_learn_rails/
---
Hey guys. I'm looking to get into rails development and I'm wondering if there are any recommended up to date resources. The JHU Coursera course seemed tempting but it is from 2015 and uses rails 4 which I feel like is a bad place to start. There's also the learn enough [railstutorial.org](https://railstutorial.org) which seems tempting, but I'd like to look around for a free resource first. 

Any recommendations and advice are appreciated.

Cheers.
