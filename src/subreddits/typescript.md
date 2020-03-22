# typescript
## [1][Who's hiring Typescript developers - March](https://www.reddit.com/r/typescript/comments/fbll8c/whos_hiring_typescript_developers_march/)
- url: https://www.reddit.com/r/typescript/comments/fbll8c/whos_hiring_typescript_developers_march/
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
## [2][Confused by generic related compilation error.](https://www.reddit.com/r/typescript/comments/fmgw5y/confused_by_generic_related_compilation_error/)
- url: https://www.reddit.com/r/typescript/comments/fmgw5y/confused_by_generic_related_compilation_error/
---
Basically the code that I have that is leading to the compilation error is similar to this:

```
class Box&lt;T&gt; {

}

class Gen&lt;T extends Box&lt;number&gt;&gt; {
    private action(array: Array&lt;T&gt;): Array&lt;T&gt; {
        return new Array()
    }

    private useAction() {
        let x:Box&lt;number&gt;[] = new Array&lt;Box&lt;number&gt;&gt;();
        // This leads to the compilation error
        this.action(x)
    } 
}
```

I would expect this to work.

I defined `action` to take an array of any type, hence the `array:Array&lt;T&gt;`

In `useAction` I am calling the `action` method.

I am calling it with an array of `Box&lt;number&gt;`

Since I defined it to take any type, I expected passing a variable of type `Box&lt;number&gt;` to `action` should then work. 

But it fails with a compilation error:

&gt; Argument of type 'Box&lt;number&gt;[]' is not assignable to parameter of
&gt; type 'T[]'.   Type 'Box&lt;number&gt;' is not assignable to type 'T'.
&gt;     'Box&lt;number&gt;' is assignable to the constraint of type 'T', but 'T' could be instantiated with a different subtype of constraint
&gt; 'Box&lt;number&gt;'.

I can't decipher the error message. Why is it failing and how do I fix it?

I asked the question on StackOverflow, if you know how to crack this riddle, please feel free to answer the question over there

[https://stackoverflow.com/questions/60789275/how-to-fix-compilation-error-when-trying-to-call-a-generic-typescript-method](https://stackoverflow.com/questions/60789275/how-to-fix-compilation-error-when-trying-to-call-a-generic-typescript-method)
## [3][Proper TypeGraphQL architecture](https://www.reddit.com/r/typescript/comments/fmcymr/proper_typegraphql_architecture/)
- url: https://www.reddit.com/r/typescript/comments/fmcymr/proper_typegraphql_architecture/
---
Trying to figure out the reference setup with TypeGraphQL and my feeling is...

- `@ObjectType` + the decorated class reflect the TS Type + SDL Type
- `@Resolver` + the decorated class reflect the resolver
- `@Service` + the decorated class which is DIed with `typedi` should reflect the data source

Where I am unsure is the last one, `@Service`: 

1. Does it makes sense?
2. Do I put actual db queries (in my case mongo native driver) there or do I make another extra class?
3. Why do I have to decorate also the resolver with `@Service`?
4. How would 3 differ re ActiveRecord and DataMapper pattern?
## [4][Introducing Gretchen: Making Fetch Happen in TypeScript](https://www.reddit.com/r/typescript/comments/fm038x/introducing_gretchen_making_fetch_happen_in/)
- url: https://www.truework.com/blog/engineering/2020-03-04-introducing-gretchen/
---

## [5][Seafox: self-hosted javascript parser written in Typescript](https://www.reddit.com/r/typescript/comments/fm02aq/seafox_selfhosted_javascript_parser_written_in/)
- url: https://github.com/KFlash/seafox
---

## [6][PostSass - A Sass x PostCSS cli tool, written in TypeScript (X-Post r/webdev)](https://www.reddit.com/r/typescript/comments/fmdebi/postsass_a_sass_x_postcss_cli_tool_written_in/)
- url: https://github.com/Tanuel/postsass
---

## [7][How to create union of original type and custom type with some properties of original changed and rest the same?](https://www.reddit.com/r/typescript/comments/flxr64/how_to_create_union_of_original_type_and_custom/)
- url: https://www.reddit.com/r/typescript/comments/flxr64/how_to_create_union_of_original_type_and_custom/
---
I am trying to write typings for a function that is supposed to take an object and some of its keys as arguments and returns a new object where the value of specified keys can be the original or a Set/Array of original value type.

For example :
If we were to pass an object of the following type :

    interface Example {
        a: number
        b: number
        c: string
    }
and pass **a** and **c** as the keys, then the returned type would be

    {
        a: number | Set&lt;number&gt; | number[]
        b: number
        c: string | Set&lt;string&gt; | string[]
    } 


I have manages to do this by defining the following types

    type SetOrArray&lt;T, K extends (keyof T)[]&gt; = { [V in K[number]]: Set&lt;T[V]&gt; | T[V][] }
    type UnionType&lt;T, K extends (keyof T)[]&gt; = T | SetOrArray&lt;T, K&gt;



But when I try this with an object, then only the keys passed (in K) are preserved in the resulting UnionType and the remaining keys, which should have just been present with their original types due to **T |** union, are not existing on the UnionType.

For example:
      
    function afunc&lt;T, K extends (keyof T)[])&gt;(obj: T, ...keys: K): UnionType&lt;T, K&gt; {
        return obj // just for example
    }
    const obj = {a: 1, b: 2, c: "apple"}
    const res = afunc(obj, "a", "c")
    // res.a exists with type number | Set&lt;number&gt; | number[]
    // res.c exists with type string | Set&lt;string&gt; | string[]
    // but res.b does not exist, whereas it should have had type number
## [8][Bitwise expression computed property name in class?](https://www.reddit.com/r/typescript/comments/flk7j4/bitwise_expression_computed_property_name_in_class/)
- url: https://www.reddit.com/r/typescript/comments/flk7j4/bitwise_expression_computed_property_name_in_class/
---
I can't figure out how to use a bitwise expression as a computed property name in a TypeScript class (either through type literal, const assertion, or other means).

I'm converting legacy code to TypeScript, and this code uses integers that are powers of 2 as class identifiers. This enables using bitwise operations and computed property names to call unique functions for specific class/class interactions. For example, in JS:

    Dog.type = 2
    Cat.type = 4
    Interaction.prototype[Dog.type | Cat.type] = () =&gt; console.log('dog and cat')

Which is equivalent to:

    Interaction.prototype[6] = () =&gt; console.log('dog and cat')

It's perfectly valid JavaScript, but using a bitwise expression as a computed property name in TypeScript gives me an error:

    A computed property name in a class property declaration must refer to an expression whose type is a literal type or a 'unique symbol' type.

Both of the values are being used as a literal type - is there some way to tell the compiler that the resulting value is also a literal type?

EDIT: There is a workaround that requires reassignment and literal assertion:

    const test = (Dog.type | Cat.type) as 6
    Interaction.prototype[test] = () =&gt; console.log('dog and cat')

but the same workaround doesn't work with const assertion:

    const test = (Dog.type | Cat.type) as const

gives an error:

    A 'const' assertions can only be applied to references to enum members, or string, number, boolean, array, or object literals.

which is very inconvenient, because I'd need to remember to update the literal assertion if a class's type id ever changes.
## [9][Building Vue Enterprise Application: Part 1. Entities](https://www.reddit.com/r/typescript/comments/fljt1k/building_vue_enterprise_application_part_1/)
- url: https://medium.com/@gregsolo/building-vue-enterprise-application-part-1-entities-808077f3d2e7
---

## [10][The official TypeScript website's certificate expired today](https://www.reddit.com/r/typescript/comments/fl3ion/the_official_typescript_websites_certificate/)
- url: https://i.redd.it/n740bvlzyjn41.png
---

## [11][Inferring types with a key of a type object in typescript](https://www.reddit.com/r/typescript/comments/flkk9l/inferring_types_with_a_key_of_a_type_object_in/)
- url: https://www.reddit.com/r/typescript/comments/flkk9l/inferring_types_with_a_key_of_a_type_object_in/
---
So i have a type and I would like to dynamically update its fields based on a key value I pass into a function. However, the way I would like to do it (brokenFunction) is giving me an error on the assignment: Type '34' is not assignable to type 'never' and Type '34' is not assignable to type 'never'. workingFunction works but is not as handy. Does anyone have a way to make this work? I'm guessing it involves generics but I still dont understand why the type guards arent working.....

````
interface Todo {
  id: number;
  text: string;
}

const todo = {
  id: 1,
  text: "Buy milk",
};


function brokenFunction(x: Todo, field: keyof Todo): void {
  if (typeof field === 'number') {
    x[field] = 34
  } else if (typeof field === 'string') {
    x[field] = 'something else'
  }
}


function workingFunction(x: Todo, field: keyof Todo): void {
  if (field == 'id') {
    x[field] = 2345
  } else if (field == 'text') {
    x[field] = 'asdfsf'
  }
}
````
playground [link][1]


  [1]: https://www.typescriptlang.org/play/?ssl=28&amp;ssc=2&amp;pln=1&amp;pc=1#code/FASwdgLgpgTgZgQwMZQAQBUD2ATTqDewqqI2AXKmAK4C2ARrANxGrQAeEFAzhDOAObMAvsGBJMYHqxx4AvARakKARgA0Ldp1QAiAEJUAnqhogANgGtt6oc1FwqYJBBATUdGJnNQwAMQdOXMAAKNgosXFVUOBAoU3JULwNMOAwZAEoKADdMUgViEBSgiAMAByhkqJi41Fla1AByanpYerS84lQ2AG1o2OwAXRrUAGYAFhYhVFiuNALUItLylN7q2vl6nj4wflb24m6VgaGNzBooCAALASnTGfqJ4BE7f2dXAHdMGHMBP0dX4NCqQilT6FESFXCmAyqGyuUI%20UKhxq61Iu3hHQOVSO8gATGMAKwTG4zEiIrHIhqaNEsfY9LGDdYILjYOBcOD3YgiERAA
