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
## [2][Good typescript function programming library](https://www.reddit.com/r/typescript/comments/gwvw9q/good_typescript_function_programming_library/)
- url: https://www.reddit.com/r/typescript/comments/gwvw9q/good_typescript_function_programming_library/
---
Hello, I'm finding a good functional programming library for typescript, currently I found [purify](https://gigobyte.github.io/purify) and [fp-ts](https://gcanti.github.io/fp-ts) and wondering which one I should learn or are there any other libraries that worth looking at. Oh and I'm just finding a library to learn functional programming, not working on any projects. Thank you.
## [3][@tsdotnet: A JavaScript-Friendly .NET Inspired TypeScript Library](https://www.reddit.com/r/typescript/comments/gwpu15/tsdotnet_a_javascriptfriendly_net_inspired/)
- url: https://www.reddit.com/r/typescript/comments/gwpu15/tsdotnet_a_javascriptfriendly_net_inspired/
---
[https://github.com/tsdotnet](https://github.com/tsdotnet)

Originally a port of `linq.js`, then evolved into a monolithic library to help those transitioning from C#/.NET to TypeScript.  Now updated with a clear intention to follow current / modern Node and Web development standards.

Classes are selectively grouped into individual packages to minimize unnecessary imports and improve tree-shaking.  Simply install any of the following using NPM.

# Highlights

Inspired by .NET:

* [@tsdotnet/linq](https://github.com/tsdotnet/linq): Written from the ground up to use JavaScript iterables and maximize modularity.
* [@tsdotnet/exceptions](https://github.com/tsdotnet/exceptions): Commonly used .NET style exceptions like `ArgumentNullException`, etc.
* [@tsdotnet/regex](https://github.com/tsdotnet/regex): Designed to mimic .NET `Regex` and enable named matching.
* [@tsdotnet/queue](https://github.com/tsdotnet/queue): Implementation of `Queue&lt;T&gt;` that performs much better than just an array.
* [@tsdotnet/linked-list](https://github.com/tsdotnet/linked-list):  A doubly (bidirectional) linked list.
* [@tsdotnet/string-builder](https://github.com/tsdotnet/string-builder): A simple implementation of a `StringBuilder`.
* [@tsdotnet/disposable](https://github.com/tsdotnet/disposable): `DisposableBase` and other utilities for managing object life-cycles.
* [@todotnet/lazy](https://github.com/tsdotnet/lazy): A disposable `Lazy&lt;T&gt;` and `ResettableLazy&lt;T&gt;` for simplifying lazy-initialization.
* [@tsdotnet/date-time](https://github.com/tsdotnet/date-time): Extensive set of date and time classes including `DateTime` and `TimeSpan`.
* [@tsdotnet/stopwatch](https://github.com/tsdotnet/stopwatch): A `Stopwatch` class with lap timing and easy to use static methods.
* [@tsdotnet/uri](https://github.com/tsdotnet/uri): A useful set of utilities for building URLs and managing query string values.

Other:

* [https://github.com/tsdotnet/linked-node-list](https://github.com/tsdotnet/linked-node-list)
* [https://github.com/tsdotnet/object-pool](https://github.com/tsdotnet/object-pool)
* [https://github.com/tsdotnet/ordered-registry](https://github.com/tsdotnet/ordered-registry)
* [https://github.com/tsdotnet/event-factory](https://github.com/tsdotnet/event-factory)
* [https://github.com/tsdotnet/tween-factory](https://github.com/tsdotnet/tween-factory)

And more... [https://github.com/tsdotnet](https://github.com/tsdotnet)

[tsdotnet](https://preview.redd.it/b7bstd0ley251.png?width=200&amp;format=png&amp;auto=webp&amp;s=8fe37df47a3c6707c5ba1f2d3685d5a041ec9ced)
## [4][Is it possible to get an interface key by value?](https://www.reddit.com/r/typescript/comments/gx1452/is_it_possible_to_get_an_interface_key_by_value/)
- url: https://www.reddit.com/r/typescript/comments/gx1452/is_it_possible_to_get_an_interface_key_by_value/
---
Hi everyone,

I'm trying to write a type (or interface) that receives an `HTMLElement` and returns its tag. For example:

```typescript
const divtag: Foo&lt;HTMLDivElement&gt; // "div"
```

We have dom's built in `HTMLElementTagNameMap` that does the opposite, i.e.:

```typescript
const divelem: HTMLElementTagNameMap["div"] // HTMLDivElement
```

How would you go about it?

Thanks!
## [5][Where do you keep your types?](https://www.reddit.com/r/typescript/comments/gwni63/where_do_you_keep_your_types/)
- url: https://www.reddit.com/r/typescript/comments/gwni63/where_do_you_keep_your_types/
---
We have just merged our client and server projects so we can easily share types across the two projects. This means we have \`client\`, \`server\` &amp; \`types\` dir. 

However, now the discussion has begun as to wether all types should live in the top level \`types\` directory or they should live to the closest place where they are used, so if in a single file - then in that file, in a module, in that module, in both projects, only then in types....

Where do people keep their types and what are the pros/cons of the approach?
## [6][Help with generic function call signatures](https://www.reddit.com/r/typescript/comments/gx2ft7/help_with_generic_function_call_signatures/)
- url: https://www.reddit.com/r/typescript/comments/gx2ft7/help_with_generic_function_call_signatures/
---
The goal of the below method is to continue looping until the callback returns a boolean. If/when this happens, the conditional should activate switching off the loop and assigning `userData` to the `rawUserInput` property in `callbackArgs`.

I don't fully understand how to handle typing a call signature to a generic method. I'm also not confident I typed `callbackArgs` correctly as I'm new to generics in general. Any feedback is appreciated, thanks.

    export default class Utilities {
       public loopUserInputDataUntilValid &lt;Callback, Args&gt;(
          callback: Callback,
          callbackArgs: Args extends {rawUserInput : string}) {
    
          let continueLooping = true;
    
          while (continueLooping) {
             const isValid: boolean = callback(callbackArgs);
             
             if (isValid) {
                continueLooping = false;
             }
          }
          // user input is validated
          
          const validatedInput = callbackArgs.rawUserInput;
          return validatedInput;
       }
    }
## [7][What filenames do you capitalize?](https://www.reddit.com/r/typescript/comments/gwykim/what_filenames_do_you_capitalize/)
- url: https://www.reddit.com/r/typescript/comments/gwykim/what_filenames_do_you_capitalize/
---
I capitalize filenames containing classes or interfaces. A couple years ago when I learned mongoose models, the course instructor capitalized his model files. I have stuck to that ever since. 

In TS I guess interfaces felt OOP enough that I leaked that personal convention over to them as well.

Curious what you guys do, and whether there is a more standard convention for which files should or should not be capitalized.
## [8][Is this even possible?](https://www.reddit.com/r/typescript/comments/gwzff8/is_this_even_possible/)
- url: https://www.reddit.com/r/typescript/comments/gwzff8/is_this_even_possible/
---
Check first:
```
class Animal {
  legs?: number
}

class Dog extends Animal {
  name: string
  barkingVolume?: number

  constructor(input: Dog) {
    super()
    Object.assign(this, input)
  }
}
```
1. Is there a way to get `Object.assign` type-safe without letting TS complain that `name` in line 6 has no initializer? 
2. Is there a way to abstract away the `Object.assign`/put it into the `Animal` class constructor?

I ask question 1 to get 2 working, then I could just eg. `new Cat({ tigerStripes: true })`, means I have subclasses with different constructor signatures but still having type-safety. This without  doing tiring many eg. `this.tigerStripes = input.tigerStripes` in every subclass constructor.

*Sandbox: https://codesandbox.io/s/clever-rosalind-4bqu2*
## [9][React's useMemo requires all dependencies to be of same type](https://www.reddit.com/r/typescript/comments/gwofi8/reacts_usememo_requires_all_dependencies_to_be_of/)
- url: https://www.reddit.com/r/typescript/comments/gwofi8/reacts_usememo_requires_all_dependencies_to_be_of/
---
I have a React component that has 3 props.  Their types are string, string, and MyInterface[].  If any of the props change, I want to recompute a new value.  I am trying to use useMemo:

```javascript
const computedValue = useMemo&lt;ComputedValueType&gt;(
  () =&gt; computeValue(props.string1, props.string2, props.myArray), [ 
    props.string1, 
    props.string2, 
    props.myArray
  ]);
```

I get the following error:

```
Type '(string | MyInterface[])[]' is not assignable to type 'ReadonlyArray&lt;any&gt;'.
  Types of property 'findIndex' are incompatible.
```

Is there a work around for this?  I am new to Typescript.  Would I have to compute my own memoization based on those three props and pass that in as the dependency?  Or at that point, just not use useMemo and do my own memoization?  Any help would be great.  Thanks.

UPDATE:
It looks like this issue is only an error that VS Code is showing.  Typescript from the command line works fine.  My VS Code and command line are both using Typescript 3.9.5.

UPDATE 2:
It goes away if I disable the VS Code plugin for JavaScript and TypeScript IntelliSense.
## [10][How can I import all interfaces from a file without a namespace?](https://www.reddit.com/r/typescript/comments/gwofcn/how_can_i_import_all_interfaces_from_a_file/)
- url: https://www.reddit.com/r/typescript/comments/gwofcn/how_can_i_import_all_interfaces_from_a_file/
---
I have a list of `interface`s I want to define in a file that I can use throughout the project. Ideally I'd like to be able to avoid needing to write `MyAPI.MyType` everywhere and just write `MyType`.

If there is a syntax that supports this—something like `import * from "module"`—then I haven't been able to find it or get it to work.

- How should I declare my interfaces? Top-level with `export` in front of each? Wrapped inside `declare module "x" { ... }`? etc
- What file type should my interfaces be in? `.ts` or `.d.ts`?
- And lastly, what would be the proper import statement? If this is at all possible
## [11][How would I properly write the example for React-Transition-Group's 'Transition' so that eslint stops yelling at me?](https://www.reddit.com/r/typescript/comments/gwgfzk/how_would_i_properly_write_the_example_for/)
- url: https://www.reddit.com/r/typescript/comments/gwgfzk/how_would_i_properly_write_the_example_for/
---
[ https://reactcommunity.org/react-transition-group/transition]( https://reactcommunity.org/react-transition-group/transition)

In particular this part:
    
    const defaultStyle = {
      transition: 'opacity ${duration}ms ease-in-out'
      opacity: 0,
    }
    
    const transitionStyles = {
      entering: { opacity: 1 },
      entered:  { opacity: 1 },
      exiting:  { opacity: 0 },
      exited:  { opacity: 0 },
    };
    
    const Fade = ({ in: inProp }) =&gt; (
      &lt;Transition in={inProp} timeout={duration}&gt;
        {state =&gt; (
          &lt;div style={{
            ...defaultStyle,
            ...transitionStyles[state]
          }}&gt;
            I'm a fade Transition!
          &lt;/div&gt;
        )}
      &lt;/Transition&gt;
    );                
    


I'm not really sure how to type the state in the arrow func or the transitionStyles[state] in the style
