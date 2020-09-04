# ruby
## [1][Ractor landed in Ruby master. That's it. That's the news \o/](https://www.reddit.com/r/ruby/comments/imbm3t/ractor_landed_in_ruby_master_thats_it_thats_the/)
- url: https://github.com/ruby/ruby/commit/79df14c04b452411b9d17e26a398e491bca1a811
---

## [2][RubyKaigi Takeout 2020 - ractor presentation by koichi sasada](https://www.reddit.com/r/ruby/comments/imcro1/rubykaigi_takeout_2020_ractor_presentation_by/)
- url: https://youtu.be/2ZcdiVSERuY?t=473
---

## [3][Proportional dispatching using Ruby](https://www.reddit.com/r/ruby/comments/imd2w3/proportional_dispatching_using_ruby/)
- url: https://medium.com/rubycademy/proportional-dispatching-using-ruby-378afbdeb32d
---

## [4][Results: 2020 Ruby on Rails Community Survey](https://www.reddit.com/r/ruby/comments/im3cvo/results_2020_ruby_on_rails_community_survey/)
- url: https://rails-hosting.com/2020/
---

## [5]['require' cannot load such file (loaderror) ... help! :)](https://www.reddit.com/r/ruby/comments/imcujt/require_cannot_load_such_file_loaderror_help/)
- url: https://www.reddit.com/r/ruby/comments/imcujt/require_cannot_load_such_file_loaderror_help/
---
Just got a new laptop and I copied my scripts over and installed Ruby and the gems I want to use. When I try to run a script I get this error. I have been searching Google all night and find similar problems others are having but I can't seem to get my problem fixed. I didn't have this issue on the other laptop.  I am on Windows 10. The only thing that I know I did differently on the new computer was also to install rails. I'm at my wit's end, any help would be appreciated. 

&amp;#x200B;

    RubyGems Environment:
      - RUBYGEMS VERSION: 3.1.4
      - RUBY VERSION: 2.6.6 (2020-03-31 patchlevel 146) [x64-mingw32]
      - INSTALLATION DIRECTORY: C:/Ruby26-x64/lib/ruby/gems/2.6.0
      - USER INSTALLATION DIRECTORY: C:/Users/gonzo/.gem/ruby/2.6.0
      - RUBY EXECUTABLE: C:/Ruby26-x64/bin/ruby.exe
      - GIT EXECUTABLE:
      - EXECUTABLE DIRECTORY: C:/Ruby26-x64/bin
      - SPEC CACHE DIRECTORY: C:/Users/gonzo/.gem/specs
      - SYSTEM CONFIGURATION DIRECTORY: C:/ProgramData
      - RUBYGEMS PLATFORMS:
        - ruby
        - x64-mingw32
      - GEM PATHS:
         - C:/Ruby26-x64/lib/ruby/gems/2.6.0
         - C:/Users/gonzo/.gem/ruby/2.6.0
      - GEM CONFIGURATION:
         - :update_sources =&gt; true
         - :verbose =&gt; true
         - :backtrace =&gt; false
         - :bulk_threshold =&gt; 1000
      - REMOTE SOURCES:
         - https://rubygems.org/
      - SHELL PATH:
         - C:\windows\system32
         - C:\windows
         - C:\windows\System32\Wbem
         - C:\windows\System32\WindowsPowerShell\v1.0\
         - C:\windows\System32\OpenSSH\
         - C:\Ruby26-x64\bin
         - C:\Users\gonzo\AppData\Local\Microsoft\WindowsApps
         - C:\Users\gonzo\AppData\Local\GitHubDesktop\bin
         - C:\Users\gonzo\AppData\Local\atom\bin
    
    

Gem list:

    *** LOCAL GEMS ***
    
    actioncable (6.0.3.2)
    actionmailbox (6.0.3.2)
    actionmailer (6.0.3.2)
    actionpack (6.0.3.2)
    actiontext (6.0.3.2)
    actionview (6.0.3.2)
    activejob (6.0.3.2)
    activemodel (6.0.3.2)
    activerecord (6.0.3.2)
    activestorage (6.0.3.2)
    activesupport (6.0.3.2)
    bigdecimal (default: 1.4.1)
    builder (3.2.4)
    bundler (2.1.4)
    childprocess (3.0.0)
    cmath (default: 1.0.0)
    concurrent-ruby (1.1.7)
    crass (1.0.6)
    csv (default: 3.0.9)
    date (3.0.1, default: 2.0.0)
    dbm (default: 1.0.0)
    did_you_mean (1.3.0)
    diff-lcs (1.4.4)
    e2mmap (default: 0.1.0)
    erubi (1.9.0)
    etc (default: 1.0.1)
    fcntl (default: 1.0.0)
    fiddle (default: 1.0.0)
    fileutils (default: 1.1.0)
    forwardable (default: 1.2.0)
    gdbm (default: 2.0.0)
    globalid (0.4.2)
    i18n (1.8.5)
    io-console (default: 0.4.7)
    ipaddr (default: 1.2.2)
    irb (default: 1.0.0)
    json (default: 2.1.0)
    logger (default: 1.3.0)
    loofah (2.7.0)
    mail (2.7.1)
    marcel (0.3.3)
    matrix (default: 0.1.0)
    method_source (1.0.0)
    mimemagic (0.3.5)
    mini_mime (1.0.2)
    mini_portile2 (2.4.0)
    minitest (5.11.3)
    mutex_m (default: 0.1.0)
    net-telnet (0.2.0)
    nio4r (2.5.2)
    nokogiri (1.10.10 x64-mingw32)
    openssl (default: 2.1.2)
    ostruct (default: 0.1.0)
    power_assert (1.1.3)
    prime (default: 0.1.0)
    psych (default: 3.1.0)
    rack (2.2.3, 2.0.9)
    rack-test (1.1.0)
    rails (6.0.3.2)
    rails-dom-testing (2.0.3)
    rails-html-sanitizer (1.3.0)
    railties (6.0.3.2)
    rake (12.3.3)
    rdoc (default: 6.1.2)
    regexp_parser (1.7.1)
    rexml (default: 3.1.9)
    rspec (3.9.0)
    rspec-core (3.9.2)
    rspec-expectations (3.9.2)
    rspec-mocks (3.9.1)
    rspec-support (3.9.3)
    rss (default: 0.2.7)
    rubyzip (2.3.0)
    scanf (default: 1.0.0)
    sdbm (default: 1.0.0)
    selenium-webdriver (3.142.7)
    shell (default: 0.7)
    sprockets (4.0.2)
    sprockets-rails (3.2.1)
    stringio (default: 0.0.2)
    strscan (default: 1.0.0)
    sync (default: 0.5.0)
    test-unit (3.2.9)
    thor (1.0.1)
    thread_safe (0.3.6)
    thwait (default: 0.1.0)
    tracer (default: 0.1.0)
    tzinfo (1.2.7)
    watir (6.17.0)
    webdrivers (4.4.1)
    webrick (default: 1.4.2)
    websocket-driver (0.7.3)
    websocket-extensions (0.1.5)
    xmlrpc (0.3.0)
    zeitwerk (2.4.0)
    zlib (default: 1.0.0)

&amp;#x200B;

Here's the error:

    C:\Users\gonzo\OneDrive\Desktop\Ruby&gt;ruby test.rb
    Traceback (most recent call last):
            2: from test.rb:5:in `&lt;main&gt;'
            1: from C:/Ruby26-x64/lib/ruby/site_ruby/2.6.0/rubygems/core_ext/kernel_require.rb:92:in `require'
    C:/Ruby26-x64/lib/ruby/site_ruby/2.6.0/rubygems/core_ext/kernel_require.rb:92:in `require': cannot load such file -- watir (LoadError)
            18: from test.rb:5:in `&lt;main&gt;'
            17: from C:/Ruby26-x64/lib/ruby/site_ruby/2.6.0/rubygems/core_ext/kernel_require.rb:156:in `require'
            16: from C:/Ruby26-x64/lib/ruby/site_ruby/2.6.0/rubygems/core_ext/kernel_require.rb:168:in `rescue in require'
            15: from C:/Ruby26-x64/lib/ruby/site_ruby/2.6.0/rubygems/core_ext/kernel_require.rb:168:in `require'
            14: from C:/Ruby26-x64/lib/ruby/gems/2.6.0/gems/watir-6.17.0/lib/watir.rb:1:in `&lt;top (required)&gt;'
            13: from C:/Ruby26-x64/lib/ruby/site_ruby/2.6.0/rubygems/core_ext/kernel_require.rb:72:in `require'
            12: from C:/Ruby26-x64/lib/ruby/site_ruby/2.6.0/rubygems/core_ext/kernel_require.rb:72:in `require'
            11: from C:/Ruby26-x64/lib/ruby/gems/2.6.0/gems/selenium-webdriver-3.142.7/lib/selenium-webdriver.rb:20:in `&lt;top (required)&gt;'
            10: from C:/Ruby26-x64/lib/ruby/site_ruby/2.6.0/rubygems/core_ext/kernel_require.rb:72:in `require'
             9: from C:/Ruby26-x64/lib/ruby/site_ruby/2.6.0/rubygems/core_ext/kernel_require.rb:72:in `require'
             8: from C:/Ruby26-x64/lib/ruby/gems/2.6.0/gems/selenium-webdriver-3.142.7/lib/selenium/webdriver.rb:20:in `&lt;top (required)&gt;'
             7: from C:/Ruby26-x64/lib/ruby/site_ruby/2.6.0/rubygems/core_ext/kernel_require.rb:72:in `require'
             6: from C:/Ruby26-x64/lib/ruby/site_ruby/2.6.0/rubygems/core_ext/kernel_require.rb:72:in `require'
             5: from C:/Ruby26-x64/lib/ruby/gems/2.6.0/gems/childprocess-3.0.0/lib/childprocess.rb:209:in `&lt;top (required)&gt;'
             4: from C:/Ruby26-x64/lib/ruby/site_ruby/2.6.0/rubygems/core_ext/kernel_require.rb:72:in `require'
             3: from C:/Ruby26-x64/lib/ruby/site_ruby/2.6.0/rubygems/core_ext/kernel_require.rb:72:in `require'
             2: from C:/Ruby26-x64/lib/ruby/gems/2.6.0/gems/childprocess-3.0.0/lib/childprocess/windows.rb:4:in `&lt;top (required)&gt;'
             1: from C:/Ruby26-x64/lib/ruby/site_ruby/2.6.0/rubygems/core_ext/kernel_require.rb:92:in `require'
    C:/Ruby26-x64/lib/ruby/site_ruby/2.6.0/rubygems/core_ext/kernel_require.rb:92:in `require': cannot load such file -- ffi (LoadError)
            17: from test.rb:5:in `&lt;main&gt;'
            16: from C:/Ruby26-x64/lib/ruby/site_ruby/2.6.0/rubygems/core_ext/kernel_require.rb:156:in `require'
            15: from C:/Ruby26-x64/lib/ruby/site_ruby/2.6.0/rubygems/core_ext/kernel_require.rb:168:in `rescue in require'
            14: from C:/Ruby26-x64/lib/ruby/site_ruby/2.6.0/rubygems/core_ext/kernel_require.rb:168:in `require'
            13: from C:/Ruby26-x64/lib/ruby/gems/2.6.0/gems/watir-6.17.0/lib/watir.rb:1:in `&lt;top (required)&gt;'
            12: from C:/Ruby26-x64/lib/ruby/site_ruby/2.6.0/rubygems/core_ext/kernel_require.rb:72:in `require'
            11: from C:/Ruby26-x64/lib/ruby/site_ruby/2.6.0/rubygems/core_ext/kernel_require.rb:72:in `require'
            10: from C:/Ruby26-x64/lib/ruby/gems/2.6.0/gems/selenium-webdriver-3.142.7/lib/selenium-webdriver.rb:20:in `&lt;top (required)&gt;'
             9: from C:/Ruby26-x64/lib/ruby/site_ruby/2.6.0/rubygems/core_ext/kernel_require.rb:72:in `require'
             8: from C:/Ruby26-x64/lib/ruby/site_ruby/2.6.0/rubygems/core_ext/kernel_require.rb:72:in `require'
             7: from C:/Ruby26-x64/lib/ruby/gems/2.6.0/gems/selenium-webdriver-3.142.7/lib/selenium/webdriver.rb:20:in `&lt;top (required)&gt;'
             6: from C:/Ruby26-x64/lib/ruby/site_ruby/2.6.0/rubygems/core_ext/kernel_require.rb:72:in `require'
             5: from C:/Ruby26-x64/lib/ruby/site_ruby/2.6.0/rubygems/core_ext/kernel_require.rb:72:in `require'
             4: from C:/Ruby26-x64/lib/ruby/gems/2.6.0/gems/childprocess-3.0.0/lib/childprocess.rb:209:in `&lt;top (required)&gt;'
             3: from C:/Ruby26-x64/lib/ruby/site_ruby/2.6.0/rubygems/core_ext/kernel_require.rb:72:in `require'
             2: from C:/Ruby26-x64/lib/ruby/site_ruby/2.6.0/rubygems/core_ext/kernel_require.rb:72:in `require'
             1: from C:/Ruby26-x64/lib/ruby/gems/2.6.0/gems/childprocess-3.0.0/lib/childprocess/windows.rb:3:in `&lt;top (required)&gt;'
    C:/Ruby26-x64/lib/ruby/gems/2.6.0/gems/childprocess-3.0.0/lib/childprocess/windows.rb:6:in `rescue in &lt;top (required)&gt;': FFI is a required pre-requisite for Windows or posix_spawn support in the ChildProcess gem. Ensure the `ffi` gem is installed. If you believe this is an error, please file a bug at http://github.com/enkessler/childprocess/issues (ChildProcess::MissingFFIError)
## [6][The latest tty-markdown gem adds support for definition lists, footnotes, XML comments and common HTML elements.](https://www.reddit.com/r/ruby/comments/im544w/the_latest_ttymarkdown_gem_adds_support_for/)
- url: https://github.com/piotrmurach/tty-markdown
---

## [7][Anti-IF framework - if/else based on type](https://www.reddit.com/r/ruby/comments/ilt25m/antiif_framework_ifelse_based_on_type/)
- url: https://blog.arkency.com/anti-if-framework---if-slash-else-based-on-type/
---

## [8][Using methods with super when there is no parent class other than active record](https://www.reddit.com/r/ruby/comments/ilocdh/using_methods_with_super_when_there_is_no_parent/)
- url: https://www.reddit.com/r/ruby/comments/ilocdh/using_methods_with_super_when_there_is_no_parent/
---
Hi. I came across this piece of code:

`class XYZ &lt; ActiveRecord::Base`

  `def method1`

`print("hello world")`

`super`

  `end`

My question is this. Since there is no defined parent class other than active record, what does this super do ? is it a bug in the system or does it actually do something. This code is a few years old.

Edit: the answer is that there was a module included in the model which had 'method1' defined.
## [9][[Ruby, Rails, Sorbet] Cleo and Robb clean up a Ruby on Rails codebase with Sorbet](https://www.reddit.com/r/ruby/comments/ilk9f4/ruby_rails_sorbet_cleo_and_robb_clean_up_a_ruby/)
- url: https://www.reddit.com/r/ruby/comments/ilk9f4/ruby_rails_sorbet_cleo_and_robb_clean_up_a_ruby/
---
[https://www.twitch.tv/lawiscode](https://www.twitch.tv/lawiscode)
## [10][Anyone tried making a NonEmptyArray class?](https://www.reddit.com/r/ruby/comments/ilx1id/anyone_tried_making_a_nonemptyarray_class/)
- url: https://www.reddit.com/r/ruby/comments/ilx1id/anyone_tried_making_a_nonemptyarray_class/
---
I learned about this idea from Haskell. I realized it'd help in some of my code where I **know** that the array has at least one item. It'd also help make my code safer by using a NonEmptyArrays instead of Arrays because I wouldn't need if/then checks and exception raising.
