# androiddev
## [1][Weekly "anything goes" thread!](https://www.reddit.com/r/androiddev/comments/jtg9hc/weekly_anything_goes_thread/)
- url: https://www.reddit.com/r/androiddev/comments/jtg9hc/weekly_anything_goes_thread/
---
Here's your chance to talk about whatever!

Although if you're thinking about getting feedback on an app, you should wait until tomorrow's App Feedback thread.

Remember that while you can talk about any topic, being a jerk is [still not allowed](https://www.reddit.com/r/androiddev/wiki/rules#wiki_rules_for_comments).
## [2][Weekly Questions Thread - November 09, 2020](https://www.reddit.com/r/androiddev/comments/jqvlp4/weekly_questions_thread_november_09_2020/)
- url: https://www.reddit.com/r/androiddev/comments/jqvlp4/weekly_questions_thread_november_09_2020/
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
## [3][Has anyone updated to MacOS Big Sur and is Android Studio working fine with it?](https://www.reddit.com/r/androiddev/comments/jtbl4m/has_anyone_updated_to_macos_big_sur_and_is/)
- url: https://www.reddit.com/r/androiddev/comments/jtbl4m/has_anyone_updated_to_macos_big_sur_and_is/
---
Since a lot has changed in Big Sur and Macbook being my development machine, I just want to make sure before upgrading that Android Studio works fine or not otherwise will have to take leave from work.
## [4][More WindowInsetsController APIs backported!](https://www.reddit.com/r/androiddev/comments/jt77s6/more_windowinsetscontroller_apis_backported/)
- url: https://i.redd.it/8uw29rn1nwy51.png
---

## [5][AOSP is migrating from its current build systems to Bazel](https://www.reddit.com/r/androiddev/comments/jsztsc/aosp_is_migrating_from_its_current_build_systems/)
- url: https://blog.bazel.build/2020/11/12/aosp_migrating_to_bazel.html
---

## [6][AdMob users. What do you think about keeping test devices in your production builds? Would you risk getting your devices spoofed and getting banned from AdMob? Or would you rather risk getting banned from invalid traffic due to forgetfulness?](https://www.reddit.com/r/androiddev/comments/jtfvsl/admob_users_what_do_you_think_about_keeping_test/)
- url: https://www.reddit.com/r/androiddev/comments/jtfvsl/admob_users_what_do_you_think_about_keeping_test/
---
I can't figure out which is riskier. The choice is binary and either choices could result in getting banned. My test devices are usually removed from my official release APKs. I recently decided to resume work on one of my AdMob apps, so I played around with the release from the Play Store and forgot that my devices were no longer interpreted as test devices. I panicked and rushed to run the the dev build on all my devices.

I thought about just keeping the test device IDs in my future releases, but I googled to make sure if I should do it. And people said someone can then decompile the APK and extract the device IDs, then spoof my devices. I believe malicious actors can then generate invalid traffic with my spoofed device IDs.

On the other hand people have said they were suspended or demonetized because they didn't use test IDs on their official releases.

It seems like both choices are bad. And hearing about how Google treats Android devs here, I feel hopeless! What do you guys usually do with your test device IDs? Keep or remove?
## [7][Modern approach to list and details screens](https://www.reddit.com/r/androiddev/comments/jtcht4/modern_approach_to_list_and_details_screens/)
- url: https://www.reddit.com/r/androiddev/comments/jtcht4/modern_approach_to_list_and_details_screens/
---
I created a sample project here
https://github.com/lawloretienne/CryptoCoins

It's a really basic app but i just wanted to showcase an approach to building a list screen and details screen.

Let me know what you think.
## [8][Kotlin Linters: which should I choose?](https://www.reddit.com/r/androiddev/comments/jtervp/kotlin_linters_which_should_i_choose/)
- url: https://akuleshov7.com/2020-11-10-kotlin-static-analyzers.html
---

## [9][Monetisation strategy: nsfw game on patreon, sfw on Google Play](https://www.reddit.com/r/androiddev/comments/jtcswd/monetisation_strategy_nsfw_game_on_patreon_sfw_on/)
- url: https://www.reddit.com/r/androiddev/comments/jtcswd/monetisation_strategy_nsfw_game_on_patreon_sfw_on/
---
Will Google play allow such thing?

On Google there will be a sfw version of the game with link to patreon where user can subscribe to get the nsfw version.
## [10][Hi can anyone help me how to implement this project I have no idea on tasks...etc I know how to design front end](https://www.reddit.com/r/androiddev/comments/jtfw5f/hi_can_anyone_help_me_how_to_implement_this/)
- url: https://i.redd.it/by1mzhfh00z51.png
---

## [11][whats the best approach to making an iso of my android tv ?](https://www.reddit.com/r/androiddev/comments/jtfq4m/whats_the_best_approach_to_making_an_iso_of_my/)
- url: https://www.reddit.com/r/androidhacking/comments/jtfohe/dumping_my_androids_tv_install/?utm_source=share&amp;utm_medium=ios_app&amp;utm_name=iossmf
---

## [12][Mockito unit test strange behaviour for custom class](https://www.reddit.com/r/androiddev/comments/jtffkf/mockito_unit_test_strange_behaviour_for_custom/)
- url: https://www.reddit.com/r/androiddev/comments/jtffkf/mockito_unit_test_strange_behaviour_for_custom/
---
 I have a custom class:

    data class ContainerModel(   
      ...   
      val data: List&lt;MainModel&gt;? 
    ) {    
      fun get(id: String) = data?.firstOrNull { it.id == id } }   
    
    
    data class MainModel(   
      val id: String?,   
      val documentGroups: List&lt;DocumentGroup&gt;?,   
      .... 
    ) {   
      fun hasAnyDocument() = documentGroups != null 
    } 

In my unit tests, I setup up like this:

    import org.mockito.Mockito.`when` as mwhen     
    
    private lateinit var containerModel: ContainerModel    
    
    @Mock lateinit var m1: MainModel    
    
    @Mock lateinit var documentGroup: DocumentGroup    
    
    @Before fun setup() {     
      m1.apply {       
        mwhen(id).thenReturn("m1")  
        mwhen(documentGroups).thenReturn(listOf(documentGroup))     
      }      
    
      containerModel = Mockito.spy(ContainerModel(data = arrayListOf(m1)))   
    } 

This test is successful

    @Test 
    fun get() {     
      Truth.assertThat(containerModel.get("m0")).isNull() // successful  
      Truth.assertThat(containerModel.get("m1")).isEqualTo(m1) // successful   
    } 

But this is strange

    @Test fun checkDocuments() {    
      Truth.assertThat(containerModel.get("m1")!!.documentGroups).isNotNull() // successful       
      Truth.assertThat(containerModel.get("m1")!!.documentGroups?.size).isEqualTo(1) // successful     
      Truth.assertThat(containerModel.get("m1")!!.hasAnyDocument()).isTrue() // FAILS  
    } 

It's super strange to me as if I check on documentGroups  
 it's clearly not null and has one item, but calling a function that uses it fails? what could be the issue?

Thanks!
