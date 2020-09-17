# typescript
## [1][Who's hiring Typescript developers - September](https://www.reddit.com/r/typescript/comments/ik9rft/whos_hiring_typescript_developers_september/)
- url: https://www.reddit.com/r/typescript/comments/ik9rft/whos_hiring_typescript_developers_september/
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
## [2][Tricking Typescript into Spellchecking](https://www.reddit.com/r/typescript/comments/itzquc/tricking_typescript_into_spellchecking/)
- url: https://github.com/kkuchta/TSpell
---

## [3][How are .ts files run directly?](https://www.reddit.com/r/typescript/comments/iuj8fp/how_are_ts_files_run_directly/)
- url: https://www.reddit.com/r/typescript/comments/iuj8fp/how_are_ts_files_run_directly/
---
I'm using ts-node-dev to run an express.js server and just curious how these types of libraries are able to run .ts files without compilation. Do they have their own interpreter that emulates the JS interpreter, but with full Typescript support?
## [4][Dependency Injection in TypeScript from scratch](https://www.reddit.com/r/typescript/comments/iucio0/dependency_injection_in_typescript_from_scratch/)
- url: https://medium.com/@Sentinelone_tech/dependency-injection-in-typescript-from-scratch-d1a4422043a0
---

## [5][Naming convention for plain objects vs classed objects?](https://www.reddit.com/r/typescript/comments/iu8pqf/naming_convention_for_plain_objects_vs_classed/)
- url: https://www.reddit.com/r/typescript/comments/iu8pqf/naming_convention_for_plain_objects_vs_classed/
---
Let's say that I have an interface representing the attributes of a Point {x: number; y: number;} and I also have a Point class which has these attributes (implements the interface) but also has methods such as shift, rotate, etc.  

Is there any standardized convention on how to name the attributes/properties interface vs. the class itself?  I use this pattern a lot and I've been inconsistent with it.  PointProps for the properties and Point for the class?  Point for the properties and PointEditor for the class?
## [6][I've been using Typescript (migrated after 15 years of Java) for about 1.5 years. Here's what I've learned and how my programming style has changed.](https://www.reddit.com/r/typescript/comments/itmcvp/ive_been_using_typescript_migrated_after_15_years/)
- url: https://www.reddit.com/r/typescript/comments/itmcvp/ive_been_using_typescript_migrated_after_15_years/
---
- I've found I really like inner functions.  I'll put functions within functions to group similar code.  In Java this wasn't possible (at least 8 ... haven't checked lately).  It's better to keep things close together especially if the function isn't reusable.

- I'm calling all interfaces as IFoo rather than just Foo. Interfaces are a bit different than in Java and are far more important.

- Almost all my code is migrating to namespaces and I'm migrating almost entirely from static functions and static classes.  Webpack can't tree shake them and I think this is just cleaner either way.

- I'm moving to a lot more functions, maps, streams, etc.  

- I use undefined a lot more than null.  

- Types and generics work pretty well but every once in a while I get stuck and they can kind of get confusing with edge cases.

Curious how it's changed your programming style.
## [7][Typescript For Express](https://www.reddit.com/r/typescript/comments/iu3tlm/typescript_for_express/)
- url: https://www.reddit.com/r/typescript/comments/iu3tlm/typescript_for_express/
---
For type checking, do most people mostly use normal javascript (such as interfaces) or use an express framework such as NestJS? It looks like NestJS makes it much more complicated.
## [8][HTML 5 Canvas, React Refs and TypeScript](https://www.reddit.com/r/typescript/comments/itucpy/html_5_canvas_react_refs_and_typescript/)
- url: https://hashnode.blainegarrett.com/html-5-canvas-react-refs-and-typescript-ckf4jju8r00eypos1gyisenyf
---

## [9][Tech Notes: Interfaces generally belong with users](https://www.reddit.com/r/typescript/comments/ityqpp/tech_notes_interfaces_generally_belong_with_users/)
- url: http://neugierig.org/software/blog/2019/11/interface-pattern.html
---

## [10][Do you use explicit return types? Why or why not?](https://www.reddit.com/r/typescript/comments/itxos7/do_you_use_explicit_return_types_why_or_why_not/)
- url: https://www.reddit.com/r/typescript/comments/itxos7/do_you_use_explicit_return_types_why_or_why_not/
---
I've always thought using inference capabilities was a better way of programming. In addition to being terser, it allows updating the function body to happen more easily. If somehow this update led to a breaking change elsewhere in your code, the compiler would immediately report it, after which you have the option to either update the function again to keep its return type as what it was originally or update the function's callers' usage of the return type. With explicit return types, you don't have this option anymore.

&amp;#x200B;

Proponents of explicit return types will argue that it's good for documentation, but I doubt any of us are writing TypeScript in an IDE that doesn't tell you the inferred return types with a simple mouse hover
## [11][Assign dynamic properties from constructor to class (typechecking)](https://www.reddit.com/r/typescript/comments/itcssj/assign_dynamic_properties_from_constructor_to/)
- url: https://www.reddit.com/r/typescript/comments/itcssj/assign_dynamic_properties_from_constructor_to/
---
In short, what I'd like to do:

```
class Person {

  name: string

  // dynamic properties from constructor param


  constructor(name: string, dynProps: any) {

    this.name = name

    Object.keys(dynProps).forEach(key =&gt; this[key] = dynProps[key]) 

 } 

}


const dynamicProperties = { age: 25 } 


const someone = new Person('Greg', dynamicProperties)


someone.name // 'Greg'

someone.age // it should aknowledge it as number type and return 25
```
