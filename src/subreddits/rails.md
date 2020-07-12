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
## [2][Namespaced Active Record models in models/ar folder? Does anyone here do this?](https://www.reddit.com/r/rails/comments/hpomzh/namespaced_active_record_models_in_modelsar/)
- url: https://www.reddit.com/r/rails/comments/hpomzh/namespaced_active_record_models_in_modelsar/
---
I just listened to the excellent [Rails with Jason podcast with Sandi Metz](https://www.codewithjason.com/sandi-metz-author-poodr-special-guest-tj-stankus-rails-jason-podcast/). 

Sandi mentioned that when she starts a new Rails app, she always moves AR models to a "models/ar" subfolder and namespaces them.

The idea being that AR models are no longer the central concept in the app, and become dedicated to specifically database access. Leaving space for PORO business objects taking the center stage, which use the AR models when necessary. One of the outcomes being that AR are no longer the default place for functionality and thus do not grow into god objects.

This makes a lot of sense and I tried it in a test app, but am noticing a few things:

1. To properly namespace models in "models/ar", I need to name them like Ar::Post and not AR::Post. A small thing but it bugs me. What's a way to fix this?
2. Some Rails conventions no longer work, like `link_to 'Show', post` must become `link_to 'Show', post_path(post)` because the rails magic gets confused by the absence of the default AR name.

Luckily, associations seem to work OK, so `user.posts` still works out of the box.

**Does anyone here use this kind of conventions and what kinds of other adjustments did you need to make?**


Edit: The answer to 1. is adding an inflector rule in config/inflections.rb


    ActiveSupport::Inflector.inflections(:en) do |inflect|
      inflect.acronym 'AR'
    end
## [3][Rails creating development database with different name?](https://www.reddit.com/r/rails/comments/hpsv0n/rails_creating_development_database_with/)
- url: https://www.reddit.com/r/rails/comments/hpsv0n/rails_creating_development_database_with/
---
My database.yml file says this:

    development:
      &lt;&lt;: *default
      database: project_name_development
    
    test:
      &lt;&lt;: *default
      database: project_name_test

However, when I run bundle exec rails db:create, the databases created are:

    Created database '123'
    Created database 'project_name_test'

I'm kind of new to Rails so I don't quite know what is happening here. Isn't the '123' database supposed to be named 'project\_name\_development'? I made no changes after running 'rails new' command, so why is it different? Does it matter if it is different or can I just move on?
## [4][Help with running RSpec with spring in rails 4.2.3 application](https://www.reddit.com/r/rails/comments/hprklk/help_with_running_rspec_with_spring_in_rails_423/)
- url: https://www.reddit.com/r/rails/comments/hprklk/help_with_running_rspec_with_spring_in_rails_423/
---
I wanted to run the RSpec via the spring preloader so I added this and this after bundle install i can see I'm using these two gems

    Using spring 2.0.2
    Using spring-commands-rspec 1.0.4

I then generate the binstub 

    bundle exec spring binstub --all

and when I do \`bin/rspec path\_to\_spec.rb\` I get this error  
 

    activesupport-4.2.11.3/lib/active_support/concern.rb:126:in `included': Cannot define multiple 'included' blocks for a Concern (ActiveSupport::Concern::MultipleIncludedBlocks)

When I try it without the bin/rspec and do just rspec it works fine. Can someone help me with this?
## [5][How to use find_by_username with case_sensitive: false?](https://www.reddit.com/r/rails/comments/hprbfc/how_to_use_find_by_username_with_case_sensitive/)
- url: https://www.reddit.com/r/rails/comments/hprbfc/how_to_use_find_by_username_with_case_sensitive/
---
How to use `find_by_username` with `case_sensitive: false`?

I explain my issue. I have this script to take the username tagged (after @) and to check if it is a "existing user" or not.

        text.gsub! (/@(\S+)/) do |match|
          user = User.find_by_username($1)
            if
            else
            end
          end

In this way if the text is `@Mike` and the user signed is `Mike`, it is ok.

BUT if the text is `@mike` and the user signed is `Mike`, it doesn't work.

**Hot to add a case\_sensitive false in find\_by\_username?**
## [6][[Help] remote: true not sending an AJAX request even though data-remote-"true" is there](https://www.reddit.com/r/rails/comments/hpm2b6/help_remote_true_not_sending_an_ajax_request_even/)
- url: https://www.reddit.com/r/rails/comments/hpm2b6/help_remote_true_not_sending_an_ajax_request_even/
---
Hey everyone, I've been having some trouble recently with a course I've been following. 

I'm trying to render a partial through a respond_to. I have been following the course strictly, and I'm very confused about why my project is now showing an error while the instructor's project up to this point is running fine.

*app/controllers/stocks_controller.rb*

    &lt;h1&gt;My Portfolio&lt;/h1&gt;
    &lt;div class='search-area'&gt;
      &lt;h3&gt;Search Stocks&lt;/h3&gt;
      &lt;%= form_tag search_stock_path, remote: true, method: :get do %&gt;
        &lt;div class="form-group row"&gt;
          &lt;div class="col-sm-9 no-right-padding"&gt;
            &lt;%= text_field_tag :stock, params[:stock], placeholder: "Stock ticker symbol", autofocus: true, class: "form-control form-control-lg" %&gt;
          &lt;/div&gt;
          &lt;div class="col-sm-3 no-left-padding"&gt;
            &lt;%= button_tag type: :submit, class: "btn btn-success" do %&gt;
              &lt;%= fa_icon 'search 2x' %&gt;
            &lt;% end %&gt;
          &lt;/div&gt;
        &lt;/div&gt;
      &lt;% end %&gt;
    &lt;/div&gt;
 
    &lt;div id="results"&gt;
     
    &lt;/div&gt;

It seems like the **remote: true** part of this is not working at all. When I follow the instructor's commands before the section this question pertains to, no ajax is called, and the page reloads again.

What I mean is that the page is loading the search_stock endpoint from the /my_portfolio endpoint (where I start the search from) instead of staying on /my_portfolio and waiting for a js response. The form tag does correctly have the data-remote="true" show up in the inspector though.

[Here is an image of what I'm describing](https://i.stack.imgur.com/olQDO.png)

The goal is to get the controller to respond to js and return a partial js.erb file

*app/controllers/stocks_controller.rb*

    class StocksController &lt;ApplicationController
        
        def search
            if params[:stock].present?            
                 @stock = Stock.new_lookup(params[:stock])
                 if @stock            
                    respond_to do |format|
                        format.js { render partial: 'users/result' }
                    end         
                 else
                    flash[:alert] = "Please enter a valid symbol to search"
                    redirect_to my_portfolio_path
                 end             
            else 
                flash[:alert] = "Please enter a symbol to search"
               redirect_to my_portfolio_path
            end
        end
    
    end

There is both a _result.html.erb and _result.js.erb in my app/views/users folder. _result.js.erb simply contains an alert to make sure the ajax call is working. I already have rails-ujs in my application.js file, and I have not modified anything related to rails-ujs myself.

    // This file is automatically compiled by Webpack, along with any other files
    // present in this directory. You're encouraged to place your actual application logic in
    // a relevant structure within app/javascript and only use these pack files to reference
    // that code so it'll be compiled.
    
    require("@rails/ujs").start()
    require("turbolinks").start()
    require("@rails/activestorage").start()
    require("channels")
    
    // Uncomment to copy all static images under ../images to the output folder and reference
    // them with the image_pack_tag helper in views (e.g &lt;%= image_pack_tag 'rails.png' %&gt;)
    // or the `imagePath` JavaScript helper below.
    //
    // const images = require.context('../images', true)
    // const imagePath = (name) =&gt; images(name, true)
    
    import "bootstrap"


I am using Rails 6.0.3.2, so I don't think there should be an issue I hope.
## [7][What route should I use?](https://www.reddit.com/r/rails/comments/hpj5we/what_route_should_i_use/)
- url: https://www.reddit.com/r/rails/comments/hpj5we/what_route_should_i_use/
---
I have a Category model and a WordlistEntry model. They have a has_and_belongs_to_many relationship between them.

I'm trying to decide what request to make from the frontend when the user adds some Categories to a WordlistEntry.

The Categories may or may not already exist in the database.

I'm thinking it should be PATCH /wordlist_entries/:id/categories

But I'm really not sure on this. What do you think?
## [8][Can't get test to pass without 'reload' in after_create instance method.](https://www.reddit.com/r/rails/comments/hpe3a6/cant_get_test_to_pass_without_reload_in_after/)
- url: https://www.reddit.com/r/rails/comments/hpe3a6/cant_get_test_to_pass_without_reload_in_after/
---
Hey I'm having an issue with a method, which, when I test it, requires a `reload` in the model in order to pass my RSpec test.  


Here is my test, which is essentially checks to see if the venue correctly updates it's own boolean value (in this case "serves\_food") if the truthy/falsy declaration for the two most recent reviews for that venue are the same.  


    describe 'truthy/falsy venue values if the last two review values are the same' do
      let!(:rev1) { create(:review, user: u, venue: ven, serves_food: true) }
      let!(:rev2) { create(:review, user: u, venue: ven, serves_food: false) }
      let!(:rev3) { create(:review, user: u, venue: ven, serves_food: false) }
      it 'should update' do
        expect(ven.serves_food).to be false
      end
    end

  
Here is the method which handles that logic (which is on the venue model, but is triggered in an "after\_update" action on the review model):  


    def truth_checked_val(attribute, value)
      if self.reviews.count == 0
        nil
      # executes if there is only one review, or the last review value matches current
      elsif self.reviews.count == 1 || self.reviews[-2].send(attribute) == value
        value
      else
        # returns the previous value if it does not match the current
        self.reviews[-2].send(attribute)
      end
    end

  
Now, I know why the test fails. It is because, for some reason, when the third `Review` factory is created, when I put a `byebug` at the beginning of the method, even though when I test it with `self.review``s.count` I get the correct number returned (3), if I try to return those reviews with `self.reviews`, only 2 of them (the first two) are returned.  


So the only thing I have found which will fix this is to put a `.reload` on the first line of the method like so:

    def truth_checked_val(attribute, value)
      self.reviews.reload
      ...
    end

This is a solution and the test will pass, but it seems very hacky. However I have as yet found no other way to get it to work.  


TLDR; My test is failing because my conditional in my instance method is returning FALSE for the line `self.reviews[-2].send(attribute) == value` because it is returning the third to last review for `self.reviews[-2]` rather than the second to last. Only a `.reload` in the method is working, which is hacky but I can't find another way to resolve it.  


Help appreciated.
## [9][Building a blockchain app without using Etheream](https://www.reddit.com/r/rails/comments/hpfw4c/building_a_blockchain_app_without_using_etheream/)
- url: https://www.reddit.com/r/rails/comments/hpfw4c/building_a_blockchain_app_without_using_etheream/
---
I am a Rails developer and I have been reading a lot about Blockchain a lot since last year.

I saw a video of how a transaction works in a chain manner in Blockchain, based on a sha256 byte hash string.  Let us say I wanted to build a "Funding tracker" app. Do I have to learn Solidity to do that? Is it possible to build a blockchain concept app without using Solidity and instead using other backend languages such as ruby or python?
## [10][best practices on creating users with Devise?](https://www.reddit.com/r/rails/comments/hpbhjd/best_practices_on_creating_users_with_devise/)
- url: https://www.reddit.com/r/rails/comments/hpbhjd/best_practices_on_creating_users_with_devise/
---
\[I'm a newb\]. going through Hartl's tutorial, i decided to try to use Devise Gem instead of manually creating users/authentication.

Is the best practice to use Devise and then manually generate a migration to add a column to users and then update code in the application controller?
## [11][Retrieve the most liked article](https://www.reddit.com/r/rails/comments/hp8npu/retrieve_the_most_liked_article/)
- url: https://www.reddit.com/r/rails/comments/hp8npu/retrieve_the_most_liked_article/
---
So I have add a Like feature on my blog app and now I'm trying to retrieve the most liked article to display on my home page.

I have been following this documentation on active record calculations [here](https://api.rubyonrails.org/classes/ActiveRecord/Calculations.html).

I have a Like model which belongs\_to an article and belongs\_to an user.

I don't know if this is the issue but I tried to retrieve the most liked article by doing Article.maximum(:likes) but it says the column "likes" does not exist. And yeah its perfectly fine since the article I followed up does not implemented that way.

The thing is I think I should check the most liked article by using my Like model  and check who is the article\_id that most appears. I couldn't grasp yet how to do that. It always results in error whatever I try.

My [repo](https://github.com/Gregory280/alpha-blog-5.1.4)

By the way, should I put this on my view or on a "popular" method in likes controller?
