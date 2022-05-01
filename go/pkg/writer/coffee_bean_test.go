package writer_test

// func TestCoffeeBean_Create(t *testing.T) {
// 	t.Parallel()

// 	impl := writer.NewCoffeeBeanWriter(testClient)
// 	user := db_helper.InsertAndDeleteUsers(t, testClient, func(u *ent.User) {
// 		u.Email = "create@example.com"
// 	})

// 	cases := map[string]struct {
// 		coffeeBean *model.UserCoffeeBean
// 		user       *model.User
// 	}{
// 		"success": {
// 			coffeeBean: &model.UserCoffeeBean{
// 				Name:        "test name",
// 				FarmName:    "test farm name",
// 				Country:     "test country",
// 				RoastDegree: model.RoastDegreeFrench,
// 				RoastedAt:   time.Date(2022, time.February, 2, 0, 0, 0, 0, time.UTC),
// 				CreatedAt:   time.Date(2022, time.February, 20, 0, 0, 0, 0, time.UTC),
// 				UpdatedAt:   time.Date(2022, time.February, 20, 0, 0, 0, 0, time.UTC),
// 			},
// 			user: &model.User{
// 				ID: int(user.ID),
// 			},
// 		},
// 	}

// 	for name, c := range cases {
// 		c := c
// 		t.Run(name, func(t *testing.T) {
// 			t.Parallel()

// 			ctx := context.Background()
// 			if err := impl.Create(ctx, c.coffeeBean, c.user); err != nil {
// 				t.Errorf("err should be nil, but got %q\n", err)
// 			}

// 			exists, err := testClient.UserCoffeeBean.Query().
// 				Where(coffeebean.IDEQ(int32(c.coffeeBean.ID))).Exist(ctx)
// 			if err != nil {
// 				t.Errorf("err should be nil, but got %v", err)
// 			}
// 			if diff := cmp.Diff(true, exists); diff != "" {
// 				t.Errorf("coffee bean should be created")
// 			}

// 			exists, err = testClient.UserCoffeeBean.Query().
// 				Where(userscoffeebean.CoffeeBeanIDEQ(int32(c.coffeeBean.ID))).Exist(ctx)
// 			if err != nil {
// 				t.Errorf("err should be nil, but got %v", err)
// 			}
// 			if diff := cmp.Diff(true, exists); diff != "" {
// 				t.Errorf("users coffee beans should be created")
// 			}
// 		})
// 	}
// }

// func TestCoffeeBean_DeleteByID(t *testing.T) {
// 	t.Parallel()

// 	impl := writer.NewCoffeeBeanWriter(testClient)
// 	user := db_helper.InsertAndDeleteUsers(t, testClient, func(u *ent.User) {
// 		u.Email = "delete-by-id"
// 	})
// 	userCoffeeBean := db_helper.InsertAndDeleteUserCoffeeBean(t, testClient, user)

// 	cases := map[string]struct {
// 		userCoffeeBean model.UserCoffeeBean
// 	}{
// 		"success": {
// 			userCoffeeBean: model.NewCoffeeBean(userCoffeeBean),
// 		},
// 	}

// 	for name, c := range cases {
// 		c := c
// 		t.Run(name, func(t *testing.T) {
// 			t.Parallel()

// 			ctx := context.Background()
// 			if err := impl.DeleteByID(ctx, &c.userCoffeeBean); err != nil {
// 				t.Errorf("err should be nil but, got %q", err)
// 			}

// 			ucb, err := testClient.UserCoffeeBean.Query().
// 				Where(userscoffeebean.IDEQ(int32(userCoffeeBean.ID))).
// 				Only(ctx)
// 			if err != nil {
// 				t.Errorf("err should be nil, but got %v", err)
// 			}
// 			if diff := cmp.Diff(model.CoffeeBeanStatusDeleted.Num(), int(ucb.Status)); diff != "" {
// 				t.Errorf("status should be deleted, but got %v", ucb.Status)
// 			}
// 		})
// 	}
// }
