# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/i8dsvv/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/i8dsvv/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/igyvm1/personal_projects_show_off_your_own_project_andor/)
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
## [3][Setting up a blog on AWS Cloud9](https://www.reddit.com/r/rails/comments/igw3bw/setting_up_a_blog_on_aws_cloud9/)
- url: https://www.reddit.com/r/rails/comments/igw3bw/setting_up_a_blog_on_aws_cloud9/
---
Hi everyone,

I'm new to programming, and i've been following the upskill course for about 2 weeks now. Everything was going great untill yesterday. I started the Deep Dive: Build a Blog section that uses Cloud9 and Ruby to build a blog. I have tried many times to follow the exact same path as the instructor but i keep on having the same problem.

The commands that I input in the terminal are the following (I do this on a new environment with the default settings).

    $ rails new blog

Then the instructor ask us to change the 'sqlite3' line in the file called "gemfile" into gem 'sqlite3', '1.3.13'. To use the same version as him.

I then inpunt :

    cd
    
    cd environment
    
    cd blog
    
    bundle install
    
    bundle update
    
    rails generate scaffold Post title:string body:text
    
    rake db:migrate
    
    rails server

It is at this point that the problem occurs. The instructor terminal's response 3 more lines that i don't have (I don't think I can put screenshots so there is basically 3 lines after the line control c to shut down the server and their start with the date and the hour followed by some text)

Plus after that the instructor click preview running application and is sent to the "ruby welcome page". He then get rid of a part of the URL to go to the blog. In my case the URL is completely different from his, it looks like that ([https://c187d78accd944209c8f91023e991d71.vfs.cloud9.us-east-2.amazonaws.com/](https://c187d78accd944209c8f91023e991d71.vfs.cloud9.us-east-2.amazonaws.com/)).

Whereas his mostly made of words ( I think with the name of his environment) at the beggining and then numbers (which he erases to get to the blog).

Do you guys know what I've been doing wrong.

Thank you for reading all of these it's my first time posting so i hope it's kind of understandable.

Have a great day
## [4][Small script to reduce the image size of Ruby and Ruby on Rails Docker images](https://www.reddit.com/r/rails/comments/igxju5/small_script_to_reduce_the_image_size_of_ruby_and/)
- url: https://www.reddit.com/r/rails/comments/igxju5/small_script_to_reduce_the_image_size_of_ruby_and/
---
In the recent months I was migrating one of my Rails applications I was maintaining in the past years from capistrano to Docker. I did know that the gems are leaving files behind and therefore the Docker images became quite large but I was shocked when I realized how big the difference is.

So I sat down and I wrote a small gem called [cleanup\_vendor](https://github.com/raszi/cleanup_vendor) which cleans up the leftover files and reduces the Docker image size significantly.

Comments and suggestions are welcome.
## [5][ActiveStorage custom key for blobs](https://www.reddit.com/r/rails/comments/igvbx7/activestorage_custom_key_for_blobs/)
- url: https://www.reddit.com/r/rails/comments/igvbx7/activestorage_custom_key_for_blobs/
---
Is there a way to specify custom keys for blobs in ActiveStorage?

[I can modify the key on my model's before\_save callback but i'm getting \`duplicate key value violates unique constraint \\"index\_active\_storage\_blobs\_on\_key\\" \` error on logo updates \(it's obvious indeed\). ](https://preview.redd.it/816hacqr7bj51.png?width=702&amp;format=png&amp;auto=webp&amp;s=c0cdd7e7d218e3ba8d88be4089579d987ad85f51)

&amp;#x200B;

[I basically want to store my logos in s3 like this](https://preview.redd.it/gn5jlyr78bj51.png?width=414&amp;format=png&amp;auto=webp&amp;s=3563ffeecc2b3152018f22338cbf4454a8d17af6)

&amp;#x200B;

[Not like this.](https://preview.redd.it/vopx9p4h8bj51.png?width=594&amp;format=png&amp;auto=webp&amp;s=ef91199af5c6cbf7e9adbe7d411fc10bd481b0ef)

&amp;#x200B;

Is there a way to configure naming convention with active\_storage?
## [6][Rails 6 Index view loaded with limitations?](https://www.reddit.com/r/rails/comments/igih4h/rails_6_index_view_loaded_with_limitations/)
- url: https://www.reddit.com/r/rails/comments/igih4h/rails_6_index_view_loaded_with_limitations/
---
I have a `people` index that I want to load with just a specific set of records shown filtered by a specific field `(position)`. I currently have it that way as I am filtering out the position entries I do not want shown, but this is limiting me to not being able to show the other positions at all. I know my approach to this is probably very bad, as I have cobbled it together with very little knowledge.

What it boils down to is I have an index of people, with filtering buttons along the top of the page, to only show the people in the desired position. I have a group(s) of people that do not need to be listed expect when their filter button is specifically selected.

index.html.erb

    &lt;% provide(:title, "Directory") %&gt;
    
    &lt;h1&gt;Directory&lt;/h1&gt;
    
    &lt;div class="filter_div"&gt;
    	&lt;div class="filter btn btn-default btn-directory active"&gt;&lt;%= link_to "Full Listing", people_path %&gt;&lt;/div&gt;
    	&lt;div class="filter btn btn-default btn-directory active"&gt;&lt;%= link_to "Position1", people_path(:filter_by =&gt; :position1), {:method =&gt; :get} %&gt;&lt;/div&gt;
    	&lt;div class="filter btn btn-default btn-directory active"&gt;&lt;%= link_to "Position2", people_path(:filter_by =&gt; :position2), {:method =&gt; :get} %&gt;&lt;/div&gt;
    	&lt;div class="filter btn btn-default btn-directory active"&gt;&lt;%= link_to "Position3", people_path(:filter_by =&gt; :'position3'), {:method =&gt; :get} %&gt;&lt;/div&gt;
    	&lt;div class="filter btn btn-default btn-directory active"&gt;&lt;%= link_to "Position4", people_path(:filter_by =&gt; :position4), {:method =&gt; :get} %&gt;&lt;/div&gt;
    	&lt;div class="filter btn btn-default btn-directory active"&gt;&lt;%= link_to "Position5", people_path(:filter_by =&gt; :position5), {:method =&gt; :get} %&gt;&lt;/div&gt; 	
    	&lt;div class="filter btn btn-default btn-directory active"&gt;&lt;%= link_to "Position6", people_path(:filter_by =&gt; :position6), {:method =&gt; :get} %&gt;&lt;/div&gt;
    	&lt;div class="filter btn btn-default btn-directory active"&gt;&lt;%= link_to "Position7", people_path(:filter_by =&gt; :position7), {:method =&gt; :get} %&gt;&lt;/div&gt; 
    
    &lt;/div&gt;
    
    &lt;table class="people table-striped"&gt;
    	
    	&lt;tbody&gt;
    	    &lt;% @people.each do |person| %&gt;
    		&lt;% if person.position == 'position1' %&gt; 
    		    &lt;%= render 'row', person: person %&gt;            
    	        &lt;% elsif person.position == 'position2' %&gt;
                	    &lt;%= render 'row', person: person %&gt;
                    &lt;% elsif person.position == 'position3' %&gt;
                	    &lt;%= render 'row', person: person %&gt;
                    &lt;% elsif person.position == 'position4' %&gt;
                	    &lt;%= render 'row', person: person %&gt;
                    &lt;% elsif person.position == 'position5' %&gt;
                	    &lt;%= render 'row', person: person %&gt;
                    &lt;% elsif person.position == 'position6' %&gt;
                	    &lt;%= render 'row', person: person %&gt;
                    &lt;% end %&gt;  	
    	    &lt;% end %&gt;
    	&lt;/tbody&gt;
    &lt;/table&gt;
    

people\_controller

    class PeopleController &lt; ApplicationController
    	def index
    	  if params[:filter_by].present?
    	    case params[:filter_by]
    	    when 'ra_cs'
    	      @people = Person.where(ra_cs: true)
    	    when 'ra_hn'
    	      @people = Person.where(ra_hn: true)
    	    when 'ra_mg'
    	      @people = Person.where(ra_mg: true)
    	    when 'ra_nb'
    	      @people = Person.where(ra_nb: true)
    	    when 'ra_ne'
    	      @people = Person.where(ra_ne: true)
    	    when 'search'
    	      @people = Person.where(search: true)
    	    else
    	      @people = Person.where(position: params[:filter_by])
    	    end
    	  else
    	    @people = Person.all
    	  end
    	end 
    
    	def show
    		if params[:id]
    	    	@person = Person.find(params[:id])
    	    else
    	    	@person = Person.find_by(position: params[:position], uname: params[:uname])
    	    end
    	end
    
    	
    
    	private
    	def person_params
    		params.require(:person).permit(
    			:uname, :prefix, :fname, :lname, :title, :position, 
    			... all fields) 
    	end
    end

&amp;#x200B;
## [7][White Screen after create-react-app S3 Static Site Deploy](https://www.reddit.com/r/rails/comments/igj3sy/white_screen_after_createreactapp_s3_static_site/)
- url: https://www.reddit.com/r/rails/comments/igj3sy/white_screen_after_createreactapp_s3_static_site/
---
My company has a create-react-app hosted in an S3 bucket as a static site with cloudfront.

However, when we do a deploy, a lot of users experience a blank white screen. This can be remedied with "disable cache" or using another browser.

What can we do to make sure this doesn't happen for our users?

Stack is Rails/React
## [8][Help with self in refactoring observers to concerns](https://www.reddit.com/r/rails/comments/igd4an/help_with_self_in_refactoring_observers_to/)
- url: https://www.reddit.com/r/rails/comments/igd4an/help_with_self_in_refactoring_observers_to/
---
I'm moving existing callbacks in observers to model concern in rails 4.2.1 application  
Existing Observer

    class TaskObserver &lt; ActiveRecord::Observer
      observer Task
      def after_commit(model)
        //do something with the model
      end
    end 

Which I'll be moving to concern like this

    module TaskConcern
      extend ActiveSupport::Concern
      included do
        after_commit :do_something
      end
      def do_something
      // here self would be equivalent to model?
      end
    end

I'll be adding this concern in a few model's and I wanted to be sure about the usage of self in the concern here.

Since the included will be executed in the context of class it is being included in self would always be the item that is being created/updated/deleted right. i.e. the model variable in Observer and the self in concern would be the same, isn't it?

EDIT: What is the difference between ActiveRecord::Observer and ActiveSupport::Concern, because in the former self returns an instance of the Observer class but in the later self returns an instance of the model in which the concern in included.
## [9][after_commit callback not triggered from Concern](https://www.reddit.com/r/rails/comments/ig9fp6/after_commit_callback_not_triggered_from_concern/)
- url: https://www.reddit.com/r/rails/comments/ig9fp6/after_commit_callback_not_triggered_from_concern/
---
I have a model class like this

    class Task &lt; ActiveRecord::Base
      include Concerns::Tasks 
      self.table_name = "tasks" 
    end

and the concern is like in the app/models/concerns directory

    module Concerns::Tasks 
      extend ActiveSupport::Concern
      included do 
        after_commit :do_something
      end
      def do_something
        byebug
      end
    end

But my after\_commit callback isn't being hit at all. I'm on rails 4.2.1. Any ideas why?

EDIT: Other concern files are working fine. But only the newly added file isn't working fine. Do I have to register is somewhere?  


FINAL\_EDIT: The issue was in the filename. Should have known how rails picks up these files . Got it from here [https://stackoverflow.com/a/12306720/3575018](https://stackoverflow.com/a/12306720/3575018)
## [10][How to maintain CRUD when accepting nested attributes?](https://www.reddit.com/r/rails/comments/ifttz6/how_to_maintain_crud_when_accepting_nested/)
- url: https://www.reddit.com/r/rails/comments/ifttz6/how_to_maintain_crud_when_accepting_nested/
---
Let’s say these are my models:

Shift

\-- has\_many :employees

\-- has\_many :holiday\_shedules

\-- accepts\_nested\_attributes\_for :employees

\-- accepts\_nested\_attributes\_for :holiday\_schedules

Employees

\-- belongs\_to :shift

HolidaySchedule

\-- belongs\_to :shift

\-- has\_many :holidays

\-- accepts\_nested\_attributes\_for :holidays

Holiday

\-- belongs\_to HolidaySchedule

User steps:

1. Creates the shift

2. Adds staffing to the shift (multiple at a time)

3. Adds two, linked holiday\_schedules to the shift at a time, each with three holidays

Everything I’ve read says that custom controller actions are bad, but how do I maintain CRUD when accepts\_nested\_attributes\_for wants me to use a single controller update action for so many things?  That doesn’t even count updating the shift itself.

These are the options I’m seeing:

1.  Create or update multiple Nurses or HolidaySchedules through their respective controllers by namespacing and picking apart the parameters

2.  Send them all to Shift’s Update action and figure out what to do with them using conditionals

3.  Send them to custom actions in the shifts\_controller, i.e. def update\_nurses, def update\_holiday\_shifts

This is my first Rails project.  Are there options I’m not seeing?  What would be the most conventional, or Railsy, path?
## [11][Javascript not firing with Turbolinks](https://www.reddit.com/r/rails/comments/ifsu1z/javascript_not_firing_with_turbolinks/)
- url: https://www.reddit.com/r/rails/comments/ifsu1z/javascript_not_firing_with_turbolinks/
---
Hi Folks,
Using rails 5.2 with turbolinks and it's driving me crazy. Attempting to hide some items on a page when a button is clicked, but the js doesn't seem to fire. Even just trying to debug by using console.log and alert() methods doesn't seem help as neither fires when the button is clicked. Code works fine if I run it in the console, but clearly some issue with turbolinks here. Any ideas? Using an event listener on turbolinks:load (see below)

    document.addEventListener("turbolinks:load", function() {
        var btnWhiskey = document.getElementById('btn-Whiskey');
        btnWhiskey.addEventListener('click', function(){
             alert("testttt");
             console.log("TEST!");
        });
    });
## [12][Calling Database fails in secondary page.](https://www.reddit.com/r/rails/comments/ifzx3c/calling_database_fails_in_secondary_page/)
- url: https://www.reddit.com/r/rails/comments/ifzx3c/calling_database_fails_in_secondary_page/
---
I have a people database, that functions fine within it's own controller group (index, show, etc)

I want to use calls to this database in another controller, but it does not seem to be getting any data. I am sure I am missing something small, but I am hoping someone can help me.

outside\_controller:

    def pagename
        @people = Person.all
      end

pagename.html.erb

    &lt;% @people.each do |person| %&gt;
        &lt;li class="clearfix"&gt;
    	&lt;% if person.groffice == 'office' %&gt;
    		&lt;div class="image"&gt;
    		    &lt;%= image_tag("profiles/#{person.uname}S.jpg") %&gt;
    		&lt;/div&gt;
    		&lt;div class="body"&gt;
    			&lt;h5&gt;&lt;%= person.fname %&gt; &lt;%= person.lname %&gt;&lt;/h5&gt;
    		&lt;/div&gt;
    	&lt;% end %&gt;
        &lt;/li&gt;
    &lt;% end %&gt;

Person.rb

    class Person &lt; ApplicationRecord
    	has_many :pubs
    	default_scope { order('lname') }
    	belongs_to :boss, class_name: 'Person'
    	has_many :subordinates, class_name: 'Person', foreign_key: 'boss_id'
    
    	
    
        validates_presence_of :uname, :position, :fname, :lname # Needed for friendly URLs
    end

When I add the `person.groffice` call to my People Index, it pulls the data from that field and displays it without issue. I have a feeling I am missing something stupid. If I run `Person.where(:groffice =&gt; 'office')` in the console it will return the proper record.

&amp;#x200B;

Error.log has this when I attempt to render the page:

      [1m[35m (0.4ms)[0m  [1m[35mSET NAMES utf8,  @@SESSION.sql_mode = CONCAT(CONCAT(@@sql_mode, ',STRICT_ALL_TABLES'), ',NO_AUTO_VALUE_ON_ZERO'),  @@SESSION.sql_auto_is_null = 0, @@SESSION.wait_timeout = 2147483[0m
    Processing by OutsideController#pagename as HTML
      Rendering outside/pagename.html.erb within layouts/application
      [1m[36mPerson Load (1.6ms)[0m  [1m[34mSELECT `people`.* FROM `people` ORDER BY lname[0m
      ↳ app/views/outside/pagename.html.erb:17
      Rendered outside/pagename.html.erb within layouts/application (Duration: 29.8ms | Allocations: 18000)
    [Webpacker] Everything's up-to-date. Nothing to do
      Rendered layouts/_header.html.erb (Duration: 0.0ms | Allocations: 6)
      Rendered layouts/_footer.html.erb (Duration: 2.6ms | Allocations: 1433)
    Completed 200 OK in 62ms (Views: 59.8ms | ActiveRecord: 1.6ms | Allocations: 36658)

people\_controller

    class PeopleController &lt; ApplicationController
        def index
          if params[:filter_by].present?
            case params[:filter_by]
            when 'ra_cs'
              @people = Person.where(ra_cs: true)
            when 'ra_hn'
              @people = Person.where(ra_hn: true)
            when 'ra_mg'
              @people = Person.where(ra_mg: true)
            when 'ra_nb'
              @people = Person.where(ra_nb: true)
            when 'ra_ne'
              @people = Person.where(ra_ne: true)
            when 'search'
              @people = Person.where(search: true)
            else
              @people = Person.where(position: params[:filter_by])
            end
          else
            @people = Person.all
          end
        end 
    
        def show
            if params[:id]
                @person = Person.find(params[:id])
            else
                @person = Person.find_by(position: params[:position], uname: params[:uname])
            end
        end
    
        
    
        private
        def person_params
            params.require(:person).permit(
                :uname, :prefix, :fname, :lname, :title, :position, 
                ... {every field in people} ... :groffice) 
        end
    end

&amp;#x200B;
