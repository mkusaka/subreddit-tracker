# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/ghmv5w/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/ghmv5w/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - May 11, 2020](https://www.reddit.com/r/androiddev/comments/ghlel5/weekly_questions_thread_may_11_2020/)
- url: https://www.reddit.com/r/androiddev/comments/ghlel5/weekly_questions_thread_may_11_2020/
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
## [3][Gradle migration to Kotlin DSL](https://www.reddit.com/r/androiddev/comments/giv28v/gradle_migration_to_kotlin_dsl/)
- url: https://www.reddit.com/r/androiddev/comments/giv28v/gradle_migration_to_kotlin_dsl/
---
I recently migrated Gradle to Kotlin DSL scripts in my side project to check it out. So I thought I will share a link to commit diff, maybe someone will find it useful in case of doing that.    
[Commit on GitHub](https://github.com/OMIsie11/CovidNow/commit/760afeef1f90dc0199146d27b379a43ba3788f34)
## [4][High performance graphing/plotting library](https://www.reddit.com/r/androiddev/comments/giwdrl/high_performance_graphingplotting_library/)
- url: https://www.reddit.com/r/androiddev/comments/giwdrl/high_performance_graphingplotting_library/
---
I am developing an android app that is in part displaying measurements from several sensors (e.g. acceleration). I want to display the last seconds of measurements by using some kind of plotting library. Right now Im using AndroidPlot which is a nice library and gives me ok-ish results. I usually have up to 2000 datapoints and want to refresh 1-4 plots with a refreshrate somewhere between 10-30 Hz. On my mid-range device this results in UI lags when using AndroidPlot.


Things I have tried so far or thought about:

- Using GraphView -&gt; Performance was worse
- Webview and Plot.ly -&gt; Great performance when just having one plot, but much worse when I had more than one
- MPAndroidChart -&gt; Havent tried it, but does not seem to be optimized for dynamic plots 
- Downsampling -&gt; Not trivial to do it right, complex algorithms take to much computation time, simple ones to produce adequate results
- AAChartKit-kotlin -&gt; Havent tried it, but seems to be using Webviews as well, so I expect same performance as with plot.ly

Does anyone maybe know a library that might be more performant or has some tips what could improve the performance of the graphs?

I started Android-development half a year ago (part-time) so I do not feel experienced enough and do not have the resources to develop something on my own (e.g. using OpenGL). I also do not want to use a commercial library.
## [5][Earth Imagery App following Jetpack](https://www.reddit.com/r/androiddev/comments/giwvg2/earth_imagery_app_following_jetpack/)
- url: https://github.com/nomanr/earth-imagery-android
---

## [6][Massive Android Ecosystem Cheat Sheet 2020](https://www.reddit.com/r/androiddev/comments/giheiy/massive_android_ecosystem_cheat_sheet_2020/)
- url: https://www.reddit.com/r/androiddev/comments/giheiy/massive_android_ecosystem_cheat_sheet_2020/
---
&amp;#x200B;

https://i.redd.it/mhhz4m3xtdy41.gif

Link to repo [https://github.com/igorwojda/android-ecosystem-cheat-sheet](https://github.com/igorwojda/android-ecosystem-cheat-sheet)

Link to retweet [https://twitter.com/igorwojda/status/1260275259107835905](https://twitter.com/igorwojda/status/1260275259107835905)

Feedback highly welcome

Enjoy
## [7][Tutorial: Intro Slider with pretty animations, state management and usage of architecture components](https://www.reddit.com/r/androiddev/comments/giy43x/tutorial_intro_slider_with_pretty_animations/)
- url: https://www.reddit.com/r/androiddev/comments/giy43x/tutorial_intro_slider_with_pretty_animations/
---
Recently had to develop an intro slider for one of the projects and decided to write an article on how to build one from ground. In my opinion it serves as a nice real life example of architecture components usage and state management.

[Link for part 1](https://medium.com/@alexandr.zhurkov/android-intro-slider-with-state-management-using-androidx-architecture-components-part-1-caeba1db7a32)

[Link for part 2](https://medium.com/@alexandr.zhurkov/android-intro-slider-with-state-management-using-androidx-architecture-components-part-2-2e927211b9a5)

[Link to GitHub](https://github.com/xzhorikx/slider-intro)
## [8][Android Studio 4.0 Release Candidate 1 available](https://www.reddit.com/r/androiddev/comments/gig7r8/android_studio_40_release_candidate_1_available/)
- url: https://androidstudio.googleblog.com/2020/05/android-studio-40-release-candidate-1.html
---

## [9][First Look at Kotlin 1.4-M2: Standard Library Improvements](https://www.reddit.com/r/androiddev/comments/gicmd3/first_look_at_kotlin_14m2_standard_library/)
- url: https://blog.jetbrains.com/kotlin/2020/05/1-4-m2-standard-library/
---

## [10][Struck with Firestore Query in Android](https://www.reddit.com/r/androiddev/comments/giwsbl/struck_with_firestore_query_in_android/)
- url: https://www.reddit.com/r/androiddev/comments/giwsbl/struck_with_firestore_query_in_android/
---
I have Firestore DB structure like this. In this app, I'm able to display all Time slots of Lynda to the Customer.

&amp;#x200B;

[Customer Collection ](https://preview.redd.it/385xntj8hiy41.png?width=956&amp;format=png&amp;auto=webp&amp;s=40070d66edccfd30f1878b7436fa2e1f81850bcc)

&amp;#x200B;

[Timeslots Collection for User \(which means Service provider\)](https://preview.redd.it/k1a4udqchiy41.png?width=875&amp;format=png&amp;auto=webp&amp;s=c200734161f1c4aa81ae7b6165aad3564345b235)

&amp;#x200B;

[User Collection](https://preview.redd.it/lqfwc58ghiy41.png?width=905&amp;format=png&amp;auto=webp&amp;s=4c6d61ea758624dd1cda535e2fba1c86ebdc6fea)

&amp;#x200B;

[User Appointments by Date](https://preview.redd.it/89fhglcjhiy41.png?width=955&amp;format=png&amp;auto=webp&amp;s=5b7efcd0a2504fa4f312ea5769042cfe5ff61fe4)

&amp;#x200B;

**The query for getting all Timeslots of Lynda1102**

`// get all time slots`

`override fun getAllSlots(): MutableLiveData&lt;List&lt;Slots&gt;&gt; {`

`remoteDB.collection("Timeslots").whereEqualTo("id", "Lynda1102")`

`.addSnapshotListener { querySnapshot, e -&gt;`

&amp;#x200B;

`// check if null`

`if (e != null) {`

`Timber.e("Listen failed. $e")`

`savedAllTimeSlots.value = null`

`}`

`val saveTimeSlots: MutableList&lt;Slots&gt; = mutableListOf()`

`for (doc in querySnapshot!!) {`

`val slotItem = doc.toObject(RemoteTimeSlots::class.java).apply { id =` [`doc.id`](https://doc.id) `}`

`.mapToTimeSlots()`

`saveTimeSlots.add(slotItem)`

`}`

`// Post value`

`savedAllTimeSlots.value = saveTimeSlots`

`}`

`return savedAllTimeSlots`

`}`

&amp;#x200B;

**Slots.kt (Model class)**

`data class Slots(val id:String,val slot:Long)`

`data class RemoteTimeSlots(var id: String = "",var slot: Long = 0)`

`fun RemoteTimeSlots.mapToTimeSlots(): Slots {`

`return Slots(id,slot)`

`}`

&amp;#x200B;

**Observering Data in MainActivity**

`// observe for service changes`

`viewModel.getAllSlots().observe(this,` [`androidx.lifecycle.Observer`](https://androidx.lifecycle.Observer) `{`

`timeSlotsAdapter.differ.submitList(it)`

&amp;#x200B;

`})`

&amp;#x200B;

I'm trying to figure out solution for 2 days 😢. Your small advice or suggestion can motivate me to finish off this things. Thanks in advance 🙏.

1. How to query &amp; show available slots for customer on a particular date.
2. Ex- On 13-05-2020 Lynda has a total of three slots. The customer booked slot no - 2. So now the other customers will have to see only the remaining two slots(slot1 &amp; slot3) on 13-05-2020.
## [11][Any way to remove an unpublished app from 'All Applications' in Google Play Console?](https://www.reddit.com/r/androiddev/comments/giya6b/any_way_to_remove_an_unpublished_app_from_all/)
- url: https://www.reddit.com/r/androiddev/comments/giya6b/any_way_to_remove_an_unpublished_app_from_all/
---
Is there any way to remove an unpublished app from 'All Applications' in Google Play Console? 

I have 1 unpublished app listed in my Console since 2014. Have emailed Google about it, and their advice is that it cannot be removed due to a very low number of downloads (which was done by our team for testing purposes). 

I am wondering if there are any possible way at all, as it really clutters my 'All Applications' section.
## [12][My admob account will immediately get suspended, what can I do ?](https://www.reddit.com/r/androiddev/comments/gilyp7/my_admob_account_will_immediately_get_suspended/)
- url: https://www.reddit.com/r/androiddev/comments/gilyp7/my_admob_account_will_immediately_get_suspended/
---
Today, I noticed that my admob clicks and estimated revenues are growing very quickly (1$ to 90$) and I don't know the reason, I'm really afraid of my account getting blocked, can I get some help please ?
