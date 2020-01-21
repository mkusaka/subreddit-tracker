# Kotlin
## [1][[Video] Visualising concurrency with Kotlin by Alexey Soshin](https://www.reddit.com/r/Kotlin/comments/erufjm/video_visualising_concurrency_with_kotlin_by/)
- url: https://codesync.global/media/visualising-concurrency-with-kotlin-cmldn19/?utm_source=Reddit&amp;utm_medium=Code%20Sync&amp;utm_campaign=Code%20MESH%20LDN%2019
---

## [2][Is it bad practice to directly set the value of an object's field?](https://www.reddit.com/r/Kotlin/comments/ermx0p/is_it_bad_practice_to_directly_set_the_value_of/)
- url: https://www.reddit.com/r/Kotlin/comments/ermx0p/is_it_bad_practice_to_directly_set_the_value_of/
---
If you have something like this:

    val a = A() // where A is a constructor
    a.b = "c"

Is it a bad idea? Should you use something like this?

    val a = A()
    a.setB("c")
## [3][Kotlin Algorithm Challenge No. 3](https://www.reddit.com/r/Kotlin/comments/erk4ae/kotlin_algorithm_challenge_no_3/)
- url: https://www.alexhwoods.com/blog/reverse-integer
---

## [4][Creating a Kotlin MP library for Bintray Publishing](https://www.reddit.com/r/Kotlin/comments/erhgdc/creating_a_kotlin_mp_library_for_bintray/)
- url: https://www.reddit.com/r/Kotlin/comments/erhgdc/creating_a_kotlin_mp_library_for_bintray/
---
I want to publish a kotlin multiplatform library on bintray/maven which itself  depends on another library. The problem I have is with some artifacts which get created without me being able to configure them. Let me explain this with this example

**Build.gradle**

    buildscript {
        repositories {
            mavenLocal()
            maven { url = uri("https://dl.bintray.com/korlibs/korlibs") }
            maven { url = uri("https://plugins.gradle.org/m2/") }
            mavenCentral()
        }
        dependencies {
            classpath("com.soywiz.korlibs.korge.plugins:korge-gradle-plugin:$korgePluginVersion")
            classpath("org.jetbrains.dokka:dokka-gradle-plugin:0.10.0")
        }
    }
    
    plugins {
        id 'org.jetbrains.kotlin.multiplatform' version '1.3.61'
        id "com.jfrog.bintray" version "1.8.4"
    
    }
    
    group 'me.CodingProjectg'
    version engineCodingProjectVersion
    
    apply plugin : "korge"
    apply plugin : "org.jetbrains.dokka"
    
    korge {
        id = "me.CodingProjectg.engineCodingProject"
        name = "EngineCodingProject"
        description = ""
        jvmMainClassName = "MainKt"
        entryPoint = "engineCodingProject.main"
        supportBox2d()
    }
    
    dokka {
        outputFormat = "html"
        outputDirectory = "dokka"
    }
    
    allprojects {
        repositories {
            jcenter()
        }
    
        apply plugin: 'maven-publish'
        apply plugin: 'java'
    
    
    }
    
    task sourcesJar(type: Jar) {
        archiveClassifier.set("sources")
        from sourceSets.main.allSource
    }
    
    task javadocJar(type: Jar) {
        archiveClassifier.set("javadoc")
    }
    
    
    publishing {
        publications {
            MyPublication(MavenPublication) {
                groupId = groupId
                artifactId = artifactId
                version = engineCodingProjectVersion
                from components.java
                artifact(sourcesJar)
                artifact(javadocJar)
    
                pom {
                    name = "EngineCodingProject"
                    description = "EngineCodingProjectDes"
                    url = "https://CodingProjectg.me"
                    licenses {
                        license {
                            name = "The Apache License, Version 2.0"
                            url = "http://www.apache.org/licenses/LICENSE-2.0.txt"
                        }
                    }
                    scm {
                        url = engineCodingProjectVcs
                    }
                    developers {
                        developer {
                            id = "CodingProjectgn"
                            name = "jondoe CodingProjectg"
                            email = "jondoe@CodingProjectg.me"
                        }
                    }
                }
            }
        }
    }
    
    bintray {
        user = bintrayUser
        key = bintrayApiKey
        publications = ["myPublication"]
        pkg {
            repo = 'engineCodingProject'
            name = 'engineCodingProject'
            userOrg = 'CodingProjectgn'
            licenses = ['Apache-2.0']
            vcsUrl = 'https://github.com/CodingProjectgn/engineCodingProjectLib.git'
            websiteUrl = 'https://CodingProjectg.me'
            issueTrackerUrl = 'https://github.com/CodingProjectgn/engineCodingProjectLib/issues'
            version {
                name = engineCodingProjectVersion
                desc = 'EngineCodingProject Library'
                vcsTag = engineCodingProjectVersion
                released = new Date()
            }
        }
    }
    
    bintrayUpload.doFirst {
        publications = publishing.publications.collect {
            it.name
        }.findAll {
            it != "kotlinMultiplatform"
        }
    }
    
    bintrayUpload.dependsOn publishToMavenLocal
    
    kotlin {
        jvm()
        js {
            browser {
            }
            nodejs {
            }
        }
        // For ARM, should be changed to iosArm32 or iosArm64
        // For Linux, should be changed to e.g. linuxX64
        // For MacOS, should be changed to e.g. macosX64
        // For Windows, should be changed to e.g. mingwX64
        macosX64("macos")
        sourceSets {
            commonMain {
                dependencies {
                    implementation kotlin('stdlib-common')
                }
            }
            commonTest {
                dependencies {
                    implementation kotlin('test-common')
                    implementation kotlin('test-annotations-common')
                }
            }
            jvmMain {
                dependencies {
                    implementation kotlin('stdlib-jdk8')
                }
            }
            jvmTest {
                dependencies {
                    implementation kotlin('test')
                    implementation kotlin('test-junit')
                }
            }
            jsMain {
                dependencies {
                    implementation kotlin('stdlib-js')
                }
            }
            jsTest {
                dependencies {
                    implementation kotlin('test-js')
                }
            }
            macosMain {
            }
            macosTest {
            }
        }
    }
    
    afterEvaluate {
        project.publishing.publications.all {
            // rename artifacts
            groupId = groupId
            if (it.name.contains('metadata')) {
                artifactId = artifactId
            } else {
                artifactId = "$artifactId-$name"
            }
        }
    }

Gradle Output:

    &gt; Task :bintrayUpload
    Uploading to https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-MyPublication/0.29e/engineCodingProject-MyPublication-0.29e-sources.jar...
    Uploaded to 'https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-MyPublication/0.29e/engineCodingProject-MyPublication-0.29e-sources.jar'.
    Uploading to https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-MyPublication/0.29e/engineCodingProject-MyPublication-0.29e-javadoc.jar...
    Uploaded to 'https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-MyPublication/0.29e/engineCodingProject-MyPublication-0.29e-javadoc.jar'.
    Uploading to https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-MyPublication/0.29e/engineCodingProject-MyPublication-0.29e.jar...
    Uploaded to 'https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-MyPublication/0.29e/engineCodingProject-MyPublication-0.29e.jar'.
    Uploading to https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-MyPublication/0.29e/engineCodingProject-MyPublication-0.29e.pom...
    Uploaded to 'https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-MyPublication/0.29e/engineCodingProject-MyPublication-0.29e.pom'.
    Uploading to https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-iosarm32-iosArm32/4.10.1/engineCodingProject-iosarm32-iosArm32-4.10.1-sources.jar...
    Uploaded to 'https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-iosarm32-iosArm32/4.10.1/engineCodingProject-iosarm32-iosArm32-4.10.1-sources.jar'.
    Uploading to https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-iosarm32-iosArm32/4.10.1/engineCodingProject-iosarm32-iosArm32-4.10.1.klib...
    Uploaded to 'https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-iosarm32-iosArm32/4.10.1/engineCodingProject-iosarm32-iosArm32-4.10.1.klib'.
    Uploading to https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-iosarm32-iosArm32/4.10.1/engineCodingProject-iosarm32-iosArm32-4.10.1.pom...
    Uploaded to 'https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-iosarm32-iosArm32/4.10.1/engineCodingProject-iosarm32-iosArm32-4.10.1.pom'.
    Uploading to https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-iosarm64-iosArm64/4.10.1/engineCodingProject-iosarm64-iosArm64-4.10.1-sources.jar...
    Uploaded to 'https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-iosarm64-iosArm64/4.10.1/engineCodingProject-iosarm64-iosArm64-4.10.1-sources.jar'.
    Uploading to https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-iosarm64-iosArm64/4.10.1/engineCodingProject-iosarm64-iosArm64-4.10.1.klib...
    Uploaded to 'https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-iosarm64-iosArm64/4.10.1/engineCodingProject-iosarm64-iosArm64-4.10.1.klib'.
    Uploading to https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-iosarm64-iosArm64/4.10.1/engineCodingProject-iosarm64-iosArm64-4.10.1.pom...
    Uploaded to 'https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-iosarm64-iosArm64/4.10.1/engineCodingProject-iosarm64-iosArm64-4.10.1.pom'.
    Uploading to https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-iosx64-iosX64/4.10.1/engineCodingProject-iosx64-iosX64-4.10.1-sources.jar...
    Uploaded to 'https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-iosx64-iosX64/4.10.1/engineCodingProject-iosx64-iosX64-4.10.1-sources.jar'.
    Uploading to https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-iosx64-iosX64/4.10.1/engineCodingProject-iosx64-iosX64-4.10.1.klib...
    Uploaded to 'https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-iosx64-iosX64/4.10.1/engineCodingProject-iosx64-iosX64-4.10.1.klib'.
    Uploading to https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-iosx64-iosX64/4.10.1/engineCodingProject-iosx64-iosX64-4.10.1.pom...
    Uploaded to 'https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-iosx64-iosX64/4.10.1/engineCodingProject-iosx64-iosX64-4.10.1.pom'.
    Uploading to https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-js-js/4.10.1/engineCodingProject-js-js-4.10.1-sources.jar...
    Uploaded to 'https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-js-js/4.10.1/engineCodingProject-js-js-4.10.1-sources.jar'.
    Uploading to https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-js-js/4.10.1/engineCodingProject-js-js-4.10.1.jar...
    Uploaded to 'https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-js-js/4.10.1/engineCodingProject-js-js-4.10.1.jar'.
    Uploading to https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-js-js/4.10.1/engineCodingProject-js-js-4.10.1.pom...
    Uploaded to 'https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-js-js/4.10.1/engineCodingProject-js-js-4.10.1.pom'.
    Uploading to https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-jvm-jvm/4.10.1/engineCodingProject-jvm-jvm-4.10.1-sources.jar...
    Uploaded to 'https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-jvm-jvm/4.10.1/engineCodingProject-jvm-jvm-4.10.1-sources.jar'.
    Uploading to https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-jvm-jvm/4.10.1/engineCodingProject-jvm-jvm-4.10.1.jar...
    Uploaded to 'https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-jvm-jvm/4.10.1/engineCodingProject-jvm-jvm-4.10.1.jar'.
    Uploading to https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-jvm-jvm/4.10.1/engineCodingProject-jvm-jvm-4.10.1.pom...
    Uploaded to 'https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-jvm-jvm/4.10.1/engineCodingProject-jvm-jvm-4.10.1.pom'.
    Uploading to https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-macos-macos/4.10.1/engineCodingProject-macos-macos-4.10.1-sources.jar...
    Uploaded to 'https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-macos-macos/4.10.1/engineCodingProject-macos-macos-4.10.1-sources.jar'.
    Uploading to https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-macos-macos/4.10.1/engineCodingProject-macos-macos-4.10.1.klib...
    Uploaded to 'https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-macos-macos/4.10.1/engineCodingProject-macos-macos-4.10.1.klib'.
    Uploading to https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-macos-macos/4.10.1/engineCodingProject-macos-macos-4.10.1.pom...
    Uploaded to 'https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-macos-macos/4.10.1/engineCodingProject-macos-macos-4.10.1.pom'.
    Uploading to https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-macosx64-macosX64/4.10.1/engineCodingProject-macosx64-macosX64-4.10.1-sources.jar...
    Uploaded to 'https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-macosx64-macosX64/4.10.1/engineCodingProject-macosx64-macosX64-4.10.1-sources.jar'.
    Uploading to https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-macosx64-macosX64/4.10.1/engineCodingProject-macosx64-macosX64-4.10.1.klib...
    Uploaded to 'https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-macosx64-macosX64/4.10.1/engineCodingProject-macosx64-macosX64-4.10.1.klib'.
    Uploading to https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-macosx64-macosX64/4.10.1/engineCodingProject-macosx64-macosX64-4.10.1.pom...
    Uploaded to 'https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-macosx64-macosX64/4.10.1/engineCodingProject-macosx64-macosX64-4.10.1.pom'.
    Uploading to https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-metadata/4.10.1/engineCodingProject-metadata-4.10.1-sources.jar...
    
    &gt; Task :bintrayUpload FAILED
    Execution failed for task ':bintrayUpload'.
    &gt; Could not upload to 'https://api.bintray.com/content/CodingProjectn/engineCodingProject/engineCodingProject/0.29e/me/CodingProject/engineCodingProject-metadata/4.10.1/engineCodingProject-metadata-4.10.1-sources.jar': HTTP/1.1 409 Conflict [message:Unable to upload files: An artifact with the path 'me/CodingProject/engineCodingProject-metadata/4.10.1/engineCodingProject-metadata-4.10.1-sources.jar' already exists under another version]
    
    
    

As you can see, most of the artifacts are created as intended. But some (iosarm32, iosxX64,...) seem to follow a different pattern. They are not put into the correct subfolder and they carry the 4.10.1 version number I have never configured.

And also: I alsways get this \*already exists under another version\* error.

Whats going wrong here? Why are these artifacts setup with another version number? And why do I get the copying conflict?
## [5][Kotlin behaves weirdly for initialising int variable with 0](https://www.reddit.com/r/Kotlin/comments/erbz73/kotlin_behaves_weirdly_for_initialising_int/)
- url: https://stackoverflow.com/questions/59822125/kotlin-variable-initialization-for-child-class-behaves-weird-for-initializing-va
---

## [6][[Question] How to restore fragment data?](https://www.reddit.com/r/Kotlin/comments/erkqdf/question_how_to_restore_fragment_data/)
- url: https://www.reddit.com/r/Kotlin/comments/erkqdf/question_how_to_restore_fragment_data/
---
Hi! I'm pretty new to Kotlin/Android development and I've been building an app, but now I noticed that when I change Fragments (using a navController) the state of my fragment isn't restored.

Let's say I have a fragment with a simple button that is disabled on click. How would I keep this button disabled when going to a different fragment and then coming back?

Do I also need to account for closing and reopening the app? And what about 'hard resetting' the app?

    class Fragment : Fragment() {
    
        companion object {
            fun newInstance() = Fragment()
        }
    
        private lateinit var viewModel: ViewModel
    
        override fun onCreateView(
            inflater: LayoutInflater, container: ViewGroup?,
            savedInstanceState: Bundle?
        ): View? {
    
            val view: View = inflater.inflate(R.layout.fragment, container, false)
    
            view.fab_start.setOnClickListener { view -&gt;
                updateButton(start = false)
            
            return view
        }
    
        fun updateButton(start: Boolean?){        
            fab_start.isEnabled = start!!
        }
       
    
        override fun onActivityCreated(savedInstanceState: Bundle?) {
            super.onActivityCreated(savedInstanceState)
            viewModel = ViewModelProviders.of(this).get(viewModel::class.java)
            // TODO: Use the ViewModel
        }
    
    }

btw. I have a viewModel in my code because I was following along with a tutorial party before abandoning that, not sure how to use that either (also not really sure I need it?)

Quick Edit: I forgot to mention I read a couple blogs and some stackoverflow posts but I can't make sense of it.

Thanks for the help!
## [7][The Power of Types for Errors](https://www.reddit.com/r/Kotlin/comments/ercrj1/the_power_of_types_for_errors/)
- url: https://developers.soundcloud.com/blog/power-of-types-for-errors
---

## [8][Kotlin getter](https://www.reddit.com/r/Kotlin/comments/erhop2/kotlin_getter/)
- url: https://www.reddit.com/r/Kotlin/comments/erhop2/kotlin_getter/
---
Is there a way to obtain values of members/attributes of an object in the order they are declared in the Class of the Object.
## [9][Any Kotlinx Serialization User Can Help Me With This Problem?](https://www.reddit.com/r/Kotlin/comments/erivu7/any_kotlinx_serialization_user_can_help_me_with/)
- url: https://github.com/Kotlin/kotlinx.serialization/issues/660
---

## [10][Please give this sub an appropriate title image of the kotlin logo!](https://www.reddit.com/r/Kotlin/comments/eqtoru/please_give_this_sub_an_appropriate_title_image/)
- url: https://www.reddit.com/r/Kotlin/comments/eqtoru/please_give_this_sub_an_appropriate_title_image/
---

