# Kotlin
## [1][Indentation when operating on the result of a block - what would you do here?](https://www.reddit.com/r/Kotlin/comments/f05m7b/indentation_when_operating_on_the_result_of_a/)
- url: https://www.reddit.com/r/Kotlin/comments/f05m7b/indentation_when_operating_on_the_result_of_a/
---
Given the example block of code, which style do you prefer:

&amp;#x200B;

https://preview.redd.it/thhtiauxwff41.png?width=356&amp;format=png&amp;auto=webp&amp;s=05765122c8bc14a8051f2d027588f2e3173f2056

I encounter this a lot, usually when working with RxJava and I'm never sure what the best approach is, none of them feel very neat to me.
## [2][Kotlin overtakes Scala and Clojure, to become the 2nd most popular language on the JVM](https://www.reddit.com/r/Kotlin/comments/ezqfpn/kotlin_overtakes_scala_and_clojure_to_become_the/)
- url: https://snyk.io/blog/kotlin-overtakes-scala-and-clojure-to-become-the-2nd-most-popular-language-on-the-jvm/
---

## [3][It’s Nothing](https://www.reddit.com/r/Kotlin/comments/ezuhgz/its_nothing/)
- url: https://www.zacsweers.dev/its-nothing/
---

## [4][Intellisense or auto-completion not working (Kotlin)](https://www.reddit.com/r/Kotlin/comments/ezwpyo/intellisense_or_autocompletion_not_working_kotlin/)
- url: /r/vscode/comments/ezv4f7/intellisense_or_autocompletion_not_working_kotlin/
---

## [5][vscode problem](https://www.reddit.com/r/Kotlin/comments/ezwq1w/vscode_problem/)
- url: /r/vscode/comments/ezv4f7/intellisense_or_autocompletion_not_working_kotlin/
---

## [6][Detekt-hint is a tool for spotting programming principles violations. I need some feedback on what rules to implement and if you think this could be useful. Please have a look! https://github.com/Mkohm/detekt-hint](https://www.reddit.com/r/Kotlin/comments/ezaoda/detekthint_is_a_tool_for_spotting_programming/)
- url: https://i.redd.it/4roaqhr0c4f41.png
---

## [7][Can I get a random user out of the firebase realtime database as chat partner as a second option besides a selected user?](https://www.reddit.com/r/Kotlin/comments/ezrkhf/can_i_get_a_random_user_out_of_the_firebase/)
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
## [8][How to learn Kotlin fast with 8 years experience of Android Java?](https://www.reddit.com/r/Kotlin/comments/ezlbmi/how_to_learn_kotlin_fast_with_8_years_experience/)
- url: https://www.reddit.com/r/Kotlin/comments/ezlbmi/how_to_learn_kotlin_fast_with_8_years_experience/
---
Hi all, I am new here. I have been trying to learn Kotlin from https://kotlinlang.org/docs/reference since 24 houes ago but I am still not able to grab the concepts into my heart and mind. It is like a totally different concept that I never learn in other programming language (sorry, but except Swift that I still struggle until now too).
## [9][Where do you guys put extensions in your projects?](https://www.reddit.com/r/Kotlin/comments/ezf5k8/where_do_you_guys_put_extensions_in_your_projects/)
- url: https://www.reddit.com/r/Kotlin/comments/ezf5k8/where_do_you_guys_put_extensions_in_your_projects/
---
And what do you name the file(s) that have the extension functions?

Right now I'm including an "extensions" package at whatever package scope the functions are useful for. So that means I have a "top level" extensions package (e.g., com.me.product.extensions) as well as something like com.me.product.http.extensions.

I usually just name the files after the class I'm extending. But that's not that helpful because it doesn't tell me how I'm extending the class...
## [10][How can I tell the Kotlin compiler were to put my generate ExampleProgject.jar files?](https://www.reddit.com/r/Kotlin/comments/ezish2/how_can_i_tell_the_kotlin_compiler_were_to_put_my/)
- url: https://www.reddit.com/r/Kotlin/comments/ezish2/how_can_i_tell_the_kotlin_compiler_were_to_put_my/
---
Hello, all just wondering if it is at all possible to tell Kotlin where to but Jar archive files.
