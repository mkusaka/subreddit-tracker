# androiddev
## [1][Weekly "anything goes" thread!](https://www.reddit.com/r/androiddev/comments/g77oez/weekly_anything_goes_thread/)
- url: https://www.reddit.com/r/androiddev/comments/g77oez/weekly_anything_goes_thread/
---
Here's your chance to talk about whatever!

Although if you're thinking about getting feedback on an app, you should wait until tomorrow's App Feedback thread.

Remember that while you can talk about any topic, being a jerk is [still not allowed](https://www.reddit.com/r/androiddev/wiki/rules#wiki_rules_for_comments).
## [2][Weekly Questions Thread - April 20, 2020](https://www.reddit.com/r/androiddev/comments/g4qoms/weekly_questions_thread_april_20_2020/)
- url: https://www.reddit.com/r/androiddev/comments/g4qoms/weekly_questions_thread_april_20_2020/
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
## [3][Android 11 DP3 new function summary: automatically cancel application permissions / share recent apps / Ethernet hotspot / ...](https://www.reddit.com/r/androiddev/comments/g75le9/android_11_dp3_new_function_summary_automatically/)
- url: https://i.redd.it/zfpsc763mqu41.png
---

## [4][What was the hardest concept to learn when you started Android?](https://www.reddit.com/r/androiddev/comments/g7093j/what_was_the_hardest_concept_to_learn_when_you/)
- url: https://www.reddit.com/r/androiddev/comments/g7093j/what_was_the_hardest_concept_to_learn_when_you/
---
Noob here, just wanted to see what tripped up the veterans in your early days. I think the initial listview, recyclerview adapter thing tripped me up the most
## [5][Google Maps SDK Error started popping on last hour](https://www.reddit.com/r/androiddev/comments/g6t8fu/google_maps_sdk_error_started_popping_on_last_hour/)
- url: https://www.reddit.com/r/androiddev/comments/g6t8fu/google_maps_sdk_error_started_popping_on_last_hour/
---
    :com.google.android.gms.dynamite_mapsdynamite@201216046@20.12.16 (040306-0) line 9
    com.google.maps.api.android.lib6.gmm6.vector.ct.&lt;init&gt;

This error started happening on random devices. Does anyone have the same issue?

I didn't release anything, didn't change anything about the app.

&amp;#x200B;

My map fragment XML:

    &lt;fragment
    android:id="@+id/map_stops" android:name="com.google.android.gms.maps.SupportMapFragment" android:layout_width="match_parent" android:layout_height="match_parent" tools:context="com.comprovei.entregas.fragments.TabTripsMapFragment" /&gt;

&amp;#x200B;

The fragment code:

    public class TabTripsMapFragment extends Fragment implements OnMapReadyCallback {
    
        @Override
        public View onCreateView(LayoutInflater inflater, ViewGroup container, Bundle savedInstanceState) {
    
            View view = inflater.inflate(R.layout.tab_fragment_stops_map, null, false);
    
            ChildFragmentManager()
                    .findFragmentById(R.id.map_stops);
            mapFragment.getMapAsync(this);
    
            return view;
        }
    
        @Override
        public void onMapReady(final GoogleMap googleMap) {
    
            ...
    
        }
    
    }

\---

I'm trying to gather info here:

[https://stackoverflow.com/questions/61395654/google-maps-sdk-error-started-popping-on-last-hour](https://stackoverflow.com/questions/61395654/google-maps-sdk-error-started-popping-on-last-hour)

[https://issuetracker.google.com/issues?q=maps](https://issuetracker.google.com/issues?q=maps)

\--- edit: answer by google

Our engineering team continues to investigate the issue.We will provide an update by Thursday, 2020-04-23 14:30 US/Pacific with current details.Diagnosis: If you see a stack dump starting with the lines:    java.lang.ArrayIndexOutOfBoundsException: length=1; index=12          at com.google.maps.api.android.lib6.gmm6.vector.ct.&lt;init&gt;  ... you are affected by this error.Workaround: None at this time

[https://issuetracker.google.com/issues/154855417](https://issuetracker.google.com/issues/154855417)

\--- update

    ah...@google.com ah...@google.com #5823-04-2020 23:48
    Update: We have identified a possible root cause for the crash and are undoing the change.

**--- "solution" update**

    en...@google.com&lt;en...@google.com&gt; #201Apr 23, 2020 10:30PM
    10:30PMSummary: Google Maps SDK is crashing; partially resolved  
    Description: We believe the root cause of the crashes of Google Maps SDK has been fixed. The fix is being propagated to the affected applications and it is continuing toward resolution at the expected pace. Full resolution is expected to complete by Thursday, 2020-04-23 19:45 US/Pacific.  Customers for whom clearing application data is safe can recommend their users clear data for the applications (not just the cache). 
    If there is uncertainty about the safety of clearing application data, users can wait for new data to be fetched within 3 hours (many users will see resolution of the problem sooner than this).  
    
    We will provide an update by Thursday, 2020-04-23 19:30 US/Pacific with current details.  
    
    Workaround: Clear application data (not just the cache).

&amp;#x200B;
## [6][Which is better on Android: divide by 2 or shift by 1? – Jake Wharton](https://www.reddit.com/r/androiddev/comments/g6npe8/which_is_better_on_android_divide_by_2_or_shift/)
- url: https://jakewharton.com/which-is-better-on-android-divide-by-two-or-shift-by-one/
---

## [7][How do alarm clock apps work?](https://www.reddit.com/r/androiddev/comments/g728hu/how_do_alarm_clock_apps_work/)
- url: https://www.reddit.com/r/androiddev/comments/g728hu/how_do_alarm_clock_apps_work/
---
I feel like I’ve read a thousand posts just like [this one](https://stackoverflow.com/questions/29299470/alarmmanager-rtc-wakeup-not-wakes-up-my-android-devices-when-it-is-switched-off) or [this one](https://proandroiddev.com/android-alarmmanager-as-deep-as-possible-909bd5b64792). I certainly have tried both `AlarmManager.setExactAndAllowWhileIdle` and `AlarmManager.setAlarmClock` after reading [the only documentation that exists](https://stackoverflow.com/questions/34699662/how-does-alarmmanager-alarmclockinfos-pendingintent-work) for `setAlarmClock`. You may also rest assured that on Android 10 I tried scheduling consecutive alarms that fire a notification with a full screen intent for a full-screen activity and the `USE_FULL_SCREEN_INTENT` manifest permission, as described by [the missing manual](https://android.jlelse.eu/full-screen-intent-notifications-android-85ea2f5b5dc1) to full-screen intent notifications.

None of that allows subsequent alarms to turn the screen on from a screen-off state after the first alarm within a 9-minute time period. 

All of my test devices ranging from Android 6 to Android 10 have a built-in Clock app that is capable of setting two alarms, one minute apart. Both alarms wake up the screen within milliseconds of the minute I wished it to fire. Both alarms appear to turn the screen on and draw a full-screen activity over the lock screen and gives me a snooze/dismiss affordance, otherwise keeping the phone locked. Amazingly, the alarm I scheduled for a minute after the first one provides the exact same experience as the first one: it too wakes up the phone while the screen is off.

I don't understand how the clock apps' alarm features could possibly be driven by any of the `AlarmManager.set...` methods, if there are power restrictions inhibiting wake behavior on subsequent `AlarmManager.set...` fires [within 9 minutes](https://developer.android.com/training/monitoring-device-state/doze-standby). So what are they doing instead?
## [8][How to test my app on my computer](https://www.reddit.com/r/androiddev/comments/g75iy3/how_to_test_my_app_on_my_computer/)
- url: https://www.reddit.com/r/androiddev/comments/g75iy3/how_to_test_my_app_on_my_computer/
---
Hi all, I'm building an Android app using Python/Kivy (and then pushing it to my phone using Buildozer).

All my code seems to work fine, however I noticed that when I put the app on my phone, there are a few things that the Android environment does, which don't show up on my laptop. One example is the hidden keyboard that will pop up whenever you try and type something, causing things to move.

Instead of constantly making one tweak, then loading it to my phone to check, and then doing it again, I was wandering if there was something I can download that mimics my phone on my laptop, so I can run my testing there?

If it helps, I am pushing my app to my phone using Ubuntu from a virtual box, so I will most likely download this there.

&amp;#x200B;

Thanks all!
## [9][The ProtonMail Android app is open source](https://www.reddit.com/r/androiddev/comments/g6rd0d/the_protonmail_android_app_is_open_source/)
- url: https://protonmail.com/blog/android-open-source/
---

## [10][What is your FCM delivery ratio?](https://www.reddit.com/r/androiddev/comments/g716i9/what_is_your_fcm_delivery_ratio/)
- url: https://www.reddit.com/r/androiddev/comments/g716i9/what_is_your_fcm_delivery_ratio/
---
Hey - anybody track their FCM delivery ratio? We've been trying for a couple of years to get it close to 100% but the highest we've had it hit is 89%. 

&amp;#x200B;

We randomly get users complaining about not having received a notification, even though we sent it to the FCM server and get an OK response back. It just doesn't seem to make it to the device for those 11%.
## [11][Seems like Google Maps API is down and causes crashes.](https://www.reddit.com/r/androiddev/comments/g6uq37/seems_like_google_maps_api_is_down_and_causes/)
- url: https://outage.report/google-maps
---

## [12][PrivacyBreacher - an app showcasing the privacy issues in Android operating system!](https://www.reddit.com/r/androiddev/comments/g73bct/privacybreacher_an_app_showcasing_the_privacy/)
- url: https://www.reddit.com/r/androiddev/comments/g73bct/privacybreacher_an_app_showcasing_the_privacy/
---
This Android app can access the following information from your phone *without requesting any permissions*:

1. **Figure out at what time your phone screen turned on/off.**
2. **Figure out at what time you plugged in or removed your phone charger and wired headphones.**
3. **Figure out at what time you switched on/off your phone (i.e., it captures the device uptime and ACTION\_SHUTDOWN broadcasts).**
4. **Access most of your device related information like your phone model, manufacturer etc.**
5. **Keep track of your WiFi/Mobile data usage.**
6. **Get a list of all the apps installed on your phone.**
7. **Construct a 3D visualization of your body movements.**

**Check out more on here:** [**https://github.com/databurn-in/PrivacyBreacher**](https://github.com/databurn-in/PrivacyBreacher)
