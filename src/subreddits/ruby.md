# ruby
## [1][Exploring Hanami 2.0 application template - part 2](https://www.reddit.com/r/ruby/comments/hyxg2r/exploring_hanami_20_application_template_part_2/)
- url: https://youtu.be/H1eKIO39ob8
---

## [2][Ractor parallel by ko1 · Pull Request #3365](https://www.reddit.com/r/ruby/comments/hyvby9/ractor_parallel_by_ko1_pull_request_3365/)
- url: https://github.com/ruby/ruby/pull/3365
---

## [3][A silly thing you can do with the Ruby parser](https://www.reddit.com/r/ruby/comments/hyog8p/a_silly_thing_you_can_do_with_the_ruby_parser/)
- url: https://penelope.zone/2019/12/22/a-silly-thing-you-can-do-with-the-ruby-parser.html
---

## [4][How to Dockerize a Sinatra application](https://www.reddit.com/r/ruby/comments/hytyt5/how_to_dockerize_a_sinatra_application/)
- url: https://www.reddit.com/r/ruby/comments/hytyt5/how_to_dockerize_a_sinatra_application/
---
Recently I've been trying to figure out how to Dockerize a Rails application. I'm finding it really hard and the documentation available online doesn't seem that good. So I decided to make things a little easier and try to Dockerize a Sinatra application first.

Here's how I did it. It's a step-by-step tutorial which also includes a video and Git repo.

[https://www.codewithjason.com/dockerize-sinatra-application/](https://www.codewithjason.com/dockerize-sinatra-application/)
## [5][I am taking a Thinkpad T480s from my workplace (migration from a MacBook setup) Do you think I can work nicely from a Windows 10 directly?](https://www.reddit.com/r/ruby/comments/hypim3/i_am_taking_a_thinkpad_t480s_from_my_workplace/)
- url: https://www.reddit.com/r/ruby/comments/hypim3/i_am_taking_a_thinkpad_t480s_from_my_workplace/
---
Hey channel,  
After \~10y working with Ruby, always using ubuntu and osx, I am now planning to work/code from the windows 10 directly.   
What do you think about this approach?  


I could have a Virtualbox with ubuntu for code, but not sure if the T480s i5 processor will be enough to have a nice experience on it.  


I heard that windows now have some new terminal stuff.. Anyone with similar situation that wanna share some knowledge here?
## [6][SerpApi vs Zenserp vs Serpstack vs SerpWow vs ScaleSerp Benchmarks](https://www.reddit.com/r/ruby/comments/hz2r8n/serpapi_vs_zenserp_vs_serpstack_vs_serpwow_vs/)
- url: https://medium.com/@hartator/serpapi-vs-zenserp-vs-serpstack-vs-serpwow-vs-scaleserp-benchmarks-6ef5fd467f7d?sk=153832e48264c08bb90aa73f5941de98
---

## [7][Fun Facts about Ruby #12: unusual ways to call a method](https://www.reddit.com/r/ruby/comments/hyn2ng/fun_facts_about_ruby_12_unusual_ways_to_call_a/)
- url: https://i.redd.it/put2urv9hcd51.png
---

## [8][Optimally distribute and execute RSpec suites among parallel workers; for faster CI builds](https://www.reddit.com/r/ruby/comments/hy4dvk/optimally_distribute_and_execute_rspec_suites/)
- url: https://github.com/skroutz/rspecq
---

## [9][MRuby faster than both YARV and MJIT](https://www.reddit.com/r/ruby/comments/hyd73z/mruby_faster_than_both_yarv_and_mjit/)
- url: https://www.reddit.com/r/ruby/comments/hyd73z/mruby_faster_than_both_yarv_and_mjit/
---
I am creating a little project with some benchmark across different languages, and I found MRuby a little bit faster than well, both YARV and MJIT, here's the code of my Mandelbrot algorithm: [https://termbin.com/i91bj](https://termbin.com/i91bj)

Ran this with 8000 as first parameter and and got this, from my "benchmarking algorithm", but for sure you can use \`time\` for measuring it:

`YARV:`

`862.2862342359986`

&amp;#x200B;

`MJIT:`

`847.2985952870004`

&amp;#x200B;

`MRuby:`

`644.4368903609975`

&amp;#x200B;

`MRubyAOT:`

`642.4403260649997`

&amp;#x200B;

MRuby isn't supposed to be just like MRI but smaller and embeddable, and MRI, the core implementation bigger and faster?

EDIT: the first line of the Mandelbrot set, is missing when I ran it with MRuby, naturally.
## [10][Is there no good way to upload multiple images with Ruby on Rails?](https://www.reddit.com/r/ruby/comments/hyhi4z/is_there_no_good_way_to_upload_multiple_images/)
- url: https://www.reddit.com/r/ruby/comments/hyhi4z/is_there_no_good_way_to_upload_multiple_images/
---
TLDR: Uploading multiple pictures to the cloud is getting annoying, anyone know a better way than Shrinerb?

I've been using Carrierwave on my Heroku-hosted picture gallery app for a bit, and started to notice that large images fail. I found out it was because my server times out after 30 seconds, so I decided I should upload directly to Google Cloud Services. ActiveStorage doesn't let you process images, which would cause preposterous load times, so that's out. [This gem](https://github.com/dwilkie/carrierwave_direct) lets you direct upload with Carrierwave, but only one image at a time, so that's out. The only viable solution I've found seems to be Shrinerb, but I've found direct upload to Google Cloud to be a bottomless rabbit hole, with different documentations pointing to each other but not detailing how to integrate into each other. Is there really no easy way to upload multiple images to the cloud or should I keep going down this rabbit hole?
