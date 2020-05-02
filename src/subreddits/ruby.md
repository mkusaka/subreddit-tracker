# ruby
## [1][How to build first telegram bot with Ruby? 8 lines of code and 5 minutes is enough!](https://www.reddit.com/r/ruby/comments/gc3g2v/how_to_build_first_telegram_bot_with_ruby_8_lines/)
- url: https://youtu.be/WefRd8W5KuA
---

## [2][Can't use ruby installed gem](https://www.reddit.com/r/ruby/comments/gc5jmu/cant_use_ruby_installed_gem/)
- url: https://www.reddit.com/r/ruby/comments/gc5jmu/cant_use_ruby_installed_gem/
---
I am new to ruby and would like to use the ruby gem [Word-Search Puzzle Generator](https://www.rubydoc.info/gems/wordsearch-puzzle)

I am using Mac OS X 10.14.6.

This is what I did:

&amp;#x200B;

1. Updated brew using `brew update &amp;&amp; brew upgrade`
2. Ran `brew doctor` to ensure there were no problems (no problems were found).
3. Ran `brew install ruby`
4. Ran echo `'export PATH="/usr/local/opt/ruby/bin:$PATH"' &gt;&gt; ~/.bash_profile`
5. Ran `source ~/.bash_profile`
6. Ran `gem install wordsearch`. This returned the following error:  
`ERROR:  While executing gem ... (Errno::EACCES)`

`Permission denied @ rb_sysopen - /usr/local/lib/ruby/gems/2.7.0/gems/wordsearch-1.0.0/ext/wordsearch/accent.c`

7. So I tried `sudo gem install wordsearch`. This worked.

8. However, when I run w`ordsearch nitwit blubber oddment tweak` it says ***command not found.***

Any ideas? Thanks!
## [3][Seeking advice on the Ruby language learning curves](https://www.reddit.com/r/ruby/comments/gc1pjc/seeking_advice_on_the_ruby_language_learning/)
- url: https://www.reddit.com/r/ruby/comments/gc1pjc/seeking_advice_on_the_ruby_language_learning/
---
For the people that already into the in-depth level of the Ruby language, what is your thought on the current standing of Ruby? and what was the point of realization that makes you feel like the needs for Ruby?

I have had experiences with the Ruby on Rails framework and only come to learn the Ruby language on-demand for debugging. I have been spoiled by the framework convention itself and left me with an uncertain/unfulfillment feeling. 

I'm all ear for any advice and starting-point recommendation that might not just simply on the topic of learning the coding language.
## [4][Guild renamed to Ractor](https://www.reddit.com/r/ruby/comments/gbhkjn/guild_renamed_to_ractor/)
- url: https://www.reddit.com/r/ruby/comments/gbhkjn/guild_renamed_to_ractor/
---
[https://github.com/ko1/ruby/blob/ractor/ractor.ja.md](https://github.com/ko1/ruby/blob/ractor/ractor.ja.md)
## [5][Socket programming / select](https://www.reddit.com/r/ruby/comments/gbiyaq/socket_programming_select/)
- url: https://www.reddit.com/r/ruby/comments/gbiyaq/socket_programming_select/
---
Is **select** useful for handling a single socket?

What would be the best way to manage a non-blocking connection of an IRC bot to a single server?
## [6][The Ruby Blend: Parentheses and typosquatting](https://www.reddit.com/r/ruby/comments/gbf4zj/the_ruby_blend_parentheses_and_typosquatting/)
- url: https://www.reddit.com/r/ruby/comments/gbf4zj/the_ruby_blend_parentheses_and_typosquatting/
---
In this episode, Nate has to check his emails, Andrew uses parentheses, and Ron drowns in the DigitalOcean. We also find out why Nate was right (again). Also, what happened with typosquatting? Nate dives into the Manager vs. Maker's schedule. And why does Andrew have "needless paranoia?" \[Listen to this week's episode to find out about all these and more\]([https://fireside.fm/s/ouBAUjGy+4ZxpTNPN](https://fireside.fm/s/ouBAUjGy+4ZxpTNPN)).
## [7][Do accurate Ruby memory profilers exist?](https://www.reddit.com/r/ruby/comments/gb4q3h/do_accurate_ruby_memory_profilers_exist/)
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
## [8][Bloody fast jaro winkler](https://www.reddit.com/r/ruby/comments/gax1nk/bloody_fast_jaro_winkler/)
- url: https://www.reddit.com/r/ruby/comments/gax1nk/bloody_fast_jaro_winkler/
---
I just released batch\_jaro\_winkler: [link](https://github.com/dbousque/batch_jaro_winkler). It is at least 20x faster than the fastest non-batch alternatives (there are no other batch libraries as far as I am aware). The goal is to make *jaro* and *jaro winkler* distance calculations over a set of predefined strings as fast as possible. How are you currently using fuzzy matching? I'm thinking about extending it with new features. In particular, would the ability to add new strings to an exportable model (see [The exportable model](https://github.com/dbousque/batch_jaro_winkler#the-exportable-model)) on the fly be valuable?
## [9][NodeRunner: execute Javascript in a Ruby context via Node](https://www.reddit.com/r/ruby/comments/gaxu1e/noderunner_execute_javascript_in_a_ruby_context/)
- url: https://www.reddit.com/r/ruby/comments/gaxu1e/noderunner_execute_javascript_in_a_ruby_context/
---
Hey all, I just released a new gem called NodeRunner. It provides a simple way to execute Javascript in a Ruby context via Node. (Loosely based on the Node Runtime module from ExecJS.) It supports Node require statements and dynamic function calls, and is nearly as fast as executing Node scripts directly on the command line.

I have some ideas of how it use it as part of a static site generator's build process, but I'm sure there are many other applications. I wouldn't necessarily recommend using it directly in a Rails request/response cycle with high traffic, but it'd be perfect for background jobs.

Let me know what you think!

https://github.com/bridgetownrb/node-runner
## [10][Why Can't I Manually Resize or Drag My GUI Window? (Shoes 3.3)](https://www.reddit.com/r/ruby/comments/gaulro/why_cant_i_manually_resize_or_drag_my_gui_window/)
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
