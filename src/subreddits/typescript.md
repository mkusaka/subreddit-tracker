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
## [2][Announcing TypeScript 4.0](https://www.reddit.com/r/typescript/comments/ideto6/announcing_typescript_40/)
- url: https://devblogs.microsoft.com/typescript/announcing-typescript-4-0/
---

## [3][I hate mocking Typescript classes with Jest](https://www.reddit.com/r/typescript/comments/idvumf/i_hate_mocking_typescript_classes_with_jest/)
- url: https://www.reddit.com/r/typescript/comments/idvumf/i_hate_mocking_typescript_classes_with_jest/
---
I find mocking typescript classes with Jest, so painful, does anyone else find this?

I'm trying to mock a single class, and it just turns into so much code.

Are there are good resources / libraries for mocking classes with typescript?
## [4][Command line tool to organize your Typescript imports](https://www.reddit.com/r/typescript/comments/idqc9r/command_line_tool_to_organize_your_typescript/)
- url: https://www.npmjs.com/package/import-conductor
---

## [5][Getting Started with TypeScript : Set up and Intro](https://www.reddit.com/r/typescript/comments/idupum/getting_started_with_typescript_set_up_and_intro/)
- url: https://blogs.rajankalwar.com.np/getting-started-with-typescript-set-up-and-intro-cke2ca7db00f7b8s1doa6817v
---

## [6][Deriving sub interface using utility type Omit and Pick . Some issues with refactoring and linting .](https://www.reddit.com/r/typescript/comments/idu0k5/deriving_sub_interface_using_utility_type_omit/)
- url: https://www.reddit.com/r/typescript/comments/idu0k5/deriving_sub_interface_using_utility_type_omit/
---
Here is one example :

    interface mega {
    	a : string,
    	bb : number,
    	c : boolean
    }
    
    type mini1 = Pick&lt;mega,"a"|"bb"/*does give intellisense*/&gt;;
    type mini2 = Omit&lt;mega,"a"/*does not give intellisense*/,"d"/*does not even exist and does not lint error*/&gt;;

1)You get intellisense for the properties for Pick but not for Omit . Why ? 

2)You refactor a to aa and it does not work in both Pick and Omit . Why ?

3)Omit does not lint error if you omit properties that do not exist . Why ?
## [7][Learning typescript](https://www.reddit.com/r/typescript/comments/id8ju3/learning_typescript/)
- url: https://www.reddit.com/r/typescript/comments/id8ju3/learning_typescript/
---
Hi,

I want to learn typescript, not for a server-side or a client-side application but for writing standalone applications. I am not very familiar with javascript, I mainly write in Python. What topics should I learn before typescript? From my understanding, I need at least to learn the basics of node.js and javascript. Is there a good tutorial or course that covers all these topics? 

Thanks!
## [8][For a function that gets as an argument a single object with optional properties that have initial values , how to type the argument ?](https://www.reddit.com/r/typescript/comments/idacac/for_a_function_that_gets_as_an_argument_a_single/)
- url: https://www.reddit.com/r/typescript/comments/idacac/for_a_function_that_gets_as_an_argument_a_single/
---
For example :

    function foo({a = 1 , b} = {}) {
        console.log(a,b);
    }
    foo({});//1 undefined
    foo();//1 undefined
    foo({a : 2});//2 undefined
    foo({b : 2});//1 2
    foo({a : 3 ,b : 2});//3 2

How do I type the argument of that function in typescript ?
## [9][Does interface segregation principle has anything to do with preventing collisions of interfaces that are merged?](https://www.reddit.com/r/typescript/comments/ideebf/does_interface_segregation_principle_has_anything/)
- url: https://www.reddit.com/r/typescript/comments/ideebf/does_interface_segregation_principle_has_anything/
---
The question in the title .

Sometimes I find my self having to create bigger interfaces by composing the smaller ones . The issue here is that the smaller ones get edited once in a while and nothing guarantees that they do not collide when they are merged .

As far as I understand ISP has nothing to do with preventing collisions .

Then how does someone prevents the collisions ? Unit testing ? 

I did [a post](https://www.reddit.com/r/typescript/comments/hjif2w/checking_for_interface_collisions/) some time ago , similar to this one and got no real help .
## [10][Reflect type of generic to class](https://www.reddit.com/r/typescript/comments/idbd90/reflect_type_of_generic_to_class/)
- url: https://www.reddit.com/r/typescript/comments/idbd90/reflect_type_of_generic_to_class/
---
Hey there TypeScript  wizards of the internet! :3

I'm trying to do the following and I'm not sure if it's somehow possible to archive with TypeScript. I'd like to write a database repository like class that takes a generic and then uses it in its constructor to create the database connection to the table provided as the generic.

I'm using TypeOrm which has a getRepository function that takes an Entity class which represents a table. My problem is that this function takes a class as a parameter and I'm not able to use the type provided as a generic.

I'll try to illustrate what the end result in my head would look like: 

    declare class BaseEntity {};
    // Dummy class that represents a table
    declare class User extends BaseEntity {};
    
    // Imitates the functionality of the getRepository function
    type DBConnection&lt;T&gt; = { getOne: () =&gt; Promise&lt;T&gt; }
    type getRepository &lt;T extends BaseEntity&gt; = (entity: T) =&gt; DBConnection&lt;T&gt;;
    
    // Imitates the getRepository function provided by TypeOrm
    declare const getRepository: getRepository&lt;any&gt;;
    
    class DatabaseConnector&lt;T&gt; {
        connection: DBConnection&lt;T&gt; 
        
        constructor() {
            // Establish database connection for class provided as the generic
            // Here I get the error that 'T' only refers to a type but is being used as
            // value which is obvious bcs. it is just a type.
            this.connection = getRepository(T);
        }
    
        // Some dummy function to use the connection
        findOne(): Promise&lt;T&gt; {
            return this.connection.getOne();
        }
    }
    
    // Bootstrapping function for an async context
    const bootstrap = async () =&gt; {
        const db = new DatabaseConnector&lt;User&gt;();
        // This should now hold a user from the db;
        const user = await db.findOne();
    }

My question now refers to the problem lying in the constructor. How can I reflect from the generic type to the actual class it represents? Is this even possible using TypeScript? Can I use reflect-metadata to accomplish what I'm trying to do? 

I know that some other languages like C# or Java have offer similar features but I haven't had any experience with that either. 

Thank you so much for your help!
## [11][How to correct relative paths to other file types in same directory when compiling TS?](https://www.reddit.com/r/typescript/comments/icz5ig/how_to_correct_relative_paths_to_other_file_types/)
- url: https://www.reddit.com/r/typescript/comments/icz5ig/how_to_correct_relative_paths_to_other_file_types/
---
So let's say we have ts and css file in the "src" directory. In ts file we have import './style.css'. After compiling TS to "lib" directory, that path is no longer valid for compiled file, because css file is not compiled. How would you make it change to '../src/style.css'?
