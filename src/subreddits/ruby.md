# ruby
## [1][A silly thing you can do with the Ruby parser](https://www.reddit.com/r/ruby/comments/hyog8p/a_silly_thing_you_can_do_with_the_ruby_parser/)
- url: https://penelope.zone/2019/12/22/a-silly-thing-you-can-do-with-the-ruby-parser.html
---

## [2][I am taking a Thinkpad T480s from my workplace (migration from a MacBook setup) Do you think I can work nicely from a Windows 10 directly?](https://www.reddit.com/r/ruby/comments/hypim3/i_am_taking_a_thinkpad_t480s_from_my_workplace/)
- url: https://www.reddit.com/r/ruby/comments/hypim3/i_am_taking_a_thinkpad_t480s_from_my_workplace/
---
Hey channel,  
After \~10y working with Ruby, always using ubuntu and osx, I am now planning to work/code from the windows 10 directly.   
What do you think about this approach?  


I could have a Virtualbox with ubuntu for code, but not sure if the T480s i5 processor will be enough to have a nice experience on it.  


I heard that windows now have some new terminal stuff.. Anyone with similar situation that wanna share some knowledge here?
## [3][Fun Facts about Ruby #12: unusual ways to call a method](https://www.reddit.com/r/ruby/comments/hyn2ng/fun_facts_about_ruby_12_unusual_ways_to_call_a/)
- url: https://i.redd.it/put2urv9hcd51.png
---

## [4][Optimally distribute and execute RSpec suites among parallel workers; for faster CI builds](https://www.reddit.com/r/ruby/comments/hy4dvk/optimally_distribute_and_execute_rspec_suites/)
- url: https://github.com/skroutz/rspecq
---

## [5][MRuby faster than both YARV and MJIT](https://www.reddit.com/r/ruby/comments/hyd73z/mruby_faster_than_both_yarv_and_mjit/)
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
## [6][Is there no good way to upload multiple images with Ruby on Rails?](https://www.reddit.com/r/ruby/comments/hyhi4z/is_there_no_good_way_to_upload_multiple_images/)
- url: https://www.reddit.com/r/ruby/comments/hyhi4z/is_there_no_good_way_to_upload_multiple_images/
---
TLDR: Uploading multiple pictures to the cloud is getting annoying, anyone know a better way than Shrinerb?

I've been using Carrierwave on my Heroku-hosted picture gallery app for a bit, and started to notice that large images fail. I found out it was because my server times out after 30 seconds, so I decided I should upload directly to Google Cloud Services. ActiveStorage doesn't let you process images, which would cause preposterous load times, so that's out. [This gem](https://github.com/dwilkie/carrierwave_direct) lets you direct upload with Carrierwave, but only one image at a time, so that's out. The only viable solution I've found seems to be Shrinerb, but I've found direct upload to Google Cloud to be a bottomless rabbit hole, with different documentations pointing to each other but not detailing how to integrate into each other. Is there really no easy way to upload multiple images to the cloud or should I keep going down this rabbit hole?
## [7][Make API payloads generic - API only ruby on rails course (chapter 6)](https://www.reddit.com/r/ruby/comments/hxm279/make_api_payloads_generic_api_only_ruby_on_rails/)
- url: https://duetcode.io/rails-api-only-course/make-api-payloads-generic
---

## [8][Starting a monthly newsletter for cool stuff I encounter on Rails, React and Graphql.](https://www.reddit.com/r/ruby/comments/hxpads/starting_a_monthly_newsletter_for_cool_stuff_i/)
- url: https://www.reddit.com/r/ruby/comments/hxpads/starting_a_monthly_newsletter_for_cool_stuff_i/
---
Starting a monthly newsletter on Ruby on Rails, React and GraphQL developers to read about some cool stuff happening in the community.

In a month we come across a lot cool stuff happening in the community, learn new things and contribute to the open-source world and newsletters is the best way to share some of those finding/blogs/videos.

This newsletter would consist of the following things:

* Dev Joke
* My Blogs.
* Reading/watching list.
* Open-source libraries/contributions.
* Rails changelog.
* Watching/Reading/Reviews.

If you'd like to read my monthly newsletter. Please do subscribe here [https://buttondown.email/abhaynikam](https://buttondown.email/abhaynikam)
## [9][Failing to Deploy RoR at DO Droplet](https://www.reddit.com/r/ruby/comments/hxvixo/failing_to_deploy_ror_at_do_droplet/)
- url: /r/rubyonrails/comments/hxtwhv/failing_from_reach_index_page_on_ror_deployed_at/
---

## [10][Installing Ruby on VSCode Please Help!](https://www.reddit.com/r/ruby/comments/hxazpb/installing_ruby_on_vscode_please_help/)
- url: https://www.reddit.com/r/ruby/comments/hxazpb/installing_ruby_on_vscode_please_help/
---
Hey guys I've been trying to figure out how to use ruby on VSCode for 2 hours on windows and I have no idea. I have Ruby, ruby extension pack, and vscode ruby extensions installed. I made a test.rb file and wrote some code but theres no run option when i right click. If I click run - start debugging on the top all I dont see any code executed.  It says Node.js (preview) in the top right. I dont know what that means?

Please help! What do I do?

&amp;#x200B;

https://preview.redd.it/mr6nxvzvnvc51.png?width=1920&amp;format=png&amp;auto=webp&amp;s=f349e3f3bb85d4f3201e4fb3354a25425071eed6
