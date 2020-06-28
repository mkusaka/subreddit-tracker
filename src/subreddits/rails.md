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
## [3][Opinions on various AWS hosting options - ElasticBeanstalk vs ECS/Fargate vs Chef](https://www.reddit.com/r/rails/comments/hgwy07/opinions_on_various_aws_hosting_options/)
- url: https://www.reddit.com/r/rails/comments/hgwy07/opinions_on_various_aws_hosting_options/
---
Hello everyone,

I've been tasked with building a RESTful API.

Initially, the prototype was various Lambda functions backed by AWS API Gateway

We are pivoting to Rails because of all the amazing bells and whistles that come with a framework.

Right now, I have the Rails API running on localhost. I am also running Sidekiq in a separate terminal session. We are leveraging PostgreSQL as our database.

We are ready to start deploying the application to our AWS environment. The team is highly-skilled in AWS and so we open to different types of deployment and hosting options. 

I'm hoping to receive some feedback on the Pros/Cons of hosting a Rails app via the following AWS Services

# Option #1

## Bootstrapped via Chef

This is our current design specification. We wrote some Terraform IaC to build the following resources

- Application Load Balancer

- Auto Scaling Group

- PostgreSQL RDS Instance

The Chef cookbook, bootstraps the instances in the ASG by performing the following commands

- Clones the source code repo

- Runs bundle install, db:migrate, etc

- Runs redis-server in daemon mode, sidekiq, and rails server

- Nginx fronted in each app server to perform reverse proxy



__Pros__ - Design and code is already done, we can begin testing deployment

__Cons__ - Coming from a Heroku background, it does not seem intuitive enough to perform rolling updates/blue green deployments. We will have to write some automation to perform that. And possible rollbacks seem difficult. And that is why I started to research Option #2


## Option #2 ElasticBeanstalk

ElasticBeanstalk looks like it will give us the closest usability like Heroku. We can simply run \`eb deploy\` at first, and then incorporate a CI/CD pipeline to perform the deployment for us. I like the fact that application versions are stored so that we can easily rollback if needed.

__Pros__

- Most "Heroku"-like service in AWS. 

- Leverage the eb cli for deployment

- Application Versions allow for easy rollback

__Cons__

- There will be development work needed to figure out .ebextensions, e.g. how to run sidekiq. A simple default Rails 6 app does not simply "deploy" to ElasticBeanstalk successfully out of the box. Theres problems with Node.js versions, webpacker, etc. 


## Option #3 ECS/Fargate

I'm all for containerizing our workloads to help mitigate against differences in envrionment e.g. localhost, dev, test, prod, etc. So thats why I started to look into ECS/Fargate. AWS has released a ruby gem for provisioning a cluster. But unfortunately, its in preview mode/experimental. I spent a few hours last night playing around with it. I believe it will take some development time to massage the CDK output to match our infrastructure requirements, e.g. naming conventions, tagging, instance sizes etc. It seems like every change I make is overwritten via `aws-rails-provisioner build`.

[https://github.com/awslabs/aws-rails-provisioner](https://github.com/awslabs/aws-rails-provisioner)

__Pros__
- Containerizing our application can mitigate differences between environments.
- ECS supports various deployment options
- Closets "single click" deployment option I've found

__Cons__
- The `aws-rails-provisioner` rubygem is in experimental stage. 
- The infrastructure output from the CDK does not match our requirements and will need to be updated and ensured that it is not overwritten via `aws-rails-provisioner build`


It's hard to look through a crystal ball to understand the long-term tradeoffs for each option. I am hoping to find the most "Heroku-like" experience so that we can iterate quickly and efficiently without having to write a bunch of automation code. 

Thank you
## [4][Rails &amp; Amazon SES for emails. Best practices and throttling?](https://www.reddit.com/r/rails/comments/hgrr27/rails_amazon_ses_for_emails_best_practices_and/)
- url: https://www.reddit.com/r/rails/comments/hgrr27/rails_amazon_ses_for_emails_best_practices_and/
---
Hi all,

For a Rails site I'd like to switch to Amazon SES for sending emails because we're having some trouble delivering to the big providers from our own servers.

Now, SES has a pretty low throttling threshold by default (1/sec) which is a problem due to the nature of my application. My users send out invoices once a month, in bursts of thousands. Once I exceed Amazon's rate, emails get rejected. I need to do some throttling / queuing *somewhere*.

The way I see it, I can do that in 2 different ways:

1. Have the Rails app talk directly to the SES SMTP and do the throttling in-app, perhaps through something like DelayedJob?
2. Have the Rails app talk to a local MTA like Exim, which sets up a relay to SES, including throttling and retries.

Options 2 seems to be more robust, but it requires yet another service to be configured and maintained.

&amp;#x200B;

Secondly, anyone know any good guides how to set up SES with SNS on AWS to get insight into the individual bounces, rejects, etc.? I found their original documentation a bit lacking.
## [5][Exploring Rails View Components](https://www.reddit.com/r/rails/comments/hgbe2n/exploring_rails_view_components/)
- url: https://www.reddit.com/r/rails/comments/hgbe2n/exploring_rails_view_components/
---
I recorded a short exploration of the view\_component gem from Github. I have to say I'm pretty excited about it. There's some good stuff going on in the Stimulus Reflex community around ViewComponentReflex([https://github.com/joshleblanc/view\_component\_reflex](https://github.com/joshleblanc/view_component_reflex)), so I figured I'd start here before jumping into that: [https://youtu.be/oKqJmEAn-X0](https://youtu.be/oKqJmEAn-X0)
## [6][Acts as Tracked gem: selectively track changes made on your AR models](https://www.reddit.com/r/rails/comments/hg8fe0/acts_as_tracked_gem_selectively_track_changes/)
- url: https://www.reddit.com/r/rails/comments/hg8fe0/acts_as_tracked_gem_selectively_track_changes/
---
Hi, i've made a gem to selectively track changes on AR models, where audited gem would be an overkill. ActsAsTracked can be plugged into ActiveRecord model, and then used whenever you need a history of changes and actors made on the record.

You can find docs in the [repository](https://github.com/ramblingcode/acts-as-tracked)

Blog post [here](https://www.ramblingcode.dev/posts/track_changes_on_your_activerecord_models/)

Example usage in this [project](https://github.com/ramblingcode/rails6-acts-as-tracked-usage)

Hope you find it useful.
## [7][Glimmer DSL for Opal v0.0.8 (Rails Web GUI Adapter for Desktop Apps)](https://www.reddit.com/r/rails/comments/hgb3q5/glimmer_dsl_for_opal_v008_rails_web_gui_adapter/)
- url: https://www.reddit.com/r/rails/comments/hgb3q5/glimmer_dsl_for_opal_v008_rails_web_gui_adapter/
---
Glimmer DSL for Opal v0.0.8 is out with table data-binding support, including selection, sorting, and editing (can be easily styled with standard CSS). Glimmer DSL for Opal is an experimental web GUI adaptor for webifying [Glimmer](https://github.com/AndyObtiva/glimmer) desktop apps (i.e. apps built with [Glimmer DSL for SWT](https://github.com/AndyObtiva/glimmer-dsl-swt)) via [Rails](https://rubyonrails.org/) and [Opal](https://opalrb.com/) without changing a line of code. Apps may then be custom-styled for the web via standard CSS.

[https://github.com/AndyObtiva/glimmer-dsl-opal](https://github.com/AndyObtiva/glimmer-dsl-opal)

&amp;#x200B;

[Glimmer DSL for SWT \(Desktop App\)](https://preview.redd.it/3bapevj4aa751.png?width=393&amp;format=png&amp;auto=webp&amp;s=a7d788bbe54350ef0378d507d72d3ed4c26be914)

[Glimmer DSL for Opal \(Rails Web App\)](https://preview.redd.it/asyhz831aa751.png?width=612&amp;format=png&amp;auto=webp&amp;s=df6751b0a364bf836524992f3c979a864e7604e2)
## [8][How can I get amount from Braintree in a safe way?](https://www.reddit.com/r/rails/comments/hgcqvv/how_can_i_get_amount_from_braintree_in_a_safe_way/)
- url: https://www.reddit.com/r/rails/comments/hgcqvv/how_can_i_get_amount_from_braintree_in_a_safe_way/
---
I'm currently trying to integrate Braintree into my app. Basically, I want to allow customers to top up their account's credit via PayPal. The way I wanted it to work would be that the customers visit `Billing#new`, are able to select/type the amount that they, then click the "Pay with PayPal" button and after that was successful they get their selected amount as credit to their account. 

After reading through the Braintree ["Get started guide"](https://developers.braintreepayments.com/start/overview) I got nearly everything working in sandbox mode. The only thing that I don't have is a safe way to get the selected/typed amount from the client. Basically what they teach in their starter guide and what I have done so far is to create a form in `Billing#new` where they render the PayPal button via their JS SDK. After the customer has clicked on the button and payed with PayPal, the SDK provides a callback function where I then submit the form with the payment_nonce to `Billing#create`. Inside `Billing#create` I currently have this code:

&gt; result = gateway.transaction.sale(amount: "10.00", payment_method_nonce: nonce_from_the_client)

This code creates the entire transaction but hardcodes the required amount. [In their example app](https://github.com/braintree/braintree_rails_example/blob/master/app/controllers/checkouts_controller.rb#L22) they handle the amount via `params[:amount]` but add a note to not do this in production. But they don't say what to use instead.

I'm feeling a little bit lost here. They say not to use `params[:amount]` but every other option that I have explored quickly falls apart too. I'd like to use something like webhooks but can't find any usage example in their documentation or online.
## [9][How to access a variable from a ruby controller action in JS?](https://www.reddit.com/r/rails/comments/hgcijg/how_to_access_a_variable_from_a_ruby_controller/)
- url: https://www.reddit.com/r/rails/comments/hgcijg/how_to_access_a_variable_from_a_ruby_controller/
---
I have a variable created in an action within one of my ruby controllers. I can access it from the corresponding view, but how can I pass it so that my JS can grab it? Does it need to be embedded in the view as a hidden tag/variable?
## [10][Rails: One-to-One relationship not working](https://www.reddit.com/r/rails/comments/hg9kdj/rails_onetoone_relationship_not_working/)
- url: https://www.reddit.com/r/rails/comments/hg9kdj/rails_onetoone_relationship_not_working/
---
I am working on a Rails project and I am using namespacing for the models and controllers. That is the child models (`admin` and `student`) are put in a directory called `user` and the controllers are put in a directory called `users`.

I also have `admins_controller` and `students_controller` that use the `admin` model and the `student` model respectively. These controllers are namespaced using  `users` directory.

&amp;#x200B;

I then have a `personal_info` model that contains more details about the `user`, such as gender, age, date of birth. The `personal_info` table has a **one-to-one** relationship with the `user` model.

&amp;#x200B;

**Here's my code**;

&amp;#x200B;

**Personal Info model**:

    class PersonalInfo &lt; ApplicationRecord
    
      belongs_to :user
    end

&amp;#x200B;

**User model**:

    class User &lt; ApplicationRecord
    
      has_secure_password
      has_one :personal_info, class_name: 'PersonalInfo', dependent: :destroy
      accepts_nested_attributes_for :personal_info, allow_destroy: true
    end

&amp;#x200B;

**Admin model**:

    class User::Admin &lt; User
    end

&amp;#x200B;

**Admin Controller**:

    class Users::AdminsController &lt; ApplicationController
      before_action :set_admin, only: [:show, :edit, :update, :destroy]
    
      # GET /admins
      # GET /admins.json
      def index
        @admins = User::Admin.all
      end
    
      # GET /admins/1
      # GET /admins/1.json
      def show
      end
    
      # GET /admins/new
      def new
        @admin = User::Admin.new
        @admin.build_personal_info
      end
    
      # GET /admins/1/edit
      def edit
      end
    
      # POST /admins
      # POST /admins.json
      def create
        @admin = User::Admin.new(admin_params)
        @admin.build_personal_info
    
        respond_to do |format|
          if @admin.save
            format.html { redirect_to users_admin_path(@admin), notice: 'Admin was successfully created.' }
            format.json { render :show, status: :created, location: users_admin_path(@admin) }
          else
            format.html { render :new }
            format.json { render json: @admin.errors, status: :unprocessable_entity }
          end
        end
      end
    
      # PATCH/PUT /admins/1
      # PATCH/PUT /admins/1.json
      def update
        respond_to do |format|
          if @admin.update(admin_params)
            format.html { redirect_to users_admin_path(@admin), notice: 'Admin was successfully updated.' }
            format.json { render :show, status: :ok, location: users_admin_path(@admin) }
          else
            format.html { render :edit }
            format.json { render json: @admin.errors, status: :unprocessable_entity }
          end
        end
      end
    
      # DELETE /admins/1
      # DELETE /admins/1.json
      def destroy
        @admin.destroy
        respond_to do |format|
          format.html { redirect_to users_admins_url, notice: 'Admin was successfully destroyed.' }
          format.json { head :no_content }
        end
      end
    
      private
      # Use callbacks to share common setup or constraints between actions.
      def set_admin
        @admin = User::Admin.find(params[:id])
      end
    
      # Only allow a list of trusted parameters through.
      def admin_params
        params.require(:user_admin).permit(
          :email, :password, :role_id, 
          personal_info_attributes: [ :id, :first_name, :last_name, :phone, 
                                       :gender, :dob, :address, :city, :state, 
                                       :country ]
        )
      end
    end

&amp;#x200B;

**Routes**:

    Rails.application.routes.draw do
     namespace :users do
      resources :admins
      resources :students
     end
    end

&amp;#x200B;

**\_form.html.erb:**

    &lt;% if @admin.errors.any? %&gt;
      &lt;div id="error_explanation"&gt;
        &lt;h2&gt;&lt;%= pluralize(@admin.errors.count, "error") %&gt; prohibited this admin from being saved:&lt;/h2&gt;
    
        &lt;ul&gt;
          &lt;% @admin.errors.full_messages.each do |message| %&gt;
            &lt;li&gt;&lt;%= message %&gt;&lt;/li&gt;
          &lt;% end %&gt;
        &lt;/ul&gt;
      &lt;/div&gt;
    &lt;% end %&gt;
    
    &lt;%= fields_for :personal_info do |form| %&gt;
      &lt;div class="field"&gt;
        &lt;%= form.label :first_name %&gt;
        &lt;%= form.text_field :first_name %&gt;
      &lt;/div&gt;
    
      &lt;div class="field"&gt;
        &lt;%= form.label :last_name %&gt;
        &lt;%= form.text_field :last_name %&gt;
      &lt;/div&gt;
    &lt;% end %&gt;
    
    &lt;div class="field"&gt;
      &lt;%= form.label :email %&gt;
      &lt;%= form.text_field :email %&gt;
    &lt;/div&gt;
    
    &lt;%= fields_for :personal_info do |form| %&gt;
      &lt;div class="field"&gt;
        &lt;%= form.label :phone %&gt;
        &lt;%= form.text_field :phone %&gt;
      &lt;/div&gt;
    &lt;% end %&gt;
    
    &lt;div class="field"&gt;
      &lt;%= form.label :password %&gt;
      &lt;%= form.text_field :password %&gt;
    &lt;/div&gt;
    
    &lt;div class="actions"&gt;
      &lt;%= form.submit %&gt;
    &lt;/div&gt;

&amp;#x200B;

**new.html.erb**:

    &lt;h1&gt;New Admin&lt;/h1&gt;
    
    
    &lt;%= form_with(model: @admin, url: users_admins_path, local: true) do |form| %&gt;
      &lt;%= render partial: 'form', admin: @admin, locals: { form: form } %&gt;
    &lt;% end %&gt;
    
    &lt;%= link_to 'Back', users_admins_path %&gt;

&amp;#x200B;

However, when I try to create a new `admin` or update an already existing admin after adding inputs to the displayed form no `Personal Info` data is saved on the database. They are all **nil**.

&amp;#x200B;

    PersonalInfo Load (0.3ms) SELECT "personal_infos".* FROM "personal_infos" WHERE "personal_infos"."user_id" = $1 LIMIT $2 [["user_id", 6], ["LIMIT", 1]]
    
    =&gt; #&lt;PersonalInfo id: 2, first_name: nil, last_name: nil, phone: nil, gender: nil, dob: nil, address: nil, city: nil, state: nil, country: nil, user_id: 6, created_at: "2020-06-26 13:37:16", updated_at: "2020-06-26 13:37:16"&gt;

&amp;#x200B;

I have tried to get this resolved, but no luck yet. Any form of help will be highly appreciated.
## [11][Help choosing model associations](https://www.reddit.com/r/rails/comments/hgb12a/help_choosing_model_associations/)
- url: https://www.reddit.com/r/rails/comments/hgb12a/help_choosing_model_associations/
---
So i have a multi user blog platform, it's a simple crud app but my issue is with the association because i have many types of users (i have simple user that can write blog post and it show his name) and i have a company account (which is like the user but have other information like company website instead of the user username)


my issue is how to handle the association (i already implemented the user model using device) which have_many blog 

so every blog have the belong to associated to the user

how to add another user type (Company) which have the the same many blog 

i searched and found the STI but it's not what i need as there is differences between the users types ( i may add another type later) 



should i use polymorphic association or are there another method to this relationship    





TL;DR i have user, CompanyUser, blablaUser which all of them can make blog post how to make the association ?
## [12][I need help logging someone out using devise on the backend, and react on the frontend.](https://www.reddit.com/r/rails/comments/hgagz6/i_need_help_logging_someone_out_using_devise_on/)
- url: https://www.reddit.com/r/rails/comments/hgagz6/i_need_help_logging_someone_out_using_devise_on/
---
Hello all,

I'm making a simple messaging application using React as my frontend and devise/rails as my backend.  Making a fetch call to sign in and register work just fine but signing out does not.  Whenever I make a delete request to /users/sign\_out I get the following error: 

Filter chain halted as :verify\_signed\_out\_user rendered or redirected.

My fetch call is written as such: 

 

const *logout* = (*event*) =&gt; {  
 *event.preventDefault*();  
 *fetch*('http://localhost:3000/users/sign\_out'  
   
 , {  
        *method*: 'DELETE',  
        *headers*: {  
 'Content-type': 'application/json',  
 'Accept': 'application/json'  
 },  
        *body*: *JSON.stringify*({ *data*: *JSON.parse*(*localStorage.getItem*('user'))})  
 }  
   
 )  
    *.then*(*res* =&gt; *res.json*())  
    *.then*(*resp* =&gt; *console.log*(resp))  
}

I cannot seem to find much help elsewhere online, and any help is much appreciated!
