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
## [2][Announcing TypeScript 3.8](https://www.reddit.com/r/typescript/comments/f71hpa/announcing_typescript_38/)
- url: https://devblogs.microsoft.com/typescript/announcing-typescript-3-8/
---

## [3][Typing the Technical Interview in TypeScript: solve an interview with no runtime code!](https://www.reddit.com/r/typescript/comments/f6v4np/typing_the_technical_interview_in_typescript/)
- url: https://gal.hagever.com/posts/typing-the-technical-interview-in-typescript/
---

## [4][GitHub - twitchtv/twirp: A simple RPC framework with protobuf service definitions](https://www.reddit.com/r/typescript/comments/f6rqis/github_twitchtvtwirp_a_simple_rpc_framework_with/)
- url: https://github.com/twitchtv/twirp
---

## [5][I programmed in TypeScript like in Haskell (Lazy Evaluation)](https://www.reddit.com/r/typescript/comments/f6t3hj/i_programmed_in_typescript_like_in_haskell_lazy/)
- url: https://www.youtube.com/watch?v=E5yAoMaVCp0&amp;feature=share
---

## [6][AdminBro v2 beta is out!!!](https://www.reddit.com/r/typescript/comments/f6szkd/adminbro_v2_beta_is_out/)
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
## [7][Fun with Functors in TypeScript](https://www.reddit.com/r/typescript/comments/f6t6mr/fun_with_functors_in_typescript/)
- url: https://medium.com/@wim.jongeneel1/fun-with-functors-in-typescript-2c3268853d69?source=friends_link&amp;sk=722004cc2e378b41a5a211f47a882432
---

## [8][OpenAPI Generator now supports Angular 9](https://www.reddit.com/r/typescript/comments/f6u6jr/openapi_generator_now_supports_angular_9/)
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
## [9][Building an application with Vue and TypeScript. Best practices, thoughts and recommendations.](https://www.reddit.com/r/typescript/comments/f6ghfx/building_an_application_with_vue_and_typescript/)
- url: https://stefan-bauer.online/building-an-application-with-vue-and-type-script-best-practices-thoughts-and-recommendations/
---

## [10][Typesafe replacement for node's EventEmitter embracing functional programming](https://www.reddit.com/r/typescript/comments/f6d6r7/typesafe_replacement_for_nodes_eventemitter/)
- url: https://github.com/garronej/ts-evt
---

## [11][How do I disable ESLint typescript for JS files?](https://www.reddit.com/r/typescript/comments/f6p73o/how_do_i_disable_eslint_typescript_for_js_files/)
- url: https://www.reddit.com/r/typescript/comments/f6p73o/how_do_i_disable_eslint_typescript_for_js_files/
---
Hello all,

I am starting with Typescript and I would like to disable ESLint errors related to Typescript for JS files, my .eslintrc.json contains the following:

```
    "parser":  "@typescript-eslint/parser",
    "extends": [
        "airbnb",
        "prettier",
        "prettier/react",
        "plugin:jsx-a11y/recommended",
        "plugin:jest/recommended",
        "plugin:@typescript-eslint/recommended"
    ],
    "overrides": [
      {
        "files": ["*.js", "*.jsx"],
        "rules": {
            "@typescript-eslint/...": "off"
        }
      }
    ],

```

But I continue getting eslint messages due to "plugin:@typescript-eslint/recommended". Any help with this please?

Thank you in advance and regards.
