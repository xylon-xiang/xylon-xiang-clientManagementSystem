package module

import (
	"clientManagementSystem/config"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
	"time"
)

var (
	StudentStatusCol *mongo.Collection
	StudentCol       *mongo.Collection
)

const (
	STUDENTSTATUS = "studentStatus"
	STUDENTINFO   = "studentInfo"
)

// connect to db and initialize the collection
func init() {

	client, err := mongo.NewClient(options.Client().
		ApplyURI(config.Config.DatabaseConfig.MongoConfig.DBAddress))
	if err != nil {
		log.Fatal(err)
	}

	expTime := time.Duration(config.Config.DatabaseConfig.MongoConfig.TimeOut) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), expTime)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	clientManagementSysDB := client.Database(config.Config.DatabaseConfig.MongoConfig.DBName)
	StudentStatusCol = clientManagementSysDB.
		Collection(config.Config.DatabaseConfig.MongoConfig.ClassCollection)
	StudentCol = clientManagementSysDB.
		Collection(config.Config.DatabaseConfig.MongoConfig.StudentCollection)

}

// trans values not ptr
func Save(colName string, content interface{}) (err error) {

	switch colName {
	case STUDENTSTATUS:
		{
			studentStatuses := content.([]StudentStatus)
			insertData := make([]interface{}, len(studentStatuses))
			for key, value := range studentStatuses {
				insertData[key] = value
			}
			_, err = StudentStatusCol.InsertMany(context.TODO(), insertData)
		}

	case STUDENTINFO:
		{
			studentInfo := content.([]StudentInfo)
			insertData := make([]interface{}, len(studentInfo))
			for key, value := range insertData {
				insertData[key] = value
			}
			_, err = StudentCol.InsertMany(context.TODO(), insertData)
		}

	}

	if err != nil {
		return err
	}

	return nil
}

func UpdateOne(colName string, content interface{}) (err error) {

	opt := options.Update().SetUpsert(true)

	switch colName {
	case STUDENTSTATUS:
		{
			studentStatus := content.(StudentStatus)
			filter := bson.M{
				"student_id":       studentStatus.StudentId,
				"class_name":       studentStatus.ClassName,
				"class_start_date": studentStatus.ClassStartDate,
			}
			update := bson.M{"$set": studentStatus}
			_, err = StudentStatusCol.UpdateOne(context.TODO(), filter, update, opt)

		}
		//
		//case STUDENTINFO:{
		//	studentInfo := content.(StudentInfo)
		//	filter :=
		//}
	}

	if err != nil {
		return err
	}

	return nil
}

// find all filter must contain the studentId, ClassName, ClassStartDate
// match student password should just contain studentId
// the return interface should be transform into StudentStatus or StudentInfo
//
// for example:
//
// filterMap:= map[string]string{
//		"studentId" = "Uxx",
//		"ClassName" = "shixun",
//		"ClassStartDate" = "555555",
// }
func FindOne(colName string, filterMap map[string]string) (result interface{}, err error) {

	switch colName {
	case STUDENTSTATUS:
		{
			var (
				classStartDate int64
				results        StudentStatus
			)
			classStartDate, err = strconv.ParseInt(filterMap["ClassStartDate"], 10, 64)
			filter := bson.M{
				"student_id":       filterMap["StudentId"],
				"class_name":       filterMap["ClassName"],
				"class_start_date": classStartDate,
			}

			err = StudentStatusCol.FindOne(context.TODO(), filter).Decode(&results)
			if err != nil {
				return nil, err
			}

			return results, nil

		}

	case STUDENTINFO:
		{
			var results StudentInfo
			filter := bson.M{
				"student_id": filterMap["StudentId"],
			}

			err := StudentCol.FindOne(context.TODO(), filter).Decode(&results)
			if err != nil {
				return nil, err
			}

			return results, nil
		}

	}

	return result, nil
}

// colValue is just used for search the specific person's all tips, so the value is StudentId
// the return interface should be just transform into [] StudentStatus or [] StudentInfo
func FindAll(colName string, colValue string) (result interface{}, err error) {

	switch colName {
	case STUDENTSTATUS:
		{

			cursor, err := StudentStatusCol.Find(context.TODO(), bson.M{"student_id": colValue})
			if err != nil {
				return nil, err
			}

			var results []StudentStatus
			if err = cursor.All(context.TODO(), &results); err != nil {
				return nil, err
			}

			return results, nil
		}

	case STUDENTINFO:

		cursor, err := StudentCol.Find(context.TODO(), bson.M{})
		if err != nil {
			return nil, err
		}

		var results []StudentInfo
		if err = cursor.All(context.TODO(), &results); err != nil {
			return nil, err
		}

		return results, nil
	}

	return nil, nil
}
