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
## [2][What are you using for requests?](https://www.reddit.com/r/rails/comments/h0vva7/what_are_you_using_for_requests/)
- url: https://www.reddit.com/r/rails/comments/h0vva7/what_are_you_using_for_requests/
---
A lot of times API's won't have Ruby SDK's or gems.

I'm wondering what you all do in these cases...

&amp;#x200B;

I see some API's do:

    require 'uri' 
    require 'net/http' 
    require 'openssl' 

What is your normal procedure for doing this?

&amp;#x200B;

Basically, something similar to Pythons requests library
## [3][How to trigger JS directly after link_to in Rails 5](https://www.reddit.com/r/rails/comments/h0zaux/how_to_trigger_js_directly_after_link_to_in_rails/)
- url: https://www.reddit.com/r/rails/comments/h0zaux/how_to_trigger_js_directly_after_link_to_in_rails/
---
Hi - I'm looking to trigger some JS immediately after someone clicks a link_to button within my application. Tried doing the following with jQuery, but nothing is firing even if I just do a basic console.log(). Any idea what turbolinks:load should be replaced with?

    $(document).on('turbolinks:load', function(){
        console.log("test");
    });
## [4][Test User devise model with RSpec](https://www.reddit.com/r/rails/comments/h0ys91/test_user_devise_model_with_rspec/)
- url: https://www.reddit.com/r/rails/comments/h0ys91/test_user_devise_model_with_rspec/
---
Hello guys, I am testing user model with rspec. I want to test the database table columns. I added column firstname and lastname in user devise model. So i have the user\_spec.rb of:  
`describe "database table" do`

   `it { is_expected.to have_db_column(:id).of_type(:uuid) }`

   `it { is_expected.to have_db_column(:first_name).of_type(:string) }`

   `it { is_expected.to have_db_column(:last_name).of_type(:string) }`

   `it { is_expected.to have_db_column(:email).of_type(:string) }`

 `end`

And then when i rspec it the error is "User does not have a db column named last\_name and first\_name."  
I have a column of first\_name and last\_name in my database and in schema.  
How to solve that? Thankss.
## [5][ActiveRecord Help!](https://www.reddit.com/r/rails/comments/h0mrup/activerecord_help/)
- url: https://www.reddit.com/r/rails/comments/h0mrup/activerecord_help/
---
 

 **I have an sql query and I keep getting back this error** \----

result = ActiveRecord::Base.connection.execute("Select Countries.capital, Max(features\_shipped) AS feat, Countries.name, Max(age) As age, Teams.name As teamname From Engineers Inner Join Teams on Engineers.team\_id = Teams.id Inner Join Countries on Countries.id = Engineers.id Group by countries.name,countries.capital, teamname Order by age desc, feat desc").result.rows

&amp;#x200B;

***\*\*\*\*\*\* I keep getting back this error ----***

**Traceback** (most recent call last):

1: from (irb):71

**NoMethodError (undefined method \`result' for #&lt;PG::Result:0x00007f9a944d6918&gt;)**
## [6][Is it possible to have several domains on one rails app with Devise ? (and no SSO)](https://www.reddit.com/r/rails/comments/h0edb3/is_it_possible_to_have_several_domains_on_one/)
- url: https://www.reddit.com/r/rails/comments/h0edb3/is_it_possible_to_have_several_domains_on_one/
---
Hello,

I have a rails 5.2 app in production with domain www.mydomain.com

We would like to roll out a new product which use the heavy-lifting from the Rails app.
This product would be branded and used under www.anotherdomain.com

We host on Heroku. I have no problem to add the domain and access the app but I can't login through Devise. I don't even want to share cookies across domains - so I just have a quick question : it is doable to have several domains on one rails app with Devise ?
I don't want to implement SSO/ oauth.

Any other gotcha besides authentication ?

Thank you !
## [7][Tailing heroku with `heroku logs --tail`](https://www.reddit.com/r/rails/comments/h0caj8/tailing_heroku_with_heroku_logs_tail/)
- url: https://www.reddit.com/r/rails/comments/h0caj8/tailing_heroku_with_heroku_logs_tail/
---
  
I had two issues with the vanilla heroku logs --tail command 

   \- it wouldn't save it to a file  
   \- it would often abruptly stop, specially if you have it on for more than 30 mins  


I went about wrapping the command in a bash script like this  

```
#!/bin/bash
heroku logs --tail
```

saved it as ~/bin/heroku_logs_clientname
with a chmod of 755.
Now you can watch the logs by typing heroku_logs_clientname
in any shell prompt.
(assuming you ~/bin/ is in your PATH)

### Adding persistence

`tee` is a handy unix command. it copies whatever you're seeing to file.

```
#!/bin/bash
heroku logs --tail  | tee /tmp/heroku_logs_clientname.log
```

### Better persistence

In this current state, logs saved tomorrow would overwrite the logs saved
today. Lets' timestamp the filename to prevent overwriting:

```
#!/bin/bash
heroku logs --tail  | tee /tmp/heroku_logs_clientname_$(date +%s).log
```

### Handling timeouts

heroku log tails time out after a certain period of time. When that happens,
we want to sleep a bit and then re-initiate the log tailing. Using the `while`
command for this:

```
#!/bin/bash
while true; do echo sleep 1; sleep 1; heroku logs --tail| tee /tmp/heroku_logs_clientname_$(date +%s).log; echo sleep 300; sleep 300; done
```

Infinite loops are dangerous imho, except if you put in a generous sleep amount like 5 minutes.

### Here' my final script

```
#!/bin/bash
cd ~/clientname; while true; do echo sleep 1; sleep 1; heroku logs --app appname --tail 2&gt;&amp;1 | tee /tmp/heroku_logs_clientname_$(date +%s); echo sleep 300; sleep 300; done
```
## [8][Update Product without replacing Thumbnail by Default](https://www.reddit.com/r/rails/comments/h0ccbp/update_product_without_replacing_thumbnail_by/)
- url: https://www.reddit.com/r/rails/comments/h0ccbp/update_product_without_replacing_thumbnail_by/
---
Hi Folks,

I have a model where I attach a thumbnail image to a product when I upload it (through has_one_attached in model). No problem there and have configured it with an S3 bucket so that images are stored there. My issue becomes when I want to update an item through a form. By default it removes the image and I have to reupload each time which proves tedious. Any way to maintain that integrity by default as it does with other text or number fields?
## [9][rails-foundation issue](https://www.reddit.com/r/rails/comments/h0c5cl/railsfoundation_issue/)
- url: https://www.reddit.com/r/rails/comments/h0c5cl/railsfoundation_issue/
---
Hi folks,

I inherited a rails project 6 months ago and up until now i've not needed to make any SASS/CSS changes. The project is using the foundation-rails gem. So i setup a watcher in order to auto-generate the css whenever a scss/sass file is changed.

The watcer seems to be functioning ok, however, i get a fundamental error in the \_settings.scss file:  
\---------------------------------

cmd.exe /D /C call C:\\Users\\xazos\\AppData\\Roaming\\npm\\sass.cmd application.scss:application.css

Error: Undefined mixin.

╷

105 │ u/include add-foundation-colors;  
\------------------------------

I'm not sure why this mixin is undefined. It seems like a fundamental part of foundation. I've tried "re-installing" with the -force switch but it doesn't seem to make a difference. Any help would be massively appreciated!

Cheers!
## [10][Issue with Joining to Another Model in Admin Namespace](https://www.reddit.com/r/rails/comments/h0a27t/issue_with_joining_to_another_model_in_admin/)
- url: https://www.reddit.com/r/rails/comments/h0a27t/issue_with_joining_to_another_model_in_admin/
---
UPDATE: Got it working. I ended up adding admin\_category.references through a migration and saw that it added the admin\_categories\_id field to the DB. I had admin\_category\_id. After that I was able to pull the the category name via the rails console and hooked it up through the view. 

\_\_\_\_\_\_

I have products which belong to a category. Categories have many products. Both of these are in an admin namespace. I am trying to show the category name on the product show page when you look at the product but haven't been able to get it to work. What am I doing wrong here?

Category Model:

    class Admin::Category &lt; ApplicationRecord
      has_many :admin_products
    end

Product Model:

    class Admin::Product &lt; ApplicationRecord
      belongs_to :admin_category, class_name: "Admin::Category"
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
## [11][How I can import a multistep loginn to ROR?](https://www.reddit.com/r/rails/comments/h0ivs8/how_i_can_import_a_multistep_loginn_to_ror/)
- url: https://www.reddit.com/r/rails/comments/h0ivs8/how_i_can_import_a_multistep_loginn_to_ror/
---
Hello  I'a starter RoR  developer and I'm trying to import this [https://bbbootstrap.com/snippets/multi-step-form-wizard-30467045](https://bbbootstrap.com/snippets/multi-step-form-wizard-30467045)  multistep  login to rails with the studying goal, and when I  run the server it becomes like that and I can't jump to the next tabs of the form when I click next step: [https://imgur.com/IWAXX0C](https://imgur.com/IWAXX0C)  but why it  don't become like the snnippets?  The app  is here:  [https://github.com/LeoFragozo/wizard\_bootstrap](https://github.com/LeoFragozo/wizard_bootstrap),  can someone kindly please help me?

Here is the terminal: [https://imgur.com/UlWvCrN](https://imgur.com/UlWvCrN)
