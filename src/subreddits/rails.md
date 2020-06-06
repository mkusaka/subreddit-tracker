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
## [2][Sidekiq file access](https://www.reddit.com/r/rails/comments/gxptw0/sidekiq_file_access/)
- url: https://www.reddit.com/r/rails/comments/gxptw0/sidekiq_file_access/
---
I have an application on Heroku that has a Postgres database and sidekiq running on a hosted Redis instance.

A tool I want to rely on is able to generate JSON output when an  output file is specified but it will also print to stdout the non-JSON  output. Before deploying the app to Heroku, I was passing a temporary  file as output to the tool, opening the file and processing the JSON  contents and finally deleting it, but now on Heroku, sidekiq doesn't  have access to files.

Since I just want the JSON data, is there a way to run the tool and  just process the JSON output without using an object storage service  like S3 and having to deal with tmp files? Alternatively, is there a way  to do something like passing a memory mapped file to the output file  option of the tool and read it from memory?

Thank you
## [3][Access Rails server in guets VM without knowing IP?](https://www.reddit.com/r/rails/comments/gxee38/access_rails_server_in_guets_vm_without_knowing_ip/)
- url: https://www.reddit.com/r/rails/comments/gxee38/access_rails_server_in_guets_vm_without_knowing_ip/
---
This isn't necessarily a Rails question, but I assume some Rails developers have run into this.

I could swear that I used to be able to access my Rails server in my Ubuntu guest VM by typing the following into Chrome on my host (Windows 10):

[http://0.0.0.0:3000](http://0.0.0.0:3000)

But now, I just get "This site can't be reached". This was after upgrading from VMWare Player to VMWare Workstation Pro. I called VMWare support, and they gave me the run-around changing a bunch of settings but they ended by saying that this doesn't appear to be a VMWare issue, and I should check application/OS settings.

I should mention that when I upgraded, I had to use the "Repair" tool in the installer. The Repair seemed to be successful, and I was able to launch my VM, but it seems that the ability to access my VM server at this path has been broken.

I'm about 98% sure I was using the [0.0.0.0](https://0.0.0.0) path on my host machine. I have my VM configured to use a NAT network adapter.

From within the guest machine, I can still access my Rails server at [0.0.0.0:3000](https://0.0.0.0:3000).

Any idea what might have changed? Does anyone know how to configure my networking so that my guest responds to [0.0.0.0](https://0.0.0.0) from my host?
## [4][How to get Rails default log path any config reader present?](https://www.reddit.com/r/rails/comments/gxkqg9/how_to_get_rails_default_log_path_any_config/)
- url: https://www.reddit.com/r/rails/comments/gxkqg9/how_to_get_rails_default_log_path_any_config/
---
I am recently looking into some third party gem [https://www.zoho.com/crm/developer/docs/ruby-sdk/config.html](https://www.zoho.com/crm/developer/docs/ruby-sdk/config.html) where it wants me to link the log file if I want any via `application_log_file_path`. How do I get the Rails default log path without doing some manual work to get there?
## [5][How do I make IDs non sequential?](https://www.reddit.com/r/rails/comments/gx8h4o/how_do_i_make_ids_non_sequential/)
- url: https://www.reddit.com/r/rails/comments/gx8h4o/how_do_i_make_ids_non_sequential/
---
Hi all,  


At the moment my IDs are just sequential numbers starting at 1.  


I don't want people to know how many entries there are, so I'd like my IDs to look like a random assortment of numbers (or letters).  


I don't really care about the IDs in the database, but in particular in the URL string. Currently I'm using this method to generate the URL string:

    def to_param
      "#{id}-#{name.parameterize}"
    end

This is generating URLs like:

    /1-some-name

I would like this to be something like:

    /19987004213-some-name

Or

    /a8gY4b97a-some-name

  
I tried the obfuscate-id gem, which seems to do exactly what I want, but when I try to bundle install it, I get this error:

    Bundler could not find compatible versions for gem "rails":
      In snapshot (Gemfile.lock):
        rails (= 6.0.3.1)
    
      In Gemfile:
        rails (= 6.0.3.1)
    
        obfuscate_id was resolved to 0.2.0, which depends on
          rails (~&gt; 4.2.0)
    
    Running `bundle update` will rebuild your snapshot from scratch, using only
    the gems in your Gemfile, which may resolve the conflict.

So I assume this means it's no longer supported and I need to go with an alternative (either a gem, or my own way of doing this)  


Help appreciated.
## [6][Quick question about charts and analytics..](https://www.reddit.com/r/rails/comments/gxl3se/quick_question_about_charts_and_analytics/)
- url: https://www.reddit.com/r/rails/comments/gxl3se/quick_question_about_charts_and_analytics/
---
I want to add some static numbers, with the ability to manually update these numbers or auto-update periodically (really depends on cost). The goal of these numbers would be to create both line and pie charts to measure a market trend/growth-over-time/current status on a number of different areas, I have the numbers I need for this already. Is there an API that you would recommend for this? Bonus points if you help me recall an API for some Alpha\_\_\_\_\_ that I saw a while back but can't find for the life of me.
## [7][Is that possible using multiple versions of bootstrap in Rails ?](https://www.reddit.com/r/rails/comments/gxkvgz/is_that_possible_using_multiple_versions_of/)
- url: https://www.reddit.com/r/rails/comments/gxkvgz/is_that_possible_using_multiple_versions_of/
---
Yesterday till today I have tried to override the classes of Bootstrap 4 in my rails app
But it still read the class of the Bootstrap 3 
I downloaded Bootstrap 3 from gem and for the Bootstrap 4 I download source files on their official website
And after it I copy those files into my styles directory 
And then trying to create new directory called "new-bootstrap" and create file "_override_boostrap.scss"
And then I @import this file into my root application scss , and on _override_boostrap.scss I import the rouse file I downloaded Wich from dist/css/bootstrap.css and do created new class and @extend the class of new Bootstrapt
After I called it on my html , itu still using the Bootstrap 3 Andi think this is because the Bootstrap 3 in my app do global 
Anyone has any idea?
## [8][Using ActiveRecord Without Rails, Weird Attribute Errors](https://www.reddit.com/r/rails/comments/gxei30/using_activerecord_without_rails_weird_attribute/)
- url: https://www.reddit.com/r/rails/comments/gxei30/using_activerecord_without_rails_weird_attribute/
---
I'm building an app and using ActiveRecord outside of Rails. So far things are OK, but I'm trying to be able to create a user object like this:

```
u = User.new(name: "Bob", email: "bob@bobloblaw.com")
```

The problem is, when I do this it tells me those attributes don't exist. Here's my user.rb model:

```
class User &lt; ApplicationRecord
  include ActiveModel::SecurePassword
  has_secure_password
  belongs_to :org

  # attr_accessor :name, :phone, :email, :password, :org_id

  def visible_attributes
    attrs = attributes
    attrs.delete('password_digest')
    attrs
  end
end
```

As you can see I commented out the `attr_accessor` line. That's because with it on, I'm able to mass assign on object creation, but I get weird shit like this happening:

https://imgur.com/a/vzNSZPn

Here's the ApplicationRecord class:

```
class ApplicationRecord &lt; ActiveRecord::Base
  include ActiveModel::Model
  self.abstract_class = true

  def initialize(args)
    super
    args.each do |k, v|
      send("#{k}=", v)
    end
  end
end
```

Any ideas what I might be doing wrong here? It's pretty weird to create an object but it has no visible attributes even though the column names are a perfect match.
## [9][Stimulus course in context of Rails?](https://www.reddit.com/r/rails/comments/gx2ca6/stimulus_course_in_context_of_rails/)
- url: https://www.reddit.com/r/rails/comments/gx2ca6/stimulus_course_in_context_of_rails/
---
Are you aware of any online courses or resources that teach how to use Stimulus in the context of a Rails 6 site? I know Rails and JS in general, and I would like to pick up Stimulus, and learn best practices. Thank you!
## [10][Query to group and count by multiple columns and get a nice representation of the data?](https://www.reddit.com/r/rails/comments/gx8op3/query_to_group_and_count_by_multiple_columns_and/)
- url: https://www.reddit.com/r/rails/comments/gx8op3/query_to_group_and_count_by_multiple_columns_and/
---
I have a model with the `created_at` column and a `state` column (which may have different values, like `A`, nil, etc.), I want to get the amount of records I have in that model, first categorized by `created_at` and then by the `state` column, something like the following...

```
{
  [Wed, 01 Jan 2020]=&gt; { "A" =&gt; 662, nil =&gt; 1091 },
  [Thu, 02 Jan 2020]=&gt;{ "A" =&gt; 2895, nil =&gt; 5058 },                                                                                                                                                                                    
  ...                                                                                                                                                                          
```

The closest I've gotten was this:

```
{
  [Wed, 01 Jan 2020, "A"]=&gt;662,
  [Wed, 01 Jan 2020, nil]=&gt;1091,
  [Thu, 02 Jan 2020, "A"]=&gt;2895,
  [Thu, 02 Jan 2020, nil]=&gt;5058,                                                                                                                                                                                                                                                                                                                                                                   
  ...
```

With this query: `MyModel.order('DATE(created_at) ASC').group('DATE(created_at)', 'state').count`

How can I get what I want in one query? I'd like it to be as efficient as possible.
## [11][Awesome-Ruby: A Curated Collection of 851 Ruby Links and Resources](https://www.reddit.com/r/rails/comments/gwv338/awesomeruby_a_curated_collection_of_851_ruby/)
- url: https://www.reddit.com/r/rails/comments/gwv338/awesomeruby_a_curated_collection_of_851_ruby/
---
Check this on github. It really helps. Awesome-Ruby: https://github.com/markets/awesome-ruby#readme
