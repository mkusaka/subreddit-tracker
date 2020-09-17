# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/isk58n/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/isk58n/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - September 14, 2020](https://www.reddit.com/r/androiddev/comments/isiooc/weekly_questions_thread_september_14_2020/)
- url: https://www.reddit.com/r/androiddev/comments/isiooc/weekly_questions_thread_september_14_2020/
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
## [3][The internals of Lottie Animation Library - Article](https://www.reddit.com/r/androiddev/comments/iufv1z/the_internals_of_lottie_animation_library_article/)
- url: https://codingwithmitch.com/blog/understanding-the-internals-of-lottie-rendering-the-animation-file/
---

## [4][How to use AnimatedVectorDrawable as windowBackground for splash screen?](https://www.reddit.com/r/androiddev/comments/iua8ap/how_to_use_animatedvectordrawable_as/)
- url: https://www.reddit.com/r/androiddev/comments/iua8ap/how_to_use_animatedvectordrawable_as/
---
I usually implement splash screen using theme with `android:windowBackground` set to the app icon. I was wondering if it is possible to use **AnimatedVectorDrawable** as the window background and make it's animation start automatically. I can use the **AnimatedVectorDrawable** as the background but it's animation won't start automatically.
## [5][It takes a lot time to review updating](https://www.reddit.com/r/androiddev/comments/iuezwk/it_takes_a_lot_time_to_review_updating/)
- url: https://i.redd.it/egg8qcnkynn51.jpg
---

## [6][Dependency management for locally built .aar](https://www.reddit.com/r/androiddev/comments/iufrj2/dependency_management_for_locally_built_aar/)
- url: https://www.reddit.com/r/androiddev/comments/iufrj2/dependency_management_for_locally_built_aar/
---
Hello,  


if I build a local .aar file which has some dependencies and distribute it to people, they will have to add the same dependencies in their build.gradle file so the library doesn't crash upon usage.   
Is there any way to ship a list of my dependencies as a .pom file within my custom library alongside with the .aar file, so gradle knows which dependencies the project is using?  
Or is there any other way to solve this nicely, without having people changing their projects to match the dependencies within my library?  


I need this specifically for local usage of the library, not through maven publish plugin
## [7][Does anyone know of a library to remove the background from an image?](https://www.reddit.com/r/androiddev/comments/iuip7a/does_anyone_know_of_a_library_to_remove_the/)
- url: https://www.reddit.com/r/androiddev/comments/iuip7a/does_anyone_know_of_a_library_to_remove_the/
---
Hi,

I am looking for a library to remove the image from a background like in the linked image.

I have found [https://github.com/GabrielBB/Android-CutOut](https://github.com/GabrielBB/Android-CutOut) but cannot use it for my project.

any help would be greatly appreciated.

Thanks in advance

https://preview.redd.it/2soyme1tepn51.jpg?width=285&amp;format=pjpg&amp;auto=webp&amp;s=65656683eddf930d778fe39f0d23c0e229110f39
## [8][Correct way to use dynamic feature](https://www.reddit.com/r/androiddev/comments/iui02g/correct_way_to_use_dynamic_feature/)
- url: https://www.reddit.com/r/androiddev/comments/iui02g/correct_way_to_use_dynamic_feature/
---
Hi everyone,

&amp;#x200B;

i am confused about usage of dynamic feature. I have been working on developing a new application. So I want to separate all of my feature to modules. I can do that with regular modules but as i read document of dynamic feature, it has lot of benefits. it provides these;

&amp;#x200B;

\-Faster build time

\-Enhancing refactoring

\-Improving process of development

&amp;#x200B;

I think that the most important one is faster build time for me. Therefore, i want to make all of my feature module to dynamic.

I have checked plaid example which is used dagger.

&amp;#x200B;

I am confused about Is it right approach to making all of your modules to dynamic also all of them will be install-time.

&amp;#x200B;

What are the disadvantages?

I can write one of disadvantages:

\-If you want to implement Dependency Injection such as Dagger. It would be hard and complex Or you can use Koin which is much more easy and understandable.

\- UI testing is hard to write; if your app has only dynamic feature, your instrumentation test must be isolated with which activity you test.

\- It is not enough to getting unit test coverage of application / or i couldn't find a way for that.

&amp;#x200B;

To sum up, What do you think?
## [9][Gradle excluding specific class](https://www.reddit.com/r/androiddev/comments/iuhwjo/gradle_excluding_specific_class/)
- url: https://www.reddit.com/r/androiddev/comments/iuhwjo/gradle_excluding_specific_class/
---
Hi, i'm pulling my hairs out for this issue.

I'im importing a library and it use the same library i use so some classes create this error:

&amp;#x200B;

Duplicate class okio.-Base64 found in modules jetified-intercom-sdk-base-8.1.0-runtime.jar ([io.intercom.android](https://io.intercom.android):intercom-sdk-base:8.1.0) and jetified-okio-jvm-2.4.3.jar (com.squareup.okio:okio:2.4.3)

Duplicate class okio.-DeflaterSinkExtensions found in modules jetified-intercom-sdk-base-8.1.0-runtime.jar ([io.intercom.android](https://io.intercom.android):intercom-sdk-base:8.1.0) and jetified-okio-jvm-2.4.3.jar (com.squareup.okio:okio:2.4.3)

Duplicate class okio.-DeprecatedOkio found in modules jetified-intercom-sdk-base-8.1.0-runtime.jar ([io.intercom.android](https://io.intercom.android):intercom-sdk-base:8.1.0) and jetified-okio-jvm-2.4.3.jar (com.squareup.okio:okio:2.4.3)

&amp;#x200B;

I tried with exclusion rules but i cannot find a solution. Can anyone help?

&amp;#x200B;

Thanks
## [10][My First Real App](https://www.reddit.com/r/androiddev/comments/itrsko/my_first_real_app/)
- url: https://www.reddit.com/r/androiddev/comments/itrsko/my_first_real_app/
---
Hello everyone, this is my first time writing in this subreddit.

I wrote a pokédex like app in kotlin and I wanted to showcase it and get your advice as everything I did came from self learning.

&amp;#x200B;

[Link to the app](https://github.com/FabZanna/FabulousPokedex)

&amp;#x200B;

I hope to receive many feedbacks!

Thank you for your time
## [11][is there any book for intermediate to advance android development](https://www.reddit.com/r/androiddev/comments/iugohj/is_there_any_book_for_intermediate_to_advance/)
- url: https://www.reddit.com/r/androiddev/comments/iugohj/is_there_any_book_for_intermediate_to_advance/
---
 i know the basics of android development like (layouts, fetching data from website to app, fragments, list views, recycler view, and a lot more ) and I have a decent hold on java. I completed the head first android book and also completed free basic course of udacity and google on android development but now I wanna go from intermediate to advance. So is there any book for that?
## [12][Moving to Kotlin, how?](https://www.reddit.com/r/androiddev/comments/iuadxq/moving_to_kotlin_how/)
- url: https://www.reddit.com/r/androiddev/comments/iuadxq/moving_to_kotlin_how/
---
Hi everyone, i'm been coding in Java for a long time (around 8 years) and i think the change to Kotlin for Android Development is inevitable. 
My biggest problem is how to start the change, theres a lot of courses and YouTube videos but Is a little overwhelming.
Any tip for the most experienced in this?
