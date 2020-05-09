# typescript
## [1][Who's hiring Typescript developers - May](https://www.reddit.com/r/typescript/comments/gb7km3/whos_hiring_typescript_developers_may/)
- url: https://www.reddit.com/r/typescript/comments/gb7km3/whos_hiring_typescript_developers_may/
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
## [2][Prepare your Angular interview with this quick and complete outline](https://www.reddit.com/r/typescript/comments/ggecix/prepare_your_angular_interview_with_this_quick/)
- url: https://medium.com//prepare-your-angular-interview-with-this-quick-and-complete-outline-1e9b5d761166?source=friends_link&amp;sk=8a4d78eadeeae8aa9b10a0f9e7c47f29
---

## [3][How can I limit the strings used to index into an object?](https://www.reddit.com/r/typescript/comments/gg1r6z/how_can_i_limit_the_strings_used_to_index_into_an/)
- url: https://www.reddit.com/r/typescript/comments/gg1r6z/how_can_i_limit_the_strings_used_to_index_into_an/
---
Say I have a Movie type:

    interface Movie {
      title: string;
      year: number;
      studio: string;
      actors: string[];
      writers: string[];
    }
    
    const m: Movie = {
      title: "Silence of the Lambs",
      year: 1991,
      studio: "Strong Heart",
      actors: ["Jodie Foster", "Anthony Hopkins", "Ted Levine"],
      writers: ["Ted Tally"],
    };

and there are a few properties I want to iterate over, like

    for (let jobName of ["actors", "writers"]) {
      m[jobName].forEach(console.log);
    }

I can't do this unless I add  `[index: string]: string[] | string | number` to my interface and then assert `m[jobName] as string`, but then I can use *any* string to index in, and I don't want that.

In this case I could just do

    m.actors.forEach(console.log);
    m.writers.forEach(console.log);

but what if I'm not iterating over just two properties, what if I'm iterating over 50 and don't want pages and pages of manually specified iterations? This isn't exactly a real problem I'm having but I'd like to know what the correct response to this problem would be, because it's definitely something where I feel like I don't know the smart path.
## [4][StackOfTheFuture.com - my new blog about TypeScript and GraphQL](https://www.reddit.com/r/typescript/comments/gftk35/stackofthefuturecom_my_new_blog_about_typescript/)
- url: https://stackofthefuture.com/
---

## [5][I need some help with my code here](https://www.reddit.com/r/typescript/comments/ggcvi3/i_need_some_help_with_my_code_here/)
- url: https://www.reddit.com/r/typescript/comments/ggcvi3/i_need_some_help_with_my_code_here/
---
So I don't know what the issue is, but Visual Studio Code is being retarded and telling me there are errors and then proceeding to tell me that there are no errors. So there's that. I would see what the errors are myself, but I cannot. Anyways, that besides the point. I need help with two things. I need help with the errors I have. The second thing I need help with is how to have another "let path = document.location.pathname;  
 title = document.querySelector("#maincontentWrapper &gt; div &gt; div:nth-child(2) &gt; section &gt; div &gt; article &gt; div &gt; h1 &gt; span")" as I have one, but I make another, it underlines "let" and path. 

&amp;#x200B;

 Code: 

    var presence = new Presence({
      clientId: "707804267146117120"
    });
    var browsingStamp = Math.floor(Date.now() / 1000);
    var title;
    presence.on("UpdateData", async () =&gt; { // You had it like {} and forgot to put the code used in those two brackets
      const presenceData: presenceData = {
          largeImageKey: "logo"
      };
      if (document.location.pathname == "/") {
        presenceData.startTimestamp = browsingStamp;
        presenceData.details = ("Viewing the Rovio Entertainment Homepage");
      }
      let path = document.location.pathname;
      title = document.querySelector("#main &gt; div.wrap &gt; div.main-wrap &gt; article &gt; header &gt; h1");
      let path = document.location.pathname;
      title = document.querySelector("#maincontentWrapper &gt; div &gt; div:nth-child(2) &gt; section &gt; div &gt; article &gt; div &gt; h1 &gt; span")
      
      if (path.includes("/news/") &amp;&amp; title) {
          presenceData.startTimestamp = browsingStamp;
          presenceData.details = ("Looking at some news from Rovio");
      }
      else if (document.location.pathname.includes ("/games")){
        presenceData.startTimestamp = browsingStamp;
        presenceData.details = ("Looking at the games Rovio make");
      }
      else if (document.location.pathname.includes ("/games/angry-birds")){
        presenceData.startTimestamp = browsingStamp;
        presenceData.details = ("Looking at the Angry Birds games Rovio makes ");
      }
      else if (document.location.pathname.includes ("/sugar-blast")){
        presenceData.startTimestamp = browsingStamp;
        presenceData.details = ("Looking at the Rovio made game Sugar Blast...");
      }
      else if (document.location.pathname.includes ("/more")){
        presenceData.startTimestamp = browsingStamp;
        presenceData.details = ("Viewing the other crappy games they have...");
      }
    //Their employee stories
    else if (document.location.pathname.includes ("/careers/employee-stories")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Viewing the avaiable employee stories");
    }
    else if (document.location.pathname.includes ("/careers/employee-story/laura-rokkanen")){
      presenceData.startTimestamp = browsingStamp;
    presenceData.details = ("Viewing Laura Rokkanen's story")
    }
    else if (document.location.pathname.includes("/careers/employee-story/nim-frydman")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Viewing Nim Frydman's story")
    }
    
    else if (document.location.pathname.includes ("/careers/employee-story/di-zhong")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Viewing Di Zhong's story")
    }
    else if (document.location.pathname.includes ("/careers/employee-story/ignacio-amaya")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Viewing Ignacio Amaya's story")
    }
    else if (document.location.pathname.includes ("/careers/employee-story/jesse-lempiainen")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Viewing Jesse Lempiäinen's story")
    }
    else if (document.location.pathname.includes("/careers/employee-story/jenna-linden")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Viewing Jenna Lindén's story")
    }
    else if (document.location.pathname.includes ("/careers/employee-story/matteo-spiri")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Viewing Matto Spiri's story")
    }
    else if (document.location.pathname.includes ("/careers/employee-story/marc-olivier")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Viewing Marc Olivier's story")
    }
    else if (document.location.pathname.includes ("/careers/employee-story/nicole-yang")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Viewing Nicole Yang's story")  
    }
    
    //Main career page. I forgot to add it before all of the stories. Oh fucking well...
    else if (document.location.pathname.includes ("/careers")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Viewing the careers at Rovio")
    }
    
    //Rovio's teams stories
    
    else if (document.location.pathname.includes ("our-teams")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Viewing the avaiable stories about Rovio's teams")
    }
    else if (document.location.pathname.includes ("/careers/our-teams/puzzle-studio")){
    presenceData.startTimestamp = browsingStamp;
    presenceData.details = ("Viewing Rovio's Puzzle Studio story")
    }
    else if (document.location.pathname.includes("/careers/our-teams/stockholm-studio")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Viewing Rovio's Stockholm Studio's story")
    }
    else if (document.location.pathname.includes("/careers/our-teams/games-technology")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Viewing Rovio's Games Technology's story")
    }
    else if (document.location.pathname.includes ("/careers/our-teams/delta")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Viewing Delta's story")
    }
    else if (document.location.pathname.includes ("/careers/our-teams/battle-studio")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Viewing Battle Studio's story")  
    }
    //Trainee shit...
    else if (document.location.pathname.includes ("/careers/trainee-faq")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Viewing Rovio's Frequently Asked Questions for trainees")
    }
    //Rovio loves fucking me over. So, they have a tab called "Investors..." It's a whole other fucking subdomain. It's "Investors.Rovio.com" Fucking kill me now... 
    else if (document.location.pathname.includes ("invesstors.rovio.com")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Viewing Rovio's homepage for Investors")
    }
    else if (document.location.pathname.includes ("investors.rovio.com/en/about-rovio/who-we-are#Yeartab")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Seeing who Rovio is")
    }
    else if (document.location.pathname.includes ("investors.rovio.com/en/about-rovio/why-invest-rovio#Yeartab")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Seeing why they should invest in Rovio")
    }
    else if (document.location.pathname.includes ("investors.rovio.com/en/about-rovio/our-strategy#Yeartab")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Seeing Rovio's stragety")
    }
    else if (document.location.pathname.includes ("investors.rovio.com/en/about-rovio/values-and-culture#Yeartab")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Seeing Rovio's values and culture")
    }
    else if (document.location.pathname.includes ("investors.rovio.com/en/about-rovio/leadership-team#Yeartab")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Seeing Rovio's leadership team")
    }
    else if (document.location.pathname.includes ("investors.rovio.com/en/about-rovio/investor-relations#Yeartab")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Seeing Rovio's relationship with investors")
    }
    else if (document.location.pathname.includes ("investors.rovio.com/en/our-business")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Seeing what Rovio's business is about")
    }
    else if (document.location.pathname.includes ("https://investors.rovio.com/en/our-business/brand-licensing")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Seeing Rovio's licensing for their brands (Angry Birds mostly...)")
    }
    else if (document.location.pathname.includes ("investors.rovio.com/en/financials/key-financials#Yeartab")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Seeing Rovio's financials")
    }
    else if (document.location.pathname.includes ("https://investors.rovio.com/en/financials/key-financials#Outlook-guidance")){
      presenceData.startTimestamp = browsingStamp;
      presenceData.details = ("Looking at Rovio's outlook guidance")
    }
    else if (document.location.pathname.includes ("https://investors.rovio.com/en/financials/interim-reports/year/2019#Yeartab")){
      presenceData.startTimestamp = browsingStamp; 
      presenceData.details =  ("Looking at Rovio's interim reports")  
    }
    else if (document.location.pathname.includes ("https://investors.rovio.com/en/financials/reports-and-presentations/year/2020#Yeartab")){
      presenceData.startTimestamp = browsingStamp; 
      presenceData.details = ("Looking at Rovio's reports and presentations")
    }
    })

Any and all help is appreciated! 

Thank you in advanced!

Red Prez
## [6][Passing a type value 'upwards'...?](https://www.reddit.com/r/typescript/comments/gg1gm6/passing_a_type_value_upwards/)
- url: https://www.reddit.com/r/typescript/comments/gg1gm6/passing_a_type_value_upwards/
---
So obviously this needs some explaining but I hope you guys can understand it. I'm writing types declaration files for various npm libraries, and one in particular I had trouble with yesterday is a Vue component. It is passed as an argument into the function `Vue.use`, which is defined in Vue's own types declaration file as the following:

     use&lt;T&gt;(plugin: PluginObject&lt;T&gt; | PluginFunction&lt;T&gt;, options?: T): VueConstructor&lt;V&gt;;

The component I was writing a declaration file for yesterday had its own set of options which I had been intending to pass as a type for the second argument of `Vue.use`, but I *could not* have the options object recognised by `Vue.use`, no matter what I did (unless I pass in the type manually with `Vue.use&lt;{ ... }&gt;`, but that obviously is not ideal).

My main approach was to add a type variable in the final exported function of the component I was attempting to type, so like this:

    export default function fooVueComponent&lt;T = fooVueComponentOptions&gt;(): new () =&gt; fooVueComponentClass&lt;T&gt;;

so that the `T` from `fooVueComponent` would take the place of the `T` in `Vue.use` in order to utilise the type as the second argument in `Vue.use`. But passing the type value 'upwards' like so did not come to fruition, as `Vue.use` still did not recognise the `fooVueComponentOptions` type. Is there anything I can do to pass the type up to `Vue.use` or is it futile? Thanks in advance to any respondents and if you guys need anything cleared up, just let me know.

EDIT: A simple tl;dr: I want the `T` from `fooVueComponent`, which will be the first argument in `Vue.use`, to be the `T` (and therefore the second argument) in `Vue.use`.
## [7][Need to read through a JSON file](https://www.reddit.com/r/typescript/comments/gg35j5/need_to_read_through_a_json_file/)
- url: https://www.reddit.com/r/typescript/comments/gg35j5/need_to_read_through_a_json_file/
---
I need to know to to read through a JSON file so I can import each item into a mongoDB collection.

I've tried a few tutorials and articles and I just can't seem to get it to work.

Anyone have a good link?
## [8][Convert parameters to destructured object](https://www.reddit.com/r/typescript/comments/gfre6t/convert_parameters_to_destructured_object/)
- url: https://www.reddit.com/r/typescript/comments/gfre6t/convert_parameters_to_destructured_object/
---
I just noticed this refactoring was available in one of my methods. Is this a refactor I should do? And what really does it do? Tried to read about it but I can't quite get it. What's the benefit of it?    

Here's the method:    

    private methodName(exampleString: string, array: any) {
        const json = JSON.parse(exampleString);
        json.map((value: any) =&gt; {
            array.map((v: { Type: any; Value: boolean; }) =&gt; {
                if (v.Type === value) {
                    v.Value = true;
                }
            });
        });
    }

If I use the refactoring it looks like this:    

    private methodName({ exampleString, array }: { exampleString: string; array: any; }) {
        const json = JSON.parse(exampleString);
        json.map((value: any) =&gt; {
            array.map((v: { Type: any; Value: boolean; }) =&gt; {
                if (v.Type === value) {
                    v.Value = true;
                }
            });
        });
    }

What's the benefit?
## [9][​​Setting Up Webpack and Typescript - TypeScript Webpack example](https://www.reddit.com/r/typescript/comments/gft88l/setting_up_webpack_and_typescript_typescript/)
- url: https://codesource.io/setting-up-webpack-and-typescript/
---

## [10][Publishing TypeScript NPM Libraries using NX and Github Actions](https://www.reddit.com/r/typescript/comments/gfejco/publishing_typescript_npm_libraries_using_nx_and/)
- url: https://tane.dev/2020/05/publishing-npm-libraries-using-nx-and-github-actions/
---

## [11][Front-end Software Engineer opening @Romania (Iasi)](https://www.reddit.com/r/typescript/comments/gfr66i/frontend_software_engineer_opening_romania_iasi/)
- url: https://www.reddit.com/r/typescript/comments/gfr66i/frontend_software_engineer_opening_romania_iasi/
---
Hi there! :)

I am currently looking for great Front-end Software Engineers for our new project in Iasi (Romania).

Our aim is to build Single Page Applications that drive decision making, with big care about writing high-quality code and providing excellent user experience in our product.

Key required skills: #JavaScript, #ReactJS, #TypeScript, #Redux &amp; familiarity with #CI/CD tools.

Want to discuss more? Just leave me a message! :)
