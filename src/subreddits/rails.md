# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/f9t9kq/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/f9t9kq/personal_projects_show_off_your_own_project_andor/
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
## [2][Gimme Gems Thursdays - Found an awesome new gem? Post it here!](https://www.reddit.com/r/rails/comments/fduax9/gimme_gems_thursdays_found_an_awesome_new_gem/)
- url: https://www.reddit.com/r/rails/comments/fduax9/gimme_gems_thursdays_found_an_awesome_new_gem/
---
Please use this thread to discuss **cool** but relatively **unknown** gems you've found.

You **should not** post popular gems such as [those listed in wiki](https://www.reddit.com/r/rails/wiki/index#wiki_popular_gems) that are already well known.

Please include a **description** and a **link** to the gem's homepage in your comment.
## [3][How to view all available methods in template or controller?](https://www.reddit.com/r/rails/comments/ff8spe/how_to_view_all_available_methods_in_template_or/)
- url: https://www.reddit.com/r/rails/comments/ff8spe/how_to_view_all_available_methods_in_template_or/
---
Hi, I have been learning Rails for about a week and I am having very hard time with understanding which methods can be used when because everything seems like magically appearing in templates and controllers.

Is there a way to dump all available methods in debug statement or something to see and better understand them?

Also, if you have other tips on understanding this topic, your comment would be really helpful!
## [4][How do I handle this join?](https://www.reddit.com/r/rails/comments/ff3raq/how_do_i_handle_this_join/)
- url: https://www.reddit.com/r/rails/comments/ff3raq/how_do_i_handle_this_join/
---
I have a side project with an interesting problem. I have two tables:

1. Books
2. Restrictions

I want to be able to restrict certain books with a scope. I'm pretty new to joins, and I feel like this would be easy, *except* I also have another unique constraint. I have set up restrictions to be polymorphic. So, there is a record\_id and record\_type column on restrictions. This is so, that in the future, I have the ability to restrict other things than just books.

My attempt is something like this:

`Book.all.joins("INNER JOIN restrictions ON restrictions.record_id =` `books.id` `AND restrictions.record_type = 'Book'")`

But this is just returning all books, and I am trying to get a list of books that isn't restricted. Any ideas?
## [5][How do I efficiently query telemetry data?](https://www.reddit.com/r/rails/comments/feswoh/how_do_i_efficiently_query_telemetry_data/)
- url: https://www.reddit.com/r/rails/comments/feswoh/how_do_i_efficiently_query_telemetry_data/
---
I have this project where there are N devices in the field that beam back telemetry data **once every minute** per device. The ingest API for that is fine, but I'm running into some concerns about _reading_ that data, specifically for building a graph.

Data is in PostgreSQL and roughly resembles this:

```
    create_table :telemetry do |t|
      # some unrelated stuff here
      t.float         :some_value, null: false, default: 0.0
      t.timestamps
    end
```

So there's one row per minute per device in this database. At 24 hours in one day times 60 rows per hour (one per minute), we have **1,440 rows of telemetry data!**

Now, querying this is easy enough if we don't care about how evenly distributed that data is:

```
@telemetry = Telemetry.where('created_at &gt; ?', 24.hours.ago).order(created_at: :asc)
```

I could of course throw a limit on that, but then we'd just get the first N minutes of the data. So for example if I were to limit that to 50 rows, I'd only get back less than the first hour of all 24 hours.

Is there a better way to do this? For example, let's say I want two data points per hour (so one every 30 minutes). I _could_ try creating a loop then using the loop to craft and execute a new query for every trip through said loop, but that also sounds pretty inefficient.

How do I _efficiently_ do something like this? What should I be googling for?
## [6][How do I delete a record and it's associated records and only skip one callback action for each, that sends a notification?](https://www.reddit.com/r/rails/comments/fethwh/how_do_i_delete_a_record_and_its_associated/)
- url: https://www.reddit.com/r/rails/comments/fethwh/how_do_i_delete_a_record_and_its_associated/
---
So I have a record that I would like to delete, and it has many associated records. And every record has multiple after_destroy callbacks including one that sends a notification to the affected users.

The problem here is that I want to fire all the callbacks actions but only skip the one that sends the notification. How can I achieve that?
## [7][Question about wrapping destroy inside a transaction](https://www.reddit.com/r/rails/comments/fepzsn/question_about_wrapping_destroy_inside_a/)
- url: https://www.reddit.com/r/rails/comments/fepzsn/question_about_wrapping_destroy_inside_a/
---
I need to call two separate methods that do a commit to the database and I'm wondering if I wrap these methods inside a transaction will it work properly. 

    ActiveRecord::Base.transaction do 
        destroy_roles
        save_user
    end

and destroy roles looks like this 

    def destroy_roles
      current_user.roles.destroy(role_to_be_destroy)
    end

and save\_user is 

    def save_user 
      current_user.update_attributes(user_params)
    end 

So when `destroy_roles` is called within the transaction the role is deleted with `begin` and `commit` which will usually succeed but it's unlikely the `save_user` will succeed which is why I want `destroy_roles` to be rolled back but since each `destroy` happens withing a `begin` and `commit` I'm wondering if this will be rolled back if `save_user` fails.  

I also have another question about if we should user the `!` version inside the transaction like `update_attributes!` as opposed to `update_attribute` and so on. What difference would it make inside a transaction?

P.S in impossible to avoid writing this as a single commit as the roles have to be deleted first so that the before\_update callbacks have the correct roles for further processing before user is saved
## [8][Video as a medium for rails](https://www.reddit.com/r/rails/comments/feri5d/video_as_a_medium_for_rails/)
- url: https://www.reddit.com/r/rails/comments/feri5d/video_as_a_medium_for_rails/
---
Hey guys I’m starting my journey into rails. The one problem I have is the lack of high quality videos out there that show everything. I come from front end and there’s so many videos on setting up react and your environment and basically projects and they show how everything is done and why. With rails there’s only these huge 5-7 hour projects with no set up or anything. 

My question is do you guys have any good YouTubers you recommend that helped you? 

I looked through the side bar and history and everyone just says read books and stuff but I learn the best with videos. Thanks guys
## [9][Lost master key rails 6](https://www.reddit.com/r/rails/comments/feo3gu/lost_master_key_rails_6/)
- url: https://www.reddit.com/r/rails/comments/feo3gu/lost_master_key_rails_6/
---
Hopefully someone can help me. I put my rails 6 project on github, then did a fresh install of my os deleting everything. After getting my project back on my “new” computer I was trying to add some secret keys in my credentials.yml file but get an error saying my master key doesn’t match. After some googling and no information on how to reset the key. Here I am. Thanks in advance for any help.
## [10][Assigning a test case to an integer doesn't work - how do I test for a non string value?](https://www.reddit.com/r/rails/comments/fendcw/assigning_a_test_case_to_an_integer_doesnt_work/)
- url: https://www.reddit.com/r/rails/comments/fendcw/assigning_a_test_case_to_an_integer_doesnt_work/
---
Hi all,  


I'm new to TDD and I'm trying to write a very simple test - checking if a value is a string.  


Here's the test (which I'm trying to make fail):  


    test "test model name should be a string" do
        test_model(:one).name = 1
        assert_equal true, test_model(:one).name.is_a?(String)
      end
    
    =&gt; 0 Failures

I've tried using byebug to check the value of test\_model(:one).name, and even though I'm assigning as 1, it's actually assigning it as "1" (A string rather than an integer).  


1) Why is it doing this?  
2) How do I test for this case?  


Thanks.
## [11][Data modeling - is single table inheritance approach good here](https://www.reddit.com/r/rails/comments/fegw2e/data_modeling_is_single_table_inheritance/)
- url: https://www.reddit.com/r/rails/comments/fegw2e/data_modeling_is_single_table_inheritance/
---
Hello guys, I need an opinion. I am creating a shopping platform, and basic functionality is as follows: suppliers have a product page, buyers can buy from them.

Other details: 

* I have the Account model which can be of two types: Buyer and Supplier. 

* Product model is associated with Account of the type 'Supplier'. 

* Also, there is User model which belongs to Account (both suppliers and buyers can have Users).

How would you model this?

I was thinking about Single table inheritance - creating Account model which inherits from ActiveRecord, and two Ruby classes for Buyer and Supplier. Putting mutual fields in Account model along with the User association, and different functionality in Account's sub-classes. What bothers me is that I don't know how to avoid having Account of type Buyer associated with Products.

Is this a good approach?
## [12][Keeping controller DRY with methods?](https://www.reddit.com/r/rails/comments/fen6xg/keeping_controller_dry_with_methods/)
- url: https://www.reddit.com/r/rails/comments/fen6xg/keeping_controller_dry_with_methods/
---
Hey guys, wanted some quick thoughts on this:

Suppose my controller has both a "create" and "update" action, and in both actions an instance variable needs to be instantiated the same way:

    \@purchase_date = Date.civil(params[:application]["purchase_date(1i)"].to_i,
    params[:application]["purchase_date(2i)"].to_i, params[:application]["purchase_date(3i)"].to_i)

I think there's three ways to do this:

**\\1. Simply copy paste this code in both create and update action. Improves readability of code but is it DRY?**

    def create
      \@purchase_date = Date.civil(params[:application]["purchase_date(1i)"].to_i,
     params[:application]["purchase_date(2i)"].to_i,
     params[:application]["purchase_date(3i)"].to_i)
      // more code
    end
    
    def update
      \@purchase_date = Date.civil(params[:application]["purchase_date(1i)"].to_i,
     params[:application]["purchase_date(2i)"].to_i,
     params[:application]["purchase_date(3i)"].to_i)
      // more code
    end

**\\2. Make a private method in controller that populates the instance variable and just stick the method in both create and update**

    def create
      set_purchase_date
      // more code
    end
    
    def update
      set_purchase_date
      // more code
    end
    
    private
    
    def set_purchase_date
      @purchase_date = Date.civil(params[:application]["purchase_date(1i)"].to_i, params[:application]["purchase_date(2i)"].to_i, params[:application]["purchase_date(3i)"].to_i) 
    end

**\\3. Make a private method in controller that returns the object needed by instance variable and make instance variable equal to that**

    def create
      \@purchase_date = set_purchase_date
      // more code
    end
    
    def update
      \@purchase_date = set_purchase_date
      // more code
    end
    
    private
    
    def set_purchase_date
      Date.civil(params[:application]["purchase_date(1i)"].to_i, params[:application]  ["purchase_date(2i)"].to_i, params[:application]["purchase_date(3i)"].to_i) 
    end

&amp;#x200B;

What's the best way to go here? Which do you prefer? Is there another way?

&amp;#x200B;

Edit: Ignore the \\@ thing, I have no idea how to escape the auto formatting reddit does for @
