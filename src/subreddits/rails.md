# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/hja8kx/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/hja8kx/personal_projects_show_off_your_own_project_andor/
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
## [2][Draper decoration seems to disappear..? I've looked at this took long.](https://www.reddit.com/r/rails/comments/hp9xnz/draper_decoration_seems_to_disappear_ive_looked/)
- url: https://www.reddit.com/r/rails/comments/hp9xnz/draper_decoration_seems_to_disappear_ive_looked/
---
Hey all -

I'm new to using Draper, and I'm really confused by this behaviour below.

I have a really simple setup:

Post model, belongs\_to :user. Has a decorator, decorates\_association :user.

User model, has\_many :posts, also has a decorator.

In my PostController I declare `decorates_assigned :post`, and all I do in #edit is assign `@post`.

This is where my confusion starts.

From the `post` variable (the decorated Post object) I'm able to get a collection of peer post objects, and then I want to render them all through a partial. So I did it like this:

    render partial: 'posts/post', collection: posts

And the \_post.html.haml partial displays the post with some author details (using `post.user`), and importantly uses one of the User decorator methods: `UserDecorator#full_name` as `post.user.full_name`.

When I do it this way I get an error: #full\_name not found on the User object. Other User model methods work fine, so I guess my User is no longer decorated?

Out of curiosity, I switched to this in my top level view where the partial is called - I enumerate the collection rather than getting the rails view helper to do it:

    - posts.each do |p|
      = render partial: 'posts/post', post: p

AND THAT WORKS! In the partial, the call to UserDecorator#full\_name works perfectly! Literally the only change I made was to introduce the use of Enum#each in the top level view and it works.

Can someone please explain that to me? I'm totally confused.

Thanks!
## [3][Retrieve the most liked article](https://www.reddit.com/r/rails/comments/hp8npu/retrieve_the_most_liked_article/)
- url: https://www.reddit.com/r/rails/comments/hp8npu/retrieve_the_most_liked_article/
---
So I have add a Like feature on my blog app and now I'm trying to retrieve the most liked article to display on my home page.

I have been following this documentation on active record calculations [here](https://api.rubyonrails.org/classes/ActiveRecord/Calculations.html).

I have a Like model which belongs\_to an article and belongs\_to an user.

I don't know if this is the issue but I tried to retrieve the most liked article by doing Article.maximum(:likes) but it says the column "likes" does not exist. And yeah its perfectly fine since the article I followed up does not implemented that way.

The thing is I think I should check the most liked article by using my Like model  and check who is the article\_id that most appears. I couldn't grasp yet how to do that. It always results in error whatever I try.

My [repo](https://github.com/Gregory280/alpha-blog-5.1.4)

By the way, should I put this on my view or on a "popular" method in likes controller?
## [4][random GET request for "/.git/HEAD"](https://www.reddit.com/r/rails/comments/hp22vr/random_get_request_for_githead/)
- url: https://www.reddit.com/r/rails/comments/hp22vr/random_get_request_for_githead/
---
 I just checked the console for my app and saw this:

    2020-07-10T01:26:50.166768+00:00 app[web.1]: Started GET "/.git/HEAD" for 88.99.161.235 at 2020-07-10 01:26:50 +0000
    2020-07-10T01:26:50.169376+00:00 app[web.1]: 
    2020-07-10T01:26:50.169428+00:00 app[web.1]: ActionController::RoutingError (No route matches [GET] "/.git/HEAD"):

The route doesn't exist on my site because I have zero need to create it, and I see no indication of anyone else with similar requests. Is this request for `/.git/HEAD` *shady business* or is it nothing to worry about?
## [5][Running Rails App inside Vagrant](https://www.reddit.com/r/rails/comments/hovx43/running_rails_app_inside_vagrant/)
- url: https://www.reddit.com/r/rails/comments/hovx43/running_rails_app_inside_vagrant/
---
I am trying to run a working app inside vagrant file, but when I hit localhost:3000 I get default rails error:

'The page you were looking for doesn't exist.'

I assume I need to set domain here:

    class ApplicationController &lt; ActionController::Base
    ...
    
    def set_current_website
        @current_website = Website.where(domain: request.domain).first!
    end
    
    ...

how do I pass the domain name to the app if I am running it on localhost and not on an actual server with a domain?

&amp;#x200B;

EDIT:

SOLUTION: I have created website with name 'localhost' and it worked. Thanks for all who helped ;)
## [6][Creating form with calculated fields](https://www.reddit.com/r/rails/comments/hos1c1/creating_form_with_calculated_fields/)
- url: https://www.reddit.com/r/rails/comments/hos1c1/creating_form_with_calculated_fields/
---
I am a bit of a noobie with rails so bear with me...

I want to create a simple webpage with some fields. Let's say there are just two of them for the sake of learning. Field one is Cost1, and field two is Cost2. I want a total, Total.

In the first instance I don't need to store the data in a database, although in future, I might want to change that.

How would you do this in Ruby on Rails? Would you have to click a Submit button or can it instantly update via some schenanegans?
## [7][Accept TCP connections from remote devices in a Rails 5 project](https://www.reddit.com/r/rails/comments/hoqc47/accept_tcp_connections_from_remote_devices_in_a/)
- url: https://www.reddit.com/r/rails/comments/hoqc47/accept_tcp_connections_from_remote_devices_in_a/
---
Hello all,

I have a Rails 5 API-only project that deals with IRL physical devices that control remote farming equipment. The remote devices offer 2 methods of communication. I was originally using one that talked over HTTP, but there is a bug with the devices where over longer periods of time, they "stall" if using only 1 method and begin to not send/receive payloads until you physically turn them on and off. Apparently this is either technically or economically unfixable by the device manufacturer, so they suggested that I move to a dual-communication strategy where I send outbound messages via HTTP and receive the inbound messages via TCP.

I have never worked with TCP connections before beyond a sending a few notifications with ActionCable, and don't really know where to start.

I currently receive the string message the device sends back in a webhook, save the message to the DB, and fire off a series of background processes using the stored message. I shouldn't have to change much there, the biggest change that needs to occur is how the message enters my system to begin with.

I need to provide the devices with a TCP endpoint they can hit to send their payloads. How do I receive a TCP connection in a Rails 5 project? I'm assuming I can use all of the normal Rails functionality from a TCP connection (saving a model to the DB, firing off a background process, etc). I have never done such a thing before and there's not much guidance online I could find for implementing such a thing.
## [8][After an action is executed trigger a download using rails](https://www.reddit.com/r/rails/comments/hov9a8/after_an_action_is_executed_trigger_a_download/)
- url: https://www.reddit.com/r/rails/comments/hov9a8/after_an_action_is_executed_trigger_a_download/
---
I'm looking for a user to input some information (ex: email, phone number, etc) in a form and after they hit the submit button it triggers a pdf download in their browser. Are there any common architectures or gems I can use to accomplish this?
## [9][How to implement refresh tokens?](https://www.reddit.com/r/rails/comments/hohl6i/how_to_implement_refresh_tokens/)
- url: https://www.reddit.com/r/rails/comments/hohl6i/how_to_implement_refresh_tokens/
---
Hey, I current use devise-jwt with an access token that expires in 2 weeks. I want to reduce the time it takes for access tokens to live, but I also don't want users logging in every 30 minutes. devise-jwt doesn't support refresh tokens out of the box. How do you folks do it?
## [10][Vagrant alternatives? Coding environment failures](https://www.reddit.com/r/rails/comments/hoeo0a/vagrant_alternatives_coding_environment_failures/)
- url: https://www.reddit.com/r/rails/comments/hoeo0a/vagrant_alternatives_coding_environment_failures/
---
Hi there. I'm just fed up with continuing trying to use Vagrant at this point. I've had many troubles in trying to build my own Vagrantbox for my Rails web dev environment. I had one going from what my bootcamp gave me, basically just a small file system with their vagrant box they provided and their own instructions for things. But sine a couple months ago I was trying to branch out and see what else is out there, and I think initially my issues started when I tried to use Docker. That's a whole other thing I won't get into, but will still probably try again sometime cause I was having issues there too. 

Anyway, I'm now kind of at my last resort using this box using Ubuntu 20.04 that's at this github link:  [https://github.com/rails/rails-dev-box](https://github.com/rails/rails-dev-box) 

I cloned it, and let it do its thing. But now it's getting hung up on the default: SSH auth method: private key line. Really this has been my main issue after setting up the box. The first time I run vagrant up, it gets hung up here. I've tried running ssh-keygen again to replace my keys in my .ssh folder, but it doesn't do anything. I'm still a newb though and may be missing something. 

Here's my full first "vagrant up" output: 

    C:\Users\Kyle_\Dev\rails-dev-box&gt;vagrant up
    Bringing machine 'default' up with 'virtualbox' provider...
    ==&gt; default: Box 'ubuntu/focal64' could not be found. Attempting to find and install...
        default: Box Provider: virtualbox
        default: Box Version: &gt;= 0
    ==&gt; default: Loading metadata for box 'ubuntu/focal64'
        default: URL: https://vagrantcloud.com/ubuntu/focal64
    ==&gt; default: Adding box 'ubuntu/focal64' (v20200707.0.0) for provider: virtualbox
        default: Downloading: https://vagrantcloud.com/ubuntu/boxes/focal64/versions/20200707.0.0/providers/virtualbox.box
    Download redirected to host: cloud-images.ubuntu.com
        default:
    ==&gt; default: Successfully added box 'ubuntu/focal64' (v20200707.0.0) for 'virtualbox'!
    ==&gt; default: Importing base box 'ubuntu/focal64'...
    ==&gt; default: Matching MAC address for NAT networking...
    ==&gt; default: Checking if box 'ubuntu/focal64' version '20200707.0.0' is up to date...
    ==&gt; default: Setting the name of the VM: rails-dev-box_default_1594339592541_33376
    ==&gt; default: Clearing any previously set network interfaces...
    ==&gt; default: Preparing network interfaces based on configuration...
        default: Adapter 1: nat
    ==&gt; default: Forwarding ports...
        default: 3000 (guest) =&gt; 3000 (host) (adapter 1)
        default: 22 (guest) =&gt; 2222 (host) (adapter 1)
    ==&gt; default: Running 'pre-boot' VM customizations...
    ==&gt; default: Booting VM...
    ==&gt; default: Waiting for machine to boot. This may take a few minutes...
        default: SSH address: 127.0.0.1:2222
        default: SSH username: vagrant
        default: SSH auth method: private key
    Timed out while waiting for the machine to boot. This means that
    Vagrant was unable to communicate with the guest machine within
    the configured ("config.vm.boot_timeout" value) time period.
    
    If you look above, you should be able to see the error(s) that
    Vagrant had when attempting to connect to the machine. These errors
    are usually good hints as to what may be wrong.
    
    If you're using a custom box, make sure that networking is properly
    working and you're able to connect to the machine. It is a common
    problem that networking isn't setup properly in these boxes.
    Verify that authentication configurations are also setup properly,
    as well.
    
    If the box appears to be booting properly, you may want to increase
    the timeout ("config.vm.boot_timeout") value.
    
    C:\Users\Kyle_\Dev\rails-dev-box&gt;
    

Any help would be great.  And honestly, if I can't get this to work, I would love other suggestions on how to set up an environment without Vagrant. I know I should still use VMs with Linux for Rails development. I was perfectly comfortable during bootcamp just vagrant up, vagrant ssh, and boom. IDK what happened and it's so discouraging :(
## [11][Efficiently searching for text from several .pdf](https://www.reddit.com/r/rails/comments/ho8g9m/efficiently_searching_for_text_from_several_pdf/)
- url: https://www.reddit.com/r/rails/comments/ho8g9m/efficiently_searching_for_text_from_several_pdf/
---
Hi,  
A friend of mine needs a search engine that searches from a lot of PDF. I though it could be a great challenge (I'm not a professional at all). I'm looking for advise since I've never deal with such a huge amount of data.  


Here's my plan:

* allow him to upload as many .pdf as he wants
* extract text from PDF using [this gem](https://github.com/yob/pdf-reader) (pdf-reader) and async jobs
* store extracted text into a database
* set up Elasticsearch to search from extracted text (never done that before)

Beyond the challenge, if there's any working (not necessarily online) tool, I'd glad to test it and share it with him

&amp;#x200B;

Thank you in advance ! :)
