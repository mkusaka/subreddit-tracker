# typescript
## [1][Who's hiring Typescript developers - January](https://www.reddit.com/r/typescript/comments/eib2jh/whos_hiring_typescript_developers_january/)
- url: https://www.reddit.com/r/typescript/comments/eib2jh/whos_hiring_typescript_developers_january/
---
The monthly thread for people to post openings at their companies.

* Please state the job location and include the keywords REMOTE, INTERNS and/or VISA when the corresponding sort of candidate is welcome. When remote work is not an option, include ONSITE.

* Please only post if you personally are part of the hiring company—no recruiting firms or job boards **Please report recruiters or job boards**. 

* Only one post per company. 

* If it isn't a household name, explain what your company does. Sell it.

* Please add the company email that applications should be sent to, or the companies application web form/job posting (needless to say this should be on the company website, not a third party site).


Commenters: please don't reply to job posts to complain about something. It's off topic here.

Readers: please only email if you are personally interested in the job. 

[Previous Hiring Threads](https://www.reddit.com/r/typescript/search?sort=new&amp;restrict_sr=on&amp;q=flair%3AMonthly%2BHiring%2BThread)
## [2][ReactFormHelper: Simple way to build forms in React](https://www.reddit.com/r/typescript/comments/el9zqo/reactformhelper_simple_way_to_build_forms_in_react/)
- url: https://github.com/EvandroLG/ReactFormHelper
---

## [3][Absolute path resolver for Typescript with Babel and Webpack](https://www.reddit.com/r/typescript/comments/ekxkoc/absolute_path_resolver_for_typescript_with_babel/)
- url: https://github.com/turbothinh/ts-babel-webpack-boilerplate
---

## [4][How can i override a property type of a merged type?](https://www.reddit.com/r/typescript/comments/ekv0wd/how_can_i_override_a_property_type_of_a_merged/)
- url: https://www.reddit.com/r/typescript/comments/ekv0wd/how_can_i_override_a_property_type_of_a_merged/
---
Hej Typescript people. :)

I need to override an existing type of a property. How can i achieve this?

I've made an codesandbox example: [https://codesandbox.io/s/typescript-playground-export-u2203](https://codesandbox.io/s/typescript-playground-export-u2203)

At the end, the test prop should accept strings and numbers.
## [5][how to import TypeScript library main function and typings from same location](https://www.reddit.com/r/typescript/comments/ekm021/how_to_import_typescript_library_main_function/)
- url: https://www.reddit.com/r/typescript/comments/ekm021/how_to_import_typescript_library_main_function/
---
I am writing a TypeScript library + an example-app, each in separate codebases. I am having trouble getting the library's main function and its typings to be accessible from its import root; I want to be able to consume the library like this:

    import myMainFunc, { InterfaceA, InterfaceB } from 'my-lib';

However, I can only import the types from my-lib, not the main func. I have an entrypoint `index.ts` containing a default export, and a `types.ts` containing all relevant types:

    src/
      index.ts
      types.ts

My package.json has

      "files": [
        "/dist/"
      ],
      "main": "dist/index.js",
      "typings": "dist/types.d.ts",

and my tsconfig has

    {
      "compilerOptions": {
        "lib": ["es6", "dom", "es2017"],
        "target": "es5",
        "module": "commonjs",
        "declaration": true,
        "outDir": "./dist",
        "strict": true
      },
      "include": ["src"],
      "exclude": ["node_modules", "src/test/*"]
    }

The built dist directory looks like:

    dist/
      index.js
      index.d.ts
      types.js
      types.d.ts

Note: the library is not yet published, so I use `yarn link` to pull it in as a dependency to the example app.
## [6][My best VSCode Linting/Formatting configuration for Typescript Projects](https://www.reddit.com/r/typescript/comments/eks5qx/my_best_vscode_lintingformatting_configuration/)
- url: https://medium.com/javascript-in-plain-english/my-best-vscode-linting-formatting-configuration-for-typescript-projects-ef400ed9b78f?source=friends_link&amp;sk=d1a797a25de1e668bcd4e69d247f2be4
---

## [7][How can I use TS in a PHP project](https://www.reddit.com/r/typescript/comments/ekttnn/how_can_i_use_ts_in_a_php_project/)
- url: https://www.reddit.com/r/typescript/comments/ekttnn/how_can_i_use_ts_in_a_php_project/
---
In my project, some of thr values are gained by php echo  , like let abc=&lt;?php echo $d?&gt;. What is the correct way to handle such scenaio such that I can use TS in these projects?
## [8][Resources to learn Redux with Typescript](https://www.reddit.com/r/typescript/comments/ekb72e/resources_to_learn_redux_with_typescript/)
- url: https://www.reddit.com/r/typescript/comments/ekb72e/resources_to_learn_redux_with_typescript/
---
I am trying to learn redux but almost everything is using Raw javascript. is there any resource which teaches redux with typescript (free or paid)?
## [9][Question: why this won't work?](https://www.reddit.com/r/typescript/comments/ekbl3g/question_why_this_wont_work/)
- url: https://www.reddit.com/r/typescript/comments/ekbl3g/question_why_this_wont_work/
---
Trying to play with generic type:


```typescript

function AorB&lt;T extends "a" | "b"&gt;(x: T) {

  const f: T extends "a" 

           ? (x: "b") =&gt; void 

           : (x: "a") =&gt; void = ((x: T) =&gt; {});



  return f

}
```

And got an error:

&gt; Type '(x: T) =&gt; void' is not assignable to type 'T extends "a" ? (x: "b") =&gt; void : (x: "a") =&gt; void'.ts(2322)

I have no idea why this happened, why TypeScript can't infer the type of the function? confused for a whole day...
## [10][Self-published my first book! 2020 goals complete :)](https://www.reddit.com/r/typescript/comments/ejydn9/selfpublished_my_first_book_2020_goals_complete/)
- url: https://www.amazon.com/gp/product/B083C5S9LV?pf_rd_p=ab873d20-a0ca-439b-ac45-cd78f07a84d8&amp;pf_rd_r=7R80GRNQ6N1GH5H9TFC4
---

## [11][Things I've Learnt Building Firebase Functions Declarations](https://www.reddit.com/r/typescript/comments/ek3wv7/things_ive_learnt_building_firebase_functions/)
- url: https://dutzi.party/building-firebase-functions-declarations/
---

