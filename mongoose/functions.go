package main

import (
	"log"
)

// bulkInsert : Function inserts the data in bulk to the collection
// Input Parameters
// 		collection (Mgo Object) : Mongo Collection Object
// 		data(Array of Objects) : the object which has to be inserted
// Output Parameters
// 		boolean : whether operation was successfull or not
// 		error : if it was error then return error else nil
func bulkInsert(bulkInsertStruct *BulkInsert) (bool, error) {
	var records bool = true
	bulk := bulkInsertStruct.Collection.Bulk()
	bulk.Insert(bulkInsertStruct.Data)
	_, err := bulk.Run()
	if err != nil {
		log.Fatal(err)
		records = false
	}
	return records, err
}

// insert : Function inserts the data object into the collection
// Input Parameters
// 		collection (Mgo Object) : Mongo Collection Object
// 		data(Object) : the object which has to be inserted
// Output Parameters
// 		boolean : whether operation was successfull or not
// 		error : if it was error then return error else nil
func insert(insertStruct *Insert) (bool, error) {
	err := insertStruct.Collection.Insert(&insertStruct.Data)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	return true, err
}

// insertAsync : Function inserts the data object into the collection
// Input Parameters
// 		collection (Mgo Object) : Mongo Collection Object
// 		data(Object) : the object which has to be inserted
//		callback (channel) : which returns data to goroutine
// Output Parameters
// 		callback : sends output back to channel

func insertAsync(insertStruct *Insert, callback chan *Callback) {
	records, err := insert(insertStruct)
	cb := new(Callback)
	cb.Data = records
	cb.Error = err
	callback <- cb
}

// update : Function Updates the record into the collection
// Input Parameters
// 		collection (Mgo Object) : Mongo Collection Object
// 		query(Object) : the bson object which holds the update criteria
// 		data(Object) : the object which has to be inserted
// Output Parameters
// 		boolean : whether operation was successfull or not
// 		error : if it was error then return error else nil

func update(updateStruct *Update) (bool, error) {
	err := updateStruct.Collection.UpdateId(updateStruct.Id, updateStruct.Data)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	return true, nil
}

// updateAsync : Function Updates the record into the collection
// Input Parameters
// 		collection (Mgo Object) : Mongo Collection Object
// 		query(Object) : the bson object which holds the update criteria
// 		data(Object) : the object which has to be inserted
//		callback (channel) : which returns data to goroutine
// Output Parameters
// 		Callback : sends output back to channel
func updateAsync(updateStruct *Update, callback chan *Callback) {
	records, err := update(updateStruct)
	cb := new(Callback)
	cb.Data = records
	cb.Error = err
	callback <- cb
}

// updateAll : Function Updates all the record into the collection
// Input Parameters
// 		collection (Mgo Object) : Mongo Collection Object
// 		query(Object) : the bson object which holds the update criteria
// 		data(Object) : the object which has to be inserted
// Output Parameters
// 		info : details about the operation affected
// 		error : if it was error then return error else nil

func updateAll(updateAllStruct *UpdateAll) (interface{}, error) {
	records, err := updateAllStruct.Collection.UpdateAll(updateAllStruct.Query, updateAllStruct.Data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return records, nil
}

// updateAllAsync : Function Updates all the record into the collection
// Input Parameters
// 		collection (Mgo Object) : Mongo Collection Object
// 		query(Object) : the bson object which holds the update criteria
// 		data(Object) : the object which has to be inserted
//		callback (channel) : which returns data to goroutine
// Output Parameters
// 		callback : sends output back to channel

func updateAllAsync(updateAllStruct *UpdateAll, callback chan *Callback) {
	records, err := updateAll(updateAllStruct)
	cb := new(Callback)
	cb.Data = records
	cb.Error = err
	callback <- cb
}

//	findByID : Function FindByID finds and returns record by Hexadecimal ID
//	Input Parameters :
//		collection *mgo.Collection : Mongo Collection Object
//		id : bson.ObjectId
//	Output :
//		interface{} : Returns Mongo Objects
func findByID(findByIDStruct *FindByID) (interface{}, error) {
	var records interface{}
	err := findByIDStruct.Collection.FindId(findByIDStruct.Id).One(&records)
	if err != nil {
		log.Fatal(err)
	}
	return records, err
}

//	findByIDAsync : Function FindByIDAsync finds and returns record by Hexadecimal ID
//	Input Parameters :
//		collection *mgo.Collection : Mongo Collection Object
//		id : bson.ObjectId
//		callback (channel) : which returns data to goroutine
//	Output :
//		callback : sends output back to channel
func findByIDAsync(findByIDStruct *FindByID, callback chan *Callback) {
	records, err := findByID(findByIDStruct)
	cb := new(Callback)
	cb.Data = records
	cb.Error = err
	callback <- cb
}

// find : Function finds the record into the collection according to the query/criteria
// Input Parameters
// 		collection (Mgo Object) : Mongo Collection Object
// 		query(Bson Object) : the criteria which has to be evaluated
//		options (map[string]int) : other parameters like limit. skip, sort, etc
// Output Parameters
// 		error : if it was error then return error else nil
// 		boolean : whether operation was successfull or not
func find(findStruct *Find) ([]interface{}, error) {

	var records []interface{}
	var err error

	limit, isLimit := findStruct.Options["limit"]
	skip, isSkip := findStruct.Options["isSkip"]

	if isLimit && isSkip {
		err = findStruct.Collection.Find(findStruct.Query).Skip(skip).Limit(limit).All(&records)
	} else if isLimit {
		err = findStruct.Collection.Find(findStruct.Query).Limit(limit).All(&records)
	} else if isSkip {
		err = findStruct.Collection.Find(findStruct.Query).Skip(skip).All(&records)
	} else {
		err = findStruct.Collection.Find(findStruct.Query).All(&records)
	}

	if err != nil {
		log.Fatal(err)
	}

	return records, err
}

// find : Function finds the record into the collection according to the query/criteria
// Input Parameters
// 		collection (Mgo Object) : Mongo Collection Object
// 		query(Bson Object) : the criteria which has to be evaluated
//		options (map[string]int) : other parameters like limit. skip, sort, etc
//		callback (channel) : which returns data to goroutine
// Output Parameters
// 		callback : returns data to channel
func findAsync(findStruct *Find, callback chan *Callback) {
	records, err := find(findStruct)
	cb := new(Callback)
	cb.Data = records
	cb.Error = err
	callback <- cb
}

// findAll : Function finds all the records into the collection
// Input Parameters
// 		collection (Mgo Object) : Mongo Collection Object
// Output Parameters
// 		records : all the records present in database
// 		error : if it was error then return error else nil
func findAll(findAllStruct *FindAll) ([]interface{}, error) {
	var records []interface{}
	err := findAllStruct.Collection.Find(nil).All(&records)
	if err != nil {
		log.Fatal(err)
	}
	return records, err
}

// findAllAsync : Function finds all the records into the collection
// Input Parameters
// 		collection (Mgo Object) : Mongo Collection Object
// Output Parameters
// 		Callback : returns data to channel
func findAllAsync(findAllStruct *FindAll, callback chan *Callback) {
	records, err := findAll(findAllStruct)
	cb := new(Callback)
	cb.Data = records
	cb.Error = err
	callback <- cb
}

// remove : Function removes the record from the collection as per criteria/query
// Input Parameters
// 		collection (Mgo Object) : Mongo Collection Object
// 		query(Bson Object) : the criteria which has to be evaluated
// Output Parameters
// 		boolean : returns true / false depending on output of operation
// 		error : if it was error then return error else nil
func remove(removeStruct *Remove) (bool, error) {
	err := removeStruct.Collection.Remove(removeStruct.Query)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	return true, err
}

// removeAsync : Function removes the record from the collection as per criteria/query
// Input Parameters
// 		collection (Mgo Object) : Mongo Collection Object
// 		query(Bson Object) : the criteria which has to be evaluated
//		callback (channel) : which returns data to goroutine
// Output Parameters
// 		Callback : returns data to channel
func removeAsync(removeStruct *Remove, callback chan *Callback) {
	records, err := remove(removeStruct)
	cb := new(Callback)
	cb.Data = records
	cb.Error = err
	callback <- cb
}

// removeAll : Function removes all the record from the collection
// Input Parameters
// 		collection (Mgo Object) : Mongo Collection Object
// Output Parameters
// 		records : returns info object as output of operation
// 		error : if it was error then return error else nil
func removeAll(removeAllStruct *RemoveAll) (interface{}, error) {
	records, err := removeAllStruct.Collection.RemoveAll(nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return records, nil
}

// removeAll : Function removes all the record from the collection
// Input Parameters
// 		collection (Mgo Object) : Mongo Collection Object
//		Callback (channel) : which returns data to goroutine
// Output Parameters
// 		Callback : returns data to channel
func removeAllAsync(removeAllStruct *RemoveAll, callback chan *Callback) {
	records, err := removeAll(removeAllStruct)
	cb := new(Callback)
	cb.Data = records
	cb.Error = err
	callback <- cb
}
