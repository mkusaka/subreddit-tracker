# rails
## [1][Personal Projects - Show off your own project and/or ask for advice](https://www.reddit.com/r/rails/comments/gek2rk/personal_projects_show_off_your_own_project_andor/)
- url: https://www.reddit.com/r/rails/comments/gek2rk/personal_projects_show_off_your_own_project_andor/
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
## [2][How to handle social login from a react native app to rails api](https://www.reddit.com/r/rails/comments/givaz1/how_to_handle_social_login_from_a_react_native/)
- url: https://www.reddit.com/r/rails/comments/givaz1/how_to_handle_social_login_from_a_react_native/
---
Hi guys

I'm trying to set up social login for my react native app which has a rails backend.

I've read a lot of tutorials but most of them focus on rails as a web app solution, not so much as an api. For example, I think I might be needing the omniauth gem but I'm not sure if I also need devise when I'm not making use of most of its features (views, controllers, ...) and when I'm just issuing JWT's.  


So I'm not sure what else I need or how best to approach this.

I was also contemplating using something like auth0 if using something like omniauth is not secure/scalable enough.

&amp;#x200B;

Any thoughts or help would be greatly appreciated!  
Stay safe!
## [3][How to create file(output.pdf) to table database in rails console?](https://www.reddit.com/r/rails/comments/givqm5/how_to_create_fileoutputpdf_to_table_database_in/)
- url: https://www.reddit.com/r/rails/comments/givqm5/how_to_create_fileoutputpdf_to_table_database_in/
---
I have Document table having a column of file:binary status:integer and expiration\_date:datetime.   
Document.create(file: Rack::Test::UploadedFile.new("#{Rails.root}/spec/fixtures/files/testing\_file\_#{rand(100)}.csv", 'application/csv'), status: 0, expiration\_date= "2020-05-13 16:48:21")  
After that, when i want to see the record the file is nil and the rest column has a record.
## [4][Deploying Rails 6 w/ Webpacker to Nginx VPS](https://www.reddit.com/r/rails/comments/gitcbu/deploying_rails_6_w_webpacker_to_nginx_vps/)
- url: https://www.reddit.com/r/rails/comments/gitcbu/deploying_rails_6_w_webpacker_to_nginx_vps/
---
I have a Rails 6 project that uses webpacker to manage its assets which I'm trying to deploy in a VPS w/ nginx + passenger as my application server. However, the css assets won't seem to display properly. JS files seem to be working ok.

What I tried to do:

```

$ RAILS_ENV=production bundle exec rails assets:precompile

$ RAILS_ENV=production RACK_ENV=production NODE_ENV=production ./bin/webpack

```

Why won't my css load correctly?
## [5][Transacted db:migrate:redo rake task](https://www.reddit.com/r/rails/comments/givoeb/transacted_dbmigrateredo_rake_task/)
- url: https://www.reddit.com/r/rails/comments/givoeb/transacted_dbmigrateredo_rake_task/
---
If you are familiar with the `db:migrate:redo` rake task you've probably noticed how it's convenient for fixing one migration, usually that's latest one.

If run without params, at first the task rolls back the latest applied migration, i.e. runs `down` , and then immediately migrates it, i.e. runs `up`. Now, guess what happens if `up` step fails? Right, the `down` step remains performed, what means the latest applied migration now is not the one was a moment ago but the previous one. Next `db:migrate:redo` run, usually after the latest migration fix, surprisingly tries to run the previous migration. That's something unexpected, drives made, and makes the task useless in the scenario when a migration is being fixed that usually requires numerous of changes and `redo`'ing it several times.

I promote a [fix](https://github.com/rails/rails/pull/39026) for that problem. That wraps entire `db:migrate:redo` task into a transaction. It won't change current migration to the previous one when `down` step fails as in the scenario above. That makes fixing a certain migration with more comfort since redo task is idempotent and can run as many times as needed.

This post aims to gather feedback on the specified problem and the [pull request](https://github.com/rails/rails/pull/39026).

All participants, thank you for your attention!
## [6][How do sites like Freecodecamp work?](https://www.reddit.com/r/rails/comments/gixx9u/how_do_sites_like_freecodecamp_work/)
- url: https://www.reddit.com/r/rails/comments/gixx9u/how_do_sites_like_freecodecamp_work/
---
How do I go about building a coding challenge site like freecodecamp. I know how to embed code editors using code mirror or ace.c9. But these editors don’t compile the code and I also don’t know how to add the “run the tests” button.

Where can I learn to implement a site like this (preferably using ruby on rails)? Is there anything I need to learn or is there any API or ruby gem I can use?
## [7][Useful Rake Tasks - Submit yours :-)](https://www.reddit.com/r/rails/comments/gissru/useful_rake_tasks_submit_yours/)
- url: https://www.reddit.com/r/rails/comments/gissru/useful_rake_tasks_submit_yours/
---
I was wondering if anyone has any useful rake tasks they would like to share. Here are a few I find useful.

Check Database connection

    namespace :db  do
      desc "Test if you can connect to the database"
      task :test =&gt; :environment do
        begin
          ActiveRecord::Base.establish_connection # Establishes connection
          ActiveRecord::Base.connection # Calls connection object
          puts "CONNECTED!" if ActiveRecord::Base.connected?
          puts "NOT CONNECTED!" unless ActiveRecord::Base.connected?
        rescue
          puts "NOT CONNECTED!"
        end
      end
    end

Delete All Tables

    namespace :db do
      desc "Erase all tables"
      task :erase =&gt; :environment do
        conn = ActiveRecord::Base.connection
        tables = conn.tables
        tables.each do |table|
    
          puts "Deleting #{table}"
    
          # Disable key checks
          if conn.instance_of?(ActiveRecord::ConnectionAdapters::Mysql2Adapter)
            conn.execute("SET FOREIGN_KEY_CHECKS = 0;")
          else
            conn.execute("ALTER TABLE #{table} DISABLE TRIGGER ALL;")
          end
    
          conn.drop_table(table, force: :cascade)
    
          # Enable key checks
          if conn.instance_of?(ActiveRecord::ConnectionAdapters::Mysql2Adapter)
            conn.execute("SET FOREIGN_KEY_CHECKS = 1;")
          else
            conn.execute("ALTER TABLE #{table} ENABLE TRIGGER ALL;")
          end
        end
    
      end
    end

Quick Rebuild of Database

    namespace :db  do
      desc "Rebuild the data in the databae"
      task :rebuild =&gt; :environment do
        puts "-- Rebuild the Database --"
    
        puts "- Dropping Database"
        Rake::Task["db:drop"].invoke
    
        puts "- Creating Tables"
        Rake::Task["db:create"].invoke
    
        puts "- Migrating"
        Rake::Task["db:migrate"].invoke
    
        puts "- Seeding"
        Rake::Task["db:seed"].invoke
    
        puts "- Finish"
      end
    end

Truncate All Tables

    namespace :db do
      desc "Truncate all tables"
      task :truncate_all =&gt; :environment do
        conn = ActiveRecord::Base.connection
        tables = conn.tables
        tables.each do |table|
          unless ["schema_migrations", "ar_internal_metadata"].include?(table)
            puts "Truncating #{table}"
    
            # Disable key checks
            if conn.instance_of?(ActiveRecord::ConnectionAdapters::Mysql2Adapter)
              conn.execute("SET FOREIGN_KEY_CHECKS = 0;")
            else
              conn.execute("ALTER TABLE #{table} DISABLE TRIGGER ALL;")
            end
    
            conn.execute("TRUNCATE #{table}")
            #conn.drop_table(table, force: :cascade)
    
            # Enable key checks
            if conn.instance_of?(ActiveRecord::ConnectionAdapters::Mysql2Adapter)
              conn.execute("SET FOREIGN_KEY_CHECKS = 1;")
            else
              conn.execute("ALTER TABLE #{table} ENABLE TRIGGER ALL;")
            end
          end
        end
    
      end
    end

I am still fairly new to rake so let me know if any of these can be improved upon
## [8][What you use in RoR for background and cron job?](https://www.reddit.com/r/rails/comments/gikvol/what_you_use_in_ror_for_background_and_cron_job/)
- url: https://www.reddit.com/r/rails/comments/gikvol/what_you_use_in_ror_for_background_and_cron_job/
---
How do you implement background and cron job in RoR ? Do you use Sidekiq or Sneaker, Bunny via RabbitMQ ?
I have read Sidekiq but to get all features we should paid for it :(, does Sidekiq free still Worth it ?

Have anyone using RabbitMQ for this case here ? Not many people out there share the article for this in RoR 

Or do you use Whener gem for cronjob?
## [9][RubyMine IDE "Errors"...this is a long shot. (Double click in console not leading to error, variables no longer changing colour from grey to white when used in code etc)](https://www.reddit.com/r/rails/comments/gim432/rubymine_ide_errorsthis_is_a_long_shot_double/)
- url: https://www.reddit.com/r/rails/comments/gim432/rubymine_ide_errorsthis_is_a_long_shot_double/
---
Hi,

I'm really new to programming and even more so to Ruby. 

I don't know what has happened or what i clicked, but recently my RubyMine IDE stopped highlighting (whitening) variables once I have used them in my code. For example a placeholder in an .each function would previously go from grey to white once I used it; now it does nothing.

In addition, whenever there were errors in the console I could double click and it would take me right there...saving me time and effort.  This functionality has suddenly stopped being available and now I have to go to the line of code manually.

I was wondering what I may have done and how I can remedy this? I know it's a long shot but I would absolutely be appreciative of any assistance.
## [10][How to have auto increment of uuid or default id](https://www.reddit.com/r/rails/comments/gipsf4/how_to_have_auto_increment_of_uuid_or_default_id/)
- url: https://www.reddit.com/r/rails/comments/gipsf4/how_to_have_auto_increment_of_uuid_or_default_id/
---
Do i need to rails migration or edit my migration after i generate a model?
## [11][How to optimize memory usage in Glacier Archiving?](https://www.reddit.com/r/rails/comments/gikdj1/how_to_optimize_memory_usage_in_glacier_archiving/)
- url: https://www.reddit.com/r/rails/comments/gikdj1/how_to_optimize_memory_usage_in_glacier_archiving/
---
Hi all, 

I'm currently running the following script against 2 million database rows in order to move the associated S3 resources to Deep Archive.

However, this process is causing our Kubernetes pods to crash and restart over and over, probably because I'm processing so much data in and out. 

Look at the job pods failing:
```
jobs-954db9b68-6qd5n                 1/1     Running     47         4d5h
jobs-954db9b68-9cmsf                 1/1     Running     249        2d14h
jobs-954db9b68-hz7jr                 1/1     Running     18         9h
jobs-954db9b68-kw8jd                 1/1     Running     182        2d14h
jobs-954db9b68-kw8vs                 1/1     Running     13         7h1m
jobs-954db9b68-mjnvs                 1/1     Running     3          7h1m
jobs-954db9b68-qwbw8                 1/1     Running     61         20h
jobs-954db9b68-sgmq2                 1/1     Running     11         14h
jobs-954db9b68-tx5nr                 1/1     Running     22         9h
jobs-954db9b68-vzcss                 1/1     Running     23         9h
jobs-954db9b68-w879w                 1/1     Running     22         20h
jobs-954db9b68-x9nc5                 1/1     Running     20         9h

```

I've attached the code below, and I've started trying to profile the memory usage using the gems outlined in this excellent guide from Toptal:
https://www.toptal.com/ruby/hunting-ruby-memory-issues

It looks like the memory usage is sailing to 3000MB while profiling? 

What can I do to reduce memory usage and move all this **** to Glacier? Previously, I had a bunch of small jobs, one for each photo_album, and it seemed to create fewer memory issues, but it could just be that some of our customers have larger video files to archive, some of them come in at like 80MB. 

Code snippets below:

```
  def schedule_media_archivals
    photo_albums_for_archival.find_each(batch_size: 1000) do |photo_album|
      photo_album.move_to_archive
    end
  end


    def move_to_archive
    return unless active?
    # We might revisit the open_photo_album line below later but will turn it off for now, see BE-75
    # return update_columns(archive_status: 4) if open_photo_album?

    # If photo_album has no media, list it as restored so that archiving will ignore it and UI hides restore button
    all_media_for_txn = all_media
    return self.update_columns(archive_status: 3) if all_media_for_txn.empty?

    all_media_for_txn.each do |media|
      media.archive
    end

    # prevent "user belongs to different Company" error
    self.update_columns(archive_status: 1) # :archived
  end


    def archive
    return unless self.byte_size&amp;.nonzero? if has_attribute?(:byte_size)
    storage_class = Rails.env == 'production' ? "DEEP_ARCHIVE" : "GLACIER"
    return if self.class.to_s == 'Signature' &amp;&amp; self.media_store._type.to_s == "MediaStores::S3Legacy"
    s3_objects.each do |obj|
      begin
        obj.put({
          acl: "public-read",
          body: obj.get.body.read,
          content_length: obj.content_length,
          content_type: obj.content_type,
          storage_class: storage_class,
        })
      rescue Aws::S3::Errors::NoSuchKey =&gt; ex
        Rails.logger.info "Missing S3 Object: #{self.class}:#{id} #{url} #{ex}"
      rescue Aws::S3::Errors::InvalidObjectState =&gt; ex
        Rails.logger.info "Invalid S3 Object State: #{self.class}:#{id}"
      end
    end
    remove_thumbnail
  end
```
