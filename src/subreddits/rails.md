# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/ep2dw9/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/ep2dw9/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/evmx0w/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/evmx0w/personal_projects_show_off_your_own_project_andor/
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
## [3][One authentication system to cover all modern methods?](https://www.reddit.com/r/rails/comments/ewt1nd/one_authentication_system_to_cover_all_modern/)
- url: https://www.reddit.com/r/rails/comments/ewt1nd/one_authentication_system_to_cover_all_modern/
---
I have a react app with rails backend which needs to support the following:

1. all the devise stuff like sign up, sign in, sign out, confirmation via email, forget password
2. Need a way to authenticate\_user! based on JWT Tokens
3. Support for refresh\_token so that we can keep the user logged in
4. External authentication systems like Login via Google and Facebook

As I understand these 4 requirements are pretty standard for a modern web app. I was surprised to see the lack of a single out of the box solution for it. Here's what I know.

* Devise is necessary. We need something that extends it to do the rest.
* Omniauth plugin for devise covers 4
* devise-jwt covers 2 but not 3
* doorkeeper covers 2 and 3, but struggles with 4 and steps on devise's shoes in places.

Does anyone have any suggestions to consider?
## [4][Is it possible to use Rails 6 with SQL Server (MSSQL)?](https://www.reddit.com/r/rails/comments/ewx2q2/is_it_possible_to_use_rails_6_with_sql_server/)
- url: https://www.reddit.com/r/rails/comments/ewx2q2/is_it_possible_to_use_rails_6_with_sql_server/
---
We were hoping to jump to the latest version of rails to rectify a bunch of identified vulnerabilities.  This Rails code is heavily hand edited, so if necessary, switching from tiny_tds would not be unthinkable, but the data we are fetching is, and will always be in an MS SQL database.  Must we use rails 5.x?
## [5][perform the follower or follower function](https://www.reddit.com/r/rails/comments/ex1mjg/perform_the_follower_or_follower_function/)
- url: https://www.reddit.com/r/rails/comments/ex1mjg/perform_the_follower_or_follower_function/
---
Good afternoon a question how can I perform the follower or follower function for my project, a tutorial or documentation

Thank you
## [6][Fonts haven't been changed.](https://www.reddit.com/r/rails/comments/ewshzl/fonts_havent_been_changed/)
- url: https://www.reddit.com/r/rails/comments/ewshzl/fonts_havent_been_changed/
---
A long time ago, I made a font-api, and I wanted to use that in my current rails project. So, I tested two scenarios and non of them worked. 

1. Font API :   
my API works like this, you add this to your html file and then call fonts in CSS stylings.   
&lt;link rel="stylesheet" type="text/css" href="[http://api.azadqalam.ir/fonts?font=FONTNAME](http://api.azadqalam.ir/fonts?font=FONTNAME) /&gt; 
2. This : [https://gist.github.com/anotheruiguy/7379570](https://gist.github.com/anotheruiguy/7379570)   


I did both of them and I performed : 

    rake assets:precompile
    

But I still see system's default font in the project!
## [7][ActionCable (Redis) -&gt; reconnect on redis connection error](https://www.reddit.com/r/rails/comments/ewn1e4/actioncable_redis_reconnect_on_redis_connection/)
- url: https://www.reddit.com/r/rails/comments/ewn1e4/actioncable_redis_reconnect_on_redis_connection/
---
Hello! 

I work on an application which has ActionCable implemented with Redis.

I would like to know whether it is possible to reset the entire ActionCable setup in case the Redis throws Connection error.

I would like to connect to a different Redis url when the connection error occurs. The infrastructure of this Redis service is that there are two Redis databases mirrored, and there are frequent planned maintanance on the service which can cause the shut down of one of the databases. Unfortunately the switch is not solved by the service itself, however it provides us 2 connection urls. These are the ones I would like to switch between automatically.

Is it possible to do with ActionCable.

I know it is possible with ActiveRecord database connection
## [8][How can I convert this polymorphic association and what is the difference in the 2 methods?](https://www.reddit.com/r/rails/comments/ewpf9f/how_can_i_convert_this_polymorphic_association/)
- url: https://www.reddit.com/r/rails/comments/ewpf9f/how_can_i_convert_this_polymorphic_association/
---
Hello, I'm trying to convert a polymorphic association that is written as (for example):

    # models
    class Exam
      has_many :questions, -&gt; { Question.where(question_type: "graded") }, foreign_key: :question_type_id, dependent: :destroy
    end
    
    class Survey
      has_many :questions, -&gt; { Question.where(question_type: "accumulated") }, foreign_key: :question_type_id, dependent: destroy
    end
    
    class Question
      has_many :choices, dependent: :destroy
    end

I'd like to convert this using Rails' polymorphic association, so it'd look probably like:

    # models
    class Exam
      has_many :questions, as: :question_able  
    end
    
    class Survey
      has_many :questions, as: :question_able
    end
    
    class Question
      belongs_to :question_able, polymorphic: true   ## not sure how to name it..
      has_many :choices, dependent: :destroy
    end

My questions are..

1. How would I go about making the proper migration to make this change? I'm confused how I can make the existing fields \`question\_type\` and \`question\_type\_id\`  convert to the rails polymorphic association. I guess in this example, it would then have to change to \`question\_able\_id\` and \`question\_able\_type\`.
2. My models before passed a scope by \`question\_type\` where it could be either "accumulated" or "graded". How can  \`question\_type\` be scoped in the rails polymorphic association?
3. What is the difference between both ways?
4. How do you name the polymorphic association when the class name is sort of vague in terms of what it is? In my app, I have classes that are named following a given industries lingo, so simply adding the suffix "-able" as the polymorphic association name makes it seem weird.
## [9][Never worked directly with Redis, anything I could be doing better in this code?](https://www.reddit.com/r/rails/comments/ewsfgh/never_worked_directly_with_redis_anything_i_could/)
- url: https://www.reddit.com/r/rails/comments/ewsfgh/never_worked_directly_with_redis_anything_i_could/
---
I wrote a throttler in Redis and it seems to work out okay for the most part, but the connections to the Redis server seem to shoot up randomly until it is exceeded and starts throwing errors. I'm wondering if there is a mistake in the way I've written this little library.

    module Throttle
        MAX_WAIT = 20
        INCREMENT = 1
    
        class &lt;&lt; self
            def limit(user, &amp;block)
                wait_for_credit(user)
    
                yield
            end
    
            def wait_for_credit(user)
                redis = Redis.new
    
                key = "rate-limit-#{user.username}"
                current = redis.get(key).to_i
                time_waited = 0
                
                while current != nil &amp;&amp; current &gt;= 2
                    raise StandardError, "Could not get credit in time" and return if time_waited &gt; MAX_WAIT
        
                    time_waited = time_waited + 1
                    sleep(INCREMENT)
    
                    current = redis.get(key).to_i
                end
    
                redis.multi do
                    redis.incr(key)
                    redis.expire(key, 1)
                end
            end
        end
    end

Does this generally look okay, or am I making some dumb mistake somewhere in this code?
## [10][What is everyone using as wrapper for hybrid app?](https://www.reddit.com/r/rails/comments/ewlmre/what_is_everyone_using_as_wrapper_for_hybrid_app/)
- url: https://www.reddit.com/r/rails/comments/ewlmre/what_is_everyone_using_as_wrapper_for_hybrid_app/
---
I've got a working PWA built with Rails and Turbolinks.  It has notifications too (only working on Android).

But now a requirement has arisen that we need to be able to use Bluetooth (both on Android and iOS) so I'm afraid a PWA will no longer suffice.

My first thought was to use the Turbolinks wrappers:
https://github.com/turbolinks/turbolinks-android and https://github.com/turbolinks/turbolinks-ios but the one for Android has been deprecated for a long time and there don't seem to come any updates.

What would be good alternatives?  Has anyone tried Cordova in combination with Rails (and Turbolinks)?  Any other suggestions.  Mind you: it's not a single page app, it's a normal Rails app, that currently also works as a PWA.

Curious to hear what others are using.
## [11][How to do a rest consult for api with a keystore](https://www.reddit.com/r/rails/comments/ewq8c1/how_to_do_a_rest_consult_for_api_with_a_keystore/)
- url: https://www.reddit.com/r/rails/comments/ewq8c1/how_to_do_a_rest_consult_for_api_with_a_keystore/
---
Hello guys Im trying to do a connection with a key but i'm not sure how to use the keystore in ruby, can you help me to understand their use, one example with rest and keystore can be very good for me, thanks
## [12][Server failed to process the image](https://www.reddit.com/r/rails/comments/ewtd7w/server_failed_to_process_the_image/)
- url: https://www.reddit.com/r/rails/comments/ewtd7w/server_failed_to_process_the_image/
---
I was trying to upload an image for a product using solidus and I keep getting the same error which is: " Server failed to process the image". Does anyone know how to fix this?
