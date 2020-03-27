# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/fgx7fz/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/fgx7fz/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/foqc07/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/foqc07/personal_projects_show_off_your_own_project_andor/
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
## [3][Database - is there any benefit to index a list of names?](https://www.reddit.com/r/rails/comments/fpv4h6/database_is_there_any_benefit_to_index_a_list_of/)
- url: https://www.reddit.com/r/rails/comments/fpv4h6/database_is_there_any_benefit_to_index_a_list_of/
---
Say I have a model Company, it has one single :name column. The purpose of this model is just to store company names. The only use for this model is to create a select list when needed (i.e Select All), and to **find by id** when a company name is required (i.e I won't search the Company model directly). The user can manage the company list (create/delete/etc).

Is there any reason to (or not to) index the :name column?
## [4][Solopreneur owner of 20k/month Rails app gives inside view.](https://www.reddit.com/r/rails/comments/fpenft/solopreneur_owner_of_20kmonth_rails_app_gives/)
- url: https://www.reddit.com/r/rails/comments/fpenft/solopreneur_owner_of_20kmonth_rails_app_gives/
---
Do you remember the [Solopreneur owner of 20k/month Rails app who wanted to give an inside view](https://www.reddit.com/r/rails/comments/ev2s0n/solopreneur_owner_of_20kmonth_rails_app_wants_to/)?

His first video is up on [Youtube](https://youtu.be/aCzT-LQI6x0)

EDIT: Not my work - The developer is /u/semicolonandsons
## [5][Nested has_many not calling validations -- sometimes?](https://www.reddit.com/r/rails/comments/fpre8u/nested_has_many_not_calling_validations_sometimes/)
- url: https://www.reddit.com/r/rails/comments/fpre8u/nested_has_many_not_calling_validations_sometimes/
---
I'm currently looking at a pretty simple scenario that kinda has me stumped. Here is the setup:

    class modelA &lt; ApplicationRecord
        has_many :model_b, autosave: true
        has_many :model_c, through: :model_b
    end
    
    class modelB &lt; ApplicationRecord
        has_many :model_c, autosave: true
    end
    
    class modelC &lt; ApplicationRecord
        validate :may_return_false
    end

I found that when a given model\_c is invalid the respective model\_a can have this behavior:

    model_a_instance.valid? # true
    model_a_instance.model_bs.first.valid? # true
    model_a_instance.model_bs.map(&amp;:model_cs)
    model_a_instance.valid? # false
    model_a_instance.model_bs.first.valid? # false

Does anyone have any ideas why I need to "touch" the child model in order to make the parent invalid?
## [6][Bulma.io Omniauth](https://www.reddit.com/r/rails/comments/fpspms/bulmaio_omniauth/)
- url: https://www.reddit.com/r/rails/comments/fpspms/bulmaio_omniauth/
---
Hello! I used the framework Bulma to the Index/login, and i use Omniauth for sign in with facebook, but the button for redirect to FB is color Black, and i want to change to Blue
How i can do this? Sorry If i have grammatical errors:/
## [7][How do you do search autocompletes in Rails 6 (with webpacker)?](https://www.reddit.com/r/rails/comments/fps9zt/how_do_you_do_search_autocompletes_in_rails_6/)
- url: https://www.reddit.com/r/rails/comments/fps9zt/how_do_you_do_search_autocompletes_in_rails_6/
---
Everyone on the internet seems to post something that's kinda right, but nothing seems to be working?

What is the best way to get autocomplete working with Rails 6? For clarity, I don't care if I'm using jquery-ui, one of the fifty js packages called autocomplete.js, or whatever.

// just want to add a simple autocomplete search bar with minimal code (that's css customizable)
## [8][Authentication and authorization in microservices](https://www.reddit.com/r/rails/comments/fpua7o/authentication_and_authorization_in_microservices/)
- url: https://www.reddit.com/r/rails/comments/fpua7o/authentication_and_authorization_in_microservices/
---
I'm working on app which will contain some microservices, I've made an auth service which handles user creation and authentication, that service returns doorkeeper token with proper scope. And I'm not sure how handle tokens with requests for another microservices. I feel most comfortable with making microservices as independent as possible, when microservice get request with token then it verifies token by making call to Auth service, Auth service returns user Identity and then microservice performs own operations. But I'm wondering about security, performance and elegancy of this solution. Is api gateway solution cleaner and safer?
## [9][How do you guys typically use Stimulus + Rails?](https://www.reddit.com/r/rails/comments/fpdypd/how_do_you_guys_typically_use_stimulus_rails/)
- url: https://www.reddit.com/r/rails/comments/fpdypd/how_do_you_guys_typically_use_stimulus_rails/
---
Hey guys, I've been in search for an incredibly productive stack for solo dev (I know this is highly subjective). I've been enjoying Rails due to the developer productivity it provides. I tried Elixir + Phoenix, and I think they are INCREDIBLY cool tech, but their package maturity and community size leave much to be desired.

Currently my stack entails:

* Ember (Ember-paper for UI, Auth0 for auth)
* Rails + Graphiti gem for JSON-API (amazing)



I very much enjoy how much I can lean on convention with these technologies, but I'm curious about what dev workflow would look like to use Rails, Stimulus, Turbolinks, Graphiti (https://www.graphiti.dev/quickstart), and Spraypaint (https://www.graphiti.dev/js/).


Are you guys still typically hitting a highly structured internal api? What do you all do for frontend themes/animations/etc?
## [10][[Help] Google Analytics works on restricted pages?](https://www.reddit.com/r/rails/comments/fpos88/help_google_analytics_works_on_restricted_pages/)
- url: https://www.reddit.com/r/rails/comments/fpos88/help_google_analytics_works_on_restricted_pages/
---
As the title said, can GA analyze statics on restricted pages (pages that demand login from the user)?. 

If analysts doesn't, what companies usually use to keep track on user behaviors on these pages? Do they roll their on implementation?
## [11][Does Rails 6 work well with Postgres 12?](https://www.reddit.com/r/rails/comments/fpk8ks/does_rails_6_work_well_with_postgres_12/)
- url: https://www.reddit.com/r/rails/comments/fpk8ks/does_rails_6_work_well_with_postgres_12/
---
Or should I stick with 11 for now? I am going to rebuild my k8s cluster so was wondering which version is recommended. Thanks
## [12][How to get crowd funding on a project?](https://www.reddit.com/r/rails/comments/fpi01w/how_to_get_crowd_funding_on_a_project/)
- url: https://www.reddit.com/r/rails/comments/fpi01w/how_to_get_crowd_funding_on_a_project/
---
What is the best strategy? If you read my earlier posts you will see I am not into scam business, so don't insult me, you can, and we will laugh it off, but what is the point? So I am thinking it would be unethical to just post an idea. Oh I got one which I don't like to share just yet :) So ANYWAY let's say I got a project at Iteration #1-2. It is really baffling to me that ANYONE would put investment into that, let alone donation. Are some people that Rich? But maybe there are people with money and thus I should just post first iteration, give my ideas for next, and hope they would pitch in? I am Really lost here, mates.....
