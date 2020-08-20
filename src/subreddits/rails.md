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
## [2][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/id911w/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/id911w/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [3][Rails on Heroku: Guide to how many dynos and which size](https://www.reddit.com/r/rails/comments/icnl4q/rails_on_heroku_guide_to_how_many_dynos_and_which/)
- url: https://www.reddit.com/r/rails/comments/icnl4q/rails_on_heroku_guide_to_how_many_dynos_and_which/
---
I just published an [exhaustive and opinionated guide](https://railsautoscale.com/how-many-dynos/) to dynos on Heroku. It answers the questions I've been hearing over and over for years:

* How many dynos should you be running?
* Which dyno type is right for your app?

I hope you find it helpful!
## [4][How to implement a Quiz functionality in Rails ?](https://www.reddit.com/r/rails/comments/iczwgn/how_to_implement_a_quiz_functionality_in_rails/)
- url: https://www.reddit.com/r/rails/comments/iczwgn/how_to_implement_a_quiz_functionality_in_rails/
---
I was thinking in how could I make something similar of [EngVid](https://www.engvid.com/) platform in Rails.

Each video has a "quiz" attached to train and test what you have learned in the video.

It's quite hard for me to grasp on how this quiz could be make in Rails.

I could have a Video model which could contain a title, description and has\_one\_attached :clip (which would be the uploaded video) and each Video has\_one Quiz. But I'm not sure how to implement the Quiz model.

Could I apply the same logic as I did in my ToDo list ? Each Quiz model would have a nested Question model. But again I'm not sure how to implement the options and correct answer for each Question.

Any ideas or tips ?
## [5][Using scopes on unsaved relations?](https://www.reddit.com/r/rails/comments/id2k1u/using_scopes_on_unsaved_relations/)
- url: https://www.reddit.com/r/rails/comments/id2k1u/using_scopes_on_unsaved_relations/
---
Hey all,

I have a question that I don't know how to formulate a Google search for.

Consider we have an Invoice model which `has_many :receivables`.
There is a `sub_total` field on the invoice which is supposed to stay in sync with summing most of the receivables, so I added logic like:


    before_validation do  
      self.sub_total = receivables.not_freight.sum(&amp;:line_price)  
    end

And this works great for existing records. However, when creating a new invoice, `receivables.not_freight` always returns an empty collection (even when `receivables` has many entries which would satisfy the scope). Interestingly, `receivables.not_freight.to_sql` returns an empty string.

I've rewritten the logic to be:

    self.sub_total = receivables.reject(&amp;:freight?).sum(&amp;:line_price)

and this works fine, as expected.  

  

What is this behaviour called, and where could I learn more about it? That code snippet looks innocuous but leads to weird data consistency bugs for new records.
## [6][Rails and Devise current_user method](https://www.reddit.com/r/rails/comments/icxg4c/rails_and_devise_current_user_method/)
- url: https://www.reddit.com/r/rails/comments/icxg4c/rails_and_devise_current_user_method/
---
Super noob with Ruby and Rails here, so I apologize in advance if this is a basic question. So, I am using the "devise" gem for authentication.

In my controller, I have this:
    
      def new_current_user
        puts "*" * 50
        puts current_user
        puts "*" * 50
        if current_user == "admin"    &lt;------ this is not correct
            @my_user = "admin"
            @my_role = ["admin"]
        else
            @my_user = "dev"
            @my_role = ["dev"]
        end
    
        puts @my_user
        puts @my_role
    
        @current_user = {
          :username =&gt; @my_user,
          :capabilities =&gt; @my_role,
        }
      end
In the database, I have these rows.

    mysql&gt; select id, email from users;
    +----+----------------+
    | id | email          |
    +----+----------------+
    |  3 | admin@test.com |
    |  2 | dev@test.com   |
    +----+----------------+
    2 rows in set (0.00 sec)
    
Basically, I don't know how to write if the current_user is "admin@test.com", then let the my_user equal to admin?
## [7][Nginx Mina Deployment strangeness.](https://www.reddit.com/r/rails/comments/icv5si/nginx_mina_deployment_strangeness/)
- url: https://www.reddit.com/r/rails/comments/icv5si/nginx_mina_deployment_strangeness/
---
I have a rails 6 app with is deployed via Mina. Until this week I deployed without an issue. Now when I deploy, it sticks on the old deployment version until I restart Nginx. Is there something that I am missing? Has anyone seen this before?
## [8][Rails API 401 Unauthorized](https://www.reddit.com/r/rails/comments/icmvdf/rails_api_401_unauthorized/)
- url: https://www.reddit.com/r/rails/comments/icmvdf/rails_api_401_unauthorized/
---
Hey, I have a Rails API that works wonderfully on localhost.

Authentication with jwt and knock session tokens.

Just pushed my code to Netlify and Heroku. Changed the config for it to communicate with Netlify app and only call that gets 401 Unauthorized is get users/me endpoint.

It lets you login and it sets the jwt token as expected but it can not set the user.

&amp;#x200B;

&amp;#x200B;

[The failing request](https://preview.redd.it/uqwgsamzeyh51.png?width=1014&amp;format=png&amp;auto=webp&amp;s=6d05f993b3eccd907b5001637ef4d2821fa9bc4f)

&amp;#x200B;

&amp;#x200B;

[Request headers of the failing request](https://preview.redd.it/p9seyqzokyh51.png?width=1204&amp;format=png&amp;auto=webp&amp;s=ca85f2b62fa281b5c1360ec6d27c401735fc4e1c)

&amp;#x200B;

[Login and users\/me API calls](https://preview.redd.it/zk4rdkm1fyh51.png?width=1210&amp;format=png&amp;auto=webp&amp;s=917e76f293de4f85dddd777853d01e603f6734d8)

&amp;#x200B;

[My User Controller](https://preview.redd.it/y7ytgnl5fyh51.png?width=1089&amp;format=png&amp;auto=webp&amp;s=b00e548279397a95654f293d86910f4c1004ee0b)

&amp;#x200B;

&amp;#x200B;

[Heroku logs](https://preview.redd.it/wdm4vf0cgyh51.png?width=1351&amp;format=png&amp;auto=webp&amp;s=f0836fd76df74379730947b4d02912724a9f4f30)

Full code can be found in here :

Frontend =&gt; [https://github.com/gitwithuli/white-curtain-frontend](https://github.com/gitwithuli/white-curtain-frontend)

Backend =&gt; [https://github.com/gitwithuli/white-curtain-backend](https://github.com/gitwithuli/white-curtain-backend)

&amp;#x200B;

Any help would be greatly appreciated.

Thanks.
## [9][How to include a concern inside of my controller_spec file?](https://www.reddit.com/r/rails/comments/ic9sof/how_to_include_a_concern_inside_of_my_controller/)
- url: https://www.reddit.com/r/rails/comments/ic9sof/how_to_include_a_concern_inside_of_my_controller/
---
RSpec newbie here. I have a before action method 'get\_user'  that I moved into it's own concern from the user\_controller file. Now my user\_controller\_spec file is broken. How do I include the concern, and thus my 'get\_user' method, in the spec file?
## [10][Parsing Input for Employee Data](https://www.reddit.com/r/rails/comments/ic3uuk/parsing_input_for_employee_data/)
- url: https://www.reddit.com/r/rails/comments/ic3uuk/parsing_input_for_employee_data/
---
I’m working on an app that helps with employee scheduling and want to find an easy way for users to load employee data into the application (I currently have a form that adds them one at a time).  The data would include name, start date, role, ect, and be in an unpredictable format.  I was thinking of having a textarea and finding a way for the application to parse relevant data and create resources from it.  

Is such a feature a pipe dream?  This is my first web-dev personal project.  Any advice would be appreciated.
## [11][how can I mix slim and erb in same file?](https://www.reddit.com/r/rails/comments/ick27c/how_can_i_mix_slim_and_erb_in_same_file/)
- url: https://www.reddit.com/r/rails/comments/ick27c/how_can_i_mix_slim_and_erb_in_same_file/
---
    
    = form_for @reg, url: update_tbd_path(@reg.state) do |f|
    
      .form-group
        = f.submit "Next", class: 'form-control btn btn-green tbd', id: 'form-submit', 'data-test': 'register-submit-button'
    
    
          &lt;div class="form-group hidden" id="tbd-date-picker"&gt;
            &lt;p class="question-text"&gt;What was the date?&lt;/p&gt;
              &lt;%= ft.label 'tbd_start_date', 'Date', id: 'tbd-start-date' %&gt;
            &lt;div class="input-group date"&gt;
              &lt;%= ft.text_field 'tbd_start_date', class: 'form-control tbd datepicker', placeholder: 'MM/DD/YYYY' %&gt;
              &lt;span class="input-group-addon calendar-icon"&gt;&lt;%= image_tag('icons/calendar@3x.png')%&gt;&lt;/span&gt;
            &lt;/div&gt;
          &lt;/div&gt;
    
    
    h1 Info
    
    p When do you intend to?
    
    p DatePicker
    
    p TimePicker
## [12][Help me decide on a search solution (between algolia, searchkick and pg_search)](https://www.reddit.com/r/rails/comments/ibx0pz/help_me_decide_on_a_search_solution_between/)
- url: https://www.reddit.com/r/rails/comments/ibx0pz/help_me_decide_on_a_search_solution_between/
---
Hi guys

&amp;#x200B;

I'm looking for a good search solution for my application.

The project has rails as a graphql api and a react-native app as client.

What I am specifically looking for is a way to search for users but also users in specific contexts, like search a list of friends or make sure that blocked users are not showing up. Preferrably, while also providing some other data from the db, such as the amount of posts or other data, as part of the returned results.

Currently I was looking at the following solutions:

\- **Algolia:** Easy to implement to some extent but it seems a bit finicky having to add a list of friends/blocked users etc. as part of the saved indices in Algolia itself. Plus it seems more difficult to include other data from the db when needed. Plus (and this is the biggest downside) it's not free and cost quite a lot when an application scales. 

\- **Searchkick/Elasticsearch**: 2 months ago I asked a similar yet more vague question regarding search in Rails and some people pointed out to me how easy it is to use elasticsearch with searchkick. At the same time there were people who thought it was overkill and warned me of the complexities of Elastic Search. So I'm not sure what to believe but I feel this is not necessarily more complex than setting something similar up on Algolia?

\- **pg\_search**: This is a gem I just discovered today and although I'm not sure, it might be good enough for my purposes? It uses postgres full text search but also allows search scopes/facetting, which would be the answer when searching through someone's friends or blocked users.

I would love to know what your experience is. I'm not sure if it matters too much but this a production app and not a simple learning project, so having a scalable solution is needed.   


Thank you!
