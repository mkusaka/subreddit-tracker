# typescript
## [1][Who's hiring Typescript developers - April](https://www.reddit.com/r/typescript/comments/fsojx3/whos_hiring_typescript_developers_april/)
- url: https://www.reddit.com/r/typescript/comments/fsojx3/whos_hiring_typescript_developers_april/
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
## [2][Problem Solving with the TypeScript Compiler (Recorded Meetup/Webinar talk)](https://www.reddit.com/r/typescript/comments/g8dlhy/problem_solving_with_the_typescript_compiler/)
- url: https://www.youtube.com/watch?v=ZHiT33F11mk&amp;feature=youtu.be
---

## [3][How to correctly make type files for a typescript module](https://www.reddit.com/r/typescript/comments/g8d2cr/how_to_correctly_make_type_files_for_a_typescript/)
- url: https://www.reddit.com/r/typescript/comments/g8d2cr/how_to_correctly_make_type_files_for_a_typescript/
---
Hello. 

So I'm making a typescript module that will be used by other project (ts and js). And **I CANNOT** make the declaration file work the way it is supposed to.

My module is made out of 3 files :

```
- mainclass.ts
- otherclass.ts
- types.d.ts
```
All my types are generated in the `types.d.ts` like this at top level.
```javascript
export type mainClassConstructor = new (
  config: Config,
) =&gt; void

export interface mainClassInterface {
  config: Config
  getIndex(indexUid: string): Index
  ...
}
```

in my `mainClass.ts` which is the entry point of my module my class is like this: 
```javascript

const mainClass: Types.mainClassConstructor = class implements Types.mainClassInterface {
  config: Types.Config
  constructor(config: Types.Config) {
    this.config = config
  }
  getIndex(indexUid: string): Index {
     return new Index(indexUid);
  }
}
```

Now, when I add this module to another typescript project using `types.d.ts` as typing in `package.json`. Upon creating my first object I receive this error:

```javascript
const myClass = new MainClass(config);
```

&gt; Type 'typeof import("/mymodule-js/src/types")' has no construct signatures.  TS2351


but when I generate the declaration files with the typescript module of rollup and 
use `mainClass.d.ts` as typing file in package.json

Although I can finally use my MainClass, I cannot import any type that is found in types.d.ts 

```
import MainClass, { IndexObject } from 'mymodule';
```
I get this error:

&gt; Module '"/mymodule-js/dist/types/mainclass"' has no exported member 'IndexObject'. Did you mean to use 'import IndexObject from "/mymodule-js/dist/types/meilisearch"'

So my question is: How do I create a proper declaration file that  
- recognizes all types
- Lets you import all types
## [4][6 ways to marry C# with TypeScript](https://www.reddit.com/r/typescript/comments/g7stb9/6_ways_to_marry_c_with_typescript/)
- url: https://www.reddit.com/r/typescript/comments/g7stb9/6_ways_to_marry_c_with_typescript/
---
A popular bundle of .NET + SPA framework written in TypeScript has a notorious problem of enforcing contracts between the back-end and the front-end. 

[Here is my analysis of 6 options](https://alex-klaus.com/marry-csharp-typescript/) to generate TypeScript code from C#: NSwag, Swagger Codegen, OpenAPI Generator, AutoRest, WebApiClientGen and TypeWriter.

Have you used any of these tools, did you try others? Any feedback welcome
## [5][Using fetch with Typescript and the Todoist API](https://www.reddit.com/r/typescript/comments/g7vfkz/using_fetch_with_typescript_and_the_todoist_api/)
- url: https://medium.com/@ricardo.trindade743/using-fetch-with-typescript-and-the-todoist-api-5203c5177ed5
---

## [6][expect-type - intuitive compile-time test assertions, inspired by jest's expect API](https://www.reddit.com/r/typescript/comments/g7uk48/expecttype_intuitive_compiletime_test_assertions/)
- url: https://www.reddit.com/r/typescript/comments/g7uk48/expecttype_intuitive_compiletime_test_assertions/
---
I wrote a small library for validating typescript types: https://github.com/mmkal/ts/blob/master/packages/expect-type/readme.md

I had previously used both `dtslint` and `ts-expect` but found them lacking for a few reasons - dtslint depends on the deprecated tslint, and the assertions are quite crude and comment-based. ts-expect is good but involves a fair amount of boilerplate, making anything but very basic assertions hard to read - and it doesn't support "negative" assertions out of the box. 

- a fluent, jest-inspired API, making the difference between `actual` and `expected` clear. This is helpful with complex types and assertions.
- inverting assertions intuitively and easily via `expectType(...).not`
- first-class support for:
    - `any` (as well as `unknown` and `never`).
      - This can be especially useful in combination with `not`, to protect against functions returning too-permissive types. For example, `const parseFile = (filename: string) =&gt; JSON.parse(readFileSync(filename).toString())` returns `any`, which could lead to errors. After giving it a proper return-type, you can add a test for this with `expect(parseFile).returns.not.toBeAny()`
    - object properties
    - function parameters
    - function return values
    - array item values
    - nullable types
    - assertions on types "matching" rather than exact type equality, for "is-a" relationships e.g. `expectTypeOf(square).toMatchTypeOf&lt;Shape&gt;()`
- built into existing tooling with no dependencies. No extra build step, cli tool, or lint plugin is needed. Just import the function and start writing tests.
## [7][Browser screenshots library written with fp-ts and Fluture](https://www.reddit.com/r/typescript/comments/g7nm1d/browser_screenshots_library_written_with_fpts_and/)
- url: https://github.com/gripeless/pico
---

## [8][A terminal emulator built with TS &amp; React.](https://www.reddit.com/r/typescript/comments/g7hfvy/a_terminal_emulator_built_with_ts_react/)
- url: https://github.com/ctaylo21/termy-the-terminal/
---

## [9][Angular Components with Extracted Immutable State](https://www.reddit.com/r/typescript/comments/g7pu6s/angular_components_with_extracted_immutable_state/)
- url: https://medium.com/@0x1000000/angular-components-with-extracted-immutable-state-86ae1a4c9237?source=friends_link&amp;sk=3d9422a57d8ac49a4b1c8de39d6fc0b3
---

## [10][Compilation error driving me nuts. Any help would be appreciated.](https://www.reddit.com/r/typescript/comments/g7o385/compilation_error_driving_me_nuts_any_help_would/)
- url: https://www.reddit.com/r/typescript/comments/g7o385/compilation_error_driving_me_nuts_any_help_would/
---
I set up a project recently and I'm having issues compiling it and it seems to be because of my tsconfig file and node\_modules. As of now, my project contains a tsconfig file that looks like this:

    {
      "compilerOptions": {
        "target": "es5",
        "lib": ["es6", "dom"],
        "module": "commonjs",
        "sourceMap": true,
        "types": [],
        "skipLibCheck": true
      },
      "exclude": [
        "node_modules"
      ]
    }

as well as a `package.json` and a `package-lock.json` file. Because I installed jest with npm, I have a `node_modules` folder as well as my own `src` folder. In the src folder, I only have one small file that I'm trying to run. When I use the command tsc src/main.ts I get an error: 

    node_modules/@types/babel__template/index.d.ts:16:28 - error TS2583: Cannot find name 'Set'. Do you need to change your target library? Try changing the `lib` compiler option to es2015 or later.
    
    16     placeholderWhitelist?: Set&lt;string&gt;;

I've searched the entire internet for a solution but none seem to work. Any suggestions would be really appreciated.
## [11][Yarn workspaces, create-react-app and express](https://www.reddit.com/r/typescript/comments/g7eq4k/yarn_workspaces_createreactapp_and_express/)
- url: https://www.reddit.com/r/typescript/comments/g7eq4k/yarn_workspaces_createreactapp_and_express/
---
Has anyone gotten yarn workspaces to work well for a single frontend + backend project? I want to have a structure like:

* package.json
* yarn.lock
* client/
   * package.json
   * src/
* server/
   * package.json
   * src/

&amp;#x200B;

But, I don't know how to have the "static files" generated by the client/ "create-react-app" be included as static files in the server/ "express" app. Do I have to do some kind of manual "cp" to get them over? I don't really understand how to configure the "static file" exports on the "client/" project so that the "server/" project can somehow include them and serve them.
