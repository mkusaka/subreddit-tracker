# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gnbebg/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gnbebg/personal_projects_show_off_your_own_project_andor/
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
## [2][Web App Development](https://www.reddit.com/r/rails/comments/gqahgr/web_app_development/)
- url: https://www.reddit.com/r/rails/comments/gqahgr/web_app_development/
---
Hello Everyone. I've been surfing through the internet on how to create a web app but so far I haven't gotten enough Information. 
First Question: Can I create a cross platform web app with Ruby on Rails?
2nd: Can I create the app alone with no other team member?
3rd: Will the back end,frontend and database be created on one the same framework?
## [3][dynamically creating ActionCable subscriptions in asset pipeline](https://www.reddit.com/r/rails/comments/gpx1hi/dynamically_creating_actioncable_subscriptions_in/)
- url: https://www.reddit.com/r/rails/comments/gpx1hi/dynamically_creating_actioncable_subscriptions_in/
---
In expanding that DHH YouTube tutorial for an ActionCable chatroom, I came across a [more fleshed out example](https://www.thegreatcodeadventure.com/rails-5-action-cable-with-multiple-chatroom-subscriptions/) that loops through all existing conversations and creates a subscription for each. This approach "works," but it essentially hardcodes the app's current number of conversations into the asset pipeline, and then leaves it as is until the assets are regenerated. So really, it's not too usable.  

What approaches have you taken to dynamically create ActionCable subscriptions?

**assets/javascripts/channels/messages.coffee**

    &lt;% Conversation.all.each do |conversation| %&gt;
        App['conversation_' + &lt;%= conversation.id %&gt;] = App.cable.subscriptions.create { 
          channel: "MessagesChannel", conversation: &lt;%= conversation.id %&gt; 
        }, 
        {
          connected: -&gt;
            console.log 'Connected'
    
          disconnected: -&gt;
            console.log 'Disconnected'
    
          received: (data) -&gt;
            $("[data-conversation-id='" + data.conversation_id + "']").append(data.message)
            $("#conversation-main").scrollTop($("#conversation-main")[0].scrollHeight);
    
          speak: (body, conversation_id) -&gt;
            @perform 'speak', body: body, conversation_id: conversation_id
    
          set_conversation_id: (conversation_id) -&gt;
            console.log conversation_id
            this.conversation_id = conversation_id
        }
    
      &lt;% end %&gt;
    
      $(document).on 'turbolinks:load', -&gt;
        submit_message()
        scroll_bottom()
    
      submit_message = () -&gt;
        $('#response').on 'keydown', (event) -&gt;
          if event.keyCode is 13
            conversation_id = $("[data-conversation-id]").data("conversation-id")
            # values = $(this).serializeArray()
            App['conversation_' + conversation_id].set_conversation_id(conversation_id)
            App['conversation_' + conversation_id].speak(event.target.value, conversation_id)
            event.target.value = ""
            event.preventDefault()
    
      scroll_bottom = () -&gt;
        if $('#messages').length &gt; 0
          $('#messages').scrollTop($('#messages')[0].scrollHeight)
## [4][Action Text Question](https://www.reddit.com/r/rails/comments/gq5s7r/action_text_question/)
- url: https://www.reddit.com/r/rails/comments/gq5s7r/action_text_question/
---
Anybody happen to know whether Action Text fields take in html code and display it on the view?
## [5][How to properly add Bulma to Rails 6 project](https://www.reddit.com/r/rails/comments/gpq3lt/how_to_properly_add_bulma_to_rails_6_project/)
- url: https://www.reddit.com/r/rails/comments/gpq3lt/how_to_properly_add_bulma_to_rails_6_project/
---
Hello, been hearing great things about the Bulma css framework and decided to try it out this weekend. Started a new Rails 6 app and run "yarn add bulma", then added this line to the application layout file &lt;%= stylesheet\_pack\_tag 'application' %&gt;, then also added the file "application.scss" to the packs directory with the contents as follows:

    import "~bulma";

...and with that, somehow Bulma's styles are not getting pulled...what could I be doing wrong?
## [6][Link to page section.](https://www.reddit.com/r/rails/comments/gpw9e3/link_to_page_section/)
- url: https://www.reddit.com/r/rails/comments/gpw9e3/link_to_page_section/
---
Hoping to add some navbar links that will link to a section of my home page. A little bit of background on this: 

Have welcome#index root path that is to be a portfolio. The portfolio has different sections to it: about, projects, recent articles, contact form, and resume link. I've figured all of it out thus far, but the navbar looks rather empty with just a link to the blog portion (articles/index). So, I'm hoping to add some extra links to the navbar that can auto-scroll to a section of the page. However, I'm at a bit of a loss as to what is best practice for this.

Do I just write up some JS for this and set some section ID's to auto-scroll to? If I am in the blog page and click one of these links, would it still load if I am using some JS for this functionality? Is there a gem that can handle something this simple with little to no marginal cost on performance? What would you do for this?
## [7][How to use react-pdf to rails 6 with active storage?](https://www.reddit.com/r/rails/comments/gpqwxg/how_to_use_reactpdf_to_rails_6_with_active_storage/)
- url: https://www.reddit.com/r/rails/comments/gpqwxg/how_to_use_reactpdf_to_rails_6_with_active_storage/
---
I want to implement react-pdf as previewer. I use active storage to upload pdf then i want to preview.
## [8][Folder Hierarchy for active storage](https://www.reddit.com/r/rails/comments/gpul39/folder_hierarchy_for_active_storage/)
- url: https://www.reddit.com/r/rails/comments/gpul39/folder_hierarchy_for_active_storage/
---
Hi all,

is there any gem available which is recommendable to create a folder hierarchy by the user for a file storage?

I would like to let users upload their document within a document structure created by themselves.

thanks for your tips.
## [9][Rails Misspeled Route](https://www.reddit.com/r/rails/comments/gpqhzj/rails_misspeled_route/)
- url: https://www.reddit.com/r/rails/comments/gpqhzj/rails_misspeled_route/
---
So I have generated a 'niches' controller, but instead of 'niche' I got 'nich' for the routes, how can I fix this?
## [10][Chrome Extension for context on what you read online - We were just featured on Product Hunt!](https://www.reddit.com/r/rails/comments/gppu7x/chrome_extension_for_context_on_what_you_read/)
- url: https://www.reddit.com/r/rails/comments/gppu7x/chrome_extension_for_context_on_what_you_read/
---
Hello Makers and Hunters,

We     just launched WebCheck AI on ProductHunt and have been featured! We    use  NLP to identify named entities on a webpage that the user may  need   more  information on, and allow the user to hover over text and  easily   view  more info, definitions, and relevant news in a tooltip  format.

I would appreciate your feedback and support :)

[https://www.producthunt.com/posts/webcheck-ai](https://www.producthunt.com/posts/webcheck-ai)
## [11][Show Posts by Date in URL](https://www.reddit.com/r/rails/comments/gp2fhi/show_posts_by_date_in_url/)
- url: https://www.reddit.com/r/rails/comments/gp2fhi/show_posts_by_date_in_url/
---
Let's say I want to show only posts for a particular post creation date, which is received from

URL like this (year/month/day): http://localhost:3000/launches/2020/5/24

How do I go about achieving this? 

PS I'm using Rails 6
