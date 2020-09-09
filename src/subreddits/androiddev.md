# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/io6s7l/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/io6s7l/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - September 07, 2020](https://www.reddit.com/r/androiddev/comments/io5e9j/weekly_questions_thread_september_07_2020/)
- url: https://www.reddit.com/r/androiddev/comments/io5e9j/weekly_questions_thread_september_07_2020/
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
## [3][Rendering Markdown using Jetpack Compose](https://www.reddit.com/r/androiddev/comments/ipcftx/rendering_markdown_using_jetpack_compose/)
- url: https://www.reddit.com/r/androiddev/comments/ipcftx/rendering_markdown_using_jetpack_compose/
---
I wrote a post and a sample app for how to render Markdown content using Jetpack Compose. 

Hope you like it!

[https://www.hellsoft.se/rendering-markdown-with-jetpack-compose/](https://www.hellsoft.se/rendering-markdown-with-jetpack-compose/)
## [4][PSA: BuildCompat.isAtLeastR() incompatible with released version of 11/R](https://www.reddit.com/r/androiddev/comments/ipdv64/psa_buildcompatisatleastr_incompatible_with/)
- url: https://www.reddit.com/r/androiddev/comments/ipdv64/psa_buildcompatisatleastr_incompatible_with/
---
This is defined as:
```
public static boolean isAtLeastR() {  
 return VERSION.CODENAME.length() == 1  
 &amp;&amp; VERSION.CODENAME.charAt(0) &gt;= 'R'  
 &amp;&amp; VERSION.CODENAME.charAt(0) &lt;= 'Z';  
}
```
The release version of 11 `CODENAME` is `"REL"` so `length() &gt; 1` and returns `false`.

This is using:
```
implementation 'androidx.core:core:1.3.1'
implementation 'androidx.core:core-ktx:1.3.1'
```
The latest stable version (at time of writing).

Be warned!

Workaround:
`Build.VERSION.SDK_INT &gt;= 30` or `Build.VERSION.SDK_INT &gt;= Build.VERSION_CODES.R`

Edit:
Markdown had italics highlighting by mistake. Forgot `=` in workaround.

Additional edit:
I now realise that `BuildCompat` is precisely for targeting pre-release versions of Android. So using it in production code is not intended. Either way it's likely to be fixed with a real version check in future.
## [5][Android 11 has arrived](https://www.reddit.com/r/androiddev/comments/ioxlor/android_11_has_arrived/)
- url: https://blog.google/products/android/android-11
---

## [6][Do people still use MPAndroidChart?](https://www.reddit.com/r/androiddev/comments/ipe4oe/do_people_still_use_mpandroidchart/)
- url: https://www.reddit.com/r/androiddev/comments/ipe4oe/do_people_still_use_mpandroidchart/
---
I was just going to suggest using it for a project and saw the last release was well over a year ago. As well as that, there's \~1700 issues on the [repo.](https://github.com/PhilJay/MPAndroidChart) Just wondering if people are still actively adding this to projects, and if not, what's the best alternative? I'm looking for something that has a few different graphs, and ideally supports some interaction.
## [7][ShapeableImageView - Material components for android](https://www.reddit.com/r/androiddev/comments/ip6q3q/shapeableimageview_material_components_for_android/)
- url: https://howtodoandroid.com/shapeableimageview-material-components-android/
---

## [8][ViewBinding Delegate — one line](https://www.reddit.com/r/androiddev/comments/ipe6tm/viewbinding_delegate_one_line/)
- url: https://medium.com/@hoc081098/viewbinding-delegate-one-line-4d0cdcbf53ba
---

## [9][Android 11 Wireless Debugging](https://www.reddit.com/r/androiddev/comments/ipd0ph/android_11_wireless_debugging/)
- url: https://www.reddit.com/r/androiddev/comments/ipd0ph/android_11_wireless_debugging/
---
Here's a video on how to use wireless debugging in Android 11.

[https://www.youtube.com/watch?v=azZPiGY9lFs](https://www.youtube.com/watch?v=azZPiGY9lFs)
## [10][Questions about Square Workflow](https://www.reddit.com/r/androiddev/comments/ipfnvu/questions_about_square_workflow/)
- url: https://www.reddit.com/r/androiddev/comments/ipfnvu/questions_about_square_workflow/
---
I've seen the video about Square Workflow at Droidcon and tried it out with a dummy project, but I still have a few questions about it and think this is the right place to discuss about it.

* Do you use it in production ? What is your experience with it ?
* Does it work well with asynchronous source of data ? I'm thinking about an app that heavily rely on Firestore to write and read data.
* Does it work well with more complex Android UI Components like ViewPager2 ?
## [11][Memoji Stickers For Android (Proper Catagories)](https://www.reddit.com/r/androiddev/comments/ipfhh4/memoji_stickers_for_android_proper_catagories/)
- url: https://www.reddit.com/r/androiddev/comments/ipfhh4/memoji_stickers_for_android_proper_catagories/
---
Switching from Iphone to Android ,users miss one thing while texting, MEMOJIs. So i decided to find a way to use the same memojis of apple on android.
Now there are alot of apps already providing you with memojis on android, problem is that they havent catagorized it with a proper plain, or most of them havent catagorized them at all. 
So , i have grouped the memojis on the basis of proper catagories based on the Hair color, Skin Tone and wearing accessories. The first build has 9 catagories, new catagories will be uploaded every week.

Question from developers ?

I need suggestion in code,  on how to add tags to each memoji, so that users can choose on the basis of those tags. The tags might include Hair colors, Face Color , Accessories etc


Google Play Store Link : 
[Memoji stickers for android](https://play.google.com/store/apps/details?id=com.memoji_wasticekrs_applememoji_sortscript.stickers)

Thank you :)
## [12][Getting started and developing an application with Jetpack Compose](https://www.reddit.com/r/androiddev/comments/ipf2u3/getting_started_and_developing_an_application/)
- url: https://blog.codemagic.io/getting-started-with-jetpack-compose/
---

