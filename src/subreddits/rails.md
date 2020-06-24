# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/har6r7/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/har6r7/personal_projects_show_off_your_own_project_andor/
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
## [2][Solidus won't display images out of the box - Paperclip and ImageMagick error](https://www.reddit.com/r/rails/comments/hf0te0/solidus_wont_display_images_out_of_the_box/)
- url: https://www.reddit.com/r/rails/comments/hf0te0/solidus_wont_display_images_out_of_the_box/
---
I'm just starting with Solidus, though I do know my way around Rails. I've created a Solidus project on my Mac, and I installed all the required gems as well as ImageMagick through Homebrew.  Both `identify` and `magick` work as terminal commands as my `/usr/local/bin` directory is in my PATH. I've also specified the path to `identify` in the `application.rb` file (I'm guessing it's supposed to go in side the `Application` class). However, none of the images show up in Solidus, and when I try and manually add an image, I get the following error (it's related to paperclip). 

Any help here would be - Google has not been my friend this time. BTW, the same thing happens on both Mac and Windows running WSL (Debian). The error message is shown below:

    [paperclip] Trying to link /tmp/RackMultipart20200622-9041-8gndhx.jpg to /tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-1t35lnu.jpg
    [paperclip] Trying to link /tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-1t35lnu.jpg to /tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-62aumc.jpg
    Command :: PATH=/usr/local/bin/identify:$PATH; file -b --mime '/tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-62aumc.jpg'
    Command :: PATH=/usr/local/bin/identify:$PATH; identify -format '%wx%h,%[exif:orientation]' '/tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-1t35lnu.jpg[0]' 2&gt;/dev/null
    [paperclip] An error was received while processing: #&lt;Paperclip::Errors::NotIdentifiedByImageMagickError: Paperclip::Errors::NotIdentifiedByImageMagickError&gt;
    Command :: PATH=/usr/local/bin/identify:$PATH; identify -format '%wx%h,%[exif:orientation]' '/tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-1t35lnu.jpg[0]' 2&gt;/dev/null
    [paperclip] An error was received while processing: #&lt;Paperclip::Errors::NotIdentifiedByImageMagickError: Paperclip::Errors::NotIdentifiedByImageMagickError&gt;
    Command :: PATH=/usr/local/bin/identify:$PATH; identify -format '%wx%h,%[exif:orientation]' '/tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-1t35lnu.jpg[0]' 2&gt;/dev/null
    [paperclip] An error was received while processing: #&lt;Paperclip::Errors::NotIdentifiedByImageMagickError: Paperclip::Errors::NotIdentifiedByImageMagickError&gt;
    Command :: PATH=/usr/local/bin/identify:$PATH; identify -format '%wx%h,%[exif:orientation]' '/tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-1t35lnu.jpg[0]' 2&gt;/dev/null
    [paperclip] An error was received while processing: #&lt;Paperclip::Errors::NotIdentifiedByImageMagickError: Paperclip::Errors::NotIdentifiedByImageMagickError&gt;
    [paperclip] Trying to link /tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-1t35lnu.jpg to /tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-1dlkgdw.jpg
    Command :: PATH=/usr/local/bin/identify:$PATH; file -b --mime '/tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-1dlkgdw.jpg'
    [paperclip] Trying to link /tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-1t35lnu.jpg to /tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-16yjulo.jpg
    Command :: PATH=/usr/local/bin/identify:$PATH; file -b --mime '/tmp/cdcf58ad03be2998ea64da8e6cd3e43c20200622-9041-16yjulo.jpg'
      Rendering /Library/Ruby/Gems/2.6.0/gems/solidus_backend-2.10.1/app/views/spree/admin/images/create.js.erb
      Rendered /Library/Ruby/Gems/2.6.0/gems/solidus_backend-2.10.1/app/views/spree/admin/images/create.js.erb (Duration: 3.4ms | Allocations: 874)
    Completed 200 OK in 836ms (Views: 20.7ms | ActiveRecord: 7.6ms | Allocations: 29100)
## [3][Live coding: refactoring video with coding tips](https://www.reddit.com/r/rails/comments/hesedi/live_coding_refactoring_video_with_coding_tips/)
- url: https://www.reddit.com/r/rails/comments/hesedi/live_coding_refactoring_video_with_coding_tips/
---
Hello everyone

It's a refactoring video with tips to make our code more maintainable.

Last year I presented a refactoring at Tokyo Rubyist Meetup. On stage I only used slides, for the video I wanted to do a live coding which I think is more appropriate for an online video.

That was my first presentation and it's my first video! 🎉

[https://youtu.be/X0542oOs1Qg](https://youtu.be/X0542oOs1Qg)

&amp;#x200B;

Feedback are welcome. Prononciation is a struggle (in any language), I'm working on it, hope it's ok enough.
## [4][Advanced Sinatra tutorials?](https://www.reddit.com/r/rails/comments/heknov/advanced_sinatra_tutorials/)
- url: https://www.reddit.com/r/rails/comments/heknov/advanced_sinatra_tutorials/
---
Would appreciate someone pointing me to any tutorials that show the advanced use of Sinatra. In a nutshell what am trying to do is have a scrapper pull some data from various websites for me (have this covered using another service), but then I need to poll that data and use a Messenger bot to send messages to a subscriber list I have (it is this particular use case that has me thinking of using Sinatra). Thanks.
## [5][How would you solve a "easy to edit complex html" situation like ....](https://www.reddit.com/r/rails/comments/hej9yo/how_would_you_solve_a_easy_to_edit_complex_html/)
- url: https://www.reddit.com/r/rails/comments/hej9yo/how_would_you_solve_a_easy_to_edit_complex_html/
---
Use case: Our sales guys like using things like [https://instapage.com/](https://instapage.com/) to create frequently edited sales type content. 

Problem: But of a lot of this is VERY repetitive for them (like they need to do a similar page for 1,000 clients) - so it makes sense to rails/erb that up and host it ourselves. They also want to do a ton of subdomain stuff that services like instapage don't really support well. But - since this is now hosted in rails the sales guys can't really easily edit/publish anymore (since we don't have a wiz-bang wysiwyg type editor for complex html, don't have a simple way for them to go through deploy -&gt; staging -&gt; prod, etc.... they're not rails devs).

Does anyone have any great solutions to something like this? The existing wysiwyg get editors for rails (afaik) are mostly just for doing word processing type editing and not for doing HTML editing with divs, js-plugins, etc.
## [6][Could not found leave without an id](https://www.reddit.com/r/rails/comments/hersml/could_not_found_leave_without_an_id/)
- url: https://www.reddit.com/r/rails/comments/hersml/could_not_found_leave_without_an_id/
---
i am using rails 5.2 and i am new to rails but i find error Couldn't find Leave without an ID please guide mn

Controller code

def leave\_to\_approveleave = Leave.find(params\[:id\])status = params\[:rejected\].present? ? params\[:rejected\] : "Approved"if leave.update\_attribute("status",status)flash\[:notice\] = "Leave updated successfully"endend

view code

&lt;% u/leaves.each\_with\_index do |leave, index| %&gt;

&lt;tr&gt;    
&lt;td&gt;&lt;%= index + 1 %&gt;&lt;/td&gt;    
&lt;td&gt;&lt;%= leave.leave\\\_date %&gt;&lt;/td&gt;    
&lt;td&gt;&lt;%= leave.leave\\\_description %&gt;&lt;/td&gt;    
&lt;td&gt;&lt;%= leave.status %&gt;&lt;/td&gt;    
&lt;td&gt;&lt;%= link\\\_to 'Approve', dashbord\\\_leave\\\_to\\\_approve\\\_path(leave),method: :put,class: "btn btn-large btn-success" %&gt;&lt;td&gt;    
&lt;td&gt;&lt;%= link\\\_to "Reject", dashbord\\\_leave\\\_to\\\_approve\\\_path(leave), method: :put, class: "btn btn-large btn-danger" %&gt;&lt;td&gt;    
&lt;/tr&gt;    
 &lt;% end %&gt;  


routes code

get 'dashbord/leaves'  
 put 'dashbord/leave\_to\_approve'
## [7][Integrate a bootstrap theme into a Rails app?](https://www.reddit.com/r/rails/comments/heht08/integrate_a_bootstrap_theme_into_a_rails_app/)
- url: https://www.reddit.com/r/rails/comments/heht08/integrate_a_bootstrap_theme_into_a_rails_app/
---
I'm currently learning Rails and I'm working to add a theme to my app to make it look, well, not terrible.  Specifically, I'm looking at [this](https://startbootstrap.com/themes/sb-admin-2/) theme.  I've tried following along with this [Medium post](https://medium.com/@yli0607x/how-to-use-bootstrap-themes-on-ruby-on-rails-in-5-minutes-8e6f9542f6d8), and I got it mostly working, but the fonts are way small and some of the JS isn't firing (see my [repo](https://github.com/kroe761/sb-admin-2) here).  Thanks!
## [8][best way to get range of keys using specific pattern in Redis](https://www.reddit.com/r/rails/comments/hehqsq/best_way_to_get_range_of_keys_using_specific/)
- url: https://www.reddit.com/r/rails/comments/hehqsq/best_way_to_get_range_of_keys_using_specific/
---
I work with Redis but I need best way to get keys from Redis .

I need search by range of date `(2020-06-01 to 2020-06-20)` for every company\_id **(id is unique)**

**example of key** `stat:company:2:date:2020-06-13`

I think to solve problem i get key for every day but if need 200 numbers of key .

it take many times and i make some operation with value before display user .

i try use scan  
 &amp; keys

`Environment Ruby 2.7.0`  
`, Ruby on Rails  6.0.3`  
 `Redis`
## [9][Where do I put this code?](https://www.reddit.com/r/rails/comments/heku2i/where_do_i_put_this_code/)
- url: https://www.reddit.com/r/rails/comments/heku2i/where_do_i_put_this_code/
---
I am working on an app that has models:

Accounts -&gt; Transactions

I want to add a "Transfer" feature to my app which will essentially create a debit Transaction in one Account, and a credit Transaction in another Account.  

My thought was to create a transfer controller action in the AccountController, and put the code there, but I'm reading about helpers, service objects, domain objects, and just putting the code in the Model instead. 

I'm getting confused and wondering where would be the best place to code this logic? 

Thanks!
## [10][NoMethodError (undefined method `email' for #&lt;User:0x00007f8d981f3c90&gt;)](https://www.reddit.com/r/rails/comments/hei2lp/nomethoderror_undefined_method_email_for/)
- url: https://www.reddit.com/r/rails/comments/hei2lp/nomethoderror_undefined_method_email_for/
---
 

I am using devise for sign-up and have removed all :email references from migration files, updated devise.rb to :username for authentication and I get this error. It won't let me create a user in the console. Has anyone seen this before and know how to fix it? Any help is appreciated, I've tried everything.
## [11][change status of leave from pending to approve](https://www.reddit.com/r/rails/comments/hegnkx/change_status_of_leave_from_pending_to_approve/)
- url: https://www.reddit.com/r/rails/comments/hegnkx/change_status_of_leave_from_pending_to_approve/
---
def leave\_to\_approve  
 leave = Leave.find(params\[:id\])  
 status = params\[:rejected\].present? ? params\[:rejected\] : "Approved"  
 if leave.update\_attribute("status",status)  
 flash\[:notice\] = "Leave updated successfully"  
 end  
 end

view code

 &lt;tbody&gt;  
 &lt;% leaves.each\_with\_index do |leave, index| %&gt;  
&lt;tr&gt;  
&lt;td&gt;&lt;%= index + 1 %&gt;&lt;/td&gt;  
&lt;td&gt;&lt;%= leave.leave\_date %&gt;&lt;/td&gt;  
&lt;td&gt;&lt;%= leave.leave\_description %&gt;&lt;/td&gt;  
&lt;td&gt;&lt;%= leave.status %&gt;&lt;/td&gt;  
&lt;td&gt;&lt;%= link\_to 'Approve', dashbord\_leave\_to\_approve\_path(leave),method: :put,class: "btn btn-large btn-success" %&gt;&lt;td&gt;  
&lt;td&gt;&lt;%= link\_to "Reject", dashbord\_leave\_to\_approve\_path(leave), method: :put, class: "btn btn-large btn-danger" %&gt;&lt;td&gt;  
&lt;/tr&gt;  
 &lt;% end %&gt;  
  &lt;/tbody&gt;

routes code 

get 'dashbord/leaves'  
 put 'dashbord/leave\_to\_approve'

but it give error

# ActiveRecord::RecordNotFound in DashbordController#leave_to_approve

## Couldn't find Leave without an ID
