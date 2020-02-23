# typescript
## [1][Who's hiring Typescript developers - February](https://www.reddit.com/r/typescript/comments/ewxjh2/whos_hiring_typescript_developers_february/)
- url: https://www.reddit.com/r/typescript/comments/ewxjh2/whos_hiring_typescript_developers_february/
---
The monthly thread for people to post openings at their companies.

* Please state the job location and include the keywords REMOTE, INTERNS and/or VISA when the corresponding sort of candidate is welcome. When remote work is not an option, include ONSITE.

* Please only post if you personally are part of the hiring company—no recruiting firms or job boards **Please report recruiters or job boards**. 

* Only one post per company. 

* If it isn't a household name, explain what your company does. Sell it.

* Please add the company email that applications should be sent to, or the companies application web form/job posting (needless to say this should be on the company website, not a third party site).


Commenters: please don't reply to job posts to complain about something. It's off topic here.

Readers: please only email if you are personally interested in the job. 

Posting top level comments that aren't job postings, [that's a paddlin](https://i.imgur.com/FxMKfnY.jpg)

[Previous Hiring Threads](https://www.reddit.com/r/typescript/search?sort=new&amp;restrict_sr=on&amp;q=flair%3AMonthly%2BHiring%2BThread)
## [2][A naive approach to functional programming using TypeScript](https://www.reddit.com/r/typescript/comments/f7pr4k/a_naive_approach_to_functional_programming_using/)
- url: https://github.com/alohawav/naive_functional_programming
---

## [3][How not to trick TypeScript compiler and not be tricked by it](https://www.reddit.com/r/typescript/comments/f7c0x4/how_not_to_trick_typescript_compiler_and_not_be/)
- url: https://indepth.dev/how-not-to-trick-typescript-compiler-and-not-be-tricked-by-it/
---

## [4][redturn.ts - Event-based/Queueing redis mutex lock in typescript](https://www.reddit.com/r/typescript/comments/f7inmh/redturnts_eventbasedqueueing_redis_mutex_lock_in/)
- url: https://github.com/kevinwilson541/redturn.ts
---

## [5][Location.state typing with React-Router](https://www.reddit.com/r/typescript/comments/f7fk08/locationstate_typing_with_reactrouter/)
- url: https://www.reddit.com/r/typescript/comments/f7fk08/locationstate_typing_with_reactrouter/
---
Can anyone help me out with using react-router with Typescript? See this link: [https://stackoverflow.com/questions/60343598/react-typescript-how-add-types-to-location-state-when-passed-in-a-route](https://stackoverflow.com/questions/60343598/react-typescript-how-add-types-to-location-state-when-passed-in-a-route)
## [6][Announcing TypeScript 3.8](https://www.reddit.com/r/typescript/comments/f71hpa/announcing_typescript_38/)
- url: https://devblogs.microsoft.com/typescript/announcing-typescript-3-8/
---

## [7][Typing the Technical Interview in TypeScript: solve an interview with no runtime code!](https://www.reddit.com/r/typescript/comments/f6v4np/typing_the_technical_interview_in_typescript/)
- url: https://gal.hagever.com/posts/typing-the-technical-interview-in-typescript/
---

## [8][AdminBro v2 beta is out!!!](https://www.reddit.com/r/typescript/comments/f6szkd/adminbro_v2_beta_is_out/)
- url: https://www.reddit.com/r/typescript/comments/f6szkd/adminbro_v2_beta_is_out/
---
AdminBro is an automated admin interface for Node.js apps based on React, written in TypeScript = Django admin for node on steroids !!!

Now, we almost finish the entire scope for v2.0 and we released it along with the demo version: 

Demo page can be found here:

[https://admin-bro-example-app-staging.herokuapp.com/admin](https://admin-bro-example-app-staging.herokuapp.com/admin)

Github repo is here: [https://github.com/SoftwareBrothers/admin-bro](https://github.com/SoftwareBrothers/admin-bro)

and you can install it using: 

```
yarn add admin-bro@2.0.0-beta.26
```

In this version: 

* there is an entire Design System written specially for AdminBro
* i18n support
* hooks support
* the entire UI change to Drawer'like exploring resources

Now we are working on the documentation which (WiP) can be fond here:

[https://softwarebrothers.github.io/admin-bro-dev/v2/](https://softwarebrothers.github.io/admin-bro-dev/v2/)

Let me know what do you think - and star the repo if you like the idea :)
## [9][OpenAPI Generator now supports Angular 9](https://www.reddit.com/r/typescript/comments/f6u6jr/openapi_generator_now_supports_angular_9/)
- url: https://www.reddit.com/r/typescript/comments/f6u6jr/openapi_generator_now_supports_angular_9/
---
Thanks for the enhancement by [topce](https://github.com/topce)

Here are 3 simple steps to generate TypeScript Angular 9 client given an OpenAPI spec (e.g. [Petstore API](https://raw.githubusercontent.com/OpenAPITools/openapi-generator/master/modules/openapi-generator/src/test/resources/2_0/petstore.yaml)):

1. Download the Java JAR (Snapshot) ([https://oss.sonatype.org/content/repositories/snapshots/org/openapitools/openapi-generator-cli/4.3.0-SNAPSHOT/openapi-generator-cli-4.3.0-20200220.141445-113.jar](https://oss.sonatype.org/content/repositories/snapshots/org/openapitools/openapi-generator-cli/4.3.0-SNAPSHOT/openapi-generator-cli-4.3.0-20200220.141445-113.jar))
2. Rename the JAR as "openapi-generator-cli.jar"
3. Run the following command to generate a TypeScript Angular 9 client for the Petstore API ([https://raw.githubusercontent.com/OpenAPITools/openapi-generator/master/modules/openapi-generator/src/test/resources/2_0/petstore.yaml](https://raw.githubusercontent.com/OpenAPITools/openapi-generator/master/modules/openapi-generator/src/test/resources/2_0/petstore.yaml)):

Mac/Linux:

\- java -jar openapi-generator-cli.jar generate -g typescript-angular -i [https://raw.githubusercontent.com/OpenAPITools/openapi-generator/master/modules/openapi-generator/src/test/resources/2\_0/petstore.yaml](https://raw.githubusercontent.com/OpenAPITools/openapi-generator/master/modules/openapi-generator/src/test/resources/2_0/petstore.yaml) \-o /var/tmp/typescript-angular/

Windows:

\- java -jar openapi-generator-cli.jar generate -g typescript-angular -i [https://raw.githubusercontent.com/OpenAPITools/openapi-generator/master/modules/openapi-generator/src/test/resources/2\_0/petstore.yaml](https://raw.githubusercontent.com/OpenAPITools/openapi-generator/master/modules/openapi-generator/src/test/resources/2_0/petstore.yaml) \-o C:\\tmp\\typescript-angular

To generate an older version of TypeScript Angular client, use `ngVersion`. For example, to generate Angular 8 client:

\- java -jar openapi-generator-cli.jar generate -g typescript-angular -i [https://raw.githubusercontent.com/OpenAPITools/openapi-generator/master/modules/openapi-generator/src/test/resources/2\_0/petstore.yaml](https://raw.githubusercontent.com/OpenAPITools/openapi-generator/master/modules/openapi-generator/src/test/resources/2_0/petstore.yaml) \-o /var/tmp/typescript-angular/ --additional-properties ngVersion=8


For any questions/feedback, please let us know via [https://github.com/OpenAPITools/openapi-generator/issues/new](https://github.com/OpenAPITools/openapi-generator/issues/new)

Ref: https://github.com/OpenAPITools/openapi-generator/pull/5370
## [10][I programmed in TypeScript like in Haskell (Lazy Evaluation)](https://www.reddit.com/r/typescript/comments/f6t3hj/i_programmed_in_typescript_like_in_haskell_lazy/)
- url: https://www.youtube.com/watch?v=E5yAoMaVCp0&amp;feature=share
---

## [11][Fun with Functors in TypeScript](https://www.reddit.com/r/typescript/comments/f6t6mr/fun_with_functors_in_typescript/)
- url: https://medium.com/@wim.jongeneel1/fun-with-functors-in-typescript-2c3268853d69?source=friends_link&amp;sk=722004cc2e378b41a5a211f47a882432
---

