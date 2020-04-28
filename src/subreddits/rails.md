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
## [2][Ruby adds experimental support for end-less method definition](https://www.reddit.com/r/rails/comments/g9hmac/ruby_adds_experimental_support_for_endless_method/)
- url: /r/ruby/comments/g9hlar/ruby_adds_experimental_support_for_endless_method/
---

## [3][Markdown redcarpet and link_to](https://www.reddit.com/r/rails/comments/g9iik3/markdown_redcarpet_and_link_to/)
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
## [4][HTTParty, converting curl command into a ruby request and general API confusion](https://www.reddit.com/r/rails/comments/g95gjd/httparty_converting_curl_command_into_a_ruby/)
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
## [5][Graph in rails](https://www.reddit.com/r/rails/comments/g9hoer/graph_in_rails/)
- url: https://www.reddit.com/r/rails/comments/g9hoer/graph_in_rails/
---
I have income, expense and user controller i want to plot graph of profit loss of current user i have following code in user model.

    class User &lt; ApplicationRecord   has_many :incomes   has_many :expenses    def total_expense     expenses.pluck(:amount).sum   end def total_incomes     incomes.pluck(:amount).sum   end def profit_loss     total_incomes - total_expense   end # Include default devise modules. Others available are: # :confirmable, :lockable, :timeoutable, :trackable and :omniauthable   devise :database_authenticatable, :registerable, :recoverable, :rememberable, :validatable end

but proft\_loss  
 is not saved in database.
## [6][how to write a very basic rails 5 api](https://www.reddit.com/r/rails/comments/g99h69/how_to_write_a_very_basic_rails_5_api/)
- url: https://www.reddit.com/r/rails/comments/g99h69/how_to_write_a_very_basic_rails_5_api/
---
how to write a **Rails 5 (or 6) API that is so simple** it has no database and all it does is  that it gets two parameters and run a method on them to return a value.

all tutorials i found had database, or involve testing which make things complicated for me. I just need to a finish a little project then i will learn rails api the right way. I have very little experience with Rails. 

But writing a little guide for my task will help me a lot. to understand the basics, including where to put the method..etc
## [7][Junior Dev Question: How do you approach database/model redesign?](https://www.reddit.com/r/rails/comments/g98zri/junior_dev_question_how_do_you_approach/)
- url: https://www.reddit.com/r/rails/comments/g98zri/junior_dev_question_how_do_you_approach/
---
I'm a new dev (roughly 6 months) on a rails project that the team has deemed in need of a back-end overhaul: expanding the scope of some models, reducing the scope of others, adding new relations, etc. As the requirements are being developed, I've tried to start formulating a high level outline of "things I need to do/consider before a major backend refactor," and I feel like I'm hitting that plateau of "not knowing what I need to know." Do you have any recommendations/resources to consider when refactoring your models/databases? Things to look out for, easy performance improvements, etc?  


TLDR: I'm a n00b that doesn't know what search terms I should be using when considering refactoring our rails models.
## [8][Removing routes that I didn't specify](https://www.reddit.com/r/rails/comments/g91wbq/removing_routes_that_i_didnt_specify/)
- url: https://www.reddit.com/r/rails/comments/g91wbq/removing_routes_that_i_didnt_specify/
---
I'm using Rails 6.0.2.2, and I added a route inside of config/routes.rb.  When I ran the 'rake routes' command, a bunch of routes appeared that I didn't specify.

They are:

/rails/action_mailbox/...

/rails/conductor/action_mailbox/inbound_emails/...

/rails/active_storage/...

I can't find any way to remove these, or prevent rails from generating these routes.  From what I've read on them, I don't need them for my intended website to work as I want it.  

How can these routes be removed / not-generated?
## [9][Beginning with Rails and I have a few questions](https://www.reddit.com/r/rails/comments/g8xv6i/beginning_with_rails_and_i_have_a_few_questions/)
- url: https://www.reddit.com/r/rails/comments/g8xv6i/beginning_with_rails_and_i_have_a_few_questions/
---
Hi everyone, I started to learn Rails last week, and I am so happy with it. I managed to get from a "hello world" level, to an actually usable app in less than a week, great experience.
That said, I have plenty of questions, and I can't really find some clear answers online. I read a lot of threads here, and learned a lot, so I was hoping some of you might help answer these questions.

First, some context:
I'm an artist manager, and my rails app is basically a CRM for my business (a team of 4 people). It has some basic CRUD features + a calendar, and allows me to create invoices etc. There is not a tons of datas (few hundred of entries on each table) but there is a lot of relations between these datas (e.g: a manager has many artists, an artist has many performances, a performance has many events (rehearsals, setup, travel), a performance has an invoice etc.).

I would like to actually use this project for my business, but before involving all my team, it would help a lot to have some answers to the following questions:


1 - For my database I use SQLite3. Being file-based, I think it is perfect for a beginner like who might benefit from having this file under version control, and in case of problem I could just revert to the previous version. I'm not that confident in doing this easily and quickly with MySQL. So is there a real benefit from using a better database for my small business ? Will I really see some better performances given that there is not a lof of entries, and 4 simultaneous user maximum ?

2 - What is the best strategy for assets, should I manage these on a per-page basis or should I just send the whole pack on the first page load ?

3 - I installed a few gems, and now my JS asset file is around 3Mb. Is that considered normal for a rails app (there is not a tons of gems: wicked_pdf, shrine, materializejs, jquery, actioncable, cocoon) ?

4 - I plan to add to my app a "settings" page where I can define some settings such as menu color etc. Basically, it is an array of settings. 
What is the best way to store this kind of datas ? It feels weird to create a table given that I will only use the first entry of this table. What do you think?

Edit:
5 - For my app I made a dashboard which display some general statistics. In order to display all the required values, I need to pass to this view all my datas (e.g: Artist.all, Performance.all, Invoice.all etc.).
Does this affect performance ? 

Many thanks in advance !
## [10][am rendering a partial from a custom controller action, but now it is only showing html and not erb](https://www.reddit.com/r/rails/comments/g9417v/am_rendering_a_partial_from_a_custom_controller/)
- url: https://www.reddit.com/r/rails/comments/g9417v/am_rendering_a_partial_from_a_custom_controller/
---
so I have made a controller custom action and have the route for it, and it routes properly(I've checked using the Rails Panel chrome plugin), but if put _only_ html tags like `&lt;p&gt;what's up&lt;/p&gt;` those work, but if I try to use erb like `&lt;% %&gt;` or `&lt;%= %&gt;` things don't render any more - even html tags. I'm not sure the reason for that and was wondering if anyone may have an idea as to why
## [11][Validates. Removing the url video.](https://www.reddit.com/r/rails/comments/g8xplp/validates_removing_the_url_video/)
- url: https://www.reddit.com/r/rails/comments/g8xplp/validates_removing_the_url_video/
---
Hi guys, I added this lines to validate the youtube video id.

It check only if the video is added and the length of the video is of 11 characters. Very easy.

        validate :check_videos_url!
    
        def check_videos_url!
          errors.add(
            :videos, :wrong_youtube_id
          ) if videos.any? &amp;&amp; videos.first.try(:url).length != 11
        end

But now the users reported me a bug.

If I create a page and I don't add any video -&gt; It works and creates the page.

If I create a page and I add a video of only 5 characters -&gt; It works, reporting me the error.

If I create a page and I add a video with 11 characters -&gt; It works and creates the page.

If I edit the page with no video, adding a new video -&gt; It works and edits the page.

If I edit the page with a video, adding a new video -&gt; It works and edits the page.

**If I edit the page with a video and I want to remove it -&gt; BUG. It say "it should be 11 characters" and I can't remove it.**

&amp;#x200B;

How to solve?
