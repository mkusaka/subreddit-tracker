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
## [3][How can I use &amp; instead of try! in this case?](https://www.reddit.com/r/rails/comments/f7ngfi/how_can_i_use_instead_of_try_in_this_case/)
- url: https://www.reddit.com/r/rails/comments/f7ngfi/how_can_i_use_instead_of_try_in_this_case/
---
In a the view I show the first 120 letters of the description of some posts, but I have to safe navigate these since `post.description` can be `nil`, for now I'm using `post.description.try!(:[], 0..120)`.

Is there a way to do the same with `&amp;` syntax?
## [4][SQL queries in rails application](https://www.reddit.com/r/rails/comments/f7jg3l/sql_queries_in_rails_application/)
- url: https://www.reddit.com/r/rails/comments/f7jg3l/sql_queries_in_rails_application/
---
Hello!

New to this subreddit.

I just have a quick question when writing sql queries in a rails application.

Where and how do we write the queries? Where should the queries go? In the Models or Controllers?

\--EDIT--

So in an interview setting, with rails and SQL. How do you think they will make me write sql queries in rails?

I know a little bit of sql but I don't know how I will be tested when creating a rails application and using sql.

I appreciate all the help!

Thank you so much!
## [5][I created a step-by-step tutorial demonstrating how to integrate React with Ruby on Rails](https://www.reddit.com/r/rails/comments/f7a0v8/i_created_a_stepbystep_tutorial_demonstrating_how/)
- url: https://www.reddit.com/r/rails/comments/f7a0v8/i_created_a_stepbystep_tutorial_demonstrating_how/
---
- [Live Demo](https://rails-react-example.herokuapp.com/)
- [Tutorial](https://stevepolito.design/blog/rails-react-tutorial/)

I really wanted to learn React and API development, so I went head first into building a simple application, and documented my experience. I think what sets this apart from other Rails and React tutorials is that I cover...

- API authorization
- API versioning
- Setting HTTP status codes
- Form validation on the front-end
- Handling errors
- Debouncing requests
- CSRF Countermeasures
## [6][Help with Active Record query](https://www.reddit.com/r/rails/comments/f7khml/help_with_active_record_query/)
- url: https://www.reddit.com/r/rails/comments/f7khml/help_with_active_record_query/
---
I'm working on an application and struggling with creating an Active Record query.

I'm working on a reservation system that allows the owner to blackout certain dates on certain properties. So, to generate availability across all properties for a specific date, I need to first see if there are any  property\_blackout\_dates on that date.

    class PropertyBlackoutDate &lt; ApplicationRecord
      belongs_to :property
    end 
    
    class Property &lt; ApplicationRecord
      has_many :property_blackout_dates
    end 

Since I want all the Properties without a property\_blackout\_date for a specific date, I tried doing this:

    date = "2020-02-22"
    Property.includes(:property_blackout_dates).where.not(property_blackout_dates: {date: date}).references(:property_blackout_dates)

It just returns an empty association, but

    date = "2020-02-22"
    PropertyBlackoutDate.where(date: date)

does return an object, so I don't think it's an issue with the date parameter being wrong.

What am I doing wrong?
## [7][How do I validate foreign keys?](https://www.reddit.com/r/rails/comments/f7jr23/how_do_i_validate_foreign_keys/)
- url: https://www.reddit.com/r/rails/comments/f7jr23/how_do_i_validate_foreign_keys/
---
Let's say I have three models:

class Brand  &lt; ***ApplicationRecord***

has\_many :cars

has\_many :dealer\_brands

has\_many :dealers,  through: :dealer\_brands

end

class Dealer &lt; ***ApplicationRecord***

has\_many :cars

has\_many : dealer\_brands # Bridge table that states all the brands that the dealer carries.

has\_many :brands,  through: :dealer\_brands

end

class Car  &lt; ***ApplicationRecord***

belongs\_to :dealer

belongs\_to :brand

end

When I create a car, how do I validate and make sure that the dealer carries a particular brand before saving the car entry? Do I just use a validation in the car model to make sure that the dealer carries the brand before saving the record into the database or is there a better way?
## [8][Different assets for different route namespaces](https://www.reddit.com/r/rails/comments/f7lcj2/different_assets_for_different_route_namespaces/)
- url: https://www.reddit.com/r/rails/comments/f7lcj2/different_assets_for_different_route_namespaces/
---
Hello, I wanted to ask how r/rails would implement this. 

I have a rails app that has 2 distinct sections. A frontend/customer facing side &amp; a backend/vendor side. I have 2 bootstrap themes that I like. However, they obviously have css name conflicts on card or some other basic names. 

My idea is to do a before_action in the application controller. If the controller namespace is admin, set '@namespace' to 'admin', else, 'frontend'. Then in the application.html do `javascript_pack_tag @namespace`. 

Another idea is to make the vendor area a separate rails application connected to the main DB but I would like to keep everything in the same repo.

Tbh, I'm not really feeling this idea. Is there a better way to implement this?
## [9][Why I don't use React with Rails](https://www.reddit.com/r/rails/comments/f7rzxg/why_i_dont_use_react_with_rails/)
- url: https://www.reddit.com/r/rails/comments/f7rzxg/why_i_dont_use_react_with_rails/
---
I'd like to share the top reasons why I don't use React:

* React is meant for large teams as it was built for Facebook which is a large company.
* React forces you to build and run \*two\* applications just for the browser: micro-service for the back-end, and a \*separate\* UI front-end.
* Designers pay a complexity tax when working with JSX.
* Developers pay a complexity tax when working with JSX.
* You can't use the keyword "class" in your HTML.

To me, React literally defeats the largest principle in Rails which is being one \*monolith\* application.

I'd like to have a high-quality discussion about this important topic. So please share your thoughts if you don't agree.

Down-voting without commenting on \*why\* you don't agree is a disservice for the community.  JavaScript is a huge part of our Rails applications, and the way we choose to go about it is fundamental.
## [10][Converted Date object to integer for purposes of a graph, struggling to get that integer back to the original date](https://www.reddit.com/r/rails/comments/f7hvzs/converted_date_object_to_integer_for_purposes_of/)
- url: https://www.reddit.com/r/rails/comments/f7hvzs/converted_date_object_to_integer_for_purposes_of/
---
So I'm working with highcharts, and for storage of plot point data, I converted a Date object to an integer to get it to display properly in highcharts.  But in the `tooltip` `pointFormat`, I am trying to interpolate a data xaxis data point, `point.x` for example into that popup when you hover over the datapoint(which I guess is what `pointFormat` and `tooltip` are for).  It needs to be a string it appears, so for that purpose, I need to get back _from_ the integer date stored in the data points data, to the original date object, or perhaps as a string date.

**TLDR**

I have a date object, such as `Mon, 31 Dec 2018`, and I converted it to an integer for the highchart datapoint plot, via

    my_date_object.strftime('%Q')

which leads to `"1545214400000"`.

When I try to go back to a date object (or perhaps I need a string, I just need to plain interpolate into a string) with month, day, year, I found and tried this, but I get some whacky year when trying to convert back(and I don't want the time appended at the end either to boot):

    Time.at(1545214400000).to_formatted_s(:long)
    =&gt; "August 01, 50967 19:00"

The above is way different than my original `Mon, 31 Dec 2018` date object, that went through conversion to an integer originally.

Any help appreciated, thanks!
## [11][Would it be Ethical to ask people to pitch in?](https://www.reddit.com/r/rails/comments/f7fzbr/would_it_be_ethical_to_ask_people_to_pitch_in/)
- url: https://www.reddit.com/r/rails/comments/f7fzbr/would_it_be_ethical_to_ask_people_to_pitch_in/
---
I am developing this pet project of mine, which I would think be useful to many. I already envision it to be free for all, but you got to pay small monthly sub (a couple of euros) if you want to see your "Reports", an advanced feature. Would it be Ethical for me to put it out on Kickstarter and ask people for donations towards the final product? I am unemployed now, just working on some stuff to make companies hire me :) But I still want to eat meanwhile.
## [12][Feature flag gems](https://www.reddit.com/r/rails/comments/f7f1w6/feature_flag_gems/)
- url: https://www.reddit.com/r/rails/comments/f7f1w6/feature_flag_gems/
---
We're thinking about implementing feature flags for control of UI gated features.  There looks like a handful of them out there (Flipper, Flipflop, rollout etc).  Anyone used these before?  Which gems are good?  What are any of the gotchas you've encountered?
