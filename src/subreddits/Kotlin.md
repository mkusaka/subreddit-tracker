# Kotlin
## [1][Kotlin Actors – No Drama Concurrency](https://www.reddit.com/r/Kotlin/comments/htlgjt/kotlin_actors_no_drama_concurrency/)
- url: https://handstandsam.com/2020/07/18/kotlin-actors-no-drama-concurrency/
---

## [2][How to update mutable state in a gRPC client app?](https://www.reddit.com/r/Kotlin/comments/htu9af/how_to_update_mutable_state_in_a_grpc_client_app/)
- url: https://www.reddit.com/r/Kotlin/comments/htu9af/how_to_update_mutable_state_in_a_grpc_client_app/
---
This is a design question about the optimal/recommended way of maintaining shared mutable state using Kotlin concurrency primitives. For that, I need to establish the context first, please bear with me.  
I’ve a Spring Boot web application that has a gRPC client. Following are the requirements/constraints on the system:

1. Messages from the gRPC server could be informational or *commands*.
2. The client has internal state, that may change depending on the commands received from the server. Informational messages don’t affect client state.
3. Based on the commands received from the server, the client may generate more messages for the server, or update its internal state, or do both.
4. Every client message generated due to a server command *must happen after* any state changes.

My current design is as follows;

1. The client is a Spring bean. It has a reference to its internal state, which is a POJO (or POKO, if you want)
2. Upon receiving a server message, it constructs a command instance and passes the message content and the client state to the command. This happens in a gRPC streaming receiver, which runs on a gRPC thread pool (default).
3. The command instance may update the client state, and returns one or messages to the client.
4. The client, running within an Unconfined dispatcher, puts the messages in a buffered Channel, and waits for the next server message.
5. A Channel consumer, running on Kotlin default dispatcher pool, picks up the messages and sends them to the server.

Notice that requirement #4 is satisfied by design steps #3 and #5.

Questions:

1. I keep hearing that Flow is recommended over Channel. In this case though, the Channel  
 needs to stay open throughout the lifecycle of the client, and is very easy to use as an async queue. What benefits, if any, are there for using a Flow in this case?
2. I don’t feel particularly good about passing the client state to the commands. Is there a different way of achieving requirement #4 without doing this? (actors, may be)

*Disclaimer: Also posted on kotlin forum, but haven't received any replies. People tend to avoid long posts.*
## [3][Firebase competitors.](https://www.reddit.com/r/Kotlin/comments/htfidl/firebase_competitors/)
- url: https://www.reddit.com/r/Kotlin/comments/htfidl/firebase_competitors/
---
I'm interested in building my first Android app using Kotlin.  Base on my research thus far, Firebase seem to be a good solution to help minimize some of the development start up work.

--

I was wondering if there are alternatives / competitors to Firebase that I should do some research on before committing to Firebase.

--

Thanks.
## [4][📣 Calling all Kotlin/JS developers! 📣](https://www.reddit.com/r/Kotlin/comments/hth3e9/calling_all_kotlinjs_developers/)
- url: https://twitter.com/kotlin/status/1284114694702346240?s=19
---

## [5][Android Kotlin RecyclerView Tutorial - In-Depth Guide incl. Different View Types](https://www.reddit.com/r/Kotlin/comments/hth7c2/android_kotlin_recyclerview_tutorial_indepth/)
- url: https://www.reddit.com/r/Kotlin/comments/hth7c2/android_kotlin_recyclerview_tutorial_indepth/
---
I watched this video. It is amazing it is really very depth knowledge as I am familiar with android development but if any beginner wants to learn android it is really amazing. I loved it.

Video Link: [https://youtu.be/oDfl-xLXiac](https://youtu.be/oDfl-xLXiac)
## [6][Java Slack SDK 1.1.0 includes a Kotlin DSL for building messages](https://www.reddit.com/r/Kotlin/comments/hswsr4/java_slack_sdk_110_includes_a_kotlin_dsl_for/)
- url: https://slack.dev/java-slack-sdk/guides/composing-messages#block-kit-kotlin-dsl
---

## [7][Learn Android and Kotlin with no programming experience (google)](https://www.reddit.com/r/Kotlin/comments/ht35ju/learn_android_and_kotlin_with_no_programming/)
- url: https://android-developers.googleblog.com/2020/07/learn-android-and-kotlin-with-no-experience.html
---

## [8][Framework similar to Blazor/Meteor/LiveView?](https://www.reddit.com/r/Kotlin/comments/ht0o9w/framework_similar_to_blazormeteorliveview/)
- url: https://www.reddit.com/r/Kotlin/comments/ht0o9w/framework_similar_to_blazormeteorliveview/
---
Do we have some framework like C#'s Blazor, or JS's Meteor or Elixir's Phoenix LiveView in Kotlin or JVM worlds? In a sence that framework will do automatic client side and server side communications using WebSockets for me. While still allowing me to write templates in HTML and CSS.

I'd prefer some framework with embedded server (like Javalin for example) so I could deploy my web app as a single jar-file without setting up a server. 

I know of Vaadin and ZK. But they are like doing web in a desktop way of doing things. And both of them requier serverlet containers.
## [9][Google Home for Android seeing fewer crashes after adopting Kotlin](https://www.reddit.com/r/Kotlin/comments/hsn47u/google_home_for_android_seeing_fewer_crashes/)
- url: https://9to5google.com/2020/07/16/google-home-crashing/
---

## [10][Help with in coming notifications](https://www.reddit.com/r/Kotlin/comments/ht3n4p/help_with_in_coming_notifications/)
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
