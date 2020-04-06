# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/foqc07/personal_projects_show_off_your_own_project_andor/)
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
## [2][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/ftl6xc/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/ftl6xc/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [3][Is Stimulus the best option to add modern Javascript to an existing Rails app?](https://www.reddit.com/r/rails/comments/fvxiqx/is_stimulus_the_best_option_to_add_modern/)
- url: https://www.reddit.com/r/rails/comments/fvxiqx/is_stimulus_the_best_option_to_add_modern/
---
I've got an existing Rails app, and I want to add some modern Javascript features to existing pages. Think taking existing elements and forms and making them feel more interactive.

I saw the [GoRails video on using Vue slots](https://gorails.com/episodes/vue-js-slots), and started going down that route, until I realised the way it works is quite a hack (it takes your existing HTML and recreates it all in Vue which is bad for performance).

Stimulus seems to be somewhat unmaintained (last release was a year ago) so I'm wondering if there is something else I should be using.
## [4][How to unsubscribe from a model](https://www.reddit.com/r/rails/comments/fvy7rc/how_to_unsubscribe_from_a_model/)
- url: https://www.reddit.com/r/rails/comments/fvy7rc/how_to_unsubscribe_from_a_model/
---
Hi Guys! I'm adding a subscription feature where users can subscribe/unsubscribe to/from a genre. I've been able to create subscription but I don't know how to unsubscribe if a user has already subscribed.

    subscription table
    
    create_table "subscriptions", force: :cascade do |t|
        t.integer "user_id"
        t.integer "genre_id"
        t.datetime "created_at", precision: 6, null: false
        t.datetime "updated_at", precision: 6, null: false
        t.index ["genre_id"], name: "index_subscriptions_on_genre_id"
        t.index ["user_id"], name: "index_subscriptions_on_user_id"
     end
    
    subscription.rb
    
    belongs_to :user
    belongs_to :genre
    
    subscriptions_controller.rb
    
    class SubscriptionsController &lt; ApplicationController
    
        def create
            @subscription = Subscription.new(subscription_params)
            @subscription.user_id = current_user.id
            @subscription.save
    
            redirect_to genre_path(@subscription.genre_id)
        end
    
        def destroy
            if @subscription.destroy
                redirect_back(fallback_location: root_path)
                flash[:danger] = %Q[&lt;i class="fa fa-trash"&gt;&lt;/i&gt; You have successfully 
            unsubscribed from this genre.]
            end
        end
    
        private
    
        def subscription_params
            params.require(:subscription).permit(:genre_id)
        end
    end
    
    user.rb
    
    has_many :subscriptions, dependent: :destroy
    has_many :genres, through: :subscriptions
    
    genre.rb
    
    has_many :subscriptions
    has_many :subscribers, through: :subscriptions, source: :user
    
    genres_controller.rb
    
    def show
          @subscriber_count = @genre.subscribers.count
          @is_subscribed = user_signed_in? ? Subscription.where(genre_id: @genre.id, 
       user_id: current_user.id).any? : false
          @subscription = Subscription.new
          @sub_genres = Subscription.where(genre_id: @genre)
    end
    
    genres/show
    
    &lt;div class="subscribe"&gt;
            &lt;% if @is_subscribed %&gt;
                &lt;%= link_to @subscription, method: :post, class: "btn btn-danger btn-sm" 
            do %&gt;
                    &lt;%= fa_icon "times" %&gt; Unsubscribe
                &lt;% end %&gt;
            &lt;% else %&gt;
                &lt;%= render partial: "genres/subscribe", locals: { genre_id: @genre.id } 
            %&gt;
            &lt;% end %&gt;
    &lt;/div&gt;
    
    genres/_subscribe
    
    &lt;%= form_with(model: @subscription, local: true) do |form| %&gt;
        &lt;%= form.hidden_field :genre_id, value: genre_id %&gt;
        &lt;%= form.submit "Subscribe", class: "btn-primary" %&gt;
    &lt;% end %&gt;
    
    routes.rb
    
    resources :subscriptions
## [5][Collecting JavaScript code coverage with Capybara in Ruby on Rails application](https://www.reddit.com/r/rails/comments/fvtuqd/collecting_javascript_code_coverage_with_capybara/)
- url: https://jtway.co/collecting-javascript-code-coverage-with-capybara-in-ruby-on-rails-application-d0cb83a86a90
---

## [6][Benefits to making model over text schema field?](https://www.reddit.com/r/rails/comments/fvri9z/benefits_to_making_model_over_text_schema_field/)
- url: https://www.reddit.com/r/rails/comments/fvri9z/benefits_to_making_model_over_text_schema_field/
---
Sorry for the confusing question. Essentially I am a rails "intermediate" and I am making a sort of device/sensor type dashboard/web-app.

Devices have 3 types: Humidity/Temp/Both. Devices will have measurements being recorded/historical data/etc..

When it comes to the model/schema, I had wondered: "Is it REALLY worth it to make a WHOLE other model with "has\_one" just for the device type, should I just make it a text field?". As in do I really need a model defined for something that isn't going to have a whole lot of interaction (Besides maybe a drop-down selection for choosing the device type when integrating a sensor).

This also goes as a general rails question I suppose: When do I make a model for something vs it just being a field?
## [7][Heroku build failing because "package-lock.json found", but here is no package-lock.json file in my repo.](https://www.reddit.com/r/rails/comments/fvkjyp/heroku_build_failing_because_packagelockjson/)
- url: https://www.reddit.com/r/rails/comments/fvkjyp/heroku_build_failing_because_packagelockjson/
---
I am getting the following error on build:

```
warning package-lock.json found. Your project contains lock files generated by tools other than Yarn. It is advised not to mix package managers in order to avoid resolution inconsistencies caused by unsynchronized lock files. To clear this warning, remove package-lock.json.
```

However, there is no package-lock.json in [my repository](https://github.com/stevepolitodesign/po-notes).

Any advice on how this can be solved?
## [8][I wrote a post about how to test your software](https://www.reddit.com/r/rails/comments/fvdb4s/i_wrote_a_post_about_how_to_test_your_software/)
- url: https://www.reddit.com/r/rails/comments/fvdb4s/i_wrote_a_post_about_how_to_test_your_software/
---
I recently asked the Rails community on Reddit the following [question](https://www.reddit.com/r/rails/comments/fgcy4s/is_there_any_part_of_rails_you_wished_you/): Is there any part of Rails you wish you understood better? The most common answer was… testing.

Ship software with confidence ––&gt;

[https://twitter.com/liroyleshed/status/1246777273317044225](https://twitter.com/liroyleshed/status/1246777273317044225)
## [9][Integrating React and Rails in 2020](https://www.reddit.com/r/rails/comments/fvd4dc/integrating_react_and_rails_in_2020/)
- url: https://www.reddit.com/r/rails/comments/fvd4dc/integrating_react_and_rails_in_2020/
---
I'm aware there's been topics like this, but they are at best 1 year old. Rails 6 came out since, webpacker is bundled in by default.

I am novice to working with React/JS (mostly some jQuery before), but I have a need to make an app for which jQuery/StimulusJS won't be enough (I tried).

My main question is:

**Should I use** `react-rails`, `react_on_rails` **or create two separate apps (Rails-BE, React-FE)?**

Seems that `react_on_rails`  had bad reviews here, but this was before Rails 6 with default webpacker. I'm also somewhat hesitant to go for the last options as communication done via JSON sounds like a bad choice performance-wise.  Thoughts?
## [10][Stimulus Reflex](https://www.reddit.com/r/rails/comments/fvd1cs/stimulus_reflex/)
- url: https://www.reddit.com/r/rails/comments/fvd1cs/stimulus_reflex/
---
I am looking into Stimulus Reflex and it looks nice, but I am trying to figure out what I could use it for. Are you using it in real apps? Thanks
## [11][Open Source Social Network built in Rails and Jquery](https://www.reddit.com/r/rails/comments/fvccj0/open_source_social_network_built_in_rails_and/)
- url: https://www.reddit.com/r/rails/comments/fvccj0/open_source_social_network_built_in_rails_and/
---
Hey! I am building a social network for coders to share and interact called [Devolio](https://www.devol.io). It's still in its early stages so there isnt much content on the platform yet but I really want to grow it! I've open sourced [the code](https://github.com/Uzay-G/devolio) and I hope to build something for devs alongside other passionate devs. Please tell me what you think because it really means a lot!
## [12][Redis for permanent data storage](https://www.reddit.com/r/rails/comments/fvdvan/redis_for_permanent_data_storage/)
- url: https://www.reddit.com/r/rails/comments/fvdvan/redis_for_permanent_data_storage/
---
I'm working on a project in which I need a number of features, mainly the ability to recommend users and products to other users based on a number of factors - likes, dislikes, purchase count etc.

I came across the [recommendable](https://davidcel.is/recommendable/) gem which covers most of my needs, and has a few nice features added such as hiding or bookmarking. The issue though, is that the data is persisted only in Redis, which as far as I know of, is primarily used as a temporary data store.

Are there any drawbacks to this set up? One thing that crossed my mind immediately that was as you have more and more users, the memory usage would probably spike up out of the room and it would be much more costly to store it in Redis rather than an RDBMS. Not only that, but the consistency of the data seemed to worry me.

I was considering modifying the gem to use active record to store the like objects but I wanted to verify the veracity of my concerns first.
