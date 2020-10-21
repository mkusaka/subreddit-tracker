# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/jbmadq/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/jbmadq/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/jfcv1r/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/jfcv1r/personal_projects_show_off_your_own_project_andor/
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
## [3][Same base project in three different instances. What do you suggest?](https://www.reddit.com/r/rails/comments/jfa5h0/same_base_project_in_three_different_instances/)
- url: https://www.reddit.com/r/rails/comments/jfa5h0/same_base_project_in_three_different_instances/
---
Let's say you have the project X which is a simple Rails app. It's open source in github. 

Then you decide you will host the project on a custom VPS with capistrano for example and sell it to people as a SaaS. 

And then, you build another instance with small differences and sell it as a business app on a different domain. 

What is the right way to accomplish that without going crazy? 

The use case is that I've build an app and I want to try and see If it gets any customers. It will have two plans, a personal and a business one. Obviously, the differences with the open source app would be to have extra models as `subscriptions` and more security. 

Should I keep the same *Core* and build everything else on top of gems and engines ? Or is there any other common solution I am missing ?

Thanks in advance!
## [4][Issue with unpermitted parameter](https://www.reddit.com/r/rails/comments/jfbvgm/issue_with_unpermitted_parameter/)
- url: https://www.reddit.com/r/rails/comments/jfbvgm/issue_with_unpermitted_parameter/
---
I am trying to build a stripped down version of a spree type application to learn. In my product form I want to be able to attached option\_types to it. I do this by having a select field that has select2 on it. I choose the option types needed from the list and it adds them in as a sort of tags type text field. I think it should come out in the form params as a comma delimited field based on how I have it set up. 

When I update a product after I have selected some option types the product updates but I am getting the following error on an unpermitted parameter:

&amp;#x200B;

https://preview.redd.it/zc3a48ojyfu51.png?width=2509&amp;format=png&amp;auto=webp&amp;s=c7ad26b2e9b688d14497e61ccc97620ac8f91104

I've tried everything and can't figure out why that is happening. Any thoughts on what the issue is based on the info below?

product controller:

    class ProductsController &lt; ApplicationController
      before_action :set_product, only: [:show, :edit, :update, :destroy]
      before_action :load_data
      # GET /products
      # GET /products.json
      def index
        u/products = Product.all
        u/option_types = OptionType.all.map{|c| [ c.name, c.id ] }
      end
    
      # GET /products/1
      # GET /products/1.json
      def show
        puts params.inspect
      end
    
      # GET /products/new
      def new
        u/product = Product.new
        u/product.option_types.new
        u/categories = Category.all.map{|c| [ c.title, c.id ] }
        u/variants = Variant.all.map{|c| [ c.name, c.id ] }
      end
    
      # GET /products/1/edit
      def edit
        u/categories = Category.all.map{|c| [ c.title, c.id ] }
        u/option_types = u/product.option_types
        u/product.option_types.new
      end
    
      # POST /products
      # POST /products.json
      def create
        u/product = Product.new(product_params)
        u/product.user_id = current_user.id
        u/product.category_id = params[:category_id]
        respond_to do |format|
          if u/product.save
            format.html { redirect_to u/product, notice: 'Product was successfully created.' }
            format.json { render :show, status: :created, location: u/product }
          else
            format.html { render :new }
            format.json { render json: u/product.errors, status: :unprocessable_entity }
          end
        end
      end
    
      # PATCH/PUT /products/1
      # PATCH/PUT /products/1.json
      def update
        u/product.category_id = params[:category_id]
    
        if u/product.option_types.present?
          u/product.option_types = u/product.option_types.split(',')
        end
    
        respond_to do |format|
          if u/product.update(product_params)
            format.html { redirect_to u/product, notice: 'Product was successfully updated.' }
            format.json { render :show, status: :ok, location: u/product }
          else
            format.html { render :edit }
            format.json { render json: u/product.errors, status: :unprocessable_entity }
          end
        end
      end
    
      # DELETE /products/1
      # DELETE /products/1.json
      def destroy
        u/product.destroy
        respond_to do |format|
          format.html { redirect_to products_url, notice: 'Product was successfully destroyed.' }
          format.json { head :no_content }
        end
      end
    
      private
        # Use callbacks to share common setup or constraints between actions.
        def set_product
          u/product = Product.find(params[:id])
        end
    
        def load_data
          u/option_types = OptionType.order(:name)
        end
    
        # Only allow a list of trusted parameters through.
        def product_params
          params.require(:product).permit(:name, :description, :stock, :available_on, :price, :user_id, :main_image, :option_type_ids, :option_type_id, option_type_attributes: [:id, :presentation])
        end
    end

product model:

    class Product &lt; ApplicationRecord
      belongs_to :user
      belongs_to :category
      has_one_attached :main_image, dependent: :destroy
      has_many_attached :images, dependent: :destroy
    
      has_many :product_option_types, dependent: :destroy
      has_many :option_types, through: :product_option_types
    
      accepts_nested_attributes_for :option_types, allow_destroy: true
    
      has_many :variants, inverse_of: :product
    
    end
    

option\_types model:

    class OptionType &lt; ApplicationRecord
      with_options dependent: :destroy, inverse_of: :option_type do
        has_many :option_values, -&gt; { order(:position) }
        has_many :product_option_types
      end
    
      accepts_nested_attributes_for :option_values, reject_if: :all_blank, allow_destroy: true
    
      has_many :products, through: :product_option_types
    end
    

product\_option\_types model:

    class ProductOptionType &lt; ApplicationRecord
      with_options inverse_of: :product_option_types do
          belongs_to :product
          belongs_to :option_type
      end
    
      validates :product, :option_type, presence: true
      validates :product_id, uniqueness: { scope: :option_type_id }, allow_nil: true
    end
    

product \_form

    &lt;%= form_with(model: product, local: true) do |form| %&gt;
      &lt;div class="row"&gt;
        &lt;div class="col-md-9"&gt;
    
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
    
      &lt;div class="form-group"&gt;
        &lt;%= form.label :name %&gt;
        &lt;%= form.text_field :name, class: 'form-control' %&gt;
      &lt;/div&gt;
    
      &lt;div class="form-group"&gt;
        &lt;%= form.label :description %&gt;
        &lt;%= form.text_area :description, class: 'form-control' %&gt;
      &lt;/div&gt;
    
      &lt;div class="form-group"&gt;
        &lt;%= form.label :stock %&gt;
        &lt;%= form.text_field :stock, class: 'form-control' %&gt;
      &lt;/div&gt;
    
      &lt;div class="form-group"&gt;
        &lt;%= form.label :available_on %&gt;
        &lt;%= form.text_field :available_on, class: 'form-control', data: { behavior: "flatpickr" } %&gt;
      &lt;/div&gt;
    
      &lt;div class="form-group"&gt;
        &lt;%= form.label :price %&gt;
        &lt;%= form.text_field :price, class: 'form-control' %&gt;
      &lt;/div&gt;
    
      &lt;div class="form-group"&gt;
        &lt;%= form.label 'Category' %&gt;
      &lt;%= select_tag(:category_id, options_for_select(@categories, u/product.category_id), class: 'form-control', :prompt =&gt; 'Select category') %&gt;
      &lt;/div&gt;
    
      &lt;div class="form-group" data-controller='select2'&gt;
        &lt;%= form.label :option_type_ids %&gt;
    
        &lt;%= form.select :option_type_id, OptionType.all.map { |type| type.presentation }, {include_blank: false}, class: 'form-control content-search', multiple: 'multiple' %&gt;
    
      &lt;/div&gt;
    
        &lt;hr&gt;
        &lt;div class="form-group"&gt;
        &lt;%= form.submit class: 'btn btn-primary' %&gt;
    
        &lt;% if product.persisted? %&gt;
        &lt;div class="float-right"&gt;
          &lt;%= link_to 'Destroy', product, method: :delete, class: "text-danger", data: { confirm: 'Are you sure?' } %&gt;
        &lt;/div&gt;
          &lt;%= link_to "Cancel", product, class: "btn btn-link" %&gt;
        &lt;% else %&gt;
          &lt;%= link_to "Cancel", products_path, class: "btn btn-link" %&gt;
        &lt;% end %&gt;
    
      &lt;/div&gt;
    &lt;/div&gt;
    
    &lt;div class="col-md-3"&gt;
      &lt;div class="form-group"&gt;
        &lt;%= form.label 'Main Product Image' %&gt;
        &lt;%= form.file_field :main_image, classs: 'form-control' %&gt;
      &lt;/div&gt;
    
      &lt;% if product.persisted? %&gt;
    
          &lt;div class="form-group"&gt;
              &lt;%= link_to "Manage Variants", product_variants_path(@product), class: "btn btn-link" %&gt;
          &lt;/div&gt;
    
        &lt;/div&gt;
        &lt;% end %&gt;
      &lt;/div&gt;
      &lt;% end %&gt;
## [5][Find a Trending user on the basis of recently 3 days score](https://www.reddit.com/r/rails/comments/jfb8fj/find_a_trending_user_on_the_basis_of_recently_3/)
- url: https://www.reddit.com/r/rails/comments/jfb8fj/find_a_trending_user_on_the_basis_of_recently_3/
---
Hi, guys hope you are fine,

I'm implementing an  API in which against certain actions user win some points and these points are saving in a column in User table &gt; Score.  by `user.score += 100`    


but now I want to get the trending user who wins the maximum score recently in 3 days.

can anyone please guide me ?
## [6][New rails project postgres issue](https://www.reddit.com/r/rails/comments/jf2n05/new_rails_project_postgres_issue/)
- url: https://www.reddit.com/r/rails/comments/jf2n05/new_rails_project_postgres_issue/
---
 As the default configuration, a user called postgres is made and the user postgres has full superadmin access to the entire PostgreSQL instance running on your OS.The default Postgres user is postgres and a password is not required for authentication. Thus, to add a password, we must first login and connect as the postgres user   


In a new rails project, whenever I perform rake db it either tells me: 

 `peer authentication failed for user 'postgres'` or `no password supplied for postgres`

I shouldnt have to necessarily touch `pg_hba.conf` file or \password
## [7][Where do people host their Rails 6 apps in 2020?](https://www.reddit.com/r/rails/comments/jeo0xy/where_do_people_host_their_rails_6_apps_in_2020/)
- url: https://www.reddit.com/r/rails/comments/jeo0xy/where_do_people_host_their_rails_6_apps_in_2020/
---
I’m a returning Rails developer from the 4.2 days. Been out a couple years. 

I used to use AWS Elastic Beanstalk, because we had other AWS resources that our app used.

But here I am in October 2020, trying to get even the Rails 6 “yay you’re on Rails” default app running on Beanstalk, and it’s an endless slog of deploy, fail, check the logs, google the error, try a fix, repeat.

Is Beanstalk the wrong solution now? How do people host Rails apps on AWS? I’ve never used containers - is that the approach I should take?
## [8][ask Rails: Where are these routes coming from?](https://www.reddit.com/r/rails/comments/jf187e/ask_rails_where_are_these_routes_coming_from/)
- url: https://www.reddit.com/r/rails/comments/jf187e/ask_rails_where_are_these_routes_coming_from/
---
Ran `rake routes` and this  

[https://i.imgur.com/ZzJKiv7.png](https://i.imgur.com/ZzJKiv7.png)  

Its a new project and I only used `rails g scaffold` to create some model   

I did nothing with action_mailbox, rails conductor, active_storage yet there a bunch of routes showing. presumably I expected them not to appear
## [9][Upgrading from rails 4.2.0 to 5.0.0 and Ruby from 2.2.4 to 2.5.8 - Circular dependency detected while autoloading constant](https://www.reddit.com/r/rails/comments/jew20l/upgrading_from_rails_420_to_500_and_ruby_from_224/)
- url: https://www.reddit.com/r/rails/comments/jew20l/upgrading_from_rails_420_to_500_and_ruby_from_224/
---
So I'm upgrading from 4.2 to 5.0.0 and it's not going so well.

Here's the stackoverflow question I posted: https://stackoverflow.com/questions/64428360/upgrading-from-rails-4-2-0-to-5-0-0-and-ruby-from-2-2-4-to-2-5-8-circular-depe

If anyone has any idea what I'm doing wrong I would be forever in your debt.

Basically I need to upgrade my ruby/rails version before heroku EOLs their 14.04 stack. I've follow guides on how to upgrade, but I'm obviously missing something and I can start a server, but it throws an exception about a circular dependency...and I have no idea how to solve it.

Thanks in advance.
## [10][Setting up a dev environment database to work on app locally?](https://www.reddit.com/r/rails/comments/jekdzx/setting_up_a_dev_environment_database_to_work_on/)
- url: https://www.reddit.com/r/rails/comments/jekdzx/setting_up_a_dev_environment_database_to_work_on/
---
Say you join a project and you need to set up your local dev environment. You install Rails, bundle install etc., but how exactly does the database portion work?

What I mean is, there's a production database that's pretty large with a couple dozen tables. When I work on the app locally, it won't run without connecting to a database for logins etc. Should I clone the whole production database and reproduce it locally, or what is the best practice here?
## [11][data model for charting stock-prices/changes?](https://www.reddit.com/r/rails/comments/jedezf/data_model_for_charting_stockpriceschanges/)
- url: https://www.reddit.com/r/rails/comments/jedezf/data_model_for_charting_stockpriceschanges/
---
Stocks are often charted out with prices and their given date of change. 

IF you have a single stock and wanted to show the price changing over time, how would you model that in your database? 

Sounds like TONS of data...

**edit!** 

Thanks for the comments. I ended up doing basically what u/UwRandom had recommended!

Main table has generic/aggregate information and a separate table stores the price changes.
## [12][Question: What are you guys using for centralized logging that is not papertrail app](https://www.reddit.com/r/rails/comments/jehhqd/question_what_are_you_guys_using_for_centralized/)
- url: https://www.reddit.com/r/rails/comments/jehhqd/question_what_are_you_guys_using_for_centralized/
---
Looking for an open sources solution to collect logs (centralize logging). Please suggest if you are using something in production. What has worked out well for you. 

\- graylog  
\- logstash  
\- fluend  
\- something else (hopefully)
