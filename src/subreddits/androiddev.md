# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/g0i9da/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/g0i9da/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - April 13, 2020](https://www.reddit.com/r/androiddev/comments/g0grat/weekly_questions_thread_april_13_2020/)
- url: https://www.reddit.com/r/androiddev/comments/g0grat/weekly_questions_thread_april_13_2020/
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
## [3][FastLayout: An android library for making row, column interface quickly](https://www.reddit.com/r/androiddev/comments/g297il/fastlayout_an_android_library_for_making_row/)
- url: https://www.reddit.com/r/androiddev/comments/g297il/fastlayout_an_android_library_for_making_row/
---
I was making a suduko game on android and had lots of issues in making suduko buttons.So I made a library with which you can quickly build any tabular/Table layout with equal cells.Here is a [link](https://github.com/alihamuh/fastLayout) to my github repo.I would very much appreciate your comments or if you think the library has any issues or needs improvement I am open to suggestions.This is my first ever library in any language.
## [4][Venom - a lightweight tool that simplifies testing of the process death scenario.](https://www.reddit.com/r/androiddev/comments/g2cdid/venom_a_lightweight_tool_that_simplifies_testing/)
- url: https://github.com/YarikSOffice/venom
---

## [5][Jetpack Compose 0.1.0-dev09 is released](https://www.reddit.com/r/androiddev/comments/g2ajja/jetpack_compose_010dev09_is_released/)
- url: https://www.reddit.com/r/androiddev/comments/g2ajja/jetpack_compose_010dev09_is_released/
---
Release notes:

* Jetpack Compose: [https://developer.android.com/jetpack/androidx/releases/compose#0.1.0-dev09](https://developer.android.com/jetpack/androidx/releases/compose#0.1.0-dev09)
* Compose UI: [https://developer.android.com/jetpack/androidx/releases/ui#0.1.0-dev09](https://developer.android.com/jetpack/androidx/releases/ui#0.1.0-dev09)

Other than these, the most interesting stuff is that they also added RxJava and LiveData adapters for Compose. Here are some examples

* [LiveData usage example](https://android-review.googlesource.com/c/platform/frameworks/support/+/1278632/11/ui/ui-livedata/samples/src/main/java/androidx/ui/livedata/samples/Samples.kt)
* [RxJava example](https://android-review.googlesource.com/c/platform/frameworks/support/+/1279164/8/ui/ui-rxjava2/samples/src/main/java/androidx/ui/rxjava2/samples/Samples.kt)

There is an open issues for adding Kotlin coroutine Flow support as well so I suppose we can expect an adapter for that soon.
## [6][Why more than two android flavor dimensions don't work properly?](https://www.reddit.com/r/androiddev/comments/g2eay6/why_more_than_two_android_flavor_dimensions_dont/)
- url: https://www.reddit.com/r/androiddev/comments/g2eay6/why_more_than_two_android_flavor_dimensions_dont/
---
There is a powerful mechanism in Gradle for android that allows the developer to create application flavors and split code, configurations and override resources. I use it in more than two years in different applications, but all that time I had only one or two flavor dimensions. Today I tried t add the third dimension to my project and still can't build it.

  
Here is a problem, for example, I need the next three dimensions for my app: "brand", "environment", "market". The configuration will look like this: 

    flavorDimensions "brand", "environment", "market"
    
    productFlavors {
        // brand
        nike{
            dimension "brand"
            applicationId "com.android.nike"
        }
        adidas{
            dimension "brand"
            applicationId "com.android.adidas"
        }
    
        // environment
        test{
            dimension "environment"
            applicationIdSuffix ".test"
        }
        prod {
            dimension "environment"
        }
    
        // market
        google {
            dimension "market"
        }
        amazon {
            dimension "services"
        }
    }

As a result, I can build the next 8 applications (I ignoring build types in this example):

    nikeTestGoogle
    nikeTestAmazon
    nikeProdGoogle
    nikeProdAmazon
    adidasTestGoogle
    adidasTestAmazon
    adidasProdGoogle
    adidasProdAmazon

As I understand documentation about flavor dimensions, application nikeTestGoogle use files from the next folders: main, nike, test, google, nikeTest, nikeTestGoogle. But it doesn't work, Android studio sees only files in main, nike, test, google, and nikeTestGoogle. Files in nikeTest ignored and even doesn't mark as currently active.

Does anyone know why it happened? Am I missing something?

P/S/ Sorry for my bad English, I suppose you understand my main thought
## [7][Coroutines: First things first](https://www.reddit.com/r/androiddev/comments/g2bxfl/coroutines_first_things_first/)
- url: https://medium.com/androiddevelopers/coroutines-first-things-first-e6187bf3bb21
---

## [8][Multidexing, methods counts, is it still a thing?](https://www.reddit.com/r/androiddev/comments/g2eo7z/multidexing_methods_counts_is_it_still_a_thing/)
- url: https://www.reddit.com/r/androiddev/comments/g2eo7z/multidexing_methods_counts_is_it_still_a_thing/
---
I've been out of touch on multidexing lately, is it still a thing for minSdk above 21? How about methods counting, are we still counting, if yes, how?
## [9][Refactoring an old app](https://www.reddit.com/r/androiddev/comments/g2eepq/refactoring_an_old_app/)
- url: https://www.reddit.com/r/androiddev/comments/g2eepq/refactoring_an_old_app/
---
I am wondering if there is any solid advice for refactoring android apps. My company inherited a project and the clients app is very very buggy and unstable. The code is in java/MVP - but they didn't use the presenters to do anything other than make HTTP requests and hand that data to the activity/fragment. The activity/fragments are gigantic and have tons of business logic in them. The application class is also bigger than any of the ones I have written with tons of services instantiated and managed there. Dagger present - but not injecting anything. A lot of libraries in the gradle file that aren't being used. Null Pointer Exceptions everywhere due to life cycle stuff. I can't really say more than that - but where do you start for this type of refactor where throwing everything out is not an option.
## [10][Android Jetpack: CameraX Beta](https://www.reddit.com/r/androiddev/comments/g1v2kl/android_jetpack_camerax_beta/)
- url: https://www.youtube.com/watch?v=HGSj-JRpuxo
---

## [11][Android API version where app shortcuts are required to be supported by launchers?](https://www.reddit.com/r/androiddev/comments/g2exgz/android_api_version_where_app_shortcuts_are/)
- url: https://www.reddit.com/r/androiddev/comments/g2exgz/android_api_version_where_app_shortcuts_are/
---
I am developing an app where its main feature is to create pinned shortcuts.

However, not all launchers support shortcuts. The problem is that since I cannot know this before the user installs the app, some users will get an error message that "Shortcuts are not supported by your launcher".

Is there an API version (e.g. Android 10 - API 29), where all launchers are required to support shortcuts?
## [12][I have an app idea that solves a common local problem, but I can't code or design. Any ideas? Also discuss what else could be in this app.](https://www.reddit.com/r/androiddev/comments/g2eop1/i_have_an_app_idea_that_solves_a_common_local/)
- url: https://www.reddit.com/r/androiddev/comments/g2eop1/i_have_an_app_idea_that_solves_a_common_local/
---
So my app idea is dependent on Google Maps, and some math.

So, here's the concept. There is a friend group of X people, and they want to minimize the travel time to go to a gathering. So this app would choose a venue using specifications by the user (cost, type of venue) and mathematical optimization techniques (that I have yet to fully learn)

It will be a while before I learn these, so I wonder what I should do for now. Should I have this patented or should I try finding people to help in making this idea into a reality?
