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
## [2][Which ORM should I use?](https://www.reddit.com/r/typescript/comments/icbey2/which_orm_should_i_use/)
- url: https://www.reddit.com/r/typescript/comments/icbey2/which_orm_should_i_use/
---
Which ORM should I use to handle a SQL database in a TypeScript project?
## [3][Dynamic type systems are not inherently more open](https://www.reddit.com/r/typescript/comments/icd0sn/dynamic_type_systems_are_not_inherently_more_open/)
- url: https://lexi-lambda.github.io/blog/2020/01/19/no-dynamic-type-systems-are-not-inherently-more-open/
---

## [4][How to enforce nested objects have matching keys?](https://www.reddit.com/r/typescript/comments/iccp8r/how_to_enforce_nested_objects_have_matching_keys/)
- url: https://www.reddit.com/r/typescript/comments/iccp8r/how_to_enforce_nested_objects_have_matching_keys/
---
I am trying to create a localization object for my React Native app. This is the structure:

    const strings = {
      ...
    
      en: {
        string_1: "foo",
        string_2: "bar",
        ...
      },
    
      fr: { ... },
    
      ...
    }

I want to enforce through TypeScript that each locale matches the other in terms of keys. For example, if `"en"` has key `"string_1"`, then every other locale should also have key `"string_1"`. Or if `"fr"` has key `"string_3"`, then every other locale should also have key `"string_3"`.

My reasoning is that I don't want to add support for a language that does not have all the keys used in the app, so it would be nice to know if/what is missing.

Is this possible to do? And is it possible to do for an arbitrary amount of languages and keys without having to manually create/update an interface?
## [5][How common are decorators in the wild?](https://www.reddit.com/r/typescript/comments/ic921d/how_common_are_decorators_in_the_wild/)
- url: https://www.reddit.com/r/typescript/comments/ic921d/how_common_are_decorators_in_the_wild/
---
I'm curious. I've been doing a deep dive on them lately, and find them an interesting concept. However I know they are "experimental' so I'd like to know outside of Angular, are they used all that often?
## [6][My TypeScript Exercise Repository.](https://www.reddit.com/r/typescript/comments/icejbf/my_typescript_exercise_repository/)
- url: https://github.com/JaganGanesh/typescript
---

## [7][Trying to use Typescript/Mocha/NYC together but am confused about how coverage is detected](https://www.reddit.com/r/typescript/comments/ic5aa3/trying_to_use_typescriptmochanyc_together_but_am/)
- url: https://www.reddit.com/r/typescript/comments/ic5aa3/trying_to_use_typescriptmochanyc_together_but_am/
---
I've recently attached NYC to my Typescript project but am confused about how it determines when code is covered. Below is an example of NYC claiming that lines 5-6 are not covered in my `src/Indexer/InvalidIndexError.ts`.

I have the following setup:

`.nycrc`

    {
      "extends": "@istanbuljs/nyc-config-typescript",
      "include": [
        "src/**/*.js",
        "src/**/*.ts",
        "src/**/*.tsx"
      ],
      "require": [
        "ts-node/register"
      ],
      "reporter": [
        "text",
        "html"
      ],
      "all": true,
      "check-coverage": true,
      "sourceMap": true,
      "instrument": true
    }

`src/Indexer/InvalidIndexError.ts`

    export default class InvalidIndexError extends Error {
        public static ErrorName = 'InvalidIndexError'
        constructor(message) {
            super(message);
            this.name = InvalidIndexError.ErrorName;
        }
    }

`test/InvalidIndexError.ts`

    import InvalidIndexError from "../src/Indexer/InvalidIndexError";
    
    const chai = require('chai');
    const chaiAsPromised = require("chai-as-promised");
    const {expect} = chai;
    
    chai.use(chaiAsPromised)
    
    describe('InvalidIndexError', () =&gt; {
        it('should have static error name',  () =&gt; {
            return expect(InvalidIndexError).to.have.property('ErrorName')
        })
    
        it('should have a name matching static name',  () =&gt; {
            const error = new InvalidIndexError('');
    
            return expect(error.name).to.equal(InvalidIndexError.name)
        })
    
        it('should be instance of Error', () =&gt; {
            return expect((new InvalidIndexError('')) instanceof Error).to.be.true
        })
    });

After executing  `nyc mocha -r ts-node/register/transpile-only -r source-map-support/register --full-trace --ui bdd tests/**/*.ts` I end up with:

    -----------------------|---------|----------|---------|---------|-------------------
    File                   | % Stmts | % Branch | % Funcs | % Lines | Uncovered Line #s 
    -----------------------|---------|----------|---------|---------|-------------------
    All files              |   38.02 |    39.29 |      35 |   40.27 |                   
     src                   |   30.95 |       40 |   36.36 |   33.33 |  
     ...                   | Excluding other files for Reddit brevity                            
     src/Indexer           |   92.86 |    52.94 |    87.5 |   95.52 |                  
      InvalidIndexError.ts |   80.95 |    46.15 |      75 |   89.47 | 5-6
## [8][Getting [tsl] errors for files inside the node_modules folder.](https://www.reddit.com/r/typescript/comments/ic9xvp/getting_tsl_errors_for_files_inside_the_node/)
- url: https://www.reddit.com/r/typescript/comments/ic9xvp/getting_tsl_errors_for_files_inside_the_node/
---
# Error:

    ERROR in C:\Users\ME\Desktop\TSLbrokenrepo\node_modules\jest-worker\build\index.js
    [tsl] ERROR in C:\Users\ME\Desktop\TSLbrokenrepo\node_modules\jest-worker\build\ind
          TS2578: Unused '@ts-expect-error' directive.
    
    ERROR in C:\Users\ME\Desktop\TSLbrokenrepo\node_modules\jest-worker\build\index.js
    [tsl] ERROR in C:\Users\ME\Desktop\TSLbrokenrepo\node_modules\jest-worker\build\ind
          TS2578: Unused '@ts-expect-error' directive.
    
    ERROR in C:\Users\ME\Desktop\TSLbrokenrepo\node_modules\jest-worker\build\index.js
    [tsl] ERROR in C:\Users\ME\Desktop\TSLbrokenrepo\node_modules\jest-worker\build\ind
          TS2578: Unused '@ts-expect-error' directive.

## webpack.config.js

    const path = require('path');
    const TerserPlugin = require('terser-webpack-plugin');
    
    module.exports = (env, argv) =&gt; {
    
      const isDev = argv.mode === "development";
    
      if (!isDev) {
        console.log("production mode");
      }
    
      return {
        entry: './src/index.ts',
        module: {
          rules: [
            {
              test: /\.tsx?$/,
              use: 'ts-loader',
              exclude: /node_modules/,
            }
          ]
        },
        resolve: {
          extensions: [ '.tsx', '.ts', '.js' ],
        },
        output: {
          filename: 'bundle.js',
          path: path.resolve(__dirname, 'dist'),
        }
      }
    };

## tsconfig.json

    {
      "compilerOptions": {
        "outDir": "./dist/",  
        "noImplicitAny": true,
        "module": "es6",
        "target": "es6",
        "allowJs": true
      }
    }

## index.ts

    console.log("Hello World");

## CLI 

    webpack --mode production --display verbose --display-modules


I dont understand this at all. What is [tsl] even? linter? loader? I have tried everything to disable this. The problem goes away entirely if I comment out `require('terser-webpack-plugin');` but I need this plugin.
## [9][Parse, don’t validate](https://www.reddit.com/r/typescript/comments/ibqabo/parse_dont_validate/)
- url: https://lexi-lambda.github.io/blog/2019/11/05/parse-don-t-validate/
---

## [10][Need help with type error](https://www.reddit.com/r/typescript/comments/ibqhic/need_help_with_type_error/)
- url: https://www.reddit.com/r/typescript/comments/ibqhic/need_help_with_type_error/
---
Im receiving this error

&gt;Type '{ foo: "bar"; }' is not assignable to type 'Partial&lt;T&gt;' 

If every property in `T` is optional and `T` should contain at least `foo: string` then why is the return value invalid?

Maybe the easy solution is to just replace `Partial&lt;T&gt;` with `IFooBar` or `Partial&lt;IFooBar&gt;`, but what if I want this interface to be an abstract method for somebody else to implement?

&amp;#x200B;

BTW I have every strict option enabled

https://preview.redd.it/hagbvyh9knh51.png?width=731&amp;format=png&amp;auto=webp&amp;s=a6a0a6c614845952beb10b7decf8eee430488025
## [11][Let's Write a JavaScript Library in ES6 using Webpack and Babel](https://www.reddit.com/r/typescript/comments/ic0imb/lets_write_a_javascript_library_in_es6_using/)
- url: https://9sh.re/dHgv80rDp
---

