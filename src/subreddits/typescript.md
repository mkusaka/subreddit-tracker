# typescript
## [1][Who's hiring Typescript developers - February](https://www.reddit.com/r/typescript/comments/ewxjh2/whos_hiring_typescript_developers_february/)
- url: https://www.reddit.com/r/typescript/comments/ewxjh2/whos_hiring_typescript_developers_february/
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
## [2][Search your codebase in natural language (Metacode, a vscode extension)](https://www.reddit.com/r/typescript/comments/f8vc6y/search_your_codebase_in_natural_language_metacode/)
- url: https://i.redd.it/hbys8ycwwwi41.gif
---

## [3][Why is reflect-metadata so important in the decorator pattern?](https://www.reddit.com/r/typescript/comments/f931o3/why_is_reflectmetadata_so_important_in_the/)
- url: https://www.reddit.com/r/typescript/comments/f931o3/why_is_reflectmetadata_so_important_in_the/
---
Relevant: https://www.reddit.com/r/typescript/comments/b8fphq/routing_with_typescript_decorators_for_node/

I'm looking for ways on `express` to make my controller life a little bit easier because the way I've been using it so far has been very verbose in nature. I've read that the decorator pattern solves that but I don't get what `reflect-metadata` has to do with it. Is it simply a convenience tool to get things done faster?

I've read the post and I've run the source code ([it's in GitHub](https://github.com/nehalist/ts-decorator-routing), by the way), but I still don't see why I shouldn't simply shove the route parameters into the constructor as properties without using `Reflect.getMetadata` / `Reflect.hasMetadata` / `Reflect.defineMetadata` to read / write / check. I've experimented the code to not use `reflect-metadata` and it still works.

Could someone explain this to me, preferably like I'm five?
## [4][Having difficulty in writing a types declaration file for a JS library](https://www.reddit.com/r/typescript/comments/f987n0/having_difficulty_in_writing_a_types_declaration/)
- url: https://www.reddit.com/r/typescript/comments/f987n0/having_difficulty_in_writing_a_types_declaration/
---
Its `index.js` file is as follows:

    module.exports = require("./src/foolibrary.js"); // class
    module.exports.bar = require("./src/subdir/bar.js"); // class
    module.exports.baz = require("./src/subdir/baz.js"); // class

And developers are supposed to use this library as follows:

    import Foo from 'foo-library';
    
    const bar = new Foo.bar('hello world');
    const foo = new Foo(bar);

So as you can see, the `bar` part (and `baz` as well) are attached to `foo`. I've made this work in the types declaration file by adding `bar` and `baz` as static methods of the main `foo` class, which each return their respective classes. However, my IDE is telling me `Only a void function can be called with the 'new' keyword.ts(2350)` on `new foo.bar('hello world')`. I need to have it so that the developer can instantiate `foo`, `bar` and `baz` with the `new` keyword but the way that the `index.js` file is written for this library kind of gets in the way of that since there is the unnamed 'default' export of `foo` itself and then `bar` and `baz` are tacked onto `module.exports` as non-default exported classes. My types file is currently as follows:

    declare module 'foo-library' {
      class Foo {
        constructor(bar: Bar);

        // work-around static methods
        static bar(path: string, options?: object): Bar;
        static baz(path: string, options?: object): Baz;
        ...
      }

      class Bar {
        constructor(path: string, options?: object);
      }

      class Baz {
        constructor(path: string, options?: object);
      }

      export = Foo;
    }

How do I declare the classes `Bar` and `Baz` as named exports with the default export of `Foo` as well?
## [5][Keeping Original Value When Transforming in RxJS](https://www.reddit.com/r/typescript/comments/f90mek/keeping_original_value_when_transforming_in_rxjs/)
- url: https://medium.com/@aerabi/keeping-original-value-when-transforming-in-rxjs-f4650e12c4cf?source=friends_link&amp;sk=7dff6d828ab33591fb8a9b91e11e90c0
---

## [6][Is it possible to generate a Tuple Type algorithmically (for recursive sub-keyof)](https://www.reddit.com/r/typescript/comments/f94lln/is_it_possible_to_generate_a_tuple_type/)
- url: https://www.reddit.com/r/typescript/comments/f94lln/is_it_possible_to_generate_a_tuple_type/
---
Right now, I have these overloads:

    interface Chainer&lt;Subject&gt; {
        its2&lt;K extends keyof Subject&gt;(propertyName: K, options?: Loggable): Chainable&lt;Subject[K]&gt;;
        its2&lt;K1 extends keyof Subject, K2 extends keyof Subject[K1]&gt;(key1: K1, key2: K2, options?: Loggable): Chainable&lt;Subject[K1][K2]&gt;;
        its2&lt;K1 extends keyof Subject, K2 extends keyof Subject[K1], K3 extends keyof Subject[K1][K2]&gt;(key1: K1, key2: K2, key3: K3, options?: Loggable): Chainable&lt;Subject[K1][K2][K3]&gt;;
    }

This works... but it's kinda hideous. And I have to hardcode an override for every possible number of parameters.

I'd like to do something along these lines (this obviously doesn't compile):

    type SubKeys&lt;T, K1 extends keyof T, K2 extends keyof T[K1], K3 extends keyof T[K1][K2], ...&gt; = [keyof T1, keyof T[K1], keyof T[K1][K2], ...]
    
    interface Chainer&lt;Subject&gt; {
        its2(...propertyNames: SubKeys&lt;Subject&gt;, options?: Loggable): Chainable&lt;Subject[K]&gt;;
    }

Is there any way to do this in TypeScript 3.8?

&amp;#x200B;

Assuming it's possible, an even better solution would be one that allows an optional number to follow each key. That is... the types match `propertyNames[0]: keyof Subject` and `propertyNames[i] : [keyof Subject[K1][K2]...[Ki], number?]` .

I'm imagining usage along the lines of `its2('key1', 'key2', ['key3', 5], 'key4')` , which (more or less) returns `somethingPrivate.key1.key2.key3[5].key4` and additionally gives compile errors if keys are not valid. (I say more or less because the return value is wrapped)
## [7][How do I check that a caught error matches a certain interface?](https://www.reddit.com/r/typescript/comments/f91zlt/how_do_i_check_that_a_caught_error_matches_a/)
- url: https://www.reddit.com/r/typescript/comments/f91zlt/how_do_i_check_that_a_caught_error_matches_a/
---
At first I tried `catch(e: Foo)` but realized that didn't work:

&gt; ts(1196): Catch clause variable cannot have a type annotation.

It makes sense I can't know the type of error that might be caught because it might be a system error.

So I assume I need a type guard, but I'm not sure how do this with an interface (not a value).

In my case, I'm using axios:

```
} catch (error) {
        if (error instanceof AxiosError) {
        if (typeof error == AxiosError) {
}
```
## [8][Is it possible to use the child's type in a parent class?](https://www.reddit.com/r/typescript/comments/f8zg6p/is_it_possible_to_use_the_childs_type_in_a_parent/)
- url: https://www.reddit.com/r/typescript/comments/f8zg6p/is_it_possible_to_use_the_childs_type_in_a_parent/
---
I want to be able to make a "createable" superclass that attaches a static "create" function on a class. This is to basically initialise a complex class with a "field initializer" or "object initialiser" style.

[https://github.com/microsoft/TypeScript/issues/16737](https://github.com/microsoft/TypeScript/issues/16737)

I've kinda come close I think, but I can't get the \`obj\` argument in create to be typed checked, it accepts anything. I want it to be the type of the child class bar any methods.

    type ExcludeMethods&lt;T&gt; =
        Pick&lt;T, { [K in keyof T]: T[K] extends (_: any) =&gt; any ? never : K }[keyof T]&gt;;
    
    export type StaticThis&lt;T&gt; = { new (): T };
    
    export class Createable {
        static create&lt;T&gt;(this: StaticThis&lt;T&gt;, obj: ExcludeMethods&lt;T&gt;): InstanceType&lt;T&gt; {
            return Object.assign(new this(), obj)
        }
    }
    
    export class Person extends Createable {
        constructor(public name: string) {
            super();
        }
    }
    const p = Person.create({name: 'ihsan'});

&amp;#x200B;

EDIT got it to work

    type DataPropertyNames&lt;T&gt; = {
        [K in keyof T]: T[K] extends Function ? never : K;
    }[keyof T];
    
    type DataPropertiesOnly&lt;T&gt; = {
        [P in DataPropertyNames&lt;T&gt;]: T[P] extends object ? DTO&lt;T[P]&gt; : T[P]
    };
    
    export type DTO&lt;T&gt; = DataPropertiesOnly&lt;T&gt;;
    
    
    export class Createable {
        static create&lt;T&gt;(this: {new (): T}, obj:DTO&lt;T&gt;): T {
            return Object.assign(new this(), obj)
        }
    }
    
    export class Person extends Createable {
        private name: string;
    
    
        logName(){
            console.log(this.name)
        }
    }
    
    const p = Person.create({name: 'ihsan'});
    p.logName(); // "ihsan"
## [9][Using type guards in TypeScript](https://www.reddit.com/r/typescript/comments/f8wwrg/using_type_guards_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/f8wwrg/using_type_guards_in_typescript/
---
I and a co-worker got into an argument some time ago about using "type guards" in TS i.e

```ts
function foo(bar: string): void {
  if (typeof bar !== 'string') throw Error('We don't do that here!');

  // more code...
}
```

We were building a library in TS to be used in a JS app.

I am of the opinion that the guard is needless since we were using TS.

He's of the opinion that it's necessary since the consumers of our library are using JS.

What's best practice? Thanks.
## [10][Flagged enum, why and how](https://www.reddit.com/r/typescript/comments/f8smod/flagged_enum_why_and_how/)
- url: https://timdeschryver.dev/blog/flagged-enum-what-why-and-how
---

## [11][I feel like it’s hard to learn typescript without first learning a framework line angular or react. Is there anyone out there who learned type script without first learning a framework? If so how and why?](https://www.reddit.com/r/typescript/comments/f94c4b/i_feel_like_its_hard_to_learn_typescript_without/)
- url: /r/LearnTypescript/comments/f947f9/i_feel_like_its_hard_to_learn_typescript_without/
---

