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
## [3][Why is this migration giving a 'no such table' error?](https://www.reddit.com/r/rails/comments/hs4rgg/why_is_this_migration_giving_a_no_such_table_error/)
- url: https://www.reddit.com/r/rails/comments/hs4rgg/why_is_this_migration_giving_a_no_such_table_error/
---
I have a table named object\_classes with a column named ClassList\_id. From schema.rb:

      create_table "object_classes", force: :cascade do |t|
        t.string "name"
        t.integer "ClassList_id", null: false
        t.datetime "created_at", precision: 6, null: false
        t.datetime "updated_at", precision: 6, null: false
        t.index ["ClassList_id"], name: "index_object_classes_on_ClassList_id"
      end

I've realized that the column should be named class\_list\_id to conform with Rails expected naming convention. Therefore, I have generated a new migration:

    class FixColumnName &lt; ActiveRecord::Migration[6.0]
      def change
        rename_column :object_classes, :ClassList_id, :class_list_id
      end
    end

However, when I run this migration, I get the following error:

    /bin/bash -c "env RBENV_VERSION=2.6.1 /home/asfarley/.rbenv/libexec/rbenv exec bundle exec ruby /home/asfarley/imgseq/bin/spring rails 'db:migrate'"
    == 20200716060501 FixColumnName: migrating ====================================
    -- rename_column(:object_classes, :ClassList_id, :class_list_id)
    rails aborted!
    StandardError: An error has occurred, this and all later migrations canceled:
    
    SQLite3::SQLException: no such table: main.ClassLists
    /home/asfarley/imgseq/db/migrate/20200716060501_fix_column_name.rb:3:in `change'
    /home/asfarley/imgseq/bin/rails:9:in `&lt;top (required)&gt;'
    /home/asfarley/imgseq/bin/spring:15:in `require'
    /home/asfarley/imgseq/bin/spring:15:in `&lt;main&gt;'
    
    Caused by:
    ActiveRecord::StatementInvalid: SQLite3::SQLException: no such table: main.ClassLists
    /home/asfarley/imgseq/db/migrate/20200716060501_fix_column_name.rb:3:in `change'
    /home/asfarley/imgseq/bin/rails:9:in `&lt;top (required)&gt;'
    /home/asfarley/imgseq/bin/spring:15:in `require'
    /home/asfarley/imgseq/bin/spring:15:in `&lt;main&gt;'
    
    Caused by:
    SQLite3::SQLException: no such table: main.ClassLists
    /home/asfarley/imgseq/db/migrate/20200716060501_fix_column_name.rb:3:in `change'
    /home/asfarley/imgseq/bin/rails:9:in `&lt;top (required)&gt;'
    /home/asfarley/imgseq/bin/spring:15:in `require'
    /home/asfarley/imgseq/bin/spring:15:in `&lt;main&gt;'
    Tasks: TOP =&gt; db:migrate
    (See full trace by running task with --trace)
    
    Process finished with exit code 1
    

What am I doing wrong here?
## [4][Is RailsCasts.com still relevant for learning Rails?](https://www.reddit.com/r/rails/comments/hrr3y3/is_railscastscom_still_relevant_for_learning_rails/)
- url: https://www.reddit.com/r/rails/comments/hrr3y3/is_railscastscom_still_relevant_for_learning_rails/
---
Seeing as how the videos have not been updated for years, can this site still be used for learning, or have the Rails APIs changed too much since 2012-2013?
## [5][Resources](https://www.reddit.com/r/rails/comments/hs2kv7/resources/)
- url: https://www.reddit.com/r/rails/comments/hs2kv7/resources/
---
Just dropping this repo again. Would love to get a bunch of contributors to add. 

Cheers!

Heres the repo

[https://github.com/tylertomlinson/crucial\_resources](https://github.com/tylertomlinson/crucial_resources)
## [6][Chill full project videos](https://www.reddit.com/r/rails/comments/hs2g9n/chill_full_project_videos/)
- url: https://www.reddit.com/r/rails/comments/hs2g9n/chill_full_project_videos/
---
I'm looking for chill videos of people doing rails projects from start to finish (it could be a long series of videos), preferably by someone with a lot of experience, so I don't pick up bad habits. I'm not looking for tutorials exactly. Just something I can de-stress to after a long day of work. 
Does anyone have any suggestions?
## [7][Is there a more beautiful way to declare and populate an array from ActiveRecord data?](https://www.reddit.com/r/rails/comments/hs3e6g/is_there_a_more_beautiful_way_to_declare_and/)
- url: https://www.reddit.com/r/rails/comments/hs3e6g/is_there_a_more_beautiful_way_to_declare_and/
---
Still learning Ruby and Rails. Ruby is so elegant! However, it pains me that I have to declare an array and populate it from ActiveRecord on two lines. Surely, Ruby has a more majestic way to do this on a single line?

```
underlyings = []
Underlying.all.each { |underlying| underlyings &lt;&lt; underlying.symbol }
```
## [8][Rails 6 and Active Job : a simple tutorial](https://www.reddit.com/r/rails/comments/hrnisf/rails_6_and_active_job_a_simple_tutorial/)
- url: https://www.reddit.com/r/rails/comments/hrnisf/rails_6_and_active_job_a_simple_tutorial/
---
Active Job is quite difficult to start with. The documentation is very good, but to just to start an "hello world" job example, it requires redis, sidekiq (or another implementation), and some configuration. I tried to create a small tutorial in which 1) you have everything that just works in a couple of minutes and 2) everything works in complete isolation of your environment. [http://bdavidxyz.com/blog/rails-6-activejob-tutorial/](http://bdavidxyz.com/blog/rails-6-activejob-tutorial/)
## [9][General advise on creating and updating a model instance](https://www.reddit.com/r/rails/comments/hrqyka/general_advise_on_creating_and_updating_a_model/)
- url: https://www.reddit.com/r/rails/comments/hrqyka/general_advise_on_creating_and_updating_a_model/
---
Hi. I'm a beginner at web programming. **I want to make the user navigate different questions, each updating an instance of the model Item.** 

My plan was to create two links that create this instance, one adding the value "found" and the other the value "lost" to a "Category" parameter. I also wanted this link to redirect to an update form, which will be a search box and a submit button that will edit the parameter "Name" value.

So far, what I could do is, in my home view:

&amp;#x200B;

&gt;&lt;button type="button" class="btn btn-info mb-4 "&gt;&lt;%= link\_to "I lost something", new\_item\_path(category: "lost")%&gt; &lt;/button&gt;

 Then on my Item controller,

&amp;#x200B;

&gt;def new  
 item = Item.new  
 item.category = params\[:category\]  
 item.user = current\_user  
 item.save  
 end

This allowed me to create an instance and get the "lost" value for the :category parameter. 

This is my .new view, where I was thinking to put a search box and a submit button. Right now there is a form.

&gt;&lt;%= simple\_form\_for(item) do |f| %&gt;  
&lt;%= f.input :name %&gt;  
&lt;%= f.submit %&gt;  
&lt;% end %&gt;

Now, I want to update the same instance with that form and I got completely lost. How do I make this submit button update the instance? Should I call an update action from my ItemsController with the SimpleForm? Anyone knows how to do that?

I hope my question is clear. Maybe there is a more convenient way to do this. 

Any help is appreciated, with our without code
## [10][How to transform my app layout to the desired idea?](https://www.reddit.com/r/rails/comments/hrt03n/how_to_transform_my_app_layout_to_the_desired_idea/)
- url: https://www.reddit.com/r/rails/comments/hrt03n/how_to_transform_my_app_layout_to_the_desired_idea/
---
Hello, I'm developing an app (I'm more familiar with the backend side) and its current state is like this: [https://postimg.cc/cKLBhSnt](https://postimg.cc/cKLBhSnt) but, my superior edited the picture in PS and it ideally should look like this: [https://postimg.cc/T56ykV4N](https://postimg.cc/T56ykV4N) how I can fix the buttons to the desired positions? What is this, a templating pattern? here is the code: [https://github.com/LeoFragozo/sistema\_loja/blob/master/app/views/products/index.html.slim](https://github.com/LeoFragozo/sistema_loja/blob/master/app/views/products/index.html.slim)
## [11][OMG! I just got my first user on my rails app](https://www.reddit.com/r/rails/comments/hr6ms0/omg_i_just_got_my_first_user_on_my_rails_app/)
- url: https://www.reddit.com/r/rails/comments/hr6ms0/omg_i_just_got_my_first_user_on_my_rails_app/
---
Doing webdev for almost a 2 years, started with JS, drove me crazy. Just wanted to learn to build something without doing to much my self. Found out about rails, started building some simple stuff, like a blog. After that build a app for my brother, and wrote some SEO about it.... and finally got my first user today! even if it is just a free user (have a free restricted plan and a premium plan) it is a great feeling!
## [12][TDD, I can't create post model because some association are empty](https://www.reddit.com/r/rails/comments/hrpcp2/tdd_i_cant_create_post_model_because_some/)
- url: https://www.reddit.com/r/rails/comments/hrpcp2/tdd_i_cant_create_post_model_because_some/
---
Hi everyone, I'm trying to use rails in api mode with TDD behavior.

I use rspec, factoryBot , Faker for my test. So here is my code

    #post_request_spec
    require 'rails_helper'
    
    RSpec.describe "Posts", type: :request do
        #initialize test data
        
        let!(:user) { FactoryBot.create(:user)}
        let!(:category) { FactoryBot.create(:category)}
        let!(:posts) { create_list(:post, 10) }
        let(:post_id) { post.first.id }
        
        #Test suite for Get api/v0/posts
        describe 'GET/api/v0/posts' do
            #make HTTP get request before each example
            before { get '/posts' }
    
            it 'returns posts' do
                # Note 'json' is a custom helper to parse JSON response
                expect(json).not_to be_empty
            end
    
            it 'returns status code 200' do
                expect(response).to have_http_status(200)
            end
        end
    ...

there are a lot more test but they've got the same problem:

    Failures:
    
      1) Posts GET/api/v0/posts returns posts
         Failure/Error: let!(:posts) { create_list(:post, 10) }
         
         ActiveRecord::RecordInvalid:
           Validation failed: Category can't be blank

I try to create Category before Post but it look like my Category is empty even if i create it with FactoryBot.

    2.6.0 :003 &gt; Category
     =&gt; Category(id: integer, title: string, created_at: datetime, updated_at: datetime)
     
    #app/models/category.rb
    class Category &lt; ApplicationRecord
        #model association
        has_many :posts, dependent: :destroy
        
        #Validations
        validates_presence_of :title
    end
    
    #spec/factories/category.rb
    FactoryBot.define do
        factory :category do
          title { Faker::Games::SonicTheHedgehog.character }
        end
      end

I also have a FactoryBot for my Post

    2.6.0 :002 &gt; Post
     =&gt; Post(id: integer, title: string, content: text, created_by: string, entry: integer, category_id: integer, rdv: datetime, tag1: string, tag2: string, tag3: string, created_at: datetime, updated_at: datetime, user_id: integer) 
     
    #app/models/post.rb
    class Post &lt; ApplicationRecord
      #model association
      belongs_to :category
      belongs_to :user
    
      #Validations
      validates_presence_of :title, :content, :created_by, :entry, :category_id, :rdv, :tag1
      
    end
    
    #spec/factories/posts.rb
    FactoryBot.define do
        factory :post do
          title { Faker::Lorem.word }
          content{ Faker::Lorem.paragraph }
          created_by "User"
          entry 5
          association :category, strategy: :build #builds user factory for test but doesn't save to test database
          category_id {category.id} #foreign key 
          rdv {Faker::Time.between(from: DateTime.now + 1, to: DateTime.now + 4, format: :default)}
          tag1 "tag1"
          tag2 nil
          tag3 nil
          association :user, strategy: :build #builds user factory for test but doesn't save to test database
          user_id {user.id} #foreign key 
        end
      end

So if you have any tips for learning TDD don't hesitate thank u.
