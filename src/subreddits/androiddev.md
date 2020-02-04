# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/ey6mqi/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/ey6mqi/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
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
## [3][Karumi Android 2020 development stack](https://www.reddit.com/r/androiddev/comments/eypany/karumi_android_2020_development_stack/)
- url: https://blog.karumi.com/our-android-2020-development-stack/
---

## [4][Material Components 1.1.0 stable version released](https://www.reddit.com/r/androiddev/comments/eyb7qp/material_components_110_stable_version_released/)
- url: https://github.com/material-components/material-components-android/releases/tag/1.1.0
---

## [5][2020 Java Technology Report](https://www.reddit.com/r/androiddev/comments/eynjf3/2020_java_technology_report/)
- url: https://www.jrebel.com/blog/2020-java-technology-report
---

## [6][Examples for multi-module apps with one module per feature](https://www.reddit.com/r/androiddev/comments/eynjkq/examples_for_multimodule_apps_with_one_module_per/)
- url: https://www.reddit.com/r/androiddev/comments/eynjkq/examples_for_multimodule_apps_with_one_module_per/
---
Recently, I have [often](https://engineering.tiki.vn/o-1-android-build-time-at-tiki-1a27a804bb6a) [read](https://github.com/jeziellago/FlowNav) [about apps](https://medium.com/@vlastobrecka/dagger-in-a-multi-module-project-without-dagger-android-5d9cc5ae981c) with multiple modules. Often in the context of decreasing build times but also when it came to better collaboration and clean(er) code.

What are examples of apps that incorporate this multi-module structure, preferrably open source? I’d like to learn more about how to do that on scale.
## [7][Is it possible to implement Multiple gestures or tranistions on a same view in MotionScene of a MotionLayout?](https://www.reddit.com/r/androiddev/comments/eymgdy/is_it_possible_to_implement_multiple_gestures_or/)
- url: https://www.reddit.com/r/androiddev/comments/eymgdy/is_it_possible_to_implement_multiple_gestures_or/
---
Here is the thing that I'm trying to do:

https://stackoverflow.com/questions/60040018/android-motionlayout-how-to-handle-both-click-and-swipe-events-on-same-view-v

I need this for my app, can someone please tell me if it's possible to do it with motion layout, or should I find alternative solution? If so can you give me suggestions.
## [8][GC overhead limit exceeded](https://www.reddit.com/r/androiddev/comments/eypy3x/gc_overhead_limit_exceeded/)
- url: https://www.reddit.com/r/androiddev/comments/eypy3x/gc_overhead_limit_exceeded/
---
Recently I have assigned to a project that is already started and published on the playstore and from time to time while working on it an exception arise when I build the application: 

**e: java.lang.OutOfMemoryError: GC overhead limit exceeded**

before saying google it, I have already googled it and fixed it. But now I want to know why this happens for this project specifically?
## [9][How can i make an APP like this?](https://www.reddit.com/r/androiddev/comments/eypw2u/how_can_i_make_an_app_like_this/)
- url: https://www.reddit.com/r/androiddev/comments/eypw2u/how_can_i_make_an_app_like_this/
---
 Hello guys, first of all sorry my bad English.

My coding knowledge is very limited and i need your help.

I  would like to make an iOS / Android app that allows to the customer  choose and buy a Service Plan,then allow them to send me private  coordinates with Google Maps to go personally (physically) and offer the  service,**there's any company that offers a app maker with this purchase system?**

Thanks in advance guys!
## [10][Do you usually construct ViewModelProvider by using Activity or Fragment?](https://www.reddit.com/r/androiddev/comments/eypkgn/do_you_usually_construct_viewmodelprovider_by/)
- url: https://www.reddit.com/r/androiddev/comments/eypkgn/do_you_usually_construct_viewmodelprovider_by/
---
It is common to have a (or multiple) Fragment(s) within single Activity.

I was wondering, usually, when you construct ViewModelProvider within a Fragment, do you pass in Activity, or Fragment into ViewModelProvider's constructor?

&amp;#x200B;

I was wondering, for the following 2 code snippets, which is the recommended way?

## Use Fragment when constructing ViewModelProvider

    public class MyFragment extends Fragment {     
        @Override     
        public void onCreate(Bundle savedInstanceState) {         
        super.onCreate(savedInstanceState);          
           final ViewModelProvider viewModelProvider = new ViewModelProvider(this);         
        tabInfoViewModel = viewModelProvider.get(TabInfoViewModel.class);     
        } 

## Use Activity when constructing ViewModelProvider

    public class MyFragment extends Fragment {     
        @Override     
        public void onCreate(Bundle savedInstanceState) {         
        super.onCreate(savedInstanceState);          
           final ViewModelProvider viewModelProvider = new ViewModelProvider(getActivity());         
        tabInfoViewModel = viewModelProvider.get(TabInfoViewModel.class);     
        } 

Both code are workable. But, which is the "more correct" way?
## [11][What android studio plugins do you use to help coding a bit faster?](https://www.reddit.com/r/androiddev/comments/eypcxl/what_android_studio_plugins_do_you_use_to_help/)
- url: https://www.reddit.com/r/androiddev/comments/eypcxl/what_android_studio_plugins_do_you_use_to_help/
---

## [12][Complex UI/Animations on Android — featuring MotionLayout](https://www.reddit.com/r/androiddev/comments/ey5lu3/complex_uianimations_on_android_featuring/)
- url: https://medium.com/@nikhilpanju22/complex-ui-animations-on-android-featuring-motionlayout-aa82d83b8660?source=friends_link&amp;sk=5b924ea26bc2ae4735483760f3c62409
---

