package main

import (
	"fmt"
	"os"
	"github.com/sky-uk/gonsx"
	"github.com/sky-uk/gonsx/api/securitytags"
	"errors"
)

// getAllSecurityTags - gets all securitytags
func getAllSecurityTags(nsxclient *gonsx.NSXClient) (*securitytags.SecurityTags, error){
	api := securitytags.NewGetAll()
	err := nsxclient.Do(api)

	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	} else{
		if api.StatusCode() == 200 {
			return api.GetResponse(), nil
		} else {
			return nil, errors.New(string(api.RawResponse()))
		}

	}
}

// createSecurityTag - creates securitytags
func createSecurityTag(name, desc string, nsxclient *gonsx.NSXClient) (string, error){
	api := securitytags.NewCreate(name, desc)
	err := nsxclient.Do(api)

	if err != nil {
		fmt.Println("Error: ", err)
		return "", err
	} else{
		fmt.Println("Creating security tag with id ", api.GetResponse())
		return api.GetResponse(), nil
	}
}

// deleteSecurityTag - deletes securitytags
func deleteSecurityTag(ID string, nsxclient *gonsx.NSXClient)(error){
	api := securitytags.NewDelete(ID)
	err := nsxclient.Do(api)

	if err != nil {
		fmt.Println("Error: ", err)
		return err
	} else{
		fmt.Println("Deleting security tag with id" , ID)
		return nil
	}
}

// RunSecurityTagExample - runs securitytag example
func RunSecurityTagExample(nsxManager, nsxUser, nsxPassword string, debug bool){
	nsxclient := gonsx.NewNSXClient(nsxManager, nsxUser, nsxPassword, true, debug)

	getTags, err := getAllSecurityTags(nsxclient)
	if err != nil {
		fmt.Println("Failed to get tags. error response:", err)
		os.Exit(1)
	}

	if !getTags.CheckByName("test") {
		_, err := createSecurityTag("test", "t", nsxclient)
		if err != nil {
			fmt.Println("Error", err)
		}
	} else{
		fmt.Println("Tag already exists")
	}

	getTags, err = getAllSecurityTags(nsxclient)
	if err != nil {
		fmt.Println("Failed to get tags")
		os.Exit(1)
	}

	if getTags.CheckByName("test") {
		ID := getTags.FilterByName("test").ObjectID
		err := deleteSecurityTag(ID, nsxclient)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}


}