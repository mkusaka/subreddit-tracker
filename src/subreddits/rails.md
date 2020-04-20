# rails
## [1][I've published a new post about how to use Stimulus.js](https://www.reddit.com/r/rails/comments/g4pwp6/ive_published_a_new_post_about_how_to_use/)
- url: https://www.reddit.com/r/rails/comments/g4pwp6/ive_published_a_new_post_about_how_to_use/
---
I wrote a new post about how you can integrate Stimulus into your application.

How to use Stimulus.js ––&gt;

[https://www.thatweeklytech.com/posts/15-how-to-use-stimulus-js](https://www.thatweeklytech.com/posts/15-how-to-use-stimulus-js)

Thanks for reading!
## [2][Webpacker : how to hot-reload HTML ?](https://www.reddit.com/r/rails/comments/g4p458/webpacker_how_to_hotreload_html/)
- url: https://www.reddit.com/r/rails/comments/g4p458/webpacker_how_to_hotreload_html/
---
With Docker +  Webpacker, it is possible to hot-reload JS, it is as simple as set hmr: true in the webpacker.yml file.

It is also possible to hot-reload CSS, even if they are sprockets-generated : a few configuration in development.js (as some listed here [https://github.com/rails/webpacker/issues/1879](https://github.com/rails/webpacker/issues/1879)) and you're done.

But I'm still struggling to hot-reload HTML (in .html.erb files). In old days it was achieved thanks to guard, but I'm looking for a solution with webpacker to avoid a tool-stacking nightmare. How can it be achieved ?
## [3][Realtime, reactive web apps without Javascript using Stimulus Reflex](https://www.reddit.com/r/rails/comments/g48fm2/realtime_reactive_web_apps_without_javascript/)
- url: https://www.reddit.com/r/rails/comments/g48fm2/realtime_reactive_web_apps_without_javascript/
---
Thought you guys would be interested in this.

Stimulus Reflex lets you build realtime, reactive apps in Rails similar to Phoenix LiveView. You don't have to write any Javascript and it will use Rails to render all the HTML updates server-side. It uses ActionCable to trigger updates and then tells Rails to render the current page again and send that back over the websocket. It'll then use DOM diffing to update the page automatically for you.

Definitely one of the coolest projects going on right now in the Rails world I'd say. It looks like we might see some similar things in Turbolinks 6 when that comes out too which is exciting.

[https://gorails.com/episodes/stimulus-reflex-basics?autoplay=1](https://gorails.com/episodes/stimulus-reflex-basics?autoplay=1)
## [4][Having difficulty understanding `has_many :something, through: :something_else, source: :another_thing`, here is a good brief lesson](https://www.reddit.com/r/rails/comments/g4dgsj/having_difficulty_understanding_has_many/)
- url: https://www.reddit.com/r/rails/comments/g4dgsj/having_difficulty_understanding_has_many/
---
primarily recommend reading this lesson

https://open.appacademy.io/learn/full-stack-online/sql/associations--has_many--through

What's important is to note that _existing associations_ are what is being referred to in the `through` and `source` parts of the association, and they are not references to a particular model, but to associations/methods already declared.

take this code:

    class Physician &lt; ApplicationRecord
      has_many(
        :appointments,
        class_name: 'Appointment',
        foreign_key: :physician_id,
        primary_key: :id
      )

      has_many :patients, through: :appointments, source: :patient
    end

    class Appointment &lt; ApplicationRecord
      belongs_to(
        :physician,
        class_name: 'Physician',
        foreign_key: :physician_id,
        primary_key: :id
      )

      belongs_to(
        :patient,
        class_name: 'Patient',
        foreign_key: :patient_id,
        primary_key: :id
      )
    end

    class Patient &lt; ApplicationRecord
      has_many(
        :appointments
        class_name: 'Appointment',
        foreign_key: :patient_id,
        primary_key: :id
      )

      has_many :physicians, through: :appointments, source: :physician
    end


In `Physician`, the `has_many :patients, through: :appointments, source: :patient`:

1. the `:patients` will be a new method for your `Physician` model, `Physician#patients`.  

2. the `through: :appointments` is using your existing `Physician#appointments` method/association already declared in the `Physician` model of course.

3. the `source: :physician` is using an _existing_ association/method that is already declared too, namely, the join model `Appointments#physician` method/association in the `Appointments` model.

I struggle sometime to conceptualize this, above is a very brief synopsis, but I recommend going to the lesson, it's free.
## [5][GitHub - styd/pry-diff-routes: Inspect routes changes in Rails console.](https://www.reddit.com/r/rails/comments/g45csl/github_stydprydiffroutes_inspect_routes_changes/)
- url: https://www.reddit.com/r/rails/comments/g45csl/github_stydprydiffroutes_inspect_routes_changes/
---
Hi fellow Rails developer,

First of all, let's keep being safe and healthy and enjoy staying at home. Secondly, I want to say thank you to those of you who have contributed to my older gem, [ApexCharts.RB](https://github.com/styd/apexcharts.rb), by using, reporting issues, sending pull requests, or just giving the repo a star (I couldn't believe it at first but even [Chartkick](https://chartkick.com) author Andrew Kane and [ApexCharts](https://apexcharts.com) author Juned Chhipa also gave a star. Thank you!). I sincerely appreciate all of your efforts.

Today, I want to tell you about a new gem that I've been working on since last month called [PryDiffRoutes](https://github.com/styd/pry-diff-routes). It's a Pry plugin that works in Rails console to see the route changes that we make.

The way it works is that it would save the current routes as Rails console is starting and then from that point on the routes you change -- don't forget to save the route file -- is comparable with the routes when you started the Rails console. When you type the Pry command `diff-routes`, it will output your changes as removed, modified, or new routes.

What's "modified"? And why don't I just display the **before** and the **after**?
IMHO, when we talk about an endpoint, what we really meant usually is a specific combination of a _verb_ and a _path_. When they don't change, we consider it as the same endpoint. By changing either the _verb_ or the _path_, we can say that we are removing the old route and creating a new route. When we just change where it routes to, we are modifying the route. So, what I'm doing is just group routes changes based on our intention for those changes.

I hope it is useful for you and let me know what you think about it.
Happy weekend!
## [6][Can I use Active Storage to store files in a user's personal storage service?](https://www.reddit.com/r/rails/comments/g4c9se/can_i_use_active_storage_to_store_files_in_a/)
- url: https://www.reddit.com/r/rails/comments/g4c9se/can_i_use_active_storage_to_store_files_in_a/
---
I would like to entertain the idea of "not owning" the users' uploaded files. Some applications store the data or files related to your account in your own Google Drive etc. Is this possible with Active Storage?
## [7][How do you organise your request specs?](https://www.reddit.com/r/rails/comments/g4anq7/how_do_you_organise_your_request_specs/)
- url: https://www.reddit.com/r/rails/comments/g4anq7/how_do_you_organise_your_request_specs/
---
I only have experience with rspec controller specs for testing my  controllers. These are typically organised with 1 spec file per controller, and then a describe block per method/action, such as:

```
describe UsersController do
  describe 'GET #show' do
    it 'is successful' do
      ....
    end
  end
end
```

People who have experience writing request specs, do you organise them in a similar way? 1 file per controller? 1 describe block per action?  What naming conventions do you use for the describe blocks and files?

Thanks
## [8][Rails Associations best practices?](https://www.reddit.com/r/rails/comments/g46rgf/rails_associations_best_practices/)
- url: https://www.reddit.com/r/rails/comments/g46rgf/rails_associations_best_practices/
---
So this is something i've been curious about for awhile, and that is the best way to structure associations. Specifically when it comes to chaining has\_many down a tree or using a more spread out hierarchy.

For example:

Lets say I have a User that can leave comments. Comments can also have likes. I could say that a User has\_many comments and that a Comment has\_many likes and that a User has\_many Likes through comments.

OR

I Could say a User has\_many Likes and has\_many Comments (and I suppose also could say a Comment has\_many Likes as well).

Is one way strictly better than the other? And when we can avoid join tables should we?
## [9][How to seek a rails dev?](https://www.reddit.com/r/rails/comments/g492s2/how_to_seek_a_rails_dev/)
- url: https://www.reddit.com/r/rails/comments/g492s2/how_to_seek_a_rails_dev/
---
Hello all,

So, for a few years now, i have been thinking about an app. And i feel, now is the time to have it realized. For the backend i decided to go with rails, because, well, thats what i know. Unfortunatly, i am not yet this good in rails that i feel confident i could build it all on my own. With that in mind we decided to look for a crowdfunding project, and hire a dev.

But where does one start?? 

I was thinking, i should first find a dev before setting up crowdfunding, because i would need to know the pricetag. I have never hired anyone, let alone a dev that could be somewhere on the world.. so i wonder what anyone here thinks?
## [10][Need help properly configuring my database.yml](https://www.reddit.com/r/rails/comments/g4gzth/need_help_properly_configuring_my_databaseyml/)
- url: https://www.reddit.com/r/rails/comments/g4gzth/need_help_properly_configuring_my_databaseyml/
---
Aloha new rails jr dev here. Ive been stressing for days trying to run this repo: [https://github.com/chinweenie/Full-Stack-Project-Etsy](https://github.com/chinweenie/Full-Stack-Project-Etsy)

i have learned A LOT so far , but i am struggling to with my database.yml configuration, here is my code : [https://pastebin.com/YkhkhYwb](https://pastebin.com/YkhkhYwb)

and here is the error im getting in the browser: Puma caught this error: Cannot load database configuration: undefined method \`each' for "gem 'pg'":String (NoMethodError)

ive busted my as to get this far, here is the guide that helped me the most : [https://vitux.com/how-to-install-latest-ruby-on-rails-on-ubuntu/](https://vitux.com/how-to-install-latest-ruby-on-rails-on-ubuntu/)

i know im close  can anyone take the time and patience to help me run this app ? 

mahalo
