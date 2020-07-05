# Kotlin
## [1][Kotlin 1.4-M3-eap-* of the IDEA Plugin](https://www.reddit.com/r/Kotlin/comments/hljfbh/kotlin_14m3eap_of_the_idea_plugin/)
- url: https://www.reddit.com/r/Kotlin/comments/hljfbh/kotlin_14m3eap_of_the_idea_plugin/
---
I can see Kotlin 1.4-M3 artifacts are being pushed to the `kotlin-dev` channel.  

However I can't find a build of the IDEA Plugin. Any idea where I could look? I would like to avoid building it myself.
## [2][Released KHipster 1.9.0](https://www.reddit.com/r/Kotlin/comments/hld2qu/released_khipster_190/)
- url: https://www.npmjs.com/package/generator-jhipster-kotlin
---

## [3][lndmflngs/line-to-telegram - Tool to prepare a set of Line stickers into Telegram stickers [Rewrote on the use of Executors, Need some code review]](https://www.reddit.com/r/Kotlin/comments/hlkmj5/lndmflngslinetotelegram_tool_to_prepare_a_set_of/)
- url: https://github.com/lndmflngs/line-to-telegram
---

## [4][java.lang.ExceptionInInitializerError in self-referential objects?](https://www.reddit.com/r/Kotlin/comments/hlhuvu/javalangexceptionininitializererror_in/)
- url: https://www.reddit.com/r/Kotlin/comments/hlhuvu/javalangexceptionininitializererror_in/
---
I'm having problems with the jvm compiler.

I'm trying to write a factory method for classes. The factory method has an init() block that helps to define behaviour for the new object. While this method compiles for JVM, I encounter a problem when running it:

java.lang.ExceptionInInitializerError

Caused by: java.lang.IllegalArgumentException: Parameter specified as non-null is null: method type.ProblematicKt.nullable, parameter $this$nullable

&amp;#x200B;

Apparently, the object isn't yet defined when I attempt to run the problematicInit() block. How do I fix this?

Seems to be a JVM problem. It seems to work for Javascript compilations. My understanding was that getProblematic would be hoisted, but what's inside the scope would be deferred until it's run designed to be run later - after the factory method is completed.

    interface ProblematicBuilderScope {
      fun problematicInit(getX: () -&gt; ProblematicInterface)
    }
    
    fun getProblematic() = X
    class Problematic(...): ProblematicInterface
    
    // Factory method with init() block
    val X = Problematic.factory(...) {
      problematicInit{ getProblematic() }
    } 
    
    fun factory(init: ProblematicBuilderScope.() -&gt; Unit): Problematic {
      val newObject = Problematic(...)
      val scope = ProblematicBuilderScope(newObject)
      scope.init()
      return newObject
    }
## [5][Can Kotlin.JS And Javascript be in the same project?](https://www.reddit.com/r/Kotlin/comments/hl1440/can_kotlinjs_and_javascript_be_in_the_same_project/)
- url: https://www.reddit.com/r/Kotlin/comments/hl1440/can_kotlinjs_and_javascript_be_in_the_same_project/
---
I was making some basic website after going through the React and Kotlin.JS tutorial of JetBrains and then my friend expressed his desire to contribute. But the problem is that he doesn't knows Kotlin and only Javascript and I have just started so I only know Javascript.
AFAIK, Java and Kotlin can be used in the same project but is there any way to use Kotlin.JS with simple Javascript. Or will one of us have to learn a new language?
## [6][Announcing LA4K: Logging API for Kotlin](https://www.reddit.com/r/Kotlin/comments/hkrz4x/announcing_la4k_logging_api_for_kotlin/)
- url: https://www.reddit.com/r/Kotlin/comments/hkrz4x/announcing_la4k_logging_api_for_kotlin/
---
Greetings!

I have been working on an original logging API for Kotlin Multiplatform for over a year now. I have done this because every other API I have found has really just been an extension of an existing JVM-only API, usually SLF4J. Right now, it can be bound to Android (Logcat), Log4j2, SLF4J and JUL. JS and native targets are planned for future releases.

I invite everyone to [check out its GitHub repo](https://github.com/wswartzendruber/la4k) (which includes a ReadMe). I have also [published the initial release, v0.1.0, to Bintray](https://bintray.com/wswartzendruber/la4k). I am especially interested in the community's assessment of the project, as any shortcommings need to be identified.

It is my intention for this to mature as Kotlin Multiplatform does the same. I would like for the v1.0.0 release to coincide with Multiplatform's initial, stable release.

I appreciate any feedback on this project.
## [7][Multiplatform Persistence with SQLDelight](https://www.reddit.com/r/Kotlin/comments/hkqf1e/multiplatform_persistence_with_sqldelight/)
- url: http://gh.jdoneill.com/2020/06/28/sqldelight/
---

## [8][Null Safety in Kotlin](https://www.reddit.com/r/Kotlin/comments/hl0ich/null_safety_in_kotlin/)
- url: https://medium.com/swlh/null-safety-in-kotlin-88298e64a1dc?source=friends_link&amp;sk=a3f75e55a492b26416fe36907f69e09c
---

## [9][Is there a good way to share code between projects in Kotlin Multiplatform?](https://www.reddit.com/r/Kotlin/comments/hkrkul/is_there_a_good_way_to_share_code_between/)
- url: https://www.reddit.com/r/Kotlin/comments/hkrkul/is_there_a_good_way_to_share_code_between/
---
I have a library I wrote all in Kotlin, which currently builds using Gradle and uploads an artifact to Maven. I can then use it in JVM projects, and it works fine.  

However if I use it in a multiplatform project (based on https://korge.org/), build.gradle.kts throws an error on the implementation.  I assume this is because Maven libraries are compiled into a jar and the build.gradle.kts script doesn't support that.

Is there a way to make a Kotlin specific library, in such a way that it can be imported into other projects?  I could use git submodules to include the src directory, but that's kind of messy.  Is there a better solution?
## [10][Kotlin Spring Boot throws an error getting a DateTime out of a MongoDB Database using Graph-QL Java](https://www.reddit.com/r/Kotlin/comments/hks4ea/kotlin_spring_boot_throws_an_error_getting_a/)
- url: https://www.reddit.com/r/Kotlin/comments/hks4ea/kotlin_spring_boot_throws_an_error_getting_a/
---
I have been working on a Kotlin App using Spring Boot and GraphQL with all my data stored in a MongoDB database. I have started using the GraphQL-Java-Extended-Scalars to try to deal with my DateTimes in my Database but it keeps throwing the error Expected something we can convert to 'java.time.OffsetDateTime' but was 'LocalDateTime'.  I was wonder how do you correctly setup GraphQL-Java-Extended-Scalars and what data type should I use for my DateTimes?

My GraphQL Schema is set up like this:

    scalar DateTime
    scalar Date
    
    type Query {
       users: [User]
    }
    
    type User {
      id: ID!
      username: String!
      password: String!
      email: String!
      dateOfBirth: Date
      dateCreated: DateTime!
    }

I defined the custom mapping and extended scalars like this:

    @Component
    class scalars() {
    
    
    @Bean
    fun objectMapper(): ObjectMapper? {
        return ObjectMapper().registerModule(JavaTimeModule())
                .disable(SerializationFeature.WRITE_DATES_AS_TIMESTAMPS)
    }
    
    @Bean
    fun schemaParserOptions(): SchemaParserOptions? {
        return SchemaParserOptions.newOptions()
                .objectMapperProvider(PerFieldObjectMapperProvider { fieldDefinition: FieldDefinition? -&gt; objectMapper() }).build()
    }
    
    @Bean
    fun date(): GraphQLScalarType? {
        return ExtendedScalars.Date
    }
    
    @Bean
    fun datetime(): GraphQLScalarType? {
        return ExtendedScalars.DateTime
    }
    }

Then  I created the user object with these data type set up:

    @Document(collection = "users")
    data class UserEntity(
        var username: String,
        var password: String,
        var email: String,
        var dateOfBirth: LocalDate,
        var dateCreated: LocalDateTime
     ) {
        @Id
        var id: String = ""
     }

and lastly, my DateTime in the MongoDB database is in this form:

    "dateCreated" : ISODate("2020-0703T04:30:52.237Z")

Hopefully, this is information to help solve the error it keeps throwing when I run the query for dateCreated of the user. Thank you for all your help in advance
