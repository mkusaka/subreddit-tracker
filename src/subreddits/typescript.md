# typescript
## [1][Who's hiring Typescript developers - May](https://www.reddit.com/r/typescript/comments/gb7km3/whos_hiring_typescript_developers_may/)
- url: https://www.reddit.com/r/typescript/comments/gb7km3/whos_hiring_typescript_developers_may/
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
## [2][Making GraphQL Magic with Sqlmancer](https://www.reddit.com/r/typescript/comments/gk2jby/making_graphql_magic_with_sqlmancer/)
- url: https://www.reddit.com/r/typescript/comments/gk2jby/making_graphql_magic_with_sqlmancer/
---
The official beta version of Sqlmancer has been released! Sqlmancer is a Node.js library that empowers you to effortlessly and efficiently translate GraphQL queries into SQL statements. I just wrote up a brief article that shows how to get started and showcases some basic features.

[https://medium.com/@danielrearden/making-graphql-magic-with-sqlmancer-d69c8203c087?sk=e4a66985889eee34f0b67d39f3bd5741](https://medium.com/@danielrearden/making-graphql-magic-with-sqlmancer-d69c8203c087?sk=e4a66985889eee34f0b67d39f3bd5741)

I also put together a CodeSandbox container with the code from the article. You can see the library in action without forking or installing anything :)

The library's still in development, so any feedback is always welcome. If you're interested in contributing, let me know or just open an issue. And if you like the library, please star it on GitHub!
## [3][How to I type a function that it has to not contain return ?](https://www.reddit.com/r/typescript/comments/gk63ea/how_to_i_type_a_function_that_it_has_to_not/)
- url: https://www.reddit.com/r/typescript/comments/gk63ea/how_to_i_type_a_function_that_it_has_to_not/
---
I thought that this :

    const foo:() =&gt; void = function() {
    	return true;
    }
    foo();

would lint error,but it does not.Am I doing something wrong with the type?How to make it lint error ?
## [4][TypeScript v3.9 has a problem when parsing recursive generic type maybe.](https://www.reddit.com/r/typescript/comments/gk5nfm/typescript_v39_has_a_problem_when_parsing/)
- url: https://github.com/samchon/tgrid/issues/34
---

## [5][Using any in TypeScript gives us a false sense of safety](https://www.reddit.com/r/typescript/comments/gk5lit/using_any_in_typescript_gives_us_a_false_sense_of/)
- url: https://mariosfakiolas.com/blog/using-any-in-typescript-gives-us-a-false-sense-of-safety
---

## [6][ts-engine 1.5.0 released with new-package command to bootstrap packages even more quickly, ESLint 7 other some quality of life fixes](https://www.reddit.com/r/typescript/comments/gjvh31/tsengine_150_released_with_newpackage_command_to/)
- url: https://github.com/ts-engine/ts-engine/blob/master/packages/cli/README.md
---

## [7][How do you guys lint a ts file while you are coding it in vscode ?](https://www.reddit.com/r/typescript/comments/gk5lrz/how_do_you_guys_lint_a_ts_file_while_you_are/)
- url: https://www.reddit.com/r/typescript/comments/gk5lrz/how_do_you_guys_lint_a_ts_file_while_you_are/
---
It seems to take way much more time than needed for me to figure it out . I have installed eslint extension in vscode I have also done `npx eslint --init;` . Here is my `eslintrc.json` file :

    {
        "env": {
            "browser": true,
            "es6": true,
    		"node": true
        },
        "extends": [
            "eslint:recommended",
            "plugin:@typescript-eslint/eslint-recommended"
        ],
        "globals": {
            "Atomics": "readonly",
            "SharedArrayBuffer": "readonly"
        },
        "parser": "@typescript-eslint/parser",
        "parserOptions": {
            "ecmaVersion": 11,
            "sourceType": "module"
        },
        "plugins": [
            "@typescript-eslint"
        ],
        "rules": {
    	}
    }

But I do not get all linting in errors that I should in a ts file . For example lets say I have a file with this code :

    let a;

Since I do not use `a` anywhere it should be red highlighted but it is not . When I hover over `a` I can see the error about not being used anywhere;
## [8][Deno 1.0 Released](https://www.reddit.com/r/typescript/comments/gja7lo/deno_10_released/)
- url: https://deno.land/v1
---

## [9][Question: How to structure an NPM package with general utilities](https://www.reddit.com/r/typescript/comments/gjmf3y/question_how_to_structure_an_npm_package_with/)
- url: https://stackoverflow.com/q/61798645/39321
---

## [10][How to get the JSON type of an object type .](https://www.reddit.com/r/typescript/comments/gjpxk9/how_to_get_the_json_type_of_an_object_type/)
- url: https://www.reddit.com/r/typescript/comments/gjpxk9/how_to_get_the_json_type_of_an_object_type/
---
i.e. I have an object type and I want to get the same type without the keys that have function values.
## [11][exporting typescript types in a package](https://www.reddit.com/r/typescript/comments/gjlumo/exporting_typescript_types_in_a_package/)
- url: https://www.reddit.com/r/typescript/comments/gjlumo/exporting_typescript_types_in_a_package/
---
***Outline***

I'm writing a package (let's call it myPackage) for shared react components and their corresponding typescript types. I'm using npm symlink to link myPackage to a typescript/react app I'm working on.

the structure of myPackage is as follows:

    multilist
    --index.tsx
    --types.d.ts
    index.tx 

this is the content of the multilist/types.d.ts file

    type ILookup = {
      readonly id: number;
      readonly text: string;
    };
    
    export type MultilistProps = {
      readonly isEditing: boolean;
      readonly id: string;
      readonly labelText: string;
      readonly value: ILookup | undefined;
      readonly responseOptions: ILookup[];
      readonly singleSelect: boolean;
      readonly invalid: boolean;
      readonly onChangeHandler: (value: any) =&gt; void;
    };

this is the content of the index.ts file

    import { Multilist } from './multilist';
    import { MultilistProps as _MultilistProps } from './multilist/types';
    export { Multilist };
    export type MultilistProps = _MultilistProps;

***Problem***

When I import from myPackage, the type MultilistPropsdoesn't retain it's type structure. A mouseover in VSCode looks like: type MulitlistProps: Any.

If I declare and export the MultilistProps type directly in myPackage/index.ts file; the problem is resolved.

If anybody has any idea's I would be very grateful!  


edit: formatting
