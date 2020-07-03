# typescript
## [1][Who's hiring Typescript developers - July](https://www.reddit.com/r/typescript/comments/hizg5z/whos_hiring_typescript_developers_july/)
- url: https://www.reddit.com/r/typescript/comments/hizg5z/whos_hiring_typescript_developers_july/
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
## [2][Valuable compiler options](https://www.reddit.com/r/typescript/comments/hkhdur/valuable_compiler_options/)
- url: https://www.reddit.com/r/typescript/comments/hkhdur/valuable_compiler_options/
---
After forgetting `strictNullChecks` I've forgotten how helpful it is to constantly have to account for `null | undefined` during my initial drafting phase.

Any others you guys would promote for most projects? I also have `noImplicitAny`, and of course the `esModuleInterop`
## [3][TS2464: A computed property name must be of type 'string', 'number', 'symbol', or 'any'. 54 [x.y] : z](https://www.reddit.com/r/typescript/comments/hkiodn/ts2464_a_computed_property_name_must_be_of_type/)
- url: https://www.reddit.com/r/typescript/comments/hkiodn/ts2464_a_computed_property_name_must_be_of_type/
---
    // TS2464: A computed property name must be of type 'string', 'number', 'symbol', or 'any'.
    
    54  [currentStepInstance.configDataKey] : validatedUserInput

I'm trying to use -- I think it's called "scope resolution" -- here on the key name, since property access requires dot notation. The key name is a string. 

The below syntax didn't work, how should it be written?

             if (currentStepInstance.hasSaveableData) {
                Object.assign(returnObject, {
                   [currentStepInstance.configDataKey: string] : validatedUserInput
                });
             }
## [4][How do I de-structure express request.body without "Unsafe assignment of an any value." error?](https://www.reddit.com/r/typescript/comments/hk1mp8/how_do_i_destructure_express_requestbody_without/)
- url: https://www.reddit.com/r/typescript/comments/hk1mp8/how_do_i_destructure_express_requestbody_without/
---
I am very new to typescript and when vscode threw that error after doing some search I mixed few stuffs from the results and came up with this solution.

`interface Body {`  
   `name: string;`  
   `email: string;`  
   `password: string;`  
 `}`  
 `const body: Body = request.body as Body;`  
 `const { name, email, password } = body;`

How can I improve this as it seems a lot of code for de-structuring.
## [5][How to deal with array of different class instances in typescript](https://www.reddit.com/r/typescript/comments/hk2z2a/how_to_deal_with_array_of_different_class/)
- url: https://www.reddit.com/r/typescript/comments/hk2z2a/how_to_deal_with_array_of_different_class/
---
So my code looks like this:

An abstract class named Component with nothing inside which all component inherit from.

An abstract generic base Entity class functioning as an Array abstraction with methods to add/get instances of classes,

A Entity class which extends BaseEntity with all the getters, one for each component, the reason I created it is because I must always force the type like &lt;Health&gt; this.getComponent(Health), I want to get rid of it, is there a way to type check BaseEntity so that when I do  this.getComponent(Health), the compiler knows its Health instance?

&amp;#x200B;

    export class Component {}
    export class Health extends Component { ............}
    export class Position extends Component { ............. }
    
    .....................
    export abstract class BaseEntity&lt;R&gt; {
        private readonly entity: Array&lt;R&gt; = [];
    
        protected getComponent&lt;T extends { new(...args: any[]): R }&gt;(component: T): R {
            return this.entity.find(
                currentComponent =&gt; currentComponent instanceof component
            );     
        } 
    
        addComponent(component: R): R {
            const index = this.entity.indexOf(component)
    
            if (index === -1) {
                 this.entity.push(component)
    
                 return component
            }
        }
    }
    
    .....................
    export abstract class Entity extends BaseEntity&lt;Component | Entity&gt; {
        get getHealth(): Health {
            return &lt;Health&gt; this.getComponent(Health)
        }
    
        get getPosition(): Position {
            return &lt;Position&gt; this.getComponent(Position)
        }
    }
    export class Player extends Entity {}
    export class Monster extends Entity {}
## [6][Type 'string' is not assignable to type '"X" | "Y" | "Z"'](https://www.reddit.com/r/typescript/comments/hjvw3n/type_string_is_not_assignable_to_type_x_y_z/)
- url: https://www.reddit.com/r/typescript/comments/hjvw3n/type_string_is_not_assignable_to_type_x_y_z/
---
Recently I had to use the `ReturnType` generic type constructor to build a type for me, that was the return shape of an obsurely documented method in Google Text To Speech's API:

    protected async fetchAndWriteAudio(options: {
          input : { text : string }
          , voice : { languageCode : string, ssmlGender: voiceGender }
          , audioConfig : { 
             audioEncoding : "AUDIO_ENCODING_UNSPECIFIED" | "LINEAR16" | "MP3" | "OGG_OPUS" }
       }, fileNameAndPath: string)
          : Promise&lt;ReturnType&lt;typeof TextToSpeechClient.prototype.synthesizeSpeech&gt;&gt; {

But now the title error shows on the error-commented lines below. I thought a specific string value, depicted as a quad union from the `ReturnType` constructor, should match one of those specific strings. Instead, none match. Why is that and what is the proper fix?

       public async makeSentenceAudio(subFolderPath: string): Promise&lt;void&gt; {
          // setup the foreign and english text string
          const foreignSentenceText = this.textTranslate(
             this.sentence.englishVersion
             , translationDirection.toForeign);
          const englishText = this.sentence.englishVersion;
    
          
          /**** Generate Audios ****/
          // define audio options (including text)
          const sharedOptions = {
             input: { text : "" }
             , voice: { languageCode : "", ssmlGender : voiceGender.male }
             , audioConfig : {audioEncoding : "OGG_OPUS"}
          }
    
          const foreignAudioOptions = {
             ...sharedOptions
             , input: {text: ""}
          }
    
          const englishAudioOptions = {
             ...sharedOptions
             , input: {text: ""}
          }
    
          this.fetchAndWriteAudio(foreignAudioOptions, "placeholder");
          this.fetchAndWriteAudio(englishAudioOptions, "placeholder");

.

    // the error message. There was also one on englishAudioOptions
    
    const foreignAudioOptions: {
        input: {
            text: string;
        };
        voice: {
            languageCode: string;
            ssmlGender: voiceGender;
        };
        audioConfig: {
            audioEncoding: string;
        };
    }
    Argument of type '{ input: { text: string; }; voice: { languageCode: string; ssmlGender: voiceGender; }; audioConfig: { audioEncoding: string; }; }' is not assignable to parameter of type '{ input: { text: string; }; voice: { languageCode: string; ssmlGender: voiceGender; }; audioConfig: { audioEncoding: "OGG_OPUS" | "AUDIO_ENCODING_UNSPECIFIED" | "LINEAR16" | "MP3"; }; }'.
    
      The types of 'audioConfig.audioEncoding' are incompatible between these types.
        Type 'string' is not assignable to type '"OGG_OPUS" | "AUDIO_ENCODING_UNSPECIFIED" | "LINEAR16" | "MP3"'. ts(2345)
## [7][epubjs library](https://www.reddit.com/r/typescript/comments/hjrtwl/epubjs_library/)
- url: https://www.reddit.com/r/typescript/comments/hjrtwl/epubjs_library/
---
I am new with typescript and I had before .net / C# experience, nowadays I am dealing with typescript epubjs library. I developed epub-reader component and it is working as it should work, but there are cases with some epub files, where first page of book is always aligned on left side, when normally is on center. Those alignments are setup on popup window where you can choose "spread to one page" OR "spread to two pages"

&amp;#x200B;

dialog example : 

    &lt;mat-radio-group [(ngModel)]="epubTextOptions.selectedTwoPages" class="inline-button"&gt; &lt;mat-radio-button [value]="false" [checked]="!epubTextOptions.selectedTwoPages" class="inline-button"&gt;{{'reader.textOptions.dialog.onePage' | translate}}&lt;/mat-radio-button&gt; &lt;mat-radio-button [value]="true" [checked]="epubTextOptions.selectedTwoPages" class="inline-button"&gt;{{'reader.textOptions.dialog.twoPages' | translate}}&lt;/mat-radio-button&gt; &lt;/mat-radio-group&gt;

 typescript example:  

    export class MyClass implements OnInit {   book: Book;   rendition: Rendition; private myMethod() { this.rendition = this.book.renderTo(this.ePub.nativeElement, {       width: '100%',       height: '100%' }); this.rendition.spread(this.epubTextOptions.selectedTwoPages ? 'always' : 'none'); } }

 html component : 

    &lt;div id="epubContent" class="d-flex align-items-center" [style.height]="contentHeight + 'px'"&gt; &lt;app-button *ngIf="displayPrevPageButton() &amp;&amp; !isMobileDevice()" (btnClick)="prevPage()"               btnClass="mat-icon-button epub-action"&gt;&lt;span class="far fa-chevron-left"&gt;&lt;/span&gt;&lt;/app-button&gt; &lt;div #ePub class="epub-container"&gt;&lt;/div&gt; &lt;app-button *ngIf="displayNextPageButton() &amp;&amp; !isMobileDevice()" (btnClick)="nextPage()"               btnClass="mat-icon-button epub-action"&gt;&lt;span class="far fa-chevron-right"&gt;&lt;/span&gt;&lt;/app-button&gt; &lt;/div&gt;

 Everything is implemented from epubjs library examples and adapted to reader, but as I said with most of epub spread() functionality is working, but there are epub files, where first page align to left side when "one page" is selected instead of center align. Maybe some how I can check if file have no more than one page and if yes, then force to align to center?
## [8][TIL What TypeScript Is](https://www.reddit.com/r/typescript/comments/hj8azp/til_what_typescript_is/)
- url: https://www.reddit.com/r/typescript/comments/hj8azp/til_what_typescript_is/
---
So last night, I finally watched a short intro on what exactly typescript is, and... Wow.

I wish I had known about this years ago; as someone who loves languages like C++, and never really was a fan of JS, TS makes the web world so much more inviting.

Anyway, happy to be here, sad it took me so long to learn about it.
## [9][How often do you guys use the NotNullable utility type constructor?](https://www.reddit.com/r/typescript/comments/hjssdd/how_often_do_you_guys_use_the_notnullable_utility/)
- url: https://www.reddit.com/r/typescript/comments/hjssdd/how_often_do_you_guys_use_the_notnullable_utility/
---
Seems especially useful for kicking undesired null or undefined values out of an expected API response type.

For my own defined types, I thought it was not needed since they are already sharply defined. But this video says TS allows any type definition to be overridden by null/undefined by default, which surprised me: [https://www.youtube.com/watch?v=Fgcu\_iB2X04?m=16&amp;s=20](https://www.youtube.com/watch?v=Fgcu_iB2X04?m=16&amp;s=20)
## [10][Checking for interface collisions .](https://www.reddit.com/r/typescript/comments/hjif2w/checking_for_interface_collisions/)
- url: https://www.reddit.com/r/typescript/comments/hjif2w/checking_for_interface_collisions/
---
Imagine having multiple different interfaces that all together are mixed into a big interface .

Before some time I [asked](https://www.reddit.com/r/typescript/comments/gab4ic/how_to_make_ts_or_eslint_or_vscode_warn_me_about/) how can I check for collisions when I merge two interfaces and I got an answer that worked . Here is a more simple version of it :

    type isNever&lt;T extends never&gt; = T;
    //I just look if I get a linting error in the following type 
    //If yes then there is a collision
    type mustBeNever = isNever&lt;keyof interface1 &amp; keyof interface2&gt;;

Now I have to merge more than two interfaces . With that way , for n interfaces I need to do `n!/(2!(n-2)!) = n(n-1)/2` manual checks (if you are interested in how that number is calculated take a look [here](https://en.wikipedia.org/wiki/Combination)) , something that is not practical .

Is there any solution via using a type function like I have already done that is not unpractical or should I create a unit test ?
## [11][How to ensure the right function is called with the right arguments?](https://www.reddit.com/r/typescript/comments/hjlrr1/how_to_ensure_the_right_function_is_called_with/)
- url: https://www.reddit.com/r/typescript/comments/hjlrr1/how_to_ensure_the_right_function_is_called_with/
---
Suppose I've got a login and a logout function:

    login(creds: { user: string, pass: string }): string;
    logout(token: string): void;

And a handler decides which function to call depending on a variable:

    let f = command === 'login' ? login : logout;
    let x = command === 'login' ? creds : token;
    return f(x);

I know that if f is logout, then x is going to be a string. Likewise if f is login, then x is going to be a credentials object. Can you help to model that?

Thanks!
