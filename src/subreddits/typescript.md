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
## [2][Assign nested Object as copy, not reference](https://www.reddit.com/r/typescript/comments/ftl7pi/assign_nested_object_as_copy_not_reference/)
- url: https://www.reddit.com/r/typescript/comments/ftl7pi/assign_nested_object_as_copy_not_reference/
---
Hi guys, i seem to not understand the assignment operator completly, when it comes to nested Objects.

I have initialization objects for my classes data, which is defined by an interface.

    export const MY_DATA_INIT {
        device: {
            id: 0,
            ip: "unknown"
        }
        anotherProperty: "unknown"
    }
    
    export interface MyData {
        device: Device;
        anotherProperty: string;
    }
    
    interface Device {
        id: number;
        ip: string;
    }

When i make an Object of type MyData and use MY\_DATA\_INIT to initialize it, i am only pointing to  MY\_DATA\_INIT

    let myDataObject: MyData = MY_DATA_INIT;
    
    myDataObject.device.ip = "192.168.0.2";    // This updates on MY_DATA_INIT
    myDataObject.anotherProperty = "value";    // This does'nt
    
    console.log(JSON.stringify(MY_DATA_INIT)).
    /*    {
     *        device: {
     *            id: 0,
     *            ip: "192.168.0.2",
     *        }
     *        anotherProperty: "unknown"
     *     }
     */

Using the spread operator {...MY\_DATA\_INIT} does not make a difference.

Same using the spread operator in the MY\_DATA\_INIT itself.

    export const MY_DATA_INIT {
        device: {...MY_DEVICE_INIT}
        anotherProperty: "unknown"
    }
    
    export const MY_DEVICE_INIT {
        id: 0,
        ip: "unknown"
    }

I don't get why this is failing. The direct Object properties of MyData seems to work as expected, but the nested device object is passed as reference.

Any ideas?

Edit: you're not seeing the class itself cause i stumbled upon this while writing Mocha unit tests for my DB access &amp; i only need to pass the data interface MyData.
## [3][Best way of figuring out the proper type to use?](https://www.reddit.com/r/typescript/comments/ftfqcf/best_way_of_figuring_out_the_proper_type_to_use/)
- url: https://www.reddit.com/r/typescript/comments/ftfqcf/best_way_of_figuring_out_the_proper_type_to_use/
---
TypeScript can be occasionally fussy about types, which is what I want, but at the same time it can be annoying, some things I've basically given up on figuring out and just put "any." Which kind of defeats the purpose of using it.   I'm learning React with TypeScript at the moment. Mostly it's okay, there are a few times that I've been defeated by React/TypesScript typings.  So where do you go for figuring out the proper types for variables, methods etc?
## [4][Prisma 2.0 is Now in Beta: Type-safe Database Access with Prisma Client](https://www.reddit.com/r/typescript/comments/ft21zn/prisma_20_is_now_in_beta_typesafe_database_access/)
- url: https://www.prisma.io/blog/prisma-2-beta-b7bcl0gd8d8e
---

## [5][Tetris clone built with TypeScript - https://tetris-js.now.sh/](https://www.reddit.com/r/typescript/comments/fszs1c/tetris_clone_built_with_typescript/)
- url: https://www.reddit.com/r/typescript/comments/fszs1c/tetris_clone_built_with_typescript/
---
I built a Tetris clone with TypeScript, not quite finished but there's enough meat there to play it. Expanding my OOP knowledge, coming from a functional programming background. Enjoy!

Game: [https://tetris-js.now.sh/](https://tetris-js.now.sh/)

Repo: [https://github.com/grantsru/tetris-js](https://github.com/grantsru/tetris-js)
## [6][Access Private/Protected via Utility Type?](https://www.reddit.com/r/typescript/comments/ft9tyy/access_privateprotected_via_utility_type/)
- url: https://www.reddit.com/r/typescript/comments/ft9tyy/access_privateprotected_via_utility_type/
---
I recently discovered [Utility Types](https://www.typescriptlang.org/docs/handbook/utility-types.html) and found one that allows modification of readonly properties via [TypeScript/issues/24509](https://github.com/microsoft/TypeScript/issues/24509).

```typescript
export type Mutable&lt;Type&gt; = {
  -readonly [key in keyof Type]: Type[key]
}

class Foo{
  readonly bar: boolean = true
}

const foo = new Foo()

;(foo as Mutable&lt;Foo&gt;).bar = false
```
 
This is very useful because it maintains IntelliSense &amp; refactoring capabilities in VSCode, and allows you to make assignments in your library code while still keeping the end-user's property access as read-only. 

Here are some examples demonstrating the differences of several implementations:

- no cast, has error:  
[@Imgur](https://i.imgur.com/f7I6phu.png)  
- cast to `any`, no error, but removes IntelliSense &amp; refactoring:  
[@Imgur](https://i.imgur.com/khysPQB.png)  
- cast to `Mutable`, no error, maintains IntelliSense &amp; refactoring:  
[@Imgur](https://i.imgur.com/LMutHUA.png)  
 
&amp;nbsp;  

-----
&amp;nbsp;  
 
Similarly, I'd like to be able to modify private and protected properties while maintaining IntelliSense &amp; refactoring.

Is it possible to write a utility type that allows this?
## [7][Beginner learning TypeScript? What are your favorite links to learn TS ASAP?](https://www.reddit.com/r/typescript/comments/ftjhc9/beginner_learning_typescript_what_are_your/)
- url: https://www.reddit.com/r/typescript/comments/ftjhc9/beginner_learning_typescript_what_are_your/
---

## [8][Webpack can't tree-shake class methods? How are you guys dealing with this?](https://www.reddit.com/r/typescript/comments/ft7wj6/webpack_cant_treeshake_class_methods_how_are_you/)
- url: https://www.reddit.com/r/typescript/comments/ft7wj6/webpack_cant_treeshake_class_methods_how_are_you/
---
I migrated over from Java so still have a bunch of Java conventions - one of which being I generally prefer OO programming.

The issue is that webpack can't tree-shake class methods because they can't be statically analyzed. 

It IS able to tree shake functions though.  

What's the best way to deal with this? Sometimes I use static classes as more like namespaces so I'm thinking of just using namespaces with functions instead.

This way they SEEM like static classes and tree-shaking can work.

Thoughts?
## [9][Help with a Styled-Components media query function](https://www.reddit.com/r/typescript/comments/ft7jhv/help_with_a_styledcomponents_media_query_function/)
- url: https://www.reddit.com/r/typescript/comments/ft7jhv/help_with_a_styledcomponents_media_query_function/
---
I'm following along with LevelupTutorials styled-components and most of it doesn't concern typescript vey much , but the lesson about media queries has me stumped about typing this function (mostly because I don't understand typing reduce functions:

`type SizesProps = {`  
`[key: string]: number`  
`}`

`const sizes: SizesProps = {`  
`small: 400,`  
`medium: 960,`  
`large: 1140,`  
`xlarge: 1792`  
`}`  
`const media = Object.keys(sizes).reduce((accumulator, label)=&gt; {`  
`accumulator[label] = (...args)=&gt; css\\`\``  
`u/media(min-width: ${sizes[label]}px) {`  
`${css(...args)}`  
`}`  
`\``  
`return accumulator`  
`}, {} )\``

The `accumulator[label]`  
 is telling me: Element implicitly has an 'any' type because expression of type 'string' can't be used to index type '{}'.  
No index signature with a parameter of type 'string' was found on type '{}'.

`${css(...args)}`  
is saying Expected at least 1 arguments, but got 0.ts(2555)

index.d.ts(277, 9): An argument for 'first' was not provided.

I've been trying to fiddle around with this but I just don't get it. Please help
## [10][Where should type declarations go for libraries with no type definition files?](https://www.reddit.com/r/typescript/comments/ft8elz/where_should_type_declarations_go_for_libraries/)
- url: https://www.reddit.com/r/typescript/comments/ft8elz/where_should_type_declarations_go_for_libraries/
---
...including nothing available on the atTypes repository? I recall having this issue a couple months ago. I created a global.d.ts file and put the needed type declaration in there. It worked.

But I remember when I changed the file name, it stoppered working even after ensuring that the declaration import was updated. It made me think there is a strict place (project root?) or naming convention for this to work, at least with default tsconfig.json settings.
## [11][Jetbrains vs Visual Studio Code with Typescript](https://www.reddit.com/r/typescript/comments/ft36lv/jetbrains_vs_visual_studio_code_with_typescript/)
- url: https://www.reddit.com/r/typescript/comments/ft36lv/jetbrains_vs_visual_studio_code_with_typescript/
---
Do you guys have any preferences when coding in Typescript? VSC is free I think and has autocomplete on types. Very useful in big libraries with non-ideal levels of distinction (I'm looking at you Firebase SDK) .I'm using Phpstorm, a Jetbrains IDE, and am considering jumping ship for those features.
