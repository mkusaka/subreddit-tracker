# androiddev
## [1][App Feedback Thread - March 07, 2020](https://www.reddit.com/r/androiddev/comments/feuigx/app_feedback_thread_march_07_2020/)
- url: https://www.reddit.com/r/androiddev/comments/feuigx/app_feedback_thread_march_07_2020/
---
This thread is for getting feedback on your own apps.

####Developers:

- must **provide feedback** for others
- must include **Play Store**, **GitHub**, or **BitBucket** link
- must make top level comment
- must make effort to respond to questions and feedback from commenters
- may be open or closed source

####Commenters:

- must give **constructive feedback** in replies to top level comments
- must not include links to other apps

To cut down on spam, accounts who are too young or do not have enough karma to post will be removed. Please make an effort to contribute to the community before asking for feedback.

As always, the mod team is only a small group of people, and we rely on the readers to help us maintain this subreddit. Please report any rule breakers. Thank you.

\- Da Mods
## [2][Weekly Questions Thread - March 02, 2020](https://www.reddit.com/r/androiddev/comments/fcaf42/weekly_questions_thread_march_02_2020/)
- url: https://www.reddit.com/r/androiddev/comments/fcaf42/weekly_questions_thread_march_02_2020/
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
## [3][Rant - it took me over a year to finally get my app to this state.](https://www.reddit.com/r/androiddev/comments/ff47te/rant_it_took_me_over_a_year_to_finally_get_my_app/)
- url: https://i.redd.it/6f30b71xccl41.png
---

## [4][Unable to cast RippleDrawable to StateListDrawable after migrating to androidx](https://www.reddit.com/r/androiddev/comments/ffbh9m/unable_to_cast_rippledrawable_to/)
- url: https://www.reddit.com/r/androiddev/comments/ffbh9m/unable_to_cast_rippledrawable_to/
---
keep getting

    android.graphics.drawable.RippleDrawable cannot be cast to android.graphics.drawable.StateListDrawable

at the following line

    StateListDrawable gradientDrawable = (StateListDrawable) view.getBackground();

This wasn't happening before attempting to migrate where the view was

    android.support.v7.widget.AppCompatImageButton

but now that its

    androidx.appcompat.widget.AppCompatImageButton

it doesn't work

this is the XML for the view in question

    &lt;androidx.appcompat.widget.AppCompatImageButton android:id="@+id/view"
     android:layout_width="@dimen/bottomSheetPeekHeight" 
    
    android:layout_height="@dimen/bottomSheetPeekHeight" 
    
    android:layout_alignTop="@+id/dataEditorValueLayouts" 
    
    android:layout_alignBottom="@+id/dataEditorValueLayouts" 
    
    android:layout_marginEnd="@dimen/small_pad" android:scaleType="fitCenter" 
    
    android:padding="@dimen/small_pad" app:srcCompat="@drawable/download" 
    
    android:tint="@color/white" android:layout_toEndOf="@+id/dataEditorValueLayouts"
    
     style="@style/ButtonDefaultRuntime" android:visibility="visible" /&gt;

and here is the background defined in the style

    &lt;?xml version="1.0" encoding="utf-8"?&gt; 
    
    &lt;selector xmlns:android="http://schemas.android.com/apk/res/android"&gt;
    
     &lt;item android:state_pressed="true"&gt; 
    
    &lt;shape android:shape="oval"&gt; &lt;solid android:color="@color/faBlue" /&gt;
    
     &lt;stroke android:color="@color/button_border" android:width="3px" /&gt; 
    
    &lt;/shape&gt; 
    &lt;/item&gt; 
    &lt;item&gt; 
    &lt;shape android:shape="oval"&gt; &lt;solid android:color="@android:color/transparent" /&gt; &lt;stroke android:color="@color/button_border" android:width="3px" /&gt; &lt;/shape&gt; &lt;/item&gt; &lt;/selector&gt;
## [5][I created the Particle Clock using Jetpack Compose (source in comments!)](https://www.reddit.com/r/androiddev/comments/feysgd/i_created_the_particle_clock_using_jetpack/)
- url: https://imgur.com/b8qAz53
---

## [6][I'd like to buy testing devices, suggestions?](https://www.reddit.com/r/androiddev/comments/ffbg5d/id_like_to_buy_testing_devices_suggestions/)
- url: https://www.reddit.com/r/androiddev/comments/ffbg5d/id_like_to_buy_testing_devices_suggestions/
---
I'd like to get couple of cheap devices for testing purposes. You guys have any suggestion on what should I get?

Like, I suppose I'd want to have API21 device, small screen older device, latest API device, tablet.
## [7][Motion layout.](https://www.reddit.com/r/androiddev/comments/ferrqr/motion_layout/)
- url: https://v.redd.it/rg8y779sa7l41
---

## [8][Workmanager `setRequiresCharging(true)` not triggered untill battery level reaches 100%](https://www.reddit.com/r/androiddev/comments/ff7r6q/workmanager_setrequireschargingtrue_not_triggered/)
- url: https://www.reddit.com/r/androiddev/comments/ff7r6q/workmanager_setrequireschargingtrue_not_triggered/
---

## [9][how I get started with android development in 2020?](https://www.reddit.com/r/androiddev/comments/ffc5c0/how_i_get_started_with_android_development_in_2020/)
- url: https://www.reddit.com/r/androiddev/comments/ffc5c0/how_i_get_started_with_android_development_in_2020/
---
Some of this subreddit's wiki content seems to have dated APIs and patterns. 

Then, thought about checking [these](https://classroom.udacity.com/courses/ud9012) [two](https://classroom.udacity.com/courses/ud940) udacity classes or maybe this channel's [content](https://www.youtube.com/watch?v=jS0buQyfJfs&amp;list=PL0dzCUj1L5JGfHj1lwxOq67zAJV3e1S9S), but, now, I'm just worried about studying outdated information. 

Could y'all help me with tips or point me in a right direction on where to get started with native development? I'll seriously appreciate your input and don't really know which place is the right one to ask this question. Thanks in advance :[
## [10][Paid vs Ad supported apps](https://www.reddit.com/r/androiddev/comments/ffbx66/paid_vs_ad_supported_apps/)
- url: https://www.reddit.com/r/androiddev/comments/ffbx66/paid_vs_ad_supported_apps/
---
I have an app with adverts in, and a paid version which is exactly the same but no adverts. I have loads of people downloading the free ad-supported one, but even at less the $2 NOBODY buys it! I considered making my adverts show more regularly, but decided against it in case it meant fewer downloads - sometimes you just can't win!

[JumblyApps](http://jumblyapps.com/)
## [11][How are apps that contain a lot of information created and maintained?](https://www.reddit.com/r/androiddev/comments/ffbodz/how_are_apps_that_contain_a_lot_of_information/)
- url: https://www.reddit.com/r/androiddev/comments/ffbodz/how_are_apps_that_contain_a_lot_of_information/
---
For example with let's say a library app. There are a lot of books and you have to have the summaries, authors, who has it on hold, etc. all on the app and new books are added as well. But is all of that done by scratch or is there a way to have the information be auto added and layout by itself?
## [12][Spoof device specs for Kinemaster](https://www.reddit.com/r/androiddev/comments/ffazr9/spoof_device_specs_for_kinemaster/)
- url: https://www.reddit.com/r/androiddev/comments/ffazr9/spoof_device_specs_for_kinemaster/
---
I use Kinemaster on my phone to edit video, but due to my devices relatively low specs (Moto g5s plus), the app only allows me to edit videos 720p or below

How can I spoof my devices specs for this app to pretend to be a device with higher specs?
