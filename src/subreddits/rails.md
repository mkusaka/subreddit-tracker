# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/iui4p8/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/iui4p8/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/iya619/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/iya619/personal_projects_show_off_your_own_project_andor/
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
## [3][Chat App](https://www.reddit.com/r/rails/comments/j05vx9/chat_app/)
- url: https://www.reddit.com/r/rails/comments/j05vx9/chat_app/
---
Hi everyone,

What is the best approach to create a user to user chat app using action cable? 

There are tons of guides out there using chat rooms instead of usernames to create chats. 

I want to create this chat app so users can message each other as long as they know each other's usernames. Wickr like but without encryption lol.

I also need to add a feature to block users. Is it ideal to use devise authentication to achieve this?

The challenge also asks me to keep logs of login and invalid login.

Any advice is appreciated.

 Have a great weekend guys.
## [4][Hiding password in database.yml file?](https://www.reddit.com/r/rails/comments/izucjd/hiding_password_in_databaseyml_file/)
- url: https://www.reddit.com/r/rails/comments/izucjd/hiding_password_in_databaseyml_file/
---
Hi there. I'm wondering if there's a way to either hide the file on github, or stop it from being pushed or what the best way to hide that would be? Since it just shows the password under "default: &amp;default" after adding the username and password, anyone an see it. Is it necessary to hide this password in the first place? I'd greatly appreciate the help!
## [5][Will AWS CloudHSM work with this digital signature script?](https://www.reddit.com/r/rails/comments/izoidm/will_aws_cloudhsm_work_with_this_digital/)
- url: https://www.reddit.com/r/rails/comments/izoidm/will_aws_cloudhsm_work_with_this_digital/
---
[https://gist.github.com/matiaskorhonen/223bd527279cf49bed1e](https://gist.github.com/matiaskorhonen/223bd527279cf49bed1e)

I'm not 100% on how CloudHSM works, can I pull the necessary keys to use with this script using the aws SDK?
## [6][Sending partials over javascript](https://www.reddit.com/r/rails/comments/izbmit/sending_partials_over_javascript/)
- url: https://www.reddit.com/r/rails/comments/izbmit/sending_partials_over_javascript/
---
I'm trying to customize a rails scaffold so it doesnt go to a new page every time I want to edit or create a new resource. 

Right now what I'm doing is using stimulus to prevent default behavior of links, then fetch `e. target.href` to display the form partial within a bootstrap modal. I then use `remote: true` to send the forms w/ javascript and respond with the resource partial like `format.js { render partial: "partial", locals: { resource: @resource }, content_type: "text" } `. I respond with the resource partial because it gets added to the DOM with `insertAdjacentHTML`.

Now that works and all but it feels kinda hacked together and dirty. Is there a better way to do this?
## [7][uploading multiple images to s3 with presign endpoint using shrine and uppy](https://www.reddit.com/r/rails/comments/iz9ux5/uploading_multiple_images_to_s3_with_presign/)
- url: https://www.reddit.com/r/rails/comments/iz9ux5/uploading_multiple_images_to_s3_with_presign/
---
Sorry in advance for the overly broad question, but this isn't about bugs so much as it is about architecture and confusion brought on by trying to twist two docs into one objective. My hope is that someone could help untangle my brain.

I would like to: upload (to s3) and process multiple images using Shrine and Uppy with a nested form in Rails.

However, what I am getting is a POST request that only creates a parent. The child `photos` are neither uploaded to s3 nor saved to the database.

`Processing by ItemsController#create as HTML`

`Parameters: {"utf8"=&gt;"✓", "authenticity_token"=&gt;"[FILTERED]", "files"=&gt;[#&lt;ActionDispatch::Http::UploadedFile:0x00007f9cbc59a050 u/tempfile=#&lt;Tempfile:/var/folders/hj/g_90jx_n7s58m73qlf9h0c4w0000gn/T/RackMultipart20200924-5853-1pb0sdf.jpg&gt;, u/original_filename="in.jpg", u/content_type="image/jpeg", u/headers="Content-Disposition: form-data; name=\"files[]\"; filename=\"in.jpg\"\r\nContent-Type: image/jpeg\r\n"&gt;], "item"=&gt;{"title"=&gt;"dip", "price"=&gt;"900"}, "commit"=&gt;"Submit"}`

`Item Create (10.1ms) INSERT INTO "items" ("created_at", "updated_at", "title", "price", "user_id") VALUES ($1, $2, $3, $4, $5) RETURNING "id" [["created_at", "2020-09-24 22:04:41.333195"], ["updated_at", "2020-09-24 22:04:41.333195"], ["title", "dip"], ["price", 900], ["user_id", 1]]`

The Shrine docs for [multiple uploads](https://shrinerb.com/docs/multiple-files#4b-direct-upload) suggest a model, controller, and view structure that looks like this:

**Models**

    class Photo &lt; ApplicationRecord
     include ImagesUploader::Attachment(:image)
     belongs_to :item
     validates_presence_of :image
    end
    
    class Item &lt; ApplicationRecord
     has_many :photos, dependent: :destroy
     accepts_nested_attributes_for :photos, allow_destroy: true
    end

**Controller**

    class ItemsController &lt; ApplicationController
     def create
      @item = Item.create(item_params)
      @item.save
     end
    
     private
     def item_params
      params.require(:item).permit(:title, photos_attributes: { image: [] })
     end
    end

**View**

    &lt;%= form_for @item, html: { enctype: "multipart/form-data" } do |f| %&gt;
     &lt;%= f.text_field :title %&gt;
     &lt;%= f.fields_for :photos do |i| %&gt;
      &lt;%= i.label :image %&gt;
      &lt;%= i.hidden_field :image, value: i.object.cached_photos_data, class: "upload-data" %&gt;
      &lt;%= i.file_field :image, class: "upload-file" %&gt;
     &lt;% end %&gt;
     &lt;%= file_field_tag "files[]", multiple: true %&gt;
     &lt;%= f.submit "Submit" %&gt;
    &lt;% end %&gt;

Ignoring my confusion over how this double `file_field/file_field_tag` works (and why I only see the latter), I go to [this github wiki](https://github.com/shrinerb/shrine/wiki/Adding-Direct-S3-Uploads) to set up the presign endpoint with an initializer and a route:

**shrine.rb**

    require "shrine"
    require "shrine/storage/file_system"
    require "shrine/storage/s3"
    
    Shrine.plugin :presign_endpoint, presign_options: -&gt; (request) {
      filename = request.params["filename"]
      type     = request.params["type"]
    
      {
        content_disposition:    ContentDisposition.inline(filename), 
        content_type:           type,                                
        content_length_range:   0..(5*1024*1024),                  
      }
    }
    
    Shrine.plugin :activerecord
    Shrine.plugin :cached_attachment_data
    Shrine.plugin :restore_cached_data
    Shrine.plugin :validation
    Shrine.plugin :validation_helpers
    Shrine.plugin :determine_mime_type, analyzer: :marcel
    Shrine.plugin :remove_invalid

**routes.rb**

    get 'presign/s3/params', to: 'presigns#image'

Moved response to a controller for authorization:

    class PresignsController &lt; InheritedResources::Base
     before_action :authenticate_user!
    
     def image
      set_rack_response ImagesUploader.presign_response(:cache, request.env)
     end
    
     private
     def set_rack_response((status, headers, body))
      self.status = status
      self.headers.merge!(headers)
      self.response_body = body
     end
    end

To confirm that this is set up correctly, I visit `/presign/s3/params?filename=test&amp;type=jpg` and see JSON and my AWS credentials in the browser. Although the browser tab says `ActionController: Exception caught`, it seems to work as expected.

Thinking I'm good, I add the uploader and the javascript for handling direct uploads:

**uploader**

    class ImagesUploader &lt; Shrine
      require "image_processing/mini_magick"
      plugin :pretty_location
      plugin :derivatives
    
      s3 = { 
       bucket: ENV['S3_BUCKET_NAME'],
       region: ENV['AWS_REGION'],
       access_key_id: ENV['AWS_ACCESS_KEY_ID'],
       secret_access_key: ENV['AWS_SECRET_ACCESS_KEY']
      }
      Shrine.storages = { 
       cache: Shrine::Storage::S3.new(prefix: "cache", **s3), 
       store: Shrine::Storage::S3.new(prefix: "store", **s3)
      }
    
      Attacher.validate do
       validate_min_size	10.kilobytes	
       validate_max_size	5.megabytes, message: 'must be smaller than 5MB'
       validate_mime_type 	%w[image/jpeg image/png]
       validate_extension 	%w[jpg jpeg png]
      end
    
      Attacher.derivatives do |original|
       magick = ImageProcessing::MiniMagick.source(original)
       {
    	large: magick.resize_to_limit!(1000, 1000),
    	small: magick.resize_and_pad!(225, 220)
       }
      end
    end

**app/javascript/fileUpload.js**

    import 'uppy/dist/uppy.min.css'
    
    import {
      Core,
      FileInput,
      Informer,
      ProgressBar,
      ThumbnailGenerator,
      AwsS3,
    } from 'uppy'
    
    function fileUpload(fileInput) {
      const hiddenInput = document.querySelector('.upload-data'),
            imagePreview = document.querySelector('.upload-preview img'),
            formGroup = fileInput.parentNode
    
      formGroup.removeChild(fileInput)
    
      const uppy = Core({
          autoProceed: true,
        })
        .use(FileInput, {
          target: formGroup,
        })
        .use(Informer, {
          target: formGroup,
        })
        .use(ProgressBar, {
          target: imagePreview.parentNode,
        })
        .use(ThumbnailGenerator, {
          thumbnailWidth: 600,
        })
        .use(AwsS3, {
          companionUrl: '/presign/s3/params', // i set this to my presign endpoint, which I checked by visting site.com/presign/s3/params?filename=test&amp;type=jpg
        })
    
      uppy.on('thumbnail:generated', (file, preview) =&gt; {
        imagePreview.src = preview
      })
    
      uppy.on('upload-success', (file, response) =&gt; {
        const uploadedFileData = {
          id: file.meta['key'].match(/^cache\/(.+)/)[1],
          storage: 'cache',
          metadata: {
            size: file.size,
            filename: file.name,
            mime_type: file.type,
          }
        }
    
        hiddenInput.value = JSON.stringify(uploadedFileData)
      })
    }
    
    export default fileUpload

**javascript/packs/application.js**

    import fileUpload from 'fileUpload'
    
    document.addEventListener('turbolinks:load', () =&gt; {
      document.querySelectorAll('.upload-file').forEach(fileInput =&gt; {
        fileUpload(fileInput)
      })
    })

Some questions.

1. How does that double `file_field` work and why is it that the stray `file_field_tag` is the one that submits the files? I do not intuitively grasp how the two are connected. If the tag is named `files[]`, shouldn't my whitelisted params be something like `photos_attributes: { files: [] }`?
2. If uppy and shrine upload the files to an s3 cache, how do I get that information into my controller? Presumably I would get a JSON response. The doc has this to say, but I'm confused by what it means:

&gt;Once files are uploaded asynchronously, you can dynamically insert photo attachment fields for the `image` attachment attribute into the form, where the hidden field is filled with uploaded file data in JSON format, just like when doing single direct uploads. The attachment field names should be namespaced according to the convention that the nested attributes feature expects. In this case it should be `album[photos_attributes][&lt;uid&gt;]`, where `&lt;uid&gt;` is any unique string.
## [8][Relational DB tutorial?](https://www.reddit.com/r/rails/comments/iz6o8m/relational_db_tutorial/)
- url: https://www.reddit.com/r/rails/comments/iz6o8m/relational_db_tutorial/
---
Anyone have a good resource to understand Relational DB linking in rails? 

I’ve read the docs, handheld a few tutorials and just cannot wrap my brain around how to properly understand it.
## [9][Auto-save the current form without refresh](https://www.reddit.com/r/rails/comments/iyyhzh/autosave_the_current_form_without_refresh/)
- url: https://www.reddit.com/r/rails/comments/iyyhzh/autosave_the_current_form_without_refresh/
---
Is it possible to use AJAX to auto-save the form that's currently in view, but leave the default submit button active?  


I'm using wicked, and I have a fairly large form that needed to be sliced into pieces. Now, I want to save each individual form step with autosave using a timed function. Is it possible?
## [10][How to improve my tests?](https://www.reddit.com/r/rails/comments/iyz9e7/how_to_improve_my_tests/)
- url: https://www.reddit.com/r/rails/comments/iyz9e7/how_to_improve_my_tests/
---
Hi! I'm in the process of writing tests for a new Rails app.

This is one my controller tests: [Gist](https://gist.github.com/ImMaax/6bd4802cb283e1cf4fe69bcd0a1f3eaf)

Even though this test works perfectly fine, it just feels pretty wrong, as I repeat myself a lot and I'm unsure whether that's "the right way" to write a controller test. I've also seen integration tests, but I'm not quite sure where to put which tests now. Can someone help me? How would you optimize those tests?
## [11][Indie Hacker : a possible meaning](https://www.reddit.com/r/rails/comments/iyyv2q/indie_hacker_a_possible_meaning/)
- url: https://www.reddit.com/r/rails/comments/iyyv2q/indie_hacker_a_possible_meaning/
---
I just wrote an article about what it means to be an Indie Hacker, in a dev point of view.  
Spoiler : I use Rails for this kind of project :) All thought are based on experience, which means very subjective and open to criticism :) [http://bdavidxyz.com/blog/indie-hacker-meaning/](http://bdavidxyz.com/blog/indie-hacker-meaning/)
## [12][automatically getting more inputs and how to handle them in the backend](https://www.reddit.com/r/rails/comments/iynsdf/automatically_getting_more_inputs_and_how_to/)
- url: https://www.reddit.com/r/rails/comments/iynsdf/automatically_getting_more_inputs_and_how_to/
---
I know this might be more of a frontend issue, but I have the habit of having anything in one place (I think this might be a harmful thought, but I do most of my projects alone, so I keep them the way I can handle!). So, I'm a musician and when I was uploading an album on bandcamp, I noticed something. 

When you use "album", the damn thing automatically gives you a field of new song. I thing that be a really good feature to have, but I don't know what that's called and how I can have on that a rails project. If you guys have any experiences, I'd love to know. 

Thanks!
