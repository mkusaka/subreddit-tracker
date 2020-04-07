# Kotlin
## [1][Migrating Duolingo’s Android app to 100% Kotlin](https://www.reddit.com/r/Kotlin/comments/fw3wkt/migrating_duolingos_android_app_to_100_kotlin/)
- url: https://blog.duolingo.com/migrating-duolingos-android-app-to-100-kotlin/
---

## [2][Rawky v0.19.5.15-alpha has just been released!](https://www.reddit.com/r/Kotlin/comments/fw8i2r/rawky_v019515alpha_has_just_been_released/)
- url: /r/rawky/comments/fw8fm5/rawky_v019515alpha_has_just_been_released/
---

## [3][Will Kotlin ever be able to compile to a stand-alone exe?](https://www.reddit.com/r/Kotlin/comments/fwch4x/will_kotlin_ever_be_able_to_compile_to_a/)
- url: https://www.reddit.com/r/Kotlin/comments/fwch4x/will_kotlin_ever_be_able_to_compile_to_a/
---
I really like the syntax of Kotlin and I'm wondering if it'll ever be possible to easily compile to a stand-alone executable (like you'd get with "gcc myprog.c").  I'm aware I can do things like append the .jar to a shell script to create an "executable", but that feels pretty clunky.
## [4][Help with Rest Api for noob](https://www.reddit.com/r/Kotlin/comments/fvurqz/help_with_rest_api_for_noob/)
- url: https://www.reddit.com/r/Kotlin/comments/fvurqz/help_with_rest_api_for_noob/
---
Hey. I am asking for help what I am doing wrong. I have recently been learning the Kotlin for Andorid. I'm trying to make a simple application that connects to api to download a list of beers, their photos, name and description. I'm trying to connect to this api: https://punkapi.com/documentation/v2 

[https://pastebin.com/hf2HdpPK](https://pastebin.com/hf2HdpPK) \-interface

[https://pastebin.com/q0DAQynn](https://pastebin.com/q0DAQynn) \- class Response

[https://pastebin.com/0JKBXQ0w](https://pastebin.com/0JKBXQ0w) \- class Beer

[https://pastebin.com/XLXvQazj](https://pastebin.com/XLXvQazj) \- RecyclerView Adapter

[https://pastebin.com/0cbcVatX](https://pastebin.com/0cbcVatX) \-MainActivity

[https://pastebin.com/2bnPF6gP](https://pastebin.com/2bnPF6gP) MainViewModel

[https://pastebin.com/jWYU3sJX](https://pastebin.com/jWYU3sJX) \- Repository
## [5][[Question] How to change my helper function so that is collects the results of the parallel processing tasks](https://www.reddit.com/r/Kotlin/comments/fvxd1a/question_how_to_change_my_helper_function_so_that/)
- url: https://www.reddit.com/r/Kotlin/comments/fvxd1a/question_how_to_change_my_helper_function_so_that/
---
[https://stackoverflow.com/questions/61058974/how-to-change-my-helper-function-so-that-is-collects-the-results-of-the-parallel](https://stackoverflow.com/questions/61058974/how-to-change-my-helper-function-so-that-is-collects-the-results-of-the-parallel)

&amp;#x200B;

I would really appreciate if anybody has a suggestion for me with this issue. (sorry for not copying the whole question over here, but I am on mobile and code snippets are hard to do without markdown on reddit)
## [6][Do you use local functions?](https://www.reddit.com/r/Kotlin/comments/fvlrtb/do_you_use_local_functions/)
- url: https://www.reddit.com/r/Kotlin/comments/fvlrtb/do_you_use_local_functions/
---
I know this question has been asked before, but not very recently, and I'm curious of whether people's views have changed here.

Do you use local functions? Personally, I like the idea of a function being hidden inside another function, but I don't like Kotlin's implementation of local functions. In particular, I dislike the fact that they must be defined *before* they are used, and I also don't like how much bytecode they generate compared to a normal function.

Here's an example of a scenario which might call for local functions. I'm making a mathematics game, using MVP, and my presenter (which responds to the user pressing buttons) has code like this:

    class FooPresenter {
        // ...
    
        fun onDigitClick(digit: Char) {
            inputSoFar += digit
            view.showInput(inputSoFar)
            if (inputSoFar.length == answerLength) {
                if (inputSoFar.toInt() == expression.evaluation) {
                    onCorrectAnswer()
                } else {
                    onIncorrectAnswer()
                }
            }
        }
    
        private fun onCorrectAnswer() {
            numQuestionsAnswered++
            if (numQuestionsAnswered &lt; numQuestions) {
                expression = expressionGenerator.generate()
                view.showExpression(expression)
                clearViewInput()
            } else {
                view.onGameOver()
            }
        }
    
        private fun onIncorrectAnswer() {
            clearViewInput()
        }
    }

I like the idea of making `onCorrectAnswer` and `onIncorrectAnswer` local functions, because they will only ever be called from `onDigitClick`, but for the reasons given above I'm reluctant to.
## [7][Coroutines and callbacks](https://www.reddit.com/r/Kotlin/comments/fvd35o/coroutines_and_callbacks/)
- url: https://www.coroutinedispatcher.com/2020/04/coroutines-and-callbacks.html
---

## [8][Unable to "interact(?)" with buttons etc... outside of the MainAcitivty](https://www.reddit.com/r/Kotlin/comments/fvxkd1/unable_to_interact_with_buttons_etc_outside_of/)
- url: https://www.reddit.com/r/Kotlin/comments/fvxkd1/unable_to_interact_with_buttons_etc_outside_of/
---
Title. ( And please see the comment in LoginActivity.kt) Here are the codes:

MainActivity.kt

    package com.andrycoder.kotlinmessages
    
    import android.content.Intent
    import androidx.appcompat.app.AppCompatActivity
    import android.os.Bundle
    import android.util.Log
    import kotlinx.android.synthetic.main.activity_login.*
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
## [9][Boilerplate free proxy fakes](https://www.reddit.com/r/Kotlin/comments/fvb2z1/boilerplate_free_proxy_fakes/)
- url: https://www.reddit.com/r/Kotlin/comments/fvb2z1/boilerplate_free_proxy_fakes/
---
Where do you stand on mocks vs fakes? Would you go for more fakes if you had less boilerplate and noisy code?

Inspired by retrofit, I tried out a proxy based approach to solve this issue and make fakes a bit more pleasant. Have a look below.

“Boilerplate free proxy fakes” by Fanis Veizis https://link.medium.com/Rql07BXVq5
## [10][Template-Oriented-Programming to Ship Faster | Kotlin, Arrow | KotlinHyderabad](https://www.reddit.com/r/Kotlin/comments/fvd5dh/templateorientedprogramming_to_ship_faster_kotlin/)
- url: https://youtu.be/_QBlKtUY6ac
---

