# androiddev
## [1][App Feedback Thread - September 26, 2020](https://www.reddit.com/r/androiddev/comments/j056re/app_feedback_thread_september_26_2020/)
- url: https://www.reddit.com/r/androiddev/comments/j056re/app_feedback_thread_september_26_2020/
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
## [2][Weekly Questions Thread - September 21, 2020](https://www.reddit.com/r/androiddev/comments/iwy6cm/weekly_questions_thread_september_21_2020/)
- url: https://www.reddit.com/r/androiddev/comments/iwy6cm/weekly_questions_thread_september_21_2020/
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
## [3][Is flutter really the future?](https://www.reddit.com/r/androiddev/comments/j024of/is_flutter_really_the_future/)
- url: https://www.reddit.com/r/androiddev/comments/j024of/is_flutter_really_the_future/
---
I am new android dev, and started working at my first job as android dev around 8 months back. I've been working with MVP/MVVM at this job and using Kotlin. Honestly I've started to love Kotlin, since during my undergraduate days, I did a lot of sports programming in Java and Kotlin seems way more easier and elegant to use. Recently I was planning to look into **Jetpack** and **KMM,** when I stumbled upon Flutter and a lot of articles how flutter is the future. I don't understand how can writing in native be the future? Also I agree cross platform is an advantage but with KMM coming out, I guess Kotlin can be used too. Since I'm pretty new, all the experienced devs out there, should I really learn Flutter to be employable in future or continue with Jetpack and KMM?
## [4][Floating Windows on Android 1: Jetpack Compose &amp; Room](https://www.reddit.com/r/androiddev/comments/j03g09/floating_windows_on_android_1_jetpack_compose_room/)
- url: https://localazy.com/blog/floating-windows-on-android-1-jetpack-compose-and-room
---

## [5][Rant: Surprised how uncomfortable/bad Room DB is for developers](https://www.reddit.com/r/androiddev/comments/izjbhq/rant_surprised_how_uncomfortablebad_room_db_is/)
- url: https://www.reddit.com/r/androiddev/comments/izjbhq/rant_surprised_how_uncomfortablebad_room_db_is/
---
After developing backend applications with Spring Boot for a while, I recently joined a mobile dev team again. They use Room as DB.    
I can see the pros: like being developed by Google, optimised for mobile and so on....   


But i was shocked how complicated such a trivial task as one-to-many relationships are with Room.  


You can't just have a list of items inside an entity. No, you need to manually persist all items in their own table and manually set the foreign keys. But if you let Room generate the ID of the "One" for you, it's pretty complicated to get this generated ID back, because room will only tell you the row number as result of the insert, but not the ID of the inserted item.   


Fortunately Room has at least some sort of support for reading one-to-many relationships. So you can define a class with the relation and room will at least create a join query.   


Maybe i miss something but based on the [official documentation](https://developer.android.com/training/data-storage/room/relationships) and some research i did that's how it needs to be done :(  


Honestly, in 2020 i would have expected way better tooling for RDBMS even on mobile platforms.
## [6][Analyzing the Source Code - How Glide Image Loading Library works Internally - Article](https://www.reddit.com/r/androiddev/comments/izofxs/analyzing_the_source_code_how_glide_image_loading/)
- url: https://codingwithmitch.com/blog/image-loading-with-glide-how-it-works/
---

## [7][Floating Windows on Android 1: Jetpack Compose &amp; Room](https://www.reddit.com/r/androiddev/comments/j03j5d/floating_windows_on_android_1_jetpack_compose_room/)
- url: https://www.reddit.com/r/androiddev/comments/j03j5d/floating_windows_on_android_1_jetpack_compose_room/
---
Hi, Guys,

I published a series of articles about how to make floating windows (such as Facebook Heads and similar) on Android. 

For 7 years, I work on Floating Apps ([https://floatingapps.net](https://floatingapps.net)) and now I'm sharing all tips I learned along the way. 

**The first lesson: Jetpack Compose &amp; Room**

[https://localazy.com/blog/floating-windows-on-android-1-jetpack-compose-and-room](https://localazy.com/blog/floating-windows-on-android-1-jetpack-compose-and-room)

Do you like it? Feel free to get in touch!
## [8][Any good libraries/tutorials for working with the android file system?](https://www.reddit.com/r/androiddev/comments/j03hr2/any_good_librariestutorials_for_working_with_the/)
- url: https://www.reddit.com/r/androiddev/comments/j03hr2/any_good_librariestutorials_for_working_with_the/
---
Working with "internal" files is fine for me but anything outside of that I find myself crawling through out of date answers on stackoverflow, confusing documentation from Google and running into constant issues trying to understand how all the various concepts work together. Anyone know any good tutorials or libraries that can help me out? Thanks!
## [9][DataStore: Security - Styling Android](https://www.reddit.com/r/androiddev/comments/izk0bb/datastore_security_styling_android/)
- url: https://blog.stylingandroid.com/datastore-security/
---

## [10][Android Modularization: Make a reusable C++ module.( Intermediate )](https://www.reddit.com/r/androiddev/comments/izym0v/android_modularization_make_a_reusable_c_module/)
- url: https://medium.com/@shalommathews05/android-modularization-make-a-reusable-c-module-intermediate-45c921f2608e
---

## [11][Google Play Billing Library v2/v3: How to stop someone maliciously preventing an acknowledgement and therefore get their purchase refunded?](https://www.reddit.com/r/androiddev/comments/izo6kg/google_play_billing_library_v2v3_how_to_stop/)
- url: https://www.reddit.com/r/androiddev/comments/izo6kg/google_play_billing_library_v2v3_how_to_stop/
---
I've been implementing Google Play Billing Library v3, and noticed that starting v2 we must [acknowledge a purchase after granting entitlement to the purchased feature][1], otherwise the user is refunded within 3 days. So essentially the lifecycle is something like:

1) User purchases something, e.g. subscription to a premium account. 

2) The client app sends the purchase and signature to the backend so the user is entitled to the purchase, e.g. premium account.

3) The backend responds with a success, we then acknowledge with Google Play the purchase was entitled (otherwise a refund is given). 

However, that got me thinking, what is stopping a malicious user from modifying the APK (or even blocking the network call) so that the acknowledgment call is _not_ made, therefore the user is refunded within 3 days, but entitlement (step 2 above) is still given. What can I do the stop this within my app? How can I make sure the acknowledgment definitely goes through if the entitlement is given?

  [1]: https://developer.android.com/google/play/billing/integrate#process
## [12][Is it a bad practice to read from LiveData inside ViewModel functions?](https://www.reddit.com/r/androiddev/comments/izmccf/is_it_a_bad_practice_to_read_from_livedata_inside/)
- url: https://www.reddit.com/r/androiddev/comments/izmccf/is_it_a_bad_practice_to_read_from_livedata_inside/
---
Let's say I have a list of fetched comments: `val comments = MutableLiveData&lt;Comment&gt;()`. When a comment is clicked a ViewModel function `fun onCommentClicked(commentIndex: Int)` is invoked, inside which I have to acquire clicked comment in order to do something. I would have to write `val comment = comments.value!!.filter {` [`it.id`](https://it.id) `== commentIndex }`. 

Generally speaking I found that very often I read from LiveData variables inside ViewModel code. My SignInViewModel has two live datas `val username = MutableLiveData&lt;String&gt;(); val password = MutableLiveData&lt;String&gt;()` so when *sign in button* is pressed I would invoke a method like this one:

    fun signIn() {
        val params = SignInParams(username.value!!, password.value!!)
        signInUseCase(params)
        ...
    }

For some reason it seems a bad practice to me that I so often explicitly read from LiveData. Source code seems to be polluted with `.value!!` statements. 

 I thought about keeping some data in additional backing lateinit properties so whenever I fetch list of comments  I would have to set two variables, and later in code I could just easily read from `_comments` variable without a need to get content from livedata and assert it's not null:

    private lateinit var _comments: List&lt;Comments&gt;
    val comments = MutableLiveData&lt;List&lt;Comments&gt;()
    fun onCommentsFetched(fetchedComments: List&lt;Comments&gt;) {
      _comments = fetchedComments
      comments.value = _comments
    }
    
    fun someFun(){
      val comment = _comments.filter { ....
    }

So is it a bad practice to read from LiveData inside ViewModel functions? Should their content be only exposed to view? How to solve this problem?
