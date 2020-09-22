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
## [2][Convert class to interface during runtime](https://www.reddit.com/r/typescript/comments/ixngm0/convert_class_to_interface_during_runtime/)
- url: https://www.reddit.com/r/typescript/comments/ixngm0/convert_class_to_interface_during_runtime/
---
Hello everyone,
I'm using react/redux with typescript (functional) and integrating a 3rd party api client. The models defined in the 3rd party are class definitions, however the redux store complaints when trying to add class objects to it. Is there a way to create an interface that uses the field definitions of a class or turn the class definition into a native object during runtime? I want to avoid having to redefine the models as interfaces in order to prevent possible schema desync between my and the 3rd party library.
## [3][A SQL database implemented purely in TypeScript type annotations.](https://www.reddit.com/r/typescript/comments/iww4hs/a_sql_database_implemented_purely_in_typescript/)
- url: https://github.com/codemix/ts-sql
---

## [4][When are we likely to see updated Clipboard access from typescript?](https://www.reddit.com/r/typescript/comments/ix861c/when_are_we_likely_to_see_updated_clipboard/)
- url: https://www.reddit.com/r/typescript/comments/ix861c/when_are_we_likely_to_see_updated_clipboard/
---
Currently the Clipboard interface in lib.d.ts only has readText and writeText.  When are we likely to see read and write, given they are pretty widely supported now?

Trying to decide if I should try to hack around this or wait.

Thanks
## [5][Debug Visualizer A VS Code extension for visualizing data structures while debugging. Like the VS Code's watch view, but with rich visualizations of the watched value.](https://www.reddit.com/r/typescript/comments/iwpzbm/debug_visualizer_a_vs_code_extension_for/)
- url: https://marketplace.visualstudio.com/items?itemName=hediet.debug-visualizer&amp;1
---

## [6][Jest Testing: How to import react components when the whole react app is in a typescript namespace and using C# ScriptBundle](https://www.reddit.com/r/typescript/comments/ix2rmw/jest_testing_how_to_import_react_components_when/)
- url: https://www.reddit.com/r/typescript/comments/ix2rmw/jest_testing_how_to_import_react_components_when/
---
 I'm having my whole ts react project in a namespace and now I want to cover this project with tests but I can't figure out how to import props so, I can access desired components that I want to test or import 'render from testing-library/react' in my test. For Example, my component looks like this:  
**MyComponent.tsx**

    module MainModule {
    
        interface Props extends ReactTabComponentProps {
        }
    
        interface State extends ReactComponentState {
        }
    
        export class MyComponent extends ReactTabComponent&lt;Props, State&gt; {
            constructor(props: Props) {
                super(props);
    
                this.state = {
                };
            }
    
            setTitle = () =&gt; {
                this.props.tab.setTitle(
                    &lt;span&gt;
                        &lt;Icon name="calculator" fixedWidth /&gt;
                        MainComponent
                    &lt;/span&gt;
                );
            }
    
            render() {
                return (
                    
                );
            }
        }
    }


My test file looks like this:  
**MyComponent.test.tsx**

    module MainModule {
    
     import { MainComponent } from '../Main/components/MainComponent';
     import { render } from '@testing-library/react';
     
        describe("mainComponentTest", () =&gt; {
            it("should renders correctly", async () =&gt; { 
               
            });
    
    
        });
    }

 With this, I get TS errors, Error: `(TS)Import declarations in a namespace cannot reference a module.`   
 And if I remove module from my test or write imports outside of module then I get the error:   
 

`(TS)Cannot compile modules using option 'outFile' unless the '--module'flag is 'amd' or 'system'.`  
Any help would be highly appreciated!
## [7][Safe Lookup Table in typescript](https://www.reddit.com/r/typescript/comments/iwsfwl/safe_lookup_table_in_typescript/)
- url: https://www.reddit.com/r/typescript/comments/iwsfwl/safe_lookup_table_in_typescript/
---
Is it possible to create a lookup table in typescript with functions that accept different type arguments?
I'm have a function that returns  a discriminated union with a type property that determines the type of data it holds, I would like to then use the type property as a key to an object with functions that accept types that the discriminated union was formed from
For example

    type a = {
        type:"obj1";
        foo:string;
    }
    type b={
        type:"obj2";
        bar:number;
    }
    type union= a or b;
    const lookupTable={
        obj1:({foo}:a)=&gt;foo.toLowerCase();
        obj2:({bar}:b)=&gt;(bar**2).toString(); // all the functions return the same type
    }
    //That all works, my issue comes when attempting to do something like
    const unionObj:union = {type:"obj1", foo:"str"}; // my code returns this from a function
    const result = lookupTable[unionObj.type](unionObj);
    // complicated type errors cause it can't tell that unionObj has been narrowed down to the right type

Can anyone please help with this
## [8][Compiler - Extract the type of union type tuple](https://www.reddit.com/r/typescript/comments/iwkt0i/compiler_extract_the_type_of_union_type_tuple/)
- url: https://www.reddit.com/r/typescript/comments/iwkt0i/compiler_extract_the_type_of_union_type_tuple/
---
If I have an array of type `string[]` and I want to extract the type from the array, I must do: 

    const typeArgs = type.typeArguments();
    typeArgs[0]; // string

But what if I have a tuple of type \[number, string\] and want to get the type union of `string | number`?
## [9][Can we use TypeScript in place of JavaScript?](https://www.reddit.com/r/typescript/comments/iw8ukw/can_we_use_typescript_in_place_of_javascript/)
- url: https://www.reddit.com/r/typescript/comments/iw8ukw/can_we_use_typescript_in_place_of_javascript/
---
I have a confession to make, I am one of those guys who tried to learn JavaScript many times but failed miserably despite having coding skills in many other languages. Recently I heard about TypeScript and really liked its way of programming. Can I ask if we can replace JS and in its place use TS.

Lets say I have to built a simple calculator with HTML+JS+CSS only. Could I use TS to write the code,  then compile it to JS and use this for my above calculator project? Will TS code be able to interact with the DOM elements like native JS?
## [10][How can I type a decorator ?](https://www.reddit.com/r/typescript/comments/iwbbq7/how_can_i_type_a_decorator/)
- url: https://www.reddit.com/r/typescript/comments/iwbbq7/how_can_i_type_a_decorator/
---
[playground link of what I have done so far.](https://www.typescriptlang.org/play?#code/GYVwdgxgLglg9mABAQ2vMAeAKogpgDylzABMBnRAMXDQQD4AKYMALkSwEoWcBvAKACQAJ1xQQQpAwB0M5EIDmZDogC8dRPwECA9NrJwAtrkRkxwYIIC+Abj42gA)
## [11][I created a Web Development Discord server for TypeScript!](https://www.reddit.com/r/typescript/comments/iwk6m9/i_created_a_web_development_discord_server_for/)
- url: https://www.reddit.com/r/typescript/comments/iwk6m9/i_created_a_web_development_discord_server_for/
---
Wonderful news!

&amp;#x200B;

I created my first Discord server and I'm so happy for what I've made. It's called "The Web Dev Heaven" it's a friendly community where web developers and designers are welcomed to talk about the latest technologies and languages. So we can all learn about different ( Databases, Frameworks, Adobe XD plugins, Algorithms, CMS's, etc. ) to improve our code. We are here for supporting our projects ( So if you are looking to gain traffic with your website or feedback on your project, we are here for you! ).

&amp;#x200B;

[https://discord.gg/uXE4E7n](https://discord.gg/uXE4E7n)

&amp;#x200B;

Thank you, if you join!  
Thank you, for the support!  
Thank you, for the upvote!
