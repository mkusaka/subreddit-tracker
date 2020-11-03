# typescript
## [1][Who's hiring Typescript developers - November](https://www.reddit.com/r/typescript/comments/jlsywo/whos_hiring_typescript_developers_november/)
- url: https://www.reddit.com/r/typescript/comments/jlsywo/whos_hiring_typescript_developers_november/
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
## [2][Strategy for hydrating/deserializing binary data into TypeScript objects?](https://www.reddit.com/r/typescript/comments/jmp44e/strategy_for_hydratingdeserializing_binary_data/)
- url: https://www.reddit.com/r/typescript/comments/jmp44e/strategy_for_hydratingdeserializing_binary_data/
---
I'm building a library which receives short binary messages from external hardware. (These happen to be MIDI messages, obtained via the Web MIDI API.) These messages range in size from three bytes to a few dozen.

There are many different types and subtypes of messages; these various types are structured hierarchically. For example:  


* On the most coarse level, they can be categorized as "channel messages" and "system messages". 
* Channel messages can be further categorized as "channel voice messages" and "channel mode messages".
* The "channel voice messages" category includes a number of concrete message types: "note on", "note off", "program change", etc.

So I have an analogous tree of TypeScript classes, all descended from the same `AbstractMidiMessage` base class.

**My goal:** Given the raw byte data for an individual message, I want to create an instance of the proper class. Can you recommend an approach? Are there standard patterns for this kind of deserialization?

In many cases, the message type can be inferred from the first byte of the message – but there are a number of subtler cases, where the first byte alone is insufficient to determine the message type.

The brute-force approach would be to write a giant if-elseif construct, to test for the various cases and instantiate the appropriate class.

But I'd prefer to avoid this approach. It isn't extensible. What if someone else wants to build a companion library, which provides additional, vendor-specific message types? They would have to patch my library to modify the if-elseif construct.

So I'm thinking that, instead, my library should delegate deserialization to the MidiMessage classes themselves.

Perhaps each class could implement canDeserialize() and deserialize() methods, like this: 

    class ExampleMidiMessage extends AbstractMidiMessage {
        
        public static canDeserialize(messageData: Uint8Array): boolean {
            // examine messageData, and return true if it looks like
            // an ExampleMidiMessage.
        }
        
        public static deserialize(messageData: Uint8Array): AbstractMidiMessage {
            const color = messageData[0];
            const flavor = messageData[1];
            const message = new ExampleMidiMessage(color, flavor);
            return message;
        }
        
    }

...and then I would call:

    const message = AbstractMidiMessage.deserialize(messageData);

...and then `AbstractMidiMessage` would delegate down its inheritance tree, until it found a concrete subclass which says that it can deserialize the data.

Of course, since TypeScript doesn't support introspection, I would need to provide a mechanism for subclasses to make parent classes aware of their existence.

At that point, this approach is starting to feel a little clunky. Perhaps it's good enough – but do you know of a better way? I've never really done this kind of thing before.

Thanks for reading!
## [3][How can I declare a const as a function that takes generics?](https://www.reddit.com/r/typescript/comments/jmch8b/how_can_i_declare_a_const_as_a_function_that/)
- url: https://www.reddit.com/r/typescript/comments/jmch8b/how_can_i_declare_a_const_as_a_function_that/
---
Given the following types, how can I declare a const `foo` of type `F` where generics of `F` must be given when using it (i.e. I want to call `foo` like this: `foo&lt;number&gt;(someValue)`)?

```
type P&lt;T&gt; = { a: T };
type R&lt;T&gt; = { b: T };
type F&lt;T&gt; = (param: P&lt;T&gt;) =&gt; R&lt;T&gt;;

// Define the dummy function here to reduce code dup.
const fn = (param: any) =&gt; ({ b: param.a });
```

Simply saying foo is of type `F` results in "Generic type 'F' requires 1 type argument(s). (2314)" so that's no good:

```
const foo: F = fn;
```

If I redefine `F` inline, it's fine:

```
const foo: &lt;T&gt;(param: P&lt;T&gt;) =&gt; R&lt;T&gt; = fn;
```

But it's a syntax error if I try declare some generic must be passed to `F`:

```
const foo: &lt;T&gt;F&lt;T&gt; = fn;
```

I think that last option is close, but I just need to figure out the right syntax. Is there a correct syntax? If not, does anyone know if there is an issue open for this (I couldn't find one)?
## [4][Continuations as collections](https://www.reddit.com/r/typescript/comments/jmaxe8/continuations_as_collections/)
- url: https://medium.com/@wim_jongeneel/continuations-as-collections-f4a5172a88a3?source=friends_link&amp;sk=1f57f96d6659aa11c0c5a33fac86db8d
---

## [5][Practice for employment algorithm tests in JS or TS?](https://www.reddit.com/r/typescript/comments/jm6rlq/practice_for_employment_algorithm_tests_in_js_or/)
- url: https://www.reddit.com/r/typescript/comments/jm6rlq/practice_for_employment_algorithm_tests_in_js_or/
---
I have an easier time keeping track of data shape in Typescript. I will take any job though.

Is Typescript normally offered as an option during an employment algorithm test? Or is its availability during those types of tests restricted, compared to Javascript?

If the later i may practice in Javascript.
## [6][Query regarding Excess Property Checks in TypeScript](https://www.reddit.com/r/typescript/comments/jlyvu3/query_regarding_excess_property_checks_in/)
- url: https://www.reddit.com/r/typescript/comments/jlyvu3/query_regarding_excess_property_checks_in/
---
Hi Everyone, 

I am unable to understand on why is `mySquare1` valid while `mySquare2` ends up in an error?

    interface SquareConfig {
      color: string;
      width: number;
    }

    function createSquare(config: SquareConfig): { color: string; area: number } {
      return { color: config.color || "red", area: config.width ? config.width*config.width : 20 };
    }

    let mySquare1 = { color:'red', width: 10, opacity: 15}
    let mySquare2 = createSquare(mySquare1);
    let mySquare = createSquare({ width: 100, color: 'white', opacity: 15 }); // Error
    // Argument of type '{ width: number; color: string; opacity: number; }' is not assignable to parameter of type 
    'SquareConfig'. Object literal may only specify known properties, and 'opacity' does not exist in type 
    'SquareConfig'.
## [7][Semicolons, yay or nay?](https://www.reddit.com/r/typescript/comments/jlsqxj/semicolons_yay_or_nay/)
- url: https://www.reddit.com/r/typescript/comments/jlsqxj/semicolons_yay_or_nay/
---
What is your preference for typescript?

[View Poll](https://www.reddit.com/poll/jlsqxj)
## [8][Strict version of built-in types (which use `any`)](https://www.reddit.com/r/typescript/comments/jlhzuv/strict_version_of_builtin_types_which_use_any/)
- url: https://www.reddit.com/r/typescript/comments/jlhzuv/strict_version_of_builtin_types_which_use_any/
---
Hello,

in Typescript, some function calls like

    const user = JSON.parse(something}
    
    console.log(user.metadata.name)

Or 

    fetch('/')
      .then(res =&gt; res.json())
      .then(user =&gt; { console.log(user.metadata.name) })

might result in runtime TypeError as they use \`any\` type for native JS and DOM APIs (instead of \`unknown\` as it didn't exist back then I suppose) .

  
(Links to definitions for the examples above: [https://github.com/microsoft/TypeScript/blob/71a2c59c557d871a1d38d10df83cfc36dc10d887/lib/lib.es5.d.ts#L1059](https://github.com/microsoft/TypeScript/blob/71a2c59c557d871a1d38d10df83cfc36dc10d887/lib/lib.es5.d.ts#L1059) and [https://github.com/microsoft/TypeScript/blob/71a2c59c557d871a1d38d10df83cfc36dc10d887/lib/lib.dom.d.ts#L2557](https://github.com/microsoft/TypeScript/blob/71a2c59c557d871a1d38d10df83cfc36dc10d887/lib/lib.dom.d.ts#L2557))

&amp;#x200B;

Is there a community version of these native types which would use \`unknown\` instead of  \`any\` for these kinds of functions?  


Or is there any other way to tackle this problem without requiring me to remember all the cases where typescript has got \`any\` as the built-in type?
## [9][Good practices to follow for a Typescript project?](https://www.reddit.com/r/typescript/comments/jlhuc4/good_practices_to_follow_for_a_typescript_project/)
- url: https://www.reddit.com/r/typescript/comments/jlhuc4/good_practices_to_follow_for_a_typescript_project/
---
Hi I'm going to start working on my first big project with Typescript using React Native. I'm still planning the whole system architecture but I need some help with how I should structure the code and files, document it, etc. Are there any resources that you would recommend for me?
## [10][Pathing packages in a TS monorepo?](https://www.reddit.com/r/typescript/comments/jl9ud4/pathing_packages_in_a_ts_monorepo/)
- url: https://www.reddit.com/r/typescript/comments/jl9ud4/pathing_packages_in_a_ts_monorepo/
---
Anyone have experience with pathing packages in monorepos?  I've got a monorepo with several packages, including a component library and a dashboard. I'm trying to import the component library in my dashboard package as import { Button } from '@company/ui' but it doesnt give me type definitions on my Button component unless I import from @company/ui/lib

I of course can always just import from /lib, but I feel like I'm doing something wrong here.
## [11][Justified use of an exclamation mark](https://www.reddit.com/r/typescript/comments/jl87jc/justified_use_of_an_exclamation_mark/)
- url: https://www.reddit.com/r/typescript/comments/jl87jc/justified_use_of_an_exclamation_mark/
---
Hey! I just found a justified use case for an exclamation mark:

    function parse(tokens: Token[]) {
        while (tokens.length) {
            const token = tokens.shift()!;
            ...
        }
    }

Since we check `tokens.length` as part of the while loop, we can be sure `tokens.shift()` returns a token.

So far, I haven't found many other cases of exclamation mark that I could justify. Have you found any that make sense?
