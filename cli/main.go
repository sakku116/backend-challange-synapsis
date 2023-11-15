package cli

import (
	"fmt"
	"synapsis/domain/model"
	"synapsis/exception"
	"synapsis/repository"

	"golang.org/x/crypto/bcrypt"
)

// args should be empty for default seed (superuser;superuser)
// or args must be containing 2 strings for custom username, and passwords
func SeedSuperuser(userRepo repository.IUserRepo, args ...string) {
	// validate args
	if len(args) != 3 && len(args) != 0 {
		fmt.Println("invalid args, should be empty for default seed (superuser;superuser) or 2 strings for custom username, and passwords")
		return
	}

	username := "superuser"
	password := "superuser"
	if len(args) == 2 {
		username = args[0]
		password = args[1]
	}

	// check for existing user
	existingUser, err := userRepo.GetByUsername(username)
	if err != nil && err != exception.DbObjNotFound {
		panic(err)
	}
	if existingUser != nil {
		fmt.Printf("%s already exists\n", username)
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user := &model.User{
		Username: username,
		Password: string(hashedPass),
	}
	userRepo.Create(user)
}

func SeedProduct(productRepo repository.IProductService) {
	// products
	product1 := &model.Product{
		Name:  "susu",
		Price: 100,
	}
	product3 := &model.Product{
		Name:  "kopi",
		Price: 300,
	}
	product4 := &model.Product{
		Name:  "jeruk",
		Price: 400,
	}
	products := []model.Product{*product1, *product3, *product4}

	// var productsResult []model.Product
	for _, product := range products {
		_, err := productRepo.GetByNameAndPrice(product.Name, product.Price)
		if err == exception.DbObjNotFound {
			err = productRepo.Create(&product)
			if err != nil {
				panic(err)
			}
			// append(productsResult, product)
		} else if err != nil {
			fmt.Printf("error: %v\n", err)
			// append(productsResult, existing)
			continue
		}
	}
}
