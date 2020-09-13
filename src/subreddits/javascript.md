# javascript
## [1][Showoff Saturday (September 12, 2020)](https://www.reddit.com/r/javascript/comments/irb676/showoff_saturday_september_12_2020/)
- url: https://www.reddit.com/r/javascript/comments/irb676/showoff_saturday_september_12_2020/
---
Did you find or create something cool this week in javascript? 

Show us here!
## [2][Introducing Rome: A linter for JavaScript and TypeScript.](https://www.reddit.com/r/javascript/comments/iru0l8/introducing_rome_a_linter_for_javascript_and/)
- url: https://romefrontend.dev/blog/2020/08/08/introducing-rome.html
---

## [3][I built a PWA to open WhatsApp chat without saving their contact first (it is painful for us to save contact just so that we can send a WhatsApp to another person)](https://www.reddit.com/r/javascript/comments/irrez7/i_built_a_pwa_to_open_whatsapp_chat_without/)
- url: http://whatsapp-anyone.netlify.com
---

## [4][[AskJS] What classless library/repo's code you like because of its clean and readable code?](https://www.reddit.com/r/javascript/comments/irly2m/askjs_what_classless_libraryrepos_code_you_like/)
- url: https://www.reddit.com/r/javascript/comments/irly2m/askjs_what_classless_libraryrepos_code_you_like/
---
I have never been a fan of classes and some other OOP concepts. I am trying to find the right balance between FP and OOP. I'm an expert at none of them :)

It is hard to find good examples or real applications of JS, as it's a very flexible language and easy to make a mess with it. What are good examples of that I can read? (no huge libraries if possible)
## [5][How To Build A Twitter Clone With NestJS, Prisma And React ( Part 1 )](https://www.reddit.com/r/javascript/comments/irvu03/how_to_build_a_twitter_clone_with_nestjs_prisma/)
- url: https://dev.to/darioielardi/how-to-build-a-twitter-clone-with-nestjs-prisma-and-react-part-1-lgd
---

## [6][A little bit of plain Javascript can do a lot](https://www.reddit.com/r/javascript/comments/irxy37/a_little_bit_of_plain_javascript_can_do_a_lot/)
- url: https://jvns.ca/blog/2020/06/19/a-little-bit-of-plain-javascript-can-do-a-lot/
---

## [7][CyberCode Online; I made a discreet MMORPG game that disguises as VS Code](https://www.reddit.com/r/javascript/comments/irg78z/cybercode_online_i_made_a_discreet_mmorpg_game/)
- url: https://cybercodeonline.com/
---

## [8][[AskJS] How would you go about creating javascript code into ASCII art?](https://www.reddit.com/r/javascript/comments/irv263/askjs_how_would_you_go_about_creating_javascript/)
- url: https://www.reddit.com/r/javascript/comments/irv263/askjs_how_would_you_go_about_creating_javascript/
---
For example i have seen this been done in C

    #  include&lt;stdio.h&gt;//  .IOCCC                                         Fluid-  #
    #  include &lt;unistd.h&gt;  //2012                                         _Sim!_  #
    #  include&lt;complex.h&gt;  //||||                     ,____.              IOCCC-  #
    #  define              h for(                     x=011;              2012/*  #
    #  */-1&gt;x              ++;)b[                     x]//-'              winner  #
    #  define              f(p,e)                                         for(/*  #
    #  */p=a;              e,p&lt;r;                                        p+=5)//  #
    #  define              z(e,i)                                        f(p,p/*  #
    ## */[i]=e)f(q,w=cabs  (d=*p-  *q)/2-     1)if(0  &lt;(x=1-      w))p[i]+=w*/// ##
       double complex a [  97687]  ,*p,*q     ,*r=a,  w=0,d;    int x,y;char b/* ##
    ## */[6856]="\x1b[2J"  "\x1b"  "[1;1H     ", *o=  b, *t;   int main   (){/** ##
    ## */for(              ;0&lt;(x=  getc (     stdin)  );)w=x  &gt;10?32&lt;     x?4[/* ##
    ## */*r++              =w,r]=  w+1,*r     =r[5]=  x==35,  r+=9:0      ,w-I/* ##
    ## */:(x=              w+2);;  for(;;     puts(o  ),o=b+  4){z(p      [1]*/* ##
    ## */9,2)              w;z(G,  3)(d*(     3-p[2]  -q[2])  *P+p[4      ]*V-/* ##
    ## */q[4]              *V)/p[  2];h=0     ;f(p,(  t=b+10  +(x=*p      *I)+/* ##
    ## */80*(              y=*p/2  ),*p+=p    [4]+=p  [3]/10  *!p[1])     )x=0/* ##
    ## */ &lt;=x              &amp;&amp;x&lt;79   &amp;&amp;0&lt;=y&amp;&amp;y&lt;23?1[1  [*t|=8   ,t]|=4,t+=80]=1/* ##
    ## */, *t              |=2:0;    h=" '`-.|//,\\"  "|\\_"    "\\/\x23\n"[x/** ##
    ## */%80-              9?x[b]      :16];;usleep(  12321)      ;}return 0;}/* ##
    ####                                                                       ####
    ###############################################################################
    **###########################################################################*/Y

And im pretty sure you can do it in JavaScript. i tried using figlet to generate ASCII art then replacing the characters but it just doesn't seem to be working. i found a 5 year old project which is in clojure but cant get it running too ([https://github.com/asivokon/harnocode](https://github.com/asivokon/harnocode)) This is more of a puzzle kinda thing :P
## [9][[AskJS] A huge problem with .offsetTop](https://www.reddit.com/r/javascript/comments/irx0mx/askjs_a_huge_problem_with_offsettop/)
- url: https://www.reddit.com/r/javascript/comments/irx0mx/askjs_a_huge_problem_with_offsettop/
---
So I know that without a positioned ancestor, .offsetTop is measured from the top of the &lt;body&gt; element. 

I am trying to to slide in elements when they are half shown with JavaScript. 

So the Javascript code calculates how far the image is from the top of the window and if the image is half shown, a CSS class is applied to the element.

But, since the parent element of the image (footer) is set to position:relative the offsetTop doesn't work anymore.

The code looks like this

    const slideInAt = (window.scrollY + window.innerHeight) - sliderImage.height / 2;
    const imageBottom = sliderImage.offsetTop + sliderImage.height;   
    const isHalfShown = slideInAt &gt; sliderImage.offsetTop;   
    const isNotScrolledPast = window.scrollY &lt; imageBottom;  

    if (isHalfShown &amp;&amp; isNotScrolledPast) {
      sliderImage.classList.add('active');
    } else {
      sliderImage.classList.remove('active');
    }

Here is the real life example: https://bluishbanana.github.io/

Any ideas how to solve that?
## [10][[AskJS] Do you use exact or range versions for your dependencies?](https://www.reddit.com/r/javascript/comments/ira5gz/askjs_do_you_use_exact_or_range_versions_for_your/)
- url: https://www.reddit.com/r/javascript/comments/ira5gz/askjs_do_you_use_exact_or_range_versions_for_your/
---
I just wanted to to have a discussion about how you guys mange your dependencies in package.json

I know that a lot of people just use the range operator \^ which is what npm uses by default.  
I  personally don't like it and use exact instead. I have been burned in the past by packages not following semver and breaking their libraries  in minor updates.

What do you use in your projects? What are your thoughts ?
## [11][JavaScript alert on close tab](https://www.reddit.com/r/javascript/comments/irwd4e/javascript_alert_on_close_tab/)
- url: https://www.foxinfotech.in/2020/09/javascript-on-close-tab-show-warning-message.html
---

