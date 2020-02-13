# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/f1psjk/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/f1psjk/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - February 10, 2020](https://www.reddit.com/r/androiddev/comments/f1opq1/weekly_questions_thread_february_10_2020/)
- url: https://www.reddit.com/r/androiddev/comments/f1opq1/weekly_questions_thread_february_10_2020/
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
## [3][How do you test updating of third party libraries?](https://www.reddit.com/r/androiddev/comments/f36soc/how_do_you_test_updating_of_third_party_libraries/)
- url: https://www.reddit.com/r/androiddev/comments/f36soc/how_do_you_test_updating_of_third_party_libraries/
---
Imagine that you have to update some third party libraries in your project. How will you check that nothing has been broken? Manual testing? Automated UI-tests? Maybe some other approach (like isolated tests for a certain library)?
## [4][Using dagger in multi-module apps](https://www.reddit.com/r/androiddev/comments/f2z9xb/using_dagger_in_multimodule_apps/)
- url: https://developer.android.com/training/dependency-injection/dagger-multi-module
---

## [5][For people using Appodeal, what does your "prominent disclosure" look like?](https://www.reddit.com/r/androiddev/comments/f390m3/for_people_using_appodeal_what_does_your/)
- url: https://www.reddit.com/r/androiddev/comments/f390m3/for_people_using_appodeal_what_does_your/
---
My app was removed with this as the reason:

&gt;Your app is using the Appodeal SDK, which is uploading user’s phone number, social accounts and installed packages information without a prominent disclosure. Prior to the collection and transmission, it must prominently highlight how the user data will be used, describe the type of data being collected and have the user provide affirmative consent for such use.

I handwrote a crappy privacy policy that def deserved to get taken down, so I'm curious what a good Appodeal disclosure looks like before trying to make a new one and potentially getting another strike. I understand there are generators out there but it would be great to see what one looks like that is actually in a live app that's acceptable by Google.

Edit: also just read Appodeal recommends linking [their privacy policy](https://www.appodeal.com/home/privacy-policy/) in yours, so that must be a part of it.
## [6][Are interstitial ads on back button allowed?](https://www.reddit.com/r/androiddev/comments/f38fae/are_interstitial_ads_on_back_button_allowed/)
- url: https://www.reddit.com/r/androiddev/comments/f38fae/are_interstitial_ads_on_back_button_allowed/
---
Is it allowed to use Admob's interstitial ads when Back Button is pressed on a specific activity? Can't find a recent legit answer.
## [7][Long press for minor shakes](https://www.reddit.com/r/androiddev/comments/f31kp4/long_press_for_minor_shakes/)
- url: https://www.reddit.com/r/androiddev/comments/f31kp4/long_press_for_minor_shakes/
---
Getting old got a little bit of the shakes and when I try to do a long press sometimes I hold it to the icon but it interprets it as a short press or a click.

  


Is there anything out there that will help me?

  


  


TIA
## [8][Fragments overlapping](https://www.reddit.com/r/androiddev/comments/f39qxc/fragments_overlapping/)
- url: https://www.reddit.com/r/androiddev/comments/f39qxc/fragments_overlapping/
---
Hello, I've been searching everywhere and the only solution I found was to add a background and clickable to true in the fragment...

Any other tip ir solution besides that?
## [9][How to test Android Intent's Bundle using Mockk?](https://www.reddit.com/r/androiddev/comments/f3967y/how_to_test_android_intents_bundle_using_mockk/)
- url: https://www.reddit.com/r/androiddev/comments/f3967y/how_to_test_android_intents_bundle_using_mockk/
---
I would like to test a custom handler class which is starting a new Activity. I would like to test the Intent's Bundle if contains the pre defined parameters.

**The test class**:

    @MockK
    lateinit var activity: ActivityCalendar
    
    @Before
    fun setUp() {
        MockKAnnotations.init(this)
    }
    
    @Test
    fun testActivityBundles() {
        val book = mockk&lt;Book&gt;()
    
        every { book.releaseDate } returns GregorianCalendar().apply { this.timeInMillis = 1423825586000 }
        every { activity.startActivity(any()) } just Runs
    
        val handler = ActivityHandler(activity)
        handler.startRequiredActivity(book)
    
        verify { activity.startActivity(
                withArg { intent -&gt;
                    val bundle = intent.extras!!
                    val releaseDateTimeMillis = bundle.getLong("release_date", 0L)
    
                    Assert.assertEquals(1423825586000, releaseDateTimeMillis)
                }
        ) }
    }

The code above is crashing at line: **val bundle = intent.extras!!** but it shouldn't.

**The class that I want to test**:

    class ActivityHandler(val activity: Activity) {
        fun startRequiredActivity(book: Book) {
            val intent = buildIntent(book)
    
            activity.startActivity(intent)
        }
    
        private fun buildIntent(book: Book): Intent {
            val extras = Bundle().apply {
                this.putLong("release_Date", book.releaseDate.timeInMillis)
            }
    
            return Intent(activity, ActivityBookDetails::class.java).apply {
                putExtras(extras)
            }
        }
    }
    
    data class Book(
            val releaseDate: GregorianCalendar
    )

I debugged the code and I found out that function **private fun buildIntent(book: Book): Intent** is returning an "null" object (string "null" and not Java NULL).
## [10][Integrating PayPal with NativeScript (Android only)](https://www.reddit.com/r/androiddev/comments/f39250/integrating_paypal_with_nativescript_android_only/)
- url: https://www.reddit.com/r/androiddev/comments/f39250/integrating_paypal_with_nativescript_android_only/
---
I'm just trying to get PayPal to work on Android right now. Many of the solutions online recommend using the PayPal sdk in the Gradle compile, but as far as I know the sdk has moved from the native to a CDN. Therefore, I'm stuck in a standoff.

1. Implement PayPal with WebView and just had a bridge.
2. Implement PayPal on the serverside since nodejs seems to have more support.
3. Implement PayPal on Android, which is the most confusing at this point.

Any help with this would be awesome and so would any resources to help me out.
## [11][How to use nfc in flutter](https://www.reddit.com/r/androiddev/comments/f38qqo/how_to_use_nfc_in_flutter/)
- url: https://www.reddit.com/r/androiddev/comments/f38qqo/how_to_use_nfc_in_flutter/
---
Hello I'm a beginner in flutter and I have to use NFC in my project but I don't have any clue how and where to start ...can anyone please help me how can I use NFC in my application. It would help me to get Internship
## [12][Android Styling: Common Theme Attributes](https://www.reddit.com/r/androiddev/comments/f2rlr8/android_styling_common_theme_attributes/)
- url: https://medium.com/androiddevelopers/android-styling-common-theme-attributes-8f7c50c9eaba
---

