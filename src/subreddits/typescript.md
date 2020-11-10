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
## [2][10 Insights from Adopting TypeScript at Scale](https://www.reddit.com/r/typescript/comments/jrgi8z/10_insights_from_adopting_typescript_at_scale/)
- url: https://www.techatbloomberg.com/blog/10-insights-adopting-typescript-at-scale/
---

## [3][Adding 3. party tools to VSCode extension](https://www.reddit.com/r/typescript/comments/jrhivw/adding_3_party_tools_to_vscode_extension/)
- url: https://www.reddit.com/r/typescript/comments/jrhivw/adding_3_party_tools_to_vscode_extension/
---
This may not be the correct reddit but I feel like there are only users at /vscode so I'm hoping there are VSCode extension developers here who might have an answer to my question.

First I'm trying to find a library to extract a .app file. It's an archive file that can be extracted in windows no problem. But I've not be able to find any that can extract my files and I have no idea how to write my own extract code. Either they fail out with messages like archive is of an unsupported format or they just dont do anything.

Therefor I'm thinking I just write a .Net commandline tool that can extract the file and I'll call that from extension code with child\_process but then my second problem occurs. No matter where I place the exe tool in my project it is just never included into the vsix file.
## [4][Terraform with TypeScript](https://www.reddit.com/r/typescript/comments/jqxmoq/terraform_with_typescript/)
- url: https://medium.com/francisvitullo/terraform-with-typescript-7643defb4eb1?source=friends_link&amp;sk=975dc30cd7d48d989f2d7f0c16882c2e
---

## [5][Help pass - Quiz: "Parse nullable string"](https://www.reddit.com/r/typescript/comments/jrdaok/help_pass_quiz_parse_nullable_string/)
- url: https://www.reddit.com/r/typescript/comments/jrdaok/help_pass_quiz_parse_nullable_string/
---
Hi All, 

I am trying to learn coding in my spare time and have run into a bit of trouble.  Can anyone help?

P.S. This is not homework.  This is from a website called execute program.  

\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_

This is the prompt:

*Write a function that turns a* *string | undefined* *into a* *number.*

* *If the argument was* *undefined*  
*, the function should return* *undefined.*
* *If the argument was a string containing a valid number, it should return that number.*
* *If the argument was a string without a number, it should return undefined.*

*(You'll want to use the built-in functions* *parseInt(s: string)*  
 *and* *isNaN(n: number)**.)*

&amp;#x200B;

These are the tests you have to pass.  Currently I am failing test #1: 

&amp;#x200B;

 \&gt;maybeParseString(undefined) Expected: 

    undefined

but got: 

    type error: type error: Argument of type 'undefined' is not assignable to parameter of type 'string'.

\&gt;maybeParseString("3") Expected: 

    3

OK!

\&gt;maybeParseString("2701") Expected: 

    2701

OK!

\&gt;maybeParseString("junk") Expected: 

    undefined

OK!

\&gt;maybeParseString("not-a-number") Expected: 

    undefined

OK!

\&gt;maybeParseString(null) Expected: 

    type error

OK!

\&gt;maybeParseString(3) Expected: 

    type error

OK!

7 tests, 1 failure 

\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_\_

This is my code so far (which is wrong):

&amp;#x200B;

**function maybeParseString(s: string ): number | undefined {**

  **const n = parseInt(s)**

  **if (isNaN(n)) {**

**return undefined**

  **}**

  **else {return n}**

**}**

&amp;#x200B;

I have tried 1000 and 1 different ways and can't seem to crack the code. 

&amp;#x200B;

Thanks!
## [6][TypeORM: EntityMetadataNotFound: No metadata for "User" was found.](https://www.reddit.com/r/typescript/comments/jr65m4/typeorm_entitymetadatanotfound_no_metadata_for/)
- url: https://www.reddit.com/r/typescript/comments/jr65m4/typeorm_entitymetadatanotfound_no_metadata_for/
---
Hey all! 

Hoping some typeORM experts can help me out, burned wayyy to much time on this. 

Was wondering if anyone can shed some light on this error, I know there are a lot of questions and answers relating to this that are mainly around not having the correct destination in the \`entities\` section of the connection or just not including the entity at all. In my case this only happens when I try run tests (Jest + Supertest). I've tried adding the path to my entities instead of them individually but also doesn't work Wondering if it is something to do with the way I am setting up the test connection and that the rest of my code can't see the new connection. Below is my \`createConnection\` function: 

    export const createTestConnection = async () =&gt; createConnection({
        dropSchema: true,
        entities: [User], // There will be more
        logging: ["error"],
        maxQueryExecutionTime: 500,
        name: "test",
        schema: "test",
        synchronize: true,
        type: "postgres",
        url: process.env.DB_URL,
    });

I checked the DB (removed \`dropSchema\`) and the entity is there but when I run \`find()\` method in my controller it throws the error:

    import { EntityRepository, Repository } from "typeorm";
    import { User } from "../entities/User";
    
    @EntityRepository(User)
    export class UserRepository extends Repository&lt;User&gt; {
    
        findAll(): Promise&lt;User[]&gt; {
            return this.find()
        }
    
        findUserByKey(key: string): Promise&lt;User | undefined&gt; {
            return this.findOne({ key });
        }
    }
    

Stacktrace:

        EntityMetadataNotFound: No metadata for "User" was found.
    
          at new EntityMetadataNotFoundError (src/error/EntityMetadataNotFoundError.ts:9:9)
          at Connection.Object.&lt;anonymous&gt;.Connection.getMetadata (src/connection/Connection.ts:333:19)
          at EntityManager.Object.&lt;anonymous&gt;.EntityManager.getCustomRepository (src/entity-manager/EntityManager.ts:1255:86)
          at Connection.Object.&lt;anonymous&gt;.Connection.getCustomRepository (src/connection/Connection.ts:368:29)
          at Object.getCustomRepository (src/index.ts:300:55)
          at getAllUsers (controllers/User.ts:7:46)
          at Layer.handle [as handle_request] (node_modules/express/lib/router/layer.js:95:5)
          at next (node_modules/express/lib/router/route.js:137:13)
          at Route.dispatch (node_modules/express/lib/router/route.js:112:3)
          at Layer.handle [as handle_request] (node_modules/express/lib/router/layer.js:95:5)
## [7][[Help] Interface with function which can accept classes instance as an argument.](https://www.reddit.com/r/typescript/comments/jr52n0/help_interface_with_function_which_can_accept/)
- url: https://www.reddit.com/r/typescript/comments/jr52n0/help_interface_with_function_which_can_accept/
---
I am new with Typescript and having some problems. I have a task where I need to create a base class. Then I need to inhered three classes. In those classes I have to use generic types, unions etc. Easy so far, but the problem is that after this I have to create an interface that has a function which can accept each of those four classes instance as an argument. And that function must only have one argument. Return type of function must be Boolean. Finally I have to create an instance of the interface and implement the function which should print the name of input class to console.

Any help with interface and function appreciated!
## [8][Are the socket.io and socket.io-redis @types files fundamentally broken or (much more likely) am I doing something silly?](https://www.reddit.com/r/typescript/comments/jr0mv8/are_the_socketio_and_socketioredis_types_files/)
- url: https://www.reddit.com/r/typescript/comments/jr0mv8/are_the_socketio_and_socketioredis_types_files/
---
I have an empty project into which I installed \`@types/socket.io\` and \`@types/socket.io-redis\` as dev dependencies. Add in a do-nothing index.ts file, try to compile and ...

    error TS2694: Namespace '"/Users/[...]/node_modules/socket.io/dist/index"' has no exported member 'Adapter'.

Now listen here you little shit, etc. SocketIO.Adapter looks to me as though it's exported on line 835 of said file.

As I say, I certainly imagine it's me doing something daft here but given there are so few moving parts involved, what on earth is it?
## [9][[Help] Parameter in function must be a key of a type](https://www.reddit.com/r/typescript/comments/jr38uu/help_parameter_in_function_must_be_a_key_of_a_type/)
- url: https://www.reddit.com/r/typescript/comments/jr38uu/help_parameter_in_function_must_be_a_key_of_a_type/
---
Let's say I have something like this:

    type BookColor = {
      red: string;
      orange: string;
      blue: string;
      green: string;
    }
    
    function newBook(color: BookColor[keyof BookColor]) {
    
    }
    
    newBook('pink'); // Should throw error
    
    newBook('orange'); // SHOULD BE GOOD

I am trying to make it so the parameter in `newBook()` can only be the type of `BookColor` (as a `string`) but I don't think what i have is working.

How can I fix this?
## [10][Stator: A full-stack boilerplate – releases, deployments, enforced conventions](https://www.reddit.com/r/typescript/comments/jqpizp/stator_a_fullstack_boilerplate_releases/)
- url: https://www.reddit.com/r/typescript/comments/jqpizp/stator_a_fullstack_boilerplate_releases/
---
Have you ever started a new project by yourself?

If so, you probably know that it is tedious to set up all the necessary tools. 

Just like you, the part I enjoy the most is coding, not boilerplate.

&amp;#x200B;

Say hi to [stator](https://github.com/chocolat-chaud-io/stator), a full-stack TypeScript template that enforces conventions, handles releases, deployments and many more features!
## [11][How to use Contact Picker API in typescript?](https://www.reddit.com/r/typescript/comments/jqu6yo/how_to_use_contact_picker_api_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/jqu6yo/how_to_use_contact_picker_api_in_typescript/
---
Hi.I need to use [Contact Picker API](https://wicg.github.io/contact-api/spec/) in my angular project. but it doesn't work because contacts and ContactsManager is not declared in Navigator in typescript dom library. and the error say "property contatcs does not exist on Navigator"

here's some working examples in js:[https://whatwebcando.today/contacts.html](https://whatwebcando.today/contacts.html)[https://contact-picker.glitch.me](https://contact-picker.glitch.me/)

here's my test angular app:  [https://codesandbox.io/s/angular-contacts-85x29-85x29](https://codesandbox.io/s/angular-contacts-85x29-85x29)

I copied one of the examples and it doesn't work. I tried to add interface for contactManager and contact and merge it with Navigator in contacts.d.ts and added it in types in tsconfig, but I think it's not correct. I appreciate if you check my code or give me some hint on what to do to make it work.

thanks in advance.
