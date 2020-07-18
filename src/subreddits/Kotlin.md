# Kotlin
## [1][Firebase competitors.](https://www.reddit.com/r/Kotlin/comments/htfidl/firebase_competitors/)
- url: https://www.reddit.com/r/Kotlin/comments/htfidl/firebase_competitors/
---
I'm interested in building my first Android app using Kotlin.  Base on my research thus far, Firebase seem to be a good solution to help minimize some of the development start up work.

--

I was wondering if there are alternatives / competitors to Firebase that I should do some research on before committing to Firebase.

--

Thanks.
## [2][Java Slack SDK 1.1.0 includes a Kotlin DSL for building messages](https://www.reddit.com/r/Kotlin/comments/hswsr4/java_slack_sdk_110_includes_a_kotlin_dsl_for/)
- url: https://slack.dev/java-slack-sdk/guides/composing-messages#block-kit-kotlin-dsl
---

## [3][Learn Android and Kotlin with no programming experience (google)](https://www.reddit.com/r/Kotlin/comments/ht35ju/learn_android_and_kotlin_with_no_programming/)
- url: https://android-developers.googleblog.com/2020/07/learn-android-and-kotlin-with-no-experience.html
---

## [4][Framework similar to Blazor/Meteor/LiveView?](https://www.reddit.com/r/Kotlin/comments/ht0o9w/framework_similar_to_blazormeteorliveview/)
- url: https://www.reddit.com/r/Kotlin/comments/ht0o9w/framework_similar_to_blazormeteorliveview/
---
Do we have some framework like C#'s Blazor, or JS's Meteor or Elixir's Phoenix LiveView in Kotlin or JVM worlds? In a sence that framework will do automatic client side and server side communications using WebSockets for me. While still allowing me to write templates in HTML and CSS.

I'd prefer some framework with embedded server (like Javalin for example) so I could deploy my web app as a single jar-file without setting up a server. 

I know of Vaadin and ZK. But they are like doing web in a desktop way of doing things. And both of them requier serverlet containers.
## [5][Google Home for Android seeing fewer crashes after adopting Kotlin](https://www.reddit.com/r/Kotlin/comments/hsn47u/google_home_for_android_seeing_fewer_crashes/)
- url: https://9to5google.com/2020/07/16/google-home-crashing/
---

## [6][Help with in coming notifications](https://www.reddit.com/r/Kotlin/comments/ht3n4p/help_with_in_coming_notifications/)
- url: https://www.reddit.com/r/Kotlin/comments/ht3n4p/help_with_in_coming_notifications/
---
So I'm using the firebase console to send a message notification to users, those notifications are being displayed in a notificationsActivity using  recyclerView and cardView. The problem I have is that those notifications can't appear automatically. You have to go to the previous activity and then enter to the notifications activity.

Here is my code

    class ventanaNotif : AppCompatActivity() {
    
    
        val notifList = ArrayList&lt;Notificaciones&gt;()
        lateinit var adapter: NotifAdapter
        lateinit var shared: SharedPreferences
        private val PREF = "Notificaciones"
        private val COUNT_KEY = "Conteo"
        var borrar: ImageButton? = null
    
        var swipeRefreshLayout: SwipeRefreshLayout? = null
    
    
        override fun onCreate(savedInstanceState: Bundle?) {
            super.onCreate(savedInstanceState)
            setContentView(R.layout.activity_ventana_notif)
    
            setRecyclerView()
    
            shared = getSharedPreferences("Notificaciones", Context.MODE_PRIVATE)
            var count = shared.getInt(COUNT_KEY, 0)
    
    
            notificaciones()
            adapter.notifyItemInserted(adapter.itemCount)
    
        }
    
    
        fun setRecyclerView() {
    
            adapter = NotifAdapter(notifList)
    
            rv_notif.adapter = adapter
            rv_notif.setHasFixedSize(true)
        }
    
    
    
        fun notificaciones() {
    
            shared = getSharedPreferences("Notificaciones", Context.MODE_PRIVATE)
            var count = shared.getInt(COUNT_KEY, 0)
    
            for ((x, valor) in (1..count).withIndex()) {
    
    
                notifList.add(
                    Notificaciones(
                        "${shared.getString("Titulo${valor}", "Todo bien")}",
                        "${shared.getString("Cuerpo${valor}", "No hay más notificaciones")}",
                        "${shared.getString("Fecha${valor}", "Sin fecha ")}"
                    )
                )
    
            }
    
    
        }
    }

Here I'm trying to use a notifyItemInserted to show last notification automatically, but it doesn't work. I have to go to the previous activity and then notifications activity to show the entire notifications.

&amp;#x200B;

I was hopping you could find the error here, any help will be very much appreciated.

Thank you
## [7][Where have I messed up?](https://www.reddit.com/r/Kotlin/comments/hswi5b/where_have_i_messed_up/)
- url: https://www.reddit.com/r/Kotlin/comments/hswi5b/where_have_i_messed_up/
---
Hey
  
I'm new to the game and just having a play around watching videos and using Android Studio (with a view to something more in the future) and I'm trying a few function out. One that has perplexed me is making a button open a new activity. Now, I did a quick search, found some code and copied it into my Main activity and sure enough both buttons link to a new activity. Sweet. I understand what the code is doing and it works. Success! Or not. For some reason I can't get the code to work in any other activity but the Main one. Nothing happens when the button is click. The code I'm using is (apologies for formatting I'm on mobile and reddit is blocked):
  
var button1 = findViewById&lt;Button&gt;(R.id.button1)
  
button1.setOnClickListener(){  
setContentView(R.id.Activity2)  
  
Is there something I'm missing here? The function works perfectly moving from Main Activity to Activity2 but won't work when moving from Activity2 to Activity3 for example.
## [8][Kotlin 1.4-M3: Generating Default Methods in Interfaces](https://www.reddit.com/r/Kotlin/comments/hs92c0/kotlin_14m3_generating_default_methods_in/)
- url: https://blog.jetbrains.com/kotlin/2020/07/kotlin-1-4-m3-generating-default-methods-in-interfaces/
---

## [9][Which Kotlin web development framework to learn?](https://www.reddit.com/r/Kotlin/comments/hs80qw/which_kotlin_web_development_framework_to_learn/)
- url: https://www.reddit.com/r/Kotlin/comments/hs80qw/which_kotlin_web_development_framework_to_learn/
---
I have been using Java/Spring for web development for years.

I recently have started learning Kotlin. I want to explore more and use/try different web development framework with Kotlin.

I know Spring  has a good support for Kotlin but I want to learn new framework that I can use in Production if chance permits.

Any suggestion between Ktor, Vertx or Micronaut and why?
## [10][TextView Coding Help](https://www.reddit.com/r/Kotlin/comments/hsph20/textview_coding_help/)
- url: https://www.reddit.com/r/Kotlin/comments/hsph20/textview_coding_help/
---
I'm stumped on how I can make a textview display a certain string based on an outcome. For instance, if a dice rolled a 4, then the textview changes to "You rolled a 4!". I feel like the answer is something I should definitely know but I'm having a brainfart. Please point me in the right direction.
