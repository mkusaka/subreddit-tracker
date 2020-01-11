# androiddev
## [1][App Feedback Thread - January 11, 2020](https://www.reddit.com/r/androiddev/comments/en6v2r/app_feedback_thread_january_11_2020/)
- url: https://www.reddit.com/r/androiddev/comments/en6v2r/app_feedback_thread_january_11_2020/
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
## [2][Weekly Questions Thread - January 06, 2020](https://www.reddit.com/r/androiddev/comments/ekslrd/weekly_questions_thread_january_06_2020/)
- url: https://www.reddit.com/r/androiddev/comments/ekslrd/weekly_questions_thread_january_06_2020/
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
## [3][[Possible rant] Why Android development seems like a mess?](https://www.reddit.com/r/androiddev/comments/en5of4/possible_rant_why_android_development_seems_like/)
- url: https://www.reddit.com/r/androiddev/comments/en5of4/possible_rant_why_android_development_seems_like/
---
New here. And new to Android development. Developing in .NET is my job.

So I developed so far 3 Android apps by myself, and publish 2 of them to Play Store.

I just finished one's latest version and published, and I feel I want to share some of my experience on my Android dev journey. Maybe some rant and misunderstanding here, I apologize for this in advance (and please comment).

Simply put, I feel Android \*seems\* like a mess. Maybe I am still not proficient enough to get everything right at once.

\- For very simple apps or scenarios, I think the default setup of Android IDE &amp; SDK, as well as docs, is all right. But once I want to use more system features, or apply more designs, Android seems to be really hard or complicated to get everything right. The official Android dev docs seem good for learning, but if you have some special scenarios, it seems you are on your own.

(For example, if you want to send network / HTTP request, the official docs recommended using Volley / Cronet libraries, which are quite complicated to use if I want to achieve something like JSON-Object serialization (maybe I didn't get it right). This makes me wonder "Why the system itself does not have built-in network / HTTP API like Windows or Linux?". Not to mention, Retrofit+OkHttp is (IMHO) way much easier than either of them. Why this was not mentioned in docs?)

\- Support / Compat libraries / classes. OH MY GOD. Maybe I don't fully know the logic behind this, but I hate these things. Every time when I try to use something, like, Activity or Fragment, I have to figure out whether I should use support libraries or system native ones. I know usually I should use support libraries. The problem is, some classes from libraries (maybe system also? I cannot remember) derived from the native ones. In some scenarios, if I want to use support library ones, I will get errors. I don't really know anymore.

\- Outside Android system territory, it surprises me a lot that a lot of libraries are lack of proper documentation. I know this happens everywhere, not only in Android, but one thing I feel about Android is the lack of "conceptual consistency".

Allow me to present an example. When you are working on Windows, say, .NET, even though there are many differences in .NET BCL (Base Class Library) and 3rdparty libraries or even BCL itself, you still can have some expectation on how they approach the objective, since all .NET stuff has more or less a unified "feel" on design rules, so that you will not get lost very easy. Linux is kinda similar (but maybe this has my own factors - I did a bunch of Linux system programming).

But in Android, system does not provide a lot of stuff (either non-exist or simply broken), so a lot of 3rdparty libraries took over the roles what I think should be a part of system (like high-level network stuff). These libraries have their own (and sometimes hacky) ways to do things. But they sometimes don't provide good documentation to explain them, even for the new features. Examples such as:

\- Retrofit2's Kotlin suspend function feature (what I can do, and what I cannot do, can I return a list instead of one object, or such)

\- Dagger's DaggerApplication, DaggerAppCompatActivity, DaggerFragment, etc., and how to do ViewModels properly (I didn't know these classes until searched through tons of results online)

\- Epoxy's generated models and integration with Android ViewModels ("I see red in Android Studio, it says Unknown symbol! What should I do???" / "why I cannot compile? - oh, different class name in [package-info.java](https://package-info.java)")

\- Gson's ProGuard rules (Where is it? - Production build got my model classes stripped out and caused Retrofit stop working properly)

\- The integration among all these libraries ("How can I use Epoxy + ViewModels + Dagger altogether???")

And it is quite insane that I have to use all these libraries in an Android to achieve "maintainability". Why Android doesn't have a built-in way to do these? Why Android Jetpack???

\- Maybe I'm really unlucky. But for some reason, every time when I try to figure out how to solve a particular problem in StackOverflow, I always bumped into a page that has either no answers (10%), or has a bunch of answers without that green mark (70%, the rest 20% are jackpot). I thought tons of people use this System, so statistically most of the problems should already have answers. But nope.

(Not to mention tons of pages have outdated information. Gosh)

(...and because of the rapid changes on all Android stuff, it is impossible to find a friendly beginner book talking all these things I mentioned above - I like to read a book to get my hands dirty, and I think books are more trustworthy)

\- Tons of blogs. No offense. I just never see this kind of things before. A lot of developers from all over the world write blogs in Medium or such to explain the same things in many different ways.

For example, Dagger. To be fair, Dagger is complicated because it tries to solve a complicated problem (dependency injection). But man, I read a lot, I wrote a lot, but when I first read these blogs, they just gave me tons of confusion (especially when someone uses outdated information). Even Google's official docs are confusing as hell (Thermosiphon? REALLY?) Where is the really official documentation???

\- System bugs. This really shocks me. Like, a crashed NotificationListenerService must be restarted by restarting my device. Or, broken EditTextPreference not being able to limit inputs to numbers even if I set it so.

\- Flutter / Fuchsia. I don't know whether Google is seriously considering using these to solve Android's issues...

Well, not quite pleasant. I guess maybe because I am still a beginner. I may misunderstand most of things here.

I can only hope I will get better on these things and not be frustrated easily like now.
## [4][Why Microsoft use React Native over Xamarin](https://www.reddit.com/r/androiddev/comments/en1onu/why_microsoft_use_react_native_over_xamarin/)
- url: https://www.reddit.com/r/androiddev/comments/en1onu/why_microsoft_use_react_native_over_xamarin/
---
I've seened a twitter post recently about the use React Native for various Mirosoft products such as Teams, Skype, Outlook etc.

My question is why they don't used Xamarin.Forms at all to build a cross platform app since they invested in it and bought the company instead they use RN. I'm just wondering if Xamarin in the near feature would die?
## [5][What happened to android-arsenal.com ?](https://www.reddit.com/r/androiddev/comments/en4q3d/what_happened_to_androidarsenalcom/)
- url: https://www.reddit.com/r/androiddev/comments/en4q3d/what_happened_to_androidarsenalcom/
---

## [6][Constant deprecated even before being added](https://www.reddit.com/r/androiddev/comments/en6v3i/constant_deprecated_even_before_being_added/)
- url: https://www.reddit.com/r/androiddev/comments/en6v3i/constant_deprecated_even_before_being_added/
---
Yeah, this sums up Android development pretty well. A constant has been deprecated even before added. 

https://preview.redd.it/xbc6mxwz75a41.png?width=879&amp;format=png&amp;auto=webp&amp;s=85def75c2589b53a5c189d9f0962b5c5ab2b64de

It is not fake, you can find it at [https://developer.android.com/reference/android/webkit/WebSettings.LayoutAlgorithm#TEXT\_AUTOSIZING](https://developer.android.com/reference/android/webkit/WebSettings.LayoutAlgorithm#TEXT_AUTOSIZING)
## [7][New: List of KTX extensions](https://www.reddit.com/r/androiddev/comments/emr8sj/new_list_of_ktx_extensions/)
- url: https://developer.android.com/kotlin/ktx/extensions-list
---

## [8][Shared Element transition return?](https://www.reddit.com/r/androiddev/comments/en6l9v/shared_element_transition_return/)
- url: https://streamable.com/r5gww
---

## [9][Logo maker for your app?](https://www.reddit.com/r/androiddev/comments/en1kun/logo_maker_for_your_app/)
- url: https://www.reddit.com/r/androiddev/comments/en1kun/logo_maker_for_your_app/
---
I was recently looking for a simple enough logo for my app and found out about Shopify's App:

 [https://hatchful.shopify.com/](https://hatchful.shopify.com/)   


I have no success with the Hatchful Mobile App, but the Web App works okay :)
## [10][microMVCompose - MVC+navigation library for Jetpack Compose](https://www.reddit.com/r/androiddev/comments/en5js0/micromvcompose_mvcnavigation_library_for_jetpack/)
- url: https://www.reddit.com/r/androiddev/comments/en5js0/micromvcompose_mvcnavigation_library_for_jetpack/
---
Hello, I made this library: [microMVCompose](https://github.com/andob/microMVCompose). It's a MVC implementation + simple navigation system for JetpackCompose. What do you guys think? :)
## [11][How Do request only While In Use Permission? App is currently asking for All The Time Background permission.](https://www.reddit.com/r/androiddev/comments/en2awi/how_do_request_only_while_in_use_permission_app/)
- url: https://www.reddit.com/r/androiddev/comments/en2awi/how_do_request_only_while_in_use_permission_app/
---
&amp;#x200B;

[How do you\* my bad](https://preview.redd.it/8atazxxem2a41.png?width=1080&amp;format=png&amp;auto=webp&amp;s=94769a22435b144d95f72e83e952ea389a1a7c98)
## [12][App published on play store, not available for download](https://www.reddit.com/r/androiddev/comments/en5fyq/app_published_on_play_store_not_available_for/)
- url: https://www.reddit.com/r/androiddev/comments/en5fyq/app_published_on_play_store_not_available_for/
---
Hi dear Reddit community!
As a freelancer I have developed an Android/iOS app for my client (React Native). About a month ago I finished the project and my client was very happy with it, so he decided to publish it on the play store. He created his account and uploaded the aab. Now, after a month, I can see the app on the play store, but it is not available for download. There's written: "Coming soon. Pre-register". What can I do?

Thanks
