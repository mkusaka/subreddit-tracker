# Kotlin
## [1][Kotlin Code Organization](https://www.reddit.com/r/Kotlin/comments/eumckv/kotlin_code_organization/)
- url: https://arturdryomov.dev/posts/kotlin-code-organization/
---

## [2][Looking for good resources for asking questions (learning Kotlin)](https://www.reddit.com/r/Kotlin/comments/euk51n/looking_for_good_resources_for_asking_questions/)
- url: https://www.reddit.com/r/Kotlin/comments/euk51n/looking_for_good_resources_for_asking_questions/
---
I'm coming from Python where there are great resources to ask questions(eg /r/learnpython), where do I go for Kotlin? I have sent a request for the jetbrains slack channel, just waiting for reply.

&amp;#x200B;

The specific question I was looking to ask about this time is re statistical combinations and permutations. Python has the Itertools library, is there something similar for java/kotlin?

&amp;#x200B;

Thank you
## [3][Kotlin/JS + Ethereum](https://www.reddit.com/r/Kotlin/comments/eulcgn/kotlinjs_ethereum/)
- url: https://www.reddit.com/r/Kotlin/comments/eulcgn/kotlinjs_ethereum/
---
Thought maybe some of you guys doing Kotlin/JS might find it interesting to see how we used it to build a URL shortener on top of Ethereum.

[Source code here](https://github.com/mohamedelshami/AKAP-url-shortener), and if your browser supports ENS head on over to [redir.eth](https://redir.eth).
## [4][How to make drawers in Jetpack Compose](https://www.reddit.com/r/Kotlin/comments/eun0lg/how_to_make_drawers_in_jetpack_compose/)
- url: https://android.jlelse.eu/exploring-drawers-in-jetpack-compose-3131e6f1b07b
---

## [5][GSON fromJson doesn't compute - please help](https://www.reddit.com/r/Kotlin/comments/euno0i/gson_fromjson_doesnt_compute_please_help/)
- url: https://www.reddit.com/r/Kotlin/comments/euno0i/gson_fromjson_doesnt_compute_please_help/
---
I'm successfully receiving a JSON Object as request and passing it on to my parser. Code runs through until I call fromJson and then gets stuck. What am I doing wrong?

Here's the corresponding class:

    class User(@SerializedName("mac") private val phoneMac: String) : Comparable&lt;User&gt; {
    	
    	@SerializedName("values")
    	private val measurements: MutableSet&lt;Measurement&gt; = mutableSetOf()
    	
    	fun getPhoneMac(): String = phoneMac
    	fun getMeasurements(): Set&lt;Measurement&gt; = measurements
    	
    	//etc.
    }

Which refers to this class:

    class Measurement (@SerializedName("mac") val deviceMac: String, val timestamp: String, val value: Double, val valueType: ValueType) : Comparable&lt;Measurement&gt;{
    	
    	fun getDeviceMac(): String = deviceMac
    	fun getTimestamp(): String = timestamp
    	fun getValue(): Double = value
    	fun getValueType(): ValueType = valueType
    
            //etc.
    }

And here is how I try to parse it:

    fun fromJson(json: String): User {
    	val builder = GsonBuilder()
    	builder.setPrettyPrinting()
    		
    	return builder.create().fromJson(json, User::class.java)
    }

&amp;#x200B;

Had the fromJson-function spreaded out to make sure where it gets stuck: create() still works, fromJson() doesn't

Also, I know that the JSON file is correct and doesn't contain missing values or nulls.

For verification:

    {
    "mac": "00-80-41-ae-fd-b1",
    "values": 
    [
    	{
    	"mac": "ab-cd-ef-98-76-13",
    	"timestamp": "2012-04-23T18:25:43",
    	"value": 68,
    	"valuetype": "HR"
    	},
    	{
    	"mac": "ab-cd-ef-98-76-13",
    	"timestamp": "2012-04-23T18:35:43",
    	"value": 65,
    	"valuetype": "HR"
    	}
    ]
    }
## [6][Any idea why abstract class having functions with default argument doesn't generate their corresponding $default(...) synthetic method in Kotlin Multi-Platform?](https://www.reddit.com/r/Kotlin/comments/eu82n8/any_idea_why_abstract_class_having_functions_with/)
- url: https://stackoverflow.com/questions/59919778/expect-abstract-class-with-default-argument-doesnt-generate-default-static-met
---

## [7][Reimport Gradle Project" functionality is failing](https://www.reddit.com/r/Kotlin/comments/eubmz7/reimport_gradle_project_functionality_is_failing/)
- url: https://www.reddit.com/r/Kotlin/comments/eubmz7/reimport_gradle_project_functionality_is_failing/
---
I hope this the right sub.  I couldn't find a r/askKotlin or anything.  I asked this on the JetBrains forum but I haven't heard back.

I'm new to Kotlin. Trying to get tornadoFX to work in IDEA.  I'm  running Manjaro linux and community IDEA 2019.3. On my OS, I am using Gradle 6.0.1, but I don't know if IDEA is using  that or it's own internal copy of Gradle.

I have been able to get  simple "hello world" apps to compile and run (both with Gradle and without).  I'm now trying to  make my first tornadoFX app.

I  installed the tornadoFX plugin and then created a blank gradle  tornadoFX app. IDEA created the project with a MyApp and MainView (with  "Hello TornadoFX" title) for me. I wanted to build it without making any  changes to see what happens, and then add code after I got it to work out of the box.

The first thing I notice is that "build: sync" fails. Clicking  around I gather that is the "Reimport Gradle Project" functionality that  is failing.  I can click on the error (the text, red exclamation point,  etc.) but I do not see any additional information about the failure.   When I try to build my project (clicking the hammer icon) nothing  happens and it says "All files are up-to-date". I can add comments or  inject errors into the source and it always behaves  the same.  I assume this is a result of the sync failing and that it's not even attempting to build.

Any ideas what is going on?  (Or is there a better sub to ask this question?)
## [8][Web development - Kotlin full stack](https://www.reddit.com/r/Kotlin/comments/etztnn/web_development_kotlin_full_stack/)
- url: https://www.reddit.com/r/Kotlin/comments/etztnn/web_development_kotlin_full_stack/
---
Hi Kotliners,

I'm wondering if anyone has experience with Kotlin full-stack for web development projects. I've just started looking into Thinkter (basically Ktor + React in kotlin, [https://github.com/Kotlin/kotlin-fullstack-sample](https://github.com/Kotlin/kotlin-fullstack-sample)), but I'm not really sure if it's a true alternative to JS-based technologies.

I previously used JSNode, Express, React + JSX + Redux, Axios and Babel. It worked nice, but coming from the mobile software development, I really missed Kotlin and dynamic typing really killed me, I just never got used to it. Although I haven't been using it for that long, so I might one day, but if I could, I'd switch to Kotlin without ever looking back.

So, what are your opinions / do you have any experience with Kotlin in full-stack web development?
## [9][Help me pick a framework](https://www.reddit.com/r/Kotlin/comments/eu7nk2/help_me_pick_a_framework/)
- url: https://www.reddit.com/r/Kotlin/comments/eu7nk2/help_me_pick_a_framework/
---
For a school project, i have to make a Content Management System on android using Kotlin and microservices. Both microservices and Kotlin is new to me, so which framework would be best to use for this? Ktor, Spring Boot, Micronaut etc? Whats the difference between them?
## [10][Is anyone using the Java module system with Kotlin? What are you using it for?](https://www.reddit.com/r/Kotlin/comments/etsxgt/is_anyone_using_the_java_module_system_with/)
- url: https://www.reddit.com/r/Kotlin/comments/etsxgt/is_anyone_using_the_java_module_system_with/
---

