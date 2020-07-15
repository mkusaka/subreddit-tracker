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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/hrnm2o/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/hrnm2o/personal_projects_show_off_your_own_project_andor/
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
## [3][OMG! I just got my first user on my rails app](https://www.reddit.com/r/rails/comments/hr6ms0/omg_i_just_got_my_first_user_on_my_rails_app/)
- url: https://www.reddit.com/r/rails/comments/hr6ms0/omg_i_just_got_my_first_user_on_my_rails_app/
---
Doing webdev for almost a 2 years, started with JS, drove me crazy. Just wanted to learn to build something without doing to much my self. Found out about rails, started building some simple stuff, like a blog. After that build a app for my brother, and wrote some SEO about it.... and finally got my first user today! even if it is just a free user (have a free restricted plan and a premium plan) it is a great feeling!
## [4][Rails 6 and Active Job : a simple tutorial](https://www.reddit.com/r/rails/comments/hrnisf/rails_6_and_active_job_a_simple_tutorial/)
- url: https://www.reddit.com/r/rails/comments/hrnisf/rails_6_and_active_job_a_simple_tutorial/
---
Active Job is quite difficult to start with. The documentation is very good, but to just to start an "hello world" job example, it requires redis, sidekiq (or another implementation), and some configuration. I tried to create a small tutorial in which 1) you have everything that just works in a couple of minutes and 2) everything works in complete isolation of your environment. [http://bdavidxyz.com/blog/rails-6-activejob-tutorial/](http://bdavidxyz.com/blog/rails-6-activejob-tutorial/)
## [5][Is action cable in rails 6 production ready?](https://www.reddit.com/r/rails/comments/hrhjuw/is_action_cable_in_rails_6_production_ready/)
- url: https://www.reddit.com/r/rails/comments/hrhjuw/is_action_cable_in_rails_6_production_ready/
---
Has any one used action cable for over 1000+ connections and had it work reliably with low latency? I'm just testing websockets in my app but it seems to be incredibly slow, is there something I could be doing wrong?
## [6][How do I link to a :delete action on a resource inside a module?](https://www.reddit.com/r/rails/comments/hrmhrv/how_do_i_link_to_a_delete_action_on_a_resource/)
- url: https://www.reddit.com/r/rails/comments/hrmhrv/how_do_i_link_to_a_delete_action_on_a_resource/
---
Hi All - I can't figure out the syntax for a `link_to` for the `#destroy` action on a controller in a module.

I have a `Post` resource with a `posts_controller` for normal user interaction with the most. I created another `posts_controller` in an `admin` module for Administrator actions.

    resources :posts 
    namespace :admin do
      resources :posts
    end

The `controllers/admin/posts_controller.rb` is defined as:

    module Admin
      class PostsController &lt; ApplicationController
        # normal resource methods
      end
    end

That all works exactly as expected, and I can create a link to the admin's view of all posts with:

    link_to 'Admin View of Posts', admin_posts_path

What I can't figure out is how to create a link to `admin/posts_controller.rb#destroy` in my `admin/posts/show.html.haml` view. If I do:

    link_to 'Delete', post, method: :delete, data: { confirm: '...' }

Then that generates a link to the normal user `posts_controller` not the admin one. How do I create the destroy action link on the admin posts controller?

Thanks!
## [7][Print PDF with Page Number](https://www.reddit.com/r/rails/comments/hrizi8/print_pdf_with_page_number/)
- url: https://www.reddit.com/r/rails/comments/hrizi8/print_pdf_with_page_number/
---
Any recommendation for PDF generation with Rails? Need one with custom header and page number. Simply printing a page from javascript print is nice but I ran into some issues with formatting and page numbers. `wkhtmltopdf` seems to be outdated and doesn't place nice as well. `prawn` is a little too low level. Any other solutions out there? Something like crystal reports would be nice.
## [8][Everything you want to know about writing your own form objects in rails](https://www.reddit.com/r/rails/comments/hr1aol/everything_you_want_to_know_about_writing_your/)
- url: https://www.reddit.com/r/rails/comments/hr1aol/everything_you_want_to_know_about_writing_your/
---
Hey r/rails, I've written up a 3-part series on form objects (3 years in the making). If you've wrestled with rails' form helpers and builders to trying to make some frontend design work or with models that don't seem to cooperate, then you've probably already tried implementing form objects, or using a form object gem.

Thing is, there are many bad ways to do form objects, so I've compiled techniques and patterns I think are most compliant with The Rails Way (remember the days when this was still a thing). Here's what it covers:

[Part 1 - Ground rules for form objects](https://medium.com/@jaryl/disciplined-rails-form-object-techniques-patterns-part-1-23cfffcaf429)

* Implementing `ActiveModel` and its lifecycle
* Jumping from virtual objects to `ActiveRecord` instances
* Working with multiple models in one form object using `fields_for`

[Part 2 - Dealing with Collections](https://medium.com/@jaryl/disciplined-rails-form-object-techniques-patterns-part-2-12b8d530143d)

* Collections of primitive values, to virtual objects, to `ActiveRecord` instances
* Taking control of nested\_attributes and your form object's interface
* Using `ActiveModel` validations and promoting errors to your UI

[Part 3 - Tying up loose ends](https://medium.com/@jaryl/disciplined-rails-form-object-techniques-patterns-part-3-8ed1e4f62ce4)

* Simplifying complex forms (like advanced search) by working with `Arel`
* Some additional thoughts on validations and nesting
* A look at form object (as a opposed to form builder) gems, and whether to use them

It's a bit of a slog, but I promise you it's less of a slog than wading through 1000-line ERB files on a daily basis!
## [9][Using Greek Alphabet](https://www.reddit.com/r/rails/comments/hr7j3b/using_greek_alphabet/)
- url: https://www.reddit.com/r/rails/comments/hr7j3b/using_greek_alphabet/
---
I have a couple of fields in my database that I need to be able to copy and paste Greek letters into and then display on their respective pages, is there an easy way to accomplish this?
## [10][Tips for building a "subscribe &amp; save" ecommerce store?](https://www.reddit.com/r/rails/comments/hr7azo/tips_for_building_a_subscribe_save_ecommerce_store/)
- url: https://www.reddit.com/r/rails/comments/hr7azo/tips_for_building_a_subscribe_save_ecommerce_store/
---
I'm looking to create a basic "subscribe and save" subscription for a physical product. At a high level, the user will signup for a monthly, 3, 6, 9 or 12 month subscription. Nothing groundbreaking or innovative there... (right?)

To date, all my RoR efforts have been for random side projects. Now that I'm thinking about dabbling with payments and physical products, I'm finding it to seem a little more daunting.

Should I be building on top of something like Spree? Or rolling my own billing, inventory, shipping code? Or approaching this problem a different way?

Edit: I realize both Spree and Solidus have subscription add-ons for products... I'm getting hung up on what happens if the subscription renews but, for example, my suppliers haven't sent me any inventory.
## [11][Renaming webpackers "javascript" folder...](https://www.reddit.com/r/rails/comments/hr3cw7/renaming_webpackers_javascript_folder/)
- url: https://www.reddit.com/r/rails/comments/hr3cw7/renaming_webpackers_javascript_folder/
---
I like webpacker but it annoys me that everything goes in a folder "javascript", stylesheets, (images?), etc.  Might this folder be better renamed to "webpacker"?  Is that possible?
## [12][Having a problem with has_many through joins](https://www.reddit.com/r/rails/comments/hr5k5s/having_a_problem_with_has_many_through_joins/)
- url: https://www.reddit.com/r/rails/comments/hr5k5s/having_a_problem_with_has_many_through_joins/
---
EDIT: Got it.

In my controller index method:

    @option_values = @variants.joins(option_values: {option_value_variants: :option_value}).uniq

In my index page:

     &lt;% @option_values.each do |option_type| %&gt;
       &lt;%= option_type.option_values.map {|ov| ov.name}.join(' , ') %&gt;
     &lt;% end %&gt;

&amp;#x200B;

Say I have these models:

    class OptionType &lt; ApplicationRecord
      belongs_to :product
      has_many :option_values, dependent: :destroy
    end
    
    class OptionValue &lt; ApplicationRecord
      belongs_to :option_type
      has_many :option_value_variants, dependent: :destroy
      has_many :variants, through: :option_value_variants
    end
    
    class Variant &lt; ApplicationRecord
      belongs_to :product
      has_many :option_value_variants, dependent: :destroy
      has_many :option_values, through: :option_value_variants
    end
    
    class OptionValueVariant &lt; ApplicationRecord
      belongs_to :option_value
      belongs_to :variant
    end

In the Variant index page I am trying to output a list of the option\_value names that are associated with that particular variant. How would I get that in the controller? Everything I have been trying produces active record join errors.

ovv = OptionValueVariant.first

ovv.option\_value gives me the correct answer for that particular one, but when I try to get all option\_value\_variants for a variant I get active record association join errors.

Any ideas?

    class Admin::VariantsController &lt; Admin::ApplicationController
      before_action :set_variant, only: [:show, :edit, :update, :destroy]
      before_action :set_product
      # GET /variants
      # GET /variants.json
      def index
        prodid = params[:product_id]
        if !prodid.nil?
          @variants = Variant.where(:product_id =&gt; prodid)
        else
          @variants = Variant.all
        end
        @option_types = @product.option_types
        @option_values = @variants.joins(option_values: {option_value_variants: :option_value}).uniq
      end
    
    
      # GET /variants/1
      # GET /variants/1.json
      def show
      end
    
      # GET /variants/new
      def new
        @variant = Variant.new
        @option_types = @product.option_types
      end
    
      # GET /variants/1/edit
      def edit
        @option_types = @product.option_types
      end
    
      # POST /variants
      # POST /variants.json
      def create
        @variant = Variant.new(variant_params)
    
        respond_to do |format|
          if @variant.save
    
            @option_types = @product.option_types
            format.html { redirect_to admin_product_variants_url, notice: 'Variant was successfully created.' }
            format.json { render :show, status: :created, location: @variant }
          else
            format.html { render :new }
            format.json { render json: @variant.errors, status: :unprocessable_entity }
          end
        end
      end
    
      # PATCH/PUT /variants/1
      # PATCH/PUT /variants/1.json
      def update
        @option_types = @product.option_types
        respond_to do |format|
          if @variant.update(variant_params)
    
            format.html { redirect_to admin_product_variants_url, notice: 'Variant was successfully updated.' }
            format.json { render :show, status: :ok, location: @variant }
          else
            format.html { render :edit }
            format.json { render json: @variant.errors, status: :unprocessable_entity }
          end
        end
      end
    
      # DELETE /variants/1
      # DELETE /variants/1.json
      def destroy
        @variant.destroy
        respond_to do |format|
          format.html { redirect_to admin_product_variants_url, notice: 'Variant was successfully destroyed.' }
          format.json { head :no_content }
        end
      end
    
    
      private
        # Use callbacks to share common setup or constraints between actions.
        def set_variant
          @variant = Variant.find(params[:id])
        end
        def set_product
          @product = Product.find(params[:product_id])
        end
        # Only allow a list of trusted parameters through.
        def variant_params
          params.require(:variant).permit(:sku, :price_cents, :price_currency, :is_master, :product_id, :count_in_stock, option_value_ids: [])
        end
    end

&amp;#x200B;
