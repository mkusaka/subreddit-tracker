# ruby
## [1][Benchmark Ruby Code](https://www.reddit.com/r/ruby/comments/j5ik0n/benchmark_ruby_code/)
- url: https://www.driftingruby.com/episodes/benchmark-ruby-code?utm_medium=social&amp;utm_campaign=weekly_episode&amp;utm_source=reddit
---

## [2][Full page caching in Rails with a memcached cluster and middleware](https://www.reddit.com/r/ruby/comments/j4ybw8/full_page_caching_in_rails_with_a_memcached/)
- url: https://vitobotta.com/2020/10/04/full-page-caching-in-rails-part-2-memcached-and-middleware/
---

## [3][Why am I getting a truetrue output on the final line. It should be separated like the others](https://www.reddit.com/r/ruby/comments/j5cibh/why_am_i_getting_a_truetrue_output_on_the_final/)
- url: https://i.redd.it/zglmc86r57r51.jpg
---

## [4][Super Bombinhas v0.7.0](https://www.reddit.com/r/ruby/comments/j4hglr/super_bombinhas_v070/)
- url: https://www.reddit.com/r/ruby/comments/j4hglr/super_bombinhas_v070/
---
Hi!

Here's an updated version of my open source platformer game Super Bombinhas:

[Download on itch.io](https://victords.itch.io/super-bombinhas).

This new version features 5 complete worlds.

Thanks!
## [5][Can someone point to me to the right direction](https://www.reddit.com/r/ruby/comments/j4tjbn/can_someone_point_to_me_to_the_right_direction/)
- url: https://www.reddit.com/r/ruby/comments/j4tjbn/can_someone_point_to_me_to_the_right_direction/
---
This is a lab hw. So when I run the code in my terminal i am getting the following error: 1) Associations — Song and Artist: Artist #add_song adds the song to the current artist's 'songs' collection Failure/Error: expect(artist.songs).to include(song) expected ["In the Aeroplane Over the Sea"] to include #&lt;Song:0x0000000001496e88 @name="In the Aeroplane Over the Sea", @artist=#&lt;Artist:0x0000000001496f50 @name="Neutral Milk Hotel", @songs=["In the Aeroplane Over the Sea"]&gt;&gt; Diff: @@ -1,2 +1,2 @@ -[#&lt;Song:0x0000000001496e88 @name="In the Aeroplane Over the Sea", @artist=#&lt;Artist:0x0000000001496f50 @name="Neutral Milk Hotel", @songs=["In the Aeroplane Over the Sea"]&gt;&gt;] +["In the Aeroplane Over the Sea"]

But when I run my code through pry, I can see the song added into the array of the instance variable, @songs. I am working with two classes, a song class and artist class.

    class Artist 
      attr_accessor :name
      attr_reader :songs 
      @@all = [] 
  
     def initialize(name)
       @name = name 
        save
       @songs = []
       end
     def self.all 
      @@all 
      end
  
      def save 
      @@all &lt;&lt; self 
      end
      def self.destroy_all
       @@all.clear
      end
     def self.create(name)
        self.new(name)
     end

     def add_song(song)
         if song.artist == nil
         song.artist = self
     end 
    #checks if the same song has already been added, otherwise adds new song
    if @songs.include?(song.name) == false 
      @songs &lt;&lt; song.name
    end
        binding.pry
        end
       end

    class Song 
        attr_accessor :name, :artist
       @@all =[]
  
       def initialize(name, artist = nil) 
          @name = name
          @artist = artist
          save
       end
  
       def self. all 
          @@all 
       end
  
       def save 
         @@all &lt;&lt; self 
       end
  
      def self.destroy_all
         @@all.clear 
       end
  
      def self.create(name)
          self.new(name)
      end
    end
## [6][Ruby Gem](https://www.reddit.com/r/ruby/comments/j4sxac/ruby_gem/)
- url: https://www.reddit.com/r/ruby/comments/j4sxac/ruby_gem/
---
Hello All, 

I've been trying to bundle this repo that i cloned from my professor but i keep getting this. 

&amp;#x200B;

Could not find gem 'ruby  (&lt; 2.6, &gt;= 2.2)', which is required by gem 'nokogiri

(\~&gt; 1.8)', in any of the relevant sources:

  the local ruby installation

&amp;#x200B;

I'm on a windows...ive spent hours trying to fix this
## [7][A Little Bit of Automation for your Backups using Ruby](https://www.reddit.com/r/ruby/comments/j4majc/a_little_bit_of_automation_for_your_backups_using/)
- url: https://www.reddit.com/r/ruby/comments/j4majc/a_little_bit_of_automation_for_your_backups_using/
---
I have published a few old of my old Ruby articles at Medium.  Here is the "friends link" so you can view the article without a membership.

[https://medium.com/@bradtrupp/a-little-bit-of-automation-for-your-backups-using-ruby-fccc948b2f7a?source=friends\_link&amp;sk=a986493dc910a34274b5844ab24206ca](https://medium.com/@bradtrupp/a-little-bit-of-automation-for-your-backups-using-ruby-fccc948b2f7a?source=friends_link&amp;sk=a986493dc910a34274b5844ab24206ca)

There are likely thousands of ways to backup your computers data on Windows. One I like, since it is simple, is using a batch file that calls 7zip to create a zip backup with the current date embedded in the zip file name.

This method can be described as “so 1980's” since it is done with batch or command files calling batch programs with command line parameters.
## [8][Bundler](https://www.reddit.com/r/ruby/comments/j4m3c3/bundler/)
- url: https://www.reddit.com/r/ruby/comments/j4m3c3/bundler/
---
Hello all,

I'm working on a HW for class and when I run bundle I get this

&amp;#x200B;

`Warning: the running version of Bundler (1.17.1) is older than the version that created the lockfile (1.17.3). We suggest you upgrade to the latest version of Bundler by running \`gem install bundler\`.`

&amp;#x200B;

`Fetching gem metadata from` [`https://rubygems.org/`](https://rubygems.org/)`...........`

`Fetching gem metadata from` [`https://rubygems.org/`](https://rubygems.org/)`.`

`Resolving dependencies...`

`Bundler could not find compatible versions for gem "ruby ":`

`In Gemfile:`

`ruby  (&gt;= 2.3.1)`

&amp;#x200B;

`rerun (= 0.10.0) was resolved to 0.10.0, which depends on`

`listen (~&gt; 2.7, &gt;= 2.7.3) was resolved to 2.10.1, which depends on`

`celluloid (~&gt; 0.16.0) was resolved to 0.16.0, which depends on`

`timers (~&gt; 4.0.0) was resolved to 4.0.4, which depends on`

`hitimes was resolved to 1.3.0, which depends on`

`ruby  (&lt; 2.6, &gt;= 2.0) x64-mingw32`

&amp;#x200B;

`capybara (= 3.1) was resolved to 3.1.0, which depends on`

`nokogiri (~&gt; 1.8) was resolved to 1.8.4, which depends on`

`ruby  (&lt; 2.6, &gt;= 2.2) x64-mingw32`

&amp;#x200B;

`Could not find gem 'ruby  (&lt; 2.6, &gt;= 2.2)', which is required by gem 'nokogiri`

`(~&gt; 1.8)', in any of the relevant sources:`

`the local ruby installation`

&amp;#x200B;

but I have bundler 2.1.4 installed.

&amp;#x200B;

&amp;#x200B;

https://preview.redd.it/ql1ddojivxq51.png?width=498&amp;format=png&amp;auto=webp&amp;s=ed41b058c995116f5a43bcdcf6b99f6523a66126
## [9][Understanding an example in the new Ractor documentation](https://www.reddit.com/r/ruby/comments/j47oux/understanding_an_example_in_the_new_ractor/)
- url: https://www.reddit.com/r/ruby/comments/j47oux/understanding_an_example_in_the_new_ractor/
---
In the Ractor documentation (https://github.com/ruby/ruby/blob/master/doc/ractor.md), there is an 'examples' section at the bottom that shows various samples. Looking at the 'traditional ring example in actor-model' it has the following code: 

```

    RN = 1000 
    CR = Ractor.current 

    r = Ractor.new do p 
        Ractor.recv CR &lt;&lt; :fin 
    end 

    RN.times{ 
        Ractor.new r do |next_r| 
            next_r &lt;&lt; Ractor.recv 
        end 
    } 

    p :setup_ok 
    r &lt;&lt; 1 
    p Ractor.recv

```

How is this creating a ring? Based on the example name i would expect each ractor to send the message to the next one, but this looks like its creating 1000 Ractors that are all waiting to receive and send something to the first one (r). It then just sends a message only to r. 

Am I misunderstanding this, or is the example incorrect?
## [10][Ruby Monk down?](https://www.reddit.com/r/ruby/comments/j43r2f/ruby_monk_down/)
- url: https://www.reddit.com/r/ruby/comments/j43r2f/ruby_monk_down/
---
Hi all,

Been learning Ruby and Rails through the Odin Project and was using RubyMonk as a secondary resource but I haven't been able to access it for about a week. Just checking to see if others have been experiencing this and if there's any info out there as to why it might be down?

Thanks!
