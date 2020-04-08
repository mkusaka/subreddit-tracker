# ruby
## [1][How to build a conference line with Twilio and Ruby](https://www.reddit.com/r/ruby/comments/fwzum1/how_to_build_a_conference_line_with_twilio_and/)
- url: https://www.reddit.com/r/ruby/comments/fwzum1/how_to_build_a_conference_line_with_twilio_and/
---
It's a weird world we live in right now, so much remote work, so many people working from home. There's been an increase in the number of video calls, but what if you just want a voice conference call without choosing an app? You [can set up a conference line easily with Twilio, and this is how to use Ruby and Rails to do so](https://www.twilio.com/blog/build-conference-line-twilio-ruby).
## [2][TIL of a gem that exists just to drive the point home that you should not install gems with sudo.](https://www.reddit.com/r/ruby/comments/fwnt4j/til_of_a_gem_that_exists_just_to_drive_the_point/)
- url: https://github.com/wmorgan/killergem/blob/master/extconf.rb
---

## [3][Easter Egg!](https://www.reddit.com/r/ruby/comments/fx0efn/easter_egg/)
- url: https://www.reddit.com/r/ruby/comments/fx0efn/easter_egg/
---
It's nearly that time of year! Lock-down chocolate! My favourite type.

Anyway, by total coincidence, I just found an easter egg in MRI 2.7!

&gt;!`IRB.send :easter_egg`!&lt;

Enjoy the Eastery goodness - Stay home, Save lives!
## [4][What do matz,k0kubun or other prominent Ruby members think of Crystal?](https://www.reddit.com/r/ruby/comments/fx5jni/what_do_matzk0kubun_or_other_prominent_ruby/)
- url: https://www.reddit.com/r/ruby/comments/fx5jni/what_do_matzk0kubun_or_other_prominent_ruby/
---
To me some form of Ruby on LLVM feels like natural progression. It's not the language that is important. It's what I can do with the language that's important.
## [5][sportdb-source-footballdata 2020 Update - download &amp; import &amp; query 22+ top football leagues from 25 seasons back to 1993/94 - 1000+ clubs, 200000+ matches](https://www.reddit.com/r/ruby/comments/fwjgfs/sportdbsourcefootballdata_2020_update_download/)
- url: https://github.com/yorobot/football.csv/tree/master/sportdb-source-footballdata
---

## [6][Rough and ready replacement for Datakick](https://www.reddit.com/r/ruby/comments/fwkkbh/rough_and_ready_replacement_for_datakick/)
- url: https://www.reddit.com/r/ruby/comments/fwkkbh/rough_and_ready_replacement_for_datakick/
---
Hey...

[Datakick](https://www.datakick.org/) - an open database of barcode / product data - announced it was shutting down at the end of March. I managed to whip up a rough and ready replacement, using data from the original site. If anyone was using the [datakick](https://github.com/ankane/datakick) gem, you should find it will still work - at least in read-only mode - if you retarget the `master` branch version to point to the service running at https://www.brocade.io. 

The code for the service is open source: https://github.com/ferrisoxide/brocade.io

Eventually we'll roll out a more specific-use gem and add editing support back in. 

Right now I'm not looking a critique for the Brocade code. I know it's awful - it was written in a rush to get something up and running before Datakick disappeared. Having said that, if there's anyone else working on barcode / product data projects I'd be keen to connect and share ideas.
## [7][Salim Semaoune: Scale Background Queues with Asynchronous I/O](https://www.reddit.com/r/ruby/comments/fwedhf/salim_semaoune_scale_background_queues_with/)
- url: https://www.youtube.com/watch?v=w5ABH3mZAxU
---

## [8][Yet another active form](https://www.reddit.com/r/ruby/comments/fwk5ae/yet_another_active_form/)
- url: /r/rails/comments/fwk4ip/yet_another_active_form/
---

## [9][XSD to Ruby classes](https://www.reddit.com/r/ruby/comments/fwh4ln/xsd_to_ruby_classes/)
- url: https://www.reddit.com/r/ruby/comments/fwh4ln/xsd_to_ruby_classes/
---
Been trying to convert a stack of .xsd schemas to Ruby classes. There appears to have been some work in this area, way back when XML was king (\~10 years ago) but I can't find any gems or code that has been maintained over the years. I've searched github for ruby+xsd, tried installing and running what I found - slowly going insane as everything seems to be borked/undocumented/unsupported.

I note that there used to be support for XSD directly in the Ruby standard lib, circa v1.8.7, but since removed. Am I just barking up the wrong tree here? Should I just try to do this by hand, and not look for a pre-existing solution?
## [10][https://github.com/brnnkt/local_images](https://www.reddit.com/r/ruby/comments/fwjzha/httpsgithubcombrnnktlocal_images/)
- url: https://www.reddit.com/r/ruby/comments/fwjzha/httpsgithubcombrnnktlocal_images/
---
My First rubygem. Saved some time for me on routine action of downloading images for seeds.rb (RoR) (I don't like remote sources of images, because of internet connection can be slow or even off)
