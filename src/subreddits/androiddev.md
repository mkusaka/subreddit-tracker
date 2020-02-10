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
## [3][Me: So what's wrong? Google: Yes.](https://www.reddit.com/r/androiddev/comments/f1hopo/me_so_whats_wrong_google_yes/)
- url: https://i.redd.it/qoc8cf9mmzf41.png
---

## [4][Better ways of using Room Database](https://www.reddit.com/r/androiddev/comments/f1mz6w/better_ways_of_using_room_database/)
- url: https://www.reddit.com/r/androiddev/comments/f1mz6w/better_ways_of_using_room_database/
---
[https://jakubpchmiel.com/better-ways-of-using-room/](https://jakubpchmiel.com/better-ways-of-using-room/)
## [5][New ViewBinding sample in the Architecture Components Samples github repo](https://www.reddit.com/r/androiddev/comments/f1i5ps/new_viewbinding_sample_in_the_architecture/)
- url: https://github.com/android/architecture-components-samples/tree/9f021451fd64362c7c227802bacf8cfe476af0be/ViewBindingSample
---

## [6][How can I fix this?](https://www.reddit.com/r/androiddev/comments/f1o5bq/how_can_i_fix_this/)
- url: https://i.redd.it/dz4hlmx9j2g41.jpg
---

## [7][Need help with app removed from Play Store](https://www.reddit.com/r/androiddev/comments/f1pgip/need_help_with_app_removed_from_play_store/)
- url: https://www.reddit.com/r/androiddev/comments/f1pgip/need_help_with_app_removed_from_play_store/
---
My app was removed from Play Store , for not filling out Target Audience and Content form ,

Now my app is an informative app , so it has not particular age group that it specifically targets , So generally it is an app for Adults and Children , So while filling Target Audience form , I am selecting all age groups , but I am serving ads in my app via Facebook Ads network , and I have found that Google requires to use only Google Certified Ad network for serving ads to children , now I cant see Facebook Ads network as Google Certified

I have removed Facebook Ads (Right Now , I don't show any ads at all) and made fixes to comply with Google Play Policies , but I cant release my app to rollout or even resubmit my app since first I have to fill Target Audience and Content Form.

So , my question is what should I do while filling Target Audience and Content form , because If I select all age groups (Adults and Children ) , and If I submit my app then I will be violating ads via Facebook Ads Network 

Please help !
## [8][Is there any image morphing transition library available that i can use in android studio?](https://www.reddit.com/r/androiddev/comments/f1pfdt/is_there_any_image_morphing_transition_library/)
- url: https://www.reddit.com/r/androiddev/comments/f1pfdt/is_there_any_image_morphing_transition_library/
---
So when i switch from one image to another, the picture should morph itself to the next image by keeping the constant features of the image unchanged and rest of the image just warping to the new one. The two images are pretty much similar to one another (i.e they have mostly have same background or 80% remains same). If it's not available then what's the best way to achieve this?
## [9][Android sensors don't gather data when phone is idle](https://www.reddit.com/r/androiddev/comments/f1p0pe/android_sensors_dont_gather_data_when_phone_is/)
- url: https://www.reddit.com/r/androiddev/comments/f1p0pe/android_sensors_dont_gather_data_when_phone_is/
---
I’m trying to develop an app for Android in which I pick up data from several sensors (if available on the device) and write it down to a file which will later be analyzed for certain uses.

I’m facing several problems, a minor one which I can kind of ignore and a major one that I haven’t been able to solve and makes the app not work properly.

&amp;#x200B;

**- Minor problem**

&amp;#x200B;

I’m gathering data from: Accelerometer, Linear Accelerometer, Gyroscope and Magnetometer and also from the GPS but that works quite differently and can only be sampled at much lower frequencies, so I’ll ignore it for now.

I gather the data by implementing a listener for each sensor:

&amp;#x200B;

    public class AccelerometerWatcher implements SensorEventListener
    {
        private SensorManager sm;
        private Sensor accelerometer;
        AccelerometerWatcher(Context context) {
            sm = (SensorManager)context.getSystemService(Context.SENSOR_SERVICE);
            assert sm != null;
            if (sm.getDefaultSensor(Sensor.TYPE_ACCELEROMETER) != null) {
                accelerometer = sm.getDefaultSensor(Sensor.TYPE_ACCELEROMETER);
            }
        }
    }

&amp;#x200B;

And I’m setting the frequency to \~50Hz by using:

&amp;#x200B;

    sm.registerListener(this, accelerometer, SensorManager.SENSOR_DELAY_GAME);

&amp;#x200B;

When gathering data, I understand the frequency can’t be 100% stable, but the weird thing is it stays more or less stable on every sensor (at around 50Hz) except on the Accelerometer, where most of the time it samples at 100Hz and sometimes drops down to 50Hz.

&amp;#x200B;

Is there something I might be doing wrong or any way to control this? So far it’s happened in every device I tried, although they don’t all behave in exactly the same way.

&amp;#x200B;

**- Major problem**

&amp;#x200B;

I’m writing down the info to a file by first writing everything I pick up from the sensors to a string and then every X seconds, writing what’s on the string to a file and clearing it so the sensor listeners can keep on writing on it but it doesn’t become infinitely long.

&amp;#x200B;

I write on the string like this:

&amp;#x200B;

    @Override
    public void onSensorChanged(SensorEvent event) {
        if (event.sensor.getType() != Sensor.TYPE_ACCELEROMETER)
        return;
        if(initTime == -1)
            initTime = event.timestamp;
        MyConfig.SENSOR_ACCEL_READINGS += ((event.timestamp - initTime) / 1000000L) +     MyConfig.DELIMITER + event.values[0] + MyConfig.DELIMITER + event.values[1] + MyConfig.DELIMITER + event.values[2] + "\n";
    }

&amp;#x200B;

And then save it to a file using this:

&amp;#x200B;

    public class Utils {
    private static Timer timer;
    private static TimerTask timerTask;
    public static void startRecording() {
        timer = new Timer();
        timerTask = new TimerTask()
        {
            @Override
            public void run()
            {
                // THIS CODE RUNS EVERY x SECONDS
                writeDataToFile();
            }
        };
        timer.scheduleAtFixedRate(timerTask, 0, MyConfig.SAVE_TIMER_PERIOD);
    }
    public static void stopRecording()
    {
        if(timer != null)
            timer.cancel();
        if(timerTask != null)
            timerTask.cancel();
        writeDataToFile();
    }
    private static void writeDataToFile()
    {
        String temp_accel =     String.copyValueOf(MyConfig.SENSOR_ACCEL_READINGS.toCharArray());
        WriteData.write(MyConfig.RECORDING_FOLDER, MyConfig.FILENAME_ACCEL, temp_accel);
        MyConfig.SENSOR_ACCEL_READINGS =     MyConfig.SENSOR_ACCEL_READINGS.replaceFirst(temp_accel, "");
    }

&amp;#x200B;

In the listener, every time I stop listening, I set “initTime” to -1 so the samples always start at 0 and go up to the duration of the listening period in miliseconds. (Ignore the DELIMITER it’s just a matter of formatting).

&amp;#x200B;

**My main app-breaking problem, is the following:**

&amp;#x200B;

In most phones (a few lucky ones work flawlessly) 1 or 2 things fail.

&amp;#x200B;

In some, after being idle for a while (locked and in your pocket for example) the sensors stop recording data so the app just writes blank values until I wake the phone up again.

&amp;#x200B;

In others, it’s even worse, not only do the sensors stop recording data, but the timer / writing to file, seems to stop working too, and when the phone wakes up again, it tries to write what it should’ve written while it wasn’t working and messes up all the timestamps, writing the same samples at different points “in the past” until it catches up to the current time. (If you visualize it as a graph, it basically looks as if the data gathering travelled back in time).

&amp;#x200B;

**Is there any way in which I can make sure that the app keeps on working no matter what, whether the phone is locked, dozing, the app is minimized, on the background, foreground, etc.?**

&amp;#x200B;

I tried a method I googled that consists of setting and alarm to "wake up the process" every X seconds (no matter what time I set to it, it only worked max once per minute).

I saw how for a few miliseconds every time the alarm went off, it captured samples again but then went to sleep right away, it didn't keep the phone "awake" for a longer period of time.

It solved nothing and even for the brief period it forced the sensors to gather data, it only helped wake up the sensors, the problem with the timer / writing to file still persisted.

&amp;#x200B;

Hope someone can shed some light on how to keep the phone gathering data no matter what, I've been trying everything I could think of and I'm not getting anywhere. Sorry for the brick of text, but I didn't really know how to explain it in a shorter way.

&amp;#x200B;

P.S: I saw that having the Battery Saver ON made it even worse, even on the phones where it usually worked properly, it started messing things up. So another question would be... How can I stop it from interfering?
## [10][How do I hook the physical buttons of a phone with the camera?](https://www.reddit.com/r/androiddev/comments/f1n1ne/how_do_i_hook_the_physical_buttons_of_a_phone/)
- url: https://www.reddit.com/r/androiddev/comments/f1n1ne/how_do_i_hook_the_physical_buttons_of_a_phone/
---
I'm learning android app programming and I want to learn how to hook the physical buttons, for e.g., the volume up/down buttons, with the camera, and, whenever I press a combination of the physical button, the camera should start.

Can anyone please tell me how do I go about doing that?
## [11][What's the difference between button.findViewById() &amp; button= findViewById()](https://www.reddit.com/r/androiddev/comments/f1ol3i/whats_the_difference_between_buttonfindviewbyid/)
- url: https://www.reddit.com/r/androiddev/comments/f1ol3i/whats_the_difference_between_buttonfindviewbyid/
---
I was working on a small app for my own and i faced crashes i didn't know from what later on i realized it was because i was using   button.findViewById() instead off  button = findViewById() so what are the deference between those two, i searched for it but couldn't find answers.
## [12][How many downloads / users did you get when you launched your app (stories please).](https://www.reddit.com/r/androiddev/comments/f1d0bu/how_many_downloads_users_did_you_get_when_you/)
- url: https://www.reddit.com/r/androiddev/comments/f1d0bu/how_many_downloads_users_did_you_get_when_you/
---
I'm trying to decide how we should prioritize shipping our app in the Android app store.  

I think part of the challenge is that some apps explode, but other apps sort of don't really get much distribution.

I'd like to figure out WHY some apps explode and others don't.

I have some theories:

- some apps leverage trends (either deliberately or accidentally) and then they show up #1 for some key search terms.

- some apps already have a large user base, just not an app, so when they email their users they get a HUGE explosion of initial downloads.

- some people just PPC their app to the top and then they end up showing up in the trending section.  I think this would work but probably costs a fair amount of $$... just not sure how much.

Would love to hear some real world stories of how you got your app to take off.
