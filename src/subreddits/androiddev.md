# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/f8qpbi/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/f8qpbi/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - February 24, 2020](https://www.reddit.com/r/androiddev/comments/f8pltv/weekly_questions_thread_february_24_2020/)
- url: https://www.reddit.com/r/androiddev/comments/f8pltv/weekly_questions_thread_february_24_2020/
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
## [3][Justifying the Google punishment ladder - on the utility of keeping the escalation ladder from app bans to account perma-bans a secret from developers - can a "policy" be a secret policy? - and whether Google Play management employs this tactic as a defence mechanism to hide incompetence](https://www.reddit.com/r/androiddev/comments/f8loc7/justifying_the_google_punishment_ladder_on_the/)
- url: https://www.reddit.com/r/androiddev/comments/f8loc7/justifying_the_google_punishment_ladder_on_the/
---
We all know the general way that app bans work, and how they escalate to account perma-bans, and then "associated account bans".

The bigger question is, what is the utility of perma-banning developers with a policy that is inspired by the archaic **"three strikes and you are out" thinking of the Reagan era?**

Even though now we know 3 app bans does not always lead to perma-ban, and sometimes 1 app ban for a new dev leads to perma-ban, **what is the logic and motivation for such arbitrariness?**

Since a number of app bans are caused by dev unawareness or sloppiness, **how does an escalating punishment regime ensure better compliance?** Which sadist at Google though up this scheme?

&amp;nbsp;


**Unpredictable "policies"**

The current app ban to perma-ban escalation ladder is completely arbitrary and unpredictable from the developer point of view!

So how does it function as "policy", or as deterrent, if it is completely opaque to the developer?

The perception it creates is of arbitrary cruelty, under an unpredictable and temperamental boss.

&amp;nbsp;


**Who thought up this regime at Google?**

Which creative (but out of touch with reality) mind at Google thought all this up?

And which misanthrope then extended it to the **notoriously unfair "associated account ban"**, where a wife can be banned for the behavior of her husband, and the employer and friends can be banned without explanation out of the blue?

&amp;nbsp;


**Google's infinite memory**

As the understood behavior stands, there is an accumulation of ill will that a developer can accumulate, which leads to a tipping point where Google finally perma-bans a developer.

However if a developer has accumulated 3 app bans some time ago, due to mistakes, stupidity and even bad behavior, does Google start forgiving them after some time?  What is the motivation for infinite memory of developer sins - so an 18 year old's infractions are remembered well into his 20s?

Under what motivation should such a developer proceed to develop for Google?

If they have accumulated 3 app bans already, yet not been banned yet, should they continue developing for Google, or are they on the brink of an account perma-ban?

If there is no more margin of error - should they embark on a new project ?

Now it may be that Google erases the bad-will over time, so after some time, the dev is again given leeway to make mistakes.  Even if Google does have such an internal policy/algorithm, that is not known by the devs.

**It seems Google doesn't want you to have any visibility into your status at Google.**  Evidently it sees this as some sort of a leg up - a strategic advantage Google has, but devs don't have.

The end result if this opaqueness is, that a dev who has accumulated a few app bans may be better off making apps for a newer platform that telegraphs policies better to it's devs.

&amp;nbsp;


**Developer perceptions**

It is no wonder that a perception exists that Google Play sees **app developers as a nuisance**.  Or as some devs report, how their contacts within Google dry up at a hint of criticism.

This suggests there is a problem within Google Play management, where they have allowed, or possibly even **cultivated a culture of loathing towards app developers** - which is a sign of exasperation, and possibly incompetence within those corporate ranks - where **insularity is chosen as a defence** mechanism over openness (with openness, outsiders get a better handle to challenge management misbehavior).
## [4][Fabric is deprecated and will shut down after March 31, 2020](https://www.reddit.com/r/androiddev/comments/f8oh3s/fabric_is_deprecated_and_will_shut_down_after/)
- url: https://www.reddit.com/r/androiddev/comments/f8oh3s/fabric_is_deprecated_and_will_shut_down_after/
---
Fabric is deprecated and will shut down after March 31, 2020, and is being integrated into Firebase. Fabric's Beta will be replaced by Firebase App Distribution and Crashlytics will be replaced by Firebase App Distribution. I found a simple and concise article on how to setup Firebase App Distribution with fastlane. Check it out at   
[https://medium.com/@clementozemoya/automated-android-deployments-with-fastlane-and-firebase-app-distribution-b1d1905a4fe6](https://medium.com/@clementozemoya/automated-android-deployments-with-fastlane-and-firebase-app-distribution-b1d1905a4fe6)
## [5][How do you prevent fragment recreation with Android Navigation Component?](https://www.reddit.com/r/androiddev/comments/f8opw0/how_do_you_prevent_fragment_recreation_with/)
- url: https://www.reddit.com/r/androiddev/comments/f8opw0/how_do_you_prevent_fragment_recreation_with/
---
I have seen people try to save ViewModel, create custom Navigator, etc. What do you guys use and prefer and why?
## [6][A simple Disney app using the material transformation motion system.](https://www.reddit.com/r/androiddev/comments/f899oq/a_simple_disney_app_using_the_material/)
- url: https://github.com/skydoves/DisneyMotions
---

## [7][Android 11 funny things: android.R.string.yes and and android.R.string.no deprecated](https://www.reddit.com/r/androiddev/comments/f8aj0f/android_11_funny_things_androidrstringyes_and_and/)
- url: https://www.reddit.com/r/androiddev/comments/f8aj0f/android_11_funny_things_androidrstringyes_and_and/
---
`android.R.string.yes` and `android.R.string.no` are finally deprecated because they incorrectly match `android.R.string.ok` and `android.R.string.cancel`
## [8][Adobe Xd Design to Real Android/iOS App | Design Weekly](https://www.reddit.com/r/androiddev/comments/f8r88o/adobe_xd_design_to_real_androidios_app_design/)
- url: https://youtu.be/DWkFLlV_Q-Q
---

## [9][Suggestion: Make scoped storage better for devs](https://www.reddit.com/r/androiddev/comments/f8qook/suggestion_make_scoped_storage_better_for_devs/)
- url: https://www.reddit.com/r/Android/comments/f8qg4u/suggestion_make_scoped_storage_better_for_devs/
---

## [10][ItemDecoration doesn't work inside NestedScrollView](https://www.reddit.com/r/androiddev/comments/f8qlc5/itemdecoration_doesnt_work_inside_nestedscrollview/)
- url: https://www.reddit.com/r/androiddev/comments/f8qlc5/itemdecoration_doesnt_work_inside_nestedscrollview/
---
I have implemented RecyclerView with sticky headers and I did it through [ItemDecoration](https://stackoverflow.com/questions/32949971/how-can-i-make-sticky-headers-in-recyclerview-without-external-lib). It works as expected in the case of stand-alone RecyclerView. But I have two RecyclerViews, and I need to use NestedScrollView

    &lt;androidx.core.widget.NestedScrollView xmlns:android="http://schemas.android.com/apk/res/android"
        xmlns:app="http://schemas.android.com/apk/res-auto"
        android:layout_width="match_parent"
        android:layout_height="match_parent"
        android:fillViewport="true"
        android:scrollbars="none"&gt;
        &lt;LinearLayout
            android:layout_width="match_parent"
            android:layout_height="wrap_content"
            android:focusableInTouchMode="true"
            android:orientation="vertical"&gt;
            &lt;androidx.recyclerview.widget.RecyclerView
                android:id="@+id/menuNewsList"
                android:layout_width="match_parent"
                android:layout_height="170dp" /&gt;
            &lt;androidx.recyclerview.widget.RecyclerView
                android:id="@+id/listMenu"
                android:layout_width="match_parent"
                android:layout_height="match_parent"
                app:layoutManager="androidx.recyclerview.widget.LinearLayoutManager" /&gt;
        &lt;/LinearLayout&gt;
    &lt;/androidx.core.widget.NestedScrollView&gt;

Code initialization: 

    listMenu.apply {
         adapter = dishAdapter
         isNestedScrollingEnabled = false 
    } 
    listMenu.addItemDecoration(HeaderItemDecoration(listMenu, isHeader = isHeader()))

&amp;#x200B;

Without  NestedScrollView  it works excellent, but in my case ItemDecoration doesn't work.

I've found out next information:

* onDrawOver's are not called on scrolling because RecyclerView.draw() is not called as well.
* All items are created at the same time (so Adapter creates view holders for all items in the data). I think it happens because of NestedScrollView.
* I tried to force calling recyclerview's redraw on scrolling but it doesn't work

I considered combining 2 RecyclerViews into one, but first one has a horizontal scroll, so I'd like to avoid this solution.

Do you have any advice how to deal with it?
## [11][Alphaanimation with surface view.](https://www.reddit.com/r/androiddev/comments/f8q4yg/alphaanimation_with_surface_view/)
- url: https://www.reddit.com/r/androiddev/comments/f8q4yg/alphaanimation_with_surface_view/
---
I implemented alphaanimation in imageview to get some sort of dissolve effect for transition.
I want the same in surface view when transitioning from one to another.but using alphaanimation in surface view gives me black screen or not animating . Please suggest me a solution for this.
## [12][A/B test in android kotlin](https://www.reddit.com/r/androiddev/comments/f8ppqr/ab_test_in_android_kotlin/)
- url: https://www.reddit.com/r/androiddev/comments/f8ppqr/ab_test_in_android_kotlin/
---
In my app there are several a/b experiments running write now in onCreate() i have several if else for handling the flow, is there any particular coding style/practice that I can follow to make it cleaner
