# androiddev
## [1][App Feedback Thread - September 05, 2020](https://www.reddit.com/r/androiddev/comments/in0ang/app_feedback_thread_september_05_2020/)
- url: https://www.reddit.com/r/androiddev/comments/in0ang/app_feedback_thread_september_05_2020/
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
## [2][Weekly Questions Thread - August 31, 2020](https://www.reddit.com/r/androiddev/comments/ijvutk/weekly_questions_thread_august_31_2020/)
- url: https://www.reddit.com/r/androiddev/comments/ijvutk/weekly_questions_thread_august_31_2020/
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
## [3][How suitable is Android development for a woman with disabilities?](https://www.reddit.com/r/androiddev/comments/ine828/how_suitable_is_android_development_for_a_woman/)
- url: https://www.reddit.com/r/androiddev/comments/ine828/how_suitable_is_android_development_for_a_woman/
---
This will probably get buried, but I don't have anything to lose. :)

My friend has disabilities that require a lot of time and effort to manage. She hasn't been able to work for many years and lives a meagre existence.

Recently, she has expressed an interest in developing apps for Android (she has a basic smartphone). She's very underconfident and has no past experience in computers but can work hard and is quite clever.

When I mentioned giving it a try she looked sad and said that no one would ever give a job to someone like her with her health issues. I'm not an expert in this area by any means, so I decided to put my feelers out here.  Is there any encouragement I can offer her? Would she never get hired? Even part time? She's 47 if that matters.

I appreciate the time if you read this.
## [4][Social Network for Mars Rovers](https://www.reddit.com/r/androiddev/comments/inle62/social_network_for_mars_rovers/)
- url: https://www.reddit.com/r/androiddev/comments/inle62/social_network_for_mars_rovers/
---
Hi, I develop a sample app created with Nasa API that shows photos taken on Mars by the rovers: Curiosity, Opportunity and Spirit, [here is the result](https://github.com/r-apostol-9618/Mars_Rover_SocialNetwork)  


I used coroutines with retrofit2 to make the calls, I'm kinda new to kotlin and git.  


Any feedback is appreciate!
## [5][Duofolio: An open-source ebook reader that helps you learn new languages](https://www.reddit.com/r/androiddev/comments/inl6xs/duofolio_an_opensource_ebook_reader_that_helps/)
- url: https://i.redd.it/o8pbutrptil51.png
---

## [6][Does the model in MVC/MVP/MVVM depend on the presentation component? The short answer is No.](https://www.reddit.com/r/androiddev/comments/in7oam/does_the_model_in_mvcmvpmvvm_depend_on_the/)
- url: https://medium.com/@johnleeroy/clean-architecture-approach-to-looking-at-the-model-2020de376130?source=friends_link&amp;sk=e96e70db60aa497cf08543313677f4ac
---

## [7][Sample with Dynamic Feature Module fragments as base of BottomNavigationView and ViewPager2 with Dagger Hilt, offline first/last with custom Flow and LiveData TestObservers](https://www.reddit.com/r/androiddev/comments/inlqbv/sample_with_dynamic_feature_module_fragments_as/)
- url: https://www.reddit.com/r/androiddev/comments/inlqbv/sample_with_dynamic_feature_module_fragments_as/
---
Working on this sample app, looking for a job and to show as a showcase app as a tech demo, which includes Dynamic Feature Module fragments as base of ViewPager2, BottomNavigation with Dagger Hilt using both  RxJava3 and Flow with both offline-first and last approach, Room one to one and migrations , and unit tests with TDD approach.

Also created custom FlowTestObserver and LiveDataTestObserver to test with RxJava style observers

such as

                    // WHEN
                    val testObserver = useCase.getPostFlowOfflineLast().test(this)
    
                    // THEN
                    testObserver
                        .assertComplete()
                        .assertNoErrors()
                        .assertValueAt(0, postList)
                        .assertValues {
                            it.first().containsAll(postList)
                        }
                        .dispose()

for a function like

        fun getPostFlowOfflineLast(): Flow&lt;List&lt;Post&gt;&gt; {
            return flow { emit(repository.fetchEntitiesFromRemote()) }
                .map {
                    if (it.isNullOrEmpty()) {
                        throw EmptyDataException("No Data is available in Remote source!")
                    } else {
                        repository.run {
                            deletePostEntities()
                            savePostEntities(it)
                            getPostEntitiesFromLocal()
                        }
                    }
                }
                .flowOn(dispatcherProvider.ioDispatcher)
                // This is where remote exception or least likely db exceptions are caught
                .catch { throwable -&gt;
                    emitAll(flowOf(repository.getPostEntitiesFromLocal()))
                }
                .map {
                    toPostListOrError(it)
                }

&amp;#x200B;

[https://github.com/SmartToolFactory/Posts-MVVM-DaggerHilt-Dynamic-Feature-RxJava3-Flow-Sample](https://github.com/SmartToolFactory/Posts-MVVM-DaggerHilt-Dynamic-Feature-RxJava3-Flow-Sample)

It's not done yet, planning to add dashboard charts, login navigation, and search with a merge adapter to display multiple different layout with RecyclerVie. Any suggestions or feedback is more than welcome.
## [8][Full Administrative Privileges](https://www.reddit.com/r/androiddev/comments/inll8p/full_administrative_privileges/)
- url: https://www.reddit.com/r/androiddev/comments/inll8p/full_administrative_privileges/
---
Say hypothetically you gave an app from the play store full Admin privileges to your phone, what is the biggest risk, and what could be potentially compromised.
## [9][30 August - 5 September Android Newsletter](https://www.reddit.com/r/androiddev/comments/injhpy/30_august_5_september_android_newsletter/)
- url: https://www.reddit.com/r/androiddev/comments/injhpy/30_august_5_september_android_newsletter/
---
Stay up to date with Android development, in this week's edition:   
🌟 Get more feedback with in-app reviews   
💾 Jetpack DataStore   
🔻 Reduce your APK size easily   
😮 20 handy Android tools   
and much more!  

🔥Featuring [@fmuntenescu](https://twitter.com/FMuntenescu) [@JakeWharton](https://twitter.com/JakeWharton) [@prafullmishra09](https://twitter.com/prafullmishra09) [@moforemmanuel](https://twitter.com/moforemmanuel) [@mmbialas](https://twitter.com/mmbialas)

💚Enjoy [https://vladsonkin.com/android-newsletter-10/](https://vladsonkin.com/android-newsletter-10/)
## [10][Retrosheet: Turn Google Spreadsheet to JSON endpoint (for Android and JVM)](https://www.reddit.com/r/androiddev/comments/imv4u6/retrosheet_turn_google_spreadsheet_to_json/)
- url: https://github.com/theapache64/retrosheet
---

## [11][Xiaomi HCE not launching app](https://www.reddit.com/r/androiddev/comments/inhh7o/xiaomi_hce_not_launching_app/)
- url: https://www.reddit.com/r/androiddev/comments/inhh7o/xiaomi_hce_not_launching_app/
---
Hello, I'm a developer and I made an HCE app which works perfectly in other phones in all modes (started app, background or closed). But I have a problem with the Xiaomi 9T miui 12, it cannot launch the app through NFC. I tried different settings to make it work, but no one.

- Allow it in background mode battery.
- Display over other apps.
- Auto launch
- Change NFC to HCE mode.

But still not working when the app is closed, does anybody know what could be the problem?
## [12][Try out my new Game 🔥❤️ Save the CUBE from the angry Eagles 🦅 Shoot, Jump, Hide and Fight 😎](https://www.reddit.com/r/androiddev/comments/inj8sd/try_out_my_new_game_save_the_cube_from_the_angry/)
- url: https://v.redd.it/u36zgk8vyhl51
---

