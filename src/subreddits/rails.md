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
## [2][How to implement refresh tokens?](https://www.reddit.com/r/rails/comments/hohl6i/how_to_implement_refresh_tokens/)
- url: https://www.reddit.com/r/rails/comments/hohl6i/how_to_implement_refresh_tokens/
---
Hey, I current use devise-jwt with an access token that expires in 2 weeks. I want to reduce the time it takes for access tokens to live, but I also don't want users logging in every 30 minutes. devise-jwt doesn't support refresh tokens out of the box. How do you folks do it?
## [3][Vagrant alternatives? Coding environment failures](https://www.reddit.com/r/rails/comments/hoeo0a/vagrant_alternatives_coding_environment_failures/)
- url: https://www.reddit.com/r/rails/comments/hoeo0a/vagrant_alternatives_coding_environment_failures/
---
Hi there. I'm just fed up with continuing trying to use Vagrant at this point. I've had many troubles in trying to build my own Vagrantbox for my Rails web dev environment. I had one going from what my bootcamp gave me, basically just a small file system with their vagrant box they provided and their own instructions for things. But sine a couple months ago I was trying to branch out and see what else is out there, and I think initially my issues started when I tried to use Docker. That's a whole other thing I won't get into, but will still probably try again sometime cause I was having issues there too. 

Anyway, I'm now kind of at my last resort using this box using Ubuntu 20.04 that's at this github link:  [https://github.com/rails/rails-dev-box](https://github.com/rails/rails-dev-box) 

I cloned it, and let it do its thing. But now it's getting hung up on the default: SSH auth method: private key line. Really this has been my main issue after setting up the box. The first time I run vagrant up, it gets hung up here. I've tried running ssh-keygen again to replace my keys in my .ssh folder, but it doesn't do anything. I'm still a newb though and may be missing something. 

Here's my full first "vagrant up" output: 

    C:\Users\Kyle_\Dev\rails-dev-box&gt;vagrant up
    Bringing machine 'default' up with 'virtualbox' provider...
    ==&gt; default: Box 'ubuntu/focal64' could not be found. Attempting to find and install...
        default: Box Provider: virtualbox
        default: Box Version: &gt;= 0
    ==&gt; default: Loading metadata for box 'ubuntu/focal64'
        default: URL: https://vagrantcloud.com/ubuntu/focal64
    ==&gt; default: Adding box 'ubuntu/focal64' (v20200707.0.0) for provider: virtualbox
        default: Downloading: https://vagrantcloud.com/ubuntu/boxes/focal64/versions/20200707.0.0/providers/virtualbox.box
    Download redirected to host: cloud-images.ubuntu.com
        default:
    ==&gt; default: Successfully added box 'ubuntu/focal64' (v20200707.0.0) for 'virtualbox'!
    ==&gt; default: Importing base box 'ubuntu/focal64'...
    ==&gt; default: Matching MAC address for NAT networking...
    ==&gt; default: Checking if box 'ubuntu/focal64' version '20200707.0.0' is up to date...
    ==&gt; default: Setting the name of the VM: rails-dev-box_default_1594339592541_33376
    ==&gt; default: Clearing any previously set network interfaces...
    ==&gt; default: Preparing network interfaces based on configuration...
        default: Adapter 1: nat
    ==&gt; default: Forwarding ports...
        default: 3000 (guest) =&gt; 3000 (host) (adapter 1)
        default: 22 (guest) =&gt; 2222 (host) (adapter 1)
    ==&gt; default: Running 'pre-boot' VM customizations...
    ==&gt; default: Booting VM...
    ==&gt; default: Waiting for machine to boot. This may take a few minutes...
        default: SSH address: 127.0.0.1:2222
        default: SSH username: vagrant
        default: SSH auth method: private key
    Timed out while waiting for the machine to boot. This means that
    Vagrant was unable to communicate with the guest machine within
    the configured ("config.vm.boot_timeout" value) time period.
    
    If you look above, you should be able to see the error(s) that
    Vagrant had when attempting to connect to the machine. These errors
    are usually good hints as to what may be wrong.
    
    If you're using a custom box, make sure that networking is properly
    working and you're able to connect to the machine. It is a common
    problem that networking isn't setup properly in these boxes.
    Verify that authentication configurations are also setup properly,
    as well.
    
    If the box appears to be booting properly, you may want to increase
    the timeout ("config.vm.boot_timeout") value.
    
    C:\Users\Kyle_\Dev\rails-dev-box&gt;
    

Any help would be great.  And honestly, if I can't get this to work, I would love other suggestions on how to set up an environment without Vagrant. I know I should still use VMs with Linux for Rails development. I was perfectly comfortable during bootcamp just vagrant up, vagrant ssh, and boom. IDK what happened and it's so discouraging :(
## [4][Efficiently searching for text from several .pdf](https://www.reddit.com/r/rails/comments/ho8g9m/efficiently_searching_for_text_from_several_pdf/)
- url: https://www.reddit.com/r/rails/comments/ho8g9m/efficiently_searching_for_text_from_several_pdf/
---
Hi,  
A friend of mine needs a search engine that searches from a lot of PDF. I though it could be a great challenge (I'm not a professional at all). I'm looking for advise since I've never deal with such a huge amount of data.  


Here's my plan:

* allow him to upload as many .pdf as he wants
* extract text from PDF using [this gem](https://github.com/yob/pdf-reader) (pdf-reader) and async jobs
* store extracted text into a database
* set up Elasticsearch to search from extracted text (never done that before)

Beyond the challenge, if there's any working (not necessarily online) tool, I'd glad to test it and share it with him

&amp;#x200B;

Thank you in advance ! :)
## [5][What is the best way to do full CRUD views for a model within a form of another?](https://www.reddit.com/r/rails/comments/ho6nwa/what_is_the_best_way_to_do_full_crud_views_for_a/)
- url: https://www.reddit.com/r/rails/comments/ho6nwa/what_is_the_best_way_to_do_full_crud_views_for_a/
---
Say I have a form for products and within products I want to be able create, update, delete option\_types. I would like to do all those actions within the product form. I have the create set up so it works, but when I transition between index, edit, show for option\_types it moves me out of the product form to the respective view pages. 

What is the best way to stay within the product form always? Iframe, or something else?
## [6][Job uniqueness for ActiveJob](https://www.reddit.com/r/rails/comments/hnmbq8/job_uniqueness_for_activejob/)
- url: https://www.reddit.com/r/rails/comments/hnmbq8/job_uniqueness_for_activejob/
---
The [activejob-uniqueness](https://github.com/veeqo/activejob-uniqueness) is an attempt to implement something similar to sidekiq-unique-jobs, but working on more high-level abstraction, like ActiveJob callbacks, what makes it compatible with any ActiveJob adapter (including Sidekiq). It uses [redlock-rb](https://github.com/leandromoreira/redlock-rb) (implementation of [Redlock algorithm](https://redis.io/topics/distlock)) and therefore depends on Redis.
## [7][Migrate WordPress site to Rails CMS](https://www.reddit.com/r/rails/comments/hnmder/migrate_wordpress_site_to_rails_cms/)
- url: https://www.reddit.com/r/rails/comments/hnmder/migrate_wordpress_site_to_rails_cms/
---
Hey everyone! 

I just finished a Full Stack Development bootcamp (based on the Rails framework) and I loved everything about Rails !

Before doing the bootcamp, I was working on my own website about dogs, built with WordPress. Now, I look at it and I think about all the stuff I could add with my Rails knowledge (instead of only being a pure content website, i could really add web app features, like a way to find the nearest vets, breeders, etc.)

I could try to learn some PHP and do it on WordPress but I honestly really don't want to do that, and Rails and PHP work differently.

I was then thinking about taking all the already existing content and migrate it in a Rails CMS. That could be doable but i dont really know which CMS suits my needs.

Do you guys have any experience with this? Thanks a lot in advance !
## [8][Ruby on rails and React stack](https://www.reddit.com/r/rails/comments/hnjauk/ruby_on_rails_and_react_stack/)
- url: https://www.reddit.com/r/rails/comments/hnjauk/ruby_on_rails_and_react_stack/
---
Just curious. While going for a Rails and React webapp which code architecture is recommended.

1. React stack in the same Rails codebase OR

2.  a separate Rails api code base and a separate React stack.  

When it comes to community support and efficiency, which one is recommended?
## [9][How to make friendly_id backfilling migration faster? You can skip all the callbacks.](https://www.reddit.com/r/rails/comments/hndvj5/how_to_make_friendly_id_backfilling_migration/)
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
## [10][Rails Frontend Development Resources](https://www.reddit.com/r/rails/comments/hnhez6/rails_frontend_development_resources/)
- url: https://www.reddit.com/r/rails/comments/hnhez6/rails_frontend_development_resources/
---
Hi everyone, 

I've been doing Ruby and RoR for a couple of years now, but have always worked on API only projects with little or no frontend at all. Im looking for resources to learn the concepts of structuring a frontend rails app, guides on how to properly use layouts, view helpers etc. If anyone has any tutorials, books or videos that would be helpful I would really appreciate if you could share them. Thanks
## [11][Best architecture for preferred sort for e-commerce products?](https://www.reddit.com/r/rails/comments/hn78ue/best_architecture_for_preferred_sort_for/)
- url: https://www.reddit.com/r/rails/comments/hn78ue/best_architecture_for_preferred_sort_for/
---
I have an ecommerce app on rails and was wondering what's the best architecture for ordering products on a page. I'm aware of how to use .order(:table_attribute), but was wondering if there is a cleaner/more dynamic way to bubble offerings to the top that you want to highlight. Any recommended practices for this?
