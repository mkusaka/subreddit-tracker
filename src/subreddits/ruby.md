# ruby
## [1][Optimizing full-text search with Postgres materialized view in Rails](https://www.reddit.com/r/ruby/comments/fctdqz/optimizing_fulltext_search_with_postgres/)
- url: https://caspg.com/blog/optimizing-full-text-search-with-postgres-materialized-view-in-rails
---

## [2][JRuby 9.2.11.0 released! Embedded gems and caller perf regressions fixed](https://www.reddit.com/r/ruby/comments/fckozt/jruby_92110_released_embedded_gems_and_caller/)
- url: https://www.reddit.com/r/ruby/comments/fckozt/jruby_92110_released_embedded_gems_and_caller/
---
The JRuby community is pleased to announce the release of JRuby 9.2.11.0

- Homepage: [http://www.jruby.org/](http://www.jruby.org/)
- Download: [http://www.jruby.org/download](http://www.jruby.org/download)

JRuby 9.2.x is compatible with Ruby 2.5.x and stays in sync with C Ruby.  As always there is a mix of miscellaneous fixes so be sure to read the issue list below.  All users are encouraged to upgrade.

Java Integration
----------------

* Gems and gem paths packaged inside jar files will properly load now. This 9.2.10.0 regression was caused by a change in RubyGems 3 that exposed a bug in JRuby. ([#6060](https://github.com/jruby/jruby/issues/6060), [#6082](https://github.com/jruby/jruby/issues/6082), [#6083](https://github.com/jruby/jruby/issues/6083), [#6084](https://github.com/jruby/jruby/pull/6084))

Performance Improvements
------------------------

* Fixed a major performance and memory bottleneck in backtrace generation, used by `Kernel#caller` and friends. This was a regression in 9.2.10.0. ([#6089](https://github.com/jruby/jruby/issues/6089))

### Github Issues resolved for 9.2.11.0

* [#6097](https://github.com/jruby/jruby/pull/6097) - CGI.escapeHTML restore the original encoding
* [#6093](https://github.com/jruby/jruby/issues/6093) - CGI.escapeHTML crashes on invalid byte sequence where CRuby does not
* [#6091](https://github.com/jruby/jruby/pull/6091) - Check arity earlier in JavaConstructor#new_instance
* [#6090](https://github.com/jruby/jruby/issues/6090) - [9.2.10.0] bogus revision regression
* [#6089](https://github.com/jruby/jruby/issues/6089) - Pontential performance issue
* [#6084](https://github.com/jruby/jruby/pull/6084) - Use expand_path logic for Dir.glob base path
* [#6083](https://github.com/jruby/jruby/issues/6083) - Basic smoke test with warbler no longer works with jruby-jars-9.2.10
* [#6082](https://github.com/jruby/jruby/issues/6082) - Upgrade to 9.2.10.0 breaks project using asciidoctor-maven-plugin
* [#6081](https://github.com/jruby/jruby/issues/6081) - Dependency convergence errors in 9.2.10.0
* [#6079](https://github.com/jruby/jruby/issues/6079) - newInstance on protected/private classes sometimes give an Index OOB Exception
* [#6060](https://github.com/jruby/jruby/issues/6060) - Globbed jar contents with expand path failing in RubyGems
## [3][Writing a Ruby Gem Specification](https://www.reddit.com/r/ruby/comments/fcl2ik/writing_a_ruby_gem_specification/)
- url: https://piotrmurach.com/articles/writing-a-ruby-gem-specification/
---

## [4][solnic.codes / Blog / Open Source Status Update](https://www.reddit.com/r/ruby/comments/fcbdkh/solniccodes_blog_open_source_status_update/)
- url: https://solnic.codes/2020/03/02/open-source-status-update/
---

## [5][[Feeling like a newbie] Help with Hash construction using standard libs.](https://www.reddit.com/r/ruby/comments/fcgdme/feeling_like_a_newbie_help_with_hash_construction/)
- url: https://www.reddit.com/r/ruby/comments/fcgdme/feeling_like_a_newbie_help_with_hash_construction/
---
I am either not asking Google the right questions or I am dense.

I have the following  information that I want to put into a Hash so I can sort and do other things with the output.  This is for managing backup information in AWS/Azure.

Here is the information:

Instance : "INSTANCE\_1"  
   Backup\_ID : "BKUP\_1"  
VOL : "VOL\_1"  
COMMENT: "COMMENT\_1"  
DATE: "DATE\_1"  
   Backup\_ID : "BKUP\_2"  
VOL : "VOL\_1"  
COMMENT: "COMMENT\_2"  
DATE: "DATE\_2"  


The syntax is what is dragging me down.   
Trying to figure this out in IRB:  
bkInst = "INSTANCE\_1"  
bkName\_1 = "BACKUP\_1"  
bkVol = "VOL\_1"  
bkCmt\_1 = "COMMENT\_1"  
bkDate\_1 = "DATE\_1"  
bkName\_2 = "BACKUP\_2"  
bkCmt\_2 = "COMMENT\_2"  
bkDate\_2 = "DATE\_2"

\_output\_ = {}  
\_output\_ = { bkInst =&gt; { bkName =&gt; {"VOL" =&gt; bkVol, "CMT" =&gt; bkCmt, "DATE" =&gt; bkDate} } }  
\_output\_ = { bkInst =&gt; { bkName\_2 =&gt; {"VOL" =&gt; bkVol, "CMT" =&gt; bkCmt\_2, "DATE" =&gt; bkDate\_2} } }  
puts \_output\_

{"INSTANCE\_1"=&gt;{"BACKUP\_2"=&gt;{"VOL"=&gt;"VOL\_1", "CMT"=&gt;"COMMENT\_2", "DATE"=&gt;"DATE\_2"}}}  
 =&gt; nil

I am trying to get it so both sets of data for Backp\_ID is in the data set.  
How can I do this programmatically using the same syntax for each set of data?

Sorry if this is a dumb question.
## [6][The state of ruby blogs cont. - Top topics / words in headlines, top ruby &amp; rails version in headlines and more](https://www.reddit.com/r/ruby/comments/fchutu/the_state_of_ruby_blogs_cont_top_topics_words_in/)
- url: https://www.reddit.com/r/ruby/comments/fchutu/the_state_of_ruby_blogs_cont_top_topics_words_in/
---
Hello,

   I've added a couple of more insights to the [little survey about the state of the ruby feed-iverse](https://github.com/planetruby/planet#the-state-of-ruby-blogs-and-news---36-channels-1464-items) that includes personal blogs, ruby project news and more.

  See the [planet.ini](https://github.com/planetruby/planet/blob/master/planet.ini) for all feeds included in the survey.

Q: What are the top topics / words in headlines?

    Top Words in Headlines 2020  (n=47)
    ---------------------------------
      rails           | ***************** 13
      ruby            | **************** 12
      conferences     | ******* 6
      dry             | ****** 5
      rom             | ***** 4
      ml              | ** 2
      rubymine        | ** 2
    
    Top Words in Headlines 2019  (n=228)
    ---------------------------------
      rails           | ********************** 64
      ruby            | ************* 38
      rubymine        | ***** 16
      activerecord    | ** 8
      bundler         | * 5
      hanami          | * 5
      jruby           | * 5
      activemodel     | * 4
      jekyll          | * 4
      passenger       | * 4
      rgsoc           | * 4
      conferences     |  3
      rubyconf        |  3
      rubygems        |  3
      rubyinstaller   |  3

Q: What are the top ruby versions in headlines?

    Top Ruby Versions in Headlines 2020  (n=93)
    ----------------------------------------------
      ruby 2.7        | ************************************************************ 3
    
    Top Ruby Versions in Headlines 2019  (n=412)
    ----------------------------------------------
      ruby 2.7        | ************************************* 12
      ruby 2.6        | ********* 3
      ruby 3.0        | ****** 2
      ruby 2.4        | *** 1
      ruby 2.5        | *** 1

Q: What are the top rails versions in headlines?

    Top Rails Versions in Headlines 2020  (n=93)
    -----------------------------------------------
      rails 6.0       | ****************************** 1
      rails 6.1       | ****************************** 1
    
    Top Rails Versions in Headlines 2019  (n=412)
    -----------------------------------------------
      rails 6.0       | ******************************************************* 58
      rails 6.1       | ** 3
      rails 5.2       |  1


That's all for now. Comments and ideas on how to improve
the stats are more than welcome. Cheers. Prost.
## [7][Syntax Highlighting with Action Text](https://www.reddit.com/r/ruby/comments/fcb0or/syntax_highlighting_with_action_text/)
- url: https://www.driftingruby.com/episodes/syntax-highlighting-with-action-text?utm_medium=social&amp;utm_campaign=weekly_episode&amp;utm_source=reddit
---

## [8][How a well-planned schema can change your project?](https://www.reddit.com/r/ruby/comments/fccjz9/how_a_wellplanned_schema_can_change_your_project/)
- url: https://blog.graphqleditor.com/well-planned-graphql-schema/
---

## [9][Parsing timestamps in Ruby](https://www.reddit.com/r/ruby/comments/fcekkq/parsing_timestamps_in_ruby/)
- url: https://prathamesh.tech/2020/03/02/parsing-timestamps-in-ruby/
---

## [10][Are (+var+) and #{var} the same thing?](https://www.reddit.com/r/ruby/comments/fc9w4h/are_var_and_var_the_same_thing/)
- url: https://www.reddit.com/r/ruby/comments/fc9w4h/are_var_and_var_the_same_thing/
---
basically the tittle, I'm learning ruby and see a lot of people use (+var+) but isn't #{var} a lot easier? is there situations where you'd need to use a specific one or is it just preference?

thank you :)
