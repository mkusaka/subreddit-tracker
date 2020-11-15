# Kotlin
## [1][How to avoid passing objects between activities](https://www.reddit.com/r/Kotlin/comments/juiyk4/how_to_avoid_passing_objects_between_activities/)
- url: https://www.reddit.com/r/Kotlin/comments/juiyk4/how_to_avoid_passing_objects_between_activities/
---
I suspect that I am using the wrong code design for my app. I am making a note taking app where each note is an object. In the main activity I create an empty `Notebook` object. In the `Notebook` is an empty mutableListOf of `Notes`s. Each `Note` has the context of the note. (text, date created, date updated ...). When the user clicks on the "Add a note" button it jumps to the CreateNote activity. 

I can think of two ways of adding that note. Either, 
 1. Serialize the `Notebook` object
 2. Send the `Notebook` by using 
```
// In main activity
intent.putExtra(NOTEBOOK, notebook)
// In CreateNote activity
val notebook = intent.getSerializableExtra(NOTEBOOK) as? Notebook
```
 3. Add a newly created note to the notebook
 4. Send back the updated notebook
 5. Update the notebook in main activity.

Alternatively, I can create a new `Note`, serialize it, and send it back. 


Is that the best coding style? If there was a way of storing the `Notebook` object as a shared object between activities I think that would be the best. I could then simply access that object (and potentially send an index if I want to edit a specific note). Is there a way to do that?
## [2][Any doc/article/reference for the new Kotlin's IR (intermediate representation)?](https://www.reddit.com/r/Kotlin/comments/ju1ixb/any_docarticlereference_for_the_new_kotlins_ir/)
- url: https://www.reddit.com/r/Kotlin/comments/ju1ixb/any_docarticlereference_for_the_new_kotlins_ir/
---
My mind was blown after seeing the [kotlin-power-assert](https://github.com/bnorm/kotlin-power-assert), and I'm eager to know about this (IR).

The [sources](https://github.com/JetBrains/kotlin/tree/master/compiler/ir) are so wide (even though well documented almost one third of those), its hard to follow a path i.e. where to see and understand the things.

I was trying to find a way to write a compiler plugin which can modify the intermediate code, so that it will easily target multiplatform, but not able to find any quick-start or something for where to start or start looking at the sources.

If anybody know already about this, he/she can share his experience and help for how to get started (basically I want to know how to traverse all the function/classes and then some way to find linked/called functions, etc).
## [3][Kotlin Coroutines in Spring Webflux](https://www.reddit.com/r/Kotlin/comments/ju7td5/kotlin_coroutines_in_spring_webflux/)
- url: https://medium.com/@tien.donam/asynchronous-spring-using-kotlin-coroutines-and-r2dbc-93b3a079ac22
---

## [4][Help, Why isnt my smile moving? I made the mouse event and the key event but nothing happens](https://www.reddit.com/r/Kotlin/comments/juet5m/help_why_isnt_my_smile_moving_i_made_the_mouse/)
- url: https://www.reddit.com/r/Kotlin/comments/juet5m/help_why_isnt_my_smile_moving_i_made_the_mouse/
---
import pt.isel.canvas.\*  
import java.nio.file.WatchEvent  
import kotlin.math.round  
import kotlin.random.Random  


data class Point(val x: Int, val y: Int)  


data class Smile(val center: Point, val radius: Int, val startAng: Int , val endAng: Int ,val color: Int, val thickness: Int)  


fun createSmile(smile: Smile, x: Int = smile.center.x, y: Int = smile.center.y, radius: Int = smile.radius, startAng: Int = smile.startAng, endAng: Int = smile.endAng, color: Int = smile.color, thickness: Int = smile.thickness): Smile = *createSmile*(x, y, radius, startAng, endAng, color, thickness)  


fun createSmile(x: Int, y: Int, radius: Int, startAng: Int, endAng: Int, color: Int, thickness: Int): Smile {  
 val center = Point(x, y)  
 return Smile(center, radius, startAng, endAng, color, thickness)  


}  


fun drawSmile(smile: Smile, canvas: Canvas): Unit {  
canvas.drawCircle(smile.center.x, smile.center.y, smile.radius, smile.color)  
canvas.drawCircle(smile.center.x, smile.center.y, smile.radius, *BLACK*, smile.thickness)  
canvas.drawArc(smile.center.x, smile.center.y, 30, 200, 340, *BLACK*, smile.thickness)  
canvas.drawCircle(smile.center.x - 17, smile.center.y - 17, 6, *BLACK* )  
canvas.drawCircle(smile.center.x + 17, smile.center.y - 17, 6, *BLACK* )  


}  


fun main() {  
 val WIDTH = 600  
 val HEIGHT = 400  
 val RADIUS = 50  
 *onStart* **{**  
 val cv = Canvas(WIDTH, HEIGHT, *CYAN*)  


var smile = *createSmile*(WIDTH/2, HEIGHT / 2, 50, 0, 360, *YELLOW*, 3)  
 *drawSmile*(smile, cv)  


cv.onMouseDown **{** mouse **-&gt;**  
 smile = *createSmile*(smile, x = mouse.x, y = mouse.y)  


**}**  
 cv.onKeyPressed **{** ke **-&gt;**  
 smile = when (ke.code) {  
 *LEFT\_CODE* \-&gt; *createSmile*(smile, (smile.center.x - 4), smile.center.y)  
 *RIGHT\_CODE* \-&gt; *createSmile*(smile, (smile.center.x + 4), smile.center.y)  
 *UP\_CODE* \-&gt; *createSmile*(smile, smile.center.x, (smile.center.y - 4))  
 *DOWN\_CODE* \-&gt; *createSmile*(smile, smile.center.x, (smile.center.y + 4))  


else -&gt; when (ke.char) {  
 '+' -&gt; *createSmile*(smile, (smile.radius + 1))  
 '-' -&gt; *createSmile*(smile, (smile.radius - 1))  
 else -&gt; smile  
}  


}  


**}**  
**}**  
 *onFinish* **{**  
**}**  
 }
## [5][Plan from a theoretical point of view and implement in Kotlin the solution to the following problem:](https://www.reddit.com/r/Kotlin/comments/juboc8/plan_from_a_theoretical_point_of_view_and/)
- url: https://www.reddit.com/r/Kotlin/comments/juboc8/plan_from_a_theoretical_point_of_view_and/
---
Plan from a theoretical point of view and implement in Java the solution to the following problem:

We want to know how many different ways exist to obtain certain integer number as a sum of smaller positive numbers, taking into account a certain list of numbers that cannot be used.

Considerations

• If A + B = B + A, only one of them should be indicated

• In addition, A &gt; 0 and B &gt; 0

Input

Each case is described in one line with an integer number X, which is the value to obtain, and a list of M positive numbers N0, N1 ... NM-1 separated by a space. These values cannot be used as operands.

Output

For each case T, the output must be the string "#t:" followed by the total number of possible ways. For example

Sample input

3 1

6 1 3

Sample output

\#1: 0

\#2: 2

In the first case (3), there is no way to get 3 from 1 and 2 without using operand 1. The second case (6) can be calculated with operations 2 + 2 + 2 and 4 + 2, so there are 2 possible ways without using operands 1 and 3.

The objective is to create the most efficient algorithm able to perform this calculation

Any tips )
## [6][Hey #Kotliners💡, ever wondered how #Coroutines work under the hood?🤔 Cool! This November, Mohit Sarveiya (Android GDE) will dissect 🗡️ coroutines and take us through some of the lesser known and advanced features of #Kotlin coroutines! RSVP](https://www.reddit.com/r/Kotlin/comments/jtkdnq/hey_kotliners_ever_wondered_how_coroutines_work/)
- url: https://www.meetup.com/Kotlin-Mumbai/events/274576859/
---

## [7][Sonarnet : Android library for seamless awareness of true Internet access and captive portals. [Feedback are welcome for improvement]](https://www.reddit.com/r/Kotlin/comments/jtkn5c/sonarnet_android_library_for_seamless_awareness/)
- url: https://github.com/fabricethilaw/sonarnet
---

## [8][hello, I need a bit of help with threads and handling serial data](https://www.reddit.com/r/Kotlin/comments/jts3ey/hello_i_need_a_bit_of_help_with_threads_and/)
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
## [9][I want to develop a custom app for my DJI Drone. DJI SDK is in Java. However, i'm attracted to learning Kotlin rather than Java as it seems the future (for Android). So will i be able to develop an app for my drone in Kotlin?](https://www.reddit.com/r/Kotlin/comments/jtdvp7/i_want_to_develop_a_custom_app_for_my_dji_drone/)
- url: https://www.reddit.com/r/Kotlin/comments/jtdvp7/i_want_to_develop_a_custom_app_for_my_dji_drone/
---
Hi guys, I'm an absolute beginner Android dev, with an end goal to develop a custom app for my DJI drone. I know very little Java &amp; none of Kotlin. The [DJI SDK](https://github.com/dji-sdk/Mobile-SDK-Android) is written in Java.

With my end goal in mind, will it be possible to make a custom app for my drone in Kotlin? Or do I need to stick to Java since the DJI SDK itself is in Java?

Basis the answer to the above question I intend to start learning Android Studio tutorials either specifically for Java or Kotlin. Any guidance would be appreciated, thank you!

EDIT - THANKS A LOT guys, appreciate all the answers, look forward to coding in Kotlin!
## [10][New kotlin newsletter - kotlin.news](https://www.reddit.com/r/Kotlin/comments/jtkmu9/new_kotlin_newsletter_kotlinnews/)
- url: https://www.reddit.com/r/Kotlin/comments/jtkmu9/new_kotlin_newsletter_kotlinnews/
---
Hi I am Android Developer and huge Kotlin enthiusiast (before I start any pet project I am checking if it is possible to be done in Kotlin). I am trying a lot kotlin frameworks, so I thought that I could share my experiences with them. Currently I am trying Compose framework for web development in kotlin - Wanna learn about it? Join at [https://kotlin.news](https://kotlin.news)  
If it does not convince You - I am also creating Intellij Plugins (recently created project wizard for kVision framework) so maybe some of newly released plugins will convince You? Anyway, it is Your choice, You know where to find me - [https://kotlin.news](https://kotlin.news)
