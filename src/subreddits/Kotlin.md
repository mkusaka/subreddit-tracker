# Kotlin
## [1][Single Responsibility with Kotlin](https://www.reddit.com/r/Kotlin/comments/i88r1p/single_responsibility_with_kotlin/)
- url: https://medium.com/@fanisveizis/kode-quality-single-responsibility-9fb9599fd4e2
---

## [2][Current best frontend for a kotlin backend?](https://www.reddit.com/r/Kotlin/comments/i89q9p/current_best_frontend_for_a_kotlin_backend/)
- url: https://www.reddit.com/r/Kotlin/comments/i89q9p/current_best_frontend_for_a_kotlin_backend/
---
Title. I've been working on a project with a ktor backend, but I'm having trouble finding a decent frontend. Is the best frontend for my project even within the Kotlin realm?
## [3][Turning a Camera Feed into a Solved Sudoku](https://www.reddit.com/r/Kotlin/comments/i8dl9w/turning_a_camera_feed_into_a_solved_sudoku/)
- url: https://www.kotlindevelopment.com/turning-a-camera-feed-into-a-solved-sudoku/
---

## [4][Kotlin - coroutines and channels](https://www.reddit.com/r/Kotlin/comments/i8dgtn/kotlin_coroutines_and_channels/)
- url: https://www.reddit.com/r/Kotlin/comments/i8dgtn/kotlin_coroutines_and_channels/
---
I have an app where one of the features is bulk import. I want to do it so that the user can upload 10k records, another one can make a similar request at the same time, and I want to do it all in the background. the flow should go like this, user1 imports 10k records which I will validate/produce and send it to the Channel -&gt; from another function I should receive data and save it to my DB -&gt; while data is being saved to the DB, another user makes a request for 5k records, again data goes through validate/produce and after all is good I send it to the same Channel -&gt; when consumer is done with the saving of user1 data, he should start with save of the next data from Channel. I guess my question is how can I make this flow that consumer always listens when Channel got some data so he can save i to the DB, and I would like to do it in the background so user can still use app, and another users can also import bulks.

Here is the code I currently have, I know its not good, and for some reason it works. My due date is very close, so I would appreciate any help. Thanks

[CODE](https://pl.kotl.in/_7VHabQ6y)
## [5][Type Aliases vs. Inline Classes](https://www.reddit.com/r/Kotlin/comments/i8aawg/type_aliases_vs_inline_classes/)
- url: https://medium.com/@forketyfork/kotlin-type-aliases-and-inline-classes-dd1b42e24bfb
---

## [6][Native Binary for Kotlin+TonrnadoFX app using GraalVM](https://www.reddit.com/r/Kotlin/comments/i8cthy/native_binary_for_kotlintonrnadofx_app_using/)
- url: https://www.reddit.com/r/Kotlin/comments/i8cthy/native_binary_for_kotlintonrnadofx_app_using/
---
Has anyone generated native binaries  (MacOS) for a Kotlin +TornadoFX App using GraalVM native image capabilities? I am trying with JDK 11 but am open to switch to any JDK from 8-14 if the binary can be created. The other suboptimal option is using \`jlink \` to create a lighter distribution of the application, which also didn't work for me :'(
## [7][Katlib - extensions functions library](https://www.reddit.com/r/Kotlin/comments/i7oa3d/katlib_extensions_functions_library/)
- url: https://www.reddit.com/r/Kotlin/comments/i7oa3d/katlib_extensions_functions_library/
---
[https://github.com/LukasForst/katlib](https://github.com/LukasForst/katlib)

&amp;#x200B;

Allow me to introduce you to **Katlib** \- collection of extension functions I and my colleges wrote for last three years of server side Kotlin development. It contains around 70 extensions or functions that we're missing in the Kotlin standard library.

Fully opensourced, with test coverage between 60-70% and with minimum (1) of dependencies.
## [8][I wrote an article about higher-order functions and how to compose them. Hope it's useful.](https://www.reddit.com/r/Kotlin/comments/i7tdo3/i_wrote_an_article_about_higherorder_functions/)
- url: https://farhanpatel.dev/index.php/2020/08/10/composing-higher-order-functions-in-kotlin/
---

## [9][Quick intro into http4k — a Kotlin library for writing HTTP servers and clients](https://www.reddit.com/r/Kotlin/comments/i7ojmg/quick_intro_into_http4k_a_kotlin_library_for/)
- url: https://dmitrykandalov.com/hello-http4k
---

## [10][Midifunk - a library for processing MIDI events easily](https://www.reddit.com/r/Kotlin/comments/i7ctex/midifunk_a_library_for_processing_midi_events/)
- url: https://github.com/stashymane/midifunk
---

