# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/ezrfed/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/ezrfed/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/f2r9mk/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/f2r9mk/personal_projects_show_off_your_own_project_andor/
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
## [3][Question About Railties, rails repo and rails gems](https://www.reddit.com/r/rails/comments/f5i1am/question_about_railties_rails_repo_and_rails_gems/)
- url: https://www.reddit.com/r/rails/comments/f5i1am/question_about_railties_rails_repo_and_rails_gems/
---
I am having a hard time understanding the rails repo. Active record, railties and all others(action view, etc) are in a monorepo. rails/rails. 

They are all separate gems as well? Can someone explain this to me? Each action/active is a gem &amp; available on rubygems. 

Or is the rails repo /activerecord boilerplate to connect to the actual gem? 

Tldr: what's the difference between activerecord gem and the folder that sits in rails repo titled "activerecord"? 

Thanks!

Edit: it looks like the individual gems (activerecord) is not on rubygems. Now I am more confused. When I run "bundler list" it treats each one as a separate gem. Why is this?
## [4][How do you remove a user's personally identifiable information but retain user activity?](https://www.reddit.com/r/rails/comments/f5lafp/how_do_you_remove_a_users_personally_identifiable/)
- url: https://www.reddit.com/r/rails/comments/f5lafp/how_do_you_remove_a_users_personally_identifiable/
---
I'd like to remove Jane Doe's email and full name from the system but retain Jane D.'s activity -- all comments, tasks etc., similar to how Basecamp functions.

Has any of you solved this problem and how did you do it?

Thank you!
## [5][Mocha Stubbing with any parameter](https://www.reddit.com/r/rails/comments/f5m16x/mocha_stubbing_with_any_parameter/)
- url: https://www.reddit.com/r/rails/comments/f5m16x/mocha_stubbing_with_any_parameter/
---
I am trying to stub a method with two parameters which can be anything to return true. I am currently using mocha 0.13.3 and I couldn't find any  help from the documentation here [https://mocha.jamesmead.org/Mocha/ParameterMatchers.html](https://mocha.jamesmead.org/Mocha/ParameterMatchers.html)

    User.any_instance.stubs(:method?).with(:any, :any).return true

So when the method is being called with any two objects it should return true. Could you help me with how to do this?
## [6][Dynamically upload images to active storage inside client](https://www.reddit.com/r/rails/comments/f57u6d/dynamically_upload_images_to_active_storage/)
- url: https://www.reddit.com/r/rails/comments/f57u6d/dynamically_upload_images_to_active_storage/
---
I am trying to setup Active Storage so that users can upload images to their posts. I can get it to work but I want to enhance it a bit:
- when a user adds images to the post form, they should automatically be added to his post body
- This means I need to upload the post images from the client side as they are added to the form

Do you know if this is possible?

Edit: I chose for my post bodies to be formatted with markdown so I cannot use ActionText.

Edit 2: 
Basically my end goal is to be able to add this to the markdown post body when I input an image in the form:
"![alt text](new image url)".
## [7][Documenting documents?](https://www.reddit.com/r/rails/comments/f5cian/documenting_documents/)
- url: https://www.reddit.com/r/rails/comments/f5cian/documenting_documents/
---
I'm working on a self-study project to improve my Rails habits and explore new/alternative ways of doing things. The project is a Jeopardy clone.

So normally when I model my apps, I unconditionally embrace the ActiveRecord pattern.

Following that habit, I would create a Game model which represents a specific game. That game would has\_and\_belongs\_to\_many categories, since a game consists of 6 (randomly chosen) categories and a category can be recycled for many games. Each category has\_many answers (remember, this is Jeopardy). And each answer has\_many questions, with only one of the questions being correct. I understand that in the TV-version, the question is basically a free text field, which the host has to process to determine correctness, but in my version of the game, the player will be able to choose one of five suggested questions, with 1 being correct and the 4 other being invalid.

Since my game will driven by Vue JS on the client side and Vue will be loading the entire round (6 categories including answers and questions) in one operation, it would like to explore the document approach to modelling.

So instead of making a handful of tables and a bunch of ActiveRecord objects that will inevitably turn into a nested JSON structure to be delivered to Vue, I would like to just have one model called a Category, which contains a json array of answers which in turn contains an array of questions.

Now, doing this is not hard. But I think it makes the application less transparent, because schema.rb won't document the data structure and neither will the ActiveRecord models and their defined relationships.

So my question is: How would you go about documenting the data structure of these json columns? So I don't end up with defensive code, because I forget the (implicit) data schema?
## [8][is it possible to add emoji using redcarpet gem for the markdown?](https://www.reddit.com/r/rails/comments/f5cv6l/is_it_possible_to_add_emoji_using_redcarpet_gem/)
- url: https://www.reddit.com/r/rails/comments/f5cv6l/is_it_possible_to_add_emoji_using_redcarpet_gem/
---

## [9][Multi-tenant with Apartment](https://www.reddit.com/r/rails/comments/f546p9/multitenant_with_apartment/)
- url: https://www.reddit.com/r/rails/comments/f546p9/multitenant_with_apartment/
---
I've started to play with the Apartment gem hoping to modify an existing app to go this route for multi-tenant.

I'm familiar with Devise so I'm able to get the user to login and redirected to their subdomain and create records in their own schema.

How do I prevent the user from changing the subdomain and seeing other users data or just trying see what subdomains are available ? I tried adding a before\_hook to the ApplicationController but seems that won't catch it and the page just failed.

Would I need to add a Generic elevator with the \`Subdomain\` elevator in the \`initializers/apartment.rb\` file?

    # initializers/apartment.rb
    Rails.application.config.middleware.insert_after Warden::Manager, Apartment::Elevators::Generic, -&gt; (request) {
       # code to check whether user is going to correct subdomain
    }

Thanks and any help/suggestions appreciated.
## [10][When should we use inverse_of?](https://www.reddit.com/r/rails/comments/f4qt9g/when_should_we_use_inverse_of/)
- url: https://www.reddit.com/r/rails/comments/f4qt9g/when_should_we_use_inverse_of/
---
I've seen articles like [this one](https://www.viget.com/articles/exploring-the-inverse-of-option-on-rails-model-associations) where they say `inverse_of` is useful to optimize memory when fetching associated records. But do we have to use it for each association in our models? Rails don't do this automatically in some cases?
## [11][Storing attachment file type when using active storage](https://www.reddit.com/r/rails/comments/f4plor/storing_attachment_file_type_when_using_active/)
- url: https://www.reddit.com/r/rails/comments/f4plor/storing_attachment_file_type_when_using_active/
---
Hey,

I was wondering if there’s a way to store, in the database, the attachment file type when using active storage. Basically a way to differentiate if a file is an Image Video etc.

I was thinking of making a monkey patch Active Storage Blob Model that adds this functionality based on MIME types but I’m assuming there must be a better way around it.

Thanks!
## [12][Paying for help with ruby on rails I'm on a project that is driving me nuts.](https://www.reddit.com/r/rails/comments/f4ydjz/paying_for_help_with_ruby_on_rails_im_on_a/)
- url: https://www.reddit.com/r/rails/comments/f4ydjz/paying_for_help_with_ruby_on_rails_im_on_a/
---
Hop on a zoom call tonight will pay $40 an hour. 
I need help asap will explain on call.
