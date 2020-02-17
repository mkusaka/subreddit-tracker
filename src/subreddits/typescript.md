# typescript
## [1][Who's hiring Typescript developers - February](https://www.reddit.com/r/typescript/comments/ewxjh2/whos_hiring_typescript_developers_february/)
- url: https://www.reddit.com/r/typescript/comments/ewxjh2/whos_hiring_typescript_developers_february/
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
## [2][swc is a typescript / javascript compiler. It's 20x faster than babel.](https://www.reddit.com/r/typescript/comments/f53z29/swc_is_a_typescript_javascript_compiler_its_20x/)
- url: https://swc-project.github.io/
---

## [3][TypeScript Generics. Discussing naming conventions](https://www.reddit.com/r/typescript/comments/f58ox4/typescript_generics_discussing_naming_conventions/)
- url: https://wanago.io/2020/02/17/typescript-generics-discussing-naming-conventions/
---

## [4][Am I doing it right?](https://www.reddit.com/r/typescript/comments/f560e3/am_i_doing_it_right/)
- url: https://www.reddit.com/r/typescript/comments/f560e3/am_i_doing_it_right/
---
I'm working on a Firefox extension with TypeScript, and I have `webpack` set up with `ts-loader` to produce several `.js` files, each corresponding to its own entry point within the FF extension.

I'm using `webpack` in watch mode (with `yarn webpack-cli`), so the files get recompiled whenever there's a change.

The `webpack` config for each file has this part:

    module: {rules: [{use: 'ts-loader', exclude: /node_modules/,},],},
    resolve: {extensions: ['.tsx', '.ts'],},

My question is, am I doing something standard or am I over-complicating things? My suspicion that I'm over-complicating things is due to the fact that I had to install three different things (`webpack`, `ts-loader`, `webpack-cli`) to do just one thing. I know TypeScript itself can also produce a single transpiled `.js` file. Is that preferred? Or do both work? Thanks.
## [5][Typescript Generics - why doesn't this throw an error?](https://www.reddit.com/r/typescript/comments/f4v69j/typescript_generics_why_doesnt_this_throw_an_error/)
- url: https://www.reddit.com/r/typescript/comments/f4v69j/typescript_generics_why_doesnt_this_throw_an_error/
---
I'm playing around with typescript just on the command line and am confused as to why the following compiles without any type errors.  If I deliberately type my mapper to use the same generic for input and output but then make use of it with an input and output of *different* types, shouldn't it fail to compile? (Note: formatting with backticks got messed up, assume the makeString function is valid syntax)

`function map&lt;T&gt;(array: T[], f: (item: T) =&gt; T): T[] {`

  `let result = []`  
  `for (let i = 0; i &lt; array.length; i++) {`  
`result[i] = f(array[i])`  
`}`   
  `return result`

`}`

`const arrayOfInts = [1,2,3]`

`function makeString(item) {return \`${item}\`}\`

`let mapped = map(arrayOfInts, makeString);`

`console.log(mapped)`

It's only when I add typing to \`makeString\` indicating that it will accept one type T and return a string that this will fail to compile.
## [6][Frontend build tools and typescript](https://www.reddit.com/r/typescript/comments/f4rc5f/frontend_build_tools_and_typescript/)
- url: https://www.reddit.com/r/typescript/comments/f4rc5f/frontend_build_tools_and_typescript/
---
Hello! So figured i was going to try and learn typescript by doing a frontend project in react and typescript. As i was setting up dev dependencies it hit me that, as far as i understand, the typescript compiler both transpiles and bundles your code. My question is: what role does tools like Babel and Webpack play when you are developing frontend with typescript?
## [7][Why are namespaces regarded as deprecated?](https://www.reddit.com/r/typescript/comments/f4x6co/why_are_namespaces_regarded_as_deprecated/)
- url: https://www.reddit.com/r/typescript/comments/f4x6co/why_are_namespaces_regarded_as_deprecated/
---
When asked what features he regretted integrating into the language Anders Hejlsberg mentioned ``namespace`` invoking the fact that it was made deprecated by the ECMAScript module system.

This came as a shock to me as I am making extensive use of ``namespace``.
I do not see how ECMAScript module can replace namespace, using module imply splitting the code into different files.


Consider this example:
```typescript
export type Person = Person.Student | Person.Teacher;

export namespace Person {

    export type Common_ = {
        name: string;
        age: number;
    };

    export type Student = Common_ &amp; {
        type: "STUDENT";
        grades: {
            maths: number;
            compSci: number;
        };
    };

    export namespace Student {

        export const match = (person: Person): person is Student =&gt;
            person.type === "STUDENT";

        export const getAverageScore = (student: Student) =&gt;
            (student.grades.compSci + student.grades.maths) / 2;

    }

    export type Teacher = Common_ &amp; {
        type: "TEACHER";
        subject: "MATH" | "COMP-SCI";
    };

    export namespace Teacher {

        export const match = (person: Person): person is Teacher =&gt;
            person.type === "TEACHER";

    }

}

{

    const logPerson = (person: Person) =&gt;
        console.log([
            person.name,
            Person.Student.match(person) ?
            `is a student, averageScore is: ${Person.Student.getAverageScore(person)}` :
            `is a ${person.subject} teacher`
       ].join(" "))
       ;


    const teacher: Person.Teacher = {
        "type": "TEACHER",
        "age": 53,
        "name": "Bob",
        "subject": "MATH"
    };

    //Prints: "Bob is a MATH teacher"
    logPerson(teacher);

}
```

I see many advantages to structuring the code types definition this way.  
First I can just import { Person } and get all the type hierarchy.  

If I do not put ``Student`` and ``Teacher`` in the ``Person``'s namespace nothing tells me that they are a subset of ``Person``.  

When I need a function that I know to be tightly coupled with a type I can easily find it and it doesn't matter if I don't remember exactly its name, IntelliSense will guide me to ``Person.Student.getAverageScrore(student)`` but if I define ``getStudentAverageScore(student)`` I will have to remember in which file I have defined it and Intellisense won't be of any help to find it.  

At last, when I am browsing to my code if I don't care about the type ``Person`` and everything that is related to it I can simply fold its namespace.  


Maybe my approach is not sound there is an obviously better design pattern?
What do you guys think about namespaces?
## [8][I just made my first NodeJS oriented package. Nodehawk - A development watcher for your Node servers.](https://www.reddit.com/r/typescript/comments/f4rkrx/i_just_made_my_first_nodejs_oriented_package/)
- url: https://github.com/samrith-s/nodehawk
---

## [9][debog - A simple drop-in decorator to track timing of class methods](https://www.reddit.com/r/typescript/comments/f4j7i2/debog_a_simple_dropin_decorator_to_track_timing/)
- url: https://github.com/UnicornHeartClub/debog
---

## [10][Quick Yet NOT dirty — lazy form of the Observer Pattern](https://www.reddit.com/r/typescript/comments/f4he51/quick_yet_not_dirty_lazy_form_of_the_observer/)
- url: https://medium.com/@lironhazan/quick-yet-not-dirty-lazy-form-of-the-observer-pattern-6d2672ead884?source=friends_link&amp;sk=b6c300d0f210375c201c5828c538bdbc
---

## [11][How to refactor this code using Optional Chaining?](https://www.reddit.com/r/typescript/comments/f48m23/how_to_refactor_this_code_using_optional_chaining/)
- url: https://www.reddit.com/r/typescript/comments/f48m23/how_to_refactor_this_code_using_optional_chaining/
---
`res.data.items ? res.data.items[0].id : undefined`

I've tried this but didn't work

`res.data.items[0]?.id`

`res.data.items?[0].id`
