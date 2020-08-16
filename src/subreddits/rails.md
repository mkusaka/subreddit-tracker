# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/i8dsvv/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/i8dsvv/personal_projects_show_off_your_own_project_andor/
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
## [2][Updating An Existing App?](https://www.reddit.com/r/rails/comments/iapuzt/updating_an_existing_app/)
- url: https://www.reddit.com/r/rails/comments/iapuzt/updating_an_existing_app/
---
Hey guys, just a quick question. I uploaded my initial version of my app, and to be honest it was to mainly see how deployment works.

How do I go about updating the app when I have a major change I want to upload to the server?

For Heroku, I used the CLI to do this, but for servers like Digital Ocean and AWS are different.

Any tips or workflow you guys use will be very helpful!

Thanks.
## [3][Managing usage for a commercial gem](https://www.reddit.com/r/rails/comments/iam6tf/managing_usage_for_a_commercial_gem/)
- url: https://www.reddit.com/r/rails/comments/iam6tf/managing_usage_for_a_commercial_gem/
---
How could I manage usage of a gem or engine as a product?

Paid Sidekiq for example is tied to threads, but I don't know how this is managed. 

It's loaded in the Gemfile so is referenced when bundled, yet Ruby is of course an interpreted and open language so even if it is reading other configs, that can all be manipulated in different ways at runtime, right? Is it more the honor system? (There is always trust involved somewhere, though this is handled to different degrees depending on the product.)

This is not meant to be the 'hack Sidekiq' thread, I am more curious how this could be done with a commercial product, not even with threads but with number of users for instance to prevent abuse.

[I tried Google but didn't find much, it returns whatever it feels like now.]
## [4][How do I configure my models in this situation?](https://www.reddit.com/r/rails/comments/iaov9r/how_do_i_configure_my_models_in_this_situation/)
- url: https://www.reddit.com/r/rails/comments/iaov9r/how_do_i_configure_my_models_in_this_situation/
---
I'm writing an email web app. The app uses email templates to create dynamic emails depending on who they are sent to.

An email template has variables. For example an email template would look like the following:

    EmailTemplate
    
    name: string
    subject: string
    body: string
    variables: [string]

and an instance would look like this:

    email_template
    
    name: "Cool email"
    subject: "Hi, FIRST_NAME,"
    body: "I see you are a POSITION at COMPANY...."
    variables: ["NAME", "POSITION", "COMPANY"]

Then a recipient model to the email would look something like this:

    Recipient
    name: "Some guy"
    email: "guy@email.com"
    variables: ??????
    
The *variables* attribute in EmailTemplate are dynamic so how do I set up the variables on the recipient model?
## [5][Efficient to run SQL command instead of ruby?](https://www.reddit.com/r/rails/comments/iaex30/efficient_to_run_sql_command_instead_of_ruby/)
- url: https://www.reddit.com/r/rails/comments/iaex30/efficient_to_run_sql_command_instead_of_ruby/
---
I've having trouble recreating the trigam `similarity` function in ruby (due to Asian characters in my text). Is the following a good idea in my rails code?

`sql = "SELECT similarity('#{src}', '#{targ}');"`

`tm_ratio = ActiveRecord::Base.connection.execute(sql).getvalue(0,0)`
## [6][TDD Course or tutorial](https://www.reddit.com/r/rails/comments/iaejjw/tdd_course_or_tutorial/)
- url: https://www.reddit.com/r/rails/comments/iaejjw/tdd_course_or_tutorial/
---
Hey y'all, I feel that I have a pretty decent grasp of ruby and rails and I want to learn TDD what are some good resources for learning TDD?

Thanks in advance
## [7][Elasticsearch keyword search orientation.](https://www.reddit.com/r/rails/comments/iaftov/elasticsearch_keyword_search_orientation/)
- url: https://www.reddit.com/r/rails/comments/iaftov/elasticsearch_keyword_search_orientation/
---
Hello, I've landed my first job and I have to do a text search built in elasticsearch (but I can't run anything on db since our db is kind of big (about \~300k lines) the db is kind of slow, the project is run by my boss and me a complete newbie developer, what are some things that I can do, please be kind with me that I'm a mere starter developer.
## [8][What is the difference about a controller test and a integration test?](https://www.reddit.com/r/rails/comments/iab5w3/what_is_the_difference_about_a_controller_test/)
- url: https://www.reddit.com/r/rails/comments/iab5w3/what_is_the_difference_about_a_controller_test/
---
I had looking about a principal difference but I feel it is the same
## [9][New Tutorial Video: Pagination with the Kaminari Gem](https://www.reddit.com/r/rails/comments/iabx3c/new_tutorial_video_pagination_with_the_kaminari/)
- url: https://www.reddit.com/r/rails/comments/iabx3c/new_tutorial_video_pagination_with_the_kaminari/
---
In this tutorial I walk through how to add pagination to your Rails 6 application using the Kaminari gem.  We will install the gem, generate the views, get those views styled via Bootstrap 4.  Finally will paginate our users and videos pages.  

As always please let me know if you see opportunities for me to improve my tutorial videos, thanks! 

[https://youtu.be/c0hrNQJhKww](https://youtu.be/c0hrNQJhKww)
## [10][Is there a way where I can do an specific query with PGCLIENTENCODING="ISO88591"?](https://www.reddit.com/r/rails/comments/iabymz/is_there_a_way_where_i_can_do_an_specific_query/)
- url: https://www.reddit.com/r/rails/comments/iabymz/is_there_a_way_where_i_can_do_an_specific_query/
---
**TL:DR;** My database is in UTF8, but I need to produce a file to be consumed by a COBOL program where 1 character should be 1 bytesize.

**Longer description**: My boss told me to setup a cronjob to periodically make some queries and produce files which should be consumed by a COBOL program, this program accepts file where each is line/record is 200 characters and bytes long (in other words if `line` is a variable which represents any line in the file, then `line.length` should be equal to `line.bytesize` and both should be equal to `200`).

I would say in the database 99.9% of characters are normal ones, no weird ones like Ü or á (or at least I haven't seen them), but to prepare myself in case there are weird characters I would like to query with `PGCLIENTENCODING="ISO88591"`, is this possible?

By the way I'm latino, and I don't care calling `á` a weird character, don't be so annoyed like [these guys](https://www.reddit.com/r/ruby/comments/i9vlu5/replace_weird_characters_such_as_%C3%A1_a_or_%C3%BC_u/).
## [11][How can I pass a list of allowed value with custom each validator?](https://www.reddit.com/r/rails/comments/ia3rxu/how_can_i_pass_a_list_of_allowed_value_with/)
- url: https://www.reddit.com/r/rails/comments/ia3rxu/how_can_i_pass_a_list_of_allowed_value_with/
---
I have been reading the example from [ActiveModel::EachValidator](https://api.rubyonrails.org/classes/ActiveModel/Validator.html) .

    class TitleValidator &lt; ActiveModel::EachValidator
      def validate_each(record, attribute, value)
        record.errors.add attribute, 'must be Mr., Mrs., or Dr.' unless %w(Mr. Mrs. Dr.).include?(value)
      end
    end
    
    class Person
      include ActiveModel::Validations
      attr_accessor :title
    
      validates :title, presence: true, title: true
    end

I understood this example. But how to write the code when I want to pass the list `%w(Mr. Mrs. Dr.)` from client code like: 

    validates :title, title: { in: %w(Mr. Mrs. Dr.) }
