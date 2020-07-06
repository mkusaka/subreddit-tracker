# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/hm6yzb/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/hm6yzb/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - July 06, 2020](https://www.reddit.com/r/androiddev/comments/hm5itj/weekly_questions_thread_july_06_2020/)
- url: https://www.reddit.com/r/androiddev/comments/hm5itj/weekly_questions_thread_july_06_2020/
---
This thread is for simple questions that don't warrant their own thread (although we suggest checking the sidebar, [the wiki](http://www.reddit.com/r/androiddev/wiki/), [our Discord](https://discord.gg/D2cNrqX), or [Stack Overflow](http://stackoverflow.com) before posting). Examples of questions:

* How do I pass data between my Activities?
* Does anyone have a link to the source for the AOSP messaging app?
* Is it possible to programmatically change the color of the status bar without targeting API 21?

**Important: Downvotes are strongly discouraged in this thread. Sorting by new is strongly encouraged.**

Large code snippets don't read well on reddit and take up a lot of space, so please don't paste them in your comments. Consider linking [Gists](https://gist.github.com) instead.

Have a question about the subreddit or otherwise for /r/androiddev mods? [We welcome your mod mail!](http://www.reddit.com/message/compose?to=%2Fr%2Fandroiddev)

Also, please don't link to Play Store pages or ask for feedback on this thread. Save those for the App Feedback threads we host on Saturdays.

Looking for all the Questions threads? Want an easy way to locate this week's thread? Click [this link](https://www.reddit.com/r/androiddev/search?q=title%3A%22questions+thread%22+author%3A%22AutoModerator%22&amp;restrict_sr=on&amp;sort=new&amp;t=all)!
## [3][Optimize the build speed for your Android project](https://www.reddit.com/r/androiddev/comments/hm4mzz/optimize_the_build_speed_for_your_android_project/)
- url: https://www.crazylegend.dev/2020/07/optimize-build-speeds-for-your-android.html
---

## [4][Why are simple things sometimes so complex [mini rant]](https://www.reddit.com/r/androiddev/comments/hm6vya/why_are_simple_things_sometimes_so_complex_mini/)
- url: https://www.reddit.com/r/androiddev/comments/hm6vya/why_are_simple_things_sometimes_so_complex_mini/
---
So I have a spinner (drop-down select, the naming of this thing is another rant) and I have a blue background and want to change the text colour to white, sounds simple, right?

to do that I have to create TWO layout resources (one for selected item and one for dropdown) and in code declare an adapter which needs the item array (cause the reference in XML is not enough), context, and the resource for the style of the selected item. then on a separate line I have to declare a separate resource for the items in the drop-down menu.

I JUST WANTED TO CHANGE THE TEXT COLOUR!!!!! why do does this need two extra XML files and 3 lines of code!? it sometimes feels like a lot of the android API is specifically made to have as many files in a project as possible
## [5][How to design creative apps? [practitioners' inputs!]](https://www.reddit.com/r/androiddev/comments/hm3xus/how_to_design_creative_apps_practitioners_inputs/)
- url: https://www.reddit.com/r/androiddev/comments/hm3xus/how_to_design_creative_apps_practitioners_inputs/
---
Hi everyone!

For  our research on fostering creativity in app design, we are looking for mobile app designers (18+ y.o.), with any level of experience, who are willing to participate in our study via short questionnaire (10-15 minutes).

You can participate here: [https://survey.uu.nl/jfe/form/SV\_eMdNbBVZUabmohD](https://survey.uu.nl/jfe/form/SV_eMdNbBVZUabmohD)

Your  help will be much appreciated and it would help us a lot with our  research! We will be happy to keep your posted on the results and to share our tools. Also, feel free to comment on this post!

Thanks!

Fabiano Dalpiaz and Ilse van 't Hul  
Utrecht University, NL
## [6][Building a multi-module project](https://www.reddit.com/r/androiddev/comments/hm35x4/building_a_multimodule_project/)
- url: https://www.reddit.com/r/androiddev/comments/hm35x4/building_a_multimodule_project/
---
Hi guys, sorry to be posting this here but couldn't find an answer for this anywhere:

What I have:

1. a multi-module project that works fine.
2. Project is stored in Git and each sub-module has it's own repository.
3. No problems with references or anything like that, again - functionally it works just fine.

The "problem":

My team and I work on the project by opening all the modules in the same Android Studio window, we see all the module in the same tree and we can open each one just fine (noted here since that might be related, IDK). If we change a file/class/asset/something in one of the submodules, and then want to run the application with the new change, we have to "rebuild" instead of "build" which takes extra time and doesn't happen automatically like regular "build" which occurs before each run.

If we "build" the specific sub-module that was changed, the change isn't recognised when we run the app.

Is that normal? did we miss something in the configuration? Is there a way to make the project build all the submodules with the "build" command instead of the "rebuild"? Is it not supposed to? Any input would be appreciated :)
## [7][How to do a splashscreen the right way?](https://www.reddit.com/r/androiddev/comments/hlq6fu/how_to_do_a_splashscreen_the_right_way/)
- url: https://www.reddit.com/r/androiddev/comments/hlq6fu/how_to_do_a_splashscreen_the_right_way/
---
I've seen countless videos where people are saying to make a splash screen occur for a set amount of time. I know this is wrong, yet I can't find any resources to do a splash screen the right way, any help?
## [8][Payment declined for ~ 70% of in-app purchases](https://www.reddit.com/r/androiddev/comments/hm5vnt/payment_declined_for_70_of_inapp_purchases/)
- url: https://www.reddit.com/r/androiddev/comments/hm5vnt/payment_declined_for_70_of_inapp_purchases/
---
I have an Android app with about 300k installs with in-app purchases (one time and recurring).

I have one issue though, about 60-70% of the orders are being marked as "Payment declined".

When I click on the declined orders I see "There was an issue charging the customer’s payment method" and "This order could not be completed and was canceled".

They're not coming from a specific country/user, and 70% decline rate is very high.

I'd appreciate any help regarding this issue.
## [9][How to best debug bluetooth connectivity?](https://www.reddit.com/r/androiddev/comments/hm4pdr/how_to_best_debug_bluetooth_connectivity/)
- url: https://www.reddit.com/r/androiddev/comments/hm4pdr/how_to_best_debug_bluetooth_connectivity/
---
Hey guys, I started a new job as android dev, and the app Im working on utilises BLE and iBeacon to connect to a small device. It usually works pretty well, but some users are reporting connectivity issues, constant disconnects and reconnects, cannot connect at all, etc.

I cannot replicate these issues on my own test devices, even if they are the same models being reported. My experience with BLE is pretty low, so I'd love some resources on how best to work with it on android. Any suggestions?
## [10][Course for experienced developers](https://www.reddit.com/r/androiddev/comments/hm7mro/course_for_experienced_developers/)
- url: https://www.reddit.com/r/androiddev/comments/hm7mro/course_for_experienced_developers/
---
Hi, I am looking for an Android with Kotlin course for experienced developers. I have been in the field for 10 years and I am not wasting 60 hours listening to how to declare a variable. Do you know of any good courses?  Couldn't find any by just googling.
## [11][Need help with FRP on a S8+](https://www.reddit.com/r/androiddev/comments/hm72xx/need_help_with_frp_on_a_s8/)
- url: https://www.reddit.com/r/androiddev/comments/hm72xx/need_help_with_frp_on_a_s8/
---
 

So my son has reset his Samsung S8+ and now has been met with the FRP and can't remember his details to log in. we have tried flashing in odin but the FRP is still there. I see there is a method to us a combination file but can't find the combination file for my phone as the only ones I find are for the U1 - U8 models and never for my S8DTC6. So if anyone could point me in the right direction it would be a massive help.

Phone - Samsung S8+

Model Number - SM-G955FZKAVOD

Firmware - G955FXXS8DTC6

Security Patch - 01/04/2020
## [12][The Path To Success As A Software Developer](https://www.reddit.com/r/androiddev/comments/hm6k9i/the_path_to_success_as_a_software_developer/)
- url: https://medium.com/vinsloev-academy/the-path-to-success-as-a-software-developer-61f01aad89f9?source=friends_link&amp;sk=eba99334d4d473d82e17643dcb52032e
---

