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
## [2][Pagination and Scroll Restoration with Turbolinks](https://www.reddit.com/r/rails/comments/hqo8c3/pagination_and_scroll_restoration_with_turbolinks/)
- url: https://www.reddit.com/r/rails/comments/hqo8c3/pagination_and_scroll_restoration_with_turbolinks/
---
A lot of JS folks seemed to criticize some strange aspects of Hey's different approach to the front-end. This is a good post about a possible solution for one of the issues raised:

[http://blog.graykemmey.com/2020/07/07/pagination-and-scroll-restoration-with-turbolinks/](http://blog.graykemmey.com/2020/07/07/pagination-and-scroll-restoration-with-turbolinks/)
## [3][Everything you want to know about writing your own form objects in rails](https://www.reddit.com/r/rails/comments/hr1aol/everything_you_want_to_know_about_writing_your/)
- url: https://www.reddit.com/r/rails/comments/hr1aol/everything_you_want_to_know_about_writing_your/
---
Hey r/rails, I've written up a 3-part series on form objects (3 years in the making). If you've wrestled with rails' form helpers and builders to trying to make some frontend design work or with models that don't seem to cooperate, then you've probably already tried implementing form objects, or using a form object gem.

Thing is, there are many bad ways to do form objects, so I've compiled techniques and patterns I think are most compliant with The Rails Way (remember the days when this was still a thing). Here's what it covers:

[Part 1 - Ground rules for form objects](https://medium.com/@jaryl/disciplined-rails-form-object-techniques-patterns-part-1-23cfffcaf429)

* Implementing `ActiveModel` and its lifecycle
* Jumping from virtual objects to `ActiveRecord` instances
* Working with multiple models in one form object using `fields_for`

[Part 2 - Dealing with Collections](https://medium.com/@jaryl/disciplined-rails-form-object-techniques-patterns-part-2-12b8d530143d)

* Collections of primitive values, to virtual objects, to `ActiveRecord` instances
* Taking control of nested\_attributes and your form object's interface
* Using `ActiveModel` validations and promoting errors to your UI

[Part 3 - Tying up loose ends](https://medium.com/@jaryl/disciplined-rails-form-object-techniques-patterns-part-3-8ed1e4f62ce4)

* Simplifying complex forms (like advanced search) by working with `Arel`
* Some additional thoughts on validations and nesting
* A look at form object (as a opposed to form builder) gems, and whether to use them

It's a bit of a slog, but I promise you it's less of a slog than wading through 1000-line ERB files on a daily basis!
## [4][Crucial Resources](https://www.reddit.com/r/rails/comments/hqwfyg/crucial_resources/)
- url: https://www.reddit.com/r/rails/comments/hqwfyg/crucial_resources/
---
Hey everyone! I posted about two months ago a [repo](https://github.com/tylertomlinson/crucial_resources) with all the resources I've collected while starting out. I'm somewhat new to programming (10 months now!). Learning rails and so far its awesome! I had a bunch of local files and was told to put them in a [repo](https://github.com/tylertomlinson/crucial_resources) for others to use. They have helped me tremendously. They are also great when I'm traveling or without internet as they are local. Feel free to use and would love if you had any helpful so called "cheat sheets" you would want to throw in there.

Cheers!

Heres the repo

[https://github.com/tylertomlinson/crucial\_resources](https://github.com/tylertomlinson/crucial_resources)
## [5][redirect_to in conjunction with send_file?](https://www.reddit.com/r/rails/comments/hqsyhy/redirect_to_in_conjunction_with_send_file/)
- url: https://www.reddit.com/r/rails/comments/hqsyhy/redirect_to_in_conjunction_with_send_file/
---
I'm trying to allow a user to download a file and then redirect them to a separate page, but I'm reading that this would trigger two responses so the rails controller won't allow it. Are there any standard workarounds for this?
## [6][Good service pattern architecture guides?](https://www.reddit.com/r/rails/comments/hqe0zs/good_service_pattern_architecture_guides/)
- url: https://www.reddit.com/r/rails/comments/hqe0zs/good_service_pattern_architecture_guides/
---
When first getting into Rails (or even web development in general) there is great documentation that explains the basics as well as the Rails framework itself doing a lot of the heavy lifting such as scaffolding, migrations, routing, etc.

If the Rails guides are to be believed, all you need are some models doing simple validations with all logic happening in the controllers wiring them together. At the same time, as a best practice, models should not get too fat and no business logic should happen in the controller. It seems like we're missing a layer in-between: the service. (for convenience lets consider an API-only project without the views)

Coming from the Java / Spring world, it's very clear-cut where the separation happens between controller, service and 'models'. I don't see how this would be applied in Rails, if it is even necessary or what best practices are. Creating something like a *hexagonal architecture* seems completely alien to me in the Rails world. Does that even make sense?

Can you recommend any good guides that touch this of bridging the gap between Controller and Model for anything a bit fancier than *just* a CRUD? Or maybe explain whether I might be thinking with the completely wrong mindset?
## [7][Using Action Mailbox in Rails 6 to Receive Mail](https://www.reddit.com/r/rails/comments/hqi0lg/using_action_mailbox_in_rails_6_to_receive_mail/)
- url: https://www.reddit.com/r/rails/comments/hqi0lg/using_action_mailbox_in_rails_6_to_receive_mail/
---
A “mostly real-life” tutorial on using ActionMailbox to receive and parse incoming mail.

[https://robrace.dev/using-action-mailbox-in-rails-6-to-receive-mail/](https://robrace.dev/using-action-mailbox-in-rails-6-to-receive-mail/)
## [8][How to Helpers and API's](https://www.reddit.com/r/rails/comments/hqgh3w/how_to_helpers_and_apis/)
- url: https://www.reddit.com/r/rails/comments/hqgh3w/how_to_helpers_and_apis/
---
I am in the process of upgrading my knowledge as I upgrade my project from Rails 4 to Rails 6. I would like to know more about how to utilize **Helpers** and pulling in data from **External API's**. Can anyone point me to resources for these subjects.

When I google API's and rails I get results about how to create good API's and none on how to pulling data from other API's.
## [9][Rails API + React Authentication](https://www.reddit.com/r/rails/comments/hq9oja/rails_api_react_authentication/)
- url: https://www.reddit.com/r/rails/comments/hq9oja/rails_api_react_authentication/
---
So im pretty familiar with rails and "somewhat" familiar with using rails as an API. I've done a bit with react (enough to know the basics/etc..)

However im a bit confused on the react side on how to handle authentication. Im guessing JWT is probably the most popular? I've used devise before, but it was years ago and only with rails.

I assume I just essentially want to store the user on react side on the state (I guess I would just store the user and whether they are logged in?) Seems like this would be ideal for redux store (I just need to get more familiar with redux) so it'd be pretty annoying to have to keep track of that user/logged in state for every component.

Anyways, is there a standard way to do this currently?

Thanks! (FWIW Im re-writing a "facebook clone" app with rails/react that I created for TheOdinProject)
## [10][app has different sessions for www and non-www](https://www.reddit.com/r/rails/comments/hq7tua/app_has_different_sessions_for_www_and_nonwww/)
- url: https://www.reddit.com/r/rails/comments/hq7tua/app_has_different_sessions_for_www_and_nonwww/
---
The www and non-www versions of my Rails app on Heroku are two completely different sessions (I'm using Devise), and I want to prevent www by redirecting to https.

[This](https://help.heroku.com/J2R1S4T8/can-heroku-force-an-application-to-use-ssl-tls) says to use `config.force_ssl = true` in `config/environments/production.rb` to force https, but that just redirects http to https. How can I redirect all www requests to https as well?
## [11][read messages when deliver Immediately in rabbitmq](https://www.reddit.com/r/rails/comments/hqfcc9/read_messages_when_deliver_immediately_in_rabbitmq/)
- url: https://www.reddit.com/r/rails/comments/hqfcc9/read_messages_when_deliver_immediately_in_rabbitmq/
---
i want way to read messages when deliver Immediately in rabbitmq . for example when services send message to exchange , another services receives this message Immediately .
