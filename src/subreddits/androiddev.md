# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/ercqvf/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/ercqvf/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - January 20, 2020](https://www.reddit.com/r/androiddev/comments/erbsa6/weekly_questions_thread_january_20_2020/)
- url: https://www.reddit.com/r/androiddev/comments/erbsa6/weekly_questions_thread_january_20_2020/
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
## [3][A couple of questions regarding Clean + MVVM (big project rewrite)](https://www.reddit.com/r/androiddev/comments/ers0kp/a_couple_of_questions_regarding_clean_mvvm_big/)
- url: https://www.reddit.com/r/androiddev/comments/ers0kp/a_couple_of_questions_regarding_clean_mvvm_big/
---
I'm planning to rewrite my app into a new architecture to make it scalable and testable. Currently, the majority of the code is in views and different manager classes because it's a 5+ year-old project and it's simply becoming too much to handle. It works well and the code is manageable, but this is a project that should last another 5+ years, so I need to make some decisions about scalability.

I've read as much info about Clean + MVVM as I possibly could and even though I still don't feel like I understand it completely, I gave it a try and wanted to start rewriting it. I ran into some issues almost immediately. Now I need to decide between blindly following Clean architecture, come up with my own architecture which works best for our project, or make some sacrifices which will go against the main Clean principles.

Here are the main points I want to discuss:

* Our API is made to support our mobile client only, meaning all of the models are almost perfectly ready to be served to the UI. At this point, it simply doesn't make sense to have mappers. We don't cache anything other than "User" object for obvious purposes. Everything else is live data straight from API. App doesn't work without internet connection and it never will. However I still need GSON because I want to rename some (although rare) parameters when mapping from API response. This means I would need to use GSON dependency in the Domain layer. Is this acceptable?
* I am running into the same issue with LiveData. I googled this issue and I noticed I'm not the only one with that problem. I want to have Observable pattern for streaming Data right off the bat from the Domain level. People are recommending to use Rx observable for domain and data layer, and only convert it into LiveData on the Presentation layer. This just seems like overhead work to me. I don't see how I would "break" the architecture if I simply decided to add LiveData as a dependency on Domain layer.
* As far as I understand, UseCases are written on Domain layer. After all, they simply use Repositories to get data, handle it somehow and then pass it on to Presenters. But Dagger (DI) is usually on Presentation layer. How do I inject UseCases with Repositories in this case? This is one thing I seemed to have missed when researching.
* How "specific" should use cases be? Like I said, most of our API calls (&gt;90%) already serve the data exactly as we need it to, so the vast majority of our UseCases would be "useless". Now I know this is a common question when working with UseCases for the first time, but the vast majority of our use cases really would be simply calling the repository method. I can only think of a couple of instances where a UseCase would take 2 or maybe 3 repositories, make 3 separate calls and then serve the data to UI. (for example, when displaying our user profile, we make one call to GET user profile and then another one to GET some additional user info - this "ShowUserProfileUseCase" would simply make these 2 calls consecutively and then serve all data at once). Having said this, is it acceptable to ditch UseCases in my case?
* Last point I want to discuss is simply what's the best approach to start rewriting architecture on a big project? Like I said, this is a 5+ years old (huge) project. Current architecture is simple. Network calls are grouped in "client" classes which access a singleton Retrofit instance. "Clients" are initialized directly on Activities/Fragments and once the network calls are made, the data is displayed on UI. It works, there are no bugs or performance issues. Considering the "old school" architecture, the project works great. Even the code is more scalable than it sounds. But it isn't testable at all, which is the main reason I want to do a rewrite. But it's hard to even start a rewrite on a project this big.
## [4][Kotlin typealias: use alternate names for class or functions types without introducing a new type](https://www.reddit.com/r/androiddev/comments/ern95v/kotlin_typealias_use_alternate_names_for_class_or/)
- url: https://medium.com/androiddevelopers/alter-type-with-typealias-4c03302fbe43
---

## [5][Concat XML strings at build time](https://www.reddit.com/r/androiddev/comments/ersk7b/concat_xml_strings_at_build_time/)
- url: https://www.reddit.com/r/androiddev/comments/ersk7b/concat_xml_strings_at_build_time/
---
[https://github.com/LikeTheSalad/android-string-reference](https://github.com/LikeTheSalad/android-string-reference)
## [6][Using Google drive api on F-Droid app?](https://www.reddit.com/r/androiddev/comments/erq475/using_google_drive_api_on_fdroid_app/)
- url: https://www.reddit.com/r/androiddev/comments/erq475/using_google_drive_api_on_fdroid_app/
---
I have created an app "Notes Sync". The app is open source on github. Many users have asked me to make the app available on F-Droid. I looked through the rules and found that google services isn't supported on F-Droid apps. My app uses Google OAuth and Google Drive API for saving notes and images to the user's Google Drive storage. So is it okay to submit my app to F-Droid? If not is there an alternative to the Google Drive API?
## [7][Super Simple Bouncy RecyclerView in Kotlin](https://www.reddit.com/r/androiddev/comments/erth62/super_simple_bouncy_recyclerview_in_kotlin/)
- url: https://www.reddit.com/r/androiddev/comments/erth62/super_simple_bouncy_recyclerview_in_kotlin/
---
[https://github.com/IainS1986/SimpleBouncyRecyclerView](https://github.com/IainS1986/SimpleBouncyRecyclerView)

To be honest, I did this project mainly to learn around producing Android aar libraries and how to share them (via Bintray -&gt; JCenter), but thought some folks might be interested in seeing what I did.

I'm a professional Xamarin developer by day but shifting to native development on the side. I did this in Xamarin originally so ported it over as a test.

Its called "Simple" really because I 100% know there are much better solutions out there that are probably much safer (especially with regards to animating cells on add/remove - that might not work while over-scrolling :/ ). This is also completely \*not\* physics based, its just faked via animation and a simple dampening function on drag. But, some people might be interested in the code and the approach I took regardless.
## [8][How do you Test your Android apps mostly Android bases classes like Activity,Fragment with Robolectric or Mockito/JUnit ?](https://www.reddit.com/r/androiddev/comments/erjbc0/how_do_you_test_your_android_apps_mostly_android/)
- url: https://www.reddit.com/r/androiddev/comments/erjbc0/how_do_you_test_your_android_apps_mostly_android/
---

## [9][Any sources for working with Tensorflow Lite on Android Studio ?](https://www.reddit.com/r/androiddev/comments/erp9p9/any_sources_for_working_with_tensorflow_lite_on/)
- url: https://www.reddit.com/r/androiddev/comments/erp9p9/any_sources_for_working_with_tensorflow_lite_on/
---
Hey guys a newbie here. Recently I've been engaged in a project where I have to classify Bluetooth Signals though a machine learning model. A SVM classifier to be exact. And the python model is done but I got no clue how to run the model in Android and I failed to find a goos resource. Any ideas or good sources ?
## [10][Room: "top star projection"](https://www.reddit.com/r/androiddev/comments/eru31c/room_top_star_projection/)
- url: https://www.reddit.com/r/androiddev/comments/eru31c/room_top_star_projection/
---
Does anyone know what does the word "top" mean in context of "top star projection"?

I've found it here: [https://developer.android.com/jetpack/androidx/releases/room#compiler-options](https://developer.android.com/jetpack/androidx/releases/room#compiler-options)

&gt;room.expandProjection: Configures Room to rewrite queries such that their ***top*** star projection is expanded to only contain the columns defined in the DAO method return type.

More details about Expanding Projections are here: [https://developer.android.com/jetpack/androidx/releases/room#2.2.0](https://developer.android.com/jetpack/androidx/releases/room#2.2.0)

&gt;**Expanding Projections:** A new experimental compiler option room.expandProjection was added that causes Room to rewrite a query with a star projection to only contain the columns in the returning type POJO. For example, for a DAO method with @'Query("SELECT \* FROM Song") that returns a POJO named SongIdAndTitle with only two fields. Then Room will rewrite the query to SELECT id, title FROM Song such that the minimum set of columns to satisfy the return type are fetched. This essentially eliminates the CURSOR\_MISMATCH warning that is presented when the query returns extra columns that do not match any field in the returning POJO type.

but they don't mention word "top". Do you have any clue?
## [11][Dependency Injection in a library project](https://www.reddit.com/r/androiddev/comments/erre46/dependency_injection_in_a_library_project/)
- url: https://www.reddit.com/r/androiddev/comments/erre46/dependency_injection_in_a_library_project/
---
I've been tasked with creating a few libraries that will be used by Android devs within the organisation. One of these libraries is a wrapper for AppAuth with some extra functionality. 

I wish to know how to have a DI framework within these libraries without the need for the consumer of the library to have that framework as well. For example with Koin, the app using the library would have to perform \`startKoin { ... } \` in their Application class in order to use it. Would it be the same if I used Dagger? I'm not too experienced with developing libraries and I don't wish for the library setup to be too cumbersome for the user but at the same time I need DI within the libraries.
## [12][Videoview doubt](https://www.reddit.com/r/androiddev/comments/ert964/videoview_doubt/)
- url: https://www.reddit.com/r/androiddev/comments/ert964/videoview_doubt/
---
I have two different videoview playing at the same time now.
So right now i have two different media controller that pops up depending on which videoview is in focus. I want to combine these two video controllers into one so two separate controllers doesn't pop up.
