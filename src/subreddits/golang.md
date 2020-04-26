# golang
## [1][What's coming in Go 1.15 — Slides By Daniel Martí (mvdan)](https://www.reddit.com/r/golang/comments/g8d8jk/whats_coming_in_go_115_slides_by_daniel_martí/)
- url: https://docs.google.com/presentation/d/1veyF0y6Ynr6AFzd9gXi4foaURlgbMxM-tmB4StDrdAM/edit#slide=id.g550f852d27_228_0
---

## [2][Chat with Go, React and k8s](https://www.reddit.com/r/golang/comments/g84yg2/chat_with_go_react_and_k8s/)
- url: https://www.reddit.com/r/golang/comments/g84yg2/chat_with_go_react_and_k8s/
---
Hey everybody! 
I built this simple chat in Golang and React and hosted it in a Kubernetes cluster in AWS. 
Though it might be useful to share a project with the full CI/CD pipeline.

[github repo](https://github.com/leartgjoni/go-chat-api)
## [3][[discussion] Why I'm hesitant / afraid of adopting go as main language](https://www.reddit.com/r/golang/comments/g8aujj/discussion_why_im_hesitant_afraid_of_adopting_go/)
- url: https://www.reddit.com/r/golang/comments/g8aujj/discussion_why_im_hesitant_afraid_of_adopting_go/
---
Is this normal behavior in Golang community? [https://i.imgur.com/wu9NAOb.png](https://i.imgur.com/wu9NAOb.png)

What's wrong with that framework? I'd love to use only stdlib but it feel too low level. Are we really expected for each project to manually parse the URL paths and map requests to their handlers? Or do we end up anyway making our own "framework" that we just keep private due to this kind of remarks?

What's wrong with having nice APIs like:

    app.Get("/", func (c *fiber.Ctx) {
        c.Send("Hello, World!")
    })

vs. needing to manually parse URL paths and their params?

Why don't we have a better router/mux in the stdlib if people hate so much on these frameworks?
## [4][How We Created a Realtime Patient Monitoring System With Go and Vue in 3 days - kasvith.me](https://www.reddit.com/r/golang/comments/g7px4d/how_we_created_a_realtime_patient_monitoring/)
- url: https://kasvith.me/posts/how-we-created-a-realtime-patient-monitoring-system-with-go-and-vue
---

## [5][Noise Gate - Golang test runner to get faster test results](https://www.reddit.com/r/golang/comments/g8b4br/noise_gate_golang_test_runner_to_get_faster_test/)
- url: https://github.com/go-noisegate/noisegate
---

## [6][GoRemoteFest April 26th 10:00 14:00 GMT](https://www.reddit.com/r/golang/comments/g7u7xc/goremotefest_april_26th_1000_1400_gmt/)
- url: https://www.reddit.com/r/golang/comments/g7u7xc/goremotefest_april_26th_1000_1400_gmt/
---
[https://goremotefest.com/](https://goremotefest.com/)
## [7][Proper testing pattern a la Python custom testing classes](https://www.reddit.com/r/golang/comments/g81hc0/proper_testing_pattern_a_la_python_custom_testing/)
- url: https://www.reddit.com/r/golang/comments/g81hc0/proper_testing_pattern_a_la_python_custom_testing/
---
I'll just duck in advance, but I'm having a hard time searching for this and getting a positive/negative answer.

Is there a pattern where it's acceptable to embed testing.T in a custom type in order to add common helpers for setup, assertions, etc?
## [8][Inlining optimisations in Go by Dave Cheney](https://www.reddit.com/r/golang/comments/g7pvp3/inlining_optimisations_in_go_by_dave_cheney/)
- url: https://dave.cheney.net/2020/04/25/inlining-optimisations-in-go
---

## [9][Why for-range behave differently depending on the size of the element](https://www.reddit.com/r/golang/comments/g7wgir/why_forrange_behave_differently_depending_on_the/)
- url: https://labs.yulrizka.com/en/why-for-range-behave-differently-depending-on-the-size-of-the-element/
---

## [10][COVID-19 Gophers made by my GF](https://www.reddit.com/r/golang/comments/g79py7/covid19_gophers_made_by_my_gf/)
- url: https://i.redd.it/84e8h6x74su41.jpg
---

