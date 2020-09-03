# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/igyvm1/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/igyvm1/personal_projects_show_off_your_own_project_andor/
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
## [2][How to add hours to datetime?](https://www.reddit.com/r/rails/comments/ilsvo3/how_to_add_hours_to_datetime/)
- url: https://www.reddit.com/r/rails/comments/ilsvo3/how_to_add_hours_to_datetime/
---
I want to add (or subtract) hours like these  '+0330' to datetime. Let's say if I have '2020-09-03 15:30:20'

for '+0330' I want it to become '2020-09-03 19:00:20'. 

for '-0330' I want it to become '2020-09-03 12:00:20'. 

I came up with this simple method, but it seems to be prone to errors, any suggestions? :)

    def timezoneize(timedate, timezone_digits)
    digits = timezone_digits.gsub('0', '')
     if digits[0] = '+'
       timedate + digits.to_i.hours
     else
       timedate - digits.to_i.hours
     end
 end
## [3][Connection pool size](https://www.reddit.com/r/rails/comments/ilptrw/connection_pool_size/)
- url: https://www.reddit.com/r/rails/comments/ilptrw/connection_pool_size/
---
Hi everyone! I’m working on an application that takes payments for user registrations. The transaction is completed through stripe checkout and I have a webhook listening for the checkout.session.completed event. Now this seems to work fine in most cases. However, on local host, sometimes I get a BusyException error because I’m using MySQL a which doesn’t support concurrent access. I have a PostgreSQL database setup in prod which I believe does support concurrent access. Currently the pool size of my application is at its default of 5. Should be be increasing this number as I’m expecting many users to use the site at the same time? Are there any drawbacks of doing so? Since the payments are processed on our site using webhooks (a new thread is created for each transaction), what should I set the pool count as?

I appreciate your help as I’m a newbie.
## [4][Incremental, by year, order number](https://www.reddit.com/r/rails/comments/iltbm1/incremental_by_year_order_number/)
- url: https://www.reddit.com/r/rails/comments/iltbm1/incremental_by_year_order_number/
---
Once again I'm writing the code to generate an order number and every time I do, I have doubts about my implementation hence my question here: how would you create an order number, composed by the year and an incremental number? The incremental number should reset itself every year.

The order numbers should be: 2020-0001, 2020-0002, 2020-0003, .... and so on.  
When 2021 starts, the sequence should start again: 2021-0001, 2021-0002, ...

How would you do that, also taking care of possible race conditions and performance (it should be really fast) ?
## [5][How I built a "URL to image" microsite over the weekend with Rails](https://www.reddit.com/r/rails/comments/il8cpz/how_i_built_a_url_to_image_microsite_over_the/)
- url: https://www.reddit.com/r/rails/comments/il8cpz/how_i_built_a_url_to_image_microsite_over_the/
---
I've grown really tired of manually creating social images for every single blog post. They take way too long to create and online tools always end up looking too generic. How many stock photos can I scroll through before they all start to look the same?

So I built [Mugshot Bot](https://www.mugshotbot.com). An automated, zero effort social image generator. You pass it a URL and it generates a perfectly sized, unique, beautiful social image.

Here's what they look like! The color and background pattern are randomized from a hand-tuned selection. The title and subtitle come directly from the HTML.

## Overall approach

My goal is to design in HTML and CSS and then convert it to a PNG. This worked pretty well with some `wkhtmlto*` magic but there were a few hoops I had to jump through. Here's what I did.

### Fetch the content

All of the content comes directly from the URL's HTML. So the first step is to fetch the website and parse the DOM. I'm using `HTTParty` and `Nokogiri` and then looking for specific markup.

```ruby
body = HTTParty.get(@url).body
html = Nokogiri.parse(body)
title = html.at_css("meta[property='og:title']")
  .attr("content")
description = html.at_css("meta[property='og:description']")
  .attr("content")
```

### Render and style the HTML

Now that we have the copy we can drop it into some HTML. In Rails we can render an arbitrary view and pass in some variables via `ApplicationController#render`.

```ruby
mugshot = Mugshot.new(title: title, description: description)
rendered_html = ApplicationController.render(
  "mugshots/show",
  assigns: { title: title, description: description },
  formats: [:html],
)
```

The rendered HTML uses the default layout so we have all of the CSS and fonts normally added in `&lt;head&gt;`.

## Convert to an image

Where the magic happens: `wkhtmlto*`. Or, as it is usually known, `wkhtmltopdf`. This library is bundled with a lesser known tool `wkhtmltoimage` that does exactly what we need.

If you have the library installed you can call directly into it with `Open3`. This works a bit better than backticks because you can handle stderr.

```ruby
result, error = Open3.capture3(
  "wkhtmltoimage jpeg - -",
  stdin_data: rendered_html
)
```

The two dashes (`- -`) at the end of the command tell the tool to render _from_ stdin and render _to_ stdout. `Open3` will write stdout to `result` and `stderr` to `error`.

## Render from the controller

`result` is the actual image, as data. We can render this directly from the controller. Ideally, this would be uploaded to S3 and/or put behind a CDN.

```ruby
def show
  # ...
  send_data(result, type: "image/jpeg", disposition: "inline")
end
```

## What a weekend!

Thanks for reading, I hope you enjoyed how I built a little side project over the weekend.

If you give [Mugshot Bot](https://www.mugshotbot.com) a try please let me know what you think in the comments! I'm open to feature requests, too.
## [6][How do async concepts apply to rails(4.2.1)?](https://www.reddit.com/r/rails/comments/il800i/how_do_async_concepts_apply_to_rails421/)
- url: https://www.reddit.com/r/rails/comments/il800i/how_do_async_concepts_apply_to_rails421/
---
Can someone help me with how async will look in rails code?   
Say for example for a controller action to return I need to make five external API calls and I can do something like async/await in C#. I haven't come across any such code like that.   


Can someone please help me with how async/await concepts like that in C#, JS or other languages apply to ruby and also to rails.
## [7][How can you distribute input elements outside a form using UJS?](https://www.reddit.com/r/rails/comments/ilglc6/how_can_you_distribute_input_elements_outside_a/)
- url: https://www.reddit.com/r/rails/comments/ilglc6/how_can_you_distribute_input_elements_outside_a/
---
I’m using Rails 5 and Bootstrap 4.  I have a table that lists employees and all of their relevant information.  The last column of the table has ‘modify’ buttons that turn their respective columns into forms for modifying that employee’s data.

My current solution uses UJS to remove the td’s of the clicked row, and replace them with a single table-spanning td that contains the edit form in it.  Functionally, it works perfectly.  Stylistically, I would prefer if the table’s borders remained.  I also can’t reliably keep the input elements aligned with the surrounding columns.

After research, the only solution I've found is to keep the form external to the table and distribute the input elements throughout the table row, each one pointing back to the form with a form attribute.  However, when I tried to actually implement this in the js.erb file, putting form_helper elements into javascript elements, it felt clunky and inarticulate, a feeling I usually get when I’m doing something in a non-railsy way.

This is largely a learning project, so I'm trying to work through these smaller obstacles without gems. Does anyone have any suggestions?
## [8][Multi-Tenant - Sign up as company/organization under existing Tenant-Account](https://www.reddit.com/r/rails/comments/ilfebg/multitenant_sign_up_as_companyorganization_under/)
- url: https://www.reddit.com/r/rails/comments/ilfebg/multitenant_sign_up_as_companyorganization_under/
---
Hello Rails community,

I am trying to customize my app but have been stuck for the last 4 days working on the data model. So I really appreciate your hints / input regarding my problem:

**About the app:**

The app should be a multi-tenant app to run multiple classified ads sites.

Each site has similar UI (slightly customizable), but shows different data.

User can sign up to create a new classified ads site (= new Account) under [myapp.com/registration](http://myapp.com/registration)

**Models:**

*Account*

* Each new site is represented by an Account with its own URL path ( e.g. [myapp.com/:account\_id](http://myapp.com/:account_id))
* Current Tenant / Account is set by [account.id](http://account.id/) in URL for the start, later a custom domain should point to each Account
* scoping of resources that belong to an Account happens by an \`account\_id \` field on a resource

*User - Site/Account Owner*

* User who is the admin / owner of the account after registration (Account belongs\_to :owner, class: "User")

*Category - Site can have one or many Categories (scoped via: account\_id)*

* id
* account\_id
* name

**-- So far so good, but I am struggling with the following requirement --**

Companies/Non-Profits/Organizations should be able to create Posts on an existing site (Account).

Therefore they need to sign up as a Company/Organization profile under an existing site.

Example: [myapp.com/:account\_id/organizations/new](http://myapp.com/:account_id/organizations/new)

*Organization/Company*

* id
* name
* address
* account\_id
* owner\_id (class: User)

Organization/Company belongs to User who becomes the owner after registration. He can add / invite other team mates to join this organization.

*Post*

* id
* account\_id
* organization\_id
* category\_id
* title

Posts can be created by Organizations/Companies on a classified ads site (Account).   
 Each Site display only Posts that belong to it.

**My questions are:**

* How would you implement the Organization/Company sign up under an existing Site(Account)?
* Are there any pitfalls / errors in my data model?
* Should a Company/Organization be scoped to an Account or be "outside" like the User model?

English is not my mother tongue, so please apologize grammar and other mistakes ;)
## [9][Question about hiring someone for a Rails project](https://www.reddit.com/r/rails/comments/ildiql/question_about_hiring_someone_for_a_rails_project/)
- url: https://www.reddit.com/r/rails/comments/ildiql/question_about_hiring_someone_for_a_rails_project/
---
Good afternoon all, I hope this is the right place to post this. 

Looking to speak with someone about getting a quote for a project. I need a web app created using RoR and Active Admin. I've never done this before so I have no idea exactly what criteria I need to have laid out to enable the potential developer to provide me with an accurate quote. I also have no idea what costs typically run for this (I assume there is a somewhat standard hourly rate maybe?).

Basically, I don't know what I don't know and could use some guidance. Thanks all!
## [10][Email from contact form in Rails 6](https://www.reddit.com/r/rails/comments/iky3nn/email_from_contact_form_in_rails_6/)
- url: https://www.reddit.com/r/rails/comments/iky3nn/email_from_contact_form_in_rails_6/
---
Hello. I'm adding a contact form in my web development portfolio. I would like people to be able to send me a message to my email directly from this form. I'm trying to find a way to do this with Rails 6 and I'm not finding anything new. While Googlilng, I notice that the JS has changed slightly in Rails. It's moved from assets to it's own folder because of webpacker. I'm such a newbie so I don't know much about why that's important. I'm pretty sure I have to add the "mail\_form", "simple\_form", "jquery-rails", and I guess "bootstrap\_sass" gems. Now I'm just wondering what to do next.

Older ways I see are basically adding a controller for the contact form. However, I my form will just be at the bottom of my homepage. Most of my portfolio will just be on this page, so when people are done looking at it, I want them to be able to just scroll to the bottom and see the form there, then message me if that's how they want to contact me quickly.  If someone could point me in the right direction of doing this for Rails 6, I'd appreciate it!
## [11][Is it a good idea learning Rails by the book of 2017 year?](https://www.reddit.com/r/rails/comments/ikph1s/is_it_a_good_idea_learning_rails_by_the_book_of/)
- url: https://www.reddit.com/r/rails/comments/ikph1s/is_it_a_good_idea_learning_rails_by_the_book_of/
---
I am reading Ruby on Rails Tutorial book by Michael Hartl (2017) now. I am feeling some troubles with Ruby version and Gems version. 

So, is it a good idea or I need to find some newer resource? What resource you can recommend?

P.S. I went from Python &amp; Django. I got new project and part of it wrote on Ruby. Stack: Django, Rails, React - looks like Frankenstein. :)
