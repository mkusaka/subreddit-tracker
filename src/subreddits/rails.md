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
## [2][The case of the disappearing db column :(](https://www.reddit.com/r/rails/comments/h7ww81/the_case_of_the_disappearing_db_column/)
- url: https://www.reddit.com/r/rails/comments/h7ww81/the_case_of_the_disappearing_db_column/
---
I have been programming in RoR since version 2.0, but I have never encountered a problem like this til now.  This is my second RoR API with React as the front end app. RoR v6.02

The problem is in one particular table **Invitations** which has, among many others, a string column originally named **state**.  I have many other tables with similar columns (same name, same type) that work fine, BUT, the **Invitations** table is in a world of its own.  


* 1 the column does exist in the table.  I can see it and populate it using
pgAdmin 4.
* 2 I can see the field in the log, with the proper values.  
* 3 I have it in my parameters whitelist.  
* 4 I can see it when printing params.inspect.

In the line immediately following the inspection, as in:

    puts "#{params.inspect}"

= "... updated_at"=&gt;"2020-06-12T17:32:46.420Z", "state"=&gt;"invite"} permitted: true&gt;}

    puts  "#{params.state}"

=NoMethodError (undefined method `state' for &lt;ActionController::Parameters...&gt;)

Any code that executes after the whitelist gives me the same NoMethodError.  For example,

    @state = params.state;
    Invitation.update(invitation_params).

I have looked all over the web for a similar problem but maybe I just don't know the right question.

Lastly, just in case, I renamed the column to **status**, but nothing changed except for the column name in the logs.

Any ideas will be appreciated.

P.S.  I have been working in digital electronics since 1954, with a degree in Electrical Engineering.  The first machine I worked on was a LGP-30 with 8k words of drum memory.  The OS was a graduate student with a cot (and miles of paper tape).  I am probably among the few survivors who knows how to program around drum latency :) /brag
## [3][Rails Routes in Rails 6 looks kinda terrible in the terminal](https://www.reddit.com/r/rails/comments/h7lozw/rails_routes_in_rails_6_looks_kinda_terrible_in/)
- url: https://www.reddit.com/r/rails/comments/h7lozw/rails_routes_in_rails_6_looks_kinda_terrible_in/
---
Does anyone have a solution for the rails routes command that doesn't require grepping a specific controller or using the expand flag, that would pretty up the output of the rails routes command? It seems to be thrown off by action mail and active storage routes because of their length. Sure, I could use the expanded flag but then the console window can go for days. I would like to just see the formatting locked like it used to be or maybe even suppress the action mail stuff. 

https://preview.redd.it/5pyumixzoh451.png?width=1706&amp;format=png&amp;auto=webp&amp;s=c4e08fe5012b36a9f1957842de27ba785805d509
## [4][heroku[logplex]: Error L10 (output buffer overflow): drain 'd.cfcd3264-7259-4b12-ac5b-016032978174' dropped 1 messages since 2020-06-12T14:40:59.280571+00:00.](https://www.reddit.com/r/rails/comments/h7lyqq/herokulogplex_error_l10_output_buffer_overflow/)
- url: https://www.reddit.com/r/rails/comments/h7lyqq/herokulogplex_error_l10_output_buffer_overflow/
---
  
How can I even begin troubleshooting this? it would help to know what the contents of the dropped messages were, or at least which method/process/dyno these dropped messages originated from.
## [5][How to handle dynamic, nested WHERE AND/OR queries using active record](https://www.reddit.com/r/rails/comments/h7pxmb/how_to_handle_dynamic_nested_where_andor_queries/)
- url: https://www.reddit.com/r/rails/comments/h7pxmb/how_to_handle_dynamic_nested_where_andor_queries/
---
I'm currently building a feature that requires me to loop over an hash, and for each key in the hash, dynamically modify an SQL query.

The actual SQL query should look something like this:

```sql
select * from space_dates d
	inner join space_prices p on p.space_date_id = d.id
	where d.space_id = ?
	and d.date between ? and ?
	and (
		(p.price_type = 'monthly' and p.price_cents &lt;&gt; 9360) or
		(p.price_type = 'daily' and p.price_cents &lt;&gt; 66198) or
		(p.price_type = 'hourly' and p.price_cents &lt;&gt; 66198) # This part should be added in dynamically
	)
```

The last `and` query is to be added dynamically, as you can see, I basically need only one of the conditions to be true but not all.

```
query = space.dates
      .joins(:price)
      .where('date between ? and ?', start_date, end_date)

# We are looping over the rails enum (hash) and getting the key for each key value pair, alongside the index

SpacePrice.price_types.each_with_index do |(price_type, _), index|
  amount_cents = space.send("#{price_type}_price").price_cents
  query = if index.positive? # It's not the first item so we want to chain it as an 'OR'
            query.or(
              space.dates
               .joins(:price)
               .where('space_prices.price_type = ?', price_type)
               .where('space_prices.price_cents &lt;&gt; ?', amount_cents)
             )
           else
             query # It's the first item, chain it as an and
               .where('space_prices.price_type = ?', price_type)
               .where('space_prices.price_cents &lt;&gt; ?', amount_cents)
           end
end
```

The output of this in rails is:

```
SELECT "space_dates".* FROM "space_dates"
  INNER JOIN "space_prices" ON "space_prices"."space_date_id" = "space_dates"."id"
  WHERE "space_dates"."space_id" = $1 AND (
   (
     (date between '2020-06-11' and '2020-06-11') AND
     (space_prices.price_type = 'hourly') AND (space_prices.price_cents &lt;&gt; 9360) OR
     (space_prices.price_type = 'daily') AND (space_prices.price_cents &lt;&gt; 66198)) OR
     (space_prices.price_type = 'monthly') AND (space_prices.price_cents &lt;&gt; 5500)
   ) LIMIT $2 
```

Which isn't as expected. I need to wrap the last few lines in another set of round brackets in order to produce the same output. I'm not sure how to go about this using ActiveRecord.

It's not possible for me to use `find_by_sql` since this would be dynamically generated SQL too.
## [6][Form_for PUT instead of PATCH in the edit form](https://www.reddit.com/r/rails/comments/h7qr9d/form_for_put_instead_of_patch_in_the_edit_form/)
- url: https://www.reddit.com/r/rails/comments/h7qr9d/form_for_put_instead_of_patch_in_the_edit_form/
---
So I am using partial in new and edit part of my User object, and by default, Rails applies \_method as PATCH while submitting the edit option. How can I use PUT instead of the patch.  
I cant add method as then it will be an issue while creating a new user.
## [7][Deploying subdirectory projects to Heroku](https://www.reddit.com/r/rails/comments/h7iazy/deploying_subdirectory_projects_to_heroku/)
- url: https://www.reddit.com/r/rails/comments/h7iazy/deploying_subdirectory_projects_to_heroku/
---
When you have your project services (API backend, frontend, background processing, etc.) in one monorepo you might struggle with a simple way of deploying them to Heroku independently. Here's an article containing a way of doing just that.

TL;DR: The easiest solution is 
```
git subtree push --prefix path/to/app-subdir heroku master
```

To get more details and a more complex solution that provides the ability to set up a deployment from CI and utilizing Heroku Review Apps you can read the article.

https://jtway.co/deploying-subdirectory-projects-to-heroku-f31ed65f3f2
## [8][MVC vs 3-tier architecture](https://www.reddit.com/r/rails/comments/h17sez/mvc_vs_3tier_architecture/)
- url: https://www.reddit.com/r/rails/comments/h17sez/mvc_vs_3tier_architecture/
---
**So this has always been a debatable topic.**  
**Approach 1:** We use the classic rails MVC architecture with Server Side rendering.  
**Approach 2:** We use client-side rendering, with rails working as a backend providing the JSON response and client to be made using react or any frontend framework.

  
I would love to get some insights and have a discussion on what are the pros and cons of both these approaches.  


Some of my thoughts are:  
\* Using client-side rendering restricts us to leverage the magic provided by rails. (form auto completions, an automatic mapping between endpoint and controller methods)  
\* Client-side rendering makes the application a bit lighter, as the server is not responsible for rendering the UI for every instance.  
\* I personally am not very comfortable in Js, so using Rails MVC is always an easy way out for me.  


Let's discuss this!
## [9][Can I use a conditional to toggle between rendered partials?](https://www.reddit.com/r/rails/comments/h79ljs/can_i_use_a_conditional_to_toggle_between/)
- url: https://www.reddit.com/r/rails/comments/h79ljs/can_i_use_a_conditional_to_toggle_between/
---
I'm trying to render a partial/hide another partial on a page when a button is clicked. Is this possible with a conditional? Or possible at all? Something like:   


    -if @toggle == 'show'
      = render :partial =&gt; "partial_one"
    
    -elsif @toggle == 'hide'
      = render :partial =&gt; "partial_two"
## [10][Implementing search, should I go for Algolia or is Rails/Postgres enough?](https://www.reddit.com/r/rails/comments/h14rpx/implementing_search_should_i_go_for_algolia_or_is/)
- url: https://www.reddit.com/r/rails/comments/h14rpx/implementing_search_should_i_go_for_algolia_or_is/
---
As the title says, I am implementing search functionality on a social network project.

I want it to feel snappy and provide good suggestions on typing and since I just came back to Rails after a long hiatus in javascript land, I'm not sure if Rails provides something for this out of the box or if should use something like Algolia which is a search-as-a-service to save me the headache.

Any ideas/feeback are welcome!

&amp;#x200B;

Regards
## [11][Appropriate way to validate association](https://www.reddit.com/r/rails/comments/h128cd/appropriate_way_to_validate_association/)
- url: https://www.reddit.com/r/rails/comments/h128cd/appropriate_way_to_validate_association/
---
I'm building a Rails API app that deals with orchestrating Users within a Company, and I want to check my logic for attaching Settings to a User.

The idea is, an admin at a Company will create a Settings entry in the database, and then assign that to a bunch of users.

&amp;#x200B;

Because this is an API, I have one call to create a Setting, one call to create a User, and one call to assign a Setting to a User.

&amp;#x200B;

Here's what the User controller looks like:

    class UsersController &lt; ApplicationController
    
      def set_setting
        if @user.update(setting: Setting.find(params[:setting_id]))
          render :show
        else
          render json: @user.errors, status: :unprocessable_entity
        end
      end
    
    end

And here's how I'm doing the validation in the User model:

    class User &lt; ApplicationRecord
      belongs_to :organisation
      has_one :user_setting
      has_one :setting, through: :user_setting
    
      validate :setting_in_organisation, on: :update
    
      private
    
        def setting_in_organisation
          errors.add(:setting_id, "You can't assign a setting that's not part of your organisation") unless setting.organisation == organisation
        end
    end

So, my question is, is this a good way to approach the issue while sticking to Rails ethos and DRY principals? I think it's important to highlight the views for this app are handled in a different repo. And because it's an API I definitely want a separate call to assign a Setting to a User, rather than chucking everything in user\_params.

What's everyone's thoughts?
