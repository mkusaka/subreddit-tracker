# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gnbebg/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gnbebg/personal_projects_show_off_your_own_project_andor/
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
## [2][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/gs526t/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/gs526t/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [3][Ruby + TestUnit Setup](https://www.reddit.com/r/rails/comments/gtci9g/ruby_testunit_setup/)
- url: https://www.reddit.com/r/rails/comments/gtci9g/ruby_testunit_setup/
---
Hi everyone,  


Whenever I've done testing before, I've always done this in the context of a rails project using RSpec, and all the dependencies, imports, requires etc just kind of work.  


But I am about to begin a backend only coding challenge which requires only pure Ruby along with some TestUnit tests.  


So my question is, how do I set this up without all the unnecessary bloat that comes with rails? Is there an equivalent of `"rails new"` but that will just give me a pure Ruby setup with a testing framework ready to go? If not, how do I set this up manually without the "magic" of rails?  


Thanks.
## [4][Stuck in the beginner-intermediate phase of a rails developer.](https://www.reddit.com/r/rails/comments/gt5ofd/stuck_in_the_beginnerintermediate_phase_of_a/)
- url: https://www.reddit.com/r/rails/comments/gt5ofd/stuck_in_the_beginnerintermediate_phase_of_a/
---
Sorry if this is a cliche post. I've been working remotely as a part time web developer at a small media agency for about a year now. It's nice because it helps me pay for college, partially at least. 

Anyway, because I have three more years of this, I really want to be as good a rails developer as I can get. Especially because I would hate to lose this job half way through school. 

What resources or insight do y'all recommend or have to give? Ideally, I would like to read, but any video series or blog posts about all things rails would be much appreciated if you could share. 

The problem I have with books is that they are seemingly outdated. We're currently working on updating all of our websites from rails 4 to rails 6, finally. This means that I'd like to read books that are written in a rails 6 context, but so far there only appears to be one such book.

I've heard excellent things about Rails Anti-Patterns and the Rails 5 Way, but Anti-Patterns was written in Rails 3, I believe, and Rails 5 Way is, self evidently, one version out of date. 

I'm particularly interested in books dealing with software construction, such as Practical Object Oriented Design in Ruby by Sandi Metz, but I'm not sure if that book will necessarily make me a better rails developer. More specifically, I'm not sure if it will get me out of the beginner-intermediate phase that I'm currently in. 

I'm not sure which direction I'd have to go in to be a better rails developer, hence my posting here. I interested to hear your all's experience and what worked for you.

thanks!
## [5][should i replace my javascript_pack_tag 'application' with javascript_packs_with_chunks_tag 'application', 'data-turbolinks-track': 'reload' %&gt;](https://www.reddit.com/r/rails/comments/gt19vr/should_i_replace_my_javascript_pack_tag/)
- url: https://www.reddit.com/r/rails/comments/gt19vr/should_i_replace_my_javascript_pack_tag/
---
  
should i replace [my](https://github.com/la-ruby/web-common-core/blob/master/app/views/layouts/application.html.erb)   

    &lt;%= javascript_pack_tag 'application', 'data-turbolinks-track': 'reload' %&gt;

with

    &lt;%= javascript_packs_with_chunks_tag 'application', 'data-turbolinks-track': 'reload' %&gt;


to fix this issue:


```
       WARNING in asset size limit: The following asset(s) exceed the recommended size limit (244 KiB).
       This can impact web performance.
       Assets: 
         css/application-337604d1.css (300 KiB)
         js/application-dfe7703552d43888f93f.js (2.34 MiB)
         js/application-dfe7703552d43888f93f.js.gz (720 KiB)
         js/application-dfe7703552d43888f93f.js.map.gz (1.47 MiB)
       
       WARNING in entrypoint size limit: The following entrypoint(s) combined asset size exceeds the recommended limit (244 KiB). This can impact web performance.
       Entrypoints:
         application (2.64 MiB)
             css/application-337604d1.css
             js/application-dfe7703552d43888f93f.js
```
## [6][Ruby gem to work with SEPA Reason Codes](https://www.reddit.com/r/rails/comments/gszxjj/ruby_gem_to_work_with_sepa_reason_codes/)
- url: https://www.reddit.com/r/rails/comments/gszxjj/ruby_gem_to_work_with_sepa_reason_codes/
---
I wrote a super minimal gem to work with SEPA Reason Codes.

When  implementing payments with SEPA, I had a pain point of not finding  information regarding reason codes. So, I found and combined them into a  gem to make working with them easier.

[https://www.ramblingcode.dev/posts/sepa\_reason\_codes\_ruby\_gem/](https://www.ramblingcode.dev/posts/sepa_reason_codes_ruby_gem/)
## [7][Help with rails project that uses action cable](https://www.reddit.com/r/rails/comments/gt4to5/help_with_rails_project_that_uses_action_cable/)
- url: https://www.reddit.com/r/rails/comments/gt4to5/help_with_rails_project_that_uses_action_cable/
---
Any rails developers who are free and want to help a college student with a rails app please message me . I have been having a hard time using action cable and would love to get someones help.  I am making a customer service application. So basically right now I have a guest account and like 4 admin accounts hardcoded into a database. I want to have a different chatrooms for each admin so they can help the customers.  Right now I have only one chat room which everyone enters when they log in which is not what a customer service app should be like . Right now it’s more like a group chat where everyone is in the same room and I want to make it where there are chatrooms for each admin. Any help is appreciated thank you!
## [8][I've got a bootstrap template/theme to import in RoR 6](https://www.reddit.com/r/rails/comments/gt48mz/ive_got_a_bootstrap_templatetheme_to_import_in/)
- url: https://www.reddit.com/r/rails/comments/gt48mz/ive_got_a_bootstrap_templatetheme_to_import_in/
---
What should I do? I have to install bootstrap? I've see something at drifting ruby but it's paid content, and my country's currency is literally bs in dollar value comparison? there are some free tutorials?  I don't need the whole template but some parts, I'm new to web dev, so please be considerate of it.

Thanks in advance!
## [9][Gem to Handler Common Errors in Client's Accepted Format?](https://www.reddit.com/r/rails/comments/gt5kp0/gem_to_handler_common_errors_in_clients_accepted/)
- url: https://www.reddit.com/r/rails/comments/gt5kp0/gem_to_handler_common_errors_in_clients_accepted/
---
I thought I saw something recently to handle things such as `ActiveRecord::RecordNotFound`, `ActionController::UnknownFormat`, etc... and respond with a JSON or HTML error message based on client's `Accept` header.

I think it just used the appropriate value from `Rack::Utils::HTTP_STATUS_CODES` as the message. 

Does this ring a bell to anyone? I could not find it. I don't think it was any of these:

* https://github.com/hassox/rack-rescue
* https://github.com/kuboon/restful_error
* https://github.com/jamesstonehill/api_error_handler
* https://github.com/mirego/gaffe
* https://github.com/richpeck/exception_handler
* https://github.com/yuki24/rambulance

Though the last 3 do look very nice!
## [10][Simple search form finds text_field nill class, confused, please help!](https://www.reddit.com/r/rails/comments/gszvwd/simple_search_form_finds_text_field_nill_class/)
- url: https://www.reddit.com/r/rails/comments/gszvwd/simple_search_form_finds_text_field_nill_class/
---
&amp;#x200B;

[I don't understand why it is bill tIhought the form tag can take anything? ](https://preview.redd.it/htwx52x4cr151.png?width=1920&amp;format=png&amp;auto=webp&amp;s=517c6d567b3303f7e3f59da837c858a16ec90eb0)

&amp;#x200B;

    #This is the index for the articles controller
    
    &lt;div class="container"&gt;
      &lt;div class="row"&gt;
        &lt;div class="col"&gt;
         &lt;%= form_tag search_articles_path do |f| %&gt;
            &lt;%= f.text_field :search %&gt;
            &lt;%= f.submit "submit" %&gt;
         &lt;% end %&gt;
        &lt;/div&gt;
      &lt;/div&gt;
    &lt;/div&gt;
    

&amp;#x200B;

    #This is in my article model
     
    def self.search(params)
            articles = Article.where("body LIKE ? or title LIKE ?", "%#{params[:search]}%", "%#{params[:search]}%")
            if params[:search].present?
                articles
            end
        end

&amp;#x200B;

    #This is the articles controller
    
    class ArticlesController &lt; ApplicationController
      before_action :set_article, only: [:show, :edit, :update, :destroy, :search]
    
      # GET /articles
      # GET /articles.json
      def index
        @articles = Article.all
        
      end
    
      def search
        if params[:search].blank?
          @articles = Article.all
    
        else 
          @articles = Article.search(params)
        end
      end
    
      # GET /articles/1
      # GET /articles/1.json
      def show
        @user = current_user.id
        @account = Account.find_by(user_id: @user)
        @company = Company.find_by(account_id: @account.id)
        @art = Article.friendly.find(params[:id])
        @category = Category.find_by(article_id: "#{@art.id}")
        @advert = Advert.find_by(article_id: "#{@art.id}")
      end
    
      # GET /articles/new
      def new
        @article = Article.new
        @article.category.build
        
        @user = current_user.id
        @account = Account.find_by(user_id: @user)
        @article.advert.build
        
      end
    
      # GET /articles/1/edit
      def edit
      end
    
      # POST /articles
      # POST /articles.json
      def create
        @article = Article.new(article_params)
    
        respond_to do |format|
          if @article.save
            format.html { redirect_to @article, notice: 'Article was successfully created.' }
            format.json { render :show, status: :created, location: @article }
          else
            format.html { render :new }
            format.json { render json: @article.errors, status: :unprocessable_entity }
          end
        end
      end
    
      # PATCH/PUT /articles/1
      # PATCH/PUT /articles/1.json
      def update
        respond_to do |format|
          if @article.update(article_params)
            format.html { redirect_to @article, notice: 'Article was successfully updated.' }
            format.json { render :show, status: :ok, location: @article }
          else
            format.html { render :edit }
            format.json { render json: @article.errors, status: :unprocessable_entity }
          end
        end
      end
    
      # DELETE /articles/1
      # DELETE /articles/1.json
      def destroy
        @article.destroy
        respond_to do |format|
          format.html { redirect_to articles_url, notice: 'Article was successfully destroyed.' }
          format.json { head :no_content }
        end
      end
    
      private
        # Use callbacks to share common setup or constraints between actions.
        def set_article
          @article = Article.friendly.find(params[:id])
        end
    
        # Only allow a list of trusted parameters through.
        def article_params
          params.require(:article).permit(
            :title,
            :body,
            :image,
            :advetise,
            :description,
            :account_id,
            :search,
            category_attributes: [
              :category_name,
              :id,
              :article_id
            ],
            advert_attributes: [
              :title,
              :id,
              :description,
              :web_link,
              :article_id
            ]
          )
        end
    end
## [11][Can we parallel upload to s3 while store is file system using carrierwave?](https://www.reddit.com/r/rails/comments/gsr6wp/can_we_parallel_upload_to_s3_while_store_is_file/)
- url: https://www.reddit.com/r/rails/comments/gsr6wp/can_we_parallel_upload_to_s3_while_store_is_file/
---
Hi,

We are using `carrierwave` uploader to upload files and storing them to file system. Now we are thinking to move them to s3 from file system. But before doing that, we want a parallel upload to local file system and s3 as well. So when I have a upload like this:

    # encoding: utf-8
    
    class FileUploader &lt; CarrierWave::Uploader::Base
    
      # Include RMagick or MiniMagick support:
      # include CarrierWave::RMagick
      include CarrierWave::MiniMagick
      # include CarrierWave::Processing::MiniMagick
    
      # Choose what kind of storage to use for this uploader:
      # storage :file
      storage :file
    
    
      # Override the directory where uploaded files will be stored.
      # This is a sensible default for uploaders that are meant to be mounted:
      def store_dir
        # "uploads/#{model.class.to_s.underscore}/#{mounted_as}/#{model.id}"
        "#{model.class.to_s.underscore}/#{model.company.id}/#{mounted_as}/#{model.id}"
      end
    
      version :report, if: :is_image? do
        process :quality =&gt; 80
        process resize_to_fill: [250, 250]
      end
    
      def is_image?(image)
        # causing error, so commented it out for now.
        ## ["png", "jpg", "jpeg"].any? {|mime| image.content_type.include?(mime)}
      end
    
      # Provide a default URL as a default if there hasn't been a file uploaded:
      # def default_url
      #   # For Rails 3.1+ asset pipeline compatibility:
      #   # ActionController::Base.helpers.asset_path("fallback/" + [version_name, "default.png"].compact.join('_'))
      #
      #   "/images/fallback/" + [version_name, "default.png"].compact.join('_')
      # end
    
      # Process files as they are uploaded:
      # process :scale =&gt; [200, 300]
      #
      # def scale(width, height)
      #   # do something
      # end
    
      # Create different versions of your uploaded files:
      # version :thumb do
      #   process :resize_to_fit =&gt; [50, 50]
      # end
    
      # Add a white list of extensions which are allowed to be uploaded.
      # For images you might use something like this:
      # def extension_white_list
      #   %w(jpg jpeg gif png)
      # end
    
      # Override the filename of the uploaded files:
      # Avoid using model.id or version_name here, see uploader/store.rb for details.
      # def filename
      #   "something.jpg" if original_filename
      # end
    
    end

How do I upload to S3 in background when local filesystem store is done? I was checking the [store!](https://github.com/carrierwaveuploader/carrierwave/blob/0f733d25e2aa9c3541c4a04fb114ee526c5ec78e/lib/carrierwave/uploader/store.rb#L62) method, but I don't see any way to tell uploader that I want to upload to s3 even if the it is mentioned as `file` statically inside the code. Is there any of doing what I am looking for?
## [12][Active Record: Count users where orders with a status is higher &gt; 1 (repeat customers)](https://www.reddit.com/r/rails/comments/gst5ry/active_record_count_users_where_orders_with_a/)
- url: https://www.reddit.com/r/rails/comments/gst5ry/active_record_count_users_where_orders_with_a/
---
I'd like to count how many `one_time` customers I have as well as `repeat_customers` where an order's status (state machine, not scope) is the attribute that determines if the order was completed, cancelled or currently in progress.

I can count how many `users` have at least one completed `order` with:

`User.includes("orders").where(orders: {global_status: 'completed'}).distinct.count`

But what if I want to know how many customers have only 1 completed order?

Something like:

`User.joins(:orders).group('users.id').having('count(orders) = 1').where(orders: {global_status: 'completed'}).distinct.count`

Is there a simpler (more straightforward way) to do this?

Thanks for reading.
