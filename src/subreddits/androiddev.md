# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/guk02c/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/guk02c/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - June 01, 2020](https://www.reddit.com/r/androiddev/comments/guij23/weekly_questions_thread_june_01_2020/)
- url: https://www.reddit.com/r/androiddev/comments/guij23/weekly_questions_thread_june_01_2020/
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
## [3][Gradle 6.5 Released](https://www.reddit.com/r/androiddev/comments/gvqaeb/gradle_65_released/)
- url: https://docs.gradle.org/6.5/release-notes.html
---

## [4][Material Components for Android 1.3.0-alpha01](https://www.reddit.com/r/androiddev/comments/gvg8s0/material_components_for_android_130alpha01/)
- url: https://github.com/material-components/material-components-android/releases/tag/1.3.0-alpha01
---

## [5][Restore RecyclerView scroll position](https://www.reddit.com/r/androiddev/comments/gvt3s7/restore_recyclerview_scroll_position/)
- url: https://medium.com/androiddevelopers/restore-recyclerview-scroll-position-a8fbdc9a9334
---

## [6][A Google Play anomaly explained: How a trivial update for existing users can cause a sharp drop in daily *new* users.](https://www.reddit.com/r/androiddev/comments/gvcx1f/a_google_play_anomaly_explained_how_a_trivial/)
- url: https://www.reddit.com/r/androiddev/comments/gvcx1f/a_google_play_anomaly_explained_how_a_trivial/
---
It's a known fact that Google play factors in 'uninstalls' in your app ranking.

The reasoning behind this is to factor in 'retention' rate to predict if users like your app. or otherwise 80% uninstall it first or second day.

But here is an unwanted side effect:

I have a game with very good retention, low uninstalls in first weeks of download, therefore  it's ranked in a good place and gets a very healthy and failry constant daily downloads for last 10 months. Now I've released a minor 'typo' update.

As expected for a game that user do "finish playing with" and it continue setting for months as dead weight in user device, users do uninstall once they notice the game updating. ( We all do)

So after 8 months, the game accumulated tens of thousands "finished users"' who uninstalled in same day.

But should a user that played my game for weeks\\months and uninstalls it when I update , hurt app ranking, like same day uninstall ?

Retention among new users is still very high, but Google play interpreted those 50k uninstalls in one day as a bad sign.  rank plummeted and downloads where hurt considerably.

This is anomaly because:1- game still has hight retention for weeks.

2. if users who finished after weeks uninstalled it ASAP and didn't accumulate my app ranking wouldn't get punished.

3. if I updated more frequently and those 'dead weight' we're spread, ranking wouldn't hurt.

Key lessons:

\- if you have a "finite puzzle game" , and you update after a long time. consider this to a cause to download drop.

\- leave that typo alone !

I hope Play team, implements an heuristic that distinguished between "uninstalls" due to bad experience and "uninstall" from finished an happy users that accumulated over time.

&amp;#x200B;

An update can hurt you
## [7][Is Jetpack Navigation good?](https://www.reddit.com/r/androiddev/comments/gvl1gm/is_jetpack_navigation_good/)
- url: https://www.reddit.com/r/androiddev/comments/gvl1gm/is_jetpack_navigation_good/
---
I just went over the Codelab Jetpack Navigation project. I don't really feel it is very intuitive. Especially with all the XMLs. What do you guys think? What kinds of benefit does it brings to your app/projects?
## [8][how to implement adview with databinding correctly.](https://www.reddit.com/r/androiddev/comments/gvs4ef/how_to_implement_adview_with_databinding_correctly/)
- url: https://www.reddit.com/r/androiddev/comments/gvs4ef/how_to_implement_adview_with_databinding_correctly/
---
Hi, fellow android developers, I hope all of your doing great Currently I am developing news android app and I have implemented adview with databinding but when I run the code I am getting following exception java.lang.IllegalStateException: The ad size and ad unit ID must be set before loadAd is called.?

even I posted my question to StackOverflow https://stackoverflow.com/questions/62150185/java-lang-illegalstateexception-the-ad-size-and-ad-unit-id-must-be-set-before-l but  I have tried suggested answer it did not work at all any help will be appreciated I have tried set ad unit size and id XML as well as programmatically using kotlin it did not work at all
## [9][Playing with Github Actions](https://www.reddit.com/r/androiddev/comments/gvrwe1/playing_with_github_actions/)
- url: http://github.com/volkansahin45/github-actions
---

## [10][Create an expandable Recyclerview with the MergeAdapter](https://www.reddit.com/r/androiddev/comments/gvpv3c/create_an_expandable_recyclerview_with_the/)
- url: https://github.com/OHoussein/Android-Expandable-MergeAdapter
---

## [11][Complexity of requesting permissions - the default Kotlin example](https://www.reddit.com/r/androiddev/comments/gvtyzb/complexity_of_requesting_permissions_the_default/)
- url: https://www.reddit.com/r/androiddev/comments/gvtyzb/complexity_of_requesting_permissions_the_default/
---
I'm looking at the Kotlin default example here - halfway down the page

[https://developer.android.com/training/permissions/requesting#kotlin](https://developer.android.com/training/permissions/requesting#kotlin)

why should I need an alpha release android x package ([Activity 1.2.0-alpha02](https://developer.android.com/jetpack/androidx/releases/activity#1.2.0-alpha02)) just to get this to compile?

Whole thing seems overly complex for what should be pretty simple bit of code,

I mean at least they should make the samples valid code, not with "..." scattered about, which should really be // TODO comment.
## [12][Resources$NotFoundException on Android 10 devices.](https://www.reddit.com/r/androiddev/comments/gvquc5/resourcesnotfoundexception_on_android_10_devices/)
- url: https://www.reddit.com/r/androiddev/comments/gvquc5/resourcesnotfoundexception_on_android_10_devices/
---
Hey community!

**Problem**I've noticed a lot of unexpected crashes only on Android 10 devices. Crashes are pointing to missing resource ID in different layouts but I was not able to reproduce it even on same device and OS as I saw in crash reports.Weirdly enough, after hours of digging through google search results page 4 and further(I've got lost and results were in Mandarin) i found post pointing to `Developer options -&gt; Enable view attribute inspection`. After enabling it on my device I've could reproduce it but it's still not clear how to fix it. Disabling it helped and app users are not seeing crashes there.

**Question**Even though that helped I want to know what's the culprit of it and what is the right approach to fix it.

**Crash log**

    android.content.res.Resources$NotFoundException: Unable to find resource ID #0x7f040378
        at android.content.res.ResourcesImpl.getResourceTypeName(ResourcesImpl.java:305)
        at android.content.res.Resources.getResourceTypeName(Resources.java:2111)
        at android.content.res.Resources$Theme.getExplicitStyle(Resources.java:1769)
        at android.view.View.retrieveExplicitStyle(View.java:6731)
        at android.view.View.&lt;init&gt;(View.java:5639)
        at android.view.ViewGroup.&lt;init&gt;(ViewGroup.java:687)
        at android.widget.FrameLayout.&lt;init&gt;(FrameLayout.java:99)
        at android.widget.FrameLayout.&lt;init&gt;(FrameLayout.java:94)
        at androidx.cardview.widget.CardView.&lt;init&gt;(SourceFile:121)
        at com.google.android.material.card.MaterialCardView.&lt;init&gt;(SourceFile:52)
        at com.google.android.material.card.MaterialCardView.&lt;init&gt;(SourceFile:48)
        at java.lang.reflect.Constructor.newInstance0(Constructor.java)
        at java.lang.reflect.Constructor.newInstance(Constructor.java:343)
        at android.view.LayoutInflater.createView(LayoutInflater.java:854)
        at android.view.LayoutInflater.createViewFromTag(LayoutInflater.java:1006)
        at android.view.LayoutInflater.createViewFromTag(LayoutInflater.java:961)
        at android.view.LayoutInflater.rInflate(LayoutInflater.java:1123)
        at android.view.LayoutInflater.rInflateChildren(LayoutInflater.java:1084)
        at android.view.LayoutInflater.rInflate(LayoutInflater.java:1126)
        at android.view.LayoutInflater.rInflateChildren(LayoutInflater.java:1084)
        at android.view.LayoutInflater.inflate(LayoutInflater.java:682)
        at android.view.LayoutInflater.inflate(LayoutInflater.java:534)
        at androidx.databinding.DataBindingUtil.inflate$45cd3863(SourceFile:126)
        at androidx.databinding.ViewDataBinding.inflateInternal$5250870e(SourceFile:1366)
        at com.example.app.LoginFragmentBinding.inflate(SourceFile:1090)
        at com.example.app.LoginFragment.onCreateView(SourceFile:66)
        at androidx.fragment.app.Fragment.performCreateView(SourceFile:2439)
        at androidx.fragment.app.FragmentManagerImpl.removeRedundantOperationsAndExecute(SourceFile:1460)
        at androidx.fragment.app.FragmentManagerImpl.moveFragmentToExpectedState(SourceFile:1784)
        at androidx.fragment.app.BackStackRecord.executeOps(SourceFile:797)
        at androidx.fragment.app.FragmentManagerImpl.executeOps(SourceFile:2625)
        at androidx.fragment.app.FragmentManagerImpl.moveToState(SourceFile:2411)
        at androidx.fragment.app.FragmentManagerImpl.removeRedundantOperationsAndExecute(SourceFile:2366)
        at androidx.fragment.app.FragmentManagerImpl.execPendingActions(SourceFile:2273)
        at androidx.fragment.app.FragmentManagerImpl.dispatchStateChange(SourceFile:3273)
        at androidx.fragment.app.FragmentManagerImpl.dispatchActivityCreated(SourceFile:3229)
        at androidx.fragment.app.Fragment.performActivityCreated(SourceFile:2466)
        at androidx.fragment.app.FragmentManagerImpl.moveToState(SourceFile:1483)
        at androidx.fragment.app.FragmentManagerImpl.moveFragmentToExpectedState(SourceFile:1784)
        at androidx.fragment.app.BackStackRecord.executeOps(SourceFile:797)
        at androidx.fragment.app.FragmentManagerImpl.executeOps(SourceFile:2625)
        at androidx.fragment.app.FragmentManagerImpl.executeOpsTogether(SourceFile:2411)
        at androidx.fragment.app.FragmentManagerImpl.removeRedundantOperationsAndExecute(SourceFile:2366)
        at androidx.fragment.app.FragmentManagerImpl.execPendingActions(SourceFile:2273)
        at androidx.fragment.app.FragmentManagerImpl.dispatchStateChange(SourceFile:3273)
        at androidx.fragment.app.FragmentActivity.androidx.fragment.app.FragmentManagerImpl.dispatchActivityCreated(SourceFile:17229)
        at androidx.appcompat.app.AppCompatActivity.onStart(SourceFile:178)
        at com.ing.mobile.app.activities.INGMobileActivity.onStart(SourceFile:145)
        at android.app.Instrumentation.callActivityOnStart(Instrumentation.java:1433)
        at android.app.Activity.performStart(Activity.java:7978)
        at android.app.ActivityThread.handleStartActivity(ActivityThread.java:3472)
        at android.app.servertransaction.TransactionExecutor.performLifecycleSequence(TransactionExecutor.java:221)
        at android.app.servertransaction.TransactionExecutor.cycleToPath(TransactionExecutor.java:201)
        at android.app.servertransaction.TransactionExecutor.executeLifecycleState(TransactionExecutor.java:173)
        at android.app.servertransaction.TransactionExecutor.execute(TransactionExecutor.java:97)
        at android.app.ActivityThread$H.handleMessage(ActivityThread.java:2147)
        at android.os.Handler.dispatchMessage(Handler.java:107)
        at android.os.Looper.loop(Looper.java:237)
        at android.app.ActivityThread.main(ActivityThread.java:7814)
        at java.lang.reflect.Method.invoke(Method.java)
        at com.android.internal.os.RuntimeInit$MethodAndArgsCaller.run(RuntimeInit.java:493)
        at com.android.internal.os.ZygoteInit.main(ZygoteInit.java:1075)

**Layout**

    &lt;?xml version="1.0" encoding="utf-8"?&gt;
    &lt;layout
        xmlns:android="http://schemas.android.com/apk/res/android"
        xmlns:app="http://schemas.android.com/apk/res-auto"&gt;
    
        &lt;data&gt;
    
            &lt;import type="model.CustomerType" /&gt;
    
            &lt;import type="android.view.View" /&gt;
    
            &lt;import type="AccessibilityUtils" /&gt;
    
            &lt;variable
                name="view"
                type="login.LoginView" /&gt;
    
            &lt;variable
                name="presenter"
                type="login.LoginPresenter" /&gt;
        &lt;/data&gt;
    
        &lt;androidx.coordinatorlayout.widget.CoordinatorLayout
            android:layout_width="match_parent"
            android:layout_height="match_parent"&gt;
    
            &lt;com.google.android.material.appbar.AppBarLayout
                android:layout_width="match_parent"
                android:layout_height="wrap_content"
                app:liftOnScroll="?attr/myAppBarLayoutScrollingBehaviorEnabled"&gt;
                &lt;androidx.appcompat.widget.Toolbar
                    android:id="@+id/toolbar"
                    android:layout_width="match_parent"
                    android:layout_height="?android:attr/actionBarSize"
                    app:navigationContentDescription="@string/login__cancel_content_description"
                    app:navigationIcon="?attr/navigationCloseDrawable"&gt;
                    &lt;ImageView
                        android:layout_width="wrap_content"
                        android:layout_height="wrap_content"
                        android:layout_gravity="center"
                        android:contentDescription="@string/login__header_content_description"
                        app:srcCompat="@drawable/logo" /&gt;
                &lt;/androidx.appcompat.widget.Toolbar&gt;
            &lt;/com.google.android.material.appbar.AppBarLayout&gt;
    
            &lt;androidx.core.widget.NestedScrollView
                android:layout_width="match_parent"
                android:layout_height="match_parent"
                app:layout_behavior="?attr/myAppBarLayoutScrollingBehavior"&gt;
    
                &lt;com.google.android.material.card.MaterialCardView
                    style="?attr/myCardViewFullWidthStyle"
                    android:layout_height="wrap_content"
                    android:layout_gravity="center_horizontal"
                    app:contentPadding="0dp"&gt;
    
                    &lt;LinearLayout
                        android:layout_width="match_parent"
                        android:layout_height="wrap_content"
                        android:orientation="vertical"&gt;
    
                        &lt;com.google.android.material.tabs.TabLayout
                            android:id="@+id/tabs"
                            android:layout_width="match_parent"
                            android:layout_height="wrap_content"
                            app:tabGravity="fill"
                            app:tabMode="fixed" /&gt;
    
                        &lt;LinearLayout
                            android:layout_width="match_parent"
                            android:layout_height="wrap_content"
                            android:orientation="vertical"
                            android:padding="@dimen/card_content_padding"&gt;
    
                            &lt;LinearLayout
                                android:layout_width="match_parent"
                                android:layout_height="wrap_content"
                                android:focusable="true"
                                android:orientation="vertical"
                                android:visibility="@{view.customerType == CustomerType.PRIVATE ? View.VISIBLE : View.GONE}"&gt;
    
                                &lt;TextView
                                    android:layout_width="match_parent"
                                    android:layout_height="wrap_content"
                                    android:layout_marginBottom="@dimen/text_vertical_spacing"
                                    android:text="@string/login__title_private"
                                    android:textAppearance="?attr/textAppearanceHeadline5" /&gt;
    
                                &lt;TextView
                                    android:layout_width="match_parent"
                                    android:layout_height="wrap_content"
                                    android:contentDescription="@string/login__body_private_accessibility"
                                    android:text="@{AccessibilityUtils.makeWordAccessible(@string/login__body_private)}" /&gt;
    
                            &lt;/LinearLayout&gt;
    
                            &lt;LinearLayout
                                android:layout_width="match_parent"
                                android:layout_height="wrap_content"
                                android:focusable="true"
                                android:orientation="vertical"
                                android:visibility="@{view.customerType == CustomerType.BUSINESS ? View.VISIBLE : View.GONE}"&gt;
    
                                &lt;TextView
                                    android:layout_width="match_parent"
                                    android:layout_height="wrap_content"
                                    android:layout_marginBottom="@dimen/text_vertical_spacing"
                                    android:text="@string/login__title_business"
                                    android:textAppearance="?attr/textAppearanceHeadline5" /&gt;
    
                                &lt;TextView
                                    android:layout_width="match_parent"
                                    android:layout_height="wrap_content"
                                    android:contentDescription="@string/login__body__accessibility"
                                    android:text="@{AccessibilityUtils.makeWordAccessible(@string/login__body)}" /&gt;
    
                            &lt;/LinearLayout&gt;
    
                            &lt;com.google.android.material.textfield.TextInputLayout
                                android:layout_width="match_parent"
                                android:layout_height="wrap_content"
                                android:layout_marginStart="-4dp"
                                android:layout_marginTop="16dp"
                                android:layout_marginEnd="-4dp"
                                android:layout_marginBottom="@dimen/input_vertical_spacing"
                                android:hint="@string/login__username"&gt;
    
                                &lt;com.google.android.material.textfield.TextInputEditText
                                    android:id="@+id/username"
                                    android:layout_width="match_parent"
                                    android:layout_height="wrap_content"
                                    android:singleLine="true"
                                    android:text="@={view.username}" /&gt;
    
                            &lt;/com.google.android.material.textfield.TextInputLayout&gt;
    
                            &lt;com.google.android.material.textfield.TextInputLayout
                                android:layout_width="match_parent"
                                android:layout_height="wrap_content"
                                android:layout_marginStart="-4dp"
                                android:layout_marginEnd="-4dp"
                                android:hint="@string/login_password_hint"&gt;
    
                                &lt;com.google.android.material.textfield.TextInputEditText
                                    android:id="@+id/password"
                                    android:layout_width="match_parent"
                                    android:layout_height="wrap_content"
                                    android:inputType="textPassword"
                                    android:singleLine="true"
                                    android:text="@={view.password}" /&gt;
                            &lt;/com.google.android.material.textfield.TextInputLayout&gt;
    
                            &lt;com.google.android.material.button.MaterialButton
                                android:id="@+id/login_button"
                                android:layout_width="match_parent"
                                android:layout_height="wrap_content"
                                android:layout_marginTop="@dimen/buttons_margin_top"
                                android:layout_marginBottom="@dimen/button_vertical_spacing"
                                android:onClick="@{() -&gt; presenter.onLoginClicked(view.username, view.password)}"
                                android:text="@string/login_label" /&gt;
    
                            &lt;com.google.android.material.button.MaterialButton
                                android:id="@+id/help_button"
                                style="?attr/buttonTextStyle"
                                android:layout_width="match_parent"
                                android:layout_height="wrap_content"
                                android:onClick="@{() -&gt; presenter.onForgotCredentialsClicked()}"
                                android:text="@string/login_help" /&gt;
                        &lt;/LinearLayout&gt;
                    &lt;/LinearLayout&gt;
                &lt;/com.google.android.material.card.MaterialCardView&gt;
            &lt;/androidx.core.widget.NestedScrollView&gt;
        &lt;/androidx.coordinatorlayout.widget.CoordinatorLayout&gt;
    &lt;/layout&gt;

**Attribute/Style**

    &lt;item name="myCardViewFullWidthStyle"&gt;@style/CardView.FullWidth&lt;/item&gt;
        &lt;style name="CardView.FullWidth" parent="Widget.MaterialComponents.CardView"&gt;
            &lt;item name="cardCornerRadius"&gt;@dimen/card_fullwidth_corner_radius&lt;/item&gt;
            &lt;item name="cardElevation"&gt;@dimen/card_fullwidth_elevation&lt;/item&gt;
            &lt;item name="contentPadding"&gt;@dimen/card_fullwidth_content_padding&lt;/item&gt;
            &lt;item name="android:layout_width"&gt;@dimen/card_fullwidth_width&lt;/item&gt;
            &lt;item name="android:layout_marginTop"&gt;@dimen/card_fullwidth_margin_top&lt;/item&gt;
            &lt;item name="android:layout_marginBottom"&gt;@dimen/card_fullwidth_margin_bottom&lt;/item&gt;
        &lt;/style&gt;

And the resource ID #0x7f040378 that system is unable to find seems to be style="?attr/myCardViewFullWidthStyle" or at least something in related to CardView because that's where the crash is pointing in the layout.

                &lt;com.google.android.material.card.MaterialCardView
                    style="?attr/myCardViewFullWidthStyle"
                    android:layout_height="wrap_content"
                    android:layout_gravity="center_horizontal"
                    app:contentPadding="0dp"&gt;

Hope I've put enough details and it's understable what my goal is. Looking forward for suggestion, ideas or fix!  


Edit:  
It happens on prod app version with proguard.
