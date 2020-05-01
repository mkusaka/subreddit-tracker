# ruby
## [1][Guild renamed to Ractor](https://www.reddit.com/r/ruby/comments/gbhkjn/guild_renamed_to_ractor/)
- url: https://www.reddit.com/r/ruby/comments/gbhkjn/guild_renamed_to_ractor/
---
[https://github.com/ko1/ruby/blob/ractor/ractor.ja.md](https://github.com/ko1/ruby/blob/ractor/ractor.ja.md)
## [2][The Ruby Blend: Parentheses and typosquatting](https://www.reddit.com/r/ruby/comments/gbf4zj/the_ruby_blend_parentheses_and_typosquatting/)
- url: https://www.reddit.com/r/ruby/comments/gbf4zj/the_ruby_blend_parentheses_and_typosquatting/
---
In this episode, Nate has to check his emails, Andrew uses parentheses, and Ron drowns in the DigitalOcean. We also find out why Nate was right (again). Also, what happened with typosquatting? Nate dives into the Manager vs. Maker's schedule. And why does Andrew have "needless paranoia?" \[Listen to this week's episode to find out about all these and more\]([https://fireside.fm/s/ouBAUjGy+4ZxpTNPN](https://fireside.fm/s/ouBAUjGy+4ZxpTNPN)).
## [3][Do accurate Ruby memory profilers exist?](https://www.reddit.com/r/ruby/comments/gb4q3h/do_accurate_ruby_memory_profilers_exist/)
- url: https://www.reddit.com/r/ruby/comments/gb4q3h/do_accurate_ruby_memory_profilers_exist/
---
We have some tasks that occasionally skirt our memory limitations. I know tallying memory in Ruby is usually a sign Ruby isn't the right choice, but here I am.

In messing around, I was trying to gauge what the profiler was measuring &amp; was surprised by the results.

```
require "memory_profiler"

nums1 = (0...100_000).map.to_a.sample(100_000); nil
nums2 = (1...100_000).map.to_a.sample(100_000); nil

report = MemoryProfiler.report do
  diff = nums1 - nums2
end

report.pretty_print

```

yields

```
Total allocated: 40 bytes (1 objects)
Total retained:  0 bytes (0 objects)
```

Effectively, it only reports allocating the final array.

It runs to fast for this array difference to be a nested loop, but I fail to see how this could be calculated quickly while not allocating some temporary memory.

What am I missing? Does ruby preallocate some space for temp math? Do the allocations happen outside ruby so it can't be profiled? Some other reason the profiler is missing?

It's difficult to assess memory usage without understanding what memory *isn't* being measured.
## [4][Choosing Your First Programming Language: Python and Ruby](https://www.reddit.com/r/ruby/comments/gbgqwj/choosing_your_first_programming_language_python/)
- url: https://www.artiba.org/blog/choosing-your-first-programming-language-python-and-ruby
---

## [5][Bloody fast jaro winkler](https://www.reddit.com/r/ruby/comments/gax1nk/bloody_fast_jaro_winkler/)
- url: https://www.reddit.com/r/ruby/comments/gax1nk/bloody_fast_jaro_winkler/
---
I just released batch\_jaro\_winkler: [link](https://github.com/dbousque/batch_jaro_winkler). It is at least 20x faster than the fastest non-batch alternatives (there are no other batch libraries as far as I am aware). The goal is to make *jaro* and *jaro winkler* distance calculations over a set of predefined strings as fast as possible. How are you currently using fuzzy matching? I'm thinking about extending it with new features. In particular, would the ability to add new strings to an exportable model (see [The exportable model](https://github.com/dbousque/batch_jaro_winkler#the-exportable-model)) on the fly be valuable?
## [6][NodeRunner: execute Javascript in a Ruby context via Node](https://www.reddit.com/r/ruby/comments/gaxu1e/noderunner_execute_javascript_in_a_ruby_context/)
- url: https://www.reddit.com/r/ruby/comments/gaxu1e/noderunner_execute_javascript_in_a_ruby_context/
---
Hey all, I just released a new gem called NodeRunner. It provides a simple way to execute Javascript in a Ruby context via Node. (Loosely based on the Node Runtime module from ExecJS.) It supports Node require statements and dynamic function calls, and is nearly as fast as executing Node scripts directly on the command line.

I have some ideas of how it use it as part of a static site generator's build process, but I'm sure there are many other applications. I wouldn't necessarily recommend using it directly in a Rails request/response cycle with high traffic, but it'd be perfect for background jobs.

Let me know what you think!

https://github.com/bridgetownrb/node-runner
## [7][Why Can't I Manually Resize or Drag My GUI Window? (Shoes 3.3)](https://www.reddit.com/r/ruby/comments/gaulro/why_cant_i_manually_resize_or_drag_my_gui_window/)
- url: https://www.reddit.com/r/ruby/comments/gaulro/why_cant_i_manually_resize_or_drag_my_gui_window/
---
I just started learning how to code GUIs using Shoes, and I have a problem.  For some reason, I cannot drag my window or resize it manually.  I am using Windows 10, and the code is posted below:

&amp;#x200B;

&gt;\#filename = ask\_open\_file  
&gt;  
&gt;[Shoes.app](https://Shoes.app) :width =&gt; 300, :height =&gt; 165 do  
&gt;  
&gt;	\#para [File.read](https://File.read)(filename)  
&gt;  
&gt;	button "Open File to be Read"  
&gt;  
&gt;end
## [8][[Dumb Question] Struggling to figure out if I can do this block of code better. Looking for advise.](https://www.reddit.com/r/ruby/comments/gb0swr/dumb_question_struggling_to_figure_out_if_i_can/)
- url: https://www.reddit.com/r/ruby/comments/gb0swr/dumb_question_struggling_to_figure_out_if_i_can/
---
&amp;#x200B;

    require 'net/http'
    require 'net/https'
    require 'uri'
    require 'json'
    require 'open-uri'
    
    cm_password_url = URI.parse("&lt;HOSTNAME&gt;/&lt;URI_PATH&gt;/")  
    cm_password_http = Net::HTTP.new(cm_password_url.host, cm_password_url.port)  
    cm_password_http.read_timeout = 5  
    cm_password_request = Net::HTTP::Get.new(cm_password_url.request_uri)
    cm_password_response = cm_password_http.request(cm_password_request) rescue false
    cm_password_resources = JSON.parse(cm_password_response.body) rescue false
## [9][Build a Twitter clone in 10 minutes with Rails, CableReady, and StimulusReflex](https://www.reddit.com/r/ruby/comments/gaaefh/build_a_twitter_clone_in_10_minutes_with_rails/)
- url: https://www.youtube.com/watch?v=F5hA79vKE_E
---

## [10][Railsconf 2020 Couch Edition on Tuesday 5/5](https://www.reddit.com/r/ruby/comments/gajxmi/railsconf_2020_couch_edition_on_tuesday_55/)
- url: http://railsconf.com
---

