# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/jfcv1r/personal_projects_show_off_your_own_project_andor/)
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/jnwqje/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/jnwqje/personal_projects_show_off_your_own_project_andor/
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
## [3][Golang crash course for RoR developer?](https://www.reddit.com/r/rails/comments/jp07ro/golang_crash_course_for_ror_developer/)
- url: https://www.reddit.com/r/rails/comments/jp07ro/golang_crash_course_for_ror_developer/
---
In my current job I'm building and maintaining a Ruby on Rails API. But I'm expecting to start a new job in a few weeks and they're using Go. I think the syntax shouldn't be much of a problem, but I should understand the basic concept of API framework using Go. Is there any source to learn the framework quickly for a Rails developer?
## [4][Open source tool for deploying Rails on Kubernetes](https://www.reddit.com/r/rails/comments/joswbw/open_source_tool_for_deploying_rails_on_kubernetes/)
- url: https://www.reddit.com/r/rails/comments/joswbw/open_source_tool_for_deploying_rails_on_kubernetes/
---
I'm thinking about creating some open source tools that makes it easy to deploy and manage rails apps on Kubernetes and wanted to get some thoughts from the community. The tools will be mainly geared towards people who want to use Kubernetes but are a bit intimidated by the steep learning curve. The tools abstract away the details of k8s so that you can get started easily - at a high level, the goal is to build tools that give you the "Heroku experience" on Kubernetes without the exorbitant cost of Heroku. It can also be used by people who are already using Kubernetes but want to save time/want a smoother experience.

I started building a few of these tools for myself while using kubernetes, and I am just trying to gauge if there's enough interest from the community for me to start an open source project. 

What do you guys think?
## [5][[Help] Validate belong_to Uniqueness](https://www.reddit.com/r/rails/comments/josqhu/help_validate_belong_to_uniqueness/)
- url: https://www.reddit.com/r/rails/comments/josqhu/help_validate_belong_to_uniqueness/
---
I have a `tools` table which has 2 references (and 2 indexes).

    create_table "tools", force: :cascade do |t|
        t.string "name", null: false
        t.bigint "product_id", null: false
        t.bigint "manufacturer_id", null: false
        t.datetime "created_at", precision: 6, null: false
        t.datetime "updated_at", precision: 6, null: false
        t.index ["product_id", "manufacturer_id"], name: "index_tools_on_product_id_and_manufacturer_id", unique: true
        t.index ["product_id"], name: "index_tools_on_product_id"
        t.index ["manufacturer_id"], name: "index_tools_on_manufacturer_id"
      end

And the `Tool` model:

    class Tool &lt; ApplicationRecord
      belongs_to :product
      belongs_to :manufacturer
    
      validates :product_id, uniqueness: { scope: :manufacturer_id }
    end

I'm trying to make sure that there is no tool with the same product and manufacturer. Even if I have this uniqueness validator, I get the PG Error:

    ActiveRecord::RecordNotUnique (PG::UniqueViolation: ERROR:  duplicate key value violates unique constraint "index_tools_on_product_id_and_manufacturer_id"
    DETAIL:  Key (product_id, manufacturer_id)=(50, 28) already exists.
    ):

How can I get the validation to work? What am I doing wrong?

Thanks
## [6][Performance impact of rails view partials](https://www.reddit.com/r/rails/comments/jomrzp/performance_impact_of_rails_view_partials/)
- url: https://www.reddit.com/r/rails/comments/jomrzp/performance_impact_of_rails_view_partials/
---
https://scoutapm.com/blog/performance-impact-of-using-ruby-on-rails-view-partials

Came across this article today. I’m a little dumbfounded by the results. Anyone care to explain when to use and not use rails partials?
## [7][Viewing/highlighting grids in the browser (including flexbox)?](https://www.reddit.com/r/rails/comments/joizh3/viewinghighlighting_grids_in_the_browser/)
- url: https://www.reddit.com/r/rails/comments/joizh3/viewinghighlighting_grids_in_the_browser/
---
I tried to ask this question on r/webdev but I'm too new to Reddit so it wasn't approved. I thought I'd ask here since my project is a Rails project and some of you might have experience with this.

I'm looking for a way to view/highlight grids in the browser. I found a Chrome extension but it only appears to work with CSS grid, and the code is using flexbox.

Ideally I could activate the extension/plugin and the entire grid(s) for the page would be highlighted. That way I can see at a glance if all the elements are aligning on the grid per the design (for which I can see the whole grid like this in my design tool).

Does anyone know of a solution for this?
## [8][How to dynamically style button background with DB color](https://www.reddit.com/r/rails/comments/jolkae/how_to_dynamically_style_button_background_with/)
- url: https://www.reddit.com/r/rails/comments/jolkae/how_to_dynamically_style_button_background_with/
---
Say I have a link\_to with btn-primary class. what I want to do is color the background different based on the hex number stored in the DB. I am trying something like this but keep getting syntax errors. Any thoughts on what is wrong?

          &lt;%= link_to compitem, class: 'btn btn-sm', style: "background-color: &lt;%= compitem.color %&gt;" do %&gt;
            Start: &lt;%= compitem.name %&gt;
          &lt;% end %&gt;
## [9][Has Pundit been abandoned?](https://www.reddit.com/r/rails/comments/jon842/has_pundit_been_abandoned/)
- url: https://www.reddit.com/r/rails/comments/jon842/has_pundit_been_abandoned/
---
The last release is from 11 months ago.

It's still worth using it? Or should I use cancancan?
## [10][Quick question about how to create a non-RESTful route](https://www.reddit.com/r/rails/comments/joe6dg/quick_question_about_how_to_create_a_nonrestful/)
- url: https://www.reddit.com/r/rails/comments/joe6dg/quick_question_about_how_to_create_a_nonrestful/
---
I want to be able to use a URL like `/goals/:date`.  

My route currently looks like this:  

      resources :goals do
        get ':date', to: 'goals#filter_by_date'
      end


But that makes the URL `/goals/:goal_id/:date`. How do I remove `:goal_id`?
## [11][undefined method `auto_strip_attributes` in rails console?](https://www.reddit.com/r/rails/comments/jo67j8/undefined_method_auto_strip_attributes_in_rails/)
- url: https://www.reddit.com/r/rails/comments/jo67j8/undefined_method_auto_strip_attributes_in_rails/
---
I'm using this gem [https://rubygems.org/gems/auto\_strip\_attributes/versions/2.0.6](https://rubygems.org/gems/auto_strip_attributes/versions/2.0.6) for stripping white space before saving the field to the db. I'm using it in my model like:

    class Brand &lt; ApplicationRecord
      auto_strip_attributes :name
    end

No idea what I did, but now whenever I go into rails console, I'm getting:

    Traceback (most recent call last):
            3: from (irb):1
            2: from app/models/person.rb:1:in `&lt;top (required)&gt;'
            1: from app/models/person.rb:12:in `&lt;class:Brand&gt;'
    NoMethodError (undefined method `auto_strip_attributes' for Person (call 'Person.connection' to establish a connection):Class)

Version of gem I'm using:

    $ bundle list | grep auto_strip
      * auto_strip_attributes (2.6.0)

I tried:

\- To force bundle install: `$ bundle install --redownload`

\- stop any servers and only use rails console
## [12][Big open source API project built with Rails to learn good patterns from?](https://www.reddit.com/r/rails/comments/jnso6v/big_open_source_api_project_built_with_rails_to/)
- url: https://www.reddit.com/r/rails/comments/jnso6v/big_open_source_api_project_built_with_rails_to/
---
Hi guys, I recently picked rails for my next side project since it offers everything I need out of the box :)

I need to develop a rest api backend, so I created an app with the "only api" flag. I'm following some tutorials on GoRails but I'm looking for a good open source project to learn from.

Since I'm building an API, I'm looking for a project that includes a API.

Thanks :)
