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
## [2][Types as axioms, or: playing god with static types](https://www.reddit.com/r/typescript/comments/ia6d8s/types_as_axioms_or_playing_god_with_static_types/)
- url: https://lexi-lambda.github.io/blog/2020/08/13/types-as-axioms-or-playing-god-with-static-types/
---

## [3][For people who think changing types constitutes a breaking change, do you consider upgrading the Typescript version you compile your library with a breaking change?](https://www.reddit.com/r/typescript/comments/ia6zvy/for_people_who_think_changing_types_constitutes_a/)
- url: https://www.reddit.com/r/typescript/comments/ia6zvy/for_people_who_think_changing_types_constitutes_a/
---
This is branching of another post, but it seemed a lot of people thought a change to the types that may cause typings to fail consortia a breaking change. To follow up, I'm curious to know what people's thoughts is on the actually Typescript library.

I can see an argument being made that if I upgrade Typescript library then the way my types are generated can include types that only exist in a later version of typescript. Therefore, a consumer of the library who is on an older version of typescript would have an error and would need to upgrade Typescript.

On the other hand, it's usually not the case as per how npm works that you release a breaking change if an internal dev dependency library upgrades.

I personally don't think it should be considered a breaking change, but I would still love to hear more sides to this.
## [4][For libraries distributed with TypeScript types, is changing or removing a Type meant to be a breaking change for the library?](https://www.reddit.com/r/typescript/comments/i9rhrw/for_libraries_distributed_with_typescript_types/)
- url: https://www.reddit.com/r/typescript/comments/i9rhrw/for_libraries_distributed_with_typescript_types/
---
Wouldn't have any impact on js consumers, but could potentially break TypeScript apps?
## [5][Overload function can't find input properties](https://www.reddit.com/r/typescript/comments/i9nmdy/overload_function_cant_find_input_properties/)
- url: https://www.reddit.com/r/typescript/comments/i9nmdy/overload_function_cant_find_input_properties/
---
This is my first attempt at a function overload so the error may be basic. The intent is that this will normally be called via cron job (overload 1). But on an initial account creation, I want this to work as a route handler callback as well (overload 2):

    interface VoipMsProperties {
      userName: string,
      apiPassword: string
    }
    
    async function fetchCallData(voipMsData: VoipMsProperties, res: Response): Promise&lt;void&gt;;
    async function fetchCallData(req: Request, res: Response): Promise&lt;void&gt;;
    
    async function fetchCallData(input: VoipMsProperties | Request, res: Response) {
      let userName;
      let apiPassword;
    
      if (input instanceof Request) {
        userName = input.body.userName;
        apiPassword = input.body.apiPassword;
      } else {
        userName = input.userName; // error 1 below
        apiPassword = input.apiPassword; // error 2 below
      }
      &lt;continues ...&gt;
    
    /*
    any
    Property 'userName' does not exist on type 'VoipMsProperties | Request&lt;ParamsDictionary, any, any, ParsedQs&gt;'.
      Property 'userName' does not exist on type 'Request&lt;ParamsDictionary, any, any, ParsedQs&gt;'.ts(2339)
    
    any
    Property 'apiPassword' does not exist on type 'VoipMsProperties | Request&lt;ParamsDictionary, any, any, ParsedQs&gt;'.
      Property 'apiPassword' does not exist on type 'Request&lt;ParamsDictionary, any, any, ParsedQs&gt;'.ts(2339)
    */

I can see the interface right above the overloads have both `userName` and `apiPassword` so why is it not being found on the interface type `VoipMsProperties` ?

Side question: Does each overload signature have to be exported? Or just the final one with the function definition?
## [6][I created a video to share 3 typing basics to improve your TypeScript code](https://www.reddit.com/r/typescript/comments/i98t3e/i_created_a_video_to_share_3_typing_basics_to/)
- url: https://www.youtube.com/watch?v=UuBJrAZsp4Y
---

## [7][TypeScipt is just another effort by Anders Hejlsberg to statically type a dynamic world.](https://www.reddit.com/r/typescript/comments/ia0lui/typescipt_is_just_another_effort_by_anders/)
- url: https://medium.com/@lonniebiz/typescipt-is-just-another-effort-by-anders-hejlsberg-to-statically-type-a-dynamic-world-324e4320e13a
---

## [8][Working on TS focused React form lib - feedback needed](https://www.reddit.com/r/typescript/comments/i9a5s2/working_on_ts_focused_react_form_lib_feedback/)
- url: https://www.reddit.com/r/typescript/comments/i9a5s2/working_on_ts_focused_react_form_lib_feedback/
---
Proof of concept and API proposal for Typescript-focused, schema-based React form library:

[https://codesandbox.io/s/formts-poc-3ryqv?file=/examples/example\_1.tsx](https://codesandbox.io/s/formts-poc-3ryqv?file=/examples/example_1.tsx)

The sandbox contains README with more info and 5 examples of usage. The API is just a prototype but everything should be feasible to implement at this point.

I'd like to know if there is any interest in such lib being implemented?

EDIT: please see README in the sandbox for explanation why I though another form lib might be needed :)
## [9][Error on creating module using monorepo react structure](https://www.reddit.com/r/typescript/comments/i9bkdz/error_on_creating_module_using_monorepo_react/)
- url: https://www.reddit.com/r/typescript/comments/i9bkdz/error_on_creating_module_using_monorepo_react/
---
Hey guys, i'm working on a project that will be a monorepo using yarn workspace, with micro-front-end using react-typescript.  To create the frontend structure i used  Create-React-App and everything was working fine.  Then i decided to create a priveta-route module using just  a basic index.tsx and package.json when i use this module i get this error:

    C:/Users/johnvidal/Documents/quero-docura/quero-docuras/packages/web-private-route/src/index.tsx 10:0
    Module parse failed: The keyword 'interface' is reserved (10:0)
    You may need an appropriate loader to handle this file type, currently no loaders are configured to process this file. See https://webpack.js.org/concepts#loaders
    | 
    | //@ts-ignore
    &gt; interface PrivateRouteProps extends RouteProps {
    |   component: React.ComponentType&lt;RouteComponentProps&lt;any&gt;&gt;
    | } 

How do i fix this?
## [10][Is there an 'Exact&lt;T&gt;' advanced type? Something that gurantees you can only put in the exact attributes required and no more?](https://www.reddit.com/r/typescript/comments/i8vxz2/is_there_an_exactt_advanced_type_something_that/)
- url: https://www.reddit.com/r/typescript/comments/i8vxz2/is_there_an_exactt_advanced_type_something_that/
---
Eg. I have the following object:

    type User = {
        id: string
        email: string
        password: string
    }

    type UserProfile = User &amp; { name: string, dob: string }

    function insert(user: User) {
        UserDB.insert(user)
    }

So I have a User, a UserProfile, and an insert function that inserts Users into the DB.

Now my problem is that the insert function connects to a NoSQL DB. So I can technically pass in a UserProfile, and it will end up inserting the extra attributes into my NoSQL DB. Over time my NoSQL DB might end up with a lot of extra unnecessary attributes.

I was wondering if there was an Exact&lt;T&gt; type that ensures you only have User.id, email &amp; password added in and nothing more.
## [11][Problem with request handlers in node.](https://www.reddit.com/r/typescript/comments/i96454/problem_with_request_handlers_in_node/)
- url: https://www.reddit.com/r/typescript/comments/i96454/problem_with_request_handlers_in_node/
---
I have this piece of code: 

    router.route('/signup').post(validateRegister);

But I get an error message saying that `validateRegister` is not assignable to type `RequestHandler&lt;ParamsDictionary, any, any, ParsedQs&gt;` The code for validate register is this: 

    import { findUserByUsername } from '../helpers';
    import { NextFunction } from 'express';
    import { UserRegisterRequest } from '../interfaces';
    import { Response } from 'express';
    
    export default async function validateRegister(
      req: UserRegisterRequest,
      res: Response,
      next: NextFunction
    ): Promise&lt;void&gt; {
      const { fullname, username, email, password, password2 } = req.body;
      // bunch of if statments ....
    }


Any help with this issue would be appreciated.
