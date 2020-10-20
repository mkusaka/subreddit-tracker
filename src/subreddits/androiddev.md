# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/je0ybe/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/je0ybe/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - October 19, 2020](https://www.reddit.com/r/androiddev/comments/jdziez/weekly_questions_thread_october_19_2020/)
- url: https://www.reddit.com/r/androiddev/comments/jdziez/weekly_questions_thread_october_19_2020/
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
## [3][AGP 4.1 breaking change for libraries: did it hit you?](https://www.reddit.com/r/androiddev/comments/jekj8e/agp_41_breaking_change_for_libraries_did_it_hit/)
- url: https://www.reddit.com/r/androiddev/comments/jekj8e/agp_41_breaking_change_for_libraries_did_it_hit/
---
The recent release of Android Studio 4.1 and the respective Gradle Plugin update introduce a small but significant [change](https://developer.android.com/studio/releases/gradle-plugin#version_properties_removed_from_buildconfig_class_in_library_projects) for library modules.  


&gt;**Version properties removed from BuildConfig class in library projects**  
&gt;  
&gt;For library projects only, the `BuildConfig.VERSION_NAME` and `BuildConfig.VERSION_CODE` properties have been removed from the generated `BuildConfig` class because these static values did not reflect the final values of the application’s version code and name, and were therefore misleading. Additionally, these values were discarded during manifest merging.

For myself, this breaks code of some of the libraries that I have built in the past few years for different projects and different companies. They will need to change this code when they upgrade to the new AGP. Luckily, a workaround is [available](https://issuetracker.google.com/issues/154275579):

`android.defaultConfig.buildConfigField "String", "VERSION_NAME", "\"$versionName\""`

Unfortunately, the Android team [do not plan](https://issuetracker.google.com/issues/154275579#comment10) to introduce an alternative versioning convention, and instead they suggest *'to create something adhoc'.*

Do you think this is the proper way to handle the library version reporting? Or you believe that the approach of Cocoa Touch frameworks with its auto-generated `ExampleFWVersionNumber` is preferrable?

Note that in the near future you will need to change your **build.gradle** even if you do not rely on the `versionCode`

&gt;In a future version of Android Gradle plugin, the `versionName` and `versionCode` properties will also be removed from the DSL for libraries.
## [4][🍭🚀💗 Tutorials about animations with Animators, Animated Vector Drawables, Shared Transitions, and more](https://www.reddit.com/r/androiddev/comments/je8qq8/tutorials_about_animations_with_animators/)
- url: https://github.com/SmartToolFactory/Animation-Tutorials
---

## [5][Very Simple Preview-only Camera2 Example](https://www.reddit.com/r/androiddev/comments/jenyxb/very_simple_previewonly_camera2_example/)
- url: https://gist.github.com/jbendtsen/a1c7fca87444621cb065fcb94a89d820
---

## [6][Automation bot for Android emulating the human actions?](https://www.reddit.com/r/androiddev/comments/jeos92/automation_bot_for_android_emulating_the_human/)
- url: https://www.reddit.com/r/androiddev/comments/jeos92/automation_bot_for_android_emulating_the_human/
---
I need an automation bot for performing repetitive actions inside multiple apps, copy/paste text, creating multiple accounts etc, what options there are available?

Thanks
## [7][Admob alternative?](https://www.reddit.com/r/androiddev/comments/jeopr5/admob_alternative/)
- url: https://www.reddit.com/r/androiddev/comments/jeopr5/admob_alternative/
---
I am an android developer and I cant use AdMob because it's still saying "your account is being verified" I created the AdMob in September! and still verifying, and I cant communicate with google it has bad support, so the question now, what are the best AdMob alternative to using, except Facebook
## [8][Google makes changes to Android faster than any solo dev or user can ever keep up with. Apple does this too with iOS. Their SF echo chamber causes this and outside SF people are very confused.](https://www.reddit.com/r/androiddev/comments/jeofk2/google_makes_changes_to_android_faster_than_any/)
- url: https://www.reddit.com/r/androiddev/comments/jeofk2/google_makes_changes_to_android_faster_than_any/
---
I watch my wife constantly guessing about which new swipe feature her apps are using and not being able to do something that should just be a button on the screen. I have the same issue and I'm a full time android dev. It is impossible to keep up with the changes to Android for both users and developers . Google needs to slow down and let people digest what they've done in the last five years. New versions of android should not break apps that are targeting one version older. I'm sure you guys can come up with tons of other stuff too. Discuss.
## [9][Best practices for developing with feature flags](https://www.reddit.com/r/androiddev/comments/jeeqlr/best_practices_for_developing_with_feature_flags/)
- url: https://www.reddit.com/r/androiddev/comments/jeeqlr/best_practices_for_developing_with_feature_flags/
---
Hello all,

I've been looking at ways to speed up development and testing via feature flags, and would love guidance from the broader community here on how everyone does it.

I've come across this neat little library with support for local and remote feature flags, which dynamically generates a UI around your features which you can drive tests from : https://github.com/JeroenMols/FeatureFlagExample

(The UK NHS COVID app is actually using this library!) 

Would love to hear people's approaches and/or recommended libraries you've used successfully to speed up testing with feature flags.
## [10][Videoview not playing](https://www.reddit.com/r/androiddev/comments/jekktn/videoview_not_playing/)
- url: https://www.reddit.com/r/androiddev/comments/jekktn/videoview_not_playing/
---
Hi. I´m having some problem with my videoView in a  recycleView. The video won´t play unless I call **binding**.**vvDiscoverDetail**.setZOrderOnTop(**true**). I have some texts on top of the view so I can´t really set it to top on Z axis. Does anyone know what might be the problem?

**val** uri = Uri.parse(item?.**video**)  
**binding**.**vvDiscoverDetail**.setMediaController(**null**)  
**binding**.**vvDiscoverDetail**.setVideoURI(uri)  
**binding**.**vvDiscoverDetail**.requestFocus()  
**binding**.**vvDiscoverDetail**.setZOrderOnTop(**true**)  
**binding**.**vvDiscoverDetail**.start()
## [11][Is there any scalable, customizable order management system for universal purpose (e.g. restaurant business; car wash business)](https://www.reddit.com/r/androiddev/comments/jemt2a/is_there_any_scalable_customizable_order/)
- url: https://www.reddit.com/r/androiddev/comments/jemt2a/is_there_any_scalable_customizable_order/
---
I'm an app dev and recently I built two apps that have the same module in their core.

1. Private app for a single restaurant for customers to place orders. Payments via app. Separate admin app to accept, decline, view and manage orders, etc. Feedback about the order. Order history
2. Private app for a car wash company, allowing users to select their vehicle location and place an order to have an employee arrive and wash their car. Payments via app. Also separate admin app to accept, decline, view and manage orders, etc. Feedback about the order. Order history

For backend I used firebase, firestore and some google cloud functions written in python. However, my question is, I wonder if I could have used instead some kind of a  3rd party order management system to integrate with, so that I wouldn't have to write the order placement and management backend code myself.

It should not be tied to a specific industry, should be customizable and allow to perform many  integrations (e.g. via API)

Is there anything good for this kind of purpose?

I don't even know if "order management system" is the right keyword for this, CRM also sounds like something of at least a partial fit for this.

I only found Hydra OMS, which from the first sight looks like it’s what I'm looking for, however, it doesn't seem widely known and I can barely find any referenced on the internet about it. Any alternative suggestions? Maybe an ecommerce platform would work for this?
## [12][Open source android development](https://www.reddit.com/r/androiddev/comments/jejydk/open_source_android_development/)
- url: https://www.reddit.com/r/androiddev/comments/jejydk/open_source_android_development/
---
I am a computer science student and i have a passion about android development. I have never done open source before i want to know which organisation should i choose as i also want to be a part of google summer of code next year how can i start making  contributions and what all skill set do i need before starting this ?
