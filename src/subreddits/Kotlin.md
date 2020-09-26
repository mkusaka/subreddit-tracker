# Kotlin
## [1][Use Vim as a Kotlin IDE](https://www.reddit.com/r/Kotlin/comments/izy8w0/use_vim_as_a_kotlin_ide/)
- url: https://spacevim.org/use-vim-as-a-kotlin-ide/#.X26y9qBwVCU.reddit
---

## [2][Floating Windows on Android 1: Jetpack Compose &amp; Room](https://www.reddit.com/r/Kotlin/comments/j03jmd/floating_windows_on_android_1_jetpack_compose_room/)
- url: https://www.reddit.com/r/Kotlin/comments/j03jmd/floating_windows_on_android_1_jetpack_compose_room/
---
Hi, guys,

I published a series of articles about how to make floating windows (such as Facebook Heads and similar) on Android.

For 7 years, I work on Floating Apps ([https://floatingapps.net](https://floatingapps.net/)) and now I'm sharing all tips I learned along the way.

**The first lesson: Jetpack Compose &amp; Room**

[https://localazy.com/blog/floating-windows-on-android-1-jetpack-compose-and-room](https://localazy.com/blog/floating-windows-on-android-1-jetpack-compose-and-room)

Do you like it? Feel free to get in touch!
## [3][PostgreSQL in -&gt; Reactive Non-Blocking Type-Safe Kotlin bindings out](https://www.reddit.com/r/Kotlin/comments/j00jlk/postgresql_in_reactive_nonblocking_typesafe/)
- url: https://www.reddit.com/r/Kotlin/comments/j00jlk/postgresql_in_reactive_nonblocking_typesafe/
---
I wanted to learn how ANTLR works and at the same time dodge the blocking nature of JDBC while avoiding the boilerplate of binding handwritten SQL, so i modifed an existing ANTLR grammar to allow generating Kotlin code around the Jasync drivers while also validating the SQL queries (currently SQL itself is checked by the grammar, table names and column names are checked against the table model and in the case of INSERT statements, the number of columns and values are counted so that too can be validated)

The generator reads the table CREATE statement to understand the table structure and then based on that, figures out which Kotlin types to use 

&amp;#x200B;

Sample input:

    -- fun insertListingId
INSERT INTO listing (id, created, modified) VALUES (?, now(), now());


Output:

    	suspend fun insertListingId(
    		listingId: String, 
    	) {
    		con.execute(
    			// language=SQL
    			"""
    			INSERT INTO listing ( id , created , modified ) 
    			VALUES ( ? , now ( ) , now ( ) ) 
    			""",
    			listingId,
    		)
    	}


For SELECT statements i'm generating a data class as well, 

    -- fun selectAllWhereIdIs1
SELECT l.id, title FROM listing AS l WHERE l.id = ?;

Output:

    	data class SelectAllWhereIdIs1Result(
    		val listingId : String,
    		val listingTitle : String?,
    	)
    	
    	suspend fun selectAllWhereIdIs1(
    		listingId: String, 
    	): List&lt;SelectAllWhereIdIs1Result&gt; {
    		return con.execute(
    			// language=SQL
    			"""
    			SELECT l . id , title FROM listing AS l WHERE l . id = 
    			? 
    			""",
    			listingId,
    		).rows.map { row -&gt;
    			SelectAllWhereIdIs1Result(
    				listingId = row.getString(0) ?: throw Exception(
    					"listing.id is declared as non-nullable, but is returning null"
    				),
    				listingTitle = row.getString(1),
    			)
    		}
    	}

SQL Delight also reads raw SQL and generates Kotlin code, but it's around JDBC at the moment, but probably much more battle tested against a large number of queries

Anyways, just thought i'd share my fun little side-project from the last 2 weeks:  
[https://github.com/JVAAS/jvaas-postgresql](https://github.com/JVAAS/jvaas-postgresql)
## [4][Kotlin Mumbai event](https://www.reddit.com/r/Kotlin/comments/j01l97/kotlin_mumbai_event/)
- url: https://www.reddit.com/r/Kotlin/comments/j01l97/kotlin_mumbai_event/
---
Kotlin Mumbai is all geared up for our September Event going live in 90 minutes. We will have Carlos Mota and Wajahat Karim speaking on Scoped Storage in Android and Building a Game with Jetpack Compose.

You can join at [https://www.meetup.com/Kotlin-Mumbai/events/273000330/](https://www.meetup.com/Kotlin-Mumbai/events/273000330/)
## [5][Announcing a painless Kotlin/Multiplatform NoSQL embedded database](https://www.reddit.com/r/Kotlin/comments/iziu2q/announcing_a_painless_kotlinmultiplatform_nosql/)
- url: https://medium.com/kodein-koders/announcing-a-painless-kotlin-multiplatform-nosql-embedded-database-30fed677549c
---

## [6][Help - With Kotlin serialization, can ProtoBuf serialize objects with nullable values?](https://www.reddit.com/r/Kotlin/comments/j00o5g/help_with_kotlin_serialization_can_protobuf/)
- url: https://www.reddit.com/r/Kotlin/comments/j00o5g/help_with_kotlin_serialization_can_protobuf/
---
I've tried this simple code:

    fun main() {
        @Serializable data class Foo(val nn : Int?)
    
        val s = ProtoBuf.encodeToByteArray(Foo.serializer(), Foo(null))
    }

But I get:

    Exception in thread "main" kotlinx.serialization.SerializationException:
     'null' is not supported in ProtoBuf

I've also tried using `encodeDefaults` based on some research:

    fun main() {
        @Serializable data class Foo(val nn : Int?)
    
        val s = ProtoBuf { encodeDefaults = false }.encodeToByteArray(Foo.serializer(), Foo(null))
    }

Unfortunately the same result:

       Exception in thread "main" kotlinx.serialization.SerializationException:
     'null' is not supported in ProtoBuf

This is quite a surprising limitation and I'm sure it will trip up a lot of people, is anyone aware of a solution?  I found a few [issues](https://github.com/Kotlin/kotlinx.serialization/issues/71) mentioning the issue but the only suggestion I found was to try `encodeDefaults` which was unsuccessful.

Any help appreciated.
## [7][Ktor Design Process](https://www.reddit.com/r/Kotlin/comments/izivwv/ktor_design_process/)
- url: https://blog.jetbrains.com/ktor/2020/09/24/ktor-design-process/
---

## [8][Publishing and Subscribing with Halcyon - Kotlin Native XMPP Library](https://www.reddit.com/r/Kotlin/comments/izrb7x/publishing_and_subscribing_with_halcyon_kotlin/)
- url: https://tigase.net/halcyon-publishing-and-subscribing-with-halcyon/
---

## [9][Any Kotlin/Google tech speakers?](https://www.reddit.com/r/Kotlin/comments/izgwwb/any_kotlingoogle_tech_speakers/)
- url: https://www.reddit.com/r/Kotlin/comments/izgwwb/any_kotlingoogle_tech_speakers/
---
\#gdg #devfest #google  
We know there are great speakers among us, technology guru’s, leaders of the pack… Yeah! That’s you we are speaking about!

We want to offer you to take part in GDG Devfest POLAND 2020 as a speaker and share your precious knowledge with our attendees.  
Kotlin is already well settled in Android, well it's a primary language there already, so if you would like to share your story please join us, and submit your talk at [https://sessionize.com/devfest-poland-2020-with-gdg/](https://sessionize.com/devfest-poland-2020-with-gdg/)  
The conference will be held online, but it's still a great opportunity for speakers and attendees to meet other awesome people :)
## [10][Are there any free courses that combined learning Kotlin and Java](https://www.reddit.com/r/Kotlin/comments/izhq2y/are_there_any_free_courses_that_combined_learning/)
- url: https://www.reddit.com/r/Kotlin/comments/izhq2y/are_there_any_free_courses_that_combined_learning/
---
I learned Kotlin first, and now learning Java to help me with Kotlin, and feel a course that combines the two or teaches you them simultaneously first java then what kotklin took and improved upon in Java would help me pick up Java quicker
