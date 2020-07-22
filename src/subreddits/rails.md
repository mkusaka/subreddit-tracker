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
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/hrnm2o/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/hrnm2o/personal_projects_show_off_your_own_project_andor/
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
## [3][[GUIDE] Fomantic / Semantic UI from Scratch in Rails 6 with Webpacker and Sprockets](https://www.reddit.com/r/rails/comments/hvrjzf/guide_fomantic_semantic_ui_from_scratch_in_rails/)
- url: https://www.reddit.com/r/rails/comments/hvrjzf/guide_fomantic_semantic_ui_from_scratch_in_rails/
---
I have spent the last several hours trying to use Fomantic / Semantic in my Rails 6 app from NPM. That means that I could theme it (can't use fomantic-ui-sass gem for themes) and also not rely on gems when adding extra JS modules.

Unfortunately, it sucks and it's a pain in the ass. I can see why frontend devs eat so many crayons now.


1. `yarn add fomantic-ui`
2. `cd node_modules/fomantic-ui/`
3. `gulp`
4. At this stage, it will ask you for a default path directory. I put mine in `app/javascript/semantic` as that's the easiest way to do it.
5. `cd ../../app/javascript/semantic`
6. `gulp build`. You can try changing paths in semantic.json but it won't do anything other than fuck up your CSS paths even more (:
7. In `app/javascript/packs/application.js` add the lines:
* `require('semantic/dist/semantic.min.js')` (or full JS if you want, I'm not a cop)
* `import 'stylesheets/application'`
8. In `app/javascript/stylesheets/application.scss` add the line `@import "../semantic/dist/semantic.min.css"`
9. In `app/views/layouts/application.html.erb` change the line that has `stylesheet_link_tag` to `stylesheet_pack_tag`
* Note: this needs to call the same filename as step 8. I.e if it says 'application' after 'stylesheet_pack_tag' then you need to have a file called `app/javascript/stylesheets/application.scss`
10. Now here comes the magic. Remember step 6 and how we built all the crap? Go back into that directory - `cd /app/javascript/semantic`
11. You need to move `dist/themes` to `../stylesheets/` -&gt; `mv dist/themes ../stylesheets/`
* I setup a simple bash script to do this for me in case I forget when regenerating.
12. Hopefully, when you run rails server now you will have a functioning theme



I don't do front-end stuff and I spent far too long on this.
I've tried so many things it's not funny. `resolve-url-loader` doesn't work, setting the path in `webpacker.yml` doesn't work (for image resources), changing `semantic.json` output path doesn't work.
This solution isn't great because it still relies on manually moving a folder and the root issue of botched url paths isn't fixed. If someone knows how to fix that, please let me know. Also, in the web page the CSS seems to all load at once in the tag? Like it's not in a file. Not sure what that's about.


Apologies that I haven't explained any steps. If I feel less defeated tomorrow, I'll edit the post with some more information.


Edit: the Javascript doesn't appear to work as JQuery is not loading. I added `yarn add jquery popper.js` and then the stuff to initialize it in `app/config/webpacker/environment.js` but it doesn't appear to be loading :/


Edit 2: Yolo, I added ```import JQuery from 'jquery';
window.$ = window.JQuery = JQuery;``` to `app/javascript/packs/application.js`
## [4][Qualys/penetration test - Path disclosure vulnerabilities](https://www.reddit.com/r/rails/comments/hvt3r0/qualyspenetration_test_path_disclosure/)
- url: https://www.reddit.com/r/rails/comments/hvt3r0/qualyspenetration_test_path_disclosure/
---
Hi all, recently one of my clients, do a Qualys test on the application and the test shows up that there is a 5 “Path disclosure” vulnerabilities and indicate to 5 paths, then he did a second-round of the test and on the second round, the test shows up also 5 “Path disclosure” vulnerabilities but indicate to different paths, even if I didn’t take any action to resolve it. But now I should resolve all of them and I’m not sure how.  


Report from Qualys test:  
**Threat -** A potentially sensitive file, directory, or directory listing was discovered on the Web server. 

**Impact -** The contents of this file or directory may disclose sensitive information. 

**Solution  -** Verify that access to this file or directory is permitted. If necessary, remove it or apply access controls to it. 

&amp;#x200B;

Rails - 5.2.2, core JS (there is no API), deployed to the Heroku server.

Does someone familiar with this kind of vulnerabilities or testing? I will be very appreciative of someone's help.
## [5][Soft delete Active Text in Rails 6](https://www.reddit.com/r/rails/comments/hvecjx/soft_delete_active_text_in_rails_6/)
- url: https://www.reddit.com/r/rails/comments/hvecjx/soft_delete_active_text_in_rails_6/
---
I added ActionText to my Post to have the \`content\` be rich text but when I soft-delete a post the \`content\` is deleted. If I restore the post the content is no longer there.

&amp;#x200B;

This is my Post model

&amp;#x200B;

    class Post &lt; ApplicationRecord
        acts_as_paranoid
        has_rich_text :content
        extend FriendlyId
        friendly_id :title, use: :slugged
        validates_presence_of :title, :content
        belongs_to :user, -&gt; { with_deleted }
    end

&amp;#x200B;

I run: `rails action_text:install`

&amp;#x200B;

This are the migrations Action Text generated

    class CreateActiveStorageTables &lt; ActiveRecord::Migration[5.2]
      def change
        create_table :active_storage_blobs do |t|
          t.string   :key,        null: false
          t.string   :filename,   null: false
          t.string   :content_type
          t.text     :metadata
          t.bigint   :byte_size,  null: false
          t.string   :checksum,   null: false
          t.datetime :created_at, null: false
    
          t.index [ :key ], unique: true
        end
    
        create_table :active_storage_attachments do |t|
          t.string     :name,     null: false
          t.references :record,   null: false, polymorphic: true, index: false
          t.references :blob,     null: false
    
          t.datetime :created_at, null: false
    
          t.index [ :record_type, :record_id, :name, :blob_id ], name: "index_active_storage_attachments_uniqueness", unique: true
          t.foreign_key :active_storage_blobs, column: :blob_id
        end
      end
    end

&amp;#x200B;

&amp;#x200B;

    class CreateActionTextTables &lt; ActiveRecord::Migration[6.0]
      def change
        create_table :action_text_rich_texts do |t|
          t.string     :name, null: false
          t.text       :body, size: :long
          t.references :record, null: false, polymorphic: true, index: false
    
          t.timestamps
    
          t.index [ :record_type, :record_id, :name ], name: "index_action_text_rich_texts_uniqueness", unique: true
        end
      end
    end

&amp;#x200B;

&amp;#x200B;

How can I soft delete the rich text? or at least not deleted completely?
## [6][Looking for a mentor/coding buddies](https://www.reddit.com/r/rails/comments/hv74j4/looking_for_a_mentorcoding_buddies/)
- url: https://www.reddit.com/r/rails/comments/hv74j4/looking_for_a_mentorcoding_buddies/
---
Hi, everyone. Over the past 5 years I've been working as a solo in-house dev for two companies. In the first role I developed a CMS-backed corporate website, and in the second I developed a web app. Both projects are Rails-based and both are up and running in production (but the CMS and app parts are not public).

While I'm comfortable with getting a Rails app to production, there are many aspects of Rails and webdev in general that I'm not familiar with due to being a solo dev.

If any experienced developers would be willing to give me some guidance on how things are done in bigger companies/teams or in enterprise level apps, I'd really appreciate the help. As some examples, I'd like to know more about user authentication (e.g. a Rails API backend with a SPA frontend), error handling and TDD.

I'd also really enjoy chatting with anyone who is in a similar position as me, as I don't really know anyone else who codes with Rails and I know how frustrating and lonely it can be as a solo dev.

A bit more about me: I'm an Australian living in Hong Kong and apart from webdev, I research computational linguistics and NLP.

Feel free to get in touch.
## [7][ActionController::ParameterMissing: param is missing or the value is empty: photo](https://www.reddit.com/r/rails/comments/hvdxpk/actioncontrollerparametermissing_param_is_missing/)
- url: https://www.reddit.com/r/rails/comments/hvdxpk/actioncontrollerparametermissing_param_is_missing/
---
Hi all,  


I'm experiencing an error when I try to submit a form WITHOUT a photo attached, after setting up photo uploads using AWS. It's all working WITH the photo, but now when I submit the form without a photo, an error is thrown.  


Here's what my params look like when the form is submitted:  


    =&gt; &lt;ActionController::Parameters {"utf8"=&gt;"✓", "authenticity_token"=&gt;"1N5+BJSCBeKChrX73gZ5IKSh6V3ynIM83XHvWPKb94Iw7QZocgforAZ9L9NrmwVF54es+tZBNut1c8jbGo/KQw==", "venue"=&gt;&lt;ActionController::Parameters {"name"=&gt;"St. Stephens Tavern", "address_attributes"=&gt;&lt;ActionController::Parameters {"address"=&gt;"10B Bridge St (at Canon Row), London, Greater London, SW1A 2JR, United Kingdom", "longitude"=&gt;"-0.1255105271062834", "latitude"=&gt;"51.50104512723709"} permitted: false&gt;, "category"=&gt;"Pub", "description"=&gt;"", "has_wifi"=&gt;"0", "wifi_restrictions"=&gt;"0", "foursquare_id"=&gt;"4ac518c5f964a520d4a420e3"} permitted: false&gt;, "commit"=&gt;"Submit", "controller"=&gt;"venues", "action"=&gt;"create"} permitted: false&gt;

Here's my strong params method:

    def new_venue_photo_params
      params.require(:venue).require(:photo).require(:image)
    end

Now, I have tried playing around with every combination of "require" and "permit", but the only way I've managed to get the photo uploading to work is with the above combination.  


If I change the strong params to:

    params.require(:venue).permit(:photo).permit(:image)

I get the error (regardless of whether an attachment is submitted or not):

    Could not find or build blob: expected attachable, got &lt;ActionController::Parameters {} permitted: true&gt;

  
If I change them to:

    params.require(:venue).require(:photo).permit(:image)

I get the error (regardless of whether an attachment is submitted or not):

    Could not find or build blob: expected attachable, got &lt;ActionController::Parameters {"image"=&gt;#&lt;ActionDispatch::Http::UploadedFile:0x00007f94f3732af0 @tempfile=#&lt;Tempfile:/var/folders/_t/mz_hqns15gd6bhhmj6_qm_w40000gn/T/RackMultipart20200721-29127-7eynw6.png&gt;, @original_filename="stack_to_the_future_profile_image_v02.png", @content_type="image/png", @headers="Content-Disposition: form-data; name=\"venue[photo][image]\"; filename=\"stack_to_the_future_profile_image_v02.png\"\r\nContent-Type: image/png\r\n"&gt;} permitted: true&gt;

  
So basically the only combination that works for image attachments (but throws an error when there is no attachment) is:

    params.require(:venue).require(:photo).require(:image)

  
Here's how I'm saving my photo in the create action:

    @venue = Venue.new(new_venue_params)
    if new_venue_photo_params
      @photo = Photo.new(imageable: @venue)
      @photo.image.attach(new_venue_photo_params)
      @photo.save
    end

  
Pretty lost here so I apologise for the somewhat confusing post, but hopefully someone can identify where I'm going wrong here.  


Thanks.
## [8][Using external image location in my application. What should I be concern ?](https://www.reddit.com/r/rails/comments/hvdy2b/using_external_image_location_in_my_application/)
- url: https://www.reddit.com/r/rails/comments/hvdy2b/using_external_image_location_in_my_application/
---
I have been using a web application which if you want to share a image or video you have to paste the image location into the application in order to show on a post.

I have a comic tracking application that I have been building for practice and it's pretty simple. The application does not have a database with a bunch of comics data. The user add manually a new comic to his collection.

I would like to display the comics covers in my application. I test out and 99% of the time I could display the correct cover from given comic if I make a search on Google Images using the publisher company, the comic's name and the volume.

So I build a function to build a google search URL using this information from each comic. Using rest-client and nokogiri gems I could retrieve the first image from the results search which is 99% the correct image that I want display.

I would like to know if this solution would be pretty bad or not as I think.

Making these request for let's say about 500 comics  each time the user load his collection or someone else it's seems pretty inefficient way to handle things. I was thinking if using Inifity Scroll to display only about 30 comics and prevent from overloading the application doing 300,400, requests to retrieve the comic covers.
## [9][Trigger spinner on button click until action completed](https://www.reddit.com/r/rails/comments/hvdbei/trigger_spinner_on_button_click_until_action/)
- url: https://www.reddit.com/r/rails/comments/hvdbei/trigger_spinner_on_button_click_until_action/
---
Hi, I've got a rails app where I'm looking to trigger a spinner on a button-click that'll show some message like "Processing" until the action/redirect/call has been made by the rails controller. Tried doing something like the below code but it actually blocks the controller action from taking place. Any thoughts?
    
    &lt;button class="btn btn-outline-success btn-mobile async-disable"&gt;Submit&lt;/button&gt;


    $(document).on("turbolinks:load", function() {
        $('.async-disable').on('click', function() {
            $(this).prop("disabled", true);
            // add spinner to button
            $(this).html(`&lt;span class="spinner-border spinner-border-sm" role="status" aria-hidden="true"&gt;&lt;/span&gt; Processing...`);
        });
    });
## [10][Question regarding queueing tasks to node](https://www.reddit.com/r/rails/comments/hvdy2w/question_regarding_queueing_tasks_to_node/)
- url: https://www.reddit.com/r/rails/comments/hvdy2w/question_regarding_queueing_tasks_to_node/
---
I have a Rails app where users can request a report to be generated AND the report gets generated automatically once every x days.

Currently, I send a POST request to an Express (node) endpoint that runs its audit and then sends the results back. I'm not sure what the best route is, to queue up these different tasks and make sure:

1. too many aren't running at once (they need to "wait their turn")
2. the http response doesn't time out because the node endpoint isn't ready yet
3. etc

Does it make sense to have a secret key shared between both services where Rails basically goes, "Hey node, run your report when you have time. Respond with this ID when you're done." and node goes "Ok, I'll add it to my queue". Once node is done, it uses the secret key to send a request back to the Rails app with the results. Rails then stores the results in the database and marks the report/audit as "complete". This would mean that the real meat and potatoes of the job queueing would be on the node end.

What other options might there be?
## [11][Create a loading state for your async buttons with rails &amp; css](https://www.reddit.com/r/rails/comments/hv9z6q/create_a_loading_state_for_your_async_buttons/)
- url: https://www.reddit.com/r/rails/comments/hv9z6q/create_a_loading_state_for_your_async_buttons/
---
Hi - I have a form that gets submitted by users which then triggers a number of API calls. I don't want the user to be able to click the submit button again and so I was wondering what the best approach to this was. Was thinking of potentially using a bootstrap spinner... any thoughts?
## [12][Is it true that Rails can be harder to debug?](https://www.reddit.com/r/rails/comments/hv9sox/is_it_true_that_rails_can_be_harder_to_debug/)
- url: https://www.reddit.com/r/rails/comments/hv9sox/is_it_true_that_rails_can_be_harder_to_debug/
---
Compared to other frameworks, I’ve heard people complain that because of all of the stuff Rails does for you (the “magic”), it makes it harder to debug issues than with other frameworks.
