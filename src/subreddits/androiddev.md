# androiddev
## [1][App Feedback Thread - April 25, 2020](https://www.reddit.com/r/androiddev/comments/g7sq25/app_feedback_thread_april_25_2020/)
- url: https://www.reddit.com/r/androiddev/comments/g7sq25/app_feedback_thread_april_25_2020/
---
This thread is for getting feedback on your own apps.

####Developers:

- must **provide feedback** for others
- must include **Play Store**, **GitHub**, or **BitBucket** link
- must make top level comment
- must make effort to respond to questions and feedback from commenters
- may be open or closed source

####Commenters:

- must give **constructive feedback** in replies to top level comments
- must not include links to other apps

To cut down on spam, accounts who are too young or do not have enough karma to post will be removed. Please make an effort to contribute to the community before asking for feedback.

As always, the mod team is only a small group of people, and we rely on the readers to help us maintain this subreddit. Please report any rule breakers. Thank you.

\- Da Mods
## [2][Weekly Questions Thread - April 20, 2020](https://www.reddit.com/r/androiddev/comments/g4qoms/weekly_questions_thread_april_20_2020/)
- url: https://www.reddit.com/r/androiddev/comments/g4qoms/weekly_questions_thread_april_20_2020/
---
This thread is for simple questions that don't warrant their own thread (although we suggest checking the sidebar, [the wiki](http://www.reddit.com/r/androiddev/wiki/), [our Discord](https://discord.gg/D2cNrqX), or [Stack Overflow](http://stackoverflow.com) before posting). Examples of questions:

* How do I pass data between my Activities?
* Does anyone have a link to the source for the AOSP messaging app?
* Is it possible to programmatically change the color of the status bar without targeting API 21?

**Important: Downvotes are strongly discouraged in this thread. Sorting by new is strongly encouraged.**

Large code snippets don't read well on reddit and take up a lot of space, so please don't paste them in your comments. Consider linking [Gists](https://gist.github.com) instead.

Have a question about the subreddit or otherwise for /r/androiddev mods? [We welcome your mod mail!](http://www.reddit.com/message/compose?to=%2Fr%2Fandroiddev)

Also, please don't link to Play Store pages or ask for feedback on this thread. Save those for the App Feedback threads we host on Saturdays.

Looking for all the Questions threads? Want an easy way to locate this week's thread? Click [this link](https://www.reddit.com/r/androiddev/search?q=title%3A%22questions+thread%22+author%3A%22AutoModerator%22&amp;restrict_sr=on&amp;sort=new&amp;t=all)!
## [3][Github Template for starting an Android app project with: 100% Kotlin + Github Actions + ktlint + Detekt + Gradle Kotlin DSL + buildSrc dependencies already set up.](https://www.reddit.com/r/androiddev/comments/g7rhi8/github_template_for_starting_an_android_app/)
- url: https://github.com/cortinico/kotlin-android-template
---

## [4][AsyncAndroid - "Be Together even When We're Apart". A new YouTube channel gathering collections of developer-created content from just some of the amazing members of the Android community and publishes them in drops so that we all can share, teach, learn, and connect remotely.](https://www.reddit.com/r/androiddev/comments/g7ncez/asyncandroid_be_together_even_when_were_apart_a/)
- url: http://youtube.com/asyncandroid
---

## [5][How to decouple threading and execution (e.g: RxJava, Coroutine) from repositories and use cases?](https://www.reddit.com/r/androiddev/comments/g7qgab/how_to_decouple_threading_and_execution_eg_rxjava/)
- url: https://www.reddit.com/r/androiddev/comments/g7qgab/how_to_decouple_threading_and_execution_eg_rxjava/
---
Decoupling the dependencies into modules will help you to easily maintain the code, test and **replace** dependencies. That's cool, but all the tutorials that I'm watching they are implementing directly RxJava and Coroutine directly into the modules.

High level modules will look something like this:

RxJava:

    fun getUsers(): Observable&lt;List&lt;User&gt;&gt;

Coroutine:

    suspend fun getUsers() = suspendCancellableCoroutine { continuation -&gt;
       if (!continuation.isCancelled) {
        ....

And in order to interrupt or cancel the execution, we need to store the RxJava's disposable into a CompositeDisposable and for the Coroutine a list of jobs that we can cancel.

Let's say we're having a RegisterUserUseCase within our ViewModel, and when the onCleared is called we call the use case's class dispose function: 

    fun dispose() {
       // Rx Java
       this.compositeDisposable().dispose()
    
       // Coroutine
       this.jobs.forEach { job -&gt; job.cancel() }
    }

What if in the future I want to replace RxJava or Coroutine with something else? Let's say I want to use Realm/Firebase or some other 3rd party library. Realm and firebase are using their own observable pattern/callbacks and if I remove RxJava or Coroutine I need to refactor the whole module(s), because I directly used RxJava and Coroutine (all functions are declared suspend and I need to remove the suspend name; etc).

Declaring RxJava or Coroutines directly doesn't contradict with the ideaology of clean architecture/S.O.L.I.D? All I can think now is having another layer of abstraction where the interface class are having the basic CRUD operations and lifecycle calls (create/destroy) and within this layer to implement the data storing logic.
## [6][Dagger SPI - Extending Dagger with custom Dependency Graph validations](https://www.reddit.com/r/androiddev/comments/g7rec1/dagger_spi_extending_dagger_with_custom/)
- url: https://arunkumar.dev/dagger-spi-building-custom-validations-for-dependency-graphs/
---

## [7][[Flutter] Google Sign in with Firebase Authentication Tutorial](https://www.reddit.com/r/androiddev/comments/g7oy8c/flutter_google_sign_in_with_firebase/)
- url: https://youtu.be/FmD1c5DIMYI
---

## [8][Android 11 DP3 new function summary: automatically cancel application permissions / share recent apps / Ethernet hotspot / ...](https://www.reddit.com/r/androiddev/comments/g75le9/android_11_dp3_new_function_summary_automatically/)
- url: https://i.redd.it/zfpsc763mqu41.png
---

## [9][Anyone know a way to snoop on creative ad assets (ads under which they advertise) from competitors?](https://www.reddit.com/r/androiddev/comments/g7th1o/anyone_know_a_way_to_snoop_on_creative_ad_assets/)
- url: https://www.reddit.com/r/androiddev/comments/g7th1o/anyone_know_a_way_to_snoop_on_creative_ad_assets/
---

## [10][Anyone who Android development with WSL 2 and a real Android device? How do you get the USB to be recognized by Ubuntu?](https://www.reddit.com/r/androiddev/comments/g7szca/anyone_who_android_development_with_wsl_2_and_a/)
- url: https://www.reddit.com/r/androiddev/comments/g7szca/anyone_who_android_development_with_wsl_2_and_a/
---

## [11][In App Purchases implementation](https://www.reddit.com/r/androiddev/comments/g7swm9/in_app_purchases_implementation/)
- url: https://www.reddit.com/r/androiddev/comments/g7swm9/in_app_purchases_implementation/
---
New to android development and I'm trying to implement a simple non - consumable one time in app purchase for my app, specifically to remove the ads. I have added the managed product in google play console and the 'com.android.billingclient:billing:2.2.0' on my gradle file. After lots of research and random YouTube videos i ended up with this piece of code:

    public class MainActivity extends AppCompatActivity implements PurchasesUpdatedListener {
    
        private BillingClient billingClient;
        private List skuList = new ArrayList();
        private String sku = "remove_ads";
        private Button buyBtn;
       
       Boolean b = getBoolFromPref(this, "myPref", sku);
            if (b) {
    		
                buyBtn = findViewById(R.id.buyBtn);
                buyBtn.setVisibility(View.INVISIBLE);
                buyBtn.setEnabled(false)
    		
    		// remove ads here
    		
    		} else {
    
    
                buyBtn = findViewById(R.id.buyBtn);
                buyBtn.setVisibility(View.VISIBLE);
                buyBtn.setEnabled(true)
    
                        // enable the ads here
    
    }
       
       
        private void setupBillingClient() {
    
            billingClient = BillingClient.newBuilder(this).enablePendingPurchases().setListener(this).build();
            billingClient.startConnection(new BillingClientStateListener() {
                @Override
                public void onBillingSetupFinished(BillingResult billingResult) {
    
                    if (billingResult.getResponseCode() == BillingClient.BillingResponseCode.OK) {
    
                        loadAllSkus();
                        
                    }
    
                }
    
                @Override
                public void onBillingServiceDisconnected() {
    			
    
                }
            });
        }
    
        private void loadAllSkus() {
    
            if (billingClient.isReady()) {
                final SkuDetailsParams params = SkuDetailsParams.newBuilder().setSkusList(skuList).setType(BillingClient.SkuType.INAPP).build();
    
                billingClient.querySkuDetailsAsync(params, new SkuDetailsResponseListener() {
                    @Override
                    public void onSkuDetailsResponse(BillingResult billingResult, List&lt;SkuDetails&gt; skuDetailsList) {
                        if (billingResult.getResponseCode() == BillingClient.BillingResponseCode.OK) {
    
                            for (Object skuDetailsObject : skuDetailsList) {
    
                                final SkuDetails skuDetails = (SkuDetails) skuDetailsObject;
                                if (skuDetails.getSku().equals(sku)) {
    
                                    buyBtn = findViewById(R.id.buyBtn);
                                    buyBtn.setOnClickListener(new View.OnClickListener() {
                                        @Override
                                        public void onClick(View v) {
                                            BillingFlowParams params = BillingFlowParams.newBuilder().setSkuDetails(skuDetails).build();
                                            billingClient.launchBillingFlow(MainActivity.this, params);
                                        }
                                    });
    
                                }
                            }
                        }
                    }
                });
            }
        }
    
        @Override
        public void onPurchasesUpdated(BillingResult billingResult, @Nullable List&lt;Purchase&gt; purchases) {
    
            int responseCode = billingResult.getResponseCode();
    
            if (responseCode == BillingClient.BillingResponseCode.OK &amp;&amp; purchases != null) {
    
                for (Purchase purchase : purchases) {
    
                    handlePurchase(purchase);
    
                }
    
    
            } else if (responseCode == BillingClient.BillingResponseCode.ITEM_ALREADY_OWNED) {
    
                setBoolInPref(this, "myPref", sku, true);
    			
            } else if (responseCode == BillingClient.BillingResponseCode.USER_CANCELED) {
    		
    
            }
    
        }
    
        private Boolean getBoolFromPref(Context context, String prefName, String constantName) {
            SharedPreferences pref = context.getSharedPreferences(prefName, 0); // 0 - for private mode
    
            return pref.getBoolean(constantName, false);
    
        }
    
        private void handlePurchase(Purchase purchase) {
    
            if (purchase.getSku().equals(sku) &amp;&amp; purchase.getPurchaseState() == Purchase.PurchaseState.PURCHASED) {
    
                setBoolInPref(this, "myPref", sku, true);
    
                AcknowledgePurchaseParams acknowledgePurchaseParams =
                        AcknowledgePurchaseParams.newBuilder()
                                .setPurchaseToken(purchase.getPurchaseToken())
                                .build();
    
                AcknowledgePurchaseResponseListener acknowledgePurchaseResponseListener = new AcknowledgePurchaseResponseListener() {
                    @Override
                    public void onAcknowledgePurchaseResponse(BillingResult billingResult) {
    
                       
    					
                    }
    
                };
    
                billingClient.acknowledgePurchase(acknowledgePurchaseParams, acknowledgePurchaseResponseListener);
    
                Toast.makeText(this, "Thanks for buying. Ads are now removed. Enjoy.", Toast.LENGTH_SHORT).show();
    
    
            } else if (purchase.getPurchaseState() == Purchase.PurchaseState.PENDING) {
    		
                // Here you can confirm to the user that they've started the pending
                // purchase, and to complete it, they should follow instructions that
                // are given to them. You can also choose to remind the user in the
                // future to complete the purchase if you detect that it is still
                // pending.
    			
            }
    
        }
    
    
        private void setBoolInPref(Context context, String prefName, String constantName, Boolean val) {
            SharedPreferences pref = context.getSharedPreferences(prefName, 0); // 0 - for private mode
    
            SharedPreferences.Editor editor = pref.edit();
            editor.putBoolean(constantName, val);
            editor.commit();
    		
        }

Using the test purchase with myself it seems to work fine. It launches the billing flow fine and after the purchase the ads are removed and also the buyBtn disappears.

However there is a small weird problem.

If i uninstall and then install the app again, the ads re appear, as well the buyBtn. 

But, (and this is where I am really confused) if I click on the buyBtn and restart the app, the ads are removed together with the buyBtn. The billing flow doesn't launch nor there is any other notification.
## [12][Android 11 Toast Updates](https://www.reddit.com/r/androiddev/comments/g7anno/android_11_toast_updates/)
- url: https://medium.com/@myrickchow32/android-11-toast-updates-7f1cd2245bc4?source=friends_link&amp;sk=231403e77b7409125b8398c9246b7eb7
---

