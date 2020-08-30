# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/igyvm1/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/igyvm1/personal_projects_show_off_your_own_project_andor/
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
## [2][Rails 5.2.4, Active Storage, and S3: is it possible to implement an effective IAM/ACL policy when the app's default opinion dumps everything into the root of a single bucket?](https://www.reddit.com/r/rails/comments/ij5e6v/rails_524_active_storage_and_s3_is_it_possible_to/)
- url: https://www.reddit.com/r/rails/comments/ij5e6v/rails_524_active_storage_and_s3_is_it_possible_to/
---
I have switched from Paperclip to Active Storage. While Active Storage is great, it presents a new set of problems. Paperclip let me set a bucket directory (or even a new bucket), whereas Active Storage dumps everything into the root of a single bucket. I fear that this could create security issues. Improperly set IAM policies are an easy to make no-no, and my app had been using two buckets to avoid the problem: one for public content, and one for private admin accessible content.

In making that two bucket setup work with Active Storage, I tried to use the `service` option for `has_one_attached`. However, I get an ArgumentError: `unknown keyword: service`. The [edge guide](https://edgeguides.rubyonrails.org/active_storage_overview.html#setup) implies that this option works in Rails 5.2, but the [guide for 5.2.4](https://guides.rubyonrails.org/v5.2.4/active_storage_overview.html#has-one-attached) doesn't mention the `service` option at *all*.

How do you lock down access to certain S3 files while being less restrictive with everything else? Is this even possible? Do I have to move to Shrine?
## [3][Beginner ActiveStorage Help](https://www.reddit.com/r/rails/comments/ij4hh0/beginner_activestorage_help/)
- url: https://www.reddit.com/r/rails/comments/ij4hh0/beginner_activestorage_help/
---
# Routes

    
    # routes.rb
      resources :properties,          only: [:new, :create, :edit, :update, :index, :show] do
        resources :expense_items
      end
    

# Model

    # model
    class ExpenseItem &lt; ApplicationRecord
      has_many_attached :attachments
    end

    # Works fine in rails console:
    @expense_item.attachments.map { |f| f.filename } 
    # =&gt;[#&lt;ActiveStorage::Filename:0x000056033874b940 @filename="test1.png"&gt;, 
      #&lt;ActiveStorage::Filename:0x0000560338748d08 @filename="test2.mov"&gt;, 
      #&lt;ActiveStorage::Filename:0x000056033893e540 @filename="test3.pdf"&gt;] 


# Error
But when I try to use it in my edit view, I get an error saying:
    
    undefined method `service_name' for #&lt;ActiveStorage::Blob:0x00007fdd4844eaa8&gt;
    Did you mean?  service

    
# View

    # edit view
      &lt;div class="col-sm-6"&gt;
        &lt;%= render 'shared/error_messages', object: @expense_item %&gt;
          &lt;ul&gt;
            &lt;% @expense_item.attachments.each do |file| %&gt;
              &lt;li&gt;
                &lt;% if file.previewable? %&gt;
                  &lt;p&gt;file.filename&lt;/p&gt;
                  &lt;%#= image_tag file.preview(resize_to_limit: [100,100]) %&gt;
                &lt;% end %&gt;
              &lt;/li&gt;
            &lt;% end %&gt;
          &lt;/ul&gt;      &lt;%= f.submit 'Update Expense', class: 'btn btn-primary mt-sm-3' %&gt;
        &lt;% end %&gt;

The exception gets thrown at the if file.previewable? line.

# Controller

    # expense items controller
      def show
        @expense_item = property.expense_items.find(params[:id])
      end  

      def edit
        @expense_item = property.expense_items.find(params[:id])
      end
  
      def update
        @expense_item = property.expense_items.find(params[:id])
        if @expense_item&amp;.update_attributes(expense_item_params)
          flash[:success] = "Updated Expense Item: #{@expense_item.name}"
          redirect_to @expense_item.property
        else
          render 'edit'
        end
      end
    
      def expense_item_params
        params.require(:expense_item).permit(:name, :cost, :expense_date, :property_id, attachments: [])
      end
## [4][Javascript + Rails](https://www.reddit.com/r/rails/comments/ij4gvx/javascript_rails/)
- url: https://www.reddit.com/r/rails/comments/ij4gvx/javascript_rails/
---
I need help guys, i have to do a form where i select a commune (a big list of this) and, when i have to selectthe condominium  it only shows me those of the selected commune. I try with onChange(Javascript), but i dont know how get this value and use in the condominium place.

&amp;#x200B;

I really hope you can understand to me, is difficult to me explain this in english
## [5][Anyone taking Ruby on Rails courses?](https://www.reddit.com/r/rails/comments/iisg11/anyone_taking_ruby_on_rails_courses/)
- url: https://www.reddit.com/r/rails/comments/iisg11/anyone_taking_ruby_on_rails_courses/
---
Hey guys, anyone taking [*Ruby on Rails Courses*](https://www.google.com/amp/s/onlinecoursesgalore.com/ruby-rails-udemy/amp/) on Udemy?
## [6][Advice for multideployed app](https://www.reddit.com/r/rails/comments/iiynxg/advice_for_multideployed_app/)
- url: https://www.reddit.com/r/rails/comments/iiynxg/advice_for_multideployed_app/
---
Hi guys!

I have developed an application targeted to municipalities. Each municipality has its own server so in the end,  I will have one code base, and I will have sereval deployments to different servers. So far I have just one client and its working fine. I Believe in the future I will get more clients and I need to plan everything the best I can.

Using Capistrano, I have set a second environment, "staging" to test before production.wn server so in the end,  I will have one code base, and I will have sereval deployments to different servers. So far I have just one client and its working fine. I Believe in the future I will get more clients and I need to plan everything the best I can.

I won't lie: I feel I won't have enough control and knowledge before this step, but I need to grow the business.

Any advice or tips to manage all theses servers and configurations without getting crazy?  


Any pf you having a similar business model?
## [7][Seed upload local files from within Heroku](https://www.reddit.com/r/rails/comments/iiutoa/seed_upload_local_files_from_within_heroku/)
- url: https://www.reddit.com/r/rails/comments/iiutoa/seed_upload_local_files_from_within_heroku/
---
I have images that I need to be seeded on to my heroku database. here's how I achieve this on development server

&amp;#x200B;

    #seed.rb
    ...
    path = Dir.glob("path/to/images/logos/#{row['url']}/*.*").first
    if path file_io = File.open(path) 
    file_name = File.basename(path) 
    file_ext = file_name.split('.')[1] 
    t.logo.attach(io: file_io, filename: file_name, content_type: "image/{file_ext}") 
    end 
    ...

Would there be a way to do it on heroku production server?

UPDATE: Forgot to mention that i use GCS for file storage
## [8][I applied for a RoR job!](https://www.reddit.com/r/rails/comments/iiddg4/i_applied_for_a_ror_job/)
- url: https://www.reddit.com/r/rails/comments/iiddg4/i_applied_for_a_ror_job/
---
After leeching for almost 2 years, while learning and building own projects I applied for a REAL RoR job. I felt confident enough to do so. This is one of those rare companies who had a decent job description without requiring 5+ years experience with none existing language 

Thanks u/rails community for your alwasy friendly help

Wish me luck!
## [9][Issue with GCS Private Key using dotenv](https://www.reddit.com/r/rails/comments/iiqtkg/issue_with_gcs_private_key_using_dotenv/)
- url: https://www.reddit.com/r/rails/comments/iiqtkg/issue_with_gcs_private_key_using_dotenv/
---
I am trying to use active storage with GCS.  It all worked when I used keyfile.json, but when I wanted to make it work using dotenv-rails I am getting error with private key entry.

&amp;#x200B;

storage.yml

    ...
    google:
      service: GCS 
      project: &lt;%= ENV['GCS_PROJECT'] %&gt; 
      bucket: &lt;%= ENV['GCS_BUCKET'] %&gt; 
      credentials: type: "service_account" 
        project_id: &lt;%= ENV['GCS_PROJECT_ID'] %&gt; 
        private_key_id: &lt;%= ENV['GCS_PRIVATE_KEY_ID'] %&gt; 
        private_key: &lt;%= ENV['GCS_PRIVATE_KEY'] %&gt; 
        client_email: &lt;%= ENV['GCS_CLIENT_EMAIL'] %&gt; 
        client_id: &lt;%= ENV['GCS_CLIENT_ID'] %&gt; 
        auth_uri: "https://accounts.google.com/o/oauth2/auth" 
        token_uri: "https://accounts.google.com/o/oauth2/token"    
        auth_provider_x509_cert_url: "https://www.googleapis.com/oauth2/v1/certs"  
        client_x509_cert_url: &lt;%= ENV['GCS_CLIENT_X509_CERT_URL'] %&gt; 
    ...

.env

    ...
    GCS_PRIVATE_KEY="-----BEGIN PRIVATE KEY-----\n&lt; ACTUAL SECRET KEY &gt;\n-----END PRIVATE KEY-----\n"
    ...

&amp;#x200B;

error I get

    YAML syntax error occurred while parsing /Users/admin/Desktop/ruby/rails/launchzilla/config/storage.yml. Please note that YAML must be consistently indented using spaces. Tabs are not allowed. Error: (&lt;unknown&gt;): could not find expected ':' while scanning a simple key at line 29 column 1

I am pretty sure its the Private key part, because I have tested all the scenarios, and when I remove private key part the yml seem to parse corectly, but of course I get error when trying to upload, since GCS setting are incorrect then. Any ideas?
## [10][Newb trying to fix a CORS issue in WEBrick](https://www.reddit.com/r/rails/comments/iin972/newb_trying_to_fix_a_cors_issue_in_webrick/)
- url: https://www.reddit.com/r/rails/comments/iin972/newb_trying_to_fix_a_cors_issue_in_webrick/
---
I am getting a cors issue with DELETE, so I am trying to fix it, and can't understand how to do it with WEBrick. It looks like there is an application.rb file with environment development.rb. I am dealing with a legacy system I don't know about...

    'http://localhost:3006' has been blocked by CORS policy: Method DELETE is not allowed by Access-Control-Allow-Methods in preflight response.
    
    
    System:
      [2020-08-28 22:14:26] INFO  WEBrick 1.3.1 [2020-08-28 22:14:26] INFO 
      ruby 2.3.8 (2018-10-18) [x86_64-darwin19]

My development.rb looks like:

Where the heck do I add DELETE/PUT header info????

    *********::Application.configure do
      config.cache_classes = false
    
      config.whiny_nils = true
    
      config.logger = SkipResourcesLogger.new(File.join(Rails.root, 'log', 'development.log'))
    
      config.lograge.formatter = Lograge::Formatters::Logstash.new
      config.lograge.enabled = true
      config.lograge.custom_options = lambda do |event| stacktrace = {}
      if event.payload[:exception] &amp;&amp; strace = event.payload[:stacktrace]
          quoted_stacktrace = %('#{Array(strace).reject{|l| l.include?('/gems/') || 
      l.include?('/instrumenter')}.join("\n\t")}') # depending on if I'm logging 
      as logstash or not
          stacktrace = {stacktrace: quoted_stacktrace}
      end
        exceptions = %w(controller action format id)
        {
          id: event.payload[:ip],
          user_id: event.payload[:user_id],
          params: event.payload[:params].except(*exceptions)
        }.merge(stacktrace)
     end
    
      config.consider_all_requests_local       = true
      config.action_controller.perform_caching = true
    
      config.cache_store = :dalli_store, 'localhost:11211', { :compress =&gt; true, :value_max_bytes =&gt; 1024*1024*32 }
    
      config.active_support.deprecation = :log
    
      config.action_dispatch.best_standards_support = :builtin
    
      config.active_record.mass_assignment_sanitizer = :strict
    
      config.active_record.auto_explain_threshold_in_seconds = 0.5
    
      config.assets.compress = false
      config.assets.debug = true
    
       unless defined?(Rails::Console)
       end
       config.dependency_loading = true if $rails_rake_task
    
       config.assets.logger = false
    
       config.action_mailer.default_url_options = { host: 'localhost' }
       Rails.application.routes.default_url_options = { host: 'localhost' }
    
       config.after_initialize do
       Bullet.enable = true
       Bullet.alert = true
       Bullet.rails_logger = true
       Bullet.console = true
       Bullet.add_footer = true
      end
    
       config.action_controller.asset_host = "http://localhost"
    
     end
## [11][Why Rails and Not...](https://www.reddit.com/r/rails/comments/iia406/why_rails_and_not/)
- url: https://www.reddit.com/r/rails/comments/iia406/why_rails_and_not/
---
Hey all, I work for a company that's considering a major redo of their e-commerce platform. We're considering a bunch of frameworks for this: Rails, Laravel, Django, and Node. So I thought I'd ask the community here what would be some good reasons for Rails over these others? Thanks so much in advance. :)
