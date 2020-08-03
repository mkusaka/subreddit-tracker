# typescript
## [1][Who's hiring Typescript developers - August](https://www.reddit.com/r/typescript/comments/i1ikj5/whos_hiring_typescript_developers_august/)
- url: https://www.reddit.com/r/typescript/comments/i1ikj5/whos_hiring_typescript_developers_august/
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
## [2][Reflecting business logic rules into the domain models using typescript by Mohsen Sareminia](https://www.reddit.com/r/typescript/comments/i2j9tg/reflecting_business_logic_rules_into_the_domain/)
- url: https://www.reddit.com/r/typescript/comments/i2j9tg/reflecting_business_logic_rules_into_the_domain/
---
Typescript is a great and powerful tool for type checking but it could bother you if you don’t define your type according to business logic rules. If your models don’t reflect business logic rules, after some time it will create a gap between what business logic says and how your code behaves. A three part series explores this topic.

* [Part 1](https://medium.com/@mohsen.sareminia/reflecting-business-logic-rules-into-the-domain-models-using-typescript-part-1-e696773ae4ae?source=friends_link&amp;sk=76085024d8b1a731ed47cabb761d0e3a) Describes how to create self-documented types and models.

* [Part 2](https://medium.com/@mohsen.sareminia/reflecting-business-logic-rules-into-the-domain-models-using-typescript-part-2-61a19fba069d?source=friends_link&amp;sk=d596490ba47c075a0648a31294b5fa33) Explains how to write some code that creates and validates those types.

* [Part 3](https://medium.com/@mohsen.sareminia/reflecting-business-logic-rules-into-the-domain-models-using-typescript-part-3-aa8998bc6d29?source=friends_link&amp;sk=bdec956a522701f196bd8a83a5a50838) Uses io-ts library to automate these boring tasks.
## [3][Is there a smarter way to filter out duplicates?](https://www.reddit.com/r/typescript/comments/i29ou1/is_there_a_smarter_way_to_filter_out_duplicates/)
- url: https://www.reddit.com/r/typescript/comments/i29ou1/is_there_a_smarter_way_to_filter_out_duplicates/
---
I want to remove objects from an array if the `playerId` is already contained in that array.

Is there a smarter way to do this in TypeScript than my current implementation?

    private filterDuplicateLeaderBoardLines(result: Result): Result {
      const leaderBoardLines: LeaderBoardLineResult[] = [];
      const playerIds = new Set&lt;string&gt;();
    
      result.sessionResult.leaderBoardLines.forEach((leaderBoardLine: LeaderBoardLineResult) =&gt; {
        const playerId = leaderBoardLine.currentDriver.playerId;
        playerIds.add(playerId);
      });
    
      playerIds.forEach((playerId) =&gt; {
        leaderBoardLines.push(
          result.sessionResult.leaderBoardLines.find(
            (leaderBoardLine) =&gt; leaderBoardLine.currentDriver.playerId === playerId
          )
        );
      });
    
      result.sessionResult.leaderBoardLines = leaderBoardLines;
    
      return result;
    }
## [4][[Question] ReactJS + TypeScript = larger bundle size](https://www.reddit.com/r/typescript/comments/i2ck8c/question_reactjs_typescript_larger_bundle_size/)
- url: https://www.reddit.com/r/typescript/comments/i2ck8c/question_reactjs_typescript_larger_bundle_size/
---
I am looking forward to using TypeScript + ReactJS in my next project. I know TS adds a lot of good stuff into the app but I wonder about the bundled size of the project.

How much bigger in a comparison between a normal ReactJS and ReactJS + TypeScript?

Thank you.
## [5][Help with getting IDE support working for a .d.ts for a JS project, in another JS project?](https://www.reddit.com/r/typescript/comments/i2fqvw/help_with_getting_ide_support_working_for_a_dts/)
- url: https://www.reddit.com/r/typescript/comments/i2fqvw/help_with_getting_ide_support_working_for_a_dts/
---
Situation:

I have a node module, pure node javascript, no bundler, packager, build step etc.  I've written a .d.ts file for it.

I want to use it from another node javascript project, but I want IDE support to work correctly.

Difficulty:  When I do const x=require('lib'); ... vs code tells me that "x" only has one property: default .  If I do 'x.default.' then I get proper intellisense, however, x.default is not how the library is exported, so that code doesn't actually work.

If I change it to "import x from 'lib';" then intellisense works properly, but of course, that's not valid usage from a CJS node package, so that also doesn't work.

Any help how to configure tsconfig / jsconfig / VSC into such a way that this will act how I expect it to?  

Someone has suggested that I add a line
    /** @type {typeof import("lib").default} */
to each place where it's imported, which apparently forces TS to understand it properly... but that's not exactly sustainable across a large project..
## [6][Gamedev Patterns and Algorithms in Action with TypeScript. Game Loop 2/2](https://www.reddit.com/r/typescript/comments/i20o2k/gamedev_patterns_and_algorithms_in_action_with/)
- url: https://medium.com/@gregsolo/gamedev-patterns-and-algorithms-in-action-with-typescript-game-loop-2-2-c0d57a8e5ec2
---

## [7][How to infer type of Array.prototype.map.call()?](https://www.reddit.com/r/typescript/comments/i24akt/how_to_infer_type_of_arrayprototypemapcall/)
- url: https://www.reddit.com/r/typescript/comments/i24akt/how_to_infer_type_of_arrayprototypemapcall/
---
It appears that TypeScript does not infer the return type of `Array.prototype.map.call()` -- it believes that it returns `unknown[]`.

```
let input: string[] = ["1", "2", "3"];
// OK
let result1: number[] = input.map(str =&gt; Number(str));
// error TS2322: Type 'unknown[]' is not assignable to type 'number[]'
let result2: number[] = Array.prototype.map.call(input, str =&gt; Number(str));
```

What can I do to make TypeScript correctly infer the type of `result2`? I'd rather not use type assertions to explicitly override the return type of `Array.prototype.map.call`.

(Running TypeScript 3.9.7, btw.)
## [8][Git ignore js build files?](https://www.reddit.com/r/typescript/comments/i1yhu4/git_ignore_js_build_files/)
- url: https://www.reddit.com/r/typescript/comments/i1yhu4/git_ignore_js_build_files/
---
Pretty new to this all just wondering if it's a good idea to gitignore the build folder all the JS gets rendered into?

Looked around on stack overflow and not a lot of guidance there. Thanks!
## [9][Why intellisense is not working here ?](https://www.reddit.com/r/typescript/comments/i1tq7w/why_intellisense_is_not_working_here/)
- url: https://www.reddit.com/r/typescript/comments/i1tq7w/why_intellisense_is_not_working_here/
---
    const oldValue = [
    	{ a: 1 },
    	{
    		b: {
    			c: 2,
    			d: [3, 4],
    		},
    		f: 5,
    	},
    	6,
    	[7, 8, 9],
    ];
    
    oldValue[0]. //no intellisense
## [10][How do you generate, organize and maintain your seeds/fixtures data for integration or e2e testing](https://www.reddit.com/r/typescript/comments/i1nmsk/how_do_you_generate_organize_and_maintain_your/)
- url: https://www.reddit.com/r/typescript/comments/i1nmsk/how_do_you_generate_organize_and_maintain_your/
---
Hello all,

I'd like to write mainly e2e and integration tests but the main challenge to me is to put my system in a certain state.

I'm unable to find good resources about how to generate, organize and maintain fixtures / seeds data.

Everytime I need to start writing an e2e test this part is the most challenging.

I just want to know if there are better alternatives than manually writing your set of fixtures that you reuse and modify per test case.

Is reusing other parts of your own API to generate data in other tests a bad thing? Like imagine testing the case "As a user I can only see my own books", do you insert a user and books in the database directly before running the test or do you use your own API operations like "user register", "user add books" to populate the database?

Do you always start from an empty database or do you have a consistent set of data that you reset between tests? (The consistent set seems easier from a test preparation perspective but is also less explicit as you can't see the whole picture in the test case).
## [11][Problem exporting typings from a package: module is imported from 'package-name/build'](https://www.reddit.com/r/typescript/comments/i1ozwt/problem_exporting_typings_from_a_package_module/)
- url: https://www.reddit.com/r/typescript/comments/i1ozwt/problem_exporting_typings_from_a_package_module/
---
Hi, I'm having trouble configuring the package.json and tsconfig.json for a package with the following structure:

    package-name
        build               \
            lib             |
                index.js    |
            test            | tsc output
                specs.js    |
            index.js        /

        src                 \
            lib             |
                index.ts    |
            test            | source files
                specs.ts    |
            index.ts        /
        package.json
        tsconfig.json

When I'm trying to import something from such a package, the auto-import (I'm using Webstorm if that makes a difference)
is always something like `import { Something } from 'package-name/build';` while I would want it to be simply
`import { Something } from 'package-name';`.

Imports from `package-name/build` seems to work at runtime, but my IDE can not resolve it properly. If I manually fix
the import from `package-name/build` to `package-name`, my IDE can resolve the types and it also works fine at runtime.

I tried a lot of things with the `main` and `types` properties in the package.json, also with the `sourceRoot`,
`baseUrl` and `rootDir` properties in the tsconfig.json but to no avail, I never succeeded at making auto-imports be
imported from `package-name`.

You can try it yourself by pulling [this repo](https://gitlab.com/gpascal/export-types-example) (which uses yarn workspaces to link the packages locally):

    git clone https://gitlab.com/gpascal/export-types-example
    cd export-types-example
    yarn install

Then open the project in your IDE, go to `packages/package-import/src/lib/index.ts`, delete the import statement and try to auto-import `TestType` or `testFunction`.

Any help would be greatly appreciated.

---------------------

For reference, the tsconfig.json and package.json look like this:

tsconfig.json:

    {
        "compileOnSave": true,
        "compilerOptions": {
            "target": "es2017",
            "module": "commonjs",
            "outDir": "build",
            "noImplicitAny": false,
            "moduleResolution": "node",
            "inlineSources": true,
            "inlineSourceMap": true
        },
        "include": ["src"]
    }

package.json:

    {
        "name": "package-name",
        "version": "1.0.0",
        "main": "build/index.js",
        "types": "src/index.ts",
        "devDependencies": {
            "@types/node": "^14.0.27",
            "typescript": "^3.9.7"
        }
    }


Edit: added link to the test repo
