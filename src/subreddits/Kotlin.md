# Kotlin
## [1][New to Kotlin](https://www.reddit.com/r/Kotlin/comments/ftjw23/new_to_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/ftjw23/new_to_kotlin/
---
Hey,

I am currently studying Kotlin but having a hard time since I feel every tutorial I find somehow expects you to be a bit experienced in Java. In which I have zero experience. Are there any books which really offer the reader an extensive overview of the language? I prefer reading over videos.  
Thanks!
## [2][Math solver android app question](https://www.reddit.com/r/Kotlin/comments/ftkzsu/math_solver_android_app_question/)
- url: https://www.reddit.com/r/Kotlin/comments/ftkzsu/math_solver_android_app_question/
---
I am thinking of a making a math solver app, mostly linear algebra stuff, matrixes, vectors etc.

If i were to make something like symbolab/wolfram (much less complex, only some features), would it be resonable to do the computation on the mobile app, or should i do it with api calls to a backend?

Thank you :)
## [3][Setting up Google Authentication in a Kotlin Android App](https://www.reddit.com/r/Kotlin/comments/ft1wft/setting_up_google_authentication_in_a_kotlin/)
- url: https://codesource.io/setting-up-google-authentication-in-a-kotlin-android-app/
---

## [4][`scan()` and `scanReduce()` in Kotlin 1.3.70!](https://www.reddit.com/r/Kotlin/comments/fsywl7/scan_and_scanreduce_in_kotlin_1370/)
- url: https://twitter.com/kotlin/status/1245300637916594179?s=20
---

## [5][🎉 GraphQL Kotlin 2.0.0!](https://www.reddit.com/r/Kotlin/comments/fsru6c/graphql_kotlin_200/)
- url: https://github.com/ExpediaGroup/graphql-kotlin/releases/tag/2.0.0
---

## [6][Async Coroutine error handling](https://www.reddit.com/r/Kotlin/comments/ft1ro1/async_coroutine_error_handling/)
- url: https://www.reddit.com/r/Kotlin/comments/ft1ro1/async_coroutine_error_handling/
---
I'd like to say "all good 'til Coroutines".  
Jokes asides, in Java I'm used to `CompletableFuture`s.  
So I was using that one in Kotlin too

```
myList
  .map {
    CompletableFuture
      .supplyAsync { it.fn() }
      .orTimeout(1, SECONDS)
      .exceptionally { emptyList() }
  }
  .toList()
  .flatMap { it.join() }

```

Now, in Kotlin, all what I could come up is

```
myList
  .map {
    try {
      withTimeout(1000) {
        async {
          it.fn()
        }
      }
    } catch (e: Exception) {
      async { emptyList() }
    }
  }
  .awaitAll()
  .asSequence()
  .flatten()
```

Honestly this looks a bit weird on a functional flow.  
What am I missing?
## [7][[Help] View/Button ins't clickable](https://www.reddit.com/r/Kotlin/comments/ft5vwi/help_viewbutton_inst_clickable/)
- url: https://www.reddit.com/r/Kotlin/comments/ft5vwi/help_viewbutton_inst_clickable/
---
So, im making an for the purpose of calculating numbers in a game while the user plays the game.

For that i've made a floating button that draws over all apps.

The issue im facing is that the button isn't clickable, and when i set a listener it crashes.

Heres the code.. thanks for help.

MainActivity.kt

    package com.example.rokcalkot
    
    import android.content.Intent
    import android.app.Activity
    import android.os.Build
    import android.provider.Settings
    import android.view.View
    import android.content.pm.PackageManager
    import android.os.Bundle
    import android.util.Log
    import android.widget.Button
    import android.widget.Toast
    import androidx.appcompat.app.AlertDialog
    import androidx.appcompat.app.AppCompatActivity
    import android.net.Uri
    import androidx.core.app.ActivityCompat
    import androidx.core.content.ContextCompat
    import kotlinx.android.synthetic.main.onscreen.*
    
    import android.graphics.Canvas
    import android.graphics.Color
    import android.graphics.Bitmap
    
    
    class MainActivity : AppCompatActivity() {
    
        val requestCode = 201
    
        override fun onCreate(savedInstanceState: Bundle?) {
            super.onCreate(savedInstanceState)
            setContentView(R.layout.activity_main)
    
            //-------------------------- LAUNCH BUTTON BTN ---------------------------
            val btnL = findViewById&lt;Button&gt;(R.id.btnShow)
            val btnS= findViewById&lt;Button&gt;(R.id.btnService)
    
            /*btnS.setOnClickListener{
                Toast.makeText(this@MainActivity, "CLICKBLE", Toast.LENGTH_LONG).show()
            }*/
            //btnShow?.setOnClickListener { Toast.makeText(this@MainActivity, "Bannaa", Toast.LENGTH_LONG).show() }
            btnL.setOnClickListener {
                if (Build.VERSION.SDK_INT &gt;= Build.VERSION_CODES.M &amp;&amp; !Settings.canDrawOverlays(this)) {
                    val intent = Intent(
                        Settings.ACTION_MANAGE_OVERLAY_PERMISSION,
                        Uri.parse("package:$packageName")
                    )
    
                    startActivityForResult(intent, requestCode)
                    //startActivityForResult(intent2,requestCode1)
                } else {
                    startService(Intent(this, MyFloatingWindowService::class.java))
    
                    finish()
                }
                //Toast.makeText(this@MainActivity, "Bannaa", Toast.LENGTH_LONG).show()
                }
    
            fun onActivityResult(requestCode: Int, resultCode: Int, data: Intent?) {
                super.onActivityResult(requestCode, resultCode, data)
                if (resultCode == Activity.RESULT_OK) {
                    startService(Intent(this, MyFloatingWindowService::class.java))
                    finish()
                }
            }
        }
    }

&amp;#x200B;

Activity\_Main.xml

    &lt;?xml version="1.0" encoding="utf-8"?&gt;
    &lt;androidx.constraintlayout.widget.ConstraintLayout xmlns:android="http://schemas.android.com/apk/res/android"
        xmlns:app="http://schemas.android.com/apk/res-auto"
        xmlns:tools="http://schemas.android.com/tools"
        android:layout_width="match_parent"
        android:layout_height="match_parent"
        android:theme="@style/AppTheme"
        tools:context=".MainActivity"&gt;
    
        &lt;Button
            android:id="@+id/btnShow"
            android:layout_width="95dp"
            android:layout_height="96dp"
            android:text="Start Rok Cal"
            app:layout_constraintBottom_toBottomOf="parent"
            app:layout_constraintEnd_toEndOf="parent"
            app:layout_constraintHorizontal_bias="0.5"
            app:layout_constraintStart_toStartOf="parent"
            app:layout_constraintTop_toTopOf="parent" /&gt;
    
        &lt;TextView
            android:id="@+id/textView"
            android:layout_width="141dp"
            android:layout_height="45dp"
            android:text="Welcome to Rok Cal"
            app:layout_constraintBottom_toTopOf="@+id/btnShow"
            app:layout_constraintEnd_toEndOf="parent"
            app:layout_constraintHorizontal_bias="0.5"
            app:layout_constraintStart_toStartOf="parent" /&gt;
    
    &lt;/androidx.constraintlayout.widget.ConstraintLayout&gt;

&amp;#x200B;

MyFloatingWindowService.kt

    package com.example.rokcalkot
    import android.annotation.SuppressLint
    import android.app.Service
    import android.content.Context
    import android.content.Intent
    import android.graphics.PixelFormat
    import android.os.Build
    import android.os.IBinder
    import android.view.*
    import android.widget.Toast
    import kotlin.math.roundToInt
    
    class MyFloatingWindowService : Service() {
    
        var LAYOUT_FLAG: Int = 0
    
        lateinit var floatingView: View
        lateinit var floatingViewMain: View
        lateinit var manager: WindowManager
        lateinit var params: WindowManager.LayoutParams
        override fun onBind(intent: Intent?): IBinder? {
            TODO("not implemented") //To change body of created functions use File | Settings | File Templates.
    
        }
    
    
        override fun onStartCommand(intent: Intent?, flags: Int, startId: Int): Int {
    
            LAYOUT_FLAG = if (Build.VERSION.SDK_INT &gt;= Build.VERSION_CODES.O) {
                WindowManager.LayoutParams.TYPE_APPLICATION_OVERLAY
            } else {
                WindowManager.LayoutParams.TYPE_PHONE
            }
            val params = WindowManager.LayoutParams(
                WindowManager.LayoutParams.WRAP_CONTENT,
                WindowManager.LayoutParams.WRAP_CONTENT,
                LAYOUT_FLAG,
                WindowManager.LayoutParams.FLAG_NOT_FOCUSABLE,
                PixelFormat.TRANSLUCENT
            )
    
            this.params = params
            //Specify the view position
            params.gravity =
                Gravity.TOP or Gravity.CENTER //Initially view will be added to top-left corner
            params.x = 0
            params.y = 0
    
    
            manager = getSystemService(Context.WINDOW_SERVICE) as WindowManager
            floatingView = LayoutInflater.from(this).inflate(R.layout.onscreen, null)
            floatingViewMain = LayoutInflater.from(this).inflate(R.layout.activity_main, null)
            manager.addView(floatingView, params)
            floatingView.findViewById&lt;View&gt;(R.id.btnService)?.setOnTouchListener(object :
                View.OnTouchListener {
                var initialX: Int? = null
                var initialY: Int? = null
                var initialTouchX: Float? = null
                var initialTouchY: Float? = null
    
    
                @SuppressLint("ClickableViewAccessibility")
                override fun onTouch(view: View?, motionEvent: MotionEvent?): Boolean {
                    //Toast.makeText(this@MyFloatingWindowService, "Clickble", Toast.LENGTH_SHORT).show()
                    when (motionEvent!!.action) {
                        MotionEvent.ACTION_DOWN -&gt; {
                            //remember the initial position.
                            initialX = params.x
                            initialY = params.y
    
    
                            //get the touch location
                            //Toast.makeText(this@MyFloatingWindowService, "ACTION DOWN", Toast.LENGTH_SHORT).show()
                            initialTouchX = motionEvent!!.getRawX()
                            initialTouchY = motionEvent!!.getRawY()
                            return true
                        }
                        MotionEvent.ACTION_UP -&gt; {
                            //Toast.makeText(this@MyFloatingWindowService, "ACTION UP", Toast.LENGTH_SHORT).show()
                            val Xdiff = (motionEvent.getRawX() - initialTouchX!!)
                            val Ydiff = (motionEvent.getRawY() - initialTouchY!!)
                            return true
                        }
                        MotionEvent.ACTION_MOVE -&gt; {
                            //Calculate the X and Y coordinates of the view.
    
                            params.x = initialX!!.plus((motionEvent.getRawX() - initialTouchX!!)).roundToInt()
                            params.y = initialY!!.plus((motionEvent.getRawY() - initialTouchY!!).roundToInt())
                            manager.updateViewLayout(floatingView, params)
                            return true
                        }
    
                    }
                    return false
                }
            })
    
    
            return START_NOT_STICKY
        }
    
    
    
        fun CHECK() {
            Toast.makeText(this@MyFloatingWindowService, "Clickble", Toast.LENGTH_LONG).show()
        }
    
        override fun onDestroy() {
            super.onDestroy()
            manager.removeView(floatingView)
        }
    
    
    }

&amp;#x200B;

onscreen.xml --&gt; the xml that the service interacts with which draws over all apps.

    &lt;?xml version="1.0" encoding="utf-8"?&gt;
    &lt;androidx.constraintlayout.widget.ConstraintLayout xmlns:android="http://schemas.android.com/apk/res/android"
        xmlns:app="http://schemas.android.com/apk/res-auto"
        xmlns:tools="http://schemas.android.com/tools"
        android:layout_width="match_parent"
        android:layout_height="match_parent"
        android:theme="@style/AppTheme"
        tools:context=".MainActivity"
        android:launchMode="singleTask"&gt;
    
        &lt;Button
            android:id="@+id/btnService"
            android:layout_width="51dp"
            android:layout_height="62dp"
            android:clickable="true"
            android:text="Start Service"
            app:layout_constraintBottom_toBottomOf="parent"
            app:layout_constraintEnd_toEndOf="parent"
            app:layout_constraintHorizontal_bias="0.0"
            app:layout_constraintStart_toStartOf="parent"
            app:layout_constraintTop_toTopOf="parent"
            app:layout_constraintVertical_bias="0.522" /&gt;
    
    &lt;/androidx.constraintlayout.widget.ConstraintLayout&gt;
## [8][Kotlin Coroutined](https://www.reddit.com/r/Kotlin/comments/ft5i2s/kotlin_coroutined/)
- url: https://medium.com/@shelbyc0hen/kotlin-coroutined-8c57acd3c45e
---

## [9][“How to Combine Kotlin Flows” #androiddev #androiddevelopment #kotlin #roomdatabase #kotlin_multiplatform #programming #androiddeveloper](https://www.reddit.com/r/Kotlin/comments/fssg5b/how_to_combine_kotlin_flows_androiddev/)
- url: https://medium.com/better-programming/learn-how-to-combine-kotlin-flows-317849a71d3e?source=friends_link&amp;sk=fc6ee71b885f67afe4dae695cb179a0b
---

## [10][Why does Kotlin look so good ? Am I missing something?](https://www.reddit.com/r/Kotlin/comments/frr50s/why_does_kotlin_look_so_good_am_i_missing/)
- url: https://www.reddit.com/r/Kotlin/comments/frr50s/why_does_kotlin_look_so_good_am_i_missing/
---
Hello there, 

 So I am software engineer and I started with C++, transitioned to Java in Uni, and I used a lot of Perl and Bash during my job that started around 6 months ago. I started developing Android apps about 1 year ago in Java, and never looked at Kotlin. A couple of days ago, I finally gave it a chance before starting a new project (might as well start with a new programming language as well). 

 I am amazed by the syntax and the language features. Everything that I like in C-like languages and that is good in scripting languages is combined in Kotlin. I even had an idea to implement a new language of my own (just as a learning exercise) and some stuff that I wanted to bring "new" to the table, they are implemented in Kotlin. 

 My questions are, is it actually that much better of a coding experience as it looks? What are some drawbacks that you experience ? Also, why do you thing it's not that popular ? 

&amp;#x200B;

Just joined the subreddit and I am still having a look around so please be kind :)
