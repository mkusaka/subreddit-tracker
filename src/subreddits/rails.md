# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gvu083/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gvu083/personal_projects_show_off_your_own_project_andor/
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
## [2][Quick Demo of Reactive UI with Stimulus Reflex](https://www.reddit.com/r/rails/comments/h8j14f/quick_demo_of_reactive_ui_with_stimulus_reflex/)
- url: https://www.reddit.com/r/rails/comments/h8j14f/quick_demo_of_reactive_ui_with_stimulus_reflex/
---
[https://youtu.be/MlnNkcz-oIc](https://youtu.be/MlnNkcz-oIc)
## [3][Out of curiosity, how many of you are finding that your professional Rails-based environment is moving more toward JavaScript?](https://www.reddit.com/r/rails/comments/h8mdaz/out_of_curiosity_how_many_of_you_are_finding_that/)
- url: https://www.reddit.com/r/rails/comments/h8mdaz/out_of_curiosity_how_many_of_you_are_finding_that/
---
The reason I ask this is because I am - and have been for quite some time - a professional JavaScript developer.

However, I love working with Rails and am positioning myself to move into my first Rails gig within the next 6 months.

And I had a very interesting conversation with a pro Rails developer the other day.  This guy is a senior engineer who has been doing Rails for over 12 years.  Recently, though, he was laid off from a remote job due to his company taking a big economic hit because of Covid-19.

He told me that he immediately started sending out resumes and went through a battery of tests and interviews and what struck his as interesting is that they were FAR more interested in his JavaScript (React mainly) experience than they were his 12 years of Rails development.  In fact, he said that because of my JS experience, I would probably have an easier time finding a Rails job than he did - which sounds ludicrous. 

Are any of you folks starting to witness the shift as well?
## [4][Calculating average order frequency in days with activerecord](https://www.reddit.com/r/rails/comments/h8pzlo/calculating_average_order_frequency_in_days_with/)
- url: https://www.reddit.com/r/rails/comments/h8pzlo/calculating_average_order_frequency_in_days_with/
---
For every user with more than one order, I want to calculate the average amount of time in between order's `created_at` in days (how often a user places an order):

7th grade math way: Calculate the average by taking the total sum of the numbers divided by the total number count:

`[1,2,3,4].sum / [1,2,3,4].length`

This is a bit more tricky when dealing with dates (more specifically time in seconds). Can I accomplish this with the same level of accuracy using the `.minimum` and `.maximum` activerecord methods?

The only way I've been able to get what appears to be an accurate number is by breaking down the dates by `seconds` and converting to `days`

my code so far (very ugly and expensive):

    def self.average_order_frequency
        frequency_collection = []
        customers = User.joins(:orders)
        .group('users.id')
        .having('count(orders) &gt; 1') 
    
        if customers.any?
          customers.each do |customer|
            orders = customer.orders
    
            seconds = orders.maximum(:created_at) - orders.minimum(:created_at)
            avg_seconds = seconds / (orders.count - 1)
            avg_days = avg_seconds / (24 * 60 * 60)
    
            frequency_collection.push(avg_days)
          end
    
          "#{(frequency_collection.sum / customers.length).round(2)} days in between orders on average"
        else
          "none to show"
        end
      end

I tried to get a day count using a range:

    (order[num+1].created_at.to_date..order[num].created_at.to_date).to_i

but it requires an additional nested block and several more lines in an already ugly method.

How do you calculate the average amount of days in between records creation?
## [5][What is the best way to access the joined model's data in has_may :through?](https://www.reddit.com/r/rails/comments/h8ghu7/what_is_the_best_way_to_access_the_joined_models/)
- url: https://www.reddit.com/r/rails/comments/h8ghu7/what_is_the_best_way_to_access_the_joined_models/
---
I'm creating a standard has\_many :through association like the one outlined in the [rails guide.](https://guides.rubyonrails.org/association_basics.html#the-has-many-through-association)

I have `Users` and `Gardens` joined through `GardenUsers` as shown below.  


    class User &lt; ApplicationRecord
      has_many :garden_users, dependent: :destroy
      has_many :gardens, through: :garden_users
    end
    
    class Garden &lt; ApplicationRecord
      has_many :garden_users, dependent: :destroy
      has_many :users, through: :garden_users
    end
    
    class GardenUser &lt; ApplicationRecord
      belongs_to :user
      belongs_to :garden
      enum garden_role: {owner: 0, maintainer: 1, harvester:2}
    end

My first question:

This seems like the best association because `Users` will have multiple `Gardens` and visa versa, but is there a better way, given that there will only be one connection between any given user/garden? If it is the best, what would be best way to ensure that there is only one instance of any `garden_id/user_id` combination in `garden_users`?

My second question is:  
What is the best way to access a user's `garden_role` for a specific garden? The best way that I could come up with is

    class User &lt; ApplicationRecord
      def garden_role(garden)
        garden_user.find_by(garden_id: garden.id).garden_role
      end
    end

but that feels a bit clunky and seems like a common enough thing that there would be a more "rails magic" way to access whatever data is in the `GardenUser` model.
## [6][Experience with the Shopify API?](https://www.reddit.com/r/rails/comments/h8c3cv/experience_with_the_shopify_api/)
- url: https://www.reddit.com/r/rails/comments/h8c3cv/experience_with_the_shopify_api/
---
Hey - does anyone have experience with the Shopify API? If so, I'd love some help. 

When a Shopify Merchant installs my app (written in RoR), we ask the Merchant to manually add some Liquid code to their Theme files. But I'd like to automate that process and can't figure out how to do it.
## [7][ActiveStorage: How to convert images to WebP properly?](https://www.reddit.com/r/rails/comments/h87nfs/activestorage_how_to_convert_images_to_webp/)
- url: https://www.reddit.com/r/rails/comments/h87nfs/activestorage_how_to_convert_images_to_webp/
---
Hey there,

currently I'm experiment with ActiveStorage and image optimization. To convert my JPGs to WebPs is my main goal at the moment to increase the performance and loading time.

Problem: converting JPG to WebP works out of the box but the generated WebPs acts like JPG (MIME type and file extension)

One of my models has an active storage association: `has_one_attached :profile_picture`

The view:

&amp;#x200B;

    - Photomodel.all.each do |model|
     - if model.profile_picture.attached?
      .pictures{style: "display: flex"}
        = image_tag(model.profile_picture.variant(convert: :webp, resize_to_limit: [300,300]), alt: 'webp')
        = image_tag(model.profile_picture.variant(resize_to_limit: [300,300]), alt: 'jpg')

Actually, rails converts the uploaded JPG to WebP but ignores to change the mime type and the file extension.

https://preview.redd.it/0i7ppk1tjo451.png?width=449&amp;format=png&amp;auto=webp&amp;s=0f8fa1b585203e5e21fd2059cec3315b76dc6101

&amp;#x200B;

https://preview.redd.it/wdosqg84jo451.jpg?width=912&amp;format=pjpg&amp;auto=webp&amp;s=67ab5c4ae4a4a93621d1e7aef55b8b45cc3ff054

With `pry` I tried to inspect the active storage object to change those parameters but found nothing in this direction.

Is it possible to change the parameter of mime type and file extention to 'image/webp' and .webp?

Many thanks!
## [8][Hello Everyone, new rails user, I’ve built a functional CRUD rails app, now I need tips on styling](https://www.reddit.com/r/rails/comments/h8ikx9/hello_everyone_new_rails_user_ive_built_a/)
- url: https://www.reddit.com/r/rails/comments/h8ikx9/hello_everyone_new_rails_user_ive_built_a/
---
I am done with my app, I’ve given it full Crud functionality, but it’s looking very plain. I’ve designed the home page and about page quite nicely, but is there any way to style my views *after* the app is completed? This might sound like a basic and stupid question, but I am new to this. Most of the resources online are for those who are starting from scratch and I am not. Thanks in advance.
## [9][The case of the disappearing db column :(](https://www.reddit.com/r/rails/comments/h7ww81/the_case_of_the_disappearing_db_column/)
- url: https://www.reddit.com/r/rails/comments/h7ww81/the_case_of_the_disappearing_db_column/
---
I have been programming in RoR since version 2.0, but I have never encountered a problem like this til now.  This is my second RoR API with React as the front end app. RoR v6.02

The problem is in one particular table **Invitations** which has, among many others, a string column originally named **state**.  I have many other tables with similar columns (same name, same type) that work fine, BUT, the **Invitations** table is in a world of its own.  


* 1 the column does exist in the table.  I can see it and populate it using
pgAdmin 4.
* 2 I can see the field in the log, with the proper values.  
* 3 I have it in my parameters whitelist.  
* 4 I can see it when printing params.inspect.

In the line immediately following the inspection, as in:

    puts "#{params.inspect}"

= "... updated_at"=&gt;"2020-06-12T17:32:46.420Z", "state"=&gt;"invite"} permitted: true&gt;}

    puts  "#{params.state}"

=NoMethodError (undefined method `state' for &lt;ActionController::Parameters...&gt;)

Any code that executes after the whitelist gives me the same NoMethodError.  For example,

    @state = params.state;
    Invitation.update(invitation_params).

I have looked all over the web for a similar problem but maybe I just don't know the right question.

Lastly, just in case, I renamed the column to **status**, but nothing changed except for the column name in the logs.

Any ideas will be appreciated.

P.S.  I have been working in digital electronics since 1954, with a degree in Electrical Engineering.  The first machine I worked on was a LGP-30 with 8k words of drum memory.  The OS was a graduate student with a cot (and miles of paper tape).  I am probably among the few survivors who knows how to program around drum latency :) /brag
## [10][Rails Routes in Rails 6 looks kinda terrible in the terminal](https://www.reddit.com/r/rails/comments/h7lozw/rails_routes_in_rails_6_looks_kinda_terrible_in/)
- url: https://www.reddit.com/r/rails/comments/h7lozw/rails_routes_in_rails_6_looks_kinda_terrible_in/
---
Does anyone have a solution for the rails routes command that doesn't require grepping a specific controller or using the expand flag, that would pretty up the output of the rails routes command? It seems to be thrown off by action mail and active storage routes because of their length. Sure, I could use the expanded flag but then the console window can go for days. I would like to just see the formatting locked like it used to be or maybe even suppress the action mail stuff. 

https://preview.redd.it/5pyumixzoh451.png?width=1706&amp;format=png&amp;auto=webp&amp;s=c4e08fe5012b36a9f1957842de27ba785805d509
## [11][heroku[logplex]: Error L10 (output buffer overflow): drain 'd.cfcd3264-7259-4b12-ac5b-016032978174' dropped 1 messages since 2020-06-12T14:40:59.280571+00:00.](https://www.reddit.com/r/rails/comments/h7lyqq/herokulogplex_error_l10_output_buffer_overflow/)
- url: https://www.reddit.com/r/rails/comments/h7lyqq/herokulogplex_error_l10_output_buffer_overflow/
---
  
How can I even begin troubleshooting this? it would help to know what the contents of the dropped messages were, or at least which method/process/dyno these dropped messages originated from.
