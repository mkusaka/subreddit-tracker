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
## [2][I created a video to share 3 typing basics to improve your TypeScript code](https://www.reddit.com/r/typescript/comments/i98t3e/i_created_a_video_to_share_3_typing_basics_to/)
- url: https://www.youtube.com/watch?v=UuBJrAZsp4Y
---

## [3][Working on TS focused React form lib - feedback needed](https://www.reddit.com/r/typescript/comments/i9a5s2/working_on_ts_focused_react_form_lib_feedback/)
- url: https://www.reddit.com/r/typescript/comments/i9a5s2/working_on_ts_focused_react_form_lib_feedback/
---
Proof of concept and API proposal for Typescript-focused, schema-based React form library:

[https://codesandbox.io/s/formts-poc-3ryqv?file=/examples/example\_1.tsx](https://codesandbox.io/s/formts-poc-3ryqv?file=/examples/example_1.tsx)

The sandbox contains README with more info and 5 examples of usage. The API is just a prototype but everything should be feasible to implement at this point.

I'd like to know if there is any interest in such lib being implemented?

EDIT: please see README in the sandbox for explanation why I though another form lib might be needed :)
## [4][Is there an 'Exact&lt;T&gt;' advanced type? Something that gurantees you can only put in the exact attributes required and no more?](https://www.reddit.com/r/typescript/comments/i8vxz2/is_there_an_exactt_advanced_type_something_that/)
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
## [5][Error on creating module using monorepo react structure](https://www.reddit.com/r/typescript/comments/i9bkdz/error_on_creating_module_using_monorepo_react/)
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
## [6][Problem with request handlers in node.](https://www.reddit.com/r/typescript/comments/i96454/problem_with_request_handlers_in_node/)
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
## [7][Type annotation erroneously overrides generic](https://www.reddit.com/r/typescript/comments/i8zupk/type_annotation_erroneously_overrides_generic/)
- url: https://www.reddit.com/r/typescript/comments/i8zupk/type_annotation_erroneously_overrides_generic/
---
I'm currently trying to refactor a function's signature to remove the need for loads of overloads. I've done this by creating a conditional generic type to handle the same cases that the overloads did.

The function works fine but for some reason adding annotations to constrain the return type from `unknown` to something else is giving me errors that didn't exist, as if the function is incorrectly inferring `T` from the annotation and I just can't figure out why. There are some times in my project where I need to declare a variable before assigning it and it just breaks.

This is the function's signature:

    const createMatrix = &lt;D extends number, T&gt;(
        dimensions: D,
        initialValues: T = null
    ): Matrix&lt;D, T&gt;
    
    type Matrix&lt;D extends number, T&gt; = D extends 1
        ? T[]
        : D extends 2
        ? T[][]
        // ...

And a TS playground link with the full code, tests and old version [is here](https://www.staging-typescript.org/play?#code/PTAEBUE8AcFNQIYDskHsAuD0EtVNAO4BOeA5gDaSioBusRR2AJrAM6imxL3YDGh2dAAtQ9EkVYAoSSFABRAB5xe6WE1AAjWEIQ1cRAFwywAQSZNsSUqHQx4yNJhx526VKBoJGCDeTahWIVQAV3J1LVBeVABbaCxsX3gCQRFheAAzYKQVXCR2Ilh0YKIkYwDsUiQsYtgAGkiEYNZLazQbO1EGVAlpWRMVYIRyTW1dfSM+1mbKlsR8BwwsNQ8vbB8-G3c00AKikpXyYPhUdJshDKycvAam-2RO8QA6MvBzji4efnBQbHZiMkoPyQ6TEyzcIzmVBOoAA5F4iAgqARzgVQCxoMJQAAeUAAERhhBR8EEASCoSYZQilhBDDB7jwgO2GgQrHgtjg1FO224rFU6nhiMeoAASrBorRZtsFk5cgFUOQ6FJZNtoCREtFNMF0DsxbR-Oz4KwECDbISuCtvIlEFMKkholxtTpoHA8oh0qoiJECvE8M9egAqf2SUD+kWFYquhCgaJYRgKUAACmioRwAFoLPa8rkhogGIiAJQCTGQEKeljNArqDNcZouObqTyHNjPEPB0MAATiCI11azddToAABrjBwYIG9y9hK2jsJna67odsY+g423QJ2vAgNZZBGtyAA1IZHdgD17wRtHM5YUQIXgiaCoSzayxneDLuNA3dDX761BC3GwOkjTkOgrjuIOSChOQg4th2uwRuwJjRrG2DxtCw6DjOc65Kw9TJJisC3vej5IM+SBfuQP7qOC2yDjuOBDIeTasJhXZboU9CwcAkhRHk2q8N6qgALIofGAC82K4qICiqEgTDsJB0RaEQ9TgAAfAmwagNpvbzqwY64rUWnafRe5MceY7fAAPqAkHkMMEl2eQkj5mOIkrqhWKGRAamgGJvkAN7GZELj8cUBSkbis41rkAAyXCkJiEm6ThADcwW8byOoxjuVhRdhdbJdFfaugOACM6XadpmXatwaisKKvDFPOfnZQguWkPlMV1r5AAM6UZaFkThQ67kfhJJh5pACZNbSkXFfO8VWMI+aPOk2D2QmpmMUebD5gNVUhXxoDrVU5BjahrV1fJjXNbkwWHQA-MNc3oBdCiPDG0AJgmhb+Q9h2AwJhHCaJCYFDl5F5QtOH1NtB67aw+YA4dyOA1VY6zRFb2iQdVXwfsp1DO91qgO9XmqWp6UAL4DbIZ6gKqqDqgEsBimBELCNeCbIua0pLEwhbbAaPzsNSoLUe4ETbFksonGUhF3u0HKqrAvBqJKby8RYzhnVe2oE66K5HM8Ivk1JsAyVw8m2cESn0JTrUW1bcnsGVwXPeAADaAC6wUGdJsk2wATB7EC+77-t4oH1vsAAzGH3s+xHfuHQHltB+wAAsicpynUfIJAefJyXA01TYbDoGVY6KcpvutcDSzvQmZX7aAshYgOBDdAA1ie7zcIwXyi6z2rghBdvKYOPFDaovLBzXk-0CnDeCbAzfB23HcDkwqBsEgMLat3RA96AA6cIPfAQCPrJj+Btf0KANkP0QvuYQOgQhGEEKDlkPdoAQJAg56hpHwDVBElg6S23tkQGex057oDjovGBxdV4g3XmDOOW8wCdzRHvVgB8j69zPgPT419fij02EOF+T9oF1x9rQl+Kd36ki-uEeAv8kD-1QIA4Br4wGhQgdwSWdD6DSHLggrOaCm5g03qlduOCBwFFIKELwQIaRcHVoQXu-cL5kO+BQzh3DeH1ENsrYk7A-4AKQPnCRlcACsyD6H1wko3UGHkFAJmDvUXq2DsRdx0WafAdFyIMQRsxTCFDb4kL0UPch7AJ4wL4WYkWFCmElzgVlBBAA2McVieE2JLtI9xcYvF+NwcfPuJ1IH1BZOY+JoB8mAL9L0MAAAhWA5AeEj22PKJgeFe4tHqJkbIushQAEkYQahXJASU7gIZ6moHQIgXSEA23SN0RAIx0AejxAADXqOcbIbI3jcAICdS4utEw+D1MjSQIyrhgLXu9AA8mELE6kEwpRcGOMqcNQlmURo9Syrlw4+wGg8q5biMEeLeUwD5GlgrfLyGObxwV4bmTYMCiALlLIpwhZc2U0LXnvM+UimGPzQBxyModDFQLLK4rBfi6QkKiXPNEnChFmlDrIv0qALONKqp0uYti8AjKk7FwJaMtl6CSXwrJTyilKLREqXRQCnaIqGWgrlQiqVjyvSyo5aSxFiqCrKpfoKky6rwkWWvs-KCV0oKMp1epUAQVDrlyxg6LqJUlqJREEVM1rBKpVXLhDdqUNOpKvYIG7qpVQAVUGsda6DU1Z3WuBJcNHUfV6VAH1PG1UhpetIiTCaU0ZojXmmav1K01obXIFta1mKkYFqOllIm51RJXTZjdNNEh7ro20s9YtOMPGfQQN9X6fk1Io0HcSo1TBwZigjS0HNsNPxhObWjQd2lt3o0xpW0dcZW1mI7e9Gm0gzYLoRa1CVJcV42RdWpX2A1Mlj0rnC6uKqXEGpkbCsILc-GVPYKdWAb6K68jhQvb9RTXHsv-YuuRCjtEnxA5A8DCC4VIJgyvODhqEMJiwfI2QwHqncHEbPD9YQpF4b-XGOFZTiNgFI6BjDVGmCOJwww2jJTUIMe8aAXxTGUNVNY3YyDYRcmNK4dY3Dv7eMKH40BwJoGgA).

Does anyone have any ideas for what the problem might be?
## [8][VSCode Snippets](https://www.reddit.com/r/typescript/comments/i9160u/vscode_snippets/)
- url: https://www.reddit.com/r/typescript/comments/i9160u/vscode_snippets/
---
Hello, I was looking for a snippet extension to vscode to help me speed my writing but all I found was this [one](https://marketplace.visualstudio.com/items?itemName=dsznajder.es7-react-js-snippets) but it doesn't include Typescript. Any suggestions?
## [9][Validating object's type at runtime?](https://www.reddit.com/r/typescript/comments/i8yk6i/validating_objects_type_at_runtime/)
- url: https://www.reddit.com/r/typescript/comments/i8yk6i/validating_objects_type_at_runtime/
---
For a project I am currently working on, an SDK for consuming a backend API (written in Go), I would like to implement runtime type checking. As I've worked intensively with Typescript over the last couple of years, I know type safety is not guaranteed at runtime, as all type annotations will get lost during the compilation stage.

I have some projects using Typescript as a production dependency, to for example generate custom JSON decoders to allow code generation to take place based on Typescript interfaces: https://github.com/PicnicSupermarket/aegis. 

A thought that came to mind, was using Typescript as a dependency and use it to check against an interface, but as I'm not really familiar with the structure of the Typescript compiler, I don't know what's possible.

So my question is: **is it possible to take a Javascript object, together with a Typescript interface, and validate whether the Javascript object satisfies the Typescript interface**? And if that can be done by calling the Typescript compiler as a production dependency somehow.

To give some further context, this is the function where results pass through in the SDK I am working on:

```
const throwOrReturn = &lt;T&gt;(result: APIResult&lt;T&gt;): T =&gt; {
    if (result.error) {
        throw new Error(result.error);
    }
    // We can assume that data is valid (type T) if no error was found
    return result.data as T;
};

interface APIResultSuccess&lt;T&gt; { data: T; error?: undefined; }
interface APIResultFailure { data?: undefined; error: string; }
type APIResult&lt;T&gt; = APIResultSuccess&lt;T&gt; | APIResultFailure;
```

Before `return result.data as T;` I would like to check whether `result.data` is valid if assigned type `T` at runtime. So currently I am using generics, but taking the checking one layer up, where I have access to the actual interface, instead of a generic, would not be an issue.

An example of how the `throwOrReturn` is implemented:

```
interface ReturnValues {
    // Some interface code
}

const exampleCall = async (config: Config): Promise&lt;interface&gt; =&gt; {
    const res = await get&lt;interface&gt;(config, GetEndpoints.EndpointWeNeed); // Wrapper for making API calls and 4xx/5xx error handling 
    return throwOrReturn(res);
};

export default exampleCall;
```
## [10][Trying to get rid of the any; need help](https://www.reddit.com/r/typescript/comments/i93a9s/trying_to_get_rid_of_the_any_need_help/)
- url: https://www.reddit.com/r/typescript/comments/i93a9s/trying_to_get_rid_of_the_any_need_help/
---
I am trying to define types for my little pub/sub service. The service has a  `subscribe` function that is giving me troubles.

This function subscribes callbacks to messages. Each message has the following shape:

```
{
  action: string;
  payload: one_of_the_predefined_payloads
}
```

So what the subscribe function does is it takes two arguments — action and callback — and puts the callback in the set of callbacks that respond to this action. When it the pub/sub service receives a message with a given `action`, it will call the callbacks interested in this action passing to them the payload from the message.

The type of the callback function is currently defined as `type Callback =  (...args: any[]) =&gt; void`; but I am trying to find a way to properly define its argument, because, goddamit, I know the full list of actions and I know which payloads are expected for those actions:

```
type FooPayload = {
    foo: number;
}

type BarPayload = {
    bar: string;
}

interface ActionToPayloadMap {
    foo: FooPayload;
    bar: BarPayload
}

type Callback = &lt;A extends Action&gt;(payload: ActionToPayloadMap[A]) =&gt; void;
```

But somewhere along the way my reasoning falls apart as can be seen in [this typescript playground](https://www.typescriptlang.org/play?#code/C4TwDgpgBAYg9nACgQxAGzsgJlAvFAbwCgpSoAzBALigDsBXAWwCMIAnAbiIF8ijRIUAELI2KdJhz5iZKM1E0AzsDYBLWgHMuvfuGjw4AWQiLFyDdGkkyyAMbBVcWjQDklOC66ywqDNhoG4n5YPFy6giJsxqbmloTWpHYOTq7ybJ4JUD4S-sKiQZKhfOrA7OR20ACC9o60ACpIvpKGyGDxsu4BCAXYXmRpNJE9IToCVTVOeFAA1hAgcORQ1cn1jTlYLWBhY1DDU4FN2FAAPnlih1jbelAAyvTMirZqrGyKUwBKELZwbFgAPMtagAaW4QYB-ADCyDQaHktmmAD4EXwdlCYXDplMAVAIAAPUq0LBvQFOBEACmywRoJNWw02AG1KgBdACUeARUAAbnBVJcUdc7g8nqpWFjKjj8RBCcSJrRyUlatSQbZobC7NMaGi1fCAQi2bgOdzeVdBIh7mhVIoABZiiUEolLWXy2VKrIXamyhp01qM1nsrk8vlEb60ZRQRT3R7PdhvaTcKDIN6CqMimNhENhiNC6M0ZPC0X4MkKpzK1UY-UcmRkDPAcNgqZZlMvRT04u0JknU60CAAd1B4K1GPJLL6pEUYIAdNgsGSVej1SPMo38zHW7KO-hx8BtHwa1lzZarTQzcwLdapkXZSDKZIK+1q04w3PtdNY+HIyvXmuVkzR1BnxiigTpQbAAKJ2Fas7MP6tjMBSFwsiyRTBo+tbuHUJi1oWN65Ac6x3hmcBoBAE4YBo8HrMBCCLkQy7Rn8bgIC45KMR4ILoZhI5AA), where I've captured my attempt at defining these relationships. Could anyone help me out?


**P.S.:** Also, I really don't like the `ActionToPayloadMap` interface from the snippet above. What I would much prefer would be to define the actual message types:

```
type FooMessage = {
    action: 'foo';
    payload: FooPayload
};

type BarMessage = {
    action: 'bar';
    payload: BarPayload
};

type Message = FooMessage | BarMessage
```
and then to somehow tell typescript to lookup the payload type based on the received action. Is typescript capable of this?
## [11][Is it possible to make one function argument optional depending on another argument's value?](https://www.reddit.com/r/typescript/comments/i8wona/is_it_possible_to_make_one_function_argument/)
- url: https://www.reddit.com/r/typescript/comments/i8wona/is_it_possible_to_make_one_function_argument/
---
Hi there! Just starting with typescript and I was guessing if this is possible to achieve. Any help or alternative solutions will be appreciated, thanks in advance!

So I have this method:

    triggerInteraction({
        parent,
        action,
        custom = false,
      }: {
        parent: HTMLElement;
        action: string;
        custom?: boolean;
      }) {
        // Search for trigger and click it if found
        const trigger = custom
          ? (document.querySelector(`[data-logic="${action}"]`) as HTMLElement)
          : (parent.querySelector(`[data-logic="${action}"]`) as HTMLElement);
    
        if (trigger) {
          trigger.click();
          return true;
        } else return false;
      }

Where the ***parent*** argument is only needed if the ***custom*** argument is true. Otherwise, it won't be used by the method.

So my goal is to be able to make this a valid call:

`triggerInteraction({ action, custom: true });`

But not this (***parent*** should be required here, because ***custom*** is false and therefore will be needed in the method):

`triggerInteraction({ action, custom: false});`

Is there any way to set this conditionally optional parameter? Or is there a better approach for this kind of solution?

&amp;#x200B;

Thanks!
