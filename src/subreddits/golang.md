# golang
## [1][What feature of Go is used very often by experienced programmers, but not so much by a newbie?](https://www.reddit.com/r/golang/comments/hp4mk3/what_feature_of_go_is_used_very_often_by/)
- url: https://www.reddit.com/r/golang/comments/hp4mk3/what_feature_of_go_is_used_very_often_by/
---

## [2][Resources to learn Web Programming with Go](https://www.reddit.com/r/golang/comments/hozqjs/resources_to_learn_web_programming_with_go/)
- url: https://www.reddit.com/r/golang/comments/hozqjs/resources_to_learn_web_programming_with_go/
---
Suggestions please!
## [3][Under the hood of Go's context](https://www.reddit.com/r/golang/comments/hp8s68/under_the_hood_of_gos_context/)
- url: https://vishnubharathi.codes/blog/go-contexts/
---

## [4][Bool vs enum](https://www.reddit.com/r/golang/comments/hp6a7h/bool_vs_enum/)
- url: https://www.reddit.com/r/golang/comments/hp6a7h/bool_vs_enum/
---
Hey there,

I would like to get your opinion on something which happend to me last week:

A coworker wrote a function which initially gave back a bool based on some checking. Then he realised that he needs three states (start, stop and nothing).
He then changed the func to return a pointer to a bool which was not initialized when to do nothing.

I totally disagree with this and told him to use enums. But he said he likes his idea better.

So, what do you think is the better approach?

Thanks in advance
## [5][Web framework that can automatically generate OpenAPI/Swagger](https://www.reddit.com/r/golang/comments/hp94nf/web_framework_that_can_automatically_generate/)
- url: https://www.reddit.com/r/golang/comments/hp94nf/web_framework_that_can_automatically_generate/
---
I'm coming to the Go world from Java. And I'm especially new to the "web" frameworks. I've been using Echo on 2 services, and found that it's pretty nice.   


The biggest thing I'm missing is the automatic generation of Swagger/OpenAPI spec from source.  
I found tools that can read the source code and look for special annotation, but it's still a lot of manual work to add around, where in Java it's done without those.

I'm not sure if it's possible in Go..
## [6][DNF5 needs your help with Go bindings](https://www.reddit.com/r/golang/comments/hoq5p3/dnf5_needs_your_help_with_go_bindings/)
- url: https://www.reddit.com/r/golang/comments/hoq5p3/dnf5_needs_your_help_with_go_bindings/
---
This [PR](https://github.com/rpm-software-management/libdnf/pull/988) adds Swig bindings to libdnf (the lib behind the Redhat/Fedora/CentOS/Mageia package manager), but currently there is nobody checking if this is actually working for Go. So if you've got a use case for that, please go check it out and provide some feedback!
## [7][Google's go safeweb. Collection of libraries to write safe web servers in go by default.](https://www.reddit.com/r/golang/comments/hp6aqy/googles_go_safeweb_collection_of_libraries_to/)
- url: https://github.com/google/go-safeweb
---

## [8][An Analyzer makes the go/analysis.Analyzer be able to ignore diagnostics with //nolint comment](https://www.reddit.com/r/golang/comments/hp41af/an_analyzer_makes_the_goanalysisanalyzer_be_able/)
- url: https://github.com/kyoh86/nolint/blob/master/README.md
---

## [9][Adaptive routing for load balancing](https://www.reddit.com/r/golang/comments/homsvy/adaptive_routing_for_load_balancing/)
- url: https://www.reddit.com/r/golang/comments/homsvy/adaptive_routing_for_load_balancing/
---
Couple years back, I worked on serviceq ([https://github.com/gptankit/serviceq](https://github.com/gptankit/serviceq)) which is a load balancer written in Go, that aims to dynamically adapt to errors in downstream nodes. It does this by   
a) adaptive routing  
b) queueing and retrying failed requests

Have exposed the adaptive routing logic as a package - harmonic ([https://github.com/gptankit/harmonic](https://github.com/gptankit/harmonic)) - which can be used from client apps, load balancers or test suites, if needed.

Comments/suggestions are welcome.
## [10][How to reduce unit test function size ?](https://www.reddit.com/r/golang/comments/hp0qc0/how_to_reduce_unit_test_function_size/)
- url: https://www.reddit.com/r/golang/comments/hp0qc0/how_to_reduce_unit_test_function_size/
---
Hey,

&amp;#x200B;

I'm building a web app and trying to learn unit test at the same time. I end up having alot of boiler plate in each test because of having to mock and inject dependencies alot. How could I write better test ?

Small example below :

	func TestHandlerGetUser(t *testing.T) {
		service := new(mocks.Service)
		handler := user.NewHandler(service)
		e := echo.New()

		var u domain.User
		faker.FakeData(&amp;u)

		t.Run("should return user with given id", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			service.On("GetByID", mock.Anything, u.ID).Return(u, nil).Once()
			c.SetPath("/api/users/:id")
			c.SetParamNames("id")
			c.SetParamValues(strconv.Itoa(u.ID))

			err := handler.GetUser(c)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)

			body, _ := ioutil.ReadAll(rec.Result().Body)
			var responseUser domain.User
			json.Unmarshal(body, &amp;responseUser)
			assert.EqualValues(t, u, responseUser)
		})

		t.Run("should return bad request with wrong params", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/api/users/:id")
			c.SetParamNames("id")
			c.SetParamValues(u.Email)

			err := handler.GetUser(c)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		})

		t.Run("should return not found", func(t *testing.T) {
			service.On("GetByID", mock.Anything, u.ID).Return(domain.User{}, sql.ErrNoRows).Once()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/api/users/:id")
			c.SetParamNames("id")
			c.SetParamValues(strconv.Itoa(u.ID))

			err := handler.GetUser(c)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusNotFound, rec.Code)
		})

		t.Run("should return internal server error", func(t *testing.T) {
			service.On("GetByID", mock.Anything, u.ID).Return(domain.User{}, errors.New("fail")).Once()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/api/users/:id")
			c.SetParamNames("id")
			c.SetParamValues(strconv.Itoa(u.ID))

			err := handler.GetUser(c)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		})

	}

My handler function is the following :

	func (h *handler) GetUser(c echo.Context) error {
		ctx := c.Request().Context()
		userID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "invalid request")
		}
		u, err := h.service.GetByID(ctx, userID)
		if err == sql.ErrNoRows {
			return c.String(http.StatusNotFound, "user not found")
		}
		if err != nil {
			return c.String(http.StatusInternalServerError, "unable to get user")
		}
		return c.JSON(http.StatusOK, u)
	}

Little update, tried to follow faabiosr comment :

	func TestHandlerGetUser(t *testing.T) {
		e := echo.New()
		var u domain.User
		faker.FakeData(&amp;u)

		type test struct {
			name         string
			param        string
			status       int
			serviceError error
			body         interface{}
		}

		table := []test{
			{"should return user", strconv.Itoa(u.ID), 200, nil, u},
			{"should return wrong parameters", u.Email, http.StatusBadRequest, nil, "invalid request"},
			{"should return no user found", strconv.Itoa(u.ID), http.StatusNotFound, sql.ErrNoRows, "user not found"},
			{"should return internal server error", strconv.Itoa(u.ID), http.StatusInternalServerError, errors.New("fail"), "unable to get user"},
		}

		for _, test := range table {
			t.Run(test.name, func(t *testing.T) {
				service := new(mocks.Service)
				handler := user.NewHandler(service)
				service.On("GetByID", mock.Anything, u.ID).Return(u, test.serviceError).Once()

				req := httptest.NewRequest(http.MethodGet, "/", nil)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetPath("/api/users/:id")
				c.SetParamNames("id")
				c.SetParamValues(test.param)

				err := handler.GetUser(c)
				assert.NoError(t, err)
				assert.Equal(t, test.status, rec.Code)

				recBody, _ := ioutil.ReadAll(rec.Result().Body)
				var recBodyContent interface{}
				var testBodyContent interface{}
				testBody, _ := json.Marshal(test.body)
				json.Unmarshal(recBody, &amp;recBodyContent)
				json.Unmarshal(testBody, &amp;testBodyContent)
				assert.EqualValues(t, testBodyContent, recBodyContent)
			})
		}
	}
