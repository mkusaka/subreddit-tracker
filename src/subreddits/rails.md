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
## [2][Audited vs papertrail gem for auditing table changes?](https://www.reddit.com/r/rails/comments/jk9i0y/audited_vs_papertrail_gem_for_auditing_table/)
- url: https://www.reddit.com/r/rails/comments/jk9i0y/audited_vs_papertrail_gem_for_auditing_table/
---
I'm currently looking to use either one of these two for auditing changes made by specific user roles on a rails 6 ecommerce app using mysql. They both seem similar and I'm currently leaning towards papertrail due to the more frequent updates, but I was curious if there is any specific advantage of using one over the other(performance, ease of use...etc) or anyone knows of any pros/cons that stick out.

Would be keen to know if people have used any other gems that seemed good as well.
## [3][How is github able to provide go to definition and find references while I can't do it in VS Code](https://www.reddit.com/r/rails/comments/jjpqnw/how_is_github_able_to_provide_go_to_definition/)
- url: https://www.reddit.com/r/rails/comments/jjpqnw/how_is_github_able_to_provide_go_to_definition/
---
In VS code I struggle to get something like below to get to work. I have installed solargraph but it takes a long time to find something and most of the time it doesn't work?   


Any idea on how this works in github and if it's possible to get this to work in vs code? 

https://preview.redd.it/iuy85bnvruv51.png?width=3096&amp;format=png&amp;auto=webp&amp;s=628207d307a41c12920cf466b5d76b483f520774
## [4][A Gem for Cloudimage Image Optimizer](https://www.reddit.com/r/rails/comments/jjzf5z/a_gem_for_cloudimage_image_optimizer/)
- url: https://www.reddit.com/r/rails/comments/jjzf5z/a_gem_for_cloudimage_image_optimizer/
---
**Hi! Our tech team at Scaleflex just published a gem for our image optimizer Cloudimage** ([https://www.cloudimage.io/](https://www.cloudimage.io/))-&gt; Discover it here: Cloudimage-gem ([https://rubygems.org/gems/cloudimage](https://rubygems.org/gems/cloudimage))

What you can achieve with this gem:

* supports Ruby 2.4 and above, JRuby, and TruffleRuby
* on-the-fly resizing of all your images (150+ transformations)
* image optimization and WebP/AVIF (tba) compression
* responsive design, progressive/lazy loading
* worldwide media delivery via CDN

Happy to get your feedback!

Greetings!
## [5][Techmaker TV YouTube](https://www.reddit.com/r/rails/comments/jjoevj/techmaker_tv_youtube/)
- url: https://www.reddit.com/r/rails/comments/jjoevj/techmaker_tv_youtube/
---
Hey guys, I've been watching the Techmaker TV stuff for Rails on YouTube and the content is fantastic: https://www.youtube.com/c/TechmakerTV/videos

I noticed, though, that there haven't been any videos in over a month and the producer usually created about 2 videos a week.  

Anyone know anything about it?  Just curious if he was still making videos and everything was alright.
## [6][What are some of the Questions/tests you’ve been asked in technical interviews for Rails positions?](https://www.reddit.com/r/rails/comments/jjj6zi/what_are_some_of_the_questionstests_youve_been/)
- url: https://www.reddit.com/r/rails/comments/jjj6zi/what_are_some_of_the_questionstests_youve_been/
---

## [7][Destroy 1 mil records in a few minutes?](https://www.reddit.com/r/rails/comments/jjmtub/destroy_1_mil_records_in_a_few_minutes/)
- url: https://www.reddit.com/r/rails/comments/jjmtub/destroy_1_mil_records_in_a_few_minutes/
---
```ruby
Order.where("end_date &lt;= ?", date).in_batches(of: 100).destroy_all
```

I tried this code with 1 mil records (500,000 of `orders` and 500,000 of associated records) and it takes over an hour to finish. Is there any better way in ActiveRecord to handle this problem? Can we achieve it in just a few minutes?
## [8][Formception!?](https://www.reddit.com/r/rails/comments/jjo68f/formception/)
- url: https://www.reddit.com/r/rails/comments/jjo68f/formception/
---
Hi there,  


I am relatively new to rails and I am looking to add a form within a form, or simply just two models within the same form. I have plans, which I have a new plan view. And for every-time a user makes a new plan I want them to be able to associate people with that plan. So on creation of the plan, I want the user to also have the option to add any number of people (name, dob, gender etc.)  


What would be the best way to do this?  


Cheers
## [9][Rails MODEL.where returns nil, but records exists](https://www.reddit.com/r/rails/comments/jjkehy/rails_modelwhere_returns_nil_but_records_exists/)
- url: https://www.reddit.com/r/rails/comments/jjkehy/rails_modelwhere_returns_nil_but_records_exists/
---
 I have a model with several associations.

A topic has several questions and each question can have several answers.

I am trying to get an answer associated with a question here:

    def answer_for_question(question_id)  
       answers.where(question_id: question_id).first 
    end

 In order to show this on a page with: 

 `&lt;% answer = @survey_lever.answer_for_question(question.id)%&gt;` 

But  this returns an empty array. If I check 

    answers

in the console there is 10 elements in the association cache. 

`answers` are being created on `@survey_lever.answers` here

      def find_or_create_survey_lever(lever)
        survey_lever = survey_levers.find_or_create_by(lever_id: lever.id)
    
        lever.questions.each do |question|
          if survey_lever.no_answer?(question)
            survey_lever.answers.create(question: question)
          end
        end
    
        survey_lever
      end

`question_id` does exist and is being passed. If I manually traverse the object tree I can get to the first element in the answers array that this should return. There is nothing being created in the database.

Can someone tell me what I am missing?
## [10][Some initializers not working in a Rails 6 app](https://www.reddit.com/r/rails/comments/jjk19o/some_initializers_not_working_in_a_rails_6_app/)
- url: https://www.reddit.com/r/rails/comments/jjk19o/some_initializers_not_working_in_a_rails_6_app/
---
I am trying to get an initializer working to configure a couple of things in a custom library. This is a fairly recent Rails 6 project, pretty much vanilla stuff.

I have done this in the past but in this project the initializer does not seem to be working as I remember it would.

Let's start with the module I am trying to configure :

    module MyModule
      class &lt;&lt; self
        attr_accessor(:configuration)
      end
      def self.configure
        self.configuration || Configuration.new
        yield(configuration)
      end
      class Configuration
        attr_accessor :my_setting
        def initialize
          @my_setting
        end
      end
    end

This is based on Thoughtbot's example, it's a bit more than other examples ou there but it's the same melody in the end.

So, the way to use this is the following :

    MyModule.configure { |c| c.my_setting = "a value" }

This works ok within the console : I can access the value after running this bloc.

It seems that the initializer is run when the app starts (if I add a `raise StandardError` the app crashes) but  if I try to access the value at a later time to use it, then it's empty.

As there been changes around initializer in recent Rails releases that could explain this ?

related posts describing this pattern :

\- [https://www.skcript.com/svr/the-easiest-configuration-block-for-your-ruby-gems/](https://www.skcript.com/svr/the-easiest-configuration-block-for-your-ruby-gems/)

\- [https://www.leighhalliday.com/block-based-configuration](https://www.leighhalliday.com/block-based-configuration)

\- [https://thoughtbot.com/blog/mygem-configure-block](https://thoughtbot.com/blog/mygem-configure-block)

issue on GitHub on this : [https://github.com/rails/rails/issues/37186](https://github.com/rails/rails/issues/37186) ; I also check production vs development env locally and indeed production env works as expected.

&amp;#x200B;

Rails version : [6.0.3.4](https://6.0.3.4)

Ruby version : 2.7.1p83

OS : Linux

&amp;#x200B;

update : fix auto correct on attr\_accessor

update2 : related posts

update3 : related issue in Rails' github
## [11][Error with keys and model relations](https://www.reddit.com/r/rails/comments/jjayr9/error_with_keys_and_model_relations/)
- url: https://www.reddit.com/r/rails/comments/jjayr9/error_with_keys_and_model_relations/
---
I have 3 models, a client model, a product model and a catalog model. I basically want the catalog model to have many products and for each client the model changes. I get this error although I think there's more that I'm missing. Any approaches that you guys recommend?

    No route matches {:action=&gt;"new", :controller=&gt;"catalogs"}, missing required keys: [:client_id]
