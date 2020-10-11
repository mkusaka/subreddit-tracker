# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/iui4p8/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/iui4p8/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/j6qvuh/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/j6qvuh/personal_projects_show_off_your_own_project_andor/
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
## [3][What are the Limits of Ruby?](https://www.reddit.com/r/rails/comments/j93e2j/what_are_the_limits_of_ruby/)
- url: https://www.reddit.com/r/rails/comments/j93e2j/what_are_the_limits_of_ruby/
---
Hi Friends,

As a beginner, i'm reading that Ruby is an excellent language to learn web programming.

What are the Limits of Ruby?

I am interested in creating a 'Shopping Site' similar to those seen when shopping LIFE INSURANCE or HEALTH INSURANCE.

Here's a Lead company that does what i'm talking about. HOWEVER, it looks like its written in 'ASP' and 'PHP'
(not sure but i checked it in Buildwith.com)

[Here's the Site]( https://www.selectquote.com/)


1. -You can see it follows some-kinda Tree System....Anyone know the Term for this design?

1. -Can i do this in Ruby?

1. -Is it possible for a 1 user to accomplish this?


Thanks very much!
## [4][Difference between to: [rating.page.user] and to: page.user.followers.to_a to exclude a notification](https://www.reddit.com/r/rails/comments/j93lmd/difference_between_to_ratingpageuser_and_to/)
- url: https://www.reddit.com/r/rails/comments/j93lmd/difference_between_to_ratingpageuser_and_to/
---
If I rate a page, the author of the page will receive a notification.

BUT, as you can see, I added an if condition: if you are the author of the page (`rating.page.user`) you will not receive it.

      def after_create(rating)
        send_update_notifications(rating)
        assign_score rating.user, :create if rating.user != rating.page.user
        assign_score rating.page.user, :create if rating.user != rating.page.user
      end
    
      def send_update_notifications(rating)
        notify! rating, 'create', from: rating.user, to: [rating.page.user] if rating.user != rating.page.user
      end

&amp;#x200B;

In another area of the website, if an author will create a new page, the followers of the other pages created by the same author, will receive a notification about this new work

      def after_create(page)
        notify_followers page
        notify_staff page, 'create'
        assign_score(page.user, :create)
      end
    
      def notify_followers(page)
        notify! page, 'create', from: page.user, to: page.user.followers.to_a, scope: :follower
      end

&amp;#x200B;

But there is a BUG. In this way if the author is in the "followers list", he will receive a notification about it.

How to "exclude" the notification to the page author?

In this case I have `to: page.user.followers.to_a`, and not a *single user* as in rating area. **How to do?**
## [5][Beginner: Database design advice](https://www.reddit.com/r/rails/comments/j93bhi/beginner_database_design_advice/)
- url: https://www.reddit.com/r/rails/comments/j93bhi/beginner_database_design_advice/
---
So, I am scraping an ecommerce website and trying to save product name and prodct price to the pg db.  


I have to keep a daily record of price of each product as well.  


So, how should the database look like?
## [6][Newrelic APM tools for Rails APP](https://www.reddit.com/r/rails/comments/j92pdz/newrelic_apm_tools_for_rails_app/)
- url: https://www.reddit.com/r/rails/comments/j92pdz/newrelic_apm_tools_for_rails_app/
---
Does anyone know any good documentation or users of new relic abilities for a Rails app? All the videos on new relic are from at least 5 years ago.   


Could someone help me with how new relic is used in their application?
## [7][Return Forms from backend ?](https://www.reddit.com/r/rails/comments/j957e4/return_forms_from_backend/)
- url: https://www.reddit.com/r/rails/comments/j957e4/return_forms_from_backend/
---
Hello folks, any one tried to return forms as json and in frontend just check the type of field if it is an file input or a text area or a text , number ? any one tried that before ? any guide ?

backend rails 

frontend using react
## [8][Multi threaded uploads with active storage + s3](https://www.reddit.com/r/rails/comments/j8p4pz/multi_threaded_uploads_with_active_storage_s3/)
- url: https://www.reddit.com/r/rails/comments/j8p4pz/multi_threaded_uploads_with_active_storage_s3/
---
Hi,I am trying to upload multiple small files to s3 from active storage. I have about 4000 files to upload each with less than 1 mb. I am trying to upload .ts files encoded from ffmpeg to s3. This is how I am attaching. I loop through the folder and attach the files. I noticed that uploading the files takes longer than encoding. Will it make better with multiple threads uploading?. Is that also possible?. It took about 1.5 hours to encode but took about 6 hours to upload. Initially, the uploads are fast but it become slower.  
`video.hls_parts.attach(io:File.open(File.join(dir, file)), filename: File.basename(file))`
## [9][How does twitter use rails for its frontend?](https://www.reddit.com/r/rails/comments/j8ktha/how_does_twitter_use_rails_for_its_frontend/)
- url: https://www.reddit.com/r/rails/comments/j8ktha/how_does_twitter_use_rails_for_its_frontend/
---
I keep hearing, twitter still uses rails for its frontend. Does this mean twitter in its entirety is comprised of rails views and the controllers call their other non rails rest APIs (eg: scala) ?
## [10][A Rails bug that makes me question my own sanity and think about deleting my repo. Advice appreciated.](https://www.reddit.com/r/rails/comments/j8nibh/a_rails_bug_that_makes_me_question_my_own_sanity/)
- url: https://www.reddit.com/r/rails/comments/j8nibh/a_rails_bug_that_makes_me_question_my_own_sanity/
---
Hello and thanks for reading. I've been trying to fix this for the last week. I think I've made some progress in understanding whats going on, but I need some help in finding the right solution. 

**THE ISSUE**

Whenever I submit a POST with a rails form and DO NOT redirect afterwards, I get a strange popup on mobile browsers which on safari looks like this. (This is only an issue on phones, the desktop is fine):

[Why does this happen when I submit a rails form and don't render something?](https://preview.redd.it/oyz2fb6njas51.png?width=750&amp;format=png&amp;auto=webp&amp;s=f6f0358242d11c1ef6cbf98fc7d930eddac35c0f)

**MY CODE**

Here i'm using javascript to submit a POST form. Here is the form tag:

    &lt;%= form_with(url: "/stock/createOrderBuy", method: "post", id: "buyForm", html: { class: "needs-validation", novalidate: true }) do %&gt;

On the front end it renders to this:

    &lt;form id="buyForm" class="needs-validation" novalidate="novalidate" action="/stock/createOrderBuy" accept-charset="UTF-8" data-remote="true" method="post"&gt;

Once the the user has hit a confirm button, I use JQuery to submit the form like this:

    $("#buyForm").submit();

Heres the createOrderBuy function its posting to. It just receives the post and returns or renders nothing. The user is kept on the same page. 

    def createOrderBuy
    		
    		shares = params[:shares_buy][:number]
    		type = params[:buyType]
    		limit = 0 
    		if params[:limitBuy][:limitPrice] != ""
    			limit = params[:limitBuy][:limitPrice].to_f
    		end
    		stop = 0
    		if params[:stopBuy][:stopPrice] != ""
    			stop = params[:stopBuy][:stopPrice] .to_f
    		end
    
    
    		order = Order.create(typeId: type, quantity: shares, symbol: params[:symbolHidden], portfolioId: params[:portfolioIdHidden], processed: false, limit: limit, stop: stop, crypto: false, company_name: params[:companyHidden])
    		puts "typeid = " + type.to_s
    
    		if Stock.isMarketOpen &amp;&amp; type.to_s == "1"
    
    				quoteObj = StockQuote::Stock.new(api_key: 'APIKEYHERE')
    				stock = StockQuote::Stock.quote(order.symbol)
    				latestPrice = stock.latest_price
    				orderCost = (stock.latest_price.to_f*shares.to_f).round(2)
    				order.cost = orderCost
    				order.save
    				#this will go in a processor later
    				Stock.buyStock(latestPrice, order)
    		end
    
    	end

**THEORIES**

I posted about this issue four days ago and received some good advice. /u/signalmage suggested that there could be an issue with the content type headers, that this createOrderBuy needs to include a response header like:

    response.headers['Content-Type'] = 'application/vnd.api+json'

I've tried several content type options though, and still get the popup. That said, I'm not sure which type exactly I should be using. 

/u/JustinCampbell offered that when you don't explicitly state a response, its like using 

    head :no_content

Both head :created and head :ok allow me to submit the form without causing a popup! WOOOO! However - both redirect the user to an empty page at /stock/createOrderBuy, stopping the experience totally. I don't want to redirect the user at all after submitting. This behavior really surprised me. 

So thats the issue. If you made it this far then thank you. I appreciate all advice and ideas.
## [11][Generating multi-page pdf reports consisting of headers and tables](https://www.reddit.com/r/rails/comments/j8b6fb/generating_multipage_pdf_reports_consisting_of/)
- url: https://www.reddit.com/r/rails/comments/j8b6fb/generating_multipage_pdf_reports_consisting_of/
---
What do you guys use to accomplish this? I looked into wicked\_pdf but generating 30 pages took way too long. Am I doing it wrong because I was forced to use combine\_pdf gem to combine 30 files to one. Should I be using Prawn gem instead? Is there another method?
## [12][How to hook standalone graphiql client to Rails routes?](https://www.reddit.com/r/rails/comments/j8gvig/how_to_hook_standalone_graphiql_client_to_rails/)
- url: https://www.reddit.com/r/rails/comments/j8gvig/how_to_hook_standalone_graphiql_client_to_rails/
---
Graphiql-Rails hasn't been updated for almost two years and it's missing quite some features (like request header). 

I downloaded the standalone graphiql client, but having trouble how to hook it up to the routes. Do I need to define some additional routes in rails for it to work? Currently I only have 

`post "/graphql", to: "graphql#execute"`
