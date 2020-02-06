# Kotlin
## [1][Kotlin overtakes Scala and Clojure, to become the 2nd most popular language on the JVM](https://www.reddit.com/r/Kotlin/comments/ezqfpn/kotlin_overtakes_scala_and_clojure_to_become_the/)
- url: https://snyk.io/blog/kotlin-overtakes-scala-and-clojure-to-become-the-2nd-most-popular-language-on-the-jvm/
---

## [2][Detekt-hint is a tool for spotting programming principles violations. I need some feedback on what rules to implement and if you think this could be useful. Please have a look! https://github.com/Mkohm/detekt-hint](https://www.reddit.com/r/Kotlin/comments/ezaoda/detekthint_is_a_tool_for_spotting_programming/)
- url: https://i.redd.it/4roaqhr0c4f41.png
---

## [3][Can I get a random user out of the firebase realtime database as chat partner as a second option besides a selected user?](https://www.reddit.com/r/Kotlin/comments/ezrkhf/can_i_get_a_random_user_out_of_the_firebase/)
- url: https://www.reddit.com/r/Kotlin/comments/ezrkhf/can_i_get_a_random_user_out_of_the_firebase/
---
I have a problem with coding a Chat application. One of the planned features is getting connected with a random user from the database. Currently I have one button that loads an overview of all users, you can chose one user, and you can chat with that person (message send/receive works, also when i close the app and open it again the messages are still viewable and saved).

For that, I have a NewMessageActivity, which on click on one user leads to the ChatLogActivity and takes the user information with it (toUser and USER_KEY).

Now my plan is to basically skip the NewMessage part and get right into the ChatLogActivity, and assign a random user to that.

I was thinking to just add something like an if/else to my code (so, if toUser is assigned by NewMessageActivity load that, otherwise load random), but I can't get it to work. My users have each individual uid's, that get created randomly when signing up.
Here's my code:

    companion object {
        val TAG = "ChatLog"
    }
    val adapter = GroupAdapter&lt;GroupieViewHolder&gt;()

    var toUser: User? = null

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_chat_log)


        recyclerview_chat_log.adapter = adapter

        keyboardManagement()




        toUser = intent.getParcelableExtra&lt;User&gt;(NewMessageActivity.USER_KEY)

        supportActionBar?.title = toUser?.username



        listenForMessages()

        send_button_chat_log.setOnClickListener {
            Log.d(TAG, "Attempt to send message")
            performSendMessage()
        }

    } 

Here's how I usually get the users for the chat, in case that is important:

        val ref = FirebaseDatabase.getInstance().getReference("/users")
        ref.addListenerForSingleValueEvent(object: ValueEventListener{

            override fun onDataChange(p0: DataSnapshot) {
                val adapter = GroupAdapter&lt;GroupieViewHolder&gt;()


                p0.children.forEach{
                    Log.d("NewMessage", it.toString())
                    val user = it.getValue(User::class.java)
                    if (user != null) {
                        adapter.add(UserItem(user))
                    }
                }

                adapter.setOnItemClickListener{ item, view -&gt;

                    val userItem = item as UserItem

                    val intent = Intent(view.context, ChatLogActivity::class.java)
                    //intent.putExtra(USER_KEY, userItem.user!)
                    intent.putExtra(USER_KEY, userItem.user)

                    startActivity(intent)

                    //Zurück zum Main Menu statt zur User Auswahl
                    finish()
                }

                recyclerView_newmessage.adapter = adapter
            }

            override fun onCancelled(p0: DatabaseError) {

            }
        })
    }
    }

    class UserItem(val user: User): Item&lt;GroupieViewHolder&gt;() {
        //wird aufgerufen für die einzelnen Userobjekte
      override fun bind(viewHolder: GroupieViewHolder, position: Int) {
        viewHolder.itemView.username_textView_new_message.text = user.username

            Picasso.get().load(user.profileImageURL).into(viewHolder.itemView.picture_imageView_new_message)
    }
    //Einzelne Zeilen gestalten
    override fun getLayout(): Int {
    return R.layout.user_row_new_messages
    }
    }

I'm still a beginner, so I'm pretty clueless what to do by now. If you have any ideas and could help me, I'd really appreciate that! :)

Thanks!
## [4][How to learn Kotlin fast with 8 years experience of Android Java?](https://www.reddit.com/r/Kotlin/comments/ezlbmi/how_to_learn_kotlin_fast_with_8_years_experience/)
- url: https://www.reddit.com/r/Kotlin/comments/ezlbmi/how_to_learn_kotlin_fast_with_8_years_experience/
---
Hi all, I am new here. I have been trying to learn Kotlin from https://kotlinlang.org/docs/reference since 24 houes ago but I am still not able to grab the concepts into my heart and mind. It is like a totally different concept that I never learn in other programming language (sorry, but except Swift that I still struggle until now too).
## [5][Where do you guys put extensions in your projects?](https://www.reddit.com/r/Kotlin/comments/ezf5k8/where_do_you_guys_put_extensions_in_your_projects/)
- url: https://www.reddit.com/r/Kotlin/comments/ezf5k8/where_do_you_guys_put_extensions_in_your_projects/
---
And what do you name the file(s) that have the extension functions?

Right now I'm including an "extensions" package at whatever package scope the functions are useful for. So that means I have a "top level" extensions package (e.g., com.me.product.extensions) as well as something like com.me.product.http.extensions.

I usually just name the files after the class I'm extending. But that's not that helpful because it doesn't tell me how I'm extending the class...
## [6][How can I tell the Kotlin compiler were to put my generate ExampleProgject.jar files?](https://www.reddit.com/r/Kotlin/comments/ezish2/how_can_i_tell_the_kotlin_compiler_were_to_put_my/)
- url: https://www.reddit.com/r/Kotlin/comments/ezish2/how_can_i_tell_the_kotlin_compiler_were_to_put_my/
---
Hello, all just wondering if it is at all possible to tell Kotlin where to but Jar archive files.
## [7][Opened the Call for Papers for KotlinConf 2020, until 31st of March](https://www.reddit.com/r/Kotlin/comments/ez8gpg/opened_the_call_for_papers_for_kotlinconf_2020/)
- url: https://sessionize.com/kotlinconf/
---

## [8][How to declare unitialized map with value type of Array&lt;String&gt;](https://www.reddit.com/r/Kotlin/comments/eyuz3n/how_to_declare_unitialized_map_with_value_type_of/)
- url: https://www.reddit.com/r/Kotlin/comments/eyuz3n/how_to_declare_unitialized_map_with_value_type_of/
---
Hi, general noob here. Having trouble declaring a mutable map with key type Char and value type Array&lt;String&gt; of size 10 (values unitialized). I think it should look something like the following:

`val example = mutableMapOf&lt;Char, Array&lt;String&gt;&gt;()`

Help greatly appreciated ⌨️
## [9][Kotlin Coroutines: What happens to the caller thread when context changes?](https://www.reddit.com/r/Kotlin/comments/eyjfz7/kotlin_coroutines_what_happens_to_the_caller/)
- url: https://www.reddit.com/r/Kotlin/comments/eyjfz7/kotlin_coroutines_what_happens_to_the_caller/
---
I have a doubt that even after a long time of studying coroutines, I never managed to understand it in a simple way. The doubt is about the difference between suspending and blocking code and how suspending functions relate to threads and coroutines.

If we define a function like this

    suspend fun wait() { // some suspending code here }

From what I understand, this function definition does not create new threads, nor coroutines. It's only meant to signal to the compiler that this function might suspend, so it should only be called from a suspendable context.

But if we declare it like this instead:

    suspend fun waitOnContext() = withContext(Dispatchers.IO) {
        // some suspending code here
    }

I believe (and correct if I'm wrong) that three things happens when a function defined like this is called:

1. Since the function is suspending (not blocking), the caller thread is free to do whaterver else it needs to do.
2. The sequential execution flow stops, meaning that the code after the function will only run after it completes, like a callback.
3. A new thread is created so that the code runs in the background.

How do the environment (I don't know which name to use here) decides what to do next with the caller thread? Does it uses an event loop, OS scheduling facilities or maybe the compiler organizes the code in such a way so that other things can be done while the new thread is running on the background.

One simple example is the Android main thread: How does the environment knows that it needs to run view rendering and gesture handling code after the new thread is created? And how does it knows to come back to the rest of the coroutine once the background thread finishes?

Another doubt that I have is if `CoroutineScope` is meant to group related coroutines so that they doesn't leak or if it's just meant to represent parent-child relationships so that when `Job` is cancelled, all of the children coroutines are cancelled as well.

One thing that made me confused was [this example](https://kotlinlang.org/docs/reference/coroutines/basics.html#scope-builder) from the Kotlin documentation. In it one `coroutineScope` is defined that suspends while waiting for a child coroutines to finish. Does this happen because the `CoroutineScope` needs to wait for all children coroutines to finish so it can be finished as well?

Am I correct in assuming that a `CoroutineScope` works similarly to a function scope in the sense that while a function needs to clear its local variables once it finishes, a `CoroutineScope` needs to clear it's children coroutines by waiting for them?

Another similarity is that when an exception happens inside a function, it has to clear all its local variables (stack unwinding), while a `CoroutineScope` needs to clear all it's coroutines by canceling them. Am I correct in using this analogy?
## [10][Kotlin Multiplatform header files (iOS)](https://www.reddit.com/r/Kotlin/comments/eyp9ue/kotlin_multiplatform_header_files_ios/)
- url: https://www.reddit.com/r/Kotlin/comments/eyp9ue/kotlin_multiplatform_header_files_ios/
---
So, I am working on a kotlin multiplatform project and I need to use some objective C classes for iOS part. I would like to know if there is a way to import and use Header (.h) files in kotlin MPP?
