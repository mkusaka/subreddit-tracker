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
## [2][How to handle multiple required strong params](https://www.reddit.com/r/rails/comments/ghhnkv/how_to_handle_multiple_required_strong_params/)
- url: https://www.reddit.com/r/rails/comments/ghhnkv/how_to_handle_multiple_required_strong_params/
---
Hello, what should params for a structure like the following look like?

\`{"person": {"first\_name": "Bob"}, "address": {"street": "123 Main St."}}\`

I've tried several different options based on different posts, including the array form of \`params.require()\`, but ran into different issues no matter which version I used (including 'last permit wins' for one of the two entity's attributes).

So there are two (or more) entities defined in parallel--not one top-level entity that can nest everything beneath itself. Anybody have an example of that?

For the array form, API docs show a parallel assignment, but then trying to \`permit()\` on each one of those param set variables "worked," but I ended up with the "last permit wins" issue mentioned above: only one set of attributes arrived, the other entity's were discarded.
## [3][sending an automated email that has an xlsx attachment through a scheduler job](https://www.reddit.com/r/rails/comments/ghnizc/sending_an_automated_email_that_has_an_xlsx/)
- url: https://www.reddit.com/r/rails/comments/ghnizc/sending_an_automated_email_that_has_an_xlsx/
---
So I've been going over documentation/articles on it.  It's scheduler job and I want to get a list of say all the products from the Product model and various aspects/data related to it, and send it off every X day/s to recipients via the scheduler job.

Will be using caxlsx:

https://github.com/caxlsx/caxlsx_rails

So say I have a Product model, will I be having most of the logic in the controller for building the xlsx sheet such as the below, or since nothing is needed for rendering(it's simply data sent out in an automated email attachment), is it going to just be a Product class method that I call in the scheduler job?

    # products_controller.rb 

    def product_listing
      @products = Product.all

      respond_to do |format| # is this even necessary, does all my code just need to be in a model since I'm not rendering anything?
        format.xlsx
      end

      wb = xlsx_package.workbook #create excel book
      wb.add_worksheet(name: "Product List")
      sheet.add_row ["Product Name ", "Product Price", "Product Inventory Date"]
      @products.each to |product|
        sheet.add_row [product.name, product.price, product.inventory_date]
      end
    end

I guess thinking this out further, regardless of whether it will be in the model or the controller, I will need a method/action in a mailer model to kick off an email, and in that need to use the controller action or model class method, correct?  Just trying to conceptually understand how to do the email along with the xlsx attachment, and I guess focusing on building out the xlsx sheet primarily to start.  And I guess I'll need that mailer method/action to be kicked off in the scheduler job as well.
## [4][Create Your First Rails Template For Devise + Bootstrap](https://www.reddit.com/r/rails/comments/gh6leo/create_your_first_rails_template_for_devise/)
- url: https://www.reddit.com/r/rails/comments/gh6leo/create_your_first_rails_template_for_devise/
---
Hey everyone, Deanin here!

I've been working on this 20 projects in 20 weeks challenge where I learn to build something new every week and then make a tutorial covering it.

For week 8 I finally snapped and couldn't take adding usernames for Devise accounts to login with and bootstrap anymore, so I set out to learn how to build a template that I could run with one command. 

After you create the template you can host it somewhere and use a link to the raw file to run your command, meaning no download necessary. You could also use my template if you like, just to see how it works.

The rails command from my blog uses a bit.ly link shortened URL. I'll include that command so you can test drive what it does:

`rails new app_name -m bit.ly/deanin_template`

I created both a text version of the tutorial on my blog that you can find here, allowing you to just get up and running faster, as well as a YouTube tutorial if you'd like things explained by me a bit more.

Hopefully either of these resources will help at least one of you. I definitely learned a lot making from doing this project, for what it's worth.

[Blog Post](https://deandehart.com/blog/rails-template-for-devise-and-bootstrap/)

[YouTube Tutorial](https://youtu.be/JR8u5gfmDxU)
## [5][Confusion over push notification](https://www.reddit.com/r/rails/comments/gha4p8/confusion_over_push_notification/)
- url: https://www.reddit.com/r/rails/comments/gha4p8/confusion_over_push_notification/
---
Hi guys, I come here in hope for clarification.

I would like to implement in my app push notifications (the one that appear in your phone when you get a new whatsapp). Originally, I did some research and saw that people are using firebase, in what seems to me a very complex sets of step, but I also found a gem, rpush, which seems to be the standard. Are these notifications like a post to a specific routes with params?  


To be honest, after reading all that I am not even sure of what I need of if I have to rely on third services. In mi case I am working toward a rails backend appi and a mobile app with Ionic or similar. Can you help me through?
## [6][why does moving my ActionCable broadcast to a background worker prevent the client side received method from executing?](https://www.reddit.com/r/rails/comments/ghbkka/why_does_moving_my_actioncable_broadcast_to_a/)
- url: https://www.reddit.com/r/rails/comments/ghbkka/why_does_moving_my_actioncable_broadcast_to_a/
---
If I do it the bad way, and run the broadcast from the model's `after_create_commit` directly, the client side `received` method is called, and my browser goes like `alert("wow hey there")`

    class Message &lt; ApplicationRecord
      after_create_commit { ActionCable.server.broadcast("messages_channel", message: self) }
    end

But if I delegate that code to a background worker....suddenly `received` is no longer called:

    class Message &lt; ApplicationRecord
      after_create_commit { MessageBroadcastWorker.perform_async self.id }
      #the self.id isn't the issue - you shouldn't be passing whole objects to sidekiq
    end

This is the background worker:

    class MessageBroadcastWorker
      include Sidekiq::Worker
    	
      def perform(message_id)
        message = Message.find message_id
        ActionCable.server.broadcast("messages_channel", message: render_message(message))
      end
    
      private
      def render_message(message)
        ApplicationController.renderer.render(partial: 'messages/message', locals: { message: message })
      end
    end

The worker runs, it just doesn't execute the code:

    2020-05-10T22:23:34.679Z pid=35945 tid=owo1k87ct class=MessageBroadcastWorker jid=5d0fac3c80268538d582b7b4 INFO: start
    2020-05-10T22:23:34.883Z pid=35945 tid=owo1k87ct class=MessageBroadcastWorker jid=5d0fac3c80268538d582b7b4 elapsed=0.203 INFO: done

What gives?
## [7][Setting up an Automatic Book reader with Devise + Rails - 1](https://www.reddit.com/r/rails/comments/ggzr52/setting_up_an_automatic_book_reader_with_devise/)
- url: https://www.reddit.com/r/rails/comments/ggzr52/setting_up_an_automatic_book_reader_with_devise/
---
Hey guys,

I hate how much coding channels focus on a basic 'ToDo list / blog' so  i wanted wanted to share my progress in building an automatic book reader (not that good in making vids but meh). Probably will do several more videos until the end goal of having a tracker for how many books I've read.

&amp;#x200B;

 [https://www.youtube.com/watch?v=AnDLaaTXeWg&amp;feature=youtu.be](https://www.youtube.com/watch?v=AnDLaaTXeWg&amp;feature=youtu.be) 

&amp;#x200B;

Thx!
## [8][How do you boost speed for upload multi files in Active Storage?](https://www.reddit.com/r/rails/comments/gh8p9b/how_do_you_boost_speed_for_upload_multi_files_in/)
- url: https://www.reddit.com/r/rails/comments/gh8p9b/how_do_you_boost_speed_for_upload_multi_files_in/
---
I am using Active Storage wit GCP, it takes times usually to upload multi files for me, and sometimes it was uploaded but I got respond Cant build URI ; it seems like that the files already on my Bucket, and it saved also on database, but it said my attachment from database zero, I saw my reford_id on active_storage_blobs always zero, why his happened?

- how do you Boost upload with Active storage?
- why my record_id is zero ??
- how to speed to remove files also with .purge?
## [9][What is the most pleasant frontend stack of your choice?](https://www.reddit.com/r/rails/comments/gglkff/what_is_the_most_pleasant_frontend_stack_of_your/)
- url: https://www.reddit.com/r/rails/comments/gglkff/what_is_the_most_pleasant_frontend_stack_of_your/
---

## [10][Rails with React](https://www.reddit.com/r/rails/comments/ggvbcn/rails_with_react/)
- url: https://www.reddit.com/r/rails/comments/ggvbcn/rails_with_react/
---
Hi just learnt to build things rails as back end. Is there any good tut where I can learn front end with rails? I want to learn react and node js with rails.
## [11][How do I set up this ActiveRecord association where a model has both a "owner" and a "user"?](https://www.reddit.com/r/rails/comments/ggjwsa/how_do_i_set_up_this_activerecord_association/)
- url: https://www.reddit.com/r/rails/comments/ggjwsa/how_do_i_set_up_this_activerecord_association/
---
Hi all,

I've just created the following migration to allow a Venue to have an owner (by association with the users table) according to [these instructions](https://stackoverflow.com/questions/27809342/rails-migration-add-reference-to-table-but-different-column-name-for-foreign-ke).

    class AddOwnerToVenue &lt; ActiveRecord::Migration[6.0]
      def change
        add_reference :venues, :owner, foreign_key: { to_table: :users }
      end
    end

&amp; I have the following association set up in venue.rb:

    class Venue &lt; ApplicationRecord
      has_one :owner belongs_to :user
    end

Now for User, I have the corresponding association:

    class User &lt; ApplicationRecord
      has_many :venues
    end

But in order for my new owner association to work, I need to do something similar to this, but for it to look in the venue.owner column, not the venue.user.

Something like:

    class User &lt; ApplicationRecord
      has_many :venues, as: owner
    end

I can't find a way to do this. Hopefully someone has come across this kind of association before and can help me out?

Thanks.
