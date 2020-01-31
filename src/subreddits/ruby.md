# ruby
## [1][Has anyone started using ruby 2.7?](https://www.reddit.com/r/ruby/comments/ewjhjb/has_anyone_started_using_ruby_27/)
- url: https://www.reddit.com/r/ruby/comments/ewjhjb/has_anyone_started_using_ruby_27/
---
Recently, I've been wondering what's the adoption of **ruby 2.7**. Has anyone here migrated a production app to 2.7?

I've done some rudimentary tests on SaaSHub (development), and it's a touch slower in most cases. I am aware that there are some GC improvements. It could be the case that switching to 2.7 improves memory consumption on long running processes (puma, sidekiq, etc). If you have any, what's your experience?
## [2][Working with tempfiles](https://www.reddit.com/r/ruby/comments/ewm3mk/working_with_tempfiles/)
- url: https://remimercier.com/working-with-tempfiles/
---

## [3][Let's talk about Roda](https://www.reddit.com/r/ruby/comments/ewh4hf/lets_talk_about_roda/)
- url: https://www.reddit.com/r/ruby/comments/ewh4hf/lets_talk_about_roda/
---
Roda is great, I've deployed a webhook with some backend to view the data and fairly confident it can hold the load. 

it was a breath of fresh air compared to Rails. Would love to hear about what r/ruby done with Roda.
## [4][Help with using FFmpeg and Ruby Rails](https://www.reddit.com/r/ruby/comments/ewj776/help_with_using_ffmpeg_and_ruby_rails/)
- url: https://www.reddit.com/r/ruby/comments/ewj776/help_with_using_ffmpeg_and_ruby_rails/
---
Hey all. I'm new to Ruby / Rails and I am doing a project where I need to upload a video into a Rails web app which would take screen shots, process them put them back together as a Gif or another file format. Ideally, I'd like this to be a web app where users can upload videos, have them be processed with various effects and be available for download or simply to see. 

I feel like this all can be done using FFmpeg via the various gems that use it. My ignorance of coding (I'm a super newbie) is that I don't believe I can deploy web app that uses FFmpeg. To my understanding, the user would need to have FFmpeg installed on their computer in order for my web app to function as it is required for the code to work. Does anyone have any advice or work around or words of wisdom?? 

Thanks in advance!!!
## [5][help. I cannot install ruby on my macbook](https://www.reddit.com/r/ruby/comments/ewiakx/help_i_cannot_install_ruby_on_my_macbook/)
- url: https://www.reddit.com/r/ruby/comments/ewiakx/help_i_cannot_install_ruby_on_my_macbook/
---
can anyone tell me how to install ruby on a mac. I have tried and failed 3 or 4 tutorials. I always get an error and then more errors when i try to fix the previous errors. 

the current tutorial i am failing miserably on is a youtube video found: [https://www.youtube.com/results?search\_query=install+ruby+on+macos](https://www.youtube.com/results?search_query=install+ruby+on+macos)

i seem to have rvm installed when i run rvm -v i get

`rvm 1.29.9 (latest) by Michal Papis, Piotr Kuczynski, Wayne E. Seguin [`[`https://rvm.io`](https://rvm.io)`]`

in the tutorial i am told to enter

`open ~/.bash_profile`

then add some stuff to the bottom. which i have done.

then it asks me to install ruby using 

`rvm install 2.6.3`

from there it all goes to hell. I get this

`Checking requirements for osx.`

`Installing requirements for osx.`

`Updating system - please wait`

`Installing required packages: libtool, coreutils, libyaml, zlib - please wait`

`There were package installation errors, make sure to read the log.`

&amp;#x200B;

`Try \`brew tap --repair\` and make sure \`brew doctor\` looks reasonable.`

&amp;#x200B;

`Check Homebrew requirements` [`https://docs.brew.sh/Installation`](https://docs.brew.sh/Installation)

`Error running 'requirements_osx_brew_libs_install libtool coreutils libyaml zlib',`

`please read /Users/douglas/.rvm/log/1580442167_ruby-2.6.3/package_install_libtool_coreutils_libyaml_zlib.log`

`Requirements installation failed with status: 1.`

so, i check the log like it says and this is what i get

`Error: Permission denied @ apply2files - /usr/local/lib/docker/cli-plugins`

`+requirements_osx_brew_libs_install:4&gt; typeset 'ret=1'`

`+requirements_osx_brew_libs_install:6&gt; requirements_osx_brew_libs_error installation`

`+requirements_osx_brew_libs_error:2&gt; rvm_warn $'There were package installation errors, make sure to read the log.\n\nTry \`brew tap --repair\` and make sure \`brew doctor\` looks reasonable.\n\nCheck Homebrew requirements` [`https://docs.brew.sh/Installation`](https://docs.brew.sh/Installation)`'`

`+rvm_warn:2&gt; rvm_pretty_print stderr`

`+rvm_pretty_print:2&gt; case auto (0|no)`

`+rvm_pretty_print:2&gt; case auto (1|auto)`

`+rvm_pretty_print:8&gt; case xterm-256color (dumb|unknown)`

`+rvm_pretty_print:12&gt; case stderr (stdout)`

`+rvm_pretty_print:12&gt; case stderr (stderr)`

`+rvm_pretty_print:14&gt; [[ -t 2 ]]`

`+rvm_pretty_print:14&gt; return 1`

`+rvm_warn:4&gt; printf %b $'There were package installation errors, make sure to read the log.\n\nTry \`brew tap --repair\` and make sure \`brew doctor\` looks reasonable.\n\nCheck Homebrew requirements` [`https://docs.brew.sh/Installation\\n`](https://docs.brew.sh/Installation\\n)`'`

`+requirements_osx_brew_libs_install:8&gt; return 1`

i am so lost. can anyone help explain this too me???
## [6][Cheat Sheet comparing all the different ways to set attributes in ActiveRecord on Rails 6](https://www.reddit.com/r/ruby/comments/ewb57e/cheat_sheet_comparing_all_the_different_ways_to/)
- url: https://scottbartell.com/2020/01/30/set-attributes-in-active-record-rails-6/
---

## [7][Ruby Concurrency Progress Report for January.](https://www.reddit.com/r/ruby/comments/ew58e5/ruby_concurrency_progress_report_for_january/)
- url: https://www.codeotaku.com/journal/2020-01/ruby-concurrency-progress-report/index
---

## [8][An alternative search bar for RubyGems, made with MeiliSearch](https://www.reddit.com/r/ruby/comments/ew98xa/an_alternative_search_bar_for_rubygems_made_with/)
- url: https://blog.meilisearch.com/meilisearch-finds-rubygems/
---

## [9][New gem for slack sign-ins for Rails](https://www.reddit.com/r/ruby/comments/ew5f5w/new_gem_for_slack_signins_for_rails/)
- url: /r/rails/comments/evmyyu/new_gem_for_slack_signins_for_rails/
---

## [10][Enum implementation code explanation](https://www.reddit.com/r/ruby/comments/ew678y/enum_implementation_code_explanation/)
- url: https://www.reddit.com/r/ruby/comments/ew678y/enum_implementation_code_explanation/
---
Hi guys, I was googling about Enums in Ruby and I came across [this code snippet](https://stackoverflow.com/a/5675566)

```ruby
class MyClass
  MY_ENUM = [MY_VALUE_1 = 'value1', MY_VALUE_2 = 'value2']
end
```

I can't really understand how the values of the `MY_ENUM` array can be accessed like this `MyClass::MY_VALUE_1` Is there something related to constants?

The author also mentioned about the example: `It does some compilation-time checking (in contrast with just using symbols)`  Can we do compile time stuff with Ruby?

Thanks!
