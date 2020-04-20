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
## [2][Deonify: For NPM module authors that would like to support Deno but do not want to write and maintain a port.](https://www.reddit.com/r/typescript/comments/g4op75/deonify_for_npm_module_authors_that_would_like_to/)
- url: https://github.com/garronej/denoify
---

## [3][a TypeScript library that helps you build normalized state faster](https://www.reddit.com/r/typescript/comments/g4obby/a_typescript_library_that_helps_you_build/)
- url: https://www.reddit.com/r/typescript/comments/g4obby/a_typescript_library_that_helps_you_build/
---
If you've have ever worked with [normalized reducer state](https://redux.js.org/recipes/structuring-reducers/normalizing-state-shape/) (e.g. Redux or React's `useReducer`), you'll know that relational state updates involve boilerplate and [complexity](https://github.com/brietsparks/normalized-reducer#the-problem).

To save you time, try using [Normalized Reducer](https://github.com/brietsparks/normalized-reducer).

It is a zero-boilerplate higher-order-reducer that takes a schema of data relationships and returns the reducers, actions, and selectors to manage the normalized state. It [integrates with normalizr](https://github.com/brietsparks/normalized-reducer#normalizr-integration) and supports many common use-cases for relational data. To see its features in action, check out its [demo app](https://brietsparks.github.io/normalized-reducer-demo/).

Here is an example from the [doc's quick start](https://github.com/brietsparks/normalized-reducer#quick-start) to show its simplicity:

    import makeNormalizedSlice from 'normalized-reducer'
    
    const mySchema = {
      list: {
        'itemIds': { type: 'item', cardinality: 'many', reciprocal: 'listId' }
      },
      item: {
        'listId': { type: 'list', cardinality: 'one', reciprocal: 'itemIds' },
        'tagIds': { type: 'tag', cardinality: 'many', reciprocal: 'itemIds'}
      },
      tag: {
        'itemIds': { type: 'item', cardinality: 'many', reciprocal: 'tagIds' }
      }
    }
    
    const {
      reducer,
      actionCreators,
      actionTypes,
      selectors,
      emptyState,
    } = makeNormalizedSlice(mySchema)

Thanks for reading, and I hope this helps build things faster!
## [4][Why does this generic constructor not seem to get type checked at all? Can it be?](https://www.reddit.com/r/typescript/comments/g4rcdz/why_does_this_generic_constructor_not_seem_to_get/)
- url: https://www.reddit.com/r/typescript/comments/g4rcdz/why_does_this_generic_constructor_not_seem_to_get/
---
Hi I am trying to use a generic factory like class to create 1 type many times with the same passed in arguments. My issue is that the passed in arguments don't seem to be type checked against the generic type's constructor that is being created. I tried to boil it down to as simple as possible example:

&amp;#x200B;

    abstract class SharedBaseClass {
        constructor(public name: string) {}
        abstract sayHi(): void;
    }
    interface IConstructor&lt;T&gt;
    {
        new (...args: any[]): T;
    }
    class Creator&lt;InputType extends SharedBaseClass&gt; extends SharedBaseClass
    {
        inputs: Array&lt;InputType&gt;;
    
        constructor(private inputNames: string[],
                    inputConstructor: IConstructor&lt;InputType&gt;,
                    ...params: ConstructorParameters&lt;IConstructor&lt;InputType&gt;&gt;) {
            super("Creator");
            this.inputs = new Array&lt;InputType&gt;(this.inputNames.length);
            this.inputNames.forEach((name: string, idx: number) =&gt; {
                this.inputs[idx] = new inputConstructor(name, ...params);
            });
        }
        runAll() {
            this.inputs.forEach((input: InputType, idx: number) =&gt; {
                input.sayHi();
            });
        }
        sayHi() {
            console.log("Hi from Creator");
        }
    }
    
    class Input1 extends SharedBaseClass {
        constructor(name: string, public age: number, public otherArg: number) {
            super(name);
        }
        sayHi() {
            console.log("Hi from ", this.name, " ", this.age, " ", this.otherArg);
        }
    }

Allowed invocations:

     let c = new Creator&lt;Input1&gt;(["a", "b"], Input1); // no args
     let c = new Creator&lt;Input1&gt;(["a", "b"], Input1, 11); // too few args
     let c = new Creator&lt;Input1&gt;(["a", "b"], Input1, 11, 14); // correct # args
     let c = new Creator&lt;Input1&gt;(["a", "b"], Input1, 11, 14, 15); // too many args 
    
     c.runAll();

&amp;#x200B;

Is there anyway I can have the generic type's constructor arguments actually validated correctly here?

&amp;#x200B;

Thanks!
## [5][Decoupled State Interface](https://www.reddit.com/r/typescript/comments/g4nwag/decoupled_state_interface/)
- url: https://github.com/gactjs/store/blob/master/docs/decoupled-state-interface.md
---

## [6][Need some clarification on generics](https://www.reddit.com/r/typescript/comments/g4bivi/need_some_clarification_on_generics/)
- url: https://www.reddit.com/r/typescript/comments/g4bivi/need_some_clarification_on_generics/
---
Let's say that I have a function like this (taken from the handbook on the official website):

    function identity&lt;T&gt;(arg: T): T {
      return arg;
    }

What's going on here, exactly? I get that we don't want to use `arg: any` here because we lose all context on the actual type of the argument, so the above is a solution to that. I'm wondering what's going on here, exactly, though. Here's my guess; please tell me if I'm right or wrong. If I'm wrong, please correct me:

1. You pass in an argument into the `identity` function (let's say it's a string)
2. `&lt;T&gt;` *captures* the argument's type (i.e. string)
3. Thanks to `&lt;T&gt;` capturing the type (string), the function now knows that it should return a string

In other words, the order of operations here is:

1. A type is passed in through `(arg: T)`
2. The type is captured by `&lt;T&gt;` so that you can reuse that type context (i.e. string) throughout the function
3. The one (and only) instance, where we're reusing the type context captured by `&lt;T&gt;` in this function is when we're typing the return value of the function `:T`. In other words, the way we type the return value of the function here is an example of how we can utilize a generic `&lt;T&gt;` to improve our experience with typing.
## [7][Made a Sengled smart light api wrapper for nodeJS](https://www.reddit.com/r/typescript/comments/g4jnil/made_a_sengled_smart_light_api_wrapper_for_nodejs/)
- url: https://www.reddit.com/r/typescript/comments/g4jnil/made_a_sengled_smart_light_api_wrapper_for_nodejs/
---
So I made a wrapper for Sengled devices. You can control both rooms and devices individually. Setting brightness, power and viewing usage statistics. Hope this helps with your smart home adventures! [https://github.com/callmekory/sengled-api](https://github.com/callmekory/sengled-api)
## [8][Death of Component State](https://www.reddit.com/r/typescript/comments/g45jk8/death_of_component_state/)
- url: https://github.com/gactjs/store/blob/master/docs/death-of-component-state.md
---

## [9][10 less known Angular features you've probably never used - W3Radar](https://www.reddit.com/r/typescript/comments/g4gncq/10_less_known_angular_features_youve_probably/)
- url: https://blog.w3radar.com/less-known-angular-features-probably-never-used/
---

## [10][Fancy Emitter - A new take on JavaScript's EventEmitter class. Strongly Typed and makes use of the newest features offered in ES6+](https://www.reddit.com/r/typescript/comments/g3nhha/fancy_emitter_a_new_take_on_javascripts/)
- url: https://github.com/mothepro/fancy-emitter
---

## [11][Use Vim as a TypeScript IDE](https://www.reddit.com/r/typescript/comments/g3pb1a/use_vim_as_a_typescript_ide/)
- url: https://spacevim.org/use-vim-as-a-typescript-ide/#.Xpsb7HnFIw0.reddit
---

