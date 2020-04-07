# androiddev
## [1][Weekly "who's hiring" thread!](https://www.reddit.com/r/androiddev/comments/fvy3fg/weekly_whos_hiring_thread/)
- url: https://www.reddit.com/r/androiddev/comments/fvy3fg/weekly_whos_hiring_thread/
---
Looking for Android developers? Heard about a cool job posting? Let people know!

Here is a suggested posting template:

&gt; Company: &lt;Best Company Ever&gt;  
&gt; Job: [&lt;Title&gt;]\(https://example.com/job)  
&gt; Location: &lt;City, State, Country&gt;  
&gt; Allows remote: &lt;Yes/No&gt;  
&gt; Visa: &lt;Yes/No&gt;  

Feel free to include any other information about the job.
## [2][Weekly Questions Thread - April 06, 2020](https://www.reddit.com/r/androiddev/comments/fvwq7t/weekly_questions_thread_april_06_2020/)
- url: https://www.reddit.com/r/androiddev/comments/fvwq7t/weekly_questions_thread_april_06_2020/
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
## [3][I just finished my first full-size MVVM app - Upcoming Games. Can I get a code review?](https://www.reddit.com/r/androiddev/comments/fwissh/i_just_finished_my_first_fullsize_mvvm_app/)
- url: https://www.reddit.com/r/androiddev/comments/fwissh/i_just_finished_my_first_fullsize_mvvm_app/
---
[https://github.com/gavingt/upcoming-games](https://github.com/gavingt/upcoming-games)

[https://youtu.be/RS\_684kyOEk](https://youtu.be/RS_684kyOEk)

The app lets you view release dates and other info for every video game ever made.  It allows you to easily track your favorite and most anticipated games. Contains full sorting and search capabilities. The database is updated periodically using WorkManager.

This is my first time using several modern Android libraries, so I'm sure I've made some mistakes along the way. Any feedback would be appreciated.

Also, what more do I need to learn to become hireable?

Thanks!
## [4][Migrating Duolingo’s Android app to 100% Kotlin](https://www.reddit.com/r/androiddev/comments/fw45de/migrating_duolingos_android_app_to_100_kotlin/)
- url: https://blog.duolingo.com/migrating-duolingos-android-app-to-100-kotlin/
---

## [5][Playstore App Publication process when you push put multiple updates](https://www.reddit.com/r/androiddev/comments/fwhnir/playstore_app_publication_process_when_you_push/)
- url: https://www.reddit.com/r/androiddev/comments/fwhnir/playstore_app_publication_process_when_you_push/
---
I published a new update for my app on Playstore on 3rd April. However I found an bug in the latest code and published a fix for that on 5th April. Since my 3rd April update  wasn't rolled out, what happens now ? Will my latest update be rolled out first or it will happen sequentially , i.e. 3rd April one first and then 5th April one ?

I've been unable to find anything regarding this in the documentation, so any info would be helpful. 

Thanks
## [6][A step-by-step guide to creating an IoT based Flutter app that interacts with any home electrical equipment.](https://www.reddit.com/r/androiddev/comments/fwjjif/a_stepbystep_guide_to_creating_an_iot_based/)
- url: https://blog.codemagic.io/creating-iot-based-flutter-app/
---

## [7][Should i memorize codes?](https://www.reddit.com/r/androiddev/comments/fwikuu/should_i_memorize_codes/)
- url: https://www.reddit.com/r/androiddev/comments/fwikuu/should_i_memorize_codes/
---
Okay so im a learning Android Development. I learned kotlin as well as java. These language are really easy to remember for me at least. But when i get to the actual development there are more codes that i have to write in Android studio. So should i memorize these code. For example there are lots of codes in recycler view. Should i memorize all of em. I keep notes of everything. So when i start developing apps i can look at them when i forget the codes. I hope my post makes sense. Would appreciate help from fellow developers. Thabks in advance
## [8][Unable to "interact(?)" with buttons etc... outside of the MainAcitivty](https://www.reddit.com/r/androiddev/comments/fwijfd/unable_to_interact_with_buttons_etc_outside_of/)
- url: https://www.reddit.com/r/androiddev/comments/fwijfd/unable_to_interact_with_buttons_etc_outside_of/
---
Title. ( And please see the comment in LoginActivity.kt) Here are the codes:

MainActivity.kt

    package com.andrycoder.kotlinmessages
    
    import android.content.Intent
    import androidx.appcompat.app.AppCompatActivity
    import android.os.Bundle
    import android.util.Log
    import kotlinx.android.synthetic.main.activity_main.*
    
    class MainActivity : AppCompatActivity() {
        override fun onCreate(savedInstanceState: Bundle?) {
            super.onCreate(savedInstanceState)
            setContentView(R.layout.activity_main)
    
            register_button_register.setOnClickListener {
                val email = email_edittext_register.text.toString()
                val password = password_edittext_register.text.toString()
                Log.d("MainActivity Activity", "The email is: \"$email\".")
                Log.d("MainActivity Activity", "The password is: \"$password\".")
            }
    
            login_button_register.setOnClickListener {
                Log.d("MainActivity Activity", "Login button clicked, try to show login activity.")
                val intent = Intent(this, LoginActivity::class.java)
                startActivity(intent)
                Log.d("LoginActivity", "LoginActivity has been opened")
            }
    
        }
    }

activity\_main.xml

    &lt;?xml version="1.0" encoding="utf-8"?&gt;
    &lt;androidx.constraintlayout.widget.ConstraintLayout xmlns:android="http://schemas.android.com/apk/res/android"
        xmlns:app="http://schemas.android.com/apk/res-auto"
        xmlns:tools="http://schemas.android.com/tools"
        android:id="@+id/editText"
        android:layout_width="match_parent"
        android:layout_height="match_parent"
        android:background="#FFFFFF"
        android:scrollbarSize="24dp"
        tools:context=".MainActivity"&gt;
    
        &lt;TextView
            android:id="@+id/textView2"
            android:layout_width="363dp"
            android:layout_height="36dp"
            android:layout_marginTop="104dp"
            android:gravity="center"
            android:text="@string/kotlinmessages"
            android:textColor="@color/colorAccent"
            android:textSize="24sp"
            android:textStyle="bold"
            app:layout_constraintEnd_toEndOf="parent"
            app:layout_constraintStart_toStartOf="parent"
            app:layout_constraintTop_toTopOf="parent" /&gt;
    
        &lt;EditText
            android:id="@+id/username_edittext_register"
            android:layout_width="363dp"
            android:layout_height="wrap_content"
            android:layout_marginTop="94dp"
            android:ems="10"
            android:hint="@string/username"
            android:inputType="textPersonName"
            app:layout_constraintEnd_toEndOf="parent"
            app:layout_constraintStart_toStartOf="parent"
            app:layout_constraintTop_toBottomOf="@+id/textView2"
            android:autofillHints="" /&gt;
    
        &lt;EditText
            android:id="@+id/email_edittext_register"
            android:layout_width="363dp"
            android:layout_height="wrap_content"
            android:layout_marginTop="16dp"
            android:ems="10"
            android:hint="@string/email"
            android:inputType="textPersonName"
            app:layout_constraintEnd_toEndOf="@+id/username_edittext_register"
            app:layout_constraintStart_toStartOf="@+id/username_edittext_register"
            app:layout_constraintTop_toBottomOf="@+id/username_edittext_register"
            android:autofillHints="" /&gt;
    
        &lt;EditText
            android:id="@+id/password_edittext_register"
            android:layout_width="363dp"
            android:layout_height="wrap_content"
            android:layout_marginTop="16dp"
            android:ems="10"
            android:hint="@string/password"
            android:inputType="textPassword"
            app:layout_constraintEnd_toEndOf="@+id/email_edittext_register"
            app:layout_constraintStart_toStartOf="@+id/email_edittext_register"
            app:layout_constraintTop_toBottomOf="@+id/email_edittext_register"
            android:autofillHints="" /&gt;
    
        &lt;Button
            android:id="@+id/register_button_register"
            android:layout_width="363dp"
            android:layout_height="wrap_content"
            android:layout_marginTop="64dp"
            android:background="@color/colorPrimary"
            android:text="@string/register"
            android:textColor="#FFFFFF"
            app:layout_constraintEnd_toEndOf="@+id/password_edittext_register"
            app:layout_constraintHorizontal_bias="1.0"
            app:layout_constraintStart_toStartOf="@+id/password_edittext_register"
            app:layout_constraintTop_toBottomOf="@+id/password_edittext_register" /&gt;
    
        &lt;Button
            android:id="@+id/login_button_register"
            android:layout_width="92dp"
            android:layout_height="32dp"
            android:layout_marginTop="8dp"
            android:background="#00FFFFFF"
            android:gravity="center"
            android:text="@string/login"
            android:textColor="#A1A1A1"
            android:textSize="14sp"
            app:layout_constraintEnd_toEndOf="@+id/register_button_register"
            app:layout_constraintStart_toStartOf="@+id/register_button_register"
            app:layout_constraintTop_toBottomOf="@+id/register_button_register" /&gt;
    &lt;/androidx.constraintlayout.widget.ConstraintLayout&gt;

LoginActivity.kt

    package com.andrycoder.kotlinmessages
    
    import android.content.BroadcastReceiver
    import android.content.Intent
    import android.os.Bundle
    import android.util.Log
    import androidx.appcompat.*
    import kotlinx.android.synthetic.main.activity_login.*
    import androidx.appcompat.app.AppCompatActivity
    
    
    class LoginActivity : AppCompatActivity() {
        override fun onCreate(savedInstanceState: Bundle?) {
            super.onCreate(savedInstanceState)
            setContentView(R.layout.activity_login)
            Log.d("MainActivity Activity", "LoginActivity opened.")
        }
    
        // from here I cannot do stuff like "button.onClickSetListener {}..."
    
    } 

activity\_login.xml

    package com.andrycoder.kotlinmessages
    
    import android.content.BroadcastReceiver
    import android.content.Intent
    import android.os.Bundle
    import android.util.Log
    import androidx.appcompat.*
    import androidx.appcompat.app.AppCompatActivity
    
    
    class LoginActivity : AppCompatActivity() {
        override fun onCreate(savedInstanceState: Bundle?) {
            super.onCreate(savedInstanceState)
            setContentView(R.layout.activity_login)
            Log.d("MainActivity Activity", "LoginActivity opened.")
        }
    
        // from here I cannot do stuff like "button.onClickSetListener {}..."
    
    }

AndroidManifest.xml

    &lt;?xml version="1.0" encoding="utf-8"?&gt;
    &lt;manifest xmlns:android="http://schemas.android.com/apk/res/android"
        package="com.andrycoder.kotlinmessages"&gt;
    
        &lt;application
            android:allowBackup="true"
            android:icon="@mipmap/ic_launcher"
            android:label="@string/app_name"
            android:roundIcon="@mipmap/ic_launcher_round"
            android:supportsRtl="true"
            android:theme="@style/AppTheme"&gt;
            &lt;activity android:name=".MainActivity"&gt;
                &lt;intent-filter&gt;
                    &lt;action android:name="android.intent.action.MAIN" /&gt;
    
                    &lt;category android:name="android.intent.category.LAUNCHER" /&gt;
                &lt;/intent-filter&gt;
            &lt;/activity&gt;
            &lt;activity android:name=".LoginActivity" /&gt;
        &lt;/application&gt;
    &lt;/manifest&gt;
## [9][ViewBinding – the New Standard for View Interaction Handling in Android](https://www.reddit.com/r/androiddev/comments/fvx3n8/viewbinding_the_new_standard_for_view_interaction/)
- url: https://infinum.com/the-capsized-eight/viewbinding-the-new-standard-for-view-interaction-handling-in-android
---

## [10][Is libGdx the best framework for Java 2d game development?](https://www.reddit.com/r/androiddev/comments/fwi5gn/is_libgdx_the_best_framework_for_java_2d_game/)
- url: https://www.reddit.com/r/androiddev/comments/fwi5gn/is_libgdx_the_best_framework_for_java_2d_game/
---
Hello, I'm new to the Java &amp; Android development screen and nearly done with my first full app. When I finish that I'd like to try making a simple 2d game and was wondering if anyone here has experience with libGdx?

I'd prefer to stick to a Java framework, because I'm still learning Java and Android, and would rather not try adding learning an engine at the same time.
## [11][Mobile application security testing methodology and key mobile vulnerabilities](https://www.reddit.com/r/androiddev/comments/fwi0nt/mobile_application_security_testing_methodology/)
- url: https://www.immuniweb.com/resources/mobile-application-security-testing/
---

## [12][Decompiling system APK's and modify'ing them to save battery](https://www.reddit.com/r/androiddev/comments/fwhp1c/decompiling_system_apks_and_modifying_them_to/)
- url: https://www.reddit.com/r/androiddev/comments/fwhp1c/decompiling_system_apks_and_modifying_them_to/
---
Has anyone gone to this extend to save battery? If so, what did you do, how much battery did it save, and what advice do you have for me? My battery lasts about a day default, arround 2 days with a few tweaks i made, and i think i can get it to 2,5 to 3 days now. Any tips to save even more battery?
