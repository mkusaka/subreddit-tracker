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
## [2][If you had to teach 30+ JS devs TypeScript, what resources would you give them?](https://www.reddit.com/r/typescript/comments/g9j5rb/if_you_had_to_teach_30_js_devs_typescript_what/)
- url: https://www.reddit.com/r/typescript/comments/g9j5rb/if_you_had_to_teach_30_js_devs_typescript_what/
---
Hello!

My company is currently in the architectural stage of building the new version of our web app. For that, our stack is now:

- Next.js
- preact
- Apollo
- mobx

As a way for us to ease our development process and get some assurances about our data structures, we're moving to TypeScript as well. We ran a poll recently, and about half have our JS devs used TypeScript before, and would like to use it again. The other half have heard of it and want to learn. 

For those 15+ JS devs that haven't used TypeScript before, we need some resources for teaching them. For the most part, they have at least 2 years of experience with JS each, and potentially more development experience. 

What resources would you all recommend to get JS devs up to speed as quickly as possible? Right now we front-end masters, which has a decent TypeScript course, and I've also suggested the TS Handbook on the docs site.

Thanks!
## [3][Infer types from object properties?](https://www.reddit.com/r/typescript/comments/g9i0r7/infer_types_from_object_properties/)
- url: https://www.reddit.com/r/typescript/comments/g9i0r7/infer_types_from_object_properties/
---
I am trying to use return type of an object's property as input for another property function:

    type XDefinitions = object;
    type YDefinitions = object;
    
    type SomeObject = {
    	create: (X: XDefinitions, Y: YDefinitions) =&gt; any;
    	getXDefinitions: (configuration: object) =&gt; XDefinitions;
    	getYDefinitions: (configuration: object) =&gt; YDefinitions;
    };

In the best case, I would like to create SomeObject with the X and Y Definitions specified in the according \`getX/YDefinitions\` as input for the the \`create\` function.

I *could* do something like:

    type SomeObject&lt;X, Y&gt; = {
    	create: (X: X, Y: Y) =&gt; any;
    	getXDefinitions: (configuration: object) =&gt; X;
    	getYDefinitions: (configuration: object) =&gt; Y;
    };

And then instantiate my object:

    const myObject = {
    	create: (x, y) =&gt; "someResult",
    	getXDefinitions: (configuration) =&gt; ({
    		"Ax": null,
    		"Bx": null,
    	}),
    	getYDefinitions: (configuration) =&gt; ({
    		"Ay": null,
    		"By": null,
    	}),
    };

But to have myObject correctly typed to \`SomeObject\`, I would have to specify X and Y on creation, but I would like TypeScript to infer those Generics from \`getXDefinitions\` and \`getYDefinitions\`. Is this at all possible, and if so, can anybody explain how?
## [4][Exploring TypeScript from a Business and Software Development Perspective](https://www.reddit.com/r/typescript/comments/g9lg1m/exploring_typescript_from_a_business_and_software/)
- url: https://www.monterail.com/blog/typescript-business-development?utm_medium=social&amp;utm_source=reddit&amp;utm_campaign=typescript
---

## [5][Using async/await to avoid stack overflow error.](https://www.reddit.com/r/typescript/comments/g9kml7/using_asyncawait_to_avoid_stack_overflow_error/)
- url: https://gist.github.com/Jason5Lee/938b451b62e47f52806f5d24b9820644
---

## [6][Job Alerts for Programmers from Reddit and Hacker News (free service)](https://www.reddit.com/r/typescript/comments/g9k5gs/job_alerts_for_programmers_from_reddit_and_hacker/)
- url: https://tryjobalerts.com/?utm_source=reddit&amp;utm_medium=post&amp;utm_campaign=typescript
---

## [7][Is it possible to alias @apollo/react-components QueryResult?](https://www.reddit.com/r/typescript/comments/g958qs/is_it_possible_to_alias_apolloreactcomponents/)
- url: https://www.reddit.com/r/typescript/comments/g958qs/is_it_possible_to_alias_apolloreactcomponents/
---
I am pretty much a newbie to Typescript, overhauling my codebase from .js to .ts/.tsx.

My code is this: -

`interface IQueryResultData {`  
 `currentUser: {`  
 `id: string`  
 `};`  
`};`

`&lt;Query&lt;IQueryResultData&gt; query={ currentUserQuery }&gt;`  
  `{({ data: userData } : IQueryResultData) =&gt; ( &lt;div&gt;[...]&lt;/div&gt; )}`  
`&lt;/Query&gt;`

And I get the following error: -

`TS2322: Type '({ data: userData }: IQueryResultData) =&gt; JSX.Element' is not assignable to type '(result: QueryResult&lt;IQueryResultData, Record&lt;string, any&gt;&gt;) =&gt; Element | null'.   Types of parameters '__0' and 'result' are incompatible.     Property 'currentUser' is missing in type 'QueryResult&lt;IQueryResultData, Record&lt;string, any&gt;&gt;' but required in type 'IQueryResultData'.  MyComponent.tsx(24, 5): 'currentUser' is declared here. types.d.ts(6, 5): The expected type comes from property 'children' which is declared here on type 'IntrinsicAttributes &amp; Pick&lt;QueryComponentOptions&lt;IQueryResultData, Record&lt;string, any&gt;&gt;, "children" | ...`

Is there anything that can be done to remedy this? Thanks.
## [8][Can i import variables from a file to another file?](https://www.reddit.com/r/typescript/comments/g97jfo/can_i_import_variables_from_a_file_to_another_file/)
- url: https://www.reddit.com/r/typescript/comments/g97jfo/can_i_import_variables_from_a_file_to_another_file/
---
Hello,so all im trying is to do a little sign up thing with Typescript. I can import variables and get the stuff from the other file. What im trying to do is:For example i try to use readline to type "name: myname123" and i want in a file called "data.ts/data.txt/data.js idk which one doesnt matter" to be imported. So in short: When i type something with readline i want this thing to be exported in another file and have this line of code there, so i can use it for later.

Sorry if it was kinda bad explaining, im not that good at stuff like this.

Just type something with readline and i want the thing i have written to appear in another file. Thanks in advance! :)

&amp;#x200B;

EDIT: Thanks to everyone for helping me! Much appreciated!   
Solution: fs.writeFile(), fs.appendFile()
## [9][TypeScript doesn't understand url imports](https://www.reddit.com/r/typescript/comments/g9317v/typescript_doesnt_understand_url_imports/)
- url: https://www.reddit.com/r/typescript/comments/g9317v/typescript_doesnt_understand_url_imports/
---
Let's say you want to build a simple app with just ESModules and no transpile step. To do this you can use a CDN like \`unpkg.com\` or \`cdn.pika.dev\` (if you've not tried these you should, it's liberating!). Since you aren't transpiling anymore you aren't using local npm modules, so TypeScript in VS Code starts to throw errors like this:

[Stupid TypeScript](https://preview.redd.it/zym029twqdv41.png?width=549&amp;format=png&amp;auto=webp&amp;s=46ee77a67d63ff41ce20ff3f331d1f0cc17666b4)

Can anything be done to tell TypeScript where to find type definitions for modules like these, or is TypeScript limited to being used as part of a local build process? My gut tells me there ought to be a config property like \`definition-alias: '[https://cdn.pika.dev/@types/'\`](https://cdn.pika.dev/@types/'`) or something to that effect, but I can't seem to find anything.
## [10][Problem Solving with the TypeScript Compiler (Recorded Meetup/Webinar talk)](https://www.reddit.com/r/typescript/comments/g8dlhy/problem_solving_with_the_typescript_compiler/)
- url: https://www.youtube.com/watch?v=ZHiT33F11mk&amp;feature=youtu.be
---

## [11][Class pattern matching library useful in useReducer](https://www.reddit.com/r/typescript/comments/g8j92q/class_pattern_matching_library_useful_in/)
- url: https://www.reddit.com/r/typescript/comments/g8j92q/class_pattern_matching_library_useful_in/
---
Hi all, I have written my first npm library  https://www.npmjs.com/package/@neal83/typeswitch it's only small at 53 lines of code in the main function including type definitions. github.com/neal83/TypeSwitch

But having used redux then moving over to hooks this is my own way of reducing boiler plate while making the data the action type.

You're able to match on class types with the option of using a "when" in combination with default. Like this:

    class MyOwnMarkerClass
    {
        public someProperty: string = "some property value"
    }
    
    const result = TypeSwitch(new MyOwnMarkerClass())(
        Case(Number, When(n: number =&gt; n === 36))
            ((n: Number) =&gt; `you gave me a number and it was 36`),
        Case(Number, When(n: number =&gt; n &gt; 0))
            ((n: Number) =&gt; `you gave me a number and it was positive`),
        Case(String)
            ((s: String) =&gt; `you gave me a string ${s}`),
        Case(MyOwnMarkerClass, When(m =&gt; m.someProperty === "some property value")) //the type is inferred
            ((m: MyOwnMarkerClass) =&gt; `you gave me MyOwnClass with value ${m.someProperty}`)
    );

Hope someone likes it and finds it useful. Thanks
