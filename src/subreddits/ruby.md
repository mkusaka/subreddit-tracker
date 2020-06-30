# ruby
## [1][Rails 6.1 tracks Active Storage variant in the database](https://www.reddit.com/r/ruby/comments/hii694/rails_61_tracks_active_storage_variant_in_the/)
- url: https://blog.bigbinary.com/2020/06/30/rails-6-1-tracks-active-storage-variant-in-the-database.html
---

## [2][Bubble sort](https://www.reddit.com/r/ruby/comments/hike73/bubble_sort/)
- url: https://www.reddit.com/r/ruby/comments/hike73/bubble_sort/
---
Hi! Beginner here.

So currently I'm trying to write a simple bubble sort code. In fact it works pretty well, but there's a problem with sorting at this point: gets.chomp.split(",") .

    def bubble_sort(array)
    n = array.length
    loop do 
    swapped = false
    (n-1).times do |i| 
    if array[i] &gt; array[i+1] 
    array[i], array[i+1] = array[i+1], array[i] 
    swapped = true end
    end
     break if not swapped
    end 
    array 
    end
    puts "Type your array:" array = gets.chomp.split(",")
    p bubble_sort(array)

When I type for example: 100,1,10 the output is correct. But then if I type: 10,1,4,66,0 the output is \["0", "1", "10", "4", "66"\].
## [3][Fun Facts about Ruby #7](https://www.reddit.com/r/ruby/comments/hik1wb/fun_facts_about_ruby_7/)
- url: https://i.redd.it/vb5p3jt4o0851.png
---

## [4][Build The World’s Simplest ETL (Extract, Transform, Load) Pipeline in Ruby With Kiba](https://www.reddit.com/r/ruby/comments/hi6loi/build_the_worlds_simplest_etl_extract_transform/)
- url: https://towardsdatascience.com/build-the-worlds-simplest-etl-extract-transform-load-pipeline-in-ruby-with-kiba-e7093a29d35
---

## [5][The Ruby Spec Suite Compatibility Report](https://www.reddit.com/r/ruby/comments/hi5xnh/the_ruby_spec_suite_compatibility_report/)
- url: https://eregon.me/blog/2020/06/27/ruby-spec-compatibility-report.html
---

## [6]["Feels good to be in Rails right now"](https://www.reddit.com/r/ruby/comments/hhm3uj/feels_good_to_be_in_rails_right_now/)
- url: https://www.reddit.com/r/ruby/comments/hhm3uj/feels_good_to_be_in_rails_right_now/
---
New rails tech is coming around and new big projects are launched on rails (like basecamp has released the new awesome [hey.com](https://hey.com) email platform running on rails without any major JS frameworks).

Chris Oliver from gorails discussed it in the [freshest remote ruby podcast](https://remoteruby.transistor.fm/83) and said...

&gt;"It feels good to be in Rails right now"

Do you share this feeling?
## [7][Is there a cleaner way of splitting strings multiple times](https://www.reddit.com/r/ruby/comments/hhyjzx/is_there_a_cleaner_way_of_splitting_strings/)
- url: https://www.reddit.com/r/ruby/comments/hhyjzx/is_there_a_cleaner_way_of_splitting_strings/
---
Hi all,

Newbie at Ruby here, created a small script to do some media conversion to allow my TV to play the media files. Most of it I'm ok with apart from the new\_filename creation, is there a cleaner way to split the pathname down multiple times?

Also any general feedback on construction would be appreciated as well.

&amp;#x200B;

    # frozen_string_literal: true
    
    require 'pathname'
    require 'streamio-ffmpeg'
    require 'slop'
    
    # Set path to recursively search directories
    opts = Slop.parse do |o|
      o.string '-d', '--directory', 'a directory to search'
      o.on '-h', '--help' do
        puts "Usage: ruby Check-MediaFiles.rb '/enter/path/to/search/'"
        exit
      end
    end
    
    path = opts[:directory]
    
    # Search all subdirectories for mp4 and mkv files to process and begin loop
    Pathname.glob(path + '/*/*.{mp4,mkv,avi}').each do |file|
      # Change to each directory
      Dir.chdir(file.dirname)
    
      # Instantiate movie object for processing
      media = FFMPEG::Movie.new(file.to_s)
    
      # Create new filename from metadata
      temp1_filename = file.basename.to_s
      temp2_filename = temp1_filename.split('/')[-1]
      temp3_filename = temp2_filename.split('.')[0]
      new_filename = temp3_filename + '.mp4'
    
      # Check got ac3 audio in media file and convert if not found.
      if file.extname == '.avi'
        puts 'AVI detected. Re-encoding to mp4 for compatibility with Samsung TV.'
        media.transcode(new_filename.to_s, %W[-c:v copy -c:a ac3])
        File.delete(file)
        next
      elsif media.audio_codec != 'ac3'
        puts "AC3 audio codec not detected on #{file.basename} converting audio to AC3."
        media.transcode(new_filename.to_s, %w[-c:v copy -c:a ac3])
        puts "Completed conversion to ac3 audio for #{new_filename}"
        File.delete(file)
      else
        puts "No conversion required for #{file.basename}"
      end
    end

&amp;#x200B;
## [8][Noob to Ruby master recommendations](https://www.reddit.com/r/ruby/comments/hhtfl0/noob_to_ruby_master_recommendations/)
- url: https://www.reddit.com/r/ruby/comments/hhtfl0/noob_to_ruby_master_recommendations/
---
I'm just getting started with programming, my current job is willing to pay for me to get a course or attend a conference. All though I'll probably would go for an interactive course so I can ask questions. Can I get some recommendations from the community? 

I like to learn the following:
Ruby 
Ruby on rails
SQL advance queries
ActiveRecord advance queries

I work with a lot of data so having a depth knowledge of the last two would be awesome. Thank you in advance!
## [9][Just started to publish API-only ruby on rails course. We’re building an Instapaper clone from scratch.](https://www.reddit.com/r/ruby/comments/hhb1ib/just_started_to_publish_apionly_ruby_on_rails/)
- url: https://www.reddit.com/r/ruby/comments/hhb1ib/just_started_to_publish_apionly_ruby_on_rails/
---
I'm happy that I just published the first two chapters of **API-only** ruby on rails course that I [mentioned](https://www.reddit.com/r/ruby/comments/gpebe0/im_planning_to_build_a_rails_course_would_you/) almost one month ago on this subreddit. 🎉🎉

In this course, we will build a functional clone of Instapaper from scratch. Also, it targets the developers already familiar with rails tech stack and has a written format. It's an ongoing course, and I'm planning to share new chapters weekly basis.

You can reach the content from [https://duetcode.io/rails-api-only-course](https://duetcode.io/rails-api-only-course). Even though it's just two chapters, I wanted to share as early as possible to get some insights and feedback from the audience.

Please, let me know if you have any thoughts/feedback about the content or the website in general.
## [10][tapping_device - A new tool that helps you debug with ease](https://www.reddit.com/r/ruby/comments/hh93ss/tapping_device_a_new_tool_that_helps_you_debug/)
- url: https://www.reddit.com/r/ruby/comments/hh93ss/tapping_device_a_new_tool_that_helps_you_debug/
---
Knowing what your objects "exactly" do can be hard when debugging a ruby program. Either using \`puts\` or debuggers requires you to jump between files and insert debugging code everywhere.

That's why I created a gem called [tapping\_device](https://github.com/st0012/tapping_device), which can track objects and print all the traces in a nice format with just 1 line of code.  


https://preview.redd.it/9m2n2eib7l751.png?width=1200&amp;format=png&amp;auto=webp&amp;s=a69b519a2f8ba78a2e576e58ffe3c5f732570a40
