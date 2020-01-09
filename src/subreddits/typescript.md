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
## [2][A form that maintains its state between pages. Feedback needed.](https://www.reddit.com/r/typescript/comments/eltcl2/a_form_that_maintains_its_state_between_pages/)
- url: https://www.reddit.com/r/typescript/comments/eltcl2/a_form_that_maintains_its_state_between_pages/
---
I have been recently asked to implement a form with pagination, dynamical elements that would maintain its state.

I have some experience in the Front End field, but I definitely need to learn a lot. 

I would appreciate it if you could take a look at the code and let me know what you think.

https://github.com/shpotainna/passengers
## [3][Missing basic DOM types in TypeScript project](https://www.reddit.com/r/typescript/comments/em13kc/missing_basic_dom_types_in_typescript_project/)
- url: https://www.reddit.com/r/typescript/comments/em13kc/missing_basic_dom_types_in_typescript_project/
---
When running tests I'm seeing errors for basic browser DOM functions (`ReferenceError: btoa is not defined`), classes, and operations.

```
TS2569: Type 'Uint8Array' is not an array type or a string type. Use compiler option '--downlevelIteration' to allow iterating of iterators.

let base64String = btoa(String.fromCharCode(...new Uint8Array(buffer)));
```

I've tried setting my `tsconfig.json` to may [variations of the compilerOptions](https://www.typescriptlang.org/docs/handbook/compiler-options.html) but none of them seem to work.

```
{
    "compilerOptions": {
        "lib": [
            "es2016",
            "dom"
        ]
    }
}
or
{
    "compilerOptions": {
        "module": "commonjs",
        "target": "es5",
        "lib": ["es6", "dom"],
    }
}

```

What am I missing? Do I need a reference to `lib.dom.d.ts` somewhere?

- https://github.com/basarat/typescript-book/blob/master/docs/types/lib.d.ts.md#lib-option
- https://github.com/Microsoft/TypeScript/blob/master/lib/lib.dom.d.ts
## [4][How do I get the intersection of generic types?](https://www.reddit.com/r/typescript/comments/elnme2/how_do_i_get_the_intersection_of_generic_types/)
- url: https://www.reddit.com/r/typescript/comments/elnme2/how_do_i_get_the_intersection_of_generic_types/
---
Suppose I have:

    type union = string | number;
    type subset = union &amp; string; // subset = string;

I'd like to do something like:

    interface A&lt;T&gt; { }
    type union = A&lt;string&gt; | A&lt;number&gt;;
    type subset = union &amp; A&lt;string&gt;; // expect subset = A&lt;string&gt;

The subset type in this case becomes `(A&lt;string&gt; &amp; A&lt;string&gt;) | (A&lt;string&gt; &amp; A&lt;number&gt;)`, which I don't think is what I want.

Update: I solved the problem using Extract&lt;T, U&gt; and by defining a property on my interface with the generic type:

    interface A&lt;T&gt; { _type: T };
    type union = A&lt;string&gt; | A&lt;number&gt;;
    type subset = Extract&lt;union, A&lt;string&gt;&gt;; // A&lt;string&gt; as expected

It is a bit hacky - in my actual code I have a class, so I defined a private setter with that type and it seems to work. The actual use case is a much more complex type whose details don't belong here, but the above sums it up well.
## [5][How to extend the ExpressJS Request object with Typescript](https://www.reddit.com/r/typescript/comments/elqazk/how_to_extend_the_expressjs_request_object_with/)
- url: https://medium.com/@andrey.palatnyi/how-to-extend-the-expressjs-request-object-with-typescript-26675cda0632?source=friends_link&amp;sk=400142c8ffed21c1354018ab85f28b39
---

## [6][Should you use TypeScript if you know JavaScript well?](https://www.reddit.com/r/typescript/comments/eljud4/should_you_use_typescript_if_you_know_javascript/)
- url: https://medium.com/@Charles_Stover/should-you-use-typescript-if-you-know-javascript-well-85042756e594
---

## [7][Absolute path resolver for Typescript with Babel and Webpack](https://www.reddit.com/r/typescript/comments/ekxkoc/absolute_path_resolver_for_typescript_with_babel/)
- url: https://github.com/turbothinh/ts-babel-webpack-boilerplate
---

## [8][ReactFormHelper: Simple way to build forms in React](https://www.reddit.com/r/typescript/comments/el9zqo/reactformhelper_simple_way_to_build_forms_in_react/)
- url: https://github.com/EvandroLG/ReactFormHelper
---

## [9][How can i override a property type of a merged type?](https://www.reddit.com/r/typescript/comments/ekv0wd/how_can_i_override_a_property_type_of_a_merged/)
- url: https://www.reddit.com/r/typescript/comments/ekv0wd/how_can_i_override_a_property_type_of_a_merged/
---
Hej Typescript people. :)

I need to override an existing type of a property. How can i achieve this?

I've made an codesandbox example: [https://codesandbox.io/s/typescript-playground-export-u2203](https://codesandbox.io/s/typescript-playground-export-u2203)

At the end, the test prop should accept strings and numbers.
## [10][how to import TypeScript library main function and typings from same location](https://www.reddit.com/r/typescript/comments/ekm021/how_to_import_typescript_library_main_function/)
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
## [11][My best VSCode Linting/Formatting configuration for Typescript Projects](https://www.reddit.com/r/typescript/comments/eks5qx/my_best_vscode_lintingformatting_configuration/)
- url: https://medium.com/javascript-in-plain-english/my-best-vscode-linting-formatting-configuration-for-typescript-projects-ef400ed9b78f?source=friends_link&amp;sk=d1a797a25de1e668bcd4e69d247f2be4
---

