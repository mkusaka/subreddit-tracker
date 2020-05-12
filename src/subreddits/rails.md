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
## [2][Anyone provisioning Ubuntu 18.04 for Rails 6 using Ansible?](https://www.reddit.com/r/rails/comments/gi3k09/anyone_provisioning_ubuntu_1804_for_rails_6_using/)
- url: https://www.reddit.com/r/rails/comments/gi3k09/anyone_provisioning_ubuntu_1804_for_rails_6_using/
---
I'm trying to get this working based on an Ansible playlist/set of roles I've tested for Rails 5. However, I seem to be hitting endless problems:

* geerlingguy.mysql seems to have some 'caveats' for Ubuntu 18.04 and doesn't work out of the box, requires some extra python mysql packages to be installed
* geerlingguy.nodejs requires a new respository to be added to apt capistrano deployment is failing after I fix the above issues with some mysterious errors
* ansible-galaxy yarn repository only installs for login shells, which doesn't work for Capistrano

To clarify: my original set of Ansible scripts was working for Rails 5, but it seems to be *very* difficult to get it working for Rails 6.

It feels like I'm going at this the wrong way. Every step is pain! Isn't ansible-galaxy supposed to make it so I can grab a role, change one or two variables, and keep moving?

Do most people write their roles from scratch?

Is ansible (for provisioning) and capistrano (for deployment) still a popular combo for Rails 6?

Why does it feel like all of these tools are out of sync?

I'm not trying to be critical, I'm genuinely unsure if:

* This is just how it is, and I need to figure it out
* I'm doing something to complicate provisioning unnecessarily

I started looking into recommended Ansible textbooks, but the most popular one is Ansible for Devops by /u/geerlingguy.

To my horror, some of the reviews mentioend that this book refers readers to use /u/geerlingguy's ansible-galaxy packages which are exactly the ones I'm fumbling around with, so I'm not sure if it's going to help me at all.

Although his packages appear to be the most popular, some of the big ones (mysql, nodejs) seem to have issues with Ubuntu 18.04 despite showing 'xenial' on the list of supported distros.

What's going on here?
## [3][Rails 6 and Webpacker Frustration](https://www.reddit.com/r/rails/comments/ghw5cq/rails_6_and_webpacker_frustration/)
- url: https://www.reddit.com/r/rails/comments/ghw5cq/rails_6_and_webpacker_frustration/
---
I am in the process of learning Rails 6 after spending years fumbling around in Rails 4. I am striving to be a better developer and programmer, but I have a habit of trying to rely on tutorials to make things happen. So as I set out on this new venture, I am finding that Webpacker and the handling of assets vastly confusing, but as of today I am trying to resolve what is the best approach to setting up my new project. Do I let webpacker handle my css, bootstrap and images? Or do I try to work with the asset pipeline still? I want to learn what the best way to handle this situation, but everything I seem to google up just seems to add to my frustration at this point. 

As it stands, I have muscled my way through using Webpacker to handle all of it, but calls seem longer and look more clunky, while I do like the idea of having all of these assets in one place. And while I do seem to call my scss files, I cannot seem to figure out what order they are supposed to go in so that I get the desired effect that I was getting from my Rails 4 asset pipeline. 

Does anyone have any advice on how they approach this and why? Am I just missing something small?

Thanks for your time!
## [4][Storing money balances or calculating them dynamically](https://www.reddit.com/r/rails/comments/gi7v26/storing_money_balances_or_calculating_them/)
- url: https://www.reddit.com/r/rails/comments/gi7v26/storing_money_balances_or_calculating_them/
---
Hey,

This is more of an architecture question, I think, but maybe there's a good way to do this in Rails.

Let's say we have a Customer model which has many orders, some paid some not, but new orders are stored every minute, so there are many of them.

The problem is:You want to quickly show a table of customers with totals: paid, unpaid, refunds, etc for hundreds of customers, sort customers by those totals, have API endpoints that list those customers - hundreds per request.

Is it better to simply have a balance column for each type (paid, unpaid, refunds) and update it every time a new invoice is added, removed, updated OR always use group by (which is not that easy to implement from my knowledge in Rails, calculating totals all the time might be slow as well) or calculate sums dynamically ?. Normally I would go with a balance column, but I have to be super careful on orders data when and might get out of sync.

What do you guys think? What's the "industry standard" for this sort of thing?

Thanks
## [5][Problems with Chat App with Rails 6, ActionCable, and Stimulus](https://www.reddit.com/r/rails/comments/ghtudy/problems_with_chat_app_with_rails_6_actioncable/)
- url: https://www.reddit.com/r/rails/comments/ghtudy/problems_with_chat_app_with_rails_6_actioncable/
---
Hi,

I have a chat up that is set up with ActionCable and Stimulus. The Rails side is working fine.  I can post to the db, then refresh the browser to receive the message. I have the ability to create and post to multiple chatrooms (Hangouts).  All of the basic Rails stuff is working fine.

Here is what I have:

Based on the logfiles (bottom of post), the message seems to be broadcasted, but not rendered on the  page.  when I check the Network tab in Chrome Dev, I only see a “Subscribed” from cable and not a “Subscribed and Confirm\_Subscription” .

Message Controller ([app](https://github.com/jemiller1963/ink3/tree/feature-group-and-direct-message/app)/[controllers](https://github.com/jemiller1963/ink3/tree/feature-group-and-direct-message/app/controllers)/messages\_controller.rb )

    class MessagesController &lt; ApplicationController
      before_action :authenticate_user!
      before_action :set_hangout
    
      def create
        @message = @hangout.messages.create(message_params)
        MessageChannel.broadcast_to @hangout, message: render_to_string(@message)
      end
    
    
      private
    
      def set_hangout
        @hangout = current_user.hangouts.find(params[:hangout_id])
      end
    
      def message_params
        params.require(:message).permit(:body).merge(user: current_user)
      end
    
    end

Message Channel ([app](https://github.com/jemiller1963/ink3/tree/feature-group-and-direct-message/app)/[channels](https://github.com/jemiller1963/ink3/tree/feature-group-and-direct-message/app/channels)/message\_channel.rb)

    class MessageChannel &lt; ApplicationCable::Channel
      def subscribed
        stop_all_streams
        stream_for Hangout.find(params['id'])
      end
    
      def unsubscribed
        # Any cleanup needed when channel is unsubscribed
      end
    end

Channel Controller ([app](https://github.com/jemiller1963/ink3/tree/feature-group-and-direct-message/app)/[javascript](https://github.com/jemiller1963/ink3/tree/feature-group-and-direct-message/app/javascript)/[controllers](https://github.com/jemiller1963/ink3/tree/feature-group-and-direct-message/app/javascript/controllers)/channel\_controller.js)

    import { Controller } from "stimulus";
    import consumer from "channels/consumer";
    
    export default class extends Controller {
      static targets = [];
    
      connect() {
        console.log("Is this working", this.element)  //This works!
        this.subscription = consumer.subscriptions.create(
          {
            channel: "MessageChannel",
            id: this.data.get("id"),
          },
          {
            connected: this._connected.bind(this),
            disconnected: this._disconnected.bind(this),
            received: this._received.bind(this),
          }
        );
      }
    
      _connected() {}
    
      _disconnected() {}
    
      _received(data) {
        console.log(data); //Never receive console output
      }
    }

View

    &lt;div data-controller="channel" data-channel-id="&lt;%= @hangout.id %&gt;"&gt;
      &lt;p&gt;
        &lt;strong&gt;Name:&lt;/strong&gt;
        &lt;%= @hangout.name %&gt;
      &lt;/p&gt;
    &lt;div data-target="hangout.messages"&gt;
      &lt;div&gt;
        &lt;div&gt;&lt;strong&gt;&lt;%= message.user.username %&gt;&lt;/strong&gt;: &lt;%= message.body %&gt;&lt;/div&gt;
      &lt;/div&gt;
      &lt;% if current_user.hangouts.include?(@hangout) %&gt;
        &lt;%= form_with model: [@hangout, Message.new] do |form| %&gt;
          &lt;%= form.text_field :body, class: 'shadow appearance-none border rounded w-full py-2 px-3 mt-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline', autofocus: true %&gt;
        &lt;% end %&gt;
      &lt;% else %&gt;
        &lt;%= link_to "Join #{@hangout.name}", hangout_hangout_user_path(@hangout), method: :post %&gt;
      &lt;% end %&gt;
    &lt;/div&gt;

Logfile

    Processing by MessagesController#create as JS
      Parameters: {"authenticity_token"=&gt;"HmM5JMkqlIa2FxvXG7IWcvDaoyxyUfCD8C3vr/N2y/Jg8g4gCcQGtneky/Q98k8CkTEdJ3Ln/RAY0XKOzo6BPA==", "message"=&gt;{"body"=&gt;"dfsdfdsaf"}, "hangout_id"=&gt;"2"}
      [1m[36mUser Load (0.3ms)[0m  [1m[34mSELECT "users".* FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" ASC LIMIT $2[0m  [["id", 1], ["LIMIT", 1]]
      [1m[36mHangout Load (0.3ms)[0m  [1m[34mSELECT "hangouts".* FROM "hangouts" INNER JOIN "hangout_users" ON "hangouts"."id" = "hangout_users"."hangout_id" WHERE "hangout_users"."user_id" = $1 AND "hangouts"."id" = $2 LIMIT $3[0m  [["user_id", 1], ["id", 2], ["LIMIT", 1]]
      ↳ app/controllers/messages_controller.rb:14:in `set_hangout'
      [1m[35m (0.2ms)[0m  [1m[35mBEGIN[0m
      ↳ app/controllers/messages_controller.rb:6:in `create'
      [1m[36mMessage Create (0.7ms)[0m  [1m[32mINSERT INTO "messages" ("hangout_id", "user_id", "body", "created_at", "updated_at") VALUES ($1, $2, $3, $4, $5) RETURNING "id"[0m  [["hangout_id", 2], ["user_id", 1], ["body", "dfsdfdsaf"], ["created_at", "2020-05-11 18:17:56.787486"], ["updated_at", "2020-05-11 18:17:56.787486"]]
      ↳ app/controllers/messages_controller.rb:6:in `create'
      [1m[35m (0.6ms)[0m  [1m[35mCOMMIT[0m
      ↳ app/controllers/messages_controller.rb:6:in `create'
      Rendered messages/_message.html.erb (Duration: 0.7ms | Allocations: 139)
    [ActionCable] Broadcasting to message:: {:message=&gt;"&lt;div&gt;\n  &lt;div&gt;\n    &lt;div&gt;&lt;strong&gt;jem&lt;/strong&gt;: dfsdfdsaf&lt;/div&gt;\n  &lt;/div&gt;\n&lt;/div&gt;\n"}
    No template found for MessagesController#create, rendering head :no_content
    Completed 204 No Content in 78ms (ActiveRecord: 10.9ms | Allocations: 47587)

&amp;#x200B;

I'm sure that there is a problem with my JS, but I don't see it. I appreciate any help!

Thanks!
## [6][friendly_id and slug_candidates](https://www.reddit.com/r/rails/comments/ghqdbp/friendly_id_and_slug_candidates/)
- url: https://www.reddit.com/r/rails/comments/ghqdbp/friendly_id_and_slug_candidates/
---
Hi guys, i'm using the friendly\_id gem (version 5).

I tried to edit it in this way to have an url as title-YY, title-YY-2, title-YY-3, etc

      extend FriendlyId
      friendly_id :slug_candidates, use: [:slugged, :finders]
    
      def slug_candidates
        [:title, :title_and_sequence]
      end
    
      def title_and_sequence
        slug = title.to_param
        sequence = Book.where("slug like #{slug}--%").count + 2
        "#{slug}--#{sequence}"
      end

But I have this error:

    PG::SyntaxError: ERROR: syntax error at or near "YY" LINE 1: SELECT COUNT(*) FROM "books" WHERE (slug like test YY--%) ^ : SELECT COUNT(*) FROM "books" WHERE (slug like test YY--%)

Why? How to fix?
## [7][sending an automated email that has an xlsx attachment through a scheduler job](https://www.reddit.com/r/rails/comments/ghnizc/sending_an_automated_email_that_has_an_xlsx/)
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
## [8][How to handle multiple required strong params](https://www.reddit.com/r/rails/comments/ghhnkv/how_to_handle_multiple_required_strong_params/)
- url: https://www.reddit.com/r/rails/comments/ghhnkv/how_to_handle_multiple_required_strong_params/
---
Hello, what should params for a structure like the following look like?

\`{"person": {"first\_name": "Bob"}, "address": {"street": "123 Main St."}}\`

I've tried several different options based on different posts, including the array form of \`params.require()\`, but ran into different issues no matter which version I used (including 'last permit wins' for one of the two entity's attributes).

So there are two (or more) entities defined in parallel--not one top-level entity that can nest everything beneath itself. Anybody have an example of that?

For the array form, API docs show a parallel assignment, but then trying to \`permit()\` on each one of those param set variables "worked," but I ended up with the "last permit wins" issue mentioned above: only one set of attributes arrived, the other entity's were discarded.
## [9][Create Your First Rails Template For Devise + Bootstrap](https://www.reddit.com/r/rails/comments/gh6leo/create_your_first_rails_template_for_devise/)
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
## [10][Confusion over push notification](https://www.reddit.com/r/rails/comments/gha4p8/confusion_over_push_notification/)
- url: https://www.reddit.com/r/rails/comments/gha4p8/confusion_over_push_notification/
---
Hi guys, I come here in hope for clarification.

I would like to implement in my app push notifications (the one that appear in your phone when you get a new whatsapp). Originally, I did some research and saw that people are using firebase, in what seems to me a very complex sets of step, but I also found a gem, rpush, which seems to be the standard. Are these notifications like a post to a specific routes with params?  


To be honest, after reading all that I am not even sure of what I need of if I have to rely on third services. In mi case I am working toward a rails backend appi and a mobile app with Ionic or similar. Can you help me through?
## [11][why does moving my ActionCable broadcast to a background worker prevent the client side received method from executing?](https://www.reddit.com/r/rails/comments/ghbkka/why_does_moving_my_actioncable_broadcast_to_a/)
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
