# androiddev
## [1][App Feedback Thread - February 08, 2020](https://www.reddit.com/r/androiddev/comments/f0qz4w/app_feedback_thread_february_08_2020/)
- url: https://www.reddit.com/r/androiddev/comments/f0qz4w/app_feedback_thread_february_08_2020/
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
## [2][Weekly Questions Thread - February 03, 2020](https://www.reddit.com/r/androiddev/comments/ey5han/weekly_questions_thread_february_03_2020/)
- url: https://www.reddit.com/r/androiddev/comments/ey5han/weekly_questions_thread_february_03_2020/
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
## [3][A simple DIY library to generate your own custom Calendar View using RecyclerView, written in Kotlin](https://www.reddit.com/r/androiddev/comments/f165ul/a_simple_diy_library_to_generate_your_own_custom/)
- url: https://github.com/tejpratap46/RecyclerCalendarAndroid
---

## [4][Anyone done white-boarding in Kotlin?](https://www.reddit.com/r/androiddev/comments/f133r5/anyone_done_whiteboarding_in_kotlin/)
- url: https://www.reddit.com/r/androiddev/comments/f133r5/anyone_done_whiteboarding_in_kotlin/
---
I have an interview at Facebook and a few other companies coming up and I am working through the Cracking the Coding Interview book. Has anyone done white boarding at any company in Kotlin? Did you use some of the built in Kotlin utility functions (reverse, sorted, filter, map, flatmap, toCharArray()...) and if so what did you for Big O for them? Memorize them, best guess, or just implement yourself?

Also thoughts on using ArrayDeque for both Stacks and Queue questions? 

Do you use coroutines, or just do single threaded solutions unless asked? 

And any other things to consider for Kotlin vs Java white boarding? One thing I have noticed is that white board questions seem to often have inputs/outputs as Arrays rather than Lists/MutableLists which I generally use when coding in Kotlin.
## [5][Google Play NOT indexing new apps again.](https://www.reddit.com/r/androiddev/comments/f16cq6/google_play_not_indexing_new_apps_again/)
- url: https://www.reddit.com/r/androiddev/comments/f16cq6/google_play_not_indexing_new_apps_again/
---
It seems the problem from November is back. Apps published from \~Jan 27 does not appear in search results. You can find them only by removing spaces in the title or changing them to underscore.
## [6][KitKat and TLSv1.2](https://www.reddit.com/r/androiddev/comments/f17h6v/kitkat_and_tlsv12/)
- url: https://publicobject.com/2020/02/09/kitkat-and-tlsv1-2/
---

## [7][Google and Dark Mode](https://www.reddit.com/r/androiddev/comments/f17r1w/google_and_dark_mode/)
- url: https://www.reddit.com/r/androiddev/comments/f17r1w/google_and_dark_mode/
---
Has anybody noticed that most apps by Google e.g Keep, Play Store, Calendar, etc when switched to dark mode use colors which are way too light compared to those suggested by [the Material Design page on Dark Mode](https://material.io/design/color/dark-theme.html#usage)

Does anyone know why this is the case?
## [8][How to host app-ads.txt on Github Pages?](https://www.reddit.com/r/androiddev/comments/f15zty/how_to_host_appadstxt_on_github_pages/)
- url: https://www.reddit.com/r/androiddev/comments/f15zty/how_to_host_appadstxt_on_github_pages/
---
I uploaded app-ads.txt to my Github Page: [https://gmikhail.github.io/app-ads.txt](https://gmikhail.github.io/app-ads.txt)

I set my website [https://gmikhail.github.io/](https://gmikhail.github.io/) in Google Dev Console (Main page with apps list -&gt; Settings -&gt; Developer account -&gt; Account details -&gt; Website).

I also tried to set the website on the app page (App control panel -&gt; App page -&gt; App description -&gt; Website).

But so far (three days have already passed) on [https://apps.admob.com/v2/apps/appadstxt](https://apps.admob.com/v2/apps/appadstxt) instead of urls are dashes is displayed.

What am I doing wrong? Is it even possible to host app-ads.txt on Github Pages?
## [9][Firestore alternatives? Is there an 'industry standard'?](https://www.reddit.com/r/androiddev/comments/f16u8x/firestore_alternatives_is_there_an_industry/)
- url: https://www.reddit.com/r/androiddev/comments/f16u8x/firestore_alternatives_is_there_an_industry/
---
Not familiar with databases at all, other than firebase. But I'm developing an app that would require store of data and login authentication.

Let's the app records workouts and sorts them into subcategories. 

Using firestore, I would do something like have a document of all the login auth encrypted, each login and has a subdocument of each workout year, which has subdocument of each workout month, which holds an instance of the workout class. The workout class has various attributes, like total time, sets time of each exercise etc.

I've made a similar app in this way for practicing development, which is why I've described it that way. While this work implement fine, I feel like this isn't the 'proper' way. Login auth seems sketchy, and firestore implementation does give alot of headaches. 

What would be the best way/service to use to go about this?
## [10][Learning iOS as an Android developer](https://www.reddit.com/r/androiddev/comments/f18lc9/learning_ios_as_an_android_developer/)
- url: https://www.reddit.com/r/androiddev/comments/f18lc9/learning_ios_as_an_android_developer/
---
I have been developing for Android for the past 3 years and would like to dive into the iOS market.   


I know that there's a lot of Swift and iOS tutorials on the internet, but what I'm looking for are those with comparisons to Android development. For example a tutorial saying the equivalent for Android stuff  : "A view in android is \_\_\_ in iOS"  


Anyone know of a good one?
## [11][How to capture screenshots in AndroidStudio without, capturing bottom soft buttons?](https://www.reddit.com/r/androiddev/comments/f177ob/how_to_capture_screenshots_in_androidstudio/)
- url: https://www.reddit.com/r/androiddev/comments/f177ob/how_to_capture_screenshots_in_androidstudio/
---
All the while, I was able to capture screenshots in AndroidStudio + Emulator, without capturing the bottom soft buttons (Back, Home, Recent)

&amp;#x200B;

I was using the bottom left camera icon button to capture screenshot - [https://i.imgur.com/BrpQB03.png](https://i.imgur.com/BrpQB03.png)

&amp;#x200B;

However, I'm not sure what I have done, which causes the bottom soft buttons (Back, Home, Recent) always be capture. For instance, this is one of the capture screenshot. I do not want the bottom soft buttons bar - [https://i.imgur.com/fRGyTpI.png](https://i.imgur.com/fRGyTpI.png)

&amp;#x200B;

This is my AVD manager - [https://i.imgur.com/emqPv5n.png](https://i.imgur.com/emqPv5n.png)

&amp;#x200B;

Any idea how I can avoid from capturing bottom soft buttons, when I was capturing screenshots using AndroidStudio + Emulator?

&amp;#x200B;

Thanks.
## [12][Setting Window background to @null in XML doesn't work after night mode is toggled](https://www.reddit.com/r/androiddev/comments/f16usx/setting_window_background_to_null_in_xml_doesnt/)
- url: https://www.reddit.com/r/androiddev/comments/f16usx/setting_window_background_to_null_in_xml_doesnt/
---
I'm using DayNight theme.

I set the window background like this
```
&lt;item name="android:windowBackground"&gt;@color/window_bg&lt;/item&gt;
```

and set the colors like this
in colors.xml (default)
```
&lt;color name="window_bg"&gt;@null&lt;/color&gt;
```
in colors.xml (night)
```
&lt;color name="window_bg"&gt;@andoid:color/black&lt;/color&gt;
```

Also, I have this thing in my Main Activity's onCreate()
```
AppCompatDelegate.setDefaultNightMode(AppCompatDelegate.MODE_NIGHT_FOLLOW_SYSTEM)
```

And the activity declaration in the manifest is just this
```
&lt;activity
    android:name=".MainActivity"
    android:configChanges="orientation|keyboardHidden"
    android:screenOrientation="portrait"&gt;
```

So, when I toggle the night mode (from quick setting tiles) from dark to light, only the window background color doesn't change. I expected the color to be white but it was black.
