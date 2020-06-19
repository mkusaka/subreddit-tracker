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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/har6r7/personal_projects_show_off_your_own_project_andor/)
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
## [3][Would you build your app in Rails in 2020?](https://www.reddit.com/r/rails/comments/hbpmc1/would_you_build_your_app_in_rails_in_2020/)
- url: https://www.reddit.com/r/rails/comments/hbpmc1/would_you_build_your_app_in_rails_in_2020/
---
I have a couple of (what I think are) good ideas floating around in my head, and trying to decide what to build them in.

I've dabbled in Rails, as well as full stack JS, but never built anything **serious** with either.

A few years ago I was really looking into becoming a professional Web Developer. I was working tech support for a web design company but had hit a dead end with my career there. So I was doing the typical self studying thing, Coursera, Codecademy, Hartl's tutorial, etc, getting ready to go down that route, before I kind of fell into an Application Support engineer position with a startup my friend was working in, and I have been doing that ever since. Now I support the same application for a self driving car startup making well over 6 figures.

Now that I am in a pretty good place financially and career wise, and I am a little older and have a little better work ethic, I want to take a serious crack at actually making something out of these ideas I have. Back when I was first thinking about this, a few years ago, I was pretty sure I was going to build them in Rails. Rails was pretty popular then, I was learning about it and found it really easy to use and I really liked how fast you could get something up. But it seems like in the few years I wasn't really paying attention, Rails has really fallen out of vogue, as it were. It doesn't scale. People are listing Ruby as their most hated language.  It's really hard to make anything other than an old fashioned Monolith style app with it. Etc. etc.

If you were starting a brand new project in 2020 one that you were hoping to actually monetize and go commercial with, would still doing it in Rails be a decision that you'd regret later, or is Rails fine, just everyone likes to crap on it because it's not the new hotness anymore?
## [4][Help me to figure out ActiveRecord logic](https://www.reddit.com/r/rails/comments/hbz770/help_me_to_figure_out_activerecord_logic/)
- url: https://www.reddit.com/r/rails/comments/hbz770/help_me_to_figure_out_activerecord_logic/
---
So here's what I want:

* Seller that has many Testimonials;
* Author that has many Testimonials;
* ConfirmedProfile of The User that gets access to all testimonials as Seller and Author, the confirmation would be requested by user and confirmed by Admin.

I am not sure which association types to choose, I believe the polymorphic? and probably I need 'through' association for confirmation part too? Any thoughts?
## [5][Need help understanding module inclusion](https://www.reddit.com/r/rails/comments/hbylab/need_help_understanding_module_inclusion/)
- url: https://www.reddit.com/r/rails/comments/hbylab/need_help_understanding_module_inclusion/
---
I'm trying to understand why I'm not getting the intended results:

I have two models, CustomerAddress and SupplierAddress that each include the module AddressValidation.

In the AddressValidation module, I am trying to determine if I am dealing with a CustomerAddress or a SupplierAddress.

&amp;nbsp;&amp;nbsp;

So in AddressValidation I wanted to have:

**def self.included(base)**

&amp;nbsp;  if base.is_a?(CustomerAddress)

 &amp;nbsp; &amp;nbsp;    [run code 1]

&amp;nbsp;   else

&amp;nbsp; &amp;nbsp;     [run code 2]

&amp;nbsp;   end

**end**

&amp;nbsp;&amp;nbsp;

The problem is, that is_a?(CustomerAddress) always returns false regardless which model is including the module. It also always returns false if I use instance_of?(CustomerAddress)

&amp;nbsp;&amp;nbsp;

So what I did instead is:

**def self.included(base)**

&amp;nbsp;  if base.name == "CustomerAddress"

&amp;nbsp;&amp;nbsp;     [run_code_1]

&amp;nbsp;  else

&amp;nbsp;&amp;nbsp;     [run_code_2]

&amp;nbsp;  end

**end**

&amp;nbsp;&amp;nbsp;

This method works as intended.

So my question is why does calling .name work as intended but not .is_a? or .instance_of?


**EDIT: Formatting**

**EDIT 2** 

Here is the detailed answer to the question, for future reference.

My class CustomerAddress, as far as Ruby is concerned, is actually an instance of the class **Class** and has an attributed called "name" that has the value "CustomerAddress".

Thus base.is_a?(Class) will return true, but nothing else will.

To ensure that it is the right class, I have to either compare its name attribute, or us the equality operator as mentioned below by /u/xire28:

if base == CustomerAddress
## [6][Gems to build order-tracking software?](https://www.reddit.com/r/rails/comments/hbqtiz/gems_to_build_ordertracking_software/)
- url: https://www.reddit.com/r/rails/comments/hbqtiz/gems_to_build_ordertracking_software/
---
Curious what gems you've used to build order-tracking software whether it be to integrate with shipping like UPS/FedEx etc or to track local deliveries.
## [7][Rails 6 multiple databases support in Rails Event Store](https://www.reddit.com/r/rails/comments/hbaxq8/rails_6_multiple_databases_support_in_rails_event/)
- url: https://www.reddit.com/r/rails/comments/hbaxq8/rails_6_multiple_databases_support_in_rails_event/
---
Rails 6 released in August 2019 has brought us several new features. One of the notable changes is support for multiple databases. All details have been described in Rails guides and I’ve read already several blog posts describing how to do it. But how to use this feature to allow Rails Event Store data to be stored in a separate database? Check my [blog post where I document my experiments with it](https://blog.arkency.com/rails-multiple-databases-support-in-rails-event-store/).
## [8][Docker and rails issue](https://www.reddit.com/r/rails/comments/hbney4/docker_and_rails_issue/)
- url: https://www.reddit.com/r/rails/comments/hbney4/docker_and_rails_issue/
---
Hello Folks, i was trying to follow docker guide to setup a new rails app with, but it in rails 5 and i wanna rails 6, is there any guide you can share with me guys.

the docker example with rails 5 : [https://docs.docker.com/compose/rails](https://docs.docker.com/compose/rails/)

**note i don't wanna install rails in my machine or any other gem just like docker docs did!!** 
## [9][Webpacker and Rails relative root for images in SCSS](https://www.reddit.com/r/rails/comments/hbkg6q/webpacker_and_rails_relative_root_for_images_in/)
- url: https://www.reddit.com/r/rails/comments/hbkg6q/webpacker_and_rails_relative_root_for_images_in/
---
Hello, I've got an app that will be accessed via a reverse-proxied subdirectory, eg. 

    https://example.com/subdir -&gt; http://10.0.0.1:3000/

This works (by setting the `RAILS_RELATIVE_URL_ROOT` env var), with helpers like `image_pack_tag` outputting the correct path (eg. "/subdir/packs/media/images/image.jpg")

However, if I reference an image in a Webpacker SCSS file, eg.

    url('../images/background.jpg')

The generated CSS file (and manifest.json) generated by Webpacker will not include the relative root path `/subdir`, eg:

    background: url('/packs/media/images/background.jpg')

So the browser will hit `example.com/packs/media....` instead of `https://example.com/subdir/packs/media...` and get a 404.

If I change `webpacker.yml` to have
 
    public_output_path: subdir/packs

then the paths/manifest will be as desired, but the generated files will now be created in `public/subdir/packs/media...`. meaning the correct URL will now need to be `https://example.com/subdir/subdir/packs/media/...` so it remains broken.

Any thoughts on how I get around this? I've tried `WEBPACKER_RELATIVE_URL_ROOT`, but other than that I'm stumped!

Thanks
## [10][Can you help me debug my super simple controller?](https://www.reddit.com/r/rails/comments/hbhey3/can_you_help_me_debug_my_super_simple_controller/)
- url: https://www.reddit.com/r/rails/comments/hbhey3/can_you_help_me_debug_my_super_simple_controller/
---
I am building a simple blog API. There's a method that I need, but it crashes my whole app. If I uncomment it, the app works fine:

* This is my application controller:

        class ApplicationController &lt; ActionController::API
        
            def token
                request.headers["Authorization"].split(" ")[1]
            end
        
            def secret
                ENV['jwt_secret']
            end
        
            def decoded_token
                JWT.decode(token, secret, true, { algorithm: 'HS256' })
            end
        
            def current_user
                User.find(decoded_token[0]["id"])
            end
        
            def create_token(id, username, email)
                payload = { id: id, username: username, email: email }
                JWT.encode(payload, secret, 'HS256')
            end
        
        end


* And this is my user controller:
        
        class UsersController &lt; ApplicationController
            before_action :get_user, only: [:show, :update, :destroy]
        
            def profile
                render json: current_user
            end
        
            def index
                @users = User.all
                render json: @users
            end
        
            def show
                render json: @user
            end
        
            def create
                @user = User.create(user_params)
                if @user.valid?
                    render json: {token: create_token(user.id, user.username, user.email)}
                else
                    render json: {errors: @user.errors.full_messages}, status: 422
                end
            end
        
            def update
                if @user.update(user_params)
                    render json: @user
                else
                    render json:{errors: @user.errors.full_messages}, status: 422
                end
            end
        
            def destroy
                @user.destroy
            end
        
            private
        
            def get_user
                @user = User.find(params[:id])
            end
        
            def user_params
                params.require(:user).permit(:username, :name, :email, :password)
            end
        
        end


I'm also using the serializer gem:

    class UserSerializer &lt; ActiveModel::Serializer
      attributes :id, :name, :username, :email, :posts, :comments
    
      def posts
        object.posts.size
      end
    
      def comments
        object.comments.size
      end
    
    end

For some reason, if I comment out the `current_user` method in the application controller, the app runs fine, but if I uncomment it, I get all sorts of errors:

        Completed 500 Internal Server Error in 25ms (ActiveRecord: 7.0ms | Allocations: 5463)
        
        
          
        NoMethodError (undefined method `split' for nil:NilClass):
          
        app/controllers/application_controller.rb:4:in `token'
        app/controllers/application_controller.rb:12:in `decoded_token'
        app/controllers/application_controller.rb:16:in `current_user'
        app/controllers/users_controller.rb:14:in `show'
        
I've been trying to fix it for days to no avail, can anyone point me in the right direction?

**EDIT**: You've all been super helpful and I thank you all for that. I'm still not seeing it. I uploaded it into a repo: https://gitlab.com/PMmeYourFlipFlops/help-me-reddit
## [11][How Cloud 66 &amp; OVHcloud makes the deployment of your Rails app in production, easier, faster &amp; cost-effective?](https://www.reddit.com/r/rails/comments/hbh7g1/how_cloud_66_ovhcloud_makes_the_deployment_of/)
- url: https://www.reddit.com/r/rails/comments/hbh7g1/how_cloud_66_ovhcloud_makes_the_deployment_of/
---
\[Webinar\] How Cloud 66 &amp; OVHcloud makes the deployment of your Rails app in production, easier, faster &amp; cost-effective? Register here: [https://register.gotowebinar.com/register/6186841286209414926](https://register.gotowebinar.com/register/6186841286209414926) \#webinar #ovh #cloud66 #deploy #Rails
## [12][Would You Build a Near Real-Time Feed/Channel Using Action Cable?](https://www.reddit.com/r/rails/comments/hay1n8/would_you_build_a_near_realtime_feedchannel_using/)
- url: https://www.reddit.com/r/rails/comments/hay1n8/would_you_build_a_near_realtime_feedchannel_using/
---
I have a newbie question.  I was watching a GoRails episode that shows how Action Cable and Stimulus could be used to [create real-time group chats](https://gorails.com/episodes/rails-group-chat-revised-part-1).

I'm thinking about creating something with feeds on topics - like a Reddit, Twitter, or Slack.

However, the channels/feeds would be around events occurring now. So I'd like to have something that is showing new messages/posts automatically near real-time (the channel could update every 30 seconds to one minute instead of real-time) at the bottom of the screen.

I'd like to build something that has the ability to handle thousands, if not tens of thousands of people viewing and posting on the channels concurrently.

I'm wondering where in the stack I might right into bottlenecks from volume (if any place), if I built it with Action Cable and Stimulus. I'm also wondering how much more expensive it would be than a different approach that may be near real-time (some sort of polling), if it at all.

I'm curious how others would approach building something like this.  Thank you!
