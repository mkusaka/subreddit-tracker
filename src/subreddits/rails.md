# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gek2rk/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gek2rk/personal_projects_show_off_your_own_project_andor/
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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gnbebg/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gnbebg/personal_projects_show_off_your_own_project_andor/
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
## [3][Rails, and David Graeber's book "Bullshit Jobs"](https://www.reddit.com/r/rails/comments/goeib9/rails_and_david_graebers_book_bullshit_jobs/)
- url: https://www.reddit.com/r/rails/comments/goeib9/rails_and_david_graebers_book_bullshit_jobs/
---
For those who have read the book (others can read wikipedia summary, it is enough for this discussion), what do you think about the 3rd kind of bullshit job : "duct tapers". The example given in the book states that programmers are happy to code open-source projects for free, but paid jobs consist in aggregate those free projects, who are not made to work together. Thus, programmer job is not very interesting because it is just integrating tools, the interesting part is left to "free time hobby".  
For me, Ruby-on-Rails is precisely built against this bias, this is why Rails fans are **both** happy and paid  for their job. Why the majority of the industry accept the "duct tapers" as a mainstream remains a mystery to me.
## [4][Capistrano Deploy and Active Storage issue](https://www.reddit.com/r/rails/comments/go4amp/capistrano_deploy_and_active_storage_issue/)
- url: https://www.reddit.com/r/rails/comments/go4amp/capistrano_deploy_and_active_storage_issue/
---
When deploying using Capistrano, I lose all of my assets stored in Rails Active Storage. An older question on this sub (link below) addressed this, however I'm still experiencing the issue after using the fix. 

My deploy.rb has 'public/uploads'  in `append :linked_dirs` . Whenever I run `cap production deploy`, all of my Active Storage photos are wiped.

&amp;#x200B;

Similar question from a year ago:

[https://www.reddit.com/r/rails/comments/bn0x19/capistrano\_deploy\_and\_active\_storage\_question/](https://www.reddit.com/r/rails/comments/bn0x19/capistrano_deploy_and_active_storage_question/)
## [5][Using Rails 4 and Angular - Angular is turning my string into a number? HALP](https://www.reddit.com/r/rails/comments/go3hms/using_rails_4_and_angular_angular_is_turning_my/)
- url: https://www.reddit.com/r/rails/comments/go3hms/using_rails_4_and_angular_angular_is_turning_my/
---
I am new to Angular and pretty new to rails, and very confused.

On one of my views, I have a simple input checkbox. There is an ng-click on this checkbox which calls this function: toggle(&lt;%= id.to\_s %&gt;) 

**Note**: This problem only happens when the ID exceeds JavaScripts max number, I have confirmed this. Unfortunately, most of the ids I'm dealing with are gigantic and greatly exceed this value, so the appropriate fix seems to be to pass them as strings, this is why I'm doing a "to\_s".

The toggle function in my view then calls my Javascript controller in my public/javascripts:

$scope.toggle = function(id)

This function has some more logic which does document.getElementById, and more things to flip the toggle. THIS WORKS for all integer values which fit in JavaScripts "number" type.

However, for some reason that is beyond me, passing in the id as a STRING to the toggle function in the view always seems to immediately convert it to a NUMBER when it enters the toggle javascript function. 

I am just so confused and have spent hours and hours on this problem. If I straight up pass in a massive string in my view like toggle('2283842384328374297439287372839') THIS WILL WORK and it will stay a string in Javascript.

So why is passing it with to\_s in my view NOT working?

Please help me. Thanks.
## [6][New to JSON and GraphQL. Which version to store in db?](https://www.reddit.com/r/rails/comments/gnz8qd/new_to_json_and_graphql_which_version_to_store_in/)
- url: https://www.reddit.com/r/rails/comments/gnz8qd/new_to_json_and_graphql_which_version_to_store_in/
---
As the title shows, I am new to both JSON and GraphQL. Currently I am messing around with the shopify\_api and as such I am using their GraphQL API to gather order data. However, I am a tad bit confused as to how I should be storing this data in my postgres db. I am currently storing requests as t.jsonb :payload

&amp;#x200B;

For the example queries below, my question is which version should I be storing in my database, and why?

      SHOP_NAME_QUERY = client.parse &lt;&lt;-'GRAPHQL'
          {
        shop {
          name
        }
        orders(first: 100) {
          edges {
            node {
              id
              name
              createdAt
              shippingAddress {
                address1
                address2
                city
                province
                provinceCode
                zip
              }
            }
          }
        }
        }
      GRAPHQL
    
      result = client.query(SHOP_NAME_QUERY)
      result_data = result.data
      result_shop_name = result.data.shop.name
      result_json = result.data.to_json
      result_as_json = result.data.as_json
    
    
      puts "result:         #{result}"
      puts "result shop:    #{result_shop_name}"
      puts "result data:    #{result_data}"
      puts "result json:    #{result_json}"
      puts "result as json: #{result_as_json}"

**Responses from the above:**

    result:         #&lt;GraphQL::Client::Response:0x00007f814c131620&gt;
    result shop:    storename
    result data:    #&lt;#&lt;Class:0x00007f8149fc1af8&gt;:0x00007f814c131710&gt;
    result json:
    
    {
      "data": {
        "shop": {
          "name": "shopname"
        },
        "orders": {
          "edges": [
            {
              "node": {
                "id": "gid://shopify/Order/1234567890",
                "name": "#1001",
                "createdAt": "2020-05-01T18:46:04Z",
                "shippingAddress": {
                  "address1": "1234 Elm St",
                  "address2": "",
                  "city": "Chicago",
                  "province": "Illinois",
                  "provinceCode": "IL",
                  "zip": "12345"
                }
              }
            }
          ]
        }
      },
      "casted_data": {
        "shop": {
          "data": {
            "name": "shopname"
          },
          "casted_data": {
            "name": "shopname"
          },
          "errors": [
            
          ]
        }
      },
      "errors": [
        
      ]
    }
    
    result as json:
    
    {
      "data"=&gt;{
        "shop"=&gt;{
          "name"=&gt;"storename"
        },
        "orders"=&gt;{
          "edges"=&gt;[
            {
              "node"=&gt;{
                "id"=&gt;"gid://shopify/Order/1234567890",
                "name"=&gt;"#1001",
                "createdAt"=&gt;"2020-05-01T18:46:04Z",
                "shippingAddress"=&gt;{
                  "address1"=&gt;"1234 Elm St",
                  "address2"=&gt;"",
                  "city"=&gt;"Chicago",
                  "province"=&gt;"Illinois",
                  "provinceCode"=&gt;"IL",
                  "zip"=&gt;"12345"
                }
              }
            }
          ]
        }
      },
      "casted_data"=&gt;{
        "shop"=&gt;{
          "data"=&gt;{
            "name"=&gt;"storename"
          },
          "casted_data"=&gt;{
            "name"=&gt;"storename"
          },
          "errors"=&gt;[
            
          ]
        }
      },
      "errors"=&gt;[
        
      ]
    }
## [7][ActiveJob::SerializationError on create model for uploading](https://www.reddit.com/r/rails/comments/gnwl80/activejobserializationerror_on_create_model_for/)
- url: https://www.reddit.com/r/rails/comments/gnwl80/activejobserializationerror_on_create_model_for/
---
i am trying to do background job for uploading pics ,

the job something like this:

```
  def perform(model, new_image, file)
    model.meta_gallery.attach(
          io: new_image, 
          filename: file.original_filename, 
          content_type:   file.content_type
     )
  end
```

what I got is `ActiveJob::SerializationError (Unsupported argument type: Tempfile):`

I called that job in my controller like this:
```
 params[:gallery][:files].each do |file|
      new_image = helpers.upload_resized_image(file.tempfile, 400, 400)
      UploadJob(@model, new_image, file)
 end

```
I tried to change: file.tempfile.path as well but doesn't work
does anyone has same issue before like this?
## [8][Any free Rails hosting options that don't use Postgres?](https://www.reddit.com/r/rails/comments/gnwlmf/any_free_rails_hosting_options_that_dont_use/)
- url: https://www.reddit.com/r/rails/comments/gnwlmf/any_free_rails_hosting_options_that_dont_use/
---
I simply cannot get Postgres to play ball on Windows 10, and I need a free hosting option for pitching/testing purposes. I know it's far from ideal, but are there any free hosting services that will accept Rails with Sqlite?
## [9][Authenticate wordpress users using Devise](https://www.reddit.com/r/rails/comments/gnvfiu/authenticate_wordpress_users_using_devise/)
- url: https://www.reddit.com/r/rails/comments/gnvfiu/authenticate_wordpress_users_using_devise/
---
I have a Rails API and I'm currently building a wordpress site.  When users sign in on the wordpress site, I'd like to authenticate their credentials using Devise and my Rails API.  Are there any solutions for this type of authentication?  Wordpress offers plugins that utilize [OAuth single sign on](https://wordpress.org/plugins/miniorange-login-with-eve-online-google-facebook/) but I'm not sure how to integrate that with Devise.  

Any ideas on how to solve this?
## [10][Give your Rails app a fast GraphQL API without writing any code.](https://www.reddit.com/r/rails/comments/gni988/give_your_rails_app_a_fast_graphql_api_without/)
- url: https://www.reddit.com/r/rails/comments/gni988/give_your_rails_app_a_fast_graphql_api_without/
---
I run a popular Rails App (on App Engine) it gets significant traffic (movnorth.com). Recently I started introducing React into the frontend and also decided to start using GraphQL in the backend for queries. This started me down a rabbit hole and here is the result.

I built Super Graph an automatic GraphQL to SQL translator (in GO). Just run it along side your Rails app it will learn your database schema, relationships and allow you to query your database using just simple GraphQL.

Super Graph understands Rails cookies and works with session stores to get the current authenticated users ID. 

Github:
https://github.com/dosco/super-graph

Documentation:
https://supergraph.dev/docs/home
## [11][You must use Bundler 2 or greater with this lockfile](https://www.reddit.com/r/rails/comments/gnr5nk/you_must_use_bundler_2_or_greater_with_this/)
- url: https://www.reddit.com/r/rails/comments/gnr5nk/you_must_use_bundler_2_or_greater_with_this/
---
I am trying to deploy an application on Linux Debian 9 and I am getting an error  
`You must use Bundler 2 or greater with this lockfile`  
`Ruby -v returns ruby 2.6.3p62`  
`rails -v returns Rails` [`6.0.3.1`](https://6.0.3.1)  
`which bundle returns /usr/local/rvm/gems/ruby-2.6.3/bin/bundle`  
`bundle version returns Bundler version 2.1.4 (2020-01-05 commit 32a4159325)`

So, I have tried various things such as

1. Updating the bundler gem(gem install bundler -v 2.1.4)
2. Updating the rubygems Package manager( gem update --system)
3. Reinstalling bundler gem

but they dont seem to work and  everytime I try to work it throws the same error.  
Anyone who has been facing the same error and could help?
## [12][Image upload.](https://www.reddit.com/r/rails/comments/gnkn0e/image_upload/)
- url: https://www.reddit.com/r/rails/comments/gnkn0e/image_upload/
---
Hit a block here. I've handled image uploading before via action text, but not as an element in a db that I can call back. Wondering, would I just add the params/permit and migration into the post table? If so, what would be the syntax for the migration? And as far as views go, it's pretty straight forward via simple\_for &amp; a framework right?

It's for a post/index page which will have a title/description linking to a full post. Hoping I can get a db-driven image since I don't want them all to have a static image. Will be 1 square image per post for the card.
