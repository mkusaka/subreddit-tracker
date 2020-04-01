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
## [2][Which is a better pattern: `declare global` or `.d.ts` file?](https://www.reddit.com/r/typescript/comments/fskubj/which_is_a_better_pattern_declare_global_or_dts/)
- url: https://www.reddit.com/r/typescript/comments/fskubj/which_is_a_better_pattern_declare_global_or_dts/
---
Both `declare global` and `.d.ts` files let you declare Typescript definitions globally.  But which one is a better pattern in general? 

`.d.ts` seems to make more sense since it is all in one place but I want to know what good cases there are to prefer `declare global` as well.
## [3][Tetris clone built with TypeScript - https://tetris-js.now.sh/](https://www.reddit.com/r/typescript/comments/fszs1c/tetris_clone_built_with_typescript/)
- url: https://www.reddit.com/r/typescript/comments/fszs1c/tetris_clone_built_with_typescript/
---
I built a Tetris clone with TypeScript, not quite finished but there's enough meat there to play it. Expanding my OOP knowledge, coming from a functional programming background. Enjoy!

Game: [https://tetris-js.now.sh/](https://tetris-js.now.sh/)

Repo: [https://github.com/grantsru/tetris-js](https://github.com/grantsru/tetris-js)
## [4][Are closures still useful?](https://www.reddit.com/r/typescript/comments/fsz0sc/are_closures_still_useful/)
- url: https://www.reddit.com/r/typescript/comments/fsz0sc/are_closures_still_useful/
---
Closures return a function that has access to private state. Typescript allows for private variables. 

Do you guys still use closures for anything? If so what are the still relevant use cases for them?
## [5][Since typescript compiles itself can it optimize functional staff?](https://www.reddit.com/r/typescript/comments/fsv79v/since_typescript_compiles_itself_can_it_optimize/)
- url: https://www.reddit.com/r/typescript/comments/fsv79v/since_typescript_compiles_itself_can_it_optimize/
---
Staff like arr.map().filter().map().reduce() - can theoretically be compiled to single for loop. Should it be?  TypeScript can even give users some superset of standart   library
## [6][Myzod: Schema Validation and Type Inference Package](https://www.reddit.com/r/typescript/comments/fsew7n/myzod_schema_validation_and_type_inference_package/)
- url: https://www.reddit.com/r/typescript/comments/fsew7n/myzod_schema_validation_and_type_inference_package/
---
I want to acknowledge [zod](https://www.npmjs.com/package/zod) for their work and inspiring me to write this package.

[Myzod](https://www.npmjs.com/package/myzod) is a joi like object validation library with built-in typescript support that can infer the types defined by your schemas. The purpose of myzod is to no longer have to maintain types and validation separately so that they will always be in sync.  


Here is a basic example of how the package is used:  

```
import myzod, { Infer } from 'myzod';

const personSchema = myzod.object({
  id: myzod.number(),
  name: myzod.string().pattern(/^[A-Z]/),
  age: myzod.number().min(0),
  birthdate: myzod.number().or(myzod.string()),
  employed: myzod.boolean(),
  friendIds: myzod.array(myzod.number()).nullable()
});

type Person = Infer&lt;typeof personSchema&gt;;
// same as
type Person = {
  id: number;
  name: string;
  age: number;
  birthdate: number | string;
  employed: boolean;
  friendIds: number[] | null;
};

const person = personSchema.parse({ ... }); // here person is statically inferred with the correct type (Person)
```

Use it if you find it can be useful for you :) 
All feedback welcome.

[Myzod](https://www.npmjs.com/package/myzod)
## [7][Deno 1.0 will be released on May 13](https://www.reddit.com/r/typescript/comments/fry3ct/deno_10_will_be_released_on_may_13/)
- url: https://twitter.com/deno_land/status/1244707313006624772
---

## [8][First project in TypeScript - https://github.com/vydimitrov/react-random-reveal](https://www.reddit.com/r/typescript/comments/fse2c2/first_project_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/fse2c2/first_project_in_typescript/
---
 Hey all,

Just want to share with you my first project written in TypeScript - [https://github.com/vydimitrov/react-random-reveal](https://github.com/vydimitrov/react-random-reveal). It is a small library that adds a little thrill before revealing the truth so to speak. The effect is achieved with a character animation that shows random characters before revealing the text you want to emphasize. 

I must say that I fell in love with TypeScript, the compiler helped me several times to find some issues with my code. Of course, there were a few times where the compiler and I did not agree .... but after some time it manage to convince me :)

Any feedback on the TypeScript part would be appreciated, I hope you also find it useful.
## [9][How to Use Plain Functions to Validate Complex Structures](https://www.reddit.com/r/typescript/comments/fsajao/how_to_use_plain_functions_to_validate_complex/)
- url: https://medium.com/@moshesimantov/how-to-use-plain-functions-to-validate-complex-structures-85af5c8d4b93
---

## [10][Combine multiple basic types?](https://www.reddit.com/r/typescript/comments/fscok7/combine_multiple_basic_types/)
- url: https://www.reddit.com/r/typescript/comments/fscok7/combine_multiple_basic_types/
---
Hello everyone,

let's say I have the following string:

&gt;'133962:22:last'

The colons are used as separators. The first two parts are always only digits, the last one is a known string value.

Now I'd like to define this as its own type, preferably by combining multiple subtypes, e.g. (pseudocode)

&gt;type knownString = 'first' | 'second' | 'last' 

&gt;type groupedType = number &amp; ':' &amp; number &amp; ':' &amp; knownString

is something like this at all possible?

Thanks!
## [11][The Omit Helper Type in TypeScript](https://www.reddit.com/r/typescript/comments/frz2qm/the_omit_helper_type_in_typescript/)
- url: https://mariusschulz.com/blog/the-omit-helper-type-in-typescript
---

