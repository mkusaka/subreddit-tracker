# Kotlin
## [1][Any doc/article/reference for the new Kotlin's IR (intermediate representation)?](https://www.reddit.com/r/Kotlin/comments/ju1ixb/any_docarticlereference_for_the_new_kotlins_ir/)
- url: https://www.reddit.com/r/Kotlin/comments/ju1ixb/any_docarticlereference_for_the_new_kotlins_ir/
---
My mind was blown after seeing the [kotlin-power-assert](https://github.com/bnorm/kotlin-power-assert), and I'm eager to know about this (IR).

The [sources](https://github.com/JetBrains/kotlin/tree/master/compiler/ir) are so wide (even though well documented almost one third of those), its hard to follow a path i.e. where to see and understand the things.

I was trying to find a way to write a compiler plugin which can modify the intermediate code, so that it will easily target multiplatform, but not able to find any quick-start or something for where to start or start looking at the sources.

If anybody know already about this, he/she can share his experience and help for how to get started (basically I want to know how to traverse all the function/classes and then some way to find linked/called functions, etc).
## [2][Hey #Kotliners💡, ever wondered how #Coroutines work under the hood?🤔 Cool! This November, Mohit Sarveiya (Android GDE) will dissect 🗡️ coroutines and take us through some of the lesser known and advanced features of #Kotlin coroutines! RSVP](https://www.reddit.com/r/Kotlin/comments/jtkdnq/hey_kotliners_ever_wondered_how_coroutines_work/)
- url: https://www.meetup.com/Kotlin-Mumbai/events/274576859/
---

## [3][Sonarnet : Android library for seamless awareness of true Internet access and captive portals. [Feedback are welcome for improvement]](https://www.reddit.com/r/Kotlin/comments/jtkn5c/sonarnet_android_library_for_seamless_awareness/)
- url: https://github.com/fabricethilaw/sonarnet
---

## [4][hello, I need a bit of help with threads and handling serial data](https://www.reddit.com/r/Kotlin/comments/jts3ey/hello_i_need_a_bit_of_help_with_threads_and/)
- url: https://www.reddit.com/r/Kotlin/comments/jts3ey/hello_i_need_a_bit_of_help_with_threads_and/
---
 Hello! I am currently working on acheiving certain tasks depending on the data received from a Bluetooth module using the input buffer. Currently I am trying to play a sound depending on the input using MediaPlayer.Create, but only seems to work when I have the device plugged in and hit "Apply Changes and Restart Activity." It does not seem to work if I were to just open the app and test, only after going through the process stated before. Is there a better way of reading the inputstream and handling that data? Is there also a better way to go about how I am playing a sound in the run() function of the ConnectedThread thread? This is my current code. I am also thinking of switching from MediaPlayer to SoundPool because I would like to play multiple sounds simultaneously, would that also be a good option?   


    class ControlActivity: AppCompatActivity() {
     
       companion object {
           var m_myUUID: UUID = UUID.fromString("00001101-0000-1000-8000-00805F9B34FB")
           var m_bluetoothSocket: BluetoothSocket? = null
           lateinit var m_progress: ProgressDialog
           lateinit var m_bluetoothAdapater: BluetoothAdapter
           var m_isConnected: Boolean = false
           lateinit var  m_address: String
           private const val TAG = "MY_APP_DEBUG_TAG"
           private lateinit var handler: Handler
       }
       var mMediaPlayer: MediaPlayer? = null
       override fun onCreate(savedInstanceState: Bundle?) {
     
           super.onCreate(savedInstanceState)
           setContentView(R.layout.control_layout)
           m_address = intent.getStringExtra(SettingsActivity.EXTRA_ADDRESS)!!
           ConnectToDevice(this).execute()
     
           control_led_on.setOnClickListener { sendCommand("1") }
           control_led_off.setOnClickListener { sendCommand("0") }
           control_led_disconnect.setOnClickListener { disconnect() }
       }
     
       private fun sendCommand(input: String){
           if (m_bluetoothSocket != null){
               try {
                   m_bluetoothSocket!!.outputStream.write(input.toByteArray())
                   Log.i("data", "sending..")
               } catch (e: IOException) {
                   e.printStackTrace()
                   Log.i("data", "couldn't send")
                   }
                   return
     
               }
     
       }
       private fun disconnect(){
           if (m_bluetoothSocket != null){
               try {
                   m_bluetoothSocket!!.close()
                   m_bluetoothSocket = null
                   m_isConnected = false
               } catch (e: IOException) {
                   e.printStackTrace()
               }
           }
           finish()
       }
     
       private class ConnectedThread(private val mmSocket: BluetoothSocket?, c : Context) : Thread() {
           private val mmBuffer: ByteArray = ByteArray(10) // mmBuffer store for the stream
           var mMediaPlayer: MediaPlayer? = null
           var context: Context
           init{
               this.context = c
           }
           override fun run() {
               var numBytes: Int = 0// bytes returned from read()
               //var begin: Int = 0
               // Keep listening to the InputStream until an exception occurs.
               while (true) {
                   // Read from the InputStream.
                   try {
                       numBytes =
                           mmSocket!!.inputStream.read(mmBuffer, numBytes, mmBuffer.size - numBytes)
     
     
                   } catch (e: IOException) {
                       Log.d(TAG, "Input stream was disconnected", e)
                       break
                   }
     
                   playNote(context)
     
                   println(numBytes)
                   println(String(mmBuffer))
     
     
     
     
     
               }
           }
           fun playNote(c: Context) {
               if (mMediaPlayer == null) {
                   mMediaPlayer = MediaPlayer.create(c, R.raw.c3)
                   mMediaPlayer!!.start()
               } else if (mMediaPlayer != null) {
                   mMediaPlayer!!.stop()
                   mMediaPlayer!!.release()
                   mMediaPlayer = null
               } else mMediaPlayer!!.start()
           }
     
       }
     
     
     
     
       private class ConnectToDevice(c: Context) : AsyncTask&lt;Void, Void, String&gt;(){
           private var connectSuccess: Boolean = true
           private val context: Context
     
           init {
               this.context = c
           }
           override fun onPreExecute() {
               super.onPreExecute()
               m_progress = ProgressDialog.show(context, "Connecting...", "please wait")
           }
           override fun doInBackground(vararg p0: Void?) : String? {
               try {
                   if (m_bluetoothSocket == null || !m_isConnected){
                       m_bluetoothAdapater = BluetoothAdapter.getDefaultAdapter()
                       val device: BluetoothDevice = m_bluetoothAdapater.getRemoteDevice(m_address)
                       m_bluetoothSocket = device.createInsecureRfcommSocketToServiceRecord(m_myUUID)
                       BluetoothAdapter.getDefaultAdapter().cancelDiscovery()
                       m_bluetoothSocket!!.connect()
     
                   }
               } catch (e: IOException){
                   connectSuccess = false
                   e.printStackTrace()
               }
               return null
           }
           override fun onPostExecute(result: String?) {
               super.onPostExecute(result)
               if(!connectSuccess){
                   Log.i("data", "couldn't connect")
               } else {
                   m_isConnected = true
                   var connectedthread= ConnectedThread(m_bluetoothSocket, context)
                   connectedthread.start()
                   Log.i("data", "connected")
               }
               m_progress.dismiss()
     
     
           }
     
       }
    }

 I feel like the problem I have might be the way I am passing the context to the MediaPlayer in the ConnectedThread thread, but I'm honestly not too sure because i just started learning kotlin. What other options do I have for evaluating the bytearray being sent from the Bluetooth?
## [5][I want to develop a custom app for my DJI Drone. DJI SDK is in Java. However, i'm attracted to learning Kotlin rather than Java as it seems the future (for Android). So will i be able to develop an app for my drone in Kotlin?](https://www.reddit.com/r/Kotlin/comments/jtdvp7/i_want_to_develop_a_custom_app_for_my_dji_drone/)
- url: https://www.reddit.com/r/Kotlin/comments/jtdvp7/i_want_to_develop_a_custom_app_for_my_dji_drone/
---
Hi guys, I'm an absolute beginner Android dev, with an end goal to develop a custom app for my DJI drone. I know very little Java &amp; none of Kotlin. The [DJI SDK](https://github.com/dji-sdk/Mobile-SDK-Android) is written in Java.

With my end goal in mind, will it be possible to make a custom app for my drone in Kotlin? Or do I need to stick to Java since the DJI SDK itself is in Java?

Basis the answer to the above question I intend to start learning Android Studio tutorials either specifically for Java or Kotlin. Any guidance would be appreciated, thank you!

EDIT - THANKS A LOT guys, appreciate all the answers, look forward to coding in Kotlin!
## [6][New kotlin newsletter - kotlin.news](https://www.reddit.com/r/Kotlin/comments/jtkmu9/new_kotlin_newsletter_kotlinnews/)
- url: https://www.reddit.com/r/Kotlin/comments/jtkmu9/new_kotlin_newsletter_kotlinnews/
---
Hi I am Android Developer and huge Kotlin enthiusiast (before I start any pet project I am checking if it is possible to be done in Kotlin). I am trying a lot kotlin frameworks, so I thought that I could share my experiences with them. Currently I am trying Compose framework for web development in kotlin - Wanna learn about it? Join at [https://kotlin.news](https://kotlin.news)  
If it does not convince You - I am also creating Intellij Plugins (recently created project wizard for kVision framework) so maybe some of newly released plugins will convince You? Anyway, it is Your choice, You know where to find me - [https://kotlin.news](https://kotlin.news)
## [7][Kotlin for loop](https://www.reddit.com/r/Kotlin/comments/jtjx6j/kotlin_for_loop/)
- url: https://www.reddit.com/r/Kotlin/comments/jtjx6j/kotlin_for_loop/
---
Hello there, I've recently learned Basic kotlin and I'm switching from Java. I am unable to write kotlin code for :-

 for (int i = 1; i &lt; t; i++) {

for (int j = i; (i\*i) + (j\*j) &lt;= (t\*t); j++) 

P.S. If you have some tips for writing Kotlin please share as I'm just converting java into kotlin
## [8][Is there a way to do nested destructuring?](https://www.reddit.com/r/Kotlin/comments/jt37lg/is_there_a_way_to_do_nested_destructuring/)
- url: https://www.reddit.com/r/Kotlin/comments/jt37lg/is_there_a_way_to_do_nested_destructuring/
---
There are two places where it comes up quasi-frequently for me: lambda arguments, and unwrapping data classes.

Example 1:

    data class Foo (val s: String, val i: Int)
    data class Bar (val f: Foo)
    
    fun doStuff () {
        val (f) = makeBar()
        val (_, i) = f
        useI(i)
    }
    
    fun useI(i: Int) {
        TODO()
    }
    
    fun makeBar(): Bar {
        TODO()
    }

I'd like it if I could've written `val ((_, i)) = makeBar()` instead of having `f` as an intermediate.

Example 2:

    data class Foo (val s: String, val i: Int)
    
    fun doStuff () {
        val sum = makeListOfFoo().fold(0) { acc, foo -&gt;
            acc + foo.i
        }
        
        println(sum)
    }
    
    fun makeListOfFoo(): List&lt;Foo&gt; {
        TODO()
    }

I'd like it if I could write my lambda args like `fold(acc) { acc, (_, i) -&gt; `

None of that is possible, right?
## [9][Kotlin Compose + web](https://www.reddit.com/r/Kotlin/comments/jsiggw/kotlin_compose_web/)
- url: https://www.reddit.com/r/Kotlin/comments/jsiggw/kotlin_compose_web/
---
Since Kotlin supports JS, is it possible for Kotlin Compose to be ported to web? (I was thinking whether is possible to match Flutter's platform set)
## [10][The best books to learn Kotlin](https://www.reddit.com/r/Kotlin/comments/jskqd6/the_best_books_to_learn_kotlin/)
- url: https://medium.com/p/the-best-books-to-learn-kotlin-894f9a83e64a
---

