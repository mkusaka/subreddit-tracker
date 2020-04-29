# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/g616hm/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/g616hm/personal_projects_show_off_your_own_project_andor/
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
## [2][Is my understanding of Delayed Job and Sidekiq correct?](https://www.reddit.com/r/rails/comments/g9yy6q/is_my_understanding_of_delayed_job_and_sidekiq/)
- url: https://www.reddit.com/r/rails/comments/g9yy6q/is_my_understanding_of_delayed_job_and_sidekiq/
---
I've used [Sidekiq](https://github.com/mperham/sidekiq) in the past, but am taking a look at [Delayed Job](https://github.com/collectiveidea/delayed_job/tree/v4.1.8). 

Below is how I understand the two. Please let me know if this is correct.

## Sidekiq
* Performs jobs created by [Active Job](https://guides.rubyonrails.org/active_job_basics.html#creating-a-job)
* Requires Redis

## Delayed Job
* Used primarily to [call methods](https://github.com/collectiveidea/delayed_job/tree/v4.1.8#queuing-jobs) directly on models, rather than calling an [Active Job](https://guides.rubyonrails.org/active_job_basics.html#creating-a-job)
* Saves jobs to the database.
## [3][Where can you find most of the Rails/Ruby community?](https://www.reddit.com/r/rails/comments/g9r5d4/where_can_you_find_most_of_the_railsruby_community/)
- url: https://www.reddit.com/r/rails/comments/g9r5d4/where_can_you_find_most_of_the_railsruby_community/
---
I’ve been learning Rails recently and would really love to see more of the community. I really enjoy the Rail and Ruby conference talks on YouTube (even if some are over my head) but would enjoy getting to interact with the community more. 

Where can I find it? Is there a forum, discord, etc? I’d also appreciate any other related media (i.e. podcast recommendations, news outlets, etc.)
## [4][How to handle multiple User types with completely different access rights?](https://www.reddit.com/r/rails/comments/ga5sdk/how_to_handle_multiple_user_types_with_completely/)
- url: https://www.reddit.com/r/rails/comments/ga5sdk/how_to_handle_multiple_user_types_with_completely/
---
I have 3 `User` types: a `Customer`, a `Vendor` and an `Admin`. I also have a couple of models, one of them being `Shipment`. Now I want to create an index action to list all the shipments that a user can see. However, because every user type has access to different shipments I am wondering about the best way to achieve that. Here's the problem:

A customer has access to all the shipments that he created (something like `current_customer.shipments`). An admin has access to all shipments that all the customers that he "controls" created (something like `current_admin.customers.shipments`). A vendor is even more specialized, as he can access the shipments for a user by passing in a user_id (something like `User.find(params[:user_id].shipments`). There is a little bit more logic to a vendor but that is the rough idea.

I basically see two ways to solve this problem:

1. I could either create a `ShipmentsController#index` action together with Pundit Scopes like they are demonstrating [here](https://github.com/varvet/pundit#scopes). That would mean that the index action only has a single line (`@shipments = policy_scope(Shipment)`) and all the authorization logic would happen inside Pundit. This would be nice because I would only need to create a single controller but with the downside that I think this controller wouldn't really adhere to the single responsibilty principle as it would essentially do way more. 

2. The other solution would be to create a `User::`, `Vendor::` and `Admin::` namespace and then create a `ShipmentsController` in each of them. This would be nice because each controller would have a clear cut responsibility but at the cost that if a change with Shipment happens I might need to change three different files which doesn't feel all to DRY.

I have this problem in multiple models, not just shipments. What would be the best way to handle this? Put everything inside one controller with Pundit or create multiple namespaces?
## [5][Ruby adds experimental support for end-less method definition](https://www.reddit.com/r/rails/comments/g9hmac/ruby_adds_experimental_support_for_endless_method/)
- url: /r/ruby/comments/g9hlar/ruby_adds_experimental_support_for_endless_method/
---

## [6][[Inquiry] Rails Way of Response](https://www.reddit.com/r/rails/comments/g9sunp/inquiry_rails_way_of_response/)
- url: https://www.reddit.com/r/rails/comments/g9sunp/inquiry_rails_way_of_response/
---
Hi Everyone,

Maybe you read my previous post (maybe not) as a beginner I'm trying to read as much as possible and adapt on the concept before doing the hard-coding part.

&amp;#x200B;

I've really learn alot thanks to this community (as a silent reader and sometimes answer on some that I think relevant ) .

&amp;#x200B;

I've tried searching for an answer my self different forums/google answer etc.

Hoping you can give some advise on what's the considered 'rails-way' on sending back to response.

&amp;#x200B;

Note: I understood it will be base on what kind of application but if you can give some pro tips on pro's n cons, it will be really helpful.

&amp;#x200B;

Thanks!!!

&amp;#x200B;

[View Poll](https://www.reddit.com/poll/g9sunp)
## [7][Markdown redcarpet and link_to](https://www.reddit.com/r/rails/comments/g9iik3/markdown_redcarpet_and_link_to/)
- url: https://www.reddit.com/r/rails/comments/g9iik3/markdown_redcarpet_and_link_to/
---
 I'm improving my markdown render.   

    class MarkdownRenderer &lt; Redcarpet::Render::HTML
       def paragraph(text)
         text.gsub! (/@(\w+)/) do |match|
           %(&lt;a target="_blank" href="/users/#{match[1..-1]}"&gt;#{match}&lt;/a&gt;)
         end
         text
       end

One question... as you can see I used an html code 

    %(&lt;a target="_blank" href="/users/#{match[1..-1]}"&gt;etc. etc.

 **I want to turn it in link\_to. Is it possible?** I tried with 

       text.gsub! (/@(\w+)/) do |match|
           &lt;%= link_to match, user_path(match) %&gt;
         end

 but it doesn't work.

It said  No method for class match

So I   used `#{1}`

    &lt;%= link_to match, user_path($1) %&gt;

But it said : `syntax error "&lt;"`

I tried to add directly `link_to match, user_path($1)` but it says me `undefined method 'user_path'` 

How to solve?!
## [8][HTTParty, converting curl command into a ruby request and general API confusion](https://www.reddit.com/r/rails/comments/g95gjd/httparty_converting_curl_command_into_a_ruby/)
- url: https://www.reddit.com/r/rails/comments/g95gjd/httparty_converting_curl_command_into_a_ruby/
---
Hi everyone,

I'm pretty confused right now, and I've been chasing my tail for a while now trying to do something and just getting further and further away from resolving it. People in this group have been extremely helpful in the past so I'm hoping I can get some guidance on my current hangup.

I think the best way to explain my current confusion is to start at my original problem to demonstrate how I got to where I am now.

Original problem:I'm trying to hit the Algolia Places API, but need to hide my API key

So I know how to work with a basic API when it comes to the Javascript fetch API, and normally the JS documentation is very easy so I can just copy their code, but since I need to hide my API key, I need to do it in the back end of my rails app.

Ok, so the last time I did this, this is is how I did it in my controller:

    response = open(url).read
    @response_json = JSON.parse(response)
    render json: @response_json

I don't know where "open" or ".read" actually come from, because I can't recall where I learned how to do a request this way, and I can't seem to find any rails documentation on this, so my first question is where are these methods coming from and where can I read the documentation?Now for this particular api call the documentation gives clear instructions on how to structure the URL, and it is only a GET request, so all I am giving this "open" method is a URL string. Nice and simple.

However Algolia places requires a POST request, and also the documentation don't demonstrate with a simple URL string (if that's even possible with a post request?), they show how to make the request with a curl command like so:

    $ curl -X POST 'https://places-dsn.algolia.net/1/places/query' \
      -H 'X-Algolia-Application-Id: YourApplicationId' \
      -H 'X-Algolia-API-Key: YourAPIKey' \
      --data '{"query": "Paris"}'

Now, I'm not going to be doing this request in my terminal, so I need to somehow convert this to something I can do in my controller.

This brings me to HTTParty. After doing some googling it seems like a popular tool for making API requests. Most of the tutorials seem to show HTTParty integrating with a class (as does the documentation). This seems like overkill to me. All I want to do is hit the api with a query, I don't see why that can't just happen in my controller? So question number 2 is what is the point in having a class just to do a seemingly basic thing of hitting a URL and getting some data?

However I did discover [this article](https://www.rubyguides.com/2018/08/ruby-http-request/) which seems to simplify HTTParty a bit.

So I've tried to convert that curl command to using HTTParty similar to what is demonstrated in that article, and this is how far I've gotten:

    HTTParty.post("https://places-dsn.algolia.net/1/places/query",
        :query =&gt; { :query =&gt; "paris" },
        :headers =&gt; {
          "X-Algolia-Application-Id" =&gt; ENV["PLACES_APP_ID"],
          "X-Algolia-API-Key" =&gt; ENV["FOURSQUARE_CLIENT_ID"]
        })

However when I try this in the console I get (and I don't know where to begin debugging..)

    {"message"=&gt;"No content in POST request", "status"=&gt;400}

If you've read this far then thanks for sticking with me. I appreciate this is a rather convoluted post with a series of questions, but I think there's something I'm just fundamentally not understanding about how to hit APIs and work with sending/receiving data, because I just seem to be hitting walls all over the place.

I've done my share of googling but I just can't seem to demystify any of this, so any help would be hugely appreciated.  


Edit: I've managed to get it working using a class now, however I can't seem to translate that curl request into a HTTParty request. Here's what I could figure after looking at the examples in the documentation:  


    def initialize(query)
        @options = {
          headers: {
          "X-Algolia-Application-Id" =&gt; ENV["PLACES_APP_ID"],
            "X-Algolia-API-Key" =&gt; ENV["FOURSQUARE_CLIENT_ID"]
          },
          query: {
            query: query
          }
        }
      end

This results in the error:

    Failed to open TCP connection to places-dsn.algolia.net:443 (getaddrinfo: nodename nor servname provided, or not known)

&amp;#x200B;
## [9][Graph in rails](https://www.reddit.com/r/rails/comments/g9hoer/graph_in_rails/)
- url: https://www.reddit.com/r/rails/comments/g9hoer/graph_in_rails/
---
I have income, expense and user controller i want to plot graph of profit loss of current user i have following code in user model.

    class User &lt; ApplicationRecord   has_many :incomes   has_many :expenses    def total_expense     expenses.pluck(:amount).sum   end def total_incomes     incomes.pluck(:amount).sum   end def profit_loss     total_incomes - total_expense   end # Include default devise modules. Others available are: # :confirmable, :lockable, :timeoutable, :trackable and :omniauthable   devise :database_authenticatable, :registerable, :recoverable, :rememberable, :validatable end

but proft\_loss  
 is not saved in database.
## [10][how to write a very basic rails 5 api](https://www.reddit.com/r/rails/comments/g99h69/how_to_write_a_very_basic_rails_5_api/)
- url: https://www.reddit.com/r/rails/comments/g99h69/how_to_write_a_very_basic_rails_5_api/
---
how to write a **Rails 5 (or 6) API that is so simple** it has no database and all it does is  that it gets two parameters and run a method on them to return a value.

all tutorials i found had database, or involve testing which make things complicated for me. I just need to a finish a little project then i will learn rails api the right way. I have very little experience with Rails. 

But writing a little guide for my task will help me a lot. to understand the basics, including where to put the method..etc
## [11][Junior Dev Question: How do you approach database/model redesign?](https://www.reddit.com/r/rails/comments/g98zri/junior_dev_question_how_do_you_approach/)
- url: https://www.reddit.com/r/rails/comments/g98zri/junior_dev_question_how_do_you_approach/
---
I'm a new dev (roughly 6 months) on a rails project that the team has deemed in need of a back-end overhaul: expanding the scope of some models, reducing the scope of others, adding new relations, etc. As the requirements are being developed, I've tried to start formulating a high level outline of "things I need to do/consider before a major backend refactor," and I feel like I'm hitting that plateau of "not knowing what I need to know." Do you have any recommendations/resources to consider when refactoring your models/databases? Things to look out for, easy performance improvements, etc?  


TLDR: I'm a n00b that doesn't know what search terms I should be using when considering refactoring our rails models.
