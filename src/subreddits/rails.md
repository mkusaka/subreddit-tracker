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
## [3][How can I clean up this code adding a CSS background: url() dynamically from model?](https://www.reddit.com/r/rails/comments/f6rmwn/how_can_i_clean_up_this_code_adding_a_css/)
- url: https://www.reddit.com/r/rails/comments/f6rmwn/how_can_i_clean_up_this_code_adding_a_css/
---
Hi everyone,  


I have a a model `coffee_shops` which `has_many` "`review_photos`".  


I want to display the `review_photo` with a gradient, using the following CSS:  


    background: linear-gradient(rgba(0,0,0,0.2),rgba(0,0,0,0.8)), url('www.mydynamicurl.com') center / cover;

Now, I can't send the image URL to the CSS as far as I'm aware. So I assume I have to do this in HTML. I have done this like so:  


    &lt;section class="hero is-medium is-bold" style="background: linear-gradient(rgba(0,0,0,0.2),rgba(0,0,0,0.8)), url('http://my-url-prefix/&lt;%= @coffee_shop.review_photos.first[:photo] %&gt;') center / cover;"&gt;

This works visually, but there's 2 problems.  


1) This seems kind of ugly and hacky.  
2) I might want to use this gradient on other pages. Putting it in the HTML like this means that I have to copy it each time, and if I make a change to the gradient then I have to find all instances of it and change it on each.  


What's the best way of doing this, (bearing in mind that the image has to be assigned in CSS and can't be a HTML &lt;img&gt; element.  


Thanks.
## [4][What do you do to stay sharp?](https://www.reddit.com/r/rails/comments/f6l733/what_do_you_do_to_stay_sharp/)
- url: https://www.reddit.com/r/rails/comments/f6l733/what_do_you_do_to_stay_sharp/
---
It's something everyone agrees is absolutely crucial to maintain relevancy as a programmer, constant learning and self development.

What are some of the things you do to stay up to date on trends and constantly learn and develop some skills? What are some things you're working on in that regard? What are things that have helped you in the past?

This is something I've struggled with especially recently. It's easy to settle into a comfortable rhythm at work and go months without adding any new meaningful knowledge or skills.
## [5][Active Storage File Upload Behind The Scenes](https://www.reddit.com/r/rails/comments/f6t0fs/active_storage_file_upload_behind_the_scenes/)
- url: https://www.reddit.com/r/rails/comments/f6t0fs/active_storage_file_upload_behind_the_scenes/
---
[Active Storage](https://github.com/rails/rails/tree/master/activestorage)   is a framework in Ruby that makes it a breeze to upload files and   reference them in the cloud (or a local disk). It’s built into [Ruby On Rails 6](https://weblog.rubyonrails.org/2019/8/15/Rails-6-0-final-release/), but it’s also got a [JavaScript library](https://github.com/rails/rails/tree/master/activestorage/app/javascript/activestorage). In fact, that’s what I like most about Rails. It’s got your back. It delivers *complete* packages out of the box. Packages that work well *together*. Client to server — it’s got everything you need! And it’s got beautiful code. Making it a pleasure to work with.

In this post, you will **see how Active Storage really works from the inside**.   We will track the main flow of the program, and see how it processes a   file uploaded by the user through the browser with JavaScript. And  then,  how the file is uploaded to a local disk with Ruby. *Let’s get started.*

[Continue reading on Medium...](https://medium.com/rubyinside/active-storage-file-upload-behind-the-scenes-59a660c43781)
## [6][Invite Links with Devise](https://www.reddit.com/r/rails/comments/f6qfrq/invite_links_with_devise/)
- url: https://www.reddit.com/r/rails/comments/f6qfrq/invite_links_with_devise/
---
I want to build an invite system where each registered user has a unique invite link. And when their friends use that link to sign up they both get benefits. You should also be able to send the link through social media or email.

An example of what I mean is airbnb's invite system:

https://preview.redd.it/13i9orrbo1i41.png?width=3354&amp;format=png&amp;auto=webp&amp;s=f9b4d59c39a1219548dfb8223fa5f277a485a439

I have tried using devise invitable but it seems to work differently. It seems to register an user with an email first and then signs them up if they accept the invite. This does not allow for social media sharing.

So, how do I go about building such an invite system and how do I integrate it with devise?
## [7][Should I start rails or something else?](https://www.reddit.com/r/rails/comments/f6q582/should_i_start_rails_or_something_else/)
- url: https://www.reddit.com/r/rails/comments/f6q582/should_i_start_rails_or_something_else/
---
Hey How are you all guys?? I hope you all doing fine.. I have a question.. I have a hard time deciding which language to study.. I just wanna get into software development.. I really like ruby and its framework rails.. But the problem is there are not enough jobs in rails.. So my question is Shoud I..  
1) start rails and go with my interest  
2) start with something else.. lets say laravel or react.. because there are tons of jobs in both of these..  


So I want advice from you people.. can I pursue my career in freelancing in rails world.. what are the prerequisite.. I mean how much your skill is must to start freelancing and get success.. because if there are not enough jobs in my area .. freelancing is the way to go..
## [8][What Ruby gems would you recommend for Rails API development](https://www.reddit.com/r/rails/comments/f6di8l/what_ruby_gems_would_you_recommend_for_rails_api/)
- url: https://www.reddit.com/r/rails/comments/f6di8l/what_ruby_gems_would_you_recommend_for_rails_api/
---
I am looking into making API only Rails app and I wanted to make a post here to ask you guys whether you know any good gems that would make my life easier for developing API only Rails apps. Maybe something for authentication, responses, etc. Thanks!
## [9][Rails migration for belongs_to with custom table name](https://www.reddit.com/r/rails/comments/f6hunj/rails_migration_for_belongs_to_with_custom_table/)
- url: https://www.reddit.com/r/rails/comments/f6hunj/rails_migration_for_belongs_to_with_custom_table/
---
Check out my recent [post](https://railsguides.net/foreign-key-for-custom-table-name/) describes how to specify a belongs to association in Rails migration with a foreign key to custom table name. Spoiler: as the result it has improved Rails docs!
## [10][I wrote a post about how geared pagination works behind-the-scenes —&gt;](https://www.reddit.com/r/rails/comments/f6jhza/i_wrote_a_post_about_how_geared_pagination_works/)
- url: https://www.reddit.com/r/rails/comments/f6jhza/i_wrote_a_post_about_how_geared_pagination_works/
---
There’s  a better, more user-friendly-driven approach to writing pagination for  your application. The idea is simple: Load more records every time the user is asking to see more.

[Continue reading on Medium...](https://medium.com/@liroy/geared-pagination-in-rails-behind-the-scenes-61d9e227540e)
## [11][Ruby 2.7 removes taint checking mechanism](https://www.reddit.com/r/rails/comments/f67vxa/ruby_27_removes_taint_checking_mechanism/)
- url: https://blog.saeloun.com/2020/02/18/ruby-2-7-access-and-setting-of-safe-warned-will-become-global-variable
---

## [12][[Help] Optimize model validations](https://www.reddit.com/r/rails/comments/f6bf1b/help_optimize_model_validations/)
- url: https://www.reddit.com/r/rails/comments/f6bf1b/help_optimize_model_validations/
---
Hi guys, I have the following model and I looking to optimize it...

    # some_model.rb
    
    belongs_to :some_association
    
    validates :some_association_id, presence: true
    
    validate :check_status_1
    validate :check_status_2
    
    
    def check_status_1
      if self.some_association_id.present?
        begin
          some_association = Some::Association.find(self.some_association_id)
          # check status_1... else raise error validation
        
          rescue ActiveRecord::RecordNotFound =&gt; e
            # case treated by model validation - so nothing goes here
          end  
        end
      end
    end
    
    
    def check_status_2
      if self.some_association_id.present?
        begin
          some_association = Some::Association.find(self.some_association_id)
          # check status_2... else raise error validation
        
          rescue ActiveRecord::RecordNotFound =&gt; e
            # case treated by model validation - so nothing goes here
          end  
        end
      end
    end

Because I see a lot of hits on database for the same query (Some::Association.find(...)  and belongs\_to validation), and I not sure if Rails reuse them from the memory, I was thinking if it would be possible to reduce this hits on db.

Any tip to improvement is well accepted.

My suggestion to solve the 2 queries Some::Association.find(...) problem is: an instance variable, something like this:

    def check_status_1/2
      if @some_association.present? &amp;&amp; self.some_association_id.present?
        begin
          @some_association = Some::Association.find(self.some_association_id)
          ...
        end
      end
    end

Another way to solve the query problem would be put them together in one method (check\_status for status\_1 &amp; status\_2), but that would make hard to test each situation.

About belongs\_to validation hitting the database for the same query that check\_status does, I'm not sure how to solve it in a pleasant way.
