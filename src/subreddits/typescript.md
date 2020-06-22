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
## [2][[Newbie] Are getters syntatic sugar?](https://www.reddit.com/r/typescript/comments/hdmea1/newbie_are_getters_syntatic_sugar/)
- url: https://www.reddit.com/r/typescript/comments/hdmea1/newbie_are_getters_syntatic_sugar/
---
Instead of 

&amp;#x200B;

    constructor(public x: string){
    
    }
    get x () {
       
      return x;
    }

Why not just do 

&amp;#x200B;

    constructor(public x: string){
    
    }
    
    console.log(class.x)
    
    or even
    
    y = "y";
    constructor(public x: string){
    
    }
    console.log(class.y);
## [3][Generic constraints in constructor not working](https://www.reddit.com/r/typescript/comments/hdrdv7/generic_constraints_in_constructor_not_working/)
- url: https://www.reddit.com/r/typescript/comments/hdrdv7/generic_constraints_in_constructor_not_working/
---
[playground link](https://www.staging-typescript.org/play?#code/KYDwDg9gTgLgBASwHY2FAZgQwMbDgCUwGcALTAIwBs8BvAKDjjNIGEIATYACgEoAuOEgCuAW3Jo6AXzp1QkWHGyViROAFEUUAJ4AeANIAaOADUAfHHqNmJAcLESGcANbAtAvY4BumSkOADjGUZsCCQiGCghbBhoLhc3OD04UFQkdlVCUgpqI29ffxMeC0dGGBIEIgA6azgAXmdXauISNk5eErgyisr4uoatDq6qvL8+keBHaWkgA)

Trying to create an `Entry` class to be used in a hash map. I need to constrain a generic type in a constructor, but Typescript is giving me a weird type error.

Parameter 'key' of constructor from exported class has or is using private name ''

I'm exporting `Hashable` though.

If you replace `Hashable` with string: `key: K extends string` it still gives the same type error.

Why?
## [4][Can I annotate all TypeORM queries with a filter condition?](https://www.reddit.com/r/typescript/comments/hdph92/can_i_annotate_all_typeorm_queries_with_a_filter/)
- url: https://www.reddit.com/r/typescript/comments/hdph92/can_i_annotate_all_typeorm_queries_with_a_filter/
---
I'm trying to create a multi-tenant architecture. For this, I want to create a new EntityManager  
or Repository that automatically annotates all queries with an additional filter for the tenant\_id.

A query like:

    createQueryBuilder() 
        .select("user") 
        .from(User, "user") 
        .where("user.id = :id", { id: "[...]" })

Should be automatically transformed by the Repository or Manager to a query like:

    createQueryBuilder() 
        .select("user") .from(User, "user") 
        .where("user.id = :id", { id: "[...]" }) 
        .andWhere("user.tenantId = :tenantId", { tenantId: "[...]" });

Is this possible with TypeORM?

(The benefit with this approach is that it is more fault-tolerant since it would be impossible to forget filtering for the right tenant)
## [5][Cannot find type definition file for 'babel__generator'.](https://www.reddit.com/r/typescript/comments/hdpeag/cannot_find_type_definition_file_for_babel/)
- url: https://www.reddit.com/r/typescript/comments/hdpeag/cannot_find_type_definition_file_for_babel/
---
I'm struggling with this error and I have tried `yarn add @types/babel__generator`  it didn't work? How should deal with this error?
## [6][[Help] Getting property names of types](https://www.reddit.com/r/typescript/comments/hdpc2p/help_getting_property_names_of_types/)
- url: https://www.reddit.com/r/typescript/comments/hdpc2p/help_getting_property_names_of_types/
---
Hi there. I just wonder if there a possible way to get a property names from a type, not an object. I tried do somethng like this:

    export type Names&lt;T&gt; = {readonly [K in keyof T]: K};
    
    export interface Test {
      myLuck2: number;
      myBad: string;
    };
    
    export function nameof&lt;TClass&gt;(callback: (obj: Names&lt;TClass&gt;) =&gt; keyof TClass): keyof TClass {
      const names: Names&lt;TClass&gt; = {
        // Here I can't define same properties as in Names&lt;TClass&gt; dynamicly =(
      };
      return callback(names);
    };
    
    export const TestNameOfSomeProperty = nameof&lt;Test&gt;(x =&gt; x.myLuck2);

And, of course, I can't convert string type literal to value string... AFAIK. Maybe there is an issue or pull request?
## [7][How can improve my tsconfig?](https://www.reddit.com/r/typescript/comments/hd176e/how_can_improve_my_tsconfig/)
- url: https://www.reddit.com/r/typescript/comments/hd176e/how_can_improve_my_tsconfig/
---
Hi everybody. I have this repo in typescript and I would like to know how it can be improved. Looking for feedback. This is a package for managing permissions. [Link](https://github.com/roggervalf/iam-policies)
## [8][Having a real tough time "optionally" using typescript in my react app. see configs](https://www.reddit.com/r/typescript/comments/hd30z4/having_a_real_tough_time_optionally_using/)
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
## [9][How would you implement a Map with objects (vectors) as keys?](https://www.reddit.com/r/typescript/comments/hcrtdx/how_would_you_implement_a_map_with_objects/)
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
## [10][Noobie trying to do something advanced I guess - conditional arguments](https://www.reddit.com/r/typescript/comments/hcpm4d/noobie_trying_to_do_something_advanced_i_guess/)
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
## [11][Help with handling JSON array of objects (merge two arrays)](https://www.reddit.com/r/typescript/comments/hc5odl/help_with_handling_json_array_of_objects_merge/)
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
