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
## [2][Why do I get "Object is possibly 'undefined'"?](https://www.reddit.com/r/typescript/comments/jtznnp/why_do_i_get_object_is_possibly_undefined/)
- url: https://i.redd.it/40hy3jwyg6z51.png
---

## [3][Is it possible to recursively traverse the type of an object and match it against another type](https://www.reddit.com/r/typescript/comments/jttj2s/is_it_possible_to_recursively_traverse_the_type/)
- url: https://www.reddit.com/r/typescript/comments/jttj2s/is_it_possible_to_recursively_traverse_the_type/
---
Guys. I have a tricky TS question. I tried to distill it as much as possible...
 
I have an object of this type
```ts
interface Obj {
  key1: SubType1;
  key2: SubType2;
}
```
Where `SubType1` and `SubType2` are all given. 
```ts
type Name = "LOL" | "HAHA";
type Statu = "Active" | "Inactive";
interface SubType1 {
  key: {
    name: Name;
  };
}
interface SubType2 {
  status: Statu;
}
```

I implemented a function which takes the object and other object called `transformObject` as its argument. The transform object has a recursive generic type so it knows the structure of the object provided. The point of this function is to conveniently traverse the object with `transformObject` and change a specific field of that object while conforming to the original type.

Here is the implementation
```ts
type MappedTransform&lt;T&gt; = {
  [K in keyof T]?: MappedTransform&lt;T[K]&gt; | ((params: T[K]) =&gt; T[K]);
};

function foo&lt;T&gt;(obj: T, transformObject: MappedTransform&lt;T&gt;) {
  // I omit the implementation details for this function to make the question simple
}

foo(obj, {
  key1: {
    key: {
      name: "HAHA" // 👈🏻 here the value that `name` field can change to can only be 
    // either `HAHA` or `LOL` because `MappedTransform` enforces that it has to be of the original type
    }
  }
});

```

My question is, if I want to use another type e.g. `Name` , `Status`  to lock where `transformObject` can reach, in other words, ideally if I do this
```ts
foo&lt;Obj, Name&gt;(obj, transformObject);
```
the `transformObject` can only be at the level of `Name` type, not on any other types e.g. `Status`.

so this should give me error because `status` is not of type `Name`, its type is `Status`
```ts
foo&lt;Obj, Name&gt;(obj, {
  key2: {
    status: 'Inactive' // 🚨 this should give me error since it is not of the type `Name` provided in the generic
  }
});
```

https://codesandbox.io/s/ts-recursive-function-tmo11?file=/src/index.ts

This is a live demo you can play with
## [4][Dev Experience: Is it possible to create a type that accepts any string but also provides autocompletion for specific strings?](https://www.reddit.com/r/typescript/comments/jtg7kt/dev_experience_is_it_possible_to_create_a_type/)
- url: https://www.reddit.com/r/typescript/comments/jtg7kt/dev_experience_is_it_possible_to_create_a_type/
---
I have a data object that contains certain keys but it can also contain random keys, for the keys I know about I want to return their type, for not known keys I want to return `unknown`.

Like this:  


    interface IData {
        a: string[];
        b: boolean;
    }
    
    function lookup&lt;K extends string&gt;(key: K): K extends keyof IData ? IData[K] : unknown {
        //do lookup
    }
    
    let a = lookup("a"); //infers string[]
    let b = lookup("b"); //infers boolean
    let c = lookup("c"); //infers unknown

If the argument for the `lookup` function matches a key from `IData` it will set the return type accordingly, else it will return `unknown`.

This works, however I don't get any autocomplete in the IDE for the known keys in `IData`. But for improved DX I would like to while still being able to enter any string as an argument, is there a way to do this currently?
## [5][How to type function arguments when return type is not known ahead of time](https://www.reddit.com/r/typescript/comments/jtotih/how_to_type_function_arguments_when_return_type/)
- url: https://www.reddit.com/r/typescript/comments/jtotih/how_to_type_function_arguments_when_return_type/
---
`function AFunction&lt;T, L&gt;(`  
 `arg: L,`  
 `callback: (x: T) =&gt; void,`  
  `{`  
 `mapArg = () =&gt; {`  
 `return {} as T;`  
`}`  
  `}: { mapArg: (x: L) =&gt; T }`  
`) {`  
 `callback(mapArg(arg));`  
`}`  
`AFunction(`  
 `45,`  
 `incorrectlyTypedArg =&gt; {`  
 `type Incorrect = typeof incorrectlyTypedArg;`  
  `},`  
  `{`  
 `mapArg(y) {`  
 `return \`${y}\`;`  
`}`  
  `}`  
`);`  


[Typescript Playground Link](https://www.typescriptlang.org/play?#code/GYVwdgxgLglg9mABAQQGLmvMAeAKgGkQBkA+ACgChFEBDAJwHMAuY-KxCGgGy4CMaIAaxZkAHi1wBKRAF4SiAG5wYAEzbUA3u2oBbGgAdkjWYjLS5iLdWuI6AUygg6SDQF9aAZ0S4A3NsSu7K4sGoh6howi4sTm8rgBFNJWHNx8AoJk4UYMZPQMkpJ+gRRoGLAIlNQALACs6ogwkHB09tBcAJ647fp2Ktmy8snUUN12iACSTS120F09JiM9cMANU61QHXO92X42APR7E2szUFsNXsuI4IJgcADuSIt2QfVDYQbZZO1J-tT2js5EAADAAkGnariBu2sgWogUKFCAA) I'm pretty new to typescript and I can't figure out how to type this function. The first parameter is some argument that can optionally be transformed by the third parameter \`mapArg\` otherwise it just is an empty object. This intermediate value is then passed into callback as it's only parameter.
## [6][How do I make this button properly functional?](https://www.reddit.com/r/typescript/comments/jtjljg/how_do_i_make_this_button_properly_functional/)
- url: https://www.reddit.com/r/typescript/comments/jtjljg/how_do_i_make_this_button_properly_functional/
---
Hey guys, I want to change the visibility of a flexBox until a button is pressed. The problem is, that React gives me errors and errors and errors. Maybe you know what to do. ( Object is possibly 'null'.ts(2531) ) -- I dont found any posibility that work for me on StackOverflow so I'm trying it here. \[Im using tsx in react\]

&amp;#x200B;

`const Explore = () =&gt; {`

`const ProfilesOpen = useRef(null);`

`const ProfilesClose = useRef(null);`

`function OpenProfiles() {`

`OpenProfiles.current.visibility = "hidden";`

`CloseProfiles.current.visibility = "visible";`

`}`

`return (`

`&lt;&gt;`

`&lt;div *className*="ProfilesOpen" *ref*={OpenProfiles}&gt;&lt;img *src*={ProfilesOpen} *alt*="Open Profiles"/&gt;&lt;/div&gt;&lt;div *className*="ProfilesOpen" *ref*={CloseProfiles}&gt;&lt;img *src*={ProfilesClose} *alt*="Open Profiles"/&gt;&lt;/div&gt;`

`&lt;/&gt;`

`)`

`}`

`export default Explore;`
## [7][A required key appears after transpiling to JS](https://www.reddit.com/r/typescript/comments/jtg2nc/a_required_key_appears_after_transpiling_to_js/)
- url: https://www.reddit.com/r/typescript/comments/jtg2nc/a_required_key_appears_after_transpiling_to_js/
---
Hi all,

I've been breaking my head over a property (`css`) which is not on my TSX component but becomes required after transpiling it to JS

    import React from 'react'
    import { default as MuiTooltip, TooltipProps as MuiTooltipProps } from '@material-ui/core/Tooltip'

    export type TooltipProps = MuiTooltipProps
    export default React.forwardRef(function Tooltip(props: TooltipProps, ref): JSX.Element {
      return &lt;MuiTooltip ref={ref} {...props} /&gt;
    })


Here I recreated the issue with the TS and JS version of the same component:  https://codesandbox.io/s/naughty-lichterman-mtx9l?file=/src/App.tsx 

Anyone any idea where `css` comes from?
## [8][Why does this property not exist when I destructure?](https://www.reddit.com/r/typescript/comments/jtkocz/why_does_this_property_not_exist_when_i/)
- url: https://www.reddit.com/r/typescript/comments/jtkocz/why_does_this_property_not_exist_when_i/
---
I have a type defined like this:

    interface A {
        type: 'EMAIL'
        email: string
    }

    interface B {
        type: 'USER'
        id: string
        email?: string
    }

    type CombinedType = A | B

Then I have a function that takes this type as an argument and destructures it:

    function C({
        type,
        id,
        email
    }: CombinedType) {
        ...
    }

I assumed that `id` would have type `string | undefined` but instead I get the following error.

    Property 'id' does not exist on type 'CombinedType'

Can anyone help me understand this error?
## [9][Having difficulty typing a function that can filter on a specific level inside of an object while preserving its original structure](https://www.reddit.com/r/typescript/comments/jt9n7j/having_difficulty_typing_a_function_that_can/)
- url: https://www.reddit.com/r/typescript/comments/jt9n7j/having_difficulty_typing_a_function_that_can/
---
I have an object called `myCached` and its type is known as 
```ts
interface MyCache {
  a: {
    b: {
      c: Target[];
      d: string;
    };
  };
}
```
Where `Target` is the type of the item in one of its property on a specific level, this type `Target` is known too. Assume `Target` is like this

```ts
type Target = "toRemove1" | "toRemove2" | "toPreserve1" | "toPreserve2";
```
In the example here, `myCache` is 
```ts
const cache: MyCache = {
  a: {
    b: { 
      c: ['toRemove1', 'toPreserve1', 'toRemove2', 'toPreserve2'], 
      d: "irrelevant" 
    },
  }
}
``` 
And I have an array of items that I need to remove. They also have the type of `Target` as in 
```ts
const thingsToRemove: Target[] = ["toRemove1", "toRemove2"];
```

And I am trying to come up with a function that can traverse this object and filter out items on a specific level. The way I designed this function is that it takes a transformer object and it travers the object and it provides the function for filtering. I want to type it properly so the users of this function can reply on the auto complete provided by the TS compiler and type checking to make sure that this function lands correctly on the correct level, in this case it is the `Target` level.

This is my code

```ts

type MappedTransform&lt;T&gt; = {
  [K in keyof T]?: MappedTransform&lt;T[K]&gt; | ((params: T[K]) =&gt; T[K]);
};

type Entries&lt;T&gt; = { [K in keyof T]-?: [K, T[K]] }[keyof T];

function traverse&lt;R&gt;(cache: R, transformObject: MappedTransform&lt;R&gt;): R {
  return (Object.entries(transformObject) as Array&lt;
    Entries&lt;MappedTransform&lt;R&gt;&gt;
  &gt;).reduce(reduceTransformNode, cache);
}

const reduceTransformNode = &lt;R, K extends keyof R&gt;(
  cacheNode: R,
  [transformKey, transformValue]: [
    K,
    MappedTransform&lt;R[K]&gt; | ((params: R[K]) =&gt; R[K]) | undefined
  ]
): R =&gt; {
  const { [transformKey]: node } = cacheNode;

  if (typeof transformValue === "undefined") return cacheNode;

  const newCacheValue =
    typeof transformValue === "function"
      ? (transformValue as (params: R[K]) =&gt; R[K])(node)
      : traverse(transformValue as R[K], node);

  return {
    ...cacheNode,
    [transformKey]: newCacheValue
  };
};

const x = traverse(cache, {
  a: {
    b: {
      // 👇Need to use Type `Target` to make sure that the transformer function lands on the type of `Target` exactly
      c: (node) =&gt; node.filter((s) =&gt; !thingsToRemove.includes(s))
    }
  }
});
```

This works fine and I have been really close to what I wanted to achieve except that **I cannot seem to find a way to add the type of `Target` to make sure that the function the user provides to do the filtering actually is on the right level i.e. `node` is of the right type `Target`**

Here is the [live demo][1]. Can someone tell me how to achieve that last bit of type safety I am seeking for here?


  [1]: https://codesandbox.io/s/superhardts-116bf?file=/src/index.ts
## [10][[vscode] Is it possible (if so, how) to add additional documentation for builtins, vendor, etc.?](https://www.reddit.com/r/typescript/comments/jszxvp/vscode_is_it_possible_if_so_how_to_add_additional/)
- url: https://www.reddit.com/r/typescript/comments/jszxvp/vscode_is_it_possible_if_so_how_to_add_additional/
---
I don't know what this would be called, I don't think any other IDE I've used supports this.

There are cases where I'd like to add additional documentation to builtin/externals -- or maybe "personalised"? Like if I'm stupid and continue misusing an API (like "hey I know it's common sense to be able to use \`.start()\` and \`.stop()\` on this interface, once stopped it can not be started again.").

I doubt this would be a "language" feature, but is there maybe a plugin for vscode?

I could see cases for having the documentation part of the repo and out of repo. Obviously another solution would be add an interface (and note exceptions locally). This was a "hmm I wonder if this is someone else has thought of" post.
## [11][Weird behaviour on VS-Code](https://www.reddit.com/r/typescript/comments/jsseoh/weird_behaviour_on_vscode/)
- url: https://www.reddit.com/r/typescript/comments/jsseoh/weird_behaviour_on_vscode/
---
Hello everyone!

Since beginning of the week something doesnt seem right.  
I dont know if its the fault of TypeScript or VS-Code (or something else) but TS seems to be the problem here.

Since the beginning of this week small "beauty" problems are marked as warnings in VS-Code such as indentations or double quotes instead of single quotes.

The warning message always points to tslint and googling the warning messages also seem to point out tslint - but nobody from the project changed the tslint file since the beginning of the project 2 months ago.

some of these messages look like this :

 `space indentation expected (indent)tslint(1)`   
or  
 `Shadowed name: 'ClassName' (no-shadowed-variable)tslint(1)` 

these werent there last week and makes the code very ugly and chaotic to look at.

I would share some code but its literally on any TS file when I press the tab key or use double quotes or instanciate a Class and on and on

EDIT :
I asked the other one working on that project who is using VSCode and he says he doesn't have the same issue. That's makes it even weirder
