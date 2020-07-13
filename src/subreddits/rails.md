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
## [2][Good service pattern architecture guides?](https://www.reddit.com/r/rails/comments/hqe0zs/good_service_pattern_architecture_guides/)
- url: https://www.reddit.com/r/rails/comments/hqe0zs/good_service_pattern_architecture_guides/
---
When first getting into Rails (or even web development in general) there is great documentation that explains the basics as well as the Rails framework itself doing a lot of the heavy lifting such as scaffolding, migrations, routing, etc.

If the Rails guides are to be believed, all you need are some models doing simple validations with all logic happening in the controllers wiring them together. At the same time, as a best practice, models should not get too fat and no business logic should happen in the controller. It seems like we're missing a layer in-between: the service. (for convenience lets consider an API-only project without the views)

Coming from the Java / Spring world, it's very clear-cut where the separation happens between controller, service and 'models'. I don't see how this would be applied in Rails, if it is even necessary or what best practices are. Creating something like a *hexagonal architecture* seems completely alien to me in the Rails world. Does that even make sense?

Can you recommend any good guides that touch this of bridging the gap between Controller and Model for anything a bit fancier than *just* a CRUD? Or maybe explain whether I might be thinking with the completely wrong mindset?
## [3][app has different sessions for www and non-www](https://www.reddit.com/r/rails/comments/hq7tua/app_has_different_sessions_for_www_and_nonwww/)
- url: https://www.reddit.com/r/rails/comments/hq7tua/app_has_different_sessions_for_www_and_nonwww/
---
The www and non-www versions of my Rails app on Heroku are two completely different sessions (I'm using Devise), and I want to prevent www by redirecting to https.

[This](https://help.heroku.com/J2R1S4T8/can-heroku-force-an-application-to-use-ssl-tls) says to use `config.force_ssl = true` in `config/environments/production.rb` to force https, but that just redirects http to https. How can I redirect all www requests to https as well?
## [4][Rails API + React Authentication](https://www.reddit.com/r/rails/comments/hq9oja/rails_api_react_authentication/)
- url: https://www.reddit.com/r/rails/comments/hq9oja/rails_api_react_authentication/
---
So im pretty familiar with rails and "somewhat" familiar with using rails as an API. I've done a bit with react (enough to know the basics/etc..)

However im a bit confused on the react side on how to handle authentication. Im guessing JWT is probably the most popular? I've used devise before, but it was years ago and only with rails.

I assume I just essentially want to store the user on react side on the state (I guess I would just store the user and whether they are logged in?) Seems like this would be ideal for redux store (I just need to get more familiar with redux) so it'd be pretty annoying to have to keep track of that user/logged in state for every component.

Anyways, is there a standard way to do this currently?

Thanks! (FWIW Im re-writing a "facebook clone" app with rails/react that I created for TheOdinProject)
## [5][read messages when deliver Immediately in rabbitmq](https://www.reddit.com/r/rails/comments/hqfcc9/read_messages_when_deliver_immediately_in_rabbitmq/)
- url: https://www.reddit.com/r/rails/comments/hqfcc9/read_messages_when_deliver_immediately_in_rabbitmq/
---
i want way to read messages when deliver Immediately in rabbitmq . for example when services send message to exchange , another services receives this message Immediately .
## [6][How are Rails' "scaling issues" different from any other framework?](https://www.reddit.com/r/rails/comments/hq0744/how_are_rails_scaling_issues_different_from_any/)
- url: https://www.reddit.com/r/rails/comments/hq0744/how_are_rails_scaling_issues_different_from_any/
---
I've heard people complain before that "Rails doesn't scale", because: 

* the fact that Ruby/Rails processes seem to consume a lot of memory
* "Ruby is slow" which I don't think is true... always... [https://benchmarksgame-team.pages.debian.net/benchmarksgame/fastest/yarv-python3.html](https://benchmarksgame-team.pages.debian.net/benchmarksgame/fastest/yarv-python3.html) (I guess Node is faster than Ruby JIT now?)
* You have to run workers, either on the same instance as your web server daemons or on other instance. 
* You have to be able to scale said workers and web servers for high loads
* Rails codebases encourage monolithic design patterns that causes them to grow too much, even if Ruby can handle boilerplate code really well (through Procs and Metaprogramming).

So these all seem like reasonable complaints, but what about Django, or even Phoenix?

Just to be clear: I've never developed in Django or Phoenix, but I'd like to hear about the experiences of people here who worked with Django. It's a very similar framework, built on an interpreted language, using the same architectural patterns and migration system. Maybe this is some sort of selection bias, but I haven't heard anyone say "Django doesn't scale". Wouldn't these issues be present in any framework that follows MVC and encourages codebases like this? Laravel? Ember? Etc.
## [7]['="?](https://www.reddit.com/r/rails/comments/hqd9m8/_/)
- url: https://www.reddit.com/r/rails/comments/hqd9m8/_/
---
Very simple one this. In code, I see ' and ". Can you use either for the same thing? i.e. are they the same?
## [8][Having trouble getting a view helper method to not return an array using haml](https://www.reddit.com/r/rails/comments/hq6o28/having_trouble_getting_a_view_helper_method_to/)
- url: https://www.reddit.com/r/rails/comments/hq6o28/having_trouble_getting_a_view_helper_method_to/
---
Hello all,

I am working on a small project where I am needing to loop through an array to decide which partial to render.

The primary view has the following

    =sections_list

Which corresponds to the following code

    def sections_list
      sections = []
        
      (@document.sections + [@document]).sort_by(&amp;:sort_order).each do |record|
        if record.respond_to?(:document_id)
          sections &lt;&lt; render('admin/sections/section_row', section: record)
        else
          sections &lt;&lt; render('admin/documents/approve_button_row', document: record)
        end
      end
      raw sections
    end

In the view, the partials show correctly but the array is being returned with the following in it:

    ["\n\n\n\n\n\n\n", "\n\n\n\n\n\n\n"]

The array has the same number of elements as what is being displayed in the view. 

As I know it, if I use a `-` instead of `=` in the view, the array wouldn't be returned, but since this is in a helper I have not figured out how to not get the array to return. I have tried doing things like return an empty value but since that is the last thing evaluated, the records don't show up in the view at all.

Does anyone have any ideas I could try?
## [9][would you rather write (A) or (B) ?](https://www.reddit.com/r/rails/comments/hq936d/would_you_rather_write_a_or_b/)
- url: https://www.reddit.com/r/rails/comments/hq936d/would_you_rather_write_a_or_b/
---
  

    # exhibit (A)
    module ApplicationHelper
      def title(str)
        content_for(:title) { str }
      end
    end
    
    # exhibit (B)
    module ApplicationHelper
      def title(title)
        content_for(:title) { title }
      end
    end
## [10][I want to run an old rails 3.3.2 project, seeking particular and specific steps for that](https://www.reddit.com/r/rails/comments/hq3f8s/i_want_to_run_an_old_rails_332_project_seeking/)
- url: https://www.reddit.com/r/rails/comments/hq3f8s/i_want_to_run_an_old_rails_332_project_seeking/
---

I want to run this project, https://github.com/ebidadmin/ebid33 I tried rails s and bundle install but was instead faced upon a series of different bundler version compatability and dependency issues. It's been a week or so that I'm trying to run. I just feel like I've reached a dead end.

My Ruby Version is 2.6.3, I also have 1.8.7 Installed. My rails version is 6.0.3.2
## [11][When, if ever, should Associations be avoided?](https://www.reddit.com/r/rails/comments/hq85h8/when_if_ever_should_associations_be_avoided/)
- url: https://www.reddit.com/r/rails/comments/hq85h8/when_if_ever_should_associations_be_avoided/
---
Learning RoR now. My prototype is a stock app. I have one table that holds all the trade information (e.g. symbol, date, price etc...). Now I want to add another table that holds all the information about the individual stocks (e.g. symbol, last price, description, beta, etc...).

My plan was just to have the trade model refer to the stock table just with a `stock_id` and without an association, because a trade isn’t a child of a stock, nor is it a parent.

However, their is some sort of ‘relationship’ going on there. The stock table has master data that the trade table refers to. The info in the stock table **should** never be deleted (since stocks are rarely de-listed), so no need to cascade deletes, which is where I see associations and foreign keys being necessary.

In this master/referrer relationship, does establishing has_many/belongs_to associations in the models provide any real benefit? Or just add unneeded complexity?

Thanks!
