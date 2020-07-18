# typescript
## [1][Who's hiring Typescript developers - July](https://www.reddit.com/r/typescript/comments/hizg5z/whos_hiring_typescript_developers_july/)
- url: https://www.reddit.com/r/typescript/comments/hizg5z/whos_hiring_typescript_developers_july/
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
## [2][What is needed to get source maps working?](https://www.reddit.com/r/typescript/comments/htg5jx/what_is_needed_to_get_source_maps_working/)
- url: https://www.reddit.com/r/typescript/comments/htg5jx/what_is_needed_to_get_source_maps_working/
---
I''m using NDB which is like the google chrome debugger but for Node.js apps. I'm able to set breakpoints in compiled .js files but not source files in .ts. I see this notice in the app:

[ndb source map notice](https://preview.redd.it/czi5qwyx4mb51.jpg?width=1912&amp;format=pjpg&amp;auto=webp&amp;s=a43c92123fb5807a5d001bfa85e5a9c347f9d99d)

I don't see any additional config options in the settings menu. Anyone know what step I'm likely missing for my .ts file breakpoints to be ignored here? 

I am running the bottom most script in the bottom left panel, `node compiled/index.js -b`  I want it to map to my ts files in the `src` folder so I can figure out where the problems are.

    // tsconfig.json for this project
    {
       "compilerOptions": {
           "target" : "ES5",
           "module": "commonjs",
           "noImplicitAny": true,
           "removeComments": true,
           "sourceMap": true
           , "outDir": "compiled/"
           , "esModuleInterop": true
           , "strict": true
           , "typeRoots": ["./src/types", "node_modules/@types"]
       },
       "include": ["src/**/*"]
       , "exclude": ["node_modules", "compiled", "__tests__", "types"]
    }
## [3][How to dynamically assign a property](https://www.reddit.com/r/typescript/comments/htfsm7/how_to_dynamically_assign_a_property/)
- url: https://www.reddit.com/r/typescript/comments/htfsm7/how_to_dynamically_assign_a_property/
---
So I'm working on a React/Typescript project, and I am trying to make this piece of code work:

    const inputChange = (event: ChangeEvent&lt;HTMLInputElement&gt;) =&gt; {
            setState((draft) =&gt; {
                draft.client[event.target.name] = event.target.value;
            });
        };

First, I'm using Immer and useImmer for my state setting, so the above code is safe and properly modifying immutable state.

Second, my main problem is in the square brackets. This function receives a ChangeEvent from an input element. The "name" property on that element matches one of the keys on the "client" object. I want to be able to assign the input's value to that property using this dynamic syntax.

The problem is TypeScript errors out because of issues trying to figure out the types here. I'm really not sure how to handle this properly. Help would be appreciated. Here is the error:

    Element implicitly has an 'any' type because expression of type 'string' can't be used to index type '{ accessTokenTimeoutSecs?: number | undefined; allowAuthCode?: boolean | undefined; allowClientCredentials?: boolean | undefined; allowPassword?: boolean | undefined; clientKey?: string | undefined; ... 4 more ...; refreshTokenTimeoutSecs?: number | undefined; }'.
      No index signature with a parameter of type 'string' was found on type '{ accessTokenTimeoutSecs?: number | undefined; allowAuthCode?: boolean | undefined; allowClientCredentials?: boolean | undefined; allowPassword?: boolean | undefined; clientKey?: string | undefined; ... 4 more ...; refreshTokenTimeoutSecs?: number | undefined; }'

Edit: More investigation, the problem seems twofold. First, knowing what property [event.target.name](https://event.target.name) that I'm passing into the square brackets matches. Second, ensuring that event.target.value, which has type 'string', is acceptable by that property. Again, help would be appreciated. Thanks.
## [4][Excess property check question](https://www.reddit.com/r/typescript/comments/htbn6t/excess_property_check_question/)
- url: https://www.reddit.com/r/typescript/comments/htbn6t/excess_property_check_question/
---
https://www.typescriptlang.org/docs/handbook/interfaces.html#excess-property-checks

It says "If an object literal has any properties that the “target type” doesn’t have, you’ll get an error:". The error is only because of the typo of color/colour. So by excess property checking does it mean it typescript checks for what it thinks is an error, i.e. a typo error? Because in the previous example when it was

    interface LabeledValue {
        label: string;
    }
    
    function printLabel(labeledObj: LabeledValue) {
        console.log(labeledObj.label);
    }
    
    let myObj = {size: 10, label: "Size 10 Object"};
    printLabel(myObj);

"it’s only the shape that matters. If the object we pass to the function meets the requirements listed, then it’s allowed.", so it's allowed here even though there's no size. So what I'm trying to figure out is what's the criteria for the typescript amigo to judge? Possible typos and?
## [5][Foal TS - Node.JS and TypeScript framework - July Release - Management of several environments &amp; Simplified authentication](https://www.reddit.com/r/typescript/comments/hsup8w/foal_ts_nodejs_and_typescript_framework_july/)
- url: https://www.reddit.com/r/typescript/comments/hsup8w/foal_ts_nodejs_and_typescript_framework_july/
---
Foal TS version 1.11 is here!

This version facilitates the management of several environments thanks to its abstract services and it also reduces the code to produce to build an authentication.

The documentation of the new features can be found here:

\- [https://foalts.gitbook.io/docs/topic-guides/architecture/services-and-dependency-injection#abstract-services](https://foalts.gitbook.io/docs/topic-guides/architecture/services-and-dependency-injection#abstract-services)

\- [https://foalts.gitbook.io/docs/topic-guides/authentication-and-access-control/session-tokens#specify-the-name-of-the-session-store-in-the-configuration](https://foalts.gitbook.io/docs/topic-guides/authentication-and-access-control/session-tokens#specify-the-name-of-the-session-store-in-the-configuration)

\- [https://foalts.gitbook.io/docs/topic-guides/validation-and-sanitization#usage-with-a-hook](https://foalts.gitbook.io/docs/topic-guides/validation-and-sanitization#usage-with-a-hook)

&amp;#x200B;

In a few words, Foal TS is a [**Node.Js**](http://node.js/) and **TypeScript** framework that **provides the bricks to build a complete web application** while keeping a **simple and intuitive code**.

Backed by **thousands of tests**, it offers more than **11,000 lines of documentation**. [**#typescript**](https://www.linkedin.com/feed/hashtag/?keywords=typescript&amp;highlightedUpdateUrns=urn%3Ali%3Aactivity%3A6689863897433829376) [**#javascript**](https://www.linkedin.com/feed/hashtag/?keywords=javascript&amp;highlightedUpdateUrns=urn%3Ali%3Aactivity%3A6689863897433829376) [**#nodejs**](https://www.linkedin.com/feed/hashtag/?keywords=nodejs&amp;highlightedUpdateUrns=urn%3Ali%3Aactivity%3A6689863897433829376)

&amp;#x200B;

https://preview.redd.it/8fpseb3rteb51.png?width=1064&amp;format=png&amp;auto=webp&amp;s=659f65321229a00a7ae887a53cc246f54709bcaf
## [6][Converting the few left at my work to TS](https://www.reddit.com/r/typescript/comments/hsr3ri/converting_the_few_left_at_my_work_to_ts/)
- url: https://www.reddit.com/r/typescript/comments/hsr3ri/converting_the_few_left_at_my_work_to_ts/
---
Every other week, folks in my group are able to host a 40 minute technical talk. I’m looking at hosting one in an attempt to convert some of the late adopters in the group to TS. Some background, we primarily work in a rapid prototyping space so the old “Typescript is slower to develop with” argument is what has kept the last few from adopting. I’m trying to attack that head on while also talking about other benefits to push my argument over the edge. I’m thinking about half of the talk will be demonstration and questions so I want to keep to the most obvious wins.

The rough outline of my presentation is this:

- Intellisense: by identifying and defining your types up front you earn intellisense which speeds up development 
- Runtimes errors/refactoring: Property changes and differences are caught at compile time rather than runtime.
- Coordination/documentation: TS improves the ability to communicate through though code. Often we return to prototypes a year after the initial work was done and have to spin up and develop on a code rot ridden app.

Any advice is very welcome. Specifically, does anyone have any additional points they think I should make or angles I should address?
## [7][Intersection of Enum Types](https://www.reddit.com/r/typescript/comments/hswsoo/intersection_of_enum_types/)
- url: https://www.reddit.com/r/typescript/comments/hswsoo/intersection_of_enum_types/
---
I'm trying to make an intersection of two enum types, but I have trouble making an instance of that new type. This is the smallest example I could think of.

    enum Element {
        WATER,
        FIRE
    }
    
    enum Target {
        PLAYER,
        ENEMY
    }
    
    type Attack = Element &amp; Target;

​Element.FIRE &amp; Target.ENEMY // doesn't work. treated like numbers  
{} as Attack // no error, but doesn't seem to have anything to do with element or target.

**Question: How would I make an Attack object?**

If that is not possible, I might just go with a tuple type that contains the two enums.
## [8][Typing a recursive camelize function?](https://www.reddit.com/r/typescript/comments/hstaxp/typing_a_recursive_camelize_function/)
- url: https://www.reddit.com/r/typescript/comments/hstaxp/typing_a_recursive_camelize_function/
---
It would be nice to not lose type of the object but IDK if this is far beyond what TypeScript is capable of:

    export function camelize&lt;T&gt;(source: T, reverse = false) {
      // console.log('snakeCaseObject', source);
      const dest: Record&lt;string, any&gt; = {};
      const fn = reverse ? snakeCase : camelCase;

      for (let [key, value] of Object.entries(source)) {
        if (isPlainObject(value)) {
          // checks that a value is a plain object or an array - for recursive key conversion
          value = camelize(value, reverse); // recursively update keys of any values that are also objects
        }
        if (isArray(value)) {
          value = value.map(v =&gt; (isPlainObject(v) ? camelize(v, reverse) : v));
        }
        dest[fn(key)] = value;
      }

      return dest;
    }


This is what I have so far, but surely the actual return type is T but with camelCased(or uncamelcased) keys.
## [9][Use mongoose FilterQuery&lt;T&gt; + check if the keys exists](https://www.reddit.com/r/typescript/comments/hsrstk/use_mongoose_filterqueryt_check_if_the_keys_exists/)
- url: https://www.reddit.com/r/typescript/comments/hsrstk/use_mongoose_filterqueryt_check_if_the_keys_exists/
---
Right now in our project (using NestJS) we do have some methods in our service like this:

```
import { Model, CreateQuery, FilterQuery } from 'mongoose';
import { User } from './user.schema';

class UsersService {
  constructor(
    @InjectModel(User.name) private readonly usersModel: Model&lt;User&gt;,
  ) {}

  async findOne(conditions: Partial&lt;Record&lt;keyof User, unknown&gt;&gt;) {
    this.usersModel.findOne(conditions as FilterQuery&lt;User&gt;);
  }

  async create(fields: Partial&lt;Record&lt;keyof User, unknown&gt;&gt;) {
    this.usersModel.create(fields as CreateQuery&lt;User&gt;);
  }

  async updateOne(id, fields: Partial&lt;Record&lt;keyof User, unknown&gt;&gt;) {
    this.usersModel.updateById(id, fields as UpdateQuery&lt;User&gt;);
  }
}
```

In the above code it makes sure through TypeScript that when I do:

```
this.usersService.create({
  name: 'John',
  countryy: 'USA', // &lt;-- error because it should be 'country'
})
```

This works perfectly, but advanced query syntax is not supported through this method. For example:

```
this.usersService.findOne({
  'myObject.nestedKey.nestedKey2': 'test',
})
```

due to this I decided to just change the `conditions` parameter in UsersService.findOne() function to be `Record&lt;string, unknown&gt;`.

Is there a possibility to allow this types of advanced syntax and still check whether those keys are valid? This means that TypeScript should know, when doing `myObject.nestedKey.nestedKey2`, that `myObject` is a valid key, `nestedKey` is also a valid key inside `myObject` etc ...

Mongoose is using `FilterQuery&lt;T&gt;` but it doesn't check for valid keys I believe

Would love to know some solutions.
## [10][Currying function parameters - Convert to Numbers](https://www.reddit.com/r/typescript/comments/hshhsn/currying_function_parameters_convert_to_numbers/)
- url: https://www.reddit.com/r/typescript/comments/hshhsn/currying_function_parameters_convert_to_numbers/
---
I have function like:

*export* const isGreaterThan = (operand1: unknown, operand2: unknown) =&gt;

Number(operand1) &gt;Number(operand2)

I would like to have a curried function in the middle that takes \`operand1\` and \`operand2\` and converts them into numbers so I don't have to do that explictly in my function body.

Is there a way to do this with function currying? Thank you
## [11][Is there a way to implement a custom behaviour when casting a class to a boolean?](https://www.reddit.com/r/typescript/comments/hsn96u/is_there_a_way_to_implement_a_custom_behaviour/)
- url: https://www.reddit.com/r/typescript/comments/hsn96u/is_there_a_way_to_implement_a_custom_behaviour/
---
If you have a class in TypeScript where whether it casts to true or false depends on a custom set of conditions, is there a way to implement that like you can implement something like `toString()`, where it will also work with regular if statements?
