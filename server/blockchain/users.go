package blockchain

import (
	"encoding/json"
	"fmt"
	"strconv"
	"io/ioutil"
	"net/http"
	//TO DO later : "golang.org/x/crypto/bcrypt"
)

// User : équipe
type User struct {
	// userId
	ID string `json:"userId"`
	// userName
	UserName string `json:"userName"`
	// userPassword
	UserPassword string `json:"userPassword"`
}

// ConnectUser
func ConnectUser(username, password string) (User, error) {
	var usrs []User
	usrs, err := GetUsers()
	if err != nil {
		return User{}, fmt.Errorf("Error getting users", err)
	}
	for _, user := range usrs {
		fmt.Println(user)
		fmt.Println(usrs)
		if(user.UserName == username){
			// TO DO later : hashedPassword, err := bcrypt.GenerateFromPassword(, bcrypt.DefaultCost)
		    //if err != nil {
		    //    return false,fmt.Errorf("Error generate hash password", err)
		    //}

		    // Comparing the password with the hash
		    // TO DO later : err = bcrypt.CompareHashAndPassword([]byte(user.userPassword), []byte(password))
		    
		    if user.UserPassword != password {
		    	return User{}, fmt.Errorf("Error wrong password")
		    }
		    return user, nil
		}
	}
	return User{}, fmt.Errorf("User not found")
}


// GetUsers
func GetUsers() ([]User, error) {
	client := &http.Client{}
	uri := "world.alphabets.User"

	url := apiURI + uri

	// Prepare the request
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return []User{}, fmt.Errorf("Error creating the request: %v", err)
	}

	// Execute the request
	res, err := client.Do(req)

	if err != nil {
		return []User{}, fmt.Errorf("Error executing the request: %v", err)
	}

	defer res.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []User{}, fmt.Errorf("Error reading response: %v", err)
	}

	var users []User

	if err = json.Unmarshal(body, &users); err != nil {
		return []User{}, fmt.Errorf("Error unmarshalling the response data: %v", err)
	}

	fmt.Println(users, string(body))
	return users, nil
}

// GetUser : retrieves a user from the sportmonks API
func GetUser(userId int) (User, error) {
	uri := "world.alphabets.User/" + strconv.Itoa(userId)

	response, err := getBcAnything(uri)

	if err != nil {
		return User{}, fmt.Errorf("Error getting the data: %v", err)
	}

	var user User

	// Marshal the data part in order to decode it from JSON later
	jsonEncodedUser, err := json.Marshal(response.Data)

	if err != nil {
		return User{}, fmt.Errorf("Error marshalling the user: %v", err)
	}

	if err = json.Unmarshal(jsonEncodedUser, &user); err != nil {
		return User{}, fmt.Errorf("Error unmarshalling the response data: %v", err)
	}

	return user, nil
}
