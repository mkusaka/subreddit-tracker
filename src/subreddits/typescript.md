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
## [2][Any reason not to use ts-node for local utility scripts?](https://www.reddit.com/r/typescript/comments/ju90sy/any_reason_not_to_use_tsnode_for_local_utility/)
- url: https://www.reddit.com/r/typescript/comments/ju90sy/any_reason_not_to_use_tsnode_for_local_utility/
---
A service that is supposed to deliver a list of data somewhat lazily crunches an important field into an array of results. I'm writing a utility script to more properly create multiple rows with one property per row, for that particular field.

For little scripts like this do you guys feel content to use ts-node to skip the compile step or do you prefer compiling for everything? I lean towards ts-node being fine but wanted to ask.
## [3][Is there a way to type a function so that it returns different types of results dynamically according to it params?](https://www.reddit.com/r/typescript/comments/juegyk/is_there_a_way_to_type_a_function_so_that_it/)
- url: https://www.reddit.com/r/typescript/comments/juegyk/is_there_a_way_to_type_a_function_so_that_it/
---
This is a very contrived example so please bear with me:

I have a function `foo`. It takes an object of this type as its parameters 
```ts
interface Foo {
  bar: number,
  baz: number,
  enabled?: boolean,
}
```
When `enabled` is not provided by the user or when it is false, `foo` would return an object of type `Result`
```ts
interface Result {
  result: number
}
```
so it just adds these two numbers as in
```ts

function foo({bar, baz}: Foo): Result {
  return {
    result: bar + baz
  }
}

```

Here is the question: When `enabled` is provide in the params and it is set to `true`, I want the function returns another type of result, where it has another field `isEqual`, which indicates if `bar` and `baz` is equal. so like 
```ts
interface AnotherResult {
   result : number,
   isEqual: boolean
}
``` 

Here is how we can use this function
```ts
const {result} = foo({baz: 1, bar: 2})

// or

const {result, isEqual} = foo({baz: 1, bar: 2, enabled: true})
```

Ideally when the user include `enabled: true` in the params to `foo`, they can destructure  `isEqual` from the return value of `foo` and when they don't provide `enabled` or set `enabled: false`, they **cannot** destructure  `isEqual` from the return value of `foo`.

Below is my attempt. It doesn't fulfill the use case and I couldn't think of a way to dynamically determine the return value's type based on the params. I wonder if that is even possible with TypeScript

```ts
nterface Foo {
  bar: number,
  baz: number,
  enabled?: boolean,
}

interface Result {
  result: number
}

interface IsEqual {
  isEqual: boolean
}

type ResultWithIsEqualEnabled = Result &amp; IsEqual


function foo({bar, baz, enabled = false}: Foo): ResultWithIsEqualEnabled {
  return {
    result: bar + baz
  }
}

```
## [4][Is there a way to type imports?](https://www.reddit.com/r/typescript/comments/juhwkv/is_there_a_way_to_type_imports/)
- url: https://www.reddit.com/r/typescript/comments/juhwkv/is_there_a_way_to_type_imports/
---
Right now I'm doing this and it feels redundant. The import is from a .js file.

    import { theme } from '../path/to/themes/index.js';
    
    const Theme: Record&lt;string, any&gt; = theme;
    
    const App = () =&gt; {
        const themeName = useSelector(state =&gt; state.Theme.themeName)
    
        console.log(Theme[themeName].color.red);
    
        return null;
    }

if I don't declare `Theme` with `Record&lt;string, any&gt;`, it will give me an error

    Element implicitly has an 'any' type because expression of type 'string' can't be used to index type

It's not a big deal, but was wondering that extra declaration is mandatory.
## [5][Database connector](https://www.reddit.com/r/typescript/comments/juc9o9/database_connector/)
- url: https://www.reddit.com/r/typescript/comments/juc9o9/database_connector/
---
Hi, I'm about to start a project for a service that I'd like to optimise for heavy traffic and database queries, while still running in NodeJS (TypeScript + Fastify) as it is high performance and easier to find contributors.

I usually work on serverless environments so DynamoDB and Athena are my choices when it comes to persistent data. However, for this project I will need Postgresql as the project should not depend on AWS.

I am looking at sql builders or other libraries helping with sql. My preference would be to write statements myself but I don't want to write raw sqls because of sql injections.

I've looked into TypeORM, Klex and Sequelize.  I have to say I'm not impressed by any of them. The first is too magic, the second and the third are somewhat lacking in TypeScript support.

What do you use for sql in your backend projects?

Cheers!
## [6][Why do I get "Object is possibly 'undefined'"?](https://www.reddit.com/r/typescript/comments/jtznnp/why_do_i_get_object_is_possibly_undefined/)
- url: https://i.redd.it/40hy3jwyg6z51.png
---

## [7][Parsing JSON to typescript interface?](https://www.reddit.com/r/typescript/comments/ju5zlr/parsing_json_to_typescript_interface/)
- url: https://www.reddit.com/r/typescript/comments/ju5zlr/parsing_json_to_typescript_interface/
---
Which one is better? Mapping each individually, checking for undefined on optional properties and parsing it like:

Const user : User { ... }



or just do this:


Const user = JSON.parse(req.body.user) as User


0r am I doing this entirely wrong lol
## [8][Why do I got "Type 'string | number' is not assignable to type 'never'. Type 'string' is not assignable to type 'never'. " it's make me misleading](https://www.reddit.com/r/typescript/comments/ju34f0/why_do_i_got_type_string_number_is_not_assignable/)
- url: https://www.reddit.com/r/typescript/comments/ju34f0/why_do_i_got_type_string_number_is_not_assignable/
---
&amp;#x200B;

https://preview.redd.it/qkatakpqy7z51.png?width=832&amp;format=png&amp;auto=webp&amp;s=ae59555d1985b09c3405790e4bfd4c0891db1772

the object data only have two fields but typescript report this error is "data\[k\]" never?
## [9][Is it possible to recursively traverse the type of an object and match it against another type](https://www.reddit.com/r/typescript/comments/jttj2s/is_it_possible_to_recursively_traverse_the_type/)
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
## [10][Dev Experience: Is it possible to create a type that accepts any string but also provides autocompletion for specific strings?](https://www.reddit.com/r/typescript/comments/jtg7kt/dev_experience_is_it_possible_to_create_a_type/)
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
## [11][How to type function arguments when return type is not known ahead of time](https://www.reddit.com/r/typescript/comments/jtotih/how_to_type_function_arguments_when_return_type/)
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
