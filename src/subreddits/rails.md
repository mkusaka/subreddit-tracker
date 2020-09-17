# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/ipfxw8/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/ipfxw8/personal_projects_show_off_your_own_project_andor/
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
## [2][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/iui4p8/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/iui4p8/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [3][Zipping active storage has_one_attached &amp; has_many_attached files in one zip file for a model](https://www.reddit.com/r/rails/comments/iui5g4/zipping_active_storage_has_one_attached_has_many/)
- url: https://www.reddit.com/r/rails/comments/iui5g4/zipping_active_storage_has_one_attached_has_many/
---
I'm using the active\_storage-send-zip, but it seems to only focus on has\_many\_attached.  I need to zip all files from the model at once. I can't seem to combine the two, to get the results that I want.

&amp;#x200B;

For example:

User.rb

has\_one\_attached :avatar

has\_many\_attached :photos

&amp;#x200B;

zip all photos + their avatar image
## [4][How to upload and access a file?](https://www.reddit.com/r/rails/comments/iucars/how_to_upload_and_access_a_file/)
- url: https://www.reddit.com/r/rails/comments/iucars/how_to_upload_and_access_a_file/
---
This is a bit of a n00b question (I'm sure), but ...

I'm writing a simple Rails application that allows the user to upload a PDF file with two pages of scanned text. Once uploaded, the application should perform OCR on the PDF and display the "extracted" text. This sounds simple.

I'm using Rails 6, Active Storage, and the local disk during development.

I'm able to upload a PDF and attach it to a Case, but I don't know how to access it so I can send it to the OCR method (I'm using Docsplit with Tesseract for now).

I want to be able to say something like this in my controller:

    @c = Case.first
    @c.ocr_text = Docsplit.extract_text(@c.attachement)

And something like this in my view:

    &lt;p&gt;The text from the PDF is:&lt;/p&gt;
    &lt;%= @c.ocr_text %&gt;

It's so easy to upload a file. I'm surprised it isn't as easy (apparently) to get access to the file. Am I misunderstanding something?

AtDhVaAnNkCsE (thanks in ADVANCE)
## [5][I am using carrierwave to upload the image and parsley for validation. when I tried to edit the form it shows required value in image field though the image is uploaded.](https://www.reddit.com/r/rails/comments/iuhi7w/i_am_using_carrierwave_to_upload_the_image_and/)
- url: https://www.reddit.com/r/rails/comments/iuhi7w/i_am_using_carrierwave_to_upload_the_image_and/
---

## [6][Rails + WordPress integration](https://www.reddit.com/r/rails/comments/ityz8r/rails_wordpress_integration/)
- url: https://www.reddit.com/r/rails/comments/ityz8r/rails_wordpress_integration/
---
Hi there,

We are planning on having a WordPress site to manage content and a rails app as our platform after the user signs-in.
However, on our WordPress site we have the need of showing a list of items that is inside our platform's database. 
What's the best way of doing that? Is there a way of making our platform return the HTML to WordPress or even return a json and generate the HTML inside WordPress?
## [7][upload files in a specific folder in AWS bucket](https://www.reddit.com/r/rails/comments/iturlf/upload_files_in_a_specific_folder_in_aws_bucket/)
- url: https://www.reddit.com/r/rails/comments/iturlf/upload_files_in_a_specific_folder_in_aws_bucket/
---
im using **gem 'aws-sdk-s3'** &amp; im successfully able to upload my files on **AWS** bucket but I don't know how can I upload my files in a specific folder in my bucket ?

I'm using Active-Storage 
## [8][Where did concept of service object come from?](https://www.reddit.com/r/rails/comments/itivdn/where_did_concept_of_service_object_come_from/)
- url: https://www.reddit.com/r/rails/comments/itivdn/where_did_concept_of_service_object_come_from/
---
I ve gone through a bit of rails open source codes and am observing a lot of service objects - a class with one static function 'call'. Never really heard of this in Java Python or Node. If im correct Ruby is OOP language in the first place. I see this service object things is annoying since it looks like its just throwing the OOP concept away, and this doesnt seem to take any advantages of Ruby OOP. Some people say its because SRP but SRP doesnt strictly mean that a file or a class should only have one function that does only one thing. Well, if used correctly, It'd be the best but feels like SO concept combined with SRP is leading some to wrong way and to no good practice. Why is it popular in rails?
## [9][Trying to deploy my web app on DigitalOcean (always used Heroku before)](https://www.reddit.com/r/rails/comments/ito2ou/trying_to_deploy_my_web_app_on_digitalocean/)
- url: https://www.reddit.com/r/rails/comments/ito2ou/trying_to_deploy_my_web_app_on_digitalocean/
---
I don't understand why I have problems with the Nginx configuration...

I followed [this guide](https://gorails.com/deploy/ubuntu/18.04#nginx) (not to the letter) and everything seems to be fine, except for one thing, when I do a request to my web app I get an [`ERR_CONNECTION_REFUSED` error](https://i.imgur.com/l0P5AWu.png), and when I look at the nginx error logs I see the following lines for each request:

```
[ N 2020-09-16 02:18:46.2063 6335/T1 age/Wat/WatchdogMain.cpp:1373 ]: Starting Passenger watchdog...
[ N 2020-09-16 02:18:46.2363 6342/T1 age/Cor/CoreMain.cpp:1340 ]: Starting Passenger core...
[ N 2020-09-16 02:18:46.2365 6342/T1 age/Cor/CoreMain.cpp:256 ]: Passenger core running in multi-application mode.
[ N 2020-09-16 02:18:46.2445 6342/T1 age/Cor/CoreMain.cpp:1015 ]: Passenger core online, PID 6342
[ N 2020-09-16 02:18:48.3100 6342/T5 age/Cor/SecurityUpdateChecker.h:519 ]: Security update check: no update found (next check in 24 hours)
2020/09/16 02:19:08 [info] 6375#6375: Using 32768KiB of shared memory for nchan in /etc/nginx/nginx.conf:63
```

The strange thing is that if I use standalone passenger it seems to work correctly...
## [10][Ok... how the F do I test model concerns in isolation (Minitest/mocha)?](https://www.reddit.com/r/rails/comments/itjoev/ok_how_the_f_do_i_test_model_concerns_in/)
- url: https://www.reddit.com/r/rails/comments/itjoev/ok_how_the_f_do_i_test_model_concerns_in/
---
I’m working on a feature for work that lends itself very well to a concern. It works fantastic but I can’t figure out the best way to test it in isolation. 

It’s basically a locking feature for individual records based on domain logic. To make it simple for an example let’s say that we have a `lock!(locked_by:)` method which just calls `update!(locked_at: Time.zone.now, locked_by: locked_by)` and an `unlock!` method to nullifies those two columns in the database. Sorry for the formatting, I’m on my phone at the moment. 

There’s more to it than this but how would I go about testing this concern in isolation?
## [11][Using ActiveStorage with Aws in Rails Api](https://www.reddit.com/r/rails/comments/itfzhr/using_activestorage_with_aws_in_rails_api/)
- url: https://www.reddit.com/r/rails/comments/itfzhr/using_activestorage_with_aws_in_rails_api/
---
Hi, guys hope you are fine!

I'm learning rails and I have a task to make an API in rails that can take image params from the front-end and save it to the backend in the image model with its caption.

1. I have credentials from Aws secret id &amp; key as well.
2. yes, I know how to install Active Storage , gem "aws-sdk-s3",  gem "active storage validation.
3. I know to configure storage.yml &amp; production.rb.
4. I'll add a line has\_one\_attached :image in my Restaurant Model.

But i don't know what kind of params I'll get from the front-end developer &amp;  how can I test that params or URL in my postman from my side . &amp; how to return that image in json format to the frontend developer after saving it.

i want to upload restaurant images from users. 
## [12][Need someone to critique my Product inventory models.](https://www.reddit.com/r/rails/comments/itg4uz/need_someone_to_critique_my_product_inventory/)
- url: https://www.reddit.com/r/rails/comments/itg4uz/need_someone_to_critique_my_product_inventory/
---
So, basically I've been trying to finally use my newly learned Rails knowledge and actually build something useful. The basic idea is this: I want to be able to build a Product model that should have the following attributes: a name, color, sizes, quantity, and a couple of pictures attached. I'll be adding more attributes in the future but for now I just want feedback.


An example of what I want to keep track is this:


* I have a Product named "Jordans" that is color "Red", and I have '2 pairs' (quantity) of 'size 6; and a 4 pictures of red Jordans
* I have a Product named "Jordans" that is color "Red", and I have '1 pairs' (quantity) of 'size 7' and THE SAME picture of the red Jordans used above.
* I have a Product named "Jordans" that is color "Blue", and I have '1 pairs' (quantity) of 'size 9' and 4 pictures of blue Jordans.
* I have a Product named "Nike" that is color "Yellow", and I have '4 pairs' (quantity) of 'size 10' and 4 pictures of yellow Nike.


And so on and so fourth. Below are the Models that I came up with. This is where I need help. I'm not sure if I have the right idea or not or it can be done better.


    #This will hold every Product based on the color and size. Ex: '2 pairs' of 'red' 'Jordans' 
    class ColorBySize &lt; ActiveRecord::Base
        belongs_to :product
        
        validates :color, presence: :true
        validates :quantity, presence: :true
        validates :US_size, presence: :true
    end 


    #will be having an association with ColorBySize
    class Product &lt; ActiveRecord::Base
        has_many :color_by_sizes

        validates :name, presence: :true
    end 


At the moment, this is working well. It keeps track of the Product based on the color and size. I was wondering if there needs any refactoring done or if it can be done better? 


Another issue that I have is attaching an image. I'm aware that I can use ActiveStorage for this, but I'm not sure if I should create an association on the Product or ColorBySize model.  As the example above:


* I have a Product named "Jordans" that is color red, and I have 2 pairs (quantity) of size 6 and a 4 pictures of red Jordans
* I have a Product named "Jordans" that is color red, and I have 1 pairs (quantity) of size 7 and THE SAME picture of the red Jordans used above.


Since Red Jordan's Image already exists, I shouldn't create a new image, I just used the already existing one. The last issue that I'm facing is limiting the number of pictures I can upload. I only want 4 done. I did reaserch on this and a StackOverFlow thread mentioned that Javascript is needed to restrict the number of images that can be uploaded. Not sure if this is the only solution or if there is another? 


Thanks.
