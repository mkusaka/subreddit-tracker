# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gek2rk/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gek2rk/personal_projects_show_off_your_own_project_andor/
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
## [2][How to Fix Slow Code in Ruby](https://www.reddit.com/r/rails/comments/gjwyl9/how_to_fix_slow_code_in_ruby/)
- url: https://engineering.shopify.com/blogs/engineering/how-fix-slow-code-ruby
---

## [3][Design Rails JSON API with performance in mind](https://www.reddit.com/r/rails/comments/gk5yzg/design_rails_json_api_with_performance_in_mind/)
- url: https://www.reddit.com/r/rails/comments/gk5yzg/design_rails_json_api_with_performance_in_mind/
---
[https://jtway.co/design-rails-json-api-with-performance-in-mind-427e0f0e6f04](https://jtway.co/design-rails-json-api-with-performance-in-mind-427e0f0e6f04)
## [4][301 redirect](https://www.reddit.com/r/rails/comments/gk6soo/301_redirect/)
- url: https://www.reddit.com/r/rails/comments/gk6soo/301_redirect/
---
I'm using Nginx and I would like to do a 301 redirect to one domain as I have 4. I'm using the word example as my domain example in this cry for help 😥. 

I don't understand why it is not going to the root file path when I use 301 redirect, can someone help please? 

Inside the server block I have this:

server_name www.example.com example.com www.example.co.uk;

return 301 example.co.uk;

root /var/www/example;
## [5][Generating and storing reports](https://www.reddit.com/r/rails/comments/gk933x/generating_and_storing_reports/)
- url: https://www.reddit.com/r/rails/comments/gk933x/generating_and_storing_reports/
---
Let's say you have multiple models with different kind of data and you want to create daily, monthly or weekly reports for a data range.  


The model tables can have millions of rows so the report should be saved for quick access and/or download.  


I'm wondering if you have some tips or design patterns on how to implement this.  I need to save the reports somehow in the database, but not sure what's the best way to handle this. What's sure is that the result is not instant and the query might take a while to compute.
## [6][Creating a Quiz in Rails](https://www.reddit.com/r/rails/comments/gk5aiv/creating_a_quiz_in_rails/)
- url: https://www.reddit.com/r/rails/comments/gk5aiv/creating_a_quiz_in_rails/
---
Hey everyone! 

I am currently trying to build a little mental math trainer in rails, that will throw random math questions (e.g. add two numbers, subtract two numbers, multiply, square etc...) at me. My main motivation behind this is that for each question I want to record the number of try's and time taken to get it correct, so that I can alter analyse these to see my improvement (hopefully). 

Right now I am terribly stuck at implementing this though. I have created a Question model that contains a try and time taken field. My idea was that I would create a random question on the home page, display a simple\_form\_for underneath and use its input to check whether the question is correct. in the Questions create method.

My simple form for looks like this:

    &lt;% @question = Question.new() %&gt;
    
    &lt;% num1, num2, answ = @question.add_two_numbers %&gt;
    &lt;h3&gt;&lt;%= num1 %&gt; + &lt;%= num2 %&gt;&lt;/h3&gt;
    
    &lt;p&gt;Trys: &lt;%= @question.trys %&gt;&lt;/p&gt;
    
    &lt;%= simple_form_for @question do |f| %&gt;
      &lt;%= f.hidden_field :trys, :value =&gt; @question.trys || 1 %&gt;
      &lt;%= f.hidden_field :answ, :value =&gt; answ %&gt;
      &lt;%= f.input :user_answ, input_html: {value: ''}  %&gt;
      &lt;%= f.button :submit %&gt;
    &lt;% end %&gt;

My Questions create method:

    def create
        @question = Question.new(trys: question_params["trys"], time: question_params["time"])
        @question.user_id = current_user.id
        respond_to do |format|
          if @question.save &amp;&amp; (question_params["answ"] == question_params["user_answ"])
            format.html { render :new, notice: 'Correct!' }
            format.json { render :show, status: :created, location: @question }
          else
            format.html { render :new, notice: 'Thats wrong' }
            format.json { render json: @question.errors, status: :unprocessable_entity }
          end
        end
      end

Whenever I hit the point where question\_params\["answ"\] == question\_params\["user\_answ"\], meaning that the user gave a wrong answer, I would like to increment the try's count by one and redirect him to the same question. As of right now, I can not get this to work however.

I would be super glad I somebody could point me in the right direction, as I would really love to get this to work :)
## [7][Serializing/Deserializing Data between Model and Various JSON Responses](https://www.reddit.com/r/rails/comments/gk0mzk/serializingdeserializing_data_between_model_and/)
- url: https://www.reddit.com/r/rails/comments/gk0mzk/serializingdeserializing_data_between_model_and/
---
Suppose, we have multiple APIs that returns a JSON response. Both APIs correspond to the same ActiveRecord model(s).


For example, given the
JSON response A:

    { 
       'model': 'Volvo',
       'wheels': 4,
       'drivers': ['John', 'Mary', 'Kelly'],
       ....
    }


and the JSON response B:

    { 
       'producer': 'Volvo',
       'size': 4,
       'names': ['John', 'Mary', 'Kelly'],
       ....
    }



and lastly, the models: `Car(id, make, wheels), Driver(id, name), CarDriver(car_id, driver_id)`

---


What would be the most idiomatic-Railsy way to go about implementing serialization and deserialization between the Car model and the various JSON responses? In my real use-case, these responses are not trivial and have significant nesting.


Currently, I'm using POROs with a format similar to DB migrations: `to_json`, `from_json`. I don't really want to put this logic in the model because these JSON responses are dependent on various services and have no impact on my internal representation. 

Thank you!
## [8][Application Search Feature more that ActiveRecord](https://www.reddit.com/r/rails/comments/gjxym1/application_search_feature_more_that_activerecord/)
- url: https://www.reddit.com/r/rails/comments/gjxym1/application_search_feature_more_that_activerecord/
---
Is there a way to add in a search feature to an app that will not only search your ActiveRecords, but also your static pages? My higher ups want to be able to search everything and they do not understand the significant differences between the two, nor do they care. I am working at creating this in Rails 6 with Bootstrap 4. Any pointers or gem references would be greatly appreciated.
## [9][React /Rails API dependent destroy?](https://www.reddit.com/r/rails/comments/gjwxty/react_rails_api_dependent_destroy/)
- url: https://www.reddit.com/r/rails/comments/gjwxty/react_rails_api_dependent_destroy/
---
Hey, first time posting and trying to wrap my head around some stuff.  


Working with:   
Rails 6.0.3  
React 16.13.1  
Postgresql 12

Created a couple of models and controllers with rails g scaffold api/{name of thing goes here}  


in this case I have a Session and its child is Attendees. Now when I delete a session the regular "  


has\_many :attendees, dependent: :destroy or  
has\_many :api\_attendees, dependent: :destroy  
doesn't work  


Any thoughts on directions to go or maybe a good article to read on the subject? Haven't been able to find anything.  


I guess I'm trying to figure out how rails associations work(or the specific syntax) when setting up as API only and through scaffolds.
## [10][Database-driven authorization in Rails using CanCanCan - Abilities in DB and MetaProgramming](https://www.reddit.com/r/rails/comments/gjkha1/databasedriven_authorization_in_rails_using/)
- url: https://www.reddit.com/r/rails/comments/gjkha1/databasedriven_authorization_in_rails_using/
---
Hi ruby family,

As an initiative to give back to the community, I have started writing a series of blogs on ruby and ruby on rails. A few days back, I published a post on **authorizing resources in rails using CanCanCan**. As a continuation of the previous post, I have recently published another post on how to **implement database-driven authorization using CanCanCan**.

Some of the key issues that I tried to solve was :

1. The Growing size of the ability file
2. Abilities being hard to maintain.
3. Redeployment of the application for every change in the ability file
4. Storing abilities in a database

If you think this can be extrapolated and be made into a gem, let me know, and let's work together to create an awesome library.

[https://addytalks.tech/2020/05/14/rails-cancancan-database-driven-authorization/](https://addytalks.tech/2020/05/14/rails-cancancan-database-driven-authorization/)

You check out my previous post here - 

[https://addytalks.tech/2020/05/03/ruby-on-rails-authorization-with-cancancan/](https://addytalks.tech/2020/05/03/ruby-on-rails-authorization-with-cancancan/)
## [11][Avoid coupling between Bounded Contexts using Weak Schema](https://www.reddit.com/r/rails/comments/gjo1hw/avoid_coupling_between_bounded_contexts_using/)
- url: https://www.reddit.com/r/rails/comments/gjo1hw/avoid_coupling_between_bounded_contexts_using/
---
See how to avoid accidental coupling on many levels while developing modular monolith application using event driven approach.

[https://blog.arkency.com/avoid-coupling-between-bounded-contexts-using-weak-schema/](https://blog.arkency.com/avoid-coupling-between-bounded-contexts-using-weak-schema/)
