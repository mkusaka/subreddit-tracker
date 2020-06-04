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
## [2][100 days of code to learn rails](https://www.reddit.com/r/rails/comments/gw2y2l/100_days_of_code_to_learn_rails/)
- url: https://www.reddit.com/r/rails/comments/gw2y2l/100_days_of_code_to_learn_rails/
---
just wanted to share I'm starting my web dev journey and choosing to go through Hartl's rails 6 tutorial. a lot of negatives on rails hype dying, but really looking forward to get a foundation of being able to build MVP's for some of my ideas I have. I've been stalking this sub a bit and it's really cool that you guys are so practical and helpful. I hope it's okay to share here and probably will be posting here for some help along the way :)

starting a substack if anyone is interested in following a newb: [https://joelchoi.substack.com/p/day-1](https://joelchoi.substack.com/p/day-1)
## [3][Active record queries](https://www.reddit.com/r/rails/comments/gwgps7/active_record_queries/)
- url: https://www.reddit.com/r/rails/comments/gwgps7/active_record_queries/
---
Beginner in Ruby and Activerecord, I wanted to know if anyone could help me write this out and explain what happened. Here's what I figured out so far but I need confirmation or a guide. \\ I have more questions I need to figure out but I thought I should post one at a time. 

The question

1. Find who published the book with the highest average rating.

Avg(rating) = avg\_rating

Book.find\_by(:authour).order(avg\_rating desc)
## [4][Product with price history](https://www.reddit.com/r/rails/comments/gw47vj/product_with_price_history/)
- url: https://www.reddit.com/r/rails/comments/gw47vj/product_with_price_history/
---
Hello folks.

I have some issue to solve, hope you will help me. I'd like to store history of price. I've model products (every time I buy something I add it into my warehouse - name, price, quantity or so on) also I want to get last added price - how can I solve it  in the proper way? To have last price, history of changed prices and to by DRY. So ie: 

warehouse has\_any products (with quantity)

product has\_many prices

also I have meal recipes where I'd like to know actual food cost (using last added price) and I'd like to have history, that given meal or product cost me $30 few month ago and today it cost me $40...

any ideas?
## [5][Slides What Comes After MVC - Peter Harkins](https://www.reddit.com/r/rails/comments/gvs0dh/slides_what_comes_after_mvc_peter_harkins/)
- url: https://www.reddit.com/r/rails/comments/gvs0dh/slides_what_comes_after_mvc_peter_harkins/
---
Peter talked so fast. I can't keep up, let alone digest the ideas presented.

Does anyone have the slides and notes to [https://push.cx/2015/railsconf](https://push.cx/2015/railsconf) that they can share?

Thanks.
## [6][Is there a ajax html table sortable gem?](https://www.reddit.com/r/rails/comments/gw6g1w/is_there_a_ajax_html_table_sortable_gem/)
- url: https://www.reddit.com/r/rails/comments/gw6g1w/is_there_a_ajax_html_table_sortable_gem/
---
I'm looking around for something that allows me to search a lot of data (currently using searchkick and will\_paginate) and let me add ajax and sortable columns.

Maybe I'm just looking in the wrong places but many of the gems either are discontinued or don't appear to take object sets from searchkick, but want to control the search themseleves.

Thanks for any help, reddit!
## [7][ActionMailer - Email not arriving, no error thrown. How to debug?](https://www.reddit.com/r/rails/comments/gvv63s/actionmailer_email_not_arriving_no_error_thrown/)
- url: https://www.reddit.com/r/rails/comments/gvv63s/actionmailer_email_not_arriving_no_error_thrown/
---
Hello,

I'm trying to get ActionMailer working. First off I have the following feature test:

    it 'sends an email', focus: true do
      test_form.visit_page.fill_in_with.submit
      expect(ActionMailer::Base.deliveries.count).to eq(1)
      expect(ActionMailer::Base.deliveries.last.to).to include(u.email)
    end

This test passes.

Here's my development.rb settings for action\_mailer:

    config.active_record.dump_schema_after_migration = false
    
    config.action_mailer.raise_delivery_errors = true
    
    config.action_mailer.perform_deliveries = true
    
    config.action_mailer.delivery_method = :smtp
      config.action_mailer.smtp_settings = {
        :address              =&gt; "mail.privateemail.com",
        :port                 =&gt; 465,
        :user_name            =&gt; ENV['EMAIL_USERNAME'],
        :password             =&gt; ENV['REMAIL_PASSWORD'],
      }

My username is my full email: [me@mydomain.com](mailto:me@mydomain.com)These settings are from my email provider which can be found on [this page](https://www.namecheap.com/support/knowledgebase/article.aspx/1179/2175/general-private-email-configuration-for-mail-clients-and-mobile-devices).

Here's my appliction\_mailer.rb

    class ApplicationMailer &lt; ActionMailer::Base
      default from: 'me@mydomain.com'
      layout 'mailer'
    end

Here's mailer/test\_mailer.rb

    class TestMailer &lt; ApplicationMailer
      def test(user:)
        mail(to: user.email,
        subject: "This is a test",
        body: "body...")
      end
    end

Here's views/test\_mailer/test.html.erb

    &lt;h1&gt; THIS IS A TEST &lt;/h1&gt;

When I go through the process manually, everything appears to go through. This is what I get in the console:

    VenueMailer#new_venue_listed: processed outbound mail in 0.6ms
    Delivered mail 5ed7b46cbe880_4c811367826e@Jonnys-MacBook-Pro.local.mail (1.7ms)
    Date: Wed, 03 Jun 2020 15:32:12 +0100
    From: mail.privateemail.com
    To: myemal@gmail.com
    Message-ID: &lt;5ed7b46cbe880_4c811367826e@Jonnys-MacBook-Pro.local.mail&gt;
    Subject: This is a test
    Mime-Version: 1.0
    Content-Type: text/plain;
     charset=UTF-8
    Content-Transfer-Encoding: 7bit
    
    body...
    Redirected to http://localhost:3000/
    Completed 302 Found in 179ms (ActiveRecord: 4.8ms | Allocations: 11556)

So I don't get any errors, but the email simply never arrives.

How do I debug this?

Thanks.
## [8][Baserails course still relevant?](https://www.reddit.com/r/rails/comments/gw3epa/baserails_course_still_relevant/)
- url: https://www.reddit.com/r/rails/comments/gw3epa/baserails_course_still_relevant/
---
The marketplace course found here:

[https://www.baserails.com/marketplace](https://www.baserails.com/marketplace)

is exactly the sort of project I'm looking to build with Rails. I'm still learning, but this fits perfectly with the product I'd like to eventually build. I noticed the copyright on the site is from 2014, so it's a bit older now. Any thoughts on if this material would still be relevant?  Any other suggestions I might look at instead?
## [9]["--webpack" flag, some explanation needed.](https://www.reddit.com/r/rails/comments/gvuw61/webpack_flag_some_explanation_needed/)
- url: https://www.reddit.com/r/rails/comments/gvuw61/webpack_flag_some_explanation_needed/
---
Last night I was reading something about rails apps. Something in the article attracted my attention. It was this :

`rails new myapp --webpack=react`

And it actually made a few questions in my head. So I'm asking them here.

1. Is it only a way to load needed frontend framework to the app? Or it also loads a basic theme or something? (in case we may want some beautiful themes for the app).
2. I saw another tutorial used `vue` instead of react. Which frameworks does it exactly support?

I'm thankful for your kind answers.
## [10][Services and tools we use for development](https://www.reddit.com/r/rails/comments/gvpvlp/services_and_tools_we_use_for_development/)
- url: https://www.reddit.com/r/rails/comments/gvpvlp/services_and_tools_we_use_for_development/
---
Tools and services that help to automatize development flow. Easy to install with increasing team effectiveness.

[https://jtway.co/services-and-tools-we-use-for-development-2749af5ad08a](https://jtway.co/services-and-tools-we-use-for-development-2749af5ad08a)
## [11][Question about how a Rails backend works...](https://www.reddit.com/r/rails/comments/gvm3r6/question_about_how_a_rails_backend_works/)
- url: https://www.reddit.com/r/rails/comments/gvm3r6/question_about_how_a_rails_backend_works/
---
This is a general programming question and may not be specific to Rails. However, it's an issue I've run into building an app with a Rails backend.

Say I make an app and store user information in the Rails server-- such as a login and pw-- and push it all to github... if somebody clones that repo, runs it and tries to login with a login/pw that I saved will they be able to do it?
