# Kotlin
## [1][Having worked with kotlin sealed classes for some time now, I feel that the Java 15 version of sealed classes are more flexible and allows better readability with the offence of being verbose on the declaration. Which one will you prefer?](https://www.reddit.com/r/Kotlin/comments/iwx5j4/having_worked_with_kotlin_sealed_classes_for_some/)
- url: https://www.reddit.com/gallery/iwx5j4
---

## [2][[P] ML With Android 11](https://www.reddit.com/r/Kotlin/comments/ivqd4f/p_ml_with_android_11/)
- url: https://www.reddit.com/r/Kotlin/comments/ivqd4f/p_ml_with_android_11/
---
Hey,

I open-sourced this repo by me demonstrating how you can get started with Machine Learning in Android 11 with Kotlin, all that is new. I further even explain how the code works so you can very easily get started with implementing Machine Learning in your Android apps with cool examples!

(PS: Leave me a star if this looks good to you, helps me a lot.)

[https://github.com/Rishit-dagli/ML-with-Android-11](https://github.com/Rishit-dagli/ML-with-Android-11)
## [3][Solicitar Bloque de Anuncios Admob y Android #1 ✅Curso de Programación A...](https://www.reddit.com/r/Kotlin/comments/iw2tur/solicitar_bloque_de_anuncios_admob_y_android_1/)
- url: https://www.youtube.com/watch?v=0U75khKzJ50&amp;feature=share
---

## [4][Moving to kotlin for native android Java Developers](https://www.reddit.com/r/Kotlin/comments/iv159n/moving_to_kotlin_for_native_android_java/)
- url: https://gwaza.hashnode.dev/moving-to-kotlin-for-native-android-java-developers-ckf7p3ovz00v06ds1ah8r9c7l
---

## [5][Firebase Admin Kotlin SDK](https://www.reddit.com/r/Kotlin/comments/iv58tw/firebase_admin_kotlin_sdk/)
- url: https://www.reddit.com/r/Kotlin/comments/iv58tw/firebase_admin_kotlin_sdk/
---
Is there a firebase admin SDK solely developed in Kotlin for adding admin SDK to the Server?  Now since Kotlin is interoperable with Java i can also use the Firebase Admin Java SDK but then obviously it wouldn't provide me the coroutine support for asynchronous execution of my code.
## [6][Kotlin Developer Course](https://www.reddit.com/r/Kotlin/comments/iv91z3/kotlin_developer_course/)
- url: https://www.google.com/amp/s/onlinecoursesgalore.com/complete-kotlin-developer-course-rob-percival/amp/
---

## [7][Kotlin for Backend](https://www.reddit.com/r/Kotlin/comments/iuxzsa/kotlin_for_backend/)
- url: https://www.reddit.com/r/Kotlin/comments/iuxzsa/kotlin_for_backend/
---
Hello everyone!
I was thinking of using GraphQL with Kotlin for my server side of my web application. Now there are a lot of frameworks and all of them have their own features but I just wanted to know which of them will be most compatible with GraphQL. Have anyone of you worked with GraphQL and built the server side without any trouble?
For now I'm thinking about Javalin or Micronaut. I'm considering Vert.x (but listened it's steeper to learn) and Spring (didn't find anything related to Graphql on their website) too.
## [8][How to set up Kotlin for NodeJS with TS dependencies?](https://www.reddit.com/r/Kotlin/comments/iujq6a/how_to_set_up_kotlin_for_nodejs_with_ts/)
- url: https://www.reddit.com/r/Kotlin/comments/iujq6a/how_to_set_up_kotlin_for_nodejs_with_ts/
---
Hey, I've been using Kotlin in the JVM for a few years and I decided to try it on nodejs (I'm also working with node for a few years, so I'm not newbie in the platform).

&amp;#x200B;

I went to IDEA &gt; New Project &gt; Gradle &gt; Kotlin (NodeJS). Added a few dependencies and enable the experimental Dukat:

build.gradle.kts`dependencies {implementation(kotlin("stdlib-js"))implementation(npm("firebase", "7.20.0"))implementation(npm("left-pad", "1.3.0"))}`

&amp;#x200B;

gradle.properties

`kotlin.code.style=officialkotlin.js.experimental.generateKotlinExternals=true`

Both firebase and left-pad (obviously I'm using this for test) does have typescript mappings, so the kotlin mappings should be generated, right?

I refreshed Gradle and went to the src/main trying to import the dependencies, but nothing happened and I don't know why. Looking at build/js/node\_modules the dependencies are there.

&amp;#x200B;

I tried to run the Dukat gradle tasks and nothing happened too.

&amp;#x200B;

This section about external dependencies is very bad documented, I need some help to set it up.

btw I'm on Kotlin 1.4, my full build.gradle looks like this:`plugins {id("org.jetbrains.kotlin.js") version "1.4.10"}group = "dev.nathanpb"version = "1.0.0"repositories {mavenCentral()}dependencies {implementation(kotlin("stdlib-js"))implementation(npm("firebase", "7.20.0"))implementation(npm("left-pad", "1.3.0"))}kotlin {js {nodejs {}binaries.executable()}}`

&amp;#x200B;

Thank you all, hope I will be using this in production soon

&amp;#x200B;

**SOLUTION:** I'm stupid and can't read documentations: [https://kotlinlang.org/docs/reference/js-external-declarations-with-dukat.html#generating-external-declarations-at-build-time](https://kotlinlang.org/docs/reference/js-external-declarations-with-dukat.html#generating-external-declarations-at-build-time)
## [9][Using createPortal in functionalComponent of Kotlin-react](https://www.reddit.com/r/Kotlin/comments/iui3th/using_createportal_in_functionalcomponent_of/)
- url: https://www.reddit.com/r/Kotlin/comments/iui3th/using_createportal_in_functionalcomponent_of/
---
Does anyone have some idea how to use createPortal in kotlin-react for example to render a modal? i have tried several ways and nothing is working. I´m new at Kotlin. Here is one way i tryed:

  
`interface Modalprops:RProps{`  
 `var visible:Boolean`  
`}`  


`val modal = functionalComponent&lt;Modalprops&gt; {props-&gt;`  
 `val (container) = useState(document.create.div{+"asdasdsa"})`  
 `useEffectWithCleanup(emptyList()) {`  
 `document.body?.appendChild(container)`  
 `return@useEffectWithCleanup { document.body?.removeChild(container)}`  
`}`  
 `createPortal(container){`  
 `if(props.visible) {`  
 `div { +"asdasd" }`  
 `}`  
 `}`  
`}`  
`fun RBuilder.OModal(isvisible:Boolean, handler:Modalprops.()-&gt;Unit) = child(modal){`  
 `attrs{`  
 `visible=isvisible`  
 `}`  
`}`  


its just how its working in good old TS and i tried to transfer it to kotlin-react but createPortal seems to do nothing in this case.

&amp;#x200B;

It would be nice if some could give me an example of how implement this example correct in react-kotlin.
## [10][Thread within coroutine](https://www.reddit.com/r/Kotlin/comments/iuqh7t/thread_within_coroutine/)
- url: https://www.reddit.com/r/Kotlin/comments/iuqh7t/thread_within_coroutine/
---
Is this a bad idea or is it fine?

Because when I'm done writing to a file I flush it but the flush function uses synchronized block so I start a thread within a coroutine to run the flush function but this has weird side effects (very strange bugs I can't really explain right now).
