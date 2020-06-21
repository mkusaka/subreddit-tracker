# typescript
## [1][Who's hiring Typescript developers - June](https://www.reddit.com/r/typescript/comments/gua247/whos_hiring_typescript_developers_june/)
- url: https://www.reddit.com/r/typescript/comments/gua247/whos_hiring_typescript_developers_june/
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
## [2][How can improve my tsconfig?](https://www.reddit.com/r/typescript/comments/hd176e/how_can_improve_my_tsconfig/)
- url: https://www.reddit.com/r/typescript/comments/hd176e/how_can_improve_my_tsconfig/
---
Hi everybody. I have this repo in typescript and I would like to know how it can be improved. Looking for feedback. This is a package for managing permissions.
## [3][How would you implement a Map with objects (vectors) as keys?](https://www.reddit.com/r/typescript/comments/hcrtdx/how_would_you_implement_a_map_with_objects/)
- url: https://www.reddit.com/r/typescript/comments/hcrtdx/how_would_you_implement_a_map_with_objects/
---
I want to store where figures are on a game board.

Positions on the board are of the type 'Vector' which has an x and an y attribute. The figures are of type Entity, but for simplicitys sake here they are just strings:

    class Vector {
        constructor(readonly x: number, readonly y: number) {
        }
    
        equals(other: this): boolean {
            return this.x === other.x &amp;&amp; this.y === other.y;
        }
    }
    
    const board = new Map&lt;Vector, string&gt;();
    board.set(new Vector(0, 1), "white pawn");
    board.set(new Vector(4, 0), "white king");
    
    console.log(board.get(new Vector(4, 0)));  // prints "undefined" - not what I want

I understand now that the Map class compares keys by their identity / memory address. Two distinct objects are considered not the same key, even if the equals method returns true.

*What would be the most elegant and efficient way to implement a type of Map that considers keys to be equivalent if they are structurally equal or an "equals" method would return true?*

On [stackoverflow](https://stackoverflow.com/questions/57262315/how-to-get-deep-equality-of-keys-in-es6-map-alternative-to-using-complex-object), they suggested using strings as keys. I'm uncomfortable with using strings in program logic. Maybe that's just a bad practice in other programming languages but okay in Javascript/Typescript? The typechecker couldn't warn me if I tried to insert an invalid key like `"Vecdor(y:  4,x=2)"`.

One other way, I could imagine, is a custom Map type that requires keys to implement a "Hashable" interface. Every key would be transformed to a `number` and then that number would be used as a key in an internal, regular `Map&lt;number, Value&gt;`.

Thirdly, I don't think I am the only person who would have a use for such a Map variant. Is there maybe something on npm that acts like that?
## [4][Having a real tough time "optionally" using typescript in my react app. see configs](https://www.reddit.com/r/typescript/comments/hd30z4/having_a_real_tough_time_optionally_using/)
- url: https://www.reddit.com/r/typescript/comments/hd30z4/having_a_real_tough_time_optionally_using/
---
So, I have a react app and my Webstorm is just not appropriately linting things. Everything seems to be building fine, but all my ".js" files are littered with linting errors EVEN though I pre-commit with linting and its fine.

Can you look over my babel, eslintrc file and let me know if anything jumps out as incorrect or "out of ordering". I am trying TRYING to ONLY use typescript on ".ts/.tsx" files and ignore the '.js' files for typing... basically "oh its a .js, ignore typescript linting etc...".

ESLINTRC

const path = require('path')

    module.exports = {
      parserOptions: {
        ecmaVersion: 2019,
        sourceType: 'module',
        ecmaFeatures: {
          jsx: true,
        },
      },
      extends: [
        'eslint:recommended',
        'plugin:react-hooks/recommended',
        'plugin:react/recommended',
        'eslint-config-prettier',
        'eslint-config-prettier/@typescript-eslint',
        'airbnb-base',
        'prettier',
        'prettier/react',
      ],
      globals: {},
      rules: {
        strict: ['error', 'never'],
        'import/prefer-default-export': 1,
        'global-require': 1,
        'react/jsx-key': 1,
        'prefer-destructuring': 1,
      },
      env: {
        browser: true,
        jest: true,
        es2020: true,
      },
      settings: {
        'import/resolver': {
          'babel-module': {
            'import/resolver': {
              src: path.join(__dirname, '/src'),
            },
          },
        },
        react: {
          version: 'detect',
        },
      },
      overrides: [
        {
          files: '**/*.+(ts|tsx)',
          parser: '@typescript-eslint/parser',
          parserOptions: {
            project: './tsconfig.json',
            jsx: true,
          },
          plugins: [
            '@typescript-eslint/eslint-plugin',
            '@typescript-eslint/eslint-recommended',
            '@typescript-eslint/recommended',
          ],
        },
        {
          files: ['**/__tests__/**'],
          settings: {
            'import/resolver': {
              jest: {
                jestConfigFile: path.join(__dirname, '/jest.config.js'),
              },
            },
          },
        },
      ],
    }

HERE IS MY BABELRC FILE

      {
      "presets": [
          "@babel/preset-react",
          "@babel/preset-env",
          "@babel/preset-typescript"
      ],
      "plugins": [
        "@babel/plugin-proposal-optional-chaining",
        "@babel/plugin-proposal-nullish-coalescing-operator",
        "@babel/plugin-syntax-dynamic-import",
        ["module-resolver", {
          "root": ["./src"],
          "alias": {
            "src": "./src"
          }
        }]
      ]
    }

HERE IS WHAT MY WEBSTORM IS SHOWING AS ERRORS. These are ".js" files. I have used Webstorm on all my projects and never any issues... even when jump to other repos, all good. So, I am thinking its my introducing with typescript to this project,  as my other projects do not use it. Anybody see any issues? Remember, I want to use both ".js" and ".ts" extensions.. and only "type" the ts files.

errors webstorm show. Its crazy. It doesn't seem to recognize the js/jsx, YET I verified over and over that i have the proper settings (see last image).

[https://imgur.com/a/xYJowaO](https://imgur.com/a/xYJowaO)

[https://imgur.com/uJSIFyj](https://imgur.com/uJSIFyj)

[https://imgur.com/a/Of3f8ah](https://imgur.com/a/Of3f8ah)

  
P.S. --- IF you want to work thru this with me, let me know and I will venmo you some $ for your time. We can google hangout and I can share my screen etc... message me. I do NOT need coding help, I am in need of configurations help.
## [5][Noobie trying to do something advanced I guess - conditional arguments](https://www.reddit.com/r/typescript/comments/hcpm4d/noobie_trying_to_do_something_advanced_i_guess/)
- url: https://www.reddit.com/r/typescript/comments/hcpm4d/noobie_trying_to_do_something_advanced_i_guess/
---
I'm learning typescript and trying to use it on a personal project.I have this function:

```
const fn = ({
   resourceId,
   select = resourceId ? state =&gt; state : id =&gt; state =&gt; state[id],
}) =&gt; {
   const selector = state =&gt;`
   resourceId ? select(resourceId(state) : select(state);)`
};
```
and I cannot make it build successfully. Here is my best attempt:

```
type Attr = string | number;
type LikeState = {
   [key: string]: any;
   [key: number]: any;
};

interface IfcFn&lt;ResourceId extends Attr | undefined&gt; {
   resourceId?: ResourceId;
   select: ResourceId extends Attr
      ? (x: Attr) =&gt; (y: LikeState) =&gt; LikeState
      : (x: LikeState) =&gt; LikeState;
}

const fn = ({
   resourceId,
   select = resourceId ? state =&gt; state : id =&gt; state =&gt; state[id],
}: IfcFn&lt;Attr | undefined&gt;) =&gt; {
   const selector = (state: LikeState) =&gt; 
      resourceId ? select(resourceId)(state) : select(state);
};
```

I know I'm not there. So, any suggestions? Thank you!
## [6][Help with handling JSON array of objects (merge two arrays)](https://www.reddit.com/r/typescript/comments/hc5odl/help_with_handling_json_array_of_objects_merge/)
- url: https://www.reddit.com/r/typescript/comments/hc5odl/help_with_handling_json_array_of_objects_merge/
---
I have a json array of Objects and need to pick a corresponding object in that array according to some calculations.  

My json have a following structure: {"OD": number, "SCH" {"SCH40": number, "SCH80": number}}  

I need to do some math with OD and SCH numbers, compare it with an outside number and find corresponding value for OD key.  

I can get a new array (of numbers I compare the outside number with):  

`this.Pipeareas = this.Pipesizes.map (x =&gt; Math.pow((x.OD - 2 * x.SCH[this.selectedSCH])/2000,2) * Math.PI);`   

How to merge existing and new arrays?  

I tried (this add extra existing array and new object to new array (not key:value pair))     

`this.Pipeareas = {...this.Pipesizes, Area: this.Pipesizes.map (x =&gt; Math.pow((x.OD - 2 * x.SCH[this.selectedSCH])/2000,2) * Math.PI) }`  

and (this adds new key "Area" to each object but the value is full array, not corresponding numbers from new array)     

`this.Pipesizes.map (x =&gt; x.Area = this.Pipesizes.map (x =&gt; Math.pow((x.OD - 2 * x.SCH[this.selectedSCH])/2000,2) * Math.PI))`  

The result should be like:  

`return this.Pipeareas.find(x =&gt; x &gt;= this.OutsideNumber);`
## [7][Abstract class "TypeError: this is null"](https://www.reddit.com/r/typescript/comments/hca7qq/abstract_class_typeerror_this_is_null/)
- url: https://www.reddit.com/r/typescript/comments/hca7qq/abstract_class_typeerror_this_is_null/
---
Going through one of my m8s code I found two classes with basically the same properties and methods and I though "Why not abstractifizzle this shizzle" which I thought would be all rainbows and flowers until this was null on line 75

      abstract addEvent(hour: number, minute: number, offset: number): void;
      abstract calculateEventPositions(): void;

      refresh() {
          this.calculateEventPositions();
      }

      eventPositionsChanged(handled: () =&gt; void) {
          this.calculateEventPositions(); // This is line 75
          handled();
      }

which is implemented like this in the class that extends the abstract one

        calculateEventPositions() {
            CellUtils.calculateEventPositions(this.entity, this.events, CalendarMode.Vertical);
        }

but aparrently not?

It probably or not helps to say that this in a Vue component where the abstract class extends Vue. And Vue Property Decorators are used
## [8][Typing promisify with variadic tuple types](https://www.reddit.com/r/typescript/comments/hbpbwq/typing_promisify_with_variadic_tuple_types/)
- url: https://oida.dev/variadic-tuple-types-preview/
---

## [9][Why You Should Use TypeScript in 2020](https://www.reddit.com/r/typescript/comments/hbhfsh/why_you_should_use_typescript_in_2020/)
- url: https://serokell.io/blog/why-typescript
---

## [10][[Help Wanted] Mutation Testing with Typescript](https://www.reddit.com/r/typescript/comments/hbqgfg/help_wanted_mutation_testing_with_typescript/)
- url: https://www.reddit.com/r/typescript/comments/hbqgfg/help_wanted_mutation_testing_with_typescript/
---
Supposedly Stryker as it working for nearly 3 years: https://stryker-mutator.io/blog/2017-10-06/typescript-support

There's even a [short guide](https://github.com/stryker-mutator/stryker-handbook/blob/master/stryker/guides/vuejs.md#jest-configuration) suitable for [my small project](https://gitlab.com/bss03/vue-webpack-ts-lambda).

Unfortunately I'm experiencing a failure before even the first test: https://pastebin.com/WWdPEXf4

I commented on a closed [issue](https://github.com/stryker-mutator/stryker/issues/1229#issuecomment-645721873), but it hasn't been reopened (nor has anyone commented on it since I did).

Anyone got mutation testing (Skryker or otherwise) working with typescript that could point me the right direction?  It looks like I might need to tell stryker to do some transpiling, but I'm not sure.  I think I tried `"transpiler": [ "typescript" ]` in the Stryker configuration, but it didn't help.

Thanks for reading, and thanks in advance for any help you can provide.
## [11][Those of you who came from dynamic languages, how did you survive without compile-time checks and autocomplete?](https://www.reddit.com/r/typescript/comments/hbhajh/those_of_you_who_came_from_dynamic_languages_how/)
- url: https://www.reddit.com/r/typescript/comments/hbhajh/those_of_you_who_came_from_dynamic_languages_how/
---
So, I come from a Java / Android / Kotlin background, where the IDE is top notch and warns you over every little mistake. Maybe it warns you a bit too much, but at least I know what I'm doing wrong.

For the last 2 years I've been working with Javascript and Ruby folk. Fortunately, on my own projects I'm allowed to write Typescript, but once in awhile I have to cross over to their projects and it's so frustrating, both for me and for them.

- Sometimes they import files that don't exist. 
- ~~Sometimes~~ Most of the time they spell variable / function names wrong and the code doesn't work.
- They can't remember the whether it's an array or an object or a hash or if it's "user_id" or "userID" or "userId" or "user.id" or "user['id']".
- null pointer bugs. So many.
- Forgetting to pass variables into functions / passing in the wrong kind
- If the documentation isn't good they're screwed, because they have no autocomplete and they can't navigate to the source file to read it.
- Renaming files. Painful Ctrl-F everywhere. Half the time they'll still miss out some references anyway.
- Moving file is more painful because I have to watch them do mental sums on how many '../../' to type.
- Nobody can remember what data their objects contain after going on vacation and coming back.

I get really annoyed whenever I have to interact with their code bases. Either I make so many stupid bugs that would never happen in Typescript, or I have to sit through and painfully watch them make stupid bugs that would never happen in Typescript and waste half a day over stupid things like wrong variable names.

Been trying to convince them to use Typescript because it feels like you're coding blindly without it, but all of them just don't see any point to it? I don't know. Are they just too blind to realize the advantages? Or maybe I'm the one that's too dumb and everyone else is really smart and that's why I need a crutch like Intellisense.

Those of you who came from non-typed backgrounds like Javascript / Ruby, how on earth did you survive?
