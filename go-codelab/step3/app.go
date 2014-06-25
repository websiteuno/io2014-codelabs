//  Package todo is an App Engine app providing a REST API to manage todo
//  lists.
//
//  The REST API provides the following handlers:
//
//  /api/lists GET
//    Gets all the list names and ids and creators if no user is logged in.
//    If the user is logged in, only their lists.
//  /api/lists POST
//    Creates a new list.
//  /api/list/{list} GET
//    Gets the name, creator, and id of a list with id {list}.
//  /api/list/{list} DELETE
//    Deletes the list with id {list}.
//
package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"

	"appengine"
	"appengine/datastore"
)

// datastore entity kind for List
const listKind = "List"

// List represents a lists of tasks.
type List struct {
	// Autogenerated id, not stored in the datastore.
	ID string `datastore:"-"`

	Name string // Name of the list
}

func init() {
	// Register all the handlers.
	r := mux.NewRouter().PathPrefix("/api/").Subrouter()

	// List management
	r.Handle("/list", appHandler(getAllLists)).Methods("GET")
	r.Handle("/list", appHandler(createList)).Methods("POST")
	r.Handle("/list/{list}", appHandler(getList)).Methods("GET")
	r.Handle("/list/{list}", appHandler(deleteList)).Methods("DELETE")

	http.Handle("/api/", r)
}

// getAllLists fetches all the lists in the datastore and encodes them
// in JSON format into the http response.
func getAllLists(w io.Writer, r *http.Request) error {
	c := appengine.NewContext(r)

	lists := []List{}
	keys, err := datastore.NewQuery(listKind).GetAll(c, &lists)
	if err != nil {
		return fmt.Errorf("fetch all lists: %v", err)
	}

	// Update the encoded keys and encode the lists.
	for i, k := range keys {
		lists[i].ID = k.Encode()
	}
	return json.NewEncoder(w).Encode(lists)
}

// createList creates a new list. It reads a JSON encoded list from the request
// body.
func createList(w io.Writer, r *http.Request) error {
	c := appengine.NewContext(r)

	// Decode a list from the request body.
	list := List{}
	err := json.NewDecoder(r.Body).Decode(&list)
	if err != nil {
		return appErrorf(http.StatusBadRequest, "decode list: %v", err)
	}

	if list.Name == "" {
		return appErrorf(http.StatusBadRequest, "missing list name")
	}

	// Put the List in the datastore.
	key := datastore.NewIncompleteKey(c, listKind, nil)
	key, err = datastore.Put(c, key, &list)
	if err != nil {
		return fmt.Errorf("create list: %v", err)
	}

	// Update the encoded key and encode the list.
	list.ID = key.Encode()
	return json.NewEncoder(w).Encode(list)
}

// getList fetches the list with the id given in the url and encodes it in
// JSON format into the http response.
func getList(w io.Writer, r *http.Request) error {
	return errors.New("getList not implemented")

	// Get the list id from the URL, identified by the name "list" in the url.
	// - mux.Vars: http://godoc.org/github.com/gorilla/mux#Vars

	// Decode the obtained id into a datastore key.
	// - datastore.DecodeKey: http://golang.org/s/datastore#DecodeKey

	// Fetch the list from the datastore using the Get method.
	// If the retured error is datastore.ErrNoSuchEntity you should return an
	// appError with http.StatusNotFound code.
	// For other errors just return the obtained error and it will be handled
	// as a http.StatusServerInternalError (500).
	// - datastore.Get: http://golang.org/s/datastore#Get

	// Set the ID field with the id from the request url and encode the list.
}

// deleteList deletes the list with the id given in the url.
func deleteList(w io.Writer, r *http.Request) error {
	return errors.New("deleteList not implemented")

	// Get the list id from the URL, identified by the name "list" in the url.
	// - mux.Vars: http://godoc.org/github.com/gorilla/mux#Vars

	// Decode the obtained id into a datastore key.
	// - datastore.DecodeKey: http://golang.org/s/datastore#DecodeKey

	// Delete the list.
	// - datastore.Delete: http://golang.org/s/datastore#Delete
}
