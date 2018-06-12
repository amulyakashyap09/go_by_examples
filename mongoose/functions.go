package functions

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Callback struct {
	data interface{}
	err  error
}

// Insert : Function inserts the data object into the collection
// Input Parameters
// 		collection (Mgo Object) : Mongo Collection Object
// 		data(Object) : the object which has to be inserted
// Output Parameters
// 		boolean : whether operation was successfull or not
// 		error : if it was error then return error else nil

func Insert(collection *mgo.Collection, data interface{}) (bool, error) {
	err := collection.Insert(&data)
	if err != nil {
		log.Println("Error in the Insert function: ", err)
		return false, err
	}
	return true, err
}

// InsertAsync : Function inserts the data object into the collection
// Input Parameters
// 		collection (Mgo Object) : Mongo Collection Object
// 		data(Object) : the object which has to be inserted
// 		callback(channel) : passes the results to the callback after
// Output Parameters
// 		boolean : whether operation was successfull or not
// 		error : if it was error then return error else nil

func InsertAsync(collection *mgo.Collection, data interface{}, callback chan *Callback) {
	var result bool = true
	err := collection.Insert(&data)
	if err != nil {
		log.Println("Error in the Insert function: ", err)
		result = false
	}

	callback <- &Callback{result, err}
}

// Update : Function Updates the record into the collection
// Input Parameters
// 		collection (Mgo Object) : Mongo Collection Object
// 		query(Object) : the bson object which holds the update criteria
// 		data(Object) : the object which has to be inserted
// Output Parameters
// 		boolean : whether operation was successfull or not
// 		error : if it was error then return error else nil

func Update(collection *mgo.Collection, query bson.M, data interface{}) (bool, error) {
	err := collection.Update(query, data)
	if err != nil {
		log.Println("Error in the Insert function: ", err)
		return false, err
	}
	return true, nil
}

// UpdateAll : Function Updates all the record into the collection
// Input Parameters
// 		collection (Mgo Object) : Mongo Collection Object
// 		query(Object) : the bson object which holds the update criteria
// 		data(Object) : the object which has to be inserted
// Output Parameters
// 		info : details about the operation affected
// 		error : if it was error then return error else nil

func UpdateAll(collection *mgo.Collection, query bson.M, data interface{}) (interface{}, error) {
	info, err := collection.UpdateAll(query, data)
	if err != nil {
		log.Println("Error in the Insert function: ", info)
		return nil, err
	}
	return info, nil
}

//	FindByID : Function FindByID finds and returns record by Hexadecimal ID
//	Input Parameters :
//		collection *mgo.Collection : Mongo Collection Object
//		id : bson.ObjectId
//	Output :
//		interface{} : Returns Mongo Objects
func FindByID(collection *mgo.Collection, id bson.ObjectId) (interface{}, error) {
	var record interface{}
	err := collection.FindId(id).One(&record)
	if err != nil {
		log.Fatal(err)
	}
	return record, err
}

// Find : Function finds the record into the collection according to the query/criteria
// Input Parameters
// 		collection (Mgo Object) : Mongo Collection Object
// 		query(Bson Object) : the criteria which has to be evaluated
//		options (map[string]int) : other parameters like limit. skip, sort, etc
// Output Parameters
// 		error : if it was error then return error else nil
// 		boolean : whether operation was successfull or not
func Find(collection *mgo.Collection, query bson.M, options map[string]int) ([]interface{}, error) {

	var records []interface{}
	var err error

	limit, isLimit := options["limit"]
	skip, isSkip := options["isSkip"]

	if isLimit && isSkip {
		err = collection.Find(query).Skip(skip).Limit(limit).All(&records)
	} else if isLimit {
		err = collection.Find(query).Limit(limit).All(&records)
	} else if isSkip {
		err = collection.Find(query).Skip(skip).All(&records)
	} else {
		err = collection.Find(query).All(&records)
	}

	if err != nil {
		log.Fatal(err)
	}

	return records, err
}

// FindAll : Function finds all the records into the collection
// Input Parameters
// 		collection (Mgo Object) : Mongo Collection Object
// Output Parameters
// 		records : all the records present in database
// 		error : if it was error then return error else nil
func FindAll(collection *mgo.Collection) ([]interface{}, error) {
	var records []interface{}
	err := collection.Find(nil).All(&records)
	if err != nil {
		log.Fatal(err)
	}
	return records, err
}

// Remove : Function removes the record from the collection as per criteria/query
// Input Parameters
// 		collection (Mgo Object) : Mongo Collection Object
// 		query(Bson Object) : the criteria which has to be evaluated
// Output Parameters
// 		boolean : returns true / false depending on output of operation
// 		error : if it was error then return error else nil
func Remove(collection *mgo.Collection, query bson.M) (bool, error) {
	err := collection.Remove(query)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	return true, err
}

// RemoveAll : Function removes all the record from the collection
// Input Parameters
// 		collection (Mgo Object) : Mongo Collection Object
// Output Parameters
// 		info : returns info object as output of operation
// 		error : if it was error then return error else nil
func RemoveAll(collection *mgo.Collection) (interface{}, error) {
	info, err := collection.RemoveAll(nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return info, nil
}
