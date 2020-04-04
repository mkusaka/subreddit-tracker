# androiddev
## [1][App Feedback Thread - April 04, 2020](https://www.reddit.com/r/androiddev/comments/fusw1l/app_feedback_thread_april_04_2020/)
- url: https://www.reddit.com/r/androiddev/comments/fusw1l/app_feedback_thread_april_04_2020/
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
## [2][Weekly Questions Thread - March 30, 2020](https://www.reddit.com/r/androiddev/comments/fronhm/weekly_questions_thread_march_30_2020/)
- url: https://www.reddit.com/r/androiddev/comments/fronhm/weekly_questions_thread_march_30_2020/
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
## [3][Removing app from Amazon app store](https://www.reddit.com/r/androiddev/comments/fuktlz/removing_app_from_amazon_app_store/)
- url: https://www.reddit.com/r/androiddev/comments/fuktlz/removing_app_from_amazon_app_store/
---
Hi,

A few months ago I submitted an app to the Amazon app store. It is a paid app and has only had one download in that time. I would like to delete the listing and only focus on the Google Play Store for Android as it's not worth continually updating the Amazon listing whenever I change something when there is only one user.

My question is, if I remove the app from the store using the 'Remove from Appstore' button, what happens for the one user that has purchased the app? I can't tell from Amazon's reporting if the user has updated to the latest version. Will they still be able to update to the latest version even when it has been de-listed? Will they be able to re-download the app later if they remove it or switch devices, seeing as how they have already purchased it?

Thanks!
## [4][Kotlin Resources for Beginners](https://www.reddit.com/r/androiddev/comments/furu3w/kotlin_resources_for_beginners/)
- url: https://www.reddit.com/r/androiddev/comments/furu3w/kotlin_resources_for_beginners/
---
I have basic programming knowledge. I was trying out learning Java but it seems Kotlin is more geared towards android development. 

I cant find beginner resources for Kotlin. Please do recommend.
## [5][Modifying the default View touch/click behaviour - is there anything wrong with this tweak?](https://www.reddit.com/r/androiddev/comments/fumupu/modifying_the_default_view_touchclick_behaviour/)
- url: https://www.reddit.com/r/androiddev/comments/fumupu/modifying_the_default_view_touchclick_behaviour/
---
For a long time during the development of my app I've noticed a perceptible bit of input lag in the ripple animation that occurs when a View is touched or pressed. As I fiddled around with it yesterday, I decided to investigate.

I found these comments/code in the OnTouchEvent method of the default View class:

    // Walk up the hierarchy to determine if we're inside a scrolling container.
    boolean isInScrollingContainer = isInScrollingContainer();
    
    // For views inside a scrolling container, delay the pressed feedback for
    // a short period in case this is a scroll.

So essentially, if the view you're clicking is inside a scrolling container, the ripple is delayed by exactly 100ms (ViewConfiguration.getTapTimeout()). To me this delay was ruining the effect of the ripple entirely in my Recycler and Scroll views, and meant that even long presses or press and slides (where you slide your finger off a view to avoid triggering a click) had a perceptible "touch latency".

The 'fix' was very simple: override the 'shouldDelayChildPressedState()' method in your scroll containers to return false always, like so:

    class ImprovedRecyclerView @JvmOverloads constructor(
    context: Context, attrs: AttributeSet? = null, defStyleAttr: Int = 0 ) : RecyclerView(context, attrs, defStyleAttr) {
        override fun shouldDelayChildPressedState(): Boolean {
         return false
    }
    
    }

This method is very tightly constrained to the pressed state behaviour, so no scrolling or other touch behaviour is affected. The result in my opinion is much nicer, views feel more interactive and the ripple feels more connected to the touch. The only negative outcome is that when you start a scroll, you can trigger a ripple on the view your finger started scrolling from. It doesn't bother me though, it feels like the screen is just registering your input like a hover state in CSS.

I'm not sure whether I'm sharing this because it's interesting or whether I'm asking if this is fine to change, but I feel strange modifying such a low level behaviour. I think the changed behaviour is much nicer - is that all the justification I need?

Edit: also worth noting that for whatever reason, Coordinator and Frame Layouts also delay pressed states, so views inside them have the 100ms delay.
## [6][Passing data ownership to your app’s customers](https://www.reddit.com/r/androiddev/comments/fut7mr/passing_data_ownership_to_your_apps_customers/)
- url: https://medium.com/@andrei.lupic/passing-data-ownership-to-your-apps-customers-e57de14e3459
---

## [7][The settings panel in Android Q](https://www.reddit.com/r/androiddev/comments/fu6vc2/the_settings_panel_in_android_q/)
- url: https://youtu.be/kv8a5aaH9rc
---

## [8][Does income really depend that much on the logic of ad display and network performance? Check out the ultimate guide on essential in-app advertising metrics.](https://www.reddit.com/r/androiddev/comments/fursuz/does_income_really_depend_that_much_on_the_logic/)
- url: https://www.appodeal.com/home/blog/essential-metrics-for-mobile-ad-revenue-success/?utm_campaign=appodeal_comm&amp;utm_source=reddit&amp;utm_term=androiddev
---

## [9][Confused by In-App-Purchase Data Provided by Google (gross/net)](https://www.reddit.com/r/androiddev/comments/furqbf/confused_by_inapppurchase_data_provided_by_google/)
- url: https://www.reddit.com/r/androiddev/comments/furqbf/confused_by_inapppurchase_data_provided_by_google/
---
Hey guys,

I am a app publisher and I’m very confused by the In-App-Purchase data (Both In-App-Purchases and Subscriptions) provided by Google. Sometimes they provide gross (revenue without taxes and Google Play cut) and sometimes net (What I receive on my bank account). So far I have several data views:

-Acquisition reports in Google Play Developer console (net revenue, also ARPU KPI is net revenue but only calculates IAP?)
-Analytics in Firebase Dashboard (gross revenue, however Admob is net)
-Revenue Chart in Google Play Developer console Android app (gross revenue)
-Financial Reports - Revenue in Google Play Developer console (gross revenue)
-Google Play Developer Console - Payment Settings Shows net revenue 

Guys I’m a bit confused, Google seems to mix data as they like, however I need this data to make my strategic decisions on Marketing spent.
## [10][Internal Test app still pending Publication even after 6 days.](https://www.reddit.com/r/androiddev/comments/furq28/internal_test_app_still_pending_publication_even/)
- url: https://www.reddit.com/r/androiddev/comments/furq28/internal_test_app_still_pending_publication_even/
---
I uploaded and internal test version of my app to the Internal Test Track on 30th March and today 6th April it is Still Pending Publication. Is anyone else facing the same problem. I have to test a lot of things before I resume app development and for the past 6 days I have just been sitting idle waiting for my test app to get published. Is this normal
## [11][Jetpack ViewModel initialization](https://www.reddit.com/r/androiddev/comments/furo3r/jetpack_viewmodel_initialization/)
- url: https://www.rockandnull.com/jetpack-viewmodel-initialization/
---

## [12][Android Studio 4.0 - Translate Editor not work](https://www.reddit.com/r/androiddev/comments/fuql2a/android_studio_40_translate_editor_not_work/)
- url: https://www.reddit.com/r/androiddev/comments/fuql2a/android_studio_40_translate_editor_not_work/
---
Translate Editor stuck at screen "Loading string resource data". I tried to "Invalided &amp; Restart" but same result. This issue just happens on Android Studio 4.0. 

&amp;#x200B;

https://preview.redd.it/xi2togysprq41.png?width=1778&amp;format=png&amp;auto=webp&amp;s=1a2a9ae03165339d6ef3cd22ad631e5e76cd167a
