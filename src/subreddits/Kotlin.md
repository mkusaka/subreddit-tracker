# Kotlin
## [1][Role-based authorization in Ktor](https://www.reddit.com/r/Kotlin/comments/j0mrli/rolebased_authorization_in_ktor/)
- url: https://www.reddit.com/r/Kotlin/comments/j0mrli/rolebased_authorization_in_ktor/
---
I recently implemented role-based authorization in our Ktor project, and wrote a blog about its implementation and what we learned about Ktor's internals in the process. Always happy to receive your comments!  


[https://www.ximedes.com/2020-09-17/role-based-authorization-in-ktor/](https://www.ximedes.com/2020-09-17/role-based-authorization-in-ktor/)
## [2][[Academic] Open Source Development Survey (Software Developers using GitHub)](https://www.reddit.com/r/Kotlin/comments/j0q7vd/academic_open_source_development_survey_software/)
- url: /r/SampleSize/comments/iybrqr/academic_open_source_development_survey_software/
---

## [3][Kotlin console with corutines and Mssql. What to use for data layer.](https://www.reddit.com/r/Kotlin/comments/j0ozcw/kotlin_console_with_corutines_and_mssql_what_to/)
- url: https://www.reddit.com/r/Kotlin/comments/j0ozcw/kotlin_console_with_corutines_and_mssql_what_to/
---
Hi I'am looking for something like a dapper in C# for a console application in kotlin with corutines and Mssql. So far, it looks like all the tutorials are for SpringBoot R2DBC. I found Kotlin Exposed, but it is not async rn. Any tips?
## [4][Comparability of Tinder’s finite state machine with Android MVVM](https://www.reddit.com/r/Kotlin/comments/j08wiq/comparability_of_tinders_finite_state_machine/)
- url: https://www.reddit.com/r/Kotlin/comments/j08wiq/comparability_of_tinders_finite_state_machine/
---
Any good examples for implementing in an MVVM system ? Is it a good choice ?
## [5][Use Vim as a Kotlin IDE](https://www.reddit.com/r/Kotlin/comments/izy8w0/use_vim_as_a_kotlin_ide/)
- url: https://spacevim.org/use-vim-as-a-kotlin-ide/#.X26y9qBwVCU.reddit
---

## [6][What's the type of projects you are working?](https://www.reddit.com/r/Kotlin/comments/j078hb/whats_the_type_of_projects_you_are_working/)
- url: https://www.reddit.com/r/Kotlin/comments/j078hb/whats_the_type_of_projects_you_are_working/
---
I want to begin work with Kotlin, but for now with a focus on backend applications. But always I search for jobs I see only positions for Android developers.

Currently, I live in Brazil working as a NodeJS development (3+ years of experience) and have a nice Java background (4+ years of experience) working on projects with Spring and Hibernate. I am curious too about the stack used for backend apps. Ktor, Kafka, JUnit, what more?

I will suffer to get jobs with this focus?
I will play with the Android ecosystem late but for now, something like microservices is my target cause I feel more comfortable.

Sorry for my very bad English! 😐

[View Poll](https://www.reddit.com/poll/j078hb)
## [7][Floating Windows on Android 1: Jetpack Compose &amp; Room](https://www.reddit.com/r/Kotlin/comments/j03jmd/floating_windows_on_android_1_jetpack_compose_room/)
- url: https://www.reddit.com/r/Kotlin/comments/j03jmd/floating_windows_on_android_1_jetpack_compose_room/
---
Hi, guys,

I published a series of articles about how to make floating windows (such as Facebook Heads and similar) on Android.

For 7 years, I work on Floating Apps ([https://floatingapps.net](https://floatingapps.net/)) and now I'm sharing all tips I learned along the way.

**The first lesson: Jetpack Compose &amp; Room**

[https://localazy.com/blog/floating-windows-on-android-1-jetpack-compose-and-room](https://localazy.com/blog/floating-windows-on-android-1-jetpack-compose-and-room)

Do you like it? Feel free to get in touch!
## [8][PostgreSQL in -&gt; Reactive Non-Blocking Type-Safe Kotlin bindings out](https://www.reddit.com/r/Kotlin/comments/j00jlk/postgresql_in_reactive_nonblocking_typesafe/)
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
## [9][Kotlin Mumbai event](https://www.reddit.com/r/Kotlin/comments/j01l97/kotlin_mumbai_event/)
- url: https://www.reddit.com/r/Kotlin/comments/j01l97/kotlin_mumbai_event/
---
Kotlin Mumbai is all geared up for our September Event going live in 90 minutes. We will have Carlos Mota and Wajahat Karim speaking on Scoped Storage in Android and Building a Game with Jetpack Compose.

You can join at [https://www.meetup.com/Kotlin-Mumbai/events/273000330/](https://www.meetup.com/Kotlin-Mumbai/events/273000330/)
## [10][Kotlin board game?](https://www.reddit.com/r/Kotlin/comments/j088at/kotlin_board_game/)
- url: https://www.reddit.com/r/Kotlin/comments/j088at/kotlin_board_game/
---
I was thinking about converting a board game into a mobile app in Kotlin as a fun project and a learning experience.

But currently I struggle even with search for tutorials, libraries or general Kotlin game ideas and tips.

Is it possible to create board game (somewhat) elegantly in Kotlin?  
If so, could you please refer me some materials?
