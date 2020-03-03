# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/f9t9kq/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/f9t9kq/personal_projects_show_off_your_own_project_andor/
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
## [2][Do I need to get user info before checking user data with Devise/Rails + React?](https://www.reddit.com/r/rails/comments/fckvp3/do_i_need_to_get_user_info_before_checking_user/)
- url: https://www.reddit.com/r/rails/comments/fckvp3/do_i_need_to_get_user_info_before_checking_user/
---
Greetings folks,

I'm new to Devise, and so far it's been really cool to learn and work with. I'm building a Rails app that has a React front-end, and I'm developing the API for the front-end to check if a logged in user is an Admin.

I have this React code that hits a "is_admin method"


       componentWillMount(){
         this.checkIfUserIsAdmin();
       }

       checkIfUserIsAdmin() {
         const url = "http://localhost:3000/users/is_admin";

         const token = document.querySelector('meta[name="csrf-token"]').content;
         console.log(token)
         fetch(url, {
           method: "GET",
           headers: {
             "X-CSRF-Token": token,
             "Content-Type": "application/json"
           }
         })
           .then(response =&gt; {
             if (response.ok) {
               return response.json();
             }
             throw new Error("Network response was not ok.");
           })
           .catch(error =&gt; console.log(error.message));
           //eventually this will update the component state depending on the response
       }

The API seems to be hitting the proper end point just fine, but I am getting the response from the server:

     Started GET "/users/is_admin" for ::1 at 2020-03-02 14:59:05 -0800
     Processing by SessionsController#is_admin as */*
       Parameters: {"session"=&gt;{}}
     Completed 500 Internal Server Error in 2ms (ActiveRecord: 0.0ms | Allocations: 1235)


  
     NoMethodError (undefined method `admin?' for nil:NilClass):

This seems to indicate that there is no"user" for Devise to check if admin on. This would make sense as I'm not passing anything to ID the user to the Rails method from React. I was under the impression that Devise would manage the session, and that passing the CSRF token would be enough.

What would be the proper way to work this out? Should I first get the logged in user's data, and then set @user to the result of that?
## [3][ActiveModel::Errors question](https://www.reddit.com/r/rails/comments/fct5m5/activemodelerrors_question/)
- url: https://www.reddit.com/r/rails/comments/fct5m5/activemodelerrors_question/
---
From [https://medium.com/@rfleury2/a-quick-guide-to-model-errors-in-rails-965e2be3ac93](https://medium.com/@rfleury2/a-quick-guide-to-model-errors-in-rails-965e2be3ac93) this article I know that once a model is derived from `ActiveRecord::Base` it will have access to `ActiveModel::Errors` object which gets populated in case of validation or if we manually insert into it using `class.errors.add()` syntax.I want to know if there are any other ways that `ActiveModel::Errors` will be populated? Like when there is a rollback when that model is being saved and so on. Or when `update_attributes` is called on the model and so on
## [4][Best practices for long-term API auth tokens?](https://www.reddit.com/r/rails/comments/fcopoz/best_practices_for_longterm_api_auth_tokens/)
- url: https://www.reddit.com/r/rails/comments/fcopoz/best_practices_for_longterm_api_auth_tokens/
---
I'm working on an Android app that will be connected to a Rails backend. I would like the user to sign in once for the entire duration that they use the app - potentially for years. Because of this, I want to use sessions, so that access can be revoked. My thinking was to use a JWT that would contain a session ID, which would correspond to a column in a session table in the database. There would also be a 'is_valid' attribute, so that the session could be disabled.

Is that a good/safe/smart way to do what I need? I've been researching this for a few hours and can't find much on it. I've seen that some people will have a short-lived token that refreshes itself, but if the user doesn't use the app for a few days they would be logged out.
## [5][How do I persist previous params to a new request?](https://www.reddit.com/r/rails/comments/fcl8oc/how_do_i_persist_previous_params_to_a_new_request/)
- url: https://www.reddit.com/r/rails/comments/fcl8oc/how_do_i_persist_previous_params_to_a_new_request/
---
Hi all,

So I have a button which will order my search results by rating:

    &lt;%= link_to 'Top rated', coffee_shops_path(:order_by =&gt; "rating") %&gt;

This goes to my coffee\_shops#index controller, where I call another method if an order\_by param is given, called order\_coffee\_shops\_by\_param:

    def order_coffee_shops_by_param(coffee_shops, param)
      if param == "rating"
        @coffee_shops = @coffee_shops.order("#{param} DESC NULLS LAST")
      end
    end

The problem is, I'm losing all my other params in this process. Everything that was already in the URL string is being lost. So rather than this just reorganising all the results that are already on the page, it's doing a clean search, returning all results, and then ordering these results. My params look like this.

    &lt;ActionController::Parameters {"order_by"=&gt;"rating", "controller"=&gt;"coffee_shops", "action"=&gt;"index"} permitted: false&gt;

I've tried doing this using [this stackoverflow answer](https://stackoverflow.com/questions/2695538/add-querystring-parameters-to-link-to) but I just get the error:

    unable to convert unpermitted parameters to hash

Help appreciated as always.
## [6][What is the fastest JSON serializer, excluding those that follow the JSON API standard?](https://www.reddit.com/r/rails/comments/fcmw93/what_is_the_fastest_json_serializer_excluding/)
- url: https://www.reddit.com/r/rails/comments/fcmw93/what_is_the_fastest_json_serializer_excluding/
---
At my job we will make an API (no JSON API, just old school REST) and a Front-end app to break down an old web app we have. I want it to be as fast and less memory consuming as possible. From what I've researched I think [panko_serializer](https://github.com/yosiat/panko_serializer) and [blueprinter](https://github.com/procore/blueprinter) are both solid choices. But I don't know if there's another one that's even better.

Greetings.
## [7][Large react app with webpacker, experiences and opinions wanted](https://www.reddit.com/r/rails/comments/fchdel/large_react_app_with_webpacker_experiences_and/)
- url: https://www.reddit.com/r/rails/comments/fchdel/large_react_app_with_webpacker_experiences_and/
---
Folks,

I have a pretty large React app that is backed by a relatively vanilla Rails app. A few years ago, I converted it to use Webpacker from the very first day it was released. I've been keeping with upgrading it, but feel that its getting problematic due to the size of my application. 

In terms of UI, there are about 10 separate packs files which are all individual react applications that map to different sections of the site. It seems that compile times have become slower and slower as the number of packs has grown. Additionally, on every save, all packs are compiled even though changes to a single one are made. Is there no way around this? I've been able to setup some code splitting so that everything coming from node\_modules is placed into a separate chunk, but things still seem slow during development.

My question is... has anyone here with large client side apps experienced the same...and what did you do? Should i be looking to ditch webpacker and just start hand rolling some optimized custom webpack configuration solution?  I personally hate dealing with webpack config and would love to avoid it, but not sure if that is just a pipedream.  Do any of you have examples of webpack / webpacker configurations that could be useful for me to look through? Thanks all
## [8][Rails / React app fails in production - parsing errors](https://www.reddit.com/r/rails/comments/fchkzz/rails_react_app_fails_in_production_parsing_errors/)
- url: https://www.reddit.com/r/rails/comments/fchkzz/rails_react_app_fails_in_production_parsing_errors/
---
I'm having a strange issue where my Rails 6/React app renders fine in development using \`bin/webpack-dev-server\` but fails in production using \`RAILS\_ENV=production assets:precompile\` and \`rails s -e production\`. It's as if the JSX no longer is recognized and view as a syntax error. 

&amp;#x200B;

    11 |   render() {
      12 |     return (
    &gt; 13 |       &lt;&gt;
         |       ^
      14 |         {this.props.recurringOutages &amp;&amp; this.props.recurringOutages.map(recurringOutage =&gt; (
      15 |           &lt;Card.Group key={uuidv4()}&gt;
      16 |             &lt;Card
        at Parser.raise (/Users/demiansims/Development/tml_info/tml_dashboard/node_modules/@babel/parser/lib/index.js:7044:17)
        at Parser.unexpected (/Users/demiansims/Development/tml_info/tml_dashboard/node_modules/@babel/parser/lib/index.js:8422:16)
        at Parser.parseExprAtom (/Users/demiansims/Development/tml_info/tml_dashboard/node_modules/@babel/parser/lib/index.js:9701:20)
        at Parser.parseExprSubscripts (/Users/demiansims/Development/tml_info/tml_dashboard/node_modules/@babel/parser/lib/index.js:9287:23)
        at Parser.parseMaybeUnary (/Users/demiansims/Development/tml_info/tml_dashboard/node_modules/@babel/parser/lib/index.js:9267:21)
        at Parser.parseExprOps (/Users/demiansims/Development/tml_info/tml_dashboard/node_modules/@babel/parser/lib/index.js:9137:23)
        at Parser.parseMaybeConditional (/Users/demiansims/Development/tml_info/tml_dashboard/node_modules/@babel/parser/lib/index.js:9110:23)
        at Parser.parseMaybeAssign (/Users/demiansims/Development/tml_info/tml_dashboard/node_modules/@babel/parser/lib/index.js:9065:21)
        at Parser.parseParenAndDistinguishExpression (/Users/demiansims/Development/tml_info/tml_dashboard/node_modules/@babel/parser/lib/index.js:9842:28)
        at Parser.parseExprAtom (/Users/demiansims/Development/tml_info/tml_dashboard/node_modules/@babel/parser/lib/index.js:9622:21)
    ℹ ｢wdm｣: Failed to compile.
## [9]["VFS Connection does not exist" on Cloud9](https://www.reddit.com/r/rails/comments/fcdlpf/vfs_connection_does_not_exist_on_cloud9/)
- url: https://www.reddit.com/r/rails/comments/fcdlpf/vfs_connection_does_not_exist_on_cloud9/
---
As the title says, when i try to run the $rails server command and   
preview running application, it tells me "Oops, VFS Connection does not   
exist". Any way to fix this?
## [10][match url to a rack app](https://www.reddit.com/r/rails/comments/fcelj3/match_url_to_a_rack_app/)
- url: https://www.reddit.com/r/rails/comments/fcelj3/match_url_to_a_rack_app/
---
When I match a url to a rack app proc in the router (without controller), how do I get the url/form params?
## [11][Does anyone build build their rails API using the JSON API standard? What server side library do you use?](https://www.reddit.com/r/rails/comments/fbzmuj/does_anyone_build_build_their_rails_api_using_the/)
- url: https://www.reddit.com/r/rails/comments/fbzmuj/does_anyone_build_build_their_rails_api_using_the/
---
The [JSON API](https://jsonapi.org/) standard is a general standard for formatting API responses for your JSON API. It's one of many standards, but I kind of like the simplicity of it and it seems fairly popular. 

Their website mentions [several ruby/rails gems](https://jsonapi.org/implementations/#server-libraries-ruby) that implement the standard. All of them promise that you can just add them and it generates links, routes, responses, etc... for you. 

I looked through a few of them and found the amount of choices overwhelming, and a few of them weren't really maintained that well. 

Do any of you use the JSON API standard for your own APIs, and if so what would you recommend? 

Thanks!
