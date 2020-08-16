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
## [2][How do I publish a library with multiple index.d.ts?](https://www.reddit.com/r/typescript/comments/iarrr4/how_do_i_publish_a_library_with_multiple_indexdts/)
- url: https://www.reddit.com/r/typescript/comments/iarrr4/how_do_i_publish_a_library_with_multiple_indexdts/
---
So I want to publish a components library where the consumer can use components from the main index or a sub-folder where legacy components are stored.

`import Button from "my-library"`

`import Button from "my-library/legacy"`

In tsconfig.json, I have it generate the types:

        "declaration": true /* Generates corresponding '.d.ts' file. */,
        "declarationMap": true /* Generates a sourcemap for each corresponding '.d.ts' file. */,
        "declarationDir": "typing",

So it generates a typings folder with 2 index.d.ts files, one in root and one in \`legacy\` folder.  


My question is I could only define one entry point in the \`types\` field in package.json. How do I include the other typing file in the \`legacy\` subfolder? Is this where I use a triple slash directive?   


Thanks!
## [3][Types as axioms, or: playing god with static types](https://www.reddit.com/r/typescript/comments/ia6d8s/types_as_axioms_or_playing_god_with_static_types/)
- url: https://lexi-lambda.github.io/blog/2020/08/13/types-as-axioms-or-playing-god-with-static-types/
---

## [4][How do you write shared libraries?](https://www.reddit.com/r/typescript/comments/iapnkw/how_do_you_write_shared_libraries/)
- url: https://www.reddit.com/r/typescript/comments/iapnkw/how_do_you_write_shared_libraries/
---
When I write a library for use in a browser I use `"lib": ["DOM"]` in my tsconfig, and for node I install `@types/node`, but what's the best practice for libraries that run on both?

I need to use types that exist in both of these, but not in the non-DOM lib files.
Is there a project that allows me to get the intersection of these types or do I just have to compile twice to check I'm not using a platform-specific API?
## [5][How do you sync one-off ts node scripts between your server and remote?](https://www.reddit.com/r/typescript/comments/iajbyp/how_do_you_sync_oneoff_ts_node_scripts_between/)
- url: https://www.reddit.com/r/typescript/comments/iajbyp/how_do_you_sync_oneoff_ts_node_scripts_between/
---
Long story short, I want to use TS on client and server for everything when it comes to personal work. I don't want to context switch, even though some node command APIs are finicky. I'm writing wrappers around strange node apis to simplify it, such as a promise spawner for processes, and i'm keeping this in git. I made a template TS project that gets generated from a bash script so I don't have to worry about configuration anymore either (and it includes my util node helpers by default)..

Specifically, I have a dist/ and src/ folder where dist has my js generated code. I have another helper to symlink a generated run script that points to the dist/index.js to my \~/bin folder! The thing is, how do I keep all this organized and syncornized with my droplet server? I generally have a git \`bin\` folder for zsh scripts that I backup with git and sync to my local \~/bin/ folder and remote \~/bin folder (I'll just run a [sync.sh](https://sync.sh/) script each time I pull on either server). But because typescript has file dependencies, it's a multi-file process. I was thinking I could rsync each ts-helper-script-folder to the remote server, and run the same helper mentioned earlier to link the run script (pointing to dist/index.js) to the remote servers \~/bin folder, but then if I make changes remotely, I don't see how I'd synchronize it back to my local machine.

😅thanks for hearing out all my thoughts on this :). Basically I have .sh scripts already "easily" synced in \~/bin on both servers because I also keep the scripts together in the same git repo, but TS is a bit tricker because of the file dependencies and I'm not sure about also keeping them in my .hidden git repo given size.

What do you think?

&amp;#x200B;

Update this is some progress? [https://www.reddit.com/r/typescript/comments/iajbyp/how\_do\_you\_sync\_oneoff\_ts\_node\_scripts\_between/g1qa0vx?utm\_source=share&amp;utm\_medium=web2x](https://www.reddit.com/r/typescript/comments/iajbyp/how_do_you_sync_oneoff_ts_node_scripts_between/g1qa0vx?utm_source=share&amp;utm_medium=web2x)
## [6][How to launch a Typescript transformer plugin?](https://www.reddit.com/r/typescript/comments/iablmc/how_to_launch_a_typescript_transformer_plugin/)
- url: https://www.reddit.com/r/typescript/comments/iablmc/how_to_launch_a_typescript_transformer_plugin/
---
Hi folks! I have a really cool idea for a Typescript transformer for JSX syntax. Do you have any advice on how to open source it to the community so that others use it and give feedback?  


Thanks in advance!
## [7][For people who think changing types constitutes a breaking change, do you consider upgrading the Typescript version you compile your library with a breaking change?](https://www.reddit.com/r/typescript/comments/ia6zvy/for_people_who_think_changing_types_constitutes_a/)
- url: https://www.reddit.com/r/typescript/comments/ia6zvy/for_people_who_think_changing_types_constitutes_a/
---
This is branching of another post, but it seemed a lot of people thought a change to the types that may cause typings to fail consortia a breaking change. To follow up, I'm curious to know what people's thoughts is on the actually Typescript library.

I can see an argument being made that if I upgrade Typescript library then the way my types are generated can include types that only exist in a later version of typescript. Therefore, a consumer of the library who is on an older version of typescript would have an error and would need to upgrade Typescript.

On the other hand, it's usually not the case as per how npm works that you release a breaking change if an internal dev dependency library upgrades.

I personally don't think it should be considered a breaking change, but I would still love to hear more sides to this.
## [8][Typescript sometimes seems like more of a hassle than it's worth](https://www.reddit.com/r/typescript/comments/iae2zl/typescript_sometimes_seems_like_more_of_a_hassle/)
- url: https://www.reddit.com/r/typescript/comments/iae2zl/typescript_sometimes_seems_like_more_of_a_hassle/
---
I like typescript a lot, but only when I know I'm not going to need many libraries. Anything on the back end I handle with typescript. 

But typescript seems like a hassle for things like react. What happens when you want a package that has no types?
## [9][For libraries distributed with TypeScript types, is changing or removing a Type meant to be a breaking change for the library?](https://www.reddit.com/r/typescript/comments/i9rhrw/for_libraries_distributed_with_typescript_types/)
- url: https://www.reddit.com/r/typescript/comments/i9rhrw/for_libraries_distributed_with_typescript_types/
---
Wouldn't have any impact on js consumers, but could potentially break TypeScript apps?
## [10][Overload function can't find input properties](https://www.reddit.com/r/typescript/comments/i9nmdy/overload_function_cant_find_input_properties/)
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
## [11][I created a video to share 3 typing basics to improve your TypeScript code](https://www.reddit.com/r/typescript/comments/i98t3e/i_created_a_video_to_share_3_typing_basics_to/)
- url: https://www.youtube.com/watch?v=UuBJrAZsp4Y
---

