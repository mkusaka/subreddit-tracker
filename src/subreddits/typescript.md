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
## [2][React Native Multibar in TS](https://www.reddit.com/r/typescript/comments/g90g5a/react_native_multibar_in_ts/)
- url: https://medium.com//react-native-multibar-in-ts-d38a26ef4208?source=friends_link&amp;sk=8fe44dad01689ee485bcd4d974a864ef
---

## [3][Problem Solving with the TypeScript Compiler (Recorded Meetup/Webinar talk)](https://www.reddit.com/r/typescript/comments/g8dlhy/problem_solving_with_the_typescript_compiler/)
- url: https://www.youtube.com/watch?v=ZHiT33F11mk&amp;feature=youtu.be
---

## [4][Class pattern matching library useful in useReducer](https://www.reddit.com/r/typescript/comments/g8j92q/class_pattern_matching_library_useful_in/)
- url: https://www.reddit.com/r/typescript/comments/g8j92q/class_pattern_matching_library_useful_in/
---
Hi all, I have written my first npm library npmjs.com/package/@neal83/typeswitch it's only small at 53 lines of code in the main function including type definitions. github.com/neal83/TypeSwitch

But having used redux then moving over to hooks this is my own way of reducing boiler plate while making the data the action type.

You're able to match on class types with the option of using a "when" in combination with default. Like this:

    class MyOwnMarkerClass
    {
        public someProperty: string = "some property value"
    }
    
    const result = TypeSwitch(new MyOwnMarkerClass())(
        Case(Number, When(n: number =&gt; n === 36))
            ((n: Number) =&gt; `you gave me a number and it was 36`),
        Case(Number, When(n: number =&gt; n &gt; 0))
            ((n: Number) =&gt; `you gave me a number and it was positive`),
        Case(String)
            ((s: String) =&gt; `you gave me a string ${s}`),
        Case(MyOwnMarkerClass, When(m =&gt; m.someProperty === "some property value")) //the type is inferred
            ((m: MyOwnMarkerClass) =&gt; `you gave me MyOwnClass with value ${m.someProperty}`)
    );

Hope someone likes it and finds it useful. Thanks
## [5][Trying to use Jest and running into some issues.](https://www.reddit.com/r/typescript/comments/g8gyk3/trying_to_use_jest_and_running_into_some_issues/)
- url: https://www.reddit.com/r/typescript/comments/g8gyk3/trying_to_use_jest_and_running_into_some_issues/
---
I'm just trying to run a simple test file using jest but I keep getting this error:

    Test suite failed to run
    TypeError: PointT_1.default is not a constructor

My PointT.ts file looks like:

    export class PointT {
      readonly x: number;
      readonly y: number;
      public constructor(xc: number, yc: number) {
          this.x = xc;
          this.y = yc;
      }
    ...
    }
    
    const p1: PointT = new PointT(5,5);
    console.log(p1.x);
    
    // running `tsc &amp;&amp; node PointT.js runs perfectly fine

and my PointT.test.ts file in my test folder looks like:

    import { PointT } from "../PointT";
    const p1: PointT = new PointT(5, 5); // the error happens here
    const d1: Interval = { min: 3, max: 5 };
    
    test("Is PointT within Domain?", () =&gt; {
      expect(p1.isWithinDomain(d1)).toBe(true);
    })

any help would be appreciated. If you need to see my config files. Let me know.

&amp;#x200B;

edit: typo in code and added constructor function for clarification
## [6][Trying have our js app communicate with out new angular 9 app](https://www.reddit.com/r/typescript/comments/g8flfy/trying_have_our_js_app_communicate_with_out_new/)
- url: https://www.reddit.com/r/typescript/comments/g8flfy/trying_have_our_js_app_communicate_with_out_new/
---
We are currently upgrading our app to angular 9 and while upgrading, we need the ability for our js to talk to angular 9. We have our js in iframes. I though of dispatching an event.. I also tried setting a window function for each app to access. But am having trouble. Any help? Thanks!
## [7][How to correctly make type files for a typescript module](https://www.reddit.com/r/typescript/comments/g8d2cr/how_to_correctly_make_type_files_for_a_typescript/)
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
## [8][6 ways to marry C# with TypeScript](https://www.reddit.com/r/typescript/comments/g7stb9/6_ways_to_marry_c_with_typescript/)
- url: https://www.reddit.com/r/typescript/comments/g7stb9/6_ways_to_marry_c_with_typescript/
---
A popular bundle of .NET + SPA framework written in TypeScript has a notorious problem of enforcing contracts between the back-end and the front-end. 

[Here is my analysis of 6 options](https://alex-klaus.com/marry-csharp-typescript/) to generate TypeScript code from C#: NSwag, Swagger Codegen, OpenAPI Generator, AutoRest, WebApiClientGen and TypeWriter.

Have you used any of these tools, did you try others? Any feedback welcome
## [9][Using fetch with Typescript and the Todoist API](https://www.reddit.com/r/typescript/comments/g7vfkz/using_fetch_with_typescript_and_the_todoist_api/)
- url: https://medium.com/@ricardo.trindade743/using-fetch-with-typescript-and-the-todoist-api-5203c5177ed5
---

## [10][expect-type - intuitive compile-time test assertions, inspired by jest's expect API](https://www.reddit.com/r/typescript/comments/g7uk48/expecttype_intuitive_compiletime_test_assertions/)
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
## [11][Browser screenshots library written with fp-ts and Fluture](https://www.reddit.com/r/typescript/comments/g7nm1d/browser_screenshots_library_written_with_fpts_and/)
- url: https://github.com/gripeless/pico
---

