# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/jbmadq/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/jbmadq/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/jfcv1r/personal_projects_show_off_your_own_project_andor/)
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
## [3][What are some of the Questions/tests you’ve been asked in technical interviews for Rails positions?](https://www.reddit.com/r/rails/comments/jjj6zi/what_are_some_of_the_questionstests_youve_been/)
- url: https://www.reddit.com/r/rails/comments/jjj6zi/what_are_some_of_the_questionstests_youve_been/
---

## [4][Destroy 1 mil records in a few minutes?](https://www.reddit.com/r/rails/comments/jjmtub/destroy_1_mil_records_in_a_few_minutes/)
- url: https://www.reddit.com/r/rails/comments/jjmtub/destroy_1_mil_records_in_a_few_minutes/
---
```ruby
Order.where("end_date &lt;= ?", date).in_batches(of: 100).destroy_all
```

I tried this code with 1 mil records (500,000 of `orders` and 500,000 of associated records) and it takes over an hour to finish. Is there any better way in ActiveRecord to handle this problem? Can we achieve it in just a few minutes?
## [5][Rails MODEL.where returns nil, but records exists](https://www.reddit.com/r/rails/comments/jjkehy/rails_modelwhere_returns_nil_but_records_exists/)
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
## [6][Error with keys and model relations](https://www.reddit.com/r/rails/comments/jjayr9/error_with_keys_and_model_relations/)
- url: https://www.reddit.com/r/rails/comments/jjayr9/error_with_keys_and_model_relations/
---
I have 3 models, a client model, a product model and a catalog model. I basically want the catalog model to have many products and for each client the model changes. I get this error although I think there's more that I'm missing. Any approaches that you guys recommend?

    No route matches {:action=&gt;"new", :controller=&gt;"catalogs"}, missing required keys: [:client_id]
## [7][Some initializers not working in a Rails 6 app](https://www.reddit.com/r/rails/comments/jjk19o/some_initializers_not_working_in_a_rails_6_app/)
- url: https://www.reddit.com/r/rails/comments/jjk19o/some_initializers_not_working_in_a_rails_6_app/
---
I am trying to get an initializer working to configure a couple of things in a custom library. This is a fairly recent Rails 6 project, pretty much vanilla stuff.

I have done this in the past but in this project the initializer does not seem to be working as I remember it would.

Let's start with the module I am trying to configure :

    module MyModule
      class &lt;&lt; self
        attractive_accessor(:configuration)
      end
      def self.configure
        self.configuration || Configuration.new
        yield(configuration)
      end
      class Configuration
        attractive_accessor :my_setting
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

Rails version : [6.0.3.4](https://6.0.3.4)

Ruby version : 2.7.1p83

OS : Linux
## [8][HOWTO: highlight link_to current_page](https://www.reddit.com/r/rails/comments/jiznev/howto_highlight_link_to_current_page/)
- url: https://www.reddit.com/r/rails/comments/jiznev/howto_highlight_link_to_current_page/
---
Short post on how to highlight `link_to` `current_page`: [https://blog.corsego.com/2020/10/ruby-on-rails-highlight-linkto.html](https://blog.corsego.com/2020/10/ruby-on-rails-highlight-linkto.html) I hope you find it useful :)

P.S. There's an old gem [active\_link\_to](https://github.com/comfy/active_link_to), but what I offer is a very simple alternative😎

https://preview.redd.it/awsrnkrnlmv51.png?width=338&amp;format=png&amp;auto=webp&amp;s=05b5ace0671b8a100ada1fb7fe4d7082230fe532
## [9][Using rails engines/plugins in a rails application](https://www.reddit.com/r/rails/comments/jj33dh/using_rails_enginesplugins_in_a_rails_application/)
- url: https://www.reddit.com/r/rails/comments/jj33dh/using_rails_enginesplugins_in_a_rails_application/
---
So, I bumped into this for the first time and I'm clueless about how to proceed with it. I've to build a rails plugin/engine and then use it in another application. I've researched and found obsolete videos articles dealing with older versions of rails(they happen to have a different directory structure for plugins). And I'm unable to wrap my head around what's really going on here and What I should be doing to set it up.   
My main concerns are how to create a plugin and use it in another rails application. Need to figure out which code goes where because they folder structure is very ambiguous and I have no idea about which is the correct folder structure.   
See my stackoverflow post for details: [https://stackoverflow.com/questions/64555911/how-to-include-a-plugin-engine-in-a-rails-application](https://stackoverflow.com/questions/64555911/how-to-include-a-plugin-engine-in-a-rails-application)
## [10][Good guides/articles regarding Rails 6 deployment to AWS](https://www.reddit.com/r/rails/comments/jiob7g/good_guidesarticles_regarding_rails_6_deployment/)
- url: https://www.reddit.com/r/rails/comments/jiob7g/good_guidesarticles_regarding_rails_6_deployment/
---
Hello everyone!

I decided to keep all my Rails stuff inside of one service(AWS S3 and hosting on AWS) but I struggle to find a good guide on how to deploy Rails to AWS and that the guide is up to date.

It would be really cool if you could point me towards some good articles about this topic.

Thank you!
## [11][Any good courses/tutorials for Rails API + Vue out there?](https://www.reddit.com/r/rails/comments/jiozti/any_good_coursestutorials_for_rails_api_vue_out/)
- url: https://www.reddit.com/r/rails/comments/jiozti/any_good_coursestutorials_for_rails_api_vue_out/
---
Hi 

I'd like to get more into working with the combination of a Rails API and Vue, but I'm looking to follow a tutorial or course to get the hang of it some more before I dive into my own project. 

I already followed a [tutorial from WebCrunch](https://web-crunch.com/posts/ruby-on-rails-api-vue-js) but felt like a lot of corners were being cut there. 

Anybody have some up-to-date recommendations for where I could get started? Would really appreciate it!
## [12][Introducing jQuery in Rails 6 Using Webpacker](https://www.reddit.com/r/rails/comments/jiyqvc/introducing_jquery_in_rails_6_using_webpacker/)
- url: https://www.reddit.com/r/rails/comments/jiyqvc/introducing_jquery_in_rails_6_using_webpacker/
---
Hello Everyone,

First release candidate of Rails 6 is out with exciting features and refinements in existing features. Today we will look into one of those features, which is webpacker. It is now by default in Rails 6 as a Javascript Compiler.

Here’s a detailed blog on [Introducing jQuery in Rails 6 Using Webpacker](https://www.botreetechnologies.com/blog/introducing-jquery-in-rails-6-using-webpacker)

Hope this will be helpful to #railscommunity.
