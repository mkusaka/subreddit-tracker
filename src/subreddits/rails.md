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
## [2][Rails 6.1's ActiveModel Errors Revamp](https://www.reddit.com/r/rails/comments/gy80or/rails_61s_activemodel_errors_revamp/)
- url: https://www.reddit.com/r/rails/comments/gy80or/rails_61s_activemodel_errors_revamp/
---
[https://code.lulalala.com/2020/0531-1013.html](https://code.lulalala.com/2020/0531-1013.html)

As Rails developers, we are all used to the \`book.errors(:title)\` interface. This has remained relatively stable up until now, but is soon going to change.

I'd like to share the new model errors changes, before Rails 6.1rc1 gets released. The article contains a list of deprecation and recommended replacements offered in the new implementation. I hope to get some feedback, and see if we need to improve the upgrade guide a bit, to make the migration process less painful.

And if you have any suggestion on the actual code changes it self, please also let me know. Thanks you!
## [3][How to identify between new and edit actions in javascript with webpack](https://www.reddit.com/r/rails/comments/gy6o5o/how_to_identify_between_new_and_edit_actions_in/)
- url: https://www.reddit.com/r/rails/comments/gy6o5o/how_to_identify_between_new_and_edit_actions_in/
---
I am hiding some elements in the new action, but I don't want to do on the edit action. Is there a way to do this, without checking whether the URL has edit in it.  


I am using webpack, so there is no way to pass instance variables inside js
## [4][How to configure a private app in Shopify](https://www.reddit.com/r/rails/comments/gxxiu1/how_to_configure_a_private_app_in_shopify/)
- url: https://www.reddit.com/r/rails/comments/gxxiu1/how_to_configure_a_private_app_in_shopify/
---
Hi! I'm new to this community. I have a quick question about configuring a private app in Shopify. My app is in Ruby on Rails and I am trying to install it on to Shopify. I created a private app on Shopify through the partners dashboard. When I enter the domain (example.shopify.com) I created in Shopify in the login page, I get this error:

Oauth error invalid\_request: The Shopify API application does not support oauth

The url i'm redirected to is this: [https://example.myshopify.com/admin/oauth/authorize?client\_id=c54817f9011be80129703e1987a4a37e&amp;redirect\_uri=http%3A%2F%2Flocalhost%3A3000%2Fauth%2Fshopify%2Fcallback&amp;response\_type=code&amp;scope=read\_products&amp;state=2c0c64c657dd9a212e088f7018c6e8fe4ba87c39e4542ec1](https://example.myshopify.com/admin/oauth/authorize?client_id=c54817f9011be80129703e1987a4a37e&amp;redirect_uri=http%3A%2F%2Flocalhost%3A3000%2Fauth%2Fshopify%2Fcallback&amp;response_type=code&amp;scope=read_products&amp;state=2c0c64c657dd9a212e088f7018c6e8fe4ba87c39e4542ec1)

My code for configuration is this:

\`\`\`

ShopifyApp.configure do |config|  
  config.application\_name = "APP NAME"  
  config.api\_key = ENV\['SHOPIFY\_API\_KEY'\]  
  config.secret = ENV\['SHOPIFY\_API\_SECRET'\]  
  config.old\_secret = ""  
  config.scope = ENV\['SHOPIFY\_PERMISSIONS'\] # Consult this page for more scope options:  
 \# https://help.shopify.com/en/api/getting-started/authentication/oauth/scopes  
  config.embedded\_app = true  
  config.after\_authenticate\_job = false  
  config.api\_version = "2020-01"  
  config.shop\_session\_repository = 'Shop'  
  config.webhooks = \[  
{topic: 'orders/create', address: "#{ENV\['APP\_HOST'\]}/webhooks/orders/order\_create", format: 'json'},  
{topic: 'fulfillments/create', address: "#{ENV\['APP\_HOST'\]}/webhooks/orders/order\_fulfilled", format: 'json'}  
  \]  
end

\`\`\`

Am i using the wrong url format? I'm not sure what to set 'APP\_HOST' to be in env file either.

Would be grateful for any answers. Thanks!
## [5][specifying ruby version in gemfile](https://www.reddit.com/r/rails/comments/gy07ne/specifying_ruby_version_in_gemfile/)
- url: https://www.reddit.com/r/rails/comments/gy07ne/specifying_ruby_version_in_gemfile/
---
howdy, working through the Hartl book and I noticed in his gem file he says to copy, doesn't include Ruby whereas it is in my default new rails app gem file. Is the best practice the leave out Ruby in your gem file? how does it work with including Ruby in your gem file vs. not including? Will the rails app look for Ruby/Ruby version in a different file?
## [6][Rails + ROM-rb](https://www.reddit.com/r/rails/comments/gxyb0d/rails_romrb/)
- url: https://www.reddit.com/r/rails/comments/gxyb0d/rails_romrb/
---
Hey,

Is there anybody here that uses ROM + Rails in production?

Since ROM uses Sequel, its a safe and mature usage? (replace ActiveRecord for ROM)
## [7][Sidekiq file access](https://www.reddit.com/r/rails/comments/gxptw0/sidekiq_file_access/)
- url: https://www.reddit.com/r/rails/comments/gxptw0/sidekiq_file_access/
---
I have an application on Heroku that has a Postgres database and sidekiq running on a hosted Redis instance.

A tool I want to rely on is able to generate JSON output when an  output file is specified but it will also print to stdout the non-JSON  output. Before deploying the app to Heroku, I was passing a temporary  file as output to the tool, opening the file and processing the JSON  contents and finally deleting it, but now on Heroku, sidekiq doesn't  have access to files.

Since I just want the JSON data, is there a way to run the tool and  just process the JSON output without using an object storage service  like S3 and having to deal with tmp files? Alternatively, is there a way  to do something like passing a memory mapped file to the output file  option of the tool and read it from memory?

Thank you
## [8][How to get Rails default log path any config reader present?](https://www.reddit.com/r/rails/comments/gxkqg9/how_to_get_rails_default_log_path_any_config/)
- url: https://www.reddit.com/r/rails/comments/gxkqg9/how_to_get_rails_default_log_path_any_config/
---
I am recently looking into some third party gem [https://www.zoho.com/crm/developer/docs/ruby-sdk/config.html](https://www.zoho.com/crm/developer/docs/ruby-sdk/config.html) where it wants me to link the log file if I want any via `application_log_file_path`. How do I get the Rails default log path without doing some manual work to get there?
## [9][Access Rails server in guets VM without knowing IP?](https://www.reddit.com/r/rails/comments/gxee38/access_rails_server_in_guets_vm_without_knowing_ip/)
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
## [10][How do I make IDs non sequential?](https://www.reddit.com/r/rails/comments/gx8h4o/how_do_i_make_ids_non_sequential/)
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
## [11][Using ActiveRecord Without Rails, Weird Attribute Errors](https://www.reddit.com/r/rails/comments/gxei30/using_activerecord_without_rails_weird_attribute/)
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
