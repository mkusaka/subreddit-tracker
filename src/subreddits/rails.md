# rails
## [1][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/ezrfed/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/ezrfed/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [2][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/f2r9mk/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/f2r9mk/personal_projects_show_off_your_own_project_andor/
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
## [3][Storing attachment file type when using active storage](https://www.reddit.com/r/rails/comments/f4plor/storing_attachment_file_type_when_using_active/)
- url: https://www.reddit.com/r/rails/comments/f4plor/storing_attachment_file_type_when_using_active/
---
Hey,

I was wondering if there’s a way to store, in the database, the attachment file type when using active storage. Basically a way to differentiate if a file is an Image Video etc.

I was thinking of making a monkey patch Active Storage Blob Model that adds this functionality based on MIME types but I’m assuming there must be a better way around it.

Thanks!
## [4][A beginners guide on how to implement Pagination &amp; Search via JSON into a Select2 drop-down on Rails 6](https://www.reddit.com/r/rails/comments/f4h609/a_beginners_guide_on_how_to_implement_pagination/)
- url: https://www.reddit.com/r/rails/comments/f4h609/a_beginners_guide_on_how_to_implement_pagination/
---
Caveat:  I still consider myself a newbie, and I hope there's no glaring errors in the code below.  I'm open to any suggestions if something looks off, but for my case, everything worked fine.  I'm writing this up in hopes that it'll help somebody else out down the line that may be searching for a similar solution to speed up their apps.

---

I'm writing this mini-guide for a few reasons.

1. the Select2 documentation on Ajax doesn't apply very well to Rails conventions
2. the [Select2-Rails gem](https://github.com/argerim/select2-rails) is severely out-of-date
3. all Google &amp; StackOverflow searches didn't bring up anything that applied to a modern implementation of Select2 into Rails

Versions of all dependencies are up to date as of Feb 15, 2020 on Rails 6 using Webpacker.  

Gems installed are [Pagy for Pagination](https://github.com/ddnexus/pagy), and [Ransack for Searching](https://github.com/activerecord-hackery/ransack).  [Select2](https://github.com/select2/select2) was installed [via Yarn](https://www.npmjs.com/package/select2).  This also assumes jQuery is installed, of course!

My use case is that I want people to be able to create a new Address record, and within the form, to add an associated record ID in a Select2 dropdown via its polymorphic `addressable_id` attribute.  In this case, we're focusing on the Address being associated with a `supplier`.

---
    # schema.rb
    create_table "addresses", force: :cascade do |t|
      t.string "address_line_1"
      # etc, etc.  Cutting out a bunch of superfluous columns
      t.string "addressable_type"
      t.bigint "addressable_id"
      t.datetime "created_at", precision: 6, null: false
      t.datetime "updated_at", precision: 6, null: false
    end
    
    create_table "suppliers", force: :cascade do |t|
      t.string "name"
      t.datetime "created_at", precision: 6, null: false
      t.datetime "updated_at", precision: 6, null: false
    end
---
    # address.rb
    class Address &lt; ApplicationRecord
      belongs_to :addressable, polymorphic: true, optional: true
    end
---
    # supplier.rb
    class Supplier &lt; ApplicationRecord
      has_many :addresses, as: :addressable
    end

Your use case may vary wildly, but this happens to be the first one I needed dynamic dropdowns for.  First up, the view.  In the New Address form, I've added a standard Select form field with a blank `options_for_select` tag as `:addressable_id`:

    # /app/views/addresses/_form.html.erb

    &lt;%= f.select :addressable_id, "", {}, {class: "addressSupplierSearch"} %&gt;

Next up, we need to get the relevant JSON up and running by creating a new jbuilder file (this was the least documented part of the whole process)

    # /app/views/addresses/addressable_dropdowns.json.jbuilder

    # 'json.page' &amp; 'json.count_filtered' relays the correct information to the select2 API
    # '@pagy.page' and '@pagy.count' are grabbed via the Pagy API
    json.page @pagy.page
    json.count_filtered @pagy.count
    
    # '@suppliers' is populated from the 'addresses_controller'
    json.items do
      json.array!(@suppliers) do |s|
        json.name s.name
        json.id s.id
      end
    end

We are using the most basic objects here with just the `name` and `id` being grabbed.  You can go way more in depth though, and even go as far as adding custom html with images, etc.  A great example of that is found in the select2 documentation via [the code at the bottom of this linked example.](https://select2.org/data-sources/ajax#additional-examples)

This is the output of some [Faker](https://github.com/faker-ruby/faker) data from the above JSON:

    {
      "page": 1,
      "count_filtered": 204,
      "items": [
        {
          "name": "Alternative Dispute Resolution",
          "id": 199
        },
        {
          "name": "Government Relations",
          "id": 160
        },
        {
          "name": "Bosco-Tillman",
          "id": 157
        },
        {
          "name": "Botsford Group",
          "id": 95
        },
        {
          "name": "Glover-Carter",
          "id": 80
        },
        {
          "name": "Electrical / Electronic Manufacturing",
          "id": 204
        },
        {
          "name": "Beier LLC",
          "id": 183
        },
        {
          "name": "Accounting",
          "id": 139
        },
        {
          "name": "Hirthe, Renner and Breitenberg",
          "id": 45
        },
        {
          "name": "Apparel &amp; Fashion",
          "id": 191
        }
      ]
    }
Of note here is that it defaults to Page 1, and that `count_filtered` is required for later in the Javascript so that the select2 API can do some calculations for how many pages are required.  In this case, I have 204 records, so with 10 records per page, it can extrapolate 21 pages for us.

Don't forget to add this new file to your routes:

    # routes.rb

    get :addressable_dropdowns, controller: :addresses
and in your controller:

    # addresses_controller.rb

    def addressable_dropdowns
      @pagy, @suppliers = pagy(Supplier.ransack(name_cont: params[:term]).result(distinct: true), items: 10)
    end

This is where the Pagination &amp; Search kick in.  We are using both the Pagy &amp; Ransack gems, and they seem to work flawlessly together without any extra coercion. `pagy` indicates that each pagination batch should be 10 items long, and Ransack allows the select2 dropdown input to be searchable via the Supplier's `name`.  Note that I originally set up Pagy via the [GoRails video found here](https://gorails.com/episodes/pagination-with-pagy-gem), and the basic setup for Pagy should be done as a prerequisite.  Ransack should work right out of the box.

The final step is to wire up the dropdown through Javascript:

    # application.js

    $(document).on("turbolinks:load", function() {
    
      $('.addressSupplierSearch').select2({
        placeholder: "Select one...",
        ajax: {
          url: "/addressable_dropdowns.json",
          dataType: "json",
          type: "GET",
          processResults: function (data, params) {
            params.page = data.page || 1;
            return {
              results: $.map(data.items, function (item) {
                return {
                  text: item.name,
                  id: item.id
                };
              }),
              pagination: {
                more: (params.page * 10) &lt; data.count_filtered
              }
            };
          },
          cache: true
        }
      });
    
    });

`processResults` is the biggest change here compared to all the tutorials I found that relied on a much older select2 API.

A few notes of how this Javascript is working:

* `params.page = data.page || 1` relays the information from the Pagy API via the `json.page @pagy.page` line in the JSON file.
* on the `results: $.map(data.items, function (item) {` line, `data.items` is the important bit here, as I noticed a lot of older guides were simply calling `data` instead.
* The `text` and `id` columns are specified in the select2 API.  Reminder too that `item.name` can be replaced by just about anything that you want.
* Read up more on how select2 pagination works [at this link](https://select2.org/data-sources/ajax#pagination), and note that the `data.count_filtered` line is grabbed from the Pagy API via `@pagy.count` in the JSON file.

---

One curiousity that I'd like to figure out and got stumped on;  For all of my normal non-ajax Select2 dropdowns within a form, I would have been able to automatically populate it in the Edit view by calling the `selected` option within the form object:

    &lt;%= f.select :addressable_id, "", {selected: f.object.addressable_id}, {class: "form-control addressSupplierSearch"} %&gt;

But since the dropdown options in this case are generated dynamically, it no longer works and the field remains blank on Edit, and I don't know how to automatically populate it if there's already a value there.  If anybody has any ideas on how to fix this, I'd appreciate it.
## [5][Tips and suggestions for managing multiple crawlers?](https://www.reddit.com/r/rails/comments/f4avxv/tips_and_suggestions_for_managing_multiple/)
- url: https://www.reddit.com/r/rails/comments/f4avxv/tips_and_suggestions_for_managing_multiple/
---
I’m working on a side project that aggregates training courses from various instructors. Some of those instructors publish Google calendars, others list their courses on Shopify, some just list them on their website... you get the idea. 

So, for now, I’ve been writing an individual rake task for each one... which is fine for a few but now starting to feel like it’s getting crazy. 

So, my questions:

1. Has anyone had experience (or seen example projects) of this sort of project? 
2. On NomadList, I notice that there’s a “crowd edit” feature where users can point out that somethings wrong... and somehow, those changes are either reviewed by the admin or merged in with other changes. Any idea how that feature was implemented?
3. Has anyone successfully brought on a developer via Upwork or similar to write these sort of crawlers?

I’m getting ~50 new signups a day now and, as a one man team, I’m really starting to think about how I can free up the majority of my time to ramp up user acquisition channels...
## [6][How to find new entries in two XML files](https://www.reddit.com/r/rails/comments/f4cq8a/how_to_find_new_entries_in_two_xml_files/)
- url: https://www.reddit.com/r/rails/comments/f4cq8a/how_to_find_new_entries_in_two_xml_files/
---
I need to find new entries in a XML sitemap compared to its last version in a rails app. I wonder how to tackle this challenge. I could parse the XML sitemap and check each line to see if it was already in the previous version but the process will be quite slow and memory-intensive.  


Any suggestion about how I could do that? Thanks.
## [7][How do you include the lib folder in your spec ?](https://www.reddit.com/r/rails/comments/f44p47/how_do_you_include_the_lib_folder_in_your_spec/)
- url: https://www.reddit.com/r/rails/comments/f44p47/how_do_you_include_the_lib_folder_in_your_spec/
---
I have some custom PORO in the lib/ folder. I am using Rails 5 and Rspec. I would like to include files from lib in the spec. I found answers [1](https://stackoverflow.com/questions/4073856/rails-3-autoload), [2](https://stackoverflow.com/questions/16954989/how-to-include-lib-directory-in-rspec-tests) which is too old. So I am seeking some advices majority of rails community follows.
## [8][First project finished, but badly in need of criticism!!](https://www.reddit.com/r/rails/comments/f3z6zt/first_project_finished_but_badly_in_need_of/)
- url: https://www.reddit.com/r/rails/comments/f3z6zt/first_project_finished_but_badly_in_need_of/
---
##Hey there you wonderfully helpful subreddit, you!

Thanks to so much help from you guys over the last month+ I've finished and deployed my first webapp. [You can check it out here](https://www.chucksef.com) but it really won't be terribly interesting since none of you are admins (and hopefully there's no way for you to sneak your way in!)

[Here's a link to the repo on github](https://github.com/Chucksef/photo_app) if you're interested in looking at the code.

If you didn't read the title, I **very badly want criticism**! I'm very new to rails (just started learning Jan 3rd) so I'm sure there are an overwhelming number of things that I've done 100% wrong, and I'd like to know about them!

#Website concept: 
Squarespace-y platform I built myself to allow me to dynamically build my own portfolio without having to go back into my development environment and re-code pages as I finish more projects to show off!

Right now, however, I've filled it with some generic photo-studio stuff.

However, here are a few categories where I'm NOT looking for criticism:

* Webpage is very generic looking bootstrap bullshit - I know. That's not the point of this project!
* I didn't use ANY testing tools while building this. Oh man do I really regret glossing over this stuff. Now that I'm finished every time I make a change I cringe at not knowing if something else broke! I WILL be learning this in the near future.
* Website copy is all bad - I'm not trying to become a copywriter.

Anything else is fair game! So come on and help me become better by destroying my month+ of studying and hard work!
## [9][Highcharts in Ruby syntax for Rails app / Also, trying to get two series of data to display on a single highchart and it's not taking](https://www.reddit.com/r/rails/comments/f3ybw0/highcharts_in_ruby_syntax_for_rails_app_also/)
- url: https://www.reddit.com/r/rails/comments/f3ybw0/highcharts_in_ruby_syntax_for_rails_app_also/
---
Does anyone know of any good docs for highcharts in rails?  Seems sparse if much at all.  I don't have a JS background yet to translate much from JS to Ruby, if that is useful (I notice lots of demos in JS fiddle).

Additionally, Im trying to get two series of data to load on a highchart in rails.  One column type and one scatter type, each with nested arrays like `[[x1, y1], [x2, y2], [x3, y3],  ...]` and so on

I'm making accessibly charts, and I've tried `chart.series( { &lt; chart 1&gt; }, { &lt;chart 2&gt; })`, I've tried making separate `chart.series`, I've tried making two separate Accessible charts.  One chart is working but can't get to seem to get the other one going, trying to add the second plot in various was as mentioned above.  Page won't load properly when I try to get a second one in
## [10][How nested_attributes works?](https://www.reddit.com/r/rails/comments/f3ui2c/how_nested_attributes_works/)
- url: https://www.reddit.com/r/rails/comments/f3ui2c/how_nested_attributes_works/
---
Hi guys, I was wondering how Rails create &amp; update nested attributes.

I had a problem where I was creating 2 associated models at the same time. For that, I was using the rails callbacks :after\_create/update, and while I was doing that, I knowledge that Rails pushes the code to the Stack (execute asynchronous), hence, inside the method called by the callback, we only have access to database's data. Also, apparently, we can't pass arguments (I tried and didn't work) to those triggers.

For this problem, the only way I was able to pass data from one model to another was creating database attributes to the first model and passing them to the second model. I did try to pass attr\_acessor first, but that didn't work for the reason I described above.

Besides using nested\_attributes to this problem, there was something that I'm missing? I guess that Rails has more privileges to pass arguments to these triggers or maybe the arguments should be feasible of serialization - similar to how works Sidekiq's Workers.

I hope this example can clarify my explanation

    # inside First::Model.rb
    
    attr_accessor :second_model_attr_4
    
    after_create :create_solicitation
    
    def create_solicitation
        puts second_model_attr_4 # returns nothing
     
        Second::Model.create!(
          first_model_id: self.id,
          second_model_attr_1: self.second_model_attr_1 
          second_model_attr_2: self.second_model_attr_2
          second_model_attr_3: Second::Model::Status.find(
            self.second_model_attr_3
         ),
          # this attr_accessor won't be seem by the function  
          second_model_attr_4: self.second_model_attr_4 
        )
    end
    
    # P.s second_model_attr_1 to 3 were added as database attributes to the First::Model (did I migration and add these methods to First::Model)

P.s It's callbacks, not triggers - sorry.
## [11][bundle exec foreman start fails](https://www.reddit.com/r/rails/comments/f436wc/bundle_exec_foreman_start_fails/)
- url: https://www.reddit.com/r/rails/comments/f436wc/bundle_exec_foreman_start_fails/
---
&gt; \`parse\_numeric': Expected numeric value for '--port'; got "${PORT-3000}" (Thor::MalformattedArgumentError)  


All I am doing is creating the environment for an open source project.  
 [https://github.com/huginn/huginn](https://github.com/huginn/huginn)   


Gemfile:  
 [https://github.com/huginn/huginn/blob/master/Gemfile](https://github.com/huginn/huginn/blob/master/Gemfile)   


Gemfile.lock  
 [https://github.com/huginn/huginn/blob/master/Gemfile.lock](https://github.com/huginn/huginn/blob/master/Gemfile.lock)   


I don't really know why I am getting this.  


From online sources they told me to try downgrading thor from 0.19.4 to 0.19.1, didn't work, another told me to try doing 0.20.0, but didn't work.  


I'm running this on Windows 10.
## [12][Should ActiveRecord be used for already existing files?](https://www.reddit.com/r/rails/comments/f3yvdp/should_activerecord_be_used_for_already_existing/)
- url: https://www.reddit.com/r/rails/comments/f3yvdp/should_activerecord_be_used_for_already_existing/
---
Hey,

I was reading a lot about ActiveStorage in the past days and I'm still not sure if it's the right solution for my problem or if it's more aimed at simple user file uploads.

I have an Image model which has an "external_url" field which contains a random direct URL to an image scraped from the internet. It's linked to another model called "Show". I don't want to store that image blob as soon as I insert a new Image as it would be too many (in my case).

What I want is that as soon as someone requests that image by visiting a show page like: http://localhost:3000/show/example-show it'll download it from the "external_url", then upload it to S3 and create a variant of it. Just like ActiveStorage usually does, just with the difference that it has to download it first and then do it's usual thing.

Is that an unusual approach and should I just use an external image server like https://github.com/willnorris/imageproxy instead?

Any pointers into the right direction or usual approaches are appreciated!

Thanks!
