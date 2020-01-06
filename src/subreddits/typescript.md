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
## [2][My best VSCode Linting/Formatting configuration for Typescript Projects](https://www.reddit.com/r/typescript/comments/eks5qx/my_best_vscode_lintingformatting_configuration/)
- url: https://medium.com/javascript-in-plain-english/my-best-vscode-linting-formatting-configuration-for-typescript-projects-ef400ed9b78f?source=friends_link&amp;sk=d1a797a25de1e668bcd4e69d247f2be4
---

## [3][How can I use TS in a PHP project](https://www.reddit.com/r/typescript/comments/ekttnn/how_can_i_use_ts_in_a_php_project/)
- url: https://www.reddit.com/r/typescript/comments/ekttnn/how_can_i_use_ts_in_a_php_project/
---
In my project, some of thr values are gained by php echo  , like let abc=&lt;?php echo $d?&gt;. What is the correct way to handle such scenaio such that I can use TS in these projects?
## [4][how to import TypeScript library main function and typings from same location](https://www.reddit.com/r/typescript/comments/ekm021/how_to_import_typescript_library_main_function/)
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
## [5][Resources to learn Redux with Typescript](https://www.reddit.com/r/typescript/comments/ekb72e/resources_to_learn_redux_with_typescript/)
- url: https://www.reddit.com/r/typescript/comments/ekb72e/resources_to_learn_redux_with_typescript/
---
I am trying to learn redux but almost everything is using Raw javascript. is there any resource which teaches redux with typescript (free or paid)?
## [6][Question: why this won't work?](https://www.reddit.com/r/typescript/comments/ekbl3g/question_why_this_wont_work/)
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
## [7][Self-published my first book! 2020 goals complete :)](https://www.reddit.com/r/typescript/comments/ejydn9/selfpublished_my_first_book_2020_goals_complete/)
- url: https://www.amazon.com/gp/product/B083C5S9LV?pf_rd_p=ab873d20-a0ca-439b-ac45-cd78f07a84d8&amp;pf_rd_r=7R80GRNQ6N1GH5H9TFC4
---

## [8][Things I've Learnt Building Firebase Functions Declarations](https://www.reddit.com/r/typescript/comments/ek3wv7/things_ive_learnt_building_firebase_functions/)
- url: https://dutzi.party/building-firebase-functions-declarations/
---

## [9][Module suffixes problem - VSCode refactoring/servers not using ".js" ?](https://www.reddit.com/r/typescript/comments/ejw340/module_suffixes_problem_vscode_refactoringservers/)
- url: https://www.reddit.com/r/typescript/comments/ejw340/module_suffixes_problem_vscode_refactoringservers/
---
Hi 👋 **Problem**:  I'm using es2015 modules, and targeting web browsers. When I move .ts files in VSCode, the updated module import paths will not have a .js suffix. E.g.: I had a module reference:

* "/js/event.js", upon move to a new directory, the reference was update to
* "/js/util/event"

In other [VSCode extention](https://github.com/stringham/move-ts.git)s to move and refactor files they also seem to use no suffix. I'm using IIS or the *python -m http.server* command to test my code, and these seem to have no easy way to automatically rewrite all /thing to /thing.js.

What am I missing? The servers don't seem to like the syntax VSCode gives, yet I find no way to change the settings so the refactoring retains the .js suffix. Are modules generated by typescript usually not served as .js? If so, what servers should I be using?

Sorry for the broad scope of this question, but it's causing me some issues. I've searched for the answer and also tried rewrite rules.

\---

**Update**: seems there is no easy way to alter this refactoring behavior, but using webpack might be a solution, and one that does not require a large build pipeline and loads of dependencies. Thanks guys!

**Update 2**: the following webpack command makes debugging work as well, though you get the .js source maps instead of the .ts `webpack --entry .\js\main.js -o library.js --mode development --watch --devtool source-map`
## [10][Distributing Pick&lt;T, K&gt;/Omit&lt;T,K&gt; over union types in TypeScript](https://www.reddit.com/r/typescript/comments/ejiudx/distributing_pickt_komittk_over_union_types_in/)
- url: https://davidgomes.com/pick-omit-over-union-types-in-typescript/
---

## [11][A short time learning Haskell can shortly IMPROVE your Typescript skills!](https://www.reddit.com/r/typescript/comments/ejllr2/a_short_time_learning_haskell_can_shortly_improve/)
- url: https://medium.com/@lironhazan/a-short-time-learning-haskell-can-shortly-improve-your-typescript-skills-523505900ac0?source=friends_link&amp;sk=daf9f13302e63282bc7664574873fa8c
---

