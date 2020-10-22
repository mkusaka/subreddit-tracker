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
## [3][Introducing Stimulus components with a first class support for Rails](https://www.reddit.com/r/rails/comments/jfvgem/introducing_stimulus_components_with_a_first/)
- url: https://www.reddit.com/r/rails/comments/jfvgem/introducing_stimulus_components_with_a_first/
---
Stimulus deserves to have a big and qualitative ecosystem with plug'n'play controllers like in other modern JS frameworks. 

More info here 👉 [https://guillaumebriday.fr/introducing-stimulus-components](https://guillaumebriday.fr/introducing-stimulus-components)

All the available controllers are here 👉 [https://github.com/stimulus-components](https://github.com/stimulus-components)

Feel free to open PRs and issues 🥳
## [4][Ruby on Rails: templates and generators in 2020](https://www.reddit.com/r/rails/comments/jfmas5/ruby_on_rails_templates_and_generators_in_2020/)
- url: https://www.reddit.com/r/rails/comments/jfmas5/ruby_on_rails_templates_and_generators_in_2020/
---
2020 is a being a very rich year for Rails **boilerplates and generators**. I've written a short article describing the most prominent ones: [https://blog.corsego.com/2020/10/ruby-on-rails-templates-and-generators.html](https://blog.corsego.com/2020/10/ruby-on-rails-templates-and-generators.html)

Hope you find it useful :)
## [5][Simple question I'm losing it](https://www.reddit.com/r/rails/comments/jfv6qz/simple_question_im_losing_it/)
- url: https://www.reddit.com/r/rails/comments/jfv6qz/simple_question_im_losing_it/
---
I have a table with a "status" string column where the values are either "pending" or "complete" and I'd like to know if there is a way to basically display two tables where each iterates each "status" value

EDIT: 
Only thing I've tried is:

model.where("column" =&gt; 'string') do 
end

EDIT 2: 
Okay I figured it out, all i dead was put the .each at the end
## [6][Anyone using Svelte,Elm or anything more obscure?How's your experience?](https://www.reddit.com/r/rails/comments/jfwvu2/anyone_using_svelteelm_or_anything_more/)
- url: https://www.reddit.com/r/rails/comments/jfwvu2/anyone_using_svelteelm_or_anything_more/
---

## [7][Hosting / Deployment Advice for Very Low Traffic Application](https://www.reddit.com/r/rails/comments/jfmy30/hosting_deployment_advice_for_very_low_traffic/)
- url: https://www.reddit.com/r/rails/comments/jfmy30/hosting_deployment_advice_for_very_low_traffic/
---
Hello good people,  


Having read through this post from yesterday and today:  
[https://www.reddit.com/r/rails/comments/jeo0xy/where\_do\_people\_host\_their\_rails\_6\_apps\_in\_2020/](https://www.reddit.com/r/rails/comments/jeo0xy/where_do_people_host_their_rails_6_apps_in_2020/)  


I thought I would pose a similar question, I run a website ([https://www.marincricketclub.com/](https://www.marincricketclub.com/)) for a sports club I'm a part of, our website lays mostly dormant for 5 months a year and the rest of the year, we would be lucky if we got 500 hits a month.   


We are a non profit so any way I can reduce cost is great! We use Heroku Hobby dyno's at $7 per month and they have been more than adequate for our use case and been zero hassle as well.  


But, after reading the above thread and having done some other research on my own, I believe I can save us some money by moving to the Digital Ocean App Marketplace (probably $5 per month) or Google Cloud Services which based on our use case, should be free or worst case scenario maybe cost $1 a month and this would include our attached storage bucket.   


I'm about to rebuild the existing site and have put together a (mostly) static site to serve as a place holder until the next season gets closer. Going forward, the app will send email, have an active storage solution and I want to get the deploy situation sorted now.  


So any advice, ideas or suggestions are very appreciated!  
Red
## [8][Same base project in three different instances. What do you suggest?](https://www.reddit.com/r/rails/comments/jfa5h0/same_base_project_in_three_different_instances/)
- url: https://www.reddit.com/r/rails/comments/jfa5h0/same_base_project_in_three_different_instances/
---
Let's say you have the project X which is a simple Rails app. It's open source in github. 

Then you decide you will host the project on a custom VPS with capistrano for example and sell it to people as a SaaS. 

And then, you build another instance with small differences and sell it as a business app on a different domain. 

What is the right way to accomplish that without going crazy? 

The use case is that I've build an app and I want to try and see If it gets any customers. It will have two plans, a personal and a business one. Obviously, the differences with the open source app would be to have extra models as `subscriptions` and more security. 

Should I keep the same *Core* and build everything else on top of gems and engines ? Or is there any other common solution I am missing ?

Thanks in advance!
## [9][Anyone know how to retrieve nested attributes with Grape?](https://www.reddit.com/r/rails/comments/jfmfqh/anyone_know_how_to_retrieve_nested_attributes/)
- url: https://www.reddit.com/r/rails/comments/jfmfqh/anyone_know_how_to_retrieve_nested_attributes/
---
I'm setting up my API, and its only returning one table -- when it should be including the nested relationships between my models. There are at least 6 different other relationships that I want to include with my initial api request. The are nested attributes for the provider\_form.

**api.rb**

&amp;#x200B;

`require 'grape-entity'`

`module Sims`

`class API &lt; Grape::API`

   `format :json`

   `prefix :api`

  `version 'v1', :path`

 `mount Sims::V1::ProviderForms`

 `end`

`end`

&amp;#x200B;

**provider\_form.rb**

`module Sims`

 `module V1`

  `class ProviderForms &lt; Sims::API`

  `include Grape::Kamari`

 `params do` 

   `use :pagination, per_page: 20, max_per_page: 30`

 `resources :provider_forms do`

 `desc 'return provider forms'` 

 `get do`

`present paginate(ProviderForm.all), with: PersonalInfoEntity`

`end` 

`end`

&amp;#x200B;

**personal\_info\_entity.rb**

`class PersonalInfoEntity &lt; Grape::Entity`

`expose :personal_info do`

`expose :first_name`

`expose :last_name`

  `end`

`end`

&amp;#x200B;

&amp;#x200B;

**provider\_form.rb**

`has_one :personal_info, dependent: :destroy`

**personal\_info.rb**

`belongs_to :provider_form`
## [10][Issue with unpermitted parameter](https://www.reddit.com/r/rails/comments/jfbvgm/issue_with_unpermitted_parameter/)
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
## [11][Find a Trending user on the basis of recently 3 days score](https://www.reddit.com/r/rails/comments/jfb8fj/find_a_trending_user_on_the_basis_of_recently_3/)
- url: https://www.reddit.com/r/rails/comments/jfb8fj/find_a_trending_user_on_the_basis_of_recently_3/
---
Hi, guys hope you are fine,

I'm implementing an  API in which against certain actions user win some points and these points are saving in a column in User table &gt; Score.  by `user.score += 100`    


but now I want to get the trending user who wins the maximum score recently in 3 days.

can anyone please guide me ?
## [12][New rails project postgres issue](https://www.reddit.com/r/rails/comments/jf2n05/new_rails_project_postgres_issue/)
- url: https://www.reddit.com/r/rails/comments/jf2n05/new_rails_project_postgres_issue/
---
 As the default configuration, a user called postgres is made and the user postgres has full superadmin access to the entire PostgreSQL instance running on your OS.The default Postgres user is postgres and a password is not required for authentication. Thus, to add a password, we must first login and connect as the postgres user   


In a new rails project, whenever I perform rake db it either tells me: 

 `peer authentication failed for user 'postgres'` or `no password supplied for postgres`

I shouldnt have to necessarily touch `pg_hba.conf` file or \password
