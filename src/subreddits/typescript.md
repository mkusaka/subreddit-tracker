# typescript
## [1][Who's hiring Typescript developers - August](https://www.reddit.com/r/typescript/comments/i1ikj5/whos_hiring_typescript_developers_august/)
- url: https://www.reddit.com/r/typescript/comments/i1ikj5/whos_hiring_typescript_developers_august/
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
## [2][Did you use typescript first and then JavaScript?](https://www.reddit.com/r/typescript/comments/i5pp1l/did_you_use_typescript_first_and_then_javascript/)
- url: https://www.reddit.com/r/typescript/comments/i5pp1l/did_you_use_typescript_first_and_then_javascript/
---
Curiously I had to do this for a university homework in Angular Ionic, without any JS experience it was quite difficult at the time, but since we were learning Java I liked the "types" aspect, now I'm learning more JS when I have time and I want to dedicate more time to TS again later.
## [3][Other languages with mapped/conditional types?](https://www.reddit.com/r/typescript/comments/i5vmni/other_languages_with_mappedconditional_types/)
- url: https://www.reddit.com/r/typescript/comments/i5vmni/other_languages_with_mappedconditional_types/
---
I’ve been using TS for a while now after only writing JS in the past. I’m currently working polymorphic data clients that have conditionally typed response types that would not be possible without conditional/mapped types.

This made me wonder, are there any other languages that use (structural) conditonal/mapped types? And if not, does anyone know why?

(I’ve glanced at C#, Swift, and Rust, and searched the web so perhaps I’ve glossed over some things)

Edit: I’ve added a more detailed example in the comments to explain what I’m after. Perhaps polymorphic is the wrong term here.
## [4][Array type of unknown length with first element type1 and the rest of the elements type2 .](https://www.reddit.com/r/typescript/comments/i5wgu2/array_type_of_unknown_length_with_first_element/)
- url: https://www.reddit.com/r/typescript/comments/i5wgu2/array_type_of_unknown_length_with_first_element/
---
How to define that type ?
## [5][Type parameter within a function call?](https://www.reddit.com/r/typescript/comments/i5vvwq/type_parameter_within_a_function_call/)
- url: https://www.reddit.com/r/typescript/comments/i5vvwq/type_parameter_within_a_function_call/
---
    import { useForm } from "react-hook-form";
    
    interface IFormInput {
      firstName: String;
      gender: GenderEnum;
    }
    
    export default function App() {
      const { register, handleSubmit } = useForm&lt;IFormInput&gt;();

I'm uncomfortable with this syntax. From the object destructure I can tell `useForm` returns an object that has keys for `register` and `handleSubmit`. So it is unclear what `IFormInput`  is doing here, if anything.

Any clarification is appreciated, thanks.

PS: Also, I vaguely recall reading that standard data typings (other than maybe `Function`)  should not be capitalized. Is the capitalized `String` above probably a minor mistake? [https://react-hook-form.com/get-started#Registerfields](https://react-hook-form.com/get-started#Registerfields)
## [6][Is it possible to access static methods on generic types?](https://www.reddit.com/r/typescript/comments/i5ytvz/is_it_possible_to_access_static_methods_on/)
- url: https://www.reddit.com/r/typescript/comments/i5ytvz/is_it_possible_to_access_static_methods_on/
---
I need to know if something like this is possible:

`class Database {`  
 `static newInstance(): Database {`  
 `return new Database();`  
`}`  
`}`  
`function create&lt;T extends Database&gt;(): T {`  
 `return T.newInstance();`  
`}`
## [7][typecasting clutter](https://www.reddit.com/r/typescript/comments/i5x5s7/typecasting_clutter/)
- url: https://www.reddit.com/r/typescript/comments/i5x5s7/typecasting_clutter/
---
    foo = &lt;string&gt;foo
    bar = bar as string

compiles to:

    foo = foo;
    bar = bar;

isn't the compiler supposed to discard these? am I missing something?
## [8][Managing Types across multiple front-ends and services](https://www.reddit.com/r/typescript/comments/i5rutg/managing_types_across_multiple_frontends_and/)
- url: https://www.reddit.com/r/typescript/comments/i5rutg/managing_types_across_multiple_frontends_and/
---
Hey all,

I’m curious if anyone has any good suggestions on ways to manage types across multiple frontends and services that have overlapping uses of types or use similar types. I’m trying to slowly move our org to typescript and would like a good way to manage types that won’t get super unruly.
## [9][Announcing TypeScript 4.0 RC](https://www.reddit.com/r/typescript/comments/i51era/announcing_typescript_40_rc/)
- url: https://devblogs.microsoft.com/typescript/announcing-typescript-4-0-rc/
---

## [10][Extracting interface with type parameters from classFactory that also has type parameters .](https://www.reddit.com/r/typescript/comments/i5eirf/extracting_interface_with_type_parameters_from/)
- url: https://www.reddit.com/r/typescript/comments/i5eirf/extracting_interface_with_type_parameters_from/
---
Suppose you have this :
```
function classFactory&lt;T1,T2,T3&gt;() {
	return class treeNode {
		prop1 : T1;
		prop2 : T2;
		prop3 : T3;
		/*some other prop that are not type functions*/
	};
}
```
How can I extract from that function this :
```
interface treeNode&lt;T1,T2,T3&gt; {
	prop1 : T1;
	prop2 : T2;
	prop3 : T3;
}
```
## [11][Using the new variadic tuple types to convert callback functions to promises](https://www.reddit.com/r/typescript/comments/i58b83/using_the_new_variadic_tuple_types_to_convert/)
- url: https://www.reddit.com/r/typescript/comments/i58b83/using_the_new_variadic_tuple_types_to_convert/
---
I've been messing around with the new variadic tuples in typescript 4.0 and was able to write a higher order function that converts the popular (err, result) "callback" format fn's into a type-safe promise.

Thought this was a really cool use-case of what's now possible, so wanted to share:

[Typescript Playground Link](https://www.typescriptlang.org/play?ts=4.0.0-beta#code/FAFwngDgpgBAwgQwDZIEYIMYGsCiAnPGAXhnzwHtCAfGAZxDwEsA7AcxhuYFcUBuUSLEQp02ADwAlAHzEYACigEAXPGRpMuAgBoYeKLR4gVEgJTEZCZmGADoMAGJdmGEI3LMA6oxAALYevEAFRgoAA8QKGYAE1oYSzAAbQBdHWlZOQA6LIQ8VloVBKyMwJ1VEQ1JKSSzIgsrfmAMd3oYCAoAW0ZaRgAzMH9RLFkxAEkQ8MiYuKtk1Kk5HqcMFUdnV3cvXwGKkbmamUzs3PyYEZMVAAUOrqhK8xgAb2AYXSgQLjxmGGYoAHcYK7kTq0KByOR6WjkJAANygOj0ACsoC59o9ni8YItnIcMjk8joFMoygFNHh4fpDMZUU8MRi9O9PiECDAAPyvJEuQl4MxKCFQ2HgilIEAmfi0gC+JnRkv44pswCxLjcXx8UBQ5E2fjUgzk6OYCHaUBU9CYbC06IQrCN3y47VQinNLwwqBUXJUZEoHBtKHJBmFxoYLFYqOh5EYUWAZhpMGdcm4PpgAAMRgBydowAAkD31hvFOktsCzBfFidFwDlwCQbxgqvVAEFaIDgbASG0gV1ev1tRo5LWkBrvFrytgy8A5AhaGBnPJqeiqyAYKw9G8g7IEL8EN4a2r+w2mzc5AAiPvkQ86ADMAEYyy9kIoQHIl1AV2wy5K5GWgA)

Also release it as an NPM package if anyone wants to help improve it (or fix my mistakes)!

[https://github.com/ChuckJonas/ts-promisify-callback](https://github.com/ChuckJonas/ts-promisify-callback)
